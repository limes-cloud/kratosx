syntax = "proto3";

package {{.Package}};

option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";
option java_outer_classname = "{{.JavaClass}}";

import "validate/validate.proto";

message Get{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gt: 0}];
}

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
    uint32 updateTotal = 1;
    uint32 createTotal = 2;
}

{{.ExportRequest}}

message Export{{.Object}}Reply {
  string src = 1;
}

{{.UpdateRequest}}

message Update{{.Object}}Reply {
}

message Delete{{.Object}}Request {
  uint32 id = 1[(validate.rules).uint32 = {gt: 0}];
}

message Delete{{.Object}}Reply {
}

message BatchDelete{{.Object}}Request {
  repeated uint32 ids = 1[(validate.rules).repeated = {min_items: 0, unique:true, max_items:50}];
}

message BatchDelete{{.Object}}Reply {
  uint32 total = 1;
}