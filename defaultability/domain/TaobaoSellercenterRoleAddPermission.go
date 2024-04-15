package domain


type TaobaoSellercenterRoleAddPermission struct {
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

func (s *TaobaoSellercenterRoleAddPermission) SetPermissionCode(v string) *TaobaoSellercenterRoleAddPermission {
    s.PermissionCode = &v
    return s
}
func (s *TaobaoSellercenterRoleAddPermission) SetIsAuthorize(v int64) *TaobaoSellercenterRoleAddPermission {
    s.IsAuthorize = &v
    return s
}
func (s *TaobaoSellercenterRoleAddPermission) SetPermissionName(v string) *TaobaoSellercenterRoleAddPermission {
    s.PermissionName = &v
    return s
}
func (s *TaobaoSellercenterRoleAddPermission) SetParentCode(v string) *TaobaoSellercenterRoleAddPermission {
    s.ParentCode = &v
    return s
}
