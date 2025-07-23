CREATE TABLE "locations" (
	"id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "lat" float NOT NULL,
    "lon" float NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
)
