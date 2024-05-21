syntax = "proto3";

package errors;

import "errors/errors.proto";
option go_package = "./;errors";


enum ErrorReason {
  // 设置缺省错误码
  option (errors.default_code) = 500;

  NotFound = 0[(errors.message)="数据不存在"];
  Database = 1[(errors.message)="数据库错误"];
  Transform = 2[(errors.message)="数据转换失败"];
  List = 3[(errors.message)="获取列表数据失败"];
  Create = 4[(errors.message)="创建数据失败"];
  Import = 5[(errors.message)="导入数据失败"];
  Export = 6[(errors.message)="导出数据失败"];
  Update = 7[(errors.message)="更新数据失败"];
  Delete = 8[(errors.message)="删除数据失败"];
  BatchDelete = 9[(errors.message)="批量删除数据失败"];
}
