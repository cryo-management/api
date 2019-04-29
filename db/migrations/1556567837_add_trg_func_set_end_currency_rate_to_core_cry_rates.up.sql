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