package domain


type TaobaoJushitaJdpUsersGetJdpUser struct {
    /*
        用户id     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        rds的id，平台appkey会返回rds_id而不是rds_name     */
    RdsId  *int64 `json:"rds_id,omitempty" `

    /*
        0:暂停1：正常2：sessoin失效，停止3：已删除     */
    Status  *int64 `json:"status,omitempty" `

    /*
        rds数据库的实例名     */
    RdsName  *string `json:"rds_name,omitempty" `

    /*
        用户昵称     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        卖家类型，B表示商城卖家，C表示淘宝卖家     */
    SellerType  *string `json:"seller_type,omitempty" `

    /*
        用户等级，用于区分大卖家，值越大则订单量越大     */
    Level  *int64 `json:"level,omitempty" `

    /*
        和数据回流绑定的appkey，用户的数据推到此appkey对应的rds后，会回传X_DOWNLOADED消息     */
    HlAppkey  *string `json:"hl_appkey,omitempty" `

}

func (s *TaobaoJushitaJdpUsersGetJdpUser) SetUserId(v int64) *TaobaoJushitaJdpUsersGetJdpUser {
    s.UserId = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetJdpUser) SetRdsId(v int64) *TaobaoJushitaJdpUsersGetJdpUser {
    s.RdsId = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetJdpUser) SetStatus(v int64) *TaobaoJushitaJdpUsersGetJdpUser {
    s.Status = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetJdpUser) SetRdsName(v string) *TaobaoJushitaJdpUsersGetJdpUser {
    s.RdsName = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetJdpUser) SetUserNick(v string) *TaobaoJushitaJdpUsersGetJdpUser {
    s.UserNick = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetJdpUser) SetSellerType(v string) *TaobaoJushitaJdpUsersGetJdpUser {
    s.SellerType = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetJdpUser) SetLevel(v int64) *TaobaoJushitaJdpUsersGetJdpUser {
    s.Level = &v
    return s
}
func (s *TaobaoJushitaJdpUsersGetJdpUser) SetHlAppkey(v string) *TaobaoJushitaJdpUsersGetJdpUser {
    s.HlAppkey = &v
    return s
}
