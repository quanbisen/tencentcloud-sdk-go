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

package v20221229

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Filter struct {
	// 过滤不满足分辨率下限的训练图像，默认开启过滤
	// 开启后将过滤横边<512或竖边<720的图片，横、竖边上限均为2000，不支持调整
	// 
	// 1：开启过滤
	// 0：关闭过滤
	Resolution *int64 `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 过滤脸部区域过小的训练图像，默认开启过滤
	// 
	// 1：开启过滤
	// 0：关闭过滤
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 过滤脸部存在明显遮挡、偏转角度过大等质量较差的训练图像，默认开启过滤
	// 
	// 1：开启过滤
	// 0：关闭过滤
	Occlusion *int64 `json:"Occlusion,omitnil,omitempty" name:"Occlusion"`
}

// Predefined struct for user
type ImageToImageRequestParams struct {
	// 输入图 Base64 数据。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于5000且大于50，转成 Base64 字符串后小于 8MB。
	InputImage *string `json:"InputImage,omitnil,omitempty" name:"InputImage"`

	// 输入图 Url。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于5000且大于50，转成 Base64 字符串后小于8MB。
	InputUrl *string `json:"InputUrl,omitnil,omitempty" name:"InputUrl"`

	// 文本描述。
	// 用于在输入图的基础上引导生成图效果，增加生成结果中出现描述内容的可能。
	// 推荐使用中文。最多支持256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在  [智能图生图风格列表](https://cloud.tencent.com/document/product/1668/86250) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	Styles []*string `json:"Styles,omitnil,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	// 支持生成以下分辨率的图片：origin（与输入图分辨率一致，长边最高为2000，超出将做等比例缩小）、768:768（1:1）、768:1024（3:4）、1024:768（4:3）。
	// 不传默认使用origin。
	ResultConfig *ResultConfig `json:"ResultConfig,omitnil,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 生成自由度。
	// Strength 值越小，生成图和原图越接近，取值范围(0, 1]，不传使用模型内置的默认值。
	// 推荐的取值范围为0.6 - 0.8。
	Strength *float64 `json:"Strength,omitnil,omitempty" name:"Strength"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type ImageToImageRequest struct {
	*tchttp.BaseRequest
	
	// 输入图 Base64 数据。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于5000且大于50，转成 Base64 字符串后小于 8MB。
	InputImage *string `json:"InputImage,omitnil,omitempty" name:"InputImage"`

	// 输入图 Url。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于5000且大于50，转成 Base64 字符串后小于8MB。
	InputUrl *string `json:"InputUrl,omitnil,omitempty" name:"InputUrl"`

	// 文本描述。
	// 用于在输入图的基础上引导生成图效果，增加生成结果中出现描述内容的可能。
	// 推荐使用中文。最多支持256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在  [智能图生图风格列表](https://cloud.tencent.com/document/product/1668/86250) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	Styles []*string `json:"Styles,omitnil,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	// 支持生成以下分辨率的图片：origin（与输入图分辨率一致，长边最高为2000，超出将做等比例缩小）、768:768（1:1）、768:1024（3:4）、1024:768（4:3）。
	// 不传默认使用origin。
	ResultConfig *ResultConfig `json:"ResultConfig,omitnil,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 生成自由度。
	// Strength 值越小，生成图和原图越接近，取值范围(0, 1]，不传使用模型内置的默认值。
	// 推荐的取值范围为0.6 - 0.8。
	Strength *float64 `json:"Strength,omitnil,omitempty" name:"Strength"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

func (r *ImageToImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputImage")
	delete(f, "InputUrl")
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Styles")
	delete(f, "ResultConfig")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "Strength")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToImageResponseParams struct {
	// 根据入参 RspImgType 填入不同，返回不同的内容。
	// 如果传入 base64 则返回生成图 Base64 编码。
	// 如果传入 url 则返回的生成图 URL , 有效期1小时，请及时保存。
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageToImageResponse struct {
	*tchttp.BaseResponse
	Response *ImageToImageResponseParams `json:"Response"`
}

func (r *ImageToImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogoParam struct {
	// 水印url
	LogoUrl *string `json:"LogoUrl,omitnil,omitempty" name:"LogoUrl"`

	// 水印base64，url和base64二选一传入
	LogoImage *string `json:"LogoImage,omitnil,omitempty" name:"LogoImage"`

	// 水印图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配
	LogoRect *LogoRect `json:"LogoRect,omitnil,omitempty" name:"LogoRect"`
}

type LogoRect struct {
	// 左上角X坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上角Y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 方框宽度
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 方框高度
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

// Predefined struct for user
type QueryDrawPortraitJobRequestParams struct {
	// 查询生成写真图片任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryDrawPortraitJobRequest struct {
	*tchttp.BaseRequest
	
	// 查询生成写真图片任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryDrawPortraitJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDrawPortraitJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryDrawPortraitJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryDrawPortraitJobResponseParams struct {
	// 任务状态码。
	// INIT: 初始化、WAIT：等待中、RUN：运行中、FAIL：处理失败、DONE：处理完成。
	JobStatusCode *string `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// 任务状态信息。
	JobStatusMsg *string `json:"JobStatusMsg,omitnil,omitempty" name:"JobStatusMsg"`

	// 任务错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 结果 URL 数组。
	// URL 有效期1小时，请及时保存。
	ResultUrls []*string `json:"ResultUrls,omitnil,omitempty" name:"ResultUrls"`

	// 结果描述数组。
	ResultDetails []*string `json:"ResultDetails,omitnil,omitempty" name:"ResultDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryDrawPortraitJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryDrawPortraitJobResponseParams `json:"Response"`
}

func (r *QueryDrawPortraitJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryDrawPortraitJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryTextToImageProJobRequestParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

type QueryTextToImageProJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`
}

func (r *QueryTextToImageProJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTextToImageProJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTextToImageProJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryTextToImageProJobResponseParams struct {
	// 当前任务状态码：
	// 1：等待中、2：运行中、4：处理失败、5：处理完成。
	JobStatusCode *string `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// 当前任务状态：排队中、处理中、处理失败或者处理完成。
	JobStatusMsg *string `json:"JobStatusMsg,omitnil,omitempty" name:"JobStatusMsg"`

	// 任务处理失败错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务处理失败错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 生成图 URL 列表，有效期1小时，请及时保存。
	ResultImage []*string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 结果 detail 数组，Success 代表成功。
	ResultDetails []*string `json:"ResultDetails,omitnil,omitempty" name:"ResultDetails"`

	// 对应 SubmitTextToImageProJob 接口中 Revise 参数。开启扩写时，返回扩写后的 prompt 文本。 如果关闭扩写，将直接返回原始输入的 prompt。
	RevisedPrompt []*string `json:"RevisedPrompt,omitnil,omitempty" name:"RevisedPrompt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryTextToImageProJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryTextToImageProJobResponseParams `json:"Response"`
}

func (r *QueryTextToImageProJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTextToImageProJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryTrainPortraitModelJobRequestParams struct {
	// 写真模型 ID。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type QueryTrainPortraitModelJobRequest struct {
	*tchttp.BaseRequest
	
	// 写真模型 ID。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

func (r *QueryTrainPortraitModelJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTrainPortraitModelJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTrainPortraitModelJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryTrainPortraitModelJobResponseParams struct {
	// 任务状态码。
	// INIT: 初始化、WAIT：等待中、RUN：运行中、FAIL：处理失败、DONE：处理完成。
	JobStatusCode *string `json:"JobStatusCode,omitnil,omitempty" name:"JobStatusCode"`

	// 任务状态信息。
	JobStatusMsg *string `json:"JobStatusMsg,omitnil,omitempty" name:"JobStatusMsg"`

	// 任务错误码。
	JobErrorCode *string `json:"JobErrorCode,omitnil,omitempty" name:"JobErrorCode"`

	// 任务错误信息。
	JobErrorMsg *string `json:"JobErrorMsg,omitnil,omitempty" name:"JobErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryTrainPortraitModelJobResponse struct {
	*tchttp.BaseResponse
	Response *QueryTrainPortraitModelJobResponseParams `json:"Response"`
}

func (r *QueryTrainPortraitModelJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTrainPortraitModelJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResultConfig struct {
	// 生成图分辨率
	// 
	// 智能文生图支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3）、1080:1920（9:16）、1920:1080（16:9），不传默认使用768:768。
	// 
	// 智能图生图支持生成以下分辨率的图片：origin（与输入图分辨率一致，长边最高为2000，超出将做等比例缩小）、768:768（1:1）、768:1024（3:4）、1024:768（4:3），不传默认使用origin，如果指定生成的长宽比与输入图长宽比差异过大可能导致图片内容被裁剪。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`
}

// Predefined struct for user
type SubmitDrawPortraitJobRequestParams struct {
	// 写真模型 ID。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 写真风格模板。
	// 请在[ AI 写真风格列表](https://cloud.tencent.com/document/product/1668/105740) 中选择期望的风格，传入风格编号。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 本次生成的图片数量，取值范围[1,4]
	ImageNum *int64 `json:"ImageNum,omitnil,omitempty" name:"ImageNum"`

	// 为生成结果图添加标识的开关，默认为1。 
	// 1：添加标识。
	//  0：不添加标识。 
	// 其他数值：默认按1处理。 
	// 建议您使用显著标识来提示结果图是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 清晰度，支持以下选项：
	// sd：基础版，分辨率512:640
	// hd：高清畅享版，分辨率1024:1280
	// hdpro：高清优享版，分辨率1024:1280（推荐）
	// uhd：超清版，分辨率2048:2560
	// 不填默认为sd。
	Definition *string `json:"Definition,omitnil,omitempty" name:"Definition"`
}

type SubmitDrawPortraitJobRequest struct {
	*tchttp.BaseRequest
	
	// 写真模型 ID。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 写真风格模板。
	// 请在[ AI 写真风格列表](https://cloud.tencent.com/document/product/1668/105740) 中选择期望的风格，传入风格编号。
	StyleId *string `json:"StyleId,omitnil,omitempty" name:"StyleId"`

	// 本次生成的图片数量，取值范围[1,4]
	ImageNum *int64 `json:"ImageNum,omitnil,omitempty" name:"ImageNum"`

	// 为生成结果图添加标识的开关，默认为1。 
	// 1：添加标识。
	//  0：不添加标识。 
	// 其他数值：默认按1处理。 
	// 建议您使用显著标识来提示结果图是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。 
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 清晰度，支持以下选项：
	// sd：基础版，分辨率512:640
	// hd：高清畅享版，分辨率1024:1280
	// hdpro：高清优享版，分辨率1024:1280（推荐）
	// uhd：超清版，分辨率2048:2560
	// 不填默认为sd。
	Definition *string `json:"Definition,omitnil,omitempty" name:"Definition"`
}

func (r *SubmitDrawPortraitJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitDrawPortraitJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "StyleId")
	delete(f, "ImageNum")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "Definition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitDrawPortraitJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitDrawPortraitJobResponseParams struct {
	// 提交生成写真图片任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitDrawPortraitJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitDrawPortraitJobResponseParams `json:"Response"`
}

func (r *SubmitDrawPortraitJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitDrawPortraitJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextToImageProJobRequestParams struct {
	// 文本描述。 
	// 算法将根据输入的文本智能生成与之相关的图像。 
	// 不能为空，推荐使用中文。最多可传100个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 绘画风格。
	// 请在 [文生图（高级版）风格列表](https://cloud.tencent.com/document/product/1668/104567) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 文生图模型，默认使用engine1。
	// 取值：
	// engine1
	// engine2
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`
}

type SubmitTextToImageProJobRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。 
	// 算法将根据输入的文本智能生成与之相关的图像。 
	// 不能为空，推荐使用中文。最多可传100个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 绘画风格。
	// 请在 [文生图（高级版）风格列表](https://cloud.tencent.com/document/product/1668/104567) 中选择期望的风格，传入风格编号。
	// 不传默认不指定风格。
	Style *string `json:"Style,omitnil,omitempty" name:"Style"`

	// 生成图分辨率。
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3），不传默认使用1024:1024。
	Resolution *string `json:"Resolution,omitnil,omitempty" name:"Resolution"`

	// 为生成结果图添加显式水印标识的开关，默认为1。  
	// 1：添加。  
	// 0：不添加。  
	// 其他数值：默认按1处理。  
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 文生图模型，默认使用engine1。
	// 取值：
	// engine1
	// engine2
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// prompt 扩写开关。1为开启，0为关闭，不传默认开启。
	// 开启扩写后，将自动扩写原始输入的 prompt 并使用扩写后的 prompt 生成图片，返回生成图片结果时将一并返回扩写后的 prompt 文本。
	// 如果关闭扩写，将直接使用原始输入的 prompt 生成图片。
	// 建议开启，在多数场景下可提升生成图片效果、丰富生成图片细节。
	Revise *int64 `json:"Revise,omitnil,omitempty" name:"Revise"`
}

func (r *SubmitTextToImageProJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextToImageProJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "Style")
	delete(f, "Resolution")
	delete(f, "LogoAdd")
	delete(f, "Engine")
	delete(f, "Revise")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTextToImageProJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTextToImageProJobResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTextToImageProJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTextToImageProJobResponseParams `json:"Response"`
}

func (r *SubmitTextToImageProJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTextToImageProJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTrainPortraitModelJobRequestParams struct {
	// 在上传写真训练图片时指定的写真模型 ID。 
	// 每个 AI 写真模型自训练完成起1年内有效，有效期内可使用模型生成图片，期满后需要重新训练模型。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

type SubmitTrainPortraitModelJobRequest struct {
	*tchttp.BaseRequest
	
	// 在上传写真训练图片时指定的写真模型 ID。 
	// 每个 AI 写真模型自训练完成起1年内有效，有效期内可使用模型生成图片，期满后需要重新训练模型。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`
}

func (r *SubmitTrainPortraitModelJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTrainPortraitModelJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTrainPortraitModelJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTrainPortraitModelJobResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTrainPortraitModelJobResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTrainPortraitModelJobResponseParams `json:"Response"`
}

func (r *SubmitTrainPortraitModelJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTrainPortraitModelJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToImageRequestParams struct {
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [智能文生图风格列表](https://cloud.tencent.com/document/product/1668/86249) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	Styles []*string `json:"Styles,omitnil,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	// 
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3）、1080:1920（9:16）、1920:1080（16:9），不传默认使用768:768。
	ResultConfig *ResultConfig `json:"ResultConfig,omitnil,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

type TextToImageRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传256个 utf-8 字符。
	Prompt *string `json:"Prompt,omitnil,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传256个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitnil,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [智能文生图风格列表](https://cloud.tencent.com/document/product/1668/86249) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	Styles []*string `json:"Styles,omitnil,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	// 
	// 支持生成以下分辨率的图片：768:768（1:1）、768:1024（3:4）、1024:768（4:3）、1024:1024（1:1）、720:1280（9:16）、1280:720（16:9）、768:1280（3:5）、1280:768（5:3）、1080:1920（9:16）、1920:1080（16:9），不传默认使用768:768。
	ResultConfig *ResultConfig `json:"ResultConfig,omitnil,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitnil,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitnil,omitempty" name:"LogoParam"`

	// 返回图像方式（base64 或 url) ，二选一，默认为 base64。url 有效期为1小时。
	RspImgType *string `json:"RspImgType,omitnil,omitempty" name:"RspImgType"`
}

func (r *TextToImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Styles")
	delete(f, "ResultConfig")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "RspImgType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToImageResponseParams struct {
	// 根据入参 RspImgType 填入不同，返回不同的内容。
	// 如果传入 base64 则返回生成图 Base64 编码。
	// 如果传入 url 则返回的生成图 URL , 有效期1小时，请及时保存。
	ResultImage *string `json:"ResultImage,omitnil,omitempty" name:"ResultImage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextToImageResponse struct {
	*tchttp.BaseResponse
	Response *TextToImageResponseParams `json:"Response"`
}

func (r *TextToImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadTrainPortraitImagesRequestParams struct {
	// 写真模型 ID。由英文大小写字母、数字及下划线组成。
	// 用于唯一标识一个写真模型，一个写真模型只能用于一个人物的写真图片生成。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 写真模型训练用的基础图像 URL，用于固定写真模型可生成的人物。
	// 图片数量：1张。
	// 图片内容：单人，脸部清晰。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	BaseUrl *string `json:"BaseUrl,omitnil,omitempty" name:"BaseUrl"`

	// 写真模型训练用的图像 URL 列表。
	// 图片数量：19 - 24 张。
	// 图片内容：单人，脸部清晰，和基础图像中的人物为同一人。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// 训练图像质量过滤开关配置。
	// 支持开启或关闭对训练图像分辨率下限、脸部区域大小、脸部遮挡的过滤，默认开启以上过滤。
	// 如果训练图像内包含多人脸或无人脸、和 Base 人像不为同一人也将被过滤，不可关闭该过滤条件。
	// 建议：关闭以上过滤可能导致写真生成效果受损，建议使用单人、正脸、脸部区域占比较大、脸部清晰无遮挡、无大角度偏转、无夸张表情的图像进行训练。
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 是否开启快速训练模式。
	// 默认不开启。开启后只需要在 BaseUrl 中传入1张图片，Urls.N 中无需传入图片。  
	// 0：不开启  
	// 1：开启
	TrainMode *int64 `json:"TrainMode,omitnil,omitempty" name:"TrainMode"`
}

type UploadTrainPortraitImagesRequest struct {
	*tchttp.BaseRequest
	
	// 写真模型 ID。由英文大小写字母、数字及下划线组成。
	// 用于唯一标识一个写真模型，一个写真模型只能用于一个人物的写真图片生成。
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 写真模型训练用的基础图像 URL，用于固定写真模型可生成的人物。
	// 图片数量：1张。
	// 图片内容：单人，脸部清晰。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	BaseUrl *string `json:"BaseUrl,omitnil,omitempty" name:"BaseUrl"`

	// 写真模型训练用的图像 URL 列表。
	// 图片数量：19 - 24 张。
	// 图片内容：单人，脸部清晰，和基础图像中的人物为同一人。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// 训练图像质量过滤开关配置。
	// 支持开启或关闭对训练图像分辨率下限、脸部区域大小、脸部遮挡的过滤，默认开启以上过滤。
	// 如果训练图像内包含多人脸或无人脸、和 Base 人像不为同一人也将被过滤，不可关闭该过滤条件。
	// 建议：关闭以上过滤可能导致写真生成效果受损，建议使用单人、正脸、脸部区域占比较大、脸部清晰无遮挡、无大角度偏转、无夸张表情的图像进行训练。
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 是否开启快速训练模式。
	// 默认不开启。开启后只需要在 BaseUrl 中传入1张图片，Urls.N 中无需传入图片。  
	// 0：不开启  
	// 1：开启
	TrainMode *int64 `json:"TrainMode,omitnil,omitempty" name:"TrainMode"`
}

func (r *UploadTrainPortraitImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadTrainPortraitImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "BaseUrl")
	delete(f, "Urls")
	delete(f, "Filter")
	delete(f, "TrainMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadTrainPortraitImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadTrainPortraitImagesResponseParams struct {
	// 用于提示对应上传的Urls训练图片是否符合要求，如果未通过需要重新上传。如果基础图像不符合要求会直接通过ErrorCode提示。如果您选择了快速模式，该参数返回为空数组。
	ResultDetails []*string `json:"ResultDetails,omitnil,omitempty" name:"ResultDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadTrainPortraitImagesResponse struct {
	*tchttp.BaseResponse
	Response *UploadTrainPortraitImagesResponseParams `json:"Response"`
}

func (r *UploadTrainPortraitImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadTrainPortraitImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}