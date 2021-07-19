CREATE DATABASE `msql_test`;
CREATE TABLE `tests` (
    `id` BIGINT AUTO_INCREMENT NOT NULL,
    `name` varchar(20) NOT NULL DEFAULT '',
    `created_at` DATETIME NOT NULL,
    `disable` BOOL DEFAULT FALSE,
    `flag` TINYINT UNSIGNED NOT NULL DEFAULT 0,
    `sex` TINYINT NOT NULL DEFAULT 0,
    `score` INT UNSIGNED NOT NULL DEFAULT 0,
    `fields` JSON NOT NULL,
    PRIMARY KEY(`id`)
);

insert into tests (name,created_at,fields) values ("小明",now(),'{"ip":"","login_times":20}'),("小王",now(),'{"ip":"","login_times":11}');