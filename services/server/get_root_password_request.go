/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-08-31T04:58:16Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetRootPasswordRequest struct {

	// 개인키
PrivateKey *string `json:"privateKey,omitempty"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`
}
