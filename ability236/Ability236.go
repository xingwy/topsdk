package ability236

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/ability236/request"
	"github.com/xingwy/topsdk/ability236/response"
	"github.com/xingwy/topsdk/util"
)

type Ability236 struct {
	Client *topsdk.TopClient
}

func NewAbility236(client *topsdk.TopClient) *Ability236 {
	return &Ability236{client}
}

/*
ToB仓储发货
*/
func (ability *Ability236) TaobaoUopTobOrderCreate(req *request.TaobaoUopTobOrderCreateRequest, session string) (*response.TaobaoUopTobOrderCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability236 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.uop.tob.order.create", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoUopTobOrderCreateResponse{}
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
云打印客户端更新配置服务
*/
func (ability *Ability236) CainiaoWaybillprintClientupdateGetconfig(req *request.CainiaoWaybillprintClientupdateGetconfigRequest) (*response.CainiaoWaybillprintClientupdateGetconfigResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability236 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("cainiao.waybillprint.clientupdate.getconfig", req.ToMap(), req.ToFileMap())
	var respStruct = response.CainiaoWaybillprintClientupdateGetconfigResponse{}
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
自定义区内容更新
*/
func (ability *Ability236) CainiaoCloudprintCustomareaUpdate(req *request.CainiaoCloudprintCustomareaUpdateRequest, session string) (*response.CainiaoCloudprintCustomareaUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability236 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("cainiao.cloudprint.customarea.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.CainiaoCloudprintCustomareaUpdateResponse{}
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
商家库存调整
*/
func (ability *Ability236) CainiaoMerchantInventoryAdjust(req *request.CainiaoMerchantInventoryAdjustRequest, session string) (*response.CainiaoMerchantInventoryAdjustResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability236 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("cainiao.merchant.inventory.adjust", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.CainiaoMerchantInventoryAdjustResponse{}
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
快递改约api
*/
func (ability *Ability236) TaobaoLogisticsExpressModifyAppoint(req *request.TaobaoLogisticsExpressModifyAppointRequest, session string) (*response.TaobaoLogisticsExpressModifyAppointResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability236 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.logistics.express.modify.appoint", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoLogisticsExpressModifyAppointResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
