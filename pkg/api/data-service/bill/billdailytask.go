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

package bill

import (
	"hcm/pkg/api/core"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/validator"
	"hcm/pkg/dal/table/types"
	"hcm/pkg/runtime/filter"

	"github.com/shopspring/decimal"
)

// BillDailyPullTaskCreateReq ...
type BillDailyPullTaskCreateReq struct {
	FirstAccountID  string          `json:"first_account_id" validate:"required"`
	SecondAccountID string          `json:"second_account_id" validate:"required"`
	Vendor          enumor.Vendor   `json:"vendor" validate:"required"`
	ProductID       int64           `json:"product_id" validate:"omitempty"`
	BkBizID         int64           `json:"bk_biz_id" validate:"omitempty"`
	BillYear        int             `json:"bill_year" validate:"required"`
	BillMonth       int             `json:"bill_month" validate:"required"`
	BillDay         int             `json:"bill_day" validate:"required"`
	VersionID       string          `json:"version_id" validate:"required"`
	State           string          `json:"state" vaildate:"required"`
	Message         string          `json:"message" validate:"required"`
	Count           int64           `json:"count" validate:"required"`
	Currency        string          `json:"currency" validate:"required"`
	Cost            decimal.Decimal `json:"cost" validate:"required"`
}

// Validate BillDailyPullTaskCreateReq.
func (c *BillDailyPullTaskCreateReq) Validate() error {
	return validator.Validate.Struct(c)
}

// BillDailyPullTaskListReq ...
type BillDailyPullTaskListReq struct {
	Filter *filter.Expression `json:"filter" validate:"required"`
	Page   *core.BasePage     `json:"page" validate:"required"`
	Fields []string           `json:"fields" validate:"omitempty"`
}

// Validate ...
func (req *BillDailyPullTaskListReq) Validate() error {
	return validator.Validate.Struct(req)
}

// BillDailyPullTaskListResult ...
type BillDailyPullTaskListResult struct {
	Count   *uint64                    `json:"count,omitempty"`
	Details []*BillDailyPullTaskResult `json:"details"`
}

// BillDailyPullTaskResult ...
type BillDailyPullTaskResult struct {
	ID              string          `json:"id,omitempty"`
	FirstAccountID  string          `json:"first_account_id" validate:"required"`
	SecondAccountID string          `json:"second_account_id" validate:"required"`
	Vendor          enumor.Vendor   `json:"vendor" validate:"required"`
	ProductID       int64           `json:"product_id" validate:"omitempty"`
	BkBizID         int64           `json:"bk_biz_id" validate:"omitempty"`
	BillYear        int             `json:"bill_year" validate:"required"`
	BillMonth       int             `json:"bill_month" validate:"required"`
	BillDay         int             `json:"bill_day" validate:"required"`
	VersionID       string          `json:"version_id" validate:"required"`
	State           string          `json:"state" vaildate:"required"`
	Message         string          `json:"message" validate:"required"`
	Count           int64           `json:"count" validate:"required"`
	Currency        string          `json:"currency" validate:"required"`
	Cost            decimal.Decimal `json:"cost" validate:"required"`
	CreatedAt       types.Time      `json:"created_at,omitempty"`
	UpdatedAt       types.Time      `json:"updated_at,omitempty"`
}

// BillDailyPullTaskUpdateReq ...
type BillDailyPullTaskUpdateReq struct {
	ID              string          `json:"id,omitempty" validate:"required"`
	FirstAccountID  string          `json:"first_account_id" validate:"required"`
	SecondAccountID string          `json:"second_account_id" validate:"required"`
	Vendor          enumor.Vendor   `json:"vendor" validate:"required"`
	ProductID       int64           `json:"product_id" validate:"omitempty"`
	BkBizID         int64           `json:"bk_biz_id" validate:"omitempty"`
	BillYear        int             `json:"bill_year" validate:"required"`
	BillMonth       int             `json:"bill_month" validate:"required"`
	BillDay         int             `json:"bill_day" validate:"required"`
	VersionID       string          `json:"version_id" validate:"required"`
	State           string          `json:"state" vaildate:"required"`
	Message         string          `json:"message" validate:"required"`
	Count           int64           `json:"count" validate:"required"`
	Currency        string          `json:"currency" validate:"required"`
	Cost            decimal.Decimal `json:"cost" validate:"required"`
}

// Validate ...
func (req *BillDailyPullTaskUpdateReq) Validate() error {
	return validator.Validate.Struct(req)
}
