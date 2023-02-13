CREATE TABLE profiles (
   id int(11) not null primary key AUTO_INCREMENT,
   username varchar(32) not null,
   name varchar(32) not null,
   publickey_base64 text default null,
   privatekey_base64 text default null,
   client_config text default null,
   client_qrcode text default null,
   is_active bool default false
);