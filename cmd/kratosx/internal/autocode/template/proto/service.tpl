syntax = "proto3";

package {{.Package}};

option go_package = "{{.GoPackage}}";
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";
option java_outer_classname = "{{.JavaClass}}";

import "google/api/annotations.proto";

service Service {
  // Get{{.Object}} 获取指定的{{.Title}}
  rpc Get{{.Object}} (Get{{.Object}}Request) returns (Get{{.Object}}Reply) {
    option (google.api.http) = {
      get: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}",
    };
  }

  // List{{.Object}} 获取{{.Title}}列表
  rpc List{{.Object}} (List{{.Object}}Request) returns (List{{.Object}}Reply) {
    option (google.api.http) = {
      get: "/{{.ServerLowerCase}}/api/v1/{{.ObjectPluralizeLowerCase}}",
    };
  }

  // Create{{.Object}} 创建{{.Title}}
  rpc Create{{.Object}} (Create{{.Object}}Request) returns (Create{{.Object}}Reply) {
    option (google.api.http) = {
      post: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}",
      body: "*"
    };
  }

  // Import{{.Object}} 导入{{.Title}}
  rpc Import{{.Object}} (Import{{.Object}}Request) returns (Import{{.Object}}Reply) {
    option (google.api.http) = {
      post: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}/import",
      body: "*"
    };
  }

  // Export{{.Object}} 导出{{.Title}}
  rpc Export{{.Object}} (Export{{.Object}}Request) returns (Export{{.Object}}Reply) {
    option (google.api.http) = {
      post: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}/export",
      body: "*"
    };
  }

  // Update{{.Object}} 更新{{.Title}}
  rpc Update{{.Object}} (Update{{.Object}}Request) returns (Update{{.Object}}Reply) {
    option (google.api.http) = {
      put: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}",
      body: "*"
    };
  }

  // Update{{.Object}}Status 更新{{.Title}}状态
  rpc Update{{.Object}}Status (Update{{.Object}}StatusRequest) returns (Update{{.Object}}StatusReply) {
    option (google.api.http) = {
      put: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}/status",
      body: "*"
    };
  }

  // Delete{{.Object}} 删除{{.Title}}
  rpc Delete{{.Object}} (Delete{{.Object}}Request) returns (Delete{{.Object}}Reply) {
    option (google.api.http) = {
      delete: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}",
    };
  }


  // GetTrash{{.Object}} 查看指定{{.Title}}回收站数据
  rpc GetTrash{{.Object}} (GetTrash{{.Object}}Request) returns (GetTrash{{.Object}}Reply) {
    option (google.api.http) = {
      get: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}/trash",
    };
  }

  // ListTrash{{.Object}} 查看{{.Title}}列表回收站数据
  rpc ListTrash{{.Object}} (ListTrash{{.Object}}Request) returns (ListTrash{{.Object}}Reply) {
    option (google.api.http) = {
      get: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}/trashes",
    };
  }

  // DeleteTrash{{.Object}} 彻底删除{{.Title}}
  rpc DeleteTrash{{.Object}} (DeleteTrash{{.Object}}Request) returns (DeleteTrash{{.Object}}Reply) {
    option (google.api.http) = {
      delete: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}/trash",
    };
  }

  // RevertTrash{{.Object}} 还原{{.Title}}
  rpc RevertTrash{{.Object}} (RevertTrash{{.Object}}Request) returns (RevertTrash{{.Object}}Reply) {
    option (google.api.http) = {
      put: "/{{.ServerLowerCase}}/api/v1/{{.ObjectLowerCase}}/trash",
      body: "*"
    };
  }
}