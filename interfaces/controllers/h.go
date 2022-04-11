package controllers

type H struct {
	message string      `json:"message"`
	data    interface{} `json:"data"`
}

func NewH(message string, data interface{}) *H {
	H := new(H)
	H.message = message
	H.data = data
	return H
}
