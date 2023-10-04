CREATE TABLE "public.User" (
	"id" serial NOT NULL,
	"Name" TEXT NOT NULL,
	"Surname" TEXT NOT NULL,
	"Middle_Name" TEXT NOT NULL,
	"email" TEXT NOT NULL,
	"Phone_number" integer NOT NULL,
	"Region" TEXT NOT NULL,
	"Registration_date" timestamptz NOT NULL,
	"tg_id" integer NOT NULL,
	"Date_of_approval" timestamptz NOT NULL,
	"Who_approved" integer NOT NULL,
	"id_Channel" integer,
	"Status_admin" TEXT NOT NULL DEFAULT 'no',
	"status" integer NOT NULL,
	"event_id" integer,
	CONSTRAINT "User_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.Channel" (
	"id_Channel" serial NOT NULL,
	"Region" integer NOT NULL,
	"Post" integer NOT NULL,
	"channel_id_tg" integer NOT NULL,
	CONSTRAINT "Channel_pk" PRIMARY KEY ("id_Channel")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.Posts" (
	"id_post" serial NOT NULL,
	"img" bytea NOT NULL,
	"text" integer NOT NULL,
	"id_admin" integer NOT NULL,
	"Date_added" timestamptz NOT NULL,
	"Date_of_publication" timestamptz NOT NULL,
	CONSTRAINT "Posts_pk" PRIMARY KEY ("id_post")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "public.Event" (
	"id_Event" serial NOT NULL,
	"Title" TEXT NOT NULL,
	"data" timestamptz NOT NULL,
	"Description" TEXT NOT NULL,
	CONSTRAINT "Event_pk" PRIMARY KEY ("id_Event")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "public.User" ADD CONSTRAINT "User_fk0" FOREIGN KEY ("id_Channel") REFERENCES "public.Channel"("id_Channel");
ALTER TABLE "public.User" ADD CONSTRAINT "User_fk1" FOREIGN KEY ("event_id") REFERENCES "public.Event"("id_Event");

ALTER TABLE "public.Channel" ADD CONSTRAINT "Channel_fk0" FOREIGN KEY ("Post") REFERENCES "public.Posts"("id_post");

ALTER TABLE "public.Posts" ADD CONSTRAINT "Posts_fk0" FOREIGN KEY ("id_admin") REFERENCES "public.User"("id");






