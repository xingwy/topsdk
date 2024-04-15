package domain

import (
        "topsdk/util"
    )

type TaobaoAppstoreSubscribeGetUserSubscribe struct {
    /*
        订购状况。应用订购者：subscribeUser;尚未订购：unsubscribeUser；非法用户：invalidateUser     */
    Status  *string `json:"status,omitempty" `

    /*
        订购开始时间。格式:yyyy-MM-dd HH:mm:ss     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" `

    /*
        订购结束时间。格式:yyyy-MM-dd HH:mm:ss     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" `

    /*
        0-无版本信息；1-初级版；2-中级版；3-高级版     */
    VersionNo  *int64 `json:"version_no,omitempty" `

}

func (s *TaobaoAppstoreSubscribeGetUserSubscribe) SetStatus(v string) *TaobaoAppstoreSubscribeGetUserSubscribe {
    s.Status = &v
    return s
}
func (s *TaobaoAppstoreSubscribeGetUserSubscribe) SetStartDate(v util.LocalTime) *TaobaoAppstoreSubscribeGetUserSubscribe {
    s.StartDate = &v
    return s
}
func (s *TaobaoAppstoreSubscribeGetUserSubscribe) SetEndDate(v util.LocalTime) *TaobaoAppstoreSubscribeGetUserSubscribe {
    s.EndDate = &v
    return s
}
func (s *TaobaoAppstoreSubscribeGetUserSubscribe) SetVersionNo(v int64) *TaobaoAppstoreSubscribeGetUserSubscribe {
    s.VersionNo = &v
    return s
}
