package request

import (
        "topsdk/util"
    )

type TaobaoOcTradetagsGetRequest struct {
    /*
        是否查询历史标签     */
    History  *int64 `json:"history,omitempty" required:"false" `
    /*
        交易主订单id     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        不填，则默认只查询1,2。1为官方标，2为自定义标，3为主站只读标签     */
    TagTypes  *[]string `json:"tag_types,omitempty" required:"false" `
    /*
        不填，则不做标签名称过滤     */
    TagNames  *[]string `json:"tag_names,omitempty" required:"false" `
}

func (s *TaobaoOcTradetagsGetRequest) SetHistory(v int64) *TaobaoOcTradetagsGetRequest {
    s.History = &v
    return s
}
func (s *TaobaoOcTradetagsGetRequest) SetTid(v int64) *TaobaoOcTradetagsGetRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoOcTradetagsGetRequest) SetTagTypes(v []string) *TaobaoOcTradetagsGetRequest {
    s.TagTypes = &v
    return s
}
func (s *TaobaoOcTradetagsGetRequest) SetTagNames(v []string) *TaobaoOcTradetagsGetRequest {
    s.TagNames = &v
    return s
}

func (req *TaobaoOcTradetagsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.History != nil) {
        paramMap["history"] = *req.History
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.TagTypes != nil) {
        paramMap["tag_types"] = util.ConvertBasicList(*req.TagTypes)
    }
    if(req.TagNames != nil) {
        paramMap["tag_names"] = util.ConvertBasicList(*req.TagNames)
    }
    return paramMap
}

func (req *TaobaoOcTradetagsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}