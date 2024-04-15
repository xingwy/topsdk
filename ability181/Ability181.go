package ability181

import (
	"errors"
	"log"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability181/request"
	"github.com/xingwy/topsdk/ability181/response"
	"github.com/xingwy/topsdk/util"
)

type Ability181 struct {
	Client *topsdk.TopClient
}

func NewAbility181(client *topsdk.TopClient) *Ability181 {
	return &Ability181{client}
}

/*
sku上下架管理
*/
func (ability *Ability181) TaobaoSkuUpdateListingTmall(req *request.TaobaoSkuUpdateListingTmallRequest, session string) (*response.TaobaoSkuUpdateListingTmallResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability181 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sku.update.listing.tmall", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSkuUpdateListingTmallResponse{}
	if err != nil {
		log.Println("taobaoSkuUpdateListingTmall error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询商品类目属性变更
*/
func (ability *Ability181) TaobaoItemCatpropsModificationGet(req *request.TaobaoItemCatpropsModificationGetRequest, session string) (*response.TaobaoItemCatpropsModificationGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability181 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.item.catprops.modification.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemCatpropsModificationGetResponse{}
	if err != nil {
		log.Println("taobaoItemCatpropsModificationGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
