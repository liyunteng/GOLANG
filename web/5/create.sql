-- for mysql

-- CREATE TABLE `userinfo` (
--         `uid` INT(10) NOT NULL AUTO_INCREMENT,
--         `username` VARCHAR(64) NULL DEFAULT NULL,
--         `departname` VARCHAR(64) NULL DEFAULT NULL,
--         `created` DATE NULL DEFAULT NULL,
--         PRIMARY KEY (`uid`)
--         )

-- CREATE TABLE `userdetail` (
--         `uid` INT(10) NOT NULL DEFAULT '0',
--         `intro` TEXT NULL,
--         `profile` TEXT NULL,
--         PRIMARY KEY (`uid`)
--         )


-- for sqlite3
CREATE TABLE `userinfo` (
        `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
        `username` VARCHAR(64) NULL,
        `departname` VARCHAR(64) NULL,
        `created` DATE NULL
        );

CREATE TABLE `userdetail` (
        `uid` INT(10) NULL,
        `intro` TEXT NULL,
        `profile` TEXT NULL,
        PRIMARY KEY (`uid`)
        );
