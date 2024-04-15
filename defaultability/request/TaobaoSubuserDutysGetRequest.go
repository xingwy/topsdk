package request


type TaobaoSubuserDutysGetRequest struct {
    /*
        主账号用户名     */
    UserNick  *string `json:"user_nick" required:"true" `
}

func (s *TaobaoSubuserDutysGetRequest) SetUserNick(v string) *TaobaoSubuserDutysGetRequest {
    s.UserNick = &v
    return s
}

func (req *TaobaoSubuserDutysGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UserNick != nil) {
        paramMap["user_nick"] = *req.UserNick
    }
    return paramMap
}

func (req *TaobaoSubuserDutysGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}