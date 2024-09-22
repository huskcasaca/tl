package tl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

func Marshal(v any) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := NewEncoder(buf)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// MarshalState provides set of different methods to marshal binary message into
// specific struct. It's working absolutely like fmt.State, you just need to
// write data to this abstraction.
type MarshalState interface {
	// Writer will throw specific error if object will write bytes the size of which is not divided by the length of the word
	io.Writer

	PutBool(bool) error
	PutInt32(int32) error
	PutInt64(int64) error
	PutCrc32(crc32) error
	PutBytes([]byte) error
}

type RealEncoder interface {
	Encode(v any) error
}

// Encoder is a type, which allows you to decode serialized message.
type Encoder struct {
	Writer io.Writer

	ByteOrder binary.ByteOrder
}

func NewEncoder(w io.Writer) RealEncoder {
	return &Encoder{Writer: w, ByteOrder: binary.LittleEndian}
}

func (e *Encoder) Encode(value any) error {
	if value == nil {
		return ErrUnexpectedNil
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("value is not pointer as expected. got %v", v.Type())
	}

	return e.encodeValue(v)
}

// writeErr works like write (also throwing panic), but without count of
// written bytes.
func (e *Encoder) writeErr(b []byte) error {
	if n, err := e.write(b); err != nil {
		return err
	} else if n != len(b) {
		return &ErrPartialWrite{Has: n, Want: len(b)}
	}

	return nil
}

func (e *Encoder) Write(b []byte) (int, error) {
	if len(b)%WordLen != 0 {
		return 0, errors.New("value can't be divided by word length")
	}

	return e.write(b)
}

// write is private, cause this function might panic.
func (e *Encoder) write(b []byte) (int, error) {
	if len(b)%WordLen != 0 { //revive:disable-line:add-constant // makes no sense
		// it's panic, because it's internal method, and we must not write in
		// any case data, which is not divided by word length
		panic("raw bytes does not divide by word size of protocol")
	}

	if len(b) == 0 {
		return 0, nil
	}

	return e.Writer.Write(b) //nolint:wrapcheck // write() is a wrapper
}

//nolint:cyclop // it contains only assertion and switch statement
//revive:disable:function-length // same: can't make better
func (e *Encoder) encodeValue(value reflect.Value) error {
	if maybeNil(value) {
		return ErrUnexpectedNil
	}

	if marshaler, ok := value.Interface().(Marshaler); ok {
		//nolint:wrapcheck // object implements Marshaler must throw unwrapped
		//                    error
		return marshaler.MarshalTL(e)
	}

	switch k := value.Type().Kind(); k { //nolint:exhaustive // has default case
	case reflect.Uint32:
		return e.PutUint32(uint32(value.Uint()))

	case reflect.Int32:
		return e.PutUint32(uint32(value.Int()))

	case reflect.Uint64:
		return e.PutUint64(value.Uint())

	case reflect.Int64:
		return e.PutInt64(value.Int())

	case reflect.Float64:
		return e.PutFloat64(value.Float())

	case reflect.Bool:
		return e.PutBool(value.Bool())

	case reflect.String:
		return e.PutString(value.String())

	case reflect.Struct:
		return e.encodeStruct(value, false)

	case reflect.Ptr, reflect.Interface:
		return e.encodeValue(value.Elem())

	// case reflect.Map:
	//	return e.encodeMap(value)

	case reflect.Slice:
		return e.encodeVector(value)
	case reflect.Array:
		if value.Type().Elem() == byteTyp { // [N]byte
			return e.encodeRaw(value)
		}

		return e.encodeVector(value)

	default:
		return ErrUnsupportedType{Type: value.Type()}
	}
}

//revive:enable

// v must be pointer to struct.
func (e *Encoder) encodeStruct(v reflect.Value, ignoreCRC bool) error {
	o, ok := v.Interface().(Any)
	if !ok {
		// Trying to look implementation by pointer
		//
		// Since just struct might be non addressable, we are creating new
		// instance, and setting it to provided value.
		vcopy := reflect.New(v.Type()).Elem()
		vcopy.Set(v)

		var ok bool
		if o, ok = vcopy.Addr().Interface().(Any); !ok {
			return errors.New(v.Type().String() + " doesn't implement tl.Any interface")
		}
	}

	typ := v.Type()

	properties, bitflags, err := parseStructTags(typ)
	if err != nil {
		return fmt.Errorf("parsing struct flags: %w", err)
	}

	optFlags := make(map[int]crc32)
	for i, target := range bitflags {
		// even if we don't have any non null optional values, we still need to
		// initialize bitflags
		if _, ok := optFlags[target.FieldIndex]; !ok {
			optFlags[target.FieldIndex] = 0
		}
		if isFieldContainsData(v.Field(i)) {
			optFlags[target.FieldIndex] |= 1 << target.BitIndex
		}
	}

	// what we checked and what we know about value:
	// 1) it's not Marshaler (marshaler method if exist used already in c.encodeValue())
	// 2) implements tl.Any
	// 3) it's addressable
	// 4) definitely struct (we don't call encodeStruct() but in c.encodeValue())
	// 5) not nil (structs can't be nil, only pointers and interfaces)
	if !ignoreCRC {
		if err := e.PutCrc32(o.CRC()); err != nil {
			return err
		}
	}

	for i := 0; i < v.NumField(); i++ {
		// putting bitflags, if this field is bitflag
		if flags, ok := optFlags[i]; ok {
			if err := e.PutUint32(flags); err != nil {
				return err
			}

			continue
		}

		tag := properties[i]
		_, fieldOptional := bitflags[i]

		// if ignore or field is unexported, then go on
		if tag.Ignore() ||
			!v.Field(i).CanSet() ||
			(fieldOptional && (!isFieldContainsData(v.Field(i)) || tag.isImplicit())) {
			continue
		}

		err := e.encodeValue(v.Field(i))
		if err != nil {
			return fmt.Errorf("encoding %v: %w", v.Type().Field(i).Name, err)
		}
	}

	return nil
}

/*
func (e *Encoder) encodeMap(m reflect.Value) error {
	if m.Type().Key().Kind() != reflect.String {
		return errors.New("map keys are not string")
	}

	crc, err := getCRCFromMap(m)
	if err != nil {
		return err
	}

	definition, ok := e.registry.Tags(crc)
	if !ok {
		//nolint:goerr113 // it's an internal error
		return fmt.Errorf("crc code 0x%08x is not found in registry", crc)
	}

	// TODO: need to cache encoded non empty objects in slice, then write
	//       everything after we will be sure that cached bitflag will not be
	//       changed (means that we checked already most right optional field
	//       for this bitflag)
	bitflags, err := definition.collectBitflags(m)
	if err != nil {
		return err
	}

	if err := e.putCRC(crc); err != nil {
		return err
	}

	for i, field := range definition.fields {
		// putting bitflags, if this field is bitflag
		if flags, ok := bitflags[uint8(i)]; ok {
			if err := e.PutUint32(flags); err != nil {
				return err
			}

			continue
		}

		val := m.MapIndex(reflect.ValueOf(field.name))
		if !val.IsValid() || field.noEncode {
			continue
		}

		if err := e.encodeValue(val); err != nil {
			return errors.Wrapf(err, "encoding %q", field.name)
		}
	}

	return nil
}
*/

func (e *Encoder) encodeRaw(v reflect.Value) error {
	if v.Kind() != reflect.Array {
		panic("raw must be array")
	} else if v.Type().Elem() != byteTyp {
		panic("raw must be array of bytes")
	} else if n := v.Len(); n%WordLen != 0 {
		// special case: this means that we want to take exact N of bytes and pop it from reader
		// n%WordLen == 0, cause we can't read less or more than word
		return fmt.Errorf("array of bytes must be divided by %v, got %v", WordLen, n)
	}

	_, err := e.Write(v.Slice(0, v.Len()).Interface().([]byte))

	return err
}

func (e *Encoder) encodeVector(slice reflect.Value) error {
	if b, ok := slice.Interface().([]byte); ok {
		return e.PutBytes(b)
	}

	if err := e.PutCrc32(crcVector); err != nil {
		return err
	}
	if err := e.PutUint32(uint32(slice.Len())); err != nil {
		return err
	}

	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		err := e.encodeValue(item)
		if err != nil {
			return fmt.Errorf("[%v]: %w", i, err)
		}
	}

	return nil
}

func (e *Encoder) PutUint32(v uint32) error   { return e.writeErr(u32b(e.ByteOrder, v)) }
func (e *Encoder) PutUint64(v uint64) error   { return e.writeErr(u64b(e.ByteOrder, v)) }
func (e *Encoder) PutInt64(v int64) error     { return e.writeErr(u64b(e.ByteOrder, uint64(v))) }
func (e *Encoder) PutFloat64(v float64) error { return e.writeErr(f64b(e.ByteOrder, v)) }
func (e *Encoder) PutCrc32(v uint32) error    { return e.PutUint32(v) } // for selfdoc code
func (e *Encoder) PutInt32(v int32) error     { return e.PutUint32(uint32(v)) }
func (e *Encoder) PutBool(v bool) error       { return e.PutUint32(boolToCRC(v)) }
func (e *Encoder) PutString(v string) error   { return e.PutBytes([]byte(v)) }

func (e *Encoder) PutBytes(msg []byte) error {
	// 3 left bytes of word, which is 4 bytes
	const maxLen = 1 << ((WordLen - 1) * bitsInByte)
	if len(msg) > maxLen {
		//nolint:goerr113 // it's an internal error
		return fmt.Errorf("message entity too large: expect less than %v, got %v", maxLen, len(msg))
	}

	var lenBytes []byte

	// how does it work:
	// any object can be putted to byte set ONLY with length, without modula
	// after dividing to word length. e.g. bytes 'Hi!' can be written as:
	//            | 0x03 0x48 0x6A 0x21 |
	// Divides by 32 bits? Yes, so it's good.
	//
	// BUT! bytes 'Hello!' MUST be written as
	//            | 0x06 0x48 0x65 0x6C | 0x6C 0x6F 0x21 0x00 |
	// See? We added extra empty byte to pad message to length of word. That is
	// the most important part of putting bytes to buffer.
	//
	// So we must create a buffer with length mod to 32 == 0. To not add
	// extra bytes manually. They could be random, but who needs that, right?
	if len(msg) < fuckingMagicNumber {
		lenBytes = []byte{byte(len(msg))}
	} else {
		lenBytes = append([]byte{fuckingMagicNumber}, littleUint24Bytes(len(msg))...)
	}

	return e.writeErr(appendMany(
		lenBytes,
		msg,
		make([]byte, pad(len(lenBytes)+len(msg), WordLen)),
	))
}

// m only map.
func getCRCFromMap(m reflect.Value) (uint32, error) {
	crcVal := m.MapIndex(reflect.ValueOf(MapCrcKey))
	if !crcVal.IsValid() {
		return 0, errors.New("key " + MapCrcKey + " not exist in map")
	}
	if !crcVal.Type().ConvertibleTo(uint32Typ) {
		return 0, errors.New(MapCrcKey + " is not convertible to uint32")
	}

	return uint32(crcVal.Convert(uint32Typ).Uint()), nil
}
