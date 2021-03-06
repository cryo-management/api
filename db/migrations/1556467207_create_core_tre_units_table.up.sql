CREATE TABLE core_tre_units (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  tree_id CHARACTER VARYING NOT NULL,
  parent_id CHARACTER VARYING,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);