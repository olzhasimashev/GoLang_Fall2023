CREATE TABLE IF NOT EXISTS blenders (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    year integer NOT NULL,
    capacity integer NOT NULL,
    material text NOT NULL,
    categories text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
);