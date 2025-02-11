package boomer

import "boomer/globals"

func BoomerError(requestType string, name string, responseTime int64, exception string) {
	globals.GlobalBoomer.RecordFailure(requestType, name, responseTime, exception)
}
