CREATE TABLE core_groups_users (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  user_id CHARACTER VARYING NOT NULL,
  group_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(user_id, group_id)
);