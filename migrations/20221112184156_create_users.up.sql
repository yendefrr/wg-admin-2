CREATE TABLE users (
    id int(11) not null primary key AUTO_INCREMENT,
    username varchar(32) not null unique,
    has_telegram bool default false
);