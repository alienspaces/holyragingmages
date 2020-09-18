-- item
CREATE TABLE "item" (
    "id"          UUID CONSTRAINT item_pk PRIMARY KEY DEFAULT gen_random_uuid(),
    "name"        TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "created_at"  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at"  TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    "deleted_at"  TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
