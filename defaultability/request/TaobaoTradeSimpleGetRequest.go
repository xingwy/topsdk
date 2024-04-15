package request

import (
        "topsdk/util"
    )

type TaobaoTradeSimpleGetRequest struct {
    /*
        需要返回的字段列表，多个字段用半角逗号分隔，可选值为返回示例中能看到的所有字段。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        交易编号     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoTradeSimpleGetRequest) SetFields(v []string) *TaobaoTradeSimpleGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTradeSimpleGetRequest) SetTid(v int64) *TaobaoTradeSimpleGetRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoTradeSimpleGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoTradeSimpleGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}