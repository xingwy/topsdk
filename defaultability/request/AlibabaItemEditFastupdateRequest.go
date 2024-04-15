package request


type AlibabaItemEditFastupdateRequest struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品类目ID。若不需要修改商品类目，则不用填写     */
    CatId  *int64 `json:"cat_id,omitempty" required:"false" `
    /*
        产品ID，若不需要修改关联的产品信息，则不需要填写     */
    SpuId  *int64 `json:"spu_id,omitempty" required:"false" `
    /*
        编辑后的schema信息(增量更新，只填写需要更新的字段)     */
    Schema  *string `json:"schema" required:"true" `
}

func (s *AlibabaItemEditFastupdateRequest) SetItemId(v int64) *AlibabaItemEditFastupdateRequest {
    s.ItemId = &v
    return s
}
func (s *AlibabaItemEditFastupdateRequest) SetCatId(v int64) *AlibabaItemEditFastupdateRequest {
    s.CatId = &v
    return s
}
func (s *AlibabaItemEditFastupdateRequest) SetSpuId(v int64) *AlibabaItemEditFastupdateRequest {
    s.SpuId = &v
    return s
}
func (s *AlibabaItemEditFastupdateRequest) SetSchema(v string) *AlibabaItemEditFastupdateRequest {
    s.Schema = &v
    return s
}

func (req *AlibabaItemEditFastupdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
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

func (req *AlibabaItemEditFastupdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}