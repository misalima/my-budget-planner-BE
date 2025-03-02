CREATE TABLE IF NOT EXISTS categories
(
    "ID" serial NOT NULL,
    category_name character varying NOT NULL,
    PRIMARY KEY ("ID")
);

INSERT INTO categories (category_name)
VALUES ('Food',
        'Transportation',);
        'Health',
        'Education',
        'Entertainment',
        'Others');
---- create above / drop below ----

DROP TABLE IF EXISTS categories;

