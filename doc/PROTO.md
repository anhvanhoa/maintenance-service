syntax = "proto3";

package harvest_record.v1;

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";
import "common/v1/common.proto";

option go_package = "github.com/anhvanhoa/sf-proto/gen/harvest_record/v1;proto_harvest_record";

service HarvestRecordService {
  rpc CreateHarvestRecord(CreateHarvestRecordRequest)
      returns (CreateHarvestRecordResponse);
  rpc GetHarvestRecord(GetHarvestRecordRequest)
      returns (GetHarvestRecordResponse);
  rpc UpdateHarvestRecord(UpdateHarvestRecordRequest)
      returns (UpdateHarvestRecordResponse);
  rpc DeleteHarvestRecord(DeleteHarvestRecordRequest)
      returns (DeleteHarvestRecordResponse);
  rpc ListHarvestRecords(ListHarvestRecordsRequest)
      returns (ListHarvestRecordsResponse);
  rpc GetHarvestRecordsByPlantingCycle(GetHarvestRecordsByPlantingCycleRequest)
      returns (GetHarvestRecordsByPlantingCycleResponse);
}

message HarvestRecord {
  string id = 1;
  string planting_cycle_id = 2;
  google.protobuf.Timestamp harvest_date = 3;
  google.protobuf.Timestamp harvest_time = 4;
  double quantity_kg = 5;
  string quality_grade = 6;
  string size_classification = 7;
  double market_price_per_kg = 8;
  double total_revenue = 9;
  double labor_hours = 10;
  double labor_cost = 11;
  double packaging_cost = 12;
  string storage_location = 13;
  double storage_temperature = 14;
  string buyer_information = 15;
  google.protobuf.Timestamp delivery_date = 16;
  string weather_at_harvest = 17;
  int32 plant_health_rating = 18;
  string notes = 19;
  string images = 20;
  string created_by = 21;
  google.protobuf.Timestamp created_at = 22;
  google.protobuf.Timestamp updated_at = 23;
}

message CreateHarvestRecordRequest {
  string planting_cycle_id = 1 [ (buf.validate.field).string.min_len = 1 ];
  google.protobuf.Timestamp harvest_date = 2;
  google.protobuf.Timestamp harvest_time = 3;
  double quantity_kg = 4 [ (buf.validate.field).double.gt = 0 ];
  string quality_grade = 5 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  string size_classification = 6 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  double market_price_per_kg = 7 [ (buf.validate.field).double.gte = 0 ];
  double labor_hours = 8 [ (buf.validate.field).double.gte = 0 ];
  double labor_cost = 9 [ (buf.validate.field).double.gte = 0 ];
  double packaging_cost = 10 [ (buf.validate.field).double.gte = 0 ];
  string storage_location = 11 [ (buf.validate.field).string.max_len = 200 ];
  double storage_temperature = 12
      [ (buf.validate.field).double.gte = -50, (buf.validate.field).double.lte = 50 ];
  string buyer_information = 13 [ (buf.validate.field).string.max_len = 500 ];
  google.protobuf.Timestamp delivery_date = 14;
  string weather_at_harvest = 15 [ (buf.validate.field).string.max_len = 100 ];
  int32 plant_health_rating = 16
      [ (buf.validate.field).int32.gte = 1, (buf.validate.field).int32.lte = 10 ];
  string notes = 17 [ (buf.validate.field).string.max_len = 1000 ];
  string images = 18 [ (buf.validate.field).string.max_len = 2000 ];
  string created_by = 19 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
}

message CreateHarvestRecordResponse { HarvestRecord harvest_record = 1; }

message GetHarvestRecordRequest {
  string id = 1 [ (buf.validate.field).string.min_len = 1 ];
}

message GetHarvestRecordResponse { HarvestRecord harvest_record = 1; }

message UpdateHarvestRecordRequest {
  string id = 1 [ (buf.validate.field).string.min_len = 1 ];
  google.protobuf.Timestamp harvest_date = 2;
  google.protobuf.Timestamp harvest_time = 3;
  double quantity_kg = 4 [ (buf.validate.field).double.gt = 0 ];
  string quality_grade = 5 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  string size_classification = 6 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  double market_price_per_kg = 7 [ (buf.validate.field).double.gte = 0 ];
  double labor_hours = 8 [ (buf.validate.field).double.gte = 0 ];
  double labor_cost = 9 [ (buf.validate.field).double.gte = 0 ];
  double packaging_cost = 10 [ (buf.validate.field).double.gte = 0 ];
  string storage_location = 11 [ (buf.validate.field).string.max_len = 200 ];
  double storage_temperature = 12
      [ (buf.validate.field).double.gte = -50, (buf.validate.field).double.lte = 50 ];
  string buyer_information = 13 [ (buf.validate.field).string.max_len = 500 ];
  google.protobuf.Timestamp delivery_date = 14;
  string weather_at_harvest = 15 [ (buf.validate.field).string.max_len = 100 ];
  int32 plant_health_rating = 16
      [ (buf.validate.field).int32.gte = 1, (buf.validate.field).int32.lte = 10 ];
  string notes = 17 [ (buf.validate.field).string.max_len = 1000 ];
  string images = 18 [ (buf.validate.field).string.max_len = 2000 ];
}

message UpdateHarvestRecordResponse { HarvestRecord harvest_record = 1; }

message DeleteHarvestRecordRequest {
  string id = 1 [ (buf.validate.field).string.min_len = 1 ];
}

message DeleteHarvestRecordResponse { string message = 1; }

message ListHarvestRecordsRequest {
  common.PaginationRequest pagination = 1;
  HarvestRecordFilter filter = 2;
}

message ListHarvestRecordsResponse {
  repeated HarvestRecord harvest_records = 1;
  common.PaginationResponse pagination = 2;
}

message GetHarvestRecordsByPlantingCycleRequest {
  string planting_cycle_id = 1 [ (buf.validate.field).string.min_len = 1 ];
  common.PaginationRequest pagination = 2;
  HarvestRecordFilter filter = 3;
}

message GetHarvestRecordsByPlantingCycleResponse {
  repeated HarvestRecord harvest_records = 1;
  int64 total = 2;
}

message HarvestRecordFilter {
  google.protobuf.Timestamp harvest_date = 1;
  string quality_grade = 2;
  string size_classification = 3;
  double market_price_per_kg = 4;
  double total_revenue = 5;
  int32 plant_health_rating = 6;
  string notes = 7;
  string images = 8;
  string created_by = 9;
  google.protobuf.Timestamp created_at = 10;
}


-------------------------------------------


syntax = "proto3";

package production_service.pest_disease_record;

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";

option go_package = "production_service/proto/pest_disease_record";

service PestDiseaseRecordService {
  rpc CreatePestDiseaseRecord(CreatePestDiseaseRecordRequest)
      returns (CreatePestDiseaseRecordResponse);
  rpc GetPestDiseaseRecord(GetPestDiseaseRecordRequest)
      returns (GetPestDiseaseRecordResponse);
  rpc UpdatePestDiseaseRecord(UpdatePestDiseaseRecordRequest)
      returns (UpdatePestDiseaseRecordResponse);
  rpc DeletePestDiseaseRecord(DeletePestDiseaseRecordRequest)
      returns (DeletePestDiseaseRecordResponse);
  rpc ListPestDiseaseRecords(ListPestDiseaseRecordsRequest)
      returns (ListPestDiseaseRecordsResponse);
  rpc GetPestDiseaseRecordsByPlantingCycle(
      GetPestDiseaseRecordsByPlantingCycleRequest)
      returns (GetPestDiseaseRecordsByPlantingCycleResponse);
}

message PestDiseaseRecord {
  string id = 1;
  string planting_cycle_id = 2;
  string type = 3;
  string name = 4;
  string scientific_name = 5;
  string severity = 6;
  double affected_area_percentage = 7;
  int32 affected_plant_count = 8;
  google.protobuf.Timestamp detection_date = 9;
  string detection_method = 10;
  string symptoms = 11;
  string treatment_applied = 12;
  google.protobuf.Timestamp treatment_date = 13;
  double treatment_cost = 14;
  int32 treatment_duration_days = 15;
  string recovery_status = 16;
  int32 effectiveness_rating = 17;
  google.protobuf.Timestamp follow_up_date = 18;
  string prevention_measures = 19;
  string environmental_factors = 20;
  string images = 21;
  string notes = 22;
  string created_by = 23;
  google.protobuf.Timestamp created_at = 24;
  google.protobuf.Timestamp updated_at = 25;
}

message CreatePestDiseaseRecordRequest {
  string planting_cycle_id = 1 [ (buf.validate.field).string.min_len = 1 ];
  string type = 2 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  string name = 3 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string scientific_name = 4 [ (buf.validate.field).string.max_len = 200 ];
  string severity = 5 [
    (buf.validate.field).string.in = "low",
    (buf.validate.field).string.in = "medium",
    (buf.validate.field).string.in = "high",
    (buf.validate.field).string.in = "critical"
  ];
  double affected_area_percentage = 6 [
    (buf.validate.field).double.gte = 0,
    (buf.validate.field).double.lte = 100
  ];
  int32 affected_plant_count = 7 [ (buf.validate.field).int32.gte = 0 ];
  google.protobuf.Timestamp detection_date = 8;
  string detection_method = 9 [ (buf.validate.field).string.max_len = 200 ];
  string symptoms = 10 [ (buf.validate.field).string.max_len = 1000 ];
  string treatment_applied = 11 [ (buf.validate.field).string.max_len = 1000 ];
  google.protobuf.Timestamp treatment_date = 12;
  double treatment_cost = 13 [ (buf.validate.field).double.gte = 0 ];
  int32 treatment_duration_days = 14 [
    (buf.validate.field).int32.gte = 0,
    (buf.validate.field).int32.lte = 365
  ];
  string recovery_status = 15 [
    (buf.validate.field).string.in = "recovered",
    (buf.validate.field).string.in = "recovering",
    (buf.validate.field).string.in = "not_recovered",
    (buf.validate.field).string.in = "unknown"
  ];
  int32 effectiveness_rating = 16 [
    (buf.validate.field).int32.gte = 1,
    (buf.validate.field).int32.lte = 10
  ];
  google.protobuf.Timestamp follow_up_date = 17;
  string prevention_measures = 18
      [ (buf.validate.field).string.max_len = 1000 ];
  string environmental_factors = 19
      [ (buf.validate.field).string.max_len = 500 ];
  string images = 20 [ (buf.validate.field).string.max_len = 2000 ];
  string notes = 21 [ (buf.validate.field).string.max_len = 2000 ];
  string created_by = 22 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
}

message CreatePestDiseaseRecordResponse {
  PestDiseaseRecord pest_disease_record = 1;
}

message GetPestDiseaseRecordRequest {
  string id = 1 [ (buf.validate.field).string.min_len = 1 ];
}

message GetPestDiseaseRecordResponse {
  PestDiseaseRecord pest_disease_record = 1;
}

message UpdatePestDiseaseRecordRequest {
  string id = 1 [ (buf.validate.field).string.min_len = 1 ];
  string type = 2 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 50
  ];
  string name = 3 [
    (buf.validate.field).string.min_len = 1,
    (buf.validate.field).string.max_len = 100
  ];
  string scientific_name = 4 [ (buf.validate.field).string.max_len = 200 ];
  string severity = 5 [
    (buf.validate.field).string.in = "low",
    (buf.validate.field).string.in = "medium",
    (buf.validate.field).string.in = "high",
    (buf.validate.field).string.in = "critical"
  ];
  double affected_area_percentage = 6 [
    (buf.validate.field).double.gte = 0,
    (buf.validate.field).double.lte = 100
  ];
  int32 affected_plant_count = 7 [ (buf.validate.field).int32.gte = 0 ];
  google.protobuf.Timestamp detection_date = 8;
  string detection_method = 9 [ (buf.validate.field).string.max_len = 200 ];
  string symptoms = 10 [ (buf.validate.field).string.max_len = 1000 ];
  string treatment_applied = 11 [ (buf.validate.field).string.max_len = 1000 ];
  google.protobuf.Timestamp treatment_date = 12;
  double treatment_cost = 13 [ (buf.validate.field).double.gte = 0 ];
  int32 treatment_duration_days = 14 [
    (buf.validate.field).int32.gte = 0,
    (buf.validate.field).int32.lte = 365
  ];
  string recovery_status = 15 [
    (buf.validate.field).string.in = "recovered",
    (buf.validate.field).string.in = "recovering",
    (buf.validate.field).string.in = "not_recovered",
    (buf.validate.field).string.in = "unknown"
  ];
  int32 effectiveness_rating = 16 [
    (buf.validate.field).int32.gte = 1,
    (buf.validate.field).int32.lte = 10
  ];
  google.protobuf.Timestamp follow_up_date = 17;
  string prevention_measures = 18
      [ (buf.validate.field).string.max_len = 1000 ];
  string environmental_factors = 19
      [ (buf.validate.field).string.max_len = 500 ];
  string images = 20 [ (buf.validate.field).string.max_len = 2000 ];
  string notes = 21 [ (buf.validate.field).string.max_len = 2000 ];
}

message UpdatePestDiseaseRecordResponse {
  PestDiseaseRecord pest_disease_record = 1;
}

message DeletePestDiseaseRecordRequest {
  string id = 1 [ (buf.validate.field).string.min_len = 1 ];
}

message DeletePestDiseaseRecordResponse {
  string message = 1;
}

message ListPestDiseaseRecordsRequest {
  Pagination pagination = 1;
  PestDiseaseRecordFilter filter = 2;
}

message ListPestDiseaseRecordsResponse {
  repeated PestDiseaseRecord pest_disease_records = 1;
  int64 total = 2;
}

message GetPestDiseaseRecordsByPlantingCycleRequest {
  string planting_cycle_id = 1 [ (buf.validate.field).string.min_len = 1 ];
  Pagination pagination = 2;
  PestDiseaseRecordFilter filter = 3;
}

message GetPestDiseaseRecordsByPlantingCycleResponse {
  repeated PestDiseaseRecord pest_disease_records = 1;
  int64 total = 2;
}

message PestDiseaseRecordFilter {
  google.protobuf.Timestamp detection_date = 1;
  string detection_method = 2;
  google.protobuf.Timestamp treatment_date = 3;
  int32 treatment_duration_days = 4;
  google.protobuf.Timestamp follow_up_date = 5;
}

message Pagination {
  int32 page = 1 [ (buf.validate.field).int32.gte = 1 ];
  int32 page_size = 2 [
    (buf.validate.field).int32.gte = 1,
    (buf.validate.field).int32.lte = 1000
  ];
  string sort_by = 3 [ (buf.validate.field).string.max_len = 50 ];
  string sort_order = 4 [
    (buf.validate.field).string.in = "asc",
    (buf.validate.field).string.in = "desc"
  ];
}
