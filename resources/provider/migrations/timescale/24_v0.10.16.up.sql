-- Autogenerated by migration tool on 2022-04-04 17:13:07
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: codepipeline.pipelines
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipelines" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "created" timestamp without time zone,
    "arn" text,
    "updated" timestamp without time zone,
    "name" text,
    "role_arn" text,
    "artifact_store_location" text,
    "artifact_store_type" text,
    "artifact_store_encryption_key_id" text,
    "artifact_store_encryption_key_type" text,
    "artifact_stores" jsonb,
    "version" integer,
    CONSTRAINT aws_codepipeline_pipelines_pk PRIMARY KEY(cq_fetch_date,arn),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_codepipeline_pipelines');
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipeline_stages" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "pipeline_cq_id" uuid,
    "stage_order" integer,
    "name" text,
    "blockers" jsonb,
    CONSTRAINT aws_codepipeline_pipeline_stages_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_codepipeline_pipeline_stages (cq_fetch_date, pipeline_cq_id);
SELECT setup_tsdb_child('aws_codepipeline_pipeline_stages', 'pipeline_cq_id', 'aws_codepipeline_pipelines', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipeline_stage_actions" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "pipeline_stage_cq_id" uuid,
    "category" text,
    "owner" text,
    "provider" text,
    "version" text,
    "name" text,
    "configuration" jsonb,
    "input_artifacts" text[],
    "namespace" text,
    "output_artifacts" text[],
    "region" text,
    "role_arn" text,
    "run_order" integer,
    CONSTRAINT aws_codepipeline_pipeline_stage_actions_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_codepipeline_pipeline_stage_actions (cq_fetch_date, pipeline_stage_cq_id);
SELECT setup_tsdb_child('aws_codepipeline_pipeline_stage_actions', 'pipeline_stage_cq_id', 'aws_codepipeline_pipeline_stages', 'cq_id');

-- Resource: iam.virtual_mfa_devices
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices DROP CONSTRAINT aws_iam_virtual_mfa_devices_pk;
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices ADD CONSTRAINT aws_iam_virtual_mfa_devices_pk PRIMARY KEY (cq_fetch_date,serial_number,enable_date);

-- Resource: mq.brokers
ALTER TABLE IF EXISTS aws_mq_brokers
    DROP CONSTRAINT aws_mq_brokers_pk;
ALTER TABLE IF EXISTS "aws_mq_brokers"
    RENAME COLUMN broker_id TO id;
ALTER TABLE IF EXISTS aws_mq_brokers
    ADD CONSTRAINT aws_mq_brokers_pk PRIMARY KEY (cq_fetch_date, account_id, id);
ALTER TABLE IF EXISTS aws_mq_broker_configurations
    DROP CONSTRAINT aws_mq_broker_configurations_pk;
ALTER TABLE IF EXISTS aws_mq_broker_configurations
    ADD CONSTRAINT aws_mq_broker_configurations_pk PRIMARY KEY (cq_fetch_date, cq_id);
CREATE TABLE IF NOT EXISTS "aws_mq_broker_configuration_revisions"
(
    "cq_id"                      uuid                        NOT NULL,
    "cq_meta"                    jsonb,
    "cq_fetch_date"              timestamp without time zone NOT NULL,
    "broker_configuration_cq_id" uuid,
    "configuration_id"           text,
    "created"                    timestamp without time zone,
    "data"                       jsonb,
    "description"                text,
    CONSTRAINT aws_mq_broker_configuration_revisions_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_mq_broker_configuration_revisions (cq_fetch_date, broker_configuration_cq_id);
SELECT setup_tsdb_child('aws_mq_broker_configuration_revisions', 'broker_configuration_cq_id',
                        'aws_mq_broker_configurations', 'cq_id');
ALTER TABLE IF EXISTS aws_mq_broker_users
    DROP CONSTRAINT aws_mq_broker_users_pk;
ALTER TABLE IF EXISTS aws_mq_broker_users
    ADD CONSTRAINT aws_mq_broker_users_pk PRIMARY KEY (cq_fetch_date, cq_id);

-- Resource: sns.subscriptions
ALTER TABLE IF EXISTS aws_sns_subscriptions DROP CONSTRAINT aws_sns_subscriptions_pk;
ALTER TABLE IF EXISTS aws_sns_subscriptions ADD CONSTRAINT aws_sns_subscriptions_pk PRIMARY KEY (cq_fetch_date,endpoint,owner,protocol,arn,topic_arn);

-- Resource: ecs.clusters
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_tasks"
(
    "cq_id"                         uuid                        NOT NULL,
    "cq_meta"                       jsonb,
    "cq_fetch_date"                 timestamp without time zone NOT NULL,
    "cluster_cq_id"                 uuid,
    "attributes"                    jsonb,
    "availability_zone"             text,
    "capacity_provider_name"        text,
    "cluster_arn"                   text,
    "connectivity"                  text,
    "connectivity_at"               timestamp without time zone,
    "container_instance_arn"        text,
    "cpu"                           text,
    "created_at"                    timestamp without time zone,
    "desired_status"                text,
    "enable_execute_command"        boolean,
    "ephemeral_storage_size_in_gib" integer,
    "execution_stopped_at"          timestamp without time zone,
    "group"                         text,
    "health_status"                 text,
    "inference_accelerators"        jsonb,
    "last_status"                   text,
    "launch_type"                   text,
    "memory"                        text,
    "overrides"                     jsonb,
    "platform_family"               text,
    "platform_version"              text,
    "pull_started_at"               timestamp without time zone,
    "pull_stopped_at"               timestamp without time zone,
    "started_at"                    timestamp without time zone,
    "started_by"                    text,
    "stop_code"                     text,
    "stopped_at"                    timestamp without time zone,
    "stopped_reason"                text,
    "stopping_at"                   timestamp without time zone,
    "tags"                          jsonb,
    "arn"                           text,
    "task_definition_arn"           text,
    "version"                       bigint,
    CONSTRAINT aws_ecs_cluster_tasks_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_tasks (cq_fetch_date, cluster_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_tasks', 'cluster_cq_id', 'aws_ecs_clusters', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_task_attachments"
(
    "cq_id"              uuid                        NOT NULL,
    "cq_meta"            jsonb,
    "cq_fetch_date"      timestamp without time zone NOT NULL,
    "cluster_task_cq_id" uuid,
    "details"            jsonb,
    "id"                 text,
    "status"             text,
    "type"               text,
    CONSTRAINT aws_ecs_cluster_task_attachments_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_task_attachments (cq_fetch_date, cluster_task_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_task_attachments', 'cluster_task_cq_id', 'aws_ecs_cluster_tasks', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_task_containers"
(
    "cq_id"              uuid                        NOT NULL,
    "cq_meta"            jsonb,
    "cq_fetch_date"      timestamp without time zone NOT NULL,
    "cluster_task_cq_id" uuid,
    "container_arn"      text,
    "cpu"                text,
    "exit_code"          integer,
    "gpu_ids"            text[],
    "health_status"      text,
    "image"              text,
    "image_digest"       text,
    "last_status"        text,
    "managed_agents"     jsonb,
    "memory"             text,
    "memory_reservation" text,
    "name"               text,
    "network_bindings"   jsonb,
    "network_interfaces" jsonb,
    "reason"             text,
    "runtime_id"         text,
    "task_arn"           text,
    CONSTRAINT aws_ecs_cluster_task_containers_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_task_containers (cq_fetch_date, cluster_task_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_task_containers', 'cluster_task_cq_id', 'aws_ecs_cluster_tasks', 'cq_id');
ALTER TABLE IF EXISTS "aws_ecs_cluster_services"
    ADD COLUMN IF NOT EXISTS "platform_family" text;
ALTER TABLE IF EXISTS "aws_ecs_cluster_service_deployments"
    ADD COLUMN IF NOT EXISTS "platform_family" text;
ALTER TABLE IF EXISTS "aws_ecs_cluster_service_task_sets"
    ADD COLUMN IF NOT EXISTS "platform_family" text;
ALTER TABLE IF EXISTS "aws_ecs_cluster_container_instances"
    ADD COLUMN IF NOT EXISTS "health_status_overall_status" text;
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_container_instance_health_status_details"
(
    "cq_id"                            uuid                        NOT NULL,
    "cq_meta"                          jsonb,
    "cq_fetch_date"                    timestamp without time zone NOT NULL,
    "cluster_container_instance_cq_id" uuid,
    "last_status_change"               timestamp without time zone,
    "last_updated"                     timestamp without time zone,
    "status"                           text,
    "type"                             text,
    CONSTRAINT aws_ecs_cluster_container_instance_health_status_details_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_container_instance_health_status_details (cq_fetch_date, cluster_container_instance_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_container_instance_health_status_details', 'cluster_container_instance_cq_id',
                        'aws_ecs_cluster_container_instances', 'cq_id');

-- Resource: backup.plans
CREATE TABLE IF NOT EXISTS "aws_backup_plans" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "account_id" text,
    "region" text,
    "arn" text,
    "id" text,
    "name" text,
    "creation_date" timestamp without time zone,
    "creator_request_id" text,
    "last_execution_date" timestamp without time zone,
    "version_id" text,
    "advanced_backup_settings" jsonb,
    "tags" jsonb,
    CONSTRAINT aws_backup_plans_pk PRIMARY KEY(cq_fetch_date,arn),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_backup_plans');
CREATE TABLE IF NOT EXISTS "aws_backup_plan_rules" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "plan_cq_id" uuid,
    "name" text,
    "target_backup_vault_name" text,
    "completion_window_minutes" bigint,
    "copy_actions" jsonb,
    "enable_continuous_backup" boolean,
    "delete_after_days" bigint,
    "move_to_cold_storage_after_days" bigint,
    "recovery_point_tags" jsonb,
    "id" text,
    "schedule_expression" text,
    "start_window_minutes" bigint,
    CONSTRAINT aws_backup_plan_rules_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_backup_plan_rules (cq_fetch_date, plan_cq_id);
SELECT setup_tsdb_child('aws_backup_plan_rules', 'plan_cq_id', 'aws_backup_plans', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_backup_plan_selections" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "plan_cq_id" uuid,
    "creation_date" timestamp without time zone,
    "creator_request_id" text,
    "iam_role_arn" text,
    "selection_id" text,
    "selection_name" text,
    "conditions" jsonb,
    "list_of_tags" jsonb,
    "not_resources" text[],
    "resources" text[],
    CONSTRAINT aws_backup_plan_selections_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_backup_plan_selections (cq_fetch_date, plan_cq_id);
SELECT setup_tsdb_child('aws_backup_plan_selections', 'plan_cq_id', 'aws_backup_plans', 'cq_id');

-- Resource: backup.vaults
CREATE TABLE IF NOT EXISTS "aws_backup_vaults" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "account_id" text,
    "region" text,
    "arn" text,
    "name" text,
    "creation_date" timestamp without time zone,
    "creator_request_id" text,
    "encryption_key_arn" text,
    "lock_date" timestamp without time zone,
    "locked" boolean,
    "max_retention_days" bigint,
    "min_retention_days" bigint,
    "number_of_recovery_points" bigint,
    "access_policy" jsonb,
    "notification_events" text[],
    "notification_sns_topic_arn" text,
    "tags" jsonb,
    CONSTRAINT aws_backup_vaults_pk PRIMARY KEY(cq_fetch_date,arn),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_backup_vaults');
CREATE TABLE IF NOT EXISTS "aws_backup_vault_recovery_points" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "vault_cq_id" uuid,
    "backup_size" bigint,
    "calculated_delete_at" timestamp without time zone,
    "calculated_move_to_cold_storage_at" timestamp without time zone,
    "completion_date" timestamp without time zone,
    "created_by" jsonb,
    "creation_date" timestamp without time zone,
    "encryption_key_arn" text,
    "iam_role_arn" text,
    "is_encrypted" boolean,
    "last_restore_time" timestamp without time zone,
    "delete_after" bigint,
    "move_to_cold_storage_after" bigint,
    "arn" text,
    "resource_arn" text,
    "resource_type" text,
    "source_backup_vault_arn" text,
    "status" text,
    "status_message" text,
    "tags" jsonb,
    CONSTRAINT aws_backup_vault_recovery_points_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_backup_vault_recovery_points (cq_fetch_date, vault_cq_id);
SELECT setup_tsdb_child('aws_backup_vault_recovery_points', 'vault_cq_id', 'aws_backup_vaults', 'cq_id');

-- Resource: codepipeline.webhooks
CREATE TABLE IF NOT EXISTS "aws_codepipeline_webhooks" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "account_id" text,
    "region" text,
    "authentication" text,
    "authentication_allowed_ip_range" text,
    "authentication_secret_token" text,
    "name" text,
    "target_action" text,
    "target_pipeline" text,
    "url" text,
    "arn" text,
    "error_code" text,
    "error_message" text,
    "last_triggered" timestamp without time zone,
    "tags" jsonb,
    CONSTRAINT aws_codepipeline_webhooks_pk PRIMARY KEY(cq_fetch_date,arn),
    UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_codepipeline_webhooks');
CREATE TABLE IF NOT EXISTS "aws_codepipeline_webhook_filters" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp without time zone NOT NULL,
    "webhook_cq_id" uuid,
    "json_path" text,
    "match_equals" text,
    CONSTRAINT aws_codepipeline_webhook_filters_pk PRIMARY KEY(cq_fetch_date,cq_id),
    UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_codepipeline_webhook_filters (cq_fetch_date, webhook_cq_id);
SELECT setup_tsdb_child('aws_codepipeline_webhook_filters', 'webhook_cq_id', 'aws_codepipeline_webhooks', 'cq_id');