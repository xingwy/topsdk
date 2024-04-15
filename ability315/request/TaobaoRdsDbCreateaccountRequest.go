package request

import (
	"github.com/xingwy/topsdk/ability315/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoRdsDbCreateaccountRequest struct {
	/*
	   入参对象     */
	Param0 *domain.TaobaoRdsDbCreateaccountRequestDbAccountModel `json:"param0,omitempty" required:"false" `
}

func (s *TaobaoRdsDbCreateaccountRequest) SetParam0(v domain.TaobaoRdsDbCreateaccountRequestDbAccountModel) *TaobaoRdsDbCreateaccountRequest {
	s.Param0 = &v
	return s
}

func (req *TaobaoRdsDbCreateaccountRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Param0 != nil {
		paramMap["param0"] = util.ConvertStruct(*req.Param0)
	}
	return paramMap
}

func (req *TaobaoRdsDbCreateaccountRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
