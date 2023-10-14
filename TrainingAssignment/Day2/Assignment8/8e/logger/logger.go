package logger

import (
	"log"
	"os"
)

type LoggerStore struct {
	Product string
	Items   string
	log     *log.Logger
}

func New() LoggerStore {
	return LoggerStore{
		Product: "Brush",
		Items:   "5",
		log:     log.New(os.Stdin, "satyam", 4),
	}
}
