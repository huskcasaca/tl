package main

import (
	"context"
	"fmt"
	tl "github.com/xelaj/tl"
)

type AccountDaysTTL interface {
	tl.Object
	_AccountDaysTTL()
}

var (
	_ AccountDaysTTL = (*AccountDaysTTLPredict)(nil)
)

type AccountDaysTTLPredict struct {
	Days int32
}

func (*AccountDaysTTLPredict) CRC() uint32 {
	return 0xb8d0afdf
}
func (*AccountDaysTTLPredict) _AccountDaysTTL() {}

type AttachMenuBot interface {
	tl.Object
	_AttachMenuBot()
}

var (
	_ AttachMenuBot = (*AttachMenuBotPredict)(nil)
)

type AttachMenuBotPredict struct {
	_                        struct{} `tl:"flags,bitflag"`
	Inactive                 bool     `tl:",omitempty:flags:0,implicit"`
	HasSettings              bool     `tl:",omitempty:flags:1,implicit"`
	RequestWriteAccess       bool     `tl:",omitempty:flags:2,implicit"`
	ShowInAttachMenu         bool     `tl:",omitempty:flags:3,implicit"`
	ShowInSideMenu           bool     `tl:",omitempty:flags:4,implicit"`
	SideMenuDisclaimerNeeded bool     `tl:",omitempty:flags:5,implicit"`
	BotID                    int64
	ShortName                string
	PeerTypes                []AttachMenuPeerType `tl:",omitempty:flags:3"`
	Icons                    []AttachMenuBotIcon
}

func (*AttachMenuBotPredict) CRC() uint32 {
	return 0xd90d8dfe
}
func (*AttachMenuBotPredict) _AttachMenuBot() {}

type AttachMenuBotIcon interface {
	tl.Object
	_AttachMenuBotIcon()
}

var (
	_ AttachMenuBotIcon = (*AttachMenuBotIconPredict)(nil)
)

type AttachMenuBotIconPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Name   string
	Icon   Document
	Colors []AttachMenuBotIconColor `tl:",omitempty:flags:0"`
}

func (*AttachMenuBotIconPredict) CRC() uint32 {
	return 0xb2a7386b
}
func (*AttachMenuBotIconPredict) _AttachMenuBotIcon() {}

type AttachMenuBotIconColor interface {
	tl.Object
	_AttachMenuBotIconColor()
}

var (
	_ AttachMenuBotIconColor = (*AttachMenuBotIconColorPredict)(nil)
)

type AttachMenuBotIconColorPredict struct {
	Name  string
	Color int32
}

func (*AttachMenuBotIconColorPredict) CRC() uint32 {
	return 0x4576f3f0
}
func (*AttachMenuBotIconColorPredict) _AttachMenuBotIconColor() {}

type AttachMenuBots interface {
	tl.Object
	_AttachMenuBots()
}

var (
	_ AttachMenuBots = (*AttachMenuBotsNotModifiedPredict)(nil)
	_ AttachMenuBots = (*AttachMenuBotsPredict)(nil)
)

type AttachMenuBotsNotModifiedPredict struct{}

func (*AttachMenuBotsNotModifiedPredict) CRC() uint32 {
	return 0xf1d88a5c
}
func (*AttachMenuBotsNotModifiedPredict) _AttachMenuBots() {}

type AttachMenuBotsPredict struct {
	Hash  int64
	Bots  []AttachMenuBot
	Users []User
}

func (*AttachMenuBotsPredict) CRC() uint32 {
	return 0x3c4301c0
}
func (*AttachMenuBotsPredict) _AttachMenuBots() {}

type AttachMenuBotsBot interface {
	tl.Object
	_AttachMenuBotsBot()
}

var (
	_ AttachMenuBotsBot = (*AttachMenuBotsBotPredict)(nil)
)

type AttachMenuBotsBotPredict struct {
	Bot   AttachMenuBot
	Users []User
}

func (*AttachMenuBotsBotPredict) CRC() uint32 {
	return 0x93bf667f
}
func (*AttachMenuBotsBotPredict) _AttachMenuBotsBot() {}

type Authorization interface {
	tl.Object
	_Authorization()
}

var (
	_ Authorization = (*AuthorizationPredict)(nil)
)

type AuthorizationPredict struct {
	_                         struct{} `tl:"flags,bitflag"`
	Current                   bool     `tl:",omitempty:flags:0,implicit"`
	OfficialApp               bool     `tl:",omitempty:flags:1,implicit"`
	PasswordPending           bool     `tl:",omitempty:flags:2,implicit"`
	EncryptedRequestsDisabled bool     `tl:",omitempty:flags:3,implicit"`
	CallRequestsDisabled      bool     `tl:",omitempty:flags:4,implicit"`
	Unconfirmed               bool     `tl:",omitempty:flags:5,implicit"`
	Hash                      int64
	DeviceModel               string
	Platform                  string
	SystemVersion             string
	APIID                     int32
	AppName                   string
	AppVersion                string
	DateCreated               int32
	DateActive                int32
	Ip                        string
	Country                   string
	Region                    string
}

func (*AuthorizationPredict) CRC() uint32 {
	return 0xad01d61d
}
func (*AuthorizationPredict) _Authorization() {}

type AutoDownloadSettings interface {
	tl.Object
	_AutoDownloadSettings()
}

var (
	_ AutoDownloadSettings = (*AutoDownloadSettingsPredict)(nil)
)

type AutoDownloadSettingsPredict struct {
	_                             struct{} `tl:"flags,bitflag"`
	Disabled                      bool     `tl:",omitempty:flags:0,implicit"`
	VideoPreloadLarge             bool     `tl:",omitempty:flags:1,implicit"`
	AudioPreloadNext              bool     `tl:",omitempty:flags:2,implicit"`
	PhonecallsLessData            bool     `tl:",omitempty:flags:3,implicit"`
	StoriesPreload                bool     `tl:",omitempty:flags:4,implicit"`
	PhotoSizeMax                  int32
	VideoSizeMax                  int64
	FileSizeMax                   int64
	VideoUploadMaxbitrate         int32
	SmallQueueActiveOperationsMax int32
	LargeQueueActiveOperationsMax int32
}

func (*AutoDownloadSettingsPredict) CRC() uint32 {
	return 0xbaa57628
}
func (*AutoDownloadSettingsPredict) _AutoDownloadSettings() {}

type AutoSaveException interface {
	tl.Object
	_AutoSaveException()
}

var (
	_ AutoSaveException = (*AutoSaveExceptionPredict)(nil)
)

type AutoSaveExceptionPredict struct {
	Peer     Peer
	Settings AutoSaveSettings
}

func (*AutoSaveExceptionPredict) CRC() uint32 {
	return 0x81602d47
}
func (*AutoSaveExceptionPredict) _AutoSaveException() {}

type AutoSaveSettings interface {
	tl.Object
	_AutoSaveSettings()
}

var (
	_ AutoSaveSettings = (*AutoSaveSettingsPredict)(nil)
)

type AutoSaveSettingsPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Photos       bool     `tl:",omitempty:flags:0,implicit"`
	Videos       bool     `tl:",omitempty:flags:1,implicit"`
	VideoMaxSize *int64   `tl:",omitempty:flags:2"`
}

func (*AutoSaveSettingsPredict) CRC() uint32 {
	return 0xc84834ce
}
func (*AutoSaveSettingsPredict) _AutoSaveSettings() {}

type AvailableEffect interface {
	tl.Object
	_AvailableEffect()
}

var (
	_ AvailableEffect = (*AvailableEffectPredict)(nil)
)

type AvailableEffectPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	PremiumRequired   bool     `tl:",omitempty:flags:2,implicit"`
	ID                int64
	Emoticon          string
	StaticIconID      *int64 `tl:",omitempty:flags:0"`
	EffectStickerID   int64
	EffectAnimationID *int64 `tl:",omitempty:flags:1"`
}

func (*AvailableEffectPredict) CRC() uint32 {
	return 0x93c3e27e
}
func (*AvailableEffectPredict) _AvailableEffect() {}

type AvailableReaction interface {
	tl.Object
	_AvailableReaction()
}

var (
	_ AvailableReaction = (*AvailableReactionPredict)(nil)
)

type AvailableReactionPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Inactive          bool     `tl:",omitempty:flags:0,implicit"`
	Premium           bool     `tl:",omitempty:flags:2,implicit"`
	Reaction          string
	Title             string
	StaticIcon        Document
	AppearAnimation   Document
	SelectAnimation   Document
	ActivateAnimation Document
	EffectAnimation   Document
	AroundAnimation   Document `tl:",omitempty:flags:1"`
	CenterIcon        Document `tl:",omitempty:flags:1"`
}

func (*AvailableReactionPredict) CRC() uint32 {
	return 0xc077ec01
}
func (*AvailableReactionPredict) _AvailableReaction() {}

type BankCardOpenURL interface {
	tl.Object
	_BankCardOpenURL()
}

var (
	_ BankCardOpenURL = (*BankCardOpenURLPredict)(nil)
)

type BankCardOpenURLPredict struct {
	URL  string
	Name string
}

func (*BankCardOpenURLPredict) CRC() uint32 {
	return 0xf568028a
}
func (*BankCardOpenURLPredict) _BankCardOpenURL() {}

type Birthday interface {
	tl.Object
	_Birthday()
}

var (
	_ Birthday = (*BirthdayPredict)(nil)
)

type BirthdayPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Day   int32
	Month int32
	Year  *int32 `tl:",omitempty:flags:0"`
}

func (*BirthdayPredict) CRC() uint32 {
	return 0x6c8e1e06
}
func (*BirthdayPredict) _Birthday() {}

type Boost interface {
	tl.Object
	_Boost()
}

var (
	_ Boost = (*BoostPredict)(nil)
)

type BoostPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Gift          bool     `tl:",omitempty:flags:1,implicit"`
	Giveaway      bool     `tl:",omitempty:flags:2,implicit"`
	Unclaimed     bool     `tl:",omitempty:flags:3,implicit"`
	ID            string
	UserID        *int64 `tl:",omitempty:flags:0"`
	GiveawayMsgID *int32 `tl:",omitempty:flags:2"`
	Date          int32
	Expires       int32
	UsedGiftSlug  *string `tl:",omitempty:flags:4"`
	Multiplier    *int32  `tl:",omitempty:flags:5"`
}

func (*BoostPredict) CRC() uint32 {
	return 0x2a1c8c71
}
func (*BoostPredict) _Boost() {}

type BotApp interface {
	tl.Object
	_BotApp()
}

var (
	_ BotApp = (*BotAppNotModifiedPredict)(nil)
	_ BotApp = (*BotAppPredict)(nil)
)

type BotAppNotModifiedPredict struct{}

func (*BotAppNotModifiedPredict) CRC() uint32 {
	return 0x5da674b7
}
func (*BotAppNotModifiedPredict) _BotApp() {}

type BotAppPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          int64
	AccessHash  int64
	ShortName   string
	Title       string
	Description string
	Photo       Photo
	Document    Document `tl:",omitempty:flags:0"`
	Hash        int64
}

func (*BotAppPredict) CRC() uint32 {
	return 0x95fcd1d6
}
func (*BotAppPredict) _BotApp() {}

type BotBusinessConnection interface {
	tl.Object
	_BotBusinessConnection()
}

var (
	_ BotBusinessConnection = (*BotBusinessConnectionPredict)(nil)
)

type BotBusinessConnectionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	CanReply     bool     `tl:",omitempty:flags:0,implicit"`
	Disabled     bool     `tl:",omitempty:flags:1,implicit"`
	ConnectionID string
	UserID       int64
	DcID         int32
	Date         int32
}

func (*BotBusinessConnectionPredict) CRC() uint32 {
	return 0x896433b4
}
func (*BotBusinessConnectionPredict) _BotBusinessConnection() {}

type BotCommand interface {
	tl.Object
	_BotCommand()
}

var (
	_ BotCommand = (*BotCommandPredict)(nil)
)

type BotCommandPredict struct {
	Command     string
	Description string
}

func (*BotCommandPredict) CRC() uint32 {
	return 0xc27ac8c7
}
func (*BotCommandPredict) _BotCommand() {}

type BotCommandScope interface {
	tl.Object
	_BotCommandScope()
}

var (
	_ BotCommandScope = (*BotCommandScopeDefaultPredict)(nil)
	_ BotCommandScope = (*BotCommandScopeUsersPredict)(nil)
	_ BotCommandScope = (*BotCommandScopeChatsPredict)(nil)
	_ BotCommandScope = (*BotCommandScopeChatAdminsPredict)(nil)
	_ BotCommandScope = (*BotCommandScopePeerPredict)(nil)
	_ BotCommandScope = (*BotCommandScopePeerAdminsPredict)(nil)
	_ BotCommandScope = (*BotCommandScopePeerUserPredict)(nil)
)

type BotCommandScopeDefaultPredict struct{}

func (*BotCommandScopeDefaultPredict) CRC() uint32 {
	return 0x2f6cb2ab
}
func (*BotCommandScopeDefaultPredict) _BotCommandScope() {}

type BotCommandScopeUsersPredict struct{}

func (*BotCommandScopeUsersPredict) CRC() uint32 {
	return 0x3c4f04d8
}
func (*BotCommandScopeUsersPredict) _BotCommandScope() {}

type BotCommandScopeChatsPredict struct{}

func (*BotCommandScopeChatsPredict) CRC() uint32 {
	return 0x6fe1a881
}
func (*BotCommandScopeChatsPredict) _BotCommandScope() {}

type BotCommandScopeChatAdminsPredict struct{}

func (*BotCommandScopeChatAdminsPredict) CRC() uint32 {
	return 0xb9aa606a
}
func (*BotCommandScopeChatAdminsPredict) _BotCommandScope() {}

type BotCommandScopePeerPredict struct {
	Peer InputPeer
}

func (*BotCommandScopePeerPredict) CRC() uint32 {
	return 0xdb9d897d
}
func (*BotCommandScopePeerPredict) _BotCommandScope() {}

type BotCommandScopePeerAdminsPredict struct {
	Peer InputPeer
}

func (*BotCommandScopePeerAdminsPredict) CRC() uint32 {
	return 0x3fd863d1
}
func (*BotCommandScopePeerAdminsPredict) _BotCommandScope() {}

type BotCommandScopePeerUserPredict struct {
	Peer   InputPeer
	UserID InputUser
}

func (*BotCommandScopePeerUserPredict) CRC() uint32 {
	return 0xa1321f3
}
func (*BotCommandScopePeerUserPredict) _BotCommandScope() {}

type BotInfo interface {
	tl.Object
	_BotInfo()
}

var (
	_ BotInfo = (*BotInfoPredict)(nil)
)

type BotInfoPredict struct {
	_                   struct{}      `tl:"flags,bitflag"`
	HasPreviewMedias    bool          `tl:",omitempty:flags:6,implicit"`
	UserID              *int64        `tl:",omitempty:flags:0"`
	Description         *string       `tl:",omitempty:flags:1"`
	DescriptionPhoto    Photo         `tl:",omitempty:flags:4"`
	DescriptionDocument Document      `tl:",omitempty:flags:5"`
	Commands            []BotCommand  `tl:",omitempty:flags:2"`
	MenuButton          BotMenuButton `tl:",omitempty:flags:3"`
}

func (*BotInfoPredict) CRC() uint32 {
	return 0x8f300b57
}
func (*BotInfoPredict) _BotInfo() {}

type BotInlineMessage interface {
	tl.Object
	_BotInlineMessage()
}

var (
	_ BotInlineMessage = (*BotInlineMessageMediaAutoPredict)(nil)
	_ BotInlineMessage = (*BotInlineMessageTextPredict)(nil)
	_ BotInlineMessage = (*BotInlineMessageMediaGeoPredict)(nil)
	_ BotInlineMessage = (*BotInlineMessageMediaVenuePredict)(nil)
	_ BotInlineMessage = (*BotInlineMessageMediaContactPredict)(nil)
	_ BotInlineMessage = (*BotInlineMessageMediaInvoicePredict)(nil)
	_ BotInlineMessage = (*BotInlineMessageMediaWebPagePredict)(nil)
)

type BotInlineMessageMediaAutoPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []MessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaAutoPredict) CRC() uint32 {
	return 0x764cf810
}
func (*BotInlineMessageMediaAutoPredict) _BotInlineMessage() {}

type BotInlineMessageTextPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	NoWebpage   bool     `tl:",omitempty:flags:0,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []MessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageTextPredict) CRC() uint32 {
	return 0x8c7f65e2
}
func (*BotInlineMessageTextPredict) _BotInlineMessage() {}

type BotInlineMessageMediaGeoPredict struct {
	_                           struct{} `tl:"flags,bitflag"`
	Geo                         GeoPoint
	Heading                     *int32      `tl:",omitempty:flags:0"`
	Period                      *int32      `tl:",omitempty:flags:1"`
	ProximityNotificationRadius *int32      `tl:",omitempty:flags:3"`
	ReplyMarkup                 ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaGeoPredict) CRC() uint32 {
	return 0x51846fd
}
func (*BotInlineMessageMediaGeoPredict) _BotInlineMessage() {}

type BotInlineMessageMediaVenuePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Geo         GeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaVenuePredict) CRC() uint32 {
	return 0x8a86659c
}
func (*BotInlineMessageMediaVenuePredict) _BotInlineMessage() {}

type BotInlineMessageMediaContactPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaContactPredict) CRC() uint32 {
	return 0x18d1cdc2
}
func (*BotInlineMessageMediaContactPredict) _BotInlineMessage() {}

type BotInlineMessageMediaInvoicePredict struct {
	_                        struct{} `tl:"flags,bitflag"`
	ShippingAddressRequested bool     `tl:",omitempty:flags:1,implicit"`
	Test                     bool     `tl:",omitempty:flags:3,implicit"`
	Title                    string
	Description              string
	Photo                    WebDocument `tl:",omitempty:flags:0"`
	Currency                 string
	TotalAmount              int64
	ReplyMarkup              ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaInvoicePredict) CRC() uint32 {
	return 0x354a9b09
}
func (*BotInlineMessageMediaInvoicePredict) _BotInlineMessage() {}

type BotInlineMessageMediaWebPagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	InvertMedia     bool     `tl:",omitempty:flags:3,implicit"`
	ForceLargeMedia bool     `tl:",omitempty:flags:4,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:5,implicit"`
	Manual          bool     `tl:",omitempty:flags:7,implicit"`
	Safe            bool     `tl:",omitempty:flags:8,implicit"`
	Message         string
	Entities        []MessageEntity `tl:",omitempty:flags:1"`
	URL             string
	ReplyMarkup     ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*BotInlineMessageMediaWebPagePredict) CRC() uint32 {
	return 0x809ad9a6
}
func (*BotInlineMessageMediaWebPagePredict) _BotInlineMessage() {}

type BotInlineResult interface {
	tl.Object
	_BotInlineResult()
}

var (
	_ BotInlineResult = (*BotInlineResultPredict)(nil)
	_ BotInlineResult = (*BotInlineMediaResultPredict)(nil)
)

type BotInlineResultPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Title       *string     `tl:",omitempty:flags:1"`
	Description *string     `tl:",omitempty:flags:2"`
	URL         *string     `tl:",omitempty:flags:3"`
	Thumb       WebDocument `tl:",omitempty:flags:4"`
	Content     WebDocument `tl:",omitempty:flags:5"`
	SendMessage BotInlineMessage
}

func (*BotInlineResultPredict) CRC() uint32 {
	return 0x11965f3a
}
func (*BotInlineResultPredict) _BotInlineResult() {}

type BotInlineMediaResultPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Photo       Photo    `tl:",omitempty:flags:0"`
	Document    Document `tl:",omitempty:flags:1"`
	Title       *string  `tl:",omitempty:flags:2"`
	Description *string  `tl:",omitempty:flags:3"`
	SendMessage BotInlineMessage
}

func (*BotInlineMediaResultPredict) CRC() uint32 {
	return 0x17db940b
}
func (*BotInlineMediaResultPredict) _BotInlineResult() {}

type BotMenuButton interface {
	tl.Object
	_BotMenuButton()
}

var (
	_ BotMenuButton = (*BotMenuButtonDefaultPredict)(nil)
	_ BotMenuButton = (*BotMenuButtonCommandsPredict)(nil)
	_ BotMenuButton = (*BotMenuButtonPredict)(nil)
)

type BotMenuButtonDefaultPredict struct{}

func (*BotMenuButtonDefaultPredict) CRC() uint32 {
	return 0x7533a588
}
func (*BotMenuButtonDefaultPredict) _BotMenuButton() {}

type BotMenuButtonCommandsPredict struct{}

func (*BotMenuButtonCommandsPredict) CRC() uint32 {
	return 0x4258c205
}
func (*BotMenuButtonCommandsPredict) _BotMenuButton() {}

type BotMenuButtonPredict struct {
	Text string
	URL  string
}

func (*BotMenuButtonPredict) CRC() uint32 {
	return 0xc7b57ce6
}
func (*BotMenuButtonPredict) _BotMenuButton() {}

type BotPreviewMedia interface {
	tl.Object
	_BotPreviewMedia()
}

var (
	_ BotPreviewMedia = (*BotPreviewMediaPredict)(nil)
)

type BotPreviewMediaPredict struct {
	Date  int32
	Media MessageMedia
}

func (*BotPreviewMediaPredict) CRC() uint32 {
	return 0x23e91ba3
}
func (*BotPreviewMediaPredict) _BotPreviewMedia() {}

type BroadcastRevenueBalances interface {
	tl.Object
	_BroadcastRevenueBalances()
}

var (
	_ BroadcastRevenueBalances = (*BroadcastRevenueBalancesPredict)(nil)
)

type BroadcastRevenueBalancesPredict struct {
	CurrentBalance   int64
	AvailableBalance int64
	OverallRevenue   int64
}

func (*BroadcastRevenueBalancesPredict) CRC() uint32 {
	return 0x8438f1c6
}
func (*BroadcastRevenueBalancesPredict) _BroadcastRevenueBalances() {}

type BroadcastRevenueTransaction interface {
	tl.Object
	_BroadcastRevenueTransaction()
}

var (
	_ BroadcastRevenueTransaction = (*BroadcastRevenueTransactionProceedsPredict)(nil)
	_ BroadcastRevenueTransaction = (*BroadcastRevenueTransactionWithdrawalPredict)(nil)
	_ BroadcastRevenueTransaction = (*BroadcastRevenueTransactionRefundPredict)(nil)
)

type BroadcastRevenueTransactionProceedsPredict struct {
	Amount   int64
	FromDate int32
	ToDate   int32
}

func (*BroadcastRevenueTransactionProceedsPredict) CRC() uint32 {
	return 0x557e2cc4
}
func (*BroadcastRevenueTransactionProceedsPredict) _BroadcastRevenueTransaction() {}

type BroadcastRevenueTransactionWithdrawalPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Pending         bool     `tl:",omitempty:flags:0,implicit"`
	Failed          bool     `tl:",omitempty:flags:2,implicit"`
	Amount          int64
	Date            int32
	Provider        string
	TransactionDate *int32  `tl:",omitempty:flags:1"`
	TransactionURL  *string `tl:",omitempty:flags:1"`
}

func (*BroadcastRevenueTransactionWithdrawalPredict) CRC() uint32 {
	return 0x5a590978
}
func (*BroadcastRevenueTransactionWithdrawalPredict) _BroadcastRevenueTransaction() {}

type BroadcastRevenueTransactionRefundPredict struct {
	Amount   int64
	Date     int32
	Provider string
}

func (*BroadcastRevenueTransactionRefundPredict) CRC() uint32 {
	return 0x42d30d2e
}
func (*BroadcastRevenueTransactionRefundPredict) _BroadcastRevenueTransaction() {}

type BusinessAwayMessage interface {
	tl.Object
	_BusinessAwayMessage()
}

var (
	_ BusinessAwayMessage = (*BusinessAwayMessagePredict)(nil)
)

type BusinessAwayMessagePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	OfflineOnly bool     `tl:",omitempty:flags:0,implicit"`
	ShortcutID  int32
	Schedule    BusinessAwayMessageSchedule
	Recipients  BusinessRecipients
}

func (*BusinessAwayMessagePredict) CRC() uint32 {
	return 0xef156a5c
}
func (*BusinessAwayMessagePredict) _BusinessAwayMessage() {}

type BusinessAwayMessageSchedule interface {
	tl.Object
	_BusinessAwayMessageSchedule()
}

var (
	_ BusinessAwayMessageSchedule = (*BusinessAwayMessageScheduleAlwaysPredict)(nil)
	_ BusinessAwayMessageSchedule = (*BusinessAwayMessageScheduleOutsideWorkHoursPredict)(nil)
	_ BusinessAwayMessageSchedule = (*BusinessAwayMessageScheduleCustomPredict)(nil)
)

type BusinessAwayMessageScheduleAlwaysPredict struct{}

func (*BusinessAwayMessageScheduleAlwaysPredict) CRC() uint32 {
	return 0xc9b9e2b9
}
func (*BusinessAwayMessageScheduleAlwaysPredict) _BusinessAwayMessageSchedule() {}

type BusinessAwayMessageScheduleOutsideWorkHoursPredict struct{}

func (*BusinessAwayMessageScheduleOutsideWorkHoursPredict) CRC() uint32 {
	return 0xc3f2f501
}
func (*BusinessAwayMessageScheduleOutsideWorkHoursPredict) _BusinessAwayMessageSchedule() {}

type BusinessAwayMessageScheduleCustomPredict struct {
	StartDate int32
	EndDate   int32
}

func (*BusinessAwayMessageScheduleCustomPredict) CRC() uint32 {
	return 0xcc4d9ecc
}
func (*BusinessAwayMessageScheduleCustomPredict) _BusinessAwayMessageSchedule() {}

type BusinessBotRecipients interface {
	tl.Object
	_BusinessBotRecipients()
}

var (
	_ BusinessBotRecipients = (*BusinessBotRecipientsPredict)(nil)
)

type BusinessBotRecipientsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ExistingChats   bool     `tl:",omitempty:flags:0,implicit"`
	NewChats        bool     `tl:",omitempty:flags:1,implicit"`
	Contacts        bool     `tl:",omitempty:flags:2,implicit"`
	NonContacts     bool     `tl:",omitempty:flags:3,implicit"`
	ExcludeSelected bool     `tl:",omitempty:flags:5,implicit"`
	Users           []int64  `tl:",omitempty:flags:4"`
	ExcludeUsers    []int64  `tl:",omitempty:flags:6"`
}

func (*BusinessBotRecipientsPredict) CRC() uint32 {
	return 0xb88cf373
}
func (*BusinessBotRecipientsPredict) _BusinessBotRecipients() {}

type BusinessChatLink interface {
	tl.Object
	_BusinessChatLink()
}

var (
	_ BusinessChatLink = (*BusinessChatLinkPredict)(nil)
)

type BusinessChatLinkPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Link     string
	Message  string
	Entities []MessageEntity `tl:",omitempty:flags:0"`
	Title    *string         `tl:",omitempty:flags:1"`
	Views    int32
}

func (*BusinessChatLinkPredict) CRC() uint32 {
	return 0xb4ae666f
}
func (*BusinessChatLinkPredict) _BusinessChatLink() {}

type BusinessGreetingMessage interface {
	tl.Object
	_BusinessGreetingMessage()
}

var (
	_ BusinessGreetingMessage = (*BusinessGreetingMessagePredict)(nil)
)

type BusinessGreetingMessagePredict struct {
	ShortcutID     int32
	Recipients     BusinessRecipients
	NoActivityDays int32
}

func (*BusinessGreetingMessagePredict) CRC() uint32 {
	return 0xe519abab
}
func (*BusinessGreetingMessagePredict) _BusinessGreetingMessage() {}

type BusinessIntro interface {
	tl.Object
	_BusinessIntro()
}

var (
	_ BusinessIntro = (*BusinessIntroPredict)(nil)
)

type BusinessIntroPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       string
	Description string
	Sticker     Document `tl:",omitempty:flags:0"`
}

func (*BusinessIntroPredict) CRC() uint32 {
	return 0x5a0a066d
}
func (*BusinessIntroPredict) _BusinessIntro() {}

type BusinessLocation interface {
	tl.Object
	_BusinessLocation()
}

var (
	_ BusinessLocation = (*BusinessLocationPredict)(nil)
)

type BusinessLocationPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	GeoPoint GeoPoint `tl:",omitempty:flags:0"`
	Address  string
}

func (*BusinessLocationPredict) CRC() uint32 {
	return 0xac5c1af7
}
func (*BusinessLocationPredict) _BusinessLocation() {}

type BusinessRecipients interface {
	tl.Object
	_BusinessRecipients()
}

var (
	_ BusinessRecipients = (*BusinessRecipientsPredict)(nil)
)

type BusinessRecipientsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ExistingChats   bool     `tl:",omitempty:flags:0,implicit"`
	NewChats        bool     `tl:",omitempty:flags:1,implicit"`
	Contacts        bool     `tl:",omitempty:flags:2,implicit"`
	NonContacts     bool     `tl:",omitempty:flags:3,implicit"`
	ExcludeSelected bool     `tl:",omitempty:flags:5,implicit"`
	Users           []int64  `tl:",omitempty:flags:4"`
}

func (*BusinessRecipientsPredict) CRC() uint32 {
	return 0x21108ff7
}
func (*BusinessRecipientsPredict) _BusinessRecipients() {}

type BusinessWeeklyOpen interface {
	tl.Object
	_BusinessWeeklyOpen()
}

var (
	_ BusinessWeeklyOpen = (*BusinessWeeklyOpenPredict)(nil)
)

type BusinessWeeklyOpenPredict struct {
	StartMinute int32
	EndMinute   int32
}

func (*BusinessWeeklyOpenPredict) CRC() uint32 {
	return 0x120b1ab9
}
func (*BusinessWeeklyOpenPredict) _BusinessWeeklyOpen() {}

type BusinessWorkHours interface {
	tl.Object
	_BusinessWorkHours()
}

var (
	_ BusinessWorkHours = (*BusinessWorkHoursPredict)(nil)
)

type BusinessWorkHoursPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	OpenNow    bool     `tl:",omitempty:flags:0,implicit"`
	TimezoneID string
	WeeklyOpen []BusinessWeeklyOpen
}

func (*BusinessWorkHoursPredict) CRC() uint32 {
	return 0x8c92b098
}
func (*BusinessWorkHoursPredict) _BusinessWorkHours() {}

type CdnConfig interface {
	tl.Object
	_CdnConfig()
}

var (
	_ CdnConfig = (*CdnConfigPredict)(nil)
)

type CdnConfigPredict struct {
	PublicKeys []CdnPublicKey
}

func (*CdnConfigPredict) CRC() uint32 {
	return 0x5725e40a
}
func (*CdnConfigPredict) _CdnConfig() {}

type CdnPublicKey interface {
	tl.Object
	_CdnPublicKey()
}

var (
	_ CdnPublicKey = (*CdnPublicKeyPredict)(nil)
)

type CdnPublicKeyPredict struct {
	DcID      int32
	PublicKey string
}

func (*CdnPublicKeyPredict) CRC() uint32 {
	return 0xc982eaba
}
func (*CdnPublicKeyPredict) _CdnPublicKey() {}

type ChannelAdminLogEvent interface {
	tl.Object
	_ChannelAdminLogEvent()
}

var (
	_ ChannelAdminLogEvent = (*ChannelAdminLogEventPredict)(nil)
)

type ChannelAdminLogEventPredict struct {
	ID     int64
	Date   int32
	UserID int64
	Action ChannelAdminLogEventAction
}

func (*ChannelAdminLogEventPredict) CRC() uint32 {
	return 0x1fad68cd
}
func (*ChannelAdminLogEventPredict) _ChannelAdminLogEvent() {}

type ChannelAdminLogEventAction interface {
	tl.Object
	_ChannelAdminLogEventAction()
}

var (
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeTitlePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeAboutPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeUsernamePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangePhotoPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleInvitesPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleSignaturesPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionUpdatePinnedPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionEditMessagePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionDeleteMessagePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantJoinPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantLeavePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantInvitePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantToggleBanPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantToggleAdminPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeStickerSetPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionTogglePreHistoryHiddenPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionDefaultBannedRightsPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionStopPollPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeLinkedChatPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeLocationPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleSlowModePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionStartGroupCallPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionDiscardGroupCallPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantMutePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantUnmutePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleGroupCallSettingPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantJoinByInvitePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionExportedInviteDeletePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionExportedInviteRevokePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionExportedInviteEditPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantVolumePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeHistoryTTLPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionParticipantJoinByRequestPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleNoForwardsPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionSendMessagePredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeAvailableReactionsPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeUsernamesPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleForumPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionCreateTopicPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionEditTopicPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionDeleteTopicPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionPinTopicPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionToggleAntiSpamPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangePeerColorPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeProfilePeerColorPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeWallpaperPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeEmojiStatusPredict)(nil)
	_ ChannelAdminLogEventAction = (*ChannelAdminLogEventActionChangeEmojiStickerSetPredict)(nil)
)

type ChannelAdminLogEventActionChangeTitlePredict struct {
	PrevValue string
	NewValue  string
}

func (*ChannelAdminLogEventActionChangeTitlePredict) CRC() uint32 {
	return 0xe6dfb825
}
func (*ChannelAdminLogEventActionChangeTitlePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeAboutPredict struct {
	PrevValue string
	NewValue  string
}

func (*ChannelAdminLogEventActionChangeAboutPredict) CRC() uint32 {
	return 0x55188a2e
}
func (*ChannelAdminLogEventActionChangeAboutPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeUsernamePredict struct {
	PrevValue string
	NewValue  string
}

func (*ChannelAdminLogEventActionChangeUsernamePredict) CRC() uint32 {
	return 0x6a4afc38
}
func (*ChannelAdminLogEventActionChangeUsernamePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangePhotoPredict struct {
	PrevPhoto Photo
	NewPhoto  Photo
}

func (*ChannelAdminLogEventActionChangePhotoPredict) CRC() uint32 {
	return 0x434bd2af
}
func (*ChannelAdminLogEventActionChangePhotoPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleInvitesPredict struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleInvitesPredict) CRC() uint32 {
	return 0x1b7907ae
}
func (*ChannelAdminLogEventActionToggleInvitesPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleSignaturesPredict struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleSignaturesPredict) CRC() uint32 {
	return 0x26ae0971
}
func (*ChannelAdminLogEventActionToggleSignaturesPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionUpdatePinnedPredict struct {
	Message Message
}

func (*ChannelAdminLogEventActionUpdatePinnedPredict) CRC() uint32 {
	return 0xe9e82c18
}
func (*ChannelAdminLogEventActionUpdatePinnedPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionEditMessagePredict struct {
	PrevMessage Message
	NewMessage  Message
}

func (*ChannelAdminLogEventActionEditMessagePredict) CRC() uint32 {
	return 0x709b2405
}
func (*ChannelAdminLogEventActionEditMessagePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDeleteMessagePredict struct {
	Message Message
}

func (*ChannelAdminLogEventActionDeleteMessagePredict) CRC() uint32 {
	return 0x42e047bb
}
func (*ChannelAdminLogEventActionDeleteMessagePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoinPredict struct{}

func (*ChannelAdminLogEventActionParticipantJoinPredict) CRC() uint32 {
	return 0x183040d3
}
func (*ChannelAdminLogEventActionParticipantJoinPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantLeavePredict struct{}

func (*ChannelAdminLogEventActionParticipantLeavePredict) CRC() uint32 {
	return 0xf89777f2
}
func (*ChannelAdminLogEventActionParticipantLeavePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantInvitePredict struct {
	Participant ChannelParticipant
}

func (*ChannelAdminLogEventActionParticipantInvitePredict) CRC() uint32 {
	return 0xe31c34d8
}
func (*ChannelAdminLogEventActionParticipantInvitePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantToggleBanPredict struct {
	PrevParticipant ChannelParticipant
	NewParticipant  ChannelParticipant
}

func (*ChannelAdminLogEventActionParticipantToggleBanPredict) CRC() uint32 {
	return 0xe6d83d7e
}
func (*ChannelAdminLogEventActionParticipantToggleBanPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantToggleAdminPredict struct {
	PrevParticipant ChannelParticipant
	NewParticipant  ChannelParticipant
}

func (*ChannelAdminLogEventActionParticipantToggleAdminPredict) CRC() uint32 {
	return 0xd5676710
}
func (*ChannelAdminLogEventActionParticipantToggleAdminPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeStickerSetPredict struct {
	PrevStickerset InputStickerSet
	NewStickerset  InputStickerSet
}

func (*ChannelAdminLogEventActionChangeStickerSetPredict) CRC() uint32 {
	return 0xb1c3caa7
}
func (*ChannelAdminLogEventActionChangeStickerSetPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionTogglePreHistoryHiddenPredict struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionTogglePreHistoryHiddenPredict) CRC() uint32 {
	return 0x5f5c95f1
}
func (*ChannelAdminLogEventActionTogglePreHistoryHiddenPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDefaultBannedRightsPredict struct {
	PrevBannedRights ChatBannedRights
	NewBannedRights  ChatBannedRights
}

func (*ChannelAdminLogEventActionDefaultBannedRightsPredict) CRC() uint32 {
	return 0x2df5fc0a
}
func (*ChannelAdminLogEventActionDefaultBannedRightsPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionStopPollPredict struct {
	Message Message
}

func (*ChannelAdminLogEventActionStopPollPredict) CRC() uint32 {
	return 0x8f079643
}
func (*ChannelAdminLogEventActionStopPollPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeLinkedChatPredict struct {
	PrevValue int64
	NewValue  int64
}

func (*ChannelAdminLogEventActionChangeLinkedChatPredict) CRC() uint32 {
	return 0x50c7ac8
}
func (*ChannelAdminLogEventActionChangeLinkedChatPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeLocationPredict struct {
	PrevValue ChannelLocation
	NewValue  ChannelLocation
}

func (*ChannelAdminLogEventActionChangeLocationPredict) CRC() uint32 {
	return 0xe6b76ae
}
func (*ChannelAdminLogEventActionChangeLocationPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleSlowModePredict struct {
	PrevValue int32
	NewValue  int32
}

func (*ChannelAdminLogEventActionToggleSlowModePredict) CRC() uint32 {
	return 0x53909779
}
func (*ChannelAdminLogEventActionToggleSlowModePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionStartGroupCallPredict struct {
	Call InputGroupCall
}

func (*ChannelAdminLogEventActionStartGroupCallPredict) CRC() uint32 {
	return 0x23209745
}
func (*ChannelAdminLogEventActionStartGroupCallPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDiscardGroupCallPredict struct {
	Call InputGroupCall
}

func (*ChannelAdminLogEventActionDiscardGroupCallPredict) CRC() uint32 {
	return 0xdb9f9140
}
func (*ChannelAdminLogEventActionDiscardGroupCallPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantMutePredict struct {
	Participant GroupCallParticipant
}

func (*ChannelAdminLogEventActionParticipantMutePredict) CRC() uint32 {
	return 0xf92424d2
}
func (*ChannelAdminLogEventActionParticipantMutePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantUnmutePredict struct {
	Participant GroupCallParticipant
}

func (*ChannelAdminLogEventActionParticipantUnmutePredict) CRC() uint32 {
	return 0xe64429c0
}
func (*ChannelAdminLogEventActionParticipantUnmutePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleGroupCallSettingPredict struct {
	JoinMuted bool
}

func (*ChannelAdminLogEventActionToggleGroupCallSettingPredict) CRC() uint32 {
	return 0x56d6a247
}
func (*ChannelAdminLogEventActionToggleGroupCallSettingPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoinByInvitePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ViaChatlist bool     `tl:",omitempty:flags:0,implicit"`
	Invite      ExportedChatInvite
}

func (*ChannelAdminLogEventActionParticipantJoinByInvitePredict) CRC() uint32 {
	return 0xfe9fc158
}
func (*ChannelAdminLogEventActionParticipantJoinByInvitePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteDeletePredict struct {
	Invite ExportedChatInvite
}

func (*ChannelAdminLogEventActionExportedInviteDeletePredict) CRC() uint32 {
	return 0x5a50fca4
}
func (*ChannelAdminLogEventActionExportedInviteDeletePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteRevokePredict struct {
	Invite ExportedChatInvite
}

func (*ChannelAdminLogEventActionExportedInviteRevokePredict) CRC() uint32 {
	return 0x410a134e
}
func (*ChannelAdminLogEventActionExportedInviteRevokePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteEditPredict struct {
	PrevInvite ExportedChatInvite
	NewInvite  ExportedChatInvite
}

func (*ChannelAdminLogEventActionExportedInviteEditPredict) CRC() uint32 {
	return 0xe90ebb59
}
func (*ChannelAdminLogEventActionExportedInviteEditPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantVolumePredict struct {
	Participant GroupCallParticipant
}

func (*ChannelAdminLogEventActionParticipantVolumePredict) CRC() uint32 {
	return 0x3e7f6847
}
func (*ChannelAdminLogEventActionParticipantVolumePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeHistoryTTLPredict struct {
	PrevValue int32
	NewValue  int32
}

func (*ChannelAdminLogEventActionChangeHistoryTTLPredict) CRC() uint32 {
	return 0x6e941a38
}
func (*ChannelAdminLogEventActionChangeHistoryTTLPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoinByRequestPredict struct {
	Invite     ExportedChatInvite
	ApprovedBy int64
}

func (*ChannelAdminLogEventActionParticipantJoinByRequestPredict) CRC() uint32 {
	return 0xafb6144a
}
func (*ChannelAdminLogEventActionParticipantJoinByRequestPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleNoForwardsPredict struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleNoForwardsPredict) CRC() uint32 {
	return 0xcb2ac766
}
func (*ChannelAdminLogEventActionToggleNoForwardsPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionSendMessagePredict struct {
	Message Message
}

func (*ChannelAdminLogEventActionSendMessagePredict) CRC() uint32 {
	return 0x278f2868
}
func (*ChannelAdminLogEventActionSendMessagePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeAvailableReactionsPredict struct {
	PrevValue ChatReactions
	NewValue  ChatReactions
}

func (*ChannelAdminLogEventActionChangeAvailableReactionsPredict) CRC() uint32 {
	return 0xbe4e0ef8
}
func (*ChannelAdminLogEventActionChangeAvailableReactionsPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeUsernamesPredict struct {
	PrevValue []string
	NewValue  []string
}

func (*ChannelAdminLogEventActionChangeUsernamesPredict) CRC() uint32 {
	return 0xf04fb3a9
}
func (*ChannelAdminLogEventActionChangeUsernamesPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleForumPredict struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleForumPredict) CRC() uint32 {
	return 0x2cc6383
}
func (*ChannelAdminLogEventActionToggleForumPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionCreateTopicPredict struct {
	Topic ForumTopic
}

func (*ChannelAdminLogEventActionCreateTopicPredict) CRC() uint32 {
	return 0x58707d28
}
func (*ChannelAdminLogEventActionCreateTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionEditTopicPredict struct {
	PrevTopic ForumTopic
	NewTopic  ForumTopic
}

func (*ChannelAdminLogEventActionEditTopicPredict) CRC() uint32 {
	return 0xf06fe208
}
func (*ChannelAdminLogEventActionEditTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDeleteTopicPredict struct {
	Topic ForumTopic
}

func (*ChannelAdminLogEventActionDeleteTopicPredict) CRC() uint32 {
	return 0xae168909
}
func (*ChannelAdminLogEventActionDeleteTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionPinTopicPredict struct {
	_         struct{}   `tl:"flags,bitflag"`
	PrevTopic ForumTopic `tl:",omitempty:flags:0"`
	NewTopic  ForumTopic `tl:",omitempty:flags:1"`
}

func (*ChannelAdminLogEventActionPinTopicPredict) CRC() uint32 {
	return 0x5d8d353b
}
func (*ChannelAdminLogEventActionPinTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleAntiSpamPredict struct {
	NewValue bool
}

func (*ChannelAdminLogEventActionToggleAntiSpamPredict) CRC() uint32 {
	return 0x64f36dfc
}
func (*ChannelAdminLogEventActionToggleAntiSpamPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangePeerColorPredict struct {
	PrevValue PeerColor
	NewValue  PeerColor
}

func (*ChannelAdminLogEventActionChangePeerColorPredict) CRC() uint32 {
	return 0x5796e780
}
func (*ChannelAdminLogEventActionChangePeerColorPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeProfilePeerColorPredict struct {
	PrevValue PeerColor
	NewValue  PeerColor
}

func (*ChannelAdminLogEventActionChangeProfilePeerColorPredict) CRC() uint32 {
	return 0x5e477b25
}
func (*ChannelAdminLogEventActionChangeProfilePeerColorPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeWallpaperPredict struct {
	PrevValue WallPaper
	NewValue  WallPaper
}

func (*ChannelAdminLogEventActionChangeWallpaperPredict) CRC() uint32 {
	return 0x31bb5d52
}
func (*ChannelAdminLogEventActionChangeWallpaperPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeEmojiStatusPredict struct {
	PrevValue EmojiStatus
	NewValue  EmojiStatus
}

func (*ChannelAdminLogEventActionChangeEmojiStatusPredict) CRC() uint32 {
	return 0x3ea9feb1
}
func (*ChannelAdminLogEventActionChangeEmojiStatusPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeEmojiStickerSetPredict struct {
	PrevStickerset InputStickerSet
	NewStickerset  InputStickerSet
}

func (*ChannelAdminLogEventActionChangeEmojiStickerSetPredict) CRC() uint32 {
	return 0x46d840ab
}
func (*ChannelAdminLogEventActionChangeEmojiStickerSetPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventsFilter interface {
	tl.Object
	_ChannelAdminLogEventsFilter()
}

var (
	_ ChannelAdminLogEventsFilter = (*ChannelAdminLogEventsFilterPredict)(nil)
)

type ChannelAdminLogEventsFilterPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Join      bool     `tl:",omitempty:flags:0,implicit"`
	Leave     bool     `tl:",omitempty:flags:1,implicit"`
	Invite    bool     `tl:",omitempty:flags:2,implicit"`
	Ban       bool     `tl:",omitempty:flags:3,implicit"`
	Unban     bool     `tl:",omitempty:flags:4,implicit"`
	Kick      bool     `tl:",omitempty:flags:5,implicit"`
	Unkick    bool     `tl:",omitempty:flags:6,implicit"`
	Promote   bool     `tl:",omitempty:flags:7,implicit"`
	Demote    bool     `tl:",omitempty:flags:8,implicit"`
	Info      bool     `tl:",omitempty:flags:9,implicit"`
	Settings  bool     `tl:",omitempty:flags:10,implicit"`
	Pinned    bool     `tl:",omitempty:flags:11,implicit"`
	Edit      bool     `tl:",omitempty:flags:12,implicit"`
	Delete    bool     `tl:",omitempty:flags:13,implicit"`
	GroupCall bool     `tl:",omitempty:flags:14,implicit"`
	Invites   bool     `tl:",omitempty:flags:15,implicit"`
	Send      bool     `tl:",omitempty:flags:16,implicit"`
	Forums    bool     `tl:",omitempty:flags:17,implicit"`
}

func (*ChannelAdminLogEventsFilterPredict) CRC() uint32 {
	return 0xea107ae4
}
func (*ChannelAdminLogEventsFilterPredict) _ChannelAdminLogEventsFilter() {}

type ChannelLocation interface {
	tl.Object
	_ChannelLocation()
}

var (
	_ ChannelLocation = (*ChannelLocationEmptyPredict)(nil)
	_ ChannelLocation = (*ChannelLocationPredict)(nil)
)

type ChannelLocationEmptyPredict struct{}

func (*ChannelLocationEmptyPredict) CRC() uint32 {
	return 0xbfb5ad8b
}
func (*ChannelLocationEmptyPredict) _ChannelLocation() {}

type ChannelLocationPredict struct {
	GeoPoint GeoPoint
	Address  string
}

func (*ChannelLocationPredict) CRC() uint32 {
	return 0x209b82db
}
func (*ChannelLocationPredict) _ChannelLocation() {}

type ChannelMessagesFilter interface {
	tl.Object
	_ChannelMessagesFilter()
}

var (
	_ ChannelMessagesFilter = (*ChannelMessagesFilterEmptyPredict)(nil)
	_ ChannelMessagesFilter = (*ChannelMessagesFilterPredict)(nil)
)

type ChannelMessagesFilterEmptyPredict struct{}

func (*ChannelMessagesFilterEmptyPredict) CRC() uint32 {
	return 0x94d42ee7
}
func (*ChannelMessagesFilterEmptyPredict) _ChannelMessagesFilter() {}

type ChannelMessagesFilterPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	ExcludeNewMessages bool     `tl:",omitempty:flags:1,implicit"`
	Ranges             []MessageRange
}

func (*ChannelMessagesFilterPredict) CRC() uint32 {
	return 0xcd77d957
}
func (*ChannelMessagesFilterPredict) _ChannelMessagesFilter() {}

type ChannelParticipant interface {
	tl.Object
	_ChannelParticipant()
}

var (
	_ ChannelParticipant = (*ChannelParticipantPredict)(nil)
	_ ChannelParticipant = (*ChannelParticipantSelfPredict)(nil)
	_ ChannelParticipant = (*ChannelParticipantCreatorPredict)(nil)
	_ ChannelParticipant = (*ChannelParticipantAdminPredict)(nil)
	_ ChannelParticipant = (*ChannelParticipantBannedPredict)(nil)
	_ ChannelParticipant = (*ChannelParticipantLeftPredict)(nil)
)

type ChannelParticipantPredict struct {
	UserID int64
	Date   int32
}

func (*ChannelParticipantPredict) CRC() uint32 {
	return 0xc00c07c0
}
func (*ChannelParticipantPredict) _ChannelParticipant() {}

type ChannelParticipantSelfPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	ViaRequest bool     `tl:",omitempty:flags:0,implicit"`
	UserID     int64
	InviterID  int64
	Date       int32
}

func (*ChannelParticipantSelfPredict) CRC() uint32 {
	return 0x35a8bfa7
}
func (*ChannelParticipantSelfPredict) _ChannelParticipant() {}

type ChannelParticipantCreatorPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	UserID      int64
	AdminRights ChatAdminRights
	Rank        *string `tl:",omitempty:flags:0"`
}

func (*ChannelParticipantCreatorPredict) CRC() uint32 {
	return 0x2fe601d3
}
func (*ChannelParticipantCreatorPredict) _ChannelParticipant() {}

type ChannelParticipantAdminPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	CanEdit     bool     `tl:",omitempty:flags:0,implicit"`
	Self        bool     `tl:",omitempty:flags:1,implicit"`
	UserID      int64
	InviterID   *int64 `tl:",omitempty:flags:1"`
	PromotedBy  int64
	Date        int32
	AdminRights ChatAdminRights
	Rank        *string `tl:",omitempty:flags:2"`
}

func (*ChannelParticipantAdminPredict) CRC() uint32 {
	return 0x34c3bb53
}
func (*ChannelParticipantAdminPredict) _ChannelParticipant() {}

type ChannelParticipantBannedPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Left         bool     `tl:",omitempty:flags:0,implicit"`
	Peer         Peer
	KickedBy     int64
	Date         int32
	BannedRights ChatBannedRights
}

func (*ChannelParticipantBannedPredict) CRC() uint32 {
	return 0x6df8014e
}
func (*ChannelParticipantBannedPredict) _ChannelParticipant() {}

type ChannelParticipantLeftPredict struct {
	Peer Peer
}

func (*ChannelParticipantLeftPredict) CRC() uint32 {
	return 0x1b03f006
}
func (*ChannelParticipantLeftPredict) _ChannelParticipant() {}

type ChannelParticipantsFilter interface {
	tl.Object
	_ChannelParticipantsFilter()
}

var (
	_ ChannelParticipantsFilter = (*ChannelParticipantsRecentPredict)(nil)
	_ ChannelParticipantsFilter = (*ChannelParticipantsAdminsPredict)(nil)
	_ ChannelParticipantsFilter = (*ChannelParticipantsKickedPredict)(nil)
	_ ChannelParticipantsFilter = (*ChannelParticipantsBotsPredict)(nil)
	_ ChannelParticipantsFilter = (*ChannelParticipantsBannedPredict)(nil)
	_ ChannelParticipantsFilter = (*ChannelParticipantsSearchPredict)(nil)
	_ ChannelParticipantsFilter = (*ChannelParticipantsContactsPredict)(nil)
	_ ChannelParticipantsFilter = (*ChannelParticipantsMentionsPredict)(nil)
)

type ChannelParticipantsRecentPredict struct{}

func (*ChannelParticipantsRecentPredict) CRC() uint32 {
	return 0xde3f3c79
}
func (*ChannelParticipantsRecentPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsAdminsPredict struct{}

func (*ChannelParticipantsAdminsPredict) CRC() uint32 {
	return 0xb4608969
}
func (*ChannelParticipantsAdminsPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsKickedPredict struct {
	Q string
}

func (*ChannelParticipantsKickedPredict) CRC() uint32 {
	return 0xa3b54985
}
func (*ChannelParticipantsKickedPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsBotsPredict struct{}

func (*ChannelParticipantsBotsPredict) CRC() uint32 {
	return 0xb0d1865b
}
func (*ChannelParticipantsBotsPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsBannedPredict struct {
	Q string
}

func (*ChannelParticipantsBannedPredict) CRC() uint32 {
	return 0x1427a5e1
}
func (*ChannelParticipantsBannedPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsSearchPredict struct {
	Q string
}

func (*ChannelParticipantsSearchPredict) CRC() uint32 {
	return 0x656ac4b
}
func (*ChannelParticipantsSearchPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsContactsPredict struct {
	Q string
}

func (*ChannelParticipantsContactsPredict) CRC() uint32 {
	return 0xbb6ae88d
}
func (*ChannelParticipantsContactsPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsMentionsPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Q        *string  `tl:",omitempty:flags:0"`
	TopMsgID *int32   `tl:",omitempty:flags:1"`
}

func (*ChannelParticipantsMentionsPredict) CRC() uint32 {
	return 0xe04b5ceb
}
func (*ChannelParticipantsMentionsPredict) _ChannelParticipantsFilter() {}

type Chat interface {
	tl.Object
	_Chat()
}

var (
	_ Chat = (*ChatEmptyPredict)(nil)
	_ Chat = (*ChatPredict)(nil)
	_ Chat = (*ChatForbiddenPredict)(nil)
	_ Chat = (*ChannelPredict)(nil)
	_ Chat = (*ChannelForbiddenPredict)(nil)
)

type ChatEmptyPredict struct {
	ID int64
}

func (*ChatEmptyPredict) CRC() uint32 {
	return 0x29562865
}
func (*ChatEmptyPredict) _Chat() {}

type ChatPredict struct {
	_                   struct{} `tl:"flags,bitflag"`
	Creator             bool     `tl:",omitempty:flags:0,implicit"`
	Left                bool     `tl:",omitempty:flags:2,implicit"`
	Deactivated         bool     `tl:",omitempty:flags:5,implicit"`
	CallActive          bool     `tl:",omitempty:flags:23,implicit"`
	CallNotEmpty        bool     `tl:",omitempty:flags:24,implicit"`
	Noforwards          bool     `tl:",omitempty:flags:25,implicit"`
	ID                  int64
	Title               string
	Photo               ChatPhoto
	ParticipantsCount   int32
	Date                int32
	Version             int32
	MigratedTo          InputChannel     `tl:",omitempty:flags:6"`
	AdminRights         ChatAdminRights  `tl:",omitempty:flags:14"`
	DefaultBannedRights ChatBannedRights `tl:",omitempty:flags:18"`
}

func (*ChatPredict) CRC() uint32 {
	return 0x41cbf256
}
func (*ChatPredict) _Chat() {}

type ChatForbiddenPredict struct {
	ID    int64
	Title string
}

func (*ChatForbiddenPredict) CRC() uint32 {
	return 0x6592a1a7
}
func (*ChatForbiddenPredict) _Chat() {}

type ChannelPredict struct {
	_                   struct{} `tl:"flags,bitflag"`
	Creator             bool     `tl:",omitempty:flags:0,implicit"`
	Left                bool     `tl:",omitempty:flags:2,implicit"`
	Broadcast           bool     `tl:",omitempty:flags:5,implicit"`
	Verified            bool     `tl:",omitempty:flags:7,implicit"`
	Megagroup           bool     `tl:",omitempty:flags:8,implicit"`
	Restricted          bool     `tl:",omitempty:flags:9,implicit"`
	Signatures          bool     `tl:",omitempty:flags:11,implicit"`
	Min                 bool     `tl:",omitempty:flags:12,implicit"`
	Scam                bool     `tl:",omitempty:flags:19,implicit"`
	HasLink             bool     `tl:",omitempty:flags:20,implicit"`
	HasGeo              bool     `tl:",omitempty:flags:21,implicit"`
	SlowmodeEnabled     bool     `tl:",omitempty:flags:22,implicit"`
	CallActive          bool     `tl:",omitempty:flags:23,implicit"`
	CallNotEmpty        bool     `tl:",omitempty:flags:24,implicit"`
	Fake                bool     `tl:",omitempty:flags:25,implicit"`
	Gigagroup           bool     `tl:",omitempty:flags:26,implicit"`
	Noforwards          bool     `tl:",omitempty:flags:27,implicit"`
	JoinToSend          bool     `tl:",omitempty:flags:28,implicit"`
	JoinRequest         bool     `tl:",omitempty:flags:29,implicit"`
	Forum               bool     `tl:",omitempty:flags:30,implicit"`
	_                   struct{} `tl:"flags2,bitflag"`
	StoriesHidden       bool     `tl:",omitempty:flags2:1,implicit"`
	StoriesHiddenMin    bool     `tl:",omitempty:flags2:2,implicit"`
	StoriesUnavailable  bool     `tl:",omitempty:flags2:3,implicit"`
	ID                  int64
	AccessHash          *int64 `tl:",omitempty:flags:13"`
	Title               string
	Username            *string `tl:",omitempty:flags:6"`
	Photo               ChatPhoto
	Date                int32
	RestrictionReason   []RestrictionReason `tl:",omitempty:flags:9"`
	AdminRights         ChatAdminRights     `tl:",omitempty:flags:14"`
	BannedRights        ChatBannedRights    `tl:",omitempty:flags:15"`
	DefaultBannedRights ChatBannedRights    `tl:",omitempty:flags:18"`
	ParticipantsCount   *int32              `tl:",omitempty:flags:17"`
	Usernames           []Username          `tl:",omitempty:flags2:0"`
	StoriesMaxID        *int32              `tl:",omitempty:flags2:4"`
	Color               PeerColor           `tl:",omitempty:flags2:7"`
	ProfileColor        PeerColor           `tl:",omitempty:flags2:8"`
	EmojiStatus         EmojiStatus         `tl:",omitempty:flags2:9"`
	Level               *int32              `tl:",omitempty:flags2:10"`
}

func (*ChannelPredict) CRC() uint32 {
	return 0xaadfc8f
}
func (*ChannelPredict) _Chat() {}

type ChannelForbiddenPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Broadcast  bool     `tl:",omitempty:flags:5,implicit"`
	Megagroup  bool     `tl:",omitempty:flags:8,implicit"`
	ID         int64
	AccessHash int64
	Title      string
	UntilDate  *int32 `tl:",omitempty:flags:16"`
}

func (*ChannelForbiddenPredict) CRC() uint32 {
	return 0x17d493d5
}
func (*ChannelForbiddenPredict) _Chat() {}

type ChatAdminRights interface {
	tl.Object
	_ChatAdminRights()
}

var (
	_ ChatAdminRights = (*ChatAdminRightsPredict)(nil)
)

type ChatAdminRightsPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ChangeInfo     bool     `tl:",omitempty:flags:0,implicit"`
	PostMessages   bool     `tl:",omitempty:flags:1,implicit"`
	EditMessages   bool     `tl:",omitempty:flags:2,implicit"`
	DeleteMessages bool     `tl:",omitempty:flags:3,implicit"`
	BanUsers       bool     `tl:",omitempty:flags:4,implicit"`
	InviteUsers    bool     `tl:",omitempty:flags:5,implicit"`
	PinMessages    bool     `tl:",omitempty:flags:7,implicit"`
	AddAdmins      bool     `tl:",omitempty:flags:9,implicit"`
	Anonymous      bool     `tl:",omitempty:flags:10,implicit"`
	ManageCall     bool     `tl:",omitempty:flags:11,implicit"`
	Other          bool     `tl:",omitempty:flags:12,implicit"`
	ManageTopics   bool     `tl:",omitempty:flags:13,implicit"`
	PostStories    bool     `tl:",omitempty:flags:14,implicit"`
	EditStories    bool     `tl:",omitempty:flags:15,implicit"`
	DeleteStories  bool     `tl:",omitempty:flags:16,implicit"`
}

func (*ChatAdminRightsPredict) CRC() uint32 {
	return 0x5fb224d5
}
func (*ChatAdminRightsPredict) _ChatAdminRights() {}

type ChatAdminWithInvites interface {
	tl.Object
	_ChatAdminWithInvites()
}

var (
	_ ChatAdminWithInvites = (*ChatAdminWithInvitesPredict)(nil)
)

type ChatAdminWithInvitesPredict struct {
	AdminID             int64
	InvitesCount        int32
	RevokedInvitesCount int32
}

func (*ChatAdminWithInvitesPredict) CRC() uint32 {
	return 0xf2ecef23
}
func (*ChatAdminWithInvitesPredict) _ChatAdminWithInvites() {}

type ChatBannedRights interface {
	tl.Object
	_ChatBannedRights()
}

var (
	_ ChatBannedRights = (*ChatBannedRightsPredict)(nil)
)

type ChatBannedRightsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ViewMessages    bool     `tl:",omitempty:flags:0,implicit"`
	SendMessages    bool     `tl:",omitempty:flags:1,implicit"`
	SendMedia       bool     `tl:",omitempty:flags:2,implicit"`
	SendStickers    bool     `tl:",omitempty:flags:3,implicit"`
	SendGifs        bool     `tl:",omitempty:flags:4,implicit"`
	SendGames       bool     `tl:",omitempty:flags:5,implicit"`
	SendInline      bool     `tl:",omitempty:flags:6,implicit"`
	EmbedLinks      bool     `tl:",omitempty:flags:7,implicit"`
	SendPolls       bool     `tl:",omitempty:flags:8,implicit"`
	ChangeInfo      bool     `tl:",omitempty:flags:10,implicit"`
	InviteUsers     bool     `tl:",omitempty:flags:15,implicit"`
	PinMessages     bool     `tl:",omitempty:flags:17,implicit"`
	ManageTopics    bool     `tl:",omitempty:flags:18,implicit"`
	SendPhotos      bool     `tl:",omitempty:flags:19,implicit"`
	SendVideos      bool     `tl:",omitempty:flags:20,implicit"`
	SendRoundvideos bool     `tl:",omitempty:flags:21,implicit"`
	SendAudios      bool     `tl:",omitempty:flags:22,implicit"`
	SendVoices      bool     `tl:",omitempty:flags:23,implicit"`
	SendDocs        bool     `tl:",omitempty:flags:24,implicit"`
	SendPlain       bool     `tl:",omitempty:flags:25,implicit"`
	UntilDate       int32
}

func (*ChatBannedRightsPredict) CRC() uint32 {
	return 0x9f120418
}
func (*ChatBannedRightsPredict) _ChatBannedRights() {}

type ChatFull interface {
	tl.Object
	_ChatFull()
}

var (
	_ ChatFull = (*ChatFullPredict)(nil)
	_ ChatFull = (*ChannelFullPredict)(nil)
)

type ChatFullPredict struct {
	_                      struct{} `tl:"flags,bitflag"`
	CanSetUsername         bool     `tl:",omitempty:flags:7,implicit"`
	HasScheduled           bool     `tl:",omitempty:flags:8,implicit"`
	TranslationsDisabled   bool     `tl:",omitempty:flags:19,implicit"`
	ID                     int64
	About                  string
	Participants           ChatParticipants
	ChatPhoto              Photo `tl:",omitempty:flags:2"`
	NotifySettings         PeerNotifySettings
	ExportedInvite         ExportedChatInvite `tl:",omitempty:flags:13"`
	BotInfo                []BotInfo          `tl:",omitempty:flags:3"`
	PinnedMsgID            *int32             `tl:",omitempty:flags:6"`
	FolderID               *int32             `tl:",omitempty:flags:11"`
	Call                   InputGroupCall     `tl:",omitempty:flags:12"`
	TTLPeriod              *int32             `tl:",omitempty:flags:14"`
	GroupcallDefaultJoinAs Peer               `tl:",omitempty:flags:15"`
	ThemeEmoticon          *string            `tl:",omitempty:flags:16"`
	RequestsPending        *int32             `tl:",omitempty:flags:17"`
	RecentRequesters       []int64            `tl:",omitempty:flags:17"`
	AvailableReactions     ChatReactions      `tl:",omitempty:flags:18"`
	ReactionsLimit         *int32             `tl:",omitempty:flags:20"`
}

func (*ChatFullPredict) CRC() uint32 {
	return 0x2633421b
}
func (*ChatFullPredict) _ChatFull() {}

type ChannelFullPredict struct {
	_                      struct{} `tl:"flags,bitflag"`
	CanViewParticipants    bool     `tl:",omitempty:flags:3,implicit"`
	CanSetUsername         bool     `tl:",omitempty:flags:6,implicit"`
	CanSetStickers         bool     `tl:",omitempty:flags:7,implicit"`
	HiddenPrehistory       bool     `tl:",omitempty:flags:10,implicit"`
	CanSetLocation         bool     `tl:",omitempty:flags:16,implicit"`
	HasScheduled           bool     `tl:",omitempty:flags:19,implicit"`
	CanViewStats           bool     `tl:",omitempty:flags:20,implicit"`
	Blocked                bool     `tl:",omitempty:flags:22,implicit"`
	_                      struct{} `tl:"flags2,bitflag"`
	CanDeleteChannel       bool     `tl:",omitempty:flags2:0,implicit"`
	Antispam               bool     `tl:",omitempty:flags2:1,implicit"`
	ParticipantsHidden     bool     `tl:",omitempty:flags2:2,implicit"`
	TranslationsDisabled   bool     `tl:",omitempty:flags2:3,implicit"`
	StoriesPinnedAvailable bool     `tl:",omitempty:flags2:5,implicit"`
	ViewForumAsMessages    bool     `tl:",omitempty:flags2:6,implicit"`
	RestrictedSponsored    bool     `tl:",omitempty:flags2:11,implicit"`
	CanViewRevenue         bool     `tl:",omitempty:flags2:12,implicit"`
	PaidMediaAllowed       bool     `tl:",omitempty:flags2:14,implicit"`
	CanViewStarsRevenue    bool     `tl:",omitempty:flags2:15,implicit"`
	ID                     int64
	About                  string
	ParticipantsCount      *int32 `tl:",omitempty:flags:0"`
	AdminsCount            *int32 `tl:",omitempty:flags:1"`
	KickedCount            *int32 `tl:",omitempty:flags:2"`
	BannedCount            *int32 `tl:",omitempty:flags:2"`
	OnlineCount            *int32 `tl:",omitempty:flags:13"`
	ReadInboxMaxID         int32
	ReadOutboxMaxID        int32
	UnreadCount            int32
	ChatPhoto              Photo
	NotifySettings         PeerNotifySettings
	ExportedInvite         ExportedChatInvite `tl:",omitempty:flags:23"`
	BotInfo                []BotInfo
	MigratedFromChatID     *int64          `tl:",omitempty:flags:4"`
	MigratedFromMaxID      *int32          `tl:",omitempty:flags:4"`
	PinnedMsgID            *int32          `tl:",omitempty:flags:5"`
	Stickerset             StickerSet      `tl:",omitempty:flags:8"`
	AvailableMinID         *int32          `tl:",omitempty:flags:9"`
	FolderID               *int32          `tl:",omitempty:flags:11"`
	LinkedChatID           *int64          `tl:",omitempty:flags:14"`
	Location               ChannelLocation `tl:",omitempty:flags:15"`
	SlowmodeSeconds        *int32          `tl:",omitempty:flags:17"`
	SlowmodeNextSendDate   *int32          `tl:",omitempty:flags:18"`
	StatsDc                *int32          `tl:",omitempty:flags:12"`
	Pts                    int32
	Call                   InputGroupCall `tl:",omitempty:flags:21"`
	TTLPeriod              *int32         `tl:",omitempty:flags:24"`
	PendingSuggestions     []string       `tl:",omitempty:flags:25"`
	GroupcallDefaultJoinAs Peer           `tl:",omitempty:flags:26"`
	ThemeEmoticon          *string        `tl:",omitempty:flags:27"`
	RequestsPending        *int32         `tl:",omitempty:flags:28"`
	RecentRequesters       []int64        `tl:",omitempty:flags:28"`
	DefaultSendAs          Peer           `tl:",omitempty:flags:29"`
	AvailableReactions     ChatReactions  `tl:",omitempty:flags:30"`
	ReactionsLimit         *int32         `tl:",omitempty:flags2:13"`
	Stories                PeerStories    `tl:",omitempty:flags2:4"`
	Wallpaper              WallPaper      `tl:",omitempty:flags2:7"`
	BoostsApplied          *int32         `tl:",omitempty:flags2:8"`
	BoostsUnrestrict       *int32         `tl:",omitempty:flags2:9"`
	Emojiset               StickerSet     `tl:",omitempty:flags2:10"`
}

func (*ChannelFullPredict) CRC() uint32 {
	return 0xbbab348d
}
func (*ChannelFullPredict) _ChatFull() {}

type ChatInvite interface {
	tl.Object
	_ChatInvite()
}

var (
	_ ChatInvite = (*ChatInviteAlreadyPredict)(nil)
	_ ChatInvite = (*ChatInvitePredict)(nil)
	_ ChatInvite = (*ChatInvitePeekPredict)(nil)
)

type ChatInviteAlreadyPredict struct {
	Chat Chat
}

func (*ChatInviteAlreadyPredict) CRC() uint32 {
	return 0x5a686d7c
}
func (*ChatInviteAlreadyPredict) _ChatInvite() {}

type ChatInvitePredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Channel           bool     `tl:",omitempty:flags:0,implicit"`
	Broadcast         bool     `tl:",omitempty:flags:1,implicit"`
	Public            bool     `tl:",omitempty:flags:2,implicit"`
	Megagroup         bool     `tl:",omitempty:flags:3,implicit"`
	RequestNeeded     bool     `tl:",omitempty:flags:6,implicit"`
	Verified          bool     `tl:",omitempty:flags:7,implicit"`
	Scam              bool     `tl:",omitempty:flags:8,implicit"`
	Fake              bool     `tl:",omitempty:flags:9,implicit"`
	Title             string
	About             *string `tl:",omitempty:flags:5"`
	Photo             Photo
	ParticipantsCount int32
	Participants      []User `tl:",omitempty:flags:4"`
	Color             int32
}

func (*ChatInvitePredict) CRC() uint32 {
	return 0xcde0ec40
}
func (*ChatInvitePredict) _ChatInvite() {}

type ChatInvitePeekPredict struct {
	Chat    Chat
	Expires int32
}

func (*ChatInvitePeekPredict) CRC() uint32 {
	return 0x61695cb0
}
func (*ChatInvitePeekPredict) _ChatInvite() {}

type ChatInviteImporter interface {
	tl.Object
	_ChatInviteImporter()
}

var (
	_ ChatInviteImporter = (*ChatInviteImporterPredict)(nil)
)

type ChatInviteImporterPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Requested   bool     `tl:",omitempty:flags:0,implicit"`
	ViaChatlist bool     `tl:",omitempty:flags:3,implicit"`
	UserID      int64
	Date        int32
	About       *string `tl:",omitempty:flags:2"`
	ApprovedBy  *int64  `tl:",omitempty:flags:1"`
}

func (*ChatInviteImporterPredict) CRC() uint32 {
	return 0x8c5adfd9
}
func (*ChatInviteImporterPredict) _ChatInviteImporter() {}

type ChatOnlines interface {
	tl.Object
	_ChatOnlines()
}

var (
	_ ChatOnlines = (*ChatOnlinesPredict)(nil)
)

type ChatOnlinesPredict struct {
	Onlines int32
}

func (*ChatOnlinesPredict) CRC() uint32 {
	return 0xf041e250
}
func (*ChatOnlinesPredict) _ChatOnlines() {}

type ChatParticipant interface {
	tl.Object
	_ChatParticipant()
}

var (
	_ ChatParticipant = (*ChatParticipantPredict)(nil)
	_ ChatParticipant = (*ChatParticipantCreatorPredict)(nil)
	_ ChatParticipant = (*ChatParticipantAdminPredict)(nil)
)

type ChatParticipantPredict struct {
	UserID    int64
	InviterID int64
	Date      int32
}

func (*ChatParticipantPredict) CRC() uint32 {
	return 0xc02d4007
}
func (*ChatParticipantPredict) _ChatParticipant() {}

type ChatParticipantCreatorPredict struct {
	UserID int64
}

func (*ChatParticipantCreatorPredict) CRC() uint32 {
	return 0xe46bcee4
}
func (*ChatParticipantCreatorPredict) _ChatParticipant() {}

type ChatParticipantAdminPredict struct {
	UserID    int64
	InviterID int64
	Date      int32
}

func (*ChatParticipantAdminPredict) CRC() uint32 {
	return 0xa0933f5b
}
func (*ChatParticipantAdminPredict) _ChatParticipant() {}

type ChatParticipants interface {
	tl.Object
	_ChatParticipants()
}

var (
	_ ChatParticipants = (*ChatParticipantsForbiddenPredict)(nil)
	_ ChatParticipants = (*ChatParticipantsPredict)(nil)
)

type ChatParticipantsForbiddenPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ChatID          int64
	SelfParticipant ChatParticipant `tl:",omitempty:flags:0"`
}

func (*ChatParticipantsForbiddenPredict) CRC() uint32 {
	return 0x8763d3e1
}
func (*ChatParticipantsForbiddenPredict) _ChatParticipants() {}

type ChatParticipantsPredict struct {
	ChatID       int64
	Participants []ChatParticipant
	Version      int32
}

func (*ChatParticipantsPredict) CRC() uint32 {
	return 0x3cbc93f8
}
func (*ChatParticipantsPredict) _ChatParticipants() {}

type ChatPhoto interface {
	tl.Object
	_ChatPhoto()
}

var (
	_ ChatPhoto = (*ChatPhotoEmptyPredict)(nil)
	_ ChatPhoto = (*ChatPhotoPredict)(nil)
)

type ChatPhotoEmptyPredict struct{}

func (*ChatPhotoEmptyPredict) CRC() uint32 {
	return 0x37c1011c
}
func (*ChatPhotoEmptyPredict) _ChatPhoto() {}

type ChatPhotoPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	HasVideo      bool     `tl:",omitempty:flags:0,implicit"`
	PhotoID       int64
	StrippedThumb *[]byte `tl:",omitempty:flags:1"`
	DcID          int32
}

func (*ChatPhotoPredict) CRC() uint32 {
	return 0x1c6e1c11
}
func (*ChatPhotoPredict) _ChatPhoto() {}

type ChatReactions interface {
	tl.Object
	_ChatReactions()
}

var (
	_ ChatReactions = (*ChatReactionsNonePredict)(nil)
	_ ChatReactions = (*ChatReactionsAllPredict)(nil)
	_ ChatReactions = (*ChatReactionsSomePredict)(nil)
)

type ChatReactionsNonePredict struct{}

func (*ChatReactionsNonePredict) CRC() uint32 {
	return 0xeafc32bc
}
func (*ChatReactionsNonePredict) _ChatReactions() {}

type ChatReactionsAllPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	AllowCustom bool     `tl:",omitempty:flags:0,implicit"`
}

func (*ChatReactionsAllPredict) CRC() uint32 {
	return 0x52928bca
}
func (*ChatReactionsAllPredict) _ChatReactions() {}

type ChatReactionsSomePredict struct {
	Reactions []Reaction
}

func (*ChatReactionsSomePredict) CRC() uint32 {
	return 0x661d4037
}
func (*ChatReactionsSomePredict) _ChatReactions() {}

type CodeSettings interface {
	tl.Object
	_CodeSettings()
}

var (
	_ CodeSettings = (*CodeSettingsPredict)(nil)
)

type CodeSettingsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	AllowFlashcall  bool     `tl:",omitempty:flags:0,implicit"`
	CurrentNumber   bool     `tl:",omitempty:flags:1,implicit"`
	AllowAppHash    bool     `tl:",omitempty:flags:4,implicit"`
	AllowMissedCall bool     `tl:",omitempty:flags:5,implicit"`
	AllowFirebase   bool     `tl:",omitempty:flags:7,implicit"`
	UnknownNumber   bool     `tl:",omitempty:flags:9,implicit"`
	LogoutTokens    [][]byte `tl:",omitempty:flags:6"`
	Token           *string  `tl:",omitempty:flags:8"`
	AppSandbox      *bool    `tl:",omitempty:flags:8"`
}

func (*CodeSettingsPredict) CRC() uint32 {
	return 0xad253d78
}
func (*CodeSettingsPredict) _CodeSettings() {}

type Config interface {
	tl.Object
	_Config()
}

var (
	_ Config = (*ConfigPredict)(nil)
)

type ConfigPredict struct {
	_                       struct{} `tl:"flags,bitflag"`
	DefaultP2PContacts      bool     `tl:",omitempty:flags:3,implicit"`
	PreloadFeaturedStickers bool     `tl:",omitempty:flags:4,implicit"`
	RevokePmInbox           bool     `tl:",omitempty:flags:6,implicit"`
	BlockedMode             bool     `tl:",omitempty:flags:8,implicit"`
	ForceTryIpv6            bool     `tl:",omitempty:flags:14,implicit"`
	Date                    int32
	Expires                 int32
	TestMode                bool
	ThisDc                  int32
	DcOptions               []DcOption
	DcTxtDomainName         string
	ChatSizeMax             int32
	MegagroupSizeMax        int32
	ForwardedCountMax       int32
	OnlineUpdatePeriodMs    int32
	OfflineBlurTimeoutMs    int32
	OfflineIdleTimeoutMs    int32
	OnlineCloudTimeoutMs    int32
	NotifyCloudDelayMs      int32
	NotifyDefaultDelayMs    int32
	PushChatPeriodMs        int32
	PushChatLimit           int32
	EditTimeLimit           int32
	RevokeTimeLimit         int32
	RevokePmTimeLimit       int32
	RatingEDecay            int32
	StickersRecentLimit     int32
	ChannelsReadMediaPeriod int32
	TmpSessions             *int32 `tl:",omitempty:flags:0"`
	CallReceiveTimeoutMs    int32
	CallRingTimeoutMs       int32
	CallConnectTimeoutMs    int32
	CallPacketTimeoutMs     int32
	MeURLPrefix             string
	AutoupdateURLPrefix     *string `tl:",omitempty:flags:7"`
	GifSearchUsername       *string `tl:",omitempty:flags:9"`
	VenueSearchUsername     *string `tl:",omitempty:flags:10"`
	ImgSearchUsername       *string `tl:",omitempty:flags:11"`
	StaticMapsProvider      *string `tl:",omitempty:flags:12"`
	CaptionLengthMax        int32
	MessageLengthMax        int32
	WebfileDcID             int32
	SuggestedLangCode       *string  `tl:",omitempty:flags:2"`
	LangPackVersion         *int32   `tl:",omitempty:flags:2"`
	BaseLangPackVersion     *int32   `tl:",omitempty:flags:2"`
	ReactionsDefault        Reaction `tl:",omitempty:flags:15"`
	AutologinToken          *string  `tl:",omitempty:flags:16"`
}

func (*ConfigPredict) CRC() uint32 {
	return 0xcc1a241e
}
func (*ConfigPredict) _Config() {}

type ConnectedBot interface {
	tl.Object
	_ConnectedBot()
}

var (
	_ ConnectedBot = (*ConnectedBotPredict)(nil)
)

type ConnectedBotPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	CanReply   bool     `tl:",omitempty:flags:0,implicit"`
	BotID      int64
	Recipients BusinessBotRecipients
}

func (*ConnectedBotPredict) CRC() uint32 {
	return 0xbd068601
}
func (*ConnectedBotPredict) _ConnectedBot() {}

type Contact interface {
	tl.Object
	_Contact()
}

var (
	_ Contact = (*ContactPredict)(nil)
)

type ContactPredict struct {
	UserID int64
	Mutual bool
}

func (*ContactPredict) CRC() uint32 {
	return 0x145ade0b
}
func (*ContactPredict) _Contact() {}

type ContactBirthday interface {
	tl.Object
	_ContactBirthday()
}

var (
	_ ContactBirthday = (*ContactBirthdayPredict)(nil)
)

type ContactBirthdayPredict struct {
	ContactID int64
	Birthday  Birthday
}

func (*ContactBirthdayPredict) CRC() uint32 {
	return 0x1d998733
}
func (*ContactBirthdayPredict) _ContactBirthday() {}

type ContactStatus interface {
	tl.Object
	_ContactStatus()
}

var (
	_ ContactStatus = (*ContactStatusPredict)(nil)
)

type ContactStatusPredict struct {
	UserID int64
	Status UserStatus
}

func (*ContactStatusPredict) CRC() uint32 {
	return 0x16d9703b
}
func (*ContactStatusPredict) _ContactStatus() {}

type DataJSON interface {
	tl.Object
	_DataJSON()
}

var (
	_ DataJSON = (*DataJSONPredict)(nil)
)

type DataJSONPredict struct {
	Data string
}

func (*DataJSONPredict) CRC() uint32 {
	return 0x7d748d04
}
func (*DataJSONPredict) _DataJSON() {}

type DcOption interface {
	tl.Object
	_DcOption()
}

var (
	_ DcOption = (*DcOptionPredict)(nil)
)

type DcOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Ipv6         bool     `tl:",omitempty:flags:0,implicit"`
	MediaOnly    bool     `tl:",omitempty:flags:1,implicit"`
	TcpoOnly     bool     `tl:",omitempty:flags:2,implicit"`
	Cdn          bool     `tl:",omitempty:flags:3,implicit"`
	Static       bool     `tl:",omitempty:flags:4,implicit"`
	ThisPortOnly bool     `tl:",omitempty:flags:5,implicit"`
	ID           int32
	IpAddress    string
	Port         int32
	Secret       *[]byte `tl:",omitempty:flags:10"`
}

func (*DcOptionPredict) CRC() uint32 {
	return 0x18b7a10d
}
func (*DcOptionPredict) _DcOption() {}

type DefaultHistoryTTL interface {
	tl.Object
	_DefaultHistoryTTL()
}

var (
	_ DefaultHistoryTTL = (*DefaultHistoryTTLPredict)(nil)
)

type DefaultHistoryTTLPredict struct {
	Period int32
}

func (*DefaultHistoryTTLPredict) CRC() uint32 {
	return 0x43b46b20
}
func (*DefaultHistoryTTLPredict) _DefaultHistoryTTL() {}

type Dialog interface {
	tl.Object
	_Dialog()
}

var (
	_ Dialog = (*DialogPredict)(nil)
	_ Dialog = (*DialogFolderPredict)(nil)
)

type DialogPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Pinned               bool     `tl:",omitempty:flags:2,implicit"`
	UnreadMark           bool     `tl:",omitempty:flags:3,implicit"`
	ViewForumAsMessages  bool     `tl:",omitempty:flags:6,implicit"`
	Peer                 Peer
	TopMessage           int32
	ReadInboxMaxID       int32
	ReadOutboxMaxID      int32
	UnreadCount          int32
	UnreadMentionsCount  int32
	UnreadReactionsCount int32
	NotifySettings       PeerNotifySettings
	Pts                  *int32       `tl:",omitempty:flags:0"`
	Draft                DraftMessage `tl:",omitempty:flags:1"`
	FolderID             *int32       `tl:",omitempty:flags:4"`
	TTLPeriod            *int32       `tl:",omitempty:flags:5"`
}

func (*DialogPredict) CRC() uint32 {
	return 0xd58a08c6
}
func (*DialogPredict) _Dialog() {}

type DialogFolderPredict struct {
	_                          struct{} `tl:"flags,bitflag"`
	Pinned                     bool     `tl:",omitempty:flags:2,implicit"`
	Folder                     Folder
	Peer                       Peer
	TopMessage                 int32
	UnreadMutedPeersCount      int32
	UnreadUnmutedPeersCount    int32
	UnreadMutedMessagesCount   int32
	UnreadUnmutedMessagesCount int32
}

func (*DialogFolderPredict) CRC() uint32 {
	return 0x71bd134c
}
func (*DialogFolderPredict) _Dialog() {}

type DialogFilter interface {
	tl.Object
	_DialogFilter()
}

var (
	_ DialogFilter = (*DialogFilterPredict)(nil)
	_ DialogFilter = (*DialogFilterDefaultPredict)(nil)
	_ DialogFilter = (*DialogFilterChatlistPredict)(nil)
)

type DialogFilterPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Contacts        bool     `tl:",omitempty:flags:0,implicit"`
	NonContacts     bool     `tl:",omitempty:flags:1,implicit"`
	Groups          bool     `tl:",omitempty:flags:2,implicit"`
	Broadcasts      bool     `tl:",omitempty:flags:3,implicit"`
	Bots            bool     `tl:",omitempty:flags:4,implicit"`
	ExcludeMuted    bool     `tl:",omitempty:flags:11,implicit"`
	ExcludeRead     bool     `tl:",omitempty:flags:12,implicit"`
	ExcludeArchived bool     `tl:",omitempty:flags:13,implicit"`
	ID              int32
	Title           string
	Emoticon        *string `tl:",omitempty:flags:25"`
	Color           *int32  `tl:",omitempty:flags:27"`
	PinnedPeers     []InputPeer
	IncludePeers    []InputPeer
	ExcludePeers    []InputPeer
}

func (*DialogFilterPredict) CRC() uint32 {
	return 0x5fb5523b
}
func (*DialogFilterPredict) _DialogFilter() {}

type DialogFilterDefaultPredict struct{}

func (*DialogFilterDefaultPredict) CRC() uint32 {
	return 0x363293ae
}
func (*DialogFilterDefaultPredict) _DialogFilter() {}

type DialogFilterChatlistPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	HasMyInvites bool     `tl:",omitempty:flags:26,implicit"`
	ID           int32
	Title        string
	Emoticon     *string `tl:",omitempty:flags:25"`
	Color        *int32  `tl:",omitempty:flags:27"`
	PinnedPeers  []InputPeer
	IncludePeers []InputPeer
}

func (*DialogFilterChatlistPredict) CRC() uint32 {
	return 0x9fe28ea4
}
func (*DialogFilterChatlistPredict) _DialogFilter() {}

type DialogFilterSuggested interface {
	tl.Object
	_DialogFilterSuggested()
}

var (
	_ DialogFilterSuggested = (*DialogFilterSuggestedPredict)(nil)
)

type DialogFilterSuggestedPredict struct {
	Filter      DialogFilter
	Description string
}

func (*DialogFilterSuggestedPredict) CRC() uint32 {
	return 0x77744d4a
}
func (*DialogFilterSuggestedPredict) _DialogFilterSuggested() {}

type DialogPeer interface {
	tl.Object
	_DialogPeer()
}

var (
	_ DialogPeer = (*DialogPeerPredict)(nil)
	_ DialogPeer = (*DialogPeerFolderPredict)(nil)
)

type DialogPeerPredict struct {
	Peer Peer
}

func (*DialogPeerPredict) CRC() uint32 {
	return 0xe56dbf05
}
func (*DialogPeerPredict) _DialogPeer() {}

type DialogPeerFolderPredict struct {
	FolderID int32
}

func (*DialogPeerFolderPredict) CRC() uint32 {
	return 0x514519e2
}
func (*DialogPeerFolderPredict) _DialogPeer() {}

type Document interface {
	tl.Object
	_Document()
}

var (
	_ Document = (*DocumentEmptyPredict)(nil)
	_ Document = (*DocumentPredict)(nil)
)

type DocumentEmptyPredict struct {
	ID int64
}

func (*DocumentEmptyPredict) CRC() uint32 {
	return 0x36f8c871
}
func (*DocumentEmptyPredict) _Document() {}

type DocumentPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	MimeType      string
	Size          int64
	Thumbs        []PhotoSize `tl:",omitempty:flags:0"`
	VideoThumbs   []VideoSize `tl:",omitempty:flags:1"`
	DcID          int32
	Attributes    []DocumentAttribute
}

func (*DocumentPredict) CRC() uint32 {
	return 0x8fd4c4d8
}
func (*DocumentPredict) _Document() {}

type DocumentAttribute interface {
	tl.Object
	_DocumentAttribute()
}

var (
	_ DocumentAttribute = (*DocumentAttributeImageSizePredict)(nil)
	_ DocumentAttribute = (*DocumentAttributeAnimatedPredict)(nil)
	_ DocumentAttribute = (*DocumentAttributeStickerPredict)(nil)
	_ DocumentAttribute = (*DocumentAttributeVideoPredict)(nil)
	_ DocumentAttribute = (*DocumentAttributeAudioPredict)(nil)
	_ DocumentAttribute = (*DocumentAttributeFilenamePredict)(nil)
	_ DocumentAttribute = (*DocumentAttributeHasStickersPredict)(nil)
	_ DocumentAttribute = (*DocumentAttributeCustomEmojiPredict)(nil)
)

type DocumentAttributeImageSizePredict struct {
	W int32
	H int32
}

func (*DocumentAttributeImageSizePredict) CRC() uint32 {
	return 0x6c37c15c
}
func (*DocumentAttributeImageSizePredict) _DocumentAttribute() {}

type DocumentAttributeAnimatedPredict struct{}

func (*DocumentAttributeAnimatedPredict) CRC() uint32 {
	return 0x11b58939
}
func (*DocumentAttributeAnimatedPredict) _DocumentAttribute() {}

type DocumentAttributeStickerPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Mask       bool     `tl:",omitempty:flags:1,implicit"`
	Alt        string
	Stickerset InputStickerSet
	MaskCoords MaskCoords `tl:",omitempty:flags:0"`
}

func (*DocumentAttributeStickerPredict) CRC() uint32 {
	return 0x6319d612
}
func (*DocumentAttributeStickerPredict) _DocumentAttribute() {}

type DocumentAttributeVideoPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	RoundMessage      bool     `tl:",omitempty:flags:0,implicit"`
	SupportsStreaming bool     `tl:",omitempty:flags:1,implicit"`
	Nosound           bool     `tl:",omitempty:flags:3,implicit"`
	Duration          float64
	W                 int32
	H                 int32
	PreloadPrefixSize *int32   `tl:",omitempty:flags:2"`
	VideoStartTs      *float64 `tl:",omitempty:flags:4"`
}

func (*DocumentAttributeVideoPredict) CRC() uint32 {
	return 0x17399fad
}
func (*DocumentAttributeVideoPredict) _DocumentAttribute() {}

type DocumentAttributeAudioPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Voice     bool     `tl:",omitempty:flags:10,implicit"`
	Duration  int32
	Title     *string `tl:",omitempty:flags:0"`
	Performer *string `tl:",omitempty:flags:1"`
	Waveform  *[]byte `tl:",omitempty:flags:2"`
}

func (*DocumentAttributeAudioPredict) CRC() uint32 {
	return 0x9852f9c6
}
func (*DocumentAttributeAudioPredict) _DocumentAttribute() {}

type DocumentAttributeFilenamePredict struct {
	FileName string
}

func (*DocumentAttributeFilenamePredict) CRC() uint32 {
	return 0x15590068
}
func (*DocumentAttributeFilenamePredict) _DocumentAttribute() {}

type DocumentAttributeHasStickersPredict struct{}

func (*DocumentAttributeHasStickersPredict) CRC() uint32 {
	return 0x9801d2f7
}
func (*DocumentAttributeHasStickersPredict) _DocumentAttribute() {}

type DocumentAttributeCustomEmojiPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Free       bool     `tl:",omitempty:flags:0,implicit"`
	TextColor  bool     `tl:",omitempty:flags:1,implicit"`
	Alt        string
	Stickerset InputStickerSet
}

func (*DocumentAttributeCustomEmojiPredict) CRC() uint32 {
	return 0xfd149899
}
func (*DocumentAttributeCustomEmojiPredict) _DocumentAttribute() {}

type DraftMessage interface {
	tl.Object
	_DraftMessage()
}

var (
	_ DraftMessage = (*DraftMessageEmptyPredict)(nil)
	_ DraftMessage = (*DraftMessagePredict)(nil)
)

type DraftMessageEmptyPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Date *int32   `tl:",omitempty:flags:0"`
}

func (*DraftMessageEmptyPredict) CRC() uint32 {
	return 0x1b0c841a
}
func (*DraftMessageEmptyPredict) _DraftMessage() {}

type DraftMessagePredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	NoWebpage   bool         `tl:",omitempty:flags:1,implicit"`
	InvertMedia bool         `tl:",omitempty:flags:6,implicit"`
	ReplyTo     InputReplyTo `tl:",omitempty:flags:4"`
	Message     string
	Entities    []MessageEntity `tl:",omitempty:flags:3"`
	Media       InputMedia      `tl:",omitempty:flags:5"`
	Date        int32
	Effect      *int64 `tl:",omitempty:flags:7"`
}

func (*DraftMessagePredict) CRC() uint32 {
	return 0x2d65321f
}
func (*DraftMessagePredict) _DraftMessage() {}

type EmailVerification interface {
	tl.Object
	_EmailVerification()
}

var (
	_ EmailVerification = (*EmailVerificationCodePredict)(nil)
	_ EmailVerification = (*EmailVerificationGooglePredict)(nil)
	_ EmailVerification = (*EmailVerificationApplePredict)(nil)
)

type EmailVerificationCodePredict struct {
	Code string
}

func (*EmailVerificationCodePredict) CRC() uint32 {
	return 0x922e55a9
}
func (*EmailVerificationCodePredict) _EmailVerification() {}

type EmailVerificationGooglePredict struct {
	Token string
}

func (*EmailVerificationGooglePredict) CRC() uint32 {
	return 0xdb909ec2
}
func (*EmailVerificationGooglePredict) _EmailVerification() {}

type EmailVerificationApplePredict struct {
	Token string
}

func (*EmailVerificationApplePredict) CRC() uint32 {
	return 0x96d074fd
}
func (*EmailVerificationApplePredict) _EmailVerification() {}

type EmailVerifyPurpose interface {
	tl.Object
	_EmailVerifyPurpose()
}

var (
	_ EmailVerifyPurpose = (*EmailVerifyPurposeLoginSetupPredict)(nil)
	_ EmailVerifyPurpose = (*EmailVerifyPurposeLoginChangePredict)(nil)
	_ EmailVerifyPurpose = (*EmailVerifyPurposePassportPredict)(nil)
)

type EmailVerifyPurposeLoginSetupPredict struct {
	PhoneNumber   string
	PhoneCodeHash string
}

func (*EmailVerifyPurposeLoginSetupPredict) CRC() uint32 {
	return 0x4345be73
}
func (*EmailVerifyPurposeLoginSetupPredict) _EmailVerifyPurpose() {}

type EmailVerifyPurposeLoginChangePredict struct{}

func (*EmailVerifyPurposeLoginChangePredict) CRC() uint32 {
	return 0x527d22eb
}
func (*EmailVerifyPurposeLoginChangePredict) _EmailVerifyPurpose() {}

type EmailVerifyPurposePassportPredict struct{}

func (*EmailVerifyPurposePassportPredict) CRC() uint32 {
	return 0xbbf51685
}
func (*EmailVerifyPurposePassportPredict) _EmailVerifyPurpose() {}

type EmojiGroup interface {
	tl.Object
	_EmojiGroup()
}

var (
	_ EmojiGroup = (*EmojiGroupPredict)(nil)
	_ EmojiGroup = (*EmojiGroupGreetingPredict)(nil)
	_ EmojiGroup = (*EmojiGroupPremiumPredict)(nil)
)

type EmojiGroupPredict struct {
	Title       string
	IconEmojiID int64
	Emoticons   []string
}

func (*EmojiGroupPredict) CRC() uint32 {
	return 0x7a9abda9
}
func (*EmojiGroupPredict) _EmojiGroup() {}

type EmojiGroupGreetingPredict struct {
	Title       string
	IconEmojiID int64
	Emoticons   []string
}

func (*EmojiGroupGreetingPredict) CRC() uint32 {
	return 0x80d26cc7
}
func (*EmojiGroupGreetingPredict) _EmojiGroup() {}

type EmojiGroupPremiumPredict struct {
	Title       string
	IconEmojiID int64
}

func (*EmojiGroupPremiumPredict) CRC() uint32 {
	return 0x93bcf34
}
func (*EmojiGroupPremiumPredict) _EmojiGroup() {}

type EmojiKeyword interface {
	tl.Object
	_EmojiKeyword()
}

var (
	_ EmojiKeyword = (*EmojiKeywordPredict)(nil)
	_ EmojiKeyword = (*EmojiKeywordDeletedPredict)(nil)
)

type EmojiKeywordPredict struct {
	Keyword   string
	Emoticons []string
}

func (*EmojiKeywordPredict) CRC() uint32 {
	return 0xd5b3b9f9
}
func (*EmojiKeywordPredict) _EmojiKeyword() {}

type EmojiKeywordDeletedPredict struct {
	Keyword   string
	Emoticons []string
}

func (*EmojiKeywordDeletedPredict) CRC() uint32 {
	return 0x236df622
}
func (*EmojiKeywordDeletedPredict) _EmojiKeyword() {}

type EmojiKeywordsDifference interface {
	tl.Object
	_EmojiKeywordsDifference()
}

var (
	_ EmojiKeywordsDifference = (*EmojiKeywordsDifferencePredict)(nil)
)

type EmojiKeywordsDifferencePredict struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Keywords    []EmojiKeyword
}

func (*EmojiKeywordsDifferencePredict) CRC() uint32 {
	return 0x5cc761bd
}
func (*EmojiKeywordsDifferencePredict) _EmojiKeywordsDifference() {}

type EmojiLanguage interface {
	tl.Object
	_EmojiLanguage()
}

var (
	_ EmojiLanguage = (*EmojiLanguagePredict)(nil)
)

type EmojiLanguagePredict struct {
	LangCode string
}

func (*EmojiLanguagePredict) CRC() uint32 {
	return 0xb3fb5361
}
func (*EmojiLanguagePredict) _EmojiLanguage() {}

type EmojiList interface {
	tl.Object
	_EmojiList()
}

var (
	_ EmojiList = (*EmojiListNotModifiedPredict)(nil)
	_ EmojiList = (*EmojiListPredict)(nil)
)

type EmojiListNotModifiedPredict struct{}

func (*EmojiListNotModifiedPredict) CRC() uint32 {
	return 0x481eadfa
}
func (*EmojiListNotModifiedPredict) _EmojiList() {}

type EmojiListPredict struct {
	Hash       int64
	DocumentID []int64
}

func (*EmojiListPredict) CRC() uint32 {
	return 0x7a1e11d1
}
func (*EmojiListPredict) _EmojiList() {}

type EmojiStatus interface {
	tl.Object
	_EmojiStatus()
}

var (
	_ EmojiStatus = (*EmojiStatusEmptyPredict)(nil)
	_ EmojiStatus = (*EmojiStatusPredict)(nil)
	_ EmojiStatus = (*EmojiStatusUntilPredict)(nil)
)

type EmojiStatusEmptyPredict struct{}

func (*EmojiStatusEmptyPredict) CRC() uint32 {
	return 0x2de11aae
}
func (*EmojiStatusEmptyPredict) _EmojiStatus() {}

type EmojiStatusPredict struct {
	DocumentID int64
}

func (*EmojiStatusPredict) CRC() uint32 {
	return 0x929b619d
}
func (*EmojiStatusPredict) _EmojiStatus() {}

type EmojiStatusUntilPredict struct {
	DocumentID int64
	Until      int32
}

func (*EmojiStatusUntilPredict) CRC() uint32 {
	return 0xfa30a8c7
}
func (*EmojiStatusUntilPredict) _EmojiStatus() {}

type EmojiURL interface {
	tl.Object
	_EmojiURL()
}

var (
	_ EmojiURL = (*EmojiURLPredict)(nil)
)

type EmojiURLPredict struct {
	URL string
}

func (*EmojiURLPredict) CRC() uint32 {
	return 0xa575739d
}
func (*EmojiURLPredict) _EmojiURL() {}

type EncryptedChat interface {
	tl.Object
	_EncryptedChat()
}

var (
	_ EncryptedChat = (*EncryptedChatEmptyPredict)(nil)
	_ EncryptedChat = (*EncryptedChatWaitingPredict)(nil)
	_ EncryptedChat = (*EncryptedChatRequestedPredict)(nil)
	_ EncryptedChat = (*EncryptedChatPredict)(nil)
	_ EncryptedChat = (*EncryptedChatDiscardedPredict)(nil)
)

type EncryptedChatEmptyPredict struct {
	ID int32
}

func (*EncryptedChatEmptyPredict) CRC() uint32 {
	return 0xab7ec0a0
}
func (*EncryptedChatEmptyPredict) _EncryptedChat() {}

type EncryptedChatWaitingPredict struct {
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
}

func (*EncryptedChatWaitingPredict) CRC() uint32 {
	return 0x66b25953
}
func (*EncryptedChatWaitingPredict) _EncryptedChat() {}

type EncryptedChatRequestedPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	FolderID      *int32   `tl:",omitempty:flags:0"`
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	GA            []byte
}

func (*EncryptedChatRequestedPredict) CRC() uint32 {
	return 0x48f1d94c
}
func (*EncryptedChatRequestedPredict) _EncryptedChat() {}

type EncryptedChatPredict struct {
	ID             int32
	AccessHash     int64
	Date           int32
	AdminID        int64
	ParticipantID  int64
	GAOrB          []byte
	KeyFingerprint int64
}

func (*EncryptedChatPredict) CRC() uint32 {
	return 0x61f0d4c7
}
func (*EncryptedChatPredict) _EncryptedChat() {}

type EncryptedChatDiscardedPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	HistoryDeleted bool     `tl:",omitempty:flags:0,implicit"`
	ID             int32
}

func (*EncryptedChatDiscardedPredict) CRC() uint32 {
	return 0x1e1c7c45
}
func (*EncryptedChatDiscardedPredict) _EncryptedChat() {}

type EncryptedFile interface {
	tl.Object
	_EncryptedFile()
}

var (
	_ EncryptedFile = (*EncryptedFileEmptyPredict)(nil)
	_ EncryptedFile = (*EncryptedFilePredict)(nil)
)

type EncryptedFileEmptyPredict struct{}

func (*EncryptedFileEmptyPredict) CRC() uint32 {
	return 0xc21f497e
}
func (*EncryptedFileEmptyPredict) _EncryptedFile() {}

type EncryptedFilePredict struct {
	ID             int64
	AccessHash     int64
	Size           int64
	DcID           int32
	KeyFingerprint int32
}

func (*EncryptedFilePredict) CRC() uint32 {
	return 0xa8008cd8
}
func (*EncryptedFilePredict) _EncryptedFile() {}

type EncryptedMessage interface {
	tl.Object
	_EncryptedMessage()
}

var (
	_ EncryptedMessage = (*EncryptedMessagePredict)(nil)
	_ EncryptedMessage = (*EncryptedMessageServicePredict)(nil)
)

type EncryptedMessagePredict struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
	File     EncryptedFile
}

func (*EncryptedMessagePredict) CRC() uint32 {
	return 0xed18c118
}
func (*EncryptedMessagePredict) _EncryptedMessage() {}

type EncryptedMessageServicePredict struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
}

func (*EncryptedMessageServicePredict) CRC() uint32 {
	return 0x23734b06
}
func (*EncryptedMessageServicePredict) _EncryptedMessage() {}

type Error interface {
	tl.Object
	_Error()
}

var (
	_ Error = (*ErrorPredict)(nil)
)

type ErrorPredict struct {
	Code int32
	Text string
}

func (*ErrorPredict) CRC() uint32 {
	return 0xc4b9f9bb
}
func (*ErrorPredict) _Error() {}

type ExportedChatInvite interface {
	tl.Object
	_ExportedChatInvite()
}

var (
	_ ExportedChatInvite = (*ChatInviteExportedPredict)(nil)
	_ ExportedChatInvite = (*ChatInvitePublicJoinRequestsPredict)(nil)
)

type ChatInviteExportedPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Revoked       bool     `tl:",omitempty:flags:0,implicit"`
	Permanent     bool     `tl:",omitempty:flags:5,implicit"`
	RequestNeeded bool     `tl:",omitempty:flags:6,implicit"`
	Link          string
	AdminID       int64
	Date          int32
	StartDate     *int32  `tl:",omitempty:flags:4"`
	ExpireDate    *int32  `tl:",omitempty:flags:1"`
	UsageLimit    *int32  `tl:",omitempty:flags:2"`
	Usage         *int32  `tl:",omitempty:flags:3"`
	Requested     *int32  `tl:",omitempty:flags:7"`
	Title         *string `tl:",omitempty:flags:8"`
}

func (*ChatInviteExportedPredict) CRC() uint32 {
	return 0xab4a819
}
func (*ChatInviteExportedPredict) _ExportedChatInvite() {}

type ChatInvitePublicJoinRequestsPredict struct{}

func (*ChatInvitePublicJoinRequestsPredict) CRC() uint32 {
	return 0xed107ab7
}
func (*ChatInvitePublicJoinRequestsPredict) _ExportedChatInvite() {}

type ExportedChatlistInvite interface {
	tl.Object
	_ExportedChatlistInvite()
}

var (
	_ ExportedChatlistInvite = (*ExportedChatlistInvitePredict)(nil)
)

type ExportedChatlistInvitePredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Title string
	URL   string
	Peers []Peer
}

func (*ExportedChatlistInvitePredict) CRC() uint32 {
	return 0xc5181ac
}
func (*ExportedChatlistInvitePredict) _ExportedChatlistInvite() {}

type ExportedContactToken interface {
	tl.Object
	_ExportedContactToken()
}

var (
	_ ExportedContactToken = (*ExportedContactTokenPredict)(nil)
)

type ExportedContactTokenPredict struct {
	URL     string
	Expires int32
}

func (*ExportedContactTokenPredict) CRC() uint32 {
	return 0x41bf109b
}
func (*ExportedContactTokenPredict) _ExportedContactToken() {}

type ExportedMessageLink interface {
	tl.Object
	_ExportedMessageLink()
}

var (
	_ ExportedMessageLink = (*ExportedMessageLinkPredict)(nil)
)

type ExportedMessageLinkPredict struct {
	Link string
	Html string
}

func (*ExportedMessageLinkPredict) CRC() uint32 {
	return 0x5dab1af4
}
func (*ExportedMessageLinkPredict) _ExportedMessageLink() {}

type ExportedStoryLink interface {
	tl.Object
	_ExportedStoryLink()
}

var (
	_ ExportedStoryLink = (*ExportedStoryLinkPredict)(nil)
)

type ExportedStoryLinkPredict struct {
	Link string
}

func (*ExportedStoryLinkPredict) CRC() uint32 {
	return 0x3fc9053b
}
func (*ExportedStoryLinkPredict) _ExportedStoryLink() {}

type FactCheck interface {
	tl.Object
	_FactCheck()
}

var (
	_ FactCheck = (*FactCheckPredict)(nil)
)

type FactCheckPredict struct {
	_         struct{}         `tl:"flags,bitflag"`
	NeedCheck bool             `tl:",omitempty:flags:0,implicit"`
	Country   *string          `tl:",omitempty:flags:1"`
	Text      TextWithEntities `tl:",omitempty:flags:1"`
	Hash      int64
}

func (*FactCheckPredict) CRC() uint32 {
	return 0xb89bfccf
}
func (*FactCheckPredict) _FactCheck() {}

type FileHash interface {
	tl.Object
	_FileHash()
}

var (
	_ FileHash = (*FileHashPredict)(nil)
)

type FileHashPredict struct {
	Offset int64
	Limit  int32
	Hash   []byte
}

func (*FileHashPredict) CRC() uint32 {
	return 0xf39b035c
}
func (*FileHashPredict) _FileHash() {}

type Folder interface {
	tl.Object
	_Folder()
}

var (
	_ Folder = (*FolderPredict)(nil)
)

type FolderPredict struct {
	_                         struct{} `tl:"flags,bitflag"`
	AutofillNewBroadcasts     bool     `tl:",omitempty:flags:0,implicit"`
	AutofillPublicGroups      bool     `tl:",omitempty:flags:1,implicit"`
	AutofillNewCorrespondents bool     `tl:",omitempty:flags:2,implicit"`
	ID                        int32
	Title                     string
	Photo                     ChatPhoto `tl:",omitempty:flags:3"`
}

func (*FolderPredict) CRC() uint32 {
	return 0xff544e65
}
func (*FolderPredict) _Folder() {}

type FolderPeer interface {
	tl.Object
	_FolderPeer()
}

var (
	_ FolderPeer = (*FolderPeerPredict)(nil)
)

type FolderPeerPredict struct {
	Peer     Peer
	FolderID int32
}

func (*FolderPeerPredict) CRC() uint32 {
	return 0xe9baa668
}
func (*FolderPeerPredict) _FolderPeer() {}

type ForumTopic interface {
	tl.Object
	_ForumTopic()
}

var (
	_ ForumTopic = (*ForumTopicDeletedPredict)(nil)
	_ ForumTopic = (*ForumTopicPredict)(nil)
)

type ForumTopicDeletedPredict struct {
	ID int32
}

func (*ForumTopicDeletedPredict) CRC() uint32 {
	return 0x23f109b
}
func (*ForumTopicDeletedPredict) _ForumTopic() {}

type ForumTopicPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	My                   bool     `tl:",omitempty:flags:1,implicit"`
	Closed               bool     `tl:",omitempty:flags:2,implicit"`
	Pinned               bool     `tl:",omitempty:flags:3,implicit"`
	Short                bool     `tl:",omitempty:flags:5,implicit"`
	Hidden               bool     `tl:",omitempty:flags:6,implicit"`
	ID                   int32
	Date                 int32
	Title                string
	IconColor            int32
	IconEmojiID          *int64 `tl:",omitempty:flags:0"`
	TopMessage           int32
	ReadInboxMaxID       int32
	ReadOutboxMaxID      int32
	UnreadCount          int32
	UnreadMentionsCount  int32
	UnreadReactionsCount int32
	FromID               Peer
	NotifySettings       PeerNotifySettings
	Draft                DraftMessage `tl:",omitempty:flags:4"`
}

func (*ForumTopicPredict) CRC() uint32 {
	return 0x71701da9
}
func (*ForumTopicPredict) _ForumTopic() {}

type FoundStory interface {
	tl.Object
	_FoundStory()
}

var (
	_ FoundStory = (*FoundStoryPredict)(nil)
)

type FoundStoryPredict struct {
	Peer  Peer
	Story StoryItem
}

func (*FoundStoryPredict) CRC() uint32 {
	return 0xe87acbc0
}
func (*FoundStoryPredict) _FoundStory() {}

type Game interface {
	tl.Object
	_Game()
}

var (
	_ Game = (*GamePredict)(nil)
)

type GamePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          int64
	AccessHash  int64
	ShortName   string
	Title       string
	Description string
	Photo       Photo
	Document    Document `tl:",omitempty:flags:0"`
}

func (*GamePredict) CRC() uint32 {
	return 0xbdf9653b
}
func (*GamePredict) _Game() {}

type GeoPoint interface {
	tl.Object
	_GeoPoint()
}

var (
	_ GeoPoint = (*GeoPointEmptyPredict)(nil)
	_ GeoPoint = (*GeoPointPredict)(nil)
)

type GeoPointEmptyPredict struct{}

func (*GeoPointEmptyPredict) CRC() uint32 {
	return 0x1117dd5f
}
func (*GeoPointEmptyPredict) _GeoPoint() {}

type GeoPointPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Long           float64
	Lat            float64
	AccessHash     int64
	AccuracyRadius *int32 `tl:",omitempty:flags:0"`
}

func (*GeoPointPredict) CRC() uint32 {
	return 0xb2a2f663
}
func (*GeoPointPredict) _GeoPoint() {}

type GeoPointAddress interface {
	tl.Object
	_GeoPointAddress()
}

var (
	_ GeoPointAddress = (*GeoPointAddressPredict)(nil)
)

type GeoPointAddressPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	CountryIso2 string
	State       *string `tl:",omitempty:flags:0"`
	City        *string `tl:",omitempty:flags:1"`
	Street      *string `tl:",omitempty:flags:2"`
}

func (*GeoPointAddressPredict) CRC() uint32 {
	return 0xde4c5d93
}
func (*GeoPointAddressPredict) _GeoPointAddress() {}

type GlobalPrivacySettings interface {
	tl.Object
	_GlobalPrivacySettings()
}

var (
	_ GlobalPrivacySettings = (*GlobalPrivacySettingsPredict)(nil)
)

type GlobalPrivacySettingsPredict struct {
	_                                struct{} `tl:"flags,bitflag"`
	ArchiveAndMuteNewNoncontactPeers bool     `tl:",omitempty:flags:0,implicit"`
	KeepArchivedUnmuted              bool     `tl:",omitempty:flags:1,implicit"`
	KeepArchivedFolders              bool     `tl:",omitempty:flags:2,implicit"`
	HideReadMarks                    bool     `tl:",omitempty:flags:3,implicit"`
	NewNoncontactPeersRequirePremium bool     `tl:",omitempty:flags:4,implicit"`
}

func (*GlobalPrivacySettingsPredict) CRC() uint32 {
	return 0x734c4ccb
}
func (*GlobalPrivacySettingsPredict) _GlobalPrivacySettings() {}

type GroupCall interface {
	tl.Object
	_GroupCall()
}

var (
	_ GroupCall = (*GroupCallDiscardedPredict)(nil)
	_ GroupCall = (*GroupCallPredict)(nil)
)

type GroupCallDiscardedPredict struct {
	ID         int64
	AccessHash int64
	Duration   int32
}

func (*GroupCallDiscardedPredict) CRC() uint32 {
	return 0x7780bcb4
}
func (*GroupCallDiscardedPredict) _GroupCall() {}

type GroupCallPredict struct {
	_                       struct{} `tl:"flags,bitflag"`
	JoinMuted               bool     `tl:",omitempty:flags:1,implicit"`
	CanChangeJoinMuted      bool     `tl:",omitempty:flags:2,implicit"`
	JoinDateAsc             bool     `tl:",omitempty:flags:6,implicit"`
	ScheduleStartSubscribed bool     `tl:",omitempty:flags:8,implicit"`
	CanStartVideo           bool     `tl:",omitempty:flags:9,implicit"`
	RecordVideoActive       bool     `tl:",omitempty:flags:11,implicit"`
	RtmpStream              bool     `tl:",omitempty:flags:12,implicit"`
	ListenersHidden         bool     `tl:",omitempty:flags:13,implicit"`
	ID                      int64
	AccessHash              int64
	ParticipantsCount       int32
	Title                   *string `tl:",omitempty:flags:3"`
	StreamDcID              *int32  `tl:",omitempty:flags:4"`
	RecordStartDate         *int32  `tl:",omitempty:flags:5"`
	ScheduleDate            *int32  `tl:",omitempty:flags:7"`
	UnmutedVideoCount       *int32  `tl:",omitempty:flags:10"`
	UnmutedVideoLimit       int32
	Version                 int32
}

func (*GroupCallPredict) CRC() uint32 {
	return 0xd597650c
}
func (*GroupCallPredict) _GroupCall() {}

type GroupCallParticipant interface {
	tl.Object
	_GroupCallParticipant()
}

var (
	_ GroupCallParticipant = (*GroupCallParticipantPredict)(nil)
)

type GroupCallParticipantPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Muted           bool     `tl:",omitempty:flags:0,implicit"`
	Left            bool     `tl:",omitempty:flags:1,implicit"`
	CanSelfUnmute   bool     `tl:",omitempty:flags:2,implicit"`
	JustJoined      bool     `tl:",omitempty:flags:4,implicit"`
	Versioned       bool     `tl:",omitempty:flags:5,implicit"`
	Min             bool     `tl:",omitempty:flags:8,implicit"`
	MutedByYou      bool     `tl:",omitempty:flags:9,implicit"`
	VolumeByAdmin   bool     `tl:",omitempty:flags:10,implicit"`
	Self            bool     `tl:",omitempty:flags:12,implicit"`
	VideoJoined     bool     `tl:",omitempty:flags:15,implicit"`
	Peer            Peer
	Date            int32
	ActiveDate      *int32 `tl:",omitempty:flags:3"`
	Source          int32
	Volume          *int32                    `tl:",omitempty:flags:7"`
	About           *string                   `tl:",omitempty:flags:11"`
	RaiseHandRating *int64                    `tl:",omitempty:flags:13"`
	Video           GroupCallParticipantVideo `tl:",omitempty:flags:6"`
	Presentation    GroupCallParticipantVideo `tl:",omitempty:flags:14"`
}

func (*GroupCallParticipantPredict) CRC() uint32 {
	return 0xeba636fe
}
func (*GroupCallParticipantPredict) _GroupCallParticipant() {}

type GroupCallParticipantVideo interface {
	tl.Object
	_GroupCallParticipantVideo()
}

var (
	_ GroupCallParticipantVideo = (*GroupCallParticipantVideoPredict)(nil)
)

type GroupCallParticipantVideoPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Paused       bool     `tl:",omitempty:flags:0,implicit"`
	Endpoint     string
	SourceGroups []GroupCallParticipantVideoSourceGroup
	AudioSource  *int32 `tl:",omitempty:flags:1"`
}

func (*GroupCallParticipantVideoPredict) CRC() uint32 {
	return 0x67753ac8
}
func (*GroupCallParticipantVideoPredict) _GroupCallParticipantVideo() {}

type GroupCallParticipantVideoSourceGroup interface {
	tl.Object
	_GroupCallParticipantVideoSourceGroup()
}

var (
	_ GroupCallParticipantVideoSourceGroup = (*GroupCallParticipantVideoSourceGroupPredict)(nil)
)

type GroupCallParticipantVideoSourceGroupPredict struct {
	Semantics string
	Sources   []int32
}

func (*GroupCallParticipantVideoSourceGroupPredict) CRC() uint32 {
	return 0xdcb118b7
}
func (*GroupCallParticipantVideoSourceGroupPredict) _GroupCallParticipantVideoSourceGroup() {}

type GroupCallStreamChannel interface {
	tl.Object
	_GroupCallStreamChannel()
}

var (
	_ GroupCallStreamChannel = (*GroupCallStreamChannelPredict)(nil)
)

type GroupCallStreamChannelPredict struct {
	Channel         int32
	Scale           int32
	LastTimestampMs int64
}

func (*GroupCallStreamChannelPredict) CRC() uint32 {
	return 0x80eb48af
}
func (*GroupCallStreamChannelPredict) _GroupCallStreamChannel() {}

type HighScore interface {
	tl.Object
	_HighScore()
}

var (
	_ HighScore = (*HighScorePredict)(nil)
)

type HighScorePredict struct {
	Pos    int32
	UserID int64
	Score  int32
}

func (*HighScorePredict) CRC() uint32 {
	return 0x73a379eb
}
func (*HighScorePredict) _HighScore() {}

type ImportedContact interface {
	tl.Object
	_ImportedContact()
}

var (
	_ ImportedContact = (*ImportedContactPredict)(nil)
)

type ImportedContactPredict struct {
	UserID   int64
	ClientID int64
}

func (*ImportedContactPredict) CRC() uint32 {
	return 0xc13e3c50
}
func (*ImportedContactPredict) _ImportedContact() {}

type InlineBotSwitchPm interface {
	tl.Object
	_InlineBotSwitchPm()
}

var (
	_ InlineBotSwitchPm = (*InlineBotSwitchPmPredict)(nil)
)

type InlineBotSwitchPmPredict struct {
	Text       string
	StartParam string
}

func (*InlineBotSwitchPmPredict) CRC() uint32 {
	return 0x3c20629f
}
func (*InlineBotSwitchPmPredict) _InlineBotSwitchPm() {}

type InlineBotWebView interface {
	tl.Object
	_InlineBotWebView()
}

var (
	_ InlineBotWebView = (*InlineBotWebViewPredict)(nil)
)

type InlineBotWebViewPredict struct {
	Text string
	URL  string
}

func (*InlineBotWebViewPredict) CRC() uint32 {
	return 0xb57295d5
}
func (*InlineBotWebViewPredict) _InlineBotWebView() {}

type InputAppEvent interface {
	tl.Object
	_InputAppEvent()
}

var (
	_ InputAppEvent = (*InputAppEventPredict)(nil)
)

type InputAppEventPredict struct {
	Time float64
	Type string
	Peer int64
	Data JSONValue
}

func (*InputAppEventPredict) CRC() uint32 {
	return 0x1d1b1245
}
func (*InputAppEventPredict) _InputAppEvent() {}

type InputBotApp interface {
	tl.Object
	_InputBotApp()
}

var (
	_ InputBotApp = (*InputBotAppIDPredict)(nil)
	_ InputBotApp = (*InputBotAppShortNamePredict)(nil)
)

type InputBotAppIDPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputBotAppIDPredict) CRC() uint32 {
	return 0xa920bd7a
}
func (*InputBotAppIDPredict) _InputBotApp() {}

type InputBotAppShortNamePredict struct {
	BotID     InputUser
	ShortName string
}

func (*InputBotAppShortNamePredict) CRC() uint32 {
	return 0x908c0407
}
func (*InputBotAppShortNamePredict) _InputBotApp() {}

type InputBotInlineMessage interface {
	tl.Object
	_InputBotInlineMessage()
}

var (
	_ InputBotInlineMessage = (*InputBotInlineMessageMediaAutoPredict)(nil)
	_ InputBotInlineMessage = (*InputBotInlineMessageTextPredict)(nil)
	_ InputBotInlineMessage = (*InputBotInlineMessageMediaGeoPredict)(nil)
	_ InputBotInlineMessage = (*InputBotInlineMessageMediaVenuePredict)(nil)
	_ InputBotInlineMessage = (*InputBotInlineMessageMediaContactPredict)(nil)
	_ InputBotInlineMessage = (*InputBotInlineMessageGamePredict)(nil)
	_ InputBotInlineMessage = (*InputBotInlineMessageMediaInvoicePredict)(nil)
	_ InputBotInlineMessage = (*InputBotInlineMessageMediaWebPagePredict)(nil)
)

type InputBotInlineMessageMediaAutoPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []MessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaAutoPredict) CRC() uint32 {
	return 0x3380c786
}
func (*InputBotInlineMessageMediaAutoPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageTextPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	NoWebpage   bool     `tl:",omitempty:flags:0,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:3,implicit"`
	Message     string
	Entities    []MessageEntity `tl:",omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageTextPredict) CRC() uint32 {
	return 0x3dcd7a87
}
func (*InputBotInlineMessageTextPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaGeoPredict struct {
	_                           struct{} `tl:"flags,bitflag"`
	GeoPoint                    InputGeoPoint
	Heading                     *int32      `tl:",omitempty:flags:0"`
	Period                      *int32      `tl:",omitempty:flags:1"`
	ProximityNotificationRadius *int32      `tl:",omitempty:flags:3"`
	ReplyMarkup                 ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaGeoPredict) CRC() uint32 {
	return 0x96929a85
}
func (*InputBotInlineMessageMediaGeoPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaVenuePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	GeoPoint    InputGeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaVenuePredict) CRC() uint32 {
	return 0x417bbf11
}
func (*InputBotInlineMessageMediaVenuePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaContactPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaContactPredict) CRC() uint32 {
	return 0xa6edbffd
}
func (*InputBotInlineMessageMediaContactPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageGamePredict struct {
	_           struct{}    `tl:"flags,bitflag"`
	ReplyMarkup ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageGamePredict) CRC() uint32 {
	return 0x4b425864
}
func (*InputBotInlineMessageGamePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaInvoicePredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Title        string
	Description  string
	Photo        InputWebDocument `tl:",omitempty:flags:0"`
	Invoice      Invoice
	Payload      []byte
	Provider     string
	ProviderData DataJSON
	ReplyMarkup  ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaInvoicePredict) CRC() uint32 {
	return 0xd7e78225
}
func (*InputBotInlineMessageMediaInvoicePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaWebPagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	InvertMedia     bool     `tl:",omitempty:flags:3,implicit"`
	ForceLargeMedia bool     `tl:",omitempty:flags:4,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:5,implicit"`
	Optional        bool     `tl:",omitempty:flags:6,implicit"`
	Message         string
	Entities        []MessageEntity `tl:",omitempty:flags:1"`
	URL             string
	ReplyMarkup     ReplyMarkup `tl:",omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaWebPagePredict) CRC() uint32 {
	return 0xbddcc510
}
func (*InputBotInlineMessageMediaWebPagePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageID interface {
	tl.Object
	_InputBotInlineMessageID()
}

var (
	_ InputBotInlineMessageID = (*InputBotInlineMessageIDPredict)(nil)
	_ InputBotInlineMessageID = (*InputBotInlineMessageID64Predict)(nil)
)

type InputBotInlineMessageIDPredict struct {
	DcID       int32
	ID         int64
	AccessHash int64
}

func (*InputBotInlineMessageIDPredict) CRC() uint32 {
	return 0x890c3d89
}
func (*InputBotInlineMessageIDPredict) _InputBotInlineMessageID() {}

type InputBotInlineMessageID64Predict struct {
	DcID       int32
	OwnerID    int64
	ID         int32
	AccessHash int64
}

func (*InputBotInlineMessageID64Predict) CRC() uint32 {
	return 0xb6d915d7
}
func (*InputBotInlineMessageID64Predict) _InputBotInlineMessageID() {}

type InputBotInlineResult interface {
	tl.Object
	_InputBotInlineResult()
}

var (
	_ InputBotInlineResult = (*InputBotInlineResultPredict)(nil)
	_ InputBotInlineResult = (*InputBotInlineResultPhotoPredict)(nil)
	_ InputBotInlineResult = (*InputBotInlineResultDocumentPredict)(nil)
	_ InputBotInlineResult = (*InputBotInlineResultGamePredict)(nil)
)

type InputBotInlineResultPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Title       *string          `tl:",omitempty:flags:1"`
	Description *string          `tl:",omitempty:flags:2"`
	URL         *string          `tl:",omitempty:flags:3"`
	Thumb       InputWebDocument `tl:",omitempty:flags:4"`
	Content     InputWebDocument `tl:",omitempty:flags:5"`
	SendMessage InputBotInlineMessage
}

func (*InputBotInlineResultPredict) CRC() uint32 {
	return 0x88bf9319
}
func (*InputBotInlineResultPredict) _InputBotInlineResult() {}

type InputBotInlineResultPhotoPredict struct {
	ID          string
	Type        string
	Photo       InputPhoto
	SendMessage InputBotInlineMessage
}

func (*InputBotInlineResultPhotoPredict) CRC() uint32 {
	return 0xa8d864a7
}
func (*InputBotInlineResultPhotoPredict) _InputBotInlineResult() {}

type InputBotInlineResultDocumentPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          string
	Type        string
	Title       *string `tl:",omitempty:flags:1"`
	Description *string `tl:",omitempty:flags:2"`
	Document    InputDocument
	SendMessage InputBotInlineMessage
}

func (*InputBotInlineResultDocumentPredict) CRC() uint32 {
	return 0xfff8fdc4
}
func (*InputBotInlineResultDocumentPredict) _InputBotInlineResult() {}

type InputBotInlineResultGamePredict struct {
	ID          string
	ShortName   string
	SendMessage InputBotInlineMessage
}

func (*InputBotInlineResultGamePredict) CRC() uint32 {
	return 0x4fa417f2
}
func (*InputBotInlineResultGamePredict) _InputBotInlineResult() {}

type InputBusinessAwayMessage interface {
	tl.Object
	_InputBusinessAwayMessage()
}

var (
	_ InputBusinessAwayMessage = (*InputBusinessAwayMessagePredict)(nil)
)

type InputBusinessAwayMessagePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	OfflineOnly bool     `tl:",omitempty:flags:0,implicit"`
	ShortcutID  int32
	Schedule    BusinessAwayMessageSchedule
	Recipients  InputBusinessRecipients
}

func (*InputBusinessAwayMessagePredict) CRC() uint32 {
	return 0x832175e0
}
func (*InputBusinessAwayMessagePredict) _InputBusinessAwayMessage() {}

type InputBusinessBotRecipients interface {
	tl.Object
	_InputBusinessBotRecipients()
}

var (
	_ InputBusinessBotRecipients = (*InputBusinessBotRecipientsPredict)(nil)
)

type InputBusinessBotRecipientsPredict struct {
	_               struct{}    `tl:"flags,bitflag"`
	ExistingChats   bool        `tl:",omitempty:flags:0,implicit"`
	NewChats        bool        `tl:",omitempty:flags:1,implicit"`
	Contacts        bool        `tl:",omitempty:flags:2,implicit"`
	NonContacts     bool        `tl:",omitempty:flags:3,implicit"`
	ExcludeSelected bool        `tl:",omitempty:flags:5,implicit"`
	Users           []InputUser `tl:",omitempty:flags:4"`
	ExcludeUsers    []InputUser `tl:",omitempty:flags:6"`
}

func (*InputBusinessBotRecipientsPredict) CRC() uint32 {
	return 0xc4e5921e
}
func (*InputBusinessBotRecipientsPredict) _InputBusinessBotRecipients() {}

type InputBusinessChatLink interface {
	tl.Object
	_InputBusinessChatLink()
}

var (
	_ InputBusinessChatLink = (*InputBusinessChatLinkPredict)(nil)
)

type InputBusinessChatLinkPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Message  string
	Entities []MessageEntity `tl:",omitempty:flags:0"`
	Title    *string         `tl:",omitempty:flags:1"`
}

func (*InputBusinessChatLinkPredict) CRC() uint32 {
	return 0x11679fa7
}
func (*InputBusinessChatLinkPredict) _InputBusinessChatLink() {}

type InputBusinessGreetingMessage interface {
	tl.Object
	_InputBusinessGreetingMessage()
}

var (
	_ InputBusinessGreetingMessage = (*InputBusinessGreetingMessagePredict)(nil)
)

type InputBusinessGreetingMessagePredict struct {
	ShortcutID     int32
	Recipients     InputBusinessRecipients
	NoActivityDays int32
}

func (*InputBusinessGreetingMessagePredict) CRC() uint32 {
	return 0x194cb3b
}
func (*InputBusinessGreetingMessagePredict) _InputBusinessGreetingMessage() {}

type InputBusinessIntro interface {
	tl.Object
	_InputBusinessIntro()
}

var (
	_ InputBusinessIntro = (*InputBusinessIntroPredict)(nil)
)

type InputBusinessIntroPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       string
	Description string
	Sticker     InputDocument `tl:",omitempty:flags:0"`
}

func (*InputBusinessIntroPredict) CRC() uint32 {
	return 0x9c469cd
}
func (*InputBusinessIntroPredict) _InputBusinessIntro() {}

type InputBusinessRecipients interface {
	tl.Object
	_InputBusinessRecipients()
}

var (
	_ InputBusinessRecipients = (*InputBusinessRecipientsPredict)(nil)
)

type InputBusinessRecipientsPredict struct {
	_               struct{}    `tl:"flags,bitflag"`
	ExistingChats   bool        `tl:",omitempty:flags:0,implicit"`
	NewChats        bool        `tl:",omitempty:flags:1,implicit"`
	Contacts        bool        `tl:",omitempty:flags:2,implicit"`
	NonContacts     bool        `tl:",omitempty:flags:3,implicit"`
	ExcludeSelected bool        `tl:",omitempty:flags:5,implicit"`
	Users           []InputUser `tl:",omitempty:flags:4"`
}

func (*InputBusinessRecipientsPredict) CRC() uint32 {
	return 0x6f8b32aa
}
func (*InputBusinessRecipientsPredict) _InputBusinessRecipients() {}

type InputChannel interface {
	tl.Object
	_InputChannel()
}

var (
	_ InputChannel = (*InputChannelEmptyPredict)(nil)
	_ InputChannel = (*InputChannelPredict)(nil)
	_ InputChannel = (*InputChannelFromMessagePredict)(nil)
)

type InputChannelEmptyPredict struct{}

func (*InputChannelEmptyPredict) CRC() uint32 {
	return 0xee8c1e86
}
func (*InputChannelEmptyPredict) _InputChannel() {}

type InputChannelPredict struct {
	ChannelID  int64
	AccessHash int64
}

func (*InputChannelPredict) CRC() uint32 {
	return 0xf35aec28
}
func (*InputChannelPredict) _InputChannel() {}

type InputChannelFromMessagePredict struct {
	Peer      InputPeer
	MsgID     int32
	ChannelID int64
}

func (*InputChannelFromMessagePredict) CRC() uint32 {
	return 0x5b934f9d
}
func (*InputChannelFromMessagePredict) _InputChannel() {}

type InputChatPhoto interface {
	tl.Object
	_InputChatPhoto()
}

var (
	_ InputChatPhoto = (*InputChatPhotoEmptyPredict)(nil)
	_ InputChatPhoto = (*InputChatUploadedPhotoPredict)(nil)
	_ InputChatPhoto = (*InputChatPhotoPredict)(nil)
)

type InputChatPhotoEmptyPredict struct{}

func (*InputChatPhotoEmptyPredict) CRC() uint32 {
	return 0x1ca48f57
}
func (*InputChatPhotoEmptyPredict) _InputChatPhoto() {}

type InputChatUploadedPhotoPredict struct {
	_                struct{}  `tl:"flags,bitflag"`
	File             InputFile `tl:",omitempty:flags:0"`
	Video            InputFile `tl:",omitempty:flags:1"`
	VideoStartTs     *float64  `tl:",omitempty:flags:2"`
	VideoEmojiMarkup VideoSize `tl:",omitempty:flags:3"`
}

func (*InputChatUploadedPhotoPredict) CRC() uint32 {
	return 0xbdcdaec0
}
func (*InputChatUploadedPhotoPredict) _InputChatPhoto() {}

type InputChatPhotoPredict struct {
	ID InputPhoto
}

func (*InputChatPhotoPredict) CRC() uint32 {
	return 0x8953ad37
}
func (*InputChatPhotoPredict) _InputChatPhoto() {}

type InputChatlist interface {
	tl.Object
	_InputChatlist()
}

var (
	_ InputChatlist = (*InputChatlistDialogFilterPredict)(nil)
)

type InputChatlistDialogFilterPredict struct {
	FilterID int32
}

func (*InputChatlistDialogFilterPredict) CRC() uint32 {
	return 0xf3e0da33
}
func (*InputChatlistDialogFilterPredict) _InputChatlist() {}

type InputCheckPasswordSRP interface {
	tl.Object
	_InputCheckPasswordSRP()
}

var (
	_ InputCheckPasswordSRP = (*InputCheckPasswordEmptyPredict)(nil)
	_ InputCheckPasswordSRP = (*InputCheckPasswordSRPPredict)(nil)
)

type InputCheckPasswordEmptyPredict struct{}

func (*InputCheckPasswordEmptyPredict) CRC() uint32 {
	return 0x9880f658
}
func (*InputCheckPasswordEmptyPredict) _InputCheckPasswordSRP() {}

type InputCheckPasswordSRPPredict struct {
	SRPID int64
	A     []byte
	M1    []byte
}

func (*InputCheckPasswordSRPPredict) CRC() uint32 {
	return 0xd27ff082
}
func (*InputCheckPasswordSRPPredict) _InputCheckPasswordSRP() {}

type InputClientProxy interface {
	tl.Object
	_InputClientProxy()
}

var (
	_ InputClientProxy = (*InputClientProxyPredict)(nil)
)

type InputClientProxyPredict struct {
	Address string
	Port    int32
}

func (*InputClientProxyPredict) CRC() uint32 {
	return 0x75588b3f
}
func (*InputClientProxyPredict) _InputClientProxy() {}

type InputCollectible interface {
	tl.Object
	_InputCollectible()
}

var (
	_ InputCollectible = (*InputCollectibleUsernamePredict)(nil)
	_ InputCollectible = (*InputCollectiblePhonePredict)(nil)
)

type InputCollectibleUsernamePredict struct {
	Username string
}

func (*InputCollectibleUsernamePredict) CRC() uint32 {
	return 0xe39460a9
}
func (*InputCollectibleUsernamePredict) _InputCollectible() {}

type InputCollectiblePhonePredict struct {
	Phone string
}

func (*InputCollectiblePhonePredict) CRC() uint32 {
	return 0xa2e214a4
}
func (*InputCollectiblePhonePredict) _InputCollectible() {}

type InputContact interface {
	tl.Object
	_InputContact()
}

var (
	_ InputContact = (*InputPhoneContactPredict)(nil)
)

type InputPhoneContactPredict struct {
	ClientID  int64
	Phone     string
	FirstName string
	LastName  string
}

func (*InputPhoneContactPredict) CRC() uint32 {
	return 0xf392b7f4
}
func (*InputPhoneContactPredict) _InputContact() {}

type InputDialogPeer interface {
	tl.Object
	_InputDialogPeer()
}

var (
	_ InputDialogPeer = (*InputDialogPeerPredict)(nil)
	_ InputDialogPeer = (*InputDialogPeerFolderPredict)(nil)
)

type InputDialogPeerPredict struct {
	Peer InputPeer
}

func (*InputDialogPeerPredict) CRC() uint32 {
	return 0xfcaafeb7
}
func (*InputDialogPeerPredict) _InputDialogPeer() {}

type InputDialogPeerFolderPredict struct {
	FolderID int32
}

func (*InputDialogPeerFolderPredict) CRC() uint32 {
	return 0x64600527
}
func (*InputDialogPeerFolderPredict) _InputDialogPeer() {}

type InputDocument interface {
	tl.Object
	_InputDocument()
}

var (
	_ InputDocument = (*InputDocumentEmptyPredict)(nil)
	_ InputDocument = (*InputDocumentPredict)(nil)
)

type InputDocumentEmptyPredict struct{}

func (*InputDocumentEmptyPredict) CRC() uint32 {
	return 0x72f0eaae
}
func (*InputDocumentEmptyPredict) _InputDocument() {}

type InputDocumentPredict struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

func (*InputDocumentPredict) CRC() uint32 {
	return 0x1abfb575
}
func (*InputDocumentPredict) _InputDocument() {}

type InputEncryptedChat interface {
	tl.Object
	_InputEncryptedChat()
}

var (
	_ InputEncryptedChat = (*InputEncryptedChatPredict)(nil)
)

type InputEncryptedChatPredict struct {
	ChatID     int32
	AccessHash int64
}

func (*InputEncryptedChatPredict) CRC() uint32 {
	return 0xf141b5e1
}
func (*InputEncryptedChatPredict) _InputEncryptedChat() {}

type InputEncryptedFile interface {
	tl.Object
	_InputEncryptedFile()
}

var (
	_ InputEncryptedFile = (*InputEncryptedFileEmptyPredict)(nil)
	_ InputEncryptedFile = (*InputEncryptedFileUploadedPredict)(nil)
	_ InputEncryptedFile = (*InputEncryptedFilePredict)(nil)
	_ InputEncryptedFile = (*InputEncryptedFileBigUploadedPredict)(nil)
)

type InputEncryptedFileEmptyPredict struct{}

func (*InputEncryptedFileEmptyPredict) CRC() uint32 {
	return 0x1837c364
}
func (*InputEncryptedFileEmptyPredict) _InputEncryptedFile() {}

type InputEncryptedFileUploadedPredict struct {
	ID             int64
	Parts          int32
	Md5Checksum    string
	KeyFingerprint int32
}

func (*InputEncryptedFileUploadedPredict) CRC() uint32 {
	return 0x64bd0306
}
func (*InputEncryptedFileUploadedPredict) _InputEncryptedFile() {}

type InputEncryptedFilePredict struct {
	ID         int64
	AccessHash int64
}

func (*InputEncryptedFilePredict) CRC() uint32 {
	return 0x5a17b5e5
}
func (*InputEncryptedFilePredict) _InputEncryptedFile() {}

type InputEncryptedFileBigUploadedPredict struct {
	ID             int64
	Parts          int32
	KeyFingerprint int32
}

func (*InputEncryptedFileBigUploadedPredict) CRC() uint32 {
	return 0x2dc173c8
}
func (*InputEncryptedFileBigUploadedPredict) _InputEncryptedFile() {}

type InputFile interface {
	tl.Object
	_InputFile()
}

var (
	_ InputFile = (*InputFilePredict)(nil)
	_ InputFile = (*InputFileBigPredict)(nil)
)

type InputFilePredict struct {
	ID          int64
	Parts       int32
	Name        string
	Md5Checksum string
}

func (*InputFilePredict) CRC() uint32 {
	return 0xf52ff27f
}
func (*InputFilePredict) _InputFile() {}

type InputFileBigPredict struct {
	ID    int64
	Parts int32
	Name  string
}

func (*InputFileBigPredict) CRC() uint32 {
	return 0xfa4f0bb5
}
func (*InputFileBigPredict) _InputFile() {}

type InputFileLocation interface {
	tl.Object
	_InputFileLocation()
}

var (
	_ InputFileLocation = (*InputFileLocationPredict)(nil)
	_ InputFileLocation = (*InputEncryptedFileLocationPredict)(nil)
	_ InputFileLocation = (*InputDocumentFileLocationPredict)(nil)
	_ InputFileLocation = (*InputSecureFileLocationPredict)(nil)
	_ InputFileLocation = (*InputTakeoutFileLocationPredict)(nil)
	_ InputFileLocation = (*InputPhotoFileLocationPredict)(nil)
	_ InputFileLocation = (*InputPhotoLegacyFileLocationPredict)(nil)
	_ InputFileLocation = (*InputPeerPhotoFileLocationPredict)(nil)
	_ InputFileLocation = (*InputStickerSetThumbPredict)(nil)
	_ InputFileLocation = (*InputGroupCallStreamPredict)(nil)
)

type InputFileLocationPredict struct {
	VolumeID      int64
	LocalID       int32
	Secret        int64
	FileReference []byte
}

func (*InputFileLocationPredict) CRC() uint32 {
	return 0xdfdaabe1
}
func (*InputFileLocationPredict) _InputFileLocation() {}

type InputEncryptedFileLocationPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputEncryptedFileLocationPredict) CRC() uint32 {
	return 0xf5235d55
}
func (*InputEncryptedFileLocationPredict) _InputFileLocation() {}

type InputDocumentFileLocationPredict struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

func (*InputDocumentFileLocationPredict) CRC() uint32 {
	return 0xbad07584
}
func (*InputDocumentFileLocationPredict) _InputFileLocation() {}

type InputSecureFileLocationPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputSecureFileLocationPredict) CRC() uint32 {
	return 0xcbc7ee28
}
func (*InputSecureFileLocationPredict) _InputFileLocation() {}

type InputTakeoutFileLocationPredict struct{}

func (*InputTakeoutFileLocationPredict) CRC() uint32 {
	return 0x29be5899
}
func (*InputTakeoutFileLocationPredict) _InputFileLocation() {}

type InputPhotoFileLocationPredict struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

func (*InputPhotoFileLocationPredict) CRC() uint32 {
	return 0x40181ffe
}
func (*InputPhotoFileLocationPredict) _InputFileLocation() {}

type InputPhotoLegacyFileLocationPredict struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	VolumeID      int64
	LocalID       int32
	Secret        int64
}

func (*InputPhotoLegacyFileLocationPredict) CRC() uint32 {
	return 0xd83466f3
}
func (*InputPhotoLegacyFileLocationPredict) _InputFileLocation() {}

type InputPeerPhotoFileLocationPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Big     bool     `tl:",omitempty:flags:0,implicit"`
	Peer    InputPeer
	PhotoID int64
}

func (*InputPeerPhotoFileLocationPredict) CRC() uint32 {
	return 0x37257e99
}
func (*InputPeerPhotoFileLocationPredict) _InputFileLocation() {}

type InputStickerSetThumbPredict struct {
	Stickerset   InputStickerSet
	ThumbVersion int32
}

func (*InputStickerSetThumbPredict) CRC() uint32 {
	return 0x9d84f3db
}
func (*InputStickerSetThumbPredict) _InputFileLocation() {}

type InputGroupCallStreamPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Call         InputGroupCall
	TimeMs       int64
	Scale        int32
	VideoChannel *int32 `tl:",omitempty:flags:0"`
	VideoQuality *int32 `tl:",omitempty:flags:0"`
}

func (*InputGroupCallStreamPredict) CRC() uint32 {
	return 0x598a92a
}
func (*InputGroupCallStreamPredict) _InputFileLocation() {}

type InputFolderPeer interface {
	tl.Object
	_InputFolderPeer()
}

var (
	_ InputFolderPeer = (*InputFolderPeerPredict)(nil)
)

type InputFolderPeerPredict struct {
	Peer     InputPeer
	FolderID int32
}

func (*InputFolderPeerPredict) CRC() uint32 {
	return 0xfbd2c296
}
func (*InputFolderPeerPredict) _InputFolderPeer() {}

type InputGame interface {
	tl.Object
	_InputGame()
}

var (
	_ InputGame = (*InputGameIDPredict)(nil)
	_ InputGame = (*InputGameShortNamePredict)(nil)
)

type InputGameIDPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputGameIDPredict) CRC() uint32 {
	return 0x32c3e77
}
func (*InputGameIDPredict) _InputGame() {}

type InputGameShortNamePredict struct {
	BotID     InputUser
	ShortName string
}

func (*InputGameShortNamePredict) CRC() uint32 {
	return 0xc331e80a
}
func (*InputGameShortNamePredict) _InputGame() {}

type InputGeoPoint interface {
	tl.Object
	_InputGeoPoint()
}

var (
	_ InputGeoPoint = (*InputGeoPointEmptyPredict)(nil)
	_ InputGeoPoint = (*InputGeoPointPredict)(nil)
)

type InputGeoPointEmptyPredict struct{}

func (*InputGeoPointEmptyPredict) CRC() uint32 {
	return 0xe4c123d6
}
func (*InputGeoPointEmptyPredict) _InputGeoPoint() {}

type InputGeoPointPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Lat            float64
	Long           float64
	AccuracyRadius *int32 `tl:",omitempty:flags:0"`
}

func (*InputGeoPointPredict) CRC() uint32 {
	return 0x48222faf
}
func (*InputGeoPointPredict) _InputGeoPoint() {}

type InputGroupCall interface {
	tl.Object
	_InputGroupCall()
}

var (
	_ InputGroupCall = (*InputGroupCallPredict)(nil)
)

type InputGroupCallPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputGroupCallPredict) CRC() uint32 {
	return 0xd8aa840f
}
func (*InputGroupCallPredict) _InputGroupCall() {}

type InputInvoice interface {
	tl.Object
	_InputInvoice()
}

var (
	_ InputInvoice = (*InputInvoiceMessagePredict)(nil)
	_ InputInvoice = (*InputInvoiceSlugPredict)(nil)
	_ InputInvoice = (*InputInvoicePremiumGiftCodePredict)(nil)
	_ InputInvoice = (*InputInvoiceStarsPredict)(nil)
)

type InputInvoiceMessagePredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*InputInvoiceMessagePredict) CRC() uint32 {
	return 0xc5b56859
}
func (*InputInvoiceMessagePredict) _InputInvoice() {}

type InputInvoiceSlugPredict struct {
	Slug string
}

func (*InputInvoiceSlugPredict) CRC() uint32 {
	return 0xc326caef
}
func (*InputInvoiceSlugPredict) _InputInvoice() {}

type InputInvoicePremiumGiftCodePredict struct {
	Purpose InputStorePaymentPurpose
	Option  PremiumGiftCodeOption
}

func (*InputInvoicePremiumGiftCodePredict) CRC() uint32 {
	return 0x98986c0d
}
func (*InputInvoicePremiumGiftCodePredict) _InputInvoice() {}

type InputInvoiceStarsPredict struct {
	Purpose InputStorePaymentPurpose
}

func (*InputInvoiceStarsPredict) CRC() uint32 {
	return 0x65f00ce3
}
func (*InputInvoiceStarsPredict) _InputInvoice() {}

type InputMedia interface {
	tl.Object
	_InputMedia()
}

var (
	_ InputMedia = (*InputMediaEmptyPredict)(nil)
	_ InputMedia = (*InputMediaUploadedPhotoPredict)(nil)
	_ InputMedia = (*InputMediaPhotoPredict)(nil)
	_ InputMedia = (*InputMediaGeoPointPredict)(nil)
	_ InputMedia = (*InputMediaContactPredict)(nil)
	_ InputMedia = (*InputMediaUploadedDocumentPredict)(nil)
	_ InputMedia = (*InputMediaDocumentPredict)(nil)
	_ InputMedia = (*InputMediaVenuePredict)(nil)
	_ InputMedia = (*InputMediaPhotoExternalPredict)(nil)
	_ InputMedia = (*InputMediaDocumentExternalPredict)(nil)
	_ InputMedia = (*InputMediaGamePredict)(nil)
	_ InputMedia = (*InputMediaInvoicePredict)(nil)
	_ InputMedia = (*InputMediaGeoLivePredict)(nil)
	_ InputMedia = (*InputMediaPollPredict)(nil)
	_ InputMedia = (*InputMediaDicePredict)(nil)
	_ InputMedia = (*InputMediaStoryPredict)(nil)
	_ InputMedia = (*InputMediaWebPagePredict)(nil)
	_ InputMedia = (*InputMediaPaidMediaPredict)(nil)
)

type InputMediaEmptyPredict struct{}

func (*InputMediaEmptyPredict) CRC() uint32 {
	return 0x9664f57f
}
func (*InputMediaEmptyPredict) _InputMedia() {}

type InputMediaUploadedPhotoPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:2,implicit"`
	File       InputFile
	Stickers   []InputDocument `tl:",omitempty:flags:0"`
	TTLSeconds *int32          `tl:",omitempty:flags:1"`
}

func (*InputMediaUploadedPhotoPredict) CRC() uint32 {
	return 0x1e287d04
}
func (*InputMediaUploadedPhotoPredict) _InputMedia() {}

type InputMediaPhotoPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:1,implicit"`
	ID         InputPhoto
	TTLSeconds *int32 `tl:",omitempty:flags:0"`
}

func (*InputMediaPhotoPredict) CRC() uint32 {
	return 0xb3ba0635
}
func (*InputMediaPhotoPredict) _InputMedia() {}

type InputMediaGeoPointPredict struct {
	GeoPoint InputGeoPoint
}

func (*InputMediaGeoPointPredict) CRC() uint32 {
	return 0xf9c44144
}
func (*InputMediaGeoPointPredict) _InputMedia() {}

type InputMediaContactPredict struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
}

func (*InputMediaContactPredict) CRC() uint32 {
	return 0xf8ab7dfb
}
func (*InputMediaContactPredict) _InputMedia() {}

type InputMediaUploadedDocumentPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	NosoundVideo bool     `tl:",omitempty:flags:3,implicit"`
	ForceFile    bool     `tl:",omitempty:flags:4,implicit"`
	Spoiler      bool     `tl:",omitempty:flags:5,implicit"`
	File         InputFile
	Thumb        InputFile `tl:",omitempty:flags:2"`
	MimeType     string
	Attributes   []DocumentAttribute
	Stickers     []InputDocument `tl:",omitempty:flags:0"`
	TTLSeconds   *int32          `tl:",omitempty:flags:1"`
}

func (*InputMediaUploadedDocumentPredict) CRC() uint32 {
	return 0x5b38c6c1
}
func (*InputMediaUploadedDocumentPredict) _InputMedia() {}

type InputMediaDocumentPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:2,implicit"`
	ID         InputDocument
	TTLSeconds *int32  `tl:",omitempty:flags:0"`
	Query      *string `tl:",omitempty:flags:1"`
}

func (*InputMediaDocumentPredict) CRC() uint32 {
	return 0x33473058
}
func (*InputMediaDocumentPredict) _InputMedia() {}

type InputMediaVenuePredict struct {
	GeoPoint  InputGeoPoint
	Title     string
	Address   string
	Provider  string
	VenueID   string
	VenueType string
}

func (*InputMediaVenuePredict) CRC() uint32 {
	return 0xc13d1c11
}
func (*InputMediaVenuePredict) _InputMedia() {}

type InputMediaPhotoExternalPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:1,implicit"`
	URL        string
	TTLSeconds *int32 `tl:",omitempty:flags:0"`
}

func (*InputMediaPhotoExternalPredict) CRC() uint32 {
	return 0xe5bbfe1a
}
func (*InputMediaPhotoExternalPredict) _InputMedia() {}

type InputMediaDocumentExternalPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:1,implicit"`
	URL        string
	TTLSeconds *int32 `tl:",omitempty:flags:0"`
}

func (*InputMediaDocumentExternalPredict) CRC() uint32 {
	return 0xfb52dc99
}
func (*InputMediaDocumentExternalPredict) _InputMedia() {}

type InputMediaGamePredict struct {
	ID InputGame
}

func (*InputMediaGamePredict) CRC() uint32 {
	return 0xd33f43f3
}
func (*InputMediaGamePredict) _InputMedia() {}

type InputMediaInvoicePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Title         string
	Description   string
	Photo         InputWebDocument `tl:",omitempty:flags:0"`
	Invoice       Invoice
	Payload       []byte
	Provider      *string `tl:",omitempty:flags:3"`
	ProviderData  DataJSON
	StartParam    *string    `tl:",omitempty:flags:1"`
	ExtendedMedia InputMedia `tl:",omitempty:flags:2"`
}

func (*InputMediaInvoicePredict) CRC() uint32 {
	return 0x405fef0d
}
func (*InputMediaInvoicePredict) _InputMedia() {}

type InputMediaGeoLivePredict struct {
	_                           struct{} `tl:"flags,bitflag"`
	Stopped                     bool     `tl:",omitempty:flags:0,implicit"`
	GeoPoint                    InputGeoPoint
	Heading                     *int32 `tl:",omitempty:flags:2"`
	Period                      *int32 `tl:",omitempty:flags:1"`
	ProximityNotificationRadius *int32 `tl:",omitempty:flags:3"`
}

func (*InputMediaGeoLivePredict) CRC() uint32 {
	return 0x971fa843
}
func (*InputMediaGeoLivePredict) _InputMedia() {}

type InputMediaPollPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	Poll             Poll
	CorrectAnswers   [][]byte        `tl:",omitempty:flags:0"`
	Solution         *string         `tl:",omitempty:flags:1"`
	SolutionEntities []MessageEntity `tl:",omitempty:flags:1"`
}

func (*InputMediaPollPredict) CRC() uint32 {
	return 0xf94e5f1
}
func (*InputMediaPollPredict) _InputMedia() {}

type InputMediaDicePredict struct {
	Emoticon string
}

func (*InputMediaDicePredict) CRC() uint32 {
	return 0xe66fbf7b
}
func (*InputMediaDicePredict) _InputMedia() {}

type InputMediaStoryPredict struct {
	Peer InputPeer
	ID   int32
}

func (*InputMediaStoryPredict) CRC() uint32 {
	return 0x89fdd778
}
func (*InputMediaStoryPredict) _InputMedia() {}

type InputMediaWebPagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ForceLargeMedia bool     `tl:",omitempty:flags:0,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:1,implicit"`
	Optional        bool     `tl:",omitempty:flags:2,implicit"`
	URL             string
}

func (*InputMediaWebPagePredict) CRC() uint32 {
	return 0xc21b8849
}
func (*InputMediaWebPagePredict) _InputMedia() {}

type InputMediaPaidMediaPredict struct {
	StarsAmount   int64
	ExtendedMedia []InputMedia
}

func (*InputMediaPaidMediaPredict) CRC() uint32 {
	return 0xaa661fc3
}
func (*InputMediaPaidMediaPredict) _InputMedia() {}

type InputMessage interface {
	tl.Object
	_InputMessage()
}

var (
	_ InputMessage = (*InputMessageIDPredict)(nil)
	_ InputMessage = (*InputMessageReplyToPredict)(nil)
	_ InputMessage = (*InputMessagePinnedPredict)(nil)
	_ InputMessage = (*InputMessageCallbackQueryPredict)(nil)
)

type InputMessageIDPredict struct {
	ID int32
}

func (*InputMessageIDPredict) CRC() uint32 {
	return 0xa676a322
}
func (*InputMessageIDPredict) _InputMessage() {}

type InputMessageReplyToPredict struct {
	ID int32
}

func (*InputMessageReplyToPredict) CRC() uint32 {
	return 0xbad88395
}
func (*InputMessageReplyToPredict) _InputMessage() {}

type InputMessagePinnedPredict struct{}

func (*InputMessagePinnedPredict) CRC() uint32 {
	return 0x86872538
}
func (*InputMessagePinnedPredict) _InputMessage() {}

type InputMessageCallbackQueryPredict struct {
	ID      int32
	QueryID int64
}

func (*InputMessageCallbackQueryPredict) CRC() uint32 {
	return 0xacfa1a7e
}
func (*InputMessageCallbackQueryPredict) _InputMessage() {}

type InputNotifyPeer interface {
	tl.Object
	_InputNotifyPeer()
}

var (
	_ InputNotifyPeer = (*InputNotifyPeerPredict)(nil)
	_ InputNotifyPeer = (*InputNotifyUsersPredict)(nil)
	_ InputNotifyPeer = (*InputNotifyChatsPredict)(nil)
	_ InputNotifyPeer = (*InputNotifyBroadcastsPredict)(nil)
	_ InputNotifyPeer = (*InputNotifyForumTopicPredict)(nil)
)

type InputNotifyPeerPredict struct {
	Peer InputPeer
}

func (*InputNotifyPeerPredict) CRC() uint32 {
	return 0xb8bc5b0c
}
func (*InputNotifyPeerPredict) _InputNotifyPeer() {}

type InputNotifyUsersPredict struct{}

func (*InputNotifyUsersPredict) CRC() uint32 {
	return 0x193b4417
}
func (*InputNotifyUsersPredict) _InputNotifyPeer() {}

type InputNotifyChatsPredict struct{}

func (*InputNotifyChatsPredict) CRC() uint32 {
	return 0x4a95e84e
}
func (*InputNotifyChatsPredict) _InputNotifyPeer() {}

type InputNotifyBroadcastsPredict struct{}

func (*InputNotifyBroadcastsPredict) CRC() uint32 {
	return 0xb1db7c7e
}
func (*InputNotifyBroadcastsPredict) _InputNotifyPeer() {}

type InputNotifyForumTopicPredict struct {
	Peer     InputPeer
	TopMsgID int32
}

func (*InputNotifyForumTopicPredict) CRC() uint32 {
	return 0x5c467992
}
func (*InputNotifyForumTopicPredict) _InputNotifyPeer() {}

type InputPaymentCredentials interface {
	tl.Object
	_InputPaymentCredentials()
}

var (
	_ InputPaymentCredentials = (*InputPaymentCredentialsSavedPredict)(nil)
	_ InputPaymentCredentials = (*InputPaymentCredentialsPredict)(nil)
	_ InputPaymentCredentials = (*InputPaymentCredentialsApplePayPredict)(nil)
	_ InputPaymentCredentials = (*InputPaymentCredentialsGooglePayPredict)(nil)
)

type InputPaymentCredentialsSavedPredict struct {
	ID          string
	TmpPassword []byte
}

func (*InputPaymentCredentialsSavedPredict) CRC() uint32 {
	return 0xc10eb2cf
}
func (*InputPaymentCredentialsSavedPredict) _InputPaymentCredentials() {}

type InputPaymentCredentialsPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Save bool     `tl:",omitempty:flags:0,implicit"`
	Data DataJSON
}

func (*InputPaymentCredentialsPredict) CRC() uint32 {
	return 0x3417d728
}
func (*InputPaymentCredentialsPredict) _InputPaymentCredentials() {}

type InputPaymentCredentialsApplePayPredict struct {
	PaymentData DataJSON
}

func (*InputPaymentCredentialsApplePayPredict) CRC() uint32 {
	return 0xaa1c39f
}
func (*InputPaymentCredentialsApplePayPredict) _InputPaymentCredentials() {}

type InputPaymentCredentialsGooglePayPredict struct {
	PaymentToken DataJSON
}

func (*InputPaymentCredentialsGooglePayPredict) CRC() uint32 {
	return 0x8ac32801
}
func (*InputPaymentCredentialsGooglePayPredict) _InputPaymentCredentials() {}

type InputPeer interface {
	tl.Object
	_InputPeer()
}

var (
	_ InputPeer = (*InputPeerEmptyPredict)(nil)
	_ InputPeer = (*InputPeerSelfPredict)(nil)
	_ InputPeer = (*InputPeerChatPredict)(nil)
	_ InputPeer = (*InputPeerUserPredict)(nil)
	_ InputPeer = (*InputPeerChannelPredict)(nil)
	_ InputPeer = (*InputPeerUserFromMessagePredict)(nil)
	_ InputPeer = (*InputPeerChannelFromMessagePredict)(nil)
)

type InputPeerEmptyPredict struct{}

func (*InputPeerEmptyPredict) CRC() uint32 {
	return 0x7f3b18ea
}
func (*InputPeerEmptyPredict) _InputPeer() {}

type InputPeerSelfPredict struct{}

func (*InputPeerSelfPredict) CRC() uint32 {
	return 0x7da07ec9
}
func (*InputPeerSelfPredict) _InputPeer() {}

type InputPeerChatPredict struct {
	ChatID int64
}

func (*InputPeerChatPredict) CRC() uint32 {
	return 0x35a95cb9
}
func (*InputPeerChatPredict) _InputPeer() {}

type InputPeerUserPredict struct {
	UserID     int64
	AccessHash int64
}

func (*InputPeerUserPredict) CRC() uint32 {
	return 0xdde8a54c
}
func (*InputPeerUserPredict) _InputPeer() {}

type InputPeerChannelPredict struct {
	ChannelID  int64
	AccessHash int64
}

func (*InputPeerChannelPredict) CRC() uint32 {
	return 0x27bcbbfc
}
func (*InputPeerChannelPredict) _InputPeer() {}

type InputPeerUserFromMessagePredict struct {
	Peer   InputPeer
	MsgID  int32
	UserID int64
}

func (*InputPeerUserFromMessagePredict) CRC() uint32 {
	return 0xa87b0a1c
}
func (*InputPeerUserFromMessagePredict) _InputPeer() {}

type InputPeerChannelFromMessagePredict struct {
	Peer      InputPeer
	MsgID     int32
	ChannelID int64
}

func (*InputPeerChannelFromMessagePredict) CRC() uint32 {
	return 0xbd2a0840
}
func (*InputPeerChannelFromMessagePredict) _InputPeer() {}

type InputPeerNotifySettings interface {
	tl.Object
	_InputPeerNotifySettings()
}

var (
	_ InputPeerNotifySettings = (*InputPeerNotifySettingsPredict)(nil)
)

type InputPeerNotifySettingsPredict struct {
	_                 struct{}          `tl:"flags,bitflag"`
	ShowPreviews      *bool             `tl:",omitempty:flags:0"`
	Silent            *bool             `tl:",omitempty:flags:1"`
	MuteUntil         *int32            `tl:",omitempty:flags:2"`
	Sound             NotificationSound `tl:",omitempty:flags:3"`
	StoriesMuted      *bool             `tl:",omitempty:flags:6"`
	StoriesHideSender *bool             `tl:",omitempty:flags:7"`
	StoriesSound      NotificationSound `tl:",omitempty:flags:8"`
}

func (*InputPeerNotifySettingsPredict) CRC() uint32 {
	return 0xcacb6ae2
}
func (*InputPeerNotifySettingsPredict) _InputPeerNotifySettings() {}

type InputPhoneCall interface {
	tl.Object
	_InputPhoneCall()
}

var (
	_ InputPhoneCall = (*InputPhoneCallPredict)(nil)
)

type InputPhoneCallPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputPhoneCallPredict) CRC() uint32 {
	return 0x1e36fded
}
func (*InputPhoneCallPredict) _InputPhoneCall() {}

type InputPhoto interface {
	tl.Object
	_InputPhoto()
}

var (
	_ InputPhoto = (*InputPhotoEmptyPredict)(nil)
	_ InputPhoto = (*InputPhotoPredict)(nil)
)

type InputPhotoEmptyPredict struct{}

func (*InputPhotoEmptyPredict) CRC() uint32 {
	return 0x1cd7bf0d
}
func (*InputPhotoEmptyPredict) _InputPhoto() {}

type InputPhotoPredict struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

func (*InputPhotoPredict) CRC() uint32 {
	return 0x3bb3b94a
}
func (*InputPhotoPredict) _InputPhoto() {}

type InputPrivacyRule interface {
	tl.Object
	_InputPrivacyRule()
}

var (
	_ InputPrivacyRule = (*InputPrivacyValueAllowContactsPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueAllowAllPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueAllowUsersPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueDisallowContactsPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueDisallowAllPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueDisallowUsersPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueAllowChatParticipantsPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueDisallowChatParticipantsPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueAllowCloseFriendsPredict)(nil)
	_ InputPrivacyRule = (*InputPrivacyValueAllowPremiumPredict)(nil)
)

type InputPrivacyValueAllowContactsPredict struct{}

func (*InputPrivacyValueAllowContactsPredict) CRC() uint32 {
	return 0xd09e07b
}
func (*InputPrivacyValueAllowContactsPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowAllPredict struct{}

func (*InputPrivacyValueAllowAllPredict) CRC() uint32 {
	return 0x184b35ce
}
func (*InputPrivacyValueAllowAllPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowUsersPredict struct {
	Users []InputUser
}

func (*InputPrivacyValueAllowUsersPredict) CRC() uint32 {
	return 0x131cc67f
}
func (*InputPrivacyValueAllowUsersPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowContactsPredict struct{}

func (*InputPrivacyValueDisallowContactsPredict) CRC() uint32 {
	return 0xba52007
}
func (*InputPrivacyValueDisallowContactsPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowAllPredict struct{}

func (*InputPrivacyValueDisallowAllPredict) CRC() uint32 {
	return 0xd66b66c9
}
func (*InputPrivacyValueDisallowAllPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowUsersPredict struct {
	Users []InputUser
}

func (*InputPrivacyValueDisallowUsersPredict) CRC() uint32 {
	return 0x90110467
}
func (*InputPrivacyValueDisallowUsersPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowChatParticipantsPredict struct {
	Chats []int64
}

func (*InputPrivacyValueAllowChatParticipantsPredict) CRC() uint32 {
	return 0x840649cf
}
func (*InputPrivacyValueAllowChatParticipantsPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowChatParticipantsPredict struct {
	Chats []int64
}

func (*InputPrivacyValueDisallowChatParticipantsPredict) CRC() uint32 {
	return 0xe94f0f86
}
func (*InputPrivacyValueDisallowChatParticipantsPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowCloseFriendsPredict struct{}

func (*InputPrivacyValueAllowCloseFriendsPredict) CRC() uint32 {
	return 0x2f453e49
}
func (*InputPrivacyValueAllowCloseFriendsPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowPremiumPredict struct{}

func (*InputPrivacyValueAllowPremiumPredict) CRC() uint32 {
	return 0x77cdc9f1
}
func (*InputPrivacyValueAllowPremiumPredict) _InputPrivacyRule() {}

type InputQuickReplyShortcut interface {
	tl.Object
	_InputQuickReplyShortcut()
}

var (
	_ InputQuickReplyShortcut = (*InputQuickReplyShortcutPredict)(nil)
	_ InputQuickReplyShortcut = (*InputQuickReplyShortcutIDPredict)(nil)
)

type InputQuickReplyShortcutPredict struct {
	Shortcut string
}

func (*InputQuickReplyShortcutPredict) CRC() uint32 {
	return 0x24596d41
}
func (*InputQuickReplyShortcutPredict) _InputQuickReplyShortcut() {}

type InputQuickReplyShortcutIDPredict struct {
	ShortcutID int32
}

func (*InputQuickReplyShortcutIDPredict) CRC() uint32 {
	return 0x1190cf1
}
func (*InputQuickReplyShortcutIDPredict) _InputQuickReplyShortcut() {}

type InputReplyTo interface {
	tl.Object
	_InputReplyTo()
}

var (
	_ InputReplyTo = (*InputReplyToMessagePredict)(nil)
	_ InputReplyTo = (*InputReplyToStoryPredict)(nil)
)

type InputReplyToMessagePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ReplyToMsgID  int32
	TopMsgID      *int32          `tl:",omitempty:flags:0"`
	ReplyToPeerID InputPeer       `tl:",omitempty:flags:1"`
	QuoteText     *string         `tl:",omitempty:flags:2"`
	QuoteEntities []MessageEntity `tl:",omitempty:flags:3"`
	QuoteOffset   *int32          `tl:",omitempty:flags:4"`
}

func (*InputReplyToMessagePredict) CRC() uint32 {
	return 0x22c0f6d5
}
func (*InputReplyToMessagePredict) _InputReplyTo() {}

type InputReplyToStoryPredict struct {
	Peer    InputPeer
	StoryID int32
}

func (*InputReplyToStoryPredict) CRC() uint32 {
	return 0x5881323a
}
func (*InputReplyToStoryPredict) _InputReplyTo() {}

type InputSecureFile interface {
	tl.Object
	_InputSecureFile()
}

var (
	_ InputSecureFile = (*InputSecureFileUploadedPredict)(nil)
	_ InputSecureFile = (*InputSecureFilePredict)(nil)
)

type InputSecureFileUploadedPredict struct {
	ID          int64
	Parts       int32
	Md5Checksum string
	FileHash    []byte
	Secret      []byte
}

func (*InputSecureFileUploadedPredict) CRC() uint32 {
	return 0x3334b0f0
}
func (*InputSecureFileUploadedPredict) _InputSecureFile() {}

type InputSecureFilePredict struct {
	ID         int64
	AccessHash int64
}

func (*InputSecureFilePredict) CRC() uint32 {
	return 0x5367e5be
}
func (*InputSecureFilePredict) _InputSecureFile() {}

type InputSecureValue interface {
	tl.Object
	_InputSecureValue()
}

var (
	_ InputSecureValue = (*InputSecureValuePredict)(nil)
)

type InputSecureValuePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Type        SecureValueType
	Data        SecureData        `tl:",omitempty:flags:0"`
	FrontSide   InputSecureFile   `tl:",omitempty:flags:1"`
	ReverseSide InputSecureFile   `tl:",omitempty:flags:2"`
	Selfie      InputSecureFile   `tl:",omitempty:flags:3"`
	Translation []InputSecureFile `tl:",omitempty:flags:6"`
	Files       []InputSecureFile `tl:",omitempty:flags:4"`
	PlainData   SecurePlainData   `tl:",omitempty:flags:5"`
}

func (*InputSecureValuePredict) CRC() uint32 {
	return 0xdb21d0a7
}
func (*InputSecureValuePredict) _InputSecureValue() {}

type InputSingleMedia interface {
	tl.Object
	_InputSingleMedia()
}

var (
	_ InputSingleMedia = (*InputSingleMediaPredict)(nil)
)

type InputSingleMediaPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Media    InputMedia
	RandomID int64
	Message  string
	Entities []MessageEntity `tl:",omitempty:flags:0"`
}

func (*InputSingleMediaPredict) CRC() uint32 {
	return 0x1cc6e91f
}
func (*InputSingleMediaPredict) _InputSingleMedia() {}

type InputStarsTransaction interface {
	tl.Object
	_InputStarsTransaction()
}

var (
	_ InputStarsTransaction = (*InputStarsTransactionPredict)(nil)
)

type InputStarsTransactionPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Refund bool     `tl:",omitempty:flags:0,implicit"`
	ID     string
}

func (*InputStarsTransactionPredict) CRC() uint32 {
	return 0x206ae6d1
}
func (*InputStarsTransactionPredict) _InputStarsTransaction() {}

type InputStickerSet interface {
	tl.Object
	_InputStickerSet()
}

var (
	_ InputStickerSet = (*InputStickerSetEmptyPredict)(nil)
	_ InputStickerSet = (*InputStickerSetIDPredict)(nil)
	_ InputStickerSet = (*InputStickerSetShortNamePredict)(nil)
	_ InputStickerSet = (*InputStickerSetAnimatedEmojiPredict)(nil)
	_ InputStickerSet = (*InputStickerSetDicePredict)(nil)
	_ InputStickerSet = (*InputStickerSetAnimatedEmojiAnimationsPredict)(nil)
	_ InputStickerSet = (*InputStickerSetPremiumGiftsPredict)(nil)
	_ InputStickerSet = (*InputStickerSetEmojiGenericAnimationsPredict)(nil)
	_ InputStickerSet = (*InputStickerSetEmojiDefaultStatusesPredict)(nil)
	_ InputStickerSet = (*InputStickerSetEmojiDefaultTopicIconsPredict)(nil)
	_ InputStickerSet = (*InputStickerSetEmojiChannelDefaultStatusesPredict)(nil)
)

type InputStickerSetEmptyPredict struct{}

func (*InputStickerSetEmptyPredict) CRC() uint32 {
	return 0xffb62b95
}
func (*InputStickerSetEmptyPredict) _InputStickerSet() {}

type InputStickerSetIDPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputStickerSetIDPredict) CRC() uint32 {
	return 0x9de7a269
}
func (*InputStickerSetIDPredict) _InputStickerSet() {}

type InputStickerSetShortNamePredict struct {
	ShortName string
}

func (*InputStickerSetShortNamePredict) CRC() uint32 {
	return 0x861cc8a0
}
func (*InputStickerSetShortNamePredict) _InputStickerSet() {}

type InputStickerSetAnimatedEmojiPredict struct{}

func (*InputStickerSetAnimatedEmojiPredict) CRC() uint32 {
	return 0x28703c8
}
func (*InputStickerSetAnimatedEmojiPredict) _InputStickerSet() {}

type InputStickerSetDicePredict struct {
	Emoticon string
}

func (*InputStickerSetDicePredict) CRC() uint32 {
	return 0xe67f520e
}
func (*InputStickerSetDicePredict) _InputStickerSet() {}

type InputStickerSetAnimatedEmojiAnimationsPredict struct{}

func (*InputStickerSetAnimatedEmojiAnimationsPredict) CRC() uint32 {
	return 0xcde3739
}
func (*InputStickerSetAnimatedEmojiAnimationsPredict) _InputStickerSet() {}

type InputStickerSetPremiumGiftsPredict struct{}

func (*InputStickerSetPremiumGiftsPredict) CRC() uint32 {
	return 0xc88b3b02
}
func (*InputStickerSetPremiumGiftsPredict) _InputStickerSet() {}

type InputStickerSetEmojiGenericAnimationsPredict struct{}

func (*InputStickerSetEmojiGenericAnimationsPredict) CRC() uint32 {
	return 0x4c4d4ce
}
func (*InputStickerSetEmojiGenericAnimationsPredict) _InputStickerSet() {}

type InputStickerSetEmojiDefaultStatusesPredict struct{}

func (*InputStickerSetEmojiDefaultStatusesPredict) CRC() uint32 {
	return 0x29d0f5ee
}
func (*InputStickerSetEmojiDefaultStatusesPredict) _InputStickerSet() {}

type InputStickerSetEmojiDefaultTopicIconsPredict struct{}

func (*InputStickerSetEmojiDefaultTopicIconsPredict) CRC() uint32 {
	return 0x44c1f8e9
}
func (*InputStickerSetEmojiDefaultTopicIconsPredict) _InputStickerSet() {}

type InputStickerSetEmojiChannelDefaultStatusesPredict struct{}

func (*InputStickerSetEmojiChannelDefaultStatusesPredict) CRC() uint32 {
	return 0x49748553
}
func (*InputStickerSetEmojiChannelDefaultStatusesPredict) _InputStickerSet() {}

type InputStickerSetItem interface {
	tl.Object
	_InputStickerSetItem()
}

var (
	_ InputStickerSetItem = (*InputStickerSetItemPredict)(nil)
)

type InputStickerSetItemPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Document   InputDocument
	Emoji      string
	MaskCoords MaskCoords `tl:",omitempty:flags:0"`
	Keywords   *string    `tl:",omitempty:flags:1"`
}

func (*InputStickerSetItemPredict) CRC() uint32 {
	return 0x32da9e9c
}
func (*InputStickerSetItemPredict) _InputStickerSetItem() {}

type InputStickeredMedia interface {
	tl.Object
	_InputStickeredMedia()
}

var (
	_ InputStickeredMedia = (*InputStickeredMediaPhotoPredict)(nil)
	_ InputStickeredMedia = (*InputStickeredMediaDocumentPredict)(nil)
)

type InputStickeredMediaPhotoPredict struct {
	ID InputPhoto
}

func (*InputStickeredMediaPhotoPredict) CRC() uint32 {
	return 0x4a992157
}
func (*InputStickeredMediaPhotoPredict) _InputStickeredMedia() {}

type InputStickeredMediaDocumentPredict struct {
	ID InputDocument
}

func (*InputStickeredMediaDocumentPredict) CRC() uint32 {
	return 0x438865b
}
func (*InputStickeredMediaDocumentPredict) _InputStickeredMedia() {}

type InputStorePaymentPurpose interface {
	tl.Object
	_InputStorePaymentPurpose()
}

var (
	_ InputStorePaymentPurpose = (*InputStorePaymentPremiumSubscriptionPredict)(nil)
	_ InputStorePaymentPurpose = (*InputStorePaymentGiftPremiumPredict)(nil)
	_ InputStorePaymentPurpose = (*InputStorePaymentPremiumGiftCodePredict)(nil)
	_ InputStorePaymentPurpose = (*InputStorePaymentPremiumGiveawayPredict)(nil)
	_ InputStorePaymentPurpose = (*InputStorePaymentStarsTopupPredict)(nil)
	_ InputStorePaymentPurpose = (*InputStorePaymentStarsGiftPredict)(nil)
)

type InputStorePaymentPremiumSubscriptionPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Restore bool     `tl:",omitempty:flags:0,implicit"`
	Upgrade bool     `tl:",omitempty:flags:1,implicit"`
}

func (*InputStorePaymentPremiumSubscriptionPredict) CRC() uint32 {
	return 0xa6751e66
}
func (*InputStorePaymentPremiumSubscriptionPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentGiftPremiumPredict struct {
	UserID   InputUser
	Currency string
	Amount   int64
}

func (*InputStorePaymentGiftPremiumPredict) CRC() uint32 {
	return 0x616f7fe8
}
func (*InputStorePaymentGiftPremiumPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentPremiumGiftCodePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Users     []InputUser
	BoostPeer InputPeer `tl:",omitempty:flags:0"`
	Currency  string
	Amount    int64
}

func (*InputStorePaymentPremiumGiftCodePredict) CRC() uint32 {
	return 0xa3805f3f
}
func (*InputStorePaymentPremiumGiftCodePredict) _InputStorePaymentPurpose() {}

type InputStorePaymentPremiumGiveawayPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers bool     `tl:",omitempty:flags:0,implicit"`
	WinnersAreVisible  bool     `tl:",omitempty:flags:3,implicit"`
	BoostPeer          InputPeer
	AdditionalPeers    []InputPeer `tl:",omitempty:flags:1"`
	CountriesIso2      []string    `tl:",omitempty:flags:2"`
	PrizeDescription   *string     `tl:",omitempty:flags:4"`
	RandomID           int64
	UntilDate          int32
	Currency           string
	Amount             int64
}

func (*InputStorePaymentPremiumGiveawayPredict) CRC() uint32 {
	return 0x160544ca
}
func (*InputStorePaymentPremiumGiveawayPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentStarsTopupPredict struct {
	Stars    int64
	Currency string
	Amount   int64
}

func (*InputStorePaymentStarsTopupPredict) CRC() uint32 {
	return 0xdddd0f56
}
func (*InputStorePaymentStarsTopupPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentStarsGiftPredict struct {
	UserID   InputUser
	Stars    int64
	Currency string
	Amount   int64
}

func (*InputStorePaymentStarsGiftPredict) CRC() uint32 {
	return 0x1d741ef7
}
func (*InputStorePaymentStarsGiftPredict) _InputStorePaymentPurpose() {}

type InputTheme interface {
	tl.Object
	_InputTheme()
}

var (
	_ InputTheme = (*InputThemePredict)(nil)
	_ InputTheme = (*InputThemeSlugPredict)(nil)
)

type InputThemePredict struct {
	ID         int64
	AccessHash int64
}

func (*InputThemePredict) CRC() uint32 {
	return 0x3c5693e9
}
func (*InputThemePredict) _InputTheme() {}

type InputThemeSlugPredict struct {
	Slug string
}

func (*InputThemeSlugPredict) CRC() uint32 {
	return 0xf5890df1
}
func (*InputThemeSlugPredict) _InputTheme() {}

type InputThemeSettings interface {
	tl.Object
	_InputThemeSettings()
}

var (
	_ InputThemeSettings = (*InputThemeSettingsPredict)(nil)
)

type InputThemeSettingsPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	MessageColorsAnimated bool     `tl:",omitempty:flags:2,implicit"`
	BaseTheme             BaseTheme
	AccentColor           int32
	OutboxAccentColor     *int32            `tl:",omitempty:flags:3"`
	MessageColors         []int32           `tl:",omitempty:flags:0"`
	Wallpaper             InputWallPaper    `tl:",omitempty:flags:1"`
	WallpaperSettings     WallPaperSettings `tl:",omitempty:flags:1"`
}

func (*InputThemeSettingsPredict) CRC() uint32 {
	return 0x8fde504f
}
func (*InputThemeSettingsPredict) _InputThemeSettings() {}

type InputUser interface {
	tl.Object
	_InputUser()
}

var (
	_ InputUser = (*InputUserEmptyPredict)(nil)
	_ InputUser = (*InputUserSelfPredict)(nil)
	_ InputUser = (*InputUserPredict)(nil)
	_ InputUser = (*InputUserFromMessagePredict)(nil)
)

type InputUserEmptyPredict struct{}

func (*InputUserEmptyPredict) CRC() uint32 {
	return 0xb98886cf
}
func (*InputUserEmptyPredict) _InputUser() {}

type InputUserSelfPredict struct{}

func (*InputUserSelfPredict) CRC() uint32 {
	return 0xf7c1b13f
}
func (*InputUserSelfPredict) _InputUser() {}

type InputUserPredict struct {
	UserID     int64
	AccessHash int64
}

func (*InputUserPredict) CRC() uint32 {
	return 0xf21158c6
}
func (*InputUserPredict) _InputUser() {}

type InputUserFromMessagePredict struct {
	Peer   InputPeer
	MsgID  int32
	UserID int64
}

func (*InputUserFromMessagePredict) CRC() uint32 {
	return 0x1da448e2
}
func (*InputUserFromMessagePredict) _InputUser() {}

type InputWallPaper interface {
	tl.Object
	_InputWallPaper()
}

var (
	_ InputWallPaper = (*InputWallPaperPredict)(nil)
	_ InputWallPaper = (*InputWallPaperSlugPredict)(nil)
	_ InputWallPaper = (*InputWallPaperNoFilePredict)(nil)
)

type InputWallPaperPredict struct {
	ID         int64
	AccessHash int64
}

func (*InputWallPaperPredict) CRC() uint32 {
	return 0xe630b979
}
func (*InputWallPaperPredict) _InputWallPaper() {}

type InputWallPaperSlugPredict struct {
	Slug string
}

func (*InputWallPaperSlugPredict) CRC() uint32 {
	return 0x72091c80
}
func (*InputWallPaperSlugPredict) _InputWallPaper() {}

type InputWallPaperNoFilePredict struct {
	ID int64
}

func (*InputWallPaperNoFilePredict) CRC() uint32 {
	return 0x967a462e
}
func (*InputWallPaperNoFilePredict) _InputWallPaper() {}

type InputWebDocument interface {
	tl.Object
	_InputWebDocument()
}

var (
	_ InputWebDocument = (*InputWebDocumentPredict)(nil)
)

type InputWebDocumentPredict struct {
	URL        string
	Size       int32
	MimeType   string
	Attributes []DocumentAttribute
}

func (*InputWebDocumentPredict) CRC() uint32 {
	return 0x9bed434d
}
func (*InputWebDocumentPredict) _InputWebDocument() {}

type InputWebFileLocation interface {
	tl.Object
	_InputWebFileLocation()
}

var (
	_ InputWebFileLocation = (*InputWebFileLocationPredict)(nil)
	_ InputWebFileLocation = (*InputWebFileGeoPointLocationPredict)(nil)
	_ InputWebFileLocation = (*InputWebFileAudioAlbumThumbLocationPredict)(nil)
)

type InputWebFileLocationPredict struct {
	URL        string
	AccessHash int64
}

func (*InputWebFileLocationPredict) CRC() uint32 {
	return 0xc239d686
}
func (*InputWebFileLocationPredict) _InputWebFileLocation() {}

type InputWebFileGeoPointLocationPredict struct {
	GeoPoint   InputGeoPoint
	AccessHash int64
	W          int32
	H          int32
	Zoom       int32
	Scale      int32
}

func (*InputWebFileGeoPointLocationPredict) CRC() uint32 {
	return 0x9f2221c9
}
func (*InputWebFileGeoPointLocationPredict) _InputWebFileLocation() {}

type InputWebFileAudioAlbumThumbLocationPredict struct {
	_         struct{}      `tl:"flags,bitflag"`
	Small     bool          `tl:",omitempty:flags:2,implicit"`
	Document  InputDocument `tl:",omitempty:flags:0"`
	Title     *string       `tl:",omitempty:flags:1"`
	Performer *string       `tl:",omitempty:flags:1"`
}

func (*InputWebFileAudioAlbumThumbLocationPredict) CRC() uint32 {
	return 0xf46fe924
}
func (*InputWebFileAudioAlbumThumbLocationPredict) _InputWebFileLocation() {}

type Invoice interface {
	tl.Object
	_Invoice()
}

var (
	_ Invoice = (*InvoicePredict)(nil)
)

type InvoicePredict struct {
	_                        struct{} `tl:"flags,bitflag"`
	Test                     bool     `tl:",omitempty:flags:0,implicit"`
	NameRequested            bool     `tl:",omitempty:flags:1,implicit"`
	PhoneRequested           bool     `tl:",omitempty:flags:2,implicit"`
	EmailRequested           bool     `tl:",omitempty:flags:3,implicit"`
	ShippingAddressRequested bool     `tl:",omitempty:flags:4,implicit"`
	Flexible                 bool     `tl:",omitempty:flags:5,implicit"`
	PhoneToProvider          bool     `tl:",omitempty:flags:6,implicit"`
	EmailToProvider          bool     `tl:",omitempty:flags:7,implicit"`
	Recurring                bool     `tl:",omitempty:flags:9,implicit"`
	Currency                 string
	Prices                   []LabeledPrice
	MaxTipAmount             *int64  `tl:",omitempty:flags:8"`
	SuggestedTipAmounts      []int64 `tl:",omitempty:flags:8"`
	TermsURL                 *string `tl:",omitempty:flags:10"`
}

func (*InvoicePredict) CRC() uint32 {
	return 0x5db95a15
}
func (*InvoicePredict) _Invoice() {}

type JSONObjectValue interface {
	tl.Object
	_JSONObjectValue()
}

var (
	_ JSONObjectValue = (*JSONObjectValuePredict)(nil)
)

type JSONObjectValuePredict struct {
	Key   string
	Value JSONValue
}

func (*JSONObjectValuePredict) CRC() uint32 {
	return 0xc0de1bd9
}
func (*JSONObjectValuePredict) _JSONObjectValue() {}

type JSONValue interface {
	tl.Object
	_JSONValue()
}

var (
	_ JSONValue = (*JSONNullPredict)(nil)
	_ JSONValue = (*JSONBoolPredict)(nil)
	_ JSONValue = (*JSONNumberPredict)(nil)
	_ JSONValue = (*JSONStringPredict)(nil)
	_ JSONValue = (*JSONArrayPredict)(nil)
	_ JSONValue = (*JSONObjectPredict)(nil)
)

type JSONNullPredict struct{}

func (*JSONNullPredict) CRC() uint32 {
	return 0x3f6d7b68
}
func (*JSONNullPredict) _JSONValue() {}

type JSONBoolPredict struct {
	Value bool
}

func (*JSONBoolPredict) CRC() uint32 {
	return 0xc7345e6a
}
func (*JSONBoolPredict) _JSONValue() {}

type JSONNumberPredict struct {
	Value float64
}

func (*JSONNumberPredict) CRC() uint32 {
	return 0x2be0dfa4
}
func (*JSONNumberPredict) _JSONValue() {}

type JSONStringPredict struct {
	Value string
}

func (*JSONStringPredict) CRC() uint32 {
	return 0xb71e767a
}
func (*JSONStringPredict) _JSONValue() {}

type JSONArrayPredict struct {
	Value []JSONValue
}

func (*JSONArrayPredict) CRC() uint32 {
	return 0xf7444763
}
func (*JSONArrayPredict) _JSONValue() {}

type JSONObjectPredict struct {
	Value []JSONObjectValue
}

func (*JSONObjectPredict) CRC() uint32 {
	return 0x99c1d49d
}
func (*JSONObjectPredict) _JSONValue() {}

type KeyboardButton interface {
	tl.Object
	_KeyboardButton()
}

var (
	_ KeyboardButton = (*KeyboardButtonPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonURLPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonCallbackPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonRequestPhonePredict)(nil)
	_ KeyboardButton = (*KeyboardButtonRequestGeoLocationPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonSwitchInlinePredict)(nil)
	_ KeyboardButton = (*KeyboardButtonGamePredict)(nil)
	_ KeyboardButton = (*KeyboardButtonBuyPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonURLAuthPredict)(nil)
	_ KeyboardButton = (*InputKeyboardButtonURLAuthPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonRequestPollPredict)(nil)
	_ KeyboardButton = (*InputKeyboardButtonUserProfilePredict)(nil)
	_ KeyboardButton = (*KeyboardButtonUserProfilePredict)(nil)
	_ KeyboardButton = (*KeyboardButtonWebViewPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonSimpleWebViewPredict)(nil)
	_ KeyboardButton = (*KeyboardButtonRequestPeerPredict)(nil)
	_ KeyboardButton = (*InputKeyboardButtonRequestPeerPredict)(nil)
)

type KeyboardButtonPredict struct {
	Text string
}

func (*KeyboardButtonPredict) CRC() uint32 {
	return 0xa2fa4880
}
func (*KeyboardButtonPredict) _KeyboardButton() {}

type KeyboardButtonURLPredict struct {
	Text string
	URL  string
}

func (*KeyboardButtonURLPredict) CRC() uint32 {
	return 0x258aff05
}
func (*KeyboardButtonURLPredict) _KeyboardButton() {}

type KeyboardButtonCallbackPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	RequiresPassword bool     `tl:",omitempty:flags:0,implicit"`
	Text             string
	Data             []byte
}

func (*KeyboardButtonCallbackPredict) CRC() uint32 {
	return 0x35bbdb6b
}
func (*KeyboardButtonCallbackPredict) _KeyboardButton() {}

type KeyboardButtonRequestPhonePredict struct {
	Text string
}

func (*KeyboardButtonRequestPhonePredict) CRC() uint32 {
	return 0xb16a6c29
}
func (*KeyboardButtonRequestPhonePredict) _KeyboardButton() {}

type KeyboardButtonRequestGeoLocationPredict struct {
	Text string
}

func (*KeyboardButtonRequestGeoLocationPredict) CRC() uint32 {
	return 0xfc796b3f
}
func (*KeyboardButtonRequestGeoLocationPredict) _KeyboardButton() {}

type KeyboardButtonSwitchInlinePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	SamePeer  bool     `tl:",omitempty:flags:0,implicit"`
	Text      string
	Query     string
	PeerTypes []InlineQueryPeerType `tl:",omitempty:flags:1"`
}

func (*KeyboardButtonSwitchInlinePredict) CRC() uint32 {
	return 0x93b9fbb5
}
func (*KeyboardButtonSwitchInlinePredict) _KeyboardButton() {}

type KeyboardButtonGamePredict struct {
	Text string
}

func (*KeyboardButtonGamePredict) CRC() uint32 {
	return 0x50f41ccf
}
func (*KeyboardButtonGamePredict) _KeyboardButton() {}

type KeyboardButtonBuyPredict struct {
	Text string
}

func (*KeyboardButtonBuyPredict) CRC() uint32 {
	return 0xafd93fbb
}
func (*KeyboardButtonBuyPredict) _KeyboardButton() {}

type KeyboardButtonURLAuthPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Text     string
	FwdText  *string `tl:",omitempty:flags:0"`
	URL      string
	ButtonID int32
}

func (*KeyboardButtonURLAuthPredict) CRC() uint32 {
	return 0x10b78d29
}
func (*KeyboardButtonURLAuthPredict) _KeyboardButton() {}

type InputKeyboardButtonURLAuthPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	RequestWriteAccess bool     `tl:",omitempty:flags:0,implicit"`
	Text               string
	FwdText            *string `tl:",omitempty:flags:1"`
	URL                string
	Bot                InputUser
}

func (*InputKeyboardButtonURLAuthPredict) CRC() uint32 {
	return 0xd02e7fd4
}
func (*InputKeyboardButtonURLAuthPredict) _KeyboardButton() {}

type KeyboardButtonRequestPollPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Quiz *bool    `tl:",omitempty:flags:0"`
	Text string
}

func (*KeyboardButtonRequestPollPredict) CRC() uint32 {
	return 0xbbc7515d
}
func (*KeyboardButtonRequestPollPredict) _KeyboardButton() {}

type InputKeyboardButtonUserProfilePredict struct {
	Text   string
	UserID InputUser
}

func (*InputKeyboardButtonUserProfilePredict) CRC() uint32 {
	return 0xe988037b
}
func (*InputKeyboardButtonUserProfilePredict) _KeyboardButton() {}

type KeyboardButtonUserProfilePredict struct {
	Text   string
	UserID int64
}

func (*KeyboardButtonUserProfilePredict) CRC() uint32 {
	return 0x308660c1
}
func (*KeyboardButtonUserProfilePredict) _KeyboardButton() {}

type KeyboardButtonWebViewPredict struct {
	Text string
	URL  string
}

func (*KeyboardButtonWebViewPredict) CRC() uint32 {
	return 0x13767230
}
func (*KeyboardButtonWebViewPredict) _KeyboardButton() {}

type KeyboardButtonSimpleWebViewPredict struct {
	Text string
	URL  string
}

func (*KeyboardButtonSimpleWebViewPredict) CRC() uint32 {
	return 0xa0c0505c
}
func (*KeyboardButtonSimpleWebViewPredict) _KeyboardButton() {}

type KeyboardButtonRequestPeerPredict struct {
	Text        string
	ButtonID    int32
	PeerType    RequestPeerType
	MaxQuantity int32
}

func (*KeyboardButtonRequestPeerPredict) CRC() uint32 {
	return 0x53d7bfd8
}
func (*KeyboardButtonRequestPeerPredict) _KeyboardButton() {}

type InputKeyboardButtonRequestPeerPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	NameRequested     bool     `tl:",omitempty:flags:0,implicit"`
	UsernameRequested bool     `tl:",omitempty:flags:1,implicit"`
	PhotoRequested    bool     `tl:",omitempty:flags:2,implicit"`
	Text              string
	ButtonID          int32
	PeerType          RequestPeerType
	MaxQuantity       int32
}

func (*InputKeyboardButtonRequestPeerPredict) CRC() uint32 {
	return 0xc9662d05
}
func (*InputKeyboardButtonRequestPeerPredict) _KeyboardButton() {}

type KeyboardButtonRow interface {
	tl.Object
	_KeyboardButtonRow()
}

var (
	_ KeyboardButtonRow = (*KeyboardButtonRowPredict)(nil)
)

type KeyboardButtonRowPredict struct {
	Buttons []KeyboardButton
}

func (*KeyboardButtonRowPredict) CRC() uint32 {
	return 0x77608b83
}
func (*KeyboardButtonRowPredict) _KeyboardButtonRow() {}

type LabeledPrice interface {
	tl.Object
	_LabeledPrice()
}

var (
	_ LabeledPrice = (*LabeledPricePredict)(nil)
)

type LabeledPricePredict struct {
	Label  string
	Amount int64
}

func (*LabeledPricePredict) CRC() uint32 {
	return 0xcb296bf8
}
func (*LabeledPricePredict) _LabeledPrice() {}

type LangPackDifference interface {
	tl.Object
	_LangPackDifference()
}

var (
	_ LangPackDifference = (*LangPackDifferencePredict)(nil)
)

type LangPackDifferencePredict struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Strings     []LangPackString
}

func (*LangPackDifferencePredict) CRC() uint32 {
	return 0xf385c1f6
}
func (*LangPackDifferencePredict) _LangPackDifference() {}

type LangPackLanguage interface {
	tl.Object
	_LangPackLanguage()
}

var (
	_ LangPackLanguage = (*LangPackLanguagePredict)(nil)
)

type LangPackLanguagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Official        bool     `tl:",omitempty:flags:0,implicit"`
	Rtl             bool     `tl:",omitempty:flags:2,implicit"`
	Beta            bool     `tl:",omitempty:flags:3,implicit"`
	Name            string
	NativeName      string
	LangCode        string
	BaseLangCode    *string `tl:",omitempty:flags:1"`
	PluralCode      string
	StringsCount    int32
	TranslatedCount int32
	TranslationsURL string
}

func (*LangPackLanguagePredict) CRC() uint32 {
	return 0xeeca5ce3
}
func (*LangPackLanguagePredict) _LangPackLanguage() {}

type LangPackString interface {
	tl.Object
	_LangPackString()
}

var (
	_ LangPackString = (*LangPackStringPredict)(nil)
	_ LangPackString = (*LangPackStringPluralizedPredict)(nil)
	_ LangPackString = (*LangPackStringDeletedPredict)(nil)
)

type LangPackStringPredict struct {
	Key   string
	Value string
}

func (*LangPackStringPredict) CRC() uint32 {
	return 0xcad181f6
}
func (*LangPackStringPredict) _LangPackString() {}

type LangPackStringPluralizedPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Key        string
	ZeroValue  *string `tl:",omitempty:flags:0"`
	OneValue   *string `tl:",omitempty:flags:1"`
	TwoValue   *string `tl:",omitempty:flags:2"`
	FewValue   *string `tl:",omitempty:flags:3"`
	ManyValue  *string `tl:",omitempty:flags:4"`
	OtherValue string
}

func (*LangPackStringPluralizedPredict) CRC() uint32 {
	return 0x6c47ac9f
}
func (*LangPackStringPluralizedPredict) _LangPackString() {}

type LangPackStringDeletedPredict struct {
	Key string
}

func (*LangPackStringDeletedPredict) CRC() uint32 {
	return 0x2979eeb2
}
func (*LangPackStringDeletedPredict) _LangPackString() {}

type MaskCoords interface {
	tl.Object
	_MaskCoords()
}

var (
	_ MaskCoords = (*MaskCoordsPredict)(nil)
)

type MaskCoordsPredict struct {
	N    int32
	X    float64
	Y    float64
	Zoom float64
}

func (*MaskCoordsPredict) CRC() uint32 {
	return 0xaed6dbb2
}
func (*MaskCoordsPredict) _MaskCoords() {}

type MediaArea interface {
	tl.Object
	_MediaArea()
}

var (
	_ MediaArea = (*MediaAreaVenuePredict)(nil)
	_ MediaArea = (*InputMediaAreaVenuePredict)(nil)
	_ MediaArea = (*MediaAreaGeoPointPredict)(nil)
	_ MediaArea = (*MediaAreaSuggestedReactionPredict)(nil)
	_ MediaArea = (*MediaAreaChannelPostPredict)(nil)
	_ MediaArea = (*InputMediaAreaChannelPostPredict)(nil)
	_ MediaArea = (*MediaAreaURLPredict)(nil)
	_ MediaArea = (*MediaAreaWeatherPredict)(nil)
)

type MediaAreaVenuePredict struct {
	Coordinates MediaAreaCoordinates
	Geo         GeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
}

func (*MediaAreaVenuePredict) CRC() uint32 {
	return 0xbe82db9c
}
func (*MediaAreaVenuePredict) _MediaArea() {}

type InputMediaAreaVenuePredict struct {
	Coordinates MediaAreaCoordinates
	QueryID     int64
	ResultID    string
}

func (*InputMediaAreaVenuePredict) CRC() uint32 {
	return 0xb282217f
}
func (*InputMediaAreaVenuePredict) _MediaArea() {}

type MediaAreaGeoPointPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Coordinates MediaAreaCoordinates
	Geo         GeoPoint
	Address     GeoPointAddress `tl:",omitempty:flags:0"`
}

func (*MediaAreaGeoPointPredict) CRC() uint32 {
	return 0xcad5452d
}
func (*MediaAreaGeoPointPredict) _MediaArea() {}

type MediaAreaSuggestedReactionPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Dark        bool     `tl:",omitempty:flags:0,implicit"`
	Flipped     bool     `tl:",omitempty:flags:1,implicit"`
	Coordinates MediaAreaCoordinates
	Reaction    Reaction
}

func (*MediaAreaSuggestedReactionPredict) CRC() uint32 {
	return 0x14455871
}
func (*MediaAreaSuggestedReactionPredict) _MediaArea() {}

type MediaAreaChannelPostPredict struct {
	Coordinates MediaAreaCoordinates
	ChannelID   int64
	MsgID       int32
}

func (*MediaAreaChannelPostPredict) CRC() uint32 {
	return 0x770416af
}
func (*MediaAreaChannelPostPredict) _MediaArea() {}

type InputMediaAreaChannelPostPredict struct {
	Coordinates MediaAreaCoordinates
	Channel     InputChannel
	MsgID       int32
}

func (*InputMediaAreaChannelPostPredict) CRC() uint32 {
	return 0x2271f2bf
}
func (*InputMediaAreaChannelPostPredict) _MediaArea() {}

type MediaAreaURLPredict struct {
	Coordinates MediaAreaCoordinates
	URL         string
}

func (*MediaAreaURLPredict) CRC() uint32 {
	return 0x37381085
}
func (*MediaAreaURLPredict) _MediaArea() {}

type MediaAreaWeatherPredict struct {
	Coordinates  MediaAreaCoordinates
	Emoji        string
	TemperatureC float64
	Color        int32
}

func (*MediaAreaWeatherPredict) CRC() uint32 {
	return 0x49a6549c
}
func (*MediaAreaWeatherPredict) _MediaArea() {}

type MediaAreaCoordinates interface {
	tl.Object
	_MediaAreaCoordinates()
}

var (
	_ MediaAreaCoordinates = (*MediaAreaCoordinatesPredict)(nil)
)

type MediaAreaCoordinatesPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	X        float64
	Y        float64
	W        float64
	H        float64
	Rotation float64
	Radius   *float64 `tl:",omitempty:flags:0"`
}

func (*MediaAreaCoordinatesPredict) CRC() uint32 {
	return 0xcfc9e002
}
func (*MediaAreaCoordinatesPredict) _MediaAreaCoordinates() {}

type Message interface {
	tl.Object
	_Message()
}

var (
	_ Message = (*MessageEmptyPredict)(nil)
	_ Message = (*MessagePredict)(nil)
	_ Message = (*MessageServicePredict)(nil)
)

type MessageEmptyPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	ID     int32
	PeerID Peer `tl:",omitempty:flags:0"`
}

func (*MessageEmptyPredict) CRC() uint32 {
	return 0x90a6ca84
}
func (*MessageEmptyPredict) _Message() {}

type MessagePredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Out                  bool     `tl:",omitempty:flags:1,implicit"`
	Mentioned            bool     `tl:",omitempty:flags:4,implicit"`
	MediaUnread          bool     `tl:",omitempty:flags:5,implicit"`
	Silent               bool     `tl:",omitempty:flags:13,implicit"`
	Post                 bool     `tl:",omitempty:flags:14,implicit"`
	FromScheduled        bool     `tl:",omitempty:flags:18,implicit"`
	Legacy               bool     `tl:",omitempty:flags:19,implicit"`
	EditHide             bool     `tl:",omitempty:flags:21,implicit"`
	Pinned               bool     `tl:",omitempty:flags:24,implicit"`
	Noforwards           bool     `tl:",omitempty:flags:26,implicit"`
	InvertMedia          bool     `tl:",omitempty:flags:27,implicit"`
	_                    struct{} `tl:"flags2,bitflag"`
	Offline              bool     `tl:",omitempty:flags2:1,implicit"`
	ID                   int32
	FromID               Peer   `tl:",omitempty:flags:8"`
	FromBoostsApplied    *int32 `tl:",omitempty:flags:29"`
	PeerID               Peer
	SavedPeerID          Peer               `tl:",omitempty:flags:28"`
	FwdFrom              MessageFwdHeader   `tl:",omitempty:flags:2"`
	ViaBotID             *int64             `tl:",omitempty:flags:11"`
	ViaBusinessBotID     *int64             `tl:",omitempty:flags2:0"`
	ReplyTo              MessageReplyHeader `tl:",omitempty:flags:3"`
	Date                 int32
	Message              string
	Media                MessageMedia        `tl:",omitempty:flags:9"`
	ReplyMarkup          ReplyMarkup         `tl:",omitempty:flags:6"`
	Entities             []MessageEntity     `tl:",omitempty:flags:7"`
	Views                *int32              `tl:",omitempty:flags:10"`
	Forwards             *int32              `tl:",omitempty:flags:10"`
	Replies              MessageReplies      `tl:",omitempty:flags:23"`
	EditDate             *int32              `tl:",omitempty:flags:15"`
	PostAuthor           *string             `tl:",omitempty:flags:16"`
	GroupedID            *int64              `tl:",omitempty:flags:17"`
	Reactions            MessageReactions    `tl:",omitempty:flags:20"`
	RestrictionReason    []RestrictionReason `tl:",omitempty:flags:22"`
	TTLPeriod            *int32              `tl:",omitempty:flags:25"`
	QuickReplyShortcutID *int32              `tl:",omitempty:flags:30"`
	Effect               *int64              `tl:",omitempty:flags2:2"`
	Factcheck            FactCheck           `tl:",omitempty:flags2:3"`
}

func (*MessagePredict) CRC() uint32 {
	return 0x94345242
}
func (*MessagePredict) _Message() {}

type MessageServicePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Out         bool     `tl:",omitempty:flags:1,implicit"`
	Mentioned   bool     `tl:",omitempty:flags:4,implicit"`
	MediaUnread bool     `tl:",omitempty:flags:5,implicit"`
	Silent      bool     `tl:",omitempty:flags:13,implicit"`
	Post        bool     `tl:",omitempty:flags:14,implicit"`
	Legacy      bool     `tl:",omitempty:flags:19,implicit"`
	ID          int32
	FromID      Peer `tl:",omitempty:flags:8"`
	PeerID      Peer
	ReplyTo     MessageReplyHeader `tl:",omitempty:flags:3"`
	Date        int32
	Action      MessageAction
	TTLPeriod   *int32 `tl:",omitempty:flags:25"`
}

func (*MessageServicePredict) CRC() uint32 {
	return 0x2b085862
}
func (*MessageServicePredict) _Message() {}

type MessageAction interface {
	tl.Object
	_MessageAction()
}

var (
	_ MessageAction = (*MessageActionEmptyPredict)(nil)
	_ MessageAction = (*MessageActionChatCreatePredict)(nil)
	_ MessageAction = (*MessageActionChatEditTitlePredict)(nil)
	_ MessageAction = (*MessageActionChatEditPhotoPredict)(nil)
	_ MessageAction = (*MessageActionChatDeletePhotoPredict)(nil)
	_ MessageAction = (*MessageActionChatAddUserPredict)(nil)
	_ MessageAction = (*MessageActionChatDeleteUserPredict)(nil)
	_ MessageAction = (*MessageActionChatJoinedByLinkPredict)(nil)
	_ MessageAction = (*MessageActionChannelCreatePredict)(nil)
	_ MessageAction = (*MessageActionChatMigrateToPredict)(nil)
	_ MessageAction = (*MessageActionChannelMigrateFromPredict)(nil)
	_ MessageAction = (*MessageActionPinMessagePredict)(nil)
	_ MessageAction = (*MessageActionHistoryClearPredict)(nil)
	_ MessageAction = (*MessageActionGameScorePredict)(nil)
	_ MessageAction = (*MessageActionPaymentSentMePredict)(nil)
	_ MessageAction = (*MessageActionPaymentSentPredict)(nil)
	_ MessageAction = (*MessageActionPhoneCallPredict)(nil)
	_ MessageAction = (*MessageActionScreenshotTakenPredict)(nil)
	_ MessageAction = (*MessageActionCustomActionPredict)(nil)
	_ MessageAction = (*MessageActionBotAllowedPredict)(nil)
	_ MessageAction = (*MessageActionSecureValuesSentMePredict)(nil)
	_ MessageAction = (*MessageActionSecureValuesSentPredict)(nil)
	_ MessageAction = (*MessageActionContactSignUpPredict)(nil)
	_ MessageAction = (*MessageActionGeoProximityReachedPredict)(nil)
	_ MessageAction = (*MessageActionGroupCallPredict)(nil)
	_ MessageAction = (*MessageActionInviteToGroupCallPredict)(nil)
	_ MessageAction = (*MessageActionSetMessagesTTLPredict)(nil)
	_ MessageAction = (*MessageActionGroupCallScheduledPredict)(nil)
	_ MessageAction = (*MessageActionSetChatThemePredict)(nil)
	_ MessageAction = (*MessageActionChatJoinedByRequestPredict)(nil)
	_ MessageAction = (*MessageActionWebViewDataSentMePredict)(nil)
	_ MessageAction = (*MessageActionWebViewDataSentPredict)(nil)
	_ MessageAction = (*MessageActionGiftPremiumPredict)(nil)
	_ MessageAction = (*MessageActionTopicCreatePredict)(nil)
	_ MessageAction = (*MessageActionTopicEditPredict)(nil)
	_ MessageAction = (*MessageActionSuggestProfilePhotoPredict)(nil)
	_ MessageAction = (*MessageActionRequestedPeerPredict)(nil)
	_ MessageAction = (*MessageActionSetChatWallPaperPredict)(nil)
	_ MessageAction = (*MessageActionGiftCodePredict)(nil)
	_ MessageAction = (*MessageActionGiveawayLaunchPredict)(nil)
	_ MessageAction = (*MessageActionGiveawayResultsPredict)(nil)
	_ MessageAction = (*MessageActionBoostApplyPredict)(nil)
	_ MessageAction = (*MessageActionRequestedPeerSentMePredict)(nil)
	_ MessageAction = (*MessageActionPaymentRefundedPredict)(nil)
	_ MessageAction = (*MessageActionGiftStarsPredict)(nil)
)

type MessageActionEmptyPredict struct{}

func (*MessageActionEmptyPredict) CRC() uint32 {
	return 0xb6aef7b0
}
func (*MessageActionEmptyPredict) _MessageAction() {}

type MessageActionChatCreatePredict struct {
	Title string
	Users []int64
}

func (*MessageActionChatCreatePredict) CRC() uint32 {
	return 0xbd47cbad
}
func (*MessageActionChatCreatePredict) _MessageAction() {}

type MessageActionChatEditTitlePredict struct {
	Title string
}

func (*MessageActionChatEditTitlePredict) CRC() uint32 {
	return 0xb5a1ce5a
}
func (*MessageActionChatEditTitlePredict) _MessageAction() {}

type MessageActionChatEditPhotoPredict struct {
	Photo Photo
}

func (*MessageActionChatEditPhotoPredict) CRC() uint32 {
	return 0x7fcb13a8
}
func (*MessageActionChatEditPhotoPredict) _MessageAction() {}

type MessageActionChatDeletePhotoPredict struct{}

func (*MessageActionChatDeletePhotoPredict) CRC() uint32 {
	return 0x95e3fbef
}
func (*MessageActionChatDeletePhotoPredict) _MessageAction() {}

type MessageActionChatAddUserPredict struct {
	Users []int64
}

func (*MessageActionChatAddUserPredict) CRC() uint32 {
	return 0x15cefd00
}
func (*MessageActionChatAddUserPredict) _MessageAction() {}

type MessageActionChatDeleteUserPredict struct {
	UserID int64
}

func (*MessageActionChatDeleteUserPredict) CRC() uint32 {
	return 0xa43f30cc
}
func (*MessageActionChatDeleteUserPredict) _MessageAction() {}

type MessageActionChatJoinedByLinkPredict struct {
	InviterID int64
}

func (*MessageActionChatJoinedByLinkPredict) CRC() uint32 {
	return 0x31224c3
}
func (*MessageActionChatJoinedByLinkPredict) _MessageAction() {}

type MessageActionChannelCreatePredict struct {
	Title string
}

func (*MessageActionChannelCreatePredict) CRC() uint32 {
	return 0x95d2ac92
}
func (*MessageActionChannelCreatePredict) _MessageAction() {}

type MessageActionChatMigrateToPredict struct {
	ChannelID int64
}

func (*MessageActionChatMigrateToPredict) CRC() uint32 {
	return 0xe1037f92
}
func (*MessageActionChatMigrateToPredict) _MessageAction() {}

type MessageActionChannelMigrateFromPredict struct {
	Title  string
	ChatID int64
}

func (*MessageActionChannelMigrateFromPredict) CRC() uint32 {
	return 0xea3948e9
}
func (*MessageActionChannelMigrateFromPredict) _MessageAction() {}

type MessageActionPinMessagePredict struct{}

func (*MessageActionPinMessagePredict) CRC() uint32 {
	return 0x94bd38ed
}
func (*MessageActionPinMessagePredict) _MessageAction() {}

type MessageActionHistoryClearPredict struct{}

func (*MessageActionHistoryClearPredict) CRC() uint32 {
	return 0x9fbab604
}
func (*MessageActionHistoryClearPredict) _MessageAction() {}

type MessageActionGameScorePredict struct {
	GameID int64
	Score  int32
}

func (*MessageActionGameScorePredict) CRC() uint32 {
	return 0x92a72876
}
func (*MessageActionGameScorePredict) _MessageAction() {}

type MessageActionPaymentSentMePredict struct {
	_                struct{} `tl:"flags,bitflag"`
	RecurringInit    bool     `tl:",omitempty:flags:2,implicit"`
	RecurringUsed    bool     `tl:",omitempty:flags:3,implicit"`
	Currency         string
	TotalAmount      int64
	Payload          []byte
	Info             PaymentRequestedInfo `tl:",omitempty:flags:0"`
	ShippingOptionID *string              `tl:",omitempty:flags:1"`
	Charge           PaymentCharge
}

func (*MessageActionPaymentSentMePredict) CRC() uint32 {
	return 0x8f31b327
}
func (*MessageActionPaymentSentMePredict) _MessageAction() {}

type MessageActionPaymentSentPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	RecurringInit bool     `tl:",omitempty:flags:2,implicit"`
	RecurringUsed bool     `tl:",omitempty:flags:3,implicit"`
	Currency      string
	TotalAmount   int64
	InvoiceSlug   *string `tl:",omitempty:flags:0"`
}

func (*MessageActionPaymentSentPredict) CRC() uint32 {
	return 0x96163f56
}
func (*MessageActionPaymentSentPredict) _MessageAction() {}

type MessageActionPhoneCallPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Video    bool     `tl:",omitempty:flags:2,implicit"`
	CallID   int64
	Reason   PhoneCallDiscardReason `tl:",omitempty:flags:0"`
	Duration *int32                 `tl:",omitempty:flags:1"`
}

func (*MessageActionPhoneCallPredict) CRC() uint32 {
	return 0x80e11a7f
}
func (*MessageActionPhoneCallPredict) _MessageAction() {}

type MessageActionScreenshotTakenPredict struct{}

func (*MessageActionScreenshotTakenPredict) CRC() uint32 {
	return 0x4792929b
}
func (*MessageActionScreenshotTakenPredict) _MessageAction() {}

type MessageActionCustomActionPredict struct {
	Message string
}

func (*MessageActionCustomActionPredict) CRC() uint32 {
	return 0xfae69f56
}
func (*MessageActionCustomActionPredict) _MessageAction() {}

type MessageActionBotAllowedPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	AttachMenu  bool     `tl:",omitempty:flags:1,implicit"`
	FromRequest bool     `tl:",omitempty:flags:3,implicit"`
	Domain      *string  `tl:",omitempty:flags:0"`
	App         BotApp   `tl:",omitempty:flags:2"`
}

func (*MessageActionBotAllowedPredict) CRC() uint32 {
	return 0xc516d679
}
func (*MessageActionBotAllowedPredict) _MessageAction() {}

type MessageActionSecureValuesSentMePredict struct {
	Values      []SecureValue
	Credentials SecureCredentialsEncrypted
}

func (*MessageActionSecureValuesSentMePredict) CRC() uint32 {
	return 0x1b287353
}
func (*MessageActionSecureValuesSentMePredict) _MessageAction() {}

type MessageActionSecureValuesSentPredict struct {
	Types []SecureValueType
}

func (*MessageActionSecureValuesSentPredict) CRC() uint32 {
	return 0xd95c6154
}
func (*MessageActionSecureValuesSentPredict) _MessageAction() {}

type MessageActionContactSignUpPredict struct{}

func (*MessageActionContactSignUpPredict) CRC() uint32 {
	return 0xf3f25f76
}
func (*MessageActionContactSignUpPredict) _MessageAction() {}

type MessageActionGeoProximityReachedPredict struct {
	FromID   Peer
	ToID     Peer
	Distance int32
}

func (*MessageActionGeoProximityReachedPredict) CRC() uint32 {
	return 0x98e0d697
}
func (*MessageActionGeoProximityReachedPredict) _MessageAction() {}

type MessageActionGroupCallPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Call     InputGroupCall
	Duration *int32 `tl:",omitempty:flags:0"`
}

func (*MessageActionGroupCallPredict) CRC() uint32 {
	return 0x7a0d7f42
}
func (*MessageActionGroupCallPredict) _MessageAction() {}

type MessageActionInviteToGroupCallPredict struct {
	Call  InputGroupCall
	Users []int64
}

func (*MessageActionInviteToGroupCallPredict) CRC() uint32 {
	return 0x502f92f7
}
func (*MessageActionInviteToGroupCallPredict) _MessageAction() {}

type MessageActionSetMessagesTTLPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Period          int32
	AutoSettingFrom *int64 `tl:",omitempty:flags:0"`
}

func (*MessageActionSetMessagesTTLPredict) CRC() uint32 {
	return 0x3c134d7b
}
func (*MessageActionSetMessagesTTLPredict) _MessageAction() {}

type MessageActionGroupCallScheduledPredict struct {
	Call         InputGroupCall
	ScheduleDate int32
}

func (*MessageActionGroupCallScheduledPredict) CRC() uint32 {
	return 0xb3a07661
}
func (*MessageActionGroupCallScheduledPredict) _MessageAction() {}

type MessageActionSetChatThemePredict struct {
	Emoticon string
}

func (*MessageActionSetChatThemePredict) CRC() uint32 {
	return 0xaa786345
}
func (*MessageActionSetChatThemePredict) _MessageAction() {}

type MessageActionChatJoinedByRequestPredict struct{}

func (*MessageActionChatJoinedByRequestPredict) CRC() uint32 {
	return 0xebbca3cb
}
func (*MessageActionChatJoinedByRequestPredict) _MessageAction() {}

type MessageActionWebViewDataSentMePredict struct {
	Text string
	Data string
}

func (*MessageActionWebViewDataSentMePredict) CRC() uint32 {
	return 0x47dd8079
}
func (*MessageActionWebViewDataSentMePredict) _MessageAction() {}

type MessageActionWebViewDataSentPredict struct {
	Text string
}

func (*MessageActionWebViewDataSentPredict) CRC() uint32 {
	return 0xb4c38cb5
}
func (*MessageActionWebViewDataSentPredict) _MessageAction() {}

type MessageActionGiftPremiumPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Currency       string
	Amount         int64
	Months         int32
	CryptoCurrency *string `tl:",omitempty:flags:0"`
	CryptoAmount   *int64  `tl:",omitempty:flags:0"`
}

func (*MessageActionGiftPremiumPredict) CRC() uint32 {
	return 0xc83d6aec
}
func (*MessageActionGiftPremiumPredict) _MessageAction() {}

type MessageActionTopicCreatePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       string
	IconColor   int32
	IconEmojiID *int64 `tl:",omitempty:flags:0"`
}

func (*MessageActionTopicCreatePredict) CRC() uint32 {
	return 0xd999256
}
func (*MessageActionTopicCreatePredict) _MessageAction() {}

type MessageActionTopicEditPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       *string  `tl:",omitempty:flags:0"`
	IconEmojiID *int64   `tl:",omitempty:flags:1"`
	Closed      *bool    `tl:",omitempty:flags:2"`
	Hidden      *bool    `tl:",omitempty:flags:3"`
}

func (*MessageActionTopicEditPredict) CRC() uint32 {
	return 0xc0944820
}
func (*MessageActionTopicEditPredict) _MessageAction() {}

type MessageActionSuggestProfilePhotoPredict struct {
	Photo Photo
}

func (*MessageActionSuggestProfilePhotoPredict) CRC() uint32 {
	return 0x57de635e
}
func (*MessageActionSuggestProfilePhotoPredict) _MessageAction() {}

type MessageActionRequestedPeerPredict struct {
	ButtonID int32
	Peers    []Peer
}

func (*MessageActionRequestedPeerPredict) CRC() uint32 {
	return 0x31518e9b
}
func (*MessageActionRequestedPeerPredict) _MessageAction() {}

type MessageActionSetChatWallPaperPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Same      bool     `tl:",omitempty:flags:0,implicit"`
	ForBoth   bool     `tl:",omitempty:flags:1,implicit"`
	Wallpaper WallPaper
}

func (*MessageActionSetChatWallPaperPredict) CRC() uint32 {
	return 0x5060a3f4
}
func (*MessageActionSetChatWallPaperPredict) _MessageAction() {}

type MessageActionGiftCodePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ViaGiveaway    bool     `tl:",omitempty:flags:0,implicit"`
	Unclaimed      bool     `tl:",omitempty:flags:2,implicit"`
	BoostPeer      Peer     `tl:",omitempty:flags:1"`
	Months         int32
	Slug           string
	Currency       *string `tl:",omitempty:flags:2"`
	Amount         *int64  `tl:",omitempty:flags:2"`
	CryptoCurrency *string `tl:",omitempty:flags:3"`
	CryptoAmount   *int64  `tl:",omitempty:flags:3"`
}

func (*MessageActionGiftCodePredict) CRC() uint32 {
	return 0x678c2e09
}
func (*MessageActionGiftCodePredict) _MessageAction() {}

type MessageActionGiveawayLaunchPredict struct{}

func (*MessageActionGiveawayLaunchPredict) CRC() uint32 {
	return 0x332ba9ed
}
func (*MessageActionGiveawayLaunchPredict) _MessageAction() {}

type MessageActionGiveawayResultsPredict struct {
	WinnersCount   int32
	UnclaimedCount int32
}

func (*MessageActionGiveawayResultsPredict) CRC() uint32 {
	return 0x2a9fadc5
}
func (*MessageActionGiveawayResultsPredict) _MessageAction() {}

type MessageActionBoostApplyPredict struct {
	Boosts int32
}

func (*MessageActionBoostApplyPredict) CRC() uint32 {
	return 0xcc02aa6d
}
func (*MessageActionBoostApplyPredict) _MessageAction() {}

type MessageActionRequestedPeerSentMePredict struct {
	ButtonID int32
	Peers    []RequestedPeer
}

func (*MessageActionRequestedPeerSentMePredict) CRC() uint32 {
	return 0x93b31848
}
func (*MessageActionRequestedPeerSentMePredict) _MessageAction() {}

type MessageActionPaymentRefundedPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Peer        Peer
	Currency    string
	TotalAmount int64
	Payload     *[]byte `tl:",omitempty:flags:0"`
	Charge      PaymentCharge
}

func (*MessageActionPaymentRefundedPredict) CRC() uint32 {
	return 0x41b3e202
}
func (*MessageActionPaymentRefundedPredict) _MessageAction() {}

type MessageActionGiftStarsPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Currency       string
	Amount         int64
	Stars          int64
	CryptoCurrency *string `tl:",omitempty:flags:0"`
	CryptoAmount   *int64  `tl:",omitempty:flags:0"`
	TransactionID  *string `tl:",omitempty:flags:1"`
}

func (*MessageActionGiftStarsPredict) CRC() uint32 {
	return 0x45d5b021
}
func (*MessageActionGiftStarsPredict) _MessageAction() {}

type MessageEntity interface {
	tl.Object
	_MessageEntity()
}

var (
	_ MessageEntity = (*MessageEntityUnknownPredict)(nil)
	_ MessageEntity = (*MessageEntityMentionPredict)(nil)
	_ MessageEntity = (*MessageEntityHashtagPredict)(nil)
	_ MessageEntity = (*MessageEntityBotCommandPredict)(nil)
	_ MessageEntity = (*MessageEntityURLPredict)(nil)
	_ MessageEntity = (*MessageEntityEmailPredict)(nil)
	_ MessageEntity = (*MessageEntityBoldPredict)(nil)
	_ MessageEntity = (*MessageEntityItalicPredict)(nil)
	_ MessageEntity = (*MessageEntityCodePredict)(nil)
	_ MessageEntity = (*MessageEntityPrePredict)(nil)
	_ MessageEntity = (*MessageEntityTextURLPredict)(nil)
	_ MessageEntity = (*MessageEntityMentionNamePredict)(nil)
	_ MessageEntity = (*InputMessageEntityMentionNamePredict)(nil)
	_ MessageEntity = (*MessageEntityPhonePredict)(nil)
	_ MessageEntity = (*MessageEntityCashtagPredict)(nil)
	_ MessageEntity = (*MessageEntityUnderlinePredict)(nil)
	_ MessageEntity = (*MessageEntityStrikePredict)(nil)
	_ MessageEntity = (*MessageEntityBankCardPredict)(nil)
	_ MessageEntity = (*MessageEntitySpoilerPredict)(nil)
	_ MessageEntity = (*MessageEntityCustomEmojiPredict)(nil)
	_ MessageEntity = (*MessageEntityBlockquotePredict)(nil)
)

type MessageEntityUnknownPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityUnknownPredict) CRC() uint32 {
	return 0xbb92ba95
}
func (*MessageEntityUnknownPredict) _MessageEntity() {}

type MessageEntityMentionPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityMentionPredict) CRC() uint32 {
	return 0xfa04579d
}
func (*MessageEntityMentionPredict) _MessageEntity() {}

type MessageEntityHashtagPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityHashtagPredict) CRC() uint32 {
	return 0x6f635b0d
}
func (*MessageEntityHashtagPredict) _MessageEntity() {}

type MessageEntityBotCommandPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityBotCommandPredict) CRC() uint32 {
	return 0x6cef8ac7
}
func (*MessageEntityBotCommandPredict) _MessageEntity() {}

type MessageEntityURLPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityURLPredict) CRC() uint32 {
	return 0x6ed02538
}
func (*MessageEntityURLPredict) _MessageEntity() {}

type MessageEntityEmailPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityEmailPredict) CRC() uint32 {
	return 0x64e475c2
}
func (*MessageEntityEmailPredict) _MessageEntity() {}

type MessageEntityBoldPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityBoldPredict) CRC() uint32 {
	return 0xbd610bc9
}
func (*MessageEntityBoldPredict) _MessageEntity() {}

type MessageEntityItalicPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityItalicPredict) CRC() uint32 {
	return 0x826f8b60
}
func (*MessageEntityItalicPredict) _MessageEntity() {}

type MessageEntityCodePredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityCodePredict) CRC() uint32 {
	return 0x28a20571
}
func (*MessageEntityCodePredict) _MessageEntity() {}

type MessageEntityPrePredict struct {
	Offset   int32
	Length   int32
	Language string
}

func (*MessageEntityPrePredict) CRC() uint32 {
	return 0x73924be0
}
func (*MessageEntityPrePredict) _MessageEntity() {}

type MessageEntityTextURLPredict struct {
	Offset int32
	Length int32
	URL    string
}

func (*MessageEntityTextURLPredict) CRC() uint32 {
	return 0x76a6d327
}
func (*MessageEntityTextURLPredict) _MessageEntity() {}

type MessageEntityMentionNamePredict struct {
	Offset int32
	Length int32
	UserID int64
}

func (*MessageEntityMentionNamePredict) CRC() uint32 {
	return 0xdc7b1140
}
func (*MessageEntityMentionNamePredict) _MessageEntity() {}

type InputMessageEntityMentionNamePredict struct {
	Offset int32
	Length int32
	UserID InputUser
}

func (*InputMessageEntityMentionNamePredict) CRC() uint32 {
	return 0x208e68c9
}
func (*InputMessageEntityMentionNamePredict) _MessageEntity() {}

type MessageEntityPhonePredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityPhonePredict) CRC() uint32 {
	return 0x9b69e34b
}
func (*MessageEntityPhonePredict) _MessageEntity() {}

type MessageEntityCashtagPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityCashtagPredict) CRC() uint32 {
	return 0x4c4e743f
}
func (*MessageEntityCashtagPredict) _MessageEntity() {}

type MessageEntityUnderlinePredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityUnderlinePredict) CRC() uint32 {
	return 0x9c4e7e8b
}
func (*MessageEntityUnderlinePredict) _MessageEntity() {}

type MessageEntityStrikePredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityStrikePredict) CRC() uint32 {
	return 0xbf0693d4
}
func (*MessageEntityStrikePredict) _MessageEntity() {}

type MessageEntityBankCardPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntityBankCardPredict) CRC() uint32 {
	return 0x761e6af4
}
func (*MessageEntityBankCardPredict) _MessageEntity() {}

type MessageEntitySpoilerPredict struct {
	Offset int32
	Length int32
}

func (*MessageEntitySpoilerPredict) CRC() uint32 {
	return 0x32ca960f
}
func (*MessageEntitySpoilerPredict) _MessageEntity() {}

type MessageEntityCustomEmojiPredict struct {
	Offset     int32
	Length     int32
	DocumentID int64
}

func (*MessageEntityCustomEmojiPredict) CRC() uint32 {
	return 0xc8cf05f8
}
func (*MessageEntityCustomEmojiPredict) _MessageEntity() {}

type MessageEntityBlockquotePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Collapsed bool     `tl:",omitempty:flags:0,implicit"`
	Offset    int32
	Length    int32
}

func (*MessageEntityBlockquotePredict) CRC() uint32 {
	return 0xf1ccaaac
}
func (*MessageEntityBlockquotePredict) _MessageEntity() {}

type MessageExtendedMedia interface {
	tl.Object
	_MessageExtendedMedia()
}

var (
	_ MessageExtendedMedia = (*MessageExtendedMediaPreviewPredict)(nil)
	_ MessageExtendedMedia = (*MessageExtendedMediaPredict)(nil)
)

type MessageExtendedMediaPreviewPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	W             *int32    `tl:",omitempty:flags:0"`
	H             *int32    `tl:",omitempty:flags:0"`
	Thumb         PhotoSize `tl:",omitempty:flags:1"`
	VideoDuration *int32    `tl:",omitempty:flags:2"`
}

func (*MessageExtendedMediaPreviewPredict) CRC() uint32 {
	return 0xad628cc8
}
func (*MessageExtendedMediaPreviewPredict) _MessageExtendedMedia() {}

type MessageExtendedMediaPredict struct {
	Media MessageMedia
}

func (*MessageExtendedMediaPredict) CRC() uint32 {
	return 0xee479c64
}
func (*MessageExtendedMediaPredict) _MessageExtendedMedia() {}

type MessageFwdHeader interface {
	tl.Object
	_MessageFwdHeader()
}

var (
	_ MessageFwdHeader = (*MessageFwdHeaderPredict)(nil)
)

type MessageFwdHeaderPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Imported       bool     `tl:",omitempty:flags:7,implicit"`
	SavedOut       bool     `tl:",omitempty:flags:11,implicit"`
	FromID         Peer     `tl:",omitempty:flags:0"`
	FromName       *string  `tl:",omitempty:flags:5"`
	Date           int32
	ChannelPost    *int32  `tl:",omitempty:flags:2"`
	PostAuthor     *string `tl:",omitempty:flags:3"`
	SavedFromPeer  Peer    `tl:",omitempty:flags:4"`
	SavedFromMsgID *int32  `tl:",omitempty:flags:4"`
	SavedFromID    Peer    `tl:",omitempty:flags:8"`
	SavedFromName  *string `tl:",omitempty:flags:9"`
	SavedDate      *int32  `tl:",omitempty:flags:10"`
	PsaType        *string `tl:",omitempty:flags:6"`
}

func (*MessageFwdHeaderPredict) CRC() uint32 {
	return 0x4e4df4bb
}
func (*MessageFwdHeaderPredict) _MessageFwdHeader() {}

type MessageMedia interface {
	tl.Object
	_MessageMedia()
}

var (
	_ MessageMedia = (*MessageMediaEmptyPredict)(nil)
	_ MessageMedia = (*MessageMediaPhotoPredict)(nil)
	_ MessageMedia = (*MessageMediaGeoPredict)(nil)
	_ MessageMedia = (*MessageMediaContactPredict)(nil)
	_ MessageMedia = (*MessageMediaUnsupportedPredict)(nil)
	_ MessageMedia = (*MessageMediaDocumentPredict)(nil)
	_ MessageMedia = (*MessageMediaWebPagePredict)(nil)
	_ MessageMedia = (*MessageMediaVenuePredict)(nil)
	_ MessageMedia = (*MessageMediaGamePredict)(nil)
	_ MessageMedia = (*MessageMediaInvoicePredict)(nil)
	_ MessageMedia = (*MessageMediaGeoLivePredict)(nil)
	_ MessageMedia = (*MessageMediaPollPredict)(nil)
	_ MessageMedia = (*MessageMediaDicePredict)(nil)
	_ MessageMedia = (*MessageMediaStoryPredict)(nil)
	_ MessageMedia = (*MessageMediaGiveawayPredict)(nil)
	_ MessageMedia = (*MessageMediaGiveawayResultsPredict)(nil)
	_ MessageMedia = (*MessageMediaPaidMediaPredict)(nil)
)

type MessageMediaEmptyPredict struct{}

func (*MessageMediaEmptyPredict) CRC() uint32 {
	return 0x3ded6320
}
func (*MessageMediaEmptyPredict) _MessageMedia() {}

type MessageMediaPhotoPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:",omitempty:flags:3,implicit"`
	Photo      Photo    `tl:",omitempty:flags:0"`
	TTLSeconds *int32   `tl:",omitempty:flags:2"`
}

func (*MessageMediaPhotoPredict) CRC() uint32 {
	return 0x695150d7
}
func (*MessageMediaPhotoPredict) _MessageMedia() {}

type MessageMediaGeoPredict struct {
	Geo GeoPoint
}

func (*MessageMediaGeoPredict) CRC() uint32 {
	return 0x56e0d474
}
func (*MessageMediaGeoPredict) _MessageMedia() {}

type MessageMediaContactPredict struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	UserID      int64
}

func (*MessageMediaContactPredict) CRC() uint32 {
	return 0x70322949
}
func (*MessageMediaContactPredict) _MessageMedia() {}

type MessageMediaUnsupportedPredict struct{}

func (*MessageMediaUnsupportedPredict) CRC() uint32 {
	return 0x9f84f49e
}
func (*MessageMediaUnsupportedPredict) _MessageMedia() {}

type MessageMediaDocumentPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Nopremium   bool     `tl:",omitempty:flags:3,implicit"`
	Spoiler     bool     `tl:",omitempty:flags:4,implicit"`
	Video       bool     `tl:",omitempty:flags:6,implicit"`
	Round       bool     `tl:",omitempty:flags:7,implicit"`
	Voice       bool     `tl:",omitempty:flags:8,implicit"`
	Document    Document `tl:",omitempty:flags:0"`
	AltDocument Document `tl:",omitempty:flags:5"`
	TTLSeconds  *int32   `tl:",omitempty:flags:2"`
}

func (*MessageMediaDocumentPredict) CRC() uint32 {
	return 0x4cf4d72d
}
func (*MessageMediaDocumentPredict) _MessageMedia() {}

type MessageMediaWebPagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ForceLargeMedia bool     `tl:",omitempty:flags:0,implicit"`
	ForceSmallMedia bool     `tl:",omitempty:flags:1,implicit"`
	Manual          bool     `tl:",omitempty:flags:3,implicit"`
	Safe            bool     `tl:",omitempty:flags:4,implicit"`
	Webpage         WebPage
}

func (*MessageMediaWebPagePredict) CRC() uint32 {
	return 0xddf10c3b
}
func (*MessageMediaWebPagePredict) _MessageMedia() {}

type MessageMediaVenuePredict struct {
	Geo       GeoPoint
	Title     string
	Address   string
	Provider  string
	VenueID   string
	VenueType string
}

func (*MessageMediaVenuePredict) CRC() uint32 {
	return 0x2ec0533f
}
func (*MessageMediaVenuePredict) _MessageMedia() {}

type MessageMediaGamePredict struct {
	Game Game
}

func (*MessageMediaGamePredict) CRC() uint32 {
	return 0xfdb19008
}
func (*MessageMediaGamePredict) _MessageMedia() {}

type MessageMediaInvoicePredict struct {
	_                        struct{} `tl:"flags,bitflag"`
	ShippingAddressRequested bool     `tl:",omitempty:flags:1,implicit"`
	Test                     bool     `tl:",omitempty:flags:3,implicit"`
	Title                    string
	Description              string
	Photo                    WebDocument `tl:",omitempty:flags:0"`
	ReceiptMsgID             *int32      `tl:",omitempty:flags:2"`
	Currency                 string
	TotalAmount              int64
	StartParam               string
	ExtendedMedia            MessageExtendedMedia `tl:",omitempty:flags:4"`
}

func (*MessageMediaInvoicePredict) CRC() uint32 {
	return 0xf6a548d3
}
func (*MessageMediaInvoicePredict) _MessageMedia() {}

type MessageMediaGeoLivePredict struct {
	_                           struct{} `tl:"flags,bitflag"`
	Geo                         GeoPoint
	Heading                     *int32 `tl:",omitempty:flags:0"`
	Period                      int32
	ProximityNotificationRadius *int32 `tl:",omitempty:flags:1"`
}

func (*MessageMediaGeoLivePredict) CRC() uint32 {
	return 0xb940c666
}
func (*MessageMediaGeoLivePredict) _MessageMedia() {}

type MessageMediaPollPredict struct {
	Poll    Poll
	Results PollResults
}

func (*MessageMediaPollPredict) CRC() uint32 {
	return 0x4bd6e798
}
func (*MessageMediaPollPredict) _MessageMedia() {}

type MessageMediaDicePredict struct {
	Value    int32
	Emoticon string
}

func (*MessageMediaDicePredict) CRC() uint32 {
	return 0x3f7ee58b
}
func (*MessageMediaDicePredict) _MessageMedia() {}

type MessageMediaStoryPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	ViaMention bool     `tl:",omitempty:flags:1,implicit"`
	Peer       Peer
	ID         int32
	Story      StoryItem `tl:",omitempty:flags:0"`
}

func (*MessageMediaStoryPredict) CRC() uint32 {
	return 0x68cb6283
}
func (*MessageMediaStoryPredict) _MessageMedia() {}

type MessageMediaGiveawayPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers bool     `tl:",omitempty:flags:0,implicit"`
	WinnersAreVisible  bool     `tl:",omitempty:flags:2,implicit"`
	Channels           []int64
	CountriesIso2      []string `tl:",omitempty:flags:1"`
	PrizeDescription   *string  `tl:",omitempty:flags:3"`
	Quantity           int32
	Months             int32
	UntilDate          int32
}

func (*MessageMediaGiveawayPredict) CRC() uint32 {
	return 0xdaad85b0
}
func (*MessageMediaGiveawayPredict) _MessageMedia() {}

type MessageMediaGiveawayResultsPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers   bool     `tl:",omitempty:flags:0,implicit"`
	Refunded             bool     `tl:",omitempty:flags:2,implicit"`
	ChannelID            int64
	AdditionalPeersCount *int32 `tl:",omitempty:flags:3"`
	LaunchMsgID          int32
	WinnersCount         int32
	UnclaimedCount       int32
	Winners              []int64
	Months               int32
	PrizeDescription     *string `tl:",omitempty:flags:1"`
	UntilDate            int32
}

func (*MessageMediaGiveawayResultsPredict) CRC() uint32 {
	return 0xc6991068
}
func (*MessageMediaGiveawayResultsPredict) _MessageMedia() {}

type MessageMediaPaidMediaPredict struct {
	StarsAmount   int64
	ExtendedMedia []MessageExtendedMedia
}

func (*MessageMediaPaidMediaPredict) CRC() uint32 {
	return 0xa8852491
}
func (*MessageMediaPaidMediaPredict) _MessageMedia() {}

type MessagePeerReaction interface {
	tl.Object
	_MessagePeerReaction()
}

var (
	_ MessagePeerReaction = (*MessagePeerReactionPredict)(nil)
)

type MessagePeerReactionPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Big      bool     `tl:",omitempty:flags:0,implicit"`
	Unread   bool     `tl:",omitempty:flags:1,implicit"`
	My       bool     `tl:",omitempty:flags:2,implicit"`
	PeerID   Peer
	Date     int32
	Reaction Reaction
}

func (*MessagePeerReactionPredict) CRC() uint32 {
	return 0x8c79b63c
}
func (*MessagePeerReactionPredict) _MessagePeerReaction() {}

type MessagePeerVote interface {
	tl.Object
	_MessagePeerVote()
}

var (
	_ MessagePeerVote = (*MessagePeerVotePredict)(nil)
	_ MessagePeerVote = (*MessagePeerVoteInputOptionPredict)(nil)
	_ MessagePeerVote = (*MessagePeerVoteMultiplePredict)(nil)
)

type MessagePeerVotePredict struct {
	Peer   Peer
	Option []byte
	Date   int32
}

func (*MessagePeerVotePredict) CRC() uint32 {
	return 0xb6cc2d5c
}
func (*MessagePeerVotePredict) _MessagePeerVote() {}

type MessagePeerVoteInputOptionPredict struct {
	Peer Peer
	Date int32
}

func (*MessagePeerVoteInputOptionPredict) CRC() uint32 {
	return 0x74cda504
}
func (*MessagePeerVoteInputOptionPredict) _MessagePeerVote() {}

type MessagePeerVoteMultiplePredict struct {
	Peer    Peer
	Options [][]byte
	Date    int32
}

func (*MessagePeerVoteMultiplePredict) CRC() uint32 {
	return 0x4628f6e6
}
func (*MessagePeerVoteMultiplePredict) _MessagePeerVote() {}

type MessageRange interface {
	tl.Object
	_MessageRange()
}

var (
	_ MessageRange = (*MessageRangePredict)(nil)
)

type MessageRangePredict struct {
	MinID int32
	MaxID int32
}

func (*MessageRangePredict) CRC() uint32 {
	return 0xae30253
}
func (*MessageRangePredict) _MessageRange() {}

type MessageReactions interface {
	tl.Object
	_MessageReactions()
}

var (
	_ MessageReactions = (*MessageReactionsPredict)(nil)
)

type MessageReactionsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Min             bool     `tl:",omitempty:flags:0,implicit"`
	CanSeeList      bool     `tl:",omitempty:flags:2,implicit"`
	ReactionsAsTags bool     `tl:",omitempty:flags:3,implicit"`
	Results         []ReactionCount
	RecentReactions []MessagePeerReaction `tl:",omitempty:flags:1"`
}

func (*MessageReactionsPredict) CRC() uint32 {
	return 0x4f2b9479
}
func (*MessageReactionsPredict) _MessageReactions() {}

type MessageReplies interface {
	tl.Object
	_MessageReplies()
}

var (
	_ MessageReplies = (*MessageRepliesPredict)(nil)
)

type MessageRepliesPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Comments       bool     `tl:",omitempty:flags:0,implicit"`
	Replies        int32
	RepliesPts     int32
	RecentRepliers []Peer `tl:",omitempty:flags:1"`
	ChannelID      *int64 `tl:",omitempty:flags:0"`
	MaxID          *int32 `tl:",omitempty:flags:2"`
	ReadMaxID      *int32 `tl:",omitempty:flags:3"`
}

func (*MessageRepliesPredict) CRC() uint32 {
	return 0x83d60fc2
}
func (*MessageRepliesPredict) _MessageReplies() {}

type MessageReplyHeader interface {
	tl.Object
	_MessageReplyHeader()
}

var (
	_ MessageReplyHeader = (*MessageReplyHeaderPredict)(nil)
	_ MessageReplyHeader = (*MessageReplyStoryHeaderPredict)(nil)
)

type MessageReplyHeaderPredict struct {
	_                struct{}         `tl:"flags,bitflag"`
	ReplyToScheduled bool             `tl:",omitempty:flags:2,implicit"`
	ForumTopic       bool             `tl:",omitempty:flags:3,implicit"`
	Quote            bool             `tl:",omitempty:flags:9,implicit"`
	ReplyToMsgID     *int32           `tl:",omitempty:flags:4"`
	ReplyToPeerID    Peer             `tl:",omitempty:flags:0"`
	ReplyFrom        MessageFwdHeader `tl:",omitempty:flags:5"`
	ReplyMedia       MessageMedia     `tl:",omitempty:flags:8"`
	ReplyToTopID     *int32           `tl:",omitempty:flags:1"`
	QuoteText        *string          `tl:",omitempty:flags:6"`
	QuoteEntities    []MessageEntity  `tl:",omitempty:flags:7"`
	QuoteOffset      *int32           `tl:",omitempty:flags:10"`
}

func (*MessageReplyHeaderPredict) CRC() uint32 {
	return 0xafbc09db
}
func (*MessageReplyHeaderPredict) _MessageReplyHeader() {}

type MessageReplyStoryHeaderPredict struct {
	Peer    Peer
	StoryID int32
}

func (*MessageReplyStoryHeaderPredict) CRC() uint32 {
	return 0xe5af939
}
func (*MessageReplyStoryHeaderPredict) _MessageReplyHeader() {}

type MessageViews interface {
	tl.Object
	_MessageViews()
}

var (
	_ MessageViews = (*MessageViewsPredict)(nil)
)

type MessageViewsPredict struct {
	_        struct{}       `tl:"flags,bitflag"`
	Views    *int32         `tl:",omitempty:flags:0"`
	Forwards *int32         `tl:",omitempty:flags:1"`
	Replies  MessageReplies `tl:",omitempty:flags:2"`
}

func (*MessageViewsPredict) CRC() uint32 {
	return 0x455b853d
}
func (*MessageViewsPredict) _MessageViews() {}

type MessagesFilter interface {
	tl.Object
	_MessagesFilter()
}

var (
	_ MessagesFilter = (*InputMessagesFilterEmptyPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterPhotosPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterVideoPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterPhotoVideoPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterDocumentPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterURLPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterGifPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterVoicePredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterMusicPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterChatPhotosPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterPhoneCallsPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterRoundVoicePredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterRoundVideoPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterMyMentionsPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterGeoPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterContactsPredict)(nil)
	_ MessagesFilter = (*InputMessagesFilterPinnedPredict)(nil)
)

type InputMessagesFilterEmptyPredict struct{}

func (*InputMessagesFilterEmptyPredict) CRC() uint32 {
	return 0x57e2f66c
}
func (*InputMessagesFilterEmptyPredict) _MessagesFilter() {}

type InputMessagesFilterPhotosPredict struct{}

func (*InputMessagesFilterPhotosPredict) CRC() uint32 {
	return 0x9609a51c
}
func (*InputMessagesFilterPhotosPredict) _MessagesFilter() {}

type InputMessagesFilterVideoPredict struct{}

func (*InputMessagesFilterVideoPredict) CRC() uint32 {
	return 0x9fc00e65
}
func (*InputMessagesFilterVideoPredict) _MessagesFilter() {}

type InputMessagesFilterPhotoVideoPredict struct{}

func (*InputMessagesFilterPhotoVideoPredict) CRC() uint32 {
	return 0x56e9f0e4
}
func (*InputMessagesFilterPhotoVideoPredict) _MessagesFilter() {}

type InputMessagesFilterDocumentPredict struct{}

func (*InputMessagesFilterDocumentPredict) CRC() uint32 {
	return 0x9eddf188
}
func (*InputMessagesFilterDocumentPredict) _MessagesFilter() {}

type InputMessagesFilterURLPredict struct{}

func (*InputMessagesFilterURLPredict) CRC() uint32 {
	return 0x7ef0dd87
}
func (*InputMessagesFilterURLPredict) _MessagesFilter() {}

type InputMessagesFilterGifPredict struct{}

func (*InputMessagesFilterGifPredict) CRC() uint32 {
	return 0xffc86587
}
func (*InputMessagesFilterGifPredict) _MessagesFilter() {}

type InputMessagesFilterVoicePredict struct{}

func (*InputMessagesFilterVoicePredict) CRC() uint32 {
	return 0x50f5c392
}
func (*InputMessagesFilterVoicePredict) _MessagesFilter() {}

type InputMessagesFilterMusicPredict struct{}

func (*InputMessagesFilterMusicPredict) CRC() uint32 {
	return 0x3751b49e
}
func (*InputMessagesFilterMusicPredict) _MessagesFilter() {}

type InputMessagesFilterChatPhotosPredict struct{}

func (*InputMessagesFilterChatPhotosPredict) CRC() uint32 {
	return 0x3a20ecb8
}
func (*InputMessagesFilterChatPhotosPredict) _MessagesFilter() {}

type InputMessagesFilterPhoneCallsPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Missed bool     `tl:",omitempty:flags:0,implicit"`
}

func (*InputMessagesFilterPhoneCallsPredict) CRC() uint32 {
	return 0x80c99768
}
func (*InputMessagesFilterPhoneCallsPredict) _MessagesFilter() {}

type InputMessagesFilterRoundVoicePredict struct{}

func (*InputMessagesFilterRoundVoicePredict) CRC() uint32 {
	return 0x7a7c17a4
}
func (*InputMessagesFilterRoundVoicePredict) _MessagesFilter() {}

type InputMessagesFilterRoundVideoPredict struct{}

func (*InputMessagesFilterRoundVideoPredict) CRC() uint32 {
	return 0xb549da53
}
func (*InputMessagesFilterRoundVideoPredict) _MessagesFilter() {}

type InputMessagesFilterMyMentionsPredict struct{}

func (*InputMessagesFilterMyMentionsPredict) CRC() uint32 {
	return 0xc1f8e69a
}
func (*InputMessagesFilterMyMentionsPredict) _MessagesFilter() {}

type InputMessagesFilterGeoPredict struct{}

func (*InputMessagesFilterGeoPredict) CRC() uint32 {
	return 0xe7026d0d
}
func (*InputMessagesFilterGeoPredict) _MessagesFilter() {}

type InputMessagesFilterContactsPredict struct{}

func (*InputMessagesFilterContactsPredict) CRC() uint32 {
	return 0xe062db83
}
func (*InputMessagesFilterContactsPredict) _MessagesFilter() {}

type InputMessagesFilterPinnedPredict struct{}

func (*InputMessagesFilterPinnedPredict) CRC() uint32 {
	return 0x1bb00451
}
func (*InputMessagesFilterPinnedPredict) _MessagesFilter() {}

type MissingInvitee interface {
	tl.Object
	_MissingInvitee()
}

var (
	_ MissingInvitee = (*MissingInviteePredict)(nil)
)

type MissingInviteePredict struct {
	_                       struct{} `tl:"flags,bitflag"`
	PremiumWouldAllowInvite bool     `tl:",omitempty:flags:0,implicit"`
	PremiumRequiredForPm    bool     `tl:",omitempty:flags:1,implicit"`
	UserID                  int64
}

func (*MissingInviteePredict) CRC() uint32 {
	return 0x628c9224
}
func (*MissingInviteePredict) _MissingInvitee() {}

type MyBoost interface {
	tl.Object
	_MyBoost()
}

var (
	_ MyBoost = (*MyBoostPredict)(nil)
)

type MyBoostPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Slot              int32
	Peer              Peer `tl:",omitempty:flags:0"`
	Date              int32
	Expires           int32
	CooldownUntilDate *int32 `tl:",omitempty:flags:1"`
}

func (*MyBoostPredict) CRC() uint32 {
	return 0xc448415c
}
func (*MyBoostPredict) _MyBoost() {}

type NearestDc interface {
	tl.Object
	_NearestDc()
}

var (
	_ NearestDc = (*NearestDcPredict)(nil)
)

type NearestDcPredict struct {
	Country   string
	ThisDc    int32
	NearestDc int32
}

func (*NearestDcPredict) CRC() uint32 {
	return 0x8e1a1775
}
func (*NearestDcPredict) _NearestDc() {}

type NotificationSound interface {
	tl.Object
	_NotificationSound()
}

var (
	_ NotificationSound = (*NotificationSoundDefaultPredict)(nil)
	_ NotificationSound = (*NotificationSoundNonePredict)(nil)
	_ NotificationSound = (*NotificationSoundLocalPredict)(nil)
	_ NotificationSound = (*NotificationSoundRingtonePredict)(nil)
)

type NotificationSoundDefaultPredict struct{}

func (*NotificationSoundDefaultPredict) CRC() uint32 {
	return 0x97e8bebe
}
func (*NotificationSoundDefaultPredict) _NotificationSound() {}

type NotificationSoundNonePredict struct{}

func (*NotificationSoundNonePredict) CRC() uint32 {
	return 0x6f0c34df
}
func (*NotificationSoundNonePredict) _NotificationSound() {}

type NotificationSoundLocalPredict struct {
	Title string
	Data  string
}

func (*NotificationSoundLocalPredict) CRC() uint32 {
	return 0x830b9ae4
}
func (*NotificationSoundLocalPredict) _NotificationSound() {}

type NotificationSoundRingtonePredict struct {
	ID int64
}

func (*NotificationSoundRingtonePredict) CRC() uint32 {
	return 0xff6c8049
}
func (*NotificationSoundRingtonePredict) _NotificationSound() {}

type NotifyPeer interface {
	tl.Object
	_NotifyPeer()
}

var (
	_ NotifyPeer = (*NotifyPeerPredict)(nil)
	_ NotifyPeer = (*NotifyUsersPredict)(nil)
	_ NotifyPeer = (*NotifyChatsPredict)(nil)
	_ NotifyPeer = (*NotifyBroadcastsPredict)(nil)
	_ NotifyPeer = (*NotifyForumTopicPredict)(nil)
)

type NotifyPeerPredict struct {
	Peer Peer
}

func (*NotifyPeerPredict) CRC() uint32 {
	return 0x9fd40bd8
}
func (*NotifyPeerPredict) _NotifyPeer() {}

type NotifyUsersPredict struct{}

func (*NotifyUsersPredict) CRC() uint32 {
	return 0xb4c83b4c
}
func (*NotifyUsersPredict) _NotifyPeer() {}

type NotifyChatsPredict struct{}

func (*NotifyChatsPredict) CRC() uint32 {
	return 0xc007cec3
}
func (*NotifyChatsPredict) _NotifyPeer() {}

type NotifyBroadcastsPredict struct{}

func (*NotifyBroadcastsPredict) CRC() uint32 {
	return 0xd612e8ef
}
func (*NotifyBroadcastsPredict) _NotifyPeer() {}

type NotifyForumTopicPredict struct {
	Peer     Peer
	TopMsgID int32
}

func (*NotifyForumTopicPredict) CRC() uint32 {
	return 0x226e6308
}
func (*NotifyForumTopicPredict) _NotifyPeer() {}

type OutboxReadDate interface {
	tl.Object
	_OutboxReadDate()
}

var (
	_ OutboxReadDate = (*OutboxReadDatePredict)(nil)
)

type OutboxReadDatePredict struct {
	Date int32
}

func (*OutboxReadDatePredict) CRC() uint32 {
	return 0x3bb842ac
}
func (*OutboxReadDatePredict) _OutboxReadDate() {}

type Page interface {
	tl.Object
	_Page()
}

var (
	_ Page = (*PagePredict)(nil)
)

type PagePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Part      bool     `tl:",omitempty:flags:0,implicit"`
	Rtl       bool     `tl:",omitempty:flags:1,implicit"`
	V2        bool     `tl:",omitempty:flags:2,implicit"`
	URL       string
	Blocks    []PageBlock
	Photos    []Photo
	Documents []Document
	Views     *int32 `tl:",omitempty:flags:3"`
}

func (*PagePredict) CRC() uint32 {
	return 0x98657f0d
}
func (*PagePredict) _Page() {}

type PageBlock interface {
	tl.Object
	_PageBlock()
}

var (
	_ PageBlock = (*PageBlockUnsupportedPredict)(nil)
	_ PageBlock = (*PageBlockTitlePredict)(nil)
	_ PageBlock = (*PageBlockSubtitlePredict)(nil)
	_ PageBlock = (*PageBlockAuthorDatePredict)(nil)
	_ PageBlock = (*PageBlockHeaderPredict)(nil)
	_ PageBlock = (*PageBlockSubheaderPredict)(nil)
	_ PageBlock = (*PageBlockParagraphPredict)(nil)
	_ PageBlock = (*PageBlockPreformattedPredict)(nil)
	_ PageBlock = (*PageBlockFooterPredict)(nil)
	_ PageBlock = (*PageBlockDividerPredict)(nil)
	_ PageBlock = (*PageBlockAnchorPredict)(nil)
	_ PageBlock = (*PageBlockListPredict)(nil)
	_ PageBlock = (*PageBlockBlockquotePredict)(nil)
	_ PageBlock = (*PageBlockPullquotePredict)(nil)
	_ PageBlock = (*PageBlockPhotoPredict)(nil)
	_ PageBlock = (*PageBlockVideoPredict)(nil)
	_ PageBlock = (*PageBlockCoverPredict)(nil)
	_ PageBlock = (*PageBlockEmbedPredict)(nil)
	_ PageBlock = (*PageBlockEmbedPostPredict)(nil)
	_ PageBlock = (*PageBlockCollagePredict)(nil)
	_ PageBlock = (*PageBlockSlideshowPredict)(nil)
	_ PageBlock = (*PageBlockChannelPredict)(nil)
	_ PageBlock = (*PageBlockAudioPredict)(nil)
	_ PageBlock = (*PageBlockKickerPredict)(nil)
	_ PageBlock = (*PageBlockTablePredict)(nil)
	_ PageBlock = (*PageBlockOrderedListPredict)(nil)
	_ PageBlock = (*PageBlockDetailsPredict)(nil)
	_ PageBlock = (*PageBlockRelatedArticlesPredict)(nil)
	_ PageBlock = (*PageBlockMapPredict)(nil)
)

type PageBlockUnsupportedPredict struct{}

func (*PageBlockUnsupportedPredict) CRC() uint32 {
	return 0x13567e8a
}
func (*PageBlockUnsupportedPredict) _PageBlock() {}

type PageBlockTitlePredict struct {
	Text RichText
}

func (*PageBlockTitlePredict) CRC() uint32 {
	return 0x70abc3fd
}
func (*PageBlockTitlePredict) _PageBlock() {}

type PageBlockSubtitlePredict struct {
	Text RichText
}

func (*PageBlockSubtitlePredict) CRC() uint32 {
	return 0x8ffa9a1f
}
func (*PageBlockSubtitlePredict) _PageBlock() {}

type PageBlockAuthorDatePredict struct {
	Author        RichText
	PublishedDate int32
}

func (*PageBlockAuthorDatePredict) CRC() uint32 {
	return 0xbaafe5e0
}
func (*PageBlockAuthorDatePredict) _PageBlock() {}

type PageBlockHeaderPredict struct {
	Text RichText
}

func (*PageBlockHeaderPredict) CRC() uint32 {
	return 0xbfd064ec
}
func (*PageBlockHeaderPredict) _PageBlock() {}

type PageBlockSubheaderPredict struct {
	Text RichText
}

func (*PageBlockSubheaderPredict) CRC() uint32 {
	return 0xf12bb6e1
}
func (*PageBlockSubheaderPredict) _PageBlock() {}

type PageBlockParagraphPredict struct {
	Text RichText
}

func (*PageBlockParagraphPredict) CRC() uint32 {
	return 0x467a0766
}
func (*PageBlockParagraphPredict) _PageBlock() {}

type PageBlockPreformattedPredict struct {
	Text     RichText
	Language string
}

func (*PageBlockPreformattedPredict) CRC() uint32 {
	return 0xc070d93e
}
func (*PageBlockPreformattedPredict) _PageBlock() {}

type PageBlockFooterPredict struct {
	Text RichText
}

func (*PageBlockFooterPredict) CRC() uint32 {
	return 0x48870999
}
func (*PageBlockFooterPredict) _PageBlock() {}

type PageBlockDividerPredict struct{}

func (*PageBlockDividerPredict) CRC() uint32 {
	return 0xdb20b188
}
func (*PageBlockDividerPredict) _PageBlock() {}

type PageBlockAnchorPredict struct {
	Name string
}

func (*PageBlockAnchorPredict) CRC() uint32 {
	return 0xce0d37b0
}
func (*PageBlockAnchorPredict) _PageBlock() {}

type PageBlockListPredict struct {
	Items []PageListItem
}

func (*PageBlockListPredict) CRC() uint32 {
	return 0xe4e88011
}
func (*PageBlockListPredict) _PageBlock() {}

type PageBlockBlockquotePredict struct {
	Text    RichText
	Caption RichText
}

func (*PageBlockBlockquotePredict) CRC() uint32 {
	return 0x263d7c26
}
func (*PageBlockBlockquotePredict) _PageBlock() {}

type PageBlockPullquotePredict struct {
	Text    RichText
	Caption RichText
}

func (*PageBlockPullquotePredict) CRC() uint32 {
	return 0x4f4456d3
}
func (*PageBlockPullquotePredict) _PageBlock() {}

type PageBlockPhotoPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	PhotoID   int64
	Caption   PageCaption
	URL       *string `tl:",omitempty:flags:0"`
	WebpageID *int64  `tl:",omitempty:flags:0"`
}

func (*PageBlockPhotoPredict) CRC() uint32 {
	return 0x1759c560
}
func (*PageBlockPhotoPredict) _PageBlock() {}

type PageBlockVideoPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Autoplay bool     `tl:",omitempty:flags:0,implicit"`
	Loop     bool     `tl:",omitempty:flags:1,implicit"`
	VideoID  int64
	Caption  PageCaption
}

func (*PageBlockVideoPredict) CRC() uint32 {
	return 0x7c8fe7b6
}
func (*PageBlockVideoPredict) _PageBlock() {}

type PageBlockCoverPredict struct {
	Cover PageBlock
}

func (*PageBlockCoverPredict) CRC() uint32 {
	return 0x39f23300
}
func (*PageBlockCoverPredict) _PageBlock() {}

type PageBlockEmbedPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	FullWidth      bool     `tl:",omitempty:flags:0,implicit"`
	AllowScrolling bool     `tl:",omitempty:flags:3,implicit"`
	URL            *string  `tl:",omitempty:flags:1"`
	Html           *string  `tl:",omitempty:flags:2"`
	PosterPhotoID  *int64   `tl:",omitempty:flags:4"`
	W              *int32   `tl:",omitempty:flags:5"`
	H              *int32   `tl:",omitempty:flags:5"`
	Caption        PageCaption
}

func (*PageBlockEmbedPredict) CRC() uint32 {
	return 0xa8718dc5
}
func (*PageBlockEmbedPredict) _PageBlock() {}

type PageBlockEmbedPostPredict struct {
	URL           string
	WebpageID     int64
	AuthorPhotoID int64
	Author        string
	Date          int32
	Blocks        []PageBlock
	Caption       PageCaption
}

func (*PageBlockEmbedPostPredict) CRC() uint32 {
	return 0xf259a80b
}
func (*PageBlockEmbedPostPredict) _PageBlock() {}

type PageBlockCollagePredict struct {
	Items   []PageBlock
	Caption PageCaption
}

func (*PageBlockCollagePredict) CRC() uint32 {
	return 0x65a0fa4d
}
func (*PageBlockCollagePredict) _PageBlock() {}

type PageBlockSlideshowPredict struct {
	Items   []PageBlock
	Caption PageCaption
}

func (*PageBlockSlideshowPredict) CRC() uint32 {
	return 0x31f9590
}
func (*PageBlockSlideshowPredict) _PageBlock() {}

type PageBlockChannelPredict struct {
	Channel Chat
}

func (*PageBlockChannelPredict) CRC() uint32 {
	return 0xef1751b5
}
func (*PageBlockChannelPredict) _PageBlock() {}

type PageBlockAudioPredict struct {
	AudioID int64
	Caption PageCaption
}

func (*PageBlockAudioPredict) CRC() uint32 {
	return 0x804361ea
}
func (*PageBlockAudioPredict) _PageBlock() {}

type PageBlockKickerPredict struct {
	Text RichText
}

func (*PageBlockKickerPredict) CRC() uint32 {
	return 0x1e148390
}
func (*PageBlockKickerPredict) _PageBlock() {}

type PageBlockTablePredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Bordered bool     `tl:",omitempty:flags:0,implicit"`
	Striped  bool     `tl:",omitempty:flags:1,implicit"`
	Title    RichText
	Rows     []PageTableRow
}

func (*PageBlockTablePredict) CRC() uint32 {
	return 0xbf4dea82
}
func (*PageBlockTablePredict) _PageBlock() {}

type PageBlockOrderedListPredict struct {
	Items []PageListOrderedItem
}

func (*PageBlockOrderedListPredict) CRC() uint32 {
	return 0x9a8ae1e1
}
func (*PageBlockOrderedListPredict) _PageBlock() {}

type PageBlockDetailsPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Open   bool     `tl:",omitempty:flags:0,implicit"`
	Blocks []PageBlock
	Title  RichText
}

func (*PageBlockDetailsPredict) CRC() uint32 {
	return 0x76768bed
}
func (*PageBlockDetailsPredict) _PageBlock() {}

type PageBlockRelatedArticlesPredict struct {
	Title    RichText
	Articles []PageRelatedArticle
}

func (*PageBlockRelatedArticlesPredict) CRC() uint32 {
	return 0x16115a96
}
func (*PageBlockRelatedArticlesPredict) _PageBlock() {}

type PageBlockMapPredict struct {
	Geo     GeoPoint
	Zoom    int32
	W       int32
	H       int32
	Caption PageCaption
}

func (*PageBlockMapPredict) CRC() uint32 {
	return 0xa44f3ef6
}
func (*PageBlockMapPredict) _PageBlock() {}

type PageCaption interface {
	tl.Object
	_PageCaption()
}

var (
	_ PageCaption = (*PageCaptionPredict)(nil)
)

type PageCaptionPredict struct {
	Text   RichText
	Credit RichText
}

func (*PageCaptionPredict) CRC() uint32 {
	return 0x6f747657
}
func (*PageCaptionPredict) _PageCaption() {}

type PageListItem interface {
	tl.Object
	_PageListItem()
}

var (
	_ PageListItem = (*PageListItemTextPredict)(nil)
	_ PageListItem = (*PageListItemBlocksPredict)(nil)
)

type PageListItemTextPredict struct {
	Text RichText
}

func (*PageListItemTextPredict) CRC() uint32 {
	return 0xb92fb6cd
}
func (*PageListItemTextPredict) _PageListItem() {}

type PageListItemBlocksPredict struct {
	Blocks []PageBlock
}

func (*PageListItemBlocksPredict) CRC() uint32 {
	return 0x25e073fc
}
func (*PageListItemBlocksPredict) _PageListItem() {}

type PageListOrderedItem interface {
	tl.Object
	_PageListOrderedItem()
}

var (
	_ PageListOrderedItem = (*PageListOrderedItemTextPredict)(nil)
	_ PageListOrderedItem = (*PageListOrderedItemBlocksPredict)(nil)
)

type PageListOrderedItemTextPredict struct {
	Num  string
	Text RichText
}

func (*PageListOrderedItemTextPredict) CRC() uint32 {
	return 0x5e068047
}
func (*PageListOrderedItemTextPredict) _PageListOrderedItem() {}

type PageListOrderedItemBlocksPredict struct {
	Num    string
	Blocks []PageBlock
}

func (*PageListOrderedItemBlocksPredict) CRC() uint32 {
	return 0x98dd8936
}
func (*PageListOrderedItemBlocksPredict) _PageListOrderedItem() {}

type PageRelatedArticle interface {
	tl.Object
	_PageRelatedArticle()
}

var (
	_ PageRelatedArticle = (*PageRelatedArticlePredict)(nil)
)

type PageRelatedArticlePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	URL           string
	WebpageID     int64
	Title         *string `tl:",omitempty:flags:0"`
	Description   *string `tl:",omitempty:flags:1"`
	PhotoID       *int64  `tl:",omitempty:flags:2"`
	Author        *string `tl:",omitempty:flags:3"`
	PublishedDate *int32  `tl:",omitempty:flags:4"`
}

func (*PageRelatedArticlePredict) CRC() uint32 {
	return 0xb390dc08
}
func (*PageRelatedArticlePredict) _PageRelatedArticle() {}

type PageTableCell interface {
	tl.Object
	_PageTableCell()
}

var (
	_ PageTableCell = (*PageTableCellPredict)(nil)
)

type PageTableCellPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Header       bool     `tl:",omitempty:flags:0,implicit"`
	AlignCenter  bool     `tl:",omitempty:flags:3,implicit"`
	AlignRight   bool     `tl:",omitempty:flags:4,implicit"`
	ValignMiddle bool     `tl:",omitempty:flags:5,implicit"`
	ValignBottom bool     `tl:",omitempty:flags:6,implicit"`
	Text         RichText `tl:",omitempty:flags:7"`
	Colspan      *int32   `tl:",omitempty:flags:1"`
	Rowspan      *int32   `tl:",omitempty:flags:2"`
}

func (*PageTableCellPredict) CRC() uint32 {
	return 0x34566b6a
}
func (*PageTableCellPredict) _PageTableCell() {}

type PageTableRow interface {
	tl.Object
	_PageTableRow()
}

var (
	_ PageTableRow = (*PageTableRowPredict)(nil)
)

type PageTableRowPredict struct {
	Cells []PageTableCell
}

func (*PageTableRowPredict) CRC() uint32 {
	return 0xe0c0c5e5
}
func (*PageTableRowPredict) _PageTableRow() {}

type PasswordKdfAlgo interface {
	tl.Object
	_PasswordKdfAlgo()
}

var (
	_ PasswordKdfAlgo = (*PasswordKdfAlgoUnknownPredict)(nil)
	_ PasswordKdfAlgo = (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPowPredict)(nil)
)

type PasswordKdfAlgoUnknownPredict struct{}

func (*PasswordKdfAlgoUnknownPredict) CRC() uint32 {
	return 0xd45ab096
}
func (*PasswordKdfAlgoUnknownPredict) _PasswordKdfAlgo() {}

type PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPowPredict struct {
	Salt1 []byte
	Salt2 []byte
	G     int32
	P     []byte
}

func (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPowPredict) CRC() uint32 {
	return 0x3a912d4a
}
func (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPowPredict) _PasswordKdfAlgo() {}

type PaymentCharge interface {
	tl.Object
	_PaymentCharge()
}

var (
	_ PaymentCharge = (*PaymentChargePredict)(nil)
)

type PaymentChargePredict struct {
	ID               string
	ProviderChargeID string
}

func (*PaymentChargePredict) CRC() uint32 {
	return 0xea02c27e
}
func (*PaymentChargePredict) _PaymentCharge() {}

type PaymentFormMethod interface {
	tl.Object
	_PaymentFormMethod()
}

var (
	_ PaymentFormMethod = (*PaymentFormMethodPredict)(nil)
)

type PaymentFormMethodPredict struct {
	URL   string
	Title string
}

func (*PaymentFormMethodPredict) CRC() uint32 {
	return 0x88f8f21b
}
func (*PaymentFormMethodPredict) _PaymentFormMethod() {}

type PaymentRequestedInfo interface {
	tl.Object
	_PaymentRequestedInfo()
}

var (
	_ PaymentRequestedInfo = (*PaymentRequestedInfoPredict)(nil)
)

type PaymentRequestedInfoPredict struct {
	_               struct{}    `tl:"flags,bitflag"`
	Name            *string     `tl:",omitempty:flags:0"`
	Phone           *string     `tl:",omitempty:flags:1"`
	Email           *string     `tl:",omitempty:flags:2"`
	ShippingAddress PostAddress `tl:",omitempty:flags:3"`
}

func (*PaymentRequestedInfoPredict) CRC() uint32 {
	return 0x909c3f94
}
func (*PaymentRequestedInfoPredict) _PaymentRequestedInfo() {}

type PaymentSavedCredentials interface {
	tl.Object
	_PaymentSavedCredentials()
}

var (
	_ PaymentSavedCredentials = (*PaymentSavedCredentialsCardPredict)(nil)
)

type PaymentSavedCredentialsCardPredict struct {
	ID    string
	Title string
}

func (*PaymentSavedCredentialsCardPredict) CRC() uint32 {
	return 0xcdc27a1f
}
func (*PaymentSavedCredentialsCardPredict) _PaymentSavedCredentials() {}

type Peer interface {
	tl.Object
	_Peer()
}

var (
	_ Peer = (*PeerUserPredict)(nil)
	_ Peer = (*PeerChatPredict)(nil)
	_ Peer = (*PeerChannelPredict)(nil)
)

type PeerUserPredict struct {
	UserID int64
}

func (*PeerUserPredict) CRC() uint32 {
	return 0x59511722
}
func (*PeerUserPredict) _Peer() {}

type PeerChatPredict struct {
	ChatID int64
}

func (*PeerChatPredict) CRC() uint32 {
	return 0x36c6019a
}
func (*PeerChatPredict) _Peer() {}

type PeerChannelPredict struct {
	ChannelID int64
}

func (*PeerChannelPredict) CRC() uint32 {
	return 0xa2a5371e
}
func (*PeerChannelPredict) _Peer() {}

type PeerBlocked interface {
	tl.Object
	_PeerBlocked()
}

var (
	_ PeerBlocked = (*PeerBlockedPredict)(nil)
)

type PeerBlockedPredict struct {
	PeerID Peer
	Date   int32
}

func (*PeerBlockedPredict) CRC() uint32 {
	return 0xe8fd8014
}
func (*PeerBlockedPredict) _PeerBlocked() {}

type PeerColor interface {
	tl.Object
	_PeerColor()
}

var (
	_ PeerColor = (*PeerColorPredict)(nil)
)

type PeerColorPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Color             *int32   `tl:",omitempty:flags:0"`
	BackgroundEmojiID *int64   `tl:",omitempty:flags:1"`
}

func (*PeerColorPredict) CRC() uint32 {
	return 0xb54b5acf
}
func (*PeerColorPredict) _PeerColor() {}

type PeerLocated interface {
	tl.Object
	_PeerLocated()
}

var (
	_ PeerLocated = (*PeerLocatedPredict)(nil)
	_ PeerLocated = (*PeerSelfLocatedPredict)(nil)
)

type PeerLocatedPredict struct {
	Peer     Peer
	Expires  int32
	Distance int32
}

func (*PeerLocatedPredict) CRC() uint32 {
	return 0xca461b5d
}
func (*PeerLocatedPredict) _PeerLocated() {}

type PeerSelfLocatedPredict struct {
	Expires int32
}

func (*PeerSelfLocatedPredict) CRC() uint32 {
	return 0xf8ec284b
}
func (*PeerSelfLocatedPredict) _PeerLocated() {}

type PeerNotifySettings interface {
	tl.Object
	_PeerNotifySettings()
}

var (
	_ PeerNotifySettings = (*PeerNotifySettingsPredict)(nil)
)

type PeerNotifySettingsPredict struct {
	_                   struct{}          `tl:"flags,bitflag"`
	ShowPreviews        *bool             `tl:",omitempty:flags:0"`
	Silent              *bool             `tl:",omitempty:flags:1"`
	MuteUntil           *int32            `tl:",omitempty:flags:2"`
	IosSound            NotificationSound `tl:",omitempty:flags:3"`
	AndroidSound        NotificationSound `tl:",omitempty:flags:4"`
	OtherSound          NotificationSound `tl:",omitempty:flags:5"`
	StoriesMuted        *bool             `tl:",omitempty:flags:6"`
	StoriesHideSender   *bool             `tl:",omitempty:flags:7"`
	StoriesIosSound     NotificationSound `tl:",omitempty:flags:8"`
	StoriesAndroidSound NotificationSound `tl:",omitempty:flags:9"`
	StoriesOtherSound   NotificationSound `tl:",omitempty:flags:10"`
}

func (*PeerNotifySettingsPredict) CRC() uint32 {
	return 0x99622c0c
}
func (*PeerNotifySettingsPredict) _PeerNotifySettings() {}

type PeerSettings interface {
	tl.Object
	_PeerSettings()
}

var (
	_ PeerSettings = (*PeerSettingsPredict)(nil)
)

type PeerSettingsPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	ReportSpam            bool     `tl:",omitempty:flags:0,implicit"`
	AddContact            bool     `tl:",omitempty:flags:1,implicit"`
	BlockContact          bool     `tl:",omitempty:flags:2,implicit"`
	ShareContact          bool     `tl:",omitempty:flags:3,implicit"`
	NeedContactsException bool     `tl:",omitempty:flags:4,implicit"`
	ReportGeo             bool     `tl:",omitempty:flags:5,implicit"`
	Autoarchived          bool     `tl:",omitempty:flags:7,implicit"`
	InviteMembers         bool     `tl:",omitempty:flags:8,implicit"`
	RequestChatBroadcast  bool     `tl:",omitempty:flags:10,implicit"`
	BusinessBotPaused     bool     `tl:",omitempty:flags:11,implicit"`
	BusinessBotCanReply   bool     `tl:",omitempty:flags:12,implicit"`
	GeoDistance           *int32   `tl:",omitempty:flags:6"`
	RequestChatTitle      *string  `tl:",omitempty:flags:9"`
	RequestChatDate       *int32   `tl:",omitempty:flags:9"`
	BusinessBotID         *int64   `tl:",omitempty:flags:13"`
	BusinessBotManageURL  *string  `tl:",omitempty:flags:13"`
}

func (*PeerSettingsPredict) CRC() uint32 {
	return 0xacd66c5e
}
func (*PeerSettingsPredict) _PeerSettings() {}

type PeerStories interface {
	tl.Object
	_PeerStories()
}

var (
	_ PeerStories = (*PeerStoriesPredict)(nil)
)

type PeerStoriesPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      Peer
	MaxReadID *int32 `tl:",omitempty:flags:0"`
	Stories   []StoryItem
}

func (*PeerStoriesPredict) CRC() uint32 {
	return 0x9a35e999
}
func (*PeerStoriesPredict) _PeerStories() {}

type PhoneCall interface {
	tl.Object
	_PhoneCall()
}

var (
	_ PhoneCall = (*PhoneCallEmptyPredict)(nil)
	_ PhoneCall = (*PhoneCallWaitingPredict)(nil)
	_ PhoneCall = (*PhoneCallRequestedPredict)(nil)
	_ PhoneCall = (*PhoneCallAcceptedPredict)(nil)
	_ PhoneCall = (*PhoneCallPredict)(nil)
	_ PhoneCall = (*PhoneCallDiscardedPredict)(nil)
)

type PhoneCallEmptyPredict struct {
	ID int64
}

func (*PhoneCallEmptyPredict) CRC() uint32 {
	return 0x5366c915
}
func (*PhoneCallEmptyPredict) _PhoneCall() {}

type PhoneCallWaitingPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Video         bool     `tl:",omitempty:flags:6,implicit"`
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	Protocol      PhoneCallProtocol
	ReceiveDate   *int32 `tl:",omitempty:flags:0"`
}

func (*PhoneCallWaitingPredict) CRC() uint32 {
	return 0xc5226f17
}
func (*PhoneCallWaitingPredict) _PhoneCall() {}

type PhoneCallRequestedPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Video         bool     `tl:",omitempty:flags:6,implicit"`
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	GAHash        []byte
	Protocol      PhoneCallProtocol
}

func (*PhoneCallRequestedPredict) CRC() uint32 {
	return 0x14b0ed0c
}
func (*PhoneCallRequestedPredict) _PhoneCall() {}

type PhoneCallAcceptedPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Video         bool     `tl:",omitempty:flags:6,implicit"`
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int64
	ParticipantID int64
	GB            []byte
	Protocol      PhoneCallProtocol
}

func (*PhoneCallAcceptedPredict) CRC() uint32 {
	return 0x3660c311
}
func (*PhoneCallAcceptedPredict) _PhoneCall() {}

type PhoneCallPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	P2PAllowed       bool     `tl:",omitempty:flags:5,implicit"`
	Video            bool     `tl:",omitempty:flags:6,implicit"`
	ID               int64
	AccessHash       int64
	Date             int32
	AdminID          int64
	ParticipantID    int64
	GAOrB            []byte
	KeyFingerprint   int64
	Protocol         PhoneCallProtocol
	Connections      []PhoneConnection
	StartDate        int32
	CustomParameters DataJSON `tl:",omitempty:flags:7"`
}

func (*PhoneCallPredict) CRC() uint32 {
	return 0x30535af5
}
func (*PhoneCallPredict) _PhoneCall() {}

type PhoneCallDiscardedPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	NeedRating bool     `tl:",omitempty:flags:2,implicit"`
	NeedDebug  bool     `tl:",omitempty:flags:3,implicit"`
	Video      bool     `tl:",omitempty:flags:6,implicit"`
	ID         int64
	Reason     PhoneCallDiscardReason `tl:",omitempty:flags:0"`
	Duration   *int32                 `tl:",omitempty:flags:1"`
}

func (*PhoneCallDiscardedPredict) CRC() uint32 {
	return 0x50ca4de1
}
func (*PhoneCallDiscardedPredict) _PhoneCall() {}

type PhoneCallProtocol interface {
	tl.Object
	_PhoneCallProtocol()
}

var (
	_ PhoneCallProtocol = (*PhoneCallProtocolPredict)(nil)
)

type PhoneCallProtocolPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	UdpP2P          bool     `tl:",omitempty:flags:0,implicit"`
	UdpReflector    bool     `tl:",omitempty:flags:1,implicit"`
	MinLayer        int32
	MaxLayer        int32
	LibraryVersions []string
}

func (*PhoneCallProtocolPredict) CRC() uint32 {
	return 0xfc878fc8
}
func (*PhoneCallProtocolPredict) _PhoneCallProtocol() {}

type PhoneConnection interface {
	tl.Object
	_PhoneConnection()
}

var (
	_ PhoneConnection = (*PhoneConnectionPredict)(nil)
	_ PhoneConnection = (*PhoneConnectionWebrtcPredict)(nil)
)

type PhoneConnectionPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Tcp     bool     `tl:",omitempty:flags:0,implicit"`
	ID      int64
	Ip      string
	Ipv6    string
	Port    int32
	PeerTag []byte
}

func (*PhoneConnectionPredict) CRC() uint32 {
	return 0x9cc123c7
}
func (*PhoneConnectionPredict) _PhoneConnection() {}

type PhoneConnectionWebrtcPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Turn     bool     `tl:",omitempty:flags:0,implicit"`
	Stun     bool     `tl:",omitempty:flags:1,implicit"`
	ID       int64
	Ip       string
	Ipv6     string
	Port     int32
	Username string
	Password string
}

func (*PhoneConnectionWebrtcPredict) CRC() uint32 {
	return 0x635fe375
}
func (*PhoneConnectionWebrtcPredict) _PhoneConnection() {}

type Photo interface {
	tl.Object
	_Photo()
}

var (
	_ Photo = (*PhotoEmptyPredict)(nil)
	_ Photo = (*PhotoPredict)(nil)
)

type PhotoEmptyPredict struct {
	ID int64
}

func (*PhotoEmptyPredict) CRC() uint32 {
	return 0x2331b22d
}
func (*PhotoEmptyPredict) _Photo() {}

type PhotoPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	HasStickers   bool     `tl:",omitempty:flags:0,implicit"`
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	Sizes         []PhotoSize
	VideoSizes    []VideoSize `tl:",omitempty:flags:1"`
	DcID          int32
}

func (*PhotoPredict) CRC() uint32 {
	return 0xfb197a65
}
func (*PhotoPredict) _Photo() {}

type PhotoSize interface {
	tl.Object
	_PhotoSize()
}

var (
	_ PhotoSize = (*PhotoSizeEmptyPredict)(nil)
	_ PhotoSize = (*PhotoSizePredict)(nil)
	_ PhotoSize = (*PhotoCachedSizePredict)(nil)
	_ PhotoSize = (*PhotoStrippedSizePredict)(nil)
	_ PhotoSize = (*PhotoSizeProgressivePredict)(nil)
	_ PhotoSize = (*PhotoPathSizePredict)(nil)
)

type PhotoSizeEmptyPredict struct {
	Type string
}

func (*PhotoSizeEmptyPredict) CRC() uint32 {
	return 0xe17e23c
}
func (*PhotoSizeEmptyPredict) _PhotoSize() {}

type PhotoSizePredict struct {
	Type string
	W    int32
	H    int32
	Size int32
}

func (*PhotoSizePredict) CRC() uint32 {
	return 0x75c78e60
}
func (*PhotoSizePredict) _PhotoSize() {}

type PhotoCachedSizePredict struct {
	Type  string
	W     int32
	H     int32
	Bytes []byte
}

func (*PhotoCachedSizePredict) CRC() uint32 {
	return 0x21e1ad6
}
func (*PhotoCachedSizePredict) _PhotoSize() {}

type PhotoStrippedSizePredict struct {
	Type  string
	Bytes []byte
}

func (*PhotoStrippedSizePredict) CRC() uint32 {
	return 0xe0b0bc2e
}
func (*PhotoStrippedSizePredict) _PhotoSize() {}

type PhotoSizeProgressivePredict struct {
	Type  string
	W     int32
	H     int32
	Sizes []int32
}

func (*PhotoSizeProgressivePredict) CRC() uint32 {
	return 0xfa3efb95
}
func (*PhotoSizeProgressivePredict) _PhotoSize() {}

type PhotoPathSizePredict struct {
	Type  string
	Bytes []byte
}

func (*PhotoPathSizePredict) CRC() uint32 {
	return 0xd8214d41
}
func (*PhotoPathSizePredict) _PhotoSize() {}

type Poll interface {
	tl.Object
	_Poll()
}

var (
	_ Poll = (*PollPredict)(nil)
)

type PollPredict struct {
	ID             int64
	_              struct{} `tl:"flags,bitflag"`
	Closed         bool     `tl:",omitempty:flags:0,implicit"`
	PublicVoters   bool     `tl:",omitempty:flags:1,implicit"`
	MultipleChoice bool     `tl:",omitempty:flags:2,implicit"`
	Quiz           bool     `tl:",omitempty:flags:3,implicit"`
	Question       TextWithEntities
	Answers        []PollAnswer
	ClosePeriod    *int32 `tl:",omitempty:flags:4"`
	CloseDate      *int32 `tl:",omitempty:flags:5"`
}

func (*PollPredict) CRC() uint32 {
	return 0x58747131
}
func (*PollPredict) _Poll() {}

type PollAnswer interface {
	tl.Object
	_PollAnswer()
}

var (
	_ PollAnswer = (*PollAnswerPredict)(nil)
)

type PollAnswerPredict struct {
	Text   TextWithEntities
	Option []byte
}

func (*PollAnswerPredict) CRC() uint32 {
	return 0xff16e2ca
}
func (*PollAnswerPredict) _PollAnswer() {}

type PollAnswerVoters interface {
	tl.Object
	_PollAnswerVoters()
}

var (
	_ PollAnswerVoters = (*PollAnswerVotersPredict)(nil)
)

type PollAnswerVotersPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Chosen  bool     `tl:",omitempty:flags:0,implicit"`
	Correct bool     `tl:",omitempty:flags:1,implicit"`
	Option  []byte
	Voters  int32
}

func (*PollAnswerVotersPredict) CRC() uint32 {
	return 0x3b6ddad2
}
func (*PollAnswerVotersPredict) _PollAnswerVoters() {}

type PollResults interface {
	tl.Object
	_PollResults()
}

var (
	_ PollResults = (*PollResultsPredict)(nil)
)

type PollResultsPredict struct {
	_                struct{}           `tl:"flags,bitflag"`
	Min              bool               `tl:",omitempty:flags:0,implicit"`
	Results          []PollAnswerVoters `tl:",omitempty:flags:1"`
	TotalVoters      *int32             `tl:",omitempty:flags:2"`
	RecentVoters     []Peer             `tl:",omitempty:flags:3"`
	Solution         *string            `tl:",omitempty:flags:4"`
	SolutionEntities []MessageEntity    `tl:",omitempty:flags:4"`
}

func (*PollResultsPredict) CRC() uint32 {
	return 0x7adf2420
}
func (*PollResultsPredict) _PollResults() {}

type PopularContact interface {
	tl.Object
	_PopularContact()
}

var (
	_ PopularContact = (*PopularContactPredict)(nil)
)

type PopularContactPredict struct {
	ClientID  int64
	Importers int32
}

func (*PopularContactPredict) CRC() uint32 {
	return 0x5ce14175
}
func (*PopularContactPredict) _PopularContact() {}

type PostAddress interface {
	tl.Object
	_PostAddress()
}

var (
	_ PostAddress = (*PostAddressPredict)(nil)
)

type PostAddressPredict struct {
	StreetLine1 string
	StreetLine2 string
	City        string
	State       string
	CountryIso2 string
	PostCode    string
}

func (*PostAddressPredict) CRC() uint32 {
	return 0x1e8caaeb
}
func (*PostAddressPredict) _PostAddress() {}

type PostInteractionCounters interface {
	tl.Object
	_PostInteractionCounters()
}

var (
	_ PostInteractionCounters = (*PostInteractionCountersMessagePredict)(nil)
	_ PostInteractionCounters = (*PostInteractionCountersStoryPredict)(nil)
)

type PostInteractionCountersMessagePredict struct {
	MsgID     int32
	Views     int32
	Forwards  int32
	Reactions int32
}

func (*PostInteractionCountersMessagePredict) CRC() uint32 {
	return 0xe7058e7f
}
func (*PostInteractionCountersMessagePredict) _PostInteractionCounters() {}

type PostInteractionCountersStoryPredict struct {
	StoryID   int32
	Views     int32
	Forwards  int32
	Reactions int32
}

func (*PostInteractionCountersStoryPredict) CRC() uint32 {
	return 0x8a480e27
}
func (*PostInteractionCountersStoryPredict) _PostInteractionCounters() {}

type PremiumGiftCodeOption interface {
	tl.Object
	_PremiumGiftCodeOption()
}

var (
	_ PremiumGiftCodeOption = (*PremiumGiftCodeOptionPredict)(nil)
)

type PremiumGiftCodeOptionPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Users         int32
	Months        int32
	StoreProduct  *string `tl:",omitempty:flags:0"`
	StoreQuantity *int32  `tl:",omitempty:flags:1"`
	Currency      string
	Amount        int64
}

func (*PremiumGiftCodeOptionPredict) CRC() uint32 {
	return 0x257e962b
}
func (*PremiumGiftCodeOptionPredict) _PremiumGiftCodeOption() {}

type PremiumGiftOption interface {
	tl.Object
	_PremiumGiftOption()
}

var (
	_ PremiumGiftOption = (*PremiumGiftOptionPredict)(nil)
)

type PremiumGiftOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Months       int32
	Currency     string
	Amount       int64
	BotURL       string
	StoreProduct *string `tl:",omitempty:flags:0"`
}

func (*PremiumGiftOptionPredict) CRC() uint32 {
	return 0x74c34319
}
func (*PremiumGiftOptionPredict) _PremiumGiftOption() {}

type PremiumSubscriptionOption interface {
	tl.Object
	_PremiumSubscriptionOption()
}

var (
	_ PremiumSubscriptionOption = (*PremiumSubscriptionOptionPredict)(nil)
)

type PremiumSubscriptionOptionPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Current            bool     `tl:",omitempty:flags:1,implicit"`
	CanPurchaseUpgrade bool     `tl:",omitempty:flags:2,implicit"`
	Transaction        *string  `tl:",omitempty:flags:3"`
	Months             int32
	Currency           string
	Amount             int64
	BotURL             string
	StoreProduct       *string `tl:",omitempty:flags:0"`
}

func (*PremiumSubscriptionOptionPredict) CRC() uint32 {
	return 0x5f2d1df2
}
func (*PremiumSubscriptionOptionPredict) _PremiumSubscriptionOption() {}

type PrepaidGiveaway interface {
	tl.Object
	_PrepaidGiveaway()
}

var (
	_ PrepaidGiveaway = (*PrepaidGiveawayPredict)(nil)
)

type PrepaidGiveawayPredict struct {
	ID       int64
	Months   int32
	Quantity int32
	Date     int32
}

func (*PrepaidGiveawayPredict) CRC() uint32 {
	return 0xb2539d54
}
func (*PrepaidGiveawayPredict) _PrepaidGiveaway() {}

type PrivacyRule interface {
	tl.Object
	_PrivacyRule()
}

var (
	_ PrivacyRule = (*PrivacyValueAllowContactsPredict)(nil)
	_ PrivacyRule = (*PrivacyValueAllowAllPredict)(nil)
	_ PrivacyRule = (*PrivacyValueAllowUsersPredict)(nil)
	_ PrivacyRule = (*PrivacyValueDisallowContactsPredict)(nil)
	_ PrivacyRule = (*PrivacyValueDisallowAllPredict)(nil)
	_ PrivacyRule = (*PrivacyValueDisallowUsersPredict)(nil)
	_ PrivacyRule = (*PrivacyValueAllowChatParticipantsPredict)(nil)
	_ PrivacyRule = (*PrivacyValueDisallowChatParticipantsPredict)(nil)
	_ PrivacyRule = (*PrivacyValueAllowCloseFriendsPredict)(nil)
	_ PrivacyRule = (*PrivacyValueAllowPremiumPredict)(nil)
)

type PrivacyValueAllowContactsPredict struct{}

func (*PrivacyValueAllowContactsPredict) CRC() uint32 {
	return 0xfffe1bac
}
func (*PrivacyValueAllowContactsPredict) _PrivacyRule() {}

type PrivacyValueAllowAllPredict struct{}

func (*PrivacyValueAllowAllPredict) CRC() uint32 {
	return 0x65427b82
}
func (*PrivacyValueAllowAllPredict) _PrivacyRule() {}

type PrivacyValueAllowUsersPredict struct {
	Users []int64
}

func (*PrivacyValueAllowUsersPredict) CRC() uint32 {
	return 0xb8905fb2
}
func (*PrivacyValueAllowUsersPredict) _PrivacyRule() {}

type PrivacyValueDisallowContactsPredict struct{}

func (*PrivacyValueDisallowContactsPredict) CRC() uint32 {
	return 0xf888fa1a
}
func (*PrivacyValueDisallowContactsPredict) _PrivacyRule() {}

type PrivacyValueDisallowAllPredict struct{}

func (*PrivacyValueDisallowAllPredict) CRC() uint32 {
	return 0x8b73e763
}
func (*PrivacyValueDisallowAllPredict) _PrivacyRule() {}

type PrivacyValueDisallowUsersPredict struct {
	Users []int64
}

func (*PrivacyValueDisallowUsersPredict) CRC() uint32 {
	return 0xe4621141
}
func (*PrivacyValueDisallowUsersPredict) _PrivacyRule() {}

type PrivacyValueAllowChatParticipantsPredict struct {
	Chats []int64
}

func (*PrivacyValueAllowChatParticipantsPredict) CRC() uint32 {
	return 0x6b134e8e
}
func (*PrivacyValueAllowChatParticipantsPredict) _PrivacyRule() {}

type PrivacyValueDisallowChatParticipantsPredict struct {
	Chats []int64
}

func (*PrivacyValueDisallowChatParticipantsPredict) CRC() uint32 {
	return 0x41c87565
}
func (*PrivacyValueDisallowChatParticipantsPredict) _PrivacyRule() {}

type PrivacyValueAllowCloseFriendsPredict struct{}

func (*PrivacyValueAllowCloseFriendsPredict) CRC() uint32 {
	return 0xf7e8d89b
}
func (*PrivacyValueAllowCloseFriendsPredict) _PrivacyRule() {}

type PrivacyValueAllowPremiumPredict struct{}

func (*PrivacyValueAllowPremiumPredict) CRC() uint32 {
	return 0xece9814b
}
func (*PrivacyValueAllowPremiumPredict) _PrivacyRule() {}

type PublicForward interface {
	tl.Object
	_PublicForward()
}

var (
	_ PublicForward = (*PublicForwardMessagePredict)(nil)
	_ PublicForward = (*PublicForwardStoryPredict)(nil)
)

type PublicForwardMessagePredict struct {
	Message Message
}

func (*PublicForwardMessagePredict) CRC() uint32 {
	return 0x1f2bf4a
}
func (*PublicForwardMessagePredict) _PublicForward() {}

type PublicForwardStoryPredict struct {
	Peer  Peer
	Story StoryItem
}

func (*PublicForwardStoryPredict) CRC() uint32 {
	return 0xedf3add0
}
func (*PublicForwardStoryPredict) _PublicForward() {}

type QuickReply interface {
	tl.Object
	_QuickReply()
}

var (
	_ QuickReply = (*QuickReplyPredict)(nil)
)

type QuickReplyPredict struct {
	ShortcutID int32
	Shortcut   string
	TopMessage int32
	Count      int32
}

func (*QuickReplyPredict) CRC() uint32 {
	return 0x697102b
}
func (*QuickReplyPredict) _QuickReply() {}

type Reaction interface {
	tl.Object
	_Reaction()
}

var (
	_ Reaction = (*ReactionEmptyPredict)(nil)
	_ Reaction = (*ReactionEmojiPredict)(nil)
	_ Reaction = (*ReactionCustomEmojiPredict)(nil)
)

type ReactionEmptyPredict struct{}

func (*ReactionEmptyPredict) CRC() uint32 {
	return 0x79f5d419
}
func (*ReactionEmptyPredict) _Reaction() {}

type ReactionEmojiPredict struct {
	Emoticon string
}

func (*ReactionEmojiPredict) CRC() uint32 {
	return 0x1b2286b8
}
func (*ReactionEmojiPredict) _Reaction() {}

type ReactionCustomEmojiPredict struct {
	DocumentID int64
}

func (*ReactionCustomEmojiPredict) CRC() uint32 {
	return 0x8935fc73
}
func (*ReactionCustomEmojiPredict) _Reaction() {}

type ReactionCount interface {
	tl.Object
	_ReactionCount()
}

var (
	_ ReactionCount = (*ReactionCountPredict)(nil)
)

type ReactionCountPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ChosenOrder *int32   `tl:",omitempty:flags:0"`
	Reaction    Reaction
	Count       int32
}

func (*ReactionCountPredict) CRC() uint32 {
	return 0xa3d1cb80
}
func (*ReactionCountPredict) _ReactionCount() {}

type ReactionsNotifySettings interface {
	tl.Object
	_ReactionsNotifySettings()
}

var (
	_ ReactionsNotifySettings = (*ReactionsNotifySettingsPredict)(nil)
)

type ReactionsNotifySettingsPredict struct {
	_                  struct{}                  `tl:"flags,bitflag"`
	MessagesNotifyFrom ReactionNotificationsFrom `tl:",omitempty:flags:0"`
	StoriesNotifyFrom  ReactionNotificationsFrom `tl:",omitempty:flags:1"`
	Sound              NotificationSound
	ShowPreviews       bool
}

func (*ReactionsNotifySettingsPredict) CRC() uint32 {
	return 0x56e34970
}
func (*ReactionsNotifySettingsPredict) _ReactionsNotifySettings() {}

type ReadParticipantDate interface {
	tl.Object
	_ReadParticipantDate()
}

var (
	_ ReadParticipantDate = (*ReadParticipantDatePredict)(nil)
)

type ReadParticipantDatePredict struct {
	UserID int64
	Date   int32
}

func (*ReadParticipantDatePredict) CRC() uint32 {
	return 0x4a4ff172
}
func (*ReadParticipantDatePredict) _ReadParticipantDate() {}

type ReceivedNotifyMessage interface {
	tl.Object
	_ReceivedNotifyMessage()
}

var (
	_ ReceivedNotifyMessage = (*ReceivedNotifyMessagePredict)(nil)
)

type ReceivedNotifyMessagePredict struct {
	ID    int32
	Flags int32
}

func (*ReceivedNotifyMessagePredict) CRC() uint32 {
	return 0xa384b779
}
func (*ReceivedNotifyMessagePredict) _ReceivedNotifyMessage() {}

type RecentMeURL interface {
	tl.Object
	_RecentMeURL()
}

var (
	_ RecentMeURL = (*RecentMeURLUnknownPredict)(nil)
	_ RecentMeURL = (*RecentMeURLUserPredict)(nil)
	_ RecentMeURL = (*RecentMeURLChatPredict)(nil)
	_ RecentMeURL = (*RecentMeURLChatInvitePredict)(nil)
	_ RecentMeURL = (*RecentMeURLStickerSetPredict)(nil)
)

type RecentMeURLUnknownPredict struct {
	URL string
}

func (*RecentMeURLUnknownPredict) CRC() uint32 {
	return 0x46e1d13d
}
func (*RecentMeURLUnknownPredict) _RecentMeURL() {}

type RecentMeURLUserPredict struct {
	URL    string
	UserID int64
}

func (*RecentMeURLUserPredict) CRC() uint32 {
	return 0xb92c09e2
}
func (*RecentMeURLUserPredict) _RecentMeURL() {}

type RecentMeURLChatPredict struct {
	URL    string
	ChatID int64
}

func (*RecentMeURLChatPredict) CRC() uint32 {
	return 0xb2da71d2
}
func (*RecentMeURLChatPredict) _RecentMeURL() {}

type RecentMeURLChatInvitePredict struct {
	URL        string
	ChatInvite ChatInvite
}

func (*RecentMeURLChatInvitePredict) CRC() uint32 {
	return 0xeb49081d
}
func (*RecentMeURLChatInvitePredict) _RecentMeURL() {}

type RecentMeURLStickerSetPredict struct {
	URL string
	Set StickerSetCovered
}

func (*RecentMeURLStickerSetPredict) CRC() uint32 {
	return 0xbc0a57dc
}
func (*RecentMeURLStickerSetPredict) _RecentMeURL() {}

type ReplyMarkup interface {
	tl.Object
	_ReplyMarkup()
}

var (
	_ ReplyMarkup = (*ReplyKeyboardHidePredict)(nil)
	_ ReplyMarkup = (*ReplyKeyboardForceReplyPredict)(nil)
	_ ReplyMarkup = (*ReplyKeyboardMarkupPredict)(nil)
	_ ReplyMarkup = (*ReplyInlineMarkupPredict)(nil)
)

type ReplyKeyboardHidePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Selective bool     `tl:",omitempty:flags:2,implicit"`
}

func (*ReplyKeyboardHidePredict) CRC() uint32 {
	return 0xa03e5b85
}
func (*ReplyKeyboardHidePredict) _ReplyMarkup() {}

type ReplyKeyboardForceReplyPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	SingleUse   bool     `tl:",omitempty:flags:1,implicit"`
	Selective   bool     `tl:",omitempty:flags:2,implicit"`
	Placeholder *string  `tl:",omitempty:flags:3"`
}

func (*ReplyKeyboardForceReplyPredict) CRC() uint32 {
	return 0x86b40b08
}
func (*ReplyKeyboardForceReplyPredict) _ReplyMarkup() {}

type ReplyKeyboardMarkupPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Resize      bool     `tl:",omitempty:flags:0,implicit"`
	SingleUse   bool     `tl:",omitempty:flags:1,implicit"`
	Selective   bool     `tl:",omitempty:flags:2,implicit"`
	Persistent  bool     `tl:",omitempty:flags:4,implicit"`
	Rows        []KeyboardButtonRow
	Placeholder *string `tl:",omitempty:flags:3"`
}

func (*ReplyKeyboardMarkupPredict) CRC() uint32 {
	return 0x85dd99d1
}
func (*ReplyKeyboardMarkupPredict) _ReplyMarkup() {}

type ReplyInlineMarkupPredict struct {
	Rows []KeyboardButtonRow
}

func (*ReplyInlineMarkupPredict) CRC() uint32 {
	return 0x48a30254
}
func (*ReplyInlineMarkupPredict) _ReplyMarkup() {}

type RequestPeerType interface {
	tl.Object
	_RequestPeerType()
}

var (
	_ RequestPeerType = (*RequestPeerTypeUserPredict)(nil)
	_ RequestPeerType = (*RequestPeerTypeChatPredict)(nil)
	_ RequestPeerType = (*RequestPeerTypeBroadcastPredict)(nil)
)

type RequestPeerTypeUserPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Bot     *bool    `tl:",omitempty:flags:0"`
	Premium *bool    `tl:",omitempty:flags:1"`
}

func (*RequestPeerTypeUserPredict) CRC() uint32 {
	return 0x5f3b8a00
}
func (*RequestPeerTypeUserPredict) _RequestPeerType() {}

type RequestPeerTypeChatPredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	Creator         bool            `tl:",omitempty:flags:0,implicit"`
	BotParticipant  bool            `tl:",omitempty:flags:5,implicit"`
	HasUsername     *bool           `tl:",omitempty:flags:3"`
	Forum           *bool           `tl:",omitempty:flags:4"`
	UserAdminRights ChatAdminRights `tl:",omitempty:flags:1"`
	BotAdminRights  ChatAdminRights `tl:",omitempty:flags:2"`
}

func (*RequestPeerTypeChatPredict) CRC() uint32 {
	return 0xc9f06e1b
}
func (*RequestPeerTypeChatPredict) _RequestPeerType() {}

type RequestPeerTypeBroadcastPredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	Creator         bool            `tl:",omitempty:flags:0,implicit"`
	HasUsername     *bool           `tl:",omitempty:flags:3"`
	UserAdminRights ChatAdminRights `tl:",omitempty:flags:1"`
	BotAdminRights  ChatAdminRights `tl:",omitempty:flags:2"`
}

func (*RequestPeerTypeBroadcastPredict) CRC() uint32 {
	return 0x339bef6c
}
func (*RequestPeerTypeBroadcastPredict) _RequestPeerType() {}

type RequestedPeer interface {
	tl.Object
	_RequestedPeer()
}

var (
	_ RequestedPeer = (*RequestedPeerUserPredict)(nil)
	_ RequestedPeer = (*RequestedPeerChatPredict)(nil)
	_ RequestedPeer = (*RequestedPeerChannelPredict)(nil)
)

type RequestedPeerUserPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	UserID    int64
	FirstName *string `tl:",omitempty:flags:0"`
	LastName  *string `tl:",omitempty:flags:0"`
	Username  *string `tl:",omitempty:flags:1"`
	Photo     Photo   `tl:",omitempty:flags:2"`
}

func (*RequestedPeerUserPredict) CRC() uint32 {
	return 0xd62ff46a
}
func (*RequestedPeerUserPredict) _RequestedPeer() {}

type RequestedPeerChatPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	ChatID int64
	Title  *string `tl:",omitempty:flags:0"`
	Photo  Photo   `tl:",omitempty:flags:2"`
}

func (*RequestedPeerChatPredict) CRC() uint32 {
	return 0x7307544f
}
func (*RequestedPeerChatPredict) _RequestedPeer() {}

type RequestedPeerChannelPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	Title     *string `tl:",omitempty:flags:0"`
	Username  *string `tl:",omitempty:flags:1"`
	Photo     Photo   `tl:",omitempty:flags:2"`
}

func (*RequestedPeerChannelPredict) CRC() uint32 {
	return 0x8ba403e4
}
func (*RequestedPeerChannelPredict) _RequestedPeer() {}

type RestrictionReason interface {
	tl.Object
	_RestrictionReason()
}

var (
	_ RestrictionReason = (*RestrictionReasonPredict)(nil)
)

type RestrictionReasonPredict struct {
	Platform string
	Reason   string
	Text     string
}

func (*RestrictionReasonPredict) CRC() uint32 {
	return 0xd072acb4
}
func (*RestrictionReasonPredict) _RestrictionReason() {}

type RichText interface {
	tl.Object
	_RichText()
}

var (
	_ RichText = (*TextEmptyPredict)(nil)
	_ RichText = (*TextPlainPredict)(nil)
	_ RichText = (*TextBoldPredict)(nil)
	_ RichText = (*TextItalicPredict)(nil)
	_ RichText = (*TextUnderlinePredict)(nil)
	_ RichText = (*TextStrikePredict)(nil)
	_ RichText = (*TextFixedPredict)(nil)
	_ RichText = (*TextURLPredict)(nil)
	_ RichText = (*TextEmailPredict)(nil)
	_ RichText = (*TextConcatPredict)(nil)
	_ RichText = (*TextSubscriptPredict)(nil)
	_ RichText = (*TextSuperscriptPredict)(nil)
	_ RichText = (*TextMarkedPredict)(nil)
	_ RichText = (*TextPhonePredict)(nil)
	_ RichText = (*TextImagePredict)(nil)
	_ RichText = (*TextAnchorPredict)(nil)
)

type TextEmptyPredict struct{}

func (*TextEmptyPredict) CRC() uint32 {
	return 0xdc3d824f
}
func (*TextEmptyPredict) _RichText() {}

type TextPlainPredict struct {
	Text string
}

func (*TextPlainPredict) CRC() uint32 {
	return 0x744694e0
}
func (*TextPlainPredict) _RichText() {}

type TextBoldPredict struct {
	Text RichText
}

func (*TextBoldPredict) CRC() uint32 {
	return 0x6724abc4
}
func (*TextBoldPredict) _RichText() {}

type TextItalicPredict struct {
	Text RichText
}

func (*TextItalicPredict) CRC() uint32 {
	return 0xd912a59c
}
func (*TextItalicPredict) _RichText() {}

type TextUnderlinePredict struct {
	Text RichText
}

func (*TextUnderlinePredict) CRC() uint32 {
	return 0xc12622c4
}
func (*TextUnderlinePredict) _RichText() {}

type TextStrikePredict struct {
	Text RichText
}

func (*TextStrikePredict) CRC() uint32 {
	return 0x9bf8bb95
}
func (*TextStrikePredict) _RichText() {}

type TextFixedPredict struct {
	Text RichText
}

func (*TextFixedPredict) CRC() uint32 {
	return 0x6c3f19b9
}
func (*TextFixedPredict) _RichText() {}

type TextURLPredict struct {
	Text      RichText
	URL       string
	WebpageID int64
}

func (*TextURLPredict) CRC() uint32 {
	return 0x3c2884c1
}
func (*TextURLPredict) _RichText() {}

type TextEmailPredict struct {
	Text  RichText
	Email string
}

func (*TextEmailPredict) CRC() uint32 {
	return 0xde5a0dd6
}
func (*TextEmailPredict) _RichText() {}

type TextConcatPredict struct {
	Texts []RichText
}

func (*TextConcatPredict) CRC() uint32 {
	return 0x7e6260d7
}
func (*TextConcatPredict) _RichText() {}

type TextSubscriptPredict struct {
	Text RichText
}

func (*TextSubscriptPredict) CRC() uint32 {
	return 0xed6a8504
}
func (*TextSubscriptPredict) _RichText() {}

type TextSuperscriptPredict struct {
	Text RichText
}

func (*TextSuperscriptPredict) CRC() uint32 {
	return 0xc7fb5e01
}
func (*TextSuperscriptPredict) _RichText() {}

type TextMarkedPredict struct {
	Text RichText
}

func (*TextMarkedPredict) CRC() uint32 {
	return 0x34b8621
}
func (*TextMarkedPredict) _RichText() {}

type TextPhonePredict struct {
	Text  RichText
	Phone string
}

func (*TextPhonePredict) CRC() uint32 {
	return 0x1ccb966a
}
func (*TextPhonePredict) _RichText() {}

type TextImagePredict struct {
	DocumentID int64
	W          int32
	H          int32
}

func (*TextImagePredict) CRC() uint32 {
	return 0x81ccf4f
}
func (*TextImagePredict) _RichText() {}

type TextAnchorPredict struct {
	Text RichText
	Name string
}

func (*TextAnchorPredict) CRC() uint32 {
	return 0x35553762
}
func (*TextAnchorPredict) _RichText() {}

type SavedContact interface {
	tl.Object
	_SavedContact()
}

var (
	_ SavedContact = (*SavedPhoneContactPredict)(nil)
)

type SavedPhoneContactPredict struct {
	Phone     string
	FirstName string
	LastName  string
	Date      int32
}

func (*SavedPhoneContactPredict) CRC() uint32 {
	return 0x1142bd56
}
func (*SavedPhoneContactPredict) _SavedContact() {}

type SavedDialog interface {
	tl.Object
	_SavedDialog()
}

var (
	_ SavedDialog = (*SavedDialogPredict)(nil)
)

type SavedDialogPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Pinned     bool     `tl:",omitempty:flags:2,implicit"`
	Peer       Peer
	TopMessage int32
}

func (*SavedDialogPredict) CRC() uint32 {
	return 0xbd87cb6c
}
func (*SavedDialogPredict) _SavedDialog() {}

type SavedReactionTag interface {
	tl.Object
	_SavedReactionTag()
}

var (
	_ SavedReactionTag = (*SavedReactionTagPredict)(nil)
)

type SavedReactionTagPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Reaction Reaction
	Title    *string `tl:",omitempty:flags:0"`
	Count    int32
}

func (*SavedReactionTagPredict) CRC() uint32 {
	return 0xcb6ff828
}
func (*SavedReactionTagPredict) _SavedReactionTag() {}

type SearchResultsCalendarPeriod interface {
	tl.Object
	_SearchResultsCalendarPeriod()
}

var (
	_ SearchResultsCalendarPeriod = (*SearchResultsCalendarPeriodPredict)(nil)
)

type SearchResultsCalendarPeriodPredict struct {
	Date     int32
	MinMsgID int32
	MaxMsgID int32
	Count    int32
}

func (*SearchResultsCalendarPeriodPredict) CRC() uint32 {
	return 0xc9b0539f
}
func (*SearchResultsCalendarPeriodPredict) _SearchResultsCalendarPeriod() {}

type SearchResultsPosition interface {
	tl.Object
	_SearchResultsPosition()
}

var (
	_ SearchResultsPosition = (*SearchResultPositionPredict)(nil)
)

type SearchResultPositionPredict struct {
	MsgID  int32
	Date   int32
	Offset int32
}

func (*SearchResultPositionPredict) CRC() uint32 {
	return 0x7f648b67
}
func (*SearchResultPositionPredict) _SearchResultsPosition() {}

type SecureCredentialsEncrypted interface {
	tl.Object
	_SecureCredentialsEncrypted()
}

var (
	_ SecureCredentialsEncrypted = (*SecureCredentialsEncryptedPredict)(nil)
)

type SecureCredentialsEncryptedPredict struct {
	Data   []byte
	Hash   []byte
	Secret []byte
}

func (*SecureCredentialsEncryptedPredict) CRC() uint32 {
	return 0x33f0ea47
}
func (*SecureCredentialsEncryptedPredict) _SecureCredentialsEncrypted() {}

type SecureData interface {
	tl.Object
	_SecureData()
}

var (
	_ SecureData = (*SecureDataPredict)(nil)
)

type SecureDataPredict struct {
	Data     []byte
	DataHash []byte
	Secret   []byte
}

func (*SecureDataPredict) CRC() uint32 {
	return 0x8aeabec3
}
func (*SecureDataPredict) _SecureData() {}

type SecureFile interface {
	tl.Object
	_SecureFile()
}

var (
	_ SecureFile = (*SecureFileEmptyPredict)(nil)
	_ SecureFile = (*SecureFilePredict)(nil)
)

type SecureFileEmptyPredict struct{}

func (*SecureFileEmptyPredict) CRC() uint32 {
	return 0x64199744
}
func (*SecureFileEmptyPredict) _SecureFile() {}

type SecureFilePredict struct {
	ID         int64
	AccessHash int64
	Size       int64
	DcID       int32
	Date       int32
	FileHash   []byte
	Secret     []byte
}

func (*SecureFilePredict) CRC() uint32 {
	return 0x7d09c27e
}
func (*SecureFilePredict) _SecureFile() {}

type SecurePasswordKdfAlgo interface {
	tl.Object
	_SecurePasswordKdfAlgo()
}

var (
	_ SecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoUnknownPredict)(nil)
	_ SecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict)(nil)
	_ SecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoSHA512Predict)(nil)
)

type SecurePasswordKdfAlgoUnknownPredict struct{}

func (*SecurePasswordKdfAlgoUnknownPredict) CRC() uint32 {
	return 0x4a8537
}
func (*SecurePasswordKdfAlgoUnknownPredict) _SecurePasswordKdfAlgo() {}

type SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict struct {
	Salt []byte
}

func (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict) CRC() uint32 {
	return 0xbbf2dda0
}
func (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict) _SecurePasswordKdfAlgo() {}

type SecurePasswordKdfAlgoSHA512Predict struct {
	Salt []byte
}

func (*SecurePasswordKdfAlgoSHA512Predict) CRC() uint32 {
	return 0x86471d92
}
func (*SecurePasswordKdfAlgoSHA512Predict) _SecurePasswordKdfAlgo() {}

type SecurePlainData interface {
	tl.Object
	_SecurePlainData()
}

var (
	_ SecurePlainData = (*SecurePlainPhonePredict)(nil)
	_ SecurePlainData = (*SecurePlainEmailPredict)(nil)
)

type SecurePlainPhonePredict struct {
	Phone string
}

func (*SecurePlainPhonePredict) CRC() uint32 {
	return 0x7d6099dd
}
func (*SecurePlainPhonePredict) _SecurePlainData() {}

type SecurePlainEmailPredict struct {
	Email string
}

func (*SecurePlainEmailPredict) CRC() uint32 {
	return 0x21ec5a5f
}
func (*SecurePlainEmailPredict) _SecurePlainData() {}

type SecureRequiredType interface {
	tl.Object
	_SecureRequiredType()
}

var (
	_ SecureRequiredType = (*SecureRequiredTypePredict)(nil)
	_ SecureRequiredType = (*SecureRequiredTypeOneOfPredict)(nil)
)

type SecureRequiredTypePredict struct {
	_                   struct{} `tl:"flags,bitflag"`
	NativeNames         bool     `tl:",omitempty:flags:0,implicit"`
	SelfieRequired      bool     `tl:",omitempty:flags:1,implicit"`
	TranslationRequired bool     `tl:",omitempty:flags:2,implicit"`
	Type                SecureValueType
}

func (*SecureRequiredTypePredict) CRC() uint32 {
	return 0x829d99da
}
func (*SecureRequiredTypePredict) _SecureRequiredType() {}

type SecureRequiredTypeOneOfPredict struct {
	Types []SecureRequiredType
}

func (*SecureRequiredTypeOneOfPredict) CRC() uint32 {
	return 0x27477b4
}
func (*SecureRequiredTypeOneOfPredict) _SecureRequiredType() {}

type SecureSecretSettings interface {
	tl.Object
	_SecureSecretSettings()
}

var (
	_ SecureSecretSettings = (*SecureSecretSettingsPredict)(nil)
)

type SecureSecretSettingsPredict struct {
	SecureAlgo     SecurePasswordKdfAlgo
	SecureSecret   []byte
	SecureSecretID int64
}

func (*SecureSecretSettingsPredict) CRC() uint32 {
	return 0x1527bcac
}
func (*SecureSecretSettingsPredict) _SecureSecretSettings() {}

type SecureValue interface {
	tl.Object
	_SecureValue()
}

var (
	_ SecureValue = (*SecureValuePredict)(nil)
)

type SecureValuePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Type        SecureValueType
	Data        SecureData      `tl:",omitempty:flags:0"`
	FrontSide   SecureFile      `tl:",omitempty:flags:1"`
	ReverseSide SecureFile      `tl:",omitempty:flags:2"`
	Selfie      SecureFile      `tl:",omitempty:flags:3"`
	Translation []SecureFile    `tl:",omitempty:flags:6"`
	Files       []SecureFile    `tl:",omitempty:flags:4"`
	PlainData   SecurePlainData `tl:",omitempty:flags:5"`
	Hash        []byte
}

func (*SecureValuePredict) CRC() uint32 {
	return 0x187fa0ca
}
func (*SecureValuePredict) _SecureValue() {}

type SecureValueError interface {
	tl.Object
	_SecureValueError()
}

var (
	_ SecureValueError = (*SecureValueErrorDataPredict)(nil)
	_ SecureValueError = (*SecureValueErrorFrontSidePredict)(nil)
	_ SecureValueError = (*SecureValueErrorReverseSidePredict)(nil)
	_ SecureValueError = (*SecureValueErrorSelfiePredict)(nil)
	_ SecureValueError = (*SecureValueErrorFilePredict)(nil)
	_ SecureValueError = (*SecureValueErrorFilesPredict)(nil)
	_ SecureValueError = (*SecureValueErrorPredict)(nil)
	_ SecureValueError = (*SecureValueErrorTranslationFilePredict)(nil)
	_ SecureValueError = (*SecureValueErrorTranslationFilesPredict)(nil)
)

type SecureValueErrorDataPredict struct {
	Type     SecureValueType
	DataHash []byte
	Field    string
	Text     string
}

func (*SecureValueErrorDataPredict) CRC() uint32 {
	return 0xe8a40bd9
}
func (*SecureValueErrorDataPredict) _SecureValueError() {}

type SecureValueErrorFrontSidePredict struct {
	Type     SecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorFrontSidePredict) CRC() uint32 {
	return 0xbe3dfa
}
func (*SecureValueErrorFrontSidePredict) _SecureValueError() {}

type SecureValueErrorReverseSidePredict struct {
	Type     SecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorReverseSidePredict) CRC() uint32 {
	return 0x868a2aa5
}
func (*SecureValueErrorReverseSidePredict) _SecureValueError() {}

type SecureValueErrorSelfiePredict struct {
	Type     SecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorSelfiePredict) CRC() uint32 {
	return 0xe537ced6
}
func (*SecureValueErrorSelfiePredict) _SecureValueError() {}

type SecureValueErrorFilePredict struct {
	Type     SecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorFilePredict) CRC() uint32 {
	return 0x7a700873
}
func (*SecureValueErrorFilePredict) _SecureValueError() {}

type SecureValueErrorFilesPredict struct {
	Type     SecureValueType
	FileHash [][]byte
	Text     string
}

func (*SecureValueErrorFilesPredict) CRC() uint32 {
	return 0x666220e9
}
func (*SecureValueErrorFilesPredict) _SecureValueError() {}

type SecureValueErrorPredict struct {
	Type SecureValueType
	Hash []byte
	Text string
}

func (*SecureValueErrorPredict) CRC() uint32 {
	return 0x869d758f
}
func (*SecureValueErrorPredict) _SecureValueError() {}

type SecureValueErrorTranslationFilePredict struct {
	Type     SecureValueType
	FileHash []byte
	Text     string
}

func (*SecureValueErrorTranslationFilePredict) CRC() uint32 {
	return 0xa1144770
}
func (*SecureValueErrorTranslationFilePredict) _SecureValueError() {}

type SecureValueErrorTranslationFilesPredict struct {
	Type     SecureValueType
	FileHash [][]byte
	Text     string
}

func (*SecureValueErrorTranslationFilesPredict) CRC() uint32 {
	return 0x34636dd8
}
func (*SecureValueErrorTranslationFilesPredict) _SecureValueError() {}

type SecureValueHash interface {
	tl.Object
	_SecureValueHash()
}

var (
	_ SecureValueHash = (*SecureValueHashPredict)(nil)
)

type SecureValueHashPredict struct {
	Type SecureValueType
	Hash []byte
}

func (*SecureValueHashPredict) CRC() uint32 {
	return 0xed1ecdb0
}
func (*SecureValueHashPredict) _SecureValueHash() {}

type SendAsPeer interface {
	tl.Object
	_SendAsPeer()
}

var (
	_ SendAsPeer = (*SendAsPeerPredict)(nil)
)

type SendAsPeerPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	PremiumRequired bool     `tl:",omitempty:flags:0,implicit"`
	Peer            Peer
}

func (*SendAsPeerPredict) CRC() uint32 {
	return 0xb81c7034
}
func (*SendAsPeerPredict) _SendAsPeer() {}

type SendMessageAction interface {
	tl.Object
	_SendMessageAction()
}

var (
	_ SendMessageAction = (*SendMessageTypingActionPredict)(nil)
	_ SendMessageAction = (*SendMessageCancelActionPredict)(nil)
	_ SendMessageAction = (*SendMessageRecordVideoActionPredict)(nil)
	_ SendMessageAction = (*SendMessageUploadVideoActionPredict)(nil)
	_ SendMessageAction = (*SendMessageRecordAudioActionPredict)(nil)
	_ SendMessageAction = (*SendMessageUploadAudioActionPredict)(nil)
	_ SendMessageAction = (*SendMessageUploadPhotoActionPredict)(nil)
	_ SendMessageAction = (*SendMessageUploadDocumentActionPredict)(nil)
	_ SendMessageAction = (*SendMessageGeoLocationActionPredict)(nil)
	_ SendMessageAction = (*SendMessageChooseContactActionPredict)(nil)
	_ SendMessageAction = (*SendMessageGamePlayActionPredict)(nil)
	_ SendMessageAction = (*SendMessageRecordRoundActionPredict)(nil)
	_ SendMessageAction = (*SendMessageUploadRoundActionPredict)(nil)
	_ SendMessageAction = (*SpeakingInGroupCallActionPredict)(nil)
	_ SendMessageAction = (*SendMessageHistoryImportActionPredict)(nil)
	_ SendMessageAction = (*SendMessageChooseStickerActionPredict)(nil)
	_ SendMessageAction = (*SendMessageEmojiInteractionPredict)(nil)
	_ SendMessageAction = (*SendMessageEmojiInteractionSeenPredict)(nil)
)

type SendMessageTypingActionPredict struct{}

func (*SendMessageTypingActionPredict) CRC() uint32 {
	return 0x16bf744e
}
func (*SendMessageTypingActionPredict) _SendMessageAction() {}

type SendMessageCancelActionPredict struct{}

func (*SendMessageCancelActionPredict) CRC() uint32 {
	return 0xfd5ec8f5
}
func (*SendMessageCancelActionPredict) _SendMessageAction() {}

type SendMessageRecordVideoActionPredict struct{}

func (*SendMessageRecordVideoActionPredict) CRC() uint32 {
	return 0xa187d66f
}
func (*SendMessageRecordVideoActionPredict) _SendMessageAction() {}

type SendMessageUploadVideoActionPredict struct {
	Progress int32
}

func (*SendMessageUploadVideoActionPredict) CRC() uint32 {
	return 0xe9763aec
}
func (*SendMessageUploadVideoActionPredict) _SendMessageAction() {}

type SendMessageRecordAudioActionPredict struct{}

func (*SendMessageRecordAudioActionPredict) CRC() uint32 {
	return 0xd52f73f7
}
func (*SendMessageRecordAudioActionPredict) _SendMessageAction() {}

type SendMessageUploadAudioActionPredict struct {
	Progress int32
}

func (*SendMessageUploadAudioActionPredict) CRC() uint32 {
	return 0xf351d7ab
}
func (*SendMessageUploadAudioActionPredict) _SendMessageAction() {}

type SendMessageUploadPhotoActionPredict struct {
	Progress int32
}

func (*SendMessageUploadPhotoActionPredict) CRC() uint32 {
	return 0xd1d34a26
}
func (*SendMessageUploadPhotoActionPredict) _SendMessageAction() {}

type SendMessageUploadDocumentActionPredict struct {
	Progress int32
}

func (*SendMessageUploadDocumentActionPredict) CRC() uint32 {
	return 0xaa0cd9e4
}
func (*SendMessageUploadDocumentActionPredict) _SendMessageAction() {}

type SendMessageGeoLocationActionPredict struct{}

func (*SendMessageGeoLocationActionPredict) CRC() uint32 {
	return 0x176f8ba1
}
func (*SendMessageGeoLocationActionPredict) _SendMessageAction() {}

type SendMessageChooseContactActionPredict struct{}

func (*SendMessageChooseContactActionPredict) CRC() uint32 {
	return 0x628cbc6f
}
func (*SendMessageChooseContactActionPredict) _SendMessageAction() {}

type SendMessageGamePlayActionPredict struct{}

func (*SendMessageGamePlayActionPredict) CRC() uint32 {
	return 0xdd6a8f48
}
func (*SendMessageGamePlayActionPredict) _SendMessageAction() {}

type SendMessageRecordRoundActionPredict struct{}

func (*SendMessageRecordRoundActionPredict) CRC() uint32 {
	return 0x88f27fbc
}
func (*SendMessageRecordRoundActionPredict) _SendMessageAction() {}

type SendMessageUploadRoundActionPredict struct {
	Progress int32
}

func (*SendMessageUploadRoundActionPredict) CRC() uint32 {
	return 0x243e1c66
}
func (*SendMessageUploadRoundActionPredict) _SendMessageAction() {}

type SpeakingInGroupCallActionPredict struct{}

func (*SpeakingInGroupCallActionPredict) CRC() uint32 {
	return 0xd92c2285
}
func (*SpeakingInGroupCallActionPredict) _SendMessageAction() {}

type SendMessageHistoryImportActionPredict struct {
	Progress int32
}

func (*SendMessageHistoryImportActionPredict) CRC() uint32 {
	return 0xdbda9246
}
func (*SendMessageHistoryImportActionPredict) _SendMessageAction() {}

type SendMessageChooseStickerActionPredict struct{}

func (*SendMessageChooseStickerActionPredict) CRC() uint32 {
	return 0xb05ac6b1
}
func (*SendMessageChooseStickerActionPredict) _SendMessageAction() {}

type SendMessageEmojiInteractionPredict struct {
	Emoticon    string
	MsgID       int32
	Interaction DataJSON
}

func (*SendMessageEmojiInteractionPredict) CRC() uint32 {
	return 0x25972bcb
}
func (*SendMessageEmojiInteractionPredict) _SendMessageAction() {}

type SendMessageEmojiInteractionSeenPredict struct {
	Emoticon string
}

func (*SendMessageEmojiInteractionSeenPredict) CRC() uint32 {
	return 0xb665902e
}
func (*SendMessageEmojiInteractionSeenPredict) _SendMessageAction() {}

type ShippingOption interface {
	tl.Object
	_ShippingOption()
}

var (
	_ ShippingOption = (*ShippingOptionPredict)(nil)
)

type ShippingOptionPredict struct {
	ID     string
	Title  string
	Prices []LabeledPrice
}

func (*ShippingOptionPredict) CRC() uint32 {
	return 0xb6213cdf
}
func (*ShippingOptionPredict) _ShippingOption() {}

type SmsJob interface {
	tl.Object
	_SmsJob()
}

var (
	_ SmsJob = (*SmsJobPredict)(nil)
)

type SmsJobPredict struct {
	JobID       string
	PhoneNumber string
	Text        string
}

func (*SmsJobPredict) CRC() uint32 {
	return 0xe6a1eeb8
}
func (*SmsJobPredict) _SmsJob() {}

type SponsoredMessage interface {
	tl.Object
	_SponsoredMessage()
}

var (
	_ SponsoredMessage = (*SponsoredMessagePredict)(nil)
)

type SponsoredMessagePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Recommended    bool     `tl:",omitempty:flags:5,implicit"`
	CanReport      bool     `tl:",omitempty:flags:12,implicit"`
	RandomID       []byte
	URL            string
	Title          string
	Message        string
	Entities       []MessageEntity `tl:",omitempty:flags:1"`
	Photo          Photo           `tl:",omitempty:flags:6"`
	Color          PeerColor       `tl:",omitempty:flags:13"`
	ButtonText     string
	SponsorInfo    *string `tl:",omitempty:flags:7"`
	AdditionalInfo *string `tl:",omitempty:flags:8"`
}

func (*SponsoredMessagePredict) CRC() uint32 {
	return 0xbdedf566
}
func (*SponsoredMessagePredict) _SponsoredMessage() {}

type SponsoredMessageReportOption interface {
	tl.Object
	_SponsoredMessageReportOption()
}

var (
	_ SponsoredMessageReportOption = (*SponsoredMessageReportOptionPredict)(nil)
)

type SponsoredMessageReportOptionPredict struct {
	Text   string
	Option []byte
}

func (*SponsoredMessageReportOptionPredict) CRC() uint32 {
	return 0x430d3150
}
func (*SponsoredMessageReportOptionPredict) _SponsoredMessageReportOption() {}

type StarsGiftOption interface {
	tl.Object
	_StarsGiftOption()
}

var (
	_ StarsGiftOption = (*StarsGiftOptionPredict)(nil)
)

type StarsGiftOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Extended     bool     `tl:",omitempty:flags:1,implicit"`
	Stars        int64
	StoreProduct *string `tl:",omitempty:flags:0"`
	Currency     string
	Amount       int64
}

func (*StarsGiftOptionPredict) CRC() uint32 {
	return 0x5e0589f1
}
func (*StarsGiftOptionPredict) _StarsGiftOption() {}

type StarsRevenueStatus interface {
	tl.Object
	_StarsRevenueStatus()
}

var (
	_ StarsRevenueStatus = (*StarsRevenueStatusPredict)(nil)
)

type StarsRevenueStatusPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	WithdrawalEnabled bool     `tl:",omitempty:flags:0,implicit"`
	CurrentBalance    int64
	AvailableBalance  int64
	OverallRevenue    int64
	NextWithdrawalAt  *int32 `tl:",omitempty:flags:1"`
}

func (*StarsRevenueStatusPredict) CRC() uint32 {
	return 0x79342946
}
func (*StarsRevenueStatusPredict) _StarsRevenueStatus() {}

type StarsTopupOption interface {
	tl.Object
	_StarsTopupOption()
}

var (
	_ StarsTopupOption = (*StarsTopupOptionPredict)(nil)
)

type StarsTopupOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Extended     bool     `tl:",omitempty:flags:1,implicit"`
	Stars        int64
	StoreProduct *string `tl:",omitempty:flags:0"`
	Currency     string
	Amount       int64
}

func (*StarsTopupOptionPredict) CRC() uint32 {
	return 0xbd915c0
}
func (*StarsTopupOptionPredict) _StarsTopupOption() {}

type StarsTransaction interface {
	tl.Object
	_StarsTransaction()
}

var (
	_ StarsTransaction = (*StarsTransactionPredict)(nil)
)

type StarsTransactionPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Refund          bool     `tl:",omitempty:flags:3,implicit"`
	Pending         bool     `tl:",omitempty:flags:4,implicit"`
	Failed          bool     `tl:",omitempty:flags:6,implicit"`
	Gift            bool     `tl:",omitempty:flags:10,implicit"`
	ID              string
	Stars           int64
	Date            int32
	Peer            StarsTransactionPeer
	Title           *string        `tl:",omitempty:flags:0"`
	Description     *string        `tl:",omitempty:flags:1"`
	Photo           WebDocument    `tl:",omitempty:flags:2"`
	TransactionDate *int32         `tl:",omitempty:flags:5"`
	TransactionURL  *string        `tl:",omitempty:flags:5"`
	BotPayload      *[]byte        `tl:",omitempty:flags:7"`
	MsgID           *int32         `tl:",omitempty:flags:8"`
	ExtendedMedia   []MessageMedia `tl:",omitempty:flags:9"`
}

func (*StarsTransactionPredict) CRC() uint32 {
	return 0x2db5418f
}
func (*StarsTransactionPredict) _StarsTransaction() {}

type StarsTransactionPeer interface {
	tl.Object
	_StarsTransactionPeer()
}

var (
	_ StarsTransactionPeer = (*StarsTransactionPeerUnsupportedPredict)(nil)
	_ StarsTransactionPeer = (*StarsTransactionPeerAppStorePredict)(nil)
	_ StarsTransactionPeer = (*StarsTransactionPeerPlayMarketPredict)(nil)
	_ StarsTransactionPeer = (*StarsTransactionPeerPremiumBotPredict)(nil)
	_ StarsTransactionPeer = (*StarsTransactionPeerFragmentPredict)(nil)
	_ StarsTransactionPeer = (*StarsTransactionPeerPredict)(nil)
	_ StarsTransactionPeer = (*StarsTransactionPeerAdsPredict)(nil)
)

type StarsTransactionPeerUnsupportedPredict struct{}

func (*StarsTransactionPeerUnsupportedPredict) CRC() uint32 {
	return 0x95f2bfe4
}
func (*StarsTransactionPeerUnsupportedPredict) _StarsTransactionPeer() {}

type StarsTransactionPeerAppStorePredict struct{}

func (*StarsTransactionPeerAppStorePredict) CRC() uint32 {
	return 0xb457b375
}
func (*StarsTransactionPeerAppStorePredict) _StarsTransactionPeer() {}

type StarsTransactionPeerPlayMarketPredict struct{}

func (*StarsTransactionPeerPlayMarketPredict) CRC() uint32 {
	return 0x7b560a0b
}
func (*StarsTransactionPeerPlayMarketPredict) _StarsTransactionPeer() {}

type StarsTransactionPeerPremiumBotPredict struct{}

func (*StarsTransactionPeerPremiumBotPredict) CRC() uint32 {
	return 0x250dbaf8
}
func (*StarsTransactionPeerPremiumBotPredict) _StarsTransactionPeer() {}

type StarsTransactionPeerFragmentPredict struct{}

func (*StarsTransactionPeerFragmentPredict) CRC() uint32 {
	return 0xe92fd902
}
func (*StarsTransactionPeerFragmentPredict) _StarsTransactionPeer() {}

type StarsTransactionPeerPredict struct {
	Peer Peer
}

func (*StarsTransactionPeerPredict) CRC() uint32 {
	return 0xd80da15d
}
func (*StarsTransactionPeerPredict) _StarsTransactionPeer() {}

type StarsTransactionPeerAdsPredict struct{}

func (*StarsTransactionPeerAdsPredict) CRC() uint32 {
	return 0x60682812
}
func (*StarsTransactionPeerAdsPredict) _StarsTransactionPeer() {}

type StatsAbsValueAndPrev interface {
	tl.Object
	_StatsAbsValueAndPrev()
}

var (
	_ StatsAbsValueAndPrev = (*StatsAbsValueAndPrevPredict)(nil)
)

type StatsAbsValueAndPrevPredict struct {
	Current  float64
	Previous float64
}

func (*StatsAbsValueAndPrevPredict) CRC() uint32 {
	return 0xcb43acde
}
func (*StatsAbsValueAndPrevPredict) _StatsAbsValueAndPrev() {}

type StatsDateRangeDays interface {
	tl.Object
	_StatsDateRangeDays()
}

var (
	_ StatsDateRangeDays = (*StatsDateRangeDaysPredict)(nil)
)

type StatsDateRangeDaysPredict struct {
	MinDate int32
	MaxDate int32
}

func (*StatsDateRangeDaysPredict) CRC() uint32 {
	return 0xb637edaf
}
func (*StatsDateRangeDaysPredict) _StatsDateRangeDays() {}

type StatsGraph interface {
	tl.Object
	_StatsGraph()
}

var (
	_ StatsGraph = (*StatsGraphAsyncPredict)(nil)
	_ StatsGraph = (*StatsGraphErrorPredict)(nil)
	_ StatsGraph = (*StatsGraphPredict)(nil)
)

type StatsGraphAsyncPredict struct {
	Token string
}

func (*StatsGraphAsyncPredict) CRC() uint32 {
	return 0x4a27eb2d
}
func (*StatsGraphAsyncPredict) _StatsGraph() {}

type StatsGraphErrorPredict struct {
	Error string
}

func (*StatsGraphErrorPredict) CRC() uint32 {
	return 0xbedc9822
}
func (*StatsGraphErrorPredict) _StatsGraph() {}

type StatsGraphPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	JSON      DataJSON
	ZoomToken *string `tl:",omitempty:flags:0"`
}

func (*StatsGraphPredict) CRC() uint32 {
	return 0x8ea464b6
}
func (*StatsGraphPredict) _StatsGraph() {}

type StatsGroupTopAdmin interface {
	tl.Object
	_StatsGroupTopAdmin()
}

var (
	_ StatsGroupTopAdmin = (*StatsGroupTopAdminPredict)(nil)
)

type StatsGroupTopAdminPredict struct {
	UserID  int64
	Deleted int32
	Kicked  int32
	Banned  int32
}

func (*StatsGroupTopAdminPredict) CRC() uint32 {
	return 0xd7584c87
}
func (*StatsGroupTopAdminPredict) _StatsGroupTopAdmin() {}

type StatsGroupTopInviter interface {
	tl.Object
	_StatsGroupTopInviter()
}

var (
	_ StatsGroupTopInviter = (*StatsGroupTopInviterPredict)(nil)
)

type StatsGroupTopInviterPredict struct {
	UserID      int64
	Invitations int32
}

func (*StatsGroupTopInviterPredict) CRC() uint32 {
	return 0x535f779d
}
func (*StatsGroupTopInviterPredict) _StatsGroupTopInviter() {}

type StatsGroupTopPoster interface {
	tl.Object
	_StatsGroupTopPoster()
}

var (
	_ StatsGroupTopPoster = (*StatsGroupTopPosterPredict)(nil)
)

type StatsGroupTopPosterPredict struct {
	UserID   int64
	Messages int32
	AvgChars int32
}

func (*StatsGroupTopPosterPredict) CRC() uint32 {
	return 0x9d04af9b
}
func (*StatsGroupTopPosterPredict) _StatsGroupTopPoster() {}

type StatsPercentValue interface {
	tl.Object
	_StatsPercentValue()
}

var (
	_ StatsPercentValue = (*StatsPercentValuePredict)(nil)
)

type StatsPercentValuePredict struct {
	Part  float64
	Total float64
}

func (*StatsPercentValuePredict) CRC() uint32 {
	return 0xcbce2fe0
}
func (*StatsPercentValuePredict) _StatsPercentValue() {}

type StatsURL interface {
	tl.Object
	_StatsURL()
}

var (
	_ StatsURL = (*StatsURLPredict)(nil)
)

type StatsURLPredict struct {
	URL string
}

func (*StatsURLPredict) CRC() uint32 {
	return 0x47a971e0
}
func (*StatsURLPredict) _StatsURL() {}

type StickerKeyword interface {
	tl.Object
	_StickerKeyword()
}

var (
	_ StickerKeyword = (*StickerKeywordPredict)(nil)
)

type StickerKeywordPredict struct {
	DocumentID int64
	Keyword    []string
}

func (*StickerKeywordPredict) CRC() uint32 {
	return 0xfcfeb29c
}
func (*StickerKeywordPredict) _StickerKeyword() {}

type StickerPack interface {
	tl.Object
	_StickerPack()
}

var (
	_ StickerPack = (*StickerPackPredict)(nil)
)

type StickerPackPredict struct {
	Emoticon  string
	Documents []int64
}

func (*StickerPackPredict) CRC() uint32 {
	return 0x12b299d4
}
func (*StickerPackPredict) _StickerPack() {}

type StickerSet interface {
	tl.Object
	_StickerSet()
}

var (
	_ StickerSet = (*StickerSetPredict)(nil)
)

type StickerSetPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Archived           bool     `tl:",omitempty:flags:1,implicit"`
	Official           bool     `tl:",omitempty:flags:2,implicit"`
	Masks              bool     `tl:",omitempty:flags:3,implicit"`
	Emojis             bool     `tl:",omitempty:flags:7,implicit"`
	TextColor          bool     `tl:",omitempty:flags:9,implicit"`
	ChannelEmojiStatus bool     `tl:",omitempty:flags:10,implicit"`
	Creator            bool     `tl:",omitempty:flags:11,implicit"`
	InstalledDate      *int32   `tl:",omitempty:flags:0"`
	ID                 int64
	AccessHash         int64
	Title              string
	ShortName          string
	Thumbs             []PhotoSize `tl:",omitempty:flags:4"`
	ThumbDcID          *int32      `tl:",omitempty:flags:4"`
	ThumbVersion       *int32      `tl:",omitempty:flags:4"`
	ThumbDocumentID    *int64      `tl:",omitempty:flags:8"`
	Count              int32
	Hash               int32
}

func (*StickerSetPredict) CRC() uint32 {
	return 0x2dd14edc
}
func (*StickerSetPredict) _StickerSet() {}

type StickerSetCovered interface {
	tl.Object
	_StickerSetCovered()
}

var (
	_ StickerSetCovered = (*StickerSetCoveredPredict)(nil)
	_ StickerSetCovered = (*StickerSetMultiCoveredPredict)(nil)
	_ StickerSetCovered = (*StickerSetFullCoveredPredict)(nil)
	_ StickerSetCovered = (*StickerSetNoCoveredPredict)(nil)
)

type StickerSetCoveredPredict struct {
	Set   StickerSet
	Cover Document
}

func (*StickerSetCoveredPredict) CRC() uint32 {
	return 0x6410a5d2
}
func (*StickerSetCoveredPredict) _StickerSetCovered() {}

type StickerSetMultiCoveredPredict struct {
	Set    StickerSet
	Covers []Document
}

func (*StickerSetMultiCoveredPredict) CRC() uint32 {
	return 0x3407e51b
}
func (*StickerSetMultiCoveredPredict) _StickerSetCovered() {}

type StickerSetFullCoveredPredict struct {
	Set       StickerSet
	Packs     []StickerPack
	Keywords  []StickerKeyword
	Documents []Document
}

func (*StickerSetFullCoveredPredict) CRC() uint32 {
	return 0x40d13c0e
}
func (*StickerSetFullCoveredPredict) _StickerSetCovered() {}

type StickerSetNoCoveredPredict struct {
	Set StickerSet
}

func (*StickerSetNoCoveredPredict) CRC() uint32 {
	return 0x77b15d1c
}
func (*StickerSetNoCoveredPredict) _StickerSetCovered() {}

type StoriesStealthMode interface {
	tl.Object
	_StoriesStealthMode()
}

var (
	_ StoriesStealthMode = (*StoriesStealthModePredict)(nil)
)

type StoriesStealthModePredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	ActiveUntilDate   *int32   `tl:",omitempty:flags:0"`
	CooldownUntilDate *int32   `tl:",omitempty:flags:1"`
}

func (*StoriesStealthModePredict) CRC() uint32 {
	return 0x712e27fd
}
func (*StoriesStealthModePredict) _StoriesStealthMode() {}

type StoryFwdHeader interface {
	tl.Object
	_StoryFwdHeader()
}

var (
	_ StoryFwdHeader = (*StoryFwdHeaderPredict)(nil)
)

type StoryFwdHeaderPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Modified bool     `tl:",omitempty:flags:3,implicit"`
	From     Peer     `tl:",omitempty:flags:0"`
	FromName *string  `tl:",omitempty:flags:1"`
	StoryID  *int32   `tl:",omitempty:flags:2"`
}

func (*StoryFwdHeaderPredict) CRC() uint32 {
	return 0xb826e150
}
func (*StoryFwdHeaderPredict) _StoryFwdHeader() {}

type StoryItem interface {
	tl.Object
	_StoryItem()
}

var (
	_ StoryItem = (*StoryItemDeletedPredict)(nil)
	_ StoryItem = (*StoryItemSkippedPredict)(nil)
	_ StoryItem = (*StoryItemPredict)(nil)
)

type StoryItemDeletedPredict struct {
	ID int32
}

func (*StoryItemDeletedPredict) CRC() uint32 {
	return 0x51e6ee4f
}
func (*StoryItemDeletedPredict) _StoryItem() {}

type StoryItemSkippedPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	CloseFriends bool     `tl:",omitempty:flags:8,implicit"`
	ID           int32
	Date         int32
	ExpireDate   int32
}

func (*StoryItemSkippedPredict) CRC() uint32 {
	return 0xffadc913
}
func (*StoryItemSkippedPredict) _StoryItem() {}

type StoryItemPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	Pinned           bool     `tl:",omitempty:flags:5,implicit"`
	Public           bool     `tl:",omitempty:flags:7,implicit"`
	CloseFriends     bool     `tl:",omitempty:flags:8,implicit"`
	Min              bool     `tl:",omitempty:flags:9,implicit"`
	Noforwards       bool     `tl:",omitempty:flags:10,implicit"`
	Edited           bool     `tl:",omitempty:flags:11,implicit"`
	Contacts         bool     `tl:",omitempty:flags:12,implicit"`
	SelectedContacts bool     `tl:",omitempty:flags:13,implicit"`
	Out              bool     `tl:",omitempty:flags:16,implicit"`
	ID               int32
	Date             int32
	FromID           Peer           `tl:",omitempty:flags:18"`
	FwdFrom          StoryFwdHeader `tl:",omitempty:flags:17"`
	ExpireDate       int32
	Caption          *string         `tl:",omitempty:flags:0"`
	Entities         []MessageEntity `tl:",omitempty:flags:1"`
	Media            MessageMedia
	MediaAreas       []MediaArea   `tl:",omitempty:flags:14"`
	Privacy          []PrivacyRule `tl:",omitempty:flags:2"`
	Views            StoryViews    `tl:",omitempty:flags:3"`
	SentReaction     Reaction      `tl:",omitempty:flags:15"`
}

func (*StoryItemPredict) CRC() uint32 {
	return 0x79b26a24
}
func (*StoryItemPredict) _StoryItem() {}

type StoryReaction interface {
	tl.Object
	_StoryReaction()
}

var (
	_ StoryReaction = (*StoryReactionPredict)(nil)
	_ StoryReaction = (*StoryReactionPublicForwardPredict)(nil)
	_ StoryReaction = (*StoryReactionPublicRepostPredict)(nil)
)

type StoryReactionPredict struct {
	PeerID   Peer
	Date     int32
	Reaction Reaction
}

func (*StoryReactionPredict) CRC() uint32 {
	return 0x6090d6d5
}
func (*StoryReactionPredict) _StoryReaction() {}

type StoryReactionPublicForwardPredict struct {
	Message Message
}

func (*StoryReactionPublicForwardPredict) CRC() uint32 {
	return 0xbbab2643
}
func (*StoryReactionPublicForwardPredict) _StoryReaction() {}

type StoryReactionPublicRepostPredict struct {
	PeerID Peer
	Story  StoryItem
}

func (*StoryReactionPublicRepostPredict) CRC() uint32 {
	return 0xcfcd0f13
}
func (*StoryReactionPublicRepostPredict) _StoryReaction() {}

type StoryView interface {
	tl.Object
	_StoryView()
}

var (
	_ StoryView = (*StoryViewPredict)(nil)
	_ StoryView = (*StoryViewPublicForwardPredict)(nil)
	_ StoryView = (*StoryViewPublicRepostPredict)(nil)
)

type StoryViewPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	UserID               int64
	Date                 int32
	Reaction             Reaction `tl:",omitempty:flags:2"`
}

func (*StoryViewPredict) CRC() uint32 {
	return 0xb0bdeac5
}
func (*StoryViewPredict) _StoryView() {}

type StoryViewPublicForwardPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	Message              Message
}

func (*StoryViewPublicForwardPredict) CRC() uint32 {
	return 0x9083670b
}
func (*StoryViewPublicForwardPredict) _StoryView() {}

type StoryViewPublicRepostPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	PeerID               Peer
	Story                StoryItem
}

func (*StoryViewPublicRepostPredict) CRC() uint32 {
	return 0xbd74cf49
}
func (*StoryViewPublicRepostPredict) _StoryView() {}

type StoryViews interface {
	tl.Object
	_StoryViews()
}

var (
	_ StoryViews = (*StoryViewsPredict)(nil)
)

type StoryViewsPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	HasViewers     bool     `tl:",omitempty:flags:1,implicit"`
	ViewsCount     int32
	ForwardsCount  *int32          `tl:",omitempty:flags:2"`
	Reactions      []ReactionCount `tl:",omitempty:flags:3"`
	ReactionsCount *int32          `tl:",omitempty:flags:4"`
	RecentViewers  []int64         `tl:",omitempty:flags:0"`
}

func (*StoryViewsPredict) CRC() uint32 {
	return 0x8d595cd6
}
func (*StoryViewsPredict) _StoryViews() {}

type TextWithEntities interface {
	tl.Object
	_TextWithEntities()
}

var (
	_ TextWithEntities = (*TextWithEntitiesPredict)(nil)
)

type TextWithEntitiesPredict struct {
	Text     string
	Entities []MessageEntity
}

func (*TextWithEntitiesPredict) CRC() uint32 {
	return 0x751f3146
}
func (*TextWithEntitiesPredict) _TextWithEntities() {}

type Theme interface {
	tl.Object
	_Theme()
}

var (
	_ Theme = (*ThemePredict)(nil)
)

type ThemePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Creator       bool     `tl:",omitempty:flags:0,implicit"`
	Default       bool     `tl:",omitempty:flags:1,implicit"`
	ForChat       bool     `tl:",omitempty:flags:5,implicit"`
	ID            int64
	AccessHash    int64
	Slug          string
	Title         string
	Document      Document        `tl:",omitempty:flags:2"`
	Settings      []ThemeSettings `tl:",omitempty:flags:3"`
	Emoticon      *string         `tl:",omitempty:flags:6"`
	InstallsCount *int32          `tl:",omitempty:flags:4"`
}

func (*ThemePredict) CRC() uint32 {
	return 0xa00e67d6
}
func (*ThemePredict) _Theme() {}

type ThemeSettings interface {
	tl.Object
	_ThemeSettings()
}

var (
	_ ThemeSettings = (*ThemeSettingsPredict)(nil)
)

type ThemeSettingsPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	MessageColorsAnimated bool     `tl:",omitempty:flags:2,implicit"`
	BaseTheme             BaseTheme
	AccentColor           int32
	OutboxAccentColor     *int32    `tl:",omitempty:flags:3"`
	MessageColors         []int32   `tl:",omitempty:flags:0"`
	Wallpaper             WallPaper `tl:",omitempty:flags:1"`
}

func (*ThemeSettingsPredict) CRC() uint32 {
	return 0xfa58b6d4
}
func (*ThemeSettingsPredict) _ThemeSettings() {}

type Timezone interface {
	tl.Object
	_Timezone()
}

var (
	_ Timezone = (*TimezonePredict)(nil)
)

type TimezonePredict struct {
	ID        string
	Name      string
	UtcOffset int32
}

func (*TimezonePredict) CRC() uint32 {
	return 0xff9289f5
}
func (*TimezonePredict) _Timezone() {}

type TopPeer interface {
	tl.Object
	_TopPeer()
}

var (
	_ TopPeer = (*TopPeerPredict)(nil)
)

type TopPeerPredict struct {
	Peer   Peer
	Rating float64
}

func (*TopPeerPredict) CRC() uint32 {
	return 0xedcdc05b
}
func (*TopPeerPredict) _TopPeer() {}

type TopPeerCategoryPeers interface {
	tl.Object
	_TopPeerCategoryPeers()
}

var (
	_ TopPeerCategoryPeers = (*TopPeerCategoryPeersPredict)(nil)
)

type TopPeerCategoryPeersPredict struct {
	Category TopPeerCategory
	Count    int32
	Peers    []TopPeer
}

func (*TopPeerCategoryPeersPredict) CRC() uint32 {
	return 0xfb834291
}
func (*TopPeerCategoryPeersPredict) _TopPeerCategoryPeers() {}

type Update interface {
	tl.Object
	_Update()
}

var (
	_ Update = (*UpdateNewMessagePredict)(nil)
	_ Update = (*UpdateMessageIDPredict)(nil)
	_ Update = (*UpdateDeleteMessagesPredict)(nil)
	_ Update = (*UpdateUserTypingPredict)(nil)
	_ Update = (*UpdateChatUserTypingPredict)(nil)
	_ Update = (*UpdateChatParticipantsPredict)(nil)
	_ Update = (*UpdateUserStatusPredict)(nil)
	_ Update = (*UpdateUserNamePredict)(nil)
	_ Update = (*UpdateNewAuthorizationPredict)(nil)
	_ Update = (*UpdateNewEncryptedMessagePredict)(nil)
	_ Update = (*UpdateEncryptedChatTypingPredict)(nil)
	_ Update = (*UpdateEncryptionPredict)(nil)
	_ Update = (*UpdateEncryptedMessagesReadPredict)(nil)
	_ Update = (*UpdateChatParticipantAddPredict)(nil)
	_ Update = (*UpdateChatParticipantDeletePredict)(nil)
	_ Update = (*UpdateDcOptionsPredict)(nil)
	_ Update = (*UpdateNotifySettingsPredict)(nil)
	_ Update = (*UpdateServiceNotificationPredict)(nil)
	_ Update = (*UpdatePrivacyPredict)(nil)
	_ Update = (*UpdateUserPhonePredict)(nil)
	_ Update = (*UpdateReadHistoryInboxPredict)(nil)
	_ Update = (*UpdateReadHistoryOutboxPredict)(nil)
	_ Update = (*UpdateWebPagePredict)(nil)
	_ Update = (*UpdateReadMessagesContentsPredict)(nil)
	_ Update = (*UpdateChannelTooLongPredict)(nil)
	_ Update = (*UpdateChannelPredict)(nil)
	_ Update = (*UpdateNewChannelMessagePredict)(nil)
	_ Update = (*UpdateReadChannelInboxPredict)(nil)
	_ Update = (*UpdateDeleteChannelMessagesPredict)(nil)
	_ Update = (*UpdateChannelMessageViewsPredict)(nil)
	_ Update = (*UpdateChatParticipantAdminPredict)(nil)
	_ Update = (*UpdateNewStickerSetPredict)(nil)
	_ Update = (*UpdateStickerSetsOrderPredict)(nil)
	_ Update = (*UpdateStickerSetsPredict)(nil)
	_ Update = (*UpdateSavedGifsPredict)(nil)
	_ Update = (*UpdateBotInlineQueryPredict)(nil)
	_ Update = (*UpdateBotInlineSendPredict)(nil)
	_ Update = (*UpdateEditChannelMessagePredict)(nil)
	_ Update = (*UpdateBotCallbackQueryPredict)(nil)
	_ Update = (*UpdateEditMessagePredict)(nil)
	_ Update = (*UpdateInlineBotCallbackQueryPredict)(nil)
	_ Update = (*UpdateReadChannelOutboxPredict)(nil)
	_ Update = (*UpdateDraftMessagePredict)(nil)
	_ Update = (*UpdateReadFeaturedStickersPredict)(nil)
	_ Update = (*UpdateRecentStickersPredict)(nil)
	_ Update = (*UpdateConfigPredict)(nil)
	_ Update = (*UpdatePtsChangedPredict)(nil)
	_ Update = (*UpdateChannelWebPagePredict)(nil)
	_ Update = (*UpdateDialogPinnedPredict)(nil)
	_ Update = (*UpdatePinnedDialogsPredict)(nil)
	_ Update = (*UpdateBotWebhookJSONPredict)(nil)
	_ Update = (*UpdateBotWebhookJSONQueryPredict)(nil)
	_ Update = (*UpdateBotShippingQueryPredict)(nil)
	_ Update = (*UpdateBotPrecheckoutQueryPredict)(nil)
	_ Update = (*UpdatePhoneCallPredict)(nil)
	_ Update = (*UpdateLangPackTooLongPredict)(nil)
	_ Update = (*UpdateLangPackPredict)(nil)
	_ Update = (*UpdateFavedStickersPredict)(nil)
	_ Update = (*UpdateChannelReadMessagesContentsPredict)(nil)
	_ Update = (*UpdateContactsResetPredict)(nil)
	_ Update = (*UpdateChannelAvailableMessagesPredict)(nil)
	_ Update = (*UpdateDialogUnreadMarkPredict)(nil)
	_ Update = (*UpdateMessagePollPredict)(nil)
	_ Update = (*UpdateChatDefaultBannedRightsPredict)(nil)
	_ Update = (*UpdateFolderPeersPredict)(nil)
	_ Update = (*UpdatePeerSettingsPredict)(nil)
	_ Update = (*UpdatePeerLocatedPredict)(nil)
	_ Update = (*UpdateNewScheduledMessagePredict)(nil)
	_ Update = (*UpdateDeleteScheduledMessagesPredict)(nil)
	_ Update = (*UpdateThemePredict)(nil)
	_ Update = (*UpdateGeoLiveViewedPredict)(nil)
	_ Update = (*UpdateLoginTokenPredict)(nil)
	_ Update = (*UpdateMessagePollVotePredict)(nil)
	_ Update = (*UpdateDialogFilterPredict)(nil)
	_ Update = (*UpdateDialogFilterOrderPredict)(nil)
	_ Update = (*UpdateDialogFiltersPredict)(nil)
	_ Update = (*UpdatePhoneCallSignalingDataPredict)(nil)
	_ Update = (*UpdateChannelMessageForwardsPredict)(nil)
	_ Update = (*UpdateReadChannelDiscussionInboxPredict)(nil)
	_ Update = (*UpdateReadChannelDiscussionOutboxPredict)(nil)
	_ Update = (*UpdatePeerBlockedPredict)(nil)
	_ Update = (*UpdateChannelUserTypingPredict)(nil)
	_ Update = (*UpdatePinnedMessagesPredict)(nil)
	_ Update = (*UpdatePinnedChannelMessagesPredict)(nil)
	_ Update = (*UpdateChatPredict)(nil)
	_ Update = (*UpdateGroupCallParticipantsPredict)(nil)
	_ Update = (*UpdateGroupCallPredict)(nil)
	_ Update = (*UpdatePeerHistoryTTLPredict)(nil)
	_ Update = (*UpdateChatParticipantPredict)(nil)
	_ Update = (*UpdateChannelParticipantPredict)(nil)
	_ Update = (*UpdateBotStoppedPredict)(nil)
	_ Update = (*UpdateGroupCallConnectionPredict)(nil)
	_ Update = (*UpdateBotCommandsPredict)(nil)
	_ Update = (*UpdatePendingJoinRequestsPredict)(nil)
	_ Update = (*UpdateBotChatInviteRequesterPredict)(nil)
	_ Update = (*UpdateMessageReactionsPredict)(nil)
	_ Update = (*UpdateAttachMenuBotsPredict)(nil)
	_ Update = (*UpdateWebViewResultSentPredict)(nil)
	_ Update = (*UpdateBotMenuButtonPredict)(nil)
	_ Update = (*UpdateSavedRingtonesPredict)(nil)
	_ Update = (*UpdateTranscribedAudioPredict)(nil)
	_ Update = (*UpdateReadFeaturedEmojiStickersPredict)(nil)
	_ Update = (*UpdateUserEmojiStatusPredict)(nil)
	_ Update = (*UpdateRecentEmojiStatusesPredict)(nil)
	_ Update = (*UpdateRecentReactionsPredict)(nil)
	_ Update = (*UpdateMoveStickerSetToTopPredict)(nil)
	_ Update = (*UpdateMessageExtendedMediaPredict)(nil)
	_ Update = (*UpdateChannelPinnedTopicPredict)(nil)
	_ Update = (*UpdateChannelPinnedTopicsPredict)(nil)
	_ Update = (*UpdateUserPredict)(nil)
	_ Update = (*UpdateAutoSaveSettingsPredict)(nil)
	_ Update = (*UpdateStoryPredict)(nil)
	_ Update = (*UpdateReadStoriesPredict)(nil)
	_ Update = (*UpdateStoryIDPredict)(nil)
	_ Update = (*UpdateStoriesStealthModePredict)(nil)
	_ Update = (*UpdateSentStoryReactionPredict)(nil)
	_ Update = (*UpdateBotChatBoostPredict)(nil)
	_ Update = (*UpdateChannelViewForumAsMessagesPredict)(nil)
	_ Update = (*UpdatePeerWallpaperPredict)(nil)
	_ Update = (*UpdateBotMessageReactionPredict)(nil)
	_ Update = (*UpdateBotMessageReactionsPredict)(nil)
	_ Update = (*UpdateSavedDialogPinnedPredict)(nil)
	_ Update = (*UpdatePinnedSavedDialogsPredict)(nil)
	_ Update = (*UpdateSavedReactionTagsPredict)(nil)
	_ Update = (*UpdateSmsJobPredict)(nil)
	_ Update = (*UpdateQuickRepliesPredict)(nil)
	_ Update = (*UpdateNewQuickReplyPredict)(nil)
	_ Update = (*UpdateDeleteQuickReplyPredict)(nil)
	_ Update = (*UpdateQuickReplyMessagePredict)(nil)
	_ Update = (*UpdateDeleteQuickReplyMessagesPredict)(nil)
	_ Update = (*UpdateBotBusinessConnectPredict)(nil)
	_ Update = (*UpdateBotNewBusinessMessagePredict)(nil)
	_ Update = (*UpdateBotEditBusinessMessagePredict)(nil)
	_ Update = (*UpdateBotDeleteBusinessMessagePredict)(nil)
	_ Update = (*UpdateNewStoryReactionPredict)(nil)
	_ Update = (*UpdateBroadcastRevenueTransactionsPredict)(nil)
	_ Update = (*UpdateStarsBalancePredict)(nil)
	_ Update = (*UpdateBusinessBotCallbackQueryPredict)(nil)
	_ Update = (*UpdateStarsRevenueStatusPredict)(nil)
)

type UpdateNewMessagePredict struct {
	Message  Message
	Pts      int32
	PtsCount int32
}

func (*UpdateNewMessagePredict) CRC() uint32 {
	return 0x1f2b0afd
}
func (*UpdateNewMessagePredict) _Update() {}

type UpdateMessageIDPredict struct {
	ID       int32
	RandomID int64
}

func (*UpdateMessageIDPredict) CRC() uint32 {
	return 0x4e90bfd6
}
func (*UpdateMessageIDPredict) _Update() {}

type UpdateDeleteMessagesPredict struct {
	Messages []int32
	Pts      int32
	PtsCount int32
}

func (*UpdateDeleteMessagesPredict) CRC() uint32 {
	return 0xa20db0e5
}
func (*UpdateDeleteMessagesPredict) _Update() {}

type UpdateUserTypingPredict struct {
	UserID int64
	Action SendMessageAction
}

func (*UpdateUserTypingPredict) CRC() uint32 {
	return 0xc01e857f
}
func (*UpdateUserTypingPredict) _Update() {}

type UpdateChatUserTypingPredict struct {
	ChatID int64
	FromID Peer
	Action SendMessageAction
}

func (*UpdateChatUserTypingPredict) CRC() uint32 {
	return 0x83487af0
}
func (*UpdateChatUserTypingPredict) _Update() {}

type UpdateChatParticipantsPredict struct {
	Participants ChatParticipants
}

func (*UpdateChatParticipantsPredict) CRC() uint32 {
	return 0x7761198
}
func (*UpdateChatParticipantsPredict) _Update() {}

type UpdateUserStatusPredict struct {
	UserID int64
	Status UserStatus
}

func (*UpdateUserStatusPredict) CRC() uint32 {
	return 0xe5bdf8de
}
func (*UpdateUserStatusPredict) _Update() {}

type UpdateUserNamePredict struct {
	UserID    int64
	FirstName string
	LastName  string
	Usernames []Username
}

func (*UpdateUserNamePredict) CRC() uint32 {
	return 0xa7848924
}
func (*UpdateUserNamePredict) _Update() {}

type UpdateNewAuthorizationPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Unconfirmed bool     `tl:",omitempty:flags:0,implicit"`
	Hash        int64
	Date        *int32  `tl:",omitempty:flags:0"`
	Device      *string `tl:",omitempty:flags:0"`
	Location    *string `tl:",omitempty:flags:0"`
}

func (*UpdateNewAuthorizationPredict) CRC() uint32 {
	return 0x8951abef
}
func (*UpdateNewAuthorizationPredict) _Update() {}

type UpdateNewEncryptedMessagePredict struct {
	Message EncryptedMessage
	Qts     int32
}

func (*UpdateNewEncryptedMessagePredict) CRC() uint32 {
	return 0x12bcbd9a
}
func (*UpdateNewEncryptedMessagePredict) _Update() {}

type UpdateEncryptedChatTypingPredict struct {
	ChatID int32
}

func (*UpdateEncryptedChatTypingPredict) CRC() uint32 {
	return 0x1710f156
}
func (*UpdateEncryptedChatTypingPredict) _Update() {}

type UpdateEncryptionPredict struct {
	Chat EncryptedChat
	Date int32
}

func (*UpdateEncryptionPredict) CRC() uint32 {
	return 0xb4a2e88d
}
func (*UpdateEncryptionPredict) _Update() {}

type UpdateEncryptedMessagesReadPredict struct {
	ChatID  int32
	MaxDate int32
	Date    int32
}

func (*UpdateEncryptedMessagesReadPredict) CRC() uint32 {
	return 0x38fe25b7
}
func (*UpdateEncryptedMessagesReadPredict) _Update() {}

type UpdateChatParticipantAddPredict struct {
	ChatID    int64
	UserID    int64
	InviterID int64
	Date      int32
	Version   int32
}

func (*UpdateChatParticipantAddPredict) CRC() uint32 {
	return 0x3dda5451
}
func (*UpdateChatParticipantAddPredict) _Update() {}

type UpdateChatParticipantDeletePredict struct {
	ChatID  int64
	UserID  int64
	Version int32
}

func (*UpdateChatParticipantDeletePredict) CRC() uint32 {
	return 0xe32f3d77
}
func (*UpdateChatParticipantDeletePredict) _Update() {}

type UpdateDcOptionsPredict struct {
	DcOptions []DcOption
}

func (*UpdateDcOptionsPredict) CRC() uint32 {
	return 0x8e5e9873
}
func (*UpdateDcOptionsPredict) _Update() {}

type UpdateNotifySettingsPredict struct {
	Peer           NotifyPeer
	NotifySettings PeerNotifySettings
}

func (*UpdateNotifySettingsPredict) CRC() uint32 {
	return 0xbec268ef
}
func (*UpdateNotifySettingsPredict) _Update() {}

type UpdateServiceNotificationPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Popup       bool     `tl:",omitempty:flags:0,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:2,implicit"`
	InboxDate   *int32   `tl:",omitempty:flags:1"`
	Type        string
	Message     string
	Media       MessageMedia
	Entities    []MessageEntity
}

func (*UpdateServiceNotificationPredict) CRC() uint32 {
	return 0xebe46819
}
func (*UpdateServiceNotificationPredict) _Update() {}

type UpdatePrivacyPredict struct {
	Key   PrivacyKey
	Rules []PrivacyRule
}

func (*UpdatePrivacyPredict) CRC() uint32 {
	return 0xee3b272a
}
func (*UpdatePrivacyPredict) _Update() {}

type UpdateUserPhonePredict struct {
	UserID int64
	Phone  string
}

func (*UpdateUserPhonePredict) CRC() uint32 {
	return 0x5492a13
}
func (*UpdateUserPhonePredict) _Update() {}

type UpdateReadHistoryInboxPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	FolderID         *int32   `tl:",omitempty:flags:0"`
	Peer             Peer
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
	PtsCount         int32
}

func (*UpdateReadHistoryInboxPredict) CRC() uint32 {
	return 0x9c974fdf
}
func (*UpdateReadHistoryInboxPredict) _Update() {}

type UpdateReadHistoryOutboxPredict struct {
	Peer     Peer
	MaxID    int32
	Pts      int32
	PtsCount int32
}

func (*UpdateReadHistoryOutboxPredict) CRC() uint32 {
	return 0x2f2f21bf
}
func (*UpdateReadHistoryOutboxPredict) _Update() {}

type UpdateWebPagePredict struct {
	Webpage  WebPage
	Pts      int32
	PtsCount int32
}

func (*UpdateWebPagePredict) CRC() uint32 {
	return 0x7f891213
}
func (*UpdateWebPagePredict) _Update() {}

type UpdateReadMessagesContentsPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Messages []int32
	Pts      int32
	PtsCount int32
	Date     *int32 `tl:",omitempty:flags:0"`
}

func (*UpdateReadMessagesContentsPredict) CRC() uint32 {
	return 0xf8227181
}
func (*UpdateReadMessagesContentsPredict) _Update() {}

type UpdateChannelTooLongPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	Pts       *int32 `tl:",omitempty:flags:0"`
}

func (*UpdateChannelTooLongPredict) CRC() uint32 {
	return 0x108d941f
}
func (*UpdateChannelTooLongPredict) _Update() {}

type UpdateChannelPredict struct {
	ChannelID int64
}

func (*UpdateChannelPredict) CRC() uint32 {
	return 0x635b4c09
}
func (*UpdateChannelPredict) _Update() {}

type UpdateNewChannelMessagePredict struct {
	Message  Message
	Pts      int32
	PtsCount int32
}

func (*UpdateNewChannelMessagePredict) CRC() uint32 {
	return 0x62ba04d9
}
func (*UpdateNewChannelMessagePredict) _Update() {}

type UpdateReadChannelInboxPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	FolderID         *int32   `tl:",omitempty:flags:0"`
	ChannelID        int64
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
}

func (*UpdateReadChannelInboxPredict) CRC() uint32 {
	return 0x922e6e10
}
func (*UpdateReadChannelInboxPredict) _Update() {}

type UpdateDeleteChannelMessagesPredict struct {
	ChannelID int64
	Messages  []int32
	Pts       int32
	PtsCount  int32
}

func (*UpdateDeleteChannelMessagesPredict) CRC() uint32 {
	return 0xc32d5b12
}
func (*UpdateDeleteChannelMessagesPredict) _Update() {}

type UpdateChannelMessageViewsPredict struct {
	ChannelID int64
	ID        int32
	Views     int32
}

func (*UpdateChannelMessageViewsPredict) CRC() uint32 {
	return 0xf226ac08
}
func (*UpdateChannelMessageViewsPredict) _Update() {}

type UpdateChatParticipantAdminPredict struct {
	ChatID  int64
	UserID  int64
	IsAdmin bool
	Version int32
}

func (*UpdateChatParticipantAdminPredict) CRC() uint32 {
	return 0xd7ca61a2
}
func (*UpdateChatParticipantAdminPredict) _Update() {}

type UpdateNewStickerSetPredict struct {
	Stickerset MessagesStickerSet
}

func (*UpdateNewStickerSetPredict) CRC() uint32 {
	return 0x688a30aa
}
func (*UpdateNewStickerSetPredict) _Update() {}

type UpdateStickerSetsOrderPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:",omitempty:flags:0,implicit"`
	Emojis bool     `tl:",omitempty:flags:1,implicit"`
	Order  []int64
}

func (*UpdateStickerSetsOrderPredict) CRC() uint32 {
	return 0xbb2d201
}
func (*UpdateStickerSetsOrderPredict) _Update() {}

type UpdateStickerSetsPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:",omitempty:flags:0,implicit"`
	Emojis bool     `tl:",omitempty:flags:1,implicit"`
}

func (*UpdateStickerSetsPredict) CRC() uint32 {
	return 0x31c24808
}
func (*UpdateStickerSetsPredict) _Update() {}

type UpdateSavedGifsPredict struct{}

func (*UpdateSavedGifsPredict) CRC() uint32 {
	return 0x9375341e
}
func (*UpdateSavedGifsPredict) _Update() {}

type UpdateBotInlineQueryPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	QueryID  int64
	UserID   int64
	Query    string
	Geo      GeoPoint            `tl:",omitempty:flags:0"`
	PeerType InlineQueryPeerType `tl:",omitempty:flags:1"`
	Offset   string
}

func (*UpdateBotInlineQueryPredict) CRC() uint32 {
	return 0x496f379c
}
func (*UpdateBotInlineQueryPredict) _Update() {}

type UpdateBotInlineSendPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	UserID int64
	Query  string
	Geo    GeoPoint `tl:",omitempty:flags:0"`
	ID     string
	MsgID  InputBotInlineMessageID `tl:",omitempty:flags:1"`
}

func (*UpdateBotInlineSendPredict) CRC() uint32 {
	return 0x12f12a07
}
func (*UpdateBotInlineSendPredict) _Update() {}

type UpdateEditChannelMessagePredict struct {
	Message  Message
	Pts      int32
	PtsCount int32
}

func (*UpdateEditChannelMessagePredict) CRC() uint32 {
	return 0x1b3f4df7
}
func (*UpdateEditChannelMessagePredict) _Update() {}

type UpdateBotCallbackQueryPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	QueryID       int64
	UserID        int64
	Peer          Peer
	MsgID         int32
	ChatInstance  int64
	Data          *[]byte `tl:",omitempty:flags:0"`
	GameShortName *string `tl:",omitempty:flags:1"`
}

func (*UpdateBotCallbackQueryPredict) CRC() uint32 {
	return 0xb9cfc48d
}
func (*UpdateBotCallbackQueryPredict) _Update() {}

type UpdateEditMessagePredict struct {
	Message  Message
	Pts      int32
	PtsCount int32
}

func (*UpdateEditMessagePredict) CRC() uint32 {
	return 0xe40370a3
}
func (*UpdateEditMessagePredict) _Update() {}

type UpdateInlineBotCallbackQueryPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	QueryID       int64
	UserID        int64
	MsgID         InputBotInlineMessageID
	ChatInstance  int64
	Data          *[]byte `tl:",omitempty:flags:0"`
	GameShortName *string `tl:",omitempty:flags:1"`
}

func (*UpdateInlineBotCallbackQueryPredict) CRC() uint32 {
	return 0x691e9052
}
func (*UpdateInlineBotCallbackQueryPredict) _Update() {}

type UpdateReadChannelOutboxPredict struct {
	ChannelID int64
	MaxID     int32
}

func (*UpdateReadChannelOutboxPredict) CRC() uint32 {
	return 0xb75f99a9
}
func (*UpdateReadChannelOutboxPredict) _Update() {}

type UpdateDraftMessagePredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     Peer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
	Draft    DraftMessage
}

func (*UpdateDraftMessagePredict) CRC() uint32 {
	return 0x1b49ec6d
}
func (*UpdateDraftMessagePredict) _Update() {}

type UpdateReadFeaturedStickersPredict struct{}

func (*UpdateReadFeaturedStickersPredict) CRC() uint32 {
	return 0x571d2742
}
func (*UpdateReadFeaturedStickersPredict) _Update() {}

type UpdateRecentStickersPredict struct{}

func (*UpdateRecentStickersPredict) CRC() uint32 {
	return 0x9a422c20
}
func (*UpdateRecentStickersPredict) _Update() {}

type UpdateConfigPredict struct{}

func (*UpdateConfigPredict) CRC() uint32 {
	return 0xa229dd06
}
func (*UpdateConfigPredict) _Update() {}

type UpdatePtsChangedPredict struct{}

func (*UpdatePtsChangedPredict) CRC() uint32 {
	return 0x3354678f
}
func (*UpdatePtsChangedPredict) _Update() {}

type UpdateChannelWebPagePredict struct {
	ChannelID int64
	Webpage   WebPage
	Pts       int32
	PtsCount  int32
}

func (*UpdateChannelWebPagePredict) CRC() uint32 {
	return 0x2f2ba99f
}
func (*UpdateChannelWebPagePredict) _Update() {}

type UpdateDialogPinnedPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Pinned   bool     `tl:",omitempty:flags:0,implicit"`
	FolderID *int32   `tl:",omitempty:flags:1"`
	Peer     DialogPeer
}

func (*UpdateDialogPinnedPredict) CRC() uint32 {
	return 0x6e6fe51c
}
func (*UpdateDialogPinnedPredict) _Update() {}

type UpdatePinnedDialogsPredict struct {
	_        struct{}     `tl:"flags,bitflag"`
	FolderID *int32       `tl:",omitempty:flags:1"`
	Order    []DialogPeer `tl:",omitempty:flags:0"`
}

func (*UpdatePinnedDialogsPredict) CRC() uint32 {
	return 0xfa0f3ca2
}
func (*UpdatePinnedDialogsPredict) _Update() {}

type UpdateBotWebhookJSONPredict struct {
	Data DataJSON
}

func (*UpdateBotWebhookJSONPredict) CRC() uint32 {
	return 0x8317c0c3
}
func (*UpdateBotWebhookJSONPredict) _Update() {}

type UpdateBotWebhookJSONQueryPredict struct {
	QueryID int64
	Data    DataJSON
	Timeout int32
}

func (*UpdateBotWebhookJSONQueryPredict) CRC() uint32 {
	return 0x9b9240a6
}
func (*UpdateBotWebhookJSONQueryPredict) _Update() {}

type UpdateBotShippingQueryPredict struct {
	QueryID         int64
	UserID          int64
	Payload         []byte
	ShippingAddress PostAddress
}

func (*UpdateBotShippingQueryPredict) CRC() uint32 {
	return 0xb5aefd7d
}
func (*UpdateBotShippingQueryPredict) _Update() {}

type UpdateBotPrecheckoutQueryPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	QueryID          int64
	UserID           int64
	Payload          []byte
	Info             PaymentRequestedInfo `tl:",omitempty:flags:0"`
	ShippingOptionID *string              `tl:",omitempty:flags:1"`
	Currency         string
	TotalAmount      int64
}

func (*UpdateBotPrecheckoutQueryPredict) CRC() uint32 {
	return 0x8caa9a96
}
func (*UpdateBotPrecheckoutQueryPredict) _Update() {}

type UpdatePhoneCallPredict struct {
	PhoneCall PhoneCall
}

func (*UpdatePhoneCallPredict) CRC() uint32 {
	return 0xab0f6b1e
}
func (*UpdatePhoneCallPredict) _Update() {}

type UpdateLangPackTooLongPredict struct {
	LangCode string
}

func (*UpdateLangPackTooLongPredict) CRC() uint32 {
	return 0x46560264
}
func (*UpdateLangPackTooLongPredict) _Update() {}

type UpdateLangPackPredict struct {
	Difference LangPackDifference
}

func (*UpdateLangPackPredict) CRC() uint32 {
	return 0x56022f4d
}
func (*UpdateLangPackPredict) _Update() {}

type UpdateFavedStickersPredict struct{}

func (*UpdateFavedStickersPredict) CRC() uint32 {
	return 0xe511996d
}
func (*UpdateFavedStickersPredict) _Update() {}

type UpdateChannelReadMessagesContentsPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	Messages  []int32
}

func (*UpdateChannelReadMessagesContentsPredict) CRC() uint32 {
	return 0xea29055d
}
func (*UpdateChannelReadMessagesContentsPredict) _Update() {}

type UpdateContactsResetPredict struct{}

func (*UpdateContactsResetPredict) CRC() uint32 {
	return 0x7084a7be
}
func (*UpdateContactsResetPredict) _Update() {}

type UpdateChannelAvailableMessagesPredict struct {
	ChannelID      int64
	AvailableMinID int32
}

func (*UpdateChannelAvailableMessagesPredict) CRC() uint32 {
	return 0xb23fc698
}
func (*UpdateChannelAvailableMessagesPredict) _Update() {}

type UpdateDialogUnreadMarkPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Unread bool     `tl:",omitempty:flags:0,implicit"`
	Peer   DialogPeer
}

func (*UpdateDialogUnreadMarkPredict) CRC() uint32 {
	return 0xe16459c3
}
func (*UpdateDialogUnreadMarkPredict) _Update() {}

type UpdateMessagePollPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	PollID  int64
	Poll    Poll `tl:",omitempty:flags:0"`
	Results PollResults
}

func (*UpdateMessagePollPredict) CRC() uint32 {
	return 0xaca1657b
}
func (*UpdateMessagePollPredict) _Update() {}

type UpdateChatDefaultBannedRightsPredict struct {
	Peer                Peer
	DefaultBannedRights ChatBannedRights
	Version             int32
}

func (*UpdateChatDefaultBannedRightsPredict) CRC() uint32 {
	return 0x54c01850
}
func (*UpdateChatDefaultBannedRightsPredict) _Update() {}

type UpdateFolderPeersPredict struct {
	FolderPeers []FolderPeer
	Pts         int32
	PtsCount    int32
}

func (*UpdateFolderPeersPredict) CRC() uint32 {
	return 0x19360dc0
}
func (*UpdateFolderPeersPredict) _Update() {}

type UpdatePeerSettingsPredict struct {
	Peer     Peer
	Settings PeerSettings
}

func (*UpdatePeerSettingsPredict) CRC() uint32 {
	return 0x6a7e7366
}
func (*UpdatePeerSettingsPredict) _Update() {}

type UpdatePeerLocatedPredict struct {
	Peers []PeerLocated
}

func (*UpdatePeerLocatedPredict) CRC() uint32 {
	return 0xb4afcfb0
}
func (*UpdatePeerLocatedPredict) _Update() {}

type UpdateNewScheduledMessagePredict struct {
	Message Message
}

func (*UpdateNewScheduledMessagePredict) CRC() uint32 {
	return 0x39a51dfb
}
func (*UpdateNewScheduledMessagePredict) _Update() {}

type UpdateDeleteScheduledMessagesPredict struct {
	Peer     Peer
	Messages []int32
}

func (*UpdateDeleteScheduledMessagesPredict) CRC() uint32 {
	return 0x90866cee
}
func (*UpdateDeleteScheduledMessagesPredict) _Update() {}

type UpdateThemePredict struct {
	Theme Theme
}

func (*UpdateThemePredict) CRC() uint32 {
	return 0x8216fba3
}
func (*UpdateThemePredict) _Update() {}

type UpdateGeoLiveViewedPredict struct {
	Peer  Peer
	MsgID int32
}

func (*UpdateGeoLiveViewedPredict) CRC() uint32 {
	return 0x871fb939
}
func (*UpdateGeoLiveViewedPredict) _Update() {}

type UpdateLoginTokenPredict struct{}

func (*UpdateLoginTokenPredict) CRC() uint32 {
	return 0x564fe691
}
func (*UpdateLoginTokenPredict) _Update() {}

type UpdateMessagePollVotePredict struct {
	PollID  int64
	Peer    Peer
	Options [][]byte
	Qts     int32
}

func (*UpdateMessagePollVotePredict) CRC() uint32 {
	return 0x24f40e77
}
func (*UpdateMessagePollVotePredict) _Update() {}

type UpdateDialogFilterPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	ID     int32
	Filter DialogFilter `tl:",omitempty:flags:0"`
}

func (*UpdateDialogFilterPredict) CRC() uint32 {
	return 0x26ffde7d
}
func (*UpdateDialogFilterPredict) _Update() {}

type UpdateDialogFilterOrderPredict struct {
	Order []int32
}

func (*UpdateDialogFilterOrderPredict) CRC() uint32 {
	return 0xa5d72105
}
func (*UpdateDialogFilterOrderPredict) _Update() {}

type UpdateDialogFiltersPredict struct{}

func (*UpdateDialogFiltersPredict) CRC() uint32 {
	return 0x3504914f
}
func (*UpdateDialogFiltersPredict) _Update() {}

type UpdatePhoneCallSignalingDataPredict struct {
	PhoneCallID int64
	Data        []byte
}

func (*UpdatePhoneCallSignalingDataPredict) CRC() uint32 {
	return 0x2661bf09
}
func (*UpdatePhoneCallSignalingDataPredict) _Update() {}

type UpdateChannelMessageForwardsPredict struct {
	ChannelID int64
	ID        int32
	Forwards  int32
}

func (*UpdateChannelMessageForwardsPredict) CRC() uint32 {
	return 0xd29a27f4
}
func (*UpdateChannelMessageForwardsPredict) _Update() {}

type UpdateReadChannelDiscussionInboxPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ChannelID     int64
	TopMsgID      int32
	ReadMaxID     int32
	BroadcastID   *int64 `tl:",omitempty:flags:0"`
	BroadcastPost *int32 `tl:",omitempty:flags:0"`
}

func (*UpdateReadChannelDiscussionInboxPredict) CRC() uint32 {
	return 0xd6b19546
}
func (*UpdateReadChannelDiscussionInboxPredict) _Update() {}

type UpdateReadChannelDiscussionOutboxPredict struct {
	ChannelID int64
	TopMsgID  int32
	ReadMaxID int32
}

func (*UpdateReadChannelDiscussionOutboxPredict) CRC() uint32 {
	return 0x695c9e7c
}
func (*UpdateReadChannelDiscussionOutboxPredict) _Update() {}

type UpdatePeerBlockedPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:",omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:",omitempty:flags:1,implicit"`
	PeerID               Peer
}

func (*UpdatePeerBlockedPredict) CRC() uint32 {
	return 0xebe07752
}
func (*UpdatePeerBlockedPredict) _Update() {}

type UpdateChannelUserTypingPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	FromID    Peer
	Action    SendMessageAction
}

func (*UpdateChannelUserTypingPredict) CRC() uint32 {
	return 0x8c88c923
}
func (*UpdateChannelUserTypingPredict) _Update() {}

type UpdatePinnedMessagesPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Pinned   bool     `tl:",omitempty:flags:0,implicit"`
	Peer     Peer
	Messages []int32
	Pts      int32
	PtsCount int32
}

func (*UpdatePinnedMessagesPredict) CRC() uint32 {
	return 0xed85eab5
}
func (*UpdatePinnedMessagesPredict) _Update() {}

type UpdatePinnedChannelMessagesPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Pinned    bool     `tl:",omitempty:flags:0,implicit"`
	ChannelID int64
	Messages  []int32
	Pts       int32
	PtsCount  int32
}

func (*UpdatePinnedChannelMessagesPredict) CRC() uint32 {
	return 0x5bb98608
}
func (*UpdatePinnedChannelMessagesPredict) _Update() {}

type UpdateChatPredict struct {
	ChatID int64
}

func (*UpdateChatPredict) CRC() uint32 {
	return 0xf89a6a4e
}
func (*UpdateChatPredict) _Update() {}

type UpdateGroupCallParticipantsPredict struct {
	Call         InputGroupCall
	Participants []GroupCallParticipant
	Version      int32
}

func (*UpdateGroupCallParticipantsPredict) CRC() uint32 {
	return 0xf2ebdb4e
}
func (*UpdateGroupCallParticipantsPredict) _Update() {}

type UpdateGroupCallPredict struct {
	ChatID int64
	Call   GroupCall
}

func (*UpdateGroupCallPredict) CRC() uint32 {
	return 0x14b24500
}
func (*UpdateGroupCallPredict) _Update() {}

type UpdatePeerHistoryTTLPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      Peer
	TTLPeriod *int32 `tl:",omitempty:flags:0"`
}

func (*UpdatePeerHistoryTTLPredict) CRC() uint32 {
	return 0xbb9bb9a5
}
func (*UpdatePeerHistoryTTLPredict) _Update() {}

type UpdateChatParticipantPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ChatID          int64
	Date            int32
	ActorID         int64
	UserID          int64
	PrevParticipant ChatParticipant    `tl:",omitempty:flags:0"`
	NewParticipant  ChatParticipant    `tl:",omitempty:flags:1"`
	Invite          ExportedChatInvite `tl:",omitempty:flags:2"`
	Qts             int32
}

func (*UpdateChatParticipantPredict) CRC() uint32 {
	return 0xd087663a
}
func (*UpdateChatParticipantPredict) _Update() {}

type UpdateChannelParticipantPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ViaChatlist     bool     `tl:",omitempty:flags:3,implicit"`
	ChannelID       int64
	Date            int32
	ActorID         int64
	UserID          int64
	PrevParticipant ChannelParticipant `tl:",omitempty:flags:0"`
	NewParticipant  ChannelParticipant `tl:",omitempty:flags:1"`
	Invite          ExportedChatInvite `tl:",omitempty:flags:2"`
	Qts             int32
}

func (*UpdateChannelParticipantPredict) CRC() uint32 {
	return 0x985d3abb
}
func (*UpdateChannelParticipantPredict) _Update() {}

type UpdateBotStoppedPredict struct {
	UserID  int64
	Date    int32
	Stopped bool
	Qts     int32
}

func (*UpdateBotStoppedPredict) CRC() uint32 {
	return 0xc4870a49
}
func (*UpdateBotStoppedPredict) _Update() {}

type UpdateGroupCallConnectionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Presentation bool     `tl:",omitempty:flags:0,implicit"`
	Params       DataJSON
}

func (*UpdateGroupCallConnectionPredict) CRC() uint32 {
	return 0xb783982
}
func (*UpdateGroupCallConnectionPredict) _Update() {}

type UpdateBotCommandsPredict struct {
	Peer     Peer
	BotID    int64
	Commands []BotCommand
}

func (*UpdateBotCommandsPredict) CRC() uint32 {
	return 0x4d712f2e
}
func (*UpdateBotCommandsPredict) _Update() {}

type UpdatePendingJoinRequestsPredict struct {
	Peer             Peer
	RequestsPending  int32
	RecentRequesters []int64
}

func (*UpdatePendingJoinRequestsPredict) CRC() uint32 {
	return 0x7063c3db
}
func (*UpdatePendingJoinRequestsPredict) _Update() {}

type UpdateBotChatInviteRequesterPredict struct {
	Peer   Peer
	Date   int32
	UserID int64
	About  string
	Invite ExportedChatInvite
	Qts    int32
}

func (*UpdateBotChatInviteRequesterPredict) CRC() uint32 {
	return 0x11dfa986
}
func (*UpdateBotChatInviteRequesterPredict) _Update() {}

type UpdateMessageReactionsPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      Peer
	MsgID     int32
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	Reactions MessageReactions
}

func (*UpdateMessageReactionsPredict) CRC() uint32 {
	return 0x5e1b3cb8
}
func (*UpdateMessageReactionsPredict) _Update() {}

type UpdateAttachMenuBotsPredict struct{}

func (*UpdateAttachMenuBotsPredict) CRC() uint32 {
	return 0x17b7a20b
}
func (*UpdateAttachMenuBotsPredict) _Update() {}

type UpdateWebViewResultSentPredict struct {
	QueryID int64
}

func (*UpdateWebViewResultSentPredict) CRC() uint32 {
	return 0x1592b79d
}
func (*UpdateWebViewResultSentPredict) _Update() {}

type UpdateBotMenuButtonPredict struct {
	BotID  int64
	Button BotMenuButton
}

func (*UpdateBotMenuButtonPredict) CRC() uint32 {
	return 0x14b85813
}
func (*UpdateBotMenuButtonPredict) _Update() {}

type UpdateSavedRingtonesPredict struct{}

func (*UpdateSavedRingtonesPredict) CRC() uint32 {
	return 0x74d8be99
}
func (*UpdateSavedRingtonesPredict) _Update() {}

type UpdateTranscribedAudioPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Pending         bool     `tl:",omitempty:flags:0,implicit"`
	Peer            Peer
	MsgID           int32
	TranscriptionID int64
	Text            string
}

func (*UpdateTranscribedAudioPredict) CRC() uint32 {
	return 0x84cd5a
}
func (*UpdateTranscribedAudioPredict) _Update() {}

type UpdateReadFeaturedEmojiStickersPredict struct{}

func (*UpdateReadFeaturedEmojiStickersPredict) CRC() uint32 {
	return 0xfb4c496c
}
func (*UpdateReadFeaturedEmojiStickersPredict) _Update() {}

type UpdateUserEmojiStatusPredict struct {
	UserID      int64
	EmojiStatus EmojiStatus
}

func (*UpdateUserEmojiStatusPredict) CRC() uint32 {
	return 0x28373599
}
func (*UpdateUserEmojiStatusPredict) _Update() {}

type UpdateRecentEmojiStatusesPredict struct{}

func (*UpdateRecentEmojiStatusesPredict) CRC() uint32 {
	return 0x30f443db
}
func (*UpdateRecentEmojiStatusesPredict) _Update() {}

type UpdateRecentReactionsPredict struct{}

func (*UpdateRecentReactionsPredict) CRC() uint32 {
	return 0x6f7863f4
}
func (*UpdateRecentReactionsPredict) _Update() {}

type UpdateMoveStickerSetToTopPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Masks      bool     `tl:",omitempty:flags:0,implicit"`
	Emojis     bool     `tl:",omitempty:flags:1,implicit"`
	Stickerset int64
}

func (*UpdateMoveStickerSetToTopPredict) CRC() uint32 {
	return 0x86fccf85
}
func (*UpdateMoveStickerSetToTopPredict) _Update() {}

type UpdateMessageExtendedMediaPredict struct {
	Peer          Peer
	MsgID         int32
	ExtendedMedia []MessageExtendedMedia
}

func (*UpdateMessageExtendedMediaPredict) CRC() uint32 {
	return 0xd5a41724
}
func (*UpdateMessageExtendedMediaPredict) _Update() {}

type UpdateChannelPinnedTopicPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Pinned    bool     `tl:",omitempty:flags:0,implicit"`
	ChannelID int64
	TopicID   int32
}

func (*UpdateChannelPinnedTopicPredict) CRC() uint32 {
	return 0x192efbe3
}
func (*UpdateChannelPinnedTopicPredict) _Update() {}

type UpdateChannelPinnedTopicsPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64
	Order     []int32 `tl:",omitempty:flags:0"`
}

func (*UpdateChannelPinnedTopicsPredict) CRC() uint32 {
	return 0xfe198602
}
func (*UpdateChannelPinnedTopicsPredict) _Update() {}

type UpdateUserPredict struct {
	UserID int64
}

func (*UpdateUserPredict) CRC() uint32 {
	return 0x20529438
}
func (*UpdateUserPredict) _Update() {}

type UpdateAutoSaveSettingsPredict struct{}

func (*UpdateAutoSaveSettingsPredict) CRC() uint32 {
	return 0xec05b097
}
func (*UpdateAutoSaveSettingsPredict) _Update() {}

type UpdateStoryPredict struct {
	Peer  Peer
	Story StoryItem
}

func (*UpdateStoryPredict) CRC() uint32 {
	return 0x75b3b798
}
func (*UpdateStoryPredict) _Update() {}

type UpdateReadStoriesPredict struct {
	Peer  Peer
	MaxID int32
}

func (*UpdateReadStoriesPredict) CRC() uint32 {
	return 0xf74e932b
}
func (*UpdateReadStoriesPredict) _Update() {}

type UpdateStoryIDPredict struct {
	ID       int32
	RandomID int64
}

func (*UpdateStoryIDPredict) CRC() uint32 {
	return 0x1bf335b9
}
func (*UpdateStoryIDPredict) _Update() {}

type UpdateStoriesStealthModePredict struct {
	StealthMode StoriesStealthMode
}

func (*UpdateStoriesStealthModePredict) CRC() uint32 {
	return 0x2c084dc1
}
func (*UpdateStoriesStealthModePredict) _Update() {}

type UpdateSentStoryReactionPredict struct {
	Peer     Peer
	StoryID  int32
	Reaction Reaction
}

func (*UpdateSentStoryReactionPredict) CRC() uint32 {
	return 0x7d627683
}
func (*UpdateSentStoryReactionPredict) _Update() {}

type UpdateBotChatBoostPredict struct {
	Peer  Peer
	Boost Boost
	Qts   int32
}

func (*UpdateBotChatBoostPredict) CRC() uint32 {
	return 0x904dd49c
}
func (*UpdateBotChatBoostPredict) _Update() {}

type UpdateChannelViewForumAsMessagesPredict struct {
	ChannelID int64
	Enabled   bool
}

func (*UpdateChannelViewForumAsMessagesPredict) CRC() uint32 {
	return 0x7b68920
}
func (*UpdateChannelViewForumAsMessagesPredict) _Update() {}

type UpdatePeerWallpaperPredict struct {
	_                   struct{} `tl:"flags,bitflag"`
	WallpaperOverridden bool     `tl:",omitempty:flags:1,implicit"`
	Peer                Peer
	Wallpaper           WallPaper `tl:",omitempty:flags:0"`
}

func (*UpdatePeerWallpaperPredict) CRC() uint32 {
	return 0xae3f101d
}
func (*UpdatePeerWallpaperPredict) _Update() {}

type UpdateBotMessageReactionPredict struct {
	Peer         Peer
	MsgID        int32
	Date         int32
	Actor        Peer
	OldReactions []Reaction
	NewReactions []Reaction
	Qts          int32
}

func (*UpdateBotMessageReactionPredict) CRC() uint32 {
	return 0xac21d3ce
}
func (*UpdateBotMessageReactionPredict) _Update() {}

type UpdateBotMessageReactionsPredict struct {
	Peer      Peer
	MsgID     int32
	Date      int32
	Reactions []ReactionCount
	Qts       int32
}

func (*UpdateBotMessageReactionsPredict) CRC() uint32 {
	return 0x9cb7759
}
func (*UpdateBotMessageReactionsPredict) _Update() {}

type UpdateSavedDialogPinnedPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Pinned bool     `tl:",omitempty:flags:0,implicit"`
	Peer   DialogPeer
}

func (*UpdateSavedDialogPinnedPredict) CRC() uint32 {
	return 0xaeaf9e74
}
func (*UpdateSavedDialogPinnedPredict) _Update() {}

type UpdatePinnedSavedDialogsPredict struct {
	_     struct{}     `tl:"flags,bitflag"`
	Order []DialogPeer `tl:",omitempty:flags:0"`
}

func (*UpdatePinnedSavedDialogsPredict) CRC() uint32 {
	return 0x686c85a6
}
func (*UpdatePinnedSavedDialogsPredict) _Update() {}

type UpdateSavedReactionTagsPredict struct{}

func (*UpdateSavedReactionTagsPredict) CRC() uint32 {
	return 0x39c67432
}
func (*UpdateSavedReactionTagsPredict) _Update() {}

type UpdateSmsJobPredict struct {
	JobID string
}

func (*UpdateSmsJobPredict) CRC() uint32 {
	return 0xf16269d4
}
func (*UpdateSmsJobPredict) _Update() {}

type UpdateQuickRepliesPredict struct {
	QuickReplies []QuickReply
}

func (*UpdateQuickRepliesPredict) CRC() uint32 {
	return 0xf9470ab2
}
func (*UpdateQuickRepliesPredict) _Update() {}

type UpdateNewQuickReplyPredict struct {
	QuickReply QuickReply
}

func (*UpdateNewQuickReplyPredict) CRC() uint32 {
	return 0xf53da717
}
func (*UpdateNewQuickReplyPredict) _Update() {}

type UpdateDeleteQuickReplyPredict struct {
	ShortcutID int32
}

func (*UpdateDeleteQuickReplyPredict) CRC() uint32 {
	return 0x53e6f1ec
}
func (*UpdateDeleteQuickReplyPredict) _Update() {}

type UpdateQuickReplyMessagePredict struct {
	Message Message
}

func (*UpdateQuickReplyMessagePredict) CRC() uint32 {
	return 0x3e050d0f
}
func (*UpdateQuickReplyMessagePredict) _Update() {}

type UpdateDeleteQuickReplyMessagesPredict struct {
	ShortcutID int32
	Messages   []int32
}

func (*UpdateDeleteQuickReplyMessagesPredict) CRC() uint32 {
	return 0x566fe7cd
}
func (*UpdateDeleteQuickReplyMessagesPredict) _Update() {}

type UpdateBotBusinessConnectPredict struct {
	Connection BotBusinessConnection
	Qts        int32
}

func (*UpdateBotBusinessConnectPredict) CRC() uint32 {
	return 0x8ae5c97a
}
func (*UpdateBotBusinessConnectPredict) _Update() {}

type UpdateBotNewBusinessMessagePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ConnectionID   string
	Message        Message
	ReplyToMessage Message `tl:",omitempty:flags:0"`
	Qts            int32
}

func (*UpdateBotNewBusinessMessagePredict) CRC() uint32 {
	return 0x9ddb347c
}
func (*UpdateBotNewBusinessMessagePredict) _Update() {}

type UpdateBotEditBusinessMessagePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ConnectionID   string
	Message        Message
	ReplyToMessage Message `tl:",omitempty:flags:0"`
	Qts            int32
}

func (*UpdateBotEditBusinessMessagePredict) CRC() uint32 {
	return 0x7df587c
}
func (*UpdateBotEditBusinessMessagePredict) _Update() {}

type UpdateBotDeleteBusinessMessagePredict struct {
	ConnectionID string
	Peer         Peer
	Messages     []int32
	Qts          int32
}

func (*UpdateBotDeleteBusinessMessagePredict) CRC() uint32 {
	return 0xa02a982e
}
func (*UpdateBotDeleteBusinessMessagePredict) _Update() {}

type UpdateNewStoryReactionPredict struct {
	StoryID  int32
	Peer     Peer
	Reaction Reaction
}

func (*UpdateNewStoryReactionPredict) CRC() uint32 {
	return 0x1824e40b
}
func (*UpdateNewStoryReactionPredict) _Update() {}

type UpdateBroadcastRevenueTransactionsPredict struct {
	Peer     Peer
	Balances BroadcastRevenueBalances
}

func (*UpdateBroadcastRevenueTransactionsPredict) CRC() uint32 {
	return 0xdfd961f5
}
func (*UpdateBroadcastRevenueTransactionsPredict) _Update() {}

type UpdateStarsBalancePredict struct {
	Balance int64
}

func (*UpdateStarsBalancePredict) CRC() uint32 {
	return 0xfb85198
}
func (*UpdateStarsBalancePredict) _Update() {}

type UpdateBusinessBotCallbackQueryPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	QueryID        int64
	UserID         int64
	ConnectionID   string
	Message        Message
	ReplyToMessage Message `tl:",omitempty:flags:2"`
	ChatInstance   int64
	Data           *[]byte `tl:",omitempty:flags:0"`
}

func (*UpdateBusinessBotCallbackQueryPredict) CRC() uint32 {
	return 0x1ea2fda7
}
func (*UpdateBusinessBotCallbackQueryPredict) _Update() {}

type UpdateStarsRevenueStatusPredict struct {
	Peer   Peer
	Status StarsRevenueStatus
}

func (*UpdateStarsRevenueStatusPredict) CRC() uint32 {
	return 0xa584b019
}
func (*UpdateStarsRevenueStatusPredict) _Update() {}

type Updates interface {
	tl.Object
	_Updates()
}

var (
	_ Updates = (*UpdatesTooLongPredict)(nil)
	_ Updates = (*UpdateShortMessagePredict)(nil)
	_ Updates = (*UpdateShortChatMessagePredict)(nil)
	_ Updates = (*UpdateShortPredict)(nil)
	_ Updates = (*UpdatesCombinedPredict)(nil)
	_ Updates = (*UpdatesPredict)(nil)
	_ Updates = (*UpdateShortSentMessagePredict)(nil)
)

type UpdatesTooLongPredict struct{}

func (*UpdatesTooLongPredict) CRC() uint32 {
	return 0xe317af7e
}
func (*UpdatesTooLongPredict) _Updates() {}

type UpdateShortMessagePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Out         bool     `tl:",omitempty:flags:1,implicit"`
	Mentioned   bool     `tl:",omitempty:flags:4,implicit"`
	MediaUnread bool     `tl:",omitempty:flags:5,implicit"`
	Silent      bool     `tl:",omitempty:flags:13,implicit"`
	ID          int32
	UserID      int64
	Message     string
	Pts         int32
	PtsCount    int32
	Date        int32
	FwdFrom     MessageFwdHeader   `tl:",omitempty:flags:2"`
	ViaBotID    *int64             `tl:",omitempty:flags:11"`
	ReplyTo     MessageReplyHeader `tl:",omitempty:flags:3"`
	Entities    []MessageEntity    `tl:",omitempty:flags:7"`
	TTLPeriod   *int32             `tl:",omitempty:flags:25"`
}

func (*UpdateShortMessagePredict) CRC() uint32 {
	return 0x313bc7f8
}
func (*UpdateShortMessagePredict) _Updates() {}

type UpdateShortChatMessagePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Out         bool     `tl:",omitempty:flags:1,implicit"`
	Mentioned   bool     `tl:",omitempty:flags:4,implicit"`
	MediaUnread bool     `tl:",omitempty:flags:5,implicit"`
	Silent      bool     `tl:",omitempty:flags:13,implicit"`
	ID          int32
	FromID      int64
	ChatID      int64
	Message     string
	Pts         int32
	PtsCount    int32
	Date        int32
	FwdFrom     MessageFwdHeader   `tl:",omitempty:flags:2"`
	ViaBotID    *int64             `tl:",omitempty:flags:11"`
	ReplyTo     MessageReplyHeader `tl:",omitempty:flags:3"`
	Entities    []MessageEntity    `tl:",omitempty:flags:7"`
	TTLPeriod   *int32             `tl:",omitempty:flags:25"`
}

func (*UpdateShortChatMessagePredict) CRC() uint32 {
	return 0x4d6deea5
}
func (*UpdateShortChatMessagePredict) _Updates() {}

type UpdateShortPredict struct {
	Update Update
	Date   int32
}

func (*UpdateShortPredict) CRC() uint32 {
	return 0x78d4dec1
}
func (*UpdateShortPredict) _Updates() {}

type UpdatesCombinedPredict struct {
	Updates  []Update
	Users    []User
	Chats    []Chat
	Date     int32
	SeqStart int32
	Seq      int32
}

func (*UpdatesCombinedPredict) CRC() uint32 {
	return 0x725b04c3
}
func (*UpdatesCombinedPredict) _Updates() {}

type UpdatesPredict struct {
	Updates []Update
	Users   []User
	Chats   []Chat
	Date    int32
	Seq     int32
}

func (*UpdatesPredict) CRC() uint32 {
	return 0x74ae4240
}
func (*UpdatesPredict) _Updates() {}

type UpdateShortSentMessagePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Out       bool     `tl:",omitempty:flags:1,implicit"`
	ID        int32
	Pts       int32
	PtsCount  int32
	Date      int32
	Media     MessageMedia    `tl:",omitempty:flags:9"`
	Entities  []MessageEntity `tl:",omitempty:flags:7"`
	TTLPeriod *int32          `tl:",omitempty:flags:25"`
}

func (*UpdateShortSentMessagePredict) CRC() uint32 {
	return 0x9015e101
}
func (*UpdateShortSentMessagePredict) _Updates() {}

type URLAuthResult interface {
	tl.Object
	_URLAuthResult()
}

var (
	_ URLAuthResult = (*URLAuthResultRequestPredict)(nil)
	_ URLAuthResult = (*URLAuthResultAcceptedPredict)(nil)
	_ URLAuthResult = (*URLAuthResultDefaultPredict)(nil)
)

type URLAuthResultRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	RequestWriteAccess bool     `tl:",omitempty:flags:0,implicit"`
	Bot                User
	Domain             string
}

func (*URLAuthResultRequestPredict) CRC() uint32 {
	return 0x92d33a0e
}
func (*URLAuthResultRequestPredict) _URLAuthResult() {}

type URLAuthResultAcceptedPredict struct {
	URL string
}

func (*URLAuthResultAcceptedPredict) CRC() uint32 {
	return 0x8f8c0e4e
}
func (*URLAuthResultAcceptedPredict) _URLAuthResult() {}

type URLAuthResultDefaultPredict struct{}

func (*URLAuthResultDefaultPredict) CRC() uint32 {
	return 0xa9d6db1f
}
func (*URLAuthResultDefaultPredict) _URLAuthResult() {}

type User interface {
	tl.Object
	_User()
}

var (
	_ User = (*UserEmptyPredict)(nil)
	_ User = (*UserPredict)(nil)
)

type UserEmptyPredict struct {
	ID int64
}

func (*UserEmptyPredict) CRC() uint32 {
	return 0xd3bc4b7a
}
func (*UserEmptyPredict) _User() {}

type UserPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	Self                  bool     `tl:",omitempty:flags:10,implicit"`
	Contact               bool     `tl:",omitempty:flags:11,implicit"`
	MutualContact         bool     `tl:",omitempty:flags:12,implicit"`
	Deleted               bool     `tl:",omitempty:flags:13,implicit"`
	Bot                   bool     `tl:",omitempty:flags:14,implicit"`
	BotChatHistory        bool     `tl:",omitempty:flags:15,implicit"`
	BotNochats            bool     `tl:",omitempty:flags:16,implicit"`
	Verified              bool     `tl:",omitempty:flags:17,implicit"`
	Restricted            bool     `tl:",omitempty:flags:18,implicit"`
	Min                   bool     `tl:",omitempty:flags:20,implicit"`
	BotInlineGeo          bool     `tl:",omitempty:flags:21,implicit"`
	Support               bool     `tl:",omitempty:flags:23,implicit"`
	Scam                  bool     `tl:",omitempty:flags:24,implicit"`
	ApplyMinPhoto         bool     `tl:",omitempty:flags:25,implicit"`
	Fake                  bool     `tl:",omitempty:flags:26,implicit"`
	BotAttachMenu         bool     `tl:",omitempty:flags:27,implicit"`
	Premium               bool     `tl:",omitempty:flags:28,implicit"`
	AttachMenuEnabled     bool     `tl:",omitempty:flags:29,implicit"`
	_                     struct{} `tl:"flags2,bitflag"`
	BotCanEdit            bool     `tl:",omitempty:flags2:1,implicit"`
	CloseFriend           bool     `tl:",omitempty:flags2:2,implicit"`
	StoriesHidden         bool     `tl:",omitempty:flags2:3,implicit"`
	StoriesUnavailable    bool     `tl:",omitempty:flags2:4,implicit"`
	ContactRequirePremium bool     `tl:",omitempty:flags2:10,implicit"`
	BotBusiness           bool     `tl:",omitempty:flags2:11,implicit"`
	BotHasMainApp         bool     `tl:",omitempty:flags2:13,implicit"`
	ID                    int64
	AccessHash            *int64              `tl:",omitempty:flags:0"`
	FirstName             *string             `tl:",omitempty:flags:1"`
	LastName              *string             `tl:",omitempty:flags:2"`
	Username              *string             `tl:",omitempty:flags:3"`
	Phone                 *string             `tl:",omitempty:flags:4"`
	Photo                 UserProfilePhoto    `tl:",omitempty:flags:5"`
	Status                UserStatus          `tl:",omitempty:flags:6"`
	BotInfoVersion        *int32              `tl:",omitempty:flags:14"`
	RestrictionReason     []RestrictionReason `tl:",omitempty:flags:18"`
	BotInlinePlaceholder  *string             `tl:",omitempty:flags:19"`
	LangCode              *string             `tl:",omitempty:flags:22"`
	EmojiStatus           EmojiStatus         `tl:",omitempty:flags:30"`
	Usernames             []Username          `tl:",omitempty:flags2:0"`
	StoriesMaxID          *int32              `tl:",omitempty:flags2:5"`
	Color                 PeerColor           `tl:",omitempty:flags2:8"`
	ProfileColor          PeerColor           `tl:",omitempty:flags2:9"`
	BotActiveUsers        *int32              `tl:",omitempty:flags2:12"`
}

func (*UserPredict) CRC() uint32 {
	return 0x83314fca
}
func (*UserPredict) _User() {}

type UserFull interface {
	tl.Object
	_UserFull()
}

var (
	_ UserFull = (*UserFullPredict)(nil)
)

type UserFullPredict struct {
	_                       struct{} `tl:"flags,bitflag"`
	Blocked                 bool     `tl:",omitempty:flags:0,implicit"`
	PhoneCallsAvailable     bool     `tl:",omitempty:flags:4,implicit"`
	PhoneCallsPrivate       bool     `tl:",omitempty:flags:5,implicit"`
	CanPinMessage           bool     `tl:",omitempty:flags:7,implicit"`
	HasScheduled            bool     `tl:",omitempty:flags:12,implicit"`
	VideoCallsAvailable     bool     `tl:",omitempty:flags:13,implicit"`
	VoiceMessagesForbidden  bool     `tl:",omitempty:flags:20,implicit"`
	TranslationsDisabled    bool     `tl:",omitempty:flags:23,implicit"`
	StoriesPinnedAvailable  bool     `tl:",omitempty:flags:26,implicit"`
	BlockedMyStoriesFrom    bool     `tl:",omitempty:flags:27,implicit"`
	WallpaperOverridden     bool     `tl:",omitempty:flags:28,implicit"`
	ContactRequirePremium   bool     `tl:",omitempty:flags:29,implicit"`
	ReadDatesPrivate        bool     `tl:",omitempty:flags:30,implicit"`
	_                       struct{} `tl:"flags2,bitflag"`
	SponsoredEnabled        bool     `tl:",omitempty:flags2:7,implicit"`
	ID                      int64
	About                   *string `tl:",omitempty:flags:1"`
	Settings                PeerSettings
	PersonalPhoto           Photo `tl:",omitempty:flags:21"`
	ProfilePhoto            Photo `tl:",omitempty:flags:2"`
	FallbackPhoto           Photo `tl:",omitempty:flags:22"`
	NotifySettings          PeerNotifySettings
	BotInfo                 BotInfo `tl:",omitempty:flags:3"`
	PinnedMsgID             *int32  `tl:",omitempty:flags:6"`
	CommonChatsCount        int32
	FolderID                *int32                  `tl:",omitempty:flags:11"`
	TTLPeriod               *int32                  `tl:",omitempty:flags:14"`
	ThemeEmoticon           *string                 `tl:",omitempty:flags:15"`
	PrivateForwardName      *string                 `tl:",omitempty:flags:16"`
	BotGroupAdminRights     ChatAdminRights         `tl:",omitempty:flags:17"`
	BotBroadcastAdminRights ChatAdminRights         `tl:",omitempty:flags:18"`
	PremiumGifts            []PremiumGiftOption     `tl:",omitempty:flags:19"`
	Wallpaper               WallPaper               `tl:",omitempty:flags:24"`
	Stories                 PeerStories             `tl:",omitempty:flags:25"`
	BusinessWorkHours       BusinessWorkHours       `tl:",omitempty:flags2:0"`
	BusinessLocation        BusinessLocation        `tl:",omitempty:flags2:1"`
	BusinessGreetingMessage BusinessGreetingMessage `tl:",omitempty:flags2:2"`
	BusinessAwayMessage     BusinessAwayMessage     `tl:",omitempty:flags2:3"`
	BusinessIntro           BusinessIntro           `tl:",omitempty:flags2:4"`
	Birthday                Birthday                `tl:",omitempty:flags2:5"`
	PersonalChannelID       *int64                  `tl:",omitempty:flags2:6"`
	PersonalChannelMessage  *int32                  `tl:",omitempty:flags2:6"`
}

func (*UserFullPredict) CRC() uint32 {
	return 0xcc997720
}
func (*UserFullPredict) _UserFull() {}

type UserProfilePhoto interface {
	tl.Object
	_UserProfilePhoto()
}

var (
	_ UserProfilePhoto = (*UserProfilePhotoEmptyPredict)(nil)
	_ UserProfilePhoto = (*UserProfilePhotoPredict)(nil)
)

type UserProfilePhotoEmptyPredict struct{}

func (*UserProfilePhotoEmptyPredict) CRC() uint32 {
	return 0x4f11bae1
}
func (*UserProfilePhotoEmptyPredict) _UserProfilePhoto() {}

type UserProfilePhotoPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	HasVideo      bool     `tl:",omitempty:flags:0,implicit"`
	Personal      bool     `tl:",omitempty:flags:2,implicit"`
	PhotoID       int64
	StrippedThumb *[]byte `tl:",omitempty:flags:1"`
	DcID          int32
}

func (*UserProfilePhotoPredict) CRC() uint32 {
	return 0x82d1f706
}
func (*UserProfilePhotoPredict) _UserProfilePhoto() {}

type UserStatus interface {
	tl.Object
	_UserStatus()
}

var (
	_ UserStatus = (*UserStatusEmptyPredict)(nil)
	_ UserStatus = (*UserStatusOnlinePredict)(nil)
	_ UserStatus = (*UserStatusOfflinePredict)(nil)
	_ UserStatus = (*UserStatusRecentlyPredict)(nil)
	_ UserStatus = (*UserStatusLastWeekPredict)(nil)
	_ UserStatus = (*UserStatusLastMonthPredict)(nil)
)

type UserStatusEmptyPredict struct{}

func (*UserStatusEmptyPredict) CRC() uint32 {
	return 0x9d05049
}
func (*UserStatusEmptyPredict) _UserStatus() {}

type UserStatusOnlinePredict struct {
	Expires int32
}

func (*UserStatusOnlinePredict) CRC() uint32 {
	return 0xedb93949
}
func (*UserStatusOnlinePredict) _UserStatus() {}

type UserStatusOfflinePredict struct {
	WasOnline int32
}

func (*UserStatusOfflinePredict) CRC() uint32 {
	return 0x8c703f
}
func (*UserStatusOfflinePredict) _UserStatus() {}

type UserStatusRecentlyPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ByMe bool     `tl:",omitempty:flags:0,implicit"`
}

func (*UserStatusRecentlyPredict) CRC() uint32 {
	return 0x7b197dc8
}
func (*UserStatusRecentlyPredict) _UserStatus() {}

type UserStatusLastWeekPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ByMe bool     `tl:",omitempty:flags:0,implicit"`
}

func (*UserStatusLastWeekPredict) CRC() uint32 {
	return 0x541a1d1a
}
func (*UserStatusLastWeekPredict) _UserStatus() {}

type UserStatusLastMonthPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ByMe bool     `tl:",omitempty:flags:0,implicit"`
}

func (*UserStatusLastMonthPredict) CRC() uint32 {
	return 0x65899777
}
func (*UserStatusLastMonthPredict) _UserStatus() {}

type Username interface {
	tl.Object
	_Username()
}

var (
	_ Username = (*UsernamePredict)(nil)
)

type UsernamePredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Editable bool     `tl:",omitempty:flags:0,implicit"`
	Active   bool     `tl:",omitempty:flags:1,implicit"`
	Username string
}

func (*UsernamePredict) CRC() uint32 {
	return 0xb4073647
}
func (*UsernamePredict) _Username() {}

type VideoSize interface {
	tl.Object
	_VideoSize()
}

var (
	_ VideoSize = (*VideoSizePredict)(nil)
	_ VideoSize = (*VideoSizeEmojiMarkupPredict)(nil)
	_ VideoSize = (*VideoSizeStickerMarkupPredict)(nil)
)

type VideoSizePredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Type         string
	W            int32
	H            int32
	Size         int32
	VideoStartTs *float64 `tl:",omitempty:flags:0"`
}

func (*VideoSizePredict) CRC() uint32 {
	return 0xde33b094
}
func (*VideoSizePredict) _VideoSize() {}

type VideoSizeEmojiMarkupPredict struct {
	EmojiID          int64
	BackgroundColors []int32
}

func (*VideoSizeEmojiMarkupPredict) CRC() uint32 {
	return 0xf85c413c
}
func (*VideoSizeEmojiMarkupPredict) _VideoSize() {}

type VideoSizeStickerMarkupPredict struct {
	Stickerset       InputStickerSet
	StickerID        int64
	BackgroundColors []int32
}

func (*VideoSizeStickerMarkupPredict) CRC() uint32 {
	return 0xda082fe
}
func (*VideoSizeStickerMarkupPredict) _VideoSize() {}

type WallPaper interface {
	tl.Object
	_WallPaper()
}

var (
	_ WallPaper = (*WallPaperPredict)(nil)
	_ WallPaper = (*WallPaperNoFilePredict)(nil)
)

type WallPaperPredict struct {
	ID         int64
	_          struct{} `tl:"flags,bitflag"`
	Creator    bool     `tl:",omitempty:flags:0,implicit"`
	Default    bool     `tl:",omitempty:flags:1,implicit"`
	Pattern    bool     `tl:",omitempty:flags:3,implicit"`
	Dark       bool     `tl:",omitempty:flags:4,implicit"`
	AccessHash int64
	Slug       string
	Document   Document
	Settings   WallPaperSettings `tl:",omitempty:flags:2"`
}

func (*WallPaperPredict) CRC() uint32 {
	return 0xa437c3ed
}
func (*WallPaperPredict) _WallPaper() {}

type WallPaperNoFilePredict struct {
	ID       int64
	_        struct{}          `tl:"flags,bitflag"`
	Default  bool              `tl:",omitempty:flags:1,implicit"`
	Dark     bool              `tl:",omitempty:flags:4,implicit"`
	Settings WallPaperSettings `tl:",omitempty:flags:2"`
}

func (*WallPaperNoFilePredict) CRC() uint32 {
	return 0xe0804116
}
func (*WallPaperNoFilePredict) _WallPaper() {}

type WallPaperSettings interface {
	tl.Object
	_WallPaperSettings()
}

var (
	_ WallPaperSettings = (*WallPaperSettingsPredict)(nil)
)

type WallPaperSettingsPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	Blur                  bool     `tl:",omitempty:flags:1,implicit"`
	Motion                bool     `tl:",omitempty:flags:2,implicit"`
	BackgroundColor       *int32   `tl:",omitempty:flags:0"`
	SecondBackgroundColor *int32   `tl:",omitempty:flags:4"`
	ThirdBackgroundColor  *int32   `tl:",omitempty:flags:5"`
	FourthBackgroundColor *int32   `tl:",omitempty:flags:6"`
	Intensity             *int32   `tl:",omitempty:flags:3"`
	Rotation              *int32   `tl:",omitempty:flags:4"`
	Emoticon              *string  `tl:",omitempty:flags:7"`
}

func (*WallPaperSettingsPredict) CRC() uint32 {
	return 0x372efcd0
}
func (*WallPaperSettingsPredict) _WallPaperSettings() {}

type WebAuthorization interface {
	tl.Object
	_WebAuthorization()
}

var (
	_ WebAuthorization = (*WebAuthorizationPredict)(nil)
)

type WebAuthorizationPredict struct {
	Hash        int64
	BotID       int64
	Domain      string
	Browser     string
	Platform    string
	DateCreated int32
	DateActive  int32
	Ip          string
	Region      string
}

func (*WebAuthorizationPredict) CRC() uint32 {
	return 0xa6f8f452
}
func (*WebAuthorizationPredict) _WebAuthorization() {}

type WebDocument interface {
	tl.Object
	_WebDocument()
}

var (
	_ WebDocument = (*WebDocumentPredict)(nil)
	_ WebDocument = (*WebDocumentNoProxyPredict)(nil)
)

type WebDocumentPredict struct {
	URL        string
	AccessHash int64
	Size       int32
	MimeType   string
	Attributes []DocumentAttribute
}

func (*WebDocumentPredict) CRC() uint32 {
	return 0x1c570ed1
}
func (*WebDocumentPredict) _WebDocument() {}

type WebDocumentNoProxyPredict struct {
	URL        string
	Size       int32
	MimeType   string
	Attributes []DocumentAttribute
}

func (*WebDocumentNoProxyPredict) CRC() uint32 {
	return 0xf9c8bcc6
}
func (*WebDocumentNoProxyPredict) _WebDocument() {}

type WebPage interface {
	tl.Object
	_WebPage()
}

var (
	_ WebPage = (*WebPageEmptyPredict)(nil)
	_ WebPage = (*WebPagePendingPredict)(nil)
	_ WebPage = (*WebPagePredict)(nil)
	_ WebPage = (*WebPageNotModifiedPredict)(nil)
)

type WebPageEmptyPredict struct {
	_   struct{} `tl:"flags,bitflag"`
	ID  int64
	URL *string `tl:",omitempty:flags:0"`
}

func (*WebPageEmptyPredict) CRC() uint32 {
	return 0x211a1788
}
func (*WebPageEmptyPredict) _WebPage() {}

type WebPagePendingPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ID   int64
	URL  *string `tl:",omitempty:flags:0"`
	Date int32
}

func (*WebPagePendingPredict) CRC() uint32 {
	return 0xb0d13e47
}
func (*WebPagePendingPredict) _WebPage() {}

type WebPagePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	HasLargeMedia bool     `tl:",omitempty:flags:13,implicit"`
	ID            int64
	URL           string
	DisplayURL    string
	Hash          int32
	Type          *string            `tl:",omitempty:flags:0"`
	SiteName      *string            `tl:",omitempty:flags:1"`
	Title         *string            `tl:",omitempty:flags:2"`
	Description   *string            `tl:",omitempty:flags:3"`
	Photo         Photo              `tl:",omitempty:flags:4"`
	EmbedURL      *string            `tl:",omitempty:flags:5"`
	EmbedType     *string            `tl:",omitempty:flags:5"`
	EmbedWidth    *int32             `tl:",omitempty:flags:6"`
	EmbedHeight   *int32             `tl:",omitempty:flags:6"`
	Duration      *int32             `tl:",omitempty:flags:7"`
	Author        *string            `tl:",omitempty:flags:8"`
	Document      Document           `tl:",omitempty:flags:9"`
	CachedPage    Page               `tl:",omitempty:flags:10"`
	Attributes    []WebPageAttribute `tl:",omitempty:flags:12"`
}

func (*WebPagePredict) CRC() uint32 {
	return 0xe89c45b2
}
func (*WebPagePredict) _WebPage() {}

type WebPageNotModifiedPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	CachedPageViews *int32   `tl:",omitempty:flags:0"`
}

func (*WebPageNotModifiedPredict) CRC() uint32 {
	return 0x7311ca11
}
func (*WebPageNotModifiedPredict) _WebPage() {}

type WebPageAttribute interface {
	tl.Object
	_WebPageAttribute()
}

var (
	_ WebPageAttribute = (*WebPageAttributeThemePredict)(nil)
	_ WebPageAttribute = (*WebPageAttributeStoryPredict)(nil)
	_ WebPageAttribute = (*WebPageAttributeStickerSetPredict)(nil)
)

type WebPageAttributeThemePredict struct {
	_         struct{}      `tl:"flags,bitflag"`
	Documents []Document    `tl:",omitempty:flags:0"`
	Settings  ThemeSettings `tl:",omitempty:flags:1"`
}

func (*WebPageAttributeThemePredict) CRC() uint32 {
	return 0x54b56617
}
func (*WebPageAttributeThemePredict) _WebPageAttribute() {}

type WebPageAttributeStoryPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Peer  Peer
	ID    int32
	Story StoryItem `tl:",omitempty:flags:0"`
}

func (*WebPageAttributeStoryPredict) CRC() uint32 {
	return 0x2e94c3e7
}
func (*WebPageAttributeStoryPredict) _WebPageAttribute() {}

type WebPageAttributeStickerSetPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Emojis    bool     `tl:",omitempty:flags:0,implicit"`
	TextColor bool     `tl:",omitempty:flags:1,implicit"`
	Stickers  []Document
}

func (*WebPageAttributeStickerSetPredict) CRC() uint32 {
	return 0x50cc03d3
}
func (*WebPageAttributeStickerSetPredict) _WebPageAttribute() {}

type WebViewMessageSent interface {
	tl.Object
	_WebViewMessageSent()
}

var (
	_ WebViewMessageSent = (*WebViewMessageSentPredict)(nil)
)

type WebViewMessageSentPredict struct {
	_     struct{}                `tl:"flags,bitflag"`
	MsgID InputBotInlineMessageID `tl:",omitempty:flags:0"`
}

func (*WebViewMessageSentPredict) CRC() uint32 {
	return 0xc94511c
}
func (*WebViewMessageSentPredict) _WebViewMessageSent() {}

type WebViewResult interface {
	tl.Object
	_WebViewResult()
}

var (
	_ WebViewResult = (*WebViewResultURLPredict)(nil)
)

type WebViewResultURLPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Fullsize bool     `tl:",omitempty:flags:1,implicit"`
	QueryID  *int64   `tl:",omitempty:flags:0"`
	URL      string
}

func (*WebViewResultURLPredict) CRC() uint32 {
	return 0x4d22ff98
}
func (*WebViewResultURLPredict) _WebViewResult() {}

type AccountAuthorizationForm interface {
	tl.Object
	_AccountAuthorizationForm()
}

var (
	_ AccountAuthorizationForm = (*AccountAuthorizationFormPredict)(nil)
)

type AccountAuthorizationFormPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	RequiredTypes    []SecureRequiredType
	Values           []SecureValue
	Errors           []SecureValueError
	Users            []User
	PrivacyPolicyURL *string `tl:",omitempty:flags:0"`
}

func (*AccountAuthorizationFormPredict) CRC() uint32 {
	return 0xad2e1cd8
}
func (*AccountAuthorizationFormPredict) _AccountAuthorizationForm() {}

type AccountAuthorizations interface {
	tl.Object
	_AccountAuthorizations()
}

var (
	_ AccountAuthorizations = (*AccountAuthorizationsPredict)(nil)
)

type AccountAuthorizationsPredict struct {
	AuthorizationTTLDays int32
	Authorizations       []Authorization
}

func (*AccountAuthorizationsPredict) CRC() uint32 {
	return 0x4bff8ea0
}
func (*AccountAuthorizationsPredict) _AccountAuthorizations() {}

type AccountAutoDownloadSettings interface {
	tl.Object
	_AccountAutoDownloadSettings()
}

var (
	_ AccountAutoDownloadSettings = (*AccountAutoDownloadSettingsPredict)(nil)
)

type AccountAutoDownloadSettingsPredict struct {
	Low    AutoDownloadSettings
	Medium AutoDownloadSettings
	High   AutoDownloadSettings
}

func (*AccountAutoDownloadSettingsPredict) CRC() uint32 {
	return 0x63cacf26
}
func (*AccountAutoDownloadSettingsPredict) _AccountAutoDownloadSettings() {}

type AccountAutoSaveSettings interface {
	tl.Object
	_AccountAutoSaveSettings()
}

var (
	_ AccountAutoSaveSettings = (*AccountAutoSaveSettingsPredict)(nil)
)

type AccountAutoSaveSettingsPredict struct {
	UsersSettings      AutoSaveSettings
	ChatsSettings      AutoSaveSettings
	BroadcastsSettings AutoSaveSettings
	Exceptions         []AutoSaveException
	Chats              []Chat
	Users              []User
}

func (*AccountAutoSaveSettingsPredict) CRC() uint32 {
	return 0x4c3e069d
}
func (*AccountAutoSaveSettingsPredict) _AccountAutoSaveSettings() {}

type AccountBusinessChatLinks interface {
	tl.Object
	_AccountBusinessChatLinks()
}

var (
	_ AccountBusinessChatLinks = (*AccountBusinessChatLinksPredict)(nil)
)

type AccountBusinessChatLinksPredict struct {
	Links []BusinessChatLink
	Chats []Chat
	Users []User
}

func (*AccountBusinessChatLinksPredict) CRC() uint32 {
	return 0xec43a2d1
}
func (*AccountBusinessChatLinksPredict) _AccountBusinessChatLinks() {}

type AccountConnectedBots interface {
	tl.Object
	_AccountConnectedBots()
}

var (
	_ AccountConnectedBots = (*AccountConnectedBotsPredict)(nil)
)

type AccountConnectedBotsPredict struct {
	ConnectedBots []ConnectedBot
	Users         []User
}

func (*AccountConnectedBotsPredict) CRC() uint32 {
	return 0x17d7f87b
}
func (*AccountConnectedBotsPredict) _AccountConnectedBots() {}

type AccountContentSettings interface {
	tl.Object
	_AccountContentSettings()
}

var (
	_ AccountContentSettings = (*AccountContentSettingsPredict)(nil)
)

type AccountContentSettingsPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	SensitiveEnabled   bool     `tl:",omitempty:flags:0,implicit"`
	SensitiveCanChange bool     `tl:",omitempty:flags:1,implicit"`
}

func (*AccountContentSettingsPredict) CRC() uint32 {
	return 0x57e28221
}
func (*AccountContentSettingsPredict) _AccountContentSettings() {}

type AccountEmailVerified interface {
	tl.Object
	_AccountEmailVerified()
}

var (
	_ AccountEmailVerified = (*AccountEmailVerifiedPredict)(nil)
	_ AccountEmailVerified = (*AccountEmailVerifiedLoginPredict)(nil)
)

type AccountEmailVerifiedPredict struct {
	Email string
}

func (*AccountEmailVerifiedPredict) CRC() uint32 {
	return 0x2b96cd1b
}
func (*AccountEmailVerifiedPredict) _AccountEmailVerified() {}

type AccountEmailVerifiedLoginPredict struct {
	Email    string
	SentCode AuthSentCode
}

func (*AccountEmailVerifiedLoginPredict) CRC() uint32 {
	return 0xe1bb0d61
}
func (*AccountEmailVerifiedLoginPredict) _AccountEmailVerified() {}

type AccountEmojiStatuses interface {
	tl.Object
	_AccountEmojiStatuses()
}

var (
	_ AccountEmojiStatuses = (*AccountEmojiStatusesNotModifiedPredict)(nil)
	_ AccountEmojiStatuses = (*AccountEmojiStatusesPredict)(nil)
)

type AccountEmojiStatusesNotModifiedPredict struct{}

func (*AccountEmojiStatusesNotModifiedPredict) CRC() uint32 {
	return 0xd08ce645
}
func (*AccountEmojiStatusesNotModifiedPredict) _AccountEmojiStatuses() {}

type AccountEmojiStatusesPredict struct {
	Hash     int64
	Statuses []EmojiStatus
}

func (*AccountEmojiStatusesPredict) CRC() uint32 {
	return 0x90c467d1
}
func (*AccountEmojiStatusesPredict) _AccountEmojiStatuses() {}

type AccountPassword interface {
	tl.Object
	_AccountPassword()
}

var (
	_ AccountPassword = (*AccountPasswordPredict)(nil)
)

type AccountPasswordPredict struct {
	_                       struct{}        `tl:"flags,bitflag"`
	HasRecovery             bool            `tl:",omitempty:flags:0,implicit"`
	HasSecureValues         bool            `tl:",omitempty:flags:1,implicit"`
	HasPassword             bool            `tl:",omitempty:flags:2,implicit"`
	CurrentAlgo             PasswordKdfAlgo `tl:",omitempty:flags:2"`
	SRPB                    *[]byte         `tl:",omitempty:flags:2"`
	SRPID                   *int64          `tl:",omitempty:flags:2"`
	Hint                    *string         `tl:",omitempty:flags:3"`
	EmailUnconfirmedPattern *string         `tl:",omitempty:flags:4"`
	NewAlgo                 PasswordKdfAlgo
	NewSecureAlgo           SecurePasswordKdfAlgo
	SecureRandom            []byte
	PendingResetDate        *int32  `tl:",omitempty:flags:5"`
	LoginEmailPattern       *string `tl:",omitempty:flags:6"`
}

func (*AccountPasswordPredict) CRC() uint32 {
	return 0x957b50fb
}
func (*AccountPasswordPredict) _AccountPassword() {}

type AccountPasswordInputSettings interface {
	tl.Object
	_AccountPasswordInputSettings()
}

var (
	_ AccountPasswordInputSettings = (*AccountPasswordInputSettingsPredict)(nil)
)

type AccountPasswordInputSettingsPredict struct {
	_                 struct{}             `tl:"flags,bitflag"`
	NewAlgo           PasswordKdfAlgo      `tl:",omitempty:flags:0"`
	NewPasswordHash   *[]byte              `tl:",omitempty:flags:0"`
	Hint              *string              `tl:",omitempty:flags:0"`
	Email             *string              `tl:",omitempty:flags:1"`
	NewSecureSettings SecureSecretSettings `tl:",omitempty:flags:2"`
}

func (*AccountPasswordInputSettingsPredict) CRC() uint32 {
	return 0xc23727c9
}
func (*AccountPasswordInputSettingsPredict) _AccountPasswordInputSettings() {}

type AccountPasswordSettings interface {
	tl.Object
	_AccountPasswordSettings()
}

var (
	_ AccountPasswordSettings = (*AccountPasswordSettingsPredict)(nil)
)

type AccountPasswordSettingsPredict struct {
	_              struct{}             `tl:"flags,bitflag"`
	Email          *string              `tl:",omitempty:flags:0"`
	SecureSettings SecureSecretSettings `tl:",omitempty:flags:1"`
}

func (*AccountPasswordSettingsPredict) CRC() uint32 {
	return 0x9a5c33e5
}
func (*AccountPasswordSettingsPredict) _AccountPasswordSettings() {}

type AccountPrivacyRules interface {
	tl.Object
	_AccountPrivacyRules()
}

var (
	_ AccountPrivacyRules = (*AccountPrivacyRulesPredict)(nil)
)

type AccountPrivacyRulesPredict struct {
	Rules []PrivacyRule
	Chats []Chat
	Users []User
}

func (*AccountPrivacyRulesPredict) CRC() uint32 {
	return 0x50a04e45
}
func (*AccountPrivacyRulesPredict) _AccountPrivacyRules() {}

type AccountResetPasswordResult interface {
	tl.Object
	_AccountResetPasswordResult()
}

var (
	_ AccountResetPasswordResult = (*AccountResetPasswordFailedWaitPredict)(nil)
	_ AccountResetPasswordResult = (*AccountResetPasswordRequestedWaitPredict)(nil)
	_ AccountResetPasswordResult = (*AccountResetPasswordOkPredict)(nil)
)

type AccountResetPasswordFailedWaitPredict struct {
	RetryDate int32
}

func (*AccountResetPasswordFailedWaitPredict) CRC() uint32 {
	return 0xe3779861
}
func (*AccountResetPasswordFailedWaitPredict) _AccountResetPasswordResult() {}

type AccountResetPasswordRequestedWaitPredict struct {
	UntilDate int32
}

func (*AccountResetPasswordRequestedWaitPredict) CRC() uint32 {
	return 0xe9effc7d
}
func (*AccountResetPasswordRequestedWaitPredict) _AccountResetPasswordResult() {}

type AccountResetPasswordOkPredict struct{}

func (*AccountResetPasswordOkPredict) CRC() uint32 {
	return 0xe926d63e
}
func (*AccountResetPasswordOkPredict) _AccountResetPasswordResult() {}

type AccountResolvedBusinessChatLinks interface {
	tl.Object
	_AccountResolvedBusinessChatLinks()
}

var (
	_ AccountResolvedBusinessChatLinks = (*AccountResolvedBusinessChatLinksPredict)(nil)
)

type AccountResolvedBusinessChatLinksPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     Peer
	Message  string
	Entities []MessageEntity `tl:",omitempty:flags:0"`
	Chats    []Chat
	Users    []User
}

func (*AccountResolvedBusinessChatLinksPredict) CRC() uint32 {
	return 0x9a23af21
}
func (*AccountResolvedBusinessChatLinksPredict) _AccountResolvedBusinessChatLinks() {}

type AccountSavedRingtone interface {
	tl.Object
	_AccountSavedRingtone()
}

var (
	_ AccountSavedRingtone = (*AccountSavedRingtonePredict)(nil)
	_ AccountSavedRingtone = (*AccountSavedRingtoneConvertedPredict)(nil)
)

type AccountSavedRingtonePredict struct{}

func (*AccountSavedRingtonePredict) CRC() uint32 {
	return 0xb7263f6d
}
func (*AccountSavedRingtonePredict) _AccountSavedRingtone() {}

type AccountSavedRingtoneConvertedPredict struct {
	Document Document
}

func (*AccountSavedRingtoneConvertedPredict) CRC() uint32 {
	return 0x1f307eb7
}
func (*AccountSavedRingtoneConvertedPredict) _AccountSavedRingtone() {}

type AccountSavedRingtones interface {
	tl.Object
	_AccountSavedRingtones()
}

var (
	_ AccountSavedRingtones = (*AccountSavedRingtonesNotModifiedPredict)(nil)
	_ AccountSavedRingtones = (*AccountSavedRingtonesPredict)(nil)
)

type AccountSavedRingtonesNotModifiedPredict struct{}

func (*AccountSavedRingtonesNotModifiedPredict) CRC() uint32 {
	return 0xfbf6e8b1
}
func (*AccountSavedRingtonesNotModifiedPredict) _AccountSavedRingtones() {}

type AccountSavedRingtonesPredict struct {
	Hash      int64
	Ringtones []Document
}

func (*AccountSavedRingtonesPredict) CRC() uint32 {
	return 0xc1e92cc5
}
func (*AccountSavedRingtonesPredict) _AccountSavedRingtones() {}

type AccountSentEmailCode interface {
	tl.Object
	_AccountSentEmailCode()
}

var (
	_ AccountSentEmailCode = (*AccountSentEmailCodePredict)(nil)
)

type AccountSentEmailCodePredict struct {
	EmailPattern string
	Length       int32
}

func (*AccountSentEmailCodePredict) CRC() uint32 {
	return 0x811f854f
}
func (*AccountSentEmailCodePredict) _AccountSentEmailCode() {}

type AccountTakeout interface {
	tl.Object
	_AccountTakeout()
}

var (
	_ AccountTakeout = (*AccountTakeoutPredict)(nil)
)

type AccountTakeoutPredict struct {
	ID int64
}

func (*AccountTakeoutPredict) CRC() uint32 {
	return 0x4dba4501
}
func (*AccountTakeoutPredict) _AccountTakeout() {}

type AccountThemes interface {
	tl.Object
	_AccountThemes()
}

var (
	_ AccountThemes = (*AccountThemesNotModifiedPredict)(nil)
	_ AccountThemes = (*AccountThemesPredict)(nil)
)

type AccountThemesNotModifiedPredict struct{}

func (*AccountThemesNotModifiedPredict) CRC() uint32 {
	return 0xf41eb622
}
func (*AccountThemesNotModifiedPredict) _AccountThemes() {}

type AccountThemesPredict struct {
	Hash   int64
	Themes []Theme
}

func (*AccountThemesPredict) CRC() uint32 {
	return 0x9a3d8c6d
}
func (*AccountThemesPredict) _AccountThemes() {}

type AccountTmpPassword interface {
	tl.Object
	_AccountTmpPassword()
}

var (
	_ AccountTmpPassword = (*AccountTmpPasswordPredict)(nil)
)

type AccountTmpPasswordPredict struct {
	TmpPassword []byte
	ValidUntil  int32
}

func (*AccountTmpPasswordPredict) CRC() uint32 {
	return 0xdb64fd34
}
func (*AccountTmpPasswordPredict) _AccountTmpPassword() {}

type AccountWallPapers interface {
	tl.Object
	_AccountWallPapers()
}

var (
	_ AccountWallPapers = (*AccountWallPapersNotModifiedPredict)(nil)
	_ AccountWallPapers = (*AccountWallPapersPredict)(nil)
)

type AccountWallPapersNotModifiedPredict struct{}

func (*AccountWallPapersNotModifiedPredict) CRC() uint32 {
	return 0x1c199183
}
func (*AccountWallPapersNotModifiedPredict) _AccountWallPapers() {}

type AccountWallPapersPredict struct {
	Hash       int64
	Wallpapers []WallPaper
}

func (*AccountWallPapersPredict) CRC() uint32 {
	return 0xcdc3858c
}
func (*AccountWallPapersPredict) _AccountWallPapers() {}

type AccountWebAuthorizations interface {
	tl.Object
	_AccountWebAuthorizations()
}

var (
	_ AccountWebAuthorizations = (*AccountWebAuthorizationsPredict)(nil)
)

type AccountWebAuthorizationsPredict struct {
	Authorizations []WebAuthorization
	Users          []User
}

func (*AccountWebAuthorizationsPredict) CRC() uint32 {
	return 0xed56c9fc
}
func (*AccountWebAuthorizationsPredict) _AccountWebAuthorizations() {}

type AuthAuthorization interface {
	tl.Object
	_AuthAuthorization()
}

var (
	_ AuthAuthorization = (*AuthAuthorizationPredict)(nil)
	_ AuthAuthorization = (*AuthAuthorizationSignUpRequiredPredict)(nil)
)

type AuthAuthorizationPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	SetupPasswordRequired bool     `tl:",omitempty:flags:1,implicit"`
	OtherwiseReloginDays  *int32   `tl:",omitempty:flags:1"`
	TmpSessions           *int32   `tl:",omitempty:flags:0"`
	FutureAuthToken       *[]byte  `tl:",omitempty:flags:2"`
	User                  User
}

func (*AuthAuthorizationPredict) CRC() uint32 {
	return 0x2ea2c0d4
}
func (*AuthAuthorizationPredict) _AuthAuthorization() {}

type AuthAuthorizationSignUpRequiredPredict struct {
	_              struct{}           `tl:"flags,bitflag"`
	TermsOfService HelpTermsOfService `tl:",omitempty:flags:0"`
}

func (*AuthAuthorizationSignUpRequiredPredict) CRC() uint32 {
	return 0x44747e9a
}
func (*AuthAuthorizationSignUpRequiredPredict) _AuthAuthorization() {}

type AuthExportedAuthorization interface {
	tl.Object
	_AuthExportedAuthorization()
}

var (
	_ AuthExportedAuthorization = (*AuthExportedAuthorizationPredict)(nil)
)

type AuthExportedAuthorizationPredict struct {
	ID    int64
	Bytes []byte
}

func (*AuthExportedAuthorizationPredict) CRC() uint32 {
	return 0xb434e2b8
}
func (*AuthExportedAuthorizationPredict) _AuthExportedAuthorization() {}

type AuthLoggedOut interface {
	tl.Object
	_AuthLoggedOut()
}

var (
	_ AuthLoggedOut = (*AuthLoggedOutPredict)(nil)
)

type AuthLoggedOutPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	FutureAuthToken *[]byte  `tl:",omitempty:flags:0"`
}

func (*AuthLoggedOutPredict) CRC() uint32 {
	return 0xc3a2835f
}
func (*AuthLoggedOutPredict) _AuthLoggedOut() {}

type AuthLoginToken interface {
	tl.Object
	_AuthLoginToken()
}

var (
	_ AuthLoginToken = (*AuthLoginTokenPredict)(nil)
	_ AuthLoginToken = (*AuthLoginTokenMigrateToPredict)(nil)
	_ AuthLoginToken = (*AuthLoginTokenSuccessPredict)(nil)
)

type AuthLoginTokenPredict struct {
	Expires int32
	Token   []byte
}

func (*AuthLoginTokenPredict) CRC() uint32 {
	return 0x629f1980
}
func (*AuthLoginTokenPredict) _AuthLoginToken() {}

type AuthLoginTokenMigrateToPredict struct {
	DcID  int32
	Token []byte
}

func (*AuthLoginTokenMigrateToPredict) CRC() uint32 {
	return 0x68e9916
}
func (*AuthLoginTokenMigrateToPredict) _AuthLoginToken() {}

type AuthLoginTokenSuccessPredict struct {
	Authorization AuthAuthorization
}

func (*AuthLoginTokenSuccessPredict) CRC() uint32 {
	return 0x390d5c5e
}
func (*AuthLoginTokenSuccessPredict) _AuthLoginToken() {}

type AuthPasswordRecovery interface {
	tl.Object
	_AuthPasswordRecovery()
}

var (
	_ AuthPasswordRecovery = (*AuthPasswordRecoveryPredict)(nil)
)

type AuthPasswordRecoveryPredict struct {
	EmailPattern string
}

func (*AuthPasswordRecoveryPredict) CRC() uint32 {
	return 0x137948a5
}
func (*AuthPasswordRecoveryPredict) _AuthPasswordRecovery() {}

type AuthSentCode interface {
	tl.Object
	_AuthSentCode()
}

var (
	_ AuthSentCode = (*AuthSentCodePredict)(nil)
	_ AuthSentCode = (*AuthSentCodeSuccessPredict)(nil)
)

type AuthSentCodePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Type          AuthSentCodeType
	PhoneCodeHash string
	NextType      AuthCodeType `tl:",omitempty:flags:1"`
	Timeout       *int32       `tl:",omitempty:flags:2"`
}

func (*AuthSentCodePredict) CRC() uint32 {
	return 0x5e002502
}
func (*AuthSentCodePredict) _AuthSentCode() {}

type AuthSentCodeSuccessPredict struct {
	Authorization AuthAuthorization
}

func (*AuthSentCodeSuccessPredict) CRC() uint32 {
	return 0x2390fe44
}
func (*AuthSentCodeSuccessPredict) _AuthSentCode() {}

type AuthSentCodeType interface {
	tl.Object
	_AuthSentCodeType()
}

var (
	_ AuthSentCodeType = (*AuthSentCodeTypeAppPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeSmsPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeCallPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeFlashCallPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeMissedCallPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeEmailCodePredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeSetUpEmailRequiredPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeFragmentSmsPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeFirebaseSmsPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeSmsWordPredict)(nil)
	_ AuthSentCodeType = (*AuthSentCodeTypeSmsPhrasePredict)(nil)
)

type AuthSentCodeTypeAppPredict struct {
	Length int32
}

func (*AuthSentCodeTypeAppPredict) CRC() uint32 {
	return 0x3dbb5986
}
func (*AuthSentCodeTypeAppPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSmsPredict struct {
	Length int32
}

func (*AuthSentCodeTypeSmsPredict) CRC() uint32 {
	return 0xc000bba2
}
func (*AuthSentCodeTypeSmsPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeCallPredict struct {
	Length int32
}

func (*AuthSentCodeTypeCallPredict) CRC() uint32 {
	return 0x5353e5a7
}
func (*AuthSentCodeTypeCallPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeFlashCallPredict struct {
	Pattern string
}

func (*AuthSentCodeTypeFlashCallPredict) CRC() uint32 {
	return 0xab03c6d9
}
func (*AuthSentCodeTypeFlashCallPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeMissedCallPredict struct {
	Prefix string
	Length int32
}

func (*AuthSentCodeTypeMissedCallPredict) CRC() uint32 {
	return 0x82006484
}
func (*AuthSentCodeTypeMissedCallPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeEmailCodePredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	AppleSigninAllowed   bool     `tl:",omitempty:flags:0,implicit"`
	GoogleSigninAllowed  bool     `tl:",omitempty:flags:1,implicit"`
	EmailPattern         string
	Length               int32
	ResetAvailablePeriod *int32 `tl:",omitempty:flags:3"`
	ResetPendingDate     *int32 `tl:",omitempty:flags:4"`
}

func (*AuthSentCodeTypeEmailCodePredict) CRC() uint32 {
	return 0xf450f59b
}
func (*AuthSentCodeTypeEmailCodePredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSetUpEmailRequiredPredict struct {
	_                   struct{} `tl:"flags,bitflag"`
	AppleSigninAllowed  bool     `tl:",omitempty:flags:0,implicit"`
	GoogleSigninAllowed bool     `tl:",omitempty:flags:1,implicit"`
}

func (*AuthSentCodeTypeSetUpEmailRequiredPredict) CRC() uint32 {
	return 0xa5491dea
}
func (*AuthSentCodeTypeSetUpEmailRequiredPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeFragmentSmsPredict struct {
	URL    string
	Length int32
}

func (*AuthSentCodeTypeFragmentSmsPredict) CRC() uint32 {
	return 0xd9565c39
}
func (*AuthSentCodeTypeFragmentSmsPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeFirebaseSmsPredict struct {
	_                      struct{} `tl:"flags,bitflag"`
	Nonce                  *[]byte  `tl:",omitempty:flags:0"`
	PlayIntegrityProjectID *int64   `tl:",omitempty:flags:2"`
	PlayIntegrityNonce     *[]byte  `tl:",omitempty:flags:2"`
	Receipt                *string  `tl:",omitempty:flags:1"`
	PushTimeout            *int32   `tl:",omitempty:flags:1"`
	Length                 int32
}

func (*AuthSentCodeTypeFirebaseSmsPredict) CRC() uint32 {
	return 0x9fd736
}
func (*AuthSentCodeTypeFirebaseSmsPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSmsWordPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Beginning *string  `tl:",omitempty:flags:0"`
}

func (*AuthSentCodeTypeSmsWordPredict) CRC() uint32 {
	return 0xa416ac81
}
func (*AuthSentCodeTypeSmsWordPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSmsPhrasePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Beginning *string  `tl:",omitempty:flags:0"`
}

func (*AuthSentCodeTypeSmsPhrasePredict) CRC() uint32 {
	return 0xb37794af
}
func (*AuthSentCodeTypeSmsPhrasePredict) _AuthSentCodeType() {}

type BotsBotInfo interface {
	tl.Object
	_BotsBotInfo()
}

var (
	_ BotsBotInfo = (*BotsBotInfoPredict)(nil)
)

type BotsBotInfoPredict struct {
	Name        string
	About       string
	Description string
}

func (*BotsBotInfoPredict) CRC() uint32 {
	return 0xe8a775b0
}
func (*BotsBotInfoPredict) _BotsBotInfo() {}

type BotsPopularAppBots interface {
	tl.Object
	_BotsPopularAppBots()
}

var (
	_ BotsPopularAppBots = (*BotsPopularAppBotsPredict)(nil)
)

type BotsPopularAppBotsPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	NextOffset *string  `tl:",omitempty:flags:0"`
	Users      []User
}

func (*BotsPopularAppBotsPredict) CRC() uint32 {
	return 0x1991b13b
}
func (*BotsPopularAppBotsPredict) _BotsPopularAppBots() {}

type BotsPreviewInfo interface {
	tl.Object
	_BotsPreviewInfo()
}

var (
	_ BotsPreviewInfo = (*BotsPreviewInfoPredict)(nil)
)

type BotsPreviewInfoPredict struct {
	Media     []BotPreviewMedia
	LangCodes []string
}

func (*BotsPreviewInfoPredict) CRC() uint32 {
	return 0xca71d64
}
func (*BotsPreviewInfoPredict) _BotsPreviewInfo() {}

type ChannelsAdminLogResults interface {
	tl.Object
	_ChannelsAdminLogResults()
}

var (
	_ ChannelsAdminLogResults = (*ChannelsAdminLogResultsPredict)(nil)
)

type ChannelsAdminLogResultsPredict struct {
	Events []ChannelAdminLogEvent
	Chats  []Chat
	Users  []User
}

func (*ChannelsAdminLogResultsPredict) CRC() uint32 {
	return 0xed8af74d
}
func (*ChannelsAdminLogResultsPredict) _ChannelsAdminLogResults() {}

type ChannelsChannelParticipant interface {
	tl.Object
	_ChannelsChannelParticipant()
}

var (
	_ ChannelsChannelParticipant = (*ChannelsChannelParticipantPredict)(nil)
)

type ChannelsChannelParticipantPredict struct {
	Participant ChannelParticipant
	Chats       []Chat
	Users       []User
}

func (*ChannelsChannelParticipantPredict) CRC() uint32 {
	return 0xdfb80317
}
func (*ChannelsChannelParticipantPredict) _ChannelsChannelParticipant() {}

type ChannelsChannelParticipants interface {
	tl.Object
	_ChannelsChannelParticipants()
}

var (
	_ ChannelsChannelParticipants = (*ChannelsChannelParticipantsPredict)(nil)
	_ ChannelsChannelParticipants = (*ChannelsChannelParticipantsNotModifiedPredict)(nil)
)

type ChannelsChannelParticipantsPredict struct {
	Count        int32
	Participants []ChannelParticipant
	Chats        []Chat
	Users        []User
}

func (*ChannelsChannelParticipantsPredict) CRC() uint32 {
	return 0x9ab0feaf
}
func (*ChannelsChannelParticipantsPredict) _ChannelsChannelParticipants() {}

type ChannelsChannelParticipantsNotModifiedPredict struct{}

func (*ChannelsChannelParticipantsNotModifiedPredict) CRC() uint32 {
	return 0xf0173fe9
}
func (*ChannelsChannelParticipantsNotModifiedPredict) _ChannelsChannelParticipants() {}

type ChannelsSendAsPeers interface {
	tl.Object
	_ChannelsSendAsPeers()
}

var (
	_ ChannelsSendAsPeers = (*ChannelsSendAsPeersPredict)(nil)
)

type ChannelsSendAsPeersPredict struct {
	Peers []SendAsPeer
	Chats []Chat
	Users []User
}

func (*ChannelsSendAsPeersPredict) CRC() uint32 {
	return 0xf496b0c6
}
func (*ChannelsSendAsPeersPredict) _ChannelsSendAsPeers() {}

type ChannelsSponsoredMessageReportResult interface {
	tl.Object
	_ChannelsSponsoredMessageReportResult()
}

var (
	_ ChannelsSponsoredMessageReportResult = (*ChannelsSponsoredMessageReportResultChooseOptionPredict)(nil)
	_ ChannelsSponsoredMessageReportResult = (*ChannelsSponsoredMessageReportResultAdsHiddenPredict)(nil)
	_ ChannelsSponsoredMessageReportResult = (*ChannelsSponsoredMessageReportResultReportedPredict)(nil)
)

type ChannelsSponsoredMessageReportResultChooseOptionPredict struct {
	Title   string
	Options []SponsoredMessageReportOption
}

func (*ChannelsSponsoredMessageReportResultChooseOptionPredict) CRC() uint32 {
	return 0x846f9e42
}
func (*ChannelsSponsoredMessageReportResultChooseOptionPredict) _ChannelsSponsoredMessageReportResult() {
}

type ChannelsSponsoredMessageReportResultAdsHiddenPredict struct{}

func (*ChannelsSponsoredMessageReportResultAdsHiddenPredict) CRC() uint32 {
	return 0x3e3bcf2f
}
func (*ChannelsSponsoredMessageReportResultAdsHiddenPredict) _ChannelsSponsoredMessageReportResult() {
}

type ChannelsSponsoredMessageReportResultReportedPredict struct{}

func (*ChannelsSponsoredMessageReportResultReportedPredict) CRC() uint32 {
	return 0xad798849
}
func (*ChannelsSponsoredMessageReportResultReportedPredict) _ChannelsSponsoredMessageReportResult() {}

type ChatlistsChatlistInvite interface {
	tl.Object
	_ChatlistsChatlistInvite()
}

var (
	_ ChatlistsChatlistInvite = (*ChatlistsChatlistInviteAlreadyPredict)(nil)
	_ ChatlistsChatlistInvite = (*ChatlistsChatlistInvitePredict)(nil)
)

type ChatlistsChatlistInviteAlreadyPredict struct {
	FilterID     int32
	MissingPeers []Peer
	AlreadyPeers []Peer
	Chats        []Chat
	Users        []User
}

func (*ChatlistsChatlistInviteAlreadyPredict) CRC() uint32 {
	return 0xfa87f659
}
func (*ChatlistsChatlistInviteAlreadyPredict) _ChatlistsChatlistInvite() {}

type ChatlistsChatlistInvitePredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Title    string
	Emoticon *string `tl:",omitempty:flags:0"`
	Peers    []Peer
	Chats    []Chat
	Users    []User
}

func (*ChatlistsChatlistInvitePredict) CRC() uint32 {
	return 0x1dcd839d
}
func (*ChatlistsChatlistInvitePredict) _ChatlistsChatlistInvite() {}

type ChatlistsChatlistUpdates interface {
	tl.Object
	_ChatlistsChatlistUpdates()
}

var (
	_ ChatlistsChatlistUpdates = (*ChatlistsChatlistUpdatesPredict)(nil)
)

type ChatlistsChatlistUpdatesPredict struct {
	MissingPeers []Peer
	Chats        []Chat
	Users        []User
}

func (*ChatlistsChatlistUpdatesPredict) CRC() uint32 {
	return 0x93bd878d
}
func (*ChatlistsChatlistUpdatesPredict) _ChatlistsChatlistUpdates() {}

type ChatlistsExportedChatlistInvite interface {
	tl.Object
	_ChatlistsExportedChatlistInvite()
}

var (
	_ ChatlistsExportedChatlistInvite = (*ChatlistsExportedChatlistInvitePredict)(nil)
)

type ChatlistsExportedChatlistInvitePredict struct {
	Filter DialogFilter
	Invite ExportedChatlistInvite
}

func (*ChatlistsExportedChatlistInvitePredict) CRC() uint32 {
	return 0x10e6e3a6
}
func (*ChatlistsExportedChatlistInvitePredict) _ChatlistsExportedChatlistInvite() {}

type ChatlistsExportedInvites interface {
	tl.Object
	_ChatlistsExportedInvites()
}

var (
	_ ChatlistsExportedInvites = (*ChatlistsExportedInvitesPredict)(nil)
)

type ChatlistsExportedInvitesPredict struct {
	Invites []ExportedChatlistInvite
	Chats   []Chat
	Users   []User
}

func (*ChatlistsExportedInvitesPredict) CRC() uint32 {
	return 0x10ab6dc7
}
func (*ChatlistsExportedInvitesPredict) _ChatlistsExportedInvites() {}

type ContactsBlocked interface {
	tl.Object
	_ContactsBlocked()
}

var (
	_ ContactsBlocked = (*ContactsBlockedPredict)(nil)
	_ ContactsBlocked = (*ContactsBlockedSlicePredict)(nil)
)

type ContactsBlockedPredict struct {
	Blocked []PeerBlocked
	Chats   []Chat
	Users   []User
}

func (*ContactsBlockedPredict) CRC() uint32 {
	return 0xade1591
}
func (*ContactsBlockedPredict) _ContactsBlocked() {}

type ContactsBlockedSlicePredict struct {
	Count   int32
	Blocked []PeerBlocked
	Chats   []Chat
	Users   []User
}

func (*ContactsBlockedSlicePredict) CRC() uint32 {
	return 0xe1664194
}
func (*ContactsBlockedSlicePredict) _ContactsBlocked() {}

type ContactsContactBirthdays interface {
	tl.Object
	_ContactsContactBirthdays()
}

var (
	_ ContactsContactBirthdays = (*ContactsContactBirthdaysPredict)(nil)
)

type ContactsContactBirthdaysPredict struct {
	Contacts []ContactBirthday
	Users    []User
}

func (*ContactsContactBirthdaysPredict) CRC() uint32 {
	return 0x114ff30d
}
func (*ContactsContactBirthdaysPredict) _ContactsContactBirthdays() {}

type ContactsContacts interface {
	tl.Object
	_ContactsContacts()
}

var (
	_ ContactsContacts = (*ContactsContactsNotModifiedPredict)(nil)
	_ ContactsContacts = (*ContactsContactsPredict)(nil)
)

type ContactsContactsNotModifiedPredict struct{}

func (*ContactsContactsNotModifiedPredict) CRC() uint32 {
	return 0xb74ba9d2
}
func (*ContactsContactsNotModifiedPredict) _ContactsContacts() {}

type ContactsContactsPredict struct {
	Contacts   []Contact
	SavedCount int32
	Users      []User
}

func (*ContactsContactsPredict) CRC() uint32 {
	return 0xeae87e42
}
func (*ContactsContactsPredict) _ContactsContacts() {}

type ContactsFound interface {
	tl.Object
	_ContactsFound()
}

var (
	_ ContactsFound = (*ContactsFoundPredict)(nil)
)

type ContactsFoundPredict struct {
	MyResults []Peer
	Results   []Peer
	Chats     []Chat
	Users     []User
}

func (*ContactsFoundPredict) CRC() uint32 {
	return 0xb3134d9d
}
func (*ContactsFoundPredict) _ContactsFound() {}

type ContactsImportedContacts interface {
	tl.Object
	_ContactsImportedContacts()
}

var (
	_ ContactsImportedContacts = (*ContactsImportedContactsPredict)(nil)
)

type ContactsImportedContactsPredict struct {
	Imported       []ImportedContact
	PopularInvites []PopularContact
	RetryContacts  []int64
	Users          []User
}

func (*ContactsImportedContactsPredict) CRC() uint32 {
	return 0x77d01c3b
}
func (*ContactsImportedContactsPredict) _ContactsImportedContacts() {}

type ContactsResolvedPeer interface {
	tl.Object
	_ContactsResolvedPeer()
}

var (
	_ ContactsResolvedPeer = (*ContactsResolvedPeerPredict)(nil)
)

type ContactsResolvedPeerPredict struct {
	Peer  Peer
	Chats []Chat
	Users []User
}

func (*ContactsResolvedPeerPredict) CRC() uint32 {
	return 0x7f077ad9
}
func (*ContactsResolvedPeerPredict) _ContactsResolvedPeer() {}

type ContactsTopPeers interface {
	tl.Object
	_ContactsTopPeers()
}

var (
	_ ContactsTopPeers = (*ContactsTopPeersNotModifiedPredict)(nil)
	_ ContactsTopPeers = (*ContactsTopPeersPredict)(nil)
	_ ContactsTopPeers = (*ContactsTopPeersDisabledPredict)(nil)
)

type ContactsTopPeersNotModifiedPredict struct{}

func (*ContactsTopPeersNotModifiedPredict) CRC() uint32 {
	return 0xde266ef5
}
func (*ContactsTopPeersNotModifiedPredict) _ContactsTopPeers() {}

type ContactsTopPeersPredict struct {
	Categories []TopPeerCategoryPeers
	Chats      []Chat
	Users      []User
}

func (*ContactsTopPeersPredict) CRC() uint32 {
	return 0x70b772a8
}
func (*ContactsTopPeersPredict) _ContactsTopPeers() {}

type ContactsTopPeersDisabledPredict struct{}

func (*ContactsTopPeersDisabledPredict) CRC() uint32 {
	return 0xb52c939d
}
func (*ContactsTopPeersDisabledPredict) _ContactsTopPeers() {}

type FragmentCollectibleInfo interface {
	tl.Object
	_FragmentCollectibleInfo()
}

var (
	_ FragmentCollectibleInfo = (*FragmentCollectibleInfoPredict)(nil)
)

type FragmentCollectibleInfoPredict struct {
	PurchaseDate   int32
	Currency       string
	Amount         int64
	CryptoCurrency string
	CryptoAmount   int64
	URL            string
}

func (*FragmentCollectibleInfoPredict) CRC() uint32 {
	return 0x6ebdff91
}
func (*FragmentCollectibleInfoPredict) _FragmentCollectibleInfo() {}

type HelpAppConfig interface {
	tl.Object
	_HelpAppConfig()
}

var (
	_ HelpAppConfig = (*HelpAppConfigNotModifiedPredict)(nil)
	_ HelpAppConfig = (*HelpAppConfigPredict)(nil)
)

type HelpAppConfigNotModifiedPredict struct{}

func (*HelpAppConfigNotModifiedPredict) CRC() uint32 {
	return 0x7cde641d
}
func (*HelpAppConfigNotModifiedPredict) _HelpAppConfig() {}

type HelpAppConfigPredict struct {
	Hash   int32
	Config JSONValue
}

func (*HelpAppConfigPredict) CRC() uint32 {
	return 0xdd18782e
}
func (*HelpAppConfigPredict) _HelpAppConfig() {}

type HelpAppUpdate interface {
	tl.Object
	_HelpAppUpdate()
}

var (
	_ HelpAppUpdate = (*HelpAppUpdatePredict)(nil)
	_ HelpAppUpdate = (*HelpNoAppUpdatePredict)(nil)
)

type HelpAppUpdatePredict struct {
	_          struct{} `tl:"flags,bitflag"`
	CanNotSkip bool     `tl:",omitempty:flags:0,implicit"`
	ID         int32
	Version    string
	Text       string
	Entities   []MessageEntity
	Document   Document `tl:",omitempty:flags:1"`
	URL        *string  `tl:",omitempty:flags:2"`
	Sticker    Document `tl:",omitempty:flags:3"`
}

func (*HelpAppUpdatePredict) CRC() uint32 {
	return 0xccbbce30
}
func (*HelpAppUpdatePredict) _HelpAppUpdate() {}

type HelpNoAppUpdatePredict struct{}

func (*HelpNoAppUpdatePredict) CRC() uint32 {
	return 0xc45a6536
}
func (*HelpNoAppUpdatePredict) _HelpAppUpdate() {}

type HelpCountriesList interface {
	tl.Object
	_HelpCountriesList()
}

var (
	_ HelpCountriesList = (*HelpCountriesListNotModifiedPredict)(nil)
	_ HelpCountriesList = (*HelpCountriesListPredict)(nil)
)

type HelpCountriesListNotModifiedPredict struct{}

func (*HelpCountriesListNotModifiedPredict) CRC() uint32 {
	return 0x93cc1f32
}
func (*HelpCountriesListNotModifiedPredict) _HelpCountriesList() {}

type HelpCountriesListPredict struct {
	Countries []HelpCountry
	Hash      int32
}

func (*HelpCountriesListPredict) CRC() uint32 {
	return 0x87d0759e
}
func (*HelpCountriesListPredict) _HelpCountriesList() {}

type HelpCountry interface {
	tl.Object
	_HelpCountry()
}

var (
	_ HelpCountry = (*HelpCountryPredict)(nil)
)

type HelpCountryPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Hidden       bool     `tl:",omitempty:flags:0,implicit"`
	Iso2         string
	DefaultName  string
	Name         *string `tl:",omitempty:flags:1"`
	CountryCodes []HelpCountryCode
}

func (*HelpCountryPredict) CRC() uint32 {
	return 0xc3878e23
}
func (*HelpCountryPredict) _HelpCountry() {}

type HelpCountryCode interface {
	tl.Object
	_HelpCountryCode()
}

var (
	_ HelpCountryCode = (*HelpCountryCodePredict)(nil)
)

type HelpCountryCodePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	CountryCode string
	Prefixes    []string `tl:",omitempty:flags:0"`
	Patterns    []string `tl:",omitempty:flags:1"`
}

func (*HelpCountryCodePredict) CRC() uint32 {
	return 0x4203c5ef
}
func (*HelpCountryCodePredict) _HelpCountryCode() {}

type HelpDeepLinkInfo interface {
	tl.Object
	_HelpDeepLinkInfo()
}

var (
	_ HelpDeepLinkInfo = (*HelpDeepLinkInfoEmptyPredict)(nil)
	_ HelpDeepLinkInfo = (*HelpDeepLinkInfoPredict)(nil)
)

type HelpDeepLinkInfoEmptyPredict struct{}

func (*HelpDeepLinkInfoEmptyPredict) CRC() uint32 {
	return 0x66afa166
}
func (*HelpDeepLinkInfoEmptyPredict) _HelpDeepLinkInfo() {}

type HelpDeepLinkInfoPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	UpdateApp bool     `tl:",omitempty:flags:0,implicit"`
	Message   string
	Entities  []MessageEntity `tl:",omitempty:flags:1"`
}

func (*HelpDeepLinkInfoPredict) CRC() uint32 {
	return 0x6a4ee832
}
func (*HelpDeepLinkInfoPredict) _HelpDeepLinkInfo() {}

type HelpInviteText interface {
	tl.Object
	_HelpInviteText()
}

var (
	_ HelpInviteText = (*HelpInviteTextPredict)(nil)
)

type HelpInviteTextPredict struct {
	Message string
}

func (*HelpInviteTextPredict) CRC() uint32 {
	return 0x18cb9f78
}
func (*HelpInviteTextPredict) _HelpInviteText() {}

type HelpPassportConfig interface {
	tl.Object
	_HelpPassportConfig()
}

var (
	_ HelpPassportConfig = (*HelpPassportConfigNotModifiedPredict)(nil)
	_ HelpPassportConfig = (*HelpPassportConfigPredict)(nil)
)

type HelpPassportConfigNotModifiedPredict struct{}

func (*HelpPassportConfigNotModifiedPredict) CRC() uint32 {
	return 0xbfb9f457
}
func (*HelpPassportConfigNotModifiedPredict) _HelpPassportConfig() {}

type HelpPassportConfigPredict struct {
	Hash           int32
	CountriesLangs DataJSON
}

func (*HelpPassportConfigPredict) CRC() uint32 {
	return 0xa098d6af
}
func (*HelpPassportConfigPredict) _HelpPassportConfig() {}

type HelpPeerColorOption interface {
	tl.Object
	_HelpPeerColorOption()
}

var (
	_ HelpPeerColorOption = (*HelpPeerColorOptionPredict)(nil)
)

type HelpPeerColorOptionPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Hidden          bool     `tl:",omitempty:flags:0,implicit"`
	ColorID         int32
	Colors          HelpPeerColorSet `tl:",omitempty:flags:1"`
	DarkColors      HelpPeerColorSet `tl:",omitempty:flags:2"`
	ChannelMinLevel *int32           `tl:",omitempty:flags:3"`
	GroupMinLevel   *int32           `tl:",omitempty:flags:4"`
}

func (*HelpPeerColorOptionPredict) CRC() uint32 {
	return 0xadec6ebe
}
func (*HelpPeerColorOptionPredict) _HelpPeerColorOption() {}

type HelpPeerColorSet interface {
	tl.Object
	_HelpPeerColorSet()
}

var (
	_ HelpPeerColorSet = (*HelpPeerColorSetPredict)(nil)
	_ HelpPeerColorSet = (*HelpPeerColorProfileSetPredict)(nil)
)

type HelpPeerColorSetPredict struct {
	Colors []int32
}

func (*HelpPeerColorSetPredict) CRC() uint32 {
	return 0x26219a58
}
func (*HelpPeerColorSetPredict) _HelpPeerColorSet() {}

type HelpPeerColorProfileSetPredict struct {
	PaletteColors []int32
	BgColors      []int32
	StoryColors   []int32
}

func (*HelpPeerColorProfileSetPredict) CRC() uint32 {
	return 0x767d61eb
}
func (*HelpPeerColorProfileSetPredict) _HelpPeerColorSet() {}

type HelpPeerColors interface {
	tl.Object
	_HelpPeerColors()
}

var (
	_ HelpPeerColors = (*HelpPeerColorsNotModifiedPredict)(nil)
	_ HelpPeerColors = (*HelpPeerColorsPredict)(nil)
)

type HelpPeerColorsNotModifiedPredict struct{}

func (*HelpPeerColorsNotModifiedPredict) CRC() uint32 {
	return 0x2ba1f5ce
}
func (*HelpPeerColorsNotModifiedPredict) _HelpPeerColors() {}

type HelpPeerColorsPredict struct {
	Hash   int32
	Colors []HelpPeerColorOption
}

func (*HelpPeerColorsPredict) CRC() uint32 {
	return 0xf8ed08
}
func (*HelpPeerColorsPredict) _HelpPeerColors() {}

type HelpPremiumPromo interface {
	tl.Object
	_HelpPremiumPromo()
}

var (
	_ HelpPremiumPromo = (*HelpPremiumPromoPredict)(nil)
)

type HelpPremiumPromoPredict struct {
	StatusText     string
	StatusEntities []MessageEntity
	VideoSections  []string
	Videos         []Document
	PeriodOptions  []PremiumSubscriptionOption
	Users          []User
}

func (*HelpPremiumPromoPredict) CRC() uint32 {
	return 0x5334759c
}
func (*HelpPremiumPromoPredict) _HelpPremiumPromo() {}

type HelpPromoData interface {
	tl.Object
	_HelpPromoData()
}

var (
	_ HelpPromoData = (*HelpPromoDataEmptyPredict)(nil)
	_ HelpPromoData = (*HelpPromoDataPredict)(nil)
)

type HelpPromoDataEmptyPredict struct {
	Expires int32
}

func (*HelpPromoDataEmptyPredict) CRC() uint32 {
	return 0x98f6ac75
}
func (*HelpPromoDataEmptyPredict) _HelpPromoData() {}

type HelpPromoDataPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Proxy      bool     `tl:",omitempty:flags:0,implicit"`
	Expires    int32
	Peer       Peer
	Chats      []Chat
	Users      []User
	PsaType    *string `tl:",omitempty:flags:1"`
	PsaMessage *string `tl:",omitempty:flags:2"`
}

func (*HelpPromoDataPredict) CRC() uint32 {
	return 0x8c39793f
}
func (*HelpPromoDataPredict) _HelpPromoData() {}

type HelpRecentMeUrls interface {
	tl.Object
	_HelpRecentMeUrls()
}

var (
	_ HelpRecentMeUrls = (*HelpRecentMeUrlsPredict)(nil)
)

type HelpRecentMeUrlsPredict struct {
	Urls  []RecentMeURL
	Chats []Chat
	Users []User
}

func (*HelpRecentMeUrlsPredict) CRC() uint32 {
	return 0xe0310d7
}
func (*HelpRecentMeUrlsPredict) _HelpRecentMeUrls() {}

type HelpSupport interface {
	tl.Object
	_HelpSupport()
}

var (
	_ HelpSupport = (*HelpSupportPredict)(nil)
)

type HelpSupportPredict struct {
	PhoneNumber string
	User        User
}

func (*HelpSupportPredict) CRC() uint32 {
	return 0x17c6b5f6
}
func (*HelpSupportPredict) _HelpSupport() {}

type HelpSupportName interface {
	tl.Object
	_HelpSupportName()
}

var (
	_ HelpSupportName = (*HelpSupportNamePredict)(nil)
)

type HelpSupportNamePredict struct {
	Name string
}

func (*HelpSupportNamePredict) CRC() uint32 {
	return 0x8c05f1c9
}
func (*HelpSupportNamePredict) _HelpSupportName() {}

type HelpTermsOfService interface {
	tl.Object
	_HelpTermsOfService()
}

var (
	_ HelpTermsOfService = (*HelpTermsOfServicePredict)(nil)
)

type HelpTermsOfServicePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Popup         bool     `tl:",omitempty:flags:0,implicit"`
	ID            DataJSON
	Text          string
	Entities      []MessageEntity
	MinAgeConfirm *int32 `tl:",omitempty:flags:1"`
}

func (*HelpTermsOfServicePredict) CRC() uint32 {
	return 0x780a0310
}
func (*HelpTermsOfServicePredict) _HelpTermsOfService() {}

type HelpTermsOfServiceUpdate interface {
	tl.Object
	_HelpTermsOfServiceUpdate()
}

var (
	_ HelpTermsOfServiceUpdate = (*HelpTermsOfServiceUpdateEmptyPredict)(nil)
	_ HelpTermsOfServiceUpdate = (*HelpTermsOfServiceUpdatePredict)(nil)
)

type HelpTermsOfServiceUpdateEmptyPredict struct {
	Expires int32
}

func (*HelpTermsOfServiceUpdateEmptyPredict) CRC() uint32 {
	return 0xe3309f7f
}
func (*HelpTermsOfServiceUpdateEmptyPredict) _HelpTermsOfServiceUpdate() {}

type HelpTermsOfServiceUpdatePredict struct {
	Expires        int32
	TermsOfService HelpTermsOfService
}

func (*HelpTermsOfServiceUpdatePredict) CRC() uint32 {
	return 0x28ecf961
}
func (*HelpTermsOfServiceUpdatePredict) _HelpTermsOfServiceUpdate() {}

type HelpTimezonesList interface {
	tl.Object
	_HelpTimezonesList()
}

var (
	_ HelpTimezonesList = (*HelpTimezonesListNotModifiedPredict)(nil)
	_ HelpTimezonesList = (*HelpTimezonesListPredict)(nil)
)

type HelpTimezonesListNotModifiedPredict struct{}

func (*HelpTimezonesListNotModifiedPredict) CRC() uint32 {
	return 0x970708cc
}
func (*HelpTimezonesListNotModifiedPredict) _HelpTimezonesList() {}

type HelpTimezonesListPredict struct {
	Timezones []Timezone
	Hash      int32
}

func (*HelpTimezonesListPredict) CRC() uint32 {
	return 0x7b74ed71
}
func (*HelpTimezonesListPredict) _HelpTimezonesList() {}

type HelpUserInfo interface {
	tl.Object
	_HelpUserInfo()
}

var (
	_ HelpUserInfo = (*HelpUserInfoEmptyPredict)(nil)
	_ HelpUserInfo = (*HelpUserInfoPredict)(nil)
)

type HelpUserInfoEmptyPredict struct{}

func (*HelpUserInfoEmptyPredict) CRC() uint32 {
	return 0xf3ae2eed
}
func (*HelpUserInfoEmptyPredict) _HelpUserInfo() {}

type HelpUserInfoPredict struct {
	Message  string
	Entities []MessageEntity
	Author   string
	Date     int32
}

func (*HelpUserInfoPredict) CRC() uint32 {
	return 0x1eb3758
}
func (*HelpUserInfoPredict) _HelpUserInfo() {}

type MessagesAffectedFoundMessages interface {
	tl.Object
	_MessagesAffectedFoundMessages()
}

var (
	_ MessagesAffectedFoundMessages = (*MessagesAffectedFoundMessagesPredict)(nil)
)

type MessagesAffectedFoundMessagesPredict struct {
	Pts      int32
	PtsCount int32
	Offset   int32
	Messages []int32
}

func (*MessagesAffectedFoundMessagesPredict) CRC() uint32 {
	return 0xef8d3e6c
}
func (*MessagesAffectedFoundMessagesPredict) _MessagesAffectedFoundMessages() {}

type MessagesAffectedHistory interface {
	tl.Object
	_MessagesAffectedHistory()
}

var (
	_ MessagesAffectedHistory = (*MessagesAffectedHistoryPredict)(nil)
)

type MessagesAffectedHistoryPredict struct {
	Pts      int32
	PtsCount int32
	Offset   int32
}

func (*MessagesAffectedHistoryPredict) CRC() uint32 {
	return 0xb45c69d1
}
func (*MessagesAffectedHistoryPredict) _MessagesAffectedHistory() {}

type MessagesAffectedMessages interface {
	tl.Object
	_MessagesAffectedMessages()
}

var (
	_ MessagesAffectedMessages = (*MessagesAffectedMessagesPredict)(nil)
)

type MessagesAffectedMessagesPredict struct {
	Pts      int32
	PtsCount int32
}

func (*MessagesAffectedMessagesPredict) CRC() uint32 {
	return 0x84d19185
}
func (*MessagesAffectedMessagesPredict) _MessagesAffectedMessages() {}

type MessagesAllStickers interface {
	tl.Object
	_MessagesAllStickers()
}

var (
	_ MessagesAllStickers = (*MessagesAllStickersNotModifiedPredict)(nil)
	_ MessagesAllStickers = (*MessagesAllStickersPredict)(nil)
)

type MessagesAllStickersNotModifiedPredict struct{}

func (*MessagesAllStickersNotModifiedPredict) CRC() uint32 {
	return 0xe86602c3
}
func (*MessagesAllStickersNotModifiedPredict) _MessagesAllStickers() {}

type MessagesAllStickersPredict struct {
	Hash int64
	Sets []StickerSet
}

func (*MessagesAllStickersPredict) CRC() uint32 {
	return 0xcdbbcebb
}
func (*MessagesAllStickersPredict) _MessagesAllStickers() {}

type MessagesArchivedStickers interface {
	tl.Object
	_MessagesArchivedStickers()
}

var (
	_ MessagesArchivedStickers = (*MessagesArchivedStickersPredict)(nil)
)

type MessagesArchivedStickersPredict struct {
	Count int32
	Sets  []StickerSetCovered
}

func (*MessagesArchivedStickersPredict) CRC() uint32 {
	return 0x4fcba9c8
}
func (*MessagesArchivedStickersPredict) _MessagesArchivedStickers() {}

type MessagesAvailableEffects interface {
	tl.Object
	_MessagesAvailableEffects()
}

var (
	_ MessagesAvailableEffects = (*MessagesAvailableEffectsNotModifiedPredict)(nil)
	_ MessagesAvailableEffects = (*MessagesAvailableEffectsPredict)(nil)
)

type MessagesAvailableEffectsNotModifiedPredict struct{}

func (*MessagesAvailableEffectsNotModifiedPredict) CRC() uint32 {
	return 0xd1ed9a5b
}
func (*MessagesAvailableEffectsNotModifiedPredict) _MessagesAvailableEffects() {}

type MessagesAvailableEffectsPredict struct {
	Hash      int32
	Effects   []AvailableEffect
	Documents []Document
}

func (*MessagesAvailableEffectsPredict) CRC() uint32 {
	return 0xbddb616e
}
func (*MessagesAvailableEffectsPredict) _MessagesAvailableEffects() {}

type MessagesAvailableReactions interface {
	tl.Object
	_MessagesAvailableReactions()
}

var (
	_ MessagesAvailableReactions = (*MessagesAvailableReactionsNotModifiedPredict)(nil)
	_ MessagesAvailableReactions = (*MessagesAvailableReactionsPredict)(nil)
)

type MessagesAvailableReactionsNotModifiedPredict struct{}

func (*MessagesAvailableReactionsNotModifiedPredict) CRC() uint32 {
	return 0x9f071957
}
func (*MessagesAvailableReactionsNotModifiedPredict) _MessagesAvailableReactions() {}

type MessagesAvailableReactionsPredict struct {
	Hash      int32
	Reactions []AvailableReaction
}

func (*MessagesAvailableReactionsPredict) CRC() uint32 {
	return 0x768e3aad
}
func (*MessagesAvailableReactionsPredict) _MessagesAvailableReactions() {}

type MessagesBotApp interface {
	tl.Object
	_MessagesBotApp()
}

var (
	_ MessagesBotApp = (*MessagesBotAppPredict)(nil)
)

type MessagesBotAppPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Inactive           bool     `tl:",omitempty:flags:0,implicit"`
	RequestWriteAccess bool     `tl:",omitempty:flags:1,implicit"`
	HasSettings        bool     `tl:",omitempty:flags:2,implicit"`
	App                BotApp
}

func (*MessagesBotAppPredict) CRC() uint32 {
	return 0xeb50adf5
}
func (*MessagesBotAppPredict) _MessagesBotApp() {}

type MessagesBotCallbackAnswer interface {
	tl.Object
	_MessagesBotCallbackAnswer()
}

var (
	_ MessagesBotCallbackAnswer = (*MessagesBotCallbackAnswerPredict)(nil)
)

type MessagesBotCallbackAnswerPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Alert     bool     `tl:",omitempty:flags:1,implicit"`
	HasURL    bool     `tl:",omitempty:flags:3,implicit"`
	NativeUi  bool     `tl:",omitempty:flags:4,implicit"`
	Message   *string  `tl:",omitempty:flags:0"`
	URL       *string  `tl:",omitempty:flags:2"`
	CacheTime int32
}

func (*MessagesBotCallbackAnswerPredict) CRC() uint32 {
	return 0x36585ea4
}
func (*MessagesBotCallbackAnswerPredict) _MessagesBotCallbackAnswer() {}

type MessagesBotResults interface {
	tl.Object
	_MessagesBotResults()
}

var (
	_ MessagesBotResults = (*MessagesBotResultsPredict)(nil)
)

type MessagesBotResultsPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Gallery       bool     `tl:",omitempty:flags:0,implicit"`
	QueryID       int64
	NextOffset    *string           `tl:",omitempty:flags:1"`
	SwitchPm      InlineBotSwitchPm `tl:",omitempty:flags:2"`
	SwitchWebview InlineBotWebView  `tl:",omitempty:flags:3"`
	Results       []BotInlineResult
	CacheTime     int32
	Users         []User
}

func (*MessagesBotResultsPredict) CRC() uint32 {
	return 0xe021f2f6
}
func (*MessagesBotResultsPredict) _MessagesBotResults() {}

type MessagesChatAdminsWithInvites interface {
	tl.Object
	_MessagesChatAdminsWithInvites()
}

var (
	_ MessagesChatAdminsWithInvites = (*MessagesChatAdminsWithInvitesPredict)(nil)
)

type MessagesChatAdminsWithInvitesPredict struct {
	Admins []ChatAdminWithInvites
	Users  []User
}

func (*MessagesChatAdminsWithInvitesPredict) CRC() uint32 {
	return 0xb69b72d7
}
func (*MessagesChatAdminsWithInvitesPredict) _MessagesChatAdminsWithInvites() {}

type MessagesChatFull interface {
	tl.Object
	_MessagesChatFull()
}

var (
	_ MessagesChatFull = (*MessagesChatFullPredict)(nil)
)

type MessagesChatFullPredict struct {
	FullChat ChatFull
	Chats    []Chat
	Users    []User
}

func (*MessagesChatFullPredict) CRC() uint32 {
	return 0xe5d7d19c
}
func (*MessagesChatFullPredict) _MessagesChatFull() {}

type MessagesChatInviteImporters interface {
	tl.Object
	_MessagesChatInviteImporters()
}

var (
	_ MessagesChatInviteImporters = (*MessagesChatInviteImportersPredict)(nil)
)

type MessagesChatInviteImportersPredict struct {
	Count     int32
	Importers []ChatInviteImporter
	Users     []User
}

func (*MessagesChatInviteImportersPredict) CRC() uint32 {
	return 0x81b6b00a
}
func (*MessagesChatInviteImportersPredict) _MessagesChatInviteImporters() {}

type MessagesChats interface {
	tl.Object
	_MessagesChats()
}

var (
	_ MessagesChats = (*MessagesChatsPredict)(nil)
	_ MessagesChats = (*MessagesChatsSlicePredict)(nil)
)

type MessagesChatsPredict struct {
	Chats []Chat
}

func (*MessagesChatsPredict) CRC() uint32 {
	return 0x64ff9fd5
}
func (*MessagesChatsPredict) _MessagesChats() {}

type MessagesChatsSlicePredict struct {
	Count int32
	Chats []Chat
}

func (*MessagesChatsSlicePredict) CRC() uint32 {
	return 0x9cd81144
}
func (*MessagesChatsSlicePredict) _MessagesChats() {}

type MessagesCheckedHistoryImportPeer interface {
	tl.Object
	_MessagesCheckedHistoryImportPeer()
}

var (
	_ MessagesCheckedHistoryImportPeer = (*MessagesCheckedHistoryImportPeerPredict)(nil)
)

type MessagesCheckedHistoryImportPeerPredict struct {
	ConfirmText string
}

func (*MessagesCheckedHistoryImportPeerPredict) CRC() uint32 {
	return 0xa24de717
}
func (*MessagesCheckedHistoryImportPeerPredict) _MessagesCheckedHistoryImportPeer() {}

type MessagesDhConfig interface {
	tl.Object
	_MessagesDhConfig()
}

var (
	_ MessagesDhConfig = (*MessagesDhConfigNotModifiedPredict)(nil)
	_ MessagesDhConfig = (*MessagesDhConfigPredict)(nil)
)

type MessagesDhConfigNotModifiedPredict struct {
	Random []byte
}

func (*MessagesDhConfigNotModifiedPredict) CRC() uint32 {
	return 0xc0e24635
}
func (*MessagesDhConfigNotModifiedPredict) _MessagesDhConfig() {}

type MessagesDhConfigPredict struct {
	G       int32
	P       []byte
	Version int32
	Random  []byte
}

func (*MessagesDhConfigPredict) CRC() uint32 {
	return 0x2c221edd
}
func (*MessagesDhConfigPredict) _MessagesDhConfig() {}

type MessagesDialogFilters interface {
	tl.Object
	_MessagesDialogFilters()
}

var (
	_ MessagesDialogFilters = (*MessagesDialogFiltersPredict)(nil)
)

type MessagesDialogFiltersPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	TagsEnabled bool     `tl:",omitempty:flags:0,implicit"`
	Filters     []DialogFilter
}

func (*MessagesDialogFiltersPredict) CRC() uint32 {
	return 0x2ad93719
}
func (*MessagesDialogFiltersPredict) _MessagesDialogFilters() {}

type MessagesDialogs interface {
	tl.Object
	_MessagesDialogs()
}

var (
	_ MessagesDialogs = (*MessagesDialogsPredict)(nil)
	_ MessagesDialogs = (*MessagesDialogsSlicePredict)(nil)
	_ MessagesDialogs = (*MessagesDialogsNotModifiedPredict)(nil)
)

type MessagesDialogsPredict struct {
	Dialogs  []Dialog
	Messages []Message
	Chats    []Chat
	Users    []User
}

func (*MessagesDialogsPredict) CRC() uint32 {
	return 0x15ba6c40
}
func (*MessagesDialogsPredict) _MessagesDialogs() {}

type MessagesDialogsSlicePredict struct {
	Count    int32
	Dialogs  []Dialog
	Messages []Message
	Chats    []Chat
	Users    []User
}

func (*MessagesDialogsSlicePredict) CRC() uint32 {
	return 0x71e094f3
}
func (*MessagesDialogsSlicePredict) _MessagesDialogs() {}

type MessagesDialogsNotModifiedPredict struct {
	Count int32
}

func (*MessagesDialogsNotModifiedPredict) CRC() uint32 {
	return 0xf0e3e596
}
func (*MessagesDialogsNotModifiedPredict) _MessagesDialogs() {}

type MessagesDiscussionMessage interface {
	tl.Object
	_MessagesDiscussionMessage()
}

var (
	_ MessagesDiscussionMessage = (*MessagesDiscussionMessagePredict)(nil)
)

type MessagesDiscussionMessagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Messages        []Message
	MaxID           *int32 `tl:",omitempty:flags:0"`
	ReadInboxMaxID  *int32 `tl:",omitempty:flags:1"`
	ReadOutboxMaxID *int32 `tl:",omitempty:flags:2"`
	UnreadCount     int32
	Chats           []Chat
	Users           []User
}

func (*MessagesDiscussionMessagePredict) CRC() uint32 {
	return 0xa6341782
}
func (*MessagesDiscussionMessagePredict) _MessagesDiscussionMessage() {}

type MessagesEmojiGroups interface {
	tl.Object
	_MessagesEmojiGroups()
}

var (
	_ MessagesEmojiGroups = (*MessagesEmojiGroupsNotModifiedPredict)(nil)
	_ MessagesEmojiGroups = (*MessagesEmojiGroupsPredict)(nil)
)

type MessagesEmojiGroupsNotModifiedPredict struct{}

func (*MessagesEmojiGroupsNotModifiedPredict) CRC() uint32 {
	return 0x6fb4ad87
}
func (*MessagesEmojiGroupsNotModifiedPredict) _MessagesEmojiGroups() {}

type MessagesEmojiGroupsPredict struct {
	Hash   int32
	Groups []EmojiGroup
}

func (*MessagesEmojiGroupsPredict) CRC() uint32 {
	return 0x881fb94b
}
func (*MessagesEmojiGroupsPredict) _MessagesEmojiGroups() {}

type MessagesExportedChatInvite interface {
	tl.Object
	_MessagesExportedChatInvite()
}

var (
	_ MessagesExportedChatInvite = (*MessagesExportedChatInvitePredict)(nil)
	_ MessagesExportedChatInvite = (*MessagesExportedChatInviteReplacedPredict)(nil)
)

type MessagesExportedChatInvitePredict struct {
	Invite ExportedChatInvite
	Users  []User
}

func (*MessagesExportedChatInvitePredict) CRC() uint32 {
	return 0x1871be50
}
func (*MessagesExportedChatInvitePredict) _MessagesExportedChatInvite() {}

type MessagesExportedChatInviteReplacedPredict struct {
	Invite    ExportedChatInvite
	NewInvite ExportedChatInvite
	Users     []User
}

func (*MessagesExportedChatInviteReplacedPredict) CRC() uint32 {
	return 0x222600ef
}
func (*MessagesExportedChatInviteReplacedPredict) _MessagesExportedChatInvite() {}

type MessagesExportedChatInvites interface {
	tl.Object
	_MessagesExportedChatInvites()
}

var (
	_ MessagesExportedChatInvites = (*MessagesExportedChatInvitesPredict)(nil)
)

type MessagesExportedChatInvitesPredict struct {
	Count   int32
	Invites []ExportedChatInvite
	Users   []User
}

func (*MessagesExportedChatInvitesPredict) CRC() uint32 {
	return 0xbdc62dcc
}
func (*MessagesExportedChatInvitesPredict) _MessagesExportedChatInvites() {}

type MessagesFavedStickers interface {
	tl.Object
	_MessagesFavedStickers()
}

var (
	_ MessagesFavedStickers = (*MessagesFavedStickersNotModifiedPredict)(nil)
	_ MessagesFavedStickers = (*MessagesFavedStickersPredict)(nil)
)

type MessagesFavedStickersNotModifiedPredict struct{}

func (*MessagesFavedStickersNotModifiedPredict) CRC() uint32 {
	return 0x9e8fa6d3
}
func (*MessagesFavedStickersNotModifiedPredict) _MessagesFavedStickers() {}

type MessagesFavedStickersPredict struct {
	Hash     int64
	Packs    []StickerPack
	Stickers []Document
}

func (*MessagesFavedStickersPredict) CRC() uint32 {
	return 0x2cb51097
}
func (*MessagesFavedStickersPredict) _MessagesFavedStickers() {}

type MessagesFeaturedStickers interface {
	tl.Object
	_MessagesFeaturedStickers()
}

var (
	_ MessagesFeaturedStickers = (*MessagesFeaturedStickersNotModifiedPredict)(nil)
	_ MessagesFeaturedStickers = (*MessagesFeaturedStickersPredict)(nil)
)

type MessagesFeaturedStickersNotModifiedPredict struct {
	Count int32
}

func (*MessagesFeaturedStickersNotModifiedPredict) CRC() uint32 {
	return 0xc6dc0c66
}
func (*MessagesFeaturedStickersNotModifiedPredict) _MessagesFeaturedStickers() {}

type MessagesFeaturedStickersPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Premium bool     `tl:",omitempty:flags:0,implicit"`
	Hash    int64
	Count   int32
	Sets    []StickerSetCovered
	Unread  []int64
}

func (*MessagesFeaturedStickersPredict) CRC() uint32 {
	return 0xbe382906
}
func (*MessagesFeaturedStickersPredict) _MessagesFeaturedStickers() {}

type MessagesForumTopics interface {
	tl.Object
	_MessagesForumTopics()
}

var (
	_ MessagesForumTopics = (*MessagesForumTopicsPredict)(nil)
)

type MessagesForumTopicsPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	OrderByCreateDate bool     `tl:",omitempty:flags:0,implicit"`
	Count             int32
	Topics            []ForumTopic
	Messages          []Message
	Chats             []Chat
	Users             []User
	Pts               int32
}

func (*MessagesForumTopicsPredict) CRC() uint32 {
	return 0x367617d3
}
func (*MessagesForumTopicsPredict) _MessagesForumTopics() {}

type MessagesFoundStickerSets interface {
	tl.Object
	_MessagesFoundStickerSets()
}

var (
	_ MessagesFoundStickerSets = (*MessagesFoundStickerSetsNotModifiedPredict)(nil)
	_ MessagesFoundStickerSets = (*MessagesFoundStickerSetsPredict)(nil)
)

type MessagesFoundStickerSetsNotModifiedPredict struct{}

func (*MessagesFoundStickerSetsNotModifiedPredict) CRC() uint32 {
	return 0xd54b65d
}
func (*MessagesFoundStickerSetsNotModifiedPredict) _MessagesFoundStickerSets() {}

type MessagesFoundStickerSetsPredict struct {
	Hash int64
	Sets []StickerSetCovered
}

func (*MessagesFoundStickerSetsPredict) CRC() uint32 {
	return 0x8af09dd2
}
func (*MessagesFoundStickerSetsPredict) _MessagesFoundStickerSets() {}

type MessagesHighScores interface {
	tl.Object
	_MessagesHighScores()
}

var (
	_ MessagesHighScores = (*MessagesHighScoresPredict)(nil)
)

type MessagesHighScoresPredict struct {
	Scores []HighScore
	Users  []User
}

func (*MessagesHighScoresPredict) CRC() uint32 {
	return 0x9a3bfd99
}
func (*MessagesHighScoresPredict) _MessagesHighScores() {}

type MessagesHistoryImport interface {
	tl.Object
	_MessagesHistoryImport()
}

var (
	_ MessagesHistoryImport = (*MessagesHistoryImportPredict)(nil)
)

type MessagesHistoryImportPredict struct {
	ID int64
}

func (*MessagesHistoryImportPredict) CRC() uint32 {
	return 0x1662af0b
}
func (*MessagesHistoryImportPredict) _MessagesHistoryImport() {}

type MessagesHistoryImportParsed interface {
	tl.Object
	_MessagesHistoryImportParsed()
}

var (
	_ MessagesHistoryImportParsed = (*MessagesHistoryImportParsedPredict)(nil)
)

type MessagesHistoryImportParsedPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Pm    bool     `tl:",omitempty:flags:0,implicit"`
	Group bool     `tl:",omitempty:flags:1,implicit"`
	Title *string  `tl:",omitempty:flags:2"`
}

func (*MessagesHistoryImportParsedPredict) CRC() uint32 {
	return 0x5e0fb7b9
}
func (*MessagesHistoryImportParsedPredict) _MessagesHistoryImportParsed() {}

type MessagesInactiveChats interface {
	tl.Object
	_MessagesInactiveChats()
}

var (
	_ MessagesInactiveChats = (*MessagesInactiveChatsPredict)(nil)
)

type MessagesInactiveChatsPredict struct {
	Dates []int32
	Chats []Chat
	Users []User
}

func (*MessagesInactiveChatsPredict) CRC() uint32 {
	return 0xa927fec5
}
func (*MessagesInactiveChatsPredict) _MessagesInactiveChats() {}

type MessagesInvitedUsers interface {
	tl.Object
	_MessagesInvitedUsers()
}

var (
	_ MessagesInvitedUsers = (*MessagesInvitedUsersPredict)(nil)
)

type MessagesInvitedUsersPredict struct {
	Updates         Updates
	MissingInvitees []MissingInvitee
}

func (*MessagesInvitedUsersPredict) CRC() uint32 {
	return 0x7f5defa6
}
func (*MessagesInvitedUsersPredict) _MessagesInvitedUsers() {}

type MessagesMessageEditData interface {
	tl.Object
	_MessagesMessageEditData()
}

var (
	_ MessagesMessageEditData = (*MessagesMessageEditDataPredict)(nil)
)

type MessagesMessageEditDataPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Caption bool     `tl:",omitempty:flags:0,implicit"`
}

func (*MessagesMessageEditDataPredict) CRC() uint32 {
	return 0x26b5dde6
}
func (*MessagesMessageEditDataPredict) _MessagesMessageEditData() {}

type MessagesMessageReactionsList interface {
	tl.Object
	_MessagesMessageReactionsList()
}

var (
	_ MessagesMessageReactionsList = (*MessagesMessageReactionsListPredict)(nil)
)

type MessagesMessageReactionsListPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Reactions  []MessagePeerReaction
	Chats      []Chat
	Users      []User
	NextOffset *string `tl:",omitempty:flags:0"`
}

func (*MessagesMessageReactionsListPredict) CRC() uint32 {
	return 0x31bd492d
}
func (*MessagesMessageReactionsListPredict) _MessagesMessageReactionsList() {}

type MessagesMessageViews interface {
	tl.Object
	_MessagesMessageViews()
}

var (
	_ MessagesMessageViews = (*MessagesMessageViewsPredict)(nil)
)

type MessagesMessageViewsPredict struct {
	Views []MessageViews
	Chats []Chat
	Users []User
}

func (*MessagesMessageViewsPredict) CRC() uint32 {
	return 0xb6c4f543
}
func (*MessagesMessageViewsPredict) _MessagesMessageViews() {}

type MessagesMessages interface {
	tl.Object
	_MessagesMessages()
}

var (
	_ MessagesMessages = (*MessagesMessagesPredict)(nil)
	_ MessagesMessages = (*MessagesMessagesSlicePredict)(nil)
	_ MessagesMessages = (*MessagesChannelMessagesPredict)(nil)
	_ MessagesMessages = (*MessagesMessagesNotModifiedPredict)(nil)
)

type MessagesMessagesPredict struct {
	Messages []Message
	Chats    []Chat
	Users    []User
}

func (*MessagesMessagesPredict) CRC() uint32 {
	return 0x8c718e87
}
func (*MessagesMessagesPredict) _MessagesMessages() {}

type MessagesMessagesSlicePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Inexact        bool     `tl:",omitempty:flags:1,implicit"`
	Count          int32
	NextRate       *int32 `tl:",omitempty:flags:0"`
	OffsetIDOffset *int32 `tl:",omitempty:flags:2"`
	Messages       []Message
	Chats          []Chat
	Users          []User
}

func (*MessagesMessagesSlicePredict) CRC() uint32 {
	return 0x3a54685e
}
func (*MessagesMessagesSlicePredict) _MessagesMessages() {}

type MessagesChannelMessagesPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Inexact        bool     `tl:",omitempty:flags:1,implicit"`
	Pts            int32
	Count          int32
	OffsetIDOffset *int32 `tl:",omitempty:flags:2"`
	Messages       []Message
	Topics         []ForumTopic
	Chats          []Chat
	Users          []User
}

func (*MessagesChannelMessagesPredict) CRC() uint32 {
	return 0xc776ba4e
}
func (*MessagesChannelMessagesPredict) _MessagesMessages() {}

type MessagesMessagesNotModifiedPredict struct {
	Count int32
}

func (*MessagesMessagesNotModifiedPredict) CRC() uint32 {
	return 0x74535f21
}
func (*MessagesMessagesNotModifiedPredict) _MessagesMessages() {}

type MessagesMyStickers interface {
	tl.Object
	_MessagesMyStickers()
}

var (
	_ MessagesMyStickers = (*MessagesMyStickersPredict)(nil)
)

type MessagesMyStickersPredict struct {
	Count int32
	Sets  []StickerSetCovered
}

func (*MessagesMyStickersPredict) CRC() uint32 {
	return 0xfaff629d
}
func (*MessagesMyStickersPredict) _MessagesMyStickers() {}

type MessagesPeerDialogs interface {
	tl.Object
	_MessagesPeerDialogs()
}

var (
	_ MessagesPeerDialogs = (*MessagesPeerDialogsPredict)(nil)
)

type MessagesPeerDialogsPredict struct {
	Dialogs  []Dialog
	Messages []Message
	Chats    []Chat
	Users    []User
	State    UpdatesState
}

func (*MessagesPeerDialogsPredict) CRC() uint32 {
	return 0x3371c354
}
func (*MessagesPeerDialogsPredict) _MessagesPeerDialogs() {}

type MessagesPeerSettings interface {
	tl.Object
	_MessagesPeerSettings()
}

var (
	_ MessagesPeerSettings = (*MessagesPeerSettingsPredict)(nil)
)

type MessagesPeerSettingsPredict struct {
	Settings PeerSettings
	Chats    []Chat
	Users    []User
}

func (*MessagesPeerSettingsPredict) CRC() uint32 {
	return 0x6880b94d
}
func (*MessagesPeerSettingsPredict) _MessagesPeerSettings() {}

type MessagesQuickReplies interface {
	tl.Object
	_MessagesQuickReplies()
}

var (
	_ MessagesQuickReplies = (*MessagesQuickRepliesPredict)(nil)
	_ MessagesQuickReplies = (*MessagesQuickRepliesNotModifiedPredict)(nil)
)

type MessagesQuickRepliesPredict struct {
	QuickReplies []QuickReply
	Messages     []Message
	Chats        []Chat
	Users        []User
}

func (*MessagesQuickRepliesPredict) CRC() uint32 {
	return 0xc68d6695
}
func (*MessagesQuickRepliesPredict) _MessagesQuickReplies() {}

type MessagesQuickRepliesNotModifiedPredict struct{}

func (*MessagesQuickRepliesNotModifiedPredict) CRC() uint32 {
	return 0x5f91eb5b
}
func (*MessagesQuickRepliesNotModifiedPredict) _MessagesQuickReplies() {}

type MessagesReactions interface {
	tl.Object
	_MessagesReactions()
}

var (
	_ MessagesReactions = (*MessagesReactionsNotModifiedPredict)(nil)
	_ MessagesReactions = (*MessagesReactionsPredict)(nil)
)

type MessagesReactionsNotModifiedPredict struct{}

func (*MessagesReactionsNotModifiedPredict) CRC() uint32 {
	return 0xb06fdbdf
}
func (*MessagesReactionsNotModifiedPredict) _MessagesReactions() {}

type MessagesReactionsPredict struct {
	Hash      int64
	Reactions []Reaction
}

func (*MessagesReactionsPredict) CRC() uint32 {
	return 0xeafdf716
}
func (*MessagesReactionsPredict) _MessagesReactions() {}

type MessagesRecentStickers interface {
	tl.Object
	_MessagesRecentStickers()
}

var (
	_ MessagesRecentStickers = (*MessagesRecentStickersNotModifiedPredict)(nil)
	_ MessagesRecentStickers = (*MessagesRecentStickersPredict)(nil)
)

type MessagesRecentStickersNotModifiedPredict struct{}

func (*MessagesRecentStickersNotModifiedPredict) CRC() uint32 {
	return 0xb17f890
}
func (*MessagesRecentStickersNotModifiedPredict) _MessagesRecentStickers() {}

type MessagesRecentStickersPredict struct {
	Hash     int64
	Packs    []StickerPack
	Stickers []Document
	Dates    []int32
}

func (*MessagesRecentStickersPredict) CRC() uint32 {
	return 0x88d37c56
}
func (*MessagesRecentStickersPredict) _MessagesRecentStickers() {}

type MessagesSavedDialogs interface {
	tl.Object
	_MessagesSavedDialogs()
}

var (
	_ MessagesSavedDialogs = (*MessagesSavedDialogsPredict)(nil)
	_ MessagesSavedDialogs = (*MessagesSavedDialogsSlicePredict)(nil)
	_ MessagesSavedDialogs = (*MessagesSavedDialogsNotModifiedPredict)(nil)
)

type MessagesSavedDialogsPredict struct {
	Dialogs  []SavedDialog
	Messages []Message
	Chats    []Chat
	Users    []User
}

func (*MessagesSavedDialogsPredict) CRC() uint32 {
	return 0xf83ae221
}
func (*MessagesSavedDialogsPredict) _MessagesSavedDialogs() {}

type MessagesSavedDialogsSlicePredict struct {
	Count    int32
	Dialogs  []SavedDialog
	Messages []Message
	Chats    []Chat
	Users    []User
}

func (*MessagesSavedDialogsSlicePredict) CRC() uint32 {
	return 0x44ba9dd9
}
func (*MessagesSavedDialogsSlicePredict) _MessagesSavedDialogs() {}

type MessagesSavedDialogsNotModifiedPredict struct {
	Count int32
}

func (*MessagesSavedDialogsNotModifiedPredict) CRC() uint32 {
	return 0xc01f6fe8
}
func (*MessagesSavedDialogsNotModifiedPredict) _MessagesSavedDialogs() {}

type MessagesSavedGifs interface {
	tl.Object
	_MessagesSavedGifs()
}

var (
	_ MessagesSavedGifs = (*MessagesSavedGifsNotModifiedPredict)(nil)
	_ MessagesSavedGifs = (*MessagesSavedGifsPredict)(nil)
)

type MessagesSavedGifsNotModifiedPredict struct{}

func (*MessagesSavedGifsNotModifiedPredict) CRC() uint32 {
	return 0xe8025ca2
}
func (*MessagesSavedGifsNotModifiedPredict) _MessagesSavedGifs() {}

type MessagesSavedGifsPredict struct {
	Hash int64
	Gifs []Document
}

func (*MessagesSavedGifsPredict) CRC() uint32 {
	return 0x84a02a0d
}
func (*MessagesSavedGifsPredict) _MessagesSavedGifs() {}

type MessagesSavedReactionTags interface {
	tl.Object
	_MessagesSavedReactionTags()
}

var (
	_ MessagesSavedReactionTags = (*MessagesSavedReactionTagsNotModifiedPredict)(nil)
	_ MessagesSavedReactionTags = (*MessagesSavedReactionTagsPredict)(nil)
)

type MessagesSavedReactionTagsNotModifiedPredict struct{}

func (*MessagesSavedReactionTagsNotModifiedPredict) CRC() uint32 {
	return 0x889b59ef
}
func (*MessagesSavedReactionTagsNotModifiedPredict) _MessagesSavedReactionTags() {}

type MessagesSavedReactionTagsPredict struct {
	Tags []SavedReactionTag
	Hash int64
}

func (*MessagesSavedReactionTagsPredict) CRC() uint32 {
	return 0x3259950a
}
func (*MessagesSavedReactionTagsPredict) _MessagesSavedReactionTags() {}

type MessagesSearchCounter interface {
	tl.Object
	_MessagesSearchCounter()
}

var (
	_ MessagesSearchCounter = (*MessagesSearchCounterPredict)(nil)
)

type MessagesSearchCounterPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Inexact bool     `tl:",omitempty:flags:1,implicit"`
	Filter  MessagesFilter
	Count   int32
}

func (*MessagesSearchCounterPredict) CRC() uint32 {
	return 0xe844ebff
}
func (*MessagesSearchCounterPredict) _MessagesSearchCounter() {}

type MessagesSearchResultsCalendar interface {
	tl.Object
	_MessagesSearchResultsCalendar()
}

var (
	_ MessagesSearchResultsCalendar = (*MessagesSearchResultsCalendarPredict)(nil)
)

type MessagesSearchResultsCalendarPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Inexact        bool     `tl:",omitempty:flags:0,implicit"`
	Count          int32
	MinDate        int32
	MinMsgID       int32
	OffsetIDOffset *int32 `tl:",omitempty:flags:1"`
	Periods        []SearchResultsCalendarPeriod
	Messages       []Message
	Chats          []Chat
	Users          []User
}

func (*MessagesSearchResultsCalendarPredict) CRC() uint32 {
	return 0x147ee23c
}
func (*MessagesSearchResultsCalendarPredict) _MessagesSearchResultsCalendar() {}

type MessagesSearchResultsPositions interface {
	tl.Object
	_MessagesSearchResultsPositions()
}

var (
	_ MessagesSearchResultsPositions = (*MessagesSearchResultsPositionsPredict)(nil)
)

type MessagesSearchResultsPositionsPredict struct {
	Count     int32
	Positions []SearchResultsPosition
}

func (*MessagesSearchResultsPositionsPredict) CRC() uint32 {
	return 0x53b22baf
}
func (*MessagesSearchResultsPositionsPredict) _MessagesSearchResultsPositions() {}

type MessagesSentEncryptedMessage interface {
	tl.Object
	_MessagesSentEncryptedMessage()
}

var (
	_ MessagesSentEncryptedMessage = (*MessagesSentEncryptedMessagePredict)(nil)
	_ MessagesSentEncryptedMessage = (*MessagesSentEncryptedFilePredict)(nil)
)

type MessagesSentEncryptedMessagePredict struct {
	Date int32
}

func (*MessagesSentEncryptedMessagePredict) CRC() uint32 {
	return 0x560f8935
}
func (*MessagesSentEncryptedMessagePredict) _MessagesSentEncryptedMessage() {}

type MessagesSentEncryptedFilePredict struct {
	Date int32
	File EncryptedFile
}

func (*MessagesSentEncryptedFilePredict) CRC() uint32 {
	return 0x9493ff32
}
func (*MessagesSentEncryptedFilePredict) _MessagesSentEncryptedMessage() {}

type MessagesSponsoredMessages interface {
	tl.Object
	_MessagesSponsoredMessages()
}

var (
	_ MessagesSponsoredMessages = (*MessagesSponsoredMessagesPredict)(nil)
	_ MessagesSponsoredMessages = (*MessagesSponsoredMessagesEmptyPredict)(nil)
)

type MessagesSponsoredMessagesPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	PostsBetween *int32   `tl:",omitempty:flags:0"`
	Messages     []SponsoredMessage
	Chats        []Chat
	Users        []User
}

func (*MessagesSponsoredMessagesPredict) CRC() uint32 {
	return 0xc9ee1d87
}
func (*MessagesSponsoredMessagesPredict) _MessagesSponsoredMessages() {}

type MessagesSponsoredMessagesEmptyPredict struct{}

func (*MessagesSponsoredMessagesEmptyPredict) CRC() uint32 {
	return 0x1839490f
}
func (*MessagesSponsoredMessagesEmptyPredict) _MessagesSponsoredMessages() {}

type MessagesStickerSet interface {
	tl.Object
	_MessagesStickerSet()
}

var (
	_ MessagesStickerSet = (*MessagesStickerSetPredict)(nil)
	_ MessagesStickerSet = (*MessagesStickerSetNotModifiedPredict)(nil)
)

type MessagesStickerSetPredict struct {
	Set       StickerSet
	Packs     []StickerPack
	Keywords  []StickerKeyword
	Documents []Document
}

func (*MessagesStickerSetPredict) CRC() uint32 {
	return 0x6e153f16
}
func (*MessagesStickerSetPredict) _MessagesStickerSet() {}

type MessagesStickerSetNotModifiedPredict struct{}

func (*MessagesStickerSetNotModifiedPredict) CRC() uint32 {
	return 0xd3f924eb
}
func (*MessagesStickerSetNotModifiedPredict) _MessagesStickerSet() {}

type MessagesStickerSetInstallResult interface {
	tl.Object
	_MessagesStickerSetInstallResult()
}

var (
	_ MessagesStickerSetInstallResult = (*MessagesStickerSetInstallResultSuccessPredict)(nil)
	_ MessagesStickerSetInstallResult = (*MessagesStickerSetInstallResultArchivePredict)(nil)
)

type MessagesStickerSetInstallResultSuccessPredict struct{}

func (*MessagesStickerSetInstallResultSuccessPredict) CRC() uint32 {
	return 0x38641628
}
func (*MessagesStickerSetInstallResultSuccessPredict) _MessagesStickerSetInstallResult() {}

type MessagesStickerSetInstallResultArchivePredict struct {
	Sets []StickerSetCovered
}

func (*MessagesStickerSetInstallResultArchivePredict) CRC() uint32 {
	return 0x35e410a8
}
func (*MessagesStickerSetInstallResultArchivePredict) _MessagesStickerSetInstallResult() {}

type MessagesStickers interface {
	tl.Object
	_MessagesStickers()
}

var (
	_ MessagesStickers = (*MessagesStickersNotModifiedPredict)(nil)
	_ MessagesStickers = (*MessagesStickersPredict)(nil)
)

type MessagesStickersNotModifiedPredict struct{}

func (*MessagesStickersNotModifiedPredict) CRC() uint32 {
	return 0xf1749a22
}
func (*MessagesStickersNotModifiedPredict) _MessagesStickers() {}

type MessagesStickersPredict struct {
	Hash     int64
	Stickers []Document
}

func (*MessagesStickersPredict) CRC() uint32 {
	return 0x30a6ec7e
}
func (*MessagesStickersPredict) _MessagesStickers() {}

type MessagesTranscribedAudio interface {
	tl.Object
	_MessagesTranscribedAudio()
}

var (
	_ MessagesTranscribedAudio = (*MessagesTranscribedAudioPredict)(nil)
)

type MessagesTranscribedAudioPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	Pending               bool     `tl:",omitempty:flags:0,implicit"`
	TranscriptionID       int64
	Text                  string
	TrialRemainsNum       *int32 `tl:",omitempty:flags:1"`
	TrialRemainsUntilDate *int32 `tl:",omitempty:flags:1"`
}

func (*MessagesTranscribedAudioPredict) CRC() uint32 {
	return 0xcfb9d957
}
func (*MessagesTranscribedAudioPredict) _MessagesTranscribedAudio() {}

type MessagesTranslatedText interface {
	tl.Object
	_MessagesTranslatedText()
}

var (
	_ MessagesTranslatedText = (*MessagesTranslateResultPredict)(nil)
)

type MessagesTranslateResultPredict struct {
	Result []TextWithEntities
}

func (*MessagesTranslateResultPredict) CRC() uint32 {
	return 0x33db32f8
}
func (*MessagesTranslateResultPredict) _MessagesTranslatedText() {}

type MessagesVotesList interface {
	tl.Object
	_MessagesVotesList()
}

var (
	_ MessagesVotesList = (*MessagesVotesListPredict)(nil)
)

type MessagesVotesListPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Votes      []MessagePeerVote
	Chats      []Chat
	Users      []User
	NextOffset *string `tl:",omitempty:flags:0"`
}

func (*MessagesVotesListPredict) CRC() uint32 {
	return 0x4899484e
}
func (*MessagesVotesListPredict) _MessagesVotesList() {}

type MessagesWebPage interface {
	tl.Object
	_MessagesWebPage()
}

var (
	_ MessagesWebPage = (*MessagesWebPagePredict)(nil)
)

type MessagesWebPagePredict struct {
	Webpage WebPage
	Chats   []Chat
	Users   []User
}

func (*MessagesWebPagePredict) CRC() uint32 {
	return 0xfd5e12bd
}
func (*MessagesWebPagePredict) _MessagesWebPage() {}

type PaymentsBankCardData interface {
	tl.Object
	_PaymentsBankCardData()
}

var (
	_ PaymentsBankCardData = (*PaymentsBankCardDataPredict)(nil)
)

type PaymentsBankCardDataPredict struct {
	Title    string
	OpenUrls []BankCardOpenURL
}

func (*PaymentsBankCardDataPredict) CRC() uint32 {
	return 0x3e24e573
}
func (*PaymentsBankCardDataPredict) _PaymentsBankCardData() {}

type PaymentsCheckedGiftCode interface {
	tl.Object
	_PaymentsCheckedGiftCode()
}

var (
	_ PaymentsCheckedGiftCode = (*PaymentsCheckedGiftCodePredict)(nil)
)

type PaymentsCheckedGiftCodePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ViaGiveaway   bool     `tl:",omitempty:flags:2,implicit"`
	FromID        Peer     `tl:",omitempty:flags:4"`
	GiveawayMsgID *int32   `tl:",omitempty:flags:3"`
	ToID          *int64   `tl:",omitempty:flags:0"`
	Date          int32
	Months        int32
	UsedDate      *int32 `tl:",omitempty:flags:1"`
	Chats         []Chat
	Users         []User
}

func (*PaymentsCheckedGiftCodePredict) CRC() uint32 {
	return 0x284a1096
}
func (*PaymentsCheckedGiftCodePredict) _PaymentsCheckedGiftCode() {}

type PaymentsExportedInvoice interface {
	tl.Object
	_PaymentsExportedInvoice()
}

var (
	_ PaymentsExportedInvoice = (*PaymentsExportedInvoicePredict)(nil)
)

type PaymentsExportedInvoicePredict struct {
	URL string
}

func (*PaymentsExportedInvoicePredict) CRC() uint32 {
	return 0xaed0cbd9
}
func (*PaymentsExportedInvoicePredict) _PaymentsExportedInvoice() {}

type PaymentsGiveawayInfo interface {
	tl.Object
	_PaymentsGiveawayInfo()
}

var (
	_ PaymentsGiveawayInfo = (*PaymentsGiveawayInfoPredict)(nil)
	_ PaymentsGiveawayInfo = (*PaymentsGiveawayInfoResultsPredict)(nil)
)

type PaymentsGiveawayInfoPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	Participating         bool     `tl:",omitempty:flags:0,implicit"`
	PreparingResults      bool     `tl:",omitempty:flags:3,implicit"`
	StartDate             int32
	JoinedTooEarlyDate    *int32  `tl:",omitempty:flags:1"`
	AdminDisallowedChatID *int64  `tl:",omitempty:flags:2"`
	DisallowedCountry     *string `tl:",omitempty:flags:4"`
}

func (*PaymentsGiveawayInfoPredict) CRC() uint32 {
	return 0x4367daa0
}
func (*PaymentsGiveawayInfoPredict) _PaymentsGiveawayInfo() {}

type PaymentsGiveawayInfoResultsPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Winner         bool     `tl:",omitempty:flags:0,implicit"`
	Refunded       bool     `tl:",omitempty:flags:1,implicit"`
	StartDate      int32
	GiftCodeSlug   *string `tl:",omitempty:flags:0"`
	FinishDate     int32
	WinnersCount   int32
	ActivatedCount int32
}

func (*PaymentsGiveawayInfoResultsPredict) CRC() uint32 {
	return 0xcd5570
}
func (*PaymentsGiveawayInfoResultsPredict) _PaymentsGiveawayInfo() {}

type PaymentsPaymentForm interface {
	tl.Object
	_PaymentsPaymentForm()
}

var (
	_ PaymentsPaymentForm = (*PaymentsPaymentFormPredict)(nil)
	_ PaymentsPaymentForm = (*PaymentsPaymentFormStarsPredict)(nil)
)

type PaymentsPaymentFormPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	CanSaveCredentials bool     `tl:",omitempty:flags:2,implicit"`
	PasswordMissing    bool     `tl:",omitempty:flags:3,implicit"`
	FormID             int64
	BotID              int64
	Title              string
	Description        string
	Photo              WebDocument `tl:",omitempty:flags:5"`
	Invoice            Invoice
	ProviderID         int64
	URL                string
	NativeProvider     *string                   `tl:",omitempty:flags:4"`
	NativeParams       DataJSON                  `tl:",omitempty:flags:4"`
	AdditionalMethods  []PaymentFormMethod       `tl:",omitempty:flags:6"`
	SavedInfo          PaymentRequestedInfo      `tl:",omitempty:flags:0"`
	SavedCredentials   []PaymentSavedCredentials `tl:",omitempty:flags:1"`
	Users              []User
}

func (*PaymentsPaymentFormPredict) CRC() uint32 {
	return 0xa0058751
}
func (*PaymentsPaymentFormPredict) _PaymentsPaymentForm() {}

type PaymentsPaymentFormStarsPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	FormID      int64
	BotID       int64
	Title       string
	Description string
	Photo       WebDocument `tl:",omitempty:flags:5"`
	Invoice     Invoice
	Users       []User
}

func (*PaymentsPaymentFormStarsPredict) CRC() uint32 {
	return 0x7bf6b15c
}
func (*PaymentsPaymentFormStarsPredict) _PaymentsPaymentForm() {}

type PaymentsPaymentReceipt interface {
	tl.Object
	_PaymentsPaymentReceipt()
}

var (
	_ PaymentsPaymentReceipt = (*PaymentsPaymentReceiptPredict)(nil)
	_ PaymentsPaymentReceipt = (*PaymentsPaymentReceiptStarsPredict)(nil)
)

type PaymentsPaymentReceiptPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	Date             int32
	BotID            int64
	ProviderID       int64
	Title            string
	Description      string
	Photo            WebDocument `tl:",omitempty:flags:2"`
	Invoice          Invoice
	Info             PaymentRequestedInfo `tl:",omitempty:flags:0"`
	Shipping         ShippingOption       `tl:",omitempty:flags:1"`
	TipAmount        *int64               `tl:",omitempty:flags:3"`
	Currency         string
	TotalAmount      int64
	CredentialsTitle string
	Users            []User
}

func (*PaymentsPaymentReceiptPredict) CRC() uint32 {
	return 0x70c4fe03
}
func (*PaymentsPaymentReceiptPredict) _PaymentsPaymentReceipt() {}

type PaymentsPaymentReceiptStarsPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Date          int32
	BotID         int64
	Title         string
	Description   string
	Photo         WebDocument `tl:",omitempty:flags:2"`
	Invoice       Invoice
	Currency      string
	TotalAmount   int64
	TransactionID string
	Users         []User
}

func (*PaymentsPaymentReceiptStarsPredict) CRC() uint32 {
	return 0xdabbf83a
}
func (*PaymentsPaymentReceiptStarsPredict) _PaymentsPaymentReceipt() {}

type PaymentsPaymentResult interface {
	tl.Object
	_PaymentsPaymentResult()
}

var (
	_ PaymentsPaymentResult = (*PaymentsPaymentResultPredict)(nil)
	_ PaymentsPaymentResult = (*PaymentsPaymentVerificationNeededPredict)(nil)
)

type PaymentsPaymentResultPredict struct {
	Updates Updates
}

func (*PaymentsPaymentResultPredict) CRC() uint32 {
	return 0x4e5f810d
}
func (*PaymentsPaymentResultPredict) _PaymentsPaymentResult() {}

type PaymentsPaymentVerificationNeededPredict struct {
	URL string
}

func (*PaymentsPaymentVerificationNeededPredict) CRC() uint32 {
	return 0xd8411139
}
func (*PaymentsPaymentVerificationNeededPredict) _PaymentsPaymentResult() {}

type PaymentsSavedInfo interface {
	tl.Object
	_PaymentsSavedInfo()
}

var (
	_ PaymentsSavedInfo = (*PaymentsSavedInfoPredict)(nil)
)

type PaymentsSavedInfoPredict struct {
	_                   struct{}             `tl:"flags,bitflag"`
	HasSavedCredentials bool                 `tl:",omitempty:flags:1,implicit"`
	SavedInfo           PaymentRequestedInfo `tl:",omitempty:flags:0"`
}

func (*PaymentsSavedInfoPredict) CRC() uint32 {
	return 0xfb8fe43c
}
func (*PaymentsSavedInfoPredict) _PaymentsSavedInfo() {}

type PaymentsStarsRevenueAdsAccountURL interface {
	tl.Object
	_PaymentsStarsRevenueAdsAccountURL()
}

var (
	_ PaymentsStarsRevenueAdsAccountURL = (*PaymentsStarsRevenueAdsAccountURLPredict)(nil)
)

type PaymentsStarsRevenueAdsAccountURLPredict struct {
	URL string
}

func (*PaymentsStarsRevenueAdsAccountURLPredict) CRC() uint32 {
	return 0x394e7f21
}
func (*PaymentsStarsRevenueAdsAccountURLPredict) _PaymentsStarsRevenueAdsAccountURL() {}

type PaymentsStarsRevenueStats interface {
	tl.Object
	_PaymentsStarsRevenueStats()
}

var (
	_ PaymentsStarsRevenueStats = (*PaymentsStarsRevenueStatsPredict)(nil)
)

type PaymentsStarsRevenueStatsPredict struct {
	RevenueGraph StatsGraph
	Status       StarsRevenueStatus
	UsdRate      float64
}

func (*PaymentsStarsRevenueStatsPredict) CRC() uint32 {
	return 0xc92bb73b
}
func (*PaymentsStarsRevenueStatsPredict) _PaymentsStarsRevenueStats() {}

type PaymentsStarsRevenueWithdrawalURL interface {
	tl.Object
	_PaymentsStarsRevenueWithdrawalURL()
}

var (
	_ PaymentsStarsRevenueWithdrawalURL = (*PaymentsStarsRevenueWithdrawalURLPredict)(nil)
)

type PaymentsStarsRevenueWithdrawalURLPredict struct {
	URL string
}

func (*PaymentsStarsRevenueWithdrawalURLPredict) CRC() uint32 {
	return 0x1dab80b7
}
func (*PaymentsStarsRevenueWithdrawalURLPredict) _PaymentsStarsRevenueWithdrawalURL() {}

type PaymentsStarsStatus interface {
	tl.Object
	_PaymentsStarsStatus()
}

var (
	_ PaymentsStarsStatus = (*PaymentsStarsStatusPredict)(nil)
)

type PaymentsStarsStatusPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Balance    int64
	History    []StarsTransaction
	NextOffset *string `tl:",omitempty:flags:0"`
	Chats      []Chat
	Users      []User
}

func (*PaymentsStarsStatusPredict) CRC() uint32 {
	return 0x8cf4ee60
}
func (*PaymentsStarsStatusPredict) _PaymentsStarsStatus() {}

type PaymentsValidatedRequestedInfo interface {
	tl.Object
	_PaymentsValidatedRequestedInfo()
}

var (
	_ PaymentsValidatedRequestedInfo = (*PaymentsValidatedRequestedInfoPredict)(nil)
)

type PaymentsValidatedRequestedInfoPredict struct {
	_               struct{}         `tl:"flags,bitflag"`
	ID              *string          `tl:",omitempty:flags:0"`
	ShippingOptions []ShippingOption `tl:",omitempty:flags:1"`
}

func (*PaymentsValidatedRequestedInfoPredict) CRC() uint32 {
	return 0xd1451883
}
func (*PaymentsValidatedRequestedInfoPredict) _PaymentsValidatedRequestedInfo() {}

type PhoneExportedGroupCallInvite interface {
	tl.Object
	_PhoneExportedGroupCallInvite()
}

var (
	_ PhoneExportedGroupCallInvite = (*PhoneExportedGroupCallInvitePredict)(nil)
)

type PhoneExportedGroupCallInvitePredict struct {
	Link string
}

func (*PhoneExportedGroupCallInvitePredict) CRC() uint32 {
	return 0x204bd158
}
func (*PhoneExportedGroupCallInvitePredict) _PhoneExportedGroupCallInvite() {}

type PhoneGroupCall interface {
	tl.Object
	_PhoneGroupCall()
}

var (
	_ PhoneGroupCall = (*PhoneGroupCallPredict)(nil)
)

type PhoneGroupCallPredict struct {
	Call                   GroupCall
	Participants           []GroupCallParticipant
	ParticipantsNextOffset string
	Chats                  []Chat
	Users                  []User
}

func (*PhoneGroupCallPredict) CRC() uint32 {
	return 0x9e727aad
}
func (*PhoneGroupCallPredict) _PhoneGroupCall() {}

type PhoneGroupCallStreamChannels interface {
	tl.Object
	_PhoneGroupCallStreamChannels()
}

var (
	_ PhoneGroupCallStreamChannels = (*PhoneGroupCallStreamChannelsPredict)(nil)
)

type PhoneGroupCallStreamChannelsPredict struct {
	Channels []GroupCallStreamChannel
}

func (*PhoneGroupCallStreamChannelsPredict) CRC() uint32 {
	return 0xd0e482b2
}
func (*PhoneGroupCallStreamChannelsPredict) _PhoneGroupCallStreamChannels() {}

type PhoneGroupCallStreamRtmpURL interface {
	tl.Object
	_PhoneGroupCallStreamRtmpURL()
}

var (
	_ PhoneGroupCallStreamRtmpURL = (*PhoneGroupCallStreamRtmpURLPredict)(nil)
)

type PhoneGroupCallStreamRtmpURLPredict struct {
	URL string
	Key string
}

func (*PhoneGroupCallStreamRtmpURLPredict) CRC() uint32 {
	return 0x2dbf3432
}
func (*PhoneGroupCallStreamRtmpURLPredict) _PhoneGroupCallStreamRtmpURL() {}

type PhoneGroupParticipants interface {
	tl.Object
	_PhoneGroupParticipants()
}

var (
	_ PhoneGroupParticipants = (*PhoneGroupParticipantsPredict)(nil)
)

type PhoneGroupParticipantsPredict struct {
	Count        int32
	Participants []GroupCallParticipant
	NextOffset   string
	Chats        []Chat
	Users        []User
	Version      int32
}

func (*PhoneGroupParticipantsPredict) CRC() uint32 {
	return 0xf47751b6
}
func (*PhoneGroupParticipantsPredict) _PhoneGroupParticipants() {}

type PhoneJoinAsPeers interface {
	tl.Object
	_PhoneJoinAsPeers()
}

var (
	_ PhoneJoinAsPeers = (*PhoneJoinAsPeersPredict)(nil)
)

type PhoneJoinAsPeersPredict struct {
	Peers []Peer
	Chats []Chat
	Users []User
}

func (*PhoneJoinAsPeersPredict) CRC() uint32 {
	return 0xafe5623f
}
func (*PhoneJoinAsPeersPredict) _PhoneJoinAsPeers() {}

type PhonePhoneCall interface {
	tl.Object
	_PhonePhoneCall()
}

var (
	_ PhonePhoneCall = (*PhonePhoneCallPredict)(nil)
)

type PhonePhoneCallPredict struct {
	PhoneCall PhoneCall
	Users     []User
}

func (*PhonePhoneCallPredict) CRC() uint32 {
	return 0xec82e140
}
func (*PhonePhoneCallPredict) _PhonePhoneCall() {}

type PhotosPhoto interface {
	tl.Object
	_PhotosPhoto()
}

var (
	_ PhotosPhoto = (*PhotosPhotoPredict)(nil)
)

type PhotosPhotoPredict struct {
	Photo Photo
	Users []User
}

func (*PhotosPhotoPredict) CRC() uint32 {
	return 0x20212ca8
}
func (*PhotosPhotoPredict) _PhotosPhoto() {}

type PhotosPhotos interface {
	tl.Object
	_PhotosPhotos()
}

var (
	_ PhotosPhotos = (*PhotosPhotosPredict)(nil)
	_ PhotosPhotos = (*PhotosPhotosSlicePredict)(nil)
)

type PhotosPhotosPredict struct {
	Photos []Photo
	Users  []User
}

func (*PhotosPhotosPredict) CRC() uint32 {
	return 0x8dca6aa5
}
func (*PhotosPhotosPredict) _PhotosPhotos() {}

type PhotosPhotosSlicePredict struct {
	Count  int32
	Photos []Photo
	Users  []User
}

func (*PhotosPhotosSlicePredict) CRC() uint32 {
	return 0x15051f54
}
func (*PhotosPhotosSlicePredict) _PhotosPhotos() {}

type PremiumBoostsList interface {
	tl.Object
	_PremiumBoostsList()
}

var (
	_ PremiumBoostsList = (*PremiumBoostsListPredict)(nil)
)

type PremiumBoostsListPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Boosts     []Boost
	NextOffset *string `tl:",omitempty:flags:0"`
	Users      []User
}

func (*PremiumBoostsListPredict) CRC() uint32 {
	return 0x86f8613c
}
func (*PremiumBoostsListPredict) _PremiumBoostsList() {}

type PremiumBoostsStatus interface {
	tl.Object
	_PremiumBoostsStatus()
}

var (
	_ PremiumBoostsStatus = (*PremiumBoostsStatusPredict)(nil)
)

type PremiumBoostsStatusPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	MyBoost            bool     `tl:",omitempty:flags:2,implicit"`
	Level              int32
	CurrentLevelBoosts int32
	Boosts             int32
	GiftBoosts         *int32            `tl:",omitempty:flags:4"`
	NextLevelBoosts    *int32            `tl:",omitempty:flags:0"`
	PremiumAudience    StatsPercentValue `tl:",omitempty:flags:1"`
	BoostURL           string
	PrepaidGiveaways   []PrepaidGiveaway `tl:",omitempty:flags:3"`
	MyBoostSlots       []int32           `tl:",omitempty:flags:2"`
}

func (*PremiumBoostsStatusPredict) CRC() uint32 {
	return 0x4959427a
}
func (*PremiumBoostsStatusPredict) _PremiumBoostsStatus() {}

type PremiumMyBoosts interface {
	tl.Object
	_PremiumMyBoosts()
}

var (
	_ PremiumMyBoosts = (*PremiumMyBoostsPredict)(nil)
)

type PremiumMyBoostsPredict struct {
	MyBoosts []MyBoost
	Chats    []Chat
	Users    []User
}

func (*PremiumMyBoostsPredict) CRC() uint32 {
	return 0x9ae228e2
}
func (*PremiumMyBoostsPredict) _PremiumMyBoosts() {}

type SmsjobsEligibilityToJoin interface {
	tl.Object
	_SmsjobsEligibilityToJoin()
}

var (
	_ SmsjobsEligibilityToJoin = (*SmsjobsEligibleToJoinPredict)(nil)
)

type SmsjobsEligibleToJoinPredict struct {
	TermsURL       string
	MonthlySentSms int32
}

func (*SmsjobsEligibleToJoinPredict) CRC() uint32 {
	return 0xdc8b44cf
}
func (*SmsjobsEligibleToJoinPredict) _SmsjobsEligibilityToJoin() {}

type SmsjobsStatus interface {
	tl.Object
	_SmsjobsStatus()
}

var (
	_ SmsjobsStatus = (*SmsjobsStatusPredict)(nil)
)

type SmsjobsStatusPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	AllowInternational bool     `tl:",omitempty:flags:0,implicit"`
	RecentSent         int32
	RecentSince        int32
	RecentRemains      int32
	TotalSent          int32
	TotalSince         int32
	LastGiftSlug       *string `tl:",omitempty:flags:1"`
	TermsURL           string
}

func (*SmsjobsStatusPredict) CRC() uint32 {
	return 0x2aee9191
}
func (*SmsjobsStatusPredict) _SmsjobsStatus() {}

type StatsBroadcastRevenueStats interface {
	tl.Object
	_StatsBroadcastRevenueStats()
}

var (
	_ StatsBroadcastRevenueStats = (*StatsBroadcastRevenueStatsPredict)(nil)
)

type StatsBroadcastRevenueStatsPredict struct {
	TopHoursGraph StatsGraph
	RevenueGraph  StatsGraph
	Balances      BroadcastRevenueBalances
	UsdRate       float64
}

func (*StatsBroadcastRevenueStatsPredict) CRC() uint32 {
	return 0x5407e297
}
func (*StatsBroadcastRevenueStatsPredict) _StatsBroadcastRevenueStats() {}

type StatsBroadcastRevenueTransactions interface {
	tl.Object
	_StatsBroadcastRevenueTransactions()
}

var (
	_ StatsBroadcastRevenueTransactions = (*StatsBroadcastRevenueTransactionsPredict)(nil)
)

type StatsBroadcastRevenueTransactionsPredict struct {
	Count        int32
	Transactions []BroadcastRevenueTransaction
}

func (*StatsBroadcastRevenueTransactionsPredict) CRC() uint32 {
	return 0x87158466
}
func (*StatsBroadcastRevenueTransactionsPredict) _StatsBroadcastRevenueTransactions() {}

type StatsBroadcastRevenueWithdrawalURL interface {
	tl.Object
	_StatsBroadcastRevenueWithdrawalURL()
}

var (
	_ StatsBroadcastRevenueWithdrawalURL = (*StatsBroadcastRevenueWithdrawalURLPredict)(nil)
)

type StatsBroadcastRevenueWithdrawalURLPredict struct {
	URL string
}

func (*StatsBroadcastRevenueWithdrawalURLPredict) CRC() uint32 {
	return 0xec659737
}
func (*StatsBroadcastRevenueWithdrawalURLPredict) _StatsBroadcastRevenueWithdrawalURL() {}

type StatsBroadcastStats interface {
	tl.Object
	_StatsBroadcastStats()
}

var (
	_ StatsBroadcastStats = (*StatsBroadcastStatsPredict)(nil)
)

type StatsBroadcastStatsPredict struct {
	Period                       StatsDateRangeDays
	Followers                    StatsAbsValueAndPrev
	ViewsPerPost                 StatsAbsValueAndPrev
	SharesPerPost                StatsAbsValueAndPrev
	ReactionsPerPost             StatsAbsValueAndPrev
	ViewsPerStory                StatsAbsValueAndPrev
	SharesPerStory               StatsAbsValueAndPrev
	ReactionsPerStory            StatsAbsValueAndPrev
	EnabledNotifications         StatsPercentValue
	GrowthGraph                  StatsGraph
	FollowersGraph               StatsGraph
	MuteGraph                    StatsGraph
	TopHoursGraph                StatsGraph
	InteractionsGraph            StatsGraph
	IvInteractionsGraph          StatsGraph
	ViewsBySourceGraph           StatsGraph
	NewFollowersBySourceGraph    StatsGraph
	LanguagesGraph               StatsGraph
	ReactionsByEmotionGraph      StatsGraph
	StoryInteractionsGraph       StatsGraph
	StoryReactionsByEmotionGraph StatsGraph
	RecentPostsInteractions      []PostInteractionCounters
}

func (*StatsBroadcastStatsPredict) CRC() uint32 {
	return 0x396ca5fc
}
func (*StatsBroadcastStatsPredict) _StatsBroadcastStats() {}

type StatsMegagroupStats interface {
	tl.Object
	_StatsMegagroupStats()
}

var (
	_ StatsMegagroupStats = (*StatsMegagroupStatsPredict)(nil)
)

type StatsMegagroupStatsPredict struct {
	Period                  StatsDateRangeDays
	Members                 StatsAbsValueAndPrev
	Messages                StatsAbsValueAndPrev
	Viewers                 StatsAbsValueAndPrev
	Posters                 StatsAbsValueAndPrev
	GrowthGraph             StatsGraph
	MembersGraph            StatsGraph
	NewMembersBySourceGraph StatsGraph
	LanguagesGraph          StatsGraph
	MessagesGraph           StatsGraph
	ActionsGraph            StatsGraph
	TopHoursGraph           StatsGraph
	WeekdaysGraph           StatsGraph
	TopPosters              []StatsGroupTopPoster
	TopAdmins               []StatsGroupTopAdmin
	TopInviters             []StatsGroupTopInviter
	Users                   []User
}

func (*StatsMegagroupStatsPredict) CRC() uint32 {
	return 0xef7ff916
}
func (*StatsMegagroupStatsPredict) _StatsMegagroupStats() {}

type StatsMessageStats interface {
	tl.Object
	_StatsMessageStats()
}

var (
	_ StatsMessageStats = (*StatsMessageStatsPredict)(nil)
)

type StatsMessageStatsPredict struct {
	ViewsGraph              StatsGraph
	ReactionsByEmotionGraph StatsGraph
}

func (*StatsMessageStatsPredict) CRC() uint32 {
	return 0x7fe91c14
}
func (*StatsMessageStatsPredict) _StatsMessageStats() {}

type StatsPublicForwards interface {
	tl.Object
	_StatsPublicForwards()
}

var (
	_ StatsPublicForwards = (*StatsPublicForwardsPredict)(nil)
)

type StatsPublicForwardsPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Forwards   []PublicForward
	NextOffset *string `tl:",omitempty:flags:0"`
	Chats      []Chat
	Users      []User
}

func (*StatsPublicForwardsPredict) CRC() uint32 {
	return 0x93037e20
}
func (*StatsPublicForwardsPredict) _StatsPublicForwards() {}

type StatsStoryStats interface {
	tl.Object
	_StatsStoryStats()
}

var (
	_ StatsStoryStats = (*StatsStoryStatsPredict)(nil)
)

type StatsStoryStatsPredict struct {
	ViewsGraph              StatsGraph
	ReactionsByEmotionGraph StatsGraph
}

func (*StatsStoryStatsPredict) CRC() uint32 {
	return 0x50cd067c
}
func (*StatsStoryStatsPredict) _StatsStoryStats() {}

type StickersSuggestedShortName interface {
	tl.Object
	_StickersSuggestedShortName()
}

var (
	_ StickersSuggestedShortName = (*StickersSuggestedShortNamePredict)(nil)
)

type StickersSuggestedShortNamePredict struct {
	ShortName string
}

func (*StickersSuggestedShortNamePredict) CRC() uint32 {
	return 0x85fea03f
}
func (*StickersSuggestedShortNamePredict) _StickersSuggestedShortName() {}

type StoriesAllStories interface {
	tl.Object
	_StoriesAllStories()
}

var (
	_ StoriesAllStories = (*StoriesAllStoriesNotModifiedPredict)(nil)
	_ StoriesAllStories = (*StoriesAllStoriesPredict)(nil)
)

type StoriesAllStoriesNotModifiedPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	State       string
	StealthMode StoriesStealthMode
}

func (*StoriesAllStoriesNotModifiedPredict) CRC() uint32 {
	return 0x1158fe3e
}
func (*StoriesAllStoriesNotModifiedPredict) _StoriesAllStories() {}

type StoriesAllStoriesPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	HasMore     bool     `tl:",omitempty:flags:0,implicit"`
	Count       int32
	State       string
	PeerStories []PeerStories
	Chats       []Chat
	Users       []User
	StealthMode StoriesStealthMode
}

func (*StoriesAllStoriesPredict) CRC() uint32 {
	return 0x6efc5e81
}
func (*StoriesAllStoriesPredict) _StoriesAllStories() {}

type StoriesFoundStories interface {
	tl.Object
	_StoriesFoundStories()
}

var (
	_ StoriesFoundStories = (*StoriesFoundStoriesPredict)(nil)
)

type StoriesFoundStoriesPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Stories    []FoundStory
	NextOffset *string `tl:",omitempty:flags:0"`
	Chats      []Chat
	Users      []User
}

func (*StoriesFoundStoriesPredict) CRC() uint32 {
	return 0xe2de7737
}
func (*StoriesFoundStoriesPredict) _StoriesFoundStories() {}

type StoriesPeerStories interface {
	tl.Object
	_StoriesPeerStories()
}

var (
	_ StoriesPeerStories = (*StoriesPeerStoriesPredict)(nil)
)

type StoriesPeerStoriesPredict struct {
	Stories PeerStories
	Chats   []Chat
	Users   []User
}

func (*StoriesPeerStoriesPredict) CRC() uint32 {
	return 0xcae68768
}
func (*StoriesPeerStoriesPredict) _StoriesPeerStories() {}

type StoriesStories interface {
	tl.Object
	_StoriesStories()
}

var (
	_ StoriesStories = (*StoriesStoriesPredict)(nil)
)

type StoriesStoriesPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Count       int32
	Stories     []StoryItem
	PinnedToTop []int32 `tl:",omitempty:flags:0"`
	Chats       []Chat
	Users       []User
}

func (*StoriesStoriesPredict) CRC() uint32 {
	return 0x63c3dd0a
}
func (*StoriesStoriesPredict) _StoriesStories() {}

type StoriesStoryReactionsList interface {
	tl.Object
	_StoriesStoryReactionsList()
}

var (
	_ StoriesStoryReactionsList = (*StoriesStoryReactionsListPredict)(nil)
)

type StoriesStoryReactionsListPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32
	Reactions  []StoryReaction
	Chats      []Chat
	Users      []User
	NextOffset *string `tl:",omitempty:flags:0"`
}

func (*StoriesStoryReactionsListPredict) CRC() uint32 {
	return 0xaa5f789c
}
func (*StoriesStoryReactionsListPredict) _StoriesStoryReactionsList() {}

type StoriesStoryViews interface {
	tl.Object
	_StoriesStoryViews()
}

var (
	_ StoriesStoryViews = (*StoriesStoryViewsPredict)(nil)
)

type StoriesStoryViewsPredict struct {
	Views []StoryViews
	Users []User
}

func (*StoriesStoryViewsPredict) CRC() uint32 {
	return 0xde9eed1d
}
func (*StoriesStoryViewsPredict) _StoriesStoryViews() {}

type StoriesStoryViewsList interface {
	tl.Object
	_StoriesStoryViewsList()
}

var (
	_ StoriesStoryViewsList = (*StoriesStoryViewsListPredict)(nil)
)

type StoriesStoryViewsListPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Count          int32
	ViewsCount     int32
	ForwardsCount  int32
	ReactionsCount int32
	Views          []StoryView
	Chats          []Chat
	Users          []User
	NextOffset     *string `tl:",omitempty:flags:0"`
}

func (*StoriesStoryViewsListPredict) CRC() uint32 {
	return 0x59d78fc5
}
func (*StoriesStoryViewsListPredict) _StoriesStoryViewsList() {}

type UpdatesChannelDifference interface {
	tl.Object
	_UpdatesChannelDifference()
}

var (
	_ UpdatesChannelDifference = (*UpdatesChannelDifferenceEmptyPredict)(nil)
	_ UpdatesChannelDifference = (*UpdatesChannelDifferenceTooLongPredict)(nil)
	_ UpdatesChannelDifference = (*UpdatesChannelDifferencePredict)(nil)
)

type UpdatesChannelDifferenceEmptyPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Final   bool     `tl:",omitempty:flags:0,implicit"`
	Pts     int32
	Timeout *int32 `tl:",omitempty:flags:1"`
}

func (*UpdatesChannelDifferenceEmptyPredict) CRC() uint32 {
	return 0x3e11affb
}
func (*UpdatesChannelDifferenceEmptyPredict) _UpdatesChannelDifference() {}

type UpdatesChannelDifferenceTooLongPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Final    bool     `tl:",omitempty:flags:0,implicit"`
	Timeout  *int32   `tl:",omitempty:flags:1"`
	Dialog   Dialog
	Messages []Message
	Chats    []Chat
	Users    []User
}

func (*UpdatesChannelDifferenceTooLongPredict) CRC() uint32 {
	return 0xa4bcc6fe
}
func (*UpdatesChannelDifferenceTooLongPredict) _UpdatesChannelDifference() {}

type UpdatesChannelDifferencePredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Final        bool     `tl:",omitempty:flags:0,implicit"`
	Pts          int32
	Timeout      *int32 `tl:",omitempty:flags:1"`
	NewMessages  []Message
	OtherUpdates []Update
	Chats        []Chat
	Users        []User
}

func (*UpdatesChannelDifferencePredict) CRC() uint32 {
	return 0x2064674e
}
func (*UpdatesChannelDifferencePredict) _UpdatesChannelDifference() {}

type UpdatesDifference interface {
	tl.Object
	_UpdatesDifference()
}

var (
	_ UpdatesDifference = (*UpdatesDifferenceEmptyPredict)(nil)
	_ UpdatesDifference = (*UpdatesDifferencePredict)(nil)
	_ UpdatesDifference = (*UpdatesDifferenceSlicePredict)(nil)
	_ UpdatesDifference = (*UpdatesDifferenceTooLongPredict)(nil)
)

type UpdatesDifferenceEmptyPredict struct {
	Date int32
	Seq  int32
}

func (*UpdatesDifferenceEmptyPredict) CRC() uint32 {
	return 0x5d75a138
}
func (*UpdatesDifferenceEmptyPredict) _UpdatesDifference() {}

type UpdatesDifferencePredict struct {
	NewMessages          []Message
	NewEncryptedMessages []EncryptedMessage
	OtherUpdates         []Update
	Chats                []Chat
	Users                []User
	State                UpdatesState
}

func (*UpdatesDifferencePredict) CRC() uint32 {
	return 0xf49ca0
}
func (*UpdatesDifferencePredict) _UpdatesDifference() {}

type UpdatesDifferenceSlicePredict struct {
	NewMessages          []Message
	NewEncryptedMessages []EncryptedMessage
	OtherUpdates         []Update
	Chats                []Chat
	Users                []User
	IntermediateState    UpdatesState
}

func (*UpdatesDifferenceSlicePredict) CRC() uint32 {
	return 0xa8fb1981
}
func (*UpdatesDifferenceSlicePredict) _UpdatesDifference() {}

type UpdatesDifferenceTooLongPredict struct {
	Pts int32
}

func (*UpdatesDifferenceTooLongPredict) CRC() uint32 {
	return 0x4afe8f6d
}
func (*UpdatesDifferenceTooLongPredict) _UpdatesDifference() {}

type UpdatesState interface {
	tl.Object
	_UpdatesState()
}

var (
	_ UpdatesState = (*UpdatesStatePredict)(nil)
)

type UpdatesStatePredict struct {
	Pts         int32
	Qts         int32
	Date        int32
	Seq         int32
	UnreadCount int32
}

func (*UpdatesStatePredict) CRC() uint32 {
	return 0xa56c2a3e
}
func (*UpdatesStatePredict) _UpdatesState() {}

type UploadCdnFile interface {
	tl.Object
	_UploadCdnFile()
}

var (
	_ UploadCdnFile = (*UploadCdnFileReuploadNeededPredict)(nil)
	_ UploadCdnFile = (*UploadCdnFilePredict)(nil)
)

type UploadCdnFileReuploadNeededPredict struct {
	RequestToken []byte
}

func (*UploadCdnFileReuploadNeededPredict) CRC() uint32 {
	return 0xeea8e46e
}
func (*UploadCdnFileReuploadNeededPredict) _UploadCdnFile() {}

type UploadCdnFilePredict struct {
	Bytes []byte
}

func (*UploadCdnFilePredict) CRC() uint32 {
	return 0xa99fca4f
}
func (*UploadCdnFilePredict) _UploadCdnFile() {}

type UploadFile interface {
	tl.Object
	_UploadFile()
}

var (
	_ UploadFile = (*UploadFilePredict)(nil)
	_ UploadFile = (*UploadFileCdnRedirectPredict)(nil)
)

type UploadFilePredict struct {
	Type  StorageFileType
	Mtime int32
	Bytes []byte
}

func (*UploadFilePredict) CRC() uint32 {
	return 0x96a18d5
}
func (*UploadFilePredict) _UploadFile() {}

type UploadFileCdnRedirectPredict struct {
	DcID          int32
	FileToken     []byte
	EncryptionKey []byte
	EncryptionIv  []byte
	FileHashes    []FileHash
}

func (*UploadFileCdnRedirectPredict) CRC() uint32 {
	return 0xf18cda44
}
func (*UploadFileCdnRedirectPredict) _UploadFile() {}

type UploadWebFile interface {
	tl.Object
	_UploadWebFile()
}

var (
	_ UploadWebFile = (*UploadWebFilePredict)(nil)
)

type UploadWebFilePredict struct {
	Size     int32
	MimeType string
	FileType StorageFileType
	Mtime    int32
	Bytes    []byte
}

func (*UploadWebFilePredict) CRC() uint32 {
	return 0x21e753bc
}
func (*UploadWebFilePredict) _UploadWebFile() {}

type UsersUserFull interface {
	tl.Object
	_UsersUserFull()
}

var (
	_ UsersUserFull = (*UsersUserFullPredict)(nil)
)

type UsersUserFullPredict struct {
	FullUser UserFull
	Chats    []Chat
	Users    []User
}

func (*UsersUserFullPredict) CRC() uint32 {
	return 0x3b6d152e
}
func (*UsersUserFullPredict) _UsersUserFull() {}

type AttachMenuPeerType uint32

const (
	AttachMenuPeerTypeSameBotPmPredict AttachMenuPeerType = 0x7d6be90e
	AttachMenuPeerTypeBotPmPredict     AttachMenuPeerType = 0xc32bfa1a
	AttachMenuPeerTypePmPredict        AttachMenuPeerType = 0xf146d31f
	AttachMenuPeerTypeChatPredict      AttachMenuPeerType = 0x509113f
	AttachMenuPeerTypeBroadcastPredict AttachMenuPeerType = 0x7bfbdefc
)

type BaseTheme uint32

const (
	BaseThemeClassicPredict BaseTheme = 0xc3a12462
	BaseThemeDayPredict     BaseTheme = 0xfbd81688
	BaseThemeNightPredict   BaseTheme = 0xb7b31ea8
	BaseThemeTintedPredict  BaseTheme = 0x6d5f77ee
	BaseThemeArcticPredict  BaseTheme = 0x5b11125a
)

type Bool uint32

const (
	BoolFalsePredict Bool = 0xbc799737
	BoolTruePredict  Bool = 0x997275b5
)

type InlineQueryPeerType uint32

const (
	InlineQueryPeerTypeSameBotPmPredict InlineQueryPeerType = 0x3081ed9d
	InlineQueryPeerTypePmPredict        InlineQueryPeerType = 0x833c0fac
	InlineQueryPeerTypeChatPredict      InlineQueryPeerType = 0xd766c50a
	InlineQueryPeerTypeMegagroupPredict InlineQueryPeerType = 0x5ec4be43
	InlineQueryPeerTypeBroadcastPredict InlineQueryPeerType = 0x6334ee9a
	InlineQueryPeerTypeBotPmPredict     InlineQueryPeerType = 0xe3b2d0c
)

type InputPrivacyKey uint32

const (
	InputPrivacyKeyStatusTimestampPredict InputPrivacyKey = 0x4f96cb18
	InputPrivacyKeyChatInvitePredict      InputPrivacyKey = 0xbdfb0426
	InputPrivacyKeyPhoneCallPredict       InputPrivacyKey = 0xfabadc5f
	InputPrivacyKeyPhoneP2PPredict        InputPrivacyKey = 0xdb9e70d2
	InputPrivacyKeyForwardsPredict        InputPrivacyKey = 0xa4dd4c08
	InputPrivacyKeyProfilePhotoPredict    InputPrivacyKey = 0x5719bacc
	InputPrivacyKeyPhoneNumberPredict     InputPrivacyKey = 0x352dafa
	InputPrivacyKeyAddedByPhonePredict    InputPrivacyKey = 0xd1219bdd
	InputPrivacyKeyVoiceMessagesPredict   InputPrivacyKey = 0xaee69d68
	InputPrivacyKeyAboutPredict           InputPrivacyKey = 0x3823cc40
	InputPrivacyKeyBirthdayPredict        InputPrivacyKey = 0xd65a11cc
)

type Null uint32

const (
	NullPredict Null = 0x56730bcc
)

type PhoneCallDiscardReason uint32

const (
	PhoneCallDiscardReasonMissedPredict     PhoneCallDiscardReason = 0x85e42301
	PhoneCallDiscardReasonDisconnectPredict PhoneCallDiscardReason = 0xe095c1a0
	PhoneCallDiscardReasonHangupPredict     PhoneCallDiscardReason = 0x57adc690
	PhoneCallDiscardReasonBusyPredict       PhoneCallDiscardReason = 0xfaf7e8c9
)

type PrivacyKey uint32

const (
	PrivacyKeyStatusTimestampPredict PrivacyKey = 0xbc2eab30
	PrivacyKeyChatInvitePredict      PrivacyKey = 0x500e6dfa
	PrivacyKeyPhoneCallPredict       PrivacyKey = 0x3d662b7b
	PrivacyKeyPhoneP2PPredict        PrivacyKey = 0x39491cc8
	PrivacyKeyForwardsPredict        PrivacyKey = 0x69ec56a3
	PrivacyKeyProfilePhotoPredict    PrivacyKey = 0x96151fed
	PrivacyKeyPhoneNumberPredict     PrivacyKey = 0xd19ae46d
	PrivacyKeyAddedByPhonePredict    PrivacyKey = 0x42ffd42b
	PrivacyKeyVoiceMessagesPredict   PrivacyKey = 0x697f414
	PrivacyKeyAboutPredict           PrivacyKey = 0xa486b761
	PrivacyKeyBirthdayPredict        PrivacyKey = 0x2000a518
)

type ReactionNotificationsFrom uint32

const (
	ReactionNotificationsFromContactsPredict ReactionNotificationsFrom = 0xbac3a61a
	ReactionNotificationsFromAllPredict      ReactionNotificationsFrom = 0x4b9e22a0
)

type ReportReason uint32

const (
	InputReportReasonSpamPredict            ReportReason = 0x58dbcab8
	InputReportReasonViolencePredict        ReportReason = 0x1e22c78d
	InputReportReasonPornographyPredict     ReportReason = 0x2e59d922
	InputReportReasonChildAbusePredict      ReportReason = 0xadf44ee3
	InputReportReasonOtherPredict           ReportReason = 0xc1e4a2b1
	InputReportReasonCopyrightPredict       ReportReason = 0x9b89f93a
	InputReportReasonGeoIrrelevantPredict   ReportReason = 0xdbd4feed
	InputReportReasonFakePredict            ReportReason = 0xf5ddd6e7
	InputReportReasonIllegalDrugsPredict    ReportReason = 0xa8eb2be
	InputReportReasonPersonalDetailsPredict ReportReason = 0x9ec7863d
)

type SecureValueType uint32

const (
	SecureValueTypePersonalDetailsPredict       SecureValueType = 0x9d2a81e3
	SecureValueTypePassportPredict              SecureValueType = 0x3dac6a00
	SecureValueTypeDriverLicensePredict         SecureValueType = 0x6e425c4
	SecureValueTypeIdentityCardPredict          SecureValueType = 0xa0d0744b
	SecureValueTypeInternalPassportPredict      SecureValueType = 0x99a48f23
	SecureValueTypeAddressPredict               SecureValueType = 0xcbe31e26
	SecureValueTypeUtilityBillPredict           SecureValueType = 0xfc36954e
	SecureValueTypeBankStatementPredict         SecureValueType = 0x89137c0d
	SecureValueTypeRentalAgreementPredict       SecureValueType = 0x8b883488
	SecureValueTypePassportRegistrationPredict  SecureValueType = 0x99e3806a
	SecureValueTypeTemporaryRegistrationPredict SecureValueType = 0xea02ec33
	SecureValueTypePhonePredict                 SecureValueType = 0xb320aadb
	SecureValueTypeEmailPredict                 SecureValueType = 0x8e3ca7ee
)

type TopPeerCategory uint32

const (
	TopPeerCategoryBotsPmPredict         TopPeerCategory = 0xab661b5b
	TopPeerCategoryBotsInlinePredict     TopPeerCategory = 0x148677e2
	TopPeerCategoryCorrespondentsPredict TopPeerCategory = 0x637b7ed
	TopPeerCategoryGroupsPredict         TopPeerCategory = 0xbd17a14a
	TopPeerCategoryChannelsPredict       TopPeerCategory = 0x161d9628
	TopPeerCategoryPhoneCallsPredict     TopPeerCategory = 0x1e76a78c
	TopPeerCategoryForwardUsersPredict   TopPeerCategory = 0xa8406ca9
	TopPeerCategoryForwardChatsPredict   TopPeerCategory = 0xfbeec0f0
	TopPeerCategoryBotsAppPredict        TopPeerCategory = 0xfd9e7bec
)

type True uint32

const (
	TruePredict True = 0x3fedd339
)

type AuthCodeType uint32

const (
	AuthCodeTypeSmsPredict         AuthCodeType = 0x72a3158c
	AuthCodeTypeCallPredict        AuthCodeType = 0x741cd3e3
	AuthCodeTypeFlashCallPredict   AuthCodeType = 0x226ccefb
	AuthCodeTypeMissedCallPredict  AuthCodeType = 0xd61ad6ee
	AuthCodeTypeFragmentSmsPredict AuthCodeType = 0x6ed998c
)

type StorageFileType uint32

const (
	StorageFileUnknownPredict StorageFileType = 0xaa963b05
	StorageFilePartialPredict StorageFileType = 0x40bc6f52
	StorageFileJpegPredict    StorageFileType = 0x7efe0e
	StorageFileGifPredict     StorageFileType = 0xcae1aadf
	StorageFilePngPredict     StorageFileType = 0xa4f63c0
	StorageFilePdfPredict     StorageFileType = 0xae1e508d
	StorageFileMp3Predict     StorageFileType = 0x528a0677
	StorageFileMovPredict     StorageFileType = 0x4b09ebbc
	StorageFileMp4Predict     StorageFileType = 0xb3cea0e4
	StorageFileWebpPredict    StorageFileType = 0x1081464c
)

type Requester interface {
	MakeRequest(ctx context.Context, msg []byte) ([]byte, error)
}

func request[IN any, OUT any](ctx context.Context, m Requester, in *IN, out *OUT) error {
	if msg, err := tl.Marshal(in); err != nil {
		return fmt.Errorf("marshaling: %w", err)
	} else if respRaw, err := m.MakeRequest(ctx, msg); err != nil {
		return fmt.Errorf("sending: %w", err)
	} else if err := tl.Unmarshal(respRaw, out); err != nil {
		return fmt.Errorf("got invalid response type: %w", err)
	}
	return nil
}

type InvokeAfterMsgRequestPredict[X any] struct {
	MsgID int64
	Query X
}

func (*InvokeAfterMsgRequestPredict[X]) CRC() uint32 {
	return 0xcb9f372d
}

func InvokeAfterMsg[X any](ctx context.Context, m Requester, i InvokeAfterMsgRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeAfterMsgsRequestPredict[X any] struct {
	MsgIds []int64
	Query  X
}

func (*InvokeAfterMsgsRequestPredict[X]) CRC() uint32 {
	return 0x3dc4b4f0
}

func InvokeAfterMsgs[X any](ctx context.Context, m Requester, i InvokeAfterMsgsRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InitConnectionRequestPredict[X any] struct {
	_              struct{} `tl:"flags,bitflag"`
	APIID          int32
	DeviceModel    string
	SystemVersion  string
	AppVersion     string
	SystemLangCode string
	LangPack       string
	LangCode       string
	Proxy          InputClientProxy `tl:",omitempty:flags:0"`
	Params         JSONValue        `tl:",omitempty:flags:1"`
	Query          X
}

func (*InitConnectionRequestPredict[X]) CRC() uint32 {
	return 0xc1cd5ea9
}

func InitConnection[X any](ctx context.Context, m Requester, i InitConnectionRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithLayerRequestPredict[X any] struct {
	Layer int32
	Query X
}

func (*InvokeWithLayerRequestPredict[X]) CRC() uint32 {
	return 0xda9b0d0d
}

func InvokeWithLayer[X any](ctx context.Context, m Requester, i InvokeWithLayerRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithoutUpdatesRequestPredict[X any] struct {
	Query X
}

func (*InvokeWithoutUpdatesRequestPredict[X]) CRC() uint32 {
	return 0xbf9459b7
}

func InvokeWithoutUpdates[X any](ctx context.Context, m Requester, i InvokeWithoutUpdatesRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithMessagesRangeRequestPredict[X any] struct {
	Range MessageRange
	Query X
}

func (*InvokeWithMessagesRangeRequestPredict[X]) CRC() uint32 {
	return 0x365275f2
}

func InvokeWithMessagesRange[X any](ctx context.Context, m Requester, i InvokeWithMessagesRangeRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithTakeoutRequestPredict[X any] struct {
	TakeoutID int64
	Query     X
}

func (*InvokeWithTakeoutRequestPredict[X]) CRC() uint32 {
	return 0xaca9fd2e
}

func InvokeWithTakeout[X any](ctx context.Context, m Requester, i InvokeWithTakeoutRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithBusinessConnectionRequestPredict[X any] struct {
	ConnectionID string
	Query        X
}

func (*InvokeWithBusinessConnectionRequestPredict[X]) CRC() uint32 {
	return 0xdd289f8e
}

func InvokeWithBusinessConnection[X any](ctx context.Context, m Requester, i InvokeWithBusinessConnectionRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithGooglePlayIntegrityRequestPredict[X any] struct {
	Nonce string
	Token string
	Query X
}

func (*InvokeWithGooglePlayIntegrityRequestPredict[X]) CRC() uint32 {
	return 0x1df92984
}

func InvokeWithGooglePlayIntegrity[X any](ctx context.Context, m Requester, i InvokeWithGooglePlayIntegrityRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithApnsSecretRequestPredict[X any] struct {
	Nonce  string
	Secret string
	Query  X
}

func (*InvokeWithApnsSecretRequestPredict[X]) CRC() uint32 {
	return 0xdae54f8
}

func InvokeWithApnsSecret[X any](ctx context.Context, m Requester, i InvokeWithApnsSecretRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type AccountRegisterDeviceRequestPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	NoMuted    bool     `tl:",omitempty:flags:0,implicit"`
	TokenType  int32
	Token      string
	AppSandbox bool
	Secret     []byte
	OtherUids  []int64
}

func (*AccountRegisterDeviceRequestPredict) CRC() uint32 {
	return 0xec86017a
}

func AccountRegisterDevice(ctx context.Context, m Requester, i AccountRegisterDeviceRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUnregisterDeviceRequestPredict struct {
	TokenType int32
	Token     string
	OtherUids []int64
}

func (*AccountUnregisterDeviceRequestPredict) CRC() uint32 {
	return 0x6a0d3206
}

func AccountUnregisterDevice(ctx context.Context, m Requester, i AccountUnregisterDeviceRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateNotifySettingsRequestPredict struct {
	Peer     InputNotifyPeer
	Settings InputPeerNotifySettings
}

func (*AccountUpdateNotifySettingsRequestPredict) CRC() uint32 {
	return 0x84be5b93
}

func AccountUpdateNotifySettings(ctx context.Context, m Requester, i AccountUpdateNotifySettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetNotifySettingsRequestPredict struct {
	Peer InputNotifyPeer
}

func (*AccountGetNotifySettingsRequestPredict) CRC() uint32 {
	return 0x12b3ad31
}

func AccountGetNotifySettings(ctx context.Context, m Requester, i AccountGetNotifySettingsRequestPredict) (PeerNotifySettings, error) {
	var res PeerNotifySettings
	return res, request(ctx, m, &i, &res)
}

type AccountResetNotifySettingsRequestPredict struct{}

func (*AccountResetNotifySettingsRequestPredict) CRC() uint32 {
	return 0xdb7e1747
}

func AccountResetNotifySettings(ctx context.Context, m Requester, i AccountResetNotifySettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateProfileRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	FirstName *string  `tl:",omitempty:flags:0"`
	LastName  *string  `tl:",omitempty:flags:1"`
	About     *string  `tl:",omitempty:flags:2"`
}

func (*AccountUpdateProfileRequestPredict) CRC() uint32 {
	return 0x78515775
}

func AccountUpdateProfile(ctx context.Context, m Requester, i AccountUpdateProfileRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateStatusRequestPredict struct {
	Offline bool
}

func (*AccountUpdateStatusRequestPredict) CRC() uint32 {
	return 0x6628562c
}

func AccountUpdateStatus(ctx context.Context, m Requester, i AccountUpdateStatusRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetWallPapersRequestPredict struct {
	Hash int64
}

func (*AccountGetWallPapersRequestPredict) CRC() uint32 {
	return 0x7967d36
}

func AccountGetWallPapers(ctx context.Context, m Requester, i AccountGetWallPapersRequestPredict) (AccountWallPapers, error) {
	var res AccountWallPapers
	return res, request(ctx, m, &i, &res)
}

type AccountReportPeerRequestPredict struct {
	Peer    InputPeer
	Reason  ReportReason
	Message string
}

func (*AccountReportPeerRequestPredict) CRC() uint32 {
	return 0xc5ba3d86
}

func AccountReportPeer(ctx context.Context, m Requester, i AccountReportPeerRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountCheckUsernameRequestPredict struct {
	Username string
}

func (*AccountCheckUsernameRequestPredict) CRC() uint32 {
	return 0x2714d86c
}

func AccountCheckUsername(ctx context.Context, m Requester, i AccountCheckUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateUsernameRequestPredict struct {
	Username string
}

func (*AccountUpdateUsernameRequestPredict) CRC() uint32 {
	return 0x3e0bdd7c
}

func AccountUpdateUsername(ctx context.Context, m Requester, i AccountUpdateUsernameRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type AccountGetPrivacyRequestPredict struct {
	Key InputPrivacyKey
}

func (*AccountGetPrivacyRequestPredict) CRC() uint32 {
	return 0xdadbc950
}

func AccountGetPrivacy(ctx context.Context, m Requester, i AccountGetPrivacyRequestPredict) (AccountPrivacyRules, error) {
	var res AccountPrivacyRules
	return res, request(ctx, m, &i, &res)
}

type AccountSetPrivacyRequestPredict struct {
	Key   InputPrivacyKey
	Rules []InputPrivacyRule
}

func (*AccountSetPrivacyRequestPredict) CRC() uint32 {
	return 0xc9f81ce8
}

func AccountSetPrivacy(ctx context.Context, m Requester, i AccountSetPrivacyRequestPredict) (AccountPrivacyRules, error) {
	var res AccountPrivacyRules
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteAccountRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Reason   string
	Password InputCheckPasswordSRP `tl:",omitempty:flags:0"`
}

func (*AccountDeleteAccountRequestPredict) CRC() uint32 {
	return 0xa2c0cf74
}

func AccountDeleteAccount(ctx context.Context, m Requester, i AccountDeleteAccountRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAccountTTLRequestPredict struct{}

func (*AccountGetAccountTTLRequestPredict) CRC() uint32 {
	return 0x8fc711d
}

func AccountGetAccountTTL(ctx context.Context, m Requester, i AccountGetAccountTTLRequestPredict) (AccountDaysTTL, error) {
	var res AccountDaysTTL
	return res, request(ctx, m, &i, &res)
}

type AccountSetAccountTTLRequestPredict struct {
	TTL AccountDaysTTL
}

func (*AccountSetAccountTTLRequestPredict) CRC() uint32 {
	return 0x2442485e
}

func AccountSetAccountTTL(ctx context.Context, m Requester, i AccountSetAccountTTLRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendChangePhoneCodeRequestPredict struct {
	PhoneNumber string
	Settings    CodeSettings
}

func (*AccountSendChangePhoneCodeRequestPredict) CRC() uint32 {
	return 0x82574ae5
}

func AccountSendChangePhoneCode(ctx context.Context, m Requester, i AccountSendChangePhoneCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountChangePhoneRequestPredict struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

func (*AccountChangePhoneRequestPredict) CRC() uint32 {
	return 0x70c32edb
}

func AccountChangePhone(ctx context.Context, m Requester, i AccountChangePhoneRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateDeviceLockedRequestPredict struct {
	Period int32
}

func (*AccountUpdateDeviceLockedRequestPredict) CRC() uint32 {
	return 0x38df3532
}

func AccountUpdateDeviceLocked(ctx context.Context, m Requester, i AccountUpdateDeviceLockedRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAuthorizationsRequestPredict struct{}

func (*AccountGetAuthorizationsRequestPredict) CRC() uint32 {
	return 0xe320c158
}

func AccountGetAuthorizations(ctx context.Context, m Requester, i AccountGetAuthorizationsRequestPredict) (AccountAuthorizations, error) {
	var res AccountAuthorizations
	return res, request(ctx, m, &i, &res)
}

type AccountResetAuthorizationRequestPredict struct {
	Hash int64
}

func (*AccountResetAuthorizationRequestPredict) CRC() uint32 {
	return 0xdf77f3bc
}

func AccountResetAuthorization(ctx context.Context, m Requester, i AccountResetAuthorizationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetPasswordRequestPredict struct{}

func (*AccountGetPasswordRequestPredict) CRC() uint32 {
	return 0x548a30f5
}

func AccountGetPassword(ctx context.Context, m Requester, i AccountGetPasswordRequestPredict) (AccountPassword, error) {
	var res AccountPassword
	return res, request(ctx, m, &i, &res)
}

type AccountGetPasswordSettingsRequestPredict struct {
	Password InputCheckPasswordSRP
}

func (*AccountGetPasswordSettingsRequestPredict) CRC() uint32 {
	return 0x9cd4eaf9
}

func AccountGetPasswordSettings(ctx context.Context, m Requester, i AccountGetPasswordSettingsRequestPredict) (AccountPasswordSettings, error) {
	var res AccountPasswordSettings
	return res, request(ctx, m, &i, &res)
}

type AccountUpdatePasswordSettingsRequestPredict struct {
	Password    InputCheckPasswordSRP
	NewSettings AccountPasswordInputSettings
}

func (*AccountUpdatePasswordSettingsRequestPredict) CRC() uint32 {
	return 0xa59b102f
}

func AccountUpdatePasswordSettings(ctx context.Context, m Requester, i AccountUpdatePasswordSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendConfirmPhoneCodeRequestPredict struct {
	Hash     string
	Settings CodeSettings
}

func (*AccountSendConfirmPhoneCodeRequestPredict) CRC() uint32 {
	return 0x1b3faa88
}

func AccountSendConfirmPhoneCode(ctx context.Context, m Requester, i AccountSendConfirmPhoneCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountConfirmPhoneRequestPredict struct {
	PhoneCodeHash string
	PhoneCode     string
}

func (*AccountConfirmPhoneRequestPredict) CRC() uint32 {
	return 0x5f2178c3
}

func AccountConfirmPhone(ctx context.Context, m Requester, i AccountConfirmPhoneRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetTmpPasswordRequestPredict struct {
	Password InputCheckPasswordSRP
	Period   int32
}

func (*AccountGetTmpPasswordRequestPredict) CRC() uint32 {
	return 0x449e0b51
}

func AccountGetTmpPassword(ctx context.Context, m Requester, i AccountGetTmpPasswordRequestPredict) (AccountTmpPassword, error) {
	var res AccountTmpPassword
	return res, request(ctx, m, &i, &res)
}

type AccountGetWebAuthorizationsRequestPredict struct{}

func (*AccountGetWebAuthorizationsRequestPredict) CRC() uint32 {
	return 0x182e6d6f
}

func AccountGetWebAuthorizations(ctx context.Context, m Requester, i AccountGetWebAuthorizationsRequestPredict) (AccountWebAuthorizations, error) {
	var res AccountWebAuthorizations
	return res, request(ctx, m, &i, &res)
}

type AccountResetWebAuthorizationRequestPredict struct {
	Hash int64
}

func (*AccountResetWebAuthorizationRequestPredict) CRC() uint32 {
	return 0x2d01b9ef
}

func AccountResetWebAuthorization(ctx context.Context, m Requester, i AccountResetWebAuthorizationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResetWebAuthorizationsRequestPredict struct{}

func (*AccountResetWebAuthorizationsRequestPredict) CRC() uint32 {
	return 0x682d2594
}

func AccountResetWebAuthorizations(ctx context.Context, m Requester, i AccountResetWebAuthorizationsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAllSecureValuesRequestPredict struct{}

func (*AccountGetAllSecureValuesRequestPredict) CRC() uint32 {
	return 0xb288bc7d
}

func AccountGetAllSecureValues(ctx context.Context, m Requester, i AccountGetAllSecureValuesRequestPredict) ([]SecureValue, error) {
	var res []SecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountGetSecureValueRequestPredict struct {
	Types []SecureValueType
}

func (*AccountGetSecureValueRequestPredict) CRC() uint32 {
	return 0x73665bc2
}

func AccountGetSecureValue(ctx context.Context, m Requester, i AccountGetSecureValueRequestPredict) ([]SecureValue, error) {
	var res []SecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountSaveSecureValueRequestPredict struct {
	Value          InputSecureValue
	SecureSecretID int64
}

func (*AccountSaveSecureValueRequestPredict) CRC() uint32 {
	return 0x899fe31d
}

func AccountSaveSecureValue(ctx context.Context, m Requester, i AccountSaveSecureValueRequestPredict) (SecureValue, error) {
	var res SecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteSecureValueRequestPredict struct {
	Types []SecureValueType
}

func (*AccountDeleteSecureValueRequestPredict) CRC() uint32 {
	return 0xb880bc4b
}

func AccountDeleteSecureValue(ctx context.Context, m Requester, i AccountDeleteSecureValueRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAuthorizationFormRequestPredict struct {
	BotID     int64
	Scope     string
	PublicKey string
}

func (*AccountGetAuthorizationFormRequestPredict) CRC() uint32 {
	return 0xa929597a
}

func AccountGetAuthorizationForm(ctx context.Context, m Requester, i AccountGetAuthorizationFormRequestPredict) (AccountAuthorizationForm, error) {
	var res AccountAuthorizationForm
	return res, request(ctx, m, &i, &res)
}

type AccountAcceptAuthorizationRequestPredict struct {
	BotID       int64
	Scope       string
	PublicKey   string
	ValueHashes []SecureValueHash
	Credentials SecureCredentialsEncrypted
}

func (*AccountAcceptAuthorizationRequestPredict) CRC() uint32 {
	return 0xf3ed4c73
}

func AccountAcceptAuthorization(ctx context.Context, m Requester, i AccountAcceptAuthorizationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendVerifyPhoneCodeRequestPredict struct {
	PhoneNumber string
	Settings    CodeSettings
}

func (*AccountSendVerifyPhoneCodeRequestPredict) CRC() uint32 {
	return 0xa5a356f9
}

func AccountSendVerifyPhoneCode(ctx context.Context, m Requester, i AccountSendVerifyPhoneCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountVerifyPhoneRequestPredict struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

func (*AccountVerifyPhoneRequestPredict) CRC() uint32 {
	return 0x4dd3a7f6
}

func AccountVerifyPhone(ctx context.Context, m Requester, i AccountVerifyPhoneRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendVerifyEmailCodeRequestPredict struct {
	Purpose EmailVerifyPurpose
	Email   string
}

func (*AccountSendVerifyEmailCodeRequestPredict) CRC() uint32 {
	return 0x98e037bb
}

func AccountSendVerifyEmailCode(ctx context.Context, m Requester, i AccountSendVerifyEmailCodeRequestPredict) (AccountSentEmailCode, error) {
	var res AccountSentEmailCode
	return res, request(ctx, m, &i, &res)
}

type AccountVerifyEmailRequestPredict struct {
	Purpose      EmailVerifyPurpose
	Verification EmailVerification
}

func (*AccountVerifyEmailRequestPredict) CRC() uint32 {
	return 0x32da4cf
}

func AccountVerifyEmail(ctx context.Context, m Requester, i AccountVerifyEmailRequestPredict) (AccountEmailVerified, error) {
	var res AccountEmailVerified
	return res, request(ctx, m, &i, &res)
}

type AccountInitTakeoutSessionRequestPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Contacts          bool     `tl:",omitempty:flags:0,implicit"`
	MessageUsers      bool     `tl:",omitempty:flags:1,implicit"`
	MessageChats      bool     `tl:",omitempty:flags:2,implicit"`
	MessageMegagroups bool     `tl:",omitempty:flags:3,implicit"`
	MessageChannels   bool     `tl:",omitempty:flags:4,implicit"`
	Files             bool     `tl:",omitempty:flags:5,implicit"`
	FileMaxSize       *int64   `tl:",omitempty:flags:5"`
}

func (*AccountInitTakeoutSessionRequestPredict) CRC() uint32 {
	return 0x8ef3eab0
}

func AccountInitTakeoutSession(ctx context.Context, m Requester, i AccountInitTakeoutSessionRequestPredict) (AccountTakeout, error) {
	var res AccountTakeout
	return res, request(ctx, m, &i, &res)
}

type AccountFinishTakeoutSessionRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Success bool     `tl:",omitempty:flags:0,implicit"`
}

func (*AccountFinishTakeoutSessionRequestPredict) CRC() uint32 {
	return 0x1d2652ee
}

func AccountFinishTakeoutSession(ctx context.Context, m Requester, i AccountFinishTakeoutSessionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountConfirmPasswordEmailRequestPredict struct {
	Code string
}

func (*AccountConfirmPasswordEmailRequestPredict) CRC() uint32 {
	return 0x8fdf1920
}

func AccountConfirmPasswordEmail(ctx context.Context, m Requester, i AccountConfirmPasswordEmailRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResendPasswordEmailRequestPredict struct{}

func (*AccountResendPasswordEmailRequestPredict) CRC() uint32 {
	return 0x7a7f2a15
}

func AccountResendPasswordEmail(ctx context.Context, m Requester, i AccountResendPasswordEmailRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountCancelPasswordEmailRequestPredict struct{}

func (*AccountCancelPasswordEmailRequestPredict) CRC() uint32 {
	return 0xc1cbd5b6
}

func AccountCancelPasswordEmail(ctx context.Context, m Requester, i AccountCancelPasswordEmailRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetContactSignUpNotificationRequestPredict struct{}

func (*AccountGetContactSignUpNotificationRequestPredict) CRC() uint32 {
	return 0x9f07c728
}

func AccountGetContactSignUpNotification(ctx context.Context, m Requester, i AccountGetContactSignUpNotificationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSetContactSignUpNotificationRequestPredict struct {
	Silent bool
}

func (*AccountSetContactSignUpNotificationRequestPredict) CRC() uint32 {
	return 0xcff43f61
}

func AccountSetContactSignUpNotification(ctx context.Context, m Requester, i AccountSetContactSignUpNotificationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetNotifyExceptionsRequestPredict struct {
	_              struct{}        `tl:"flags,bitflag"`
	CompareSound   bool            `tl:",omitempty:flags:1,implicit"`
	CompareStories bool            `tl:",omitempty:flags:2,implicit"`
	Peer           InputNotifyPeer `tl:",omitempty:flags:0"`
}

func (*AccountGetNotifyExceptionsRequestPredict) CRC() uint32 {
	return 0x53577479
}

func AccountGetNotifyExceptions(ctx context.Context, m Requester, i AccountGetNotifyExceptionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type AccountGetWallPaperRequestPredict struct {
	Wallpaper InputWallPaper
}

func (*AccountGetWallPaperRequestPredict) CRC() uint32 {
	return 0xfc8ddbea
}

func AccountGetWallPaper(ctx context.Context, m Requester, i AccountGetWallPaperRequestPredict) (WallPaper, error) {
	var res WallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountUploadWallPaperRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	ForChat  bool     `tl:",omitempty:flags:0,implicit"`
	File     InputFile
	MimeType string
	Settings WallPaperSettings
}

func (*AccountUploadWallPaperRequestPredict) CRC() uint32 {
	return 0xe39a8f03
}

func AccountUploadWallPaper(ctx context.Context, m Requester, i AccountUploadWallPaperRequestPredict) (WallPaper, error) {
	var res WallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountSaveWallPaperRequestPredict struct {
	Wallpaper InputWallPaper
	Unsave    bool
	Settings  WallPaperSettings
}

func (*AccountSaveWallPaperRequestPredict) CRC() uint32 {
	return 0x6c5a5b37
}

func AccountSaveWallPaper(ctx context.Context, m Requester, i AccountSaveWallPaperRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountInstallWallPaperRequestPredict struct {
	Wallpaper InputWallPaper
	Settings  WallPaperSettings
}

func (*AccountInstallWallPaperRequestPredict) CRC() uint32 {
	return 0xfeed5769
}

func AccountInstallWallPaper(ctx context.Context, m Requester, i AccountInstallWallPaperRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResetWallPapersRequestPredict struct{}

func (*AccountResetWallPapersRequestPredict) CRC() uint32 {
	return 0xbb3b9804
}

func AccountResetWallPapers(ctx context.Context, m Requester, i AccountResetWallPapersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAutoDownloadSettingsRequestPredict struct{}

func (*AccountGetAutoDownloadSettingsRequestPredict) CRC() uint32 {
	return 0x56da0b3f
}

func AccountGetAutoDownloadSettings(ctx context.Context, m Requester, i AccountGetAutoDownloadSettingsRequestPredict) (AccountAutoDownloadSettings, error) {
	var res AccountAutoDownloadSettings
	return res, request(ctx, m, &i, &res)
}

type AccountSaveAutoDownloadSettingsRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Low      bool     `tl:",omitempty:flags:0,implicit"`
	High     bool     `tl:",omitempty:flags:1,implicit"`
	Settings AutoDownloadSettings
}

func (*AccountSaveAutoDownloadSettingsRequestPredict) CRC() uint32 {
	return 0x76f36233
}

func AccountSaveAutoDownloadSettings(ctx context.Context, m Requester, i AccountSaveAutoDownloadSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUploadThemeRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	File     InputFile
	Thumb    InputFile `tl:",omitempty:flags:0"`
	FileName string
	MimeType string
}

func (*AccountUploadThemeRequestPredict) CRC() uint32 {
	return 0x1c3db333
}

func AccountUploadTheme(ctx context.Context, m Requester, i AccountUploadThemeRequestPredict) (Document, error) {
	var res Document
	return res, request(ctx, m, &i, &res)
}

type AccountCreateThemeRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Slug     string
	Title    string
	Document InputDocument        `tl:",omitempty:flags:2"`
	Settings []InputThemeSettings `tl:",omitempty:flags:3"`
}

func (*AccountCreateThemeRequestPredict) CRC() uint32 {
	return 0x652e4400
}

func AccountCreateTheme(ctx context.Context, m Requester, i AccountCreateThemeRequestPredict) (Theme, error) {
	var res Theme
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateThemeRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Format   string
	Theme    InputTheme
	Slug     *string              `tl:",omitempty:flags:0"`
	Title    *string              `tl:",omitempty:flags:1"`
	Document InputDocument        `tl:",omitempty:flags:2"`
	Settings []InputThemeSettings `tl:",omitempty:flags:3"`
}

func (*AccountUpdateThemeRequestPredict) CRC() uint32 {
	return 0x2bf40ccc
}

func AccountUpdateTheme(ctx context.Context, m Requester, i AccountUpdateThemeRequestPredict) (Theme, error) {
	var res Theme
	return res, request(ctx, m, &i, &res)
}

type AccountSaveThemeRequestPredict struct {
	Theme  InputTheme
	Unsave bool
}

func (*AccountSaveThemeRequestPredict) CRC() uint32 {
	return 0xf257106c
}

func AccountSaveTheme(ctx context.Context, m Requester, i AccountSaveThemeRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountInstallThemeRequestPredict struct {
	_         struct{}   `tl:"flags,bitflag"`
	Dark      bool       `tl:",omitempty:flags:0,implicit"`
	Theme     InputTheme `tl:",omitempty:flags:1"`
	Format    *string    `tl:",omitempty:flags:2"`
	BaseTheme BaseTheme  `tl:",omitempty:flags:3"`
}

func (*AccountInstallThemeRequestPredict) CRC() uint32 {
	return 0xc727bb3b
}

func AccountInstallTheme(ctx context.Context, m Requester, i AccountInstallThemeRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetThemeRequestPredict struct {
	Format string
	Theme  InputTheme
}

func (*AccountGetThemeRequestPredict) CRC() uint32 {
	return 0x3a5869ec
}

func AccountGetTheme(ctx context.Context, m Requester, i AccountGetThemeRequestPredict) (Theme, error) {
	var res Theme
	return res, request(ctx, m, &i, &res)
}

type AccountGetThemesRequestPredict struct {
	Format string
	Hash   int64
}

func (*AccountGetThemesRequestPredict) CRC() uint32 {
	return 0x7206e458
}

func AccountGetThemes(ctx context.Context, m Requester, i AccountGetThemesRequestPredict) (AccountThemes, error) {
	var res AccountThemes
	return res, request(ctx, m, &i, &res)
}

type AccountSetContentSettingsRequestPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	SensitiveEnabled bool     `tl:",omitempty:flags:0,implicit"`
}

func (*AccountSetContentSettingsRequestPredict) CRC() uint32 {
	return 0xb574b16b
}

func AccountSetContentSettings(ctx context.Context, m Requester, i AccountSetContentSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetContentSettingsRequestPredict struct{}

func (*AccountGetContentSettingsRequestPredict) CRC() uint32 {
	return 0x8b9b4dae
}

func AccountGetContentSettings(ctx context.Context, m Requester, i AccountGetContentSettingsRequestPredict) (AccountContentSettings, error) {
	var res AccountContentSettings
	return res, request(ctx, m, &i, &res)
}

type AccountGetMultiWallPapersRequestPredict struct {
	Wallpapers []InputWallPaper
}

func (*AccountGetMultiWallPapersRequestPredict) CRC() uint32 {
	return 0x65ad71dc
}

func AccountGetMultiWallPapers(ctx context.Context, m Requester, i AccountGetMultiWallPapersRequestPredict) ([]WallPaper, error) {
	var res []WallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountGetGlobalPrivacySettingsRequestPredict struct{}

func (*AccountGetGlobalPrivacySettingsRequestPredict) CRC() uint32 {
	return 0xeb2b4cf6
}

func AccountGetGlobalPrivacySettings(ctx context.Context, m Requester, i AccountGetGlobalPrivacySettingsRequestPredict) (GlobalPrivacySettings, error) {
	var res GlobalPrivacySettings
	return res, request(ctx, m, &i, &res)
}

type AccountSetGlobalPrivacySettingsRequestPredict struct {
	Settings GlobalPrivacySettings
}

func (*AccountSetGlobalPrivacySettingsRequestPredict) CRC() uint32 {
	return 0x1edaaac2
}

func AccountSetGlobalPrivacySettings(ctx context.Context, m Requester, i AccountSetGlobalPrivacySettingsRequestPredict) (GlobalPrivacySettings, error) {
	var res GlobalPrivacySettings
	return res, request(ctx, m, &i, &res)
}

type AccountReportProfilePhotoRequestPredict struct {
	Peer    InputPeer
	PhotoID InputPhoto
	Reason  ReportReason
	Message string
}

func (*AccountReportProfilePhotoRequestPredict) CRC() uint32 {
	return 0xfa8cc6f5
}

func AccountReportProfilePhoto(ctx context.Context, m Requester, i AccountReportProfilePhotoRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountResetPasswordRequestPredict struct{}

func (*AccountResetPasswordRequestPredict) CRC() uint32 {
	return 0x9308ce1b
}

func AccountResetPassword(ctx context.Context, m Requester, i AccountResetPasswordRequestPredict) (AccountResetPasswordResult, error) {
	var res AccountResetPasswordResult
	return res, request(ctx, m, &i, &res)
}

type AccountDeclinePasswordResetRequestPredict struct{}

func (*AccountDeclinePasswordResetRequestPredict) CRC() uint32 {
	return 0x4c9409f6
}

func AccountDeclinePasswordReset(ctx context.Context, m Requester, i AccountDeclinePasswordResetRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetChatThemesRequestPredict struct {
	Hash int64
}

func (*AccountGetChatThemesRequestPredict) CRC() uint32 {
	return 0xd638de89
}

func AccountGetChatThemes(ctx context.Context, m Requester, i AccountGetChatThemesRequestPredict) (AccountThemes, error) {
	var res AccountThemes
	return res, request(ctx, m, &i, &res)
}

type AccountSetAuthorizationTTLRequestPredict struct {
	AuthorizationTTLDays int32
}

func (*AccountSetAuthorizationTTLRequestPredict) CRC() uint32 {
	return 0xbf899aa0
}

func AccountSetAuthorizationTTL(ctx context.Context, m Requester, i AccountSetAuthorizationTTLRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountChangeAuthorizationSettingsRequestPredict struct {
	_                         struct{} `tl:"flags,bitflag"`
	Confirmed                 bool     `tl:",omitempty:flags:3,implicit"`
	Hash                      int64
	EncryptedRequestsDisabled *bool `tl:",omitempty:flags:0"`
	CallRequestsDisabled      *bool `tl:",omitempty:flags:1"`
}

func (*AccountChangeAuthorizationSettingsRequestPredict) CRC() uint32 {
	return 0x40f48462
}

func AccountChangeAuthorizationSettings(ctx context.Context, m Requester, i AccountChangeAuthorizationSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetSavedRingtonesRequestPredict struct {
	Hash int64
}

func (*AccountGetSavedRingtonesRequestPredict) CRC() uint32 {
	return 0xe1902288
}

func AccountGetSavedRingtones(ctx context.Context, m Requester, i AccountGetSavedRingtonesRequestPredict) (AccountSavedRingtones, error) {
	var res AccountSavedRingtones
	return res, request(ctx, m, &i, &res)
}

type AccountSaveRingtoneRequestPredict struct {
	ID     InputDocument
	Unsave bool
}

func (*AccountSaveRingtoneRequestPredict) CRC() uint32 {
	return 0x3dea5b03
}

func AccountSaveRingtone(ctx context.Context, m Requester, i AccountSaveRingtoneRequestPredict) (AccountSavedRingtone, error) {
	var res AccountSavedRingtone
	return res, request(ctx, m, &i, &res)
}

type AccountUploadRingtoneRequestPredict struct {
	File     InputFile
	FileName string
	MimeType string
}

func (*AccountUploadRingtoneRequestPredict) CRC() uint32 {
	return 0x831a83a2
}

func AccountUploadRingtone(ctx context.Context, m Requester, i AccountUploadRingtoneRequestPredict) (Document, error) {
	var res Document
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateEmojiStatusRequestPredict struct {
	EmojiStatus EmojiStatus
}

func (*AccountUpdateEmojiStatusRequestPredict) CRC() uint32 {
	return 0xfbd3de6b
}

func AccountUpdateEmojiStatus(ctx context.Context, m Requester, i AccountUpdateEmojiStatusRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultEmojiStatusesRequestPredict struct {
	Hash int64
}

func (*AccountGetDefaultEmojiStatusesRequestPredict) CRC() uint32 {
	return 0xd6753386
}

func AccountGetDefaultEmojiStatuses(ctx context.Context, m Requester, i AccountGetDefaultEmojiStatusesRequestPredict) (AccountEmojiStatuses, error) {
	var res AccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountGetRecentEmojiStatusesRequestPredict struct {
	Hash int64
}

func (*AccountGetRecentEmojiStatusesRequestPredict) CRC() uint32 {
	return 0xf578105
}

func AccountGetRecentEmojiStatuses(ctx context.Context, m Requester, i AccountGetRecentEmojiStatusesRequestPredict) (AccountEmojiStatuses, error) {
	var res AccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountClearRecentEmojiStatusesRequestPredict struct{}

func (*AccountClearRecentEmojiStatusesRequestPredict) CRC() uint32 {
	return 0x18201aae
}

func AccountClearRecentEmojiStatuses(ctx context.Context, m Requester, i AccountClearRecentEmojiStatusesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountReorderUsernamesRequestPredict struct {
	Order []string
}

func (*AccountReorderUsernamesRequestPredict) CRC() uint32 {
	return 0xef500eab
}

func AccountReorderUsernames(ctx context.Context, m Requester, i AccountReorderUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountToggleUsernameRequestPredict struct {
	Username string
	Active   bool
}

func (*AccountToggleUsernameRequestPredict) CRC() uint32 {
	return 0x58d6b376
}

func AccountToggleUsername(ctx context.Context, m Requester, i AccountToggleUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultProfilePhotoEmojisRequestPredict struct {
	Hash int64
}

func (*AccountGetDefaultProfilePhotoEmojisRequestPredict) CRC() uint32 {
	return 0xe2750328
}

func AccountGetDefaultProfilePhotoEmojis(ctx context.Context, m Requester, i AccountGetDefaultProfilePhotoEmojisRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultGroupPhotoEmojisRequestPredict struct {
	Hash int64
}

func (*AccountGetDefaultGroupPhotoEmojisRequestPredict) CRC() uint32 {
	return 0x915860ae
}

func AccountGetDefaultGroupPhotoEmojis(ctx context.Context, m Requester, i AccountGetDefaultGroupPhotoEmojisRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetAutoSaveSettingsRequestPredict struct{}

func (*AccountGetAutoSaveSettingsRequestPredict) CRC() uint32 {
	return 0xadcbbcda
}

func AccountGetAutoSaveSettings(ctx context.Context, m Requester, i AccountGetAutoSaveSettingsRequestPredict) (AccountAutoSaveSettings, error) {
	var res AccountAutoSaveSettings
	return res, request(ctx, m, &i, &res)
}

type AccountSaveAutoSaveSettingsRequestPredict struct {
	_          struct{}  `tl:"flags,bitflag"`
	Users      bool      `tl:",omitempty:flags:0,implicit"`
	Chats      bool      `tl:",omitempty:flags:1,implicit"`
	Broadcasts bool      `tl:",omitempty:flags:2,implicit"`
	Peer       InputPeer `tl:",omitempty:flags:3"`
	Settings   AutoSaveSettings
}

func (*AccountSaveAutoSaveSettingsRequestPredict) CRC() uint32 {
	return 0xd69b8361
}

func AccountSaveAutoSaveSettings(ctx context.Context, m Requester, i AccountSaveAutoSaveSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteAutoSaveExceptionsRequestPredict struct{}

func (*AccountDeleteAutoSaveExceptionsRequestPredict) CRC() uint32 {
	return 0x53bc0020
}

func AccountDeleteAutoSaveExceptions(ctx context.Context, m Requester, i AccountDeleteAutoSaveExceptionsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountInvalidateSignInCodesRequestPredict struct {
	Codes []string
}

func (*AccountInvalidateSignInCodesRequestPredict) CRC() uint32 {
	return 0xca8ae8ba
}

func AccountInvalidateSignInCodes(ctx context.Context, m Requester, i AccountInvalidateSignInCodesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateColorRequestPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	ForProfile        bool     `tl:",omitempty:flags:1,implicit"`
	Color             *int32   `tl:",omitempty:flags:2"`
	BackgroundEmojiID *int64   `tl:",omitempty:flags:0"`
}

func (*AccountUpdateColorRequestPredict) CRC() uint32 {
	return 0x7cefa15d
}

func AccountUpdateColor(ctx context.Context, m Requester, i AccountUpdateColorRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultBackgroundEmojisRequestPredict struct {
	Hash int64
}

func (*AccountGetDefaultBackgroundEmojisRequestPredict) CRC() uint32 {
	return 0xa60ab9ce
}

func AccountGetDefaultBackgroundEmojis(ctx context.Context, m Requester, i AccountGetDefaultBackgroundEmojisRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetChannelDefaultEmojiStatusesRequestPredict struct {
	Hash int64
}

func (*AccountGetChannelDefaultEmojiStatusesRequestPredict) CRC() uint32 {
	return 0x7727a7d5
}

func AccountGetChannelDefaultEmojiStatuses(ctx context.Context, m Requester, i AccountGetChannelDefaultEmojiStatusesRequestPredict) (AccountEmojiStatuses, error) {
	var res AccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountGetChannelRestrictedStatusEmojisRequestPredict struct {
	Hash int64
}

func (*AccountGetChannelRestrictedStatusEmojisRequestPredict) CRC() uint32 {
	return 0x35a9e0d5
}

func AccountGetChannelRestrictedStatusEmojis(ctx context.Context, m Requester, i AccountGetChannelRestrictedStatusEmojisRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateBusinessWorkHoursRequestPredict struct {
	_                 struct{}          `tl:"flags,bitflag"`
	BusinessWorkHours BusinessWorkHours `tl:",omitempty:flags:0"`
}

func (*AccountUpdateBusinessWorkHoursRequestPredict) CRC() uint32 {
	return 0x4b00e066
}

func AccountUpdateBusinessWorkHours(ctx context.Context, m Requester, i AccountUpdateBusinessWorkHoursRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateBusinessLocationRequestPredict struct {
	_        struct{}      `tl:"flags,bitflag"`
	GeoPoint InputGeoPoint `tl:",omitempty:flags:1"`
	Address  *string       `tl:",omitempty:flags:0"`
}

func (*AccountUpdateBusinessLocationRequestPredict) CRC() uint32 {
	return 0x9e6b131a
}

func AccountUpdateBusinessLocation(ctx context.Context, m Requester, i AccountUpdateBusinessLocationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateBusinessGreetingMessageRequestPredict struct {
	_       struct{}                     `tl:"flags,bitflag"`
	Message InputBusinessGreetingMessage `tl:",omitempty:flags:0"`
}

func (*AccountUpdateBusinessGreetingMessageRequestPredict) CRC() uint32 {
	return 0x66cdafc4
}

func AccountUpdateBusinessGreetingMessage(ctx context.Context, m Requester, i AccountUpdateBusinessGreetingMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateBusinessAwayMessageRequestPredict struct {
	_       struct{}                 `tl:"flags,bitflag"`
	Message InputBusinessAwayMessage `tl:",omitempty:flags:0"`
}

func (*AccountUpdateBusinessAwayMessageRequestPredict) CRC() uint32 {
	return 0xa26a7fa5
}

func AccountUpdateBusinessAwayMessage(ctx context.Context, m Requester, i AccountUpdateBusinessAwayMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateConnectedBotRequestPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	CanReply   bool     `tl:",omitempty:flags:0,implicit"`
	Deleted    bool     `tl:",omitempty:flags:1,implicit"`
	Bot        InputUser
	Recipients InputBusinessBotRecipients
}

func (*AccountUpdateConnectedBotRequestPredict) CRC() uint32 {
	return 0x43d8521d
}

func AccountUpdateConnectedBot(ctx context.Context, m Requester, i AccountUpdateConnectedBotRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type AccountGetConnectedBotsRequestPredict struct{}

func (*AccountGetConnectedBotsRequestPredict) CRC() uint32 {
	return 0x4ea4c80f
}

func AccountGetConnectedBots(ctx context.Context, m Requester, i AccountGetConnectedBotsRequestPredict) (AccountConnectedBots, error) {
	var res AccountConnectedBots
	return res, request(ctx, m, &i, &res)
}

type AccountGetBotBusinessConnectionRequestPredict struct {
	ConnectionID string
}

func (*AccountGetBotBusinessConnectionRequestPredict) CRC() uint32 {
	return 0x76a86270
}

func AccountGetBotBusinessConnection(ctx context.Context, m Requester, i AccountGetBotBusinessConnectionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateBusinessIntroRequestPredict struct {
	_     struct{}           `tl:"flags,bitflag"`
	Intro InputBusinessIntro `tl:",omitempty:flags:0"`
}

func (*AccountUpdateBusinessIntroRequestPredict) CRC() uint32 {
	return 0xa614d034
}

func AccountUpdateBusinessIntro(ctx context.Context, m Requester, i AccountUpdateBusinessIntroRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountToggleConnectedBotPausedRequestPredict struct {
	Peer   InputPeer
	Paused bool
}

func (*AccountToggleConnectedBotPausedRequestPredict) CRC() uint32 {
	return 0x646e1097
}

func AccountToggleConnectedBotPaused(ctx context.Context, m Requester, i AccountToggleConnectedBotPausedRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountDisablePeerConnectedBotRequestPredict struct {
	Peer InputPeer
}

func (*AccountDisablePeerConnectedBotRequestPredict) CRC() uint32 {
	return 0x5e437ed9
}

func AccountDisablePeerConnectedBot(ctx context.Context, m Requester, i AccountDisablePeerConnectedBotRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateBirthdayRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Birthday Birthday `tl:",omitempty:flags:0"`
}

func (*AccountUpdateBirthdayRequestPredict) CRC() uint32 {
	return 0xcc6e0c11
}

func AccountUpdateBirthday(ctx context.Context, m Requester, i AccountUpdateBirthdayRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountCreateBusinessChatLinkRequestPredict struct {
	Link InputBusinessChatLink
}

func (*AccountCreateBusinessChatLinkRequestPredict) CRC() uint32 {
	return 0x8851e68e
}

func AccountCreateBusinessChatLink(ctx context.Context, m Requester, i AccountCreateBusinessChatLinkRequestPredict) (BusinessChatLink, error) {
	var res BusinessChatLink
	return res, request(ctx, m, &i, &res)
}

type AccountEditBusinessChatLinkRequestPredict struct {
	Slug string
	Link InputBusinessChatLink
}

func (*AccountEditBusinessChatLinkRequestPredict) CRC() uint32 {
	return 0x8c3410af
}

func AccountEditBusinessChatLink(ctx context.Context, m Requester, i AccountEditBusinessChatLinkRequestPredict) (BusinessChatLink, error) {
	var res BusinessChatLink
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteBusinessChatLinkRequestPredict struct {
	Slug string
}

func (*AccountDeleteBusinessChatLinkRequestPredict) CRC() uint32 {
	return 0x60073674
}

func AccountDeleteBusinessChatLink(ctx context.Context, m Requester, i AccountDeleteBusinessChatLinkRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetBusinessChatLinksRequestPredict struct{}

func (*AccountGetBusinessChatLinksRequestPredict) CRC() uint32 {
	return 0x6f70dde1
}

func AccountGetBusinessChatLinks(ctx context.Context, m Requester, i AccountGetBusinessChatLinksRequestPredict) (AccountBusinessChatLinks, error) {
	var res AccountBusinessChatLinks
	return res, request(ctx, m, &i, &res)
}

type AccountResolveBusinessChatLinkRequestPredict struct {
	Slug string
}

func (*AccountResolveBusinessChatLinkRequestPredict) CRC() uint32 {
	return 0x5492e5ee
}

func AccountResolveBusinessChatLink(ctx context.Context, m Requester, i AccountResolveBusinessChatLinkRequestPredict) (AccountResolvedBusinessChatLinks, error) {
	var res AccountResolvedBusinessChatLinks
	return res, request(ctx, m, &i, &res)
}

type AccountUpdatePersonalChannelRequestPredict struct {
	Channel InputChannel
}

func (*AccountUpdatePersonalChannelRequestPredict) CRC() uint32 {
	return 0xd94305e0
}

func AccountUpdatePersonalChannel(ctx context.Context, m Requester, i AccountUpdatePersonalChannelRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountToggleSponsoredMessagesRequestPredict struct {
	Enabled bool
}

func (*AccountToggleSponsoredMessagesRequestPredict) CRC() uint32 {
	return 0xb9d9a38d
}

func AccountToggleSponsoredMessages(ctx context.Context, m Requester, i AccountToggleSponsoredMessagesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetReactionsNotifySettingsRequestPredict struct{}

func (*AccountGetReactionsNotifySettingsRequestPredict) CRC() uint32 {
	return 0x6dd654c
}

func AccountGetReactionsNotifySettings(ctx context.Context, m Requester, i AccountGetReactionsNotifySettingsRequestPredict) (ReactionsNotifySettings, error) {
	var res ReactionsNotifySettings
	return res, request(ctx, m, &i, &res)
}

type AccountSetReactionsNotifySettingsRequestPredict struct {
	Settings ReactionsNotifySettings
}

func (*AccountSetReactionsNotifySettingsRequestPredict) CRC() uint32 {
	return 0x316ce548
}

func AccountSetReactionsNotifySettings(ctx context.Context, m Requester, i AccountSetReactionsNotifySettingsRequestPredict) (ReactionsNotifySettings, error) {
	var res ReactionsNotifySettings
	return res, request(ctx, m, &i, &res)
}

type AuthSendCodeRequestPredict struct {
	PhoneNumber string
	APIID       int32
	APIHash     string
	Settings    CodeSettings
}

func (*AuthSendCodeRequestPredict) CRC() uint32 {
	return 0xa677244f
}

func AuthSendCode(ctx context.Context, m Requester, i AuthSendCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AuthSignUpRequestPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	NoJoinedNotifications bool     `tl:",omitempty:flags:0,implicit"`
	PhoneNumber           string
	PhoneCodeHash         string
	FirstName             string
	LastName              string
}

func (*AuthSignUpRequestPredict) CRC() uint32 {
	return 0xaac7b717
}

func AuthSignUp(ctx context.Context, m Requester, i AuthSignUpRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthSignInRequestPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	PhoneNumber       string
	PhoneCodeHash     string
	PhoneCode         *string           `tl:",omitempty:flags:0"`
	EmailVerification EmailVerification `tl:",omitempty:flags:1"`
}

func (*AuthSignInRequestPredict) CRC() uint32 {
	return 0x8d52a951
}

func AuthSignIn(ctx context.Context, m Requester, i AuthSignInRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthLogOutRequestPredict struct{}

func (*AuthLogOutRequestPredict) CRC() uint32 {
	return 0x3e72ba19
}

func AuthLogOut(ctx context.Context, m Requester, i AuthLogOutRequestPredict) (AuthLoggedOut, error) {
	var res AuthLoggedOut
	return res, request(ctx, m, &i, &res)
}

type AuthResetAuthorizationsRequestPredict struct{}

func (*AuthResetAuthorizationsRequestPredict) CRC() uint32 {
	return 0x9fab0d1a
}

func AuthResetAuthorizations(ctx context.Context, m Requester, i AuthResetAuthorizationsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthExportAuthorizationRequestPredict struct {
	DcID int32
}

func (*AuthExportAuthorizationRequestPredict) CRC() uint32 {
	return 0xe5bfffcd
}

func AuthExportAuthorization(ctx context.Context, m Requester, i AuthExportAuthorizationRequestPredict) (AuthExportedAuthorization, error) {
	var res AuthExportedAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthImportAuthorizationRequestPredict struct {
	ID    int64
	Bytes []byte
}

func (*AuthImportAuthorizationRequestPredict) CRC() uint32 {
	return 0xa57a7dad
}

func AuthImportAuthorization(ctx context.Context, m Requester, i AuthImportAuthorizationRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthBindTempAuthKeyRequestPredict struct {
	PermAuthKeyID    int64
	Nonce            int64
	ExpiresAt        int32
	EncryptedMessage []byte
}

func (*AuthBindTempAuthKeyRequestPredict) CRC() uint32 {
	return 0xcdd42a05
}

func AuthBindTempAuthKey(ctx context.Context, m Requester, i AuthBindTempAuthKeyRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthImportBotAuthorizationRequestPredict struct {
	Flags        int32
	APIID        int32
	APIHash      string
	BotAuthToken string
}

func (*AuthImportBotAuthorizationRequestPredict) CRC() uint32 {
	return 0x67a3ff2c
}

func AuthImportBotAuthorization(ctx context.Context, m Requester, i AuthImportBotAuthorizationRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthCheckPasswordRequestPredict struct {
	Password InputCheckPasswordSRP
}

func (*AuthCheckPasswordRequestPredict) CRC() uint32 {
	return 0xd18b4d16
}

func AuthCheckPassword(ctx context.Context, m Requester, i AuthCheckPasswordRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthRequestPasswordRecoveryRequestPredict struct{}

func (*AuthRequestPasswordRecoveryRequestPredict) CRC() uint32 {
	return 0xd897bc66
}

func AuthRequestPasswordRecovery(ctx context.Context, m Requester, i AuthRequestPasswordRecoveryRequestPredict) (AuthPasswordRecovery, error) {
	var res AuthPasswordRecovery
	return res, request(ctx, m, &i, &res)
}

type AuthRecoverPasswordRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Code        string
	NewSettings AccountPasswordInputSettings `tl:",omitempty:flags:0"`
}

func (*AuthRecoverPasswordRequestPredict) CRC() uint32 {
	return 0x37096c70
}

func AuthRecoverPassword(ctx context.Context, m Requester, i AuthRecoverPasswordRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthResendCodeRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	PhoneNumber   string
	PhoneCodeHash string
	Reason        *string `tl:",omitempty:flags:0"`
}

func (*AuthResendCodeRequestPredict) CRC() uint32 {
	return 0xcae47523
}

func AuthResendCode(ctx context.Context, m Requester, i AuthResendCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AuthCancelCodeRequestPredict struct {
	PhoneNumber   string
	PhoneCodeHash string
}

func (*AuthCancelCodeRequestPredict) CRC() uint32 {
	return 0x1f040578
}

func AuthCancelCode(ctx context.Context, m Requester, i AuthCancelCodeRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthDropTempAuthKeysRequestPredict struct {
	ExceptAuthKeys []int64
}

func (*AuthDropTempAuthKeysRequestPredict) CRC() uint32 {
	return 0x8e48a188
}

func AuthDropTempAuthKeys(ctx context.Context, m Requester, i AuthDropTempAuthKeysRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthExportLoginTokenRequestPredict struct {
	APIID     int32
	APIHash   string
	ExceptIds []int64
}

func (*AuthExportLoginTokenRequestPredict) CRC() uint32 {
	return 0xb7e085fe
}

func AuthExportLoginToken(ctx context.Context, m Requester, i AuthExportLoginTokenRequestPredict) (AuthLoginToken, error) {
	var res AuthLoginToken
	return res, request(ctx, m, &i, &res)
}

type AuthImportLoginTokenRequestPredict struct {
	Token []byte
}

func (*AuthImportLoginTokenRequestPredict) CRC() uint32 {
	return 0x95ac5ce4
}

func AuthImportLoginToken(ctx context.Context, m Requester, i AuthImportLoginTokenRequestPredict) (AuthLoginToken, error) {
	var res AuthLoginToken
	return res, request(ctx, m, &i, &res)
}

type AuthAcceptLoginTokenRequestPredict struct {
	Token []byte
}

func (*AuthAcceptLoginTokenRequestPredict) CRC() uint32 {
	return 0xe894ad4d
}

func AuthAcceptLoginToken(ctx context.Context, m Requester, i AuthAcceptLoginTokenRequestPredict) (Authorization, error) {
	var res Authorization
	return res, request(ctx, m, &i, &res)
}

type AuthCheckRecoveryPasswordRequestPredict struct {
	Code string
}

func (*AuthCheckRecoveryPasswordRequestPredict) CRC() uint32 {
	return 0xd36bf79
}

func AuthCheckRecoveryPassword(ctx context.Context, m Requester, i AuthCheckRecoveryPasswordRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthImportWebTokenAuthorizationRequestPredict struct {
	APIID        int32
	APIHash      string
	WebAuthToken string
}

func (*AuthImportWebTokenAuthorizationRequestPredict) CRC() uint32 {
	return 0x2db873a9
}

func AuthImportWebTokenAuthorization(ctx context.Context, m Requester, i AuthImportWebTokenAuthorizationRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthRequestFirebaseSmsRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	PhoneNumber        string
	PhoneCodeHash      string
	SafetyNetToken     *string `tl:",omitempty:flags:0"`
	PlayIntegrityToken *string `tl:",omitempty:flags:2"`
	IosPushSecret      *string `tl:",omitempty:flags:1"`
}

func (*AuthRequestFirebaseSmsRequestPredict) CRC() uint32 {
	return 0x8e39261e
}

func AuthRequestFirebaseSms(ctx context.Context, m Requester, i AuthRequestFirebaseSmsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthResetLoginEmailRequestPredict struct {
	PhoneNumber   string
	PhoneCodeHash string
}

func (*AuthResetLoginEmailRequestPredict) CRC() uint32 {
	return 0x7e960193
}

func AuthResetLoginEmail(ctx context.Context, m Requester, i AuthResetLoginEmailRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AuthReportMissingCodeRequestPredict struct {
	PhoneNumber   string
	PhoneCodeHash string
	Mnc           string
}

func (*AuthReportMissingCodeRequestPredict) CRC() uint32 {
	return 0xcb9deff6
}

func AuthReportMissingCode(ctx context.Context, m Requester, i AuthReportMissingCodeRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSendCustomRequestRequestPredict struct {
	CustomMethod string
	Params       DataJSON
}

func (*BotsSendCustomRequestRequestPredict) CRC() uint32 {
	return 0xaa2769ed
}

func BotsSendCustomRequest(ctx context.Context, m Requester, i BotsSendCustomRequestRequestPredict) (DataJSON, error) {
	var res DataJSON
	return res, request(ctx, m, &i, &res)
}

type BotsAnswerWebhookJSONQueryRequestPredict struct {
	QueryID int64
	Data    DataJSON
}

func (*BotsAnswerWebhookJSONQueryRequestPredict) CRC() uint32 {
	return 0xe6213f4d
}

func BotsAnswerWebhookJSONQuery(ctx context.Context, m Requester, i BotsAnswerWebhookJSONQueryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotCommandsRequestPredict struct {
	Scope    BotCommandScope
	LangCode string
	Commands []BotCommand
}

func (*BotsSetBotCommandsRequestPredict) CRC() uint32 {
	return 0x517165a
}

func BotsSetBotCommands(ctx context.Context, m Requester, i BotsSetBotCommandsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsResetBotCommandsRequestPredict struct {
	Scope    BotCommandScope
	LangCode string
}

func (*BotsResetBotCommandsRequestPredict) CRC() uint32 {
	return 0x3d8de0f9
}

func BotsResetBotCommands(ctx context.Context, m Requester, i BotsResetBotCommandsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotCommandsRequestPredict struct {
	Scope    BotCommandScope
	LangCode string
}

func (*BotsGetBotCommandsRequestPredict) CRC() uint32 {
	return 0xe34c0dd6
}

func BotsGetBotCommands(ctx context.Context, m Requester, i BotsGetBotCommandsRequestPredict) ([]BotCommand, error) {
	var res []BotCommand
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotMenuButtonRequestPredict struct {
	UserID InputUser
	Button BotMenuButton
}

func (*BotsSetBotMenuButtonRequestPredict) CRC() uint32 {
	return 0x4504d54f
}

func BotsSetBotMenuButton(ctx context.Context, m Requester, i BotsSetBotMenuButtonRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotMenuButtonRequestPredict struct {
	UserID InputUser
}

func (*BotsGetBotMenuButtonRequestPredict) CRC() uint32 {
	return 0x9c60eb28
}

func BotsGetBotMenuButton(ctx context.Context, m Requester, i BotsGetBotMenuButtonRequestPredict) (BotMenuButton, error) {
	var res BotMenuButton
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotBroadcastDefaultAdminRightsRequestPredict struct {
	AdminRights ChatAdminRights
}

func (*BotsSetBotBroadcastDefaultAdminRightsRequestPredict) CRC() uint32 {
	return 0x788464e1
}

func BotsSetBotBroadcastDefaultAdminRights(ctx context.Context, m Requester, i BotsSetBotBroadcastDefaultAdminRightsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotGroupDefaultAdminRightsRequestPredict struct {
	AdminRights ChatAdminRights
}

func (*BotsSetBotGroupDefaultAdminRightsRequestPredict) CRC() uint32 {
	return 0x925ec9ea
}

func BotsSetBotGroupDefaultAdminRights(ctx context.Context, m Requester, i BotsSetBotGroupDefaultAdminRightsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotInfoRequestPredict struct {
	_           struct{}  `tl:"flags,bitflag"`
	Bot         InputUser `tl:",omitempty:flags:2"`
	LangCode    string
	Name        *string `tl:",omitempty:flags:3"`
	About       *string `tl:",omitempty:flags:0"`
	Description *string `tl:",omitempty:flags:1"`
}

func (*BotsSetBotInfoRequestPredict) CRC() uint32 {
	return 0x10cf3123
}

func BotsSetBotInfo(ctx context.Context, m Requester, i BotsSetBotInfoRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotInfoRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Bot      InputUser `tl:",omitempty:flags:0"`
	LangCode string
}

func (*BotsGetBotInfoRequestPredict) CRC() uint32 {
	return 0xdcd914fd
}

func BotsGetBotInfo(ctx context.Context, m Requester, i BotsGetBotInfoRequestPredict) (BotsBotInfo, error) {
	var res BotsBotInfo
	return res, request(ctx, m, &i, &res)
}

type BotsReorderUsernamesRequestPredict struct {
	Bot   InputUser
	Order []string
}

func (*BotsReorderUsernamesRequestPredict) CRC() uint32 {
	return 0x9709b1c2
}

func BotsReorderUsernames(ctx context.Context, m Requester, i BotsReorderUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsToggleUsernameRequestPredict struct {
	Bot      InputUser
	Username string
	Active   bool
}

func (*BotsToggleUsernameRequestPredict) CRC() uint32 {
	return 0x53ca973
}

func BotsToggleUsername(ctx context.Context, m Requester, i BotsToggleUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsCanSendMessageRequestPredict struct {
	Bot InputUser
}

func (*BotsCanSendMessageRequestPredict) CRC() uint32 {
	return 0x1359f4e6
}

func BotsCanSendMessage(ctx context.Context, m Requester, i BotsCanSendMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsAllowSendMessageRequestPredict struct {
	Bot InputUser
}

func (*BotsAllowSendMessageRequestPredict) CRC() uint32 {
	return 0xf132e3ef
}

func BotsAllowSendMessage(ctx context.Context, m Requester, i BotsAllowSendMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type BotsInvokeWebViewCustomMethodRequestPredict struct {
	Bot          InputUser
	CustomMethod string
	Params       DataJSON
}

func (*BotsInvokeWebViewCustomMethodRequestPredict) CRC() uint32 {
	return 0x87fc5e7
}

func BotsInvokeWebViewCustomMethod(ctx context.Context, m Requester, i BotsInvokeWebViewCustomMethodRequestPredict) (DataJSON, error) {
	var res DataJSON
	return res, request(ctx, m, &i, &res)
}

type BotsGetPopularAppBotsRequestPredict struct {
	Offset string
	Limit  int32
}

func (*BotsGetPopularAppBotsRequestPredict) CRC() uint32 {
	return 0xc2510192
}

func BotsGetPopularAppBots(ctx context.Context, m Requester, i BotsGetPopularAppBotsRequestPredict) (BotsPopularAppBots, error) {
	var res BotsPopularAppBots
	return res, request(ctx, m, &i, &res)
}

type BotsAddPreviewMediaRequestPredict struct {
	Bot      InputUser
	LangCode string
	Media    InputMedia
}

func (*BotsAddPreviewMediaRequestPredict) CRC() uint32 {
	return 0x17aeb75a
}

func BotsAddPreviewMedia(ctx context.Context, m Requester, i BotsAddPreviewMediaRequestPredict) (BotPreviewMedia, error) {
	var res BotPreviewMedia
	return res, request(ctx, m, &i, &res)
}

type BotsEditPreviewMediaRequestPredict struct {
	Bot      InputUser
	LangCode string
	Media    InputMedia
	NewMedia InputMedia
}

func (*BotsEditPreviewMediaRequestPredict) CRC() uint32 {
	return 0x8525606f
}

func BotsEditPreviewMedia(ctx context.Context, m Requester, i BotsEditPreviewMediaRequestPredict) (BotPreviewMedia, error) {
	var res BotPreviewMedia
	return res, request(ctx, m, &i, &res)
}

type BotsDeletePreviewMediaRequestPredict struct {
	Bot      InputUser
	LangCode string
	Media    []InputMedia
}

func (*BotsDeletePreviewMediaRequestPredict) CRC() uint32 {
	return 0x2d0135b3
}

func BotsDeletePreviewMedia(ctx context.Context, m Requester, i BotsDeletePreviewMediaRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsReorderPreviewMediasRequestPredict struct {
	Bot      InputUser
	LangCode string
	Order    []InputMedia
}

func (*BotsReorderPreviewMediasRequestPredict) CRC() uint32 {
	return 0xb627f3aa
}

func BotsReorderPreviewMedias(ctx context.Context, m Requester, i BotsReorderPreviewMediasRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetPreviewInfoRequestPredict struct {
	Bot      InputUser
	LangCode string
}

func (*BotsGetPreviewInfoRequestPredict) CRC() uint32 {
	return 0x423ab3ad
}

func BotsGetPreviewInfo(ctx context.Context, m Requester, i BotsGetPreviewInfoRequestPredict) (BotsPreviewInfo, error) {
	var res BotsPreviewInfo
	return res, request(ctx, m, &i, &res)
}

type BotsGetPreviewMediasRequestPredict struct {
	Bot InputUser
}

func (*BotsGetPreviewMediasRequestPredict) CRC() uint32 {
	return 0xa2a5594d
}

func BotsGetPreviewMedias(ctx context.Context, m Requester, i BotsGetPreviewMediasRequestPredict) ([]BotPreviewMedia, error) {
	var res []BotPreviewMedia
	return res, request(ctx, m, &i, &res)
}

type ChannelsReadHistoryRequestPredict struct {
	Channel InputChannel
	MaxID   int32
}

func (*ChannelsReadHistoryRequestPredict) CRC() uint32 {
	return 0xcc104937
}

func ChannelsReadHistory(ctx context.Context, m Requester, i ChannelsReadHistoryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteMessagesRequestPredict struct {
	Channel InputChannel
	ID      []int32
}

func (*ChannelsDeleteMessagesRequestPredict) CRC() uint32 {
	return 0x84c1fd4e
}

func ChannelsDeleteMessages(ctx context.Context, m Requester, i ChannelsDeleteMessagesRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportSpamRequestPredict struct {
	Channel     InputChannel
	Participant InputPeer
	ID          []int32
}

func (*ChannelsReportSpamRequestPredict) CRC() uint32 {
	return 0xf44a8315
}

func ChannelsReportSpam(ctx context.Context, m Requester, i ChannelsReportSpamRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetMessagesRequestPredict struct {
	Channel InputChannel
	ID      []InputMessage
}

func (*ChannelsGetMessagesRequestPredict) CRC() uint32 {
	return 0xad8c9a23
}

func ChannelsGetMessages(ctx context.Context, m Requester, i ChannelsGetMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetParticipantsRequestPredict struct {
	Channel InputChannel
	Filter  ChannelParticipantsFilter
	Offset  int32
	Limit   int32
	Hash    int64
}

func (*ChannelsGetParticipantsRequestPredict) CRC() uint32 {
	return 0x77ced9d0
}

func ChannelsGetParticipants(ctx context.Context, m Requester, i ChannelsGetParticipantsRequestPredict) (ChannelsChannelParticipants, error) {
	var res ChannelsChannelParticipants
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetParticipantRequestPredict struct {
	Channel     InputChannel
	Participant InputPeer
}

func (*ChannelsGetParticipantRequestPredict) CRC() uint32 {
	return 0xa0ab6cc6
}

func ChannelsGetParticipant(ctx context.Context, m Requester, i ChannelsGetParticipantRequestPredict) (ChannelsChannelParticipant, error) {
	var res ChannelsChannelParticipant
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetChannelsRequestPredict struct {
	ID []InputChannel
}

func (*ChannelsGetChannelsRequestPredict) CRC() uint32 {
	return 0xa7f6bbb
}

func ChannelsGetChannels(ctx context.Context, m Requester, i ChannelsGetChannelsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetFullChannelRequestPredict struct {
	Channel InputChannel
}

func (*ChannelsGetFullChannelRequestPredict) CRC() uint32 {
	return 0x8736a09
}

func ChannelsGetFullChannel(ctx context.Context, m Requester, i ChannelsGetFullChannelRequestPredict) (MessagesChatFull, error) {
	var res MessagesChatFull
	return res, request(ctx, m, &i, &res)
}

type ChannelsCreateChannelRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Broadcast bool     `tl:",omitempty:flags:0,implicit"`
	Megagroup bool     `tl:",omitempty:flags:1,implicit"`
	ForImport bool     `tl:",omitempty:flags:3,implicit"`
	Forum     bool     `tl:",omitempty:flags:5,implicit"`
	Title     string
	About     string
	GeoPoint  InputGeoPoint `tl:",omitempty:flags:2"`
	Address   *string       `tl:",omitempty:flags:2"`
	TTLPeriod *int32        `tl:",omitempty:flags:4"`
}

func (*ChannelsCreateChannelRequestPredict) CRC() uint32 {
	return 0x91006707
}

func ChannelsCreateChannel(ctx context.Context, m Requester, i ChannelsCreateChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditAdminRequestPredict struct {
	Channel     InputChannel
	UserID      InputUser
	AdminRights ChatAdminRights
	Rank        string
}

func (*ChannelsEditAdminRequestPredict) CRC() uint32 {
	return 0xd33c8902
}

func ChannelsEditAdmin(ctx context.Context, m Requester, i ChannelsEditAdminRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditTitleRequestPredict struct {
	Channel InputChannel
	Title   string
}

func (*ChannelsEditTitleRequestPredict) CRC() uint32 {
	return 0x566decd0
}

func ChannelsEditTitle(ctx context.Context, m Requester, i ChannelsEditTitleRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditPhotoRequestPredict struct {
	Channel InputChannel
	Photo   InputChatPhoto
}

func (*ChannelsEditPhotoRequestPredict) CRC() uint32 {
	return 0xf12e57c9
}

func ChannelsEditPhoto(ctx context.Context, m Requester, i ChannelsEditPhotoRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsCheckUsernameRequestPredict struct {
	Channel  InputChannel
	Username string
}

func (*ChannelsCheckUsernameRequestPredict) CRC() uint32 {
	return 0x10e6bd2c
}

func ChannelsCheckUsername(ctx context.Context, m Requester, i ChannelsCheckUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateUsernameRequestPredict struct {
	Channel  InputChannel
	Username string
}

func (*ChannelsUpdateUsernameRequestPredict) CRC() uint32 {
	return 0x3514b3de
}

func ChannelsUpdateUsername(ctx context.Context, m Requester, i ChannelsUpdateUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsJoinChannelRequestPredict struct {
	Channel InputChannel
}

func (*ChannelsJoinChannelRequestPredict) CRC() uint32 {
	return 0x24b524c5
}

func ChannelsJoinChannel(ctx context.Context, m Requester, i ChannelsJoinChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsLeaveChannelRequestPredict struct {
	Channel InputChannel
}

func (*ChannelsLeaveChannelRequestPredict) CRC() uint32 {
	return 0xf836aa95
}

func ChannelsLeaveChannel(ctx context.Context, m Requester, i ChannelsLeaveChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsInviteToChannelRequestPredict struct {
	Channel InputChannel
	Users   []InputUser
}

func (*ChannelsInviteToChannelRequestPredict) CRC() uint32 {
	return 0xc9e33d54
}

func ChannelsInviteToChannel(ctx context.Context, m Requester, i ChannelsInviteToChannelRequestPredict) (MessagesInvitedUsers, error) {
	var res MessagesInvitedUsers
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteChannelRequestPredict struct {
	Channel InputChannel
}

func (*ChannelsDeleteChannelRequestPredict) CRC() uint32 {
	return 0xc0111fe3
}

func ChannelsDeleteChannel(ctx context.Context, m Requester, i ChannelsDeleteChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsExportMessageLinkRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Grouped bool     `tl:",omitempty:flags:0,implicit"`
	Thread  bool     `tl:",omitempty:flags:1,implicit"`
	Channel InputChannel
	ID      int32
}

func (*ChannelsExportMessageLinkRequestPredict) CRC() uint32 {
	return 0xe63fadeb
}

func ChannelsExportMessageLink(ctx context.Context, m Requester, i ChannelsExportMessageLinkRequestPredict) (ExportedMessageLink, error) {
	var res ExportedMessageLink
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleSignaturesRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsToggleSignaturesRequestPredict) CRC() uint32 {
	return 0x1f69b606
}

func ChannelsToggleSignatures(ctx context.Context, m Requester, i ChannelsToggleSignaturesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetAdminedPublicChannelsRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ByLocation  bool     `tl:",omitempty:flags:0,implicit"`
	CheckLimit  bool     `tl:",omitempty:flags:1,implicit"`
	ForPersonal bool     `tl:",omitempty:flags:2,implicit"`
}

func (*ChannelsGetAdminedPublicChannelsRequestPredict) CRC() uint32 {
	return 0xf8b036af
}

func ChannelsGetAdminedPublicChannels(ctx context.Context, m Requester, i ChannelsGetAdminedPublicChannelsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditBannedRequestPredict struct {
	Channel      InputChannel
	Participant  InputPeer
	BannedRights ChatBannedRights
}

func (*ChannelsEditBannedRequestPredict) CRC() uint32 {
	return 0x96e6cd81
}

func ChannelsEditBanned(ctx context.Context, m Requester, i ChannelsEditBannedRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetAdminLogRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Channel      InputChannel
	Q            string
	EventsFilter ChannelAdminLogEventsFilter `tl:",omitempty:flags:0"`
	Admins       []InputUser                 `tl:",omitempty:flags:1"`
	MaxID        int64
	MinID        int64
	Limit        int32
}

func (*ChannelsGetAdminLogRequestPredict) CRC() uint32 {
	return 0x33ddf480
}

func ChannelsGetAdminLog(ctx context.Context, m Requester, i ChannelsGetAdminLogRequestPredict) (ChannelsAdminLogResults, error) {
	var res ChannelsAdminLogResults
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetStickersRequestPredict struct {
	Channel    InputChannel
	Stickerset InputStickerSet
}

func (*ChannelsSetStickersRequestPredict) CRC() uint32 {
	return 0xea8ca4f9
}

func ChannelsSetStickers(ctx context.Context, m Requester, i ChannelsSetStickersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsReadMessageContentsRequestPredict struct {
	Channel InputChannel
	ID      []int32
}

func (*ChannelsReadMessageContentsRequestPredict) CRC() uint32 {
	return 0xeab5dc38
}

func ChannelsReadMessageContents(ctx context.Context, m Requester, i ChannelsReadMessageContentsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteHistoryRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ForEveryone bool     `tl:",omitempty:flags:0,implicit"`
	Channel     InputChannel
	MaxID       int32
}

func (*ChannelsDeleteHistoryRequestPredict) CRC() uint32 {
	return 0x9baa9647
}

func ChannelsDeleteHistory(ctx context.Context, m Requester, i ChannelsDeleteHistoryRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsTogglePreHistoryHiddenRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsTogglePreHistoryHiddenRequestPredict) CRC() uint32 {
	return 0xeabbb94c
}

func ChannelsTogglePreHistoryHidden(ctx context.Context, m Requester, i ChannelsTogglePreHistoryHiddenRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetLeftChannelsRequestPredict struct {
	Offset int32
}

func (*ChannelsGetLeftChannelsRequestPredict) CRC() uint32 {
	return 0x8341ecc0
}

func ChannelsGetLeftChannels(ctx context.Context, m Requester, i ChannelsGetLeftChannelsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetGroupsForDiscussionRequestPredict struct{}

func (*ChannelsGetGroupsForDiscussionRequestPredict) CRC() uint32 {
	return 0xf5dad378
}

func ChannelsGetGroupsForDiscussion(ctx context.Context, m Requester, i ChannelsGetGroupsForDiscussionRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetDiscussionGroupRequestPredict struct {
	Broadcast InputChannel
	Group     InputChannel
}

func (*ChannelsSetDiscussionGroupRequestPredict) CRC() uint32 {
	return 0x40582bb2
}

func ChannelsSetDiscussionGroup(ctx context.Context, m Requester, i ChannelsSetDiscussionGroupRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditCreatorRequestPredict struct {
	Channel  InputChannel
	UserID   InputUser
	Password InputCheckPasswordSRP
}

func (*ChannelsEditCreatorRequestPredict) CRC() uint32 {
	return 0x8f38cd1f
}

func ChannelsEditCreator(ctx context.Context, m Requester, i ChannelsEditCreatorRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditLocationRequestPredict struct {
	Channel  InputChannel
	GeoPoint InputGeoPoint
	Address  string
}

func (*ChannelsEditLocationRequestPredict) CRC() uint32 {
	return 0x58e63f6d
}

func ChannelsEditLocation(ctx context.Context, m Requester, i ChannelsEditLocationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleSlowModeRequestPredict struct {
	Channel InputChannel
	Seconds int32
}

func (*ChannelsToggleSlowModeRequestPredict) CRC() uint32 {
	return 0xedd49ef0
}

func ChannelsToggleSlowMode(ctx context.Context, m Requester, i ChannelsToggleSlowModeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetInactiveChannelsRequestPredict struct{}

func (*ChannelsGetInactiveChannelsRequestPredict) CRC() uint32 {
	return 0x11e831ee
}

func ChannelsGetInactiveChannels(ctx context.Context, m Requester, i ChannelsGetInactiveChannelsRequestPredict) (MessagesInactiveChats, error) {
	var res MessagesInactiveChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsConvertToGigagroupRequestPredict struct {
	Channel InputChannel
}

func (*ChannelsConvertToGigagroupRequestPredict) CRC() uint32 {
	return 0xb290c69
}

func ChannelsConvertToGigagroup(ctx context.Context, m Requester, i ChannelsConvertToGigagroupRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsViewSponsoredMessageRequestPredict struct {
	Channel  InputChannel
	RandomID []byte
}

func (*ChannelsViewSponsoredMessageRequestPredict) CRC() uint32 {
	return 0xbeaedb94
}

func ChannelsViewSponsoredMessage(ctx context.Context, m Requester, i ChannelsViewSponsoredMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetSponsoredMessagesRequestPredict struct {
	Channel InputChannel
}

func (*ChannelsGetSponsoredMessagesRequestPredict) CRC() uint32 {
	return 0xec210fbf
}

func ChannelsGetSponsoredMessages(ctx context.Context, m Requester, i ChannelsGetSponsoredMessagesRequestPredict) (MessagesSponsoredMessages, error) {
	var res MessagesSponsoredMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetSendAsRequestPredict struct {
	Peer InputPeer
}

func (*ChannelsGetSendAsRequestPredict) CRC() uint32 {
	return 0xdc770ee
}

func ChannelsGetSendAs(ctx context.Context, m Requester, i ChannelsGetSendAsRequestPredict) (ChannelsSendAsPeers, error) {
	var res ChannelsSendAsPeers
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteParticipantHistoryRequestPredict struct {
	Channel     InputChannel
	Participant InputPeer
}

func (*ChannelsDeleteParticipantHistoryRequestPredict) CRC() uint32 {
	return 0x367544db
}

func ChannelsDeleteParticipantHistory(ctx context.Context, m Requester, i ChannelsDeleteParticipantHistoryRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleJoinToSendRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsToggleJoinToSendRequestPredict) CRC() uint32 {
	return 0xe4cb9580
}

func ChannelsToggleJoinToSend(ctx context.Context, m Requester, i ChannelsToggleJoinToSendRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleJoinRequestRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsToggleJoinRequestRequestPredict) CRC() uint32 {
	return 0x4c2985b6
}

func ChannelsToggleJoinRequest(ctx context.Context, m Requester, i ChannelsToggleJoinRequestRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsReorderUsernamesRequestPredict struct {
	Channel InputChannel
	Order   []string
}

func (*ChannelsReorderUsernamesRequestPredict) CRC() uint32 {
	return 0xb45ced1d
}

func ChannelsReorderUsernames(ctx context.Context, m Requester, i ChannelsReorderUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleUsernameRequestPredict struct {
	Channel  InputChannel
	Username string
	Active   bool
}

func (*ChannelsToggleUsernameRequestPredict) CRC() uint32 {
	return 0x50f24105
}

func ChannelsToggleUsername(ctx context.Context, m Requester, i ChannelsToggleUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeactivateAllUsernamesRequestPredict struct {
	Channel InputChannel
}

func (*ChannelsDeactivateAllUsernamesRequestPredict) CRC() uint32 {
	return 0xa245dd3
}

func ChannelsDeactivateAllUsernames(ctx context.Context, m Requester, i ChannelsDeactivateAllUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleForumRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsToggleForumRequestPredict) CRC() uint32 {
	return 0xa4298b29
}

func ChannelsToggleForum(ctx context.Context, m Requester, i ChannelsToggleForumRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsCreateForumTopicRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Channel     InputChannel
	Title       string
	IconColor   *int32 `tl:",omitempty:flags:0"`
	IconEmojiID *int64 `tl:",omitempty:flags:3"`
	RandomID    int64
	SendAs      InputPeer `tl:",omitempty:flags:2"`
}

func (*ChannelsCreateForumTopicRequestPredict) CRC() uint32 {
	return 0xf40c0224
}

func ChannelsCreateForumTopic(ctx context.Context, m Requester, i ChannelsCreateForumTopicRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetForumTopicsRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Channel     InputChannel
	Q           *string `tl:",omitempty:flags:0"`
	OffsetDate  int32
	OffsetID    int32
	OffsetTopic int32
	Limit       int32
}

func (*ChannelsGetForumTopicsRequestPredict) CRC() uint32 {
	return 0xde560d1
}

func ChannelsGetForumTopics(ctx context.Context, m Requester, i ChannelsGetForumTopicsRequestPredict) (MessagesForumTopics, error) {
	var res MessagesForumTopics
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetForumTopicsByIDRequestPredict struct {
	Channel InputChannel
	Topics  []int32
}

func (*ChannelsGetForumTopicsByIDRequestPredict) CRC() uint32 {
	return 0xb0831eb9
}

func ChannelsGetForumTopicsByID(ctx context.Context, m Requester, i ChannelsGetForumTopicsByIDRequestPredict) (MessagesForumTopics, error) {
	var res MessagesForumTopics
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditForumTopicRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Channel     InputChannel
	TopicID     int32
	Title       *string `tl:",omitempty:flags:0"`
	IconEmojiID *int64  `tl:",omitempty:flags:1"`
	Closed      *bool   `tl:",omitempty:flags:2"`
	Hidden      *bool   `tl:",omitempty:flags:3"`
}

func (*ChannelsEditForumTopicRequestPredict) CRC() uint32 {
	return 0xf4dfa185
}

func ChannelsEditForumTopic(ctx context.Context, m Requester, i ChannelsEditForumTopicRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdatePinnedForumTopicRequestPredict struct {
	Channel InputChannel
	TopicID int32
	Pinned  bool
}

func (*ChannelsUpdatePinnedForumTopicRequestPredict) CRC() uint32 {
	return 0x6c2d9026
}

func ChannelsUpdatePinnedForumTopic(ctx context.Context, m Requester, i ChannelsUpdatePinnedForumTopicRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteTopicHistoryRequestPredict struct {
	Channel  InputChannel
	TopMsgID int32
}

func (*ChannelsDeleteTopicHistoryRequestPredict) CRC() uint32 {
	return 0x34435f2d
}

func ChannelsDeleteTopicHistory(ctx context.Context, m Requester, i ChannelsDeleteTopicHistoryRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type ChannelsReorderPinnedForumTopicsRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Force   bool     `tl:",omitempty:flags:0,implicit"`
	Channel InputChannel
	Order   []int32
}

func (*ChannelsReorderPinnedForumTopicsRequestPredict) CRC() uint32 {
	return 0x2950a18f
}

func ChannelsReorderPinnedForumTopics(ctx context.Context, m Requester, i ChannelsReorderPinnedForumTopicsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleAntiSpamRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsToggleAntiSpamRequestPredict) CRC() uint32 {
	return 0x68f3e4eb
}

func ChannelsToggleAntiSpam(ctx context.Context, m Requester, i ChannelsToggleAntiSpamRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportAntiSpamFalsePositiveRequestPredict struct {
	Channel InputChannel
	MsgID   int32
}

func (*ChannelsReportAntiSpamFalsePositiveRequestPredict) CRC() uint32 {
	return 0xa850a693
}

func ChannelsReportAntiSpamFalsePositive(ctx context.Context, m Requester, i ChannelsReportAntiSpamFalsePositiveRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleParticipantsHiddenRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsToggleParticipantsHiddenRequestPredict) CRC() uint32 {
	return 0x6a6e7854
}

func ChannelsToggleParticipantsHidden(ctx context.Context, m Requester, i ChannelsToggleParticipantsHiddenRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsClickSponsoredMessageRequestPredict struct {
	Channel  InputChannel
	RandomID []byte
}

func (*ChannelsClickSponsoredMessageRequestPredict) CRC() uint32 {
	return 0x18afbc93
}

func ChannelsClickSponsoredMessage(ctx context.Context, m Requester, i ChannelsClickSponsoredMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateColorRequestPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	ForProfile        bool     `tl:",omitempty:flags:1,implicit"`
	Channel           InputChannel
	Color             *int32 `tl:",omitempty:flags:2"`
	BackgroundEmojiID *int64 `tl:",omitempty:flags:0"`
}

func (*ChannelsUpdateColorRequestPredict) CRC() uint32 {
	return 0xd8aa3671
}

func ChannelsUpdateColor(ctx context.Context, m Requester, i ChannelsUpdateColorRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleViewForumAsMessagesRequestPredict struct {
	Channel InputChannel
	Enabled bool
}

func (*ChannelsToggleViewForumAsMessagesRequestPredict) CRC() uint32 {
	return 0x9738bb15
}

func ChannelsToggleViewForumAsMessages(ctx context.Context, m Requester, i ChannelsToggleViewForumAsMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetChannelRecommendationsRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Channel InputChannel `tl:",omitempty:flags:0"`
}

func (*ChannelsGetChannelRecommendationsRequestPredict) CRC() uint32 {
	return 0x25a71742
}

func ChannelsGetChannelRecommendations(ctx context.Context, m Requester, i ChannelsGetChannelRecommendationsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateEmojiStatusRequestPredict struct {
	Channel     InputChannel
	EmojiStatus EmojiStatus
}

func (*ChannelsUpdateEmojiStatusRequestPredict) CRC() uint32 {
	return 0xf0d3e6a8
}

func ChannelsUpdateEmojiStatus(ctx context.Context, m Requester, i ChannelsUpdateEmojiStatusRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetBoostsToUnblockRestrictionsRequestPredict struct {
	Channel InputChannel
	Boosts  int32
}

func (*ChannelsSetBoostsToUnblockRestrictionsRequestPredict) CRC() uint32 {
	return 0xad399cee
}

func ChannelsSetBoostsToUnblockRestrictions(ctx context.Context, m Requester, i ChannelsSetBoostsToUnblockRestrictionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetEmojiStickersRequestPredict struct {
	Channel    InputChannel
	Stickerset InputStickerSet
}

func (*ChannelsSetEmojiStickersRequestPredict) CRC() uint32 {
	return 0x3cd930b7
}

func ChannelsSetEmojiStickers(ctx context.Context, m Requester, i ChannelsSetEmojiStickersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportSponsoredMessageRequestPredict struct {
	Channel  InputChannel
	RandomID []byte
	Option   []byte
}

func (*ChannelsReportSponsoredMessageRequestPredict) CRC() uint32 {
	return 0xaf8ff6b9
}

func ChannelsReportSponsoredMessage(ctx context.Context, m Requester, i ChannelsReportSponsoredMessageRequestPredict) (ChannelsSponsoredMessageReportResult, error) {
	var res ChannelsSponsoredMessageReportResult
	return res, request(ctx, m, &i, &res)
}

type ChannelsRestrictSponsoredMessagesRequestPredict struct {
	Channel    InputChannel
	Restricted bool
}

func (*ChannelsRestrictSponsoredMessagesRequestPredict) CRC() uint32 {
	return 0x9ae91519
}

func ChannelsRestrictSponsoredMessages(ctx context.Context, m Requester, i ChannelsRestrictSponsoredMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsSearchPostsRequestPredict struct {
	Hashtag    string
	OffsetRate int32
	OffsetPeer InputPeer
	OffsetID   int32
	Limit      int32
}

func (*ChannelsSearchPostsRequestPredict) CRC() uint32 {
	return 0xd19f987b
}

func ChannelsSearchPosts(ctx context.Context, m Requester, i ChannelsSearchPostsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type ChatlistsExportChatlistInviteRequestPredict struct {
	Chatlist InputChatlist
	Title    string
	Peers    []InputPeer
}

func (*ChatlistsExportChatlistInviteRequestPredict) CRC() uint32 {
	return 0x8472478e
}

func ChatlistsExportChatlistInvite(ctx context.Context, m Requester, i ChatlistsExportChatlistInviteRequestPredict) (ChatlistsExportedChatlistInvite, error) {
	var res ChatlistsExportedChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsDeleteExportedInviteRequestPredict struct {
	Chatlist InputChatlist
	Slug     string
}

func (*ChatlistsDeleteExportedInviteRequestPredict) CRC() uint32 {
	return 0x719c5c5e
}

func ChatlistsDeleteExportedInvite(ctx context.Context, m Requester, i ChatlistsDeleteExportedInviteRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChatlistsEditExportedInviteRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Chatlist InputChatlist
	Slug     string
	Title    *string     `tl:",omitempty:flags:1"`
	Peers    []InputPeer `tl:",omitempty:flags:2"`
}

func (*ChatlistsEditExportedInviteRequestPredict) CRC() uint32 {
	return 0x653db63d
}

func ChatlistsEditExportedInvite(ctx context.Context, m Requester, i ChatlistsEditExportedInviteRequestPredict) (ExportedChatlistInvite, error) {
	var res ExportedChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetExportedInvitesRequestPredict struct {
	Chatlist InputChatlist
}

func (*ChatlistsGetExportedInvitesRequestPredict) CRC() uint32 {
	return 0xce03da83
}

func ChatlistsGetExportedInvites(ctx context.Context, m Requester, i ChatlistsGetExportedInvitesRequestPredict) (ChatlistsExportedInvites, error) {
	var res ChatlistsExportedInvites
	return res, request(ctx, m, &i, &res)
}

type ChatlistsCheckChatlistInviteRequestPredict struct {
	Slug string
}

func (*ChatlistsCheckChatlistInviteRequestPredict) CRC() uint32 {
	return 0x41c10fff
}

func ChatlistsCheckChatlistInvite(ctx context.Context, m Requester, i ChatlistsCheckChatlistInviteRequestPredict) (ChatlistsChatlistInvite, error) {
	var res ChatlistsChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsJoinChatlistInviteRequestPredict struct {
	Slug  string
	Peers []InputPeer
}

func (*ChatlistsJoinChatlistInviteRequestPredict) CRC() uint32 {
	return 0xa6b1e39a
}

func ChatlistsJoinChatlistInvite(ctx context.Context, m Requester, i ChatlistsJoinChatlistInviteRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetChatlistUpdatesRequestPredict struct {
	Chatlist InputChatlist
}

func (*ChatlistsGetChatlistUpdatesRequestPredict) CRC() uint32 {
	return 0x89419521
}

func ChatlistsGetChatlistUpdates(ctx context.Context, m Requester, i ChatlistsGetChatlistUpdatesRequestPredict) (ChatlistsChatlistUpdates, error) {
	var res ChatlistsChatlistUpdates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsJoinChatlistUpdatesRequestPredict struct {
	Chatlist InputChatlist
	Peers    []InputPeer
}

func (*ChatlistsJoinChatlistUpdatesRequestPredict) CRC() uint32 {
	return 0xe089f8f5
}

func ChatlistsJoinChatlistUpdates(ctx context.Context, m Requester, i ChatlistsJoinChatlistUpdatesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsHideChatlistUpdatesRequestPredict struct {
	Chatlist InputChatlist
}

func (*ChatlistsHideChatlistUpdatesRequestPredict) CRC() uint32 {
	return 0x66e486fb
}

func ChatlistsHideChatlistUpdates(ctx context.Context, m Requester, i ChatlistsHideChatlistUpdatesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetLeaveChatlistSuggestionsRequestPredict struct {
	Chatlist InputChatlist
}

func (*ChatlistsGetLeaveChatlistSuggestionsRequestPredict) CRC() uint32 {
	return 0xfdbcd714
}

func ChatlistsGetLeaveChatlistSuggestions(ctx context.Context, m Requester, i ChatlistsGetLeaveChatlistSuggestionsRequestPredict) ([]Peer, error) {
	var res []Peer
	return res, request(ctx, m, &i, &res)
}

type ChatlistsLeaveChatlistRequestPredict struct {
	Chatlist InputChatlist
	Peers    []InputPeer
}

func (*ChatlistsLeaveChatlistRequestPredict) CRC() uint32 {
	return 0x74fae13a
}

func ChatlistsLeaveChatlist(ctx context.Context, m Requester, i ChatlistsLeaveChatlistRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsGetContactIDsRequestPredict struct {
	Hash int64
}

func (*ContactsGetContactIDsRequestPredict) CRC() uint32 {
	return 0x7adc669d
}

func ContactsGetContactIDs(ctx context.Context, m Requester, i ContactsGetContactIDsRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type ContactsGetStatusesRequestPredict struct{}

func (*ContactsGetStatusesRequestPredict) CRC() uint32 {
	return 0xc4a353ee
}

func ContactsGetStatuses(ctx context.Context, m Requester, i ContactsGetStatusesRequestPredict) ([]ContactStatus, error) {
	var res []ContactStatus
	return res, request(ctx, m, &i, &res)
}

type ContactsGetContactsRequestPredict struct {
	Hash int64
}

func (*ContactsGetContactsRequestPredict) CRC() uint32 {
	return 0x5dd69e12
}

func ContactsGetContacts(ctx context.Context, m Requester, i ContactsGetContactsRequestPredict) (ContactsContacts, error) {
	var res ContactsContacts
	return res, request(ctx, m, &i, &res)
}

type ContactsImportContactsRequestPredict struct {
	Contacts []InputContact
}

func (*ContactsImportContactsRequestPredict) CRC() uint32 {
	return 0x2c800be5
}

func ContactsImportContacts(ctx context.Context, m Requester, i ContactsImportContactsRequestPredict) (ContactsImportedContacts, error) {
	var res ContactsImportedContacts
	return res, request(ctx, m, &i, &res)
}

type ContactsDeleteContactsRequestPredict struct {
	ID []InputUser
}

func (*ContactsDeleteContactsRequestPredict) CRC() uint32 {
	return 0x96a0e00
}

func ContactsDeleteContacts(ctx context.Context, m Requester, i ContactsDeleteContactsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsDeleteByPhonesRequestPredict struct {
	Phones []string
}

func (*ContactsDeleteByPhonesRequestPredict) CRC() uint32 {
	return 0x1013fd9e
}

func ContactsDeleteByPhones(ctx context.Context, m Requester, i ContactsDeleteByPhonesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsBlockRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	ID            InputPeer
}

func (*ContactsBlockRequestPredict) CRC() uint32 {
	return 0x2e2e8734
}

func ContactsBlock(ctx context.Context, m Requester, i ContactsBlockRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsUnblockRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	ID            InputPeer
}

func (*ContactsUnblockRequestPredict) CRC() uint32 {
	return 0xb550d328
}

func ContactsUnblock(ctx context.Context, m Requester, i ContactsUnblockRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsGetBlockedRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	Offset        int32
	Limit         int32
}

func (*ContactsGetBlockedRequestPredict) CRC() uint32 {
	return 0x9a868f80
}

func ContactsGetBlocked(ctx context.Context, m Requester, i ContactsGetBlockedRequestPredict) (ContactsBlocked, error) {
	var res ContactsBlocked
	return res, request(ctx, m, &i, &res)
}

type ContactsSearchRequestPredict struct {
	Q     string
	Limit int32
}

func (*ContactsSearchRequestPredict) CRC() uint32 {
	return 0x11f812d8
}

func ContactsSearch(ctx context.Context, m Requester, i ContactsSearchRequestPredict) (ContactsFound, error) {
	var res ContactsFound
	return res, request(ctx, m, &i, &res)
}

type ContactsResolveUsernameRequestPredict struct {
	Username string
}

func (*ContactsResolveUsernameRequestPredict) CRC() uint32 {
	return 0xf93ccba3
}

func ContactsResolveUsername(ctx context.Context, m Requester, i ContactsResolveUsernameRequestPredict) (ContactsResolvedPeer, error) {
	var res ContactsResolvedPeer
	return res, request(ctx, m, &i, &res)
}

type ContactsGetTopPeersRequestPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Correspondents bool     `tl:",omitempty:flags:0,implicit"`
	BotsPm         bool     `tl:",omitempty:flags:1,implicit"`
	BotsInline     bool     `tl:",omitempty:flags:2,implicit"`
	PhoneCalls     bool     `tl:",omitempty:flags:3,implicit"`
	ForwardUsers   bool     `tl:",omitempty:flags:4,implicit"`
	ForwardChats   bool     `tl:",omitempty:flags:5,implicit"`
	Groups         bool     `tl:",omitempty:flags:10,implicit"`
	Channels       bool     `tl:",omitempty:flags:15,implicit"`
	BotsApp        bool     `tl:",omitempty:flags:16,implicit"`
	Offset         int32
	Limit          int32
	Hash           int64
}

func (*ContactsGetTopPeersRequestPredict) CRC() uint32 {
	return 0x973478b6
}

func ContactsGetTopPeers(ctx context.Context, m Requester, i ContactsGetTopPeersRequestPredict) (ContactsTopPeers, error) {
	var res ContactsTopPeers
	return res, request(ctx, m, &i, &res)
}

type ContactsResetTopPeerRatingRequestPredict struct {
	Category TopPeerCategory
	Peer     InputPeer
}

func (*ContactsResetTopPeerRatingRequestPredict) CRC() uint32 {
	return 0x1ae373ac
}

func ContactsResetTopPeerRating(ctx context.Context, m Requester, i ContactsResetTopPeerRatingRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsResetSavedRequestPredict struct{}

func (*ContactsResetSavedRequestPredict) CRC() uint32 {
	return 0x879537f1
}

func ContactsResetSaved(ctx context.Context, m Requester, i ContactsResetSavedRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsGetSavedRequestPredict struct{}

func (*ContactsGetSavedRequestPredict) CRC() uint32 {
	return 0x82f1e39f
}

func ContactsGetSaved(ctx context.Context, m Requester, i ContactsGetSavedRequestPredict) ([]SavedContact, error) {
	var res []SavedContact
	return res, request(ctx, m, &i, &res)
}

type ContactsToggleTopPeersRequestPredict struct {
	Enabled bool
}

func (*ContactsToggleTopPeersRequestPredict) CRC() uint32 {
	return 0x8514bdda
}

func ContactsToggleTopPeers(ctx context.Context, m Requester, i ContactsToggleTopPeersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsAddContactRequestPredict struct {
	_                        struct{} `tl:"flags,bitflag"`
	AddPhonePrivacyException bool     `tl:",omitempty:flags:0,implicit"`
	ID                       InputUser
	FirstName                string
	LastName                 string
	Phone                    string
}

func (*ContactsAddContactRequestPredict) CRC() uint32 {
	return 0xe8f463d0
}

func ContactsAddContact(ctx context.Context, m Requester, i ContactsAddContactRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsAcceptContactRequestPredict struct {
	ID InputUser
}

func (*ContactsAcceptContactRequestPredict) CRC() uint32 {
	return 0xf831a20f
}

func ContactsAcceptContact(ctx context.Context, m Requester, i ContactsAcceptContactRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsGetLocatedRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Background  bool     `tl:",omitempty:flags:1,implicit"`
	GeoPoint    InputGeoPoint
	SelfExpires *int32 `tl:",omitempty:flags:0"`
}

func (*ContactsGetLocatedRequestPredict) CRC() uint32 {
	return 0xd348bc44
}

func ContactsGetLocated(ctx context.Context, m Requester, i ContactsGetLocatedRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsBlockFromRepliesRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	DeleteMessage bool     `tl:",omitempty:flags:0,implicit"`
	DeleteHistory bool     `tl:",omitempty:flags:1,implicit"`
	ReportSpam    bool     `tl:",omitempty:flags:2,implicit"`
	MsgID         int32
}

func (*ContactsBlockFromRepliesRequestPredict) CRC() uint32 {
	return 0x29a8962c
}

func ContactsBlockFromReplies(ctx context.Context, m Requester, i ContactsBlockFromRepliesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsResolvePhoneRequestPredict struct {
	Phone string
}

func (*ContactsResolvePhoneRequestPredict) CRC() uint32 {
	return 0x8af94344
}

func ContactsResolvePhone(ctx context.Context, m Requester, i ContactsResolvePhoneRequestPredict) (ContactsResolvedPeer, error) {
	var res ContactsResolvedPeer
	return res, request(ctx, m, &i, &res)
}

type ContactsExportContactTokenRequestPredict struct{}

func (*ContactsExportContactTokenRequestPredict) CRC() uint32 {
	return 0xf8654027
}

func ContactsExportContactToken(ctx context.Context, m Requester, i ContactsExportContactTokenRequestPredict) (ExportedContactToken, error) {
	var res ExportedContactToken
	return res, request(ctx, m, &i, &res)
}

type ContactsImportContactTokenRequestPredict struct {
	Token string
}

func (*ContactsImportContactTokenRequestPredict) CRC() uint32 {
	return 0x13005788
}

func ContactsImportContactToken(ctx context.Context, m Requester, i ContactsImportContactTokenRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type ContactsEditCloseFriendsRequestPredict struct {
	ID []int64
}

func (*ContactsEditCloseFriendsRequestPredict) CRC() uint32 {
	return 0xba6705f0
}

func ContactsEditCloseFriends(ctx context.Context, m Requester, i ContactsEditCloseFriendsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsSetBlockedRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	MyStoriesFrom bool     `tl:",omitempty:flags:0,implicit"`
	ID            []InputPeer
	Limit         int32
}

func (*ContactsSetBlockedRequestPredict) CRC() uint32 {
	return 0x94c65c76
}

func ContactsSetBlocked(ctx context.Context, m Requester, i ContactsSetBlockedRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsGetBirthdaysRequestPredict struct{}

func (*ContactsGetBirthdaysRequestPredict) CRC() uint32 {
	return 0xdaeda864
}

func ContactsGetBirthdays(ctx context.Context, m Requester, i ContactsGetBirthdaysRequestPredict) (ContactsContactBirthdays, error) {
	var res ContactsContactBirthdays
	return res, request(ctx, m, &i, &res)
}

type FoldersEditPeerFoldersRequestPredict struct {
	FolderPeers []InputFolderPeer
}

func (*FoldersEditPeerFoldersRequestPredict) CRC() uint32 {
	return 0x6847d0ab
}

func FoldersEditPeerFolders(ctx context.Context, m Requester, i FoldersEditPeerFoldersRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type FragmentGetCollectibleInfoRequestPredict struct {
	Collectible InputCollectible
}

func (*FragmentGetCollectibleInfoRequestPredict) CRC() uint32 {
	return 0xbe1e85ba
}

func FragmentGetCollectibleInfo(ctx context.Context, m Requester, i FragmentGetCollectibleInfoRequestPredict) (FragmentCollectibleInfo, error) {
	var res FragmentCollectibleInfo
	return res, request(ctx, m, &i, &res)
}

type HelpGetConfigRequestPredict struct{}

func (*HelpGetConfigRequestPredict) CRC() uint32 {
	return 0xc4f9186b
}

func HelpGetConfig(ctx context.Context, m Requester, i HelpGetConfigRequestPredict) (Config, error) {
	var res Config
	return res, request(ctx, m, &i, &res)
}

type HelpGetNearestDcRequestPredict struct{}

func (*HelpGetNearestDcRequestPredict) CRC() uint32 {
	return 0x1fb33026
}

func HelpGetNearestDc(ctx context.Context, m Requester, i HelpGetNearestDcRequestPredict) (NearestDc, error) {
	var res NearestDc
	return res, request(ctx, m, &i, &res)
}

type HelpGetAppUpdateRequestPredict struct {
	Source string
}

func (*HelpGetAppUpdateRequestPredict) CRC() uint32 {
	return 0x522d5a7d
}

func HelpGetAppUpdate(ctx context.Context, m Requester, i HelpGetAppUpdateRequestPredict) (HelpAppUpdate, error) {
	var res HelpAppUpdate
	return res, request(ctx, m, &i, &res)
}

type HelpGetInviteTextRequestPredict struct{}

func (*HelpGetInviteTextRequestPredict) CRC() uint32 {
	return 0x4d392343
}

func HelpGetInviteText(ctx context.Context, m Requester, i HelpGetInviteTextRequestPredict) (HelpInviteText, error) {
	var res HelpInviteText
	return res, request(ctx, m, &i, &res)
}

type HelpGetSupportRequestPredict struct{}

func (*HelpGetSupportRequestPredict) CRC() uint32 {
	return 0x9cdf08cd
}

func HelpGetSupport(ctx context.Context, m Requester, i HelpGetSupportRequestPredict) (HelpSupport, error) {
	var res HelpSupport
	return res, request(ctx, m, &i, &res)
}

type HelpSetBotUpdatesStatusRequestPredict struct {
	PendingUpdatesCount int32
	Message             string
}

func (*HelpSetBotUpdatesStatusRequestPredict) CRC() uint32 {
	return 0xec22cfcd
}

func HelpSetBotUpdatesStatus(ctx context.Context, m Requester, i HelpSetBotUpdatesStatusRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetCdnConfigRequestPredict struct{}

func (*HelpGetCdnConfigRequestPredict) CRC() uint32 {
	return 0x52029342
}

func HelpGetCdnConfig(ctx context.Context, m Requester, i HelpGetCdnConfigRequestPredict) (CdnConfig, error) {
	var res CdnConfig
	return res, request(ctx, m, &i, &res)
}

type HelpGetRecentMeUrlsRequestPredict struct {
	Referer string
}

func (*HelpGetRecentMeUrlsRequestPredict) CRC() uint32 {
	return 0x3dc0f114
}

func HelpGetRecentMeUrls(ctx context.Context, m Requester, i HelpGetRecentMeUrlsRequestPredict) (HelpRecentMeUrls, error) {
	var res HelpRecentMeUrls
	return res, request(ctx, m, &i, &res)
}

type HelpGetTermsOfServiceUpdateRequestPredict struct{}

func (*HelpGetTermsOfServiceUpdateRequestPredict) CRC() uint32 {
	return 0x2ca51fd1
}

func HelpGetTermsOfServiceUpdate(ctx context.Context, m Requester, i HelpGetTermsOfServiceUpdateRequestPredict) (HelpTermsOfServiceUpdate, error) {
	var res HelpTermsOfServiceUpdate
	return res, request(ctx, m, &i, &res)
}

type HelpAcceptTermsOfServiceRequestPredict struct {
	ID DataJSON
}

func (*HelpAcceptTermsOfServiceRequestPredict) CRC() uint32 {
	return 0xee72f79a
}

func HelpAcceptTermsOfService(ctx context.Context, m Requester, i HelpAcceptTermsOfServiceRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetDeepLinkInfoRequestPredict struct {
	Path string
}

func (*HelpGetDeepLinkInfoRequestPredict) CRC() uint32 {
	return 0x3fedc75f
}

func HelpGetDeepLinkInfo(ctx context.Context, m Requester, i HelpGetDeepLinkInfoRequestPredict) (HelpDeepLinkInfo, error) {
	var res HelpDeepLinkInfo
	return res, request(ctx, m, &i, &res)
}

type HelpGetAppConfigRequestPredict struct {
	Hash int32
}

func (*HelpGetAppConfigRequestPredict) CRC() uint32 {
	return 0x61e3f854
}

func HelpGetAppConfig(ctx context.Context, m Requester, i HelpGetAppConfigRequestPredict) (HelpAppConfig, error) {
	var res HelpAppConfig
	return res, request(ctx, m, &i, &res)
}

type HelpSaveAppLogRequestPredict struct {
	Events []InputAppEvent
}

func (*HelpSaveAppLogRequestPredict) CRC() uint32 {
	return 0x6f02f748
}

func HelpSaveAppLog(ctx context.Context, m Requester, i HelpSaveAppLogRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetPassportConfigRequestPredict struct {
	Hash int32
}

func (*HelpGetPassportConfigRequestPredict) CRC() uint32 {
	return 0xc661ad08
}

func HelpGetPassportConfig(ctx context.Context, m Requester, i HelpGetPassportConfigRequestPredict) (HelpPassportConfig, error) {
	var res HelpPassportConfig
	return res, request(ctx, m, &i, &res)
}

type HelpGetSupportNameRequestPredict struct{}

func (*HelpGetSupportNameRequestPredict) CRC() uint32 {
	return 0xd360e72c
}

func HelpGetSupportName(ctx context.Context, m Requester, i HelpGetSupportNameRequestPredict) (HelpSupportName, error) {
	var res HelpSupportName
	return res, request(ctx, m, &i, &res)
}

type HelpGetUserInfoRequestPredict struct {
	UserID InputUser
}

func (*HelpGetUserInfoRequestPredict) CRC() uint32 {
	return 0x38a08d3
}

func HelpGetUserInfo(ctx context.Context, m Requester, i HelpGetUserInfoRequestPredict) (HelpUserInfo, error) {
	var res HelpUserInfo
	return res, request(ctx, m, &i, &res)
}

type HelpEditUserInfoRequestPredict struct {
	UserID   InputUser
	Message  string
	Entities []MessageEntity
}

func (*HelpEditUserInfoRequestPredict) CRC() uint32 {
	return 0x66b91b70
}

func HelpEditUserInfo(ctx context.Context, m Requester, i HelpEditUserInfoRequestPredict) (HelpUserInfo, error) {
	var res HelpUserInfo
	return res, request(ctx, m, &i, &res)
}

type HelpGetPromoDataRequestPredict struct{}

func (*HelpGetPromoDataRequestPredict) CRC() uint32 {
	return 0xc0977421
}

func HelpGetPromoData(ctx context.Context, m Requester, i HelpGetPromoDataRequestPredict) (HelpPromoData, error) {
	var res HelpPromoData
	return res, request(ctx, m, &i, &res)
}

type HelpHidePromoDataRequestPredict struct {
	Peer InputPeer
}

func (*HelpHidePromoDataRequestPredict) CRC() uint32 {
	return 0x1e251c95
}

func HelpHidePromoData(ctx context.Context, m Requester, i HelpHidePromoDataRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpDismissSuggestionRequestPredict struct {
	Peer       InputPeer
	Suggestion string
}

func (*HelpDismissSuggestionRequestPredict) CRC() uint32 {
	return 0xf50dbaa1
}

func HelpDismissSuggestion(ctx context.Context, m Requester, i HelpDismissSuggestionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetCountriesListRequestPredict struct {
	LangCode string
	Hash     int32
}

func (*HelpGetCountriesListRequestPredict) CRC() uint32 {
	return 0x735787a8
}

func HelpGetCountriesList(ctx context.Context, m Requester, i HelpGetCountriesListRequestPredict) (HelpCountriesList, error) {
	var res HelpCountriesList
	return res, request(ctx, m, &i, &res)
}

type HelpGetPremiumPromoRequestPredict struct{}

func (*HelpGetPremiumPromoRequestPredict) CRC() uint32 {
	return 0xb81b93d4
}

func HelpGetPremiumPromo(ctx context.Context, m Requester, i HelpGetPremiumPromoRequestPredict) (HelpPremiumPromo, error) {
	var res HelpPremiumPromo
	return res, request(ctx, m, &i, &res)
}

type HelpGetPeerColorsRequestPredict struct {
	Hash int32
}

func (*HelpGetPeerColorsRequestPredict) CRC() uint32 {
	return 0xda80f42f
}

func HelpGetPeerColors(ctx context.Context, m Requester, i HelpGetPeerColorsRequestPredict) (HelpPeerColors, error) {
	var res HelpPeerColors
	return res, request(ctx, m, &i, &res)
}

type HelpGetPeerProfileColorsRequestPredict struct {
	Hash int32
}

func (*HelpGetPeerProfileColorsRequestPredict) CRC() uint32 {
	return 0xabcfa9fd
}

func HelpGetPeerProfileColors(ctx context.Context, m Requester, i HelpGetPeerProfileColorsRequestPredict) (HelpPeerColors, error) {
	var res HelpPeerColors
	return res, request(ctx, m, &i, &res)
}

type HelpGetTimezonesListRequestPredict struct {
	Hash int32
}

func (*HelpGetTimezonesListRequestPredict) CRC() uint32 {
	return 0x49b30240
}

func HelpGetTimezonesList(ctx context.Context, m Requester, i HelpGetTimezonesListRequestPredict) (HelpTimezonesList, error) {
	var res HelpTimezonesList
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLangPackRequestPredict struct {
	LangPack string
	LangCode string
}

func (*LangpackGetLangPackRequestPredict) CRC() uint32 {
	return 0xf2f2330a
}

func LangpackGetLangPack(ctx context.Context, m Requester, i LangpackGetLangPackRequestPredict) (LangPackDifference, error) {
	var res LangPackDifference
	return res, request(ctx, m, &i, &res)
}

type LangpackGetStringsRequestPredict struct {
	LangPack string
	LangCode string
	Keys     []string
}

func (*LangpackGetStringsRequestPredict) CRC() uint32 {
	return 0xefea3803
}

func LangpackGetStrings(ctx context.Context, m Requester, i LangpackGetStringsRequestPredict) ([]LangPackString, error) {
	var res []LangPackString
	return res, request(ctx, m, &i, &res)
}

type LangpackGetDifferenceRequestPredict struct {
	LangPack    string
	LangCode    string
	FromVersion int32
}

func (*LangpackGetDifferenceRequestPredict) CRC() uint32 {
	return 0xcd984aa5
}

func LangpackGetDifference(ctx context.Context, m Requester, i LangpackGetDifferenceRequestPredict) (LangPackDifference, error) {
	var res LangPackDifference
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLanguagesRequestPredict struct {
	LangPack string
}

func (*LangpackGetLanguagesRequestPredict) CRC() uint32 {
	return 0x42c6978f
}

func LangpackGetLanguages(ctx context.Context, m Requester, i LangpackGetLanguagesRequestPredict) ([]LangPackLanguage, error) {
	var res []LangPackLanguage
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLanguageRequestPredict struct {
	LangPack string
	LangCode string
}

func (*LangpackGetLanguageRequestPredict) CRC() uint32 {
	return 0x6a596502
}

func LangpackGetLanguage(ctx context.Context, m Requester, i LangpackGetLanguageRequestPredict) (LangPackLanguage, error) {
	var res LangPackLanguage
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesRequestPredict struct {
	ID []InputMessage
}

func (*MessagesGetMessagesRequestPredict) CRC() uint32 {
	return 0x63c66506
}

func MessagesGetMessages(ctx context.Context, m Requester, i MessagesGetMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDialogsRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ExcludePinned bool     `tl:",omitempty:flags:0,implicit"`
	FolderID      *int32   `tl:",omitempty:flags:1"`
	OffsetDate    int32
	OffsetID      int32
	OffsetPeer    InputPeer
	Limit         int32
	Hash          int64
}

func (*MessagesGetDialogsRequestPredict) CRC() uint32 {
	return 0xa0f4cb4f
}

func MessagesGetDialogs(ctx context.Context, m Requester, i MessagesGetDialogsRequestPredict) (MessagesDialogs, error) {
	var res MessagesDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesGetHistoryRequestPredict struct {
	Peer       InputPeer
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int64
}

func (*MessagesGetHistoryRequestPredict) CRC() uint32 {
	return 0x4423e6c5
}

func MessagesGetHistory(ctx context.Context, m Requester, i MessagesGetHistoryRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Peer          InputPeer
	Q             string
	FromID        InputPeer  `tl:",omitempty:flags:0"`
	SavedPeerID   InputPeer  `tl:",omitempty:flags:2"`
	SavedReaction []Reaction `tl:",omitempty:flags:3"`
	TopMsgID      *int32     `tl:",omitempty:flags:1"`
	Filter        MessagesFilter
	MinDate       int32
	MaxDate       int32
	OffsetID      int32
	AddOffset     int32
	Limit         int32
	MaxID         int32
	MinID         int32
	Hash          int64
}

func (*MessagesSearchRequestPredict) CRC() uint32 {
	return 0x29ee847a
}

func MessagesSearch(ctx context.Context, m Requester, i MessagesSearchRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadHistoryRequestPredict struct {
	Peer  InputPeer
	MaxID int32
}

func (*MessagesReadHistoryRequestPredict) CRC() uint32 {
	return 0xe306d3a
}

func MessagesReadHistory(ctx context.Context, m Requester, i MessagesReadHistoryRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteHistoryRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	JustClear bool     `tl:",omitempty:flags:0,implicit"`
	Revoke    bool     `tl:",omitempty:flags:1,implicit"`
	Peer      InputPeer
	MaxID     int32
	MinDate   *int32 `tl:",omitempty:flags:2"`
	MaxDate   *int32 `tl:",omitempty:flags:3"`
}

func (*MessagesDeleteHistoryRequestPredict) CRC() uint32 {
	return 0xb08f922a
}

func MessagesDeleteHistory(ctx context.Context, m Requester, i MessagesDeleteHistoryRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteMessagesRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Revoke bool     `tl:",omitempty:flags:0,implicit"`
	ID     []int32
}

func (*MessagesDeleteMessagesRequestPredict) CRC() uint32 {
	return 0xe58e95d2
}

func MessagesDeleteMessages(ctx context.Context, m Requester, i MessagesDeleteMessagesRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReceivedMessagesRequestPredict struct {
	MaxID int32
}

func (*MessagesReceivedMessagesRequestPredict) CRC() uint32 {
	return 0x5a954c0
}

func MessagesReceivedMessages(ctx context.Context, m Requester, i MessagesReceivedMessagesRequestPredict) ([]ReceivedNotifyMessage, error) {
	var res []ReceivedNotifyMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSetTypingRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     InputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
	Action   SendMessageAction
}

func (*MessagesSetTypingRequestPredict) CRC() uint32 {
	return 0x58943ee2
}

func MessagesSetTyping(ctx context.Context, m Requester, i MessagesSetTypingRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMessageRequestPredict struct {
	_                      struct{} `tl:"flags,bitflag"`
	NoWebpage              bool     `tl:",omitempty:flags:1,implicit"`
	Silent                 bool     `tl:",omitempty:flags:5,implicit"`
	Background             bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft             bool     `tl:",omitempty:flags:7,implicit"`
	Noforwards             bool     `tl:",omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool     `tl:",omitempty:flags:15,implicit"`
	InvertMedia            bool     `tl:",omitempty:flags:16,implicit"`
	Peer                   InputPeer
	ReplyTo                InputReplyTo `tl:",omitempty:flags:0"`
	Message                string
	RandomID               int64
	ReplyMarkup            ReplyMarkup             `tl:",omitempty:flags:2"`
	Entities               []MessageEntity         `tl:",omitempty:flags:3"`
	ScheduleDate           *int32                  `tl:",omitempty:flags:10"`
	SendAs                 InputPeer               `tl:",omitempty:flags:13"`
	QuickReplyShortcut     InputQuickReplyShortcut `tl:",omitempty:flags:17"`
	Effect                 *int64                  `tl:",omitempty:flags:18"`
}

func (*MessagesSendMessageRequestPredict) CRC() uint32 {
	return 0x983f9745
}

func MessagesSendMessage(ctx context.Context, m Requester, i MessagesSendMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMediaRequestPredict struct {
	_                      struct{} `tl:"flags,bitflag"`
	Silent                 bool     `tl:",omitempty:flags:5,implicit"`
	Background             bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft             bool     `tl:",omitempty:flags:7,implicit"`
	Noforwards             bool     `tl:",omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool     `tl:",omitempty:flags:15,implicit"`
	InvertMedia            bool     `tl:",omitempty:flags:16,implicit"`
	Peer                   InputPeer
	ReplyTo                InputReplyTo `tl:",omitempty:flags:0"`
	Media                  InputMedia
	Message                string
	RandomID               int64
	ReplyMarkup            ReplyMarkup             `tl:",omitempty:flags:2"`
	Entities               []MessageEntity         `tl:",omitempty:flags:3"`
	ScheduleDate           *int32                  `tl:",omitempty:flags:10"`
	SendAs                 InputPeer               `tl:",omitempty:flags:13"`
	QuickReplyShortcut     InputQuickReplyShortcut `tl:",omitempty:flags:17"`
	Effect                 *int64                  `tl:",omitempty:flags:18"`
}

func (*MessagesSendMediaRequestPredict) CRC() uint32 {
	return 0x7852834e
}

func MessagesSendMedia(ctx context.Context, m Requester, i MessagesSendMediaRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesForwardMessagesRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Silent             bool     `tl:",omitempty:flags:5,implicit"`
	Background         bool     `tl:",omitempty:flags:6,implicit"`
	WithMyScore        bool     `tl:",omitempty:flags:8,implicit"`
	DropAuthor         bool     `tl:",omitempty:flags:11,implicit"`
	DropMediaCaptions  bool     `tl:",omitempty:flags:12,implicit"`
	Noforwards         bool     `tl:",omitempty:flags:14,implicit"`
	FromPeer           InputPeer
	ID                 []int32
	RandomID           []int64
	ToPeer             InputPeer
	TopMsgID           *int32                  `tl:",omitempty:flags:9"`
	ScheduleDate       *int32                  `tl:",omitempty:flags:10"`
	SendAs             InputPeer               `tl:",omitempty:flags:13"`
	QuickReplyShortcut InputQuickReplyShortcut `tl:",omitempty:flags:17"`
}

func (*MessagesForwardMessagesRequestPredict) CRC() uint32 {
	return 0xd5039208
}

func MessagesForwardMessages(ctx context.Context, m Requester, i MessagesForwardMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesReportSpamRequestPredict struct {
	Peer InputPeer
}

func (*MessagesReportSpamRequestPredict) CRC() uint32 {
	return 0xcf1592db
}

func MessagesReportSpam(ctx context.Context, m Requester, i MessagesReportSpamRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPeerSettingsRequestPredict struct {
	Peer InputPeer
}

func (*MessagesGetPeerSettingsRequestPredict) CRC() uint32 {
	return 0xefd9a6a2
}

func MessagesGetPeerSettings(ctx context.Context, m Requester, i MessagesGetPeerSettingsRequestPredict) (MessagesPeerSettings, error) {
	var res MessagesPeerSettings
	return res, request(ctx, m, &i, &res)
}

type MessagesReportRequestPredict struct {
	Peer    InputPeer
	ID      []int32
	Reason  ReportReason
	Message string
}

func (*MessagesReportRequestPredict) CRC() uint32 {
	return 0x8953ab4e
}

func MessagesReport(ctx context.Context, m Requester, i MessagesReportRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetChatsRequestPredict struct {
	ID []int64
}

func (*MessagesGetChatsRequestPredict) CRC() uint32 {
	return 0x49e9528f
}

func MessagesGetChats(ctx context.Context, m Requester, i MessagesGetChatsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFullChatRequestPredict struct {
	ChatID int64
}

func (*MessagesGetFullChatRequestPredict) CRC() uint32 {
	return 0xaeb00b34
}

func MessagesGetFullChat(ctx context.Context, m Requester, i MessagesGetFullChatRequestPredict) (MessagesChatFull, error) {
	var res MessagesChatFull
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatTitleRequestPredict struct {
	ChatID int64
	Title  string
}

func (*MessagesEditChatTitleRequestPredict) CRC() uint32 {
	return 0x73783ffd
}

func MessagesEditChatTitle(ctx context.Context, m Requester, i MessagesEditChatTitleRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatPhotoRequestPredict struct {
	ChatID int64
	Photo  InputChatPhoto
}

func (*MessagesEditChatPhotoRequestPredict) CRC() uint32 {
	return 0x35ddd674
}

func MessagesEditChatPhoto(ctx context.Context, m Requester, i MessagesEditChatPhotoRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesAddChatUserRequestPredict struct {
	ChatID   int64
	UserID   InputUser
	FwdLimit int32
}

func (*MessagesAddChatUserRequestPredict) CRC() uint32 {
	return 0xcbc6d107
}

func MessagesAddChatUser(ctx context.Context, m Requester, i MessagesAddChatUserRequestPredict) (MessagesInvitedUsers, error) {
	var res MessagesInvitedUsers
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteChatUserRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	RevokeHistory bool     `tl:",omitempty:flags:0,implicit"`
	ChatID        int64
	UserID        InputUser
}

func (*MessagesDeleteChatUserRequestPredict) CRC() uint32 {
	return 0xa2185cab
}

func MessagesDeleteChatUser(ctx context.Context, m Requester, i MessagesDeleteChatUserRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesCreateChatRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Users     []InputUser
	Title     string
	TTLPeriod *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesCreateChatRequestPredict) CRC() uint32 {
	return 0x92ceddd4
}

func MessagesCreateChat(ctx context.Context, m Requester, i MessagesCreateChatRequestPredict) (MessagesInvitedUsers, error) {
	var res MessagesInvitedUsers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDhConfigRequestPredict struct {
	Version      int32
	RandomLength int32
}

func (*MessagesGetDhConfigRequestPredict) CRC() uint32 {
	return 0x26cf8950
}

func MessagesGetDhConfig(ctx context.Context, m Requester, i MessagesGetDhConfigRequestPredict) (MessagesDhConfig, error) {
	var res MessagesDhConfig
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestEncryptionRequestPredict struct {
	UserID   InputUser
	RandomID int32
	GA       []byte
}

func (*MessagesRequestEncryptionRequestPredict) CRC() uint32 {
	return 0xf64daf43
}

func MessagesRequestEncryption(ctx context.Context, m Requester, i MessagesRequestEncryptionRequestPredict) (EncryptedChat, error) {
	var res EncryptedChat
	return res, request(ctx, m, &i, &res)
}

type MessagesAcceptEncryptionRequestPredict struct {
	Peer           InputEncryptedChat
	GB             []byte
	KeyFingerprint int64
}

func (*MessagesAcceptEncryptionRequestPredict) CRC() uint32 {
	return 0x3dbc0415
}

func MessagesAcceptEncryption(ctx context.Context, m Requester, i MessagesAcceptEncryptionRequestPredict) (EncryptedChat, error) {
	var res EncryptedChat
	return res, request(ctx, m, &i, &res)
}

type MessagesDiscardEncryptionRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	DeleteHistory bool     `tl:",omitempty:flags:0,implicit"`
	ChatID        int32
}

func (*MessagesDiscardEncryptionRequestPredict) CRC() uint32 {
	return 0xf393aea0
}

func MessagesDiscardEncryption(ctx context.Context, m Requester, i MessagesDiscardEncryptionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSetEncryptedTypingRequestPredict struct {
	Peer   InputEncryptedChat
	Typing bool
}

func (*MessagesSetEncryptedTypingRequestPredict) CRC() uint32 {
	return 0x791451ed
}

func MessagesSetEncryptedTyping(ctx context.Context, m Requester, i MessagesSetEncryptedTypingRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReadEncryptedHistoryRequestPredict struct {
	Peer    InputEncryptedChat
	MaxDate int32
}

func (*MessagesReadEncryptedHistoryRequestPredict) CRC() uint32 {
	return 0x7f4b690a
}

func MessagesReadEncryptedHistory(ctx context.Context, m Requester, i MessagesReadEncryptedHistoryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Silent   bool     `tl:",omitempty:flags:0,implicit"`
	Peer     InputEncryptedChat
	RandomID int64
	Data     []byte
}

func (*MessagesSendEncryptedRequestPredict) CRC() uint32 {
	return 0x44fa7a15
}

func MessagesSendEncrypted(ctx context.Context, m Requester, i MessagesSendEncryptedRequestPredict) (MessagesSentEncryptedMessage, error) {
	var res MessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedFileRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Silent   bool     `tl:",omitempty:flags:0,implicit"`
	Peer     InputEncryptedChat
	RandomID int64
	Data     []byte
	File     InputEncryptedFile
}

func (*MessagesSendEncryptedFileRequestPredict) CRC() uint32 {
	return 0x5559481d
}

func MessagesSendEncryptedFile(ctx context.Context, m Requester, i MessagesSendEncryptedFileRequestPredict) (MessagesSentEncryptedMessage, error) {
	var res MessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedServiceRequestPredict struct {
	Peer     InputEncryptedChat
	RandomID int64
	Data     []byte
}

func (*MessagesSendEncryptedServiceRequestPredict) CRC() uint32 {
	return 0x32d439a4
}

func MessagesSendEncryptedService(ctx context.Context, m Requester, i MessagesSendEncryptedServiceRequestPredict) (MessagesSentEncryptedMessage, error) {
	var res MessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesReceivedQueueRequestPredict struct {
	MaxQts int32
}

func (*MessagesReceivedQueueRequestPredict) CRC() uint32 {
	return 0x55a5bb66
}

func MessagesReceivedQueue(ctx context.Context, m Requester, i MessagesReceivedQueueRequestPredict) ([]int64, error) {
	var res []int64
	return res, request(ctx, m, &i, &res)
}

type MessagesReportEncryptedSpamRequestPredict struct {
	Peer InputEncryptedChat
}

func (*MessagesReportEncryptedSpamRequestPredict) CRC() uint32 {
	return 0x4b0c8c0f
}

func MessagesReportEncryptedSpam(ctx context.Context, m Requester, i MessagesReportEncryptedSpamRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReadMessageContentsRequestPredict struct {
	ID []int32
}

func (*MessagesReadMessageContentsRequestPredict) CRC() uint32 {
	return 0x36a73f77
}

func MessagesReadMessageContents(ctx context.Context, m Requester, i MessagesReadMessageContentsRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetStickersRequestPredict struct {
	Emoticon string
	Hash     int64
}

func (*MessagesGetStickersRequestPredict) CRC() uint32 {
	return 0xd5a5d3a1
}

func MessagesGetStickers(ctx context.Context, m Requester, i MessagesGetStickersRequestPredict) (MessagesStickers, error) {
	var res MessagesStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAllStickersRequestPredict struct {
	Hash int64
}

func (*MessagesGetAllStickersRequestPredict) CRC() uint32 {
	return 0xb8a0a1a8
}

func MessagesGetAllStickers(ctx context.Context, m Requester, i MessagesGetAllStickersRequestPredict) (MessagesAllStickers, error) {
	var res MessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetWebPagePreviewRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Message  string
	Entities []MessageEntity `tl:",omitempty:flags:3"`
}

func (*MessagesGetWebPagePreviewRequestPredict) CRC() uint32 {
	return 0x8b68b0cc
}

func MessagesGetWebPagePreview(ctx context.Context, m Requester, i MessagesGetWebPagePreviewRequestPredict) (MessageMedia, error) {
	var res MessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesExportChatInviteRequestPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	LegacyRevokePermanent bool     `tl:",omitempty:flags:2,implicit"`
	RequestNeeded         bool     `tl:",omitempty:flags:3,implicit"`
	Peer                  InputPeer
	ExpireDate            *int32  `tl:",omitempty:flags:0"`
	UsageLimit            *int32  `tl:",omitempty:flags:1"`
	Title                 *string `tl:",omitempty:flags:4"`
}

func (*MessagesExportChatInviteRequestPredict) CRC() uint32 {
	return 0xa02ce5d5
}

func MessagesExportChatInvite(ctx context.Context, m Requester, i MessagesExportChatInviteRequestPredict) (ExportedChatInvite, error) {
	var res ExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckChatInviteRequestPredict struct {
	Hash string
}

func (*MessagesCheckChatInviteRequestPredict) CRC() uint32 {
	return 0x3eadb1bb
}

func MessagesCheckChatInvite(ctx context.Context, m Requester, i MessagesCheckChatInviteRequestPredict) (ChatInvite, error) {
	var res ChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesImportChatInviteRequestPredict struct {
	Hash string
}

func (*MessagesImportChatInviteRequestPredict) CRC() uint32 {
	return 0x6c50051c
}

func MessagesImportChatInvite(ctx context.Context, m Requester, i MessagesImportChatInviteRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetStickerSetRequestPredict struct {
	Stickerset InputStickerSet
	Hash       int32
}

func (*MessagesGetStickerSetRequestPredict) CRC() uint32 {
	return 0xc8a0ec74
}

func MessagesGetStickerSet(ctx context.Context, m Requester, i MessagesGetStickerSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type MessagesInstallStickerSetRequestPredict struct {
	Stickerset InputStickerSet
	Archived   bool
}

func (*MessagesInstallStickerSetRequestPredict) CRC() uint32 {
	return 0xc78fe460
}

func MessagesInstallStickerSet(ctx context.Context, m Requester, i MessagesInstallStickerSetRequestPredict) (MessagesStickerSetInstallResult, error) {
	var res MessagesStickerSetInstallResult
	return res, request(ctx, m, &i, &res)
}

type MessagesUninstallStickerSetRequestPredict struct {
	Stickerset InputStickerSet
}

func (*MessagesUninstallStickerSetRequestPredict) CRC() uint32 {
	return 0xf96e55de
}

func MessagesUninstallStickerSet(ctx context.Context, m Requester, i MessagesUninstallStickerSetRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesStartBotRequestPredict struct {
	Bot        InputUser
	Peer       InputPeer
	RandomID   int64
	StartParam string
}

func (*MessagesStartBotRequestPredict) CRC() uint32 {
	return 0xe6df7378
}

func MessagesStartBot(ctx context.Context, m Requester, i MessagesStartBotRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesViewsRequestPredict struct {
	Peer      InputPeer
	ID        []int32
	Increment bool
}

func (*MessagesGetMessagesViewsRequestPredict) CRC() uint32 {
	return 0x5784d3e1
}

func MessagesGetMessagesViews(ctx context.Context, m Requester, i MessagesGetMessagesViewsRequestPredict) (MessagesMessageViews, error) {
	var res MessagesMessageViews
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatAdminRequestPredict struct {
	ChatID  int64
	UserID  InputUser
	IsAdmin bool
}

func (*MessagesEditChatAdminRequestPredict) CRC() uint32 {
	return 0xa85bd1c2
}

func MessagesEditChatAdmin(ctx context.Context, m Requester, i MessagesEditChatAdminRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesMigrateChatRequestPredict struct {
	ChatID int64
}

func (*MessagesMigrateChatRequestPredict) CRC() uint32 {
	return 0xa2875319
}

func MessagesMigrateChat(ctx context.Context, m Requester, i MessagesMigrateChatRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchGlobalRequestPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	BroadcastsOnly bool     `tl:",omitempty:flags:1,implicit"`
	FolderID       *int32   `tl:",omitempty:flags:0"`
	Q              string
	Filter         MessagesFilter
	MinDate        int32
	MaxDate        int32
	OffsetRate     int32
	OffsetPeer     InputPeer
	OffsetID       int32
	Limit          int32
}

func (*MessagesSearchGlobalRequestPredict) CRC() uint32 {
	return 0x4bc6589a
}

func MessagesSearchGlobal(ctx context.Context, m Requester, i MessagesSearchGlobalRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderStickerSetsRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:",omitempty:flags:0,implicit"`
	Emojis bool     `tl:",omitempty:flags:1,implicit"`
	Order  []int64
}

func (*MessagesReorderStickerSetsRequestPredict) CRC() uint32 {
	return 0x78337739
}

func MessagesReorderStickerSets(ctx context.Context, m Requester, i MessagesReorderStickerSetsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDocumentByHashRequestPredict struct {
	SHA256   []byte
	Size     int64
	MimeType string
}

func (*MessagesGetDocumentByHashRequestPredict) CRC() uint32 {
	return 0xb1f2061f
}

func MessagesGetDocumentByHash(ctx context.Context, m Requester, i MessagesGetDocumentByHashRequestPredict) (Document, error) {
	var res Document
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedGifsRequestPredict struct {
	Hash int64
}

func (*MessagesGetSavedGifsRequestPredict) CRC() uint32 {
	return 0x5cf09635
}

func MessagesGetSavedGifs(ctx context.Context, m Requester, i MessagesGetSavedGifsRequestPredict) (MessagesSavedGifs, error) {
	var res MessagesSavedGifs
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveGifRequestPredict struct {
	ID     InputDocument
	Unsave bool
}

func (*MessagesSaveGifRequestPredict) CRC() uint32 {
	return 0x327a30cb
}

func MessagesSaveGif(ctx context.Context, m Requester, i MessagesSaveGifRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetInlineBotResultsRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Bot      InputUser
	Peer     InputPeer
	GeoPoint InputGeoPoint `tl:",omitempty:flags:0"`
	Query    string
	Offset   string
}

func (*MessagesGetInlineBotResultsRequestPredict) CRC() uint32 {
	return 0x514e999d
}

func MessagesGetInlineBotResults(ctx context.Context, m Requester, i MessagesGetInlineBotResultsRequestPredict) (MessagesBotResults, error) {
	var res MessagesBotResults
	return res, request(ctx, m, &i, &res)
}

type MessagesSetInlineBotResultsRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Gallery       bool     `tl:",omitempty:flags:0,implicit"`
	Private       bool     `tl:",omitempty:flags:1,implicit"`
	QueryID       int64
	Results       []InputBotInlineResult
	CacheTime     int32
	NextOffset    *string           `tl:",omitempty:flags:2"`
	SwitchPm      InlineBotSwitchPm `tl:",omitempty:flags:3"`
	SwitchWebview InlineBotWebView  `tl:",omitempty:flags:4"`
}

func (*MessagesSetInlineBotResultsRequestPredict) CRC() uint32 {
	return 0xbb12a419
}

func MessagesSetInlineBotResults(ctx context.Context, m Requester, i MessagesSetInlineBotResultsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendInlineBotResultRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Silent             bool     `tl:",omitempty:flags:5,implicit"`
	Background         bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft         bool     `tl:",omitempty:flags:7,implicit"`
	HideVia            bool     `tl:",omitempty:flags:11,implicit"`
	Peer               InputPeer
	ReplyTo            InputReplyTo `tl:",omitempty:flags:0"`
	RandomID           int64
	QueryID            int64
	ID                 string
	ScheduleDate       *int32                  `tl:",omitempty:flags:10"`
	SendAs             InputPeer               `tl:",omitempty:flags:13"`
	QuickReplyShortcut InputQuickReplyShortcut `tl:",omitempty:flags:17"`
}

func (*MessagesSendInlineBotResultRequestPredict) CRC() uint32 {
	return 0x3ebee86a
}

func MessagesSendInlineBotResult(ctx context.Context, m Requester, i MessagesSendInlineBotResultRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageEditDataRequestPredict struct {
	Peer InputPeer
	ID   int32
}

func (*MessagesGetMessageEditDataRequestPredict) CRC() uint32 {
	return 0xfda68d36
}

func MessagesGetMessageEditData(ctx context.Context, m Requester, i MessagesGetMessageEditDataRequestPredict) (MessagesMessageEditData, error) {
	var res MessagesMessageEditData
	return res, request(ctx, m, &i, &res)
}

type MessagesEditMessageRequestPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	NoWebpage            bool     `tl:",omitempty:flags:1,implicit"`
	InvertMedia          bool     `tl:",omitempty:flags:16,implicit"`
	Peer                 InputPeer
	ID                   int32
	Message              *string         `tl:",omitempty:flags:11"`
	Media                InputMedia      `tl:",omitempty:flags:14"`
	ReplyMarkup          ReplyMarkup     `tl:",omitempty:flags:2"`
	Entities             []MessageEntity `tl:",omitempty:flags:3"`
	ScheduleDate         *int32          `tl:",omitempty:flags:15"`
	QuickReplyShortcutID *int32          `tl:",omitempty:flags:17"`
}

func (*MessagesEditMessageRequestPredict) CRC() uint32 {
	return 0xdfd14005
}

func MessagesEditMessage(ctx context.Context, m Requester, i MessagesEditMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesEditInlineBotMessageRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	NoWebpage   bool     `tl:",omitempty:flags:1,implicit"`
	InvertMedia bool     `tl:",omitempty:flags:16,implicit"`
	ID          InputBotInlineMessageID
	Message     *string         `tl:",omitempty:flags:11"`
	Media       InputMedia      `tl:",omitempty:flags:14"`
	ReplyMarkup ReplyMarkup     `tl:",omitempty:flags:2"`
	Entities    []MessageEntity `tl:",omitempty:flags:3"`
}

func (*MessagesEditInlineBotMessageRequestPredict) CRC() uint32 {
	return 0x83557dba
}

func MessagesEditInlineBotMessage(ctx context.Context, m Requester, i MessagesEditInlineBotMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetBotCallbackAnswerRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Game     bool     `tl:",omitempty:flags:1,implicit"`
	Peer     InputPeer
	MsgID    int32
	Data     *[]byte               `tl:",omitempty:flags:0"`
	Password InputCheckPasswordSRP `tl:",omitempty:flags:2"`
}

func (*MessagesGetBotCallbackAnswerRequestPredict) CRC() uint32 {
	return 0x9342ca07
}

func MessagesGetBotCallbackAnswer(ctx context.Context, m Requester, i MessagesGetBotCallbackAnswerRequestPredict) (MessagesBotCallbackAnswer, error) {
	var res MessagesBotCallbackAnswer
	return res, request(ctx, m, &i, &res)
}

type MessagesSetBotCallbackAnswerRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Alert     bool     `tl:",omitempty:flags:1,implicit"`
	QueryID   int64
	Message   *string `tl:",omitempty:flags:0"`
	URL       *string `tl:",omitempty:flags:2"`
	CacheTime int32
}

func (*MessagesSetBotCallbackAnswerRequestPredict) CRC() uint32 {
	return 0xd58f130a
}

func MessagesSetBotCallbackAnswer(ctx context.Context, m Requester, i MessagesSetBotCallbackAnswerRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPeerDialogsRequestPredict struct {
	Peers []InputDialogPeer
}

func (*MessagesGetPeerDialogsRequestPredict) CRC() uint32 {
	return 0xe470bcfd
}

func MessagesGetPeerDialogs(ctx context.Context, m Requester, i MessagesGetPeerDialogsRequestPredict) (MessagesPeerDialogs, error) {
	var res MessagesPeerDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveDraftRequestPredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	NoWebpage   bool         `tl:",omitempty:flags:1,implicit"`
	InvertMedia bool         `tl:",omitempty:flags:6,implicit"`
	ReplyTo     InputReplyTo `tl:",omitempty:flags:4"`
	Peer        InputPeer
	Message     string
	Entities    []MessageEntity `tl:",omitempty:flags:3"`
	Media       InputMedia      `tl:",omitempty:flags:5"`
	Effect      *int64          `tl:",omitempty:flags:7"`
}

func (*MessagesSaveDraftRequestPredict) CRC() uint32 {
	return 0xd372c5ce
}

func MessagesSaveDraft(ctx context.Context, m Requester, i MessagesSaveDraftRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAllDraftsRequestPredict struct{}

func (*MessagesGetAllDraftsRequestPredict) CRC() uint32 {
	return 0x6a3f8d65
}

func MessagesGetAllDrafts(ctx context.Context, m Requester, i MessagesGetAllDraftsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFeaturedStickersRequestPredict struct {
	Hash int64
}

func (*MessagesGetFeaturedStickersRequestPredict) CRC() uint32 {
	return 0x64780b14
}

func MessagesGetFeaturedStickers(ctx context.Context, m Requester, i MessagesGetFeaturedStickersRequestPredict) (MessagesFeaturedStickers, error) {
	var res MessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesReadFeaturedStickersRequestPredict struct {
	ID []int64
}

func (*MessagesReadFeaturedStickersRequestPredict) CRC() uint32 {
	return 0x5b118126
}

func MessagesReadFeaturedStickers(ctx context.Context, m Requester, i MessagesReadFeaturedStickersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentStickersRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Attached bool     `tl:",omitempty:flags:0,implicit"`
	Hash     int64
}

func (*MessagesGetRecentStickersRequestPredict) CRC() uint32 {
	return 0x9da9403b
}

func MessagesGetRecentStickers(ctx context.Context, m Requester, i MessagesGetRecentStickersRequestPredict) (MessagesRecentStickers, error) {
	var res MessagesRecentStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveRecentStickerRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Attached bool     `tl:",omitempty:flags:0,implicit"`
	ID       InputDocument
	Unsave   bool
}

func (*MessagesSaveRecentStickerRequestPredict) CRC() uint32 {
	return 0x392718f8
}

func MessagesSaveRecentSticker(ctx context.Context, m Requester, i MessagesSaveRecentStickerRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesClearRecentStickersRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Attached bool     `tl:",omitempty:flags:0,implicit"`
}

func (*MessagesClearRecentStickersRequestPredict) CRC() uint32 {
	return 0x8999602d
}

func MessagesClearRecentStickers(ctx context.Context, m Requester, i MessagesClearRecentStickersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetArchivedStickersRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Masks    bool     `tl:",omitempty:flags:0,implicit"`
	Emojis   bool     `tl:",omitempty:flags:1,implicit"`
	OffsetID int64
	Limit    int32
}

func (*MessagesGetArchivedStickersRequestPredict) CRC() uint32 {
	return 0x57f17692
}

func MessagesGetArchivedStickers(ctx context.Context, m Requester, i MessagesGetArchivedStickersRequestPredict) (MessagesArchivedStickers, error) {
	var res MessagesArchivedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMaskStickersRequestPredict struct {
	Hash int64
}

func (*MessagesGetMaskStickersRequestPredict) CRC() uint32 {
	return 0x640f82b8
}

func MessagesGetMaskStickers(ctx context.Context, m Requester, i MessagesGetMaskStickersRequestPredict) (MessagesAllStickers, error) {
	var res MessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachedStickersRequestPredict struct {
	Media InputStickeredMedia
}

func (*MessagesGetAttachedStickersRequestPredict) CRC() uint32 {
	return 0xcc5b67cc
}

func MessagesGetAttachedStickers(ctx context.Context, m Requester, i MessagesGetAttachedStickersRequestPredict) ([]StickerSetCovered, error) {
	var res []StickerSetCovered
	return res, request(ctx, m, &i, &res)
}

type MessagesSetGameScoreRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	EditMessage bool     `tl:",omitempty:flags:0,implicit"`
	Force       bool     `tl:",omitempty:flags:1,implicit"`
	Peer        InputPeer
	ID          int32
	UserID      InputUser
	Score       int32
}

func (*MessagesSetGameScoreRequestPredict) CRC() uint32 {
	return 0x8ef8ecc0
}

func MessagesSetGameScore(ctx context.Context, m Requester, i MessagesSetGameScoreRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSetInlineGameScoreRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	EditMessage bool     `tl:",omitempty:flags:0,implicit"`
	Force       bool     `tl:",omitempty:flags:1,implicit"`
	ID          InputBotInlineMessageID
	UserID      InputUser
	Score       int32
}

func (*MessagesSetInlineGameScoreRequestPredict) CRC() uint32 {
	return 0x15ad9f64
}

func MessagesSetInlineGameScore(ctx context.Context, m Requester, i MessagesSetInlineGameScoreRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetGameHighScoresRequestPredict struct {
	Peer   InputPeer
	ID     int32
	UserID InputUser
}

func (*MessagesGetGameHighScoresRequestPredict) CRC() uint32 {
	return 0xe822649d
}

func MessagesGetGameHighScores(ctx context.Context, m Requester, i MessagesGetGameHighScoresRequestPredict) (MessagesHighScores, error) {
	var res MessagesHighScores
	return res, request(ctx, m, &i, &res)
}

type MessagesGetInlineGameHighScoresRequestPredict struct {
	ID     InputBotInlineMessageID
	UserID InputUser
}

func (*MessagesGetInlineGameHighScoresRequestPredict) CRC() uint32 {
	return 0xf635e1b
}

func MessagesGetInlineGameHighScores(ctx context.Context, m Requester, i MessagesGetInlineGameHighScoresRequestPredict) (MessagesHighScores, error) {
	var res MessagesHighScores
	return res, request(ctx, m, &i, &res)
}

type MessagesGetCommonChatsRequestPredict struct {
	UserID InputUser
	MaxID  int64
	Limit  int32
}

func (*MessagesGetCommonChatsRequestPredict) CRC() uint32 {
	return 0xe40ca104
}

func MessagesGetCommonChats(ctx context.Context, m Requester, i MessagesGetCommonChatsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type MessagesGetWebPageRequestPredict struct {
	URL  string
	Hash int32
}

func (*MessagesGetWebPageRequestPredict) CRC() uint32 {
	return 0x8d9692a3
}

func MessagesGetWebPage(ctx context.Context, m Requester, i MessagesGetWebPageRequestPredict) (MessagesWebPage, error) {
	var res MessagesWebPage
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleDialogPinRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Pinned bool     `tl:",omitempty:flags:0,implicit"`
	Peer   InputDialogPeer
}

func (*MessagesToggleDialogPinRequestPredict) CRC() uint32 {
	return 0xa731e257
}

func MessagesToggleDialogPin(ctx context.Context, m Requester, i MessagesToggleDialogPinRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderPinnedDialogsRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Force    bool     `tl:",omitempty:flags:0,implicit"`
	FolderID int32
	Order    []InputDialogPeer
}

func (*MessagesReorderPinnedDialogsRequestPredict) CRC() uint32 {
	return 0x3b1adf37
}

func MessagesReorderPinnedDialogs(ctx context.Context, m Requester, i MessagesReorderPinnedDialogsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPinnedDialogsRequestPredict struct {
	FolderID int32
}

func (*MessagesGetPinnedDialogsRequestPredict) CRC() uint32 {
	return 0xd6b94df2
}

func MessagesGetPinnedDialogs(ctx context.Context, m Requester, i MessagesGetPinnedDialogsRequestPredict) (MessagesPeerDialogs, error) {
	var res MessagesPeerDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesSetBotShippingResultsRequestPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	QueryID         int64
	Error           *string          `tl:",omitempty:flags:0"`
	ShippingOptions []ShippingOption `tl:",omitempty:flags:1"`
}

func (*MessagesSetBotShippingResultsRequestPredict) CRC() uint32 {
	return 0xe5f672fa
}

func MessagesSetBotShippingResults(ctx context.Context, m Requester, i MessagesSetBotShippingResultsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSetBotPrecheckoutResultsRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Success bool     `tl:",omitempty:flags:1,implicit"`
	QueryID int64
	Error   *string `tl:",omitempty:flags:0"`
}

func (*MessagesSetBotPrecheckoutResultsRequestPredict) CRC() uint32 {
	return 0x9c2dd95
}

func MessagesSetBotPrecheckoutResults(ctx context.Context, m Requester, i MessagesSetBotPrecheckoutResultsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadMediaRequestPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	BusinessConnectionID *string  `tl:",omitempty:flags:0"`
	Peer                 InputPeer
	Media                InputMedia
}

func (*MessagesUploadMediaRequestPredict) CRC() uint32 {
	return 0x14967978
}

func MessagesUploadMedia(ctx context.Context, m Requester, i MessagesUploadMediaRequestPredict) (MessageMedia, error) {
	var res MessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesSendScreenshotNotificationRequestPredict struct {
	Peer     InputPeer
	ReplyTo  InputReplyTo
	RandomID int64
}

func (*MessagesSendScreenshotNotificationRequestPredict) CRC() uint32 {
	return 0xa1405817
}

func MessagesSendScreenshotNotification(ctx context.Context, m Requester, i MessagesSendScreenshotNotificationRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFavedStickersRequestPredict struct {
	Hash int64
}

func (*MessagesGetFavedStickersRequestPredict) CRC() uint32 {
	return 0x4f1aaa9
}

func MessagesGetFavedStickers(ctx context.Context, m Requester, i MessagesGetFavedStickersRequestPredict) (MessagesFavedStickers, error) {
	var res MessagesFavedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesFaveStickerRequestPredict struct {
	ID     InputDocument
	Unfave bool
}

func (*MessagesFaveStickerRequestPredict) CRC() uint32 {
	return 0xb9ffc55b
}

func MessagesFaveSticker(ctx context.Context, m Requester, i MessagesFaveStickerRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetUnreadMentionsRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      InputPeer
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
}

func (*MessagesGetUnreadMentionsRequestPredict) CRC() uint32 {
	return 0xf107e790
}

func MessagesGetUnreadMentions(ctx context.Context, m Requester, i MessagesGetUnreadMentionsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadMentionsRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     InputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesReadMentionsRequestPredict) CRC() uint32 {
	return 0x36e5bf4d
}

func MessagesReadMentions(ctx context.Context, m Requester, i MessagesReadMentionsRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentLocationsRequestPredict struct {
	Peer  InputPeer
	Limit int32
	Hash  int64
}

func (*MessagesGetRecentLocationsRequestPredict) CRC() uint32 {
	return 0x702a40e0
}

func MessagesGetRecentLocations(ctx context.Context, m Requester, i MessagesGetRecentLocationsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMultiMediaRequestPredict struct {
	_                      struct{} `tl:"flags,bitflag"`
	Silent                 bool     `tl:",omitempty:flags:5,implicit"`
	Background             bool     `tl:",omitempty:flags:6,implicit"`
	ClearDraft             bool     `tl:",omitempty:flags:7,implicit"`
	Noforwards             bool     `tl:",omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool     `tl:",omitempty:flags:15,implicit"`
	InvertMedia            bool     `tl:",omitempty:flags:16,implicit"`
	Peer                   InputPeer
	ReplyTo                InputReplyTo `tl:",omitempty:flags:0"`
	MultiMedia             []InputSingleMedia
	ScheduleDate           *int32                  `tl:",omitempty:flags:10"`
	SendAs                 InputPeer               `tl:",omitempty:flags:13"`
	QuickReplyShortcut     InputQuickReplyShortcut `tl:",omitempty:flags:17"`
	Effect                 *int64                  `tl:",omitempty:flags:18"`
}

func (*MessagesSendMultiMediaRequestPredict) CRC() uint32 {
	return 0x37b74355
}

func MessagesSendMultiMedia(ctx context.Context, m Requester, i MessagesSendMultiMediaRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadEncryptedFileRequestPredict struct {
	Peer InputEncryptedChat
	File InputEncryptedFile
}

func (*MessagesUploadEncryptedFileRequestPredict) CRC() uint32 {
	return 0x5057c497
}

func MessagesUploadEncryptedFile(ctx context.Context, m Requester, i MessagesUploadEncryptedFileRequestPredict) (EncryptedFile, error) {
	var res EncryptedFile
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchStickerSetsRequestPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ExcludeFeatured bool     `tl:",omitempty:flags:0,implicit"`
	Q               string
	Hash            int64
}

func (*MessagesSearchStickerSetsRequestPredict) CRC() uint32 {
	return 0x35705b8a
}

func MessagesSearchStickerSets(ctx context.Context, m Requester, i MessagesSearchStickerSetsRequestPredict) (MessagesFoundStickerSets, error) {
	var res MessagesFoundStickerSets
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSplitRangesRequestPredict struct{}

func (*MessagesGetSplitRangesRequestPredict) CRC() uint32 {
	return 0x1cff7e08
}

func MessagesGetSplitRanges(ctx context.Context, m Requester, i MessagesGetSplitRangesRequestPredict) ([]MessageRange, error) {
	var res []MessageRange
	return res, request(ctx, m, &i, &res)
}

type MessagesMarkDialogUnreadRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Unread bool     `tl:",omitempty:flags:0,implicit"`
	Peer   InputDialogPeer
}

func (*MessagesMarkDialogUnreadRequestPredict) CRC() uint32 {
	return 0xc286d98f
}

func MessagesMarkDialogUnread(ctx context.Context, m Requester, i MessagesMarkDialogUnreadRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDialogUnreadMarksRequestPredict struct{}

func (*MessagesGetDialogUnreadMarksRequestPredict) CRC() uint32 {
	return 0x22e24e22
}

func MessagesGetDialogUnreadMarks(ctx context.Context, m Requester, i MessagesGetDialogUnreadMarksRequestPredict) ([]DialogPeer, error) {
	var res []DialogPeer
	return res, request(ctx, m, &i, &res)
}

type MessagesClearAllDraftsRequestPredict struct{}

func (*MessagesClearAllDraftsRequestPredict) CRC() uint32 {
	return 0x7e58ee9c
}

func MessagesClearAllDrafts(ctx context.Context, m Requester, i MessagesClearAllDraftsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdatePinnedMessageRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Silent    bool     `tl:",omitempty:flags:0,implicit"`
	Unpin     bool     `tl:",omitempty:flags:1,implicit"`
	PmOneside bool     `tl:",omitempty:flags:2,implicit"`
	Peer      InputPeer
	ID        int32
}

func (*MessagesUpdatePinnedMessageRequestPredict) CRC() uint32 {
	return 0xd2aaf7ec
}

func MessagesUpdatePinnedMessage(ctx context.Context, m Requester, i MessagesUpdatePinnedMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSendVoteRequestPredict struct {
	Peer    InputPeer
	MsgID   int32
	Options [][]byte
}

func (*MessagesSendVoteRequestPredict) CRC() uint32 {
	return 0x10ea6184
}

func MessagesSendVote(ctx context.Context, m Requester, i MessagesSendVoteRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPollResultsRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*MessagesGetPollResultsRequestPredict) CRC() uint32 {
	return 0x73bb643b
}

func MessagesGetPollResults(ctx context.Context, m Requester, i MessagesGetPollResultsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOnlinesRequestPredict struct {
	Peer InputPeer
}

func (*MessagesGetOnlinesRequestPredict) CRC() uint32 {
	return 0x6e2be050
}

func MessagesGetOnlines(ctx context.Context, m Requester, i MessagesGetOnlinesRequestPredict) (ChatOnlines, error) {
	var res ChatOnlines
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatAboutRequestPredict struct {
	Peer  InputPeer
	About string
}

func (*MessagesEditChatAboutRequestPredict) CRC() uint32 {
	return 0xdef60797
}

func MessagesEditChatAbout(ctx context.Context, m Requester, i MessagesEditChatAboutRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatDefaultBannedRightsRequestPredict struct {
	Peer         InputPeer
	BannedRights ChatBannedRights
}

func (*MessagesEditChatDefaultBannedRightsRequestPredict) CRC() uint32 {
	return 0xa5866b41
}

func MessagesEditChatDefaultBannedRights(ctx context.Context, m Requester, i MessagesEditChatDefaultBannedRightsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsRequestPredict struct {
	LangCode string
}

func (*MessagesGetEmojiKeywordsRequestPredict) CRC() uint32 {
	return 0x35a0e062
}

func MessagesGetEmojiKeywords(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsRequestPredict) (EmojiKeywordsDifference, error) {
	var res EmojiKeywordsDifference
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsDifferenceRequestPredict struct {
	LangCode    string
	FromVersion int32
}

func (*MessagesGetEmojiKeywordsDifferenceRequestPredict) CRC() uint32 {
	return 0x1508b6af
}

func MessagesGetEmojiKeywordsDifference(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsDifferenceRequestPredict) (EmojiKeywordsDifference, error) {
	var res EmojiKeywordsDifference
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsLanguagesRequestPredict struct {
	LangCodes []string
}

func (*MessagesGetEmojiKeywordsLanguagesRequestPredict) CRC() uint32 {
	return 0x4e9963b2
}

func MessagesGetEmojiKeywordsLanguages(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsLanguagesRequestPredict) ([]EmojiLanguage, error) {
	var res []EmojiLanguage
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiURLRequestPredict struct {
	LangCode string
}

func (*MessagesGetEmojiURLRequestPredict) CRC() uint32 {
	return 0xd5b10c26
}

func MessagesGetEmojiURL(ctx context.Context, m Requester, i MessagesGetEmojiURLRequestPredict) (EmojiURL, error) {
	var res EmojiURL
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchCountersRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Peer        InputPeer
	SavedPeerID InputPeer `tl:",omitempty:flags:2"`
	TopMsgID    *int32    `tl:",omitempty:flags:0"`
	Filters     []MessagesFilter
}

func (*MessagesGetSearchCountersRequestPredict) CRC() uint32 {
	return 0x1bbcf300
}

func MessagesGetSearchCounters(ctx context.Context, m Requester, i MessagesGetSearchCountersRequestPredict) ([]MessagesSearchCounter, error) {
	var res []MessagesSearchCounter
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestURLAuthRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Peer     InputPeer `tl:",omitempty:flags:1"`
	MsgID    *int32    `tl:",omitempty:flags:1"`
	ButtonID *int32    `tl:",omitempty:flags:1"`
	URL      *string   `tl:",omitempty:flags:2"`
}

func (*MessagesRequestURLAuthRequestPredict) CRC() uint32 {
	return 0x198fb446
}

func MessagesRequestURLAuth(ctx context.Context, m Requester, i MessagesRequestURLAuthRequestPredict) (URLAuthResult, error) {
	var res URLAuthResult
	return res, request(ctx, m, &i, &res)
}

type MessagesAcceptURLAuthRequestPredict struct {
	_            struct{}  `tl:"flags,bitflag"`
	WriteAllowed bool      `tl:",omitempty:flags:0,implicit"`
	Peer         InputPeer `tl:",omitempty:flags:1"`
	MsgID        *int32    `tl:",omitempty:flags:1"`
	ButtonID     *int32    `tl:",omitempty:flags:1"`
	URL          *string   `tl:",omitempty:flags:2"`
}

func (*MessagesAcceptURLAuthRequestPredict) CRC() uint32 {
	return 0xb12c7125
}

func MessagesAcceptURLAuth(ctx context.Context, m Requester, i MessagesAcceptURLAuthRequestPredict) (URLAuthResult, error) {
	var res URLAuthResult
	return res, request(ctx, m, &i, &res)
}

type MessagesHidePeerSettingsBarRequestPredict struct {
	Peer InputPeer
}

func (*MessagesHidePeerSettingsBarRequestPredict) CRC() uint32 {
	return 0x4facb138
}

func MessagesHidePeerSettingsBar(ctx context.Context, m Requester, i MessagesHidePeerSettingsBarRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetScheduledHistoryRequestPredict struct {
	Peer InputPeer
	Hash int64
}

func (*MessagesGetScheduledHistoryRequestPredict) CRC() uint32 {
	return 0xf516760b
}

func MessagesGetScheduledHistory(ctx context.Context, m Requester, i MessagesGetScheduledHistoryRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetScheduledMessagesRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesGetScheduledMessagesRequestPredict) CRC() uint32 {
	return 0xbdbb0464
}

func MessagesGetScheduledMessages(ctx context.Context, m Requester, i MessagesGetScheduledMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendScheduledMessagesRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesSendScheduledMessagesRequestPredict) CRC() uint32 {
	return 0xbd38850a
}

func MessagesSendScheduledMessages(ctx context.Context, m Requester, i MessagesSendScheduledMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteScheduledMessagesRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesDeleteScheduledMessagesRequestPredict) CRC() uint32 {
	return 0x59ae2b16
}

func MessagesDeleteScheduledMessages(ctx context.Context, m Requester, i MessagesDeleteScheduledMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPollVotesRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Peer   InputPeer
	ID     int32
	Option *[]byte `tl:",omitempty:flags:0"`
	Offset *string `tl:",omitempty:flags:1"`
	Limit  int32
}

func (*MessagesGetPollVotesRequestPredict) CRC() uint32 {
	return 0xb86e380e
}

func MessagesGetPollVotes(ctx context.Context, m Requester, i MessagesGetPollVotesRequestPredict) (MessagesVotesList, error) {
	var res MessagesVotesList
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleStickerSetsRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Uninstall   bool     `tl:",omitempty:flags:0,implicit"`
	Archive     bool     `tl:",omitempty:flags:1,implicit"`
	Unarchive   bool     `tl:",omitempty:flags:2,implicit"`
	Stickersets []InputStickerSet
}

func (*MessagesToggleStickerSetsRequestPredict) CRC() uint32 {
	return 0xb5052fea
}

func MessagesToggleStickerSets(ctx context.Context, m Requester, i MessagesToggleStickerSetsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDialogFiltersRequestPredict struct{}

func (*MessagesGetDialogFiltersRequestPredict) CRC() uint32 {
	return 0xefd48c89
}

func MessagesGetDialogFilters(ctx context.Context, m Requester, i MessagesGetDialogFiltersRequestPredict) (MessagesDialogFilters, error) {
	var res MessagesDialogFilters
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSuggestedDialogFiltersRequestPredict struct{}

func (*MessagesGetSuggestedDialogFiltersRequestPredict) CRC() uint32 {
	return 0xa29cd42c
}

func MessagesGetSuggestedDialogFilters(ctx context.Context, m Requester, i MessagesGetSuggestedDialogFiltersRequestPredict) ([]DialogFilterSuggested, error) {
	var res []DialogFilterSuggested
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdateDialogFilterRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	ID     int32
	Filter DialogFilter `tl:",omitempty:flags:0"`
}

func (*MessagesUpdateDialogFilterRequestPredict) CRC() uint32 {
	return 0x1ad4a04a
}

func MessagesUpdateDialogFilter(ctx context.Context, m Requester, i MessagesUpdateDialogFilterRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdateDialogFiltersOrderRequestPredict struct {
	Order []int32
}

func (*MessagesUpdateDialogFiltersOrderRequestPredict) CRC() uint32 {
	return 0xc563c1e4
}

func MessagesUpdateDialogFiltersOrder(ctx context.Context, m Requester, i MessagesUpdateDialogFiltersOrderRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOldFeaturedStickersRequestPredict struct {
	Offset int32
	Limit  int32
	Hash   int64
}

func (*MessagesGetOldFeaturedStickersRequestPredict) CRC() uint32 {
	return 0x7ed094a1
}

func MessagesGetOldFeaturedStickers(ctx context.Context, m Requester, i MessagesGetOldFeaturedStickersRequestPredict) (MessagesFeaturedStickers, error) {
	var res MessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRepliesRequestPredict struct {
	Peer       InputPeer
	MsgID      int32
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int64
}

func (*MessagesGetRepliesRequestPredict) CRC() uint32 {
	return 0x22ddd30c
}

func MessagesGetReplies(ctx context.Context, m Requester, i MessagesGetRepliesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDiscussionMessageRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*MessagesGetDiscussionMessageRequestPredict) CRC() uint32 {
	return 0x446972fd
}

func MessagesGetDiscussionMessage(ctx context.Context, m Requester, i MessagesGetDiscussionMessageRequestPredict) (MessagesDiscussionMessage, error) {
	var res MessagesDiscussionMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesReadDiscussionRequestPredict struct {
	Peer      InputPeer
	MsgID     int32
	ReadMaxID int32
}

func (*MessagesReadDiscussionRequestPredict) CRC() uint32 {
	return 0xf731a9f4
}

func MessagesReadDiscussion(ctx context.Context, m Requester, i MessagesReadDiscussionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUnpinAllMessagesRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     InputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesUnpinAllMessagesRequestPredict) CRC() uint32 {
	return 0xee22b9a8
}

func MessagesUnpinAllMessages(ctx context.Context, m Requester, i MessagesUnpinAllMessagesRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteChatRequestPredict struct {
	ChatID int64
}

func (*MessagesDeleteChatRequestPredict) CRC() uint32 {
	return 0x5bd0ee50
}

func MessagesDeleteChat(ctx context.Context, m Requester, i MessagesDeleteChatRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesDeletePhoneCallHistoryRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Revoke bool     `tl:",omitempty:flags:0,implicit"`
}

func (*MessagesDeletePhoneCallHistoryRequestPredict) CRC() uint32 {
	return 0xf9cbe409
}

func MessagesDeletePhoneCallHistory(ctx context.Context, m Requester, i MessagesDeletePhoneCallHistoryRequestPredict) (MessagesAffectedFoundMessages, error) {
	var res MessagesAffectedFoundMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckHistoryImportRequestPredict struct {
	ImportHead string
}

func (*MessagesCheckHistoryImportRequestPredict) CRC() uint32 {
	return 0x43fe19f3
}

func MessagesCheckHistoryImport(ctx context.Context, m Requester, i MessagesCheckHistoryImportRequestPredict) (MessagesHistoryImportParsed, error) {
	var res MessagesHistoryImportParsed
	return res, request(ctx, m, &i, &res)
}

type MessagesInitHistoryImportRequestPredict struct {
	Peer       InputPeer
	File       InputFile
	MediaCount int32
}

func (*MessagesInitHistoryImportRequestPredict) CRC() uint32 {
	return 0x34090c3b
}

func MessagesInitHistoryImport(ctx context.Context, m Requester, i MessagesInitHistoryImportRequestPredict) (MessagesHistoryImport, error) {
	var res MessagesHistoryImport
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadImportedMediaRequestPredict struct {
	Peer     InputPeer
	ImportID int64
	FileName string
	Media    InputMedia
}

func (*MessagesUploadImportedMediaRequestPredict) CRC() uint32 {
	return 0x2a862092
}

func MessagesUploadImportedMedia(ctx context.Context, m Requester, i MessagesUploadImportedMediaRequestPredict) (MessageMedia, error) {
	var res MessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesStartHistoryImportRequestPredict struct {
	Peer     InputPeer
	ImportID int64
}

func (*MessagesStartHistoryImportRequestPredict) CRC() uint32 {
	return 0xb43df344
}

func MessagesStartHistoryImport(ctx context.Context, m Requester, i MessagesStartHistoryImportRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExportedChatInvitesRequestPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Revoked    bool     `tl:",omitempty:flags:3,implicit"`
	Peer       InputPeer
	AdminID    InputUser
	OffsetDate *int32  `tl:",omitempty:flags:2"`
	OffsetLink *string `tl:",omitempty:flags:2"`
	Limit      int32
}

func (*MessagesGetExportedChatInvitesRequestPredict) CRC() uint32 {
	return 0xa2b5a3f6
}

func MessagesGetExportedChatInvites(ctx context.Context, m Requester, i MessagesGetExportedChatInvitesRequestPredict) (MessagesExportedChatInvites, error) {
	var res MessagesExportedChatInvites
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExportedChatInviteRequestPredict struct {
	Peer InputPeer
	Link string
}

func (*MessagesGetExportedChatInviteRequestPredict) CRC() uint32 {
	return 0x73746f5c
}

func MessagesGetExportedChatInvite(ctx context.Context, m Requester, i MessagesGetExportedChatInviteRequestPredict) (MessagesExportedChatInvite, error) {
	var res MessagesExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesEditExportedChatInviteRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Revoked       bool     `tl:",omitempty:flags:2,implicit"`
	Peer          InputPeer
	Link          string
	ExpireDate    *int32  `tl:",omitempty:flags:0"`
	UsageLimit    *int32  `tl:",omitempty:flags:1"`
	RequestNeeded *bool   `tl:",omitempty:flags:3"`
	Title         *string `tl:",omitempty:flags:4"`
}

func (*MessagesEditExportedChatInviteRequestPredict) CRC() uint32 {
	return 0xbdca2f75
}

func MessagesEditExportedChatInvite(ctx context.Context, m Requester, i MessagesEditExportedChatInviteRequestPredict) (MessagesExportedChatInvite, error) {
	var res MessagesExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteRevokedExportedChatInvitesRequestPredict struct {
	Peer    InputPeer
	AdminID InputUser
}

func (*MessagesDeleteRevokedExportedChatInvitesRequestPredict) CRC() uint32 {
	return 0x56987bd5
}

func MessagesDeleteRevokedExportedChatInvites(ctx context.Context, m Requester, i MessagesDeleteRevokedExportedChatInvitesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteExportedChatInviteRequestPredict struct {
	Peer InputPeer
	Link string
}

func (*MessagesDeleteExportedChatInviteRequestPredict) CRC() uint32 {
	return 0xd464a42b
}

func MessagesDeleteExportedChatInvite(ctx context.Context, m Requester, i MessagesDeleteExportedChatInviteRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAdminsWithInvitesRequestPredict struct {
	Peer InputPeer
}

func (*MessagesGetAdminsWithInvitesRequestPredict) CRC() uint32 {
	return 0x3920e6ef
}

func MessagesGetAdminsWithInvites(ctx context.Context, m Requester, i MessagesGetAdminsWithInvitesRequestPredict) (MessagesChatAdminsWithInvites, error) {
	var res MessagesChatAdminsWithInvites
	return res, request(ctx, m, &i, &res)
}

type MessagesGetChatInviteImportersRequestPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Requested  bool     `tl:",omitempty:flags:0,implicit"`
	Peer       InputPeer
	Link       *string `tl:",omitempty:flags:1"`
	Q          *string `tl:",omitempty:flags:2"`
	OffsetDate int32
	OffsetUser InputUser
	Limit      int32
}

func (*MessagesGetChatInviteImportersRequestPredict) CRC() uint32 {
	return 0xdf04dd4e
}

func MessagesGetChatInviteImporters(ctx context.Context, m Requester, i MessagesGetChatInviteImportersRequestPredict) (MessagesChatInviteImporters, error) {
	var res MessagesChatInviteImporters
	return res, request(ctx, m, &i, &res)
}

type MessagesSetHistoryTTLRequestPredict struct {
	Peer   InputPeer
	Period int32
}

func (*MessagesSetHistoryTTLRequestPredict) CRC() uint32 {
	return 0xb80e5fe4
}

func MessagesSetHistoryTTL(ctx context.Context, m Requester, i MessagesSetHistoryTTLRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckHistoryImportPeerRequestPredict struct {
	Peer InputPeer
}

func (*MessagesCheckHistoryImportPeerRequestPredict) CRC() uint32 {
	return 0x5dc60f03
}

func MessagesCheckHistoryImportPeer(ctx context.Context, m Requester, i MessagesCheckHistoryImportPeerRequestPredict) (MessagesCheckedHistoryImportPeer, error) {
	var res MessagesCheckedHistoryImportPeer
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatThemeRequestPredict struct {
	Peer     InputPeer
	Emoticon string
}

func (*MessagesSetChatThemeRequestPredict) CRC() uint32 {
	return 0xe63be13f
}

func MessagesSetChatTheme(ctx context.Context, m Requester, i MessagesSetChatThemeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageReadParticipantsRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*MessagesGetMessageReadParticipantsRequestPredict) CRC() uint32 {
	return 0x31c1c44f
}

func MessagesGetMessageReadParticipants(ctx context.Context, m Requester, i MessagesGetMessageReadParticipantsRequestPredict) ([]ReadParticipantDate, error) {
	var res []ReadParticipantDate
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchResultsCalendarRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Peer        InputPeer
	SavedPeerID InputPeer `tl:",omitempty:flags:2"`
	Filter      MessagesFilter
	OffsetID    int32
	OffsetDate  int32
}

func (*MessagesGetSearchResultsCalendarRequestPredict) CRC() uint32 {
	return 0x6aa3f6bd
}

func MessagesGetSearchResultsCalendar(ctx context.Context, m Requester, i MessagesGetSearchResultsCalendarRequestPredict) (MessagesSearchResultsCalendar, error) {
	var res MessagesSearchResultsCalendar
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchResultsPositionsRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Peer        InputPeer
	SavedPeerID InputPeer `tl:",omitempty:flags:2"`
	Filter      MessagesFilter
	OffsetID    int32
	Limit       int32
}

func (*MessagesGetSearchResultsPositionsRequestPredict) CRC() uint32 {
	return 0x9c7f2f10
}

func MessagesGetSearchResultsPositions(ctx context.Context, m Requester, i MessagesGetSearchResultsPositionsRequestPredict) (MessagesSearchResultsPositions, error) {
	var res MessagesSearchResultsPositions
	return res, request(ctx, m, &i, &res)
}

type MessagesHideChatJoinRequestRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Approved bool     `tl:",omitempty:flags:0,implicit"`
	Peer     InputPeer
	UserID   InputUser
}

func (*MessagesHideChatJoinRequestRequestPredict) CRC() uint32 {
	return 0x7fe7e815
}

func MessagesHideChatJoinRequest(ctx context.Context, m Requester, i MessagesHideChatJoinRequestRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesHideAllChatJoinRequestsRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Approved bool     `tl:",omitempty:flags:0,implicit"`
	Peer     InputPeer
	Link     *string `tl:",omitempty:flags:1"`
}

func (*MessagesHideAllChatJoinRequestsRequestPredict) CRC() uint32 {
	return 0xe085f4ea
}

func MessagesHideAllChatJoinRequests(ctx context.Context, m Requester, i MessagesHideAllChatJoinRequestsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleNoForwardsRequestPredict struct {
	Peer    InputPeer
	Enabled bool
}

func (*MessagesToggleNoForwardsRequestPredict) CRC() uint32 {
	return 0xb11eafa2
}

func MessagesToggleNoForwards(ctx context.Context, m Requester, i MessagesToggleNoForwardsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveDefaultSendAsRequestPredict struct {
	Peer   InputPeer
	SendAs InputPeer
}

func (*MessagesSaveDefaultSendAsRequestPredict) CRC() uint32 {
	return 0xccfddf96
}

func MessagesSaveDefaultSendAs(ctx context.Context, m Requester, i MessagesSaveDefaultSendAsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendReactionRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Big         bool     `tl:",omitempty:flags:1,implicit"`
	AddToRecent bool     `tl:",omitempty:flags:2,implicit"`
	Peer        InputPeer
	MsgID       int32
	Reaction    []Reaction `tl:",omitempty:flags:0"`
}

func (*MessagesSendReactionRequestPredict) CRC() uint32 {
	return 0xd30d78d4
}

func MessagesSendReaction(ctx context.Context, m Requester, i MessagesSendReactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesReactionsRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesGetMessagesReactionsRequestPredict) CRC() uint32 {
	return 0x8bba90e6
}

func MessagesGetMessagesReactions(ctx context.Context, m Requester, i MessagesGetMessagesReactionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageReactionsListRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     InputPeer
	ID       int32
	Reaction Reaction `tl:",omitempty:flags:0"`
	Offset   *string  `tl:",omitempty:flags:1"`
	Limit    int32
}

func (*MessagesGetMessageReactionsListRequestPredict) CRC() uint32 {
	return 0x461b3f48
}

func MessagesGetMessageReactionsList(ctx context.Context, m Requester, i MessagesGetMessageReactionsListRequestPredict) (MessagesMessageReactionsList, error) {
	var res MessagesMessageReactionsList
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatAvailableReactionsRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Peer               InputPeer
	AvailableReactions ChatReactions
	ReactionsLimit     *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesSetChatAvailableReactionsRequestPredict) CRC() uint32 {
	return 0x5a150bd4
}

func MessagesSetChatAvailableReactions(ctx context.Context, m Requester, i MessagesSetChatAvailableReactionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAvailableReactionsRequestPredict struct {
	Hash int32
}

func (*MessagesGetAvailableReactionsRequestPredict) CRC() uint32 {
	return 0x18dea0ac
}

func MessagesGetAvailableReactions(ctx context.Context, m Requester, i MessagesGetAvailableReactionsRequestPredict) (MessagesAvailableReactions, error) {
	var res MessagesAvailableReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesSetDefaultReactionRequestPredict struct {
	Reaction Reaction
}

func (*MessagesSetDefaultReactionRequestPredict) CRC() uint32 {
	return 0x4f47a016
}

func MessagesSetDefaultReaction(ctx context.Context, m Requester, i MessagesSetDefaultReactionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesTranslateTextRequestPredict struct {
	_      struct{}           `tl:"flags,bitflag"`
	Peer   InputPeer          `tl:",omitempty:flags:0"`
	ID     []int32            `tl:",omitempty:flags:0"`
	Text   []TextWithEntities `tl:",omitempty:flags:1"`
	ToLang string
}

func (*MessagesTranslateTextRequestPredict) CRC() uint32 {
	return 0x63183030
}

func MessagesTranslateText(ctx context.Context, m Requester, i MessagesTranslateTextRequestPredict) (MessagesTranslatedText, error) {
	var res MessagesTranslatedText
	return res, request(ctx, m, &i, &res)
}

type MessagesGetUnreadReactionsRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      InputPeer
	TopMsgID  *int32 `tl:",omitempty:flags:0"`
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
}

func (*MessagesGetUnreadReactionsRequestPredict) CRC() uint32 {
	return 0x3223495b
}

func MessagesGetUnreadReactions(ctx context.Context, m Requester, i MessagesGetUnreadReactionsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadReactionsRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Peer     InputPeer
	TopMsgID *int32 `tl:",omitempty:flags:0"`
}

func (*MessagesReadReactionsRequestPredict) CRC() uint32 {
	return 0x54aa7f8e
}

func MessagesReadReactions(ctx context.Context, m Requester, i MessagesReadReactionsRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchSentMediaRequestPredict struct {
	Q      string
	Filter MessagesFilter
	Limit  int32
}

func (*MessagesSearchSentMediaRequestPredict) CRC() uint32 {
	return 0x107e31a0
}

func MessagesSearchSentMedia(ctx context.Context, m Requester, i MessagesSearchSentMediaRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachMenuBotsRequestPredict struct {
	Hash int64
}

func (*MessagesGetAttachMenuBotsRequestPredict) CRC() uint32 {
	return 0x16fcc2cb
}

func MessagesGetAttachMenuBots(ctx context.Context, m Requester, i MessagesGetAttachMenuBotsRequestPredict) (AttachMenuBots, error) {
	var res AttachMenuBots
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachMenuBotRequestPredict struct {
	Bot InputUser
}

func (*MessagesGetAttachMenuBotRequestPredict) CRC() uint32 {
	return 0x77216192
}

func MessagesGetAttachMenuBot(ctx context.Context, m Requester, i MessagesGetAttachMenuBotRequestPredict) (AttachMenuBotsBot, error) {
	var res AttachMenuBotsBot
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleBotInAttachMenuRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	WriteAllowed bool     `tl:",omitempty:flags:0,implicit"`
	Bot          InputUser
	Enabled      bool
}

func (*MessagesToggleBotInAttachMenuRequestPredict) CRC() uint32 {
	return 0x69f59d69
}

func MessagesToggleBotInAttachMenu(ctx context.Context, m Requester, i MessagesToggleBotInAttachMenuRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestWebViewRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	FromBotMenu bool     `tl:",omitempty:flags:4,implicit"`
	Silent      bool     `tl:",omitempty:flags:5,implicit"`
	Compact     bool     `tl:",omitempty:flags:7,implicit"`
	Peer        InputPeer
	Bot         InputUser
	URL         *string  `tl:",omitempty:flags:1"`
	StartParam  *string  `tl:",omitempty:flags:3"`
	ThemeParams DataJSON `tl:",omitempty:flags:2"`
	Platform    string
	ReplyTo     InputReplyTo `tl:",omitempty:flags:0"`
	SendAs      InputPeer    `tl:",omitempty:flags:13"`
}

func (*MessagesRequestWebViewRequestPredict) CRC() uint32 {
	return 0x269dc2c1
}

func MessagesRequestWebView(ctx context.Context, m Requester, i MessagesRequestWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesProlongWebViewRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Silent  bool     `tl:",omitempty:flags:5,implicit"`
	Peer    InputPeer
	Bot     InputUser
	QueryID int64
	ReplyTo InputReplyTo `tl:",omitempty:flags:0"`
	SendAs  InputPeer    `tl:",omitempty:flags:13"`
}

func (*MessagesProlongWebViewRequestPredict) CRC() uint32 {
	return 0xb0d81a83
}

func MessagesProlongWebView(ctx context.Context, m Requester, i MessagesProlongWebViewRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestSimpleWebViewRequestPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	FromSwitchWebview bool     `tl:",omitempty:flags:1,implicit"`
	FromSideMenu      bool     `tl:",omitempty:flags:2,implicit"`
	Compact           bool     `tl:",omitempty:flags:7,implicit"`
	Bot               InputUser
	URL               *string  `tl:",omitempty:flags:3"`
	StartParam        *string  `tl:",omitempty:flags:4"`
	ThemeParams       DataJSON `tl:",omitempty:flags:0"`
	Platform          string
}

func (*MessagesRequestSimpleWebViewRequestPredict) CRC() uint32 {
	return 0x413a3e73
}

func MessagesRequestSimpleWebView(ctx context.Context, m Requester, i MessagesRequestSimpleWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesSendWebViewResultMessageRequestPredict struct {
	BotQueryID string
	Result     InputBotInlineResult
}

func (*MessagesSendWebViewResultMessageRequestPredict) CRC() uint32 {
	return 0xa4314f5
}

func MessagesSendWebViewResultMessage(ctx context.Context, m Requester, i MessagesSendWebViewResultMessageRequestPredict) (WebViewMessageSent, error) {
	var res WebViewMessageSent
	return res, request(ctx, m, &i, &res)
}

type MessagesSendWebViewDataRequestPredict struct {
	Bot        InputUser
	RandomID   int64
	ButtonText string
	Data       string
}

func (*MessagesSendWebViewDataRequestPredict) CRC() uint32 {
	return 0xdc0242c8
}

func MessagesSendWebViewData(ctx context.Context, m Requester, i MessagesSendWebViewDataRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesTranscribeAudioRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*MessagesTranscribeAudioRequestPredict) CRC() uint32 {
	return 0x269e9a49
}

func MessagesTranscribeAudio(ctx context.Context, m Requester, i MessagesTranscribeAudioRequestPredict) (MessagesTranscribedAudio, error) {
	var res MessagesTranscribedAudio
	return res, request(ctx, m, &i, &res)
}

type MessagesRateTranscribedAudioRequestPredict struct {
	Peer            InputPeer
	MsgID           int32
	TranscriptionID int64
	Good            bool
}

func (*MessagesRateTranscribedAudioRequestPredict) CRC() uint32 {
	return 0x7f1d072f
}

func MessagesRateTranscribedAudio(ctx context.Context, m Requester, i MessagesRateTranscribedAudioRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetCustomEmojiDocumentsRequestPredict struct {
	DocumentID []int64
}

func (*MessagesGetCustomEmojiDocumentsRequestPredict) CRC() uint32 {
	return 0xd9ab0f54
}

func MessagesGetCustomEmojiDocuments(ctx context.Context, m Requester, i MessagesGetCustomEmojiDocumentsRequestPredict) ([]Document, error) {
	var res []Document
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStickersRequestPredict struct {
	Hash int64
}

func (*MessagesGetEmojiStickersRequestPredict) CRC() uint32 {
	return 0xfbfca18f
}

func MessagesGetEmojiStickers(ctx context.Context, m Requester, i MessagesGetEmojiStickersRequestPredict) (MessagesAllStickers, error) {
	var res MessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFeaturedEmojiStickersRequestPredict struct {
	Hash int64
}

func (*MessagesGetFeaturedEmojiStickersRequestPredict) CRC() uint32 {
	return 0xecf6736
}

func MessagesGetFeaturedEmojiStickers(ctx context.Context, m Requester, i MessagesGetFeaturedEmojiStickersRequestPredict) (MessagesFeaturedStickers, error) {
	var res MessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesReportReactionRequestPredict struct {
	Peer         InputPeer
	ID           int32
	ReactionPeer InputPeer
}

func (*MessagesReportReactionRequestPredict) CRC() uint32 {
	return 0x3f64c076
}

func MessagesReportReaction(ctx context.Context, m Requester, i MessagesReportReactionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetTopReactionsRequestPredict struct {
	Limit int32
	Hash  int64
}

func (*MessagesGetTopReactionsRequestPredict) CRC() uint32 {
	return 0xbb8125ba
}

func MessagesGetTopReactions(ctx context.Context, m Requester, i MessagesGetTopReactionsRequestPredict) (MessagesReactions, error) {
	var res MessagesReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentReactionsRequestPredict struct {
	Limit int32
	Hash  int64
}

func (*MessagesGetRecentReactionsRequestPredict) CRC() uint32 {
	return 0x39461db2
}

func MessagesGetRecentReactions(ctx context.Context, m Requester, i MessagesGetRecentReactionsRequestPredict) (MessagesReactions, error) {
	var res MessagesReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesClearRecentReactionsRequestPredict struct{}

func (*MessagesClearRecentReactionsRequestPredict) CRC() uint32 {
	return 0x9dfeefb4
}

func MessagesClearRecentReactions(ctx context.Context, m Requester, i MessagesClearRecentReactionsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExtendedMediaRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*MessagesGetExtendedMediaRequestPredict) CRC() uint32 {
	return 0x84f80814
}

func MessagesGetExtendedMedia(ctx context.Context, m Requester, i MessagesGetExtendedMediaRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSetDefaultHistoryTTLRequestPredict struct {
	Period int32
}

func (*MessagesSetDefaultHistoryTTLRequestPredict) CRC() uint32 {
	return 0x9eb51445
}

func MessagesSetDefaultHistoryTTL(ctx context.Context, m Requester, i MessagesSetDefaultHistoryTTLRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDefaultHistoryTTLRequestPredict struct{}

func (*MessagesGetDefaultHistoryTTLRequestPredict) CRC() uint32 {
	return 0x658b7188
}

func MessagesGetDefaultHistoryTTL(ctx context.Context, m Requester, i MessagesGetDefaultHistoryTTLRequestPredict) (DefaultHistoryTTL, error) {
	var res DefaultHistoryTTL
	return res, request(ctx, m, &i, &res)
}

type MessagesSendBotRequestedPeerRequestPredict struct {
	Peer           InputPeer
	MsgID          int32
	ButtonID       int32
	RequestedPeers []InputPeer
}

func (*MessagesSendBotRequestedPeerRequestPredict) CRC() uint32 {
	return 0x91b2d060
}

func MessagesSendBotRequestedPeer(ctx context.Context, m Requester, i MessagesSendBotRequestedPeerRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiGroupsRequestPredict struct {
	Hash int32
}

func (*MessagesGetEmojiGroupsRequestPredict) CRC() uint32 {
	return 0x7488ce5b
}

func MessagesGetEmojiGroups(ctx context.Context, m Requester, i MessagesGetEmojiGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStatusGroupsRequestPredict struct {
	Hash int32
}

func (*MessagesGetEmojiStatusGroupsRequestPredict) CRC() uint32 {
	return 0x2ecd56cd
}

func MessagesGetEmojiStatusGroups(ctx context.Context, m Requester, i MessagesGetEmojiStatusGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiProfilePhotoGroupsRequestPredict struct {
	Hash int32
}

func (*MessagesGetEmojiProfilePhotoGroupsRequestPredict) CRC() uint32 {
	return 0x21a548f3
}

func MessagesGetEmojiProfilePhotoGroups(ctx context.Context, m Requester, i MessagesGetEmojiProfilePhotoGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchCustomEmojiRequestPredict struct {
	Emoticon string
	Hash     int64
}

func (*MessagesSearchCustomEmojiRequestPredict) CRC() uint32 {
	return 0x2c11c0d7
}

func MessagesSearchCustomEmoji(ctx context.Context, m Requester, i MessagesSearchCustomEmojiRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type MessagesTogglePeerTranslationsRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Disabled bool     `tl:",omitempty:flags:0,implicit"`
	Peer     InputPeer
}

func (*MessagesTogglePeerTranslationsRequestPredict) CRC() uint32 {
	return 0xe47cb579
}

func MessagesTogglePeerTranslations(ctx context.Context, m Requester, i MessagesTogglePeerTranslationsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetBotAppRequestPredict struct {
	App  InputBotApp
	Hash int64
}

func (*MessagesGetBotAppRequestPredict) CRC() uint32 {
	return 0x34fdc5c3
}

func MessagesGetBotApp(ctx context.Context, m Requester, i MessagesGetBotAppRequestPredict) (MessagesBotApp, error) {
	var res MessagesBotApp
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestAppWebViewRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	WriteAllowed bool     `tl:",omitempty:flags:0,implicit"`
	Compact      bool     `tl:",omitempty:flags:7,implicit"`
	Peer         InputPeer
	App          InputBotApp
	StartParam   *string  `tl:",omitempty:flags:1"`
	ThemeParams  DataJSON `tl:",omitempty:flags:2"`
	Platform     string
}

func (*MessagesRequestAppWebViewRequestPredict) CRC() uint32 {
	return 0x53618bce
}

func MessagesRequestAppWebView(ctx context.Context, m Requester, i MessagesRequestAppWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatWallPaperRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ForBoth   bool     `tl:",omitempty:flags:3,implicit"`
	Revert    bool     `tl:",omitempty:flags:4,implicit"`
	Peer      InputPeer
	Wallpaper InputWallPaper    `tl:",omitempty:flags:0"`
	Settings  WallPaperSettings `tl:",omitempty:flags:2"`
	ID        *int32            `tl:",omitempty:flags:1"`
}

func (*MessagesSetChatWallPaperRequestPredict) CRC() uint32 {
	return 0x8ffacae1
}

func MessagesSetChatWallPaper(ctx context.Context, m Requester, i MessagesSetChatWallPaperRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchEmojiStickerSetsRequestPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ExcludeFeatured bool     `tl:",omitempty:flags:0,implicit"`
	Q               string
	Hash            int64
}

func (*MessagesSearchEmojiStickerSetsRequestPredict) CRC() uint32 {
	return 0x92b4494c
}

func MessagesSearchEmojiStickerSets(ctx context.Context, m Requester, i MessagesSearchEmojiStickerSetsRequestPredict) (MessagesFoundStickerSets, error) {
	var res MessagesFoundStickerSets
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedDialogsRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ExcludePinned bool     `tl:",omitempty:flags:0,implicit"`
	OffsetDate    int32
	OffsetID      int32
	OffsetPeer    InputPeer
	Limit         int32
	Hash          int64
}

func (*MessagesGetSavedDialogsRequestPredict) CRC() uint32 {
	return 0x5381d21a
}

func MessagesGetSavedDialogs(ctx context.Context, m Requester, i MessagesGetSavedDialogsRequestPredict) (MessagesSavedDialogs, error) {
	var res MessagesSavedDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedHistoryRequestPredict struct {
	Peer       InputPeer
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int64
}

func (*MessagesGetSavedHistoryRequestPredict) CRC() uint32 {
	return 0x3d9a414d
}

func MessagesGetSavedHistory(ctx context.Context, m Requester, i MessagesGetSavedHistoryRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteSavedHistoryRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Peer    InputPeer
	MaxID   int32
	MinDate *int32 `tl:",omitempty:flags:2"`
	MaxDate *int32 `tl:",omitempty:flags:3"`
}

func (*MessagesDeleteSavedHistoryRequestPredict) CRC() uint32 {
	return 0x6e98102b
}

func MessagesDeleteSavedHistory(ctx context.Context, m Requester, i MessagesDeleteSavedHistoryRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPinnedSavedDialogsRequestPredict struct{}

func (*MessagesGetPinnedSavedDialogsRequestPredict) CRC() uint32 {
	return 0xd63d94e0
}

func MessagesGetPinnedSavedDialogs(ctx context.Context, m Requester, i MessagesGetPinnedSavedDialogsRequestPredict) (MessagesSavedDialogs, error) {
	var res MessagesSavedDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleSavedDialogPinRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Pinned bool     `tl:",omitempty:flags:0,implicit"`
	Peer   InputDialogPeer
}

func (*MessagesToggleSavedDialogPinRequestPredict) CRC() uint32 {
	return 0xac81bbde
}

func MessagesToggleSavedDialogPin(ctx context.Context, m Requester, i MessagesToggleSavedDialogPinRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderPinnedSavedDialogsRequestPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Force bool     `tl:",omitempty:flags:0,implicit"`
	Order []InputDialogPeer
}

func (*MessagesReorderPinnedSavedDialogsRequestPredict) CRC() uint32 {
	return 0x8b716587
}

func MessagesReorderPinnedSavedDialogs(ctx context.Context, m Requester, i MessagesReorderPinnedSavedDialogsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedReactionTagsRequestPredict struct {
	_    struct{}  `tl:"flags,bitflag"`
	Peer InputPeer `tl:",omitempty:flags:0"`
	Hash int64
}

func (*MessagesGetSavedReactionTagsRequestPredict) CRC() uint32 {
	return 0x3637e05b
}

func MessagesGetSavedReactionTags(ctx context.Context, m Requester, i MessagesGetSavedReactionTagsRequestPredict) (MessagesSavedReactionTags, error) {
	var res MessagesSavedReactionTags
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdateSavedReactionTagRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Reaction Reaction
	Title    *string `tl:",omitempty:flags:0"`
}

func (*MessagesUpdateSavedReactionTagRequestPredict) CRC() uint32 {
	return 0x60297dec
}

func MessagesUpdateSavedReactionTag(ctx context.Context, m Requester, i MessagesUpdateSavedReactionTagRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDefaultTagReactionsRequestPredict struct {
	Hash int64
}

func (*MessagesGetDefaultTagReactionsRequestPredict) CRC() uint32 {
	return 0xbdf93428
}

func MessagesGetDefaultTagReactions(ctx context.Context, m Requester, i MessagesGetDefaultTagReactionsRequestPredict) (MessagesReactions, error) {
	var res MessagesReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOutboxReadDateRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*MessagesGetOutboxReadDateRequestPredict) CRC() uint32 {
	return 0x8c4bfe5d
}

func MessagesGetOutboxReadDate(ctx context.Context, m Requester, i MessagesGetOutboxReadDateRequestPredict) (OutboxReadDate, error) {
	var res OutboxReadDate
	return res, request(ctx, m, &i, &res)
}

type MessagesGetQuickRepliesRequestPredict struct {
	Hash int64
}

func (*MessagesGetQuickRepliesRequestPredict) CRC() uint32 {
	return 0xd483f2a8
}

func MessagesGetQuickReplies(ctx context.Context, m Requester, i MessagesGetQuickRepliesRequestPredict) (MessagesQuickReplies, error) {
	var res MessagesQuickReplies
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderQuickRepliesRequestPredict struct {
	Order []int32
}

func (*MessagesReorderQuickRepliesRequestPredict) CRC() uint32 {
	return 0x60331907
}

func MessagesReorderQuickReplies(ctx context.Context, m Requester, i MessagesReorderQuickRepliesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckQuickReplyShortcutRequestPredict struct {
	Shortcut string
}

func (*MessagesCheckQuickReplyShortcutRequestPredict) CRC() uint32 {
	return 0xf1d0fbd3
}

func MessagesCheckQuickReplyShortcut(ctx context.Context, m Requester, i MessagesCheckQuickReplyShortcutRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesEditQuickReplyShortcutRequestPredict struct {
	ShortcutID int32
	Shortcut   string
}

func (*MessagesEditQuickReplyShortcutRequestPredict) CRC() uint32 {
	return 0x5c003cef
}

func MessagesEditQuickReplyShortcut(ctx context.Context, m Requester, i MessagesEditQuickReplyShortcutRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteQuickReplyShortcutRequestPredict struct {
	ShortcutID int32
}

func (*MessagesDeleteQuickReplyShortcutRequestPredict) CRC() uint32 {
	return 0x3cc04740
}

func MessagesDeleteQuickReplyShortcut(ctx context.Context, m Requester, i MessagesDeleteQuickReplyShortcutRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetQuickReplyMessagesRequestPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	ShortcutID int32
	ID         []int32 `tl:",omitempty:flags:0"`
	Hash       int64
}

func (*MessagesGetQuickReplyMessagesRequestPredict) CRC() uint32 {
	return 0x94a495c3
}

func MessagesGetQuickReplyMessages(ctx context.Context, m Requester, i MessagesGetQuickReplyMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendQuickReplyMessagesRequestPredict struct {
	Peer       InputPeer
	ShortcutID int32
	ID         []int32
	RandomID   []int64
}

func (*MessagesSendQuickReplyMessagesRequestPredict) CRC() uint32 {
	return 0x6c750de1
}

func MessagesSendQuickReplyMessages(ctx context.Context, m Requester, i MessagesSendQuickReplyMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteQuickReplyMessagesRequestPredict struct {
	ShortcutID int32
	ID         []int32
}

func (*MessagesDeleteQuickReplyMessagesRequestPredict) CRC() uint32 {
	return 0xe105e910
}

func MessagesDeleteQuickReplyMessages(ctx context.Context, m Requester, i MessagesDeleteQuickReplyMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleDialogFilterTagsRequestPredict struct {
	Enabled bool
}

func (*MessagesToggleDialogFilterTagsRequestPredict) CRC() uint32 {
	return 0xfd2dda49
}

func MessagesToggleDialogFilterTags(ctx context.Context, m Requester, i MessagesToggleDialogFilterTagsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMyStickersRequestPredict struct {
	OffsetID int64
	Limit    int32
}

func (*MessagesGetMyStickersRequestPredict) CRC() uint32 {
	return 0xd0b5e1fc
}

func MessagesGetMyStickers(ctx context.Context, m Requester, i MessagesGetMyStickersRequestPredict) (MessagesMyStickers, error) {
	var res MessagesMyStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStickerGroupsRequestPredict struct {
	Hash int32
}

func (*MessagesGetEmojiStickerGroupsRequestPredict) CRC() uint32 {
	return 0x1dd840f5
}

func MessagesGetEmojiStickerGroups(ctx context.Context, m Requester, i MessagesGetEmojiStickerGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAvailableEffectsRequestPredict struct {
	Hash int32
}

func (*MessagesGetAvailableEffectsRequestPredict) CRC() uint32 {
	return 0xdea20a39
}

func MessagesGetAvailableEffects(ctx context.Context, m Requester, i MessagesGetAvailableEffectsRequestPredict) (MessagesAvailableEffects, error) {
	var res MessagesAvailableEffects
	return res, request(ctx, m, &i, &res)
}

type MessagesEditFactCheckRequestPredict struct {
	Peer  InputPeer
	MsgID int32
	Text  TextWithEntities
}

func (*MessagesEditFactCheckRequestPredict) CRC() uint32 {
	return 0x589ee75
}

func MessagesEditFactCheck(ctx context.Context, m Requester, i MessagesEditFactCheckRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteFactCheckRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*MessagesDeleteFactCheckRequestPredict) CRC() uint32 {
	return 0xd1da940c
}

func MessagesDeleteFactCheck(ctx context.Context, m Requester, i MessagesDeleteFactCheckRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFactCheckRequestPredict struct {
	Peer  InputPeer
	MsgID []int32
}

func (*MessagesGetFactCheckRequestPredict) CRC() uint32 {
	return 0xb9cdc5ee
}

func MessagesGetFactCheck(ctx context.Context, m Requester, i MessagesGetFactCheckRequestPredict) ([]FactCheck, error) {
	var res []FactCheck
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestMainWebViewRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Compact     bool     `tl:",omitempty:flags:7,implicit"`
	Peer        InputPeer
	Bot         InputUser
	StartParam  *string  `tl:",omitempty:flags:1"`
	ThemeParams DataJSON `tl:",omitempty:flags:0"`
	Platform    string
}

func (*MessagesRequestMainWebViewRequestPredict) CRC() uint32 {
	return 0xc9e01e7b
}

func MessagesRequestMainWebView(ctx context.Context, m Requester, i MessagesRequestMainWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPaymentFormRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Invoice     InputInvoice
	ThemeParams DataJSON `tl:",omitempty:flags:0"`
}

func (*PaymentsGetPaymentFormRequestPredict) CRC() uint32 {
	return 0x37148dbb
}

func PaymentsGetPaymentForm(ctx context.Context, m Requester, i PaymentsGetPaymentFormRequestPredict) (PaymentsPaymentForm, error) {
	var res PaymentsPaymentForm
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPaymentReceiptRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*PaymentsGetPaymentReceiptRequestPredict) CRC() uint32 {
	return 0x2478d1cc
}

func PaymentsGetPaymentReceipt(ctx context.Context, m Requester, i PaymentsGetPaymentReceiptRequestPredict) (PaymentsPaymentReceipt, error) {
	var res PaymentsPaymentReceipt
	return res, request(ctx, m, &i, &res)
}

type PaymentsValidateRequestedInfoRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Save    bool     `tl:",omitempty:flags:0,implicit"`
	Invoice InputInvoice
	Info    PaymentRequestedInfo
}

func (*PaymentsValidateRequestedInfoRequestPredict) CRC() uint32 {
	return 0xb6c8f12b
}

func PaymentsValidateRequestedInfo(ctx context.Context, m Requester, i PaymentsValidateRequestedInfoRequestPredict) (PaymentsValidatedRequestedInfo, error) {
	var res PaymentsValidatedRequestedInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsSendPaymentFormRequestPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	FormID           int64
	Invoice          InputInvoice
	RequestedInfoID  *string `tl:",omitempty:flags:0"`
	ShippingOptionID *string `tl:",omitempty:flags:1"`
	Credentials      InputPaymentCredentials
	TipAmount        *int64 `tl:",omitempty:flags:2"`
}

func (*PaymentsSendPaymentFormRequestPredict) CRC() uint32 {
	return 0x2d03522f
}

func PaymentsSendPaymentForm(ctx context.Context, m Requester, i PaymentsSendPaymentFormRequestPredict) (PaymentsPaymentResult, error) {
	var res PaymentsPaymentResult
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetSavedInfoRequestPredict struct{}

func (*PaymentsGetSavedInfoRequestPredict) CRC() uint32 {
	return 0x227d824b
}

func PaymentsGetSavedInfo(ctx context.Context, m Requester, i PaymentsGetSavedInfoRequestPredict) (PaymentsSavedInfo, error) {
	var res PaymentsSavedInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsClearSavedInfoRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Credentials bool     `tl:",omitempty:flags:0,implicit"`
	Info        bool     `tl:",omitempty:flags:1,implicit"`
}

func (*PaymentsClearSavedInfoRequestPredict) CRC() uint32 {
	return 0xd83d70c1
}

func PaymentsClearSavedInfo(ctx context.Context, m Requester, i PaymentsClearSavedInfoRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetBankCardDataRequestPredict struct {
	Number string
}

func (*PaymentsGetBankCardDataRequestPredict) CRC() uint32 {
	return 0x2e79d779
}

func PaymentsGetBankCardData(ctx context.Context, m Requester, i PaymentsGetBankCardDataRequestPredict) (PaymentsBankCardData, error) {
	var res PaymentsBankCardData
	return res, request(ctx, m, &i, &res)
}

type PaymentsExportInvoiceRequestPredict struct {
	InvoiceMedia InputMedia
}

func (*PaymentsExportInvoiceRequestPredict) CRC() uint32 {
	return 0xf91b065
}

func PaymentsExportInvoice(ctx context.Context, m Requester, i PaymentsExportInvoiceRequestPredict) (PaymentsExportedInvoice, error) {
	var res PaymentsExportedInvoice
	return res, request(ctx, m, &i, &res)
}

type PaymentsAssignAppStoreTransactionRequestPredict struct {
	Receipt []byte
	Purpose InputStorePaymentPurpose
}

func (*PaymentsAssignAppStoreTransactionRequestPredict) CRC() uint32 {
	return 0x80ed747d
}

func PaymentsAssignAppStoreTransaction(ctx context.Context, m Requester, i PaymentsAssignAppStoreTransactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsAssignPlayMarketTransactionRequestPredict struct {
	Receipt DataJSON
	Purpose InputStorePaymentPurpose
}

func (*PaymentsAssignPlayMarketTransactionRequestPredict) CRC() uint32 {
	return 0xdffd50d3
}

func PaymentsAssignPlayMarketTransaction(ctx context.Context, m Requester, i PaymentsAssignPlayMarketTransactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsCanPurchasePremiumRequestPredict struct {
	Purpose InputStorePaymentPurpose
}

func (*PaymentsCanPurchasePremiumRequestPredict) CRC() uint32 {
	return 0x9fc19eb6
}

func PaymentsCanPurchasePremium(ctx context.Context, m Requester, i PaymentsCanPurchasePremiumRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPremiumGiftCodeOptionsRequestPredict struct {
	_         struct{}  `tl:"flags,bitflag"`
	BoostPeer InputPeer `tl:",omitempty:flags:0"`
}

func (*PaymentsGetPremiumGiftCodeOptionsRequestPredict) CRC() uint32 {
	return 0x2757ba54
}

func PaymentsGetPremiumGiftCodeOptions(ctx context.Context, m Requester, i PaymentsGetPremiumGiftCodeOptionsRequestPredict) ([]PremiumGiftCodeOption, error) {
	var res []PremiumGiftCodeOption
	return res, request(ctx, m, &i, &res)
}

type PaymentsCheckGiftCodeRequestPredict struct {
	Slug string
}

func (*PaymentsCheckGiftCodeRequestPredict) CRC() uint32 {
	return 0x8e51b4c1
}

func PaymentsCheckGiftCode(ctx context.Context, m Requester, i PaymentsCheckGiftCodeRequestPredict) (PaymentsCheckedGiftCode, error) {
	var res PaymentsCheckedGiftCode
	return res, request(ctx, m, &i, &res)
}

type PaymentsApplyGiftCodeRequestPredict struct {
	Slug string
}

func (*PaymentsApplyGiftCodeRequestPredict) CRC() uint32 {
	return 0xf6e26854
}

func PaymentsApplyGiftCode(ctx context.Context, m Requester, i PaymentsApplyGiftCodeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetGiveawayInfoRequestPredict struct {
	Peer  InputPeer
	MsgID int32
}

func (*PaymentsGetGiveawayInfoRequestPredict) CRC() uint32 {
	return 0xf4239425
}

func PaymentsGetGiveawayInfo(ctx context.Context, m Requester, i PaymentsGetGiveawayInfoRequestPredict) (PaymentsGiveawayInfo, error) {
	var res PaymentsGiveawayInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsLaunchPrepaidGiveawayRequestPredict struct {
	Peer       InputPeer
	GiveawayID int64
	Purpose    InputStorePaymentPurpose
}

func (*PaymentsLaunchPrepaidGiveawayRequestPredict) CRC() uint32 {
	return 0x5ff58f20
}

func PaymentsLaunchPrepaidGiveaway(ctx context.Context, m Requester, i PaymentsLaunchPrepaidGiveawayRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsTopupOptionsRequestPredict struct{}

func (*PaymentsGetStarsTopupOptionsRequestPredict) CRC() uint32 {
	return 0xc00ec7d3
}

func PaymentsGetStarsTopupOptions(ctx context.Context, m Requester, i PaymentsGetStarsTopupOptionsRequestPredict) ([]StarsTopupOption, error) {
	var res []StarsTopupOption
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsStatusRequestPredict struct {
	Peer InputPeer
}

func (*PaymentsGetStarsStatusRequestPredict) CRC() uint32 {
	return 0x104fcfa7
}

func PaymentsGetStarsStatus(ctx context.Context, m Requester, i PaymentsGetStarsStatusRequestPredict) (PaymentsStarsStatus, error) {
	var res PaymentsStarsStatus
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsTransactionsRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Inbound   bool     `tl:",omitempty:flags:0,implicit"`
	Outbound  bool     `tl:",omitempty:flags:1,implicit"`
	Ascending bool     `tl:",omitempty:flags:2,implicit"`
	Peer      InputPeer
	Offset    string
	Limit     int32
}

func (*PaymentsGetStarsTransactionsRequestPredict) CRC() uint32 {
	return 0x97938d5a
}

func PaymentsGetStarsTransactions(ctx context.Context, m Requester, i PaymentsGetStarsTransactionsRequestPredict) (PaymentsStarsStatus, error) {
	var res PaymentsStarsStatus
	return res, request(ctx, m, &i, &res)
}

type PaymentsSendStarsFormRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	FormID  int64
	Invoice InputInvoice
}

func (*PaymentsSendStarsFormRequestPredict) CRC() uint32 {
	return 0x2bb731d
}

func PaymentsSendStarsForm(ctx context.Context, m Requester, i PaymentsSendStarsFormRequestPredict) (PaymentsPaymentResult, error) {
	var res PaymentsPaymentResult
	return res, request(ctx, m, &i, &res)
}

type PaymentsRefundStarsChargeRequestPredict struct {
	UserID   InputUser
	ChargeID string
}

func (*PaymentsRefundStarsChargeRequestPredict) CRC() uint32 {
	return 0x25ae8f4a
}

func PaymentsRefundStarsCharge(ctx context.Context, m Requester, i PaymentsRefundStarsChargeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsRevenueStatsRequestPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Dark bool     `tl:",omitempty:flags:0,implicit"`
	Peer InputPeer
}

func (*PaymentsGetStarsRevenueStatsRequestPredict) CRC() uint32 {
	return 0xd91ffad6
}

func PaymentsGetStarsRevenueStats(ctx context.Context, m Requester, i PaymentsGetStarsRevenueStatsRequestPredict) (PaymentsStarsRevenueStats, error) {
	var res PaymentsStarsRevenueStats
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsRevenueWithdrawalURLRequestPredict struct {
	Peer     InputPeer
	Stars    int64
	Password InputCheckPasswordSRP
}

func (*PaymentsGetStarsRevenueWithdrawalURLRequestPredict) CRC() uint32 {
	return 0x13bbe8b3
}

func PaymentsGetStarsRevenueWithdrawalURL(ctx context.Context, m Requester, i PaymentsGetStarsRevenueWithdrawalURLRequestPredict) (PaymentsStarsRevenueWithdrawalURL, error) {
	var res PaymentsStarsRevenueWithdrawalURL
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsRevenueAdsAccountURLRequestPredict struct {
	Peer InputPeer
}

func (*PaymentsGetStarsRevenueAdsAccountURLRequestPredict) CRC() uint32 {
	return 0xd1d7efc5
}

func PaymentsGetStarsRevenueAdsAccountURL(ctx context.Context, m Requester, i PaymentsGetStarsRevenueAdsAccountURLRequestPredict) (PaymentsStarsRevenueAdsAccountURL, error) {
	var res PaymentsStarsRevenueAdsAccountURL
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsTransactionsByIDRequestPredict struct {
	Peer InputPeer
	ID   []InputStarsTransaction
}

func (*PaymentsGetStarsTransactionsByIDRequestPredict) CRC() uint32 {
	return 0x27842d2e
}

func PaymentsGetStarsTransactionsByID(ctx context.Context, m Requester, i PaymentsGetStarsTransactionsByIDRequestPredict) (PaymentsStarsStatus, error) {
	var res PaymentsStarsStatus
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsGiftOptionsRequestPredict struct {
	_      struct{}  `tl:"flags,bitflag"`
	UserID InputUser `tl:",omitempty:flags:0"`
}

func (*PaymentsGetStarsGiftOptionsRequestPredict) CRC() uint32 {
	return 0xd3c96bc8
}

func PaymentsGetStarsGiftOptions(ctx context.Context, m Requester, i PaymentsGetStarsGiftOptionsRequestPredict) ([]StarsGiftOption, error) {
	var res []StarsGiftOption
	return res, request(ctx, m, &i, &res)
}

type PhoneGetCallConfigRequestPredict struct{}

func (*PhoneGetCallConfigRequestPredict) CRC() uint32 {
	return 0x55451fa9
}

func PhoneGetCallConfig(ctx context.Context, m Requester, i PhoneGetCallConfigRequestPredict) (DataJSON, error) {
	var res DataJSON
	return res, request(ctx, m, &i, &res)
}

type PhoneRequestCallRequestPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Video    bool     `tl:",omitempty:flags:0,implicit"`
	UserID   InputUser
	RandomID int32
	GAHash   []byte
	Protocol PhoneCallProtocol
}

func (*PhoneRequestCallRequestPredict) CRC() uint32 {
	return 0x42ff96ed
}

func PhoneRequestCall(ctx context.Context, m Requester, i PhoneRequestCallRequestPredict) (PhonePhoneCall, error) {
	var res PhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneAcceptCallRequestPredict struct {
	Peer     InputPhoneCall
	GB       []byte
	Protocol PhoneCallProtocol
}

func (*PhoneAcceptCallRequestPredict) CRC() uint32 {
	return 0x3bd2b4a0
}

func PhoneAcceptCall(ctx context.Context, m Requester, i PhoneAcceptCallRequestPredict) (PhonePhoneCall, error) {
	var res PhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneConfirmCallRequestPredict struct {
	Peer           InputPhoneCall
	GA             []byte
	KeyFingerprint int64
	Protocol       PhoneCallProtocol
}

func (*PhoneConfirmCallRequestPredict) CRC() uint32 {
	return 0x2efe1722
}

func PhoneConfirmCall(ctx context.Context, m Requester, i PhoneConfirmCallRequestPredict) (PhonePhoneCall, error) {
	var res PhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneReceivedCallRequestPredict struct {
	Peer InputPhoneCall
}

func (*PhoneReceivedCallRequestPredict) CRC() uint32 {
	return 0x17d54f61
}

func PhoneReceivedCall(ctx context.Context, m Requester, i PhoneReceivedCallRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneDiscardCallRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Video        bool     `tl:",omitempty:flags:0,implicit"`
	Peer         InputPhoneCall
	Duration     int32
	Reason       PhoneCallDiscardReason
	ConnectionID int64
}

func (*PhoneDiscardCallRequestPredict) CRC() uint32 {
	return 0xb2cbc1c0
}

func PhoneDiscardCall(ctx context.Context, m Requester, i PhoneDiscardCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneSetCallRatingRequestPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	UserInitiative bool     `tl:",omitempty:flags:0,implicit"`
	Peer           InputPhoneCall
	Rating         int32
	Comment        string
}

func (*PhoneSetCallRatingRequestPredict) CRC() uint32 {
	return 0x59ead627
}

func PhoneSetCallRating(ctx context.Context, m Requester, i PhoneSetCallRatingRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveCallDebugRequestPredict struct {
	Peer  InputPhoneCall
	Debug DataJSON
}

func (*PhoneSaveCallDebugRequestPredict) CRC() uint32 {
	return 0x277add7e
}

func PhoneSaveCallDebug(ctx context.Context, m Requester, i PhoneSaveCallDebugRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneSendSignalingDataRequestPredict struct {
	Peer InputPhoneCall
	Data []byte
}

func (*PhoneSendSignalingDataRequestPredict) CRC() uint32 {
	return 0xff7a9383
}

func PhoneSendSignalingData(ctx context.Context, m Requester, i PhoneSendSignalingDataRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneCreateGroupCallRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	RtmpStream   bool     `tl:",omitempty:flags:2,implicit"`
	Peer         InputPeer
	RandomID     int32
	Title        *string `tl:",omitempty:flags:0"`
	ScheduleDate *int32  `tl:",omitempty:flags:1"`
}

func (*PhoneCreateGroupCallRequestPredict) CRC() uint32 {
	return 0x48cdc6d8
}

func PhoneCreateGroupCall(ctx context.Context, m Requester, i PhoneCreateGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneJoinGroupCallRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Muted        bool     `tl:",omitempty:flags:0,implicit"`
	VideoStopped bool     `tl:",omitempty:flags:2,implicit"`
	Call         InputGroupCall
	JoinAs       InputPeer
	InviteHash   *string `tl:",omitempty:flags:1"`
	Params       DataJSON
}

func (*PhoneJoinGroupCallRequestPredict) CRC() uint32 {
	return 0xb132ff7b
}

func PhoneJoinGroupCall(ctx context.Context, m Requester, i PhoneJoinGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneLeaveGroupCallRequestPredict struct {
	Call   InputGroupCall
	Source int32
}

func (*PhoneLeaveGroupCallRequestPredict) CRC() uint32 {
	return 0x500377f9
}

func PhoneLeaveGroupCall(ctx context.Context, m Requester, i PhoneLeaveGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneInviteToGroupCallRequestPredict struct {
	Call  InputGroupCall
	Users []InputUser
}

func (*PhoneInviteToGroupCallRequestPredict) CRC() uint32 {
	return 0x7b393160
}

func PhoneInviteToGroupCall(ctx context.Context, m Requester, i PhoneInviteToGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneDiscardGroupCallRequestPredict struct {
	Call InputGroupCall
}

func (*PhoneDiscardGroupCallRequestPredict) CRC() uint32 {
	return 0x7a777135
}

func PhoneDiscardGroupCall(ctx context.Context, m Requester, i PhoneDiscardGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallSettingsRequestPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ResetInviteHash bool     `tl:",omitempty:flags:1,implicit"`
	Call            InputGroupCall
	JoinMuted       *bool `tl:",omitempty:flags:0"`
}

func (*PhoneToggleGroupCallSettingsRequestPredict) CRC() uint32 {
	return 0x74bbb43d
}

func PhoneToggleGroupCallSettings(ctx context.Context, m Requester, i PhoneToggleGroupCallSettingsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallRequestPredict struct {
	Call  InputGroupCall
	Limit int32
}

func (*PhoneGetGroupCallRequestPredict) CRC() uint32 {
	return 0x41845db
}

func PhoneGetGroupCall(ctx context.Context, m Requester, i PhoneGetGroupCallRequestPredict) (PhoneGroupCall, error) {
	var res PhoneGroupCall
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupParticipantsRequestPredict struct {
	Call    InputGroupCall
	Ids     []InputPeer
	Sources []int32
	Offset  string
	Limit   int32
}

func (*PhoneGetGroupParticipantsRequestPredict) CRC() uint32 {
	return 0xc558d8ab
}

func PhoneGetGroupParticipants(ctx context.Context, m Requester, i PhoneGetGroupParticipantsRequestPredict) (PhoneGroupParticipants, error) {
	var res PhoneGroupParticipants
	return res, request(ctx, m, &i, &res)
}

type PhoneCheckGroupCallRequestPredict struct {
	Call    InputGroupCall
	Sources []int32
}

func (*PhoneCheckGroupCallRequestPredict) CRC() uint32 {
	return 0xb59cf977
}

func PhoneCheckGroupCall(ctx context.Context, m Requester, i PhoneCheckGroupCallRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallRecordRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Start         bool     `tl:",omitempty:flags:0,implicit"`
	Video         bool     `tl:",omitempty:flags:2,implicit"`
	Call          InputGroupCall
	Title         *string `tl:",omitempty:flags:1"`
	VideoPortrait *bool   `tl:",omitempty:flags:2"`
}

func (*PhoneToggleGroupCallRecordRequestPredict) CRC() uint32 {
	return 0xf128c708
}

func PhoneToggleGroupCallRecord(ctx context.Context, m Requester, i PhoneToggleGroupCallRecordRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneEditGroupCallParticipantRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Call               InputGroupCall
	Participant        InputPeer
	Muted              *bool  `tl:",omitempty:flags:0"`
	Volume             *int32 `tl:",omitempty:flags:1"`
	RaiseHand          *bool  `tl:",omitempty:flags:2"`
	VideoStopped       *bool  `tl:",omitempty:flags:3"`
	VideoPaused        *bool  `tl:",omitempty:flags:4"`
	PresentationPaused *bool  `tl:",omitempty:flags:5"`
}

func (*PhoneEditGroupCallParticipantRequestPredict) CRC() uint32 {
	return 0xa5273abf
}

func PhoneEditGroupCallParticipant(ctx context.Context, m Requester, i PhoneEditGroupCallParticipantRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneEditGroupCallTitleRequestPredict struct {
	Call  InputGroupCall
	Title string
}

func (*PhoneEditGroupCallTitleRequestPredict) CRC() uint32 {
	return 0x1ca6ac0a
}

func PhoneEditGroupCallTitle(ctx context.Context, m Requester, i PhoneEditGroupCallTitleRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallJoinAsRequestPredict struct {
	Peer InputPeer
}

func (*PhoneGetGroupCallJoinAsRequestPredict) CRC() uint32 {
	return 0xef7c213a
}

func PhoneGetGroupCallJoinAs(ctx context.Context, m Requester, i PhoneGetGroupCallJoinAsRequestPredict) (PhoneJoinAsPeers, error) {
	var res PhoneJoinAsPeers
	return res, request(ctx, m, &i, &res)
}

type PhoneExportGroupCallInviteRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	CanSelfUnmute bool     `tl:",omitempty:flags:0,implicit"`
	Call          InputGroupCall
}

func (*PhoneExportGroupCallInviteRequestPredict) CRC() uint32 {
	return 0xe6aa647f
}

func PhoneExportGroupCallInvite(ctx context.Context, m Requester, i PhoneExportGroupCallInviteRequestPredict) (PhoneExportedGroupCallInvite, error) {
	var res PhoneExportedGroupCallInvite
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallStartSubscriptionRequestPredict struct {
	Call       InputGroupCall
	Subscribed bool
}

func (*PhoneToggleGroupCallStartSubscriptionRequestPredict) CRC() uint32 {
	return 0x219c34e6
}

func PhoneToggleGroupCallStartSubscription(ctx context.Context, m Requester, i PhoneToggleGroupCallStartSubscriptionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneStartScheduledGroupCallRequestPredict struct {
	Call InputGroupCall
}

func (*PhoneStartScheduledGroupCallRequestPredict) CRC() uint32 {
	return 0x5680e342
}

func PhoneStartScheduledGroupCall(ctx context.Context, m Requester, i PhoneStartScheduledGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveDefaultGroupCallJoinAsRequestPredict struct {
	Peer   InputPeer
	JoinAs InputPeer
}

func (*PhoneSaveDefaultGroupCallJoinAsRequestPredict) CRC() uint32 {
	return 0x575e1f8c
}

func PhoneSaveDefaultGroupCallJoinAs(ctx context.Context, m Requester, i PhoneSaveDefaultGroupCallJoinAsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneJoinGroupCallPresentationRequestPredict struct {
	Call   InputGroupCall
	Params DataJSON
}

func (*PhoneJoinGroupCallPresentationRequestPredict) CRC() uint32 {
	return 0xcbea6bc4
}

func PhoneJoinGroupCallPresentation(ctx context.Context, m Requester, i PhoneJoinGroupCallPresentationRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneLeaveGroupCallPresentationRequestPredict struct {
	Call InputGroupCall
}

func (*PhoneLeaveGroupCallPresentationRequestPredict) CRC() uint32 {
	return 0x1c50d144
}

func PhoneLeaveGroupCallPresentation(ctx context.Context, m Requester, i PhoneLeaveGroupCallPresentationRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallStreamChannelsRequestPredict struct {
	Call InputGroupCall
}

func (*PhoneGetGroupCallStreamChannelsRequestPredict) CRC() uint32 {
	return 0x1ab21940
}

func PhoneGetGroupCallStreamChannels(ctx context.Context, m Requester, i PhoneGetGroupCallStreamChannelsRequestPredict) (PhoneGroupCallStreamChannels, error) {
	var res PhoneGroupCallStreamChannels
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallStreamRtmpURLRequestPredict struct {
	Peer   InputPeer
	Revoke bool
}

func (*PhoneGetGroupCallStreamRtmpURLRequestPredict) CRC() uint32 {
	return 0xdeb3abbf
}

func PhoneGetGroupCallStreamRtmpURL(ctx context.Context, m Requester, i PhoneGetGroupCallStreamRtmpURLRequestPredict) (PhoneGroupCallStreamRtmpURL, error) {
	var res PhoneGroupCallStreamRtmpURL
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveCallLogRequestPredict struct {
	Peer InputPhoneCall
	File InputFile
}

func (*PhoneSaveCallLogRequestPredict) CRC() uint32 {
	return 0x41248786
}

func PhoneSaveCallLog(ctx context.Context, m Requester, i PhoneSaveCallLogRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhotosUpdateProfilePhotoRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Fallback bool      `tl:",omitempty:flags:0,implicit"`
	Bot      InputUser `tl:",omitempty:flags:1"`
	ID       InputPhoto
}

func (*PhotosUpdateProfilePhotoRequestPredict) CRC() uint32 {
	return 0x9e82039
}

func PhotosUpdateProfilePhoto(ctx context.Context, m Requester, i PhotosUpdateProfilePhotoRequestPredict) (PhotosPhoto, error) {
	var res PhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PhotosUploadProfilePhotoRequestPredict struct {
	_                struct{}  `tl:"flags,bitflag"`
	Fallback         bool      `tl:",omitempty:flags:3,implicit"`
	Bot              InputUser `tl:",omitempty:flags:5"`
	File             InputFile `tl:",omitempty:flags:0"`
	Video            InputFile `tl:",omitempty:flags:1"`
	VideoStartTs     *float64  `tl:",omitempty:flags:2"`
	VideoEmojiMarkup VideoSize `tl:",omitempty:flags:4"`
}

func (*PhotosUploadProfilePhotoRequestPredict) CRC() uint32 {
	return 0x388a3b5
}

func PhotosUploadProfilePhoto(ctx context.Context, m Requester, i PhotosUploadProfilePhotoRequestPredict) (PhotosPhoto, error) {
	var res PhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PhotosDeletePhotosRequestPredict struct {
	ID []InputPhoto
}

func (*PhotosDeletePhotosRequestPredict) CRC() uint32 {
	return 0x87cf7f2f
}

func PhotosDeletePhotos(ctx context.Context, m Requester, i PhotosDeletePhotosRequestPredict) ([]int64, error) {
	var res []int64
	return res, request(ctx, m, &i, &res)
}

type PhotosGetUserPhotosRequestPredict struct {
	UserID InputUser
	Offset int32
	MaxID  int64
	Limit  int32
}

func (*PhotosGetUserPhotosRequestPredict) CRC() uint32 {
	return 0x91cd32a8
}

func PhotosGetUserPhotos(ctx context.Context, m Requester, i PhotosGetUserPhotosRequestPredict) (PhotosPhotos, error) {
	var res PhotosPhotos
	return res, request(ctx, m, &i, &res)
}

type PhotosUploadContactProfilePhotoRequestPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	Suggest          bool     `tl:",omitempty:flags:3,implicit"`
	Save             bool     `tl:",omitempty:flags:4,implicit"`
	UserID           InputUser
	File             InputFile `tl:",omitempty:flags:0"`
	Video            InputFile `tl:",omitempty:flags:1"`
	VideoStartTs     *float64  `tl:",omitempty:flags:2"`
	VideoEmojiMarkup VideoSize `tl:",omitempty:flags:5"`
}

func (*PhotosUploadContactProfilePhotoRequestPredict) CRC() uint32 {
	return 0xe14c4a71
}

func PhotosUploadContactProfilePhoto(ctx context.Context, m Requester, i PhotosUploadContactProfilePhotoRequestPredict) (PhotosPhoto, error) {
	var res PhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PremiumGetBoostsListRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Gifts  bool     `tl:",omitempty:flags:0,implicit"`
	Peer   InputPeer
	Offset string
	Limit  int32
}

func (*PremiumGetBoostsListRequestPredict) CRC() uint32 {
	return 0x60f67660
}

func PremiumGetBoostsList(ctx context.Context, m Requester, i PremiumGetBoostsListRequestPredict) (PremiumBoostsList, error) {
	var res PremiumBoostsList
	return res, request(ctx, m, &i, &res)
}

type PremiumGetMyBoostsRequestPredict struct{}

func (*PremiumGetMyBoostsRequestPredict) CRC() uint32 {
	return 0xbe77b4a
}

func PremiumGetMyBoosts(ctx context.Context, m Requester, i PremiumGetMyBoostsRequestPredict) (PremiumMyBoosts, error) {
	var res PremiumMyBoosts
	return res, request(ctx, m, &i, &res)
}

type PremiumApplyBoostRequestPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Slots []int32  `tl:",omitempty:flags:0"`
	Peer  InputPeer
}

func (*PremiumApplyBoostRequestPredict) CRC() uint32 {
	return 0x6b7da746
}

func PremiumApplyBoost(ctx context.Context, m Requester, i PremiumApplyBoostRequestPredict) (PremiumMyBoosts, error) {
	var res PremiumMyBoosts
	return res, request(ctx, m, &i, &res)
}

type PremiumGetBoostsStatusRequestPredict struct {
	Peer InputPeer
}

func (*PremiumGetBoostsStatusRequestPredict) CRC() uint32 {
	return 0x42f1f61
}

func PremiumGetBoostsStatus(ctx context.Context, m Requester, i PremiumGetBoostsStatusRequestPredict) (PremiumBoostsStatus, error) {
	var res PremiumBoostsStatus
	return res, request(ctx, m, &i, &res)
}

type PremiumGetUserBoostsRequestPredict struct {
	Peer   InputPeer
	UserID InputUser
}

func (*PremiumGetUserBoostsRequestPredict) CRC() uint32 {
	return 0x39854d1f
}

func PremiumGetUserBoosts(ctx context.Context, m Requester, i PremiumGetUserBoostsRequestPredict) (PremiumBoostsList, error) {
	var res PremiumBoostsList
	return res, request(ctx, m, &i, &res)
}

type SmsjobsIsEligibleToJoinRequestPredict struct{}

func (*SmsjobsIsEligibleToJoinRequestPredict) CRC() uint32 {
	return 0xedc39d0
}

func SmsjobsIsEligibleToJoin(ctx context.Context, m Requester, i SmsjobsIsEligibleToJoinRequestPredict) (SmsjobsEligibilityToJoin, error) {
	var res SmsjobsEligibilityToJoin
	return res, request(ctx, m, &i, &res)
}

type SmsjobsJoinRequestPredict struct{}

func (*SmsjobsJoinRequestPredict) CRC() uint32 {
	return 0xa74ece2d
}

func SmsjobsJoin(ctx context.Context, m Requester, i SmsjobsJoinRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type SmsjobsLeaveRequestPredict struct{}

func (*SmsjobsLeaveRequestPredict) CRC() uint32 {
	return 0x9898ad73
}

func SmsjobsLeave(ctx context.Context, m Requester, i SmsjobsLeaveRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type SmsjobsUpdateSettingsRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	AllowInternational bool     `tl:",omitempty:flags:0,implicit"`
}

func (*SmsjobsUpdateSettingsRequestPredict) CRC() uint32 {
	return 0x93fa0bf
}

func SmsjobsUpdateSettings(ctx context.Context, m Requester, i SmsjobsUpdateSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type SmsjobsGetStatusRequestPredict struct{}

func (*SmsjobsGetStatusRequestPredict) CRC() uint32 {
	return 0x10a698e8
}

func SmsjobsGetStatus(ctx context.Context, m Requester, i SmsjobsGetStatusRequestPredict) (SmsjobsStatus, error) {
	var res SmsjobsStatus
	return res, request(ctx, m, &i, &res)
}

type SmsjobsGetSmsJobRequestPredict struct {
	JobID string
}

func (*SmsjobsGetSmsJobRequestPredict) CRC() uint32 {
	return 0x778d902f
}

func SmsjobsGetSmsJob(ctx context.Context, m Requester, i SmsjobsGetSmsJobRequestPredict) (SmsJob, error) {
	var res SmsJob
	return res, request(ctx, m, &i, &res)
}

type SmsjobsFinishJobRequestPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	JobID string
	Error *string `tl:",omitempty:flags:0"`
}

func (*SmsjobsFinishJobRequestPredict) CRC() uint32 {
	return 0x4f1ebf24
}

func SmsjobsFinishJob(ctx context.Context, m Requester, i SmsjobsFinishJobRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastStatsRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Dark    bool     `tl:",omitempty:flags:0,implicit"`
	Channel InputChannel
}

func (*StatsGetBroadcastStatsRequestPredict) CRC() uint32 {
	return 0xab42441a
}

func StatsGetBroadcastStats(ctx context.Context, m Requester, i StatsGetBroadcastStatsRequestPredict) (StatsBroadcastStats, error) {
	var res StatsBroadcastStats
	return res, request(ctx, m, &i, &res)
}

type StatsLoadAsyncGraphRequestPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Token string
	X     *int64 `tl:",omitempty:flags:0"`
}

func (*StatsLoadAsyncGraphRequestPredict) CRC() uint32 {
	return 0x621d5fa0
}

func StatsLoadAsyncGraph(ctx context.Context, m Requester, i StatsLoadAsyncGraphRequestPredict) (StatsGraph, error) {
	var res StatsGraph
	return res, request(ctx, m, &i, &res)
}

type StatsGetMegagroupStatsRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Dark    bool     `tl:",omitempty:flags:0,implicit"`
	Channel InputChannel
}

func (*StatsGetMegagroupStatsRequestPredict) CRC() uint32 {
	return 0xdcdf8607
}

func StatsGetMegagroupStats(ctx context.Context, m Requester, i StatsGetMegagroupStatsRequestPredict) (StatsMegagroupStats, error) {
	var res StatsMegagroupStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetMessagePublicForwardsRequestPredict struct {
	Channel InputChannel
	MsgID   int32
	Offset  string
	Limit   int32
}

func (*StatsGetMessagePublicForwardsRequestPredict) CRC() uint32 {
	return 0x5f150144
}

func StatsGetMessagePublicForwards(ctx context.Context, m Requester, i StatsGetMessagePublicForwardsRequestPredict) (StatsPublicForwards, error) {
	var res StatsPublicForwards
	return res, request(ctx, m, &i, &res)
}

type StatsGetMessageStatsRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Dark    bool     `tl:",omitempty:flags:0,implicit"`
	Channel InputChannel
	MsgID   int32
}

func (*StatsGetMessageStatsRequestPredict) CRC() uint32 {
	return 0xb6e0a3f5
}

func StatsGetMessageStats(ctx context.Context, m Requester, i StatsGetMessageStatsRequestPredict) (StatsMessageStats, error) {
	var res StatsMessageStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetStoryStatsRequestPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Dark bool     `tl:",omitempty:flags:0,implicit"`
	Peer InputPeer
	ID   int32
}

func (*StatsGetStoryStatsRequestPredict) CRC() uint32 {
	return 0x374fef40
}

func StatsGetStoryStats(ctx context.Context, m Requester, i StatsGetStoryStatsRequestPredict) (StatsStoryStats, error) {
	var res StatsStoryStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetStoryPublicForwardsRequestPredict struct {
	Peer   InputPeer
	ID     int32
	Offset string
	Limit  int32
}

func (*StatsGetStoryPublicForwardsRequestPredict) CRC() uint32 {
	return 0xa6437ef6
}

func StatsGetStoryPublicForwards(ctx context.Context, m Requester, i StatsGetStoryPublicForwardsRequestPredict) (StatsPublicForwards, error) {
	var res StatsPublicForwards
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastRevenueStatsRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Dark    bool     `tl:",omitempty:flags:0,implicit"`
	Channel InputChannel
}

func (*StatsGetBroadcastRevenueStatsRequestPredict) CRC() uint32 {
	return 0x75dfb671
}

func StatsGetBroadcastRevenueStats(ctx context.Context, m Requester, i StatsGetBroadcastRevenueStatsRequestPredict) (StatsBroadcastRevenueStats, error) {
	var res StatsBroadcastRevenueStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastRevenueWithdrawalURLRequestPredict struct {
	Channel  InputChannel
	Password InputCheckPasswordSRP
}

func (*StatsGetBroadcastRevenueWithdrawalURLRequestPredict) CRC() uint32 {
	return 0x2a65ef73
}

func StatsGetBroadcastRevenueWithdrawalURL(ctx context.Context, m Requester, i StatsGetBroadcastRevenueWithdrawalURLRequestPredict) (StatsBroadcastRevenueWithdrawalURL, error) {
	var res StatsBroadcastRevenueWithdrawalURL
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastRevenueTransactionsRequestPredict struct {
	Channel InputChannel
	Offset  int32
	Limit   int32
}

func (*StatsGetBroadcastRevenueTransactionsRequestPredict) CRC() uint32 {
	return 0x69280f
}

func StatsGetBroadcastRevenueTransactions(ctx context.Context, m Requester, i StatsGetBroadcastRevenueTransactionsRequestPredict) (StatsBroadcastRevenueTransactions, error) {
	var res StatsBroadcastRevenueTransactions
	return res, request(ctx, m, &i, &res)
}

type StickersCreateStickerSetRequestPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Masks     bool     `tl:",omitempty:flags:0,implicit"`
	Emojis    bool     `tl:",omitempty:flags:5,implicit"`
	TextColor bool     `tl:",omitempty:flags:6,implicit"`
	UserID    InputUser
	Title     string
	ShortName string
	Thumb     InputDocument `tl:",omitempty:flags:2"`
	Stickers  []InputStickerSetItem
	Software  *string `tl:",omitempty:flags:3"`
}

func (*StickersCreateStickerSetRequestPredict) CRC() uint32 {
	return 0x9021ab67
}

func StickersCreateStickerSet(ctx context.Context, m Requester, i StickersCreateStickerSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersRemoveStickerFromSetRequestPredict struct {
	Sticker InputDocument
}

func (*StickersRemoveStickerFromSetRequestPredict) CRC() uint32 {
	return 0xf7760f51
}

func StickersRemoveStickerFromSet(ctx context.Context, m Requester, i StickersRemoveStickerFromSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersChangeStickerPositionRequestPredict struct {
	Sticker  InputDocument
	Position int32
}

func (*StickersChangeStickerPositionRequestPredict) CRC() uint32 {
	return 0xffb6d4ca
}

func StickersChangeStickerPosition(ctx context.Context, m Requester, i StickersChangeStickerPositionRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersAddStickerToSetRequestPredict struct {
	Stickerset InputStickerSet
	Sticker    InputStickerSetItem
}

func (*StickersAddStickerToSetRequestPredict) CRC() uint32 {
	return 0x8653febe
}

func StickersAddStickerToSet(ctx context.Context, m Requester, i StickersAddStickerToSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersSetStickerSetThumbRequestPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Stickerset      InputStickerSet
	Thumb           InputDocument `tl:",omitempty:flags:0"`
	ThumbDocumentID *int64        `tl:",omitempty:flags:1"`
}

func (*StickersSetStickerSetThumbRequestPredict) CRC() uint32 {
	return 0xa76a5392
}

func StickersSetStickerSetThumb(ctx context.Context, m Requester, i StickersSetStickerSetThumbRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersCheckShortNameRequestPredict struct {
	ShortName string
}

func (*StickersCheckShortNameRequestPredict) CRC() uint32 {
	return 0x284b3639
}

func StickersCheckShortName(ctx context.Context, m Requester, i StickersCheckShortNameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StickersSuggestShortNameRequestPredict struct {
	Title string
}

func (*StickersSuggestShortNameRequestPredict) CRC() uint32 {
	return 0x4dafc503
}

func StickersSuggestShortName(ctx context.Context, m Requester, i StickersSuggestShortNameRequestPredict) (StickersSuggestedShortName, error) {
	var res StickersSuggestedShortName
	return res, request(ctx, m, &i, &res)
}

type StickersChangeStickerRequestPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Sticker    InputDocument
	Emoji      *string    `tl:",omitempty:flags:0"`
	MaskCoords MaskCoords `tl:",omitempty:flags:1"`
	Keywords   *string    `tl:",omitempty:flags:2"`
}

func (*StickersChangeStickerRequestPredict) CRC() uint32 {
	return 0xf5537ebc
}

func StickersChangeSticker(ctx context.Context, m Requester, i StickersChangeStickerRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersRenameStickerSetRequestPredict struct {
	Stickerset InputStickerSet
	Title      string
}

func (*StickersRenameStickerSetRequestPredict) CRC() uint32 {
	return 0x124b1c00
}

func StickersRenameStickerSet(ctx context.Context, m Requester, i StickersRenameStickerSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersDeleteStickerSetRequestPredict struct {
	Stickerset InputStickerSet
}

func (*StickersDeleteStickerSetRequestPredict) CRC() uint32 {
	return 0x87704394
}

func StickersDeleteStickerSet(ctx context.Context, m Requester, i StickersDeleteStickerSetRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StickersReplaceStickerRequestPredict struct {
	Sticker    InputDocument
	NewSticker InputStickerSetItem
}

func (*StickersReplaceStickerRequestPredict) CRC() uint32 {
	return 0x4696459a
}

func StickersReplaceSticker(ctx context.Context, m Requester, i StickersReplaceStickerRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StoriesCanSendStoryRequestPredict struct {
	Peer InputPeer
}

func (*StoriesCanSendStoryRequestPredict) CRC() uint32 {
	return 0xc7dfdfdd
}

func StoriesCanSendStory(ctx context.Context, m Requester, i StoriesCanSendStoryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesSendStoryRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Pinned       bool     `tl:",omitempty:flags:2,implicit"`
	Noforwards   bool     `tl:",omitempty:flags:4,implicit"`
	FwdModified  bool     `tl:",omitempty:flags:7,implicit"`
	Peer         InputPeer
	Media        InputMedia
	MediaAreas   []MediaArea     `tl:",omitempty:flags:5"`
	Caption      *string         `tl:",omitempty:flags:0"`
	Entities     []MessageEntity `tl:",omitempty:flags:1"`
	PrivacyRules []InputPrivacyRule
	RandomID     int64
	Period       *int32    `tl:",omitempty:flags:3"`
	FwdFromID    InputPeer `tl:",omitempty:flags:6"`
	FwdFromStory *int32    `tl:",omitempty:flags:6"`
}

func (*StoriesSendStoryRequestPredict) CRC() uint32 {
	return 0xe4e6694b
}

func StoriesSendStory(ctx context.Context, m Requester, i StoriesSendStoryRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesEditStoryRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Peer         InputPeer
	ID           int32
	Media        InputMedia         `tl:",omitempty:flags:0"`
	MediaAreas   []MediaArea        `tl:",omitempty:flags:3"`
	Caption      *string            `tl:",omitempty:flags:1"`
	Entities     []MessageEntity    `tl:",omitempty:flags:1"`
	PrivacyRules []InputPrivacyRule `tl:",omitempty:flags:2"`
}

func (*StoriesEditStoryRequestPredict) CRC() uint32 {
	return 0xb583ba46
}

func StoriesEditStory(ctx context.Context, m Requester, i StoriesEditStoryRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesDeleteStoriesRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*StoriesDeleteStoriesRequestPredict) CRC() uint32 {
	return 0xae59db5f
}

func StoriesDeleteStories(ctx context.Context, m Requester, i StoriesDeleteStoriesRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesTogglePinnedRequestPredict struct {
	Peer   InputPeer
	ID     []int32
	Pinned bool
}

func (*StoriesTogglePinnedRequestPredict) CRC() uint32 {
	return 0x9a75a1ef
}

func StoriesTogglePinned(ctx context.Context, m Requester, i StoriesTogglePinnedRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesGetAllStoriesRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Next   bool     `tl:",omitempty:flags:1,implicit"`
	Hidden bool     `tl:",omitempty:flags:2,implicit"`
	State  *string  `tl:",omitempty:flags:0"`
}

func (*StoriesGetAllStoriesRequestPredict) CRC() uint32 {
	return 0xeeb0d625
}

func StoriesGetAllStories(ctx context.Context, m Requester, i StoriesGetAllStoriesRequestPredict) (StoriesAllStories, error) {
	var res StoriesAllStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPinnedStoriesRequestPredict struct {
	Peer     InputPeer
	OffsetID int32
	Limit    int32
}

func (*StoriesGetPinnedStoriesRequestPredict) CRC() uint32 {
	return 0x5821a5dc
}

func StoriesGetPinnedStories(ctx context.Context, m Requester, i StoriesGetPinnedStoriesRequestPredict) (StoriesStories, error) {
	var res StoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesArchiveRequestPredict struct {
	Peer     InputPeer
	OffsetID int32
	Limit    int32
}

func (*StoriesGetStoriesArchiveRequestPredict) CRC() uint32 {
	return 0xb4352016
}

func StoriesGetStoriesArchive(ctx context.Context, m Requester, i StoriesGetStoriesArchiveRequestPredict) (StoriesStories, error) {
	var res StoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesByIDRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*StoriesGetStoriesByIDRequestPredict) CRC() uint32 {
	return 0x5774ca74
}

func StoriesGetStoriesByID(ctx context.Context, m Requester, i StoriesGetStoriesByIDRequestPredict) (StoriesStories, error) {
	var res StoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesToggleAllStoriesHiddenRequestPredict struct {
	Hidden bool
}

func (*StoriesToggleAllStoriesHiddenRequestPredict) CRC() uint32 {
	return 0x7c2557c4
}

func StoriesToggleAllStoriesHidden(ctx context.Context, m Requester, i StoriesToggleAllStoriesHiddenRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesReadStoriesRequestPredict struct {
	Peer  InputPeer
	MaxID int32
}

func (*StoriesReadStoriesRequestPredict) CRC() uint32 {
	return 0xa556dac8
}

func StoriesReadStories(ctx context.Context, m Requester, i StoriesReadStoriesRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesIncrementStoryViewsRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*StoriesIncrementStoryViewsRequestPredict) CRC() uint32 {
	return 0xb2028afb
}

func StoriesIncrementStoryViews(ctx context.Context, m Requester, i StoriesIncrementStoryViewsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoryViewsListRequestPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	JustContacts   bool     `tl:",omitempty:flags:0,implicit"`
	ReactionsFirst bool     `tl:",omitempty:flags:2,implicit"`
	ForwardsFirst  bool     `tl:",omitempty:flags:3,implicit"`
	Peer           InputPeer
	Q              *string `tl:",omitempty:flags:1"`
	ID             int32
	Offset         string
	Limit          int32
}

func (*StoriesGetStoryViewsListRequestPredict) CRC() uint32 {
	return 0x7ed23c57
}

func StoriesGetStoryViewsList(ctx context.Context, m Requester, i StoriesGetStoryViewsListRequestPredict) (StoriesStoryViewsList, error) {
	var res StoriesStoryViewsList
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesViewsRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*StoriesGetStoriesViewsRequestPredict) CRC() uint32 {
	return 0x28e16cc8
}

func StoriesGetStoriesViews(ctx context.Context, m Requester, i StoriesGetStoriesViewsRequestPredict) (StoriesStoryViews, error) {
	var res StoriesStoryViews
	return res, request(ctx, m, &i, &res)
}

type StoriesExportStoryLinkRequestPredict struct {
	Peer InputPeer
	ID   int32
}

func (*StoriesExportStoryLinkRequestPredict) CRC() uint32 {
	return 0x7b8def20
}

func StoriesExportStoryLink(ctx context.Context, m Requester, i StoriesExportStoryLinkRequestPredict) (ExportedStoryLink, error) {
	var res ExportedStoryLink
	return res, request(ctx, m, &i, &res)
}

type StoriesReportRequestPredict struct {
	Peer    InputPeer
	ID      []int32
	Reason  ReportReason
	Message string
}

func (*StoriesReportRequestPredict) CRC() uint32 {
	return 0x1923fa8c
}

func StoriesReport(ctx context.Context, m Requester, i StoriesReportRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesActivateStealthModeRequestPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Past   bool     `tl:",omitempty:flags:0,implicit"`
	Future bool     `tl:",omitempty:flags:1,implicit"`
}

func (*StoriesActivateStealthModeRequestPredict) CRC() uint32 {
	return 0x57bbd166
}

func StoriesActivateStealthMode(ctx context.Context, m Requester, i StoriesActivateStealthModeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesSendReactionRequestPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	AddToRecent bool     `tl:",omitempty:flags:0,implicit"`
	Peer        InputPeer
	StoryID     int32
	Reaction    Reaction
}

func (*StoriesSendReactionRequestPredict) CRC() uint32 {
	return 0x7fd736b2
}

func StoriesSendReaction(ctx context.Context, m Requester, i StoriesSendReactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPeerStoriesRequestPredict struct {
	Peer InputPeer
}

func (*StoriesGetPeerStoriesRequestPredict) CRC() uint32 {
	return 0x2c4ada50
}

func StoriesGetPeerStories(ctx context.Context, m Requester, i StoriesGetPeerStoriesRequestPredict) (StoriesPeerStories, error) {
	var res StoriesPeerStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetAllReadPeerStoriesRequestPredict struct{}

func (*StoriesGetAllReadPeerStoriesRequestPredict) CRC() uint32 {
	return 0x9b5ae7f9
}

func StoriesGetAllReadPeerStories(ctx context.Context, m Requester, i StoriesGetAllReadPeerStoriesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPeerMaxIDsRequestPredict struct {
	ID []InputPeer
}

func (*StoriesGetPeerMaxIDsRequestPredict) CRC() uint32 {
	return 0x535983c3
}

func StoriesGetPeerMaxIDs(ctx context.Context, m Requester, i StoriesGetPeerMaxIDsRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesGetChatsToSendRequestPredict struct{}

func (*StoriesGetChatsToSendRequestPredict) CRC() uint32 {
	return 0xa56a8b60
}

func StoriesGetChatsToSend(ctx context.Context, m Requester, i StoriesGetChatsToSendRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type StoriesTogglePeerStoriesHiddenRequestPredict struct {
	Peer   InputPeer
	Hidden bool
}

func (*StoriesTogglePeerStoriesHiddenRequestPredict) CRC() uint32 {
	return 0xbd0415c4
}

func StoriesTogglePeerStoriesHidden(ctx context.Context, m Requester, i StoriesTogglePeerStoriesHiddenRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoryReactionsListRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ForwardsFirst bool     `tl:",omitempty:flags:2,implicit"`
	Peer          InputPeer
	ID            int32
	Reaction      Reaction `tl:",omitempty:flags:0"`
	Offset        *string  `tl:",omitempty:flags:1"`
	Limit         int32
}

func (*StoriesGetStoryReactionsListRequestPredict) CRC() uint32 {
	return 0xb9b2881f
}

func StoriesGetStoryReactionsList(ctx context.Context, m Requester, i StoriesGetStoryReactionsListRequestPredict) (StoriesStoryReactionsList, error) {
	var res StoriesStoryReactionsList
	return res, request(ctx, m, &i, &res)
}

type StoriesTogglePinnedToTopRequestPredict struct {
	Peer InputPeer
	ID   []int32
}

func (*StoriesTogglePinnedToTopRequestPredict) CRC() uint32 {
	return 0xb297e9b
}

func StoriesTogglePinnedToTop(ctx context.Context, m Requester, i StoriesTogglePinnedToTopRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesSearchPostsRequestPredict struct {
	_       struct{}  `tl:"flags,bitflag"`
	Hashtag *string   `tl:",omitempty:flags:0"`
	Area    MediaArea `tl:",omitempty:flags:1"`
	Offset  string
	Limit   int32
}

func (*StoriesSearchPostsRequestPredict) CRC() uint32 {
	return 0x6cea116a
}

func StoriesSearchPosts(ctx context.Context, m Requester, i StoriesSearchPostsRequestPredict) (StoriesFoundStories, error) {
	var res StoriesFoundStories
	return res, request(ctx, m, &i, &res)
}

type UpdatesGetStateRequestPredict struct{}

func (*UpdatesGetStateRequestPredict) CRC() uint32 {
	return 0xedd4882a
}

func UpdatesGetState(ctx context.Context, m Requester, i UpdatesGetStateRequestPredict) (UpdatesState, error) {
	var res UpdatesState
	return res, request(ctx, m, &i, &res)
}

type UpdatesGetDifferenceRequestPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Pts           int32
	PtsLimit      *int32 `tl:",omitempty:flags:1"`
	PtsTotalLimit *int32 `tl:",omitempty:flags:0"`
	Date          int32
	Qts           int32
	QtsLimit      *int32 `tl:",omitempty:flags:2"`
}

func (*UpdatesGetDifferenceRequestPredict) CRC() uint32 {
	return 0x19c2f763
}

func UpdatesGetDifference(ctx context.Context, m Requester, i UpdatesGetDifferenceRequestPredict) (UpdatesDifference, error) {
	var res UpdatesDifference
	return res, request(ctx, m, &i, &res)
}

type UpdatesGetChannelDifferenceRequestPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Force   bool     `tl:",omitempty:flags:0,implicit"`
	Channel InputChannel
	Filter  ChannelMessagesFilter
	Pts     int32
	Limit   int32
}

func (*UpdatesGetChannelDifferenceRequestPredict) CRC() uint32 {
	return 0x3173d78
}

func UpdatesGetChannelDifference(ctx context.Context, m Requester, i UpdatesGetChannelDifferenceRequestPredict) (UpdatesChannelDifference, error) {
	var res UpdatesChannelDifference
	return res, request(ctx, m, &i, &res)
}

type UploadSaveFilePartRequestPredict struct {
	FileID   int64
	FilePart int32
	Bytes    []byte
}

func (*UploadSaveFilePartRequestPredict) CRC() uint32 {
	return 0xb304a621
}

func UploadSaveFilePart(ctx context.Context, m Requester, i UploadSaveFilePartRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UploadGetFileRequestPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Precise      bool     `tl:",omitempty:flags:0,implicit"`
	CdnSupported bool     `tl:",omitempty:flags:1,implicit"`
	Location     InputFileLocation
	Offset       int64
	Limit        int32
}

func (*UploadGetFileRequestPredict) CRC() uint32 {
	return 0xbe5335be
}

func UploadGetFile(ctx context.Context, m Requester, i UploadGetFileRequestPredict) (UploadFile, error) {
	var res UploadFile
	return res, request(ctx, m, &i, &res)
}

type UploadSaveBigFilePartRequestPredict struct {
	FileID         int64
	FilePart       int32
	FileTotalParts int32
	Bytes          []byte
}

func (*UploadSaveBigFilePartRequestPredict) CRC() uint32 {
	return 0xde7b673d
}

func UploadSaveBigFilePart(ctx context.Context, m Requester, i UploadSaveBigFilePartRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UploadGetWebFileRequestPredict struct {
	Location InputWebFileLocation
	Offset   int32
	Limit    int32
}

func (*UploadGetWebFileRequestPredict) CRC() uint32 {
	return 0x24e6818d
}

func UploadGetWebFile(ctx context.Context, m Requester, i UploadGetWebFileRequestPredict) (UploadWebFile, error) {
	var res UploadWebFile
	return res, request(ctx, m, &i, &res)
}

type UploadGetCdnFileRequestPredict struct {
	FileToken []byte
	Offset    int64
	Limit     int32
}

func (*UploadGetCdnFileRequestPredict) CRC() uint32 {
	return 0x395f69da
}

func UploadGetCdnFile(ctx context.Context, m Requester, i UploadGetCdnFileRequestPredict) (UploadCdnFile, error) {
	var res UploadCdnFile
	return res, request(ctx, m, &i, &res)
}

type UploadReuploadCdnFileRequestPredict struct {
	FileToken    []byte
	RequestToken []byte
}

func (*UploadReuploadCdnFileRequestPredict) CRC() uint32 {
	return 0x9b2754a8
}

func UploadReuploadCdnFile(ctx context.Context, m Requester, i UploadReuploadCdnFileRequestPredict) ([]FileHash, error) {
	var res []FileHash
	return res, request(ctx, m, &i, &res)
}

type UploadGetCdnFileHashesRequestPredict struct {
	FileToken []byte
	Offset    int64
}

func (*UploadGetCdnFileHashesRequestPredict) CRC() uint32 {
	return 0x91dc3f31
}

func UploadGetCdnFileHashes(ctx context.Context, m Requester, i UploadGetCdnFileHashesRequestPredict) ([]FileHash, error) {
	var res []FileHash
	return res, request(ctx, m, &i, &res)
}

type UploadGetFileHashesRequestPredict struct {
	Location InputFileLocation
	Offset   int64
}

func (*UploadGetFileHashesRequestPredict) CRC() uint32 {
	return 0x9156982a
}

func UploadGetFileHashes(ctx context.Context, m Requester, i UploadGetFileHashesRequestPredict) ([]FileHash, error) {
	var res []FileHash
	return res, request(ctx, m, &i, &res)
}

type UsersGetUsersRequestPredict struct {
	ID []InputUser
}

func (*UsersGetUsersRequestPredict) CRC() uint32 {
	return 0xd91a548
}

func UsersGetUsers(ctx context.Context, m Requester, i UsersGetUsersRequestPredict) ([]User, error) {
	var res []User
	return res, request(ctx, m, &i, &res)
}

type UsersGetFullUserRequestPredict struct {
	ID InputUser
}

func (*UsersGetFullUserRequestPredict) CRC() uint32 {
	return 0xb60f5918
}

func UsersGetFullUser(ctx context.Context, m Requester, i UsersGetFullUserRequestPredict) (UsersUserFull, error) {
	var res UsersUserFull
	return res, request(ctx, m, &i, &res)
}

type UsersSetSecureValueErrorsRequestPredict struct {
	ID     InputUser
	Errors []SecureValueError
}

func (*UsersSetSecureValueErrorsRequestPredict) CRC() uint32 {
	return 0x90c894b5
}

func UsersSetSecureValueErrors(ctx context.Context, m Requester, i UsersSetSecureValueErrorsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UsersGetIsPremiumRequiredToContactRequestPredict struct {
	ID []InputUser
}

func (*UsersGetIsPremiumRequiredToContactRequestPredict) CRC() uint32 {
	return 0xa622aa10
}

func UsersGetIsPremiumRequiredToContact(ctx context.Context, m Requester, i UsersGetIsPremiumRequiredToContactRequestPredict) ([]bool, error) {
	var res []bool
	return res, request(ctx, m, &i, &res)
}
