/*
 * vmssql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmssql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmssql

type GetDmsOperationResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

	// DMS상태
Status *CommonCode `json:"status,omitempty"`
}