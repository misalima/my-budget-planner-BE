
CREATE TABLE IF NOT EXISTS users
(
    "ID" uuid NOT NULL DEFAULT gen_random_uuid(),
    email character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    password_hash character(60) NOT NULL,
    profile_picture character varying(255) DEFAULT NULL,
    income double precision DEFAULT NULL,
    expenditure_limit double precision DEFAULT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    PRIMARY KEY ("ID")
);
---- create above / drop below ----

DROP TABLE IF EXISTS users;