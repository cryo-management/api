CREATE TABLE core_tree_levels (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);