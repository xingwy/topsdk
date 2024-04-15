package request


type TaobaoOcTradetagAttachRequest struct {
    /*
        标签名称     */
    TagName  *string `json:"tag_name" required:"true" `
    /*
        标签类型       1：官方标签      2：自定义标签 defalutValue��2    */
    TagType  *int64 `json:"tag_type,omitempty" required:"false" `
    /*
        标签值，json格式     */
    TagValue  *string `json:"tag_value" required:"true" `
    /*
        订单id     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        该标签在消费者端是否显示,0:不显示,1：显示 defalutValue��0    */
    Visible  *int64 `json:"visible,omitempty" required:"false" `
}

func (s *TaobaoOcTradetagAttachRequest) SetTagName(v string) *TaobaoOcTradetagAttachRequest {
    s.TagName = &v
    return s
}
func (s *TaobaoOcTradetagAttachRequest) SetTagType(v int64) *TaobaoOcTradetagAttachRequest {
    s.TagType = &v
    return s
}
func (s *TaobaoOcTradetagAttachRequest) SetTagValue(v string) *TaobaoOcTradetagAttachRequest {
    s.TagValue = &v
    return s
}
func (s *TaobaoOcTradetagAttachRequest) SetTid(v int64) *TaobaoOcTradetagAttachRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoOcTradetagAttachRequest) SetVisible(v int64) *TaobaoOcTradetagAttachRequest {
    s.Visible = &v
    return s
}

func (req *TaobaoOcTradetagAttachRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TagName != nil) {
        paramMap["tag_name"] = *req.TagName
    }
    if(req.TagType != nil) {
        paramMap["tag_type"] = *req.TagType
    }
    if(req.TagValue != nil) {
        paramMap["tag_value"] = *req.TagValue
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.Visible != nil) {
        paramMap["visible"] = *req.Visible
    }
    return paramMap
}

func (req *TaobaoOcTradetagAttachRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}