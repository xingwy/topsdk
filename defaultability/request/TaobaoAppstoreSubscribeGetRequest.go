package request


type TaobaoAppstoreSubscribeGetRequest struct {
    /*
        用户昵称     */
    Nick  *string `json:"nick" required:"true" `
    /*
        插件实例ID defalutValue��0    */
    LeaseId  *int64 `json:"lease_id,omitempty" required:"false" `
}

func (s *TaobaoAppstoreSubscribeGetRequest) SetNick(v string) *TaobaoAppstoreSubscribeGetRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoAppstoreSubscribeGetRequest) SetLeaseId(v int64) *TaobaoAppstoreSubscribeGetRequest {
    s.LeaseId = &v
    return s
}

func (req *TaobaoAppstoreSubscribeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.LeaseId != nil) {
        paramMap["lease_id"] = *req.LeaseId
    }
    return paramMap
}

func (req *TaobaoAppstoreSubscribeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}