-- +goose Up
-- modify "accounts" table
ALTER TABLE "accounts" ALTER COLUMN "account_id" DROP NOT NULL;

-- +goose Down
-- reverse: modify "accounts" table
ALTER TABLE "accounts" ALTER COLUMN "account_id" SET NOT NULL;
