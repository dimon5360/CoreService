package logger

import (
	"app/main/kafka"
	"app/main/utils"
)

type LoggerCore struct {
	handler *kafka.Handler
}

func Init(configPath string) *LoggerCore {
	var logger LoggerCore
	utils.ParseJsonConfig("config/kafka.json", &logger)

	var kafkaHandler = kafka.InitKafkaHandler(configPath)

	kafkaHandler.StartKafkaHandler()
	logger.handler = kafkaHandler

	return &logger
}

func (logger *LoggerCore) Write(record string) {
	logger.handler.Produce(record)
}
