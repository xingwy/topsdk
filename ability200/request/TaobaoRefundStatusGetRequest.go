package request

import (
	"github.com/xingwy/topsdk/ability200/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoRefundStatusGetRequest struct {
	/*
	   入参对象     */
	QueryParam *domain.TaobaoRefundStatusGetRefundQueryByOrderIdRequest `json:"query_param" required:"true" `
}

func (s *TaobaoRefundStatusGetRequest) SetQueryParam(v domain.TaobaoRefundStatusGetRefundQueryByOrderIdRequest) *TaobaoRefundStatusGetRequest {
	s.QueryParam = &v
	return s
}

func (req *TaobaoRefundStatusGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryParam != nil {
		paramMap["query_param"] = util.ConvertStruct(*req.QueryParam)
	}
	return paramMap
}

func (req *TaobaoRefundStatusGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
