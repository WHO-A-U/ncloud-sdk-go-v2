/*
 * vserver
 *
 * VPC Compute 관련 API<br/>https://ncloud.apigw.ntruss.com/vserver/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vserver

type DeleteInitScriptsRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 초기화스크립트번호리스트
InitScriptNoList []*string `json:"initScriptNoList"`
}
