/*
 * vmssql
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmssql/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmssql

type GetDmsObjectStorageBackupListResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

	// DMS파일리스트
DmsFileList *DmsFileList `json:"dmsFileList,omitempty"`
}