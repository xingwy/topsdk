package ability302

import (
	"errors"
	"log"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability302/request"
	"github.com/xingwy/topsdk/ability302/response"
	"github.com/xingwy/topsdk/util"
)

type Ability302 struct {
	Client *topsdk.TopClient
}

func NewAbility302(client *topsdk.TopClient) *Ability302 {
	return &Ability302{client}
}

/*
获取内购服务及SKU详情
*/
func (ability *Ability302) TaobaoFuwuSkuGet(req *request.TaobaoFuwuSkuGetRequest) (*response.TaobaoFuwuSkuGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability302 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.fuwu.sku.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoFuwuSkuGetResponse{}
	if err != nil {
		log.Println("taobaoFuwuSkuGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
内购服务确认单申请接口
*/
func (ability *Ability302) TaobaoFuwuSpConfirmApply(req *request.TaobaoFuwuSpConfirmApplyRequest, session string) (*response.TaobaoFuwuSpConfirmApplyResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability302 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.fuwu.sp.confirm.apply", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoFuwuSpConfirmApplyResponse{}
	if err != nil {
		log.Println("taobaoFuwuSpConfirmApply error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
内购服务确认单明细上传接口
*/
func (ability *Ability302) TaobaoFuwuSpBillreordAdd(req *request.TaobaoFuwuSpBillreordAddRequest, session string) (*response.TaobaoFuwuSpBillreordAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability302 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.fuwu.sp.billreord.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoFuwuSpBillreordAddResponse{}
	if err != nil {
		log.Println("taobaoFuwuSpBillreordAdd error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
内购服务订单付款页获取接口
*/
func (ability *Ability302) TaobaoFuwuPurchaseOrderPay(req *request.TaobaoFuwuPurchaseOrderPayRequest) (*response.TaobaoFuwuPurchaseOrderPayResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability302 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.fuwu.purchase.order.pay", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoFuwuPurchaseOrderPayResponse{}
	if err != nil {
		log.Println("taobaoFuwuPurchaseOrderPay error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
