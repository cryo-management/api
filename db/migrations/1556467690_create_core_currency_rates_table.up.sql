CREATE TABLE core_currency_rates (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  currency_id CHARACTER VARYING NOT NULL,
  value integer NOT NULL,
  start_at TIMESTAMP NOT NULL,
  end_at TIMESTAMP NOT NULL, 
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);