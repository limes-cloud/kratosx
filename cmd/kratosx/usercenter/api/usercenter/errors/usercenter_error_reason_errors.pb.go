// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errors

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsSystemError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SystemError.String() && e.Code == 400
}

func SystemError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_SystemError.String(), "系统异常")
	case 1:
		return errors.New(400, ErrorReason_SystemError.String(), "系统异常:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_SystemError.String(), "系统异常:"+msg)
	}
}

func IsParamsError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ParamsError.String() && e.Code == 400
}

func ParamsError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_ParamsError.String(), "参数错误")
	case 1:
		return errors.New(400, ErrorReason_ParamsError.String(), "参数错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_ParamsError.String(), "参数错误:"+msg)
	}
}

func IsDatabaseError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DatabaseError.String() && e.Code == 400
}

func DatabaseError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_DatabaseError.String(), "数据库错误")
	case 1:
		return errors.New(400, ErrorReason_DatabaseError.String(), "数据库错误:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_DatabaseError.String(), "数据库错误:"+msg)
	}
}

func IsTransformError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TransformError.String() && e.Code == 400
}

func TransformError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_TransformError.String(), "数据转换失败")
	case 1:
		return errors.New(400, ErrorReason_TransformError.String(), "数据转换失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_TransformError.String(), "数据转换失败:"+msg)
	}
}

func IsGetError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GetError.String() && e.Code == 400
}

func GetError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_GetError.String(), "获取数据失败")
	case 1:
		return errors.New(400, ErrorReason_GetError.String(), "获取数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_GetError.String(), "获取数据失败:"+msg)
	}
}

func IsListError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ListError.String() && e.Code == 400
}

func ListError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_ListError.String(), "获取列表数据失败")
	case 1:
		return errors.New(400, ErrorReason_ListError.String(), "获取列表数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_ListError.String(), "获取列表数据失败:"+msg)
	}
}

func IsCreateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CreateError.String() && e.Code == 400
}

func CreateError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_CreateError.String(), "创建数据失败")
	case 1:
		return errors.New(400, ErrorReason_CreateError.String(), "创建数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_CreateError.String(), "创建数据失败:"+msg)
	}
}

func IsImportError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ImportError.String() && e.Code == 400
}

func ImportError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_ImportError.String(), "导入数据失败")
	case 1:
		return errors.New(400, ErrorReason_ImportError.String(), "导入数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_ImportError.String(), "导入数据失败:"+msg)
	}
}

func IsExportError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ExportError.String() && e.Code == 400
}

func ExportError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_ExportError.String(), "导出数据失败")
	case 1:
		return errors.New(400, ErrorReason_ExportError.String(), "导出数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_ExportError.String(), "导出数据失败:"+msg)
	}
}

func IsUpdateError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UpdateError.String() && e.Code == 400
}

func UpdateError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_UpdateError.String(), "更新数据失败")
	case 1:
		return errors.New(400, ErrorReason_UpdateError.String(), "更新数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_UpdateError.String(), "更新数据失败:"+msg)
	}
}

func IsDeleteError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DeleteError.String() && e.Code == 400
}

func DeleteError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_DeleteError.String(), "删除数据失败")
	case 1:
		return errors.New(400, ErrorReason_DeleteError.String(), "删除数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_DeleteError.String(), "删除数据失败:"+msg)
	}
}

func IsGetTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GetTrashError.String() && e.Code == 400
}

func GetTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_GetTrashError.String(), "获取回收站数据失败")
	case 1:
		return errors.New(400, ErrorReason_GetTrashError.String(), "获取回收站数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_GetTrashError.String(), "获取回收站数据失败:"+msg)
	}
}

func IsListTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ListTrashError.String() && e.Code == 400
}

func ListTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_ListTrashError.String(), "获取回收站列表数据失败")
	case 1:
		return errors.New(400, ErrorReason_ListTrashError.String(), "获取回收站列表数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_ListTrashError.String(), "获取回收站列表数据失败:"+msg)
	}
}

func IsDeleteTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DeleteTrashError.String() && e.Code == 400
}

func DeleteTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_DeleteTrashError.String(), "删除回收站数据失败")
	case 1:
		return errors.New(400, ErrorReason_DeleteTrashError.String(), "删除回收站数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_DeleteTrashError.String(), "删除回收站数据失败:"+msg)
	}
}

func IsRevertTrashError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RevertTrashError.String() && e.Code == 400
}

func RevertTrashError(args ...any) *errors.Error {
	switch len(args) {
	case 0:
		return errors.New(400, ErrorReason_RevertTrashError.String(), "还原回收站数据失败")
	case 1:
		return errors.New(400, ErrorReason_RevertTrashError.String(), "还原回收站数据失败:"+fmt.Sprint(args[0]))
	default:
		msg := fmt.Sprintf(fmt.Sprint(args[0]), args[1:]...)
		return errors.New(400, ErrorReason_RevertTrashError.String(), "还原回收站数据失败:"+msg)
	}
}