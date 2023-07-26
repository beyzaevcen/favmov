--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: adil
--

INSERT INTO public.movies VALUES (1, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (2, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (3, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: adil
--

INSERT INTO public.users VALUES (1, 'Ezel', '', '', 'ezelimo', '2023-07-26 10:25:29.029228+00');


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: adil
--



--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: adil
--

INSERT INTO public.schema_migrations VALUES (2, false);


--
-- Data for Name: watched_movie; Type: TABLE DATA; Schema: public; Owner: adil
--



--
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: adil
--

SELECT pg_catalog.setval('public.comments_id_seq', 1, false);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: adil
--

SELECT pg_catalog.setval('public.movies_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: adil
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


--
-- PostgreSQL database dump complete
--
