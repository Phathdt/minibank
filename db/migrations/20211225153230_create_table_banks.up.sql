CREATE TABLE "banks" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "created_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX ON "banks" ("name");
