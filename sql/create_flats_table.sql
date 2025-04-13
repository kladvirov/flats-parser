CREATE TABLE `mydb`.`flats`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `remote_id`  INT UNSIGNED,
    `type` 		 TINYINT UNSIGNED,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP() NOT NULL,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP() NOT NULL ON
             UPDATE CURRENT_TIMESTAMP ()
);

CREATE INDEX remote_id
    ON `mydb`.flats (remote_id);