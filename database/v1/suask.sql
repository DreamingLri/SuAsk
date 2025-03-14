/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80403
 Source Host           : localhost:3306
 Source Schema         : suask

 Target Server Type    : MySQL
 Target Server Version : 80403
 File Encoding         : 65001

 Date: 27/12/2024 12:13:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for answers
-- ----------------------------
DROP TABLE IF EXISTS `answers`;
CREATE TABLE `answers`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '回答ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `question_id` int NOT NULL COMMENT '问题ID',
  `in_reply_to` int NULL DEFAULT NULL COMMENT '回复的回答ID，可为空',
  `contents` text CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '回答内容',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `upvotes` int NOT NULL DEFAULT 0 COMMENT '点赞量',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `question_id`(`question_id` ASC) USING BTREE,
  INDEX `in_reply_to`(`in_reply_to` ASC) USING BTREE,
  CONSTRAINT `answers_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `answers_ibfk_2` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `answers_ibfk_3` FOREIGN KEY (`in_reply_to`) REFERENCES `answers` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for attachments
-- ----------------------------
DROP TABLE IF EXISTS `attachments`;
CREATE TABLE `attachments`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '附件ID',
  `question_id` int NULL DEFAULT NULL COMMENT '问题ID',
  `answer_id` int NULL DEFAULT NULL COMMENT '回答ID',
  `type` enum('picture') CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '附件类型（目前仅支持图片）',
  `file_id` int NOT NULL COMMENT '文件ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `question_id`(`question_id` ASC) USING BTREE,
  INDEX `answer_id`(`answer_id` ASC) USING BTREE,
  INDEX `file_id`(`file_id` ASC) USING BTREE,
  CONSTRAINT `attachments_ibfk_1` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `attachments_ibfk_2` FOREIGN KEY (`answer_id`) REFERENCES `answers` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `attachments_ibfk_3` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `attachments_chk_1` CHECK (((`question_id` is not null) + (`answer_id` is not null)) = 1)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config`  (
  `id` bit(1) NOT NULL DEFAULT b'0' COMMENT '配置ID，限制为0',
  `default_avatar_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '默认头像文件路径',
  `default_theme_id` int NOT NULL COMMENT '默认主题ID',
  PRIMARY KEY (`id`) USING BTREE,
  CONSTRAINT `config_chk_1` CHECK (`id` = 0)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '收藏（置顶）ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `question_id` int NOT NULL COMMENT '问题ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `package` varchar(50) COLLATE utf8mb4_zh_0900_as_cs NOT NULL DEFAULT '默认收藏夹' COMMENT '收藏夹',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`user_id` ASC, `question_id` ASC) USING BTREE COMMENT '每个用户收藏同个问题最多一次',
  INDEX `question_id`(`question_id` ASC) USING BTREE,
  CONSTRAINT `favorites_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `favorites_ibfk_2` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '文件ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '文件名，不得包含非法字符例如斜杠',
  `hash` binary(32) NOT NULL COMMENT '文件哈希，算法暂定为BLAKE2b',
  `uploader_id` int NULL DEFAULT NULL COMMENT '上传者用户ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '文件上传时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uploader_id`(`uploader_id` ASC) USING BTREE,
  CONSTRAINT `files_ibfk_1` FOREIGN KEY (`uploader_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for notifications
-- ----------------------------
DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '提醒ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `question_id` int NOT NULL COMMENT '问题ID',
  `reply_to_id` int COMMENT '回复问题的ID',
  `answer_id` int COMMENT '问题ID',
  `type` enum('new_question','new_reply','new_answer') CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '提醒类型（新提问、新回复、新回答）',
  `is_read` bit(1) NOT NULL COMMENT '是否已读' DEFAULT 0,
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `question_id`(`question_id` ASC) USING BTREE,
  CONSTRAINT `notifications_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `notifications_ibfk_2` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for questions
-- ----------------------------
DROP TABLE IF EXISTS `questions`;
CREATE TABLE `questions`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '问题ID',
  `src_user_id` int NOT NULL COMMENT '发起提问的用户ID',
  `dst_user_id` int NULL DEFAULT NULL COMMENT '被提问的用户ID，为空时问大家，不为空时问教师',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '问题标题',
  `contents` text CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '问题内容',
  `is_private` bit(1) NOT NULL COMMENT '是否私密提问，仅在问教师时可为是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `views` int NOT NULL DEFAULT 0 COMMENT '浏览量',
  `reply_cnt` int NOT NULL DEFAULT 0 COMMENT '回复数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `src_user_id`(`src_user_id` ASC) USING BTREE,
  INDEX `dst_user_id`(`dst_user_id` ASC) USING BTREE,
  FULLTEXT INDEX `title`(`title`) WITH PARSER `ngram` COMMENT '内容支持全文搜索，使用ngram parser以支持中文，默认token size为2',
  CONSTRAINT `questions_ibfk_1` FOREIGN KEY (`src_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `questions_ibfk_2` FOREIGN KEY (`dst_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `questions_chk_1` CHECK ((`dst_user_id` is not null) or (`is_private` = 0))
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for teachers
-- ----------------------------
DROP TABLE IF EXISTS `teachers`;
CREATE TABLE `teachers`  (
  `id` int NOT NULL,
  `responses` int NULL DEFAULT NULL COMMENT '回复数',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '老师名字',
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '老师头像链接',
  `introduction` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '老师简介',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '老师邮箱',
  `perm` enum('public','protected','private') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '提问箱权限',
  PRIMARY KEY (`id`) USING BTREE,
  CONSTRAINT `teachers_ibfk_1` FOREIGN KEY (`id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for upvotes
-- ----------------------------
DROP TABLE IF EXISTS `upvotes`;
CREATE TABLE `upvotes`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '点赞ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `question_id` int NULL DEFAULT NULL COMMENT '问题ID',
  `answer_id` int NULL DEFAULT NULL COMMENT '回复ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `upvotes_ibfk_1`(`user_id` ASC) USING BTREE,
  INDEX `upvotes_ibfk_2`(`question_id` ASC) USING BTREE,
  INDEX `upvotes_ibfk_3`(`answer_id` ASC) USING BTREE,
  CONSTRAINT `upvotes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `upvotes_ibfk_2` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `upvotes_ibfk_3` FOREIGN KEY (`answer_id`) REFERENCES `answers` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `upvotes_chk_1` CHECK (((`question_id` is not null) + (`answer_id` is not null)) = 1)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for user_relation
-- ----------------------------
DROP TABLE IF EXISTS `user_relation`;
CREATE TABLE `user_relation`  (
  `question_id` int NOT NULL,
  `user_id` int NOT NULL,
  INDEX `question_id`(`question_id` ASC) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `user_relation_ibfk_1` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_relation_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '记录每个问题的前n个回复' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '用户名',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '邮箱',
  `salt` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '加密盐',
  `password_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '密码哈希，算法暂定为Argon2id',
  `role` enum('admin','teacher','student') CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '角色',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '昵称',
  `introduction` text CHARACTER SET utf8mb4 COLLATE utf8mb4_zh_0900_as_cs NOT NULL COMMENT '简介',
  `avatar_file_id` int NULL DEFAULT NULL COMMENT '头像文件ID，为空时为配置的默认头像',
  `theme_id` int NULL DEFAULT NULL COMMENT '主题ID，为空时为配置的默认主题',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE COMMENT '用户名唯一',
  UNIQUE INDEX `email`(`email` ASC) USING BTREE COMMENT '邮箱唯一',
  INDEX `avatar_file_id`(`avatar_file_id` ASC) USING BTREE,
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`avatar_file_id`) REFERENCES `files` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for settings
-- ----------------------------

DROP TABLE IF EXISTS `settings`;
CREATE TABLE `settings` (
    `id` int NOT NULL COMMENT '设置id',
    `theme_id` int NULL DEFAULT NULL COMMENT '主题ID，为空时为配置的默认主题',
    `question_box_perm` enum('public','protected','private') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '提问箱权限',
    PRIMARY KEY (`id`) USING BTREE,
    CONSTRAINT `settings_ibfk_1` FOREIGN KEY (`id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for viewed
-- ----------------------------

DROP TABLE IF EXISTS `viewed`;
CREATE TABLE `viewed` (
    `user_id` int NOT NULL COMMENT '用户id',
    `question_id` int NOT NULL COMMENT '问题id',
    PRIMARY KEY (`user_id`, `question_id`) USING BTREE,
    CONSTRAINT `viewed_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT `viewed_ibfk_2` FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    INDEX `viewed_id` (`user_id` ASC , `question_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_zh_0900_as_cs ROW_FORMAT = DYNAMIC;