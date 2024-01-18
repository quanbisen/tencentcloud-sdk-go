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

package v20230308

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ConsumeGroupItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil" name:"ConsumeMessageOrderly"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil" name:"MaxRetryTimes"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

// Predefined struct for user
type CreateConsumerGroupRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil" name:"MaxRetryTimes"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil" name:"ConsumeMessageOrderly"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type CreateConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil" name:"MaxRetryTimes"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil" name:"ConsumeMessageOrderly"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *CreateConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	delete(f, "MaxRetryTimes")
	delete(f, "ConsumeEnable")
	delete(f, "ConsumeMessageOrderly")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerGroupResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsumerGroupResponseParams `json:"Response"`
}

func (r *CreateConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 商品规格，可用规格如下：
	// experiment_500,
	// basic_1k,
	// basic_2k,
	// basic_4k,
	// basic_6k
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil" name:"EnablePublic"`

	// 公网带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil" name:"IpRules"`

	// 消息保留时长，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil" name:"MessageRetention"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 商品规格，可用规格如下：
	// experiment_500,
	// basic_1k,
	// basic_2k,
	// basic_4k,
	// basic_6k
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil" name:"EnablePublic"`

	// 公网带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil" name:"IpRules"`

	// 消息保留时长，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil" name:"MessageRetention"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "Name")
	delete(f, "SkuCode")
	delete(f, "Remark")
	delete(f, "TagList")
	delete(f, "VpcList")
	delete(f, "EnablePublic")
	delete(f, "Bandwidth")
	delete(f, "IpRules")
	delete(f, "MessageRetention")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTInstanceRequestParams struct {
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 商品规格，可用规格如下：
	// experiment_500,
	// basic_1k,
	// basic_2k,
	// basic_4k,
	// basic_6k,
	// pro_4k,
	// pro_6k,
	// pro_1w,
	// pro_2w,
	// pro_3w,
	// pro_4w,
	// pro_5w,
	// platinum_6k,
	// platinum_1w,
	// platinum_2w,
	// platinum_4w,
	// platinum_10w,
	// platinum_15w,
	// platinum_20w,
	// platinum_40w,
	// platinum_60w,
	// platinum_100w
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil" name:"EnablePublic"`

	// 公网带宽（单位：兆）
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil" name:"IpRules"`

	// 是否自动续费（0: 不自动续费；1: 自动续费）
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 购买时长（单位：月）
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`
}

type CreateMQTTInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 商品规格，可用规格如下：
	// experiment_500,
	// basic_1k,
	// basic_2k,
	// basic_4k,
	// basic_6k,
	// pro_4k,
	// pro_6k,
	// pro_1w,
	// pro_2w,
	// pro_3w,
	// pro_4w,
	// pro_5w,
	// platinum_6k,
	// platinum_1w,
	// platinum_2w,
	// platinum_4w,
	// platinum_10w,
	// platinum_15w,
	// platinum_20w,
	// platinum_40w,
	// platinum_60w,
	// platinum_100w
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil" name:"EnablePublic"`

	// 公网带宽（单位：兆）
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil" name:"IpRules"`

	// 是否自动续费（0: 不自动续费；1: 自动续费）
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 购买时长（单位：月）
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`
}

func (r *CreateMQTTInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "Name")
	delete(f, "SkuCode")
	delete(f, "Remark")
	delete(f, "TagList")
	delete(f, "VpcList")
	delete(f, "EnablePublic")
	delete(f, "Bandwidth")
	delete(f, "IpRules")
	delete(f, "RenewFlag")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMQTTInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateMQTTInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateMQTTInstanceResponseParams `json:"Response"`
}

func (r *CreateMQTTInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色名称
	Role *string `json:"Role,omitnil" name:"Role"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否开启生产权限
	PermWrite *bool `json:"PermWrite,omitnil" name:"PermWrite"`

	// 是否开启消费权限
	PermRead *bool `json:"PermRead,omitnil" name:"PermRead"`
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色名称
	Role *string `json:"Role,omitnil" name:"Role"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否开启生产权限
	PermWrite *bool `json:"PermWrite,omitnil" name:"PermWrite"`

	// 是否开启消费权限
	PermRead *bool `json:"PermRead,omitnil" name:"PermRead"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Role")
	delete(f, "Remark")
	delete(f, "PermWrite")
	delete(f, "PermRead")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleResponseParams struct {
	// 角色名
	Role *string `json:"Role,omitnil" name:"Role"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleResponseParams `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 主题类型
	// UNSPECIFIED:未指定,
	// NORMAL:普通消息,
	// FIFO:顺序消息,
	// DELAY:延时消息,
	// TRANSACTION:事务消息
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// 队列数量
	QueueNum *int64 `json:"QueueNum,omitnil" name:"QueueNum"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 主题类型
	// UNSPECIFIED:未指定,
	// NORMAL:普通消息,
	// FIFO:顺序消息,
	// DELAY:延时消息,
	// TRANSACTION:事务消息
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// 队列数量
	QueueNum *int64 `json:"QueueNum,omitnil" name:"QueueNum"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "TopicType")
	delete(f, "QueueNum")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsumerGroupRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`
}

type DeleteConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`
}

func (r *DeleteConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsumerGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConsumerGroupResponseParams `json:"Response"`
}

func (r *DeleteConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色名称
	Role *string `json:"Role,omitnil" name:"Role"`
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色名称
	Role *string `json:"Role,omitnil" name:"Role"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleResponseParams `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicResponseParams `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询指定主题下的消费组
	FromTopic *string `json:"FromTopic,omitnil" name:"FromTopic"`
}

type DescribeConsumerGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询指定主题下的消费组
	FromTopic *string `json:"FromTopic,omitnil" name:"FromTopic"`
}

func (r *DescribeConsumerGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "FromTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 消费组列表
	Data []*ConsumeGroupItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeConsumerGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerGroupListResponseParams `json:"Response"`
}

func (r *DescribeConsumerGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`
}

func (r *DescribeConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupResponseParams struct {
	// 在线消费者数量
	ConsumerNum *int64 `json:"ConsumerNum,omitnil" name:"ConsumerNum"`

	// TPS
	Tps *int64 `json:"Tps,omitnil" name:"Tps"`

	// 消息堆积数量
	ConsumerLag *int64 `json:"ConsumerLag,omitnil" name:"ConsumerLag"`

	// 消费者类型
	ConsumeType *string `json:"ConsumeType,omitnil" name:"ConsumeType"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil" name:"ConsumeMessageOrderly"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil" name:"ConsumeEnable"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil" name:"MaxRetryTimes"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerGroupResponseParams `json:"Response"`
}

func (r *DescribeConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例列表
	Data []*InstanceItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil" name:"TopicNum"`

	// 实例最大主题数量
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil" name:"TopicNumLimit"`

	// 消费组数量
	GroupNum *int64 `json:"GroupNum,omitnil" name:"GroupNum"`

	// 实例最大消费组数量
	GroupNumLimit *int64 `json:"GroupNumLimit,omitnil" name:"GroupNumLimit"`

	// 消息保留时间，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil" name:"MessageRetention"`

	// 最大可调整消息保留时间，小时为单位
	RetentionUpperLimit *int64 `json:"RetentionUpperLimit,omitnil" name:"RetentionUpperLimit"`

	// 最小可调整消息保留时间，小时为单位
	RetentionLowerLimit *int64 `json:"RetentionLowerLimit,omitnil" name:"RetentionLowerLimit"`

	// TPS限流值
	TpsLimit *int64 `json:"TpsLimit,omitnil" name:"TpsLimit"`

	// 弹性TPS限流值
	ScaledTpsLimit *int64 `json:"ScaledTpsLimit,omitnil" name:"ScaledTpsLimit"`

	// 延迟消息最长时间，小时为单位
	MaxMessageDelay *int64 `json:"MaxMessageDelay,omitnil" name:"MaxMessageDelay"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 消息发送接收比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil" name:"SendReceiveRatio"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitnil" name:"TagList"`

	// 接入点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointList []*Endpoint `json:"EndpointList,omitnil" name:"EndpointList"`

	// 主题队列数上限
	TopicQueueNumUpperLimit *int64 `json:"TopicQueueNumUpperLimit,omitnil" name:"TopicQueueNumUpperLimit"`

	// 主题队列数下限
	TopicQueueNumLowerLimit *int64 `json:"TopicQueueNumLowerLimit,omitnil" name:"TopicQueueNumLowerLimit"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// 实例规格
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 计费模式
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceListRequestParams struct {
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeMQTTInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeMQTTInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例列表
	Data []*MQTTInstanceItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMQTTInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTInstanceListResponseParams `json:"Response"`
}

func (r *DescribeMQTTInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTProductSKUListRequestParams struct {

}

type DescribeMQTTProductSKUListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeMQTTProductSKUListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTProductSKUListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTProductSKUListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTProductSKUListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// mqtt商品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MQTTProductSkuList []*MQTTProductSkuItem `json:"MQTTProductSkuList,omitnil" name:"MQTTProductSkuList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMQTTProductSKUListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTProductSKUListResponseParams `json:"Response"`
}

func (r *DescribeMQTTProductSKUListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTProductSKUListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 角色信息列表
	Data []*RoleItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoleListResponseParams `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 主题列表
	Data []*TopicItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicListResponseParams `json:"Response"`
}

func (r *DescribeTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 主题类型
	// UNSPECIFIED:未指定,
	// NORMAL:普通消息,
	// FIFO:顺序消息,
	// DELAY:延时消息,
	// TRANSACTION:事务消息
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 最后写入时间，秒为单位
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil" name:"LastUpdateTime"`

	// 订阅数量
	SubscriptionCount *int64 `json:"SubscriptionCount,omitnil" name:"SubscriptionCount"`

	// 订阅关系列表
	SubscriptionData []*SubscriptionData `json:"SubscriptionData,omitnil" name:"SubscriptionData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicResponseParams `json:"Response"`
}

func (r *DescribeTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicStatsOpRequestParams struct {
	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`
}

type DescribeTopicStatsOpRequest struct {
	*tchttp.BaseRequest
	
	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`
}

func (r *DescribeTopicStatsOpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicStatsOpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicStatsOpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicStatsOpResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopicStatsOpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicStatsOpResponseParams `json:"Response"`
}

func (r *DescribeTopicStatsOpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicStatsOpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Endpoint struct {
	// 接入点类型，
	// VPC，
	// PUBLIC 公网
	Type *string `json:"Type,omitnil" name:"Type"`

	// 状态，
	// OPEN 开启，
	// CLOSE 关闭，
	// PROCESSING 操作中，
	Status *string `json:"Status,omitnil" name:"Status"`

	// 付费类型，仅公网
	// PREPAID 包年包月
	// POSTPAID 按量付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 接入点地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointUrl *string `json:"EndpointUrl,omitnil" name:"EndpointUrl"`

	// VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 公网带宽，Mbps为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *int64 `json:"Bandwidth,omitnil" name:"Bandwidth"`

	// 公网放通规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpRules []*IpRule `json:"IpRules,omitnil" name:"IpRules"`
}

type Filter struct {
	// 过滤条件名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤条件的值
	Values []*string `json:"Values,omitnil" name:"Values"`
}

// Predefined struct for user
type ImportSourceClusterConsumerGroupsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 待导入的消费组列表
	GroupList []*SourceClusterGroupConfig `json:"GroupList,omitnil" name:"GroupList"`
}

type ImportSourceClusterConsumerGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 待导入的消费组列表
	GroupList []*SourceClusterGroupConfig `json:"GroupList,omitnil" name:"GroupList"`
}

func (r *ImportSourceClusterConsumerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterConsumerGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "GroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportSourceClusterConsumerGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSourceClusterConsumerGroupsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImportSourceClusterConsumerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ImportSourceClusterConsumerGroupsResponseParams `json:"Response"`
}

func (r *ImportSourceClusterConsumerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterConsumerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSourceClusterTopicsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 待导入的主题列表
	TopicList []*SourceClusterTopicConfig `json:"TopicList,omitnil" name:"TopicList"`
}

type ImportSourceClusterTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 待导入的主题列表
	TopicList []*SourceClusterTopicConfig `json:"TopicList,omitnil" name:"TopicList"`
}

func (r *ImportSourceClusterTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TopicList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportSourceClusterTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSourceClusterTopicsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ImportSourceClusterTopicsResponse struct {
	*tchttp.BaseResponse
	Response *ImportSourceClusterTopicsResponseParams `json:"Response"`
}

func (r *ImportSourceClusterTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 实例类型，
	// EXPERIMENT，体验版
	// BASIC，基础版
	// PRO，专业版
	// PLATINUM，铂金版
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例状态，
	// RUNNING, 运行中
	// MAINTAINING，维护中
	// ABNORMAL，异常
	// OVERDUE，欠费
	// DESTROYED，已删除
	// CREATING，创建中
	// MODIFYING，变配中
	// CREATE_FAILURE，创建失败
	// MODIFY_FAILURE，变配失败
	// DELETING，删除中
	InstanceStatus *string `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// 实例主题数上限
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil" name:"TopicNumLimit"`

	// 实例消费组数量上限
	GroupNumLimit *int64 `json:"GroupNumLimit,omitnil" name:"GroupNumLimit"`

	// 计费模式，
	// POSTPAID，按量计费
	// PREPAID，包年包月
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 到期时间，秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiryTime *int64 `json:"ExpiryTime,omitnil" name:"ExpiryTime"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil" name:"TopicNum"`

	// 消费组数量
	GroupNum *int64 `json:"GroupNum,omitnil" name:"GroupNum"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitnil" name:"TagList"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitnil" name:"TpsLimit"`

	// 弹性TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaledTpsLimit *int64 `json:"ScaledTpsLimit,omitnil" name:"ScaledTpsLimit"`

	// 消息保留时间，小时为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageRetention *int64 `json:"MessageRetention,omitnil" name:"MessageRetention"`

	// 延迟消息最大时长，小时为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMessageDelay *int64 `json:"MaxMessageDelay,omitnil" name:"MaxMessageDelay"`
}

type IpRule struct {
	// IP地址
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 是否允许放行
	Allow *bool `json:"Allow,omitnil" name:"Allow"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type MQTTInstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitnil" name:"Version"`

	// 实例类型，
	// EXPERIMENT，体验版
	// BASIC，基础版
	// PRO，专业版
	// PLATINUM，铂金版
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 实例状态，
	// RUNNING, 运行中
	// MAINTAINING，维护中
	// ABNORMAL，异常
	// OVERDUE，欠费
	// DESTROYED，已删除
	// CREATING，创建中
	// MODIFYING，变配中
	// CREATE_FAILURE，创建失败
	// MODIFY_FAILURE，变配失败
	// DELETING，删除中
	InstanceStatus *string `json:"InstanceStatus,omitnil" name:"InstanceStatus"`

	// 实例主题数上限
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil" name:"TopicNumLimit"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil" name:"TopicNum"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 弹性TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitnil" name:"TpsLimit"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// 订阅关系上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionNumLimit *int64 `json:"SubscriptionNumLimit,omitnil" name:"SubscriptionNumLimit"`

	// 客户端连接数上线
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil" name:"ClientNumLimit"`
}

type MQTTProductSkuItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// cide
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// sale
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnSale *bool `json:"OnSale,omitnil" name:"OnSale"`

	// topic num限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil" name:"TopicNumLimit"`

	// tps
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitnil" name:"TpsLimit"`

	// 客户端连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil" name:"ClientNumLimit"`

	// 订阅数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionNumLimit *int64 `json:"SubscriptionNumLimit,omitnil" name:"SubscriptionNumLimit"`

	// 代理核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySpecCore *int64 `json:"ProxySpecCore,omitnil" name:"ProxySpecCore"`

	// 代理内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySpecMemory *int64 `json:"ProxySpecMemory,omitnil" name:"ProxySpecMemory"`

	// 代理总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySpecCount *int64 `json:"ProxySpecCount,omitnil" name:"ProxySpecCount"`
}

// Predefined struct for user
type ModifyConsumerGroupRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil" name:"ConsumeMessageOrderly"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil" name:"MaxRetryTimes"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifyConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil" name:"ConsumeMessageOrderly"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil" name:"MaxRetryTimes"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifyConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	delete(f, "ConsumeEnable")
	delete(f, "ConsumeMessageOrderly")
	delete(f, "MaxRetryTimes")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsumerGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConsumerGroupResponseParams `json:"Response"`
}

func (r *ModifyConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 消息发送和接收的比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil" name:"SendReceiveRatio"`

	// 调整实例规格的商品代号
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 消息保留时长，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil" name:"MessageRetention"`

	// 是否开启弹性TPS
	ScaledTpsEnabled *bool `json:"ScaledTpsEnabled,omitnil" name:"ScaledTpsEnabled"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 消息发送和接收的比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil" name:"SendReceiveRatio"`

	// 调整实例规格的商品代号
	SkuCode *string `json:"SkuCode,omitnil" name:"SkuCode"`

	// 消息保留时长，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil" name:"MessageRetention"`

	// 是否开启弹性TPS
	ScaledTpsEnabled *bool `json:"ScaledTpsEnabled,omitnil" name:"ScaledTpsEnabled"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "SendReceiveRatio")
	delete(f, "SkuCode")
	delete(f, "MessageRetention")
	delete(f, "ScaledTpsEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoleRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色名称
	Role *string `json:"Role,omitnil" name:"Role"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil" name:"PermWrite"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 角色名称
	Role *string `json:"Role,omitnil" name:"Role"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil" name:"PermWrite"`

	// 备注
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifyRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Role")
	delete(f, "PermRead")
	delete(f, "PermWrite")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoleResponseParams `json:"Response"`
}

func (r *ModifyRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 队列数量
	QueueNum *int64 `json:"QueueNum,omitnil" name:"QueueNum"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 队列数量
	QueueNum *int64 `json:"QueueNum,omitnil" name:"QueueNum"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "QueueNum")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicResponseParams `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleItem struct {
	// 角色名称
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Access Key
	AccessKey *string `json:"AccessKey,omitnil" name:"AccessKey"`

	// Secret Key
	SecretKey *string `json:"SecretKey,omitnil" name:"SecretKey"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil" name:"PermWrite"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 修改时间，秒为单位
	ModifiedTime *int64 `json:"ModifiedTime,omitnil" name:"ModifiedTime"`
}

type SourceClusterGroupConfig struct {
	// 消费组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否已导入，作为入参时无效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Imported *bool `json:"Imported,omitnil" name:"Imported"`

	// 命名空间，仅4.x集群有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 导入状态
	// Unknown 未知
	// Success 成功
	// Failure 失败
	// AlreadyExists 已存在
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportStatus *string `json:"ImportStatus,omitnil" name:"ImportStatus"`
}

type SourceClusterTopicConfig struct {
	// 主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 主题类型，
	// 5.x版本
	// UNSPECIFIED 未指定
	// NORMAL 普通消息
	// FIFO 顺序消息
	// DELAY 延迟消息
	// TRANSACTION 事务消息
	// 
	// 4.x版本
	// Normal 普通消息
	// PartitionedOrder 分区顺序消息
	// Transaction 事务消息
	// DelayScheduled 延时消息
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// 队列数
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueNum *int64 `json:"QueueNum,omitnil" name:"QueueNum"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否已导入，作为入参时无效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Imported *bool `json:"Imported,omitnil" name:"Imported"`

	// 命名空间，仅4.x集群有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 导入状态，
	// Unknown 未知，
	// AlreadyExists 已存在，
	// Success 成功，
	// Failure 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportStatus *string `json:"ImportStatus,omitnil" name:"ImportStatus"`
}

type SubscriptionData struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 主题类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// 单个节点上主题队列数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicQueueNum *int64 `json:"TopicQueueNum,omitnil" name:"TopicQueueNum"`

	// 消费组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil" name:"ConsumerGroup"`

	// 是否在线
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOnline *bool `json:"IsOnline,omitnil" name:"IsOnline"`

	// 消费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumeType *string `json:"ConsumeType,omitnil" name:"ConsumeType"`

	// 订阅规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubString *string `json:"SubString,omitnil" name:"SubString"`

	// 过滤类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpressionType *string `json:"ExpressionType,omitnil" name:"ExpressionType"`

	// 订阅一致性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *int64 `json:"Consistency,omitnil" name:"Consistency"`

	// 消费堆积
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLag *int64 `json:"ConsumerLag,omitnil" name:"ConsumerLag"`

	// 最后消费进度更新时间，秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil" name:"LastUpdateTime"`

	// 最大重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil" name:"MaxRetryTimes"`

	// 是否顺序消费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil" name:"ConsumeMessageOrderly"`
}

type Tag struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type TagFilter struct {
	// 标签键名称
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值列表
	TagValues []*string `json:"TagValues,omitnil" name:"TagValues"`
}

type TopicItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// 主题类型
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// 队列数量
	QueueNum *int64 `json:"QueueNum,omitnil" name:"QueueNum"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type VpcInfo struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}