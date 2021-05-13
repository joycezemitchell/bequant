--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.20
-- Dumped by pg_dump version 13.1

-- Started on 2021-05-13 13:57:37

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

--
-- TOC entry 185 (class 1259 OID 16476)
-- Name: display; Type: TABLE; Schema: public; Owner: mitch
--

CREATE TABLE public.display (
    change24hour character varying(100),
    changepct24hour character varying(100),
    open24hour character varying(100),
    volume24hour character varying(100),
    volume24hourto character varying(100),
    low24hour character varying(100),
    high24hour character varying(100),
    price character varying(100),
    lastupdate character varying(100),
    mktcap character varying(100),
    currency character varying(50),
    id integer NOT NULL,
    ptype character varying(100),
    supply character varying(100)
);


ALTER TABLE public.display OWNER TO mitch;

--
-- TOC entry 186 (class 1259 OID 16482)
-- Name: display_id_seq; Type: SEQUENCE; Schema: public; Owner: mitch
--

CREATE SEQUENCE public.display_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.display_id_seq OWNER TO mitch;

--
-- TOC entry 2768 (class 0 OID 0)
-- Dependencies: 186
-- Name: display_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mitch
--

ALTER SEQUENCE public.display_id_seq OWNED BY public.display.id;


--
-- TOC entry 187 (class 1259 OID 16493)
-- Name: raw; Type: TABLE; Schema: public; Owner: mitch
--

CREATE TABLE public.raw (
    change24hour numeric(100,14),
    changepct24hour numeric(100,14),
    open24hour numeric(100,14),
    volume24hour numeric(100,14),
    volume24hourto numeric(100,14),
    low24hour numeric(100,14),
    high24hour numeric(100,14),
    price numeric(100,14),
    lastupdate integer,
    supply integer,
    mktcap numeric(100,14),
    currency character varying(100),
    ptype character varying(100),
    id integer NOT NULL
);


ALTER TABLE public.raw OWNER TO mitch;

--
-- TOC entry 188 (class 1259 OID 16499)
-- Name: raw_id_seq; Type: SEQUENCE; Schema: public; Owner: mitch
--

CREATE SEQUENCE public.raw_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.raw_id_seq OWNER TO mitch;

--
-- TOC entry 2769 (class 0 OID 0)
-- Dependencies: 188
-- Name: raw_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: mitch
--

ALTER SEQUENCE public.raw_id_seq OWNED BY public.raw.id;


--
-- TOC entry 2636 (class 2604 OID 16484)
-- Name: display id; Type: DEFAULT; Schema: public; Owner: mitch
--

ALTER TABLE ONLY public.display ALTER COLUMN id SET DEFAULT nextval('public.display_id_seq'::regclass);


--
-- TOC entry 2637 (class 2604 OID 16501)
-- Name: raw id; Type: DEFAULT; Schema: public; Owner: mitch
--

ALTER TABLE ONLY public.raw ALTER COLUMN id SET DEFAULT nextval('public.raw_id_seq'::regclass);


--
-- TOC entry 2759 (class 0 OID 16476)
-- Dependencies: 185
-- Data for Name: display; Type: TABLE DATA; Schema: public; Owner: mitch
--

COPY public.display (change24hour, changepct24hour, open24hour, volume24hour, volume24hourto, low24hour, high24hour, price, lastupdate, mktcap, currency, id, ptype, supply) FROM stdin;
$ -7,777.10	-13.45	$ 57,818.6	Ƀ 97,225.5	$ 5,101,094,057.7	$ 45,871.7	$ 58,011.6	$ 50,041.5	Just now	$ 936.21 B	USD	34	BTC	Ƀ 18,708,625.0
€ -6,175.79	-12.95	€ 47,691.8	Ƀ 26,933.9	€ 1,173,468,731.0	€ 38,665.4	€ 47,847.6	€ 41,516.0	Just now	€ 776.71 B	EUR	35	BTC	Ƀ 18,708,625.0
$ -6.14	-12.35	$ 49.73	LINK 7,596,347.8	$ 341,549,093.3	$ 39.30	$ 49.81	$ 43.59	Just now	$ 43.59 B	USD	36	LINK	LINK 1,000,000,000.0
€ -4.88	-11.91	€ 40.99	LINK 1,112,376.8	€ 41,544,979.2	€ 32.92	€ 41.05	€ 36.11	Just now	€ 36.11 B	EUR	37	LINK	LINK 1,000,000,000.0
\.


--
-- TOC entry 2761 (class 0 OID 16493)
-- Dependencies: 187
-- Data for Name: raw; Type: TABLE DATA; Schema: public; Owner: mitch
--

COPY public.raw (change24hour, changepct24hour, open24hour, volume24hour, volume24hourto, low24hour, high24hour, price, lastupdate, supply, mktcap, currency, ptype, id) FROM stdin;
-7777.09999999999850	-13.45085252783531	57818.64000000000000	97225.47488755000000	5101094057.71541500000000	45871.69000000000000	58011.60000000000000	50041.54000000000000	1620876350	18708625	936208406282.50000000000000	USD	BTC	51
-6175.79000000000100	-12.94936679930294	47691.83000000000000	26933.90796927000000	1173468731.03863760000000	38665.35000000000000	47847.64000000000000	41516.04000000000000	1620876348	18708625	776708023845.00000000000000	EUR	BTC	52
-6.13999999999999	-12.34667202895635	49.73000000000000	7596347.81818117900000	341549093.28405917000000	39.30000000000000	49.81000000000000	43.59000000000000	1620876351	1000000000	43590000000.00000000000000	USD	LINK	53
-4.88000000000000	-11.90534276652843	40.99000000000000	1112376.80598850000000	41544979.19783466000000	32.92000000000000	41.05000000000000	36.11000000000000	1620876333	1000000000	36110000000.00000000000000	EUR	LINK	54
\.


--
-- TOC entry 2770 (class 0 OID 0)
-- Dependencies: 186
-- Name: display_id_seq; Type: SEQUENCE SET; Schema: public; Owner: mitch
--

SELECT pg_catalog.setval('public.display_id_seq', 37, true);


--
-- TOC entry 2771 (class 0 OID 0)
-- Dependencies: 188
-- Name: raw_id_seq; Type: SEQUENCE SET; Schema: public; Owner: mitch
--

SELECT pg_catalog.setval('public.raw_id_seq', 54, true);


--
-- TOC entry 2639 (class 2606 OID 16492)
-- Name: display display_pkey; Type: CONSTRAINT; Schema: public; Owner: mitch
--

ALTER TABLE ONLY public.display
    ADD CONSTRAINT display_pkey PRIMARY KEY (id);


--
-- TOC entry 2641 (class 2606 OID 16506)
-- Name: raw raw_pkey; Type: CONSTRAINT; Schema: public; Owner: mitch
--

ALTER TABLE ONLY public.raw
    ADD CONSTRAINT raw_pkey PRIMARY KEY (id);


-- Completed on 2021-05-13 13:57:42

--
-- PostgreSQL database dump complete
--

