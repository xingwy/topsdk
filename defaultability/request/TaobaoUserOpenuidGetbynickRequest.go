package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoUserOpenuidGetbynickRequest struct {
	/*
	   买家nick列表     */
	BuyerNicks *[]string `json:"buyer_nicks" required:"true" `
}

func (s *TaobaoUserOpenuidGetbynickRequest) SetBuyerNicks(v []string) *TaobaoUserOpenuidGetbynickRequest {
	s.BuyerNicks = &v
	return s
}

func (req *TaobaoUserOpenuidGetbynickRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.BuyerNicks != nil {
		paramMap["buyer_nicks"] = util.ConvertBasicList(*req.BuyerNicks)
	}
	return paramMap
}

func (req *TaobaoUserOpenuidGetbynickRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
