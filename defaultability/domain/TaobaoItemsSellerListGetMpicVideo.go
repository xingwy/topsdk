package domain


type TaobaoItemsSellerListGetMpicVideo struct {
    /*
        主图视频记录所关联的商品的数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        主图视频的时长，单位：秒     */
    VideoDuaration  *int64 `json:"video_duaration,omitempty" `

    /*
        主图视频的在淘视频中的ID     */
    VideoId  *int64 `json:"video_id,omitempty" `

    /*
        主图视频的缩略图URL     */
    VideoPic  *string `json:"video_pic,omitempty" `

    /*
        主图视频的状态     */
    VideoStatus  *int64 `json:"video_status,omitempty" `

}

func (s *TaobaoItemsSellerListGetMpicVideo) SetNumIid(v int64) *TaobaoItemsSellerListGetMpicVideo {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemsSellerListGetMpicVideo) SetVideoDuaration(v int64) *TaobaoItemsSellerListGetMpicVideo {
    s.VideoDuaration = &v
    return s
}
func (s *TaobaoItemsSellerListGetMpicVideo) SetVideoId(v int64) *TaobaoItemsSellerListGetMpicVideo {
    s.VideoId = &v
    return s
}
func (s *TaobaoItemsSellerListGetMpicVideo) SetVideoPic(v string) *TaobaoItemsSellerListGetMpicVideo {
    s.VideoPic = &v
    return s
}
func (s *TaobaoItemsSellerListGetMpicVideo) SetVideoStatus(v int64) *TaobaoItemsSellerListGetMpicVideo {
    s.VideoStatus = &v
    return s
}
