package request

import (
        "topsdk/util"
    )

type TaobaoCrmMembersGroupBatchaddPrivyRequest struct {
    /*
        分组id     */
    GroupIds  *[]int64 `json:"group_ids" required:"true" `
    /*
        ouid列表     */
    Ouids  *[]string `json:"ouids" required:"true" `
}

func (s *TaobaoCrmMembersGroupBatchaddPrivyRequest) SetGroupIds(v []int64) *TaobaoCrmMembersGroupBatchaddPrivyRequest {
    s.GroupIds = &v
    return s
}
func (s *TaobaoCrmMembersGroupBatchaddPrivyRequest) SetOuids(v []string) *TaobaoCrmMembersGroupBatchaddPrivyRequest {
    s.Ouids = &v
    return s
}

func (req *TaobaoCrmMembersGroupBatchaddPrivyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupIds != nil) {
        paramMap["group_ids"] = util.ConvertBasicList(*req.GroupIds)
    }
    if(req.Ouids != nil) {
        paramMap["ouids"] = util.ConvertBasicList(*req.Ouids)
    }
    return paramMap
}

func (req *TaobaoCrmMembersGroupBatchaddPrivyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}