package domain

import (
        "topsdk/util"
    )

type TaobaoItemsSellerListGetVideo struct {
    /*
        视频关联记录创建时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        视频关联记录的id，和商品相对应     */
    Id  *int64 `json:"id,omitempty" `

    /*
        视频关联记录修改时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        video的url连接地址。淘秀里视频记录里面存储的url地址     */
    Url  *string `json:"url,omitempty" `

    /*
        video的id，对应于视频在淘秀的存储记录     */
    VideoId  *int64 `json:"video_id,omitempty" `

}

func (s *TaobaoItemsSellerListGetVideo) SetCreated(v util.LocalTime) *TaobaoItemsSellerListGetVideo {
    s.Created = &v
    return s
}
func (s *TaobaoItemsSellerListGetVideo) SetId(v int64) *TaobaoItemsSellerListGetVideo {
    s.Id = &v
    return s
}
func (s *TaobaoItemsSellerListGetVideo) SetModified(v util.LocalTime) *TaobaoItemsSellerListGetVideo {
    s.Modified = &v
    return s
}
func (s *TaobaoItemsSellerListGetVideo) SetUrl(v string) *TaobaoItemsSellerListGetVideo {
    s.Url = &v
    return s
}
func (s *TaobaoItemsSellerListGetVideo) SetVideoId(v int64) *TaobaoItemsSellerListGetVideo {
    s.VideoId = &v
    return s
}
