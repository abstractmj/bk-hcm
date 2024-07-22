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

package monthtask

import (
	rawjson "encoding/json"

	"hcm/pkg/api/data-service/bill"
	"hcm/pkg/kit"
)

func newGcpRunner() *Gcp {
	return &Gcp{}
}

type Gcp struct {
}

func (gcp *Gcp) GetBatchSize(kt *kit.Kit) uint64 {
	return 1000
}

func (gcp *Gcp) Pull(kt *kit.Kit, rootAccountID string, billYear, billMonth int, index uint64) (
	itemList []bill.RawBillItem, isFinished bool, err error) {
	return nil, false, nil
}
func (gcp *Gcp) Split(kt *kit.Kit, rootAccountID string,
	rawItemList []*bill.RawBillItem) ([]bill.BillItemCreateReq[rawjson.RawMessage], error) {
	return nil, nil
}
