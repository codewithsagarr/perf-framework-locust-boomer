package influxdb

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"boomer/model"
	"boomer/utils"
)

var influxdbAddress = fmt.Sprintf("http://%s/api/v2/write?bucket=%s&org=%s",
	utils.GetEnv("INFLUXDB_ADDRESS", "localhost:8086"),
	utils.GetEnv("INFLUXDB_BUCKET", "my-bucket"),
	utils.GetEnv("INFLUXDB_ORG", "my-org"),
)
var ticker *time.Ticker
var logChan chan string
var closeSignalChan chan int
var channelOpen = false

func InitLogging() {
	if !channelOpen {
		logChan = make(chan string, 1000)
		closeSignalChan = make(chan int)
		ticker = time.NewTicker(1 * time.Second)

		go logReceiver(logChan, closeSignalChan)
		channelOpen = true
	}
}

func CloseLogging() {
	if channelOpen {
		closeSignalChan <- 1

		close(logChan)
		close(closeSignalChan)
		ticker.Stop()
		channelOpen = false
	}
}

func logReceiver(inMessage <-chan string, inClose <-chan int) {
	var messages strings.Builder

	for {
		select {
		case <-ticker.C:
			if messages.Len() > 0 {
				logBatchWriter(messages)
				messages.Reset()
			}

		case message := <-inMessage:
			messages.WriteString(message)

		case <-inClose:
			log.Printf("Influxdb Closing channel: %d\n", messages.Len())
			if messages.Len() > 0 {
				logBatchWriter(messages)
			}
			messages.Reset()
		}
	}
}

func logBatchWriter(messages strings.Builder) {
	client := http.Client{
		Timeout: time.Second * 120,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	defer client.CloseIdleConnections()

	request, err := http.NewRequest("POST", influxdbAddress, strings.NewReader(messages.String()))
	if err != nil {
		fmt.Printf("influxdb logBatchWriter new request error: %s\n", err)
		return
	}
	request.Header.Set("Content-Type", "text/plain")
	// TODO - need to environment variable this
	request.Header.Set("Authorization", "Token my-super-secret-auth-token")

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("influxdb logBatchWriter new response error: %s\n", err)
		return
	}
	if response != nil {
		defer response.Body.Close()
	}
}

func Log(
	meta *model.TestMeta,
	responseCode string,
	success bool,
	elapsed time.Duration,
	responseSize int,
	timestamp time.Time,
) {
	message := InfluxMessage{
		measure:   meta.Workflow,
		timestamp: timestamp,
	}
	message.AddTag("run_id", meta.RunId)
	message.AddTag("summary", strconv.FormatBool(meta.Summary))
	message.AddTag("step", fmt.Sprintf("%03d", meta.Step))
	message.AddTag("label", meta.Label)
	message.AddTag("response_code", responseCode)
	message.AddTag("success", strconv.FormatBool(success))
	if meta.Summary {
		message.AddTag("total_elapsed_time", fmt.Sprintf("%d", meta.TotalElapsedTime.Milliseconds()))
	}

	message.AddField("elapsed", fmt.Sprintf("%d", elapsed.Milliseconds()))

	if responseSize > -1 {
		message.AddField("response_size", fmt.Sprintf("%d", responseSize))
	}

	logChan <- message.LineProtocol()
}
