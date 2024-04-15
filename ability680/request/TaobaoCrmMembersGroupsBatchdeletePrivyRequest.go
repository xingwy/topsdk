package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoCrmMembersGroupsBatchdeletePrivyRequest struct {
	/*
	   会员需要删除的分组     */
	GroupIds *[]int64 `json:"group_ids" required:"true" `
	/*
	   ouid列表     */
	Ouids *[]string `json:"ouids,omitempty" required:"false" `
}

func (s *TaobaoCrmMembersGroupsBatchdeletePrivyRequest) SetGroupIds(v []int64) *TaobaoCrmMembersGroupsBatchdeletePrivyRequest {
	s.GroupIds = &v
	return s
}
func (s *TaobaoCrmMembersGroupsBatchdeletePrivyRequest) SetOuids(v []string) *TaobaoCrmMembersGroupsBatchdeletePrivyRequest {
	s.Ouids = &v
	return s
}

func (req *TaobaoCrmMembersGroupsBatchdeletePrivyRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.GroupIds != nil {
		paramMap["group_ids"] = util.ConvertBasicList(*req.GroupIds)
	}
	if req.Ouids != nil {
		paramMap["ouids"] = util.ConvertBasicList(*req.Ouids)
	}
	return paramMap
}

func (req *TaobaoCrmMembersGroupsBatchdeletePrivyRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
