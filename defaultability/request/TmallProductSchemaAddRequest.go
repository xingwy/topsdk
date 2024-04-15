package request


type TmallProductSchemaAddRequest struct {
    /*
        商品发布的目标类目，必须是叶子类目     */
    CategoryId  *int64 `json:"category_id" required:"true" `
    /*
        品牌ID     */
    BrandId  *int64 `json:"brand_id,omitempty" required:"false" `
    /*
        根据tmall.product.add.schema.get生成的产品发布规则入参数据     */
    XmlData  *string `json:"xml_data" required:"true" `
}

func (s *TmallProductSchemaAddRequest) SetCategoryId(v int64) *TmallProductSchemaAddRequest {
    s.CategoryId = &v
    return s
}
func (s *TmallProductSchemaAddRequest) SetBrandId(v int64) *TmallProductSchemaAddRequest {
    s.BrandId = &v
    return s
}
func (s *TmallProductSchemaAddRequest) SetXmlData(v string) *TmallProductSchemaAddRequest {
    s.XmlData = &v
    return s
}

func (req *TmallProductSchemaAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.BrandId != nil) {
        paramMap["brand_id"] = *req.BrandId
    }
    if(req.XmlData != nil) {
        paramMap["xml_data"] = *req.XmlData
    }
    return paramMap
}

func (req *TmallProductSchemaAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}