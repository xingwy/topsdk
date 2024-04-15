package request

import (
	"github.com/xingwy/topsdk/ability302/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoFuwuSpBillreordAddRequest struct {
	/*
	   确认单的账单明细     */
	ParamBillRecordDTO *domain.TaobaoFuwuSpBillreordAddBillRecordDto `json:"param_bill_record_d_t_o" required:"true" `
}

func (s *TaobaoFuwuSpBillreordAddRequest) SetParamBillRecordDTO(v domain.TaobaoFuwuSpBillreordAddBillRecordDto) *TaobaoFuwuSpBillreordAddRequest {
	s.ParamBillRecordDTO = &v
	return s
}

func (req *TaobaoFuwuSpBillreordAddRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ParamBillRecordDTO != nil {
		paramMap["param_bill_record_d_t_o"] = util.ConvertStruct(*req.ParamBillRecordDTO)
	}
	return paramMap
}

func (req *TaobaoFuwuSpBillreordAddRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
