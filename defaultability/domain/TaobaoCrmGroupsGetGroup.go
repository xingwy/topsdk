package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoCrmGroupsGetGroup struct {
	/*
	   分组的名称     */
	GroupName *string `json:"group_name,omitempty" `

	/*
	   分组所拥有的会员数量,如果返回值为-1，表示当前服务忙或数据在初始化。     */
	MemberCount *int64 `json:"member_count,omitempty" `

	/*
	   分组的状态，1表示正常     */
	Status *string `json:"status,omitempty" `

	/*
	   分组的id     */
	GroupId *int64 `json:"group_id,omitempty" `

	/*
	   分组的修改时间     */
	GroupModify *util.LocalTime `json:"group_modify,omitempty" `

	/*
	   分组的创建时间     */
	GroupCreate *util.LocalTime `json:"group_create,omitempty" `
}

func (s *TaobaoCrmGroupsGetGroup) SetGroupName(v string) *TaobaoCrmGroupsGetGroup {
	s.GroupName = &v
	return s
}
func (s *TaobaoCrmGroupsGetGroup) SetMemberCount(v int64) *TaobaoCrmGroupsGetGroup {
	s.MemberCount = &v
	return s
}
func (s *TaobaoCrmGroupsGetGroup) SetStatus(v string) *TaobaoCrmGroupsGetGroup {
	s.Status = &v
	return s
}
func (s *TaobaoCrmGroupsGetGroup) SetGroupId(v int64) *TaobaoCrmGroupsGetGroup {
	s.GroupId = &v
	return s
}
func (s *TaobaoCrmGroupsGetGroup) SetGroupModify(v util.LocalTime) *TaobaoCrmGroupsGetGroup {
	s.GroupModify = &v
	return s
}
func (s *TaobaoCrmGroupsGetGroup) SetGroupCreate(v util.LocalTime) *TaobaoCrmGroupsGetGroup {
	s.GroupCreate = &v
	return s
}
