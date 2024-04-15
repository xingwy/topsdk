package request


type TaobaoSellercenterSubusersGetRequest struct {
    /*
        表示卖家昵称     */
    Nick  *string `json:"nick" required:"true" `
}

func (s *TaobaoSellercenterSubusersGetRequest) SetNick(v string) *TaobaoSellercenterSubusersGetRequest {
    s.Nick = &v
    return s
}

func (req *TaobaoSellercenterSubusersGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    return paramMap
}

func (req *TaobaoSellercenterSubusersGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}