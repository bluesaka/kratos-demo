package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos-demo/api/order/service/v1"
	"kratos-demo/app/order/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	v1.UnimplementedOrderServer

	oc  *biz.OrderUseCase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderUseCase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "module", "service/order")),
	}
}
func (s *OrderService) GetOrder(ctx context.Context, req *v1.GetOrderReq) (*v1.GetOrderReply, error) {
	order, err := s.oc.Get(ctx, req.Id)
	if err != nil {
		return &v1.GetOrderReply{}, err
	}

	return &v1.GetOrderReply{
		Id:     order.Id,
		UserId: order.UserId,
	}, nil
}
