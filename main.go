package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	prepare(logrus.TraceLevel)
	logrus.Infof("hello")
}

func prepare(logLevel logrus.Level) {
	customFormatter := logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.999999999 -0700"
	logrus.SetFormatter(&customFormatter)
	logrus.SetReportCaller(true)
	logrus.SetLevel(logLevel)
}
