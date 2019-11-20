/**
 * @Time : 2019-07-17 18:58
 * @Author : ygqbasic@gmail.com
 * @File : ednpoint
 * @Software: VS Code
 */

package persistentvolume

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/fuhsicloud/fuhsicloud/src/util/encode"
)

type getPvRequest struct {
	Name string
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getPvRequest)
		res, err := s.Get(ctx, req.Name)
		return encode.Response{Data: res, Err: err}, err
	}
}
