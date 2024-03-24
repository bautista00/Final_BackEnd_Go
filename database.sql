-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema database_Go
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema database_Go
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `database_Go` DEFAULT CHARACTER SET utf8mb3 ;
USE `database_Go` ;

-- -----------------------------------------------------
-- Table `database_Go`.`odontologos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `database_Go`.`odontologos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(100) NULL DEFAULT NULL,
  `apellido` VARCHAR(100) NULL DEFAULT NULL,
  `matricula` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `database_Go`.`pacientes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `database_Go`.`pacientes` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(100) NULL DEFAULT NULL,
  `apellido` VARCHAR(100) NULL DEFAULT NULL,
  `domicilio` VARCHAR(200) NULL DEFAULT NULL,
  `dni` VARCHAR(100) NULL DEFAULT NULL,
  `fecha_alta` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `database_Go`.`turnos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `database_Go`.`turnos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `paciente` INT NULL,
  `odontologo` INT NULL DEFAULT NULL,
  `fechaHora` DATETIME NULL DEFAULT NULL,
  `descripcion` VARCHAR(500) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `OT_idx` (`odontologo` ASC) VISIBLE,
  INDEX `PT_idx` (`paciente` ASC) VISIBLE,
  CONSTRAINT `OT`
    FOREIGN KEY (`odontologo`)
    REFERENCES `database_Go`.`odontologos` (`id`),
  CONSTRAINT `PT`
    FOREIGN KEY (`paciente`)
    REFERENCES `database_Go`.`pacientes` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;