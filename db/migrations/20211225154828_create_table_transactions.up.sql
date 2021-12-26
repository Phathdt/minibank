CREATE TABLE "transactions" (
    "id" BIGSERIAL PRIMARY KEY,
    "from_account_id" bigint NOT NULL,
    "to_account_id" bigint NOT NULL,
    "amount" bigint NOT NULL DEFAULT 0,
    "type" text NOT NULL DEFAULT 'deposit'::text,
    "created_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
