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

{{.ListReply}}

{{.CreateRequest}}

message Create{{.Object}}Reply {
  uint32 id = 1;
}

{{.UpdateRequest}}

message Update{{.Object}}Reply {
}

{{if .EnableBatchDelete}}
message Delete{{.Object}}Request {
  repeated uint32 ids = 1[(validate.rules).repeated = {min_items: 1, unique:true, max_items:50}];
}

message Delete{{.Object}}Reply {
  uint32 total = 1;
}
{{else }}
message Delete{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gte: 1}];
}

message Delete{{.Object}}Reply {
}
{{end}}

message GetTrash{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gte: 1}];
}

{{.GetTrashReply}}

{{.ListTrashRequest}}

{{.ListTrashReply}}

{{if .EnableBatchDelete}}
message DeleteTrash{{.Object}}Request {
  repeated uint32 ids = 1[(validate.rules).repeated = {min_items: 1, unique:true, max_items:50}];
}

message DeleteTrash{{.Object}}Reply {
  uint32 total = 1;
}
{{else }}
message DeleteTrash{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gte: 1}];
}

message DeleteTrash{{.Object}}Reply {
}
{{end}}

message RevertTrash{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gte: 1}];
}

message RevertTrash{{.Object}}Reply {
}