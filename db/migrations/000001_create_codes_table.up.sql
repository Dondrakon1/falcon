CREATE TABLE "codes" (
                       "id" bigserial PRIMARY KEY,
                       order_id bigint NOT NULL,
                       payload VARCHAR NOT NULL,
                       created_at timestamptz DEFAULT (now())
);

