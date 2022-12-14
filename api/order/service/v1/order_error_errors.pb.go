// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsOrderNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == OrderServiceErrorReason_ORDER_NOT_FOUND.String() && e.Code == 404
}

func ErrorOrderNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, OrderServiceErrorReason_ORDER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == OrderServiceErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, OrderServiceErrorReason_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}
