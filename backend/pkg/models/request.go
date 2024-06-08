package models

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Request struct {
	Method   string   `json:"method"`
	Protocol string   `json:"protocol"`
	Path     string   `json:"path"`
	Headers  []Header `json:"headers"`
	Body     string   `json:"body"`
}

type Response struct {
	StatusCode int      `json:"statusCode"`
	Protocol   string   `json:"protocol"`
	Headers    []Header `json:"headers"`
	Body       string   `json:"body"`
}

type ProxiedRequest struct {
	Id        string   `bson:"_id,omitempty" json:"id,omitempty"`
	Request   Request  `bson:"request" json:"request"`
	Response  Response `bson:"response" json:"response"`
	Target    string   `bson:"target" json:"target"`
	StartTime int64    `bson:"startTime" json:"startTime"` // Start time as Unix timestamp
	EndTime   int64    `bson:"endTime" json:"endTime"`     // End time as Unix timestamp
}

type Address struct {
	Host string `json:"host" bson:"host"`
	Port int    `json:"port" bson:"port"`
}
