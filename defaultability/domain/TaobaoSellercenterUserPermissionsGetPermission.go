package domain


type TaobaoSellercenterUserPermissionsGetPermission struct {
    /*
        注册到权限中心的code值     */
    PermissionCode  *string `json:"permission_code,omitempty" `

    /*
        权限名称     */
    PermissionName  *string `json:"permission_name,omitempty" `

    /*
        父权限code     */
    ParentCode  *string `json:"parent_code,omitempty" `

    /*
        1 :允许授权  2：不允许授权 6：不允许授权但默认已有权限     */
    IsAuthorize  *int64 `json:"is_authorize,omitempty" `

}

func (s *TaobaoSellercenterUserPermissionsGetPermission) SetPermissionCode(v string) *TaobaoSellercenterUserPermissionsGetPermission {
    s.PermissionCode = &v
    return s
}
func (s *TaobaoSellercenterUserPermissionsGetPermission) SetPermissionName(v string) *TaobaoSellercenterUserPermissionsGetPermission {
    s.PermissionName = &v
    return s
}
func (s *TaobaoSellercenterUserPermissionsGetPermission) SetParentCode(v string) *TaobaoSellercenterUserPermissionsGetPermission {
    s.ParentCode = &v
    return s
}
func (s *TaobaoSellercenterUserPermissionsGetPermission) SetIsAuthorize(v int64) *TaobaoSellercenterUserPermissionsGetPermission {
    s.IsAuthorize = &v
    return s
}
