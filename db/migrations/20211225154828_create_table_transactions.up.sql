CREATE TABLE "transactions" (
    "id" SERIAL PRIMARY KEY,
    "account_id" int8,
    "amount" bigint,
    "type" text NOT NULL DEFAULT 'deposit'::text,
    "inserted_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now(),
    CONSTRAINT "transactions_account_id_fkey" FOREIGN KEY ("account_id") REFERENCES "accounts"("id") ON DELETE CASCADE
);
