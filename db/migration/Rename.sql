CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "mobile" bigint NOT NULL,
  "type" int NOT NULL
);

CREATE TABLE "Equipment" (
  "id" SERIAL PRIMARY KEY,
  "type" int NOT NULL,
  "seller" int NOT NULL,
  "cost" int NOT NULL,
  "description" varchar NOT NULL,
  "photo" bytea NOT NULL
);

CREATE TABLE "Seller" (
  "id" SERIAL PRIMARY KEY,
  "uid" int NOT NULL,
  "location" varchar NOT NULL
);

CREATE TABLE "Staff" (
  "id" SERIAL PRIMARY KEY,
  "cost" int NOT NULL,
  "location" varchar NOT NULL,
  "photo" bytea NOT NULL,
  "uid" int NOT NULL
);

CREATE TABLE "Booking" (
  "id" SERIAL PRIMARY KEY,
  "sid" int NOT NULL,
  "uid" int NOT NULL,
  "duration" int NOT NULL,
  "cost" bigint NOT NULL
);

CREATE TABLE "Rent" (
  "id" SERIAL PRIMARY KEY,
  "equipment" int NOT NULL,
  "buyer" int NOT NULL,
  "cost" int NOT NULL
);

ALTER TABLE "Equipment" ADD FOREIGN KEY ("seller") REFERENCES "Seller" ("id");

ALTER TABLE "Seller" ADD FOREIGN KEY ("uid") REFERENCES "users" ("id");

ALTER TABLE "Staff" ADD FOREIGN KEY ("uid") REFERENCES "users" ("id");

ALTER TABLE "Booking" ADD FOREIGN KEY ("sid") REFERENCES "Staff" ("id");

ALTER TABLE "Booking" ADD FOREIGN KEY ("uid") REFERENCES "users" ("id");

ALTER TABLE "Rent" ADD FOREIGN KEY ("equipment") REFERENCES "Equipment" ("id");

ALTER TABLE "Rent" ADD FOREIGN KEY ("buyer") REFERENCES "users" ("id");

CREATE INDEX ON "Equipment" ("seller");

CREATE INDEX ON "Equipment" ("type");

CREATE INDEX ON "Booking" ("sid");

CREATE INDEX ON "Booking" ("uid");

CREATE INDEX ON "Booking" ("sid", "uid");

CREATE INDEX ON "Rent" ("equipment");

CREATE INDEX ON "Rent" ("buyer");

CREATE INDEX ON "Rent" ("equipment", "buyer");
