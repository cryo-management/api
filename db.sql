DROP TABLE IF EXISTS core_tree CASCADE;
DROP TABLE IF EXISTS core_tree_levels CASCADE;
DROP TABLE IF EXISTS core_tree_units CASCADE;
DROP TABLE IF EXISTS core_currencies CASCADE;
DROP TABLE IF EXISTS core_currency_rates CASCADE;
DROP TABLE IF EXISTS core_config_languages CASCADE;
DROP TABLE IF EXISTS core_users CASCADE;
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

CREATE TABLE core_tree (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE core_tree_levels (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE core_tree_units (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  parent_id CHARACTER VARYING,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE core_currencies (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
);

CREATE TABLE core_currency_rates (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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

CREATE TABLE core_config_languages (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id),
  UNIQUE(code)
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
  '8629ddba-f153-482f-ad9b-b7d4fae54d07',
  'pt-br',
  true,
  '57a97aaf-16da-44ef-a8be-b1caf52becd6',
  '2019-04-23 15:30:36.480864',
  '57a97aaf-16da-44ef-a8be-b1caf52becd6',
  '2019-04-23 15:30:36.480864'
);

CREATE TABLE core_users (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  '57a97aaf-16da-44ef-a8be-b1caf52becd6',
  'admin',
  'Administrator',
  'System',
  'admin@domain.com',
  '123456',
  'pt-br',
  true,
  '57a97aaf-16da-44ef-a8be-b1caf52becd6',
  '2019-04-23 15:30:36.480864',
  '57a97aaf-16da-44ef-a8be-b1caf52becd6',
  '2019-04-23 15:30:36.480864'
);

CREATE TABLE core_groups (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  view_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE core_sch_pag_sections (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
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

INSERT INTO core_translations(
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
  'core_config_languages',
  '8629ddba-f153-482f-ad9b-b7d4fae54d07',
  'name',
  'Português do Brasil',
  'pt-br',
  '57a97aaf-16da-44ef-a8be-b1caf52becd6',
  '2019-04-23 15:30:36.480864',
  '57a97aaf-16da-44ef-a8be-b1caf52becd6',
  '2019-04-23 15:30:36.480864'
);

-- ALTER TABLE "core_tree" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tree" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tree_levels" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tree_levels" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tree_units" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tree_units" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_tree_units" ADD FOREIGN KEY ("parent_id") REFERENCES "core_tree_units" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_currencies" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_currencies" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_currency_rates" ADD FOREIGN KEY ("created_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_currency_rates" ADD FOREIGN KEY ("updated_by") REFERENCES "core_users" ("id") ON DELETE CASCADE;
-- ALTER TABLE "core_currency_rates" ADD FOREIGN KEY ("currency_id") REFERENCES "core_currencies" ("id") ON DELETE CASCADE;
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