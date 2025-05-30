package param

type Response struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}
