package domain


type TaobaoRdsDbCreateaccountRequestDbAccountModel struct {
    /*
        账户描述     */
    AccountDesc  *string `json:"account_desc,omitempty" `

    /*
        账户名称     */
    AccountName  *string `json:"account_name,omitempty" `

    /*
        新建db名称     */
    DbName  *string `json:"db_name,omitempty" `

    /*
        rds实例名称     */
    InstanceName  *string `json:"instance_name,omitempty" `

    /*
        账户密码     */
    Password  *string `json:"password,omitempty" `

}

func (s *TaobaoRdsDbCreateaccountRequestDbAccountModel) SetAccountDesc(v string) *TaobaoRdsDbCreateaccountRequestDbAccountModel {
    s.AccountDesc = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountRequestDbAccountModel) SetAccountName(v string) *TaobaoRdsDbCreateaccountRequestDbAccountModel {
    s.AccountName = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountRequestDbAccountModel) SetDbName(v string) *TaobaoRdsDbCreateaccountRequestDbAccountModel {
    s.DbName = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountRequestDbAccountModel) SetInstanceName(v string) *TaobaoRdsDbCreateaccountRequestDbAccountModel {
    s.InstanceName = &v
    return s
}
func (s *TaobaoRdsDbCreateaccountRequestDbAccountModel) SetPassword(v string) *TaobaoRdsDbCreateaccountRequestDbAccountModel {
    s.Password = &v
    return s
}
