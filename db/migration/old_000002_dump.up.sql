--
-- PostgreSQL database dump
--

-- Dumped from database version 14.8 (Homebrew)
-- Dumped by pg_dump version 14.8 (Homebrew)

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

INSERT INTO public.movies VALUES (1, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 15:08:03.255262+03');
INSERT INTO public.movies VALUES (2, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 15:08:03.255262+03');
INSERT INTO public.movies VALUES (3, 'The Fault in Our Stars', 'Cancer', 8.8, 'jhhfhj', '2023-07-25 15:08:03.255262+03');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: adil
--

INSERT INTO public.users VALUES (2, 'Ezel', 'dsfjkhs', 'gfgh', 'ezelimo', '2023-07-26 12:11:43.348112+03');


--
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: adil
--

INSERT INTO public.comments VALUES (1, 2, 3, 'selam', '2023-07-26 12:39:17.619598+03');


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: adil
--



--
-- Data for Name: watched_movie; Type: TABLE DATA; Schema: public; Owner: adil
--



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

