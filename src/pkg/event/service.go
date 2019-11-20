/**
 * @Time : 2019-07-22 13:51
 * @Author : ygqbasic@gmail.com
 * @File : service
 * @Software: VS Code
 */

package event

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/fuhsicloud/fuhsicloud/src/repository"
	"github.com/fuhsicloud/fuhsicloud/src/repository/types"
)

type Service interface {
	// 获取所有events
	All(ctx context.Context) (res []*types.Event, err error)
}

type service struct {
	logger     log.Logger
	repository repository.Repository
}

func NewService(logger log.Logger, store repository.Repository) Service {
	return &service{logger, store}
}

func (c *service) All(ctx context.Context) (res []*types.Event, err error) {
	return c.repository.Event().FindAllEvents()
}
