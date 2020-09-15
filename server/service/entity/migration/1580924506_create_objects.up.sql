-- Entity
CREATE TABLE "entity" (
  "id" uuid PRIMARY KEY,
  "name" text NOT NULL,
  "strength" int NOT NULL DEFAULT 0,
  "dexterity" int NOT NULL DEFAULT 0,
  "intelligence" int NOT NULL DEFAULT 0,
  "attribute_points" int NOT NULL DEFAULT 0,
  "experience_points" bigint NOT NULL DEFAULT 0,
  "coins" bigint NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

-- Entity Instance
CREATE TABLE "entity_instance" (
  "id" uuid PRIMARY KEY,
  "entity_id" uuid NOT NULL,
  "health_points" int NOT NULL DEFAULT 0,
  "action_points" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

-- Entity Item
CREATE TABLE "entity_item" (
  "id" uuid PRIMARY KEY,
  "entity_id" uuid NOT NULL,
  "item_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

-- Entity Spell
CREATE TABLE "entity_spell" (
  "id" uuid PRIMARY KEY,
  "entity_id" uuid NOT NULL,
  "spell_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (current_timestamp),
  "updated_at" timestamp DEFAULT null,
  "deleted_at" timestamp DEFAULT null
);

ALTER TABLE "entity_instance" ADD CONSTRAINT "entity_instance_entity_id_fk" FOREIGN KEY ("entity_id") REFERENCES "entity" ("id");

ALTER TABLE "entity_item" ADD CONSTRAINT "entity_item_entity_id_fk" FOREIGN KEY ("entity_id") REFERENCES "entity" ("id");

ALTER TABLE "entity_spell" ADD CONSTRAINT "entity_spell_entity_id_fk" FOREIGN KEY ("entity_id") REFERENCES "entity" ("id");

CREATE UNIQUE INDEX ON "entity_instance" ("entity_id", "deleted_at");

COMMENT ON COLUMN "entity_instance"."health_points" IS 'Computed from attribute points';

COMMENT ON COLUMN "entity_instance"."action_points" IS 'Computed from attribute points';

COMMENT ON COLUMN "entity_item"."item_id" IS 'Remote "item" service reference';

COMMENT ON COLUMN "entity_spell"."spell_id" IS 'Remote "spell" service reference';
