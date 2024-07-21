CREATE TABLE IF NOT EXISTS expenses
(
    "ID" uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    category_id serial NOT NULL,
    amount double precision NOT NULL,
    description character varying(255),
    date date NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY ("ID")
);
---- create above / drop below ----

DROP TABLE IF EXISTS expenses