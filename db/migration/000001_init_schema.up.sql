CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_tokens" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "token" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_data" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "full_name" varchar NOT NULL,
  "sex" varchar,
  "location" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "semantic_id" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar,
  "tags" varchar[],
  "content" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "user_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_data" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "user_tokens" ("user_id");

CREATE INDEX ON "user_data" ("user_id");

CREATE INDEX ON "posts" ("author_id");

CREATE INDEX ON "posts" ("semantic_id");
