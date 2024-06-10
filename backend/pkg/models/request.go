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

func (r Response) IsBadRequest() bool {
	return r.StatusCode >= 400 && r.StatusCode < 500
}
func (r Response) IsServerError() bool {
	return r.StatusCode >= 500
}
func (r Response) IsSuccessful() bool {

	return r.StatusCode >= 200 && r.StatusCode < 300
}

type ProxiedRequest struct {
	Id        string   `bson:"_id,omitempty" json:"id,omitempty"`
	Request   Request  `bson:"request" json:"request"`
	Response  Response `bson:"response" json:"response"`
	Target    string   `bson:"target" json:"target"`
	StartTime int      `bson:"startTime" json:"startTime"` // Start time as Unix timestamp
	EndTime   int      `bson:"endTime" json:"endTime"`     // End time as Unix timestamp
}

func (r ProxiedRequest) Host() string {
	for _, header := range r.Request.Headers {
		if header.Key == "Host" {
			return header.Value
		}
	}
	return ""
}

func (r ProxiedRequest) Timestamp() int {
	return r.StartTime
}
func (r ProxiedRequest) Latency() int {
	return r.EndTime - r.StartTime
}

type Address struct {
	Host string `json:"host" bson:"host"`
	Port int    `json:"port" bson:"port"`
}
