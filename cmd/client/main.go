package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_maintenance_schedule "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Nhập trang (mặc định 1): ")
	offsetStr, _ := reader.ReadString('\n')
	offsetStr = cleanInput(offsetStr)
	offset := int32(1)
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(o)
		}
	}

	fmt.Print("Nhập số bản ghi mỗi trang (mặc định 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	return offset, limit
}

type MaintenanceServiceClient struct {
	maintenanceScheduleClient proto_maintenance_schedule.MaintenanceScheduleServiceClient
	conn                      *grpc.ClientConn
}

func NewMaintenanceServiceClient(address string) (*MaintenanceServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &MaintenanceServiceClient{
		maintenanceScheduleClient: proto_maintenance_schedule.NewMaintenanceScheduleServiceClient(conn),
		conn:                      conn,
	}, nil
}

func (c *MaintenanceServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Maintenance Schedule Service Tests ==================

func (c *MaintenanceServiceClient) TestCreateMaintenanceSchedule() {
	fmt.Println("\n=== Kiểm thử Tạo lịch bảo trì ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID thiết bị (UUID): ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập loại bảo trì (cleaning/calibration/replacement/repair/inspection/software_update): ")
	maintenanceType, _ := reader.ReadString('\n')
	maintenanceType = cleanInput(maintenanceType)
	if maintenanceType == "" {
		maintenanceType = "cleaning" // Default value
	}

	fmt.Print("Nhập danh mục bảo trì (preventive/corrective/emergency/routine): ")
	maintenanceCategory, _ := reader.ReadString('\n')
	maintenanceCategory = cleanInput(maintenanceCategory)
	if maintenanceCategory == "" {
		maintenanceCategory = "preventive" // Default value
	}

	fmt.Print("Nhập mức độ ưu tiên (low/medium/high/critical): ")
	priority, _ := reader.ReadString('\n')
	priority = cleanInput(priority)
	if priority == "" {
		priority = "medium" // Default value
	}

	fmt.Print("Nhập ngày lên lịch (YYYY-MM-DD): ")
	scheduledDateStr, _ := reader.ReadString('\n')
	scheduledDateStr = cleanInput(scheduledDateStr)
	var scheduledDate *timestamppb.Timestamp
	if scheduledDateStr != "" {
		if t, err := time.Parse("2006-01-02", scheduledDateStr); err == nil {
			scheduledDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập thời gian ước tính (giờ): ")
	durationStr, _ := reader.ReadString('\n')
	durationStr = cleanInput(durationStr)
	duration := float64(2.0)
	if durationStr != "" {
		if d, err := strconv.ParseFloat(durationStr, 64); err == nil {
			duration = d
		}
	}

	fmt.Print("Nhập tên kỹ thuật viên: ")
	technician, _ := reader.ReadString('\n')
	technician = cleanInput(technician)

	fmt.Print("Nhập liên hệ kỹ thuật viên: ")
	technicianContact, _ := reader.ReadString('\n')
	technicianContact = cleanInput(technicianContact)

	fmt.Print("Nhập chi phí: ")
	costStr, _ := reader.ReadString('\n')
	costStr = cleanInput(costStr)
	cost := float64(100.0)
	if costStr != "" {
		if c, err := strconv.ParseFloat(costStr, 64); err == nil {
			cost = c
		}
	}

	fmt.Print("Nhập linh kiện thay thế: ")
	partsReplaced, _ := reader.ReadString('\n')
	partsReplaced = cleanInput(partsReplaced)

	fmt.Print("Nhập dụng cụ cần thiết: ")
	toolsRequired, _ := reader.ReadString('\n')
	toolsRequired = cleanInput(toolsRequired)

	fmt.Print("Nhập biện pháp an toàn: ")
	safetyPrecautions, _ := reader.ReadString('\n')
	safetyPrecautions = cleanInput(safetyPrecautions)

	fmt.Print("Nhập ghi chú: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Nhập người tạo: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	// Additional required fields
	fmt.Print("Nhập trạng thái (scheduled/in_progress/completed/cancelled/postponed): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "scheduled" // Default value
	}

	fmt.Print("Nhập đánh giá hoàn thành (1-5): ")
	ratingStr, _ := reader.ReadString('\n')
	ratingStr = cleanInput(ratingStr)
	rating := int32(1) // Default value
	if ratingStr != "" {
		if r, err := strconv.Atoi(ratingStr); err == nil && r >= 1 && r <= 5 {
			rating = int32(r)
		}
	}

	fmt.Print("Nhập khoảng cách bảo trì (ngày, 1-3650): ")
	intervalStr, _ := reader.ReadString('\n')
	intervalStr = cleanInput(intervalStr)
	interval := int32(30) // Default value
	if intervalStr != "" {
		if i, err := strconv.Atoi(intervalStr); err == nil && i >= 1 && i <= 3650 {
			interval = int32(i)
		}
	}

	fmt.Print("Ảnh hưởng bảo hành (true/false): ")
	warrantyImpactStr, _ := reader.ReadString('\n')
	warrantyImpactStr = cleanInput(warrantyImpactStr)
	warrantyImpact := false
	if warrantyImpactStr == "true" {
		warrantyImpact = true
	}

	fmt.Print("Nhập thời gian ngừng hoạt động (phút, 0-10080): ")
	downtimeStr, _ := reader.ReadString('\n')
	downtimeStr = cleanInput(downtimeStr)
	downtime := int32(0) // Default value
	if downtimeStr != "" {
		if d, err := strconv.Atoi(downtimeStr); err == nil && d >= 0 && d <= 10080 {
			downtime = int32(d)
		}
	}

	fmt.Print("Nhập nhật ký bảo trì: ")
	maintenanceLog, _ := reader.ReadString('\n')
	maintenanceLog = cleanInput(maintenanceLog)

	fmt.Print("Nhập hình ảnh trước: ")
	beforeImages, _ := reader.ReadString('\n')
	beforeImages = cleanInput(beforeImages)

	fmt.Print("Nhập hình ảnh sau: ")
	afterImages, _ := reader.ReadString('\n')
	afterImages = cleanInput(afterImages)

	// Convert image strings to slices
	var beforeImagesSlice []string
	if beforeImages != "" {
		beforeImagesSlice = []string{beforeImages}
	}
	var afterImagesSlice []string
	if afterImages != "" {
		afterImagesSlice = []string{afterImages}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.maintenanceScheduleClient.CreateMaintenanceSchedule(ctx, &proto_maintenance_schedule.CreateMaintenanceScheduleRequest{
		DeviceId:                deviceId,
		MaintenanceType:         maintenanceType,
		MaintenanceCategory:     maintenanceCategory,
		Priority:                priority,
		ScheduledDate:           scheduledDate,
		EstimatedDurationHours:  duration,
		Technician:              technician,
		TechnicianContact:       technicianContact,
		Cost:                    cost,
		PartsReplaced:           partsReplaced,
		ToolsRequired:           toolsRequired,
		SafetyPrecautions:       safetyPrecautions,
		Status:                  status,
		CompletionRating:        rating,
		MaintenanceIntervalDays: interval,
		WarrantyImpact:          warrantyImpact,
		DowntimeMinutes:         downtime,
		Notes:                   notes,
		MaintenanceLog:          maintenanceLog,
		BeforeImages:            beforeImagesSlice,
		AfterImages:             afterImagesSlice,
		CreatedBy:               createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateMaintenanceSchedule: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo lịch bảo trì:\n")
	if resp.MaintenanceSchedule != nil {
		fmt.Printf("ID: %s\n", resp.MaintenanceSchedule.Id)
		fmt.Printf("Device ID: %s\n", resp.MaintenanceSchedule.DeviceId)
		fmt.Printf("Maintenance Type: %s\n", resp.MaintenanceSchedule.MaintenanceType)
		fmt.Printf("Maintenance Category: %s\n", resp.MaintenanceSchedule.MaintenanceCategory)
		fmt.Printf("Priority: %s\n", resp.MaintenanceSchedule.Priority)
		fmt.Printf("Status: %s\n", resp.MaintenanceSchedule.Status)
		fmt.Printf("Estimated Duration: %.2f hours\n", resp.MaintenanceSchedule.EstimatedDurationHours)
		fmt.Printf("Technician: %s\n", resp.MaintenanceSchedule.Technician)
		fmt.Printf("Cost: %.2f\n", resp.MaintenanceSchedule.Cost)
		fmt.Printf("Completion Rating: %d\n", resp.MaintenanceSchedule.CompletionRating)
		fmt.Printf("Maintenance Interval: %d days\n", resp.MaintenanceSchedule.MaintenanceIntervalDays)
		fmt.Printf("Warranty Impact: %t\n", resp.MaintenanceSchedule.WarrantyImpact)
		fmt.Printf("Downtime: %d minutes\n", resp.MaintenanceSchedule.DowntimeMinutes)
		fmt.Printf("Created By: %s\n", resp.MaintenanceSchedule.CreatedBy)
	}
}

func (c *MaintenanceServiceClient) TestGetMaintenanceSchedule() {
	fmt.Println("\n=== Kiểm thử Lấy lịch bảo trì ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID lịch bảo trì: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.maintenanceScheduleClient.GetMaintenanceSchedule(ctx, &proto_maintenance_schedule.GetMaintenanceScheduleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetMaintenanceSchedule: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy lịch bảo trì:\n")
	if resp.MaintenanceSchedule != nil {
		fmt.Printf("ID: %s\n", resp.MaintenanceSchedule.Id)
		fmt.Printf("Device ID: %s\n", resp.MaintenanceSchedule.DeviceId)
		fmt.Printf("Maintenance Type: %s\n", resp.MaintenanceSchedule.MaintenanceType)
		fmt.Printf("Maintenance Category: %s\n", resp.MaintenanceSchedule.MaintenanceCategory)
		fmt.Printf("Priority: %s\n", resp.MaintenanceSchedule.Priority)
		fmt.Printf("Scheduled Date: %s\n", resp.MaintenanceSchedule.ScheduledDate)
		fmt.Printf("Estimated Duration: %.2f hours\n", resp.MaintenanceSchedule.EstimatedDurationHours)
		fmt.Printf("Technician: %s\n", resp.MaintenanceSchedule.Technician)
		fmt.Printf("Technician Contact: %s\n", resp.MaintenanceSchedule.TechnicianContact)
		fmt.Printf("Cost: %.2f\n", resp.MaintenanceSchedule.Cost)
		fmt.Printf("Parts Replaced: %s\n", resp.MaintenanceSchedule.PartsReplaced)
		fmt.Printf("Tools Required: %s\n", resp.MaintenanceSchedule.ToolsRequired)
		fmt.Printf("Safety Precautions: %s\n", resp.MaintenanceSchedule.SafetyPrecautions)
		fmt.Printf("Status: %s\n", resp.MaintenanceSchedule.Status)
		fmt.Printf("Notes: %s\n", resp.MaintenanceSchedule.Notes)
		fmt.Printf("Created By: %s\n", resp.MaintenanceSchedule.CreatedBy)
	}
}

func (c *MaintenanceServiceClient) TestListMaintenanceSchedules() {
	fmt.Println("\n=== Kiểm thử Liệt kê Lịch bảo trì ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	// Input filter options
	fmt.Println("\n--- Bộ lọc (để trống để bỏ qua) ---")

	fmt.Print("Nhập Device ID (UUID): ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập trạng thái (scheduled/in_progress/completed/cancelled/postponed), phân cách bằng dấu phẩy: ")
	statusesStr, _ := reader.ReadString('\n')
	statusesStr = cleanInput(statusesStr)
	var statuses []string
	if statusesStr != "" {
		statuses = strings.Split(statusesStr, ",")
		for i, status := range statuses {
			statuses[i] = strings.TrimSpace(status)
		}
	}

	fmt.Print("Nhập loại bảo trì (cleaning/calibration/replacement/repair/inspection/software_update), phân cách bằng dấu phẩy: ")
	typesStr, _ := reader.ReadString('\n')
	typesStr = cleanInput(typesStr)
	var types []string
	if typesStr != "" {
		types = strings.Split(typesStr, ",")
		for i, t := range types {
			types[i] = strings.TrimSpace(t)
		}
	}

	fmt.Print("Nhập danh mục bảo trì (preventive/corrective/emergency/routine), phân cách bằng dấu phẩy: ")
	categoriesStr, _ := reader.ReadString('\n')
	categoriesStr = cleanInput(categoriesStr)
	var categories []string
	if categoriesStr != "" {
		categories = strings.Split(categoriesStr, ",")
		for i, category := range categories {
			categories[i] = strings.TrimSpace(category)
		}
	}

	fmt.Print("Nhập mức độ ưu tiên (low/medium/high/critical), phân cách bằng dấu phẩy: ")
	prioritiesStr, _ := reader.ReadString('\n')
	prioritiesStr = cleanInput(prioritiesStr)
	var priorities []string
	if prioritiesStr != "" {
		priorities = strings.Split(prioritiesStr, ",")
		for i, priority := range priorities {
			priorities[i] = strings.TrimSpace(priority)
		}
	}

	fmt.Print("Nhập ngày bắt đầu (YYYY-MM-DD): ")
	fromDateStr, _ := reader.ReadString('\n')
	fromDateStr = cleanInput(fromDateStr)
	var fromDate *timestamppb.Timestamp
	if fromDateStr != "" {
		if t, err := time.Parse("2006-01-02", fromDateStr); err == nil {
			fromDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập ngày kết thúc (YYYY-MM-DD): ")
	toDateStr, _ := reader.ReadString('\n')
	toDateStr = cleanInput(toDateStr)
	var toDate *timestamppb.Timestamp
	if toDateStr != "" {
		if t, err := time.Parse("2006-01-02", toDateStr); err == nil {
			toDate = timestamppb.New(t)
		}
	}

	filter := &proto_maintenance_schedule.MaintenanceScheduleFilter{}
	if deviceId != "" {
		filter.DeviceId = deviceId
	}
	if len(statuses) > 0 {
		filter.Statuses = statuses
	}
	if len(types) > 0 {
		filter.Types = types
	}
	if len(categories) > 0 {
		filter.Categories = categories
	}
	if len(priorities) > 0 {
		filter.Priorities = priorities
	}
	if fromDate != nil {
		filter.FromDate = fromDate
	}
	if toDate != nil {
		filter.ToDate = toDate
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.maintenanceScheduleClient.ListMaintenanceSchedule(ctx, &proto_maintenance_schedule.ListMaintenanceScheduleRequest{
		Pagination: &proto_common.PaginationRequest{
			Page:     offset,
			PageSize: limit,
		},
		Filter: filter,
	})
	if err != nil {
		fmt.Printf("Error calling ListMaintenanceSchedule: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê lịch bảo trì:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Trang: %d/%d\n", resp.Pagination.Page, resp.Pagination.TotalPages)
	fmt.Printf("Kích thước trang: %d\n", resp.Pagination.PageSize)
	fmt.Printf("Danh sách lịch bảo trì:\n")
	for i, schedule := range resp.MaintenanceSchedules {
		fmt.Printf("  [%d] ID: %s, Device ID: %s, Type: %s, Category: %s, Priority: %s, Status: %s\n",
			i+1, schedule.Id, schedule.DeviceId, schedule.MaintenanceType, schedule.MaintenanceCategory, schedule.Priority, schedule.Status)
	}
}

func (c *MaintenanceServiceClient) TestUpdateMaintenanceSchedule() {
	fmt.Println("\n=== Kiểm thử Cập nhật Lịch bảo trì ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID lịch bảo trì: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Nhập Device ID (UUID): ")
	deviceId, _ := reader.ReadString('\n')
	deviceId = cleanInput(deviceId)

	fmt.Print("Nhập loại bảo trì (cleaning/calibration/replacement/repair/inspection/software_update): ")
	maintenanceType, _ := reader.ReadString('\n')
	maintenanceType = cleanInput(maintenanceType)

	fmt.Print("Nhập danh mục bảo trì (preventive/corrective/emergency/routine): ")
	maintenanceCategory, _ := reader.ReadString('\n')
	maintenanceCategory = cleanInput(maintenanceCategory)

	fmt.Print("Nhập mức độ ưu tiên (low/medium/high/critical): ")
	priority, _ := reader.ReadString('\n')
	priority = cleanInput(priority)

	fmt.Print("Nhập ngày lên lịch (YYYY-MM-DD): ")
	scheduledDateStr, _ := reader.ReadString('\n')
	scheduledDateStr = cleanInput(scheduledDateStr)
	var scheduledDate *timestamppb.Timestamp
	if scheduledDateStr != "" {
		if t, err := time.Parse("2006-01-02", scheduledDateStr); err == nil {
			scheduledDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập thời gian ước tính (giờ): ")
	estimatedDurationStr, _ := reader.ReadString('\n')
	estimatedDurationStr = cleanInput(estimatedDurationStr)
	estimatedDuration := float64(2.0)
	if estimatedDurationStr != "" {
		if d, err := strconv.ParseFloat(estimatedDurationStr, 64); err == nil {
			estimatedDuration = d
		}
	}

	fmt.Print("Nhập ngày hoàn thành (YYYY-MM-DD): ")
	completedDateStr, _ := reader.ReadString('\n')
	completedDateStr = cleanInput(completedDateStr)
	var completedDate *timestamppb.Timestamp
	if completedDateStr != "" {
		if t, err := time.Parse("2006-01-02", completedDateStr); err == nil {
			completedDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập thời gian thực tế (giờ): ")
	actualDurationStr, _ := reader.ReadString('\n')
	actualDurationStr = cleanInput(actualDurationStr)
	actualDuration := float64(0.0)
	if actualDurationStr != "" {
		if d, err := strconv.ParseFloat(actualDurationStr, 64); err == nil {
			actualDuration = d
		}
	}

	fmt.Print("Nhập tên kỹ thuật viên: ")
	technician, _ := reader.ReadString('\n')
	technician = cleanInput(technician)

	fmt.Print("Nhập liên hệ kỹ thuật viên: ")
	technicianContact, _ := reader.ReadString('\n')
	technicianContact = cleanInput(technicianContact)

	fmt.Print("Nhập chi phí: ")
	costStr, _ := reader.ReadString('\n')
	costStr = cleanInput(costStr)
	cost := float64(100.0)
	if costStr != "" {
		if c, err := strconv.ParseFloat(costStr, 64); err == nil {
			cost = c
		}
	}

	fmt.Print("Nhập linh kiện thay thế: ")
	partsReplaced, _ := reader.ReadString('\n')
	partsReplaced = cleanInput(partsReplaced)

	fmt.Print("Nhập dụng cụ cần thiết: ")
	toolsRequired, _ := reader.ReadString('\n')
	toolsRequired = cleanInput(toolsRequired)

	fmt.Print("Nhập biện pháp an toàn: ")
	safetyPrecautions, _ := reader.ReadString('\n')
	safetyPrecautions = cleanInput(safetyPrecautions)

	fmt.Print("Nhập ghi chú trước bảo trì: ")
	preMaintenanceReadings, _ := reader.ReadString('\n')
	preMaintenanceReadings = cleanInput(preMaintenanceReadings)

	fmt.Print("Nhập ghi chú sau bảo trì: ")
	postMaintenanceReadings, _ := reader.ReadString('\n')
	postMaintenanceReadings = cleanInput(postMaintenanceReadings)

	fmt.Print("Nhập giá trị hiệu chuẩn: ")
	calibrationValues, _ := reader.ReadString('\n')
	calibrationValues = cleanInput(calibrationValues)

	fmt.Print("Nhập kết quả kiểm tra: ")
	testResults, _ := reader.ReadString('\n')
	testResults = cleanInput(testResults)

	fmt.Print("Nhập trạng thái mới (scheduled/in_progress/completed/cancelled/postponed): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)

	fmt.Print("Nhập đánh giá hoàn thành (1-5): ")
	ratingStr, _ := reader.ReadString('\n')
	ratingStr = cleanInput(ratingStr)
	rating := int32(0)
	if ratingStr != "" {
		if r, err := strconv.Atoi(ratingStr); err == nil {
			rating = int32(r)
		}
	}

	fmt.Print("Nhập ngày bảo trì tiếp theo (YYYY-MM-DD): ")
	nextMaintenanceDateStr, _ := reader.ReadString('\n')
	nextMaintenanceDateStr = cleanInput(nextMaintenanceDateStr)
	var nextMaintenanceDate *timestamppb.Timestamp
	if nextMaintenanceDateStr != "" {
		if t, err := time.Parse("2006-01-02", nextMaintenanceDateStr); err == nil {
			nextMaintenanceDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập khoảng cách bảo trì (ngày): ")
	intervalStr, _ := reader.ReadString('\n')
	intervalStr = cleanInput(intervalStr)
	interval := int32(30)
	if intervalStr != "" {
		if i, err := strconv.Atoi(intervalStr); err == nil {
			interval = int32(i)
		}
	}

	fmt.Print("Ảnh hưởng bảo hành (true/false): ")
	warrantyImpactStr, _ := reader.ReadString('\n')
	warrantyImpactStr = cleanInput(warrantyImpactStr)
	warrantyImpact := false
	if warrantyImpactStr == "true" {
		warrantyImpact = true
	}

	fmt.Print("Nhập thời gian ngừng hoạt động (phút): ")
	downtimeStr, _ := reader.ReadString('\n')
	downtimeStr = cleanInput(downtimeStr)
	downtime := int32(0)
	if downtimeStr != "" {
		if d, err := strconv.Atoi(downtimeStr); err == nil {
			downtime = int32(d)
		}
	}

	fmt.Print("Nhập ghi chú: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Nhập nhật ký bảo trì: ")
	maintenanceLog, _ := reader.ReadString('\n')
	maintenanceLog = cleanInput(maintenanceLog)

	fmt.Print("Nhập hình ảnh trước: ")
	beforeImages, _ := reader.ReadString('\n')
	beforeImages = cleanInput(beforeImages)

	fmt.Print("Nhập hình ảnh sau: ")
	afterImages, _ := reader.ReadString('\n')
	afterImages = cleanInput(afterImages)

	fmt.Print("Nhập người tạo (UUID): ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	// Convert image strings to slices
	var beforeImagesSlice []string
	if beforeImages != "" {
		beforeImagesSlice = []string{beforeImages}
	}
	var afterImagesSlice []string
	if afterImages != "" {
		afterImagesSlice = []string{afterImages}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.maintenanceScheduleClient.UpdateMaintenanceSchedule(ctx, &proto_maintenance_schedule.UpdateMaintenanceScheduleRequest{
		Id:                      id,
		DeviceId:                deviceId,
		MaintenanceType:         maintenanceType,
		MaintenanceCategory:     maintenanceCategory,
		Priority:                priority,
		ScheduledDate:           scheduledDate,
		EstimatedDurationHours:  estimatedDuration,
		CompletedDate:           completedDate,
		ActualDurationHours:     actualDuration,
		Technician:              technician,
		TechnicianContact:       technicianContact,
		Cost:                    cost,
		PartsReplaced:           partsReplaced,
		ToolsRequired:           toolsRequired,
		SafetyPrecautions:       safetyPrecautions,
		PreMaintenanceReadings:  preMaintenanceReadings,
		PostMaintenanceReadings: postMaintenanceReadings,
		CalibrationValues:       calibrationValues,
		TestResults:             testResults,
		Status:                  status,
		CompletionRating:        rating,
		NextMaintenanceDate:     nextMaintenanceDate,
		MaintenanceIntervalDays: interval,
		WarrantyImpact:          warrantyImpact,
		DowntimeMinutes:         downtime,
		Notes:                   notes,
		MaintenanceLog:          maintenanceLog,
		BeforeImages:            beforeImagesSlice,
		AfterImages:             afterImagesSlice,
		CreatedBy:               createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateMaintenanceSchedule: %v\n", err)
		return
	}

	fmt.Printf("Kết quả cập nhật lịch bảo trì:\n")
	if resp.MaintenanceSchedule != nil {
		fmt.Printf("ID: %s\n", resp.MaintenanceSchedule.Id)
		fmt.Printf("Device ID: %s\n", resp.MaintenanceSchedule.DeviceId)
		fmt.Printf("Maintenance Type: %s\n", resp.MaintenanceSchedule.MaintenanceType)
		fmt.Printf("Maintenance Category: %s\n", resp.MaintenanceSchedule.MaintenanceCategory)
		fmt.Printf("Priority: %s\n", resp.MaintenanceSchedule.Priority)
		fmt.Printf("Status: %s\n", resp.MaintenanceSchedule.Status)
		fmt.Printf("Estimated Duration: %.2f hours\n", resp.MaintenanceSchedule.EstimatedDurationHours)
		fmt.Printf("Actual Duration: %.2f hours\n", resp.MaintenanceSchedule.ActualDurationHours)
		fmt.Printf("Technician: %s\n", resp.MaintenanceSchedule.Technician)
		fmt.Printf("Cost: %.2f\n", resp.MaintenanceSchedule.Cost)
		fmt.Printf("Completion Rating: %d\n", resp.MaintenanceSchedule.CompletionRating)
		fmt.Printf("Maintenance Interval: %d days\n", resp.MaintenanceSchedule.MaintenanceIntervalDays)
		fmt.Printf("Warranty Impact: %t\n", resp.MaintenanceSchedule.WarrantyImpact)
		fmt.Printf("Downtime: %d minutes\n", resp.MaintenanceSchedule.DowntimeMinutes)
		fmt.Printf("Notes: %s\n", resp.MaintenanceSchedule.Notes)
		fmt.Printf("Created By: %s\n", resp.MaintenanceSchedule.CreatedBy)
	}
}

func (c *MaintenanceServiceClient) TestDeleteMaintenanceSchedule() {
	fmt.Println("\n=== Kiểm thử Xóa Lịch bảo trì ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID lịch bảo trì: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.maintenanceScheduleClient.DeleteMaintenanceSchedule(ctx, &proto_maintenance_schedule.DeleteMaintenanceScheduleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteMaintenanceSchedule: %v\n", err)
		return
	}

	fmt.Printf("Lịch bảo trì đã được xóa thành công!\n")
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== Ứng dụng kiểm thử gRPC Maintenance Service ===")
	fmt.Println("1. Dịch vụ Lịch Bảo trì")
	fmt.Println("0. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printMaintenanceScheduleMenu() {
	fmt.Println("\n=== Dịch vụ Lịch Bảo trì ===")
	fmt.Println("1. Tạo lịch bảo trì")
	fmt.Println("2. Lấy lịch bảo trì")
	fmt.Println("3. Liệt kê lịch bảo trì")
	fmt.Println("4. Cập nhật lịch bảo trì")
	fmt.Println("5. Xóa lịch bảo trì")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Đang kết nối tới máy chủ gRPC tại %s...\n", address)
	client, err := NewMaintenanceServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Kết nối thành công!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Dịch vụ Lịch Bảo trì
			for {
				printMaintenanceScheduleMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateMaintenanceSchedule()
				case "2":
					client.TestGetMaintenanceSchedule()
				case "3":
					client.TestListMaintenanceSchedules()
				case "4":
					client.TestUpdateMaintenanceSchedule()
				case "5":
					client.TestDeleteMaintenanceSchedule()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
	}
}
