CREATE TABLE sample_info(
                            id int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
                            `md5` varchar(32) NOT NULL DEFAULT '' COMMENT 'MD5',
                            `sha1` varchar(60) NOT NULL DEFAULT '' COMMENT 'SHA1',
                            `create_time` datetime DEFAULT NOW() COMMENT '创建时间',
                            `level` tinyint(1) DEFAULT 0 COMMENT '级别',
                            `operator` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `key_sha1` (`sha1`),
                            KEY `key_md5` (`md5`),
                            KEY `key_level` (`level`),
                            KEY `key_operator` (`operator`)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 COMMENT='sample_info';

INSERT INTO sample_info(`md5`,`sha1`,`level`,`operator`)VALUES("c253201e6ac8857e24b838d781c81099","068b987d3a102a69c1b8377a54c3bc7bd2a1287c",70,"admin");
INSERT INTO sample_info(`md5`,`sha1`,`level`,`operator`)VALUES("dc61f178f6e846e32cb65328c549e5a8","fec0ec0c1814703675680bf2e34dc3d14630cba0",70,"admin");

