/*
 Navicat Premium Dump SQL

 Source Server         : pg_dev
 Source Server Type    : PostgreSQL
 Source Server Version : 170005 (170005)
 Source Host           : 10.118.141.29:8432
 Source Catalog        : db_default
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 170005 (170005)
 File Encoding         : 65001

 Date: 11/08/2025 16:44:42
*/


-- ----------------------------
-- Table structure for app_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."app_20250811";
CREATE TABLE "public"."app_20250811" (
  "id" int8,
  "logo" varchar(254) COLLATE "pg_catalog"."default",
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "status" int4,
  "disable_des" text COLLATE "pg_catalog"."default",
  "version" varchar(254) COLLATE "pg_catalog"."default",
  "copyright" varchar(254) COLLATE "pg_catalog"."default",
  "extra" text COLLATE "pg_catalog"."default",
  "allow_regis" int4,
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."app_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of app_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."app_20250811" ("id", "logo", "keyword", "name", "status", "disable_des", "version", "copyright", "extra", "allow_regis", "description", "created_at", "updated_at") VALUES (4, '6dc607ee0b87559d8932377d46b9a3ea', 'dangyuan', '党员之家', 1, NULL, 'v1.0', 'limes-cloud@q.com', NULL, 1, '党员之家,幸福你我他', 1751047844, 1751118163);
COMMIT;

-- ----------------------------
-- Table structure for app_channel_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."app_channel_20250811";
CREATE TABLE "public"."app_channel_20250811" (
  "id" int8,
  "app_id" int8,
  "channel_id" int8
)
;
ALTER TABLE "public"."app_channel_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of app_channel_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."app_channel_20250811" ("id", "app_id", "channel_id") VALUES (7, 4, 1);
INSERT INTO "public"."app_channel_20250811" ("id", "app_id", "channel_id") VALUES (9, 4, 2);
INSERT INTO "public"."app_channel_20250811" ("id", "app_id", "channel_id") VALUES (8, 4, 3);
COMMIT;

-- ----------------------------
-- Table structure for app_field_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."app_field_20250811";
CREATE TABLE "public"."app_field_20250811" (
  "id" int8,
  "app_id" int8,
  "field_id" int8
)
;
ALTER TABLE "public"."app_field_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of app_field_20250811
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."casbin_rule_20250811";
CREATE TABLE "public"."casbin_rule_20250811" (
  "id" int8,
  "ptype" varchar(254) COLLATE "pg_catalog"."default",
  "v0" varchar(254) COLLATE "pg_catalog"."default",
  "v1" varchar(254) COLLATE "pg_catalog"."default",
  "v2" varchar(254) COLLATE "pg_catalog"."default",
  "v3" varchar(254) COLLATE "pg_catalog"."default",
  "v4" varchar(254) COLLATE "pg_catalog"."default",
  "v5" varchar(254) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."casbin_rule_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of casbin_rule_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3464, 'p', '3', '/ai-agent/api/v1/assessment', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3462, 'p', '3', '/ai-agent/api/v1/assessment', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3463, 'p', '3', '/ai-agent/api/v1/assessment', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3461, 'p', '3', '/ai-agent/api/v1/assessments', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3469, 'p', '3', '/ai-agent/api/v1/chat', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3465, 'p', '3', '/ai-agent/api/v1/chat', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3468, 'p', '3', '/ai-agent/api/v1/chat', 'update', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3467, 'p', '3', '/ai-agent/api/v1/chat/conversation', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3466, 'p', '3', '/ai-agent/api/v1/chat/messages', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3460, 'p', '3', '/ai-agent/api/v1/knowledge', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3458, 'p', '3', '/ai-agent/api/v1/knowledge', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3459, 'p', '3', '/ai-agent/api/v1/knowledge', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3457, 'p', '3', '/ai-agent/api/v1/knowledges', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3452, 'p', '3', '/ai-agent/api/v1/model', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3450, 'p', '3', '/ai-agent/api/v1/model', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3451, 'p', '3', '/ai-agent/api/v1/model', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3449, 'p', '3', '/ai-agent/api/v1/models', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3448, 'p', '3', '/ai-agent/api/v1/secret', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3446, 'p', '3', '/ai-agent/api/v1/secret', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3447, 'p', '3', '/ai-agent/api/v1/secret', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3444, 'p', '3', '/ai-agent/api/v1/secret_group', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3442, 'p', '3', '/ai-agent/api/v1/secret_group', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3443, 'p', '3', '/ai-agent/api/v1/secret_group', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3441, 'p', '3', '/ai-agent/api/v1/secret_groups', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3445, 'p', '3', '/ai-agent/api/v1/secrets', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3456, 'p', '3', '/ai-agent/api/v1/tool', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3454, 'p', '3', '/ai-agent/api/v1/tool', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3455, 'p', '3', '/ai-agent/api/v1/tool', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3453, 'p', '3', '/ai-agent/api/v1/tools', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3383, 'p', '3', '/configure/api/v1/business', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3381, 'p', '3', '/configure/api/v1/business', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3382, 'p', '3', '/configure/api/v1/business', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3380, 'p', '3', '/configure/api/v1/businesses', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3400, 'p', '3', '/configure/api/v1/configure', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3399, 'p', '3', '/configure/api/v1/configure/compare', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3371, 'p', '3', '/configure/api/v1/env', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3368, 'p', '3', '/configure/api/v1/env', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3369, 'p', '3', '/configure/api/v1/env', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3370, 'p', '3', '/configure/api/v1/env/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3372, 'p', '3', '/configure/api/v1/env/token', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3373, 'p', '3', '/configure/api/v1/env/token', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3366, 'p', '3', '/configure/api/v1/envs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3389, 'p', '3', '/configure/api/v1/resource', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3387, 'p', '3', '/configure/api/v1/resource', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3388, 'p', '3', '/configure/api/v1/resource', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3386, 'p', '3', '/configure/api/v1/resources', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3379, 'p', '3', '/configure/api/v1/server', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3376, 'p', '3', '/configure/api/v1/server', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3377, 'p', '3', '/configure/api/v1/server', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3378, 'p', '3', '/configure/api/v1/server/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3374, 'p', '3', '/configure/api/v1/servers', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3393, 'p', '3', '/configure/api/v1/template', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3395, 'p', '3', '/configure/api/v1/template', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3396, 'p', '3', '/configure/api/v1/template/compare', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3394, 'p', '3', '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3398, 'p', '3', '/configure/api/v1/template/preview', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3397, 'p', '3', '/configure/api/v1/template/switch', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3392, 'p', '3', '/configure/api/v1/templates', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3384, 'p', '3', '/configure/business/values', 'get', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3385, 'p', '3', '/configure/business/values', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3390, 'p', '3', '/configure/resource/values', 'get', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3391, 'p', '3', '/configure/resource/values', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3422, 'p', '3', '/cron/api/v1/log', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3421, 'p', '3', '/cron/api/v1/logs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3419, 'p', '3', '/cron/api/v1/task', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3415, 'p', '3', '/cron/api/v1/task', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3418, 'p', '3', '/cron/api/v1/task', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3413, 'p', '3', '/cron/api/v1/task_group', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3411, 'p', '3', '/cron/api/v1/task_group', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3412, 'p', '3', '/cron/api/v1/task_group', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3410, 'p', '3', '/cron/api/v1/task_groups', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3417, 'p', '3', '/cron/api/v1/task/cancel', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3416, 'p', '3', '/cron/api/v1/task/exec', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3420, 'p', '3', '/cron/api/v1/task/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3414, 'p', '3', '/cron/api/v1/tasks', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3408, 'p', '3', '/cron/api/v1/worker', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3406, 'p', '3', '/cron/api/v1/worker', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3407, 'p', '3', '/cron/api/v1/worker', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3404, 'p', '3', '/cron/api/v1/worker_group', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3402, 'p', '3', '/cron/api/v1/worker_group', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3403, 'p', '3', '/cron/api/v1/worker_group', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3401, 'p', '3', '/cron/api/v1/worker_groups', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3409, 'p', '3', '/cron/api/v1/worker/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3405, 'p', '3', '/cron/api/v1/workers', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3346, 'p', '3', '/manager/api/v1/app', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3343, 'p', '3', '/manager/api/v1/app', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3344, 'p', '3', '/manager/api/v1/app', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3345, 'p', '3', '/manager/api/v1/app/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3341, 'p', '3', '/manager/api/v1/apps', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3335, 'p', '3', '/manager/api/v1/channel', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3333, 'p', '3', '/manager/api/v1/channel', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3334, 'p', '3', '/manager/api/v1/channel', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3332, 'p', '3', '/manager/api/v1/channels', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3311, 'p', '3', '/manager/api/v1/department', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3309, 'p', '3', '/manager/api/v1/department', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3310, 'p', '3', '/manager/api/v1/department', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3314, 'p', '3', '/manager/api/v1/department_classify', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3312, 'p', '3', '/manager/api/v1/department_classify', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3313, 'p', '3', '/manager/api/v1/department_classify', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3308, 'p', '3', '/manager/api/v1/departments', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3295, 'p', '3', '/manager/api/v1/dictionaries', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3298, 'p', '3', '/manager/api/v1/dictionary', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3296, 'p', '3', '/manager/api/v1/dictionary', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3297, 'p', '3', '/manager/api/v1/dictionary', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3303, 'p', '3', '/manager/api/v1/dictionary_value', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3300, 'p', '3', '/manager/api/v1/dictionary_value', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3301, 'p', '3', '/manager/api/v1/dictionary_value', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3302, 'p', '3', '/manager/api/v1/dictionary_value/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3299, 'p', '3', '/manager/api/v1/dictionary_values', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3354, 'p', '3', '/manager/api/v1/feedback', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3352, 'p', '3', '/manager/api/v1/feedback', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3353, 'p', '3', '/manager/api/v1/feedback', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3347, 'p', '3', '/manager/api/v1/feedback_categories', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3350, 'p', '3', '/manager/api/v1/feedback_category', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3348, 'p', '3', '/manager/api/v1/feedback_category', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3349, 'p', '3', '/manager/api/v1/feedback_category', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3351, 'p', '3', '/manager/api/v1/feedbacks', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3340, 'p', '3', '/manager/api/v1/field', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3337, 'p', '3', '/manager/api/v1/field', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3338, 'p', '3', '/manager/api/v1/field', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3339, 'p', '3', '/manager/api/v1/field/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3336, 'p', '3', '/manager/api/v1/fields', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3318, 'p', '3', '/manager/api/v1/job', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3316, 'p', '3', '/manager/api/v1/job', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3317, 'p', '3', '/manager/api/v1/job', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3315, 'p', '3', '/manager/api/v1/jobs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3307, 'p', '3', '/manager/api/v1/menu', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3305, 'p', '3', '/manager/api/v1/menu', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3306, 'p', '3', '/manager/api/v1/menu', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3304, 'p', '3', '/manager/api/v1/menus', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3367, 'p', '3', '/manager/api/v1/resource/cfg_env', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3375, 'p', '3', '/manager/api/v1/resource/cfg_server', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3342, 'p', '3', '/manager/api/v1/resource/uc_app', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3323, 'p', '3', '/manager/api/v1/role', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3320, 'p', '3', '/manager/api/v1/role', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3321, 'p', '3', '/manager/api/v1/role', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3325, 'p', '3', '/manager/api/v1/role/menu', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3324, 'p', '3', '/manager/api/v1/role/menu_ids', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3322, 'p', '3', '/manager/api/v1/role/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3319, 'p', '3', '/manager/api/v1/roles', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3329, 'p', '3', '/manager/api/v1/user', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3327, 'p', '3', '/manager/api/v1/user', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3328, 'p', '3', '/manager/api/v1/user', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3355, 'p', '3', '/manager/api/v1/user/login/logs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3331, 'p', '3', '/manager/api/v1/user/password/reset', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3330, 'p', '3', '/manager/api/v1/user/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3326, 'p', '3', '/manager/api/v1/users', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3426, 'p', '3', '/notify/api/v1/channel', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3424, 'p', '3', '/notify/api/v1/channel', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3425, 'p', '3', '/notify/api/v1/channel', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3423, 'p', '3', '/notify/api/v1/channels', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3440, 'p', '3', '/notify/api/v1/logs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3431, 'p', '3', '/notify/api/v1/notifies', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3435, 'p', '3', '/notify/api/v1/notify', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3432, 'p', '3', '/notify/api/v1/notify', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3433, 'p', '3', '/notify/api/v1/notify', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3434, 'p', '3', '/notify/api/v1/notify', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3427, 'p', '3', '/notify/api/v1/notify_categories', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3430, 'p', '3', '/notify/api/v1/notify_category', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3428, 'p', '3', '/notify/api/v1/notify_category', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3429, 'p', '3', '/notify/api/v1/notify_category', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3439, 'p', '3', '/notify/api/v1/template', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3437, 'p', '3', '/notify/api/v1/template', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3438, 'p', '3', '/notify/api/v1/template', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3436, 'p', '3', '/notify/api/v1/templates', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3356, 'p', '3', '/resource/api/v1/directories', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3359, 'p', '3', '/resource/api/v1/directory', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3357, 'p', '3', '/resource/api/v1/directory', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3358, 'p', '3', '/resource/api/v1/directory', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3365, 'p', '3', '/resource/api/v1/export', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3364, 'p', '3', '/resource/api/v1/export', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3363, 'p', '3', '/resource/api/v1/exports', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3362, 'p', '3', '/resource/api/v1/file', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3361, 'p', '3', '/resource/api/v1/file', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3360, 'p', '3', '/resource/api/v1/files', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3639, 'p', 'test', '/ai-agent/api/v1/assessment', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3637, 'p', 'test', '/ai-agent/api/v1/assessment', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3638, 'p', 'test', '/ai-agent/api/v1/assessment', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3636, 'p', 'test', '/ai-agent/api/v1/assessments', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3644, 'p', 'test', '/ai-agent/api/v1/chat', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3640, 'p', 'test', '/ai-agent/api/v1/chat', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3643, 'p', 'test', '/ai-agent/api/v1/chat', 'update', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3642, 'p', 'test', '/ai-agent/api/v1/chat/conversation', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3641, 'p', 'test', '/ai-agent/api/v1/chat/messages', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3635, 'p', 'test', '/ai-agent/api/v1/knowledge', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3633, 'p', 'test', '/ai-agent/api/v1/knowledge', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3634, 'p', 'test', '/ai-agent/api/v1/knowledge', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3632, 'p', 'test', '/ai-agent/api/v1/knowledges', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3627, 'p', 'test', '/ai-agent/api/v1/model', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3625, 'p', 'test', '/ai-agent/api/v1/model', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3626, 'p', 'test', '/ai-agent/api/v1/model', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3624, 'p', 'test', '/ai-agent/api/v1/models', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3623, 'p', 'test', '/ai-agent/api/v1/secret', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3621, 'p', 'test', '/ai-agent/api/v1/secret', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3622, 'p', 'test', '/ai-agent/api/v1/secret', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3619, 'p', 'test', '/ai-agent/api/v1/secret_group', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3617, 'p', 'test', '/ai-agent/api/v1/secret_group', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3618, 'p', 'test', '/ai-agent/api/v1/secret_group', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3616, 'p', 'test', '/ai-agent/api/v1/secret_groups', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3620, 'p', 'test', '/ai-agent/api/v1/secrets', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3631, 'p', 'test', '/ai-agent/api/v1/tool', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3629, 'p', 'test', '/ai-agent/api/v1/tool', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3630, 'p', 'test', '/ai-agent/api/v1/tool', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3628, 'p', 'test', '/ai-agent/api/v1/tools', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3558, 'p', 'test', '/configure/api/v1/business', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3556, 'p', 'test', '/configure/api/v1/business', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3557, 'p', 'test', '/configure/api/v1/business', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3555, 'p', 'test', '/configure/api/v1/businesses', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3575, 'p', 'test', '/configure/api/v1/configure', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3574, 'p', 'test', '/configure/api/v1/configure/compare', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3546, 'p', 'test', '/configure/api/v1/env', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3543, 'p', 'test', '/configure/api/v1/env', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3544, 'p', 'test', '/configure/api/v1/env', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3545, 'p', 'test', '/configure/api/v1/env/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3547, 'p', 'test', '/configure/api/v1/env/token', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3548, 'p', 'test', '/configure/api/v1/env/token', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3541, 'p', 'test', '/configure/api/v1/envs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3564, 'p', 'test', '/configure/api/v1/resource', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3562, 'p', 'test', '/configure/api/v1/resource', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3563, 'p', 'test', '/configure/api/v1/resource', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3561, 'p', 'test', '/configure/api/v1/resources', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3554, 'p', 'test', '/configure/api/v1/server', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3551, 'p', 'test', '/configure/api/v1/server', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3552, 'p', 'test', '/configure/api/v1/server', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3553, 'p', 'test', '/configure/api/v1/server/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3549, 'p', 'test', '/configure/api/v1/servers', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3568, 'p', 'test', '/configure/api/v1/template', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3570, 'p', 'test', '/configure/api/v1/template', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3571, 'p', 'test', '/configure/api/v1/template/compare', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3569, 'p', 'test', '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3573, 'p', 'test', '/configure/api/v1/template/preview', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3572, 'p', 'test', '/configure/api/v1/template/switch', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3567, 'p', 'test', '/configure/api/v1/templates', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3559, 'p', 'test', '/configure/business/values', 'get', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3560, 'p', 'test', '/configure/business/values', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3565, 'p', 'test', '/configure/resource/values', 'get', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3566, 'p', 'test', '/configure/resource/values', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3597, 'p', 'test', '/cron/api/v1/log', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3596, 'p', 'test', '/cron/api/v1/logs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3594, 'p', 'test', '/cron/api/v1/task', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3590, 'p', 'test', '/cron/api/v1/task', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3593, 'p', 'test', '/cron/api/v1/task', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3588, 'p', 'test', '/cron/api/v1/task_group', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3586, 'p', 'test', '/cron/api/v1/task_group', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3587, 'p', 'test', '/cron/api/v1/task_group', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3585, 'p', 'test', '/cron/api/v1/task_groups', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3592, 'p', 'test', '/cron/api/v1/task/cancel', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3591, 'p', 'test', '/cron/api/v1/task/exec', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3595, 'p', 'test', '/cron/api/v1/task/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3589, 'p', 'test', '/cron/api/v1/tasks', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3583, 'p', 'test', '/cron/api/v1/worker', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3581, 'p', 'test', '/cron/api/v1/worker', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3582, 'p', 'test', '/cron/api/v1/worker', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3579, 'p', 'test', '/cron/api/v1/worker_group', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3577, 'p', 'test', '/cron/api/v1/worker_group', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3578, 'p', 'test', '/cron/api/v1/worker_group', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3576, 'p', 'test', '/cron/api/v1/worker_groups', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3584, 'p', 'test', '/cron/api/v1/worker/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3580, 'p', 'test', '/cron/api/v1/workers', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3521, 'p', 'test', '/manager/api/v1/app', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3518, 'p', 'test', '/manager/api/v1/app', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3519, 'p', 'test', '/manager/api/v1/app', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3520, 'p', 'test', '/manager/api/v1/app/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3516, 'p', 'test', '/manager/api/v1/apps', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3510, 'p', 'test', '/manager/api/v1/channel', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3508, 'p', 'test', '/manager/api/v1/channel', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3509, 'p', 'test', '/manager/api/v1/channel', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3507, 'p', 'test', '/manager/api/v1/channels', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3486, 'p', 'test', '/manager/api/v1/department', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3484, 'p', 'test', '/manager/api/v1/department', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3485, 'p', 'test', '/manager/api/v1/department', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3489, 'p', 'test', '/manager/api/v1/department_classify', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3487, 'p', 'test', '/manager/api/v1/department_classify', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3488, 'p', 'test', '/manager/api/v1/department_classify', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3483, 'p', 'test', '/manager/api/v1/departments', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3470, 'p', 'test', '/manager/api/v1/dictionaries', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3473, 'p', 'test', '/manager/api/v1/dictionary', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3471, 'p', 'test', '/manager/api/v1/dictionary', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3472, 'p', 'test', '/manager/api/v1/dictionary', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3478, 'p', 'test', '/manager/api/v1/dictionary_value', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3475, 'p', 'test', '/manager/api/v1/dictionary_value', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3476, 'p', 'test', '/manager/api/v1/dictionary_value', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3477, 'p', 'test', '/manager/api/v1/dictionary_value/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3474, 'p', 'test', '/manager/api/v1/dictionary_values', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3529, 'p', 'test', '/manager/api/v1/feedback', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3527, 'p', 'test', '/manager/api/v1/feedback', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3528, 'p', 'test', '/manager/api/v1/feedback', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3522, 'p', 'test', '/manager/api/v1/feedback_categories', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3525, 'p', 'test', '/manager/api/v1/feedback_category', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3523, 'p', 'test', '/manager/api/v1/feedback_category', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3524, 'p', 'test', '/manager/api/v1/feedback_category', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3526, 'p', 'test', '/manager/api/v1/feedbacks', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3515, 'p', 'test', '/manager/api/v1/field', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3512, 'p', 'test', '/manager/api/v1/field', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3513, 'p', 'test', '/manager/api/v1/field', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3514, 'p', 'test', '/manager/api/v1/field/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3511, 'p', 'test', '/manager/api/v1/fields', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3493, 'p', 'test', '/manager/api/v1/job', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3491, 'p', 'test', '/manager/api/v1/job', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3492, 'p', 'test', '/manager/api/v1/job', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3490, 'p', 'test', '/manager/api/v1/jobs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3482, 'p', 'test', '/manager/api/v1/menu', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3480, 'p', 'test', '/manager/api/v1/menu', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3481, 'p', 'test', '/manager/api/v1/menu', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3479, 'p', 'test', '/manager/api/v1/menus', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3542, 'p', 'test', '/manager/api/v1/resource/cfg_env', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3550, 'p', 'test', '/manager/api/v1/resource/cfg_server', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3517, 'p', 'test', '/manager/api/v1/resource/uc_app', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3498, 'p', 'test', '/manager/api/v1/role', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3495, 'p', 'test', '/manager/api/v1/role', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3496, 'p', 'test', '/manager/api/v1/role', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3500, 'p', 'test', '/manager/api/v1/role/menu', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3499, 'p', 'test', '/manager/api/v1/role/menu_ids', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3497, 'p', 'test', '/manager/api/v1/role/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3494, 'p', 'test', '/manager/api/v1/roles', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3504, 'p', 'test', '/manager/api/v1/user', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3502, 'p', 'test', '/manager/api/v1/user', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3503, 'p', 'test', '/manager/api/v1/user', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3530, 'p', 'test', '/manager/api/v1/user/login/logs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3506, 'p', 'test', '/manager/api/v1/user/password/reset', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3505, 'p', 'test', '/manager/api/v1/user/status', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3501, 'p', 'test', '/manager/api/v1/users', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3601, 'p', 'test', '/notify/api/v1/channel', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3599, 'p', 'test', '/notify/api/v1/channel', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3600, 'p', 'test', '/notify/api/v1/channel', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3598, 'p', 'test', '/notify/api/v1/channels', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3615, 'p', 'test', '/notify/api/v1/logs', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3606, 'p', 'test', '/notify/api/v1/notifies', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3610, 'p', 'test', '/notify/api/v1/notify', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3607, 'p', 'test', '/notify/api/v1/notify', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3608, 'p', 'test', '/notify/api/v1/notify', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3609, 'p', 'test', '/notify/api/v1/notify', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3602, 'p', 'test', '/notify/api/v1/notify_categories', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3605, 'p', 'test', '/notify/api/v1/notify_category', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3603, 'p', 'test', '/notify/api/v1/notify_category', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3604, 'p', 'test', '/notify/api/v1/notify_category', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3614, 'p', 'test', '/notify/api/v1/template', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3612, 'p', 'test', '/notify/api/v1/template', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3613, 'p', 'test', '/notify/api/v1/template', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3611, 'p', 'test', '/notify/api/v1/templates', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3531, 'p', 'test', '/resource/api/v1/directories', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3534, 'p', 'test', '/resource/api/v1/directory', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3532, 'p', 'test', '/resource/api/v1/directory', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3533, 'p', 'test', '/resource/api/v1/directory', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3540, 'p', 'test', '/resource/api/v1/export', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3539, 'p', 'test', '/resource/api/v1/export', 'POST', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3538, 'p', 'test', '/resource/api/v1/exports', 'GET', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3537, 'p', 'test', '/resource/api/v1/file', 'DELETE', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3536, 'p', 'test', '/resource/api/v1/file', 'PUT', NULL, NULL, NULL);
INSERT INTO "public"."casbin_rule_20250811" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES (3535, 'p', 'test', '/resource/api/v1/files', 'GET', NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for channel_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."channel_20250811";
CREATE TABLE "public"."channel_20250811" (
  "id" int8,
  "logo" varchar(254) COLLATE "pg_catalog"."default",
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "ak" varchar(254) COLLATE "pg_catalog"."default",
  "sk" varchar(254) COLLATE "pg_catalog"."default",
  "extra" text COLLATE "pg_catalog"."default",
  "status" int4,
  "created_at" int8,
  "updated_at" int8,
  "admin" int4,
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "type" varchar(254) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."channel_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of channel_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."channel_20250811" ("id", "logo", "keyword", "name", "ak", "sk", "extra", "status", "created_at", "updated_at", "admin", "description", "type") VALUES (1, '2229a835e0457a716906046cf5c66a96', 'yiban1', '易班', '45eff8ae0d15a3d2', '02cf8747dfb2b724a089845776a6bcd7', '{
    "callback": "http://f.yiban.cn/iapp569942"
}', 1, 1749374196, 1751118367, 1, '易班后台登陆', 'yiban');
INSERT INTO "public"."channel_20250811" ("id", "logo", "keyword", "name", "ak", "sk", "extra", "status", "created_at", "updated_at", "admin", "description", "type") VALUES (2, '2252554cf6309d2e53e95a5d40458d17', 'wx', '微信公众号', 'wxcc26972e8a4cdd79', 'd12e0c80c3ca99197efaf2132bd6f4c4', '{
    "callback": "http://192.168.1.14:5173"
}', 1, 1749996349, 1750514891, 1, '后台登陆测试', 'wx_official_account');
INSERT INTO "public"."channel_20250811" ("id", "logo", "keyword", "name", "ak", "sk", "extra", "status", "created_at", "updated_at", "admin", "description", "type") VALUES (3, '385d37202ae8f08cd8ba429eb51b5422', 'email1', '邮箱', '860808187@qq.com', 'fyudafdzqmhwbfbd', '{
    "host": "smtp.qq.com",
    "port": 25
}', 1, 1750513180, 1750514898, 1, 'qq邮箱', 'email');
COMMIT;

-- ----------------------------
-- Table structure for department_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."department_20250811";
CREATE TABLE "public"."department_20250811" (
  "id" int8,
  "classify_id" int8,
  "parent_id" int8,
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."department_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of department_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."department_20250811" ("id", "classify_id", "parent_id", "keyword", "name", "description", "created_at", "updated_at") VALUES (1, 1, 0, 'company', '贵州青橙科技有限公司', '开放合作，拥抱未来', 1713706137, 1713706137);
INSERT INTO "public"."department_20250811" ("id", "classify_id", "parent_id", "keyword", "name", "description", "created_at", "updated_at") VALUES (5, 2, 1, 'dep_1', '下级测试部门', '下级测试部门', 1720685640, 1751567775);
INSERT INTO "public"."department_20250811" ("id", "classify_id", "parent_id", "keyword", "name", "description", "created_at", "updated_at") VALUES (6, 1, 5, 'dep_2', '下级测试部门2', '下级测试部门', 1720685653, 1720685653);
INSERT INTO "public"."department_20250811" ("id", "classify_id", "parent_id", "keyword", "name", "description", "created_at", "updated_at") VALUES (7, 1, 1, 'dep3', '下级测试部门3', '下级测试部门2', 1720685663, 1751566524);
INSERT INTO "public"."department_20250811" ("id", "classify_id", "parent_id", "keyword", "name", "description", "created_at", "updated_at") VALUES (8, 1, 7, 'dep4', '下级测试部门4', '下级测试部门2', 1720685670, 1720685670);
INSERT INTO "public"."department_20250811" ("id", "classify_id", "parent_id", "keyword", "name", "description", "created_at", "updated_at") VALUES (9, 1, 6, 'dep5', '下级测试部门5', '下级测试部门2', 1720685679, 1720685679);
INSERT INTO "public"."department_20250811" ("id", "classify_id", "parent_id", "keyword", "name", "description", "created_at", "updated_at") VALUES (13, 1, 5, '测试', '测试', '测试', 1749315479, 1749315479);
COMMIT;

-- ----------------------------
-- Table structure for department_classify_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."department_classify_20250811";
CREATE TABLE "public"."department_classify_20250811" (
  "id" int8,
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."department_classify_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of department_classify_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."department_classify_20250811" ("id", "name", "description", "created_at", "updated_at") VALUES (1, '行政部门', '行政部门', 1749314470, 1749314470);
INSERT INTO "public"."department_classify_20250811" ("id", "name", "description", "created_at", "updated_at") VALUES (2, '学院', '校级单位', 1749358194, 1749358194);
COMMIT;

-- ----------------------------
-- Table structure for department_closure_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."department_closure_20250811";
CREATE TABLE "public"."department_closure_20250811" (
  "id" int8,
  "parent" int8,
  "children" int8
)
;
ALTER TABLE "public"."department_closure_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of department_closure_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."department_closure_20250811" ("id", "parent", "children") VALUES (22, 5, 13);
COMMIT;

-- ----------------------------
-- Table structure for department_role_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."department_role_20250811";
CREATE TABLE "public"."department_role_20250811" (
  "id" int8,
  "role_id" int8,
  "department_" int8
)
;
ALTER TABLE "public"."department_role_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of department_role_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."department_role_20250811" ("id", "role_id", "department_") VALUES (1, 1, 1);
INSERT INTO "public"."department_role_20250811" ("id", "role_id", "department_") VALUES (2, 5, 7);
INSERT INTO "public"."department_role_20250811" ("id", "role_id", "department_") VALUES (4, 9, 5);
COMMIT;

-- ----------------------------
-- Table structure for dictionary_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."dictionary_20250811";
CREATE TABLE "public"."dictionary_20250811" (
  "id" int8,
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8,
  "deleted_at" int8,
  "type" varchar(254) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."dictionary_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of dictionary_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."dictionary_20250811" ("id", "keyword", "name", "description", "created_at", "updated_at", "deleted_at", "type") VALUES (4, 't2', 't2', 't', 1721835689, 1721837018, 0, 'list');
INSERT INTO "public"."dictionary_20250811" ("id", "keyword", "name", "description", "created_at", "updated_at", "deleted_at", "type") VALUES (6, 'test', '1', '1', 1721964102, 1728227177, 0, 'tree');
COMMIT;

-- ----------------------------
-- Table structure for dictionary_value_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."dictionary_value_20250811";
CREATE TABLE "public"."dictionary_value_20250811" (
  "id" int8,
  "dictionary_" int8,
  "label" varchar(254) COLLATE "pg_catalog"."default",
  "value" varchar(254) COLLATE "pg_catalog"."default",
  "status" int4,
  "weight" int8,
  "type" varchar(254) COLLATE "pg_catalog"."default",
  "extra" varchar(254) COLLATE "pg_catalog"."default",
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8,
  "parent_id" int8
)
;
ALTER TABLE "public"."dictionary_value_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of dictionary_value_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."dictionary_value_20250811" ("id", "dictionary_", "label", "value", "status", "weight", "type", "extra", "description", "created_at", "updated_at", "parent_id") VALUES (3, 6, '在', 'keyword', 1, 0, NULL, NULL, NULL, 1728118961, 1728121524, 0);
INSERT INTO "public"."dictionary_value_20250811" ("id", "dictionary_", "label", "value", "status", "weight", "type", "extra", "description", "created_at", "updated_at", "parent_id") VALUES (4, 6, '测试节点', '1', 1, 0, NULL, NULL, NULL, 1728119659, 1728121526, 3);
INSERT INTO "public"."dictionary_value_20250811" ("id", "dictionary_", "label", "value", "status", "weight", "type", "extra", "description", "created_at", "updated_at", "parent_id") VALUES (5, 4, '1', '1', NULL, 0, '1', '1', NULL, 1728221359, 1728221359, 0);
COMMIT;

-- ----------------------------
-- Table structure for feedback_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."feedback_20250811";
CREATE TABLE "public"."feedback_20250811" (
  "id" int8,
  "app_id" int8,
  "user_id" int8,
  "category_id" int8,
  "title" varchar(254) COLLATE "pg_catalog"."default",
  "content" text COLLATE "pg_catalog"."default",
  "status" varchar(254) COLLATE "pg_catalog"."default",
  "images" text COLLATE "pg_catalog"."default",
  "contact" varchar(254) COLLATE "pg_catalog"."default",
  "device" text COLLATE "pg_catalog"."default",
  "platform" varchar(254) COLLATE "pg_catalog"."default",
  "version" varchar(254) COLLATE "pg_catalog"."default",
  "md5" varchar(254) COLLATE "pg_catalog"."default",
  "processed_b" int8,
  "processed_r" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."feedback_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of feedback_20250811
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for feedback_category_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."feedback_category_20250811";
CREATE TABLE "public"."feedback_category_20250811" (
  "id" int8,
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8
)
;
ALTER TABLE "public"."feedback_category_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of feedback_category_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."feedback_category_20250811" ("id", "name", "created_at") VALUES (1, '1', 1750093897);
COMMIT;

-- ----------------------------
-- Table structure for field_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."field_20250811";
CREATE TABLE "public"."field_20250811" (
  "id" int8,
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "type" varchar(254) COLLATE "pg_catalog"."default",
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "status" int4,
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."field_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of field_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."field_20250811" ("id", "keyword", "type", "name", "status", "description", "created_at", "updated_at") VALUES (1, 'key', 'string', '测试', 1, 'cc', 1749364667, 1749364791);
COMMIT;

-- ----------------------------
-- Table structure for gorm_init_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."gorm_init_20250811";
CREATE TABLE "public"."gorm_init_20250811" (
  "id" int8,
  "init" int4
)
;
ALTER TABLE "public"."gorm_init_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of gorm_init_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."gorm_init_20250811" ("id", "init") VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for job_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."job_20250811";
CREATE TABLE "public"."job_20250811" (
  "id" int8,
  "parent_id" int8,
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "weight" int8,
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8,
  "deleted_at" int8
)
;
ALTER TABLE "public"."job_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of job_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."job_20250811" ("id", "parent_id", "keyword", "name", "weight", "description", "created_at", "updated_at", "deleted_at") VALUES (1, 0, 'chairman', '董事长', 2, '董事长', 1713706137, 1721838228, 0);
INSERT INTO "public"."job_20250811" ("id", "parent_id", "keyword", "name", "weight", "description", "created_at", "updated_at", "deleted_at") VALUES (20, 1, '2', '1', NULL, '3', 1751133357, 1751133357, 0);
COMMIT;

-- ----------------------------
-- Table structure for job_closure_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."job_closure_20250811";
CREATE TABLE "public"."job_closure_20250811" (
  "id" int8,
  "parent" int8,
  "children" int8
)
;
ALTER TABLE "public"."job_closure_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of job_closure_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."job_closure_20250811" ("id", "parent", "children") VALUES (1, 1, 20);
COMMIT;

-- ----------------------------
-- Table structure for login_log_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."login_log_20250811";
CREATE TABLE "public"."login_log_20250811" (
  "id" int8,
  "username" varchar(254) COLLATE "pg_catalog"."default",
  "type" varchar(254) COLLATE "pg_catalog"."default",
  "ip" varchar(254) COLLATE "pg_catalog"."default",
  "address" varchar(254) COLLATE "pg_catalog"."default",
  "device" varchar(254) COLLATE "pg_catalog"."default",
  "browser" varchar(254) COLLATE "pg_catalog"."default",
  "code" int4,
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8
)
;
ALTER TABLE "public"."login_log_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of login_log_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (1, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1725125197);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (2, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1725163517);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (3, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1725247547);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (4, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1725331878);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (5, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1725553142);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (6, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1725592696);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (7, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1725803463);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (8, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726283101);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (9, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726310473);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (10, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726587996);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (11, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726829870);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (12, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726850129);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (13, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726896761);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (14, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726904398);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (15, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1726923972);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (16, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727010081);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (17, 'P164364', NULL, '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 500, 'error: code = 500 reason = UsernameFormatError message = 用户名格式错误 metadata = map[] cause = <nil>', 1727331524);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (18, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 500, 'error: code = 500 reason = PasswordError message = 密码错误 metadata = map[] cause = <nil>', 1727331545);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (19, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727331556);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (20, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727585872);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (21, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727593093);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (22, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727604056);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (23, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727605268);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (24, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727614418);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (25, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727623596);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (26, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727667100);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (27, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727693387);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (28, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727703232);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (29, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727711238);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (30, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727924106);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (31, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1727942333);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (32, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728053602);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (33, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728071085);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (34, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728105814);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (35, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728115267);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (36, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728221279);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (37, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728562111);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (38, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1728615199);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (39, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729134980);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (40, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729253151);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (41, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729411166);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (42, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 112.0.0.0', 200, '登陆成功', 1729413157);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (43, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 130.0.0.0', 200, '登陆成功', 1731225514);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (44, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 130.0.0.0', 200, '登陆成功', 1731255244);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (45, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 130.0.0.0', 200, '登陆成功', 1732376970);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (46, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1735738516);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (47, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1735808462);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (48, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1735890213);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (49, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1735902741);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (50, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1735982765);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (51, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736009711);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (52, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736057413);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (53, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736069049);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (54, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736237596);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (55, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736249171);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (56, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736408698);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (57, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736426230);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (58, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736477883);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (59, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736493839);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (60, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736737519);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (61, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736744836);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (62, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736757523);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (63, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736823954);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (64, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1736855850);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (65, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1737130853);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (66, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1737184097);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (67, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1737193492);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (68, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1737300782);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (69, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1739640909);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (70, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1739707386);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (71, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1739897240);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (72, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1739956596);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (73, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740120514);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (74, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740129024);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (75, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740136876);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (76, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740228186);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (77, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740245515);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (78, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740295655);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (79, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740306028);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (80, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740325586);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (81, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740333772);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (82, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740379491);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (83, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740387905);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (84, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740407996);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (85, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740451263);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (86, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740467178);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (87, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740497704);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (88, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740506231);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (89, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740584387);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (90, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740592796);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (91, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740638610);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (92, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740674753);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (93, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740710198);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (94, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740718231);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (95, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1740756072);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (96, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741022575);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (97, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741089406);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (98, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741107015);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (99, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741143628);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (100, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741186764);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (101, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741431231);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (102, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741449856);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (103, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741459635);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (104, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741465403);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (105, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741507741);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (106, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741520506);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (107, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741526101);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (108, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741717647);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (109, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741875922);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (110, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741885005);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (111, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'iOS 13.2.3', 'Safari 13.0.3', 200, '登陆成功', 1741922825);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (112, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1741940715);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (113, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742012057);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (114, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742027470);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (115, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742038997);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (116, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742039186);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (117, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742040716);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (118, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742040811);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (119, '18286219255', 'phone', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742042110);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (120, '18286219255', 'phone', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742043219);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (121, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742057388);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (122, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742099796);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (123, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742115939);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (124, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 131.0.0.0', 200, '登陆成功', 1742143844);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (125, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742223094);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (126, '18286219255', 'phone', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742223234);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (127, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742491483);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (128, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742523305);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (129, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742532558);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (130, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742543549);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (131, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742617249);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (132, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742642812);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (133, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742696617);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (134, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742914642);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (135, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742915227);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (136, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742915664);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (137, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742915756);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (138, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742951695);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (139, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742973322);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (140, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1742987178);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (141, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743001399);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (142, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743030270);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (143, 'emotion@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1743057458);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (144, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743057468);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (145, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743087231);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (146, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Safari 17.6', 200, '登陆成功', 1743093873);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (147, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743137725);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (148, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743157987);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (149, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743263010);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (150, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743321891);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (151, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743332564);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (152, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743350752);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (153, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Safari 17.6', 200, '登陆成功', 1743350773);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (154, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743392460);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (155, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743479251);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (156, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743487388);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (157, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743504455);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (158, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743592701);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (159, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743648241);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (160, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743650290);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (161, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743650575);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (162, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743673472);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (163, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743694172);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (217, '超级管理员', 'email', '127.0.0.1', '本地', 'iOS 13.2.3', 'Safari 13.0.3', 200, '登陆成功', 1751099661);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (164, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743736004);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (165, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743744620);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (166, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743855539);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (167, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1743920508);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (168, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1744008125);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (169, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1744178562);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (170, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1744197655);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (171, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1744906207);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (172, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1744951358);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (173, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1744960574);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (174, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1744999557);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (175, '18888888888', 'phone', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 134.0.0.0', 200, '登陆成功', 1745569360);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (176, 'emotion@baidu.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 135.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1747321525);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (177, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 135.0.0.0', 200, '登陆成功', 1747321539);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (178, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 135.0.0.0', 200, '登陆成功', 1747332425);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (179, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 135.0.0.0', 200, '登陆成功', 1747365341);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (180, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 135.0.0.0', 200, '登陆成功', 1747640750);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (181, 'emotion@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 135.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1747640792);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (182, 'emotion@baidu.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 135.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1747900976);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (183, 'emotion@baidu.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1749286436);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (184, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749286454);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (185, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749308304);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (186, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749357950);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (187, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749361018);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (188, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749372142);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (189, '1280291001@qq.com', 'email', '127.0.0.1', '本地登陆', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749392093);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (190, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749879812);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (191, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749900122);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (192, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749913572);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (193, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749982132);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (194, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749982149);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (195, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749982214);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (196, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749982249);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (197, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749982647);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (198, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749982768);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (199, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749982813);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (200, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749985528);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (201, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749992505);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (202, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749992588);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (203, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1749994770);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (204, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750087832);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (205, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750089796);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (206, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750260231);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (207, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750512953);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (208, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750519842);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (209, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750594110);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (210, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750607494);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (211, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750615427);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (212, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750698684);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (213, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750761257);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (214, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1750769290);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (215, NULL, 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1751047578);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (216, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751047630);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (218, '超级管理员', 'email', '127.0.0.1', '本地', 'iOS 13.2.3', 'Safari 13.0.3', 200, '登陆成功', 1751099903);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (219, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751106394);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (220, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751111223);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (221, '超级管理员', 'email', '127.0.0.1', '本地', 'iOS 13.2.3', 'Safari 13.0.3', 200, '登陆成功', 1751115190);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (222, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751118136);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (223, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751125477);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (224, '1', 'phone', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751127437);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (225, '1', 'phone', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751127462);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (226, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751133345);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (227, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751179441);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (228, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1751195450);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (229, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751196068);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (230, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751196084);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (231, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751196140);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (232, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751217223);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (233, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751306022);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (234, '超级管理员', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751561768);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (235, '31@q.com', NULL, '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = UsernameFormatError message = 用户名格式错误 metadata = map[] cause = <nil>', 1751563488);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (236, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = RoleDisableError message = 角色已被禁用 metadata = map[] cause = <nil>', 1751563528);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (237, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = SystemError message = 系统错误 metadata = map[] cause = <nil>', 1751564168);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (238, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751564248);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (239, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751564481);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (240, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751564527);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (241, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751564663);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (242, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751564894);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (243, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751565134);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (244, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751565194);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (245, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751565434);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (246, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751565727);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (247, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = PasswordError message = 密码错误 metadata = map[] cause = <nil>', 1751566372);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (248, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = PasswordError message = 密码错误 metadata = map[] cause = <nil>', 1751566391);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (249, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751566538);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (250, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751566616);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (251, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751566637);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (252, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751566713);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (253, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751566731);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (254, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751566842);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (255, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751567847);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (256, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751636885);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (257, '31@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1751638389);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (258, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1752051210);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (259, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1752075677);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (260, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1752132631);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (261, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1752148638);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (262, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1752217316);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (263, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1752473816);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (264, 'emotion@baidu.com', 'email', '127.0.0.1', '本地', 'iOS 16.6', 'Safari 16.6', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1752473843);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (265, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1752473869);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (266, 'admin@baidu.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1753112529);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (267, '1280291001@baidu.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 500, 'error: code = 500 reason = UsernameNotExistError message = 用户不存在 metadata = map[] cause = <nil>', 1753112542);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (268, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1753112553);
INSERT INTO "public"."login_log_20250811" ("id", "username", "type", "ip", "address", "device", "browser", "code", "description", "created_at") VALUES (269, '1280291001@qq.com', 'email', '127.0.0.1', '本地', 'macOS 10.15.7', 'Chrome 137.0.0.0', 200, '登陆成功', 1753285756);
COMMIT;

-- ----------------------------
-- Table structure for menu_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."menu_20250811";
CREATE TABLE "public"."menu_20250811" (
  "id" int8,
  "parent_id" int8,
  "title" varchar(254) COLLATE "pg_catalog"."default",
  "type" varchar(254) COLLATE "pg_catalog"."default",
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "icon" varchar(254) COLLATE "pg_catalog"."default",
  "api" varchar(254) COLLATE "pg_catalog"."default",
  "method" varchar(254) COLLATE "pg_catalog"."default",
  "path" varchar(254) COLLATE "pg_catalog"."default",
  "permission" varchar(254) COLLATE "pg_catalog"."default",
  "component" varchar(254) COLLATE "pg_catalog"."default",
  "redirect" varchar(254) COLLATE "pg_catalog"."default",
  "weight" int8,
  "is_hidden" int4,
  "is_cache" int4,
  "is_home" int4,
  "is_affix" int4,
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."menu_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of menu_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (1, 0, '管理平台', 'R', 'SystemPlatform', 'desktop', NULL, NULL, '/', NULL, 'Layout', NULL, 2, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (2, 1, '首页面板', 'M', 'Dashboard', 'dashboard', NULL, NULL, '/dashboard', NULL, '/dashboard/index', NULL, 0, NULL, 1, 1, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (3, 1, '管理中心', 'M', 'SystemPlatformManager', 'desktop', NULL, NULL, '/manager', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (4, 3, '基础接口', 'G', 'ManagerBaseApi', 'apps', NULL, NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (5, 4, '获取用户可见部门树', 'BA', NULL, NULL, '/manager/api/v1/current/departments', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (6, 4, '获取用户可见部门分类', 'BA', NULL, NULL, '/manager/api/v1/department_classifies', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (7, 4, '获取用户可见角色树', 'BA', NULL, NULL, '/manager/api/v1/current/roles', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (8, 4, '获取个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (9, 4, '更新个人用户信息', 'BA', NULL, NULL, '/manager/api/v1/user/current/info', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (10, 4, '更新个人用户密码', 'BA', NULL, NULL, '/manager/api/v1/user/current/password', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (11, 4, '保存个人用户设置', 'BA', NULL, NULL, '/manager/api/v1/user/current/setting', 'PUT', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (12, 4, '发送用户验证吗', 'BA', NULL, NULL, '/manager/api/v1/user/current/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (13, 4, '获取用户当前角色菜单', 'BA', NULL, NULL, '/manager/api/v1/menus/by/cur_role', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (14, 4, '退出登录', 'BA', NULL, NULL, '/manager/api/v1/user/logout', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (15, 4, '刷新token', 'BA', NULL, NULL, '/manager/api/v1/user/token/refresh', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (16, 4, '用户登录', 'BA', NULL, NULL, '/manager/api/v1/user/login', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (17, 4, '获取登录验证码', 'BA', NULL, NULL, '/manager/api/v1/user/login/captcha', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (18, 4, '获取系统基础设置', 'BA', NULL, NULL, '/manager/api/v1/system/setting', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (19, 4, '接口鉴权', 'BA', NULL, NULL, '/manager/api/v1/auth', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (20, 4, '切换用户角色', 'BA', NULL, NULL, '/manager/api/v1/user/current/role', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (21, 4, '分块上传文件', 'BA', NULL, NULL, '/resource/api/v1/upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (22, 4, '预上传文件', 'BA', NULL, NULL, '/resource/api/v1/prepare_upload', 'POST', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (23, 4, '获取字段类型', 'BA', NULL, NULL, '/application/api/v1/field/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (24, 4, '查询资源权限', 'BA', NULL, NULL, '/manager/api/v1/resource', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (25, 4, '获取渠道类型', 'BA', NULL, NULL, '/application/api/v1/channel/types', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (26, 3, '字典管理', 'M', 'ManagerDictionary', 'storage', NULL, NULL, '/manager/dictionary', NULL, '/manager/dictionary/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (27, 26, '查询字典', 'A', NULL, NULL, '/manager/api/v1/dictionaries', 'GET', NULL, 'manager:dictionary:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (28, 26, '新增字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'POST', NULL, 'manager:dictionary:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (29, 26, '修改字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'PUT', NULL, 'manager:dictionary:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (30, 26, '删除字典', 'A', NULL, NULL, '/manager/api/v1/dictionary', 'DELETE', NULL, 'manager:dictionary:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (31, 26, '获取字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_values', 'GET', NULL, 'manager:dictionary:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (32, 26, '新增字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'POST', NULL, 'manager:dictionary:value:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (33, 26, '修改字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'PUT', NULL, 'manager:dictionary:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (34, 26, '更新字典值目录状态', 'A', NULL, NULL, '/manager/api/v1/dictionary_value/status', 'PUT', NULL, 'manager:dictionary:value:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (35, 26, '删除字典值', 'A', NULL, NULL, '/manager/api/v1/dictionary_value', 'DELETE', NULL, 'manager:dictionary:value:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (36, 3, '菜单管理', 'M', 'ManagerMenu', 'menu', NULL, NULL, '/manager/menu', NULL, '/manager/menu/index', NULL, 0, NULL, 1, 1, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (37, 36, '查询菜单', 'A', NULL, NULL, '/manager/api/v1/menus', 'GET', NULL, 'manager:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (38, 36, '新增菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'POST', NULL, 'manager:menu:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (39, 36, '修改菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'PUT', NULL, 'manager:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (40, 36, '删除菜单', 'A', NULL, NULL, '/manager/api/v1/menu', 'DELETE', NULL, 'manager:menu:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (41, 3, '部门管理', 'M', 'ManagerDepartment', 'user-group', NULL, NULL, '/manager/department', NULL, '/manager/department/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (42, 41, '查询部门', 'A', NULL, NULL, '/manager/api/v1/departments', 'GET', NULL, 'manager:department:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (43, 41, '新增部门', 'A', NULL, NULL, '/manager/api/v1/department', 'POST', NULL, 'manager:department:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (44, 41, '修改部门', 'A', NULL, NULL, '/manager/api/v1/department', 'PUT', NULL, 'manager:department:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (45, 41, '删除部门', 'A', NULL, NULL, '/manager/api/v1/department', 'DELETE', NULL, 'manager:department:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (46, 41, '新增部门分类', 'A', NULL, NULL, '/manager/api/v1/department_classify', 'POST', NULL, 'manager:department:classify:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (47, 41, '修改部门分类', 'A', NULL, NULL, '/manager/api/v1/department_classify', 'PUT', NULL, 'manager:department:classify:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (48, 41, '删除部门分类', 'A', NULL, NULL, '/manager/api/v1/department_classify', 'DELETE', NULL, 'manager:department:classify:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (49, 3, '职位管理', 'M', 'ManagerJob', 'tag', NULL, NULL, '/manager/job', NULL, '/manager/job/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (50, 49, '查询职位', 'A', NULL, NULL, '/manager/api/v1/jobs', 'GET', NULL, 'manager:job:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (51, 49, '新增职位', 'A', NULL, NULL, '/manager/api/v1/job', 'POST', NULL, 'manager:job:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (52, 49, '修改职位', 'A', NULL, NULL, '/manager/api/v1/job', 'PUT', NULL, 'manager:job:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (53, 49, '删除职位', 'A', NULL, NULL, '/manager/api/v1/job', 'DELETE', NULL, 'manager:job:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (54, 3, '角色管理', 'M', 'ManagerRole', 'safe', NULL, NULL, '/manager/role', NULL, '/manager/role/index', NULL, 0, NULL, NULL, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (55, 54, '查询角色', 'A', NULL, NULL, '/manager/api/v1/roles', 'GET', NULL, 'manager:role:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (56, 54, '新增角色', 'A', NULL, NULL, '/manager/api/v1/role', 'POST', NULL, 'manager:role:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (57, 54, '修改角色', 'A', NULL, NULL, '/manager/api/v1/role', 'PUT', NULL, 'manager:role:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (58, 54, '修改角色状态', 'A', NULL, NULL, '/manager/api/v1/role/status', 'PUT', NULL, 'manager:role:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (59, 54, '删除角色', 'A', NULL, NULL, '/manager/api/v1/role', 'DELETE', NULL, 'manager:role:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (60, 54, '角色菜单管理', 'G', NULL, NULL, NULL, NULL, NULL, 'manager:role:menu', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (61, 60, '查询角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu_ids', 'GET', NULL, 'manager:role:menu:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (62, 60, '修改角色菜单', 'A', NULL, NULL, '/manager/api/v1/role/menu', 'POST', NULL, 'manager:role:menu:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (63, 3, '用户管理', 'M', 'ManagerUser', 'user', NULL, NULL, '/manager/user', NULL, '/manager/user/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (64, 63, '查询用户列表', 'A', NULL, NULL, '/manager/api/v1/users', 'GET', NULL, 'manager:user:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (65, 63, '新增用户', 'A', NULL, NULL, '/manager/api/v1/user', 'POST', NULL, 'manager:user:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (66, 63, '修改用户', 'A', NULL, NULL, '/manager/api/v1/user', 'PUT', NULL, 'manager:user:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (67, 63, '删除用户', 'A', NULL, NULL, '/manager/api/v1/user', 'DELETE', NULL, 'manager:user:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (68, 63, '修改用户状态', 'A', NULL, NULL, '/manager/api/v1/user/status', 'PUT', NULL, 'manager:user:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (69, 63, '重置账号密码', 'A', NULL, NULL, '/manager/api/v1/user/password/reset', 'POST', NULL, 'manager:user:reset:password', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (70, 3, '授权渠道', 'M', 'managerChannel', 'mind-mapping', NULL, NULL, '/manager/channel', NULL, '/manager/channel/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (71, 70, '查看渠道', 'A', NULL, NULL, '/manager/api/v1/channels', 'GET', NULL, 'manager:channel:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (72, 70, '新增渠道', 'A', NULL, NULL, '/manager/api/v1/channel', 'POST', NULL, 'manager:channel:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (73, 70, '修改渠道', 'A', NULL, NULL, '/manager/api/v1/channel', 'PUT', NULL, 'manager:channel:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (74, 70, '删除渠道', 'A', NULL, NULL, '/manager/api/v1/channel', 'DELETE', NULL, 'manager:channel:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (75, 3, '信息字段', 'M', 'managerField', 'list', NULL, NULL, '/manager/field', NULL, '/manager/field/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (76, 75, '查看字段列表', 'A', NULL, NULL, '/manager/api/v1/fields', 'GET', NULL, 'manager:field:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (77, 75, '新增字段', 'A', NULL, NULL, '/manager/api/v1/field', 'POST', NULL, 'manager:field:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (78, 75, '修改字段', 'A', NULL, NULL, '/manager/api/v1/field', 'PUT', NULL, 'manager:field:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (79, 75, '修改字段状态', 'A', NULL, NULL, '/manager/api/v1/field/status', 'PUT', NULL, 'manager:field:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (80, 75, '删除字段', 'A', NULL, NULL, '/manager/api/v1/field', 'DELETE', NULL, 'manager:field:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (81, 3, '应用管理', 'M', 'managerApp', 'apps', NULL, NULL, '/manager/app', NULL, '/manager/app/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (82, 81, '查看应用', 'A', NULL, NULL, '/manager/api/v1/apps', 'GET', NULL, 'manager:app:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (83, 81, '设置应用资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/uc_app', 'PUT', NULL, 'manager:app:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (84, 81, '新增应用', 'A', NULL, NULL, '/manager/api/v1/app', 'POST', NULL, 'manager:app:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (85, 81, '修改应用', 'A', NULL, NULL, '/manager/api/v1/app', 'PUT', NULL, 'manager:app:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (86, 81, '修改应用状态', 'A', NULL, NULL, '/manager/api/v1/app/status', 'PUT', NULL, 'manager:app:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (87, 81, '删除应用', 'A', NULL, NULL, '/manager/api/v1/app', 'DELETE', NULL, 'manager:app:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (88, 3, '反馈管理', 'M', 'managerFeedback', 'question-circle', NULL, NULL, '/manager/feedback', NULL, '/manager/feedback/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (89, 88, '查看反馈分类', 'A', NULL, NULL, '/manager/api/v1/feedback_categories', 'GET', NULL, 'manager:feedback:category:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (90, 88, '新增反馈渠道', 'A', NULL, NULL, '/manager/api/v1/feedback_category', 'POST', NULL, 'manager:feedback:category:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (91, 88, '修改反馈渠道', 'A', NULL, NULL, '/manager/api/v1/feedback_category', 'PUT', NULL, 'manager:feedback:category:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (92, 88, '删除反馈渠道', 'A', NULL, NULL, '/manager/api/v1/feedback_category', 'DELETE', NULL, 'manager:feedback:category:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (93, 88, '查看反馈', 'A', NULL, NULL, '/manager/api/v1/feedbacks', 'GET', NULL, 'manager:feedback:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (94, 88, '新增反馈', 'A', NULL, NULL, '/manager/api/v1/feedback', 'POST', NULL, 'manager:feedback:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (95, 88, '修改反馈', 'A', NULL, NULL, '/manager/api/v1/feedback', 'PUT', NULL, 'manager:feedback:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (96, 88, '删除反馈', 'A', NULL, NULL, '/manager/api/v1/feedback', 'DELETE', NULL, 'manager:feedback:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (97, 3, '登陆日志', 'M', 'ManagerLoginLog', 'history', NULL, NULL, '/manager/log', NULL, '/manager/log/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (98, 97, '查询登陆列表', 'A', NULL, NULL, '/manager/api/v1/user/login/logs', 'GET', NULL, 'manager:login:log:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (99, 1, '资源中心', 'M', 'SystemPlatformResource', 'file', NULL, NULL, '/resource', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (100, 99, '文件管理', 'M', 'ResourceFile', 'file', NULL, NULL, '/resource/file', NULL, '/resource/file/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (101, 100, '目录管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:directory:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (102, 101, '查看目录', 'A', NULL, NULL, '/resource/api/v1/directories', 'GET', NULL, 'resource:directory:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (103, 101, '新增目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'POST', NULL, 'resource:directory:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (104, 101, '修改目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'PUT', NULL, 'resource:directory:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (105, 101, '删除目录', 'A', NULL, NULL, '/resource/api/v1/directory', 'DELETE', NULL, 'resource:directory:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (106, 100, '文件管理', 'G', NULL, NULL, NULL, NULL, NULL, 'resource:file:group', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (107, 106, '查看文件', 'A', NULL, NULL, '/resource/api/v1/files', 'GET', NULL, 'resource:file:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (108, 106, '修改文件', 'A', NULL, NULL, '/resource/api/v1/file', 'PUT', NULL, 'resource:file:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (109, 106, '删除文件', 'A', NULL, NULL, '/resource/api/v1/file', 'DELETE', NULL, 'resource:file:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (110, 99, '导出管理', 'M', 'ResourceExport', 'export', NULL, NULL, '/resource/export', NULL, '/resource/export/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (111, 110, '查看导出', 'A', NULL, NULL, '/resource/api/v1/exports', 'GET', NULL, 'resource:export:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (112, 110, '新增导出', 'A', NULL, NULL, '/resource/api/v1/export', 'POST', NULL, 'resource:export:file', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (113, 110, '删除导出', 'A', NULL, NULL, '/resource/api/v1/export', 'DELETE', NULL, 'resource:export:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (114, 1, '配置中心', 'M', 'SystemPlatformConfigure', 'code-block', NULL, NULL, '/configure', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (115, 114, '环境管理', 'M', 'ConfigureEnv', 'common', NULL, NULL, '/configure/env', NULL, '/configure/env/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (116, 115, '查看环境', 'A', NULL, NULL, '/configure/api/v1/envs', 'GET', NULL, 'configure:env:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (117, 115, '设置环境资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_env', 'PUT', NULL, 'configure:env:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (118, 115, '新增环境', 'A', NULL, NULL, '/configure/api/v1/env', 'POST', NULL, 'configure:env:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (119, 115, '修改环境', 'A', NULL, NULL, '/configure/api/v1/env', 'PUT', NULL, 'configure:env:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (120, 115, '修改环境状态', 'A', NULL, NULL, '/configure/api/v1/env/status', 'PUT', NULL, 'configure:env:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (121, 115, '删除环境', 'A', NULL, NULL, '/configure/api/v1/env', 'DELETE', NULL, 'configure:env:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (122, 115, '查看环境Token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'GET', NULL, 'configure:env:token:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (123, 115, '重置环境token', 'A', NULL, NULL, '/configure/api/v1/env/token', 'PUT', NULL, 'configure:env:token:reset', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (124, 114, '服务管理', 'M', 'ConfigureServer', 'apps', NULL, NULL, '/configure/server', NULL, '/configure/server/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (125, 124, '查询服务', 'A', NULL, NULL, '/configure/api/v1/servers', 'GET', NULL, 'configure:server:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (126, 124, '设置服务资源权限', 'A', NULL, NULL, '/manager/api/v1/resource/cfg_server', 'PUT', NULL, 'configure:server:resource:permission', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (127, 124, '新增服务', 'A', NULL, NULL, '/configure/api/v1/server', 'POST', NULL, 'configure:server:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (128, 124, '修改服务', 'A', NULL, NULL, '/configure/api/v1/server', 'PUT', NULL, 'configure:server:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (129, 124, '修改服务状态', 'A', NULL, NULL, '/configure/api/v1/server/status', 'PUT', NULL, 'configure:server:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (130, 124, '删除服务', 'A', NULL, NULL, '/configure/api/v1/server', 'DELETE', NULL, 'configure:server:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (131, 114, '业务变量', 'M', 'ConfigureBusiness', 'code', NULL, NULL, '/configure/business', NULL, '/configure/business/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (132, 131, '查看业务变量', 'A', NULL, NULL, '/configure/api/v1/businesses', 'GET', NULL, 'configure:business:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (133, 131, '新增业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'POST', NULL, 'configure:business:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (134, 131, '修改业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'PUT', NULL, 'configure:business:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (135, 131, '删除业务变量', 'A', NULL, NULL, '/configure/api/v1/business', 'DELETE', NULL, 'configure:business:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (136, 131, '查看业务变量值', 'A', NULL, NULL, '/configure/business/values', 'get', NULL, 'configure:business:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (137, 131, '设置业务变量值', 'A', NULL, NULL, '/configure/business/values', 'PUT', NULL, 'configure:business:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (138, 114, '资源变量', 'M', 'ConfigureResource', 'unordered-list', NULL, NULL, '/configure/resource', NULL, '/configure/resource/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (139, 138, '查看资源', 'A', NULL, NULL, '/configure/api/v1/resources', 'GET', NULL, 'configure:resource:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (140, 138, '新增资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'POST', NULL, 'configure:resource:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (141, 138, '修改资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'PUT', NULL, 'configure:resource:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (142, 138, '删除资源', 'A', NULL, NULL, '/configure/api/v1/resource', 'DELETE', NULL, 'configure:resource:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (143, 138, '查看业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'get', NULL, 'configure:resource:value:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (144, 138, '设置业务变量值', 'A', NULL, NULL, '/configure/resource/values', 'PUT', NULL, 'configure:resource:value:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (145, 114, '配置模板', 'M', 'ConfgureTemplate', 'file', NULL, NULL, '/configure/template', NULL, '/configure/template/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (146, 145, '模板管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (147, 146, '查看模板历史版本', 'A', NULL, NULL, '/configure/api/v1/templates', 'GET', NULL, 'configure:template:history', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (148, 146, '查看指定模板详细数据', 'A', NULL, NULL, '/configure/api/v1/template', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (149, 146, '查看当前正在使用的模板', 'A', NULL, NULL, '/configure/api/v1/template/current', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (150, 146, '提交模板', 'A', NULL, NULL, '/configure/api/v1/template', 'POST', NULL, 'configure:template:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (151, 146, '模板对比', 'A', NULL, NULL, '/configure/api/v1/template/compare', 'POST', NULL, 'configure:template:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (152, 146, '切换模板', 'A', NULL, NULL, '/configure/api/v1/template/switch', 'POST', NULL, 'configure:template:switch', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (153, 146, '模板预览', 'A', NULL, NULL, '/configure/api/v1/template/preview', 'POST', NULL, 'configure:template:preview', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (154, 145, '配置管理', 'G', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (155, 154, '配置对比', 'A', NULL, NULL, '/configure/api/v1/configure/compare', 'POST', NULL, 'configure:configure:compare', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (156, 154, '同步配置', 'A', NULL, NULL, '/configure/api/v1/configure', 'PUT', NULL, 'configure:configure:sync', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (157, 1, '定时任务', 'M', 'SystemPlatformCron', 'schedule', NULL, NULL, '/cron', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (158, 157, '节点管理', 'M', 'Worker', 'common', NULL, NULL, '/cron/worker', NULL, '/cron/worker/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (159, 158, '查看节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_groups', 'GET', NULL, 'cron:worker:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (160, 158, '新增节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'POST', NULL, 'cron:worker:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (161, 158, '修改节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'PUT', NULL, 'cron:worker:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (162, 158, '删除节点分组', 'A', NULL, NULL, '/cron/api/v1/worker_group', 'DELETE', NULL, 'cron:worker:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (163, 158, '查看节点', 'A', NULL, NULL, '/cron/api/v1/workers', 'GET', NULL, 'cron:worker:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (164, 158, '新增节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'POST', NULL, 'cron:worker:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (165, 158, '修改节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'PUT', NULL, 'cron:worker:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (166, 158, '删除节点', 'A', NULL, NULL, '/cron/api/v1/worker', 'DELETE', NULL, 'cron:worker:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (167, 158, '更新节点状态', 'A', NULL, NULL, '/cron/api/v1/worker/status', 'PUT', NULL, 'cron:worker:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (168, 157, '任务管理', 'M', 'Task', 'computer', NULL, NULL, '/cron/task', NULL, '/cron/task/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (169, 168, '查看任务分组', 'A', NULL, NULL, '/cron/api/v1/task_groups', 'GET', NULL, 'cron:task:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (170, 168, '新增任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'POST', NULL, 'cron:task:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (171, 168, '修改任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'PUT', NULL, 'cron:task:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (172, 168, '删除任务分组', 'A', NULL, NULL, '/cron/api/v1/task_group', 'DELETE', NULL, 'cron:task:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (173, 168, '查看任务', 'A', NULL, NULL, '/cron/api/v1/tasks', 'GET', NULL, 'cron:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (174, 168, '新增任务', 'A', NULL, NULL, '/cron/api/v1/task', 'POST', NULL, 'cron:task:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (175, 168, '立即执行任务', 'A', NULL, NULL, '/cron/api/v1/task/exec', 'POST', NULL, 'cron:task:exec', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (176, 168, '取消执行任务', 'A', NULL, NULL, '/cron/api/v1/task/cancel', 'POST', NULL, 'cron:task:cancel', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (177, 168, '修改任务', 'A', NULL, NULL, '/cron/api/v1/task', 'PUT', NULL, 'cron:task:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (178, 168, '删除任务', 'A', NULL, NULL, '/cron/api/v1/task', 'DELETE', NULL, 'cron:task:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (179, 168, '任务状态管理', 'A', NULL, NULL, '/cron/api/v1/task/status', 'PUT', NULL, 'cron:task:update:status', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (180, 168, '任务日志', 'G', NULL, NULL, NULL, NULL, NULL, 'cron:task:log', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (181, 180, '获取任务日志分页', 'A', NULL, NULL, '/cron/api/v1/logs', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (182, 180, '获取任务日志详情', 'A', NULL, NULL, '/cron/api/v1/log', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (183, 1, '通知中心', 'M', 'Notify', 'notification', NULL, NULL, '/notify', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (184, 183, '渠道管理', 'M', 'NotifyChannel', 'mind-mapping', NULL, NULL, '/notify/channel', NULL, '/notify/channel/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (185, 184, '查看渠道', 'A', NULL, NULL, '/notify/api/v1/channels', 'GET', NULL, 'notify:channel:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (186, 184, '新增渠道', 'A', NULL, NULL, '/notify/api/v1/channel', 'POST', NULL, 'notify:channel:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (187, 184, '修改渠道', 'A', NULL, NULL, '/notify/api/v1/channel', 'PUT', NULL, 'notify:channel:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (188, 184, '删除渠道', 'A', NULL, NULL, '/notify/api/v1/channel', 'DELETE', NULL, 'notify:channel:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (189, 183, '通知管理', 'M', 'NotifyManager', 'message', NULL, NULL, '/notify/notify', NULL, '/notify/notify/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (190, 189, '查看通知分组', 'A', NULL, NULL, '/notify/api/v1/notify_categories', 'GET', NULL, 'notify:category:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (191, 189, '新增通知分组', 'A', NULL, NULL, '/notify/api/v1/notify_category', 'POST', NULL, 'notify:category:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (192, 189, '修改通知分组', 'A', NULL, NULL, '/notify/api/v1/notify_category', 'PUT', NULL, 'notify:category:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (193, 189, '删除通知分组', 'A', NULL, NULL, '/notify/api/v1/notify_category', 'DELETE', NULL, 'notify:category:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (194, 189, '查看通知列表', 'A', NULL, NULL, '/notify/api/v1/notifies', 'GET', NULL, 'notify:notify:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (195, 189, '查看指定通知', 'A', NULL, NULL, '/notify/api/v1/notify', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (196, 189, '新增通知', 'A', NULL, NULL, '/notify/api/v1/notify', 'POST', NULL, 'notify:notify:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (197, 189, '修改通知', 'A', NULL, NULL, '/notify/api/v1/notify', 'PUT', NULL, 'notify:notify:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (198, 189, '删除通知', 'A', NULL, NULL, '/notify/api/v1/notify', 'DELETE', NULL, 'notify:notify:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (199, 189, '查看通知模板', 'A', NULL, NULL, '/notify/api/v1/templates', 'GET', NULL, 'notify:template:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (200, 189, '新增通知模板', 'A', NULL, NULL, '/notify/api/v1/template', 'POST', NULL, 'notify:template:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (201, 189, '修改通知模板', 'A', NULL, NULL, '/notify/api/v1/template', 'PUT', NULL, 'notify:template:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (202, 189, '删除通知模板', 'A', NULL, NULL, '/notify/api/v1/template', 'DELETE', NULL, 'notify:template:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (203, 183, '发送日志', 'M', 'NotifyLog', 'history', NULL, NULL, '/notify/log', NULL, '/notify/log/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (204, 203, '查看发送日志', 'A', NULL, NULL, '/notify/api/v1/logs', 'GET', NULL, 'notify:log:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (205, 1, 'AI智能体', 'M', 'SystemPlatformAIAgent', 'robot', NULL, NULL, '/ai-agent', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (206, 205, '密钥管理', 'M', 'ai-agent-secret', 'safe', NULL, NULL, '/ai-agent/secret', NULL, '/ai-agent/secret/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (207, 206, '查看密钥分组', 'A', NULL, NULL, '/ai-agent/api/v1/secret_groups', 'GET', NULL, 'ai-agent:secret:group:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (208, 206, '新增密钥分组', 'A', NULL, NULL, '/ai-agent/api/v1/secret_group', 'POST', NULL, 'ai-agent:secret:group:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (209, 206, '修改密钥分组', 'A', NULL, NULL, '/ai-agent/api/v1/secret_group', 'PUT', NULL, 'ai-agent:secret:group:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (210, 206, '删除密钥分组', 'A', NULL, NULL, '/ai-agent/api/v1/secret_group', 'DELETE', NULL, 'ai-agent:secret:group:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (211, 206, '查看密钥', 'A', NULL, NULL, '/ai-agent/api/v1/secrets', 'GET', NULL, 'ai-agent:secret:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (212, 206, '新增密钥', 'A', NULL, NULL, '/ai-agent/api/v1/secret', 'POST', NULL, 'ai-agent:secret:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (213, 206, '修改密钥', 'A', NULL, NULL, '/ai-agent/api/v1/secret', 'PUT', NULL, 'ai-agent:secret:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (214, 206, '删除密钥', 'A', NULL, NULL, '/ai-agent/api/v1/secret', 'DELETE', NULL, 'ai-agent:secret:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (215, 205, '模型管理', 'M', 'ai-agent-model', 'common', NULL, NULL, '/ai-agent/model', NULL, '/ai-agent/model/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (216, 215, '查看模型', 'A', NULL, NULL, '/ai-agent/api/v1/models', 'GET', NULL, 'ai-agent:model:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (217, 215, '新增模型', 'A', NULL, NULL, '/ai-agent/api/v1/model', 'POST', NULL, 'ai-agent:model:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (218, 215, '修改模型', 'A', NULL, NULL, '/ai-agent/api/v1/model', 'PUT', NULL, 'ai-agent:model:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (219, 215, '删除模型', 'A', NULL, NULL, '/ai-agent/api/v1/model', 'DELETE', NULL, 'ai-agent:model:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (220, 205, '工具管理', 'M', 'ai-agent-tool', 'tool', NULL, NULL, '/ai-agent/tool', NULL, '/ai-agent/tool/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (221, 220, '查看工具', 'A', NULL, NULL, '/ai-agent/api/v1/tools', 'GET', NULL, 'ai-agent:tool:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (222, 220, '新增工具', 'A', NULL, NULL, '/ai-agent/api/v1/tool', 'POST', NULL, 'ai-agent:tool:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (223, 220, '修改工具', 'A', NULL, NULL, '/ai-agent/api/v1/tool', 'PUT', NULL, 'ai-agent:tool:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (224, 220, '删除工具', 'A', NULL, NULL, '/ai-agent/api/v1/tool', 'DELETE', NULL, 'ai-agent:tool:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (225, 205, '知识库管理', 'M', 'ai-agent-knowledge', 'storage', NULL, NULL, '/ai-agent/knowledge', NULL, '/ai-agent/knowledge/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (226, 225, '查看知识库', 'A', NULL, NULL, '/ai-agent/api/v1/knowledges', 'GET', NULL, 'ai-agent:knowledge:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (227, 225, '新增知识库', 'A', NULL, NULL, '/ai-agent/api/v1/knowledge', 'POST', NULL, 'ai-agent:knowledge:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (228, 225, '修改知识库', 'A', NULL, NULL, '/ai-agent/api/v1/knowledge', 'PUT', NULL, 'ai-agent:knowledge:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (229, 225, '删除知识库', 'A', NULL, NULL, '/ai-agent/api/v1/knowledge', 'DELETE', NULL, 'ai-agent:knowledge:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (230, 205, '提示词评估', 'M', 'ai-agent-assessment', 'bulb', NULL, NULL, '/ai-agent/assessment', NULL, '/ai-agent/assessment/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (231, 230, '查看assessment评估记录', 'A', NULL, NULL, '/ai-agent/api/v1/assessments', 'GET', NULL, 'ai-agent:assessment:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (232, 230, '新增assessment评估记录', 'A', NULL, NULL, '/ai-agent/api/v1/assessment', 'POST', NULL, 'ai-agent:assessment:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (233, 230, '修改assessment评估记录', 'A', NULL, NULL, '/ai-agent/api/v1/assessment', 'PUT', NULL, 'ai-agent:assessment:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (234, 230, '删除assessment评估记录', 'A', NULL, NULL, '/ai-agent/api/v1/assessment', 'DELETE', NULL, 'ai-agent:assessment:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (235, 205, '评估面板', 'M', 'ai-agent-assessment-operator', NULL, NULL, NULL, '/ai-agent/assessment/operator', NULL, '/ai-agent/assessment/operator/index', NULL, 0, 1, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (236, 235, '获取assessment评估记录会话历史', 'A', NULL, NULL, '/ai-agent/api/v1/assessment/conversation', 'GET', NULL, 'ai-agent:assessment:conversation:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (237, 235, '新增assessment评估记录会话历史', 'A', NULL, NULL, '/ai-agent/api/v1/assessment/conversation', 'POST', NULL, 'ai-agent:assessment:conversation:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (238, 205, '在线聊天', 'M', 'ai-agent-chat', 'message', NULL, NULL, '/ai-agent/chat', NULL, '/ai-agent/chat/index', NULL, 0, NULL, NULL, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (239, 238, '获取会话历史列表', 'A', NULL, NULL, '/ai-agent/api/v1/chat', 'GET', NULL, 'ai-agent:chat:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (240, 238, '获取指定会话的历史记录列表', 'A', NULL, NULL, '/ai-agent/api/v1/chat/messages', 'GET', NULL, 'ai-agent:chat:message:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (241, 238, 'chat会话', 'A', NULL, NULL, '/ai-agent/api/v1/chat/conversation', 'POST', NULL, 'ai-agent:chat:conversation', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (242, 238, '更新chat会话', 'A', NULL, NULL, '/ai-agent/api/v1/chat', 'update', NULL, 'ai-agent:chat:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (243, 238, '删除chat会话', 'A', NULL, NULL, '/ai-agent/api/v1/chat', 'DELETE', NULL, 'ai-agent:chat:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (244, 0, '应用平台', 'R', 'AppPlatform', 'apps', NULL, NULL, '/app', NULL, 'Layout', NULL, 2, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (245, 244, '首页面板', 'M', 'AppDashboard', 'dashboard', NULL, NULL, '/app/dashboard', NULL, '/dashboard/index', NULL, 0, NULL, 1, 1, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (246, 244, 'AIGC视频', 'M', 'AgentServerAIGC', 'live-broadcast', NULL, NULL, '/agent-server', NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (247, 246, '提示词管理', 'M', 'AgentServerAIGCPrompt', 'info-circle', NULL, NULL, '/agent-server/aigc/prompt', NULL, '/agent-server/aigc/prompt/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (248, 247, '查看提示词', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/prompts', 'GET', NULL, 'agent-server:aigc:prompt:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (249, 247, '新增提示词', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/prompt', 'POST', NULL, 'agent-server:aigc:prompt:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (250, 247, '更新提示词', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/prompt', 'PUT', NULL, 'agent-server:aigc:prompt:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (251, 247, '删除提示词', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/prompt', 'DELETE', NULL, 'agent-server:aigc:prompt:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (252, 247, '查看模型列表', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/prompt/models', 'GET', NULL, NULL, NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (253, 247, '提交生成', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/prompt/chat', 'POST', NULL, 'agent-server:aigc:prompt:chat', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (254, 246, '标签管理', 'M', 'AgentServerAIGCTag', 'tag', NULL, NULL, '/agent-server/aigc/tag', NULL, '/agent-server/aigc/tag/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (255, 254, '查看标签', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/tags', 'GET', NULL, 'agent-server:aigc:tag:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (256, 254, '新增标签', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/tag', 'POST', NULL, 'agent-server:aigc:tag:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (257, 254, '更新标签', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/tag', 'PUT', NULL, 'agent-server:aigc:tag:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (258, 254, '删除标签', 'A', NULL, NULL, '/agent-server/admin/v1/aigccontent/tag', 'DELETE', NULL, 'agent-server:aigc:tag:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (259, 246, '智能体管理', 'M', 'AgentServerAIGCAgent', 'robot', NULL, NULL, '/agent-server/aigc/agent', NULL, '/agent-server/aigc/agent/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (260, 259, '查看智能体', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/agents', 'GET', NULL, 'agent-server:aigc:agent:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (261, 259, '新增智能体', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/agent', 'POST', NULL, 'agent-server:aigc:agent:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (262, 259, '更新智能体', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/agent', 'PUT', NULL, 'agent-server:aigc:agent:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (263, 259, '删除智能体', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/agent', 'DELETE', NULL, 'agent-server:aigc:agent:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (264, 259, '查看预产', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/examples', 'GET', NULL, 'agent-server:aigc:example:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (265, 259, '新增预产', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/example', 'POST', NULL, 'agent-server:aigc:example:add', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (266, 259, '更新预产', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/example', 'PUT', NULL, 'agent-server:aigc:example:update', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (267, 259, '删除预产', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/example', 'DELETE', NULL, 'agent-server:aigc:example:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (268, 246, '视频管理', 'M', 'AgentServerAIGCVideo', 'live-broadcast', NULL, NULL, '/agent-server/aigc/video', NULL, '/agent-server/aigc/video/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (269, 268, '查看视频', 'A', NULL, NULL, '/agent-server/admin/v1/aigcvideo/videos', 'GET', NULL, 'agent-server:aigc:video:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (270, 244, '通用任务', 'M', 'AgentServerTask', 'public', NULL, NULL, '/agent-server/task', NULL, '/agent-server/task/index', NULL, 0, NULL, 1, NULL, 1, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (271, 270, '查看任务', 'A', NULL, NULL, '/agent-server/admin/v1/tasks', 'GET', NULL, 'agent-server:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (272, 270, '查看可执行的任务', 'A', NULL, NULL, '/agent-server/task/v1/tasks', 'GET', NULL, 'agent-server:task:query', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (273, 270, '执行任务', 'A', NULL, NULL, '/agent-server/task/v1/task', 'POST', NULL, 'agent-server:task:exec', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (274, 270, '中止任务', 'A', NULL, NULL, '/agent-server/task/v1/task/cancel', 'POST', NULL, 'agent-server:task:cancel', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
INSERT INTO "public"."menu_20250811" ("id", "parent_id", "title", "type", "keyword", "icon", "api", "method", "path", "permission", "component", "redirect", "weight", "is_hidden", "is_cache", "is_home", "is_affix", "created_at", "updated_at") VALUES (275, 270, '删除任务', 'A', NULL, NULL, '/agent-server/admin/v1/task', 'DELETE', NULL, 'agent-server:task:delete', NULL, NULL, 0, NULL, NULL, NULL, NULL, 1751133244, 1751133244);
COMMIT;

-- ----------------------------
-- Table structure for menu_closure_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."menu_closure_20250811";
CREATE TABLE "public"."menu_closure_20250811" (
  "id" int8,
  "parent" int8,
  "children" int8
)
;
ALTER TABLE "public"."menu_closure_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of menu_closure_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35619, 1, 5);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35620, 3, 5);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35621, 4, 5);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35622, 1, 6);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35623, 3, 6);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35624, 4, 6);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35625, 1, 7);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35626, 3, 7);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35627, 4, 7);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35628, 1, 8);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35629, 3, 8);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35630, 4, 8);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35631, 1, 9);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35632, 3, 9);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35633, 4, 9);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35634, 1, 10);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35635, 3, 10);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35636, 4, 10);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35637, 1, 11);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35638, 3, 11);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35639, 4, 11);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35640, 1, 12);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35641, 3, 12);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35642, 4, 12);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35643, 1, 13);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35644, 3, 13);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35645, 4, 13);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35646, 1, 14);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35647, 3, 14);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35648, 4, 14);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35649, 1, 15);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35650, 3, 15);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35651, 4, 15);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35652, 1, 16);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35653, 3, 16);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35654, 4, 16);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35655, 1, 17);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35656, 3, 17);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35657, 4, 17);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35658, 1, 18);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35659, 3, 18);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35660, 4, 18);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35661, 1, 19);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35662, 3, 19);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35663, 4, 19);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35664, 1, 20);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35665, 3, 20);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35666, 4, 20);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35667, 1, 21);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35668, 3, 21);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35669, 4, 21);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35670, 1, 22);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35671, 3, 22);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35672, 4, 22);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35673, 1, 23);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35674, 3, 23);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35675, 4, 23);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35676, 1, 24);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35677, 3, 24);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35678, 4, 24);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35679, 1, 25);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35680, 3, 25);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35681, 4, 25);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35682, 1, 27);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35683, 3, 27);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35684, 26, 27);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35685, 1, 28);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35686, 3, 28);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35687, 26, 28);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35688, 1, 29);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35689, 3, 29);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35690, 26, 29);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35691, 1, 30);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35692, 3, 30);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35693, 26, 30);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35694, 1, 31);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35695, 3, 31);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35696, 26, 31);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35697, 1, 32);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35698, 3, 32);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35699, 26, 32);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35700, 1, 33);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35701, 3, 33);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35702, 26, 33);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35703, 1, 34);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35704, 3, 34);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35705, 26, 34);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35706, 1, 35);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35707, 3, 35);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35708, 26, 35);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35709, 1, 37);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35710, 3, 37);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35711, 36, 37);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35712, 1, 38);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35713, 3, 38);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35714, 36, 38);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35715, 1, 39);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35716, 3, 39);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35717, 36, 39);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35718, 1, 40);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35719, 3, 40);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35720, 36, 40);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35721, 1, 42);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35722, 3, 42);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35723, 41, 42);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35724, 1, 43);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35725, 3, 43);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35726, 41, 43);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35727, 1, 44);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35728, 3, 44);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35729, 41, 44);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35730, 1, 45);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35731, 3, 45);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35732, 41, 45);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35733, 1, 46);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35734, 3, 46);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35735, 41, 46);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35736, 1, 47);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35737, 3, 47);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35738, 41, 47);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35739, 1, 48);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35740, 3, 48);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35741, 41, 48);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35742, 1, 50);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35743, 3, 50);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35744, 49, 50);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35745, 1, 51);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35746, 3, 51);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35747, 49, 51);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35748, 1, 52);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35749, 3, 52);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35750, 49, 52);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35751, 1, 53);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35752, 3, 53);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35753, 49, 53);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35754, 1, 61);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35755, 3, 61);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35756, 54, 61);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35757, 60, 61);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35758, 1, 62);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35759, 3, 62);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35760, 54, 62);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35761, 60, 62);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35762, 1, 55);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35763, 3, 55);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35764, 54, 55);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35765, 1, 56);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35766, 3, 56);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35767, 54, 56);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35768, 1, 57);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35769, 3, 57);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35770, 54, 57);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35771, 1, 58);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35772, 3, 58);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35773, 54, 58);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35774, 1, 59);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35775, 3, 59);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35776, 54, 59);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35777, 1, 60);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35778, 3, 60);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35779, 54, 60);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35780, 1, 64);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35781, 3, 64);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35782, 63, 64);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35783, 1, 65);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35784, 3, 65);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35785, 63, 65);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35786, 1, 66);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35787, 3, 66);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35788, 63, 66);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35789, 1, 67);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35790, 3, 67);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35791, 63, 67);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35792, 1, 68);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35793, 3, 68);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35794, 63, 68);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35795, 1, 69);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35796, 3, 69);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35797, 63, 69);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35798, 1, 71);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35799, 3, 71);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35800, 70, 71);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35801, 1, 72);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35802, 3, 72);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35803, 70, 72);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35804, 1, 73);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35805, 3, 73);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35806, 70, 73);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35807, 1, 74);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35808, 3, 74);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35809, 70, 74);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35810, 1, 76);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35811, 3, 76);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35812, 75, 76);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35813, 1, 77);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35814, 3, 77);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35815, 75, 77);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35816, 1, 78);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35817, 3, 78);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35818, 75, 78);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35819, 1, 79);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35820, 3, 79);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35821, 75, 79);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35822, 1, 80);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35823, 3, 80);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35824, 75, 80);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35825, 1, 82);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35826, 3, 82);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35827, 81, 82);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35828, 1, 83);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35829, 3, 83);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35830, 81, 83);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35831, 1, 84);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35832, 3, 84);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35833, 81, 84);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35834, 1, 85);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35835, 3, 85);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35836, 81, 85);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35837, 1, 86);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35838, 3, 86);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35839, 81, 86);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35840, 1, 87);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35841, 3, 87);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35842, 81, 87);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35843, 1, 89);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35844, 3, 89);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35845, 88, 89);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35846, 1, 90);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35847, 3, 90);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35848, 88, 90);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35849, 1, 91);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35850, 3, 91);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35851, 88, 91);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35852, 1, 92);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35853, 3, 92);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35854, 88, 92);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35855, 1, 93);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35856, 3, 93);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35857, 88, 93);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35858, 1, 94);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35859, 3, 94);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35860, 88, 94);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35861, 1, 95);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35862, 3, 95);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35863, 88, 95);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35864, 1, 96);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35865, 3, 96);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35866, 88, 96);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35867, 1, 98);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35868, 3, 98);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35869, 97, 98);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35870, 1, 4);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35871, 3, 4);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35872, 1, 26);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35873, 3, 26);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35874, 1, 36);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35875, 3, 36);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35876, 1, 41);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35877, 3, 41);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35878, 1, 49);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35879, 3, 49);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35880, 1, 54);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35881, 3, 54);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35882, 1, 63);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35883, 3, 63);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35884, 1, 70);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35885, 3, 70);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35886, 1, 75);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35887, 3, 75);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35888, 1, 81);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35889, 3, 81);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35890, 1, 88);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35891, 3, 88);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35892, 1, 97);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35893, 3, 97);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35894, 1, 102);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35895, 99, 102);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35896, 100, 102);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35897, 101, 102);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35898, 1, 103);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35899, 99, 103);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35900, 100, 103);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35901, 101, 103);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35902, 1, 104);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35903, 99, 104);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35904, 100, 104);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35905, 101, 104);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35906, 1, 105);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35907, 99, 105);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35908, 100, 105);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35909, 101, 105);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35910, 1, 107);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35911, 99, 107);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35912, 100, 107);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35913, 106, 107);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35914, 1, 108);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35915, 99, 108);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35916, 100, 108);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35917, 106, 108);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35918, 1, 109);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35919, 99, 109);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35920, 100, 109);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35921, 106, 109);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35922, 1, 101);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35923, 99, 101);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35924, 100, 101);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35925, 1, 106);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35926, 99, 106);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35927, 100, 106);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35928, 1, 111);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35929, 99, 111);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35930, 110, 111);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35931, 1, 112);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35932, 99, 112);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35933, 110, 112);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35934, 1, 113);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35935, 99, 113);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35936, 110, 113);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35937, 1, 100);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35938, 99, 100);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35939, 1, 110);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35940, 99, 110);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35941, 1, 116);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35942, 114, 116);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35943, 115, 116);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35944, 1, 117);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35945, 114, 117);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35946, 115, 117);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35947, 1, 118);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35948, 114, 118);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35949, 115, 118);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35950, 1, 119);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35951, 114, 119);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35952, 115, 119);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35953, 1, 120);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35954, 114, 120);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35955, 115, 120);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35956, 1, 121);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35957, 114, 121);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35958, 115, 121);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35959, 1, 122);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35960, 114, 122);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35961, 115, 122);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35962, 1, 123);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35963, 114, 123);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35964, 115, 123);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35965, 1, 125);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35966, 114, 125);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35967, 124, 125);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35968, 1, 126);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35969, 114, 126);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35970, 124, 126);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35971, 1, 127);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35972, 114, 127);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35973, 124, 127);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35974, 1, 128);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35975, 114, 128);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35976, 124, 128);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35977, 1, 129);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35978, 114, 129);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35979, 124, 129);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35980, 1, 130);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35981, 114, 130);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35982, 124, 130);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35983, 1, 132);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35984, 114, 132);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35985, 131, 132);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35986, 1, 133);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35987, 114, 133);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35988, 131, 133);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35989, 1, 134);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35990, 114, 134);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35991, 131, 134);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35992, 1, 135);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35993, 114, 135);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35994, 131, 135);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35995, 1, 136);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35996, 114, 136);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35997, 131, 136);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35998, 1, 137);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (35999, 114, 137);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36000, 131, 137);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36001, 1, 139);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36002, 114, 139);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36003, 138, 139);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36004, 1, 140);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36005, 114, 140);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36006, 138, 140);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36007, 1, 141);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36008, 114, 141);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36009, 138, 141);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36010, 1, 142);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36011, 114, 142);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36012, 138, 142);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36013, 1, 143);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36014, 114, 143);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36015, 138, 143);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36016, 1, 144);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36017, 114, 144);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36018, 138, 144);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36019, 1, 147);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36020, 114, 147);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36021, 145, 147);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36022, 146, 147);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36023, 1, 148);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36024, 114, 148);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36025, 145, 148);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36026, 146, 148);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36027, 1, 149);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36028, 114, 149);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36029, 145, 149);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36030, 146, 149);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36031, 1, 150);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36032, 114, 150);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36033, 145, 150);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36034, 146, 150);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36035, 1, 151);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36036, 114, 151);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36037, 145, 151);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36038, 146, 151);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36039, 1, 152);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36040, 114, 152);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36041, 145, 152);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36042, 146, 152);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36043, 1, 153);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36044, 114, 153);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36045, 145, 153);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36046, 146, 153);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36047, 1, 155);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36048, 114, 155);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36049, 145, 155);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36050, 154, 155);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36051, 1, 156);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36052, 114, 156);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36053, 145, 156);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36054, 154, 156);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36055, 1, 146);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36056, 114, 146);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36057, 145, 146);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36058, 1, 154);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36059, 114, 154);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36060, 145, 154);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36061, 1, 115);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36062, 114, 115);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36063, 1, 124);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36064, 114, 124);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36065, 1, 131);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36066, 114, 131);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36067, 1, 138);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36068, 114, 138);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36069, 1, 145);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36070, 114, 145);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36071, 1, 159);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36072, 157, 159);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36073, 158, 159);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36074, 1, 160);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36075, 157, 160);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36076, 158, 160);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36077, 1, 161);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36078, 157, 161);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36079, 158, 161);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36080, 1, 162);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36081, 157, 162);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36082, 158, 162);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36083, 1, 163);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36084, 157, 163);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36085, 158, 163);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36086, 1, 164);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36087, 157, 164);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36088, 158, 164);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36089, 1, 165);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36090, 157, 165);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36091, 158, 165);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36092, 1, 166);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36093, 157, 166);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36094, 158, 166);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36095, 1, 167);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36096, 157, 167);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36097, 158, 167);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36098, 1, 181);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36099, 157, 181);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36100, 168, 181);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36101, 180, 181);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36102, 1, 182);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36103, 157, 182);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36104, 168, 182);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36105, 180, 182);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36106, 1, 169);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36107, 157, 169);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36108, 168, 169);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36109, 1, 170);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36110, 157, 170);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36111, 168, 170);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36112, 1, 171);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36113, 157, 171);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36114, 168, 171);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36115, 1, 172);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36116, 157, 172);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36117, 168, 172);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36118, 1, 173);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36119, 157, 173);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36120, 168, 173);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36121, 1, 174);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36122, 157, 174);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36123, 168, 174);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36124, 1, 175);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36125, 157, 175);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36126, 168, 175);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36127, 1, 176);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36128, 157, 176);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36129, 168, 176);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36130, 1, 177);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36131, 157, 177);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36132, 168, 177);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36133, 1, 178);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36134, 157, 178);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36135, 168, 178);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36136, 1, 179);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36137, 157, 179);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36138, 168, 179);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36139, 1, 180);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36140, 157, 180);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36141, 168, 180);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36142, 1, 158);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36143, 157, 158);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36144, 1, 168);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36145, 157, 168);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36146, 1, 185);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36147, 183, 185);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36148, 184, 185);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36149, 1, 186);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36150, 183, 186);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36151, 184, 186);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36152, 1, 187);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36153, 183, 187);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36154, 184, 187);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36155, 1, 188);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36156, 183, 188);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36157, 184, 188);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36158, 1, 190);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36159, 183, 190);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36160, 189, 190);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36161, 1, 191);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36162, 183, 191);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36163, 189, 191);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36164, 1, 192);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36165, 183, 192);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36166, 189, 192);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36167, 1, 193);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36168, 183, 193);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36169, 189, 193);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36170, 1, 194);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36171, 183, 194);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36172, 189, 194);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36173, 1, 195);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36174, 183, 195);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36175, 189, 195);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36176, 1, 196);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36177, 183, 196);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36178, 189, 196);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36179, 1, 197);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36180, 183, 197);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36181, 189, 197);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36182, 1, 198);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36183, 183, 198);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36184, 189, 198);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36185, 1, 199);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36186, 183, 199);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36187, 189, 199);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36188, 1, 200);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36189, 183, 200);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36190, 189, 200);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36191, 1, 201);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36192, 183, 201);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36193, 189, 201);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36194, 1, 202);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36195, 183, 202);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36196, 189, 202);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36197, 1, 204);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36198, 183, 204);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36199, 203, 204);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36200, 1, 184);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36201, 183, 184);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36202, 1, 189);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36203, 183, 189);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36204, 1, 203);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36205, 183, 203);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36206, 1, 207);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36207, 205, 207);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36208, 206, 207);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36209, 1, 208);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36210, 205, 208);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36211, 206, 208);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36212, 1, 209);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36213, 205, 209);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36214, 206, 209);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36215, 1, 210);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36216, 205, 210);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36217, 206, 210);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36218, 1, 211);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36219, 205, 211);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36220, 206, 211);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36221, 1, 212);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36222, 205, 212);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36223, 206, 212);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36224, 1, 213);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36225, 205, 213);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36226, 206, 213);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36227, 1, 214);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36228, 205, 214);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36229, 206, 214);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36230, 1, 216);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36231, 205, 216);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36232, 215, 216);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36233, 1, 217);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36234, 205, 217);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36235, 215, 217);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36236, 1, 218);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36237, 205, 218);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36238, 215, 218);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36239, 1, 219);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36240, 205, 219);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36241, 215, 219);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36242, 1, 221);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36243, 205, 221);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36244, 220, 221);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36245, 1, 222);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36246, 205, 222);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36247, 220, 222);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36248, 1, 223);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36249, 205, 223);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36250, 220, 223);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36251, 1, 224);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36252, 205, 224);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36253, 220, 224);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36254, 1, 226);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36255, 205, 226);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36256, 225, 226);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36257, 1, 227);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36258, 205, 227);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36259, 225, 227);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36260, 1, 228);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36261, 205, 228);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36262, 225, 228);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36263, 1, 229);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36264, 205, 229);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36265, 225, 229);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36266, 1, 231);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36267, 205, 231);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36268, 230, 231);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36269, 1, 232);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36270, 205, 232);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36271, 230, 232);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36272, 1, 233);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36273, 205, 233);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36274, 230, 233);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36275, 1, 234);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36276, 205, 234);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36277, 230, 234);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36278, 1, 236);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36279, 205, 236);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36280, 235, 236);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36281, 1, 237);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36282, 205, 237);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36283, 235, 237);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36284, 1, 239);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36285, 205, 239);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36286, 238, 239);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36287, 1, 240);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36288, 205, 240);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36289, 238, 240);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36290, 1, 241);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36291, 205, 241);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36292, 238, 241);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36293, 1, 242);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36294, 205, 242);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36295, 238, 242);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36296, 1, 243);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36297, 205, 243);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36298, 238, 243);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36299, 1, 206);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36300, 205, 206);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36301, 1, 215);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36302, 205, 215);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36303, 1, 220);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36304, 205, 220);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36305, 1, 225);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36306, 205, 225);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36307, 1, 230);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36308, 205, 230);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36309, 1, 235);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36310, 205, 235);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36311, 1, 238);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36312, 205, 238);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36313, 1, 2);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36314, 1, 3);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36315, 1, 99);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36316, 1, 114);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36317, 1, 157);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36318, 1, 183);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36319, 1, 205);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36320, 244, 248);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36321, 246, 248);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36322, 247, 248);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36323, 244, 249);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36324, 246, 249);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36325, 247, 249);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36326, 244, 250);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36327, 246, 250);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36328, 247, 250);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36329, 244, 251);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36330, 246, 251);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36331, 247, 251);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36332, 244, 252);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36333, 246, 252);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36334, 247, 252);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36335, 244, 253);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36336, 246, 253);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36337, 247, 253);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36338, 244, 255);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36339, 246, 255);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36340, 254, 255);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36341, 244, 256);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36342, 246, 256);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36343, 254, 256);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36344, 244, 257);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36345, 246, 257);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36346, 254, 257);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36347, 244, 258);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36348, 246, 258);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36349, 254, 258);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36350, 244, 260);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36351, 246, 260);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36352, 259, 260);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36353, 244, 261);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36354, 246, 261);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36355, 259, 261);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36356, 244, 262);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36357, 246, 262);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36358, 259, 262);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36359, 244, 263);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36360, 246, 263);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36361, 259, 263);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36362, 244, 264);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36363, 246, 264);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36364, 259, 264);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36365, 244, 265);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36366, 246, 265);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36367, 259, 265);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36368, 244, 266);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36369, 246, 266);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36370, 259, 266);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36371, 244, 267);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36372, 246, 267);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36373, 259, 267);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36374, 244, 269);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36375, 246, 269);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36376, 268, 269);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36377, 244, 247);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36378, 246, 247);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36379, 244, 254);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36380, 246, 254);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36381, 244, 259);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36382, 246, 259);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36383, 244, 268);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36384, 246, 268);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36385, 244, 271);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36386, 270, 271);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36387, 244, 272);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36388, 270, 272);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36389, 244, 273);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36390, 270, 273);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36391, 244, 274);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36392, 270, 274);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36393, 244, 275);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36394, 270, 275);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36395, 244, 245);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36396, 244, 246);
INSERT INTO "public"."menu_closure_20250811" ("id", "parent", "children") VALUES (36397, 244, 270);
COMMIT;

-- ----------------------------
-- Table structure for oauth_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."oauth_20250811";
CREATE TABLE "public"."oauth_20250811" (
  "id" int8,
  "user_id" int8,
  "channel_id" int8,
  "oid" varchar(254) COLLATE "pg_catalog"."default",
  "token" varchar(254) COLLATE "pg_catalog"."default",
  "logged_at" int8,
  "extra" text COLLATE "pg_catalog"."default",
  "expired_at" int8,
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."oauth_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of oauth_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."oauth_20250811" ("id", "user_id", "channel_id", "oid", "token", "logged_at", "extra", "expired_at", "created_at", "updated_at") VALUES (3, 1, 1, '17871318', 'c39cbbbb76d7bda8c04b0ca3df6fa8fc27180e11', 1751113438, '{"union_id":""}', 1751197576, 1749994770, NULL);
INSERT INTO "public"."oauth_20250811" ("id", "user_id", "channel_id", "oid", "token", "logged_at", "extra", "expired_at", "created_at", "updated_at") VALUES (4, 1, 2, 'oC8iT7AFTJlWKvnJUDBU7VSJwNBg', '93_y2fEn5cNGE5n-LXPn_otHyCjq0R8Zwj-Ks4kewz26YzrjoLUoiJQ8to4UWfOTD3mGh3fOYBnrfRl3l0rD9RHMgAuLYjAIzDRqbAxktvaATA', 1750089796, '{"union_id":""}', 1750089796, 1750089796, NULL);
INSERT INTO "public"."oauth_20250811" ("id", "user_id", "channel_id", "oid", "token", "logged_at", "extra", "expired_at", "created_at", "updated_at") VALUES (6, 1, 3, '1280291001@qq.com', NULL, 1751115190, '{"union_id":""}', 1751115190, 1751115190, NULL);
COMMIT;

-- ----------------------------
-- Table structure for resource_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."resource_20250811";
CREATE TABLE "public"."resource_20250811" (
  "id" int8,
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "department_" int8,
  "resource_id" int8
)
;
ALTER TABLE "public"."resource_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of resource_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."resource_20250811" ("id", "keyword", "department_", "resource_id") VALUES (1, 'cfg_env', 1, 2);
INSERT INTO "public"."resource_20250811" ("id", "keyword", "department_", "resource_id") VALUES (2, 'cfg_server', 1, 1);
INSERT INTO "public"."resource_20250811" ("id", "keyword", "department_", "resource_id") VALUES (3, 'cfg_server', 5, 1);
INSERT INTO "public"."resource_20250811" ("id", "keyword", "department_", "resource_id") VALUES (4, 'cfg_server', 6, 1);
INSERT INTO "public"."resource_20250811" ("id", "keyword", "department_", "resource_id") VALUES (5, 'cfg_server', 9, 1);
INSERT INTO "public"."resource_20250811" ("id", "keyword", "department_", "resource_id") VALUES (6, 'uc_app', 1, 3);
COMMIT;

-- ----------------------------
-- Table structure for role_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."role_20250811";
CREATE TABLE "public"."role_20250811" (
  "id" int8,
  "parent_id" int8,
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "keyword" varchar(254) COLLATE "pg_catalog"."default",
  "status" int4,
  "description" varchar(254) COLLATE "pg_catalog"."default",
  "department_" text COLLATE "pg_catalog"."default",
  "data_scope" varchar(254) COLLATE "pg_catalog"."default",
  "job_ids" text COLLATE "pg_catalog"."default",
  "job_scope" varchar(254) COLLATE "pg_catalog"."default",
  "created_at" int8,
  "updated_at" int8,
  "deleted_at" int8
)
;
ALTER TABLE "public"."role_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of role_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."role_20250811" ("id", "parent_id", "name", "keyword", "status", "description", "department_", "data_scope", "job_ids", "job_scope", "created_at", "updated_at", "deleted_at") VALUES (1, 0, '超级管理员', 'superAdmin', 1, '超级管理员', NULL, 'ASSIGN_ALL', NULL, NULL, 1713706137, 1713706137, 0);
INSERT INTO "public"."role_20250811" ("id", "parent_id", "name", "keyword", "status", "description", "department_", "data_scope", "job_ids", "job_scope", "created_at", "updated_at", "deleted_at") VALUES (5, 1, '21', '3', 1, '412', '1', 'CUR', '20', 'ALL', 1719464519, 1751567747, 0);
INSERT INTO "public"."role_20250811" ("id", "parent_id", "name", "keyword", "status", "description", "department_", "data_scope", "job_ids", "job_scope", "created_at", "updated_at", "deleted_at") VALUES (9, 5, '测试', 'test', 1, '1', NULL, 'CUR', NULL, NULL, 1751562296, 1751562296, 0);
COMMIT;

-- ----------------------------
-- Table structure for role_closure_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."role_closure_20250811";
CREATE TABLE "public"."role_closure_20250811" (
  "id" int8,
  "parent" int8,
  "children" int8
)
;
ALTER TABLE "public"."role_closure_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of role_closure_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."role_closure_20250811" ("id", "parent", "children") VALUES (5, 1, 5);
INSERT INTO "public"."role_closure_20250811" ("id", "parent", "children") VALUES (12, 5, 9);
INSERT INTO "public"."role_closure_20250811" ("id", "parent", "children") VALUES (13, 1, 9);
COMMIT;

-- ----------------------------
-- Table structure for role_menu_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."role_menu_20250811";
CREATE TABLE "public"."role_menu_20250811" (
  "id" int8,
  "role_id" int8,
  "menu_id" int8
)
;
ALTER TABLE "public"."role_menu_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of role_menu_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3939, 5, 1);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3940, 5, 2);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3941, 5, 3);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3942, 5, 26);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3943, 5, 27);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3944, 5, 28);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3945, 5, 29);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3946, 5, 30);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3947, 5, 31);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3948, 5, 32);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3949, 5, 33);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3950, 5, 34);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3951, 5, 35);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3952, 5, 36);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3953, 5, 37);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3954, 5, 38);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3955, 5, 39);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3956, 5, 40);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3957, 5, 41);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3958, 5, 42);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3959, 5, 43);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3960, 5, 44);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3961, 5, 45);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3962, 5, 46);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3963, 5, 47);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3964, 5, 48);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3965, 5, 49);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3966, 5, 50);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3967, 5, 51);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3968, 5, 52);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3969, 5, 53);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3970, 5, 54);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3971, 5, 55);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3972, 5, 56);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3973, 5, 57);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3974, 5, 58);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3975, 5, 59);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3976, 5, 60);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3977, 5, 61);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3978, 5, 62);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3979, 5, 63);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3980, 5, 64);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3981, 5, 65);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3982, 5, 66);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3983, 5, 67);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3984, 5, 68);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3985, 5, 69);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3986, 5, 70);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3987, 5, 71);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3988, 5, 72);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3989, 5, 73);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3990, 5, 74);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3991, 5, 75);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3992, 5, 76);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3993, 5, 77);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3994, 5, 78);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3995, 5, 79);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3996, 5, 80);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3997, 5, 81);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3998, 5, 82);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (3999, 5, 83);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4000, 5, 84);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4001, 5, 85);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4002, 5, 86);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4003, 5, 87);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4004, 5, 88);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4005, 5, 89);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4006, 5, 90);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4007, 5, 91);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4008, 5, 92);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4009, 5, 93);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4010, 5, 94);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4011, 5, 95);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4012, 5, 96);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4013, 5, 97);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4014, 5, 98);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4015, 5, 99);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4016, 5, 100);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4017, 5, 101);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4018, 5, 102);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4019, 5, 103);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4020, 5, 104);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4021, 5, 105);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4022, 5, 106);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4023, 5, 107);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4024, 5, 108);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4025, 5, 109);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4026, 5, 110);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4027, 5, 111);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4028, 5, 112);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4029, 5, 113);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4030, 5, 114);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4031, 5, 115);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4032, 5, 116);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4033, 5, 117);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4034, 5, 118);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4035, 5, 119);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4036, 5, 120);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4037, 5, 121);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4038, 5, 122);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4039, 5, 123);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4040, 5, 124);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4041, 5, 125);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4042, 5, 126);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4043, 5, 127);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4044, 5, 128);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4045, 5, 129);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4046, 5, 130);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4047, 5, 131);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4048, 5, 132);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4049, 5, 133);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4050, 5, 134);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4051, 5, 135);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4052, 5, 136);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4053, 5, 137);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4054, 5, 138);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4055, 5, 139);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4056, 5, 140);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4057, 5, 141);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4058, 5, 142);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4059, 5, 143);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4060, 5, 144);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4061, 5, 145);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4062, 5, 146);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4063, 5, 147);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4064, 5, 148);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4065, 5, 149);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4066, 5, 150);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4067, 5, 151);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4068, 5, 152);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4069, 5, 153);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4070, 5, 154);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4071, 5, 155);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4072, 5, 156);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4073, 5, 157);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4074, 5, 158);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4075, 5, 159);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4076, 5, 160);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4077, 5, 161);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4078, 5, 162);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4079, 5, 163);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4080, 5, 164);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4081, 5, 165);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4082, 5, 166);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4083, 5, 167);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4084, 5, 168);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4085, 5, 169);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4086, 5, 170);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4087, 5, 171);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4088, 5, 172);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4089, 5, 173);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4090, 5, 174);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4091, 5, 175);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4092, 5, 176);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4093, 5, 177);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4094, 5, 178);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4095, 5, 179);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4096, 5, 180);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4097, 5, 181);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4098, 5, 182);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4099, 5, 183);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4100, 5, 184);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4101, 5, 185);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4102, 5, 186);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4103, 5, 187);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4104, 5, 188);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4105, 5, 189);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4106, 5, 190);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4107, 5, 191);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4108, 5, 192);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4109, 5, 193);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4110, 5, 194);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4111, 5, 195);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4112, 5, 196);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4113, 5, 197);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4114, 5, 198);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4115, 5, 199);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4116, 5, 200);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4117, 5, 201);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4118, 5, 202);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4119, 5, 203);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4120, 5, 204);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4121, 5, 205);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4122, 5, 206);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4123, 5, 207);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4124, 5, 208);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4125, 5, 209);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4126, 5, 210);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4127, 5, 211);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4128, 5, 212);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4129, 5, 213);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4130, 5, 214);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4131, 5, 215);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4132, 5, 216);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4133, 5, 217);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4134, 5, 218);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4135, 5, 219);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4136, 5, 220);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4137, 5, 221);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4138, 5, 222);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4139, 5, 223);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4140, 5, 224);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4141, 5, 225);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4142, 5, 226);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4143, 5, 227);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4144, 5, 228);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4145, 5, 229);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4146, 5, 230);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4147, 5, 231);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4148, 5, 232);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4149, 5, 233);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4150, 5, 234);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4151, 5, 238);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4152, 5, 239);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4153, 5, 240);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4154, 5, 241);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4155, 5, 242);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4156, 5, 243);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4157, 9, 1);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4158, 9, 2);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4159, 9, 3);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4160, 9, 26);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4161, 9, 27);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4162, 9, 28);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4163, 9, 29);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4164, 9, 30);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4165, 9, 31);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4166, 9, 32);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4167, 9, 33);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4168, 9, 34);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4169, 9, 35);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4170, 9, 36);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4171, 9, 37);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4172, 9, 38);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4173, 9, 39);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4174, 9, 40);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4175, 9, 41);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4176, 9, 42);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4177, 9, 43);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4178, 9, 44);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4179, 9, 45);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4180, 9, 46);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4181, 9, 47);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4182, 9, 48);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4183, 9, 49);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4184, 9, 50);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4185, 9, 51);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4186, 9, 52);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4187, 9, 53);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4188, 9, 54);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4189, 9, 55);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4190, 9, 56);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4191, 9, 57);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4192, 9, 58);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4193, 9, 59);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4194, 9, 60);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4195, 9, 61);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4196, 9, 62);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4197, 9, 63);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4198, 9, 64);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4199, 9, 65);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4200, 9, 66);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4201, 9, 67);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4202, 9, 68);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4203, 9, 69);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4204, 9, 70);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4205, 9, 71);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4206, 9, 72);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4207, 9, 73);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4208, 9, 74);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4209, 9, 75);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4210, 9, 76);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4211, 9, 77);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4212, 9, 78);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4213, 9, 79);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4214, 9, 80);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4215, 9, 81);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4216, 9, 82);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4217, 9, 83);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4218, 9, 84);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4219, 9, 85);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4220, 9, 86);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4221, 9, 87);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4222, 9, 88);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4223, 9, 89);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4224, 9, 90);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4225, 9, 91);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4226, 9, 92);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4227, 9, 93);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4228, 9, 94);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4229, 9, 95);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4230, 9, 96);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4231, 9, 97);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4232, 9, 98);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4233, 9, 99);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4234, 9, 100);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4235, 9, 101);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4236, 9, 102);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4237, 9, 103);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4238, 9, 104);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4239, 9, 105);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4240, 9, 106);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4241, 9, 107);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4242, 9, 108);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4243, 9, 109);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4244, 9, 110);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4245, 9, 111);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4246, 9, 112);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4247, 9, 113);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4248, 9, 114);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4249, 9, 115);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4250, 9, 116);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4251, 9, 117);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4252, 9, 118);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4253, 9, 119);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4254, 9, 120);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4255, 9, 121);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4256, 9, 122);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4257, 9, 123);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4258, 9, 124);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4259, 9, 125);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4260, 9, 126);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4261, 9, 127);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4262, 9, 128);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4263, 9, 129);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4264, 9, 130);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4265, 9, 131);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4266, 9, 132);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4267, 9, 133);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4268, 9, 134);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4269, 9, 135);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4270, 9, 136);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4271, 9, 137);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4272, 9, 138);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4273, 9, 139);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4274, 9, 140);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4275, 9, 141);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4276, 9, 142);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4277, 9, 143);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4278, 9, 144);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4279, 9, 145);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4280, 9, 146);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4281, 9, 147);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4282, 9, 148);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4283, 9, 149);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4284, 9, 150);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4285, 9, 151);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4286, 9, 152);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4287, 9, 153);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4288, 9, 154);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4289, 9, 155);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4290, 9, 156);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4291, 9, 157);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4292, 9, 158);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4293, 9, 159);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4294, 9, 160);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4295, 9, 161);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4296, 9, 162);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4297, 9, 163);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4298, 9, 164);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4299, 9, 165);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4300, 9, 166);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4301, 9, 167);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4302, 9, 168);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4303, 9, 169);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4304, 9, 170);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4305, 9, 171);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4306, 9, 172);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4307, 9, 173);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4308, 9, 174);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4309, 9, 175);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4310, 9, 176);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4311, 9, 177);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4312, 9, 178);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4313, 9, 179);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4314, 9, 180);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4315, 9, 181);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4316, 9, 182);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4317, 9, 183);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4318, 9, 184);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4319, 9, 185);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4320, 9, 186);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4321, 9, 187);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4322, 9, 188);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4323, 9, 189);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4324, 9, 190);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4325, 9, 191);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4326, 9, 192);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4327, 9, 193);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4328, 9, 194);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4329, 9, 195);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4330, 9, 196);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4331, 9, 197);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4332, 9, 198);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4333, 9, 199);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4334, 9, 200);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4335, 9, 201);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4336, 9, 202);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4337, 9, 203);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4338, 9, 204);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4339, 9, 205);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4340, 9, 206);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4341, 9, 207);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4342, 9, 208);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4343, 9, 209);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4344, 9, 210);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4345, 9, 211);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4346, 9, 212);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4347, 9, 213);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4348, 9, 214);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4349, 9, 215);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4350, 9, 216);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4351, 9, 217);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4352, 9, 218);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4353, 9, 219);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4354, 9, 220);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4355, 9, 221);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4356, 9, 222);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4357, 9, 223);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4358, 9, 224);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4359, 9, 225);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4360, 9, 226);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4361, 9, 227);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4362, 9, 228);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4363, 9, 229);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4364, 9, 230);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4365, 9, 231);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4366, 9, 232);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4367, 9, 233);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4368, 9, 234);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4369, 9, 238);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4370, 9, 239);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4371, 9, 240);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4372, 9, 241);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4373, 9, 242);
INSERT INTO "public"."role_menu_20250811" ("id", "role_id", "menu_id") VALUES (4374, 9, 243);
COMMIT;

-- ----------------------------
-- Table structure for user_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_20250811";
CREATE TABLE "public"."user_20250811" (
  "id" int8,
  "name" varchar(254) COLLATE "pg_catalog"."default",
  "nickname" varchar(254) COLLATE "pg_catalog"."default",
  "gender" varchar(254) COLLATE "pg_catalog"."default",
  "avatar" varchar(254) COLLATE "pg_catalog"."default",
  "email" varchar(254) COLLATE "pg_catalog"."default",
  "phone" varchar(254) COLLATE "pg_catalog"."default",
  "password" varchar(254) COLLATE "pg_catalog"."default",
  "status" int4,
  "setting" text COLLATE "pg_catalog"."default",
  "token" varchar(254) COLLATE "pg_catalog"."default",
  "extra" text COLLATE "pg_catalog"."default",
  "logged_at" int8,
  "expire_at" int8,
  "created_at" int8,
  "updated_at" int8
)
;
ALTER TABLE "public"."user_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of user_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."user_20250811" ("id", "name", "nickname", "gender", "avatar", "email", "phone", "password", "status", "setting", "token", "extra", "logged_at", "expire_at", "created_at", "updated_at") VALUES (1, '超级管理员', '超级管理员', 'F', 'a9f224627346905e258d771e4043f921', '1280291001@qq.com', '18888888888', '$2a$10$9qRJe9KQo6sEcU8ipKg.e.dkl2E7Wy64SigYlgraTAn.1paHFq6W.', 1, '{"theme":"light","themeColor":"#165DFF","skin":"default","tabBar":true,"menuWidth":200,"layout":"default","language":"zh_CN","animation":"gp-slide-down","popupType":"modal"}', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWRzIjpudWxsLCJleHAiOjE3NTMyOTI5NTcsImlhdCI6MTc1MzI4NTc1Niwiam9iSWRzIjpudWxsLCJuYmYiOjE3NTMyODU3NTYsInJvbGVJZHMiOlsxXSwidXNlcklkIjoxLCJ1c2VyTmFtZSI6Iui2hee6p-euoeeQhuWRmCJ9.U1-L-8QhjZx7YW6tXPkIcXE5vH3m', NULL, 1753285756, 1753292956, 1713706137, 1753285756);
INSERT INTO "public"."user_20250811" ("id", "name", "nickname", "gender", "avatar", "email", "phone", "password", "status", "setting", "token", "extra", "logged_at", "expire_at", "created_at", "updated_at") VALUES (4, '1', '1', 'M', NULL, '31@qq.com', '18286219255', '$2a$10$qZOMEg/DMkcxGdJo9FeWNunWWCaSiix/Hz0Xa4FiVLqxdo5WY2l8i', 1, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXBhcnRtZW50SWRzIjpbNV0sImV4cCI6MTc1MTY0NTU5MCwiaWF0IjoxNzUxNjM4Mzg5LCJqb2JJZHMiOm51bGwsIm5iZiI6MTc1MTYzODM4OSwicm9sZUlkcyI6WzldLCJ1c2VySWQiOjQsInVzZXJOYW1lIjoiMSJ9.W-ZrtSpvDGBA9JjwtivLYylaNU3si-6zOfz8YJXLT9w', NULL, 1751638389, 1751645589, 1721840505, 1751638389);
COMMIT;

-- ----------------------------
-- Table structure for user_app_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_app_20250811";
CREATE TABLE "public"."user_app_20250811" (
  "id" int8,
  "user_id" int8,
  "app_id" int8,
  "status" int4,
  "disable_des" varchar(254) COLLATE "pg_catalog"."default",
  "logged_at" int8,
  "expired_at" int8,
  "created_at" int8
)
;
ALTER TABLE "public"."user_app_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of user_app_20250811
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_department_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_department_20250811";
CREATE TABLE "public"."user_department_20250811" (
  "id" int8,
  "department_" int8,
  "user_id" int8
)
;
ALTER TABLE "public"."user_department_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of user_department_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."user_department_20250811" ("id", "department_", "user_id") VALUES (15, 5, 4);
COMMIT;

-- ----------------------------
-- Table structure for user_job_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_job_20250811";
CREATE TABLE "public"."user_job_20250811" (
  "id" int8,
  "job_id" int8,
  "user_id" int8
)
;
ALTER TABLE "public"."user_job_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of user_job_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."user_job_20250811" ("id", "job_id", "user_id") VALUES (1, 1, 1);
INSERT INTO "public"."user_job_20250811" ("id", "job_id", "user_id") VALUES (25, 1, 4);
COMMIT;

-- ----------------------------
-- Table structure for user_role_20250811
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_role_20250811";
CREATE TABLE "public"."user_role_20250811" (
  "id" int8,
  "role_id" int8,
  "user_id" int8
)
;
ALTER TABLE "public"."user_role_20250811" OWNER TO "postgres";

-- ----------------------------
-- Records of user_role_20250811
-- ----------------------------
BEGIN;
INSERT INTO "public"."user_role_20250811" ("id", "role_id", "user_id") VALUES (1, 1, 1);
COMMIT;
