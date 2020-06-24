-- mage
CREATE TABLE mage (
    id           UUID CONSTRAINT mage_pk PRIMARY KEY DEFAULT gen_random_uuid(),
    "name"       TEXT NOT NULL,
    strength     INTEGER NOT NULL DEFAULT 0,
    dexterity    INTEGER NOT NULL DEFAULT 0,
    intelligence INTEGER NOT NULL DEFAULT 0,
    experience   BIGINT NOT NULL DEFAULT 0,
    coin         BIGINT NOT NULL DEFAULT 0,
    created_at   timestamp with time zone DEFAULT current_timestamp,
    updated_at   timestamp with time zone DEFAULT NULL,
    deleted_at   timestamp with time zone DEFAULT NULL
);