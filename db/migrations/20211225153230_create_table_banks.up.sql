CREATE TABLE "banks" (
    "id" SERIAL PRIMARY KEY,
    "name" text NOT NULL,
    "inserted_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX banks_name_key ON banks (name);
