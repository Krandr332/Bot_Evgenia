ALTER TABLE "public.Additionally" 
ALTER COLUMN "status" SET DEFAULT 0;

ALTER TABLE "public.channel" ADD COLUMN "address" TEXT;


ALTER TABLE "public.Additionally" ALTER COLUMN "status" SET DEFAULT NULL;

ALTER TABLE "public.channel" ALTER COLUMN "region" TYPE TEXT;
ALTER TABLE "public.channel" ALTER COLUMN "channel_id_tg" TYPE TEXT;


ALTER TABLE "public.channel" ALTER COLUMN posts DROP DEFAULT;


ALTER TABLE "public.posts" ALTER COLUMN "text" TYPE TEXT;
-- Изменить тип данных поля "text" на "text" в таблице "public.posts"
ALTER TABLE "public.posts"
ALTER COLUMN text TYPE text
USING text::text;
