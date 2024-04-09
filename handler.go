package main

import (
	"context"
	"log"
	"newRpc/cache"
	"newRpc/dao"
	api "newRpc/kitex_gen/api"
)

// ShortLinkServiceImpl implements the last service interface defined in the IDL.
type ShortLinkServiceImpl struct{}

// GenerateShortLink implements the ShortLinkServiceImpl interface.
func (s *ShortLinkServiceImpl) GenerateShortLink(ctx context.Context, req *api.GenerateShortLinkRequest) (resp *api.GenerateShortLinkResponse, err error) {
	// TODO: Your code here...
	resp = new(api.GenerateShortLinkResponse)

	if resp.ShortUrl, err = dao.GenerateUrl(req.LongUrl); err != nil {
		resp.Message = err.Error()
		resp.Code = api.Code_DBErr
		return
	}

	if err = dao.InsertUrl(resp.ShortUrl, req.LongUrl); err != nil {
		resp.Message = err.Error()
		resp.Code = api.Code_DBErr
		return
	}

	resp.Code = api.Code_Success
	return
}

// DeleteShortLink implements the ShortLinkServiceImpl interface.
func (s *ShortLinkServiceImpl) DeleteShortLink(ctx context.Context, req *api.DeleteShortLinkRequest) (resp *api.DeleteShortLinkResponse, err error) {
	// TODO: Your code here...
	resp = new(api.DeleteShortLinkResponse)

	if err = dao.DeleteUrl(req.Id); err != nil {
		resp.Message = err.Error()
		resp.Code = api.Code_DBErr
		return
	}

	resp.Code = api.Code_Success
	return
}

// UpdateShortLink implements the ShortLinkServiceImpl interface.
func (s *ShortLinkServiceImpl) UpdateShortLink(ctx context.Context, req *api.UpdateShortLinkRequest) (resp *api.UpdateShortLinkResponse, err error) {
	// TODO: Your code here...
	resp = new(api.UpdateShortLinkResponse)

	if err = dao.UpdateUrl(req.OldShortUrl); err != nil {
		resp.Message = err.Error()
		resp.Code = api.Code_DBErr
		return
	}

	resp.Code = api.Code_Success
	return
}

// GetUserShortLinks implements the ShortLinkServiceImpl interface.
func (s *ShortLinkServiceImpl) GetUserShortLinks(ctx context.Context, req *api.GetUserShortLinksRequest) (resp []*api.Url, err error) {
	// TODO: Your code here...
	urls, err := dao.GetUserUrls(req.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var response []*api.Url
	for _, url := range urls {
		response = append(response, &api.Url{
			Id:       url.ID,
			LongUrl:  url.LongUrl,
			ShortUrl: url.ShortUrl,
			UserId:   url.UserID,
			Visits:   int64(url.Visits),
		})
	}

	return response, nil
}

// GetShortLinkRankings implements the ShortLinkServiceImpl interface.
func (s *ShortLinkServiceImpl) GetShortLinkRankings(ctx context.Context, req *api.GetShortLinkRankingsRequest) (resp []*api.Url, err error) {
	// TODO: Your code here...
	//获取短链数据排名
	ranks, err := cache.GetAllRank(ctx, cache.Rdb)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 构造响应数据
	resp = make([]*api.Url, 0, len(ranks))
	for shortURL, rank := range ranks {
		url := &api.Url{
			ShortUrl: shortURL,
			Rank:     int32(rank),
		}
		resp = append(resp, url)
	}

	return resp,nil
}
