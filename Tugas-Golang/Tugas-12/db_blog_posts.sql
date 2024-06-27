/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 100432 (10.4.32-MariaDB)
 Source Host           : localhost:3306
 Source Schema         : db_blog_posts

 Target Server Type    : MySQL
 Target Server Version : 100432 (10.4.32-MariaDB)
 File Encoding         : 65001

 Date: 28/06/2024 01:26:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_posts
-- ----------------------------
DROP TABLE IF EXISTS `blog_posts`;
CREATE TABLE `blog_posts`  (
  `id` int NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `author_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `author_id`(`author_id` ASC) USING BTREE,
  CONSTRAINT `blog_posts_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_posts
-- ----------------------------
INSERT INTO `blog_posts` VALUES (1, 'Algorithm', 'this is article for algorithm', 1);
INSERT INTO `blog_posts` VALUES (2, 'Lorem Ipsum', 'lorem ipsum dolor sit amet', 1);
INSERT INTO `blog_posts` VALUES (3, 'NodeJS', 'Node.js is a JavaScript runtime built on Chrome\'s V8 JavaScript engine.', 2);
INSERT INTO `blog_posts` VALUES (4, 'ReactJS', 'ReactJS is JavaScript library for building user interfaces', 4);
INSERT INTO `blog_posts` VALUES (5, 'Git', 'Git is a free and open source distributed version control system designed to handle everything from small to very large projects with speed and efficiency', 3);
INSERT INTO `blog_posts` VALUES (6, 'Basketball', 'Basketball is a team sport in which two teams, most commonly of five players each, opposing one another on a rectangular court, compete with the primary objective of shooting a basketball (approximately 9.4 inches (24 cm) in diameter) through the defender\'s hoop (a basket 18 inches (46 cm) in diameter mounted 10 feet (3.048 m) high to a backboard at each end of the court, while preventing the opposing team from shooting through their own hoop', 2);

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories`  (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES (1, 'Basic');
INSERT INTO `categories` VALUES (2, 'Library');
INSERT INTO `categories` VALUES (3, 'Unknown');
INSERT INTO `categories` VALUES (4, 'Sport');

-- ----------------------------
-- Table structure for post_categories
-- ----------------------------
DROP TABLE IF EXISTS `post_categories`;
CREATE TABLE `post_categories`  (
  `post_id` int NOT NULL,
  `category_id` int NOT NULL,
  PRIMARY KEY (`post_id`, `category_id`) USING BTREE,
  INDEX `category_id`(`category_id` ASC) USING BTREE,
  CONSTRAINT `post_categories_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `blog_posts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `post_categories_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post_categories
-- ----------------------------
INSERT INTO `post_categories` VALUES (1, 1);
INSERT INTO `post_categories` VALUES (2, 3);
INSERT INTO `post_categories` VALUES (3, 1);
INSERT INTO `post_categories` VALUES (3, 2);
INSERT INTO `post_categories` VALUES (4, 1);
INSERT INTO `post_categories` VALUES (4, 2);
INSERT INTO `post_categories` VALUES (5, 1);
INSERT INTO `post_categories` VALUES (6, 4);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'John Doe');
INSERT INTO `users` VALUES (2, 'Robert');
INSERT INTO `users` VALUES (3, 'Frank');
INSERT INTO `users` VALUES (4, 'Jane');

SET FOREIGN_KEY_CHECKS = 1;
