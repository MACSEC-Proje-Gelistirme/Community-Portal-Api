/*CREATE USER community  WITH PASSWORD 'password';

CREATE DATABASE community;

GRANT ALL PRIVILEGES ON DATABASE community TO community; */

/* You need to reconnect db before next line with user community and db community. */

CREATE TABLE IF NOT EXISTS users  (
   id  varchar PRIMARY KEY,
   school_id  varchar,
   first_name  varchar,
   last_name  varchar,
   email  varchar,
   telephone_number  varchar,
   email_preferences  varchar,
   marketing_preferences  varchar,
   created_at  timestamp,
   updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS  clubs  (
   id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
   name  varchar,
   description  text,
   email  varchar,
   member_count  varchar,
   created_at  timestamp,
   updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS events  (
   id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
   club_id  UUID,
   name  varchar,
   description  varchar,
   date  timestamp,
   tags  varchar,
   location  varchar,
   created_at  timestamp,
   updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS club_roles  (
   user_id  varchar,
   club_id  UUID,
   role  varchar,
   created_at  timestamp,
   updated_at  timestamp,
  PRIMARY KEY ( user_id ,  club_id )
);

CREATE TABLE IF NOT EXISTS feed_posts  (
   id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
   author_club_id  UUID,
   author_user_id  varchar,
   iamge  text,
   description  text,
   like_count  varchar,
   created_at  timestamp,
   updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS gallery_posts  (
   id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
   author_club_id  UUID,
   author_user_id  varchar,
   iamge  text,
   description  text,
   created_at  timestamp,
   updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS comments  (
   id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
   author_user_id  varchar,
   post_id  UUID,
   comment_id  UUID,
   comment  text,
   like_count  varchar,
   created_at  timestamp,
   updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS likes  (
   user_id  varchar,
   post_id  UUID,
   created_at  timestamp,
   updated_at  timestamp,
  PRIMARY KEY ( user_id ,  post_id )
);

CREATE TABLE IF NOT EXISTS attended_events  (
   id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
   user_id  varchar,
   event_id  UUID,
   situation  varchar,
   created_at  timestamp,
   updated_at  timestamp
);

CREATE TABLE IF NOT EXISTS mails  (
   id  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
   author_club_id  UUID,
   author_user_id  varchar,
   subject  varchar,
   content  text,
   recipients  text,
   created_at  timestamp,
   updated_at  timestamp
);

ALTER TABLE  events  ADD FOREIGN KEY ( club_id ) REFERENCES  clubs  ( id );

ALTER TABLE  club_roles  ADD FOREIGN KEY ( user_id ) REFERENCES  users  ( id );

ALTER TABLE  club_roles  ADD FOREIGN KEY ( club_id ) REFERENCES  clubs  ( id );

ALTER TABLE  feed_posts  ADD FOREIGN KEY ( author_club_id ) REFERENCES  clubs  ( id );

ALTER TABLE  feed_posts  ADD FOREIGN KEY ( author_user_id ) REFERENCES  users  ( id );

ALTER TABLE  gallery_posts  ADD FOREIGN KEY ( author_club_id ) REFERENCES  clubs  ( id );

ALTER TABLE  gallery_posts  ADD FOREIGN KEY ( author_user_id ) REFERENCES  users  ( id );

ALTER TABLE  comments  ADD FOREIGN KEY ( author_user_id ) REFERENCES  users  ( id );

ALTER TABLE  comments  ADD FOREIGN KEY ( post_id ) REFERENCES  feed_posts  ( id );

ALTER TABLE  likes  ADD FOREIGN KEY ( post_id ) REFERENCES  feed_posts  ( id );

ALTER TABLE  likes  ADD FOREIGN KEY ( user_id ) REFERENCES  users  ( id );

ALTER TABLE  attended_events  ADD FOREIGN KEY ( user_id ) REFERENCES  users  ( id );

ALTER TABLE  attended_events  ADD FOREIGN KEY ( event_id ) REFERENCES  events  ( id );

ALTER TABLE  mails  ADD FOREIGN KEY ( author_club_id ) REFERENCES  clubs  ( id );

ALTER TABLE  mails  ADD FOREIGN KEY ( author_user_id ) REFERENCES  users  ( id );