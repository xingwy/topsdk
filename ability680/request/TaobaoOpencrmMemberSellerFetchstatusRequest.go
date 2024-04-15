package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoOpencrmMemberSellerFetchstatusRequest struct {
	/*
	   商家nick列表     */
	SellerNicks *[]string `json:"seller_nicks,omitempty" required:"false" `
}

func (s *TaobaoOpencrmMemberSellerFetchstatusRequest) SetSellerNicks(v []string) *TaobaoOpencrmMemberSellerFetchstatusRequest {
	s.SellerNicks = &v
	return s
}

func (req *TaobaoOpencrmMemberSellerFetchstatusRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SellerNicks != nil {
		paramMap["seller_nicks"] = util.ConvertBasicList(*req.SellerNicks)
	}
	return paramMap
}

func (req *TaobaoOpencrmMemberSellerFetchstatusRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
