/*
 * vnas
 *
 * VPC NAS 관련 API<br/>https://ncloud.apigw.ntruss.com/vnas/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vnas

type NasVolumeSnapshot struct {

	// NAS볼륨스냅샷이름
NasVolumeSnapshotName *string `json:"nasVolumeSnapshotName;,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate;,omitempty"`

	// 스냡샷사이즈
SnapshotSize *int64 `json:"snapshotSize;,omitempty"`

	// busy여부
IsBusy *bool `json:"isBusy;,omitempty"`
}