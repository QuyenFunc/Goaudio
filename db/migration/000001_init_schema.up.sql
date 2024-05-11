CREATE TABLE "admin" (
                         "id" SERIAL PRIMARY KEY,
                         "username" VARCHAR(32) NOT NULL DEFAULT '',
                         "password" VARCHAR(32) NOT NULL DEFAULT '',
                         "name" VARCHAR(128) NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);          

CREATE TABLE "users" (
                         "id" SERIAL PRIMARY KEY,
                         "name" VARCHAR(50) NOT NULL,
                         "email" VARCHAR(50) NOT NULL,
                         "phone" VARCHAR(15) NOT NULL,
                         "address" VARCHAR(128) NOT NULL,
                         "password" VARCHAR(40) NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "cataloges" (
                                           "id" SERIAL PRIMARY KEY,
                                           "name" VARCHAR(128) NOT NULL,
                                           "parent_id" INT NOT NULL DEFAULT '0',
                                           "sort_order" SMALLINT NOT NULL DEFAULT '0'
);

CREATE TABLE "productes" (
                             "id" SERIAL PRIMARY KEY,
                             "catalog_id" INT NOT NULL,
                             "name" VARCHAR(100) NOT NULL,
                             "price" INT NOT NULL DEFAULT '0',
                             "view" INT NOT NULL DEFAULT '0',
                             "content" TEXT NOT NULL,
                             "discount" VARCHAR(10) NOT NULL,
                             "image_link" VARCHAR(50) NOT NULL,
                             "image_list" TEXT NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);
