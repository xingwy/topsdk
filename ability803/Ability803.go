package ability803

import (
	"errors"
	"log"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability803/request"
	"github.com/xingwy/topsdk/ability803/response"
	"github.com/xingwy/topsdk/util"
)

type Ability803 struct {
	Client *topsdk.TopClient
}

func NewAbility803(client *topsdk.TopClient) *Ability803 {
	return &Ability803{client}
}

/*
通过店铺id取得短链
*/
func (ability *Ability803) TaobaoWirelessBuntingShopShorturlCreate(req *request.TaobaoWirelessBuntingShopShorturlCreateRequest) (*response.TaobaoWirelessBuntingShopShorturlCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability803 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.wireless.bunting.shop.shorturl.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoWirelessBuntingShopShorturlCreateResponse{}
	if err != nil {
		log.Println("taobaoWirelessBuntingShopShorturlCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
