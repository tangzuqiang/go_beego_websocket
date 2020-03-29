package data

type Format struct {
	Status  interface{} `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
