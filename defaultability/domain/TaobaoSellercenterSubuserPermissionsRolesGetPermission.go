package domain


type TaobaoSellercenterSubuserPermissionsRolesGetPermission struct {
    /*
        注册到权限中心的code值     */
    PermissionCode  *string `json:"permission_code,omitempty" `

    /*
        1 :允许授权  2：不允许授权 6：不允许授权但默认已有权限     */
    IsAuthorize  *int64 `json:"is_authorize,omitempty" `

    /*
        权限名称     */
    PermissionName  *string `json:"permission_name,omitempty" `

    /*
        父权限code     */
    ParentCode  *string `json:"parent_code,omitempty" `

}

func (s *TaobaoSellercenterSubuserPermissionsRolesGetPermission) SetPermissionCode(v string) *TaobaoSellercenterSubuserPermissionsRolesGetPermission {
    s.PermissionCode = &v
    return s
}
func (s *TaobaoSellercenterSubuserPermissionsRolesGetPermission) SetIsAuthorize(v int64) *TaobaoSellercenterSubuserPermissionsRolesGetPermission {
    s.IsAuthorize = &v
    return s
}
func (s *TaobaoSellercenterSubuserPermissionsRolesGetPermission) SetPermissionName(v string) *TaobaoSellercenterSubuserPermissionsRolesGetPermission {
    s.PermissionName = &v
    return s
}
func (s *TaobaoSellercenterSubuserPermissionsRolesGetPermission) SetParentCode(v string) *TaobaoSellercenterSubuserPermissionsRolesGetPermission {
    s.ParentCode = &v
    return s
}
