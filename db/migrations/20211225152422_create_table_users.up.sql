CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "role" text NOT NULL DEFAULT 'customer'::text,
    "inserted_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX users_email_key ON users (email);