CREATE TABLE "admin" (
                         "id" bigserial PRIMARY KEY,
                         "username" VARCHAR(32) NOT NULL DEFAULT '',
                         "password" VARCHAR(32) NOT NULL DEFAULT '',
                         "name" VARCHAR(128) NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);          

CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "name" VARCHAR(50) NOT NULL,
                         "email" VARCHAR(50) NOT NULL,
                         "phone" VARCHAR(15) NOT NULL,
                         "address" VARCHAR(128) NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "cataloges" (
                                           "id" bigserial PRIMARY KEY,
                                           "name" VARCHAR(128) NOT NULL,
                                           "parent_id" bigint NOT NULL DEFAULT '0',
                                           "sort_order" SMALLINT NOT NULL DEFAULT '0'
);

CREATE TABLE "productes" (
                             "id" bigserial PRIMARY KEY,
                             "catalog_id" bigint NOT NULL,
                             "name" VARCHAR(100) NOT NULL,
                             "price" bigint NOT NULL DEFAULT '0',
                             "view" bigint NOT NULL DEFAULT '0',
                             "content" TEXT NOT NULL,
                             "discount" VARCHAR(10) NOT NULL,
                             "image_link" VARCHAR(50) NOT NULL,
                             "image_list" TEXT NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "admin" ("username");

CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("phone");

CREATE INDEX ON "productes" ("price");
CREATE INDEX ON "productes" ("name");
CREATE INDEX ON "productes" ("discount");
