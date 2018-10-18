/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-09-19T05:44:45Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type GetObjectStorageBackupListRequest struct {

	// 클라우드DB인스턴스번호
CloudDBInstanceNo *string `json:"cloudDBInstanceNo"`

	// 폴더이름
FolderName *string `json:"folderName,omitempty"`
}