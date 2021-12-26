CREATE TABLE "accounts" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "bank_id" bigint NOT NULL,
    "name" text NOT NULL,
    "balance" bigint NOT NULL DEFAULT 0,
    "created_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("bank_id") REFERENCES "banks" ("id");
