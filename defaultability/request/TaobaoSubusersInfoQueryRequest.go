package request


type TaobaoSubusersInfoQueryRequest struct {
    /*
        站点信息，淘宝天猫传0，1688传3 defalutValue��0    */
    Site  *int64 `json:"site" required:"true" `
}

func (s *TaobaoSubusersInfoQueryRequest) SetSite(v int64) *TaobaoSubusersInfoQueryRequest {
    s.Site = &v
    return s
}

func (req *TaobaoSubusersInfoQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Site != nil) {
        paramMap["site"] = *req.Site
    }
    return paramMap
}

func (req *TaobaoSubusersInfoQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}