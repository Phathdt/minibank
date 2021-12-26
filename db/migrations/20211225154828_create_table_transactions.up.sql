CREATE TABLE "transactions" (
    "id" BIGSERIAL PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "amount" bigint NOT NULL DEFAULT 0,
    "transaction_type" text NOT NULL DEFAULT 'deposit'::text,
    "created_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now()
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
