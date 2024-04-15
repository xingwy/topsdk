package request


type AlibabaItemEditSubmitRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        业务扩展参数，需与平台约定好     */
    BizType  *string `json:"biz_type,omitempty" required:"false" `
    /*
        商品类目ID。若不需要修改商品类目，则不用填写     */
    CatId  *int64 `json:"cat_id,omitempty" required:"false" `
    /*
        产品ID，若不需要修改关联的产品信息，则不需要填写     */
    SpuId  *int64 `json:"spu_id,omitempty" required:"false" `
    /*
        编辑后的schema信息，通过alibaba.item.edit.schema.get获取     */
    Schema  *string `json:"schema" required:"true" `
}

func (s *AlibabaItemEditSubmitRequest) SetItemId(v int64) *AlibabaItemEditSubmitRequest {
    s.ItemId = &v
    return s
}
func (s *AlibabaItemEditSubmitRequest) SetBizType(v string) *AlibabaItemEditSubmitRequest {
    s.BizType = &v
    return s
}
func (s *AlibabaItemEditSubmitRequest) SetCatId(v int64) *AlibabaItemEditSubmitRequest {
    s.CatId = &v
    return s
}
func (s *AlibabaItemEditSubmitRequest) SetSpuId(v int64) *AlibabaItemEditSubmitRequest {
    s.SpuId = &v
    return s
}
func (s *AlibabaItemEditSubmitRequest) SetSchema(v string) *AlibabaItemEditSubmitRequest {
    s.Schema = &v
    return s
}

func (req *AlibabaItemEditSubmitRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    if(req.SpuId != nil) {
        paramMap["spu_id"] = *req.SpuId
    }
    if(req.Schema != nil) {
        paramMap["schema"] = *req.Schema
    }
    return paramMap
}

func (req *AlibabaItemEditSubmitRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}