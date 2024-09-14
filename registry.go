package tl

import (
	"fmt"
	"reflect"
)

// key is crc code, value is name of constructor.
type (
	enumNames = map[crc32]string
	typeName  = string
)

type Registry interface {
	ConstructObject(code crc32) (Any, bool)
}

var _defaultRegistry = val(NewRegistry())

func DefaultRegistry() *ObjectRegistry { return &_defaultRegistry }

func RegisterObjectDefault[T Any]() { RegisterObject[T](&_defaultRegistry) }
func RegisterCustomDefault[T Any](constructor func(uint32) T, crcs ...uint32) {
	RegisterCustom[T](&_defaultRegistry, constructor, crcs...)
}

// ObjectRegistry is a type, which handles code generated schema, and could be
// useful for spawning TL objects. Unlike RawSchemaRegistry, it can work only
// with predefined go types.
//
// If you are not able to use codegen, use RawSchemaRegistry, it could be
// slower, but more flexible.
type ObjectRegistry struct {
	// in objects it's allowed to store ONLY structs and uint32.
	objects map[crc32]func(uint32) Any
}

var _ Registry = (*ObjectRegistry)(nil)

func NewRegistry() *ObjectRegistry {
	return &ObjectRegistry{
		objects: make(map[crc32]func(uint32) Any),
	}
}

// ConstructObject spawns new object by crc code from registry.
func (r *ObjectRegistry) ConstructObject(crc crc32) (Any, bool) {
	if obj, ok := r.objects[crc]; ok {
		return obj(crc), true
	}

	return nil, false
}

type field struct {
	fType       fieldType
	name        string
	flagTrigger uint8 // index of field in list of fields in orphanType
	bitTrigger  uint8
	noEncode    bool // don't encode value, if value is true. works ONLY for boolean types
	optional    bool
}

type structFields struct {
	// key is an index of optional field in list of fields in object
	// value is bit, which you need to trigger
	bitflags map[int]BitflagBit
	tags     []StructTag
}

func (s *structFields) isFieldOptional(fieldIndex int) bool {
	_, ok := s.bitflags[fieldIndex]

	return ok
}

type BitflagBit struct {
	FieldIndex int // which field to put the bit in, indicating that the field exists
	BitIndex   int // specifically which bit to put the flag in, indicating that everything is okay
}

func RegisterObject[T Any](r *ObjectRegistry) {
	r.registerObject(asObject(new[T]))
}

func RegisterCustom[T Any](r *ObjectRegistry, constructor func(uint32) T, crcs ...uint32) {
	if len(crcs) == 0 {
		panic("no enums provided")
	}
	r.registerCustom(crcs, func(u uint32) Any { return constructor(u) })
}

func (r *ObjectRegistry) registerObject(c func(uint32) Any) {
	typ := reflect.TypeOf(c(0))
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		panic("accepted only structs")
	}

	if r.objects == nil {
		r.objects = make(map[crc32]func(uint32) Any)
	}
	crc := c(0).CRC()
	if _, ok := r.objects[crc]; ok {
		panic(fmt.Sprintf("object with code %x already registered", crc))
	}

	r.objects[crc] = c
}

func (r *ObjectRegistry) registerCustom(valid []uint32, c func(uint32) Any) {
	if r.objects == nil {
		r.objects = make(map[crc32]func(uint32) Any)
	}
	for _, enum := range valid {
		r.objects[enum] = c
	}
}
