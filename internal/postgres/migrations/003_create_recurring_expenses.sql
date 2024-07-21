CREATE TABLE IF NOT EXISTS recurring_expenses
(
    expense_id uuid NOT NULL,
    card_id uuid,
    start_date date NOT NULL,
    end_date date,
    frequency character varying NOT NULL,
    PRIMARY KEY (expense_id)
);
---- create above / drop below ----

DROP TABLE IF EXISTS recurring_expenses;
