CREATE TABLE IF NOT EXISTS categories
(
    "ID" serial NOT NULL,
    category_name character varying NOT NULL,
    PRIMARY KEY ("ID")
);
---- create above / drop below ----

DROP TABLE IF EXISTS categories;

