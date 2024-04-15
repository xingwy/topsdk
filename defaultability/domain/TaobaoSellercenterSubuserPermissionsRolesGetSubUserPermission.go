package domain


type TaobaoSellercenterSubuserPermissionsRolesGetSubUserPermission struct {
    /*
        子账号被赋予的角色信息(Role)列表。列表中的角色对象只有role_id，role_name，permissions信息     */
    Roles  *[]TaobaoSellercenterSubuserPermissionsRolesGetRole `json:"roles,omitempty" `

    /*
        子账号被直接赋予的权限点列表     */
    Permissions  *[]TaobaoSellercenterSubuserPermissionsRolesGetPermission `json:"permissions,omitempty" `

}

func (s *TaobaoSellercenterSubuserPermissionsRolesGetSubUserPermission) SetRoles(v []TaobaoSellercenterSubuserPermissionsRolesGetRole) *TaobaoSellercenterSubuserPermissionsRolesGetSubUserPermission {
    s.Roles = &v
    return s
}
func (s *TaobaoSellercenterSubuserPermissionsRolesGetSubUserPermission) SetPermissions(v []TaobaoSellercenterSubuserPermissionsRolesGetPermission) *TaobaoSellercenterSubuserPermissionsRolesGetSubUserPermission {
    s.Permissions = &v
    return s
}
