package request


type TmallItemHscodeDetailGetRequest struct {
    /*
        hscode     */
    Hscode  *string `json:"hscode" required:"true" `
}

func (s *TmallItemHscodeDetailGetRequest) SetHscode(v string) *TmallItemHscodeDetailGetRequest {
    s.Hscode = &v
    return s
}

func (req *TmallItemHscodeDetailGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Hscode != nil) {
        paramMap["hscode"] = *req.Hscode
    }
    return paramMap
}

func (req *TmallItemHscodeDetailGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}