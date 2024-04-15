package request

import (
	"github.com/xingwy/topsdk/ability302/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoFuwuSpConfirmApplyRequest struct {
	/*
	   确认单申请     */
	ParamIncomeConfirmDTO *domain.TaobaoFuwuSpConfirmApplyIncomeConfirmDto `json:"param_income_confirm_d_t_o" required:"true" `
}

func (s *TaobaoFuwuSpConfirmApplyRequest) SetParamIncomeConfirmDTO(v domain.TaobaoFuwuSpConfirmApplyIncomeConfirmDto) *TaobaoFuwuSpConfirmApplyRequest {
	s.ParamIncomeConfirmDTO = &v
	return s
}

func (req *TaobaoFuwuSpConfirmApplyRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ParamIncomeConfirmDTO != nil {
		paramMap["param_income_confirm_d_t_o"] = util.ConvertStruct(*req.ParamIncomeConfirmDTO)
	}
	return paramMap
}

func (req *TaobaoFuwuSpConfirmApplyRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
