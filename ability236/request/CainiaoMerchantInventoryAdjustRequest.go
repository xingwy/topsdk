package request

import (
        "topsdk/ability236/domain"
        "topsdk/util"
    )

type CainiaoMerchantInventoryAdjustRequest struct {
    /*
        商家仓编辑库存     */
    AdjustRequest  *[]domain.CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto `json:"adjust_request" required:"true" `
    /*
        调用方应用名     */
    AppName  *string `json:"app_name" required:"true" `
    /*
        操作     */
    Operation  *string `json:"operation,omitempty" required:"false" `
}

func (s *CainiaoMerchantInventoryAdjustRequest) SetAdjustRequest(v []domain.CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto) *CainiaoMerchantInventoryAdjustRequest {
    s.AdjustRequest = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustRequest) SetAppName(v string) *CainiaoMerchantInventoryAdjustRequest {
    s.AppName = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustRequest) SetOperation(v string) *CainiaoMerchantInventoryAdjustRequest {
    s.Operation = &v
    return s
}

func (req *CainiaoMerchantInventoryAdjustRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AdjustRequest != nil) {
        paramMap["adjust_request"] = util.ConvertStructList(*req.AdjustRequest)
    }
    if(req.AppName != nil) {
        paramMap["app_name"] = *req.AppName
    }
    if(req.Operation != nil) {
        paramMap["operation"] = *req.Operation
    }
    return paramMap
}

func (req *CainiaoMerchantInventoryAdjustRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}