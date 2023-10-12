CREATE TABLE "public.user" (
	"id_user" integer NOT NULL,
	"name" TEXT NOT NULL,
	"surname" TEXT NOT NULL,
	"middle_name" TEXT NOT NULL,
	"email" TEXT NOT NULL,
	"phone_number" integer NOT NULL,
	"region" TEXT NOT NULL,
	"tg_id" integer NOT NULL,
	"status" integer,
	"Additional_information" integer NOT NULL,
	"channel" integer,
	"event_id" integer NOT NULL,
	CONSTRAINT "user_pk" PRIMARY KEY ("id_user")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.admin" (
	"id_admin" integer NOT NULL,
	"level" integer NOT NULL,
	CONSTRAINT "admin_pk" PRIMARY KEY ("id_admin")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.channel" (
	"id_channel" integer NOT NULL,
	"region" integer NOT NULL,
	"channel_id_tg" integer NOT NULL,
	"posts" integer NOT NULL,
	CONSTRAINT "channel_pk" PRIMARY KEY ("id_channel")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.posts" (
	"id_post" integer NOT NULL,
	"img" bytea,
	"text" integer,
	"date_added" timestamptz,
	"date_of_publication" timestamptz,
	CONSTRAINT "posts_pk" PRIMARY KEY ("id_post")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.event" (
	"id_event" integer NOT NULL,
	"title" TEXT NOT NULL,
	"data" timestamptz NOT NULL,
	"description" TEXT NOT NULL,
	CONSTRAINT "event_pk" PRIMARY KEY ("id_event")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.Additionally" (
	"id_additionally" serial NOT NULL,
	"registration_date" timestamptz NOT NULL,
	"date_of_approval" timestamptz,
	"who_approved" integer,
	"status_admin" TEXT NOT NULL DEFAULT 'no',
	CONSTRAINT "Additionally_pk" PRIMARY KEY ("id_additionally")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.UserEventRegistration" (
	"id_registration" serial NOT NULL,
	"user_id" integer NOT NULL,
	"event_id" integer NOT NULL,
	"registration_date" timestamptz NOT NULL,
	CONSTRAINT "UserEventRegistration_pk" PRIMARY KEY ("id_registration")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "public.user" ADD CONSTRAINT "user_fk0" FOREIGN KEY ("status") REFERENCES "public.admin"("id_admin");
ALTER TABLE "public.user" ADD CONSTRAINT "user_fk1" FOREIGN KEY ("Additional_information") REFERENCES "public.Additionally"("id_additionally");
ALTER TABLE "public.user" ADD CONSTRAINT "user_fk2" FOREIGN KEY ("channel") REFERENCES "public.channel"("id_channel");
ALTER TABLE "public.user" ADD CONSTRAINT "user_fk3" FOREIGN KEY ("event_id") REFERENCES "public.UserEventRegistration"("id_registration");


ALTER TABLE "public.channel" ADD CONSTRAINT "channel_fk0" FOREIGN KEY ("posts") REFERENCES "public.posts"("id_post");




ALTER TABLE "public.UserEventRegistration" ADD CONSTRAINT "UserEventRegistration_fk0" FOREIGN KEY ("user_id") REFERENCES "public.user"("id_user");
ALTER TABLE "public.UserEventRegistration" ADD CONSTRAINT "UserEventRegistration_fk1" FOREIGN KEY ("event_id") REFERENCES "public.event"("id_event");









