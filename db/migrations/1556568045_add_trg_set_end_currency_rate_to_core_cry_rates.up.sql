CREATE TRIGGER trg_set_end_currency_rate
  AFTER INSERT ON core_cry_rates
    FOR EACH ROW
      EXECUTE PROCEDURE trg_func_set_end_currency_rate();