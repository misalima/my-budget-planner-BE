CREATE TABLE IF NOT EXISTS credit_cards
(
    "ID" uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    card_name character varying NOT NULL,
    total_limit double precision NOT NULL,
    current_limit double precision,
    due_date integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY ("ID")
);

---- create above / drop below ----

DROP TABLE IF EXISTS credit_cards;
