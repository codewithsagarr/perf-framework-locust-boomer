package utils

import (
	"net/http"

	"boomer/model"
)

func AddHeaders(request *http.Request, headers map[string]string, session string, meta *model.TestMeta) *http.Request {
	if session != "" {
		// request.Header.Set("cookie", fmt.Sprintf("session=%s", session))
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}
	return request
}
