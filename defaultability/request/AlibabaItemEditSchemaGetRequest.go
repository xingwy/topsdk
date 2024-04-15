package request

import (
	"github.com/xingwy/topsdk/util"
)

type AlibabaItemEditSchemaGetRequest struct {
	/*
	   商品ID     */
	ItemId *int64 `json:"item_id" required:"true" `
	/*
	   业务扩展参数，需与平台约定好     */
	BizType *string `json:"biz_type,omitempty" required:"false" `
	/*
	   制定返回schema中field字段列表，可用于裁剪返回的schema信息。不填则为全部field     */
	Fields *[]string `json:"fields,omitempty" required:"false" `
}

func (s *AlibabaItemEditSchemaGetRequest) SetItemId(v int64) *AlibabaItemEditSchemaGetRequest {
	s.ItemId = &v
	return s
}
func (s *AlibabaItemEditSchemaGetRequest) SetBizType(v string) *AlibabaItemEditSchemaGetRequest {
	s.BizType = &v
	return s
}
func (s *AlibabaItemEditSchemaGetRequest) SetFields(v []string) *AlibabaItemEditSchemaGetRequest {
	s.Fields = &v
	return s
}

func (req *AlibabaItemEditSchemaGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ItemId != nil {
		paramMap["item_id"] = *req.ItemId
	}
	if req.BizType != nil {
		paramMap["biz_type"] = *req.BizType
	}
	if req.Fields != nil {
		paramMap["fields"] = util.ConvertBasicList(*req.Fields)
	}
	return paramMap
}

func (req *AlibabaItemEditSchemaGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
