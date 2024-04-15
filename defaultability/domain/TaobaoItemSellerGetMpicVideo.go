package domain


type TaobaoItemSellerGetMpicVideo struct {
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

func (s *TaobaoItemSellerGetMpicVideo) SetNumIid(v int64) *TaobaoItemSellerGetMpicVideo {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSellerGetMpicVideo) SetVideoDuaration(v int64) *TaobaoItemSellerGetMpicVideo {
    s.VideoDuaration = &v
    return s
}
func (s *TaobaoItemSellerGetMpicVideo) SetVideoId(v int64) *TaobaoItemSellerGetMpicVideo {
    s.VideoId = &v
    return s
}
func (s *TaobaoItemSellerGetMpicVideo) SetVideoPic(v string) *TaobaoItemSellerGetMpicVideo {
    s.VideoPic = &v
    return s
}
func (s *TaobaoItemSellerGetMpicVideo) SetVideoStatus(v int64) *TaobaoItemSellerGetMpicVideo {
    s.VideoStatus = &v
    return s
}
