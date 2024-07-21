CREATE TABLE IF NOT EXISTS budgets
(
    "ID" serial NOT NULL,
    user_id uuid NOT NULL,
    budget_name character varying(255) NOT NULL,
    description character varying(255),
    amount double precision NOT NULL,
    category_id integer NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    period character varying(255) NOT NULL,
    PRIMARY KEY ("ID")
);
---- create above / drop below ----

DROP TABLE IF EXISTS budgets;