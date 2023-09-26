/*
 Navicat Premium Data Transfer

 Source Server         : LOCAL
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 127.0.0.1:3306
 Source Schema         : xzs

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 10/03/2022 11:01:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_exam_paper
-- ----------------------------
DROP TABLE IF EXISTS `t_exam_paper`;
CREATE TABLE `t_exam_paper`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `subject_id` int(11) NULL DEFAULT NULL,
  `paper_type` int(11) NULL DEFAULT NULL,
  `grade_level` int(11) NULL DEFAULT NULL,
  `score` int(11) NULL DEFAULT NULL,
  `question_count` int(11) NULL DEFAULT NULL,
  `suggest_time` int(11) NULL DEFAULT NULL,
  `limit_start_time` datetime(0) NULL DEFAULT NULL,
  `limit_end_time` datetime(0) NULL DEFAULT NULL,
  `frame_text_content_id` int(11) NULL DEFAULT NULL,
  `create_user` int(11) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `task_exam_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_exam_paper
-- ----------------------------
INSERT INTO `t_exam_paper` VALUES (1, '测试', 1, 1, 1, 10, 1, 10, NULL, NULL, 2, 2, '2022-03-09 13:49:03', 0, NULL);

-- ----------------------------
-- Table structure for t_exam_paper_answer
-- ----------------------------
DROP TABLE IF EXISTS `t_exam_paper_answer`;
CREATE TABLE `t_exam_paper_answer`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `exam_paper_id` int(11) NULL DEFAULT NULL,
  `paper_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `paper_type` int(11) NULL DEFAULT NULL,
  `subject_id` int(11) NULL DEFAULT NULL,
  `system_score` int(11) NULL DEFAULT NULL,
  `user_score` int(11) NULL DEFAULT NULL,
  `paper_score` int(11) NULL DEFAULT NULL,
  `question_correct` int(11) NULL DEFAULT NULL,
  `question_count` int(11) NULL DEFAULT NULL,
  `do_time` int(11) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `create_user` int(11) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `task_exam_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_exam_paper_answer
-- ----------------------------
INSERT INTO `t_exam_paper_answer` VALUES (1, 1, '测试', 1, 1, 10, 10, 10, 1, 1, 4, 2, 1, '2022-03-09 13:50:04', NULL);

-- ----------------------------
-- Table structure for t_exam_paper_question_customer_answer
-- ----------------------------
DROP TABLE IF EXISTS `t_exam_paper_question_customer_answer`;
CREATE TABLE `t_exam_paper_question_customer_answer`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `question_id` int(11) NULL DEFAULT NULL,
  `exam_paper_id` int(11) NULL DEFAULT NULL,
  `exam_paper_answer_id` int(11) NULL DEFAULT NULL,
  `question_type` int(11) NULL DEFAULT NULL,
  `subject_id` int(11) NULL DEFAULT NULL,
  `customer_score` int(11) NULL DEFAULT NULL,
  `question_score` int(11) NULL DEFAULT NULL,
  `question_text_content_id` int(11) NULL DEFAULT NULL,
  `answer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `text_content_id` int(11) NULL DEFAULT NULL,
  `do_right` bit(1) NULL DEFAULT NULL,
  `create_user` int(11) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `item_order` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_exam_paper_question_customer_answer
-- ----------------------------
INSERT INTO `t_exam_paper_question_customer_answer` VALUES (1, 1, 1, 1, 1, 1, 10, 10, 1, 'B', NULL, b'1', 1, '2022-03-09 13:50:04', 1);

-- ----------------------------
-- Table structure for t_message
-- ----------------------------
DROP TABLE IF EXISTS `t_message`;
CREATE TABLE `t_message`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `send_user_id` int(11) NULL DEFAULT NULL,
  `send_user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `send_real_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `receive_user_count` int(11) NULL DEFAULT NULL,
  `read_count` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_message_user
-- ----------------------------
DROP TABLE IF EXISTS `t_message_user`;
CREATE TABLE `t_message_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message_id` int(11) NULL DEFAULT NULL,
  `receive_user_id` int(11) NULL DEFAULT NULL,
  `receive_user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `receive_real_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `readed` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `read_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_question
-- ----------------------------
DROP TABLE IF EXISTS `t_question`;
CREATE TABLE `t_question`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `question_type` int(11) NULL DEFAULT NULL,
  `subject_id` int(11) NULL DEFAULT NULL,
  `score` int(11) NULL DEFAULT NULL,
  `grade_level` int(11) NULL DEFAULT NULL,
  `difficult` int(11) NULL DEFAULT NULL,
  `correct` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `info_text_content_id` int(11) NULL DEFAULT NULL,
  `create_user` int(11) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_question
-- ----------------------------
INSERT INTO `t_question` VALUES (1, 1, 1, 10, 1, 1, 'B', 1, 2, 1, '2022-03-09 13:48:19', 0);

-- ----------------------------
-- Table structure for t_subject
-- ----------------------------
DROP TABLE IF EXISTS `t_subject`;
CREATE TABLE `t_subject`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `level` int(11) NULL DEFAULT NULL,
  `level_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `item_order` int(11) NULL DEFAULT NULL,
  `deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_subject
-- ----------------------------
INSERT INTO `t_subject` VALUES (1, '数学', 1, '一年级', NULL, 0);
INSERT INTO `t_subject` VALUES (2, '英语', 1, '一年级', NULL, 0);

-- ----------------------------
-- Table structure for t_task_exam
-- ----------------------------
DROP TABLE IF EXISTS `t_task_exam`;
CREATE TABLE `t_task_exam`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `grade_level` int(11) NULL DEFAULT NULL,
  `frame_text_content_id` int(11) NULL DEFAULT NULL,
  `create_user` int(11) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `create_user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_task_exam_customer_answer
-- ----------------------------
DROP TABLE IF EXISTS `t_task_exam_customer_answer`;
CREATE TABLE `t_task_exam_customer_answer`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_exam_id` int(11) NULL DEFAULT NULL,
  `create_user` int(255) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `text_content_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for t_text_content
-- ----------------------------
DROP TABLE IF EXISTS `t_text_content`;
CREATE TABLE `t_text_content`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_text_content
-- ----------------------------
INSERT INTO `t_text_content` VALUES (1, '{\"titleContent\":\"1＋１＝\",\"analyze\":\"２\",\"questionItemObjects\":[{\"prefix\":\"A\",\"content\":\"１\",\"score\":null,\"itemUuid\":null},{\"prefix\":\"B\",\"content\":\"２\",\"score\":null,\"itemUuid\":null},{\"prefix\":\"C\",\"content\":\"３\",\"score\":null,\"itemUuid\":null},{\"prefix\":\"D\",\"content\":\"４\",\"score\":null,\"itemUuid\":null}],\"correct\":\"B\"}', '2022-03-09 13:48:19');
INSERT INTO `t_text_content` VALUES (2, '[{\"name\":\"计算\",\"questionItems\":[{\"id\":1,\"itemOrder\":1}]}]', '2022-03-09 13:49:03');

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `real_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `age` int(11) NULL DEFAULT NULL,
  `sex` int(11) NULL DEFAULT NULL,
  `birth_day` datetime(0) NULL DEFAULT NULL,
  `user_level` int(11) NULL DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role` int(11) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `image_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `modify_time` datetime(0) NULL DEFAULT NULL,
  `last_active_time` datetime(0) NULL DEFAULT NULL,
  `deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0,
  `wx_open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES (1, 'd2d29da2-dcb3-4013-b874-727626236f47', 'student', 'D1AGFL+Gx37t0NPG4d6biYP5Z31cNbwhK5w1lUeiHB2zagqbk8efYfSjYoh1Z/j1dkiRjHU+b0EpwzCh8IGsksJjzD65ci5LsnodQVf4Uj6D3pwoscXGqmkjjpzvSJbx42swwNTA+QoDU8YLo7JhtbUK2X0qCjFGpd+8eJ5BGvk=', '学生', 18, 1, '2019-09-01 16:00:00', 1, '19171171610', 1, 1, 'https://www.mindskip.net:9008/image/ba607a75-83ba-4530-8e23-660b72dc4953/头像.jpg', '2019-09-07 18:55:02', '2020-02-04 08:26:54', NULL, 0, NULL);
INSERT INTO `t_user` VALUES (2, '52045f5f-a13f-4ccc-93dd-f7ee8270ad4c', 'admin', 'D1AGFL+Gx37t0NPG4d6biYP5Z31cNbwhK5w1lUeiHB2zagqbk8efYfSjYoh1Z/j1dkiRjHU+b0EpwzCh8IGsksJjzD65ci5LsnodQVf4Uj6D3pwoscXGqmkjjpzvSJbx42swwNTA+QoDU8YLo7JhtbUK2X0qCjFGpd+8eJ5BGvk=', '管理员', 30, 1, '2019-09-07 18:56:07', NULL, NULL, 3, 1, NULL, '2019-09-07 18:56:21', NULL, NULL, 0, NULL);
INSERT INTO `t_user` VALUES (3, 'd2d29da2-dcb3-4013-b874-727626236f47', 'student1', 'D1AGFL+Gx37t0NPG4d6biYP5Z31cNbwhK5w1lUeiHB2zagqbk8efYfSjYoh1Z/j1dkiRjHU+b0EpwzCh8IGsksJjzD65ci5LsnodQVf4Uj6D3pwoscXGqmkjjpzvSJbx42swwNTA+QoDU8YLo7JhtbUK2X0qCjFGpd+8eJ5BGvk=', '学生', 18, 1, '2019-09-01 16:00:00', 1, '19171171610', 1, 1, 'https://www.mindskip.net:9008/image/ba607a75-83ba-4530-8e23-660b72dc4953/头像.jpg', '2019-09-07 18:55:02', '2020-02-04 08:26:54', NULL, 0, NULL);
INSERT INTO `t_user` VALUES (4, '52045f5f-a13f-4ccc-93dd-f7ee8270ad4c', 'admin1', 'D1AGFL+Gx37t0NPG4d6biYP5Z31cNbwhK5w1lUeiHB2zagqbk8efYfSjYoh1Z/j1dkiRjHU+b0EpwzCh8IGsksJjzD65ci5LsnodQVf4Uj6D3pwoscXGqmkjjpzvSJbx42swwNTA+QoDU8YLo7JhtbUK2X0qCjFGpd+8eJ5BGvk=', '管理员', 29, 1, '2019-09-07 18:56:07', NULL, NULL, 3, 1, NULL, '2019-10-07 18:56:21', NULL, NULL, 0, NULL);
-- ----------------------------
-- Table structure for t_user_event_log
-- ----------------------------
DROP TABLE IF EXISTS `t_user_event_log`;
CREATE TABLE `t_user_event_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `real_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of t_user_event_log
-- ----------------------------
INSERT INTO `t_user_event_log` VALUES (1, 2, 'admin', '管理员', 'admin 登录了学之思开源考试系统', '2022-03-04 13:05:13');
INSERT INTO `t_user_event_log` VALUES (2, 2, 'admin', '管理员', 'admin 登出了学之思开源考试系统', '2022-03-04 13:18:32');
INSERT INTO `t_user_event_log` VALUES (3, 2, 'admin', '管理员', 'admin 登录了学之思开源考试系统', '2022-03-04 13:18:43');
INSERT INTO `t_user_event_log` VALUES (4, 2, 'admin', '管理员', 'admin 登录了学之思开源考试系统', '2022-03-08 18:03:51');
INSERT INTO `t_user_event_log` VALUES (5, 2, 'admin', '管理员', 'admin  登录了学之思开源考试系统', '2022-03-09 12:04:53');
INSERT INTO `t_user_event_log` VALUES (6, 2, 'admin', '管理员', 'admin 登录了学之思开源考试系统', '2022-03-09 13:45:12');
INSERT INTO `t_user_event_log` VALUES (7, 1, 'student', '学生', 'student 登录了学之思开源考试系统', '2022-03-09 13:49:25');
INSERT INTO `t_user_event_log` VALUES (8, 1, 'student', '学生', 'student 提交试卷：测试 得分：1 耗时：4 秒', '2022-03-09 13:50:04');
INSERT INTO `t_user_event_log` VALUES (9, 2, 'admin', '管理员', 'admin 登录了学之思开源考试系统', '2022-03-09 19:43:55');

-- ----------------------------
-- Table structure for t_user_token
-- ----------------------------
DROP TABLE IF EXISTS `t_user_token`;
CREATE TABLE `t_user_token`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `user_id` int(11) NULL DEFAULT NULL,
  `wx_open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `end_time` datetime(0) NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
