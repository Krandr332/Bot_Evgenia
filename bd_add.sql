ALTER TABLE "public.user" ALTER COLUMN "event_id" DROP NOT NULL;

ALTER TABLE "public.user" ALTER COLUMN "Additional_information" DROP NOT NULL;
ALTER TABLE "public.user" ALTER COLUMN "Additional_information" SET DATA TYPE integer USING NULL;
