package domain


type TaobaoSubuserDepartmentsGetDepartment struct {
    /*
        部门ID     */
    DepartmentId  *int64 `json:"department_id,omitempty" `

    /*
        部门名称     */
    DepartmentName  *string `json:"department_name,omitempty" `

    /*
        当前部门的父部门ID     */
    ParentId  *int64 `json:"parent_id,omitempty" `

    /*
        部门下关联的子账号id列表     */
    SubUserIds  *[]int64 `json:"sub_user_ids,omitempty" `

}

func (s *TaobaoSubuserDepartmentsGetDepartment) SetDepartmentId(v int64) *TaobaoSubuserDepartmentsGetDepartment {
    s.DepartmentId = &v
    return s
}
func (s *TaobaoSubuserDepartmentsGetDepartment) SetDepartmentName(v string) *TaobaoSubuserDepartmentsGetDepartment {
    s.DepartmentName = &v
    return s
}
func (s *TaobaoSubuserDepartmentsGetDepartment) SetParentId(v int64) *TaobaoSubuserDepartmentsGetDepartment {
    s.ParentId = &v
    return s
}
func (s *TaobaoSubuserDepartmentsGetDepartment) SetSubUserIds(v []int64) *TaobaoSubuserDepartmentsGetDepartment {
    s.SubUserIds = &v
    return s
}
