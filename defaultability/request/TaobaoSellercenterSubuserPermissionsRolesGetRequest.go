package request


type TaobaoSellercenterSubuserPermissionsRolesGetRequest struct {
    /*
        子账号昵称(子账号标识)     */
    Nick  *string `json:"nick" required:"true" `
}

func (s *TaobaoSellercenterSubuserPermissionsRolesGetRequest) SetNick(v string) *TaobaoSellercenterSubuserPermissionsRolesGetRequest {
    s.Nick = &v
    return s
}

func (req *TaobaoSellercenterSubuserPermissionsRolesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    return paramMap
}

func (req *TaobaoSellercenterSubuserPermissionsRolesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}