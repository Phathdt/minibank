CREATE TABLE "accounts" (
    "id" SERIAL PRIMARY KEY,
    "user_id" int8,
    "bank_id" int8,
    "name" text NOT NULL,
    "balance" bigint,
    "inserted_at" timestamp(0) NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) NOT NULL DEFAULT now(),
    CONSTRAINT "accounts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
    CONSTRAINT "accounts_bank_id_fkey" FOREIGN KEY ("bank_id") REFERENCES "banks"("id") ON DELETE CASCADE
);
