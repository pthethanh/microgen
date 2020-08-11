package templates

const (
	// UserEventProto event proto.
	UserEventProto = `syntax = "proto3";

package user.event;
// Note: always define package for a correct package generation.
option go_package = "{{.module_name}}/pkg/api/user/event;event";

import "api/user/rpc.proto";
import "google/protobuf/field_mask.proto";


message UserCreated {
	user.User user = 1;
}

message UserUpdated {
	user.User user = 1;
	google.protobuf.FieldMask update_mask = 2;
}

message UserDeleted {
	string user_id = 1;
}
	`
)
