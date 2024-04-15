package response

import (
	"github.com/xingwy/topsdk/ability147/domain"
)

type TmallItemHscodeAuditResultsQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   商品或sku的hscode信息审核状态。
	*/
	ResultList []domain.TmallItemHscodeAuditResultsQueryHscodeAuditInfo `json:"result_list,omitempty" `
}
