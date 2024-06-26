CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "role_id" bigint,
  "username" varchar NOT NULL,
  "email" varchar,
  "password_hash" varchar,
  "last_login" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "roles" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "active" bool DEFAULT true,
  "description" varchar(500),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "permissions" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "active" bool DEFAULT true,
  "description" varchar(500),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "role_permission" (
  "role_id" bigint NOT NULL,
  "permission_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "account_permission" (
  "account_id" bigint NOT NULL,
  "permission_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "sessions" (
  "id" bigserial PRIMARY KEY,
  "key" varchar NOT NULL,
  "value" varchar NOT NULL,
  "ip" varchar,
  "extend_data" jsonb,
  "exprired_at" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "roles" ("slug");

CREATE INDEX ON "permissions" ("slug");

CREATE INDEX ON "sessions" ("key");

COMMENT ON COLUMN "sessions"."exprired_at" IS 'Expired time in seconds';

ALTER TABLE "accounts" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "account_permission" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "account_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("id") REFERENCES "accounts" ("id");
