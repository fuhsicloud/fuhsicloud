/**
 * @Time : 2019-07-16 10:33
 * @Author : ygqbasic@gmail.com
 * @File : redis
 * @Software: VS Code
 */

package redis

import (
	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	kplredis "github.com/fuhsicloud/fuhsicloud/src/redis"
)

type adapter struct {
	rds kplredis.RedisInterface
}

func NewAdapter(rds kplredis.RedisInterface) persist.Adapter {
	return &adapter{rds: rds}
}

func (c *adapter) LoadPolicy(model model.Model) error {

	return nil
}

func (c *adapter) SavePolicy(model model.Model) error {

	return nil
}

func (c *adapter) AddPolicy(sec string, ptype string, rule []string) error {

	return nil
}

func (c *adapter) RemovePolicy(sec string, ptype string, rule []string) error {

	return nil
}

func (c *adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {

	return nil
}
