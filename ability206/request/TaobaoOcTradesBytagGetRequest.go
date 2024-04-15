package request


type TaobaoOcTradesBytagGetRequest struct {
    /*
        当前页 defalutValue��1    */
    Page  *int64 `json:"page,omitempty" required:"false" `
    /*
        分页大小 defalutValue��50    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        标签名称     */
    TagName  *string `json:"tag_name" required:"true" `
    /*
        标签类型，1官方，2自定义     */
    TagType  *int64 `json:"tag_type" required:"true" `
}

func (s *TaobaoOcTradesBytagGetRequest) SetPage(v int64) *TaobaoOcTradesBytagGetRequest {
    s.Page = &v
    return s
}
func (s *TaobaoOcTradesBytagGetRequest) SetPageSize(v int64) *TaobaoOcTradesBytagGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoOcTradesBytagGetRequest) SetTagName(v string) *TaobaoOcTradesBytagGetRequest {
    s.TagName = &v
    return s
}
func (s *TaobaoOcTradesBytagGetRequest) SetTagType(v int64) *TaobaoOcTradesBytagGetRequest {
    s.TagType = &v
    return s
}

func (req *TaobaoOcTradesBytagGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Page != nil) {
        paramMap["page"] = *req.Page
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.TagName != nil) {
        paramMap["tag_name"] = *req.TagName
    }
    if(req.TagType != nil) {
        paramMap["tag_type"] = *req.TagType
    }
    return paramMap
}

func (req *TaobaoOcTradesBytagGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}