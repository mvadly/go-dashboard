package models

type ResponseBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ReponseServices struct {
	Code     int
	Message  string
	Response interface{}
}
