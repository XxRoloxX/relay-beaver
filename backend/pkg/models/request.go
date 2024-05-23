package models

type Request struct {
	Id              string `bson:"_id,omitempty" json:"id,omitempty"`
	Source          string
	Destination     string
	StartTimestamp  int64
	Target          string
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
