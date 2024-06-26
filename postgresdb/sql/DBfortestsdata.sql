CREATE TABLE "users" (
  "id" integer PRIMARY KEY AUTOINCREMENT,
  "name" varchar(128)
);

CREATE TABLE "posts" (
  "id" integer PRIMARY KEY AUTOINCREMENT,
  "text" text,
  "userid" integer,
  "cancomment" boolean,
  FOREIGN KEY (userid) REFERENCES users(id)
);

CREATE TABLE "comments" (
  "id" integer PRIMARY KEY AUTOINCREMENT,
  "userid" integer,
  "text" varchar(2000),
  "postid" integer,
  "perentid" integer,
  FOREIGN KEY (userid) REFERENCES users(id),
  FOREIGN KEY (postid) REFERENCES posts(id),
  FOREIGN KEY (perentid) REFERENCES comments(id)
);


INSERT INTO "users" (name) VALUES ('Пользователь_1');
INSERT INTO "users" (name) VALUES ('Пользователь_2');
INSERT INTO "users" (name) VALUES ('Пользователь_3');
INSERT INTO "users" (name) VALUES ('Пользователь_4');
INSERT INTO "users" (name) VALUES ('Пользователь_5');
INSERT INTO "users" (name) VALUES ('Пользователь_6');
INSERT INTO "users" (name) VALUES ('Пользователь_7');
INSERT INTO "users" (name) VALUES ('Пользователь_8');
INSERT INTO "users" (name) VALUES ('Пользователь_9');
INSERT INTO "users" (name) VALUES ('Пользователь_10');
INSERT INTO "users" (name) VALUES ('Пользователь_11');
INSERT INTO "users" (name) VALUES ('Пользователь_12');
INSERT INTO "users" (name) VALUES ('Пользователь_13');
INSERT INTO "users" (name) VALUES ('Пользователь_14');
INSERT INTO "users" (name) VALUES ('Пользователь_15');
INSERT INTO "users" (name) VALUES ('Пользователь_16');
INSERT INTO "users" (name) VALUES ('Пользователь_17');
INSERT INTO "users" (name) VALUES ('Пользователь_18');
INSERT INTO "users" (name) VALUES ('Пользователь_19');
INSERT INTO "users" (name) VALUES ('Пользователь_20');
INSERT INTO "users" (name) VALUES ('Пользователь_21');
INSERT INTO "users" (name) VALUES ('Пользователь_22');


INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 1', 'true', id
FROM users 
WHERE "name" = 'Пользователь_1';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 2', 'true', id
FROM users 
WHERE "name" = 'Пользователь_1';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 3', 'true', id
FROM users 
WHERE "name" = 'Пользователь_1';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 4', 'true', id
FROM users 
WHERE "name" = 'Пользователь_1';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста', 'true', id
FROM users 
WHERE "name" = 'Пользователь_3';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 5', 'true', id
FROM users 
WHERE "name" = 'Пользователь_3';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 6', 'true', id
FROM users 
WHERE "name" = 'Пользователь_5';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 7', 'true', id
FROM users 
WHERE "name" = 'Пользователь_5';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 8', 'true', id
FROM users 
WHERE "name" = 'Пользователь_7';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 9', 'true', id
FROM users 
WHERE "name" = 'Пользователь_9';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 10', 'true', id
FROM users 
WHERE "name" = 'Пользователь_9';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 11', 'false', id
FROM users 
WHERE "name" = 'Пользователь_11';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 12', 'false', id
FROM users 
WHERE "name" = 'Пользователь_11';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 13', 'false', id
FROM users 
WHERE "name" = 'Пользователь_13';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 14', 'false', id
FROM users 
WHERE "name" = 'Пользователь_13';

INSERT INTO posts ("text", "cancomment", "userid") 
SELECT 'Текст поста 15', 'true', id
FROM users 
WHERE "name" = 'Пользователь_13';




INSERT INTO comments ("text", "userid", "postid") 
SELECT 'Текст каментария 00', users.id, posts.id
FROM users, posts
WHERE "name" = 'Пользователь_2' AND "text" = 'Текст поста 1';

INSERT INTO comments ("text", "userid", "postid") 
SELECT 'Текст каментария 10', users.id, posts.id
FROM users, posts
WHERE "name" = 'Пользователь_4' AND "text" = 'Текст поста 1';

INSERT INTO comments ("text", "userid", "postid") 
SELECT 'Текст каментария 20', users.id, posts.id
FROM users, posts
WHERE "name" = 'Пользователь_6' AND "text" = 'Текст поста 1';

INSERT INTO comments ("text", "userid", "postid") 
SELECT 'Текст каментария 30', users.id, posts.id
FROM users, posts
WHERE "name" = 'Пользователь_8' AND "text" = 'Текст поста 1';

INSERT INTO comments ("text", "userid", "postid", "perentid") 
SELECT 'Текст каментария 01', users.id, posts.id, comments.id
FROM users, posts, comments
WHERE "name" = 'Пользователь_8' AND posts.text = 'Текст поста 1' AND comments.text = 'Текст каментария 00';

INSERT INTO comments ("text", "userid", "postid", "perentid") 
SELECT 'Текст каментария 02', users.id, posts.id, comments.id
FROM users, posts, comments
WHERE "name" = 'Пользователь_10' AND posts.text = 'Текст поста 1' AND comments.text = 'Текст каментария 00';

INSERT INTO comments ("text", "userid", "postid", "perentid") 
SELECT 'Текст каментария 012', users.id, posts.id, comments.id
FROM users, posts, comments
WHERE "name" = 'Пользователь_12' AND posts.text = 'Текст поста 1' AND comments.text = 'Текст каментария 01';

INSERT INTO comments ("text", "userid", "postid", "perentid") 
SELECT 'Текст каментария 013', users.id, posts.id, comments.id
FROM users, posts, comments
WHERE "name" = 'Пользователь_14' AND posts.text = 'Текст поста 1' AND comments.text = 'Текст каментария 012';

INSERT INTO comments ("text", "userid", "postid", "perentid") 
SELECT 'Текст каментария 014', users.id, posts.id, comments.id
FROM users, posts, comments
WHERE "name" = 'Пользователь_16' AND posts.text = 'Текст поста 1' AND comments.text = 'Текст каментария 013';

INSERT INTO comments ("text", "userid", "postid", "perentid") 
SELECT 'Текст каментария 016', users.id, posts.id, comments.id
FROM users, posts, comments
WHERE "name" = 'Пользователь_6' AND posts.text = 'Текст поста 1' AND comments.text = 'Текст каментария 014';