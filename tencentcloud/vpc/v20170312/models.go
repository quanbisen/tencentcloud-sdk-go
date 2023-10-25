// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20170312

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AcceptAttachCcnInstancesRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 接受关联实例列表。
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

type AcceptAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 接受关联实例列表。
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

func (r *AcceptAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptAttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptAttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptAttachCcnInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AcceptAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AcceptAttachCcnInstancesResponseParams `json:"Response"`
}

func (r *AcceptAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptVpcPeeringConnectionRequestParams struct {
	// 对等连接唯一ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`
}

type AcceptVpcPeeringConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 对等连接唯一ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`
}

func (r *AcceptVpcPeeringConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PeeringConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcceptVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AcceptVpcPeeringConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AcceptVpcPeeringConnectionResponse struct {
	*tchttp.BaseResponse
	Response *AcceptVpcPeeringConnectionResponseParams `json:"Response"`
}

func (r *AcceptVpcPeeringConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcceptVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessPolicy struct {
	// 目的CIDR
	TargetCidr *string `json:"TargetCidr,omitnil" name:"TargetCidr"`

	// 策略ID
	VpnGatewayIdSslAccessPolicyId *string `json:"VpnGatewayIdSslAccessPolicyId,omitnil" name:"VpnGatewayIdSslAccessPolicyId"`

	// 是否对所有用户都生效。1 生效 0不生效
	ForAllClient *uint64 `json:"ForAllClient,omitnil" name:"ForAllClient"`

	// 用户组ID
	UserGroupIds []*string `json:"UserGroupIds,omitnil" name:"UserGroupIds"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Remark
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type AccountAttribute struct {
	// 属性名
	AttributeName *string `json:"AttributeName,omitnil" name:"AttributeName"`

	// 属性值
	AttributeValues []*string `json:"AttributeValues,omitnil" name:"AttributeValues"`
}

// Predefined struct for user
type AddBandwidthPackageResourcesRequestParams struct {
	// 资源唯一ID，当前支持EIP资源和LB资源，形如'eip-xxxx', 'lb-xxxx'
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 带宽包唯一标识ID，形如'bwp-xxxx'
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 带宽包类型，当前支持'BGP'类型，表示内部资源是BGP IP。
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// 资源类型，包括'Address', 'LoadBalance'
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 带宽包协议类型。当前支持'ipv4'和'ipv6'协议类型。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

type AddBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 资源唯一ID，当前支持EIP资源和LB资源，形如'eip-xxxx', 'lb-xxxx'
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 带宽包唯一标识ID，形如'bwp-xxxx'
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 带宽包类型，当前支持'BGP'类型，表示内部资源是BGP IP。
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// 资源类型，包括'Address', 'LoadBalance'
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 带宽包协议类型。当前支持'ipv4'和'ipv6'协议类型。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`
}

func (r *AddBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceIds")
	delete(f, "BandwidthPackageId")
	delete(f, "NetworkType")
	delete(f, "ResourceType")
	delete(f, "Protocol")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddBandwidthPackageResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddBandwidthPackageResourcesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *AddBandwidthPackageResourcesResponseParams `json:"Response"`
}

func (r *AddBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddIp6RulesRequestParams struct {
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// IPV6转换规则信息
	Ip6RuleInfos []*Ip6RuleInfo `json:"Ip6RuleInfos,omitnil" name:"Ip6RuleInfos"`

	// IPV6转换规则名称
	Ip6RuleName *string `json:"Ip6RuleName,omitnil" name:"Ip6RuleName"`
}

type AddIp6RulesRequest struct {
	*tchttp.BaseRequest
	
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// IPV6转换规则信息
	Ip6RuleInfos []*Ip6RuleInfo `json:"Ip6RuleInfos,omitnil" name:"Ip6RuleInfos"`

	// IPV6转换规则名称
	Ip6RuleName *string `json:"Ip6RuleName,omitnil" name:"Ip6RuleName"`
}

func (r *AddIp6RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddIp6RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorId")
	delete(f, "Ip6RuleInfos")
	delete(f, "Ip6RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddIp6RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddIp6RulesResponseParams struct {
	// IPV6转换规则唯一ID数组，形如rule6-xxxxxxxx
	Ip6RuleSet []*string `json:"Ip6RuleSet,omitnil" name:"Ip6RuleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddIp6RulesResponse struct {
	*tchttp.BaseResponse
	Response *AddIp6RulesResponseParams `json:"Response"`
}

func (r *AddIp6RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddIp6RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTemplateMemberRequestParams struct {
	// 参数模板实例ID，支持IP地址、协议端口、IP地址组、协议端口组四种参数模板的实例ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 需要添加的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致。
	TemplateMember []*MemberInfo `json:"TemplateMember,omitnil" name:"TemplateMember"`
}

type AddTemplateMemberRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板实例ID，支持IP地址、协议端口、IP地址组、协议端口组四种参数模板的实例ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 需要添加的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致。
	TemplateMember []*MemberInfo `json:"TemplateMember,omitnil" name:"TemplateMember"`
}

func (r *AddTemplateMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTemplateMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateMember")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTemplateMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTemplateMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddTemplateMemberResponse struct {
	*tchttp.BaseResponse
	Response *AddTemplateMemberResponseParams `json:"Response"`
}

func (r *AddTemplateMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTemplateMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Address struct {
	// `EIP`的`ID`，是`EIP`的唯一标识。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// `EIP`名称。
	AddressName *string `json:"AddressName,omitnil" name:"AddressName"`

	// `EIP`状态，包含'CREATING'(创建中),'BINDING'(绑定中),'BIND'(已绑定),'UNBINDING'(解绑中),'UNBIND'(已解绑),'OFFLINING'(释放中),'BIND_ENI'(绑定悬空弹性网卡)
	AddressStatus *string `json:"AddressStatus,omitnil" name:"AddressStatus"`

	// 外网IP地址
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`

	// 绑定的资源实例`ID`。可能是一个`CVM`，`NAT`。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 绑定的弹性网卡ID
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 绑定的资源内网ip
	PrivateAddressIp *string `json:"PrivateAddressIp,omitnil" name:"PrivateAddressIp"`

	// 资源隔离状态。true表示eip处于隔离状态，false表示资源处于未隔离状态
	IsArrears *bool `json:"IsArrears,omitnil" name:"IsArrears"`

	// 资源封堵状态。true表示eip处于封堵状态，false表示eip处于未封堵状态
	IsBlocked *bool `json:"IsBlocked,omitnil" name:"IsBlocked"`

	// eip是否支持直通模式。true表示eip支持直通模式，false表示资源不支持直通模式
	IsEipDirectConnection *bool `json:"IsEipDirectConnection,omitnil" name:"IsEipDirectConnection"`

	// EIP 资源类型，包括CalcIP、WanIP、EIP和AnycastEIP、高防EIP。其中：`CalcIP` 表示设备 IP，`WanIP` 表示普通公网 IP，`EIP` 表示弹性公网 IP，`AnycastEIP` 表示加速 EIP，`AntiDDoSEIP`表示高防EIP。
	AddressType *string `json:"AddressType,omitnil" name:"AddressType"`

	// eip是否在解绑后自动释放。true表示eip将会在解绑后自动释放，false表示eip在解绑后不会自动释放
	CascadeRelease *bool `json:"CascadeRelease,omitnil" name:"CascadeRelease"`

	// EIP ALG开启的协议类型。
	EipAlgType *AlgType `json:"EipAlgType,omitnil" name:"EipAlgType"`

	// 弹性公网IP的运营商信息，当前可能返回值包括"CMCC","CTCC","CUCC","BGP"
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil" name:"InternetServiceProvider"`

	// 是否本地带宽EIP
	LocalBgp *bool `json:"LocalBgp,omitnil" name:"LocalBgp"`

	// 弹性公网IP的带宽值。注意，传统账户类型账户的弹性公网IP没有带宽属性，值为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *uint64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 弹性公网IP的网络计费模式。注意，传统账户类型账户的弹性公网IP没有网络计费模式属性，值为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 包括：
	// <li><strong>BANDWIDTH_PREPAID_BY_MONTH</strong></li>
	// <p style="padding-left: 30px;">表示包月带宽预付费。</p>
	// <li><strong>TRAFFIC_POSTPAID_BY_HOUR</strong></li>
	// <p style="padding-left: 30px;">表示按小时流量后付费。</p>
	// <li><strong>BANDWIDTH_POSTPAID_BY_HOUR</strong></li>
	// <p style="padding-left: 30px;">表示按小时带宽后付费。</p>
	// <li><strong>BANDWIDTH_PACKAGE</strong></li>
	// <p style="padding-left: 30px;">表示共享带宽包。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// 弹性公网IP关联的标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 到期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadlineDate *string `json:"DeadlineDate,omitnil" name:"DeadlineDate"`

	// EIP绑定的实例类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 静态单线IP网络出口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Egress *string `json:"Egress,omitnil" name:"Egress"`

	// 高防包ID,当EIP类型为高防EIP时，返回EIP绑定的高防包ID.
	AntiDDoSPackageId *string `json:"AntiDDoSPackageId,omitnil" name:"AntiDDoSPackageId"`
}

type AddressChargePrepaid struct {
	// 购买实例的时长，单位是月。可支持时长：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 自动续费标志。0表示手动续费，1表示自动续费，2表示到期不续费。默认缺省为0即手动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`
}

type AddressInfo struct {
	// ip地址。
	Address *string `json:"Address,omitnil" name:"Address"`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type AddressTemplate struct {
	// IP地址模板名称。
	AddressTemplateName *string `json:"AddressTemplateName,omitnil" name:"AddressTemplateName"`

	// IP地址模板实例唯一ID。
	AddressTemplateId *string `json:"AddressTemplateId,omitnil" name:"AddressTemplateId"`

	// IP地址信息。
	AddressSet []*string `json:"AddressSet,omitnil" name:"AddressSet"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 带备注的IP地址信息。
	AddressExtraSet []*AddressInfo `json:"AddressExtraSet,omitnil" name:"AddressExtraSet"`
}

type AddressTemplateGroup struct {
	// IP地址模板集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil" name:"AddressTemplateGroupName"`

	// IP地址模板集合实例ID，例如：ipmg-dih8xdbq。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil" name:"AddressTemplateGroupId"`

	// IP地址模板ID。
	AddressTemplateIdSet []*string `json:"AddressTemplateIdSet,omitnil" name:"AddressTemplateIdSet"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// IP地址模板实例。
	AddressTemplateSet []*AddressTemplateItem `json:"AddressTemplateSet,omitnil" name:"AddressTemplateSet"`
}

type AddressTemplateItem struct {
	// ipm-xxxxxxxx
	AddressTemplateId *string `json:"AddressTemplateId,omitnil" name:"AddressTemplateId"`

	// IP模板名称
	AddressTemplateName *string `json:"AddressTemplateName,omitnil" name:"AddressTemplateName"`

	// 废弃字段
	From *string `json:"From,omitnil" name:"From"`

	// 废弃字段
	To *string `json:"To,omitnil" name:"To"`
}

type AddressTemplateSpecification struct {
	// IP地址ID，例如：ipm-2uw6ujo6。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// IP地址组ID，例如：ipmg-2uw6ujo6。
	AddressGroupId *string `json:"AddressGroupId,omitnil" name:"AddressGroupId"`
}

// Predefined struct for user
type AdjustPublicAddressRequestParams struct {
	// 标识CVM实例的唯一 ID。CVM 唯一 ID 形如：`ins-11112222`。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 标识EIP实例的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`
}

type AdjustPublicAddressRequest struct {
	*tchttp.BaseRequest
	
	// 标识CVM实例的唯一 ID。CVM 唯一 ID 形如：`ins-11112222`。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 标识EIP实例的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`
}

func (r *AdjustPublicAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustPublicAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AdjustPublicAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AdjustPublicAddressResponseParams struct {
	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AdjustPublicAddressResponse struct {
	*tchttp.BaseResponse
	Response *AdjustPublicAddressResponseParams `json:"Response"`
}

func (r *AdjustPublicAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AdjustPublicAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlgType struct {
	// Ftp协议Alg功能是否开启
	Ftp *bool `json:"Ftp,omitnil" name:"Ftp"`

	// Sip协议Alg功能是否开启
	Sip *bool `json:"Sip,omitnil" name:"Sip"`
}

// Predefined struct for user
type AllocateAddressesRequestParams struct {
	// EIP数量。默认值：1。
	AddressCount *int64 `json:"AddressCount,omitnil" name:"AddressCount"`

	// EIP线路类型。默认值：BGP。
	// <ul style="margin:0"><li>已开通静态单线IP白名单的用户，可选值：<ul><li>CMCC：中国移动</li>
	// <li>CTCC：中国电信</li>
	// <li>CUCC：中国联通</li></ul>注意：仅部分地域支持静态单线IP。</li></ul>
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil" name:"InternetServiceProvider"`

	// EIP计费方式。
	// <ul style="margin:0"><li>已开通标准账户类型白名单的用户，可选值：<ul><li>BANDWIDTH_PACKAGE：[共享带宽包](https://cloud.tencent.com/document/product/684/15255)付费（需额外开通共享带宽包白名单）</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH：包月按带宽预付费</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费</li></ul>默认值：TRAFFIC_POSTPAID_BY_HOUR。</li>
	// <li>未开通标准账户类型白名单的用户，EIP计费方式与其绑定的实例的计费方式一致，无需传递此参数。</li></ul>
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// EIP出带宽上限，单位：Mbps。
	// <ul style="margin:0"><li>已开通标准账户类型白名单的用户，可选值范围取决于EIP计费方式：<ul><li>BANDWIDTH_PACKAGE：1 Mbps 至 2000 Mbps</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH：1 Mbps 至 200 Mbps</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li></ul>默认值：1 Mbps。</li>
	// <li>未开通标准账户类型白名单的用户，EIP出带宽上限取决于与其绑定的实例的公网出带宽上限，无需传递此参数。</li></ul>
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 包月按带宽预付费EIP的计费参数。EIP为包月按带宽预付费时，该参数必传，其余场景不需传递
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil" name:"AddressChargePrepaid"`

	// EIP类型。默认值：EIP。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>AnycastEIP：加速IP，可参见 [Anycast 公网加速](https://cloud.tencent.com/document/product/644)</li></ul>注意：仅部分地域支持加速IP。</li></ul>
	// <ul style="margin:0"><li>已开通精品IP白名单的用户，可选值：<ul><li>HighQualityEIP：精品IP</li></ul>注意：仅部分地域支持精品IP。</li></ul>
	// <ul style="margin:0"><li>已开高防IP白名单的用户，可选值：<ul><li>AntiDDoSEIP：高防IP</li></ul>注意：仅部分地域支持高防IP。</li></ul>
	AddressType *string `json:"AddressType,omitnil" name:"AddressType"`

	// Anycast发布域。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>ANYCAST_ZONE_GLOBAL：全球发布域（需要额外开通Anycast全球加速白名单）</li><li>ANYCAST_ZONE_OVERSEAS：境外发布域</li><li><b>[已废弃]</b> ANYCAST_ZONE_A：发布域A（已更新为全球发布域）</li><li><b>[已废弃]</b> ANYCAST_ZONE_B：发布域B（已更新为全球发布域）</li></ul>默认值：ANYCAST_ZONE_OVERSEAS。</li></ul>
	AnycastZone *string `json:"AnycastZone,omitnil" name:"AnycastZone"`

	// <b>[已废弃]</b> AnycastEIP不再区分是否负载均衡。原参数说明如下：
	// AnycastEIP是否用于绑定负载均衡。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>TRUE：AnycastEIP可绑定对象为负载均衡</li>
	// <li>FALSE：AnycastEIP可绑定对象为云服务器、NAT网关、高可用虚拟IP等</li></ul>默认值：FALSE。</li></ul>
	ApplicableForCLB *bool `json:"ApplicableForCLB,omitnil" name:"ApplicableForCLB"`

	// 需要关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// BGP带宽包唯一ID参数。设定该参数且InternetChargeType为BANDWIDTH_PACKAGE，则表示创建的EIP加入该BGP带宽包并采用带宽包计费
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// EIP名称，用于申请EIP时用户自定义该EIP的个性化名称，默认值：未命名
	AddressName *string `json:"AddressName,omitnil" name:"AddressName"`

	// 静态单线IP网络出口，默认值：center_egress1
	Egress *string `json:"Egress,omitnil" name:"Egress"`

	// 高防包ID， 申请高防IP时，该字段必传。
	AntiDDoSPackageId *string `json:"AntiDDoSPackageId,omitnil" name:"AntiDDoSPackageId"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type AllocateAddressesRequest struct {
	*tchttp.BaseRequest
	
	// EIP数量。默认值：1。
	AddressCount *int64 `json:"AddressCount,omitnil" name:"AddressCount"`

	// EIP线路类型。默认值：BGP。
	// <ul style="margin:0"><li>已开通静态单线IP白名单的用户，可选值：<ul><li>CMCC：中国移动</li>
	// <li>CTCC：中国电信</li>
	// <li>CUCC：中国联通</li></ul>注意：仅部分地域支持静态单线IP。</li></ul>
	InternetServiceProvider *string `json:"InternetServiceProvider,omitnil" name:"InternetServiceProvider"`

	// EIP计费方式。
	// <ul style="margin:0"><li>已开通标准账户类型白名单的用户，可选值：<ul><li>BANDWIDTH_PACKAGE：[共享带宽包](https://cloud.tencent.com/document/product/684/15255)付费（需额外开通共享带宽包白名单）</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH：包月按带宽预付费</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费</li></ul>默认值：TRAFFIC_POSTPAID_BY_HOUR。</li>
	// <li>未开通标准账户类型白名单的用户，EIP计费方式与其绑定的实例的计费方式一致，无需传递此参数。</li></ul>
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// EIP出带宽上限，单位：Mbps。
	// <ul style="margin:0"><li>已开通标准账户类型白名单的用户，可选值范围取决于EIP计费方式：<ul><li>BANDWIDTH_PACKAGE：1 Mbps 至 2000 Mbps</li>
	// <li>BANDWIDTH_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li>
	// <li>BANDWIDTH_PREPAID_BY_MONTH：1 Mbps 至 200 Mbps</li>
	// <li>TRAFFIC_POSTPAID_BY_HOUR：1 Mbps 至 100 Mbps</li></ul>默认值：1 Mbps。</li>
	// <li>未开通标准账户类型白名单的用户，EIP出带宽上限取决于与其绑定的实例的公网出带宽上限，无需传递此参数。</li></ul>
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 包月按带宽预付费EIP的计费参数。EIP为包月按带宽预付费时，该参数必传，其余场景不需传递
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil" name:"AddressChargePrepaid"`

	// EIP类型。默认值：EIP。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>AnycastEIP：加速IP，可参见 [Anycast 公网加速](https://cloud.tencent.com/document/product/644)</li></ul>注意：仅部分地域支持加速IP。</li></ul>
	// <ul style="margin:0"><li>已开通精品IP白名单的用户，可选值：<ul><li>HighQualityEIP：精品IP</li></ul>注意：仅部分地域支持精品IP。</li></ul>
	// <ul style="margin:0"><li>已开高防IP白名单的用户，可选值：<ul><li>AntiDDoSEIP：高防IP</li></ul>注意：仅部分地域支持高防IP。</li></ul>
	AddressType *string `json:"AddressType,omitnil" name:"AddressType"`

	// Anycast发布域。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>ANYCAST_ZONE_GLOBAL：全球发布域（需要额外开通Anycast全球加速白名单）</li><li>ANYCAST_ZONE_OVERSEAS：境外发布域</li><li><b>[已废弃]</b> ANYCAST_ZONE_A：发布域A（已更新为全球发布域）</li><li><b>[已废弃]</b> ANYCAST_ZONE_B：发布域B（已更新为全球发布域）</li></ul>默认值：ANYCAST_ZONE_OVERSEAS。</li></ul>
	AnycastZone *string `json:"AnycastZone,omitnil" name:"AnycastZone"`

	// <b>[已废弃]</b> AnycastEIP不再区分是否负载均衡。原参数说明如下：
	// AnycastEIP是否用于绑定负载均衡。
	// <ul style="margin:0"><li>已开通Anycast公网加速白名单的用户，可选值：<ul><li>TRUE：AnycastEIP可绑定对象为负载均衡</li>
	// <li>FALSE：AnycastEIP可绑定对象为云服务器、NAT网关、高可用虚拟IP等</li></ul>默认值：FALSE。</li></ul>
	ApplicableForCLB *bool `json:"ApplicableForCLB,omitnil" name:"ApplicableForCLB"`

	// 需要关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// BGP带宽包唯一ID参数。设定该参数且InternetChargeType为BANDWIDTH_PACKAGE，则表示创建的EIP加入该BGP带宽包并采用带宽包计费
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// EIP名称，用于申请EIP时用户自定义该EIP的个性化名称，默认值：未命名
	AddressName *string `json:"AddressName,omitnil" name:"AddressName"`

	// 静态单线IP网络出口，默认值：center_egress1
	Egress *string `json:"Egress,omitnil" name:"Egress"`

	// 高防包ID， 申请高防IP时，该字段必传。
	AntiDDoSPackageId *string `json:"AntiDDoSPackageId,omitnil" name:"AntiDDoSPackageId"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *AllocateAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressCount")
	delete(f, "InternetServiceProvider")
	delete(f, "InternetChargeType")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "AddressChargePrepaid")
	delete(f, "AddressType")
	delete(f, "AnycastZone")
	delete(f, "ApplicableForCLB")
	delete(f, "Tags")
	delete(f, "BandwidthPackageId")
	delete(f, "AddressName")
	delete(f, "Egress")
	delete(f, "AntiDDoSPackageId")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateAddressesResponseParams struct {
	// 申请到的 EIP 的唯一 ID 列表。
	AddressSet []*string `json:"AddressSet,omitnil" name:"AddressSet"`

	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AllocateAddressesResponse struct {
	*tchttp.BaseResponse
	Response *AllocateAddressesResponseParams `json:"Response"`
}

func (r *AllocateAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateIp6AddressesBandwidthRequestParams struct {
	// 需要开通公网访问能力的IPV6地址
	Ip6Addresses []*string `json:"Ip6Addresses,omitnil" name:"Ip6Addresses"`

	// 带宽，单位Mbps。默认是1Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 网络计费模式。IPV6当前对标准账户类型支持"TRAFFIC_POSTPAID_BY_HOUR"，对传统账户类型支持"BANDWIDTH_PACKAGE"。默认网络计费模式是"TRAFFIC_POSTPAID_BY_HOUR"。
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// 带宽包id，上移账号，申请带宽包计费模式的ipv6地址需要传入.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`
}

type AllocateIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// 需要开通公网访问能力的IPV6地址
	Ip6Addresses []*string `json:"Ip6Addresses,omitnil" name:"Ip6Addresses"`

	// 带宽，单位Mbps。默认是1Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 网络计费模式。IPV6当前对标准账户类型支持"TRAFFIC_POSTPAID_BY_HOUR"，对传统账户类型支持"BANDWIDTH_PACKAGE"。默认网络计费模式是"TRAFFIC_POSTPAID_BY_HOUR"。
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// 带宽包id，上移账号，申请带宽包计费模式的ipv6地址需要传入.
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`
}

func (r *AllocateIp6AddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6Addresses")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "InternetChargeType")
	delete(f, "BandwidthPackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateIp6AddressesBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateIp6AddressesBandwidthResponseParams struct {
	// 弹性公网 IPV6 的唯一 ID 列表。
	AddressSet []*string `json:"AddressSet,omitnil" name:"AddressSet"`

	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AllocateIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *AllocateIp6AddressesBandwidthResponseParams `json:"Response"`
}

func (r *AllocateIp6AddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6AddressesRequestParams struct {
	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的`IPv6`地址列表，单次最多指定10个。与入参`Ipv6AddressCount`合并计算配额。与Ipv6AddressCount必填一个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`

	// 自动分配`IPv6`地址个数，内网IP地址个数总和不能超过配额数。与入参`Ipv6Addresses`合并计算配额。与Ipv6Addresses必填一个。
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitnil" name:"Ipv6AddressCount"`
}

type AssignIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的`IPv6`地址列表，单次最多指定10个。与入参`Ipv6AddressCount`合并计算配额。与Ipv6AddressCount必填一个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`

	// 自动分配`IPv6`地址个数，内网IP地址个数总和不能超过配额数。与入参`Ipv6Addresses`合并计算配额。与Ipv6Addresses必填一个。
	Ipv6AddressCount *uint64 `json:"Ipv6AddressCount,omitnil" name:"Ipv6AddressCount"`
}

func (r *AssignIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	delete(f, "Ipv6AddressCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6AddressesResponseParams struct {
	// 分配给弹性网卡的`IPv6`地址列表。
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitnil" name:"Ipv6AddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssignIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *AssignIpv6AddressesResponseParams `json:"Response"`
}

func (r *AssignIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6CidrBlockRequestParams struct {
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type AssignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

func (r *AssignIpv6CidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6CidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6CidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6CidrBlockResponseParams struct {
	// 分配的 `IPv6` 网段。形如：`3402:4e00:20:1000::/56`。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *AssignIpv6CidrBlockResponseParams `json:"Response"`
}

func (r *AssignIpv6CidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6CidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6SubnetCidrBlockRequestParams struct {
	// 子网所在私有网络`ID`。形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 分配 `IPv6` 子网段列表。
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil" name:"Ipv6SubnetCidrBlocks"`
}

type AssignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// 子网所在私有网络`ID`。形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 分配 `IPv6` 子网段列表。
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil" name:"Ipv6SubnetCidrBlocks"`
}

func (r *AssignIpv6SubnetCidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6SubnetCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignIpv6SubnetCidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignIpv6SubnetCidrBlockResponseParams struct {
	// 分配 `IPv6` 子网段列表。
	Ipv6SubnetCidrBlockSet []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlockSet,omitnil" name:"Ipv6SubnetCidrBlockSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *AssignIpv6SubnetCidrBlockResponseParams `json:"Response"`
}

func (r *AssignIpv6SubnetCidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IP信息，单次最多指定10个。与SecondaryPrivateIpAddressCount至少提供一个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 新申请的内网IP地址个数，与PrivateIpAddresses至少提供一个。内网IP地址个数总和不能超过配额数，详见<a href="/document/product/576/18527">弹性网卡使用限制</a>。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`

	// IP服务质量等级，和SecondaryPrivateIpAddressCount配合使用，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type AssignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IP信息，单次最多指定10个。与SecondaryPrivateIpAddressCount至少提供一个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 新申请的内网IP地址个数，与PrivateIpAddresses至少提供一个。内网IP地址个数总和不能超过配额数，详见<a href="/document/product/576/18527">弹性网卡使用限制</a>。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`

	// IP服务质量等级，和SecondaryPrivateIpAddressCount配合使用，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *AssignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignPrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "QosLevel")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignPrivateIpAddressesResponseParams struct {
	// 内网IP详细信息。
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitnil" name:"PrivateIpAddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *AssignPrivateIpAddressesResponseParams `json:"Response"`
}

func (r *AssignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssistantCidr struct {
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 辅助CIDR。形如：`172.16.0.0/16`
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 辅助CIDR类型（0：普通辅助CIDR，1：容器辅助CIDR），默认都是0。
	AssistantType *int64 `json:"AssistantType,omitnil" name:"AssistantType"`

	// 辅助CIDR拆分的子网。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetSet []*Subnet `json:"SubnetSet,omitnil" name:"SubnetSet"`
}

// Predefined struct for user
type AssociateAddressRequestParams struct {
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 要绑定的实例 ID。实例 ID 形如：`ins-11112222`、`lb-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：`eni-11112222`。`NetworkInterfaceId` 与 `InstanceId` 不可同时指定。弹性网卡 ID 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/15817)接口返回值中的`networkInterfaceId`获取。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 要绑定的内网 IP。如果指定了 `NetworkInterfaceId` 则也必须指定 `PrivateIpAddress` ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 `PrivateIpAddress` 是指定的 `NetworkInterfaceId` 上的一个内网 IP。指定弹性网卡的内网 IP 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/15817)接口返回值中的`privateIpAddress`获取。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 指定绑定时是否设置直通。弹性公网 IP 直通请参见 [EIP 直通](https://cloud.tencent.com/document/product/1199/41709)。取值：True、False，默认值为 False。当绑定 CVM 实例、EKS 弹性集群时，可设定此参数为 True。此参数目前处于内测中，如需使用，请提交 [工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20CLB&level3_id=1071&queue=96&scene_code=34639&step=2)。
	EipDirectConnection *bool `json:"EipDirectConnection,omitnil" name:"EipDirectConnection"`
}

type AssociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 要绑定的实例 ID。实例 ID 形如：`ins-11112222`、`lb-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 要绑定的弹性网卡 ID。 弹性网卡 ID 形如：`eni-11112222`。`NetworkInterfaceId` 与 `InstanceId` 不可同时指定。弹性网卡 ID 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/15817)接口返回值中的`networkInterfaceId`获取。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 要绑定的内网 IP。如果指定了 `NetworkInterfaceId` 则也必须指定 `PrivateIpAddress` ，表示将 EIP 绑定到指定弹性网卡的指定内网 IP 上。同时要确保指定的 `PrivateIpAddress` 是指定的 `NetworkInterfaceId` 上的一个内网 IP。指定弹性网卡的内网 IP 可通过登录[控制台](https://console.cloud.tencent.com/vpc/eni)查询，也可通过[DescribeNetworkInterfaces](https://cloud.tencent.com/document/api/215/15817)接口返回值中的`privateIpAddress`获取。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 指定绑定时是否设置直通。弹性公网 IP 直通请参见 [EIP 直通](https://cloud.tencent.com/document/product/1199/41709)。取值：True、False，默认值为 False。当绑定 CVM 实例、EKS 弹性集群时，可设定此参数为 True。此参数目前处于内测中，如需使用，请提交 [工单申请](https://console.cloud.tencent.com/workorder/category?level1_id=6&level2_id=163&source=0&data_title=%E8%B4%9F%E8%BD%BD%E5%9D%87%E8%A1%A1%20CLB&level3_id=1071&queue=96&scene_code=34639&step=2)。
	EipDirectConnection *bool `json:"EipDirectConnection,omitnil" name:"EipDirectConnection"`
}

func (r *AssociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressId")
	delete(f, "InstanceId")
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	delete(f, "EipDirectConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateAddressResponseParams struct {
	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *AssociateAddressResponseParams `json:"Response"`
}

func (r *AssociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDhcpIpWithAddressIpRequestParams struct {
	// `DhcpIp`唯一`ID`，形如：`dhcpip-9o233uri`。必须是没有绑定`EIP`的`DhcpIp`
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`

	// 弹性公网`IP`。必须是没有绑定`DhcpIp`的`EIP`
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`
}

type AssociateDhcpIpWithAddressIpRequest struct {
	*tchttp.BaseRequest
	
	// `DhcpIp`唯一`ID`，形如：`dhcpip-9o233uri`。必须是没有绑定`EIP`的`DhcpIp`
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`

	// 弹性公网`IP`。必须是没有绑定`DhcpIp`的`EIP`
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`
}

func (r *AssociateDhcpIpWithAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDhcpIpWithAddressIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DhcpIpId")
	delete(f, "AddressIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDhcpIpWithAddressIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDhcpIpWithAddressIpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateDhcpIpWithAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDhcpIpWithAddressIpResponseParams `json:"Response"`
}

func (r *AssociateDhcpIpWithAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDhcpIpWithAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDirectConnectGatewayNatGatewayRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关ID。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 专线网关ID。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`
}

type AssociateDirectConnectGatewayNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关ID。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 专线网关ID。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`
}

func (r *AssociateDirectConnectGatewayNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDirectConnectGatewayNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NatGatewayId")
	delete(f, "DirectConnectGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDirectConnectGatewayNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDirectConnectGatewayNatGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateDirectConnectGatewayNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDirectConnectGatewayNatGatewayResponseParams `json:"Response"`
}

func (r *AssociateDirectConnectGatewayNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDirectConnectGatewayNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNatGatewayAddressRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 需要申请的弹性IP个数，系统会按您的要求生产N个弹性IP, 其中AddressCount和PublicAddresses至少传递一个。
	AddressCount *uint64 `json:"AddressCount,omitnil" name:"AddressCount"`

	// 绑定NAT网关的弹性IP数组，其中AddressCount和PublicAddresses至少传递一个。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`

	// 弹性IP可用区，自动分配弹性IP时传递。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 绑定NAT网关的弹性IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil" name:"StockPublicIpAddressesBandwidthOut"`

	// 需要申请公网IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil" name:"PublicIpAddressesBandwidthOut"`

	// 公网IP是否强制与NAT网关来自同可用区，true表示需要与NAT网关同可用区；false表示可与NAT网关不是同一个可用区。此参数只有当参数Zone存在时才能生效。
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil" name:"PublicIpFromSameZone"`
}

type AssociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 需要申请的弹性IP个数，系统会按您的要求生产N个弹性IP, 其中AddressCount和PublicAddresses至少传递一个。
	AddressCount *uint64 `json:"AddressCount,omitnil" name:"AddressCount"`

	// 绑定NAT网关的弹性IP数组，其中AddressCount和PublicAddresses至少传递一个。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`

	// 弹性IP可用区，自动分配弹性IP时传递。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 绑定NAT网关的弹性IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil" name:"StockPublicIpAddressesBandwidthOut"`

	// 需要申请公网IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil" name:"PublicIpAddressesBandwidthOut"`

	// 公网IP是否强制与NAT网关来自同可用区，true表示需要与NAT网关同可用区；false表示可与NAT网关不是同一个可用区。此参数只有当参数Zone存在时才能生效。
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil" name:"PublicIpFromSameZone"`
}

func (r *AssociateNatGatewayAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNatGatewayAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "AddressCount")
	delete(f, "PublicIpAddresses")
	delete(f, "Zone")
	delete(f, "StockPublicIpAddressesBandwidthOut")
	delete(f, "PublicIpAddressesBandwidthOut")
	delete(f, "PublicIpFromSameZone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateNatGatewayAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNatGatewayAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse
	Response *AssociateNatGatewayAddressResponseParams `json:"Response"`
}

func (r *AssociateNatGatewayAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNatGatewayAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkAclSubnetsRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 子网实例ID数组。例如：[subnet-12345678]。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`
}

type AssociateNetworkAclSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 子网实例ID数组。例如：[subnet-12345678]。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`
}

func (r *AssociateNetworkAclSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkAclSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateNetworkAclSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkAclSubnetsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateNetworkAclSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateNetworkAclSubnetsResponseParams `json:"Response"`
}

func (r *AssociateNetworkAclSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkAclSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkInterfaceSecurityGroupsRequestParams struct {
	// 弹性网卡实例ID。形如：eni-pxir56ns。每次请求的实例的上限为100。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。每次请求的实例的上限为100。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type AssociateNetworkInterfaceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID。形如：eni-pxir56ns。每次请求的实例的上限为100。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。每次请求的实例的上限为100。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *AssociateNetworkInterfaceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkInterfaceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateNetworkInterfaceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateNetworkInterfaceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AssociateNetworkInterfaceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateNetworkInterfaceSecurityGroupsResponseParams `json:"Response"`
}

func (r *AssociateNetworkInterfaceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateNetworkInterfaceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCcnInstancesRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 关联网络实例列表
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`

	// CCN所属UIN（根账号），默认当前账号所属UIN
	CcnUin *string `json:"CcnUin,omitnil" name:"CcnUin"`
}

type AttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 关联网络实例列表
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`

	// CCN所属UIN（根账号），默认当前账号所属UIN
	CcnUin *string `json:"CcnUin,omitnil" name:"CcnUin"`
}

func (r *AttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	delete(f, "CcnUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachCcnInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachCcnInstancesResponseParams `json:"Response"`
}

func (r *AttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachClassicLinkVpcRequestParams struct {
	// VPC实例ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CVM实例ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type AttachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CVM实例ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *AttachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachClassicLinkVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachClassicLinkVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachClassicLinkVpcResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AttachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *AttachClassicLinkVpcResponseParams `json:"Response"`
}

func (r *AttachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachClassicLinkVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachNetworkInterfaceRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// CVM实例ID。形如：ins-r8hr2upy。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 网卡的挂载类型：0 标准型，1扩展型，默认值0。
	AttachType *uint64 `json:"AttachType,omitnil" name:"AttachType"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type AttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// CVM实例ID。形如：ins-r8hr2upy。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 网卡的挂载类型：0 标准型，1扩展型，默认值0。
	AttachType *uint64 `json:"AttachType,omitnil" name:"AttachType"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *AttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "InstanceId")
	delete(f, "AttachType")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *AttachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *AttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachSnapshotInstancesRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 关联实例信息。
	Instances []*SnapshotInstance `json:"Instances,omitnil" name:"Instances"`
}

type AttachSnapshotInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 关联实例信息。
	Instances []*SnapshotInstance `json:"Instances,omitnil" name:"Instances"`
}

func (r *AttachSnapshotInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachSnapshotInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachSnapshotInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachSnapshotInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AttachSnapshotInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachSnapshotInstancesResponseParams `json:"Response"`
}

func (r *AttachSnapshotInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachSnapshotInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuditCrossBorderComplianceRequestParams struct {
	// 服务商, 可选值：`UNICOM`。
	ServiceProvider *string `json:"ServiceProvider,omitnil" name:"ServiceProvider"`

	// 表单唯一`ID`。可通过[DescribeCrossBorderCompliance](https://cloud.tencent.com/document/product/215/47838)接口查询ComplianceId信息
	ComplianceId *uint64 `json:"ComplianceId,omitnil" name:"ComplianceId"`

	// 通过：`APPROVED `，拒绝：`DENY`。
	AuditBehavior *string `json:"AuditBehavior,omitnil" name:"AuditBehavior"`
}

type AuditCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest
	
	// 服务商, 可选值：`UNICOM`。
	ServiceProvider *string `json:"ServiceProvider,omitnil" name:"ServiceProvider"`

	// 表单唯一`ID`。可通过[DescribeCrossBorderCompliance](https://cloud.tencent.com/document/product/215/47838)接口查询ComplianceId信息
	ComplianceId *uint64 `json:"ComplianceId,omitnil" name:"ComplianceId"`

	// 通过：`APPROVED `，拒绝：`DENY`。
	AuditBehavior *string `json:"AuditBehavior,omitnil" name:"AuditBehavior"`
}

func (r *AuditCrossBorderComplianceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuditCrossBorderComplianceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProvider")
	delete(f, "ComplianceId")
	delete(f, "AuditBehavior")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuditCrossBorderComplianceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuditCrossBorderComplianceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AuditCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse
	Response *AuditCrossBorderComplianceResponseParams `json:"Response"`
}

func (r *AuditCrossBorderComplianceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuditCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BackupPolicy struct {
	// 备份周期时间，取值为monday, tuesday, wednesday, thursday, friday, saturday, sunday。
	BackupDay *string `json:"BackupDay,omitnil" name:"BackupDay"`

	// 备份时间点，格式：HH:mm:ss。
	BackupTime *string `json:"BackupTime,omitnil" name:"BackupTime"`
}

type BandwidthPackage struct {
	// 带宽包唯一标识Id
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 带宽包类型，包括'BGP','SINGLEISP','ANYCAST','SINGLEISP_CMCC','SINGLEISP_CTCC','SINGLEISP_CUCC'
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// 带宽包计费类型，包括'TOP5_POSTPAID_BY_MONTH'和'PERCENT95_POSTPAID_BY_MONTH'
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// 带宽包名称
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil" name:"BandwidthPackageName"`

	// 带宽包创建时间。按照`ISO8601`标准表示，并且使用`UTC`时间。格式为：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 带宽包状态，包括'CREATING','CREATED','DELETING','DELETED'
	Status *string `json:"Status,omitnil" name:"Status"`

	// 带宽包资源信息
	ResourceSet []*Resource `json:"ResourceSet,omitnil" name:"ResourceSet"`

	// 带宽包限速大小。单位：Mbps，-1表示不限速。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 网络出口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Egress *string `json:"Egress,omitnil" name:"Egress"`
}

type BandwidthPackageBillBandwidth struct {
	// 当前计费用量，单位为 Mbps
	BandwidthUsage *float64 `json:"BandwidthUsage,omitnil" name:"BandwidthUsage"`
}

type BatchModifySnapshotPolicy struct {
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 快照策略名称。
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil" name:"SnapshotPolicyName"`

	// 备份策略。
	BackupPolicies []*BackupPolicy `json:"BackupPolicies,omitnil" name:"BackupPolicies"`

	// 快照保留时间，支持1～365天。
	KeepTime *uint64 `json:"KeepTime,omitnil" name:"KeepTime"`
}

type CCN struct {
	// 云联网唯一ID
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网名称
	CcnName *string `json:"CcnName,omitnil" name:"CcnName"`

	// 云联网描述信息
	CcnDescription *string `json:"CcnDescription,omitnil" name:"CcnDescription"`

	// 关联实例数量
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 实例状态， 'ISOLATED': 隔离中（欠费停服），'AVAILABLE'：运行中。
	State *string `json:"State,omitnil" name:"State"`

	// 实例服务质量，’PT’：白金，'AU'：金，'AG'：银。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 付费类型，PREPAID为预付费，POSTPAID为后付费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 限速类型，`INTER_REGION_LIMIT` 为地域间限速；`OUTER_REGION_LIMIT` 为地域出口限速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil" name:"BandwidthLimitType"`

	// 标签键值对。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 是否支持云联网路由优先级的功能。`False`：不支持，`True`：支持。
	RoutePriorityFlag *bool `json:"RoutePriorityFlag,omitnil" name:"RoutePriorityFlag"`

	// 实例关联的路由表个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableCount *uint64 `json:"RouteTableCount,omitnil" name:"RouteTableCount"`

	// 是否开启云联网多路由表特性。`False`：未开启，`True`：开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableFlag *bool `json:"RouteTableFlag,omitnil" name:"RouteTableFlag"`

	// `true`：实例已被封禁，流量不通，`false`:解封禁。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSecurityLock *bool `json:"IsSecurityLock,omitnil" name:"IsSecurityLock"`

	// 是否开启云联网路由传播策略。`False` 未开启，`True` 开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteBroadcastPolicyFlag *bool `json:"RouteBroadcastPolicyFlag,omitnil" name:"RouteBroadcastPolicyFlag"`
}

type CcnAttachedInstance struct {
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 关联实例类型：
	// <li>`VPC`：私有网络</li>
	// <li>`DIRECTCONNECT`：专线网关</li>
	// <li>`BMVPC`：黑石私有网络</li>
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 关联实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 关联实例名称。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 关联实例所属大区，例如：ap-guangzhou。
	InstanceRegion *string `json:"InstanceRegion,omitnil" name:"InstanceRegion"`

	// 关联实例所属UIN（根账号）。
	InstanceUin *string `json:"InstanceUin,omitnil" name:"InstanceUin"`

	// 关联实例CIDR。
	CidrBlock []*string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 关联实例状态：
	// <li>`PENDING`：申请中</li>
	// <li>`ACTIVE`：已连接</li>
	// <li>`EXPIRED`：已过期</li>
	// <li>`REJECTED`：已拒绝</li>
	// <li>`DELETED`：已删除</li>
	// <li>`FAILED`：失败的（2小时后将异步强制解关联）</li>
	// <li>`ATTACHING`：关联中</li>
	// <li>`DETACHING`：解关联中</li>
	// <li>`DETACHFAILED`：解关联失败（2小时后将异步强制解关联）</li>
	State *string `json:"State,omitnil" name:"State"`

	// 关联时间。
	AttachedTime *string `json:"AttachedTime,omitnil" name:"AttachedTime"`

	// 云联网所属UIN（根账号）。
	CcnUin *string `json:"CcnUin,omitnil" name:"CcnUin"`

	// 关联实例所属的大地域，如: CHINA_MAINLAND
	InstanceArea *string `json:"InstanceArea,omitnil" name:"InstanceArea"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`

	// 路由表ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`
}

type CcnBandwidth struct {
	// 带宽所属的云联网ID。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 实例的过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 带宽实例的唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionFlowControlId *string `json:"RegionFlowControlId,omitnil" name:"RegionFlowControlId"`

	// 带宽是否自动续费的标记。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 描述带宽的地域和限速上限信息。在地域间限速的情况下才会返回参数，出口限速模式不返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnRegionBandwidthLimit *CcnRegionBandwidthLimitInfo `json:"CcnRegionBandwidthLimit,omitnil" name:"CcnRegionBandwidthLimit"`

	// 云市场实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarketId *string `json:"MarketId,omitnil" name:"MarketId"`

	// 实例所属用户主账号ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAccountID *string `json:"UserAccountID,omitnil" name:"UserAccountID"`

	// 是否跨境，`true`表示跨境，反之不跨境。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCrossBorder *bool `json:"IsCrossBorder,omitnil" name:"IsCrossBorder"`

	// `true`表示封禁，地域间流量不通，`false`解禁，地域间流量正常
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSecurityLock *bool `json:"IsSecurityLock,omitnil" name:"IsSecurityLock"`

	// `POSTPAID`表示后付费，`PREPAID`表示预付费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 实例更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type CcnBandwidthInfo struct {
	// 带宽所属的云联网ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 实例的过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 带宽实例的唯一ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionFlowControlId *string `json:"RegionFlowControlId,omitnil" name:"RegionFlowControlId"`

	// 带宽是否自动续费的标记。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 描述带宽的地域和限速上限信息。在地域间限速的情况下才会返回参数，出口限速模式不返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnRegionBandwidthLimit *CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimit,omitnil" name:"CcnRegionBandwidthLimit"`

	// 云市场实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarketId *string `json:"MarketId,omitnil" name:"MarketId"`

	// 资源绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`
}

type CcnFlowLock struct {
	// 带宽所属的云联网ID。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 实例所属用户主账号ID。
	UserAccountID *string `json:"UserAccountID,omitnil" name:"UserAccountID"`

	// 带宽实例的唯一ID。作为`UnlockCcnBandwidths`接口和`LockCcnBandwidths`接口的入参时，该字段必传。
	RegionFlowControlId *string `json:"RegionFlowControlId,omitnil" name:"RegionFlowControlId"`
}

type CcnInstance struct {
	// 关联实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 关联实例ID所属大区，例如：ap-guangzhou。
	InstanceRegion *string `json:"InstanceRegion,omitnil" name:"InstanceRegion"`

	// 关联实例类型，可选值：
	// <li>`VPC`：私有网络</li>
	// <li>`DIRECTCONNECT`：专线网关</li>
	// <li>`BMVPC`：黑石私有网络</li>
	// <li>`VPNGW`：VPNGW类型</li>
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`

	// 实例关联的路由表ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`
}

type CcnInstanceInfo struct {

}

type CcnRegionBandwidthLimit struct {
	// 地域，例如：ap-guangzhou
	Region *string `json:"Region,omitnil" name:"Region"`

	// 出带宽上限，单位：Mbps
	BandwidthLimit *uint64 `json:"BandwidthLimit,omitnil" name:"BandwidthLimit"`

	// 是否黑石地域，默认`false`。
	IsBm *bool `json:"IsBm,omitnil" name:"IsBm"`

	// 目的地域，例如：ap-shanghai
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 目的地域是否为黑石地域，默认`false`。
	DstIsBm *bool `json:"DstIsBm,omitnil" name:"DstIsBm"`
}

type CcnRegionBandwidthLimitInfo struct {
	// 源地域，例如：ap-shanghai
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceRegion *string `json:"SourceRegion,omitnil" name:"SourceRegion"`

	// 目的地域， 例如：ap-shanghai
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestinationRegion *string `json:"DestinationRegion,omitnil" name:"DestinationRegion"`

	// 出带宽上限，单位：Mbps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthLimit *uint64 `json:"BandwidthLimit,omitnil" name:"BandwidthLimit"`
}

type CcnRoute struct {
	// 路由策略ID
	RouteId *string `json:"RouteId,omitnil" name:"RouteId"`

	// 目的端
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil" name:"DestinationCidrBlock"`

	// 下一跳类型（关联实例类型），所有类型：VPC、DIRECTCONNECT
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 下一跳（关联实例）
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 下一跳名称（关联实例名称）
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 下一跳所属地域（关联实例所属地域）
	InstanceRegion *string `json:"InstanceRegion,omitnil" name:"InstanceRegion"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 路由是否启用
	Enabled *bool `json:"Enabled,omitnil" name:"Enabled"`

	// 关联实例所属UIN（根账号）
	InstanceUin *string `json:"InstanceUin,omitnil" name:"InstanceUin"`

	// 路由的扩展状态
	ExtraState *string `json:"ExtraState,omitnil" name:"ExtraState"`

	// 是否动态路由
	IsBgp *bool `json:"IsBgp,omitnil" name:"IsBgp"`

	// 路由优先级
	RoutePriority *uint64 `json:"RoutePriority,omitnil" name:"RoutePriority"`

	// 下一跳扩展名称（关联实例的扩展名称）
	InstanceExtraName *string `json:"InstanceExtraName,omitnil" name:"InstanceExtraName"`
}

// Predefined struct for user
type CheckAssistantCidrRequestParams struct {
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 待添加的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]。入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil" name:"NewCidrBlocks"`

	// 待删除的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]。入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil" name:"OldCidrBlocks"`
}

type CheckAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 待添加的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]。入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil" name:"NewCidrBlocks"`

	// 待删除的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]。入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil" name:"OldCidrBlocks"`
}

func (r *CheckAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NewCidrBlocks")
	delete(f, "OldCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAssistantCidrResponseParams struct {
	// 冲突资源信息数组。
	ConflictSourceSet []*ConflictSource `json:"ConflictSourceSet,omitnil" name:"ConflictSourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *CheckAssistantCidrResponseParams `json:"Response"`
}

func (r *CheckAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDefaultSubnetRequestParams struct {
	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

type CheckDefaultSubnetRequest struct {
	*tchttp.BaseRequest
	
	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone *string `json:"Zone,omitnil" name:"Zone"`
}

func (r *CheckDefaultSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDefaultSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDefaultSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDefaultSubnetResponseParams struct {
	// 检查结果。true为可以创建默认子网，false为不可以创建默认子网。
	Result *bool `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckDefaultSubnetResponse struct {
	*tchttp.BaseResponse
	Response *CheckDefaultSubnetResponseParams `json:"Response"`
}

func (r *CheckDefaultSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDefaultSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckNetDetectStateRequestParams struct {
	// 探测目的IPv4地址数组，最多两个。
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 网络探测实例ID。形如：netd-12345678。该参数与（VpcId，SubnetId，NetDetectName），至少要有一个。当NetDetectId存在时，使用NetDetectId。
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`

	// `VPC`实例`ID`。形如：`vpc-12345678`。该参数与（SubnetId，NetDetectName）配合使用，与NetDetectId至少要有一个。当NetDetectId存在时，使用NetDetectId。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。形如：subnet-12345678。该参数与（VpcId，NetDetectName）配合使用，与NetDetectId至少要有一个。当NetDetectId存在时，使用NetDetectId。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 网络探测名称，最大长度不能超过60个字节。该参数与（VpcId，SubnetId）配合使用，与NetDetectId至少要有一个。当NetDetectId存在时，使用NetDetectId。
	NetDetectName *string `json:"NetDetectName,omitnil" name:"NetDetectName"`

	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云服务器；
	// CCN：云联网网关；
	// NONEXTHOP：无下一跳；
	NextHopType *string `json:"NextHopType,omitnil" name:"NextHopType"`

	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云服务器IPv4地址，形如：10.0.0.12；
	// 下一跳类型为CCN，取值云联网ID，形如：ccn-12345678；
	// 下一跳类型为NONEXTHOP，指定网络探测为无下一跳的网络探测；
	NextHopDestination *string `json:"NextHopDestination,omitnil" name:"NextHopDestination"`
}

type CheckNetDetectStateRequest struct {
	*tchttp.BaseRequest
	
	// 探测目的IPv4地址数组，最多两个。
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 网络探测实例ID。形如：netd-12345678。该参数与（VpcId，SubnetId，NetDetectName），至少要有一个。当NetDetectId存在时，使用NetDetectId。
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`

	// `VPC`实例`ID`。形如：`vpc-12345678`。该参数与（SubnetId，NetDetectName）配合使用，与NetDetectId至少要有一个。当NetDetectId存在时，使用NetDetectId。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。形如：subnet-12345678。该参数与（VpcId，NetDetectName）配合使用，与NetDetectId至少要有一个。当NetDetectId存在时，使用NetDetectId。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 网络探测名称，最大长度不能超过60个字节。该参数与（VpcId，SubnetId）配合使用，与NetDetectId至少要有一个。当NetDetectId存在时，使用NetDetectId。
	NetDetectName *string `json:"NetDetectName,omitnil" name:"NetDetectName"`

	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云服务器；
	// CCN：云联网网关；
	// NONEXTHOP：无下一跳；
	NextHopType *string `json:"NextHopType,omitnil" name:"NextHopType"`

	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云服务器IPv4地址，形如：10.0.0.12；
	// 下一跳类型为CCN，取值云联网ID，形如：ccn-12345678；
	// 下一跳类型为NONEXTHOP，指定网络探测为无下一跳的网络探测；
	NextHopDestination *string `json:"NextHopDestination,omitnil" name:"NextHopDestination"`
}

func (r *CheckNetDetectStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNetDetectStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DetectDestinationIp")
	delete(f, "NetDetectId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "NetDetectName")
	delete(f, "NextHopType")
	delete(f, "NextHopDestination")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckNetDetectStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckNetDetectStateResponseParams struct {
	// 网络探测验证结果对象数组。
	NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitnil" name:"NetDetectIpStateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckNetDetectStateResponse struct {
	*tchttp.BaseResponse
	Response *CheckNetDetectStateResponseParams `json:"Response"`
}

func (r *CheckNetDetectStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckNetDetectStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CidrForCcn struct {
	// local cidr值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitnil" name:"Cidr"`

	// 是否发布到了云联网。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishedToVbc *bool `json:"PublishedToVbc,omitnil" name:"PublishedToVbc"`
}

type ClassicLinkInstance struct {
	// VPC实例ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 云服务器实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

// Predefined struct for user
type CloneSecurityGroupRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。未提供参数时，克隆后的安全组名称和SecurityGroupId对应的安全组名称相同。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。未提供参数时，克隆后的安全组备注和SecurityGroupId对应的安全组备注相同。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`

	// 项目ID，默认0。可在<a href="https://console.cloud.tencent.com/project">qcloud控制台项目管理页面</a>查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 源Region,跨地域克隆安全组时，需要传入源安全组所属地域信息，例如：克隆广州的安全组到上海，则这里需要传入广州安全的地域信息：ap-guangzhou。
	RemoteRegion *string `json:"RemoteRegion,omitnil" name:"RemoteRegion"`
}

type CloneSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。未提供参数时，克隆后的安全组名称和SecurityGroupId对应的安全组名称相同。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。未提供参数时，克隆后的安全组备注和SecurityGroupId对应的安全组备注相同。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`

	// 项目ID，默认0。可在<a href="https://console.cloud.tencent.com/project">qcloud控制台项目管理页面</a>查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 源Region,跨地域克隆安全组时，需要传入源安全组所属地域信息，例如：克隆广州的安全组到上海，则这里需要传入广州安全的地域信息：ap-guangzhou。
	RemoteRegion *string `json:"RemoteRegion,omitnil" name:"RemoteRegion"`
}

func (r *CloneSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	delete(f, "ProjectId")
	delete(f, "RemoteRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneSecurityGroupResponseParams struct {
	// 安全组对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CloneSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CloneSecurityGroupResponseParams `json:"Response"`
}

func (r *CloneSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConflictItem struct {
	// 冲突资源的ID
	ConfilctId *string `json:"ConfilctId,omitnil" name:"ConfilctId"`

	// 冲突目的资源
	DestinationItem *string `json:"DestinationItem,omitnil" name:"DestinationItem"`
}

type ConflictSource struct {
	// 冲突资源ID
	ConflictSourceId *string `json:"ConflictSourceId,omitnil" name:"ConflictSourceId"`

	// 冲突资源
	SourceItem *string `json:"SourceItem,omitnil" name:"SourceItem"`

	// 冲突资源条目信息
	ConflictItemSet []*ConflictItem `json:"ConflictItemSet,omitnil" name:"ConflictItemSet"`
}

// Predefined struct for user
type CreateAddressTemplateGroupRequestParams struct {
	// IP地址模板集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil" name:"AddressTemplateGroupName"`

	// IP地址模板实例ID，例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil" name:"AddressTemplateIds"`
}

type CreateAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// IP地址模板集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil" name:"AddressTemplateGroupName"`

	// IP地址模板实例ID，例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil" name:"AddressTemplateIds"`
}

func (r *CreateAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateGroupName")
	delete(f, "AddressTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateGroupResponseParams struct {
	// IP地址模板集合对象。
	AddressTemplateGroup *AddressTemplateGroup `json:"AddressTemplateGroup,omitnil" name:"AddressTemplateGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateAddressTemplateGroupResponseParams `json:"Response"`
}

func (r *CreateAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateRequestParams struct {
	// IP地址模板名称。
	AddressTemplateName *string `json:"AddressTemplateName,omitnil" name:"AddressTemplateName"`

	// 地址信息，支持 IP、CIDR、IP 范围。Addresses与AddressesExtra必填其一。
	Addresses []*string `json:"Addresses,omitnil" name:"Addresses"`

	// 地址信息，支持携带备注，支持 IP、CIDR、IP 范围。Addresses与AddressesExtra必填其一。
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil" name:"AddressesExtra"`
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// IP地址模板名称。
	AddressTemplateName *string `json:"AddressTemplateName,omitnil" name:"AddressTemplateName"`

	// 地址信息，支持 IP、CIDR、IP 范围。Addresses与AddressesExtra必填其一。
	Addresses []*string `json:"Addresses,omitnil" name:"Addresses"`

	// 地址信息，支持携带备注，支持 IP、CIDR、IP 范围。Addresses与AddressesExtra必填其一。
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil" name:"AddressesExtra"`
}

func (r *CreateAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateName")
	delete(f, "Addresses")
	delete(f, "AddressesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateResponseParams struct {
	// IP地址模板对象。
	AddressTemplate *AddressTemplate `json:"AddressTemplate,omitnil" name:"AddressTemplate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAddressTemplateResponseParams `json:"Response"`
}

func (r *CreateAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndAttachNetworkInterfaceRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil" name:"NetworkInterfaceName"`

	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 云服务器实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配额数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`

	// IP服务质量等级，和SecondaryPrivateIpAddressCount配合使用，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil" name:"NetworkInterfaceDescription"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 绑定类型：0 标准型 1 扩展型。
	AttachType *uint64 `json:"AttachType,omitnil" name:"AttachType"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type CreateAndAttachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil" name:"NetworkInterfaceName"`

	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 云服务器实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配额数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`

	// IP服务质量等级，和SecondaryPrivateIpAddressCount配合使用，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil" name:"NetworkInterfaceDescription"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 绑定类型：0 标准型 1 扩展型。
	AttachType *uint64 `json:"AttachType,omitnil" name:"AttachType"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *CreateAndAttachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndAttachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NetworkInterfaceName")
	delete(f, "SubnetId")
	delete(f, "InstanceId")
	delete(f, "PrivateIpAddresses")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "QosLevel")
	delete(f, "SecurityGroupIds")
	delete(f, "NetworkInterfaceDescription")
	delete(f, "Tags")
	delete(f, "AttachType")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAndAttachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAndAttachNetworkInterfaceResponseParams struct {
	// 弹性网卡实例。
	NetworkInterface *NetworkInterface `json:"NetworkInterface,omitnil" name:"NetworkInterface"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAndAttachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateAndAttachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *CreateAndAttachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAndAttachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssistantCidrRequestParams struct {
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitnil" name:"CidrBlocks"`
}

type CreateAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]
	CidrBlocks []*string `json:"CidrBlocks,omitnil" name:"CidrBlocks"`
}

func (r *CreateAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "CidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssistantCidrResponseParams struct {
	// 辅助CIDR数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil" name:"AssistantCidrSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssistantCidrResponseParams `json:"Response"`
}

func (r *CreateAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBandwidthPackageRequestParams struct {
	// 带宽包类型, 默认值: BGP, 可选值:
	// <li>BGP: 普通BGP共享带宽包</li>
	// <li>HIGH_QUALITY_BGP: 精品BGP共享带宽包</li>
	// <li>SINGLEISP_CMCC: 中国移动共享带宽包</li>
	// <li>SINGLEISP_CTCC: 中国电信共享带宽包</li>
	// <li>SINGLEISP_CUCC: 中国联通共享带宽包</li>
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// 带宽包计费类型, 默认为: TOP5_POSTPAID_BY_MONTH, 可选值:
	// <li>TOP5_POSTPAID_BY_MONTH: 按月后付费TOP5计费</li>
	// <li>PERCENT95_POSTPAID_BY_MONTH: 按月后付费月95计费</li>
	// <li>FIXED_PREPAID_BY_MONTH: 包月预付费计费</li>
	// <li>ENHANCED95_POSTPAID_BY_MONTH: 按月后付费增强型95计费</li>
	// <li>PEAK_BANDWIDTH_POSTPAID_BY_DAY: 后付费日结按带宽计费</li>
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// 带宽包名称。
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil" name:"BandwidthPackageName"`

	// 带宽包数量(传统账户类型只能填1), 标准账户类型取值范围为1~20。
	BandwidthPackageCount *uint64 `json:"BandwidthPackageCount,omitnil" name:"BandwidthPackageCount"`

	// 带宽包限速大小。单位：Mbps，-1表示不限速。该功能当前内测中，暂不对外开放。
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitnil" name:"InternetMaxBandwidth"`

	// 需要关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 带宽包协议类型。当前支持'ipv4'和'ipv6'协议带宽包，默认值是'ipv4'。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 预付费包月带宽包的购买时长，单位: 月，取值范围: 1~60。
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 网络出口，默认值：center_egress1
	Egress *string `json:"Egress,omitnil" name:"Egress"`
}

type CreateBandwidthPackageRequest struct {
	*tchttp.BaseRequest
	
	// 带宽包类型, 默认值: BGP, 可选值:
	// <li>BGP: 普通BGP共享带宽包</li>
	// <li>HIGH_QUALITY_BGP: 精品BGP共享带宽包</li>
	// <li>SINGLEISP_CMCC: 中国移动共享带宽包</li>
	// <li>SINGLEISP_CTCC: 中国电信共享带宽包</li>
	// <li>SINGLEISP_CUCC: 中国联通共享带宽包</li>
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// 带宽包计费类型, 默认为: TOP5_POSTPAID_BY_MONTH, 可选值:
	// <li>TOP5_POSTPAID_BY_MONTH: 按月后付费TOP5计费</li>
	// <li>PERCENT95_POSTPAID_BY_MONTH: 按月后付费月95计费</li>
	// <li>FIXED_PREPAID_BY_MONTH: 包月预付费计费</li>
	// <li>ENHANCED95_POSTPAID_BY_MONTH: 按月后付费增强型95计费</li>
	// <li>PEAK_BANDWIDTH_POSTPAID_BY_DAY: 后付费日结按带宽计费</li>
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// 带宽包名称。
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil" name:"BandwidthPackageName"`

	// 带宽包数量(传统账户类型只能填1), 标准账户类型取值范围为1~20。
	BandwidthPackageCount *uint64 `json:"BandwidthPackageCount,omitnil" name:"BandwidthPackageCount"`

	// 带宽包限速大小。单位：Mbps，-1表示不限速。该功能当前内测中，暂不对外开放。
	InternetMaxBandwidth *int64 `json:"InternetMaxBandwidth,omitnil" name:"InternetMaxBandwidth"`

	// 需要关联的标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 带宽包协议类型。当前支持'ipv4'和'ipv6'协议带宽包，默认值是'ipv4'。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 预付费包月带宽包的购买时长，单位: 月，取值范围: 1~60。
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 网络出口，默认值：center_egress1
	Egress *string `json:"Egress,omitnil" name:"Egress"`
}

func (r *CreateBandwidthPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBandwidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkType")
	delete(f, "ChargeType")
	delete(f, "BandwidthPackageName")
	delete(f, "BandwidthPackageCount")
	delete(f, "InternetMaxBandwidth")
	delete(f, "Tags")
	delete(f, "Protocol")
	delete(f, "TimeSpan")
	delete(f, "Egress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBandwidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBandwidthPackageResponseParams struct {
	// 带宽包唯一ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 带宽包唯一ID列表(申请数量大于1时有效)。
	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitnil" name:"BandwidthPackageIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateBandwidthPackageResponse struct {
	*tchttp.BaseResponse
	Response *CreateBandwidthPackageResponseParams `json:"Response"`
}

func (r *CreateBandwidthPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcnRequestParams struct {
	// CCN名称，最大长度不能超过60个字节。
	CcnName *string `json:"CcnName,omitnil" name:"CcnName"`

	// CCN描述信息，最大长度不能超过100个字节。
	CcnDescription *string `json:"CcnDescription,omitnil" name:"CcnDescription"`

	// CCN服务质量，`PT`：白金，`AU`：金，`AG`：银，默认为`AU`。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 计费模式，`PREPAID`：表示预付费，即包年包月，`POSTPAID`：表示后付费，即按量计费。默认：`POSTPAID`。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 限速类型，`OUTER_REGION_LIMIT`表示地域出口限速，`INTER_REGION_LIMIT`为地域间限速，默认为`OUTER_REGION_LIMIT`。预付费模式仅支持地域间限速，后付费模式支持地域间限速和地域出口限速。
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil" name:"BandwidthLimitType"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateCcnRequest struct {
	*tchttp.BaseRequest
	
	// CCN名称，最大长度不能超过60个字节。
	CcnName *string `json:"CcnName,omitnil" name:"CcnName"`

	// CCN描述信息，最大长度不能超过100个字节。
	CcnDescription *string `json:"CcnDescription,omitnil" name:"CcnDescription"`

	// CCN服务质量，`PT`：白金，`AU`：金，`AG`：银，默认为`AU`。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 计费模式，`PREPAID`：表示预付费，即包年包月，`POSTPAID`：表示后付费，即按量计费。默认：`POSTPAID`。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 限速类型，`OUTER_REGION_LIMIT`表示地域出口限速，`INTER_REGION_LIMIT`为地域间限速，默认为`OUTER_REGION_LIMIT`。预付费模式仅支持地域间限速，后付费模式支持地域间限速和地域出口限速。
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil" name:"BandwidthLimitType"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnName")
	delete(f, "CcnDescription")
	delete(f, "QosLevel")
	delete(f, "InstanceChargeType")
	delete(f, "BandwidthLimitType")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCcnResponseParams struct {
	// 云联网（CCN）对象。
	Ccn *CCN `json:"Ccn,omitnil" name:"Ccn"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCcnResponse struct {
	*tchttp.BaseResponse
	Response *CreateCcnResponseParams `json:"Response"`
}

func (r *CreateCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerGatewayRequestParams struct {
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil" name:"CustomerGatewayName"`

	// 对端网关公网IP。
	IpAddress *string `json:"IpAddress,omitnil" name:"IpAddress"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateCustomerGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil" name:"CustomerGatewayName"`

	// 对端网关公网IP。
	IpAddress *string `json:"IpAddress,omitnil" name:"IpAddress"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayName")
	delete(f, "IpAddress")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerGatewayResponseParams struct {
	// 对端网关对象
	CustomerGateway *CustomerGateway `json:"CustomerGateway,omitnil" name:"CustomerGateway"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomerGatewayResponseParams `json:"Response"`
}

func (r *CreateCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultSecurityGroupRequestParams struct {
	// 项目ID，默认0。可在qcloud控制台项目管理页面查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

type CreateDefaultSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID，默认0。可在qcloud控制台项目管理页面查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`
}

func (r *CreateDefaultSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefaultSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultSecurityGroupResponseParams struct {
	// 安全组对象。
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDefaultSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateDefaultSecurityGroupResponseParams `json:"Response"`
}

func (r *CreateDefaultSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultVpcRequestParams struct {
	// 子网所在的可用区，该参数可通过[DescribeZones](https://cloud.tencent.com/document/product/213/15707)接口获取，例如ap-guangzhou-1，不指定时将随机选择可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 是否强制返回默认VPC。
	Force *bool `json:"Force,omitnil" name:"Force"`
}

type CreateDefaultVpcRequest struct {
	*tchttp.BaseRequest
	
	// 子网所在的可用区，该参数可通过[DescribeZones](https://cloud.tencent.com/document/product/213/15707)接口获取，例如ap-guangzhou-1，不指定时将随机选择可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 是否强制返回默认VPC。
	Force *bool `json:"Force,omitnil" name:"Force"`
}

func (r *CreateDefaultVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Zone")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDefaultVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDefaultVpcResponseParams struct {
	// 默认VPC和子网ID。
	Vpc *DefaultVpcSubnet `json:"Vpc,omitnil" name:"Vpc"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDefaultVpcResponse struct {
	*tchttp.BaseResponse
	Response *CreateDefaultVpcResponseParams `json:"Response"`
}

func (r *CreateDefaultVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDefaultVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDhcpIpRequestParams struct {
	// 私有网络`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网`ID`。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// `DhcpIp`名称。
	DhcpIpName *string `json:"DhcpIpName,omitnil" name:"DhcpIpName"`

	// 新申请的内网IP地址个数。总数不能超过64个，为了兼容性，当前参数必填。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`
}

type CreateDhcpIpRequest struct {
	*tchttp.BaseRequest
	
	// 私有网络`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网`ID`。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// `DhcpIp`名称。
	DhcpIpName *string `json:"DhcpIpName,omitnil" name:"DhcpIpName"`

	// 新申请的内网IP地址个数。总数不能超过64个，为了兼容性，当前参数必填。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`
}

func (r *CreateDhcpIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDhcpIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "DhcpIpName")
	delete(f, "SecondaryPrivateIpAddressCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDhcpIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDhcpIpResponseParams struct {
	// 新创建的`DhcpIp`信息。
	DhcpIpSet []*DhcpIp `json:"DhcpIpSet,omitnil" name:"DhcpIpSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDhcpIpResponse struct {
	*tchttp.BaseResponse
	Response *CreateDhcpIpResponseParams `json:"Response"`
}

func (r *CreateDhcpIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDhcpIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayCcnRoutesRequestParams struct {
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 需要连通的IDC网段列表
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil" name:"Routes"`
}

type CreateDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 需要连通的IDC网段列表
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil" name:"Routes"`
}

func (r *CreateDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayCcnRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *CreateDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayRequestParams struct {
	// 专线网关名称
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil" name:"DirectConnectGatewayName"`

	// 关联网络类型，可选值：
	// <li>VPC - 私有网络</li>
	// <li>CCN - 云联网</li>
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// <li>NetworkType 为 VPC 时，这里传值为私有网络实例ID</li>
	// <li>NetworkType 为 CCN 时，这里传值为云联网实例ID</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil" name:"NetworkInstanceId"`

	// 网关类型，可选值：
	// <li>NORMAL - （默认）标准型，注：云联网只支持标准型</li>
	// <li>NAT - NAT型</li>NAT类型支持网络地址转换配置，类型确定后不能修改；一个私有网络可以创建一个NAT类型的专线网关和一个非NAT类型的专线网关
	GatewayType *string `json:"GatewayType,omitnil" name:"GatewayType"`

	// 云联网路由发布模式，可选值：`standard`（标准模式）、`exquisite`（精细模式）。只有云联网类型专线网关才支持`ModeType`。
	ModeType *string `json:"ModeType,omitnil" name:"ModeType"`

	// 专线网关可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 专线网关高可用区容灾组ID
	HaZoneGroupId *string `json:"HaZoneGroupId,omitnil" name:"HaZoneGroupId"`
}

type CreateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关名称
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil" name:"DirectConnectGatewayName"`

	// 关联网络类型，可选值：
	// <li>VPC - 私有网络</li>
	// <li>CCN - 云联网</li>
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// <li>NetworkType 为 VPC 时，这里传值为私有网络实例ID</li>
	// <li>NetworkType 为 CCN 时，这里传值为云联网实例ID</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil" name:"NetworkInstanceId"`

	// 网关类型，可选值：
	// <li>NORMAL - （默认）标准型，注：云联网只支持标准型</li>
	// <li>NAT - NAT型</li>NAT类型支持网络地址转换配置，类型确定后不能修改；一个私有网络可以创建一个NAT类型的专线网关和一个非NAT类型的专线网关
	GatewayType *string `json:"GatewayType,omitnil" name:"GatewayType"`

	// 云联网路由发布模式，可选值：`standard`（标准模式）、`exquisite`（精细模式）。只有云联网类型专线网关才支持`ModeType`。
	ModeType *string `json:"ModeType,omitnil" name:"ModeType"`

	// 专线网关可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 专线网关高可用区容灾组ID
	HaZoneGroupId *string `json:"HaZoneGroupId,omitnil" name:"HaZoneGroupId"`
}

func (r *CreateDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayName")
	delete(f, "NetworkType")
	delete(f, "NetworkInstanceId")
	delete(f, "GatewayType")
	delete(f, "ModeType")
	delete(f, "Zone")
	delete(f, "HaZoneGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDirectConnectGatewayResponseParams struct {
	// 专线网关对象。
	DirectConnectGateway *DirectConnectGateway `json:"DirectConnectGateway,omitnil" name:"DirectConnectGateway"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateDirectConnectGatewayResponseParams `json:"Response"`
}

func (r *CreateDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowLogRequestParams struct {
	// 流日志实例名字。
	FlowLogName *string `json:"FlowLogName,omitnil" name:"FlowLogName"`

	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE|CCN|NAT|DCG。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源唯一ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流日志采集类型，ACCEPT|REJECT|ALL。
	TrafficType *string `json:"TrafficType,omitnil" name:"TrafficType"`

	// 私用网络ID或者统一ID，建议使用统一ID，当ResourceType为CCN时不填，其他类型必填。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志实例描述。
	FlowLogDescription *string `json:"FlowLogDescription,omitnil" name:"FlowLogDescription"`

	// 流日志存储ID。
	CloudLogId *string `json:"CloudLogId,omitnil" name:"CloudLogId"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 消费端类型：cls、ckafka。默认值cls。
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// 流日志消费端信息，当消费端类型为ckafka时，必填。
	FlowLogStorage *FlowLogStorage `json:"FlowLogStorage,omitnil" name:"FlowLogStorage"`

	// 流日志存储ID对应的地域，不传递默认为本地域。
	CloudLogRegion *string `json:"CloudLogRegion,omitnil" name:"CloudLogRegion"`
}

type CreateFlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 流日志实例名字。
	FlowLogName *string `json:"FlowLogName,omitnil" name:"FlowLogName"`

	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE|CCN|NAT|DCG。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源唯一ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流日志采集类型，ACCEPT|REJECT|ALL。
	TrafficType *string `json:"TrafficType,omitnil" name:"TrafficType"`

	// 私用网络ID或者统一ID，建议使用统一ID，当ResourceType为CCN时不填，其他类型必填。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志实例描述。
	FlowLogDescription *string `json:"FlowLogDescription,omitnil" name:"FlowLogDescription"`

	// 流日志存储ID。
	CloudLogId *string `json:"CloudLogId,omitnil" name:"CloudLogId"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 消费端类型：cls、ckafka。默认值cls。
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// 流日志消费端信息，当消费端类型为ckafka时，必填。
	FlowLogStorage *FlowLogStorage `json:"FlowLogStorage,omitnil" name:"FlowLogStorage"`

	// 流日志存储ID对应的地域，不传递默认为本地域。
	CloudLogRegion *string `json:"CloudLogRegion,omitnil" name:"CloudLogRegion"`
}

func (r *CreateFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogName")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "TrafficType")
	delete(f, "VpcId")
	delete(f, "FlowLogDescription")
	delete(f, "CloudLogId")
	delete(f, "Tags")
	delete(f, "StorageType")
	delete(f, "FlowLogStorage")
	delete(f, "CloudLogRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlowLogResponseParams struct {
	// 创建的流日志信息。
	FlowLog []*FlowLog `json:"FlowLog,omitnil" name:"FlowLog"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlowLogResponseParams `json:"Response"`
}

func (r *CreateFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHaVipRequestParams struct {
	// `HAVIP`所在私有网络`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `HAVIP`所在子网`ID`。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// `HAVIP`名称。
	HaVipName *string `json:"HaVipName,omitnil" name:"HaVipName"`

	// 指定虚拟IP地址，必须在`VPC`网段内且未被占用。不指定则自动分配。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// `HAVIP`所在弹性网卡`ID`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`
}

type CreateHaVipRequest struct {
	*tchttp.BaseRequest
	
	// `HAVIP`所在私有网络`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `HAVIP`所在子网`ID`。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// `HAVIP`名称。
	HaVipName *string `json:"HaVipName,omitnil" name:"HaVipName"`

	// 指定虚拟IP地址，必须在`VPC`网段内且未被占用。不指定则自动分配。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// `HAVIP`所在弹性网卡`ID`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`
}

func (r *CreateHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "HaVipName")
	delete(f, "Vip")
	delete(f, "NetworkInterfaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHaVipResponseParams struct {
	// `HAVIP`对象。
	HaVip *HaVip `json:"HaVip,omitnil" name:"HaVip"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateHaVipResponse struct {
	*tchttp.BaseResponse
	Response *CreateHaVipResponseParams `json:"Response"`
}

func (r *CreateHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIp6TranslatorsRequestParams struct {
	// 转换实例名称
	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitnil" name:"Ip6TranslatorName"`

	// 创建转换实例数量，默认是1个
	Ip6TranslatorCount *int64 `json:"Ip6TranslatorCount,omitnil" name:"Ip6TranslatorCount"`

	// 转换实例运营商属性，可取"CMCC","CTCC","CUCC","BGP"
	Ip6InternetServiceProvider *string `json:"Ip6InternetServiceProvider,omitnil" name:"Ip6InternetServiceProvider"`
}

type CreateIp6TranslatorsRequest struct {
	*tchttp.BaseRequest
	
	// 转换实例名称
	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitnil" name:"Ip6TranslatorName"`

	// 创建转换实例数量，默认是1个
	Ip6TranslatorCount *int64 `json:"Ip6TranslatorCount,omitnil" name:"Ip6TranslatorCount"`

	// 转换实例运营商属性，可取"CMCC","CTCC","CUCC","BGP"
	Ip6InternetServiceProvider *string `json:"Ip6InternetServiceProvider,omitnil" name:"Ip6InternetServiceProvider"`
}

func (r *CreateIp6TranslatorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIp6TranslatorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorName")
	delete(f, "Ip6TranslatorCount")
	delete(f, "Ip6InternetServiceProvider")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIp6TranslatorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIp6TranslatorsResponseParams struct {
	// 转换实例的唯一ID数组，形如"ip6-xxxxxxxx"
	Ip6TranslatorSet []*string `json:"Ip6TranslatorSet,omitnil" name:"Ip6TranslatorSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateIp6TranslatorsResponse struct {
	*tchttp.BaseResponse
	Response *CreateIp6TranslatorsResponseParams `json:"Response"`
}

func (r *CreateIp6TranslatorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIp6TranslatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLocalGatewayRequestParams struct {
	// 本地网关名称。
	LocalGatewayName *string `json:"LocalGatewayName,omitnil" name:"LocalGatewayName"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`
}

type CreateLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 本地网关名称。
	LocalGatewayName *string `json:"LocalGatewayName,omitnil" name:"LocalGatewayName"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`
}

func (r *CreateLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalGatewayName")
	delete(f, "VpcId")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLocalGatewayResponseParams struct {
	// 本地网关信息。
	LocalGateway *LocalGateway `json:"LocalGateway,omitnil" name:"LocalGateway"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateLocalGatewayResponseParams `json:"Response"`
}

func (r *CreateLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayDestinationIpPortTranslationNatRuleRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的端口转换规则。
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil" name:"DestinationIpPortTranslationNatRules"`
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的端口转换规则。
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil" name:"DestinationIpPortTranslationNatRules"`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "DestinationIpPortTranslationNatRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatGatewayDestinationIpPortTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayDestinationIpPortTranslationNatRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatGatewayDestinationIpPortTranslationNatRuleResponseParams `json:"Response"`
}

func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayRequestParams struct {
	// NAT网关名称
	NatGatewayName *string `json:"NatGatewayName,omitnil" name:"NatGatewayName"`

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关最大外网出带宽(单位:Mbps)，支持的参数值：`20, 50, 100, 200, 500, 1000, 2000, 5000`，默认: `100Mbps`。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// NAT网关并发连接上限，支持参数值：`1000000、3000000、10000000`，默认值为`100000`。
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil" name:"MaxConcurrentConnection"`

	// 需要申请的弹性IP个数，系统会按您的要求生产N个弹性IP，其中AddressCount和PublicAddresses至少传递一个。
	AddressCount *uint64 `json:"AddressCount,omitnil" name:"AddressCount"`

	// 绑定NAT网关的弹性IP数组，其中AddressCount和PublicAddresses至少传递一个。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`

	// 可用区，形如：`ap-guangzhou-1`。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// NAT网关所属子网
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 绑定NAT网关的弹性IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil" name:"StockPublicIpAddressesBandwidthOut"`

	// 需要申请公网IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil" name:"PublicIpAddressesBandwidthOut"`

	// 公网IP是否强制与NAT网关来自同可用区，true表示需要与NAT网关同可用区；false表示可与NAT网关不是同一个可用区。此参数只有当参数Zone存在时才能生效。
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil" name:"PublicIpFromSameZone"`

	// NAT网关大版本号，1是传统型，2是标准型，默认是1
	NatProductVersion *uint64 `json:"NatProductVersion,omitnil" name:"NatProductVersion"`
}

type CreateNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关名称
	NatGatewayName *string `json:"NatGatewayName,omitnil" name:"NatGatewayName"`

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关最大外网出带宽(单位:Mbps)，支持的参数值：`20, 50, 100, 200, 500, 1000, 2000, 5000`，默认: `100Mbps`。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// NAT网关并发连接上限，支持参数值：`1000000、3000000、10000000`，默认值为`100000`。
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil" name:"MaxConcurrentConnection"`

	// 需要申请的弹性IP个数，系统会按您的要求生产N个弹性IP，其中AddressCount和PublicAddresses至少传递一个。
	AddressCount *uint64 `json:"AddressCount,omitnil" name:"AddressCount"`

	// 绑定NAT网关的弹性IP数组，其中AddressCount和PublicAddresses至少传递一个。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`

	// 可用区，形如：`ap-guangzhou-1`。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// NAT网关所属子网
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 绑定NAT网关的弹性IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	StockPublicIpAddressesBandwidthOut *uint64 `json:"StockPublicIpAddressesBandwidthOut,omitnil" name:"StockPublicIpAddressesBandwidthOut"`

	// 需要申请公网IP带宽大小（单位Mbps），默认为当前用户类型所能使用的最大值。
	PublicIpAddressesBandwidthOut *uint64 `json:"PublicIpAddressesBandwidthOut,omitnil" name:"PublicIpAddressesBandwidthOut"`

	// 公网IP是否强制与NAT网关来自同可用区，true表示需要与NAT网关同可用区；false表示可与NAT网关不是同一个可用区。此参数只有当参数Zone存在时才能生效。
	PublicIpFromSameZone *bool `json:"PublicIpFromSameZone,omitnil" name:"PublicIpFromSameZone"`

	// NAT网关大版本号，1是传统型，2是标准型，默认是1
	NatProductVersion *uint64 `json:"NatProductVersion,omitnil" name:"NatProductVersion"`
}

func (r *CreateNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayName")
	delete(f, "VpcId")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "MaxConcurrentConnection")
	delete(f, "AddressCount")
	delete(f, "PublicIpAddresses")
	delete(f, "Zone")
	delete(f, "Tags")
	delete(f, "SubnetId")
	delete(f, "StockPublicIpAddressesBandwidthOut")
	delete(f, "PublicIpAddressesBandwidthOut")
	delete(f, "PublicIpFromSameZone")
	delete(f, "NatProductVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewayResponseParams struct {
	// NAT网关对象数组。
	NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitnil" name:"NatGatewaySet"`

	// 符合条件的 NAT网关对象数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatGatewayResponseParams `json:"Response"`
}

func (r *CreateNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewaySourceIpTranslationNatRuleRequestParams struct {
	// NAT网关的ID，形如："nat-df45454"
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的SNAT转换规则
	SourceIpTranslationNatRules []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRules,omitnil" name:"SourceIpTranslationNatRules"`
}

type CreateNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如："nat-df45454"
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的SNAT转换规则
	SourceIpTranslationNatRules []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRules,omitnil" name:"SourceIpTranslationNatRules"`
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "SourceIpTranslationNatRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatGatewaySourceIpTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatGatewaySourceIpTranslationNatRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatGatewaySourceIpTranslationNatRuleResponseParams `json:"Response"`
}

func (r *CreateNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetDetectRequestParams struct {
	// `VPC`实例`ID`。形如：`vpc-12345678`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。形如：subnet-12345678。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 网络探测名称，最大长度不能超过60个字节。
	NetDetectName *string `json:"NetDetectName,omitnil" name:"NetDetectName"`

	// 探测目的IPv4地址数组。最多两个。
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云服务器；
	// CCN：云联网网关；
	// NONEXTHOP：无下一跳；
	NextHopType *string `json:"NextHopType,omitnil" name:"NextHopType"`

	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云服务器IPv4地址，形如：10.0.0.12；
	// 下一跳类型为CCN，取值云联网ID，形如：ccn-12345678；
	// 下一跳类型为NONEXTHOP，指定网络探测为无下一跳的网络探测；
	NextHopDestination *string `json:"NextHopDestination,omitnil" name:"NextHopDestination"`

	// 网络探测描述。
	NetDetectDescription *string `json:"NetDetectDescription,omitnil" name:"NetDetectDescription"`
}

type CreateNetDetectRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`。形如：`vpc-12345678`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。形如：subnet-12345678。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 网络探测名称，最大长度不能超过60个字节。
	NetDetectName *string `json:"NetDetectName,omitnil" name:"NetDetectName"`

	// 探测目的IPv4地址数组。最多两个。
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云服务器；
	// CCN：云联网网关；
	// NONEXTHOP：无下一跳；
	NextHopType *string `json:"NextHopType,omitnil" name:"NextHopType"`

	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云服务器IPv4地址，形如：10.0.0.12；
	// 下一跳类型为CCN，取值云联网ID，形如：ccn-12345678；
	// 下一跳类型为NONEXTHOP，指定网络探测为无下一跳的网络探测；
	NextHopDestination *string `json:"NextHopDestination,omitnil" name:"NextHopDestination"`

	// 网络探测描述。
	NetDetectDescription *string `json:"NetDetectDescription,omitnil" name:"NetDetectDescription"`
}

func (r *CreateNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "NetDetectName")
	delete(f, "DetectDestinationIp")
	delete(f, "NextHopType")
	delete(f, "NextHopDestination")
	delete(f, "NetDetectDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetDetectResponseParams struct {
	// 网络探测（NetDetect）对象。
	NetDetect *NetDetect `json:"NetDetect,omitnil" name:"NetDetect"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetDetectResponseParams `json:"Response"`
}

func (r *CreateNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclEntriesRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络三元组ACL规则集。
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil" name:"NetworkAclEntrySet"`
}

type CreateNetworkAclEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络三元组ACL规则集。
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil" name:"NetworkAclEntrySet"`
}

func (r *CreateNetworkAclEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclEntrySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkAclEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclEntriesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNetworkAclEntriesResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkAclEntriesResponseParams `json:"Response"`
}

func (r *CreateNetworkAclEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclQuintupleEntriesRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络五元组ACL规则集。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`
}

type CreateNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络五元组ACL规则集。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`
}

func (r *CreateNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclQuintupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclQuintupleEntriesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *CreateNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclRequestParams struct {
	// VPC实例ID。可通过<a href="https://cloud.tencent.com/document/product/215/15778">DescribeVpcs</a>接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 网络ACL名称，最大长度不能超过60个字节。
	NetworkAclName *string `json:"NetworkAclName,omitnil" name:"NetworkAclName"`

	// 网络ACL类型，三元组(TRIPLE)或五元组(QUINTUPLE)。默认值三元组(TRIPLE)。
	NetworkAclType *string `json:"NetworkAclType,omitnil" name:"NetworkAclType"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateNetworkAclRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过<a href="https://cloud.tencent.com/document/product/215/15778">DescribeVpcs</a>接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 网络ACL名称，最大长度不能超过60个字节。
	NetworkAclName *string `json:"NetworkAclName,omitnil" name:"NetworkAclName"`

	// 网络ACL类型，三元组(TRIPLE)或五元组(QUINTUPLE)。默认值三元组(TRIPLE)。
	NetworkAclType *string `json:"NetworkAclType,omitnil" name:"NetworkAclType"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateNetworkAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NetworkAclName")
	delete(f, "NetworkAclType")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkAclResponseParams struct {
	// 网络ACL实例。
	NetworkAcl *NetworkAcl `json:"NetworkAcl,omitnil" name:"NetworkAcl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNetworkAclResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkAclResponseParams `json:"Response"`
}

func (r *CreateNetworkAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil" name:"NetworkInterfaceName"`

	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil" name:"NetworkInterfaceDescription"`

	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配额数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`

	// IP服务质量等级，和SecondaryPrivateIpAddressCount配合使用，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 网卡trunking模式设置，Enable-开启，Disable--关闭，默认关闭。
	TrunkingFlag *string `json:"TrunkingFlag,omitnil" name:"TrunkingFlag"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type CreateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil" name:"NetworkInterfaceName"`

	// 弹性网卡所在的子网实例ID，例如：subnet-0ap8nwca。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil" name:"NetworkInterfaceDescription"`

	// 新申请的内网IP地址个数，内网IP地址个数总和不能超过配额数。
	SecondaryPrivateIpAddressCount *uint64 `json:"SecondaryPrivateIpAddressCount,omitnil" name:"SecondaryPrivateIpAddressCount"`

	// IP服务质量等级，和SecondaryPrivateIpAddressCount配合使用，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 指定绑定的安全组，例如：['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 网卡trunking模式设置，Enable-开启，Disable--关闭，默认关闭。
	TrunkingFlag *string `json:"TrunkingFlag,omitnil" name:"TrunkingFlag"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *CreateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NetworkInterfaceName")
	delete(f, "SubnetId")
	delete(f, "NetworkInterfaceDescription")
	delete(f, "SecondaryPrivateIpAddressCount")
	delete(f, "QosLevel")
	delete(f, "SecurityGroupIds")
	delete(f, "PrivateIpAddresses")
	delete(f, "Tags")
	delete(f, "TrunkingFlag")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNetworkInterfaceResponseParams struct {
	// 弹性网卡实例。
	NetworkInterface *NetworkInterface `json:"NetworkInterface,omitnil" name:"NetworkInterface"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNetworkInterfaceResponseParams `json:"Response"`
}

func (r *CreateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteTableRequestParams struct {
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "RouteTableName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRouteTableResponseParams struct {
	// 路由表对象。
	RouteTable *RouteTable `json:"RouteTable,omitnil" name:"RouteTable"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateRouteTableResponseParams `json:"Response"`
}

func (r *CreateRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutesRequestParams struct {
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略对象。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

type CreateRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略对象。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

func (r *CreateRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoutesResponseParams struct {
	// 新增的实例个数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 路由表对象。
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitnil" name:"RouteTableSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRoutesResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoutesResponseParams `json:"Response"`
}

func (r *CreateRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`
}

type CreateSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRequestParams struct {
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`

	// 项目ID，默认0。可在<a href="https://console.cloud.tencent.com/project">控制台项目管理页面</a>查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`

	// 项目ID，默认0。可在<a href="https://console.cloud.tencent.com/project">控制台项目管理页面</a>查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	delete(f, "ProjectId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupResponseParams struct {
	// 安全组对象。
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupWithPoliciesRequestParams struct {
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`

	// 项目ID，默认0。可在<a href="https://console.cloud.tencent.com/project">控制台项目管理页面</a>查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`
}

type CreateSecurityGroupWithPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`

	// 项目ID，默认0。可在<a href="https://console.cloud.tencent.com/project">控制台项目管理页面</a>查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`
}

func (r *CreateSecurityGroupWithPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupWithPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	delete(f, "ProjectId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupWithPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupWithPoliciesResponseParams struct {
	// 安全组对象。
	SecurityGroup *SecurityGroup `json:"SecurityGroup,omitnil" name:"SecurityGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSecurityGroupWithPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupWithPoliciesResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupWithPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupWithPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateGroupRequestParams struct {
	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil" name:"ServiceTemplateGroupName"`

	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil" name:"ServiceTemplateIds"`
}

type CreateServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil" name:"ServiceTemplateGroupName"`

	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil" name:"ServiceTemplateIds"`
}

func (r *CreateServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateGroupName")
	delete(f, "ServiceTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateGroupResponseParams struct {
	// 协议端口模板集合对象。
	ServiceTemplateGroup *ServiceTemplateGroup `json:"ServiceTemplateGroup,omitnil" name:"ServiceTemplateGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceTemplateGroupResponseParams `json:"Response"`
}

func (r *CreateServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateRequestParams struct {
	// 协议端口模板名称。
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil" name:"ServiceTemplateName"`

	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。Services与ServicesExtra必填其一。
	Services []*string `json:"Services,omitnil" name:"Services"`

	// 支持添加备注，单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。Services与ServicesExtra必填其一。
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil" name:"ServicesExtra"`
}

type CreateServiceTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 协议端口模板名称。
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil" name:"ServiceTemplateName"`

	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。Services与ServicesExtra必填其一。
	Services []*string `json:"Services,omitnil" name:"Services"`

	// 支持添加备注，单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。Services与ServicesExtra必填其一。
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil" name:"ServicesExtra"`
}

func (r *CreateServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateName")
	delete(f, "Services")
	delete(f, "ServicesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateServiceTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateServiceTemplateResponseParams struct {
	// 协议端口模板对象。
	ServiceTemplate *ServiceTemplate `json:"ServiceTemplate,omitnil" name:"ServiceTemplate"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateServiceTemplateResponseParams `json:"Response"`
}

func (r *CreateServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateServiceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotPoliciesRequestParams struct {
	// 快照策略详情。
	SnapshotPolicies []*SnapshotPolicy `json:"SnapshotPolicies,omitnil" name:"SnapshotPolicies"`
}

type CreateSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略详情。
	SnapshotPolicies []*SnapshotPolicy `json:"SnapshotPolicies,omitnil" name:"SnapshotPolicies"`
}

func (r *CreateSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSnapshotPoliciesResponseParams struct {
	// 快照策略。
	SnapshotPolicies []*SnapshotPolicy `json:"SnapshotPolicies,omitnil" name:"SnapshotPolicies"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *CreateSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetRequestParams struct {
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`
}

type CreateSubnetRequest struct {
	*tchttp.BaseRequest
	
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`
}

func (r *CreateSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetName")
	delete(f, "CidrBlock")
	delete(f, "Zone")
	delete(f, "Tags")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetResponseParams struct {
	// 子网对象。
	Subnet *Subnet `json:"Subnet,omitnil" name:"Subnet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSubnetResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubnetResponseParams `json:"Response"`
}

func (r *CreateSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetsRequestParams struct {
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网对象列表。
	Subnets []*SubnetInput `json:"Subnets,omitnil" name:"Subnets"`

	// 指定绑定的标签列表，注意这里的标签集合为列表中所有子网对象所共享，不能为每个子网对象单独指定标签，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 需要增加到的CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`
}

type CreateSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网对象列表。
	Subnets []*SubnetInput `json:"Subnets,omitnil" name:"Subnets"`

	// 指定绑定的标签列表，注意这里的标签集合为列表中所有子网对象所共享，不能为每个子网对象单独指定标签，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 需要增加到的CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`
}

func (r *CreateSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Subnets")
	delete(f, "Tags")
	delete(f, "CdcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubnetsResponseParams struct {
	// 新创建的子网列表。
	SubnetSet []*Subnet `json:"SubnetSet,omitnil" name:"SubnetSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubnetsResponseParams `json:"Response"`
}

func (r *CreateSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrafficPackagesRequestParams struct {
	// 流量包规格。可选值:
	// <li>10: 10GB流量，有效期一个月</li>
	// <li>50: 50GB流量，有效期一个月</li>
	// <li>512: 512GB流量，有效期一个月</li>
	// <li>1024: 1TB流量，有效期一个月</li>
	// <li>5120: 5TB流量，有效期一个月</li>
	// <li>51200: 50TB流量，有效期一个月</li>
	// <li>60: 60GB流量，有效期半年</li>
	// <li>300: 300GB流量，有效期半年</li>
	// <li>600: 600GB流量，有效期半年</li>
	// <li>3072: 3TB流量，有效期半年</li>
	// <li>6144: 6TB流量，有效期半年</li>
	// <li>30720: 30TB流量，有效期半年</li>
	// <li>61440: 60TB流量，有效期半年</li>
	// <li>307200: 300TB流量，有效期半年</li>
	TrafficAmount *uint64 `json:"TrafficAmount,omitnil" name:"TrafficAmount"`

	// 流量包数量，可选范围 1~20。
	TrafficPackageCount *uint64 `json:"TrafficPackageCount,omitnil" name:"TrafficPackageCount"`
}

type CreateTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 流量包规格。可选值:
	// <li>10: 10GB流量，有效期一个月</li>
	// <li>50: 50GB流量，有效期一个月</li>
	// <li>512: 512GB流量，有效期一个月</li>
	// <li>1024: 1TB流量，有效期一个月</li>
	// <li>5120: 5TB流量，有效期一个月</li>
	// <li>51200: 50TB流量，有效期一个月</li>
	// <li>60: 60GB流量，有效期半年</li>
	// <li>300: 300GB流量，有效期半年</li>
	// <li>600: 600GB流量，有效期半年</li>
	// <li>3072: 3TB流量，有效期半年</li>
	// <li>6144: 6TB流量，有效期半年</li>
	// <li>30720: 30TB流量，有效期半年</li>
	// <li>61440: 60TB流量，有效期半年</li>
	// <li>307200: 300TB流量，有效期半年</li>
	TrafficAmount *uint64 `json:"TrafficAmount,omitnil" name:"TrafficAmount"`

	// 流量包数量，可选范围 1~20。
	TrafficPackageCount *uint64 `json:"TrafficPackageCount,omitnil" name:"TrafficPackageCount"`
}

func (r *CreateTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficAmount")
	delete(f, "TrafficPackageCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrafficPackagesResponseParams struct {
	// 创建的流量包ID列表。
	TrafficPackageSet []*string `json:"TrafficPackageSet,omitnil" name:"TrafficPackageSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrafficPackagesResponseParams `json:"Response"`
}

func (r *CreateTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointRequestParams struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 终端节点名称。
	EndPointName *string `json:"EndPointName,omitnil" name:"EndPointName"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 终端节点VIP，可以指定IP申请。
	EndPointVip *string `json:"EndPointVip,omitnil" name:"EndPointVip"`

	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

type CreateVpcEndPointRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 终端节点名称。
	EndPointName *string `json:"EndPointName,omitnil" name:"EndPointName"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 终端节点VIP，可以指定IP申请。
	EndPointVip *string `json:"EndPointVip,omitnil" name:"EndPointVip"`

	// 安全组ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

func (r *CreateVpcEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "EndPointName")
	delete(f, "EndPointServiceId")
	delete(f, "EndPointVip")
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointResponseParams struct {
	// 终端节点对象详细信息。
	EndPoint *EndPoint `json:"EndPoint,omitnil" name:"EndPoint"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpcEndPointResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcEndPointResponseParams `json:"Response"`
}

func (r *CreateVpcEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceRequestParams struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 终端节点服务名称。
	EndPointServiceName *string `json:"EndPointServiceName,omitnil" name:"EndPointServiceName"`

	// 是否自动接受。
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil" name:"AutoAcceptFlag"`

	// 后端服务ID，比如lb-xxx。
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil" name:"ServiceInstanceId"`

	// ~~是否是PassService类型。该字段已废弃，请不要使用该字段。~~
	IsPassService *bool `json:"IsPassService,omitnil" name:"IsPassService"`

	// 挂载的PAAS服务类型，CLB,CDB,CRS，不填默认挂载为CLB。
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`
}

type CreateVpcEndPointServiceRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 终端节点服务名称。
	EndPointServiceName *string `json:"EndPointServiceName,omitnil" name:"EndPointServiceName"`

	// 是否自动接受。
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil" name:"AutoAcceptFlag"`

	// 后端服务ID，比如lb-xxx。
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil" name:"ServiceInstanceId"`

	// ~~是否是PassService类型。该字段已废弃，请不要使用该字段。~~
	IsPassService *bool `json:"IsPassService,omitnil" name:"IsPassService"`

	// 挂载的PAAS服务类型，CLB,CDB,CRS，不填默认挂载为CLB。
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`
}

func (r *CreateVpcEndPointServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "EndPointServiceName")
	delete(f, "AutoAcceptFlag")
	delete(f, "ServiceInstanceId")
	delete(f, "IsPassService")
	delete(f, "ServiceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcEndPointServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceResponseParams struct {
	// 终端节点服务对象详细信息。
	EndPointService *EndPointService `json:"EndPointService,omitnil" name:"EndPointService"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpcEndPointServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcEndPointServiceResponseParams `json:"Response"`
}

func (r *CreateVpcEndPointServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceWhiteListRequestParams struct {
	// UIN。
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 白名单描述。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// UIN。
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 白名单描述。
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *CreateVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUin")
	delete(f, "EndPointServiceId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcEndPointServiceWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *CreateVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcPeeringConnectionRequestParams struct {
	// 本端VPC唯一ID。
	SourceVpcId *string `json:"SourceVpcId,omitnil" name:"SourceVpcId"`

	// 对等连接名称。
	PeeringConnectionName *string `json:"PeeringConnectionName,omitnil" name:"PeeringConnectionName"`

	// 对端VPC唯一ID。
	DestinationVpcId *string `json:"DestinationVpcId,omitnil" name:"DestinationVpcId"`

	// 对端用户UIN。
	DestinationUin *string `json:"DestinationUin,omitnil" name:"DestinationUin"`

	// 对端地域。
	DestinationRegion *string `json:"DestinationRegion,omitnil" name:"DestinationRegion"`

	// 带宽上限，单位Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 互通类型，VPC_PEER：VPC间互通；VPC_BM_PEER：VPC与黑石网络互通。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 计费模式，日峰值POSTPAID_BY_DAY_MAX，月95POSTPAID_BY_MONTH_95。
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// 服务分级：PT、AU、AG。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`
}

type CreateVpcPeeringConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 本端VPC唯一ID。
	SourceVpcId *string `json:"SourceVpcId,omitnil" name:"SourceVpcId"`

	// 对等连接名称。
	PeeringConnectionName *string `json:"PeeringConnectionName,omitnil" name:"PeeringConnectionName"`

	// 对端VPC唯一ID。
	DestinationVpcId *string `json:"DestinationVpcId,omitnil" name:"DestinationVpcId"`

	// 对端用户UIN。
	DestinationUin *string `json:"DestinationUin,omitnil" name:"DestinationUin"`

	// 对端地域。
	DestinationRegion *string `json:"DestinationRegion,omitnil" name:"DestinationRegion"`

	// 带宽上限，单位Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 互通类型，VPC_PEER：VPC间互通；VPC_BM_PEER：VPC与黑石网络互通。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 计费模式，日峰值POSTPAID_BY_DAY_MAX，月95POSTPAID_BY_MONTH_95。
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// 服务分级：PT、AU、AG。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`
}

func (r *CreateVpcPeeringConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceVpcId")
	delete(f, "PeeringConnectionName")
	delete(f, "DestinationVpcId")
	delete(f, "DestinationUin")
	delete(f, "DestinationRegion")
	delete(f, "Bandwidth")
	delete(f, "Type")
	delete(f, "ChargeType")
	delete(f, "QosLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcPeeringConnectionResponseParams struct {
	// 对等连接ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpcPeeringConnectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcPeeringConnectionResponseParams `json:"Response"`
}

func (r *CreateVpcPeeringConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcRequestParams struct {
	// vpc名称，最大长度不能超过60个字节。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// vpc的cidr，仅能在10.0.0.0/12，172.16.0.0/12，192.168.0.0/16这三个内网网段内。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 是否开启组播。true: 开启, false: 不开启。
	EnableMulticast *string `json:"EnableMulticast,omitnil" name:"EnableMulticast"`

	// DNS地址，最多支持4个。
	DnsServers []*string `json:"DnsServers,omitnil" name:"DnsServers"`

	// DHCP使用的域名。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateVpcRequest struct {
	*tchttp.BaseRequest
	
	// vpc名称，最大长度不能超过60个字节。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// vpc的cidr，仅能在10.0.0.0/12，172.16.0.0/12，192.168.0.0/16这三个内网网段内。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 是否开启组播。true: 开启, false: 不开启。
	EnableMulticast *string `json:"EnableMulticast,omitnil" name:"EnableMulticast"`

	// DNS地址，最多支持4个。
	DnsServers []*string `json:"DnsServers,omitnil" name:"DnsServers"`

	// DHCP使用的域名。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcName")
	delete(f, "CidrBlock")
	delete(f, "EnableMulticast")
	delete(f, "DnsServers")
	delete(f, "DomainName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcResponseParams struct {
	// Vpc对象。
	Vpc *Vpc `json:"Vpc,omitnil" name:"Vpc"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpcResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcResponseParams `json:"Response"`
}

func (r *CreateVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnConnectionRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 对端网关ID。例如：cgw-2wqq41m9，可通过[DescribeCustomerGateways](https://cloud.tencent.com/document/product/215/17516)接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`

	// 通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName,omitnil" name:"VpnConnectionName"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitnil" name:"PreShareKey"`

	// VPC实例ID。可通过[DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)接口返回值中的VpcId获取。
	// CCN VPN 形的通道 可以不传VPCID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil" name:"SecurityPolicyDatabases"`

	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil" name:"IKEOptionsSpecification"`

	// IPSec配置，腾讯云提供IPSec安全会话设置
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil" name:"IPSECOptionsSpecification"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 是否支持隧道内健康检查，默认为False。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil" name:"EnableHealthCheck"`

	// 健康检查本端地址，默认值为随机在169.254.128.0/17分配一个IP。
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil" name:"HealthCheckLocalIp"`

	// 健康检查对端地址，默认值为随机在169.254.128.0/17分配一个IP。
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil" name:"HealthCheckRemoteIp"`

	// 通道类型, 例如:["STATIC", "StaticRoute", "Policy"]
	RouteType *string `json:"RouteType,omitnil" name:"RouteType"`

	// 协商类型，默认为active（主动协商）。可选值：active（主动协商），passive（被动协商），flowTrigger（流量协商）
	NegotiationType *string `json:"NegotiationType,omitnil" name:"NegotiationType"`

	// DPD探测开关。默认为0，表示关闭DPD探测。可选值：0（关闭），1（开启）
	DpdEnable *int64 `json:"DpdEnable,omitnil" name:"DpdEnable"`

	// DPD超时时间。即探测确认对端不存在需要的时间。dpdEnable为1（开启）时有效。默认30，单位为秒
	DpdTimeout *string `json:"DpdTimeout,omitnil" name:"DpdTimeout"`

	// DPD超时后的动作。默认为clear。dpdEnable为1（开启）时有效。可取值为clear（断开）和restart（重试）
	DpdAction *string `json:"DpdAction,omitnil" name:"DpdAction"`
}

type CreateVpnConnectionRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 对端网关ID。例如：cgw-2wqq41m9，可通过[DescribeCustomerGateways](https://cloud.tencent.com/document/product/215/17516)接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`

	// 通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName,omitnil" name:"VpnConnectionName"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitnil" name:"PreShareKey"`

	// VPC实例ID。可通过[DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)接口返回值中的VpcId获取。
	// CCN VPN 形的通道 可以不传VPCID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil" name:"SecurityPolicyDatabases"`

	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil" name:"IKEOptionsSpecification"`

	// IPSec配置，腾讯云提供IPSec安全会话设置
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil" name:"IPSECOptionsSpecification"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 是否支持隧道内健康检查，默认为False。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil" name:"EnableHealthCheck"`

	// 健康检查本端地址，默认值为随机在169.254.128.0/17分配一个IP。
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil" name:"HealthCheckLocalIp"`

	// 健康检查对端地址，默认值为随机在169.254.128.0/17分配一个IP。
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil" name:"HealthCheckRemoteIp"`

	// 通道类型, 例如:["STATIC", "StaticRoute", "Policy"]
	RouteType *string `json:"RouteType,omitnil" name:"RouteType"`

	// 协商类型，默认为active（主动协商）。可选值：active（主动协商），passive（被动协商），flowTrigger（流量协商）
	NegotiationType *string `json:"NegotiationType,omitnil" name:"NegotiationType"`

	// DPD探测开关。默认为0，表示关闭DPD探测。可选值：0（关闭），1（开启）
	DpdEnable *int64 `json:"DpdEnable,omitnil" name:"DpdEnable"`

	// DPD超时时间。即探测确认对端不存在需要的时间。dpdEnable为1（开启）时有效。默认30，单位为秒
	DpdTimeout *string `json:"DpdTimeout,omitnil" name:"DpdTimeout"`

	// DPD超时后的动作。默认为clear。dpdEnable为1（开启）时有效。可取值为clear（断开）和restart（重试）
	DpdAction *string `json:"DpdAction,omitnil" name:"DpdAction"`
}

func (r *CreateVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "CustomerGatewayId")
	delete(f, "VpnConnectionName")
	delete(f, "PreShareKey")
	delete(f, "VpcId")
	delete(f, "SecurityPolicyDatabases")
	delete(f, "IKEOptionsSpecification")
	delete(f, "IPSECOptionsSpecification")
	delete(f, "Tags")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheckLocalIp")
	delete(f, "HealthCheckRemoteIp")
	delete(f, "RouteType")
	delete(f, "NegotiationType")
	delete(f, "DpdEnable")
	delete(f, "DpdTimeout")
	delete(f, "DpdAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnConnectionResponseParams struct {
	// 通道实例对象。
	VpnConnection *VpnConnection `json:"VpnConnection,omitnil" name:"VpnConnection"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnConnectionResponseParams `json:"Response"`
}

func (r *CreateVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayRequestParams struct {
	// VPC实例ID。可通过[DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName,omitnil" name:"VpnGatewayName"`

	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100, 200, 500, 1000, 3000；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 可用区，如：ap-guangzhou-2。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// VPN网关类型，默认为IPSEC。值“IPSEC”为VPC型IPSEC VPN网关，值“SSL”为VPC型SSL VPN网关，值“CCN”为云联网型IPSEC VPN网关，值“SSL_CCN”为云联网型SSL VPN网关。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// SSL VPN连接数设置，可选规格：5, 10, 20, 50, 100, 200, 500, 1000；单位：个。仅 SSL / SSL_CCN 类型需要选这个参数。
	MaxConnection *uint64 `json:"MaxConnection,omitnil" name:"MaxConnection"`
}

type CreateVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过[DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName,omitnil" name:"VpnGatewayName"`

	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100, 200, 500, 1000, 3000；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// 可用区，如：ap-guangzhou-2。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// VPN网关类型，默认为IPSEC。值“IPSEC”为VPC型IPSEC VPN网关，值“SSL”为VPC型SSL VPN网关，值“CCN”为云联网型IPSEC VPN网关，值“SSL_CCN”为云联网型SSL VPN网关。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 指定绑定的标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// SSL VPN连接数设置，可选规格：5, 10, 20, 50, 100, 200, 500, 1000；单位：个。仅 SSL / SSL_CCN 类型需要选这个参数。
	MaxConnection *uint64 `json:"MaxConnection,omitnil" name:"MaxConnection"`
}

func (r *CreateVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "VpnGatewayName")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "Zone")
	delete(f, "Type")
	delete(f, "Tags")
	delete(f, "CdcId")
	delete(f, "MaxConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayResponseParams struct {
	// VPN网关对象
	VpnGateway *VpnGateway `json:"VpnGateway,omitnil" name:"VpnGateway"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnGatewayResponseParams `json:"Response"`
}

func (r *CreateVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayRoutesRequestParams struct {
	// VPN网关的ID
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN网关目的路由列表
	Routes []*VpnGatewayRoute `json:"Routes,omitnil" name:"Routes"`
}

type CreateVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关的ID
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN网关目的路由列表
	Routes []*VpnGatewayRoute `json:"Routes,omitnil" name:"Routes"`
}

func (r *CreateVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewayRoutesResponseParams struct {
	// VPN网关目的路由
	Routes []*VpnGatewayRoute `json:"Routes,omitnil" name:"Routes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *CreateVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewaySslClientRequestParams struct {
	// SSL-VPN-SERVER 实例ID。
	SslVpnServerId *string `json:"SslVpnServerId,omitnil" name:"SslVpnServerId"`

	// SSL-VPN-CLIENT实例Name。不可和SslVpnClientNames同时使用。
	SslVpnClientName *string `json:"SslVpnClientName,omitnil" name:"SslVpnClientName"`

	// SSL-VPN-CLIENT实例Name数字。批量创建时使用。不可和SslVpnClientName同时使用。
	SslVpnClientNames []*string `json:"SslVpnClientNames,omitnil" name:"SslVpnClientNames"`
}

type CreateVpnGatewaySslClientRequest struct {
	*tchttp.BaseRequest
	
	// SSL-VPN-SERVER 实例ID。
	SslVpnServerId *string `json:"SslVpnServerId,omitnil" name:"SslVpnServerId"`

	// SSL-VPN-CLIENT实例Name。不可和SslVpnClientNames同时使用。
	SslVpnClientName *string `json:"SslVpnClientName,omitnil" name:"SslVpnClientName"`

	// SSL-VPN-CLIENT实例Name数字。批量创建时使用。不可和SslVpnClientName同时使用。
	SslVpnClientNames []*string `json:"SslVpnClientNames,omitnil" name:"SslVpnClientNames"`
}

func (r *CreateVpnGatewaySslClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewaySslClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SslVpnServerId")
	delete(f, "SslVpnClientName")
	delete(f, "SslVpnClientNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnGatewaySslClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewaySslClientResponseParams struct {
	// 异步任务ID。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// SSL-VPN client 唯一ID
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpnGatewaySslClientResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnGatewaySslClientResponseParams `json:"Response"`
}

func (r *CreateVpnGatewaySslClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewaySslClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewaySslServerRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// SSL-VPN-SERVER 实例名称，长度不超过60个字节。
	SslVpnServerName *string `json:"SslVpnServerName,omitnil" name:"SslVpnServerName"`

	// 云端地址（CIDR）列表。
	LocalAddress []*string `json:"LocalAddress,omitnil" name:"LocalAddress"`

	// 客户端地址网段。
	RemoteAddress *string `json:"RemoteAddress,omitnil" name:"RemoteAddress"`

	// SSL VPN服务端监听协议。当前仅支持 UDP，默认UDP。
	SslVpnProtocol *string `json:"SslVpnProtocol,omitnil" name:"SslVpnProtocol"`

	// SSL VPN服务端监听协议端口，默认1194。
	SslVpnPort *int64 `json:"SslVpnPort,omitnil" name:"SslVpnPort"`

	// 认证算法。可选 'SHA1', 'MD5', 'NONE'，默认NONE。
	IntegrityAlgorithm *string `json:"IntegrityAlgorithm,omitnil" name:"IntegrityAlgorithm"`

	// 加密算法。可选 'AES-128-CBC','AES-192-CBC', 'AES-256-CBC', 'NONE'，默认NONE。
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil" name:"EncryptAlgorithm"`

	// 是否支持压缩。当前仅支持不支持压缩，默认False。
	Compress *bool `json:"Compress,omitnil" name:"Compress"`

	// 是否开启SSO认证。默认为False
	SsoEnabled *bool `json:"SsoEnabled,omitnil" name:"SsoEnabled"`

	// 是否开启策略访问控制。默认为False
	AccessPolicyEnabled *bool `json:"AccessPolicyEnabled,omitnil" name:"AccessPolicyEnabled"`

	// SAML-DATA，开启SSO时传。
	SamlData *string `json:"SamlData,omitnil" name:"SamlData"`
}

type CreateVpnGatewaySslServerRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// SSL-VPN-SERVER 实例名称，长度不超过60个字节。
	SslVpnServerName *string `json:"SslVpnServerName,omitnil" name:"SslVpnServerName"`

	// 云端地址（CIDR）列表。
	LocalAddress []*string `json:"LocalAddress,omitnil" name:"LocalAddress"`

	// 客户端地址网段。
	RemoteAddress *string `json:"RemoteAddress,omitnil" name:"RemoteAddress"`

	// SSL VPN服务端监听协议。当前仅支持 UDP，默认UDP。
	SslVpnProtocol *string `json:"SslVpnProtocol,omitnil" name:"SslVpnProtocol"`

	// SSL VPN服务端监听协议端口，默认1194。
	SslVpnPort *int64 `json:"SslVpnPort,omitnil" name:"SslVpnPort"`

	// 认证算法。可选 'SHA1', 'MD5', 'NONE'，默认NONE。
	IntegrityAlgorithm *string `json:"IntegrityAlgorithm,omitnil" name:"IntegrityAlgorithm"`

	// 加密算法。可选 'AES-128-CBC','AES-192-CBC', 'AES-256-CBC', 'NONE'，默认NONE。
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil" name:"EncryptAlgorithm"`

	// 是否支持压缩。当前仅支持不支持压缩，默认False。
	Compress *bool `json:"Compress,omitnil" name:"Compress"`

	// 是否开启SSO认证。默认为False
	SsoEnabled *bool `json:"SsoEnabled,omitnil" name:"SsoEnabled"`

	// 是否开启策略访问控制。默认为False
	AccessPolicyEnabled *bool `json:"AccessPolicyEnabled,omitnil" name:"AccessPolicyEnabled"`

	// SAML-DATA，开启SSO时传。
	SamlData *string `json:"SamlData,omitnil" name:"SamlData"`
}

func (r *CreateVpnGatewaySslServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewaySslServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "SslVpnServerName")
	delete(f, "LocalAddress")
	delete(f, "RemoteAddress")
	delete(f, "SslVpnProtocol")
	delete(f, "SslVpnPort")
	delete(f, "IntegrityAlgorithm")
	delete(f, "EncryptAlgorithm")
	delete(f, "Compress")
	delete(f, "SsoEnabled")
	delete(f, "AccessPolicyEnabled")
	delete(f, "SamlData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpnGatewaySslServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpnGatewaySslServerResponseParams struct {
	// 创建SSL-VPN server 异步任务ID。
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// SSL-VPN-SERVER 唯一ID。
	SslVpnServerId *string `json:"SslVpnServerId,omitnil" name:"SslVpnServerId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateVpnGatewaySslServerResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpnGatewaySslServerResponseParams `json:"Response"`
}

func (r *CreateVpnGatewaySslServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpnGatewaySslServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrossBorderCompliance struct {
	// 服务商，可选值：`UNICOM`。
	ServiceProvider *string `json:"ServiceProvider,omitnil" name:"ServiceProvider"`

	// 合规化审批单`ID`。
	ComplianceId *uint64 `json:"ComplianceId,omitnil" name:"ComplianceId"`

	// 公司全称。
	Company *string `json:"Company,omitnil" name:"Company"`

	// 统一社会信用代码。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// 法定代表人。
	LegalPerson *string `json:"LegalPerson,omitnil" name:"LegalPerson"`

	// 发证机关。
	IssuingAuthority *string `json:"IssuingAuthority,omitnil" name:"IssuingAuthority"`

	// 营业执照。
	BusinessLicense *string `json:"BusinessLicense,omitnil" name:"BusinessLicense"`

	// 营业执照住所。
	BusinessAddress *string `json:"BusinessAddress,omitnil" name:"BusinessAddress"`

	// 邮编。
	PostCode *uint64 `json:"PostCode,omitnil" name:"PostCode"`

	// 经办人。
	Manager *string `json:"Manager,omitnil" name:"Manager"`

	// 经办人身份证号。
	ManagerId *string `json:"ManagerId,omitnil" name:"ManagerId"`

	// 经办人身份证。
	ManagerIdCard *string `json:"ManagerIdCard,omitnil" name:"ManagerIdCard"`

	// 经办人身份证地址。
	ManagerAddress *string `json:"ManagerAddress,omitnil" name:"ManagerAddress"`

	// 经办人联系电话。
	ManagerTelephone *string `json:"ManagerTelephone,omitnil" name:"ManagerTelephone"`

	// 电子邮箱。
	Email *string `json:"Email,omitnil" name:"Email"`

	// 服务受理单。
	ServiceHandlingForm *string `json:"ServiceHandlingForm,omitnil" name:"ServiceHandlingForm"`

	// 授权函。
	AuthorizationLetter *string `json:"AuthorizationLetter,omitnil" name:"AuthorizationLetter"`

	// 信息安全承诺书。
	SafetyCommitment *string `json:"SafetyCommitment,omitnil" name:"SafetyCommitment"`

	// 服务开始时间。
	ServiceStartDate *string `json:"ServiceStartDate,omitnil" name:"ServiceStartDate"`

	// 服务截止时间。
	ServiceEndDate *string `json:"ServiceEndDate,omitnil" name:"ServiceEndDate"`

	// 状态。待审批：`PENDING`，已通过：`APPROVED`，已拒绝：`DENY`。
	State *string `json:"State,omitnil" name:"State"`

	// 审批单创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type CrossBorderFlowMonitorData struct {
	// 入带宽，单位：`bps`。
	InBandwidth []*int64 `json:"InBandwidth,omitnil" name:"InBandwidth"`

	// 出带宽，单位：`bps`。
	OutBandwidth []*int64 `json:"OutBandwidth,omitnil" name:"OutBandwidth"`

	// 入包，单位：`pps`。
	InPkg []*int64 `json:"InPkg,omitnil" name:"InPkg"`

	// 出包，单位：`pps`。
	OutPkg []*int64 `json:"OutPkg,omitnil" name:"OutPkg"`
}

type CustomerGateway struct {
	// 用户网关唯一ID
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`

	// 网关名称
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil" name:"CustomerGatewayName"`

	// 公网地址
	IpAddress *string `json:"IpAddress,omitnil" name:"IpAddress"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type CustomerGatewayVendor struct {
	// 平台。
	Platform *string `json:"Platform,omitnil" name:"Platform"`

	// 软件版本。
	SoftwareVersion *string `json:"SoftwareVersion,omitnil" name:"SoftwareVersion"`

	// 供应商名称。
	VendorName *string `json:"VendorName,omitnil" name:"VendorName"`
}

type CvmInstance struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 云主机实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 云主机名称。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 云主机状态。
	InstanceState *string `json:"InstanceState,omitnil" name:"InstanceState"`

	// 实例的CPU核数，单位：核。
	CPU *uint64 `json:"CPU,omitnil" name:"CPU"`

	// 实例内存容量，单位：GB。
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例弹性网卡配额（包含主网卡）。
	EniLimit *uint64 `json:"EniLimit,omitnil" name:"EniLimit"`

	// 实例弹性网卡内网IP配额（包含主网卡）。
	EniIpLimit *uint64 `json:"EniIpLimit,omitnil" name:"EniIpLimit"`

	// 实例已绑定弹性网卡的个数（包含主网卡）。
	InstanceEniCount *uint64 `json:"InstanceEniCount,omitnil" name:"InstanceEniCount"`
}

type DefaultVpcSubnet struct {
	// 默认VpcId。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 默认SubnetId。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 默认Vpc名字。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 默认Subnet名字。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 默认子网网段。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`
}

// Predefined struct for user
type DeleteAddressTemplateGroupRequestParams struct {
	// IP地址模板集合实例ID，例如：ipmg-90cex8mq。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil" name:"AddressTemplateGroupId"`
}

type DeleteAddressTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// IP地址模板集合实例ID，例如：ipmg-90cex8mq。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil" name:"AddressTemplateGroupId"`
}

func (r *DeleteAddressTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddressTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAddressTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddressTemplateGroupResponseParams `json:"Response"`
}

func (r *DeleteAddressTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateRequestParams struct {
	// IP地址模板实例ID，例如：ipm-09o5m8kc。
	AddressTemplateId *string `json:"AddressTemplateId,omitnil" name:"AddressTemplateId"`
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// IP地址模板实例ID，例如：ipm-09o5m8kc。
	AddressTemplateId *string `json:"AddressTemplateId,omitnil" name:"AddressTemplateId"`
}

func (r *DeleteAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddressTemplateResponseParams `json:"Response"`
}

func (r *DeleteAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssistantCidrRequestParams struct {
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]。
	CidrBlocks []*string `json:"CidrBlocks,omitnil" name:"CidrBlocks"`
}

type DeleteAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]。
	CidrBlocks []*string `json:"CidrBlocks,omitnil" name:"CidrBlocks"`
}

func (r *DeleteAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "CidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssistantCidrResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAssistantCidrResponseParams `json:"Response"`
}

func (r *DeleteAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBandwidthPackageRequestParams struct {
	// 待删除带宽包唯一ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`
}

type DeleteBandwidthPackageRequest struct {
	*tchttp.BaseRequest
	
	// 待删除带宽包唯一ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`
}

func (r *DeleteBandwidthPackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBandwidthPackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBandwidthPackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBandwidthPackageResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteBandwidthPackageResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBandwidthPackageResponseParams `json:"Response"`
}

func (r *DeleteBandwidthPackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBandwidthPackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcnRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
}

type DeleteCcnRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
}

func (r *DeleteCcnRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcnRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCcnRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCcnResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCcnResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCcnResponseParams `json:"Response"`
}

func (r *DeleteCcnResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomerGatewayRequestParams struct {
	// 对端网关ID，例如：cgw-2wqq41m9，可通过[DescribeCustomerGateways](https://cloud.tencent.com/document/api/215/17516)接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`
}

type DeleteCustomerGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 对端网关ID，例如：cgw-2wqq41m9，可通过[DescribeCustomerGateways](https://cloud.tencent.com/document/api/215/17516)接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`
}

func (r *DeleteCustomerGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomerGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomerGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomerGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCustomerGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomerGatewayResponseParams `json:"Response"`
}

func (r *DeleteCustomerGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomerGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDhcpIpRequestParams struct {
	// `DhcpIp`的`ID`，是`DhcpIp`的唯一标识。
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`
}

type DeleteDhcpIpRequest struct {
	*tchttp.BaseRequest
	
	// `DhcpIp`的`ID`，是`DhcpIp`的唯一标识。
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`
}

func (r *DeleteDhcpIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDhcpIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DhcpIpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDhcpIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDhcpIpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDhcpIpResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDhcpIpResponseParams `json:"Response"`
}

func (r *DeleteDhcpIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDhcpIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayCcnRoutesRequestParams struct {
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 路由ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

type DeleteDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 路由ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

func (r *DeleteDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayCcnRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayRequestParams struct {
	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`
}

type DeleteDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`
}

func (r *DeleteDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDirectConnectGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDirectConnectGatewayResponseParams `json:"Response"`
}

func (r *DeleteDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFlowLogRequestParams struct {
	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`

	// 私用网络ID或者统一ID，建议使用统一ID，删除云联网流日志时，可不填，其他流日志类型必填。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type DeleteFlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`

	// 私用网络ID或者统一ID，建议使用统一ID，删除云联网流日志时，可不填，其他流日志类型必填。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

func (r *DeleteFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFlowLogResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFlowLogResponseParams `json:"Response"`
}

func (r *DeleteFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHaVipRequestParams struct {
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`
}

type DeleteHaVipRequest struct {
	*tchttp.BaseRequest
	
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`
}

func (r *DeleteHaVipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHaVipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHaVipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHaVipResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteHaVipResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHaVipResponseParams `json:"Response"`
}

func (r *DeleteHaVipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHaVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIp6TranslatorsRequestParams struct {
	// 待释放的IPV6转换实例的唯一ID，形如‘ip6-xxxxxxxx’
	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitnil" name:"Ip6TranslatorIds"`
}

type DeleteIp6TranslatorsRequest struct {
	*tchttp.BaseRequest
	
	// 待释放的IPV6转换实例的唯一ID，形如‘ip6-xxxxxxxx’
	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitnil" name:"Ip6TranslatorIds"`
}

func (r *DeleteIp6TranslatorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIp6TranslatorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIp6TranslatorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIp6TranslatorsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteIp6TranslatorsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIp6TranslatorsResponseParams `json:"Response"`
}

func (r *DeleteIp6TranslatorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIp6TranslatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLocalGatewayRequestParams struct {
	// 本地网关实例ID。
	LocalGatewayId *string `json:"LocalGatewayId,omitnil" name:"LocalGatewayId"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type DeleteLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 本地网关实例ID。
	LocalGatewayId *string `json:"LocalGatewayId,omitnil" name:"LocalGatewayId"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

func (r *DeleteLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalGatewayId")
	delete(f, "CdcId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLocalGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLocalGatewayResponseParams `json:"Response"`
}

func (r *DeleteLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayDestinationIpPortTranslationNatRuleRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的端口转换规则。
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil" name:"DestinationIpPortTranslationNatRules"`
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的端口转换规则。
	DestinationIpPortTranslationNatRules []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRules,omitnil" name:"DestinationIpPortTranslationNatRules"`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "DestinationIpPortTranslationNatRules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatGatewayDestinationIpPortTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayDestinationIpPortTranslationNatRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponseParams `json:"Response"`
}

func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`
}

type DeleteNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`
}

func (r *DeleteNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatGatewayResponseParams `json:"Response"`
}

func (r *DeleteNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewaySourceIpTranslationNatRuleRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的SNAT ID列表，形如：`snat-df43254`。
	NatGatewaySnatIds []*string `json:"NatGatewaySnatIds,omitnil" name:"NatGatewaySnatIds"`
}

type DeleteNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的SNAT ID列表，形如：`snat-df43254`。
	NatGatewaySnatIds []*string `json:"NatGatewaySnatIds,omitnil" name:"NatGatewaySnatIds"`
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "NatGatewaySnatIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatGatewaySourceIpTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatGatewaySourceIpTranslationNatRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatGatewaySourceIpTranslationNatRuleResponseParams `json:"Response"`
}

func (r *DeleteNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetDetectRequestParams struct {
	// 网络探测实例`ID`。形如：`netd-12345678`。
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`
}

type DeleteNetDetectRequest struct {
	*tchttp.BaseRequest
	
	// 网络探测实例`ID`。形如：`netd-12345678`。
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`
}

func (r *DeleteNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetDetectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetDetectResponseParams `json:"Response"`
}

func (r *DeleteNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclEntriesRequestParams struct {
	// 三元组网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 三元组网络ACL规则集。
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil" name:"NetworkAclEntrySet"`
}

type DeleteNetworkAclEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 三元组网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 三元组网络ACL规则集。
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil" name:"NetworkAclEntrySet"`
}

func (r *DeleteNetworkAclEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclEntrySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkAclEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclEntriesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNetworkAclEntriesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkAclEntriesResponseParams `json:"Response"`
}

func (r *DeleteNetworkAclEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclQuintupleEntriesRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络五元组ACL规则集。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`
}

type DeleteNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络五元组ACL规则集。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`
}

func (r *DeleteNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclQuintupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclQuintupleEntriesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *DeleteNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`
}

type DeleteNetworkAclRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`
}

func (r *DeleteNetworkAclRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkAclRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkAclResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNetworkAclResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkAclResponseParams `json:"Response"`
}

func (r *DeleteNetworkAclResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`
}

type DeleteNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`
}

func (r *DeleteNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNetworkInterfaceResponseParams `json:"Response"`
}

func (r *DeleteNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTableRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`
}

type DeleteRouteTableRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`
}

func (r *DeleteRouteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRouteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRouteTableResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRouteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRouteTableResponseParams `json:"Response"`
}

func (r *DeleteRouteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutesRequestParams struct {
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略对象，删除路由策略时，仅需使用Route的RouteId字段。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

type DeleteRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略对象，删除路由策略时，仅需使用Route的RouteId字段。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

func (r *DeleteRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoutesResponseParams struct {
	// 已删除的路由策略详情。
	RouteSet []*Route `json:"RouteSet,omitnil" name:"RouteSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoutesResponseParams `json:"Response"`
}

func (r *DeleteRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合。一个请求中只能删除单个方向的一条或多条规则。支持指定索引（PolicyIndex） 匹配删除和安全组规则匹配删除两种方式，一个请求中只能使用一种匹配方式。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`
}

type DeleteSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合。一个请求中只能删除单个方向的一条或多条规则。支持指定索引（PolicyIndex） 匹配删除和安全组规则匹配删除两种方式，一个请求中只能使用一种匹配方式。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`
}

func (r *DeleteSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

type DeleteSecurityGroupRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

func (r *DeleteSecurityGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSecurityGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceTemplateGroupRequestParams struct {
	// 协议端口模板集合实例ID，例如：ppmg-n17uxvve。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil" name:"ServiceTemplateGroupId"`
}

type DeleteServiceTemplateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 协议端口模板集合实例ID，例如：ppmg-n17uxvve。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil" name:"ServiceTemplateGroupId"`
}

func (r *DeleteServiceTemplateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceTemplateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceTemplateGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteServiceTemplateGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceTemplateGroupResponseParams `json:"Response"`
}

func (r *DeleteServiceTemplateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceTemplateRequestParams struct {
	// 协议端口模板实例ID，例如：ppm-e6dy460g。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`
}

type DeleteServiceTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 协议端口模板实例ID，例如：ppm-e6dy460g。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`
}

func (r *DeleteServiceTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteServiceTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteServiceTemplateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteServiceTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteServiceTemplateResponseParams `json:"Response"`
}

func (r *DeleteServiceTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteServiceTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotPoliciesRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`
}

type DeleteSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`
}

func (r *DeleteSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSnapshotPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DeleteSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetRequestParams struct {
	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type DeleteSubnetRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

func (r *DeleteSubnetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubnetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubnetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubnetResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSubnetResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSubnetResponseParams `json:"Response"`
}

func (r *DeleteSubnetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubnetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTemplateMemberRequestParams struct {
	// 参数模板实例ID，支持IP地址、协议端口、IP地址组、协议端口组四种参数模板的实例ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 需要添加的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致。
	TemplateMember []*MemberInfo `json:"TemplateMember,omitnil" name:"TemplateMember"`
}

type DeleteTemplateMemberRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板实例ID，支持IP地址、协议端口、IP地址组、协议端口组四种参数模板的实例ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 需要添加的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致。
	TemplateMember []*MemberInfo `json:"TemplateMember,omitnil" name:"TemplateMember"`
}

func (r *DeleteTemplateMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTemplateMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "TemplateMember")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTemplateMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTemplateMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteTemplateMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTemplateMemberResponseParams `json:"Response"`
}

func (r *DeleteTemplateMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTemplateMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrafficPackagesRequestParams struct {
	// 待删除的流量包唯一ID数组
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil" name:"TrafficPackageIds"`
}

type DeleteTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的流量包唯一ID数组
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil" name:"TrafficPackageIds"`
}

func (r *DeleteTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficPackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTrafficPackagesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTrafficPackagesResponseParams `json:"Response"`
}

func (r *DeleteTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointRequestParams struct {
	// 终端节点ID。
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`
}

type DeleteVpcEndPointRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点ID。
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`
}

func (r *DeleteVpcEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpcEndPointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcEndPointResponseParams `json:"Response"`
}

func (r *DeleteVpcEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceRequestParams struct {
	// 终端节点ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`
}

type DeleteVpcEndPointServiceRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`
}

func (r *DeleteVpcEndPointServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcEndPointServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpcEndPointServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcEndPointServiceResponseParams `json:"Response"`
}

func (r *DeleteVpcEndPointServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceWhiteListRequestParams struct {
	// 用户UIN数组。
	UserUin []*string `json:"UserUin,omitnil" name:"UserUin"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`
}

type DeleteVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 用户UIN数组。
	UserUin []*string `json:"UserUin,omitnil" name:"UserUin"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`
}

func (r *DeleteVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUin")
	delete(f, "EndPointServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcEndPointServiceWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *DeleteVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcPeeringConnectionRequestParams struct {
	// 对等连接唯一ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`
}

type DeleteVpcPeeringConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 对等连接唯一ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`
}

func (r *DeleteVpcPeeringConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PeeringConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcPeeringConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpcPeeringConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcPeeringConnectionResponseParams `json:"Response"`
}

func (r *DeleteVpcPeeringConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type DeleteVpcRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

func (r *DeleteVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpcResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcResponseParams `json:"Response"`
}

func (r *DeleteVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnConnectionRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`
}

type DeleteVpnConnectionRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`
}

func (r *DeleteVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnConnectionResponseParams `json:"Response"`
}

func (r *DeleteVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`
}

type DeleteVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`
}

func (r *DeleteVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnGatewayResponseParams `json:"Response"`
}

func (r *DeleteVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayRoutesRequestParams struct {
	// VPN网关实例ID
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 路由ID信息列表，可以通过[DescribeVpnGatewayRoutes](https://cloud.tencent.com/document/api/215/57676)接口查询。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

type DeleteVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 路由ID信息列表，可以通过[DescribeVpnGatewayRoutes](https://cloud.tencent.com/document/api/215/57676)接口查询。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

func (r *DeleteVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewayRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *DeleteVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewaySslClientRequestParams struct {
	// SSL-VPN-CLIENT 实例ID。不可和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SSL-VPN-CLIENT 实例ID列表。批量删除时使用。不可和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

type DeleteVpnGatewaySslClientRequest struct {
	*tchttp.BaseRequest
	
	// SSL-VPN-CLIENT 实例ID。不可和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SSL-VPN-CLIENT 实例ID列表。批量删除时使用。不可和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

func (r *DeleteVpnGatewaySslClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewaySslClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SslVpnClientId")
	delete(f, "SslVpnClientIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnGatewaySslClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewaySslClientResponseParams struct {
	// 异步任务ID。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpnGatewaySslClientResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnGatewaySslClientResponseParams `json:"Response"`
}

func (r *DeleteVpnGatewaySslClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewaySslClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewaySslServerRequestParams struct {
	// SSL-VPN-SERVER 实例ID。
	SslVpnServerId *string `json:"SslVpnServerId,omitnil" name:"SslVpnServerId"`
}

type DeleteVpnGatewaySslServerRequest struct {
	*tchttp.BaseRequest
	
	// SSL-VPN-SERVER 实例ID。
	SslVpnServerId *string `json:"SslVpnServerId,omitnil" name:"SslVpnServerId"`
}

func (r *DeleteVpnGatewaySslServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewaySslServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SslVpnServerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpnGatewaySslServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpnGatewaySslServerResponseParams struct {
	// 异步任务ID。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpnGatewaySslServerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpnGatewaySslServerResponseParams `json:"Response"`
}

func (r *DeleteVpnGatewaySslServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpnGatewaySslServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountAttributesRequestParams struct {

}

type DescribeAccountAttributesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccountAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountAttributesResponseParams struct {
	// 用户账号属性对象。
	AccountAttributeSet []*AccountAttribute `json:"AccountAttributeSet,omitnil" name:"AccountAttributeSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountAttributesResponseParams `json:"Response"`
}

func (r *DescribeAccountAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressQuotaRequestParams struct {

}

type DescribeAddressQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAddressQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressQuotaResponseParams struct {
	// 账户 EIP 配额信息。
	QuotaSet []*Quota `json:"QuotaSet,omitnil" name:"QuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAddressQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressQuotaResponseParams `json:"Response"`
}

func (r *DescribeAddressQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplateGroupsRequestParams struct {
	// 过滤条件。
	// <li>address-template-group-name - String - （过滤条件）IP地址模板集合名称。</li>
	// <li>address-template-group-id - String - （过滤条件）IP地址模板实集合例ID，例如：ipmg-mdunqeb6。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否查询IP地址模板成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

type DescribeAddressTemplateGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>address-template-group-name - String - （过滤条件）IP地址模板集合名称。</li>
	// <li>address-template-group-id - String - （过滤条件）IP地址模板实集合例ID，例如：ipmg-mdunqeb6。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否查询IP地址模板成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

func (r *DescribeAddressTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplateGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NeedMemberInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressTemplateGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplateGroupsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// IP地址模板。
	AddressTemplateGroupSet []*AddressTemplateGroup `json:"AddressTemplateGroupSet,omitnil" name:"AddressTemplateGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAddressTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressTemplateGroupsResponseParams `json:"Response"`
}

func (r *DescribeAddressTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplateGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplatesRequestParams struct {
	// 过滤条件。
	// <li>address-template-name - IP地址模板名称。</li>
	// <li>address-template-id - IP地址模板实例ID，例如：ipm-mdunqeb6。</li>
	// <li>address-ip - IP地址。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否获取IP地址模板成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

type DescribeAddressTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>address-template-name - IP地址模板名称。</li>
	// <li>address-template-id - IP地址模板实例ID，例如：ipm-mdunqeb6。</li>
	// <li>address-ip - IP地址。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否获取IP地址模板成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

func (r *DescribeAddressTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NeedMemberInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplatesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// IP地址模板。
	AddressTemplateSet []*AddressTemplate `json:"AddressTemplateSet,omitnil" name:"AddressTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAddressTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAddressTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressesRequestParams struct {
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。参数不支持同时指定`AddressIds`和`Filters.address-id`。
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。详细的过滤条件如下：
	// <li> address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。</li>
	// <li> address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。</li>
	// <li> address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。</li>
	// <li> address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。状态包含：'CREATING'，'BINDING'，'BIND'，'UNBINDING'，'UNBIND'，'OFFLINING'，'BIND_ENI'。</li>
	// <li> instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。</li>
	// <li> network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。</li>
	// <li> is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）</li>
	// <li> address-type - String - 是否必填：否 - （过滤条件）按照 IP类型 进行过滤。可选值：'WanIP', 'EIP'，'AnycastEIP'，'HighQualityEIP'。默认值是'EIP'。</li>
	// <li> address-isp - String - 是否必填：否 - （过滤条件）按照 运营商类型 进行过滤。可选值：'BGP'，'CMCC'，'CUCC', 'CTCC'</li>
	// <li> dedicated-cluster-id - String - 是否必填：否 - （过滤条件）按照 CDC 的唯一 ID 过滤。CDC 唯一 ID 形如：cluster-11112222。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 - （过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。tag-key使用具体的标签键进行替换。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAddressesRequest struct {
	*tchttp.BaseRequest
	
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。参数不支持同时指定`AddressIds`和`Filters.address-id`。
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。详细的过滤条件如下：
	// <li> address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。</li>
	// <li> address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。</li>
	// <li> address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。</li>
	// <li> address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。状态包含：'CREATING'，'BINDING'，'BIND'，'UNBINDING'，'UNBIND'，'OFFLINING'，'BIND_ENI'。</li>
	// <li> instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。</li>
	// <li> network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。</li>
	// <li> is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）</li>
	// <li> address-type - String - 是否必填：否 - （过滤条件）按照 IP类型 进行过滤。可选值：'WanIP', 'EIP'，'AnycastEIP'，'HighQualityEIP'。默认值是'EIP'。</li>
	// <li> address-isp - String - 是否必填：否 - （过滤条件）按照 运营商类型 进行过滤。可选值：'BGP'，'CMCC'，'CUCC', 'CTCC'</li>
	// <li> dedicated-cluster-id - String - 是否必填：否 - （过滤条件）按照 CDC 的唯一 ID 过滤。CDC 唯一 ID 形如：cluster-11112222。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 - （过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。tag-key使用具体的标签键进行替换。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API 中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API 中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressesResponseParams struct {
	// 符合条件的 EIP 数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// EIP 详细信息列表。
	AddressSet []*Address `json:"AddressSet,omitnil" name:"AddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressesResponseParams `json:"Response"`
}

func (r *DescribeAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssistantCidrRequestParams struct {
	// `VPC`实例`ID`数组。形如：[`vpc-6v2ht8q5`]
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`数组。形如：[`vpc-6v2ht8q5`]
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`

	// 过滤条件，参数不支持同时指定VpcIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssistantCidrResponseParams struct {
	// 符合条件的辅助CIDR数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil" name:"AssistantCidrSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssistantCidrResponseParams `json:"Response"`
}

func (r *DescribeAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageBillUsageRequestParams struct {
	// 后付费共享带宽包的唯一ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`
}

type DescribeBandwidthPackageBillUsageRequest struct {
	*tchttp.BaseRequest
	
	// 后付费共享带宽包的唯一ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`
}

func (r *DescribeBandwidthPackageBillUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageBillUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackageBillUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageBillUsageResponseParams struct {
	// 当前计费用量
	BandwidthPackageBillBandwidthSet []*BandwidthPackageBillBandwidth `json:"BandwidthPackageBillBandwidthSet,omitnil" name:"BandwidthPackageBillBandwidthSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBandwidthPackageBillUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackageBillUsageResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackageBillUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageBillUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageQuotaRequestParams struct {

}

type DescribeBandwidthPackageQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBandwidthPackageQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackageQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageQuotaResponseParams struct {
	// 带宽包配额详细信息
	QuotaSet []*Quota `json:"QuotaSet,omitnil" name:"QuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBandwidthPackageQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackageQuotaResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackageQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageResourcesRequestParams struct {
	// 标识 共享带宽包 的唯一 ID 列表。共享带宽包 唯一 ID 形如：`bwp-11112222`。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> resource-id - String - 是否必填：否 - （过滤条件）按照 共享带宽包内资源 的唯一 ID 过滤。共享带宽包内资源 唯一 ID 形如：eip-11112222。</li>
	// <li> resource-type - String - 是否必填：否 - （过滤条件）按照 共享带宽包内资源 类型过滤，目前仅支持 弹性IP 和 负载均衡 两种类型，可选值为 Address 和 LoadBalance。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 标识 共享带宽包 的唯一 ID 列表。共享带宽包 唯一 ID 形如：`bwp-11112222`。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> resource-id - String - 是否必填：否 - （过滤条件）按照 共享带宽包内资源 的唯一 ID 过滤。共享带宽包内资源 唯一 ID 形如：eip-11112222。</li>
	// <li> resource-type - String - 是否必填：否 - （过滤条件）按照 共享带宽包内资源 类型过滤，目前仅支持 弹性IP 和 负载均衡 两种类型，可选值为 Address 和 LoadBalance。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackageResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackageResourcesResponseParams struct {
	// 符合条件的 共享带宽包内资源 数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 共享带宽包内资源 详细信息列表。
	ResourceSet []*Resource `json:"ResourceSet,omitnil" name:"ResourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackageResourcesResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackagesRequestParams struct {
	// 带宽包唯一ID列表
	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitnil" name:"BandwidthPackageIds"`

	// 每次请求的`Filters`的上限为10。参数不支持同时指定`BandwidthPackageIds`和`Filters`。详细的过滤条件如下：
	// <li> bandwidth-package_id - String - 是否必填：否 - （过滤条件）按照带宽包的唯一标识ID过滤。</li>
	// <li> bandwidth-package-name - String - 是否必填：否 - （过滤条件）按照 带宽包名称过滤。不支持模糊过滤。</li>
	// <li> network-type - String - 是否必填：否 - （过滤条件）按照带宽包的类型过滤。类型包括'HIGH_QUALITY_BGP','BGP','SINGLEISP'和'ANYCAST'。</li>
	// <li> charge-type - String - 是否必填：否 - （过滤条件）按照带宽包的计费类型过滤。计费类型包括'TOP5_POSTPAID_BY_MONTH'和'PERCENT95_POSTPAID_BY_MONTH'。</li>
	// <li> resource.resource-type - String - 是否必填：否 - （过滤条件）按照带宽包资源类型过滤。资源类型包括'Address'和'LoadBalance'</li>
	// <li> resource.resource-id - String - 是否必填：否 - （过滤条件）按照带宽包资源Id过滤。资源Id形如'eip-xxxx','lb-xxxx'</li>
	// <li> resource.address-ip - String - 是否必填：否 - （过滤条件）按照带宽包资源Ip过滤。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 - （过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。tag-key使用具体的标签键进行替换。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询带宽包偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小结。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询带宽包返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小结。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeBandwidthPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 带宽包唯一ID列表
	BandwidthPackageIds []*string `json:"BandwidthPackageIds,omitnil" name:"BandwidthPackageIds"`

	// 每次请求的`Filters`的上限为10。参数不支持同时指定`BandwidthPackageIds`和`Filters`。详细的过滤条件如下：
	// <li> bandwidth-package_id - String - 是否必填：否 - （过滤条件）按照带宽包的唯一标识ID过滤。</li>
	// <li> bandwidth-package-name - String - 是否必填：否 - （过滤条件）按照 带宽包名称过滤。不支持模糊过滤。</li>
	// <li> network-type - String - 是否必填：否 - （过滤条件）按照带宽包的类型过滤。类型包括'HIGH_QUALITY_BGP','BGP','SINGLEISP'和'ANYCAST'。</li>
	// <li> charge-type - String - 是否必填：否 - （过滤条件）按照带宽包的计费类型过滤。计费类型包括'TOP5_POSTPAID_BY_MONTH'和'PERCENT95_POSTPAID_BY_MONTH'。</li>
	// <li> resource.resource-type - String - 是否必填：否 - （过滤条件）按照带宽包资源类型过滤。资源类型包括'Address'和'LoadBalance'</li>
	// <li> resource.resource-id - String - 是否必填：否 - （过滤条件）按照带宽包资源Id过滤。资源Id形如'eip-xxxx','lb-xxxx'</li>
	// <li> resource.address-ip - String - 是否必填：否 - （过滤条件）按照带宽包资源Ip过滤。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 - （过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。tag-key使用具体的标签键进行替换。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询带宽包偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小结。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询带宽包返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小结。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeBandwidthPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBandwidthPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBandwidthPackagesResponseParams struct {
	// 符合条件的带宽包数量
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 描述带宽包详细信息
	BandwidthPackageSet []*BandwidthPackage `json:"BandwidthPackageSet,omitnil" name:"BandwidthPackageSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBandwidthPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBandwidthPackagesResponseParams `json:"Response"`
}

func (r *DescribeBandwidthPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBandwidthPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAttachedInstancesRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件：
	// <li>ccn-id - String -（过滤条件）CCN实例ID。</li>
	// <li>instance-type - String -（过滤条件）关联实例类型。</li>
	// <li>instance-region - String -（过滤条件）关联实例所属地域。</li>
	// <li>instance-id - String -（过滤条件）关联实例ID。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 云联网实例ID
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 排序字段。支持：`CcnId` `InstanceType` `InstanceId` `InstanceName` `InstanceRegion` `AttachedTime` `State`。默认值：`AttachedTime`
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

type DescribeCcnAttachedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件：
	// <li>ccn-id - String -（过滤条件）CCN实例ID。</li>
	// <li>instance-type - String -（过滤条件）关联实例类型。</li>
	// <li>instance-region - String -（过滤条件）关联实例所属地域。</li>
	// <li>instance-id - String -（过滤条件）关联实例ID。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 云联网实例ID
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 排序字段。支持：`CcnId` `InstanceType` `InstanceId` `InstanceName` `InstanceRegion` `AttachedTime` `State`。默认值：`AttachedTime`
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

func (r *DescribeCcnAttachedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAttachedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "CcnId")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnAttachedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAttachedInstancesResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 关联实例列表。
	InstanceSet []*CcnAttachedInstance `json:"InstanceSet,omitnil" name:"InstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcnAttachedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnAttachedInstancesResponseParams `json:"Response"`
}

func (r *DescribeCcnAttachedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAttachedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnRegionBandwidthLimitsRequestParams struct {
	// CCN实例ID，形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
}

type DescribeCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID，形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`
}

func (r *DescribeCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnRegionBandwidthLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnRegionBandwidthLimitsResponseParams struct {
	// 云联网（CCN）各地域出带宽上限
	CcnRegionBandwidthLimitSet []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimitSet,omitnil" name:"CcnRegionBandwidthLimitSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnRegionBandwidthLimitsResponseParams `json:"Response"`
}

func (r *DescribeCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnRoutesRequestParams struct {
	// CCN实例ID，形如：`ccn-gree226l`。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN路由策略唯一ID，形如：`ccnr-f49l6u0z`。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`

	// 过滤条件，参数不支持同时指定RouteIds和Filters。
	// <li>route-id - String -（过滤条件）路由策略ID。</li>
	// <li>cidr-block - String -（过滤条件）目的端。</li>
	// <li>instance-type - String -（过滤条件）下一跳类型。</li>
	// <li>instance-region - String -（过滤条件）下一跳所属地域。</li>
	// <li>instance-id - String -（过滤条件）下一跳实例ID。</li>
	// <li>route-table-id - String -（过滤条件）路由表ID列表，形如ccntr-1234edfr，可以根据路由表ID 过滤。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID，形如：`ccn-gree226l`。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN路由策略唯一ID，形如：`ccnr-f49l6u0z`。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`

	// 过滤条件，参数不支持同时指定RouteIds和Filters。
	// <li>route-id - String -（过滤条件）路由策略ID。</li>
	// <li>cidr-block - String -（过滤条件）目的端。</li>
	// <li>instance-type - String -（过滤条件）下一跳类型。</li>
	// <li>instance-region - String -（过滤条件）下一跳所属地域。</li>
	// <li>instance-id - String -（过滤条件）下一跳实例ID。</li>
	// <li>route-table-id - String -（过滤条件）路由表ID列表，形如ccntr-1234edfr，可以根据路由表ID 过滤。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "RouteIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnRoutesResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// CCN路由策略对象。
	RouteSet []*CcnRoute `json:"RouteSet,omitnil" name:"RouteSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnRoutesResponseParams `json:"Response"`
}

func (r *DescribeCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnsRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定CcnIds和Filters。
	CcnIds []*string `json:"CcnIds,omitnil" name:"CcnIds"`

	// 过滤条件，参数不支持同时指定CcnIds和Filters。
	// <li>ccn-id - String - （过滤条件）CCN唯一ID，形如：`ccn-f49l6u0z`。</li>
	// <li>ccn-name - String - （过滤条件）CCN名称。</li>
	// <li>ccn-description - String - （过滤条件）CCN描述。</li>
	// <li>state - String - （过滤条件）实例状态， 'ISOLATED': 隔离中（欠费停服），'AVAILABLE'：运行中。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例：查询绑定了标签的CCN列表。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持：`CcnId` `CcnName` `CreateTime` `State` `QosLevel`。默认值: `CreateTime`
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

type DescribeCcnsRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定CcnIds和Filters。
	CcnIds []*string `json:"CcnIds,omitnil" name:"CcnIds"`

	// 过滤条件，参数不支持同时指定CcnIds和Filters。
	// <li>ccn-id - String - （过滤条件）CCN唯一ID，形如：`ccn-f49l6u0z`。</li>
	// <li>ccn-name - String - （过滤条件）CCN名称。</li>
	// <li>ccn-description - String - （过滤条件）CCN描述。</li>
	// <li>state - String - （过滤条件）实例状态， 'ISOLATED': 隔离中（欠费停服），'AVAILABLE'：运行中。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例：查询绑定了标签的CCN列表。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持：`CcnId` `CcnName` `CreateTime` `State` `QosLevel`。默认值: `CreateTime`
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

func (r *DescribeCcnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnsResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// CCN对象。
	CcnSet []*CCN `json:"CcnSet,omitnil" name:"CcnSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCcnsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnsResponseParams `json:"Response"`
}

func (r *DescribeCcnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicLinkInstancesRequestParams struct {
	// 过滤条件。
	// <li>vpc-id - String - （过滤条件）VPC实例ID。</li>
	// <li>vm-ip - String - （过滤条件）基础网络云服务器IP。</li>
	Filters []*FilterObject `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认值0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

type DescribeClassicLinkInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>vpc-id - String - （过滤条件）VPC实例ID。</li>
	// <li>vm-ip - String - （过滤条件）基础网络云服务器IP。</li>
	Filters []*FilterObject `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认值0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeClassicLinkInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicLinkInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicLinkInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicLinkInstancesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 私有网络和基础网络互通设备。
	ClassicLinkInstanceSet []*ClassicLinkInstance `json:"ClassicLinkInstanceSet,omitnil" name:"ClassicLinkInstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClassicLinkInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicLinkInstancesResponseParams `json:"Response"`
}

func (r *DescribeClassicLinkInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicLinkInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderCcnRegionBandwidthLimitsRequestParams struct {
	// 过滤条件，目前`value`值个数只支持一个，可支持的字段有：
	// <li>`source-region` 源地域，值形如：`["ap-guangzhou"]`</li> <li>`destination-region` 目的地域，值形如：`["ap-shanghai"]`</li> <li>`ccn-ids` 云联网ID数组，值形如：`["ccn-12345678"]`</li> <li>`user-account-id` 用户账号ID，值形如`["12345678"]`</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数据量可选值0到100之间的整数，默认20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeCrossBorderCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，目前`value`值个数只支持一个，可支持的字段有：
	// <li>`source-region` 源地域，值形如：`["ap-guangzhou"]`</li> <li>`destination-region` 目的地域，值形如：`["ap-shanghai"]`</li> <li>`ccn-ids` 云联网ID数组，值形如：`["ccn-12345678"]`</li> <li>`user-account-id` 用户账号ID，值形如`["12345678"]`</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数据量可选值0到100之间的整数，默认20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeCrossBorderCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossBorderCcnRegionBandwidthLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderCcnRegionBandwidthLimitsResponseParams struct {
	// 符合条件的对象总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 云联网地域间限速带宽实例的信息。
	CcnBandwidthSet []*CcnBandwidth `json:"CcnBandwidthSet,omitnil" name:"CcnBandwidthSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCrossBorderCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossBorderCcnRegionBandwidthLimitsResponseParams `json:"Response"`
}

func (r *DescribeCrossBorderCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderComplianceRequestParams struct {
	// （精确匹配）服务商，可选值：`UNICOM`。
	ServiceProvider *string `json:"ServiceProvider,omitnil" name:"ServiceProvider"`

	// （精确匹配）合规化审批单`ID`。
	ComplianceId *uint64 `json:"ComplianceId,omitnil" name:"ComplianceId"`

	// （模糊查询）公司名称。
	Company *string `json:"Company,omitnil" name:"Company"`

	// （精确匹配）统一社会信用代码。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// （模糊查询）法定代表人。
	LegalPerson *string `json:"LegalPerson,omitnil" name:"LegalPerson"`

	// （模糊查询）发证机关。
	IssuingAuthority *string `json:"IssuingAuthority,omitnil" name:"IssuingAuthority"`

	// （模糊查询）营业执照住所。
	BusinessAddress *string `json:"BusinessAddress,omitnil" name:"BusinessAddress"`

	// （精确匹配）邮编。
	PostCode *uint64 `json:"PostCode,omitnil" name:"PostCode"`

	// （模糊查询）经办人。
	Manager *string `json:"Manager,omitnil" name:"Manager"`

	// （精确查询）经办人身份证号。
	ManagerId *string `json:"ManagerId,omitnil" name:"ManagerId"`

	// （模糊查询）经办人身份证地址。
	ManagerAddress *string `json:"ManagerAddress,omitnil" name:"ManagerAddress"`

	// （精确匹配）经办人联系电话。
	ManagerTelephone *string `json:"ManagerTelephone,omitnil" name:"ManagerTelephone"`

	// （精确匹配）电子邮箱。
	Email *string `json:"Email,omitnil" name:"Email"`

	// （精确匹配）服务开始日期，如：`2020-07-28`。
	ServiceStartDate *string `json:"ServiceStartDate,omitnil" name:"ServiceStartDate"`

	// （精确匹配）服务结束日期，如：`2021-07-28`。
	ServiceEndDate *string `json:"ServiceEndDate,omitnil" name:"ServiceEndDate"`

	// （精确匹配）状态。待审批：`PENDING`，通过：`APPROVED `，拒绝：`DENY`。
	State *string `json:"State,omitnil" name:"State"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeCrossBorderComplianceRequest struct {
	*tchttp.BaseRequest
	
	// （精确匹配）服务商，可选值：`UNICOM`。
	ServiceProvider *string `json:"ServiceProvider,omitnil" name:"ServiceProvider"`

	// （精确匹配）合规化审批单`ID`。
	ComplianceId *uint64 `json:"ComplianceId,omitnil" name:"ComplianceId"`

	// （模糊查询）公司名称。
	Company *string `json:"Company,omitnil" name:"Company"`

	// （精确匹配）统一社会信用代码。
	UniformSocialCreditCode *string `json:"UniformSocialCreditCode,omitnil" name:"UniformSocialCreditCode"`

	// （模糊查询）法定代表人。
	LegalPerson *string `json:"LegalPerson,omitnil" name:"LegalPerson"`

	// （模糊查询）发证机关。
	IssuingAuthority *string `json:"IssuingAuthority,omitnil" name:"IssuingAuthority"`

	// （模糊查询）营业执照住所。
	BusinessAddress *string `json:"BusinessAddress,omitnil" name:"BusinessAddress"`

	// （精确匹配）邮编。
	PostCode *uint64 `json:"PostCode,omitnil" name:"PostCode"`

	// （模糊查询）经办人。
	Manager *string `json:"Manager,omitnil" name:"Manager"`

	// （精确查询）经办人身份证号。
	ManagerId *string `json:"ManagerId,omitnil" name:"ManagerId"`

	// （模糊查询）经办人身份证地址。
	ManagerAddress *string `json:"ManagerAddress,omitnil" name:"ManagerAddress"`

	// （精确匹配）经办人联系电话。
	ManagerTelephone *string `json:"ManagerTelephone,omitnil" name:"ManagerTelephone"`

	// （精确匹配）电子邮箱。
	Email *string `json:"Email,omitnil" name:"Email"`

	// （精确匹配）服务开始日期，如：`2020-07-28`。
	ServiceStartDate *string `json:"ServiceStartDate,omitnil" name:"ServiceStartDate"`

	// （精确匹配）服务结束日期，如：`2021-07-28`。
	ServiceEndDate *string `json:"ServiceEndDate,omitnil" name:"ServiceEndDate"`

	// （精确匹配）状态。待审批：`PENDING`，通过：`APPROVED `，拒绝：`DENY`。
	State *string `json:"State,omitnil" name:"State"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeCrossBorderComplianceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderComplianceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProvider")
	delete(f, "ComplianceId")
	delete(f, "Company")
	delete(f, "UniformSocialCreditCode")
	delete(f, "LegalPerson")
	delete(f, "IssuingAuthority")
	delete(f, "BusinessAddress")
	delete(f, "PostCode")
	delete(f, "Manager")
	delete(f, "ManagerId")
	delete(f, "ManagerAddress")
	delete(f, "ManagerTelephone")
	delete(f, "Email")
	delete(f, "ServiceStartDate")
	delete(f, "ServiceEndDate")
	delete(f, "State")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossBorderComplianceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderComplianceResponseParams struct {
	// 合规化审批单列表。
	CrossBorderComplianceSet []*CrossBorderCompliance `json:"CrossBorderComplianceSet,omitnil" name:"CrossBorderComplianceSet"`

	// 合规化审批单总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCrossBorderComplianceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossBorderComplianceResponseParams `json:"Response"`
}

func (r *DescribeCrossBorderComplianceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderComplianceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderFlowMonitorRequestParams struct {
	// 源地域。
	SourceRegion *string `json:"SourceRegion,omitnil" name:"SourceRegion"`

	// 目的地域。
	DestinationRegion *string `json:"DestinationRegion,omitnil" name:"DestinationRegion"`

	// 云联网ID。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网所属账号。
	CcnUin *string `json:"CcnUin,omitnil" name:"CcnUin"`

	// 时间粒度。单位为:秒，如60为60s的时间粒度
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeCrossBorderFlowMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 源地域。
	SourceRegion *string `json:"SourceRegion,omitnil" name:"SourceRegion"`

	// 目的地域。
	DestinationRegion *string `json:"DestinationRegion,omitnil" name:"DestinationRegion"`

	// 云联网ID。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网所属账号。
	CcnUin *string `json:"CcnUin,omitnil" name:"CcnUin"`

	// 时间粒度。单位为:秒，如60为60s的时间粒度
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeCrossBorderFlowMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderFlowMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceRegion")
	delete(f, "DestinationRegion")
	delete(f, "CcnId")
	delete(f, "CcnUin")
	delete(f, "Period")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossBorderFlowMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossBorderFlowMonitorResponseParams struct {
	// 云联网跨境带宽监控数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossBorderFlowMonitorData []*CrossBorderFlowMonitorData `json:"CrossBorderFlowMonitorData,omitnil" name:"CrossBorderFlowMonitorData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCrossBorderFlowMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossBorderFlowMonitorResponseParams `json:"Response"`
}

func (r *DescribeCrossBorderFlowMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossBorderFlowMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewayVendorsRequestParams struct {

}

type DescribeCustomerGatewayVendorsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCustomerGatewayVendorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewayVendorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerGatewayVendorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewayVendorsResponseParams struct {
	// 对端网关厂商信息对象。
	CustomerGatewayVendorSet []*CustomerGatewayVendor `json:"CustomerGatewayVendorSet,omitnil" name:"CustomerGatewayVendorSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomerGatewayVendorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerGatewayVendorsResponseParams `json:"Response"`
}

func (r *DescribeCustomerGatewayVendorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewayVendorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewaysRequestParams struct {
	// 对端网关ID，例如：cgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitnil" name:"CustomerGatewayIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	// <li>customer-gateway-id - String - （过滤条件）用户网关唯一ID形如：`cgw-mgp33pll`。</li>
	// <li>customer-gateway-name - String - （过滤条件）用户网关名称形如：`test-cgw`。</li>
	// <li>ip-address - String - （过滤条件）公网地址形如：`58.211.1.12`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeCustomerGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 对端网关ID，例如：cgw-2wqq41m9。每次请求的实例的上限为100。参数不支持同时指定CustomerGatewayIds和Filters。
	CustomerGatewayIds []*string `json:"CustomerGatewayIds,omitnil" name:"CustomerGatewayIds"`

	// 过滤条件，详见下表：实例过滤条件表。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定CustomerGatewayIds和Filters。
	// <li>customer-gateway-id - String - （过滤条件）用户网关唯一ID形如：`cgw-mgp33pll`。</li>
	// <li>customer-gateway-name - String - （过滤条件）用户网关名称形如：`test-cgw`。</li>
	// <li>ip-address - String - （过滤条件）公网地址形如：`58.211.1.12`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeCustomerGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomerGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomerGatewaysResponseParams struct {
	// 对端网关对象列表。
	CustomerGatewaySet []*CustomerGateway `json:"CustomerGatewaySet,omitnil" name:"CustomerGatewaySet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomerGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomerGatewaysResponseParams `json:"Response"`
}

func (r *DescribeCustomerGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomerGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDhcpIpsRequestParams struct {
	// DhcpIp实例ID。形如：dhcpip-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定DhcpIpIds和Filters。
	DhcpIpIds []*string `json:"DhcpIpIds,omitnil" name:"DhcpIpIds"`

	// 过滤条件，参数不支持同时指定DhcpIpIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。</li>
	// <li>dhcpip-id - String - （过滤条件）DhcpIp实例ID，形如：dhcpip-pxir56ns。</li>
	// <li>dhcpip-name - String - （过滤条件）DhcpIp实例名称。</li>
	// <li>address-ip - String - （过滤条件）DhcpIp实例的IP，根据IP精确查找。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDhcpIpsRequest struct {
	*tchttp.BaseRequest
	
	// DhcpIp实例ID。形如：dhcpip-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定DhcpIpIds和Filters。
	DhcpIpIds []*string `json:"DhcpIpIds,omitnil" name:"DhcpIpIds"`

	// 过滤条件，参数不支持同时指定DhcpIpIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。</li>
	// <li>dhcpip-id - String - （过滤条件）DhcpIp实例ID，形如：dhcpip-pxir56ns。</li>
	// <li>dhcpip-name - String - （过滤条件）DhcpIp实例名称。</li>
	// <li>address-ip - String - （过滤条件）DhcpIp实例的IP，根据IP精确查找。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDhcpIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDhcpIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DhcpIpIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDhcpIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDhcpIpsResponseParams struct {
	// 实例详细信息列表。
	DhcpIpSet []*DhcpIp `json:"DhcpIpSet,omitnil" name:"DhcpIpSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDhcpIpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDhcpIpsResponseParams `json:"Response"`
}

func (r *DescribeDhcpIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDhcpIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewayCcnRoutesRequestParams struct {
	// 专线网关ID，形如：`dcg-prpqlmg1`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 云联网路由学习类型，可选值：
	// <li>`BGP` - 自动学习。</li>
	// <li>`STATIC` - 静态，即用户配置，默认值。</li>
	CcnRouteType *string `json:"CcnRouteType,omitnil" name:"CcnRouteType"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关ID，形如：`dcg-prpqlmg1`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 云联网路由学习类型，可选值：
	// <li>`BGP` - 自动学习。</li>
	// <li>`STATIC` - 静态，即用户配置，默认值。</li>
	CcnRouteType *string `json:"CcnRouteType,omitnil" name:"CcnRouteType"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "CcnRouteType")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewayCcnRoutesResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 云联网路由（IDC网段）列表。
	RouteSet []*DirectConnectGatewayCcnRoute `json:"RouteSet,omitnil" name:"RouteSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewaysRequestParams struct {
	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitnil" name:"DirectConnectGatewayIds"`

	// 过滤条件，参数不支持同时指定`DirectConnectGatewayIds`和`Filters`。
	// <li>direct-connect-gateway-id - String - 专线网关唯一`ID`，形如：`dcg-9o233uri`。</li>
	// <li>direct-connect-gateway-name - String - 专线网关名称，默认模糊查询。</li>
	// <li>direct-connect-gateway-ip - String - 专线网关`IP`。</li>
	// <li>gateway-type - String - 网关类型，可选值：`NORMAL`（普通型）、`NAT`（NAT型）。</li>
	// <li>network-type- String - 网络类型，可选值：`VPC`（私有网络类型）、`CCN`（云联网类型）。</li>
	// <li>ccn-id - String - 专线网关所在云联网`ID`。</li>
	// <li>vpc-id - String - 专线网关所在私有网络`ID`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDirectConnectGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitnil" name:"DirectConnectGatewayIds"`

	// 过滤条件，参数不支持同时指定`DirectConnectGatewayIds`和`Filters`。
	// <li>direct-connect-gateway-id - String - 专线网关唯一`ID`，形如：`dcg-9o233uri`。</li>
	// <li>direct-connect-gateway-name - String - 专线网关名称，默认模糊查询。</li>
	// <li>direct-connect-gateway-ip - String - 专线网关`IP`。</li>
	// <li>gateway-type - String - 网关类型，可选值：`NORMAL`（普通型）、`NAT`（NAT型）。</li>
	// <li>network-type- String - 网络类型，可选值：`VPC`（私有网络类型）、`CCN`（云联网类型）。</li>
	// <li>ccn-id - String - 专线网关所在云联网`ID`。</li>
	// <li>vpc-id - String - 专线网关所在私有网络`ID`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDirectConnectGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDirectConnectGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDirectConnectGatewaysResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 专线网关对象数组。
	DirectConnectGatewaySet []*DirectConnectGateway `json:"DirectConnectGatewaySet,omitnil" name:"DirectConnectGatewaySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDirectConnectGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDirectConnectGatewaysResponseParams `json:"Response"`
}

func (r *DescribeDirectConnectGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDirectConnectGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogRequestParams struct {
	// 私用网络ID或者统一ID，建议使用统一ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`
}

type DescribeFlowLogRequest struct {
	*tchttp.BaseRequest
	
	// 私用网络ID或者统一ID，建议使用统一ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`
}

func (r *DescribeFlowLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "FlowLogId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogResponseParams struct {
	// 流日志信息。
	FlowLog []*FlowLog `json:"FlowLog,omitnil" name:"FlowLog"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowLogResponseParams `json:"Response"`
}

func (r *DescribeFlowLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogsRequestParams struct {
	// 私用网络ID或者统一ID，建议使用统一ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`

	// 流日志实例名字。
	FlowLogName *string `json:"FlowLogName,omitnil" name:"FlowLogName"`

	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源唯一ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流日志采集类型，ACCEPT|REJECT|ALL。
	TrafficType *string `json:"TrafficType,omitnil" name:"TrafficType"`

	// 流日志存储ID。
	CloudLogId *string `json:"CloudLogId,omitnil" name:"CloudLogId"`

	// 流日志存储ID状态。
	CloudLogState *string `json:"CloudLogState,omitnil" name:"CloudLogState"`

	// 按某个字段排序,支持字段：flowLogName,createTime，默认按createTime。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 升序（asc）还是降序（desc）,默认：desc。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页行数，默认为10。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，参数不支持同时指定FlowLogId和Filters。
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>
	Filters *Filter `json:"Filters,omitnil" name:"Filters"`

	// 流日志存储ID对应的地域信息。
	CloudLogRegion *string `json:"CloudLogRegion,omitnil" name:"CloudLogRegion"`
}

type DescribeFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 私用网络ID或者统一ID，建议使用统一ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`

	// 流日志实例名字。
	FlowLogName *string `json:"FlowLogName,omitnil" name:"FlowLogName"`

	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源唯一ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流日志采集类型，ACCEPT|REJECT|ALL。
	TrafficType *string `json:"TrafficType,omitnil" name:"TrafficType"`

	// 流日志存储ID。
	CloudLogId *string `json:"CloudLogId,omitnil" name:"CloudLogId"`

	// 流日志存储ID状态。
	CloudLogState *string `json:"CloudLogState,omitnil" name:"CloudLogState"`

	// 按某个字段排序,支持字段：flowLogName,createTime，默认按createTime。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 升序（asc）还是降序（desc）,默认：desc。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页行数，默认为10。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，参数不支持同时指定FlowLogId和Filters。
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。</li>
	Filters *Filter `json:"Filters,omitnil" name:"Filters"`

	// 流日志存储ID对应的地域信息。
	CloudLogRegion *string `json:"CloudLogRegion,omitnil" name:"CloudLogRegion"`
}

func (r *DescribeFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "FlowLogId")
	delete(f, "FlowLogName")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "TrafficType")
	delete(f, "CloudLogId")
	delete(f, "CloudLogState")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "CloudLogRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowLogsResponseParams struct {
	// 流日志实例集合。
	FlowLog []*FlowLog `json:"FlowLog,omitnil" name:"FlowLog"`

	// 流日志总数目。
	TotalNum *uint64 `json:"TotalNum,omitnil" name:"TotalNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowLogsResponseParams `json:"Response"`
}

func (r *DescribeFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowMonitorDetailRequestParams struct {
	// 时间点。表示要查询这分钟内的明细。如：`2019-02-28 18:15:20`，将查询 `18:15` 这一分钟内的明细。
	TimePoint *string `json:"TimePoint,omitnil" name:"TimePoint"`

	// VPN网关实例ID，形如：`vpn-ltjahce6`。
	VpnId *string `json:"VpnId,omitnil" name:"VpnId"`

	// 专线网关实例ID，形如：`dcg-ltjahce6`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 对等连接实例ID，形如：`pcx-ltjahce6`。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`

	// NAT网关实例ID，形如：`nat-ltjahce6`。
	NatId *string `json:"NatId,omitnil" name:"NatId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持 `InPkg` `OutPkg` `InTraffic` `OutTraffic`。默认值`OutTraffic`。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。顺序：`ASC`，倒序：`DESC`。默认值`DESC`。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

type DescribeGatewayFlowMonitorDetailRequest struct {
	*tchttp.BaseRequest
	
	// 时间点。表示要查询这分钟内的明细。如：`2019-02-28 18:15:20`，将查询 `18:15` 这一分钟内的明细。
	TimePoint *string `json:"TimePoint,omitnil" name:"TimePoint"`

	// VPN网关实例ID，形如：`vpn-ltjahce6`。
	VpnId *string `json:"VpnId,omitnil" name:"VpnId"`

	// 专线网关实例ID，形如：`dcg-ltjahce6`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 对等连接实例ID，形如：`pcx-ltjahce6`。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`

	// NAT网关实例ID，形如：`nat-ltjahce6`。
	NatId *string `json:"NatId,omitnil" name:"NatId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持 `InPkg` `OutPkg` `InTraffic` `OutTraffic`。默认值`OutTraffic`。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。顺序：`ASC`，倒序：`DESC`。默认值`DESC`。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

func (r *DescribeGatewayFlowMonitorDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowMonitorDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimePoint")
	delete(f, "VpnId")
	delete(f, "DirectConnectGatewayId")
	delete(f, "PeeringConnectionId")
	delete(f, "NatId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayFlowMonitorDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowMonitorDetailResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 网关流量监控明细。
	GatewayFlowMonitorDetailSet []*GatewayFlowMonitorDetail `json:"GatewayFlowMonitorDetailSet,omitnil" name:"GatewayFlowMonitorDetailSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGatewayFlowMonitorDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayFlowMonitorDetailResponseParams `json:"Response"`
}

func (r *DescribeGatewayFlowMonitorDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowMonitorDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowQosRequestParams struct {
	// 网关实例ID，目前我们支持的网关实例类型有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 限流的云服务器内网IP。
	IpAddresses []*string `json:"IpAddresses,omitnil" name:"IpAddresses"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeGatewayFlowQosRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID，目前我们支持的网关实例类型有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 限流的云服务器内网IP。
	IpAddresses []*string `json:"IpAddresses,omitnil" name:"IpAddresses"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeGatewayFlowQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "IpAddresses")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayFlowQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayFlowQosResponseParams struct {
	// 实例详细信息列表。
	GatewayQosSet []*GatewayQos `json:"GatewayQosSet,omitnil" name:"GatewayQosSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeGatewayFlowQosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayFlowQosResponseParams `json:"Response"`
}

func (r *DescribeGatewayFlowQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayFlowQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsRequestParams struct {
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。
	HaVipIds []*string `json:"HaVipIds,omitnil" name:"HaVipIds"`

	// 过滤条件，参数不支持同时指定`HaVipIds`和`Filters`。
	// <li>havip-id - String - `HAVIP`唯一`ID`，形如：`havip-9o233uri`。</li>
	// <li>havip-name - String - `HAVIP`名称。</li>
	// <li>vpc-id - String - `HAVIP`所在私有网络`ID`。</li>
	// <li>subnet-id - String - `HAVIP`所在子网`ID`。</li>
	// <li>vip - String - `HAVIP`的地址`VIP`。</li>
	// <li>address-ip - String - `HAVIP`绑定的弹性公网`IP`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeHaVipsRequest struct {
	*tchttp.BaseRequest
	
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。
	HaVipIds []*string `json:"HaVipIds,omitnil" name:"HaVipIds"`

	// 过滤条件，参数不支持同时指定`HaVipIds`和`Filters`。
	// <li>havip-id - String - `HAVIP`唯一`ID`，形如：`havip-9o233uri`。</li>
	// <li>havip-name - String - `HAVIP`名称。</li>
	// <li>vpc-id - String - `HAVIP`所在私有网络`ID`。</li>
	// <li>subnet-id - String - `HAVIP`所在子网`ID`。</li>
	// <li>vip - String - `HAVIP`的地址`VIP`。</li>
	// <li>address-ip - String - `HAVIP`绑定的弹性公网`IP`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeHaVipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHaVipsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHaVipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHaVipsResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// `HAVIP`对象数组。
	HaVipSet []*HaVip `json:"HaVipSet,omitnil" name:"HaVipSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHaVipsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHaVipsResponseParams `json:"Response"`
}

func (r *DescribeHaVipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHaVipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIp6AddressesRequestParams struct {
	// 标识 IPV6 的唯一 ID 列表。IPV6 唯一 ID 形如：`eip-11112222`。参数不支持同时指定`Ip6AddressIds`和`Filters`。
	Ip6AddressIds []*string `json:"Ip6AddressIds,omitnil" name:"Ip6AddressIds"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。</li>
	// <li> network-interface-id - String - 是否必填：否 - （过滤条件）按照弹性网卡的唯一ID过滤。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeIp6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// 标识 IPV6 的唯一 ID 列表。IPV6 唯一 ID 形如：`eip-11112222`。参数不支持同时指定`Ip6AddressIds`和`Filters`。
	Ip6AddressIds []*string `json:"Ip6AddressIds,omitnil" name:"Ip6AddressIds"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AddressIds`和`Filters`。详细的过滤条件如下：
	// <li> address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。</li>
	// <li> network-interface-id - String - 是否必填：否 - （过滤条件）按照弹性网卡的唯一ID过滤。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeIp6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIp6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6AddressIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIp6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIp6AddressesResponseParams struct {
	// 符合条件的 IPV6 数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// IPV6 详细信息列表。
	AddressSet []*Address `json:"AddressSet,omitnil" name:"AddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIp6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIp6AddressesResponseParams `json:"Response"`
}

func (r *DescribeIp6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIp6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIp6TranslatorQuotaRequestParams struct {
	// 待查询IPV6转换实例的唯一ID列表，形如ip6-xxxxxxxx
	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitnil" name:"Ip6TranslatorIds"`
}

type DescribeIp6TranslatorQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 待查询IPV6转换实例的唯一ID列表，形如ip6-xxxxxxxx
	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitnil" name:"Ip6TranslatorIds"`
}

func (r *DescribeIp6TranslatorQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIp6TranslatorQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIp6TranslatorQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIp6TranslatorQuotaResponseParams struct {
	// 账户在指定地域的IPV6转换实例及规则配额信息
	// QUOTAID属性是TOTAL_TRANSLATOR_QUOTA，表示账户在指定地域的IPV6转换实例配额信息；QUOTAID属性是IPV6转换实例唯一ID（形如ip6-xxxxxxxx），表示账户在该转换实例允许创建的转换规则配额
	QuotaSet []*Quota `json:"QuotaSet,omitnil" name:"QuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIp6TranslatorQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIp6TranslatorQuotaResponseParams `json:"Response"`
}

func (r *DescribeIp6TranslatorQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIp6TranslatorQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIp6TranslatorsRequestParams struct {
	// IPV6转换实例唯一ID数组，形如ip6-xxxxxxxx
	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitnil" name:"Ip6TranslatorIds"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`Ip6TranslatorIds`和`Filters`。详细的过滤条件如下：
	// <li> ip6-translator-id - String - 是否必填：否 - （过滤条件）按照IPV6转换实例的唯一ID过滤,形如ip6-xxxxxxx。</li>
	// <li> ip6-translator-vip6 - String - 是否必填：否 - （过滤条件）按照IPV6地址过滤。不支持模糊过滤。</li>
	// <li> ip6-translator-name - String - 是否必填：否 - （过滤条件）按照IPV6转换实例名称过滤。不支持模糊过滤。</li>
	// <li> ip6-translator-status - String - 是否必填：否 - （过滤条件）按照IPV6转换实例的状态过滤。状态取值范围为"CREATING","RUNNING","DELETING","MODIFYING"
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeIp6TranslatorsRequest struct {
	*tchttp.BaseRequest
	
	// IPV6转换实例唯一ID数组，形如ip6-xxxxxxxx
	Ip6TranslatorIds []*string `json:"Ip6TranslatorIds,omitnil" name:"Ip6TranslatorIds"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`Ip6TranslatorIds`和`Filters`。详细的过滤条件如下：
	// <li> ip6-translator-id - String - 是否必填：否 - （过滤条件）按照IPV6转换实例的唯一ID过滤,形如ip6-xxxxxxx。</li>
	// <li> ip6-translator-vip6 - String - 是否必填：否 - （过滤条件）按照IPV6地址过滤。不支持模糊过滤。</li>
	// <li> ip6-translator-name - String - 是否必填：否 - （过滤条件）按照IPV6转换实例名称过滤。不支持模糊过滤。</li>
	// <li> ip6-translator-status - String - 是否必填：否 - （过滤条件）按照IPV6转换实例的状态过滤。状态取值范围为"CREATING","RUNNING","DELETING","MODIFYING"
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeIp6TranslatorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIp6TranslatorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIp6TranslatorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIp6TranslatorsResponseParams struct {
	// 符合过滤条件的IPV6转换实例数量。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 符合过滤条件的IPV6转换实例详细信息
	Ip6TranslatorSet []*Ip6Translator `json:"Ip6TranslatorSet,omitnil" name:"Ip6TranslatorSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIp6TranslatorsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIp6TranslatorsResponseParams `json:"Response"`
}

func (r *DescribeIp6TranslatorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIp6TranslatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpGeolocationDatabaseUrlRequestParams struct {
	// IP地理位置库协议类型，目前仅支持"ipv4"。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeIpGeolocationDatabaseUrlRequest struct {
	*tchttp.BaseRequest
	
	// IP地理位置库协议类型，目前仅支持"ipv4"。
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *DescribeIpGeolocationDatabaseUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationDatabaseUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpGeolocationDatabaseUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpGeolocationDatabaseUrlResponseParams struct {
	// IP地理位置库下载链接地址。
	DownLoadUrl *string `json:"DownLoadUrl,omitnil" name:"DownLoadUrl"`

	// 链接到期时间。按照`ISO8601`标准表示，并且使用`UTC`时间。
	ExpiredAt *string `json:"ExpiredAt,omitnil" name:"ExpiredAt"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIpGeolocationDatabaseUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpGeolocationDatabaseUrlResponseParams `json:"Response"`
}

func (r *DescribeIpGeolocationDatabaseUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationDatabaseUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpGeolocationInfosRequestParams struct {
	// 需查询的IP地址列表，目前仅支持IPv4地址。查询的IP地址数量上限为100个。
	AddressIps []*string `json:"AddressIps,omitnil" name:"AddressIps"`

	// 需查询的IP地址的字段信息。
	Fields *IpField `json:"Fields,omitnil" name:"Fields"`
}

type DescribeIpGeolocationInfosRequest struct {
	*tchttp.BaseRequest
	
	// 需查询的IP地址列表，目前仅支持IPv4地址。查询的IP地址数量上限为100个。
	AddressIps []*string `json:"AddressIps,omitnil" name:"AddressIps"`

	// 需查询的IP地址的字段信息。
	Fields *IpField `json:"Fields,omitnil" name:"Fields"`
}

func (r *DescribeIpGeolocationInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIps")
	delete(f, "Fields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpGeolocationInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpGeolocationInfosResponseParams struct {
	// IP地址信息列表。
	AddressInfo []*IpGeolocationInfo `json:"AddressInfo,omitnil" name:"AddressInfo"`

	// IP地址信息个数。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIpGeolocationInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpGeolocationInfosResponseParams `json:"Response"`
}

func (r *DescribeIpGeolocationInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpGeolocationInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalGatewayRequestParams struct {
	// 查询条件：
	// vpc-id：按照VPCID过滤，local-gateway-name：按照本地网关名称过滤，名称支持模糊搜索，local-gateway-id：按照本地网关实例ID过滤，cdc-id：按照cdc实例ID过滤查询。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件：
	// vpc-id：按照VPCID过滤，local-gateway-name：按照本地网关名称过滤，名称支持模糊搜索，local-gateway-id：按照本地网关实例ID过滤，cdc-id：按照cdc实例ID过滤查询。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/11646)中的相关小节。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLocalGatewayResponseParams struct {
	// 本地网关信息集合。
	LocalGatewaySet []*LocalGateway `json:"LocalGatewaySet,omitnil" name:"LocalGatewaySet"`

	// 本地网关总数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLocalGatewayResponseParams `json:"Response"`
}

func (r *DescribeLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDestinationIpPortTranslationNatRulesRequestParams struct {
	// NAT网关ID。
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil" name:"NatGatewayIds"`

	// 过滤条件:
	// 参数不支持同时指定NatGatewayIds和Filters。每次请求的Filters的上限为10，Filter.Values的上限为5
	// <li> nat-gateway-id，NAT网关的ID，如`nat-0yi4hekt`</li>
	// <li> vpc-id，私有网络VPC的ID，如`vpc-0yi4hekt`</li>
	// <li> public-ip-address， 弹性IP，如`139.199.232.238`。</li>
	// <li>public-port， 公网端口。</li>
	// <li>private-ip-address， 内网IP，如`10.0.0.1`。</li>
	// <li>private-port， 内网端口。</li>
	// <li>description，规则描述。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID。
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil" name:"NatGatewayIds"`

	// 过滤条件:
	// 参数不支持同时指定NatGatewayIds和Filters。每次请求的Filters的上限为10，Filter.Values的上限为5
	// <li> nat-gateway-id，NAT网关的ID，如`nat-0yi4hekt`</li>
	// <li> vpc-id，私有网络VPC的ID，如`vpc-0yi4hekt`</li>
	// <li> public-ip-address， 弹性IP，如`139.199.232.238`。</li>
	// <li>public-port， 公网端口。</li>
	// <li>private-ip-address， 内网IP，如`10.0.0.1`。</li>
	// <li>private-port， 内网端口。</li>
	// <li>description，规则描述。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewayDestinationIpPortTranslationNatRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDestinationIpPortTranslationNatRulesResponseParams struct {
	// NAT网关端口转发规则对象数组。
	NatGatewayDestinationIpPortTranslationNatRuleSet []*NatGatewayDestinationIpPortTranslationNatRule `json:"NatGatewayDestinationIpPortTranslationNatRuleSet,omitnil" name:"NatGatewayDestinationIpPortTranslationNatRuleSet"`

	// 符合条件的NAT网关端口转发规则对象数目。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponseParams `json:"Response"`
}

func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDirectConnectGatewayRouteRequestParams struct {
	// nat的唯一标识
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// vpc的唯一标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 0到200之间
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 大于0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeNatGatewayDirectConnectGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// nat的唯一标识
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// vpc的唯一标识
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 0到200之间
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 大于0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeNatGatewayDirectConnectGatewayRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDirectConnectGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "VpcId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewayDirectConnectGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewayDirectConnectGatewayRouteResponseParams struct {
	// 路由数据
	NatDirectConnectGatewayRouteSet []*NatDirectConnectGatewayRoute `json:"NatDirectConnectGatewayRouteSet,omitnil" name:"NatDirectConnectGatewayRouteSet"`

	// 路由总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNatGatewayDirectConnectGatewayRouteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewayDirectConnectGatewayRouteResponseParams `json:"Response"`
}

func (r *DescribeNatGatewayDirectConnectGatewayRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewayDirectConnectGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaySourceIpTranslationNatRulesRequestParams struct {
	// NAT网关统一 ID，形如：`nat-123xx454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 过滤条件:
	// <li> resource-id，Subnet的ID或者Cvm ID，如`subnet-0yi4hekt`</li>
	// <li> public-ip-address，弹性IP，如`139.199.232.238`</li>
	// <li>description，规则描述。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNatGatewaySourceIpTranslationNatRulesRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关统一 ID，形如：`nat-123xx454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 过滤条件:
	// <li> resource-id，Subnet的ID或者Cvm ID，如`subnet-0yi4hekt`</li>
	// <li> public-ip-address，弹性IP，如`139.199.232.238`</li>
	// <li>description，规则描述。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaySourceIpTranslationNatRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewaySourceIpTranslationNatRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaySourceIpTranslationNatRulesResponseParams struct {
	// NAT网关SNAT规则对象数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceIpTranslationNatRuleSet []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRuleSet,omitnil" name:"SourceIpTranslationNatRuleSet"`

	// 符合条件的NAT网关端口转发规则对象数目。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNatGatewaySourceIpTranslationNatRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewaySourceIpTranslationNatRulesResponseParams `json:"Response"`
}

func (r *DescribeNatGatewaySourceIpTranslationNatRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaySourceIpTranslationNatRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaysRequestParams struct {
	// NAT网关统一 ID，形如：`nat-123xx454`。每次请求的实例上限为100。参数不支持同时指定NatGatewayIds和Filters。
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil" name:"NatGatewayIds"`

	// 过滤条件，参数不支持同时指定NatGatewayIds和Filters。每次请求的Filters的上限为10，Filter.Values的上限为5。
	// <li>nat-gateway-id - String - （过滤条件）协议端口模板实例ID，形如：`nat-123xx454`。</li>
	// <li>vpc-id - String - （过滤条件）私有网络 唯一ID，形如：`vpc-123xx454`。</li>
	// <li>nat-gateway-name - String - （过滤条件）协议端口模板实例ID，形如：`test_nat`。</li>
	// <li>tag-key - String - （过滤条件）标签键，形如：`test-key`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNatGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关统一 ID，形如：`nat-123xx454`。每次请求的实例上限为100。参数不支持同时指定NatGatewayIds和Filters。
	NatGatewayIds []*string `json:"NatGatewayIds,omitnil" name:"NatGatewayIds"`

	// 过滤条件，参数不支持同时指定NatGatewayIds和Filters。每次请求的Filters的上限为10，Filter.Values的上限为5。
	// <li>nat-gateway-id - String - （过滤条件）协议端口模板实例ID，形如：`nat-123xx454`。</li>
	// <li>vpc-id - String - （过滤条件）私有网络 唯一ID，形如：`vpc-123xx454`。</li>
	// <li>nat-gateway-name - String - （过滤条件）协议端口模板实例ID，形如：`test_nat`。</li>
	// <li>tag-key - String - （过滤条件）标签键，形如：`test-key`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeNatGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatGatewaysResponseParams struct {
	// NAT网关对象数组。
	NatGatewaySet []*NatGateway `json:"NatGatewaySet,omitnil" name:"NatGatewaySet"`

	// 符合条件的NAT网关对象个数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNatGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatGatewaysResponseParams `json:"Response"`
}

func (r *DescribeNatGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectStatesRequestParams struct {
	// 网络探测实例`ID`数组。形如：[`netd-12345678`]。
	NetDetectIds []*string `json:"NetDetectIds,omitnil" name:"NetDetectIds"`

	// 过滤条件，参数不支持同时指定NetDetectIds和Filters。
	// <li>net-detect-id - String - （过滤条件）网络探测实例ID，形如：netd-12345678。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNetDetectStatesRequest struct {
	*tchttp.BaseRequest
	
	// 网络探测实例`ID`数组。形如：[`netd-12345678`]。
	NetDetectIds []*string `json:"NetDetectIds,omitnil" name:"NetDetectIds"`

	// 过滤条件，参数不支持同时指定NetDetectIds和Filters。
	// <li>net-detect-id - String - （过滤条件）网络探测实例ID，形如：netd-12345678。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeNetDetectStatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectStatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetDetectStatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectStatesResponseParams struct {
	// 符合条件的网络探测验证结果对象数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetectStateSet []*NetDetectState `json:"NetDetectStateSet,omitnil" name:"NetDetectStateSet"`

	// 符合条件的网络探测验证结果对象数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNetDetectStatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetDetectStatesResponseParams `json:"Response"`
}

func (r *DescribeNetDetectStatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectsRequestParams struct {
	// 网络探测实例`ID`数组。形如：[`netd-12345678`]。
	NetDetectIds []*string `json:"NetDetectIds,omitnil" name:"NetDetectIds"`

	// 过滤条件，参数不支持同时指定NetDetectIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-12345678</li>
	// <li>net-detect-id - String - （过滤条件）网络探测实例ID，形如：netd-12345678</li>
	// <li>subnet-id - String - （过滤条件）子网实例ID，形如：subnet-12345678</li>
	// <li>net-detect-name - String - （过滤条件）网络探测名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNetDetectsRequest struct {
	*tchttp.BaseRequest
	
	// 网络探测实例`ID`数组。形如：[`netd-12345678`]。
	NetDetectIds []*string `json:"NetDetectIds,omitnil" name:"NetDetectIds"`

	// 过滤条件，参数不支持同时指定NetDetectIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-12345678</li>
	// <li>net-detect-id - String - （过滤条件）网络探测实例ID，形如：netd-12345678</li>
	// <li>subnet-id - String - （过滤条件）子网实例ID，形如：subnet-12345678</li>
	// <li>net-detect-name - String - （过滤条件）网络探测名称</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeNetDetectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetDetectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDetectsResponseParams struct {
	// 符合条件的网络探测对象数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetectSet []*NetDetect `json:"NetDetectSet,omitnil" name:"NetDetectSet"`

	// 符合条件的网络探测对象数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNetDetectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetDetectsResponseParams `json:"Response"`
}

func (r *DescribeNetDetectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDetectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAccountTypeRequestParams struct {

}

type DescribeNetworkAccountTypeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNetworkAccountTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAccountTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkAccountTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAccountTypeResponseParams struct {
	// 用户账号的网络类型，STANDARD为标准用户，LEGACY为传统用户
	NetworkAccountType *string `json:"NetworkAccountType,omitnil" name:"NetworkAccountType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNetworkAccountTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkAccountTypeResponseParams `json:"Response"`
}

func (r *DescribeNetworkAccountTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAccountTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclQuintupleEntriesRequestParams struct {
	// 网络ACL实例ID。形如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最小值为1，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，参数不支持同时指定`HaVipIds`和`Filters`。
	// <li>protocol - String - 协议，形如：`TCP`。</li>
	// <li>description - String - 描述。</li>
	// <li>destination-cidr - String - 目的CIDR， 形如：'192.168.0.0/24'。</li>
	// <li>source-cidr- String - 源CIDR， 形如：'192.168.0.0/24'。</li>
	// <li>action - String - 动作，形如ACCEPT或DROP。</li>
	// <li>network-acl-quintuple-entry-id - String - 五元组唯一ID，形如：'acli45-ahnu4rv5'。</li>
	// <li>network-acl-direction - String - 方向，形如：'INGRESS'或'EGRESS'。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。形如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最小值为1，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，参数不支持同时指定`HaVipIds`和`Filters`。
	// <li>protocol - String - 协议，形如：`TCP`。</li>
	// <li>description - String - 描述。</li>
	// <li>destination-cidr - String - 目的CIDR， 形如：'192.168.0.0/24'。</li>
	// <li>source-cidr- String - 源CIDR， 形如：'192.168.0.0/24'。</li>
	// <li>action - String - 动作，形如ACCEPT或DROP。</li>
	// <li>network-acl-quintuple-entry-id - String - 五元组唯一ID，形如：'acli45-ahnu4rv5'。</li>
	// <li>network-acl-direction - String - 方向，形如：'INGRESS'或'EGRESS'。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclQuintupleEntriesResponseParams struct {
	// 网络ACL条目列表（NetworkAclTuple5Entry）
	NetworkAclQuintupleSet []*NetworkAclQuintupleEntry `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *DescribeNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclsRequestParams struct {
	// 过滤条件，参数不支持同时指定NetworkAclIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-12345678。</li>
	// <li>network-acl-id - String - （过滤条件）网络ACL实例ID，形如：acl-12345678。</li>
	// <li>network-acl-name - String - （过滤条件）网络ACL实例名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 网络ACL实例ID数组。形如：[acl-12345678]。每次请求的实例的上限为100。参数不支持同时指定NetworkAclIds和Filters。
	NetworkAclIds []*string `json:"NetworkAclIds,omitnil" name:"NetworkAclIds"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最小值为1，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持：NetworkAclId,NetworkAclName,CreatedTime
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。顺序：ASC，倒序：DESC。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

type DescribeNetworkAclsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，参数不支持同时指定NetworkAclIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-12345678。</li>
	// <li>network-acl-id - String - （过滤条件）网络ACL实例ID，形如：acl-12345678。</li>
	// <li>network-acl-name - String - （过滤条件）网络ACL实例名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 网络ACL实例ID数组。形如：[acl-12345678]。每次请求的实例的上限为100。参数不支持同时指定NetworkAclIds和Filters。
	NetworkAclIds []*string `json:"NetworkAclIds,omitnil" name:"NetworkAclIds"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最小值为1，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持：NetworkAclId,NetworkAclName,CreatedTime
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。顺序：ASC，倒序：DESC。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

func (r *DescribeNetworkAclsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "NetworkAclIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkAclsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkAclsResponseParams struct {
	// 实例详细信息列表。
	NetworkAclSet []*NetworkAcl `json:"NetworkAclSet,omitnil" name:"NetworkAclSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNetworkAclsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkAclsResponseParams `json:"Response"`
}

func (r *DescribeNetworkAclsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkAclsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfaceLimitRequestParams struct {
	// 要查询的CVM实例ID或弹性网卡ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeNetworkInterfaceLimitRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的CVM实例ID或弹性网卡ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeNetworkInterfaceLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfaceLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkInterfaceLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfaceLimitResponseParams struct {
	// 标准型弹性网卡配额。
	EniQuantity *int64 `json:"EniQuantity,omitnil" name:"EniQuantity"`

	// 每个标准型弹性网卡可以分配的IP配额。
	EniPrivateIpAddressQuantity *int64 `json:"EniPrivateIpAddressQuantity,omitnil" name:"EniPrivateIpAddressQuantity"`

	// 扩展型网卡配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendEniQuantity *int64 `json:"ExtendEniQuantity,omitnil" name:"ExtendEniQuantity"`

	// 每个扩展型弹性网卡可以分配的IP配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendEniPrivateIpAddressQuantity *int64 `json:"ExtendEniPrivateIpAddressQuantity,omitnil" name:"ExtendEniPrivateIpAddressQuantity"`

	// 中继网卡配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubEniQuantity *int64 `json:"SubEniQuantity,omitnil" name:"SubEniQuantity"`

	// 每个中继网卡可以分配的IP配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubEniPrivateIpAddressQuantity *int64 `json:"SubEniPrivateIpAddressQuantity,omitnil" name:"SubEniPrivateIpAddressQuantity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNetworkInterfaceLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkInterfaceLimitResponseParams `json:"Response"`
}

func (r *DescribeNetworkInterfaceLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfaceLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesRequestParams struct {
	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。</li>
	// <li>network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。</li>
	// <li>attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ins-3nqpdn3i。</li>
	// <li>groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。</li>
	// <li>network-interface-name - String - （过滤条件）网卡实例名称。</li>
	// <li>network-interface-description - String - （过滤条件）网卡实例描述。</li>
	// <li>address-ip - String - （过滤条件）内网IPv4地址，单IP后缀模糊匹配，多IP精确匹配。可以与`ip-exact-match`配合做单IP的精确匹配查询。</li>
	// <li>ip-exact-match - Boolean - （过滤条件）内网IPv4精确匹配查询，存在多值情况，只取第一个。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。使用请参考示例2</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>
	// <li>is-primary - Boolean - 是否必填：否 - （过滤条件）按照是否主网卡进行过滤。值为true时，仅过滤主网卡；值为false时，仅过滤辅助网卡；此过滤参数未提供时，同时过滤主网卡和辅助网卡。</li>
	// <li>eni-type - String -是否必填：否- （过滤条件）按照网卡类型进行过滤。“0”-辅助网卡，“1”-主网卡，“2”：中继网卡。</li>
	// <li>eni-qos - String -是否必填：否- （过滤条件）按照网卡服务质量进行过滤。“AG”-服务质量为云铜，“AU”-服务质量为云银。</li>
	// <li>address-ipv6 - String - 是否必填：否 -（过滤条件）内网IPv6地址过滤，支持多ipv6地址查询，如果和address-ip一起使用取交集。</li>
	// <li>public-address-ip - String - （过滤条件）公网IPv4地址，精确匹配。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNetworkInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID查询。形如：eni-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定NetworkInterfaceIds和Filters。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 过滤条件，参数不支持同时指定NetworkInterfaceIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>subnet-id - String - （过滤条件）所属子网实例ID，形如：subnet-f49l6u0z。</li>
	// <li>network-interface-id - String - （过滤条件）弹性网卡实例ID，形如：eni-5k56k7k7。</li>
	// <li>attachment.instance-id - String - （过滤条件）绑定的云服务器实例ID，形如：ins-3nqpdn3i。</li>
	// <li>groups.security-group-id - String - （过滤条件）绑定的安全组实例ID，例如：sg-f9ekbxeq。</li>
	// <li>network-interface-name - String - （过滤条件）网卡实例名称。</li>
	// <li>network-interface-description - String - （过滤条件）网卡实例描述。</li>
	// <li>address-ip - String - （过滤条件）内网IPv4地址，单IP后缀模糊匹配，多IP精确匹配。可以与`ip-exact-match`配合做单IP的精确匹配查询。</li>
	// <li>ip-exact-match - Boolean - （过滤条件）内网IPv4精确匹配查询，存在多值情况，只取第一个。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。使用请参考示例2</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>
	// <li>is-primary - Boolean - 是否必填：否 - （过滤条件）按照是否主网卡进行过滤。值为true时，仅过滤主网卡；值为false时，仅过滤辅助网卡；此过滤参数未提供时，同时过滤主网卡和辅助网卡。</li>
	// <li>eni-type - String -是否必填：否- （过滤条件）按照网卡类型进行过滤。“0”-辅助网卡，“1”-主网卡，“2”：中继网卡。</li>
	// <li>eni-qos - String -是否必填：否- （过滤条件）按照网卡服务质量进行过滤。“AG”-服务质量为云铜，“AU”-服务质量为云银。</li>
	// <li>address-ipv6 - String - 是否必填：否 -（过滤条件）内网IPv6地址过滤，支持多ipv6地址查询，如果和address-ip一起使用取交集。</li>
	// <li>public-address-ip - String - （过滤条件）公网IPv4地址，精确匹配。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeNetworkInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInterfacesResponseParams struct {
	// 实例详细信息列表。
	NetworkInterfaceSet []*NetworkInterface `json:"NetworkInterfaceSet,omitnil" name:"NetworkInterfaceSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNetworkInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkInterfacesResponseParams `json:"Response"`
}

func (r *DescribeNetworkInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductQuotaRequestParams struct {
	// 查询的网络产品名称，可查询的产品有：vpc、ccn、vpn、dc、dfw、clb、eip。
	Product *string `json:"Product,omitnil" name:"Product"`
}

type DescribeProductQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 查询的网络产品名称，可查询的产品有：vpc、ccn、vpn、dc、dfw、clb、eip。
	Product *string `json:"Product,omitnil" name:"Product"`
}

func (r *DescribeProductQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductQuotaResponseParams struct {
	// ProductQuota对象数组。
	ProductQuotaSet []*ProductQuota `json:"ProductQuotaSet,omitnil" name:"ProductQuotaSet"`

	// 符合条件的产品类型个数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeProductQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductQuotaResponseParams `json:"Response"`
}

func (r *DescribeProductQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteConflictsRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 要检查的与之冲突的目的端列表。
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitnil" name:"DestinationCidrBlocks"`
}

type DescribeRouteConflictsRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 要检查的与之冲突的目的端列表。
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitnil" name:"DestinationCidrBlocks"`
}

func (r *DescribeRouteConflictsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteConflictsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "DestinationCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteConflictsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteConflictsResponseParams struct {
	// 路由策略冲突列表。
	RouteConflictSet []*RouteConflict `json:"RouteConflictSet,omitnil" name:"RouteConflictSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRouteConflictsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteConflictsResponseParams `json:"Response"`
}

func (r *DescribeRouteConflictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteConflictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesRequestParams struct {
	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// <li>route-table-id - String - （过滤条件）路由表实例ID。</li>
	// <li>route-table-name - String - （过滤条件）路由表名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>association.main - String - （过滤条件）是否主路由表。</li>
	// <li>tag-key - String -是否必填：否 - （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>
	// <li>next-hop-type - String - 是否必填：否 - （过滤条件）按下一跳类型进行过滤。使用next-hop-type进行过滤时，必须同时携带route-table-id与vpc-id。
	// 目前我们支持的类型有：
	// LOCAL: 本地路由
	// CVM：公网网关类型的云服务器；
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// HAVIP：高可用虚拟IP；
	// NAT：NAT网关; 
	// NORMAL_CVM：普通云服务器；
	// EIP：云服务器的公网IP；
	// CCN：云联网；
	// LOCAL_GATEWAY：本地网关。
	// </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds,omitnil" name:"RouteTableIds"`

	// 偏移量。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

type DescribeRouteTablesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// <li>route-table-id - String - （过滤条件）路由表实例ID。</li>
	// <li>route-table-name - String - （过滤条件）路由表名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>association.main - String - （过滤条件）是否主路由表。</li>
	// <li>tag-key - String -是否必填：否 - （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>
	// <li>next-hop-type - String - 是否必填：否 - （过滤条件）按下一跳类型进行过滤。使用next-hop-type进行过滤时，必须同时携带route-table-id与vpc-id。
	// 目前我们支持的类型有：
	// LOCAL: 本地路由
	// CVM：公网网关类型的云服务器；
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// HAVIP：高可用虚拟IP；
	// NAT：NAT网关; 
	// NORMAL_CVM：普通云服务器；
	// EIP：云服务器的公网IP；
	// CCN：云联网；
	// LOCAL_GATEWAY：本地网关。
	// </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableIds []*string `json:"RouteTableIds,omitnil" name:"RouteTableIds"`

	// 偏移量。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeRouteTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "RouteTableIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRouteTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRouteTablesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 路由表对象。
	RouteTableSet []*RouteTable `json:"RouteTableSet,omitnil" name:"RouteTableSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRouteTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRouteTablesResponseParams `json:"Response"`
}

func (r *DescribeRouteTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupAssociationStatisticsRequestParams struct {
	// 安全实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 安全实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *DescribeSecurityGroupAssociationStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupAssociationStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupAssociationStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupAssociationStatisticsResponseParams struct {
	// 安全组关联实例统计。
	SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet,omitnil" name:"SecurityGroupAssociationStatisticsSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityGroupAssociationStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupAssociationStatisticsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupAssociationStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupAssociationStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupLimitsRequestParams struct {

}

type DescribeSecurityGroupLimitsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecurityGroupLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupLimitsResponseParams struct {
	// 用户安全组配额限制。
	SecurityGroupLimitSet *SecurityGroupLimitSet `json:"SecurityGroupLimitSet,omitnil" name:"SecurityGroupLimitSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityGroupLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupLimitsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 过滤条件。
	// <li>security-group-id - String - 规则中的安全组ID。</li>
	// <li>ip - String - IP，支持IPV4和IPV6模糊匹配。</li>
	// <li>address-module - String - IP地址模板或IP地址组模板ID。</li>
	// <li>service-module - String - 协议端口模板或协议端口组模板ID。</li>
	// <li>protocol-type - String - 安全组策略支持的协议，可选值：`TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, `ALL`。</li>
	// <li>port - String - 是否必填：否 -协议端口，支持模糊匹配，值为`ALL`时，查询所有的端口。</li>
	// <li>poly - String - 协议策略，可选值：`ALL`，所有策略；`ACCEPT`，允许；`DROP`，拒绝。</li>
	// <li>direction - String - 协议规则，可选值：`ALL`，所有策略；`INBOUND`，入站规则；`OUTBOUND`，出站规则。</li>
	// <li>description - String - 协议描述，该过滤条件支持模糊匹配。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 过滤条件。
	// <li>security-group-id - String - 规则中的安全组ID。</li>
	// <li>ip - String - IP，支持IPV4和IPV6模糊匹配。</li>
	// <li>address-module - String - IP地址模板或IP地址组模板ID。</li>
	// <li>service-module - String - 协议端口模板或协议端口组模板ID。</li>
	// <li>protocol-type - String - 安全组策略支持的协议，可选值：`TCP`, `UDP`, `ICMP`, `ICMPV6`, `GRE`, `ALL`。</li>
	// <li>port - String - 是否必填：否 -协议端口，支持模糊匹配，值为`ALL`时，查询所有的端口。</li>
	// <li>poly - String - 协议策略，可选值：`ALL`，所有策略；`ACCEPT`，允许；`DROP`，拒绝。</li>
	// <li>direction - String - 协议规则，可选值：`ALL`，所有策略；`INBOUND`，入站规则；`OUTBOUND`，出站规则。</li>
	// <li>description - String - 协议描述，该过滤条件支持模糊匹配。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupPoliciesResponseParams struct {
	// 安全组规则集合。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupReferencesRequestParams struct {
	// 安全组实例ID数组。格式如：['sg-12345678']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type DescribeSecurityGroupReferencesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID数组。格式如：['sg-12345678']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *DescribeSecurityGroupReferencesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupReferencesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupReferencesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupReferencesResponseParams struct {
	// 安全组被引用信息。
	ReferredSecurityGroupSet []*ReferredSecurityGroup `json:"ReferredSecurityGroupSet,omitnil" name:"ReferredSecurityGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityGroupReferencesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupReferencesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupReferencesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupReferencesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsRequestParams struct {
	// 安全组实例ID，例如：sg-33ocnj9n。每次请求的实例的上限为100。参数不支持同时指定SecurityGroupIds和Filters。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 过滤条件，参数不支持同时指定SecurityGroupIds和Filters。
	// <li>security-group-id - String - （过滤条件）安全组ID。</li>
	// <li>project-id - Integer - （过滤条件）项目ID。</li>
	// <li>security-group-name - String - （过滤条件）安全组名称。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 其中 tag-key 请使用具体的标签键进行替换。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持：`CreatedTime` `UpdateTime`。注意：该字段没有默认值。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

type DescribeSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如：sg-33ocnj9n。每次请求的实例的上限为100。参数不支持同时指定SecurityGroupIds和Filters。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 过滤条件，参数不支持同时指定SecurityGroupIds和Filters。
	// <li>security-group-id - String - （过滤条件）安全组ID。</li>
	// <li>project-id - Integer - （过滤条件）项目ID。</li>
	// <li>security-group-name - String - （过滤条件）安全组名称。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 其中 tag-key 请使用具体的标签键进行替换。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 排序字段。支持：`CreatedTime` `UpdateTime`。注意：该字段没有默认值。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方法。升序：`ASC`，倒序：`DESC`。默认值：`ASC`
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

func (r *DescribeSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupsResponseParams struct {
	// 安全组对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupSet []*SecurityGroup `json:"SecurityGroupSet,omitnil" name:"SecurityGroupSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupsResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceTemplateGroupsRequestParams struct {
	// 过滤条件。
	// <li>service-template-group-name - String - （过滤条件）协议端口模板集合名称。</li>
	// <li>service-template-group-id - String - （过滤条件）协议端口模板集合实例ID，例如：ppmg-e6dy460g。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否获取协议端口模板成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

type DescribeServiceTemplateGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>service-template-group-name - String - （过滤条件）协议端口模板集合名称。</li>
	// <li>service-template-group-id - String - （过滤条件）协议端口模板集合实例ID，例如：ppmg-e6dy460g。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否获取协议端口模板成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

func (r *DescribeServiceTemplateGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplateGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NeedMemberInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceTemplateGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceTemplateGroupsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 协议端口模板集合。
	ServiceTemplateGroupSet []*ServiceTemplateGroup `json:"ServiceTemplateGroupSet,omitnil" name:"ServiceTemplateGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeServiceTemplateGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceTemplateGroupsResponseParams `json:"Response"`
}

func (r *DescribeServiceTemplateGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplateGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceTemplatesRequestParams struct {
	// 过滤条件。
	// <li>service-template-name - 协议端口模板名称。</li>
	// <li>service-template-id - 协议端口模板实例ID，例如：ppm-e6dy460g。</li>
	// <li>service-port- 协议端口。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否获取协议端口成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

type DescribeServiceTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li>service-template-name - 协议端口模板名称。</li>
	// <li>service-template-id - 协议端口模板实例ID，例如：ppm-e6dy460g。</li>
	// <li>service-port- 协议端口。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// 是否获取协议端口成员标识。
	NeedMemberInfo *bool `json:"NeedMemberInfo,omitnil" name:"NeedMemberInfo"`
}

func (r *DescribeServiceTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NeedMemberInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceTemplatesResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 协议端口模板对象。
	ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet,omitnil" name:"ServiceTemplateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeServiceTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceTemplatesResponseParams `json:"Response"`
}

func (r *DescribeServiceTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSgSnapshotFileContentRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 快照文件Id。
	SnapshotFileId *string `json:"SnapshotFileId,omitnil" name:"SnapshotFileId"`

	// 安全组Id。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

type DescribeSgSnapshotFileContentRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 快照文件Id。
	SnapshotFileId *string `json:"SnapshotFileId,omitnil" name:"SnapshotFileId"`

	// 安全组Id。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`
}

func (r *DescribeSgSnapshotFileContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSgSnapshotFileContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "SnapshotFileId")
	delete(f, "SecurityGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSgSnapshotFileContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSgSnapshotFileContentResponseParams struct {
	// 实例Id，即安全组Id。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 快照文件Id。
	SnapshotFileId *string `json:"SnapshotFileId,omitnil" name:"SnapshotFileId"`

	// 备份时间。
	BackupTime *string `json:"BackupTime,omitnil" name:"BackupTime"`

	// 操作者。
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 原始数据。
	OriginalData []*SecurityGroupPolicy `json:"OriginalData,omitnil" name:"OriginalData"`

	// 备份数据。
	BackupData []*SecurityGroupPolicy `json:"BackupData,omitnil" name:"BackupData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSgSnapshotFileContentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSgSnapshotFileContentResponseParams `json:"Response"`
}

func (r *DescribeSgSnapshotFileContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSgSnapshotFileContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotAttachedInstancesRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 过滤条件。
	// 支持的过滤条件如下：
	// <li>instance-id：实例ID。</li>
	// <li>instance-region：实例所在地域。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大为200。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSnapshotAttachedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 过滤条件。
	// 支持的过滤条件如下：
	// <li>instance-id：实例ID。</li>
	// <li>instance-region：实例所在地域。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大为200。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSnapshotAttachedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotAttachedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotAttachedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotAttachedInstancesResponseParams struct {
	// 实例列表
	InstanceSet []*SnapshotInstance `json:"InstanceSet,omitnil" name:"InstanceSet"`

	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSnapshotAttachedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotAttachedInstancesResponseParams `json:"Response"`
}

func (r *DescribeSnapshotAttachedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotAttachedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotFilesRequestParams struct {
	// 业务类型，目前支持安全组：securitygroup。
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 业务实例Id，和BusinessType对应。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 开始日期，格式%Y-%m-%d %H:%M:%S。
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束日期，格式%Y-%m-%d %H:%M:%S。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSnapshotFilesRequest struct {
	*tchttp.BaseRequest
	
	// 业务类型，目前支持安全组：securitygroup。
	BusinessType *string `json:"BusinessType,omitnil" name:"BusinessType"`

	// 业务实例Id，和BusinessType对应。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 开始日期，格式%Y-%m-%d %H:%M:%S。
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束日期，格式%Y-%m-%d %H:%M:%S。
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSnapshotFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BusinessType")
	delete(f, "InstanceId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotFilesResponseParams struct {
	// 快照文件集合。
	SnapshotFileSet []*SnapshotFileInfo `json:"SnapshotFileSet,omitnil" name:"SnapshotFileSet"`

	// 符合条件的对象数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSnapshotFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotFilesResponseParams `json:"Response"`
}

func (r *DescribeSnapshotFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotPoliciesRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`

	// 过滤条件，参数不支持同时指定SnapshotPolicyIds和Filters。
	// <li>snapshot-policy-id - String -（过滤条件）快照策略ID。</li>
	// <li>snapshot-policy-name - String -（过滤条件）快照策略名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大为200。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`

	// 过滤条件，参数不支持同时指定SnapshotPolicyIds和Filters。
	// <li>snapshot-policy-id - String -（过滤条件）快照策略ID。</li>
	// <li>snapshot-policy-name - String -（过滤条件）快照策略名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大为200。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSnapshotPoliciesResponseParams struct {
	// 快照策略。
	SnapshotPolicySet []*SnapshotPolicy `json:"SnapshotPolicySet,omitnil" name:"SnapshotPolicySet"`

	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecificTrafficPackageUsedDetailsRequestParams struct {
	// 共享流量包唯一ID
	TrafficPackageId *string `json:"TrafficPackageId,omitnil" name:"TrafficPackageId"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。详细的过滤条件如下：<li> resource-id - String - 是否必填：否 - （过滤条件）按照抵扣流量资源的唯一 ID 过滤。</li><li> resource-type - String - 是否必填：否 - （过滤条件）按照资源类型过滤，资源类型包括 CVM 和 EIP </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序条件。该参数仅支持根据抵扣量排序，传值为 deduction
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序类型，仅支持0和1，0-降序，1-升序。不传默认为0
	OrderType *int64 `json:"OrderType,omitnil" name:"OrderType"`

	// 开始时间。不传默认为当前时间往前推30天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。不传默认为当前时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSpecificTrafficPackageUsedDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 共享流量包唯一ID
	TrafficPackageId *string `json:"TrafficPackageId,omitnil" name:"TrafficPackageId"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。详细的过滤条件如下：<li> resource-id - String - 是否必填：否 - （过滤条件）按照抵扣流量资源的唯一 ID 过滤。</li><li> resource-type - String - 是否必填：否 - （过滤条件）按照资源类型过滤，资源类型包括 CVM 和 EIP </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序条件。该参数仅支持根据抵扣量排序，传值为 deduction
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序类型，仅支持0和1，0-降序，1-升序。不传默认为0
	OrderType *int64 `json:"OrderType,omitnil" name:"OrderType"`

	// 开始时间。不传默认为当前时间往前推30天
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。不传默认为当前时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSpecificTrafficPackageUsedDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecificTrafficPackageUsedDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficPackageId")
	delete(f, "Filters")
	delete(f, "OrderField")
	delete(f, "OrderType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpecificTrafficPackageUsedDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpecificTrafficPackageUsedDetailsResponseParams struct {
	// 符合查询条件的共享流量包用量明细的总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 共享流量包用量明细列表
	UsedDetailSet []*UsedDetail `json:"UsedDetailSet,omitnil" name:"UsedDetailSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpecificTrafficPackageUsedDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpecificTrafficPackageUsedDetailsResponseParams `json:"Response"`
}

func (r *DescribeSpecificTrafficPackageUsedDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpecificTrafficPackageUsedDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetResourceDashboardRequestParams struct {
	// Subnet实例ID，例如：subnet-f1xjkw1b。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`
}

type DescribeSubnetResourceDashboardRequest struct {
	*tchttp.BaseRequest
	
	// Subnet实例ID，例如：subnet-f1xjkw1b。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`
}

func (r *DescribeSubnetResourceDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetResourceDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetResourceDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetResourceDashboardResponseParams struct {
	// 资源统计结果。
	ResourceStatisticsSet []*ResourceStatistics `json:"ResourceStatisticsSet,omitnil" name:"ResourceStatisticsSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSubnetResourceDashboardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetResourceDashboardResponseParams `json:"Response"`
}

func (r *DescribeSubnetResourceDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetResourceDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsRequestParams struct {
	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// <li>subnet-id - String - （过滤条件）Subnet实例名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>cidr-block - String - （过滤条件）子网网段，形如: 192.168.1.0 。</li>
	// <li>is-default - Boolean - （过滤条件）是否是默认子网。</li>
	// <li>is-remote-vpc-snat - Boolean - （过滤条件）是否为VPC SNAT地址池子网。</li>
	// <li>subnet-name - String - （过滤条件）子网名称。</li>
	// <li>zone - String - （过滤条件）可用区。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>
	// <li>cdc-id - String - 是否必填：否 - （过滤条件）按照cdc信息进行过滤。过滤出来制定cdc下的子网。</li>
	// <li>is-cdc-subnet - String - 是否必填：否 - （过滤条件）按照是否是cdc子网进行过滤。取值：“0”-非cdc子网，“1”--cdc子网</li>
	// <li>ipv6-cidr-block - String - （过滤条件）IPv6子网网段，形如: 2402:4e00:1717:8700::/64 。</li>
	// <li>isp-type  - String - （过滤条件）运营商类型，形如: BGP 。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID查询。形如：subnet-pxir56ns。每次请求的实例的上限为100。参数不支持同时指定SubnetIds和Filters。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`

	// 过滤条件，参数不支持同时指定SubnetIds和Filters。
	// <li>subnet-id - String - （过滤条件）Subnet实例名称。</li>
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>cidr-block - String - （过滤条件）子网网段，形如: 192.168.1.0 。</li>
	// <li>is-default - Boolean - （过滤条件）是否是默认子网。</li>
	// <li>is-remote-vpc-snat - Boolean - （过滤条件）是否为VPC SNAT地址池子网。</li>
	// <li>subnet-name - String - （过滤条件）子网名称。</li>
	// <li>zone - String - （过滤条件）可用区。</li>
	// <li>tag-key - String -是否必填：否- （过滤条件）按照标签键进行过滤。</li>
	// <li>tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。</li>
	// <li>cdc-id - String - 是否必填：否 - （过滤条件）按照cdc信息进行过滤。过滤出来制定cdc下的子网。</li>
	// <li>is-cdc-subnet - String - 是否必填：否 - （过滤条件）按照是否是cdc子网进行过滤。取值：“0”-非cdc子网，“1”--cdc子网</li>
	// <li>ipv6-cidr-block - String - （过滤条件）IPv6子网网段，形如: 2402:4e00:1717:8700::/64 。</li>
	// <li>isp-type  - String - （过滤条件）运营商类型，形如: BGP 。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubnetsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 子网对象。
	SubnetSet []*Subnet `json:"SubnetSet,omitnil" name:"SubnetSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubnetsResponseParams `json:"Response"`
}

func (r *DescribeSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultRequestParams struct {
	// 异步任务ID。TaskId和DealName必填一个参数
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 计费订单号。TaskId和DealName必填一个参数
	DealName *string `json:"DealName,omitnil" name:"DealName"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 异步任务ID。TaskId和DealName必填一个参数
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 计费订单号。TaskId和DealName必填一个参数
	DealName *string `json:"DealName,omitnil" name:"DealName"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "DealName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// 任务ID
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 执行结果，包括"SUCCESS", "FAILED", "RUNNING"
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResultResponseParams `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateLimitsRequestParams struct {

}

type DescribeTemplateLimitsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTemplateLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateLimitsResponseParams struct {
	// 参数模板配额对象。
	TemplateLimit *TemplateLimit `json:"TemplateLimit,omitnil" name:"TemplateLimit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTemplateLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateLimitsResponseParams `json:"Response"`
}

func (r *DescribeTemplateLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantCcnsRequestParams struct {
	// 过滤条件，目前`value`值个数只支持一个，允许可支持的字段有：
	// <li>`ccn-ids` 云联网ID数组，值形如：`["ccn-12345678"]`</li> 
	// <li>`user-account-id` 用户账号ID，值形如：`["12345678"]`</li><li>`is-security-lock` 是否锁定，值形如：`["true"]`</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数据量，可选值0到100之间的整数，默认20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeTenantCcnsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，目前`value`值个数只支持一个，允许可支持的字段有：
	// <li>`ccn-ids` 云联网ID数组，值形如：`["ccn-12345678"]`</li> 
	// <li>`user-account-id` 用户账号ID，值形如：`["12345678"]`</li><li>`is-security-lock` 是否锁定，值形如：`["true"]`</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数据量，可选值0到100之间的整数，默认20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeTenantCcnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantCcnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTenantCcnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTenantCcnsResponseParams struct {
	// 云联网（CCN）对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnSet []*CcnInstanceInfo `json:"CcnSet,omitnil" name:"CcnSet"`

	// 符合条件的对象总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTenantCcnsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTenantCcnsResponseParams `json:"Response"`
}

func (r *DescribeTenantCcnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTenantCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficPackagesRequestParams struct {
	// 共享流量包ID，支持批量
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil" name:"TrafficPackageIds"`

	// 每次请求的`Filters`的上限为10。参数不支持同时指定`TrafficPackageIds`和`Filters`。详细的过滤条件如下：
	// <li> traffic-package_id - String - 是否必填：否 - （过滤条件）按照共享流量包的唯一标识ID过滤。</li>
	// <li> traffic-package-name - String - 是否必填：否 - （过滤条件）按照共享流量包名称过滤。不支持模糊过滤。</li>
	// <li> status - String - 是否必填：否 - （过滤条件）按照共享流量包状态过滤。可选状态：[AVAILABLE|EXPIRED|EXHAUSTED]</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeTrafficPackagesRequest struct {
	*tchttp.BaseRequest
	
	// 共享流量包ID，支持批量
	TrafficPackageIds []*string `json:"TrafficPackageIds,omitnil" name:"TrafficPackageIds"`

	// 每次请求的`Filters`的上限为10。参数不支持同时指定`TrafficPackageIds`和`Filters`。详细的过滤条件如下：
	// <li> traffic-package_id - String - 是否必填：否 - （过滤条件）按照共享流量包的唯一标识ID过滤。</li>
	// <li> traffic-package-name - String - 是否必填：否 - （过滤条件）按照共享流量包名称过滤。不支持模糊过滤。</li>
	// <li> status - String - 是否必填：否 - （过滤条件）按照共享流量包状态过滤。可选状态：[AVAILABLE|EXPIRED|EXHAUSTED]</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页参数
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页参数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeTrafficPackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrafficPackageIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrafficPackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrafficPackagesResponseParams struct {
	// 按照条件查询出来的流量包数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 流量包信息
	TrafficPackageSet []*TrafficPackage `json:"TrafficPackageSet,omitnil" name:"TrafficPackageSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTrafficPackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrafficPackagesResponseParams `json:"Response"`
}

func (r *DescribeTrafficPackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrafficPackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsedIpAddressRequestParams struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 查询是否占用的ip列表，ip需要在vpc或子网内。最多允许一次查询100个IP。
	IpAddresses []*string `json:"IpAddresses,omitnil" name:"IpAddresses"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeUsedIpAddressRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 查询是否占用的ip列表，ip需要在vpc或子网内。最多允许一次查询100个IP。
	IpAddresses []*string `json:"IpAddresses,omitnil" name:"IpAddresses"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeUsedIpAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsedIpAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "IpAddresses")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsedIpAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsedIpAddressResponseParams struct {
	// 占用ip地址的资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddressStates []*IpAddressStates `json:"IpAddressStates,omitnil" name:"IpAddressStates"`

	// 返回占用资源的个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUsedIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsedIpAddressResponseParams `json:"Response"`
}

func (r *DescribeUsedIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsedIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointRequestParams struct {
	// 过滤条件。
	// <li> end-point-service-id- String - （过滤条件）终端节点服务ID。</li>
	// <li>end-point-name - String - （过滤条件）终端节点实例名称。</li>
	// <li> end-point-id- String - （过滤条件）终端节点实例ID。</li>
	// <li> vpc-id- String - （过滤条件）VPC实例ID。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 终端节点ID列表。
	EndPointId []*string `json:"EndPointId,omitnil" name:"EndPointId"`
}

type DescribeVpcEndPointRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。
	// <li> end-point-service-id- String - （过滤条件）终端节点服务ID。</li>
	// <li>end-point-name - String - （过滤条件）终端节点实例名称。</li>
	// <li> end-point-id- String - （过滤条件）终端节点实例ID。</li>
	// <li> vpc-id- String - （过滤条件）VPC实例ID。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 终端节点ID列表。
	EndPointId []*string `json:"EndPointId,omitnil" name:"EndPointId"`
}

func (r *DescribeVpcEndPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcEndPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointResponseParams struct {
	// 终端节点对象。
	EndPointSet []*EndPoint `json:"EndPointSet,omitnil" name:"EndPointSet"`

	// 符合查询条件的终端节点个数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcEndPointResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcEndPointResponseParams `json:"Response"`
}

func (r *DescribeVpcEndPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceRequestParams struct {
	// 过滤条件。不支持同时传入参数 EndPointServiceIds and Filters。
	// <li> service-id - String - （过滤条件）终端节点服务唯一ID。</li>
	// <li>service-name - String - （过滤条件）终端节点实例名称。</li>
	// <li>service-instance-id - String - （过滤条件）后端服务的唯一ID，比如lb-xxx。</li>
	// <li>service-type - String - （过滤条件）后端PAAS服务类型，CLB,CDB,CRS，不填默认查询类型为CLB。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 终端节点服务ID。不支持同时传入参数 EndPointServiceIds and Filters。
	EndPointServiceIds []*string `json:"EndPointServiceIds,omitnil" name:"EndPointServiceIds"`

	// <li>不支持同时传入参数 Filters 。</li> <li>列出授权给当前账号的的终端节点服务信息。可以配合EndPointServiceIds参数进行过滤，那些终端节点服务授权了该账户。</li>
	IsListAuthorizedEndPointService *bool `json:"IsListAuthorizedEndPointService,omitnil" name:"IsListAuthorizedEndPointService"`
}

type DescribeVpcEndPointServiceRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件。不支持同时传入参数 EndPointServiceIds and Filters。
	// <li> service-id - String - （过滤条件）终端节点服务唯一ID。</li>
	// <li>service-name - String - （过滤条件）终端节点实例名称。</li>
	// <li>service-instance-id - String - （过滤条件）后端服务的唯一ID，比如lb-xxx。</li>
	// <li>service-type - String - （过滤条件）后端PAAS服务类型，CLB,CDB,CRS，不填默认查询类型为CLB。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 终端节点服务ID。不支持同时传入参数 EndPointServiceIds and Filters。
	EndPointServiceIds []*string `json:"EndPointServiceIds,omitnil" name:"EndPointServiceIds"`

	// <li>不支持同时传入参数 Filters 。</li> <li>列出授权给当前账号的的终端节点服务信息。可以配合EndPointServiceIds参数进行过滤，那些终端节点服务授权了该账户。</li>
	IsListAuthorizedEndPointService *bool `json:"IsListAuthorizedEndPointService,omitnil" name:"IsListAuthorizedEndPointService"`
}

func (r *DescribeVpcEndPointServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "EndPointServiceIds")
	delete(f, "IsListAuthorizedEndPointService")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcEndPointServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceResponseParams struct {
	// 终端节点服务对象数组。
	EndPointServiceSet []*EndPointService `json:"EndPointServiceSet,omitnil" name:"EndPointServiceSet"`

	// 符合查询条件的个数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcEndPointServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcEndPointServiceResponseParams `json:"Response"`
}

func (r *DescribeVpcEndPointServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceWhiteListRequestParams struct {
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li> user-uin String - （过滤条件）用户UIN。</li>
	// <li> end-point-service-id String - （过滤条件）终端节点服务ID。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 单页返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件。
	// <li> user-uin String - （过滤条件）用户UIN。</li>
	// <li> end-point-service-id String - （过滤条件）终端节点服务ID。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcEndPointServiceWhiteListResponseParams struct {
	// 白名单对象数组。已废弃
	VpcEndpointServiceUserSet []*VpcEndPointServiceUser `json:"VpcEndpointServiceUserSet,omitnil" name:"VpcEndpointServiceUserSet"`

	// 白名单对象数组。
	VpcEndPointServiceUserSet []*VpcEndPointServiceUser `json:"VpcEndPointServiceUserSet,omitnil" name:"VpcEndPointServiceUserSet"`

	// 符合条件的白名单个数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *DescribeVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcInstancesRequestParams struct {
	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>instance-id - String - （过滤条件）云主机实例ID。</li>
	// <li>instance-name - String - （过滤条件）云主机名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeVpcInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，参数不支持同时指定RouteTableIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>instance-id - String - （过滤条件）云主机实例ID。</li>
	// <li>instance-name - String - （过滤条件）云主机名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeVpcInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcInstancesResponseParams struct {
	// 云主机实例列表。
	InstanceSet []*CvmInstance `json:"InstanceSet,omitnil" name:"InstanceSet"`

	// 满足条件的云主机实例个数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcInstancesResponseParams `json:"Response"`
}

func (r *DescribeVpcInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcIpv6AddressesRequestParams struct {
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `IP`地址列表，批量查询单次请求最多支持`10`个。
	Ipv6Addresses []*string `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// VPC下的子网ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

type DescribeVpcIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `IP`地址列表，批量查询单次请求最多支持`10`个。
	Ipv6Addresses []*string `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// VPC下的子网ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

func (r *DescribeVpcIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6Addresses")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubnetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcIpv6AddressesResponseParams struct {
	// `IPv6`地址列表。
	Ipv6AddressSet []*VpcIpv6Address `json:"Ipv6AddressSet,omitnil" name:"Ipv6AddressSet"`

	// `IPv6`地址总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcIpv6AddressesResponseParams `json:"Response"`
}

func (r *DescribeVpcIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcLimitsRequestParams struct {
	// 配额名称。每次最大查询100个配额类型。
	LimitTypes []*string `json:"LimitTypes,omitnil" name:"LimitTypes"`
}

type DescribeVpcLimitsRequest struct {
	*tchttp.BaseRequest
	
	// 配额名称。每次最大查询100个配额类型。
	LimitTypes []*string `json:"LimitTypes,omitnil" name:"LimitTypes"`
}

func (r *DescribeVpcLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LimitTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcLimitsResponseParams struct {
	// 私有网络配额
	VpcLimitSet []*VpcLimit `json:"VpcLimitSet,omitnil" name:"VpcLimitSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcLimitsResponseParams `json:"Response"`
}

func (r *DescribeVpcLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPeeringConnectionsRequestParams struct {
	// 对等连接唯一ID数组。
	PeeringConnectionIds []*string `json:"PeeringConnectionIds,omitnil" name:"PeeringConnectionIds"`

	// 过滤条件，参数不支持同时指定PeeringConnectionIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>state String - （过滤条件）对等连接状态，可选值有：PENDING，投放中；ACTIVE，使用中；EXPIRED，已过期；REJECTED，拒绝。</li>
	// <li>peering-connection-name - String - （过滤条件）对等连接名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，可选值有：CreatedTime，PeeringConnectionName。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方式：DESC，降序；ASC，升序。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

type DescribeVpcPeeringConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 对等连接唯一ID数组。
	PeeringConnectionIds []*string `json:"PeeringConnectionIds,omitnil" name:"PeeringConnectionIds"`

	// 过滤条件，参数不支持同时指定PeeringConnectionIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>state String - （过滤条件）对等连接状态，可选值有：PENDING，投放中；ACTIVE，使用中；EXPIRED，已过期；REJECTED，拒绝。</li>
	// <li>peering-connection-name - String - （过滤条件）对等连接名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，可选值有：CreatedTime，PeeringConnectionName。
	OrderField *string `json:"OrderField,omitnil" name:"OrderField"`

	// 排序方式：DESC，降序；ASC，升序。
	OrderDirection *string `json:"OrderDirection,omitnil" name:"OrderDirection"`
}

func (r *DescribeVpcPeeringConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPeeringConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PeeringConnectionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderField")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcPeeringConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPeeringConnectionsResponseParams struct {
	// 满足条件的对等连接实例个数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 对等连接实例列表。
	PeerConnectionSet []*PeerConnection `json:"PeerConnectionSet,omitnil" name:"PeerConnectionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcPeeringConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcPeeringConnectionsResponseParams `json:"Response"`
}

func (r *DescribeVpcPeeringConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPeeringConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPrivateIpAddressesRequestParams struct {
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 内网`IP`地址列表，批量查询单次请求最多支持`10`个。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`
}

type DescribeVpcPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 内网`IP`地址列表，批量查询单次请求最多支持`10`个。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`
}

func (r *DescribeVpcPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "PrivateIpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcPrivateIpAddressesResponseParams struct {
	// 内网`IP`地址信息列表。
	VpcPrivateIpAddressSet []*VpcPrivateIpAddress `json:"VpcPrivateIpAddressSet,omitnil" name:"VpcPrivateIpAddressSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcPrivateIpAddressesResponseParams `json:"Response"`
}

func (r *DescribeVpcPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcResourceDashboardRequestParams struct {
	// Vpc实例ID，例如：vpc-f1xjkw1b。
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`
}

type DescribeVpcResourceDashboardRequest struct {
	*tchttp.BaseRequest
	
	// Vpc实例ID，例如：vpc-f1xjkw1b。
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`
}

func (r *DescribeVpcResourceDashboardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcResourceDashboardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcResourceDashboardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcResourceDashboardResponseParams struct {
	// 资源对象列表。
	ResourceDashboardSet []*ResourceDashboard `json:"ResourceDashboardSet,omitnil" name:"ResourceDashboardSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcResourceDashboardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcResourceDashboardResponseParams `json:"Response"`
}

func (r *DescribeVpcResourceDashboardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcResourceDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcTaskResultRequestParams struct {
	// 异步任务请求返回的RequestId。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

type DescribeVpcTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 异步任务请求返回的RequestId。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

func (r *DescribeVpcTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcTaskResultResponseParams struct {
	// 异步任务执行结果。结果：SUCCESS、FAILED、RUNNING。3者其中之一。其中SUCCESS表示任务执行成功，FAILED表示任务执行失败，RUNNING表示任务执行中。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 异步任务执行输出。
	Output *string `json:"Output,omitnil" name:"Output"`

	// 异步任务详细结果。只用于特殊场景，如批量删除弹性网卡时查询成功的网卡列表和失败的列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*VpcTaskResultDetailInfo `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcTaskResultResponseParams `json:"Response"`
}

func (r *DescribeVpcTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsRequestParams struct {
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`

	// 过滤条件，不支持同时指定VpcIds和Filters参数。
	// 支持的过滤条件如下：
	// <li>vpc-name：VPC实例名称，支持模糊查询。</li>
	// <li>is-default ：是否默认VPC。</li>
	// <li>vpc-id ：VPC实例ID，例如：vpc-f49l6u0z。</li>
	// <li>cidr-block：VPC的CIDR。</li>
	// <li>tag-key ：按照标签键进行过滤，非必填参数。</li>
	// <li>tag:tag-key：按照标签键值对进行过滤，非必填参数。 其中 tag-key 请使用具体的标签键进行替换，可参考示例2。</li>
	//   **说明：**若同一个过滤条件（Filter）存在多个Values，则同一Filter下Values间的关系为逻辑或（OR）关系；若存在多个过滤条件（Filter），Filter之间的关系为逻辑与（AND）关系。
	// <li>ipv6-cidr-block - String - （过滤条件）IPv6子网网段，形如: 2402:4e00:1717:8700::/64 。</li>
	// <li>isp-type  - String - （过滤条件）运营商类型，形如: BGP 取值范围：'BGP'-默认, 'CMCC'-中国移动, 'CTCC'-中国电信, 'CUCC'-中国联调。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

type DescribeVpcsRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`

	// 过滤条件，不支持同时指定VpcIds和Filters参数。
	// 支持的过滤条件如下：
	// <li>vpc-name：VPC实例名称，支持模糊查询。</li>
	// <li>is-default ：是否默认VPC。</li>
	// <li>vpc-id ：VPC实例ID，例如：vpc-f49l6u0z。</li>
	// <li>cidr-block：VPC的CIDR。</li>
	// <li>tag-key ：按照标签键进行过滤，非必填参数。</li>
	// <li>tag:tag-key：按照标签键值对进行过滤，非必填参数。 其中 tag-key 请使用具体的标签键进行替换，可参考示例2。</li>
	//   **说明：**若同一个过滤条件（Filter）存在多个Values，则同一Filter下Values间的关系为逻辑或（OR）关系；若存在多个过滤条件（Filter），Filter之间的关系为逻辑与（AND）关系。
	// <li>ipv6-cidr-block - String - （过滤条件）IPv6子网网段，形如: 2402:4e00:1717:8700::/64 。</li>
	// <li>isp-type  - String - （过滤条件）运营商类型，形如: BGP 取值范围：'BGP'-默认, 'CMCC'-中国移动, 'CTCC'-中国电信, 'CUCC'-中国联调。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。
	Offset *string `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *string `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcsResponseParams struct {
	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// VPC对象。
	VpcSet []*Vpc `json:"VpcSet,omitnil" name:"VpcSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpcsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcsResponseParams `json:"Response"`
}

func (r *DescribeVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnConnectionsRequestParams struct {
	// VPN通道实例ID。形如：vpnx-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnConnectionIds和Filters。
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitnil" name:"VpnConnectionIds"`

	// 过滤条件。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpnConnectionIds和Filters。
	// <li>vpc-id - String - VPC实例ID，形如：`vpc-0a36uwkr`。</li>
	// <li>vpn-gateway-id - String - VPN网关实例ID，形如：`vpngw-p4lmqawn`。</li>
	// <li>customer-gateway-id - String - 对端网关实例ID，形如：`cgw-l4rblw63`。</li>
	// <li>vpn-connection-name - String - 通道名称，形如：`test-vpn`。</li>
	// <li>vpn-connection-id - String - 通道实例ID，形如：`vpnx-5p7vkch8"`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeVpnConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// VPN通道实例ID。形如：vpnx-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnConnectionIds和Filters。
	VpnConnectionIds []*string `json:"VpnConnectionIds,omitnil" name:"VpnConnectionIds"`

	// 过滤条件。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定VpnConnectionIds和Filters。
	// <li>vpc-id - String - VPC实例ID，形如：`vpc-0a36uwkr`。</li>
	// <li>vpn-gateway-id - String - VPN网关实例ID，形如：`vpngw-p4lmqawn`。</li>
	// <li>customer-gateway-id - String - 对端网关实例ID，形如：`cgw-l4rblw63`。</li>
	// <li>vpn-connection-name - String - 通道名称，形如：`test-vpn`。</li>
	// <li>vpn-connection-id - String - 通道实例ID，形如：`vpnx-5p7vkch8"`。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeVpnConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnConnectionsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// VPN通道实例。
	VpnConnectionSet []*VpnConnection `json:"VpnConnectionSet,omitnil" name:"VpnConnectionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpnConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnConnectionsResponseParams `json:"Response"`
}

func (r *DescribeVpnConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayCcnRoutesRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 偏移量。默认值：0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。默认值：20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeVpnGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 偏移量。默认值：0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。默认值：20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeVpnGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayCcnRoutesResponseParams struct {
	// 云联网路由（IDC网段）列表。
	RouteSet []*VpngwCcnRoutes `json:"RouteSet,omitnil" name:"RouteSet"`

	// 符合条件的对象数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpnGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayRoutesRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 过滤条件,  条件包括(DestinationCidr, InstanceId,InstanceType)。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量, 默认0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 单页个数, 默认20, 最大值100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 过滤条件,  条件包括(DestinationCidr, InstanceId,InstanceType)。
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量, 默认0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 单页个数, 默认20, 最大值100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewayRoutesResponseParams struct {
	// VPN网关目的路由。
	Routes []*VpnGatewayRoute `json:"Routes,omitnil" name:"Routes"`

	// 路由条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaySslClientsRequestParams struct {
	// 过滤条件，参数不支持同时指定SslVpnClientIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>vpn-gateway-id - String - （过滤条件）VPN实例ID形如：vpngw-5aluhh9t。</li>
	// <li>ssl-vpn-server-id - String - （过滤条件）SSL-VPN-SERVER实例ID形如：vpns-1j2w6xpx。</li>
	// <li>ssl-vpn-client-id - String - （过滤条件）SSL-VPN-CLIENT实例ID形如：vpnc-3rlxp4nd。</li>
	// <li>ssl-vpn-client-name - String - （过滤条件）SSL-VPN-CLIENT实例名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认值0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数，默认值20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// SSL-VPN-CLIENT实例ID。形如：	
	// vpns-1jww3xpx。每次请求的实例的上限为100。参数不支持同时指定SslVpnClientIds和Filters。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`

	// VPN门户网站使用。默认是False。
	IsVpnPortal *bool `json:"IsVpnPortal,omitnil" name:"IsVpnPortal"`
}

type DescribeVpnGatewaySslClientsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，参数不支持同时指定SslVpnClientIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>vpn-gateway-id - String - （过滤条件）VPN实例ID形如：vpngw-5aluhh9t。</li>
	// <li>ssl-vpn-server-id - String - （过滤条件）SSL-VPN-SERVER实例ID形如：vpns-1j2w6xpx。</li>
	// <li>ssl-vpn-client-id - String - （过滤条件）SSL-VPN-CLIENT实例ID形如：vpnc-3rlxp4nd。</li>
	// <li>ssl-vpn-client-name - String - （过滤条件）SSL-VPN-CLIENT实例名称。</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 偏移量，默认值0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数，默认值20。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// SSL-VPN-CLIENT实例ID。形如：	
	// vpns-1jww3xpx。每次请求的实例的上限为100。参数不支持同时指定SslVpnClientIds和Filters。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`

	// VPN门户网站使用。默认是False。
	IsVpnPortal *bool `json:"IsVpnPortal,omitnil" name:"IsVpnPortal"`
}

func (r *DescribeVpnGatewaySslClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaySslClientsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SslVpnClientIds")
	delete(f, "IsVpnPortal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewaySslClientsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaySslClientsResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// SSL-VPN-CLIENT 实例列表。
	SslVpnClientSet []*SslVpnClient `json:"SslVpnClientSet,omitnil" name:"SslVpnClientSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpnGatewaySslClientsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewaySslClientsResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewaySslClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaySslClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaySslServersRequestParams struct {
	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// SSL-VPN-SERVER实例ID。形如：vpngwSslServer-12345678。每次请求的实例的上限为100。参数不支持同时指定SslVpnServerIds和Filters。
	SslVpnServerIds []*string `json:"SslVpnServerIds,omitnil" name:"SslVpnServerIds"`

	// 过滤条件，参数不支持同时指定SslVpnServerIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>vpn-gateway-id - String - （过滤条件）VPN实例ID，形如：vpngw-5aluhh9t。</li>
	// <li>vpn-gateway-name - String - （过滤条件）VPN实例名称。</li>
	// <li>ssl-vpn-server-name - String - （过滤条件）SSL-VPN-SERVER实例名称。</li>
	// <li>ssl-vpn-server-id - String - （过滤条件）SSL-VPN-SERVER实例ID，形如：vpns-xxx。</li>
	Filters []*FilterObject `json:"Filters,omitnil" name:"Filters"`

	// vpn门户使用。 默认Flase
	IsVpnPortal *bool `json:"IsVpnPortal,omitnil" name:"IsVpnPortal"`
}

type DescribeVpnGatewaySslServersRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// SSL-VPN-SERVER实例ID。形如：vpngwSslServer-12345678。每次请求的实例的上限为100。参数不支持同时指定SslVpnServerIds和Filters。
	SslVpnServerIds []*string `json:"SslVpnServerIds,omitnil" name:"SslVpnServerIds"`

	// 过滤条件，参数不支持同时指定SslVpnServerIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID，形如：vpc-f49l6u0z。</li>
	// <li>vpn-gateway-id - String - （过滤条件）VPN实例ID，形如：vpngw-5aluhh9t。</li>
	// <li>vpn-gateway-name - String - （过滤条件）VPN实例名称。</li>
	// <li>ssl-vpn-server-name - String - （过滤条件）SSL-VPN-SERVER实例名称。</li>
	// <li>ssl-vpn-server-id - String - （过滤条件）SSL-VPN-SERVER实例ID，形如：vpns-xxx。</li>
	Filters []*FilterObject `json:"Filters,omitnil" name:"Filters"`

	// vpn门户使用。 默认Flase
	IsVpnPortal *bool `json:"IsVpnPortal,omitnil" name:"IsVpnPortal"`
}

func (r *DescribeVpnGatewaySslServersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaySslServersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SslVpnServerIds")
	delete(f, "Filters")
	delete(f, "IsVpnPortal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewaySslServersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaySslServersResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// SSL-VPN-SERVER 实例详细信息列表。
	SslVpnSeverSet []*SslVpnSever `json:"SslVpnSeverSet,omitnil" name:"SslVpnSeverSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpnGatewaySslServersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewaySslServersResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewaySslServersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaySslServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaysRequestParams struct {
	// VPN网关实例ID。形如：vpngw-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnGatewayIds和Filters。
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil" name:"VpnGatewayIds"`

	// 过滤条件，参数不支持同时指定VpnGatewayIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>vpn-gateway-id - String - （过滤条件）VPN实例ID形如：vpngw-5aluhh9t。</li>
	// <li>vpn-gateway-name - String - （过滤条件）VPN实例名称。</li>
	// <li>type - String - （过滤条件）VPN网关类型：'IPSEC', 'SSL'。</li>
	// <li>public-ip-address- String - （过滤条件）公网IP。</li>
	// <li>renew-flag - String - （过滤条件）网关续费类型，手动续费：'NOTIFY_AND_MANUAL_RENEW'、自动续费：'NOTIFY_AND_AUTO_RENEW'。</li>
	// <li>zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。</li>
	Filters []*FilterObject `json:"Filters,omitnil" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeVpnGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。形如：vpngw-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpnGatewayIds和Filters。
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil" name:"VpnGatewayIds"`

	// 过滤条件，参数不支持同时指定VpnGatewayIds和Filters。
	// <li>vpc-id - String - （过滤条件）VPC实例ID形如：vpc-f49l6u0z。</li>
	// <li>vpn-gateway-id - String - （过滤条件）VPN实例ID形如：vpngw-5aluhh9t。</li>
	// <li>vpn-gateway-name - String - （过滤条件）VPN实例名称。</li>
	// <li>type - String - （过滤条件）VPN网关类型：'IPSEC', 'SSL'。</li>
	// <li>public-ip-address- String - （过滤条件）公网IP。</li>
	// <li>renew-flag - String - （过滤条件）网关续费类型，手动续费：'NOTIFY_AND_MANUAL_RENEW'、自动续费：'NOTIFY_AND_AUTO_RENEW'。</li>
	// <li>zone - String - （过滤条件）VPN所在可用区，形如：ap-guangzhou-2。</li>
	Filters []*FilterObject `json:"Filters,omitnil" name:"Filters"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 请求对象个数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeVpnGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpnGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpnGatewaysResponseParams struct {
	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// VPN网关实例详细信息列表。
	VpnGatewaySet []*VpnGateway `json:"VpnGatewaySet,omitnil" name:"VpnGatewaySet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVpnGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpnGatewaysResponseParams `json:"Response"`
}

func (r *DescribeVpnGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpnGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestinationIpPortTranslationNatRule struct {
	// 网络协议，可选值：`TCP`、`UDP`。
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 弹性IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil" name:"PublicIpAddress"`

	// 公网端口。
	PublicPort *uint64 `json:"PublicPort,omitnil" name:"PublicPort"`

	// 内网地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 内网端口。
	PrivatePort *uint64 `json:"PrivatePort,omitnil" name:"PrivatePort"`

	// NAT网关转发规则描述。
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type DetachCcnInstancesRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 要解关联网络实例列表
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

type DetachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 要解关联网络实例列表
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

func (r *DetachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachCcnInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachCcnInstancesResponseParams `json:"Response"`
}

func (r *DetachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachClassicLinkVpcRequestParams struct {
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CVM实例ID查询。形如：ins-r8hr2upy。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DetachClassicLinkVpcRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// CVM实例ID查询。形如：ins-r8hr2upy。
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *DetachClassicLinkVpcRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachClassicLinkVpcRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachClassicLinkVpcRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachClassicLinkVpcResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetachClassicLinkVpcResponse struct {
	*tchttp.BaseResponse
	Response *DetachClassicLinkVpcResponseParams `json:"Response"`
}

func (r *DetachClassicLinkVpcResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachClassicLinkVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// CVM实例ID。形如：ins-r8hr2upy。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type DetachNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// CVM实例ID。形如：ins-r8hr2upy。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。	
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *DetachNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "InstanceId")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetachNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *DetachNetworkInterfaceResponseParams `json:"Response"`
}

func (r *DetachNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachSnapshotInstancesRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 实例信息。
	Instances []*SnapshotInstance `json:"Instances,omitnil" name:"Instances"`
}

type DetachSnapshotInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 实例信息。
	Instances []*SnapshotInstance `json:"Instances,omitnil" name:"Instances"`
}

func (r *DetachSnapshotInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachSnapshotInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachSnapshotInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachSnapshotInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetachSnapshotInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachSnapshotInstancesResponseParams `json:"Response"`
}

func (r *DetachSnapshotInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachSnapshotInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DhcpIp struct {
	// `DhcpIp`的`ID`，是`DhcpIp`的唯一标识。
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`

	// `DhcpIp`所在私有网络`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `DhcpIp`所在子网`ID`。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// `DhcpIp`的名称。
	DhcpIpName *string `json:"DhcpIpName,omitnil" name:"DhcpIpName"`

	// IP地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 绑定`EIP`。
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`

	// `DhcpIp`关联弹性网卡`ID`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 被绑定的实例`ID`。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 状态：
	// <li>`AVAILABLE`：运行中</li>
	// <li>`UNBIND`：未绑定</li>
	State *string `json:"State,omitnil" name:"State"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type DirectConnectGateway struct {
	// 专线网关`ID`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 专线网关名称。
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil" name:"DirectConnectGatewayName"`

	// 专线网关关联`VPC`实例`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 关联网络类型：
	// <li>`VPC` - 私有网络</li>
	// <li>`CCN` - 云联网</li>
	NetworkType *string `json:"NetworkType,omitnil" name:"NetworkType"`

	// 关联网络实例`ID`：
	// <li>`NetworkType`为`VPC`时，这里为私有网络实例`ID`</li>
	// <li>`NetworkType`为`CCN`时，这里为云联网实例`ID`</li>
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil" name:"NetworkInstanceId"`

	// 网关类型：
	// <li>NORMAL - 标准型，注：云联网只支持标准型</li>
	// <li>NAT - NAT型</li>
	// NAT类型支持网络地址转换配置，类型确定后不能修改；一个私有网络可以创建一个NAT类型的专线网关和一个非NAT类型的专线网关
	GatewayType *string `json:"GatewayType,omitnil" name:"GatewayType"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 专线网关IP。
	DirectConnectGatewayIp *string `json:"DirectConnectGatewayIp,omitnil" name:"DirectConnectGatewayIp"`

	// 专线网关关联`CCN`实例`ID`。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网路由学习类型：
	// <li>`BGP` - 自动学习。</li>
	// <li>`STATIC` - 静态，即用户配置。</li>
	CcnRouteType *string `json:"CcnRouteType,omitnil" name:"CcnRouteType"`

	// 是否启用BGP。
	EnableBGP *bool `json:"EnableBGP,omitnil" name:"EnableBGP"`

	// 开启和关闭BGP的community属性。
	EnableBGPCommunity *bool `json:"EnableBGPCommunity,omitnil" name:"EnableBGPCommunity"`

	// 绑定的NAT网关ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 专线网关是否支持VXLAN架构
	// 注意：此字段可能返回 null，表示取不到有效值。
	VXLANSupport []*bool `json:"VXLANSupport,omitnil" name:"VXLANSupport"`

	// 云联网路由发布模式：`standard`（标准模式）、`exquisite`（精细模式）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModeType *string `json:"ModeType,omitnil" name:"ModeType"`

	// 是否为localZone专线网关。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalZone *bool `json:"LocalZone,omitnil" name:"LocalZone"`

	// 专线网关所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 网关流控明细启用状态：
	// 0：关闭
	// 1：开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableFlowDetails *uint64 `json:"EnableFlowDetails,omitnil" name:"EnableFlowDetails"`

	// 开启、关闭网关流控明细时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowDetailsUpdateTime *string `json:"FlowDetailsUpdateTime,omitnil" name:"FlowDetailsUpdateTime"`

	// 是否支持开启网关流控明细
	// 0：不支持
	// 1：支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewAfc *uint64 `json:"NewAfc,omitnil" name:"NewAfc"`

	// 专线网关接入网络类型：
	// <li>`VXLAN` - VXLAN类型。</li>
	// <li>`MPLS` - MPLS类型。</li>
	// <li>`Hybrid` - Hybrid类型。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessNetworkType *string `json:"AccessNetworkType,omitnil" name:"AccessNetworkType"`

	// 跨可用区容灾专线网关的可用区列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HaZoneList []*string `json:"HaZoneList,omitnil" name:"HaZoneList"`
}

type DirectConnectGatewayCcnRoute struct {
	// 路由ID。
	RouteId *string `json:"RouteId,omitnil" name:"RouteId"`

	// IDC网段。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil" name:"DestinationCidrBlock"`

	// `BGP`的`AS-Path`属性。
	ASPath []*string `json:"ASPath,omitnil" name:"ASPath"`

	// 备注
	Description *string `json:"Description,omitnil" name:"Description"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type DirectConnectSubnet struct {
	// 专线网关ID
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// IDC子网网段
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`
}

// Predefined struct for user
type DisableCcnRoutesRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

type DisableCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

func (r *DisableCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableCcnRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DisableCcnRoutesResponseParams `json:"Response"`
}

func (r *DisableCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableFlowLogsRequestParams struct {
	// 流日志Id。
	FlowLogIds []*string `json:"FlowLogIds,omitnil" name:"FlowLogIds"`
}

type DisableFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 流日志Id。
	FlowLogIds []*string `json:"FlowLogIds,omitnil" name:"FlowLogIds"`
}

func (r *DisableFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableFlowLogsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DisableFlowLogsResponseParams `json:"Response"`
}

func (r *DisableFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableGatewayFlowMonitorRequestParams struct {
	// 网关实例ID，目前我们支持的网关实例类型有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type DisableGatewayFlowMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID，目前我们支持的网关实例类型有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *DisableGatewayFlowMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableGatewayFlowMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableGatewayFlowMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableGatewayFlowMonitorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableGatewayFlowMonitorResponse struct {
	*tchttp.BaseResponse
	Response *DisableGatewayFlowMonitorResponseParams `json:"Response"`
}

func (r *DisableGatewayFlowMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableGatewayFlowMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRoutesRequestParams struct {
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略ID。不能和RouteItemIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteIds []*uint64 `json:"RouteIds,omitnil" name:"RouteIds"`

	// 路由策略唯一ID。不能和RouteIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

type DisableRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略ID。不能和RouteItemIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteIds []*uint64 `json:"RouteIds,omitnil" name:"RouteIds"`

	// 路由策略唯一ID。不能和RouteIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

func (r *DisableRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteIds")
	delete(f, "RouteItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DisableRoutesResponseParams `json:"Response"`
}

func (r *DisableRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableSnapshotPoliciesRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`
}

type DisableSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`
}

func (r *DisableSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableSnapshotPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DisableSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *DisableSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableVpnGatewaySslClientCertRequestParams struct {
	// SSL-VPN-CLIENT 实例ID。不可和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SSL-VPN-CLIENT 实例ID列表。批量禁用时使用。不可和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

type DisableVpnGatewaySslClientCertRequest struct {
	*tchttp.BaseRequest
	
	// SSL-VPN-CLIENT 实例ID。不可和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SSL-VPN-CLIENT 实例ID列表。批量禁用时使用。不可和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

func (r *DisableVpnGatewaySslClientCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableVpnGatewaySslClientCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SslVpnClientId")
	delete(f, "SslVpnClientIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableVpnGatewaySslClientCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableVpnGatewaySslClientCertResponseParams struct {
	// 异步任务实例ID。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisableVpnGatewaySslClientCertResponse struct {
	*tchttp.BaseResponse
	Response *DisableVpnGatewaySslClientCertResponseParams `json:"Response"`
}

func (r *DisableVpnGatewaySslClientCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableVpnGatewaySslClientCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressRequestParams struct {
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 表示解绑 EIP 之后是否分配普通公网 IP。取值范围：<br><li>TRUE：表示解绑 EIP 之后分配普通公网 IP。<br><li>FALSE：表示解绑 EIP 之后不分配普通公网 IP。<br>默认取值：FALSE。<br><br>只有满足以下条件时才能指定该参数：<br><li> 只有在解绑主网卡的主内网 IP 上的 EIP 时才能指定该参数。<br><li>解绑 EIP 后重新分配普通公网 IP 操作一个账号每天最多操作 10 次；详情可通过 [DescribeAddressQuota](https://cloud.tencent.com/document/api/213/1378) 接口获取。
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitnil" name:"ReallocateNormalPublicIp"`
}

type DisassociateAddressRequest struct {
	*tchttp.BaseRequest
	
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 表示解绑 EIP 之后是否分配普通公网 IP。取值范围：<br><li>TRUE：表示解绑 EIP 之后分配普通公网 IP。<br><li>FALSE：表示解绑 EIP 之后不分配普通公网 IP。<br>默认取值：FALSE。<br><br>只有满足以下条件时才能指定该参数：<br><li> 只有在解绑主网卡的主内网 IP 上的 EIP 时才能指定该参数。<br><li>解绑 EIP 后重新分配普通公网 IP 操作一个账号每天最多操作 10 次；详情可通过 [DescribeAddressQuota](https://cloud.tencent.com/document/api/213/1378) 接口获取。
	ReallocateNormalPublicIp *bool `json:"ReallocateNormalPublicIp,omitnil" name:"ReallocateNormalPublicIp"`
}

func (r *DisassociateAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressId")
	delete(f, "ReallocateNormalPublicIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateAddressResponseParams struct {
	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateAddressResponseParams `json:"Response"`
}

func (r *DisassociateAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDhcpIpWithAddressIpRequestParams struct {
	// `DhcpIp`唯一`ID`，形如：`dhcpip-9o233uri`。必须是已绑定`EIP`的`DhcpIp`。
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`
}

type DisassociateDhcpIpWithAddressIpRequest struct {
	*tchttp.BaseRequest
	
	// `DhcpIp`唯一`ID`，形如：`dhcpip-9o233uri`。必须是已绑定`EIP`的`DhcpIp`。
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`
}

func (r *DisassociateDhcpIpWithAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDhcpIpWithAddressIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DhcpIpId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateDhcpIpWithAddressIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDhcpIpWithAddressIpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateDhcpIpWithAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateDhcpIpWithAddressIpResponseParams `json:"Response"`
}

func (r *DisassociateDhcpIpWithAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDhcpIpWithAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDirectConnectGatewayNatGatewayRequestParams struct {
	// 专线网关ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关ID。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`
}

type DisassociateDirectConnectGatewayNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关ID。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`
}

func (r *DisassociateDirectConnectGatewayNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDirectConnectGatewayNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NatGatewayId")
	delete(f, "DirectConnectGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateDirectConnectGatewayNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateDirectConnectGatewayNatGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateDirectConnectGatewayNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateDirectConnectGatewayNatGatewayResponseParams `json:"Response"`
}

func (r *DisassociateDirectConnectGatewayNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateDirectConnectGatewayNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNatGatewayAddressRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 待解绑NAT网关的弹性IP数组。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`
}

type DisassociateNatGatewayAddressRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 待解绑NAT网关的弹性IP数组。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`
}

func (r *DisassociateNatGatewayAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNatGatewayAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "PublicIpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateNatGatewayAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNatGatewayAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateNatGatewayAddressResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateNatGatewayAddressResponseParams `json:"Response"`
}

func (r *DisassociateNatGatewayAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNatGatewayAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkAclSubnetsRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 子网实例ID数组。例如：[subnet-12345678]。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`
}

type DisassociateNetworkAclSubnetsRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 子网实例ID数组。例如：[subnet-12345678]。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`
}

func (r *DisassociateNetworkAclSubnetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkAclSubnetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "SubnetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateNetworkAclSubnetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkAclSubnetsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateNetworkAclSubnetsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateNetworkAclSubnetsResponseParams `json:"Response"`
}

func (r *DisassociateNetworkAclSubnetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkAclSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkInterfaceSecurityGroupsRequestParams struct {
	// 弹性网卡实例ID。形如：eni-pxir56ns。每次请求的实例的上限为100。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。每次请求的实例的上限为100。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type DisassociateNetworkInterfaceSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID。形如：eni-pxir56ns。每次请求的实例的上限为100。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。每次请求的实例的上限为100。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *DisassociateNetworkInterfaceSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkInterfaceSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateNetworkInterfaceSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateNetworkInterfaceSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateNetworkInterfaceSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateNetworkInterfaceSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateNetworkInterfaceSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateNetworkInterfaceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateVpcEndPointSecurityGroupsRequestParams struct {
	// 安全组ID数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 终端节点ID。
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`
}

type DisassociateVpcEndPointSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 安全组ID数组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 终端节点ID。
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`
}

func (r *DisassociateVpcEndPointSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateVpcEndPointSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupIds")
	delete(f, "EndPointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateVpcEndPointSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateVpcEndPointSecurityGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DisassociateVpcEndPointSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateVpcEndPointSecurityGroupsResponseParams `json:"Response"`
}

func (r *DisassociateVpcEndPointSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateVpcEndPointSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomerGatewayConfigurationRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`

	// 对端网关厂商信息对象，可通过[DescribeCustomerGatewayVendors](https://cloud.tencent.com/document/api/215/17513)获取。
	CustomerGatewayVendor *CustomerGatewayVendor `json:"CustomerGatewayVendor,omitnil" name:"CustomerGatewayVendor"`

	// 通道接入设备物理接口名称。
	InterfaceName *string `json:"InterfaceName,omitnil" name:"InterfaceName"`
}

type DownloadCustomerGatewayConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`

	// 对端网关厂商信息对象，可通过[DescribeCustomerGatewayVendors](https://cloud.tencent.com/document/api/215/17513)获取。
	CustomerGatewayVendor *CustomerGatewayVendor `json:"CustomerGatewayVendor,omitnil" name:"CustomerGatewayVendor"`

	// 通道接入设备物理接口名称。
	InterfaceName *string `json:"InterfaceName,omitnil" name:"InterfaceName"`
}

func (r *DownloadCustomerGatewayConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomerGatewayConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnConnectionId")
	delete(f, "CustomerGatewayVendor")
	delete(f, "InterfaceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadCustomerGatewayConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadCustomerGatewayConfigurationResponseParams struct {
	// XML格式配置信息。
	CustomerGatewayConfiguration *string `json:"CustomerGatewayConfiguration,omitnil" name:"CustomerGatewayConfiguration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DownloadCustomerGatewayConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DownloadCustomerGatewayConfigurationResponseParams `json:"Response"`
}

func (r *DownloadCustomerGatewayConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadCustomerGatewayConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadVpnGatewaySslClientCertRequestParams struct {
	// SSL-VPN-CLIENT 实例ID。不可以和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SAML Token（SAML令牌）。
	SamlToken *string `json:"SamlToken,omitnil" name:"SamlToken"`

	// VPN门户网站使用。默认False
	IsVpnPortal *bool `json:"IsVpnPortal,omitnil" name:"IsVpnPortal"`

	// SSL-VPN-CLIENT 实例ID列表。批量下载时使用。不可以和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

type DownloadVpnGatewaySslClientCertRequest struct {
	*tchttp.BaseRequest
	
	// SSL-VPN-CLIENT 实例ID。不可以和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SAML Token（SAML令牌）。
	SamlToken *string `json:"SamlToken,omitnil" name:"SamlToken"`

	// VPN门户网站使用。默认False
	IsVpnPortal *bool `json:"IsVpnPortal,omitnil" name:"IsVpnPortal"`

	// SSL-VPN-CLIENT 实例ID列表。批量下载时使用。不可以和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

func (r *DownloadVpnGatewaySslClientCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadVpnGatewaySslClientCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SslVpnClientId")
	delete(f, "SamlToken")
	delete(f, "IsVpnPortal")
	delete(f, "SslVpnClientIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadVpnGatewaySslClientCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadVpnGatewaySslClientCertResponseParams struct {
	// SSL-VPN 客户端配置。
	SslClientConfigsSet *string `json:"SslClientConfigsSet,omitnil" name:"SslClientConfigsSet"`

	// SSL-VPN 客户端配置。
	SslClientConfig []*SslClientConfig `json:"SslClientConfig,omitnil" name:"SslClientConfig"`

	// 是否鉴权成功 只有传入SamlToken 才生效，1为成功，0为失败。
	Authenticated *uint64 `json:"Authenticated,omitnil" name:"Authenticated"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DownloadVpnGatewaySslClientCertResponse struct {
	*tchttp.BaseResponse
	Response *DownloadVpnGatewaySslClientCertResponseParams `json:"Response"`
}

func (r *DownloadVpnGatewaySslClientCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadVpnGatewaySslClientCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableCcnRoutesRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

type EnableCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN路由策略唯一ID。形如：ccnr-f49l6u0z。
	RouteIds []*string `json:"RouteIds,omitnil" name:"RouteIds"`
}

func (r *EnableCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "RouteIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableCcnRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *EnableCcnRoutesResponseParams `json:"Response"`
}

func (r *EnableCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableFlowLogsRequestParams struct {
	// 流日志Id。
	FlowLogIds []*string `json:"FlowLogIds,omitnil" name:"FlowLogIds"`
}

type EnableFlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 流日志Id。
	FlowLogIds []*string `json:"FlowLogIds,omitnil" name:"FlowLogIds"`
}

func (r *EnableFlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableFlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableFlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableFlowLogsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableFlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *EnableFlowLogsResponseParams `json:"Response"`
}

func (r *EnableFlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableFlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGatewayFlowMonitorRequestParams struct {
	// 网关实例ID，目前我们支持的网关实例有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type EnableGatewayFlowMonitorRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID，目前我们支持的网关实例有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *EnableGatewayFlowMonitorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGatewayFlowMonitorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableGatewayFlowMonitorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableGatewayFlowMonitorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableGatewayFlowMonitorResponse struct {
	*tchttp.BaseResponse
	Response *EnableGatewayFlowMonitorResponseParams `json:"Response"`
}

func (r *EnableGatewayFlowMonitorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableGatewayFlowMonitorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRoutesRequestParams struct {
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略ID。不能和RouteItemIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteIds []*uint64 `json:"RouteIds,omitnil" name:"RouteIds"`

	// 路由策略唯一ID。不能和RouteIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

type EnableRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略ID。不能和RouteItemIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteIds []*uint64 `json:"RouteIds,omitnil" name:"RouteIds"`

	// 路由策略唯一ID。不能和RouteIds同时使用，但至少输入一个。该参数取值可通过查询路由列表（[DescribeRouteTables](https://cloud.tencent.com/document/product/215/15763)）获取。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

func (r *EnableRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteIds")
	delete(f, "RouteItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableRoutesResponse struct {
	*tchttp.BaseResponse
	Response *EnableRoutesResponseParams `json:"Response"`
}

func (r *EnableRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSnapshotPoliciesRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`
}

type EnableSnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyIds []*string `json:"SnapshotPolicyIds,omitnil" name:"SnapshotPolicyIds"`
}

func (r *EnableSnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableSnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableSnapshotPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableSnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *EnableSnapshotPoliciesResponseParams `json:"Response"`
}

func (r *EnableSnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableSnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcEndPointConnectRequestParams struct {
	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 终端节点ID。
	EndPointId []*string `json:"EndPointId,omitnil" name:"EndPointId"`

	// 是否接受终端节点连接请求。
	// <li> true：自动接受。</li> <li> false：不自动接受。</li>
	AcceptFlag *bool `json:"AcceptFlag,omitnil" name:"AcceptFlag"`
}

type EnableVpcEndPointConnectRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 终端节点ID。
	EndPointId []*string `json:"EndPointId,omitnil" name:"EndPointId"`

	// 是否接受终端节点连接请求。
	// <li> true：自动接受。</li> <li> false：不自动接受。</li>
	AcceptFlag *bool `json:"AcceptFlag,omitnil" name:"AcceptFlag"`
}

func (r *EnableVpcEndPointConnectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcEndPointConnectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	delete(f, "EndPointId")
	delete(f, "AcceptFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableVpcEndPointConnectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpcEndPointConnectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableVpcEndPointConnectResponse struct {
	*tchttp.BaseResponse
	Response *EnableVpcEndPointConnectResponseParams `json:"Response"`
}

func (r *EnableVpcEndPointConnectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpcEndPointConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpnGatewaySslClientCertRequestParams struct {
	// SSL-VPN-CLIENT 实例ID。不可和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SSL-VPN-CLIENT 实例ID列表。批量启用时使用。不可和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

type EnableVpnGatewaySslClientCertRequest struct {
	*tchttp.BaseRequest
	
	// SSL-VPN-CLIENT 实例ID。不可和SslVpnClientIds同时使用。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// SSL-VPN-CLIENT 实例ID列表。批量启用时使用。不可和SslVpnClientId同时使用。
	SslVpnClientIds []*string `json:"SslVpnClientIds,omitnil" name:"SslVpnClientIds"`
}

func (r *EnableVpnGatewaySslClientCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpnGatewaySslClientCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SslVpnClientId")
	delete(f, "SslVpnClientIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableVpnGatewaySslClientCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableVpnGatewaySslClientCertResponseParams struct {
	// 异步任务实例ID。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EnableVpnGatewaySslClientCertResponse struct {
	*tchttp.BaseResponse
	Response *EnableVpnGatewaySslClientCertResponseParams `json:"Response"`
}

func (r *EnableVpnGatewaySslClientCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableVpnGatewaySslClientCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EndPoint struct {
	// 终端节点ID。
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`

	// VPCID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// APPID。
	EndPointOwner *string `json:"EndPointOwner,omitnil" name:"EndPointOwner"`

	// 终端节点名称。
	EndPointName *string `json:"EndPointName,omitnil" name:"EndPointName"`

	// 终端节点服务的VPCID。
	ServiceVpcId *string `json:"ServiceVpcId,omitnil" name:"ServiceVpcId"`

	// 终端节点服务的VIP。
	ServiceVip *string `json:"ServiceVip,omitnil" name:"ServiceVip"`

	// 终端节点服务的ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 终端节点的VIP。
	EndPointVip *string `json:"EndPointVip,omitnil" name:"EndPointVip"`

	// 终端节点状态，ACTIVE：可用，PENDING：待接受，ACCEPTING：接受中，REJECTED：已拒绝，FAILED：失败。
	State *string `json:"State,omitnil" name:"State"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 终端节点绑定的安全组实例ID列表。
	GroupSet []*string `json:"GroupSet,omitnil" name:"GroupSet"`

	// 终端节点服务名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`
}

type EndPointService struct {
	// 终端节点服务ID
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// VPCID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// APPID。
	ServiceOwner *string `json:"ServiceOwner,omitnil" name:"ServiceOwner"`

	// 终端节点服务名称。
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 后端服务的VIP。
	ServiceVip *string `json:"ServiceVip,omitnil" name:"ServiceVip"`

	// 后端服务的ID，比如lb-xxx。
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil" name:"ServiceInstanceId"`

	// 是否自动接受。
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil" name:"AutoAcceptFlag"`

	// 关联的终端节点个数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPointCount *uint64 `json:"EndPointCount,omitnil" name:"EndPointCount"`

	// 终端节点对象数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPointSet []*EndPoint `json:"EndPointSet,omitnil" name:"EndPointSet"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 挂载的PAAS服务类型，CLB,CDB,CRS
	ServiceType *string `json:"ServiceType,omitnil" name:"ServiceType"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。当值类型为布尔类型时，可直接取值为字符串"TRUE"或 "FALSE"。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FilterObject struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FlowLog struct {
	// 私用网络ID或者统一ID，建议使用统一ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`

	// 流日志实例名字。
	FlowLogName *string `json:"FlowLogName,omitnil" name:"FlowLogName"`

	// 流日志所属资源类型，VPC|SUBNET|NETWORKINTERFACE|CCN|NAT|DCG。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源唯一ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 流日志采集类型，ACCEPT|REJECT|ALL。
	TrafficType *string `json:"TrafficType,omitnil" name:"TrafficType"`

	// 流日志存储ID。
	CloudLogId *string `json:"CloudLogId,omitnil" name:"CloudLogId"`

	// 流日志存储ID状态。
	CloudLogState *string `json:"CloudLogState,omitnil" name:"CloudLogState"`

	// 流日志描述信息。
	FlowLogDescription *string `json:"FlowLogDescription,omitnil" name:"FlowLogDescription"`

	// 流日志创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 标签列表，例如：[{"Key": "city", "Value": "shanghai"}]。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 是否启用，true-启用，false-停用。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 消费端类型：cls、ckafka。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// 消费端信息，当消费端类型为ckafka时返回。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowLogStorage *FlowLogStorage `json:"FlowLogStorage,omitnil" name:"FlowLogStorage"`

	// 流日志存储ID对应的地域信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudLogRegion *string `json:"CloudLogRegion,omitnil" name:"CloudLogRegion"`
}

type FlowLogStorage struct {
	// 存储实例Id，当流日志存储类型为ckafka时，必填。
	StorageId *string `json:"StorageId,omitnil" name:"StorageId"`

	// 主题Id，当流日志存储类型为ckafka时，必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageTopic *string `json:"StorageTopic,omitnil" name:"StorageTopic"`
}

type GatewayFlowMonitorDetail struct {
	// 来源`IP`。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 入包量。
	InPkg *uint64 `json:"InPkg,omitnil" name:"InPkg"`

	// 出包量。
	OutPkg *uint64 `json:"OutPkg,omitnil" name:"OutPkg"`

	// 入流量，单位：`Byte`。
	InTraffic *uint64 `json:"InTraffic,omitnil" name:"InTraffic"`

	// 出流量，单位：`Byte`。
	OutTraffic *uint64 `json:"OutTraffic,omitnil" name:"OutTraffic"`
}

type GatewayQos struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 云服务器内网IP。
	IpAddress *string `json:"IpAddress,omitnil" name:"IpAddress"`

	// 流控带宽值。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

// Predefined struct for user
type GenerateVpnConnectionDefaultHealthCheckIpRequestParams struct {
	// VPN网关id， 例如：vpngw-1w9tue3d
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`
}

type GenerateVpnConnectionDefaultHealthCheckIpRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关id， 例如：vpngw-1w9tue3d
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`
}

func (r *GenerateVpnConnectionDefaultHealthCheckIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateVpnConnectionDefaultHealthCheckIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateVpnConnectionDefaultHealthCheckIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateVpnConnectionDefaultHealthCheckIpResponseParams struct {
	// VPN通道健康检查本端ip
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil" name:"HealthCheckLocalIp"`

	// VPN通道健康检查对端ip
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil" name:"HealthCheckRemoteIp"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GenerateVpnConnectionDefaultHealthCheckIpResponse struct {
	*tchttp.BaseResponse
	Response *GenerateVpnConnectionDefaultHealthCheckIpResponseParams `json:"Response"`
}

func (r *GenerateVpnConnectionDefaultHealthCheckIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateVpnConnectionDefaultHealthCheckIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCcnRegionBandwidthLimitsRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 过滤条件。
	// <li>sregion - String - （过滤条件）源地域，形如：ap-guangzhou。</li>
	// <li>dregion - String - （过滤条件）目的地域，形如：ap-shanghai-bm</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序条件，目前支持带宽（`BandwidthLimit`）和过期时间（`ExpireTime`），默认按 `ExpireTime` 排序。
	SortedBy *string `json:"SortedBy,omitnil" name:"SortedBy"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，'ASC':升序,'DESC':降序。默认按'ASC'排序。
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`
}

type GetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 过滤条件。
	// <li>sregion - String - （过滤条件）源地域，形如：ap-guangzhou。</li>
	// <li>dregion - String - （过滤条件）目的地域，形如：ap-shanghai-bm</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 排序条件，目前支持带宽（`BandwidthLimit`）和过期时间（`ExpireTime`），默认按 `ExpireTime` 排序。
	SortedBy *string `json:"SortedBy,omitnil" name:"SortedBy"`

	// 偏移量。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，'ASC':升序,'DESC':降序。默认按'ASC'排序。
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`
}

func (r *GetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Filters")
	delete(f, "SortedBy")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCcnRegionBandwidthLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCcnRegionBandwidthLimitsResponseParams struct {
	// 云联网（CCN）各地域出带宽详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnBandwidthSet []*CcnBandwidthInfo `json:"CcnBandwidthSet,omitnil" name:"CcnBandwidthSet"`

	// 符合条件的对象数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *GetCcnRegionBandwidthLimitsResponseParams `json:"Response"`
}

func (r *GetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HaVip struct {
	// `HAVIP`的`ID`，是`HAVIP`的唯一标识。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`

	// `HAVIP`名称。
	HaVipName *string `json:"HaVipName,omitnil" name:"HaVipName"`

	// 虚拟IP地址。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// `HAVIP`所在私有网络`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `HAVIP`所在子网`ID`。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// `HAVIP`关联弹性网卡`ID`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 被绑定的实例`ID`。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 绑定`EIP`。
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`

	// 状态：
	// <li>`AVAILABLE`：运行中</li>
	// <li>`UNBIND`：未绑定</li>
	State *string `json:"State,omitnil" name:"State"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 使用havip的业务标识。
	Business *string `json:"Business,omitnil" name:"Business"`
}

// Predefined struct for user
type HaVipAssociateAddressIpRequestParams struct {
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。必须是没有绑定`EIP`的`HAVIP`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`

	// 弹性公网`IP`。必须是没有绑定`HAVIP`的`EIP`。
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`
}

type HaVipAssociateAddressIpRequest struct {
	*tchttp.BaseRequest
	
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。必须是没有绑定`EIP`的`HAVIP`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`

	// 弹性公网`IP`。必须是没有绑定`HAVIP`的`EIP`。
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`
}

func (r *HaVipAssociateAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipAssociateAddressIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	delete(f, "AddressIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HaVipAssociateAddressIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HaVipAssociateAddressIpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type HaVipAssociateAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *HaVipAssociateAddressIpResponseParams `json:"Response"`
}

func (r *HaVipAssociateAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipAssociateAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HaVipDisassociateAddressIpRequestParams struct {
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。必须是已绑定`EIP`的`HAVIP`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`
}

type HaVipDisassociateAddressIpRequest struct {
	*tchttp.BaseRequest
	
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。必须是已绑定`EIP`的`HAVIP`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`
}

func (r *HaVipDisassociateAddressIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipDisassociateAddressIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "HaVipDisassociateAddressIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type HaVipDisassociateAddressIpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type HaVipDisassociateAddressIpResponse struct {
	*tchttp.BaseResponse
	Response *HaVipDisassociateAddressIpResponseParams `json:"Response"`
}

func (r *HaVipDisassociateAddressIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *HaVipDisassociateAddressIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IKEOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBS-192', 'AES-CBC-256', 'DES-CBC'，'SM4', 默认为3DES-CBC
	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm,omitnil" name:"PropoEncryAlgorithm"`

	// 认证算法：可选值：'MD5', 'SHA1'，'SHA-256' 默认为MD5
	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm,omitnil" name:"PropoAuthenAlgorithm"`

	// 协商模式：可选值：'AGGRESSIVE', 'MAIN'，默认为MAIN
	ExchangeMode *string `json:"ExchangeMode,omitnil" name:"ExchangeMode"`

	// 本端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	LocalIdentity *string `json:"LocalIdentity,omitnil" name:"LocalIdentity"`

	// 对端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	RemoteIdentity *string `json:"RemoteIdentity,omitnil" name:"RemoteIdentity"`

	// 本端标识，当LocalIdentity选为ADDRESS时，LocalAddress必填。localAddress默认为vpn网关公网IP
	LocalAddress *string `json:"LocalAddress,omitnil" name:"LocalAddress"`

	// 对端标识，当RemoteIdentity选为ADDRESS时，RemoteAddress必填
	RemoteAddress *string `json:"RemoteAddress,omitnil" name:"RemoteAddress"`

	// 本端标识，当LocalIdentity选为FQDN时，LocalFqdnName必填
	LocalFqdnName *string `json:"LocalFqdnName,omitnil" name:"LocalFqdnName"`

	// 对端标识，当remoteIdentity选为FQDN时，RemoteFqdnName必填
	RemoteFqdnName *string `json:"RemoteFqdnName,omitnil" name:"RemoteFqdnName"`

	// DH group，指定IKE交换密钥时使用的DH组，可选值：'GROUP1', 'GROUP2', 'GROUP5', 'GROUP14', 'GROUP24'，
	DhGroupName *string `json:"DhGroupName,omitnil" name:"DhGroupName"`

	// IKE SA Lifetime，单位：秒，设置IKE SA的生存周期，取值范围：60-604800
	IKESaLifetimeSeconds *uint64 `json:"IKESaLifetimeSeconds,omitnil" name:"IKESaLifetimeSeconds"`

	// IKE版本
	IKEVersion *string `json:"IKEVersion,omitnil" name:"IKEVersion"`
}

type IPSECOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBC-192', 'AES-CBC-256', 'DES-CBC', 'SM4', 'NULL'， 默认为AES-CBC-128
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil" name:"EncryptAlgorithm"`

	// 认证算法：可选值：'MD5', 'SHA1'，'SHA-256' 默认为
	IntegrityAlgorith *string `json:"IntegrityAlgorith,omitnil" name:"IntegrityAlgorith"`

	// IPsec SA lifetime(s)：单位秒，取值范围：180-604800
	IPSECSaLifetimeSeconds *uint64 `json:"IPSECSaLifetimeSeconds,omitnil" name:"IPSECSaLifetimeSeconds"`

	// PFS：可选值：'NULL', 'DH-GROUP1', 'DH-GROUP2', 'DH-GROUP5', 'DH-GROUP14', 'DH-GROUP24'，默认为NULL
	PfsDhGroup *string `json:"PfsDhGroup,omitnil" name:"PfsDhGroup"`

	// IPsec SA lifetime(KB)：单位KB，取值范围：2560-604800
	IPSECSaLifetimeTraffic *uint64 `json:"IPSECSaLifetimeTraffic,omitnil" name:"IPSECSaLifetimeTraffic"`
}

// Predefined struct for user
type InquirePriceCreateDirectConnectGatewayRequestParams struct {

}

type InquirePriceCreateDirectConnectGatewayRequest struct {
	*tchttp.BaseRequest
	
}

func (r *InquirePriceCreateDirectConnectGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDirectConnectGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateDirectConnectGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateDirectConnectGatewayResponseParams struct {
	// 专线网关标准接入费用
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *int64 `json:"TotalCost,omitnil" name:"TotalCost"`

	// 专线网关真实接入费用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *int64 `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquirePriceCreateDirectConnectGatewayResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateDirectConnectGatewayResponseParams `json:"Response"`
}

func (r *InquirePriceCreateDirectConnectGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateDirectConnectGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateVpnGatewayRequestParams struct {
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// SSL VPN连接数设置，可选规格：5, 10, 20, 50, 100；单位：个。
	MaxConnection *uint64 `json:"MaxConnection,omitnil" name:"MaxConnection"`

	// 查询的VPN类型，支持IPSEC和SSL两种类型，为SSL类型时，MaxConnection参数必传。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type InquiryPriceCreateVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// VPN网关计费模式，PREPAID：表示预付费，即包年包月，POSTPAID_BY_HOUR：表示后付费，即按量计费。默认：POSTPAID_BY_HOUR，如果指定预付费模式，参数InstanceChargePrepaid必填。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`

	// SSL VPN连接数设置，可选规格：5, 10, 20, 50, 100；单位：个。
	MaxConnection *uint64 `json:"MaxConnection,omitnil" name:"MaxConnection"`

	// 查询的VPN类型，支持IPSEC和SSL两种类型，为SSL类型时，MaxConnection参数必传。
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *InquiryPriceCreateVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "MaxConnection")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateVpnGatewayResponseParams struct {
	// 商品价格。
	Price *Price `json:"Price,omitnil" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceCreateVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateVpnGatewayResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewVpnGatewayRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

type InquiryPriceRenewVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

func (r *InquiryPriceRenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewVpnGatewayResponseParams struct {
	// 商品价格。
	Price *Price `json:"Price,omitnil" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceRenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewVpnGatewayResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 公网带宽设置。可选带宽规格：5, 10, 20, 50, 100；单位：Mbps。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceResetVpnGatewayInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponseParams struct {
	// 商品价格。
	Price *Price `json:"Price,omitnil" name:"Price"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36。
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// 自动续费标识。取值范围： NOTIFY_AND_AUTO_RENEW：通知过期且自动续费， NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。默认：NOTIFY_AND_AUTO_RENEW
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type InstanceStatistic struct {
	// 实例的类型
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例的个数
	InstanceCount *uint64 `json:"InstanceCount,omitnil" name:"InstanceCount"`
}

type Ip6Rule struct {
	// IPV6转换规则唯一ID，形如rule6-xxxxxxxx
	Ip6RuleId *string `json:"Ip6RuleId,omitnil" name:"Ip6RuleId"`

	// IPV6转换规则名称
	Ip6RuleName *string `json:"Ip6RuleName,omitnil" name:"Ip6RuleName"`

	// IPV6地址
	Vip6 *string `json:"Vip6,omitnil" name:"Vip6"`

	// IPV6端口号
	Vport6 *int64 `json:"Vport6,omitnil" name:"Vport6"`

	// 协议类型，支持TCP/UDP
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// IPV4地址
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPV4端口号
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// 转换规则状态，限于CREATING,RUNNING,DELETING,MODIFYING
	RuleStatus *string `json:"RuleStatus,omitnil" name:"RuleStatus"`

	// 转换规则创建时间
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type Ip6RuleInfo struct {
	// IPV6端口号，可在0~65535范围取值
	Vport6 *int64 `json:"Vport6,omitnil" name:"Vport6"`

	// 协议类型，支持TCP/UDP
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// IPV4地址
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPV4端口号，可在0~65535范围取值
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`
}

type Ip6Translator struct {
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// IPV6转换实例名称
	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitnil" name:"Ip6TranslatorName"`

	// IPV6地址
	Vip6 *string `json:"Vip6,omitnil" name:"Vip6"`

	// IPV6转换地址所属运营商
	IspName *string `json:"IspName,omitnil" name:"IspName"`

	// 转换实例状态，限于CREATING,RUNNING,DELETING,MODIFYING
	TranslatorStatus *string `json:"TranslatorStatus,omitnil" name:"TranslatorStatus"`

	// IPV6转换实例创建时间
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 绑定的IPV6转换规则数量
	Ip6RuleCount *int64 `json:"Ip6RuleCount,omitnil" name:"Ip6RuleCount"`

	// IPV6转换规则信息
	IP6RuleSet []*Ip6Rule `json:"IP6RuleSet,omitnil" name:"IP6RuleSet"`
}

type IpAddressStates struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// IP地址。
	IpAddress *string `json:"IpAddress,omitnil" name:"IpAddress"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type IpField struct {
	// 国家字段信息
	Country *bool `json:"Country,omitnil" name:"Country"`

	// 省、州、郡一级行政区域字段信息
	Province *bool `json:"Province,omitnil" name:"Province"`

	// 市一级行政区域字段信息
	City *bool `json:"City,omitnil" name:"City"`

	// 市内区域字段信息
	Region *bool `json:"Region,omitnil" name:"Region"`

	// 接入运营商字段信息
	Isp *bool `json:"Isp,omitnil" name:"Isp"`

	// 骨干运营商字段信息
	AsName *bool `json:"AsName,omitnil" name:"AsName"`

	// 骨干As号
	AsId *bool `json:"AsId,omitnil" name:"AsId"`

	// 注释字段
	Comment *bool `json:"Comment,omitnil" name:"Comment"`
}

type IpGeolocationInfo struct {
	// 国家信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitnil" name:"Country"`

	// 省、州、郡一级行政区域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Province *string `json:"Province,omitnil" name:"Province"`

	// 市一级行政区域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	City *string `json:"City,omitnil" name:"City"`

	// 市内区域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 接入运营商信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isp *string `json:"Isp,omitnil" name:"Isp"`

	// 骨干运营商名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsName *string `json:"AsName,omitnil" name:"AsName"`

	// 骨干运营商AS号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsId *string `json:"AsId,omitnil" name:"AsId"`

	// 注释信息。目前的填充值为移动接入用户的APN值，如无APN属性则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`
}

type Ipv6Address struct {
	// `IPv6`地址，形如：`3402:4e00:20:100:0:8cd9:2a67:71f3`
	Address *string `json:"Address,omitnil" name:"Address"`

	// 是否是主`IP`。
	Primary *bool `json:"Primary,omitnil" name:"Primary"`

	// `EIP`实例`ID`，形如：`eip-hxlqja90`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 公网IP是否被封堵。
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitnil" name:"IsWanIpBlocked"`

	// `IPv6`地址状态：
	// <li>`PENDING`：生产中</li>
	// <li>`MIGRATING`：迁移中</li>
	// <li>`DELETING`：删除中</li>
	// <li>`AVAILABLE`：可用的</li>
	State *string `json:"State,omitnil" name:"State"`
}

type Ipv6SubnetCidrBlock struct {
	// 子网实例`ID`。形如：`subnet-pxir56ns`。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// `IPv6`子网段。形如：`3402:4e00:20:1001::/64`
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`
}

type ItemPrice struct {
	// 按量计费后付费单价，单位：元。
	UnitPrice *float64 `json:"UnitPrice,omitnil" name:"UnitPrice"`

	// 按量计费后付费计价单元，可取值范围： HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）： GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。
	ChargeUnit *string `json:"ChargeUnit,omitnil" name:"ChargeUnit"`

	// 预付费商品的原价，单位：元。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil" name:"OriginalPrice"`

	// 预付费商品的折扣价，单位：元。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil" name:"DiscountPrice"`
}

type LocalGateway struct {
	// CDC实例ID
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// VPC实例ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 本地网关实例ID
	UniqLocalGwId *string `json:"UniqLocalGwId,omitnil" name:"UniqLocalGwId"`

	// 本地网关名称
	LocalGatewayName *string `json:"LocalGatewayName,omitnil" name:"LocalGatewayName"`

	// 本地网关IP地址
	LocalGwIp *string `json:"LocalGwIp,omitnil" name:"LocalGwIp"`

	// 本地网关创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

// Predefined struct for user
type LockCcnBandwidthsRequestParams struct {
	// 带宽实例的唯一ID数组。
	Instances []*CcnFlowLock `json:"Instances,omitnil" name:"Instances"`
}

type LockCcnBandwidthsRequest struct {
	*tchttp.BaseRequest
	
	// 带宽实例的唯一ID数组。
	Instances []*CcnFlowLock `json:"Instances,omitnil" name:"Instances"`
}

func (r *LockCcnBandwidthsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockCcnBandwidthsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LockCcnBandwidthsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockCcnBandwidthsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type LockCcnBandwidthsResponse struct {
	*tchttp.BaseResponse
	Response *LockCcnBandwidthsResponseParams `json:"Response"`
}

func (r *LockCcnBandwidthsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockCcnBandwidthsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockCcnsRequestParams struct {

}

type LockCcnsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *LockCcnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockCcnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LockCcnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockCcnsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type LockCcnsResponse struct {
	*tchttp.BaseResponse
	Response *LockCcnsResponseParams `json:"Response"`
}

func (r *LockCcnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberInfo struct {
	// 模板对象成员
	Member *string `json:"Member,omitnil" name:"Member"`

	// 模板对象成员描述信息
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type MigrateNetworkInterfaceRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 弹性网卡当前绑定的CVM实例ID。形如：ins-r8hr2upy。
	SourceInstanceId *string `json:"SourceInstanceId,omitnil" name:"SourceInstanceId"`

	// 待迁移的目的CVM实例ID。
	DestinationInstanceId *string `json:"DestinationInstanceId,omitnil" name:"DestinationInstanceId"`

	// 网卡绑定类型：0 标准型 1 扩展型。
	AttachType *uint64 `json:"AttachType,omitnil" name:"AttachType"`
}

type MigrateNetworkInterfaceRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 弹性网卡当前绑定的CVM实例ID。形如：ins-r8hr2upy。
	SourceInstanceId *string `json:"SourceInstanceId,omitnil" name:"SourceInstanceId"`

	// 待迁移的目的CVM实例ID。
	DestinationInstanceId *string `json:"DestinationInstanceId,omitnil" name:"DestinationInstanceId"`

	// 网卡绑定类型：0 标准型 1 扩展型。
	AttachType *uint64 `json:"AttachType,omitnil" name:"AttachType"`
}

func (r *MigrateNetworkInterfaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateNetworkInterfaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "SourceInstanceId")
	delete(f, "DestinationInstanceId")
	delete(f, "AttachType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateNetworkInterfaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigrateNetworkInterfaceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MigrateNetworkInterfaceResponse struct {
	*tchttp.BaseResponse
	Response *MigrateNetworkInterfaceResponseParams `json:"Response"`
}

func (r *MigrateNetworkInterfaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateNetworkInterfaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigratePrivateIpAddressRequestParams struct {
	// 当内网IP绑定的弹性网卡实例ID，例如：eni-m6dyj72l。
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitnil" name:"SourceNetworkInterfaceId"`

	// 待迁移的目的弹性网卡实例ID。
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitnil" name:"DestinationNetworkInterfaceId"`

	// 迁移的内网IP地址，例如：10.0.0.6。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`
}

type MigratePrivateIpAddressRequest struct {
	*tchttp.BaseRequest
	
	// 当内网IP绑定的弹性网卡实例ID，例如：eni-m6dyj72l。
	SourceNetworkInterfaceId *string `json:"SourceNetworkInterfaceId,omitnil" name:"SourceNetworkInterfaceId"`

	// 待迁移的目的弹性网卡实例ID。
	DestinationNetworkInterfaceId *string `json:"DestinationNetworkInterfaceId,omitnil" name:"DestinationNetworkInterfaceId"`

	// 迁移的内网IP地址，例如：10.0.0.6。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`
}

func (r *MigratePrivateIpAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigratePrivateIpAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceNetworkInterfaceId")
	delete(f, "DestinationNetworkInterfaceId")
	delete(f, "PrivateIpAddress")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigratePrivateIpAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigratePrivateIpAddressResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type MigratePrivateIpAddressResponse struct {
	*tchttp.BaseResponse
	Response *MigratePrivateIpAddressResponseParams `json:"Response"`
}

func (r *MigratePrivateIpAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigratePrivateIpAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressAttributeRequestParams struct {
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 修改后的 EIP 名称。长度上限为20个字符。
	AddressName *string `json:"AddressName,omitnil" name:"AddressName"`

	// 设定EIP是否直通，"TRUE"表示直通，"FALSE"表示非直通。注意该参数仅对EIP直通功能可见的用户可以设定。
	EipDirectConnection *string `json:"EipDirectConnection,omitnil" name:"EipDirectConnection"`
}

type ModifyAddressAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 修改后的 EIP 名称。长度上限为20个字符。
	AddressName *string `json:"AddressName,omitnil" name:"AddressName"`

	// 设定EIP是否直通，"TRUE"表示直通，"FALSE"表示非直通。注意该参数仅对EIP直通功能可见的用户可以设定。
	EipDirectConnection *string `json:"EipDirectConnection,omitnil" name:"EipDirectConnection"`
}

func (r *ModifyAddressAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressId")
	delete(f, "AddressName")
	delete(f, "EipDirectConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAddressAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressAttributeResponseParams `json:"Response"`
}

func (r *ModifyAddressAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressInternetChargeTypeRequestParams struct {
	// 弹性公网IP的唯一ID，形如eip-xxx
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 弹性公网IP调整目标计费模式，只支持"BANDWIDTH_PREPAID_BY_MONTH"和"TRAFFIC_POSTPAID_BY_HOUR"
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// 弹性公网IP调整目标带宽值
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 包月带宽网络计费模式参数。弹性公网IP的调整目标计费模式是"BANDWIDTH_PREPAID_BY_MONTH"时，必传该参数。
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil" name:"AddressChargePrepaid"`
}

type ModifyAddressInternetChargeTypeRequest struct {
	*tchttp.BaseRequest
	
	// 弹性公网IP的唯一ID，形如eip-xxx
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 弹性公网IP调整目标计费模式，只支持"BANDWIDTH_PREPAID_BY_MONTH"和"TRAFFIC_POSTPAID_BY_HOUR"
	InternetChargeType *string `json:"InternetChargeType,omitnil" name:"InternetChargeType"`

	// 弹性公网IP调整目标带宽值
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 包月带宽网络计费模式参数。弹性公网IP的调整目标计费模式是"BANDWIDTH_PREPAID_BY_MONTH"时，必传该参数。
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil" name:"AddressChargePrepaid"`
}

func (r *ModifyAddressInternetChargeTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressInternetChargeTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressId")
	delete(f, "InternetChargeType")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "AddressChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressInternetChargeTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressInternetChargeTypeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAddressInternetChargeTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressInternetChargeTypeResponseParams `json:"Response"`
}

func (r *ModifyAddressInternetChargeTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressInternetChargeTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateAttributeRequestParams struct {
	// IP地址模板实例ID，例如：ipm-mdunqeb6。
	AddressTemplateId *string `json:"AddressTemplateId,omitnil" name:"AddressTemplateId"`

	// IP地址模板名称。
	AddressTemplateName *string `json:"AddressTemplateName,omitnil" name:"AddressTemplateName"`

	// 地址信息，支持 IP、CIDR、IP 范围。
	Addresses []*string `json:"Addresses,omitnil" name:"Addresses"`

	// 支持添加备注的地址信息，支持 IP、CIDR、IP 范围。
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil" name:"AddressesExtra"`
}

type ModifyAddressTemplateAttributeRequest struct {
	*tchttp.BaseRequest
	
	// IP地址模板实例ID，例如：ipm-mdunqeb6。
	AddressTemplateId *string `json:"AddressTemplateId,omitnil" name:"AddressTemplateId"`

	// IP地址模板名称。
	AddressTemplateName *string `json:"AddressTemplateName,omitnil" name:"AddressTemplateName"`

	// 地址信息，支持 IP、CIDR、IP 范围。
	Addresses []*string `json:"Addresses,omitnil" name:"Addresses"`

	// 支持添加备注的地址信息，支持 IP、CIDR、IP 范围。
	AddressesExtra []*AddressInfo `json:"AddressesExtra,omitnil" name:"AddressesExtra"`
}

func (r *ModifyAddressTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateId")
	delete(f, "AddressTemplateName")
	delete(f, "Addresses")
	delete(f, "AddressesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressTemplateAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAddressTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressTemplateAttributeResponseParams `json:"Response"`
}

func (r *ModifyAddressTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateGroupAttributeRequestParams struct {
	// IP地址模板集合实例ID，例如：ipmg-2uw6ujo6。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil" name:"AddressTemplateGroupId"`

	// IP地址模板集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil" name:"AddressTemplateGroupName"`

	// IP地址模板实例ID， 例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil" name:"AddressTemplateIds"`
}

type ModifyAddressTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// IP地址模板集合实例ID，例如：ipmg-2uw6ujo6。
	AddressTemplateGroupId *string `json:"AddressTemplateGroupId,omitnil" name:"AddressTemplateGroupId"`

	// IP地址模板集合名称。
	AddressTemplateGroupName *string `json:"AddressTemplateGroupName,omitnil" name:"AddressTemplateGroupName"`

	// IP地址模板实例ID， 例如：ipm-mdunqeb6。
	AddressTemplateIds []*string `json:"AddressTemplateIds,omitnil" name:"AddressTemplateIds"`
}

func (r *ModifyAddressTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressTemplateGroupId")
	delete(f, "AddressTemplateGroupName")
	delete(f, "AddressTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressTemplateGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateGroupAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAddressTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressTemplateGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyAddressTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthRequestParams struct {
	// EIP唯一标识ID列表，形如'eip-xxxx'
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`

	// 调整带宽目标值
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 包月带宽起始时间(已废弃，输入无效)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 包月带宽结束时间(已废弃，输入无效)
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type ModifyAddressesBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// EIP唯一标识ID列表，形如'eip-xxxx'
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`

	// 调整带宽目标值
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 包月带宽起始时间(已废弃，输入无效)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 包月带宽结束时间(已废弃，输入无效)
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *ModifyAddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIds")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressesBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressesBandwidthResponseParams struct {
	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressesBandwidthResponseParams `json:"Response"`
}

func (r *ModifyAddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssistantCidrRequestParams struct {
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 待添加的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]，入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil" name:"NewCidrBlocks"`

	// 待删除的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]，入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil" name:"OldCidrBlocks"`
}

type ModifyAssistantCidrRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`。形如：`vpc-6v2ht8q5`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 待添加的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]，入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	NewCidrBlocks []*string `json:"NewCidrBlocks,omitnil" name:"NewCidrBlocks"`

	// 待删除的辅助CIDR。CIDR数组，格式如["10.0.0.0/16", "172.16.0.0/16"]，入参NewCidrBlocks和OldCidrBlocks至少需要其一。
	OldCidrBlocks []*string `json:"OldCidrBlocks,omitnil" name:"OldCidrBlocks"`
}

func (r *ModifyAssistantCidrRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssistantCidrRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NewCidrBlocks")
	delete(f, "OldCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssistantCidrRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssistantCidrResponseParams struct {
	// 辅助CIDR数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil" name:"AssistantCidrSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAssistantCidrResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssistantCidrResponseParams `json:"Response"`
}

func (r *ModifyAssistantCidrResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssistantCidrResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBandwidthPackageAttributeRequestParams struct {
	// 带宽包唯一标识ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 带宽包名称
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil" name:"BandwidthPackageName"`

	// 带宽包计费模式，示例 ：
	// 'TOP5_POSTPAID_BY_MONTH'（后付费-TOP5计费）
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`
}

type ModifyBandwidthPackageAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 带宽包唯一标识ID
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 带宽包名称
	BandwidthPackageName *string `json:"BandwidthPackageName,omitnil" name:"BandwidthPackageName"`

	// 带宽包计费模式，示例 ：
	// 'TOP5_POSTPAID_BY_MONTH'（后付费-TOP5计费）
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`
}

func (r *ModifyBandwidthPackageAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBandwidthPackageAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "BandwidthPackageName")
	delete(f, "ChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBandwidthPackageAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBandwidthPackageAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBandwidthPackageAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBandwidthPackageAttributeResponseParams `json:"Response"`
}

func (r *ModifyBandwidthPackageAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBandwidthPackageAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttachedInstancesAttributeRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 关联网络实例列表
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

type ModifyCcnAttachedInstancesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 关联网络实例列表
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

func (r *ModifyCcnAttachedInstancesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttachedInstancesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcnAttachedInstancesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttachedInstancesAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCcnAttachedInstancesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcnAttachedInstancesAttributeResponseParams `json:"Response"`
}

func (r *ModifyCcnAttachedInstancesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttachedInstancesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttributeRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN名称，最大长度不能超过60个字节，限制：CcnName和CcnDescription必须至少选择一个参数输入，否则报错。
	CcnName *string `json:"CcnName,omitnil" name:"CcnName"`

	// CCN描述信息，最大长度不能超过100个字节，限制：CcnName和CcnDescription必须至少选择一个参数输入，否则报错。
	CcnDescription *string `json:"CcnDescription,omitnil" name:"CcnDescription"`
}

type ModifyCcnAttributeRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN名称，最大长度不能超过60个字节，限制：CcnName和CcnDescription必须至少选择一个参数输入，否则报错。
	CcnName *string `json:"CcnName,omitnil" name:"CcnName"`

	// CCN描述信息，最大长度不能超过100个字节，限制：CcnName和CcnDescription必须至少选择一个参数输入，否则报错。
	CcnDescription *string `json:"CcnDescription,omitnil" name:"CcnDescription"`
}

func (r *ModifyCcnAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "CcnName")
	delete(f, "CcnDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcnAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCcnAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcnAttributeResponseParams `json:"Response"`
}

func (r *ModifyCcnAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnRegionBandwidthLimitsTypeRequestParams struct {
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网限速类型，INTER_REGION_LIMIT：地域间限速，OUTER_REGION_LIMIT：地域出口限速。默认值：OUTER_REGION_LIMIT。
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil" name:"BandwidthLimitType"`
}

type ModifyCcnRegionBandwidthLimitsTypeRequest struct {
	*tchttp.BaseRequest
	
	// 云联网实例ID。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网限速类型，INTER_REGION_LIMIT：地域间限速，OUTER_REGION_LIMIT：地域出口限速。默认值：OUTER_REGION_LIMIT。
	BandwidthLimitType *string `json:"BandwidthLimitType,omitnil" name:"BandwidthLimitType"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnRegionBandwidthLimitsTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "BandwidthLimitType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCcnRegionBandwidthLimitsTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCcnRegionBandwidthLimitsTypeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCcnRegionBandwidthLimitsTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCcnRegionBandwidthLimitsTypeResponseParams `json:"Response"`
}

func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCcnRegionBandwidthLimitsTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomerGatewayAttributeRequestParams struct {
	// 对端网关ID，例如：cgw-2wqq41m9，可通过[DescribeCustomerGateways](https://cloud.tencent.com/document/api/215/17516)接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`

	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil" name:"CustomerGatewayName"`
}

type ModifyCustomerGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 对端网关ID，例如：cgw-2wqq41m9，可通过[DescribeCustomerGateways](https://cloud.tencent.com/document/api/215/17516)接口查询对端网关。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`

	// 对端网关名称，可任意命名，但不得超过60个字符。
	CustomerGatewayName *string `json:"CustomerGatewayName,omitnil" name:"CustomerGatewayName"`
}

func (r *ModifyCustomerGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerGatewayId")
	delete(f, "CustomerGatewayName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomerGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomerGatewayAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomerGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomerGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyCustomerGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDhcpIpAttributeRequestParams struct {
	// `DhcpIp`唯一`ID`，形如：`dhcpip-9o233uri`。
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`

	// `DhcpIp`名称，可任意命名，但不得超过60个字符。
	DhcpIpName *string `json:"DhcpIpName,omitnil" name:"DhcpIpName"`
}

type ModifyDhcpIpAttributeRequest struct {
	*tchttp.BaseRequest
	
	// `DhcpIp`唯一`ID`，形如：`dhcpip-9o233uri`。
	DhcpIpId *string `json:"DhcpIpId,omitnil" name:"DhcpIpId"`

	// `DhcpIp`名称，可任意命名，但不得超过60个字符。
	DhcpIpName *string `json:"DhcpIpName,omitnil" name:"DhcpIpName"`
}

func (r *ModifyDhcpIpAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDhcpIpAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DhcpIpId")
	delete(f, "DhcpIpName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDhcpIpAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDhcpIpAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDhcpIpAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDhcpIpAttributeResponseParams `json:"Response"`
}

func (r *ModifyDhcpIpAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDhcpIpAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectGatewayAttributeRequestParams struct {
	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 专线网关名称，可任意命名，但不得超过60个字符。
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil" name:"DirectConnectGatewayName"`

	// 云联网路由学习类型，可选值：`BGP`（自动学习）、`STATIC`（静态，即用户配置）。只有云联网类型专线网关且开启了BGP功能才支持修改`CcnRouteType`。
	CcnRouteType *string `json:"CcnRouteType,omitnil" name:"CcnRouteType"`

	// 云联网路由发布模式，可选值：`standard`（标准模式）、`exquisite`（精细模式）。只有云联网类型专线网关才支持修改`ModeType`。
	ModeType *string `json:"ModeType,omitnil" name:"ModeType"`
}

type ModifyDirectConnectGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关唯一`ID`，形如：`dcg-9o233uri`。
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 专线网关名称，可任意命名，但不得超过60个字符。
	DirectConnectGatewayName *string `json:"DirectConnectGatewayName,omitnil" name:"DirectConnectGatewayName"`

	// 云联网路由学习类型，可选值：`BGP`（自动学习）、`STATIC`（静态，即用户配置）。只有云联网类型专线网关且开启了BGP功能才支持修改`CcnRouteType`。
	CcnRouteType *string `json:"CcnRouteType,omitnil" name:"CcnRouteType"`

	// 云联网路由发布模式，可选值：`standard`（标准模式）、`exquisite`（精细模式）。只有云联网类型专线网关才支持修改`ModeType`。
	ModeType *string `json:"ModeType,omitnil" name:"ModeType"`
}

func (r *ModifyDirectConnectGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "DirectConnectGatewayName")
	delete(f, "CcnRouteType")
	delete(f, "ModeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDirectConnectGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDirectConnectGatewayAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDirectConnectGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDirectConnectGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyDirectConnectGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDirectConnectGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlowLogAttributeRequestParams struct {
	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`

	// 私用网络ID或者统一ID，建议使用统一ID，修改云联网流日志属性时可不填，其他流日志类型必填。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志实例名字。
	FlowLogName *string `json:"FlowLogName,omitnil" name:"FlowLogName"`

	// 流日志实例描述。
	FlowLogDescription *string `json:"FlowLogDescription,omitnil" name:"FlowLogDescription"`
}

type ModifyFlowLogAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 流日志唯一ID。
	FlowLogId *string `json:"FlowLogId,omitnil" name:"FlowLogId"`

	// 私用网络ID或者统一ID，建议使用统一ID，修改云联网流日志属性时可不填，其他流日志类型必填。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 流日志实例名字。
	FlowLogName *string `json:"FlowLogName,omitnil" name:"FlowLogName"`

	// 流日志实例描述。
	FlowLogDescription *string `json:"FlowLogDescription,omitnil" name:"FlowLogDescription"`
}

func (r *ModifyFlowLogAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowLogAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FlowLogId")
	delete(f, "VpcId")
	delete(f, "FlowLogName")
	delete(f, "FlowLogDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFlowLogAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFlowLogAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyFlowLogAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFlowLogAttributeResponseParams `json:"Response"`
}

func (r *ModifyFlowLogAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFlowLogAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGatewayFlowQosRequestParams struct {
	// 网关实例ID，目前我们支持的网关实例类型有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 流控带宽值。取值大于0，表示限流到指定的Mbps；取值等于0，表示完全限流；取值为-1，不限流。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 限流的云服务器内网IP。
	IpAddresses []*string `json:"IpAddresses,omitnil" name:"IpAddresses"`
}

type ModifyGatewayFlowQosRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID，目前我们支持的网关实例类型有，
	// 专线网关实例ID，形如，`dcg-ltjahce6`；
	// Nat网关实例ID，形如，`nat-ltjahce6`；
	// VPN网关实例ID，形如，`vpn-ltjahce6`。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 流控带宽值。取值大于0，表示限流到指定的Mbps；取值等于0，表示完全限流；取值为-1，不限流。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 限流的云服务器内网IP。
	IpAddresses []*string `json:"IpAddresses,omitnil" name:"IpAddresses"`
}

func (r *ModifyGatewayFlowQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatewayFlowQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Bandwidth")
	delete(f, "IpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGatewayFlowQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGatewayFlowQosResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyGatewayFlowQosResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGatewayFlowQosResponseParams `json:"Response"`
}

func (r *ModifyGatewayFlowQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGatewayFlowQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHaVipAttributeRequestParams struct {
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`

	// `HAVIP`名称，可任意命名，但不得超过60个字符。
	HaVipName *string `json:"HaVipName,omitnil" name:"HaVipName"`
}

type ModifyHaVipAttributeRequest struct {
	*tchttp.BaseRequest
	
	// `HAVIP`唯一`ID`，形如：`havip-9o233uri`。
	HaVipId *string `json:"HaVipId,omitnil" name:"HaVipId"`

	// `HAVIP`名称，可任意命名，但不得超过60个字符。
	HaVipName *string `json:"HaVipName,omitnil" name:"HaVipName"`
}

func (r *ModifyHaVipAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHaVipAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HaVipId")
	delete(f, "HaVipName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHaVipAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHaVipAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyHaVipAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHaVipAttributeResponseParams `json:"Response"`
}

func (r *ModifyHaVipAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHaVipAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIp6AddressesBandwidthRequestParams struct {
	// 修改的目标带宽，单位Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// IPV6地址。Ip6Addresses和Ip6AddressId必须且只能传一个
	Ip6Addresses []*string `json:"Ip6Addresses,omitnil" name:"Ip6Addresses"`

	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressId必须且只能传一个
	Ip6AddressIds []*string `json:"Ip6AddressIds,omitnil" name:"Ip6AddressIds"`
}

type ModifyIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// 修改的目标带宽，单位Mbps
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// IPV6地址。Ip6Addresses和Ip6AddressId必须且只能传一个
	Ip6Addresses []*string `json:"Ip6Addresses,omitnil" name:"Ip6Addresses"`

	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressId必须且只能传一个
	Ip6AddressIds []*string `json:"Ip6AddressIds,omitnil" name:"Ip6AddressIds"`
}

func (r *ModifyIp6AddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "Ip6Addresses")
	delete(f, "Ip6AddressIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIp6AddressesBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIp6AddressesBandwidthResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIp6AddressesBandwidthResponseParams `json:"Response"`
}

func (r *ModifyIp6AddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIp6RuleRequestParams struct {
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// IPV6转换规则唯一ID，形如rule6-xxxxxxxx
	Ip6RuleId *string `json:"Ip6RuleId,omitnil" name:"Ip6RuleId"`

	// IPV6转换规则修改后的名称
	Ip6RuleName *string `json:"Ip6RuleName,omitnil" name:"Ip6RuleName"`

	// IPV6转换规则修改后的IPV4地址
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPV6转换规则修改后的IPV4端口号
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`
}

type ModifyIp6RuleRequest struct {
	*tchttp.BaseRequest
	
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// IPV6转换规则唯一ID，形如rule6-xxxxxxxx
	Ip6RuleId *string `json:"Ip6RuleId,omitnil" name:"Ip6RuleId"`

	// IPV6转换规则修改后的名称
	Ip6RuleName *string `json:"Ip6RuleName,omitnil" name:"Ip6RuleName"`

	// IPV6转换规则修改后的IPV4地址
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// IPV6转换规则修改后的IPV4端口号
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`
}

func (r *ModifyIp6RuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIp6RuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorId")
	delete(f, "Ip6RuleId")
	delete(f, "Ip6RuleName")
	delete(f, "Vip")
	delete(f, "Vport")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIp6RuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIp6RuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyIp6RuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIp6RuleResponseParams `json:"Response"`
}

func (r *ModifyIp6RuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIp6RuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIp6TranslatorRequestParams struct {
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// IPV6转换实例修改名称
	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitnil" name:"Ip6TranslatorName"`
}

type ModifyIp6TranslatorRequest struct {
	*tchttp.BaseRequest
	
	// IPV6转换实例唯一ID，形如ip6-xxxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// IPV6转换实例修改名称
	Ip6TranslatorName *string `json:"Ip6TranslatorName,omitnil" name:"Ip6TranslatorName"`
}

func (r *ModifyIp6TranslatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIp6TranslatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorId")
	delete(f, "Ip6TranslatorName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIp6TranslatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIp6TranslatorResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyIp6TranslatorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIp6TranslatorResponseParams `json:"Response"`
}

func (r *ModifyIp6TranslatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIp6TranslatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpv6AddressesAttributeRequestParams struct {
	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IPv6地址信息。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`
}

type ModifyIpv6AddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IPv6地址信息。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`
}

func (r *ModifyIpv6AddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpv6AddressesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIpv6AddressesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpv6AddressesAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyIpv6AddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIpv6AddressesAttributeResponseParams `json:"Response"`
}

func (r *ModifyIpv6AddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpv6AddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLocalGatewayRequestParams struct {
	// 本地网关名称。
	LocalGatewayName *string `json:"LocalGatewayName,omitnil" name:"LocalGatewayName"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// 本地网关实例ID。
	LocalGatewayId *string `json:"LocalGatewayId,omitnil" name:"LocalGatewayId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type ModifyLocalGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 本地网关名称。
	LocalGatewayName *string `json:"LocalGatewayName,omitnil" name:"LocalGatewayName"`

	// CDC实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// 本地网关实例ID。
	LocalGatewayId *string `json:"LocalGatewayId,omitnil" name:"LocalGatewayId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

func (r *ModifyLocalGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LocalGatewayName")
	delete(f, "CdcId")
	delete(f, "LocalGatewayId")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLocalGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLocalGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyLocalGatewayResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLocalGatewayResponseParams `json:"Response"`
}

func (r *ModifyLocalGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLocalGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayAttributeRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的名称，形如：`test_nat`。
	NatGatewayName *string `json:"NatGatewayName,omitnil" name:"NatGatewayName"`

	// NAT网关最大外网出带宽(单位:Mbps)。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 是否修改NAT网关绑定的安全组。
	ModifySecurityGroup *bool `json:"ModifySecurityGroup,omitnil" name:"ModifySecurityGroup"`

	// NAT网关绑定的安全组列表，最终状态，空列表表示删除所有安全组，形如: `['sg-1n232323', 'sg-o4242424']`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type ModifyNatGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的名称，形如：`test_nat`。
	NatGatewayName *string `json:"NatGatewayName,omitnil" name:"NatGatewayName"`

	// NAT网关最大外网出带宽(单位:Mbps)。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 是否修改NAT网关绑定的安全组。
	ModifySecurityGroup *bool `json:"ModifySecurityGroup,omitnil" name:"ModifySecurityGroup"`

	// NAT网关绑定的安全组列表，最终状态，空列表表示删除所有安全组，形如: `['sg-1n232323', 'sg-o4242424']`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *ModifyNatGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "NatGatewayName")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "ModifySecurityGroup")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNatGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyNatGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayDestinationIpPortTranslationNatRuleRequestParams struct {
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 源NAT网关的端口转换规则。
	SourceNatRule *DestinationIpPortTranslationNatRule `json:"SourceNatRule,omitnil" name:"SourceNatRule"`

	// 目的NAT网关的端口转换规则。
	DestinationNatRule *DestinationIpPortTranslationNatRule `json:"DestinationNatRule,omitnil" name:"DestinationNatRule"`
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df45454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 源NAT网关的端口转换规则。
	SourceNatRule *DestinationIpPortTranslationNatRule `json:"SourceNatRule,omitnil" name:"SourceNatRule"`

	// 目的NAT网关的端口转换规则。
	DestinationNatRule *DestinationIpPortTranslationNatRule `json:"DestinationNatRule,omitnil" name:"DestinationNatRule"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "SourceNatRule")
	delete(f, "DestinationNatRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatGatewayDestinationIpPortTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewayDestinationIpPortTranslationNatRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponseParams `json:"Response"`
}

func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewayDestinationIpPortTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewaySourceIpTranslationNatRuleRequestParams struct {
	// NAT网关的ID，形如：`nat-df453454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的SNAT转换规则。
	SourceIpTranslationNatRule *SourceIpTranslationNatRule `json:"SourceIpTranslationNatRule,omitnil" name:"SourceIpTranslationNatRule"`
}

type ModifyNatGatewaySourceIpTranslationNatRuleRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关的ID，形如：`nat-df453454`。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的SNAT转换规则。
	SourceIpTranslationNatRule *SourceIpTranslationNatRule `json:"SourceIpTranslationNatRule,omitnil" name:"SourceIpTranslationNatRule"`
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewaySourceIpTranslationNatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "SourceIpTranslationNatRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatGatewaySourceIpTranslationNatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatGatewaySourceIpTranslationNatRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNatGatewaySourceIpTranslationNatRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatGatewaySourceIpTranslationNatRuleResponseParams `json:"Response"`
}

func (r *ModifyNatGatewaySourceIpTranslationNatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatGatewaySourceIpTranslationNatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetDetectRequestParams struct {
	// 网络探测实例`ID`。形如：`netd-12345678`
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`

	// 网络探测名称，最大长度不能超过60个字节。
	NetDetectName *string `json:"NetDetectName,omitnil" name:"NetDetectName"`

	// 探测目的IPv4地址数组，最多两个。
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云服务器；
	// CCN：云联网网关；
	// NONEXTHOP：无下一跳；
	NextHopType *string `json:"NextHopType,omitnil" name:"NextHopType"`

	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云服务器IPv4地址，形如：10.0.0.12；
	// 下一跳类型为CCN，取值云联网ID，形如：ccn-12345678；
	// 下一跳类型为NONEXTHOP，指定网络探测为无下一跳的网络探测；
	NextHopDestination *string `json:"NextHopDestination,omitnil" name:"NextHopDestination"`

	// 网络探测描述。
	NetDetectDescription *string `json:"NetDetectDescription,omitnil" name:"NetDetectDescription"`
}

type ModifyNetDetectRequest struct {
	*tchttp.BaseRequest
	
	// 网络探测实例`ID`。形如：`netd-12345678`
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`

	// 网络探测名称，最大长度不能超过60个字节。
	NetDetectName *string `json:"NetDetectName,omitnil" name:"NetDetectName"`

	// 探测目的IPv4地址数组，最多两个。
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云服务器；
	// CCN：云联网网关；
	// NONEXTHOP：无下一跳；
	NextHopType *string `json:"NextHopType,omitnil" name:"NextHopType"`

	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云服务器IPv4地址，形如：10.0.0.12；
	// 下一跳类型为CCN，取值云联网ID，形如：ccn-12345678；
	// 下一跳类型为NONEXTHOP，指定网络探测为无下一跳的网络探测；
	NextHopDestination *string `json:"NextHopDestination,omitnil" name:"NextHopDestination"`

	// 网络探测描述。
	NetDetectDescription *string `json:"NetDetectDescription,omitnil" name:"NetDetectDescription"`
}

func (r *ModifyNetDetectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetDetectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetDetectId")
	delete(f, "NetDetectName")
	delete(f, "DetectDestinationIp")
	delete(f, "NextHopType")
	delete(f, "NextHopDestination")
	delete(f, "NetDetectDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetDetectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetDetectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNetDetectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetDetectResponseParams `json:"Response"`
}

func (r *ModifyNetDetectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetDetectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclAttributeRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络ACL名称，最大长度不能超过60个字节。
	NetworkAclName *string `json:"NetworkAclName,omitnil" name:"NetworkAclName"`
}

type ModifyNetworkAclAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络ACL名称，最大长度不能超过60个字节。
	NetworkAclName *string `json:"NetworkAclName,omitnil" name:"NetworkAclName"`
}

func (r *ModifyNetworkAclAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkAclAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNetworkAclAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkAclAttributeResponseParams `json:"Response"`
}

func (r *ModifyNetworkAclAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclEntriesRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络ACL规则集。NetworkAclEntrySet和NetworkAclQuintupleSet只能输入一个。
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil" name:"NetworkAclEntrySet"`

	// 网络ACL五元组规则集。NetworkAclEntrySet和NetworkAclQuintupleSet只能输入一个。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`

	// 三元组的增量更新。该接口的默认语义为全量覆盖。当需要实现增量更新语义时，设置该参数为True。
	EnableUpdateAclEntries *bool `json:"EnableUpdateAclEntries,omitnil" name:"EnableUpdateAclEntries"`
}

type ModifyNetworkAclEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络ACL规则集。NetworkAclEntrySet和NetworkAclQuintupleSet只能输入一个。
	NetworkAclEntrySet *NetworkAclEntrySet `json:"NetworkAclEntrySet,omitnil" name:"NetworkAclEntrySet"`

	// 网络ACL五元组规则集。NetworkAclEntrySet和NetworkAclQuintupleSet只能输入一个。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`

	// 三元组的增量更新。该接口的默认语义为全量覆盖。当需要实现增量更新语义时，设置该参数为True。
	EnableUpdateAclEntries *bool `json:"EnableUpdateAclEntries,omitnil" name:"EnableUpdateAclEntries"`
}

func (r *ModifyNetworkAclEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclEntrySet")
	delete(f, "NetworkAclQuintupleSet")
	delete(f, "EnableUpdateAclEntries")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkAclEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclEntriesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNetworkAclEntriesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkAclEntriesResponseParams `json:"Response"`
}

func (r *ModifyNetworkAclEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclQuintupleEntriesRequestParams struct {
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络五元组ACL规则集。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`
}

type ModifyNetworkAclQuintupleEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 网络ACL实例ID。例如：acl-12345678。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络五元组ACL规则集。
	NetworkAclQuintupleSet *NetworkAclQuintupleEntries `json:"NetworkAclQuintupleSet,omitnil" name:"NetworkAclQuintupleSet"`
}

func (r *ModifyNetworkAclQuintupleEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclQuintupleEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkAclId")
	delete(f, "NetworkAclQuintupleSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkAclQuintupleEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAclQuintupleEntriesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNetworkAclQuintupleEntriesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkAclQuintupleEntriesResponseParams `json:"Response"`
}

func (r *ModifyNetworkAclQuintupleEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAclQuintupleEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkInterfaceAttributeRequestParams struct {
	// 弹性网卡实例ID，例如：eni-pxir56ns。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil" name:"NetworkInterfaceName"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil" name:"NetworkInterfaceDescription"`

	// 指定绑定的安全组，例如:['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 网卡trunking模式设置，Enable-开启，Disable--关闭，默认关闭。
	TrunkingFlag *string `json:"TrunkingFlag,omitnil" name:"TrunkingFlag"`
}

type ModifyNetworkInterfaceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-pxir56ns。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 弹性网卡名称，最大长度不能超过60个字节。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil" name:"NetworkInterfaceName"`

	// 弹性网卡描述，可任意命名，但不得超过60个字符。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil" name:"NetworkInterfaceDescription"`

	// 指定绑定的安全组，例如:['sg-1dd51d']。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`

	// 网卡trunking模式设置，Enable-开启，Disable--关闭，默认关闭。
	TrunkingFlag *string `json:"TrunkingFlag,omitnil" name:"TrunkingFlag"`
}

func (r *ModifyNetworkInterfaceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkInterfaceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "NetworkInterfaceName")
	delete(f, "NetworkInterfaceDescription")
	delete(f, "SecurityGroupIds")
	delete(f, "TrunkingFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkInterfaceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkInterfaceAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNetworkInterfaceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkInterfaceAttributeResponseParams `json:"Response"`
}

func (r *ModifyNetworkInterfaceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkInterfaceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkInterfaceQosRequestParams struct {
	// 弹性网卡ID，支持批量修改。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 服务质量，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// DirectSend端口范围最大值。
	DirectSendMaxPort *uint64 `json:"DirectSendMaxPort,omitnil" name:"DirectSendMaxPort"`
}

type ModifyNetworkInterfaceQosRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡ID，支持批量修改。
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitnil" name:"NetworkInterfaceIds"`

	// 服务质量，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// DirectSend端口范围最大值。
	DirectSendMaxPort *uint64 `json:"DirectSendMaxPort,omitnil" name:"DirectSendMaxPort"`
}

func (r *ModifyNetworkInterfaceQosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkInterfaceQosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceIds")
	delete(f, "QosLevel")
	delete(f, "DirectSendMaxPort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkInterfaceQosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkInterfaceQosResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNetworkInterfaceQosResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkInterfaceQosResponseParams `json:"Response"`
}

func (r *ModifyNetworkInterfaceQosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkInterfaceQosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`
}

type ModifyPrivateIpAddressesAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IP信息。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`
}

func (r *ModifyPrivateIpAddressesAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateIpAddressesAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPrivateIpAddressesAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPrivateIpAddressesAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyPrivateIpAddressesAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPrivateIpAddressesAttributeResponseParams `json:"Response"`
}

func (r *ModifyPrivateIpAddressesAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPrivateIpAddressesAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableAttributeRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`
}

type ModifyRouteTableAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`
}

func (r *ModifyRouteTableAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRouteTableAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRouteTableAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRouteTableAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRouteTableAttributeResponseParams `json:"Response"`
}

func (r *ModifyRouteTableAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRouteTableAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupAttributeRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`
}

type ModifySecurityGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 安全组备注，最多100个字符。
	GroupDescription *string `json:"GroupDescription,omitnil" name:"GroupDescription"`
}

func (r *ModifySecurityGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "GroupName")
	delete(f, "GroupDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySecurityGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合。 SecurityGroupPolicySet对象必须同时指定新的出（Egress）入（Ingress）站规则。 SecurityGroupPolicy对象不支持自定义索引（PolicyIndex）。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`

	// 排序安全组标识，默认值为False。当SortPolicys为False时，不改变安全组规则排序；当SortPolicys为True时，系统将严格按照SecurityGroupPolicySet参数传入的安全组规则及顺序进行重置，考虑到人为输入参数可能存在遗漏风险，建议通过控制台对安全组规则进行排序。
	SortPolicys *bool `json:"SortPolicys,omitnil" name:"SortPolicys"`
}

type ModifySecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合。 SecurityGroupPolicySet对象必须同时指定新的出（Egress）入（Ingress）站规则。 SecurityGroupPolicy对象不支持自定义索引（PolicyIndex）。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`

	// 排序安全组标识，默认值为False。当SortPolicys为False时，不改变安全组规则排序；当SortPolicys为True时，系统将严格按照SecurityGroupPolicySet参数传入的安全组规则及顺序进行重置，考虑到人为输入参数可能存在遗漏风险，建议通过控制台对安全组规则进行排序。
	SortPolicys *bool `json:"SortPolicys,omitnil" name:"SortPolicys"`
}

func (r *ModifySecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	delete(f, "SortPolicys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceTemplateAttributeRequestParams struct {
	// 协议端口模板实例ID，例如：ppm-529nwwj8。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// 协议端口模板名称。
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil" name:"ServiceTemplateName"`

	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	Services []*string `json:"Services,omitnil" name:"Services"`

	// 支持添加备注的协议端口信息，支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil" name:"ServicesExtra"`
}

type ModifyServiceTemplateAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 协议端口模板实例ID，例如：ppm-529nwwj8。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// 协议端口模板名称。
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil" name:"ServiceTemplateName"`

	// 支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	Services []*string `json:"Services,omitnil" name:"Services"`

	// 支持添加备注的协议端口信息，支持单个端口、多个端口、连续端口及所有端口，协议支持：TCP、UDP、ICMP、GRE 协议。
	ServicesExtra []*ServicesInfo `json:"ServicesExtra,omitnil" name:"ServicesExtra"`
}

func (r *ModifyServiceTemplateAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateId")
	delete(f, "ServiceTemplateName")
	delete(f, "Services")
	delete(f, "ServicesExtra")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceTemplateAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceTemplateAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyServiceTemplateAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceTemplateAttributeResponseParams `json:"Response"`
}

func (r *ModifyServiceTemplateAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceTemplateGroupAttributeRequestParams struct {
	// 协议端口模板集合实例ID，例如：ppmg-ei8hfd9a。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil" name:"ServiceTemplateGroupId"`

	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil" name:"ServiceTemplateGroupName"`

	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil" name:"ServiceTemplateIds"`
}

type ModifyServiceTemplateGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 协议端口模板集合实例ID，例如：ppmg-ei8hfd9a。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil" name:"ServiceTemplateGroupId"`

	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil" name:"ServiceTemplateGroupName"`

	// 协议端口模板实例ID，例如：ppm-4dw6agho。
	ServiceTemplateIds []*string `json:"ServiceTemplateIds,omitnil" name:"ServiceTemplateIds"`
}

func (r *ModifyServiceTemplateGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceTemplateGroupId")
	delete(f, "ServiceTemplateGroupName")
	delete(f, "ServiceTemplateIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceTemplateGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceTemplateGroupAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyServiceTemplateGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceTemplateGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyServiceTemplateGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceTemplateGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotPoliciesRequestParams struct {
	// 快照策略修改信息。
	SnapshotPoliciesInfo []*BatchModifySnapshotPolicy `json:"SnapshotPoliciesInfo,omitnil" name:"SnapshotPoliciesInfo"`
}

type ModifySnapshotPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略修改信息。
	SnapshotPoliciesInfo []*BatchModifySnapshotPolicy `json:"SnapshotPoliciesInfo,omitnil" name:"SnapshotPoliciesInfo"`
}

func (r *ModifySnapshotPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPoliciesInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySnapshotPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySnapshotPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySnapshotPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySnapshotPoliciesResponseParams `json:"Response"`
}

func (r *ModifySnapshotPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySnapshotPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeRequestParams struct {
	// 子网实例ID。形如：subnet-pxir56ns。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 子网是否开启广播。
	EnableBroadcast *string `json:"EnableBroadcast,omitnil" name:"EnableBroadcast"`
}

type ModifySubnetAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID。形如：subnet-pxir56ns。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网名称，最大长度不能超过60个字节。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 子网是否开启广播。
	EnableBroadcast *string `json:"EnableBroadcast,omitnil" name:"EnableBroadcast"`
}

func (r *ModifySubnetAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "SubnetName")
	delete(f, "EnableBroadcast")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubnetAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubnetAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySubnetAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubnetAttributeResponseParams `json:"Response"`
}

func (r *ModifySubnetAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubnetAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTemplateMemberRequestParams struct {
	// 参数模板实例ID，支持IP地址、协议端口、IP地址组、协议端口组四种参数模板的实例ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 需要修改的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致，修改顺序与TemplateMember参数顺序一一对应，入参长度需要与TemplateMember参数保持一致。
	OriginalTemplateMember []*MemberInfo `json:"OriginalTemplateMember,omitnil" name:"OriginalTemplateMember"`

	// 新的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致，修改顺序与OriginalTemplateMember参数顺序一一对应，入参长度需要与OriginalTemplateMember参数保持一致。
	TemplateMember []*MemberInfo `json:"TemplateMember,omitnil" name:"TemplateMember"`
}

type ModifyTemplateMemberRequest struct {
	*tchttp.BaseRequest
	
	// 参数模板实例ID，支持IP地址、协议端口、IP地址组、协议端口组四种参数模板的实例ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 需要修改的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致，修改顺序与TemplateMember参数顺序一一对应，入参长度需要与TemplateMember参数保持一致。
	OriginalTemplateMember []*MemberInfo `json:"OriginalTemplateMember,omitnil" name:"OriginalTemplateMember"`

	// 新的参数模板成员信息，支持IP地址、协议端口、IP地址组、协议端口组四种类型，类型需要与TemplateId参数类型一致，修改顺序与OriginalTemplateMember参数顺序一一对应，入参长度需要与OriginalTemplateMember参数保持一致。
	TemplateMember []*MemberInfo `json:"TemplateMember,omitnil" name:"TemplateMember"`
}

func (r *ModifyTemplateMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTemplateMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "OriginalTemplateMember")
	delete(f, "TemplateMember")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTemplateMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTemplateMemberResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyTemplateMemberResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTemplateMemberResponseParams `json:"Response"`
}

func (r *ModifyTemplateMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTemplateMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeRequestParams struct {
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络名称，可任意命名，但不得超过60个字符。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 是否开启组播。true: 开启, false: 关闭。
	EnableMulticast *string `json:"EnableMulticast,omitnil" name:"EnableMulticast"`

	// DNS地址，最多支持4个，第1个默认为主，其余为备。
	DnsServers []*string `json:"DnsServers,omitnil" name:"DnsServers"`

	// 域名。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 发布cdc 子网到云联网的开关。true: 发布, false: 不发布。
	EnableCdcPublish *bool `json:"EnableCdcPublish,omitnil" name:"EnableCdcPublish"`
}

type ModifyVpcAttributeRequest struct {
	*tchttp.BaseRequest
	
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络名称，可任意命名，但不得超过60个字符。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 是否开启组播。true: 开启, false: 关闭。
	EnableMulticast *string `json:"EnableMulticast,omitnil" name:"EnableMulticast"`

	// DNS地址，最多支持4个，第1个默认为主，其余为备。
	DnsServers []*string `json:"DnsServers,omitnil" name:"DnsServers"`

	// 域名。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 发布cdc 子网到云联网的开关。true: 发布, false: 不发布。
	EnableCdcPublish *bool `json:"EnableCdcPublish,omitnil" name:"EnableCdcPublish"`
}

func (r *ModifyVpcAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "VpcName")
	delete(f, "EnableMulticast")
	delete(f, "DnsServers")
	delete(f, "DomainName")
	delete(f, "EnableCdcPublish")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpcAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpcAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointAttributeRequestParams struct {
	// 终端节点ID。
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`

	// 终端节点名称。
	EndPointName *string `json:"EndPointName,omitnil" name:"EndPointName"`

	// 安全组ID列表。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

type ModifyVpcEndPointAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点ID。
	EndPointId *string `json:"EndPointId,omitnil" name:"EndPointId"`

	// 终端节点名称。
	EndPointName *string `json:"EndPointName,omitnil" name:"EndPointName"`

	// 安全组ID列表。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil" name:"SecurityGroupIds"`
}

func (r *ModifyVpcEndPointAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointId")
	delete(f, "EndPointName")
	delete(f, "SecurityGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcEndPointAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpcEndPointAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcEndPointAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpcEndPointAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceAttributeRequestParams struct {
	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// VPCID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 终端节点服务名称。
	EndPointServiceName *string `json:"EndPointServiceName,omitnil" name:"EndPointServiceName"`

	// 是否自动接受终端节点的连接请求。<ui><li>true：自动接受<li>false：不自动接受</ul>
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil" name:"AutoAcceptFlag"`

	// 后端服务的ID，比如lb-xxx。
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil" name:"ServiceInstanceId"`
}

type ModifyVpcEndPointServiceAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// VPCID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 终端节点服务名称。
	EndPointServiceName *string `json:"EndPointServiceName,omitnil" name:"EndPointServiceName"`

	// 是否自动接受终端节点的连接请求。<ui><li>true：自动接受<li>false：不自动接受</ul>
	AutoAcceptFlag *bool `json:"AutoAcceptFlag,omitnil" name:"AutoAcceptFlag"`

	// 后端服务的ID，比如lb-xxx。
	ServiceInstanceId *string `json:"ServiceInstanceId,omitnil" name:"ServiceInstanceId"`
}

func (r *ModifyVpcEndPointServiceAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndPointServiceId")
	delete(f, "VpcId")
	delete(f, "EndPointServiceName")
	delete(f, "AutoAcceptFlag")
	delete(f, "ServiceInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcEndPointServiceAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpcEndPointServiceAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcEndPointServiceAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpcEndPointServiceAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceWhiteListRequestParams struct {
	// 用户UIN。
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 白名单描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ModifyVpcEndPointServiceWhiteListRequest struct {
	*tchttp.BaseRequest
	
	// 用户UIN。
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`

	// 白名单描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *ModifyVpcEndPointServiceWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserUin")
	delete(f, "EndPointServiceId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcEndPointServiceWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcEndPointServiceWhiteListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpcEndPointServiceWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcEndPointServiceWhiteListResponseParams `json:"Response"`
}

func (r *ModifyVpcEndPointServiceWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcEndPointServiceWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcPeeringConnectionRequestParams struct {
	// 对等连接ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`

	// 对等连接名称。
	PeeringConnectionName *string `json:"PeeringConnectionName,omitnil" name:"PeeringConnectionName"`

	// 带宽上限，单位Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 计费模式，日峰值POSTPAID_BY_DAY_MAX，月95 POSTPAID_BY_MONTH_95。
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`
}

type ModifyVpcPeeringConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 对等连接ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`

	// 对等连接名称。
	PeeringConnectionName *string `json:"PeeringConnectionName,omitnil" name:"PeeringConnectionName"`

	// 带宽上限，单位Mbps。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 计费模式，日峰值POSTPAID_BY_DAY_MAX，月95 POSTPAID_BY_MONTH_95。
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`
}

func (r *ModifyVpcPeeringConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PeeringConnectionId")
	delete(f, "PeeringConnectionName")
	delete(f, "Bandwidth")
	delete(f, "ChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcPeeringConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpcPeeringConnectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcPeeringConnectionResponseParams `json:"Response"`
}

func (r *ModifyVpcPeeringConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnConnectionAttributeRequestParams struct {
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`

	// VPN通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName,omitnil" name:"VpnConnectionName"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitnil" name:"PreShareKey"`

	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段，172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil" name:"SecurityPolicyDatabases"`

	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil" name:"IKEOptionsSpecification"`

	// IPSec配置，腾讯云提供IPSec安全会话设置。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil" name:"IPSECOptionsSpecification"`

	// 是否启用通道健康检查，默认为False。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil" name:"EnableHealthCheck"`

	// 本端通道探测IP。
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil" name:"HealthCheckLocalIp"`

	// 对端通道探测IP。
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil" name:"HealthCheckRemoteIp"`

	// 协商类型，默认为active（主动协商）。可选值：active（主动协商），passive（被动协商），flowTrigger（流量协商）
	NegotiationType *string `json:"NegotiationType,omitnil" name:"NegotiationType"`

	// DPD探测开关。默认为0，表示关闭DPD探测。可选值：0（关闭），1（开启）
	DpdEnable *int64 `json:"DpdEnable,omitnil" name:"DpdEnable"`

	// DPD超时时间。即探测确认对端不存在需要的时间。dpdEnable为1（开启）时有效。默认30，单位为秒
	DpdTimeout *string `json:"DpdTimeout,omitnil" name:"DpdTimeout"`

	// DPD超时后的动作。默认为clear。dpdEnable为1（开启）时有效。可取值为clear（断开）和restart（重试）
	DpdAction *string `json:"DpdAction,omitnil" name:"DpdAction"`

	// 对端网关ID，4.0及以上网关下的通道支持更新。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`
}

type ModifyVpnConnectionAttributeRequest struct {
	*tchttp.BaseRequest
	
	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`

	// VPN通道名称，可任意命名，但不得超过60个字符。
	VpnConnectionName *string `json:"VpnConnectionName,omitnil" name:"VpnConnectionName"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitnil" name:"PreShareKey"`

	// SPD策略组，例如：{"10.0.0.5/24":["172.123.10.5/16"]}，10.0.0.5/24是vpc内网段，172.123.10.5/16是IDC网段。用户指定VPC内哪些网段可以和您IDC中哪些网段通信。
	SecurityPolicyDatabases []*SecurityPolicyDatabase `json:"SecurityPolicyDatabases,omitnil" name:"SecurityPolicyDatabases"`

	// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自我保护机制，用户配置网络安全协议。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil" name:"IKEOptionsSpecification"`

	// IPSec配置，腾讯云提供IPSec安全会话设置。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil" name:"IPSECOptionsSpecification"`

	// 是否启用通道健康检查，默认为False。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil" name:"EnableHealthCheck"`

	// 本端通道探测IP。
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil" name:"HealthCheckLocalIp"`

	// 对端通道探测IP。
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil" name:"HealthCheckRemoteIp"`

	// 协商类型，默认为active（主动协商）。可选值：active（主动协商），passive（被动协商），flowTrigger（流量协商）
	NegotiationType *string `json:"NegotiationType,omitnil" name:"NegotiationType"`

	// DPD探测开关。默认为0，表示关闭DPD探测。可选值：0（关闭），1（开启）
	DpdEnable *int64 `json:"DpdEnable,omitnil" name:"DpdEnable"`

	// DPD超时时间。即探测确认对端不存在需要的时间。dpdEnable为1（开启）时有效。默认30，单位为秒
	DpdTimeout *string `json:"DpdTimeout,omitnil" name:"DpdTimeout"`

	// DPD超时后的动作。默认为clear。dpdEnable为1（开启）时有效。可取值为clear（断开）和restart（重试）
	DpdAction *string `json:"DpdAction,omitnil" name:"DpdAction"`

	// 对端网关ID，4.0及以上网关下的通道支持更新。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`
}

func (r *ModifyVpnConnectionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnConnectionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnConnectionId")
	delete(f, "VpnConnectionName")
	delete(f, "PreShareKey")
	delete(f, "SecurityPolicyDatabases")
	delete(f, "IKEOptionsSpecification")
	delete(f, "IPSECOptionsSpecification")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheckLocalIp")
	delete(f, "HealthCheckRemoteIp")
	delete(f, "NegotiationType")
	delete(f, "DpdEnable")
	delete(f, "DpdTimeout")
	delete(f, "DpdAction")
	delete(f, "CustomerGatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnConnectionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnConnectionAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpnConnectionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnConnectionAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpnConnectionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnConnectionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayAttributeRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName,omitnil" name:"VpnGatewayName"`

	// VPN网关计费模式，目前只支持预付费（即包年包月）到后付费（即按量计费）的转换。即参数只支持：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`
}

type ModifyVpnGatewayAttributeRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN网关名称，最大长度不能超过60个字节。
	VpnGatewayName *string `json:"VpnGatewayName,omitnil" name:"VpnGatewayName"`

	// VPN网关计费模式，目前只支持预付费（即包年包月）到后付费（即按量计费）的转换。即参数只支持：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`
}

func (r *ModifyVpnGatewayAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnGatewayName")
	delete(f, "InstanceChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpnGatewayAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnGatewayAttributeResponseParams `json:"Response"`
}

func (r *ModifyVpnGatewayAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayCcnRoutesRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 云联网路由（IDC网段）列表。
	Routes []*VpngwCcnRoutes `json:"Routes,omitnil" name:"Routes"`
}

type ModifyVpnGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 云联网路由（IDC网段）列表。
	Routes []*VpngwCcnRoutes `json:"Routes,omitnil" name:"Routes"`
}

func (r *ModifyVpnGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayCcnRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpnGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *ModifyVpnGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayRoutesRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 路由修改参数。
	Routes []*VpnGatewayRouteModify `json:"Routes,omitnil" name:"Routes"`
}

type ModifyVpnGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 路由修改参数。
	Routes []*VpnGatewayRouteModify `json:"Routes,omitnil" name:"Routes"`
}

func (r *ModifyVpnGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpnGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpnGatewayRoutesResponseParams struct {
	// VPN路由信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Routes []*VpnGatewayRoute `json:"Routes,omitnil" name:"Routes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyVpnGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpnGatewayRoutesResponseParams `json:"Response"`
}

func (r *ModifyVpnGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpnGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatDirectConnectGatewayRoute struct {
	// 子网的 `IPv4` `CIDR`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil" name:"DestinationCidrBlock"`

	// 下一跳网关的类型，目前此接口支持的类型有：
	// DIRECTCONNECT：专线网关
	GatewayType *string `json:"GatewayType,omitnil" name:"GatewayType"`

	// 下一跳网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由的创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 路由的更新时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type NatGateway struct {
	// NAT网关的ID。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关的名称。
	NatGatewayName *string `json:"NatGatewayName,omitnil" name:"NatGatewayName"`

	// NAT网关创建的时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// NAT网关的状态。
	//  'PENDING'：生产中，'DELETING'：删除中，'AVAILABLE'：运行中，'UPDATING'：升级中，
	// ‘FAILED’：失败。
	State *string `json:"State,omitnil" name:"State"`

	// 网关最大外网出带宽(单位:Mbps)。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 网关并发连接上限。
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil" name:"MaxConcurrentConnection"`

	// 绑定NAT网关的公网IP对象数组。
	PublicIpAddressSet []*NatGatewayAddress `json:"PublicIpAddressSet,omitnil" name:"PublicIpAddressSet"`

	// NAT网关网络状态。“AVAILABLE”:运行中, “UNAVAILABLE”:不可用, “INSUFFICIENT”:欠费停服。
	NetworkState *string `json:"NetworkState,omitnil" name:"NetworkState"`

	// NAT网关的端口转发规则。
	DestinationIpPortTranslationNatRuleSet []*DestinationIpPortTranslationNatRule `json:"DestinationIpPortTranslationNatRuleSet,omitnil" name:"DestinationIpPortTranslationNatRuleSet"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关所在的可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 绑定的专线网关ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DirectConnectGatewayIds []*string `json:"DirectConnectGatewayIds,omitnil" name:"DirectConnectGatewayIds"`

	// 所属子网ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 标签键值对。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// NAT网关绑定的安全组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupSet []*string `json:"SecurityGroupSet,omitnil" name:"SecurityGroupSet"`

	// NAT网关的SNAT转发规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceIpTranslationNatRuleSet []*SourceIpTranslationNatRule `json:"SourceIpTranslationNatRuleSet,omitnil" name:"SourceIpTranslationNatRuleSet"`

	// 是否独享型NAT。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExclusive *bool `json:"IsExclusive,omitnil" name:"IsExclusive"`

	// 独享型NAT所在的网关集群的带宽(单位:Mbps)，当IsExclusive为false时无此字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExclusiveGatewayBandwidth *uint64 `json:"ExclusiveGatewayBandwidth,omitnil" name:"ExclusiveGatewayBandwidth"`

	// NAT网关是否被封禁。“NORMAL”：未被封禁，“RESTRICTED”：已被封禁。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestrictState *string `json:"RestrictState,omitnil" name:"RestrictState"`

	// NAT网关大版本号，传统型=1，标准型=2
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatProductVersion *uint64 `json:"NatProductVersion,omitnil" name:"NatProductVersion"`

	// 是否启用根据目的网段选择SNAT使用的EIP功能	
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartScheduleMode *bool `json:"SmartScheduleMode,omitnil" name:"SmartScheduleMode"`

	// NAT实例归属的专属集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DedicatedClusterId *string `json:"DedicatedClusterId,omitnil" name:"DedicatedClusterId"`
}

type NatGatewayAddress struct {
	// 弹性公网IP（EIP）的唯一 ID，形如：`eip-11112222`。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 外网IP地址，形如：`123.121.34.33`。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil" name:"PublicIpAddress"`

	// 资源封堵状态。true表示弹性ip处于封堵状态，false表示弹性ip处于未封堵状态。
	IsBlocked *bool `json:"IsBlocked,omitnil" name:"IsBlocked"`

	// 资源封堵类型。NORMAL表示未封禁，SECURITY表示安全封禁，USER表示用户封禁，OTHER表示其他封禁，多个原因封禁时用&连接，比如：SECURITY&USER&OTHER。
	BlockType *string `json:"BlockType,omitnil" name:"BlockType"`
}

type NatGatewayDestinationIpPortTranslationNatRule struct {
	// 网络协议，可选值：`TCP`、`UDP`。
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 弹性IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil" name:"PublicIpAddress"`

	// 公网端口。
	PublicPort *uint64 `json:"PublicPort,omitnil" name:"PublicPort"`

	// 内网地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 内网端口。
	PrivatePort *uint64 `json:"PrivatePort,omitnil" name:"PrivatePort"`

	// NAT网关转发规则描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// NAT网关的ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 私有网络VPC的ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关转发规则创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type NetDetect struct {
	// `VPC`实例`ID`。形如：`vpc-12345678`
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `VPC`实例名称。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 子网实例ID。形如：subnet-12345678。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网实例名称。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 网络探测实例ID。形如：netd-12345678。
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`

	// 网络探测名称，最大长度不能超过60个字节。
	NetDetectName *string `json:"NetDetectName,omitnil" name:"NetDetectName"`

	// 探测目的IPv4地址数组，最多两个。
	DetectDestinationIp []*string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 系统自动分配的探测源IPv4数组。长度为2。
	DetectSourceIp []*string `json:"DetectSourceIp,omitnil" name:"DetectSourceIp"`

	// 下一跳类型，目前我们支持的类型有：
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// NAT：NAT网关；
	// NORMAL_CVM：普通云服务器；
	// CCN：云联网网关；
	// NONEXTHOP：无下一跳；
	NextHopType *string `json:"NextHopType,omitnil" name:"NextHopType"`

	// 下一跳目的网关，取值与“下一跳类型”相关：
	// 下一跳类型为VPN，取值VPN网关ID，形如：vpngw-12345678；
	// 下一跳类型为DIRECTCONNECT，取值专线网关ID，形如：dcg-12345678；
	// 下一跳类型为PEERCONNECTION，取值对等连接ID，形如：pcx-12345678；
	// 下一跳类型为NAT，取值Nat网关，形如：nat-12345678；
	// 下一跳类型为NORMAL_CVM，取值云服务器IPv4地址，形如：10.0.0.12；
	// 下一跳类型为CCN，取值云联网ID，形如：ccn-12345678；
	// 下一跳类型为NONEXTHOP，指定网络探测为无下一跳的网络探测，添加和修改时，不需要指定值，查询时值为空字符串；
	NextHopDestination *string `json:"NextHopDestination,omitnil" name:"NextHopDestination"`

	// 下一跳网关名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextHopName *string `json:"NextHopName,omitnil" name:"NextHopName"`

	// 网络探测描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetDetectDescription *string `json:"NetDetectDescription,omitnil" name:"NetDetectDescription"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type NetDetectIpState struct {
	// 探测目的IPv4地址。
	DetectDestinationIp *string `json:"DetectDestinationIp,omitnil" name:"DetectDestinationIp"`

	// 探测结果。
	// 0：成功；
	// -1：查询不到路由丢包；
	// -2：外出ACL丢包；
	// -3：IN ACL丢包；
	// -4：其他错误；
	State *int64 `json:"State,omitnil" name:"State"`

	// 时延，单位毫秒
	Delay *uint64 `json:"Delay,omitnil" name:"Delay"`

	// 丢包率
	PacketLossRate *uint64 `json:"PacketLossRate,omitnil" name:"PacketLossRate"`
}

type NetDetectState struct {
	// 网络探测实例ID。形如：netd-12345678。
	NetDetectId *string `json:"NetDetectId,omitnil" name:"NetDetectId"`

	// 网络探测目的IP验证结果对象数组。
	NetDetectIpStateSet []*NetDetectIpState `json:"NetDetectIpStateSet,omitnil" name:"NetDetectIpStateSet"`
}

type NetworkAcl struct {
	// `VPC`实例`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 网络ACL实例`ID`。
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 网络ACL名称，最大长度为60。
	NetworkAclName *string `json:"NetworkAclName,omitnil" name:"NetworkAclName"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 网络ACL关联的子网数组。
	SubnetSet []*Subnet `json:"SubnetSet,omitnil" name:"SubnetSet"`

	// 该参数仅对三元组ACL有效，网络ACl入站规则。
	IngressEntries []*NetworkAclEntry `json:"IngressEntries,omitnil" name:"IngressEntries"`

	// 该参数仅对三元组ACL有效，网络ACL出站规则。
	EgressEntries []*NetworkAclEntry `json:"EgressEntries,omitnil" name:"EgressEntries"`

	// 网络ACL类型。三元组：'TRIPLE'   五元组：'QUINTUPLE'
	NetworkAclType *string `json:"NetworkAclType,omitnil" name:"NetworkAclType"`

	// 标签键值对
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`
}

type NetworkAclEntry struct {
	// 协议, 取值: TCP,UDP, ICMP, ALL。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 端口(all, 单个port,  range)。当Protocol为ALL或ICMP时，不能指定Port。
	Port *string `json:"Port,omitnil" name:"Port"`

	// 网段或IP(互斥)。增量创建ACL规则时，CidrBlock和Ipv6CidrBlock至少提供一个。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 网段或IPv6(互斥)。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`

	// ACCEPT 或 DROP。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 规则描述，最大长度100。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 优先级，从1开始。	
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// IPv4网络ACL条目唯一ID。当修改ACL条目时，NetworkAclIpv4EntryId和NetworkAclIpv6EntryID至少提供一个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAclIpv4EntryId *string `json:"NetworkAclIpv4EntryId,omitnil" name:"NetworkAclIpv4EntryId"`

	// IPv6网络ACL条目唯一ID。当修改ACL条目时，NetworkAclIpv4EntryId和NetworkAclIpv6EntryId至少提供一个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAclIpv6EntryId *string `json:"NetworkAclIpv6EntryId,omitnil" name:"NetworkAclIpv6EntryId"`
}

type NetworkAclEntrySet struct {
	// 入站规则。
	Ingress []*NetworkAclEntry `json:"Ingress,omitnil" name:"Ingress"`

	// 出站规则。
	Egress []*NetworkAclEntry `json:"Egress,omitnil" name:"Egress"`
}

type NetworkAclQuintupleEntries struct {
	// 网络ACL五元组入站规则。
	Ingress []*NetworkAclQuintupleEntry `json:"Ingress,omitnil" name:"Ingress"`

	// 网络ACL五元组出站规则
	Egress []*NetworkAclQuintupleEntry `json:"Egress,omitnil" name:"Egress"`
}

type NetworkAclQuintupleEntry struct {
	// 协议, 取值: TCP,UDP, ICMP, ALL。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 源端口(all, 单个port,  range)。当Protocol为ALL或ICMP时，不能指定Port。
	SourcePort *string `json:"SourcePort,omitnil" name:"SourcePort"`

	// 源CIDR。
	SourceCidr *string `json:"SourceCidr,omitnil" name:"SourceCidr"`

	// 目的端口(all, 单个port,  range)。当Protocol为ALL或ICMP时，不能指定Port。
	DestinationPort *string `json:"DestinationPort,omitnil" name:"DestinationPort"`

	// 目的CIDR。
	DestinationCidr *string `json:"DestinationCidr,omitnil" name:"DestinationCidr"`

	// 动作，ACCEPT 或 DROP。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 网络ACL条目唯一ID。
	NetworkAclQuintupleEntryId *string `json:"NetworkAclQuintupleEntryId,omitnil" name:"NetworkAclQuintupleEntryId"`

	// 优先级，从1开始。
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 创建时间，用于DescribeNetworkAclQuintupleEntries的出参。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 方向，INGRESS或EGRESS，用于DescribeNetworkAclQuintupleEntries的出参。
	NetworkAclDirection *string `json:"NetworkAclDirection,omitnil" name:"NetworkAclDirection"`
}

type NetworkInterface struct {
	// 弹性网卡实例ID，例如：eni-f1xjkw1b。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 弹性网卡名称。
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitnil" name:"NetworkInterfaceName"`

	// 弹性网卡描述。
	NetworkInterfaceDescription *string `json:"NetworkInterfaceDescription,omitnil" name:"NetworkInterfaceDescription"`

	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 绑定的安全组。
	GroupSet []*string `json:"GroupSet,omitnil" name:"GroupSet"`

	// 是否是主网卡。
	Primary *bool `json:"Primary,omitnil" name:"Primary"`

	// MAC地址。
	MacAddress *string `json:"MacAddress,omitnil" name:"MacAddress"`

	// 弹性网卡状态：
	// <li>`PENDING`：创建中</li>
	// <li>`AVAILABLE`：可用的</li>
	// <li>`ATTACHING`：绑定中</li>
	// <li>`DETACHING`：解绑中</li>
	// <li>`DELETING`：删除中</li>
	State *string `json:"State,omitnil" name:"State"`

	// 内网IP信息。
	PrivateIpAddressSet []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet,omitnil" name:"PrivateIpAddressSet"`

	// 绑定的云服务器对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Attachment *NetworkInterfaceAttachment `json:"Attachment,omitnil" name:"Attachment"`

	// 可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// `IPv6`地址列表。
	Ipv6AddressSet []*Ipv6Address `json:"Ipv6AddressSet,omitnil" name:"Ipv6AddressSet"`

	// 标签键值对。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 网卡类型。0 - 弹性网卡；1 - evm弹性网卡。
	EniType *uint64 `json:"EniType,omitnil" name:"EniType"`

	// 网卡绑定的子机类型：cvm，eks。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Business *string `json:"Business,omitnil" name:"Business"`

	// 网卡所关联的CDC实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// 弹性网卡类型：0:标准型/1:扩展型。默认值为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachType *uint64 `json:"AttachType,omitnil" name:"AttachType"`

	// 用于保留网卡主IP的资源ID用于保留网卡主IP的资源ID。用于删除网卡时作为入参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 服务质量级别：
	// <li>`DEFAULT`：默认</li>
	// <li>`PT`：云金</li>
	// <li>`AU`：云银</li>
	// <li>`AG`：云铜</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`
}

type NetworkInterfaceAttachment struct {
	// 云主机实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 网卡在云主机实例内的序号。
	DeviceIndex *uint64 `json:"DeviceIndex,omitnil" name:"DeviceIndex"`

	// 云主机所有者账户信息。
	InstanceAccountId *string `json:"InstanceAccountId,omitnil" name:"InstanceAccountId"`

	// 绑定时间。
	AttachTime *string `json:"AttachTime,omitnil" name:"AttachTime"`
}

// Predefined struct for user
type NotifyRoutesRequestParams struct {
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略唯一ID。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

type NotifyRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略唯一ID。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

func (r *NotifyRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *NotifyRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "NotifyRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type NotifyRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type NotifyRoutesResponse struct {
	*tchttp.BaseResponse
	Response *NotifyRoutesResponseParams `json:"Response"`
}

func (r *NotifyRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *NotifyRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PeerConnection struct {
	// 本端VPC唯一ID。
	SourceVpcId *string `json:"SourceVpcId,omitnil" name:"SourceVpcId"`

	// 对端VPC唯一ID。
	PeerVpcId *string `json:"PeerVpcId,omitnil" name:"PeerVpcId"`

	// 对等连接唯一ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`

	// 对等连接名称。
	PeeringConnectionName *string `json:"PeeringConnectionName,omitnil" name:"PeeringConnectionName"`

	// 对等连接状态，PENDING，投放中；ACTIVE，使用中；REJECTED，已拒绝‘DELETED，已删除；FAILED，失败；EXPIRED，已过期；ISOLATED，隔离中。
	State *string `json:"State,omitnil" name:"State"`

	// 是否是新控制器，true: 是NewAfc；false:不是。
	IsNgw *bool `json:"IsNgw,omitnil" name:"IsNgw"`

	// 对等连接带宽值。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 本端地域。
	SourceRegion *string `json:"SourceRegion,omitnil" name:"SourceRegion"`

	// 对端地域。
	DestinationRegion *string `json:"DestinationRegion,omitnil" name:"DestinationRegion"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 本端APPID。
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// 对端APPID。
	PeerAppId *int64 `json:"PeerAppId,omitnil" name:"PeerAppId"`

	// 计费类型，POSTPAID_BY_DAY_MAX：日峰值计费；POSTPAID_BY_MONTH_95：月95计费。
	ChargeType *string `json:"ChargeType,omitnil" name:"ChargeType"`

	// 本端UIN。
	SourceUin *int64 `json:"SourceUin,omitnil" name:"SourceUin"`

	// 对端UIN。
	DestinationUin *int64 `json:"DestinationUin,omitnil" name:"DestinationUin"`

	// 资源标签数据。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 服务分级：PT、AU、AG。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`

	// 互通类型，VPC_PEER：VPC间互通；VPC_BM_PEER：VPC与黑石网络互通。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type Price struct {
	// 实例价格。
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil" name:"InstancePrice"`

	// 带宽价格。
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitnil" name:"BandwidthPrice"`
}

type PrivateIpAddressSpecification struct {
	// 内网IP地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 是否是主IP。
	Primary *bool `json:"Primary,omitnil" name:"Primary"`

	// 公网IP地址。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil" name:"PublicIpAddress"`

	// EIP实例ID，例如：eip-11112222。
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 内网IP描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 公网IP是否被封堵。
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitnil" name:"IsWanIpBlocked"`

	// IP状态：
	// PENDING：生产中
	// MIGRATING：迁移中
	// DELETING：删除中
	// AVAILABLE：可用的
	State *string `json:"State,omitnil" name:"State"`

	// IP服务质量等级，可选值：PT、AU、AG、DEFAULT，分别代表云金、云银、云铜、默认四个等级。
	QosLevel *string `json:"QosLevel,omitnil" name:"QosLevel"`
}

type ProductQuota struct {
	// 产品配额ID
	QuotaId *string `json:"QuotaId,omitnil" name:"QuotaId"`

	// 产品配额名称
	QuotaName *string `json:"QuotaName,omitnil" name:"QuotaName"`

	// 产品当前配额
	QuotaCurrent *int64 `json:"QuotaCurrent,omitnil" name:"QuotaCurrent"`

	// 产品配额上限
	QuotaLimit *int64 `json:"QuotaLimit,omitnil" name:"QuotaLimit"`

	// 产品配额是否有地域属性
	QuotaRegion *bool `json:"QuotaRegion,omitnil" name:"QuotaRegion"`
}

type Quota struct {
	// 配额名称，取值范围：<br><li>`TOTAL_EIP_QUOTA`：用户当前地域下EIP的配额数；<br><li>`DAILY_EIP_APPLY`：用户当前地域下今日申购次数；<br><li>`DAILY_PUBLIC_IP_ASSIGN`：用户当前地域下，重新分配公网 IP次数。
	QuotaId *string `json:"QuotaId,omitnil" name:"QuotaId"`

	// 当前数量
	QuotaCurrent *int64 `json:"QuotaCurrent,omitnil" name:"QuotaCurrent"`

	// 配额数量
	QuotaLimit *int64 `json:"QuotaLimit,omitnil" name:"QuotaLimit"`
}

type ReferredSecurityGroup struct {
	// 安全组实例ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 引用安全组实例ID（SecurityGroupId）的所有安全组实例ID。
	ReferredSecurityGroupIds []*string `json:"ReferredSecurityGroupIds,omitnil" name:"ReferredSecurityGroupIds"`
}

// Predefined struct for user
type RefreshDirectConnectGatewayRouteToNatGatewayRequestParams struct {
	// vpc的ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关ID
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 是否是预刷新；True:是， False:否
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
}

type RefreshDirectConnectGatewayRouteToNatGatewayRequest struct {
	*tchttp.BaseRequest
	
	// vpc的ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关ID
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 是否是预刷新；True:是， False:否
	DryRun *bool `json:"DryRun,omitnil" name:"DryRun"`
}

func (r *RefreshDirectConnectGatewayRouteToNatGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDirectConnectGatewayRouteToNatGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "NatGatewayId")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshDirectConnectGatewayRouteToNatGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshDirectConnectGatewayRouteToNatGatewayResponseParams struct {
	// IDC子网信息
	DirectConnectSubnetSet []*DirectConnectSubnet `json:"DirectConnectSubnetSet,omitnil" name:"DirectConnectSubnetSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RefreshDirectConnectGatewayRouteToNatGatewayResponse struct {
	*tchttp.BaseResponse
	Response *RefreshDirectConnectGatewayRouteToNatGatewayResponseParams `json:"Response"`
}

func (r *RefreshDirectConnectGatewayRouteToNatGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshDirectConnectGatewayRouteToNatGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectAttachCcnInstancesRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 拒绝关联实例列表。
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

type RejectAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 拒绝关联实例列表。
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

func (r *RejectAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectAttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectAttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectAttachCcnInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RejectAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RejectAttachCcnInstancesResponseParams `json:"Response"`
}

func (r *RejectAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectVpcPeeringConnectionRequestParams struct {
	// 对等连接唯一ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`
}

type RejectVpcPeeringConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 对等连接唯一ID。
	PeeringConnectionId *string `json:"PeeringConnectionId,omitnil" name:"PeeringConnectionId"`
}

func (r *RejectVpcPeeringConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectVpcPeeringConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PeeringConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RejectVpcPeeringConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RejectVpcPeeringConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RejectVpcPeeringConnectionResponse struct {
	*tchttp.BaseResponse
	Response *RejectVpcPeeringConnectionResponseParams `json:"Response"`
}

func (r *RejectVpcPeeringConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RejectVpcPeeringConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseAddressesRequestParams struct {
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`
}

type ReleaseAddressesRequest struct {
	*tchttp.BaseRequest
	
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：`eip-11112222`。
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`
}

func (r *ReleaseAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseAddressesResponseParams struct {
	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReleaseAddressesResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseAddressesResponseParams `json:"Response"`
}

func (r *ReleaseAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIp6AddressesBandwidthRequestParams struct {
	// IPV6地址。Ip6Addresses和Ip6AddressIds必须且只能传一个
	Ip6Addresses []*string `json:"Ip6Addresses,omitnil" name:"Ip6Addresses"`

	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressIds必须且只能传一个。
	Ip6AddressIds []*string `json:"Ip6AddressIds,omitnil" name:"Ip6AddressIds"`
}

type ReleaseIp6AddressesBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// IPV6地址。Ip6Addresses和Ip6AddressIds必须且只能传一个
	Ip6Addresses []*string `json:"Ip6Addresses,omitnil" name:"Ip6Addresses"`

	// IPV6地址对应的唯一ID，形如eip-xxxxxxxx。Ip6Addresses和Ip6AddressIds必须且只能传一个。
	Ip6AddressIds []*string `json:"Ip6AddressIds,omitnil" name:"Ip6AddressIds"`
}

func (r *ReleaseIp6AddressesBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIp6AddressesBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6Addresses")
	delete(f, "Ip6AddressIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseIp6AddressesBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseIp6AddressesBandwidthResponseParams struct {
	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReleaseIp6AddressesBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseIp6AddressesBandwidthResponseParams `json:"Response"`
}

func (r *ReleaseIp6AddressesBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseIp6AddressesBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveBandwidthPackageResourcesRequestParams struct {
	// 带宽包唯一标识ID，形如'bwp-xxxx'
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 资源类型，包括‘Address’, ‘LoadBalance’
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源ID，可支持资源形如'eip-xxxx', 'lb-xxxx'
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

type RemoveBandwidthPackageResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 带宽包唯一标识ID，形如'bwp-xxxx'
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil" name:"BandwidthPackageId"`

	// 资源类型，包括‘Address’, ‘LoadBalance’
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源ID，可支持资源形如'eip-xxxx', 'lb-xxxx'
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`
}

func (r *RemoveBandwidthPackageResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveBandwidthPackageResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "ResourceType")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveBandwidthPackageResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveBandwidthPackageResourcesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RemoveBandwidthPackageResourcesResponse struct {
	*tchttp.BaseResponse
	Response *RemoveBandwidthPackageResourcesResponseParams `json:"Response"`
}

func (r *RemoveBandwidthPackageResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveBandwidthPackageResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveIp6RulesRequestParams struct {
	// IPV6转换规则所属的转换实例唯一ID，形如ip6-xxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// 待删除IPV6转换规则，形如rule6-xxxxxxxx
	Ip6RuleIds []*string `json:"Ip6RuleIds,omitnil" name:"Ip6RuleIds"`
}

type RemoveIp6RulesRequest struct {
	*tchttp.BaseRequest
	
	// IPV6转换规则所属的转换实例唯一ID，形如ip6-xxxxxxxx
	Ip6TranslatorId *string `json:"Ip6TranslatorId,omitnil" name:"Ip6TranslatorId"`

	// 待删除IPV6转换规则，形如rule6-xxxxxxxx
	Ip6RuleIds []*string `json:"Ip6RuleIds,omitnil" name:"Ip6RuleIds"`
}

func (r *RemoveIp6RulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveIp6RulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip6TranslatorId")
	delete(f, "Ip6RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveIp6RulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveIp6RulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RemoveIp6RulesResponse struct {
	*tchttp.BaseResponse
	Response *RemoveIp6RulesResponseParams `json:"Response"`
}

func (r *RemoveIp6RulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveIp6RulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewAddressesRequestParams struct {
	// EIP唯一标识ID列表，形如'eip-xxxx'
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`

	// 续费参数
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil" name:"AddressChargePrepaid"`
}

type RenewAddressesRequest struct {
	*tchttp.BaseRequest
	
	// EIP唯一标识ID列表，形如'eip-xxxx'
	AddressIds []*string `json:"AddressIds,omitnil" name:"AddressIds"`

	// 续费参数
	AddressChargePrepaid *AddressChargePrepaid `json:"AddressChargePrepaid,omitnil" name:"AddressChargePrepaid"`
}

func (r *RenewAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIds")
	delete(f, "AddressChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewAddressesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewAddressesResponse struct {
	*tchttp.BaseResponse
	Response *RenewAddressesResponseParams `json:"Response"`
}

func (r *RenewAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewVpnGatewayRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 预付费计费模式。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

type RenewVpnGatewayRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 预付费计费模式。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil" name:"InstanceChargePrepaid"`
}

func (r *RenewVpnGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewVpnGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InstanceChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewVpnGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewVpnGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RenewVpnGatewayResponse struct {
	*tchttp.BaseResponse
	Response *RenewVpnGatewayResponseParams `json:"Response"`
}

func (r *RenewVpnGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewVpnGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceDirectConnectGatewayCcnRoutesRequestParams struct {
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 需要连通的IDC网段列表
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil" name:"Routes"`
}

type ReplaceDirectConnectGatewayCcnRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 专线网关ID，形如：dcg-prpqlmg1
	DirectConnectGatewayId *string `json:"DirectConnectGatewayId,omitnil" name:"DirectConnectGatewayId"`

	// 需要连通的IDC网段列表
	Routes []*DirectConnectGatewayCcnRoute `json:"Routes,omitnil" name:"Routes"`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceDirectConnectGatewayCcnRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DirectConnectGatewayId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceDirectConnectGatewayCcnRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceDirectConnectGatewayCcnRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReplaceDirectConnectGatewayCcnRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceDirectConnectGatewayCcnRoutesResponseParams `json:"Response"`
}

func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceDirectConnectGatewayCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationRequestParams struct {
	// 子网实例ID，例如：subnet-3x5lf5q0。可通过DescribeSubnets接口查询。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`
}

type ReplaceRouteTableAssociationRequest struct {
	*tchttp.BaseRequest
	
	// 子网实例ID，例如：subnet-3x5lf5q0。可通过DescribeSubnets接口查询。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`
}

func (r *ReplaceRouteTableAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRouteTableAssociationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "RouteTableId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceRouteTableAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRouteTableAssociationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReplaceRouteTableAssociationResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceRouteTableAssociationResponseParams `json:"Response"`
}

func (r *ReplaceRouteTableAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRouteTableAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRoutesRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略对象。需要指定路由策略ID（RouteId）。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

type ReplaceRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略对象。需要指定路由策略ID（RouteId）。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

func (r *ReplaceRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceRoutesResponseParams struct {
	// 原路由策略信息。
	OldRouteSet []*Route `json:"OldRouteSet,omitnil" name:"OldRouteSet"`

	// 修改后的路由策略信息。
	NewRouteSet []*Route `json:"NewRouteSet,omitnil" name:"NewRouteSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReplaceRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceRoutesResponseParams `json:"Response"`
}

func (r *ReplaceRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPoliciesRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合对象。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`

	// 旧的安全组规则集合对象，可选，日志记录用。
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil" name:"OriginalSecurityGroupPolicySet"`
}

type ReplaceSecurityGroupPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合对象。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`

	// 旧的安全组规则集合对象，可选，日志记录用。
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil" name:"OriginalSecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	delete(f, "OriginalSecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceSecurityGroupPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPoliciesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReplaceSecurityGroupPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceSecurityGroupPoliciesResponseParams `json:"Response"`
}

func (r *ReplaceSecurityGroupPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyRequestParams struct {
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合对象。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`

	// 旧的安全组规则集合对象，可选，日志记录用。
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil" name:"OriginalSecurityGroupPolicySet"`
}

type ReplaceSecurityGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 安全组实例ID，例如sg-33ocnj9n，可通过<a href="https://cloud.tencent.com/document/product/215/15808">DescribeSecurityGroups</a>获取。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组规则集合对象。
	SecurityGroupPolicySet *SecurityGroupPolicySet `json:"SecurityGroupPolicySet,omitnil" name:"SecurityGroupPolicySet"`

	// 旧的安全组规则集合对象，可选，日志记录用。
	OriginalSecurityGroupPolicySet *SecurityGroupPolicySet `json:"OriginalSecurityGroupPolicySet,omitnil" name:"OriginalSecurityGroupPolicySet"`
}

func (r *ReplaceSecurityGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroupId")
	delete(f, "SecurityGroupPolicySet")
	delete(f, "OriginalSecurityGroupPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceSecurityGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceSecurityGroupPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReplaceSecurityGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceSecurityGroupPolicyResponseParams `json:"Response"`
}

func (r *ReplaceSecurityGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceSecurityGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAttachCcnInstancesRequestParams struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN所属UIN（根账号）。
	CcnUin *string `json:"CcnUin,omitnil" name:"CcnUin"`

	// 重新申请关联网络实例列表。
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

type ResetAttachCcnInstancesRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// CCN所属UIN（根账号）。
	CcnUin *string `json:"CcnUin,omitnil" name:"CcnUin"`

	// 重新申请关联网络实例列表。
	Instances []*CcnInstance `json:"Instances,omitnil" name:"Instances"`
}

func (r *ResetAttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "CcnUin")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetAttachCcnInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetAttachCcnInstancesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetAttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ResetAttachCcnInstancesResponseParams `json:"Response"`
}

func (r *ResetAttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetAttachCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetNatGatewayConnectionRequestParams struct {
	// NAT网关ID。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关并发连接上限，形如：1000000、3000000、10000000。
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil" name:"MaxConcurrentConnection"`
}

type ResetNatGatewayConnectionRequest struct {
	*tchttp.BaseRequest
	
	// NAT网关ID。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// NAT网关并发连接上限，形如：1000000、3000000、10000000。
	MaxConcurrentConnection *uint64 `json:"MaxConcurrentConnection,omitnil" name:"MaxConcurrentConnection"`
}

func (r *ResetNatGatewayConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetNatGatewayConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatGatewayId")
	delete(f, "MaxConcurrentConnection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetNatGatewayConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetNatGatewayConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetNatGatewayConnectionResponse struct {
	*tchttp.BaseResponse
	Response *ResetNatGatewayConnectionResponseParams `json:"Response"`
}

func (r *ResetNatGatewayConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetNatGatewayConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRoutesRequestParams struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`

	// 路由策略。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

type ResetRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`

	// 路由策略。
	Routes []*Route `json:"Routes,omitnil" name:"Routes"`
}

func (r *ResetRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteTableName")
	delete(f, "Routes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetRoutesResponse struct {
	*tchttp.BaseResponse
	Response *ResetRoutesResponseParams `json:"Response"`
}

func (r *ResetRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnConnectionRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`
}

type ResetVpnConnectionRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPN通道实例ID。形如：vpnx-f49l6u0z。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`
}

func (r *ResetVpnConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "VpnConnectionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetVpnConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnConnectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetVpnConnectionResponse struct {
	*tchttp.BaseResponse
	Response *ResetVpnConnectionResponseParams `json:"Response"`
}

func (r *ResetVpnConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnGatewayInternetMaxBandwidthRequestParams struct {
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 新规格公网带宽设置。可选带宽规格：5, 10, 20, 50, 100, 200, 500, 1000；单位：Mbps。VPN网关带宽目前仅支持部分带宽范围内升降配，如【5,100】Mbps和【200,1000】Mbps，在各自带宽范围内可提升配额，跨范围提升配额和降配暂不支持。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`
}

type ResetVpnGatewayInternetMaxBandwidthRequest struct {
	*tchttp.BaseRequest
	
	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 新规格公网带宽设置。可选带宽规格：5, 10, 20, 50, 100, 200, 500, 1000；单位：Mbps。VPN网关带宽目前仅支持部分带宽范围内升降配，如【5,100】Mbps和【200,1000】Mbps，在各自带宽范围内可提升配额，跨范围提升配额和降配暂不支持。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnGatewayInternetMaxBandwidthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayId")
	delete(f, "InternetMaxBandwidthOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetVpnGatewayInternetMaxBandwidthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetVpnGatewayInternetMaxBandwidthResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetVpnGatewayInternetMaxBandwidthResponse struct {
	*tchttp.BaseResponse
	Response *ResetVpnGatewayInternetMaxBandwidthResponseParams `json:"Response"`
}

func (r *ResetVpnGatewayInternetMaxBandwidthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetVpnGatewayInternetMaxBandwidthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 带宽包资源类型，包括'Address'和'LoadBalance'
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 带宽包资源Id，形如'eip-xxxx', 'lb-xxxx'
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 带宽包资源Ip
	AddressIp *string `json:"AddressIp,omitnil" name:"AddressIp"`
}

type ResourceDashboard struct {
	// Vpc实例ID，例如：vpc-bq4bzxpj。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID，例如：subnet-bthucmmy。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 基础网络互通。
	Classiclink *uint64 `json:"Classiclink,omitnil" name:"Classiclink"`

	// 专线网关。
	Dcg *uint64 `json:"Dcg,omitnil" name:"Dcg"`

	// 对等连接。
	Pcx *uint64 `json:"Pcx,omitnil" name:"Pcx"`

	// 统计当前除云服务器 IP、弹性网卡IP和网络探测IP以外的所有已使用的IP总数。云服务器 IP、弹性网卡IP和网络探测IP单独计数。
	Ip *uint64 `json:"Ip,omitnil" name:"Ip"`

	// NAT网关。
	Nat *uint64 `json:"Nat,omitnil" name:"Nat"`

	// VPN网关。
	Vpngw *uint64 `json:"Vpngw,omitnil" name:"Vpngw"`

	// 流日志。
	FlowLog *uint64 `json:"FlowLog,omitnil" name:"FlowLog"`

	// 网络探测。
	NetworkDetect *uint64 `json:"NetworkDetect,omitnil" name:"NetworkDetect"`

	// 网络ACL。
	NetworkACL *uint64 `json:"NetworkACL,omitnil" name:"NetworkACL"`

	// 云主机。
	CVM *uint64 `json:"CVM,omitnil" name:"CVM"`

	// 负载均衡。
	LB *uint64 `json:"LB,omitnil" name:"LB"`

	// 关系型数据库。
	CDB *uint64 `json:"CDB,omitnil" name:"CDB"`

	// 云数据库 TencentDB for Memcached。
	Cmem *uint64 `json:"Cmem,omitnil" name:"Cmem"`

	// 时序数据库。
	CTSDB *uint64 `json:"CTSDB,omitnil" name:"CTSDB"`

	// 数据库 TencentDB for MariaDB（TDSQL）。
	MariaDB *uint64 `json:"MariaDB,omitnil" name:"MariaDB"`

	// 数据库 TencentDB for SQL Server。
	SQLServer *uint64 `json:"SQLServer,omitnil" name:"SQLServer"`

	// 云数据库 TencentDB for PostgreSQL。
	Postgres *uint64 `json:"Postgres,omitnil" name:"Postgres"`

	// 网络附加存储。
	NAS *uint64 `json:"NAS,omitnil" name:"NAS"`

	// Snova云数据仓库。
	Greenplumn *uint64 `json:"Greenplumn,omitnil" name:"Greenplumn"`

	// 消息队列 CKAFKA。
	Ckafka *uint64 `json:"Ckafka,omitnil" name:"Ckafka"`

	// Grocery。
	Grocery *uint64 `json:"Grocery,omitnil" name:"Grocery"`

	// 数据加密服务。
	HSM *uint64 `json:"HSM,omitnil" name:"HSM"`

	// 游戏存储 Tcaplus。
	Tcaplus *uint64 `json:"Tcaplus,omitnil" name:"Tcaplus"`

	// Cnas。
	Cnas *uint64 `json:"Cnas,omitnil" name:"Cnas"`

	// HTAP 数据库 TiDB。
	TiDB *uint64 `json:"TiDB,omitnil" name:"TiDB"`

	// EMR 集群。
	Emr *uint64 `json:"Emr,omitnil" name:"Emr"`

	// SEAL。
	SEAL *uint64 `json:"SEAL,omitnil" name:"SEAL"`

	// 文件存储 CFS。
	CFS *uint64 `json:"CFS,omitnil" name:"CFS"`

	// Oracle。
	Oracle *uint64 `json:"Oracle,omitnil" name:"Oracle"`

	// ElasticSearch服务。
	ElasticSearch *uint64 `json:"ElasticSearch,omitnil" name:"ElasticSearch"`

	// 区块链服务。
	TBaaS *uint64 `json:"TBaaS,omitnil" name:"TBaaS"`

	// Itop。
	Itop *uint64 `json:"Itop,omitnil" name:"Itop"`

	// 云数据库审计。
	DBAudit *uint64 `json:"DBAudit,omitnil" name:"DBAudit"`

	// 企业级云数据库 CynosDB for Postgres。
	CynosDBPostgres *uint64 `json:"CynosDBPostgres,omitnil" name:"CynosDBPostgres"`

	// 数据库 TencentDB for Redis。
	Redis *uint64 `json:"Redis,omitnil" name:"Redis"`

	// 数据库 TencentDB for MongoDB。
	MongoDB *uint64 `json:"MongoDB,omitnil" name:"MongoDB"`

	// 分布式数据库 TencentDB for TDSQL。
	DCDB *uint64 `json:"DCDB,omitnil" name:"DCDB"`

	// 企业级云数据库 CynosDB for MySQL。
	CynosDBMySQL *uint64 `json:"CynosDBMySQL,omitnil" name:"CynosDBMySQL"`

	// 子网。
	Subnet *uint64 `json:"Subnet,omitnil" name:"Subnet"`

	// 路由表。
	RouteTable *uint64 `json:"RouteTable,omitnil" name:"RouteTable"`
}

type ResourceStatistics struct {
	// Vpc实例ID，例如：vpc-f1xjkw1b。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例ID，例如：subnet-bthucmmy。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 当前已使用的IP总数。
	Ip *uint64 `json:"Ip,omitnil" name:"Ip"`

	// 资源统计信息。
	ResourceStatisticsItemSet []*ResourceStatisticsItem `json:"ResourceStatisticsItemSet,omitnil" name:"ResourceStatisticsItemSet"`
}

type ResourceStatisticsItem struct {
	// 资源类型。比如，CVM，ENI等。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源名称。
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 资源个数。
	ResourceCount *uint64 `json:"ResourceCount,omitnil" name:"ResourceCount"`
}

// Predefined struct for user
type ResumeSnapshotInstanceRequestParams struct {
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 快照文件Id。
	SnapshotFileId *string `json:"SnapshotFileId,omitnil" name:"SnapshotFileId"`

	// 实例Id。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type ResumeSnapshotInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 快照文件Id。
	SnapshotFileId *string `json:"SnapshotFileId,omitnil" name:"SnapshotFileId"`

	// 实例Id。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *ResumeSnapshotInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSnapshotInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SnapshotPolicyId")
	delete(f, "SnapshotFileId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeSnapshotInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeSnapshotInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResumeSnapshotInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ResumeSnapshotInstanceResponseParams `json:"Response"`
}

func (r *ResumeSnapshotInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSnapshotInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReturnNormalAddressesRequestParams struct {
	// EIP 的 IP 地址,示例：101.35.139.183
	AddressIps []*string `json:"AddressIps,omitnil" name:"AddressIps"`
}

type ReturnNormalAddressesRequest struct {
	*tchttp.BaseRequest
	
	// EIP 的 IP 地址,示例：101.35.139.183
	AddressIps []*string `json:"AddressIps,omitnil" name:"AddressIps"`
}

func (r *ReturnNormalAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReturnNormalAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReturnNormalAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReturnNormalAddressesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReturnNormalAddressesResponse struct {
	*tchttp.BaseResponse
	Response *ReturnNormalAddressesResponseParams `json:"Response"`
}

func (r *ReturnNormalAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReturnNormalAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Route struct {
	// 目的网段，取值不能在私有网络网段内，例如：112.20.51.0/24。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil" name:"DestinationCidrBlock"`

	// 下一跳类型，目前我们支持的类型有：
	// CVM：公网网关类型的云服务器；
	// VPN：VPN网关；
	// DIRECTCONNECT：专线网关；
	// PEERCONNECTION：对等连接；
	// HAVIP：高可用虚拟IP；
	// NAT：NAT网关; 
	// NORMAL_CVM：普通云服务器；
	// EIP：云服务器的公网IP；
	// LOCAL_GATEWAY：本地网关。
	GatewayType *string `json:"GatewayType,omitnil" name:"GatewayType"`

	// 下一跳地址，这里只需要指定不同下一跳类型的网关ID，系统会自动匹配到下一跳地址。
	// 特殊说明：GatewayType为NORMAL_CVM时，GatewayId填写实例的内网IP。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由策略ID。IPv4路由策略ID是有意义的值，IPv6路由策略是无意义的值0。后续建议完全使用字符串唯一ID `RouteItemId`操作路由策略。
	// 该字段在删除时必填，其他字段无需填写。
	RouteId *uint64 `json:"RouteId,omitnil" name:"RouteId"`

	// 路由策略描述。
	RouteDescription *string `json:"RouteDescription,omitnil" name:"RouteDescription"`

	// 是否启用
	Enabled *bool `json:"Enabled,omitnil" name:"Enabled"`

	// 路由类型，目前我们支持的类型有：
	// USER：用户路由；
	// NETD：网络探测路由，创建网络探测实例时，系统默认下发，不可编辑与删除；
	// CCN：云联网路由，系统默认下发，不可编辑与删除。
	// 用户只能添加和操作 USER 类型的路由。
	RouteType *string `json:"RouteType,omitnil" name:"RouteType"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 目的IPv6网段，取值不能在私有网络网段内，例如：2402:4e00:1000:810b::/64。
	DestinationIpv6CidrBlock *string `json:"DestinationIpv6CidrBlock,omitnil" name:"DestinationIpv6CidrBlock"`

	// 路由唯一策略ID。
	RouteItemId *string `json:"RouteItemId,omitnil" name:"RouteItemId"`

	// 路由策略是否发布到云联网。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishedToVbc *bool `json:"PublishedToVbc,omitnil" name:"PublishedToVbc"`

	// 路由策略创建时间
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type RouteConflict struct {
	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 要检查的与之冲突的目的端
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil" name:"DestinationCidrBlock"`

	// 冲突的路由策略列表
	ConflictSet []*Route `json:"ConflictSet,omitnil" name:"ConflictSet"`
}

type RouteTable struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 路由表实例ID，例如：rtb-azd4dt1c。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由表名称。
	RouteTableName *string `json:"RouteTableName,omitnil" name:"RouteTableName"`

	// 路由表关联关系。
	AssociationSet []*RouteTableAssociation `json:"AssociationSet,omitnil" name:"AssociationSet"`

	// IPv4路由策略集合。
	RouteSet []*Route `json:"RouteSet,omitnil" name:"RouteSet"`

	// 是否默认路由表。
	Main *bool `json:"Main,omitnil" name:"Main"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 标签键值对。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// local路由是否发布云联网。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalCidrForCcn []*CidrForCcn `json:"LocalCidrForCcn,omitnil" name:"LocalCidrForCcn"`
}

type RouteTableAssociation struct {
	// 子网实例ID。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 路由表实例ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`
}

type SecurityGroup struct {
	// 安全组实例ID，例如：sg-ohuuioma。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 安全组名称，可任意命名，但不得超过60个字符。
	SecurityGroupName *string `json:"SecurityGroupName,omitnil" name:"SecurityGroupName"`

	// 安全组备注，最多100个字符。
	SecurityGroupDesc *string `json:"SecurityGroupDesc,omitnil" name:"SecurityGroupDesc"`

	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 是否是默认安全组，默认安全组不支持删除。
	IsDefault *bool `json:"IsDefault,omitnil" name:"IsDefault"`

	// 安全组创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 标签键值对。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 安全组更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type SecurityGroupAssociationStatistics struct {
	// 安全组实例ID。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 云服务器实例数。
	CVM *uint64 `json:"CVM,omitnil" name:"CVM"`

	// MySQL数据库实例数。
	CDB *uint64 `json:"CDB,omitnil" name:"CDB"`

	// 弹性网卡实例数。
	ENI *uint64 `json:"ENI,omitnil" name:"ENI"`

	// 被安全组引用数。
	SG *uint64 `json:"SG,omitnil" name:"SG"`

	// 负载均衡实例数。
	CLB *uint64 `json:"CLB,omitnil" name:"CLB"`

	// 全量实例的绑定统计。
	InstanceStatistics []*InstanceStatistic `json:"InstanceStatistics,omitnil" name:"InstanceStatistics"`

	// 所有资源的总计数（不包含被安全组引用数）。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type SecurityGroupLimitSet struct {
	// 每个项目每个地域可创建安全组数
	SecurityGroupLimit *uint64 `json:"SecurityGroupLimit,omitnil" name:"SecurityGroupLimit"`

	// 安全组下的最大规则数
	SecurityGroupPolicyLimit *uint64 `json:"SecurityGroupPolicyLimit,omitnil" name:"SecurityGroupPolicyLimit"`

	// 安全组下嵌套安全组规则数
	ReferedSecurityGroupLimit *uint64 `json:"ReferedSecurityGroupLimit,omitnil" name:"ReferedSecurityGroupLimit"`

	// 单安全组关联实例数
	SecurityGroupInstanceLimit *uint64 `json:"SecurityGroupInstanceLimit,omitnil" name:"SecurityGroupInstanceLimit"`

	// 实例关联安全组数
	InstanceSecurityGroupLimit *uint64 `json:"InstanceSecurityGroupLimit,omitnil" name:"InstanceSecurityGroupLimit"`

	// 安全组展开后的规则数限制
	SecurityGroupExtendedPolicyLimit *uint64 `json:"SecurityGroupExtendedPolicyLimit,omitnil" name:"SecurityGroupExtendedPolicyLimit"`

	// 被引用的安全组关联CVM、ENI的实例配额
	SecurityGroupReferedCvmAndEniLimit *uint64 `json:"SecurityGroupReferedCvmAndEniLimit,omitnil" name:"SecurityGroupReferedCvmAndEniLimit"`

	// 被引用的安全组关联数据库、LB等服务实例配额
	SecurityGroupReferedSvcLimit *uint64 `json:"SecurityGroupReferedSvcLimit,omitnil" name:"SecurityGroupReferedSvcLimit"`
}

type SecurityGroupPolicy struct {
	// 安全组规则索引号，值会随着安全组规则的变更动态变化。使用PolicyIndex时，请先调用`DescribeSecurityGroupPolicies`获取到规则的PolicyIndex，并且结合返回值中的Version一起使用处理规则。
	PolicyIndex *int64 `json:"PolicyIndex,omitnil" name:"PolicyIndex"`

	// 协议, 取值: TCP,UDP,ICMP,ICMPv6,ALL。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 端口(all, 离散port,  range)。
	// 说明：如果Protocol设置为ALL，则Port也需要设置为all。
	Port *string `json:"Port,omitnil" name:"Port"`

	// 协议端口ID或者协议端口组ID。ServiceTemplate和Protocol+Port互斥。
	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitnil" name:"ServiceTemplate"`

	// 网段或IP(互斥)，特殊说明：0.0.0.0/n 都会映射为0.0.0.0/0。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 网段或IPv6(互斥)。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`

	// 安全组实例ID，例如：sg-ohuuioma。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// IP地址ID或者IP地址组ID。
	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitnil" name:"AddressTemplate"`

	// ACCEPT 或 DROP。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 安全组规则描述。
	PolicyDescription *string `json:"PolicyDescription,omitnil" name:"PolicyDescription"`

	// 安全组最近修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type SecurityGroupPolicySet struct {
	// 安全组规则当前版本。用户每次更新安全规则版本会自动加1，防止更新的路由规则已过期，不填不考虑冲突。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil" name:"Version"`

	// 出站规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Egress []*SecurityGroupPolicy `json:"Egress,omitnil" name:"Egress"`

	// 入站规则。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ingress []*SecurityGroupPolicy `json:"Ingress,omitnil" name:"Ingress"`
}

type SecurityPolicyDatabase struct {
	// 本端网段
	LocalCidrBlock *string `json:"LocalCidrBlock,omitnil" name:"LocalCidrBlock"`

	// 对端网段
	RemoteCidrBlock []*string `json:"RemoteCidrBlock,omitnil" name:"RemoteCidrBlock"`
}

type ServiceTemplate struct {
	// 协议端口实例ID，例如：ppm-f5n1f8da。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// 模板名称。
	ServiceTemplateName *string `json:"ServiceTemplateName,omitnil" name:"ServiceTemplateName"`

	// 协议端口信息。
	ServiceSet []*string `json:"ServiceSet,omitnil" name:"ServiceSet"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 带备注的协议端口信息。
	ServiceExtraSet []*ServicesInfo `json:"ServiceExtraSet,omitnil" name:"ServiceExtraSet"`
}

type ServiceTemplateGroup struct {
	// 协议端口模板集合实例ID，例如：ppmg-2klmrefu。
	ServiceTemplateGroupId *string `json:"ServiceTemplateGroupId,omitnil" name:"ServiceTemplateGroupId"`

	// 协议端口模板集合名称。
	ServiceTemplateGroupName *string `json:"ServiceTemplateGroupName,omitnil" name:"ServiceTemplateGroupName"`

	// 协议端口模板实例ID。
	ServiceTemplateIdSet []*string `json:"ServiceTemplateIdSet,omitnil" name:"ServiceTemplateIdSet"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 协议端口模板实例信息。
	ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet,omitnil" name:"ServiceTemplateSet"`
}

type ServiceTemplateSpecification struct {
	// 协议端口ID，例如：ppm-f5n1f8da。
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 协议端口组ID，例如：ppmg-f5n1f8da。
	ServiceGroupId *string `json:"ServiceGroupId,omitnil" name:"ServiceGroupId"`
}

type ServicesInfo struct {
	// 协议端口。
	Service *string `json:"Service,omitnil" name:"Service"`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type SetCcnRegionBandwidthLimitsRequestParams struct {
	// CCN实例ID，形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网（CCN）各地域出带宽上限。
	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitnil" name:"CcnRegionBandwidthLimits"`

	// 是否恢复云联网地域出口/地域间带宽限速为默认值（1Gbps）。false表示不恢复；true表示恢复。恢复默认值后，限速实例将不在控制台展示。该参数默认为 false，不恢复。
	SetDefaultLimitFlag *bool `json:"SetDefaultLimitFlag,omitnil" name:"SetDefaultLimitFlag"`
}

type SetCcnRegionBandwidthLimitsRequest struct {
	*tchttp.BaseRequest
	
	// CCN实例ID，形如：ccn-f49l6u0z。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云联网（CCN）各地域出带宽上限。
	CcnRegionBandwidthLimits []*CcnRegionBandwidthLimit `json:"CcnRegionBandwidthLimits,omitnil" name:"CcnRegionBandwidthLimits"`

	// 是否恢复云联网地域出口/地域间带宽限速为默认值（1Gbps）。false表示不恢复；true表示恢复。恢复默认值后，限速实例将不在控制台展示。该参数默认为 false，不恢复。
	SetDefaultLimitFlag *bool `json:"SetDefaultLimitFlag,omitnil" name:"SetDefaultLimitFlag"`
}

func (r *SetCcnRegionBandwidthLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCcnRegionBandwidthLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "CcnRegionBandwidthLimits")
	delete(f, "SetDefaultLimitFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCcnRegionBandwidthLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCcnRegionBandwidthLimitsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetCcnRegionBandwidthLimitsResponse struct {
	*tchttp.BaseResponse
	Response *SetCcnRegionBandwidthLimitsResponseParams `json:"Response"`
}

func (r *SetCcnRegionBandwidthLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCcnRegionBandwidthLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVpnGatewaysRenewFlagRequestParams struct {
	// VPNGW字符型ID列表。可通过[DescribeVpnGateways](https://cloud.tencent.com/document/api/215/17514)接口返回值VpnGatewaySet中的VpnGatewayId获取。
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil" name:"VpnGatewayIds"`

	// 自动续费标记 [0, 1, 2]
	// 0表示默认状态(初始状态)， 1表示自动续费，2表示明确不自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// VPNGW类型['IPSEC', 'SSL']， 默认为IPSEC。
	Type *string `json:"Type,omitnil" name:"Type"`
}

type SetVpnGatewaysRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// VPNGW字符型ID列表。可通过[DescribeVpnGateways](https://cloud.tencent.com/document/api/215/17514)接口返回值VpnGatewaySet中的VpnGatewayId获取。
	VpnGatewayIds []*string `json:"VpnGatewayIds,omitnil" name:"VpnGatewayIds"`

	// 自动续费标记 [0, 1, 2]
	// 0表示默认状态(初始状态)， 1表示自动续费，2表示明确不自动续费。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// VPNGW类型['IPSEC', 'SSL']， 默认为IPSEC。
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *SetVpnGatewaysRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVpnGatewaysRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpnGatewayIds")
	delete(f, "AutoRenewFlag")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetVpnGatewaysRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetVpnGatewaysRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SetVpnGatewaysRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *SetVpnGatewaysRenewFlagResponseParams `json:"Response"`
}

func (r *SetVpnGatewaysRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetVpnGatewaysRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SnapshotFileInfo struct {
	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 实例Id。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 快照文件Id。
	SnapshotFileId *string `json:"SnapshotFileId,omitnil" name:"SnapshotFileId"`

	// 备份时间。
	BackupTime *string `json:"BackupTime,omitnil" name:"BackupTime"`

	// 操作者Uin。
	Operator *string `json:"Operator,omitnil" name:"Operator"`
}

type SnapshotInstance struct {
	// 实例Id。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例类型，目前支持安全组：securitygroup。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例所在地域。
	InstanceRegion *string `json:"InstanceRegion,omitnil" name:"InstanceRegion"`

	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`
}

type SnapshotPolicy struct {
	// 快照策略名称。
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitnil" name:"SnapshotPolicyName"`

	// 备份策略类型，operate-操作备份，time-定时备份。
	BackupType *string `json:"BackupType,omitnil" name:"BackupType"`

	// 保留时间，支持1～365天。
	KeepTime *uint64 `json:"KeepTime,omitnil" name:"KeepTime"`

	// 是否创建新的cos桶，默认为False。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateNewCos *bool `json:"CreateNewCos,omitnil" name:"CreateNewCos"`

	// cos桶所在地域。
	CosRegion *string `json:"CosRegion,omitnil" name:"CosRegion"`

	// cos桶。
	CosBucket *string `json:"CosBucket,omitnil" name:"CosBucket"`

	// 快照策略Id。
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitnil" name:"SnapshotPolicyId"`

	// 时间备份策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupPolicies []*BackupPolicy `json:"BackupPolicies,omitnil" name:"BackupPolicies"`

	// 启用状态，True-启用，False-停用，默认为True。
	Enable *bool `json:"Enable,omitnil" name:"Enable"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type SourceIpTranslationNatRule struct {
	// 资源ID，如果ResourceType为USERDEFINED，可以为空
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源类型，目前包含SUBNET、NETWORKINTERFACE、USERDEFINED
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 源IP/网段
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 弹性IP地址池
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// Snat规则ID
	NatGatewaySnatId *string `json:"NatGatewaySnatId,omitnil" name:"NatGatewaySnatId"`

	// NAT网关的ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 私有网络VPC的ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// NAT网关SNAT规则创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type SslClientConfig struct {
	// 客户端配置
	SslVpnClientConfiguration *string `json:"SslVpnClientConfiguration,omitnil" name:"SslVpnClientConfiguration"`

	// 更证书
	SslVpnRootCert *string `json:"SslVpnRootCert,omitnil" name:"SslVpnRootCert"`

	// 客户端密钥
	SslVpnKey *string `json:"SslVpnKey,omitnil" name:"SslVpnKey"`

	// 客户端证书
	SslVpnCert *string `json:"SslVpnCert,omitnil" name:"SslVpnCert"`

	// SSL-VPN-CLIENT 实例ID。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`
}

type SslVpnClient struct {
	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// SSL-VPN-SERVER 实例ID。
	SslVpnServerId *string `json:"SslVpnServerId,omitnil" name:"SslVpnServerId"`

	// 证书状态。
	// 0:创建中
	// 1:正常
	// 2:已停用
	// 3.已过期
	// 4.创建出错
	CertStatus *uint64 `json:"CertStatus,omitnil" name:"CertStatus"`

	// SSL-VPN-CLIENT 实例ID。
	SslVpnClientId *string `json:"SslVpnClientId,omitnil" name:"SslVpnClientId"`

	// 证书开始时间。
	CertBeginTime *string `json:"CertBeginTime,omitnil" name:"CertBeginTime"`

	// 证书到期时间。
	CertEndTime *string `json:"CertEndTime,omitnil" name:"CertEndTime"`

	// CLIENT NAME。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 创建CLIENT 状态。
	// 0 创建中
	// 1 创建出错
	// 2 更新中
	// 3 更新出错
	// 4 销毁中
	// 5 销毁出错
	// 6 已连通
	// 7 未知
	State *string `json:"State,omitnil" name:"State"`
}

type SslVpnSever struct {
	// VPC实例ID.
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// SSL-VPN-SERVER 实例ID。
	SslVpnServerId *string `json:"SslVpnServerId,omitnil" name:"SslVpnServerId"`

	// VPN 实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// SSL-VPN-SERVER name。
	SslVpnServerName *string `json:"SslVpnServerName,omitnil" name:"SslVpnServerName"`

	// 本端地址段。
	LocalAddress []*string `json:"LocalAddress,omitnil" name:"LocalAddress"`

	// 客户端地址段。
	RemoteAddress *string `json:"RemoteAddress,omitnil" name:"RemoteAddress"`

	// 客户端最大连接数。
	MaxConnection *uint64 `json:"MaxConnection,omitnil" name:"MaxConnection"`

	// SSL-VPN 网关公网IP。
	WanIp *string `json:"WanIp,omitnil" name:"WanIp"`

	// SSL VPN服务端监听协议
	SslVpnProtocol *string `json:"SslVpnProtocol,omitnil" name:"SslVpnProtocol"`

	// SSL VPN服务端监听协议端口
	SslVpnPort *uint64 `json:"SslVpnPort,omitnil" name:"SslVpnPort"`

	// 加密算法。
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitnil" name:"EncryptAlgorithm"`

	// 认证算法。
	IntegrityAlgorithm *string `json:"IntegrityAlgorithm,omitnil" name:"IntegrityAlgorithm"`

	// 是否支持压缩。
	Compress *uint64 `json:"Compress,omitnil" name:"Compress"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// SSL-VPN-SERVER 创建状态。
	// 0 创建中
	// 1 创建出错
	// 2 更新中
	// 3 更新出错
	// 4 销毁中
	// 5 销毁出错
	// 6 已连通
	// 7 未知
	State *uint64 `json:"State,omitnil" name:"State"`

	// 是否开启SSO认证。1：开启  0： 不开启
	SsoEnabled *uint64 `json:"SsoEnabled,omitnil" name:"SsoEnabled"`

	// EIAM应用ID
	EiamApplicationId *string `json:"EiamApplicationId,omitnil" name:"EiamApplicationId"`

	// 是否开启策略控制。0：不开启 1： 开启
	AccessPolicyEnabled *uint64 `json:"AccessPolicyEnabled,omitnil" name:"AccessPolicyEnabled"`

	// 策略信息
	AccessPolicy []*AccessPolicy `json:"AccessPolicy,omitnil" name:"AccessPolicy"`
}

type Subnet struct {
	// `VPC`实例`ID`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网实例`ID`，例如：subnet-bthucmmy。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网名称。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 子网的 `IPv4` `CIDR`。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 是否默认子网。
	IsDefault *bool `json:"IsDefault,omitnil" name:"IsDefault"`

	// 是否开启广播。
	EnableBroadcast *bool `json:"EnableBroadcast,omitnil" name:"EnableBroadcast"`

	// 可用区。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 路由表实例ID，例如：rtb-l2h8d7c2。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 可用`IPv4`数。
	AvailableIpAddressCount *uint64 `json:"AvailableIpAddressCount,omitnil" name:"AvailableIpAddressCount"`

	// 子网的 `IPv6` `CIDR`。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`

	// 关联`ACL`ID
	NetworkAclId *string `json:"NetworkAclId,omitnil" name:"NetworkAclId"`

	// 是否为 `SNAT` 地址池子网。
	IsRemoteVpcSnat *bool `json:"IsRemoteVpcSnat,omitnil" name:"IsRemoteVpcSnat"`

	// 子网`IPv4`总数。
	TotalIpAddressCount *uint64 `json:"TotalIpAddressCount,omitnil" name:"TotalIpAddressCount"`

	// 标签键值对。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// CDC实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// 是否是CDC所属子网。0:否 1:是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCdcSubnet *int64 `json:"IsCdcSubnet,omitnil" name:"IsCdcSubnet"`
}

type SubnetInput struct {
	// 子网的`CIDR`。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 子网名称。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// 可用区。形如：`ap-guangzhou-2`。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 指定关联路由表，形如：`rtb-3ryrwzuu`。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type TemplateLimit struct {
	// 参数模板IP地址成员配额。
	AddressTemplateMemberLimit *uint64 `json:"AddressTemplateMemberLimit,omitnil" name:"AddressTemplateMemberLimit"`

	// 参数模板IP地址组成员配额。
	AddressTemplateGroupMemberLimit *uint64 `json:"AddressTemplateGroupMemberLimit,omitnil" name:"AddressTemplateGroupMemberLimit"`

	// 参数模板I协议端口成员配额。
	ServiceTemplateMemberLimit *uint64 `json:"ServiceTemplateMemberLimit,omitnil" name:"ServiceTemplateMemberLimit"`

	// 参数模板协议端口组成员配额。
	ServiceTemplateGroupMemberLimit *uint64 `json:"ServiceTemplateGroupMemberLimit,omitnil" name:"ServiceTemplateGroupMemberLimit"`
}

type TrafficFlow struct {
	// 实际流量，单位为 字节
	Value *uint64 `json:"Value,omitnil" name:"Value"`

	// 格式化后的流量，单位见参数 FormatUnit
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormatValue *float64 `json:"FormatValue,omitnil" name:"FormatValue"`

	// 格式化后流量的单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormatUnit *string `json:"FormatUnit,omitnil" name:"FormatUnit"`
}

type TrafficPackage struct {
	// 流量包唯一ID
	TrafficPackageId *string `json:"TrafficPackageId,omitnil" name:"TrafficPackageId"`

	// 流量包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrafficPackageName *string `json:"TrafficPackageName,omitnil" name:"TrafficPackageName"`

	// 流量包总量，单位GB
	TotalAmount *float64 `json:"TotalAmount,omitnil" name:"TotalAmount"`

	// 流量包剩余量，单位GB
	RemainingAmount *float64 `json:"RemainingAmount,omitnil" name:"RemainingAmount"`

	// 流量包状态，可能的值有: AVAILABLE-可用状态， EXPIRED-已过期， EXHAUSTED-已用完， REFUNDED-已退还， DELETED-已删除
	Status *string `json:"Status,omitnil" name:"Status"`

	// 流量包创建时间
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 流量包截止时间
	Deadline *string `json:"Deadline,omitnil" name:"Deadline"`

	// 已使用的流量，单位GB
	UsedAmount *float64 `json:"UsedAmount,omitnil" name:"UsedAmount"`

	// 流量包标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 区分闲时流量包与全时流量包
	DeductType *string `json:"DeductType,omitnil" name:"DeductType"`
}

// Predefined struct for user
type TransformAddressRequestParams struct {
	// 待操作有普通公网 IP 的实例 ID。实例 ID 形如：`ins-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9389) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type TransformAddressRequest struct {
	*tchttp.BaseRequest
	
	// 待操作有普通公网 IP 的实例 ID。实例 ID 形如：`ins-11112222`。可通过登录[控制台](https://console.cloud.tencent.com/cvm)查询，也可通过 [DescribeInstances](https://cloud.tencent.com/document/api/213/9389) 接口返回值中的`InstanceId`获取。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *TransformAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransformAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TransformAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TransformAddressResponseParams struct {
	// 异步任务TaskId。可以使用[DescribeTaskResult](https://cloud.tencent.com/document/api/215/36271)接口查询任务状态。
	TaskId *uint64 `json:"TaskId,omitnil" name:"TaskId"`

	// 转为弹性公网IP后的唯一ID
	AddressId *string `json:"AddressId,omitnil" name:"AddressId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TransformAddressResponse struct {
	*tchttp.BaseResponse
	Response *TransformAddressResponseParams `json:"Response"`
}

func (r *TransformAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TransformAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6AddressesRequestParams struct {
	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的`IPv6`地址列表，单次最多指定10个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`
}

type UnassignIpv6AddressesRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例`ID`，形如：`eni-m6dyj72l`。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的`IPv6`地址列表，单次最多指定10个。
	Ipv6Addresses []*Ipv6Address `json:"Ipv6Addresses,omitnil" name:"Ipv6Addresses"`
}

func (r *UnassignIpv6AddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6AddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "Ipv6Addresses")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignIpv6AddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6AddressesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnassignIpv6AddressesResponse struct {
	*tchttp.BaseResponse
	Response *UnassignIpv6AddressesResponseParams `json:"Response"`
}

func (r *UnassignIpv6AddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6AddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6CidrBlockRequestParams struct {
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `IPv6`网段。形如：`3402:4e00:20:1000::/56`。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`
}

type UnassignIpv6CidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// `VPC`实例`ID`，形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `IPv6`网段。形如：`3402:4e00:20:1000::/56`。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`
}

func (r *UnassignIpv6CidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6CidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6CidrBlock")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignIpv6CidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6CidrBlockResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnassignIpv6CidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *UnassignIpv6CidrBlockResponseParams `json:"Response"`
}

func (r *UnassignIpv6CidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6CidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6SubnetCidrBlockRequestParams struct {
	// 子网所在私有网络`ID`。形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `IPv6` 子网段列表。
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil" name:"Ipv6SubnetCidrBlocks"`
}

type UnassignIpv6SubnetCidrBlockRequest struct {
	*tchttp.BaseRequest
	
	// 子网所在私有网络`ID`。形如：`vpc-f49l6u0z`。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `IPv6` 子网段列表。
	Ipv6SubnetCidrBlocks []*Ipv6SubnetCidrBlock `json:"Ipv6SubnetCidrBlocks,omitnil" name:"Ipv6SubnetCidrBlocks"`
}

func (r *UnassignIpv6SubnetCidrBlockRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6SubnetCidrBlockRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Ipv6SubnetCidrBlocks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignIpv6SubnetCidrBlockRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignIpv6SubnetCidrBlockResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnassignIpv6SubnetCidrBlockResponse struct {
	*tchttp.BaseResponse
	Response *UnassignIpv6SubnetCidrBlockResponseParams `json:"Response"`
}

func (r *UnassignIpv6SubnetCidrBlockResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignIpv6SubnetCidrBlockResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignPrivateIpAddressesRequestParams struct {
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 网卡绑定的子机实例ID，该参数仅用于指定网卡退还IP并解绑子机的场景，如果不涉及解绑子机，请勿填写。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type UnassignPrivateIpAddressesRequest struct {
	*tchttp.BaseRequest
	
	// 弹性网卡实例ID，例如：eni-m6dyj72l。
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitnil" name:"NetworkInterfaceId"`

	// 指定的内网IP信息，单次最多指定10个。
	PrivateIpAddresses []*PrivateIpAddressSpecification `json:"PrivateIpAddresses,omitnil" name:"PrivateIpAddresses"`

	// 网卡绑定的子机实例ID，该参数仅用于指定网卡退还IP并解绑子机的场景，如果不涉及解绑子机，请勿填写。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *UnassignPrivateIpAddressesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignPrivateIpAddressesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkInterfaceId")
	delete(f, "PrivateIpAddresses")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnassignPrivateIpAddressesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnassignPrivateIpAddressesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnassignPrivateIpAddressesResponse struct {
	*tchttp.BaseResponse
	Response *UnassignPrivateIpAddressesResponseParams `json:"Response"`
}

func (r *UnassignPrivateIpAddressesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnassignPrivateIpAddressesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockCcnBandwidthsRequestParams struct {
	// 带宽实例对象数组。
	Instances []*CcnFlowLock `json:"Instances,omitnil" name:"Instances"`
}

type UnlockCcnBandwidthsRequest struct {
	*tchttp.BaseRequest
	
	// 带宽实例对象数组。
	Instances []*CcnFlowLock `json:"Instances,omitnil" name:"Instances"`
}

func (r *UnlockCcnBandwidthsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockCcnBandwidthsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlockCcnBandwidthsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockCcnBandwidthsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnlockCcnBandwidthsResponse struct {
	*tchttp.BaseResponse
	Response *UnlockCcnBandwidthsResponseParams `json:"Response"`
}

func (r *UnlockCcnBandwidthsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockCcnBandwidthsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockCcnsRequestParams struct {

}

type UnlockCcnsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *UnlockCcnsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockCcnsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlockCcnsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockCcnsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnlockCcnsResponse struct {
	*tchttp.BaseResponse
	Response *UnlockCcnsResponseParams `json:"Response"`
}

func (r *UnlockCcnsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockCcnsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsedDetail struct {
	// 流量包唯一ID
	TrafficPackageId *string `json:"TrafficPackageId,omitnil" name:"TrafficPackageId"`

	// 流量包名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrafficPackageName *string `json:"TrafficPackageName,omitnil" name:"TrafficPackageName"`

	// 流量包总量
	TotalAmount *TrafficFlow `json:"TotalAmount,omitnil" name:"TotalAmount"`

	// 本次抵扣
	Deduction *TrafficFlow `json:"Deduction,omitnil" name:"Deduction"`

	// 本次抵扣后剩余量
	RemainingAmount *TrafficFlow `json:"RemainingAmount,omitnil" name:"RemainingAmount"`

	// 抵扣时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 资源类型。可能的值: CVM, LB, NAT, HAVIP, EIP
	ResourceType *string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 流量包到期时间
	Deadline *string `json:"Deadline,omitnil" name:"Deadline"`
}

type Vpc struct {
	// `VPC`名称。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// `VPC`实例`ID`，例如：vpc-azd4dt1c。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// `VPC`的`IPv4` `CIDR`。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 是否默认`VPC`。
	IsDefault *bool `json:"IsDefault,omitnil" name:"IsDefault"`

	// 是否开启组播。
	EnableMulticast *bool `json:"EnableMulticast,omitnil" name:"EnableMulticast"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// `DNS`列表。
	DnsServerSet []*string `json:"DnsServerSet,omitnil" name:"DnsServerSet"`

	// `DHCP`域名选项值。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// `DHCP`选项集`ID`。
	DhcpOptionsId *string `json:"DhcpOptionsId,omitnil" name:"DhcpOptionsId"`

	// 是否开启`DHCP`。
	EnableDhcp *bool `json:"EnableDhcp,omitnil" name:"EnableDhcp"`

	// `VPC`的`IPv6` `CIDR`。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil" name:"Ipv6CidrBlock"`

	// 标签键值对
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 辅助CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssistantCidrSet []*AssistantCidr `json:"AssistantCidrSet,omitnil" name:"AssistantCidrSet"`
}

type VpcEndPointServiceUser struct {
	// AppId。
	Owner *uint64 `json:"Owner,omitnil" name:"Owner"`

	// Uin。
	UserUin *string `json:"UserUin,omitnil" name:"UserUin"`

	// 描述信息。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 终端节点服务ID。
	EndPointServiceId *string `json:"EndPointServiceId,omitnil" name:"EndPointServiceId"`
}

type VpcIpv6Address struct {
	// `VPC`内`IPv6`地址。
	Ipv6Address *string `json:"Ipv6Address,omitnil" name:"Ipv6Address"`

	// 所属子网 `IPv6` `CIDR`。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// `IPv6`类型。
	Ipv6AddressType *string `json:"Ipv6AddressType,omitnil" name:"Ipv6AddressType"`

	// `IPv6`申请时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type VpcLimit struct {
	// 私有网络配额描述
	LimitType *string `json:"LimitType,omitnil" name:"LimitType"`

	// 私有网络配额值
	LimitValue *uint64 `json:"LimitValue,omitnil" name:"LimitValue"`
}

type VpcPrivateIpAddress struct {
	// `VPC`内网`IP`。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 所属子网`CIDR`。
	CidrBlock *string `json:"CidrBlock,omitnil" name:"CidrBlock"`

	// 内网`IP`类型。
	PrivateIpAddressType *string `json:"PrivateIpAddressType,omitnil" name:"PrivateIpAddressType"`

	// `IP`申请时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type VpcTaskResultDetailInfo struct {
	// 资源ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type VpnConnection struct {
	// 通道实例ID。
	VpnConnectionId *string `json:"VpnConnectionId,omitnil" name:"VpnConnectionId"`

	// 通道名称。
	VpnConnectionName *string `json:"VpnConnectionName,omitnil" name:"VpnConnectionName"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPN网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// 对端网关实例ID。
	CustomerGatewayId *string `json:"CustomerGatewayId,omitnil" name:"CustomerGatewayId"`

	// 预共享密钥。
	PreShareKey *string `json:"PreShareKey,omitnil" name:"PreShareKey"`

	// 通道传输协议。
	VpnProto *string `json:"VpnProto,omitnil" name:"VpnProto"`

	// 通道加密协议。
	EncryptProto *string `json:"EncryptProto,omitnil" name:"EncryptProto"`

	// 路由类型。
	RouteType *string `json:"RouteType,omitnil" name:"RouteType"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 通道的生产状态，PENDING：生产中，AVAILABLE：运行中，DELETING：删除中。
	State *string `json:"State,omitnil" name:"State"`

	// 通道连接状态，AVAILABLE：已连接。
	NetStatus *string `json:"NetStatus,omitnil" name:"NetStatus"`

	// SPD。
	SecurityPolicyDatabaseSet []*SecurityPolicyDatabase `json:"SecurityPolicyDatabaseSet,omitnil" name:"SecurityPolicyDatabaseSet"`

	// IKE选项。
	IKEOptionsSpecification *IKEOptionsSpecification `json:"IKEOptionsSpecification,omitnil" name:"IKEOptionsSpecification"`

	// IPSEC选择。
	IPSECOptionsSpecification *IPSECOptionsSpecification `json:"IPSECOptionsSpecification,omitnil" name:"IPSECOptionsSpecification"`

	// 是否支持健康状态探测
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil" name:"EnableHealthCheck"`

	// 本端探测ip
	HealthCheckLocalIp *string `json:"HealthCheckLocalIp,omitnil" name:"HealthCheckLocalIp"`

	// 对端探测ip
	HealthCheckRemoteIp *string `json:"HealthCheckRemoteIp,omitnil" name:"HealthCheckRemoteIp"`

	// 通道健康检查状态，AVAILABLE：正常，UNAVAILABLE：不正常。 未配置健康检查不返回该对象
	HealthCheckStatus *string `json:"HealthCheckStatus,omitnil" name:"HealthCheckStatus"`

	// DPD探测开关。默认为0，表示关闭DPD探测。可选值：0（关闭），1（开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DpdEnable *int64 `json:"DpdEnable,omitnil" name:"DpdEnable"`

	// DPD超时时间。即探测确认对端不存在需要的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DpdTimeout *string `json:"DpdTimeout,omitnil" name:"DpdTimeout"`

	// DPD超时后的动作。默认为clear。dpdEnable为1（开启）时有效。可取值为clear（断开）和restart（重试）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DpdAction *string `json:"DpdAction,omitnil" name:"DpdAction"`

	// 标签键值对数组
	TagSet []*Tag `json:"TagSet,omitnil" name:"TagSet"`

	// 协商类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NegotiationType *string `json:"NegotiationType,omitnil" name:"NegotiationType"`
}

type VpnGateway struct {
	// 网关实例ID。
	VpnGatewayId *string `json:"VpnGatewayId,omitnil" name:"VpnGatewayId"`

	// VPC实例ID。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 网关实例名称。
	VpnGatewayName *string `json:"VpnGatewayName,omitnil" name:"VpnGatewayName"`

	// 网关实例类型：'IPSEC', 'SSL','CCN','SSL_CCN'。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 网关实例状态， 'PENDING'：生产中，'PENDING_ERROR'：生产失败，'DELETING'：删除中，'DELETING_ERROR'：删除失败，'AVAILABLE'：运行中。
	State *string `json:"State,omitnil" name:"State"`

	// 网关公网IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil" name:"PublicIpAddress"`

	// 网关续费类型：'NOTIFY_AND_MANUAL_RENEW'：手动续费，'NOTIFY_AND_AUTO_RENEW'：自动续费，'NOT_NOTIFY_AND_NOT_RENEW'：到期不续费。
	RenewFlag *string `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 网关付费类型：POSTPAID_BY_HOUR：按量计费，PREPAID：包年包月预付费。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil" name:"InstanceChargeType"`

	// 网关出带宽。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 预付费网关过期时间。
	ExpiredTime *string `json:"ExpiredTime,omitnil" name:"ExpiredTime"`

	// 公网IP是否被封堵。
	IsAddressBlocked *bool `json:"IsAddressBlocked,omitnil" name:"IsAddressBlocked"`

	// 计费模式变更，PREPAID_TO_POSTPAID：包年包月预付费到期转按小时后付费。
	NewPurchasePlan *string `json:"NewPurchasePlan,omitnil" name:"NewPurchasePlan"`

	// 网关计费状态，PROTECTIVELY_ISOLATED：被安全隔离的实例，NORMAL：正常。
	RestrictState *string `json:"RestrictState,omitnil" name:"RestrictState"`

	// 可用区，如：ap-guangzhou-2。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 网关带宽配额信息。
	VpnGatewayQuotaSet []*VpnGatewayQuota `json:"VpnGatewayQuotaSet,omitnil" name:"VpnGatewayQuotaSet"`

	// 网关实例版本信息。
	Version *string `json:"Version,omitnil" name:"Version"`

	// Type值为CCN时，该值表示云联网实例ID。
	NetworkInstanceId *string `json:"NetworkInstanceId,omitnil" name:"NetworkInstanceId"`

	// CDC 实例ID。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// SSL-VPN 客户端连接数。
	MaxConnection *uint64 `json:"MaxConnection,omitnil" name:"MaxConnection"`
}

type VpnGatewayQuota struct {
	// 带宽配额
	Bandwidth *uint64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 配额中文名称
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 配额英文名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type VpnGatewayRoute struct {
	// 目的端IDC网段。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil" name:"DestinationCidrBlock"`

	// 下一跳类型（关联实例类型）可选值："VPNCONN"（VPN通道）， "CCN"（CCN实例）。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 下一跳实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 优先级，可选值：0，100。
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 启用状态，可选值："ENABLE"（启用），"DISABLE"  (禁用)。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 路由条目ID。
	RouteId *string `json:"RouteId,omitnil" name:"RouteId"`

	// 路由类型，可选值："VPC"（VPC路由），"CCN"（云联网传播路由），"Static"（静态路由），"BGP"（BGP路由）。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type VpnGatewayRouteModify struct {
	// VPN网关路由ID。
	RouteId *string `json:"RouteId,omitnil" name:"RouteId"`

	// VPN网关状态, ENABLE 启用, DISABLE禁用。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type VpngwCcnRoutes struct {
	// 路由信息ID。
	RouteId *string `json:"RouteId,omitnil" name:"RouteId"`

	// 路由信息是否启用。
	// ENABLE：启用该路由
	// DISABLE：不启用该路由
	Status *string `json:"Status,omitnil" name:"Status"`

	// 路由CIDR。
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitnil" name:"DestinationCidrBlock"`
}

// Predefined struct for user
type WithdrawNotifyRoutesRequestParams struct {
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略唯一ID。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

type WithdrawNotifyRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 路由表唯一ID。
	RouteTableId *string `json:"RouteTableId,omitnil" name:"RouteTableId"`

	// 路由策略唯一ID。
	RouteItemIds []*string `json:"RouteItemIds,omitnil" name:"RouteItemIds"`
}

func (r *WithdrawNotifyRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawNotifyRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RouteTableId")
	delete(f, "RouteItemIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WithdrawNotifyRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WithdrawNotifyRoutesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type WithdrawNotifyRoutesResponse struct {
	*tchttp.BaseResponse
	Response *WithdrawNotifyRoutesResponseParams `json:"Response"`
}

func (r *WithdrawNotifyRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WithdrawNotifyRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}