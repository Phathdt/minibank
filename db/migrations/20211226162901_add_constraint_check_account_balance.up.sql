ALTER TABLE "accounts"
ADD CONSTRAINT balance_check
CHECK ( balance >= 0);
