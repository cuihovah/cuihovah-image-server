package api

type Return struct {
	Code int `json:"code"`
	Msg string `json:"string"`
	Data interface{} `json:"data"`
}