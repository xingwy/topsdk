package request

import (
        "topsdk/util"
    )

type TaobaoJushitaJdpUsersGetRequest struct {
    /*
        普通isv不能传入此参数     */
    TargetAppkey  *string `json:"target_appkey,omitempty" required:"false" `
    /*
        此参数一般不用传，用于查询最后更改时间在某个时间段内的用户     */
    StartModified  *util.LocalTime `json:"start_modified,omitempty" required:"false" `
    /*
        此参数一般不用传，用于查询最后更改时间在某个时间段内的用户     */
    EndModified  *util.LocalTime `json:"end_modified,omitempty" required:"false" `
    /*
        每页记录数，默认200 defalutValue��200    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        当前页数 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        如果传了user_id表示单条查询     */
    UserId  *int64 `json:"user_id,omitempty" required:"false" `
}

func (s *TaobaoJushitaJdpUsersGetRequest) SetTargetAppkey(v string) *TaobaoJushitaJdpUsersGetRequest {
    s.TargetAppkey = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetRequest) SetStartModified(v util.LocalTime) *TaobaoJushitaJdpUsersGetRequest {
    s.StartModified = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetRequest) SetEndModified(v util.LocalTime) *TaobaoJushitaJdpUsersGetRequest {
    s.EndModified = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetRequest) SetPageSize(v int64) *TaobaoJushitaJdpUsersGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetRequest) SetPageNo(v int64) *TaobaoJushitaJdpUsersGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetRequest) SetUserId(v int64) *TaobaoJushitaJdpUsersGetRequest {
    s.UserId = &v
    return s
}

func (req *TaobaoJushitaJdpUsersGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TargetAppkey != nil) {
        paramMap["target_appkey"] = *req.TargetAppkey
    }
    if(req.StartModified != nil) {
        paramMap["start_modified"] = *req.StartModified
    }
    if(req.EndModified != nil) {
        paramMap["end_modified"] = *req.EndModified
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.UserId != nil) {
        paramMap["user_id"] = *req.UserId
    }
    return paramMap
}

func (req *TaobaoJushitaJdpUsersGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}