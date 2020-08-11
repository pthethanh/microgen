package templates

const (
	// UserProto service proto.
	UserProto = `syntax = "proto3";

package user;
option go_package = "{{.module_name}}/pkg/api/user;user";

import "google/api/annotations.proto";
import "google/protobuf/field_mask.proto";

service Users {
	// RegisterUser create a new user.
	rpc RegisterUser(RegisterUserRequest) returns (RegisterUserResponse) {
		option (google.api.http) = {
			post: "/api/v1/users/registrations"
			body: "*"
		};
	}

	// VerifyEmail verify email after registered.
	rpc VerifyEmail(VerifyEmailRequest) returns (VerifyEmailResponse){
		option (google.api.http) = {
			post: "/api/v1/users/emails/verification"
			body: "*"
		};
	}

	// Login authenticate the given user.
	rpc Login(LoginRequest) returns (LoginResponse) {
		option (google.api.http) = {
			post: "/api/v1/login"
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

	// UpdateUserStatus update status of the user.
	rpc UpdateUserStatus(UpdateUserStatusRequest) returns (UpdateUserStatusResponse) {
		option (google.api.http) = {
			put: "/api/v1/users/{user_id}/attributes/status"
			body: "*"
		};
	}

	// ChangePassword change user password.
	rpc ChangePassword(ChangePasswordRequest) returns (ChangePasswordResponse) {
		option (google.api.http) = {
			put: "/api/v1/users/{user_id}/password"
			body: "*"
		};
	}

	// ForgotPassword trigger forgot password flow.
	rpc ForgotPassword(ForgotPasswordRequest) returns (ForgotPasswordResponse) {
		option (google.api.http) = {
			post: "/api/v1/users/password/forgot"
			body: "*"
		};
	}

	// ResetPassword reset user password after the ForgotPassword flow is triggered.
	rpc ResetPassword(ResetPasswordRequest) returns (ResetPasswordResponse) {
		option (google.api.http) = {
			post: "/api/v1/users/password/reset"
			body: "*"
		};
	}

	// DeleteUser delete the user.
	rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse) {
		option (google.api.http) = {
			delete: "/api/v1/users/{user_id}"
		};
	}
	
	// ListUseers list all users.
	rpc FindUsers(FindUsersRequest) returns (FindUsersResponse) {
		option (google.api.http) = {
			get: "/api/v1/users"
		};
	}
}

message RegisterUserRequest {
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

message RegisterUserResponse {
	string user_id = 1;
}

message VerifyEmailRequest {
	string token = 1;
}

message VerifyEmailResponse {
	string user_id = 2;
}

message LoginRequest {
	string password = 1;
	oneof auth_id {
		string email = 2;
		string user_id = 3;
	}
}

message LoginResponse {
	string user_id = 1;
	string first_name = 2;
	string last_name = 3;
	string avatar = 4;
	map<string,string> attributes = 5;
}

message UpdateUserRequest {
	string user_id = 1;
	string first_name = 2;
	string last_name = 3;
	Gender gender = 4;
	string avatar = 5;
	string status = 6;
	map<string,string> attributes = 7;
	google.protobuf.FieldMask update_mask = 8;
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
	map<string,string> attributes = 6;
	repeated string sort_by = 7;
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
	map<string,string> attributes = 10;
}

message ChangePasswordRequest {
	string user_id = 1;
	string old_password = 2;
	string new_password = 3;
}

message ChangePasswordResponse {
	string user_id = 1;
}

message DeleteUserRequest {
	string user_id = 1;
}

message DeleteUserResponse {}

message UpdateUserStatusRequest {
	string user_id = 1;
	Status status = 2;
}

message UpdateUserStatusResponse {}

enum Status {
	STATUS_UNSPECIFIED = 0;
	STATUS_INACTIVE = 1;
	STATUS_ACTIVE = 2;
	STATUS_LOCKED = 3;
	STATUS_DELETED = 4;
}

message ForgotPasswordRequest {
	string email = 1;
}

message ForgotPasswordResponse {
	string token = 1;
}

message ResetPasswordRequest {
	string token = 1;
	string password = 2;
}

message ResetPasswordResponse {}
	`
)
