-- MySQL dump 10.13  Distrib 5.7.36, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: comunion
-- ------------------------------------------------------
-- Server version	5.7.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `bounty`
--

DROP TABLE IF EXISTS `bounty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bounty` (
  `id` bigint(20) NOT NULL,
  `startup_id` bigint(20) NOT NULL COMMENT 'startup id',
  `comer_id` bigint(20) NOT NULL,
  `title` varchar(255) NOT NULL COMMENT 'title',
  `type` tinyint(4) NOT NULL COMMENT 'type 1 for business',
  `intro` text NOT NULL COMMENT 'intro',
  `address` varchar(50) NOT NULL COMMENT 'block chain address',
  `state` tinyint(4) NOT NULL COMMENT 'bounty state 1 for open 2 for processing 3 for closed',
  `description_url` text COMMENT 'description url',
  `contact_email` varchar(255) DEFAULT NULL COMMENT 'contract email',
  `duration_days` int(11) DEFAULT NULL COMMENT 'duration days',
  `payment_ethereum_token_id` bigint(20) DEFAULT NULL COMMENT 'payment ethereum token id',
  `payment_amount` bigint(20) DEFAULT NULL COMMENT 'payment amount',
  `started_at` datetime DEFAULT NULL COMMENT 'bounty start time',
  `closed_at` datetime DEFAULT NULL COMMENT 'bounty end time',
  `created_by` bigint(20) NOT NULL COMMENT 'comer id created by',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-state` (`state`) USING BTREE,
  KEY `idx-startup_id` (`startup_id`) USING BTREE,
  KEY `idx-created_by` (`created_by`) USING BTREE,
  KEY `idx-comer_id` (`comer_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `bounty_hunter`
--

DROP TABLE IF EXISTS `bounty_hunter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bounty_hunter` (
  `id` bigint(20) NOT NULL,
  `bounty_id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `started_at` datetime DEFAULT NULL,
  `submitted_at` datetime DEFAULT NULL,
  `paid_at` datetime DEFAULT NULL,
  `rejected_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-bounty_id` (`bounty_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `bounty_tag_rel`
--

DROP TABLE IF EXISTS `bounty_tag_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bounty_tag_rel` (
  `bounty_id` bigint(20) NOT NULL COMMENT 'bounty id',
  `tag_id` bigint(20) NOT NULL COMMENT 'tag id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`bounty_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer`
--

DROP TABLE IF EXISTS `comer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer` (
  `id` bigint(20) NOT NULL,
  `address` char(42) NOT NULL COMMENT 'comer could save some useful info on block chain with this address',
  `nick` varchar(50) NOT NULL COMMENT 'nick',
  `city` varchar(50) DEFAULT NULL COMMENT 'city',
  `avatar` varchar(255) DEFAULT NULL COMMENT 'avatar link address',
  `blog` text COMMENT 'blog address',
  `intro` text COMMENT 'intro',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account` (
  `id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL COMMENT 'comer unique identifier',
  `oin` varchar(100) NOT NULL COMMENT 'comer outer account unique identifier, wallet will be public key and Oauth is the OauthID',
  `is_primary` boolean NOT NULL COMMENT 'comer use this account as primay account',
  `nick` varchar(50) NOT NULL COMMENT 'comer nick name',
  `avatar` varchar(255) NOT NULL COMMENT 'avatar link address',
  `type` int(11) NOT NULL COMMENT '1 for github 2 for twitter 3 for facebook 4 for likedin 5 for google',
  `is_linked` boolean NOT NULL COMMENT '0 for unlink 1 for linked',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni-oin` (`oin`) USING BTREE,
  KEY `idx-comer_id` (`comer_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_followed_comer_rel`
--

DROP TABLE IF EXISTS `comer_followed_comer_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_followed_comer_rel` (
  `comer_id` bigint(20) NOT NULL,
  `followed_comer_id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`comer_id`,`followed_comer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_followed_startup_rel`
--

DROP TABLE IF EXISTS `comer_followed_startup_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_followed_startup_rel` (
  `comer_id` bigint(20) NOT NULL,
  `followed_startup_id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`comer_id`,`followed_startup_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_followed_tag_rel`
--

DROP TABLE IF EXISTS `comer_followed_tag_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_followed_tag_rel` (
  `comer_id` bigint(20) NOT NULL,
  `followed_tag_id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`comer_id`,`followed_tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_skill_rel`
--

DROP TABLE IF EXISTS `comer_skill_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_skill_rel` (
  `id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL COMMENT 'comer id',
  `skill_id` bigint(20) NOT NULL COMMENT 'skill id',
  `is_delete` tinyint(1) NOT NULL DEFAULT 0  COMMENT 'Is Delete',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `comer_id_skill_id` (`comer_id`,`skill_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `currency`
--

DROP TABLE IF EXISTS `currency`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `currency` (
  `name` char(3) NOT NULL COMMENT 'name',
  `symbol` varchar(10) DEFAULT NULL COMMENT 'symbol',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `disco`
--

DROP TABLE IF EXISTS `disco`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `disco` (
  `id` bigint(20) NOT NULL,
  `startup_id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL,
  `wallet_address` char(42) NOT NULL,
  `etherenum_token_id` bigint(20) NOT NULL,
  `description` text NOT NULL,
  `fund_raising_started_at` datetime NOT NULL,
  `fund_raising_ended_at` datetime NOT NULL,
  `investment_reward` bigint(20) NOT NULL,
  `reward_decline_rate` int(11) NOT NULL,
  `share_token` bigint(20) NOT NULL,
  `min_fund_raising` bigint(20) NOT NULL,
  `add_liquidity_pool` bigint(20) NOT NULL,
  `total_deposit_token` bigint(20) NOT NULL,
  `state` int(11) NOT NULL,
  `fund_raising_address` char(42) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-startup_id` (`startup_id`) USING BTREE,
  KEY `idx-comer_id` (`comer_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `disco_investment`
--

DROP TABLE IF EXISTS `disco_investment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `disco_investment` (
  `id` bigint(20) NOT NULL,
  `disco_id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL,
  `etherenum_token_id` bigint(20) NOT NULL,
  `amount` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-disco_id` (`disco_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ethereum_token`
--

DROP TABLE IF EXISTS `ethereum_token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ethereum_token` (
  `id` bigint(20) NOT NULL,
  `standard_type` tinyint(4) NOT NULL COMMENT 'standard type 1 for erc-20',
  `name` varchar(10) NOT NULL COMMENT 'name',
  `contract_address` char(42) NOT NULL COMMENT 'contract address',
  `symbol` varchar(10) NOT NULL COMMENT 'symbol',
  `decimals` int(11) NOT NULL COMMENT 'decimal',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `proposal`
--

DROP TABLE IF EXISTS `proposal`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `proposal` (
  `id` bigint(20) NOT NULL,
  `startup_id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL,
  `wallet_address` char(42) NOT NULL,
  `contract_address` char(42) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL,
  `type` int(11) NOT NULL,
  `contact` char(40) NOT NULL,
  `description` text NOT NULL,
  `voter_type` tinyint(4) NOT NULL,
  `supporters` int(11) NOT NULL,
  `minimum_approval_percentage` int(11) NOT NULL,
  `duration_days` int(11) NOT NULL,
  `has_payment` tinyint(1) NOT NULL,
  `payment_address` text,
  `payment_type` tinyint(4) DEFAULT NULL,
  `payment_months` int(11) DEFAULT NULL,
  `payment_date` text,
  `payment_amount` bigint(20) DEFAULT NULL,
  `total_payment_amount` bigint(20) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-comer_id` (`comer_id`) USING BTREE,
  KEY `idx-startup_id` (`startup_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `proposal_tag_rel`
--

DROP TABLE IF EXISTS `proposal_tag_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `proposal_tag_rel` (
  `proposal_id` bigint(20) NOT NULL COMMENT 'proposal id',
  `tag_id` bigint(20) NOT NULL COMMENT 'tag id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`proposal_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `proposal_term`
--

DROP TABLE IF EXISTS `proposal_term`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `proposal_term` (
  `id` bigint(20) NOT NULL,
  `proposal_id` bigint(20) NOT NULL,
  `etherenum_token_id` bigint(20) NOT NULL,
  `amount` bigint(20) NOT NULL,
  `content` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-proposal_id` (`proposal_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `proposal_vote`
--

DROP TABLE IF EXISTS `proposal_vote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `proposal_vote` (
  `id` bigint(20) NOT NULL,
  `proposal_id` bigint(20) NOT NULL,
  `amount` bigint(20) NOT NULL,
  `is_approved` tinyint(1) NOT NULL,
  `wallet_address` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-proposal_id` (`proposal_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_skill`
--

DROP TABLE IF EXISTS `comer_skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_skill` (
  `id` bigint(20) NOT NULL,
  `name` varchar(64) NOT NULL COMMENT 'name',
  `is_delete` tinyint(1) NOT NULL DEFAULT 0  COMMENT 'Is Delete',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `startup`
--

DROP TABLE IF EXISTS `startup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `startup` (
  `id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL COMMENT 'name',
  `logo` text NOT NULL COMMENT 'logo',
  `description_url` text NOT NULL,
  `mission` text NOT NULL,
  `category_id` bigint(20) NOT NULL,
  `etherenum_token_id` bigint(20) NOT NULL,
  `voter_token_limit` int(11) DEFAULT NULL,
  `proposer_type` int(11) DEFAULT NULL,
  `proposer_token_limit` int(11) DEFAULT NULL,
  `proposal_supporters` int(11) DEFAULT NULL,
  `proposal_min_approval_percent` int(11) DEFAULT NULL,
  `proposal_min_duration` int(11) DEFAULT NULL,
  `proposal_max_duration` int(11) DEFAULT NULL,
  `voter_type` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-comer_id` (`comer_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `startup_assigned_proposer`
--

DROP TABLE IF EXISTS `startup_assigned_proposer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `startup_assigned_proposer` (
  `id` bigint(20) NOT NULL,
  `startup_id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-startup_id` (`startup_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `startup_assigned_voter`
--

DROP TABLE IF EXISTS `startup_assigned_voter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `startup_assigned_voter` (
  `id` bigint(20) NOT NULL,
  `startup_id` bigint(20) NOT NULL,
  `comer_id` bigint(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-startup_id` (`startup_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `startup_member`
--

DROP TABLE IF EXISTS `startup_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `startup_member` (
  `id` bigint(20) NOT NULL,
  `startup_id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `location` text,
  `intro` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-startup_id` (`startup_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `startup_tag_rel`
--

DROP TABLE IF EXISTS `startup_tag_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `startup_tag_rel` (
  `startup_id` bigint(20) NOT NULL COMMENT 'comerup id',
  `tag_id` bigint(20) NOT NULL COMMENT 'tag id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`startup_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `startup_wallet`
--

DROP TABLE IF EXISTS `startup_wallet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `startup_wallet` (
  `id` bigint(20) NOT NULL,
  `startup_id` bigint(20) NOT NULL COMMENT 'comer unique identifier',
  `name` varchar(50) NOT NULL,
  `address` char(40) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx-startup_id` (`startup_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS `tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tag` (
  `id` bigint(20) NOT NULL,
  `name` varchar(55) NOT NULL COMMENT 'name',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transaction` (
  `hash` char(40) NOT NULL,
  `address` char(40) NOT NULL,
  `state` tinyint(4) NOT NULL,
  `target_type` tinyint(4) NOT NULL COMMENT '1.startup 2.bounty 3.disco 4.disco_investment 5.proposal 6.proposal_vote',
  `target_id` bigint(20) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_wallet`
--

DROP TABLE IF EXISTS `comer_wallet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_wallet` (
    `id` bigint(20) NOT NULL,
    `address` varchar(55) NOT NULL COMMENT 'name',
    `is_delete` tinyint(1) NOT NULL DEFAULT 0  COMMENT 'Is Delete',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `address` (`address`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_wallet_rel`
--

DROP TABLE IF EXISTS `comer_wallet_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_wallet_rel` (
    `id` bigint(20) NOT NULL,
    `comer_id` bigint(20) NOT NULL COMMENT 'comer id',
    `wallet_id` bigint(20) NOT NULL COMMENT 'wallet id',
    `is_delete` tinyint(1) NOT NULL DEFAULT 0  COMMENT 'Is Delete',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `comer_id_wallet_id` (`comer_id`,`wallet_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_social`
--

DROP TABLE IF EXISTS `comer_social`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_social` (
    `id` bigint(20) NOT NULL,
    `type` int NOT NULL DEFAULT 0 COMMENT '0 GitHub 1 Google',
    `account` varchar(64) NOT NULL COMMENT 'Social Account',
    `is_delete` tinyint(1) NOT NULL DEFAULT 0  COMMENT 'Is Delete',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `account` (`account`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_wallet_rel`
--

DROP TABLE IF EXISTS `comer_social_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_social_rel` (
    `id` bigint(20) NOT NULL,
    `comer_id` bigint(20) NOT NULL COMMENT 'comer id',
    `social_id` bigint(20) NOT NULL COMMENT 'social id',
    `is_delete` tinyint(1) NOT NULL DEFAULT 0  COMMENT 'Is Delete',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `comer_id_social_id` (`comer_id`,`social_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `comer_profile`
--

DROP TABLE IF EXISTS `comer_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comer_profile` (
    `id` bigint(20) NOT NULL,
    `comer_id` char(40) NOT NULL COMMENT 'comerId',
    `name` varchar(50) NOT NULL COMMENT 'Name',
    `location` varchar(100) DEFAULT NULL COMMENT 'Location',
    `website` varchar(100) DEFAULT NULL COMMENT 'Website',
    `bio` text DEFAULT NULL COMMENT 'Bio',
    `is_delete` tinyint(1) NOT NULL DEFAULT 0  COMMENT 'Is Delete',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `comer_id` (`comer_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;


/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-21 11:28:35
