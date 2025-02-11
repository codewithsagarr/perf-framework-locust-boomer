package http_server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"boomer/logging/boomer"
	"boomer/model"
	"boomer/utils"
)

func GetAusOpen(meta *model.TestMeta) *model.APIReturn {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// spanCtx, span := otel.NewWorkflowSpanStep(meta)
	// if span != nil {
	// 	defer span.End()
	// }
	url := "https://ausopen.com/"

	client := utils.GetHTTPClient(true, false)
	request, err := http.NewRequest("GET", url, nil)
	// request, err := http.NewRequestWithContext(spanCtx, "POST", fmt.Sprintf("http://%s/user/add", utils.GetEnv("HTTP_END_POINT", "localhost:50052")), nil)
	if err != nil {
		boomer.BoomerError("http get", url, 0, err.Error())
		// attributes := map[string]string{
		// 	"event.error_area": "http.NewRequestWithContext",
		// }
		// otel.SetSpanError(span, err, attributes)
		return &model.APIReturn{
			Success: false,
			Elapsed: 0,
			Error:   err,
			Input:   "",
			Output:  "",
		}
	}
	request = utils.AddHeaders(request, headers, "", meta)

	start := time.Now()
	response, err := client.Do(request)
	elapsed := time.Since(start)
	if err != nil {
		boomer.BoomerError("http get", url, elapsed.Milliseconds(), err.Error())
		// attributes := map[string]string{
		// 	"event.error_area": "client.Do",
		// }
		// otel.SetSpanError(span, err, attributes)
		return &model.APIReturn{
			Success: false,
			Elapsed: elapsed,
			Error:   err,
			Input:   "",
			Output:  "",
		}
	}
	if response != nil {
		defer response.Body.Close()
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		boomer.BoomerError("http get", url, elapsed.Milliseconds(), err.Error())
		// attributes := map[string]string{
		// 	"event.error_area": "io.ReadAll",
		// }
		// otel.SetSpanError(span, err, attributes)
		return &model.APIReturn{
			Success: false,
			Elapsed: elapsed,
			Error:   err,
			Input:   "",
			Output:  "",
		}
	}

	switch response.StatusCode {
	case 200:
		boomer.BoomerSuccess("http get", url, elapsed.Milliseconds(), 0)
		return &model.APIReturn{
			Success: true,
			Elapsed: elapsed,
			Error:   nil,
			Input:   "",
			Output:  string(body),
		}

	default:
		boomer.BoomerError("http get", url, elapsed.Milliseconds(), fmt.Sprintf("StatusCode: %d", response.StatusCode))
		// attributes := map[string]string{
		// 	"event.error_area": "response.StatusCode",
		// }
		// otel.SetSpanError(span, fmt.Errorf("StatusCode: %d", response.StatusCode), attributes)
		return &model.APIReturn{
			Success: false,
			Elapsed: elapsed,
			Error:   fmt.Errorf("Status Code: %d", response.StatusCode),
			Input:   "",
			Output:  "",
		}
	}
}
