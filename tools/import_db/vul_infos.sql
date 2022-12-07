CREATE TABLE `vul_infos` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
`vuln-id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
`published` date DEFAULT NULL,
`modified` date DEFAULT NULL,
`source` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL,
`severity` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
`vuln-type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
`vuln-descript` text COLLATE utf8_unicode_ci,
`cve-id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
`bugtraq-id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
`vuln-solution` text COLLATE utf8_unicode_ci,
PRIMARY KEY (`id`),
UNIQUE KEY `idx_vul_infos_cve_id` (`cve-id`)
) ENGINE=MyISAM AUTO_INCREMENT=1859 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;