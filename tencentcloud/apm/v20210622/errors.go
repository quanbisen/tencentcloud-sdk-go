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

package v20210622

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// AuthFailure.UnauthorizedOperation
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// appid和实例信息不匹配。
	FAILEDOPERATION_APPIDNOTMATCHINSTANCEINFO = "FailedOperation.AppIdNotMatchInstanceInfo"

	// 实例ID为空。
	FAILEDOPERATION_INSTANCEIDISEMPTY = "FailedOperation.InstanceIdIsEmpty"

	// apm实例不存在。
	FAILEDOPERATION_INSTANCENOTFOUND = "FailedOperation.InstanceNotFound"

	// 非内网vpc。
	FAILEDOPERATION_NOTINNERVPC = "FailedOperation.NotInnerVPC"

	// 发送查询请求失败。
	FAILEDOPERATION_SENDREQUEST = "FailedOperation.SendRequest"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// filters中的字段不存在或非法。
	INVALIDPARAMETER_FILTERSFIELDSNOTEXISTORILLEGAL = "InvalidParameter.FiltersFieldsNotExistOrIllegal"

	// groupby中的字段不存在或非法。
	INVALIDPARAMETER_GROUPBYFIELDSNOTEXISTORILLEGAL = "InvalidParameter.GroupByFieldsNotExistOrIllegal"

	// filters中必须存在service.name字段，否则会报错。
	INVALIDPARAMETER_METRICFILTERSLACKPARAMS = "InvalidParameter.MetricFiltersLackParams"

	// metrics中的字段不存在或非法。
	INVALIDPARAMETER_METRICSFIELDNOTEXISTORILLEGAL = "InvalidParameter.MetricsFieldNotExistOrIllegal"

	// metrics中不允许为空。
	INVALIDPARAMETER_METRICSFIELDSNOTALLOWEMPTY = "InvalidParameter.MetricsFieldsNotAllowEmpty"

	// period不为空，0或60。
	INVALIDPARAMETER_PERIODISILLEGAL = "InvalidParameter.PeriodIsIllegal"

	// 查询时间不支持，最多只能查询最近2天、最多一个小时的数据。
	INVALIDPARAMETER_QUERYTIMEINTERVALISNOTSUPPORTED = "InvalidParameter.QueryTimeIntervalIsNotSupported"

	// 视图名称不存在或非法。
	INVALIDPARAMETER_VIEWNAMENOTEXISTORILLEGAL = "InvalidParameter.ViewNameNotExistOrIllegal"
)
