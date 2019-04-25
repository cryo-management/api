CREATE OR REPLACE TABLE core_config_languages (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
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

CREATE OR REPLACE TABLE core_users (
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
  updated_at TIMESTAMP NOT NULL
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

CREATE OR REPLACE TABLE core_groups (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_grp_permissions (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  group_id CHARACTER VARYING NOT NULL,
  structure_type CHARACTER VARYING NOT NULL,
  structure_id CHARACTER VARYING NOT NULL,
  type integer NOT NULL,
  condition_query CHARACTER VARYING,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_groups_users (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  user_id CHARACTER VARYING NOT NULL,
  group_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_schemas (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  module BOOLEAN DEFAULT FALSE NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_schemas_modules (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  module_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_lookups (
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
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_lkp_options (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  lookup_id CHARACTER VARYING NOT NULL,
  value CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_sch_fields (
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
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_sch_fld_validations (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  field_id CHARACTER VARYING NOT NULL,
  validation CHARACTER VARYING NOT NULL,
  valid_when CHARACTER VARYING,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_widgets (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  type CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_sch_pages (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  type CHARACTER VARYING NOT NULL,
  active BOOLEAN DEFAULT FALSE NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_sch_views (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_views_pages (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  view_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_sch_pag_sections (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_sch_pag_sec_tabs (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  code CHARACTER VARYING NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  section_id CHARACTER VARYING NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_sch_pag_sec_structures (
  id CHARACTER VARYING DEFAULT uuid_generate_v4() NOT NULL,
  schema_id CHARACTER VARYING NOT NULL,
  page_id CHARACTER VARYING NOT NULL,
  container_id CHARACTER VARYING NOT NULL,
  container_type CHARACTER VARYING NOT NULL,
  structure_type CHARACTER VARYING NOT NULL,
  structure_id CHARACTER VARYING NOT NULL,
  "row" integer NOT NULL,
  "column" integer NOT NULL,
  width integer NOT NULL,
  height integer NOT NULL,
  created_by CHARACTER VARYING NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_by CHARACTER VARYING NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE OR REPLACE TABLE core_translations (
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
  updated_at TIMESTAMP NOT NULL
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

CREATE TRIGGER trg_replica_translations
  AFTER INSERT ON core_translations
    FOR EACH ROW
      WHEN NEW.replicated = false
        EXECUTE PROCEDURE trg_func_replic_new_translation();

INSERT INTO fields (id, code, schema_id, type, multivalue, lookup_id, active) VALUES
('8a09225d-ab04-49ae-ba7c-e1ecf6025752', 'start_at', 'ffa7cfb3-5220-4430-800a-c0be84421db6', 'date', null, null, true),
('1c10e475-102d-4681-b0eb-45a565603db1', 'finish_at', 'ffa7cfb3-5220-4430-800a-c0be84421db6', 'date', null, null, true),
('e2440f61-ee3c-4e5c-ac0d-1b9442ebf90d', 'value', 'ffa7cfb3-5220-4430-800a-c0be84421db6', 'number', null, null, true),
('ccd8f50c-19b8-49b7-b83c-8e8d63f8f7c6', 'manager', 'ffa7cfb3-5220-4430-800a-c0be84421db6', 'text', null, null, true),
('c32acb58-240c-47bc-ae08-1923a80bbb19', 'title', '68aafe12-bb6d-4ae7-bf75-e0aaf8bfed10', 'text', null, null, true);

INSERT INTO groups (id, code, active) VALUES
('3aee1c18-2245-4a4f-b499-8457ecb5e169', 'administrator', true),
('0e489612-b71e-47a1-ab0b-916c37cb5a25', 'pmo', true),
('b04cc232-b97c-4991-974e-190717afa246', 'view', true);

INSERT INTO groups_permissions (id, group_id, structure_type, structure_id, type, condition_query) VALUES
('fba4cafa-3a0c-4047-b28c-350a66d748f4', '0e489612-b71e-47a1-ab0b-916c37cb5a25', 'field', 'ccd8f50c-19b8-49b7-b83c-8e8d63f8f7c6', 200, null),
('9f2d0b0a-edda-415b-a296-2975183ea93f', '0e489612-b71e-47a1-ab0b-916c37cb5a25', 'field', '8a09225d-ab04-49ae-ba7c-e1ecf6025752', 100, null),
('97193257-0d8a-4fa9-be97-23f45d61cf34', '0e489612-b71e-47a1-ab0b-916c37cb5a25', 'field', '1c10e475-102d-4681-b0eb-45a565603db1', 100, null),
('76aa7047-4b1c-43ab-a910-5c9f5d1eef18', 'b04cc232-b97c-4991-974e-190717afa246', 'field', 'e2440f61-ee3c-4e5c-ac0d-1b9442ebf90d', 100, null),
('6b10b054-f33c-4754-80e7-7fd0a359ac13', '0e489612-b71e-47a1-ab0b-916c37cb5a25', 'field', 'e2440f61-ee3c-4e5c-ac0d-1b9442ebf90d', 200, null);

INSERT INTO schemas (id, code, module, active, last_modified_at) VALUES
('ffa7cfb3-5220-4430-800a-c0be84421db6', 'SC016', false, true, '2019-04-05'),
('68aafe12-bb6d-4ae7-bf75-e0aaf8bfed10', 'tarefas', false, true, null);

INSERT INTO translations (id, structure_type, structure_id, structure_field, value, language_code) VALUES
('e5f2b783-1e5a-4c93-8094-808cd020c95f', 'schemas', 'ffa7cfb3-5220-4430-800a-c0be84421db6', 'name', 'Contrato', 'pt-br'),
('25aa5a65-7f0a-4473-9d52-1087535ce0b2', 'schemas', 'ffa7cfb3-5220-4430-800a-c0be84421db6', 'description', 'Estrutura para armazenar atributos de contratos', 'pt-br'),
('a8f4028e-f4de-4534-a4b2-4547795aa200', 'fields', 'ccd8f50c-19b8-49b7-b83c-8e8d63f8f7c6', 'name', 'Gerente', 'pt-br'),
('617953df-6450-47dc-9cee-bb8f0fdb098d', 'fields', 'ccd8f50c-19b8-49b7-b83c-8e8d63f8f7c6', 'description', 'Descrição do Gerente', 'pt-br'),
('d2a69c53-1bed-45bc-8684-7ebb4659db6a', 'fields', '8a09225d-ab04-49ae-ba7c-e1ecf6025752', 'name', 'Data de Início', 'pt-br'),
('e1bc9428-1577-4545-86d1-c84f93ceb385', 'fields', '8a09225d-ab04-49ae-ba7c-e1ecf6025752', 'description', 'Descrição da Data de Início', 'pt-br'),
('6683555d-eb37-4ec4-a856-0887c2e04072', 'fields', '1c10e475-102d-4681-b0eb-45a565603db1', 'name', 'Data de Término', 'pt-br'),
('3dc472e6-feeb-423e-977c-ea56c9ba9abc', 'fields', '1c10e475-102d-4681-b0eb-45a565603db1', 'description', 'Descrição da Data de Término', 'pt-br'),
('7b5b735b-d94e-4732-b34b-a0d4c4c51a80', 'fields', 'e2440f61-ee3c-4e5c-ac0d-1b9442ebf90d', 'name', 'Valor', 'pt-br'),
('cf05bb82-f080-44ca-a921-6345d8a8424f', 'fields', 'e2440f61-ee3c-4e5c-ac0d-1b9442ebf90d', 'description', 'Descrição do Valor', 'pt-br'),
('7089afb6-2a1a-43d0-b443-f32ea74979c3', 'schemas', '68aafe12-bb6d-4ae7-bf75-e0aaf8bfed10', 'name', 'Tarefas', 'pt-br'),
('7156b6fb-9707-463a-b049-44114f3e2aa3', 'schemas', '68aafe12-bb6d-4ae7-bf75-e0aaf8bfed10', 'description', 'Estrutura para armazenar atributos de tarefas', 'pt-br'),
('cd4f010f-061a-4840-87b9-b7d24ec899dc', 'fields', 'c32acb58-240c-47bc-ae08-1923a80bbb19', 'name', 'Título', 'pt-br'),
('779a0224-f54f-4b45-881e-17deff073e3f', 'fields', 'c32acb58-240c-47bc-ae08-1923a80bbb19', 'description', 'Descrição do Título', 'pt-br')
('0d021589-0284-4898-b70b-5574a17274f7', 'groups', '3aee1c18-2245-4a4f-b499-8457ecb5e169', 'name', 'Administradores', 'pt-br'),
('1fe02a71-f491-43b9-a758-17a2a15665b5', 'groups', '3aee1c18-2245-4a4f-b499-8457ecb5e169', 'description', 'Grupo com permissão em todo o sistema', 'pt-br'),
('8ae315e3-b8a8-4182-93ca-6639139e8fc4', 'groups', '0e489612-b71e-47a1-ab0b-916c37cb5a25', 'name', 'PMO', 'pt-br'),
('568fb7cc-3a11-4026-9b30-3eb4ccceee20', 'groups', '0e489612-b71e-47a1-ab0b-916c37cb5a25', 'description', 'Grupo com permissão de PMO', 'pt-br'),
('ca29c2db-c269-45de-aab3-8883288f43c1', 'groups', 'b04cc232-b97c-4991-974e-190717afa246', 'name', 'View', 'pt-br'),
('168757ef-fe05-419e-963e-1039e1ae535a', 'groups', 'b04cc232-b97c-4991-974e-190717afa246', 'description', 'Grupo com permissão de visualização', 'pt-br');

INSERT INTO groups_users (id, user_id, group_id) VALUES
('c796884a-a401-406a-9c14-9a7e4d55d62b', '059fa339-025c-4104-ab55-c764d3028f63', '0e489612-b71e-47a1-ab0b-916c37cb5a25'),
('37bfe49b-8bb8-44dc-aafc-d780e0193c2b', '0625ad35-a49d-440a-9ea0-83f0740717ae', 'b04cc232-b97c-4991-974e-190717afa246'),
('582d71b2-7f36-47d9-933b-3897f476baf5', '059fa339-025c-4104-ab55-c764d3028f63', 'b04cc232-b97c-4991-974e-190717afa246');

INSERT INTO users (id, first_name, last_name, email, password, active, language) VALUES
('059fa339-025c-4104-ab55-c764d3028f63', 'Bruno', 'Piaui', 'brunopiaui@gmail.com', '12345', true, 'pt-br'),
('0625ad35-a49d-440a-9ea0-83f0740717ae', 'Andre', 'Mendonça', 'andreluzz@gmail.com', '12345', true, 'pt-br');

ALTER TABLE
  "fields"
ADD
  FOREIGN KEY ("schema_id") REFERENCES "schemas" ("id") ON DELETE CASCADE;
ALTER TABLE
  "fields_validations"
ADD
  FOREIGN KEY ("field_id") REFERENCES "fields" ("id") ON DELETE CASCADE;
ALTER TABLE
  "groups_users"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE;
ALTER TABLE
  "groups_users"
ADD
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE
  "groups_permissions"
ADD
  FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE;
ALTER TABLE
  "lookups_options"
ADD
  FOREIGN KEY ("lookup_id") REFERENCES "lookups" ("id") ON DELETE CASCADE;
ALTER TABLE
  "pages"
ADD
  FOREIGN KEY ("view_id") REFERENCES "views" ("id") ON DELETE CASCADE;
ALTER TABLE
  "sections"
ADD
  FOREIGN KEY ("view_id") REFERENCES "views" ("id") ON DELETE CASCADE;
ALTER TABLE
  "tabs"
ADD
  FOREIGN KEY ("view_id") REFERENCES "views" ("id") ON DELETE CASCADE;