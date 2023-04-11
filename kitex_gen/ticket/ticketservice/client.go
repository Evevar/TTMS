// Code generated by Kitex v0.4.4. DO NOT EDIT.

package ticketservice

import (
	ticket "TTMS/kitex_gen/ticket"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	BatchAddTicket(ctx context.Context, Req *ticket.BatchAddTicketRequest, callOptions ...callopt.Option) (r *ticket.BatchAddTicketResponse, err error)
	UpdateTicket(ctx context.Context, Req *ticket.UpdateTicketRequest, callOptions ...callopt.Option) (r *ticket.UpdateTicketResponse, err error)
	GetAllTicket(ctx context.Context, Req *ticket.GetAllTicketRequest, callOptions ...callopt.Option) (r *ticket.GetAllTicketResponse, err error)
	BuyTicket(ctx context.Context, Req *ticket.BuyTicketRequest, callOptions ...callopt.Option) (r *ticket.BuyTicketResponse, err error)
	ReturnTicket(ctx context.Context, Req *ticket.ReturnTicketRequest, callOptions ...callopt.Option) (r *ticket.ReturnTicketResponse, err error)
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
	return &kTicketServiceClient{
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

type kTicketServiceClient struct {
	*kClient
}

func (p *kTicketServiceClient) BatchAddTicket(ctx context.Context, Req *ticket.BatchAddTicketRequest, callOptions ...callopt.Option) (r *ticket.BatchAddTicketResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchAddTicket(ctx, Req)
}

func (p *kTicketServiceClient) UpdateTicket(ctx context.Context, Req *ticket.UpdateTicketRequest, callOptions ...callopt.Option) (r *ticket.UpdateTicketResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateTicket(ctx, Req)
}

func (p *kTicketServiceClient) GetAllTicket(ctx context.Context, Req *ticket.GetAllTicketRequest, callOptions ...callopt.Option) (r *ticket.GetAllTicketResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllTicket(ctx, Req)
}

func (p *kTicketServiceClient) BuyTicket(ctx context.Context, Req *ticket.BuyTicketRequest, callOptions ...callopt.Option) (r *ticket.BuyTicketResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BuyTicket(ctx, Req)
}

func (p *kTicketServiceClient) ReturnTicket(ctx context.Context, Req *ticket.ReturnTicketRequest, callOptions ...callopt.Option) (r *ticket.ReturnTicketResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReturnTicket(ctx, Req)
}
