package request


type TaobaoSubusersSubaccountSearchRequest struct {
    /*
        主账号登录名     */
    MainNick  *string `json:"main_nick" required:"true" `
    /*
        根据子账号冒号后缀的搜索词，支持中文单字、英文单词 分词规则搜索。该搜索词必传。如果不需要模糊搜索仅需要分页获取子账号列表，请使用taobao.sellercenter.subusers.page接口     */
    FilterKey  *string `json:"filter_key" required:"true" `
    /*
        页大小，大于1小于200 defalutValue��20    */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码，大于等于1 defalutValue��1    */
    PageNum  *int64 `json:"page_num" required:"true" `
}

func (s *TaobaoSubusersSubaccountSearchRequest) SetMainNick(v string) *TaobaoSubusersSubaccountSearchRequest {
    s.MainNick = &v
    return s
}
func (s *TaobaoSubusersSubaccountSearchRequest) SetFilterKey(v string) *TaobaoSubusersSubaccountSearchRequest {
    s.FilterKey = &v
    return s
}
func (s *TaobaoSubusersSubaccountSearchRequest) SetPageSize(v int64) *TaobaoSubusersSubaccountSearchRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoSubusersSubaccountSearchRequest) SetPageNum(v int64) *TaobaoSubusersSubaccountSearchRequest {
    s.PageNum = &v
    return s
}

func (req *TaobaoSubusersSubaccountSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MainNick != nil) {
        paramMap["main_nick"] = *req.MainNick
    }
    if(req.FilterKey != nil) {
        paramMap["filter_key"] = *req.FilterKey
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNum != nil) {
        paramMap["page_num"] = *req.PageNum
    }
    return paramMap
}

func (req *TaobaoSubusersSubaccountSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}