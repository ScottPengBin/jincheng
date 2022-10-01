/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : jin_cheng

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 01/10/2022 21:58:55
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for jc_car_info
-- ----------------------------
DROP TABLE IF EXISTS `jc_car_info`;
CREATE TABLE `jc_car_info`
(
    `id`          int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `color`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `car_no`      varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `note`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `created_at`  datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP (0),
    `updated_at`  datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP (0),
    `deleted_at`  datetime(0) NULL DEFAULT NULL,
    `delete_flag` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '逻辑删除 1删除 默认0',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for jc_member
-- ----------------------------
DROP TABLE IF EXISTS `jc_member`;
CREATE TABLE `jc_member`
(
    `id`          int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `gender`      varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '3' COMMENT '性别 ',
    `phone`       int(0) NOT NULL COMMENT '电话号码',
    `brith_day`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `car_id`      int(0) NOT NULL,
    `status`      tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1启用 2禁用',
    `wet_chat_id` int(0) NULL DEFAULT NULL COMMENT '微信',
    `note`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
    `created_at`  datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP (0) COMMENT '创建时间',
    `update_at`   datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP (0) COMMENT '更新时间',
    `delete_flag` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除 1删除 默认0',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `car_id`(`car_id`) USING BTREE,
    INDEX         `phone`(`phone`) USING BTREE,
    INDEX         `name`(`name`) USING BTREE,
    INDEX         `wet_chat_id`(`wet_chat_id`) USING BTREE
) ENGINE = InnoDB  CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET
FOREIGN_KEY_CHECKS = 1;
