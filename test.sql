-- Замените значение 1 на действительный channel_id
INSERT INTO "public.posts" (channel_id, img, text, date_added, date_of_publication)
VALUES (1, E'\\x0123456789ABCDEF', 'Ваш текст поста', '2023-10-28 15:30:00+02:00'::timestamptz, '2023-10-29 12:00:00+02:00'::timestamptz);
