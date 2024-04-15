package request


type TmallProductSchemaGetRequest struct {
    /*
        产品编号     */
    ProductId  *int64 `json:"product_id" required:"true" `
}

func (s *TmallProductSchemaGetRequest) SetProductId(v int64) *TmallProductSchemaGetRequest {
    s.ProductId = &v
    return s
}

func (req *TmallProductSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    return paramMap
}

func (req *TmallProductSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}