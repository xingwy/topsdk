package request


type TaobaoCrmGroupUpdateRequest struct {
    /*
        分组的id     */
    GroupId  *int64 `json:"group_id" required:"true" `
    /*
        新的分组名，分组名称不能包含|或者：     */
    NewGroupName  *string `json:"new_group_name" required:"true" `
}

func (s *TaobaoCrmGroupUpdateRequest) SetGroupId(v int64) *TaobaoCrmGroupUpdateRequest {
    s.GroupId = &v
    return s
}
func (s *TaobaoCrmGroupUpdateRequest) SetNewGroupName(v string) *TaobaoCrmGroupUpdateRequest {
    s.NewGroupName = &v
    return s
}

func (req *TaobaoCrmGroupUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupId != nil) {
        paramMap["group_id"] = *req.GroupId
    }
    if(req.NewGroupName != nil) {
        paramMap["new_group_name"] = *req.NewGroupName
    }
    return paramMap
}

func (req *TaobaoCrmGroupUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}