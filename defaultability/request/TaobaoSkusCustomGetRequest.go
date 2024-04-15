package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoSkusCustomGetRequest struct {
	/*
	   需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开     */
	Fields *[]string `json:"fields" required:"true" `
	/*
	   Sku的外部商家ID     */
	OuterId *string `json:"outer_id" required:"true" `
}

func (s *TaobaoSkusCustomGetRequest) SetFields(v []string) *TaobaoSkusCustomGetRequest {
	s.Fields = &v
	return s
}
func (s *TaobaoSkusCustomGetRequest) SetOuterId(v string) *TaobaoSkusCustomGetRequest {
	s.OuterId = &v
	return s
}

func (req *TaobaoSkusCustomGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	if req.OuterId != nil {
		paramMap["outer_id"] = *req.OuterId
	}
	return paramMap
}

func (req *TaobaoSkusCustomGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
