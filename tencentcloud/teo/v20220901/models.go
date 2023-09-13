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

package v20220901

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccelerateMainland struct {
	// 是否开启中国大陆加速优化配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type AccelerateType struct {
	// 加速开关。取值范围：
	// <li> on：打开;</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type AccelerationDomain struct {
	// 源站信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginDetail *OriginDetail `json:"OriginDetail,omitnil" name:"OriginDetail"`

	// 创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil" name:"ModifiedOn"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速域名状态，取值有：
	// <li>online：已生效；</li>
	// <li>process：部署中；</li>
	// <li>offline：已停用；</li>
	// <li>forbidden：已封禁；</li>
	// <li>init：未生效，待激活站点；</li>
	DomainStatus *string `json:"DomainStatus,omitnil" name:"DomainStatus"`

	// CNAME 地址。
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 加速域名归属权验证状态，取值有： <li>pending：待验证；</li> <li>finished：已完成验证。</li>	
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentificationStatus *string `json:"IdentificationStatus,omitnil" name:"IdentificationStatus"`
}

type AclCondition struct {
	// 匹配字段，取值有：
	// <li>host：请求域名；</li>
	// <li>sip：客户端IP；</li>
	// <li>ua：User-Agent；</li>
	// <li>cookie：会话 Cookie；</li>
	// <li>cgi：CGI 脚本；</li>
	// <li>xff：XFF 扩展头部；</li>
	// <li>url：请求 URL；</li>
	// <li>accept：请求内容类型；</li>
	// <li>method：请求方式；</li>
	// <li>header：请求头部；</li>
	// <li>app_proto：应用层协议；</li>
	// <li>sip_proto：网络层协议；</li>
	// <li>uabot：UA 特征规则，仅bot自定义规则可用；</li>
	// <li>idcid：IDC 规则，仅bot自定义规则可用；</li>
	// <li>sipbot：搜索引擎规则，仅bot自定义规则可用；</li>
	// <li>portrait：画像分析，仅bot自定义规则可用；</li>
	// <li>header_seq：请求头顺序，仅bot自定义规则可用；</li>
	// <li>hdr：请求正文，仅Web防护自定义规则可用。</li>
	MatchFrom *string `json:"MatchFrom,omitnil" name:"MatchFrom"`

	// 匹配字符串。当 MatchFrom 为 header 时，可以填入 header 的 key 作为参数。
	MatchParam *string `json:"MatchParam,omitnil" name:"MatchParam"`

	// 匹配关系，取值有：
	// <li>equal：字符串等于；</li>
	// <li>not_equal：数值不等于；</li>
	// <li>include：字符包含；</li>
	// <li>not_include：字符不包含；</li>
	// <li>match：ip匹配；</li>
	// <li>not_match：ip不匹配；</li>
	// <li>include_area：地域包含；</li>
	// <li>is_empty：存在字段但值为空；</li>
	// <li>not_exists：不存在关键字段；</li>
	// <li>regexp：正则匹配；</li>
	// <li>len_gt：数值大于；</li>
	// <li>len_lt：数值小于；</li>
	// <li>len_eq：数值等于；</li>
	// <li>match_prefix：前缀匹配；</li>
	// <li>match_suffix：后缀匹配；</li>
	// <li>wildcard：通配符。</li>
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 匹配内容。
	MatchContent *string `json:"MatchContent,omitnil" name:"MatchContent"`
}

type AclConfig struct {
	// 开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 用户自定义规则。
	AclUserRules []*AclUserRule `json:"AclUserRules,omitnil" name:"AclUserRules"`

	// 托管定制规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Customizes []*AclUserRule `json:"Customizes,omitnil" name:"Customizes"`
}

type AclUserRule struct {
	// 规则名。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 处罚动作，取值有：
	// <li>trans：放行；</li>
	// <li>drop：拦截；</li>
	// <li>monitor：观察；</li>
	// <li>ban：IP封禁；</li>
	// <li>redirect：重定向；</li>
	// <li>page：指定页面；</li>
	// <li>alg：Javascript挑战。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：失效。</li>
	RuleStatus *string `json:"RuleStatus,omitnil" name:"RuleStatus"`

	// 自定义规则。
	AclConditions []*AclCondition `json:"AclConditions,omitnil" name:"AclConditions"`

	// 规则优先级，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitnil" name:"RulePriority"`

	// 规则Id。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 更新时间。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// ip封禁的惩罚时间，取值范围0-2天。默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitnil" name:"PunishTime"`

	// ip封禁的惩罚时间单位，取值有：
	// <li>second：秒；</li>
	// <li>minutes：分；</li>
	// <li>hour：小时。</li>默认为second。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTimeUnit *string `json:"PunishTimeUnit,omitnil" name:"PunishTimeUnit"`

	// 自定义返回页面的名称。默认为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 自定义返回页面的实例id。默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageId *int64 `json:"PageId,omitnil" name:"PageId"`

	// 重定向时候的地址，必须为本用户接入的站点子域名。默认为空字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUrl *string `json:"RedirectUrl,omitnil" name:"RedirectUrl"`

	// 重定向时候的返回码。默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseCode *int64 `json:"ResponseCode,omitnil" name:"ResponseCode"`
}

type Action struct {
	// 常规功能操作，选择该类型的功能项有：
	// <li> 访问URL 重写（AccessUrlRedirect）；</li>
	// <li> 回源 URL 重写 （UpstreamUrlRedirect）；</li>
	// <li> QUIC（QUIC）；</li>
	// <li> WebSocket （WebSocket）；</li>
	// <li> 视频拖拽（VideoSeek）；</li>
	// <li> Token 鉴权（Authentication）；</li>
	// <li> 自定义CacheKey（CacheKey）；</li>
	// <li> 节点缓存 TTL （Cache）；</li>
	// <li> 浏览器缓存 TTL（MaxAge）；</li>
	// <li> 离线缓存（OfflineCache）；</li>
	// <li> 智能加速（SmartRouting）；</li>
	// <li> 分片回源（RangeOriginPull）；</li>
	// <li> HTTP/2 回源（UpstreamHttp2）；</li>
	// <li> Host Header 重写（HostHeader）；</li>
	// <li> 强制 HTTPS（ForceRedirect）；</li>
	// <li> 回源 HTTPS（OriginPullProtocol）；</li>
	// <li> 缓存预刷新（CachePrefresh）；</li>
	// <li> 智能压缩（Compression）；</li>
	// <li> Hsts；</li>
	// <li> ClientIpHeader；</li>
	// <li> SslTlsSecureConf；</li>
	// <li> OcspStapling；</li>
	// <li> HTTP/2 访问（Http2）；</li>
	// <li> 回源跟随重定向(UpstreamFollowRedirect)；</li>
	// <li> 修改源站(Origin)。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalAction *NormalAction `json:"NormalAction,omitnil" name:"NormalAction"`

	// 带有请求头/响应头的功能操作，选择该类型的功能项有：
	// <li> 修改 HTTP 请求头（RequestHeader）；</li>
	// <li> 修改HTTP响应头（ResponseHeader）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteAction *RewriteAction `json:"RewriteAction,omitnil" name:"RewriteAction"`

	// 带有状态码的功能操作，选择该类型的功能项有：
	// <li> 自定义错误页面（ErrorPage）；</li>
	// <li> 状态码缓存 TTL（StatusCodeCache）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeAction *CodeAction `json:"CodeAction,omitnil" name:"CodeAction"`
}

type AdvancedFilter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 是否启用模糊查询。
	Fuzzy *bool `json:"Fuzzy,omitnil" name:"Fuzzy"`
}

type AiRule struct {
	// AI规则引擎状态，取值有：
	// <li> smart_status_close：关闭；</li>
	// <li> smart_status_open：拦截处置；</li>
	// <li> smart_status_observe：观察处置。</li>
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

type AlgDetectJS struct {
	// 操作名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 工作量证明 (proof_Of-Work)校验强度，默认low，取值有：
	// <li>low：低；</li>
	// <li>middle：中；</li>
	// <li>high：高。</li>
	WorkLevel *string `json:"WorkLevel,omitnil" name:"WorkLevel"`

	// 执行方式，js延迟执行的时间。单位为ms，默认500，取值：0～1000。
	ExecuteMode *int64 `json:"ExecuteMode,omitnil" name:"ExecuteMode"`

	// 客户端末启用JS（末完成检测）统计周期。单位为秒，默认10，取值：5～3600。
	InvalidStatTime *int64 `json:"InvalidStatTime,omitnil" name:"InvalidStatTime"`

	// 客户端末启用JS（末完成检测）触发阈值。单位为次，默认300，取值：1～100000000。
	InvalidThreshold *int64 `json:"InvalidThreshold,omitnil" name:"InvalidThreshold"`

	// Bot主动特征识别客户端行为校验结果。
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitnil" name:"AlgDetectResults"`
}

type AlgDetectResult struct {
	// 校验结果，取值有：
	// <li>invalid：不合法Cookie；</li>
	// <li>cookie_empty：末携带Cookie或Cookie己过期；</li>
	// <li>js_empty：客户端末启用JS（末完成检测）；</li>
	// <li>low：会话速率和周期特征校验低风险；</li>
	// <li>middle：会话速率和周期特征校验中风险；</li>
	// <li>high：会话速率和周期特征校验高风险；</li>
	// <li>timeout：检测超时时长；</li>
	// <li>not_browser：不合法浏览器；</li>
	// <li>is_bot：Bot客户端。</li>
	Result *string `json:"Result,omitnil" name:"Result"`

	// 处罚动作，取值有：
	// <li>drop：拦截；</li>
	// <li>monitor：观察；</li>
	// <li>silence：静默；</li>
	// <li>shortdelay：（短时间）等待后响应；</li>
	// <li>longdelay：（长时间）等待后响应。</li>
	Action *string `json:"Action,omitnil" name:"Action"`
}

type AlgDetectRule struct {
	// 规则id。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 规则名。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则开关。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义规则。
	AlgConditions []*AclCondition `json:"AlgConditions,omitnil" name:"AlgConditions"`

	// Cookie校验和会话行为分析。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgDetectSession *AlgDetectSession `json:"AlgDetectSession,omitnil" name:"AlgDetectSession"`

	// 客户端行为校验。
	AlgDetectJS []*AlgDetectJS `json:"AlgDetectJS,omitnil" name:"AlgDetectJS"`

	// 更新时间。仅出参使用。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type AlgDetectSession struct {
	// 操作名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 校验方式，默认update_detect，取值有：
	// <li>detect：仅校验；</li>
	// <li>update_detect：更新Cookie并校验。</li>
	DetectMode *string `json:"DetectMode,omitnil" name:"DetectMode"`

	// 会话速率和周期特征校验开关，默认off，取值有：
	// <li>off：关闭；</li>
	// <li>on：打开。</li>
	SessionAnalyzeSwitch *string `json:"SessionAnalyzeSwitch,omitnil" name:"SessionAnalyzeSwitch"`

	// 校验结果为未携带Cookie或Cookie已过期的统计周期。单位为秒，默认10，取值：5～3600。
	InvalidStatTime *int64 `json:"InvalidStatTime,omitnil" name:"InvalidStatTime"`

	// 校验结果为未携带Cookie或Cookie已过期的触发阈值。单位为次，默认300，取值：1～100000000。
	InvalidThreshold *int64 `json:"InvalidThreshold,omitnil" name:"InvalidThreshold"`

	// Cookie校验校验结果。
	AlgDetectResults []*AlgDetectResult `json:"AlgDetectResults,omitnil" name:"AlgDetectResults"`

	// 会话速率和周期特征校验结果。
	SessionBehaviors []*AlgDetectResult `json:"SessionBehaviors,omitnil" name:"SessionBehaviors"`
}

type AliasDomain struct {
	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil" name:"TargetName"`

	// 别称域名状态，取值有：
	// <li> active：已生效； </li>
	// <li> pending：部署中；</li>
	// <li> conflict：被找回。 </li>
	// <li> stop：已停用；</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 封禁模式，取值有：
	// <li> 0：未封禁； </li>
	// <li> 11：合规封禁；</li>
	// <li> 14：未备案封禁。</li>
	ForbidMode *int64 `json:"ForbidMode,omitnil" name:"ForbidMode"`

	// 别称域名创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 别称域名修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil" name:"ModifiedOn"`
}

type ApplicationProxy struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 当ProxyType=hostname时，表示域名或子域名；
	// 当ProxyType=instance时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil" name:"ProxyName"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>
	ProxyType *string `json:"ProxyType,omitnil" name:"ProxyType"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitnil" name:"PlatType"`

	// 加速区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	// 默认值：overseas
	Area *string `json:"Area,omitnil" name:"Area"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitnil" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil" name:"AccelerateType"`

	// 会话保持时间。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 封禁状态，取值有：
	// <li>banned：已封禁;</li>
	// <li>banning：封禁中；</li>
	// <li>recover：已解封；</li>
	// <li>recovering：解封禁中。</li>
	BanStatus *string `json:"BanStatus,omitnil" name:"BanStatus"`

	// 调度信息。
	ScheduleValue []*string `json:"ScheduleValue,omitnil" name:"ScheduleValue"`

	// 当ProxyType=hostname时：
	// 表示代理加速唯一标识。
	HostId *string `json:"HostId,omitnil" name:"HostId"`

	// Ipv6访问配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 规则列表。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil" name:"ApplicationProxyRules"`

	// 中国大陆加速优化配置。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil" name:"AccelerateMainland"`
}

type ApplicationProxyRule struct {
	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 端口，支持格式：
	// <li>单个端口，如：80。</li>
	// <li>端口段，如：81-82。表示81，82两个端口。</li>
	// 注意：一条规则最多可填写20个端口。
	Port []*string `json:"Port,omitnil" name:"Port"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	OriginValue []*string `json:"OriginValue,omitnil" name:"OriginValue"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用；</li>
	// <li>progress：部署中；</li>
	// <li>stopping：停用中；</li>
	// <li>fail：部署失败/停用失败。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitnil" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口，如：80。</li>
	// <li>端口段：81-82，表示81，82两个端口。</li>
	OriginPort *string `json:"OriginPort,omitnil" name:"OriginPort"`

	// 规则标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTag *string `json:"RuleTag,omitnil" name:"RuleTag"`
}

type AscriptionInfo struct {
	// 主机记录。
	Subdomain *string `json:"Subdomain,omitnil" name:"Subdomain"`

	// 记录类型。
	RecordType *string `json:"RecordType,omitnil" name:"RecordType"`

	// 记录值。
	RecordValue *string `json:"RecordValue,omitnil" name:"RecordValue"`
}

// Predefined struct for user
type BindZoneToPlanRequestParams struct {
	// 未绑定套餐的站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 待绑定的目标套餐ID。
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

type BindZoneToPlanRequest struct {
	*tchttp.BaseRequest
	
	// 未绑定套餐的站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 待绑定的目标套餐ID。
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`
}

func (r *BindZoneToPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindZoneToPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindZoneToPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindZoneToPlanResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type BindZoneToPlanResponse struct {
	*tchttp.BaseResponse
	Response *BindZoneToPlanResponseParams `json:"Response"`
}

func (r *BindZoneToPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindZoneToPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BotConfig struct {
	// bot开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 通用详细基础规则。如果为null，默认使用历史配置。
	BotManagedRule *BotManagedRule `json:"BotManagedRule,omitnil" name:"BotManagedRule"`

	// 用户画像规则。如果为null，默认使用历史配置。
	BotPortraitRule *BotPortraitRule `json:"BotPortraitRule,omitnil" name:"BotPortraitRule"`

	// Bot智能分析。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntelligenceRule *IntelligenceRule `json:"IntelligenceRule,omitnil" name:"IntelligenceRule"`

	// Bot自定义规则。如果为null，默认使用历史配置。
	BotUserRules []*BotUserRule `json:"BotUserRules,omitnil" name:"BotUserRules"`

	// Bot主动特征识别规则。
	AlgDetectRule []*AlgDetectRule `json:"AlgDetectRule,omitnil" name:"AlgDetectRule"`

	// Bot托管定制策略，入参可不填，仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Customizes []*BotUserRule `json:"Customizes,omitnil" name:"Customizes"`
}

type BotExtendAction struct {
	// 处置动作，取值有：
	// <li>monitor：观察；</li>
	// <li>alg：JavaScript挑战；</li>
	// <li>captcha：托管挑战；</li>
	// <li>random：随机，按照ExtendActions分配处置动作和比例；</li>
	// <li>silence：静默；</li>
	// <li>shortdelay：短时响应；</li>
	// <li>longdelay：长时响应。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 处置方式的触发概率，范围0-100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitnil" name:"Percent"`
}

type BotManagedRule struct {
	// 触发规则后的处置方式，取值有：
	// <li>drop：拦截；</li>
	// <li>trans：放行；</li>
	// <li>alg：Javascript挑战；</li>
	// <li>monitor：观察。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 本规则的ID。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 放行的规则ID。默认所有规则不配置放行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransManagedIds []*int64 `json:"TransManagedIds,omitnil" name:"TransManagedIds"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil" name:"DropManagedIds"`
}

type BotPortraitRule struct {
	// 本功能的开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 本规则的ID。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// JS挑战的规则ID。默认所有规则不配置JS挑战。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgManagedIds []*int64 `json:"AlgManagedIds,omitnil" name:"AlgManagedIds"`

	// 数字验证码的规则ID。默认所有规则不配置数字验证码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CapManagedIds []*int64 `json:"CapManagedIds,omitnil" name:"CapManagedIds"`

	// 观察的规则ID。默认所有规则不配置观察。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonManagedIds []*int64 `json:"MonManagedIds,omitnil" name:"MonManagedIds"`

	// 拦截的规则ID。默认所有规则不配置拦截。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropManagedIds []*int64 `json:"DropManagedIds,omitnil" name:"DropManagedIds"`
}

type BotUserRule struct {
	// 规则名，只能以英文字符，数字，下划线组合，且不能以下划线开头。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 处置动作，取值有：
	// <li>drop：拦截；</li>
	// <li>monitor：观察；</li>
	// <li>trans：放行；</li>
	// <li>alg：JavaScript挑战；</li>
	// <li>captcha：托管挑战；</li>
	// <li>random：随机处置；</li>
	// <li>silence：静默；</li>
	// <li>shortdelay：短时响应；</li>
	// <li>longdelay：长时响应。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：不生效。</li>默认on生效。
	RuleStatus *string `json:"RuleStatus,omitnil" name:"RuleStatus"`

	// 规则详情。
	AclConditions []*AclCondition `json:"AclConditions,omitnil" name:"AclConditions"`

	// 规则权重，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitnil" name:"RulePriority"`

	// 规则id。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 随机处置的处置方式及占比，非随机处置可不填暂不支持。
	ExtendActions []*BotExtendAction `json:"ExtendActions,omitnil" name:"ExtendActions"`

	// 过滤词，取值有：
	// <li>sip：客户端ip。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreqFields []*string `json:"FreqFields,omitnil" name:"FreqFields"`

	// 更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 统计范围，字段为null时，代表source_to_eo。取值有：
	// <li>source_to_eo：（响应）源站到EdgeOne。</li>
	// <li>client_to_eo：（请求）客户端到EdgeOne；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreqScope []*string `json:"FreqScope,omitnil" name:"FreqScope"`
}

type CC struct {
	// Waf开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

type Cache struct {
	// 缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 缓存过期时间设置。
	// 单位为秒，最大可设置为 365 天。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheTime *int64 `json:"CacheTime,omitnil" name:"CacheTime"`

	// 是否开启强制缓存，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IgnoreCacheControl is deprecated.
	IgnoreCacheControl *string `json:"IgnoreCacheControl,omitnil" name:"IgnoreCacheControl"`
}

type CacheConfig struct {
	// 缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cache *Cache `json:"Cache,omitnil" name:"Cache"`

	// 不缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCache *NoCache `json:"NoCache,omitnil" name:"NoCache"`

	// 遵循源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FollowOrigin *FollowOrigin `json:"FollowOrigin,omitnil" name:"FollowOrigin"`
}

type CacheKey struct {
	// 是否开启全路径缓存，取值有：
	// <li>on：开启全路径缓存（即关闭参数忽略）；</li>
	// <li>off：关闭全路径缓存（即开启参数忽略）。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullUrlCache *string `json:"FullUrlCache,omitnil" name:"FullUrlCache"`

	// 是否忽略大小写缓存，取值有：
	// <li>on：忽略；</li>
	// <li>off：不忽略。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreCase *string `json:"IgnoreCase,omitnil" name:"IgnoreCase"`

	// CacheKey 中包含请求参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryString *QueryString `json:"QueryString,omitnil" name:"QueryString"`
}

type CachePrefresh struct {
	// 缓存预刷新配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 缓存预刷新百分比，取值范围：1-99。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *int64 `json:"Percent,omitnil" name:"Percent"`
}

// Predefined struct for user
type CheckCnameStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速域名列表。
	RecordNames []*string `json:"RecordNames,omitnil" name:"RecordNames"`
}

type CheckCnameStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速域名列表。
	RecordNames []*string `json:"RecordNames,omitnil" name:"RecordNames"`
}

func (r *CheckCnameStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCnameStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RecordNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckCnameStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckCnameStatusResponseParams struct {
	// 加速域名 CNAME 状态信息列表。
	CnameStatus []*CnameStatus `json:"CnameStatus,omitnil" name:"CnameStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CheckCnameStatusResponse struct {
	*tchttp.BaseResponse
	Response *CheckCnameStatusResponseParams `json:"Response"`
}

func (r *CheckCnameStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckCnameStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientIpCountry struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 存放客户端 IP 所属地域信息的请求头名称，当 Switch=on 时有效。
	// 为空则使用默认值：EO-Client-IPCountry。
	HeaderName *string `json:"HeaderName,omitnil" name:"HeaderName"`
}

type ClientIpHeader struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 回源时，存放客户端 IP 的请求头名称。
	// 为空则使用默认值：X-Forwarded-IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderName *string `json:"HeaderName,omitnil" name:"HeaderName"`
}

type CnameStatus struct {
	// 记录名称。
	RecordName *string `json:"RecordName,omitnil" name:"RecordName"`

	// CNAME 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// Cname状态信息，取值有：
	// <li>active：生效；</li>
	// <li>moved：不生效。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type CodeAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 操作参数。
	Parameters []*RuleCodeActionParams `json:"Parameters,omitnil" name:"Parameters"`
}

type Compression struct {
	// 智能压缩配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 支持的压缩算法列表，取值有：
	// <li>brotli：brotli算法；</li>
	// <li>gzip：gzip算法。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Algorithms []*string `json:"Algorithms,omitnil" name:"Algorithms"`
}

// Predefined struct for user
type CreateAccelerationDomainRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil" name:"OriginInfo"`
}

type CreateAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil" name:"OriginInfo"`
}

func (r *CreateAccelerationDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerationDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainName")
	delete(f, "OriginInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccelerationDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAccelerationDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccelerationDomainResponseParams `json:"Response"`
}

func (r *CreateAccelerationDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccelerationDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAliasDomainRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书。</li>默认取值为 none。
	CertType *string `json:"CertType,omitnil" name:"CertType"`

	// 当 CertType 取值为 hosting 时需填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil" name:"CertId"`
}

type CreateAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书。</li>默认取值为 none。
	CertType *string `json:"CertType,omitnil" name:"CertType"`

	// 当 CertType 取值为 hosting 时需填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil" name:"CertId"`
}

func (r *CreateAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasName")
	delete(f, "TargetName")
	delete(f, "CertType")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAliasDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAliasDomainResponseParams `json:"Response"`
}

func (r *CreateAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil" name:"ProxyName"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitnil" name:"PlatType"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitnil" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil" name:"AccelerateType"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitnil" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// Ipv6 访问配置。
	// 不填写表示关闭 Ipv6 访问。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 规则详细信息。
	// 不填写则不创建规则。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil" name:"ApplicationProxyRules"`

	// 中国大陆加速优化配置。不填写表示关闭中国大陆加速优化。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil" name:"AccelerateMainland"`
}

type CreateApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil" name:"ProxyName"`

	// 调度模式，取值有：
	// <li>ip：表示Anycast IP调度；</li>
	// <li>domain：表示CNAME调度。</li>
	PlatType *string `json:"PlatType,omitnil" name:"PlatType"`

	// 是否开启安全，取值有：
	// <li>0：关闭安全；</li>
	// <li>1：开启安全。</li>
	SecurityType *int64 `json:"SecurityType,omitnil" name:"SecurityType"`

	// 是否开启加速，取值有：
	// <li>0：关闭加速；</li>
	// <li>1：开启加速。</li>
	AccelerateType *int64 `json:"AccelerateType,omitnil" name:"AccelerateType"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写使用默认值instance。
	ProxyType *string `json:"ProxyType,omitnil" name:"ProxyType"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写使用默认值600。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// Ipv6 访问配置。
	// 不填写表示关闭 Ipv6 访问。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 规则详细信息。
	// 不填写则不创建规则。
	ApplicationProxyRules []*ApplicationProxyRule `json:"ApplicationProxyRules,omitnil" name:"ApplicationProxyRules"`

	// 中国大陆加速优化配置。不填写表示关闭中国大陆加速优化。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil" name:"AccelerateMainland"`
}

func (r *CreateApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyName")
	delete(f, "PlatType")
	delete(f, "SecurityType")
	delete(f, "AccelerateType")
	delete(f, "ProxyType")
	delete(f, "SessionPersistTime")
	delete(f, "Ipv6")
	delete(f, "ApplicationProxyRules")
	delete(f, "AccelerateMainland")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyResponseParams struct {
	// 新增的四层代理应用ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyResponseParams `json:"Response"`
}

func (r *CreateApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRuleRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil" name:"Port"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	OriginValue []*string `json:"OriginValue,omitnil" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitnil" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil" name:"OriginPort"`

	// 规则标签。默认值为空字符串。
	RuleTag *string `json:"RuleTag,omitnil" name:"RuleTag"`
}

type CreateApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil" name:"Port"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	OriginValue []*string `json:"OriginValue,omitnil" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>默认值：OFF。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>默认值：false。
	SessionPersist *bool `json:"SessionPersist,omitnil" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil" name:"OriginPort"`

	// 规则标签。默认值为空字符串。
	RuleTag *string `json:"RuleTag,omitnil" name:"RuleTag"`
}

func (r *CreateApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Proto")
	delete(f, "Port")
	delete(f, "OriginType")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	delete(f, "SessionPersistTime")
	delete(f, "OriginPort")
	delete(f, "RuleTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationProxyRuleResponseParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationProxyRuleResponseParams `json:"Response"`
}

func (r *CreateApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitnil" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitnil" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitnil" name:"OriginRecords"`

	// 回源Host，仅当OriginType=self时可以设置。
	HostHeader *string `json:"HostHeader,omitnil" name:"HostHeader"`
}

type CreateOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitnil" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitnil" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitnil" name:"OriginRecords"`

	// 回源Host，仅当OriginType=self时可以设置。
	HostHeader *string `json:"HostHeader,omitnil" name:"HostHeader"`
}

func (r *CreateOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "OriginType")
	delete(f, "OriginGroupName")
	delete(f, "ConfigurationType")
	delete(f, "OriginRecords")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOriginGroupResponseParams struct {
	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil" name:"OriginGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateOriginGroupResponseParams `json:"Response"`
}

func (r *CreateOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanForZoneRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 所要购买套餐的类型，取值有：
	// <li> sta: 全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot: 全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm: 中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot: 中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> sta_global ：全球内容分发网络（包括中国大陆）标准版套餐； </li>
	// <li> sta_global_with_bot ：全球内容分发网络（包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> ent: 全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot: 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm: 中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot: 中国大陆内容分发网络企业版套餐附带bot管理。</li>
	// <li> ent_global ：全球内容分发网络（包括中国大陆）企业版套餐； </li>
	// <li> ent_global_with_bot ：全球内容分发网络（包括中国大陆）企业版套餐附带bot管理。</li>当前账户可购买套餐类型请以<a href="https://cloud.tencent.com/document/product/1552/80606">DescribeAvailablePlans</a>返回为准。
	PlanType *string `json:"PlanType,omitnil" name:"PlanType"`
}

type CreatePlanForZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 所要购买套餐的类型，取值有：
	// <li> sta: 全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot: 全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm: 中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot: 中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> sta_global ：全球内容分发网络（包括中国大陆）标准版套餐； </li>
	// <li> sta_global_with_bot ：全球内容分发网络（包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> ent: 全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot: 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm: 中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot: 中国大陆内容分发网络企业版套餐附带bot管理。</li>
	// <li> ent_global ：全球内容分发网络（包括中国大陆）企业版套餐； </li>
	// <li> ent_global_with_bot ：全球内容分发网络（包括中国大陆）企业版套餐附带bot管理。</li>当前账户可购买套餐类型请以<a href="https://cloud.tencent.com/document/product/1552/80606">DescribeAvailablePlans</a>返回为准。
	PlanType *string `json:"PlanType,omitnil" name:"PlanType"`
}

func (r *CreatePlanForZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanForZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "PlanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePlanForZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePlanForZoneResponseParams struct {
	// 购买的资源名字列表。
	ResourceNames []*string `json:"ResourceNames,omitnil" name:"ResourceNames"`

	// 购买的订单号列表。
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePlanForZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreatePlanForZoneResponseParams `json:"Response"`
}

func (r *CreatePlanForZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePlanForZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt。
	// 注意：提交任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。
	Targets []*string `json:"Targets,omitnil" name:"Targets"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitnil" name:"EncodeUrl"`

	// 附带的http头部信息。
	Headers []*Header `json:"Headers,omitnil" name:"Headers"`
}

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt。
	// 注意：提交任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。
	Targets []*string `json:"Targets,omitnil" name:"Targets"`

	// 是否对url进行encode，若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitnil" name:"EncodeUrl"`

	// 附带的http头部信息。
	Headers []*Header `json:"Headers,omitnil" name:"Headers"`
}

func (r *CreatePrefetchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrefetchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePrefetchTaskResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 失败的任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*FailReason `json:"FailedList,omitnil" name:"FailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePrefetchTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePrefetchTaskResponseParams `json:"Response"`
}

func (r *CreatePrefetchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 节点缓存清除类型，取值有：
	// <li>purge_url：URL刷新；</li>
	// <li>purge_prefix：目录刷新；</li>
	// <li>purge_host：Hostname 刷新；</li>
	// <li>purge_all：站点下全部缓存刷新；</li>
	// <li>purge_cache_tag：cache-tag 刷新。</li>缓存清除类型详情请查看[清除缓存](https://cloud.tencent.com/document/product/1552/70759)。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 节点缓存清除方法，针对目录刷新、Hostname刷新以及刷新全部缓存 类型有效，取值有：<li> invalidate：仅刷新目录下产生了更新的资源；</li><li> delete：无论目录下资源是否更新都刷新节点资源。</li>注意：使用目录刷新时，默认值： invalidate。
	Method *string `json:"Method,omitnil" name:"Method"`

	// 要清除缓存的资源列表。每个元素格式依据清除缓存类型而定，可参考接口示例。<li>EO 默认针对内容含有非 ASCII 字符集的字符进行转义，编码规则遵循 RFC3986；</li><li>单次提交的任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。</li>
	Targets []*string `json:"Targets,omitnil" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源。
	// 若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	//
	// Deprecated: EncodeUrl is deprecated.
	EncodeUrl *bool `json:"EncodeUrl,omitnil" name:"EncodeUrl"`
}

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 节点缓存清除类型，取值有：
	// <li>purge_url：URL刷新；</li>
	// <li>purge_prefix：目录刷新；</li>
	// <li>purge_host：Hostname 刷新；</li>
	// <li>purge_all：站点下全部缓存刷新；</li>
	// <li>purge_cache_tag：cache-tag 刷新。</li>缓存清除类型详情请查看[清除缓存](https://cloud.tencent.com/document/product/1552/70759)。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 节点缓存清除方法，针对目录刷新、Hostname刷新以及刷新全部缓存 类型有效，取值有：<li> invalidate：仅刷新目录下产生了更新的资源；</li><li> delete：无论目录下资源是否更新都刷新节点资源。</li>注意：使用目录刷新时，默认值： invalidate。
	Method *string `json:"Method,omitnil" name:"Method"`

	// 要清除缓存的资源列表。每个元素格式依据清除缓存类型而定，可参考接口示例。<li>EO 默认针对内容含有非 ASCII 字符集的字符进行转义，编码规则遵循 RFC3986；</li><li>单次提交的任务数受计费套餐配额限制，请查看 [EO计费套餐](https://cloud.tencent.com/document/product/1552/77380)。</li>
	Targets []*string `json:"Targets,omitnil" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源。
	// 若内容含有非 ASCII 字符集的字符，请开启此开关进行编码转换（编码规则遵循 RFC3986）。
	EncodeUrl *bool `json:"EncodeUrl,omitnil" name:"EncodeUrl"`
}

func (r *CreatePurgeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Method")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePurgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePurgeTaskResponseParams struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 失败的任务列表及原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedList []*FailReason `json:"FailedList,omitnil" name:"FailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePurgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreatePurgeTaskResponseParams `json:"Response"`
}

func (r *CreatePurgeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 规则名称，名称字符串长度 1～255。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil" name:"Rules"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 规则名称，名称字符串长度 1～255。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil" name:"Rules"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleName")
	delete(f, "Status")
	delete(f, "Rules")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityIPGroupRequestParams struct {
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// IP 组信息。
	IPGroup *IPGroup `json:"IPGroup,omitnil" name:"IPGroup"`
}

type CreateSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// IP 组信息。
	IPGroup *IPGroup `json:"IPGroup,omitnil" name:"IPGroup"`
}

func (r *CreateSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "IPGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityIPGroupResponseParams struct {
	// IP 组 Id。
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityIPGroupResponseParams `json:"Response"`
}

func (r *CreateSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSharedCNAMERequestParams struct {
	// 共享 CNAME 所属站点的 ID。	
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 共享 CNAME 前缀。请输入合法的域名前缀，例如"test-api"、"test-api.com"，限制输入 50 个字符。
	// 共享 CNAME 完整格式为：<自定义前缀>+<zoneid中的12位随机字符串>+"share.eo.dnse[0-5].com"。例如前缀传入 example.com，EO 会为您创建共享 CNAME：example.com.sai2ig51kaa5.eo.dns2.com
	// 示例值：example.com
	SharedCNAMEPrefix *string `json:"SharedCNAMEPrefix,omitnil" name:"SharedCNAMEPrefix"`

	// 描述。可输入 1-50 个任意字符。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateSharedCNAMERequest struct {
	*tchttp.BaseRequest
	
	// 共享 CNAME 所属站点的 ID。	
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 共享 CNAME 前缀。请输入合法的域名前缀，例如"test-api"、"test-api.com"，限制输入 50 个字符。
	// 共享 CNAME 完整格式为：<自定义前缀>+<zoneid中的12位随机字符串>+"share.eo.dnse[0-5].com"。例如前缀传入 example.com，EO 会为您创建共享 CNAME：example.com.sai2ig51kaa5.eo.dns2.com
	// 示例值：example.com
	SharedCNAMEPrefix *string `json:"SharedCNAMEPrefix,omitnil" name:"SharedCNAMEPrefix"`

	// 描述。可输入 1-50 个任意字符。
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *CreateSharedCNAMERequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedCNAMERequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SharedCNAMEPrefix")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSharedCNAMERequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSharedCNAMEResponseParams struct {
	// 共享 CNAME。格式为：<自定义前缀>+<ZoneId中的12位随机字符串>+"share.eo.dnse[0-5].com"
	SharedCNAME *string `json:"SharedCNAME,omitnil" name:"SharedCNAME"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSharedCNAMEResponse struct {
	*tchttp.BaseResponse
	Response *CreateSharedCNAMEResponseParams `json:"Response"`
}

func (r *CreateSharedCNAMEResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedCNAMEResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneRequestParams struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 接入方式，取值有：
	// <li> full：NS接入；</li>
	// <li> partial：CNAME接入，请先调用认证站点API（IdentifyZone）进行站点归属权校验，校验通过后继续调用本接口创建站点；<li>noDomainAccess：无域名接入，取此值时仅Tags字段有效。</li>
	// </li>不填写使用默认值full。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 是否跳过站点现有的DNS记录扫描。默认值：false。
	JumpStart *bool `json:"JumpStart,omitnil" name:"JumpStart"`

	// 资源标签。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 是否允许重复接入。
	// <li> true：允许重复接入；</li>
	// <li> false：不允许重复接入。</li>不填写使用默认值false。
	AllowDuplicates *bool `json:"AllowDuplicates,omitnil" name:"AllowDuplicates"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	AliasZoneName *string `json:"AliasZoneName,omitnil" name:"AliasZoneName"`
}

type CreateZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 接入方式，取值有：
	// <li> full：NS接入；</li>
	// <li> partial：CNAME接入，请先调用认证站点API（IdentifyZone）进行站点归属权校验，校验通过后继续调用本接口创建站点；<li>noDomainAccess：无域名接入，取此值时仅Tags字段有效。</li>
	// </li>不填写使用默认值full。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 是否跳过站点现有的DNS记录扫描。默认值：false。
	JumpStart *bool `json:"JumpStart,omitnil" name:"JumpStart"`

	// 资源标签。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 是否允许重复接入。
	// <li> true：允许重复接入；</li>
	// <li> false：不允许重复接入。</li>不填写使用默认值false。
	AllowDuplicates *bool `json:"AllowDuplicates,omitnil" name:"AllowDuplicates"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	AliasZoneName *string `json:"AliasZoneName,omitnil" name:"AliasZoneName"`
}

func (r *CreateZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneName")
	delete(f, "Type")
	delete(f, "JumpStart")
	delete(f, "Tags")
	delete(f, "AllowDuplicates")
	delete(f, "AliasZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateZoneResponseParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateZoneResponse struct {
	*tchttp.BaseResponse
	Response *CreateZoneResponseParams `json:"Response"`
}

func (r *CreateZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DDoS struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type DDoSAttackEvent struct {
	// 事件ID。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 攻击类型(对应交互事件名称)。
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 攻击状态。
	AttackStatus *int64 `json:"AttackStatus,omitnil" name:"AttackStatus"`

	// 攻击最大带宽。
	AttackMaxBandWidth *int64 `json:"AttackMaxBandWidth,omitnil" name:"AttackMaxBandWidth"`

	// 攻击包速率峰值。
	AttackPacketMaxRate *int64 `json:"AttackPacketMaxRate,omitnil" name:"AttackPacketMaxRate"`

	// 攻击开始时间，单位为s。
	AttackStartTime *int64 `json:"AttackStartTime,omitnil" name:"AttackStartTime"`

	// 攻击结束时间，单位为s。
	AttackEndTime *int64 `json:"AttackEndTime,omitnil" name:"AttackEndTime"`

	// DDoS策略组ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`

	// 站点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 攻击事件所属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 封禁解封信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoSBlockData []*DDoSBlockData `json:"DDoSBlockData,omitnil" name:"DDoSBlockData"`
}

type DDoSBlockData struct {
	// 开始时间，采用unix时间戳。
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间，采用unix时间戳, 为0表示还处于封禁中。
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 封禁受影响区域。
	BlockArea *string `json:"BlockArea,omitnil" name:"BlockArea"`
}

type DefaultServerCertInfo struct {
	// 服务器证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// 证书备注名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 证书类型，取值有：
	// <li>default: 默认证书;</li>
	// <li>upload:用户上传;</li>
	// <li>managed:腾讯云托管。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 证书生效时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EffectiveTime *string `json:"EffectiveTime,omitnil" name:"EffectiveTime"`

	// 证书公用名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitnil" name:"CommonName"`

	// 证书SAN域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubjectAltName []*string `json:"SubjectAltName,omitnil" name:"SubjectAltName"`

	// 部署状态，取值有：
	// <li>processing: 部署中；</li>
	// <li>deployed: 已部署；</li>
	// <li>failed: 部署失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// Status为失败时,此字段返回失败原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 证书算法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignAlgo *string `json:"SignAlgo,omitnil" name:"SignAlgo"`
}

// Predefined struct for user
type DeleteAccelerationDomainsRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 需要删除的加速域名ID列表。
	DomainNames []*string `json:"DomainNames,omitnil" name:"DomainNames"`

	// 是否强制删除。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制删除该域名，取值有：
	// <li> true：删除该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许删除。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil" name:"Force"`
}

type DeleteAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 需要删除的加速域名ID列表。
	DomainNames []*string `json:"DomainNames,omitnil" name:"DomainNames"`

	// 是否强制删除。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制删除该域名，取值有：
	// <li> true：删除该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许删除。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil" name:"Force"`
}

func (r *DeleteAccelerationDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerationDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainNames")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccelerationDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccelerationDomainsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAccelerationDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccelerationDomainsResponseParams `json:"Response"`
}

func (r *DeleteAccelerationDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccelerationDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAliasDomainRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 待删除别称域名名称。如果为空，则不执行删除操作。
	AliasNames []*string `json:"AliasNames,omitnil" name:"AliasNames"`
}

type DeleteAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 待删除别称域名名称。如果为空，则不执行删除操作。
	AliasNames []*string `json:"AliasNames,omitnil" name:"AliasNames"`
}

func (r *DeleteAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAliasDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAliasDomainResponseParams `json:"Response"`
}

func (r *DeleteAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`
}

type DeleteApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`
}

func (r *DeleteApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationProxyResponseParams `json:"Response"`
}

func (r *DeleteApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyRuleRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

type DeleteApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *DeleteApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationProxyRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationProxyRuleResponseParams `json:"Response"`
}

func (r *DeleteApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil" name:"OriginGroupId"`
}

type DeleteOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil" name:"OriginGroupId"`
}

func (r *DeleteOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "OriginGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOriginGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOriginGroupResponseParams `json:"Response"`
}

func (r *DeleteOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 指定删除的规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitnil" name:"RuleIds"`
}

type DeleteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 指定删除的规则 ID 列表。
	RuleIds []*string `json:"RuleIds,omitnil" name:"RuleIds"`
}

func (r *DeleteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRulesResponseParams `json:"Response"`
}

func (r *DeleteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityIPGroupRequestParams struct {
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// IP 组 Id。
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

type DeleteSecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// IP 组 Id。
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *DeleteSecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityIPGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityIPGroupResponseParams `json:"Response"`
}

func (r *DeleteSecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
}

type DeleteZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
}

func (r *DeleteZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteZoneResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteZoneResponse struct {
	*tchttp.BaseResponse
	Response *DeleteZoneResponseParams `json:"Response"`
}

func (r *DeleteZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerationDomainsRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>domain-name<br>   按照【<strong>加速域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>origin-type<br>   按照【<strong>源站类型</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>origin<br>   按照【<strong>主源站地址</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>backup-origin<br>   按照【<strong>备用源站地址</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>domain-cname<br>   按照【<strong>加速CNAME名</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>share-cname<br>   按照【<strong>共享CNAME名</strong>】进行过滤。<br>   类型：String<br>   必选：否
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`

	// 列表排序方式，取值有：
	// <li>asc：升序排列；</li>
	// <li>desc：降序排列。</li>默认值为asc。
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 匹配方式，取值有：
	// <li>all：返回匹配所有查询条件的加速域名；</li>
	// <li>any：返回匹配任意一个查询条件的加速域名。</li>默认值为all。
	Match *string `json:"Match,omitnil" name:"Match"`

	// 分页查询限制数目，默认值：20，上限：200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序依据，取值有：
	// <li>created_on：加速域名创建时间；</li>
	// <li>domain-name：加速域名名称；</li>
	// </li>默认根据domain-name属性排序。
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeAccelerationDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>domain-name<br>   按照【<strong>加速域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>origin-type<br>   按照【<strong>源站类型</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>origin<br>   按照【<strong>主源站地址</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>backup-origin<br>   按照【<strong>备用源站地址</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>domain-cname<br>   按照【<strong>加速CNAME名</strong>】进行过滤。<br>   类型：String<br>   必选：否
	// <li>share-cname<br>   按照【<strong>共享CNAME名</strong>】进行过滤。<br>   类型：String<br>   必选：否
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`

	// 列表排序方式，取值有：
	// <li>asc：升序排列；</li>
	// <li>desc：降序排列。</li>默认值为asc。
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 匹配方式，取值有：
	// <li>all：返回匹配所有查询条件的加速域名；</li>
	// <li>any：返回匹配任意一个查询条件的加速域名。</li>默认值为all。
	Match *string `json:"Match,omitnil" name:"Match"`

	// 分页查询限制数目，默认值：20，上限：200。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序依据，取值有：
	// <li>created_on：加速域名创建时间；</li>
	// <li>domain-name：加速域名名称；</li>
	// </li>默认根据domain-name属性排序。
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *DescribeAccelerationDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerationDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	delete(f, "Direction")
	delete(f, "Match")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccelerationDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccelerationDomainsResponseParams struct {
	// 加速域名总数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 加速域名列表。
	AccelerationDomains []*AccelerationDomain `json:"AccelerationDomains,omitnil" name:"AccelerationDomains"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccelerationDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccelerationDomainsResponseParams `json:"Response"`
}

func (r *DescribeAccelerationDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccelerationDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasDomainsRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>target-name<br>   按照【<strong>目标域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>alias-name<br>   按照【<strong>别称域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>模糊查询时仅支持过滤字段名为alias-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAliasDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>target-name<br>   按照【<strong>目标域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>alias-name<br>   按照【<strong>别称域名名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>模糊查询时仅支持过滤字段名为alias-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeAliasDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAliasDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAliasDomainsResponseParams struct {
	// 符合条件的别称域名个数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 别称域名详细信息列表。
	AliasDomains []*AliasDomain `json:"AliasDomains,omitnil" name:"AliasDomains"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAliasDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAliasDomainsResponseParams `json:"Response"`
}

func (r *DescribeAliasDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAliasDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesRequestParams struct {
	// 分页查询偏移量。默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>proxy-id<br>   按照【<strong>代理ID</strong>】进行过滤。代理ID形如：proxy-ev2sawbwfd。<br>   类型：String<br>   必选：否</li><li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-vawer2vadg。<br>   类型：String<br>   必选：否</li><li>rule-tag<br>   按照【<strong>规则标签</strong>】对应用代理下的规则进行过滤。规则标签形如：rule-service-1。<br>   类型：String<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeApplicationProxiesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量。默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>proxy-id<br>   按照【<strong>代理ID</strong>】进行过滤。代理ID形如：proxy-ev2sawbwfd。<br>   类型：String<br>   必选：否</li><li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-vawer2vadg。<br>   类型：String<br>   必选：否</li><li>rule-tag<br>   按照【<strong>规则标签</strong>】对应用代理下的规则进行过滤。规则标签形如：rule-service-1。<br>   类型：String<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeApplicationProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationProxiesResponseParams struct {
	// 应用代理列表。
	ApplicationProxies []*ApplicationProxy `json:"ApplicationProxies,omitnil" name:"ApplicationProxies"`

	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeApplicationProxiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationProxiesResponseParams `json:"Response"`
}

func (r *DescribeApplicationProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailablePlansRequestParams struct {

}

type DescribeAvailablePlansRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAvailablePlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailablePlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAvailablePlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAvailablePlansResponseParams struct {
	// 当前账户可购买套餐类型及相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanInfo []*PlanInfo `json:"PlanInfo,omitnil" name:"PlanInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAvailablePlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAvailablePlansResponseParams `json:"Response"`
}

func (r *DescribeAvailablePlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAvailablePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
}

type DescribeContentQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
}

func (r *DescribeContentQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeContentQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeContentQuotaResponseParams struct {
	// 刷新相关配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PurgeQuota []*Quota `json:"PurgeQuota,omitnil" name:"PurgeQuota"`

	// 预热相关配额。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrefetchQuota []*Quota `json:"PrefetchQuota,omitnil" name:"PrefetchQuota"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeContentQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeContentQuotaResponseParams `json:"Response"`
}

func (r *DescribeContentQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeContentQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// DDoS策略组ID列表，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间与结束时间的间隔自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeDDoSAttackDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 统计指标列表，取值有：
	// <li>ddos_attackMaxBandwidth：攻击带宽峰值；</li>
	// <li>ddos_attackMaxPackageRate：攻击包速率峰值 ；</li>
	// <li>ddos_attackBandwidth：攻击带宽曲线；</li>
	// <li>ddos_attackPackageRate：攻击包速率曲线。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// DDoS策略组ID列表，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间与结束时间的间隔自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeDDoSAttackDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// DDoS攻击数据内容列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SecEntry `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDDoSAttackDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackDataResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 是否展示详细信息。
	ShowDetail *bool `json:"ShowDetail,omitnil" name:"ShowDetail"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据；</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 排序字段，取值有：
	// <li>MaxBandWidth：带宽峰值；</li>
	// <li>AttackStartTime：攻击开始时间。</li>不填默认值为：AttackStartTime。
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 排序方式，取值有：
	// <li>asc：升序方式；</li>
	// <li>desc：降序方式。</li>不填默认值为：desc。
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

type DescribeDDoSAttackEventRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// ddos策略组集合，不填默认选择全部策略。
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 分页查询的限制数目，默认值为20，最大查询条目为1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 是否展示详细信息。
	ShowDetail *bool `json:"ShowDetail,omitnil" name:"ShowDetail"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据；</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 排序字段，取值有：
	// <li>MaxBandWidth：带宽峰值；</li>
	// <li>AttackStartTime：攻击开始时间。</li>不填默认值为：AttackStartTime。
	OrderBy *string `json:"OrderBy,omitnil" name:"OrderBy"`

	// 排序方式，取值有：
	// <li>asc：升序方式；</li>
	// <li>desc：降序方式。</li>不填默认值为：desc。
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

func (r *DescribeDDoSAttackEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PolicyIds")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ShowDetail")
	delete(f, "Area")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackEventResponseParams struct {
	// DDOS攻击事件数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DDoSAttackEvent `json:"Data,omitnil" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDDoSAttackEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackEventResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackTopDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：按各协议的攻击流量排行；</li>
	// <li>ddos_attackPackageNum_protocol：按各协议的攻击包量排行；</li>
	// <li>ddos_attackNum_attackType：按各攻击类型的攻击数量排行；</li>
	// <li>ddos_attackNum_sregion：按攻击源地区的攻击数量排行；</li>
	// <li>ddos_attackFlux_sip：按攻击源IP的攻击数量排行；</li>
	// <li>ddos_attackFlux_sregion：按攻击源地区的攻击数量排行。</li>
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 站点ID集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// DDoS策略组ID集合，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 攻击类型，取值有：
	// <li>flood：洪泛攻击；</li>
	// <li>icmpFlood：icmp洪泛攻击；</li>
	// <li>all：所有的攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：所有的协议类型。</li>不填默认为all，表示查询所有协议。
	ProtocolType *string `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 端口号。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeDDoSAttackTopDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的统计指标，取值有：
	// <li>ddos_attackFlux_protocol：按各协议的攻击流量排行；</li>
	// <li>ddos_attackPackageNum_protocol：按各协议的攻击包量排行；</li>
	// <li>ddos_attackNum_attackType：按各攻击类型的攻击数量排行；</li>
	// <li>ddos_attackNum_sregion：按攻击源地区的攻击数量排行；</li>
	// <li>ddos_attackFlux_sip：按攻击源IP的攻击数量排行；</li>
	// <li>ddos_attackFlux_sregion：按攻击源地区的攻击数量排行。</li>
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 站点ID集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// DDoS策略组ID集合，不填默认选择全部策略ID。
	PolicyIds []*int64 `json:"PolicyIds,omitnil" name:"PolicyIds"`

	// 攻击类型，取值有：
	// <li>flood：洪泛攻击；</li>
	// <li>icmpFlood：icmp洪泛攻击；</li>
	// <li>all：所有的攻击类型。</li>不填默认为all，表示查询全部攻击类型。
	AttackType *string `json:"AttackType,omitnil" name:"AttackType"`

	// 协议类型，取值有：
	// <li>tcp：tcp协议；</li>
	// <li>udp：udp协议；</li>
	// <li>all：所有的协议类型。</li>不填默认为all，表示查询所有协议。
	ProtocolType *string `json:"ProtocolType,omitnil" name:"ProtocolType"`

	// 端口号。
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 查询前多少个数据，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据。</li>不填将根据用户所在地智能选择地区。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeDDoSAttackTopDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackTopDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "PolicyIds")
	delete(f, "AttackType")
	delete(f, "ProtocolType")
	delete(f, "Port")
	delete(f, "Limit")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDDoSAttackTopDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDDoSAttackTopDataResponseParams struct {
	// DDoS攻击Top数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopEntry `json:"Data,omitnil" name:"Data"`

	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDDoSAttackTopDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDDoSAttackTopDataResponseParams `json:"Response"`
}

func (r *DescribeDDoSAttackTopDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDDoSAttackTopDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesRequestParams struct {
	// 过滤条件，Filters.Values的上限为5。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeDefaultCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，Filters.Values的上限为5。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：是 </li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeDefaultCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefaultCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefaultCertificatesResponseParams struct {
	// 证书总数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 默认证书列表。
	DefaultServerCertInfo []*DefaultServerCertInfo `json:"DefaultServerCertInfo,omitnil" name:"DefaultServerCertInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDefaultCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefaultCertificatesResponseParams `json:"Response"`
}

func (r *DescribeDefaultCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefaultCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 分页查询偏移量。默认值： 0，最小值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值： 100，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>host<br>   按照【<strong>域名</strong>】进行过滤。<br>   类型：string<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeHostsSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 分页查询偏移量。默认值： 0，最小值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值： 100，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>host<br>   按照【<strong>域名</strong>】进行过滤。<br>   类型：string<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeHostsSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsSettingResponseParams struct {
	// 域名列表。
	DetailHosts []*DetailHost `json:"DetailHosts,omitnil" name:"DetailHosts"`

	// 域名数量。
	TotalNumber *int64 `json:"TotalNumber,omitnil" name:"TotalNumber"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHostsSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsSettingResponseParams `json:"Response"`
}

func (r *DescribeHostsSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationsRequestParams struct {
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：是</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeIdentificationsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：是</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeIdentificationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdentificationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdentificationsResponseParams struct {
	// 符合条件的站点个数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 站点验证信息列表。
	Identifications []*Identification `json:"Identifications,omitnil" name:"Identifications"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIdentificationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdentificationsResponseParams `json:"Response"`
}

func (r *DescribeIdentificationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdentificationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupRequestParams struct {
	// 分页查询偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为10，取值：1-1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-20hzkd4rdmy0<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>origin-group-id<br>   按照【<strong>源站组ID</strong>】进行过滤。源站组ID形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>origin-group-name<br>   按照【<strong>源站组名称</strong>】进行过滤<br>   类型：String<br>   必选：否<br>   模糊查询：支持。使用模糊查询时，仅支持填写一个源站组名称</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认为10，取值：1-1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-20hzkd4rdmy0<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>origin-group-id<br>   按照【<strong>源站组ID</strong>】进行过滤。源站组ID形如：origin-2ccgtb24-7dc5-46s2-9r3e-95825d53dwe3a<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>origin-group-name<br>   按照【<strong>源站组名称</strong>】进行过滤<br>   类型：String<br>   必选：否<br>   模糊查询：支持。使用模糊查询时，仅支持填写一个源站组名称</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginGroupResponseParams struct {
	// 记录总数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 源站组信息。
	OriginGroups []*OriginGroup `json:"OriginGroups,omitnil" name:"OriginGroups"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginGroupResponseParams `json:"Response"`
}

func (r *DescribeOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginProtectionRequestParams struct {
	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>need-update<br>   按照【<strong>站点是否需要更新源站防护IP白名单</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：需要更新<br>   false：无需更新<br></li>
	// <li>plan-support<br>   按照【<strong>站点套餐是否支持源站防护</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：支持<br>   false：不支持<br></li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeOriginProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 查询的站点集合，不填默认查询所有站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>need-update<br>   按照【<strong>站点是否需要更新源站防护IP白名单</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：需要更新<br>   false：无需更新<br></li>
	// <li>plan-support<br>   按照【<strong>站点套餐是否支持源站防护</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   true：支持<br>   false：不支持<br></li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeOriginProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOriginProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOriginProtectionResponseParams struct {
	// 源站防护信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginProtectionInfo []*OriginProtectionInfo `json:"OriginProtectionInfo,omitnil" name:"OriginProtectionInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOriginProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOriginProtectionResponseParams `json:"Response"`
}

func (r *DescribeOriginProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOriginProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Flow_outFlux: Edegone响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone请求流量；</li>
	// <li>l7Flow_outBandwidth: Edegone响应带宽；</li>
	// <li>l7Flow_inBandwidth: Edegone请求带宽；</li>
	// <li>l7Flow_hit_outFlux: 缓存命中流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合。
	// 若不填写，默认选择全部站点，且最多只能查询近30天的数据；若填写，则可查询站点绑定套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询的域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 查询的协议类型，取值有：
	// <li>http: http协议；</li>
	// <li>https: https协议；</li>
	// <li>http2: http2协议；</li>
	// <li>all:  所有协议。</li>不填默认为all，此参数暂未生效。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeOverviewL7DataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Flow_outFlux: Edegone响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone请求流量；</li>
	// <li>l7Flow_outBandwidth: Edegone响应带宽；</li>
	// <li>l7Flow_inBandwidth: Edegone请求带宽；</li>
	// <li>l7Flow_hit_outFlux: 缓存命中流量；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合。
	// 若不填写，默认选择全部站点，且最多只能查询近30天的数据；若填写，则可查询站点绑定套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询的域名集合，不填默认查询所有子域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 查询的协议类型，取值有：
	// <li>http: http协议；</li>
	// <li>https: https协议；</li>
	// <li>http2: http2协议；</li>
	// <li>all:  所有协议。</li>不填默认为all，此参数暂未生效。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeOverviewL7DataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewL7DataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Protocol")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOverviewL7DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOverviewL7DataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 七层监控类时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOverviewL7DataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOverviewL7DataResponseParams `json:"Response"`
}

func (r *DescribeOverviewL7DataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOverviewL7DataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksRequestParams struct {
	// 查询起始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点 ID</strong>】进行过滤。zone-id形如：zone-1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤。target形如：http://www.qq.com/1.txt，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>domains<br>   按照【<strong>域名</strong>】进行过滤。domains形如：www.qq.com。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤。<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目，默认值：20，上限：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-id<br>   按照【<strong>站点 ID</strong>】进行过滤。zone-id形如：zone-1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤。target形如：http://www.qq.com/1.txt，暂不支持多值。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>domains<br>   按照【<strong>域名</strong>】进行过滤。domains形如：www.qq.com。<br>   类型：String<br>   必选：否。<br>   模糊查询：不支持。</li><li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤。<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribePrefetchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePrefetchTasksResponseParams struct {
	// 该查询条件总共条目数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务结果列表。
	Tasks []*Task `json:"Tasks,omitnil" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePrefetchTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePrefetchTasksResponseParams `json:"Response"`
}

func (r *DescribePrefetchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksRequestParams struct {
	// 字段已废弃，请使用Filters中的zone-id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 查询起始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查限制数目，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>zone-id<br>   按照【<strong>站点 ID</strong>】进行过滤。zone-id形如：zone-xxx，暂不支持多值<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤，target形如：http://www.qq.com/1.txt或者tag1，暂不支持多值<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>domains<br>   按照【<strong>域名</strong>】进行过滤，domains形如：www.qq.com<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。</li><li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时</li><li>type<br>   按照【<strong>清除缓存类型</strong>】进行过滤，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持<br>   可选项：<br>   purge_url：URL<br>   purge_prefix：前缀<br>   purge_all：全部缓存内容<br>   purge_host：Hostname<br>   purge_cache_tag：CacheTag</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 字段已废弃，请使用Filters中的zone-id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 查询起始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分页查询偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查限制数目，默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：<li>zone-id<br>   按照【<strong>站点 ID</strong>】进行过滤。zone-id形如：zone-xxx，暂不支持多值<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>job-id<br>   按照【<strong>任务ID</strong>】进行过滤。job-id形如：1379afjk91u32h，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>target<br>   按照【<strong>目标资源信息</strong>】进行过滤，target形如：http://www.qq.com/1.txt或者tag1，暂不支持多值<br>   类型：String<br>   必选：否<br>   模糊查询：不支持</li><li>domains<br>   按照【<strong>域名</strong>】进行过滤，domains形如：www.qq.com<br>   类型：String<br>   必选：否<br>   模糊查询：不支持。</li><li>statuses<br>   按照【<strong>任务状态</strong>】进行过滤<br>   必选：否<br>   模糊查询：不支持。<br>   可选项：<br>   processing：处理中<br>   success：成功<br>   failed：失败<br>   timeout：超时</li><li>type<br>   按照【<strong>清除缓存类型</strong>】进行过滤，暂不支持多值。<br>   类型：String<br>   必选：否<br>   模糊查询：不支持<br>   可选项：<br>   purge_url：URL<br>   purge_prefix：前缀<br>   purge_all：全部缓存内容<br>   purge_host：Hostname<br>   purge_cache_tag：CacheTag</li>
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePurgeTasksResponseParams struct {
	// 该查询条件总共条目数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务结果列表。
	Tasks []*Task `json:"Tasks,omitnil" name:"Tasks"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribePurgeTasksResponseParams `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>rule-id<br>   按照【<strong>规则ID</strong>】进行过滤。<br>   类型：string<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>rule-id<br>   按照【<strong>规则ID</strong>】进行过滤。<br>   类型：string<br>   必选：否</li>
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesResponseParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 规则列表，按规则执行顺序从先往后排序。
	RuleItems []*RuleItem `json:"RuleItems,omitnil" name:"RuleItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesResponseParams `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesSettingRequestParams struct {

}

type DescribeRulesSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeRulesSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesSettingResponseParams struct {
	// 规则引擎可应用匹配请求的设置列表及其详细建议配置信息。
	Actions []*RulesSettingAction `json:"Actions,omitnil" name:"Actions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRulesSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesSettingResponseParams `json:"Response"`
}

func (r *DescribeRulesSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量；</li>
	// <li> l4Flow_outPkt: 访问出包量。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合。
	// 若不填写，默认选择全部站点，且最多只能查询近30天的数据；
	// 若填写，则可查询站点绑定套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 四层实例列表, 不填表示选择全部实例。
	ProxyIds []*string `json:"ProxyIds,omitnil" name:"ProxyIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟 ；</li>
	// <li>5min: 5分钟 ；</li>
	// <li>hour: 1小时 ；</li>
	// <li>day: 1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>ruleId<br>   按照【<strong>转发规则ID</strong>】进行过滤。</li>
	// <li>proxyId<br>   按照【<strong>四层代理实例ID</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeTimingL4DataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询指标，取值有：
	// <li>l4Flow_connections: 访问连接数；</li>
	// <li>l4Flow_flux: 访问总流量；</li>
	// <li>l4Flow_inFlux: 访问入流量；</li>
	// <li>l4Flow_outFlux: 访问出流量；</li>
	// <li> l4Flow_outPkt: 访问出包量。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合。
	// 若不填写，默认选择全部站点，且最多只能查询近30天的数据；
	// 若填写，则可查询站点绑定套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 四层实例列表, 不填表示选择全部实例。
	ProxyIds []*string `json:"ProxyIds,omitnil" name:"ProxyIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟 ；</li>
	// <li>5min: 5分钟 ；</li>
	// <li>hour: 1小时 ；</li>
	// <li>day: 1天 。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>ruleId<br>   按照【<strong>转发规则ID</strong>】进行过滤。</li>
	// <li>proxyId<br>   按照【<strong>四层代理实例ID</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeTimingL4DataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL4DataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "ProxyIds")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL4DataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL4DataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 四层时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTimingL4DataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL4DataResponseParams `json:"Response"`
}

func (r *DescribeTimingL4DataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL4DataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: Edgeone响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone请求流量；</li>
	// <li>l7Flow_outBandwidth: Edgeone响应带宽；</li>
	// <li>l7Flow_inBandwidth：Edgeone请求带宽；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合。
	// 若不填写，默认选择全部站点，且最多只能查询近30天的数据；若填写，则可查询站点绑定套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循<a href="https://zh.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a>规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeTimingL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 指标列表，取值有:
	// <li>l7Flow_outFlux: Edgeone响应流量；</li>
	// <li>l7Flow_inFlux: Edgeone请求流量；</li>
	// <li>l7Flow_outBandwidth: Edgeone响应带宽；</li>
	// <li>l7Flow_inBandwidth：Edgeone请求带宽；</li>
	// <li>l7Flow_request: 访问请求数；</li>
	// <li>l7Flow_flux: 访问请求上行+下行流量；</li>
	// <li>l7Flow_bandwidth：访问请求上行+下行带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合。
	// 若不填写，默认选择全部站点，且最多只能查询近30天的数据；若填写，则可查询站点绑定套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：1小时范围内以min粒度查询，2天范围内以5min粒度查询，7天范围内以hour粒度查询，超过7天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循<a href="https://zh.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a>规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeTimingL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Interval")
	delete(f, "Filters")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7AnalysisDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 时序流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTimingL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeTimingL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Cache_outFlux：响应流量；</li>
	// <li>l7Cache_request：响应请求数；</li>
	// <li> l7Cache_outBandwidth：响应带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 查询时间粒度，可选的值有：
	// <li>min：1分钟的时间粒度；</li>
	// <li>5min：5分钟的时间粒度；</li>
	// <li>hour：1小时的时间粒度；</li>
	// <li>day：1天的时间粒度。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeTimingL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li>l7Cache_outFlux：响应流量；</li>
	// <li>l7Cache_request：响应请求数；</li>
	// <li> l7Cache_outBandwidth：响应带宽。</li>
	MetricNames []*string `json:"MetricNames,omitnil" name:"MetricNames"`

	// 站点集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 查询时间粒度，可选的值有：
	// <li>min：1分钟的时间粒度；</li>
	// <li>5min：5分钟的时间粒度；</li>
	// <li>hour：1小时的时间粒度；</li>
	// <li>day：1天的时间粒度。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeTimingL7CacheDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7CacheDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricNames")
	delete(f, "ZoneIds")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTimingL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTimingL7CacheDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 七层缓存分析时序类流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TimingDataRecord `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTimingL7CacheDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTimingL7CacheDataResponseParams `json:"Response"`
}

func (r *DescribeTimingL7CacheDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTimingL7CacheDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：按国家/地区维度统计流量指标；</li>
	// <li> l7Flow_outFlux_statusCode：按状态码维度统计流量指标；</li>
	// <li> l7Flow_outFlux_domain：按域名维度统计流量指标；</li>
	// <li> l7Flow_outFlux_url：按URL维度统计流量指标; </li>
	// <li> l7Flow_outFlux_resourceType：按资源类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_sip：按客户端的源IP维度统计流量指标；</li>
	// <li> l7Flow_outFlux_referers：按refer信息维度统计流量指标；</li>
	// <li> l7Flow_outFlux_ua_device：按设备类型维度统计流量指标; </li>
	// <li> l7Flow_outFlux_ua_browser：按浏览器类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_us_os：按操作系统类型维度统计流量指标；</li>
	// <li> l7Flow_request_country：按国家/地区维度统计请求数指标；</li>
	// <li> l7Flow_request_statusCode：按状态码维度统计请求数指标；</li>
	// <li> l7Flow_request_domain：按域名维度统计请求数指标；</li>
	// <li> l7Flow_request_url：按URL维度统计请求数指标; </li>
	// <li> l7Flow_request_resourceType：按资源类型维度统计请求数指标；</li>
	// <li> l7Flow_request_sip：按客户端的源IP维度统计请求数指标；</li>
	// <li> l7Flow_request_referer：按refer信息维度统计请求数指标；</li>
	// <li> l7Flow_request_ua_device：按设备类型维度统计请求数指标; </li>
	// <li> l7Flow_request_ua_browser：按浏览器类型维度统计请求数指标；</li>
	// <li> l7Flow_request_us_os：按操作系统类型维度统计请求数指标。</li>
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认默认为: 10， 表示查询前top10的数据。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循<a href="https://zh.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a>规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeTopL7AnalysisDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Flow_outFlux_country：按国家/地区维度统计流量指标；</li>
	// <li> l7Flow_outFlux_statusCode：按状态码维度统计流量指标；</li>
	// <li> l7Flow_outFlux_domain：按域名维度统计流量指标；</li>
	// <li> l7Flow_outFlux_url：按URL维度统计流量指标; </li>
	// <li> l7Flow_outFlux_resourceType：按资源类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_sip：按客户端的源IP维度统计流量指标；</li>
	// <li> l7Flow_outFlux_referers：按refer信息维度统计流量指标；</li>
	// <li> l7Flow_outFlux_ua_device：按设备类型维度统计流量指标; </li>
	// <li> l7Flow_outFlux_ua_browser：按浏览器类型维度统计流量指标；</li>
	// <li> l7Flow_outFlux_us_os：按操作系统类型维度统计流量指标；</li>
	// <li> l7Flow_request_country：按国家/地区维度统计请求数指标；</li>
	// <li> l7Flow_request_statusCode：按状态码维度统计请求数指标；</li>
	// <li> l7Flow_request_domain：按域名维度统计请求数指标；</li>
	// <li> l7Flow_request_url：按URL维度统计请求数指标; </li>
	// <li> l7Flow_request_resourceType：按资源类型维度统计请求数指标；</li>
	// <li> l7Flow_request_sip：按客户端的源IP维度统计请求数指标；</li>
	// <li> l7Flow_request_referer：按refer信息维度统计请求数指标；</li>
	// <li> l7Flow_request_ua_device：按设备类型维度统计请求数指标; </li>
	// <li> l7Flow_request_ua_browser：按浏览器类型维度统计请求数指标；</li>
	// <li> l7Flow_request_us_os：按操作系统类型维度统计请求数指标。</li>
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认默认为: 10， 表示查询前top10的数据。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，详细的过滤条件Key值如下：
	// <li>country<br>   按照【<strong>国家/地区</strong>】进行过滤，国家/地区遵循<a href="https://zh.wikipedia.org/wiki/ISO_3166-1">ISO 3166</a>规范。</li>
	// <li>province<br>   按照【<strong>省份</strong>】进行过滤，此参数只支持服务区域为中国大陆。</li>
	// <li>isp<br>   按照【<strong>运营商</strong>】进行过滤，此参数只支持服务区域为中国大陆。<br>   对应的Value可选项如下：<br>   2：中国电信；<br>   26：中国联通；<br>   1046：中国移动；<br>   3947：中国铁通；<br>   38：教育网；<br>   43：长城宽带；<br>   0：其他运营商。</li>
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。</li>
	// <li>url<br>   按照【<strong>URL Path</strong>】进行过滤，URL Path形如：/content或/content/test.jpg。<br>   若只填写url参数，则最多可查询近30天的数据；<br>   若同时填写url+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>referer<br>   按照【<strong>Referer头信息</strong>】进行过滤, Referer形如：example.com。<br>   若只填写referer参数，则最多可查询近30天的数据；<br>   若同时填写referer+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，资源类型一般是文件后缀，形如: .jpg, .css。<br>   若只填写resourceType参数，则最多可查询近30天的数据；<br>   若同时填写resourceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。</li>
	// <li>protocol<br>   按照【<strong>HTTP协议版本</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP/1.0：HTTP 1.0；<br>   HTTP/1.1：HTTP 1.1；<br>   HTTP/2.0：HTTP 2.0；<br>   HTTP/3.0：HTTP 3.0；<br>   WebSocket：WebSocket。</li>
	// <li>socket<br>   按照【<strong>HTTP协议类型</strong>】进行过滤。<br>   对应的Value可选项如下：<br>   HTTP：HTTP 协议；<br>   HTTPS：HTTPS协议；<br>   QUIC：QUIC协议。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤。<br>   若只填写statusCode参数，则最多可查询近30天的数据；<br>   若同时填写statusCode+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应的Value可选项如下：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   205：205状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>browserType<br>   按照【<strong>浏览器类型</strong>】进行过滤。<br>   若只填写browserType参数，则最多可查询近30天的数据；<br>   若同时填写browserType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Firefox：Firefox浏览器；<br>   Chrome：Chrome浏览器；<br>   Safari：Safari浏览器；<br>   Other：其他浏览器类型；<br>   Empty：浏览器类型为空；<br>   Bot：搜索引擎爬虫；<br>   MicrosoftEdge：MicrosoftEdge浏览器；<br>   IE：IE浏览器；<br>   Opera：Opera浏览器；<br>   QQBrowser：QQ浏览器；<br>   LBBrowser：LB浏览器；<br>   MaxthonBrowser：Maxthon浏览器；<br>   SouGouBrowser：搜狗浏览器；<br>   BIDUBrowser：百度浏览器；<br>   TaoBrowser：淘浏览器；<br>   UBrowser：UC浏览器。</li>
	// <li>deviceType<br>   按照【<strong>设备类型</strong>】进行过滤。<br>   若只填写deviceType参数，则最多可查询近30天的数据；<br>   若同时填写deviceType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TV：TV设备；<br>   Tablet：Tablet设备；<br>   Mobile：Mobile设备；<br>   Desktop：Desktop设备；<br>   Other：其他设备类型；<br>   Empty：设备类型为空。</li>
	// <li>operatingSystemType<br>   按照【<strong>操作系统类型</strong>】进行过滤。<br>   若只填写operatingSystemType参数，则最多可查询近30天的数据；<br>   若同时填写operatingSystemType+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   Linux：Linux操作系统；<br>   MacOS：MacOs操作系统；<br>   Android：Android操作系统；<br>   IOS：IOS操作系统；<br>   Windows：Windows操作系统；<br>   NetBSD：NetBSD；<br>   ChromiumOS：ChromiumOS；<br>   Bot：搜索引擎爬虫；<br>   Other：其他类型的操作系统；<br>   Empty：操作系统为空。</li>
	// <li>tlsVersion<br>   按照【<strong>TLS版本</strong>】进行过滤。<br>   若只填写tlsVersion参数，则最多可查询近30天的数据；<br>   若同时填写tlsVersion+Zonelds参数，则支持的查询数据范围为套餐支持的<a href="https://cloud.tencent.com/document/product/1552/77380#edgeone-.E5.A5.97.E9.A4.90">数据分析最大查询范围</a>与30天两者中的较小值。<br>   对应Value的可选项如下：<br>   TLS1.0：TLS 1.0；<br>   TLS1.1：TLS 1.1；<br>   TLS1.2：TLS 1.2；<br>   TLS1.3：TLS 1.3。</li>
	// <li>ipVersion<br>   按照【<strong>IP版本</strong>】进行过滤。<br>   对应Value的可选项如下：<br>   4：Ipv4；<br>   6：Ipv6。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min：1分钟；</li>
	// <li>5min：5分钟；</li>
	// <li>hour：1小时；</li>
	// <li>day：1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeTopL7AnalysisDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7AnalysisDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7AnalysisDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7AnalysisDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 七层流量前topN数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopL7AnalysisDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopL7AnalysisDataResponseParams `json:"Response"`
}

func (r *DescribeTopL7AnalysisDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7AnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Cache_outFlux_domain：host/域名；</li>
	// <li> l7Cache_outFlux_url：url地址；</li>
	// <li> l7Cache_outFlux_resourceType：资源类型；</li>
	// <li> l7Cache_outFlux_statusCode：状态码。</li>
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 站点id集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeTopL7CacheDataRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的指标，取值有：
	// <li> l7Cache_outFlux_domain：host/域名；</li>
	// <li> l7Cache_outFlux_url：url地址；</li>
	// <li> l7Cache_outFlux_resourceType：资源类型；</li>
	// <li> l7Cache_outFlux_statusCode：状态码。</li>
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 站点id集合，不填默认选择全部站点。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 查询前多少个数据，最大值为1000，不填默认默认为10， 表示查询前top 10的数据。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，详细的过滤条件如下：
	// <li>domain<br>   按照【<strong>子域名</strong>】进行过滤，子域名形如： test.example.com。<br>   类型：String<br>   必选：否</li>
	// <li>url<br>   按照【<strong>URL</strong>】进行过滤，此参数只支持30天的时间范围，URL形如：/content。<br>   类型：String<br>   必选：否</li>
	// <li>resourceType<br>   按照【<strong>资源类型</strong>】进行过滤，此参数只支持30天的时间范围，资源类型形如：jpg，png。<br>   类型：String<br>   必选：否</li>
	// <li>cacheType<br>   按照【<strong>缓存类型</strong>】进行过滤。<br>   类型：String<br>   必选：否<br>   可选项：<br>   hit：命中缓存；<br>   dynamic：资源不可缓存；<br>   miss：未命中缓存。</li>
	// <li>statusCode<br>   按照【<strong>状态码</strong>】进行过滤，此参数只支持30天的时间范围。<br>   类型：String<br>   必选：否<br>   可选项：<br>   1XX：1xx类型的状态码；<br>   100：100状态码；<br>   101：101状态码；<br>   102：102状态码；<br>   2XX：2xx类型的状态码；<br>   200：200状态码；<br>   201：201状态码；<br>   202：202状态码；<br>   203：203状态码；<br>   204：204状态码；<br>   100：100状态码；<br>   206：206状态码；<br>   207：207状态码；<br>   3XX：3xx类型的状态码；<br>   300：300状态码；<br>   301：301状态码；<br>   302：302状态码；<br>   303：303状态码；<br>   304：304状态码；<br>   305：305状态码；<br>   307：307状态码；<br>   4XX：4xx类型的状态码；<br>   400：400状态码；<br>   401：401状态码；<br>   402：402状态码；<br>   403：403状态码；<br>   404：404状态码；<br>   405：405状态码；<br>   406：406状态码；<br>   407：407状态码；<br>   408：408状态码；<br>   409：409状态码；<br>   410：410状态码；<br>   411：411状态码；<br>   412：412状态码；<br>   412：413状态码；<br>   414：414状态码；<br>   415：415状态码；<br>   416：416状态码；<br>   417：417状态码；<br>   422：422状态码；<br>   423：423状态码；<br>   424：424状态码；<br>   426：426状态码；<br>   451：451状态码；<br>   5XX：5xx类型的状态码；<br>   500：500状态码；<br>   501：501状态码；<br>   502：502状态码；<br>   503：503状态码；<br>   504：504状态码；<br>   505：505状态码；<br>   506：506状态码；<br>   507：507状态码；<br>   510：510状态码；<br>   514：514状态码；<br>   544：544状态码。</li>
	// <li>tagKey<br>   按照【<strong>标签Key</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	// <li>tagValue<br>   按照【<strong>标签Value</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>
	Filters []*QueryCondition `json:"Filters,omitnil" name:"Filters"`

	// 查询时间粒度，取值有：
	// <li>min: 1分钟；</li>
	// <li>5min: 5分钟；</li>
	// <li>hour: 1小时；</li>
	// <li>day: 1天。</li>不填将根据开始时间跟结束时间的间距自动推算粒度，具体为：一小时范围内以min粒度查询，两天范围内以5min粒度查询，七天范围内以hour粒度查询，超过七天以day粒度查询。
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 数据归属地区，取值有：
	// <li>overseas：全球（除中国大陆地区）数据；</li>
	// <li>mainland：中国大陆地区数据；</li>
	// <li>global：全球数据。</li>不填默认取值为global。
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeTopL7CacheDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7CacheDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricName")
	delete(f, "ZoneIds")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Interval")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopL7CacheDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopL7CacheDataResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 七层缓存TopN流量数据列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TopDataRecord `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopL7CacheDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopL7CacheDataResponseParams `json:"Response"`
}

func (r *DescribeTopL7CacheDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopL7CacheDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneSettingRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
}

type DescribeZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`
}

func (r *DescribeZoneSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneSettingResponseParams struct {
	// 站点配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneSetting *ZoneSetting `json:"ZoneSetting,omitnil" name:"ZoneSetting"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneSettingResponseParams `json:"Response"`
}

func (r *DescribeZoneSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesRequestParams struct {
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：否</li><li>status<br>   按照【<strong>站点状态</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>tag-key<br>   按照【<strong>标签键</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>tag-value<br>   按照【<strong>标签值</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>模糊查询时仅支持过滤字段名为zone-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段，取值有：
	// <li> type：接入类型；</li>
	// <li> area：加速区域；</li>
	// <li> create-time：创建时间；</li>
	// <li> zone-name：站点名称；</li>
	// <li> use-time：最近使用时间；</li>
	// <li> active-status：生效状态。</li>不填写使用默认值create-time。
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方向，取值有：
	// <li> asc：从小到大排序；</li>
	// <li> desc：从大到小排序。</li>不填写使用默认值desc。
	Direction *string `json:"Direction,omitnil" name:"Direction"`
}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量。默认值：0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页查询限制数目。默认值：20，最大值：1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件，Filters.Values的上限为20。详细的过滤条件如下：
	// <li>zone-name<br>   按照【<strong>站点名称</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>zone-id<br>   按照【<strong>站点ID</strong>】进行过滤。站点ID形如：zone-xxx。<br>   类型：String<br>   必选：否</li><li>status<br>   按照【<strong>站点状态</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>tag-key<br>   按照【<strong>标签键</strong>】进行过滤。<br>   类型：String<br>   必选：否</li><li>tag-value<br>   按照【<strong>标签值</strong>】进行过滤。<br>   类型：String<br>   必选：否</li>模糊查询时仅支持过滤字段名为zone-name。
	Filters []*AdvancedFilter `json:"Filters,omitnil" name:"Filters"`

	// 排序字段，取值有：
	// <li> type：接入类型；</li>
	// <li> area：加速区域；</li>
	// <li> create-time：创建时间；</li>
	// <li> zone-name：站点名称；</li>
	// <li> use-time：最近使用时间；</li>
	// <li> active-status：生效状态。</li>不填写使用默认值create-time。
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方向，取值有：
	// <li> asc：从小到大排序；</li>
	// <li> desc：从大到小排序。</li>不填写使用默认值desc。
	Direction *string `json:"Direction,omitnil" name:"Direction"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 符合条件的站点个数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 站点详细信息列表。
	Zones []*Zone `json:"Zones,omitnil" name:"Zones"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailHost struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速服务状态，取值为：
	// <li> process：部署中；</li>
	// <li> online：已启动；</li>
	// <li> offline：已关闭。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 域名。
	Host *string `json:"Host,omitnil" name:"Host"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 分配的Cname域名
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 资源ID。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 锁状态。
	Lock *int64 `json:"Lock,omitnil" name:"Lock"`

	// 域名状态类型。
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 域名加速地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 加速类型配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateType *AccelerateType `json:"AccelerateType,omitnil" name:"AccelerateType"`

	// Https配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 缓存配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil" name:"CacheConfig"`

	// 源站配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 安全类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityType *SecurityType `json:"SecurityType,omitnil" name:"SecurityType"`

	// 缓存键配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 智能压缩配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// Waf防护配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Waf *Waf `json:"Waf,omitnil" name:"Waf"`

	// CC防护配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CC *CC `json:"CC,omitnil" name:"CC"`

	// DDoS防护配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DDoS *DDoS `json:"DDoS,omitnil" name:"DDoS"`

	// 智能路由配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil" name:"SmartRouting"`

	// Ipv6访问配置项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 回源时是否携带客户端IP所属地域信息的配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil" name:"ClientIpCountry"`
}

type DiffIPWhitelist struct {
	// 最新IP白名单列表。
	LatestIPWhitelist *IPWhitelist `json:"LatestIPWhitelist,omitnil" name:"LatestIPWhitelist"`

	// 最新IP白名单列表相比于当前IP白名单列表，新增部分。
	AddedIPWhitelist *IPWhitelist `json:"AddedIPWhitelist,omitnil" name:"AddedIPWhitelist"`

	// 最新IP白名单列表相比于当前IP白名单列表，删减部分。
	RemovedIPWhitelist *IPWhitelist `json:"RemovedIPWhitelist,omitnil" name:"RemovedIPWhitelist"`

	// 最新IP白名单列表相比于当前IP白名单列表，不变部分。
	NoChangeIPWhitelist *IPWhitelist `json:"NoChangeIPWhitelist,omitnil" name:"NoChangeIPWhitelist"`
}

// Predefined struct for user
type DownloadL4LogsRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 四层实例 ID 集合。
	ProxyIds []*string `json:"ProxyIds,omitnil" name:"ProxyIds"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DownloadL4LogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 四层实例 ID 集合。
	ProxyIds []*string `json:"ProxyIds,omitnil" name:"ProxyIds"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DownloadL4LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL4LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "ProxyIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL4LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL4LogsResponseParams struct {
	// 查询结果的总条数。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 四层离线日志数据列表。
	Data []*L4OfflineLog `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DownloadL4LogsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadL4LogsResponseParams `json:"Response"`
}

func (r *DownloadL4LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL4LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsRequestParams struct {
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 站点集合，此参数必填，不填默认查询为空。
	ZoneIds []*string `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// 子域名集合，不填默认选择全部子域名。
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 分页查询的限制数目，默认值为 20，最大查询条目为 1000。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 分页的偏移量，默认值为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DownloadL7LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ZoneIds")
	delete(f, "Domains")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL7LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadL7LogsResponseParams struct {
	// 查询结果的总条数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 七层离线日志数据列表。
	Data []*L7OfflineLog `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DownloadL7LogsResponse struct {
	*tchttp.BaseResponse
	Response *DownloadL7LogsResponseParams `json:"Response"`
}

func (r *DownloadL7LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DropPageConfig struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// Waf(托管规则)模块的拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafDropPageDetail *DropPageDetail `json:"WafDropPageDetail,omitnil" name:"WafDropPageDetail"`

	// 自定义页面的拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclDropPageDetail *DropPageDetail `json:"AclDropPageDetail,omitnil" name:"AclDropPageDetail"`
}

type DropPageDetail struct {
	// 拦截页面的唯一Id。系统默认包含一个自带拦截页面，Id值为0。
	// 该Id可通过创建拦截页面接口进行上传获取。如传入0，代表使用系统默认拦截页面。
	PageId *int64 `json:"PageId,omitnil" name:"PageId"`

	// 拦截页面的HTTP状态码。状态码范围是100-600。
	StatusCode *int64 `json:"StatusCode,omitnil" name:"StatusCode"`

	// 页面文件名或url。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 页面的类型，取值有：
	// <li> file：页面文件内容；</li>
	// <li> url：上传的url地址。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type ExceptConfig struct {
	// 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 例外规则详情。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRules []*ExceptUserRule `json:"ExceptUserRules,omitnil" name:"ExceptUserRules"`
}

type ExceptUserRule struct {
	// 规则名称，不可使用中文。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则的处置方式，当前仅支持skip：跳过全部托管规则。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 规则生效状态，取值有：
	// <li>on：生效；</li>
	// <li>off：失效。</li>
	RuleStatus *string `json:"RuleStatus,omitnil" name:"RuleStatus"`

	// 规则ID。仅出参使用。默认由底层生成。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 更新时间，如果为null，默认由底层按当前时间生成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 匹配条件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRuleConditions []*ExceptUserRuleCondition `json:"ExceptUserRuleConditions,omitnil" name:"ExceptUserRuleConditions"`

	// 规则生效的范围。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptUserRuleScope *ExceptUserRuleScope `json:"ExceptUserRuleScope,omitnil" name:"ExceptUserRuleScope"`

	// 优先级，取值范围0-100。如果为null，默认由底层设置为0。
	RulePriority *int64 `json:"RulePriority,omitnil" name:"RulePriority"`
}

type ExceptUserRuleCondition struct {
	// 匹配项，取值有：
	// <li>host：请求域名；</li>
	// <li>sip：客户端IP；</li>
	// <li>ua：User-Agent；</li>
	// <li>cookie：会话 Cookie；</li>
	// <li>cgi：CGI 脚本；</li>
	// <li>xff：XFF 扩展头部；</li>
	// <li>url：请求 URL；</li>
	// <li>accept：请求内容类型；</li>
	// <li>method：请求方式；</li>
	// <li>header：请求头部；</li>
	// <li>sip_proto：网络层协议。</li>
	MatchFrom *string `json:"MatchFrom,omitnil" name:"MatchFrom"`

	// 匹配项的参数。仅当 MatchFrom 为 header 时，可以使用本参数，值可填入 header 的 key 作为参数。
	MatchParam *string `json:"MatchParam,omitnil" name:"MatchParam"`

	// 匹配操作符，取值有：
	// <li>equal：字符串等于；</li>
	// <li>not_equal：数值不等于；</li>
	// <li>include：字符包含；</li>
	// <li>not_include：字符不包含；</li>
	// <li>match：ip匹配；</li>
	// <li>not_match：ip不匹配；</li>
	// <li>include_area：地域包含；</li>
	// <li>is_empty：存在字段但值为空；</li>
	// <li>not_exists：不存在关键字段；</li>
	// <li>regexp：正则匹配；</li>
	// <li>len_gt：数值大于；</li>
	// <li>len_lt：数值小于；</li>
	// <li>len_eq：数值等于；</li>
	// <li>match_prefix：前缀匹配；</li>
	// <li>match_suffix：后缀匹配；</li>
	// <li>wildcard：通配符。</li>
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 匹配值。
	MatchContent *string `json:"MatchContent,omitnil" name:"MatchContent"`
}

type ExceptUserRuleScope struct {
	// 例外规则类型。其中complete模式代表全量数据进行例外，partial模式代表可选择指定模块指定字段进行例外，该字段取值有：
	// <li>complete：完全跳过模式；</li>
	// <li>partial：部分跳过模式。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 生效的模块，该字段取值有：
	// <li>waf：托管规则；</li>
	// <li>rate：速率限制；</li>
	// <li>acl：自定义规则；</li>
	// <li>cc：cc攻击防护；</li>
	// <li>bot：Bot防护。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Modules []*string `json:"Modules,omitnil" name:"Modules"`

	// 跳过部分规则ID的例外规则详情。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartialModules []*PartialModule `json:"PartialModules,omitnil" name:"PartialModules"`

	// 跳过具体字段不去扫描的例外规则详情。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipConditions []*SkipCondition `json:"SkipConditions,omitnil" name:"SkipConditions"`
}

type FailReason struct {
	// 失败原因。
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 处理失败的资源列表。
	Targets []*string `json:"Targets,omitnil" name:"Targets"`
}

type FileAscriptionInfo struct {
	// 文件校验目录。
	IdentifyPath *string `json:"IdentifyPath,omitnil" name:"IdentifyPath"`

	// 文件校验内容。
	IdentifyContent *string `json:"IdentifyContent,omitnil" name:"IdentifyContent"`
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FirstPartConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 首段包的统计时长，单位是秒，即期望首段包的统计时长是多少，默认5秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatTime *uint64 `json:"StatTime,omitnil" name:"StatTime"`
}

type FollowOrigin struct {
	// 遵循源站配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 源站未返回 Cache-Control 头时, 设置默认的缓存时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCacheTime *int64 `json:"DefaultCacheTime,omitnil" name:"DefaultCacheTime"`

	// 源站未返回 Cache-Control 头时, 设置缓存/不缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCache *string `json:"DefaultCache,omitnil" name:"DefaultCache"`

	// 源站未返回 Cache-Control 头时, 使用/不使用默认缓存策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultCacheStrategy *string `json:"DefaultCacheStrategy,omitnil" name:"DefaultCacheStrategy"`
}

type ForceRedirect struct {
	// 访问强制跳转配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 重定向状态码，取值有：
	// <li>301：301跳转；</li>
	// <li>302：302跳转。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectStatusCode *int64 `json:"RedirectStatusCode,omitnil" name:"RedirectStatusCode"`
}

type Grpc struct {
	// 是否开启 Grpc 配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Header struct {
	// HTTP头部名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// HTTP头部值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Hsts struct {
	// 是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// MaxAge 数值。单位为秒，最大值为1天。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *int64 `json:"MaxAge,omitnil" name:"MaxAge"`

	// 是否包含子域名，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncludeSubDomains *string `json:"IncludeSubDomains,omitnil" name:"IncludeSubDomains"`

	// 是否开启预加载，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Preload *string `json:"Preload,omitnil" name:"Preload"`
}

type Https struct {
	// http2 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Http2 *string `json:"Http2,omitnil" name:"Http2"`

	// OCSP 配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcspStapling *string `json:"OcspStapling,omitnil" name:"OcspStapling"`

	// Tls 版本设置，取值有：
	// <li>TLSv1：TLSv1版本；</li>
	// <li>TLSV1.1：TLSv1.1版本；</li>
	// <li>TLSV1.2：TLSv1.2版本；</li>
	// <li>TLSv1.3：TLSv1.3版本。</li>修改时必须开启连续的版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TlsVersion []*string `json:"TlsVersion,omitnil" name:"TlsVersion"`

	// HSTS 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hsts *Hsts `json:"Hsts,omitnil" name:"Hsts"`

	// 证书配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertInfo []*ServerCertInfo `json:"CertInfo,omitnil" name:"CertInfo"`

	// 申请类型，取值有：
	// <li>apply：托管EdgeOne；</li>
	// <li>none：不托管EdgeOne。</li>不填，默认取值为none。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyType *string `json:"ApplyType,omitnil" name:"ApplyType"`

	// 密码套件，取值有：
	// <li>loose-v2023：提供最高的兼容性，安全性一般，支持 TLS 1.0-1.3 密码套件；</li>
	// <li>general-v2023：提供较高的兼容性，安全性中等，支持 TLS 1.2-1.3 密码套件；</li>
	// <li>strict-v2023：提供最高的安全性能，禁用所有含不安全隐患的加密套件，支持 TLS 1.2-1.3 密码套件。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherSuite *string `json:"CipherSuite,omitnil" name:"CipherSuite"`
}

type IPGroup struct {
	// 组 Id，创建时填 0 即可。
	GroupId *int64 `json:"GroupId,omitnil" name:"GroupId"`

	// 组名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// IP 组内容，可以填入 IP 及 IP 掩码。
	Content []*string `json:"Content,omitnil" name:"Content"`
}

type IPWhitelist struct {
	// IPv4列表。
	IPv4 []*string `json:"IPv4,omitnil" name:"IPv4"`

	// IPv6列表。
	IPv6 []*string `json:"IPv6,omitnil" name:"IPv6"`
}

type Identification struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 验证子域名。验证站点时，该值为空。验证子域名是为具体子域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 验证状态，取值有：
	// <li> pending：验证中；</li>
	// <li> finished：验证完成。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 站点归属权校验：Dns校验信息。
	Ascription *AscriptionInfo `json:"Ascription,omitnil" name:"Ascription"`

	// 域名当前的 NS 记录。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil" name:"OriginalNameServers"`

	// 站点归属权校验：文件校验信息。
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitnil" name:"FileAscription"`
}

// Predefined struct for user
type IdentifyZoneRequestParams struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 站点下的子域名。如果验证站点下的子域名，则传该值，否则为空。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type IdentifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 站点下的子域名。如果验证站点下的子域名，则传该值，否则为空。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *IdentifyZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdentifyZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneName")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IdentifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IdentifyZoneResponseParams struct {
	// 站点归属校验：Dns校验信息。
	Ascription *AscriptionInfo `json:"Ascription,omitnil" name:"Ascription"`

	// 站点归属权校验：文件校验信息。
	FileAscription *FileAscriptionInfo `json:"FileAscription,omitnil" name:"FileAscription"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IdentifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *IdentifyZoneResponseParams `json:"Response"`
}

func (r *IdentifyZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IdentifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageOptimize struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type IntelligenceRule struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 规则详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntelligenceRuleItems []*IntelligenceRuleItem `json:"IntelligenceRuleItems,omitnil" name:"IntelligenceRuleItems"`
}

type IntelligenceRuleItem struct {
	// 智能分析标签，取值有：
	// <li>evil_bot：恶意bot；</li>
	// <li>suspect_bot：疑似bot；</li>
	// <li>good_bot：良好bot；</li>
	// <li>normal：正常请求。</li>
	Label *string `json:"Label,omitnil" name:"Label"`

	// 触发智能分析标签对应的处置方式，取值有：
	// <li>drop：拦截；</li>
	// <li>trans：放行；</li>
	// <li>alg：Javascript挑战；</li>
	// <li>captcha：数字验证码；</li>
	// <li>monitor：观察。</li>
	Action *string `json:"Action,omitnil" name:"Action"`
}

type IpTableConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 基础管控规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTableRules []*IpTableRule `json:"IpTableRules,omitnil" name:"IpTableRules"`
}

type IpTableRule struct {
	// 动作，取值有：
	// <li> drop：拦截；</li>
	// <li> trans：放行；</li>
	// <li> monitor：观察。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 根据类型匹配，取值有：
	// <li>ip：对ip进行匹配；</li>
	// <li>area：对ip所属地区匹配。</li>
	MatchFrom *string `json:"MatchFrom,omitnil" name:"MatchFrom"`

	// 规则的匹配方式，默认为空代表等于。
	// 取值有：
	// <li> is_emty：配置为空；</li>
	// <li> not_exists：配置为不存在；</li>
	// <li> include：包含；</li>
	// <li> not_include：不包含；</li>
	// <li> equal：等于；</li>
	// <li> not_equal：不等于。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 规则id。仅出参使用。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 更新时间。仅出参使用。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 规则启用状态，当返回为null时，为启用。取值有：
	// <li> on：启用；</li>
	// <li> off：未启用。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规则名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 匹配内容。当 Operator为is_emty 或not_exists时，此值允许为空。
	MatchContent *string `json:"MatchContent,omitnil" name:"MatchContent"`
}

type Ipv6 struct {
	// Ipv6 访问功能配置，取值有：
	// <li>on：开启Ipv6访问功能；</li>
	// <li>off：关闭Ipv6访问功能。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type L4OfflineLog struct {
	// 四层代理实例 ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 日志所属区域，取值有：
	// <li>mainland：中国大陆境内;</li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 离线日志数据包名。
	LogPacketName *string `json:"LogPacketName,omitnil" name:"LogPacketName"`

	// 离线日志下载地址。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 日志打包时间，此参数已经废弃。
	LogTime *int64 `json:"LogTime,omitnil" name:"LogTime"`

	// 日志打包开始时间。
	LogStartTime *string `json:"LogStartTime,omitnil" name:"LogStartTime"`

	// 日志打包结束时间。
	LogEndTime *string `json:"LogEndTime,omitnil" name:"LogEndTime"`

	// 日志大小，单位为 Byte。
	Size *int64 `json:"Size,omitnil" name:"Size"`
}

type L7OfflineLog struct {
	// 离线日志域名。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 日志所属区域，取值有：
	// <li>mainland：中国大陆境内; </li>
	// <li>overseas：全球（不含中国大陆）。</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 离线日志数据包名。	
	LogPacketName *string `json:"LogPacketName,omitnil" name:"LogPacketName"`

	// 离线日志下载地址。	
	Url *string `json:"Url,omitnil" name:"Url"`

	// 日志打包时间，此参数已经废弃。
	LogTime *int64 `json:"LogTime,omitnil" name:"LogTime"`

	// 日志打包开始时间。
	LogStartTime *string `json:"LogStartTime,omitnil" name:"LogStartTime"`

	// 日志打包结束时间。
	LogEndTime *string `json:"LogEndTime,omitnil" name:"LogEndTime"`

	// 日志原始大小，单位 Byte。
	Size *int64 `json:"Size,omitnil" name:"Size"`
}

type MaxAge struct {
	// 是否遵循源站，取值有：
	// <li>on：遵循源站，忽略MaxAge 时间设置；</li>
	// <li>off：不遵循源站，使用MaxAge 时间设置。</li>
	FollowOrigin *string `json:"FollowOrigin,omitnil" name:"FollowOrigin"`

	// MaxAge 时间设置，单位秒，最大365天。
	// 注意：时间为0，即不缓存。
	MaxAgeTime *int64 `json:"MaxAgeTime,omitnil" name:"MaxAgeTime"`
}

// Predefined struct for user
type ModifyAccelerationDomainRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil" name:"OriginInfo"`
}

type ModifyAccelerationDomainRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 加速域名名称。
	DomainName *string `json:"DomainName,omitnil" name:"DomainName"`

	// 源站信息。
	OriginInfo *OriginInfo `json:"OriginInfo,omitnil" name:"OriginInfo"`
}

func (r *ModifyAccelerationDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainName")
	delete(f, "OriginInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerationDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAccelerationDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerationDomainResponseParams `json:"Response"`
}

func (r *ModifyAccelerationDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainStatusesRequestParams struct {
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 要执行状态变更的加速域名列表。
	DomainNames []*string `json:"DomainNames,omitnil" name:"DomainNames"`

	// 加速域名状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 是否强制停用。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制停用该域名，取值有：
	// <li> true：停用该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许停用。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil" name:"Force"`
}

type ModifyAccelerationDomainStatusesRequest struct {
	*tchttp.BaseRequest
	
	// 加速域名所属站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 要执行状态变更的加速域名列表。
	DomainNames []*string `json:"DomainNames,omitnil" name:"DomainNames"`

	// 加速域名状态，取值有：
	// <li>online：启用；</li>
	// <li>offline：停用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 是否强制停用。当域名存在关联资源（如马甲域名、流量调度功能）时，是否强制停用该域名，取值有：
	// <li> true：停用该域名及所有关联资源；</li>
	// <li> false：当该加速域名存在关联资源时，不允许停用。</li>不填写，默认值为：false。
	Force *bool `json:"Force,omitnil" name:"Force"`
}

func (r *ModifyAccelerationDomainStatusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainStatusesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "DomainNames")
	delete(f, "Status")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccelerationDomainStatusesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccelerationDomainStatusesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAccelerationDomainStatusesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccelerationDomainStatusesResponseParams `json:"Response"`
}

func (r *ModifyAccelerationDomainStatusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccelerationDomainStatusesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书；</li>
	// <li> apply：申请免费证书。</li>不填写保持原有配置。
	CertType *string `json:"CertType,omitnil" name:"CertType"`

	// 当 CertType 取值为 hosting 时填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil" name:"CertId"`
}

type ModifyAliasDomainRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 别称域名名称。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// 目标域名名称。
	TargetName *string `json:"TargetName,omitnil" name:"TargetName"`

	// 证书配置，取值有：
	// <li> none：不配置；</li>
	// <li> hosting：SSL托管证书；</li>
	// <li> apply：申请免费证书。</li>不填写保持原有配置。
	CertType *string `json:"CertType,omitnil" name:"CertType"`

	// 当 CertType 取值为 hosting 时填入相应证书 ID。
	CertId []*string `json:"CertId,omitnil" name:"CertId"`
}

func (r *ModifyAliasDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "AliasName")
	delete(f, "TargetName")
	delete(f, "CertType")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAliasDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAliasDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAliasDomainResponseParams `json:"Response"`
}

func (r *ModifyAliasDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 别称域名状态，取值有：
	// <li> false：开启别称域名；</li>
	// <li> true：关闭别称域名。</li>
	Paused *bool `json:"Paused,omitnil" name:"Paused"`

	// 待修改状态的别称域名名称。如果为空，则不执行修改状态操作。
	AliasNames []*string `json:"AliasNames,omitnil" name:"AliasNames"`
}

type ModifyAliasDomainStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 别称域名状态，取值有：
	// <li> false：开启别称域名；</li>
	// <li> true：关闭别称域名。</li>
	Paused *bool `json:"Paused,omitnil" name:"Paused"`

	// 待修改状态的别称域名名称。如果为空，则不执行修改状态操作。
	AliasNames []*string `json:"AliasNames,omitnil" name:"AliasNames"`
}

func (r *ModifyAliasDomainStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Paused")
	delete(f, "AliasNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAliasDomainStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAliasDomainStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAliasDomainStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAliasDomainStatusResponseParams `json:"Response"`
}

func (r *ModifyAliasDomainStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAliasDomainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理 ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil" name:"ProxyName"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写保持原有配置。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写保持原有配置。
	ProxyType *string `json:"ProxyType,omitnil" name:"ProxyType"`

	// Ipv6 访问配置，不填写保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 中国大陆加速优化配置。 不填写表示保持原有配置。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil" name:"AccelerateMainland"`
}

type ModifyApplicationProxyRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理 ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 当 ProxyType=hostname 时，表示域名或子域名；
	// 当 ProxyType=instance 时，表示代理名称。
	ProxyName *string `json:"ProxyName,omitnil" name:"ProxyName"`

	// 会话保持时间，取值范围：30-3600，单位：秒。
	// 不填写保持原有配置。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 四层代理模式，取值有：
	// <li>hostname：表示子域名模式；</li>
	// <li>instance：表示实例模式。</li>不填写保持原有配置。
	ProxyType *string `json:"ProxyType,omitnil" name:"ProxyType"`

	// Ipv6 访问配置，不填写保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 中国大陆加速优化配置。 不填写表示保持原有配置。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil" name:"AccelerateMainland"`
}

func (r *ModifyApplicationProxyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "ProxyName")
	delete(f, "SessionPersistTime")
	delete(f, "ProxyType")
	delete(f, "Ipv6")
	delete(f, "AccelerateMainland")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApplicationProxyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>不填保持原有值。
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil" name:"Port"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>不填保持原有值。
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	// 
	// 不填保持原有值。
	OriginValue []*string `json:"OriginValue,omitnil" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>不填保持原有值。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>不填为false。
	SessionPersist *bool `json:"SessionPersist,omitnil" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil" name:"OriginPort"`

	// 规则标签。不填保持原有值。
	RuleTag *string `json:"RuleTag,omitnil" name:"RuleTag"`
}

type ModifyApplicationProxyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 源站类型，取值有：
	// <li>custom：手动添加；</li>
	// <li>origins：源站组。</li>不填保持原有值。
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 端口，支持格式：
	// <li>80：80端口；</li>
	// <li>81-90：81至90端口。</li>
	Port []*string `json:"Port,omitnil" name:"Port"`

	// 协议，取值有：
	// <li>TCP：TCP协议；</li>
	// <li>UDP：UDP协议。</li>不填保持原有值。
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 源站信息：
	// <li>当 OriginType 为 custom 时，表示一个或多个源站，如`["8.8.8.8","9.9.9.9"]` 或 `OriginValue=["test.com"]`；</li>
	// <li>当 OriginType 为 origins 时，要求有且仅有一个元素，表示源站组ID，如`["origin-537f5b41-162a-11ed-abaa-525400c5da15"]`。</li>
	// 
	// 不填保持原有值。
	OriginValue []*string `json:"OriginValue,omitnil" name:"OriginValue"`

	// 传递客户端IP，取值有：
	// <li>TOA：TOA（仅Proto=TCP时可选）；</li>
	// <li>PPV1：Proxy Protocol传递，协议版本V1（仅Proto=TCP时可选）；</li>
	// <li>PPV2：Proxy Protocol传递，协议版本V2；</li>
	// <li>OFF：不传递。</li>不填保持原有值。
	ForwardClientIp *string `json:"ForwardClientIp,omitnil" name:"ForwardClientIp"`

	// 是否开启会话保持，取值有：
	// <li>true：开启；</li>
	// <li>false：关闭。</li>不填为false。
	SessionPersist *bool `json:"SessionPersist,omitnil" name:"SessionPersist"`

	// 会话保持的时间，只有当SessionPersist为true时，该值才会生效。
	SessionPersistTime *uint64 `json:"SessionPersistTime,omitnil" name:"SessionPersistTime"`

	// 源站端口，支持格式：
	// <li>单端口：80；</li>
	// <li>端口段：81-90，81至90端口。</li>
	OriginPort *string `json:"OriginPort,omitnil" name:"OriginPort"`

	// 规则标签。不填保持原有值。
	RuleTag *string `json:"RuleTag,omitnil" name:"RuleTag"`
}

func (r *ModifyApplicationProxyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	delete(f, "OriginType")
	delete(f, "Port")
	delete(f, "Proto")
	delete(f, "OriginValue")
	delete(f, "ForwardClientIp")
	delete(f, "SessionPersist")
	delete(f, "SessionPersistTime")
	delete(f, "OriginPort")
	delete(f, "RuleTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApplicationProxyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyRuleResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleStatusRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

type ModifyApplicationProxyRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyApplicationProxyRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "RuleId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyRuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApplicationProxyRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyStatusRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

type ModifyApplicationProxyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 代理ID。
	ProxyId *string `json:"ProxyId,omitnil" name:"ProxyId"`

	// 状态，取值有：
	// <li>offline: 停用；</li>
	// <li>online: 启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyApplicationProxyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "ProxyId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationProxyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationProxyStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApplicationProxyStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationProxyStatusResponseParams `json:"Response"`
}

func (r *ModifyApplicationProxyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationProxyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次变更的域名列表。
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 证书信息, 只需要传入 CertId 即可, 如果为空, 则使用默认证书。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil" name:"ServerCertInfo"`

	// 托管类型，取值有：
	// <li>apply：托管EO；</li>
	// <li>none：不托管EO；</li>不填，默认取值为none。
	ApplyType *string `json:"ApplyType,omitnil" name:"ApplyType"`
}

type ModifyHostsCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 本次变更的域名列表。
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 证书信息, 只需要传入 CertId 即可, 如果为空, 则使用默认证书。
	ServerCertInfo []*ServerCertInfo `json:"ServerCertInfo,omitnil" name:"ServerCertInfo"`

	// 托管类型，取值有：
	// <li>apply：托管EO；</li>
	// <li>none：不托管EO；</li>不填，默认取值为none。
	ApplyType *string `json:"ApplyType,omitnil" name:"ApplyType"`
}

func (r *ModifyHostsCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Hosts")
	delete(f, "ServerCertInfo")
	delete(f, "ApplyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostsCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostsCertificateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyHostsCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostsCertificateResponseParams `json:"Response"`
}

func (r *ModifyHostsCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostsCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupRequestParams struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil" name:"OriginGroupId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitnil" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitnil" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitnil" name:"OriginRecords"`

	// 回源Host，仅当OriginType=self时可以设置。
	// 不填写，表示使用已有配置。
	HostHeader *string `json:"HostHeader,omitnil" name:"HostHeader"`
}

type ModifyOriginGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil" name:"OriginGroupId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitnil" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置；</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitnil" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitnil" name:"OriginRecords"`

	// 回源Host，仅当OriginType=self时可以设置。
	// 不填写，表示使用已有配置。
	HostHeader *string `json:"HostHeader,omitnil" name:"HostHeader"`
}

func (r *ModifyOriginGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "OriginGroupId")
	delete(f, "OriginType")
	delete(f, "OriginGroupName")
	delete(f, "ConfigurationType")
	delete(f, "OriginRecords")
	delete(f, "HostHeader")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyOriginGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyOriginGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyOriginGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyOriginGroupResponseParams `json:"Response"`
}

func (r *ModifyOriginGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyOriginGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 规则名称，字符串名称长度 1~255。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil" name:"Rules"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 规则名称，字符串名称长度 1~255。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil" name:"Rules"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 规则状态，取值有：
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "RuleName")
	delete(f, "Rules")
	delete(f, "RuleId")
	delete(f, "Status")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleResponseParams `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityIPGroupRequestParams struct {
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// IP 组配置。
	IPGroup *IPGroup `json:"IPGroup,omitnil" name:"IPGroup"`

	// 操作类型，取值有：
	// <li> append: 向 IPGroup 中追加 Content 参数中内容；</li>
	// <li> remove: 从 IPGroup 中删除 Content 参数中内容；</li>
	// <li> update: 全量替换 IPGroup 内容，并可修改 IPGroup 名称。 </li>
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

type ModifySecurityIPGroupRequest struct {
	*tchttp.BaseRequest
	
	// 站点 Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// IP 组配置。
	IPGroup *IPGroup `json:"IPGroup,omitnil" name:"IPGroup"`

	// 操作类型，取值有：
	// <li> append: 向 IPGroup 中追加 Content 参数中内容；</li>
	// <li> remove: 从 IPGroup 中删除 Content 参数中内容；</li>
	// <li> update: 全量替换 IPGroup 内容，并可修改 IPGroup 名称。 </li>
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

func (r *ModifySecurityIPGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityIPGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "IPGroup")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityIPGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityIPGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySecurityIPGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityIPGroupResponseParams `json:"Response"`
}

func (r *ModifySecurityIPGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityIPGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyRequestParams struct {
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 安全配置。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil" name:"SecurityConfig"`

	// 子域名/应用名。当使用Entity时可不填写TemplateId，否则必须填写TemplateId。
	Entity *string `json:"Entity,omitnil" name:"Entity"`

	// 模板策略id。当使用模板Id时可不填Entity，否则必须填写Entity。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

type ModifySecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 站点Id。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 安全配置。
	SecurityConfig *SecurityConfig `json:"SecurityConfig,omitnil" name:"SecurityConfig"`

	// 子域名/应用名。当使用Entity时可不填写TemplateId，否则必须填写TemplateId。
	Entity *string `json:"Entity,omitnil" name:"Entity"`

	// 模板策略id。当使用模板Id时可不填Entity，否则必须填写Entity。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`
}

func (r *ModifySecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "SecurityConfig")
	delete(f, "Entity")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityPolicyResponseParams `json:"Response"`
}

func (r *ModifySecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 站点接入方式，取值有：
	// <li> full：NS 接入；</li>
	// <li> partial：CNAME 接入，如果站点当前是无域名接入，仅支持切换到CNAME接入。</li>不填写保持原有配置。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 自定义站点信息，以替代系统默认分配的名称服务器。不填写保持原有配置。当站点是无域名接入方式时不允许传此参数。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil" name:"VanityNameServers"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	AliasZoneName *string `json:"AliasZoneName,omitnil" name:"AliasZoneName"`

	// 站点接入地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>当站点是无域名接入方式时，不允许传此参数。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 站点名称。仅当站点由无域名接入方式切换到CNAME接入方式的场景下有效。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`
}

type ModifyZoneRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 站点接入方式，取值有：
	// <li> full：NS 接入；</li>
	// <li> partial：CNAME 接入，如果站点当前是无域名接入，仅支持切换到CNAME接入。</li>不填写保持原有配置。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 自定义站点信息，以替代系统默认分配的名称服务器。不填写保持原有配置。当站点是无域名接入方式时不允许传此参数。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil" name:"VanityNameServers"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	AliasZoneName *string `json:"AliasZoneName,omitnil" name:"AliasZoneName"`

	// 站点接入地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>当站点是无域名接入方式时，不允许传此参数。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 站点名称。仅当站点由无域名接入方式切换到CNAME接入方式的场景下有效。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`
}

func (r *ModifyZoneRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "VanityNameServers")
	delete(f, "AliasZoneName")
	delete(f, "Area")
	delete(f, "ZoneName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyZoneResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneResponseParams `json:"Response"`
}

func (r *ModifyZoneResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingRequestParams struct {
	// 待变更的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 缓存过期时间配置。
	// 不填写表示保持原有配置。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil" name:"CacheConfig"`

	// 节点缓存键配置。
	// 不填写表示保持原有配置。
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 浏览器缓存配置。
	// 不填写表示保持原有配置。
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// 离线缓存配置。
	// 不填写表示保持原有配置。
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// Quic 访问配置。
	// 不填写表示保持原有配置。
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// Post 请求传输配置。
	// 不填写表示保持原有配置。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil" name:"PostMaxSize"`

	// 智能压缩配置。
	// 不填写表示保持原有配置。
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// Http2 回源配置。
	// 不填写表示保持原有配置。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil" name:"UpstreamHttp2"`

	// 访问协议强制 Https 跳转配置。
	// 不填写表示保持原有配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Https 加速配置。
	// 不填写表示保持原有配置。
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 源站配置。
	// 不填写表示保持原有配置。
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 智能加速配置。
	// 不填写表示保持原有配置。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil" name:"SmartRouting"`

	// WebSocket 配置。
	// 不填写表示保持原有配置。
	WebSocket *WebSocket `json:"WebSocket,omitnil" name:"WebSocket"`

	// 客户端 IP 回源请求头配置。
	// 不填写表示保持原有配置。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil" name:"CachePrefresh"`

	// Ipv6 访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 回源时是否携带客户端 IP 所属地域信息的配置。
	// 不填写表示保持原有配置。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil" name:"ClientIpCountry"`

	// Grpc 协议支持配置。
	// 不填写表示保持原有配置。
	Grpc *Grpc `json:"Grpc,omitnil" name:"Grpc"`

	// 图片优化配置。
	// 不填写表示关闭。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil" name:"ImageOptimize"`

	// 标准 Debug 配置。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil" name:"StandardDebug"`
}

type ModifyZoneSettingRequest struct {
	*tchttp.BaseRequest
	
	// 待变更的站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 缓存过期时间配置。
	// 不填写表示保持原有配置。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil" name:"CacheConfig"`

	// 节点缓存键配置。
	// 不填写表示保持原有配置。
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// 浏览器缓存配置。
	// 不填写表示保持原有配置。
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// 离线缓存配置。
	// 不填写表示保持原有配置。
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// Quic 访问配置。
	// 不填写表示保持原有配置。
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// Post 请求传输配置。
	// 不填写表示保持原有配置。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil" name:"PostMaxSize"`

	// 智能压缩配置。
	// 不填写表示保持原有配置。
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// Http2 回源配置。
	// 不填写表示保持原有配置。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil" name:"UpstreamHttp2"`

	// 访问协议强制 Https 跳转配置。
	// 不填写表示保持原有配置。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// Https 加速配置。
	// 不填写表示保持原有配置。
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 源站配置。
	// 不填写表示保持原有配置。
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 智能加速配置。
	// 不填写表示保持原有配置。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil" name:"SmartRouting"`

	// WebSocket 配置。
	// 不填写表示保持原有配置。
	WebSocket *WebSocket `json:"WebSocket,omitnil" name:"WebSocket"`

	// 客户端 IP 回源请求头配置。
	// 不填写表示保持原有配置。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 不填写表示保持原有配置。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil" name:"CachePrefresh"`

	// Ipv6 访问配置。
	// 不填写表示保持原有配置。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// 回源时是否携带客户端 IP 所属地域信息的配置。
	// 不填写表示保持原有配置。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil" name:"ClientIpCountry"`

	// Grpc 协议支持配置。
	// 不填写表示保持原有配置。
	Grpc *Grpc `json:"Grpc,omitnil" name:"Grpc"`

	// 图片优化配置。
	// 不填写表示关闭。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil" name:"ImageOptimize"`

	// 标准 Debug 配置。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil" name:"StandardDebug"`
}

func (r *ModifyZoneSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "CacheConfig")
	delete(f, "CacheKey")
	delete(f, "MaxAge")
	delete(f, "OfflineCache")
	delete(f, "Quic")
	delete(f, "PostMaxSize")
	delete(f, "Compression")
	delete(f, "UpstreamHttp2")
	delete(f, "ForceRedirect")
	delete(f, "Https")
	delete(f, "Origin")
	delete(f, "SmartRouting")
	delete(f, "WebSocket")
	delete(f, "ClientIpHeader")
	delete(f, "CachePrefresh")
	delete(f, "Ipv6")
	delete(f, "ClientIpCountry")
	delete(f, "Grpc")
	delete(f, "ImageOptimize")
	delete(f, "StandardDebug")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneSettingResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyZoneSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneSettingResponseParams `json:"Response"`
}

func (r *ModifyZoneSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneStatusRequestParams struct {
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 站点状态，取值有：
	// <li> false：开启站点；</li>
	// <li> true：关闭站点。</li>
	Paused *bool `json:"Paused,omitnil" name:"Paused"`
}

type ModifyZoneStatusRequest struct {
	*tchttp.BaseRequest
	
	// 站点 ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 站点状态，取值有：
	// <li> false：开启站点；</li>
	// <li> true：关闭站点。</li>
	Paused *bool `json:"Paused,omitnil" name:"Paused"`
}

func (r *ModifyZoneStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Paused")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyZoneStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyZoneStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyZoneStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyZoneStatusResponseParams `json:"Response"`
}

func (r *ModifyZoneStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyZoneStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoCache struct {
	// 不缓存配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type NormalAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 参数。
	Parameters []*RuleNormalActionParams `json:"Parameters,omitnil" name:"Parameters"`
}

type OfflineCache struct {
	// 离线缓存是否开启，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Origin struct {
	// 主源站列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origins []*string `json:"Origins,omitnil" name:"Origins"`

	// 备源站列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupOrigins []*string `json:"BackupOrigins,omitnil" name:"BackupOrigins"`

	// 回源协议配置，取值有：
	// <li>http：强制 http 回源；</li>
	// <li>follow：协议跟随回源；</li>
	// <li>https：强制 https 回源。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginPullProtocol *string `json:"OriginPullProtocol,omitnil" name:"OriginPullProtocol"`

	// 源站为腾讯云 COS 时，是否为私有访问 bucket，取值有：
	// <li>on：私有访问；</li>
	// <li>off：公共访问。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPrivateAccess *string `json:"CosPrivateAccess,omitnil" name:"CosPrivateAccess"`
}

type OriginDetail struct {
	// 源站类型，取值有：
	// <li>IP_DOMAIN：IPV4、IPV6或域名类型源站；</li>
	// <li>COS：COS源。</li>
	// <li>ORIGIN_GROUP：源站组类型源站。</li>
	// <li>AWS_S3：AWS S3对象存储源站。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站地址，当OriginType参数指定为ORIGIN_GROUP时，该参数填写源站组ID，其他情况下填写源站地址。
	Origin *string `json:"Origin,omitnil" name:"Origin"`

	// 备用源站组ID，该参数在OriginType参数指定为ORIGIN_GROUP时生效，为空表示不使用备用源站。
	BackupOrigin *string `json:"BackupOrigin,omitnil" name:"BackupOrigin"`

	// 主源源站组名称，当OriginType参数指定为ORIGIN_GROUP时该参数生效。
	OriginGroupName *string `json:"OriginGroupName,omitnil" name:"OriginGroupName"`

	// 备用源站源站组名称，当OriginType参数指定为ORIGIN_GROUP，且用户指定了被用源站时该参数生效。
	BackOriginGroupName *string `json:"BackOriginGroupName,omitnil" name:"BackOriginGroupName"`

	// 指定是否允许访问私有对象存储源站。当源站类型OriginType=COS或AWS_S3时有效 取值有：
	// <li>on：使用私有鉴权；</li>
	// <li>off：不使用私有鉴权。</li>
	// 不填写，默认值为off。
	PrivateAccess *string `json:"PrivateAccess,omitnil" name:"PrivateAccess"`

	// 私有鉴权使用参数，当源站类型PrivateAccess=on时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil" name:"PrivateParameters"`
}

type OriginGroup struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 源站组ID。
	OriginGroupId *string `json:"OriginGroupId,omitnil" name:"OriginGroupId"`

	// 源站类型，取值有：
	// <li>self：自有源站；</li>
	// <li>third_party：第三方源站；</li>
	// <li>cos：腾讯云COS源站。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站组名称。
	OriginGroupName *string `json:"OriginGroupName,omitnil" name:"OriginGroupName"`

	// 源站配置类型，当OriginType=self时，取值有：
	// <li>area：按区域配置；</li>
	// <li>weight： 按权重配置。</li>
	// <li>proto： 按HTTP协议配置。</li>当OriginType=third_party/cos时放空。
	ConfigurationType *string `json:"ConfigurationType,omitnil" name:"ConfigurationType"`

	// 源站记录信息。
	OriginRecords []*OriginRecord `json:"OriginRecords,omitnil" name:"OriginRecords"`

	// 源站组更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 当OriginType=self时，表示回源Host。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostHeader *string `json:"HostHeader,omitnil" name:"HostHeader"`
}

type OriginInfo struct {
	// 源站类型，取值有：
	// <li>IP_DOMAIN：IPV4、IPV6 或域名类型源站；</li>
	// <li>COS：COS 源；</li>
	// <li>ORIGIN_GROUP：源站组类型源站；</li>
	// <li>AWS_S3：S3兼容对象存储源站；</li>
	// <li>LB: 负载均衡类型源站；</li>
	// <li>SPACE：EdgeOne Shield Space 存储。</li>
	OriginType *string `json:"OriginType,omitnil" name:"OriginType"`

	// 源站地址，当 OriginType 参数指定为 ORIGIN_GROUP 时，该参数填写源站组 ID，其他情况下填写源站地址。
	Origin *string `json:"Origin,omitnil" name:"Origin"`

	// 备用源站组 ID，该参数在 OriginType 参数指定为 ORIGIN_GROUP 时生效，为空表示不使用备用源站。
	BackupOrigin *string `json:"BackupOrigin,omitnil" name:"BackupOrigin"`

	// 指定是否允许访问私有对象存储源站，当源站类型 OriginType=COS 或 AWS_S3 时有效，取值有：
	// <li>on：使用私有鉴权；</li>
	// <li>off：不使用私有鉴权。</li>默认值：off。
	PrivateAccess *string `json:"PrivateAccess,omitnil" name:"PrivateAccess"`

	// 私有鉴权使用参数，当源站类型 PrivateAccess=on 时有效。
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil" name:"PrivateParameters"`
}

type OriginProtectionInfo struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 域名列表。
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 代理ID列表。
	ProxyIds []*string `json:"ProxyIds,omitnil" name:"ProxyIds"`

	// 当前版本的IP白名单。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentIPWhitelist *IPWhitelist `json:"CurrentIPWhitelist,omitnil" name:"CurrentIPWhitelist"`

	// 该站点是否需要更新源站白名单，取值有：
	// <li>true ：需要更新IP白名单 ；</li>
	// <li>false ：无需更新IP白名单。</li>
	NeedUpdate *bool `json:"NeedUpdate,omitnil" name:"NeedUpdate"`

	// 源站防护状态，取值有：
	// <li>online ：源站防护启用中 ；</li>
	// <li>offline ：源站防护已停用 ；</li>
	// <li>nonactivate ：源站防护未激活，仅在从未使用过源站防护功能的站点调用中返回。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 站点套餐是否支持源站防护，取值有：
	// <li>true ：支持 ；</li>
	// <li>false ：不支持。</li>
	PlanSupport *bool `json:"PlanSupport,omitnil" name:"PlanSupport"`

	// 最新IP白名单与当前IP白名单的对比。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiffIPWhitelist *DiffIPWhitelist `json:"DiffIPWhitelist,omitnil" name:"DiffIPWhitelist"`
}

type OriginRecord struct {
	// 源站记录值，不包含端口信息，可以为：IPv4，IPv6，域名格式。
	Record *string `json:"Record,omitnil" name:"Record"`

	// 源站记录ID。
	RecordId *string `json:"RecordId,omitnil" name:"RecordId"`

	// 源站端口，取值范围：[1-65535]。
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 当源站配置类型ConfigurationType=weight时，表示权重。
	// 不配置权重信息时，所有源站组记录统一填写为0或者不填写，表示多个源站轮询回源。
	// 配置权重信息时，取值为[1-100]，多个源站权重总和应为100，表示多个源站按照权重回源。
	// 当源站配置类型ConfigurationType=proto时，表示权重。
	// 不配置权重信息时，所有源站组记录统一填写为0或者不填写，表示多个源站轮询回源。
	// 配置权重信息时，取值为[1-100]，源站组内Proto相同的多个源站权重总和应为100，表示多个源站按照权重回源。
	Weight *uint64 `json:"Weight,omitnil" name:"Weight"`

	// 当源站配置类型ConfigurationType=proto时，表示源站的协议类型，将按照客户端请求协议回到相应的源站，取值有：
	// <li>http：HTTP协议源站；</li>
	// <li>https：HTTPS协议源站。</li>
	Proto *string `json:"Proto,omitnil" name:"Proto"`

	// 当源站配置类型ConfigurationType=area时，表示区域，为空表示全部地区。取值为iso-3166中alpha-2编码或者大洲区域代码。大洲区域代码取值为：
	// <li>Asia：亚洲；</li>
	// <li>Europe：欧洲；</li>
	// <li>Africa：非洲；</li>
	// <li>Oceania：大洋洲；</li>
	// <li>Americas：美洲。</li>源站组记录中，至少需要有一项为全部地区。
	Area []*string `json:"Area,omitnil" name:"Area"`

	// 当源站类型OriginType=third_part时有效
	// 是否私有鉴权，取值有：
	// <li>true：使用私有鉴权；</li>
	// <li>false：不使用私有鉴权。</li>不填写，默认值为：false。
	Private *bool `json:"Private,omitnil" name:"Private"`

	// 当源站类型Private=true时有效，表示私有鉴权使用参数。
	PrivateParameters []*PrivateParameter `json:"PrivateParameters,omitnil" name:"PrivateParameters"`
}

type PartialModule struct {
	// 模块名称，取值为：
	// <li>waf：托管规则。</li>
	Module *string `json:"Module,omitnil" name:"Module"`

	// 模块下的需要例外的具体规则ID列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Include []*int64 `json:"Include,omitnil" name:"Include"`
}

type PlanInfo struct {
	// 结算货币类型，取值有：
	// <li> CNY ：人民币结算； </li>
	// <li> USD ：美元结算。</li>
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 套餐所含流量，该流量数值为安全加速流量，内容加速流量和智能加速流量的总和（单位：字节）。
	Flux *uint64 `json:"Flux,omitnil" name:"Flux"`

	// 结算周期，取值有：
	// <li> y ：按年结算； </li>
	// <li> m ：按月结算；</li>
	// <li> h ：按小时结算； </li>
	// <li> M ：按分钟结算；</li>
	// <li> s ：按秒结算。 </li>
	Frequency *string `json:"Frequency,omitnil" name:"Frequency"`

	// 套餐类型，取值有：
	// <li> sta ：全球内容分发网络（不包括中国大陆）标准版套餐； </li>
	// <li> sta_with_bot ：全球内容分发网络（不包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> sta_cm ：中国大陆内容分发网络标准版套餐； </li>
	// <li> sta_cm_with_bot ：中国大陆内容分发网络标准版套餐附带bot管理；</li>
	// <li> sta_global ：全球内容分发网络（包括中国大陆）标准版套餐； </li>
	// <li> sta_global_with_bot ：全球内容分发网络（包括中国大陆）标准版套餐附带bot管理；</li>
	// <li> ent ：全球内容分发网络（不包括中国大陆）企业版套餐； </li>
	// <li> ent_with_bot ： 全球内容分发网络（不包括中国大陆）企业版套餐附带bot管理；</li>
	// <li> ent_cm ：中国大陆内容分发网络企业版套餐； </li>
	// <li> ent_cm_with_bot ：中国大陆内容分发网络企业版套餐附带bot管理；</li>
	// <li> ent_global ：全球内容分发网络（包括中国大陆）企业版套餐； </li>
	// <li> ent_global_with_bot ：全球内容分发网络（包括中国大陆）企业版套餐附带bot管理。</li>
	PlanType *string `json:"PlanType,omitnil" name:"PlanType"`

	// 套餐价格（单位：分）。
	Price *float64 `json:"Price,omitnil" name:"Price"`

	// 套餐所含请求次数，该请求次数为安全加速请求次数。（单位：次）。
	Request *uint64 `json:"Request,omitnil" name:"Request"`

	// 套餐所能绑定的站点个数。
	SiteNumber *uint64 `json:"SiteNumber,omitnil" name:"SiteNumber"`

	// 套餐加速区域类型，取值有：
	// <li> mainland ：中国大陆； </li>
	// <li> overseas ：全球（不包括中国大陆）；</li>
	// <li> global ：全球（包括中国大陆）。 </li>
	Area *string `json:"Area,omitnil" name:"Area"`
}

type PostMaxSize struct {
	// 是否开启 POST 请求上传文件限制，平台默认为限制为32MB，取值有：
	// <li>on：开启限制；</li>
	// <li>off：关闭限制。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 最大限制，取值在1MB和500MB之间。单位字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSize *int64 `json:"MaxSize,omitnil" name:"MaxSize"`
}

type PrivateParameter struct {
	// 私有鉴权参数名称，取值有：
	// <li>AccessKeyId：鉴权参数Access Key ID；</li>
	// <li>SecretAccessKey：鉴权参数Secret Access Key。</li>
	Name *string `json:"Name,omitnil" name:"Name"`

	// 私有鉴权参数值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type QueryCondition struct {
	// 筛选条件的key。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 查询条件操作符，操作类型有：
	// <li>equals: 等于；</li>
	// <li>notEquals: 不等于；</li>
	// <li>include: 包含；</li>
	// <li>notInclude: 不包含; </li>
	// <li>startWith: 开始的值是value；</li>
	// <li>notStartWith: 不以value的值开始；</li>
	// <li>endWith: 结尾是value值；</li>
	// <li>notEndWith: 不以value的值结尾。</li>
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 筛选条件的值。
	Value []*string `json:"Value,omitnil" name:"Value"`
}

type QueryString struct {
	// CacheKey是否由QueryString组成，取值有：
	// <li>on：是；</li>
	// <li>off：否。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// CacheKey使用QueryString的方式，取值有：
	// <li>includeCustom：使用部分url参数；</li>
	// <li>excludeCustom：排除部分url参数。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 使用/排除的url参数数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value []*string `json:"Value,omitnil" name:"Value"`
}

type Quic struct {
	// 是否开启 Quic 配置，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type Quota struct {
	// 单次批量提交配额上限。
	Batch *int64 `json:"Batch,omitnil" name:"Batch"`

	// 每日提交配额上限。
	Daily *int64 `json:"Daily,omitnil" name:"Daily"`

	// 每日剩余的可提交配额。
	DailyAvailable *int64 `json:"DailyAvailable,omitnil" name:"DailyAvailable"`

	// 刷新预热缓存类型，取值有：
	// <li> purge_prefix：按前缀刷新；</li>
	// <li> purge_url：按URL刷新；</li>
	// <li> purge_host：按Hostname刷新；</li>
	// <li> purge_all：刷新全部缓存内容；</li>
	// <li> purge_cache_tag：按CacheTag刷新；</li><li> prefetch_url：按URL预热。</li>
	Type *string `json:"Type,omitnil" name:"Type"`
}

type RateLimitConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 速率限制-用户规则列表。如果为null，默认使用历史配置。
	RateLimitUserRules []*RateLimitUserRule `json:"RateLimitUserRules,omitnil" name:"RateLimitUserRules"`

	// 速率限制模板功能。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitTemplate *RateLimitTemplate `json:"RateLimitTemplate,omitnil" name:"RateLimitTemplate"`

	// 智能客户端过滤。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitIntelligence *RateLimitIntelligence `json:"RateLimitIntelligence,omitnil" name:"RateLimitIntelligence"`

	// 速率限制-托管定制规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitCustomizes []*RateLimitUserRule `json:"RateLimitCustomizes,omitnil" name:"RateLimitCustomizes"`
}

type RateLimitIntelligence struct {
	// 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 执行动作，取值有：
	// <li>monitor：观察；</li>
	// <li>alg：挑战。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 规则id，仅出参使用。
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`
}

type RateLimitTemplate struct {
	// 模板等级名称，取值有：
	// <li>sup_loose：超级宽松；</li>
	// <li>loose：宽松；</li>
	// <li>emergency：紧急；</li>
	// <li>normal：适中；</li>
	// <li>strict：严格；</li>
	// <li>close：关闭，仅精准速率限制生效。</li>
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 模板处置方式，取值有：
	// <li>alg：JavaScript挑战；</li>
	// <li>monitor：观察。</li>不填写默认取alg。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 模板值详情。仅出参返回。
	RateLimitTemplateDetail *RateLimitTemplateDetail `json:"RateLimitTemplateDetail,omitnil" name:"RateLimitTemplateDetail"`
}

type RateLimitTemplateDetail struct {
	// 模板等级名称，取值有：
	// <li>sup_loose：超级宽松；</li>
	// <li>loose：宽松；</li>
	// <li>emergency：紧急；</li>
	// <li>normal：适中；</li>
	// <li>strict：严格；</li>
	// <li>close：关闭，仅精准速率限制生效。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 唯一id。
	ID *int64 `json:"ID,omitnil" name:"ID"`

	// 模板处置方式，取值有：
	// <li>alg：JavaScript挑战；</li>
	// <li>monitor：观察。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 惩罚时间，取值范围0-2天，单位秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PunishTime *int64 `json:"PunishTime,omitnil" name:"PunishTime"`

	// 统计阈值，单位是次，取值范围0-4294967294。
	Threshold *int64 `json:"Threshold,omitnil" name:"Threshold"`

	// 统计周期，取值范围0-120秒。
	Period *int64 `json:"Period,omitnil" name:"Period"`
}

type RateLimitUserRule struct {
	// 速率限制统计阈值，单位是次，取值范围0-4294967294。
	Threshold *int64 `json:"Threshold,omitnil" name:"Threshold"`

	// 速率限制统计时间，取值范围 10/20/30/40/50/60 单位是秒。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 规则名，只能以英文字符，数字，下划线组合，且不能以下划线开头。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 处置动作，取值有： <li>monitor：观察；</li> <li>drop：拦截；</li> <li>alg：JavaScript挑战。</li>	
	Action *string `json:"Action,omitnil" name:"Action"`

	// 惩罚时长，0-2天。
	PunishTime *int64 `json:"PunishTime,omitnil" name:"PunishTime"`

	// 处罚时长单位，取值有：
	// <li>second：秒；</li>
	// <li>minutes：分钟；</li>
	// <li>hour：小时。</li>
	PunishTimeUnit *string `json:"PunishTimeUnit,omitnil" name:"PunishTimeUnit"`

	// 规则状态，取值有：
	// <li>on：生效；</li>
	// <li>off：不生效。</li>默认on生效。
	RuleStatus *string `json:"RuleStatus,omitnil" name:"RuleStatus"`

	// 规则详情。
	AclConditions []*AclCondition `json:"AclConditions,omitnil" name:"AclConditions"`

	// 规则权重，取值范围0-100。
	RulePriority *int64 `json:"RulePriority,omitnil" name:"RulePriority"`

	// 规则 Id。仅出参使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleID *int64 `json:"RuleID,omitnil" name:"RuleID"`

	// 过滤词，取值有：
	// <li>sip：客户端ip。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreqFields []*string `json:"FreqFields,omitnil" name:"FreqFields"`

	// 更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 统计范围，字段为 null 时，代表 source_to_eo。取值有：
	// <li>source_to_eo：（响应）源站到EdgeOne。</li>
	// <li>client_to_eo：（请求）客户端到EdgeOne；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreqScope []*string `json:"FreqScope,omitnil" name:"FreqScope"`
}

type Resource struct {
	// 资源 ID。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 付费模式，取值有：
	// <li>0：后付费。</li>
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 生效时间。
	EnableTime *string `json:"EnableTime,omitnil" name:"EnableTime"`

	// 失效时间。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 套餐状态，取值有：
	// <li>normal：正常；</li>
	// <li>isolated：隔离；</li>
	// <li>destroyed：销毁。</li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 询价参数。
	Sv []*Sv `json:"Sv,omitnil" name:"Sv"`

	// 是否自动续费，取值有：
	// <li>0：默认状态；</li>
	// <li>1：自动续费；</li>
	// <li>2：不自动续费。</li>
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 套餐关联资源 ID。
	PlanId *string `json:"PlanId,omitnil" name:"PlanId"`

	// 地域，取值有：
	// <li>mainland：国内；</li>
	// <li>overseas：海外。</li>
	// <li>global：全球。</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 资源类型，取值有：
	// <li>plan：套餐类型；</li>
	// <li>pay-as-you-go：后付费类型。</li>
	// <li>value-added：增值服务类型。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group *string `json:"Group,omitnil" name:"Group"`

	// 当前资源绑定的站点数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneNumber *int64 `json:"ZoneNumber,omitnil" name:"ZoneNumber"`
}

type RewriteAction struct {
	// 功能名称，功能名称填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 参数。
	Parameters []*RuleRewriteActionParams `json:"Parameters,omitnil" name:"Parameters"`
}

type Rule struct {
	// 执行的功能。
	Actions []*Action `json:"Actions,omitnil" name:"Actions"`

	// 执行功能判断条件。
	// 注意：满足该数组内任意一项条件，功能即可执行。
	Conditions []*RuleAndConditions `json:"Conditions,omitnil" name:"Conditions"`

	// 嵌套规则。
	SubRules []*SubRuleItem `json:"SubRules,omitnil" name:"SubRules"`
}

type RuleAndConditions struct {
	// 规则引擎条件，该数组内所有项全部满足即判断该条件满足。
	Conditions []*RuleCondition `json:"Conditions,omitnil" name:"Conditions"`
}

type RuleChoicePropertiesItem struct {
	// 参数名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 ChoicesValue 中选择； </li>
	// <li> TOGGLE：参数值为开关类型，可在 ChoicesValue 中选择；</li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 参数值的可选值。
	// 注意：若参数值为用户自定义则该数组为空数组。
	ChoicesValue []*string `json:"ChoicesValue,omitnil" name:"ChoicesValue"`

	// 数值参数的最小值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Min *int64 `json:"Min,omitnil" name:"Min"`

	// 数值参数的最大值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Max *int64 `json:"Max,omitnil" name:"Max"`

	// 参数值是否支持多选或者填写多个。
	IsMultiple *bool `json:"IsMultiple,omitnil" name:"IsMultiple"`

	// 是否允许为空。
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitnil" name:"IsAllowEmpty"`

	// 特殊参数。
	// <li> 为 NULL：RuleAction 选择 NormalAction；</li>
	// <li> 成员参数 Id 为 Action：RuleAction 选择 RewirteAction；</li>
	// <li> 成员参数 Id 为 StatusCode：RuleAction 选择 CodeAction。</li>
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitnil" name:"ExtraParameter"`
}

type RuleCodeActionParams struct {
	// 状态 Code。
	StatusCode *int64 `json:"StatusCode,omitnil" name:"StatusCode"`

	// 参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type RuleCondition struct {
	// 运算符，取值有：
	// <li> equal: 等于； </li>
	// <li> notequal: 不等于；</li>
	// <li> exist: 存在； </li>
	// <li> notexist: 不存在。</li>
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 匹配类型，取值有：
	// <li> filename：文件名； </li>
	// <li> extension：文件后缀； </li>
	// <li> host：HOST； </li>
	// <li> full_url：URL Full，当前站点下完整 URL 路径，必须包含 HTTP 协议，Host 和 路径； </li>
	// <li> url：URL Path，当前站点下 URL 路径的请求； </li><li>client_country：客户端国家/地区；</li>
	// <li> query_string：查询字符串，当前站点下请求URL的查询字符串； </li>
	// <li> request_header：HTTP请求头部。 </li>
	Target *string `json:"Target,omitnil" name:"Target"`

	// 对应匹配类型的参数值，仅在匹配类型为查询字符串或HTTP请求头并且运算符取值为存在或不存在时允许传空数组，对应匹配类型有：
	// <li> 文件后缀：jpg、txt等文件后缀；</li>
	// <li> 文件名称：例如 foo.jpg 中的 foo；</li>
	// <li> 全部（站点任意请求）： all； </li>
	// <li> HOST：当前站点下的 host ，例如www.maxx55.com；</li>
	// <li> URL Path：当前站点下 URL 路径的请求，例如：/example；</li>
	// <li> URL Full：当前站点下完整 URL 请求，必须包含 HTTP 协议，Host 和 路径，例如：https://www.maxx55.cn/example；</li>
	// <li> 客户端国家/地区：符合ISO3166标准的国家/地区标识；</li>
	// <li> 查询字符串: 当前站点下URL请求中查询字符串的参数值，例如lang=cn&version=1中的cn和1； </li>
	// <li> HTTP 请求头: HTTP请求头部字段值，例如Accept-Language:zh-CN,zh;q=0.9中的zh-CN,zh;q=0.9。 </li>
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 是否忽略参数值的大小写，默认值为 false。
	IgnoreCase *bool `json:"IgnoreCase,omitnil" name:"IgnoreCase"`

	// 对应匹配类型的参数名称，在 Target 值为以下取值时有效，有效时值不能为空：
	// <li> query_string（查询字符串）: 当前站点下URL请求中查询字符串的参数名称，例如lang=cn&version=1中的lang和version； </li>
	// <li> request_header（HTTP 请求头）: HTTP请求头部字段名，例如Accept-Language:zh-CN,zh;q=0.9中的Accept-Language。 </li>
	Name *string `json:"Name,omitnil" name:"Name"`

	// 是否忽略参数名称的大小写，默认值为 false。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreNameCase *bool `json:"IgnoreNameCase,omitnil" name:"IgnoreNameCase"`
}

type RuleExtraParameter struct {
	// 参数名，取值有：
	// <li> Action：修改 HTTP 头部所需参数，RuleAction 选择 RewirteAction；</li>
	// <li> StatusCode：状态码相关功能所需参数，RuleAction 选择 CodeAction。</li>
	Id *string `json:"Id,omitnil" name:"Id"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 Values 中选择； </li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 可选参数值。
	// 注意：当 Id 的值为 StatusCode 时数组中的值为整型，填写参数值时请填写字符串的整型数值。
	Choices []*string `json:"Choices,omitnil" name:"Choices"`
}

type RuleItem struct {
	// 规则ID。
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称，名称字符串长度 1~255。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则状态，取值有:
	// <li> enable: 启用； </li>
	// <li> disable: 未启用。 </li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规则内容。
	Rules []*Rule `json:"Rules,omitnil" name:"Rules"`

	// 规则优先级, 值越大优先级越高，最小为 1。
	RulePriority *int64 `json:"RulePriority,omitnil" name:"RulePriority"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type RuleNormalActionParams struct {
	// 参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type RuleRewriteActionParams struct {
	// 功能参数名称，参数填写规范可调用接口 [查询规则引擎的设置参数](https://tcloud4api.woa.com/document/product/1657/79433?!preview&!document=1) 查看。现在只有三种取值：
	// <li> add：添加 HTTP 头部；</li>
	// <li> set：重写 HTTP 头部；</li>
	// <li> del：删除 HTTP 头部。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 参数名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 参数值。
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type RulesProperties struct {
	// 值为参数名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 数值参数的最小值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Min *int64 `json:"Min,omitnil" name:"Min"`

	// 参数值的可选值。
	// 注意：若参数值为用户自定义则该数组为空数组。
	ChoicesValue []*string `json:"ChoicesValue,omitnil" name:"ChoicesValue"`

	// 参数值类型。
	// <li> CHOICE：参数值只能在 ChoicesValue 中选择； </li>
	// <li> TOGGLE：参数值为开关类型，可在 ChoicesValue 中选择；</li>
	// <li> OBJECT：参数值为对象类型，ChoiceProperties 为改对象类型关联的属性；</li>
	// <li> CUSTOM_NUM：参数值用户自定义，整型类型；</li>
	// <li> CUSTOM_STRING：参数值用户自定义，字符串类型。</li>注意：当参数类型为 OBJECT 类型时，请注意参考 [示例2 参数为 OBJECT 类型的创建](https://tcloud4api.woa.com/document/product/1657/79382?!preview&!document=1)
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数值参数的最大值，非数值参数或 Min 和 Max 值都为 0 则此项无意义。
	Max *int64 `json:"Max,omitnil" name:"Max"`

	// 参数值是否支持多选或者填写多个。
	IsMultiple *bool `json:"IsMultiple,omitnil" name:"IsMultiple"`

	// 是否允许为空。
	IsAllowEmpty *bool `json:"IsAllowEmpty,omitnil" name:"IsAllowEmpty"`

	// 该参数对应的关联配置参数，属于调用接口的必填参数。
	// 注意：如果可选参数无特殊新增参数则该数组为空数组。
	ChoiceProperties []*RuleChoicePropertiesItem `json:"ChoiceProperties,omitnil" name:"ChoiceProperties"`

	// <li> 为 NULL：无特殊参数，RuleAction 选择 NormalAction；</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraParameter *RuleExtraParameter `json:"ExtraParameter,omitnil" name:"ExtraParameter"`
}

type RulesSettingAction struct {
	// 功能名称，取值有：
	// <li> 访问URL 重写（AccessUrlRedirect）；</li>
	// <li> 回源 URL 重写 （UpstreamUrlRedirect）；</li>
	// <li> 自定义错误页面
	// (ErrorPage)；</li>
	// <li> QUIC（QUIC）；</li>
	// <li> WebSocket （WebSocket）；</li>
	// <li> 视频拖拽（VideoSeek）；</li>
	// <li> Token 鉴权（Authentication）；</li>
	// <li> 自定义CacheKey（CacheKey）；</li>
	// <li> 节点缓存 TTL （Cache）；</li>
	// <li> 浏览器缓存 TTL（MaxAge）；</li>
	// <li> 离线缓存（OfflineCache）；</li>
	// <li> 智能加速（SmartRouting）；</li>
	// <li> 分片回源（RangeOriginPull）；</li>
	// <li> HTTP/2 回源（UpstreamHttp2）；</li>
	// <li> Host Header 重写（HostHeader）；</li>
	// <li> 强制 HTTPS（ForceRedirect）；</li>
	// <li> 回源 HTTPS（OriginPullProtocol）；</li>
	// <li> 缓存预刷新（CachePrefresh）；</li>
	// <li> 智能压缩（Compression）；</li>
	// <li> 修改 HTTP 请求头（RequestHeader）；</li>
	// <li> 修改HTTP响应头（ResponseHeader）;</li>
	// <li> 状态码缓存 TTL（StatusCodeCache）;</li>
	// <li> Hsts；</li>
	// <li> ClientIpHeader；</li>
	// <li> TlsVersion；</li>
	// <li> OcspStapling。</li>
	Action *string `json:"Action,omitnil" name:"Action"`

	// 参数信息。
	Properties []*RulesProperties `json:"Properties,omitnil" name:"Properties"`
}

type SecEntry struct {
	// 查询维度值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 查询维度下详细数据。
	Value []*SecEntryValue `json:"Value,omitnil" name:"Value"`
}

type SecEntryValue struct {
	// 指标名称。
	Metric *string `json:"Metric,omitnil" name:"Metric"`

	// 时序数据详情。
	Detail []*TimingDataItem `json:"Detail,omitnil" name:"Detail"`

	// 最大值。
	Max *int64 `json:"Max,omitnil" name:"Max"`

	// 平均值。
	Avg *float64 `json:"Avg,omitnil" name:"Avg"`

	// 数据总和。
	Sum *float64 `json:"Sum,omitnil" name:"Sum"`
}

type SecurityConfig struct {
	// 托管规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafConfig *WafConfig `json:"WafConfig,omitnil" name:"WafConfig"`

	// 速率限制。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitConfig *RateLimitConfig `json:"RateLimitConfig,omitnil" name:"RateLimitConfig"`

	// 自定义规则。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclConfig *AclConfig `json:"AclConfig,omitnil" name:"AclConfig"`

	// Bot配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotConfig *BotConfig `json:"BotConfig,omitnil" name:"BotConfig"`

	// 七层防护总开关。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchConfig *SwitchConfig `json:"SwitchConfig,omitnil" name:"SwitchConfig"`

	// 基础访问管控。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpTableConfig *IpTableConfig `json:"IpTableConfig,omitnil" name:"IpTableConfig"`

	// 例外规则配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptConfig *ExceptConfig `json:"ExceptConfig,omitnil" name:"ExceptConfig"`

	// 自定义拦截页面配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DropPageConfig *DropPageConfig `json:"DropPageConfig,omitnil" name:"DropPageConfig"`

	// 模板配置。此处仅出参数使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateConfig *TemplateConfig `json:"TemplateConfig,omitnil" name:"TemplateConfig"`

	// 慢速攻击配置。如果为null，默认使用历史配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowPostConfig *SlowPostConfig `json:"SlowPostConfig,omitnil" name:"SlowPostConfig"`
}

type SecurityType struct {
	// 安全类型开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type ServerCertInfo struct {
	// 服务器证书 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// 证书备注名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil" name:"Alias"`

	// 证书类型，取值有：
	// <li>default：默认证书；</li>
	// <li>upload：用户上传；</li>
	// <li>managed：腾讯云托管。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 证书过期时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 证书部署时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeployTime *string `json:"DeployTime,omitnil" name:"DeployTime"`

	// 签名算法。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignAlgo *string `json:"SignAlgo,omitnil" name:"SignAlgo"`

	// 证书归属域名名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonName *string `json:"CommonName,omitnil" name:"CommonName"`
}

type SkipCondition struct {
	// 例外跳过类型，取值为：
	// <li>header_fields：HTTP请求Header；</li>
	// <li>cookie：HTTP请求Cookie；</li>
	// <li>query_string：HTTP请求URL中的Query参数；</li>
	// <li>uri：HTTP请求URI；</li>
	// <li>body_raw：HTTP请求Body；</li>
	// <li>body_json： JSON格式的HTTP Body。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 选择跳过的字段，取值为：
	// <li>args：uri 下选择 query 参数: ?name1=jack&age=12；</li>
	// <li>path：uri 下选择部分路径：/path/to/resource.jpg；</li>
	// <li>full：uri 下选择完整路径：example.com/path/to/resource.jpg?name1=jack&age=12；</li>
	// <li>upload_filename：分段文件名，即分段传输文件时；</li>
	// <li>keys：所有的Key；</li>
	// <li>values：匹配Key对应的值；</li>
	// <li>key_value：匹配Key及匹配Value。</li>
	Selector *string `json:"Selector,omitnil" name:"Selector"`

	// 匹配Key所使用的匹配方式，取值为：
	// <li>equal：精准匹配，等于；</li>
	// <li>wildcard：通配符匹配，支持 * 通配。</li>
	MatchFromType *string `json:"MatchFromType,omitnil" name:"MatchFromType"`

	// 匹配Key的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchFrom []*string `json:"MatchFrom,omitnil" name:"MatchFrom"`

	// 匹配Content所使用的匹配方式，取值为：
	// <li>equal：精准匹配，等于；</li>
	// <li>wildcard：通配符匹配，支持 * 通配。</li>
	MatchContentType *string `json:"MatchContentType,omitnil" name:"MatchContentType"`

	// 匹配Value的值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchContent []*string `json:"MatchContent,omitnil" name:"MatchContent"`
}

type SlowPostConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 首包配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstPartConfig *FirstPartConfig `json:"FirstPartConfig,omitnil" name:"FirstPartConfig"`

	// 基础配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowRateConfig *SlowRateConfig `json:"SlowRateConfig,omitnil" name:"SlowRateConfig"`

	// 慢速攻击的处置动作，取值有：
	// <li>monitor：观察；</li>
	// <li>drop：拦截。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 本规则的Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

type SlowRateConfig struct {
	// 开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 统计的间隔，单位是秒，即在首段包传输结束后，将数据传输轴按照本参数切分，每个分片独立计算慢速攻击。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *uint64 `json:"Interval,omitnil" name:"Interval"`

	// 统计时应用的速率阈值，单位是bps，即如果本分片中的传输速率没达到本参数的值，则判定为慢速攻击，应用慢速攻击的处置方式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *uint64 `json:"Threshold,omitnil" name:"Threshold"`
}

type SmartRouting struct {
	// 智能加速配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type StandardDebug struct {
	// Debug 功能开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 允许的客户端来源。支持填写 IPV4 以及 IPV6 的 IP/IP 段，不填则表示允许任意客户端 IP。
	AllowClientIPList []*string `json:"AllowClientIPList,omitnil" name:"AllowClientIPList"`

	// Debug 功能到期时间。超出设置的时间，则功能失效。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`
}

type SubRule struct {
	// 执行功能判断条件。
	// 注意：满足该数组内任意一项条件，功能即可执行。
	Conditions []*RuleAndConditions `json:"Conditions,omitnil" name:"Conditions"`

	// 执行的功能。
	Actions []*Action `json:"Actions,omitnil" name:"Actions"`
}

type SubRuleItem struct {
	// 嵌套规则信息。
	Rules []*SubRule `json:"Rules,omitnil" name:"Rules"`

	// 规则标签。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type Sv struct {
	// 询价参数键。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 询价参数值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 询价参数映射的配额，取值有：
	// <li>zone：站点数；</li>
	// <li>custom-rule：自定义规则数；</li>
	// <li>rate-limiting-rule：速率限制规则数；</li>
	// <li>l4-proxy-instance：四层代理实例数。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pack *string `json:"Pack,omitnil" name:"Pack"`

	// 询价参数映射的四层代理实例Id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 询价参数对应的防护等级。
	// 取值有： <li> cm_30G：中国大陆加速区域保底防护30Gbps；</li><li> cm_60G：中国大陆加速区域保底防护60Gbps；</li><li> cm_100G：中国大陆加速区域保底防护100Gbps；</li><li> anycast_300G：全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> anycast_unlimited：全球加速区域（除中国大陆）Anycast无上限全力防护；</li><li> cm_30G_anycast_300G：中国大陆加速区域保底防护30Gbps，全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> cm_30G_anycast_unlimited：中国大陆加速区域保底防护30Gbps，全球加速区域（除中国大陆）Anycast无上限全力防护；</li><li> cm_60G_anycast_300G：中国大陆加速区域保底防护60Gbps，全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> cm_60G_anycast_unlimited：中国大陆加速区域保底防护60Gbps，全球加速区域（除中国大陆）Anycast无上限全力防护；</li><li> cm_100G_anycast_300G：中国大陆加速区域保底防护100Gbps，全球加速区域（除中国大陆）Anycast联防300Gbps；</li><li> cm_100G_anycast_unlimited：中国大陆加速区域保底防护100Gbps，全球加速区域（除中国大陆）Anycast无上限全力防护。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectionSpecs *string `json:"ProtectionSpecs,omitnil" name:"ProtectionSpecs"`
}

type SwitchConfig struct {
	// Web类型的安全总开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>不影响DDoS与Bot的开关。
	WebSwitch *string `json:"WebSwitch,omitnil" name:"WebSwitch"`
}

type Tag struct {
	// 标签键。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type Task struct {
	// 任务 ID。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 资源。
	Target *string `json:"Target,omitnil" name:"Target"`

	// 任务类型。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务完成时间。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type TemplateConfig struct {
	// 模板ID。
	TemplateId *string `json:"TemplateId,omitnil" name:"TemplateId"`

	// 模板名称。
	TemplateName *string `json:"TemplateName,omitnil" name:"TemplateName"`
}

type TimingDataItem struct {
	// 返回数据对应时间点，采用unix秒级时间戳。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 具体数值。
	Value *int64 `json:"Value,omitnil" name:"Value"`
}

type TimingDataRecord struct {
	// 查询维度值。
	TypeKey *string `json:"TypeKey,omitnil" name:"TypeKey"`

	// 详细时序数据。
	TypeValue []*TimingTypeValue `json:"TypeValue,omitnil" name:"TypeValue"`
}

type TimingTypeValue struct {
	// 数据和。
	Sum *int64 `json:"Sum,omitnil" name:"Sum"`

	// 最大值。
	Max *int64 `json:"Max,omitnil" name:"Max"`

	// 平均值。
	Avg *int64 `json:"Avg,omitnil" name:"Avg"`

	// 指标名。
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 详细数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*TimingDataItem `json:"Detail,omitnil" name:"Detail"`
}

type TopDataRecord struct {
	// 查询维度值。
	TypeKey *string `json:"TypeKey,omitnil" name:"TypeKey"`

	// top数据排行。
	DetailData []*TopDetailData `json:"DetailData,omitnil" name:"DetailData"`
}

type TopDetailData struct {
	// 字段名。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 字段值。
	Value *int64 `json:"Value,omitnil" name:"Value"`
}

type TopEntry struct {
	// top查询维度值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 查询具体数据。
	Value []*TopEntryValue `json:"Value,omitnil" name:"Value"`
}

type TopEntryValue struct {
	// 排序实体名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 排序实体数量。
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type UpstreamHttp2 struct {
	// http2 回源配置开关，取值有：
	// <li>on：开启；</li>
	// <li>off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`
}

type VanityNameServers struct {
	// 自定义 ns 开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 自定义 ns 列表。
	Servers []*string `json:"Servers,omitnil" name:"Servers"`
}

type VanityNameServersIps struct {
	// 自定义名字服务器名称。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 自定义名字服务器 IPv4 地址。
	IPv4 *string `json:"IPv4,omitnil" name:"IPv4"`
}

type Waf struct {
	// Waf开关，取值为：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 策略ID。
	PolicyId *int64 `json:"PolicyId,omitnil" name:"PolicyId"`
}

type WafConfig struct {
	// WafConfig开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>开关仅与配置是否生效有关，即使为off（关闭），也可以正常修改配置的内容。
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 上一次设置的防护级别，取值有：
	// <li> loose：宽松；</li>
	// <li> normal：正常；</li>
	// <li> strict：严格；</li>
	// <li> stricter：超严格；</li>
	// <li> custom：自定义。</li>
	Level *string `json:"Level,omitnil" name:"Level"`

	// 全局WAF模式，取值有：
	// <li> block：阻断（全局阻断，但可对详细规则配置观察）；</li>
	// <li> observe：观察（无论详细规则配置什么，都为观察）。</li>
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 托管规则详细配置。如果为null，默认使用历史配置。
	WafRule *WafRule `json:"WafRule,omitnil" name:"WafRule"`

	// AI规则引擎防护配置。如果为null，默认使用历史配置。
	AiRule *AiRule `json:"AiRule,omitnil" name:"AiRule"`
}

type WafRule struct {
	// 托管规则开关，取值有：
	// <li> on：开启；</li>
	// <li> off：关闭。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 黑名单ID列表，将规则ID加入本参数列表中代表该ID关闭，即该规则ID不再生效。
	BlockRuleIDs []*int64 `json:"BlockRuleIDs,omitnil" name:"BlockRuleIDs"`

	// 观察模式ID列表，将规则ID加入本参数列表中代表该ID使用观察模式生效，即该规则ID进入观察模式。
	ObserveRuleIDs []*int64 `json:"ObserveRuleIDs,omitnil" name:"ObserveRuleIDs"`
}

type WebSocket struct {
	// WebSocket 超时时间配置开关，取值有：
	// <li>on：使用Timeout作为WebSocket超时时间；</li>
	// <li>off：平台仍支持WebSocket连接，此时使用系统默认的15秒为超时时间。</li>
	Switch *string `json:"Switch,omitnil" name:"Switch"`

	// 超时时间，单位为秒，最大超时时间120秒。
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`
}

type Zone struct {
	// 站点ID。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 站点当前使用的 NS 列表。
	OriginalNameServers []*string `json:"OriginalNameServers,omitnil" name:"OriginalNameServers"`

	// 腾讯云分配的 NS 列表。
	NameServers []*string `json:"NameServers,omitnil" name:"NameServers"`

	// 站点状态，取值有：
	// <li> active：NS 已切换； </li>
	// <li> pending：NS 未切换；</li>
	// <li> moved：NS 已切走；</li>
	// <li> deactivated：被封禁。 </li>
	Status *string `json:"Status,omitnil" name:"Status"`

	// 站点接入方式，取值有
	// <li> full：NS 接入； </li>
	// <li> partial：CNAME 接入；</li>
	// <li> noDomainAccess：无域名接入。</li>
	Type *string `json:"Type,omitnil" name:"Type"`

	// 站点是否关闭。
	Paused *bool `json:"Paused,omitnil" name:"Paused"`

	// 是否开启 CNAME 加速，取值有：
	// <li> enabled：开启；</li>
	// <li> disabled：关闭。</li>
	CnameSpeedUp *string `json:"CnameSpeedUp,omitnil" name:"CnameSpeedUp"`

	// CNAME 接入状态，取值有：
	// <li> finished：站点已验证；</li>
	// <li> pending：站点验证中。</li>
	CnameStatus *string `json:"CnameStatus,omitnil" name:"CnameStatus"`

	// 资源标签列表。
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// 计费资源列表。
	Resources []*Resource `json:"Resources,omitnil" name:"Resources"`

	// 站点创建时间。
	CreatedOn *string `json:"CreatedOn,omitnil" name:"CreatedOn"`

	// 站点修改时间。
	ModifiedOn *string `json:"ModifiedOn,omitnil" name:"ModifiedOn"`

	// 站点接入地域，取值有：
	// <li> global：全球；</li>
	// <li> mainland：中国大陆；</li>
	// <li> overseas：境外区域。</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 用户自定义 NS 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServers *VanityNameServers `json:"VanityNameServers,omitnil" name:"VanityNameServers"`

	// 用户自定义 NS IP 信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VanityNameServersIps []*VanityNameServersIps `json:"VanityNameServersIps,omitnil" name:"VanityNameServersIps"`

	// 展示状态，取值有：
	// <li> active：已启用；</li>
	// <li> inactive：未生效；</li>
	// <li> paused：已停用。</li>
	ActiveStatus *string `json:"ActiveStatus,omitnil" name:"ActiveStatus"`

	// 站点别名。数字、英文、-和_组合，限制20个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasZoneName *string `json:"AliasZoneName,omitnil" name:"AliasZoneName"`

	// 是否伪站点，取值有：
	// <li> 0：非伪站点；</li>
	// <li> 1：伪站点。</li>
	IsFake *int64 `json:"IsFake,omitnil" name:"IsFake"`

	// 锁定状态，取值有：<li> enable：正常，允许进行修改操作；</li><li> disable：锁定中，不允许进行修改操作。</li>
	LockStatus *string `json:"LockStatus,omitnil" name:"LockStatus"`
}

type ZoneSetting struct {
	// 站点名称。
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 站点加速区域信息，取值有：
	// <li> mainland：中国境内加速；</li>
	// <li> overseas：中国境外加速。</li>
	Area *string `json:"Area,omitnil" name:"Area"`

	// 节点缓存键配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheKey *CacheKey `json:"CacheKey,omitnil" name:"CacheKey"`

	// Quic访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Quic *Quic `json:"Quic,omitnil" name:"Quic"`

	// POST请求传输配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostMaxSize *PostMaxSize `json:"PostMaxSize,omitnil" name:"PostMaxSize"`

	// 智能压缩配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compression *Compression `json:"Compression,omitnil" name:"Compression"`

	// Http2回源配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamHttp2 *UpstreamHttp2 `json:"UpstreamHttp2,omitnil" name:"UpstreamHttp2"`

	// 访问协议强制Https跳转配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ForceRedirect *ForceRedirect `json:"ForceRedirect,omitnil" name:"ForceRedirect"`

	// 缓存过期时间配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CacheConfig *CacheConfig `json:"CacheConfig,omitnil" name:"CacheConfig"`

	// 源站配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Origin *Origin `json:"Origin,omitnil" name:"Origin"`

	// 智能加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartRouting *SmartRouting `json:"SmartRouting,omitnil" name:"SmartRouting"`

	// 浏览器缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxAge *MaxAge `json:"MaxAge,omitnil" name:"MaxAge"`

	// 离线缓存配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineCache *OfflineCache `json:"OfflineCache,omitnil" name:"OfflineCache"`

	// WebSocket配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSocket *WebSocket `json:"WebSocket,omitnil" name:"WebSocket"`

	// 客户端IP回源请求头配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIpHeader *ClientIpHeader `json:"ClientIpHeader,omitnil" name:"ClientIpHeader"`

	// 缓存预刷新配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CachePrefresh *CachePrefresh `json:"CachePrefresh,omitnil" name:"CachePrefresh"`

	// Ipv6访问配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ipv6 *Ipv6 `json:"Ipv6,omitnil" name:"Ipv6"`

	// Https 加速配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Https *Https `json:"Https,omitnil" name:"Https"`

	// 回源时是否携带客户端IP所属地域信息的配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIpCountry *ClientIpCountry `json:"ClientIpCountry,omitnil" name:"ClientIpCountry"`

	// Grpc协议支持配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grpc *Grpc `json:"Grpc,omitnil" name:"Grpc"`

	// 图片优化相关配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageOptimize *ImageOptimize `json:"ImageOptimize,omitnil" name:"ImageOptimize"`

	// 中国大陆加速优化配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccelerateMainland *AccelerateMainland `json:"AccelerateMainland,omitnil" name:"AccelerateMainland"`

	// 标准 Debug 配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardDebug *StandardDebug `json:"StandardDebug,omitnil" name:"StandardDebug"`
}