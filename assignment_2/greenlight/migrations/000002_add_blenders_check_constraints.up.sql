ALTER TABLE blenders ADD CONSTRAINT blenders_capacity_check CHECK (capacity >= 0);
ALTER TABLE blenders ADD CONSTRAINT blenders_year_check CHECK (year BETWEEN 1970 AND date_part('year', now()));
ALTER TABLE blenders ADD CONSTRAINT categories_length_check CHECK (array_length(categories, 1) BETWEEN 1 AND 5);