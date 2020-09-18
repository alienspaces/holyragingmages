-- account
CREATE TYPE "provider" AS ENUM (
  'google',
  'apple',
  'facebook',
  'twitter',
  'github'
);

CREATE TABLE "account" (
  "id" uuid CONSTRAINT account_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" text NOT NULL,
  "email" text NOT NULL,
  "provider" provider NOT NULL,
  "provider_account_id" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

CREATE TABLE "account_entity" (
  "id" uuid CONSTRAINT account_entity_pk PRIMARY KEY DEFAULT gen_random_uuid(),
  "account_id" uuid NOT NULL,
  "entity_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "account_entity" ADD CONSTRAINT "account_entity_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "account" ("id");

COMMENT ON COLUMN "account_entity"."entity_id" IS 'Remote "entity" service reference';
