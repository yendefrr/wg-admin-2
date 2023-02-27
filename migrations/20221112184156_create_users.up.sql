CREATE TABLE users (
    id int(11) not null primary key AUTO_INCREMENT,
    username varchar(32) not null unique,
    password_hash varchar(512) not null,
    role varchar(15) default "guest",
    has_telegram bool default false
);