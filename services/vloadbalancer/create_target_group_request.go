/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type CreateTargetGroupRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 타겟그룹포트
TargetGroupPort *int32 `json:"targetGroupPort,omitempty"`

	// 타겟그룹프로토콜유형코드
TargetGroupProtocolTypeCode *string `json:"targetGroupProtocolTypeCode"`

	// 타겟그룹설명
TargetGroupDescription *string `json:"targetGroupDescription,omitempty"`

	// 헬스체크주기
HealthCheckCycle *int32 `json:"healthCheckCycle,omitempty"`

	// 헬스체크실패임계값
HealthCheckDownThreshold *int32 `json:"healthCheckDownThreshold,omitempty"`

	// 헬스체크HTTP메소드유형코드
HealthCheckHttpMethodTypeCode *string `json:"healthCheckHttpMethodTypeCode,omitempty"`

	// 헬스체크포트
HealthCheckPort *int32 `json:"healthCheckPort,omitempty"`

	// 헬스체크프로토콜유형코드
HealthCheckProtocolTypeCode *string `json:"healthCheckProtocolTypeCode"`

	// 헬스체크정상임계값
HealthCheckUpThreshold *int32 `json:"healthCheckUpThreshold,omitempty"`

	// 헬스체크URL경로
HealthCheckUrlPath *string `json:"healthCheckUrlPath,omitempty"`

	// 타겟그룹이름
TargetGroupName *string `json:"targetGroupName,omitempty"`

	// 타겟번호리스트
TargetNoList []*string `json:"targetNoList,omitempty"`

	// 타겟유형코드
TargetTypeCode *string `json:"targetTypeCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo"`
}
