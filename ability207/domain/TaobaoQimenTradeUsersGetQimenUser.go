package domain

import (
        "topsdk/util"
    )

type TaobaoQimenTradeUsersGetQimenUser struct {
    /*
        memo     */
    Memo  *string `json:"memo,omitempty" `

    /*
        gmtCreate     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        sellerNick     */
    SellerNick  *string `json:"seller_nick,omitempty" `

}

func (s *TaobaoQimenTradeUsersGetQimenUser) SetMemo(v string) *TaobaoQimenTradeUsersGetQimenUser {
    s.Memo = &v
    return s
}
func (s *TaobaoQimenTradeUsersGetQimenUser) SetGmtCreate(v util.LocalTime) *TaobaoQimenTradeUsersGetQimenUser {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoQimenTradeUsersGetQimenUser) SetSellerNick(v string) *TaobaoQimenTradeUsersGetQimenUser {
    s.SellerNick = &v
    return s
}
