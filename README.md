# Maintenance Service

Microservice quản lý lịch trình bảo trì thiết bị IoT trong hệ thống công nghiệp, được xây dựng bằng Go và tuân theo nguyên tắc Clean Architecture.

## 🏗️ Kiến trúc

Dự án này tuân theo **Clean Architecture** với sự phân tách rõ ràng các mối quan tâm:

```
├── domain/           # Tầng logic nghiệp vụ
│   ├── entity/       # Các thực thể nghiệp vụ cốt lõi
│   │   └── maintenance_schedule.go     # Entity lịch trình bảo trì
│   ├── repository/   # Giao diện truy cập dữ liệu
│   │   └── maintenance_schedule_repository.go
│   └── usecase/      # Các trường hợp sử dụng nghiệp vụ
│       └── maintenance_schedule/       # Use cases lịch trình bảo trì
├── infrastructure/   # Các mối quan tâm bên ngoài
│   ├── grpc_service/ # Triển khai API gRPC
│   │   └── maintenance_schedule/      # gRPC handlers lịch trình bảo trì
│   └── repo/         # Triển khai repository cơ sở dữ liệu
├── bootstrap/        # Khởi tạo ứng dụng
└── cmd/             # Điểm vào ứng dụng
```

## 🚀 Tính năng

### Quản lý Lịch trình Bảo trì
- ✅ Tạo, đọc, cập nhật, xóa lịch trình bảo trì
- ✅ Liệt kê lịch trình với bộ lọc (thiết bị, loại bảo trì, trạng thái, ngày tháng)
- ✅ Quản lý các loại bảo trì (vệ sinh, hiệu chuẩn, thay thế, sửa chữa, kiểm tra, cập nhật phần mềm)
- ✅ Phân loại bảo trì (phòng ngừa, khắc phục, khẩn cấp, định kỳ)
- ✅ Quản lý mức độ ưu tiên (thấp, trung bình, cao, nghiêm trọng)
- ✅ Theo dõi tiến độ bảo trì (đã lên lịch → đang tiến hành → hoàn thành)
- ✅ Quản lý kỹ thuật viên và thông tin liên hệ
- ✅ Theo dõi chi phí và thời gian bảo trì
- ✅ Lưu trữ hình ảnh trước và sau bảo trì
- ✅ Quản lý linh kiện thay thế và công cụ cần thiết
- ✅ Ghi nhận số liệu trước và sau bảo trì
- ✅ Đánh giá chất lượng bảo trì
- ✅ Lập lịch bảo trì định kỳ

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.24.6
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC
- **Kiến trúc**: Clean Architecture
- **Dependencies**:
  - `github.com/go-pg/pg/v10` - PostgreSQL ORM
  - `google.golang.org/grpc` - Framework gRPC
  - `github.com/spf13/viper` - Quản lý cấu hình
  - `go.uber.org/zap` - Logging có cấu trúc

## 📋 Yêu cầu hệ thống

- Go 1.24.6 trở lên
- PostgreSQL 12 trở lên
- [golang-migrate](https://github.com/golang-migrate/migrate) để quản lý migration cơ sở dữ liệu

## 🚀 Hướng dẫn nhanh

### 1. Clone repository
```bash
git clone <repository-url>
cd maintenance-service
```

### 2. Cài đặt dependencies
```bash
go mod download
```

### 3. Thiết lập cơ sở dữ liệu
```bash
# Tạo cơ sở dữ liệu
make create-db

# Chạy migrations
make up
```

### 4. Cấu hình ứng dụng
Sao chép và chỉnh sửa file cấu hình:
```bash
cp dev.config.yml config.yml
```

Cập nhật chuỗi kết nối cơ sở dữ liệu trong `config.yml`:
```yaml
node_env: "development"
url_db: "postgres://postgres:123456@localhost:5432/maintenance_service_db?sslmode=disable"
name_service: "MaintenanceService"
port_grpc: 50054
host_grpc: "localhost"
interval_check: "20s"
timeout_check: "15s"
```

### 5. Chạy ứng dụng
```bash
# Build và chạy service chính
make run

# Hoặc chạy client để test
make client
```

## 🗄️ Quản lý Cơ sở dữ liệu

Dự án sử dụng `golang-migrate` để quản lý schema cơ sở dữ liệu:

```bash
# Chạy tất cả migrations đang chờ
make up

# Rollback migration cuối cùng
make down

# Reset cơ sở dữ liệu hoàn toàn
make reset

# Tạo migration mới
make create name=migration_name

# Force migration đến phiên bản cụ thể
make force version=1
```

## 🌱 Dữ liệu mẫu

Dự án bao gồm dữ liệu mẫu để phát triển và kiểm thử:

```bash
# Chèn dữ liệu mẫu vào cơ sở dữ liệu
make seed

# Reset cơ sở dữ liệu và chèn dữ liệu mẫu
make seed-reset

# Chèn dữ liệu mẫu vào cơ sở dữ liệu Docker
make docker-seed
```

### Dữ liệu mẫu bao gồm:

**Lịch trình bảo trì thiết bị IoT với thông tin chi tiết:**
- **Loại bảo trì**: Vệ sinh, hiệu chuẩn, thay thế, sửa chữa, kiểm tra, cập nhật phần mềm
- **Phân loại**: Phòng ngừa, khắc phục, khẩn cấp, định kỳ
- **Mức độ ưu tiên**: Thấp, trung bình, cao, nghiêm trọng
- **Trạng thái**: Đã lên lịch, đang tiến hành, hoàn thành, hủy, hoãn

Mỗi lịch trình bảo trì bao gồm:
- Thông tin thiết bị và kỹ thuật viên phụ trách
- Thời gian dự kiến và thực tế thực hiện
- Chi phí và linh kiện thay thế
- Số liệu trước và sau bảo trì
- Hình ảnh minh chứng và ghi chú chi tiết
- Đánh giá chất lượng và lịch bảo trì tiếp theo

## 📁 Cấu trúc Dự án

```
maintenance-service/
├── bootstrap/                 # Khởi tạo ứng dụng
│   ├── app.go               # Khởi tạo app
│   └── env.go               # Cấu hình môi trường
├── cmd/                     # Điểm vào ứng dụng
│   ├── main.go             # Điểm vào service chính
│   └── client/             # gRPC client để test
├── domain/                  # Logic nghiệp vụ (Clean Architecture)
│   ├── entity/             # Các thực thể nghiệp vụ cốt lõi
│   │   └── maintenance_schedule.go     # Entity lịch trình bảo trì và DTOs
│   ├── repository/         # Giao diện truy cập dữ liệu
│   │   └── maintenance_schedule_repository.go
│   └── usecase/            # Các trường hợp sử dụng nghiệp vụ
│       └── maintenance_schedule/       # Use cases lịch trình bảo trì
│           ├── create_maintenance_schedule_usecase.go
│           ├── get_maintenance_schedule_usecase.go
│           ├── list_maintenance_schedule_usecase.go
│           ├── update_maintenance_schedule_usecase.go
│           └── delete_maintenance_schedule_usecase.go
├── infrastructure/          # Các mối quan tâm bên ngoài
│   ├── grpc_service/       # Triển khai API gRPC
│   │   ├── maintenance_schedule/      # gRPC handlers lịch trình bảo trì
│   │   └── server.go             # Thiết lập gRPC server
│   └── repo/               # Triển khai cơ sở dữ liệu
│       ├── maintenance_schedule_repository.go
│       └── init.go
├── migrations/              # Database migrations
│   ├── 000000_common.up.sql
│   ├── 000002_create_maintenance_schedules.up.sql
│   └── seed/                     # Dữ liệu mẫu
├── script/seed/             # Script chèn dữ liệu mẫu
├── doc/                     # Tài liệu
└── logs/                    # Log ứng dụng
```

## 🔧 Các lệnh có sẵn

```bash
# Thao tác cơ sở dữ liệu
make up              # Chạy migrations
make down            # Rollback migration
make reset           # Reset cơ sở dữ liệu
make create-db       # Tạo cơ sở dữ liệu
make drop-db         # Xóa cơ sở dữ liệu

# Ứng dụng
make build           # Build ứng dụng
make run             # Chạy service chính
make client          # Chạy client test
make test            # Chạy tests

# Trợ giúp
make help            # Hiển thị tất cả lệnh có sẵn
```

## 📊 Mô hình Dữ liệu

### Lịch trình Bảo trì (Maintenance Schedule)
- **ID**: Định danh duy nhất
- **DeviceID**: ID thiết bị IoT cần bảo trì
- **MaintenanceType**: Loại bảo trì (cleaning, calibration, replacement, repair, inspection, software_update)
- **MaintenanceCategory**: Phân loại bảo trì (preventive, corrective, emergency, routine)
- **Priority**: Mức độ ưu tiên (low, medium, high, critical)
- **ScheduledDate**: Ngày dự kiến thực hiện bảo trì
- **EstimatedDurationHours**: Thời gian dự kiến (giờ)
- **CompletedDate**: Ngày thực hiện xong bảo trì
- **ActualDurationHours**: Thời gian thực tế (giờ)
- **Technician**: Tên kỹ thuật viên phụ trách
- **TechnicianContact**: Thông tin liên hệ của kỹ thuật viên
- **Cost**: Chi phí bảo trì
- **PartsReplaced**: Danh sách linh kiện đã thay thế (JSON)
- **ToolsRequired**: Danh sách công cụ cần dùng (JSON)
- **SafetyPrecautions**: Biện pháp an toàn cần lưu ý
- **PreMaintenanceReadings**: Số liệu trước bảo trì (JSON)
- **PostMaintenanceReadings**: Số liệu sau bảo trì (JSON)
- **CalibrationValues**: Giá trị hiệu chuẩn (JSON)
- **TestResults**: Kết quả kiểm tra sau bảo trì
- **Status**: Trạng thái (scheduled, in_progress, completed, cancelled, postponed)
- **CompletionRating**: Đánh giá chất lượng bảo trì (1-5)
- **NextMaintenanceDate**: Ngày bảo trì tiếp theo
- **MaintenanceIntervalDays**: Khoảng cách giữa các lần bảo trì (ngày)
- **WarrantyImpact**: Ảnh hưởng đến bảo hành
- **DowntimeMinutes**: Thời gian thiết bị ngưng hoạt động (phút)
- **Notes**: Ghi chú thêm
- **MaintenanceLog**: Nhật ký bảo trì
- **BeforeImages**: Hình ảnh trước bảo trì (array)
- **AfterImages**: Hình ảnh sau bảo trì (array)
- **CreatedBy**: Định danh người tạo
- **Timestamps**: Thời gian tạo/cập nhật

## 🔌 API Endpoints

Service cung cấp các endpoint gRPC:

### Maintenance Schedule Service
- `CreateMaintenanceSchedule` - Tạo lịch trình bảo trì mới
- `GetMaintenanceSchedule` - Lấy thông tin lịch trình bảo trì theo ID
- `UpdateMaintenanceSchedule` - Cập nhật thông tin lịch trình bảo trì
- `DeleteMaintenanceSchedule` - Xóa lịch trình bảo trì
- `ListMaintenanceSchedules` - Liệt kê lịch trình bảo trì với bộ lọc
- `GetByDevice` - Lấy lịch trình bảo trì theo thiết bị
- `GetByTechnician` - Lấy lịch trình bảo trì theo kỹ thuật viên
- `GetByStatus` - Lấy lịch trình bảo trì theo trạng thái
- `GetByType` - Lấy lịch trình bảo trì theo loại bảo trì
- `GetByCategory` - Lấy lịch trình bảo trì theo phân loại
- `GetByPriority` - Lấy lịch trình bảo trì theo mức độ ưu tiên
- `GetByDateRange` - Lấy lịch trình bảo trì theo khoảng ngày
- `GetUpcomingMaintenance` - Lấy lịch trình bảo trì sắp tới
- `GetOverdueMaintenance` - Lấy lịch trình bảo trì quá hạn
- `GetCompletedMaintenance` - Lấy lịch trình bảo trì đã hoàn thành
- `UpdateStatus` - Cập nhật trạng thái lịch trình bảo trì
- `UpdateCompletion` - Cập nhật thông tin hoàn thành bảo trì
- `GetMaintenanceHistory` - Lấy lịch sử bảo trì của thiết bị
- `GetMaintenanceStatistics` - Lấy thống kê bảo trì

## 🧪 Testing

Chạy client test để tương tác với service:

```bash
make client
```

Điều này sẽ khởi động một client tương tác nơi bạn có thể test tất cả các endpoint gRPC.

## 📝 Cấu hình

Ứng dụng sử dụng Viper để quản lý cấu hình. Các tùy chọn cấu hình chính:

- `node_env`: Môi trường (development, production)
- `url_db`: Chuỗi kết nối PostgreSQL
- `name_service`: Tên service cho discovery
- `port_grpc`: Cổng gRPC server
- `host_grpc`: Host gRPC server
- `interval_check`: Khoảng thời gian kiểm tra sức khỏe
- `timeout_check`: Timeout kiểm tra sức khỏe

## 🚀 Triển khai

1. **Build ứng dụng**:
   ```bash
   make build
   ```

2. **Thiết lập cơ sở dữ liệu production**:
   ```bash
   make create-db
   make up
   ```

3. **Chạy service**:
   ```bash
   ./bin/app
   ```

## 🤝 Đóng góp

1. Fork repository
2. Tạo feature branch
3. Thực hiện thay đổi
4. Thêm tests nếu cần thiết
5. Submit pull request

## 📄 Giấy phép

Dự án này được cấp phép theo MIT License.

## 🆘 Hỗ trợ

Để được hỗ trợ và đặt câu hỏi, vui lòng tạo issue trong repository.

---

**Lưu ý**: Service này được thiết kế để quản lý lịch trình bảo trì thiết bị IoT trong hệ thống công nghiệp, tuân theo các nguyên tắc kiến trúc microservice để có thể mở rộng và bảo trì dễ dàng.
