-- mage
CREATE TABLE mage (
    id         UUID CONSTRAINT mage_pk PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT current_timestamp,
    updated_at timestamp with time zone DEFAULT NULL,
    deleted_at timestamp with time zone DEFAULT NULL
);