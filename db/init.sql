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

 Date: 10/10/2022 11:57:05
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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_car_info
-- ----------------------------
INSERT INTO `jc_car_info`
VALUES (2, '2022款宝马530', '白色', '渝A88888', '这是车辆信息备注', '2022-10-01 22:00:21', '2022-10-01 22:00:21', NULL,
        0);

-- ----------------------------
-- Table structure for jc_member
-- ----------------------------
DROP TABLE IF EXISTS `jc_member`;
CREATE TABLE `jc_member`
(
    `id`          int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `gender`      varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '3' COMMENT '性别 ',
    `phone`       varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '电话号码',
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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_member
-- ----------------------------
INSERT INTO `jc_member`
VALUES (1, '张三', '男', '13568383850', '1992-06-30', 2, 0, 0, '这是会员信息备注', '2022-10-01 22:00:21',
        '2022-10-01 22:00:21', 0);

-- ----------------------------
-- Table structure for jc_menus
-- ----------------------------
DROP TABLE IF EXISTS `jc_menus`;
CREATE TABLE `jc_menus`
(
    `id`        int(0) NOT NULL,
    `parentId`  int(0) NULL DEFAULT NULL,
    `title`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `icon`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `basePath`  varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `path`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `target`    varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `sort`      int(0) NULL DEFAULT 0,
    `type`      int(0) NOT NULL DEFAULT 1,
    `enabled`   tinyint(1) NOT NULL,
    `code`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `name`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `entry`     varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `createdAt` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updatedAt` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP (0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `menus_id_uindex`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_menus
-- ----------------------------
INSERT INTO `jc_menus`
VALUES (1, NULL, '系统管理', NULL, NULL, NULL, 'menu', 0, 1, 1, NULL, NULL, NULL, '2022-10-10 11:53:56',
        '2022-10-10 11:53:56');
INSERT INTO `jc_menus`
VALUES (2, 1, '用户管理', NULL, NULL, '/users', 'menu', 0, 1, 1, NULL, NULL, NULL, '2022-10-10 11:53:56',
        '2022-10-10 11:53:56');
INSERT INTO `jc_menus`
VALUES (3, 1, '角色管理', NULL, NULL, '/roles', 'menu', 0, 1, 1, NULL, NULL, NULL, '2022-10-10 11:53:56',
        '2022-10-10 11:53:56');
INSERT INTO `jc_menus`
VALUES (4, 1, '菜单管理', NULL, NULL, '/menus', 'menu', 0, 1, 1, NULL, NULL, NULL, '2022-10-10 11:53:56',
        '2022-10-10 11:53:56');
INSERT INTO `jc_menus`
VALUES (5, 2, '添加用户', NULL, NULL, NULL, NULL, 0, 2, 1, 'ADD_USER', NULL, NULL, '2022-10-10 11:53:56',
        '2022-10-10 11:53:56');
INSERT INTO `jc_menus`
VALUES (6, 2, '删除用户', NULL, NULL, NULL, NULL, 0, 2, 1, 'UPDATE_USER', NULL, NULL, '2022-10-10 11:53:56',
        '2022-10-10 11:53:56');

-- ----------------------------
-- Table structure for jc_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `jc_role_menus`;
CREATE TABLE `jc_role_menus`
(
    `id`        int(0) NOT NULL,
    `roleId`    int(0) NOT NULL,
    `menuId`    int(0) NOT NULL,
    `createdAt` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updatedAt` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `role_menus_id_uindex`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_role_menus
-- ----------------------------
INSERT INTO `jc_role_menus`
VALUES (1, 1, 1, '2022-10-10 11:53:43', '2022-10-10 11:53:43');
INSERT INTO `jc_role_menus`
VALUES (2, 1, 2, '2022-10-10 11:53:43', '2022-10-10 11:53:43');
INSERT INTO `jc_role_menus`
VALUES (3, 1, 3, '2022-10-10 11:53:43', '2022-10-10 11:53:43');
INSERT INTO `jc_role_menus`
VALUES (4, 1, 4, '2022-10-10 11:53:43', '2022-10-10 11:53:43');
INSERT INTO `jc_role_menus`
VALUES (5, 1, 5, '2022-10-10 11:53:43', '2022-10-10 11:53:43');

-- ----------------------------
-- Table structure for jc_roles
-- ----------------------------
DROP TABLE IF EXISTS `jc_roles`;
CREATE TABLE `jc_roles`
(
    `id`        int(0) NOT NULL,
    `type`      int(0) NULL DEFAULT NULL,
    `systemId`  int(0) NULL DEFAULT NULL,
    `enabled`   tinyint(1) NOT NULL,
    `name`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `remark`    varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `createdAt` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updatedAt` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP (0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `roles_id_uindex`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_roles
-- ----------------------------
INSERT INTO `jc_roles`
VALUES (1, 1, NULL, 1, '超级管理员', '超级管理员拥有系统所有权限', '2022-10-10 11:53:37', '2022-10-10 11:53:37');

-- ----------------------------
-- Table structure for jc_user_collect_menus
-- ----------------------------
DROP TABLE IF EXISTS `jc_user_collect_menus`;
CREATE TABLE `jc_user_collect_menus`
(
    `id`        int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
    `userId`    int(0) NOT NULL,
    `menuId`    int(0) NOT NULL,
    `createdAt` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updatedAt` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_collect_menus_id_uindex`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_user_collect_menus
-- ----------------------------
INSERT INTO `jc_user_collect_menus`
VALUES (1, 1, 2, '2022-10-10 11:54:24', '2022-10-10 11:54:24');

-- ----------------------------
-- Table structure for jc_user_roles
-- ----------------------------
DROP TABLE IF EXISTS `jc_user_roles`;
CREATE TABLE `jc_user_roles`
(
    `id`        int(0) NOT NULL,
    `userId`    int(0) NOT NULL,
    `roleId`    int(0) NOT NULL,
    `createdAt` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updatedAt` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user_roles_id_uindex`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_user_roles
-- ----------------------------
INSERT INTO `jc_user_roles`
VALUES (1, 1, 1, '2022-10-10 11:53:51', '2022-10-10 11:53:51');

-- ----------------------------
-- Table structure for jc_users
-- ----------------------------
DROP TABLE IF EXISTS `jc_users`;
CREATE TABLE `jc_users`
(
    `id`        int(0) NOT NULL,
    `account`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `name`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `password`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `mobile`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `email`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `enabled`   tinyint(1) NOT NULL,
    `createdAt` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `updatedAt` timestamp(0)                                                 NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `users_account_uindex`(`account`) USING BTREE,
    UNIQUE INDEX `users_id_uindex`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jc_users
-- ----------------------------
INSERT INTO `jc_users`
VALUES (1, 'admin', '管理员', '7c4a8d09ca3762af61e59520943dc26494f8941b', '18888888888', 'email@qq.com', 1,
        '2022-10-10 11:53:48', '2022-10-10 11:53:48');

SET
FOREIGN_KEY_CHECKS = 1;
