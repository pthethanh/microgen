package templates

const (
	// UserProto service proto.
	UserProto = `syntax = "proto3";

package user;
option go_package = "{{.module_name}}/pkg/api/user;user";

import "google/api/annotations.proto";
import "google/protobuf/field_mask.proto";

service Users {
	// CreateUser create a new user.
	rpc CreateUser(CreateUserRequest) returns (CreateUserResponse) {
		option (google.api.http) = {
			post: "/api/v1/users"
			body: "*"
		};
	}

	// GetUser get user detail.
	rpc GetUser(GetUserRequest) returns (GetUserResponse) {
		option (google.api.http) = {
			get: "/api/v1/users/{user_id}"
		};
	}
	
	// UpdateUser update user info.
	rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse) {
		option (google.api.http) = {
			patch: "/api/v1/users/{user_id}"
			body: "*"
		};
	}

	// DeleteUser delete the user.
	rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse) {
		option (google.api.http) = {
			delete: "/api/v1/users/{user_id}"
		};
	}
	
	// FindUsers list all users.
	rpc FindUsers(FindUsersRequest) returns (FindUsersResponse) {
		option (google.api.http) = {
			get: "/api/v1/users"
		};
	}
}

message CreateUserRequest {
	string email = 1;
	string password = 2;
	string first_name = 3;
	string last_name = 4;
	Gender gender = 5;
}

enum Gender {
	GENDER_UNSPECIFIED = 0;
	GENDER_MALE = 1;
	GENDER_FEMALE = 2;
	GENDER_OTHER = 3;
}

message CreateUserResponse {
	string user_id = 1;
}

message UpdateUserRequest {
	string user_id = 1;
	string first_name = 2;
	string last_name = 3;
	Gender gender = 4;
	string avatar = 5;
	Status status = 6;
	google.protobuf.FieldMask update_mask = 7;
}

message UpdateUserResponse {
	string user_id = 1;
}

message FindUsersRequest {
	int32 offset = 1;
	int32 limit = 2;
	repeated string user_id = 3;
	repeated Gender gender = 4;
	repeated Status status = 5;
	repeated string sort_by = 6;
}

message FindUsersResponse {
	repeated User users = 1;
}

message GetUserRequest {
	string user_id = 1;
}

message GetUserResponse {
	User user = 1;
}

message User {
	string user_id = 1;
	string email = 2;
	Gender gender = 3;
	Status status = 4;
	string first_name = 5;
	string last_name = 6;
	string avatar = 7;
	bool email_verified = 9;
}

message DeleteUserRequest {
	string user_id = 1;
}

message DeleteUserResponse {}

enum Status {
	STATUS_UNSPECIFIED = 0;
	STATUS_INACTIVE = 1;
	STATUS_ACTIVE = 2;
	STATUS_LOCKED = 3;
	STATUS_DELETED = 4;
}
	`
)
