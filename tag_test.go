package tl_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	. "github.com/xelaj/tl"
	"testing"
)

// checking that serializing and deserializing again got same result.
func TestParseTag(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		tag       string
		fieldName string
		want      StructTag
		wantErr   assert.ErrorAssertionFunc
	}{{
		tag:       "",
		fieldName: "SomeField",
		want: StructTag{
			Name: "SomeField",
		},
	}, {
		tag:       "some_field",
		fieldName: "SomeField",
		want: StructTag{
			Name: "some_field",
		},
	}, {
		tag:       "some_field,",
		fieldName: "SomeField",
		want: StructTag{
			Name: "some_field",
		},
	}, {
		tag:       ",omitempty:bitflag:30",
		fieldName: "SomeField",
		want: StructTag{
			Name: "SomeField",
			BitFlags: &Bitflag{
				TargetField: "bitflag",
				BitPosition: 30,
			},
		},
	}, {
		tag:       ",omitempty:bitflag:30,implicit",
		fieldName: "SomeField",
		want: StructTag{
			Name: "SomeField",
			BitFlags: &Bitflag{
				TargetField: "bitflag",
				BitPosition: 30,
			},
			Type: ImplicitBoolType,
		},
	}, {
		tag:       ",omitempty:otherflag",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:otherflag:1000",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:otherflag:-1",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",implicit",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",bitflag",
		fieldName: "SomeField",
		want: StructTag{
			Name: "SomeField",
			Type: BitflagType,
		},
	}, {
		tag:       "some_field,abracadabre",
		fieldName: "SomeField",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:bitflags:0,implicit,bitflag",
		fieldName: "SomeField",
		wantErr:   assert.Error,
	}, {
		tag:       ",implicit",
		fieldName: "",
		wantErr:   assert.Error,
	}, {
		tag:       ",omitempty:global_bitflags:0,bitflag",
		fieldName: "subflags",
		want: StructTag{
			Name: "subflags",
			BitFlags: &Bitflag{
				TargetField: "global_bitflags",
				BitPosition: 0,
			},
			Type: BitflagType,
		},
	}} {
		tt := tt // for parallel tests
		tt.wantErr = noErrAsDefault(tt.wantErr)

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseTag(tt.tag, tt.fieldName)
			if !tt.wantErr(t, err) || err != nil {
				return
			}

			require.Equal(t, tt.want, got)
		})
	}
}

type ParseStructTagsTestCase[T Any] struct {
	name         string
	wantTags     []StructTag
	wantBitflags map[int]BitflagBit
	wantErr      assert.ErrorAssertionFunc
}

var _ TestCase = (*ParseStructTagsTestCase[Any])(nil)

func (tt ParseStructTagsTestCase[T]) Name() string { return tt.name }

func (tt ParseStructTagsTestCase[T]) Run(t *testing.T) {
	tt.wantErr = noErrAsDefault(tt.wantErr)
	t.Parallel()

	var obj T
	gotTags, gotFlags, err := ParseStructTags(obj)
	if !tt.wantErr(t, err) || err != nil {
		return
	}

	require.Equal(t, tt.wantTags, gotTags)
	require.Equal(t, tt.wantBitflags, gotFlags)
}

func TestParseStructTags(t *testing.T) {
	for _, tt := range []TestCase{
		ParseStructTagsTestCase[*Poll]{
			name: "Poll",
			wantTags: []StructTag{
				{Name: "ID"},
				{Name: "flag", Type: BitflagType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 0}, Name: "Closed", Type: ImplicitBoolType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 1}, Name: "PublicVoters", Type: ImplicitBoolType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 2}, Name: "MultipleChoice", Type: ImplicitBoolType},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 3}, Name: "Quiz", Type: ImplicitBoolType},
				{Name: "Question"},
				{Name: "Answers"},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 4}, Name: "ClosePeriod"},
				{BitFlags: &Bitflag{TargetField: "flag", BitPosition: 5}, Name: "CloseDate"},
			},
			wantBitflags: map[int]BitflagBit{
				2: {FieldIndex: 1, BitIndex: 0},
				3: {FieldIndex: 1, BitIndex: 1},
				4: {FieldIndex: 1, BitIndex: 2},
				5: {FieldIndex: 1, BitIndex: 3},
				8: {FieldIndex: 1, BitIndex: 4},
				9: {FieldIndex: 1, BitIndex: 5},
			},
		},
	} {
		t.Run(tt.Name(), tt.Run)
	}
}

// type Poll struct {
// 	ID int64
// 	//nolint:revive // tl works with unexported tags
// 	_              null `tl:"flag,bitflag"`
// 	Closed         bool `tl:",omitempty:flag:0,implicit"`
// 	PublicVoters   bool `tl:",omitempty:flag:1,implicit"`
// 	MultipleChoice bool `tl:",omitempty:flag:2,implicit"`
// 	Quiz           bool `tl:",omitempty:flag:3,implicit"`
// 	Question       string
// 	Answers        []*PollAnswer
// 	ClosePeriod    *int32 `tl:",omitempty:flag:4"`
// 	CloseDate      *int32 `tl:",omitempty:flag:5"`
// }
