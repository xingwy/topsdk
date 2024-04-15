package response

import (
	"github.com/xingwy/topsdk/ability236/domain"
)

type CainiaoCloudprintCustomareaUpdateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   result
	*/
	Result domain.CainiaoCloudprintCustomareaUpdateCloudPrintBaseResult `json:"result,omitempty" `
}
