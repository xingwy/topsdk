package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoSellercenterRoleAddRequest struct {
	/*
	   表示卖家昵称     */
	Nick *string `json:"nick" required:"true" `
	/*
	   角色名     */
	Name *string `json:"name" required:"true" `
	/*
	   角色描述     */
	Description *string `json:"description,omitempty" required:"false" `
	/*
	   需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。     */
	PermissionCodes *[]string `json:"permission_codes,omitempty" required:"false" `
}

func (s *TaobaoSellercenterRoleAddRequest) SetNick(v string) *TaobaoSellercenterRoleAddRequest {
	s.Nick = &v
	return s
}
func (s *TaobaoSellercenterRoleAddRequest) SetName(v string) *TaobaoSellercenterRoleAddRequest {
	s.Name = &v
	return s
}
func (s *TaobaoSellercenterRoleAddRequest) SetDescription(v string) *TaobaoSellercenterRoleAddRequest {
	s.Description = &v
	return s
}
func (s *TaobaoSellercenterRoleAddRequest) SetPermissionCodes(v []string) *TaobaoSellercenterRoleAddRequest {
	s.PermissionCodes = &v
	return s
}

func (req *TaobaoSellercenterRoleAddRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Nick != nil {
		paramMap["nick"] = *req.Nick
	}
	if req.Name != nil {
		paramMap["name"] = *req.Name
	}
	if req.Description != nil {
		paramMap["description"] = *req.Description
	}
	if req.PermissionCodes != nil {
		paramMap["permission_codes"] = util.ConvertBasicList(*req.PermissionCodes)
	}
	return paramMap
}

func (req *TaobaoSellercenterRoleAddRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
