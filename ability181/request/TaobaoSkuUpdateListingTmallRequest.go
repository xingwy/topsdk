package request

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoSkuUpdateListingTmallRequest struct {
	/*
	   商品ID，必填     */
	ItemId *int64 `json:"item_id" required:"true" `
	/*
	   skuId和sku状态的映射，其中状态1代表上架，-2代表下架，必填     */
	SkuMap *map[string]interface{} `json:"sku_map" required:"true" `
}

func (s *TaobaoSkuUpdateListingTmallRequest) SetItemId(v int64) *TaobaoSkuUpdateListingTmallRequest {
	s.ItemId = &v
	return s
}
func (s *TaobaoSkuUpdateListingTmallRequest) SetSkuMap(v map[string]interface{}) *TaobaoSkuUpdateListingTmallRequest {
	s.SkuMap = &v
	return s
}

func (req *TaobaoSkuUpdateListingTmallRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.ItemId != nil {
		paramMap["item_id"] = *req.ItemId
	}
	if req.SkuMap != nil {
		paramMap["sku_map"] = util.ConvertStruct(*req.SkuMap)
	}
	return paramMap
}

func (req *TaobaoSkuUpdateListingTmallRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
