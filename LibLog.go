/*
Author: Pranav KV
Mail: pranavkvnambiar@gmail.com
*/
package golib_v1

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

var Logger StandardLogger

type StandardLogger struct {
	*log.Entry
}

func (l *StandardLogger) LogRequest(req GoLibRequest) {
	resBody, _ := json.Marshal(req)
	l.Info("Request Received: ", resBody)
}

func InitLog(serviceName string, hostName string) *StandardLogger {

	var baseLogger = log.New()
	baseLogger.Formatter = &log.JSONFormatter{}

	childLogger := baseLogger.WithFields(log.Fields{
		"service": serviceName,
		"host":    hostName,
	})

	var standardLogger = &StandardLogger{childLogger}
	Logger = *standardLogger

	return standardLogger
}
