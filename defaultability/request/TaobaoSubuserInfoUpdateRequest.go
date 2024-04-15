package request


type TaobaoSubuserInfoUpdateRequest struct {
    /*
        子账号ID     */
    SubId  *int64 `json:"sub_id" required:"true" `
    /*
        子账号是否参与分流 true:参与分流 false:不参与分流     */
    IsDispatch  *bool `json:"is_dispatch,omitempty" required:"false" `
    /*
        是否停用子账号 true:表示停用该子账号false:表示开启该子账号     */
    IsDisableSubaccount  *bool `json:"is_disable_subaccount,omitempty" required:"false" `
}

func (s *TaobaoSubuserInfoUpdateRequest) SetSubId(v int64) *TaobaoSubuserInfoUpdateRequest {
    s.SubId = &v
    return s
}
func (s *TaobaoSubuserInfoUpdateRequest) SetIsDispatch(v bool) *TaobaoSubuserInfoUpdateRequest {
    s.IsDispatch = &v
    return s
}
func (s *TaobaoSubuserInfoUpdateRequest) SetIsDisableSubaccount(v bool) *TaobaoSubuserInfoUpdateRequest {
    s.IsDisableSubaccount = &v
    return s
}

func (req *TaobaoSubuserInfoUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SubId != nil) {
        paramMap["sub_id"] = *req.SubId
    }
    if(req.IsDispatch != nil) {
        paramMap["is_dispatch"] = *req.IsDispatch
    }
    if(req.IsDisableSubaccount != nil) {
        paramMap["is_disable_subaccount"] = *req.IsDisableSubaccount
    }
    return paramMap
}

func (req *TaobaoSubuserInfoUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}