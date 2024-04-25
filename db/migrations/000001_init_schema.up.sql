CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "role_id" bigint,
  "username" varchar NOT NULL,
  "email" varchar,
  "password_hash" varchar,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "active" bool DEFAULT true,
  "description" tinytext,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz
);

CREATE TABLE "permissions" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "active" bool DEFAULT true,
  "description" tinytext,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz
);

CREATE TABLE "role_permission" (
  "role_id" bigint,
  "permission_id" bigint,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz
);

CREATE TABLE "account_permission" (
  "account_id" bigint,
  "permission_id" bigint,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestampz NOT NULL DEFAULT (now()),
  "deleted_at" timestampz
);

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "roles" ("slug");

CREATE INDEX ON "permissions" ("slug");

ALTER TABLE "accounts" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "account_permission" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "account_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");
