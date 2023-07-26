CREATE TABLE "movies" (
  "id" bigserial PRIMARY KEY,
  "title" text NOT NULL,
  "description" text NOT NULL,
  "score" float NOT NULL,
  "image" text NOT NULL,
  "release_date" timestamptz NOT NULL
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "image" text NOT NULL,
  "password_hash" text NOT NULL,
  "email" text UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "watched_movie" (
  "user_id" bigint NOT NULL,
  "movie_id" bigint NOT NULL
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "movie_id" bigint NOT NULL,
  "content" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "watched_movie" ("user_id", "movie_id");

ALTER TABLE "watched_movie" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "watched_movie" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
