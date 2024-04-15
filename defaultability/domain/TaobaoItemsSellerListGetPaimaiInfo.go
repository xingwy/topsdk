package domain

import (
        "topsdk/util"
    )

type TaobaoItemsSellerListGetPaimaiInfo struct {
    /*
        用户自定义的固定保证金。如果用户未传则说明用户选择使用淘宝默认的保证金模式10%，此时deposit等于0.     */
    Deposit  *int64 `json:"deposit,omitempty" `

    /*
        增价幅度或降价幅度     */
    Increment  *string `json:"increment,omitempty" `

    /*
        降价拍中的降价间隔     */
    Interval  *int64 `json:"interval,omitempty" `

    /*
        拍卖类型，目前包括增加拍，荷兰拍和降价拍。     */
    Mode  *int64 `json:"mode,omitempty" `

    /*
        降价拍的保留价     */
    Reserve  *string `json:"reserve,omitempty" `

    /*
        对于拍卖来说可以设定有效期，这里是有效期的小时数。     */
    ValidHour  *int64 `json:"valid_hour,omitempty" `

    /*
        对于拍卖来说可以设定有效期，这里是有效期的分钟数。     */
    ValidMinute  *int64 `json:"valid_minute,omitempty" `

    /*
        重复上拍总次数，如果不是重复上拍的，则为0     */
    Repeat  *int64 `json:"repeat,omitempty" `

    /*
        拍品开始时间     */
    Start  *util.LocalTime `json:"start,omitempty" `

    /*
        拍品结束时间     */
    End  *util.LocalTime `json:"end,omitempty" `

    /*
        拍品起拍价     */
    StartPrice  *int64 `json:"start_price,omitempty" `

    /*
        拍品封顶价（分）     */
    CeilPrice  *int64 `json:"ceil_price,omitempty" `

    /*
        增加拍延迟周期（分钟）     */
    DelayInMinute  *int64 `json:"delay_in_minute,omitempty" `

    /*
        拍卖周期（秒钟），非重复上架为（开始时间-结束时间），当为重复上拍时为一次重复上架的时间     */
    Period  *int64 `json:"period,omitempty" `

    /*
        降价时间周期（分钟）     */
    Frequency  *int64 `json:"frequency,omitempty" `

}

func (s *TaobaoItemsSellerListGetPaimaiInfo) SetDeposit(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Deposit = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetIncrement(v string) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Increment = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetInterval(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Interval = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetMode(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Mode = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetReserve(v string) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Reserve = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetValidHour(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.ValidHour = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetValidMinute(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.ValidMinute = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetRepeat(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Repeat = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetStart(v util.LocalTime) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Start = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetEnd(v util.LocalTime) *TaobaoItemsSellerListGetPaimaiInfo {
    s.End = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetStartPrice(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.StartPrice = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetCeilPrice(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.CeilPrice = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetDelayInMinute(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.DelayInMinute = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetPeriod(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Period = &v
    return s
}
func (s *TaobaoItemsSellerListGetPaimaiInfo) SetFrequency(v int64) *TaobaoItemsSellerListGetPaimaiInfo {
    s.Frequency = &v
    return s
}
