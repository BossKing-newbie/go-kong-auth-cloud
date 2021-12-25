/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : localhost:3306
 Source Schema         : generate_db

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 25/12/2021 11:56:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `login_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户登录名',
  `user_password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户密码',
  `insert_time` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `is_deleted` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'N',
  `role_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '数据权限ID',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `sys_user_login_name_uindex`(`login_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (5, 'bossking', '$2a$10$LMu4q1mgZwMGZQaI/J0ALenxrZSUobpPGCrYogHm4wxpKhiXoGop6', '2021-12-20 06:49:36', '2021-12-20 06:49:36', 'N', 'user');
INSERT INTO `sys_user` VALUES (6, 'admin', '$2a$10$9s82AKikfPYIgYtRnEHu3.BFtrdY41wzAXmqlDqz0mdsZFwAYDRMe', '2021-12-21 02:39:25', '2021-12-21 02:39:25', 'N', 'admin');

SET FOREIGN_KEY_CHECKS = 1;
