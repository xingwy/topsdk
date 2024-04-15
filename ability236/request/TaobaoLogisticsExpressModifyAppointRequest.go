package request

import (
	"github.com/xingwy/topsdk/ability236/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoLogisticsExpressModifyAppointRequest struct {
	/*
	   改约请求对象     */
	ExpressModifyAppointTopRequest *domain.TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto `json:"express_modify_appoint_top_request" required:"true" `
}

func (s *TaobaoLogisticsExpressModifyAppointRequest) SetExpressModifyAppointTopRequest(v domain.TaobaoLogisticsExpressModifyAppointExpressModifyAppointTopRequestDto) *TaobaoLogisticsExpressModifyAppointRequest {
	s.ExpressModifyAppointTopRequest = &v
	return s
}

func (req *TaobaoLogisticsExpressModifyAppointRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ExpressModifyAppointTopRequest != nil {
		paramMap["express_modify_appoint_top_request"] = util.ConvertStruct(*req.ExpressModifyAppointTopRequest)
	}
	return paramMap
}

func (req *TaobaoLogisticsExpressModifyAppointRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
