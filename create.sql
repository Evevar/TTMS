CREATE TABLE IF NOT EXISTS `users` (
    `id` INTEGER NOT NULL auto_increment,
    `type` INTEGER NOT NULL COMMENT '1-销售，2-经理，9-系统管理员',
    `name` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL COMMENT '电子邮件',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `studios` (
    `id` INTEGER NOT NULL auto_increment,
    `name` VARCHAR(255) NOT NULL,
    `rows_count` INTEGER NOT NULL,
    `cols_count` INTEGER NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `seats` (
    `id` INTEGER NOT NULL auto_increment,
    `studio_id` INTEGER NOT NULL,
    `row` INTEGER NOT NULL,
    `col` INTEGER NOT NULL,
    `status` INTEGER NOT NULL COMMENT '0-不是座位，1-有座位，9-损坏的座位',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `plays` (
    `id` INTEGER NOT NULL auto_increment,
    `name` NVARCHAR(255) NOT NULL,
    `type` INTEGER NOT NULL COMMENT '1-file 2-opear 3-音乐会',
    `area` NVARCHAR(255) NOT NULL,
    `rating` INTEGER NOT NULL COMMENT '剧目等级 1-儿童 2-青年 3-成年',
    `duration` NVARCHAR(255) NOT NULL COMMENT '如：3h30m30s',
    `start_data` NVARCHAR(255) NOT NULL COMMENT '示例格式：2006-01-02',
    `end_data` NVARCHAR(255) NOT NULL COMMENT '示例格式：2006-01-02',
    `price` INTEGER NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `schedules` (
    `id` INTEGER NOT NULL auto_increment,
    `play_id` INTEGER NOT NULL,
    `studio_id` INTEGER NOT NULL,
    `show_time` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB;
