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

type ListTeamRegionsResp struct {
	Regions []TeamRegionsResp `json:"regions"`
	Total int32 `json:"total"`
}
