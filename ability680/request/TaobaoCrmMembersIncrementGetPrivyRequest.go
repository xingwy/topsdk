package request

import (
        "topsdk/util"
    )

type TaobaoCrmMembersIncrementGetPrivyRequest struct {
    /*
        卖家修改会员信息的时间起点.     */
    StartModify  *util.LocalTime `json:"start_modify,omitempty" required:"false" `
    /*
        卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.     */
    EndModify  *util.LocalTime `json:"end_modify,omitempty" required:"false" `
    /*
        每页显示的会员数，page_size的值不能超过100，最小值要大于1     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        会员等级 defalutValue��-1    */
    Grade  *int64 `json:"grade,omitempty" required:"false" `
    /*
        显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1     */
    CurrentPage  *int64 `json:"current_page" required:"true" `
}

func (s *TaobaoCrmMembersIncrementGetPrivyRequest) SetStartModify(v util.LocalTime) *TaobaoCrmMembersIncrementGetPrivyRequest {
    s.StartModify = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyRequest) SetEndModify(v util.LocalTime) *TaobaoCrmMembersIncrementGetPrivyRequest {
    s.EndModify = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyRequest) SetPageSize(v int64) *TaobaoCrmMembersIncrementGetPrivyRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyRequest) SetGrade(v int64) *TaobaoCrmMembersIncrementGetPrivyRequest {
    s.Grade = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetPrivyRequest) SetCurrentPage(v int64) *TaobaoCrmMembersIncrementGetPrivyRequest {
    s.CurrentPage = &v
    return s
}

func (req *TaobaoCrmMembersIncrementGetPrivyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.StartModify != nil) {
        paramMap["start_modify"] = *req.StartModify
    }
    if(req.EndModify != nil) {
        paramMap["end_modify"] = *req.EndModify
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.Grade != nil) {
        paramMap["grade"] = *req.Grade
    }
    if(req.CurrentPage != nil) {
        paramMap["current_page"] = *req.CurrentPage
    }
    return paramMap
}

func (req *TaobaoCrmMembersIncrementGetPrivyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}