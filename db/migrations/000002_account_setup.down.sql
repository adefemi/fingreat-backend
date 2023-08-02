DROP TABLE IF EXISTS "transfers";
DROP TABLE IF EXISTS "entries";
ALTER TABLE "accounts" DROP CONSTRAINT "unique_user_currency";
DROP TABLE IF EXISTS "accounts";