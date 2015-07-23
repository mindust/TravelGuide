/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Version : 50625
 Source Host           : localhost
 Source Database       : kanghui

 Target Server Version : 50625
 File Encoding         : utf-8

 Date: 07/23/2015 16:07:20 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `ph_fee_record`
-- ----------------------------
DROP TABLE IF EXISTS `ph_fee_record`;
CREATE TABLE `ph_fee_record` (
	`ID` varchar(36) NOT NULL,
	`USER_ID` int(36) DEFAULT NULL,
	`TRAVEL_PACKAGES_ID` int(36) DEFAULT NULL,
	`VIEW_SPOT_ID` varchar(36) DEFAULT NULL,
	`FEE` varchar(50) DEFAULT NULL,
	`FEE_DATE` date DEFAULT NULL,
	`EQUIP_ID` varchar(36) DEFAULT NULL,
	PRIMARY KEY (`ID`)
) ENGINE=`InnoDB` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT CHECKSUM=0 DELAY_KEY_WRITE=0;

-- ----------------------------
--  Table structure for `ph_travel_package_view_spots`
-- ----------------------------
DROP TABLE IF EXISTS `ph_travel_package_view_spots`;
CREATE TABLE `ph_travel_package_view_spots` (
	`ID` varchar(36) NOT NULL,
	`TRAVEL_PACKAGES_ID` varchar(36) DEFAULT NULL,
	`VIEW_SPOT_ID` varchar(36) DEFAULT NULL,
	`DESCRIPTION` varchar(500) DEFAULT NULL,
	`DAYS` int(2) DEFAULT NULL,
	`PLAY_ITEMS` varchar(500) DEFAULT NULL,
	PRIMARY KEY (`ID`)
) ENGINE=`InnoDB` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT CHECKSUM=0 DELAY_KEY_WRITE=0;

-- ----------------------------
--  Table structure for `ph_travel_packages`
-- ----------------------------
DROP TABLE IF EXISTS `ph_travel_packages`;
CREATE TABLE `ph_travel_packages` (
	`ID` varchar(36) NOT NULL,
	`NAME` varchar(36) DEFAULT NULL,
	`DESCRIPTION` varchar(500) DEFAULT NULL,
	`FEE` varchar(300) DEFAULT NULL,
	`START_DATE` date DEFAULT NULL,
	`END_DATE` date DEFAULT NULL,
	`DAYS` int(2) DEFAULT NULL,
	`HOTELS` varchar(300) DEFAULT NULL,
	`TRANSPOT` varchar(300) DEFAULT NULL,
	`LEVEL` int(1) DEFAULT NULL,
	`PERSON_NUM` varchar(36) DEFAULT NULL,
	`LABLE` varchar(36) DEFAULT NULL,
	PRIMARY KEY (`ID`)
) ENGINE=`InnoDB` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT CHECKSUM=0 DELAY_KEY_WRITE=0;

-- ----------------------------
--  Table structure for `ph_users`
-- ----------------------------
DROP TABLE IF EXISTS `ph_users`;
CREATE TABLE `ph_users` (
	`ID` int(36) NOT NULL AUTO_INCREMENT,
	`USER_NAME` varchar(36) DEFAULT NULL,
	`PASSWORD` varchar(50) DEFAULT NULL,
	`DESCRIPTION` varchar(300) DEFAULT NULL,
	`QQ_NO` varchar(30) DEFAULT NULL,
	`MOBILE` varchar(20) DEFAULT NULL,
	`EMAIL` varchar(100) DEFAULT NULL,
	`LEVEL` varchar(30) DEFAULT NULL,
	`LOGO_PATH` varchar(200) DEFAULT NULL,
	PRIMARY KEY (`ID`)
) ENGINE=`InnoDB` AUTO_INCREMENT=2 DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT CHECKSUM=0 DELAY_KEY_WRITE=0;

-- ----------------------------
--  Records of `ph_users`
-- ----------------------------
BEGIN;
INSERT INTO `ph_users` VALUES ('1', 'Messi', '123JKLjkl', 'Messi', '52486846', '18013579965', 'messi.he@sig.biz', '1', '2');
COMMIT;

-- ----------------------------
--  Table structure for `ph_view_spot_images`
-- ----------------------------
DROP TABLE IF EXISTS `ph_view_spot_images`;
CREATE TABLE `ph_view_spot_images` (
	`ID` varchar(36) NOT NULL,
	`VIEW_SPOT_ID` varchar(36) DEFAULT NULL,
	`PATH` varchar(300) DEFAULT NULL,
	`NAME` varchar(50) DEFAULT NULL,
	`FORMAT` varchar(20) DEFAULT NULL,
	`SOURCE_NAME` varchar(200) DEFAULT NULL,
	PRIMARY KEY (`ID`)
) ENGINE=`InnoDB` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT CHECKSUM=0 DELAY_KEY_WRITE=0;

-- ----------------------------
--  Records of `ph_view_spot_images`
-- ----------------------------
BEGIN;
INSERT INTO `ph_view_spot_images` VALUES ('30375d0c-3c8d-4d29-9b22-ed9b925966bd', '75db7c06-bbee-46a5-8153-81d9c659486f', 'uploads', 'c718a66b-611d-4606-a8b5-8d820ab170fa', 'jpg', '圆融星座.jpg'), ('8a0df19e-dcc2-4890-bd81-4f27f391e74d', '75db7c06-bbee-46a5-8153-81d9c659486f', 'uploads', 'e171262c-725b-4cbe-8971-9db61190dd88', 'png', '1ED7909F5095C9C0D0F43A138D9401E6.png');
COMMIT;

-- ----------------------------
--  Table structure for `ph_view_spots`
-- ----------------------------
DROP TABLE IF EXISTS `ph_view_spots`;
CREATE TABLE `ph_view_spots` (
	`ID` varchar(100) NOT NULL,
	`COUNTRY` varchar(100) DEFAULT NULL,
	`PROVINCE` varchar(100) DEFAULT NULL,
	`CITY` varchar(100) DEFAULT NULL,
	`COUNTY` varchar(100) DEFAULT NULL,
	`NAME` varchar(300) DEFAULT NULL,
	`LEVEL` varchar(30) DEFAULT NULL,
	`LABEL` varchar(200) DEFAULT NULL,
	`PRICE` varchar(300) DEFAULT NULL,
	`STATUS` int(1) DEFAULT NULL,
	PRIMARY KEY (`ID`)
) ENGINE=`InnoDB` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT CHECKSUM=0 DELAY_KEY_WRITE=0;

-- ----------------------------
--  Records of `ph_view_spots`
-- ----------------------------
BEGIN;
INSERT INTO `ph_view_spots` VALUES ('1', 'China ', 'JiangSu', 'Suzhou', '1', 'TaiLake', '5', 'nutual view', '2000', '1'), ('2', 'Japan', 'Fuga', 'Fuga', 'Fuga', 'Changqi', '1', 'Fuga ChangQi', '11111', '1'), ('3', 'Japan', 'Tokyo', 'Tokyo', 'Tokyo', 'Name', '2', 'adsf', '121', '1'), ('4', 'Japan', 'Tokyo2', 'Tokyo2', 'Tokyo', 'Name', '3', '12', '112', '1'), ('5', 'Japan', 'Tokyo22', 'Tokyo22', 'Tokyo2', 'Changqi2', '12', '22', '222', '1'), ('75db7c06-bbee-46a5-8153-81d9c659486f', '1', '2', '3', '4', '5', '6', '8', '9', '10');
COMMIT;

-- ----------------------------
--  Table structure for `posts`
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
	`post_id` bigint(20) NOT NULL AUTO_INCREMENT,
	`Created` bigint(20) DEFAULT NULL,
	`Title` varchar(255) DEFAULT NULL,
	`Body` varchar(255) DEFAULT NULL,
	PRIMARY KEY (`post_id`)
) ENGINE=`InnoDB` AUTO_INCREMENT=4 DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ROW_FORMAT=COMPACT CHECKSUM=0 DELAY_KEY_WRITE=0;

-- ----------------------------
--  Records of `posts`
-- ----------------------------
BEGIN;
INSERT INTO `posts` VALUES ('1', '1436948769', 'Post 1', 'Lorem ipsum lorem ipsum'), ('2', '1436948769', 'Post 2', 'This is my second post'), ('3', '1436948781', 'adsf', 'asdfa');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
