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

type AppPostInfo struct {
	// 应用名称
	AppName string `json:"app_name"`
	// 应用备注
	AppNote string `json:"app_note,omitempty"`
}