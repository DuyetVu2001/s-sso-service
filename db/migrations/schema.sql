CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "avatar" varchar,
  "username" varchar,
  "full_name" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "follows" (
  "following_user_id" integer,
  "followed_user_id" integer,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "desc" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "content" text,
  "image" varchar,
  "user_id" integer,
  "category_id" integer,
  "status" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "favorites" (
  "post_id" bigint,
  "user_id" bigint,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "content" varchar,
  "post_id" integer,
  "user_id" integer,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "ratings" (
  "rate" int,
  "post_id" bigint,
  "user_id" bigint,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "notifications" (
  "id" bigserial PRIMARY KEY,
  "title" varchar,
  "content" varchar,
  "reference_id" bigint,
  "user_id" bigint,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "posts" ("title");

ALTER TABLE
  "follows"
ADD
  FOREIGN KEY ("following_user_id") REFERENCES "users" ("id");

ALTER TABLE
  "follows"
ADD
  FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id");

ALTER TABLE
  "posts"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE
  "posts"
ADD
  FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE
  "favorites"
ADD
  FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE
  "favorites"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE
  "comments"
ADD
  FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE
  "comments"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE
  "ratings"
ADD
  FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE
  "ratings"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE
  "notifications"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE authors (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  bio text
);