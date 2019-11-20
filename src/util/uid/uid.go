/**
 * @Time : 2019-08-15 14:08
 * @Author : ygqbasic@gmail.com
 * @File : uid
 * @Software: VS Code
 */

package uid

import "github.com/google/uuid"

type Uid string

const (
	RequestId Uid = "uid-request-id"
)

func GenerateUid() string {
	id, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return id.String()
}
