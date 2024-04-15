package request


type TaobaoSellercenterUserPermissionsGetRequest struct {
    /*
        用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。     */
    Nick  *string `json:"nick" required:"true" `
}

func (s *TaobaoSellercenterUserPermissionsGetRequest) SetNick(v string) *TaobaoSellercenterUserPermissionsGetRequest {
    s.Nick = &v
    return s
}

func (req *TaobaoSellercenterUserPermissionsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    return paramMap
}

func (req *TaobaoSellercenterUserPermissionsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}