CREATE TABLE core_schemas_plugins (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  plugin_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(schema_id, plugin_id)
);