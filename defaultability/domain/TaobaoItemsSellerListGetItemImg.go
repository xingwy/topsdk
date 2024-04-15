package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoItemsSellerListGetItemImg struct {
	/*
	   图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   商品图片的id，和商品相对应（主图id默认为0）     */
	Id *int64 `json:"id,omitempty" `

	/*
	   图片放在第几张（多图时可设置）     */
	Position *int64 `json:"position,omitempty" `

	/*
	   图片链接地址     */
	Url *string `json:"url,omitempty" `
}

func (s *TaobaoItemsSellerListGetItemImg) SetCreated(v util.LocalTime) *TaobaoItemsSellerListGetItemImg {
	s.Created = &v
	return s
}
func (s *TaobaoItemsSellerListGetItemImg) SetId(v int64) *TaobaoItemsSellerListGetItemImg {
	s.Id = &v
	return s
}
func (s *TaobaoItemsSellerListGetItemImg) SetPosition(v int64) *TaobaoItemsSellerListGetItemImg {
	s.Position = &v
	return s
}
func (s *TaobaoItemsSellerListGetItemImg) SetUrl(v string) *TaobaoItemsSellerListGetItemImg {
	s.Url = &v
	return s
}
