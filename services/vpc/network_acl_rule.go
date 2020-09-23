/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type NetworkAclRule struct {

	// 네트워크ACL번호
NetworkAclNo *string `json:"networkAclNo,omitempty"`

	// 우선순위
Priority *int32 `json:"priority,omitempty"`

	// 프로토콜유형
ProtocolType *CommonCode `json:"protocolType,omitempty"`

	// 포트범위
PortRange *string `json:"portRange,omitempty"`

	// Rule액션
RuleAction *CommonCode `json:"ruleAction,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// IP블록
IpBlock *string `json:"ipBlock,omitempty"`

	// 네트워크ACLRule유형
NetworkAclRuleType *CommonCode `json:"networkAclRuleType,omitempty"`

	// 네트워크ACLRule설명
NetworkAclRuleDescription *string `json:"networkAclRuleDescription,omitempty"`
}
