package ability206

import (
	"errors"
	"log"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability206/request"
	"github.com/xingwy/topsdk/ability206/response"
	"github.com/xingwy/topsdk/util"
)

type Ability206 struct {
	Client *topsdk.TopClient
}

func NewAbility206(client *topsdk.TopClient) *Ability206 {
	return &Ability206{client}
}

/*
订单全链路用户信息修改
*/
func (ability *Ability206) TaobaoJdsHluserUpdate(req *request.TaobaoJdsHluserUpdateRequest, session string) (*response.TaobaoJdsHluserUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jds.hluser.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJdsHluserUpdateResponse{}
	if err != nil {
		log.Println("taobaoJdsHluserUpdate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
订单全链路用户信息获取
*/
func (ability *Ability206) TaobaoJdsHluserGet(req *request.TaobaoJdsHluserGetRequest, session string) (*response.TaobaoJdsHluserGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jds.hluser.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJdsHluserGetResponse{}
	if err != nil {
		log.Println("taobaoJdsHluserGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取单条退款跟踪详情
*/
func (ability *Ability206) TaobaoJdsRefundTracesGet(req *request.TaobaoJdsRefundTracesGetRequest, session string) (*response.TaobaoJdsRefundTracesGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jds.refund.traces.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJdsRefundTracesGetResponse{}
	if err != nil {
		log.Println("taobaoJdsRefundTracesGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
订单全链路状态统计差异比较
*/
func (ability *Ability206) TaobaoJdsTradesStatisticsDiff(req *request.TaobaoJdsTradesStatisticsDiffRequest, session string) (*response.TaobaoJdsTradesStatisticsDiffResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jds.trades.statistics.diff", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJdsTradesStatisticsDiffResponse{}
	if err != nil {
		log.Println("taobaoJdsTradesStatisticsDiff error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取订单数量统计结果
*/
func (ability *Ability206) TaobaoJdsTradesStatisticsGet(req *request.TaobaoJdsTradesStatisticsGetRequest, session string) (*response.TaobaoJdsTradesStatisticsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jds.trades.statistics.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJdsTradesStatisticsGetResponse{}
	if err != nil {
		log.Println("taobaoJdsTradesStatisticsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
订单打标或者订单标签更新
*/
func (ability *Ability206) TaobaoOcTradetagAttach(req *request.TaobaoOcTradetagAttachRequest, session string) (*response.TaobaoOcTradetagAttachResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.oc.tradetag.attach", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOcTradetagAttachResponse{}
	if err != nil {
		log.Println("taobaoOcTradetagAttach error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
标签查询订单
*/
func (ability *Ability206) TaobaoOcTradesBytagGet(req *request.TaobaoOcTradesBytagGetRequest, session string) (*response.TaobaoOcTradesBytagGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.oc.trades.bytag.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOcTradesBytagGetResponse{}
	if err != nil {
		log.Println("taobaoOcTradesBytagGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据订单查询订单标签
*/
func (ability *Ability206) TaobaoOcTradetagsGet(req *request.TaobaoOcTradetagsGetRequest, session string) (*response.TaobaoOcTradetagsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability206 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.oc.tradetags.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOcTradetagsGetResponse{}
	if err != nil {
		log.Println("taobaoOcTradetagsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
