package request


type TaobaoCrmGrouptaskCheckRequest struct {
    /*
        分组id     */
    GroupId  *int64 `json:"group_id" required:"true" `
}

func (s *TaobaoCrmGrouptaskCheckRequest) SetGroupId(v int64) *TaobaoCrmGrouptaskCheckRequest {
    s.GroupId = &v
    return s
}

func (req *TaobaoCrmGrouptaskCheckRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupId != nil) {
        paramMap["group_id"] = *req.GroupId
    }
    return paramMap
}

func (req *TaobaoCrmGrouptaskCheckRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}