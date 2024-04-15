package request


type TaobaoFuwuSaleLinkGenRequest struct {
    /*
        用户需要营销的目标人群中的用户nick     */
    Nick  *string `json:"nick,omitempty" required:"false" `
    /*
        从服务商后台，营销链接功能中生成的参数串直接复制使用。不要修改，否则抛错。     */
    ParamStr  *string `json:"param_str" required:"true" `
}

func (s *TaobaoFuwuSaleLinkGenRequest) SetNick(v string) *TaobaoFuwuSaleLinkGenRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoFuwuSaleLinkGenRequest) SetParamStr(v string) *TaobaoFuwuSaleLinkGenRequest {
    s.ParamStr = &v
    return s
}

func (req *TaobaoFuwuSaleLinkGenRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.ParamStr != nil) {
        paramMap["param_str"] = *req.ParamStr
    }
    return paramMap
}

func (req *TaobaoFuwuSaleLinkGenRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}