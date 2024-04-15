package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoSubuserFullinfoGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息
	*/
	SubFullinfo domain.TaobaoSubuserFullinfoGetSubUserFullInfo `json:"sub_fullinfo,omitempty" `
}
