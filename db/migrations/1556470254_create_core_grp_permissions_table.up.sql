CREATE TABLE core_grp_permissions (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  group_id CHARACTER VARYING NOT NULL,
  structure_type CHARACTER VARYING NOT NULL,
  structure_id CHARACTER VARYING NOT NULL,
  type integer NOT NULL,
  condition_query CHARACTER VARYING,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);