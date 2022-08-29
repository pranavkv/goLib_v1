package LibUtils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	data "github.com/pranavkv/golib_v1/LibData"
  )


var Logger StandardLogger

type StandardLogger struct {
	*log.Entry
}

func (l *StandardLogger) LogRequest(req data.GoLibRequest) {
	resBody, _ := json.Marshal(req)
	l.Info("Request Received: ", resBody)
  }
 
func InitLog(serviceName string, hostName string) *StandardLogger {

	var baseLogger = log.New()
	baseLogger.Formatter = &log.JSONFormatter{}

	childLogger := baseLogger.WithFields(log.Fields{
		"service": serviceName,
		"host": hostName,
	  })

	var standardLogger = &StandardLogger{childLogger}
	Logger = *standardLogger

	return standardLogger
}
