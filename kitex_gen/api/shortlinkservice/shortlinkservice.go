// Code generated by Kitex v0.9.1. DO NOT EDIT.

package shortlinkservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	api "newRpc/kitex_gen/api"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"generateShortLink": kitex.NewMethodInfo(
		generateShortLinkHandler,
		newShortLinkServiceGenerateShortLinkArgs,
		newShortLinkServiceGenerateShortLinkResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"deleteShortLink": kitex.NewMethodInfo(
		deleteShortLinkHandler,
		newShortLinkServiceDeleteShortLinkArgs,
		newShortLinkServiceDeleteShortLinkResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"updateShortLink": kitex.NewMethodInfo(
		updateShortLinkHandler,
		newShortLinkServiceUpdateShortLinkArgs,
		newShortLinkServiceUpdateShortLinkResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"getUserShortLinks": kitex.NewMethodInfo(
		getUserShortLinksHandler,
		newShortLinkServiceGetUserShortLinksArgs,
		newShortLinkServiceGetUserShortLinksResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"getShortLinkRankings": kitex.NewMethodInfo(
		getShortLinkRankingsHandler,
		newShortLinkServiceGetShortLinkRankingsArgs,
		newShortLinkServiceGetShortLinkRankingsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	shortLinkServiceServiceInfo                = NewServiceInfo()
	shortLinkServiceServiceInfoForClient       = NewServiceInfoForClient()
	shortLinkServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return shortLinkServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return shortLinkServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return shortLinkServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ShortLinkService"
	handlerType := (*api.ShortLinkService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func generateShortLinkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShortLinkServiceGenerateShortLinkArgs)
	realResult := result.(*api.ShortLinkServiceGenerateShortLinkResult)
	success, err := handler.(api.ShortLinkService).GenerateShortLink(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortLinkServiceGenerateShortLinkArgs() interface{} {
	return api.NewShortLinkServiceGenerateShortLinkArgs()
}

func newShortLinkServiceGenerateShortLinkResult() interface{} {
	return api.NewShortLinkServiceGenerateShortLinkResult()
}

func deleteShortLinkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShortLinkServiceDeleteShortLinkArgs)
	realResult := result.(*api.ShortLinkServiceDeleteShortLinkResult)
	success, err := handler.(api.ShortLinkService).DeleteShortLink(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortLinkServiceDeleteShortLinkArgs() interface{} {
	return api.NewShortLinkServiceDeleteShortLinkArgs()
}

func newShortLinkServiceDeleteShortLinkResult() interface{} {
	return api.NewShortLinkServiceDeleteShortLinkResult()
}

func updateShortLinkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShortLinkServiceUpdateShortLinkArgs)
	realResult := result.(*api.ShortLinkServiceUpdateShortLinkResult)
	success, err := handler.(api.ShortLinkService).UpdateShortLink(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortLinkServiceUpdateShortLinkArgs() interface{} {
	return api.NewShortLinkServiceUpdateShortLinkArgs()
}

func newShortLinkServiceUpdateShortLinkResult() interface{} {
	return api.NewShortLinkServiceUpdateShortLinkResult()
}

func getUserShortLinksHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShortLinkServiceGetUserShortLinksArgs)
	realResult := result.(*api.ShortLinkServiceGetUserShortLinksResult)
	success, err := handler.(api.ShortLinkService).GetUserShortLinks(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortLinkServiceGetUserShortLinksArgs() interface{} {
	return api.NewShortLinkServiceGetUserShortLinksArgs()
}

func newShortLinkServiceGetUserShortLinksResult() interface{} {
	return api.NewShortLinkServiceGetUserShortLinksResult()
}

func getShortLinkRankingsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ShortLinkServiceGetShortLinkRankingsArgs)
	realResult := result.(*api.ShortLinkServiceGetShortLinkRankingsResult)
	success, err := handler.(api.ShortLinkService).GetShortLinkRankings(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newShortLinkServiceGetShortLinkRankingsArgs() interface{} {
	return api.NewShortLinkServiceGetShortLinkRankingsArgs()
}

func newShortLinkServiceGetShortLinkRankingsResult() interface{} {
	return api.NewShortLinkServiceGetShortLinkRankingsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GenerateShortLink(ctx context.Context, req *api.GenerateShortLinkRequest) (r *api.GenerateShortLinkResponse, err error) {
	var _args api.ShortLinkServiceGenerateShortLinkArgs
	_args.Req = req
	var _result api.ShortLinkServiceGenerateShortLinkResult
	if err = p.c.Call(ctx, "generateShortLink", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteShortLink(ctx context.Context, req *api.DeleteShortLinkRequest) (r *api.DeleteShortLinkResponse, err error) {
	var _args api.ShortLinkServiceDeleteShortLinkArgs
	_args.Req = req
	var _result api.ShortLinkServiceDeleteShortLinkResult
	if err = p.c.Call(ctx, "deleteShortLink", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateShortLink(ctx context.Context, req *api.UpdateShortLinkRequest) (r *api.UpdateShortLinkResponse, err error) {
	var _args api.ShortLinkServiceUpdateShortLinkArgs
	_args.Req = req
	var _result api.ShortLinkServiceUpdateShortLinkResult
	if err = p.c.Call(ctx, "updateShortLink", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserShortLinks(ctx context.Context, req *api.GetUserShortLinksRequest) (r []*api.Url, err error) {
	var _args api.ShortLinkServiceGetUserShortLinksArgs
	_args.Req = req
	var _result api.ShortLinkServiceGetUserShortLinksResult
	if err = p.c.Call(ctx, "getUserShortLinks", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetShortLinkRankings(ctx context.Context, req *api.GetShortLinkRankingsRequest) (r []*api.Url, err error) {
	var _args api.ShortLinkServiceGetShortLinkRankingsArgs
	_args.Req = req
	var _result api.ShortLinkServiceGetShortLinkRankingsResult
	if err = p.c.Call(ctx, "getShortLinkRankings", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
