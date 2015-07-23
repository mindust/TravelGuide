/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50544
Source Host           : localhost:3306
Source Database       : kanghui

Target Server Type    : MYSQL
Target Server Version : 50544
File Encoding         : 65001

Date: 2015-07-23 08:12:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for ph_fee_record
-- ----------------------------
DROP TABLE IF EXISTS `ph_fee_record`;
CREATE TABLE `ph_fee_record` (
`ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`USER_ID`  int(36) NULL DEFAULT NULL ,
`TRAVEL_PACKAGES_ID`  int(36) NULL DEFAULT NULL ,
`VIEW_SPOT_ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`FEE`  varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`FEE_DATE`  date NULL DEFAULT NULL ,
`EQUIP_ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`ID`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of ph_fee_record
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ph_travel_package_view_spots
-- ----------------------------
DROP TABLE IF EXISTS `ph_travel_package_view_spots`;
CREATE TABLE `ph_travel_package_view_spots` (
`ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`TRAVEL_PACKAGES_ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`VIEW_SPOT_ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`DESCRIPTION`  varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`DAYS`  int(2) NULL DEFAULT NULL ,
`PLAY_ITEMS`  varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`ID`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of ph_travel_package_view_spots
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ph_travel_packages
-- ----------------------------
DROP TABLE IF EXISTS `ph_travel_packages`;
CREATE TABLE `ph_travel_packages` (
`ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`NAME`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`DESCRIPTION`  varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`FEE`  varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`START_DATE`  date NULL DEFAULT NULL ,
`END_DATE`  date NULL DEFAULT NULL ,
`DAYS`  int(2) NULL DEFAULT NULL ,
`HOTELS`  varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`TRANSPOT`  varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`LEVEL`  int(1) NULL DEFAULT NULL ,
`PERSON_NUM`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`LABLE`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`ID`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of ph_travel_packages
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ph_users
-- ----------------------------
DROP TABLE IF EXISTS `ph_users`;
CREATE TABLE `ph_users` (
`ID`  int(36) NOT NULL AUTO_INCREMENT ,
`USER_NAME`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`PASSWORD`  varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`DESCRIPTION`  varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`QQ_NO`  varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`MOBILE`  varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`EMAIL`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`LEVEL`  varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`LOGO_PATH`  varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`ID`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=2

;

-- ----------------------------
-- Records of ph_users
-- ----------------------------
BEGIN;
INSERT INTO `ph_users` VALUES ('1', 'Messi', '123JKLjkl', 'Messi', '52486846', '18013579965', 'messi.he@sig.biz', '1', '2');
COMMIT;

-- ----------------------------
-- Table structure for ph_view_spot_images
-- ----------------------------
DROP TABLE IF EXISTS `ph_view_spot_images`;
CREATE TABLE `ph_view_spot_images` (
`ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL ,
`VIEW_SPOT_ID`  varchar(36) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`PATH`  varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`NAME`  varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`FORMAT`  varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`DESCRIPTION`  varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`ID`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci

;

-- ----------------------------
-- Records of ph_view_spot_images
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for ph_view_spots
-- ----------------------------
DROP TABLE IF EXISTS `ph_view_spots`;
CREATE TABLE `ph_view_spots` (
`ID`  int(100) NOT NULL AUTO_INCREMENT ,
`COUNTRY`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`PROVINCE`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`CITY`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`COUNTY`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`NAME`  varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`LEVEL`  varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`GRADE`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`HIGHLIGHTS1`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`HIGHLIGHTS2`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`HIGHLIGHTS3`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`LABEL`  varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`PRICE`  varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`STATUS`  varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`ID`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=10

;

-- ----------------------------
-- Records of ph_view_spots
-- ----------------------------
BEGIN;
INSERT INTO `ph_view_spots` VALUES ('1', 'China ', 'JiangSu', 'Suzhou', '1', 'TaiLake', 'First', 'Hot', '1', '2', '3', 'nutual view', '2000', 'enable'), ('2', 'Japan', 'Fuga', 'Fuga', 'Fuga', 'Changqi', 'First', 'Luxury', '1', '2', '3', 'Fuga ChangQi', '11111', 'enable'), ('3', 'Japan', 'Tokyo', 'Tokyo', 'Tokyo', 'Name', 'Second', 'Luxury', '1', '2', '3', 'adsf', '121', ''), ('4', 'Japan', 'Tokyo2', 'Tokyo2', 'Tokyo', 'Name', 'Second', 'Economic', '1', '2', '3', '12', '112', 'disable'), ('5', 'Japan', 'Tokyo22', 'Tokyo22', 'Tokyo2', 'Changqi2', 'Third', 'Economic', '1', '2', '4', '22', '222', 'disable'), ('6', 'New Zealand', 'Northen Island', 'Wahakatane', 'WHL', 'Ohope Beach', 'First', 'Hot', '', '', '', 'Expensive', '3999', ''), ('7', 'New Zealand', 'Northen Island', 'Wahakatane', 'WHL', 'Ohope Beach2', 'First', 'Hot', 'Bsi', 'sf', 'asdfaf', 'asdfasf', '1321', 'enable'), ('8', 'New Zealand', 'Northen Island', 'Wahakatane', 'Fuga', 'Changqi', 'First', 'Hot', 'wfs', 'sdfa', 'asdfa', 'asdfff', '12', 'enable'), ('9', 'New Zealand', 'Northen Island', 'Fuga', 'Tokyo', 'Ohope Beach', 'First', 'Hot', 'asd', 'asd', 'sada', 'asdasd', '1212', 'enable');
COMMIT;

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
`post_id`  bigint(20) NOT NULL AUTO_INCREMENT ,
`Created`  bigint(20) NULL DEFAULT NULL ,
`Title`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
`Body`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
PRIMARY KEY (`post_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=4

;

-- ----------------------------
-- Records of posts
-- ----------------------------
BEGIN;
INSERT INTO `posts` VALUES ('1', '1436948769', 'Post 1', 'Lorem ipsum lorem ipsum'), ('2', '1436948769', 'Post 2', 'This is my second post'), ('3', '1436948781', 'adsf', 'asdfa');
COMMIT;

-- ----------------------------
-- Auto increment value for ph_users
-- ----------------------------
ALTER TABLE `ph_users` AUTO_INCREMENT=2;

-- ----------------------------
-- Auto increment value for ph_view_spots
-- ----------------------------
ALTER TABLE `ph_view_spots` AUTO_INCREMENT=10;

-- ----------------------------
-- Auto increment value for posts
-- ----------------------------
ALTER TABLE `posts` AUTO_INCREMENT=4;
