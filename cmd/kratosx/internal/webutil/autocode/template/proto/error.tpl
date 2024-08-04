syntax = "proto3";

package errors;

import "errors/errors.proto";
option go_package = "./;errors";


enum ErrorReason {
  // 设置缺省错误码
  option (errors.default_code) = 400;

  SystemError = 0[(errors.message)="系统异常"];
  ParamsError = 0[(errors.message)="参数错误"];
  DatabaseError = 0[(errors.message)="数据库错误"];
  TransformError = 1[(errors.message)="数据转换失败"];
  GetError = 2[(errors.message)="获取数据失败"];
  ListError = 3[(errors.message)="获取列表数据失败"];
  CreateError = 4[(errors.message)="创建数据失败"];
  ImportError = 5[(errors.message)="导入数据失败"];
  ExportError = 6[(errors.message)="导出数据失败"];
  UpdateError = 7[(errors.message)="更新数据失败"];
  DeleteError = 8[(errors.message)="删除数据失败"];
  GetTrashError = 9[(errors.message)="获取回收站数据失败"];
  ListTrashError = 10[(errors.message)="获取回收站列表数据失败"];
  DeleteTrashError = 10[(errors.message)="删除回收站数据失败"];
  RevertTrashError = 11[(errors.message)="还原回收站数据失败"];
}
