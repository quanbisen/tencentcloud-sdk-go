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

package v20210125

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 订单状态异常。
	FAILEDOPERATION_ABNORMALORDERSTATUS = "FailedOperation.AbnormalOrderStatus"

	// 另一个数据源正在创建中。
	FAILEDOPERATION_ANOTHERCREATEPROCESSRUNNING = "FailedOperation.AnotherCreateProcessRunning"

	// 另一个操作正在处理中，请稍后再试。
	FAILEDOPERATION_ANOTHERPROCESSRUNNING = "FailedOperation.AnotherProcessRunning"

	// 另一个请求正在处理中，请稍后再试。
	FAILEDOPERATION_ANOTHERREQUESTPROCESSING = "FailedOperation.AnotherRequestProcessing"

	// 账户余额不足。
	FAILEDOPERATION_BALANCENOTENOUGH = "FailedOperation.BalanceNotEnough"

	// 计费系统异常。
	FAILEDOPERATION_BILLINGSYSTEMERROR = "FailedOperation.BillingSystemError"

	// 绑定的标签数量超出限制。
	FAILEDOPERATION_BINDTOOMANYTAGS = "FailedOperation.BindTooManyTags"

	// 创建引擎失败。
	FAILEDOPERATION_CREATEDATAENGINEFAILED = "FailedOperation.CreateDataEngineFailed"

	// 删除数据引擎失败。
	FAILEDOPERATION_DELETEDATAENGINEFAILED = "FailedOperation.DeleteDataEngineFailed"

	// 发货失败。
	FAILEDOPERATION_DELIVERGOODSFAILED = "FailedOperation.DeliverGoodsFailed"

	// 重复的标签键。
	FAILEDOPERATION_DUPLICATETAGKEY = "FailedOperation.DuplicateTagKey"

	// 扣费失败。
	FAILEDOPERATION_FEEDEDUCTIONFAILED = "FailedOperation.FeeDeductionFailed"

	// 获取鉴权策略失败。
	FAILEDOPERATION_GETPOLICYFAILED = "FailedOperation.GetPolicyFailed"

	// 获取商品信息失败。
	FAILEDOPERATION_GETPRODUCTINFORMATIONFAILED = "FailedOperation.GetProductInformationFailed"

	// 获取用户信息失败，请重试或提工单联系我们
	FAILEDOPERATION_GETUSERINFOFAILED = "FailedOperation.GetUserInfoFailed"

	// 获取工作组信息失败。
	FAILEDOPERATION_GETWORKGROUPINFOFAILED = "FailedOperation.GetWorkGroupInfoFailed"

	// 授权失败。
	FAILEDOPERATION_GRANTPOLICYFAILED = "FailedOperation.GrantPolicyFailed"

	// HTTP客户端请求失败。
	FAILEDOPERATION_HTTPCLIENTDOREQUESTFAILED = "FailedOperation.HttpClientDoRequestFailed"

	// 资源不符合规定。
	FAILEDOPERATION_ILLEGALRESOURCE = "FailedOperation.IllegalResource"

	// 标签键含有非法字符。
	FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"

	// 标签值含有非法字符。
	FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"

	// 询价失败。
	FAILEDOPERATION_INQUIREPRICEFAILED = "FailedOperation.InquirePriceFailed"

	// 元数据错误，请重试，或者提交工单联系我们
	FAILEDOPERATION_METASTOREERROR = "FailedOperation.MetastoreError"

	// 实例变配失败。
	FAILEDOPERATION_MODIFYINSTANCEFAILED = "FailedOperation.ModifyInstanceFailed"

	// 没有操作权限。
	FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"

	// 用户没有指定引擎的使用权限
	FAILEDOPERATION_NOPERMISSIONTOUSETHEDATAENGINE = "FailedOperation.NoPermissionToUseTheDataEngine"

	// 账号未进行实名认证。
	FAILEDOPERATION_NOREALNAMEAUTHENTICATION = "FailedOperation.NoRealNameAuthentication"

	// 采购数量超过限制。
	FAILEDOPERATION_NUMBEREXCEEDLIMIT = "FailedOperation.NumberExceedLimit"

	// 参数校验失败。
	FAILEDOPERATION_PARAMETERVALIDATIONFAILED = "FailedOperation.ParameterValidationFailed"

	// 退押金失败。
	FAILEDOPERATION_REFUNDDEPOSITFAILED = "FailedOperation.RefundDepositFailed"

	// 取消授权失败。
	FAILEDOPERATION_REVOKEPOLICYFAILED = "FailedOperation.RevokePolicyFailed"

	// 资源已经绑定了同名标签键。
	FAILEDOPERATION_TAGALREADYATTACHED = "FailedOperation.TagAlreadyAttached"

	// 标签键长度超过限制。
	FAILEDOPERATION_TAGKEYTOOLONG = "FailedOperation.TagKeyTooLong"

	// 标签不存在。
	FAILEDOPERATION_TAGNOTEXIST = "FailedOperation.TagNotExist"

	// 标签值长度超过限制。
	FAILEDOPERATION_TAGVALUETOOLONG = "FailedOperation.TagValueTooLong"

	// 资源数量超出限制。
	FAILEDOPERATION_TOOMANYRESOURCES = "FailedOperation.TooManyResources"

	// 标签数量超出限制。
	FAILEDOPERATION_TOOMANYTAGS = "FailedOperation.TooManyTags"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据库错误。
	INTERNALERROR_DBERROR = "InternalError.DBError"

	// 业务系统异常，请重试或提工单联系我们。
	INTERNALERROR_INTERNALSYSTEMEXCEPTION = "InternalError.InternalSystemException"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 指定集群参数已存在
	INVALIDPARAMETER_DATAENGINECONFIGPAIRSDUPLICATE = "InvalidParameter.DataEngineConfigPairsDuplicate"

	// 指定集群传参ExecType不匹配，当前仅支持SQL或BATCH
	INVALIDPARAMETER_DATAENGINEEXECTYPENOTMATCH = "InvalidParameter.DataEngineExecTypeNotMatch"

	// 指定集群镜像操作不匹配，当前仅支持: InitImage/UpgradeImage/SwitchImage/RollbackImage/ModifyResource
	INVALIDPARAMETER_DATAENGINEIMAGEOPERATENOTMATCH = "InvalidParameter.DataEngineImageOperateNotMatch"

	// 指定集群付费类型不匹配，当前仅支持: 0: 后付费, 1: 预付费
	INVALIDPARAMETER_DATAENGINEPAYMODETYPENOTMATCH = "InvalidParameter.DataEnginePayModeTypeNotMatch"

	// 指定集群类型不匹配，当前仅支持: spark/presto
	INVALIDPARAMETER_DATAENGINETYPENOTMATCH = "InvalidParameter.DataEngineTypeNotMatch"

	// 数据源类型错误。
	INVALIDPARAMETER_DATASOURCETYPEERROR = "InvalidParameter.DatasourceTypeError"

	// 重复的引擎名称。
	INVALIDPARAMETER_DUPLICATEDATAENGINENAME = "InvalidParameter.DuplicateDataEngineName"

	// 重复的工作组名称。
	INVALIDPARAMETER_DUPLICATEGROUPNAME = "InvalidParameter.DuplicateGroupName"

	// 重复的用户名。
	INVALIDPARAMETER_DUPLICATEUSERNAME = "InvalidParameter.DuplicateUserName"

	// 指定集群镜像Cluster参数格式非JSON
	INVALIDPARAMETER_IMAGECLUSTERPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageClusterParametersFormatNotJson"

	// 指定引擎类型不匹配，当前仅支持: SparkSQL, PrestoSQL, SparkBatch
	INVALIDPARAMETER_IMAGEENGINETYPENOTMATCH = "InvalidParameter.ImageEngineTypeNotMatch"

	// 指定isPublic不匹配，当前仅支持: 1:公共, 2:私有
	INVALIDPARAMETER_IMAGEISPUBLICNOTMATCH = "InvalidParameter.ImageIsPublicNotMatch"

	// 指定集群镜像参数不存在
	INVALIDPARAMETER_IMAGEPARAMETERNOTFOUND = "InvalidParameter.ImageParameterNotFound"

	// 指定集群镜像ParameterSubmitMethod不匹配，当前仅支持: User, BackGround
	INVALIDPARAMETER_IMAGEPARAMETERSUBMITMETHODNOTMATCH = "InvalidParameter.ImageParameterSubmitMethodNotMatch"

	// 指定集群镜像ParameterType不匹配，当前仅支持: 1: session , 2: common, 3: cluster
	INVALIDPARAMETER_IMAGEPARAMETERTYPENOTMATCH = "InvalidParameter.ImageParameterTypeNotMatch"

	// 指定集群镜像Session参数格式非JSON
	INVALIDPARAMETER_IMAGESESSIONPARAMETERSFORMATNOTJSON = "InvalidParameter.ImageSessionParametersFormatNotJson"

	// 指定state不匹配，当前仅支持: 1:初始化, 2:上线, 3:下线
	INVALIDPARAMETER_IMAGESTATENOTMATCH = "InvalidParameter.ImageStateNotMatch"

	// 指定集群镜像UserRecords不匹配，当前仅支持: 1: parentVersion, 2: childVersion, 3: pySpark
	INVALIDPARAMETER_IMAGEUSERRECORDSTYPENOTMATCH = "InvalidParameter.ImageUserRecordsTypeNotMatch"

	// 实例在其他流程中。
	INVALIDPARAMETER_INSTANCEINPROCESS = "InvalidParameter.InstanceInProcess"

	// 无效的访问策略。
	INVALIDPARAMETER_INVALIDACCESSPOLICY = "InvalidParameter.InvalidAccessPolicy"

	// 指定集群参数无效，请校验后重试
	INVALIDPARAMETER_INVALIDDATAENGINECONFIGPAIRS = "InvalidParameter.InvalidDataEngineConfigPairs"

	// 无效的引擎描述信息。
	INVALIDPARAMETER_INVALIDDATAENGINEDESCRIPTION = "InvalidParameter.InvalidDataEngineDescription"

	// 无效的数据引擎模式。
	INVALIDPARAMETER_INVALIDDATAENGINEMODE = "InvalidParameter.InvalidDataEngineMode"

	// 无效的数据引擎名。
	INVALIDPARAMETER_INVALIDDATAENGINENAME = "InvalidParameter.InvalidDataEngineName"

	// 无效的数据引擎规格。
	INVALIDPARAMETER_INVALIDDATAENGINESPECS = "InvalidParameter.InvalidDataEngineSpecs"

	// 无效的默认数据引擎参数。
	INVALIDPARAMETER_INVALIDDEFAULTDATAENGINE = "InvalidParameter.InvalidDefaultDataEngine"

	// 无效的描述信息。
	INVALIDPARAMETER_INVALIDDESCRIPTION = "InvalidParameter.InvalidDescription"

	// 引擎类型不合法
	INVALIDPARAMETER_INVALIDENGINETYPE = "InvalidParameter.InvalidEngineType"

	// 任务容错类型错误，当前仅支持: Proceed/Terminate
	INVALIDPARAMETER_INVALIDFAILURETOLERANCE = "InvalidParameter.InvalidFailureTolerance"

	// 不支持此过滤条件
	INVALIDPARAMETER_INVALIDFILTERKEY = "InvalidParameter.InvalidFilterKey"

	// 无效的工作组Id。
	INVALIDPARAMETER_INVALIDGROUPID = "InvalidParameter.InvalidGroupId"

	// 请求的消息类型无效。
	INVALIDPARAMETER_INVALIDINFOTYPE = "InvalidParameter.InvalidInfoType"

	// 无效的最大结果数。
	INVALIDPARAMETER_INVALIDMAXRESULTS = "InvalidParameter.InvalidMaxResults"

	// 无效的最小集群数量。
	INVALIDPARAMETER_INVALIDMINCLUSTERS = "InvalidParameter.InvalidMinClusters"

	// 无效的Offset值。
	INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"

	// 无效的计费模式。
	INVALIDPARAMETER_INVALIDPAYMODE = "InvalidParameter.InvalidPayMode"

	// 无效的CAM role arn。
	INVALIDPARAMETER_INVALIDROLEARN = "InvalidParameter.InvalidRoleArn"

	// SQL解析失败。
	INVALIDPARAMETER_INVALIDSQL = "InvalidParameter.InvalidSQL"

	// 执行SQL数量错误，SQL数量要大于等于1个且小于等于50个
	INVALIDPARAMETER_INVALIDSQLNUM = "InvalidParameter.InvalidSQLNum"

	// 不支持的排序类型。
	INVALIDPARAMETER_INVALIDSORTBYTYPE = "InvalidParameter.InvalidSortByType"

	// SparkAppParam无效。
	INVALIDPARAMETER_INVALIDSPARKAPPPARAM = "InvalidParameter.InvalidSparkAppParam"

	// 存储位置错误。
	INVALIDPARAMETER_INVALIDSTORELOCATION = "InvalidParameter.InvalidStoreLocation"

	// 无效的taskid。
	INVALIDPARAMETER_INVALIDTASKID = "InvalidParameter.InvalidTaskId"

	// 任务类型TaskType错误，Spark引擎任务类型为SparkSQLTask,Presto引擎任务类型为SQLTask
	INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"

	// 无效的计费时长。
	INVALIDPARAMETER_INVALIDTIMESPAN = "InvalidParameter.InvalidTimeSpan"

	// 无效的计费时长单位。
	INVALIDPARAMETER_INVALIDTIMEUNIT = "InvalidParameter.InvalidTimeUnit"

	// 无效用户名称。
	INVALIDPARAMETER_INVALIDUSERALIAS = "InvalidParameter.InvalidUserAlias"

	// 无效的用户名。
	INVALIDPARAMETER_INVALIDUSERNAME = "InvalidParameter.InvalidUserName"

	// 无效的用户类型。
	INVALIDPARAMETER_INVALIDUSERTYPE = "InvalidParameter.InvalidUserType"

	// 获取白名单错误，请重试，或者提交工单联系我们
	INVALIDPARAMETER_INVALIDWHITELISTKEY = "InvalidParameter.InvalidWhiteListKey"

	// 无效的工作组名。
	INVALIDPARAMETER_INVALIDWORKGROUPNAME = "InvalidParameter.InvalidWorkGroupName"

	// 找不到参数或参数为空
	INVALIDPARAMETER_PARAMETERNOTFOUNDORBENONE = "InvalidParameter.ParameterNotFoundOrBeNone"

	// 任务已经结束，不能取消。
	INVALIDPARAMETER_TASKALREADYFINISHED = "InvalidParameter.TaskAlreadyFinished"

	// Vpc cidr格式错误。
	INVALIDPARAMETER_VPCCIDRFORMATERROR = "InvalidParameter.VpcCidrFormatError"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 有SQL任务尚未执行完成。
	RESOURCEINUSE_UNFINISHEDSQLS = "ResourceInUse.UnfinishedSQLs"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 指定集群配置实例不存在
	RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTFOUND = "ResourceNotFound.DataEngineConfigInstanceNotFound"

	// 指定集群配置实例已存在
	RESOURCENOTFOUND_DATAENGINECONFIGINSTANCENOTUNIQUE = "ResourceNotFound.DataEngineConfigInstanceNotUnique"

	// 指定集群未处于运行状态
	RESOURCENOTFOUND_DATAENGINENOTACTIVITY = "ResourceNotFound.DataEngineNotActivity"

	// 指定的引擎不存在
	RESOURCENOTFOUND_DATAENGINENOTFOUND = "ResourceNotFound.DataEngineNotFound"

	// 指定集群为非多版本类型，不支持该项操作
	RESOURCENOTFOUND_DATAENGINENOTMULTIVERSION = "ResourceNotFound.DataEngineNotMultiVersion"

	// 指定的引擎未处于运行中
	RESOURCENOTFOUND_DATAENGINENOTRUNNING = "ResourceNotFound.DataEngineNotRunning"

	// 指定的引擎已存在
	RESOURCENOTFOUND_DATAENGINENOTUNIQUE = "ResourceNotFound.DataEngineNotUnique"

	// 指定集群为非Spark批作业类型，不支持该项操作
	RESOURCENOTFOUND_DATAENGINETYPEONLYSUPPORTBATCH = "ResourceNotFound.DataEngineTypeOnlySupportBatch"

	// 数据源连接不存在，请重试，或者提交工单联系我们
	RESOURCENOTFOUND_DATASOURCENOTFOUND = "ResourceNotFound.DatasourceNotFound"

	// 找不到默认引擎
	RESOURCENOTFOUND_DEFAULTDATAENGINENOTFOUND = "ResourceNotFound.DefaultDataEngineNotFound"

	// 指定集群镜像Session配置不存在
	RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTFOUND = "ResourceNotFound.ImageSessionConfigNotFound"

	// 指定集群镜像Session配置已存在
	RESOURCENOTFOUND_IMAGESESSIONCONFIGNOTUNIQUE = "ResourceNotFound.ImageSessionConfigNotUnique"

	// 指定集群镜像未激活
	RESOURCENOTFOUND_IMAGEVERSIONNOTACTIVITY = "ResourceNotFound.ImageVersionNotActivity"

	// 指定集群镜像版本不存在
	RESOURCENOTFOUND_IMAGEVERSIONNOTFOUND = "ResourceNotFound.ImageVersionNotFound"

	// 指定集群镜像版本已存在
	RESOURCENOTFOUND_IMAGEVERSIONNOTUNIQUE = "ResourceNotFound.ImageVersionNotUnique"

	// 未找到结果路径。
	RESOURCENOTFOUND_RESULTOUTPUTPATHNOTFOUND = "ResourceNotFound.ResultOutputPathNotFound"

	// 当前无资源创建session，请稍后重试或使用包年包月集群。
	RESOURCENOTFOUND_SESSIONINSUFFICIENTRESOURCES = "ResourceNotFound.SessionInsufficientResources"

	// session不存在。
	RESOURCENOTFOUND_SESSIONNOTFOUND = "ResourceNotFound.SessionNotFound"

	// session已消亡。
	RESOURCENOTFOUND_SESSIONSTATEDEAD = "ResourceNotFound.SessionStateDead"

	// 任务已经失败
	RESOURCENOTFOUND_TASKALREADYFAILED = "ResourceNotFound.TaskAlreadyFailed"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 账号余额不足，无法执行SQL任务。
	RESOURCEUNAVAILABLE_BALANCEINSUFFICIENT = "ResourceUnavailable.BalanceInsufficient"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 无权限授权策略。
	RESOURCESSOLDOUT_UNAUTHORIZEDGRANTPOLICY = "ResourcesSoldOut.UnauthorizedGrantPolicy"

	// 无权限操作。
	RESOURCESSOLDOUT_UNAUTHORIZEDOPERATION = "ResourcesSoldOut.UnauthorizedOperation"

	// 无权限回收权限。
	RESOURCESSOLDOUT_UNAUTHORIZEDREVOKEPOLICY = "ResourcesSoldOut.UnauthorizedRevokePolicy"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 子用户不是管理员，无权将用户添加到工作组。
	UNAUTHORIZEDOPERATION_ADDUSERSTOWORKGROUP = "UnauthorizedOperation.AddUsersToWorkgroup"

	// 子用户不是管理员，无权绑定工作组到用户。
	UNAUTHORIZEDOPERATION_BINDWORKGROUPSTOUSER = "UnauthorizedOperation.BindWorkgroupsToUser"

	// 子用户不是管理员，无权创建工作组。
	UNAUTHORIZEDOPERATION_CREATEWORKGROUP = "UnauthorizedOperation.CreateWorkgroup"

	// 子用户无权删除计算引擎。
	UNAUTHORIZEDOPERATION_DELETECOMPUTINGENGINE = "UnauthorizedOperation.DeleteComputingEngine"

	// 子用户不是管理员，无权删除用户。
	UNAUTHORIZEDOPERATION_DELETEUSER = "UnauthorizedOperation.DeleteUser"

	// 子用户不是管理员，无权将用户从工作组解绑。
	UNAUTHORIZEDOPERATION_DELETEUSERSFROMWORKGROUP = "UnauthorizedOperation.DeleteUsersFromWorkgroup"

	// 子用户不是管理员，无权删除工作组。
	UNAUTHORIZEDOPERATION_DELETEWORKGROUP = "UnauthorizedOperation.DeleteWorkgroup"

	// 子用户无权授予特定权限。
	UNAUTHORIZEDOPERATION_GRANTPOLICY = "UnauthorizedOperation.GrantPolicy"

	// 子用户无权修改引擎配置。
	UNAUTHORIZEDOPERATION_MODIFYCOMPUTINGENGINE = "UnauthorizedOperation.ModifyComputingEngine"

	// 子用户不是管理员，无权修改用户信息。
	UNAUTHORIZEDOPERATION_MODIFYUSERINFO = "UnauthorizedOperation.ModifyUserInfo"

	// 子用户不是管理员，无权修改用户类型。
	UNAUTHORIZEDOPERATION_MODIFYUSERTYPE = "UnauthorizedOperation.ModifyUserType"

	// 子用户不是管理员，无权修改工作组信息。
	UNAUTHORIZEDOPERATION_MODIFYWORKGROUPINFO = "UnauthorizedOperation.ModifyWorkgroupInfo"

	// 没有支付权限。
	UNAUTHORIZEDOPERATION_NOPAYMENTAUTHORITY = "UnauthorizedOperation.NoPaymentAuthority"

	// 子用户无权操作引擎。
	UNAUTHORIZEDOPERATION_OPERATECOMPUTINGENGINE = "UnauthorizedOperation.OperateComputingEngine"

	// 子用户无权取消特定权限。
	UNAUTHORIZEDOPERATION_REVOKEPOLICY = "UnauthorizedOperation.RevokePolicy"

	// 子用户不是管理员，无权将工作组和用户解绑。
	UNAUTHORIZEDOPERATION_UNBINDWORKGROUPSFROMUSER = "UnauthorizedOperation.UnbindWorkgroupsFromUser"

	// 子用户无权使用计算引擎。
	UNAUTHORIZEDOPERATION_USECOMPUTINGENGINE = "UnauthorizedOperation.UseComputingEngine"

	// 子用户不存在。
	UNAUTHORIZEDOPERATION_USERNOTEXIST = "UnauthorizedOperation.UserNotExist"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 无法修改主账号。
	UNSUPPORTEDOPERATION_MODIFYOWNERUNSUPPORTED = "UnsupportedOperation.ModifyOwnerUnsupported"
)
