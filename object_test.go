//revive:disable:add-constant       fixtures must have constants
//revive:disable:max-public-structs for test purposes

package tl_test

import (
	. "github.com/xelaj/tl"
)

type GenericRequest[T any] struct {
	MsgID          int64
	_              null `tl:"flag,bitflag"`
	IsSpecificType bool `tl:",omitempty:flag:0,implicit"`
	Msg            T    `tl:",omitempty:flag:1"`
}

func (*GenericRequest[T]) CRC() uint32 { return uint32(0x77553311) }

type MultipleChats struct {
	Chats []Chat
}

func (*MultipleChats) CRC() uint32 { return uint32(0xff1144cc) }

type Chat struct {
	//nolint:revive // tl works with unexported tags
	_                 null `tl:"flag,bitflag"`
	Creator           bool `tl:",omitempty:flag:0,implicit"`
	Kicked            bool `tl:",omitempty:flag:1,implicit"`
	Left              bool `tl:",omitempty:flag:2,implicit"`
	Deactivated       bool `tl:",omitempty:flag:5,implicit"`
	ID                int32
	Title             string
	Photo             string
	ParticipantsCount int32
	Date              int32
	Version           int32
	AdminRights       *Rights `tl:",omitempty:flag:14"`
	BannedRights      *Rights `tl:",omitempty:flag:18"`
	_                 null    `tl:"flag2,bitflag"`
	CustomField       bool    `tl:",omitempty:flag2:0,implicit"`
}

func (*Chat) CRC() uint32 { return uint32(0x3bda1bde) }

type AuthSentCode struct {
	//nolint:revive // tl works with unexported tags
	_             null `tl:"flag,bitflag"`
	Type          AuthSentCodeType
	PhoneCodeHash string
	NextType      AuthCodeType `tl:",omitempty:flag:1"`
	Timeout       *int32       `tl:",omitempty:flag:2"`
}

func (*AuthSentCode) CRC() uint32 { return uint32(0x5e002502) }

type SomeNullStruct null

func (*SomeNullStruct) CRC() uint32 { return uint32(0xc4f9186b) }

type AuthSentCodeType interface {
	Any
	_AuthSentCodeType()
}

type AuthSentCodeTypeApp struct {
	Length int32
}

var _ AuthSentCodeType = (*AuthSentCodeTypeApp)(nil)

func (*AuthSentCodeTypeApp) CRC() uint32 { return uint32(0x3dbb5986) }

func (*AuthSentCodeTypeApp) _AuthSentCodeType() {}

type Rights struct {
	//nolint:revive // tl works with unexported tags
	_              null `tl:"flag,bitflag"`
	DeleteMessages bool `tl:",omitempty:flag:3,implicit"`
	BanUsers       bool `tl:",omitempty:flag:4,implicit"`
}

func (*Rights) CRC() uint32 { return uint32(0x5fb224d5) }

type AuthCodeType interface {
	Any
	_AuthCodeType()
}

var _ AuthCodeType = (*AuthCodeTypeSms)(nil)
var _ AuthCodeType = (*AuthCodeTypeCall)(nil)
var _ AuthCodeType = (*AuthCodeTypeFlashCall)(nil)

type AuthCodeTypeSms struct {
}

func (*AuthCodeTypeSms) CRC() uint32 { return uint32(1923290508) }

func (*AuthCodeTypeSms) _AuthCodeType() {}

type AuthCodeTypeCall struct {
}

func (*AuthCodeTypeCall) CRC() uint32 { return uint32(1948046307) }

func (*AuthCodeTypeCall) _AuthCodeType() {}

type AuthCodeTypeFlashCall struct {
}

func (*AuthCodeTypeFlashCall) CRC() uint32 { return uint32(577556219) }

func (*AuthCodeTypeFlashCall) _AuthCodeType() {}

type PollResults struct {
	//nolint:revive // tl works with unexported tags
	_                null                `tl:"flag,bitflag"`
	Min              bool                `tl:",omitempty:flag:0,implicit"`
	Results          []*PollAnswerVoters `tl:",omitempty:flag:1"`
	TotalVoters      *int32              `tl:",omitempty:flag:2"`
	RecentVoters     []int32             `tl:",omitempty:flag:3"`
	Solution         *string             `tl:",omitempty:flag:4"`
	SolutionEntities []MessageEntity     `tl:",omitempty:flag:4"`
}

func (*PollResults) CRC() uint32 { return uint32(0xbadcc1a3) }

type PollAnswerVoters struct {
	//nolint:revive // tl works with unexported tags
	_       null `tl:"flag,bitflag"`
	Chosen  bool `tl:",omitempty:flag:0,implicit"`
	Correct bool `tl:",omitempty:flag:1,implicit"`
	Option  []byte
	Voters  int32
}

func (*PollAnswerVoters) CRC() uint32 { return uint32(0x3b6ddad2) }

type MessageEntity interface {
	Any
	_MessageEntity()
}

type AccountInstallThemeParams struct {
	//nolint:revive // tl works with unexported tags
	_      null       `tl:"flag,bitflag"`
	Dark   bool       `tl:",omitempty:flag:0,implicit"`
	Format *string    `tl:",omitempty:flag:1"`
	Theme  InputTheme `tl:",omitempty:flag:1"`
}

func (*AccountInstallThemeParams) CRC() uint32 { return uint32(0x7ae43737) }

type InputTheme interface {
	Any
	_InputTheme()
}

type InputThemeObj struct {
	ID         int64
	AccessHash int64
}

var _ InputTheme = (*InputThemeObj)(nil)

func (*InputThemeObj) CRC() uint32 { return uint32(0x3c5693e9) }

func (*InputThemeObj) _InputTheme() {}

type AccountUnregisterDeviceParams struct {
	TokenType int32
	Token     string
	OtherUids []int32
}

func (*AccountUnregisterDeviceParams) CRC() uint32 { return uint32(0x3076c4bf) }

type AnyStructWithAnyType struct {
	SomeInt int32
	Data    any
}

func (*AnyStructWithAnyType) CRC() uint32 { return uint32(0xfdfd4646) }

type AnyStructWithAnyObject struct {
	SomeInt int32
	Data    Any
}

func (*AnyStructWithAnyObject) CRC() uint32 { return uint32(0xfd46fd46) }

type InvokeWithLayerParams struct {
	Layer int32
	Query Any
}

func (*InvokeWithLayerParams) CRC() uint32 { return 0xda9b0d0d }

type InitConnectionParams struct {
	//nolint:revive // tl works with unexported tags
	_              null `tl:"flag,bitflag"`
	APIID          int32
	DeviceModel    string
	SystemVersion  string
	AppVersion     string
	SystemLangCode string
	LangPack       string
	LangCode       string
	Proxy          Any `tl:",omitempty:flag:0"`
	Params         Any `tl:",omitempty:flag:1"`

	Query Any
}

func (*InitConnectionParams) CRC() uint32 { return 0xc1cd5ea9 }

type ResPQ struct {
	Nonce        *Int128
	ServerNonce  *Int128
	Pq           []byte
	Fingerprints []int64
}

func (*ResPQ) CRC() uint32 { return 0x05162463 }

type ResPQRaw struct {
	Nonce        [16]byte
	ServerNonce  [16]byte
	Pq           []byte
	Fingerprints []int64
}

func (*ResPQRaw) CRC() uint32 { return 0x05162464 }

// issue #59 fixtures

type Poll struct {
	ID int64
	//nolint:revive // tl works with unexported tags
	_              null `tl:"flag,bitflag"`
	Closed         bool `tl:",omitempty:flag:0,implicit"`
	PublicVoters   bool `tl:",omitempty:flag:1,implicit"`
	MultipleChoice bool `tl:",omitempty:flag:2,implicit"`
	Quiz           bool `tl:",omitempty:flag:3,implicit"`
	Question       string
	Answers        []*PollAnswer
	ClosePeriod    *int32 `tl:",omitempty:flag:4"`
	CloseDate      *int32 `tl:",omitempty:flag:5"`
}

func (*Poll) CRC() uint32 { return 0x86e18161 }

type PollAnswer struct {
	Text   string
	Option []byte
}

func (*PollAnswer) CRC() uint32 { return 0x6ca9c2e9 }

type DHParamsOk struct {
	Nonce           *Int128
	ServerNonce     *Int128
	EncryptedAnswer []byte
}

func (*DHParamsOk) CRC() uint32 { return 0xd0e8075c }
