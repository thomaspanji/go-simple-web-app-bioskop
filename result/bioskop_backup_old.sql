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
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
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
-- Name: bioskop id; Type: DEFAULT; Schema: public; Owner: bioskop_admin
--

ALTER TABLE ONLY public.bioskop ALTER COLUMN id SET DEFAULT nextval('public.bioskop_id_seq'::regclass);


--
-- Data for Name: bioskop; Type: TABLE DATA; Schema: public; Owner: bioskop_admin
--

COPY public.bioskop (id, nama, lokasi, rating, created_at, updated_at) FROM stdin;
1	Empire XXI	Bandung	4.5	2025-04-10 02:14:51.917425	2025-04-10 02:14:51.917425
4	CGV Miko Mall	Bandung	4.4	2025-04-10 09:14:58.5246	2025-04-10 09:14:58.5246
5	CGV Vivo Mall Sentul	Bogor	4.7	2025-04-10 09:16:20.563949	2025-04-10 09:16:20.563949
6	Cinépolis Istana Plaza	Bandung	4.4	2025-04-10 09:19:46.461261	2025-04-10 09:19:46.461261
7	Cinépolis Q Mall Banjarbaru	Banjarbaru	4.5	2025-04-10 09:21:42.445754	2025-04-10 09:21:42.445754
8	Metropole XXI	Jakarta	4.6	2025-04-10 09:40:24.64648	2025-04-10 09:40:24.64648
9	 Pondok Indah Mall 1 XXI    	Jakarta 	4.6	2025-04-10 09:49:41.39277	2025-04-10 09:49:41.39277
\.


--
-- Name: bioskop_id_seq; Type: SEQUENCE SET; Schema: public; Owner: bioskop_admin
--

SELECT pg_catalog.setval('public.bioskop_id_seq', 9, true);


--
-- Name: bioskop bioskop_pkey; Type: CONSTRAINT; Schema: public; Owner: bioskop_admin
--

ALTER TABLE ONLY public.bioskop
    ADD CONSTRAINT bioskop_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

