CREATE TABLE profiles (
   id int(11) not null primary key AUTO_INCREMENT,
   username varchar(32) not null,
   name varchar(32) not null,
   public_key_b64 text default null,
   private_key_b64 text default null,
   client_config text default null,
   client_qr_code text default null,
   created_at datetime default null,
   updated_at datetime default null,
   is_active bool default false
);