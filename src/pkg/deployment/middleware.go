/**
 * @Time : 2019-06-28 10:57
 * @Author : ygqbasic@gmail.com
 * @File : middleware
 * @Software: VS Code
 */

package deployment

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/fuhsicloud/fuhsicloud/src/config"
	"github.com/fuhsicloud/fuhsicloud/src/repository"
)

func checkServiceMesh(cf *config.Config) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(meshRequest)
			if req.Model == repository.FieldMesh && !cf.GetBool("server", "service_mesh") {
				return nil, ErrDeploymentServiceMesh
			}
			// ErrDeploymentServiceMesh
			return next(ctx, request)
		}
	}
}
