ALTER TABLE IF EXISTS public.expenses
    ADD CONSTRAINT user_id FOREIGN KEY (user_id)
        REFERENCES public.users ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.expenses
    ADD CONSTRAINT category_id FOREIGN KEY (category_id)
        REFERENCES public.categories ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.recurring_expenses
    ADD CONSTRAINT expense_id FOREIGN KEY (expense_id)
        REFERENCES public.expenses ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.recurring_expenses
    ADD CONSTRAINT card_id FOREIGN KEY (card_id)
        REFERENCES public.credit_cards ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.credit_card_expenses
    ADD CONSTRAINT card_expense_id FOREIGN KEY (expense_id)
        REFERENCES public.expenses ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.credit_card_expenses
    ADD CONSTRAINT card_id FOREIGN KEY (card_id)
        REFERENCES public.credit_cards ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.credit_cards
    ADD CONSTRAINT user_id FOREIGN KEY (user_id)
        REFERENCES public.users ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.budgets
    ADD CONSTRAINT user_id FOREIGN KEY (user_id)
        REFERENCES public.users ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;


ALTER TABLE IF EXISTS public.budgets
    ADD CONSTRAINT category_id FOREIGN KEY (category_id)
        REFERENCES public.categories ("ID") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID;

---- create above / drop below ----

-- Reverse for adding user_id foreign key constraint on public.expenses
ALTER TABLE IF EXISTS public.expenses
    DROP CONSTRAINT IF EXISTS user_id;

-- Reverse for adding category_id foreign key constraint on public.expenses
ALTER TABLE IF EXISTS public.expenses
    DROP CONSTRAINT IF EXISTS category_id;

-- Reverse for adding expense_id foreign key constraint on public.recurring_expenses
ALTER TABLE IF EXISTS public.recurring_expenses
    DROP CONSTRAINT IF EXISTS expense_id;

-- Reverse for adding card_id foreign key constraint on public.recurring_expenses
ALTER TABLE IF EXISTS public.recurring_expenses
    DROP CONSTRAINT IF EXISTS card_id;

-- Reverse for adding expense_id foreign key constraint on public.credit_card_expenses
ALTER TABLE IF EXISTS public.credit_card_expenses
    DROP CONSTRAINT IF EXISTS expense_id;

-- Reverse for adding card_id foreign key constraint on public.credit_card_expenses
ALTER TABLE IF EXISTS public.credit_card_expenses
    DROP CONSTRAINT IF EXISTS card_id;

-- Reverse for adding user_id foreign key constraint on public.credit_cards
ALTER TABLE IF EXISTS public.credit_cards
    DROP CONSTRAINT IF EXISTS user_id;

-- Reverse for adding user_id foreign key constraint on public.budgets
ALTER TABLE IF EXISTS public.budgets
    DROP CONSTRAINT IF EXISTS user_id;

-- Reverse for adding category_id foreign key constraint on public.budgets
ALTER TABLE IF EXISTS public.budgets
    DROP CONSTRAINT IF EXISTS category_id;
