# Go API client for autoscaling

## Installation
```
go get https://github.com/NaverCloudPlatform/ncloud-sdk-go-v2/autoscaling
```

## Documentation for API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V2Api* | [**CreateAutoScalingGroup**](docs/V2Api.md#createautoscalinggroup) | **Post** /createAutoScalingGroup |
*V2Api* | [**CreateLaunchConfiguration**](docs/V2Api.md#createlaunchconfiguration) | **Post** /createLaunchConfiguration |
*V2Api* | [**DeleteAutoScalingGroup**](docs/V2Api.md#deleteautoscalinggroup) | **Post** /deleteAutoScalingGroup |
*V2Api* | [**DeleteAutoScalingLaunchConfiguration**](docs/V2Api.md#deleteautoscalinglaunchconfiguration) | **Post** /deleteAutoScalingLaunchConfiguration |
*V2Api* | [**DeletePolicy**](docs/V2Api.md#deletepolicy) | **Post** /deletePolicy |
*V2Api* | [**DeleteScheduledAction**](docs/V2Api.md#deletescheduledaction) | **Post** /deleteScheduledAction |
*V2Api* | [**ExecutePolicy**](docs/V2Api.md#executepolicy) | **Post** /executePolicy |
*V2Api* | [**GetAdjustmentTypeList**](docs/V2Api.md#getadjustmenttypelist) | **Post** /getAdjustmentTypeList |
*V2Api* | [**GetAutoScalingActivityLogList**](docs/V2Api.md#getautoscalingactivityloglist) | **Post** /getAutoScalingActivityLogList |
*V2Api* | [**GetAutoScalingConfigurationLogList**](docs/V2Api.md#getautoscalingconfigurationloglist) | **Post** /getAutoScalingConfigurationLogList |
*V2Api* | [**GetAutoScalingGroupList**](docs/V2Api.md#getautoscalinggrouplist) | **Post** /getAutoScalingGroupList |
*V2Api* | [**GetAutoScalingPolicyList**](docs/V2Api.md#getautoscalingpolicylist) | **Post** /getAutoScalingPolicyList |
*V2Api* | [**GetLaunchConfigurationList**](docs/V2Api.md#getlaunchconfigurationlist) | **Post** /getLaunchConfigurationList |
*V2Api* | [**GetScalingProcessTypeList**](docs/V2Api.md#getscalingprocesstypelist) | **Post** /getScalingProcessTypeList |
*V2Api* | [**GetScheduledActionList**](docs/V2Api.md#getscheduledactionlist) | **Post** /getScheduledActionList |
*V2Api* | [**PutScalingPolicy**](docs/V2Api.md#putscalingpolicy) | **Post** /putScalingPolicy |
*V2Api* | [**PutScheduledUpdateGroupAction**](docs/V2Api.md#putscheduledupdategroupaction) | **Post** /putScheduledUpdateGroupAction |
*V2Api* | [**ResumeProcesses**](docs/V2Api.md#resumeprocesses) | **Post** /resumeProcesses |
*V2Api* | [**SetDesiredCapacity**](docs/V2Api.md#setdesiredcapacity) | **Post** /setDesiredCapacity |
*V2Api* | [**SetServerInstanceHealth**](docs/V2Api.md#setserverinstancehealth) | **Post** /setServerInstanceHealth |
*V2Api* | [**SuspendProcesses**](docs/V2Api.md#suspendprocesses) | **Post** /suspendProcesses |
*V2Api* | [**TerminateServerInstanceInAutoScalingGroup**](docs/V2Api.md#terminateserverinstanceinautoscalinggroup) | **Post** /terminateServerInstanceInAutoScalingGroup |
*V2Api* | [**UpdateAutoScalingGroup**](docs/V2Api.md#updateautoscalinggroup) | **Post** /updateAutoScalingGroup |


## Documentation For Models

 - [AccessControlGroup](docs/AccessControlGroup.md)
 - [ActivityLog](docs/ActivityLog.md)
 - [AdjustmentType](docs/AdjustmentType.md)
 - [AutoScalingGroup](docs/AutoScalingGroup.md)
 - [CommonCode](docs/CommonCode.md)
 - [ConfigurationLog](docs/ConfigurationLog.md)
 - [CreateAutoScalingGroupRequest](docs/CreateAutoScalingGroupRequest.md)
 - [CreateAutoScalingGroupResponse](docs/CreateAutoScalingGroupResponse.md)
 - [CreateLaunchConfigurationRequest](docs/CreateLaunchConfigurationRequest.md)
 - [CreateLaunchConfigurationResponse](docs/CreateLaunchConfigurationResponse.md)
 - [DeleteAutoScalingGroupRequest](docs/DeleteAutoScalingGroupRequest.md)
 - [DeleteAutoScalingGroupResponse](docs/DeleteAutoScalingGroupResponse.md)
 - [DeleteAutoScalingLaunchConfigurationRequest](docs/DeleteAutoScalingLaunchConfigurationRequest.md)
 - [DeleteAutoScalingLaunchConfigurationResponse](docs/DeleteAutoScalingLaunchConfigurationResponse.md)
 - [DeletePolicyRequest](docs/DeletePolicyRequest.md)
 - [DeletePolicyResponse](docs/DeletePolicyResponse.md)
 - [DeleteScheduledActionRequest](docs/DeleteScheduledActionRequest.md)
 - [DeleteScheduledActionResponse](docs/DeleteScheduledActionResponse.md)
 - [ExecutePolicyRequest](docs/ExecutePolicyRequest.md)
 - [ExecutePolicyResponse](docs/ExecutePolicyResponse.md)
 - [GetAdjustmentTypeListRequest](docs/GetAdjustmentTypeListRequest.md)
 - [GetAdjustmentTypeListResponse](docs/GetAdjustmentTypeListResponse.md)
 - [GetAutoScalingActivityLogListRequest](docs/GetAutoScalingActivityLogListRequest.md)
 - [GetAutoScalingActivityLogListResponse](docs/GetAutoScalingActivityLogListResponse.md)
 - [GetAutoScalingConfigurationLogListRequest](docs/GetAutoScalingConfigurationLogListRequest.md)
 - [GetAutoScalingConfigurationLogListResponse](docs/GetAutoScalingConfigurationLogListResponse.md)
 - [GetAutoScalingGroupListRequest](docs/GetAutoScalingGroupListRequest.md)
 - [GetAutoScalingGroupListResponse](docs/GetAutoScalingGroupListResponse.md)
 - [GetAutoScalingPolicyListRequest](docs/GetAutoScalingPolicyListRequest.md)
 - [GetAutoScalingPolicyListResponse](docs/GetAutoScalingPolicyListResponse.md)
 - [GetLaunchConfigurationListRequest](docs/GetLaunchConfigurationListRequest.md)
 - [GetLaunchConfigurationListResponse](docs/GetLaunchConfigurationListResponse.md)
 - [GetScalingProcessTypeListRequest](docs/GetScalingProcessTypeListRequest.md)
 - [GetScalingProcessTypeListResponse](docs/GetScalingProcessTypeListResponse.md)
 - [GetScheduledActionListRequest](docs/GetScheduledActionListRequest.md)
 - [GetScheduledActionListResponse](docs/GetScheduledActionListResponse.md)
 - [InAutoScalingGroupServerInstance](docs/InAutoScalingGroupServerInstance.md)
 - [LaunchConfiguration](docs/LaunchConfiguration.md)
 - [LoadBalancerInstanceSummary](docs/LoadBalancerInstanceSummary.md)
 - [Process](docs/Process.md)
 - [PutScalingPolicyRequest](docs/PutScalingPolicyRequest.md)
 - [PutScalingPolicyResponse](docs/PutScalingPolicyResponse.md)
 - [PutScheduledUpdateGroupActionRequest](docs/PutScheduledUpdateGroupActionRequest.md)
 - [PutScheduledUpdateGroupActionResponse](docs/PutScheduledUpdateGroupActionResponse.md)
 - [ResumeProcessesRequest](docs/ResumeProcessesRequest.md)
 - [ResumeProcessesResponse](docs/ResumeProcessesResponse.md)
 - [ScalingPolicy](docs/ScalingPolicy.md)
 - [ScheduledUpdateGroupAction](docs/ScheduledUpdateGroupAction.md)
 - [SetDesiredCapacityRequest](docs/SetDesiredCapacityRequest.md)
 - [SetDesiredCapacityResponse](docs/SetDesiredCapacityResponse.md)
 - [SetServerInstanceHealthRequest](docs/SetServerInstanceHealthRequest.md)
 - [SetServerInstanceHealthResponse](docs/SetServerInstanceHealthResponse.md)
 - [SuspendProcessesRequest](docs/SuspendProcessesRequest.md)
 - [SuspendProcessesResponse](docs/SuspendProcessesResponse.md)
 - [SuspendedProcess](docs/SuspendedProcess.md)
 - [TerminateServerInstanceInAutoScalingGroupRequest](docs/TerminateServerInstanceInAutoScalingGroupRequest.md)
 - [TerminateServerInstanceInAutoScalingGroupResponse](docs/TerminateServerInstanceInAutoScalingGroupResponse.md)
 - [UpdateAutoScalingGroupRequest](docs/UpdateAutoScalingGroupRequest.md)
 - [UpdateAutoScalingGroupResponse](docs/UpdateAutoScalingGroupResponse.md)
 - [Zone](docs/Zone.md)
