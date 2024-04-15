package request


type TaobaoJushitaJmsTopicsGetRequest struct {
    /*
        卖家nick     */
    Nick  *string `json:"nick" required:"true" `
}

func (s *TaobaoJushitaJmsTopicsGetRequest) SetNick(v string) *TaobaoJushitaJmsTopicsGetRequest {
    s.Nick = &v
    return s
}

func (req *TaobaoJushitaJmsTopicsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    return paramMap
}

func (req *TaobaoJushitaJmsTopicsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}