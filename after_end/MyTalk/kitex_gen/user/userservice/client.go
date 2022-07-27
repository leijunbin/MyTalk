// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"MyTalk/kitex_gen/user"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RegisterUser(ctx context.Context, req *user.RegisterUserRequest, callOptions ...callopt.Option) (r *user.RegisterUserResponse, err error)
	LoginUser(ctx context.Context, req *user.LoginUserRequest, callOptions ...callopt.Option) (r *user.LoginUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) RegisterUser(ctx context.Context, req *user.RegisterUserRequest, callOptions ...callopt.Option) (r *user.RegisterUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RegisterUser(ctx, req)
}

func (p *kUserServiceClient) LoginUser(ctx context.Context, req *user.LoginUserRequest, callOptions ...callopt.Option) (r *user.LoginUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginUser(ctx, req)
}