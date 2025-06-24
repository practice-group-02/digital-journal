CREATE TABLE `programs`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` TEXT NOT NULL,
    `description` TEXT NOT NULL,
    `deadline` DATE NOT NULL,
    `country_id` INT NOT NULL,
    `duration` TEXT NOT NULL,
    `degree` TEXT NOT NULL,
    `funding` TEXT NOT NULL,
    `language_id` INT NOT NULL,
    `university_id` INT NOT NULL
);
CREATE TABLE `countries`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` TEXT NOT NULL
);
CREATE TABLE `tags`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` TEXT NOT NULL
);
CREATE TABLE `program_tags`(
    `program_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `tag_id` INT NOT NULL,
    PRIMARY KEY(`tag_id`)
);
CREATE TABLE `universities`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` TEXT NOT NULL,
    `website` TEXT NOT NULL
);
CREATE TABLE `eligibility`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `program_id` BIGINT NOT NULL,
    `citizenship` TEXT NOT NULL,
    `age_range` TEXT NOT NULL,
    `education_level` TEXT NOT NULL,
    `experience_required` TEXT NOT NULL
);
CREATE TABLE `documents`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `program_id` BIGINT NOT NULL,
    `name` TEXT NOT NULL
);
CREATE TABLE `languages`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` TEXT NOT NULL
);
ALTER TABLE
    `programs` ADD CONSTRAINT `programs_country_id_foreign` FOREIGN KEY(`country_id`) REFERENCES `countries`(`id`);
ALTER TABLE
    `eligibility` ADD CONSTRAINT `eligibility_program_id_foreign` FOREIGN KEY(`program_id`) REFERENCES `programs`(`id`);
ALTER TABLE
    `tags` ADD CONSTRAINT `tags_id_foreign` FOREIGN KEY(`id`) REFERENCES `program_tags`(`tag_id`);
ALTER TABLE
    `documents` ADD CONSTRAINT `documents_program_id_foreign` FOREIGN KEY(`program_id`) REFERENCES `programs`(`id`);
ALTER TABLE
    `programs` ADD CONSTRAINT `programs_university_id_foreign` FOREIGN KEY(`university_id`) REFERENCES `universities`(`id`);
ALTER TABLE
    `program_tags` ADD CONSTRAINT `program_tags_program_id_foreign` FOREIGN KEY(`program_id`) REFERENCES `programs`(`id`);
ALTER TABLE
    `programs` ADD CONSTRAINT `programs_language_id_foreign` FOREIGN KEY(`language_id`) REFERENCES `languages`(`id`);