syntax = "proto3";

package {{.Package}};

option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";
option java_outer_classname = "{{.JavaClass}}";

import "google/protobuf/empty.proto";
import "google/api/annotations.proto";

service Service {
  // Get{{.Object}} 获取指定的{{.Title}}
  rpc Get{{.Object}} (Get{{.Object}}Request) returns (Get{{.Object}}Reply) {
    option (google.api.http) = {
      get: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}",
    };
  }

  // List{{.Object}} 获取{{.Title}}列表
  rpc List{{.Object}} (List{{.Object}}Request) returns (List{{.Object}}Reply) {
    option (google.api.http) = {
      get: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}",
    };
  }

  // Create{{.Object}} 创建{{.Title}}
  rpc Create{{.Object}} (Create{{.Object}}Request) returns (Create{{.Object}}Reply) {
    option (google.api.http) = {
      post: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}",
      body: "*"
    };
  }

  // Import{{.Object}} 导入{{.Title}}
  rpc Import{{.Object}} (Import{{.Object}}Request) returns (Import{{.Object}}Reply) {
    option (google.api.http) = {
      post: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}/import",
      body: "*"
    };
  }

  // Export{{.Object}} 导出{{.Title}}
  rpc Export{{.Object}} (Export{{.Object}}Request) returns (Export{{.Object}}Reply) {
    option (google.api.http) = {
      post: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}/export",
      body: "*"
    };
  }

  // Update{{.Object}} 更新{{.Title}}
  rpc Update{{.Object}} (Update{{.Object}}Request) returns (Update{{.Object}}Reply) {
    option (google.api.http) = {
      put: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}",
      body: "*"
    };
  }

  // Delete{{.Object}} 删除{{.Title}}
  rpc Delete{{.Object}} (Delete{{.Object}}Request) returns (Delete{{.Object}}Reply) {
    option (google.api.http) = {
      delete: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}",
    };
  }

  // BatchDelete{{.Object}} 批量删除{{.Title}}
  rpc BatchDelete{{.Object}} (BatchDelete{{.Object}}Request) returns (BatchDelete{{.Object}}Reply) {
    option (google.api.http) = {
      delete: "/{{.ServerLowerCase}}/{{.ModuleLowerCase}}/api/v1/{{.ObjectLowerCase}}",
    };
  }
 }