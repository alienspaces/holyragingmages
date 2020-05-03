-- item
CREATE TABLE item (
    id         UUID CONSTRAINT item_pk PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamp with time zone DEFAULT current_timestamp,
    updated_at timestamp with time zone DEFAULT NULL,
    deleted_at timestamp with time zone DEFAULT NULL
);