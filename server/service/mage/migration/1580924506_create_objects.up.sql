-- mage
CREATE TABLE mage (
    id                UUID CONSTRAINT mage_pk PRIMARY KEY DEFAULT gen_random_uuid(),
    "name"            TEXT NOT NULL,
    strength          INTEGER NOT NULL DEFAULT 0,
    dexterity         INTEGER NOT NULL DEFAULT 0,
    intelligence      INTEGER NOT NULL DEFAULT 0,
    attribute_points  BIGINT NOT NULL DEFAULT 0,
    experience_points BIGINT NOT NULL DEFAULT 0,
    coins             BIGINT NOT NULL DEFAULT 0,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    updated_at        TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at        TIMESTAMP WITH TIME ZONE DEFAULT NULL
);