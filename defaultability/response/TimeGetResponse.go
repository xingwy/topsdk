package response

type TimeGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   是否成功
	*/
	Time string `json:"time,omitempty" `
}
