package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoItemsSellerListGetRequest struct {
	/*
	   需要返回的商品字段列表。可选值：点击返回结果中的Item结构体中能展示出来的所有字段，多个字段用“,”分隔。注：返回所有sku信息的字段名称是sku而不是skus。     */
	Fields *[]string `json:"fields" required:"true" `
	/*
	   商品ID列表，多个ID用半角逗号隔开，一次最多不超过20个。注：获取不存在的商品ID或获取别人的商品都不会报错，但没有商品数据返回。     */
	NumIids *[]string `json:"num_iids" required:"true" `
}

func (s *TaobaoItemsSellerListGetRequest) SetFields(v []string) *TaobaoItemsSellerListGetRequest {
	s.Fields = &v
	return s
}
func (s *TaobaoItemsSellerListGetRequest) SetNumIids(v []string) *TaobaoItemsSellerListGetRequest {
	s.NumIids = &v
	return s
}

func (req *TaobaoItemsSellerListGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	if req.NumIids != nil {
		paramMap["num_iids"] = util.ConvertBasicList(*req.NumIids)
	}
	return paramMap
}

func (req *TaobaoItemsSellerListGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
