syntax = "proto3";

package {{.Package}};

option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";
option java_outer_classname = "{{.JavaClass}}";

import "validate/validate.proto";

{{.GetRequest}}

{{.GetReply}}

{{.ListRequest}}

message List{{.Object}}Reply {
{{.ListReply}}

  uint32 total = 1;
  repeated {{.Object}} list = 2;
}

{{.CreateRequest}}

message Create{{.Object}}Reply {
  uint32 id = 1;
}

message Import{{.Object}}Request {
  repeated Create{{.Object}}Request list = 1;
}

message Import{{.Object}}Reply {
    uint32 total = 1;
}

{{.ExportRequest}}

message Export{{.Object}}Reply {
  string src = 1;
}

{{.UpdateRequest}}

message Update{{.Object}}Reply {
}

message Update{{.Object}}StatusRequest {
  uint32 id = 1[(validate.rules).uint32 = {gt: 0}];
  bool status = 2;
}

message Update{{.Object}}StatusReply {
}

message Delete{{.Object}}Request {
  repeated uint32 ids = 1[(validate.rules).repeated = {min_items: 1, unique:true, max_items:50}];
}

message Delete{{.Object}}Reply {
  uint32 total = 1;
}

message GetTrash{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gt: 0}];
}

{{.GetTrashReply}}

{{.ListTrashRequest}}

message ListTrash{{.Object}}Reply {
{{.ListTrashReply}}

  uint32 total = 1;
  repeated {{.Object}} list = 2;
}

message DeleteTrash{{.Object}}Request {
  repeated uint32 ids = 1[(validate.rules).repeated = {min_items: 1, unique:true, max_items:50}];
}

message DeleteTrash{{.Object}}Reply {
  uint32 total = 1;
}

message RevertTrash{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gt: 0}];
}

message RevertTrash{{.Object}}Reply {
}