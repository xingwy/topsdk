package request


type TaobaoSubusersPageRequest struct {
    /*
        主账号用户名     */
    UserNick  *string `json:"user_nick" required:"true" `
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

func (s *TaobaoSubusersPageRequest) SetUserNick(v string) *TaobaoSubusersPageRequest {
    s.UserNick = &v
    return s
}
func (s *TaobaoSubusersPageRequest) SetAccountStatus(v int64) *TaobaoSubusersPageRequest {
    s.AccountStatus = &v
    return s
}
func (s *TaobaoSubusersPageRequest) SetPageSize(v int64) *TaobaoSubusersPageRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoSubusersPageRequest) SetPageNum(v int64) *TaobaoSubusersPageRequest {
    s.PageNum = &v
    return s
}

func (req *TaobaoSubusersPageRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UserNick != nil) {
        paramMap["user_nick"] = *req.UserNick
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

func (req *TaobaoSubusersPageRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}