CREATE INDEX IF NOT EXISTS blenders_name_idx ON blenders USING GIN (to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS blenders_categories_idx ON blenders USING GIN (categories);
