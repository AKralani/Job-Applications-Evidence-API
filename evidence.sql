CREATE DATABASE jobapplicationsdb;

USE jobapplicationsdb;

CREATE TABLE `jobapplicationsdb`.`evidences` (
  `evidenceId` INT NOT NULL AUTO_INCREMENT,
  `date` DATE NOT NULL,
  `companyName` VARCHAR(255) NOT NULL,
  `link` BLOB NOT NULL,
  `jobDescription` VARCHAR(255) NOT NULL,
  `location` VARCHAR(255) NOT NULL,
  `jobType` VARCHAR(255) NOT NULL,
  `field` VARCHAR(255) NOT NULL,
  `interviewDate` DATE,
  `interviewDescription` VARCHAR(255),
  `accepted` BOOL,
  PRIMARY KEY (`evidenceId`));