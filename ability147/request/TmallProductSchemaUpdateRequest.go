package request


type TmallProductSchemaUpdateRequest struct {
    /*
        产品编号     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        根据tmall.product.update.schema.get生成的产品更新规则入参数据     */
    XmlData  *string `json:"xml_data" required:"true" `
}

func (s *TmallProductSchemaUpdateRequest) SetProductId(v int64) *TmallProductSchemaUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TmallProductSchemaUpdateRequest) SetXmlData(v string) *TmallProductSchemaUpdateRequest {
    s.XmlData = &v
    return s
}

func (req *TmallProductSchemaUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.XmlData != nil) {
        paramMap["xml_data"] = *req.XmlData
    }
    return paramMap
}

func (req *TmallProductSchemaUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}