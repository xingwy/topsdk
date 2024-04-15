package response

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoQimenTradeUserAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   创建时间
	*/
	GmtCreate util.LocalTime `json:"gmt_create,omitempty" `
	/*
	   appkey
	*/
	Appkey string `json:"appkey,omitempty" `
	/*
	   卖家备注
	*/
	Memo string `json:"memo,omitempty" `
}
