DROP TABLE IF EXISTS core_users CASCADE;
DROP TABLE IF EXISTS core_trees CASCADE;
DROP TABLE IF EXISTS core_tre_levels CASCADE;
DROP TABLE IF EXISTS core_tre_units CASCADE;
DROP TABLE IF EXISTS core_currencies CASCADE;
DROP TABLE IF EXISTS core_cry_rates CASCADE;
DROP TABLE IF EXISTS core_config_languages CASCADE;
DROP TABLE IF EXISTS core_groups CASCADE;
DROP TABLE IF EXISTS core_grp_permissions CASCADE;
DROP TABLE IF EXISTS core_groups_users CASCADE;
DROP TABLE IF EXISTS core_schemas CASCADE;
DROP TABLE IF EXISTS core_schemas_modules CASCADE;
DROP TABLE IF EXISTS core_lookups CASCADE;
DROP TABLE IF EXISTS core_lkp_options CASCADE;
DROP TABLE IF EXISTS core_sch_fields CASCADE;
DROP TABLE IF EXISTS core_sch_fld_validations CASCADE;
DROP TABLE IF EXISTS core_widgets CASCADE;
DROP TABLE IF EXISTS core_sch_pages CASCADE;
DROP TABLE IF EXISTS core_sch_views CASCADE;
DROP TABLE IF EXISTS core_views_pages CASCADE;
DROP TABLE IF EXISTS core_sch_pag_sections CASCADE;
DROP TABLE IF EXISTS core_sch_pag_sec_tabs CASCADE;
DROP TABLE IF EXISTS core_sch_pag_cnt_structures CASCADE;
DROP TABLE IF EXISTS core_translations CASCADE;
DROP TABLE IF EXISTS core_jobs CASCADE;
DROP TABLE IF EXISTS core_jobs_followers CASCADE;
DROP TABLE IF EXISTS core_job_task CASCADE;
DROP VIEW IF EXISTS core_v_job_followers;
DROP VIEW IF EXISTS core_v_users_and_groups;

CREATE TABLE core_users (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  username CHARACTER VARYING NOT NULL,
  first_name CHARACTER VARYING NOT NULL,
  last_name CHARACTER VARYING NOT NULL,
  email CHARACTER VARYING NOT NULL,
  password CHARACTER VARYING NOT NULL,
  language_code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(username)
);

CREATE TABLE core_trees (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_tre_levels (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  tree_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

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

CREATE TABLE core_currencies (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_cry_rates (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  from_currency_id CHARACTER VARYING NOT NULL,
  to_currency_id CHARACTER VARYING NOT NULL,
  from_currency_code CHARACTER VARYING NOT NULL,
  to_currency_code CHARACTER VARYING NOT NULL,
  value integer NOT NULL,
  start_at TIMESTAMP NOT NULL,
  end_at TIMESTAMP, 
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE core_config_languages (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_groups (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

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

CREATE TABLE core_schemas (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  module BOOLEAN DEFAULT FALSE NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_schemas_modules (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  module_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(schema_id, module_id)
);

CREATE TABLE core_lookups (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  type CHARACTER VARYING NOT NULL,
  query CHARACTER VARYING,
  value CHARACTER VARYING NOT NULL,
  label CHARACTER VARYING NOT NULL,
  autocomplete CHARACTER VARYING,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_lkp_options (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  lookup_id CHARACTER VARYING NOT NULL,
  value CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(id, code)
);

CREATE TABLE core_sch_fields (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  field_type CHARACTER VARYING NOT NULL,
  multivalue BOOLEAN,
  lookup_id CHARACTER VARYING,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(id, code)
);

CREATE TABLE core_sch_fld_validations (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  field_id CHARACTER VARYING NOT NULL,
  validation CHARACTER VARYING NOT NULL,
  valid_when CHARACTER VARYING,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE core_widgets (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  type CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_sch_pages (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  type CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_sch_views (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_views_pages (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  view_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE core_sch_pag_sections (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_sch_pag_sec_tabs (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  section_id CHARACTER VARYING NOT NULL,
  tab_order integer NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_sch_pag_cnt_structures (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  container_id CHARACTER VARYING NOT NULL,
  container_type CHARACTER VARYING NOT NULL,
  structure_id CHARACTER VARYING NOT NULL,
  structure_type CHARACTER VARYING NOT NULL,
  position_row integer NOT NULL,
  position_column integer NOT NULL,
  width integer NOT NULL,
  height integer NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(schema_id, page_id, container_id, container_type, structure_id, structure_type)
);

CREATE TABLE core_translations (
  id CHARACTER VARYING DEFAULT uuid_generate_v1() NOT NULL,
  structure_type CHARACTER VARYING NOT NULL,
  structure_id CHARACTER VARYING NOT NULL,
  structure_field CHARACTER VARYING NOT NULL,
  value CHARACTER VARYING NOT NULL,
  language_code CHARACTER VARYING NOT NULL,
  replicated BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(structure_type, structure_id, structure_field, language_code)
);

CREATE TABLE core_jobs (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING,
  job_type CHARACTER VARYING NOT NULL, --system, user
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE core_jobs_followers (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,	
  job_id CHARACTER VARYING NOT NULL,
  follower_id CHARACTER VARYING NOT NULL,
  follower_type CHARACTER VARYING NOT NULL, --group, user
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE core_job_tasks (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING,
  job_id CHARACTER VARYING NOT NULL,
  task_sequence INTEGER NOT NULL DEFAULT 0,
  parent_id CHARACTER VARYING NOT NULL,
  exec_action CHARACTER VARYING NOT NULL, --exec_query, api_post, api_get, api_delete, api_patch
  exec_address CHARACTER VARYING, --/api/v1/schema/{parent_id}/page
  exec_payload CHARACTER VARYING NOT NULL,
  action_on_fail CHARACTER VARYING NOT NULL, --continue, retry_and_continue, cancel, retry_and_cancel, rollback, retry_and_rollback
  max_retry_attempts INTEGER DEFAULT 2,
  rollback_action CHARACTER VARYING NOT NULL, --drop table, api_delete
  rollback_address CHARACTER VARYING, --/api/v1/schema/{parent_id}/fields/{field_id}
  rollback_payload CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE VIEW core_v_users_and_groups AS 
  SELECT * FROM (
    SELECT
      core_users.id AS id,
      core_users.first_name || ' ' || core_users.last_name AS name,
      null AS language_code,
      'user' AS ug_type,
      core_users.active AS active,
      core_users.created_by AS created_by,
      core_users.created_at AS created_at,
      core_users.updated_by AS updated_by,
      core_users.updated_at AS updated_at  
    FROM core_users
  ) AS users
  UNION ALL
  SELECT * FROM (
    SELECT
      core_groups.id AS id,
      core_translations_name.value AS name,
      core_translations_name.language_code AS language_code,
      'group' AS ug_type,
      core_groups.active AS active,
      core_groups.created_by AS created_by,
      core_groups.created_at AS created_at,
      core_groups.updated_by AS updated_by,
      core_groups.updated_at AS updated_at
    FROM core_groups
    JOIN core_translations core_translations_name
    ON core_translations_name.structure_id = core_groups.id
    AND core_translations_name.structure_field = 'name'
  ) AS groups;

CREATE VIEW core_v_job_followers AS 
  SELECT
    ug.id AS id,
    followers.job_id AS job_id,
    ug.name AS name,
    ug.language_code AS language_code,
    ug.ug_type AS follower_type,
    ug.active AS active,
    followers.created_by AS created_by,
    followers.created_at AS created_at,
    followers.updated_by AS updated_by,
    followers.updated_at AS updated_at
  FROM core_jobs_followers AS followers
  JOIN core_v_users_and_groups AS ug
  ON ug.id = followers.follower_id
  AND ug.ug_type = followers.follower_type;

INSERT INTO core_users(
  id,
  username,
  first_name,
  last_name,
  email,
  password,
  language_code,
  active,
  created_by,
  created_at,
  updated_by,
  updated_at
)
VALUES (
  '307e481c-69c5-11e9-96a0-06ea2c43bb20',
  'admin',
  'Administrator',
  'System',
  'admin@domain.com',
  '123456',
  'pt-br',
  true,
  '307e481c-69c5-11e9-96a0-06ea2c43bb20',
  '2019-04-23 15:30:36.480864',
  '307e481c-69c5-11e9-96a0-06ea2c43bb20',
  '2019-04-23 15:30:36.480864'
);

INSERT INTO core_config_languages(
  id,
  code,
  active,
  created_by,
  created_at,
  updated_by,
  updated_at
)
VALUES (
  '9b09866a-69c5-11e9-96a1-06ea2c43bb20',
  'pt-br',
  true,
  '307e481c-69c5-11e9-96a0-06ea2c43bb20',
  '2019-04-23 15:30:36.480864',
  '307e481c-69c5-11e9-96a0-06ea2c43bb20',
  '2019-04-23 15:30:36.480864'
);

INSERT INTO core_translations(
  id,
  structure_type,
  structure_id,
  structure_field,
  value,
  language_code,
  created_by,
  created_at,
  updated_by,
  updated_at
)
VALUES (
  'ff1d2822-69c6-11e9-92d9-06ea2c43bb20',
  'core_config_languages',
  '9b09866a-69c5-11e9-96a1-06ea2c43bb20',
  'name',
  'Português do Brasil',
  'pt-br',
  '307e481c-69c5-11e9-96a0-06ea2c43bb20',
  '2019-04-23 15:30:36.480864',
  '307e481c-69c5-11e9-96a0-06ea2c43bb20',
  '2019-04-23 15:30:36.480864'
);

-- ALTER TABLE "core_trees" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_trees" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tre_levels" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tre_levels" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tre_units" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tre_units" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tre_units" ADD FOREIGN KEY ("parent_id") REFERENCES "core_tre_units" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_currencies" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_currencies" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_cry_rates" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_cry_rates" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_cry_rates" ADD FOREIGN KEY ("currency_id") REFERENCES "core_currencies" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_config_languages" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_config_languages" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_users" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_users" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_groups" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_groups" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_grp_permissions" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_grp_permissions" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_grp_permissions" ADD FOREIGN KEY ("group_id") REFERENCES "core_groups" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_groups_users" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_groups_users" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_groups_users" ADD FOREIGN KEY ("user_id") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_groups_users" ADD FOREIGN KEY ("group_id") REFERENCES "core_groups" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_schemas" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_schemas" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_schemas_modules" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_schemas_modules" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_schemas_modules" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_schemas_modules" ADD FOREIGN KEY ("module_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_lookups" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_lookups" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_lkp_options" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_lkp_options" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_lkp_options" ADD FOREIGN KEY ("lookup_id") REFERENCES "core_lookups" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fields" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fields" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fields" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fields" ADD FOREIGN KEY ("lookup_id") REFERENCES "core_lookups" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fld_validations" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fld_validations" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fld_validations" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_fld_validations" ADD FOREIGN KEY ("field_id") REFERENCES "core_sch_fields" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_widgets" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_widgets" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pages" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pages" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pages" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_views" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_views" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_views" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_views_pages" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_views_pages" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_views_pages" ADD FOREIGN KEY ("view_id") REFERENCES "core_sch_views" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_views_pages" ADD FOREIGN KEY ("page_id") REFERENCES "core_sch_pages" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sections" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sections" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sections" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sections" ADD FOREIGN KEY ("page_id") REFERENCES "core_sch_pages" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sec_tabs" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sec_tabs" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sec_tabs" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sec_tabs" ADD FOREIGN KEY ("page_id") REFERENCES "core_sch_pages" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_sec_tabs" ADD FOREIGN KEY ("section_id") REFERENCES "core_sch_pag_sections" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_cnt_structures" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_cnt_structures" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_cnt_structures" ADD FOREIGN KEY ("schema_id") REFERENCES "core_schemas" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_sch_pag_cnt_structures" ADD FOREIGN KEY ("page_id") REFERENCES "core_sch_pages" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_translations" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_translations" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;

CREATE OR REPLACE FUNCTION trg_func_replic_translations() RETURNS TRIGGER AS $$
  DECLARE
    curtime TIMESTAMP := NOW();
    from_lang TEXT := 'pt-br';
  BEGIN
    INSERT INTO core_translations (
      structure_type,
      structure_id,
      structure_field,
      value,
      language_code,
      replicated,
      created_by,
      created_at,
      updated_by,
      updated_at
    )
    SELECT
      structure_type,
      structure_id,
      structure_field,
      value,
      NEW.code AS language_code,
      true,
      NEW.updated_by AS created_by,
      NOW() AS created_at,
      NEW.updated_by AS updated_by,
      NOW() AS updated_at
    FROM core_translations
    WHERE
      language_code = from_lang;
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trg_func_delete_translations() RETURNS TRIGGER AS $$
  BEGIN
    DELETE FROM core_translations WHERE language_code = NEW.code;
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_replica_translations
  AFTER UPDATE ON core_config_languages
    FOR EACH ROW
      WHEN (NEW.active != OLD.active AND NEW.active = true)
        EXECUTE PROCEDURE trg_func_replic_translations();

CREATE TRIGGER trg_delete_translations
  AFTER UPDATE ON core_config_languages
    FOR EACH ROW
      WHEN (NEW.active != OLD.active AND NEW.active = false)
        EXECUTE PROCEDURE trg_func_delete_translations();

CREATE OR REPLACE FUNCTION trg_func_replic_new_translation() RETURNS TRIGGER AS $$
  BEGIN
    INSERT INTO core_translations (
      structure_type,
      structure_id,
      structure_field,
      value,
      language_code,
      replicated,
      created_by,
      created_at,
      updated_by,
      updated_at
    )
    SELECT
      NEW.structure_type AS structure_type,
      NEW.structure_id AS structure_id,
      NEW.structure_field AS structure_field,
      NEW.value AS value,
      code,
      true AS replicated,
      NEW.created_by AS created_by,
      NEW.created_at AS created_at,
      NEW.updated_by AS updated_by,
      NEW.updated_at AS updated_at
    FROM core_config_languages
    WHERE
      active = true
    AND
      code != NEW.language_code;
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_replica_new_translations
  AFTER INSERT ON core_translations
    FOR EACH ROW
      WHEN (NEW.replicated = false)
        EXECUTE PROCEDURE trg_func_replic_new_translation();

CREATE OR REPLACE FUNCTION trg_func_set_end_currency_rate() RETURNS TRIGGER AS $$
  BEGIN
    UPDATE core_cry_rates
    SET
      end_at = NEW.start_at,
      updated_by = NEW.created_by,
      updated_at = NEW.created_at
    WHERE id = (
      SELECT
        id
      FROM core_cry_rates
      WHERE
        id != NEW.id
        AND from_currency_code = NEW.from_currency_code
        AND to_currency_code = NEW.to_currency_code
      ORDER BY
        id desc
      LIMIT 1
    );
    RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_set_end_currency_rate
  AFTER INSERT ON core_cry_rates
    FOR EACH ROW
      EXECUTE PROCEDURE trg_func_set_end_currency_rate();