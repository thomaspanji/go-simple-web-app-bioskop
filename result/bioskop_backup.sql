--
-- PostgreSQL database dump
--

-- Dumped from database version 14.17 (Ubuntu 14.17-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.17 (Ubuntu 14.17-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: update_updated_at_column(); Type: FUNCTION; Schema: public; Owner: bioskop_admin
--

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_updated_at_column() OWNER TO bioskop_admin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: bioskop; Type: TABLE; Schema: public; Owner: bioskop_admin
--

CREATE TABLE public.bioskop (
    id integer NOT NULL,
    nama character varying(255) NOT NULL,
    lokasi character varying(255) NOT NULL,
    rating double precision NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bioskop OWNER TO bioskop_admin;

--
-- Name: bioskop_id_seq; Type: SEQUENCE; Schema: public; Owner: bioskop_admin
--

CREATE SEQUENCE public.bioskop_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bioskop_id_seq OWNER TO bioskop_admin;

--
-- Name: bioskop_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: bioskop_admin
--

ALTER SEQUENCE public.bioskop_id_seq OWNED BY public.bioskop.id;


--
-- Name: gorp_migrations; Type: TABLE; Schema: public; Owner: bioskop_admin
--

CREATE TABLE public.gorp_migrations (
    id text NOT NULL,
    applied_at timestamp with time zone
);


ALTER TABLE public.gorp_migrations OWNER TO bioskop_admin;

--
-- Name: bioskop id; Type: DEFAULT; Schema: public; Owner: bioskop_admin
--

ALTER TABLE ONLY public.bioskop ALTER COLUMN id SET DEFAULT nextval('public.bioskop_id_seq'::regclass);


--
-- Data for Name: bioskop; Type: TABLE DATA; Schema: public; Owner: bioskop_admin
--

COPY public.bioskop (id, nama, lokasi, rating, created_at, updated_at) FROM stdin;
1	Pondok Indah Mall 1 XXI	Jakarta	4.6	2025-04-11 11:12:00.130096+07	2025-04-11 11:17:22.794878+07
2	Empire XXI	Bandung	4.5	2025-04-11 11:15:00.942967+07	2025-04-11 11:18:18.311562+07
3	Cinépolis Istana Plaza	Bandung	4.4	2025-04-11 11:19:14.065761+07	2025-04-11 11:19:14.065761+07
4	CGV Vivo Mall Sentul	Bogor	4.7	2025-04-11 11:19:45.709905+07	2025-04-11 11:19:45.709905+07
\.


--
-- Data for Name: gorp_migrations; Type: TABLE DATA; Schema: public; Owner: bioskop_admin
--

COPY public.gorp_migrations (id, applied_at) FROM stdin;
001_initial_create_table.sql	2025-04-11 11:10:47.094297+07
\.


--
-- Name: bioskop_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bioskop_admin
--

SELECT pg_catalog.setval('public.bioskop_id_seq', 5, true);


--
-- Name: bioskop bioskop_pkey; Type: CONSTRAINT; Schema: public; Owner: bioskop_admin
--

ALTER TABLE ONLY public.bioskop
    ADD CONSTRAINT bioskop_pkey PRIMARY KEY (id);


--
-- Name: gorp_migrations gorp_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: bioskop_admin
--

ALTER TABLE ONLY public.gorp_migrations
    ADD CONSTRAINT gorp_migrations_pkey PRIMARY KEY (id);


--
-- Name: bioskop set_updated_at; Type: TRIGGER; Schema: public; Owner: bioskop_admin
--

CREATE TRIGGER set_updated_at BEFORE UPDATE ON public.bioskop FOR EACH ROW EXECUTE FUNCTION public.update_updated_at_column();


--
-- PostgreSQL database dump complete
--

