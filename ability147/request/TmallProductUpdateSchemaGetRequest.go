package request


type TmallProductUpdateSchemaGetRequest struct {
    /*
        产品编号     */
    ProductId  *int64 `json:"product_id" required:"true" `
}

func (s *TmallProductUpdateSchemaGetRequest) SetProductId(v int64) *TmallProductUpdateSchemaGetRequest {
    s.ProductId = &v
    return s
}

func (req *TmallProductUpdateSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    return paramMap
}

func (req *TmallProductUpdateSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}