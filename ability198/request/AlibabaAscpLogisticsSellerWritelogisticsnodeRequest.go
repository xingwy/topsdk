package request

import (
	"github.com/xingwy/topsdk/ability198/domain"
	"github.com/xingwy/topsdk/util"
)

type AlibabaAscpLogisticsSellerWritelogisticsnodeRequest struct {
	/*
	   物流发货单号     */
	LpOrderId *int64 `json:"lp_order_id" required:"true" `
	/*
	   物流节点     */
	Nodes *[]domain.AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO `json:"nodes" required:"true" `
}

func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeRequest) SetLpOrderId(v int64) *AlibabaAscpLogisticsSellerWritelogisticsnodeRequest {
	s.LpOrderId = &v
	return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeRequest) SetNodes(v []domain.AlibabaAscpLogisticsSellerWritelogisticsnodeLogisticsNodeTopDTO) *AlibabaAscpLogisticsSellerWritelogisticsnodeRequest {
	s.Nodes = &v
	return s
}

func (req *AlibabaAscpLogisticsSellerWritelogisticsnodeRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.LpOrderId != nil {
		paramMap["lp_order_id"] = *req.LpOrderId
	}
	if req.Nodes != nil {
		paramMap["nodes"] = util.ConvertStructList(*req.Nodes)
	}
	return paramMap
}

func (req *AlibabaAscpLogisticsSellerWritelogisticsnodeRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
