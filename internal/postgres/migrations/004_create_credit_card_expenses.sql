CREATE TABLE IF NOT EXISTS credit_card_expenses
(
    expense_id uuid NOT NULL,
    card_id uuid NOT NULL,
    installment_amount double precision NOT NULL,
    installments_number integer NOT NULL,
    CONSTRAINT expense_id PRIMARY KEY (expense_id)
);
---- create above / drop below ----

DROP TABLE IF EXISTS credit_card_expenses;