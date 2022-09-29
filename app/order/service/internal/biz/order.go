package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	Id     int64
	UserId int64
}

type OrderRepo interface {
	GetOrder(context.Context, int64) (*Order, error)
}

type OrderUseCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUseCase(repo OrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "order-service/usecase"))}
}

func (uc *OrderUseCase) Get(ctx context.Context, id int64) (*Order, error) {
	return uc.repo.GetOrder(ctx, id)
}
