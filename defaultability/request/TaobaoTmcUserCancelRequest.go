package request


type TaobaoTmcUserCancelRequest struct {
    /*
        用户昵称     */
    Nick  *string `json:"nick" required:"true" `
    /*
        用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 defalutValue��tbUIC    */
    UserPlatform  *string `json:"user_platform,omitempty" required:"false" `
}

func (s *TaobaoTmcUserCancelRequest) SetNick(v string) *TaobaoTmcUserCancelRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoTmcUserCancelRequest) SetUserPlatform(v string) *TaobaoTmcUserCancelRequest {
    s.UserPlatform = &v
    return s
}

func (req *TaobaoTmcUserCancelRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.UserPlatform != nil) {
        paramMap["user_platform"] = *req.UserPlatform
    }
    return paramMap
}

func (req *TaobaoTmcUserCancelRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}