package request


type TaobaoSubusersGetRequest struct {
    /*
        主账号用户名     */
    UserNick  *string `json:"user_nick" required:"true" `
}

func (s *TaobaoSubusersGetRequest) SetUserNick(v string) *TaobaoSubusersGetRequest {
    s.UserNick = &v
    return s
}

func (req *TaobaoSubusersGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UserNick != nil) {
        paramMap["user_nick"] = *req.UserNick
    }
    return paramMap
}

func (req *TaobaoSubusersGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}