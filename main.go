package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.SugaredLogger
}

func NewLogger() Logger {
	logger := zap.NewExample().Sugar()
	return Logger{
		Logger: logger,
	}
}

// TODO: These should be read from environment variables
var broker = "10.0.0.228"
var port = 9002

var Zap = zap.NewLogger()

type MqttHandler struct {
	Client *mqtt.Client
}

var MqttClient = MqttInit()
