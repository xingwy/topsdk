package domain

import (
        "topsdk/util"
    )

type TaobaoSellercenterRoleAddRole struct {
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
        所拥有权限     */
    Permissions  *[]TaobaoSellercenterRoleAddPermission `json:"permissions,omitempty" `

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

func (s *TaobaoSellercenterRoleAddRole) SetRoleId(v int64) *TaobaoSellercenterRoleAddRole {
    s.RoleId = &v
    return s
}
func (s *TaobaoSellercenterRoleAddRole) SetDescription(v string) *TaobaoSellercenterRoleAddRole {
    s.Description = &v
    return s
}
func (s *TaobaoSellercenterRoleAddRole) SetModifiedTime(v util.LocalTime) *TaobaoSellercenterRoleAddRole {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoSellercenterRoleAddRole) SetPermissions(v []TaobaoSellercenterRoleAddPermission) *TaobaoSellercenterRoleAddRole {
    s.Permissions = &v
    return s
}
func (s *TaobaoSellercenterRoleAddRole) SetCreateTime(v util.LocalTime) *TaobaoSellercenterRoleAddRole {
    s.CreateTime = &v
    return s
}
func (s *TaobaoSellercenterRoleAddRole) SetSellerId(v int64) *TaobaoSellercenterRoleAddRole {
    s.SellerId = &v
    return s
}
func (s *TaobaoSellercenterRoleAddRole) SetRoleName(v string) *TaobaoSellercenterRoleAddRole {
    s.RoleName = &v
    return s
}
