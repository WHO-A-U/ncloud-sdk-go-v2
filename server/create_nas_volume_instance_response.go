/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-07-02T09:06:17Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type CreateNasVolumeInstanceResponse struct {
	RequestId *string `json:"requestId,omitempty"`

	ReturnCode *string `json:"returnCode,omitempty"`

	ReturnMessage *string `json:"returnMessage,omitempty"`

	TotalRows *int32 `json:"totalRows,omitempty"`

	NasVolumeInstanceList *[]NasVolumeInstance `json:"nasVolumeInstanceList,omitempty"`
}