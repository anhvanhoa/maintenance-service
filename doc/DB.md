CREATE TABLE maintenance_schedules (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()), -- ID duy nhất của bản ghi bảo trì
    device_id VARCHAR(36) NOT NULL,              -- Liên kết tới thiết bị IoT cần bảo trì
    maintenance_type VARCHAR(100) COMMENT 'cleaning, calibration, replacement, repair, inspection, software_update', 
        -- Loại bảo trì: vệ sinh, hiệu chuẩn, thay thế, sửa chữa, kiểm tra, cập nhật phần mềm
    maintenance_category VARCHAR(50) COMMENT 'preventive, corrective, emergency, routine', 
        -- Phân loại bảo trì: phòng ngừa, khắc phục, khẩn cấp, định kỳ
    priority VARCHAR(50) COMMENT 'low, medium, high, critical', 
        -- Mức độ ưu tiên của công việc bảo trì
    scheduled_date DATE,                         -- Ngày dự kiến thực hiện bảo trì
    estimated_duration_hours DECIMAL(5,2),       -- Thời gian dự kiến (giờ)
    completed_date DATE,                         -- Ngày thực hiện xong bảo trì
    actual_duration_hours DECIMAL(5,2),          -- Thời gian thực tế (giờ)
    technician VARCHAR(255),                     -- Tên kỹ thuật viên phụ trách
    technician_contact VARCHAR(255),             -- Thông tin liên hệ của kỹ thuật viên
    cost DECIMAL(12,2),                          -- Chi phí bảo trì
    parts_replaced JSON COMMENT 'Array of parts with details', 
        -- Danh sách linh kiện đã thay thế (JSON: tên, số lượng, model…)
    tools_required JSON COMMENT 'Array of required tools', 
        -- Danh sách công cụ cần dùng khi bảo trì
    safety_precautions TEXT,                     -- Biện pháp an toàn cần lưu ý
    pre_maintenance_readings JSON,               -- Số liệu trước bảo trì (cảm biến, thông số…)
    post_maintenance_readings JSON,              -- Số liệu sau bảo trì để so sánh
    calibration_values JSON,                     -- Giá trị hiệu chuẩn (nếu có)
    test_results TEXT,                           -- Kết quả kiểm tra sau bảo trì
    status VARCHAR(50) DEFAULT 'scheduled' COMMENT 'scheduled, in_progress, completed, cancelled, postponed', 
        -- Trạng thái công việc: đã lên lịch, đang tiến hành, hoàn thành, hủy, hoãn
    completion_rating INTEGER COMMENT '1-5 rating of maintenance quality', 
        -- Đánh giá chất lượng bảo trì (1-5)
    next_maintenance_date DATE,                  -- Ngày bảo trì tiếp theo (nếu định kỳ)
    maintenance_interval_days INTEGER,           -- Khoảng cách giữa các lần bảo trì (ngày)
    warranty_impact BOOLEAN DEFAULT FALSE,       -- Cho biết việc bảo trì có ảnh hưởng đến bảo hành không
    downtime_minutes INTEGER,                    -- Tổng thời gian thiết bị ngưng hoạt động (phút)
    notes TEXT,                                  -- Ghi chú thêm
    maintenance_log TEXT,                        -- Nhật ký bảo trì (các bước thực hiện)
    before_images JSON,                          -- Hình ảnh trước bảo trì
    after_images JSON,                           -- Hình ảnh sau bảo trì
    created_by VARCHAR(36),                      -- Người tạo bản ghi bảo trì
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo bản ghi
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Thời điểm cập nhật
    
    INDEX idx_maintenance_schedules_device (device_id), -- Index để tìm theo thiết bị
    INDEX idx_maintenance_schedules_date (scheduled_date), -- Index để tìm theo ngày lên lịch
    INDEX idx_maintenance_schedules_status (status), -- Index để tìm theo trạng thái
    INDEX idx_maintenance_schedules_type (maintenance_type) -- Index để tìm theo loại bảo trì
);
