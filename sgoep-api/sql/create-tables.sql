CREATE TABLE `skills` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `title` VARCHAR(64),
  `icon` VARCHAR(64),
  `desc` TEXT(512),
  `like` BOOLEAN
);

CREATE TABLE `author` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR(64),
  `title` VARCHAR(128),
  `subtitle` VARCHAR(128),
  `about` TEXT(512),
  `github` VARCHAR(64),
  `linkedin` VARCHAR(64)
);

CREATE TABLE `achievements` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` VARCHAR(64),
	`subtitle` VARCHAR(128),
	`desc` TEXT(512),
	`icon` VARCHAR(64),
	`link` VARCHAR(128),
	`link_icon` VARCHAR(64),
	`github` VARCHAR(128),
  `tech` TEXT(512)
);

CREATE TABLE `interests` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` VARCHAR(64),
	`icon` VARCHAR(64),
	`desc` TEXT(512)
);
