package domain

import (
        "topsdk/util"
    )

type TaobaoRefundMessagesGetRefundMessage struct {
    /*
        留言内容。最大长度: 400个字节     */
    Content  *string `json:"content,omitempty" `

    /*
        留言创建时间。格式:yyyy-MM-dd HH:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        留言编号     */
    Id  *int64 `json:"id,omitempty" `

    /*
        退款类型：NORMAL（普通留言），RETURN_GOODS_APPROVED（卖家留退货地址时留言）；如果为RETURN_GOODS_APPROVED，则退款留言中有卖家收货地址     */
    MessageType  *string `json:"message_type,omitempty" `

    /*
        留言者编号     */
    OwnerId  *int64 `json:"owner_id,omitempty" `

    /*
        留言者昵称     */
    OwnerNick  *string `json:"owner_nick,omitempty" `

    /*
        留言者身份1代表买家，2代表卖家，3代表小二     */
    OwnerRole  *string `json:"owner_role,omitempty" `

    /*
        退款编号。     */
    RefundId  *string `json:"refund_id,omitempty" `

    /*
        退款阶段，可选值：onsale（售中）, aftersale（售后）     */
    RefundPhase  *string `json:"refund_phase,omitempty" `

    /*
        图片链接     */
    PicUrls  *[]TaobaoRefundMessagesGetPicUrl `json:"pic_urls,omitempty" `

    /*
        留言者openid     */
    OwnerOpenUid  *string `json:"owner_open_uid,omitempty" `

}

func (s *TaobaoRefundMessagesGetRefundMessage) SetContent(v string) *TaobaoRefundMessagesGetRefundMessage {
    s.Content = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetCreated(v util.LocalTime) *TaobaoRefundMessagesGetRefundMessage {
    s.Created = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetId(v int64) *TaobaoRefundMessagesGetRefundMessage {
    s.Id = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetMessageType(v string) *TaobaoRefundMessagesGetRefundMessage {
    s.MessageType = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetOwnerId(v int64) *TaobaoRefundMessagesGetRefundMessage {
    s.OwnerId = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetOwnerNick(v string) *TaobaoRefundMessagesGetRefundMessage {
    s.OwnerNick = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetOwnerRole(v string) *TaobaoRefundMessagesGetRefundMessage {
    s.OwnerRole = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetRefundId(v string) *TaobaoRefundMessagesGetRefundMessage {
    s.RefundId = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetRefundPhase(v string) *TaobaoRefundMessagesGetRefundMessage {
    s.RefundPhase = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetPicUrls(v []TaobaoRefundMessagesGetPicUrl) *TaobaoRefundMessagesGetRefundMessage {
    s.PicUrls = &v
    return s
}
func (s *TaobaoRefundMessagesGetRefundMessage) SetOwnerOpenUid(v string) *TaobaoRefundMessagesGetRefundMessage {
    s.OwnerOpenUid = &v
    return s
}
