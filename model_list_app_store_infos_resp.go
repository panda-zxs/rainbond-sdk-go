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

type ListAppStoreInfosResp struct {
	Appstores []AppStoreInfo `json:"appstores"`
	// 总数
	Total int32 `json:"total"`
}
