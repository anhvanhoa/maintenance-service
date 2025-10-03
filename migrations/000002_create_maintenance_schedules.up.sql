CREATE TABLE IF NOT EXISTS maintenance_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    device_id UUID NOT NULL,
    maintenance_type VARCHAR(100),
    maintenance_category VARCHAR(50),
    priority VARCHAR(50),
    scheduled_date DATE,
    estimated_duration_hours NUMERIC(5,2),
    completed_date DATE,
    actual_duration_hours NUMERIC(5,2),
    technician VARCHAR(255),
    technician_contact VARCHAR(255),
    cost NUMERIC(12,2),
    parts_replaced JSONB,
    tools_required JSONB,
    safety_precautions TEXT,
    pre_maintenance_readings JSONB,
    post_maintenance_readings JSONB,
    calibration_values JSONB,
    test_results TEXT,
    status VARCHAR(50) DEFAULT 'scheduled',
    completion_rating INTEGER,
    next_maintenance_date DATE,
    maintenance_interval_days INTEGER,
    warranty_impact BOOLEAN DEFAULT FALSE,
    downtime_minutes INTEGER,
    notes TEXT,
    maintenance_log TEXT,
    before_images JSONB,
    after_images JSONB,
    created_by UUID,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_maintenance_schedules_device ON maintenance_schedules (device_id);
CREATE INDEX IF NOT EXISTS idx_maintenance_schedules_date ON maintenance_schedules (scheduled_date);
CREATE INDEX IF NOT EXISTS idx_maintenance_schedules_status ON maintenance_schedules (status);
CREATE INDEX IF NOT EXISTS idx_maintenance_schedules_type ON maintenance_schedules (maintenance_type);

DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM pg_proc p
        JOIN pg_namespace n ON n.oid = p.pronamespace
        WHERE p.proname = 'update_updated_at_column' AND n.nspname = 'public'
    ) THEN
        CREATE TRIGGER trg_maintenance_schedules_updated
        BEFORE UPDATE ON maintenance_schedules
        FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
    END IF;
END $$;


