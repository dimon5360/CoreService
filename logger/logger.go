package logger

import "app/main/utils"

type LoggerCore struct {
	LoggingServiceHost    string `json:"host"`
	LoggingServiceTopic   string `json:"producerTopic"`
	LoggingServiceGroudId string `json:"group.id"`
}

func Init(configPath string) *LoggerCore {
	var logger LoggerCore
	utils.ParseJsonConfig("config/kafka.json", &logger)

	// todo: init kafka producer
	return &logger
}

func (logger *LoggerCore) Write(record string) {

}
