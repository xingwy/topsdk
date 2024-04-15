package request


type TaobaoCrmGroupDeleteRequest struct {
    /*
        要删除的分组id     */
    GroupId  *int64 `json:"group_id" required:"true" `
}

func (s *TaobaoCrmGroupDeleteRequest) SetGroupId(v int64) *TaobaoCrmGroupDeleteRequest {
    s.GroupId = &v
    return s
}

func (req *TaobaoCrmGroupDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupId != nil) {
        paramMap["group_id"] = *req.GroupId
    }
    return paramMap
}

func (req *TaobaoCrmGroupDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}