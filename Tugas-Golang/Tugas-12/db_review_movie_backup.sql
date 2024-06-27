/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 100432 (10.4.32-MariaDB)
 Source Host           : localhost:3306
 Source Schema         : db_review_movie

 Target Server Type    : MySQL
 Target Server Version : 100432 (10.4.32-MariaDB)
 File Encoding         : 65001

 Date: 28/06/2024 01:09:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Cast
-- ----------------------------
DROP TABLE IF EXISTS `Cast`;
CREATE TABLE `Cast`  (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `age` int NULL DEFAULT NULL,
  `biodata` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of Cast
-- ----------------------------

-- ----------------------------
-- Table structure for Genre
-- ----------------------------
DROP TABLE IF EXISTS `Genre`;
CREATE TABLE `Genre`  (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of Genre
-- ----------------------------

-- ----------------------------
-- Table structure for Movie
-- ----------------------------
DROP TABLE IF EXISTS `Movie`;
CREATE TABLE `Movie`  (
  `id` int NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `year` int NULL DEFAULT NULL,
  `poster` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of Movie
-- ----------------------------

-- ----------------------------
-- Table structure for MovieGenre
-- ----------------------------
DROP TABLE IF EXISTS `MovieGenre`;
CREATE TABLE `MovieGenre`  (
  `movie_id` int NOT NULL,
  `genre_id` int NOT NULL,
  PRIMARY KEY (`movie_id`, `genre_id`) USING BTREE,
  INDEX `genre_id`(`genre_id` ASC) USING BTREE,
  CONSTRAINT `moviegenre_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `Movie` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `moviegenre_ibfk_2` FOREIGN KEY (`genre_id`) REFERENCES `Genre` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of MovieGenre
-- ----------------------------

-- ----------------------------
-- Table structure for Profile
-- ----------------------------
DROP TABLE IF EXISTS `Profile`;
CREATE TABLE `Profile`  (
  `id` int NOT NULL,
  `age` int NULL DEFAULT NULL,
  `biodata` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `user_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `profile_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of Profile
-- ----------------------------

-- ----------------------------
-- Table structure for Review
-- ----------------------------
DROP TABLE IF EXISTS `Review`;
CREATE TABLE `Review`  (
  `id` int NOT NULL,
  `rating` int NULL DEFAULT NULL,
  `critics` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `user_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `review_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `review_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `Movie` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of Review
-- ----------------------------

-- ----------------------------
-- Table structure for User
-- ----------------------------
DROP TABLE IF EXISTS `User`;
CREATE TABLE `User`  (
  `id` int NOT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of User
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
