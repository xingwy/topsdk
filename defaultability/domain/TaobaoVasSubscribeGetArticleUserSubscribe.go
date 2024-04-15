package domain

import (
        "topsdk/util"
    )

type TaobaoVasSubscribeGetArticleUserSubscribe struct {
    /*
        收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        订购关系到期时间     */
    Deadline  *util.LocalTime `json:"deadline,omitempty" `

}

func (s *TaobaoVasSubscribeGetArticleUserSubscribe) SetItemCode(v string) *TaobaoVasSubscribeGetArticleUserSubscribe {
    s.ItemCode = &v
    return s
}
func (s *TaobaoVasSubscribeGetArticleUserSubscribe) SetDeadline(v util.LocalTime) *TaobaoVasSubscribeGetArticleUserSubscribe {
    s.Deadline = &v
    return s
}
