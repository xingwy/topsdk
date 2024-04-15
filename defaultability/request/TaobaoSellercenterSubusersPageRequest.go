package request


type TaobaoSellercenterSubusersPageRequest struct {
    /*
        主账号登陆nick     */
    Nick  *string `json:"nick" required:"true" `
    /*
        可以不传，不传的话和老接口返回数据一样。如果传则根据子账号当前状态筛选 1正常   2冻结  3处罚，如果不传默认都返回     */
    AccountStatus  *int64 `json:"account_status,omitempty" required:"false" `
    /*
        页大小，大于1小于200 defalutValue��100    */
    PageSize  *int64 `json:"page_size" required:"true" `
    /*
        页码，大于等于1 defalutValue��1    */
    PageNum  *int64 `json:"page_num" required:"true" `
}

func (s *TaobaoSellercenterSubusersPageRequest) SetNick(v string) *TaobaoSellercenterSubusersPageRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoSellercenterSubusersPageRequest) SetAccountStatus(v int64) *TaobaoSellercenterSubusersPageRequest {
    s.AccountStatus = &v
    return s
}
func (s *TaobaoSellercenterSubusersPageRequest) SetPageSize(v int64) *TaobaoSellercenterSubusersPageRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoSellercenterSubusersPageRequest) SetPageNum(v int64) *TaobaoSellercenterSubusersPageRequest {
    s.PageNum = &v
    return s
}

func (req *TaobaoSellercenterSubusersPageRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.AccountStatus != nil) {
        paramMap["account_status"] = *req.AccountStatus
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNum != nil) {
        paramMap["page_num"] = *req.PageNum
    }
    return paramMap
}

func (req *TaobaoSellercenterSubusersPageRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}