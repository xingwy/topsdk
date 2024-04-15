package request


type TaobaoJushitaJmsUserGetRequest struct {
    /*
        需要查询的用户名     */
    UserNick  *string `json:"user_nick" required:"true" `
}

func (s *TaobaoJushitaJmsUserGetRequest) SetUserNick(v string) *TaobaoJushitaJmsUserGetRequest {
    s.UserNick = &v
    return s
}

func (req *TaobaoJushitaJmsUserGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UserNick != nil) {
        paramMap["user_nick"] = *req.UserNick
    }
    return paramMap
}

func (req *TaobaoJushitaJmsUserGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}