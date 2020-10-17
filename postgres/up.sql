
DROP TABLE IF EXISTS public.ads;

CREATE TABLE public.ads
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    "number" character varying(200) COLLATE pg_catalog."default" NOT NULL,
    price character varying(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT urls_pkey PRIMARY KEY (id),
    CONSTRAINT urls_key UNIQUE ("number")
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.ads
    OWNER to admin;

DROP TABLE IF EXISTS public.subscription;

CREATE TABLE public.subscription
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    userid integer NOT NULL,
    adid integer NOT NULL,
    CONSTRAINT subscription_pkey PRIMARY KEY (id),
    CONSTRAINT subscription_key UNIQUE (userid, adid)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.subscription
    OWNER to admin;

DROP TABLE IF EXISTS public.users;

CREATE TABLE public.users
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    email character varying(200) COLLATE pg_catalog."default" NOT NULL,
    activation boolean NOT NULL DEFAULT false,
    key character varying(15) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to admin;