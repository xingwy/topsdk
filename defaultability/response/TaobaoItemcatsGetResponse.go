package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoItemcatsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   最近修改时间(如果取增量，会返回该字段)。
	*/
	LastModified util.LocalTime `json:"last_modified,omitempty" `
	/*
	   商品类目结构
	*/
	ItemCats []domain.TaobaoItemcatsGetItemCat `json:"item_cats,omitempty" `
}
