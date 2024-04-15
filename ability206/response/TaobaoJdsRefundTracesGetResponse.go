package response

import (
	"github.com/xingwy/topsdk/ability206/domain"
)

type TaobaoJdsRefundTracesGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   退款跟踪列表
	*/
	Traces []domain.TaobaoJdsRefundTracesGetRefundTrace `json:"traces,omitempty" `
	/*
	   用户在全链路系统中的状态(比如是否存在)
	*/
	UserStatus string `json:"user_status,omitempty" `
}
