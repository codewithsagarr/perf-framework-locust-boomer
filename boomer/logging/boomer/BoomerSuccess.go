package boomer

import "boomer/globals"

func BoomerSuccess(requestType string, name string, responseTime int64, responseLength int64) {
	globals.GlobalBoomer.RecordSuccess(requestType, name, responseTime, responseLength)
}
