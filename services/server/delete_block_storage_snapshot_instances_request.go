/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-08-31T04:58:16Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type DeleteBlockStorageSnapshotInstancesRequest struct {

	// 블록스토리지스냅샷인스턴스번호리스트
BlockStorageSnapshotInstanceNoList []*string `json:"blockStorageSnapshotInstanceNoList"`
}
