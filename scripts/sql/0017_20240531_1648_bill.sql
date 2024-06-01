-- 账单同步器表
create table if not exists `account_bill_puller` (
  `id` varchar(64) not null,
  `first_account_id` varchar(64) not null,
  `second_account_id` varchar(64) not null,
  `vendor` varchar(16) not null,
  `product_id` bigint(1),
  `bk_biz_id` bigint(1),
  `pull_mode` varchar(64) not null,
  `sync_period` varchar(64) not null,
  `bill_delay` varchar(64) not null,
  `final_bill_calendar_date` bigint(1) not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  unique key `idx_account_id` (`first_account_id`, `second_account_id`)
) engine = innodb default charset = utf8mb4;

insert into id_generator(`resource`, `max_id`)
values ('account_bill_puller', '0');

-- 月账单汇总表
create table if not exists `account_bill_summary` (
  `id` varchar(64) not null,
  `first_account_id` varchar(64) not null,
  `second_account_id` varchar(64) not null,
  `vendor` varchar(16) not null,
  `product_id` bigint(1),
  `bk_biz_id` bigint(1),
  `bill_year` bigint(1) not null,
  `bill_month` tinyint(1) not null,
  `current_version` varchar(64) not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  unique key `idx_bill_date` (`first_account_id`, `second_account_id`, `bill_year`, `bill_month`)
) engine = innodb default charset = utf8mb4;

insert into id_generator(`resource`, `max_id`)
values ('account_bill_summary', '0');

-- 月账单汇总版本表
create table if not exists `account_bill_summary_version` (
  `id` varchar(64) not null,
  `first_account_id` varchar(64) not null,
  `second_account_id` varchar(64) not null,
  `vendor` varchar(16) not null,
  `product_id` bigint(1),
  `bk_biz_id` bigint(1),
  `bill_year` bigint(1) not null,
  `bill_month` tinyint(1) not null,
  `version_id` varchar(64) not null,
  `currency` varchar(64) not null,
  `cost` decimal not null,
  `rmb_cost` decimal not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  unique key `idx_bill_date_version` (`first_account_id`, `second_account_id`, `bill_year`, `bill_month`, `version_id`)
) engine = innodb default charset = utf8mb4;

insert into id_generator(`resource`, `max_id`)
values ('account_bill_summary_version', '0');

-- 每日账单汇总表
create table if not exists `account_bill_summary_daily` (
  `id` varchar(64) not null,
  `first_account_id` varchar(64) not null,
  `second_account_id` varchar(64) not null,
  `vendor` varchar(16) not null,
  `product_id` bigint(1),
  `bk_biz_id` bigint(1),
  `bill_year` bigint(1) not null,
  `bill_month` tinyint(1) not null,
  `bill_day` tinyint(1) not null,
  `version_id` varchar(64) not null,
  `currency` varchar(64) not null,
  `cost` decimal not null,
  `rmb_cost` decimal not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  unique key `idx_bill_date_version` (`first_account_id`, `second_account_id`, `bill_year`, `bill_month`, `bill_day`, `version_id`)
) engine = innodb default charset = utf8mb4;

insert into id_generator(`resource`, `max_id`)
values ('account_bill_summary_daily', '0');

-- 分账后的明细数据
create table if not exists `account_bill_item` (
  `id` varchar(64) not null,
  `first_account_id` varchar(64) not null,
  `second_account_id` varchar(64) not null,
  `vendor` varchar(16) not null,
  `product_id` bigint(1),
  `bk_biz_id` bigint(1),
  `bill_year` bigint(1) not null,
  `bill_month` tinyint(1) not null,
  `bill_day` tinyint(1) not null,
  `version_id` varchar(64) not null,
  `currency` varchar(64) not null,
  `cost` decimal not null,
  `rmb_cost` decimal not null,
  `hc_product_code` varchar(128),
  `hc_product_name` varchar(128),
  `res_amount` decimal,
  `res_amount_unit` varchar(64),
  `extension` json,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`)
) engine = innodb default charset = utf8mb4;

insert into id_generator(`resource`, `max_id`)
values ('account_bill_item', '0');

-- 调账明细数据
create table if not exists `account_bill_adjustment_item` (
  `id` varchar(64) not null,
  `first_account_id` varchar(64) not null,
  `second_account_id` varchar(64) not null,
  `vendor` varchar(16) not null,
  `product_id` bigint(1),
  `bk_biz_id` bigint(1),
  `bill_year` bigint(1) not null,
  `bill_month` tinyint(1) not null,
  `bill_day` tinyint(1) not null,
  `type` varchar(64) not null,
  `memo` varchar(255) default '',
  `operator` varchar(64) not null,
  `currency` varchar(64) not null,
  `cost` decimal not null,
  `rmb_cost` decimal not null,
  `state` varchar(64) not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`)
) engine = innodb default charset = utf8mb4;

insert into id_generator(`resource`, `max_id`)
values ('account_bill_adjustment_item', '0');

-- 每日拉取任务状态表
create table if not exists `account_bill_daily_pull_task` (
  `id` varchar(64) not null,
  `first_account_id` varchar(64) not null,
  `second_account_id` varchar(64) not null,
  `product_id` bigint(1),
  `bk_biz_id` bigint(1),
  `bill_year` bigint(1) not null,
  `bill_month` tinyint(1) not null,
  `bill_day` tinyint(1) not null,
  `version_id` varchar(64) not null,
  `state` varchar(64) not null,
  `message` varchar(255) default '',
  `count` bigint(1) not null,
  `currency` varchar(64) not null,
  `cost` decimal not null,
  `created_at` timestamp not null default current_timestamp,
  `updated_at` timestamp not null default current_timestamp on update current_timestamp,
  primary key (`id`),
  unique key `idx_bill_date_version` (`first_account_id`, `second_account_id`, `bill_year`, `bill_month`, `bill_day`, `version_id`)
) engine = innodb default charset = utf8mb4;

insert into id_generator(`resource`, `max_id`)
values ('account_bill_daily_pull_task', '0');