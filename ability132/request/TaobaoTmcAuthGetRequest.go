package request


type TaobaoTmcAuthGetRequest struct {
    /*
        tmc组名     */
    Group  *string `json:"group,omitempty" required:"false" `
}

func (s *TaobaoTmcAuthGetRequest) SetGroup(v string) *TaobaoTmcAuthGetRequest {
    s.Group = &v
    return s
}

func (req *TaobaoTmcAuthGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Group != nil) {
        paramMap["group"] = *req.Group
    }
    return paramMap
}

func (req *TaobaoTmcAuthGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}