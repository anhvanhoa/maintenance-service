syntax = "proto3";

package maintenance_schedule.v1;

import "google/protobuf/timestamp.proto";
import "buf/validate/validate.proto";
import "common/v1/common.proto";

option go_package = "github.com/anhvanhoa/sf-proto/gen/maintenance_schedule/v1;proto_maintenance_schedule";

message MaintenanceSchedule {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
  string device_id = 2 [ (buf.validate.field).string.uuid = true ];
  string maintenance_type = 3 [
    (buf.validate.field).string.in = "cleaning",
    (buf.validate.field).string.in = "calibration",
    (buf.validate.field).string.in = "replacement",
    (buf.validate.field).string.in = "repair",
    (buf.validate.field).string.in = "inspection",
    (buf.validate.field).string.in = "software_update"
  ];
  string maintenance_category = 4 [
    (buf.validate.field).string.in = "preventive",
    (buf.validate.field).string.in = "corrective",
    (buf.validate.field).string.in = "emergency",
    (buf.validate.field).string.in = "routine"
  ];
  string priority = 5 [
    (buf.validate.field).string.in = "low",
    (buf.validate.field).string.in = "medium",
    (buf.validate.field).string.in = "high",
    (buf.validate.field).string.in = "critical"
  ];
  google.protobuf.Timestamp scheduled_date = 6;
  double estimated_duration_hours = 7
      [ (buf.validate.field).double = {gte : 0, lte : 168} ];
  google.protobuf.Timestamp completed_date = 8;
  double actual_duration_hours = 9
      [ (buf.validate.field).double = {gte : 0, lte : 168} ];
  string technician = 10 [ (buf.validate.field).string = {max_len : 100} ];
  string technician_contact = 11
      [ (buf.validate.field).string = {max_len : 100} ];
  double cost = 12 [ (buf.validate.field).double = {gte : 0} ];
  string parts_replaced = 13 [ (buf.validate.field).string = {max_len : 1000} ];
  string tools_required = 14 [ (buf.validate.field).string = {max_len : 1000} ];
  string safety_precautions = 15
      [ (buf.validate.field).string = {max_len : 2000} ];
  string pre_maintenance_readings = 16
      [ (buf.validate.field).string = {max_len : 2000} ];
  string post_maintenance_readings = 17
      [ (buf.validate.field).string = {max_len : 2000} ];
  string calibration_values = 18
      [ (buf.validate.field).string = {max_len : 2000} ];
  string test_results = 19 [ (buf.validate.field).string = {max_len : 2000} ];
  string status = 20 [
    (buf.validate.field).string.in = "scheduled",
    (buf.validate.field).string.in = "in_progress",
    (buf.validate.field).string.in = "completed",
    (buf.validate.field).string.in = "cancelled",
    (buf.validate.field).string.in = "postponed"
  ];
  int32 completion_rating = 21
      [ (buf.validate.field).int32 = {gte : 1, lte : 5} ];
  google.protobuf.Timestamp next_maintenance_date = 22;
  int32 maintenance_interval_days = 23
      [ (buf.validate.field).int32 = {gte : 1, lte : 3650} ];
  bool warranty_impact = 24;
  int32 downtime_minutes = 25
      [ (buf.validate.field).int32 = {gte : 0, lte : 10080} ];
  string notes = 26 [ (buf.validate.field).string = {max_len : 2000} ];
  string maintenance_log = 27
      [ (buf.validate.field).string = {max_len : 5000} ];
  repeated string before_images = 28;
  repeated string after_images = 29;
  string created_by = 30
      [ (buf.validate.field).string = {min_len : 1, max_len : 100} ];
  google.protobuf.Timestamp created_at = 31;
  google.protobuf.Timestamp updated_at = 32;
}

message MaintenanceScheduleFilter {
  string device_id = 1 [ (buf.validate.field).string = {max_len : 50} ];
  repeated string statuses = 2 [ (buf.validate.field).repeated = {
    items : {
      string : {
        in : [
          "scheduled",
          "in_progress",
          "completed",
          "cancelled",
          "postponed"
        ]
      }
    }
  } ];
  repeated string types = 3 [ (buf.validate.field).repeated = {
    items : {
      string : {
        in : [
          "cleaning",
          "calibration",
          "replacement",
          "repair",
          "inspection",
          "software_update"
        ]
      }
    }
  } ];
  repeated string categories = 4 [ (buf.validate.field).repeated = {
    items : {
      string : {in : [ "preventive", "corrective", "emergency", "routine" ]}
    }
  } ];

  repeated string priorities = 5 [ (buf.validate.field).repeated = {
    items : {string : {in : [ "low", "medium", "high", "critical" ]}}
  } ];

  google.protobuf.Timestamp from_date = 6;
  google.protobuf.Timestamp to_date = 7;
}

message CreateMaintenanceScheduleRequest {
  string device_id = 1 [ (buf.validate.field).string.uuid = true ];
  string maintenance_type = 2 [
    (buf.validate.field).string.in = "cleaning",
    (buf.validate.field).string.in = "calibration",
    (buf.validate.field).string.in = "replacement",
    (buf.validate.field).string.in = "repair",
    (buf.validate.field).string.in = "inspection",
    (buf.validate.field).string.in = "software_update"
  ];
  string maintenance_category = 3 [
    (buf.validate.field).string.in = "preventive",
    (buf.validate.field).string.in = "corrective",
    (buf.validate.field).string.in = "emergency",
    (buf.validate.field).string.in = "routine"
  ];
  string priority = 4 [
    (buf.validate.field).string.in = "low",
    (buf.validate.field).string.in = "medium",
    (buf.validate.field).string.in = "high",
    (buf.validate.field).string.in = "critical"
  ];
  google.protobuf.Timestamp scheduled_date = 5;
  double estimated_duration_hours = 6
      [ (buf.validate.field).double = {gte : 0, lte : 168} ];
  google.protobuf.Timestamp completed_date = 7;
  double actual_duration_hours = 8
      [ (buf.validate.field).double = {gte : 0, lte : 168} ];
  string technician = 9 [ (buf.validate.field).string = {max_len : 100} ];
  string technician_contact = 10
      [ (buf.validate.field).string = {max_len : 100} ];
  double cost = 11 [ (buf.validate.field).double = {gte : 0} ];
  string parts_replaced = 12 [ (buf.validate.field).string = {max_len : 1000} ];
  string tools_required = 13 [ (buf.validate.field).string = {max_len : 1000} ];
  string safety_precautions = 14
      [ (buf.validate.field).string = {max_len : 2000} ];
  string pre_maintenance_readings = 15
      [ (buf.validate.field).string = {max_len : 2000} ];
  string post_maintenance_readings = 16
      [ (buf.validate.field).string = {max_len : 2000} ];
  string calibration_values = 17
      [ (buf.validate.field).string = {max_len : 2000} ];
  string test_results = 18 [ (buf.validate.field).string = {max_len : 2000} ];
  string status = 19 [
    (buf.validate.field).string.in = "scheduled",
    (buf.validate.field).string.in = "in_progress",
    (buf.validate.field).string.in = "completed",
    (buf.validate.field).string.in = "cancelled",
    (buf.validate.field).string.in = "postponed"
  ];
  int32 completion_rating = 20
      [ (buf.validate.field).int32 = {gte : 1, lte : 5} ];
  google.protobuf.Timestamp next_maintenance_date = 21;
  int32 maintenance_interval_days = 22
      [ (buf.validate.field).int32 = {gte : 1, lte : 3650} ];
  bool warranty_impact = 23;
  int32 downtime_minutes = 24
      [ (buf.validate.field).int32 = {gte : 0, lte : 10080} ];
  string notes = 25 [ (buf.validate.field).string = {max_len : 2000} ];
  string maintenance_log = 26
      [ (buf.validate.field).string = {max_len : 5000} ];
  repeated string before_images = 27;
  repeated string after_images = 28;
  string created_by = 29
      [ (buf.validate.field).string = {min_len : 1, max_len : 100} ];
}

message CreateMaintenanceScheduleResponse {
  MaintenanceSchedule maintenance_schedule = 1;
}

message GetMaintenanceScheduleRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}

message GetMaintenanceScheduleResponse {
  MaintenanceSchedule maintenance_schedule = 1;
}

message UpdateMaintenanceScheduleRequest {
  string id = 1 [
    (buf.validate.field).string.uuid = true,
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string device_id = 2 [
    (buf.validate.field).string.uuid = true,
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string maintenance_type = 3 [
    (buf.validate.field).string.in = "cleaning",
    (buf.validate.field).string.in = "calibration",
    (buf.validate.field).string.in = "replacement",
    (buf.validate.field).string.in = "repair",
    (buf.validate.field).string.in = "inspection",
    (buf.validate.field).string.in = "software_update",
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string maintenance_category = 4 [
    (buf.validate.field).string.in = "preventive",
    (buf.validate.field).string.in = "corrective",
    (buf.validate.field).string.in = "emergency",
    (buf.validate.field).string.in = "routine",
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string priority = 5 [
    (buf.validate.field).string.in = "low",
    (buf.validate.field).string.in = "medium",
    (buf.validate.field).string.in = "high",
    (buf.validate.field).string.in = "critical",
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  google.protobuf.Timestamp scheduled_date = 6;
  double estimated_duration_hours = 7 [
    (buf.validate.field).double = {gte : 0, lte : 168},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  google.protobuf.Timestamp completed_date = 8;
  double actual_duration_hours = 9 [
    (buf.validate.field).double = {gte : 0, lte : 168},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string technician = 10 [
    (buf.validate.field).string = {max_len : 100},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string technician_contact = 11 [
    (buf.validate.field).string = {max_len : 100},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  double cost = 12 [
    (buf.validate.field).double = {gte : 0},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string parts_replaced = 13 [ (buf.validate.field).string = {max_len : 1000} ];
  string tools_required = 14 [ (buf.validate.field).string = {max_len : 1000} ];
  string safety_precautions = 15 [
    (buf.validate.field).string = {max_len : 2000},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string pre_maintenance_readings = 16 [
    (buf.validate.field).string = {max_len : 2000},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string post_maintenance_readings = 17 [
    (buf.validate.field).string = {max_len : 2000},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string calibration_values = 18 [
    (buf.validate.field).string = {max_len : 2000},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string test_results = 19 [
    (buf.validate.field).string = {max_len : 2000},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string status = 20 [
    (buf.validate.field).string.in = "scheduled",
    (buf.validate.field).string.in = "in_progress",
    (buf.validate.field).string.in = "completed",
    (buf.validate.field).string.in = "cancelled",
    (buf.validate.field).string.in = "postponed",
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  int32 completion_rating = 21 [
    (buf.validate.field).int32 = {gte : 1, lte : 5},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  google.protobuf.Timestamp next_maintenance_date = 22;
  int32 maintenance_interval_days = 23 [
    (buf.validate.field).int32 = {gte : 1, lte : 3650},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  bool warranty_impact = 24;
  int32 downtime_minutes = 25 [
    (buf.validate.field).int32 = {gte : 0, lte : 10080},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string notes = 26 [
    (buf.validate.field).string = {max_len : 2000},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  string maintenance_log = 27 [
    (buf.validate.field).string = {max_len : 5000},
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
  repeated string before_images = 28;
  repeated string after_images = 29;
  string created_by = 30 [
    (buf.validate.field).string.uuid = true,
    (buf.validate.field).ignore = IGNORE_ALWAYS
  ];
}

message UpdateMaintenanceScheduleResponse {
  MaintenanceSchedule maintenance_schedule = 1;
}

message DeleteMaintenanceScheduleRequest {
  string id = 1 [ (buf.validate.field).string.uuid = true ];
}

message DeleteMaintenanceScheduleResponse {
  string message = 1;
  bool success = 2;
}

message ListMaintenanceScheduleRequest {
  common.PaginationRequest pagination = 1;
  MaintenanceScheduleFilter filter = 2;
}

message ListMaintenanceScheduleResponse {
  repeated MaintenanceSchedule maintenance_schedules = 1;
  common.PaginationResponse pagination = 2;
}

service MaintenanceScheduleService {
  rpc CreateMaintenanceSchedule(CreateMaintenanceScheduleRequest)
      returns (CreateMaintenanceScheduleResponse);
  rpc GetMaintenanceSchedule(GetMaintenanceScheduleRequest)
      returns (GetMaintenanceScheduleResponse);
  rpc UpdateMaintenanceSchedule(UpdateMaintenanceScheduleRequest)
      returns (UpdateMaintenanceScheduleResponse);
  rpc DeleteMaintenanceSchedule(DeleteMaintenanceScheduleRequest)
      returns (DeleteMaintenanceScheduleResponse);
  rpc ListMaintenanceSchedule(ListMaintenanceScheduleRequest)
      returns (ListMaintenanceScheduleResponse);
}
