package request

type Request struct {
	Method   string   `json:"method"`
	Protocol string   `json:"protocol"`
	Path     string   `json:"path"`
	Headers  []Header `json:"headers"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
