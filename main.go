package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TODO: These should be read from environment variables
var broker = "10.0.0.228"
var port = 9002

var Zap = zap.NewLogger()

type MqttHandler struct {
	Client *mqtt.Client
}

var MqttClient = MqttInit()
