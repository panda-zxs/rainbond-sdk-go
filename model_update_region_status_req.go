/*
 * Rainbond Open API
 *
 * Rainbond open api
 *
 * API version: v1
 * Contact: barnett@goodrain.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UpdateRegionStatusReq struct {
	// 需要设置的集群状态, 可选值为: 'ONLINE', 'OFFLINE', 'MAINTAIN'(大小写不敏感)
	Status string `json:"status"`
}