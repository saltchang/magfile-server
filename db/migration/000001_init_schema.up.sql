CREATE TABLE "blog_user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" varchar,
  "location" varchar,
  "password_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "logined_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post" (
  "id" bigserial PRIMARY KEY,
  "semantic_id" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar,
  "content" varchar,
  "tags" varchar[],
  "archived" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "post" ADD FOREIGN KEY ("author_id") REFERENCES "blog_user" ("id");

CREATE INDEX ON "blog_user" ("username");

CREATE INDEX ON "blog_user" ("email");

CREATE INDEX ON "post" ("author_id");

CREATE INDEX ON "post" ("semantic_id");

COMMENT ON COLUMN "blog_user"."username" IS 'can be change';

COMMENT ON COLUMN "blog_user"."email" IS 'can be change';

COMMENT ON COLUMN "blog_user"."gender" IS 'male, female, other';

COMMENT ON COLUMN "blog_user"."password_hash" IS 'password hashed with SHA-512';
