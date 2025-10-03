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
	proto_harvest_record "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1"
	proto_pest_disease_record "github.com/anhvanhoa/sf-proto/gen/pest_disease_record/v1"
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

type CropServiceClient struct {
	harvestRecordClient     proto_harvest_record.HarvestRecordServiceClient
	pestDiseaseRecordClient proto_pest_disease_record.PestDiseaseRecordServiceClient
	conn                    *grpc.ClientConn
}

func NewCropServiceClient(address string) (*CropServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &CropServiceClient{
		harvestRecordClient:     proto_harvest_record.NewHarvestRecordServiceClient(conn),
		pestDiseaseRecordClient: proto_pest_disease_record.NewPestDiseaseRecordServiceClient(conn),
		conn:                    conn,
	}, nil
}

func (c *CropServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Harvest Record Service Tests ==================

func (c *CropServiceClient) TestCreateHarvestRecord() {
	fmt.Println("\n=== Kiểm thử Tạo bản ghi Thu hoạch ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID chu kỳ trồng: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Nhập ngày thu hoạch (YYYY-MM-DD): ")
	harvestDateStr, _ := reader.ReadString('\n')
	harvestDateStr = cleanInput(harvestDateStr)
	var harvestDate *timestamppb.Timestamp
	if harvestDateStr != "" {
		if t, err := time.Parse("2006-01-02", harvestDateStr); err == nil {
			harvestDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập khối lượng (kg): ")
	quantityStr, _ := reader.ReadString('\n')
	quantityStr = cleanInput(quantityStr)
	quantity := float64(100.0)
	if quantityStr != "" {
		if q, err := strconv.ParseFloat(quantityStr, 64); err == nil {
			quantity = q
		}
	}

	fmt.Print("Nhập hạng chất lượng: ")
	qualityGrade, _ := reader.ReadString('\n')
	qualityGrade = cleanInput(qualityGrade)

	fmt.Print("Nhập phân loại kích cỡ: ")
	sizeClassification, _ := reader.ReadString('\n')
	sizeClassification = cleanInput(sizeClassification)

	fmt.Print("Nhập giá thị trường mỗi kg: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = cleanInput(priceStr)
	price := float64(10.0)
	if priceStr != "" {
		if p, err := strconv.ParseFloat(priceStr, 64); err == nil {
			price = p
		}
	}

	fmt.Print("Nhập số giờ lao động: ")
	laborHoursStr, _ := reader.ReadString('\n')
	laborHoursStr = cleanInput(laborHoursStr)
	laborHours := float64(8.0)
	if laborHoursStr != "" {
		if h, err := strconv.ParseFloat(laborHoursStr, 64); err == nil {
			laborHours = h
		}
	}

	fmt.Print("Nhập chi phí lao động: ")
	laborCostStr, _ := reader.ReadString('\n')
	laborCostStr = cleanInput(laborCostStr)
	laborCost := float64(50.0)
	if laborCostStr != "" {
		if c, err := strconv.ParseFloat(laborCostStr, 64); err == nil {
			laborCost = c
		}
	}

	fmt.Print("Nhập chi phí đóng gói: ")
	packagingCostStr, _ := reader.ReadString('\n')
	packagingCostStr = cleanInput(packagingCostStr)
	packagingCost := float64(10.0)
	if packagingCostStr != "" {
		if c, err := strconv.ParseFloat(packagingCostStr, 64); err == nil {
			packagingCost = c
		}
	}

	fmt.Print("Nhập địa điểm lưu trữ: ")
	storageLocation, _ := reader.ReadString('\n')
	storageLocation = cleanInput(storageLocation)

	fmt.Print("Nhập nhiệt độ lưu trữ: ")
	tempStr, _ := reader.ReadString('\n')
	tempStr = cleanInput(tempStr)
	temperature := float64(4.0)
	if tempStr != "" {
		if t, err := strconv.ParseFloat(tempStr, 64); err == nil {
			temperature = t
		}
	}

	fmt.Print("Nhập thông tin người mua: ")
	buyerInfo, _ := reader.ReadString('\n')
	buyerInfo = cleanInput(buyerInfo)

	fmt.Print("Nhập thời tiết lúc thu hoạch: ")
	weather, _ := reader.ReadString('\n')
	weather = cleanInput(weather)

	fmt.Print("Nhập đánh giá sức khỏe cây (1-5): ")
	ratingStr, _ := reader.ReadString('\n')
	ratingStr = cleanInput(ratingStr)
	rating := int32(5)
	if ratingStr != "" {
		if r, err := strconv.Atoi(ratingStr); err == nil {
			rating = int32(r)
		}
	}

	fmt.Print("Nhập ghi chú: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Nhập hình ảnh VD: [id1, id2, id3]: ")
	images, _ := reader.ReadString('\n')
	images = cleanInput(images)

	fmt.Print("Nhập người tạo: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.CreateHarvestRecord(ctx, &proto_harvest_record.CreateHarvestRecordRequest{
		PlantingCycleId:    plantingCycleId,
		HarvestDate:        harvestDate,
		QuantityKg:         quantity,
		QualityGrade:       qualityGrade,
		SizeClassification: sizeClassification,
		MarketPricePerKg:   price,
		LaborHours:         laborHours,
		LaborCost:          laborCost,
		PackagingCost:      packagingCost,
		StorageLocation:    storageLocation,
		StorageTemperature: temperature,
		BuyerInformation:   buyerInfo,
		WeatherAtHarvest:   weather,
		PlantHealthRating:  rating,
		Notes:              notes,
		Images:             images,
		CreatedBy:          createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreateHarvestRecord: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo bản ghi thu hoạch:\n")
	if resp.HarvestRecord != nil {
		fmt.Printf("ID: %s\n", resp.HarvestRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.HarvestRecord.PlantingCycleId)
		fmt.Printf("Quantity: %.2f kg\n", resp.HarvestRecord.QuantityKg)
		fmt.Printf("Quality Grade: %s\n", resp.HarvestRecord.QualityGrade)
		fmt.Printf("Size Classification: %s\n", resp.HarvestRecord.SizeClassification)
		fmt.Printf("Market Price: %.2f per kg\n", resp.HarvestRecord.MarketPricePerKg)
		fmt.Printf("Total Revenue: %.2f\n", resp.HarvestRecord.TotalRevenue)
	}
}

func (c *CropServiceClient) TestGetHarvestRecord() {
	fmt.Println("\n=== Kiểm thử Lấy bản ghi Thu hoạch ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID bản ghi thu hoạch: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.GetHarvestRecord(ctx, &proto_harvest_record.GetHarvestRecordRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetHarvestRecord: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy bản ghi thu hoạch:\n")
	if resp.HarvestRecord != nil {
		fmt.Printf("ID: %s\n", resp.HarvestRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.HarvestRecord.PlantingCycleId)
		fmt.Printf("Khối lượng: %.2f kg\n", resp.HarvestRecord.QuantityKg)
		fmt.Printf("Hạng chất lượng: %s\n", resp.HarvestRecord.QualityGrade)
		fmt.Printf("Phân loại kích cỡ: %s\n", resp.HarvestRecord.SizeClassification)
		fmt.Printf("Giá thị trường: %.2f mỗi kg\n", resp.HarvestRecord.MarketPricePerKg)
		fmt.Printf("Tổng doanh thu: %.2f\n", resp.HarvestRecord.TotalRevenue)
		fmt.Printf("Số giờ lao động: %.2f\n", resp.HarvestRecord.LaborHours)
		fmt.Printf("Chi phí lao động: %.2f\n", resp.HarvestRecord.LaborCost)
		fmt.Printf("Chi phí đóng gói: %.2f\n", resp.HarvestRecord.PackagingCost)
		fmt.Printf("Địa điểm lưu trữ: %s\n", resp.HarvestRecord.StorageLocation)
		fmt.Printf("Nhiệt độ lưu trữ: %.2f°C\n", resp.HarvestRecord.StorageTemperature)
		fmt.Printf("Thông tin người mua: %s\n", resp.HarvestRecord.BuyerInformation)
		fmt.Printf("Thời tiết lúc thu hoạch: %s\n", resp.HarvestRecord.WeatherAtHarvest)
		fmt.Printf("Đánh giá sức khỏe cây: %d\n", resp.HarvestRecord.PlantHealthRating)
		fmt.Printf("Ghi chú: %s\n", resp.HarvestRecord.Notes)
		fmt.Printf("Người tạo: %s\n", resp.HarvestRecord.CreatedBy)
	}
}

func (c *CropServiceClient) TestListHarvestRecords() {
	fmt.Println("\n=== Kiểm thử Liệt kê Bản ghi Thu hoạch ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.ListHarvestRecords(ctx, &proto_harvest_record.ListHarvestRecordsRequest{
		Pagination: &proto_common.PaginationRequest{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "created_at",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListHarvestRecords: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê bản ghi thu hoạch:\n")
	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	fmt.Printf("Danh sách bản ghi thu hoạch:\n")
	for i, record := range resp.HarvestRecords {
		fmt.Printf("  [%d] ID: %s, Chu kỳ trồng: %s, Khối lượng: %.2f kg, Chất lượng: %s\n",
			i+1, record.Id, record.PlantingCycleId, record.QuantityKg, record.QualityGrade)
	}
}

func (c *CropServiceClient) TestGetHarvestRecordsByPlantingCycle() {
	fmt.Println("\n=== Kiểm thử Lấy Bản ghi Thu hoạch theo Chu kỳ trồng ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID chu kỳ trồng: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.harvestRecordClient.GetHarvestRecordsByPlantingCycle(ctx, &proto_harvest_record.GetHarvestRecordsByPlantingCycleRequest{
		PlantingCycleId: plantingCycleId,
		Pagination: &proto_common.PaginationRequest{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "harvest_date",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling GetHarvestRecordsByPlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy bản ghi thu hoạch theo chu kỳ trồng:\n")
	fmt.Printf("Tổng số: %d\n", resp.Total)
	fmt.Printf("Danh sách bản ghi thu hoạch:\n")
	for i, record := range resp.HarvestRecords {
		fmt.Printf("  [%d] ID: %s, Khối lượng: %.2f kg, Chất lượng: %s, Doanh thu: %.2f\n",
			i+1, record.Id, record.QuantityKg, record.QualityGrade, record.TotalRevenue)
	}
}

// ================== Pest Disease Record Service Tests ==================

func (c *CropServiceClient) TestCreatePestDiseaseRecord() {
	fmt.Println("\n=== Kiểm thử Tạo Bản ghi Sâu bệnh ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID chu kỳ trồng: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	fmt.Print("Nhập loại (pest/disease/nutrient_deficiency/environmental_stress): ")
	recordType, _ := reader.ReadString('\n')
	recordType = cleanInput(recordType)

	fmt.Print("Nhập tên: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Nhập tên khoa học: ")
	scientificName, _ := reader.ReadString('\n')
	scientificName = cleanInput(scientificName)

	fmt.Print("Nhập mức độ (low/medium/high/critical): ")
	severity, _ := reader.ReadString('\n')
	severity = cleanInput(severity)

	fmt.Print("Nhập tỷ lệ diện tích bị ảnh hưởng (%): ")
	areaStr, _ := reader.ReadString('\n')
	areaStr = cleanInput(areaStr)
	area := float64(10.0)
	if areaStr != "" {
		if a, err := strconv.ParseFloat(areaStr, 64); err == nil {
			area = a
		}
	}

	fmt.Print("Nhập số lượng cây bị ảnh hưởng: ")
	countStr, _ := reader.ReadString('\n')
	countStr = cleanInput(countStr)
	count := int32(10)
	if countStr != "" {
		if c, err := strconv.Atoi(countStr); err == nil {
			count = int32(c)
		}
	}

	fmt.Print("Nhập ngày phát hiện (YYYY-MM-DD): ")
	detectionDateStr, _ := reader.ReadString('\n')
	detectionDateStr = cleanInput(detectionDateStr)
	var detectionDate *timestamppb.Timestamp
	if detectionDateStr != "" {
		if t, err := time.Parse("2006-01-02", detectionDateStr); err == nil {
			detectionDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập phương pháp phát hiện: ")
	detectionMethod, _ := reader.ReadString('\n')
	detectionMethod = cleanInput(detectionMethod)

	fmt.Print("Nhập triệu chứng: ")
	symptoms, _ := reader.ReadString('\n')
	symptoms = cleanInput(symptoms)

	fmt.Print("Nhập biện pháp điều trị đã áp dụng: ")
	treatment, _ := reader.ReadString('\n')
	treatment = cleanInput(treatment)

	fmt.Print("Nhập ngày điều trị (YYYY-MM-DD): ")
	treatmentDateStr, _ := reader.ReadString('\n')
	treatmentDateStr = cleanInput(treatmentDateStr)
	var treatmentDate *timestamppb.Timestamp
	if treatmentDateStr != "" {
		if t, err := time.Parse("2006-01-02", treatmentDateStr); err == nil {
			treatmentDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập chi phí điều trị: ")
	costStr, _ := reader.ReadString('\n')
	costStr = cleanInput(costStr)
	cost := float64(50.0)
	if costStr != "" {
		if c, err := strconv.ParseFloat(costStr, 64); err == nil {
			cost = c
		}
	}

	fmt.Print("Nhập thời gian điều trị (ngày): ")
	durationStr, _ := reader.ReadString('\n')
	durationStr = cleanInput(durationStr)
	duration := int32(7)
	if durationStr != "" {
		if d, err := strconv.Atoi(durationStr); err == nil {
			duration = int32(d)
		}
	}

	fmt.Print("Nhập trạng thái phục hồi: ")
	recoveryStatus, _ := reader.ReadString('\n')
	recoveryStatus = cleanInput(recoveryStatus)

	fmt.Print("Nhập đánh giá hiệu quả (1-5): ")
	ratingStr, _ := reader.ReadString('\n')
	ratingStr = cleanInput(ratingStr)
	rating := int32(4)
	if ratingStr != "" {
		if r, err := strconv.Atoi(ratingStr); err == nil {
			rating = int32(r)
		}
	}

	fmt.Print("Nhập ngày theo dõi (YYYY-MM-DD): ")
	followUpDateStr, _ := reader.ReadString('\n')
	followUpDateStr = cleanInput(followUpDateStr)
	var followUpDate *timestamppb.Timestamp
	if followUpDateStr != "" {
		if t, err := time.Parse("2006-01-02", followUpDateStr); err == nil {
			followUpDate = timestamppb.New(t)
		}
	}

	fmt.Print("Nhập biện pháp phòng ngừa: ")
	prevention, _ := reader.ReadString('\n')
	prevention = cleanInput(prevention)

	fmt.Print("Nhập yếu tố môi trường: ")
	environmental, _ := reader.ReadString('\n')
	environmental = cleanInput(environmental)

	fmt.Print("Nhập hình ảnh: ")
	images, _ := reader.ReadString('\n')
	images = cleanInput(images)

	fmt.Print("Nhập ghi chú: ")
	notes, _ := reader.ReadString('\n')
	notes = cleanInput(notes)

	fmt.Print("Nhập người tạo: ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.CreatePestDiseaseRecord(ctx, &proto_pest_disease_record.CreatePestDiseaseRecordRequest{
		PlantingCycleId:        plantingCycleId,
		Type:                   recordType,
		Name:                   name,
		ScientificName:         scientificName,
		Severity:               severity,
		AffectedAreaPercentage: area,
		AffectedPlantCount:     count,
		DetectionDate:          detectionDate,
		DetectionMethod:        detectionMethod,
		Symptoms:               symptoms,
		TreatmentApplied:       treatment,
		TreatmentDate:          treatmentDate,
		TreatmentCost:          cost,
		TreatmentDurationDays:  duration,
		RecoveryStatus:         recoveryStatus,
		EffectivenessRating:    rating,
		FollowUpDate:           followUpDate,
		PreventionMeasures:     prevention,
		EnvironmentalFactors:   environmental,
		Images:                 images,
		Notes:                  notes,
		CreatedBy:              createdBy,
	})
	if err != nil {
		fmt.Printf("Error calling CreatePestDiseaseRecord: %v\n", err)
		return
	}

	fmt.Printf("Kết quả tạo bản ghi sâu bệnh:\n")
	if resp.PestDiseaseRecord != nil {
		fmt.Printf("ID: %s\n", resp.PestDiseaseRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PestDiseaseRecord.PlantingCycleId)
		fmt.Printf("Type: %s\n", resp.PestDiseaseRecord.Type)
		fmt.Printf("Name: %s\n", resp.PestDiseaseRecord.Name)
		fmt.Printf("Scientific Name: %s\n", resp.PestDiseaseRecord.ScientificName)
		fmt.Printf("Severity: %s\n", resp.PestDiseaseRecord.Severity)
		fmt.Printf("Affected Area: %.2f%%\n", resp.PestDiseaseRecord.AffectedAreaPercentage)
		fmt.Printf("Affected Plants: %d\n", resp.PestDiseaseRecord.AffectedPlantCount)
		fmt.Printf("Detection Method: %s\n", resp.PestDiseaseRecord.DetectionMethod)
		fmt.Printf("Treatment Applied: %s\n", resp.PestDiseaseRecord.TreatmentApplied)
		fmt.Printf("Treatment Cost: %.2f\n", resp.PestDiseaseRecord.TreatmentCost)
		fmt.Printf("Recovery Status: %s\n", resp.PestDiseaseRecord.RecoveryStatus)
		fmt.Printf("Effectiveness Rating: %d\n", resp.PestDiseaseRecord.EffectivenessRating)
	}
}

func (c *CropServiceClient) TestGetPestDiseaseRecord() {
	fmt.Println("\n=== Kiểm thử Lấy Bản ghi Sâu bệnh ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID bản ghi sâu bệnh: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.GetPestDiseaseRecord(ctx, &proto_pest_disease_record.GetPestDiseaseRecordRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetPestDiseaseRecord: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy bản ghi sâu bệnh:\n")
	if resp.PestDiseaseRecord != nil {
		fmt.Printf("ID: %s\n", resp.PestDiseaseRecord.Id)
		fmt.Printf("Planting Cycle ID: %s\n", resp.PestDiseaseRecord.PlantingCycleId)
		fmt.Printf("Loại: %s\n", resp.PestDiseaseRecord.Type)
		fmt.Printf("Tên: %s\n", resp.PestDiseaseRecord.Name)
		fmt.Printf("Tên khoa học: %s\n", resp.PestDiseaseRecord.ScientificName)
		fmt.Printf("Mức độ: %s\n", resp.PestDiseaseRecord.Severity)
		fmt.Printf("Diện tích bị ảnh hưởng: %.2f%%\n", resp.PestDiseaseRecord.AffectedAreaPercentage)
		fmt.Printf("Số cây bị ảnh hưởng: %d\n", resp.PestDiseaseRecord.AffectedPlantCount)
		fmt.Printf("Phương pháp phát hiện: %s\n", resp.PestDiseaseRecord.DetectionMethod)
		fmt.Printf("Symptoms: %s\n", resp.PestDiseaseRecord.Symptoms)
		fmt.Printf("Biện pháp điều trị: %s\n", resp.PestDiseaseRecord.TreatmentApplied)
		fmt.Printf("Chi phí điều trị: %.2f\n", resp.PestDiseaseRecord.TreatmentCost)
		fmt.Printf("Treatment Duration: %d days\n", resp.PestDiseaseRecord.TreatmentDurationDays)
		fmt.Printf("Trạng thái phục hồi: %s\n", resp.PestDiseaseRecord.RecoveryStatus)
		fmt.Printf("Đánh giá hiệu quả: %d\n", resp.PestDiseaseRecord.EffectivenessRating)
		fmt.Printf("Biện pháp phòng ngừa: %s\n", resp.PestDiseaseRecord.PreventionMeasures)
		fmt.Printf("Yếu tố môi trường: %s\n", resp.PestDiseaseRecord.EnvironmentalFactors)
		fmt.Printf("Ghi chú: %s\n", resp.PestDiseaseRecord.Notes)
		fmt.Printf("Người tạo: %s\n", resp.PestDiseaseRecord.CreatedBy)
	}
}

func (c *CropServiceClient) TestListPestDiseaseRecords() {
	fmt.Println("\n=== Kiểm thử Liệt kê Bản ghi Sâu bệnh ===")

	reader := bufio.NewReader(os.Stdin)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.ListPestDiseaseRecords(ctx, &proto_pest_disease_record.ListPestDiseaseRecordsRequest{
		Pagination: &proto_pest_disease_record.Pagination{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "detection_date",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListPestDiseaseRecords: %v\n", err)
		return
	}

	fmt.Printf("Kết quả liệt kê bản ghi sâu bệnh:\n")
	fmt.Printf("Tổng số: %d\n", resp.Total)
	fmt.Printf("Danh sách bản ghi sâu bệnh:\n")
	for i, record := range resp.PestDiseaseRecords {
		fmt.Printf("  [%d] ID: %s, Loại: %s, Tên: %s, Mức độ: %s\n",
			i+1, record.Id, record.Type, record.Name, record.Severity)
	}
}

func (c *CropServiceClient) TestGetPestDiseaseRecordsByPlantingCycle() {
	fmt.Println("\n=== Kiểm thử Lấy Bản ghi Sâu bệnh theo Chu kỳ trồng ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID chu kỳ trồng: ")
	plantingCycleId, _ := reader.ReadString('\n')
	plantingCycleId = cleanInput(plantingCycleId)

	offset, limit := inputPaging(reader)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.pestDiseaseRecordClient.GetPestDiseaseRecordsByPlantingCycle(ctx, &proto_pest_disease_record.GetPestDiseaseRecordsByPlantingCycleRequest{
		PlantingCycleId: plantingCycleId,
		Pagination: &proto_pest_disease_record.Pagination{
			Page:      offset,
			PageSize:  limit,
			SortBy:    "detection_date",
			SortOrder: "desc",
		},
	})
	if err != nil {
		fmt.Printf("Error calling GetPestDiseaseRecordsByPlantingCycle: %v\n", err)
		return
	}

	fmt.Printf("Kết quả lấy bản ghi sâu bệnh theo chu kỳ trồng:\n")
	fmt.Printf("Tổng số: %d\n", resp.Total)
	fmt.Printf("Danh sách bản ghi sâu bệnh:\n")
	for i, record := range resp.PestDiseaseRecords {
		fmt.Printf("  [%d] ID: %s, Loại: %s, Tên: %s, Mức độ: %s, Trạng thái: %s\n",
			i+1, record.Id, record.Type, record.Name, record.Severity, record.RecoveryStatus)
	}
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== Ứng dụng kiểm thử gRPC Production Service ===")
	fmt.Println("1. Dịch vụ Bản ghi Thu hoạch")
	fmt.Println("2. Dịch vụ Bản ghi Sâu bệnh")
	fmt.Println("0. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printHarvestRecordMenu() {
	fmt.Println("\n=== Dịch vụ Bản ghi Thu hoạch ===")
	fmt.Println("1. Tạo bản ghi thu hoạch")
	fmt.Println("2. Lấy bản ghi thu hoạch")
	fmt.Println("3. Liệt kê bản ghi thu hoạch")
	fmt.Println("4. Lấy bản ghi theo chu kỳ trồng")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printPestDiseaseRecordMenu() {
	fmt.Println("\n=== Dịch vụ Bản ghi Sâu bệnh ===")
	fmt.Println("1. Tạo bản ghi sâu bệnh")
	fmt.Println("2. Lấy bản ghi sâu bệnh")
	fmt.Println("3. Liệt kê bản ghi sâu bệnh")
	fmt.Println("4. Lấy bản ghi theo chu kỳ trồng")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Đang kết nối tới máy chủ gRPC tại %s...\n", address)
	client, err := NewCropServiceClient(address)
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
			// Dịch vụ Bản ghi Thu hoạch
			for {
				printHarvestRecordMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateHarvestRecord()
				case "2":
					client.TestGetHarvestRecord()
				case "3":
					client.TestListHarvestRecords()
				case "4":
					client.TestGetHarvestRecordsByPlantingCycle()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "2":
			// Dịch vụ Bản ghi Sâu bệnh
			for {
				printPestDiseaseRecordMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreatePestDiseaseRecord()
				case "2":
					client.TestGetPestDiseaseRecord()
				case "3":
					client.TestListPestDiseaseRecords()
				case "4":
					client.TestGetPestDiseaseRecordsByPlantingCycle()
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
