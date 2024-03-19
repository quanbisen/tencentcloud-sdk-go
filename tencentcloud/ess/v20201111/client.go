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

package v20201111

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-11-11"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewBindEmployeeUserIdWithClientOpenIdRequest() (request *BindEmployeeUserIdWithClientOpenIdRequest) {
    request = &BindEmployeeUserIdWithClientOpenIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "BindEmployeeUserIdWithClientOpenId")
    
    
    return
}

func NewBindEmployeeUserIdWithClientOpenIdResponse() (response *BindEmployeeUserIdWithClientOpenIdResponse) {
    response = &BindEmployeeUserIdWithClientOpenIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// BindEmployeeUserIdWithClientOpenId
// 此接口（BindEmployeeUserIdWithClientOpenId）用于将电子签系统员工UserId与客户系统员工OpenId进行绑定。
//
// 
//
// 此OpenId只在 [更新企业员工信息 ](https://qian.tencent.com/developers/companyApis/staffs/UpdateIntegrationEmployees)、[移除企业员工](https://qian.tencent.com/developers/companyApis/staffs/DeleteIntegrationEmployees) 等场景下可以使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) BindEmployeeUserIdWithClientOpenId(request *BindEmployeeUserIdWithClientOpenIdRequest) (response *BindEmployeeUserIdWithClientOpenIdResponse, err error) {
    return c.BindEmployeeUserIdWithClientOpenIdWithContext(context.Background(), request)
}

// BindEmployeeUserIdWithClientOpenId
// 此接口（BindEmployeeUserIdWithClientOpenId）用于将电子签系统员工UserId与客户系统员工OpenId进行绑定。
//
// 
//
// 此OpenId只在 [更新企业员工信息 ](https://qian.tencent.com/developers/companyApis/staffs/UpdateIntegrationEmployees)、[移除企业员工](https://qian.tencent.com/developers/companyApis/staffs/DeleteIntegrationEmployees) 等场景下可以使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) BindEmployeeUserIdWithClientOpenIdWithContext(ctx context.Context, request *BindEmployeeUserIdWithClientOpenIdRequest) (response *BindEmployeeUserIdWithClientOpenIdResponse, err error) {
    if request == nil {
        request = NewBindEmployeeUserIdWithClientOpenIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BindEmployeeUserIdWithClientOpenId require credential")
    }

    request.SetContext(ctx)
    
    response = NewBindEmployeeUserIdWithClientOpenIdResponse()
    err = c.Send(request, response)
    return
}

func NewCancelFlowRequest() (request *CancelFlowRequest) {
    request = &CancelFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CancelFlow")
    
    
    return
}

func NewCancelFlowResponse() (response *CancelFlowResponse) {
    response = &CancelFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelFlow
// 用于撤销合同流程<br/>
//
// 适用场景：如果某个合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。<br/>
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 注:
//
// 1. 如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同，签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/companyApis/operateFlows/CreateReleaseFlow" target="_blank">发起解除合同流程</a>接口。
//
// 
//
// 2. 有对应合同撤销权限的人:  <font color='red'>合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人</font>
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1f9f07fea6a70766cd286e0d58682ee2.png)
//
// 
//
// 3. <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// 
//
// 4.  撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER_CANCELREASON = "MissingParameter.CancelReason"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelFlow(request *CancelFlowRequest) (response *CancelFlowResponse, err error) {
    return c.CancelFlowWithContext(context.Background(), request)
}

// CancelFlow
// 用于撤销合同流程<br/>
//
// 适用场景：如果某个合同流程当前至少还有一方没有签署，则可通过该接口取消该合同流程。常用于合同发错、内容填错，需要及时撤销的场景。<br/>
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 注:
//
// 1. 如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同，签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/companyApis/operateFlows/CreateReleaseFlow" target="_blank">发起解除合同流程</a>接口。
//
// 
//
// 2. 有对应合同撤销权限的人:  <font color='red'>合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人</font>
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1f9f07fea6a70766cd286e0d58682ee2.png)
//
// 
//
// 3. <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// 
//
// 4.  撤销后可以看合同PDF内容的人员： 发起方的超管， 发起方自己，发起方撤销合同的操作人员，已经签署合同、已经填写合同、邀请填写已经补充信息的参与人员
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CANCELREASON = "InvalidParameter.CancelReason"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_STATUS = "InvalidParameter.Status"
//  MISSINGPARAMETER_CANCELREASON = "MissingParameter.CancelReason"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelFlowWithContext(ctx context.Context, request *CancelFlowRequest) (response *CancelFlowResponse, err error) {
    if request == nil {
        request = NewCancelFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCancelMultiFlowSignQRCodeRequest() (request *CancelMultiFlowSignQRCodeRequest) {
    request = &CancelMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CancelMultiFlowSignQRCode")
    
    
    return
}

func NewCancelMultiFlowSignQRCodeResponse() (response *CancelMultiFlowSignQRCodeResponse) {
    response = &CancelMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelMultiFlowSignQRCode
// 此接口（CancelMultiFlowSignQRCode）用于废除一码多签签署码。
//
// 该接口所需的二维码ID，源自[创建一码多签签署码](https://qian.tencent.com/developers/companyApis/startFlows/CreateMultiFlowSignQRCode)生成的。
//
// 如果该签署码尚处于有效期内，可通过本接口将其设置为失效状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_QRCODEID = "MissingParameter.QrCodeId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QRHASEXPIRE = "OperationDenied.QrHasExpire"
//  OPERATIONDENIED_QRINVALID = "OperationDenied.QrInvalid"
//  RESOURCENOTFOUND_QRINFO = "ResourceNotFound.QrInfo"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelMultiFlowSignQRCode(request *CancelMultiFlowSignQRCodeRequest) (response *CancelMultiFlowSignQRCodeResponse, err error) {
    return c.CancelMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// CancelMultiFlowSignQRCode
// 此接口（CancelMultiFlowSignQRCode）用于废除一码多签签署码。
//
// 该接口所需的二维码ID，源自[创建一码多签签署码](https://qian.tencent.com/developers/companyApis/startFlows/CreateMultiFlowSignQRCode)生成的。
//
// 如果该签署码尚处于有效期内，可通过本接口将其设置为失效状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_QRCODEID = "MissingParameter.QrCodeId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_QRHASEXPIRE = "OperationDenied.QrHasExpire"
//  OPERATIONDENIED_QRINVALID = "OperationDenied.QrInvalid"
//  RESOURCENOTFOUND_QRINFO = "ResourceNotFound.QrInfo"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CancelMultiFlowSignQRCodeWithContext(ctx context.Context, request *CancelMultiFlowSignQRCodeRequest) (response *CancelMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewCancelMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCancelUserAutoSignEnableUrlRequest() (request *CancelUserAutoSignEnableUrlRequest) {
    request = &CancelUserAutoSignEnableUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CancelUserAutoSignEnableUrl")
    
    
    return
}

func NewCancelUserAutoSignEnableUrlResponse() (response *CancelUserAutoSignEnableUrlResponse) {
    response = &CancelUserAutoSignEnableUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelUserAutoSignEnableUrl
// 用来撤销<a href="https://qian.tencent.com/developers/companyApis/users/CreateUserAutoSignEnableUrl" target="_blank">获取个人用户自动签的开通状态</a>生成的开通链接，撤销生成的链接失效。
//
// 
//
// 注:
//
// <ul><li>若个人用户已经用生成的完成自动签署的开通，撤销链接无效不会对开通结果产生影响(此情况接口会报错)。</li>
//
// <li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERAUTOSIGNENABLEALREADY = "FailedOperation.UserAutoSignEnableAlready"
//  FAILEDOPERATION_USERAUTOSIGNENABLEURLNOTEXIST = "FailedOperation.UserAutoSignEnableUrlNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CancelUserAutoSignEnableUrl(request *CancelUserAutoSignEnableUrlRequest) (response *CancelUserAutoSignEnableUrlResponse, err error) {
    return c.CancelUserAutoSignEnableUrlWithContext(context.Background(), request)
}

// CancelUserAutoSignEnableUrl
// 用来撤销<a href="https://qian.tencent.com/developers/companyApis/users/CreateUserAutoSignEnableUrl" target="_blank">获取个人用户自动签的开通状态</a>生成的开通链接，撤销生成的链接失效。
//
// 
//
// 注:
//
// <ul><li>若个人用户已经用生成的完成自动签署的开通，撤销链接无效不会对开通结果产生影响(此情况接口会报错)。</li>
//
// <li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERAUTOSIGNENABLEALREADY = "FailedOperation.UserAutoSignEnableAlready"
//  FAILEDOPERATION_USERAUTOSIGNENABLEURLNOTEXIST = "FailedOperation.UserAutoSignEnableUrlNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CancelUserAutoSignEnableUrlWithContext(ctx context.Context, request *CancelUserAutoSignEnableUrlRequest) (response *CancelUserAutoSignEnableUrlResponse, err error) {
    if request == nil {
        request = NewCancelUserAutoSignEnableUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelUserAutoSignEnableUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelUserAutoSignEnableUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchCancelFlowUrlRequest() (request *CreateBatchCancelFlowUrlRequest) {
    request = &CreateBatchCancelFlowUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateBatchCancelFlowUrl")
    
    
    return
}

func NewCreateBatchCancelFlowUrlResponse() (response *CreateBatchCancelFlowUrlResponse) {
    response = &CreateBatchCancelFlowUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBatchCancelFlowUrl
// 指定需要批量撤回的签署流程Id，以获取批量撤销链接。
//
// 客户需指定要撤回的签署流程Id，最多可指定100个，超过100则不处理。
//
// 接口调用成功后，将返回批量撤回合同的链接。通过点击链接，可跳转至电子签小程序完成批量撤回操作。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 批量撤销结果可以通过接口返回的TaskId关联[批量撤销任务结果回调](https://qian.tencent.com/developers/company/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)或通过接口[查询批量撤销签署流程任务结果](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl)
//
// 
//
// 
//
// 注：
//
// 1. 如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同，签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/companyApis/operateFlows/CreateReleaseFlow" target="_blank">发起解除合同流程</a>接口。
//
// 
//
// 2. 有对应合同撤销权限的人:  <font color='red'>合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人</font>
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1f9f07fea6a70766cd286e0d58682ee2.png)
//
// 
//
// 3. <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BATCHCANCELFORBID = "OperationDenied.BatchCancelForbid"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchCancelFlowUrl(request *CreateBatchCancelFlowUrlRequest) (response *CreateBatchCancelFlowUrlResponse, err error) {
    return c.CreateBatchCancelFlowUrlWithContext(context.Background(), request)
}

// CreateBatchCancelFlowUrl
// 指定需要批量撤回的签署流程Id，以获取批量撤销链接。
//
// 客户需指定要撤回的签署流程Id，最多可指定100个，超过100则不处理。
//
// 接口调用成功后，将返回批量撤回合同的链接。通过点击链接，可跳转至电子签小程序完成批量撤回操作。
//
// 
//
// - **可撤回合同状态**：未全部签署完成
//
// - **不撤回合同状态**：已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态。
//
// 
//
// 批量撤销结果可以通过接口返回的TaskId关联[批量撤销任务结果回调](https://qian.tencent.com/developers/company/callback_types_contracts_sign#%E4%B9%9D-%E6%89%B9%E9%87%8F%E6%92%A4%E9%94%80%E7%BB%93%E6%9E%9C%E5%9B%9E%E8%B0%83)或通过接口[查询批量撤销签署流程任务结果](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl)
//
// 
//
// 
//
// 注：
//
// 1. 如果合同流程中的参与方均已签署完毕，则无法通过该接口撤销合同，签署完毕的合同需要双方走解除流程将合同作废，可以参考<a href="https://qian.tencent.com/developers/companyApis/operateFlows/CreateReleaseFlow" target="_blank">发起解除合同流程</a>接口。
//
// 
//
// 2. 有对应合同撤销权限的人:  <font color='red'>合同的发起人（并已经授予撤销权限）或者发起人所在企业的超管、法人</font>
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/1f9f07fea6a70766cd286e0d58682ee2.png)
//
// 
//
// 3. <font color='red'>只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。</font>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BATCHCANCELFORBID = "OperationDenied.BatchCancelForbid"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchCancelFlowUrlWithContext(ctx context.Context, request *CreateBatchCancelFlowUrlRequest) (response *CreateBatchCancelFlowUrlResponse, err error) {
    if request == nil {
        request = NewCreateBatchCancelFlowUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchCancelFlowUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchCancelFlowUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchQuickSignUrlRequest() (request *CreateBatchQuickSignUrlRequest) {
    request = &CreateBatchQuickSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateBatchQuickSignUrl")
    
    
    return
}

func NewCreateBatchQuickSignUrlResponse() (response *CreateBatchQuickSignUrlResponse) {
    response = &CreateBatchQuickSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBatchQuickSignUrl
// 该接口用于发起合同后，生成个人用户的批量签署链接, 暂时不支持企业端签署。
//
// **注意：**
//
// 1. 该接口目前仅支持签署人类型是**个人签署方的批量签署场景**(ApproverType=1)。
//
// 2. 该接口可生成批量签署链接的C端签署人**必须仅有手写签名(控件类型为SIGN_SIGNATURE)和时间类型的签署控件**，**不支持填写控件** 。
//
// 3. 请确保C端签署人在批量签署合同中**为待签署状态**，如需顺序签署请待前一位参与人签署完成后，再创建该C端用户的签署链接。
//
// 4. 该签署链接**有效期为30分钟**，过期后将失效，如需签署可重新创建批量签署链接 。
//
// 5. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。
//
// 跳转到小程序的实现，参考微信官方文档(分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式)，如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 6. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchQuickSignUrl(request *CreateBatchQuickSignUrlRequest) (response *CreateBatchQuickSignUrlResponse, err error) {
    return c.CreateBatchQuickSignUrlWithContext(context.Background(), request)
}

// CreateBatchQuickSignUrl
// 该接口用于发起合同后，生成个人用户的批量签署链接, 暂时不支持企业端签署。
//
// **注意：**
//
// 1. 该接口目前仅支持签署人类型是**个人签署方的批量签署场景**(ApproverType=1)。
//
// 2. 该接口可生成批量签署链接的C端签署人**必须仅有手写签名(控件类型为SIGN_SIGNATURE)和时间类型的签署控件**，**不支持填写控件** 。
//
// 3. 请确保C端签署人在批量签署合同中**为待签署状态**，如需顺序签署请待前一位参与人签署完成后，再创建该C端用户的签署链接。
//
// 4. 该签署链接**有效期为30分钟**，过期后将失效，如需签署可重新创建批量签署链接 。
//
// 5. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。
//
// 跳转到小程序的实现，参考微信官方文档(分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式)，如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 6. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchQuickSignUrlWithContext(ctx context.Context, request *CreateBatchQuickSignUrlRequest) (response *CreateBatchQuickSignUrlResponse, err error) {
    if request == nil {
        request = NewCreateBatchQuickSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchQuickSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchQuickSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchSignUrlRequest() (request *CreateBatchSignUrlRequest) {
    request = &CreateBatchSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateBatchSignUrl")
    
    
    return
}

func NewCreateBatchSignUrlResponse() (response *CreateBatchSignUrlResponse) {
    response = &CreateBatchSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBatchSignUrl
// 通过此接口，创建小程序批量签署链接，个人/企业员工点击此链接即可跳转小程序进行批量签署。
//
// 请确保生成链接时候的身份信息和签署合同参与方的信息保持一致。
//
// 
//
// 注：
//
// - 使用此接口生成链接，需要提前开通 `使用手机号验证签署方身份` 功能，在 `腾讯电子签网页端-企业设置-拓展服务` 中可以找到。
//
// - 参与人点击链接后需短信验证码才能查看合同内容。
//
// - 企业用户批量签署，需要传OrganizationName（参与方所在企业名称）参数生成签署链接，`请确保此企业已完成腾讯电子签企业认证`。
//
// - 个人批量签署，签名区`仅支持手写签名`。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchSignUrl(request *CreateBatchSignUrlRequest) (response *CreateBatchSignUrlResponse, err error) {
    return c.CreateBatchSignUrlWithContext(context.Background(), request)
}

// CreateBatchSignUrl
// 通过此接口，创建小程序批量签署链接，个人/企业员工点击此链接即可跳转小程序进行批量签署。
//
// 请确保生成链接时候的身份信息和签署合同参与方的信息保持一致。
//
// 
//
// 注：
//
// - 使用此接口生成链接，需要提前开通 `使用手机号验证签署方身份` 功能，在 `腾讯电子签网页端-企业设置-拓展服务` 中可以找到。
//
// - 参与人点击链接后需短信验证码才能查看合同内容。
//
// - 企业用户批量签署，需要传OrganizationName（参与方所在企业名称）参数生成签署链接，`请确保此企业已完成腾讯电子签企业认证`。
//
// - 个人批量签署，签名区`仅支持手写签名`。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateBatchSignUrlWithContext(ctx context.Context, request *CreateBatchSignUrlRequest) (response *CreateBatchSignUrlResponse, err error) {
    if request == nil {
        request = NewCreateBatchSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConvertTaskApiRequest() (request *CreateConvertTaskApiRequest) {
    request = &CreateConvertTaskApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateConvertTaskApi")
    
    
    return
}

func NewCreateConvertTaskApiResponse() (response *CreateConvertTaskApiResponse) {
    response = &CreateConvertTaskApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConvertTaskApi
// 此接口（CreateConvertTaskApi）用来将word、excel、html、图片、txt类型文件转换为PDF文件。<br />
//
// 前提条件：源文件已经通过 <a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">文件上传接口</a>完成上传，并得到了源文件的资源Id。<br />
//
// 适用场景1：已经上传了一个word文件，希望将该word文件转换成pdf文件后发起合同
//
// 适用场景2：已经上传了一个jpg图片文件，希望将该图片文件转换成pdf文件后发起合同<br />
//
// 转换文件是一个耗时操作，若想查看转换任务是否完成，可以通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/GetTaskResultApi" target="_blank">查询转换任务状态</a>接口获取任务状态。<br />
//
// 注: 
//
// 1. `支持的文件类型有doc、docx、xls、xlsx、html、jpg、jpeg、png、bmp、txt`
//
// 2. `可通过发起合同时设置预览来检查转换文件是否达到预期效果`
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateConvertTaskApi(request *CreateConvertTaskApiRequest) (response *CreateConvertTaskApiResponse, err error) {
    return c.CreateConvertTaskApiWithContext(context.Background(), request)
}

// CreateConvertTaskApi
// 此接口（CreateConvertTaskApi）用来将word、excel、html、图片、txt类型文件转换为PDF文件。<br />
//
// 前提条件：源文件已经通过 <a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles" target="_blank">文件上传接口</a>完成上传，并得到了源文件的资源Id。<br />
//
// 适用场景1：已经上传了一个word文件，希望将该word文件转换成pdf文件后发起合同
//
// 适用场景2：已经上传了一个jpg图片文件，希望将该图片文件转换成pdf文件后发起合同<br />
//
// 转换文件是一个耗时操作，若想查看转换任务是否完成，可以通过<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/GetTaskResultApi" target="_blank">查询转换任务状态</a>接口获取任务状态。<br />
//
// 注: 
//
// 1. `支持的文件类型有doc、docx、xls、xlsx、html、jpg、jpeg、png、bmp、txt`
//
// 2. `可通过发起合同时设置预览来检查转换文件是否达到预期效果`
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateConvertTaskApiWithContext(ctx context.Context, request *CreateConvertTaskApiRequest) (response *CreateConvertTaskApiResponse, err error) {
    if request == nil {
        request = NewCreateConvertTaskApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConvertTaskApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConvertTaskApiResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDocumentRequest() (request *CreateDocumentRequest) {
    request = &CreateDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateDocument")
    
    
    return
}

func NewCreateDocumentResponse() (response *CreateDocumentResponse) {
    response = &CreateDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDocument
// 创建签署流程电子文档<br />
//
// 
//
// ###  调用流程
//
// 该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。需要配置<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow" target="_blank">创建签署流程</a>和<a href="https://qian.tencent.com/developers/companyApis/startFlows/StartFlow" target="_blank">发起签署流程</a>接口使用。具体逻辑可以参考下图:
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/06f2bc0f1772d8deac2f92b5df61a5ac.png)
//
// 
//
// 
//
// ### 填充模板中定义的填写控件
//
// 模板中配置的<font color="red">发起人填充控件</font>可以通过本接口的**FormFields数组**字段填充
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/37457e0e450fc221effddfcb8b1bad55.png)
//
// 
//
// 填充的传参示例如下
//
// 
//
// ```
//
//     request.FormFields = [{
//
//             "ComponentName": "项目的名字",
//
//             "ComponentValue": "休闲山庄"
//
//         }, {
//
//             "ComponentName": "项目的地址",
//
//             "ComponentValue": "凤凰山北侧",
//
//         }, {
//
//             "ComponentName": "范围",
//
//             "ComponentValue": "凤凰山至107国道",
//
//         }, {
//
//             "ComponentName": "面积",
//
//             "ComponentValue": "100亩",
//
//         }, {
//
//             "ComponentName": "基本情况",
//
//             "ComponentValue": "完好",
//
//         }, , {
//
//             "ComponentName": "用途",
//
//             "ComponentValue": "经营农家乐",
//
//         }
//
//     ]
//
// ```
//
// 合成后合同样子示例
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/140a2fb771ac66a185d0a000d37485f6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateDocument(request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    return c.CreateDocumentWithContext(context.Background(), request)
}

// CreateDocument
// 创建签署流程电子文档<br />
//
// 
//
// ###  调用流程
//
// 该接口需要给对应的流程指定一个模板id，并且填充该模板中需要补充的信息。需要配置<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow" target="_blank">创建签署流程</a>和<a href="https://qian.tencent.com/developers/companyApis/startFlows/StartFlow" target="_blank">发起签署流程</a>接口使用。具体逻辑可以参考下图:
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/06f2bc0f1772d8deac2f92b5df61a5ac.png)
//
// 
//
// 
//
// ### 填充模板中定义的填写控件
//
// 模板中配置的<font color="red">发起人填充控件</font>可以通过本接口的**FormFields数组**字段填充
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/37457e0e450fc221effddfcb8b1bad55.png)
//
// 
//
// 填充的传参示例如下
//
// 
//
// ```
//
//     request.FormFields = [{
//
//             "ComponentName": "项目的名字",
//
//             "ComponentValue": "休闲山庄"
//
//         }, {
//
//             "ComponentName": "项目的地址",
//
//             "ComponentValue": "凤凰山北侧",
//
//         }, {
//
//             "ComponentName": "范围",
//
//             "ComponentValue": "凤凰山至107国道",
//
//         }, {
//
//             "ComponentName": "面积",
//
//             "ComponentValue": "100亩",
//
//         }, {
//
//             "ComponentName": "基本情况",
//
//             "ComponentValue": "完好",
//
//         }, , {
//
//             "ComponentName": "用途",
//
//             "ComponentValue": "经营农家乐",
//
//         }
//
//     ]
//
// ```
//
// 合成后合同样子示例
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/140a2fb771ac66a185d0a000d37485f6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_TEMPLATEHASNORESOURCE = "FailedOperation.TemplateHasNoResource"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateDocumentWithContext(ctx context.Context, request *CreateDocumentRequest) (response *CreateDocumentResponse, err error) {
    if request == nil {
        request = NewCreateDocumentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDocument require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateEmbedWebUrlRequest() (request *CreateEmbedWebUrlRequest) {
    request = &CreateEmbedWebUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateEmbedWebUrl")
    
    
    return
}

func NewCreateEmbedWebUrlResponse() (response *CreateEmbedWebUrlResponse) {
    response = &CreateEmbedWebUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateEmbedWebUrl
// 本接口（CreateEmbedWebUrl）用于创建嵌入Web的链接，支持以下类型的Web链接创建：
//
// 1. 创建印章
//
// 2. 创建模板
//
// 3. 修改模板
//
// 4. 预览模板
//
// 5. 预览合同流程
//
// 
//
// 用户可以通过这些链接快速将其集成到自己的系统中。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateEmbedWebUrl(request *CreateEmbedWebUrlRequest) (response *CreateEmbedWebUrlResponse, err error) {
    return c.CreateEmbedWebUrlWithContext(context.Background(), request)
}

// CreateEmbedWebUrl
// 本接口（CreateEmbedWebUrl）用于创建嵌入Web的链接，支持以下类型的Web链接创建：
//
// 1. 创建印章
//
// 2. 创建模板
//
// 3. 修改模板
//
// 4. 预览模板
//
// 5. 预览合同流程
//
// 
//
// 用户可以通过这些链接快速将其集成到自己的系统中。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateEmbedWebUrlWithContext(ctx context.Context, request *CreateEmbedWebUrlRequest) (response *CreateEmbedWebUrlResponse, err error) {
    if request == nil {
        request = NewCreateEmbedWebUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateEmbedWebUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateEmbedWebUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateExtendedServiceAuthInfosRequest() (request *CreateExtendedServiceAuthInfosRequest) {
    request = &CreateExtendedServiceAuthInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateExtendedServiceAuthInfos")
    
    
    return
}

func NewCreateExtendedServiceAuthInfosResponse() (response *CreateExtendedServiceAuthInfosResponse) {
    response = &CreateExtendedServiceAuthInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateExtendedServiceAuthInfos
// 创建企业扩展服务授权，当前仅支持授权 “企业自动签” 和 “批量签署” 给企业员工。
//
// 该接口作用和电子签控制台 企业设置-扩展服务-企业自动签署和批量签署授权 两个模块功能相同，可通过该接口授权给企业员工。
//
// 
//
// 注：“企业自动签授权”支持集团代子企业操作，请联系运营开通此功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateExtendedServiceAuthInfos(request *CreateExtendedServiceAuthInfosRequest) (response *CreateExtendedServiceAuthInfosResponse, err error) {
    return c.CreateExtendedServiceAuthInfosWithContext(context.Background(), request)
}

// CreateExtendedServiceAuthInfos
// 创建企业扩展服务授权，当前仅支持授权 “企业自动签” 和 “批量签署” 给企业员工。
//
// 该接口作用和电子签控制台 企业设置-扩展服务-企业自动签署和批量签署授权 两个模块功能相同，可通过该接口授权给企业员工。
//
// 
//
// 注：“企业自动签授权”支持集团代子企业操作，请联系运营开通此功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateExtendedServiceAuthInfosWithContext(ctx context.Context, request *CreateExtendedServiceAuthInfosRequest) (response *CreateExtendedServiceAuthInfosResponse, err error) {
    if request == nil {
        request = NewCreateExtendedServiceAuthInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateExtendedServiceAuthInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateExtendedServiceAuthInfosResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowRequest() (request *CreateFlowRequest) {
    request = &CreateFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlow")
    
    
    return
}

func NewCreateFlowResponse() (response *CreateFlowResponse) {
    response = &CreateFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlow
// 通过模板创建签署流程<br/>
//
// 适用场景：在标准制式的合同场景中，可通过提前预制好模板文件，每次调用模板文件的id，补充合同内容信息及签署信息生成电子合同。
//
// <table>
//
// 	<thead>
//
// 		<tr>
//
// 			<th>签署人类别</th>
//
// 			<th>需要提前准备的信息</th>
//
// 		</tr>
//
// 	</thead>
//
// 	<tbody>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（未认证加入或已认证加入）</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（已认证加入）</td>
//
// 			<td>签署企业的名字、员工在电子签平台的ID（UserId）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>其他企业的员工签署</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>个人（自然人）签署</td>
//
// 			<td>个人的真实名字、个人的触达手机号、个人的身份证（证件号非必传）</td>
//
// 		</tr>
//
// 	</tbody>
//
// </table>
//
// 
//
// 
//
// 注：配合<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument" target="_blank">创建电子文档</a>和<a href="https://qian.tencent.com/developers/companyApis/startFlows/StartFlow" target="_blank">发起签署流程</a>接口使用。整体的逻辑如下图
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/06f2bc0f1772d8deac2f92b5df61a5ac.png)
//
// 
//
// 注：**静默（自动）签署不支持合同签署方存在填写**功能
//
// <br>
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/createflow_seversign.mp4" target="_blank">创建静默（自动）签署模板和开通自动签署</a><br>
//
// 2. <a href="https://dyn.ess.tencent.cn/guide/apivideo/flow_document_start.mp4" target="_blank">用模板创建发起合同</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_NOTFOUNDSHADOWUSER = "FailedOperation.NotFoundShadowUser"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlow(request *CreateFlowRequest) (response *CreateFlowResponse, err error) {
    return c.CreateFlowWithContext(context.Background(), request)
}

// CreateFlow
// 通过模板创建签署流程<br/>
//
// 适用场景：在标准制式的合同场景中，可通过提前预制好模板文件，每次调用模板文件的id，补充合同内容信息及签署信息生成电子合同。
//
// <table>
//
// 	<thead>
//
// 		<tr>
//
// 			<th>签署人类别</th>
//
// 			<th>需要提前准备的信息</th>
//
// 		</tr>
//
// 	</thead>
//
// 	<tbody>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（未认证加入或已认证加入）</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（已认证加入）</td>
//
// 			<td>签署企业的名字、员工在电子签平台的ID（UserId）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>其他企业的员工签署</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>个人（自然人）签署</td>
//
// 			<td>个人的真实名字、个人的触达手机号、个人的身份证（证件号非必传）</td>
//
// 		</tr>
//
// 	</tbody>
//
// </table>
//
// 
//
// 
//
// 注：配合<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument" target="_blank">创建电子文档</a>和<a href="https://qian.tencent.com/developers/companyApis/startFlows/StartFlow" target="_blank">发起签署流程</a>接口使用。整体的逻辑如下图
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/06f2bc0f1772d8deac2f92b5df61a5ac.png)
//
// 
//
// 注：**静默（自动）签署不支持合同签署方存在填写**功能
//
// <br>
//
// 
//
// <font color="red">相关视频指引</font> <br>
//
// 1. <a href="https://dyn.ess.tencent.cn/guide/apivideo/createflow_seversign.mp4" target="_blank">创建静默（自动）签署模板和开通自动签署</a><br>
//
// 2. <a href="https://dyn.ess.tencent.cn/guide/apivideo/flow_document_start.mp4" target="_blank">用模板创建发起合同</a><br>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_NOTFOUNDSHADOWUSER = "FailedOperation.NotFoundShadowUser"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowWithContext(ctx context.Context, request *CreateFlowRequest) (response *CreateFlowResponse, err error) {
    if request == nil {
        request = NewCreateFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowApproversRequest() (request *CreateFlowApproversRequest) {
    request = &CreateFlowApproversRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowApprovers")
    
    
    return
}

func NewCreateFlowApproversResponse() (response *CreateFlowApproversResponse) {
    response = &CreateFlowApproversResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowApprovers
// 适用场景： 当通过模板或文件发起合同时，若未指定企业签署人信息，则可调用此接口补充或添加签署人。同一签署人可补充多个员工作为或签署人，最终实际签署人取决于谁先领取合同完成签署。
//
// 
//
// 限制条件：
//
// 
//
// 1.本企业（发起方企业）企业微信签署人仅支持通过企业微信UserId或姓名+手机号进行补充。
//
// 2.本企业（发起方企业）非企业微信签署人仅支持通过姓名+手机号进行补充。
//
// 3.他方企业仅支持通过姓名+手机号进行补充。
//
// 4.个人签署人支持通过姓名+手机号进行补充，补充动态签署人时：若个人用户已完成实名，则可通过姓名+证件号码进行补充。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SENSITIVE = "InvalidParameter.Sensitive"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowApprovers(request *CreateFlowApproversRequest) (response *CreateFlowApproversResponse, err error) {
    return c.CreateFlowApproversWithContext(context.Background(), request)
}

// CreateFlowApprovers
// 适用场景： 当通过模板或文件发起合同时，若未指定企业签署人信息，则可调用此接口补充或添加签署人。同一签署人可补充多个员工作为或签署人，最终实际签署人取决于谁先领取合同完成签署。
//
// 
//
// 限制条件：
//
// 
//
// 1.本企业（发起方企业）企业微信签署人仅支持通过企业微信UserId或姓名+手机号进行补充。
//
// 2.本企业（发起方企业）非企业微信签署人仅支持通过姓名+手机号进行补充。
//
// 3.他方企业仅支持通过姓名+手机号进行补充。
//
// 4.个人签署人支持通过姓名+手机号进行补充，补充动态签署人时：若个人用户已完成实名，则可通过姓名+证件号码进行补充。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SENSITIVE = "InvalidParameter.Sensitive"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowApproversWithContext(ctx context.Context, request *CreateFlowApproversRequest) (response *CreateFlowApproversResponse, err error) {
    if request == nil {
        request = NewCreateFlowApproversRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowApprovers require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowApproversResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowByFilesRequest() (request *CreateFlowByFilesRequest) {
    request = &CreateFlowByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowByFiles")
    
    
    return
}

func NewCreateFlowByFilesResponse() (response *CreateFlowByFilesResponse) {
    response = &CreateFlowByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowByFiles
// 此接口（CreateFlowByFiles）用来通过上传后的pdf资源编号来创建待签署的合同流程。<br/>
//
// 适用场景：适用非制式的合同文件签署。一般开发者自己有完整的签署文件，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。<br/>
//
// 
//
// <table>
//
// 	<thead>
//
// 		<tr>
//
// 			<th>签署人类别</th>
//
// 			<th>需要提前准备的信息</th>
//
// 		</tr>
//
// 	</thead>
//
// 	<tbody>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（未认证加入或已认证加入）</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（已认证加入）</td>
//
// 			<td>签署企业的名字、员工在电子签平台的ID（UserId）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>其他企业的员工签署</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>个人（自然人）签署</td>
//
// 			<td>个人的真实名字、个人的触达手机号、个人的身份证（证件号非必传）</td>
//
// 		</tr>
//
// 	</tbody>
//
// </table>
//
// 
//
// 
//
// 
//
// 该接口需要依赖[上传文件](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口生成pdf资源编号（FileIds）进行使用。（如果非pdf文件需要调用[创建文件转换任务](https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi)接口转换成pdf资源）<br/>
//
// 
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/f097a74b289e3e1acd740936bdfe9843.png)
//
// 
//
// 注：
//
// -  合同**发起后就会扣减合同的额度**, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（**过期，拒签，签署完成，解除完成等状态不会返还额度**）
//
// - **静默（自动）签署不支持合同签署方存在填写**功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOAPPROVERMOBILECHECKPERMISSION = "OperationDenied.NoApproverMobileCheckPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowByFiles(request *CreateFlowByFilesRequest) (response *CreateFlowByFilesResponse, err error) {
    return c.CreateFlowByFilesWithContext(context.Background(), request)
}

// CreateFlowByFiles
// 此接口（CreateFlowByFiles）用来通过上传后的pdf资源编号来创建待签署的合同流程。<br/>
//
// 适用场景：适用非制式的合同文件签署。一般开发者自己有完整的签署文件，可以通过该接口传入完整的PDF文件及流程信息生成待签署的合同流程。<br/>
//
// 
//
// <table>
//
// 	<thead>
//
// 		<tr>
//
// 			<th>签署人类别</th>
//
// 			<th>需要提前准备的信息</th>
//
// 		</tr>
//
// 	</thead>
//
// 	<tbody>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（未认证加入或已认证加入）</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>自己企业的员工签署（已认证加入）</td>
//
// 			<td>签署企业的名字、员工在电子签平台的ID（UserId）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>其他企业的员工签署</td>
//
// 			<td>签署企业的名字、员工的真实名字、员工的触达手机号、员工的证件号（证件号非必传）</td>
//
// 		</tr>
//
// 		<tr>
//
// 			<td>个人（自然人）签署</td>
//
// 			<td>个人的真实名字、个人的触达手机号、个人的身份证（证件号非必传）</td>
//
// 		</tr>
//
// 	</tbody>
//
// </table>
//
// 
//
// 
//
// 
//
// 该接口需要依赖[上传文件](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口生成pdf资源编号（FileIds）进行使用。（如果非pdf文件需要调用[创建文件转换任务](https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi)接口转换成pdf资源）<br/>
//
// 
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/f097a74b289e3e1acd740936bdfe9843.png)
//
// 
//
// 注：
//
// -  合同**发起后就会扣减合同的额度**, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（**过期，拒签，签署完成，解除完成等状态不会返还额度**）
//
// - **静默（自动）签署不支持合同签署方存在填写**功能
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTTYPE = "InvalidParameter.SignComponentType"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOAPPROVERMOBILECHECKPERMISSION = "OperationDenied.NoApproverMobileCheckPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowByFilesWithContext(ctx context.Context, request *CreateFlowByFilesRequest) (response *CreateFlowByFilesResponse, err error) {
    if request == nil {
        request = NewCreateFlowByFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowByFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowEvidenceReportRequest() (request *CreateFlowEvidenceReportRequest) {
    request = &CreateFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowEvidenceReport")
    
    
    return
}

func NewCreateFlowEvidenceReportResponse() (response *CreateFlowEvidenceReportResponse) {
    response = &CreateFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowEvidenceReport
// 提交申请出证报告任务并返回报告ID。
//
// 
//
// 注意：
//
// <ul><li>使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。</li>
//
// <li>操作人必须是`发起方或者签署方企业的(非走授权书认证)法人或者超管`。</li>
//
// <li>合同流程必须所有参与方`已经签署完成`。</li>
//
// <li>出证过程需一定时间，建议在`提交出证任务后的24小时之后`，通过<a href="https://qian.tencent.com/developers/companyApis/certificate/DescribeFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>接口进行查询执行结果和出证报告的下载URL。</li></ul>
//
// 
//
// <svg id="SvgjsSvg1006" width="262" height="229" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.com/svgjs"><defs id="SvgjsDefs1007"><pattern patternUnits="userSpaceOnUse" id="pattern_mark_0" width="300" height="300"><text x="150" y="100" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><pattern patternUnits="userSpaceOnUse" id="pattern_mark_1" width="300" height="300"><text x="150" y="200" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><marker id="SvgjsMarker1021" markerWidth="12" markerHeight="8" refX="9" refY="4" viewBox="0 0 12 8" orient="auto" markerUnits="userSpaceOnUse" stroke-dasharray="0,0"><path id="SvgjsPath1022" d="M0,0 L12,4 L0,8 L0,0" fill="#323232" stroke="#323232" stroke-width="1"></path></marker></defs><rect id="svgbackgroundid" width="262" height="229" fill="transparent"></rect><rect id="SvgjsRect1009" width="262" height="229" fill="url(#pattern_mark_0)"></rect><rect id="SvgjsRect1010" width="262" height="229" fill="url(#pattern_mark_1)"></rect><g id="SvgjsG1011" transform="translate(31.75,25)"><path id="SvgjsPath1012" d="M 0 0L 198 0L 198 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1013"><text id="SvgjsText1014" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="178px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1015" dy="16" x="99"><tspan id="SvgjsTspan1016" style="text-decoration:;fill: rgb(28, 30, 33);">CreateFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1017" dy="16" x="99"><tspan id="SvgjsTspan1018" style="text-decoration:;fill: rgb(51, 51, 51);">提交申请出证报告任务</tspan></tspan></text></g></g><g id="SvgjsG1019"><path id="SvgjsPath1020" d="M130.75 84.5L130.75 114.5L130.75 114.5L130.75 143.2" stroke="#323232" stroke-width="1" fill="none" marker-end="url(#SvgjsMarker1021)"></path></g><g id="SvgjsG1023" transform="translate(25,145)"><path id="SvgjsPath1024" d="M 0 0L 211.5 0L 211.5 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1025"><text id="SvgjsText1026" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="192px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1027" dy="16" x="106"><tspan id="SvgjsTspan1028" style="text-decoration:;fill: rgb(28, 30, 33);">DescribeFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1029" dy="16" x="106"><tspan id="SvgjsTspan1030" style="text-decoration:;fill: rgb(51, 51, 51);">获取出证报告任务执行结果</tspan></tspan></text></g></g></svg>
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_PROVENOQUOTA = "OperationDenied.ProveNoQuota"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowEvidenceReport(request *CreateFlowEvidenceReportRequest) (response *CreateFlowEvidenceReportResponse, err error) {
    return c.CreateFlowEvidenceReportWithContext(context.Background(), request)
}

// CreateFlowEvidenceReport
// 提交申请出证报告任务并返回报告ID。
//
// 
//
// 注意：
//
// <ul><li>使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。</li>
//
// <li>操作人必须是`发起方或者签署方企业的(非走授权书认证)法人或者超管`。</li>
//
// <li>合同流程必须所有参与方`已经签署完成`。</li>
//
// <li>出证过程需一定时间，建议在`提交出证任务后的24小时之后`，通过<a href="https://qian.tencent.com/developers/companyApis/certificate/DescribeFlowEvidenceReport" target="_blank">获取出证报告任务执行结果</a>接口进行查询执行结果和出证报告的下载URL。</li></ul>
//
// 
//
// <svg id="SvgjsSvg1006" width="262" height="229" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.com/svgjs"><defs id="SvgjsDefs1007"><pattern patternUnits="userSpaceOnUse" id="pattern_mark_0" width="300" height="300"><text x="150" y="100" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><pattern patternUnits="userSpaceOnUse" id="pattern_mark_1" width="300" height="300"><text x="150" y="200" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><marker id="SvgjsMarker1021" markerWidth="12" markerHeight="8" refX="9" refY="4" viewBox="0 0 12 8" orient="auto" markerUnits="userSpaceOnUse" stroke-dasharray="0,0"><path id="SvgjsPath1022" d="M0,0 L12,4 L0,8 L0,0" fill="#323232" stroke="#323232" stroke-width="1"></path></marker></defs><rect id="svgbackgroundid" width="262" height="229" fill="transparent"></rect><rect id="SvgjsRect1009" width="262" height="229" fill="url(#pattern_mark_0)"></rect><rect id="SvgjsRect1010" width="262" height="229" fill="url(#pattern_mark_1)"></rect><g id="SvgjsG1011" transform="translate(31.75,25)"><path id="SvgjsPath1012" d="M 0 0L 198 0L 198 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1013"><text id="SvgjsText1014" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="178px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1015" dy="16" x="99"><tspan id="SvgjsTspan1016" style="text-decoration:;fill: rgb(28, 30, 33);">CreateFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1017" dy="16" x="99"><tspan id="SvgjsTspan1018" style="text-decoration:;fill: rgb(51, 51, 51);">提交申请出证报告任务</tspan></tspan></text></g></g><g id="SvgjsG1019"><path id="SvgjsPath1020" d="M130.75 84.5L130.75 114.5L130.75 114.5L130.75 143.2" stroke="#323232" stroke-width="1" fill="none" marker-end="url(#SvgjsMarker1021)"></path></g><g id="SvgjsG1023" transform="translate(25,145)"><path id="SvgjsPath1024" d="M 0 0L 211.5 0L 211.5 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1025"><text id="SvgjsText1026" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="192px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1027" dy="16" x="106"><tspan id="SvgjsTspan1028" style="text-decoration:;fill: rgb(28, 30, 33);">DescribeFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1029" dy="16" x="106"><tspan id="SvgjsTspan1030" style="text-decoration:;fill: rgb(51, 51, 51);">获取出证报告任务执行结果</tspan></tspan></text></g></g></svg>
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_PROVENOQUOTA = "OperationDenied.ProveNoQuota"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowEvidenceReportWithContext(ctx context.Context, request *CreateFlowEvidenceReportRequest) (response *CreateFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewCreateFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowGroupByFilesRequest() (request *CreateFlowGroupByFilesRequest) {
    request = &CreateFlowGroupByFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowGroupByFiles")
    
    
    return
}

func NewCreateFlowGroupByFilesResponse() (response *CreateFlowGroupByFilesResponse) {
    response = &CreateFlowGroupByFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowGroupByFiles
// 此接口（CreateFlowGroupByFiles）可用于通过多个文件创建合同组签署流程。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// 注意事项：使用该接口需要先依赖[多文件上传](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口返回的FileIds。
//
// 
//
// 注：`合同发起后就会扣减合同的额度, 如果未签署完成时撤销合同会返还此额度（过期，拒签，签署完成，解除完成等状态不会返还额度），合同组中每个合同会扣减一个合同额度`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowGroupByFiles(request *CreateFlowGroupByFilesRequest) (response *CreateFlowGroupByFilesResponse, err error) {
    return c.CreateFlowGroupByFilesWithContext(context.Background(), request)
}

// CreateFlowGroupByFiles
// 此接口（CreateFlowGroupByFiles）可用于通过多个文件创建合同组签署流程。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// 注意事项：使用该接口需要先依赖[多文件上传](https://qian.tencent.com/developers/companyApis/templatesAndFiles/UploadFiles)接口返回的FileIds。
//
// 
//
// 注：`合同发起后就会扣减合同的额度, 如果未签署完成时撤销合同会返还此额度（过期，拒签，签署完成，解除完成等状态不会返还额度），合同组中每个合同会扣减一个合同额度`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowGroupByFilesWithContext(ctx context.Context, request *CreateFlowGroupByFilesRequest) (response *CreateFlowGroupByFilesResponse, err error) {
    if request == nil {
        request = NewCreateFlowGroupByFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowGroupByFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowGroupByFilesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowGroupByTemplatesRequest() (request *CreateFlowGroupByTemplatesRequest) {
    request = &CreateFlowGroupByTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowGroupByTemplates")
    
    
    return
}

func NewCreateFlowGroupByTemplatesResponse() (response *CreateFlowGroupByTemplatesResponse) {
    response = &CreateFlowGroupByTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowGroupByTemplates
// 此接口（CreateFlowGroupByTemplates）可用于通过多个模板创建合同组签署流程。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// 注：`合同发起后就会扣减合同的额度, 如果未签署完成时撤销合同会返还此额度（过期，拒签，签署完成，解除完成等状态不会返还额度），合同组中每个合同会扣减一个合同额度`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowGroupByTemplates(request *CreateFlowGroupByTemplatesRequest) (response *CreateFlowGroupByTemplatesResponse, err error) {
    return c.CreateFlowGroupByTemplatesWithContext(context.Background(), request)
}

// CreateFlowGroupByTemplates
// 此接口（CreateFlowGroupByTemplates）可用于通过多个模板创建合同组签署流程。
//
// 
//
// 适用场景：该接口适用于需要一次性完成多份合同签署的情况，多份合同一般具有关联性，用户以目录的形式查看合同。
//
// 
//
// 注：`合同发起后就会扣减合同的额度, 如果未签署完成时撤销合同会返还此额度（过期，拒签，签署完成，解除完成等状态不会返还额度），合同组中每个合同会扣减一个合同额度`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_PREVIEWURLFAIL = "FailedOperation.PreViewUrlFail"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_JSONMARSHAL = "InternalError.JsonMarshal"
//  INTERNALERROR_PDF = "InternalError.Pdf"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_COMPONENTFILEINDEX = "InvalidParameter.ComponentFileIndex"
//  INVALIDPARAMETER_COMPONENTPAGE = "InvalidParameter.ComponentPage"
//  INVALIDPARAMETER_COMPONENTPOSITION = "InvalidParameter.ComponentPosition"
//  INVALIDPARAMETER_COMPONENTTYPENOMATCHVALUE = "InvalidParameter.ComponentTypeNoMatchValue"
//  INVALIDPARAMETER_COMPONENTVALUE = "InvalidParameter.ComponentValue"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_ENDPOINT = "InvalidParameter.EndPoint"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_INVALIDORGANIZATIONNAME = "InvalidParameter.InvalidOrganizationName"
//  INVALIDPARAMETER_INVALIDVERIFYCHANNEL = "InvalidParameter.InvalidVerifyChannel"
//  INVALIDPARAMETER_MISSINGREQUIREDCOMPONENTVALUE = "InvalidParameter.MissingRequiredComponentValue"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETER_SEALID = "InvalidParameter.SealId"
//  INVALIDPARAMETER_SIGNCOMPONENTS = "InvalidParameter.SignComponents"
//  INVALIDPARAMETER_VERIFYCHANNEL = "InvalidParameter.VerifyChannel"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  MISSINGPARAMETER_APPROVERROLE = "MissingParameter.ApproverRole"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  MISSINGPARAMETER_FIELD = "MissingParameter.Field"
//  MISSINGPARAMETER_FILENAMES = "MissingParameter.FileNames"
//  MISSINGPARAMETER_FLOWAPPROVER = "MissingParameter.FlowApprover"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  MISSINGPARAMETER_KEYWORD = "MissingParameter.KeyWord"
//  MISSINGPARAMETER_RECIPIENT = "MissingParameter.Recipient"
//  MISSINGPARAMETER_RESOURCEID = "MissingParameter.ResourceId"
//  MISSINGPARAMETER_SERVERSIGNSEALID = "MissingParameter.ServerSignSealId"
//  MISSINGPARAMETER_SIGNCOMPONENTS = "MissingParameter.SignComponents"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERNOMATCHTEMPLATE = "OperationDenied.ApproverNoMatchTemplate"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_ERRNOSUPPORTINDIVIDUALHASORGANIZATIONNAME = "OperationDenied.ErrNoSupportIndividualHasOrganizationName"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FILENOMATCHRESOURCE = "OperationDenied.FileNoMatchResource"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_MANYRESOURCEID = "OperationDenied.ManyResourceId"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  OPERATIONDENIED_NOPERMISSIONUSESERVERSIGNSEAL = "OperationDenied.NoPermissionUseServerSignSeal"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOSUPPORTCOMPONENTTYPE = "OperationDenied.NoSupportComponentType"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_PERSONSERVERSIGNFORBID = "OperationDenied.PersonServerSignForbid"
//  OPERATIONDENIED_REQUIREDCOMPONENTNOTFILL = "OperationDenied.RequiredComponentNotFill"
//  OPERATIONDENIED_SERVERSIGNNOALLOWCOMPONENT = "OperationDenied.ServerSignNoAllowComponent"
//  OPERATIONDENIED_SERVERSIGNNOSUPPORTSIGNATURE = "OperationDenied.ServerSignNoSupportSignature"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_DOCUMENT = "ResourceNotFound.Document"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_NOTEXISTRESOURCE = "ResourceNotFound.NotExistResource"
//  RESOURCENOTFOUND_NOTEXISTTEMPLATE = "ResourceNotFound.NotExistTemplate"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFlowGroupByTemplatesWithContext(ctx context.Context, request *CreateFlowGroupByTemplatesRequest) (response *CreateFlowGroupByTemplatesResponse, err error) {
    if request == nil {
        request = NewCreateFlowGroupByTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowGroupByTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowGroupByTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowGroupSignReviewRequest() (request *CreateFlowGroupSignReviewRequest) {
    request = &CreateFlowGroupSignReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowGroupSignReview")
    
    
    return
}

func NewCreateFlowGroupSignReviewResponse() (response *CreateFlowGroupSignReviewResponse) {
    response = &CreateFlowGroupSignReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowGroupSignReview
// 提交合同组签署流程审批结果的适用场景包括：
//
// 
//
// 1. 在使用[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByFiles)或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByTemplates)创建合同组签署流程时，若指定了以下参数 为true，则可以调用此接口提交企业内部签署审批结果。即使是自动签署也需要进行审核通过才会进行签署。
//
//   - [FlowGroupInfo.NeedSignReview](https://qian.tencent.com/developers/companyApis/dataTypes/#flowgroupinfo)
//
//   - [ApproverInfo.ApproverNeedSignReview](https://qian.tencent.com/developers/companyApis/dataTypes/#approverinfo)
//
// 
//
// 
//
// 2. 同一合同组，同一签署人可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowGroupSignReview(request *CreateFlowGroupSignReviewRequest) (response *CreateFlowGroupSignReviewResponse, err error) {
    return c.CreateFlowGroupSignReviewWithContext(context.Background(), request)
}

// CreateFlowGroupSignReview
// 提交合同组签署流程审批结果的适用场景包括：
//
// 
//
// 1. 在使用[通过多文件创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByFiles)或[通过多模板创建合同组签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowGroupByTemplates)创建合同组签署流程时，若指定了以下参数 为true，则可以调用此接口提交企业内部签署审批结果。即使是自动签署也需要进行审核通过才会进行签署。
//
//   - [FlowGroupInfo.NeedSignReview](https://qian.tencent.com/developers/companyApis/dataTypes/#flowgroupinfo)
//
//   - [ApproverInfo.ApproverNeedSignReview](https://qian.tencent.com/developers/companyApis/dataTypes/#approverinfo)
//
// 
//
// 
//
// 2. 同一合同组，同一签署人可以多次提交签署审批结果，签署时的最后一个“审批结果”有效。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowGroupSignReviewWithContext(ctx context.Context, request *CreateFlowGroupSignReviewRequest) (response *CreateFlowGroupSignReviewResponse, err error) {
    if request == nil {
        request = NewCreateFlowGroupSignReviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowGroupSignReview require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowGroupSignReviewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowRemindsRequest() (request *CreateFlowRemindsRequest) {
    request = &CreateFlowRemindsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowReminds")
    
    
    return
}

func NewCreateFlowRemindsResponse() (response *CreateFlowRemindsResponse) {
    response = &CreateFlowRemindsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowReminds
// 指定需要批量催办的签署流程ID，批量催办合同，最多100个。需要符合以下条件的合同才可被催办：
//
// 
//
// 1. 发起合同时，**签署人的NotifyType需设置为sms**
//
// 2. 合同中当前状态为 **待签署** 的签署人是催办的对象
//
// 3. **每个合同只能催办一次**
//
// 
//
// **催办的效果**: 对方会收到如下的短信通知
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/3caf94b7f540fa5736270d38528d3a7b.png)
//
// 
//
// 注：`合同催办是白名单功能，请联系客户经理申请开白后使用`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowReminds(request *CreateFlowRemindsRequest) (response *CreateFlowRemindsResponse, err error) {
    return c.CreateFlowRemindsWithContext(context.Background(), request)
}

// CreateFlowReminds
// 指定需要批量催办的签署流程ID，批量催办合同，最多100个。需要符合以下条件的合同才可被催办：
//
// 
//
// 1. 发起合同时，**签署人的NotifyType需设置为sms**
//
// 2. 合同中当前状态为 **待签署** 的签署人是催办的对象
//
// 3. **每个合同只能催办一次**
//
// 
//
// **催办的效果**: 对方会收到如下的短信通知
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/3caf94b7f540fa5736270d38528d3a7b.png)
//
// 
//
// 注：`合同催办是白名单功能，请联系客户经理申请开白后使用`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowRemindsWithContext(ctx context.Context, request *CreateFlowRemindsRequest) (response *CreateFlowRemindsResponse, err error) {
    if request == nil {
        request = NewCreateFlowRemindsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowReminds require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowRemindsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowSignReviewRequest() (request *CreateFlowSignReviewRequest) {
    request = &CreateFlowSignReviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowSignReview")
    
    
    return
}

func NewCreateFlowSignReviewResponse() (response *CreateFlowSignReviewResponse) {
    response = &CreateFlowSignReviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowSignReview
// 提交企业流程审批结果 
//
// **当前存在两种审核操作：**
//
// <ul>
//
// <li>签署审核
//
// <ul>
//
// <li>在通过接口<ul>
//
// <li>CreateFlowByFiles</li>
//
// <li>CreateFlow</li>
//
// <li>CreateFlowGroupByTemplates</li>
//
// <li>CreateFlowGroupByFiles</li>
//
// <li>CreatePrepareFlow</li>
//
// </ul> 
//
// 发起签署流程时，通过指定NeedSignReview为true，则可以调用此接口，并指定operate=SignReview，以提交企业内部签署审批结果</li>
//
// <li>在通过接口
//
// <ul>
//
// <li>CreateFlowByFiles</li>
//
// <li>CreateFlow</li>
//
// <li>CreateFlowGroupByTemplates</li>
//
// <li>CreateFlowGroupByFiles</li>
//
// </ul>
//
// 发起签署流程时，通过指定签署人ApproverNeedSignReview为true，则可以调用此接口，并指定operate=SignReview，并指定RecipientId，以提交企业内部签署审批结果</li>
//
// </ul>
//
// </li>
//
// <li>发起审核
//
//  <ul>
//
// <li>通过接口CreatePrepareFlow指定发起后需要审核，那么可以调用此接口，并指定operate=CreateReview，以提交企业内部审批结果。可以多次提交审批结果，但一旦审批通过，后续提交的结果将无效
//
// </li>
//
// </ul>
//
// </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignReview(request *CreateFlowSignReviewRequest) (response *CreateFlowSignReviewResponse, err error) {
    return c.CreateFlowSignReviewWithContext(context.Background(), request)
}

// CreateFlowSignReview
// 提交企业流程审批结果 
//
// **当前存在两种审核操作：**
//
// <ul>
//
// <li>签署审核
//
// <ul>
//
// <li>在通过接口<ul>
//
// <li>CreateFlowByFiles</li>
//
// <li>CreateFlow</li>
//
// <li>CreateFlowGroupByTemplates</li>
//
// <li>CreateFlowGroupByFiles</li>
//
// <li>CreatePrepareFlow</li>
//
// </ul> 
//
// 发起签署流程时，通过指定NeedSignReview为true，则可以调用此接口，并指定operate=SignReview，以提交企业内部签署审批结果</li>
//
// <li>在通过接口
//
// <ul>
//
// <li>CreateFlowByFiles</li>
//
// <li>CreateFlow</li>
//
// <li>CreateFlowGroupByTemplates</li>
//
// <li>CreateFlowGroupByFiles</li>
//
// </ul>
//
// 发起签署流程时，通过指定签署人ApproverNeedSignReview为true，则可以调用此接口，并指定operate=SignReview，并指定RecipientId，以提交企业内部签署审批结果</li>
//
// </ul>
//
// </li>
//
// <li>发起审核
//
//  <ul>
//
// <li>通过接口CreatePrepareFlow指定发起后需要审核，那么可以调用此接口，并指定operate=CreateReview，以提交企业内部审批结果。可以多次提交审批结果，但一旦审批通过，后续提交的结果将无效
//
// </li>
//
// </ul>
//
// </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignReviewWithContext(ctx context.Context, request *CreateFlowSignReviewRequest) (response *CreateFlowSignReviewResponse, err error) {
    if request == nil {
        request = NewCreateFlowSignReviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowSignReview require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowSignReviewResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlowSignUrlRequest() (request *CreateFlowSignUrlRequest) {
    request = &CreateFlowSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateFlowSignUrl")
    
    
    return
}

func NewCreateFlowSignUrlResponse() (response *CreateFlowSignUrlResponse) {
    response = &CreateFlowSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateFlowSignUrl
// 该接口用于发起合同后，生成用户的签署链接 <br/>
//
// 
//
// **注意**
//
// 1. 该接口可生成签署链接的签署人必须仅有手写签名、时间类型和印章类型的签署控件，**不支持填写控件** 。
//
// 2. 该签署**链接有效期为30分钟**，过期后将失效，如需签署可重新创建签署链接 。
//
// 3. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。
//
// 跳转到小程序的实现，参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 4. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignUrl(request *CreateFlowSignUrlRequest) (response *CreateFlowSignUrlResponse, err error) {
    return c.CreateFlowSignUrlWithContext(context.Background(), request)
}

// CreateFlowSignUrl
// 该接口用于发起合同后，生成用户的签署链接 <br/>
//
// 
//
// **注意**
//
// 1. 该接口可生成签署链接的签署人必须仅有手写签名、时间类型和印章类型的签署控件，**不支持填写控件** 。
//
// 2. 该签署**链接有效期为30分钟**，过期后将失效，如需签署可重新创建签署链接 。
//
// 3. 该接口返回的签署链接适用于APP集成的场景，支持APP打开或浏览器直接打开，**不支持微信小程序嵌入**。
//
// 跳转到小程序的实现，参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>。
//
// 4. 因h5涉及人脸身份认证能力基于慧眼人脸核身，对Android和iOS系统均有一定要求， 因此<font color='red'>App嵌入H5签署合同需要按照慧眼提供的<a href="https://cloud.tencent.com/document/product/1007/61076">慧眼人脸核身兼容性文档</a>做兼容性适配</font>。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateFlowSignUrlWithContext(ctx context.Context, request *CreateFlowSignUrlRequest) (response *CreateFlowSignUrlResponse, err error) {
    if request == nil {
        request = NewCreateFlowSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlowSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlowSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrationDepartmentRequest() (request *CreateIntegrationDepartmentRequest) {
    request = &CreateIntegrationDepartmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateIntegrationDepartment")
    
    
    return
}

func NewCreateIntegrationDepartmentResponse() (response *CreateIntegrationDepartmentResponse) {
    response = &CreateIntegrationDepartmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrationDepartment
// 此接口（CreateIntegrationDepartment）用于创建企业的部门信息，支持绑定客户系统部门ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
func (c *Client) CreateIntegrationDepartment(request *CreateIntegrationDepartmentRequest) (response *CreateIntegrationDepartmentResponse, err error) {
    return c.CreateIntegrationDepartmentWithContext(context.Background(), request)
}

// CreateIntegrationDepartment
// 此接口（CreateIntegrationDepartment）用于创建企业的部门信息，支持绑定客户系统部门ID。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
func (c *Client) CreateIntegrationDepartmentWithContext(ctx context.Context, request *CreateIntegrationDepartmentRequest) (response *CreateIntegrationDepartmentResponse, err error) {
    if request == nil {
        request = NewCreateIntegrationDepartmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrationDepartment require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrationDepartmentResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrationEmployeesRequest() (request *CreateIntegrationEmployeesRequest) {
    request = &CreateIntegrationEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateIntegrationEmployees")
    
    
    return
}

func NewCreateIntegrationEmployeesResponse() (response *CreateIntegrationEmployeesResponse) {
    response = &CreateIntegrationEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrationEmployees
// 此接口（CreateIntegrationEmployees）用于创建企业员工。创建的员工初始化为未实名，如下图所示。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/2bdcc0d91ac3146b5e8c28811a78ffe9.png)
//
// 
//
// 支持以下场景
//
// <table>
//
// <tbody>
//
// <tr>
//
// <td>生成端</td>
//
// <td >入参</td>
//
// <td>提醒方式</td>
//
// </tr>
//
// <tr>
//
// <td>普通saas员工</td>
//
// <td>将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号</td>
//
// <td>发送短信通知员工（短信中带有认证加入企业的链接）   ![image]() </td>
//
// </tr>
//
// <tr>
//
// <td>企微员工</td>
//
// <td>将Employees 中的WeworkOpenId字段设置为企微员工明文的openid，需<font color="red">确保该企微员工在应用的可见范围内</font></td>
//
// <td>企微内部实名消息</td>
//
// </tr>
//
// <tr>
//
// <td>H5端 saas员工</td>
//
// <td>传递 InvitationNotifyType = H5，将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号，<font color="red">此场景不支持企微</font></td>
//
// <td>生成认证加入企业的H5链接，贵方可以通过自己的渠道触达到此员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 注意：
//
// 
//
// -  <b> 新增员工的手机号、OpenId不能与已加入员工重复</b>， 不管已加入员工的手机号、OpenId是否已经实名
//
// - <b>若通过手机号发现员工已经创建且信息一致（名字，openId等），则不会重复创建，但会发送短信或者生成链接提醒员工实名。</b>
//
// - jumpUrl 仅支持H5的邀请方式，回跳的url，使用前请联系对接的客户经理沟通，进行域名的配置。
//
// 
//
// 
//
// 
//
// 短信的样式
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/b6ad1b79e0adaaa41d282456c72a1ee6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationEmployees(request *CreateIntegrationEmployeesRequest) (response *CreateIntegrationEmployeesResponse, err error) {
    return c.CreateIntegrationEmployeesWithContext(context.Background(), request)
}

// CreateIntegrationEmployees
// 此接口（CreateIntegrationEmployees）用于创建企业员工。创建的员工初始化为未实名，如下图所示。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/2bdcc0d91ac3146b5e8c28811a78ffe9.png)
//
// 
//
// 支持以下场景
//
// <table>
//
// <tbody>
//
// <tr>
//
// <td>生成端</td>
//
// <td >入参</td>
//
// <td>提醒方式</td>
//
// </tr>
//
// <tr>
//
// <td>普通saas员工</td>
//
// <td>将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号</td>
//
// <td>发送短信通知员工（短信中带有认证加入企业的链接）   ![image]() </td>
//
// </tr>
//
// <tr>
//
// <td>企微员工</td>
//
// <td>将Employees 中的WeworkOpenId字段设置为企微员工明文的openid，需<font color="red">确保该企微员工在应用的可见范围内</font></td>
//
// <td>企微内部实名消息</td>
//
// </tr>
//
// <tr>
//
// <td>H5端 saas员工</td>
//
// <td>传递 InvitationNotifyType = H5，将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号，<font color="red">此场景不支持企微</font></td>
//
// <td>生成认证加入企业的H5链接，贵方可以通过自己的渠道触达到此员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 注意：
//
// 
//
// -  <b> 新增员工的手机号、OpenId不能与已加入员工重复</b>， 不管已加入员工的手机号、OpenId是否已经实名
//
// - <b>若通过手机号发现员工已经创建且信息一致（名字，openId等），则不会重复创建，但会发送短信或者生成链接提醒员工实名。</b>
//
// - jumpUrl 仅支持H5的邀请方式，回跳的url，使用前请联系对接的客户经理沟通，进行域名的配置。
//
// 
//
// 
//
// 
//
// 短信的样式
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/b6ad1b79e0adaaa41d282456c72a1ee6.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationEmployeesWithContext(ctx context.Context, request *CreateIntegrationEmployeesRequest) (response *CreateIntegrationEmployeesResponse, err error) {
    if request == nil {
        request = NewCreateIntegrationEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrationEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrationEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrationRoleRequest() (request *CreateIntegrationRoleRequest) {
    request = &CreateIntegrationRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateIntegrationRole")
    
    
    return
}

func NewCreateIntegrationRoleResponse() (response *CreateIntegrationRoleResponse) {
    response = &CreateIntegrationRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrationRole
// 此接口（CreateIntegrationRole）用来创建企业自定义的SaaS角色或集团角色。
//
// 
//
// 适用场景1：创建当前企业的自定义SaaS角色或集团角色，并且创建时不进行权限的设置（PermissionGroups 参数不传），角色中的权限内容可通过控制台编辑角色或通过接口 ModifyIntegrationRole 完成更新。
//
// 
//
// 适用场景2：创建当前企业的自定义SaaS角色或集团角色，并且创建时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 
//
// 适用场景3：创建集团角色时可同时设置角色管理的子企业列表，可通过设置 SubOrganizationIds 参数达到此效果。
//
// 
//
// 适用场景4：主企业代理子企业操作的场景，需要设置Agent参数，并且ProxyOrganizationId设置为子企业的id即可。
//
// 
//
// 注意事项：SaaS角色和集团角色对应的权限树是不一样的。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationRole(request *CreateIntegrationRoleRequest) (response *CreateIntegrationRoleResponse, err error) {
    return c.CreateIntegrationRoleWithContext(context.Background(), request)
}

// CreateIntegrationRole
// 此接口（CreateIntegrationRole）用来创建企业自定义的SaaS角色或集团角色。
//
// 
//
// 适用场景1：创建当前企业的自定义SaaS角色或集团角色，并且创建时不进行权限的设置（PermissionGroups 参数不传），角色中的权限内容可通过控制台编辑角色或通过接口 ModifyIntegrationRole 完成更新。
//
// 
//
// 适用场景2：创建当前企业的自定义SaaS角色或集团角色，并且创建时进行权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 
//
// 适用场景3：创建集团角色时可同时设置角色管理的子企业列表，可通过设置 SubOrganizationIds 参数达到此效果。
//
// 
//
// 适用场景4：主企业代理子企业操作的场景，需要设置Agent参数，并且ProxyOrganizationId设置为子企业的id即可。
//
// 
//
// 注意事项：SaaS角色和集团角色对应的权限树是不一样的。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationRoleWithContext(ctx context.Context, request *CreateIntegrationRoleRequest) (response *CreateIntegrationRoleResponse, err error) {
    if request == nil {
        request = NewCreateIntegrationRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrationRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrationRoleResponse()
    err = c.Send(request, response)
    return
}

func NewCreateIntegrationUserRolesRequest() (request *CreateIntegrationUserRolesRequest) {
    request = &CreateIntegrationUserRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateIntegrationUserRoles")
    
    
    return
}

func NewCreateIntegrationUserRolesResponse() (response *CreateIntegrationUserRolesResponse) {
    response = &CreateIntegrationUserRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateIntegrationUserRoles
// 此接口用于赋予员工指定的角色权限，如需解绑请使用 DeleteIntegrationRoleUsers 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationUserRoles(request *CreateIntegrationUserRolesRequest) (response *CreateIntegrationUserRolesResponse, err error) {
    return c.CreateIntegrationUserRolesWithContext(context.Background(), request)
}

// CreateIntegrationUserRoles
// 此接口用于赋予员工指定的角色权限，如需解绑请使用 DeleteIntegrationRoleUsers 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOPERMISSIONUSERESOURCE = "OperationDenied.NoPermissionUseResource"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateIntegrationUserRolesWithContext(ctx context.Context, request *CreateIntegrationUserRolesRequest) (response *CreateIntegrationUserRolesResponse, err error) {
    if request == nil {
        request = NewCreateIntegrationUserRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIntegrationUserRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIntegrationUserRolesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMultiFlowSignQRCodeRequest() (request *CreateMultiFlowSignQRCodeRequest) {
    request = &CreateMultiFlowSignQRCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateMultiFlowSignQRCode")
    
    
    return
}

func NewCreateMultiFlowSignQRCodeResponse() (response *CreateMultiFlowSignQRCodeResponse) {
    response = &CreateMultiFlowSignQRCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMultiFlowSignQRCode
// 此接口（CreateMultiFlowSignQRCode）用于创建一码多签签署码。
//
// 
//
// **适用场景**:
//
// 签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 
//
// **注意**:
//
// 1. 本接口适用于**发起方没有填写控件的 B2C或者单C模板**,  若是B2C模板,还要满足以下任意一个条件
//
//     - 模板中配置的签署顺序是无序
//
//     - B端企业的签署方式是静默签署
//
//     - B端企业是非首位签署
//
// 2. 通过一码多签签署码发起的合同，合同涉及到的回调消息可参考文档[合同发起及签署相关回调
//
// ]( https://qian.tencent.com/developers/company/callback_types_contracts_sign)
//
// 3. 用户通过一码多签签署码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档[签署二维码相关回调](https://qian.tencent.com/developers/company/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83)
//
// 
//
// 签署码的样式如下图:
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/27317cf5aacb094fb1dc6f94179a5148.png )
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_QREFFECTDAY = "InvalidParameter.QrEffectDay"
//  INVALIDPARAMETER_QRFLOWEFFECTDAY = "InvalidParameter.QrFlowEffectDay"
//  INVALIDPARAMETER_TEMPLATEID = "InvalidParameter.TemplateId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateMultiFlowSignQRCode(request *CreateMultiFlowSignQRCodeRequest) (response *CreateMultiFlowSignQRCodeResponse, err error) {
    return c.CreateMultiFlowSignQRCodeWithContext(context.Background(), request)
}

// CreateMultiFlowSignQRCode
// 此接口（CreateMultiFlowSignQRCode）用于创建一码多签签署码。
//
// 
//
// **适用场景**:
//
// 签署人可通过扫描二维码补充签署信息进行实名签署。常用于提前不知道签署人的身份信息场景，例如：劳务工招工、大批量员工入职等场景。
//
// 
//
// **注意**:
//
// 1. 本接口适用于**发起方没有填写控件的 B2C或者单C模板**,  若是B2C模板,还要满足以下任意一个条件
//
//     - 模板中配置的签署顺序是无序
//
//     - B端企业的签署方式是静默签署
//
//     - B端企业是非首位签署
//
// 2. 通过一码多签签署码发起的合同，合同涉及到的回调消息可参考文档[合同发起及签署相关回调
//
// ]( https://qian.tencent.com/developers/company/callback_types_contracts_sign)
//
// 3. 用户通过一码多签签署码发起合同时，因企业额度不足导致失败 会触发签署二维码相关回调,具体参考文档[签署二维码相关回调](https://qian.tencent.com/developers/company/callback_types_commons#%E7%AD%BE%E7%BD%B2%E4%BA%8C%E7%BB%B4%E7%A0%81%E7%9B%B8%E5%85%B3%E5%9B%9E%E8%B0%83)
//
// 
//
// 签署码的样式如下图:
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/27317cf5aacb094fb1dc6f94179a5148.png )
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QRCODECREATORSIGNCOMPONENTS = "FailedOperation.QrCodeCreatorSignComponents"
//  FAILEDOPERATION_QRCODESIGNUSERS = "FailedOperation.QrCodeSignUsers"
//  FAILEDOPERATION_QRCODETEMPLATEID = "FailedOperation.QrCodeTemplateId"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_QREFFECTDAY = "InvalidParameter.QrEffectDay"
//  INVALIDPARAMETER_QRFLOWEFFECTDAY = "InvalidParameter.QrFlowEffectDay"
//  INVALIDPARAMETER_TEMPLATEID = "InvalidParameter.TemplateId"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_NOVERIFY = "OperationDenied.NoVerify"
//  OPERATIONDENIED_USERNOTINORGANIZATION = "OperationDenied.UserNotInOrganization"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateMultiFlowSignQRCodeWithContext(ctx context.Context, request *CreateMultiFlowSignQRCodeRequest) (response *CreateMultiFlowSignQRCodeResponse, err error) {
    if request == nil {
        request = NewCreateMultiFlowSignQRCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMultiFlowSignQRCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMultiFlowSignQRCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationBatchSignUrlRequest() (request *CreateOrganizationBatchSignUrlRequest) {
    request = &CreateOrganizationBatchSignUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateOrganizationBatchSignUrl")
    
    
    return
}

func NewCreateOrganizationBatchSignUrlResponse() (response *CreateOrganizationBatchSignUrlResponse) {
    response = &CreateOrganizationBatchSignUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationBatchSignUrl
// 使用此接口，您可以创建企业批量签署链接，员工只需点击链接即可跳转至控制台进行批量签署。<br/>
//
// 附注：
//
// - 员工必须在企业下完成实名认证，且需作为批量签署合同的签署方。
//
// - 如有UserId，应以UserId为主要标识；如果没有UserId，则必须填写Name和Mobile信息。
//
// - 仅支持待签署状态的合同生成签署链接。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateOrganizationBatchSignUrl(request *CreateOrganizationBatchSignUrlRequest) (response *CreateOrganizationBatchSignUrlResponse, err error) {
    return c.CreateOrganizationBatchSignUrlWithContext(context.Background(), request)
}

// CreateOrganizationBatchSignUrl
// 使用此接口，您可以创建企业批量签署链接，员工只需点击链接即可跳转至控制台进行批量签署。<br/>
//
// 附注：
//
// - 员工必须在企业下完成实名认证，且需作为批量签署合同的签署方。
//
// - 如有UserId，应以UserId为主要标识；如果没有UserId，则必须填写Name和Mobile信息。
//
// - 仅支持待签署状态的合同生成签署链接。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateOrganizationBatchSignUrlWithContext(ctx context.Context, request *CreateOrganizationBatchSignUrlRequest) (response *CreateOrganizationBatchSignUrlResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationBatchSignUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationBatchSignUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationBatchSignUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrganizationInfoChangeUrlRequest() (request *CreateOrganizationInfoChangeUrlRequest) {
    request = &CreateOrganizationInfoChangeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateOrganizationInfoChangeUrl")
    
    
    return
}

func NewCreateOrganizationInfoChangeUrlResponse() (response *CreateOrganizationInfoChangeUrlResponse) {
    response = &CreateOrganizationInfoChangeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOrganizationInfoChangeUrl
// 此接口（CreateOrganizationInfoChangeUrl）用于创建企业信息变更链接，支持创建企业超管变更链接或企业基础信息变更链接，通过入参ChangeType指定。
//
// 
//
// 
//
// <h3 id="1-企业超管变更">1. 企业超管变更</h3>
//
// 
//
// <p>换成企业的其他员工来当超管</p>
//
// 
//
// <h3 id="2-企业基础信息变更">2. 企业基础信息变更</h3>
//
// 
//
// <h4 id="可以变动">可以变动</h4>
//
// 
//
// <ul>
//
// <li>企业名称<br>
//
// </li>
//
// <li>法定代表人姓名(新法人有邀请链接)<br>
//
// </li>
//
// <li>企业地址和所在地</li>
//
// </ul>
//
// 
//
// <h4 id="不可变动">不可变动</h4>
//
// 
//
// <ul>
//
// <li>统一社会信用代码<br>
//
// </li>
//
// <li>企业主体类型</li>
//
// </ul>
//
// 
//
// <p>如果企业名称变动会引起下面的变动</p>
//
// 
//
// <ul>
//
// <li>合同:   老合同不做任何处理,   新发起的合同需要用新的企业名字作为签署方, 否则无法签署</li>
//
// <li>印章:   会删除所有的印章所有的机构公章和合同专用章,  然后用新企业名称生成新的机构公章 和合同专用章,  而法人章, 财务专用章和人事专用章不会处理</li>
//
// <li>证书:   企业证书会重新请求CA机构用新企业名称生成新的证书</li>
//
// </ul>
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOrganizationInfoChangeUrl(request *CreateOrganizationInfoChangeUrlRequest) (response *CreateOrganizationInfoChangeUrlResponse, err error) {
    return c.CreateOrganizationInfoChangeUrlWithContext(context.Background(), request)
}

// CreateOrganizationInfoChangeUrl
// 此接口（CreateOrganizationInfoChangeUrl）用于创建企业信息变更链接，支持创建企业超管变更链接或企业基础信息变更链接，通过入参ChangeType指定。
//
// 
//
// 
//
// <h3 id="1-企业超管变更">1. 企业超管变更</h3>
//
// 
//
// <p>换成企业的其他员工来当超管</p>
//
// 
//
// <h3 id="2-企业基础信息变更">2. 企业基础信息变更</h3>
//
// 
//
// <h4 id="可以变动">可以变动</h4>
//
// 
//
// <ul>
//
// <li>企业名称<br>
//
// </li>
//
// <li>法定代表人姓名(新法人有邀请链接)<br>
//
// </li>
//
// <li>企业地址和所在地</li>
//
// </ul>
//
// 
//
// <h4 id="不可变动">不可变动</h4>
//
// 
//
// <ul>
//
// <li>统一社会信用代码<br>
//
// </li>
//
// <li>企业主体类型</li>
//
// </ul>
//
// 
//
// <p>如果企业名称变动会引起下面的变动</p>
//
// 
//
// <ul>
//
// <li>合同:   老合同不做任何处理,   新发起的合同需要用新的企业名字作为签署方, 否则无法签署</li>
//
// <li>印章:   会删除所有的印章所有的机构公章和合同专用章,  然后用新企业名称生成新的机构公章 和合同专用章,  而法人章, 财务专用章和人事专用章不会处理</li>
//
// <li>证书:   企业证书会重新请求CA机构用新企业名称生成新的证书</li>
//
// </ul>
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOrganizationInfoChangeUrlWithContext(ctx context.Context, request *CreateOrganizationInfoChangeUrlRequest) (response *CreateOrganizationInfoChangeUrlResponse, err error) {
    if request == nil {
        request = NewCreateOrganizationInfoChangeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOrganizationInfoChangeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOrganizationInfoChangeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonAuthCertificateImageRequest() (request *CreatePersonAuthCertificateImageRequest) {
    request = &CreatePersonAuthCertificateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreatePersonAuthCertificateImage")
    
    
    return
}

func NewCreatePersonAuthCertificateImageResponse() (response *CreatePersonAuthCertificateImageResponse) {
    response = &CreatePersonAuthCertificateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePersonAuthCertificateImage
// 获取个人用户认证证书图片下载URL
//
// 
//
// 个人用户认证证书图片样式如下图
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreatePersonAuthCertificateImage.png)
//
// 
//
// 注:  
//
// <ul>
//
// <li>只能获取个人用户证明图片, 企业员工的暂不支持</li>
//
// <li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。  </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePersonAuthCertificateImage(request *CreatePersonAuthCertificateImageRequest) (response *CreatePersonAuthCertificateImageResponse, err error) {
    return c.CreatePersonAuthCertificateImageWithContext(context.Background(), request)
}

// CreatePersonAuthCertificateImage
// 获取个人用户认证证书图片下载URL
//
// 
//
// 个人用户认证证书图片样式如下图
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreatePersonAuthCertificateImage.png)
//
// 
//
// 注:  
//
// <ul>
//
// <li>只能获取个人用户证明图片, 企业员工的暂不支持</li>
//
// <li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。  </li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreatePersonAuthCertificateImageWithContext(ctx context.Context, request *CreatePersonAuthCertificateImageRequest) (response *CreatePersonAuthCertificateImageResponse, err error) {
    if request == nil {
        request = NewCreatePersonAuthCertificateImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePersonAuthCertificateImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePersonAuthCertificateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrepareFlowRequest() (request *CreatePrepareFlowRequest) {
    request = &CreatePrepareFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreatePrepareFlow")
    
    
    return
}

func NewCreatePrepareFlowResponse() (response *CreatePrepareFlowResponse) {
    response = &CreatePrepareFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePrepareFlow
// 创建发起流程web页面
//
// <br/>适用场景：通过该接口（CreatePrepareFlow）传入合同文件/模板编号及签署人信息，可获得发起流程的可嵌入页面，在页面完成签署控件等信息的编辑与确认后，快速发起流程。
//
// <br/>注：该接口包含模板/文件发起流程的全部功能，调用接口后不会立即发起，需在可嵌入页面点击按钮进行发起流程。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreatePrepareFlow(request *CreatePrepareFlowRequest) (response *CreatePrepareFlowResponse, err error) {
    return c.CreatePrepareFlowWithContext(context.Background(), request)
}

// CreatePrepareFlow
// 创建发起流程web页面
//
// <br/>适用场景：通过该接口（CreatePrepareFlow）传入合同文件/模板编号及签署人信息，可获得发起流程的可嵌入页面，在页面完成签署控件等信息的编辑与确认后，快速发起流程。
//
// <br/>注：该接口包含模板/文件发起流程的全部功能，调用接口后不会立即发起，需在可嵌入页面点击按钮进行发起流程。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreatePrepareFlowWithContext(ctx context.Context, request *CreatePrepareFlowRequest) (response *CreatePrepareFlowResponse, err error) {
    if request == nil {
        request = NewCreatePrepareFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePrepareFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePrepareFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePreparedPersonalEsignRequest() (request *CreatePreparedPersonalEsignRequest) {
    request = &CreatePreparedPersonalEsignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreatePreparedPersonalEsign")
    
    
    return
}

func NewCreatePreparedPersonalEsignResponse() (response *CreatePreparedPersonalEsignResponse) {
    response = &CreatePreparedPersonalEsignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePreparedPersonalEsign
// 本接口（CreatePreparedPersonalEsign）用于创建导入个人印章（处方单场景专用，使用此接口请与客户经理确认）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IDCARDNUMBERCHECKFAILED = "FailedOperation.IdCardNumberCheckFailed"
//  FAILEDOPERATION_LICENSENOQUOTA = "FailedOperation.LicenseNoQuota"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOBILE = "InvalidParameterValue.InvalidMobile"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
func (c *Client) CreatePreparedPersonalEsign(request *CreatePreparedPersonalEsignRequest) (response *CreatePreparedPersonalEsignResponse, err error) {
    return c.CreatePreparedPersonalEsignWithContext(context.Background(), request)
}

// CreatePreparedPersonalEsign
// 本接口（CreatePreparedPersonalEsign）用于创建导入个人印章（处方单场景专用，使用此接口请与客户经理确认）。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IDCARDNUMBERCHECKFAILED = "FailedOperation.IdCardNumberCheckFailed"
//  FAILEDOPERATION_LICENSENOQUOTA = "FailedOperation.LicenseNoQuota"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDMOBILE = "InvalidParameterValue.InvalidMobile"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
func (c *Client) CreatePreparedPersonalEsignWithContext(ctx context.Context, request *CreatePreparedPersonalEsignRequest) (response *CreatePreparedPersonalEsignResponse, err error) {
    if request == nil {
        request = NewCreatePreparedPersonalEsignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePreparedPersonalEsign require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePreparedPersonalEsignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateReleaseFlowRequest() (request *CreateReleaseFlowRequest) {
    request = &CreateReleaseFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateReleaseFlow")
    
    
    return
}

func NewCreateReleaseFlowResponse() (response *CreateReleaseFlowResponse) {
    response = &CreateReleaseFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateReleaseFlow
// 发起解除协议的主要应用场景为：基于一份已经签署的合同（签署流程），进行解除操作。
//
// 解除协议的模板是官方提供 ，经过提供法务审核，暂不支持自定义。
//
// 
//
// 注意：
//
// <ul><li><code>原合同必须签署完</code>成后才能发起解除协议。</li>
//
// <li>只有原合同企业类型的参与人才能发起解除协议，<code>个人参与方不能发起解除协议</code>。</li>
//
// <li>原合同个人类型参与人必须是解除协议的参与人，<code>不能更换其他第三方个人</code>参与解除协议。</li>
//
// <li>如果原合同企业参与人无法参与解除协议，可以指定同企业具有同等权限的<code>企业员工代为处理</code>。</li>
//
// <li>发起解除协议同发起其他企业合同一样，也会参与合同<code>扣费</code>，扣费标准同其他类型合同。</li>
//
// <li>在解除协议发起之后，原合同的状态将转变为解除中。一旦解除协议签署完毕，原合同及解除协议均变为已解除状态。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateReleaseFlow(request *CreateReleaseFlowRequest) (response *CreateReleaseFlowResponse, err error) {
    return c.CreateReleaseFlowWithContext(context.Background(), request)
}

// CreateReleaseFlow
// 发起解除协议的主要应用场景为：基于一份已经签署的合同（签署流程），进行解除操作。
//
// 解除协议的模板是官方提供 ，经过提供法务审核，暂不支持自定义。
//
// 
//
// 注意：
//
// <ul><li><code>原合同必须签署完</code>成后才能发起解除协议。</li>
//
// <li>只有原合同企业类型的参与人才能发起解除协议，<code>个人参与方不能发起解除协议</code>。</li>
//
// <li>原合同个人类型参与人必须是解除协议的参与人，<code>不能更换其他第三方个人</code>参与解除协议。</li>
//
// <li>如果原合同企业参与人无法参与解除协议，可以指定同企业具有同等权限的<code>企业员工代为处理</code>。</li>
//
// <li>发起解除协议同发起其他企业合同一样，也会参与合同<code>扣费</code>，扣费标准同其他类型合同。</li>
//
// <li>在解除协议发起之后，原合同的状态将转变为解除中。一旦解除协议签署完毕，原合同及解除协议均变为已解除状态。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_FLOWHASDOCUMENT = "FailedOperation.FlowHasDocument"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  FAILEDOPERATION_REQUESTLIMITEXCEEDED = "FailedOperation.RequestLimitExceeded"
//  FAILEDOPERATION_USERINFONOMATCH = "FailedOperation.UserInfoNoMatch"
//  INTERNALERROR_DBINSERT = "InternalError.DbInsert"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_APPROVERTYPE = "InvalidParameter.ApproverType"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_CCNUM = "InvalidParameter.CcNum"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_CUSTOMSHOWMAP = "InvalidParameter.CustomShowMap"
//  INVALIDPARAMETER_FLOWCALLBACKURL = "InvalidParameter.FlowCallbackUrl"
//  INVALIDPARAMETER_FLOWDEADLINE = "InvalidParameter.FlowDeadLine"
//  INVALIDPARAMETER_FLOWDESCRIPTION = "InvalidParameter.FlowDescription"
//  INVALIDPARAMETER_FLOWNAME = "InvalidParameter.FlowName"
//  INVALIDPARAMETER_FLOWTYPE = "InvalidParameter.FlowType"
//  INVALIDPARAMETER_FLOWUSERDATA = "InvalidParameter.FlowUserData"
//  INVALIDPARAMETER_FROMSOURCE = "InvalidParameter.FromSource"
//  INVALIDPARAMETER_IDCARDVALIDITYOVERLIMIT = "InvalidParameter.IdCardValidityOverLimit"
//  INVALIDPARAMETER_INVALIDMOBILE = "InvalidParameter.InvalidMobile"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_NONSUPPORTMOBILE = "InvalidParameter.NonsupportMobile"
//  INVALIDPARAMETER_NOTIFYTYPE = "InvalidParameter.NotifyType"
//  INVALIDPARAMETER_ORGANIZATIONNAME = "InvalidParameter.OrganizationName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_PERSONAUTOSIGNTAG = "InvalidParameter.PersonAutoSignTag"
//  INVALIDPARAMETER_PREREADTIME = "InvalidParameter.PreReadTime"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER_APPROVERMOBILE = "MissingParameter.ApproverMobile"
//  MISSINGPARAMETER_APPROVERNAME = "MissingParameter.ApproverName"
//  MISSINGPARAMETER_APPROVERORGANIZATIONINFO = "MissingParameter.ApproverOrganizationInfo"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_APPROVERREPEAT = "OperationDenied.ApproverRepeat"
//  OPERATIONDENIED_BRANCHSENDFLOWTOPARENTNOTALLOW = "OperationDenied.BranchSendFlowToParentNotAllow"
//  OPERATIONDENIED_CCFORBID = "OperationDenied.CcForbid"
//  OPERATIONDENIED_CCUSERREPEAT = "OperationDenied.CcUserRepeat"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_ORGUNIFORMSOCIALCREDITCODEERR = "OperationDenied.OrgUniformSocialCreditCodeErr"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  OPERATIONDENIED_OVERSEAFORBID = "OperationDenied.OverSeaForbid"
//  OPERATIONDENIED_PERSONHASNOSIGNATURE = "OperationDenied.PersonHasNoSignature"
//  OPERATIONDENIED_PERSONNOOPENSERVERSIGN = "OperationDenied.PersonNoOpenServerSign"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_AUTHACTIVEORGANIZATION = "ResourceNotFound.AuthActiveOrganization"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_SUPERADMIN = "ResourceNotFound.SuperAdmin"
//  RESOURCENOTFOUND_USER = "ResourceNotFound.User"
//  RESOURCENOTFOUND_VERIFYUSER = "ResourceNotFound.VerifyUser"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateReleaseFlowWithContext(ctx context.Context, request *CreateReleaseFlowRequest) (response *CreateReleaseFlowResponse, err error) {
    if request == nil {
        request = NewCreateReleaseFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateReleaseFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateReleaseFlowResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchemeUrlRequest() (request *CreateSchemeUrlRequest) {
    request = &CreateSchemeUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateSchemeUrl")
    
    
    return
}

func NewCreateSchemeUrlResponse() (response *CreateSchemeUrlResponse) {
    response = &CreateSchemeUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSchemeUrl
// 获取跳转至腾讯电子签小程序的签署链接
//
// 
//
// 适用场景：如果需要签署人在自己的APP、小程序、H5应用中签署，可以通过此接口获取跳转腾讯电子签小程序的签署跳转链接。
//
// 
//
// 跳转到小程序的实现，参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>
//
// 
//
// 注：
//
// <ul><li>1. 如果签署人是在PC端扫码签署，可以通过生成跳转链接自主转换成二维码，让签署人在PC端扫码签署</li>
//
// <li>2. 签署链接的有效期为<font color="red">90天</font>，超过有效期链接不可用</li>
//
// <li>3. 如果需跳转详情页（即PathType值为1）进行填写或签署合同，需指定签署方信息:姓名、手机号码、企业名称等，才能生成正确的跳转链接</li>
//
// <li>4. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）</li></ul>
//
// 
//
// 其中小程序的原始Id如下，或者查看小程序信息自助获取。
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSchemeUrl(request *CreateSchemeUrlRequest) (response *CreateSchemeUrlResponse, err error) {
    return c.CreateSchemeUrlWithContext(context.Background(), request)
}

// CreateSchemeUrl
// 获取跳转至腾讯电子签小程序的签署链接
//
// 
//
// 适用场景：如果需要签署人在自己的APP、小程序、H5应用中签署，可以通过此接口获取跳转腾讯电子签小程序的签署跳转链接。
//
// 
//
// 跳转到小程序的实现，参考微信官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>
//
// 
//
// 注：
//
// <ul><li>1. 如果签署人是在PC端扫码签署，可以通过生成跳转链接自主转换成二维码，让签署人在PC端扫码签署</li>
//
// <li>2. 签署链接的有效期为<font color="red">90天</font>，超过有效期链接不可用</li>
//
// <li>3. 如果需跳转详情页（即PathType值为1）进行填写或签署合同，需指定签署方信息:姓名、手机号码、企业名称等，才能生成正确的跳转链接</li>
//
// <li>4. <font color="red">生成的链路后面不能再增加参数</font>（会出现覆盖链接中已有参数导致错误）</li></ul>
//
// 
//
// 其中小程序的原始Id如下，或者查看小程序信息自助获取。
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_MOBILE = "InvalidParameter.Mobile"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOSUPPORTJUMPPAGE = "OperationDenied.NoSupportJumpPage"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) CreateSchemeUrlWithContext(ctx context.Context, request *CreateSchemeUrlRequest) (response *CreateSchemeUrlResponse, err error) {
    if request == nil {
        request = NewCreateSchemeUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchemeUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSchemeUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSealRequest() (request *CreateSealRequest) {
    request = &CreateSealRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateSeal")
    
    
    return
}

func NewCreateSealResponse() (response *CreateSealResponse) {
    response = &CreateSealResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSeal
// 本接口（CreateSeal）用于创建企业电子印章，支持创建企业公章，合同章，财务专用章和人事专用章创建。
//
// 
//
// 1. 可以**通过图片**创建印章，图片最大5MB
//
// 2. 可以**系统创建**创建印章, 系统创建的印章样子下图(样式可以调整)
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
func (c *Client) CreateSeal(request *CreateSealRequest) (response *CreateSealResponse, err error) {
    return c.CreateSealWithContext(context.Background(), request)
}

// CreateSeal
// 本接口（CreateSeal）用于创建企业电子印章，支持创建企业公章，合同章，财务专用章和人事专用章创建。
//
// 
//
// 1. 可以**通过图片**创建印章，图片最大5MB
//
// 2. 可以**系统创建**创建印章, 系统创建的印章样子下图(样式可以调整)
//
// 
//
// ![image](https://dyn.ess.tencent.cn/guide/capi/CreateSealByImage.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
func (c *Client) CreateSealWithContext(ctx context.Context, request *CreateSealRequest) (response *CreateSealResponse, err error) {
    if request == nil {
        request = NewCreateSealRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSeal require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSealResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSealPolicyRequest() (request *CreateSealPolicyRequest) {
    request = &CreateSealPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateSealPolicy")
    
    
    return
}

func NewCreateSealPolicyResponse() (response *CreateSealPolicyResponse) {
    response = &CreateSealPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSealPolicy
// 本接口（CreateSealPolicy）用于对企业员工进行印章授权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSealPolicy(request *CreateSealPolicyRequest) (response *CreateSealPolicyResponse, err error) {
    return c.CreateSealPolicyWithContext(context.Background(), request)
}

// CreateSealPolicy
// 本接口（CreateSealPolicy）用于对企业员工进行印章授权
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateSealPolicyWithContext(ctx context.Context, request *CreateSealPolicyRequest) (response *CreateSealPolicyResponse, err error) {
    if request == nil {
        request = NewCreateSealPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSealPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSealPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserAutoSignEnableUrlRequest() (request *CreateUserAutoSignEnableUrlRequest) {
    request = &CreateUserAutoSignEnableUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateUserAutoSignEnableUrl")
    
    
    return
}

func NewCreateUserAutoSignEnableUrlResponse() (response *CreateUserAutoSignEnableUrlResponse) {
    response = &CreateUserAutoSignEnableUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserAutoSignEnableUrl
// 获取个人用户自动签的开通链接。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
func (c *Client) CreateUserAutoSignEnableUrl(request *CreateUserAutoSignEnableUrlRequest) (response *CreateUserAutoSignEnableUrlResponse, err error) {
    return c.CreateUserAutoSignEnableUrlWithContext(context.Background(), request)
}

// CreateUserAutoSignEnableUrl
// 获取个人用户自动签的开通链接。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_INVALIDAPPROVERAGE = "OperationDenied.InvalidApproverAge"
func (c *Client) CreateUserAutoSignEnableUrlWithContext(ctx context.Context, request *CreateUserAutoSignEnableUrlRequest) (response *CreateUserAutoSignEnableUrlResponse, err error) {
    if request == nil {
        request = NewCreateUserAutoSignEnableUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserAutoSignEnableUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserAutoSignEnableUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserAutoSignSealUrlRequest() (request *CreateUserAutoSignSealUrlRequest) {
    request = &CreateUserAutoSignSealUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateUserAutoSignSealUrl")
    
    
    return
}

func NewCreateUserAutoSignSealUrlResponse() (response *CreateUserAutoSignSealUrlResponse) {
    response = &CreateUserAutoSignSealUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserAutoSignSealUrl
// 获取设置自动签印章小程序链接。
//
// 
//
// 注意：
//
// <ul><li>需要<code>企业开通自动签</code>后使用。</li>
//
// <li>仅支持<code>已经开通了自动签的个人</code>更换自动签印章。</li>
//
// <li>链接有效期默认7天，<code>最多30天</code>。</li>
//
// <li>该接口的链接适用于<code>小程序</code>端。</li>
//
// <li>该接口不会扣除您的合同套餐，暂不参与计费。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateUserAutoSignSealUrl(request *CreateUserAutoSignSealUrlRequest) (response *CreateUserAutoSignSealUrlResponse, err error) {
    return c.CreateUserAutoSignSealUrlWithContext(context.Background(), request)
}

// CreateUserAutoSignSealUrl
// 获取设置自动签印章小程序链接。
//
// 
//
// 注意：
//
// <ul><li>需要<code>企业开通自动签</code>后使用。</li>
//
// <li>仅支持<code>已经开通了自动签的个人</code>更换自动签印章。</li>
//
// <li>链接有效期默认7天，<code>最多30天</code>。</li>
//
// <li>该接口的链接适用于<code>小程序</code>端。</li>
//
// <li>该接口不会扣除您的合同套餐，暂不参与计费。</li></ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateUserAutoSignSealUrlWithContext(ctx context.Context, request *CreateUserAutoSignSealUrlRequest) (response *CreateUserAutoSignSealUrlResponse, err error) {
    if request == nil {
        request = NewCreateUserAutoSignSealUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserAutoSignSealUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserAutoSignSealUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateWebThemeConfigRequest() (request *CreateWebThemeConfigRequest) {
    request = &CreateWebThemeConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "CreateWebThemeConfig")
    
    
    return
}

func NewCreateWebThemeConfigResponse() (response *CreateWebThemeConfigResponse) {
    response = &CreateWebThemeConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateWebThemeConfig
// 用来设置本企业嵌入式页面个性化主题配置（例如是否展示电子签logo、定义主题色等），设置后获取的web签署界面都会使用此配置进行展示。
//
// 
//
// 如果多次调用，会以最后一次的配置为准
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateWebThemeConfig(request *CreateWebThemeConfigRequest) (response *CreateWebThemeConfigResponse, err error) {
    return c.CreateWebThemeConfigWithContext(context.Background(), request)
}

// CreateWebThemeConfig
// 用来设置本企业嵌入式页面个性化主题配置（例如是否展示电子签logo、定义主题色等），设置后获取的web签署界面都会使用此配置进行展示。
//
// 
//
// 如果多次调用，会以最后一次的配置为准
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateWebThemeConfigWithContext(ctx context.Context, request *CreateWebThemeConfigRequest) (response *CreateWebThemeConfigResponse, err error) {
    if request == nil {
        request = NewCreateWebThemeConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateWebThemeConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateWebThemeConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteExtendedServiceAuthInfosRequest() (request *DeleteExtendedServiceAuthInfosRequest) {
    request = &DeleteExtendedServiceAuthInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DeleteExtendedServiceAuthInfos")
    
    
    return
}

func NewDeleteExtendedServiceAuthInfosResponse() (response *DeleteExtendedServiceAuthInfosResponse) {
    response = &DeleteExtendedServiceAuthInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteExtendedServiceAuthInfos
// 删除企业扩展服务授权，当前仅支持 “企业自动签” 和“批量签署”  的取消授权。
//
// 该接口作用和电子签控制台 企业设置-扩展服务-企业自动签署和批量签署授权 两个模块功能相同，可通过该接口取消企业员工授权。
//
// 
//
// 注：支持集团代子企业操作，请联系运营开通此功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteExtendedServiceAuthInfos(request *DeleteExtendedServiceAuthInfosRequest) (response *DeleteExtendedServiceAuthInfosResponse, err error) {
    return c.DeleteExtendedServiceAuthInfosWithContext(context.Background(), request)
}

// DeleteExtendedServiceAuthInfos
// 删除企业扩展服务授权，当前仅支持 “企业自动签” 和“批量签署”  的取消授权。
//
// 该接口作用和电子签控制台 企业设置-扩展服务-企业自动签署和批量签署授权 两个模块功能相同，可通过该接口取消企业员工授权。
//
// 
//
// 注：支持集团代子企业操作，请联系运营开通此功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DeleteExtendedServiceAuthInfosWithContext(ctx context.Context, request *DeleteExtendedServiceAuthInfosRequest) (response *DeleteExtendedServiceAuthInfosResponse, err error) {
    if request == nil {
        request = NewDeleteExtendedServiceAuthInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteExtendedServiceAuthInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteExtendedServiceAuthInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntegrationDepartmentRequest() (request *DeleteIntegrationDepartmentRequest) {
    request = &DeleteIntegrationDepartmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DeleteIntegrationDepartment")
    
    
    return
}

func NewDeleteIntegrationDepartmentResponse() (response *DeleteIntegrationDepartmentResponse) {
    response = &DeleteIntegrationDepartmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIntegrationDepartment
// 此接口（DeleteIntegrationDepartment）用于删除企业的部门信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteIntegrationDepartment(request *DeleteIntegrationDepartmentRequest) (response *DeleteIntegrationDepartmentResponse, err error) {
    return c.DeleteIntegrationDepartmentWithContext(context.Background(), request)
}

// DeleteIntegrationDepartment
// 此接口（DeleteIntegrationDepartment）用于删除企业的部门信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SERIALIZE = "InternalError.Serialize"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeleteIntegrationDepartmentWithContext(ctx context.Context, request *DeleteIntegrationDepartmentRequest) (response *DeleteIntegrationDepartmentResponse, err error) {
    if request == nil {
        request = NewDeleteIntegrationDepartmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntegrationDepartment require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntegrationDepartmentResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntegrationEmployeesRequest() (request *DeleteIntegrationEmployeesRequest) {
    request = &DeleteIntegrationEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DeleteIntegrationEmployees")
    
    
    return
}

func NewDeleteIntegrationEmployeesResponse() (response *DeleteIntegrationEmployeesResponse) {
    response = &DeleteIntegrationEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIntegrationEmployees
// 该接口（DeleteIntegrationEmployees）用于离职本企业员工，同时可选择是否进行离职交接。
//
// 
//
// 
//
// - 如果该员工没有未处理的合同，可不设置交接人的ReceiveUserId或ReceiveOpenId进行离职操作。
//
// - 如果该员工有未处理的合同，需要设置ReceiveUserId或ReceiveOpenId表示交接的负责人，交接后员工会进行离职操作。
//
// 
//
// 未处理的合同包括以下：
//
// 
//
// - 待签署的合同（包括顺序签署还没有轮到的合同，此类合同某些情况可能不会出现在用户的列表中）。
//
// - 待填写的合同。
//
// - 待解除的合同等。
//
// 
//
// 注：
//
// 1. <font color="red">超管或法人身份的员工不能被离职</font>， 需要在控制台或小程序更换法人和超管后进行离职删除。
//
// 2. <font color="red">员工存在待处理合同时必须交接</font>后才能离职无人交接时不能被离职删除。
//
// 3. 未实名的员工可以直接离职，不用交接合同
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteIntegrationEmployees(request *DeleteIntegrationEmployeesRequest) (response *DeleteIntegrationEmployeesResponse, err error) {
    return c.DeleteIntegrationEmployeesWithContext(context.Background(), request)
}

// DeleteIntegrationEmployees
// 该接口（DeleteIntegrationEmployees）用于离职本企业员工，同时可选择是否进行离职交接。
//
// 
//
// 
//
// - 如果该员工没有未处理的合同，可不设置交接人的ReceiveUserId或ReceiveOpenId进行离职操作。
//
// - 如果该员工有未处理的合同，需要设置ReceiveUserId或ReceiveOpenId表示交接的负责人，交接后员工会进行离职操作。
//
// 
//
// 未处理的合同包括以下：
//
// 
//
// - 待签署的合同（包括顺序签署还没有轮到的合同，此类合同某些情况可能不会出现在用户的列表中）。
//
// - 待填写的合同。
//
// - 待解除的合同等。
//
// 
//
// 注：
//
// 1. <font color="red">超管或法人身份的员工不能被离职</font>， 需要在控制台或小程序更换法人和超管后进行离职删除。
//
// 2. <font color="red">员工存在待处理合同时必须交接</font>后才能离职无人交接时不能被离职删除。
//
// 3. 未实名的员工可以直接离职，不用交接合同
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteIntegrationEmployeesWithContext(ctx context.Context, request *DeleteIntegrationEmployeesRequest) (response *DeleteIntegrationEmployeesResponse, err error) {
    if request == nil {
        request = NewDeleteIntegrationEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntegrationEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntegrationEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIntegrationRoleUsersRequest() (request *DeleteIntegrationRoleUsersRequest) {
    request = &DeleteIntegrationRoleUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DeleteIntegrationRoleUsers")
    
    
    return
}

func NewDeleteIntegrationRoleUsersResponse() (response *DeleteIntegrationRoleUsersResponse) {
    response = &DeleteIntegrationRoleUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteIntegrationRoleUsers
// 解绑员工与对应角色的关系，如需绑定请使用 CreateIntegrationUserRoles 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DEPARTUSERID = "InvalidParameter.DepartUserId"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLEID = "InvalidParameter.RoleId"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteIntegrationRoleUsers(request *DeleteIntegrationRoleUsersRequest) (response *DeleteIntegrationRoleUsersResponse, err error) {
    return c.DeleteIntegrationRoleUsersWithContext(context.Background(), request)
}

// DeleteIntegrationRoleUsers
// 解绑员工与对应角色的关系，如需绑定请使用 CreateIntegrationUserRoles 接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_DEPARTUSERID = "InvalidParameter.DepartUserId"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_ROLEID = "InvalidParameter.RoleId"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DeleteIntegrationRoleUsersWithContext(ctx context.Context, request *DeleteIntegrationRoleUsersRequest) (response *DeleteIntegrationRoleUsersResponse, err error) {
    if request == nil {
        request = NewDeleteIntegrationRoleUsersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIntegrationRoleUsers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIntegrationRoleUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSealPoliciesRequest() (request *DeleteSealPoliciesRequest) {
    request = &DeleteSealPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DeleteSealPolicies")
    
    
    return
}

func NewDeleteSealPoliciesResponse() (response *DeleteSealPoliciesResponse) {
    response = &DeleteSealPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSealPolicies
// 本接口（DeleteSealPolicies）用于撤销企业员工持有的印章权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSealPolicies(request *DeleteSealPoliciesRequest) (response *DeleteSealPoliciesResponse, err error) {
    return c.DeleteSealPoliciesWithContext(context.Background(), request)
}

// DeleteSealPolicies
// 本接口（DeleteSealPolicies）用于撤销企业员工持有的印章权限
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSealPoliciesWithContext(ctx context.Context, request *DeleteSealPoliciesRequest) (response *DeleteSealPoliciesResponse, err error) {
    if request == nil {
        request = NewDeleteSealPoliciesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSealPolicies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSealPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillUsageDetailRequest() (request *DescribeBillUsageDetailRequest) {
    request = &DescribeBillUsageDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeBillUsageDetail")
    
    
    return
}

func NewDescribeBillUsageDetailResponse() (response *DescribeBillUsageDetailResponse) {
    response = &DescribeBillUsageDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillUsageDetail
// 通过此接口（DescribeBillUsageDetail）查询该企业的套餐消耗详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDQUOTATYPE = "InvalidParameterValue.InvalidQuotaType"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  RESOURCENOTFOUND_NOTEXISTAPPLICATION = "ResourceNotFound.NotExistApplication"
func (c *Client) DescribeBillUsageDetail(request *DescribeBillUsageDetailRequest) (response *DescribeBillUsageDetailResponse, err error) {
    return c.DescribeBillUsageDetailWithContext(context.Background(), request)
}

// DescribeBillUsageDetail
// 通过此接口（DescribeBillUsageDetail）查询该企业的套餐消耗详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETERVALUE_INVALIDQUOTATYPE = "InvalidParameterValue.InvalidQuotaType"
//  INVALIDPARAMETERVALUE_INVALIDTIME = "InvalidParameterValue.InvalidTime"
//  RESOURCENOTFOUND_NOTEXISTAPPLICATION = "ResourceNotFound.NotExistApplication"
func (c *Client) DescribeBillUsageDetailWithContext(ctx context.Context, request *DescribeBillUsageDetailRequest) (response *DescribeBillUsageDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillUsageDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillUsageDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillUsageDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCancelFlowsTaskRequest() (request *DescribeCancelFlowsTaskRequest) {
    request = &DescribeCancelFlowsTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeCancelFlowsTask")
    
    
    return
}

func NewDescribeCancelFlowsTaskResponse() (response *DescribeCancelFlowsTaskResponse) {
    response = &DescribeCancelFlowsTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCancelFlowsTask
// 通过[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl)发起批量撤销任务后，可通过此接口查询批量撤销任务的结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BATCHCANCELFORBID = "OperationDenied.BatchCancelForbid"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeCancelFlowsTask(request *DescribeCancelFlowsTaskRequest) (response *DescribeCancelFlowsTaskResponse, err error) {
    return c.DescribeCancelFlowsTaskWithContext(context.Background(), request)
}

// DescribeCancelFlowsTask
// 通过[获取批量撤销签署流程腾讯电子签小程序链接](https://qian.tencent.com/developers/companyApis/operateFlows/CreateBatchCancelFlowUrl)发起批量撤销任务后，可通过此接口查询批量撤销任务的结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BATCHCANCELFORBID = "OperationDenied.BatchCancelForbid"
//  RESOURCENOTFOUND_ERRNOTEXISTROLE = "ResourceNotFound.ErrNotExistRole"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeCancelFlowsTaskWithContext(ctx context.Context, request *DescribeCancelFlowsTaskRequest) (response *DescribeCancelFlowsTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCancelFlowsTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCancelFlowsTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCancelFlowsTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtendedServiceAuthDetailRequest() (request *DescribeExtendedServiceAuthDetailRequest) {
    request = &DescribeExtendedServiceAuthDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeExtendedServiceAuthDetail")
    
    
    return
}

func NewDescribeExtendedServiceAuthDetailResponse() (response *DescribeExtendedServiceAuthDetailResponse) {
    response = &DescribeExtendedServiceAuthDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExtendedServiceAuthDetail
// 查询企业扩展服务的授权详情（列表），当前支持查询以下内容：
//
// 1. 企业自动签（本企业授权、集团企业授权、合作企业授权）
//
// 2. 批量签署能力
//
// 
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExtendedServiceAuthDetail(request *DescribeExtendedServiceAuthDetailRequest) (response *DescribeExtendedServiceAuthDetailResponse, err error) {
    return c.DescribeExtendedServiceAuthDetailWithContext(context.Background(), request)
}

// DescribeExtendedServiceAuthDetail
// 查询企业扩展服务的授权详情（列表），当前支持查询以下内容：
//
// 1. 企业自动签（本企业授权、集团企业授权、合作企业授权）
//
// 2. 批量签署能力
//
// 
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Agent.ProxyOperator.OpenId 需要传递超管或者法人的OpenId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeExtendedServiceAuthDetailWithContext(ctx context.Context, request *DescribeExtendedServiceAuthDetailRequest) (response *DescribeExtendedServiceAuthDetailResponse, err error) {
    if request == nil {
        request = NewDescribeExtendedServiceAuthDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtendedServiceAuthDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtendedServiceAuthDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtendedServiceAuthInfosRequest() (request *DescribeExtendedServiceAuthInfosRequest) {
    request = &DescribeExtendedServiceAuthInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeExtendedServiceAuthInfos")
    
    
    return
}

func NewDescribeExtendedServiceAuthInfosResponse() (response *DescribeExtendedServiceAuthInfosResponse) {
    response = &DescribeExtendedServiceAuthInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeExtendedServiceAuthInfos
// 查询企业扩展服务的开通和授权情况，当前支持查询以下内容：
//
// 
//
// 1. **企业自动签署**
//
// 2. **批量签署授权**
//
// 3. **企业与港澳台居民签署合同**
//
// 4. **拓宽签署方年龄限制**
//
// 5. **个人签署方仅校验手机号**
//
// 6. **隐藏合同经办人姓名**
//
// 7. **正楷临摹签名失败后更换其他签名类型**
//
// 8. **短信通知签署方**
//
// 9. **个人签署方手动签字**
//
// 10. **骑缝章**
//
// 11. **签署密码开通引导**
//
// 
//
// 
//
// 对应能力开通页面在Web控制台-更多-企业设置-拓展服务，如下图所示:
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7d79746ecca1c5fe878a2ec36ed69c23.jpg)
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Operator.UserId需要传递超管或者法人的UserId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthInfos(request *DescribeExtendedServiceAuthInfosRequest) (response *DescribeExtendedServiceAuthInfosResponse, err error) {
    return c.DescribeExtendedServiceAuthInfosWithContext(context.Background(), request)
}

// DescribeExtendedServiceAuthInfos
// 查询企业扩展服务的开通和授权情况，当前支持查询以下内容：
//
// 
//
// 1. **企业自动签署**
//
// 2. **批量签署授权**
//
// 3. **企业与港澳台居民签署合同**
//
// 4. **拓宽签署方年龄限制**
//
// 5. **个人签署方仅校验手机号**
//
// 6. **隐藏合同经办人姓名**
//
// 7. **正楷临摹签名失败后更换其他签名类型**
//
// 8. **短信通知签署方**
//
// 9. **个人签署方手动签字**
//
// 10. **骑缝章**
//
// 11. **签署密码开通引导**
//
// 
//
// 
//
// 对应能力开通页面在Web控制台-更多-企业设置-拓展服务，如下图所示:
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/7d79746ecca1c5fe878a2ec36ed69c23.jpg)
//
// 
//
// 注: <font color='red'>所在企业的超管、法人才有权限调用此接口</font>(Operator.UserId需要传递超管或者法人的UserId)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeExtendedServiceAuthInfosWithContext(ctx context.Context, request *DescribeExtendedServiceAuthInfosRequest) (response *DescribeExtendedServiceAuthInfosResponse, err error) {
    if request == nil {
        request = NewDescribeExtendedServiceAuthInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeExtendedServiceAuthInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeExtendedServiceAuthInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileUrlsRequest() (request *DescribeFileUrlsRequest) {
    request = &DescribeFileUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFileUrls")
    
    
    return
}

func NewDescribeFileUrlsResponse() (response *DescribeFileUrlsResponse) {
    response = &DescribeFileUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFileUrls
// 本接口（DescribeFileUrls）用于查询文件的下载URL。
//
// 适用场景：通过传参合同流程编号，下载对应的合同PDF文件流到本地。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSID = "InvalidParameter.BusinessId"
//  INVALIDPARAMETER_BUSINESSTYPE = "InvalidParameter.BusinessType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFileUrls(request *DescribeFileUrlsRequest) (response *DescribeFileUrlsResponse, err error) {
    return c.DescribeFileUrlsWithContext(context.Background(), request)
}

// DescribeFileUrls
// 本接口（DescribeFileUrls）用于查询文件的下载URL。
//
// 适用场景：通过传参合同流程编号，下载对应的合同PDF文件流到本地。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BUSINESSID = "InvalidParameter.BusinessId"
//  INVALIDPARAMETER_BUSINESSTYPE = "InvalidParameter.BusinessType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FILEDELETED = "OperationDenied.FileDeleted"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFileUrlsWithContext(ctx context.Context, request *DescribeFileUrlsRequest) (response *DescribeFileUrlsResponse, err error) {
    if request == nil {
        request = NewDescribeFileUrlsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileUrls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowBriefsRequest() (request *DescribeFlowBriefsRequest) {
    request = &DescribeFlowBriefsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowBriefs")
    
    
    return
}

func NewDescribeFlowBriefsResponse() (response *DescribeFlowBriefsResponse) {
    response = &DescribeFlowBriefsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlowBriefs
// 查询流程基础信息
//
// 适用场景：可用于主动查询某个合同流程的签署状态信息。可以配合回调通知使用。
//
// 
//
// 注: `每个企业限制日调用量限制：100W，当日超过此限制后再调用接口返回错误`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_INVALIDROLEID = "InvalidParameter.InvalidRoleId"
//  INVALIDPARAMETER_INVALIDROLENAME = "InvalidParameter.InvalidRoleName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowBriefs(request *DescribeFlowBriefsRequest) (response *DescribeFlowBriefsResponse, err error) {
    return c.DescribeFlowBriefsWithContext(context.Background(), request)
}

// DescribeFlowBriefs
// 查询流程基础信息
//
// 适用场景：可用于主动查询某个合同流程的签署状态信息。可以配合回调通知使用。
//
// 
//
// 注: `每个企业限制日调用量限制：100W，当日超过此限制后再调用接口返回错误`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CACHE = "InternalError.Cache"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_INVALIDROLEID = "InvalidParameter.InvalidRoleId"
//  INVALIDPARAMETER_INVALIDROLENAME = "InvalidParameter.InvalidRoleName"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_OUTQUERYLIMIT = "OperationDenied.OutQueryLimit"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFlowBriefsWithContext(ctx context.Context, request *DescribeFlowBriefsRequest) (response *DescribeFlowBriefsResponse, err error) {
    if request == nil {
        request = NewDescribeFlowBriefsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowBriefs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowBriefsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowComponentsRequest() (request *DescribeFlowComponentsRequest) {
    request = &DescribeFlowComponentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowComponents")
    
    
    return
}

func NewDescribeFlowComponentsResponse() (response *DescribeFlowComponentsResponse) {
    response = &DescribeFlowComponentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlowComponents
// 查询流程填写控件内容，可以根据合同流程ID查询该合同流程相关联的填写控件信息和填写内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
func (c *Client) DescribeFlowComponents(request *DescribeFlowComponentsRequest) (response *DescribeFlowComponentsResponse, err error) {
    return c.DescribeFlowComponentsWithContext(context.Background(), request)
}

// DescribeFlowComponents
// 查询流程填写控件内容，可以根据合同流程ID查询该合同流程相关联的填写控件信息和填写内容。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
func (c *Client) DescribeFlowComponentsWithContext(ctx context.Context, request *DescribeFlowComponentsRequest) (response *DescribeFlowComponentsResponse, err error) {
    if request == nil {
        request = NewDescribeFlowComponentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowComponents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowComponentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowEvidenceReportRequest() (request *DescribeFlowEvidenceReportRequest) {
    request = &DescribeFlowEvidenceReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowEvidenceReport")
    
    
    return
}

func NewDescribeFlowEvidenceReportResponse() (response *DescribeFlowEvidenceReportResponse) {
    response = &DescribeFlowEvidenceReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlowEvidenceReport
// 获取出证报告任务执行结果，返回报告 URL。
//
// 
//
// 
//
// 
//
// 注意：
//
// <ul><li>使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。</li>
//
// <li>需调用创建并返回出证报告接口<a href="https://qian.tencent.com/developers/companyApis/certificate/CreateFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>获取报告编号后调用当前接口获取报告链接。</li></ul>
//
// 
//
// <svg id="SvgjsSvg1006" width="262" height="229" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.com/svgjs"><defs id="SvgjsDefs1007"><pattern patternUnits="userSpaceOnUse" id="pattern_mark_0" width="300" height="300"><text x="150" y="100" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><pattern patternUnits="userSpaceOnUse" id="pattern_mark_1" width="300" height="300"><text x="150" y="200" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><marker id="SvgjsMarker1021" markerWidth="12" markerHeight="8" refX="9" refY="4" viewBox="0 0 12 8" orient="auto" markerUnits="userSpaceOnUse" stroke-dasharray="0,0"><path id="SvgjsPath1022" d="M0,0 L12,4 L0,8 L0,0" fill="#323232" stroke="#323232" stroke-width="1"></path></marker></defs><rect id="svgbackgroundid" width="262" height="229" fill="transparent"></rect><rect id="SvgjsRect1009" width="262" height="229" fill="url(#pattern_mark_0)"></rect><rect id="SvgjsRect1010" width="262" height="229" fill="url(#pattern_mark_1)"></rect><g id="SvgjsG1011" transform="translate(31.75,25)"><path id="SvgjsPath1012" d="M 0 0L 198 0L 198 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1013"><text id="SvgjsText1014" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="178px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1015" dy="16" x="99"><tspan id="SvgjsTspan1016" style="text-decoration:;fill: rgb(28, 30, 33);">CreateFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1017" dy="16" x="99"><tspan id="SvgjsTspan1018" style="text-decoration:;fill: rgb(51, 51, 51);">提交申请出证报告任务</tspan></tspan></text></g></g><g id="SvgjsG1019"><path id="SvgjsPath1020" d="M130.75 84.5L130.75 114.5L130.75 114.5L130.75 143.2" stroke="#323232" stroke-width="1" fill="none" marker-end="url(#SvgjsMarker1021)"></path></g><g id="SvgjsG1023" transform="translate(25,145)"><path id="SvgjsPath1024" d="M 0 0L 211.5 0L 211.5 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1025"><text id="SvgjsText1026" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="192px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1027" dy="16" x="106"><tspan id="SvgjsTspan1028" style="text-decoration:;fill: rgb(28, 30, 33);">DescribeFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1029" dy="16" x="106"><tspan id="SvgjsTspan1030" style="text-decoration:;fill: rgb(51, 51, 51);">获取出证报告任务执行结果</tspan></tspan></text></g></g></svg>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowEvidenceReport(request *DescribeFlowEvidenceReportRequest) (response *DescribeFlowEvidenceReportResponse, err error) {
    return c.DescribeFlowEvidenceReportWithContext(context.Background(), request)
}

// DescribeFlowEvidenceReport
// 获取出证报告任务执行结果，返回报告 URL。
//
// 
//
// 
//
// 
//
// 注意：
//
// <ul><li>使用此功能`需搭配出证套餐` ，使用前请联系对接的客户经理沟通。</li>
//
// <li>需调用创建并返回出证报告接口<a href="https://qian.tencent.com/developers/companyApis/certificate/CreateFlowEvidenceReport" target="_blank">提交申请出证报告任务</a>获取报告编号后调用当前接口获取报告链接。</li></ul>
//
// 
//
// <svg id="SvgjsSvg1006" width="262" height="229" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.com/svgjs"><defs id="SvgjsDefs1007"><pattern patternUnits="userSpaceOnUse" id="pattern_mark_0" width="300" height="300"><text x="150" y="100" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><pattern patternUnits="userSpaceOnUse" id="pattern_mark_1" width="300" height="300"><text x="150" y="200" fill="rgba(229,229,229,0.8)" font-size="18" transform="rotate(-45, 150, 150)" style="dominant-baseline: middle; text-anchor: middle;"></text></pattern><marker id="SvgjsMarker1021" markerWidth="12" markerHeight="8" refX="9" refY="4" viewBox="0 0 12 8" orient="auto" markerUnits="userSpaceOnUse" stroke-dasharray="0,0"><path id="SvgjsPath1022" d="M0,0 L12,4 L0,8 L0,0" fill="#323232" stroke="#323232" stroke-width="1"></path></marker></defs><rect id="svgbackgroundid" width="262" height="229" fill="transparent"></rect><rect id="SvgjsRect1009" width="262" height="229" fill="url(#pattern_mark_0)"></rect><rect id="SvgjsRect1010" width="262" height="229" fill="url(#pattern_mark_1)"></rect><g id="SvgjsG1011" transform="translate(31.75,25)"><path id="SvgjsPath1012" d="M 0 0L 198 0L 198 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1013"><text id="SvgjsText1014" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="178px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1015" dy="16" x="99"><tspan id="SvgjsTspan1016" style="text-decoration:;fill: rgb(28, 30, 33);">CreateFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1017" dy="16" x="99"><tspan id="SvgjsTspan1018" style="text-decoration:;fill: rgb(51, 51, 51);">提交申请出证报告任务</tspan></tspan></text></g></g><g id="SvgjsG1019"><path id="SvgjsPath1020" d="M130.75 84.5L130.75 114.5L130.75 114.5L130.75 143.2" stroke="#323232" stroke-width="1" fill="none" marker-end="url(#SvgjsMarker1021)"></path></g><g id="SvgjsG1023" transform="translate(25,145)"><path id="SvgjsPath1024" d="M 0 0L 211.5 0L 211.5 59L 0 59Z" stroke="rgba(86,146,48,1)" stroke-width="1" fill-opacity="1" fill="#e7ebed"></path><g id="SvgjsG1025"><text id="SvgjsText1026" font-family="微软雅黑" text-anchor="middle" font-size="13px" width="192px" fill="#323232" font-weight="400" align="middle" lineHeight="125%" anchor="middle" family="微软雅黑" size="13px" weight="400" font-style="" opacity="1" y="10.375" transform="rotate(0)"><tspan id="SvgjsTspan1027" dy="16" x="106"><tspan id="SvgjsTspan1028" style="text-decoration:;fill: rgb(28, 30, 33);">DescribeFlowEvidenceReport</tspan></tspan><tspan id="SvgjsTspan1029" dy="16" x="106"><tspan id="SvgjsTspan1030" style="text-decoration:;fill: rgb(51, 51, 51);">获取出证报告任务执行结果</tspan></tspan></text></g></g></svg>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER_FLOWID = "MissingParameter.FlowId"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_NOTBELONGSUPERADMINORLEGALPERSON = "OperationDenied.NotBelongSuperAdminOrLegalPerson"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_URL = "ResourceNotFound.Url"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowEvidenceReportWithContext(ctx context.Context, request *DescribeFlowEvidenceReportRequest) (response *DescribeFlowEvidenceReportResponse, err error) {
    if request == nil {
        request = NewDescribeFlowEvidenceReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowEvidenceReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowEvidenceReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowInfoRequest() (request *DescribeFlowInfoRequest) {
    request = &DescribeFlowInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowInfo")
    
    
    return
}

func NewDescribeFlowInfoResponse() (response *DescribeFlowInfoResponse) {
    response = &DescribeFlowInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlowInfo
// 此接口用于查询合同流程的详情信息，支持查询多个（数量不能超过100）。
//
// 
//
// 适用场景：可用于主动查询某个合同详情信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowInfo(request *DescribeFlowInfoRequest) (response *DescribeFlowInfoResponse, err error) {
    return c.DescribeFlowInfoWithContext(context.Background(), request)
}

// DescribeFlowInfo
// 此接口用于查询合同流程的详情信息，支持查询多个（数量不能超过100）。
//
// 
//
// 适用场景：可用于主动查询某个合同详情信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_ORGANIZATIONNOTACTIVATED = "OperationDenied.OrganizationNotActivated"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeFlowInfoWithContext(ctx context.Context, request *DescribeFlowInfoRequest) (response *DescribeFlowInfoResponse, err error) {
    if request == nil {
        request = NewDescribeFlowInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlowTemplatesRequest() (request *DescribeFlowTemplatesRequest) {
    request = &DescribeFlowTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeFlowTemplates")
    
    
    return
}

func NewDescribeFlowTemplatesResponse() (response *DescribeFlowTemplatesResponse) {
    response = &DescribeFlowTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeFlowTemplates
// 此接口（DescribeFlowTemplates）用于查询本企业模板列表信息。
//
// 
//
// 
//
// **适用场景**
//
// 该接口常用来配合<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument" target="_blank">模板发起合同-创建电子文档</a>接口，作为创建电子文档的前置接口使用。
//
// 通过此接口查询到模板信息后，再通过调用创建电子文档接口，指定模板ID，指定模板中需要的填写控件内容等，完成电子文档的创建。
//
// 
//
// **一个模板通常会包含以下结构信息** 
//
// 
//
// - 模板模板ID, 模板名字等基本信息
//
// - 发起方参与信息Promoter、签署参与方 Recipients，后者会在模板发起合同时用于指定参与方
//
// - 发起方和签署方的填写控件 Components
//
// - 签署方的签署控件 SignComponents
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/ab81fa948a0a6fea14f48cac91d0e36a/channel_DescribeTemplates.png)
//
// 
//
// 模板中各元素的层级关系, 所有的填写控件和签署控件都归属某一个角色(通过控件的ComponentRecipientId来关联)
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/45c638bd93f9c8024763add9ab47c27f.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUNDSHADOWUSER = "FailedOperation.NotFoundShadowUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFlowTemplates(request *DescribeFlowTemplatesRequest) (response *DescribeFlowTemplatesResponse, err error) {
    return c.DescribeFlowTemplatesWithContext(context.Background(), request)
}

// DescribeFlowTemplates
// 此接口（DescribeFlowTemplates）用于查询本企业模板列表信息。
//
// 
//
// 
//
// **适用场景**
//
// 该接口常用来配合<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument" target="_blank">模板发起合同-创建电子文档</a>接口，作为创建电子文档的前置接口使用。
//
// 通过此接口查询到模板信息后，再通过调用创建电子文档接口，指定模板ID，指定模板中需要的填写控件内容等，完成电子文档的创建。
//
// 
//
// **一个模板通常会包含以下结构信息** 
//
// 
//
// - 模板模板ID, 模板名字等基本信息
//
// - 发起方参与信息Promoter、签署参与方 Recipients，后者会在模板发起合同时用于指定参与方
//
// - 发起方和签署方的填写控件 Components
//
// - 签署方的签署控件 SignComponents
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/ab81fa948a0a6fea14f48cac91d0e36a/channel_DescribeTemplates.png)
//
// 
//
// 模板中各元素的层级关系, 所有的填写控件和签署控件都归属某一个角色(通过控件的ComponentRecipientId来关联)
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/45c638bd93f9c8024763add9ab47c27f.png)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTFOUNDSHADOWUSER = "FailedOperation.NotFoundShadowUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTENTTYPE = "InvalidParameter.ContentType"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCENOTFOUND_APPLICATION = "ResourceNotFound.Application"
//  RESOURCENOTFOUND_TEMPLATE = "ResourceNotFound.Template"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeFlowTemplatesWithContext(ctx context.Context, request *DescribeFlowTemplatesRequest) (response *DescribeFlowTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeFlowTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlowTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlowTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationDepartmentsRequest() (request *DescribeIntegrationDepartmentsRequest) {
    request = &DescribeIntegrationDepartmentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeIntegrationDepartments")
    
    
    return
}

func NewDescribeIntegrationDepartmentsResponse() (response *DescribeIntegrationDepartmentsResponse) {
    response = &DescribeIntegrationDepartmentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationDepartments
// 此接口（DescribeIntegrationDepartments）用于查询企业的部门信息列表，支持查询单个部门节点或单个部门节点及一级子节点部门列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeIntegrationDepartments(request *DescribeIntegrationDepartmentsRequest) (response *DescribeIntegrationDepartmentsResponse, err error) {
    return c.DescribeIntegrationDepartmentsWithContext(context.Background(), request)
}

// DescribeIntegrationDepartments
// 此接口（DescribeIntegrationDepartments）用于查询企业的部门信息列表，支持查询单个部门节点或单个部门节点及一级子节点部门列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeIntegrationDepartmentsWithContext(ctx context.Context, request *DescribeIntegrationDepartmentsRequest) (response *DescribeIntegrationDepartmentsResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationDepartmentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationDepartments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationDepartmentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationEmployeesRequest() (request *DescribeIntegrationEmployeesRequest) {
    request = &DescribeIntegrationEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeIntegrationEmployees")
    
    
    return
}

func NewDescribeIntegrationEmployeesResponse() (response *DescribeIntegrationEmployeesResponse) {
    response = &DescribeIntegrationEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationEmployees
// 此接口（DescribeIntegrationEmployees）用于分页查询企业员工信息列表，支持设置过滤条件以筛选员工查询结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeIntegrationEmployees(request *DescribeIntegrationEmployeesRequest) (response *DescribeIntegrationEmployeesResponse, err error) {
    return c.DescribeIntegrationEmployeesWithContext(context.Background(), request)
}

// DescribeIntegrationEmployees
// 此接口（DescribeIntegrationEmployees）用于分页查询企业员工信息列表，支持设置过滤条件以筛选员工查询结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_INVALIDCHANNEL = "InvalidParameter.InvalidChannel"
//  INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"
//  INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"
//  INVALIDPARAMETER_INVALIDOPERATORID = "InvalidParameter.InvalidOperatorId"
//  INVALIDPARAMETER_INVALIDORGANIZATIONID = "InvalidParameter.InvalidOrganizationId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeIntegrationEmployeesWithContext(ctx context.Context, request *DescribeIntegrationEmployeesRequest) (response *DescribeIntegrationEmployeesResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIntegrationRolesRequest() (request *DescribeIntegrationRolesRequest) {
    request = &DescribeIntegrationRolesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeIntegrationRoles")
    
    
    return
}

func NewDescribeIntegrationRolesResponse() (response *DescribeIntegrationRolesResponse) {
    response = &DescribeIntegrationRolesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIntegrationRoles
// 此接口（DescribeIntegrationRoles）用于分页查询企业角色列表，列表按照角色创建时间升序排列。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeIntegrationRoles(request *DescribeIntegrationRolesRequest) (response *DescribeIntegrationRolesResponse, err error) {
    return c.DescribeIntegrationRolesWithContext(context.Background(), request)
}

// DescribeIntegrationRoles
// 此接口（DescribeIntegrationRoles）用于分页查询企业角色列表，列表按照角色创建时间升序排列。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_LIMIT = "InvalidParameter.Limit"
//  INVALIDPARAMETER_OFFSET = "InvalidParameter.Offset"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeIntegrationRolesWithContext(ctx context.Context, request *DescribeIntegrationRolesRequest) (response *DescribeIntegrationRolesResponse, err error) {
    if request == nil {
        request = NewDescribeIntegrationRolesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIntegrationRoles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIntegrationRolesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationGroupOrganizationsRequest() (request *DescribeOrganizationGroupOrganizationsRequest) {
    request = &DescribeOrganizationGroupOrganizationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeOrganizationGroupOrganizations")
    
    
    return
}

func NewDescribeOrganizationGroupOrganizationsResponse() (response *DescribeOrganizationGroupOrganizationsResponse) {
    response = &DescribeOrganizationGroupOrganizationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationGroupOrganizations
// 此API接口用来查询加入集团的成员企业信息
//
// 适用场景：子企业在加入集团后，主企业可能通过此接口获取到所有的子企业列表，方便进行展示和统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
func (c *Client) DescribeOrganizationGroupOrganizations(request *DescribeOrganizationGroupOrganizationsRequest) (response *DescribeOrganizationGroupOrganizationsResponse, err error) {
    return c.DescribeOrganizationGroupOrganizationsWithContext(context.Background(), request)
}

// DescribeOrganizationGroupOrganizations
// 此API接口用来查询加入集团的成员企业信息
//
// 适用场景：子企业在加入集团后，主企业可能通过此接口获取到所有的子企业列表，方便进行展示和统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_API = "InternalError.Api"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
func (c *Client) DescribeOrganizationGroupOrganizationsWithContext(ctx context.Context, request *DescribeOrganizationGroupOrganizationsRequest) (response *DescribeOrganizationGroupOrganizationsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationGroupOrganizationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationGroupOrganizations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationGroupOrganizationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrganizationSealsRequest() (request *DescribeOrganizationSealsRequest) {
    request = &DescribeOrganizationSealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeOrganizationSeals")
    
    
    return
}

func NewDescribeOrganizationSealsResponse() (response *DescribeOrganizationSealsResponse) {
    response = &DescribeOrganizationSealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOrganizationSeals
// 查询企业印章的列表，需要操作者具有查询印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeOrganizationSeals(request *DescribeOrganizationSealsRequest) (response *DescribeOrganizationSealsResponse, err error) {
    return c.DescribeOrganizationSealsWithContext(context.Background(), request)
}

// DescribeOrganizationSeals
// 查询企业印章的列表，需要操作者具有查询印章权限
//
// 客户指定需要获取的印章数量和偏移量，数量最多100，超过100按100处理；入参InfoType控制印章是否携带授权人信息，为1则携带，为0则返回的授权人信息为空数组。接口调用成功返回印章的信息列表还有企业印章的总数。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_OPERATORHASNOPERMISSION = "OperationDenied.OperatorHasNoPermission"
//  OPERATIONDENIED_SUBORGNOTJOIN = "OperationDenied.SubOrgNotJoin"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeOrganizationSealsWithContext(ctx context.Context, request *DescribeOrganizationSealsRequest) (response *DescribeOrganizationSealsResponse, err error) {
    if request == nil {
        request = NewDescribeOrganizationSealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOrganizationSeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOrganizationSealsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePersonCertificateRequest() (request *DescribePersonCertificateRequest) {
    request = &DescribePersonCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribePersonCertificate")
    
    
    return
}

func NewDescribePersonCertificateResponse() (response *DescribePersonCertificateResponse) {
    response = &DescribePersonCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePersonCertificate
// 此接口（DescribePersonCertificate）用于查询个人数字证书信息。<br />注：`1.目前仅用于查询开通了医疗自动签署功能的个人数字证书。`<br />`2.调用此接口需要开通白名单，使用前请联系相关人员开通白名单。`
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePersonCertificate(request *DescribePersonCertificateRequest) (response *DescribePersonCertificateResponse, err error) {
    return c.DescribePersonCertificateWithContext(context.Background(), request)
}

// DescribePersonCertificate
// 此接口（DescribePersonCertificate）用于查询个人数字证书信息。<br />注：`1.目前仅用于查询开通了医疗自动签署功能的个人数字证书。`<br />`2.调用此接口需要开通白名单，使用前请联系相关人员开通白名单。`
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_EMPTYPARAMS = "InvalidParameter.EmptyParams"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribePersonCertificateWithContext(ctx context.Context, request *DescribePersonCertificateRequest) (response *DescribePersonCertificateResponse, err error) {
    if request == nil {
        request = NewDescribePersonCertificateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePersonCertificate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePersonCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSignFaceVideoRequest() (request *DescribeSignFaceVideoRequest) {
    request = &DescribeSignFaceVideoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeSignFaceVideo")
    
    
    return
}

func NewDescribeSignFaceVideoResponse() (response *DescribeSignFaceVideoResponse) {
    response = &DescribeSignFaceVideoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSignFaceVideo
// 该接口用于在使用视频认证方式签署合同后，获取用户的签署人脸认证视频。
//
// 
//
// 1. 该接口**仅适用于在H5端签署**的合同，**在通过视频认证后**获取人脸图片。
//
// 2. 该接口**不支持小程序端**的签署人脸图片获取。
//
// 3. 请在**签署完成后的三天内**获取人脸图片，**过期后将无法获取**。
//
// 
//
// **注意：该接口需要开通白名单，请联系客户经理开通后使用。**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSignFaceVideo(request *DescribeSignFaceVideoRequest) (response *DescribeSignFaceVideoResponse, err error) {
    return c.DescribeSignFaceVideoWithContext(context.Background(), request)
}

// DescribeSignFaceVideo
// 该接口用于在使用视频认证方式签署合同后，获取用户的签署人脸认证视频。
//
// 
//
// 1. 该接口**仅适用于在H5端签署**的合同，**在通过视频认证后**获取人脸图片。
//
// 2. 该接口**不支持小程序端**的签署人脸图片获取。
//
// 3. 请在**签署完成后的三天内**获取人脸图片，**过期后将无法获取**。
//
// 
//
// **注意：该接口需要开通白名单，请联系客户经理开通后使用。**
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSignFaceVideoWithContext(ctx context.Context, request *DescribeSignFaceVideoRequest) (response *DescribeSignFaceVideoResponse, err error) {
    if request == nil {
        request = NewDescribeSignFaceVideoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSignFaceVideo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSignFaceVideoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeThirdPartyAuthCodeRequest() (request *DescribeThirdPartyAuthCodeRequest) {
    request = &DescribeThirdPartyAuthCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeThirdPartyAuthCode")
    
    
    return
}

func NewDescribeThirdPartyAuthCodeResponse() (response *DescribeThirdPartyAuthCodeResponse) {
    response = &DescribeThirdPartyAuthCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeThirdPartyAuthCode
// 通过AuthCode查询个人用户是否实名
//
// 
//
// 
//
// 注意: 
//
// <ul>
//
// <li>此接口为合作引流场景使用，使用`有白名单限制`，使用前请联系对接的客户经理沟通。</li>
//
// <li>`AuthCode 只能使用一次`，查询一次再次查询会返回错误</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_AUTHCODE = "MissingParameter.AuthCode"
//  OPERATIONDENIED_AUTHCODEINVALID = "OperationDenied.AuthCodeInvalid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeThirdPartyAuthCode(request *DescribeThirdPartyAuthCodeRequest) (response *DescribeThirdPartyAuthCodeResponse, err error) {
    return c.DescribeThirdPartyAuthCodeWithContext(context.Background(), request)
}

// DescribeThirdPartyAuthCode
// 通过AuthCode查询个人用户是否实名
//
// 
//
// 
//
// 注意: 
//
// <ul>
//
// <li>此接口为合作引流场景使用，使用`有白名单限制`，使用前请联系对接的客户经理沟通。</li>
//
// <li>`AuthCode 只能使用一次`，查询一次再次查询会返回错误</li>
//
// </ul>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER_AUTHCODE = "MissingParameter.AuthCode"
//  OPERATIONDENIED_AUTHCODEINVALID = "OperationDenied.AuthCodeInvalid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) DescribeThirdPartyAuthCodeWithContext(ctx context.Context, request *DescribeThirdPartyAuthCodeRequest) (response *DescribeThirdPartyAuthCodeResponse, err error) {
    if request == nil {
        request = NewDescribeThirdPartyAuthCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeThirdPartyAuthCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeThirdPartyAuthCodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserAutoSignStatusRequest() (request *DescribeUserAutoSignStatusRequest) {
    request = &DescribeUserAutoSignStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DescribeUserAutoSignStatus")
    
    
    return
}

func NewDescribeUserAutoSignStatusResponse() (response *DescribeUserAutoSignStatusResponse) {
    response = &DescribeUserAutoSignStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserAutoSignStatus
// 通过此接口获取个人用户自动签的开通状态。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeUserAutoSignStatus(request *DescribeUserAutoSignStatusRequest) (response *DescribeUserAutoSignStatusResponse, err error) {
    return c.DescribeUserAutoSignStatusWithContext(context.Background(), request)
}

// DescribeUserAutoSignStatus
// 通过此接口获取个人用户自动签的开通状态。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeUserAutoSignStatusWithContext(ctx context.Context, request *DescribeUserAutoSignStatusRequest) (response *DescribeUserAutoSignStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUserAutoSignStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserAutoSignStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserAutoSignStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDisableUserAutoSignRequest() (request *DisableUserAutoSignRequest) {
    request = &DisableUserAutoSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "DisableUserAutoSign")
    
    
    return
}

func NewDisableUserAutoSignResponse() (response *DisableUserAutoSignResponse) {
    response = &DisableUserAutoSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DisableUserAutoSign
// 通过此接口可以关闭个人用户自动签功能。
//
// 无需对应的用户刷脸等方式同意即可关闭。
//
// 
//
// 注意: 
//
// 
//
// <ul><li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>如果此用户在开通时候绑定过个人自动签账号许可,  关闭此用户的自动签不会归还个人自动签账号许可的额度。</li></ul>
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DisableUserAutoSign(request *DisableUserAutoSignRequest) (response *DisableUserAutoSignResponse, err error) {
    return c.DisableUserAutoSignWithContext(context.Background(), request)
}

// DisableUserAutoSign
// 通过此接口可以关闭个人用户自动签功能。
//
// 无需对应的用户刷脸等方式同意即可关闭。
//
// 
//
// 注意: 
//
// 
//
// <ul><li>处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。</li>
//
// <li>如果此用户在开通时候绑定过个人自动签账号许可,  关闭此用户的自动签不会归还个人自动签账号许可的额度。</li></ul>
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLOPENCLOUDAPIERROR = "InternalError.CallOpenCloudApiError"
//  INTERNALERROR_CALLBACK = "InternalError.Callback"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DisableUserAutoSignWithContext(ctx context.Context, request *DisableUserAutoSignRequest) (response *DisableUserAutoSignResponse, err error) {
    if request == nil {
        request = NewDisableUserAutoSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DisableUserAutoSign require credential")
    }

    request.SetContext(ctx)
    
    response = NewDisableUserAutoSignResponse()
    err = c.Send(request, response)
    return
}

func NewGetTaskResultApiRequest() (request *GetTaskResultApiRequest) {
    request = &GetTaskResultApiRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "GetTaskResultApi")
    
    
    return
}

func NewGetTaskResultApiResponse() (response *GetTaskResultApiResponse) {
    response = &GetTaskResultApiResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetTaskResultApi
// 此接口（GetTaskResultApi）用来查询转换任务的状态。如需发起转换任务，请使用<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行资源文件的转换操作<br />
//
// 前提条件：已调用 <a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行文件转换，并得到了返回的转换任务Id。<br />
//
// 
//
// 适用场景：已创建一个文件转换任务，想查询该文件转换任务的状态，或获取转换后的文件资源Id。<br />
//
// 注：
//
// 1. `大文件转换所需的时间可能会比较长`
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) GetTaskResultApi(request *GetTaskResultApiRequest) (response *GetTaskResultApiResponse, err error) {
    return c.GetTaskResultApiWithContext(context.Background(), request)
}

// GetTaskResultApi
// 此接口（GetTaskResultApi）用来查询转换任务的状态。如需发起转换任务，请使用<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行资源文件的转换操作<br />
//
// 前提条件：已调用 <a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务接口</a>进行文件转换，并得到了返回的转换任务Id。<br />
//
// 
//
// 适用场景：已创建一个文件转换任务，想查询该文件转换任务的状态，或获取转换后的文件资源Id。<br />
//
// 注：
//
// 1. `大文件转换所需的时间可能会比较长`
//
// 可能返回的错误码:
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_RESOURCETYPE = "InvalidParameter.ResourceType"
//  MISSINGPARAMETER_RESOURCENAME = "MissingParameter.ResourceName"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) GetTaskResultApiWithContext(ctx context.Context, request *GetTaskResultApiRequest) (response *GetTaskResultApiResponse, err error) {
    if request == nil {
        request = NewGetTaskResultApiRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetTaskResultApi require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetTaskResultApiResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationCallbackInfoRequest() (request *ModifyApplicationCallbackInfoRequest) {
    request = &ModifyApplicationCallbackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "ModifyApplicationCallbackInfo")
    
    
    return
}

func NewModifyApplicationCallbackInfoResponse() (response *ModifyApplicationCallbackInfoResponse) {
    response = &ModifyApplicationCallbackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplicationCallbackInfo
// 新增/删除企业应用集成中的回调配置。
//
// 新增回调配置只会增加不存在的CallbackUrl；删除操作将针对找到的相同CallbackUrl的配置进行删除。
//
// 请确保回调地址能够接收并处理 HTTP POST 请求，并返回状态码 200 以表示处理正常。
//
// 更多回调相关的说明参考文档[回调通知能力](https://qian.tencent.com/developers/company/callback_types_v2)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyApplicationCallbackInfo(request *ModifyApplicationCallbackInfoRequest) (response *ModifyApplicationCallbackInfoResponse, err error) {
    return c.ModifyApplicationCallbackInfoWithContext(context.Background(), request)
}

// ModifyApplicationCallbackInfo
// 新增/删除企业应用集成中的回调配置。
//
// 新增回调配置只会增加不存在的CallbackUrl；删除操作将针对找到的相同CallbackUrl的配置进行删除。
//
// 请确保回调地址能够接收并处理 HTTP POST 请求，并返回状态码 200 以表示处理正常。
//
// 更多回调相关的说明参考文档[回调通知能力](https://qian.tencent.com/developers/company/callback_types_v2)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyApplicationCallbackInfoWithContext(ctx context.Context, request *ModifyApplicationCallbackInfoRequest) (response *ModifyApplicationCallbackInfoResponse, err error) {
    if request == nil {
        request = NewModifyApplicationCallbackInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplicationCallbackInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationCallbackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyExtendedServiceRequest() (request *ModifyExtendedServiceRequest) {
    request = &ModifyExtendedServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "ModifyExtendedService")
    
    
    return
}

func NewModifyExtendedServiceResponse() (response *ModifyExtendedServiceResponse) {
    response = &ModifyExtendedServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyExtendedService
// 管理企业扩展服务 ，企业经办人需要是企业超管或者法人。
//
// 
//
// 跳转小程序的几种方式：主要是设置不同的EndPoint
//
// 1. 通过链接Url直接跳转到小程序，不需要返回
//
// 设置EndPoint为WEIXINAPP，得到链接打开即可。
//
// 
//
// 2. 客户App直接跳转到小程序-->腾讯电子签小程序操作完成-->返回App
//
// 跳转到小程序的实现，参考官方文档<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html">打开 App</a>
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 3. 客户小程序直接跳到电子签小程序-->腾讯电子签小程序操作完成--->回到客户小程序
//
// 跳转到小程序的实现，参考官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 4.其中小程序的原始Id如下，或者查看小程序信息自助获取。
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyExtendedService(request *ModifyExtendedServiceRequest) (response *ModifyExtendedServiceResponse, err error) {
    return c.ModifyExtendedServiceWithContext(context.Background(), request)
}

// ModifyExtendedService
// 管理企业扩展服务 ，企业经办人需要是企业超管或者法人。
//
// 
//
// 跳转小程序的几种方式：主要是设置不同的EndPoint
//
// 1. 通过链接Url直接跳转到小程序，不需要返回
//
// 设置EndPoint为WEIXINAPP，得到链接打开即可。
//
// 
//
// 2. 客户App直接跳转到小程序-->腾讯电子签小程序操作完成-->返回App
//
// 跳转到小程序的实现，参考官方文档<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/launchApp.html">打开 App</a>
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 3. 客户小程序直接跳到电子签小程序-->腾讯电子签小程序操作完成--->回到客户小程序
//
// 跳转到小程序的实现，参考官方文档（分为<a href="https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.navigateToMiniProgram.html">全屏</a>、<a href="https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html">半屏</a>两种方式），如何配置也可以请参考: <a href="https://qian.tencent.com/developers/company/openwxminiprogram">跳转电子签小程序配置</a>
//
// 设置EndPoint为APP，得到path。
//
// 
//
// 4.其中小程序的原始Id如下，或者查看小程序信息自助获取。
//
// 
//
// | 小程序 | AppID | 原始ID |
//
// | ------------ | ------------ | ------------ |
//
// | 腾讯电子签（正式版） | wxa023b292fd19d41d | gh_da88f6188665 |
//
// | 腾讯电子签Demo | wx371151823f6f3edf | gh_39a5d3de69fa |
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyExtendedServiceWithContext(ctx context.Context, request *ModifyExtendedServiceRequest) (response *ModifyExtendedServiceResponse, err error) {
    if request == nil {
        request = NewModifyExtendedServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyExtendedService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyExtendedServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFlowDeadlineRequest() (request *ModifyFlowDeadlineRequest) {
    request = &ModifyFlowDeadlineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "ModifyFlowDeadline")
    
    
    return
}

func NewModifyFlowDeadlineResponse() (response *ModifyFlowDeadlineResponse) {
    response = &ModifyFlowDeadlineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyFlowDeadline
// 在已发起的签署流程中，支持对签署截止日期进行延期操作，主要涉及以下两个维度：
//
// 1. 合同（流程）维度：只需要传递签署流程ID。这将对签署流程和发起时未单独设定签署截止时间的签署人进行延期操作。
//
// 2. 签署人维度：需要传递流程ID和签署人ID。此操作将对签署人进行延期操作，尤其对于有序的合同（流程），签署截止时间不能超过后一位合同（流程）签署人的流程截止时间。
//
// 
//
// 此接口有如下限制条件：
//
// 1. 执行操作的员工需为发起方企业的超级管理员、法定代表人或签署流程的发起人。
//
// 2. 在延长整个签署流程时，签署流程应至少还有一方未签署（即签署流程不能处于已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态）。
//
// 3. 在延长整个签署流程时，新的签署截止日期应晚于合同已设定的签署截止日期和当前日期。
//
// 4. 在延长签署方的截止时间时，签署方不能处于流程完结或已终止状态（即签署人不能处于已签署、已拒签、已过期、已撤回、拒绝填写、已解除等状态）。
//
// 5. 在延长签署方的截止时间时，签署方的新签署截止日期应晚于当前日期和签署方已设定的截止日期。若为有序合同，还应早于或等于下一签署人的截止日期，且早于签署流程整体的截止日期。
//
// 6. 不支持操作合同组合同。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyFlowDeadline(request *ModifyFlowDeadlineRequest) (response *ModifyFlowDeadlineResponse, err error) {
    return c.ModifyFlowDeadlineWithContext(context.Background(), request)
}

// ModifyFlowDeadline
// 在已发起的签署流程中，支持对签署截止日期进行延期操作，主要涉及以下两个维度：
//
// 1. 合同（流程）维度：只需要传递签署流程ID。这将对签署流程和发起时未单独设定签署截止时间的签署人进行延期操作。
//
// 2. 签署人维度：需要传递流程ID和签署人ID。此操作将对签署人进行延期操作，尤其对于有序的合同（流程），签署截止时间不能超过后一位合同（流程）签署人的流程截止时间。
//
// 
//
// 此接口有如下限制条件：
//
// 1. 执行操作的员工需为发起方企业的超级管理员、法定代表人或签署流程的发起人。
//
// 2. 在延长整个签署流程时，签署流程应至少还有一方未签署（即签署流程不能处于已全部签署完成、已拒签、已过期、已撤回、拒绝填写、已解除等合同状态）。
//
// 3. 在延长整个签署流程时，新的签署截止日期应晚于合同已设定的签署截止日期和当前日期。
//
// 4. 在延长签署方的截止时间时，签署方不能处于流程完结或已终止状态（即签署人不能处于已签署、已拒签、已过期、已撤回、拒绝填写、已解除等状态）。
//
// 5. 在延长签署方的截止时间时，签署方的新签署截止日期应晚于当前日期和签署方已设定的截止日期。若为有序合同，还应早于或等于下一签署人的截止日期，且早于签署流程整体的截止日期。
//
// 6. 不支持操作合同组合同。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENOTACHIEVENORMALLEGAL = "FailedOperation.AgeNotAchieveNormalLegal"
//  FAILEDOPERATION_NOTAVAILABLESIGNREVIEW = "FailedOperation.NotAvailableSignReview"
//  FAILEDOPERATION_ORGANIZATIONEXPERIENCECHANGE = "FailedOperation.OrganizationExperienceChange"
//  FAILEDOPERATION_ORGANIZATIONNAMECHANGED = "FailedOperation.OrganizationNameChanged"
//  FAILEDOPERATION_ORGANIZATIONNAMENEEDCHANGE = "FailedOperation.OrganizationNameNeedChange"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_DECRYPTION = "InternalError.Decryption"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyFlowDeadlineWithContext(ctx context.Context, request *ModifyFlowDeadlineRequest) (response *ModifyFlowDeadlineResponse, err error) {
    if request == nil {
        request = NewModifyFlowDeadlineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyFlowDeadline require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyFlowDeadlineResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIntegrationDepartmentRequest() (request *ModifyIntegrationDepartmentRequest) {
    request = &ModifyIntegrationDepartmentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "ModifyIntegrationDepartment")
    
    
    return
}

func NewModifyIntegrationDepartmentResponse() (response *ModifyIntegrationDepartmentResponse) {
    response = &ModifyIntegrationDepartmentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIntegrationDepartment
// 此接口（ModifyIntegrationDepartment）用于更新企业的部门信息，支持更新部门名称、客户系统部门ID和部门序号等信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyIntegrationDepartment(request *ModifyIntegrationDepartmentRequest) (response *ModifyIntegrationDepartmentResponse, err error) {
    return c.ModifyIntegrationDepartmentWithContext(context.Background(), request)
}

// ModifyIntegrationDepartment
// 此接口（ModifyIntegrationDepartment）用于更新企业的部门信息，支持更新部门名称、客户系统部门ID和部门序号等信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_DATAEXISTS = "InvalidParameter.DataExists"
//  INVALIDPARAMETER_DATANOTFOUND = "InvalidParameter.DataNotFound"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) ModifyIntegrationDepartmentWithContext(ctx context.Context, request *ModifyIntegrationDepartmentRequest) (response *ModifyIntegrationDepartmentResponse, err error) {
    if request == nil {
        request = NewModifyIntegrationDepartmentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIntegrationDepartment require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIntegrationDepartmentResponse()
    err = c.Send(request, response)
    return
}

func NewModifyIntegrationRoleRequest() (request *ModifyIntegrationRoleRequest) {
    request = &ModifyIntegrationRoleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "ModifyIntegrationRole")
    
    
    return
}

func NewModifyIntegrationRoleResponse() (response *ModifyIntegrationRoleResponse) {
    response = &ModifyIntegrationRoleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyIntegrationRole
// 此接口（ModifyIntegrationRole）用来更新企业自定义的SaaS角色或集团角色。
//
// 
//
// 适用场景1：更新当前企业的自定义SaaS角色或集团角色，并且更新时不进行角色中权限的更新（PermissionGroups 参数不传）。
//
// 
//
// 适用场景2：更新当前企业的自定义SaaS角色或集团角色，并且更新时进行角色中权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 
//
// 适用场景3：更新集团角色管理的子企业列表，可通过设置 SubOrganizationIds 参数达到此效果。
//
// 
//
// 适用场景4：主企业代理子企业操作的场景，需要设置Agent参数，并且ProxyOrganizationId设置为子企业的id即可。
//
// 
//
// 注意事项：SaaS角色和集团角色对应的权限树是不一样的。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyIntegrationRole(request *ModifyIntegrationRoleRequest) (response *ModifyIntegrationRoleResponse, err error) {
    return c.ModifyIntegrationRoleWithContext(context.Background(), request)
}

// ModifyIntegrationRole
// 此接口（ModifyIntegrationRole）用来更新企业自定义的SaaS角色或集团角色。
//
// 
//
// 适用场景1：更新当前企业的自定义SaaS角色或集团角色，并且更新时不进行角色中权限的更新（PermissionGroups 参数不传）。
//
// 
//
// 适用场景2：更新当前企业的自定义SaaS角色或集团角色，并且更新时进行角色中权限的设置（PermissionGroups 参数要传），权限树内容 PermissionGroups 可参考接口 DescribeIntegrationRoles 的输出。此处注意权限树内容可能会更新，需尽量拉取最新的权限树内容，并且权限树内容 PermissionGroups 必须是一颗完整的权限树。
//
// 
//
// 适用场景3：更新集团角色管理的子企业列表，可通过设置 SubOrganizationIds 参数达到此效果。
//
// 
//
// 适用场景4：主企业代理子企业操作的场景，需要设置Agent参数，并且ProxyOrganizationId设置为子企业的id即可。
//
// 
//
// 注意事项：SaaS角色和集团角色对应的权限树是不一样的。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) ModifyIntegrationRoleWithContext(ctx context.Context, request *ModifyIntegrationRoleRequest) (response *ModifyIntegrationRoleResponse, err error) {
    if request == nil {
        request = NewModifyIntegrationRoleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyIntegrationRole require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyIntegrationRoleResponse()
    err = c.Send(request, response)
    return
}

func NewRenewAutoSignLicenseRequest() (request *RenewAutoSignLicenseRequest) {
    request = &RenewAutoSignLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "RenewAutoSignLicense")
    
    
    return
}

func NewRenewAutoSignLicenseResponse() (response *RenewAutoSignLicenseResponse) {
    response = &RenewAutoSignLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RenewAutoSignLicense
// 给医疗个人自动签许可续期。续期成功后，可对医疗自动签许可追加一年有效期，只可续期一次。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LICENSENOQUOTA = "FailedOperation.LicenseNoQuota"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) RenewAutoSignLicense(request *RenewAutoSignLicenseRequest) (response *RenewAutoSignLicenseResponse, err error) {
    return c.RenewAutoSignLicenseWithContext(context.Background(), request)
}

// RenewAutoSignLicense
// 给医疗个人自动签许可续期。续期成功后，可对医疗自动签许可追加一年有效期，只可续期一次。
//
// 
//
// 注意: `处方单等特殊场景专用，此接口为白名单功能，使用前请联系对接的客户经理沟通。`
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LICENSENOQUOTA = "FailedOperation.LicenseNoQuota"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_API = "InternalError.Api"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DEPENDSAPI = "InternalError.DependsApi"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER_CARDNUMBER = "InvalidParameter.CardNumber"
//  INVALIDPARAMETER_CARDTYPE = "InvalidParameter.CardType"
//  INVALIDPARAMETER_NAME = "InvalidParameter.Name"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) RenewAutoSignLicenseWithContext(ctx context.Context, request *RenewAutoSignLicenseRequest) (response *RenewAutoSignLicenseResponse, err error) {
    if request == nil {
        request = NewRenewAutoSignLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RenewAutoSignLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewRenewAutoSignLicenseResponse()
    err = c.Send(request, response)
    return
}

func NewStartFlowRequest() (request *StartFlowRequest) {
    request = &StartFlowRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "StartFlow")
    
    
    return
}

func NewStartFlowResponse() (response *StartFlowResponse) {
    response = &StartFlowResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartFlow
// 此接口用于启动流程。它是模板发起合同的最后一步。
//
// 在[创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)和[创建电子文档](https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument)之后，用于开始整个合同流程,  推进流程进入到签署环节。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/06f2bc0f1772d8deac2f92b5df61a5ac.png)
//
// 
//
// 注：
//
// 1.<font color="red">合同发起后就会扣减合同的额度</font>, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（过期，拒签，签署完成，解除完成等状态不会返还额度）
//
// 
//
// 2.<font color="red">静默（自动）签署不支持合同签署方存在填写</font>功能
//
// 
//
// 3.<font color="red">在发起签署流程之前，建议等待 [PDF合成完成的回调](https://qian.tencent.com/developers/company/callback_types_file_resources)</font>，尤其是当模板中存在动态表格等复杂填写控件时，因为合成过程可能会耗费秒级别的时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_NOTFOUNDSHADOWUSER = "FailedOperation.NotFoundShadowUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOCUMENTNOAVAILABLE = "OperationDenied.DocumentNoAvailable"
//  OPERATIONDENIED_FLOWCANCELFORBID = "OperationDenied.FlowCancelForbid"
//  OPERATIONDENIED_FLOWHASSTARTED = "OperationDenied.FlowHasStarted"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FLOWNONEEDREVIEW = "OperationDenied.FlowNoNeedReview"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartFlow(request *StartFlowRequest) (response *StartFlowResponse, err error) {
    return c.StartFlowWithContext(context.Background(), request)
}

// StartFlow
// 此接口用于启动流程。它是模板发起合同的最后一步。
//
// 在[创建签署流程](https://qian.tencent.com/developers/companyApis/startFlows/CreateFlow)和[创建电子文档](https://qian.tencent.com/developers/companyApis/startFlows/CreateDocument)之后，用于开始整个合同流程,  推进流程进入到签署环节。
//
// 
//
// ![image](https://qcloudimg.tencent-cloud.cn/raw/06f2bc0f1772d8deac2f92b5df61a5ac.png)
//
// 
//
// 注：
//
// 1.<font color="red">合同发起后就会扣减合同的额度</font>, 只有撤销没有参与方签署过或只有自动签署签署过的合同，才会返还合同额度。（过期，拒签，签署完成，解除完成等状态不会返还额度）
//
// 
//
// 2.<font color="red">静默（自动）签署不支持合同签署方存在填写</font>功能
//
// 
//
// 3.<font color="red">在发起签署流程之前，建议等待 [PDF合成完成的回调](https://qian.tencent.com/developers/company/callback_types_file_resources)</font>，尤其是当模板中存在动态表格等复杂填写控件时，因为合成过程可能会耗费秒级别的时间。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FLOWHASNODOCUMENT = "FailedOperation.FlowHasNoDocument"
//  FAILEDOPERATION_NOSIGNREVIEWPASS = "FailedOperation.NoSignReviewPass"
//  FAILEDOPERATION_NOTFOUNDSHADOWUSER = "FailedOperation.NotFoundShadowUser"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.Db"
//  INTERNALERROR_DBCONNECTION = "InternalError.DbConnection"
//  INTERNALERROR_DBREAD = "InternalError.DbRead"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CLIENTTOKEN = "InvalidParameter.ClientToken"
//  INVALIDPARAMETER_INVALIDID = "InvalidParameter.InvalidId"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE_MASK = "InvalidParameterValue.Mask"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPROVERSIGNCOMPONENT = "MissingParameter.ApproverSignComponent"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_DOCUMENTNOAVAILABLE = "OperationDenied.DocumentNoAvailable"
//  OPERATIONDENIED_FLOWCANCELFORBID = "OperationDenied.FlowCancelForbid"
//  OPERATIONDENIED_FLOWHASSTARTED = "OperationDenied.FlowHasStarted"
//  OPERATIONDENIED_FLOWHASTERMINATED = "OperationDenied.FlowHasTerminated"
//  OPERATIONDENIED_FLOWNONEEDREVIEW = "OperationDenied.FlowNoNeedReview"
//  OPERATIONDENIED_FLOWSTATUSFORBID = "OperationDenied.FlowStatusForbid"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  OPERATIONDENIED_NOIDENTITYVERIFY = "OperationDenied.NoIdentityVerify"
//  OPERATIONDENIED_NOLOGIN = "OperationDenied.NoLogin"
//  OPERATIONDENIED_NOOPENSERVERSIGN = "OperationDenied.NoOpenServerSign"
//  OPERATIONDENIED_NOQUOTA = "OperationDenied.NoQuota"
//  OPERATIONDENIED_WHITELISTFORBID = "OperationDenied.WhiteListForbid"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_FLOW = "ResourceNotFound.Flow"
//  RESOURCENOTFOUND_FLOWAPPROVER = "ResourceNotFound.FlowApprover"
//  RESOURCENOTFOUND_NOTEXISTDOCUMENT = "ResourceNotFound.NotExistDocument"
//  RESOURCENOTFOUND_NOTEXISTFLOW = "ResourceNotFound.NotExistFlow"
//  RESOURCENOTFOUND_ORGANIZATION = "ResourceNotFound.Organization"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartFlowWithContext(ctx context.Context, request *StartFlowRequest) (response *StartFlowResponse, err error) {
    if request == nil {
        request = NewStartFlowRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartFlow require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartFlowResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindEmployeeUserIdWithClientOpenIdRequest() (request *UnbindEmployeeUserIdWithClientOpenIdRequest) {
    request = &UnbindEmployeeUserIdWithClientOpenIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "UnbindEmployeeUserIdWithClientOpenId")
    
    
    return
}

func NewUnbindEmployeeUserIdWithClientOpenIdResponse() (response *UnbindEmployeeUserIdWithClientOpenIdResponse) {
    response = &UnbindEmployeeUserIdWithClientOpenIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnbindEmployeeUserIdWithClientOpenId
// 此接口（UnbindEmployeeUserIdWithClientOpenId）用于解除电子签系统员工UserId与客户系统员工OpenId之间的绑定关系。
//
// 
//
// 注：`在调用此接口时，需确保OpenId已通过调用`<a href="https://qian.tencent.com/developers/companyApis/staffs/BindEmployeeUserIdWithClientOpenId" target="_blank">BindEmployeeUserIdWithClientOpenId</a>`接口与电子签系统的UserId绑定过。若OpenId未经过绑定，则无法使用此接口进行解绑操作。`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UnbindEmployeeUserIdWithClientOpenId(request *UnbindEmployeeUserIdWithClientOpenIdRequest) (response *UnbindEmployeeUserIdWithClientOpenIdResponse, err error) {
    return c.UnbindEmployeeUserIdWithClientOpenIdWithContext(context.Background(), request)
}

// UnbindEmployeeUserIdWithClientOpenId
// 此接口（UnbindEmployeeUserIdWithClientOpenId）用于解除电子签系统员工UserId与客户系统员工OpenId之间的绑定关系。
//
// 
//
// 注：`在调用此接口时，需确保OpenId已通过调用`<a href="https://qian.tencent.com/developers/companyApis/staffs/BindEmployeeUserIdWithClientOpenId" target="_blank">BindEmployeeUserIdWithClientOpenId</a>`接口与电子签系统的UserId绑定过。若OpenId未经过绑定，则无法使用此接口进行解绑操作。`
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UnbindEmployeeUserIdWithClientOpenIdWithContext(ctx context.Context, request *UnbindEmployeeUserIdWithClientOpenIdRequest) (response *UnbindEmployeeUserIdWithClientOpenIdResponse, err error) {
    if request == nil {
        request = NewUnbindEmployeeUserIdWithClientOpenIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnbindEmployeeUserIdWithClientOpenId require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnbindEmployeeUserIdWithClientOpenIdResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateIntegrationEmployeesRequest() (request *UpdateIntegrationEmployeesRequest) {
    request = &UpdateIntegrationEmployeesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "UpdateIntegrationEmployees")
    
    
    return
}

func NewUpdateIntegrationEmployeesResponse() (response *UpdateIntegrationEmployeesResponse) {
    response = &UpdateIntegrationEmployeesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateIntegrationEmployees
// 此接口（UpdateIntegrationEmployees）<font color="red"><b>用于修改未实名企业员工信息(姓名，手机号，邮件、部门)</b></font>。
//
// 如果企业员工已经实名， 姓名，手机号，邮件等需要企业员工到小程序或者控制台自己修改， 部门则需要超管到控制台分配
//
// 
//
// 修改手机号的时候,支持以下场景进行提醒通知
//
// <table>
//
// <tbody>
//
// <tr>
//
// <td>生成端</td>
//
// <td >入参</td>
//
// <td>提醒方式</td>
//
// </tr>
//
// <tr>
//
// <td>普通saas员工</td>
//
// <td>将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号</td>
//
// <td>发送短信通知员工（短信中带有认证加入企业的链接）</td>
//
// </tr>
//
// <tr>
//
// <td>企微员工</td>
//
// <td>将Employees 中的WeworkOpenId字段设置为企微员工明文的openid，需<font color="red">确保该企微员工在应用的可见范围内</font></td>
//
// <td>企微内部实名消息</td>
//
// </tr>
//
// <tr>
//
// <td>H5端 saas员工</td>
//
// <td>传递 InvitationNotifyType = H5，将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号，<font color="red">此场景不支持企微</font></td>
//
// <td>生成认证加入企业的H5链接，贵方可以通过自己的渠道触达到此员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 注意：
//
// 
//
// - <b>若通过手机号发现员工已经创建，则不会重复创建，但会发送短信或者生成链接提醒员工实名。</b>
//
// - jumpUrl 仅支持H5的邀请方式，回跳的url，使用前请联系对接的客户经理沟通，进行域名的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
func (c *Client) UpdateIntegrationEmployees(request *UpdateIntegrationEmployeesRequest) (response *UpdateIntegrationEmployeesResponse, err error) {
    return c.UpdateIntegrationEmployeesWithContext(context.Background(), request)
}

// UpdateIntegrationEmployees
// 此接口（UpdateIntegrationEmployees）<font color="red"><b>用于修改未实名企业员工信息(姓名，手机号，邮件、部门)</b></font>。
//
// 如果企业员工已经实名， 姓名，手机号，邮件等需要企业员工到小程序或者控制台自己修改， 部门则需要超管到控制台分配
//
// 
//
// 修改手机号的时候,支持以下场景进行提醒通知
//
// <table>
//
// <tbody>
//
// <tr>
//
// <td>生成端</td>
//
// <td >入参</td>
//
// <td>提醒方式</td>
//
// </tr>
//
// <tr>
//
// <td>普通saas员工</td>
//
// <td>将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号</td>
//
// <td>发送短信通知员工（短信中带有认证加入企业的链接）</td>
//
// </tr>
//
// <tr>
//
// <td>企微员工</td>
//
// <td>将Employees 中的WeworkOpenId字段设置为企微员工明文的openid，需<font color="red">确保该企微员工在应用的可见范围内</font></td>
//
// <td>企微内部实名消息</td>
//
// </tr>
//
// <tr>
//
// <td>H5端 saas员工</td>
//
// <td>传递 InvitationNotifyType = H5，将Employees中的DisplayName设置员工的名字，Mobile设置成员工的手机号，<font color="red">此场景不支持企微</font></td>
//
// <td>生成认证加入企业的H5链接，贵方可以通过自己的渠道触达到此员工</td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// 注意：
//
// 
//
// - <b>若通过手机号发现员工已经创建，则不会重复创建，但会发送短信或者生成链接提醒员工实名。</b>
//
// - jumpUrl 仅支持H5的邀请方式，回跳的url，使用前请联系对接的客户经理沟通，进行域名的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOAPIAUTH = "OperationDenied.NoApiAuth"
func (c *Client) UpdateIntegrationEmployeesWithContext(ctx context.Context, request *UpdateIntegrationEmployeesRequest) (response *UpdateIntegrationEmployeesResponse, err error) {
    if request == nil {
        request = NewUpdateIntegrationEmployeesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateIntegrationEmployees require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateIntegrationEmployeesResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFilesRequest() (request *UploadFilesRequest) {
    request = &UploadFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "UploadFiles")
    
    
    return
}

func NewUploadFilesResponse() (response *UploadFilesResponse) {
    response = &UploadFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadFiles
// 此接口（UploadFiles）文件上传。<br/>
//
// 
//
// 适用场景：用于合同，印章的文件上传。文件上传以后，
//
// 如果是PDF格式文件可配合<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowByFiles" target="_blank">用PDF文件创建签署流程</a>接口进行合同流程的发起
//
// 如果是其他类型可以配合<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务</a>接口转换成PDF文件
//
// 
//
// 注: 
//
// 1. 图片类型(png/jpg/jpeg)限制大小为5M以下, PDF/word/excel等其他格式限制大小为60M以下
//
// 2. 调用此接口时需要设置单独的Domain请求域名，<font color="red">联调开发环境为: file.test.ess.tencent.cn，正式环境需要设置为:file.ess.tencent.cn</font>，代码示例
//
// ```
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// ```
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UploadFiles(request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    return c.UploadFilesWithContext(context.Background(), request)
}

// UploadFiles
// 此接口（UploadFiles）文件上传。<br/>
//
// 
//
// 适用场景：用于合同，印章的文件上传。文件上传以后，
//
// 如果是PDF格式文件可配合<a href="https://qian.tencent.com/developers/companyApis/startFlows/CreateFlowByFiles" target="_blank">用PDF文件创建签署流程</a>接口进行合同流程的发起
//
// 如果是其他类型可以配合<a href="https://qian.tencent.com/developers/companyApis/templatesAndFiles/CreateConvertTaskApi" target="_blank">创建文件转换任务</a>接口转换成PDF文件
//
// 
//
// 注: 
//
// 1. 图片类型(png/jpg/jpeg)限制大小为5M以下, PDF/word/excel等其他格式限制大小为60M以下
//
// 2. 调用此接口时需要设置单独的Domain请求域名，<font color="red">联调开发环境为: file.test.ess.tencent.cn，正式环境需要设置为:file.ess.tencent.cn</font>，代码示例
//
// ```
//
// HttpProfile httpProfile = new HttpProfile();
//
// httpProfile.setEndpoint("file.test.ess.tencent.cn");
//
// ```
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_ENCRYPTION = "InternalError.Encryption"
//  INVALIDPARAMETER_SENSITIVEFILECONTENT = "InvalidParameter.SensitiveFileContent"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) UploadFilesWithContext(ctx context.Context, request *UploadFilesRequest) (response *UploadFilesResponse, err error) {
    if request == nil {
        request = NewUploadFilesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFilesResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyPdfRequest() (request *VerifyPdfRequest) {
    request = &VerifyPdfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ess", APIVersion, "VerifyPdf")
    
    
    return
}

func NewVerifyPdfResponse() (response *VerifyPdfResponse) {
    response = &VerifyPdfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyPdf
// 对合同流程文件进行数字签名验证，判断数字签名是否有效，合同文件内容是否被篡改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) VerifyPdf(request *VerifyPdfRequest) (response *VerifyPdfResponse, err error) {
    return c.VerifyPdfWithContext(context.Background(), request)
}

// VerifyPdf
// 对合同流程文件进行数字签名验证，判断数字签名是否有效，合同文件内容是否被篡改。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  OPERATIONDENIED_ERRNORESOURCEACCESS = "OperationDenied.ErrNoResourceAccess"
//  OPERATIONDENIED_FORBID = "OperationDenied.Forbid"
//  OPERATIONDENIED_NOFLOWPERMISSION = "OperationDenied.NoFlowPermission"
//  UNAUTHORIZEDOPERATION_NOPERMISSIONFEATURE = "UnauthorizedOperation.NoPermissionFeature"
func (c *Client) VerifyPdfWithContext(ctx context.Context, request *VerifyPdfRequest) (response *VerifyPdfResponse, err error) {
    if request == nil {
        request = NewVerifyPdfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyPdf require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyPdfResponse()
    err = c.Send(request, response)
    return
}
