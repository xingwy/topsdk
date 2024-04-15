package domain

import (
        "topsdk/util"
    )

type TaobaoSubuserFullinfoGetSubUserFullInfo struct {
    /*
        工作地点     */
    WorkLocation  *string `json:"work_location,omitempty" `

    /*
        员工性别  1:男;  2:女     */
    Sex  *int64 `json:"sex,omitempty" `

    /*
        子账号Id     */
    SubId  *int64 `json:"sub_id,omitempty" `

    /*
        员工花名     */
    EmployeeNickname  *string `json:"employee_nickname,omitempty" `

    /*
        子账号当前状态：1正常，2卖家停用，3处罚冻结     */
    SubStatus  *int64 `json:"sub_status,omitempty" `

    /*
        部门Id     */
    DepartmentId  *int64 `json:"department_id,omitempty" `

    /*
        职务等级     */
    DutyLevel  *int64 `json:"duty_level,omitempty" `

    /*
        子账号是否参与分流 true:参与分流 false:未参与分流     */
    SubDispatchStatus  *bool `json:"sub_dispatch_status,omitempty" `

    /*
        主账号企业邮箱     */
    UserEmail  *string `json:"user_email,omitempty" `

    /*
        子账号是否已欠费 true:已欠费 false:未欠费     */
    SubOwedStatus  *bool `json:"sub_owed_status,omitempty" `

    /*
        职务名称     */
    DutyName  *string `json:"duty_name,omitempty" `

    /*
        员工姓名     */
    EmployeeName  *string `json:"employee_name,omitempty" `

    /*
        直接上级的员工ID     */
    LeaderId  *int64 `json:"leader_id,omitempty" `

    /*
        入职员工工号     */
    EmployeeNum  *string `json:"employee_num,omitempty" `

    /*
        父部门Id     */
    ParentDepartment  *int64 `json:"parent_department,omitempty" `

    /*
        职务Id     */
    DutyId  *int64 `json:"duty_id,omitempty" `

    /*
        员工入职时间     */
    EntryDate  *util.LocalTime `json:"entry_date,omitempty" `

    /*
        主账号Id     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        部门名称     */
    DepartmentName  *string `json:"department_name,omitempty" `

    /*
        员工ID     */
    EmployeeId  *int64 `json:"employee_id,omitempty" `

    /*
        子账号用户名     */
    SubNick  *string `json:"sub_nick,omitempty" `

    /*
        主账号用户名     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        办公电话     */
    OfficePhone  *string `json:"office_phone,omitempty" `

    /*
        子账号企业邮箱     */
    SubuserEmail  *string `json:"subuser_email,omitempty" `

}

func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetWorkLocation(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.WorkLocation = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetSex(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.Sex = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetSubId(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.SubId = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetEmployeeNickname(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.EmployeeNickname = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetSubStatus(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.SubStatus = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetDepartmentId(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.DepartmentId = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetDutyLevel(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.DutyLevel = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetSubDispatchStatus(v bool) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.SubDispatchStatus = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetUserEmail(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.UserEmail = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetSubOwedStatus(v bool) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.SubOwedStatus = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetDutyName(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.DutyName = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetEmployeeName(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.EmployeeName = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetLeaderId(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.LeaderId = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetEmployeeNum(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.EmployeeNum = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetParentDepartment(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.ParentDepartment = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetDutyId(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.DutyId = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetEntryDate(v util.LocalTime) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.EntryDate = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetUserId(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.UserId = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetDepartmentName(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.DepartmentName = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetEmployeeId(v int64) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.EmployeeId = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetSubNick(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.SubNick = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetUserNick(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.UserNick = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetOfficePhone(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.OfficePhone = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetSubUserFullInfo) SetSubuserEmail(v string) *TaobaoSubuserFullinfoGetSubUserFullInfo {
    s.SubuserEmail = &v
    return s
}
