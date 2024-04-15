package request

import (
        "topsdk/util"
    )

type TaobaoTmcGroupsGetRequest struct {
    /*
        要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。     */
    GroupNames  *[]string `json:"group_names,omitempty" required:"false" `
    /*
        页码 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页返回多少个分组 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoTmcGroupsGetRequest) SetGroupNames(v []string) *TaobaoTmcGroupsGetRequest {
    s.GroupNames = &v
    return s
}
func (s *TaobaoTmcGroupsGetRequest) SetPageNo(v int64) *TaobaoTmcGroupsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTmcGroupsGetRequest) SetPageSize(v int64) *TaobaoTmcGroupsGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoTmcGroupsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupNames != nil) {
        paramMap["group_names"] = util.ConvertBasicList(*req.GroupNames)
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoTmcGroupsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}