// Copyright (c) 2022-2024 Xelaj Software
//
// This file is a part of tl package.
// See https://github.com/xelaj/tl/blob/master/LICENSE_README.md for details.

package schema

import (
	"fmt"
	"math"
	"strings"

	"github.com/quenbyako/ext/slices"
	"golang.org/x/exp/constraints"
)

type TLParams []TLParam

func (p TLParams) String() string {
	return strings.Join(slices.Remap(p, func(p TLParam) string { return p.String() }), " ")
}

func (p TLParams) Comments() []string {
	params := slices.Filter(p, func(p TLParam) bool { return p.GetComment() != "" })
	paramMaxLen := slicesMaxFunc(params, func(p TLParam) int { return len([]rune(p.GetName())) })

	return slices.Remap(params, func(p TLParam) string {
		return fmt.Sprintf("// @param %-*v %v", paramMaxLen, p.GetName(), p.GetComment())
	})
}

type TLParam interface {
	_Parameter()

	GetName() string
	GetComment() string
	fmt.Stringer
}

var (
	_ TLParam = TLBitflagParam{}
	_ TLParam = TLRequiredParam{}
	_ TLParam = TLOptionalParam{}
	_ TLParam = TLTriggerParam{}
	_ TLParam = TLPolyParam{}
)

type TLBitflagParam struct {
	Comment string
	Name    string
}

func (_ TLBitflagParam) _Parameter()        {}
func (t TLBitflagParam) GetName() string    { return t.Name }
func (t TLBitflagParam) GetComment() string { return t.Comment }
func (t TLBitflagParam) String() string     { return t.Name + ":#" }

type TLRequiredParam struct {
	Comment string
	Name    string
	Type    TLType
}

func (_ TLRequiredParam) _Parameter()        {}
func (t TLRequiredParam) GetName() string    { return t.Name }
func (t TLRequiredParam) GetComment() string { return t.Comment }
func (t TLRequiredParam) String() string     { return t.Name + ":" + t.Type.String() }

type TLOptionalParam struct {
	Comment     string
	Name        string
	Type        TLType
	FlagTrigger string
	BitTrigger  int
}

func (_ TLOptionalParam) _Parameter()        {}
func (t TLOptionalParam) GetName() string    { return t.Name }
func (t TLOptionalParam) GetComment() string { return t.Comment }
func (t TLOptionalParam) String() string {
	return fmt.Sprintf("%v:%v.%v?%v", t.Name, t.FlagTrigger, t.BitTrigger, t.Type.String())
}

type TLTriggerParam struct {
	Comment     string
	Name        string
	FlagTrigger string
	BitTrigger  int
}

func (_ TLTriggerParam) _Parameter()        {}
func (t TLTriggerParam) GetName() string    { return t.Name }
func (t TLTriggerParam) GetComment() string { return t.Comment }
func (t TLTriggerParam) String() string {
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

type TLPolyParam struct {
	Comment string
	Name    string
	Type    TLType
}

func (_ TLPolyParam) _Parameter()        {}
func (t TLPolyParam) GetName() string    { return t.Name }
func (t TLPolyParam) GetComment() string { return t.Comment }
func (t TLPolyParam) String() string {
	return fmt.Sprintf("%v:%v", t.Name, t.Type.String())
}
