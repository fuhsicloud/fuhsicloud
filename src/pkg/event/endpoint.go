/**
 * @Time : 2019-07-22 14:24
 * @Author : ygqbasic@gmail.com
 * @File : endpoint
 * @Software: VS Code
 */

package event

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/fuhsicloud/fuhsicloud/src/util/encode"
)

func makeAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := s.All(ctx)
		return encode.Response{Err: err, Data: res}, err
	}
}
