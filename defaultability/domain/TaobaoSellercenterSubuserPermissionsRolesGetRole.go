package domain


type TaobaoSellercenterSubuserPermissionsRolesGetRole struct {
    /*
        角色id     */
    RoleId  *int64 `json:"role_id,omitempty" `

    /*
        角色名     */
    RoleName  *string `json:"role_name,omitempty" `

    /*
        角色描述     */
    Description  *string `json:"description,omitempty" `

    /*
        所拥有权限     */
    Permissions  *[]TaobaoSellercenterSubuserPermissionsRolesGetPermission `json:"permissions,omitempty" `

}

func (s *TaobaoSellercenterSubuserPermissionsRolesGetRole) SetRoleId(v int64) *TaobaoSellercenterSubuserPermissionsRolesGetRole {
    s.RoleId = &v
    return s
}
func (s *TaobaoSellercenterSubuserPermissionsRolesGetRole) SetRoleName(v string) *TaobaoSellercenterSubuserPermissionsRolesGetRole {
    s.RoleName = &v
    return s
}
func (s *TaobaoSellercenterSubuserPermissionsRolesGetRole) SetDescription(v string) *TaobaoSellercenterSubuserPermissionsRolesGetRole {
    s.Description = &v
    return s
}
func (s *TaobaoSellercenterSubuserPermissionsRolesGetRole) SetPermissions(v []TaobaoSellercenterSubuserPermissionsRolesGetPermission) *TaobaoSellercenterSubuserPermissionsRolesGetRole {
    s.Permissions = &v
    return s
}
