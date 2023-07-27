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

<<<<<<< HEAD
=======
INSERT INTO public.movies VALUES (1, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (2, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (3, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (4, 'Barbie', 'Pink', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (5, 'Barbie', 'Pink', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (6, 'Barbie', 'Pink', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
INSERT INTO public.movies VALUES (7, 'Barbie', 'Pink', 8.8, 'jhhfhj', '2023-07-25 12:08:03.255262+00');
>>>>>>> bdc9054b09e04c55c212f7486af0867a9d235cff


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: adil
--

<<<<<<< HEAD
=======
INSERT INTO public.users VALUES (1, 'Ezel', '', '', 'ezelimo', '2023-07-26 10:25:29.029228+00');
>>>>>>> bdc9054b09e04c55c212f7486af0867a9d235cff


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: adil
--



--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: adil
--



--
-- Data for Name: watched_movie; Type: TABLE DATA; Schema: public; Owner: adil
--

INSERT INTO public.watched_movie VALUES (1, 2);
INSERT INTO public.watched_movie VALUES (1, 1);
INSERT INTO public.watched_movie VALUES (1, 4);
INSERT INTO public.watched_movie VALUES (1, 3);


--
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: adil
--

SELECT pg_catalog.setval('public.comments_id_seq', (SELECT MAX(id) FROM public.comments));

--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: adil
--

SELECT pg_catalog.setval('public.movies_id_seq', (SELECT MAX(id) FROM public.movies));

--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: adil
--

SELECT pg_catalog.setval('public.users_id_seq', (SELECT MAX(id) FROM public.users));

--
-- PostgreSQL database dump complete
--

