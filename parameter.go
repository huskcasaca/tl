package tl

import (
	"fmt"
	"math"
	"strings"

	"github.com/quenbyako/ext/slices"
	"golang.org/x/exp/constraints"
)

type Params []Param

func (p Params) String() string {
	return strings.Join(slices.Remap(p, func(p Param) string { return p.String() }), " ")
}

type Param interface {
	_Parameter()

	GetName() string
	GetType() Type
	GetComment() string
	SetComment(comment string)
	fmt.Stringer
}

var (
	_ Param = BitflagParam{}
	_ Param = RequiredParam{}
	_ Param = OptionalParam{}
	_ Param = TriggerParam{}
	_ Param = GenericParam{}
)

type BitflagParam struct {
	Comment string
	Name    string
}

func (_ BitflagParam) _Parameter()               {}
func (t BitflagParam) GetName() string           { return t.Name }
func (t BitflagParam) GetType() Type             { return Type{} }
func (t BitflagParam) GetComment() string        { return t.Comment }
func (t BitflagParam) SetComment(comment string) { t.Comment = comment }
func (t BitflagParam) String() string            { return t.Name + ":#" }

type RequiredParam struct {
	Comment string
	Name    string
	Type    Type
}

func (_ RequiredParam) _Parameter()               {}
func (t RequiredParam) GetName() string           { return t.Name }
func (t RequiredParam) GetType() Type             { return t.Type }
func (t RequiredParam) GetComment() string        { return t.Comment }
func (t RequiredParam) SetComment(comment string) { t.Comment = comment }
func (t RequiredParam) String() string            { return t.Name + ":" + t.Type.String() }

type OptionalParam struct {
	Comment     string
	Name        string
	Type        Type
	FlagTrigger string
	BitTrigger  int
}

func (_ OptionalParam) _Parameter()               {}
func (t OptionalParam) GetName() string           { return t.Name }
func (t OptionalParam) GetType() Type             { return t.Type }
func (t OptionalParam) GetComment() string        { return t.Comment }
func (t OptionalParam) SetComment(comment string) { t.Comment = comment }
func (t OptionalParam) String() string {
	return fmt.Sprintf("%v:%v.%v?%v", t.Name, t.FlagTrigger, t.BitTrigger, t.Type.String())
}

type TriggerParam struct {
	Comment     string
	Name        string
	FlagTrigger string
	BitTrigger  int
}

func (_ TriggerParam) _Parameter()               {}
func (t TriggerParam) GetName() string           { return t.Name }
func (t TriggerParam) GetType() Type             { return Type{} }
func (t TriggerParam) GetComment() string        { return t.Comment }
func (t TriggerParam) SetComment(comment string) { t.Comment = comment }
func (t TriggerParam) String() string {
	return fmt.Sprintf("%v:%v.%v?true", t.Name, t.FlagTrigger, t.BitTrigger)
}

func slicesMaxFunc[S ~[]T, T any](s S, f func(T) int) int {
	if len(s) == 0 {
		return 0
	}

	max := math.MinInt
	for _, item := range s {
		if m := f(item); m > max {
			max = m
		}
	}

	return max
}

func slicesMinFunc[S ~[]T, T any](s S, f func(T) int) int {
	if len(s) == 0 {
		return 0
	}

	min := math.MaxInt
	for _, item := range s {
		if m := f(item); m < min {
			min = m
		}
	}

	return min
}

func slicesSumFunc[S ~[]T, T any, C constraints.Ordered](s S, f func(T) C) (sum C) {
	if len(s) == 0 {
		return sum
	}

	for _, item := range s {
		sum += f(item)
	}

	return sum
}

type GenericParam struct {
	Comment string
	Name    string
	Type    Type
}

func (_ GenericParam) _Parameter()               {}
func (t GenericParam) GetName() string           { return t.Name }
func (t GenericParam) GetType() Type             { return t.Type }
func (t GenericParam) GetComment() string        { return t.Comment }
func (t GenericParam) SetComment(comment string) { t.Comment = comment }
func (t GenericParam) String() string {
	return fmt.Sprintf("{%v:%v}", t.Name, t.Type.String())
}
