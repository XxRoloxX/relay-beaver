package models

import (
	UUID "github.com/google/uuid"
)

type Request struct {
	Id              UUID.UUID
	Source          string
	Destination     Address
	StartTimestamp  int64
	FinishTimestamp int64
	Headers         map[string]string
	Body            string
	Method          string
	ResponseCode    int
}

type Address struct {
	Host string `json:"host" bson:"host"`
	Port int    `json:"port" bson:"port"`
}
