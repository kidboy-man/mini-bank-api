CREATE TABLE IF NOT EXISTS "entries" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL REFERENCES "accounts" ("id"),
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "entries" ("account_id");
