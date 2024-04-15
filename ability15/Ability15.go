package ability15

import (
	"errors"
	"log"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability15/request"
	"github.com/xingwy/topsdk/ability15/response"
	"github.com/xingwy/topsdk/util"
)

type Ability15 struct {
	Client *topsdk.TopClient
}

func NewAbility15(client *topsdk.TopClient) *Ability15 {
	return &Ability15{client}
}

/*
商家按照仓的类型获取仓库接口
*/
func (ability *Ability15) TaobaoWlbStoresBaseinfoGet(req *request.TaobaoWlbStoresBaseinfoGetRequest, session string) (*response.TaobaoWlbStoresBaseinfoGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability15 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.wlb.stores.baseinfo.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoWlbStoresBaseinfoGetResponse{}
	if err != nil {
		log.Println("taobaoWlbStoresBaseinfoGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
创建订单并发货
*/
func (ability *Ability15) TaobaoLogisticsConsignOrderCreateandsend(req *request.TaobaoLogisticsConsignOrderCreateandsendRequest) (*response.TaobaoLogisticsConsignOrderCreateandsendResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability15 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.logistics.consign.order.createandsend", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoLogisticsConsignOrderCreateandsendResponse{}
	if err != nil {
		log.Println("taobaoLogisticsConsignOrderCreateandsend error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
