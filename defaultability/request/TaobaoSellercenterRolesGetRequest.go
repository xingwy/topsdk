package request


type TaobaoSellercenterRolesGetRequest struct {
    /*
        卖家昵称(只允许查询自己的信息：当前登陆者)     */
    Nick  *string `json:"nick" required:"true" `
}

func (s *TaobaoSellercenterRolesGetRequest) SetNick(v string) *TaobaoSellercenterRolesGetRequest {
    s.Nick = &v
    return s
}

func (req *TaobaoSellercenterRolesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    return paramMap
}

func (req *TaobaoSellercenterRolesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}