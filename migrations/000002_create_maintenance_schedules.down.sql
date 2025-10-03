DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM pg_trigger WHERE tgname = 'trg_maintenance_schedules_updated'
    ) THEN
        DROP TRIGGER IF EXISTS trg_maintenance_schedules_updated ON maintenance_schedules;
    END IF;
END $$;

DROP TABLE IF EXISTS maintenance_schedules;


