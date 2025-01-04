-- Write your migrate up statements here

ALTER TABLE users ALTER COLUMN profile_picture SET DEFAULT '';
UPDATE users SET profile_picture = '' WHERE profile_picture IS NULL;
ALTER TABLE users ALTER COLUMN profile_picture SET NOT NULL;

ALTER TABLE users ALTER COLUMN income SET DEFAULT 0;
UPDATE users SET income = 0 WHERE income IS NULL;
ALTER TABLE users ALTER COLUMN income SET NOT NULL;

ALTER TABLE users ALTER COLUMN expenditure_limit SET DEFAULT 0;
UPDATE users SET expenditure_limit = 0 WHERE expenditure_limit IS NULL;
ALTER TABLE users ALTER COLUMN expenditure_limit SET NOT NULL;

ALTER TABLE users ALTER COLUMN created_at SET DEFAULT NOW();
ALTER TABLE users ALTER COLUMN updated_at SET DEFAULT NOW();


---- create above / drop below ----

-- Remove DEFAULT and NOT NULL from profile_picture
ALTER TABLE users ALTER COLUMN profile_picture DROP DEFAULT;
ALTER TABLE users ALTER COLUMN profile_picture DROP NOT NULL;

-- Remove DEFAULT and NOT NULL from income
ALTER TABLE users ALTER COLUMN income DROP DEFAULT;
ALTER TABLE users ALTER COLUMN income DROP NOT NULL;

-- Remove DEFAULT and NOT NULL from expenditure_limit
ALTER TABLE users ALTER COLUMN expenditure_limit DROP DEFAULT;
ALTER TABLE users ALTER COLUMN expenditure_limit DROP NOT NULL;

-- Remove DEFAULT from created_at and updated_at
ALTER TABLE users ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE users ALTER COLUMN updated_at DROP DEFAULT;
