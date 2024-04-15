package request


type TaobaoWlbStoresBaseinfoGetRequest struct {
    /*
        0.商家仓库.1.菜鸟仓库.2全部被划分的仓库 defalutValue��0    */
    Type  *int64 `json:"type,omitempty" required:"false" `
}

func (s *TaobaoWlbStoresBaseinfoGetRequest) SetType(v int64) *TaobaoWlbStoresBaseinfoGetRequest {
    s.Type = &v
    return s
}

func (req *TaobaoWlbStoresBaseinfoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoWlbStoresBaseinfoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}