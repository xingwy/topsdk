package request

import (
        "topsdk/util"
    )

type TaobaoCrmMembersIncrementGetRequest struct {
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

func (s *TaobaoCrmMembersIncrementGetRequest) SetStartModify(v util.LocalTime) *TaobaoCrmMembersIncrementGetRequest {
    s.StartModify = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetRequest) SetEndModify(v util.LocalTime) *TaobaoCrmMembersIncrementGetRequest {
    s.EndModify = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetRequest) SetPageSize(v int64) *TaobaoCrmMembersIncrementGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetRequest) SetGrade(v int64) *TaobaoCrmMembersIncrementGetRequest {
    s.Grade = &v
    return s
}
func (s *TaobaoCrmMembersIncrementGetRequest) SetCurrentPage(v int64) *TaobaoCrmMembersIncrementGetRequest {
    s.CurrentPage = &v
    return s
}

func (req *TaobaoCrmMembersIncrementGetRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoCrmMembersIncrementGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}