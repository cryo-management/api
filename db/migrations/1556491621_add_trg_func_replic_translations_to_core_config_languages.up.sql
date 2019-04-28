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