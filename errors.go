// Copyright (c) 2022 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package tl

import (
	"fmt"
	"reflect"
)

const (
	ErrUnexpectedNil = errorConst("unexpected nil value")
	ErrImplicitInt   = errorConst("value must be converted to int32, int64 or uint32 explicitly")

	ErrBitflagTooHigh   = errorConst("trigger bit is more than 32")
	ErrImplicitNoTarget = errorConst(implicitFlag + " defined without target field to trigger")
	ErrInvalidTagOption = errorConst("invalid option")
	ErrInvalidTagFormat = errorConst("invalid tag format")
)

type ErrRegisteredObjectNotFound struct {
	Data []byte
	Crc  uint32
}

func (e ErrRegisteredObjectNotFound) Error() string {
	return fmt.Sprintf("object with provided crc not registered: 0x%08x", e.Crc)
}

type ErrorPartialWrite struct {
	Has  int
	Want int
}

func (e ErrorPartialWrite) Error() string {
	return fmt.Sprintf("write failed: writed only %v bytes, expected %v", e.Has, e.Want)
}

type ErrUnsupportedInt struct {
	Kind reflect.Kind
}

func (e ErrUnsupportedInt) Error() string {
	return fmt.Sprintf("int32, int64 or uint32 allowed, %v is unsupported by protocol", e.Kind)
}

type ErrUnsupportedFloat struct {
	Kind reflect.Kind
}

func (e ErrUnsupportedFloat) Error() string {
	return fmt.Sprintf("only float64 is allowed, %v is unsupported by protocol", e.Kind)
}

type ErrUnsupportedType struct {
	Type reflect.Type
}

func (e ErrUnsupportedType) Error() string {
	return fmt.Sprintf("invalid or unknown type %v", e.Type)
}

type ErrObjectNotRegistered crc32

func (e ErrObjectNotRegistered) Error() string {
	return fmt.Sprintf("object with crc 0x%08x not registered", crc32(e))
}

type ErrInvalidCRC struct {
	Got  crc32
	Want crc32
}

func (e ErrInvalidCRC) Error() string {
	return fmt.Sprintf("invalid crc code: got 0x%08x; want 0x%08x", e.Got, e.Want)
}

type ErrInvalidBoolCRC crc32

func (e ErrInvalidBoolCRC) Error() string {
	return fmt.Sprintf(
		"want a 0x%08x (true) or 0x%08x (false); got 0x%08x",
		crcTrue,
		crcFalse,
		crc32(e),
	)
}

type errReadCRC struct {
	Wrapped error
}

func (e errReadCRC) Error() string { return "read crc: " + e.Wrapped.Error() }
func (e errReadCRC) Unwrap() error { return e.Wrapped }

type ErrPath struct {
	Err  error
	Path string
}

func (e ErrPath) Error() string { return e.Path + ": " + e.Err.Error() }
func (e ErrPath) Unwrap() error { return e.Err }

func wrapPath(err error, path string) error {
	if err == nil {
		return nil
	}

	if err, ok := as118[ErrPath](err); ok {
		err.Path += "." + path

		return err
	}

	return ErrPath{
		Path: path,
		Err:  err,
	}
}
