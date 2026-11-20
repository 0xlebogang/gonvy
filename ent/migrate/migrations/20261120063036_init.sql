-- +goose Up
-- create "users" table
CREATE TABLE "users" ("id" character varying NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "name" character varying NOT NULL, "email" character varying NOT NULL, "email_verified" boolean NOT NULL DEFAULT false, "image" text NULL, PRIMARY KEY ("id"));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- create "verifications" table
CREATE TABLE "verifications" ("id" character varying NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "identifier" character varying NOT NULL, "value" character varying NOT NULL, "expires_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create index "verification_identifier" to table: "verifications"
CREATE INDEX "verification_identifier" ON "verifications" ("identifier");
-- create "accounts" table
CREATE TABLE "accounts" ("id" character varying NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "account_id" character varying NOT NULL, "provider_id" character varying NOT NULL, "access_token" character varying NULL, "refresh_token" character varying NULL, "access_token_expires_at" timestamptz NULL, "refresh_token_expires_at" timestamptz NULL, "scope" character varying NULL, "id_token" character varying NULL, "password" text NULL, "user_accounts" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "accounts_users_accounts" FOREIGN KEY ("user_accounts") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "accounts_account_id_key" to table: "accounts"
CREATE UNIQUE INDEX "accounts_account_id_key" ON "accounts" ("account_id");
-- create "sessions" table
CREATE TABLE "sessions" ("id" character varying NOT NULL, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "token" character varying NOT NULL, "expires_at" timestamptz NOT NULL, "ip_address" character varying NULL, "user_agent" character varying NULL, "user_sessions" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "sessions_users_sessions" FOREIGN KEY ("user_sessions") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- create index "sessions_token_key" to table: "sessions"
CREATE UNIQUE INDEX "sessions_token_key" ON "sessions" ("token");

-- +goose Down
-- reverse: create index "sessions_token_key" to table: "sessions"
DROP INDEX "sessions_token_key";
-- reverse: create "sessions" table
DROP TABLE "sessions";
-- reverse: create index "accounts_account_id_key" to table: "accounts"
DROP INDEX "accounts_account_id_key";
-- reverse: create "accounts" table
DROP TABLE "accounts";
-- reverse: create index "verification_identifier" to table: "verifications"
DROP INDEX "verification_identifier";
-- reverse: create "verifications" table
DROP TABLE "verifications";
-- reverse: create index "users_email_key" to table: "users"
DROP INDEX "users_email_key";
-- reverse: create "users" table
DROP TABLE "users";
