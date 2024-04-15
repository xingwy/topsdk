package request


type TaobaoSubuserDepartmentsGetRequest struct {
    /*
        主账号用户名     */
    UserNick  *string `json:"user_nick" required:"true" `
}

func (s *TaobaoSubuserDepartmentsGetRequest) SetUserNick(v string) *TaobaoSubuserDepartmentsGetRequest {
    s.UserNick = &v
    return s
}

func (req *TaobaoSubuserDepartmentsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.UserNick != nil) {
        paramMap["user_nick"] = *req.UserNick
    }
    return paramMap
}

func (req *TaobaoSubuserDepartmentsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}