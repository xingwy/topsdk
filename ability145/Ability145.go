package ability145

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability145/request"
	"github.com/xingwy/topsdk/ability145/response"
	"github.com/xingwy/topsdk/util"
)

type Ability145 struct {
	Client *topsdk.TopClient
}

func NewAbility145(client *topsdk.TopClient) *Ability145 {
	return &Ability145{client}
}

/*
查询订单备注列表
*/
func (ability *Ability145) TaobaoTradeSellermemosGet(req *request.TaobaoTradeSellermemosGetRequest, session string) (*response.TaobaoTradeSellermemosGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability145 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.trade.sellermemos.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTradeSellermemosGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询商家插旗
*/
func (ability *Ability145) TaobaoTradeSellerflagsGet(req *request.TaobaoTradeSellerflagsGetRequest, session string) (*response.TaobaoTradeSellerflagsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability145 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.trade.sellerflags.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTradeSellerflagsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
