package domain

import (
        "topsdk/util"
    )

type TaobaoRefundMessageAddRefundMessage struct {
    /*
        留言编号     */
    Id  *int64 `json:"id,omitempty" `

    /*
        留言创建时间。格式:yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoRefundMessageAddRefundMessage) SetId(v int64) *TaobaoRefundMessageAddRefundMessage {
    s.Id = &v
    return s
}
func (s *TaobaoRefundMessageAddRefundMessage) SetCreated(v util.LocalTime) *TaobaoRefundMessageAddRefundMessage {
    s.Created = &v
    return s
}
