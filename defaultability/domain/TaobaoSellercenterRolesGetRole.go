package domain

import (
        "topsdk/util"
    )

type TaobaoSellercenterRolesGetRole struct {
    /*
        角色id     */
    RoleId  *int64 `json:"role_id,omitempty" `

    /*
        角色描述     */
    Description  *string `json:"description,omitempty" `

    /*
        修改时间     */
    ModifiedTime  *util.LocalTime `json:"modified_time,omitempty" `

    /*
        创建时间     */
    CreateTime  *util.LocalTime `json:"create_time,omitempty" `

    /*
        卖家Id     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        角色名     */
    RoleName  *string `json:"role_name,omitempty" `

}

func (s *TaobaoSellercenterRolesGetRole) SetRoleId(v int64) *TaobaoSellercenterRolesGetRole {
    s.RoleId = &v
    return s
}
func (s *TaobaoSellercenterRolesGetRole) SetDescription(v string) *TaobaoSellercenterRolesGetRole {
    s.Description = &v
    return s
}
func (s *TaobaoSellercenterRolesGetRole) SetModifiedTime(v util.LocalTime) *TaobaoSellercenterRolesGetRole {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoSellercenterRolesGetRole) SetCreateTime(v util.LocalTime) *TaobaoSellercenterRolesGetRole {
    s.CreateTime = &v
    return s
}
func (s *TaobaoSellercenterRolesGetRole) SetSellerId(v int64) *TaobaoSellercenterRolesGetRole {
    s.SellerId = &v
    return s
}
func (s *TaobaoSellercenterRolesGetRole) SetRoleName(v string) *TaobaoSellercenterRolesGetRole {
    s.RoleName = &v
    return s
}
