CREATE TRIGGER trg_replica_new_translations
  AFTER INSERT ON core_translations
    FOR EACH ROW
      WHEN (NEW.replicated = false)
        EXECUTE PROCEDURE trg_func_replic_new_translation();