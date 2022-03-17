package model

import (
	"encoding/json"

	"strings"
)

// 创建lb实例，支持创建或者绑定ipv4弹性公网和ipv6弹性公网
type CreateLoadBalancerOption struct {
	// 负载均衡器名称。

	Name *string `json:"name,omitempty"`
	// 负载均衡器功能说明。

	Description *string `json:"description,omitempty"`
	// 负载均衡器的虚拟IP。 1.传入vip_address时必须传入vip_subnet_cidr_id 2.不传入vip_address，自动分配虚拟IP 3.传入vip_address，需要保证该ip地址在子网中未被占用

	VipAddress *string `json:"vip_address,omitempty"`
	// 负载均衡器所在的子网ID。 说明：vpc_id , vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。

	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`
	// 双栈实例对应v6的网络id 。注：默认为空，只有开启IPv6时才会传入。  说明：vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。

	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// 负载均衡器的生产者名称。只支持vlb。

	Provider *string `json:"provider,omitempty"`
	// 四层Flavor。

	L4FlavorId *string `json:"l4_flavor_id,omitempty"`
	// 负载均衡器所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 共享型：false 保障型：true，当前只支持true。

	Guaranteed *bool `json:"guaranteed,omitempty"`
	// 实例对应的vpc属性。 说明：vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。

	VpcId *string `json:"vpc_id,omitempty"`
	// 可用区列表。默认指定所有可利用的AZ。 注： 可用AZ的查询方式可用通过调用nova接口查询 /v2/{project_id}/os-availability-zone

	AvailabilityZoneList []string `json:"availability_zone_list"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 负载均衡的标签列表。示例如下：\"tags\":[{\"key\":\"aaaa\",\"value\":\"mmmaaaaa\"}]

	Tags *[]Tag `json:"tags,omitempty"`
	// 负载均衡器的管理状态。说明：负载均衡器的管理状态。只支持设定为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 七层Flavor。

	L7FlavorId *string `json:"l7_flavor_id,omitempty"`
	// 预留资源账单信息。

	BillingInfo *string `json:"billing_info,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`
	// 公网EIP的ID，目前只支持一个

	PublicipIds *[]string `json:"publicip_ids,omitempty"`

	Publicip *CreateLoadBalancerPublicIpOption `json:"publicip,omitempty"`
	// 下联面网络id列表 若该字段不指定，在loadbalancer所属的VPC中任意选一个网络id，优选双栈网络

	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`
	// 是否启用跨VPC后端转发

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// 是否开启删除保护，默认不开启

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
}

func (o CreateLoadBalancerOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerOption", string(data)}, " ")
}
