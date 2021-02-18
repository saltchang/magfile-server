CREATE TABLE "blog_user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "current_location" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "logined_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "token" (
  "id" bigserial,
  "user_id" bigint,
  "access_token" varchar NOT NULL,
  "expired_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id", "user_id")
);

CREATE TABLE "post" (
  "id" bigserial PRIMARY KEY,
  "semantic_id" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  "series_id" bigint,
  "order_in_series" int,
  "title" varchar NOT NULL,
  "abstract" varchar NOT NULL DEFAULT '',
  "content" varchar NOT NULL DEFAULT '',
  "views" bigint NOT NULL DEFAULT 0,
  "is_archived" boolean NOT NULL DEFAULT false,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tag" (
  "id" bigserial PRIMARY KEY,
  "author_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_tag" (
  "id" bigserial PRIMARY KEY,
  "post_id" bigint NOT NULL,
  "tag_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "series" (
  "id" bigserial PRIMARY KEY,
  "semantic_id" varchar NOT NULL,
  "author_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "abstract" varchar NOT NULL,
  "is_archived" boolean NOT NULL DEFAULT false,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "token" ADD FOREIGN KEY ("user_id") REFERENCES "blog_user" ("id");

ALTER TABLE "post" ADD FOREIGN KEY ("author_id") REFERENCES "blog_user" ("id");

ALTER TABLE "post" ADD FOREIGN KEY ("series_id") REFERENCES "series" ("id");

ALTER TABLE "tag" ADD FOREIGN KEY ("author_id") REFERENCES "blog_user" ("id");

ALTER TABLE "post_tag" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("id");

ALTER TABLE "post_tag" ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id");

ALTER TABLE "series" ADD FOREIGN KEY ("author_id") REFERENCES "blog_user" ("id");

CREATE INDEX ON "blog_user" ("username");

CREATE INDEX ON "blog_user" ("email");

CREATE INDEX ON "token" ("user_id");

CREATE INDEX ON "post" ("author_id");

CREATE INDEX ON "post" ("semantic_id");

CREATE INDEX ON "tag" ("name");

CREATE INDEX ON "post_tag" ("post_id");

CREATE INDEX ON "series" ("author_id");

CREATE INDEX ON "series" ("semantic_id");

COMMENT ON COLUMN "blog_user"."username" IS 'can be change';

COMMENT ON COLUMN "blog_user"."email" IS 'can be change';

COMMENT ON COLUMN "blog_user"."gender" IS 'male, female, other';

COMMENT ON COLUMN "blog_user"."password_hash" IS 'password hashed with SHA-512';

COMMENT ON COLUMN "token"."access_token" IS 'random string hashed with SHA-512';
