ALTER TABLE refresh_tokens ALTER COLUMN token TYPE character varying(512);
---- create above / drop below ----
ALTER TABLE refresh_tokens ALTER COLUMN token TYPE character varying(255);
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
