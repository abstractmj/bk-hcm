/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package huawei

import (
	typesinstancetype "hcm/pkg/adaptor/types/instance-type"
	"hcm/pkg/kit"
	"hcm/pkg/logs"

	"github.com/TencentBlueKing/gopkg/conv"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
)

// ListInstanceType ...
// reference: https://support.huaweicloud.com/api-ecs/zh-cn_topic_0020212656.html
func (h *HuaWei) ListInstanceType(kt *kit.Kit, opt *typesinstancetype.HuaWeiInstanceTypeListOption) (
	[]*typesinstancetype.HuaWeiInstanceType, error,
) {

	client, err := h.clientSet.ecsClient(opt.Region)
	if err != nil {
		return nil, err
	}

	req := &model.ListFlavorsRequest{
		AvailabilityZone: &opt.Zone,
	}

	resp, err := client.ListFlavors(req)
	if err != nil {
		logs.Errorf("list huawei instance type failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	its := make([]*typesinstancetype.HuaWeiInstanceType, 0, len(*resp.Flavors))
	for _, flavor := range *resp.Flavors {
		its = append(its, toHuaweiInstanceType(flavor))
	}

	return its, nil
}

func toHuaweiInstanceType(flavor model.Flavor) *typesinstancetype.HuaWeiInstanceType {
	cpu, _ := conv.ToInt64(flavor.Vcpus)
	return &typesinstancetype.HuaWeiInstanceType{
		InstanceType: flavor.Id,
		CPU:          cpu,
		Memory:       int64(flavor.Ram),
	}
}
