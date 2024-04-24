package ability200

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability200/request"
	"github.com/xingwy/topsdk/ability200/response"
	"github.com/xingwy/topsdk/util"
)

type Ability200 struct {
	Client *topsdk.TopClient
}

func NewAbility200(client *topsdk.TopClient) *Ability200 {
	return &Ability200{client}
}

/*
卖家发起拦截
*/
func (ability *Ability200) TaobaoRpRefundIntercept(req *request.TaobaoRpRefundInterceptRequest, session string) (*response.TaobaoRpRefundInterceptResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability200 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.rp.refund.intercept", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRpRefundInterceptResponse{}
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
协商退货退款
*/
func (ability *Ability200) TaobaoRefundNegotiatereturn(req *request.TaobaoRefundNegotiatereturnRequest, session string) (*response.TaobaoRefundNegotiatereturnResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability200 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refund.negotiatereturn", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundNegotiatereturnResponse{}
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
协商退货退款渲染
*/
func (ability *Ability200) TaobaoRefundNegotiatereturnRender(req *request.TaobaoRefundNegotiatereturnRenderRequest, session string) (*response.TaobaoRefundNegotiatereturnRenderResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability200 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refund.negotiatereturn.render", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundNegotiatereturnRenderResponse{}
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
退款详情页渲染
*/
func (ability *Ability200) TaobaoRefundDetailGet(req *request.TaobaoRefundDetailGetRequest, session string) (*response.TaobaoRefundDetailGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability200 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refund.detail.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundDetailGetResponse{}
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
查询退款状态
*/
func (ability *Ability200) TaobaoRefundStatusGet(req *request.TaobaoRefundStatusGetRequest, session string) (*response.TaobaoRefundStatusGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability200 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refund.status.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundStatusGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
