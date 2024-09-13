package telegram

import (
	"context"
	"fmt"
	tl "github.com/xelaj/tl"
)

type AccountDaysTTL interface {
	tl.TLObject
	_AccountDaysTTL()
}

var (
	_ AccountDaysTTL = (*AccountDaysTTLPredict)(nil)
)

type AccountDaysTTLPredict struct {
	Days int32 `tl:"days"`
}

func (*AccountDaysTTLPredict) CRC() uint32 {
	return 0xb8d0afdf
}
func (*AccountDaysTTLPredict) _AccountDaysTTL() {}

type AttachMenuBot interface {
	tl.TLObject
	_AttachMenuBot()
}

var (
	_ AttachMenuBot = (*AttachMenuBotPredict)(nil)
)

type AttachMenuBotPredict struct {
	_                        struct{}             `tl:"flags,bitflag"`
	Inactive                 bool                 `tl:"inactive,omitempty:flags:0,implicit"`
	HasSettings              bool                 `tl:"has_settings,omitempty:flags:1,implicit"`
	RequestWriteAccess       bool                 `tl:"request_write_access,omitempty:flags:2,implicit"`
	ShowInAttachMenu         bool                 `tl:"show_in_attach_menu,omitempty:flags:3,implicit"`
	ShowInSideMenu           bool                 `tl:"show_in_side_menu,omitempty:flags:4,implicit"`
	SideMenuDisclaimerNeeded bool                 `tl:"side_menu_disclaimer_needed,omitempty:flags:5,implicit"`
	BotID                    int64                `tl:"bot_id"`
	ShortName                string               `tl:"short_name"`
	PeerTypes                []AttachMenuPeerType `tl:"peer_types,omitempty:flags:3"`
	Icons                    []AttachMenuBotIcon  `tl:"icons"`
}

func (*AttachMenuBotPredict) CRC() uint32 {
	return 0xd90d8dfe
}
func (*AttachMenuBotPredict) _AttachMenuBot() {}

type AttachMenuBotIcon interface {
	tl.TLObject
	_AttachMenuBotIcon()
}

var (
	_ AttachMenuBotIcon = (*AttachMenuBotIconPredict)(nil)
)

type AttachMenuBotIconPredict struct {
	_      struct{}                 `tl:"flags,bitflag"`
	Name   string                   `tl:"name"`
	Icon   Document                 `tl:"icon"`
	Colors []AttachMenuBotIconColor `tl:"colors,omitempty:flags:0"`
}

func (*AttachMenuBotIconPredict) CRC() uint32 {
	return 0xb2a7386b
}
func (*AttachMenuBotIconPredict) _AttachMenuBotIcon() {}

type AttachMenuBotIconColor interface {
	tl.TLObject
	_AttachMenuBotIconColor()
}

var (
	_ AttachMenuBotIconColor = (*AttachMenuBotIconColorPredict)(nil)
)

type AttachMenuBotIconColorPredict struct {
	Name  string `tl:"name"`
	Color int32  `tl:"color"`
}

func (*AttachMenuBotIconColorPredict) CRC() uint32 {
	return 0x4576f3f0
}
func (*AttachMenuBotIconColorPredict) _AttachMenuBotIconColor() {}

type AttachMenuBots interface {
	tl.TLObject
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
	Hash  int64           `tl:"hash"`
	Bots  []AttachMenuBot `tl:"bots"`
	Users []User          `tl:"users"`
}

func (*AttachMenuBotsPredict) CRC() uint32 {
	return 0x3c4301c0
}
func (*AttachMenuBotsPredict) _AttachMenuBots() {}

type AttachMenuBotsBot interface {
	tl.TLObject
	_AttachMenuBotsBot()
}

var (
	_ AttachMenuBotsBot = (*AttachMenuBotsBotPredict)(nil)
)

type AttachMenuBotsBotPredict struct {
	Bot   AttachMenuBot `tl:"bot"`
	Users []User        `tl:"users"`
}

func (*AttachMenuBotsBotPredict) CRC() uint32 {
	return 0x93bf667f
}
func (*AttachMenuBotsBotPredict) _AttachMenuBotsBot() {}

type AttachMenuPeerType interface {
	tl.TLObject
	_AttachMenuPeerType()
}

var (
	_ AttachMenuPeerType = (*AttachMenuPeerTypeSameBotPmPredict)(nil)
	_ AttachMenuPeerType = (*AttachMenuPeerTypeBotPmPredict)(nil)
	_ AttachMenuPeerType = (*AttachMenuPeerTypePmPredict)(nil)
	_ AttachMenuPeerType = (*AttachMenuPeerTypeChatPredict)(nil)
	_ AttachMenuPeerType = (*AttachMenuPeerTypeBroadcastPredict)(nil)
)

type AttachMenuPeerTypeSameBotPmPredict struct{}

func (*AttachMenuPeerTypeSameBotPmPredict) CRC() uint32 {
	return 0x7d6be90e
}
func (*AttachMenuPeerTypeSameBotPmPredict) _AttachMenuPeerType() {}

type AttachMenuPeerTypeBotPmPredict struct{}

func (*AttachMenuPeerTypeBotPmPredict) CRC() uint32 {
	return 0xc32bfa1a
}
func (*AttachMenuPeerTypeBotPmPredict) _AttachMenuPeerType() {}

type AttachMenuPeerTypePmPredict struct{}

func (*AttachMenuPeerTypePmPredict) CRC() uint32 {
	return 0xf146d31f
}
func (*AttachMenuPeerTypePmPredict) _AttachMenuPeerType() {}

type AttachMenuPeerTypeChatPredict struct{}

func (*AttachMenuPeerTypeChatPredict) CRC() uint32 {
	return 0x0509113f
}
func (*AttachMenuPeerTypeChatPredict) _AttachMenuPeerType() {}

type AttachMenuPeerTypeBroadcastPredict struct{}

func (*AttachMenuPeerTypeBroadcastPredict) CRC() uint32 {
	return 0x7bfbdefc
}
func (*AttachMenuPeerTypeBroadcastPredict) _AttachMenuPeerType() {}

type Authorization interface {
	tl.TLObject
	_Authorization()
}

var (
	_ Authorization = (*AuthorizationPredict)(nil)
)

type AuthorizationPredict struct {
	_                         struct{} `tl:"flags,bitflag"`
	Current                   bool     `tl:"current,omitempty:flags:0,implicit"`
	OfficialApp               bool     `tl:"official_app,omitempty:flags:1,implicit"`
	PasswordPending           bool     `tl:"password_pending,omitempty:flags:2,implicit"`
	EncryptedRequestsDisabled bool     `tl:"encrypted_requests_disabled,omitempty:flags:3,implicit"`
	CallRequestsDisabled      bool     `tl:"call_requests_disabled,omitempty:flags:4,implicit"`
	Unconfirmed               bool     `tl:"unconfirmed,omitempty:flags:5,implicit"`
	Hash                      int64    `tl:"hash"`
	DeviceModel               string   `tl:"device_model"`
	Platform                  string   `tl:"platform"`
	SystemVersion             string   `tl:"system_version"`
	APIID                     int32    `tl:"api_id"`
	AppName                   string   `tl:"app_name"`
	AppVersion                string   `tl:"app_version"`
	DateCreated               int32    `tl:"date_created"`
	DateActive                int32    `tl:"date_active"`
	Ip                        string   `tl:"ip"`
	Country                   string   `tl:"country"`
	Region                    string   `tl:"region"`
}

func (*AuthorizationPredict) CRC() uint32 {
	return 0xad01d61d
}
func (*AuthorizationPredict) _Authorization() {}

type AutoDownloadSettings interface {
	tl.TLObject
	_AutoDownloadSettings()
}

var (
	_ AutoDownloadSettings = (*AutoDownloadSettingsPredict)(nil)
)

type AutoDownloadSettingsPredict struct {
	_                             struct{} `tl:"flags,bitflag"`
	Disabled                      bool     `tl:"disabled,omitempty:flags:0,implicit"`
	VideoPreloadLarge             bool     `tl:"video_preload_large,omitempty:flags:1,implicit"`
	AudioPreloadNext              bool     `tl:"audio_preload_next,omitempty:flags:2,implicit"`
	PhonecallsLessData            bool     `tl:"phonecalls_less_data,omitempty:flags:3,implicit"`
	StoriesPreload                bool     `tl:"stories_preload,omitempty:flags:4,implicit"`
	PhotoSizeMax                  int32    `tl:"photo_size_max"`
	VideoSizeMax                  int64    `tl:"video_size_max"`
	FileSizeMax                   int64    `tl:"file_size_max"`
	VideoUploadMaxbitrate         int32    `tl:"video_upload_maxbitrate"`
	SmallQueueActiveOperationsMax int32    `tl:"small_queue_active_operations_max"`
	LargeQueueActiveOperationsMax int32    `tl:"large_queue_active_operations_max"`
}

func (*AutoDownloadSettingsPredict) CRC() uint32 {
	return 0xbaa57628
}
func (*AutoDownloadSettingsPredict) _AutoDownloadSettings() {}

type AutoSaveException interface {
	tl.TLObject
	_AutoSaveException()
}

var (
	_ AutoSaveException = (*AutoSaveExceptionPredict)(nil)
)

type AutoSaveExceptionPredict struct {
	Peer     Peer             `tl:"peer"`
	Settings AutoSaveSettings `tl:"settings"`
}

func (*AutoSaveExceptionPredict) CRC() uint32 {
	return 0x81602d47
}
func (*AutoSaveExceptionPredict) _AutoSaveException() {}

type AutoSaveSettings interface {
	tl.TLObject
	_AutoSaveSettings()
}

var (
	_ AutoSaveSettings = (*AutoSaveSettingsPredict)(nil)
)

type AutoSaveSettingsPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Photos       bool     `tl:"photos,omitempty:flags:0,implicit"`
	Videos       bool     `tl:"videos,omitempty:flags:1,implicit"`
	VideoMaxSize *int64   `tl:"video_max_size,omitempty:flags:2"`
}

func (*AutoSaveSettingsPredict) CRC() uint32 {
	return 0xc84834ce
}
func (*AutoSaveSettingsPredict) _AutoSaveSettings() {}

type AvailableEffect interface {
	tl.TLObject
	_AvailableEffect()
}

var (
	_ AvailableEffect = (*AvailableEffectPredict)(nil)
)

type AvailableEffectPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	PremiumRequired   bool     `tl:"premium_required,omitempty:flags:2,implicit"`
	ID                int64    `tl:"id"`
	Emoticon          string   `tl:"emoticon"`
	StaticIconID      *int64   `tl:"static_icon_id,omitempty:flags:0"`
	EffectStickerID   int64    `tl:"effect_sticker_id"`
	EffectAnimationID *int64   `tl:"effect_animation_id,omitempty:flags:1"`
}

func (*AvailableEffectPredict) CRC() uint32 {
	return 0x93c3e27e
}
func (*AvailableEffectPredict) _AvailableEffect() {}

type AvailableReaction interface {
	tl.TLObject
	_AvailableReaction()
}

var (
	_ AvailableReaction = (*AvailableReactionPredict)(nil)
)

type AvailableReactionPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Inactive          bool     `tl:"inactive,omitempty:flags:0,implicit"`
	Premium           bool     `tl:"premium,omitempty:flags:2,implicit"`
	Reaction          string   `tl:"reaction"`
	Title             string   `tl:"title"`
	StaticIcon        Document `tl:"static_icon"`
	AppearAnimation   Document `tl:"appear_animation"`
	SelectAnimation   Document `tl:"select_animation"`
	ActivateAnimation Document `tl:"activate_animation"`
	EffectAnimation   Document `tl:"effect_animation"`
	AroundAnimation   Document `tl:"around_animation,omitempty:flags:1"`
	CenterIcon        Document `tl:"center_icon,omitempty:flags:1"`
}

func (*AvailableReactionPredict) CRC() uint32 {
	return 0xc077ec01
}
func (*AvailableReactionPredict) _AvailableReaction() {}

type BankCardOpenURL interface {
	tl.TLObject
	_BankCardOpenURL()
}

var (
	_ BankCardOpenURL = (*BankCardOpenURLPredict)(nil)
)

type BankCardOpenURLPredict struct {
	URL  string `tl:"url"`
	Name string `tl:"name"`
}

func (*BankCardOpenURLPredict) CRC() uint32 {
	return 0xf568028a
}
func (*BankCardOpenURLPredict) _BankCardOpenURL() {}

type BaseTheme interface {
	tl.TLObject
	_BaseTheme()
}

var (
	_ BaseTheme = (*BaseThemeClassicPredict)(nil)
	_ BaseTheme = (*BaseThemeDayPredict)(nil)
	_ BaseTheme = (*BaseThemeNightPredict)(nil)
	_ BaseTheme = (*BaseThemeTintedPredict)(nil)
	_ BaseTheme = (*BaseThemeArcticPredict)(nil)
)

type BaseThemeClassicPredict struct{}

func (*BaseThemeClassicPredict) CRC() uint32 {
	return 0xc3a12462
}
func (*BaseThemeClassicPredict) _BaseTheme() {}

type BaseThemeDayPredict struct{}

func (*BaseThemeDayPredict) CRC() uint32 {
	return 0xfbd81688
}
func (*BaseThemeDayPredict) _BaseTheme() {}

type BaseThemeNightPredict struct{}

func (*BaseThemeNightPredict) CRC() uint32 {
	return 0xb7b31ea8
}
func (*BaseThemeNightPredict) _BaseTheme() {}

type BaseThemeTintedPredict struct{}

func (*BaseThemeTintedPredict) CRC() uint32 {
	return 0x6d5f77ee
}
func (*BaseThemeTintedPredict) _BaseTheme() {}

type BaseThemeArcticPredict struct{}

func (*BaseThemeArcticPredict) CRC() uint32 {
	return 0x5b11125a
}
func (*BaseThemeArcticPredict) _BaseTheme() {}

type Birthday interface {
	tl.TLObject
	_Birthday()
}

var (
	_ Birthday = (*BirthdayPredict)(nil)
)

type BirthdayPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Day   int32    `tl:"day"`
	Month int32    `tl:"month"`
	Year  *int32   `tl:"year,omitempty:flags:0"`
}

func (*BirthdayPredict) CRC() uint32 {
	return 0x6c8e1e06
}
func (*BirthdayPredict) _Birthday() {}

type Bool interface {
	tl.TLObject
	_Bool()
}

var (
	_ Bool = (*BoolFalsePredict)(nil)
	_ Bool = (*BoolTruePredict)(nil)
)

type BoolFalsePredict struct{}

func (*BoolFalsePredict) CRC() uint32 {
	return 0xbc799737
}
func (*BoolFalsePredict) _Bool() {}

type BoolTruePredict struct{}

func (*BoolTruePredict) CRC() uint32 {
	return 0x997275b5
}
func (*BoolTruePredict) _Bool() {}

type Boost interface {
	tl.TLObject
	_Boost()
}

var (
	_ Boost = (*BoostPredict)(nil)
)

type BoostPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Gift          bool     `tl:"gift,omitempty:flags:1,implicit"`
	Giveaway      bool     `tl:"giveaway,omitempty:flags:2,implicit"`
	Unclaimed     bool     `tl:"unclaimed,omitempty:flags:3,implicit"`
	ID            string   `tl:"id"`
	UserID        *int64   `tl:"user_id,omitempty:flags:0"`
	GiveawayMsgID *int32   `tl:"giveaway_msg_id,omitempty:flags:2"`
	Date          int32    `tl:"date"`
	Expires       int32    `tl:"expires"`
	UsedGiftSlug  *string  `tl:"used_gift_slug,omitempty:flags:4"`
	Multiplier    *int32   `tl:"multiplier,omitempty:flags:5"`
}

func (*BoostPredict) CRC() uint32 {
	return 0x2a1c8c71
}
func (*BoostPredict) _Boost() {}

type BotApp interface {
	tl.TLObject
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
	ID          int64    `tl:"id"`
	AccessHash  int64    `tl:"access_hash"`
	ShortName   string   `tl:"short_name"`
	Title       string   `tl:"title"`
	Description string   `tl:"description"`
	Photo       Photo    `tl:"photo"`
	Document    Document `tl:"document,omitempty:flags:0"`
	Hash        int64    `tl:"hash"`
}

func (*BotAppPredict) CRC() uint32 {
	return 0x95fcd1d6
}
func (*BotAppPredict) _BotApp() {}

type BotBusinessConnection interface {
	tl.TLObject
	_BotBusinessConnection()
}

var (
	_ BotBusinessConnection = (*BotBusinessConnectionPredict)(nil)
)

type BotBusinessConnectionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	CanReply     bool     `tl:"can_reply,omitempty:flags:0,implicit"`
	Disabled     bool     `tl:"disabled,omitempty:flags:1,implicit"`
	ConnectionID string   `tl:"connection_id"`
	UserID       int64    `tl:"user_id"`
	DcID         int32    `tl:"dc_id"`
	Date         int32    `tl:"date"`
}

func (*BotBusinessConnectionPredict) CRC() uint32 {
	return 0x896433b4
}
func (*BotBusinessConnectionPredict) _BotBusinessConnection() {}

type BotCommand interface {
	tl.TLObject
	_BotCommand()
}

var (
	_ BotCommand = (*BotCommandPredict)(nil)
)

type BotCommandPredict struct {
	Command     string `tl:"command"`
	Description string `tl:"description"`
}

func (*BotCommandPredict) CRC() uint32 {
	return 0xc27ac8c7
}
func (*BotCommandPredict) _BotCommand() {}

type BotCommandScope interface {
	tl.TLObject
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
	Peer InputPeer `tl:"peer"`
}

func (*BotCommandScopePeerPredict) CRC() uint32 {
	return 0xdb9d897d
}
func (*BotCommandScopePeerPredict) _BotCommandScope() {}

type BotCommandScopePeerAdminsPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*BotCommandScopePeerAdminsPredict) CRC() uint32 {
	return 0x3fd863d1
}
func (*BotCommandScopePeerAdminsPredict) _BotCommandScope() {}

type BotCommandScopePeerUserPredict struct {
	Peer   InputPeer `tl:"peer"`
	UserID InputUser `tl:"user_id"`
}

func (*BotCommandScopePeerUserPredict) CRC() uint32 {
	return 0x0a1321f3
}
func (*BotCommandScopePeerUserPredict) _BotCommandScope() {}

type BotInfo interface {
	tl.TLObject
	_BotInfo()
}

var (
	_ BotInfo = (*BotInfoPredict)(nil)
)

type BotInfoPredict struct {
	_                   struct{}      `tl:"flags,bitflag"`
	HasPreviewMedias    bool          `tl:"has_preview_medias,omitempty:flags:6,implicit"`
	UserID              *int64        `tl:"user_id,omitempty:flags:0"`
	Description         *string       `tl:"description,omitempty:flags:1"`
	DescriptionPhoto    Photo         `tl:"description_photo,omitempty:flags:4"`
	DescriptionDocument Document      `tl:"description_document,omitempty:flags:5"`
	Commands            []BotCommand  `tl:"commands,omitempty:flags:2"`
	MenuButton          BotMenuButton `tl:"menu_button,omitempty:flags:3"`
}

func (*BotInfoPredict) CRC() uint32 {
	return 0x8f300b57
}
func (*BotInfoPredict) _BotInfo() {}

type BotInlineMessage interface {
	tl.TLObject
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
	_           struct{}        `tl:"flags,bitflag"`
	InvertMedia bool            `tl:"invert_media,omitempty:flags:3,implicit"`
	Message     string          `tl:"message"`
	Entities    []MessageEntity `tl:"entities,omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:"reply_markup,omitempty:flags:2"`
}

func (*BotInlineMessageMediaAutoPredict) CRC() uint32 {
	return 0x764cf810
}
func (*BotInlineMessageMediaAutoPredict) _BotInlineMessage() {}

type BotInlineMessageTextPredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	NoWebpage   bool            `tl:"no_webpage,omitempty:flags:0,implicit"`
	InvertMedia bool            `tl:"invert_media,omitempty:flags:3,implicit"`
	Message     string          `tl:"message"`
	Entities    []MessageEntity `tl:"entities,omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:"reply_markup,omitempty:flags:2"`
}

func (*BotInlineMessageTextPredict) CRC() uint32 {
	return 0x8c7f65e2
}
func (*BotInlineMessageTextPredict) _BotInlineMessage() {}

type BotInlineMessageMediaGeoPredict struct {
	_                           struct{}    `tl:"flags,bitflag"`
	Geo                         GeoPoint    `tl:"geo"`
	Heading                     *int32      `tl:"heading,omitempty:flags:0"`
	Period                      *int32      `tl:"period,omitempty:flags:1"`
	ProximityNotificationRadius *int32      `tl:"proximity_notification_radius,omitempty:flags:3"`
	ReplyMarkup                 ReplyMarkup `tl:"reply_markup,omitempty:flags:2"`
}

func (*BotInlineMessageMediaGeoPredict) CRC() uint32 {
	return 0x051846fd
}
func (*BotInlineMessageMediaGeoPredict) _BotInlineMessage() {}

type BotInlineMessageMediaVenuePredict struct {
	_           struct{}    `tl:"flags,bitflag"`
	Geo         GeoPoint    `tl:"geo"`
	Title       string      `tl:"title"`
	Address     string      `tl:"address"`
	Provider    string      `tl:"provider"`
	VenueID     string      `tl:"venue_id"`
	VenueType   string      `tl:"venue_type"`
	ReplyMarkup ReplyMarkup `tl:"reply_markup,omitempty:flags:2"`
}

func (*BotInlineMessageMediaVenuePredict) CRC() uint32 {
	return 0x8a86659c
}
func (*BotInlineMessageMediaVenuePredict) _BotInlineMessage() {}

type BotInlineMessageMediaContactPredict struct {
	_           struct{}    `tl:"flags,bitflag"`
	PhoneNumber string      `tl:"phone_number"`
	FirstName   string      `tl:"first_name"`
	LastName    string      `tl:"last_name"`
	Vcard       string      `tl:"vcard"`
	ReplyMarkup ReplyMarkup `tl:"reply_markup,omitempty:flags:2"`
}

func (*BotInlineMessageMediaContactPredict) CRC() uint32 {
	return 0x18d1cdc2
}
func (*BotInlineMessageMediaContactPredict) _BotInlineMessage() {}

type BotInlineMessageMediaInvoicePredict struct {
	_                        struct{}    `tl:"flags,bitflag"`
	ShippingAddressRequested bool        `tl:"shipping_address_requested,omitempty:flags:1,implicit"`
	Test                     bool        `tl:"test,omitempty:flags:3,implicit"`
	Title                    string      `tl:"title"`
	Description              string      `tl:"description"`
	Photo                    WebDocument `tl:"photo,omitempty:flags:0"`
	Currency                 string      `tl:"currency"`
	TotalAmount              int64       `tl:"total_amount"`
	ReplyMarkup              ReplyMarkup `tl:"reply_markup,omitempty:flags:2"`
}

func (*BotInlineMessageMediaInvoicePredict) CRC() uint32 {
	return 0x354a9b09
}
func (*BotInlineMessageMediaInvoicePredict) _BotInlineMessage() {}

type BotInlineMessageMediaWebPagePredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	InvertMedia     bool            `tl:"invert_media,omitempty:flags:3,implicit"`
	ForceLargeMedia bool            `tl:"force_large_media,omitempty:flags:4,implicit"`
	ForceSmallMedia bool            `tl:"force_small_media,omitempty:flags:5,implicit"`
	Manual          bool            `tl:"manual,omitempty:flags:7,implicit"`
	Safe            bool            `tl:"safe,omitempty:flags:8,implicit"`
	Message         string          `tl:"message"`
	Entities        []MessageEntity `tl:"entities,omitempty:flags:1"`
	URL             string          `tl:"url"`
	ReplyMarkup     ReplyMarkup     `tl:"reply_markup,omitempty:flags:2"`
}

func (*BotInlineMessageMediaWebPagePredict) CRC() uint32 {
	return 0x809ad9a6
}
func (*BotInlineMessageMediaWebPagePredict) _BotInlineMessage() {}

type BotInlineResult interface {
	tl.TLObject
	_BotInlineResult()
}

var (
	_ BotInlineResult = (*BotInlineResultPredict)(nil)
	_ BotInlineResult = (*BotInlineMediaResultPredict)(nil)
)

type BotInlineResultPredict struct {
	_           struct{}         `tl:"flags,bitflag"`
	ID          string           `tl:"id"`
	Type        string           `tl:"type"`
	Title       *string          `tl:"title,omitempty:flags:1"`
	Description *string          `tl:"description,omitempty:flags:2"`
	URL         *string          `tl:"url,omitempty:flags:3"`
	Thumb       WebDocument      `tl:"thumb,omitempty:flags:4"`
	Content     WebDocument      `tl:"content,omitempty:flags:5"`
	SendMessage BotInlineMessage `tl:"send_message"`
}

func (*BotInlineResultPredict) CRC() uint32 {
	return 0x11965f3a
}
func (*BotInlineResultPredict) _BotInlineResult() {}

type BotInlineMediaResultPredict struct {
	_           struct{}         `tl:"flags,bitflag"`
	ID          string           `tl:"id"`
	Type        string           `tl:"type"`
	Photo       Photo            `tl:"photo,omitempty:flags:0"`
	Document    Document         `tl:"document,omitempty:flags:1"`
	Title       *string          `tl:"title,omitempty:flags:2"`
	Description *string          `tl:"description,omitempty:flags:3"`
	SendMessage BotInlineMessage `tl:"send_message"`
}

func (*BotInlineMediaResultPredict) CRC() uint32 {
	return 0x17db940b
}
func (*BotInlineMediaResultPredict) _BotInlineResult() {}

type BotMenuButton interface {
	tl.TLObject
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
	Text string `tl:"text"`
	URL  string `tl:"url"`
}

func (*BotMenuButtonPredict) CRC() uint32 {
	return 0xc7b57ce6
}
func (*BotMenuButtonPredict) _BotMenuButton() {}

type BotPreviewMedia interface {
	tl.TLObject
	_BotPreviewMedia()
}

var (
	_ BotPreviewMedia = (*BotPreviewMediaPredict)(nil)
)

type BotPreviewMediaPredict struct {
	Date  int32        `tl:"date"`
	Media MessageMedia `tl:"media"`
}

func (*BotPreviewMediaPredict) CRC() uint32 {
	return 0x23e91ba3
}
func (*BotPreviewMediaPredict) _BotPreviewMedia() {}

type BroadcastRevenueBalances interface {
	tl.TLObject
	_BroadcastRevenueBalances()
}

var (
	_ BroadcastRevenueBalances = (*BroadcastRevenueBalancesPredict)(nil)
)

type BroadcastRevenueBalancesPredict struct {
	CurrentBalance   int64 `tl:"current_balance"`
	AvailableBalance int64 `tl:"available_balance"`
	OverallRevenue   int64 `tl:"overall_revenue"`
}

func (*BroadcastRevenueBalancesPredict) CRC() uint32 {
	return 0x8438f1c6
}
func (*BroadcastRevenueBalancesPredict) _BroadcastRevenueBalances() {}

type BroadcastRevenueTransaction interface {
	tl.TLObject
	_BroadcastRevenueTransaction()
}

var (
	_ BroadcastRevenueTransaction = (*BroadcastRevenueTransactionProceedsPredict)(nil)
	_ BroadcastRevenueTransaction = (*BroadcastRevenueTransactionWithdrawalPredict)(nil)
	_ BroadcastRevenueTransaction = (*BroadcastRevenueTransactionRefundPredict)(nil)
)

type BroadcastRevenueTransactionProceedsPredict struct {
	Amount   int64 `tl:"amount"`
	FromDate int32 `tl:"from_date"`
	ToDate   int32 `tl:"to_date"`
}

func (*BroadcastRevenueTransactionProceedsPredict) CRC() uint32 {
	return 0x557e2cc4
}
func (*BroadcastRevenueTransactionProceedsPredict) _BroadcastRevenueTransaction() {}

type BroadcastRevenueTransactionWithdrawalPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Pending         bool     `tl:"pending,omitempty:flags:0,implicit"`
	Failed          bool     `tl:"failed,omitempty:flags:2,implicit"`
	Amount          int64    `tl:"amount"`
	Date            int32    `tl:"date"`
	Provider        string   `tl:"provider"`
	TransactionDate *int32   `tl:"transaction_date,omitempty:flags:1"`
	TransactionURL  *string  `tl:"transaction_url,omitempty:flags:1"`
}

func (*BroadcastRevenueTransactionWithdrawalPredict) CRC() uint32 {
	return 0x5a590978
}
func (*BroadcastRevenueTransactionWithdrawalPredict) _BroadcastRevenueTransaction() {}

type BroadcastRevenueTransactionRefundPredict struct {
	Amount   int64  `tl:"amount"`
	Date     int32  `tl:"date"`
	Provider string `tl:"provider"`
}

func (*BroadcastRevenueTransactionRefundPredict) CRC() uint32 {
	return 0x42d30d2e
}
func (*BroadcastRevenueTransactionRefundPredict) _BroadcastRevenueTransaction() {}

type BusinessAwayMessage interface {
	tl.TLObject
	_BusinessAwayMessage()
}

var (
	_ BusinessAwayMessage = (*BusinessAwayMessagePredict)(nil)
)

type BusinessAwayMessagePredict struct {
	_           struct{}                    `tl:"flags,bitflag"`
	OfflineOnly bool                        `tl:"offline_only,omitempty:flags:0,implicit"`
	ShortcutID  int32                       `tl:"shortcut_id"`
	Schedule    BusinessAwayMessageSchedule `tl:"schedule"`
	Recipients  BusinessRecipients          `tl:"recipients"`
}

func (*BusinessAwayMessagePredict) CRC() uint32 {
	return 0xef156a5c
}
func (*BusinessAwayMessagePredict) _BusinessAwayMessage() {}

type BusinessAwayMessageSchedule interface {
	tl.TLObject
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
	StartDate int32 `tl:"start_date"`
	EndDate   int32 `tl:"end_date"`
}

func (*BusinessAwayMessageScheduleCustomPredict) CRC() uint32 {
	return 0xcc4d9ecc
}
func (*BusinessAwayMessageScheduleCustomPredict) _BusinessAwayMessageSchedule() {}

type BusinessBotRecipients interface {
	tl.TLObject
	_BusinessBotRecipients()
}

var (
	_ BusinessBotRecipients = (*BusinessBotRecipientsPredict)(nil)
)

type BusinessBotRecipientsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ExistingChats   bool     `tl:"existing_chats,omitempty:flags:0,implicit"`
	NewChats        bool     `tl:"new_chats,omitempty:flags:1,implicit"`
	Contacts        bool     `tl:"contacts,omitempty:flags:2,implicit"`
	NonContacts     bool     `tl:"non_contacts,omitempty:flags:3,implicit"`
	ExcludeSelected bool     `tl:"exclude_selected,omitempty:flags:5,implicit"`
	Users           []int64  `tl:"users,omitempty:flags:4"`
	ExcludeUsers    []int64  `tl:"exclude_users,omitempty:flags:6"`
}

func (*BusinessBotRecipientsPredict) CRC() uint32 {
	return 0xb88cf373
}
func (*BusinessBotRecipientsPredict) _BusinessBotRecipients() {}

type BusinessChatLink interface {
	tl.TLObject
	_BusinessChatLink()
}

var (
	_ BusinessChatLink = (*BusinessChatLinkPredict)(nil)
)

type BusinessChatLinkPredict struct {
	_        struct{}        `tl:"flags,bitflag"`
	Link     string          `tl:"link"`
	Message  string          `tl:"message"`
	Entities []MessageEntity `tl:"entities,omitempty:flags:0"`
	Title    *string         `tl:"title,omitempty:flags:1"`
	Views    int32           `tl:"views"`
}

func (*BusinessChatLinkPredict) CRC() uint32 {
	return 0xb4ae666f
}
func (*BusinessChatLinkPredict) _BusinessChatLink() {}

type BusinessGreetingMessage interface {
	tl.TLObject
	_BusinessGreetingMessage()
}

var (
	_ BusinessGreetingMessage = (*BusinessGreetingMessagePredict)(nil)
)

type BusinessGreetingMessagePredict struct {
	ShortcutID     int32              `tl:"shortcut_id"`
	Recipients     BusinessRecipients `tl:"recipients"`
	NoActivityDays int32              `tl:"no_activity_days"`
}

func (*BusinessGreetingMessagePredict) CRC() uint32 {
	return 0xe519abab
}
func (*BusinessGreetingMessagePredict) _BusinessGreetingMessage() {}

type BusinessIntro interface {
	tl.TLObject
	_BusinessIntro()
}

var (
	_ BusinessIntro = (*BusinessIntroPredict)(nil)
)

type BusinessIntroPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       string   `tl:"title"`
	Description string   `tl:"description"`
	Sticker     Document `tl:"sticker,omitempty:flags:0"`
}

func (*BusinessIntroPredict) CRC() uint32 {
	return 0x5a0a066d
}
func (*BusinessIntroPredict) _BusinessIntro() {}

type BusinessLocation interface {
	tl.TLObject
	_BusinessLocation()
}

var (
	_ BusinessLocation = (*BusinessLocationPredict)(nil)
)

type BusinessLocationPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	GeoPoint GeoPoint `tl:"geo_point,omitempty:flags:0"`
	Address  string   `tl:"address"`
}

func (*BusinessLocationPredict) CRC() uint32 {
	return 0xac5c1af7
}
func (*BusinessLocationPredict) _BusinessLocation() {}

type BusinessRecipients interface {
	tl.TLObject
	_BusinessRecipients()
}

var (
	_ BusinessRecipients = (*BusinessRecipientsPredict)(nil)
)

type BusinessRecipientsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ExistingChats   bool     `tl:"existing_chats,omitempty:flags:0,implicit"`
	NewChats        bool     `tl:"new_chats,omitempty:flags:1,implicit"`
	Contacts        bool     `tl:"contacts,omitempty:flags:2,implicit"`
	NonContacts     bool     `tl:"non_contacts,omitempty:flags:3,implicit"`
	ExcludeSelected bool     `tl:"exclude_selected,omitempty:flags:5,implicit"`
	Users           []int64  `tl:"users,omitempty:flags:4"`
}

func (*BusinessRecipientsPredict) CRC() uint32 {
	return 0x21108ff7
}
func (*BusinessRecipientsPredict) _BusinessRecipients() {}

type BusinessWeeklyOpen interface {
	tl.TLObject
	_BusinessWeeklyOpen()
}

var (
	_ BusinessWeeklyOpen = (*BusinessWeeklyOpenPredict)(nil)
)

type BusinessWeeklyOpenPredict struct {
	StartMinute int32 `tl:"start_minute"`
	EndMinute   int32 `tl:"end_minute"`
}

func (*BusinessWeeklyOpenPredict) CRC() uint32 {
	return 0x120b1ab9
}
func (*BusinessWeeklyOpenPredict) _BusinessWeeklyOpen() {}

type BusinessWorkHours interface {
	tl.TLObject
	_BusinessWorkHours()
}

var (
	_ BusinessWorkHours = (*BusinessWorkHoursPredict)(nil)
)

type BusinessWorkHoursPredict struct {
	_          struct{}             `tl:"flags,bitflag"`
	OpenNow    bool                 `tl:"open_now,omitempty:flags:0,implicit"`
	TimezoneID string               `tl:"timezone_id"`
	WeeklyOpen []BusinessWeeklyOpen `tl:"weekly_open"`
}

func (*BusinessWorkHoursPredict) CRC() uint32 {
	return 0x8c92b098
}
func (*BusinessWorkHoursPredict) _BusinessWorkHours() {}

type CdnConfig interface {
	tl.TLObject
	_CdnConfig()
}

var (
	_ CdnConfig = (*CdnConfigPredict)(nil)
)

type CdnConfigPredict struct {
	PublicKeys []CdnPublicKey `tl:"public_keys"`
}

func (*CdnConfigPredict) CRC() uint32 {
	return 0x5725e40a
}
func (*CdnConfigPredict) _CdnConfig() {}

type CdnPublicKey interface {
	tl.TLObject
	_CdnPublicKey()
}

var (
	_ CdnPublicKey = (*CdnPublicKeyPredict)(nil)
)

type CdnPublicKeyPredict struct {
	DcID      int32  `tl:"dc_id"`
	PublicKey string `tl:"public_key"`
}

func (*CdnPublicKeyPredict) CRC() uint32 {
	return 0xc982eaba
}
func (*CdnPublicKeyPredict) _CdnPublicKey() {}

type ChannelAdminLogEvent interface {
	tl.TLObject
	_ChannelAdminLogEvent()
}

var (
	_ ChannelAdminLogEvent = (*ChannelAdminLogEventPredict)(nil)
)

type ChannelAdminLogEventPredict struct {
	ID     int64                      `tl:"id"`
	Date   int32                      `tl:"date"`
	UserID int64                      `tl:"user_id"`
	Action ChannelAdminLogEventAction `tl:"action"`
}

func (*ChannelAdminLogEventPredict) CRC() uint32 {
	return 0x1fad68cd
}
func (*ChannelAdminLogEventPredict) _ChannelAdminLogEvent() {}

type ChannelAdminLogEventAction interface {
	tl.TLObject
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
	PrevValue string `tl:"prev_value"`
	NewValue  string `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeTitlePredict) CRC() uint32 {
	return 0xe6dfb825
}
func (*ChannelAdminLogEventActionChangeTitlePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeAboutPredict struct {
	PrevValue string `tl:"prev_value"`
	NewValue  string `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeAboutPredict) CRC() uint32 {
	return 0x55188a2e
}
func (*ChannelAdminLogEventActionChangeAboutPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeUsernamePredict struct {
	PrevValue string `tl:"prev_value"`
	NewValue  string `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeUsernamePredict) CRC() uint32 {
	return 0x6a4afc38
}
func (*ChannelAdminLogEventActionChangeUsernamePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangePhotoPredict struct {
	PrevPhoto Photo `tl:"prev_photo"`
	NewPhoto  Photo `tl:"new_photo"`
}

func (*ChannelAdminLogEventActionChangePhotoPredict) CRC() uint32 {
	return 0x434bd2af
}
func (*ChannelAdminLogEventActionChangePhotoPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleInvitesPredict struct {
	NewValue bool `tl:"new_value"`
}

func (*ChannelAdminLogEventActionToggleInvitesPredict) CRC() uint32 {
	return 0x1b7907ae
}
func (*ChannelAdminLogEventActionToggleInvitesPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleSignaturesPredict struct {
	NewValue bool `tl:"new_value"`
}

func (*ChannelAdminLogEventActionToggleSignaturesPredict) CRC() uint32 {
	return 0x26ae0971
}
func (*ChannelAdminLogEventActionToggleSignaturesPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionUpdatePinnedPredict struct {
	Message Message `tl:"message"`
}

func (*ChannelAdminLogEventActionUpdatePinnedPredict) CRC() uint32 {
	return 0xe9e82c18
}
func (*ChannelAdminLogEventActionUpdatePinnedPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionEditMessagePredict struct {
	PrevMessage Message `tl:"prev_message"`
	NewMessage  Message `tl:"new_message"`
}

func (*ChannelAdminLogEventActionEditMessagePredict) CRC() uint32 {
	return 0x709b2405
}
func (*ChannelAdminLogEventActionEditMessagePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDeleteMessagePredict struct {
	Message Message `tl:"message"`
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
	Participant ChannelParticipant `tl:"participant"`
}

func (*ChannelAdminLogEventActionParticipantInvitePredict) CRC() uint32 {
	return 0xe31c34d8
}
func (*ChannelAdminLogEventActionParticipantInvitePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantToggleBanPredict struct {
	PrevParticipant ChannelParticipant `tl:"prev_participant"`
	NewParticipant  ChannelParticipant `tl:"new_participant"`
}

func (*ChannelAdminLogEventActionParticipantToggleBanPredict) CRC() uint32 {
	return 0xe6d83d7e
}
func (*ChannelAdminLogEventActionParticipantToggleBanPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantToggleAdminPredict struct {
	PrevParticipant ChannelParticipant `tl:"prev_participant"`
	NewParticipant  ChannelParticipant `tl:"new_participant"`
}

func (*ChannelAdminLogEventActionParticipantToggleAdminPredict) CRC() uint32 {
	return 0xd5676710
}
func (*ChannelAdminLogEventActionParticipantToggleAdminPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeStickerSetPredict struct {
	PrevStickerset InputStickerSet `tl:"prev_stickerset"`
	NewStickerset  InputStickerSet `tl:"new_stickerset"`
}

func (*ChannelAdminLogEventActionChangeStickerSetPredict) CRC() uint32 {
	return 0xb1c3caa7
}
func (*ChannelAdminLogEventActionChangeStickerSetPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionTogglePreHistoryHiddenPredict struct {
	NewValue bool `tl:"new_value"`
}

func (*ChannelAdminLogEventActionTogglePreHistoryHiddenPredict) CRC() uint32 {
	return 0x5f5c95f1
}
func (*ChannelAdminLogEventActionTogglePreHistoryHiddenPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDefaultBannedRightsPredict struct {
	PrevBannedRights ChatBannedRights `tl:"prev_banned_rights"`
	NewBannedRights  ChatBannedRights `tl:"new_banned_rights"`
}

func (*ChannelAdminLogEventActionDefaultBannedRightsPredict) CRC() uint32 {
	return 0x2df5fc0a
}
func (*ChannelAdminLogEventActionDefaultBannedRightsPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionStopPollPredict struct {
	Message Message `tl:"message"`
}

func (*ChannelAdminLogEventActionStopPollPredict) CRC() uint32 {
	return 0x8f079643
}
func (*ChannelAdminLogEventActionStopPollPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeLinkedChatPredict struct {
	PrevValue int64 `tl:"prev_value"`
	NewValue  int64 `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeLinkedChatPredict) CRC() uint32 {
	return 0x050c7ac8
}
func (*ChannelAdminLogEventActionChangeLinkedChatPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeLocationPredict struct {
	PrevValue ChannelLocation `tl:"prev_value"`
	NewValue  ChannelLocation `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeLocationPredict) CRC() uint32 {
	return 0x0e6b76ae
}
func (*ChannelAdminLogEventActionChangeLocationPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleSlowModePredict struct {
	PrevValue int32 `tl:"prev_value"`
	NewValue  int32 `tl:"new_value"`
}

func (*ChannelAdminLogEventActionToggleSlowModePredict) CRC() uint32 {
	return 0x53909779
}
func (*ChannelAdminLogEventActionToggleSlowModePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionStartGroupCallPredict struct {
	Call InputGroupCall `tl:"call"`
}

func (*ChannelAdminLogEventActionStartGroupCallPredict) CRC() uint32 {
	return 0x23209745
}
func (*ChannelAdminLogEventActionStartGroupCallPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDiscardGroupCallPredict struct {
	Call InputGroupCall `tl:"call"`
}

func (*ChannelAdminLogEventActionDiscardGroupCallPredict) CRC() uint32 {
	return 0xdb9f9140
}
func (*ChannelAdminLogEventActionDiscardGroupCallPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantMutePredict struct {
	Participant GroupCallParticipant `tl:"participant"`
}

func (*ChannelAdminLogEventActionParticipantMutePredict) CRC() uint32 {
	return 0xf92424d2
}
func (*ChannelAdminLogEventActionParticipantMutePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantUnmutePredict struct {
	Participant GroupCallParticipant `tl:"participant"`
}

func (*ChannelAdminLogEventActionParticipantUnmutePredict) CRC() uint32 {
	return 0xe64429c0
}
func (*ChannelAdminLogEventActionParticipantUnmutePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleGroupCallSettingPredict struct {
	JoinMuted bool `tl:"join_muted"`
}

func (*ChannelAdminLogEventActionToggleGroupCallSettingPredict) CRC() uint32 {
	return 0x56d6a247
}
func (*ChannelAdminLogEventActionToggleGroupCallSettingPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoinByInvitePredict struct {
	_           struct{}           `tl:"flags,bitflag"`
	ViaChatlist bool               `tl:"via_chatlist,omitempty:flags:0,implicit"`
	Invite      ExportedChatInvite `tl:"invite"`
}

func (*ChannelAdminLogEventActionParticipantJoinByInvitePredict) CRC() uint32 {
	return 0xfe9fc158
}
func (*ChannelAdminLogEventActionParticipantJoinByInvitePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteDeletePredict struct {
	Invite ExportedChatInvite `tl:"invite"`
}

func (*ChannelAdminLogEventActionExportedInviteDeletePredict) CRC() uint32 {
	return 0x5a50fca4
}
func (*ChannelAdminLogEventActionExportedInviteDeletePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteRevokePredict struct {
	Invite ExportedChatInvite `tl:"invite"`
}

func (*ChannelAdminLogEventActionExportedInviteRevokePredict) CRC() uint32 {
	return 0x410a134e
}
func (*ChannelAdminLogEventActionExportedInviteRevokePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionExportedInviteEditPredict struct {
	PrevInvite ExportedChatInvite `tl:"prev_invite"`
	NewInvite  ExportedChatInvite `tl:"new_invite"`
}

func (*ChannelAdminLogEventActionExportedInviteEditPredict) CRC() uint32 {
	return 0xe90ebb59
}
func (*ChannelAdminLogEventActionExportedInviteEditPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantVolumePredict struct {
	Participant GroupCallParticipant `tl:"participant"`
}

func (*ChannelAdminLogEventActionParticipantVolumePredict) CRC() uint32 {
	return 0x3e7f6847
}
func (*ChannelAdminLogEventActionParticipantVolumePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeHistoryTTLPredict struct {
	PrevValue int32 `tl:"prev_value"`
	NewValue  int32 `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeHistoryTTLPredict) CRC() uint32 {
	return 0x6e941a38
}
func (*ChannelAdminLogEventActionChangeHistoryTTLPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionParticipantJoinByRequestPredict struct {
	Invite     ExportedChatInvite `tl:"invite"`
	ApprovedBy int64              `tl:"approved_by"`
}

func (*ChannelAdminLogEventActionParticipantJoinByRequestPredict) CRC() uint32 {
	return 0xafb6144a
}
func (*ChannelAdminLogEventActionParticipantJoinByRequestPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleNoForwardsPredict struct {
	NewValue bool `tl:"new_value"`
}

func (*ChannelAdminLogEventActionToggleNoForwardsPredict) CRC() uint32 {
	return 0xcb2ac766
}
func (*ChannelAdminLogEventActionToggleNoForwardsPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionSendMessagePredict struct {
	Message Message `tl:"message"`
}

func (*ChannelAdminLogEventActionSendMessagePredict) CRC() uint32 {
	return 0x278f2868
}
func (*ChannelAdminLogEventActionSendMessagePredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeAvailableReactionsPredict struct {
	PrevValue ChatReactions `tl:"prev_value"`
	NewValue  ChatReactions `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeAvailableReactionsPredict) CRC() uint32 {
	return 0xbe4e0ef8
}
func (*ChannelAdminLogEventActionChangeAvailableReactionsPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeUsernamesPredict struct {
	PrevValue []string `tl:"prev_value"`
	NewValue  []string `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeUsernamesPredict) CRC() uint32 {
	return 0xf04fb3a9
}
func (*ChannelAdminLogEventActionChangeUsernamesPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleForumPredict struct {
	NewValue bool `tl:"new_value"`
}

func (*ChannelAdminLogEventActionToggleForumPredict) CRC() uint32 {
	return 0x02cc6383
}
func (*ChannelAdminLogEventActionToggleForumPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionCreateTopicPredict struct {
	Topic ForumTopic `tl:"topic"`
}

func (*ChannelAdminLogEventActionCreateTopicPredict) CRC() uint32 {
	return 0x58707d28
}
func (*ChannelAdminLogEventActionCreateTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionEditTopicPredict struct {
	PrevTopic ForumTopic `tl:"prev_topic"`
	NewTopic  ForumTopic `tl:"new_topic"`
}

func (*ChannelAdminLogEventActionEditTopicPredict) CRC() uint32 {
	return 0xf06fe208
}
func (*ChannelAdminLogEventActionEditTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionDeleteTopicPredict struct {
	Topic ForumTopic `tl:"topic"`
}

func (*ChannelAdminLogEventActionDeleteTopicPredict) CRC() uint32 {
	return 0xae168909
}
func (*ChannelAdminLogEventActionDeleteTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionPinTopicPredict struct {
	_         struct{}   `tl:"flags,bitflag"`
	PrevTopic ForumTopic `tl:"prev_topic,omitempty:flags:0"`
	NewTopic  ForumTopic `tl:"new_topic,omitempty:flags:1"`
}

func (*ChannelAdminLogEventActionPinTopicPredict) CRC() uint32 {
	return 0x5d8d353b
}
func (*ChannelAdminLogEventActionPinTopicPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionToggleAntiSpamPredict struct {
	NewValue bool `tl:"new_value"`
}

func (*ChannelAdminLogEventActionToggleAntiSpamPredict) CRC() uint32 {
	return 0x64f36dfc
}
func (*ChannelAdminLogEventActionToggleAntiSpamPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangePeerColorPredict struct {
	PrevValue PeerColor `tl:"prev_value"`
	NewValue  PeerColor `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangePeerColorPredict) CRC() uint32 {
	return 0x5796e780
}
func (*ChannelAdminLogEventActionChangePeerColorPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeProfilePeerColorPredict struct {
	PrevValue PeerColor `tl:"prev_value"`
	NewValue  PeerColor `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeProfilePeerColorPredict) CRC() uint32 {
	return 0x5e477b25
}
func (*ChannelAdminLogEventActionChangeProfilePeerColorPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeWallpaperPredict struct {
	PrevValue WallPaper `tl:"prev_value"`
	NewValue  WallPaper `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeWallpaperPredict) CRC() uint32 {
	return 0x31bb5d52
}
func (*ChannelAdminLogEventActionChangeWallpaperPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeEmojiStatusPredict struct {
	PrevValue EmojiStatus `tl:"prev_value"`
	NewValue  EmojiStatus `tl:"new_value"`
}

func (*ChannelAdminLogEventActionChangeEmojiStatusPredict) CRC() uint32 {
	return 0x3ea9feb1
}
func (*ChannelAdminLogEventActionChangeEmojiStatusPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventActionChangeEmojiStickerSetPredict struct {
	PrevStickerset InputStickerSet `tl:"prev_stickerset"`
	NewStickerset  InputStickerSet `tl:"new_stickerset"`
}

func (*ChannelAdminLogEventActionChangeEmojiStickerSetPredict) CRC() uint32 {
	return 0x46d840ab
}
func (*ChannelAdminLogEventActionChangeEmojiStickerSetPredict) _ChannelAdminLogEventAction() {}

type ChannelAdminLogEventsFilter interface {
	tl.TLObject
	_ChannelAdminLogEventsFilter()
}

var (
	_ ChannelAdminLogEventsFilter = (*ChannelAdminLogEventsFilterPredict)(nil)
)

type ChannelAdminLogEventsFilterPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Join      bool     `tl:"join,omitempty:flags:0,implicit"`
	Leave     bool     `tl:"leave,omitempty:flags:1,implicit"`
	Invite    bool     `tl:"invite,omitempty:flags:2,implicit"`
	Ban       bool     `tl:"ban,omitempty:flags:3,implicit"`
	Unban     bool     `tl:"unban,omitempty:flags:4,implicit"`
	Kick      bool     `tl:"kick,omitempty:flags:5,implicit"`
	Unkick    bool     `tl:"unkick,omitempty:flags:6,implicit"`
	Promote   bool     `tl:"promote,omitempty:flags:7,implicit"`
	Demote    bool     `tl:"demote,omitempty:flags:8,implicit"`
	Info      bool     `tl:"info,omitempty:flags:9,implicit"`
	Settings  bool     `tl:"settings,omitempty:flags:10,implicit"`
	Pinned    bool     `tl:"pinned,omitempty:flags:11,implicit"`
	Edit      bool     `tl:"edit,omitempty:flags:12,implicit"`
	Delete    bool     `tl:"delete,omitempty:flags:13,implicit"`
	GroupCall bool     `tl:"group_call,omitempty:flags:14,implicit"`
	Invites   bool     `tl:"invites,omitempty:flags:15,implicit"`
	Send      bool     `tl:"send,omitempty:flags:16,implicit"`
	Forums    bool     `tl:"forums,omitempty:flags:17,implicit"`
}

func (*ChannelAdminLogEventsFilterPredict) CRC() uint32 {
	return 0xea107ae4
}
func (*ChannelAdminLogEventsFilterPredict) _ChannelAdminLogEventsFilter() {}

type ChannelLocation interface {
	tl.TLObject
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
	GeoPoint GeoPoint `tl:"geo_point"`
	Address  string   `tl:"address"`
}

func (*ChannelLocationPredict) CRC() uint32 {
	return 0x209b82db
}
func (*ChannelLocationPredict) _ChannelLocation() {}

type ChannelMessagesFilter interface {
	tl.TLObject
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
	_                  struct{}       `tl:"flags,bitflag"`
	ExcludeNewMessages bool           `tl:"exclude_new_messages,omitempty:flags:1,implicit"`
	Ranges             []MessageRange `tl:"ranges"`
}

func (*ChannelMessagesFilterPredict) CRC() uint32 {
	return 0xcd77d957
}
func (*ChannelMessagesFilterPredict) _ChannelMessagesFilter() {}

type ChannelParticipant interface {
	tl.TLObject
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
	UserID int64 `tl:"user_id"`
	Date   int32 `tl:"date"`
}

func (*ChannelParticipantPredict) CRC() uint32 {
	return 0xc00c07c0
}
func (*ChannelParticipantPredict) _ChannelParticipant() {}

type ChannelParticipantSelfPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	ViaRequest bool     `tl:"via_request,omitempty:flags:0,implicit"`
	UserID     int64    `tl:"user_id"`
	InviterID  int64    `tl:"inviter_id"`
	Date       int32    `tl:"date"`
}

func (*ChannelParticipantSelfPredict) CRC() uint32 {
	return 0x35a8bfa7
}
func (*ChannelParticipantSelfPredict) _ChannelParticipant() {}

type ChannelParticipantCreatorPredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	UserID      int64           `tl:"user_id"`
	AdminRights ChatAdminRights `tl:"admin_rights"`
	Rank        *string         `tl:"rank,omitempty:flags:0"`
}

func (*ChannelParticipantCreatorPredict) CRC() uint32 {
	return 0x2fe601d3
}
func (*ChannelParticipantCreatorPredict) _ChannelParticipant() {}

type ChannelParticipantAdminPredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	CanEdit     bool            `tl:"can_edit,omitempty:flags:0,implicit"`
	Self        bool            `tl:"self,omitempty:flags:1,implicit"`
	UserID      int64           `tl:"user_id"`
	InviterID   *int64          `tl:"inviter_id,omitempty:flags:1"`
	PromotedBy  int64           `tl:"promoted_by"`
	Date        int32           `tl:"date"`
	AdminRights ChatAdminRights `tl:"admin_rights"`
	Rank        *string         `tl:"rank,omitempty:flags:2"`
}

func (*ChannelParticipantAdminPredict) CRC() uint32 {
	return 0x34c3bb53
}
func (*ChannelParticipantAdminPredict) _ChannelParticipant() {}

type ChannelParticipantBannedPredict struct {
	_            struct{}         `tl:"flags,bitflag"`
	Left         bool             `tl:"left,omitempty:flags:0,implicit"`
	Peer         Peer             `tl:"peer"`
	KickedBy     int64            `tl:"kicked_by"`
	Date         int32            `tl:"date"`
	BannedRights ChatBannedRights `tl:"banned_rights"`
}

func (*ChannelParticipantBannedPredict) CRC() uint32 {
	return 0x6df8014e
}
func (*ChannelParticipantBannedPredict) _ChannelParticipant() {}

type ChannelParticipantLeftPredict struct {
	Peer Peer `tl:"peer"`
}

func (*ChannelParticipantLeftPredict) CRC() uint32 {
	return 0x1b03f006
}
func (*ChannelParticipantLeftPredict) _ChannelParticipant() {}

type ChannelParticipantsFilter interface {
	tl.TLObject
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
	Q string `tl:"q"`
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
	Q string `tl:"q"`
}

func (*ChannelParticipantsBannedPredict) CRC() uint32 {
	return 0x1427a5e1
}
func (*ChannelParticipantsBannedPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsSearchPredict struct {
	Q string `tl:"q"`
}

func (*ChannelParticipantsSearchPredict) CRC() uint32 {
	return 0x0656ac4b
}
func (*ChannelParticipantsSearchPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsContactsPredict struct {
	Q string `tl:"q"`
}

func (*ChannelParticipantsContactsPredict) CRC() uint32 {
	return 0xbb6ae88d
}
func (*ChannelParticipantsContactsPredict) _ChannelParticipantsFilter() {}

type ChannelParticipantsMentionsPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Q        *string  `tl:"q,omitempty:flags:0"`
	TopMsgID *int32   `tl:"top_msg_id,omitempty:flags:1"`
}

func (*ChannelParticipantsMentionsPredict) CRC() uint32 {
	return 0xe04b5ceb
}
func (*ChannelParticipantsMentionsPredict) _ChannelParticipantsFilter() {}

type Chat interface {
	tl.TLObject
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
	ID int64 `tl:"id"`
}

func (*ChatEmptyPredict) CRC() uint32 {
	return 0x29562865
}
func (*ChatEmptyPredict) _Chat() {}

type ChatPredict struct {
	_                   struct{}         `tl:"flags,bitflag"`
	Creator             bool             `tl:"creator,omitempty:flags:0,implicit"`
	Left                bool             `tl:"left,omitempty:flags:2,implicit"`
	Deactivated         bool             `tl:"deactivated,omitempty:flags:5,implicit"`
	CallActive          bool             `tl:"call_active,omitempty:flags:23,implicit"`
	CallNotEmpty        bool             `tl:"call_not_empty,omitempty:flags:24,implicit"`
	Noforwards          bool             `tl:"noforwards,omitempty:flags:25,implicit"`
	ID                  int64            `tl:"id"`
	Title               string           `tl:"title"`
	Photo               ChatPhoto        `tl:"photo"`
	ParticipantsCount   int32            `tl:"participants_count"`
	Date                int32            `tl:"date"`
	Version             int32            `tl:"version"`
	MigratedTo          InputChannel     `tl:"migrated_to,omitempty:flags:6"`
	AdminRights         ChatAdminRights  `tl:"admin_rights,omitempty:flags:14"`
	DefaultBannedRights ChatBannedRights `tl:"default_banned_rights,omitempty:flags:18"`
}

func (*ChatPredict) CRC() uint32 {
	return 0x41cbf256
}
func (*ChatPredict) _Chat() {}

type ChatForbiddenPredict struct {
	ID    int64  `tl:"id"`
	Title string `tl:"title"`
}

func (*ChatForbiddenPredict) CRC() uint32 {
	return 0x6592a1a7
}
func (*ChatForbiddenPredict) _Chat() {}

type ChannelPredict struct {
	_                   struct{}            `tl:"flags,bitflag"`
	Creator             bool                `tl:"creator,omitempty:flags:0,implicit"`
	Left                bool                `tl:"left,omitempty:flags:2,implicit"`
	Broadcast           bool                `tl:"broadcast,omitempty:flags:5,implicit"`
	Verified            bool                `tl:"verified,omitempty:flags:7,implicit"`
	Megagroup           bool                `tl:"megagroup,omitempty:flags:8,implicit"`
	Restricted          bool                `tl:"restricted,omitempty:flags:9,implicit"`
	Signatures          bool                `tl:"signatures,omitempty:flags:11,implicit"`
	Min                 bool                `tl:"min,omitempty:flags:12,implicit"`
	Scam                bool                `tl:"scam,omitempty:flags:19,implicit"`
	HasLink             bool                `tl:"has_link,omitempty:flags:20,implicit"`
	HasGeo              bool                `tl:"has_geo,omitempty:flags:21,implicit"`
	SlowmodeEnabled     bool                `tl:"slowmode_enabled,omitempty:flags:22,implicit"`
	CallActive          bool                `tl:"call_active,omitempty:flags:23,implicit"`
	CallNotEmpty        bool                `tl:"call_not_empty,omitempty:flags:24,implicit"`
	Fake                bool                `tl:"fake,omitempty:flags:25,implicit"`
	Gigagroup           bool                `tl:"gigagroup,omitempty:flags:26,implicit"`
	Noforwards          bool                `tl:"noforwards,omitempty:flags:27,implicit"`
	JoinToSend          bool                `tl:"join_to_send,omitempty:flags:28,implicit"`
	JoinRequest         bool                `tl:"join_request,omitempty:flags:29,implicit"`
	Forum               bool                `tl:"forum,omitempty:flags:30,implicit"`
	_                   struct{}            `tl:"flags2,bitflag"`
	StoriesHidden       bool                `tl:"stories_hidden,omitempty:flags2:1,implicit"`
	StoriesHiddenMin    bool                `tl:"stories_hidden_min,omitempty:flags2:2,implicit"`
	StoriesUnavailable  bool                `tl:"stories_unavailable,omitempty:flags2:3,implicit"`
	ID                  int64               `tl:"id"`
	AccessHash          *int64              `tl:"access_hash,omitempty:flags:13"`
	Title               string              `tl:"title"`
	Username            *string             `tl:"username,omitempty:flags:6"`
	Photo               ChatPhoto           `tl:"photo"`
	Date                int32               `tl:"date"`
	RestrictionReason   []RestrictionReason `tl:"restriction_reason,omitempty:flags:9"`
	AdminRights         ChatAdminRights     `tl:"admin_rights,omitempty:flags:14"`
	BannedRights        ChatBannedRights    `tl:"banned_rights,omitempty:flags:15"`
	DefaultBannedRights ChatBannedRights    `tl:"default_banned_rights,omitempty:flags:18"`
	ParticipantsCount   *int32              `tl:"participants_count,omitempty:flags:17"`
	Usernames           []Username          `tl:"usernames,omitempty:flags2:0"`
	StoriesMaxID        *int32              `tl:"stories_max_id,omitempty:flags2:4"`
	Color               PeerColor           `tl:"color,omitempty:flags2:7"`
	ProfileColor        PeerColor           `tl:"profile_color,omitempty:flags2:8"`
	EmojiStatus         EmojiStatus         `tl:"emoji_status,omitempty:flags2:9"`
	Level               *int32              `tl:"level,omitempty:flags2:10"`
}

func (*ChannelPredict) CRC() uint32 {
	return 0x0aadfc8f
}
func (*ChannelPredict) _Chat() {}

type ChannelForbiddenPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Broadcast  bool     `tl:"broadcast,omitempty:flags:5,implicit"`
	Megagroup  bool     `tl:"megagroup,omitempty:flags:8,implicit"`
	ID         int64    `tl:"id"`
	AccessHash int64    `tl:"access_hash"`
	Title      string   `tl:"title"`
	UntilDate  *int32   `tl:"until_date,omitempty:flags:16"`
}

func (*ChannelForbiddenPredict) CRC() uint32 {
	return 0x17d493d5
}
func (*ChannelForbiddenPredict) _Chat() {}

type ChatAdminRights interface {
	tl.TLObject
	_ChatAdminRights()
}

var (
	_ ChatAdminRights = (*ChatAdminRightsPredict)(nil)
)

type ChatAdminRightsPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ChangeInfo     bool     `tl:"change_info,omitempty:flags:0,implicit"`
	PostMessages   bool     `tl:"post_messages,omitempty:flags:1,implicit"`
	EditMessages   bool     `tl:"edit_messages,omitempty:flags:2,implicit"`
	DeleteMessages bool     `tl:"delete_messages,omitempty:flags:3,implicit"`
	BanUsers       bool     `tl:"ban_users,omitempty:flags:4,implicit"`
	InviteUsers    bool     `tl:"invite_users,omitempty:flags:5,implicit"`
	PinMessages    bool     `tl:"pin_messages,omitempty:flags:7,implicit"`
	AddAdmins      bool     `tl:"add_admins,omitempty:flags:9,implicit"`
	Anonymous      bool     `tl:"anonymous,omitempty:flags:10,implicit"`
	ManageCall     bool     `tl:"manage_call,omitempty:flags:11,implicit"`
	Other          bool     `tl:"other,omitempty:flags:12,implicit"`
	ManageTopics   bool     `tl:"manage_topics,omitempty:flags:13,implicit"`
	PostStories    bool     `tl:"post_stories,omitempty:flags:14,implicit"`
	EditStories    bool     `tl:"edit_stories,omitempty:flags:15,implicit"`
	DeleteStories  bool     `tl:"delete_stories,omitempty:flags:16,implicit"`
}

func (*ChatAdminRightsPredict) CRC() uint32 {
	return 0x5fb224d5
}
func (*ChatAdminRightsPredict) _ChatAdminRights() {}

type ChatAdminWithInvites interface {
	tl.TLObject
	_ChatAdminWithInvites()
}

var (
	_ ChatAdminWithInvites = (*ChatAdminWithInvitesPredict)(nil)
)

type ChatAdminWithInvitesPredict struct {
	AdminID             int64 `tl:"admin_id"`
	InvitesCount        int32 `tl:"invites_count"`
	RevokedInvitesCount int32 `tl:"revoked_invites_count"`
}

func (*ChatAdminWithInvitesPredict) CRC() uint32 {
	return 0xf2ecef23
}
func (*ChatAdminWithInvitesPredict) _ChatAdminWithInvites() {}

type ChatBannedRights interface {
	tl.TLObject
	_ChatBannedRights()
}

var (
	_ ChatBannedRights = (*ChatBannedRightsPredict)(nil)
)

type ChatBannedRightsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ViewMessages    bool     `tl:"view_messages,omitempty:flags:0,implicit"`
	SendMessages    bool     `tl:"send_messages,omitempty:flags:1,implicit"`
	SendMedia       bool     `tl:"send_media,omitempty:flags:2,implicit"`
	SendStickers    bool     `tl:"send_stickers,omitempty:flags:3,implicit"`
	SendGifs        bool     `tl:"send_gifs,omitempty:flags:4,implicit"`
	SendGames       bool     `tl:"send_games,omitempty:flags:5,implicit"`
	SendInline      bool     `tl:"send_inline,omitempty:flags:6,implicit"`
	EmbedLinks      bool     `tl:"embed_links,omitempty:flags:7,implicit"`
	SendPolls       bool     `tl:"send_polls,omitempty:flags:8,implicit"`
	ChangeInfo      bool     `tl:"change_info,omitempty:flags:10,implicit"`
	InviteUsers     bool     `tl:"invite_users,omitempty:flags:15,implicit"`
	PinMessages     bool     `tl:"pin_messages,omitempty:flags:17,implicit"`
	ManageTopics    bool     `tl:"manage_topics,omitempty:flags:18,implicit"`
	SendPhotos      bool     `tl:"send_photos,omitempty:flags:19,implicit"`
	SendVideos      bool     `tl:"send_videos,omitempty:flags:20,implicit"`
	SendRoundvideos bool     `tl:"send_roundvideos,omitempty:flags:21,implicit"`
	SendAudios      bool     `tl:"send_audios,omitempty:flags:22,implicit"`
	SendVoices      bool     `tl:"send_voices,omitempty:flags:23,implicit"`
	SendDocs        bool     `tl:"send_docs,omitempty:flags:24,implicit"`
	SendPlain       bool     `tl:"send_plain,omitempty:flags:25,implicit"`
	UntilDate       int32    `tl:"until_date"`
}

func (*ChatBannedRightsPredict) CRC() uint32 {
	return 0x9f120418
}
func (*ChatBannedRightsPredict) _ChatBannedRights() {}

type ChatFull interface {
	tl.TLObject
	_ChatFull()
}

var (
	_ ChatFull = (*ChatFullPredict)(nil)
	_ ChatFull = (*ChannelFullPredict)(nil)
)

type ChatFullPredict struct {
	_                      struct{}           `tl:"flags,bitflag"`
	CanSetUsername         bool               `tl:"can_set_username,omitempty:flags:7,implicit"`
	HasScheduled           bool               `tl:"has_scheduled,omitempty:flags:8,implicit"`
	TranslationsDisabled   bool               `tl:"translations_disabled,omitempty:flags:19,implicit"`
	ID                     int64              `tl:"id"`
	About                  string             `tl:"about"`
	Participants           ChatParticipants   `tl:"participants"`
	ChatPhoto              Photo              `tl:"chat_photo,omitempty:flags:2"`
	NotifySettings         PeerNotifySettings `tl:"notify_settings"`
	ExportedInvite         ExportedChatInvite `tl:"exported_invite,omitempty:flags:13"`
	BotInfo                []BotInfo          `tl:"bot_info,omitempty:flags:3"`
	PinnedMsgID            *int32             `tl:"pinned_msg_id,omitempty:flags:6"`
	FolderID               *int32             `tl:"folder_id,omitempty:flags:11"`
	Call                   InputGroupCall     `tl:"call,omitempty:flags:12"`
	TTLPeriod              *int32             `tl:"ttl_period,omitempty:flags:14"`
	GroupcallDefaultJoinAs Peer               `tl:"groupcall_default_join_as,omitempty:flags:15"`
	ThemeEmoticon          *string            `tl:"theme_emoticon,omitempty:flags:16"`
	RequestsPending        *int32             `tl:"requests_pending,omitempty:flags:17"`
	RecentRequesters       []int64            `tl:"recent_requesters,omitempty:flags:17"`
	AvailableReactions     ChatReactions      `tl:"available_reactions,omitempty:flags:18"`
	ReactionsLimit         *int32             `tl:"reactions_limit,omitempty:flags:20"`
}

func (*ChatFullPredict) CRC() uint32 {
	return 0x2633421b
}
func (*ChatFullPredict) _ChatFull() {}

type ChannelFullPredict struct {
	_                      struct{}           `tl:"flags,bitflag"`
	CanViewParticipants    bool               `tl:"can_view_participants,omitempty:flags:3,implicit"`
	CanSetUsername         bool               `tl:"can_set_username,omitempty:flags:6,implicit"`
	CanSetStickers         bool               `tl:"can_set_stickers,omitempty:flags:7,implicit"`
	HiddenPrehistory       bool               `tl:"hidden_prehistory,omitempty:flags:10,implicit"`
	CanSetLocation         bool               `tl:"can_set_location,omitempty:flags:16,implicit"`
	HasScheduled           bool               `tl:"has_scheduled,omitempty:flags:19,implicit"`
	CanViewStats           bool               `tl:"can_view_stats,omitempty:flags:20,implicit"`
	Blocked                bool               `tl:"blocked,omitempty:flags:22,implicit"`
	_                      struct{}           `tl:"flags2,bitflag"`
	CanDeleteChannel       bool               `tl:"can_delete_channel,omitempty:flags2:0,implicit"`
	Antispam               bool               `tl:"antispam,omitempty:flags2:1,implicit"`
	ParticipantsHidden     bool               `tl:"participants_hidden,omitempty:flags2:2,implicit"`
	TranslationsDisabled   bool               `tl:"translations_disabled,omitempty:flags2:3,implicit"`
	StoriesPinnedAvailable bool               `tl:"stories_pinned_available,omitempty:flags2:5,implicit"`
	ViewForumAsMessages    bool               `tl:"view_forum_as_messages,omitempty:flags2:6,implicit"`
	RestrictedSponsored    bool               `tl:"restricted_sponsored,omitempty:flags2:11,implicit"`
	CanViewRevenue         bool               `tl:"can_view_revenue,omitempty:flags2:12,implicit"`
	PaidMediaAllowed       bool               `tl:"paid_media_allowed,omitempty:flags2:14,implicit"`
	CanViewStarsRevenue    bool               `tl:"can_view_stars_revenue,omitempty:flags2:15,implicit"`
	ID                     int64              `tl:"id"`
	About                  string             `tl:"about"`
	ParticipantsCount      *int32             `tl:"participants_count,omitempty:flags:0"`
	AdminsCount            *int32             `tl:"admins_count,omitempty:flags:1"`
	KickedCount            *int32             `tl:"kicked_count,omitempty:flags:2"`
	BannedCount            *int32             `tl:"banned_count,omitempty:flags:2"`
	OnlineCount            *int32             `tl:"online_count,omitempty:flags:13"`
	ReadInboxMaxID         int32              `tl:"read_inbox_max_id"`
	ReadOutboxMaxID        int32              `tl:"read_outbox_max_id"`
	UnreadCount            int32              `tl:"unread_count"`
	ChatPhoto              Photo              `tl:"chat_photo"`
	NotifySettings         PeerNotifySettings `tl:"notify_settings"`
	ExportedInvite         ExportedChatInvite `tl:"exported_invite,omitempty:flags:23"`
	BotInfo                []BotInfo          `tl:"bot_info"`
	MigratedFromChatID     *int64             `tl:"migrated_from_chat_id,omitempty:flags:4"`
	MigratedFromMaxID      *int32             `tl:"migrated_from_max_id,omitempty:flags:4"`
	PinnedMsgID            *int32             `tl:"pinned_msg_id,omitempty:flags:5"`
	Stickerset             StickerSet         `tl:"stickerset,omitempty:flags:8"`
	AvailableMinID         *int32             `tl:"available_min_id,omitempty:flags:9"`
	FolderID               *int32             `tl:"folder_id,omitempty:flags:11"`
	LinkedChatID           *int64             `tl:"linked_chat_id,omitempty:flags:14"`
	Location               ChannelLocation    `tl:"location,omitempty:flags:15"`
	SlowmodeSeconds        *int32             `tl:"slowmode_seconds,omitempty:flags:17"`
	SlowmodeNextSendDate   *int32             `tl:"slowmode_next_send_date,omitempty:flags:18"`
	StatsDc                *int32             `tl:"stats_dc,omitempty:flags:12"`
	Pts                    int32              `tl:"pts"`
	Call                   InputGroupCall     `tl:"call,omitempty:flags:21"`
	TTLPeriod              *int32             `tl:"ttl_period,omitempty:flags:24"`
	PendingSuggestions     []string           `tl:"pending_suggestions,omitempty:flags:25"`
	GroupcallDefaultJoinAs Peer               `tl:"groupcall_default_join_as,omitempty:flags:26"`
	ThemeEmoticon          *string            `tl:"theme_emoticon,omitempty:flags:27"`
	RequestsPending        *int32             `tl:"requests_pending,omitempty:flags:28"`
	RecentRequesters       []int64            `tl:"recent_requesters,omitempty:flags:28"`
	DefaultSendAs          Peer               `tl:"default_send_as,omitempty:flags:29"`
	AvailableReactions     ChatReactions      `tl:"available_reactions,omitempty:flags:30"`
	ReactionsLimit         *int32             `tl:"reactions_limit,omitempty:flags2:13"`
	Stories                PeerStories        `tl:"stories,omitempty:flags2:4"`
	Wallpaper              WallPaper          `tl:"wallpaper,omitempty:flags2:7"`
	BoostsApplied          *int32             `tl:"boosts_applied,omitempty:flags2:8"`
	BoostsUnrestrict       *int32             `tl:"boosts_unrestrict,omitempty:flags2:9"`
	Emojiset               StickerSet         `tl:"emojiset,omitempty:flags2:10"`
}

func (*ChannelFullPredict) CRC() uint32 {
	return 0xbbab348d
}
func (*ChannelFullPredict) _ChatFull() {}

type ChatInvite interface {
	tl.TLObject
	_ChatInvite()
}

var (
	_ ChatInvite = (*ChatInviteAlreadyPredict)(nil)
	_ ChatInvite = (*ChatInvitePredict)(nil)
	_ ChatInvite = (*ChatInvitePeekPredict)(nil)
)

type ChatInviteAlreadyPredict struct {
	Chat Chat `tl:"chat"`
}

func (*ChatInviteAlreadyPredict) CRC() uint32 {
	return 0x5a686d7c
}
func (*ChatInviteAlreadyPredict) _ChatInvite() {}

type ChatInvitePredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Channel           bool     `tl:"channel,omitempty:flags:0,implicit"`
	Broadcast         bool     `tl:"broadcast,omitempty:flags:1,implicit"`
	Public            bool     `tl:"public,omitempty:flags:2,implicit"`
	Megagroup         bool     `tl:"megagroup,omitempty:flags:3,implicit"`
	RequestNeeded     bool     `tl:"request_needed,omitempty:flags:6,implicit"`
	Verified          bool     `tl:"verified,omitempty:flags:7,implicit"`
	Scam              bool     `tl:"scam,omitempty:flags:8,implicit"`
	Fake              bool     `tl:"fake,omitempty:flags:9,implicit"`
	Title             string   `tl:"title"`
	About             *string  `tl:"about,omitempty:flags:5"`
	Photo             Photo    `tl:"photo"`
	ParticipantsCount int32    `tl:"participants_count"`
	Participants      []User   `tl:"participants,omitempty:flags:4"`
	Color             int32    `tl:"color"`
}

func (*ChatInvitePredict) CRC() uint32 {
	return 0xcde0ec40
}
func (*ChatInvitePredict) _ChatInvite() {}

type ChatInvitePeekPredict struct {
	Chat    Chat  `tl:"chat"`
	Expires int32 `tl:"expires"`
}

func (*ChatInvitePeekPredict) CRC() uint32 {
	return 0x61695cb0
}
func (*ChatInvitePeekPredict) _ChatInvite() {}

type ChatInviteImporter interface {
	tl.TLObject
	_ChatInviteImporter()
}

var (
	_ ChatInviteImporter = (*ChatInviteImporterPredict)(nil)
)

type ChatInviteImporterPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Requested   bool     `tl:"requested,omitempty:flags:0,implicit"`
	ViaChatlist bool     `tl:"via_chatlist,omitempty:flags:3,implicit"`
	UserID      int64    `tl:"user_id"`
	Date        int32    `tl:"date"`
	About       *string  `tl:"about,omitempty:flags:2"`
	ApprovedBy  *int64   `tl:"approved_by,omitempty:flags:1"`
}

func (*ChatInviteImporterPredict) CRC() uint32 {
	return 0x8c5adfd9
}
func (*ChatInviteImporterPredict) _ChatInviteImporter() {}

type ChatOnlines interface {
	tl.TLObject
	_ChatOnlines()
}

var (
	_ ChatOnlines = (*ChatOnlinesPredict)(nil)
)

type ChatOnlinesPredict struct {
	Onlines int32 `tl:"onlines"`
}

func (*ChatOnlinesPredict) CRC() uint32 {
	return 0xf041e250
}
func (*ChatOnlinesPredict) _ChatOnlines() {}

type ChatParticipant interface {
	tl.TLObject
	_ChatParticipant()
}

var (
	_ ChatParticipant = (*ChatParticipantPredict)(nil)
	_ ChatParticipant = (*ChatParticipantCreatorPredict)(nil)
	_ ChatParticipant = (*ChatParticipantAdminPredict)(nil)
)

type ChatParticipantPredict struct {
	UserID    int64 `tl:"user_id"`
	InviterID int64 `tl:"inviter_id"`
	Date      int32 `tl:"date"`
}

func (*ChatParticipantPredict) CRC() uint32 {
	return 0xc02d4007
}
func (*ChatParticipantPredict) _ChatParticipant() {}

type ChatParticipantCreatorPredict struct {
	UserID int64 `tl:"user_id"`
}

func (*ChatParticipantCreatorPredict) CRC() uint32 {
	return 0xe46bcee4
}
func (*ChatParticipantCreatorPredict) _ChatParticipant() {}

type ChatParticipantAdminPredict struct {
	UserID    int64 `tl:"user_id"`
	InviterID int64 `tl:"inviter_id"`
	Date      int32 `tl:"date"`
}

func (*ChatParticipantAdminPredict) CRC() uint32 {
	return 0xa0933f5b
}
func (*ChatParticipantAdminPredict) _ChatParticipant() {}

type ChatParticipants interface {
	tl.TLObject
	_ChatParticipants()
}

var (
	_ ChatParticipants = (*ChatParticipantsForbiddenPredict)(nil)
	_ ChatParticipants = (*ChatParticipantsPredict)(nil)
)

type ChatParticipantsForbiddenPredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	ChatID          int64           `tl:"chat_id"`
	SelfParticipant ChatParticipant `tl:"self_participant,omitempty:flags:0"`
}

func (*ChatParticipantsForbiddenPredict) CRC() uint32 {
	return 0x8763d3e1
}
func (*ChatParticipantsForbiddenPredict) _ChatParticipants() {}

type ChatParticipantsPredict struct {
	ChatID       int64             `tl:"chat_id"`
	Participants []ChatParticipant `tl:"participants"`
	Version      int32             `tl:"version"`
}

func (*ChatParticipantsPredict) CRC() uint32 {
	return 0x3cbc93f8
}
func (*ChatParticipantsPredict) _ChatParticipants() {}

type ChatPhoto interface {
	tl.TLObject
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
	HasVideo      bool     `tl:"has_video,omitempty:flags:0,implicit"`
	PhotoID       int64    `tl:"photo_id"`
	StrippedThumb *[]byte  `tl:"stripped_thumb,omitempty:flags:1"`
	DcID          int32    `tl:"dc_id"`
}

func (*ChatPhotoPredict) CRC() uint32 {
	return 0x1c6e1c11
}
func (*ChatPhotoPredict) _ChatPhoto() {}

type ChatReactions interface {
	tl.TLObject
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
	AllowCustom bool     `tl:"allow_custom,omitempty:flags:0,implicit"`
}

func (*ChatReactionsAllPredict) CRC() uint32 {
	return 0x52928bca
}
func (*ChatReactionsAllPredict) _ChatReactions() {}

type ChatReactionsSomePredict struct {
	Reactions []Reaction `tl:"reactions"`
}

func (*ChatReactionsSomePredict) CRC() uint32 {
	return 0x661d4037
}
func (*ChatReactionsSomePredict) _ChatReactions() {}

type CodeSettings interface {
	tl.TLObject
	_CodeSettings()
}

var (
	_ CodeSettings = (*CodeSettingsPredict)(nil)
)

type CodeSettingsPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	AllowFlashcall  bool     `tl:"allow_flashcall,omitempty:flags:0,implicit"`
	CurrentNumber   bool     `tl:"current_number,omitempty:flags:1,implicit"`
	AllowAppHash    bool     `tl:"allow_app_hash,omitempty:flags:4,implicit"`
	AllowMissedCall bool     `tl:"allow_missed_call,omitempty:flags:5,implicit"`
	AllowFirebase   bool     `tl:"allow_firebase,omitempty:flags:7,implicit"`
	UnknownNumber   bool     `tl:"unknown_number,omitempty:flags:9,implicit"`
	LogoutTokens    [][]byte `tl:"logout_tokens,omitempty:flags:6"`
	Token           *string  `tl:"token,omitempty:flags:8"`
	AppSandbox      *bool    `tl:"app_sandbox,omitempty:flags:8"`
}

func (*CodeSettingsPredict) CRC() uint32 {
	return 0xad253d78
}
func (*CodeSettingsPredict) _CodeSettings() {}

type Config interface {
	tl.TLObject
	_Config()
}

var (
	_ Config = (*ConfigPredict)(nil)
)

type ConfigPredict struct {
	_                       struct{}   `tl:"flags,bitflag"`
	DefaultP2PContacts      bool       `tl:"default_p2p_contacts,omitempty:flags:3,implicit"`
	PreloadFeaturedStickers bool       `tl:"preload_featured_stickers,omitempty:flags:4,implicit"`
	RevokePmInbox           bool       `tl:"revoke_pm_inbox,omitempty:flags:6,implicit"`
	BlockedMode             bool       `tl:"blocked_mode,omitempty:flags:8,implicit"`
	ForceTryIpv6            bool       `tl:"force_try_ipv6,omitempty:flags:14,implicit"`
	Date                    int32      `tl:"date"`
	Expires                 int32      `tl:"expires"`
	TestMode                bool       `tl:"test_mode"`
	ThisDc                  int32      `tl:"this_dc"`
	DcOptions               []DcOption `tl:"dc_options"`
	DcTxtDomainName         string     `tl:"dc_txt_domain_name"`
	ChatSizeMax             int32      `tl:"chat_size_max"`
	MegagroupSizeMax        int32      `tl:"megagroup_size_max"`
	ForwardedCountMax       int32      `tl:"forwarded_count_max"`
	OnlineUpdatePeriodMs    int32      `tl:"online_update_period_ms"`
	OfflineBlurTimeoutMs    int32      `tl:"offline_blur_timeout_ms"`
	OfflineIdleTimeoutMs    int32      `tl:"offline_idle_timeout_ms"`
	OnlineCloudTimeoutMs    int32      `tl:"online_cloud_timeout_ms"`
	NotifyCloudDelayMs      int32      `tl:"notify_cloud_delay_ms"`
	NotifyDefaultDelayMs    int32      `tl:"notify_default_delay_ms"`
	PushChatPeriodMs        int32      `tl:"push_chat_period_ms"`
	PushChatLimit           int32      `tl:"push_chat_limit"`
	EditTimeLimit           int32      `tl:"edit_time_limit"`
	RevokeTimeLimit         int32      `tl:"revoke_time_limit"`
	RevokePmTimeLimit       int32      `tl:"revoke_pm_time_limit"`
	RatingEDecay            int32      `tl:"rating_e_decay"`
	StickersRecentLimit     int32      `tl:"stickers_recent_limit"`
	ChannelsReadMediaPeriod int32      `tl:"channels_read_media_period"`
	TmpSessions             *int32     `tl:"tmp_sessions,omitempty:flags:0"`
	CallReceiveTimeoutMs    int32      `tl:"call_receive_timeout_ms"`
	CallRingTimeoutMs       int32      `tl:"call_ring_timeout_ms"`
	CallConnectTimeoutMs    int32      `tl:"call_connect_timeout_ms"`
	CallPacketTimeoutMs     int32      `tl:"call_packet_timeout_ms"`
	MeURLPrefix             string     `tl:"me_url_prefix"`
	AutoupdateURLPrefix     *string    `tl:"autoupdate_url_prefix,omitempty:flags:7"`
	GifSearchUsername       *string    `tl:"gif_search_username,omitempty:flags:9"`
	VenueSearchUsername     *string    `tl:"venue_search_username,omitempty:flags:10"`
	ImgSearchUsername       *string    `tl:"img_search_username,omitempty:flags:11"`
	StaticMapsProvider      *string    `tl:"static_maps_provider,omitempty:flags:12"`
	CaptionLengthMax        int32      `tl:"caption_length_max"`
	MessageLengthMax        int32      `tl:"message_length_max"`
	WebfileDcID             int32      `tl:"webfile_dc_id"`
	SuggestedLangCode       *string    `tl:"suggested_lang_code,omitempty:flags:2"`
	LangPackVersion         *int32     `tl:"lang_pack_version,omitempty:flags:2"`
	BaseLangPackVersion     *int32     `tl:"base_lang_pack_version,omitempty:flags:2"`
	ReactionsDefault        Reaction   `tl:"reactions_default,omitempty:flags:15"`
	AutologinToken          *string    `tl:"autologin_token,omitempty:flags:16"`
}

func (*ConfigPredict) CRC() uint32 {
	return 0xcc1a241e
}
func (*ConfigPredict) _Config() {}

type ConnectedBot interface {
	tl.TLObject
	_ConnectedBot()
}

var (
	_ ConnectedBot = (*ConnectedBotPredict)(nil)
)

type ConnectedBotPredict struct {
	_          struct{}              `tl:"flags,bitflag"`
	CanReply   bool                  `tl:"can_reply,omitempty:flags:0,implicit"`
	BotID      int64                 `tl:"bot_id"`
	Recipients BusinessBotRecipients `tl:"recipients"`
}

func (*ConnectedBotPredict) CRC() uint32 {
	return 0xbd068601
}
func (*ConnectedBotPredict) _ConnectedBot() {}

type Contact interface {
	tl.TLObject
	_Contact()
}

var (
	_ Contact = (*ContactPredict)(nil)
)

type ContactPredict struct {
	UserID int64 `tl:"user_id"`
	Mutual bool  `tl:"mutual"`
}

func (*ContactPredict) CRC() uint32 {
	return 0x145ade0b
}
func (*ContactPredict) _Contact() {}

type ContactBirthday interface {
	tl.TLObject
	_ContactBirthday()
}

var (
	_ ContactBirthday = (*ContactBirthdayPredict)(nil)
)

type ContactBirthdayPredict struct {
	ContactID int64    `tl:"contact_id"`
	Birthday  Birthday `tl:"birthday"`
}

func (*ContactBirthdayPredict) CRC() uint32 {
	return 0x1d998733
}
func (*ContactBirthdayPredict) _ContactBirthday() {}

type ContactStatus interface {
	tl.TLObject
	_ContactStatus()
}

var (
	_ ContactStatus = (*ContactStatusPredict)(nil)
)

type ContactStatusPredict struct {
	UserID int64      `tl:"user_id"`
	Status UserStatus `tl:"status"`
}

func (*ContactStatusPredict) CRC() uint32 {
	return 0x16d9703b
}
func (*ContactStatusPredict) _ContactStatus() {}

type DataJSON interface {
	tl.TLObject
	_DataJSON()
}

var (
	_ DataJSON = (*DataJSONPredict)(nil)
)

type DataJSONPredict struct {
	Data string `tl:"data"`
}

func (*DataJSONPredict) CRC() uint32 {
	return 0x7d748d04
}
func (*DataJSONPredict) _DataJSON() {}

type DcOption interface {
	tl.TLObject
	_DcOption()
}

var (
	_ DcOption = (*DcOptionPredict)(nil)
)

type DcOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Ipv6         bool     `tl:"ipv6,omitempty:flags:0,implicit"`
	MediaOnly    bool     `tl:"media_only,omitempty:flags:1,implicit"`
	TcpoOnly     bool     `tl:"tcpo_only,omitempty:flags:2,implicit"`
	Cdn          bool     `tl:"cdn,omitempty:flags:3,implicit"`
	Static       bool     `tl:"static,omitempty:flags:4,implicit"`
	ThisPortOnly bool     `tl:"this_port_only,omitempty:flags:5,implicit"`
	ID           int32    `tl:"id"`
	IpAddress    string   `tl:"ip_address"`
	Port         int32    `tl:"port"`
	Secret       *[]byte  `tl:"secret,omitempty:flags:10"`
}

func (*DcOptionPredict) CRC() uint32 {
	return 0x18b7a10d
}
func (*DcOptionPredict) _DcOption() {}

type DefaultHistoryTTL interface {
	tl.TLObject
	_DefaultHistoryTTL()
}

var (
	_ DefaultHistoryTTL = (*DefaultHistoryTTLPredict)(nil)
)

type DefaultHistoryTTLPredict struct {
	Period int32 `tl:"period"`
}

func (*DefaultHistoryTTLPredict) CRC() uint32 {
	return 0x43b46b20
}
func (*DefaultHistoryTTLPredict) _DefaultHistoryTTL() {}

type Dialog interface {
	tl.TLObject
	_Dialog()
}

var (
	_ Dialog = (*DialogPredict)(nil)
	_ Dialog = (*DialogFolderPredict)(nil)
)

type DialogPredict struct {
	_                    struct{}           `tl:"flags,bitflag"`
	Pinned               bool               `tl:"pinned,omitempty:flags:2,implicit"`
	UnreadMark           bool               `tl:"unread_mark,omitempty:flags:3,implicit"`
	ViewForumAsMessages  bool               `tl:"view_forum_as_messages,omitempty:flags:6,implicit"`
	Peer                 Peer               `tl:"peer"`
	TopMessage           int32              `tl:"top_message"`
	ReadInboxMaxID       int32              `tl:"read_inbox_max_id"`
	ReadOutboxMaxID      int32              `tl:"read_outbox_max_id"`
	UnreadCount          int32              `tl:"unread_count"`
	UnreadMentionsCount  int32              `tl:"unread_mentions_count"`
	UnreadReactionsCount int32              `tl:"unread_reactions_count"`
	NotifySettings       PeerNotifySettings `tl:"notify_settings"`
	Pts                  *int32             `tl:"pts,omitempty:flags:0"`
	Draft                DraftMessage       `tl:"draft,omitempty:flags:1"`
	FolderID             *int32             `tl:"folder_id,omitempty:flags:4"`
	TTLPeriod            *int32             `tl:"ttl_period,omitempty:flags:5"`
}

func (*DialogPredict) CRC() uint32 {
	return 0xd58a08c6
}
func (*DialogPredict) _Dialog() {}

type DialogFolderPredict struct {
	_                          struct{} `tl:"flags,bitflag"`
	Pinned                     bool     `tl:"pinned,omitempty:flags:2,implicit"`
	Folder                     Folder   `tl:"folder"`
	Peer                       Peer     `tl:"peer"`
	TopMessage                 int32    `tl:"top_message"`
	UnreadMutedPeersCount      int32    `tl:"unread_muted_peers_count"`
	UnreadUnmutedPeersCount    int32    `tl:"unread_unmuted_peers_count"`
	UnreadMutedMessagesCount   int32    `tl:"unread_muted_messages_count"`
	UnreadUnmutedMessagesCount int32    `tl:"unread_unmuted_messages_count"`
}

func (*DialogFolderPredict) CRC() uint32 {
	return 0x71bd134c
}
func (*DialogFolderPredict) _Dialog() {}

type DialogFilter interface {
	tl.TLObject
	_DialogFilter()
}

var (
	_ DialogFilter = (*DialogFilterPredict)(nil)
	_ DialogFilter = (*DialogFilterDefaultPredict)(nil)
	_ DialogFilter = (*DialogFilterChatlistPredict)(nil)
)

type DialogFilterPredict struct {
	_               struct{}    `tl:"flags,bitflag"`
	Contacts        bool        `tl:"contacts,omitempty:flags:0,implicit"`
	NonContacts     bool        `tl:"non_contacts,omitempty:flags:1,implicit"`
	Groups          bool        `tl:"groups,omitempty:flags:2,implicit"`
	Broadcasts      bool        `tl:"broadcasts,omitempty:flags:3,implicit"`
	Bots            bool        `tl:"bots,omitempty:flags:4,implicit"`
	ExcludeMuted    bool        `tl:"exclude_muted,omitempty:flags:11,implicit"`
	ExcludeRead     bool        `tl:"exclude_read,omitempty:flags:12,implicit"`
	ExcludeArchived bool        `tl:"exclude_archived,omitempty:flags:13,implicit"`
	ID              int32       `tl:"id"`
	Title           string      `tl:"title"`
	Emoticon        *string     `tl:"emoticon,omitempty:flags:25"`
	Color           *int32      `tl:"color,omitempty:flags:27"`
	PinnedPeers     []InputPeer `tl:"pinned_peers"`
	IncludePeers    []InputPeer `tl:"include_peers"`
	ExcludePeers    []InputPeer `tl:"exclude_peers"`
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
	_            struct{}    `tl:"flags,bitflag"`
	HasMyInvites bool        `tl:"has_my_invites,omitempty:flags:26,implicit"`
	ID           int32       `tl:"id"`
	Title        string      `tl:"title"`
	Emoticon     *string     `tl:"emoticon,omitempty:flags:25"`
	Color        *int32      `tl:"color,omitempty:flags:27"`
	PinnedPeers  []InputPeer `tl:"pinned_peers"`
	IncludePeers []InputPeer `tl:"include_peers"`
}

func (*DialogFilterChatlistPredict) CRC() uint32 {
	return 0x9fe28ea4
}
func (*DialogFilterChatlistPredict) _DialogFilter() {}

type DialogFilterSuggested interface {
	tl.TLObject
	_DialogFilterSuggested()
}

var (
	_ DialogFilterSuggested = (*DialogFilterSuggestedPredict)(nil)
)

type DialogFilterSuggestedPredict struct {
	Filter      DialogFilter `tl:"filter"`
	Description string       `tl:"description"`
}

func (*DialogFilterSuggestedPredict) CRC() uint32 {
	return 0x77744d4a
}
func (*DialogFilterSuggestedPredict) _DialogFilterSuggested() {}

type DialogPeer interface {
	tl.TLObject
	_DialogPeer()
}

var (
	_ DialogPeer = (*DialogPeerPredict)(nil)
	_ DialogPeer = (*DialogPeerFolderPredict)(nil)
)

type DialogPeerPredict struct {
	Peer Peer `tl:"peer"`
}

func (*DialogPeerPredict) CRC() uint32 {
	return 0xe56dbf05
}
func (*DialogPeerPredict) _DialogPeer() {}

type DialogPeerFolderPredict struct {
	FolderID int32 `tl:"folder_id"`
}

func (*DialogPeerFolderPredict) CRC() uint32 {
	return 0x514519e2
}
func (*DialogPeerFolderPredict) _DialogPeer() {}

type Document interface {
	tl.TLObject
	_Document()
}

var (
	_ Document = (*DocumentEmptyPredict)(nil)
	_ Document = (*DocumentPredict)(nil)
)

type DocumentEmptyPredict struct {
	ID int64 `tl:"id"`
}

func (*DocumentEmptyPredict) CRC() uint32 {
	return 0x36f8c871
}
func (*DocumentEmptyPredict) _Document() {}

type DocumentPredict struct {
	_             struct{}            `tl:"flags,bitflag"`
	ID            int64               `tl:"id"`
	AccessHash    int64               `tl:"access_hash"`
	FileReference []byte              `tl:"file_reference"`
	Date          int32               `tl:"date"`
	MimeType      string              `tl:"mime_type"`
	Size          int64               `tl:"size"`
	Thumbs        []PhotoSize         `tl:"thumbs,omitempty:flags:0"`
	VideoThumbs   []VideoSize         `tl:"video_thumbs,omitempty:flags:1"`
	DcID          int32               `tl:"dc_id"`
	Attributes    []DocumentAttribute `tl:"attributes"`
}

func (*DocumentPredict) CRC() uint32 {
	return 0x8fd4c4d8
}
func (*DocumentPredict) _Document() {}

type DocumentAttribute interface {
	tl.TLObject
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
	W int32 `tl:"w"`
	H int32 `tl:"h"`
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
	_          struct{}        `tl:"flags,bitflag"`
	Mask       bool            `tl:"mask,omitempty:flags:1,implicit"`
	Alt        string          `tl:"alt"`
	Stickerset InputStickerSet `tl:"stickerset"`
	MaskCoords MaskCoords      `tl:"mask_coords,omitempty:flags:0"`
}

func (*DocumentAttributeStickerPredict) CRC() uint32 {
	return 0x6319d612
}
func (*DocumentAttributeStickerPredict) _DocumentAttribute() {}

type DocumentAttributeVideoPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	RoundMessage      bool     `tl:"round_message,omitempty:flags:0,implicit"`
	SupportsStreaming bool     `tl:"supports_streaming,omitempty:flags:1,implicit"`
	Nosound           bool     `tl:"nosound,omitempty:flags:3,implicit"`
	Duration          float64  `tl:"duration"`
	W                 int32    `tl:"w"`
	H                 int32    `tl:"h"`
	PreloadPrefixSize *int32   `tl:"preload_prefix_size,omitempty:flags:2"`
	VideoStartTs      *float64 `tl:"video_start_ts,omitempty:flags:4"`
}

func (*DocumentAttributeVideoPredict) CRC() uint32 {
	return 0x17399fad
}
func (*DocumentAttributeVideoPredict) _DocumentAttribute() {}

type DocumentAttributeAudioPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Voice     bool     `tl:"voice,omitempty:flags:10,implicit"`
	Duration  int32    `tl:"duration"`
	Title     *string  `tl:"title,omitempty:flags:0"`
	Performer *string  `tl:"performer,omitempty:flags:1"`
	Waveform  *[]byte  `tl:"waveform,omitempty:flags:2"`
}

func (*DocumentAttributeAudioPredict) CRC() uint32 {
	return 0x9852f9c6
}
func (*DocumentAttributeAudioPredict) _DocumentAttribute() {}

type DocumentAttributeFilenamePredict struct {
	FileName string `tl:"file_name"`
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
	_          struct{}        `tl:"flags,bitflag"`
	Free       bool            `tl:"free,omitempty:flags:0,implicit"`
	TextColor  bool            `tl:"text_color,omitempty:flags:1,implicit"`
	Alt        string          `tl:"alt"`
	Stickerset InputStickerSet `tl:"stickerset"`
}

func (*DocumentAttributeCustomEmojiPredict) CRC() uint32 {
	return 0xfd149899
}
func (*DocumentAttributeCustomEmojiPredict) _DocumentAttribute() {}

type DraftMessage interface {
	tl.TLObject
	_DraftMessage()
}

var (
	_ DraftMessage = (*DraftMessageEmptyPredict)(nil)
	_ DraftMessage = (*DraftMessagePredict)(nil)
)

type DraftMessageEmptyPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Date *int32   `tl:"date,omitempty:flags:0"`
}

func (*DraftMessageEmptyPredict) CRC() uint32 {
	return 0x1b0c841a
}
func (*DraftMessageEmptyPredict) _DraftMessage() {}

type DraftMessagePredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	NoWebpage   bool            `tl:"no_webpage,omitempty:flags:1,implicit"`
	InvertMedia bool            `tl:"invert_media,omitempty:flags:6,implicit"`
	ReplyTo     InputReplyTo    `tl:"reply_to,omitempty:flags:4"`
	Message     string          `tl:"message"`
	Entities    []MessageEntity `tl:"entities,omitempty:flags:3"`
	Media       InputMedia      `tl:"media,omitempty:flags:5"`
	Date        int32           `tl:"date"`
	Effect      *int64          `tl:"effect,omitempty:flags:7"`
}

func (*DraftMessagePredict) CRC() uint32 {
	return 0x2d65321f
}
func (*DraftMessagePredict) _DraftMessage() {}

type EmailVerification interface {
	tl.TLObject
	_EmailVerification()
}

var (
	_ EmailVerification = (*EmailVerificationCodePredict)(nil)
	_ EmailVerification = (*EmailVerificationGooglePredict)(nil)
	_ EmailVerification = (*EmailVerificationApplePredict)(nil)
)

type EmailVerificationCodePredict struct {
	Code string `tl:"code"`
}

func (*EmailVerificationCodePredict) CRC() uint32 {
	return 0x922e55a9
}
func (*EmailVerificationCodePredict) _EmailVerification() {}

type EmailVerificationGooglePredict struct {
	Token string `tl:"token"`
}

func (*EmailVerificationGooglePredict) CRC() uint32 {
	return 0xdb909ec2
}
func (*EmailVerificationGooglePredict) _EmailVerification() {}

type EmailVerificationApplePredict struct {
	Token string `tl:"token"`
}

func (*EmailVerificationApplePredict) CRC() uint32 {
	return 0x96d074fd
}
func (*EmailVerificationApplePredict) _EmailVerification() {}

type EmailVerifyPurpose interface {
	tl.TLObject
	_EmailVerifyPurpose()
}

var (
	_ EmailVerifyPurpose = (*EmailVerifyPurposeLoginSetupPredict)(nil)
	_ EmailVerifyPurpose = (*EmailVerifyPurposeLoginChangePredict)(nil)
	_ EmailVerifyPurpose = (*EmailVerifyPurposePassportPredict)(nil)
)

type EmailVerifyPurposeLoginSetupPredict struct {
	PhoneNumber   string `tl:"phone_number"`
	PhoneCodeHash string `tl:"phone_code_hash"`
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
	tl.TLObject
	_EmojiGroup()
}

var (
	_ EmojiGroup = (*EmojiGroupPredict)(nil)
	_ EmojiGroup = (*EmojiGroupGreetingPredict)(nil)
	_ EmojiGroup = (*EmojiGroupPremiumPredict)(nil)
)

type EmojiGroupPredict struct {
	Title       string   `tl:"title"`
	IconEmojiID int64    `tl:"icon_emoji_id"`
	Emoticons   []string `tl:"emoticons"`
}

func (*EmojiGroupPredict) CRC() uint32 {
	return 0x7a9abda9
}
func (*EmojiGroupPredict) _EmojiGroup() {}

type EmojiGroupGreetingPredict struct {
	Title       string   `tl:"title"`
	IconEmojiID int64    `tl:"icon_emoji_id"`
	Emoticons   []string `tl:"emoticons"`
}

func (*EmojiGroupGreetingPredict) CRC() uint32 {
	return 0x80d26cc7
}
func (*EmojiGroupGreetingPredict) _EmojiGroup() {}

type EmojiGroupPremiumPredict struct {
	Title       string `tl:"title"`
	IconEmojiID int64  `tl:"icon_emoji_id"`
}

func (*EmojiGroupPremiumPredict) CRC() uint32 {
	return 0x093bcf34
}
func (*EmojiGroupPremiumPredict) _EmojiGroup() {}

type EmojiKeyword interface {
	tl.TLObject
	_EmojiKeyword()
}

var (
	_ EmojiKeyword = (*EmojiKeywordPredict)(nil)
	_ EmojiKeyword = (*EmojiKeywordDeletedPredict)(nil)
)

type EmojiKeywordPredict struct {
	Keyword   string   `tl:"keyword"`
	Emoticons []string `tl:"emoticons"`
}

func (*EmojiKeywordPredict) CRC() uint32 {
	return 0xd5b3b9f9
}
func (*EmojiKeywordPredict) _EmojiKeyword() {}

type EmojiKeywordDeletedPredict struct {
	Keyword   string   `tl:"keyword"`
	Emoticons []string `tl:"emoticons"`
}

func (*EmojiKeywordDeletedPredict) CRC() uint32 {
	return 0x236df622
}
func (*EmojiKeywordDeletedPredict) _EmojiKeyword() {}

type EmojiKeywordsDifference interface {
	tl.TLObject
	_EmojiKeywordsDifference()
}

var (
	_ EmojiKeywordsDifference = (*EmojiKeywordsDifferencePredict)(nil)
)

type EmojiKeywordsDifferencePredict struct {
	LangCode    string         `tl:"lang_code"`
	FromVersion int32          `tl:"from_version"`
	Version     int32          `tl:"version"`
	Keywords    []EmojiKeyword `tl:"keywords"`
}

func (*EmojiKeywordsDifferencePredict) CRC() uint32 {
	return 0x5cc761bd
}
func (*EmojiKeywordsDifferencePredict) _EmojiKeywordsDifference() {}

type EmojiLanguage interface {
	tl.TLObject
	_EmojiLanguage()
}

var (
	_ EmojiLanguage = (*EmojiLanguagePredict)(nil)
)

type EmojiLanguagePredict struct {
	LangCode string `tl:"lang_code"`
}

func (*EmojiLanguagePredict) CRC() uint32 {
	return 0xb3fb5361
}
func (*EmojiLanguagePredict) _EmojiLanguage() {}

type EmojiList interface {
	tl.TLObject
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
	Hash       int64   `tl:"hash"`
	DocumentID []int64 `tl:"document_id"`
}

func (*EmojiListPredict) CRC() uint32 {
	return 0x7a1e11d1
}
func (*EmojiListPredict) _EmojiList() {}

type EmojiStatus interface {
	tl.TLObject
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
	DocumentID int64 `tl:"document_id"`
}

func (*EmojiStatusPredict) CRC() uint32 {
	return 0x929b619d
}
func (*EmojiStatusPredict) _EmojiStatus() {}

type EmojiStatusUntilPredict struct {
	DocumentID int64 `tl:"document_id"`
	Until      int32 `tl:"until"`
}

func (*EmojiStatusUntilPredict) CRC() uint32 {
	return 0xfa30a8c7
}
func (*EmojiStatusUntilPredict) _EmojiStatus() {}

type EmojiURL interface {
	tl.TLObject
	_EmojiURL()
}

var (
	_ EmojiURL = (*EmojiURLPredict)(nil)
)

type EmojiURLPredict struct {
	URL string `tl:"url"`
}

func (*EmojiURLPredict) CRC() uint32 {
	return 0xa575739d
}
func (*EmojiURLPredict) _EmojiURL() {}

type EncryptedChat interface {
	tl.TLObject
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
	ID int32 `tl:"id"`
}

func (*EncryptedChatEmptyPredict) CRC() uint32 {
	return 0xab7ec0a0
}
func (*EncryptedChatEmptyPredict) _EncryptedChat() {}

type EncryptedChatWaitingPredict struct {
	ID            int32 `tl:"id"`
	AccessHash    int64 `tl:"access_hash"`
	Date          int32 `tl:"date"`
	AdminID       int64 `tl:"admin_id"`
	ParticipantID int64 `tl:"participant_id"`
}

func (*EncryptedChatWaitingPredict) CRC() uint32 {
	return 0x66b25953
}
func (*EncryptedChatWaitingPredict) _EncryptedChat() {}

type EncryptedChatRequestedPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	FolderID      *int32   `tl:"folder_id,omitempty:flags:0"`
	ID            int32    `tl:"id"`
	AccessHash    int64    `tl:"access_hash"`
	Date          int32    `tl:"date"`
	AdminID       int64    `tl:"admin_id"`
	ParticipantID int64    `tl:"participant_id"`
	GA            []byte   `tl:"g_a"`
}

func (*EncryptedChatRequestedPredict) CRC() uint32 {
	return 0x48f1d94c
}
func (*EncryptedChatRequestedPredict) _EncryptedChat() {}

type EncryptedChatPredict struct {
	ID             int32  `tl:"id"`
	AccessHash     int64  `tl:"access_hash"`
	Date           int32  `tl:"date"`
	AdminID        int64  `tl:"admin_id"`
	ParticipantID  int64  `tl:"participant_id"`
	GAOrB          []byte `tl:"g_a_or_b"`
	KeyFingerprint int64  `tl:"key_fingerprint"`
}

func (*EncryptedChatPredict) CRC() uint32 {
	return 0x61f0d4c7
}
func (*EncryptedChatPredict) _EncryptedChat() {}

type EncryptedChatDiscardedPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	HistoryDeleted bool     `tl:"history_deleted,omitempty:flags:0,implicit"`
	ID             int32    `tl:"id"`
}

func (*EncryptedChatDiscardedPredict) CRC() uint32 {
	return 0x1e1c7c45
}
func (*EncryptedChatDiscardedPredict) _EncryptedChat() {}

type EncryptedFile interface {
	tl.TLObject
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
	ID             int64 `tl:"id"`
	AccessHash     int64 `tl:"access_hash"`
	Size           int64 `tl:"size"`
	DcID           int32 `tl:"dc_id"`
	KeyFingerprint int32 `tl:"key_fingerprint"`
}

func (*EncryptedFilePredict) CRC() uint32 {
	return 0xa8008cd8
}
func (*EncryptedFilePredict) _EncryptedFile() {}

type EncryptedMessage interface {
	tl.TLObject
	_EncryptedMessage()
}

var (
	_ EncryptedMessage = (*EncryptedMessagePredict)(nil)
	_ EncryptedMessage = (*EncryptedMessageServicePredict)(nil)
)

type EncryptedMessagePredict struct {
	RandomID int64         `tl:"random_id"`
	ChatID   int32         `tl:"chat_id"`
	Date     int32         `tl:"date"`
	Bytes    []byte        `tl:"bytes"`
	File     EncryptedFile `tl:"file"`
}

func (*EncryptedMessagePredict) CRC() uint32 {
	return 0xed18c118
}
func (*EncryptedMessagePredict) _EncryptedMessage() {}

type EncryptedMessageServicePredict struct {
	RandomID int64  `tl:"random_id"`
	ChatID   int32  `tl:"chat_id"`
	Date     int32  `tl:"date"`
	Bytes    []byte `tl:"bytes"`
}

func (*EncryptedMessageServicePredict) CRC() uint32 {
	return 0x23734b06
}
func (*EncryptedMessageServicePredict) _EncryptedMessage() {}

type Error interface {
	tl.TLObject
	_Error()
}

var (
	_ Error = (*ErrorPredict)(nil)
)

type ErrorPredict struct {
	Code int32  `tl:"code"`
	Text string `tl:"text"`
}

func (*ErrorPredict) CRC() uint32 {
	return 0xc4b9f9bb
}
func (*ErrorPredict) _Error() {}

type ExportedChatInvite interface {
	tl.TLObject
	_ExportedChatInvite()
}

var (
	_ ExportedChatInvite = (*ChatInviteExportedPredict)(nil)
	_ ExportedChatInvite = (*ChatInvitePublicJoinRequestsPredict)(nil)
)

type ChatInviteExportedPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Revoked       bool     `tl:"revoked,omitempty:flags:0,implicit"`
	Permanent     bool     `tl:"permanent,omitempty:flags:5,implicit"`
	RequestNeeded bool     `tl:"request_needed,omitempty:flags:6,implicit"`
	Link          string   `tl:"link"`
	AdminID       int64    `tl:"admin_id"`
	Date          int32    `tl:"date"`
	StartDate     *int32   `tl:"start_date,omitempty:flags:4"`
	ExpireDate    *int32   `tl:"expire_date,omitempty:flags:1"`
	UsageLimit    *int32   `tl:"usage_limit,omitempty:flags:2"`
	Usage         *int32   `tl:"usage,omitempty:flags:3"`
	Requested     *int32   `tl:"requested,omitempty:flags:7"`
	Title         *string  `tl:"title,omitempty:flags:8"`
}

func (*ChatInviteExportedPredict) CRC() uint32 {
	return 0x0ab4a819
}
func (*ChatInviteExportedPredict) _ExportedChatInvite() {}

type ChatInvitePublicJoinRequestsPredict struct{}

func (*ChatInvitePublicJoinRequestsPredict) CRC() uint32 {
	return 0xed107ab7
}
func (*ChatInvitePublicJoinRequestsPredict) _ExportedChatInvite() {}

type ExportedChatlistInvite interface {
	tl.TLObject
	_ExportedChatlistInvite()
}

var (
	_ ExportedChatlistInvite = (*ExportedChatlistInvitePredict)(nil)
)

type ExportedChatlistInvitePredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Title string   `tl:"title"`
	URL   string   `tl:"url"`
	Peers []Peer   `tl:"peers"`
}

func (*ExportedChatlistInvitePredict) CRC() uint32 {
	return 0x0c5181ac
}
func (*ExportedChatlistInvitePredict) _ExportedChatlistInvite() {}

type ExportedContactToken interface {
	tl.TLObject
	_ExportedContactToken()
}

var (
	_ ExportedContactToken = (*ExportedContactTokenPredict)(nil)
)

type ExportedContactTokenPredict struct {
	URL     string `tl:"url"`
	Expires int32  `tl:"expires"`
}

func (*ExportedContactTokenPredict) CRC() uint32 {
	return 0x41bf109b
}
func (*ExportedContactTokenPredict) _ExportedContactToken() {}

type ExportedMessageLink interface {
	tl.TLObject
	_ExportedMessageLink()
}

var (
	_ ExportedMessageLink = (*ExportedMessageLinkPredict)(nil)
)

type ExportedMessageLinkPredict struct {
	Link string `tl:"link"`
	Html string `tl:"html"`
}

func (*ExportedMessageLinkPredict) CRC() uint32 {
	return 0x5dab1af4
}
func (*ExportedMessageLinkPredict) _ExportedMessageLink() {}

type ExportedStoryLink interface {
	tl.TLObject
	_ExportedStoryLink()
}

var (
	_ ExportedStoryLink = (*ExportedStoryLinkPredict)(nil)
)

type ExportedStoryLinkPredict struct {
	Link string `tl:"link"`
}

func (*ExportedStoryLinkPredict) CRC() uint32 {
	return 0x3fc9053b
}
func (*ExportedStoryLinkPredict) _ExportedStoryLink() {}

type FactCheck interface {
	tl.TLObject
	_FactCheck()
}

var (
	_ FactCheck = (*FactCheckPredict)(nil)
)

type FactCheckPredict struct {
	_         struct{}         `tl:"flags,bitflag"`
	NeedCheck bool             `tl:"need_check,omitempty:flags:0,implicit"`
	Country   *string          `tl:"country,omitempty:flags:1"`
	Text      TextWithEntities `tl:"text,omitempty:flags:1"`
	Hash      int64            `tl:"hash"`
}

func (*FactCheckPredict) CRC() uint32 {
	return 0xb89bfccf
}
func (*FactCheckPredict) _FactCheck() {}

type FileHash interface {
	tl.TLObject
	_FileHash()
}

var (
	_ FileHash = (*FileHashPredict)(nil)
)

type FileHashPredict struct {
	Offset int64  `tl:"offset"`
	Limit  int32  `tl:"limit"`
	Hash   []byte `tl:"hash"`
}

func (*FileHashPredict) CRC() uint32 {
	return 0xf39b035c
}
func (*FileHashPredict) _FileHash() {}

type Folder interface {
	tl.TLObject
	_Folder()
}

var (
	_ Folder = (*FolderPredict)(nil)
)

type FolderPredict struct {
	_                         struct{}  `tl:"flags,bitflag"`
	AutofillNewBroadcasts     bool      `tl:"autofill_new_broadcasts,omitempty:flags:0,implicit"`
	AutofillPublicGroups      bool      `tl:"autofill_public_groups,omitempty:flags:1,implicit"`
	AutofillNewCorrespondents bool      `tl:"autofill_new_correspondents,omitempty:flags:2,implicit"`
	ID                        int32     `tl:"id"`
	Title                     string    `tl:"title"`
	Photo                     ChatPhoto `tl:"photo,omitempty:flags:3"`
}

func (*FolderPredict) CRC() uint32 {
	return 0xff544e65
}
func (*FolderPredict) _Folder() {}

type FolderPeer interface {
	tl.TLObject
	_FolderPeer()
}

var (
	_ FolderPeer = (*FolderPeerPredict)(nil)
)

type FolderPeerPredict struct {
	Peer     Peer  `tl:"peer"`
	FolderID int32 `tl:"folder_id"`
}

func (*FolderPeerPredict) CRC() uint32 {
	return 0xe9baa668
}
func (*FolderPeerPredict) _FolderPeer() {}

type ForumTopic interface {
	tl.TLObject
	_ForumTopic()
}

var (
	_ ForumTopic = (*ForumTopicDeletedPredict)(nil)
	_ ForumTopic = (*ForumTopicPredict)(nil)
)

type ForumTopicDeletedPredict struct {
	ID int32 `tl:"id"`
}

func (*ForumTopicDeletedPredict) CRC() uint32 {
	return 0x023f109b
}
func (*ForumTopicDeletedPredict) _ForumTopic() {}

type ForumTopicPredict struct {
	_                    struct{}           `tl:"flags,bitflag"`
	My                   bool               `tl:"my,omitempty:flags:1,implicit"`
	Closed               bool               `tl:"closed,omitempty:flags:2,implicit"`
	Pinned               bool               `tl:"pinned,omitempty:flags:3,implicit"`
	Short                bool               `tl:"short,omitempty:flags:5,implicit"`
	Hidden               bool               `tl:"hidden,omitempty:flags:6,implicit"`
	ID                   int32              `tl:"id"`
	Date                 int32              `tl:"date"`
	Title                string             `tl:"title"`
	IconColor            int32              `tl:"icon_color"`
	IconEmojiID          *int64             `tl:"icon_emoji_id,omitempty:flags:0"`
	TopMessage           int32              `tl:"top_message"`
	ReadInboxMaxID       int32              `tl:"read_inbox_max_id"`
	ReadOutboxMaxID      int32              `tl:"read_outbox_max_id"`
	UnreadCount          int32              `tl:"unread_count"`
	UnreadMentionsCount  int32              `tl:"unread_mentions_count"`
	UnreadReactionsCount int32              `tl:"unread_reactions_count"`
	FromID               Peer               `tl:"from_id"`
	NotifySettings       PeerNotifySettings `tl:"notify_settings"`
	Draft                DraftMessage       `tl:"draft,omitempty:flags:4"`
}

func (*ForumTopicPredict) CRC() uint32 {
	return 0x71701da9
}
func (*ForumTopicPredict) _ForumTopic() {}

type FoundStory interface {
	tl.TLObject
	_FoundStory()
}

var (
	_ FoundStory = (*FoundStoryPredict)(nil)
)

type FoundStoryPredict struct {
	Peer  Peer      `tl:"peer"`
	Story StoryItem `tl:"story"`
}

func (*FoundStoryPredict) CRC() uint32 {
	return 0xe87acbc0
}
func (*FoundStoryPredict) _FoundStory() {}

type Game interface {
	tl.TLObject
	_Game()
}

var (
	_ Game = (*GamePredict)(nil)
)

type GamePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ID          int64    `tl:"id"`
	AccessHash  int64    `tl:"access_hash"`
	ShortName   string   `tl:"short_name"`
	Title       string   `tl:"title"`
	Description string   `tl:"description"`
	Photo       Photo    `tl:"photo"`
	Document    Document `tl:"document,omitempty:flags:0"`
}

func (*GamePredict) CRC() uint32 {
	return 0xbdf9653b
}
func (*GamePredict) _Game() {}

type GeoPoint interface {
	tl.TLObject
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
	Long           float64  `tl:"long"`
	Lat            float64  `tl:"lat"`
	AccessHash     int64    `tl:"access_hash"`
	AccuracyRadius *int32   `tl:"accuracy_radius,omitempty:flags:0"`
}

func (*GeoPointPredict) CRC() uint32 {
	return 0xb2a2f663
}
func (*GeoPointPredict) _GeoPoint() {}

type GeoPointAddress interface {
	tl.TLObject
	_GeoPointAddress()
}

var (
	_ GeoPointAddress = (*GeoPointAddressPredict)(nil)
)

type GeoPointAddressPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	CountryIso2 string   `tl:"country_iso2"`
	State       *string  `tl:"state,omitempty:flags:0"`
	City        *string  `tl:"city,omitempty:flags:1"`
	Street      *string  `tl:"street,omitempty:flags:2"`
}

func (*GeoPointAddressPredict) CRC() uint32 {
	return 0xde4c5d93
}
func (*GeoPointAddressPredict) _GeoPointAddress() {}

type GlobalPrivacySettings interface {
	tl.TLObject
	_GlobalPrivacySettings()
}

var (
	_ GlobalPrivacySettings = (*GlobalPrivacySettingsPredict)(nil)
)

type GlobalPrivacySettingsPredict struct {
	_                                struct{} `tl:"flags,bitflag"`
	ArchiveAndMuteNewNoncontactPeers bool     `tl:"archive_and_mute_new_noncontact_peers,omitempty:flags:0,implicit"`
	KeepArchivedUnmuted              bool     `tl:"keep_archived_unmuted,omitempty:flags:1,implicit"`
	KeepArchivedFolders              bool     `tl:"keep_archived_folders,omitempty:flags:2,implicit"`
	HideReadMarks                    bool     `tl:"hide_read_marks,omitempty:flags:3,implicit"`
	NewNoncontactPeersRequirePremium bool     `tl:"new_noncontact_peers_require_premium,omitempty:flags:4,implicit"`
}

func (*GlobalPrivacySettingsPredict) CRC() uint32 {
	return 0x734c4ccb
}
func (*GlobalPrivacySettingsPredict) _GlobalPrivacySettings() {}

type GroupCall interface {
	tl.TLObject
	_GroupCall()
}

var (
	_ GroupCall = (*GroupCallDiscardedPredict)(nil)
	_ GroupCall = (*GroupCallPredict)(nil)
)

type GroupCallDiscardedPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
	Duration   int32 `tl:"duration"`
}

func (*GroupCallDiscardedPredict) CRC() uint32 {
	return 0x7780bcb4
}
func (*GroupCallDiscardedPredict) _GroupCall() {}

type GroupCallPredict struct {
	_                       struct{} `tl:"flags,bitflag"`
	JoinMuted               bool     `tl:"join_muted,omitempty:flags:1,implicit"`
	CanChangeJoinMuted      bool     `tl:"can_change_join_muted,omitempty:flags:2,implicit"`
	JoinDateAsc             bool     `tl:"join_date_asc,omitempty:flags:6,implicit"`
	ScheduleStartSubscribed bool     `tl:"schedule_start_subscribed,omitempty:flags:8,implicit"`
	CanStartVideo           bool     `tl:"can_start_video,omitempty:flags:9,implicit"`
	RecordVideoActive       bool     `tl:"record_video_active,omitempty:flags:11,implicit"`
	RtmpStream              bool     `tl:"rtmp_stream,omitempty:flags:12,implicit"`
	ListenersHidden         bool     `tl:"listeners_hidden,omitempty:flags:13,implicit"`
	ID                      int64    `tl:"id"`
	AccessHash              int64    `tl:"access_hash"`
	ParticipantsCount       int32    `tl:"participants_count"`
	Title                   *string  `tl:"title,omitempty:flags:3"`
	StreamDcID              *int32   `tl:"stream_dc_id,omitempty:flags:4"`
	RecordStartDate         *int32   `tl:"record_start_date,omitempty:flags:5"`
	ScheduleDate            *int32   `tl:"schedule_date,omitempty:flags:7"`
	UnmutedVideoCount       *int32   `tl:"unmuted_video_count,omitempty:flags:10"`
	UnmutedVideoLimit       int32    `tl:"unmuted_video_limit"`
	Version                 int32    `tl:"version"`
}

func (*GroupCallPredict) CRC() uint32 {
	return 0xd597650c
}
func (*GroupCallPredict) _GroupCall() {}

type GroupCallParticipant interface {
	tl.TLObject
	_GroupCallParticipant()
}

var (
	_ GroupCallParticipant = (*GroupCallParticipantPredict)(nil)
)

type GroupCallParticipantPredict struct {
	_               struct{}                  `tl:"flags,bitflag"`
	Muted           bool                      `tl:"muted,omitempty:flags:0,implicit"`
	Left            bool                      `tl:"left,omitempty:flags:1,implicit"`
	CanSelfUnmute   bool                      `tl:"can_self_unmute,omitempty:flags:2,implicit"`
	JustJoined      bool                      `tl:"just_joined,omitempty:flags:4,implicit"`
	Versioned       bool                      `tl:"versioned,omitempty:flags:5,implicit"`
	Min             bool                      `tl:"min,omitempty:flags:8,implicit"`
	MutedByYou      bool                      `tl:"muted_by_you,omitempty:flags:9,implicit"`
	VolumeByAdmin   bool                      `tl:"volume_by_admin,omitempty:flags:10,implicit"`
	Self            bool                      `tl:"self,omitempty:flags:12,implicit"`
	VideoJoined     bool                      `tl:"video_joined,omitempty:flags:15,implicit"`
	Peer            Peer                      `tl:"peer"`
	Date            int32                     `tl:"date"`
	ActiveDate      *int32                    `tl:"active_date,omitempty:flags:3"`
	Source          int32                     `tl:"source"`
	Volume          *int32                    `tl:"volume,omitempty:flags:7"`
	About           *string                   `tl:"about,omitempty:flags:11"`
	RaiseHandRating *int64                    `tl:"raise_hand_rating,omitempty:flags:13"`
	Video           GroupCallParticipantVideo `tl:"video,omitempty:flags:6"`
	Presentation    GroupCallParticipantVideo `tl:"presentation,omitempty:flags:14"`
}

func (*GroupCallParticipantPredict) CRC() uint32 {
	return 0xeba636fe
}
func (*GroupCallParticipantPredict) _GroupCallParticipant() {}

type GroupCallParticipantVideo interface {
	tl.TLObject
	_GroupCallParticipantVideo()
}

var (
	_ GroupCallParticipantVideo = (*GroupCallParticipantVideoPredict)(nil)
)

type GroupCallParticipantVideoPredict struct {
	_            struct{}                               `tl:"flags,bitflag"`
	Paused       bool                                   `tl:"paused,omitempty:flags:0,implicit"`
	Endpoint     string                                 `tl:"endpoint"`
	SourceGroups []GroupCallParticipantVideoSourceGroup `tl:"source_groups"`
	AudioSource  *int32                                 `tl:"audio_source,omitempty:flags:1"`
}

func (*GroupCallParticipantVideoPredict) CRC() uint32 {
	return 0x67753ac8
}
func (*GroupCallParticipantVideoPredict) _GroupCallParticipantVideo() {}

type GroupCallParticipantVideoSourceGroup interface {
	tl.TLObject
	_GroupCallParticipantVideoSourceGroup()
}

var (
	_ GroupCallParticipantVideoSourceGroup = (*GroupCallParticipantVideoSourceGroupPredict)(nil)
)

type GroupCallParticipantVideoSourceGroupPredict struct {
	Semantics string  `tl:"semantics"`
	Sources   []int32 `tl:"sources"`
}

func (*GroupCallParticipantVideoSourceGroupPredict) CRC() uint32 {
	return 0xdcb118b7
}
func (*GroupCallParticipantVideoSourceGroupPredict) _GroupCallParticipantVideoSourceGroup() {}

type GroupCallStreamChannel interface {
	tl.TLObject
	_GroupCallStreamChannel()
}

var (
	_ GroupCallStreamChannel = (*GroupCallStreamChannelPredict)(nil)
)

type GroupCallStreamChannelPredict struct {
	Channel         int32 `tl:"channel"`
	Scale           int32 `tl:"scale"`
	LastTimestampMs int64 `tl:"last_timestamp_ms"`
}

func (*GroupCallStreamChannelPredict) CRC() uint32 {
	return 0x80eb48af
}
func (*GroupCallStreamChannelPredict) _GroupCallStreamChannel() {}

type HighScore interface {
	tl.TLObject
	_HighScore()
}

var (
	_ HighScore = (*HighScorePredict)(nil)
)

type HighScorePredict struct {
	Pos    int32 `tl:"pos"`
	UserID int64 `tl:"user_id"`
	Score  int32 `tl:"score"`
}

func (*HighScorePredict) CRC() uint32 {
	return 0x73a379eb
}
func (*HighScorePredict) _HighScore() {}

type ImportedContact interface {
	tl.TLObject
	_ImportedContact()
}

var (
	_ ImportedContact = (*ImportedContactPredict)(nil)
)

type ImportedContactPredict struct {
	UserID   int64 `tl:"user_id"`
	ClientID int64 `tl:"client_id"`
}

func (*ImportedContactPredict) CRC() uint32 {
	return 0xc13e3c50
}
func (*ImportedContactPredict) _ImportedContact() {}

type InlineBotSwitchPm interface {
	tl.TLObject
	_InlineBotSwitchPm()
}

var (
	_ InlineBotSwitchPm = (*InlineBotSwitchPmPredict)(nil)
)

type InlineBotSwitchPmPredict struct {
	Text       string `tl:"text"`
	StartParam string `tl:"start_param"`
}

func (*InlineBotSwitchPmPredict) CRC() uint32 {
	return 0x3c20629f
}
func (*InlineBotSwitchPmPredict) _InlineBotSwitchPm() {}

type InlineBotWebView interface {
	tl.TLObject
	_InlineBotWebView()
}

var (
	_ InlineBotWebView = (*InlineBotWebViewPredict)(nil)
)

type InlineBotWebViewPredict struct {
	Text string `tl:"text"`
	URL  string `tl:"url"`
}

func (*InlineBotWebViewPredict) CRC() uint32 {
	return 0xb57295d5
}
func (*InlineBotWebViewPredict) _InlineBotWebView() {}

type InlineQueryPeerType interface {
	tl.TLObject
	_InlineQueryPeerType()
}

var (
	_ InlineQueryPeerType = (*InlineQueryPeerTypeSameBotPmPredict)(nil)
	_ InlineQueryPeerType = (*InlineQueryPeerTypePmPredict)(nil)
	_ InlineQueryPeerType = (*InlineQueryPeerTypeChatPredict)(nil)
	_ InlineQueryPeerType = (*InlineQueryPeerTypeMegagroupPredict)(nil)
	_ InlineQueryPeerType = (*InlineQueryPeerTypeBroadcastPredict)(nil)
	_ InlineQueryPeerType = (*InlineQueryPeerTypeBotPmPredict)(nil)
)

type InlineQueryPeerTypeSameBotPmPredict struct{}

func (*InlineQueryPeerTypeSameBotPmPredict) CRC() uint32 {
	return 0x3081ed9d
}
func (*InlineQueryPeerTypeSameBotPmPredict) _InlineQueryPeerType() {}

type InlineQueryPeerTypePmPredict struct{}

func (*InlineQueryPeerTypePmPredict) CRC() uint32 {
	return 0x833c0fac
}
func (*InlineQueryPeerTypePmPredict) _InlineQueryPeerType() {}

type InlineQueryPeerTypeChatPredict struct{}

func (*InlineQueryPeerTypeChatPredict) CRC() uint32 {
	return 0xd766c50a
}
func (*InlineQueryPeerTypeChatPredict) _InlineQueryPeerType() {}

type InlineQueryPeerTypeMegagroupPredict struct{}

func (*InlineQueryPeerTypeMegagroupPredict) CRC() uint32 {
	return 0x5ec4be43
}
func (*InlineQueryPeerTypeMegagroupPredict) _InlineQueryPeerType() {}

type InlineQueryPeerTypeBroadcastPredict struct{}

func (*InlineQueryPeerTypeBroadcastPredict) CRC() uint32 {
	return 0x6334ee9a
}
func (*InlineQueryPeerTypeBroadcastPredict) _InlineQueryPeerType() {}

type InlineQueryPeerTypeBotPmPredict struct{}

func (*InlineQueryPeerTypeBotPmPredict) CRC() uint32 {
	return 0x0e3b2d0c
}
func (*InlineQueryPeerTypeBotPmPredict) _InlineQueryPeerType() {}

type InputAppEvent interface {
	tl.TLObject
	_InputAppEvent()
}

var (
	_ InputAppEvent = (*InputAppEventPredict)(nil)
)

type InputAppEventPredict struct {
	Time float64   `tl:"time"`
	Type string    `tl:"type"`
	Peer int64     `tl:"peer"`
	Data JSONValue `tl:"data"`
}

func (*InputAppEventPredict) CRC() uint32 {
	return 0x1d1b1245
}
func (*InputAppEventPredict) _InputAppEvent() {}

type InputBotApp interface {
	tl.TLObject
	_InputBotApp()
}

var (
	_ InputBotApp = (*InputBotAppIDPredict)(nil)
	_ InputBotApp = (*InputBotAppShortNamePredict)(nil)
)

type InputBotAppIDPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputBotAppIDPredict) CRC() uint32 {
	return 0xa920bd7a
}
func (*InputBotAppIDPredict) _InputBotApp() {}

type InputBotAppShortNamePredict struct {
	BotID     InputUser `tl:"bot_id"`
	ShortName string    `tl:"short_name"`
}

func (*InputBotAppShortNamePredict) CRC() uint32 {
	return 0x908c0407
}
func (*InputBotAppShortNamePredict) _InputBotApp() {}

type InputBotInlineMessage interface {
	tl.TLObject
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
	_           struct{}        `tl:"flags,bitflag"`
	InvertMedia bool            `tl:"invert_media,omitempty:flags:3,implicit"`
	Message     string          `tl:"message"`
	Entities    []MessageEntity `tl:"entities,omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaAutoPredict) CRC() uint32 {
	return 0x3380c786
}
func (*InputBotInlineMessageMediaAutoPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageTextPredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	NoWebpage   bool            `tl:"no_webpage,omitempty:flags:0,implicit"`
	InvertMedia bool            `tl:"invert_media,omitempty:flags:3,implicit"`
	Message     string          `tl:"message"`
	Entities    []MessageEntity `tl:"entities,omitempty:flags:1"`
	ReplyMarkup ReplyMarkup     `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageTextPredict) CRC() uint32 {
	return 0x3dcd7a87
}
func (*InputBotInlineMessageTextPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaGeoPredict struct {
	_                           struct{}      `tl:"flags,bitflag"`
	GeoPoint                    InputGeoPoint `tl:"geo_point"`
	Heading                     *int32        `tl:"heading,omitempty:flags:0"`
	Period                      *int32        `tl:"period,omitempty:flags:1"`
	ProximityNotificationRadius *int32        `tl:"proximity_notification_radius,omitempty:flags:3"`
	ReplyMarkup                 ReplyMarkup   `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaGeoPredict) CRC() uint32 {
	return 0x96929a85
}
func (*InputBotInlineMessageMediaGeoPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaVenuePredict struct {
	_           struct{}      `tl:"flags,bitflag"`
	GeoPoint    InputGeoPoint `tl:"geo_point"`
	Title       string        `tl:"title"`
	Address     string        `tl:"address"`
	Provider    string        `tl:"provider"`
	VenueID     string        `tl:"venue_id"`
	VenueType   string        `tl:"venue_type"`
	ReplyMarkup ReplyMarkup   `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaVenuePredict) CRC() uint32 {
	return 0x417bbf11
}
func (*InputBotInlineMessageMediaVenuePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaContactPredict struct {
	_           struct{}    `tl:"flags,bitflag"`
	PhoneNumber string      `tl:"phone_number"`
	FirstName   string      `tl:"first_name"`
	LastName    string      `tl:"last_name"`
	Vcard       string      `tl:"vcard"`
	ReplyMarkup ReplyMarkup `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaContactPredict) CRC() uint32 {
	return 0xa6edbffd
}
func (*InputBotInlineMessageMediaContactPredict) _InputBotInlineMessage() {}

type InputBotInlineMessageGamePredict struct {
	_           struct{}    `tl:"flags,bitflag"`
	ReplyMarkup ReplyMarkup `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageGamePredict) CRC() uint32 {
	return 0x4b425864
}
func (*InputBotInlineMessageGamePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaInvoicePredict struct {
	_            struct{}         `tl:"flags,bitflag"`
	Title        string           `tl:"title"`
	Description  string           `tl:"description"`
	Photo        InputWebDocument `tl:"photo,omitempty:flags:0"`
	Invoice      Invoice          `tl:"invoice"`
	Payload      []byte           `tl:"payload"`
	Provider     string           `tl:"provider"`
	ProviderData DataJSON         `tl:"provider_data"`
	ReplyMarkup  ReplyMarkup      `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaInvoicePredict) CRC() uint32 {
	return 0xd7e78225
}
func (*InputBotInlineMessageMediaInvoicePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageMediaWebPagePredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	InvertMedia     bool            `tl:"invert_media,omitempty:flags:3,implicit"`
	ForceLargeMedia bool            `tl:"force_large_media,omitempty:flags:4,implicit"`
	ForceSmallMedia bool            `tl:"force_small_media,omitempty:flags:5,implicit"`
	Optional        bool            `tl:"optional,omitempty:flags:6,implicit"`
	Message         string          `tl:"message"`
	Entities        []MessageEntity `tl:"entities,omitempty:flags:1"`
	URL             string          `tl:"url"`
	ReplyMarkup     ReplyMarkup     `tl:"reply_markup,omitempty:flags:2"`
}

func (*InputBotInlineMessageMediaWebPagePredict) CRC() uint32 {
	return 0xbddcc510
}
func (*InputBotInlineMessageMediaWebPagePredict) _InputBotInlineMessage() {}

type InputBotInlineMessageID interface {
	tl.TLObject
	_InputBotInlineMessageID()
}

var (
	_ InputBotInlineMessageID = (*InputBotInlineMessageIDPredict)(nil)
	_ InputBotInlineMessageID = (*InputBotInlineMessageID64Predict)(nil)
)

type InputBotInlineMessageIDPredict struct {
	DcID       int32 `tl:"dc_id"`
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputBotInlineMessageIDPredict) CRC() uint32 {
	return 0x890c3d89
}
func (*InputBotInlineMessageIDPredict) _InputBotInlineMessageID() {}

type InputBotInlineMessageID64Predict struct {
	DcID       int32 `tl:"dc_id"`
	OwnerID    int64 `tl:"owner_id"`
	ID         int32 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputBotInlineMessageID64Predict) CRC() uint32 {
	return 0xb6d915d7
}
func (*InputBotInlineMessageID64Predict) _InputBotInlineMessageID() {}

type InputBotInlineResult interface {
	tl.TLObject
	_InputBotInlineResult()
}

var (
	_ InputBotInlineResult = (*InputBotInlineResultPredict)(nil)
	_ InputBotInlineResult = (*InputBotInlineResultPhotoPredict)(nil)
	_ InputBotInlineResult = (*InputBotInlineResultDocumentPredict)(nil)
	_ InputBotInlineResult = (*InputBotInlineResultGamePredict)(nil)
)

type InputBotInlineResultPredict struct {
	_           struct{}              `tl:"flags,bitflag"`
	ID          string                `tl:"id"`
	Type        string                `tl:"type"`
	Title       *string               `tl:"title,omitempty:flags:1"`
	Description *string               `tl:"description,omitempty:flags:2"`
	URL         *string               `tl:"url,omitempty:flags:3"`
	Thumb       InputWebDocument      `tl:"thumb,omitempty:flags:4"`
	Content     InputWebDocument      `tl:"content,omitempty:flags:5"`
	SendMessage InputBotInlineMessage `tl:"send_message"`
}

func (*InputBotInlineResultPredict) CRC() uint32 {
	return 0x88bf9319
}
func (*InputBotInlineResultPredict) _InputBotInlineResult() {}

type InputBotInlineResultPhotoPredict struct {
	ID          string                `tl:"id"`
	Type        string                `tl:"type"`
	Photo       InputPhoto            `tl:"photo"`
	SendMessage InputBotInlineMessage `tl:"send_message"`
}

func (*InputBotInlineResultPhotoPredict) CRC() uint32 {
	return 0xa8d864a7
}
func (*InputBotInlineResultPhotoPredict) _InputBotInlineResult() {}

type InputBotInlineResultDocumentPredict struct {
	_           struct{}              `tl:"flags,bitflag"`
	ID          string                `tl:"id"`
	Type        string                `tl:"type"`
	Title       *string               `tl:"title,omitempty:flags:1"`
	Description *string               `tl:"description,omitempty:flags:2"`
	Document    InputDocument         `tl:"document"`
	SendMessage InputBotInlineMessage `tl:"send_message"`
}

func (*InputBotInlineResultDocumentPredict) CRC() uint32 {
	return 0xfff8fdc4
}
func (*InputBotInlineResultDocumentPredict) _InputBotInlineResult() {}

type InputBotInlineResultGamePredict struct {
	ID          string                `tl:"id"`
	ShortName   string                `tl:"short_name"`
	SendMessage InputBotInlineMessage `tl:"send_message"`
}

func (*InputBotInlineResultGamePredict) CRC() uint32 {
	return 0x4fa417f2
}
func (*InputBotInlineResultGamePredict) _InputBotInlineResult() {}

type InputBusinessAwayMessage interface {
	tl.TLObject
	_InputBusinessAwayMessage()
}

var (
	_ InputBusinessAwayMessage = (*InputBusinessAwayMessagePredict)(nil)
)

type InputBusinessAwayMessagePredict struct {
	_           struct{}                    `tl:"flags,bitflag"`
	OfflineOnly bool                        `tl:"offline_only,omitempty:flags:0,implicit"`
	ShortcutID  int32                       `tl:"shortcut_id"`
	Schedule    BusinessAwayMessageSchedule `tl:"schedule"`
	Recipients  InputBusinessRecipients     `tl:"recipients"`
}

func (*InputBusinessAwayMessagePredict) CRC() uint32 {
	return 0x832175e0
}
func (*InputBusinessAwayMessagePredict) _InputBusinessAwayMessage() {}

type InputBusinessBotRecipients interface {
	tl.TLObject
	_InputBusinessBotRecipients()
}

var (
	_ InputBusinessBotRecipients = (*InputBusinessBotRecipientsPredict)(nil)
)

type InputBusinessBotRecipientsPredict struct {
	_               struct{}    `tl:"flags,bitflag"`
	ExistingChats   bool        `tl:"existing_chats,omitempty:flags:0,implicit"`
	NewChats        bool        `tl:"new_chats,omitempty:flags:1,implicit"`
	Contacts        bool        `tl:"contacts,omitempty:flags:2,implicit"`
	NonContacts     bool        `tl:"non_contacts,omitempty:flags:3,implicit"`
	ExcludeSelected bool        `tl:"exclude_selected,omitempty:flags:5,implicit"`
	Users           []InputUser `tl:"users,omitempty:flags:4"`
	ExcludeUsers    []InputUser `tl:"exclude_users,omitempty:flags:6"`
}

func (*InputBusinessBotRecipientsPredict) CRC() uint32 {
	return 0xc4e5921e
}
func (*InputBusinessBotRecipientsPredict) _InputBusinessBotRecipients() {}

type InputBusinessChatLink interface {
	tl.TLObject
	_InputBusinessChatLink()
}

var (
	_ InputBusinessChatLink = (*InputBusinessChatLinkPredict)(nil)
)

type InputBusinessChatLinkPredict struct {
	_        struct{}        `tl:"flags,bitflag"`
	Message  string          `tl:"message"`
	Entities []MessageEntity `tl:"entities,omitempty:flags:0"`
	Title    *string         `tl:"title,omitempty:flags:1"`
}

func (*InputBusinessChatLinkPredict) CRC() uint32 {
	return 0x11679fa7
}
func (*InputBusinessChatLinkPredict) _InputBusinessChatLink() {}

type InputBusinessGreetingMessage interface {
	tl.TLObject
	_InputBusinessGreetingMessage()
}

var (
	_ InputBusinessGreetingMessage = (*InputBusinessGreetingMessagePredict)(nil)
)

type InputBusinessGreetingMessagePredict struct {
	ShortcutID     int32                   `tl:"shortcut_id"`
	Recipients     InputBusinessRecipients `tl:"recipients"`
	NoActivityDays int32                   `tl:"no_activity_days"`
}

func (*InputBusinessGreetingMessagePredict) CRC() uint32 {
	return 0x0194cb3b
}
func (*InputBusinessGreetingMessagePredict) _InputBusinessGreetingMessage() {}

type InputBusinessIntro interface {
	tl.TLObject
	_InputBusinessIntro()
}

var (
	_ InputBusinessIntro = (*InputBusinessIntroPredict)(nil)
)

type InputBusinessIntroPredict struct {
	_           struct{}      `tl:"flags,bitflag"`
	Title       string        `tl:"title"`
	Description string        `tl:"description"`
	Sticker     InputDocument `tl:"sticker,omitempty:flags:0"`
}

func (*InputBusinessIntroPredict) CRC() uint32 {
	return 0x09c469cd
}
func (*InputBusinessIntroPredict) _InputBusinessIntro() {}

type InputBusinessRecipients interface {
	tl.TLObject
	_InputBusinessRecipients()
}

var (
	_ InputBusinessRecipients = (*InputBusinessRecipientsPredict)(nil)
)

type InputBusinessRecipientsPredict struct {
	_               struct{}    `tl:"flags,bitflag"`
	ExistingChats   bool        `tl:"existing_chats,omitempty:flags:0,implicit"`
	NewChats        bool        `tl:"new_chats,omitempty:flags:1,implicit"`
	Contacts        bool        `tl:"contacts,omitempty:flags:2,implicit"`
	NonContacts     bool        `tl:"non_contacts,omitempty:flags:3,implicit"`
	ExcludeSelected bool        `tl:"exclude_selected,omitempty:flags:5,implicit"`
	Users           []InputUser `tl:"users,omitempty:flags:4"`
}

func (*InputBusinessRecipientsPredict) CRC() uint32 {
	return 0x6f8b32aa
}
func (*InputBusinessRecipientsPredict) _InputBusinessRecipients() {}

type InputChannel interface {
	tl.TLObject
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
	ChannelID  int64 `tl:"channel_id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputChannelPredict) CRC() uint32 {
	return 0xf35aec28
}
func (*InputChannelPredict) _InputChannel() {}

type InputChannelFromMessagePredict struct {
	Peer      InputPeer `tl:"peer"`
	MsgID     int32     `tl:"msg_id"`
	ChannelID int64     `tl:"channel_id"`
}

func (*InputChannelFromMessagePredict) CRC() uint32 {
	return 0x5b934f9d
}
func (*InputChannelFromMessagePredict) _InputChannel() {}

type InputChatPhoto interface {
	tl.TLObject
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
	File             InputFile `tl:"file,omitempty:flags:0"`
	Video            InputFile `tl:"video,omitempty:flags:1"`
	VideoStartTs     *float64  `tl:"video_start_ts,omitempty:flags:2"`
	VideoEmojiMarkup VideoSize `tl:"video_emoji_markup,omitempty:flags:3"`
}

func (*InputChatUploadedPhotoPredict) CRC() uint32 {
	return 0xbdcdaec0
}
func (*InputChatUploadedPhotoPredict) _InputChatPhoto() {}

type InputChatPhotoPredict struct {
	ID InputPhoto `tl:"id"`
}

func (*InputChatPhotoPredict) CRC() uint32 {
	return 0x8953ad37
}
func (*InputChatPhotoPredict) _InputChatPhoto() {}

type InputChatlist interface {
	tl.TLObject
	_InputChatlist()
}

var (
	_ InputChatlist = (*InputChatlistDialogFilterPredict)(nil)
)

type InputChatlistDialogFilterPredict struct {
	FilterID int32 `tl:"filter_id"`
}

func (*InputChatlistDialogFilterPredict) CRC() uint32 {
	return 0xf3e0da33
}
func (*InputChatlistDialogFilterPredict) _InputChatlist() {}

type InputCheckPasswordSRP interface {
	tl.TLObject
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
	SRPID int64  `tl:"srp_id"`
	A     []byte `tl:"A"`
	M1    []byte `tl:"M1"`
}

func (*InputCheckPasswordSRPPredict) CRC() uint32 {
	return 0xd27ff082
}
func (*InputCheckPasswordSRPPredict) _InputCheckPasswordSRP() {}

type InputClientProxy interface {
	tl.TLObject
	_InputClientProxy()
}

var (
	_ InputClientProxy = (*InputClientProxyPredict)(nil)
)

type InputClientProxyPredict struct {
	Address string `tl:"address"`
	Port    int32  `tl:"port"`
}

func (*InputClientProxyPredict) CRC() uint32 {
	return 0x75588b3f
}
func (*InputClientProxyPredict) _InputClientProxy() {}

type InputCollectible interface {
	tl.TLObject
	_InputCollectible()
}

var (
	_ InputCollectible = (*InputCollectibleUsernamePredict)(nil)
	_ InputCollectible = (*InputCollectiblePhonePredict)(nil)
)

type InputCollectibleUsernamePredict struct {
	Username string `tl:"username"`
}

func (*InputCollectibleUsernamePredict) CRC() uint32 {
	return 0xe39460a9
}
func (*InputCollectibleUsernamePredict) _InputCollectible() {}

type InputCollectiblePhonePredict struct {
	Phone string `tl:"phone"`
}

func (*InputCollectiblePhonePredict) CRC() uint32 {
	return 0xa2e214a4
}
func (*InputCollectiblePhonePredict) _InputCollectible() {}

type InputContact interface {
	tl.TLObject
	_InputContact()
}

var (
	_ InputContact = (*InputPhoneContactPredict)(nil)
)

type InputPhoneContactPredict struct {
	ClientID  int64  `tl:"client_id"`
	Phone     string `tl:"phone"`
	FirstName string `tl:"first_name"`
	LastName  string `tl:"last_name"`
}

func (*InputPhoneContactPredict) CRC() uint32 {
	return 0xf392b7f4
}
func (*InputPhoneContactPredict) _InputContact() {}

type InputDialogPeer interface {
	tl.TLObject
	_InputDialogPeer()
}

var (
	_ InputDialogPeer = (*InputDialogPeerPredict)(nil)
	_ InputDialogPeer = (*InputDialogPeerFolderPredict)(nil)
)

type InputDialogPeerPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*InputDialogPeerPredict) CRC() uint32 {
	return 0xfcaafeb7
}
func (*InputDialogPeerPredict) _InputDialogPeer() {}

type InputDialogPeerFolderPredict struct {
	FolderID int32 `tl:"folder_id"`
}

func (*InputDialogPeerFolderPredict) CRC() uint32 {
	return 0x64600527
}
func (*InputDialogPeerFolderPredict) _InputDialogPeer() {}

type InputDocument interface {
	tl.TLObject
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
	ID            int64  `tl:"id"`
	AccessHash    int64  `tl:"access_hash"`
	FileReference []byte `tl:"file_reference"`
}

func (*InputDocumentPredict) CRC() uint32 {
	return 0x1abfb575
}
func (*InputDocumentPredict) _InputDocument() {}

type InputEncryptedChat interface {
	tl.TLObject
	_InputEncryptedChat()
}

var (
	_ InputEncryptedChat = (*InputEncryptedChatPredict)(nil)
)

type InputEncryptedChatPredict struct {
	ChatID     int32 `tl:"chat_id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputEncryptedChatPredict) CRC() uint32 {
	return 0xf141b5e1
}
func (*InputEncryptedChatPredict) _InputEncryptedChat() {}

type InputEncryptedFile interface {
	tl.TLObject
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
	ID             int64  `tl:"id"`
	Parts          int32  `tl:"parts"`
	Md5Checksum    string `tl:"md5_checksum"`
	KeyFingerprint int32  `tl:"key_fingerprint"`
}

func (*InputEncryptedFileUploadedPredict) CRC() uint32 {
	return 0x64bd0306
}
func (*InputEncryptedFileUploadedPredict) _InputEncryptedFile() {}

type InputEncryptedFilePredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputEncryptedFilePredict) CRC() uint32 {
	return 0x5a17b5e5
}
func (*InputEncryptedFilePredict) _InputEncryptedFile() {}

type InputEncryptedFileBigUploadedPredict struct {
	ID             int64 `tl:"id"`
	Parts          int32 `tl:"parts"`
	KeyFingerprint int32 `tl:"key_fingerprint"`
}

func (*InputEncryptedFileBigUploadedPredict) CRC() uint32 {
	return 0x2dc173c8
}
func (*InputEncryptedFileBigUploadedPredict) _InputEncryptedFile() {}

type InputFile interface {
	tl.TLObject
	_InputFile()
}

var (
	_ InputFile = (*InputFilePredict)(nil)
	_ InputFile = (*InputFileBigPredict)(nil)
)

type InputFilePredict struct {
	ID          int64  `tl:"id"`
	Parts       int32  `tl:"parts"`
	Name        string `tl:"name"`
	Md5Checksum string `tl:"md5_checksum"`
}

func (*InputFilePredict) CRC() uint32 {
	return 0xf52ff27f
}
func (*InputFilePredict) _InputFile() {}

type InputFileBigPredict struct {
	ID    int64  `tl:"id"`
	Parts int32  `tl:"parts"`
	Name  string `tl:"name"`
}

func (*InputFileBigPredict) CRC() uint32 {
	return 0xfa4f0bb5
}
func (*InputFileBigPredict) _InputFile() {}

type InputFileLocation interface {
	tl.TLObject
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
	VolumeID      int64  `tl:"volume_id"`
	LocalID       int32  `tl:"local_id"`
	Secret        int64  `tl:"secret"`
	FileReference []byte `tl:"file_reference"`
}

func (*InputFileLocationPredict) CRC() uint32 {
	return 0xdfdaabe1
}
func (*InputFileLocationPredict) _InputFileLocation() {}

type InputEncryptedFileLocationPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputEncryptedFileLocationPredict) CRC() uint32 {
	return 0xf5235d55
}
func (*InputEncryptedFileLocationPredict) _InputFileLocation() {}

type InputDocumentFileLocationPredict struct {
	ID            int64  `tl:"id"`
	AccessHash    int64  `tl:"access_hash"`
	FileReference []byte `tl:"file_reference"`
	ThumbSize     string `tl:"thumb_size"`
}

func (*InputDocumentFileLocationPredict) CRC() uint32 {
	return 0xbad07584
}
func (*InputDocumentFileLocationPredict) _InputFileLocation() {}

type InputSecureFileLocationPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
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
	ID            int64  `tl:"id"`
	AccessHash    int64  `tl:"access_hash"`
	FileReference []byte `tl:"file_reference"`
	ThumbSize     string `tl:"thumb_size"`
}

func (*InputPhotoFileLocationPredict) CRC() uint32 {
	return 0x40181ffe
}
func (*InputPhotoFileLocationPredict) _InputFileLocation() {}

type InputPhotoLegacyFileLocationPredict struct {
	ID            int64  `tl:"id"`
	AccessHash    int64  `tl:"access_hash"`
	FileReference []byte `tl:"file_reference"`
	VolumeID      int64  `tl:"volume_id"`
	LocalID       int32  `tl:"local_id"`
	Secret        int64  `tl:"secret"`
}

func (*InputPhotoLegacyFileLocationPredict) CRC() uint32 {
	return 0xd83466f3
}
func (*InputPhotoLegacyFileLocationPredict) _InputFileLocation() {}

type InputPeerPhotoFileLocationPredict struct {
	_       struct{}  `tl:"flags,bitflag"`
	Big     bool      `tl:"big,omitempty:flags:0,implicit"`
	Peer    InputPeer `tl:"peer"`
	PhotoID int64     `tl:"photo_id"`
}

func (*InputPeerPhotoFileLocationPredict) CRC() uint32 {
	return 0x37257e99
}
func (*InputPeerPhotoFileLocationPredict) _InputFileLocation() {}

type InputStickerSetThumbPredict struct {
	Stickerset   InputStickerSet `tl:"stickerset"`
	ThumbVersion int32           `tl:"thumb_version"`
}

func (*InputStickerSetThumbPredict) CRC() uint32 {
	return 0x9d84f3db
}
func (*InputStickerSetThumbPredict) _InputFileLocation() {}

type InputGroupCallStreamPredict struct {
	_            struct{}       `tl:"flags,bitflag"`
	Call         InputGroupCall `tl:"call"`
	TimeMs       int64          `tl:"time_ms"`
	Scale        int32          `tl:"scale"`
	VideoChannel *int32         `tl:"video_channel,omitempty:flags:0"`
	VideoQuality *int32         `tl:"video_quality,omitempty:flags:0"`
}

func (*InputGroupCallStreamPredict) CRC() uint32 {
	return 0x0598a92a
}
func (*InputGroupCallStreamPredict) _InputFileLocation() {}

type InputFolderPeer interface {
	tl.TLObject
	_InputFolderPeer()
}

var (
	_ InputFolderPeer = (*InputFolderPeerPredict)(nil)
)

type InputFolderPeerPredict struct {
	Peer     InputPeer `tl:"peer"`
	FolderID int32     `tl:"folder_id"`
}

func (*InputFolderPeerPredict) CRC() uint32 {
	return 0xfbd2c296
}
func (*InputFolderPeerPredict) _InputFolderPeer() {}

type InputGame interface {
	tl.TLObject
	_InputGame()
}

var (
	_ InputGame = (*InputGameIDPredict)(nil)
	_ InputGame = (*InputGameShortNamePredict)(nil)
)

type InputGameIDPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputGameIDPredict) CRC() uint32 {
	return 0x032c3e77
}
func (*InputGameIDPredict) _InputGame() {}

type InputGameShortNamePredict struct {
	BotID     InputUser `tl:"bot_id"`
	ShortName string    `tl:"short_name"`
}

func (*InputGameShortNamePredict) CRC() uint32 {
	return 0xc331e80a
}
func (*InputGameShortNamePredict) _InputGame() {}

type InputGeoPoint interface {
	tl.TLObject
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
	Lat            float64  `tl:"lat"`
	Long           float64  `tl:"long"`
	AccuracyRadius *int32   `tl:"accuracy_radius,omitempty:flags:0"`
}

func (*InputGeoPointPredict) CRC() uint32 {
	return 0x48222faf
}
func (*InputGeoPointPredict) _InputGeoPoint() {}

type InputGroupCall interface {
	tl.TLObject
	_InputGroupCall()
}

var (
	_ InputGroupCall = (*InputGroupCallPredict)(nil)
)

type InputGroupCallPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputGroupCallPredict) CRC() uint32 {
	return 0xd8aa840f
}
func (*InputGroupCallPredict) _InputGroupCall() {}

type InputInvoice interface {
	tl.TLObject
	_InputInvoice()
}

var (
	_ InputInvoice = (*InputInvoiceMessagePredict)(nil)
	_ InputInvoice = (*InputInvoiceSlugPredict)(nil)
	_ InputInvoice = (*InputInvoicePremiumGiftCodePredict)(nil)
	_ InputInvoice = (*InputInvoiceStarsPredict)(nil)
)

type InputInvoiceMessagePredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*InputInvoiceMessagePredict) CRC() uint32 {
	return 0xc5b56859
}
func (*InputInvoiceMessagePredict) _InputInvoice() {}

type InputInvoiceSlugPredict struct {
	Slug string `tl:"slug"`
}

func (*InputInvoiceSlugPredict) CRC() uint32 {
	return 0xc326caef
}
func (*InputInvoiceSlugPredict) _InputInvoice() {}

type InputInvoicePremiumGiftCodePredict struct {
	Purpose InputStorePaymentPurpose `tl:"purpose"`
	Option  PremiumGiftCodeOption    `tl:"option"`
}

func (*InputInvoicePremiumGiftCodePredict) CRC() uint32 {
	return 0x98986c0d
}
func (*InputInvoicePremiumGiftCodePredict) _InputInvoice() {}

type InputInvoiceStarsPredict struct {
	Purpose InputStorePaymentPurpose `tl:"purpose"`
}

func (*InputInvoiceStarsPredict) CRC() uint32 {
	return 0x65f00ce3
}
func (*InputInvoiceStarsPredict) _InputInvoice() {}

type InputMedia interface {
	tl.TLObject
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
	_          struct{}        `tl:"flags,bitflag"`
	Spoiler    bool            `tl:"spoiler,omitempty:flags:2,implicit"`
	File       InputFile       `tl:"file"`
	Stickers   []InputDocument `tl:"stickers,omitempty:flags:0"`
	TTLSeconds *int32          `tl:"ttl_seconds,omitempty:flags:1"`
}

func (*InputMediaUploadedPhotoPredict) CRC() uint32 {
	return 0x1e287d04
}
func (*InputMediaUploadedPhotoPredict) _InputMedia() {}

type InputMediaPhotoPredict struct {
	_          struct{}   `tl:"flags,bitflag"`
	Spoiler    bool       `tl:"spoiler,omitempty:flags:1,implicit"`
	ID         InputPhoto `tl:"id"`
	TTLSeconds *int32     `tl:"ttl_seconds,omitempty:flags:0"`
}

func (*InputMediaPhotoPredict) CRC() uint32 {
	return 0xb3ba0635
}
func (*InputMediaPhotoPredict) _InputMedia() {}

type InputMediaGeoPointPredict struct {
	GeoPoint InputGeoPoint `tl:"geo_point"`
}

func (*InputMediaGeoPointPredict) CRC() uint32 {
	return 0xf9c44144
}
func (*InputMediaGeoPointPredict) _InputMedia() {}

type InputMediaContactPredict struct {
	PhoneNumber string `tl:"phone_number"`
	FirstName   string `tl:"first_name"`
	LastName    string `tl:"last_name"`
	Vcard       string `tl:"vcard"`
}

func (*InputMediaContactPredict) CRC() uint32 {
	return 0xf8ab7dfb
}
func (*InputMediaContactPredict) _InputMedia() {}

type InputMediaUploadedDocumentPredict struct {
	_            struct{}            `tl:"flags,bitflag"`
	NosoundVideo bool                `tl:"nosound_video,omitempty:flags:3,implicit"`
	ForceFile    bool                `tl:"force_file,omitempty:flags:4,implicit"`
	Spoiler      bool                `tl:"spoiler,omitempty:flags:5,implicit"`
	File         InputFile           `tl:"file"`
	Thumb        InputFile           `tl:"thumb,omitempty:flags:2"`
	MimeType     string              `tl:"mime_type"`
	Attributes   []DocumentAttribute `tl:"attributes"`
	Stickers     []InputDocument     `tl:"stickers,omitempty:flags:0"`
	TTLSeconds   *int32              `tl:"ttl_seconds,omitempty:flags:1"`
}

func (*InputMediaUploadedDocumentPredict) CRC() uint32 {
	return 0x5b38c6c1
}
func (*InputMediaUploadedDocumentPredict) _InputMedia() {}

type InputMediaDocumentPredict struct {
	_          struct{}      `tl:"flags,bitflag"`
	Spoiler    bool          `tl:"spoiler,omitempty:flags:2,implicit"`
	ID         InputDocument `tl:"id"`
	TTLSeconds *int32        `tl:"ttl_seconds,omitempty:flags:0"`
	Query      *string       `tl:"query,omitempty:flags:1"`
}

func (*InputMediaDocumentPredict) CRC() uint32 {
	return 0x33473058
}
func (*InputMediaDocumentPredict) _InputMedia() {}

type InputMediaVenuePredict struct {
	GeoPoint  InputGeoPoint `tl:"geo_point"`
	Title     string        `tl:"title"`
	Address   string        `tl:"address"`
	Provider  string        `tl:"provider"`
	VenueID   string        `tl:"venue_id"`
	VenueType string        `tl:"venue_type"`
}

func (*InputMediaVenuePredict) CRC() uint32 {
	return 0xc13d1c11
}
func (*InputMediaVenuePredict) _InputMedia() {}

type InputMediaPhotoExternalPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:"spoiler,omitempty:flags:1,implicit"`
	URL        string   `tl:"url"`
	TTLSeconds *int32   `tl:"ttl_seconds,omitempty:flags:0"`
}

func (*InputMediaPhotoExternalPredict) CRC() uint32 {
	return 0xe5bbfe1a
}
func (*InputMediaPhotoExternalPredict) _InputMedia() {}

type InputMediaDocumentExternalPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Spoiler    bool     `tl:"spoiler,omitempty:flags:1,implicit"`
	URL        string   `tl:"url"`
	TTLSeconds *int32   `tl:"ttl_seconds,omitempty:flags:0"`
}

func (*InputMediaDocumentExternalPredict) CRC() uint32 {
	return 0xfb52dc99
}
func (*InputMediaDocumentExternalPredict) _InputMedia() {}

type InputMediaGamePredict struct {
	ID InputGame `tl:"id"`
}

func (*InputMediaGamePredict) CRC() uint32 {
	return 0xd33f43f3
}
func (*InputMediaGamePredict) _InputMedia() {}

type InputMediaInvoicePredict struct {
	_             struct{}         `tl:"flags,bitflag"`
	Title         string           `tl:"title"`
	Description   string           `tl:"description"`
	Photo         InputWebDocument `tl:"photo,omitempty:flags:0"`
	Invoice       Invoice          `tl:"invoice"`
	Payload       []byte           `tl:"payload"`
	Provider      *string          `tl:"provider,omitempty:flags:3"`
	ProviderData  DataJSON         `tl:"provider_data"`
	StartParam    *string          `tl:"start_param,omitempty:flags:1"`
	ExtendedMedia InputMedia       `tl:"extended_media,omitempty:flags:2"`
}

func (*InputMediaInvoicePredict) CRC() uint32 {
	return 0x405fef0d
}
func (*InputMediaInvoicePredict) _InputMedia() {}

type InputMediaGeoLivePredict struct {
	_                           struct{}      `tl:"flags,bitflag"`
	Stopped                     bool          `tl:"stopped,omitempty:flags:0,implicit"`
	GeoPoint                    InputGeoPoint `tl:"geo_point"`
	Heading                     *int32        `tl:"heading,omitempty:flags:2"`
	Period                      *int32        `tl:"period,omitempty:flags:1"`
	ProximityNotificationRadius *int32        `tl:"proximity_notification_radius,omitempty:flags:3"`
}

func (*InputMediaGeoLivePredict) CRC() uint32 {
	return 0x971fa843
}
func (*InputMediaGeoLivePredict) _InputMedia() {}

type InputMediaPollPredict struct {
	_                struct{}        `tl:"flags,bitflag"`
	Poll             Poll            `tl:"poll"`
	CorrectAnswers   [][]byte        `tl:"correct_answers,omitempty:flags:0"`
	Solution         *string         `tl:"solution,omitempty:flags:1"`
	SolutionEntities []MessageEntity `tl:"solution_entities,omitempty:flags:1"`
}

func (*InputMediaPollPredict) CRC() uint32 {
	return 0x0f94e5f1
}
func (*InputMediaPollPredict) _InputMedia() {}

type InputMediaDicePredict struct {
	Emoticon string `tl:"emoticon"`
}

func (*InputMediaDicePredict) CRC() uint32 {
	return 0xe66fbf7b
}
func (*InputMediaDicePredict) _InputMedia() {}

type InputMediaStoryPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   int32     `tl:"id"`
}

func (*InputMediaStoryPredict) CRC() uint32 {
	return 0x89fdd778
}
func (*InputMediaStoryPredict) _InputMedia() {}

type InputMediaWebPagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ForceLargeMedia bool     `tl:"force_large_media,omitempty:flags:0,implicit"`
	ForceSmallMedia bool     `tl:"force_small_media,omitempty:flags:1,implicit"`
	Optional        bool     `tl:"optional,omitempty:flags:2,implicit"`
	URL             string   `tl:"url"`
}

func (*InputMediaWebPagePredict) CRC() uint32 {
	return 0xc21b8849
}
func (*InputMediaWebPagePredict) _InputMedia() {}

type InputMediaPaidMediaPredict struct {
	StarsAmount   int64        `tl:"stars_amount"`
	ExtendedMedia []InputMedia `tl:"extended_media"`
}

func (*InputMediaPaidMediaPredict) CRC() uint32 {
	return 0xaa661fc3
}
func (*InputMediaPaidMediaPredict) _InputMedia() {}

type InputMessage interface {
	tl.TLObject
	_InputMessage()
}

var (
	_ InputMessage = (*InputMessageIDPredict)(nil)
	_ InputMessage = (*InputMessageReplyToPredict)(nil)
	_ InputMessage = (*InputMessagePinnedPredict)(nil)
	_ InputMessage = (*InputMessageCallbackQueryPredict)(nil)
)

type InputMessageIDPredict struct {
	ID int32 `tl:"id"`
}

func (*InputMessageIDPredict) CRC() uint32 {
	return 0xa676a322
}
func (*InputMessageIDPredict) _InputMessage() {}

type InputMessageReplyToPredict struct {
	ID int32 `tl:"id"`
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
	ID      int32 `tl:"id"`
	QueryID int64 `tl:"query_id"`
}

func (*InputMessageCallbackQueryPredict) CRC() uint32 {
	return 0xacfa1a7e
}
func (*InputMessageCallbackQueryPredict) _InputMessage() {}

type InputNotifyPeer interface {
	tl.TLObject
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
	Peer InputPeer `tl:"peer"`
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
	Peer     InputPeer `tl:"peer"`
	TopMsgID int32     `tl:"top_msg_id"`
}

func (*InputNotifyForumTopicPredict) CRC() uint32 {
	return 0x5c467992
}
func (*InputNotifyForumTopicPredict) _InputNotifyPeer() {}

type InputPaymentCredentials interface {
	tl.TLObject
	_InputPaymentCredentials()
}

var (
	_ InputPaymentCredentials = (*InputPaymentCredentialsSavedPredict)(nil)
	_ InputPaymentCredentials = (*InputPaymentCredentialsPredict)(nil)
	_ InputPaymentCredentials = (*InputPaymentCredentialsApplePayPredict)(nil)
	_ InputPaymentCredentials = (*InputPaymentCredentialsGooglePayPredict)(nil)
)

type InputPaymentCredentialsSavedPredict struct {
	ID          string `tl:"id"`
	TmpPassword []byte `tl:"tmp_password"`
}

func (*InputPaymentCredentialsSavedPredict) CRC() uint32 {
	return 0xc10eb2cf
}
func (*InputPaymentCredentialsSavedPredict) _InputPaymentCredentials() {}

type InputPaymentCredentialsPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Save bool     `tl:"save,omitempty:flags:0,implicit"`
	Data DataJSON `tl:"data"`
}

func (*InputPaymentCredentialsPredict) CRC() uint32 {
	return 0x3417d728
}
func (*InputPaymentCredentialsPredict) _InputPaymentCredentials() {}

type InputPaymentCredentialsApplePayPredict struct {
	PaymentData DataJSON `tl:"payment_data"`
}

func (*InputPaymentCredentialsApplePayPredict) CRC() uint32 {
	return 0x0aa1c39f
}
func (*InputPaymentCredentialsApplePayPredict) _InputPaymentCredentials() {}

type InputPaymentCredentialsGooglePayPredict struct {
	PaymentToken DataJSON `tl:"payment_token"`
}

func (*InputPaymentCredentialsGooglePayPredict) CRC() uint32 {
	return 0x8ac32801
}
func (*InputPaymentCredentialsGooglePayPredict) _InputPaymentCredentials() {}

type InputPeer interface {
	tl.TLObject
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
	ChatID int64 `tl:"chat_id"`
}

func (*InputPeerChatPredict) CRC() uint32 {
	return 0x35a95cb9
}
func (*InputPeerChatPredict) _InputPeer() {}

type InputPeerUserPredict struct {
	UserID     int64 `tl:"user_id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputPeerUserPredict) CRC() uint32 {
	return 0xdde8a54c
}
func (*InputPeerUserPredict) _InputPeer() {}

type InputPeerChannelPredict struct {
	ChannelID  int64 `tl:"channel_id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputPeerChannelPredict) CRC() uint32 {
	return 0x27bcbbfc
}
func (*InputPeerChannelPredict) _InputPeer() {}

type InputPeerUserFromMessagePredict struct {
	Peer   InputPeer `tl:"peer"`
	MsgID  int32     `tl:"msg_id"`
	UserID int64     `tl:"user_id"`
}

func (*InputPeerUserFromMessagePredict) CRC() uint32 {
	return 0xa87b0a1c
}
func (*InputPeerUserFromMessagePredict) _InputPeer() {}

type InputPeerChannelFromMessagePredict struct {
	Peer      InputPeer `tl:"peer"`
	MsgID     int32     `tl:"msg_id"`
	ChannelID int64     `tl:"channel_id"`
}

func (*InputPeerChannelFromMessagePredict) CRC() uint32 {
	return 0xbd2a0840
}
func (*InputPeerChannelFromMessagePredict) _InputPeer() {}

type InputPeerNotifySettings interface {
	tl.TLObject
	_InputPeerNotifySettings()
}

var (
	_ InputPeerNotifySettings = (*InputPeerNotifySettingsPredict)(nil)
)

type InputPeerNotifySettingsPredict struct {
	_                 struct{}          `tl:"flags,bitflag"`
	ShowPreviews      *bool             `tl:"show_previews,omitempty:flags:0"`
	Silent            *bool             `tl:"silent,omitempty:flags:1"`
	MuteUntil         *int32            `tl:"mute_until,omitempty:flags:2"`
	Sound             NotificationSound `tl:"sound,omitempty:flags:3"`
	StoriesMuted      *bool             `tl:"stories_muted,omitempty:flags:6"`
	StoriesHideSender *bool             `tl:"stories_hide_sender,omitempty:flags:7"`
	StoriesSound      NotificationSound `tl:"stories_sound,omitempty:flags:8"`
}

func (*InputPeerNotifySettingsPredict) CRC() uint32 {
	return 0xcacb6ae2
}
func (*InputPeerNotifySettingsPredict) _InputPeerNotifySettings() {}

type InputPhoneCall interface {
	tl.TLObject
	_InputPhoneCall()
}

var (
	_ InputPhoneCall = (*InputPhoneCallPredict)(nil)
)

type InputPhoneCallPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputPhoneCallPredict) CRC() uint32 {
	return 0x1e36fded
}
func (*InputPhoneCallPredict) _InputPhoneCall() {}

type InputPhoto interface {
	tl.TLObject
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
	ID            int64  `tl:"id"`
	AccessHash    int64  `tl:"access_hash"`
	FileReference []byte `tl:"file_reference"`
}

func (*InputPhotoPredict) CRC() uint32 {
	return 0x3bb3b94a
}
func (*InputPhotoPredict) _InputPhoto() {}

type InputPrivacyKey interface {
	tl.TLObject
	_InputPrivacyKey()
}

var (
	_ InputPrivacyKey = (*InputPrivacyKeyStatusTimestampPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyChatInvitePredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyPhoneCallPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyPhoneP2PPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyForwardsPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyProfilePhotoPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyPhoneNumberPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyAddedByPhonePredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyVoiceMessagesPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyAboutPredict)(nil)
	_ InputPrivacyKey = (*InputPrivacyKeyBirthdayPredict)(nil)
)

type InputPrivacyKeyStatusTimestampPredict struct{}

func (*InputPrivacyKeyStatusTimestampPredict) CRC() uint32 {
	return 0x4f96cb18
}
func (*InputPrivacyKeyStatusTimestampPredict) _InputPrivacyKey() {}

type InputPrivacyKeyChatInvitePredict struct{}

func (*InputPrivacyKeyChatInvitePredict) CRC() uint32 {
	return 0xbdfb0426
}
func (*InputPrivacyKeyChatInvitePredict) _InputPrivacyKey() {}

type InputPrivacyKeyPhoneCallPredict struct{}

func (*InputPrivacyKeyPhoneCallPredict) CRC() uint32 {
	return 0xfabadc5f
}
func (*InputPrivacyKeyPhoneCallPredict) _InputPrivacyKey() {}

type InputPrivacyKeyPhoneP2PPredict struct{}

func (*InputPrivacyKeyPhoneP2PPredict) CRC() uint32 {
	return 0xdb9e70d2
}
func (*InputPrivacyKeyPhoneP2PPredict) _InputPrivacyKey() {}

type InputPrivacyKeyForwardsPredict struct{}

func (*InputPrivacyKeyForwardsPredict) CRC() uint32 {
	return 0xa4dd4c08
}
func (*InputPrivacyKeyForwardsPredict) _InputPrivacyKey() {}

type InputPrivacyKeyProfilePhotoPredict struct{}

func (*InputPrivacyKeyProfilePhotoPredict) CRC() uint32 {
	return 0x5719bacc
}
func (*InputPrivacyKeyProfilePhotoPredict) _InputPrivacyKey() {}

type InputPrivacyKeyPhoneNumberPredict struct{}

func (*InputPrivacyKeyPhoneNumberPredict) CRC() uint32 {
	return 0x0352dafa
}
func (*InputPrivacyKeyPhoneNumberPredict) _InputPrivacyKey() {}

type InputPrivacyKeyAddedByPhonePredict struct{}

func (*InputPrivacyKeyAddedByPhonePredict) CRC() uint32 {
	return 0xd1219bdd
}
func (*InputPrivacyKeyAddedByPhonePredict) _InputPrivacyKey() {}

type InputPrivacyKeyVoiceMessagesPredict struct{}

func (*InputPrivacyKeyVoiceMessagesPredict) CRC() uint32 {
	return 0xaee69d68
}
func (*InputPrivacyKeyVoiceMessagesPredict) _InputPrivacyKey() {}

type InputPrivacyKeyAboutPredict struct{}

func (*InputPrivacyKeyAboutPredict) CRC() uint32 {
	return 0x3823cc40
}
func (*InputPrivacyKeyAboutPredict) _InputPrivacyKey() {}

type InputPrivacyKeyBirthdayPredict struct{}

func (*InputPrivacyKeyBirthdayPredict) CRC() uint32 {
	return 0xd65a11cc
}
func (*InputPrivacyKeyBirthdayPredict) _InputPrivacyKey() {}

type InputPrivacyRule interface {
	tl.TLObject
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
	return 0x0d09e07b
}
func (*InputPrivacyValueAllowContactsPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowAllPredict struct{}

func (*InputPrivacyValueAllowAllPredict) CRC() uint32 {
	return 0x184b35ce
}
func (*InputPrivacyValueAllowAllPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowUsersPredict struct {
	Users []InputUser `tl:"users"`
}

func (*InputPrivacyValueAllowUsersPredict) CRC() uint32 {
	return 0x131cc67f
}
func (*InputPrivacyValueAllowUsersPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowContactsPredict struct{}

func (*InputPrivacyValueDisallowContactsPredict) CRC() uint32 {
	return 0x0ba52007
}
func (*InputPrivacyValueDisallowContactsPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowAllPredict struct{}

func (*InputPrivacyValueDisallowAllPredict) CRC() uint32 {
	return 0xd66b66c9
}
func (*InputPrivacyValueDisallowAllPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowUsersPredict struct {
	Users []InputUser `tl:"users"`
}

func (*InputPrivacyValueDisallowUsersPredict) CRC() uint32 {
	return 0x90110467
}
func (*InputPrivacyValueDisallowUsersPredict) _InputPrivacyRule() {}

type InputPrivacyValueAllowChatParticipantsPredict struct {
	Chats []int64 `tl:"chats"`
}

func (*InputPrivacyValueAllowChatParticipantsPredict) CRC() uint32 {
	return 0x840649cf
}
func (*InputPrivacyValueAllowChatParticipantsPredict) _InputPrivacyRule() {}

type InputPrivacyValueDisallowChatParticipantsPredict struct {
	Chats []int64 `tl:"chats"`
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
	tl.TLObject
	_InputQuickReplyShortcut()
}

var (
	_ InputQuickReplyShortcut = (*InputQuickReplyShortcutPredict)(nil)
	_ InputQuickReplyShortcut = (*InputQuickReplyShortcutIDPredict)(nil)
)

type InputQuickReplyShortcutPredict struct {
	Shortcut string `tl:"shortcut"`
}

func (*InputQuickReplyShortcutPredict) CRC() uint32 {
	return 0x24596d41
}
func (*InputQuickReplyShortcutPredict) _InputQuickReplyShortcut() {}

type InputQuickReplyShortcutIDPredict struct {
	ShortcutID int32 `tl:"shortcut_id"`
}

func (*InputQuickReplyShortcutIDPredict) CRC() uint32 {
	return 0x01190cf1
}
func (*InputQuickReplyShortcutIDPredict) _InputQuickReplyShortcut() {}

type InputReplyTo interface {
	tl.TLObject
	_InputReplyTo()
}

var (
	_ InputReplyTo = (*InputReplyToMessagePredict)(nil)
	_ InputReplyTo = (*InputReplyToStoryPredict)(nil)
)

type InputReplyToMessagePredict struct {
	_             struct{}        `tl:"flags,bitflag"`
	ReplyToMsgID  int32           `tl:"reply_to_msg_id"`
	TopMsgID      *int32          `tl:"top_msg_id,omitempty:flags:0"`
	ReplyToPeerID InputPeer       `tl:"reply_to_peer_id,omitempty:flags:1"`
	QuoteText     *string         `tl:"quote_text,omitempty:flags:2"`
	QuoteEntities []MessageEntity `tl:"quote_entities,omitempty:flags:3"`
	QuoteOffset   *int32          `tl:"quote_offset,omitempty:flags:4"`
}

func (*InputReplyToMessagePredict) CRC() uint32 {
	return 0x22c0f6d5
}
func (*InputReplyToMessagePredict) _InputReplyTo() {}

type InputReplyToStoryPredict struct {
	Peer    InputPeer `tl:"peer"`
	StoryID int32     `tl:"story_id"`
}

func (*InputReplyToStoryPredict) CRC() uint32 {
	return 0x5881323a
}
func (*InputReplyToStoryPredict) _InputReplyTo() {}

type InputSecureFile interface {
	tl.TLObject
	_InputSecureFile()
}

var (
	_ InputSecureFile = (*InputSecureFileUploadedPredict)(nil)
	_ InputSecureFile = (*InputSecureFilePredict)(nil)
)

type InputSecureFileUploadedPredict struct {
	ID          int64  `tl:"id"`
	Parts       int32  `tl:"parts"`
	Md5Checksum string `tl:"md5_checksum"`
	FileHash    []byte `tl:"file_hash"`
	Secret      []byte `tl:"secret"`
}

func (*InputSecureFileUploadedPredict) CRC() uint32 {
	return 0x3334b0f0
}
func (*InputSecureFileUploadedPredict) _InputSecureFile() {}

type InputSecureFilePredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputSecureFilePredict) CRC() uint32 {
	return 0x5367e5be
}
func (*InputSecureFilePredict) _InputSecureFile() {}

type InputSecureValue interface {
	tl.TLObject
	_InputSecureValue()
}

var (
	_ InputSecureValue = (*InputSecureValuePredict)(nil)
)

type InputSecureValuePredict struct {
	_           struct{}          `tl:"flags,bitflag"`
	Type        SecureValueType   `tl:"type"`
	Data        SecureData        `tl:"data,omitempty:flags:0"`
	FrontSide   InputSecureFile   `tl:"front_side,omitempty:flags:1"`
	ReverseSide InputSecureFile   `tl:"reverse_side,omitempty:flags:2"`
	Selfie      InputSecureFile   `tl:"selfie,omitempty:flags:3"`
	Translation []InputSecureFile `tl:"translation,omitempty:flags:6"`
	Files       []InputSecureFile `tl:"files,omitempty:flags:4"`
	PlainData   SecurePlainData   `tl:"plain_data,omitempty:flags:5"`
}

func (*InputSecureValuePredict) CRC() uint32 {
	return 0xdb21d0a7
}
func (*InputSecureValuePredict) _InputSecureValue() {}

type InputSingleMedia interface {
	tl.TLObject
	_InputSingleMedia()
}

var (
	_ InputSingleMedia = (*InputSingleMediaPredict)(nil)
)

type InputSingleMediaPredict struct {
	_        struct{}        `tl:"flags,bitflag"`
	Media    InputMedia      `tl:"media"`
	RandomID int64           `tl:"random_id"`
	Message  string          `tl:"message"`
	Entities []MessageEntity `tl:"entities,omitempty:flags:0"`
}

func (*InputSingleMediaPredict) CRC() uint32 {
	return 0x1cc6e91f
}
func (*InputSingleMediaPredict) _InputSingleMedia() {}

type InputStarsTransaction interface {
	tl.TLObject
	_InputStarsTransaction()
}

var (
	_ InputStarsTransaction = (*InputStarsTransactionPredict)(nil)
)

type InputStarsTransactionPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Refund bool     `tl:"refund,omitempty:flags:0,implicit"`
	ID     string   `tl:"id"`
}

func (*InputStarsTransactionPredict) CRC() uint32 {
	return 0x206ae6d1
}
func (*InputStarsTransactionPredict) _InputStarsTransaction() {}

type InputStickerSet interface {
	tl.TLObject
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
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputStickerSetIDPredict) CRC() uint32 {
	return 0x9de7a269
}
func (*InputStickerSetIDPredict) _InputStickerSet() {}

type InputStickerSetShortNamePredict struct {
	ShortName string `tl:"short_name"`
}

func (*InputStickerSetShortNamePredict) CRC() uint32 {
	return 0x861cc8a0
}
func (*InputStickerSetShortNamePredict) _InputStickerSet() {}

type InputStickerSetAnimatedEmojiPredict struct{}

func (*InputStickerSetAnimatedEmojiPredict) CRC() uint32 {
	return 0x028703c8
}
func (*InputStickerSetAnimatedEmojiPredict) _InputStickerSet() {}

type InputStickerSetDicePredict struct {
	Emoticon string `tl:"emoticon"`
}

func (*InputStickerSetDicePredict) CRC() uint32 {
	return 0xe67f520e
}
func (*InputStickerSetDicePredict) _InputStickerSet() {}

type InputStickerSetAnimatedEmojiAnimationsPredict struct{}

func (*InputStickerSetAnimatedEmojiAnimationsPredict) CRC() uint32 {
	return 0x0cde3739
}
func (*InputStickerSetAnimatedEmojiAnimationsPredict) _InputStickerSet() {}

type InputStickerSetPremiumGiftsPredict struct{}

func (*InputStickerSetPremiumGiftsPredict) CRC() uint32 {
	return 0xc88b3b02
}
func (*InputStickerSetPremiumGiftsPredict) _InputStickerSet() {}

type InputStickerSetEmojiGenericAnimationsPredict struct{}

func (*InputStickerSetEmojiGenericAnimationsPredict) CRC() uint32 {
	return 0x04c4d4ce
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
	tl.TLObject
	_InputStickerSetItem()
}

var (
	_ InputStickerSetItem = (*InputStickerSetItemPredict)(nil)
)

type InputStickerSetItemPredict struct {
	_          struct{}      `tl:"flags,bitflag"`
	Document   InputDocument `tl:"document"`
	Emoji      string        `tl:"emoji"`
	MaskCoords MaskCoords    `tl:"mask_coords,omitempty:flags:0"`
	Keywords   *string       `tl:"keywords,omitempty:flags:1"`
}

func (*InputStickerSetItemPredict) CRC() uint32 {
	return 0x32da9e9c
}
func (*InputStickerSetItemPredict) _InputStickerSetItem() {}

type InputStickeredMedia interface {
	tl.TLObject
	_InputStickeredMedia()
}

var (
	_ InputStickeredMedia = (*InputStickeredMediaPhotoPredict)(nil)
	_ InputStickeredMedia = (*InputStickeredMediaDocumentPredict)(nil)
)

type InputStickeredMediaPhotoPredict struct {
	ID InputPhoto `tl:"id"`
}

func (*InputStickeredMediaPhotoPredict) CRC() uint32 {
	return 0x4a992157
}
func (*InputStickeredMediaPhotoPredict) _InputStickeredMedia() {}

type InputStickeredMediaDocumentPredict struct {
	ID InputDocument `tl:"id"`
}

func (*InputStickeredMediaDocumentPredict) CRC() uint32 {
	return 0x0438865b
}
func (*InputStickeredMediaDocumentPredict) _InputStickeredMedia() {}

type InputStorePaymentPurpose interface {
	tl.TLObject
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
	Restore bool     `tl:"restore,omitempty:flags:0,implicit"`
	Upgrade bool     `tl:"upgrade,omitempty:flags:1,implicit"`
}

func (*InputStorePaymentPremiumSubscriptionPredict) CRC() uint32 {
	return 0xa6751e66
}
func (*InputStorePaymentPremiumSubscriptionPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentGiftPremiumPredict struct {
	UserID   InputUser `tl:"user_id"`
	Currency string    `tl:"currency"`
	Amount   int64     `tl:"amount"`
}

func (*InputStorePaymentGiftPremiumPredict) CRC() uint32 {
	return 0x616f7fe8
}
func (*InputStorePaymentGiftPremiumPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentPremiumGiftCodePredict struct {
	_         struct{}    `tl:"flags,bitflag"`
	Users     []InputUser `tl:"users"`
	BoostPeer InputPeer   `tl:"boost_peer,omitempty:flags:0"`
	Currency  string      `tl:"currency"`
	Amount    int64       `tl:"amount"`
}

func (*InputStorePaymentPremiumGiftCodePredict) CRC() uint32 {
	return 0xa3805f3f
}
func (*InputStorePaymentPremiumGiftCodePredict) _InputStorePaymentPurpose() {}

type InputStorePaymentPremiumGiveawayPredict struct {
	_                  struct{}    `tl:"flags,bitflag"`
	OnlyNewSubscribers bool        `tl:"only_new_subscribers,omitempty:flags:0,implicit"`
	WinnersAreVisible  bool        `tl:"winners_are_visible,omitempty:flags:3,implicit"`
	BoostPeer          InputPeer   `tl:"boost_peer"`
	AdditionalPeers    []InputPeer `tl:"additional_peers,omitempty:flags:1"`
	CountriesIso2      []string    `tl:"countries_iso2,omitempty:flags:2"`
	PrizeDescription   *string     `tl:"prize_description,omitempty:flags:4"`
	RandomID           int64       `tl:"random_id"`
	UntilDate          int32       `tl:"until_date"`
	Currency           string      `tl:"currency"`
	Amount             int64       `tl:"amount"`
}

func (*InputStorePaymentPremiumGiveawayPredict) CRC() uint32 {
	return 0x160544ca
}
func (*InputStorePaymentPremiumGiveawayPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentStarsTopupPredict struct {
	Stars    int64  `tl:"stars"`
	Currency string `tl:"currency"`
	Amount   int64  `tl:"amount"`
}

func (*InputStorePaymentStarsTopupPredict) CRC() uint32 {
	return 0xdddd0f56
}
func (*InputStorePaymentStarsTopupPredict) _InputStorePaymentPurpose() {}

type InputStorePaymentStarsGiftPredict struct {
	UserID   InputUser `tl:"user_id"`
	Stars    int64     `tl:"stars"`
	Currency string    `tl:"currency"`
	Amount   int64     `tl:"amount"`
}

func (*InputStorePaymentStarsGiftPredict) CRC() uint32 {
	return 0x1d741ef7
}
func (*InputStorePaymentStarsGiftPredict) _InputStorePaymentPurpose() {}

type InputTheme interface {
	tl.TLObject
	_InputTheme()
}

var (
	_ InputTheme = (*InputThemePredict)(nil)
	_ InputTheme = (*InputThemeSlugPredict)(nil)
)

type InputThemePredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputThemePredict) CRC() uint32 {
	return 0x3c5693e9
}
func (*InputThemePredict) _InputTheme() {}

type InputThemeSlugPredict struct {
	Slug string `tl:"slug"`
}

func (*InputThemeSlugPredict) CRC() uint32 {
	return 0xf5890df1
}
func (*InputThemeSlugPredict) _InputTheme() {}

type InputThemeSettings interface {
	tl.TLObject
	_InputThemeSettings()
}

var (
	_ InputThemeSettings = (*InputThemeSettingsPredict)(nil)
)

type InputThemeSettingsPredict struct {
	_                     struct{}          `tl:"flags,bitflag"`
	MessageColorsAnimated bool              `tl:"message_colors_animated,omitempty:flags:2,implicit"`
	BaseTheme             BaseTheme         `tl:"base_theme"`
	AccentColor           int32             `tl:"accent_color"`
	OutboxAccentColor     *int32            `tl:"outbox_accent_color,omitempty:flags:3"`
	MessageColors         []int32           `tl:"message_colors,omitempty:flags:0"`
	Wallpaper             InputWallPaper    `tl:"wallpaper,omitempty:flags:1"`
	WallpaperSettings     WallPaperSettings `tl:"wallpaper_settings,omitempty:flags:1"`
}

func (*InputThemeSettingsPredict) CRC() uint32 {
	return 0x8fde504f
}
func (*InputThemeSettingsPredict) _InputThemeSettings() {}

type InputUser interface {
	tl.TLObject
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
	UserID     int64 `tl:"user_id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputUserPredict) CRC() uint32 {
	return 0xf21158c6
}
func (*InputUserPredict) _InputUser() {}

type InputUserFromMessagePredict struct {
	Peer   InputPeer `tl:"peer"`
	MsgID  int32     `tl:"msg_id"`
	UserID int64     `tl:"user_id"`
}

func (*InputUserFromMessagePredict) CRC() uint32 {
	return 0x1da448e2
}
func (*InputUserFromMessagePredict) _InputUser() {}

type InputWallPaper interface {
	tl.TLObject
	_InputWallPaper()
}

var (
	_ InputWallPaper = (*InputWallPaperPredict)(nil)
	_ InputWallPaper = (*InputWallPaperSlugPredict)(nil)
	_ InputWallPaper = (*InputWallPaperNoFilePredict)(nil)
)

type InputWallPaperPredict struct {
	ID         int64 `tl:"id"`
	AccessHash int64 `tl:"access_hash"`
}

func (*InputWallPaperPredict) CRC() uint32 {
	return 0xe630b979
}
func (*InputWallPaperPredict) _InputWallPaper() {}

type InputWallPaperSlugPredict struct {
	Slug string `tl:"slug"`
}

func (*InputWallPaperSlugPredict) CRC() uint32 {
	return 0x72091c80
}
func (*InputWallPaperSlugPredict) _InputWallPaper() {}

type InputWallPaperNoFilePredict struct {
	ID int64 `tl:"id"`
}

func (*InputWallPaperNoFilePredict) CRC() uint32 {
	return 0x967a462e
}
func (*InputWallPaperNoFilePredict) _InputWallPaper() {}

type InputWebDocument interface {
	tl.TLObject
	_InputWebDocument()
}

var (
	_ InputWebDocument = (*InputWebDocumentPredict)(nil)
)

type InputWebDocumentPredict struct {
	URL        string              `tl:"url"`
	Size       int32               `tl:"size"`
	MimeType   string              `tl:"mime_type"`
	Attributes []DocumentAttribute `tl:"attributes"`
}

func (*InputWebDocumentPredict) CRC() uint32 {
	return 0x9bed434d
}
func (*InputWebDocumentPredict) _InputWebDocument() {}

type InputWebFileLocation interface {
	tl.TLObject
	_InputWebFileLocation()
}

var (
	_ InputWebFileLocation = (*InputWebFileLocationPredict)(nil)
	_ InputWebFileLocation = (*InputWebFileGeoPointLocationPredict)(nil)
	_ InputWebFileLocation = (*InputWebFileAudioAlbumThumbLocationPredict)(nil)
)

type InputWebFileLocationPredict struct {
	URL        string `tl:"url"`
	AccessHash int64  `tl:"access_hash"`
}

func (*InputWebFileLocationPredict) CRC() uint32 {
	return 0xc239d686
}
func (*InputWebFileLocationPredict) _InputWebFileLocation() {}

type InputWebFileGeoPointLocationPredict struct {
	GeoPoint   InputGeoPoint `tl:"geo_point"`
	AccessHash int64         `tl:"access_hash"`
	W          int32         `tl:"w"`
	H          int32         `tl:"h"`
	Zoom       int32         `tl:"zoom"`
	Scale      int32         `tl:"scale"`
}

func (*InputWebFileGeoPointLocationPredict) CRC() uint32 {
	return 0x9f2221c9
}
func (*InputWebFileGeoPointLocationPredict) _InputWebFileLocation() {}

type InputWebFileAudioAlbumThumbLocationPredict struct {
	_         struct{}      `tl:"flags,bitflag"`
	Small     bool          `tl:"small,omitempty:flags:2,implicit"`
	Document  InputDocument `tl:"document,omitempty:flags:0"`
	Title     *string       `tl:"title,omitempty:flags:1"`
	Performer *string       `tl:"performer,omitempty:flags:1"`
}

func (*InputWebFileAudioAlbumThumbLocationPredict) CRC() uint32 {
	return 0xf46fe924
}
func (*InputWebFileAudioAlbumThumbLocationPredict) _InputWebFileLocation() {}

type Invoice interface {
	tl.TLObject
	_Invoice()
}

var (
	_ Invoice = (*InvoicePredict)(nil)
)

type InvoicePredict struct {
	_                        struct{}       `tl:"flags,bitflag"`
	Test                     bool           `tl:"test,omitempty:flags:0,implicit"`
	NameRequested            bool           `tl:"name_requested,omitempty:flags:1,implicit"`
	PhoneRequested           bool           `tl:"phone_requested,omitempty:flags:2,implicit"`
	EmailRequested           bool           `tl:"email_requested,omitempty:flags:3,implicit"`
	ShippingAddressRequested bool           `tl:"shipping_address_requested,omitempty:flags:4,implicit"`
	Flexible                 bool           `tl:"flexible,omitempty:flags:5,implicit"`
	PhoneToProvider          bool           `tl:"phone_to_provider,omitempty:flags:6,implicit"`
	EmailToProvider          bool           `tl:"email_to_provider,omitempty:flags:7,implicit"`
	Recurring                bool           `tl:"recurring,omitempty:flags:9,implicit"`
	Currency                 string         `tl:"currency"`
	Prices                   []LabeledPrice `tl:"prices"`
	MaxTipAmount             *int64         `tl:"max_tip_amount,omitempty:flags:8"`
	SuggestedTipAmounts      []int64        `tl:"suggested_tip_amounts,omitempty:flags:8"`
	TermsURL                 *string        `tl:"terms_url,omitempty:flags:10"`
}

func (*InvoicePredict) CRC() uint32 {
	return 0x5db95a15
}
func (*InvoicePredict) _Invoice() {}

type JSONObjectValue interface {
	tl.TLObject
	_JSONObjectValue()
}

var (
	_ JSONObjectValue = (*JSONObjectValuePredict)(nil)
)

type JSONObjectValuePredict struct {
	Key   string    `tl:"key"`
	Value JSONValue `tl:"value"`
}

func (*JSONObjectValuePredict) CRC() uint32 {
	return 0xc0de1bd9
}
func (*JSONObjectValuePredict) _JSONObjectValue() {}

type JSONValue interface {
	tl.TLObject
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
	Value bool `tl:"value"`
}

func (*JSONBoolPredict) CRC() uint32 {
	return 0xc7345e6a
}
func (*JSONBoolPredict) _JSONValue() {}

type JSONNumberPredict struct {
	Value float64 `tl:"value"`
}

func (*JSONNumberPredict) CRC() uint32 {
	return 0x2be0dfa4
}
func (*JSONNumberPredict) _JSONValue() {}

type JSONStringPredict struct {
	Value string `tl:"value"`
}

func (*JSONStringPredict) CRC() uint32 {
	return 0xb71e767a
}
func (*JSONStringPredict) _JSONValue() {}

type JSONArrayPredict struct {
	Value []JSONValue `tl:"value"`
}

func (*JSONArrayPredict) CRC() uint32 {
	return 0xf7444763
}
func (*JSONArrayPredict) _JSONValue() {}

type JSONObjectPredict struct {
	Value []JSONObjectValue `tl:"value"`
}

func (*JSONObjectPredict) CRC() uint32 {
	return 0x99c1d49d
}
func (*JSONObjectPredict) _JSONValue() {}

type KeyboardButton interface {
	tl.TLObject
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
	Text string `tl:"text"`
}

func (*KeyboardButtonPredict) CRC() uint32 {
	return 0xa2fa4880
}
func (*KeyboardButtonPredict) _KeyboardButton() {}

type KeyboardButtonURLPredict struct {
	Text string `tl:"text"`
	URL  string `tl:"url"`
}

func (*KeyboardButtonURLPredict) CRC() uint32 {
	return 0x258aff05
}
func (*KeyboardButtonURLPredict) _KeyboardButton() {}

type KeyboardButtonCallbackPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	RequiresPassword bool     `tl:"requires_password,omitempty:flags:0,implicit"`
	Text             string   `tl:"text"`
	Data             []byte   `tl:"data"`
}

func (*KeyboardButtonCallbackPredict) CRC() uint32 {
	return 0x35bbdb6b
}
func (*KeyboardButtonCallbackPredict) _KeyboardButton() {}

type KeyboardButtonRequestPhonePredict struct {
	Text string `tl:"text"`
}

func (*KeyboardButtonRequestPhonePredict) CRC() uint32 {
	return 0xb16a6c29
}
func (*KeyboardButtonRequestPhonePredict) _KeyboardButton() {}

type KeyboardButtonRequestGeoLocationPredict struct {
	Text string `tl:"text"`
}

func (*KeyboardButtonRequestGeoLocationPredict) CRC() uint32 {
	return 0xfc796b3f
}
func (*KeyboardButtonRequestGeoLocationPredict) _KeyboardButton() {}

type KeyboardButtonSwitchInlinePredict struct {
	_         struct{}              `tl:"flags,bitflag"`
	SamePeer  bool                  `tl:"same_peer,omitempty:flags:0,implicit"`
	Text      string                `tl:"text"`
	Query     string                `tl:"query"`
	PeerTypes []InlineQueryPeerType `tl:"peer_types,omitempty:flags:1"`
}

func (*KeyboardButtonSwitchInlinePredict) CRC() uint32 {
	return 0x93b9fbb5
}
func (*KeyboardButtonSwitchInlinePredict) _KeyboardButton() {}

type KeyboardButtonGamePredict struct {
	Text string `tl:"text"`
}

func (*KeyboardButtonGamePredict) CRC() uint32 {
	return 0x50f41ccf
}
func (*KeyboardButtonGamePredict) _KeyboardButton() {}

type KeyboardButtonBuyPredict struct {
	Text string `tl:"text"`
}

func (*KeyboardButtonBuyPredict) CRC() uint32 {
	return 0xafd93fbb
}
func (*KeyboardButtonBuyPredict) _KeyboardButton() {}

type KeyboardButtonURLAuthPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Text     string   `tl:"text"`
	FwdText  *string  `tl:"fwd_text,omitempty:flags:0"`
	URL      string   `tl:"url"`
	ButtonID int32    `tl:"button_id"`
}

func (*KeyboardButtonURLAuthPredict) CRC() uint32 {
	return 0x10b78d29
}
func (*KeyboardButtonURLAuthPredict) _KeyboardButton() {}

type InputKeyboardButtonURLAuthPredict struct {
	_                  struct{}  `tl:"flags,bitflag"`
	RequestWriteAccess bool      `tl:"request_write_access,omitempty:flags:0,implicit"`
	Text               string    `tl:"text"`
	FwdText            *string   `tl:"fwd_text,omitempty:flags:1"`
	URL                string    `tl:"url"`
	Bot                InputUser `tl:"bot"`
}

func (*InputKeyboardButtonURLAuthPredict) CRC() uint32 {
	return 0xd02e7fd4
}
func (*InputKeyboardButtonURLAuthPredict) _KeyboardButton() {}

type KeyboardButtonRequestPollPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	Quiz *bool    `tl:"quiz,omitempty:flags:0"`
	Text string   `tl:"text"`
}

func (*KeyboardButtonRequestPollPredict) CRC() uint32 {
	return 0xbbc7515d
}
func (*KeyboardButtonRequestPollPredict) _KeyboardButton() {}

type InputKeyboardButtonUserProfilePredict struct {
	Text   string    `tl:"text"`
	UserID InputUser `tl:"user_id"`
}

func (*InputKeyboardButtonUserProfilePredict) CRC() uint32 {
	return 0xe988037b
}
func (*InputKeyboardButtonUserProfilePredict) _KeyboardButton() {}

type KeyboardButtonUserProfilePredict struct {
	Text   string `tl:"text"`
	UserID int64  `tl:"user_id"`
}

func (*KeyboardButtonUserProfilePredict) CRC() uint32 {
	return 0x308660c1
}
func (*KeyboardButtonUserProfilePredict) _KeyboardButton() {}

type KeyboardButtonWebViewPredict struct {
	Text string `tl:"text"`
	URL  string `tl:"url"`
}

func (*KeyboardButtonWebViewPredict) CRC() uint32 {
	return 0x13767230
}
func (*KeyboardButtonWebViewPredict) _KeyboardButton() {}

type KeyboardButtonSimpleWebViewPredict struct {
	Text string `tl:"text"`
	URL  string `tl:"url"`
}

func (*KeyboardButtonSimpleWebViewPredict) CRC() uint32 {
	return 0xa0c0505c
}
func (*KeyboardButtonSimpleWebViewPredict) _KeyboardButton() {}

type KeyboardButtonRequestPeerPredict struct {
	Text        string          `tl:"text"`
	ButtonID    int32           `tl:"button_id"`
	PeerType    RequestPeerType `tl:"peer_type"`
	MaxQuantity int32           `tl:"max_quantity"`
}

func (*KeyboardButtonRequestPeerPredict) CRC() uint32 {
	return 0x53d7bfd8
}
func (*KeyboardButtonRequestPeerPredict) _KeyboardButton() {}

type InputKeyboardButtonRequestPeerPredict struct {
	_                 struct{}        `tl:"flags,bitflag"`
	NameRequested     bool            `tl:"name_requested,omitempty:flags:0,implicit"`
	UsernameRequested bool            `tl:"username_requested,omitempty:flags:1,implicit"`
	PhotoRequested    bool            `tl:"photo_requested,omitempty:flags:2,implicit"`
	Text              string          `tl:"text"`
	ButtonID          int32           `tl:"button_id"`
	PeerType          RequestPeerType `tl:"peer_type"`
	MaxQuantity       int32           `tl:"max_quantity"`
}

func (*InputKeyboardButtonRequestPeerPredict) CRC() uint32 {
	return 0xc9662d05
}
func (*InputKeyboardButtonRequestPeerPredict) _KeyboardButton() {}

type KeyboardButtonRow interface {
	tl.TLObject
	_KeyboardButtonRow()
}

var (
	_ KeyboardButtonRow = (*KeyboardButtonRowPredict)(nil)
)

type KeyboardButtonRowPredict struct {
	Buttons []KeyboardButton `tl:"buttons"`
}

func (*KeyboardButtonRowPredict) CRC() uint32 {
	return 0x77608b83
}
func (*KeyboardButtonRowPredict) _KeyboardButtonRow() {}

type LabeledPrice interface {
	tl.TLObject
	_LabeledPrice()
}

var (
	_ LabeledPrice = (*LabeledPricePredict)(nil)
)

type LabeledPricePredict struct {
	Label  string `tl:"label"`
	Amount int64  `tl:"amount"`
}

func (*LabeledPricePredict) CRC() uint32 {
	return 0xcb296bf8
}
func (*LabeledPricePredict) _LabeledPrice() {}

type LangPackDifference interface {
	tl.TLObject
	_LangPackDifference()
}

var (
	_ LangPackDifference = (*LangPackDifferencePredict)(nil)
)

type LangPackDifferencePredict struct {
	LangCode    string           `tl:"lang_code"`
	FromVersion int32            `tl:"from_version"`
	Version     int32            `tl:"version"`
	Strings     []LangPackString `tl:"strings"`
}

func (*LangPackDifferencePredict) CRC() uint32 {
	return 0xf385c1f6
}
func (*LangPackDifferencePredict) _LangPackDifference() {}

type LangPackLanguage interface {
	tl.TLObject
	_LangPackLanguage()
}

var (
	_ LangPackLanguage = (*LangPackLanguagePredict)(nil)
)

type LangPackLanguagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Official        bool     `tl:"official,omitempty:flags:0,implicit"`
	Rtl             bool     `tl:"rtl,omitempty:flags:2,implicit"`
	Beta            bool     `tl:"beta,omitempty:flags:3,implicit"`
	Name            string   `tl:"name"`
	NativeName      string   `tl:"native_name"`
	LangCode        string   `tl:"lang_code"`
	BaseLangCode    *string  `tl:"base_lang_code,omitempty:flags:1"`
	PluralCode      string   `tl:"plural_code"`
	StringsCount    int32    `tl:"strings_count"`
	TranslatedCount int32    `tl:"translated_count"`
	TranslationsURL string   `tl:"translations_url"`
}

func (*LangPackLanguagePredict) CRC() uint32 {
	return 0xeeca5ce3
}
func (*LangPackLanguagePredict) _LangPackLanguage() {}

type LangPackString interface {
	tl.TLObject
	_LangPackString()
}

var (
	_ LangPackString = (*LangPackStringPredict)(nil)
	_ LangPackString = (*LangPackStringPluralizedPredict)(nil)
	_ LangPackString = (*LangPackStringDeletedPredict)(nil)
)

type LangPackStringPredict struct {
	Key   string `tl:"key"`
	Value string `tl:"value"`
}

func (*LangPackStringPredict) CRC() uint32 {
	return 0xcad181f6
}
func (*LangPackStringPredict) _LangPackString() {}

type LangPackStringPluralizedPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Key        string   `tl:"key"`
	ZeroValue  *string  `tl:"zero_value,omitempty:flags:0"`
	OneValue   *string  `tl:"one_value,omitempty:flags:1"`
	TwoValue   *string  `tl:"two_value,omitempty:flags:2"`
	FewValue   *string  `tl:"few_value,omitempty:flags:3"`
	ManyValue  *string  `tl:"many_value,omitempty:flags:4"`
	OtherValue string   `tl:"other_value"`
}

func (*LangPackStringPluralizedPredict) CRC() uint32 {
	return 0x6c47ac9f
}
func (*LangPackStringPluralizedPredict) _LangPackString() {}

type LangPackStringDeletedPredict struct {
	Key string `tl:"key"`
}

func (*LangPackStringDeletedPredict) CRC() uint32 {
	return 0x2979eeb2
}
func (*LangPackStringDeletedPredict) _LangPackString() {}

type MaskCoords interface {
	tl.TLObject
	_MaskCoords()
}

var (
	_ MaskCoords = (*MaskCoordsPredict)(nil)
)

type MaskCoordsPredict struct {
	N    int32   `tl:"n"`
	X    float64 `tl:"x"`
	Y    float64 `tl:"y"`
	Zoom float64 `tl:"zoom"`
}

func (*MaskCoordsPredict) CRC() uint32 {
	return 0xaed6dbb2
}
func (*MaskCoordsPredict) _MaskCoords() {}

type MediaArea interface {
	tl.TLObject
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
	Coordinates MediaAreaCoordinates `tl:"coordinates"`
	Geo         GeoPoint             `tl:"geo"`
	Title       string               `tl:"title"`
	Address     string               `tl:"address"`
	Provider    string               `tl:"provider"`
	VenueID     string               `tl:"venue_id"`
	VenueType   string               `tl:"venue_type"`
}

func (*MediaAreaVenuePredict) CRC() uint32 {
	return 0xbe82db9c
}
func (*MediaAreaVenuePredict) _MediaArea() {}

type InputMediaAreaVenuePredict struct {
	Coordinates MediaAreaCoordinates `tl:"coordinates"`
	QueryID     int64                `tl:"query_id"`
	ResultID    string               `tl:"result_id"`
}

func (*InputMediaAreaVenuePredict) CRC() uint32 {
	return 0xb282217f
}
func (*InputMediaAreaVenuePredict) _MediaArea() {}

type MediaAreaGeoPointPredict struct {
	_           struct{}             `tl:"flags,bitflag"`
	Coordinates MediaAreaCoordinates `tl:"coordinates"`
	Geo         GeoPoint             `tl:"geo"`
	Address     GeoPointAddress      `tl:"address,omitempty:flags:0"`
}

func (*MediaAreaGeoPointPredict) CRC() uint32 {
	return 0xcad5452d
}
func (*MediaAreaGeoPointPredict) _MediaArea() {}

type MediaAreaSuggestedReactionPredict struct {
	_           struct{}             `tl:"flags,bitflag"`
	Dark        bool                 `tl:"dark,omitempty:flags:0,implicit"`
	Flipped     bool                 `tl:"flipped,omitempty:flags:1,implicit"`
	Coordinates MediaAreaCoordinates `tl:"coordinates"`
	Reaction    Reaction             `tl:"reaction"`
}

func (*MediaAreaSuggestedReactionPredict) CRC() uint32 {
	return 0x14455871
}
func (*MediaAreaSuggestedReactionPredict) _MediaArea() {}

type MediaAreaChannelPostPredict struct {
	Coordinates MediaAreaCoordinates `tl:"coordinates"`
	ChannelID   int64                `tl:"channel_id"`
	MsgID       int32                `tl:"msg_id"`
}

func (*MediaAreaChannelPostPredict) CRC() uint32 {
	return 0x770416af
}
func (*MediaAreaChannelPostPredict) _MediaArea() {}

type InputMediaAreaChannelPostPredict struct {
	Coordinates MediaAreaCoordinates `tl:"coordinates"`
	Channel     InputChannel         `tl:"channel"`
	MsgID       int32                `tl:"msg_id"`
}

func (*InputMediaAreaChannelPostPredict) CRC() uint32 {
	return 0x2271f2bf
}
func (*InputMediaAreaChannelPostPredict) _MediaArea() {}

type MediaAreaURLPredict struct {
	Coordinates MediaAreaCoordinates `tl:"coordinates"`
	URL         string               `tl:"url"`
}

func (*MediaAreaURLPredict) CRC() uint32 {
	return 0x37381085
}
func (*MediaAreaURLPredict) _MediaArea() {}

type MediaAreaWeatherPredict struct {
	Coordinates  MediaAreaCoordinates `tl:"coordinates"`
	Emoji        string               `tl:"emoji"`
	TemperatureC float64              `tl:"temperature_c"`
	Color        int32                `tl:"color"`
}

func (*MediaAreaWeatherPredict) CRC() uint32 {
	return 0x49a6549c
}
func (*MediaAreaWeatherPredict) _MediaArea() {}

type MediaAreaCoordinates interface {
	tl.TLObject
	_MediaAreaCoordinates()
}

var (
	_ MediaAreaCoordinates = (*MediaAreaCoordinatesPredict)(nil)
)

type MediaAreaCoordinatesPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	X        float64  `tl:"x"`
	Y        float64  `tl:"y"`
	W        float64  `tl:"w"`
	H        float64  `tl:"h"`
	Rotation float64  `tl:"rotation"`
	Radius   *float64 `tl:"radius,omitempty:flags:0"`
}

func (*MediaAreaCoordinatesPredict) CRC() uint32 {
	return 0xcfc9e002
}
func (*MediaAreaCoordinatesPredict) _MediaAreaCoordinates() {}

type Message interface {
	tl.TLObject
	_Message()
}

var (
	_ Message = (*MessageEmptyPredict)(nil)
	_ Message = (*MessagePredict)(nil)
	_ Message = (*MessageServicePredict)(nil)
)

type MessageEmptyPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	ID     int32    `tl:"id"`
	PeerID Peer     `tl:"peer_id,omitempty:flags:0"`
}

func (*MessageEmptyPredict) CRC() uint32 {
	return 0x90a6ca84
}
func (*MessageEmptyPredict) _Message() {}

type MessagePredict struct {
	_                    struct{}            `tl:"flags,bitflag"`
	Out                  bool                `tl:"out,omitempty:flags:1,implicit"`
	Mentioned            bool                `tl:"mentioned,omitempty:flags:4,implicit"`
	MediaUnread          bool                `tl:"media_unread,omitempty:flags:5,implicit"`
	Silent               bool                `tl:"silent,omitempty:flags:13,implicit"`
	Post                 bool                `tl:"post,omitempty:flags:14,implicit"`
	FromScheduled        bool                `tl:"from_scheduled,omitempty:flags:18,implicit"`
	Legacy               bool                `tl:"legacy,omitempty:flags:19,implicit"`
	EditHide             bool                `tl:"edit_hide,omitempty:flags:21,implicit"`
	Pinned               bool                `tl:"pinned,omitempty:flags:24,implicit"`
	Noforwards           bool                `tl:"noforwards,omitempty:flags:26,implicit"`
	InvertMedia          bool                `tl:"invert_media,omitempty:flags:27,implicit"`
	_                    struct{}            `tl:"flags2,bitflag"`
	Offline              bool                `tl:"offline,omitempty:flags2:1,implicit"`
	ID                   int32               `tl:"id"`
	FromID               Peer                `tl:"from_id,omitempty:flags:8"`
	FromBoostsApplied    *int32              `tl:"from_boosts_applied,omitempty:flags:29"`
	PeerID               Peer                `tl:"peer_id"`
	SavedPeerID          Peer                `tl:"saved_peer_id,omitempty:flags:28"`
	FwdFrom              MessageFwdHeader    `tl:"fwd_from,omitempty:flags:2"`
	ViaBotID             *int64              `tl:"via_bot_id,omitempty:flags:11"`
	ViaBusinessBotID     *int64              `tl:"via_business_bot_id,omitempty:flags2:0"`
	ReplyTo              MessageReplyHeader  `tl:"reply_to,omitempty:flags:3"`
	Date                 int32               `tl:"date"`
	Message              string              `tl:"message"`
	Media                MessageMedia        `tl:"media,omitempty:flags:9"`
	ReplyMarkup          ReplyMarkup         `tl:"reply_markup,omitempty:flags:6"`
	Entities             []MessageEntity     `tl:"entities,omitempty:flags:7"`
	Views                *int32              `tl:"views,omitempty:flags:10"`
	Forwards             *int32              `tl:"forwards,omitempty:flags:10"`
	Replies              MessageReplies      `tl:"replies,omitempty:flags:23"`
	EditDate             *int32              `tl:"edit_date,omitempty:flags:15"`
	PostAuthor           *string             `tl:"post_author,omitempty:flags:16"`
	GroupedID            *int64              `tl:"grouped_id,omitempty:flags:17"`
	Reactions            MessageReactions    `tl:"reactions,omitempty:flags:20"`
	RestrictionReason    []RestrictionReason `tl:"restriction_reason,omitempty:flags:22"`
	TTLPeriod            *int32              `tl:"ttl_period,omitempty:flags:25"`
	QuickReplyShortcutID *int32              `tl:"quick_reply_shortcut_id,omitempty:flags:30"`
	Effect               *int64              `tl:"effect,omitempty:flags2:2"`
	Factcheck            FactCheck           `tl:"factcheck,omitempty:flags2:3"`
}

func (*MessagePredict) CRC() uint32 {
	return 0x94345242
}
func (*MessagePredict) _Message() {}

type MessageServicePredict struct {
	_           struct{}           `tl:"flags,bitflag"`
	Out         bool               `tl:"out,omitempty:flags:1,implicit"`
	Mentioned   bool               `tl:"mentioned,omitempty:flags:4,implicit"`
	MediaUnread bool               `tl:"media_unread,omitempty:flags:5,implicit"`
	Silent      bool               `tl:"silent,omitempty:flags:13,implicit"`
	Post        bool               `tl:"post,omitempty:flags:14,implicit"`
	Legacy      bool               `tl:"legacy,omitempty:flags:19,implicit"`
	ID          int32              `tl:"id"`
	FromID      Peer               `tl:"from_id,omitempty:flags:8"`
	PeerID      Peer               `tl:"peer_id"`
	ReplyTo     MessageReplyHeader `tl:"reply_to,omitempty:flags:3"`
	Date        int32              `tl:"date"`
	Action      MessageAction      `tl:"action"`
	TTLPeriod   *int32             `tl:"ttl_period,omitempty:flags:25"`
}

func (*MessageServicePredict) CRC() uint32 {
	return 0x2b085862
}
func (*MessageServicePredict) _Message() {}

type MessageAction interface {
	tl.TLObject
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
	Title string  `tl:"title"`
	Users []int64 `tl:"users"`
}

func (*MessageActionChatCreatePredict) CRC() uint32 {
	return 0xbd47cbad
}
func (*MessageActionChatCreatePredict) _MessageAction() {}

type MessageActionChatEditTitlePredict struct {
	Title string `tl:"title"`
}

func (*MessageActionChatEditTitlePredict) CRC() uint32 {
	return 0xb5a1ce5a
}
func (*MessageActionChatEditTitlePredict) _MessageAction() {}

type MessageActionChatEditPhotoPredict struct {
	Photo Photo `tl:"photo"`
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
	Users []int64 `tl:"users"`
}

func (*MessageActionChatAddUserPredict) CRC() uint32 {
	return 0x15cefd00
}
func (*MessageActionChatAddUserPredict) _MessageAction() {}

type MessageActionChatDeleteUserPredict struct {
	UserID int64 `tl:"user_id"`
}

func (*MessageActionChatDeleteUserPredict) CRC() uint32 {
	return 0xa43f30cc
}
func (*MessageActionChatDeleteUserPredict) _MessageAction() {}

type MessageActionChatJoinedByLinkPredict struct {
	InviterID int64 `tl:"inviter_id"`
}

func (*MessageActionChatJoinedByLinkPredict) CRC() uint32 {
	return 0x031224c3
}
func (*MessageActionChatJoinedByLinkPredict) _MessageAction() {}

type MessageActionChannelCreatePredict struct {
	Title string `tl:"title"`
}

func (*MessageActionChannelCreatePredict) CRC() uint32 {
	return 0x95d2ac92
}
func (*MessageActionChannelCreatePredict) _MessageAction() {}

type MessageActionChatMigrateToPredict struct {
	ChannelID int64 `tl:"channel_id"`
}

func (*MessageActionChatMigrateToPredict) CRC() uint32 {
	return 0xe1037f92
}
func (*MessageActionChatMigrateToPredict) _MessageAction() {}

type MessageActionChannelMigrateFromPredict struct {
	Title  string `tl:"title"`
	ChatID int64  `tl:"chat_id"`
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
	GameID int64 `tl:"game_id"`
	Score  int32 `tl:"score"`
}

func (*MessageActionGameScorePredict) CRC() uint32 {
	return 0x92a72876
}
func (*MessageActionGameScorePredict) _MessageAction() {}

type MessageActionPaymentSentMePredict struct {
	_                struct{}             `tl:"flags,bitflag"`
	RecurringInit    bool                 `tl:"recurring_init,omitempty:flags:2,implicit"`
	RecurringUsed    bool                 `tl:"recurring_used,omitempty:flags:3,implicit"`
	Currency         string               `tl:"currency"`
	TotalAmount      int64                `tl:"total_amount"`
	Payload          []byte               `tl:"payload"`
	Info             PaymentRequestedInfo `tl:"info,omitempty:flags:0"`
	ShippingOptionID *string              `tl:"shipping_option_id,omitempty:flags:1"`
	Charge           PaymentCharge        `tl:"charge"`
}

func (*MessageActionPaymentSentMePredict) CRC() uint32 {
	return 0x8f31b327
}
func (*MessageActionPaymentSentMePredict) _MessageAction() {}

type MessageActionPaymentSentPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	RecurringInit bool     `tl:"recurring_init,omitempty:flags:2,implicit"`
	RecurringUsed bool     `tl:"recurring_used,omitempty:flags:3,implicit"`
	Currency      string   `tl:"currency"`
	TotalAmount   int64    `tl:"total_amount"`
	InvoiceSlug   *string  `tl:"invoice_slug,omitempty:flags:0"`
}

func (*MessageActionPaymentSentPredict) CRC() uint32 {
	return 0x96163f56
}
func (*MessageActionPaymentSentPredict) _MessageAction() {}

type MessageActionPhoneCallPredict struct {
	_        struct{}               `tl:"flags,bitflag"`
	Video    bool                   `tl:"video,omitempty:flags:2,implicit"`
	CallID   int64                  `tl:"call_id"`
	Reason   PhoneCallDiscardReason `tl:"reason,omitempty:flags:0"`
	Duration *int32                 `tl:"duration,omitempty:flags:1"`
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
	Message string `tl:"message"`
}

func (*MessageActionCustomActionPredict) CRC() uint32 {
	return 0xfae69f56
}
func (*MessageActionCustomActionPredict) _MessageAction() {}

type MessageActionBotAllowedPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	AttachMenu  bool     `tl:"attach_menu,omitempty:flags:1,implicit"`
	FromRequest bool     `tl:"from_request,omitempty:flags:3,implicit"`
	Domain      *string  `tl:"domain,omitempty:flags:0"`
	App         BotApp   `tl:"app,omitempty:flags:2"`
}

func (*MessageActionBotAllowedPredict) CRC() uint32 {
	return 0xc516d679
}
func (*MessageActionBotAllowedPredict) _MessageAction() {}

type MessageActionSecureValuesSentMePredict struct {
	Values      []SecureValue              `tl:"values"`
	Credentials SecureCredentialsEncrypted `tl:"credentials"`
}

func (*MessageActionSecureValuesSentMePredict) CRC() uint32 {
	return 0x1b287353
}
func (*MessageActionSecureValuesSentMePredict) _MessageAction() {}

type MessageActionSecureValuesSentPredict struct {
	Types []SecureValueType `tl:"types"`
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
	FromID   Peer  `tl:"from_id"`
	ToID     Peer  `tl:"to_id"`
	Distance int32 `tl:"distance"`
}

func (*MessageActionGeoProximityReachedPredict) CRC() uint32 {
	return 0x98e0d697
}
func (*MessageActionGeoProximityReachedPredict) _MessageAction() {}

type MessageActionGroupCallPredict struct {
	_        struct{}       `tl:"flags,bitflag"`
	Call     InputGroupCall `tl:"call"`
	Duration *int32         `tl:"duration,omitempty:flags:0"`
}

func (*MessageActionGroupCallPredict) CRC() uint32 {
	return 0x7a0d7f42
}
func (*MessageActionGroupCallPredict) _MessageAction() {}

type MessageActionInviteToGroupCallPredict struct {
	Call  InputGroupCall `tl:"call"`
	Users []int64        `tl:"users"`
}

func (*MessageActionInviteToGroupCallPredict) CRC() uint32 {
	return 0x502f92f7
}
func (*MessageActionInviteToGroupCallPredict) _MessageAction() {}

type MessageActionSetMessagesTTLPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	Period          int32    `tl:"period"`
	AutoSettingFrom *int64   `tl:"auto_setting_from,omitempty:flags:0"`
}

func (*MessageActionSetMessagesTTLPredict) CRC() uint32 {
	return 0x3c134d7b
}
func (*MessageActionSetMessagesTTLPredict) _MessageAction() {}

type MessageActionGroupCallScheduledPredict struct {
	Call         InputGroupCall `tl:"call"`
	ScheduleDate int32          `tl:"schedule_date"`
}

func (*MessageActionGroupCallScheduledPredict) CRC() uint32 {
	return 0xb3a07661
}
func (*MessageActionGroupCallScheduledPredict) _MessageAction() {}

type MessageActionSetChatThemePredict struct {
	Emoticon string `tl:"emoticon"`
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
	Text string `tl:"text"`
	Data string `tl:"data"`
}

func (*MessageActionWebViewDataSentMePredict) CRC() uint32 {
	return 0x47dd8079
}
func (*MessageActionWebViewDataSentMePredict) _MessageAction() {}

type MessageActionWebViewDataSentPredict struct {
	Text string `tl:"text"`
}

func (*MessageActionWebViewDataSentPredict) CRC() uint32 {
	return 0xb4c38cb5
}
func (*MessageActionWebViewDataSentPredict) _MessageAction() {}

type MessageActionGiftPremiumPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Currency       string   `tl:"currency"`
	Amount         int64    `tl:"amount"`
	Months         int32    `tl:"months"`
	CryptoCurrency *string  `tl:"crypto_currency,omitempty:flags:0"`
	CryptoAmount   *int64   `tl:"crypto_amount,omitempty:flags:0"`
}

func (*MessageActionGiftPremiumPredict) CRC() uint32 {
	return 0xc83d6aec
}
func (*MessageActionGiftPremiumPredict) _MessageAction() {}

type MessageActionTopicCreatePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       string   `tl:"title"`
	IconColor   int32    `tl:"icon_color"`
	IconEmojiID *int64   `tl:"icon_emoji_id,omitempty:flags:0"`
}

func (*MessageActionTopicCreatePredict) CRC() uint32 {
	return 0x0d999256
}
func (*MessageActionTopicCreatePredict) _MessageAction() {}

type MessageActionTopicEditPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Title       *string  `tl:"title,omitempty:flags:0"`
	IconEmojiID *int64   `tl:"icon_emoji_id,omitempty:flags:1"`
	Closed      *bool    `tl:"closed,omitempty:flags:2"`
	Hidden      *bool    `tl:"hidden,omitempty:flags:3"`
}

func (*MessageActionTopicEditPredict) CRC() uint32 {
	return 0xc0944820
}
func (*MessageActionTopicEditPredict) _MessageAction() {}

type MessageActionSuggestProfilePhotoPredict struct {
	Photo Photo `tl:"photo"`
}

func (*MessageActionSuggestProfilePhotoPredict) CRC() uint32 {
	return 0x57de635e
}
func (*MessageActionSuggestProfilePhotoPredict) _MessageAction() {}

type MessageActionRequestedPeerPredict struct {
	ButtonID int32  `tl:"button_id"`
	Peers    []Peer `tl:"peers"`
}

func (*MessageActionRequestedPeerPredict) CRC() uint32 {
	return 0x31518e9b
}
func (*MessageActionRequestedPeerPredict) _MessageAction() {}

type MessageActionSetChatWallPaperPredict struct {
	_         struct{}  `tl:"flags,bitflag"`
	Same      bool      `tl:"same,omitempty:flags:0,implicit"`
	ForBoth   bool      `tl:"for_both,omitempty:flags:1,implicit"`
	Wallpaper WallPaper `tl:"wallpaper"`
}

func (*MessageActionSetChatWallPaperPredict) CRC() uint32 {
	return 0x5060a3f4
}
func (*MessageActionSetChatWallPaperPredict) _MessageAction() {}

type MessageActionGiftCodePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ViaGiveaway    bool     `tl:"via_giveaway,omitempty:flags:0,implicit"`
	Unclaimed      bool     `tl:"unclaimed,omitempty:flags:2,implicit"`
	BoostPeer      Peer     `tl:"boost_peer,omitempty:flags:1"`
	Months         int32    `tl:"months"`
	Slug           string   `tl:"slug"`
	Currency       *string  `tl:"currency,omitempty:flags:2"`
	Amount         *int64   `tl:"amount,omitempty:flags:2"`
	CryptoCurrency *string  `tl:"crypto_currency,omitempty:flags:3"`
	CryptoAmount   *int64   `tl:"crypto_amount,omitempty:flags:3"`
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
	WinnersCount   int32 `tl:"winners_count"`
	UnclaimedCount int32 `tl:"unclaimed_count"`
}

func (*MessageActionGiveawayResultsPredict) CRC() uint32 {
	return 0x2a9fadc5
}
func (*MessageActionGiveawayResultsPredict) _MessageAction() {}

type MessageActionBoostApplyPredict struct {
	Boosts int32 `tl:"boosts"`
}

func (*MessageActionBoostApplyPredict) CRC() uint32 {
	return 0xcc02aa6d
}
func (*MessageActionBoostApplyPredict) _MessageAction() {}

type MessageActionRequestedPeerSentMePredict struct {
	ButtonID int32           `tl:"button_id"`
	Peers    []RequestedPeer `tl:"peers"`
}

func (*MessageActionRequestedPeerSentMePredict) CRC() uint32 {
	return 0x93b31848
}
func (*MessageActionRequestedPeerSentMePredict) _MessageAction() {}

type MessageActionPaymentRefundedPredict struct {
	_           struct{}      `tl:"flags,bitflag"`
	Peer        Peer          `tl:"peer"`
	Currency    string        `tl:"currency"`
	TotalAmount int64         `tl:"total_amount"`
	Payload     *[]byte       `tl:"payload,omitempty:flags:0"`
	Charge      PaymentCharge `tl:"charge"`
}

func (*MessageActionPaymentRefundedPredict) CRC() uint32 {
	return 0x41b3e202
}
func (*MessageActionPaymentRefundedPredict) _MessageAction() {}

type MessageActionGiftStarsPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Currency       string   `tl:"currency"`
	Amount         int64    `tl:"amount"`
	Stars          int64    `tl:"stars"`
	CryptoCurrency *string  `tl:"crypto_currency,omitempty:flags:0"`
	CryptoAmount   *int64   `tl:"crypto_amount,omitempty:flags:0"`
	TransactionID  *string  `tl:"transaction_id,omitempty:flags:1"`
}

func (*MessageActionGiftStarsPredict) CRC() uint32 {
	return 0x45d5b021
}
func (*MessageActionGiftStarsPredict) _MessageAction() {}

type MessageEntity interface {
	tl.TLObject
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
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityUnknownPredict) CRC() uint32 {
	return 0xbb92ba95
}
func (*MessageEntityUnknownPredict) _MessageEntity() {}

type MessageEntityMentionPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityMentionPredict) CRC() uint32 {
	return 0xfa04579d
}
func (*MessageEntityMentionPredict) _MessageEntity() {}

type MessageEntityHashtagPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityHashtagPredict) CRC() uint32 {
	return 0x6f635b0d
}
func (*MessageEntityHashtagPredict) _MessageEntity() {}

type MessageEntityBotCommandPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityBotCommandPredict) CRC() uint32 {
	return 0x6cef8ac7
}
func (*MessageEntityBotCommandPredict) _MessageEntity() {}

type MessageEntityURLPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityURLPredict) CRC() uint32 {
	return 0x6ed02538
}
func (*MessageEntityURLPredict) _MessageEntity() {}

type MessageEntityEmailPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityEmailPredict) CRC() uint32 {
	return 0x64e475c2
}
func (*MessageEntityEmailPredict) _MessageEntity() {}

type MessageEntityBoldPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityBoldPredict) CRC() uint32 {
	return 0xbd610bc9
}
func (*MessageEntityBoldPredict) _MessageEntity() {}

type MessageEntityItalicPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityItalicPredict) CRC() uint32 {
	return 0x826f8b60
}
func (*MessageEntityItalicPredict) _MessageEntity() {}

type MessageEntityCodePredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityCodePredict) CRC() uint32 {
	return 0x28a20571
}
func (*MessageEntityCodePredict) _MessageEntity() {}

type MessageEntityPrePredict struct {
	Offset   int32  `tl:"offset"`
	Length   int32  `tl:"length"`
	Language string `tl:"language"`
}

func (*MessageEntityPrePredict) CRC() uint32 {
	return 0x73924be0
}
func (*MessageEntityPrePredict) _MessageEntity() {}

type MessageEntityTextURLPredict struct {
	Offset int32  `tl:"offset"`
	Length int32  `tl:"length"`
	URL    string `tl:"url"`
}

func (*MessageEntityTextURLPredict) CRC() uint32 {
	return 0x76a6d327
}
func (*MessageEntityTextURLPredict) _MessageEntity() {}

type MessageEntityMentionNamePredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
	UserID int64 `tl:"user_id"`
}

func (*MessageEntityMentionNamePredict) CRC() uint32 {
	return 0xdc7b1140
}
func (*MessageEntityMentionNamePredict) _MessageEntity() {}

type InputMessageEntityMentionNamePredict struct {
	Offset int32     `tl:"offset"`
	Length int32     `tl:"length"`
	UserID InputUser `tl:"user_id"`
}

func (*InputMessageEntityMentionNamePredict) CRC() uint32 {
	return 0x208e68c9
}
func (*InputMessageEntityMentionNamePredict) _MessageEntity() {}

type MessageEntityPhonePredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityPhonePredict) CRC() uint32 {
	return 0x9b69e34b
}
func (*MessageEntityPhonePredict) _MessageEntity() {}

type MessageEntityCashtagPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityCashtagPredict) CRC() uint32 {
	return 0x4c4e743f
}
func (*MessageEntityCashtagPredict) _MessageEntity() {}

type MessageEntityUnderlinePredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityUnderlinePredict) CRC() uint32 {
	return 0x9c4e7e8b
}
func (*MessageEntityUnderlinePredict) _MessageEntity() {}

type MessageEntityStrikePredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityStrikePredict) CRC() uint32 {
	return 0xbf0693d4
}
func (*MessageEntityStrikePredict) _MessageEntity() {}

type MessageEntityBankCardPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntityBankCardPredict) CRC() uint32 {
	return 0x761e6af4
}
func (*MessageEntityBankCardPredict) _MessageEntity() {}

type MessageEntitySpoilerPredict struct {
	Offset int32 `tl:"offset"`
	Length int32 `tl:"length"`
}

func (*MessageEntitySpoilerPredict) CRC() uint32 {
	return 0x32ca960f
}
func (*MessageEntitySpoilerPredict) _MessageEntity() {}

type MessageEntityCustomEmojiPredict struct {
	Offset     int32 `tl:"offset"`
	Length     int32 `tl:"length"`
	DocumentID int64 `tl:"document_id"`
}

func (*MessageEntityCustomEmojiPredict) CRC() uint32 {
	return 0xc8cf05f8
}
func (*MessageEntityCustomEmojiPredict) _MessageEntity() {}

type MessageEntityBlockquotePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Collapsed bool     `tl:"collapsed,omitempty:flags:0,implicit"`
	Offset    int32    `tl:"offset"`
	Length    int32    `tl:"length"`
}

func (*MessageEntityBlockquotePredict) CRC() uint32 {
	return 0xf1ccaaac
}
func (*MessageEntityBlockquotePredict) _MessageEntity() {}

type MessageExtendedMedia interface {
	tl.TLObject
	_MessageExtendedMedia()
}

var (
	_ MessageExtendedMedia = (*MessageExtendedMediaPreviewPredict)(nil)
	_ MessageExtendedMedia = (*MessageExtendedMediaPredict)(nil)
)

type MessageExtendedMediaPreviewPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	W             *int32    `tl:"w,omitempty:flags:0"`
	H             *int32    `tl:"h,omitempty:flags:0"`
	Thumb         PhotoSize `tl:"thumb,omitempty:flags:1"`
	VideoDuration *int32    `tl:"video_duration,omitempty:flags:2"`
}

func (*MessageExtendedMediaPreviewPredict) CRC() uint32 {
	return 0xad628cc8
}
func (*MessageExtendedMediaPreviewPredict) _MessageExtendedMedia() {}

type MessageExtendedMediaPredict struct {
	Media MessageMedia `tl:"media"`
}

func (*MessageExtendedMediaPredict) CRC() uint32 {
	return 0xee479c64
}
func (*MessageExtendedMediaPredict) _MessageExtendedMedia() {}

type MessageFwdHeader interface {
	tl.TLObject
	_MessageFwdHeader()
}

var (
	_ MessageFwdHeader = (*MessageFwdHeaderPredict)(nil)
)

type MessageFwdHeaderPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Imported       bool     `tl:"imported,omitempty:flags:7,implicit"`
	SavedOut       bool     `tl:"saved_out,omitempty:flags:11,implicit"`
	FromID         Peer     `tl:"from_id,omitempty:flags:0"`
	FromName       *string  `tl:"from_name,omitempty:flags:5"`
	Date           int32    `tl:"date"`
	ChannelPost    *int32   `tl:"channel_post,omitempty:flags:2"`
	PostAuthor     *string  `tl:"post_author,omitempty:flags:3"`
	SavedFromPeer  Peer     `tl:"saved_from_peer,omitempty:flags:4"`
	SavedFromMsgID *int32   `tl:"saved_from_msg_id,omitempty:flags:4"`
	SavedFromID    Peer     `tl:"saved_from_id,omitempty:flags:8"`
	SavedFromName  *string  `tl:"saved_from_name,omitempty:flags:9"`
	SavedDate      *int32   `tl:"saved_date,omitempty:flags:10"`
	PsaType        *string  `tl:"psa_type,omitempty:flags:6"`
}

func (*MessageFwdHeaderPredict) CRC() uint32 {
	return 0x4e4df4bb
}
func (*MessageFwdHeaderPredict) _MessageFwdHeader() {}

type MessageMedia interface {
	tl.TLObject
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
	Spoiler    bool     `tl:"spoiler,omitempty:flags:3,implicit"`
	Photo      Photo    `tl:"photo,omitempty:flags:0"`
	TTLSeconds *int32   `tl:"ttl_seconds,omitempty:flags:2"`
}

func (*MessageMediaPhotoPredict) CRC() uint32 {
	return 0x695150d7
}
func (*MessageMediaPhotoPredict) _MessageMedia() {}

type MessageMediaGeoPredict struct {
	Geo GeoPoint `tl:"geo"`
}

func (*MessageMediaGeoPredict) CRC() uint32 {
	return 0x56e0d474
}
func (*MessageMediaGeoPredict) _MessageMedia() {}

type MessageMediaContactPredict struct {
	PhoneNumber string `tl:"phone_number"`
	FirstName   string `tl:"first_name"`
	LastName    string `tl:"last_name"`
	Vcard       string `tl:"vcard"`
	UserID      int64  `tl:"user_id"`
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
	Nopremium   bool     `tl:"nopremium,omitempty:flags:3,implicit"`
	Spoiler     bool     `tl:"spoiler,omitempty:flags:4,implicit"`
	Video       bool     `tl:"video,omitempty:flags:6,implicit"`
	Round       bool     `tl:"round,omitempty:flags:7,implicit"`
	Voice       bool     `tl:"voice,omitempty:flags:8,implicit"`
	Document    Document `tl:"document,omitempty:flags:0"`
	AltDocument Document `tl:"alt_document,omitempty:flags:5"`
	TTLSeconds  *int32   `tl:"ttl_seconds,omitempty:flags:2"`
}

func (*MessageMediaDocumentPredict) CRC() uint32 {
	return 0x4cf4d72d
}
func (*MessageMediaDocumentPredict) _MessageMedia() {}

type MessageMediaWebPagePredict struct {
	_               struct{} `tl:"flags,bitflag"`
	ForceLargeMedia bool     `tl:"force_large_media,omitempty:flags:0,implicit"`
	ForceSmallMedia bool     `tl:"force_small_media,omitempty:flags:1,implicit"`
	Manual          bool     `tl:"manual,omitempty:flags:3,implicit"`
	Safe            bool     `tl:"safe,omitempty:flags:4,implicit"`
	Webpage         WebPage  `tl:"webpage"`
}

func (*MessageMediaWebPagePredict) CRC() uint32 {
	return 0xddf10c3b
}
func (*MessageMediaWebPagePredict) _MessageMedia() {}

type MessageMediaVenuePredict struct {
	Geo       GeoPoint `tl:"geo"`
	Title     string   `tl:"title"`
	Address   string   `tl:"address"`
	Provider  string   `tl:"provider"`
	VenueID   string   `tl:"venue_id"`
	VenueType string   `tl:"venue_type"`
}

func (*MessageMediaVenuePredict) CRC() uint32 {
	return 0x2ec0533f
}
func (*MessageMediaVenuePredict) _MessageMedia() {}

type MessageMediaGamePredict struct {
	Game Game `tl:"game"`
}

func (*MessageMediaGamePredict) CRC() uint32 {
	return 0xfdb19008
}
func (*MessageMediaGamePredict) _MessageMedia() {}

type MessageMediaInvoicePredict struct {
	_                        struct{}             `tl:"flags,bitflag"`
	ShippingAddressRequested bool                 `tl:"shipping_address_requested,omitempty:flags:1,implicit"`
	Test                     bool                 `tl:"test,omitempty:flags:3,implicit"`
	Title                    string               `tl:"title"`
	Description              string               `tl:"description"`
	Photo                    WebDocument          `tl:"photo,omitempty:flags:0"`
	ReceiptMsgID             *int32               `tl:"receipt_msg_id,omitempty:flags:2"`
	Currency                 string               `tl:"currency"`
	TotalAmount              int64                `tl:"total_amount"`
	StartParam               string               `tl:"start_param"`
	ExtendedMedia            MessageExtendedMedia `tl:"extended_media,omitempty:flags:4"`
}

func (*MessageMediaInvoicePredict) CRC() uint32 {
	return 0xf6a548d3
}
func (*MessageMediaInvoicePredict) _MessageMedia() {}

type MessageMediaGeoLivePredict struct {
	_                           struct{} `tl:"flags,bitflag"`
	Geo                         GeoPoint `tl:"geo"`
	Heading                     *int32   `tl:"heading,omitempty:flags:0"`
	Period                      int32    `tl:"period"`
	ProximityNotificationRadius *int32   `tl:"proximity_notification_radius,omitempty:flags:1"`
}

func (*MessageMediaGeoLivePredict) CRC() uint32 {
	return 0xb940c666
}
func (*MessageMediaGeoLivePredict) _MessageMedia() {}

type MessageMediaPollPredict struct {
	Poll    Poll        `tl:"poll"`
	Results PollResults `tl:"results"`
}

func (*MessageMediaPollPredict) CRC() uint32 {
	return 0x4bd6e798
}
func (*MessageMediaPollPredict) _MessageMedia() {}

type MessageMediaDicePredict struct {
	Value    int32  `tl:"value"`
	Emoticon string `tl:"emoticon"`
}

func (*MessageMediaDicePredict) CRC() uint32 {
	return 0x3f7ee58b
}
func (*MessageMediaDicePredict) _MessageMedia() {}

type MessageMediaStoryPredict struct {
	_          struct{}  `tl:"flags,bitflag"`
	ViaMention bool      `tl:"via_mention,omitempty:flags:1,implicit"`
	Peer       Peer      `tl:"peer"`
	ID         int32     `tl:"id"`
	Story      StoryItem `tl:"story,omitempty:flags:0"`
}

func (*MessageMediaStoryPredict) CRC() uint32 {
	return 0x68cb6283
}
func (*MessageMediaStoryPredict) _MessageMedia() {}

type MessageMediaGiveawayPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers bool     `tl:"only_new_subscribers,omitempty:flags:0,implicit"`
	WinnersAreVisible  bool     `tl:"winners_are_visible,omitempty:flags:2,implicit"`
	Channels           []int64  `tl:"channels"`
	CountriesIso2      []string `tl:"countries_iso2,omitempty:flags:1"`
	PrizeDescription   *string  `tl:"prize_description,omitempty:flags:3"`
	Quantity           int32    `tl:"quantity"`
	Months             int32    `tl:"months"`
	UntilDate          int32    `tl:"until_date"`
}

func (*MessageMediaGiveawayPredict) CRC() uint32 {
	return 0xdaad85b0
}
func (*MessageMediaGiveawayPredict) _MessageMedia() {}

type MessageMediaGiveawayResultsPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	OnlyNewSubscribers   bool     `tl:"only_new_subscribers,omitempty:flags:0,implicit"`
	Refunded             bool     `tl:"refunded,omitempty:flags:2,implicit"`
	ChannelID            int64    `tl:"channel_id"`
	AdditionalPeersCount *int32   `tl:"additional_peers_count,omitempty:flags:3"`
	LaunchMsgID          int32    `tl:"launch_msg_id"`
	WinnersCount         int32    `tl:"winners_count"`
	UnclaimedCount       int32    `tl:"unclaimed_count"`
	Winners              []int64  `tl:"winners"`
	Months               int32    `tl:"months"`
	PrizeDescription     *string  `tl:"prize_description,omitempty:flags:1"`
	UntilDate            int32    `tl:"until_date"`
}

func (*MessageMediaGiveawayResultsPredict) CRC() uint32 {
	return 0xc6991068
}
func (*MessageMediaGiveawayResultsPredict) _MessageMedia() {}

type MessageMediaPaidMediaPredict struct {
	StarsAmount   int64                  `tl:"stars_amount"`
	ExtendedMedia []MessageExtendedMedia `tl:"extended_media"`
}

func (*MessageMediaPaidMediaPredict) CRC() uint32 {
	return 0xa8852491
}
func (*MessageMediaPaidMediaPredict) _MessageMedia() {}

type MessagePeerReaction interface {
	tl.TLObject
	_MessagePeerReaction()
}

var (
	_ MessagePeerReaction = (*MessagePeerReactionPredict)(nil)
)

type MessagePeerReactionPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Big      bool     `tl:"big,omitempty:flags:0,implicit"`
	Unread   bool     `tl:"unread,omitempty:flags:1,implicit"`
	My       bool     `tl:"my,omitempty:flags:2,implicit"`
	PeerID   Peer     `tl:"peer_id"`
	Date     int32    `tl:"date"`
	Reaction Reaction `tl:"reaction"`
}

func (*MessagePeerReactionPredict) CRC() uint32 {
	return 0x8c79b63c
}
func (*MessagePeerReactionPredict) _MessagePeerReaction() {}

type MessagePeerVote interface {
	tl.TLObject
	_MessagePeerVote()
}

var (
	_ MessagePeerVote = (*MessagePeerVotePredict)(nil)
	_ MessagePeerVote = (*MessagePeerVoteInputOptionPredict)(nil)
	_ MessagePeerVote = (*MessagePeerVoteMultiplePredict)(nil)
)

type MessagePeerVotePredict struct {
	Peer   Peer   `tl:"peer"`
	Option []byte `tl:"option"`
	Date   int32  `tl:"date"`
}

func (*MessagePeerVotePredict) CRC() uint32 {
	return 0xb6cc2d5c
}
func (*MessagePeerVotePredict) _MessagePeerVote() {}

type MessagePeerVoteInputOptionPredict struct {
	Peer Peer  `tl:"peer"`
	Date int32 `tl:"date"`
}

func (*MessagePeerVoteInputOptionPredict) CRC() uint32 {
	return 0x74cda504
}
func (*MessagePeerVoteInputOptionPredict) _MessagePeerVote() {}

type MessagePeerVoteMultiplePredict struct {
	Peer    Peer     `tl:"peer"`
	Options [][]byte `tl:"options"`
	Date    int32    `tl:"date"`
}

func (*MessagePeerVoteMultiplePredict) CRC() uint32 {
	return 0x4628f6e6
}
func (*MessagePeerVoteMultiplePredict) _MessagePeerVote() {}

type MessageRange interface {
	tl.TLObject
	_MessageRange()
}

var (
	_ MessageRange = (*MessageRangePredict)(nil)
)

type MessageRangePredict struct {
	MinID int32 `tl:"min_id"`
	MaxID int32 `tl:"max_id"`
}

func (*MessageRangePredict) CRC() uint32 {
	return 0x0ae30253
}
func (*MessageRangePredict) _MessageRange() {}

type MessageReactions interface {
	tl.TLObject
	_MessageReactions()
}

var (
	_ MessageReactions = (*MessageReactionsPredict)(nil)
)

type MessageReactionsPredict struct {
	_               struct{}              `tl:"flags,bitflag"`
	Min             bool                  `tl:"min,omitempty:flags:0,implicit"`
	CanSeeList      bool                  `tl:"can_see_list,omitempty:flags:2,implicit"`
	ReactionsAsTags bool                  `tl:"reactions_as_tags,omitempty:flags:3,implicit"`
	Results         []ReactionCount       `tl:"results"`
	RecentReactions []MessagePeerReaction `tl:"recent_reactions,omitempty:flags:1"`
}

func (*MessageReactionsPredict) CRC() uint32 {
	return 0x4f2b9479
}
func (*MessageReactionsPredict) _MessageReactions() {}

type MessageReplies interface {
	tl.TLObject
	_MessageReplies()
}

var (
	_ MessageReplies = (*MessageRepliesPredict)(nil)
)

type MessageRepliesPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Comments       bool     `tl:"comments,omitempty:flags:0,implicit"`
	Replies        int32    `tl:"replies"`
	RepliesPts     int32    `tl:"replies_pts"`
	RecentRepliers []Peer   `tl:"recent_repliers,omitempty:flags:1"`
	ChannelID      *int64   `tl:"channel_id,omitempty:flags:0"`
	MaxID          *int32   `tl:"max_id,omitempty:flags:2"`
	ReadMaxID      *int32   `tl:"read_max_id,omitempty:flags:3"`
}

func (*MessageRepliesPredict) CRC() uint32 {
	return 0x83d60fc2
}
func (*MessageRepliesPredict) _MessageReplies() {}

type MessageReplyHeader interface {
	tl.TLObject
	_MessageReplyHeader()
}

var (
	_ MessageReplyHeader = (*MessageReplyHeaderPredict)(nil)
	_ MessageReplyHeader = (*MessageReplyStoryHeaderPredict)(nil)
)

type MessageReplyHeaderPredict struct {
	_                struct{}         `tl:"flags,bitflag"`
	ReplyToScheduled bool             `tl:"reply_to_scheduled,omitempty:flags:2,implicit"`
	ForumTopic       bool             `tl:"forum_topic,omitempty:flags:3,implicit"`
	Quote            bool             `tl:"quote,omitempty:flags:9,implicit"`
	ReplyToMsgID     *int32           `tl:"reply_to_msg_id,omitempty:flags:4"`
	ReplyToPeerID    Peer             `tl:"reply_to_peer_id,omitempty:flags:0"`
	ReplyFrom        MessageFwdHeader `tl:"reply_from,omitempty:flags:5"`
	ReplyMedia       MessageMedia     `tl:"reply_media,omitempty:flags:8"`
	ReplyToTopID     *int32           `tl:"reply_to_top_id,omitempty:flags:1"`
	QuoteText        *string          `tl:"quote_text,omitempty:flags:6"`
	QuoteEntities    []MessageEntity  `tl:"quote_entities,omitempty:flags:7"`
	QuoteOffset      *int32           `tl:"quote_offset,omitempty:flags:10"`
}

func (*MessageReplyHeaderPredict) CRC() uint32 {
	return 0xafbc09db
}
func (*MessageReplyHeaderPredict) _MessageReplyHeader() {}

type MessageReplyStoryHeaderPredict struct {
	Peer    Peer  `tl:"peer"`
	StoryID int32 `tl:"story_id"`
}

func (*MessageReplyStoryHeaderPredict) CRC() uint32 {
	return 0x0e5af939
}
func (*MessageReplyStoryHeaderPredict) _MessageReplyHeader() {}

type MessageViews interface {
	tl.TLObject
	_MessageViews()
}

var (
	_ MessageViews = (*MessageViewsPredict)(nil)
)

type MessageViewsPredict struct {
	_        struct{}       `tl:"flags,bitflag"`
	Views    *int32         `tl:"views,omitempty:flags:0"`
	Forwards *int32         `tl:"forwards,omitempty:flags:1"`
	Replies  MessageReplies `tl:"replies,omitempty:flags:2"`
}

func (*MessageViewsPredict) CRC() uint32 {
	return 0x455b853d
}
func (*MessageViewsPredict) _MessageViews() {}

type MessagesFilter interface {
	tl.TLObject
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
	Missed bool     `tl:"missed,omitempty:flags:0,implicit"`
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
	tl.TLObject
	_MissingInvitee()
}

var (
	_ MissingInvitee = (*MissingInviteePredict)(nil)
)

type MissingInviteePredict struct {
	_                       struct{} `tl:"flags,bitflag"`
	PremiumWouldAllowInvite bool     `tl:"premium_would_allow_invite,omitempty:flags:0,implicit"`
	PremiumRequiredForPm    bool     `tl:"premium_required_for_pm,omitempty:flags:1,implicit"`
	UserID                  int64    `tl:"user_id"`
}

func (*MissingInviteePredict) CRC() uint32 {
	return 0x628c9224
}
func (*MissingInviteePredict) _MissingInvitee() {}

type MyBoost interface {
	tl.TLObject
	_MyBoost()
}

var (
	_ MyBoost = (*MyBoostPredict)(nil)
)

type MyBoostPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Slot              int32    `tl:"slot"`
	Peer              Peer     `tl:"peer,omitempty:flags:0"`
	Date              int32    `tl:"date"`
	Expires           int32    `tl:"expires"`
	CooldownUntilDate *int32   `tl:"cooldown_until_date,omitempty:flags:1"`
}

func (*MyBoostPredict) CRC() uint32 {
	return 0xc448415c
}
func (*MyBoostPredict) _MyBoost() {}

type NearestDc interface {
	tl.TLObject
	_NearestDc()
}

var (
	_ NearestDc = (*NearestDcPredict)(nil)
)

type NearestDcPredict struct {
	Country   string `tl:"country"`
	ThisDc    int32  `tl:"this_dc"`
	NearestDc int32  `tl:"nearest_dc"`
}

func (*NearestDcPredict) CRC() uint32 {
	return 0x8e1a1775
}
func (*NearestDcPredict) _NearestDc() {}

type NotificationSound interface {
	tl.TLObject
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
	Title string `tl:"title"`
	Data  string `tl:"data"`
}

func (*NotificationSoundLocalPredict) CRC() uint32 {
	return 0x830b9ae4
}
func (*NotificationSoundLocalPredict) _NotificationSound() {}

type NotificationSoundRingtonePredict struct {
	ID int64 `tl:"id"`
}

func (*NotificationSoundRingtonePredict) CRC() uint32 {
	return 0xff6c8049
}
func (*NotificationSoundRingtonePredict) _NotificationSound() {}

type NotifyPeer interface {
	tl.TLObject
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
	Peer Peer `tl:"peer"`
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
	Peer     Peer  `tl:"peer"`
	TopMsgID int32 `tl:"top_msg_id"`
}

func (*NotifyForumTopicPredict) CRC() uint32 {
	return 0x226e6308
}
func (*NotifyForumTopicPredict) _NotifyPeer() {}

type Null interface {
	tl.TLObject
	_Null()
}

var (
	_ Null = (*NullPredict)(nil)
)

type NullPredict struct{}

func (*NullPredict) CRC() uint32 {
	return 0x56730bcc
}
func (*NullPredict) _Null() {}

type OutboxReadDate interface {
	tl.TLObject
	_OutboxReadDate()
}

var (
	_ OutboxReadDate = (*OutboxReadDatePredict)(nil)
)

type OutboxReadDatePredict struct {
	Date int32 `tl:"date"`
}

func (*OutboxReadDatePredict) CRC() uint32 {
	return 0x3bb842ac
}
func (*OutboxReadDatePredict) _OutboxReadDate() {}

type Page interface {
	tl.TLObject
	_Page()
}

var (
	_ Page = (*PagePredict)(nil)
)

type PagePredict struct {
	_         struct{}    `tl:"flags,bitflag"`
	Part      bool        `tl:"part,omitempty:flags:0,implicit"`
	Rtl       bool        `tl:"rtl,omitempty:flags:1,implicit"`
	V2        bool        `tl:"v2,omitempty:flags:2,implicit"`
	URL       string      `tl:"url"`
	Blocks    []PageBlock `tl:"blocks"`
	Photos    []Photo     `tl:"photos"`
	Documents []Document  `tl:"documents"`
	Views     *int32      `tl:"views,omitempty:flags:3"`
}

func (*PagePredict) CRC() uint32 {
	return 0x98657f0d
}
func (*PagePredict) _Page() {}

type PageBlock interface {
	tl.TLObject
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
	Text RichText `tl:"text"`
}

func (*PageBlockTitlePredict) CRC() uint32 {
	return 0x70abc3fd
}
func (*PageBlockTitlePredict) _PageBlock() {}

type PageBlockSubtitlePredict struct {
	Text RichText `tl:"text"`
}

func (*PageBlockSubtitlePredict) CRC() uint32 {
	return 0x8ffa9a1f
}
func (*PageBlockSubtitlePredict) _PageBlock() {}

type PageBlockAuthorDatePredict struct {
	Author        RichText `tl:"author"`
	PublishedDate int32    `tl:"published_date"`
}

func (*PageBlockAuthorDatePredict) CRC() uint32 {
	return 0xbaafe5e0
}
func (*PageBlockAuthorDatePredict) _PageBlock() {}

type PageBlockHeaderPredict struct {
	Text RichText `tl:"text"`
}

func (*PageBlockHeaderPredict) CRC() uint32 {
	return 0xbfd064ec
}
func (*PageBlockHeaderPredict) _PageBlock() {}

type PageBlockSubheaderPredict struct {
	Text RichText `tl:"text"`
}

func (*PageBlockSubheaderPredict) CRC() uint32 {
	return 0xf12bb6e1
}
func (*PageBlockSubheaderPredict) _PageBlock() {}

type PageBlockParagraphPredict struct {
	Text RichText `tl:"text"`
}

func (*PageBlockParagraphPredict) CRC() uint32 {
	return 0x467a0766
}
func (*PageBlockParagraphPredict) _PageBlock() {}

type PageBlockPreformattedPredict struct {
	Text     RichText `tl:"text"`
	Language string   `tl:"language"`
}

func (*PageBlockPreformattedPredict) CRC() uint32 {
	return 0xc070d93e
}
func (*PageBlockPreformattedPredict) _PageBlock() {}

type PageBlockFooterPredict struct {
	Text RichText `tl:"text"`
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
	Name string `tl:"name"`
}

func (*PageBlockAnchorPredict) CRC() uint32 {
	return 0xce0d37b0
}
func (*PageBlockAnchorPredict) _PageBlock() {}

type PageBlockListPredict struct {
	Items []PageListItem `tl:"items"`
}

func (*PageBlockListPredict) CRC() uint32 {
	return 0xe4e88011
}
func (*PageBlockListPredict) _PageBlock() {}

type PageBlockBlockquotePredict struct {
	Text    RichText `tl:"text"`
	Caption RichText `tl:"caption"`
}

func (*PageBlockBlockquotePredict) CRC() uint32 {
	return 0x263d7c26
}
func (*PageBlockBlockquotePredict) _PageBlock() {}

type PageBlockPullquotePredict struct {
	Text    RichText `tl:"text"`
	Caption RichText `tl:"caption"`
}

func (*PageBlockPullquotePredict) CRC() uint32 {
	return 0x4f4456d3
}
func (*PageBlockPullquotePredict) _PageBlock() {}

type PageBlockPhotoPredict struct {
	_         struct{}    `tl:"flags,bitflag"`
	PhotoID   int64       `tl:"photo_id"`
	Caption   PageCaption `tl:"caption"`
	URL       *string     `tl:"url,omitempty:flags:0"`
	WebpageID *int64      `tl:"webpage_id,omitempty:flags:0"`
}

func (*PageBlockPhotoPredict) CRC() uint32 {
	return 0x1759c560
}
func (*PageBlockPhotoPredict) _PageBlock() {}

type PageBlockVideoPredict struct {
	_        struct{}    `tl:"flags,bitflag"`
	Autoplay bool        `tl:"autoplay,omitempty:flags:0,implicit"`
	Loop     bool        `tl:"loop,omitempty:flags:1,implicit"`
	VideoID  int64       `tl:"video_id"`
	Caption  PageCaption `tl:"caption"`
}

func (*PageBlockVideoPredict) CRC() uint32 {
	return 0x7c8fe7b6
}
func (*PageBlockVideoPredict) _PageBlock() {}

type PageBlockCoverPredict struct {
	Cover PageBlock `tl:"cover"`
}

func (*PageBlockCoverPredict) CRC() uint32 {
	return 0x39f23300
}
func (*PageBlockCoverPredict) _PageBlock() {}

type PageBlockEmbedPredict struct {
	_              struct{}    `tl:"flags,bitflag"`
	FullWidth      bool        `tl:"full_width,omitempty:flags:0,implicit"`
	AllowScrolling bool        `tl:"allow_scrolling,omitempty:flags:3,implicit"`
	URL            *string     `tl:"url,omitempty:flags:1"`
	Html           *string     `tl:"html,omitempty:flags:2"`
	PosterPhotoID  *int64      `tl:"poster_photo_id,omitempty:flags:4"`
	W              *int32      `tl:"w,omitempty:flags:5"`
	H              *int32      `tl:"h,omitempty:flags:5"`
	Caption        PageCaption `tl:"caption"`
}

func (*PageBlockEmbedPredict) CRC() uint32 {
	return 0xa8718dc5
}
func (*PageBlockEmbedPredict) _PageBlock() {}

type PageBlockEmbedPostPredict struct {
	URL           string      `tl:"url"`
	WebpageID     int64       `tl:"webpage_id"`
	AuthorPhotoID int64       `tl:"author_photo_id"`
	Author        string      `tl:"author"`
	Date          int32       `tl:"date"`
	Blocks        []PageBlock `tl:"blocks"`
	Caption       PageCaption `tl:"caption"`
}

func (*PageBlockEmbedPostPredict) CRC() uint32 {
	return 0xf259a80b
}
func (*PageBlockEmbedPostPredict) _PageBlock() {}

type PageBlockCollagePredict struct {
	Items   []PageBlock `tl:"items"`
	Caption PageCaption `tl:"caption"`
}

func (*PageBlockCollagePredict) CRC() uint32 {
	return 0x65a0fa4d
}
func (*PageBlockCollagePredict) _PageBlock() {}

type PageBlockSlideshowPredict struct {
	Items   []PageBlock `tl:"items"`
	Caption PageCaption `tl:"caption"`
}

func (*PageBlockSlideshowPredict) CRC() uint32 {
	return 0x031f9590
}
func (*PageBlockSlideshowPredict) _PageBlock() {}

type PageBlockChannelPredict struct {
	Channel Chat `tl:"channel"`
}

func (*PageBlockChannelPredict) CRC() uint32 {
	return 0xef1751b5
}
func (*PageBlockChannelPredict) _PageBlock() {}

type PageBlockAudioPredict struct {
	AudioID int64       `tl:"audio_id"`
	Caption PageCaption `tl:"caption"`
}

func (*PageBlockAudioPredict) CRC() uint32 {
	return 0x804361ea
}
func (*PageBlockAudioPredict) _PageBlock() {}

type PageBlockKickerPredict struct {
	Text RichText `tl:"text"`
}

func (*PageBlockKickerPredict) CRC() uint32 {
	return 0x1e148390
}
func (*PageBlockKickerPredict) _PageBlock() {}

type PageBlockTablePredict struct {
	_        struct{}       `tl:"flags,bitflag"`
	Bordered bool           `tl:"bordered,omitempty:flags:0,implicit"`
	Striped  bool           `tl:"striped,omitempty:flags:1,implicit"`
	Title    RichText       `tl:"title"`
	Rows     []PageTableRow `tl:"rows"`
}

func (*PageBlockTablePredict) CRC() uint32 {
	return 0xbf4dea82
}
func (*PageBlockTablePredict) _PageBlock() {}

type PageBlockOrderedListPredict struct {
	Items []PageListOrderedItem `tl:"items"`
}

func (*PageBlockOrderedListPredict) CRC() uint32 {
	return 0x9a8ae1e1
}
func (*PageBlockOrderedListPredict) _PageBlock() {}

type PageBlockDetailsPredict struct {
	_      struct{}    `tl:"flags,bitflag"`
	Open   bool        `tl:"open,omitempty:flags:0,implicit"`
	Blocks []PageBlock `tl:"blocks"`
	Title  RichText    `tl:"title"`
}

func (*PageBlockDetailsPredict) CRC() uint32 {
	return 0x76768bed
}
func (*PageBlockDetailsPredict) _PageBlock() {}

type PageBlockRelatedArticlesPredict struct {
	Title    RichText             `tl:"title"`
	Articles []PageRelatedArticle `tl:"articles"`
}

func (*PageBlockRelatedArticlesPredict) CRC() uint32 {
	return 0x16115a96
}
func (*PageBlockRelatedArticlesPredict) _PageBlock() {}

type PageBlockMapPredict struct {
	Geo     GeoPoint    `tl:"geo"`
	Zoom    int32       `tl:"zoom"`
	W       int32       `tl:"w"`
	H       int32       `tl:"h"`
	Caption PageCaption `tl:"caption"`
}

func (*PageBlockMapPredict) CRC() uint32 {
	return 0xa44f3ef6
}
func (*PageBlockMapPredict) _PageBlock() {}

type PageCaption interface {
	tl.TLObject
	_PageCaption()
}

var (
	_ PageCaption = (*PageCaptionPredict)(nil)
)

type PageCaptionPredict struct {
	Text   RichText `tl:"text"`
	Credit RichText `tl:"credit"`
}

func (*PageCaptionPredict) CRC() uint32 {
	return 0x6f747657
}
func (*PageCaptionPredict) _PageCaption() {}

type PageListItem interface {
	tl.TLObject
	_PageListItem()
}

var (
	_ PageListItem = (*PageListItemTextPredict)(nil)
	_ PageListItem = (*PageListItemBlocksPredict)(nil)
)

type PageListItemTextPredict struct {
	Text RichText `tl:"text"`
}

func (*PageListItemTextPredict) CRC() uint32 {
	return 0xb92fb6cd
}
func (*PageListItemTextPredict) _PageListItem() {}

type PageListItemBlocksPredict struct {
	Blocks []PageBlock `tl:"blocks"`
}

func (*PageListItemBlocksPredict) CRC() uint32 {
	return 0x25e073fc
}
func (*PageListItemBlocksPredict) _PageListItem() {}

type PageListOrderedItem interface {
	tl.TLObject
	_PageListOrderedItem()
}

var (
	_ PageListOrderedItem = (*PageListOrderedItemTextPredict)(nil)
	_ PageListOrderedItem = (*PageListOrderedItemBlocksPredict)(nil)
)

type PageListOrderedItemTextPredict struct {
	Num  string   `tl:"num"`
	Text RichText `tl:"text"`
}

func (*PageListOrderedItemTextPredict) CRC() uint32 {
	return 0x5e068047
}
func (*PageListOrderedItemTextPredict) _PageListOrderedItem() {}

type PageListOrderedItemBlocksPredict struct {
	Num    string      `tl:"num"`
	Blocks []PageBlock `tl:"blocks"`
}

func (*PageListOrderedItemBlocksPredict) CRC() uint32 {
	return 0x98dd8936
}
func (*PageListOrderedItemBlocksPredict) _PageListOrderedItem() {}

type PageRelatedArticle interface {
	tl.TLObject
	_PageRelatedArticle()
}

var (
	_ PageRelatedArticle = (*PageRelatedArticlePredict)(nil)
)

type PageRelatedArticlePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	URL           string   `tl:"url"`
	WebpageID     int64    `tl:"webpage_id"`
	Title         *string  `tl:"title,omitempty:flags:0"`
	Description   *string  `tl:"description,omitempty:flags:1"`
	PhotoID       *int64   `tl:"photo_id,omitempty:flags:2"`
	Author        *string  `tl:"author,omitempty:flags:3"`
	PublishedDate *int32   `tl:"published_date,omitempty:flags:4"`
}

func (*PageRelatedArticlePredict) CRC() uint32 {
	return 0xb390dc08
}
func (*PageRelatedArticlePredict) _PageRelatedArticle() {}

type PageTableCell interface {
	tl.TLObject
	_PageTableCell()
}

var (
	_ PageTableCell = (*PageTableCellPredict)(nil)
)

type PageTableCellPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Header       bool     `tl:"header,omitempty:flags:0,implicit"`
	AlignCenter  bool     `tl:"align_center,omitempty:flags:3,implicit"`
	AlignRight   bool     `tl:"align_right,omitempty:flags:4,implicit"`
	ValignMiddle bool     `tl:"valign_middle,omitempty:flags:5,implicit"`
	ValignBottom bool     `tl:"valign_bottom,omitempty:flags:6,implicit"`
	Text         RichText `tl:"text,omitempty:flags:7"`
	Colspan      *int32   `tl:"colspan,omitempty:flags:1"`
	Rowspan      *int32   `tl:"rowspan,omitempty:flags:2"`
}

func (*PageTableCellPredict) CRC() uint32 {
	return 0x34566b6a
}
func (*PageTableCellPredict) _PageTableCell() {}

type PageTableRow interface {
	tl.TLObject
	_PageTableRow()
}

var (
	_ PageTableRow = (*PageTableRowPredict)(nil)
)

type PageTableRowPredict struct {
	Cells []PageTableCell `tl:"cells"`
}

func (*PageTableRowPredict) CRC() uint32 {
	return 0xe0c0c5e5
}
func (*PageTableRowPredict) _PageTableRow() {}

type PasswordKdfAlgo interface {
	tl.TLObject
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
	Salt1 []byte `tl:"salt1"`
	Salt2 []byte `tl:"salt2"`
	G     int32  `tl:"g"`
	P     []byte `tl:"p"`
}

func (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPowPredict) CRC() uint32 {
	return 0x3a912d4a
}
func (*PasswordKdfAlgoSHA256SHA256Pbkdf2Hmacsha512Iter100000SHA256ModPowPredict) _PasswordKdfAlgo() {}

type PaymentCharge interface {
	tl.TLObject
	_PaymentCharge()
}

var (
	_ PaymentCharge = (*PaymentChargePredict)(nil)
)

type PaymentChargePredict struct {
	ID               string `tl:"id"`
	ProviderChargeID string `tl:"provider_charge_id"`
}

func (*PaymentChargePredict) CRC() uint32 {
	return 0xea02c27e
}
func (*PaymentChargePredict) _PaymentCharge() {}

type PaymentFormMethod interface {
	tl.TLObject
	_PaymentFormMethod()
}

var (
	_ PaymentFormMethod = (*PaymentFormMethodPredict)(nil)
)

type PaymentFormMethodPredict struct {
	URL   string `tl:"url"`
	Title string `tl:"title"`
}

func (*PaymentFormMethodPredict) CRC() uint32 {
	return 0x88f8f21b
}
func (*PaymentFormMethodPredict) _PaymentFormMethod() {}

type PaymentRequestedInfo interface {
	tl.TLObject
	_PaymentRequestedInfo()
}

var (
	_ PaymentRequestedInfo = (*PaymentRequestedInfoPredict)(nil)
)

type PaymentRequestedInfoPredict struct {
	_               struct{}    `tl:"flags,bitflag"`
	Name            *string     `tl:"name,omitempty:flags:0"`
	Phone           *string     `tl:"phone,omitempty:flags:1"`
	Email           *string     `tl:"email,omitempty:flags:2"`
	ShippingAddress PostAddress `tl:"shipping_address,omitempty:flags:3"`
}

func (*PaymentRequestedInfoPredict) CRC() uint32 {
	return 0x909c3f94
}
func (*PaymentRequestedInfoPredict) _PaymentRequestedInfo() {}

type PaymentSavedCredentials interface {
	tl.TLObject
	_PaymentSavedCredentials()
}

var (
	_ PaymentSavedCredentials = (*PaymentSavedCredentialsCardPredict)(nil)
)

type PaymentSavedCredentialsCardPredict struct {
	ID    string `tl:"id"`
	Title string `tl:"title"`
}

func (*PaymentSavedCredentialsCardPredict) CRC() uint32 {
	return 0xcdc27a1f
}
func (*PaymentSavedCredentialsCardPredict) _PaymentSavedCredentials() {}

type Peer interface {
	tl.TLObject
	_Peer()
}

var (
	_ Peer = (*PeerUserPredict)(nil)
	_ Peer = (*PeerChatPredict)(nil)
	_ Peer = (*PeerChannelPredict)(nil)
)

type PeerUserPredict struct {
	UserID int64 `tl:"user_id"`
}

func (*PeerUserPredict) CRC() uint32 {
	return 0x59511722
}
func (*PeerUserPredict) _Peer() {}

type PeerChatPredict struct {
	ChatID int64 `tl:"chat_id"`
}

func (*PeerChatPredict) CRC() uint32 {
	return 0x36c6019a
}
func (*PeerChatPredict) _Peer() {}

type PeerChannelPredict struct {
	ChannelID int64 `tl:"channel_id"`
}

func (*PeerChannelPredict) CRC() uint32 {
	return 0xa2a5371e
}
func (*PeerChannelPredict) _Peer() {}

type PeerBlocked interface {
	tl.TLObject
	_PeerBlocked()
}

var (
	_ PeerBlocked = (*PeerBlockedPredict)(nil)
)

type PeerBlockedPredict struct {
	PeerID Peer  `tl:"peer_id"`
	Date   int32 `tl:"date"`
}

func (*PeerBlockedPredict) CRC() uint32 {
	return 0xe8fd8014
}
func (*PeerBlockedPredict) _PeerBlocked() {}

type PeerColor interface {
	tl.TLObject
	_PeerColor()
}

var (
	_ PeerColor = (*PeerColorPredict)(nil)
)

type PeerColorPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Color             *int32   `tl:"color,omitempty:flags:0"`
	BackgroundEmojiID *int64   `tl:"background_emoji_id,omitempty:flags:1"`
}

func (*PeerColorPredict) CRC() uint32 {
	return 0xb54b5acf
}
func (*PeerColorPredict) _PeerColor() {}

type PeerLocated interface {
	tl.TLObject
	_PeerLocated()
}

var (
	_ PeerLocated = (*PeerLocatedPredict)(nil)
	_ PeerLocated = (*PeerSelfLocatedPredict)(nil)
)

type PeerLocatedPredict struct {
	Peer     Peer  `tl:"peer"`
	Expires  int32 `tl:"expires"`
	Distance int32 `tl:"distance"`
}

func (*PeerLocatedPredict) CRC() uint32 {
	return 0xca461b5d
}
func (*PeerLocatedPredict) _PeerLocated() {}

type PeerSelfLocatedPredict struct {
	Expires int32 `tl:"expires"`
}

func (*PeerSelfLocatedPredict) CRC() uint32 {
	return 0xf8ec284b
}
func (*PeerSelfLocatedPredict) _PeerLocated() {}

type PeerNotifySettings interface {
	tl.TLObject
	_PeerNotifySettings()
}

var (
	_ PeerNotifySettings = (*PeerNotifySettingsPredict)(nil)
)

type PeerNotifySettingsPredict struct {
	_                   struct{}          `tl:"flags,bitflag"`
	ShowPreviews        *bool             `tl:"show_previews,omitempty:flags:0"`
	Silent              *bool             `tl:"silent,omitempty:flags:1"`
	MuteUntil           *int32            `tl:"mute_until,omitempty:flags:2"`
	IosSound            NotificationSound `tl:"ios_sound,omitempty:flags:3"`
	AndroidSound        NotificationSound `tl:"android_sound,omitempty:flags:4"`
	OtherSound          NotificationSound `tl:"other_sound,omitempty:flags:5"`
	StoriesMuted        *bool             `tl:"stories_muted,omitempty:flags:6"`
	StoriesHideSender   *bool             `tl:"stories_hide_sender,omitempty:flags:7"`
	StoriesIosSound     NotificationSound `tl:"stories_ios_sound,omitempty:flags:8"`
	StoriesAndroidSound NotificationSound `tl:"stories_android_sound,omitempty:flags:9"`
	StoriesOtherSound   NotificationSound `tl:"stories_other_sound,omitempty:flags:10"`
}

func (*PeerNotifySettingsPredict) CRC() uint32 {
	return 0x99622c0c
}
func (*PeerNotifySettingsPredict) _PeerNotifySettings() {}

type PeerSettings interface {
	tl.TLObject
	_PeerSettings()
}

var (
	_ PeerSettings = (*PeerSettingsPredict)(nil)
)

type PeerSettingsPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	ReportSpam            bool     `tl:"report_spam,omitempty:flags:0,implicit"`
	AddContact            bool     `tl:"add_contact,omitempty:flags:1,implicit"`
	BlockContact          bool     `tl:"block_contact,omitempty:flags:2,implicit"`
	ShareContact          bool     `tl:"share_contact,omitempty:flags:3,implicit"`
	NeedContactsException bool     `tl:"need_contacts_exception,omitempty:flags:4,implicit"`
	ReportGeo             bool     `tl:"report_geo,omitempty:flags:5,implicit"`
	Autoarchived          bool     `tl:"autoarchived,omitempty:flags:7,implicit"`
	InviteMembers         bool     `tl:"invite_members,omitempty:flags:8,implicit"`
	RequestChatBroadcast  bool     `tl:"request_chat_broadcast,omitempty:flags:10,implicit"`
	BusinessBotPaused     bool     `tl:"business_bot_paused,omitempty:flags:11,implicit"`
	BusinessBotCanReply   bool     `tl:"business_bot_can_reply,omitempty:flags:12,implicit"`
	GeoDistance           *int32   `tl:"geo_distance,omitempty:flags:6"`
	RequestChatTitle      *string  `tl:"request_chat_title,omitempty:flags:9"`
	RequestChatDate       *int32   `tl:"request_chat_date,omitempty:flags:9"`
	BusinessBotID         *int64   `tl:"business_bot_id,omitempty:flags:13"`
	BusinessBotManageURL  *string  `tl:"business_bot_manage_url,omitempty:flags:13"`
}

func (*PeerSettingsPredict) CRC() uint32 {
	return 0xacd66c5e
}
func (*PeerSettingsPredict) _PeerSettings() {}

type PeerStories interface {
	tl.TLObject
	_PeerStories()
}

var (
	_ PeerStories = (*PeerStoriesPredict)(nil)
)

type PeerStoriesPredict struct {
	_         struct{}    `tl:"flags,bitflag"`
	Peer      Peer        `tl:"peer"`
	MaxReadID *int32      `tl:"max_read_id,omitempty:flags:0"`
	Stories   []StoryItem `tl:"stories"`
}

func (*PeerStoriesPredict) CRC() uint32 {
	return 0x9a35e999
}
func (*PeerStoriesPredict) _PeerStories() {}

type PhoneCall interface {
	tl.TLObject
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
	ID int64 `tl:"id"`
}

func (*PhoneCallEmptyPredict) CRC() uint32 {
	return 0x5366c915
}
func (*PhoneCallEmptyPredict) _PhoneCall() {}

type PhoneCallWaitingPredict struct {
	_             struct{}          `tl:"flags,bitflag"`
	Video         bool              `tl:"video,omitempty:flags:6,implicit"`
	ID            int64             `tl:"id"`
	AccessHash    int64             `tl:"access_hash"`
	Date          int32             `tl:"date"`
	AdminID       int64             `tl:"admin_id"`
	ParticipantID int64             `tl:"participant_id"`
	Protocol      PhoneCallProtocol `tl:"protocol"`
	ReceiveDate   *int32            `tl:"receive_date,omitempty:flags:0"`
}

func (*PhoneCallWaitingPredict) CRC() uint32 {
	return 0xc5226f17
}
func (*PhoneCallWaitingPredict) _PhoneCall() {}

type PhoneCallRequestedPredict struct {
	_             struct{}          `tl:"flags,bitflag"`
	Video         bool              `tl:"video,omitempty:flags:6,implicit"`
	ID            int64             `tl:"id"`
	AccessHash    int64             `tl:"access_hash"`
	Date          int32             `tl:"date"`
	AdminID       int64             `tl:"admin_id"`
	ParticipantID int64             `tl:"participant_id"`
	GAHash        []byte            `tl:"g_a_hash"`
	Protocol      PhoneCallProtocol `tl:"protocol"`
}

func (*PhoneCallRequestedPredict) CRC() uint32 {
	return 0x14b0ed0c
}
func (*PhoneCallRequestedPredict) _PhoneCall() {}

type PhoneCallAcceptedPredict struct {
	_             struct{}          `tl:"flags,bitflag"`
	Video         bool              `tl:"video,omitempty:flags:6,implicit"`
	ID            int64             `tl:"id"`
	AccessHash    int64             `tl:"access_hash"`
	Date          int32             `tl:"date"`
	AdminID       int64             `tl:"admin_id"`
	ParticipantID int64             `tl:"participant_id"`
	GB            []byte            `tl:"g_b"`
	Protocol      PhoneCallProtocol `tl:"protocol"`
}

func (*PhoneCallAcceptedPredict) CRC() uint32 {
	return 0x3660c311
}
func (*PhoneCallAcceptedPredict) _PhoneCall() {}

type PhoneCallPredict struct {
	_                struct{}          `tl:"flags,bitflag"`
	P2PAllowed       bool              `tl:"p2p_allowed,omitempty:flags:5,implicit"`
	Video            bool              `tl:"video,omitempty:flags:6,implicit"`
	ID               int64             `tl:"id"`
	AccessHash       int64             `tl:"access_hash"`
	Date             int32             `tl:"date"`
	AdminID          int64             `tl:"admin_id"`
	ParticipantID    int64             `tl:"participant_id"`
	GAOrB            []byte            `tl:"g_a_or_b"`
	KeyFingerprint   int64             `tl:"key_fingerprint"`
	Protocol         PhoneCallProtocol `tl:"protocol"`
	Connections      []PhoneConnection `tl:"connections"`
	StartDate        int32             `tl:"start_date"`
	CustomParameters DataJSON          `tl:"custom_parameters,omitempty:flags:7"`
}

func (*PhoneCallPredict) CRC() uint32 {
	return 0x30535af5
}
func (*PhoneCallPredict) _PhoneCall() {}

type PhoneCallDiscardedPredict struct {
	_          struct{}               `tl:"flags,bitflag"`
	NeedRating bool                   `tl:"need_rating,omitempty:flags:2,implicit"`
	NeedDebug  bool                   `tl:"need_debug,omitempty:flags:3,implicit"`
	Video      bool                   `tl:"video,omitempty:flags:6,implicit"`
	ID         int64                  `tl:"id"`
	Reason     PhoneCallDiscardReason `tl:"reason,omitempty:flags:0"`
	Duration   *int32                 `tl:"duration,omitempty:flags:1"`
}

func (*PhoneCallDiscardedPredict) CRC() uint32 {
	return 0x50ca4de1
}
func (*PhoneCallDiscardedPredict) _PhoneCall() {}

type PhoneCallDiscardReason interface {
	tl.TLObject
	_PhoneCallDiscardReason()
}

var (
	_ PhoneCallDiscardReason = (*PhoneCallDiscardReasonMissedPredict)(nil)
	_ PhoneCallDiscardReason = (*PhoneCallDiscardReasonDisconnectPredict)(nil)
	_ PhoneCallDiscardReason = (*PhoneCallDiscardReasonHangupPredict)(nil)
	_ PhoneCallDiscardReason = (*PhoneCallDiscardReasonBusyPredict)(nil)
)

type PhoneCallDiscardReasonMissedPredict struct{}

func (*PhoneCallDiscardReasonMissedPredict) CRC() uint32 {
	return 0x85e42301
}
func (*PhoneCallDiscardReasonMissedPredict) _PhoneCallDiscardReason() {}

type PhoneCallDiscardReasonDisconnectPredict struct{}

func (*PhoneCallDiscardReasonDisconnectPredict) CRC() uint32 {
	return 0xe095c1a0
}
func (*PhoneCallDiscardReasonDisconnectPredict) _PhoneCallDiscardReason() {}

type PhoneCallDiscardReasonHangupPredict struct{}

func (*PhoneCallDiscardReasonHangupPredict) CRC() uint32 {
	return 0x57adc690
}
func (*PhoneCallDiscardReasonHangupPredict) _PhoneCallDiscardReason() {}

type PhoneCallDiscardReasonBusyPredict struct{}

func (*PhoneCallDiscardReasonBusyPredict) CRC() uint32 {
	return 0xfaf7e8c9
}
func (*PhoneCallDiscardReasonBusyPredict) _PhoneCallDiscardReason() {}

type PhoneCallProtocol interface {
	tl.TLObject
	_PhoneCallProtocol()
}

var (
	_ PhoneCallProtocol = (*PhoneCallProtocolPredict)(nil)
)

type PhoneCallProtocolPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	UdpP2P          bool     `tl:"udp_p2p,omitempty:flags:0,implicit"`
	UdpReflector    bool     `tl:"udp_reflector,omitempty:flags:1,implicit"`
	MinLayer        int32    `tl:"min_layer"`
	MaxLayer        int32    `tl:"max_layer"`
	LibraryVersions []string `tl:"library_versions"`
}

func (*PhoneCallProtocolPredict) CRC() uint32 {
	return 0xfc878fc8
}
func (*PhoneCallProtocolPredict) _PhoneCallProtocol() {}

type PhoneConnection interface {
	tl.TLObject
	_PhoneConnection()
}

var (
	_ PhoneConnection = (*PhoneConnectionPredict)(nil)
	_ PhoneConnection = (*PhoneConnectionWebrtcPredict)(nil)
)

type PhoneConnectionPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Tcp     bool     `tl:"tcp,omitempty:flags:0,implicit"`
	ID      int64    `tl:"id"`
	Ip      string   `tl:"ip"`
	Ipv6    string   `tl:"ipv6"`
	Port    int32    `tl:"port"`
	PeerTag []byte   `tl:"peer_tag"`
}

func (*PhoneConnectionPredict) CRC() uint32 {
	return 0x9cc123c7
}
func (*PhoneConnectionPredict) _PhoneConnection() {}

type PhoneConnectionWebrtcPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Turn     bool     `tl:"turn,omitempty:flags:0,implicit"`
	Stun     bool     `tl:"stun,omitempty:flags:1,implicit"`
	ID       int64    `tl:"id"`
	Ip       string   `tl:"ip"`
	Ipv6     string   `tl:"ipv6"`
	Port     int32    `tl:"port"`
	Username string   `tl:"username"`
	Password string   `tl:"password"`
}

func (*PhoneConnectionWebrtcPredict) CRC() uint32 {
	return 0x635fe375
}
func (*PhoneConnectionWebrtcPredict) _PhoneConnection() {}

type Photo interface {
	tl.TLObject
	_Photo()
}

var (
	_ Photo = (*PhotoEmptyPredict)(nil)
	_ Photo = (*PhotoPredict)(nil)
)

type PhotoEmptyPredict struct {
	ID int64 `tl:"id"`
}

func (*PhotoEmptyPredict) CRC() uint32 {
	return 0x2331b22d
}
func (*PhotoEmptyPredict) _Photo() {}

type PhotoPredict struct {
	_             struct{}    `tl:"flags,bitflag"`
	HasStickers   bool        `tl:"has_stickers,omitempty:flags:0,implicit"`
	ID            int64       `tl:"id"`
	AccessHash    int64       `tl:"access_hash"`
	FileReference []byte      `tl:"file_reference"`
	Date          int32       `tl:"date"`
	Sizes         []PhotoSize `tl:"sizes"`
	VideoSizes    []VideoSize `tl:"video_sizes,omitempty:flags:1"`
	DcID          int32       `tl:"dc_id"`
}

func (*PhotoPredict) CRC() uint32 {
	return 0xfb197a65
}
func (*PhotoPredict) _Photo() {}

type PhotoSize interface {
	tl.TLObject
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
	Type string `tl:"type"`
}

func (*PhotoSizeEmptyPredict) CRC() uint32 {
	return 0x0e17e23c
}
func (*PhotoSizeEmptyPredict) _PhotoSize() {}

type PhotoSizePredict struct {
	Type string `tl:"type"`
	W    int32  `tl:"w"`
	H    int32  `tl:"h"`
	Size int32  `tl:"size"`
}

func (*PhotoSizePredict) CRC() uint32 {
	return 0x75c78e60
}
func (*PhotoSizePredict) _PhotoSize() {}

type PhotoCachedSizePredict struct {
	Type  string `tl:"type"`
	W     int32  `tl:"w"`
	H     int32  `tl:"h"`
	Bytes []byte `tl:"bytes"`
}

func (*PhotoCachedSizePredict) CRC() uint32 {
	return 0x021e1ad6
}
func (*PhotoCachedSizePredict) _PhotoSize() {}

type PhotoStrippedSizePredict struct {
	Type  string `tl:"type"`
	Bytes []byte `tl:"bytes"`
}

func (*PhotoStrippedSizePredict) CRC() uint32 {
	return 0xe0b0bc2e
}
func (*PhotoStrippedSizePredict) _PhotoSize() {}

type PhotoSizeProgressivePredict struct {
	Type  string  `tl:"type"`
	W     int32   `tl:"w"`
	H     int32   `tl:"h"`
	Sizes []int32 `tl:"sizes"`
}

func (*PhotoSizeProgressivePredict) CRC() uint32 {
	return 0xfa3efb95
}
func (*PhotoSizeProgressivePredict) _PhotoSize() {}

type PhotoPathSizePredict struct {
	Type  string `tl:"type"`
	Bytes []byte `tl:"bytes"`
}

func (*PhotoPathSizePredict) CRC() uint32 {
	return 0xd8214d41
}
func (*PhotoPathSizePredict) _PhotoSize() {}

type Poll interface {
	tl.TLObject
	_Poll()
}

var (
	_ Poll = (*PollPredict)(nil)
)

type PollPredict struct {
	ID             int64            `tl:"id"`
	_              struct{}         `tl:"flags,bitflag"`
	Closed         bool             `tl:"closed,omitempty:flags:0,implicit"`
	PublicVoters   bool             `tl:"public_voters,omitempty:flags:1,implicit"`
	MultipleChoice bool             `tl:"multiple_choice,omitempty:flags:2,implicit"`
	Quiz           bool             `tl:"quiz,omitempty:flags:3,implicit"`
	Question       TextWithEntities `tl:"question"`
	Answers        []PollAnswer     `tl:"answers"`
	ClosePeriod    *int32           `tl:"close_period,omitempty:flags:4"`
	CloseDate      *int32           `tl:"close_date,omitempty:flags:5"`
}

func (*PollPredict) CRC() uint32 {
	return 0x58747131
}
func (*PollPredict) _Poll() {}

type PollAnswer interface {
	tl.TLObject
	_PollAnswer()
}

var (
	_ PollAnswer = (*PollAnswerPredict)(nil)
)

type PollAnswerPredict struct {
	Text   TextWithEntities `tl:"text"`
	Option []byte           `tl:"option"`
}

func (*PollAnswerPredict) CRC() uint32 {
	return 0xff16e2ca
}
func (*PollAnswerPredict) _PollAnswer() {}

type PollAnswerVoters interface {
	tl.TLObject
	_PollAnswerVoters()
}

var (
	_ PollAnswerVoters = (*PollAnswerVotersPredict)(nil)
)

type PollAnswerVotersPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Chosen  bool     `tl:"chosen,omitempty:flags:0,implicit"`
	Correct bool     `tl:"correct,omitempty:flags:1,implicit"`
	Option  []byte   `tl:"option"`
	Voters  int32    `tl:"voters"`
}

func (*PollAnswerVotersPredict) CRC() uint32 {
	return 0x3b6ddad2
}
func (*PollAnswerVotersPredict) _PollAnswerVoters() {}

type PollResults interface {
	tl.TLObject
	_PollResults()
}

var (
	_ PollResults = (*PollResultsPredict)(nil)
)

type PollResultsPredict struct {
	_                struct{}           `tl:"flags,bitflag"`
	Min              bool               `tl:"min,omitempty:flags:0,implicit"`
	Results          []PollAnswerVoters `tl:"results,omitempty:flags:1"`
	TotalVoters      *int32             `tl:"total_voters,omitempty:flags:2"`
	RecentVoters     []Peer             `tl:"recent_voters,omitempty:flags:3"`
	Solution         *string            `tl:"solution,omitempty:flags:4"`
	SolutionEntities []MessageEntity    `tl:"solution_entities,omitempty:flags:4"`
}

func (*PollResultsPredict) CRC() uint32 {
	return 0x7adf2420
}
func (*PollResultsPredict) _PollResults() {}

type PopularContact interface {
	tl.TLObject
	_PopularContact()
}

var (
	_ PopularContact = (*PopularContactPredict)(nil)
)

type PopularContactPredict struct {
	ClientID  int64 `tl:"client_id"`
	Importers int32 `tl:"importers"`
}

func (*PopularContactPredict) CRC() uint32 {
	return 0x5ce14175
}
func (*PopularContactPredict) _PopularContact() {}

type PostAddress interface {
	tl.TLObject
	_PostAddress()
}

var (
	_ PostAddress = (*PostAddressPredict)(nil)
)

type PostAddressPredict struct {
	StreetLine1 string `tl:"street_line1"`
	StreetLine2 string `tl:"street_line2"`
	City        string `tl:"city"`
	State       string `tl:"state"`
	CountryIso2 string `tl:"country_iso2"`
	PostCode    string `tl:"post_code"`
}

func (*PostAddressPredict) CRC() uint32 {
	return 0x1e8caaeb
}
func (*PostAddressPredict) _PostAddress() {}

type PostInteractionCounters interface {
	tl.TLObject
	_PostInteractionCounters()
}

var (
	_ PostInteractionCounters = (*PostInteractionCountersMessagePredict)(nil)
	_ PostInteractionCounters = (*PostInteractionCountersStoryPredict)(nil)
)

type PostInteractionCountersMessagePredict struct {
	MsgID     int32 `tl:"msg_id"`
	Views     int32 `tl:"views"`
	Forwards  int32 `tl:"forwards"`
	Reactions int32 `tl:"reactions"`
}

func (*PostInteractionCountersMessagePredict) CRC() uint32 {
	return 0xe7058e7f
}
func (*PostInteractionCountersMessagePredict) _PostInteractionCounters() {}

type PostInteractionCountersStoryPredict struct {
	StoryID   int32 `tl:"story_id"`
	Views     int32 `tl:"views"`
	Forwards  int32 `tl:"forwards"`
	Reactions int32 `tl:"reactions"`
}

func (*PostInteractionCountersStoryPredict) CRC() uint32 {
	return 0x8a480e27
}
func (*PostInteractionCountersStoryPredict) _PostInteractionCounters() {}

type PremiumGiftCodeOption interface {
	tl.TLObject
	_PremiumGiftCodeOption()
}

var (
	_ PremiumGiftCodeOption = (*PremiumGiftCodeOptionPredict)(nil)
)

type PremiumGiftCodeOptionPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	Users         int32    `tl:"users"`
	Months        int32    `tl:"months"`
	StoreProduct  *string  `tl:"store_product,omitempty:flags:0"`
	StoreQuantity *int32   `tl:"store_quantity,omitempty:flags:1"`
	Currency      string   `tl:"currency"`
	Amount        int64    `tl:"amount"`
}

func (*PremiumGiftCodeOptionPredict) CRC() uint32 {
	return 0x257e962b
}
func (*PremiumGiftCodeOptionPredict) _PremiumGiftCodeOption() {}

type PremiumGiftOption interface {
	tl.TLObject
	_PremiumGiftOption()
}

var (
	_ PremiumGiftOption = (*PremiumGiftOptionPredict)(nil)
)

type PremiumGiftOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Months       int32    `tl:"months"`
	Currency     string   `tl:"currency"`
	Amount       int64    `tl:"amount"`
	BotURL       string   `tl:"bot_url"`
	StoreProduct *string  `tl:"store_product,omitempty:flags:0"`
}

func (*PremiumGiftOptionPredict) CRC() uint32 {
	return 0x74c34319
}
func (*PremiumGiftOptionPredict) _PremiumGiftOption() {}

type PremiumSubscriptionOption interface {
	tl.TLObject
	_PremiumSubscriptionOption()
}

var (
	_ PremiumSubscriptionOption = (*PremiumSubscriptionOptionPredict)(nil)
)

type PremiumSubscriptionOptionPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Current            bool     `tl:"current,omitempty:flags:1,implicit"`
	CanPurchaseUpgrade bool     `tl:"can_purchase_upgrade,omitempty:flags:2,implicit"`
	Transaction        *string  `tl:"transaction,omitempty:flags:3"`
	Months             int32    `tl:"months"`
	Currency           string   `tl:"currency"`
	Amount             int64    `tl:"amount"`
	BotURL             string   `tl:"bot_url"`
	StoreProduct       *string  `tl:"store_product,omitempty:flags:0"`
}

func (*PremiumSubscriptionOptionPredict) CRC() uint32 {
	return 0x5f2d1df2
}
func (*PremiumSubscriptionOptionPredict) _PremiumSubscriptionOption() {}

type PrepaidGiveaway interface {
	tl.TLObject
	_PrepaidGiveaway()
}

var (
	_ PrepaidGiveaway = (*PrepaidGiveawayPredict)(nil)
)

type PrepaidGiveawayPredict struct {
	ID       int64 `tl:"id"`
	Months   int32 `tl:"months"`
	Quantity int32 `tl:"quantity"`
	Date     int32 `tl:"date"`
}

func (*PrepaidGiveawayPredict) CRC() uint32 {
	return 0xb2539d54
}
func (*PrepaidGiveawayPredict) _PrepaidGiveaway() {}

type PrivacyKey interface {
	tl.TLObject
	_PrivacyKey()
}

var (
	_ PrivacyKey = (*PrivacyKeyStatusTimestampPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyChatInvitePredict)(nil)
	_ PrivacyKey = (*PrivacyKeyPhoneCallPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyPhoneP2PPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyForwardsPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyProfilePhotoPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyPhoneNumberPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyAddedByPhonePredict)(nil)
	_ PrivacyKey = (*PrivacyKeyVoiceMessagesPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyAboutPredict)(nil)
	_ PrivacyKey = (*PrivacyKeyBirthdayPredict)(nil)
)

type PrivacyKeyStatusTimestampPredict struct{}

func (*PrivacyKeyStatusTimestampPredict) CRC() uint32 {
	return 0xbc2eab30
}
func (*PrivacyKeyStatusTimestampPredict) _PrivacyKey() {}

type PrivacyKeyChatInvitePredict struct{}

func (*PrivacyKeyChatInvitePredict) CRC() uint32 {
	return 0x500e6dfa
}
func (*PrivacyKeyChatInvitePredict) _PrivacyKey() {}

type PrivacyKeyPhoneCallPredict struct{}

func (*PrivacyKeyPhoneCallPredict) CRC() uint32 {
	return 0x3d662b7b
}
func (*PrivacyKeyPhoneCallPredict) _PrivacyKey() {}

type PrivacyKeyPhoneP2PPredict struct{}

func (*PrivacyKeyPhoneP2PPredict) CRC() uint32 {
	return 0x39491cc8
}
func (*PrivacyKeyPhoneP2PPredict) _PrivacyKey() {}

type PrivacyKeyForwardsPredict struct{}

func (*PrivacyKeyForwardsPredict) CRC() uint32 {
	return 0x69ec56a3
}
func (*PrivacyKeyForwardsPredict) _PrivacyKey() {}

type PrivacyKeyProfilePhotoPredict struct{}

func (*PrivacyKeyProfilePhotoPredict) CRC() uint32 {
	return 0x96151fed
}
func (*PrivacyKeyProfilePhotoPredict) _PrivacyKey() {}

type PrivacyKeyPhoneNumberPredict struct{}

func (*PrivacyKeyPhoneNumberPredict) CRC() uint32 {
	return 0xd19ae46d
}
func (*PrivacyKeyPhoneNumberPredict) _PrivacyKey() {}

type PrivacyKeyAddedByPhonePredict struct{}

func (*PrivacyKeyAddedByPhonePredict) CRC() uint32 {
	return 0x42ffd42b
}
func (*PrivacyKeyAddedByPhonePredict) _PrivacyKey() {}

type PrivacyKeyVoiceMessagesPredict struct{}

func (*PrivacyKeyVoiceMessagesPredict) CRC() uint32 {
	return 0x0697f414
}
func (*PrivacyKeyVoiceMessagesPredict) _PrivacyKey() {}

type PrivacyKeyAboutPredict struct{}

func (*PrivacyKeyAboutPredict) CRC() uint32 {
	return 0xa486b761
}
func (*PrivacyKeyAboutPredict) _PrivacyKey() {}

type PrivacyKeyBirthdayPredict struct{}

func (*PrivacyKeyBirthdayPredict) CRC() uint32 {
	return 0x2000a518
}
func (*PrivacyKeyBirthdayPredict) _PrivacyKey() {}

type PrivacyRule interface {
	tl.TLObject
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
	Users []int64 `tl:"users"`
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
	Users []int64 `tl:"users"`
}

func (*PrivacyValueDisallowUsersPredict) CRC() uint32 {
	return 0xe4621141
}
func (*PrivacyValueDisallowUsersPredict) _PrivacyRule() {}

type PrivacyValueAllowChatParticipantsPredict struct {
	Chats []int64 `tl:"chats"`
}

func (*PrivacyValueAllowChatParticipantsPredict) CRC() uint32 {
	return 0x6b134e8e
}
func (*PrivacyValueAllowChatParticipantsPredict) _PrivacyRule() {}

type PrivacyValueDisallowChatParticipantsPredict struct {
	Chats []int64 `tl:"chats"`
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
	tl.TLObject
	_PublicForward()
}

var (
	_ PublicForward = (*PublicForwardMessagePredict)(nil)
	_ PublicForward = (*PublicForwardStoryPredict)(nil)
)

type PublicForwardMessagePredict struct {
	Message Message `tl:"message"`
}

func (*PublicForwardMessagePredict) CRC() uint32 {
	return 0x01f2bf4a
}
func (*PublicForwardMessagePredict) _PublicForward() {}

type PublicForwardStoryPredict struct {
	Peer  Peer      `tl:"peer"`
	Story StoryItem `tl:"story"`
}

func (*PublicForwardStoryPredict) CRC() uint32 {
	return 0xedf3add0
}
func (*PublicForwardStoryPredict) _PublicForward() {}

type QuickReply interface {
	tl.TLObject
	_QuickReply()
}

var (
	_ QuickReply = (*QuickReplyPredict)(nil)
)

type QuickReplyPredict struct {
	ShortcutID int32  `tl:"shortcut_id"`
	Shortcut   string `tl:"shortcut"`
	TopMessage int32  `tl:"top_message"`
	Count      int32  `tl:"count"`
}

func (*QuickReplyPredict) CRC() uint32 {
	return 0x0697102b
}
func (*QuickReplyPredict) _QuickReply() {}

type Reaction interface {
	tl.TLObject
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
	Emoticon string `tl:"emoticon"`
}

func (*ReactionEmojiPredict) CRC() uint32 {
	return 0x1b2286b8
}
func (*ReactionEmojiPredict) _Reaction() {}

type ReactionCustomEmojiPredict struct {
	DocumentID int64 `tl:"document_id"`
}

func (*ReactionCustomEmojiPredict) CRC() uint32 {
	return 0x8935fc73
}
func (*ReactionCustomEmojiPredict) _Reaction() {}

type ReactionCount interface {
	tl.TLObject
	_ReactionCount()
}

var (
	_ ReactionCount = (*ReactionCountPredict)(nil)
)

type ReactionCountPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	ChosenOrder *int32   `tl:"chosen_order,omitempty:flags:0"`
	Reaction    Reaction `tl:"reaction"`
	Count       int32    `tl:"count"`
}

func (*ReactionCountPredict) CRC() uint32 {
	return 0xa3d1cb80
}
func (*ReactionCountPredict) _ReactionCount() {}

type ReactionNotificationsFrom interface {
	tl.TLObject
	_ReactionNotificationsFrom()
}

var (
	_ ReactionNotificationsFrom = (*ReactionNotificationsFromContactsPredict)(nil)
	_ ReactionNotificationsFrom = (*ReactionNotificationsFromAllPredict)(nil)
)

type ReactionNotificationsFromContactsPredict struct{}

func (*ReactionNotificationsFromContactsPredict) CRC() uint32 {
	return 0xbac3a61a
}
func (*ReactionNotificationsFromContactsPredict) _ReactionNotificationsFrom() {}

type ReactionNotificationsFromAllPredict struct{}

func (*ReactionNotificationsFromAllPredict) CRC() uint32 {
	return 0x4b9e22a0
}
func (*ReactionNotificationsFromAllPredict) _ReactionNotificationsFrom() {}

type ReactionsNotifySettings interface {
	tl.TLObject
	_ReactionsNotifySettings()
}

var (
	_ ReactionsNotifySettings = (*ReactionsNotifySettingsPredict)(nil)
)

type ReactionsNotifySettingsPredict struct {
	_                  struct{}                  `tl:"flags,bitflag"`
	MessagesNotifyFrom ReactionNotificationsFrom `tl:"messages_notify_from,omitempty:flags:0"`
	StoriesNotifyFrom  ReactionNotificationsFrom `tl:"stories_notify_from,omitempty:flags:1"`
	Sound              NotificationSound         `tl:"sound"`
	ShowPreviews       bool                      `tl:"show_previews"`
}

func (*ReactionsNotifySettingsPredict) CRC() uint32 {
	return 0x56e34970
}
func (*ReactionsNotifySettingsPredict) _ReactionsNotifySettings() {}

type ReadParticipantDate interface {
	tl.TLObject
	_ReadParticipantDate()
}

var (
	_ ReadParticipantDate = (*ReadParticipantDatePredict)(nil)
)

type ReadParticipantDatePredict struct {
	UserID int64 `tl:"user_id"`
	Date   int32 `tl:"date"`
}

func (*ReadParticipantDatePredict) CRC() uint32 {
	return 0x4a4ff172
}
func (*ReadParticipantDatePredict) _ReadParticipantDate() {}

type ReceivedNotifyMessage interface {
	tl.TLObject
	_ReceivedNotifyMessage()
}

var (
	_ ReceivedNotifyMessage = (*ReceivedNotifyMessagePredict)(nil)
)

type ReceivedNotifyMessagePredict struct {
	ID    int32 `tl:"id"`
	Flags int32 `tl:"flags"`
}

func (*ReceivedNotifyMessagePredict) CRC() uint32 {
	return 0xa384b779
}
func (*ReceivedNotifyMessagePredict) _ReceivedNotifyMessage() {}

type RecentMeURL interface {
	tl.TLObject
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
	URL string `tl:"url"`
}

func (*RecentMeURLUnknownPredict) CRC() uint32 {
	return 0x46e1d13d
}
func (*RecentMeURLUnknownPredict) _RecentMeURL() {}

type RecentMeURLUserPredict struct {
	URL    string `tl:"url"`
	UserID int64  `tl:"user_id"`
}

func (*RecentMeURLUserPredict) CRC() uint32 {
	return 0xb92c09e2
}
func (*RecentMeURLUserPredict) _RecentMeURL() {}

type RecentMeURLChatPredict struct {
	URL    string `tl:"url"`
	ChatID int64  `tl:"chat_id"`
}

func (*RecentMeURLChatPredict) CRC() uint32 {
	return 0xb2da71d2
}
func (*RecentMeURLChatPredict) _RecentMeURL() {}

type RecentMeURLChatInvitePredict struct {
	URL        string     `tl:"url"`
	ChatInvite ChatInvite `tl:"chat_invite"`
}

func (*RecentMeURLChatInvitePredict) CRC() uint32 {
	return 0xeb49081d
}
func (*RecentMeURLChatInvitePredict) _RecentMeURL() {}

type RecentMeURLStickerSetPredict struct {
	URL string            `tl:"url"`
	Set StickerSetCovered `tl:"set"`
}

func (*RecentMeURLStickerSetPredict) CRC() uint32 {
	return 0xbc0a57dc
}
func (*RecentMeURLStickerSetPredict) _RecentMeURL() {}

type ReplyMarkup interface {
	tl.TLObject
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
	Selective bool     `tl:"selective,omitempty:flags:2,implicit"`
}

func (*ReplyKeyboardHidePredict) CRC() uint32 {
	return 0xa03e5b85
}
func (*ReplyKeyboardHidePredict) _ReplyMarkup() {}

type ReplyKeyboardForceReplyPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	SingleUse   bool     `tl:"single_use,omitempty:flags:1,implicit"`
	Selective   bool     `tl:"selective,omitempty:flags:2,implicit"`
	Placeholder *string  `tl:"placeholder,omitempty:flags:3"`
}

func (*ReplyKeyboardForceReplyPredict) CRC() uint32 {
	return 0x86b40b08
}
func (*ReplyKeyboardForceReplyPredict) _ReplyMarkup() {}

type ReplyKeyboardMarkupPredict struct {
	_           struct{}            `tl:"flags,bitflag"`
	Resize      bool                `tl:"resize,omitempty:flags:0,implicit"`
	SingleUse   bool                `tl:"single_use,omitempty:flags:1,implicit"`
	Selective   bool                `tl:"selective,omitempty:flags:2,implicit"`
	Persistent  bool                `tl:"persistent,omitempty:flags:4,implicit"`
	Rows        []KeyboardButtonRow `tl:"rows"`
	Placeholder *string             `tl:"placeholder,omitempty:flags:3"`
}

func (*ReplyKeyboardMarkupPredict) CRC() uint32 {
	return 0x85dd99d1
}
func (*ReplyKeyboardMarkupPredict) _ReplyMarkup() {}

type ReplyInlineMarkupPredict struct {
	Rows []KeyboardButtonRow `tl:"rows"`
}

func (*ReplyInlineMarkupPredict) CRC() uint32 {
	return 0x48a30254
}
func (*ReplyInlineMarkupPredict) _ReplyMarkup() {}

type ReportReason interface {
	tl.TLObject
	_ReportReason()
}

var (
	_ ReportReason = (*InputReportReasonSpamPredict)(nil)
	_ ReportReason = (*InputReportReasonViolencePredict)(nil)
	_ ReportReason = (*InputReportReasonPornographyPredict)(nil)
	_ ReportReason = (*InputReportReasonChildAbusePredict)(nil)
	_ ReportReason = (*InputReportReasonOtherPredict)(nil)
	_ ReportReason = (*InputReportReasonCopyrightPredict)(nil)
	_ ReportReason = (*InputReportReasonGeoIrrelevantPredict)(nil)
	_ ReportReason = (*InputReportReasonFakePredict)(nil)
	_ ReportReason = (*InputReportReasonIllegalDrugsPredict)(nil)
	_ ReportReason = (*InputReportReasonPersonalDetailsPredict)(nil)
)

type InputReportReasonSpamPredict struct{}

func (*InputReportReasonSpamPredict) CRC() uint32 {
	return 0x58dbcab8
}
func (*InputReportReasonSpamPredict) _ReportReason() {}

type InputReportReasonViolencePredict struct{}

func (*InputReportReasonViolencePredict) CRC() uint32 {
	return 0x1e22c78d
}
func (*InputReportReasonViolencePredict) _ReportReason() {}

type InputReportReasonPornographyPredict struct{}

func (*InputReportReasonPornographyPredict) CRC() uint32 {
	return 0x2e59d922
}
func (*InputReportReasonPornographyPredict) _ReportReason() {}

type InputReportReasonChildAbusePredict struct{}

func (*InputReportReasonChildAbusePredict) CRC() uint32 {
	return 0xadf44ee3
}
func (*InputReportReasonChildAbusePredict) _ReportReason() {}

type InputReportReasonOtherPredict struct{}

func (*InputReportReasonOtherPredict) CRC() uint32 {
	return 0xc1e4a2b1
}
func (*InputReportReasonOtherPredict) _ReportReason() {}

type InputReportReasonCopyrightPredict struct{}

func (*InputReportReasonCopyrightPredict) CRC() uint32 {
	return 0x9b89f93a
}
func (*InputReportReasonCopyrightPredict) _ReportReason() {}

type InputReportReasonGeoIrrelevantPredict struct{}

func (*InputReportReasonGeoIrrelevantPredict) CRC() uint32 {
	return 0xdbd4feed
}
func (*InputReportReasonGeoIrrelevantPredict) _ReportReason() {}

type InputReportReasonFakePredict struct{}

func (*InputReportReasonFakePredict) CRC() uint32 {
	return 0xf5ddd6e7
}
func (*InputReportReasonFakePredict) _ReportReason() {}

type InputReportReasonIllegalDrugsPredict struct{}

func (*InputReportReasonIllegalDrugsPredict) CRC() uint32 {
	return 0x0a8eb2be
}
func (*InputReportReasonIllegalDrugsPredict) _ReportReason() {}

type InputReportReasonPersonalDetailsPredict struct{}

func (*InputReportReasonPersonalDetailsPredict) CRC() uint32 {
	return 0x9ec7863d
}
func (*InputReportReasonPersonalDetailsPredict) _ReportReason() {}

type RequestPeerType interface {
	tl.TLObject
	_RequestPeerType()
}

var (
	_ RequestPeerType = (*RequestPeerTypeUserPredict)(nil)
	_ RequestPeerType = (*RequestPeerTypeChatPredict)(nil)
	_ RequestPeerType = (*RequestPeerTypeBroadcastPredict)(nil)
)

type RequestPeerTypeUserPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Bot     *bool    `tl:"bot,omitempty:flags:0"`
	Premium *bool    `tl:"premium,omitempty:flags:1"`
}

func (*RequestPeerTypeUserPredict) CRC() uint32 {
	return 0x5f3b8a00
}
func (*RequestPeerTypeUserPredict) _RequestPeerType() {}

type RequestPeerTypeChatPredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	Creator         bool            `tl:"creator,omitempty:flags:0,implicit"`
	BotParticipant  bool            `tl:"bot_participant,omitempty:flags:5,implicit"`
	HasUsername     *bool           `tl:"has_username,omitempty:flags:3"`
	Forum           *bool           `tl:"forum,omitempty:flags:4"`
	UserAdminRights ChatAdminRights `tl:"user_admin_rights,omitempty:flags:1"`
	BotAdminRights  ChatAdminRights `tl:"bot_admin_rights,omitempty:flags:2"`
}

func (*RequestPeerTypeChatPredict) CRC() uint32 {
	return 0xc9f06e1b
}
func (*RequestPeerTypeChatPredict) _RequestPeerType() {}

type RequestPeerTypeBroadcastPredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	Creator         bool            `tl:"creator,omitempty:flags:0,implicit"`
	HasUsername     *bool           `tl:"has_username,omitempty:flags:3"`
	UserAdminRights ChatAdminRights `tl:"user_admin_rights,omitempty:flags:1"`
	BotAdminRights  ChatAdminRights `tl:"bot_admin_rights,omitempty:flags:2"`
}

func (*RequestPeerTypeBroadcastPredict) CRC() uint32 {
	return 0x339bef6c
}
func (*RequestPeerTypeBroadcastPredict) _RequestPeerType() {}

type RequestedPeer interface {
	tl.TLObject
	_RequestedPeer()
}

var (
	_ RequestedPeer = (*RequestedPeerUserPredict)(nil)
	_ RequestedPeer = (*RequestedPeerChatPredict)(nil)
	_ RequestedPeer = (*RequestedPeerChannelPredict)(nil)
)

type RequestedPeerUserPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	UserID    int64    `tl:"user_id"`
	FirstName *string  `tl:"first_name,omitempty:flags:0"`
	LastName  *string  `tl:"last_name,omitempty:flags:0"`
	Username  *string  `tl:"username,omitempty:flags:1"`
	Photo     Photo    `tl:"photo,omitempty:flags:2"`
}

func (*RequestedPeerUserPredict) CRC() uint32 {
	return 0xd62ff46a
}
func (*RequestedPeerUserPredict) _RequestedPeer() {}

type RequestedPeerChatPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	ChatID int64    `tl:"chat_id"`
	Title  *string  `tl:"title,omitempty:flags:0"`
	Photo  Photo    `tl:"photo,omitempty:flags:2"`
}

func (*RequestedPeerChatPredict) CRC() uint32 {
	return 0x7307544f
}
func (*RequestedPeerChatPredict) _RequestedPeer() {}

type RequestedPeerChannelPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64    `tl:"channel_id"`
	Title     *string  `tl:"title,omitempty:flags:0"`
	Username  *string  `tl:"username,omitempty:flags:1"`
	Photo     Photo    `tl:"photo,omitempty:flags:2"`
}

func (*RequestedPeerChannelPredict) CRC() uint32 {
	return 0x8ba403e4
}
func (*RequestedPeerChannelPredict) _RequestedPeer() {}

type RestrictionReason interface {
	tl.TLObject
	_RestrictionReason()
}

var (
	_ RestrictionReason = (*RestrictionReasonPredict)(nil)
)

type RestrictionReasonPredict struct {
	Platform string `tl:"platform"`
	Reason   string `tl:"reason"`
	Text     string `tl:"text"`
}

func (*RestrictionReasonPredict) CRC() uint32 {
	return 0xd072acb4
}
func (*RestrictionReasonPredict) _RestrictionReason() {}

type RichText interface {
	tl.TLObject
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
	Text string `tl:"text"`
}

func (*TextPlainPredict) CRC() uint32 {
	return 0x744694e0
}
func (*TextPlainPredict) _RichText() {}

type TextBoldPredict struct {
	Text RichText `tl:"text"`
}

func (*TextBoldPredict) CRC() uint32 {
	return 0x6724abc4
}
func (*TextBoldPredict) _RichText() {}

type TextItalicPredict struct {
	Text RichText `tl:"text"`
}

func (*TextItalicPredict) CRC() uint32 {
	return 0xd912a59c
}
func (*TextItalicPredict) _RichText() {}

type TextUnderlinePredict struct {
	Text RichText `tl:"text"`
}

func (*TextUnderlinePredict) CRC() uint32 {
	return 0xc12622c4
}
func (*TextUnderlinePredict) _RichText() {}

type TextStrikePredict struct {
	Text RichText `tl:"text"`
}

func (*TextStrikePredict) CRC() uint32 {
	return 0x9bf8bb95
}
func (*TextStrikePredict) _RichText() {}

type TextFixedPredict struct {
	Text RichText `tl:"text"`
}

func (*TextFixedPredict) CRC() uint32 {
	return 0x6c3f19b9
}
func (*TextFixedPredict) _RichText() {}

type TextURLPredict struct {
	Text      RichText `tl:"text"`
	URL       string   `tl:"url"`
	WebpageID int64    `tl:"webpage_id"`
}

func (*TextURLPredict) CRC() uint32 {
	return 0x3c2884c1
}
func (*TextURLPredict) _RichText() {}

type TextEmailPredict struct {
	Text  RichText `tl:"text"`
	Email string   `tl:"email"`
}

func (*TextEmailPredict) CRC() uint32 {
	return 0xde5a0dd6
}
func (*TextEmailPredict) _RichText() {}

type TextConcatPredict struct {
	Texts []RichText `tl:"texts"`
}

func (*TextConcatPredict) CRC() uint32 {
	return 0x7e6260d7
}
func (*TextConcatPredict) _RichText() {}

type TextSubscriptPredict struct {
	Text RichText `tl:"text"`
}

func (*TextSubscriptPredict) CRC() uint32 {
	return 0xed6a8504
}
func (*TextSubscriptPredict) _RichText() {}

type TextSuperscriptPredict struct {
	Text RichText `tl:"text"`
}

func (*TextSuperscriptPredict) CRC() uint32 {
	return 0xc7fb5e01
}
func (*TextSuperscriptPredict) _RichText() {}

type TextMarkedPredict struct {
	Text RichText `tl:"text"`
}

func (*TextMarkedPredict) CRC() uint32 {
	return 0x034b8621
}
func (*TextMarkedPredict) _RichText() {}

type TextPhonePredict struct {
	Text  RichText `tl:"text"`
	Phone string   `tl:"phone"`
}

func (*TextPhonePredict) CRC() uint32 {
	return 0x1ccb966a
}
func (*TextPhonePredict) _RichText() {}

type TextImagePredict struct {
	DocumentID int64 `tl:"document_id"`
	W          int32 `tl:"w"`
	H          int32 `tl:"h"`
}

func (*TextImagePredict) CRC() uint32 {
	return 0x081ccf4f
}
func (*TextImagePredict) _RichText() {}

type TextAnchorPredict struct {
	Text RichText `tl:"text"`
	Name string   `tl:"name"`
}

func (*TextAnchorPredict) CRC() uint32 {
	return 0x35553762
}
func (*TextAnchorPredict) _RichText() {}

type SavedContact interface {
	tl.TLObject
	_SavedContact()
}

var (
	_ SavedContact = (*SavedPhoneContactPredict)(nil)
)

type SavedPhoneContactPredict struct {
	Phone     string `tl:"phone"`
	FirstName string `tl:"first_name"`
	LastName  string `tl:"last_name"`
	Date      int32  `tl:"date"`
}

func (*SavedPhoneContactPredict) CRC() uint32 {
	return 0x1142bd56
}
func (*SavedPhoneContactPredict) _SavedContact() {}

type SavedDialog interface {
	tl.TLObject
	_SavedDialog()
}

var (
	_ SavedDialog = (*SavedDialogPredict)(nil)
)

type SavedDialogPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Pinned     bool     `tl:"pinned,omitempty:flags:2,implicit"`
	Peer       Peer     `tl:"peer"`
	TopMessage int32    `tl:"top_message"`
}

func (*SavedDialogPredict) CRC() uint32 {
	return 0xbd87cb6c
}
func (*SavedDialogPredict) _SavedDialog() {}

type SavedReactionTag interface {
	tl.TLObject
	_SavedReactionTag()
}

var (
	_ SavedReactionTag = (*SavedReactionTagPredict)(nil)
)

type SavedReactionTagPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Reaction Reaction `tl:"reaction"`
	Title    *string  `tl:"title,omitempty:flags:0"`
	Count    int32    `tl:"count"`
}

func (*SavedReactionTagPredict) CRC() uint32 {
	return 0xcb6ff828
}
func (*SavedReactionTagPredict) _SavedReactionTag() {}

type SearchResultsCalendarPeriod interface {
	tl.TLObject
	_SearchResultsCalendarPeriod()
}

var (
	_ SearchResultsCalendarPeriod = (*SearchResultsCalendarPeriodPredict)(nil)
)

type SearchResultsCalendarPeriodPredict struct {
	Date     int32 `tl:"date"`
	MinMsgID int32 `tl:"min_msg_id"`
	MaxMsgID int32 `tl:"max_msg_id"`
	Count    int32 `tl:"count"`
}

func (*SearchResultsCalendarPeriodPredict) CRC() uint32 {
	return 0xc9b0539f
}
func (*SearchResultsCalendarPeriodPredict) _SearchResultsCalendarPeriod() {}

type SearchResultsPosition interface {
	tl.TLObject
	_SearchResultsPosition()
}

var (
	_ SearchResultsPosition = (*SearchResultPositionPredict)(nil)
)

type SearchResultPositionPredict struct {
	MsgID  int32 `tl:"msg_id"`
	Date   int32 `tl:"date"`
	Offset int32 `tl:"offset"`
}

func (*SearchResultPositionPredict) CRC() uint32 {
	return 0x7f648b67
}
func (*SearchResultPositionPredict) _SearchResultsPosition() {}

type SecureCredentialsEncrypted interface {
	tl.TLObject
	_SecureCredentialsEncrypted()
}

var (
	_ SecureCredentialsEncrypted = (*SecureCredentialsEncryptedPredict)(nil)
)

type SecureCredentialsEncryptedPredict struct {
	Data   []byte `tl:"data"`
	Hash   []byte `tl:"hash"`
	Secret []byte `tl:"secret"`
}

func (*SecureCredentialsEncryptedPredict) CRC() uint32 {
	return 0x33f0ea47
}
func (*SecureCredentialsEncryptedPredict) _SecureCredentialsEncrypted() {}

type SecureData interface {
	tl.TLObject
	_SecureData()
}

var (
	_ SecureData = (*SecureDataPredict)(nil)
)

type SecureDataPredict struct {
	Data     []byte `tl:"data"`
	DataHash []byte `tl:"data_hash"`
	Secret   []byte `tl:"secret"`
}

func (*SecureDataPredict) CRC() uint32 {
	return 0x8aeabec3
}
func (*SecureDataPredict) _SecureData() {}

type SecureFile interface {
	tl.TLObject
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
	ID         int64  `tl:"id"`
	AccessHash int64  `tl:"access_hash"`
	Size       int64  `tl:"size"`
	DcID       int32  `tl:"dc_id"`
	Date       int32  `tl:"date"`
	FileHash   []byte `tl:"file_hash"`
	Secret     []byte `tl:"secret"`
}

func (*SecureFilePredict) CRC() uint32 {
	return 0x7d09c27e
}
func (*SecureFilePredict) _SecureFile() {}

type SecurePasswordKdfAlgo interface {
	tl.TLObject
	_SecurePasswordKdfAlgo()
}

var (
	_ SecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoUnknownPredict)(nil)
	_ SecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict)(nil)
	_ SecurePasswordKdfAlgo = (*SecurePasswordKdfAlgoSHA512Predict)(nil)
)

type SecurePasswordKdfAlgoUnknownPredict struct{}

func (*SecurePasswordKdfAlgoUnknownPredict) CRC() uint32 {
	return 0x004a8537
}
func (*SecurePasswordKdfAlgoUnknownPredict) _SecurePasswordKdfAlgo() {}

type SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict struct {
	Salt []byte `tl:"salt"`
}

func (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict) CRC() uint32 {
	return 0xbbf2dda0
}
func (*SecurePasswordKdfAlgoPbkdf2Hmacsha512Iter100000Predict) _SecurePasswordKdfAlgo() {}

type SecurePasswordKdfAlgoSHA512Predict struct {
	Salt []byte `tl:"salt"`
}

func (*SecurePasswordKdfAlgoSHA512Predict) CRC() uint32 {
	return 0x86471d92
}
func (*SecurePasswordKdfAlgoSHA512Predict) _SecurePasswordKdfAlgo() {}

type SecurePlainData interface {
	tl.TLObject
	_SecurePlainData()
}

var (
	_ SecurePlainData = (*SecurePlainPhonePredict)(nil)
	_ SecurePlainData = (*SecurePlainEmailPredict)(nil)
)

type SecurePlainPhonePredict struct {
	Phone string `tl:"phone"`
}

func (*SecurePlainPhonePredict) CRC() uint32 {
	return 0x7d6099dd
}
func (*SecurePlainPhonePredict) _SecurePlainData() {}

type SecurePlainEmailPredict struct {
	Email string `tl:"email"`
}

func (*SecurePlainEmailPredict) CRC() uint32 {
	return 0x21ec5a5f
}
func (*SecurePlainEmailPredict) _SecurePlainData() {}

type SecureRequiredType interface {
	tl.TLObject
	_SecureRequiredType()
}

var (
	_ SecureRequiredType = (*SecureRequiredTypePredict)(nil)
	_ SecureRequiredType = (*SecureRequiredTypeOneOfPredict)(nil)
)

type SecureRequiredTypePredict struct {
	_                   struct{}        `tl:"flags,bitflag"`
	NativeNames         bool            `tl:"native_names,omitempty:flags:0,implicit"`
	SelfieRequired      bool            `tl:"selfie_required,omitempty:flags:1,implicit"`
	TranslationRequired bool            `tl:"translation_required,omitempty:flags:2,implicit"`
	Type                SecureValueType `tl:"type"`
}

func (*SecureRequiredTypePredict) CRC() uint32 {
	return 0x829d99da
}
func (*SecureRequiredTypePredict) _SecureRequiredType() {}

type SecureRequiredTypeOneOfPredict struct {
	Types []SecureRequiredType `tl:"types"`
}

func (*SecureRequiredTypeOneOfPredict) CRC() uint32 {
	return 0x027477b4
}
func (*SecureRequiredTypeOneOfPredict) _SecureRequiredType() {}

type SecureSecretSettings interface {
	tl.TLObject
	_SecureSecretSettings()
}

var (
	_ SecureSecretSettings = (*SecureSecretSettingsPredict)(nil)
)

type SecureSecretSettingsPredict struct {
	SecureAlgo     SecurePasswordKdfAlgo `tl:"secure_algo"`
	SecureSecret   []byte                `tl:"secure_secret"`
	SecureSecretID int64                 `tl:"secure_secret_id"`
}

func (*SecureSecretSettingsPredict) CRC() uint32 {
	return 0x1527bcac
}
func (*SecureSecretSettingsPredict) _SecureSecretSettings() {}

type SecureValue interface {
	tl.TLObject
	_SecureValue()
}

var (
	_ SecureValue = (*SecureValuePredict)(nil)
)

type SecureValuePredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	Type        SecureValueType `tl:"type"`
	Data        SecureData      `tl:"data,omitempty:flags:0"`
	FrontSide   SecureFile      `tl:"front_side,omitempty:flags:1"`
	ReverseSide SecureFile      `tl:"reverse_side,omitempty:flags:2"`
	Selfie      SecureFile      `tl:"selfie,omitempty:flags:3"`
	Translation []SecureFile    `tl:"translation,omitempty:flags:6"`
	Files       []SecureFile    `tl:"files,omitempty:flags:4"`
	PlainData   SecurePlainData `tl:"plain_data,omitempty:flags:5"`
	Hash        []byte          `tl:"hash"`
}

func (*SecureValuePredict) CRC() uint32 {
	return 0x187fa0ca
}
func (*SecureValuePredict) _SecureValue() {}

type SecureValueError interface {
	tl.TLObject
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
	Type     SecureValueType `tl:"type"`
	DataHash []byte          `tl:"data_hash"`
	Field    string          `tl:"field"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorDataPredict) CRC() uint32 {
	return 0xe8a40bd9
}
func (*SecureValueErrorDataPredict) _SecureValueError() {}

type SecureValueErrorFrontSidePredict struct {
	Type     SecureValueType `tl:"type"`
	FileHash []byte          `tl:"file_hash"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorFrontSidePredict) CRC() uint32 {
	return 0x00be3dfa
}
func (*SecureValueErrorFrontSidePredict) _SecureValueError() {}

type SecureValueErrorReverseSidePredict struct {
	Type     SecureValueType `tl:"type"`
	FileHash []byte          `tl:"file_hash"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorReverseSidePredict) CRC() uint32 {
	return 0x868a2aa5
}
func (*SecureValueErrorReverseSidePredict) _SecureValueError() {}

type SecureValueErrorSelfiePredict struct {
	Type     SecureValueType `tl:"type"`
	FileHash []byte          `tl:"file_hash"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorSelfiePredict) CRC() uint32 {
	return 0xe537ced6
}
func (*SecureValueErrorSelfiePredict) _SecureValueError() {}

type SecureValueErrorFilePredict struct {
	Type     SecureValueType `tl:"type"`
	FileHash []byte          `tl:"file_hash"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorFilePredict) CRC() uint32 {
	return 0x7a700873
}
func (*SecureValueErrorFilePredict) _SecureValueError() {}

type SecureValueErrorFilesPredict struct {
	Type     SecureValueType `tl:"type"`
	FileHash [][]byte        `tl:"file_hash"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorFilesPredict) CRC() uint32 {
	return 0x666220e9
}
func (*SecureValueErrorFilesPredict) _SecureValueError() {}

type SecureValueErrorPredict struct {
	Type SecureValueType `tl:"type"`
	Hash []byte          `tl:"hash"`
	Text string          `tl:"text"`
}

func (*SecureValueErrorPredict) CRC() uint32 {
	return 0x869d758f
}
func (*SecureValueErrorPredict) _SecureValueError() {}

type SecureValueErrorTranslationFilePredict struct {
	Type     SecureValueType `tl:"type"`
	FileHash []byte          `tl:"file_hash"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorTranslationFilePredict) CRC() uint32 {
	return 0xa1144770
}
func (*SecureValueErrorTranslationFilePredict) _SecureValueError() {}

type SecureValueErrorTranslationFilesPredict struct {
	Type     SecureValueType `tl:"type"`
	FileHash [][]byte        `tl:"file_hash"`
	Text     string          `tl:"text"`
}

func (*SecureValueErrorTranslationFilesPredict) CRC() uint32 {
	return 0x34636dd8
}
func (*SecureValueErrorTranslationFilesPredict) _SecureValueError() {}

type SecureValueHash interface {
	tl.TLObject
	_SecureValueHash()
}

var (
	_ SecureValueHash = (*SecureValueHashPredict)(nil)
)

type SecureValueHashPredict struct {
	Type SecureValueType `tl:"type"`
	Hash []byte          `tl:"hash"`
}

func (*SecureValueHashPredict) CRC() uint32 {
	return 0xed1ecdb0
}
func (*SecureValueHashPredict) _SecureValueHash() {}

type SecureValueType interface {
	tl.TLObject
	_SecureValueType()
}

var (
	_ SecureValueType = (*SecureValueTypePersonalDetailsPredict)(nil)
	_ SecureValueType = (*SecureValueTypePassportPredict)(nil)
	_ SecureValueType = (*SecureValueTypeDriverLicensePredict)(nil)
	_ SecureValueType = (*SecureValueTypeIdentityCardPredict)(nil)
	_ SecureValueType = (*SecureValueTypeInternalPassportPredict)(nil)
	_ SecureValueType = (*SecureValueTypeAddressPredict)(nil)
	_ SecureValueType = (*SecureValueTypeUtilityBillPredict)(nil)
	_ SecureValueType = (*SecureValueTypeBankStatementPredict)(nil)
	_ SecureValueType = (*SecureValueTypeRentalAgreementPredict)(nil)
	_ SecureValueType = (*SecureValueTypePassportRegistrationPredict)(nil)
	_ SecureValueType = (*SecureValueTypeTemporaryRegistrationPredict)(nil)
	_ SecureValueType = (*SecureValueTypePhonePredict)(nil)
	_ SecureValueType = (*SecureValueTypeEmailPredict)(nil)
)

type SecureValueTypePersonalDetailsPredict struct{}

func (*SecureValueTypePersonalDetailsPredict) CRC() uint32 {
	return 0x9d2a81e3
}
func (*SecureValueTypePersonalDetailsPredict) _SecureValueType() {}

type SecureValueTypePassportPredict struct{}

func (*SecureValueTypePassportPredict) CRC() uint32 {
	return 0x3dac6a00
}
func (*SecureValueTypePassportPredict) _SecureValueType() {}

type SecureValueTypeDriverLicensePredict struct{}

func (*SecureValueTypeDriverLicensePredict) CRC() uint32 {
	return 0x06e425c4
}
func (*SecureValueTypeDriverLicensePredict) _SecureValueType() {}

type SecureValueTypeIdentityCardPredict struct{}

func (*SecureValueTypeIdentityCardPredict) CRC() uint32 {
	return 0xa0d0744b
}
func (*SecureValueTypeIdentityCardPredict) _SecureValueType() {}

type SecureValueTypeInternalPassportPredict struct{}

func (*SecureValueTypeInternalPassportPredict) CRC() uint32 {
	return 0x99a48f23
}
func (*SecureValueTypeInternalPassportPredict) _SecureValueType() {}

type SecureValueTypeAddressPredict struct{}

func (*SecureValueTypeAddressPredict) CRC() uint32 {
	return 0xcbe31e26
}
func (*SecureValueTypeAddressPredict) _SecureValueType() {}

type SecureValueTypeUtilityBillPredict struct{}

func (*SecureValueTypeUtilityBillPredict) CRC() uint32 {
	return 0xfc36954e
}
func (*SecureValueTypeUtilityBillPredict) _SecureValueType() {}

type SecureValueTypeBankStatementPredict struct{}

func (*SecureValueTypeBankStatementPredict) CRC() uint32 {
	return 0x89137c0d
}
func (*SecureValueTypeBankStatementPredict) _SecureValueType() {}

type SecureValueTypeRentalAgreementPredict struct{}

func (*SecureValueTypeRentalAgreementPredict) CRC() uint32 {
	return 0x8b883488
}
func (*SecureValueTypeRentalAgreementPredict) _SecureValueType() {}

type SecureValueTypePassportRegistrationPredict struct{}

func (*SecureValueTypePassportRegistrationPredict) CRC() uint32 {
	return 0x99e3806a
}
func (*SecureValueTypePassportRegistrationPredict) _SecureValueType() {}

type SecureValueTypeTemporaryRegistrationPredict struct{}

func (*SecureValueTypeTemporaryRegistrationPredict) CRC() uint32 {
	return 0xea02ec33
}
func (*SecureValueTypeTemporaryRegistrationPredict) _SecureValueType() {}

type SecureValueTypePhonePredict struct{}

func (*SecureValueTypePhonePredict) CRC() uint32 {
	return 0xb320aadb
}
func (*SecureValueTypePhonePredict) _SecureValueType() {}

type SecureValueTypeEmailPredict struct{}

func (*SecureValueTypeEmailPredict) CRC() uint32 {
	return 0x8e3ca7ee
}
func (*SecureValueTypeEmailPredict) _SecureValueType() {}

type SendAsPeer interface {
	tl.TLObject
	_SendAsPeer()
}

var (
	_ SendAsPeer = (*SendAsPeerPredict)(nil)
)

type SendAsPeerPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	PremiumRequired bool     `tl:"premium_required,omitempty:flags:0,implicit"`
	Peer            Peer     `tl:"peer"`
}

func (*SendAsPeerPredict) CRC() uint32 {
	return 0xb81c7034
}
func (*SendAsPeerPredict) _SendAsPeer() {}

type SendMessageAction interface {
	tl.TLObject
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
	Progress int32 `tl:"progress"`
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
	Progress int32 `tl:"progress"`
}

func (*SendMessageUploadAudioActionPredict) CRC() uint32 {
	return 0xf351d7ab
}
func (*SendMessageUploadAudioActionPredict) _SendMessageAction() {}

type SendMessageUploadPhotoActionPredict struct {
	Progress int32 `tl:"progress"`
}

func (*SendMessageUploadPhotoActionPredict) CRC() uint32 {
	return 0xd1d34a26
}
func (*SendMessageUploadPhotoActionPredict) _SendMessageAction() {}

type SendMessageUploadDocumentActionPredict struct {
	Progress int32 `tl:"progress"`
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
	Progress int32 `tl:"progress"`
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
	Progress int32 `tl:"progress"`
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
	Emoticon    string   `tl:"emoticon"`
	MsgID       int32    `tl:"msg_id"`
	Interaction DataJSON `tl:"interaction"`
}

func (*SendMessageEmojiInteractionPredict) CRC() uint32 {
	return 0x25972bcb
}
func (*SendMessageEmojiInteractionPredict) _SendMessageAction() {}

type SendMessageEmojiInteractionSeenPredict struct {
	Emoticon string `tl:"emoticon"`
}

func (*SendMessageEmojiInteractionSeenPredict) CRC() uint32 {
	return 0xb665902e
}
func (*SendMessageEmojiInteractionSeenPredict) _SendMessageAction() {}

type ShippingOption interface {
	tl.TLObject
	_ShippingOption()
}

var (
	_ ShippingOption = (*ShippingOptionPredict)(nil)
)

type ShippingOptionPredict struct {
	ID     string         `tl:"id"`
	Title  string         `tl:"title"`
	Prices []LabeledPrice `tl:"prices"`
}

func (*ShippingOptionPredict) CRC() uint32 {
	return 0xb6213cdf
}
func (*ShippingOptionPredict) _ShippingOption() {}

type SmsJob interface {
	tl.TLObject
	_SmsJob()
}

var (
	_ SmsJob = (*SmsJobPredict)(nil)
)

type SmsJobPredict struct {
	JobID       string `tl:"job_id"`
	PhoneNumber string `tl:"phone_number"`
	Text        string `tl:"text"`
}

func (*SmsJobPredict) CRC() uint32 {
	return 0xe6a1eeb8
}
func (*SmsJobPredict) _SmsJob() {}

type SponsoredMessage interface {
	tl.TLObject
	_SponsoredMessage()
}

var (
	_ SponsoredMessage = (*SponsoredMessagePredict)(nil)
)

type SponsoredMessagePredict struct {
	_              struct{}        `tl:"flags,bitflag"`
	Recommended    bool            `tl:"recommended,omitempty:flags:5,implicit"`
	CanReport      bool            `tl:"can_report,omitempty:flags:12,implicit"`
	RandomID       []byte          `tl:"random_id"`
	URL            string          `tl:"url"`
	Title          string          `tl:"title"`
	Message        string          `tl:"message"`
	Entities       []MessageEntity `tl:"entities,omitempty:flags:1"`
	Photo          Photo           `tl:"photo,omitempty:flags:6"`
	Color          PeerColor       `tl:"color,omitempty:flags:13"`
	ButtonText     string          `tl:"button_text"`
	SponsorInfo    *string         `tl:"sponsor_info,omitempty:flags:7"`
	AdditionalInfo *string         `tl:"additional_info,omitempty:flags:8"`
}

func (*SponsoredMessagePredict) CRC() uint32 {
	return 0xbdedf566
}
func (*SponsoredMessagePredict) _SponsoredMessage() {}

type SponsoredMessageReportOption interface {
	tl.TLObject
	_SponsoredMessageReportOption()
}

var (
	_ SponsoredMessageReportOption = (*SponsoredMessageReportOptionPredict)(nil)
)

type SponsoredMessageReportOptionPredict struct {
	Text   string `tl:"text"`
	Option []byte `tl:"option"`
}

func (*SponsoredMessageReportOptionPredict) CRC() uint32 {
	return 0x430d3150
}
func (*SponsoredMessageReportOptionPredict) _SponsoredMessageReportOption() {}

type StarsGiftOption interface {
	tl.TLObject
	_StarsGiftOption()
}

var (
	_ StarsGiftOption = (*StarsGiftOptionPredict)(nil)
)

type StarsGiftOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Extended     bool     `tl:"extended,omitempty:flags:1,implicit"`
	Stars        int64    `tl:"stars"`
	StoreProduct *string  `tl:"store_product,omitempty:flags:0"`
	Currency     string   `tl:"currency"`
	Amount       int64    `tl:"amount"`
}

func (*StarsGiftOptionPredict) CRC() uint32 {
	return 0x5e0589f1
}
func (*StarsGiftOptionPredict) _StarsGiftOption() {}

type StarsRevenueStatus interface {
	tl.TLObject
	_StarsRevenueStatus()
}

var (
	_ StarsRevenueStatus = (*StarsRevenueStatusPredict)(nil)
)

type StarsRevenueStatusPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	WithdrawalEnabled bool     `tl:"withdrawal_enabled,omitempty:flags:0,implicit"`
	CurrentBalance    int64    `tl:"current_balance"`
	AvailableBalance  int64    `tl:"available_balance"`
	OverallRevenue    int64    `tl:"overall_revenue"`
	NextWithdrawalAt  *int32   `tl:"next_withdrawal_at,omitempty:flags:1"`
}

func (*StarsRevenueStatusPredict) CRC() uint32 {
	return 0x79342946
}
func (*StarsRevenueStatusPredict) _StarsRevenueStatus() {}

type StarsTopupOption interface {
	tl.TLObject
	_StarsTopupOption()
}

var (
	_ StarsTopupOption = (*StarsTopupOptionPredict)(nil)
)

type StarsTopupOptionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Extended     bool     `tl:"extended,omitempty:flags:1,implicit"`
	Stars        int64    `tl:"stars"`
	StoreProduct *string  `tl:"store_product,omitempty:flags:0"`
	Currency     string   `tl:"currency"`
	Amount       int64    `tl:"amount"`
}

func (*StarsTopupOptionPredict) CRC() uint32 {
	return 0x0bd915c0
}
func (*StarsTopupOptionPredict) _StarsTopupOption() {}

type StarsTransaction interface {
	tl.TLObject
	_StarsTransaction()
}

var (
	_ StarsTransaction = (*StarsTransactionPredict)(nil)
)

type StarsTransactionPredict struct {
	_               struct{}             `tl:"flags,bitflag"`
	Refund          bool                 `tl:"refund,omitempty:flags:3,implicit"`
	Pending         bool                 `tl:"pending,omitempty:flags:4,implicit"`
	Failed          bool                 `tl:"failed,omitempty:flags:6,implicit"`
	Gift            bool                 `tl:"gift,omitempty:flags:10,implicit"`
	ID              string               `tl:"id"`
	Stars           int64                `tl:"stars"`
	Date            int32                `tl:"date"`
	Peer            StarsTransactionPeer `tl:"peer"`
	Title           *string              `tl:"title,omitempty:flags:0"`
	Description     *string              `tl:"description,omitempty:flags:1"`
	Photo           WebDocument          `tl:"photo,omitempty:flags:2"`
	TransactionDate *int32               `tl:"transaction_date,omitempty:flags:5"`
	TransactionURL  *string              `tl:"transaction_url,omitempty:flags:5"`
	BotPayload      *[]byte              `tl:"bot_payload,omitempty:flags:7"`
	MsgID           *int32               `tl:"msg_id,omitempty:flags:8"`
	ExtendedMedia   []MessageMedia       `tl:"extended_media,omitempty:flags:9"`
}

func (*StarsTransactionPredict) CRC() uint32 {
	return 0x2db5418f
}
func (*StarsTransactionPredict) _StarsTransaction() {}

type StarsTransactionPeer interface {
	tl.TLObject
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
	Peer Peer `tl:"peer"`
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
	tl.TLObject
	_StatsAbsValueAndPrev()
}

var (
	_ StatsAbsValueAndPrev = (*StatsAbsValueAndPrevPredict)(nil)
)

type StatsAbsValueAndPrevPredict struct {
	Current  float64 `tl:"current"`
	Previous float64 `tl:"previous"`
}

func (*StatsAbsValueAndPrevPredict) CRC() uint32 {
	return 0xcb43acde
}
func (*StatsAbsValueAndPrevPredict) _StatsAbsValueAndPrev() {}

type StatsDateRangeDays interface {
	tl.TLObject
	_StatsDateRangeDays()
}

var (
	_ StatsDateRangeDays = (*StatsDateRangeDaysPredict)(nil)
)

type StatsDateRangeDaysPredict struct {
	MinDate int32 `tl:"min_date"`
	MaxDate int32 `tl:"max_date"`
}

func (*StatsDateRangeDaysPredict) CRC() uint32 {
	return 0xb637edaf
}
func (*StatsDateRangeDaysPredict) _StatsDateRangeDays() {}

type StatsGraph interface {
	tl.TLObject
	_StatsGraph()
}

var (
	_ StatsGraph = (*StatsGraphAsyncPredict)(nil)
	_ StatsGraph = (*StatsGraphErrorPredict)(nil)
	_ StatsGraph = (*StatsGraphPredict)(nil)
)

type StatsGraphAsyncPredict struct {
	Token string `tl:"token"`
}

func (*StatsGraphAsyncPredict) CRC() uint32 {
	return 0x4a27eb2d
}
func (*StatsGraphAsyncPredict) _StatsGraph() {}

type StatsGraphErrorPredict struct {
	Error string `tl:"error"`
}

func (*StatsGraphErrorPredict) CRC() uint32 {
	return 0xbedc9822
}
func (*StatsGraphErrorPredict) _StatsGraph() {}

type StatsGraphPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	JSON      DataJSON `tl:"json"`
	ZoomToken *string  `tl:"zoom_token,omitempty:flags:0"`
}

func (*StatsGraphPredict) CRC() uint32 {
	return 0x8ea464b6
}
func (*StatsGraphPredict) _StatsGraph() {}

type StatsGroupTopAdmin interface {
	tl.TLObject
	_StatsGroupTopAdmin()
}

var (
	_ StatsGroupTopAdmin = (*StatsGroupTopAdminPredict)(nil)
)

type StatsGroupTopAdminPredict struct {
	UserID  int64 `tl:"user_id"`
	Deleted int32 `tl:"deleted"`
	Kicked  int32 `tl:"kicked"`
	Banned  int32 `tl:"banned"`
}

func (*StatsGroupTopAdminPredict) CRC() uint32 {
	return 0xd7584c87
}
func (*StatsGroupTopAdminPredict) _StatsGroupTopAdmin() {}

type StatsGroupTopInviter interface {
	tl.TLObject
	_StatsGroupTopInviter()
}

var (
	_ StatsGroupTopInviter = (*StatsGroupTopInviterPredict)(nil)
)

type StatsGroupTopInviterPredict struct {
	UserID      int64 `tl:"user_id"`
	Invitations int32 `tl:"invitations"`
}

func (*StatsGroupTopInviterPredict) CRC() uint32 {
	return 0x535f779d
}
func (*StatsGroupTopInviterPredict) _StatsGroupTopInviter() {}

type StatsGroupTopPoster interface {
	tl.TLObject
	_StatsGroupTopPoster()
}

var (
	_ StatsGroupTopPoster = (*StatsGroupTopPosterPredict)(nil)
)

type StatsGroupTopPosterPredict struct {
	UserID   int64 `tl:"user_id"`
	Messages int32 `tl:"messages"`
	AvgChars int32 `tl:"avg_chars"`
}

func (*StatsGroupTopPosterPredict) CRC() uint32 {
	return 0x9d04af9b
}
func (*StatsGroupTopPosterPredict) _StatsGroupTopPoster() {}

type StatsPercentValue interface {
	tl.TLObject
	_StatsPercentValue()
}

var (
	_ StatsPercentValue = (*StatsPercentValuePredict)(nil)
)

type StatsPercentValuePredict struct {
	Part  float64 `tl:"part"`
	Total float64 `tl:"total"`
}

func (*StatsPercentValuePredict) CRC() uint32 {
	return 0xcbce2fe0
}
func (*StatsPercentValuePredict) _StatsPercentValue() {}

type StatsURL interface {
	tl.TLObject
	_StatsURL()
}

var (
	_ StatsURL = (*StatsURLPredict)(nil)
)

type StatsURLPredict struct {
	URL string `tl:"url"`
}

func (*StatsURLPredict) CRC() uint32 {
	return 0x47a971e0
}
func (*StatsURLPredict) _StatsURL() {}

type StickerKeyword interface {
	tl.TLObject
	_StickerKeyword()
}

var (
	_ StickerKeyword = (*StickerKeywordPredict)(nil)
)

type StickerKeywordPredict struct {
	DocumentID int64    `tl:"document_id"`
	Keyword    []string `tl:"keyword"`
}

func (*StickerKeywordPredict) CRC() uint32 {
	return 0xfcfeb29c
}
func (*StickerKeywordPredict) _StickerKeyword() {}

type StickerPack interface {
	tl.TLObject
	_StickerPack()
}

var (
	_ StickerPack = (*StickerPackPredict)(nil)
)

type StickerPackPredict struct {
	Emoticon  string  `tl:"emoticon"`
	Documents []int64 `tl:"documents"`
}

func (*StickerPackPredict) CRC() uint32 {
	return 0x12b299d4
}
func (*StickerPackPredict) _StickerPack() {}

type StickerSet interface {
	tl.TLObject
	_StickerSet()
}

var (
	_ StickerSet = (*StickerSetPredict)(nil)
)

type StickerSetPredict struct {
	_                  struct{}    `tl:"flags,bitflag"`
	Archived           bool        `tl:"archived,omitempty:flags:1,implicit"`
	Official           bool        `tl:"official,omitempty:flags:2,implicit"`
	Masks              bool        `tl:"masks,omitempty:flags:3,implicit"`
	Emojis             bool        `tl:"emojis,omitempty:flags:7,implicit"`
	TextColor          bool        `tl:"text_color,omitempty:flags:9,implicit"`
	ChannelEmojiStatus bool        `tl:"channel_emoji_status,omitempty:flags:10,implicit"`
	Creator            bool        `tl:"creator,omitempty:flags:11,implicit"`
	InstalledDate      *int32      `tl:"installed_date,omitempty:flags:0"`
	ID                 int64       `tl:"id"`
	AccessHash         int64       `tl:"access_hash"`
	Title              string      `tl:"title"`
	ShortName          string      `tl:"short_name"`
	Thumbs             []PhotoSize `tl:"thumbs,omitempty:flags:4"`
	ThumbDcID          *int32      `tl:"thumb_dc_id,omitempty:flags:4"`
	ThumbVersion       *int32      `tl:"thumb_version,omitempty:flags:4"`
	ThumbDocumentID    *int64      `tl:"thumb_document_id,omitempty:flags:8"`
	Count              int32       `tl:"count"`
	Hash               int32       `tl:"hash"`
}

func (*StickerSetPredict) CRC() uint32 {
	return 0x2dd14edc
}
func (*StickerSetPredict) _StickerSet() {}

type StickerSetCovered interface {
	tl.TLObject
	_StickerSetCovered()
}

var (
	_ StickerSetCovered = (*StickerSetCoveredPredict)(nil)
	_ StickerSetCovered = (*StickerSetMultiCoveredPredict)(nil)
	_ StickerSetCovered = (*StickerSetFullCoveredPredict)(nil)
	_ StickerSetCovered = (*StickerSetNoCoveredPredict)(nil)
)

type StickerSetCoveredPredict struct {
	Set   StickerSet `tl:"set"`
	Cover Document   `tl:"cover"`
}

func (*StickerSetCoveredPredict) CRC() uint32 {
	return 0x6410a5d2
}
func (*StickerSetCoveredPredict) _StickerSetCovered() {}

type StickerSetMultiCoveredPredict struct {
	Set    StickerSet `tl:"set"`
	Covers []Document `tl:"covers"`
}

func (*StickerSetMultiCoveredPredict) CRC() uint32 {
	return 0x3407e51b
}
func (*StickerSetMultiCoveredPredict) _StickerSetCovered() {}

type StickerSetFullCoveredPredict struct {
	Set       StickerSet       `tl:"set"`
	Packs     []StickerPack    `tl:"packs"`
	Keywords  []StickerKeyword `tl:"keywords"`
	Documents []Document       `tl:"documents"`
}

func (*StickerSetFullCoveredPredict) CRC() uint32 {
	return 0x40d13c0e
}
func (*StickerSetFullCoveredPredict) _StickerSetCovered() {}

type StickerSetNoCoveredPredict struct {
	Set StickerSet `tl:"set"`
}

func (*StickerSetNoCoveredPredict) CRC() uint32 {
	return 0x77b15d1c
}
func (*StickerSetNoCoveredPredict) _StickerSetCovered() {}

type StoriesStealthMode interface {
	tl.TLObject
	_StoriesStealthMode()
}

var (
	_ StoriesStealthMode = (*StoriesStealthModePredict)(nil)
)

type StoriesStealthModePredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	ActiveUntilDate   *int32   `tl:"active_until_date,omitempty:flags:0"`
	CooldownUntilDate *int32   `tl:"cooldown_until_date,omitempty:flags:1"`
}

func (*StoriesStealthModePredict) CRC() uint32 {
	return 0x712e27fd
}
func (*StoriesStealthModePredict) _StoriesStealthMode() {}

type StoryFwdHeader interface {
	tl.TLObject
	_StoryFwdHeader()
}

var (
	_ StoryFwdHeader = (*StoryFwdHeaderPredict)(nil)
)

type StoryFwdHeaderPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Modified bool     `tl:"modified,omitempty:flags:3,implicit"`
	From     Peer     `tl:"from,omitempty:flags:0"`
	FromName *string  `tl:"from_name,omitempty:flags:1"`
	StoryID  *int32   `tl:"story_id,omitempty:flags:2"`
}

func (*StoryFwdHeaderPredict) CRC() uint32 {
	return 0xb826e150
}
func (*StoryFwdHeaderPredict) _StoryFwdHeader() {}

type StoryItem interface {
	tl.TLObject
	_StoryItem()
}

var (
	_ StoryItem = (*StoryItemDeletedPredict)(nil)
	_ StoryItem = (*StoryItemSkippedPredict)(nil)
	_ StoryItem = (*StoryItemPredict)(nil)
)

type StoryItemDeletedPredict struct {
	ID int32 `tl:"id"`
}

func (*StoryItemDeletedPredict) CRC() uint32 {
	return 0x51e6ee4f
}
func (*StoryItemDeletedPredict) _StoryItem() {}

type StoryItemSkippedPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	CloseFriends bool     `tl:"close_friends,omitempty:flags:8,implicit"`
	ID           int32    `tl:"id"`
	Date         int32    `tl:"date"`
	ExpireDate   int32    `tl:"expire_date"`
}

func (*StoryItemSkippedPredict) CRC() uint32 {
	return 0xffadc913
}
func (*StoryItemSkippedPredict) _StoryItem() {}

type StoryItemPredict struct {
	_                struct{}        `tl:"flags,bitflag"`
	Pinned           bool            `tl:"pinned,omitempty:flags:5,implicit"`
	Public           bool            `tl:"public,omitempty:flags:7,implicit"`
	CloseFriends     bool            `tl:"close_friends,omitempty:flags:8,implicit"`
	Min              bool            `tl:"min,omitempty:flags:9,implicit"`
	Noforwards       bool            `tl:"noforwards,omitempty:flags:10,implicit"`
	Edited           bool            `tl:"edited,omitempty:flags:11,implicit"`
	Contacts         bool            `tl:"contacts,omitempty:flags:12,implicit"`
	SelectedContacts bool            `tl:"selected_contacts,omitempty:flags:13,implicit"`
	Out              bool            `tl:"out,omitempty:flags:16,implicit"`
	ID               int32           `tl:"id"`
	Date             int32           `tl:"date"`
	FromID           Peer            `tl:"from_id,omitempty:flags:18"`
	FwdFrom          StoryFwdHeader  `tl:"fwd_from,omitempty:flags:17"`
	ExpireDate       int32           `tl:"expire_date"`
	Caption          *string         `tl:"caption,omitempty:flags:0"`
	Entities         []MessageEntity `tl:"entities,omitempty:flags:1"`
	Media            MessageMedia    `tl:"media"`
	MediaAreas       []MediaArea     `tl:"media_areas,omitempty:flags:14"`
	Privacy          []PrivacyRule   `tl:"privacy,omitempty:flags:2"`
	Views            StoryViews      `tl:"views,omitempty:flags:3"`
	SentReaction     Reaction        `tl:"sent_reaction,omitempty:flags:15"`
}

func (*StoryItemPredict) CRC() uint32 {
	return 0x79b26a24
}
func (*StoryItemPredict) _StoryItem() {}

type StoryReaction interface {
	tl.TLObject
	_StoryReaction()
}

var (
	_ StoryReaction = (*StoryReactionPredict)(nil)
	_ StoryReaction = (*StoryReactionPublicForwardPredict)(nil)
	_ StoryReaction = (*StoryReactionPublicRepostPredict)(nil)
)

type StoryReactionPredict struct {
	PeerID   Peer     `tl:"peer_id"`
	Date     int32    `tl:"date"`
	Reaction Reaction `tl:"reaction"`
}

func (*StoryReactionPredict) CRC() uint32 {
	return 0x6090d6d5
}
func (*StoryReactionPredict) _StoryReaction() {}

type StoryReactionPublicForwardPredict struct {
	Message Message `tl:"message"`
}

func (*StoryReactionPublicForwardPredict) CRC() uint32 {
	return 0xbbab2643
}
func (*StoryReactionPublicForwardPredict) _StoryReaction() {}

type StoryReactionPublicRepostPredict struct {
	PeerID Peer      `tl:"peer_id"`
	Story  StoryItem `tl:"story"`
}

func (*StoryReactionPublicRepostPredict) CRC() uint32 {
	return 0xcfcd0f13
}
func (*StoryReactionPublicRepostPredict) _StoryReaction() {}

type StoryView interface {
	tl.TLObject
	_StoryView()
}

var (
	_ StoryView = (*StoryViewPredict)(nil)
	_ StoryView = (*StoryViewPublicForwardPredict)(nil)
	_ StoryView = (*StoryViewPublicRepostPredict)(nil)
)

type StoryViewPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:"blocked,omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:"blocked_my_stories_from,omitempty:flags:1,implicit"`
	UserID               int64    `tl:"user_id"`
	Date                 int32    `tl:"date"`
	Reaction             Reaction `tl:"reaction,omitempty:flags:2"`
}

func (*StoryViewPredict) CRC() uint32 {
	return 0xb0bdeac5
}
func (*StoryViewPredict) _StoryView() {}

type StoryViewPublicForwardPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:"blocked,omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:"blocked_my_stories_from,omitempty:flags:1,implicit"`
	Message              Message  `tl:"message"`
}

func (*StoryViewPublicForwardPredict) CRC() uint32 {
	return 0x9083670b
}
func (*StoryViewPublicForwardPredict) _StoryView() {}

type StoryViewPublicRepostPredict struct {
	_                    struct{}  `tl:"flags,bitflag"`
	Blocked              bool      `tl:"blocked,omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool      `tl:"blocked_my_stories_from,omitempty:flags:1,implicit"`
	PeerID               Peer      `tl:"peer_id"`
	Story                StoryItem `tl:"story"`
}

func (*StoryViewPublicRepostPredict) CRC() uint32 {
	return 0xbd74cf49
}
func (*StoryViewPublicRepostPredict) _StoryView() {}

type StoryViews interface {
	tl.TLObject
	_StoryViews()
}

var (
	_ StoryViews = (*StoryViewsPredict)(nil)
)

type StoryViewsPredict struct {
	_              struct{}        `tl:"flags,bitflag"`
	HasViewers     bool            `tl:"has_viewers,omitempty:flags:1,implicit"`
	ViewsCount     int32           `tl:"views_count"`
	ForwardsCount  *int32          `tl:"forwards_count,omitempty:flags:2"`
	Reactions      []ReactionCount `tl:"reactions,omitempty:flags:3"`
	ReactionsCount *int32          `tl:"reactions_count,omitempty:flags:4"`
	RecentViewers  []int64         `tl:"recent_viewers,omitempty:flags:0"`
}

func (*StoryViewsPredict) CRC() uint32 {
	return 0x8d595cd6
}
func (*StoryViewsPredict) _StoryViews() {}

type TextWithEntities interface {
	tl.TLObject
	_TextWithEntities()
}

var (
	_ TextWithEntities = (*TextWithEntitiesPredict)(nil)
)

type TextWithEntitiesPredict struct {
	Text     string          `tl:"text"`
	Entities []MessageEntity `tl:"entities"`
}

func (*TextWithEntitiesPredict) CRC() uint32 {
	return 0x751f3146
}
func (*TextWithEntitiesPredict) _TextWithEntities() {}

type Theme interface {
	tl.TLObject
	_Theme()
}

var (
	_ Theme = (*ThemePredict)(nil)
)

type ThemePredict struct {
	_             struct{}        `tl:"flags,bitflag"`
	Creator       bool            `tl:"creator,omitempty:flags:0,implicit"`
	Default       bool            `tl:"default,omitempty:flags:1,implicit"`
	ForChat       bool            `tl:"for_chat,omitempty:flags:5,implicit"`
	ID            int64           `tl:"id"`
	AccessHash    int64           `tl:"access_hash"`
	Slug          string          `tl:"slug"`
	Title         string          `tl:"title"`
	Document      Document        `tl:"document,omitempty:flags:2"`
	Settings      []ThemeSettings `tl:"settings,omitempty:flags:3"`
	Emoticon      *string         `tl:"emoticon,omitempty:flags:6"`
	InstallsCount *int32          `tl:"installs_count,omitempty:flags:4"`
}

func (*ThemePredict) CRC() uint32 {
	return 0xa00e67d6
}
func (*ThemePredict) _Theme() {}

type ThemeSettings interface {
	tl.TLObject
	_ThemeSettings()
}

var (
	_ ThemeSettings = (*ThemeSettingsPredict)(nil)
)

type ThemeSettingsPredict struct {
	_                     struct{}  `tl:"flags,bitflag"`
	MessageColorsAnimated bool      `tl:"message_colors_animated,omitempty:flags:2,implicit"`
	BaseTheme             BaseTheme `tl:"base_theme"`
	AccentColor           int32     `tl:"accent_color"`
	OutboxAccentColor     *int32    `tl:"outbox_accent_color,omitempty:flags:3"`
	MessageColors         []int32   `tl:"message_colors,omitempty:flags:0"`
	Wallpaper             WallPaper `tl:"wallpaper,omitempty:flags:1"`
}

func (*ThemeSettingsPredict) CRC() uint32 {
	return 0xfa58b6d4
}
func (*ThemeSettingsPredict) _ThemeSettings() {}

type Timezone interface {
	tl.TLObject
	_Timezone()
}

var (
	_ Timezone = (*TimezonePredict)(nil)
)

type TimezonePredict struct {
	ID        string `tl:"id"`
	Name      string `tl:"name"`
	UtcOffset int32  `tl:"utc_offset"`
}

func (*TimezonePredict) CRC() uint32 {
	return 0xff9289f5
}
func (*TimezonePredict) _Timezone() {}

type TopPeer interface {
	tl.TLObject
	_TopPeer()
}

var (
	_ TopPeer = (*TopPeerPredict)(nil)
)

type TopPeerPredict struct {
	Peer   Peer    `tl:"peer"`
	Rating float64 `tl:"rating"`
}

func (*TopPeerPredict) CRC() uint32 {
	return 0xedcdc05b
}
func (*TopPeerPredict) _TopPeer() {}

type TopPeerCategory interface {
	tl.TLObject
	_TopPeerCategory()
}

var (
	_ TopPeerCategory = (*TopPeerCategoryBotsPmPredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryBotsInlinePredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryCorrespondentsPredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryGroupsPredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryChannelsPredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryPhoneCallsPredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryForwardUsersPredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryForwardChatsPredict)(nil)
	_ TopPeerCategory = (*TopPeerCategoryBotsAppPredict)(nil)
)

type TopPeerCategoryBotsPmPredict struct{}

func (*TopPeerCategoryBotsPmPredict) CRC() uint32 {
	return 0xab661b5b
}
func (*TopPeerCategoryBotsPmPredict) _TopPeerCategory() {}

type TopPeerCategoryBotsInlinePredict struct{}

func (*TopPeerCategoryBotsInlinePredict) CRC() uint32 {
	return 0x148677e2
}
func (*TopPeerCategoryBotsInlinePredict) _TopPeerCategory() {}

type TopPeerCategoryCorrespondentsPredict struct{}

func (*TopPeerCategoryCorrespondentsPredict) CRC() uint32 {
	return 0x0637b7ed
}
func (*TopPeerCategoryCorrespondentsPredict) _TopPeerCategory() {}

type TopPeerCategoryGroupsPredict struct{}

func (*TopPeerCategoryGroupsPredict) CRC() uint32 {
	return 0xbd17a14a
}
func (*TopPeerCategoryGroupsPredict) _TopPeerCategory() {}

type TopPeerCategoryChannelsPredict struct{}

func (*TopPeerCategoryChannelsPredict) CRC() uint32 {
	return 0x161d9628
}
func (*TopPeerCategoryChannelsPredict) _TopPeerCategory() {}

type TopPeerCategoryPhoneCallsPredict struct{}

func (*TopPeerCategoryPhoneCallsPredict) CRC() uint32 {
	return 0x1e76a78c
}
func (*TopPeerCategoryPhoneCallsPredict) _TopPeerCategory() {}

type TopPeerCategoryForwardUsersPredict struct{}

func (*TopPeerCategoryForwardUsersPredict) CRC() uint32 {
	return 0xa8406ca9
}
func (*TopPeerCategoryForwardUsersPredict) _TopPeerCategory() {}

type TopPeerCategoryForwardChatsPredict struct{}

func (*TopPeerCategoryForwardChatsPredict) CRC() uint32 {
	return 0xfbeec0f0
}
func (*TopPeerCategoryForwardChatsPredict) _TopPeerCategory() {}

type TopPeerCategoryBotsAppPredict struct{}

func (*TopPeerCategoryBotsAppPredict) CRC() uint32 {
	return 0xfd9e7bec
}
func (*TopPeerCategoryBotsAppPredict) _TopPeerCategory() {}

type TopPeerCategoryPeers interface {
	tl.TLObject
	_TopPeerCategoryPeers()
}

var (
	_ TopPeerCategoryPeers = (*TopPeerCategoryPeersPredict)(nil)
)

type TopPeerCategoryPeersPredict struct {
	Category TopPeerCategory `tl:"category"`
	Count    int32           `tl:"count"`
	Peers    []TopPeer       `tl:"peers"`
}

func (*TopPeerCategoryPeersPredict) CRC() uint32 {
	return 0xfb834291
}
func (*TopPeerCategoryPeersPredict) _TopPeerCategoryPeers() {}

type True interface {
	tl.TLObject
	_True()
}

var (
	_ True = (*TruePredict)(nil)
)

type TruePredict struct{}

func (*TruePredict) CRC() uint32 {
	return 0x3fedd339
}
func (*TruePredict) _True() {}

type Update interface {
	tl.TLObject
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
	Message  Message `tl:"message"`
	Pts      int32   `tl:"pts"`
	PtsCount int32   `tl:"pts_count"`
}

func (*UpdateNewMessagePredict) CRC() uint32 {
	return 0x1f2b0afd
}
func (*UpdateNewMessagePredict) _Update() {}

type UpdateMessageIDPredict struct {
	ID       int32 `tl:"id"`
	RandomID int64 `tl:"random_id"`
}

func (*UpdateMessageIDPredict) CRC() uint32 {
	return 0x4e90bfd6
}
func (*UpdateMessageIDPredict) _Update() {}

type UpdateDeleteMessagesPredict struct {
	Messages []int32 `tl:"messages"`
	Pts      int32   `tl:"pts"`
	PtsCount int32   `tl:"pts_count"`
}

func (*UpdateDeleteMessagesPredict) CRC() uint32 {
	return 0xa20db0e5
}
func (*UpdateDeleteMessagesPredict) _Update() {}

type UpdateUserTypingPredict struct {
	UserID int64             `tl:"user_id"`
	Action SendMessageAction `tl:"action"`
}

func (*UpdateUserTypingPredict) CRC() uint32 {
	return 0xc01e857f
}
func (*UpdateUserTypingPredict) _Update() {}

type UpdateChatUserTypingPredict struct {
	ChatID int64             `tl:"chat_id"`
	FromID Peer              `tl:"from_id"`
	Action SendMessageAction `tl:"action"`
}

func (*UpdateChatUserTypingPredict) CRC() uint32 {
	return 0x83487af0
}
func (*UpdateChatUserTypingPredict) _Update() {}

type UpdateChatParticipantsPredict struct {
	Participants ChatParticipants `tl:"participants"`
}

func (*UpdateChatParticipantsPredict) CRC() uint32 {
	return 0x07761198
}
func (*UpdateChatParticipantsPredict) _Update() {}

type UpdateUserStatusPredict struct {
	UserID int64      `tl:"user_id"`
	Status UserStatus `tl:"status"`
}

func (*UpdateUserStatusPredict) CRC() uint32 {
	return 0xe5bdf8de
}
func (*UpdateUserStatusPredict) _Update() {}

type UpdateUserNamePredict struct {
	UserID    int64      `tl:"user_id"`
	FirstName string     `tl:"first_name"`
	LastName  string     `tl:"last_name"`
	Usernames []Username `tl:"usernames"`
}

func (*UpdateUserNamePredict) CRC() uint32 {
	return 0xa7848924
}
func (*UpdateUserNamePredict) _Update() {}

type UpdateNewAuthorizationPredict struct {
	_           struct{} `tl:"flags,bitflag"`
	Unconfirmed bool     `tl:"unconfirmed,omitempty:flags:0,implicit"`
	Hash        int64    `tl:"hash"`
	Date        *int32   `tl:"date,omitempty:flags:0"`
	Device      *string  `tl:"device,omitempty:flags:0"`
	Location    *string  `tl:"location,omitempty:flags:0"`
}

func (*UpdateNewAuthorizationPredict) CRC() uint32 {
	return 0x8951abef
}
func (*UpdateNewAuthorizationPredict) _Update() {}

type UpdateNewEncryptedMessagePredict struct {
	Message EncryptedMessage `tl:"message"`
	Qts     int32            `tl:"qts"`
}

func (*UpdateNewEncryptedMessagePredict) CRC() uint32 {
	return 0x12bcbd9a
}
func (*UpdateNewEncryptedMessagePredict) _Update() {}

type UpdateEncryptedChatTypingPredict struct {
	ChatID int32 `tl:"chat_id"`
}

func (*UpdateEncryptedChatTypingPredict) CRC() uint32 {
	return 0x1710f156
}
func (*UpdateEncryptedChatTypingPredict) _Update() {}

type UpdateEncryptionPredict struct {
	Chat EncryptedChat `tl:"chat"`
	Date int32         `tl:"date"`
}

func (*UpdateEncryptionPredict) CRC() uint32 {
	return 0xb4a2e88d
}
func (*UpdateEncryptionPredict) _Update() {}

type UpdateEncryptedMessagesReadPredict struct {
	ChatID  int32 `tl:"chat_id"`
	MaxDate int32 `tl:"max_date"`
	Date    int32 `tl:"date"`
}

func (*UpdateEncryptedMessagesReadPredict) CRC() uint32 {
	return 0x38fe25b7
}
func (*UpdateEncryptedMessagesReadPredict) _Update() {}

type UpdateChatParticipantAddPredict struct {
	ChatID    int64 `tl:"chat_id"`
	UserID    int64 `tl:"user_id"`
	InviterID int64 `tl:"inviter_id"`
	Date      int32 `tl:"date"`
	Version   int32 `tl:"version"`
}

func (*UpdateChatParticipantAddPredict) CRC() uint32 {
	return 0x3dda5451
}
func (*UpdateChatParticipantAddPredict) _Update() {}

type UpdateChatParticipantDeletePredict struct {
	ChatID  int64 `tl:"chat_id"`
	UserID  int64 `tl:"user_id"`
	Version int32 `tl:"version"`
}

func (*UpdateChatParticipantDeletePredict) CRC() uint32 {
	return 0xe32f3d77
}
func (*UpdateChatParticipantDeletePredict) _Update() {}

type UpdateDcOptionsPredict struct {
	DcOptions []DcOption `tl:"dc_options"`
}

func (*UpdateDcOptionsPredict) CRC() uint32 {
	return 0x8e5e9873
}
func (*UpdateDcOptionsPredict) _Update() {}

type UpdateNotifySettingsPredict struct {
	Peer           NotifyPeer         `tl:"peer"`
	NotifySettings PeerNotifySettings `tl:"notify_settings"`
}

func (*UpdateNotifySettingsPredict) CRC() uint32 {
	return 0xbec268ef
}
func (*UpdateNotifySettingsPredict) _Update() {}

type UpdateServiceNotificationPredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	Popup       bool            `tl:"popup,omitempty:flags:0,implicit"`
	InvertMedia bool            `tl:"invert_media,omitempty:flags:2,implicit"`
	InboxDate   *int32          `tl:"inbox_date,omitempty:flags:1"`
	Type        string          `tl:"type"`
	Message     string          `tl:"message"`
	Media       MessageMedia    `tl:"media"`
	Entities    []MessageEntity `tl:"entities"`
}

func (*UpdateServiceNotificationPredict) CRC() uint32 {
	return 0xebe46819
}
func (*UpdateServiceNotificationPredict) _Update() {}

type UpdatePrivacyPredict struct {
	Key   PrivacyKey    `tl:"key"`
	Rules []PrivacyRule `tl:"rules"`
}

func (*UpdatePrivacyPredict) CRC() uint32 {
	return 0xee3b272a
}
func (*UpdatePrivacyPredict) _Update() {}

type UpdateUserPhonePredict struct {
	UserID int64  `tl:"user_id"`
	Phone  string `tl:"phone"`
}

func (*UpdateUserPhonePredict) CRC() uint32 {
	return 0x05492a13
}
func (*UpdateUserPhonePredict) _Update() {}

type UpdateReadHistoryInboxPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	FolderID         *int32   `tl:"folder_id,omitempty:flags:0"`
	Peer             Peer     `tl:"peer"`
	MaxID            int32    `tl:"max_id"`
	StillUnreadCount int32    `tl:"still_unread_count"`
	Pts              int32    `tl:"pts"`
	PtsCount         int32    `tl:"pts_count"`
}

func (*UpdateReadHistoryInboxPredict) CRC() uint32 {
	return 0x9c974fdf
}
func (*UpdateReadHistoryInboxPredict) _Update() {}

type UpdateReadHistoryOutboxPredict struct {
	Peer     Peer  `tl:"peer"`
	MaxID    int32 `tl:"max_id"`
	Pts      int32 `tl:"pts"`
	PtsCount int32 `tl:"pts_count"`
}

func (*UpdateReadHistoryOutboxPredict) CRC() uint32 {
	return 0x2f2f21bf
}
func (*UpdateReadHistoryOutboxPredict) _Update() {}

type UpdateWebPagePredict struct {
	Webpage  WebPage `tl:"webpage"`
	Pts      int32   `tl:"pts"`
	PtsCount int32   `tl:"pts_count"`
}

func (*UpdateWebPagePredict) CRC() uint32 {
	return 0x7f891213
}
func (*UpdateWebPagePredict) _Update() {}

type UpdateReadMessagesContentsPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Messages []int32  `tl:"messages"`
	Pts      int32    `tl:"pts"`
	PtsCount int32    `tl:"pts_count"`
	Date     *int32   `tl:"date,omitempty:flags:0"`
}

func (*UpdateReadMessagesContentsPredict) CRC() uint32 {
	return 0xf8227181
}
func (*UpdateReadMessagesContentsPredict) _Update() {}

type UpdateChannelTooLongPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64    `tl:"channel_id"`
	Pts       *int32   `tl:"pts,omitempty:flags:0"`
}

func (*UpdateChannelTooLongPredict) CRC() uint32 {
	return 0x108d941f
}
func (*UpdateChannelTooLongPredict) _Update() {}

type UpdateChannelPredict struct {
	ChannelID int64 `tl:"channel_id"`
}

func (*UpdateChannelPredict) CRC() uint32 {
	return 0x635b4c09
}
func (*UpdateChannelPredict) _Update() {}

type UpdateNewChannelMessagePredict struct {
	Message  Message `tl:"message"`
	Pts      int32   `tl:"pts"`
	PtsCount int32   `tl:"pts_count"`
}

func (*UpdateNewChannelMessagePredict) CRC() uint32 {
	return 0x62ba04d9
}
func (*UpdateNewChannelMessagePredict) _Update() {}

type UpdateReadChannelInboxPredict struct {
	_                struct{} `tl:"flags,bitflag"`
	FolderID         *int32   `tl:"folder_id,omitempty:flags:0"`
	ChannelID        int64    `tl:"channel_id"`
	MaxID            int32    `tl:"max_id"`
	StillUnreadCount int32    `tl:"still_unread_count"`
	Pts              int32    `tl:"pts"`
}

func (*UpdateReadChannelInboxPredict) CRC() uint32 {
	return 0x922e6e10
}
func (*UpdateReadChannelInboxPredict) _Update() {}

type UpdateDeleteChannelMessagesPredict struct {
	ChannelID int64   `tl:"channel_id"`
	Messages  []int32 `tl:"messages"`
	Pts       int32   `tl:"pts"`
	PtsCount  int32   `tl:"pts_count"`
}

func (*UpdateDeleteChannelMessagesPredict) CRC() uint32 {
	return 0xc32d5b12
}
func (*UpdateDeleteChannelMessagesPredict) _Update() {}

type UpdateChannelMessageViewsPredict struct {
	ChannelID int64 `tl:"channel_id"`
	ID        int32 `tl:"id"`
	Views     int32 `tl:"views"`
}

func (*UpdateChannelMessageViewsPredict) CRC() uint32 {
	return 0xf226ac08
}
func (*UpdateChannelMessageViewsPredict) _Update() {}

type UpdateChatParticipantAdminPredict struct {
	ChatID  int64 `tl:"chat_id"`
	UserID  int64 `tl:"user_id"`
	IsAdmin bool  `tl:"is_admin"`
	Version int32 `tl:"version"`
}

func (*UpdateChatParticipantAdminPredict) CRC() uint32 {
	return 0xd7ca61a2
}
func (*UpdateChatParticipantAdminPredict) _Update() {}

type UpdateNewStickerSetPredict struct {
	Stickerset MessagesStickerSet `tl:"stickerset"`
}

func (*UpdateNewStickerSetPredict) CRC() uint32 {
	return 0x688a30aa
}
func (*UpdateNewStickerSetPredict) _Update() {}

type UpdateStickerSetsOrderPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:"masks,omitempty:flags:0,implicit"`
	Emojis bool     `tl:"emojis,omitempty:flags:1,implicit"`
	Order  []int64  `tl:"order"`
}

func (*UpdateStickerSetsOrderPredict) CRC() uint32 {
	return 0x0bb2d201
}
func (*UpdateStickerSetsOrderPredict) _Update() {}

type UpdateStickerSetsPredict struct {
	_      struct{} `tl:"flags,bitflag"`
	Masks  bool     `tl:"masks,omitempty:flags:0,implicit"`
	Emojis bool     `tl:"emojis,omitempty:flags:1,implicit"`
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
	_        struct{}            `tl:"flags,bitflag"`
	QueryID  int64               `tl:"query_id"`
	UserID   int64               `tl:"user_id"`
	Query    string              `tl:"query"`
	Geo      GeoPoint            `tl:"geo,omitempty:flags:0"`
	PeerType InlineQueryPeerType `tl:"peer_type,omitempty:flags:1"`
	Offset   string              `tl:"offset"`
}

func (*UpdateBotInlineQueryPredict) CRC() uint32 {
	return 0x496f379c
}
func (*UpdateBotInlineQueryPredict) _Update() {}

type UpdateBotInlineSendPredict struct {
	_      struct{}                `tl:"flags,bitflag"`
	UserID int64                   `tl:"user_id"`
	Query  string                  `tl:"query"`
	Geo    GeoPoint                `tl:"geo,omitempty:flags:0"`
	ID     string                  `tl:"id"`
	MsgID  InputBotInlineMessageID `tl:"msg_id,omitempty:flags:1"`
}

func (*UpdateBotInlineSendPredict) CRC() uint32 {
	return 0x12f12a07
}
func (*UpdateBotInlineSendPredict) _Update() {}

type UpdateEditChannelMessagePredict struct {
	Message  Message `tl:"message"`
	Pts      int32   `tl:"pts"`
	PtsCount int32   `tl:"pts_count"`
}

func (*UpdateEditChannelMessagePredict) CRC() uint32 {
	return 0x1b3f4df7
}
func (*UpdateEditChannelMessagePredict) _Update() {}

type UpdateBotCallbackQueryPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	QueryID       int64    `tl:"query_id"`
	UserID        int64    `tl:"user_id"`
	Peer          Peer     `tl:"peer"`
	MsgID         int32    `tl:"msg_id"`
	ChatInstance  int64    `tl:"chat_instance"`
	Data          *[]byte  `tl:"data,omitempty:flags:0"`
	GameShortName *string  `tl:"game_short_name,omitempty:flags:1"`
}

func (*UpdateBotCallbackQueryPredict) CRC() uint32 {
	return 0xb9cfc48d
}
func (*UpdateBotCallbackQueryPredict) _Update() {}

type UpdateEditMessagePredict struct {
	Message  Message `tl:"message"`
	Pts      int32   `tl:"pts"`
	PtsCount int32   `tl:"pts_count"`
}

func (*UpdateEditMessagePredict) CRC() uint32 {
	return 0xe40370a3
}
func (*UpdateEditMessagePredict) _Update() {}

type UpdateInlineBotCallbackQueryPredict struct {
	_             struct{}                `tl:"flags,bitflag"`
	QueryID       int64                   `tl:"query_id"`
	UserID        int64                   `tl:"user_id"`
	MsgID         InputBotInlineMessageID `tl:"msg_id"`
	ChatInstance  int64                   `tl:"chat_instance"`
	Data          *[]byte                 `tl:"data,omitempty:flags:0"`
	GameShortName *string                 `tl:"game_short_name,omitempty:flags:1"`
}

func (*UpdateInlineBotCallbackQueryPredict) CRC() uint32 {
	return 0x691e9052
}
func (*UpdateInlineBotCallbackQueryPredict) _Update() {}

type UpdateReadChannelOutboxPredict struct {
	ChannelID int64 `tl:"channel_id"`
	MaxID     int32 `tl:"max_id"`
}

func (*UpdateReadChannelOutboxPredict) CRC() uint32 {
	return 0xb75f99a9
}
func (*UpdateReadChannelOutboxPredict) _Update() {}

type UpdateDraftMessagePredict struct {
	_        struct{}     `tl:"flags,bitflag"`
	Peer     Peer         `tl:"peer"`
	TopMsgID *int32       `tl:"top_msg_id,omitempty:flags:0"`
	Draft    DraftMessage `tl:"draft"`
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
	ChannelID int64   `tl:"channel_id"`
	Webpage   WebPage `tl:"webpage"`
	Pts       int32   `tl:"pts"`
	PtsCount  int32   `tl:"pts_count"`
}

func (*UpdateChannelWebPagePredict) CRC() uint32 {
	return 0x2f2ba99f
}
func (*UpdateChannelWebPagePredict) _Update() {}

type UpdateDialogPinnedPredict struct {
	_        struct{}   `tl:"flags,bitflag"`
	Pinned   bool       `tl:"pinned,omitempty:flags:0,implicit"`
	FolderID *int32     `tl:"folder_id,omitempty:flags:1"`
	Peer     DialogPeer `tl:"peer"`
}

func (*UpdateDialogPinnedPredict) CRC() uint32 {
	return 0x6e6fe51c
}
func (*UpdateDialogPinnedPredict) _Update() {}

type UpdatePinnedDialogsPredict struct {
	_        struct{}     `tl:"flags,bitflag"`
	FolderID *int32       `tl:"folder_id,omitempty:flags:1"`
	Order    []DialogPeer `tl:"order,omitempty:flags:0"`
}

func (*UpdatePinnedDialogsPredict) CRC() uint32 {
	return 0xfa0f3ca2
}
func (*UpdatePinnedDialogsPredict) _Update() {}

type UpdateBotWebhookJSONPredict struct {
	Data DataJSON `tl:"data"`
}

func (*UpdateBotWebhookJSONPredict) CRC() uint32 {
	return 0x8317c0c3
}
func (*UpdateBotWebhookJSONPredict) _Update() {}

type UpdateBotWebhookJSONQueryPredict struct {
	QueryID int64    `tl:"query_id"`
	Data    DataJSON `tl:"data"`
	Timeout int32    `tl:"timeout"`
}

func (*UpdateBotWebhookJSONQueryPredict) CRC() uint32 {
	return 0x9b9240a6
}
func (*UpdateBotWebhookJSONQueryPredict) _Update() {}

type UpdateBotShippingQueryPredict struct {
	QueryID         int64       `tl:"query_id"`
	UserID          int64       `tl:"user_id"`
	Payload         []byte      `tl:"payload"`
	ShippingAddress PostAddress `tl:"shipping_address"`
}

func (*UpdateBotShippingQueryPredict) CRC() uint32 {
	return 0xb5aefd7d
}
func (*UpdateBotShippingQueryPredict) _Update() {}

type UpdateBotPrecheckoutQueryPredict struct {
	_                struct{}             `tl:"flags,bitflag"`
	QueryID          int64                `tl:"query_id"`
	UserID           int64                `tl:"user_id"`
	Payload          []byte               `tl:"payload"`
	Info             PaymentRequestedInfo `tl:"info,omitempty:flags:0"`
	ShippingOptionID *string              `tl:"shipping_option_id,omitempty:flags:1"`
	Currency         string               `tl:"currency"`
	TotalAmount      int64                `tl:"total_amount"`
}

func (*UpdateBotPrecheckoutQueryPredict) CRC() uint32 {
	return 0x8caa9a96
}
func (*UpdateBotPrecheckoutQueryPredict) _Update() {}

type UpdatePhoneCallPredict struct {
	PhoneCall PhoneCall `tl:"phone_call"`
}

func (*UpdatePhoneCallPredict) CRC() uint32 {
	return 0xab0f6b1e
}
func (*UpdatePhoneCallPredict) _Update() {}

type UpdateLangPackTooLongPredict struct {
	LangCode string `tl:"lang_code"`
}

func (*UpdateLangPackTooLongPredict) CRC() uint32 {
	return 0x46560264
}
func (*UpdateLangPackTooLongPredict) _Update() {}

type UpdateLangPackPredict struct {
	Difference LangPackDifference `tl:"difference"`
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
	ChannelID int64    `tl:"channel_id"`
	TopMsgID  *int32   `tl:"top_msg_id,omitempty:flags:0"`
	Messages  []int32  `tl:"messages"`
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
	ChannelID      int64 `tl:"channel_id"`
	AvailableMinID int32 `tl:"available_min_id"`
}

func (*UpdateChannelAvailableMessagesPredict) CRC() uint32 {
	return 0xb23fc698
}
func (*UpdateChannelAvailableMessagesPredict) _Update() {}

type UpdateDialogUnreadMarkPredict struct {
	_      struct{}   `tl:"flags,bitflag"`
	Unread bool       `tl:"unread,omitempty:flags:0,implicit"`
	Peer   DialogPeer `tl:"peer"`
}

func (*UpdateDialogUnreadMarkPredict) CRC() uint32 {
	return 0xe16459c3
}
func (*UpdateDialogUnreadMarkPredict) _Update() {}

type UpdateMessagePollPredict struct {
	_       struct{}    `tl:"flags,bitflag"`
	PollID  int64       `tl:"poll_id"`
	Poll    Poll        `tl:"poll,omitempty:flags:0"`
	Results PollResults `tl:"results"`
}

func (*UpdateMessagePollPredict) CRC() uint32 {
	return 0xaca1657b
}
func (*UpdateMessagePollPredict) _Update() {}

type UpdateChatDefaultBannedRightsPredict struct {
	Peer                Peer             `tl:"peer"`
	DefaultBannedRights ChatBannedRights `tl:"default_banned_rights"`
	Version             int32            `tl:"version"`
}

func (*UpdateChatDefaultBannedRightsPredict) CRC() uint32 {
	return 0x54c01850
}
func (*UpdateChatDefaultBannedRightsPredict) _Update() {}

type UpdateFolderPeersPredict struct {
	FolderPeers []FolderPeer `tl:"folder_peers"`
	Pts         int32        `tl:"pts"`
	PtsCount    int32        `tl:"pts_count"`
}

func (*UpdateFolderPeersPredict) CRC() uint32 {
	return 0x19360dc0
}
func (*UpdateFolderPeersPredict) _Update() {}

type UpdatePeerSettingsPredict struct {
	Peer     Peer         `tl:"peer"`
	Settings PeerSettings `tl:"settings"`
}

func (*UpdatePeerSettingsPredict) CRC() uint32 {
	return 0x6a7e7366
}
func (*UpdatePeerSettingsPredict) _Update() {}

type UpdatePeerLocatedPredict struct {
	Peers []PeerLocated `tl:"peers"`
}

func (*UpdatePeerLocatedPredict) CRC() uint32 {
	return 0xb4afcfb0
}
func (*UpdatePeerLocatedPredict) _Update() {}

type UpdateNewScheduledMessagePredict struct {
	Message Message `tl:"message"`
}

func (*UpdateNewScheduledMessagePredict) CRC() uint32 {
	return 0x39a51dfb
}
func (*UpdateNewScheduledMessagePredict) _Update() {}

type UpdateDeleteScheduledMessagesPredict struct {
	Peer     Peer    `tl:"peer"`
	Messages []int32 `tl:"messages"`
}

func (*UpdateDeleteScheduledMessagesPredict) CRC() uint32 {
	return 0x90866cee
}
func (*UpdateDeleteScheduledMessagesPredict) _Update() {}

type UpdateThemePredict struct {
	Theme Theme `tl:"theme"`
}

func (*UpdateThemePredict) CRC() uint32 {
	return 0x8216fba3
}
func (*UpdateThemePredict) _Update() {}

type UpdateGeoLiveViewedPredict struct {
	Peer  Peer  `tl:"peer"`
	MsgID int32 `tl:"msg_id"`
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
	PollID  int64    `tl:"poll_id"`
	Peer    Peer     `tl:"peer"`
	Options [][]byte `tl:"options"`
	Qts     int32    `tl:"qts"`
}

func (*UpdateMessagePollVotePredict) CRC() uint32 {
	return 0x24f40e77
}
func (*UpdateMessagePollVotePredict) _Update() {}

type UpdateDialogFilterPredict struct {
	_      struct{}     `tl:"flags,bitflag"`
	ID     int32        `tl:"id"`
	Filter DialogFilter `tl:"filter,omitempty:flags:0"`
}

func (*UpdateDialogFilterPredict) CRC() uint32 {
	return 0x26ffde7d
}
func (*UpdateDialogFilterPredict) _Update() {}

type UpdateDialogFilterOrderPredict struct {
	Order []int32 `tl:"order"`
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
	PhoneCallID int64  `tl:"phone_call_id"`
	Data        []byte `tl:"data"`
}

func (*UpdatePhoneCallSignalingDataPredict) CRC() uint32 {
	return 0x2661bf09
}
func (*UpdatePhoneCallSignalingDataPredict) _Update() {}

type UpdateChannelMessageForwardsPredict struct {
	ChannelID int64 `tl:"channel_id"`
	ID        int32 `tl:"id"`
	Forwards  int32 `tl:"forwards"`
}

func (*UpdateChannelMessageForwardsPredict) CRC() uint32 {
	return 0xd29a27f4
}
func (*UpdateChannelMessageForwardsPredict) _Update() {}

type UpdateReadChannelDiscussionInboxPredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ChannelID     int64    `tl:"channel_id"`
	TopMsgID      int32    `tl:"top_msg_id"`
	ReadMaxID     int32    `tl:"read_max_id"`
	BroadcastID   *int64   `tl:"broadcast_id,omitempty:flags:0"`
	BroadcastPost *int32   `tl:"broadcast_post,omitempty:flags:0"`
}

func (*UpdateReadChannelDiscussionInboxPredict) CRC() uint32 {
	return 0xd6b19546
}
func (*UpdateReadChannelDiscussionInboxPredict) _Update() {}

type UpdateReadChannelDiscussionOutboxPredict struct {
	ChannelID int64 `tl:"channel_id"`
	TopMsgID  int32 `tl:"top_msg_id"`
	ReadMaxID int32 `tl:"read_max_id"`
}

func (*UpdateReadChannelDiscussionOutboxPredict) CRC() uint32 {
	return 0x695c9e7c
}
func (*UpdateReadChannelDiscussionOutboxPredict) _Update() {}

type UpdatePeerBlockedPredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	Blocked              bool     `tl:"blocked,omitempty:flags:0,implicit"`
	BlockedMyStoriesFrom bool     `tl:"blocked_my_stories_from,omitempty:flags:1,implicit"`
	PeerID               Peer     `tl:"peer_id"`
}

func (*UpdatePeerBlockedPredict) CRC() uint32 {
	return 0xebe07752
}
func (*UpdatePeerBlockedPredict) _Update() {}

type UpdateChannelUserTypingPredict struct {
	_         struct{}          `tl:"flags,bitflag"`
	ChannelID int64             `tl:"channel_id"`
	TopMsgID  *int32            `tl:"top_msg_id,omitempty:flags:0"`
	FromID    Peer              `tl:"from_id"`
	Action    SendMessageAction `tl:"action"`
}

func (*UpdateChannelUserTypingPredict) CRC() uint32 {
	return 0x8c88c923
}
func (*UpdateChannelUserTypingPredict) _Update() {}

type UpdatePinnedMessagesPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Pinned   bool     `tl:"pinned,omitempty:flags:0,implicit"`
	Peer     Peer     `tl:"peer"`
	Messages []int32  `tl:"messages"`
	Pts      int32    `tl:"pts"`
	PtsCount int32    `tl:"pts_count"`
}

func (*UpdatePinnedMessagesPredict) CRC() uint32 {
	return 0xed85eab5
}
func (*UpdatePinnedMessagesPredict) _Update() {}

type UpdatePinnedChannelMessagesPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Pinned    bool     `tl:"pinned,omitempty:flags:0,implicit"`
	ChannelID int64    `tl:"channel_id"`
	Messages  []int32  `tl:"messages"`
	Pts       int32    `tl:"pts"`
	PtsCount  int32    `tl:"pts_count"`
}

func (*UpdatePinnedChannelMessagesPredict) CRC() uint32 {
	return 0x5bb98608
}
func (*UpdatePinnedChannelMessagesPredict) _Update() {}

type UpdateChatPredict struct {
	ChatID int64 `tl:"chat_id"`
}

func (*UpdateChatPredict) CRC() uint32 {
	return 0xf89a6a4e
}
func (*UpdateChatPredict) _Update() {}

type UpdateGroupCallParticipantsPredict struct {
	Call         InputGroupCall         `tl:"call"`
	Participants []GroupCallParticipant `tl:"participants"`
	Version      int32                  `tl:"version"`
}

func (*UpdateGroupCallParticipantsPredict) CRC() uint32 {
	return 0xf2ebdb4e
}
func (*UpdateGroupCallParticipantsPredict) _Update() {}

type UpdateGroupCallPredict struct {
	ChatID int64     `tl:"chat_id"`
	Call   GroupCall `tl:"call"`
}

func (*UpdateGroupCallPredict) CRC() uint32 {
	return 0x14b24500
}
func (*UpdateGroupCallPredict) _Update() {}

type UpdatePeerHistoryTTLPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Peer      Peer     `tl:"peer"`
	TTLPeriod *int32   `tl:"ttl_period,omitempty:flags:0"`
}

func (*UpdatePeerHistoryTTLPredict) CRC() uint32 {
	return 0xbb9bb9a5
}
func (*UpdatePeerHistoryTTLPredict) _Update() {}

type UpdateChatParticipantPredict struct {
	_               struct{}           `tl:"flags,bitflag"`
	ChatID          int64              `tl:"chat_id"`
	Date            int32              `tl:"date"`
	ActorID         int64              `tl:"actor_id"`
	UserID          int64              `tl:"user_id"`
	PrevParticipant ChatParticipant    `tl:"prev_participant,omitempty:flags:0"`
	NewParticipant  ChatParticipant    `tl:"new_participant,omitempty:flags:1"`
	Invite          ExportedChatInvite `tl:"invite,omitempty:flags:2"`
	Qts             int32              `tl:"qts"`
}

func (*UpdateChatParticipantPredict) CRC() uint32 {
	return 0xd087663a
}
func (*UpdateChatParticipantPredict) _Update() {}

type UpdateChannelParticipantPredict struct {
	_               struct{}           `tl:"flags,bitflag"`
	ViaChatlist     bool               `tl:"via_chatlist,omitempty:flags:3,implicit"`
	ChannelID       int64              `tl:"channel_id"`
	Date            int32              `tl:"date"`
	ActorID         int64              `tl:"actor_id"`
	UserID          int64              `tl:"user_id"`
	PrevParticipant ChannelParticipant `tl:"prev_participant,omitempty:flags:0"`
	NewParticipant  ChannelParticipant `tl:"new_participant,omitempty:flags:1"`
	Invite          ExportedChatInvite `tl:"invite,omitempty:flags:2"`
	Qts             int32              `tl:"qts"`
}

func (*UpdateChannelParticipantPredict) CRC() uint32 {
	return 0x985d3abb
}
func (*UpdateChannelParticipantPredict) _Update() {}

type UpdateBotStoppedPredict struct {
	UserID  int64 `tl:"user_id"`
	Date    int32 `tl:"date"`
	Stopped bool  `tl:"stopped"`
	Qts     int32 `tl:"qts"`
}

func (*UpdateBotStoppedPredict) CRC() uint32 {
	return 0xc4870a49
}
func (*UpdateBotStoppedPredict) _Update() {}

type UpdateGroupCallConnectionPredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Presentation bool     `tl:"presentation,omitempty:flags:0,implicit"`
	Params       DataJSON `tl:"params"`
}

func (*UpdateGroupCallConnectionPredict) CRC() uint32 {
	return 0x0b783982
}
func (*UpdateGroupCallConnectionPredict) _Update() {}

type UpdateBotCommandsPredict struct {
	Peer     Peer         `tl:"peer"`
	BotID    int64        `tl:"bot_id"`
	Commands []BotCommand `tl:"commands"`
}

func (*UpdateBotCommandsPredict) CRC() uint32 {
	return 0x4d712f2e
}
func (*UpdateBotCommandsPredict) _Update() {}

type UpdatePendingJoinRequestsPredict struct {
	Peer             Peer    `tl:"peer"`
	RequestsPending  int32   `tl:"requests_pending"`
	RecentRequesters []int64 `tl:"recent_requesters"`
}

func (*UpdatePendingJoinRequestsPredict) CRC() uint32 {
	return 0x7063c3db
}
func (*UpdatePendingJoinRequestsPredict) _Update() {}

type UpdateBotChatInviteRequesterPredict struct {
	Peer   Peer               `tl:"peer"`
	Date   int32              `tl:"date"`
	UserID int64              `tl:"user_id"`
	About  string             `tl:"about"`
	Invite ExportedChatInvite `tl:"invite"`
	Qts    int32              `tl:"qts"`
}

func (*UpdateBotChatInviteRequesterPredict) CRC() uint32 {
	return 0x11dfa986
}
func (*UpdateBotChatInviteRequesterPredict) _Update() {}

type UpdateMessageReactionsPredict struct {
	_         struct{}         `tl:"flags,bitflag"`
	Peer      Peer             `tl:"peer"`
	MsgID     int32            `tl:"msg_id"`
	TopMsgID  *int32           `tl:"top_msg_id,omitempty:flags:0"`
	Reactions MessageReactions `tl:"reactions"`
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
	QueryID int64 `tl:"query_id"`
}

func (*UpdateWebViewResultSentPredict) CRC() uint32 {
	return 0x1592b79d
}
func (*UpdateWebViewResultSentPredict) _Update() {}

type UpdateBotMenuButtonPredict struct {
	BotID  int64         `tl:"bot_id"`
	Button BotMenuButton `tl:"button"`
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
	Pending         bool     `tl:"pending,omitempty:flags:0,implicit"`
	Peer            Peer     `tl:"peer"`
	MsgID           int32    `tl:"msg_id"`
	TranscriptionID int64    `tl:"transcription_id"`
	Text            string   `tl:"text"`
}

func (*UpdateTranscribedAudioPredict) CRC() uint32 {
	return 0x0084cd5a
}
func (*UpdateTranscribedAudioPredict) _Update() {}

type UpdateReadFeaturedEmojiStickersPredict struct{}

func (*UpdateReadFeaturedEmojiStickersPredict) CRC() uint32 {
	return 0xfb4c496c
}
func (*UpdateReadFeaturedEmojiStickersPredict) _Update() {}

type UpdateUserEmojiStatusPredict struct {
	UserID      int64       `tl:"user_id"`
	EmojiStatus EmojiStatus `tl:"emoji_status"`
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
	Masks      bool     `tl:"masks,omitempty:flags:0,implicit"`
	Emojis     bool     `tl:"emojis,omitempty:flags:1,implicit"`
	Stickerset int64    `tl:"stickerset"`
}

func (*UpdateMoveStickerSetToTopPredict) CRC() uint32 {
	return 0x86fccf85
}
func (*UpdateMoveStickerSetToTopPredict) _Update() {}

type UpdateMessageExtendedMediaPredict struct {
	Peer          Peer                   `tl:"peer"`
	MsgID         int32                  `tl:"msg_id"`
	ExtendedMedia []MessageExtendedMedia `tl:"extended_media"`
}

func (*UpdateMessageExtendedMediaPredict) CRC() uint32 {
	return 0xd5a41724
}
func (*UpdateMessageExtendedMediaPredict) _Update() {}

type UpdateChannelPinnedTopicPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Pinned    bool     `tl:"pinned,omitempty:flags:0,implicit"`
	ChannelID int64    `tl:"channel_id"`
	TopicID   int32    `tl:"topic_id"`
}

func (*UpdateChannelPinnedTopicPredict) CRC() uint32 {
	return 0x192efbe3
}
func (*UpdateChannelPinnedTopicPredict) _Update() {}

type UpdateChannelPinnedTopicsPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	ChannelID int64    `tl:"channel_id"`
	Order     []int32  `tl:"order,omitempty:flags:0"`
}

func (*UpdateChannelPinnedTopicsPredict) CRC() uint32 {
	return 0xfe198602
}
func (*UpdateChannelPinnedTopicsPredict) _Update() {}

type UpdateUserPredict struct {
	UserID int64 `tl:"user_id"`
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
	Peer  Peer      `tl:"peer"`
	Story StoryItem `tl:"story"`
}

func (*UpdateStoryPredict) CRC() uint32 {
	return 0x75b3b798
}
func (*UpdateStoryPredict) _Update() {}

type UpdateReadStoriesPredict struct {
	Peer  Peer  `tl:"peer"`
	MaxID int32 `tl:"max_id"`
}

func (*UpdateReadStoriesPredict) CRC() uint32 {
	return 0xf74e932b
}
func (*UpdateReadStoriesPredict) _Update() {}

type UpdateStoryIDPredict struct {
	ID       int32 `tl:"id"`
	RandomID int64 `tl:"random_id"`
}

func (*UpdateStoryIDPredict) CRC() uint32 {
	return 0x1bf335b9
}
func (*UpdateStoryIDPredict) _Update() {}

type UpdateStoriesStealthModePredict struct {
	StealthMode StoriesStealthMode `tl:"stealth_mode"`
}

func (*UpdateStoriesStealthModePredict) CRC() uint32 {
	return 0x2c084dc1
}
func (*UpdateStoriesStealthModePredict) _Update() {}

type UpdateSentStoryReactionPredict struct {
	Peer     Peer     `tl:"peer"`
	StoryID  int32    `tl:"story_id"`
	Reaction Reaction `tl:"reaction"`
}

func (*UpdateSentStoryReactionPredict) CRC() uint32 {
	return 0x7d627683
}
func (*UpdateSentStoryReactionPredict) _Update() {}

type UpdateBotChatBoostPredict struct {
	Peer  Peer  `tl:"peer"`
	Boost Boost `tl:"boost"`
	Qts   int32 `tl:"qts"`
}

func (*UpdateBotChatBoostPredict) CRC() uint32 {
	return 0x904dd49c
}
func (*UpdateBotChatBoostPredict) _Update() {}

type UpdateChannelViewForumAsMessagesPredict struct {
	ChannelID int64 `tl:"channel_id"`
	Enabled   bool  `tl:"enabled"`
}

func (*UpdateChannelViewForumAsMessagesPredict) CRC() uint32 {
	return 0x07b68920
}
func (*UpdateChannelViewForumAsMessagesPredict) _Update() {}

type UpdatePeerWallpaperPredict struct {
	_                   struct{}  `tl:"flags,bitflag"`
	WallpaperOverridden bool      `tl:"wallpaper_overridden,omitempty:flags:1,implicit"`
	Peer                Peer      `tl:"peer"`
	Wallpaper           WallPaper `tl:"wallpaper,omitempty:flags:0"`
}

func (*UpdatePeerWallpaperPredict) CRC() uint32 {
	return 0xae3f101d
}
func (*UpdatePeerWallpaperPredict) _Update() {}

type UpdateBotMessageReactionPredict struct {
	Peer         Peer       `tl:"peer"`
	MsgID        int32      `tl:"msg_id"`
	Date         int32      `tl:"date"`
	Actor        Peer       `tl:"actor"`
	OldReactions []Reaction `tl:"old_reactions"`
	NewReactions []Reaction `tl:"new_reactions"`
	Qts          int32      `tl:"qts"`
}

func (*UpdateBotMessageReactionPredict) CRC() uint32 {
	return 0xac21d3ce
}
func (*UpdateBotMessageReactionPredict) _Update() {}

type UpdateBotMessageReactionsPredict struct {
	Peer      Peer            `tl:"peer"`
	MsgID     int32           `tl:"msg_id"`
	Date      int32           `tl:"date"`
	Reactions []ReactionCount `tl:"reactions"`
	Qts       int32           `tl:"qts"`
}

func (*UpdateBotMessageReactionsPredict) CRC() uint32 {
	return 0x09cb7759
}
func (*UpdateBotMessageReactionsPredict) _Update() {}

type UpdateSavedDialogPinnedPredict struct {
	_      struct{}   `tl:"flags,bitflag"`
	Pinned bool       `tl:"pinned,omitempty:flags:0,implicit"`
	Peer   DialogPeer `tl:"peer"`
}

func (*UpdateSavedDialogPinnedPredict) CRC() uint32 {
	return 0xaeaf9e74
}
func (*UpdateSavedDialogPinnedPredict) _Update() {}

type UpdatePinnedSavedDialogsPredict struct {
	_     struct{}     `tl:"flags,bitflag"`
	Order []DialogPeer `tl:"order,omitempty:flags:0"`
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
	JobID string `tl:"job_id"`
}

func (*UpdateSmsJobPredict) CRC() uint32 {
	return 0xf16269d4
}
func (*UpdateSmsJobPredict) _Update() {}

type UpdateQuickRepliesPredict struct {
	QuickReplies []QuickReply `tl:"quick_replies"`
}

func (*UpdateQuickRepliesPredict) CRC() uint32 {
	return 0xf9470ab2
}
func (*UpdateQuickRepliesPredict) _Update() {}

type UpdateNewQuickReplyPredict struct {
	QuickReply QuickReply `tl:"quick_reply"`
}

func (*UpdateNewQuickReplyPredict) CRC() uint32 {
	return 0xf53da717
}
func (*UpdateNewQuickReplyPredict) _Update() {}

type UpdateDeleteQuickReplyPredict struct {
	ShortcutID int32 `tl:"shortcut_id"`
}

func (*UpdateDeleteQuickReplyPredict) CRC() uint32 {
	return 0x53e6f1ec
}
func (*UpdateDeleteQuickReplyPredict) _Update() {}

type UpdateQuickReplyMessagePredict struct {
	Message Message `tl:"message"`
}

func (*UpdateQuickReplyMessagePredict) CRC() uint32 {
	return 0x3e050d0f
}
func (*UpdateQuickReplyMessagePredict) _Update() {}

type UpdateDeleteQuickReplyMessagesPredict struct {
	ShortcutID int32   `tl:"shortcut_id"`
	Messages   []int32 `tl:"messages"`
}

func (*UpdateDeleteQuickReplyMessagesPredict) CRC() uint32 {
	return 0x566fe7cd
}
func (*UpdateDeleteQuickReplyMessagesPredict) _Update() {}

type UpdateBotBusinessConnectPredict struct {
	Connection BotBusinessConnection `tl:"connection"`
	Qts        int32                 `tl:"qts"`
}

func (*UpdateBotBusinessConnectPredict) CRC() uint32 {
	return 0x8ae5c97a
}
func (*UpdateBotBusinessConnectPredict) _Update() {}

type UpdateBotNewBusinessMessagePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ConnectionID   string   `tl:"connection_id"`
	Message        Message  `tl:"message"`
	ReplyToMessage Message  `tl:"reply_to_message,omitempty:flags:0"`
	Qts            int32    `tl:"qts"`
}

func (*UpdateBotNewBusinessMessagePredict) CRC() uint32 {
	return 0x9ddb347c
}
func (*UpdateBotNewBusinessMessagePredict) _Update() {}

type UpdateBotEditBusinessMessagePredict struct {
	_              struct{} `tl:"flags,bitflag"`
	ConnectionID   string   `tl:"connection_id"`
	Message        Message  `tl:"message"`
	ReplyToMessage Message  `tl:"reply_to_message,omitempty:flags:0"`
	Qts            int32    `tl:"qts"`
}

func (*UpdateBotEditBusinessMessagePredict) CRC() uint32 {
	return 0x07df587c
}
func (*UpdateBotEditBusinessMessagePredict) _Update() {}

type UpdateBotDeleteBusinessMessagePredict struct {
	ConnectionID string  `tl:"connection_id"`
	Peer         Peer    `tl:"peer"`
	Messages     []int32 `tl:"messages"`
	Qts          int32   `tl:"qts"`
}

func (*UpdateBotDeleteBusinessMessagePredict) CRC() uint32 {
	return 0xa02a982e
}
func (*UpdateBotDeleteBusinessMessagePredict) _Update() {}

type UpdateNewStoryReactionPredict struct {
	StoryID  int32    `tl:"story_id"`
	Peer     Peer     `tl:"peer"`
	Reaction Reaction `tl:"reaction"`
}

func (*UpdateNewStoryReactionPredict) CRC() uint32 {
	return 0x1824e40b
}
func (*UpdateNewStoryReactionPredict) _Update() {}

type UpdateBroadcastRevenueTransactionsPredict struct {
	Peer     Peer                     `tl:"peer"`
	Balances BroadcastRevenueBalances `tl:"balances"`
}

func (*UpdateBroadcastRevenueTransactionsPredict) CRC() uint32 {
	return 0xdfd961f5
}
func (*UpdateBroadcastRevenueTransactionsPredict) _Update() {}

type UpdateStarsBalancePredict struct {
	Balance int64 `tl:"balance"`
}

func (*UpdateStarsBalancePredict) CRC() uint32 {
	return 0x0fb85198
}
func (*UpdateStarsBalancePredict) _Update() {}

type UpdateBusinessBotCallbackQueryPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	QueryID        int64    `tl:"query_id"`
	UserID         int64    `tl:"user_id"`
	ConnectionID   string   `tl:"connection_id"`
	Message        Message  `tl:"message"`
	ReplyToMessage Message  `tl:"reply_to_message,omitempty:flags:2"`
	ChatInstance   int64    `tl:"chat_instance"`
	Data           *[]byte  `tl:"data,omitempty:flags:0"`
}

func (*UpdateBusinessBotCallbackQueryPredict) CRC() uint32 {
	return 0x1ea2fda7
}
func (*UpdateBusinessBotCallbackQueryPredict) _Update() {}

type UpdateStarsRevenueStatusPredict struct {
	Peer   Peer               `tl:"peer"`
	Status StarsRevenueStatus `tl:"status"`
}

func (*UpdateStarsRevenueStatusPredict) CRC() uint32 {
	return 0xa584b019
}
func (*UpdateStarsRevenueStatusPredict) _Update() {}

type Updates interface {
	tl.TLObject
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
	_           struct{}           `tl:"flags,bitflag"`
	Out         bool               `tl:"out,omitempty:flags:1,implicit"`
	Mentioned   bool               `tl:"mentioned,omitempty:flags:4,implicit"`
	MediaUnread bool               `tl:"media_unread,omitempty:flags:5,implicit"`
	Silent      bool               `tl:"silent,omitempty:flags:13,implicit"`
	ID          int32              `tl:"id"`
	UserID      int64              `tl:"user_id"`
	Message     string             `tl:"message"`
	Pts         int32              `tl:"pts"`
	PtsCount    int32              `tl:"pts_count"`
	Date        int32              `tl:"date"`
	FwdFrom     MessageFwdHeader   `tl:"fwd_from,omitempty:flags:2"`
	ViaBotID    *int64             `tl:"via_bot_id,omitempty:flags:11"`
	ReplyTo     MessageReplyHeader `tl:"reply_to,omitempty:flags:3"`
	Entities    []MessageEntity    `tl:"entities,omitempty:flags:7"`
	TTLPeriod   *int32             `tl:"ttl_period,omitempty:flags:25"`
}

func (*UpdateShortMessagePredict) CRC() uint32 {
	return 0x313bc7f8
}
func (*UpdateShortMessagePredict) _Updates() {}

type UpdateShortChatMessagePredict struct {
	_           struct{}           `tl:"flags,bitflag"`
	Out         bool               `tl:"out,omitempty:flags:1,implicit"`
	Mentioned   bool               `tl:"mentioned,omitempty:flags:4,implicit"`
	MediaUnread bool               `tl:"media_unread,omitempty:flags:5,implicit"`
	Silent      bool               `tl:"silent,omitempty:flags:13,implicit"`
	ID          int32              `tl:"id"`
	FromID      int64              `tl:"from_id"`
	ChatID      int64              `tl:"chat_id"`
	Message     string             `tl:"message"`
	Pts         int32              `tl:"pts"`
	PtsCount    int32              `tl:"pts_count"`
	Date        int32              `tl:"date"`
	FwdFrom     MessageFwdHeader   `tl:"fwd_from,omitempty:flags:2"`
	ViaBotID    *int64             `tl:"via_bot_id,omitempty:flags:11"`
	ReplyTo     MessageReplyHeader `tl:"reply_to,omitempty:flags:3"`
	Entities    []MessageEntity    `tl:"entities,omitempty:flags:7"`
	TTLPeriod   *int32             `tl:"ttl_period,omitempty:flags:25"`
}

func (*UpdateShortChatMessagePredict) CRC() uint32 {
	return 0x4d6deea5
}
func (*UpdateShortChatMessagePredict) _Updates() {}

type UpdateShortPredict struct {
	Update Update `tl:"update"`
	Date   int32  `tl:"date"`
}

func (*UpdateShortPredict) CRC() uint32 {
	return 0x78d4dec1
}
func (*UpdateShortPredict) _Updates() {}

type UpdatesCombinedPredict struct {
	Updates  []Update `tl:"updates"`
	Users    []User   `tl:"users"`
	Chats    []Chat   `tl:"chats"`
	Date     int32    `tl:"date"`
	SeqStart int32    `tl:"seq_start"`
	Seq      int32    `tl:"seq"`
}

func (*UpdatesCombinedPredict) CRC() uint32 {
	return 0x725b04c3
}
func (*UpdatesCombinedPredict) _Updates() {}

type UpdatesPredict struct {
	Updates []Update `tl:"updates"`
	Users   []User   `tl:"users"`
	Chats   []Chat   `tl:"chats"`
	Date    int32    `tl:"date"`
	Seq     int32    `tl:"seq"`
}

func (*UpdatesPredict) CRC() uint32 {
	return 0x74ae4240
}
func (*UpdatesPredict) _Updates() {}

type UpdateShortSentMessagePredict struct {
	_         struct{}        `tl:"flags,bitflag"`
	Out       bool            `tl:"out,omitempty:flags:1,implicit"`
	ID        int32           `tl:"id"`
	Pts       int32           `tl:"pts"`
	PtsCount  int32           `tl:"pts_count"`
	Date      int32           `tl:"date"`
	Media     MessageMedia    `tl:"media,omitempty:flags:9"`
	Entities  []MessageEntity `tl:"entities,omitempty:flags:7"`
	TTLPeriod *int32          `tl:"ttl_period,omitempty:flags:25"`
}

func (*UpdateShortSentMessagePredict) CRC() uint32 {
	return 0x9015e101
}
func (*UpdateShortSentMessagePredict) _Updates() {}

type URLAuthResult interface {
	tl.TLObject
	_URLAuthResult()
}

var (
	_ URLAuthResult = (*URLAuthResultRequestPredict)(nil)
	_ URLAuthResult = (*URLAuthResultAcceptedPredict)(nil)
	_ URLAuthResult = (*URLAuthResultDefaultPredict)(nil)
)

type URLAuthResultRequestPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	RequestWriteAccess bool     `tl:"request_write_access,omitempty:flags:0,implicit"`
	Bot                User     `tl:"bot"`
	Domain             string   `tl:"domain"`
}

func (*URLAuthResultRequestPredict) CRC() uint32 {
	return 0x92d33a0e
}
func (*URLAuthResultRequestPredict) _URLAuthResult() {}

type URLAuthResultAcceptedPredict struct {
	URL string `tl:"url"`
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
	tl.TLObject
	_User()
}

var (
	_ User = (*UserEmptyPredict)(nil)
	_ User = (*UserPredict)(nil)
)

type UserEmptyPredict struct {
	ID int64 `tl:"id"`
}

func (*UserEmptyPredict) CRC() uint32 {
	return 0xd3bc4b7a
}
func (*UserEmptyPredict) _User() {}

type UserPredict struct {
	_                     struct{}            `tl:"flags,bitflag"`
	Self                  bool                `tl:"self,omitempty:flags:10,implicit"`
	Contact               bool                `tl:"contact,omitempty:flags:11,implicit"`
	MutualContact         bool                `tl:"mutual_contact,omitempty:flags:12,implicit"`
	Deleted               bool                `tl:"deleted,omitempty:flags:13,implicit"`
	Bot                   bool                `tl:"bot,omitempty:flags:14,implicit"`
	BotChatHistory        bool                `tl:"bot_chat_history,omitempty:flags:15,implicit"`
	BotNochats            bool                `tl:"bot_nochats,omitempty:flags:16,implicit"`
	Verified              bool                `tl:"verified,omitempty:flags:17,implicit"`
	Restricted            bool                `tl:"restricted,omitempty:flags:18,implicit"`
	Min                   bool                `tl:"min,omitempty:flags:20,implicit"`
	BotInlineGeo          bool                `tl:"bot_inline_geo,omitempty:flags:21,implicit"`
	Support               bool                `tl:"support,omitempty:flags:23,implicit"`
	Scam                  bool                `tl:"scam,omitempty:flags:24,implicit"`
	ApplyMinPhoto         bool                `tl:"apply_min_photo,omitempty:flags:25,implicit"`
	Fake                  bool                `tl:"fake,omitempty:flags:26,implicit"`
	BotAttachMenu         bool                `tl:"bot_attach_menu,omitempty:flags:27,implicit"`
	Premium               bool                `tl:"premium,omitempty:flags:28,implicit"`
	AttachMenuEnabled     bool                `tl:"attach_menu_enabled,omitempty:flags:29,implicit"`
	_                     struct{}            `tl:"flags2,bitflag"`
	BotCanEdit            bool                `tl:"bot_can_edit,omitempty:flags2:1,implicit"`
	CloseFriend           bool                `tl:"close_friend,omitempty:flags2:2,implicit"`
	StoriesHidden         bool                `tl:"stories_hidden,omitempty:flags2:3,implicit"`
	StoriesUnavailable    bool                `tl:"stories_unavailable,omitempty:flags2:4,implicit"`
	ContactRequirePremium bool                `tl:"contact_require_premium,omitempty:flags2:10,implicit"`
	BotBusiness           bool                `tl:"bot_business,omitempty:flags2:11,implicit"`
	BotHasMainApp         bool                `tl:"bot_has_main_app,omitempty:flags2:13,implicit"`
	ID                    int64               `tl:"id"`
	AccessHash            *int64              `tl:"access_hash,omitempty:flags:0"`
	FirstName             *string             `tl:"first_name,omitempty:flags:1"`
	LastName              *string             `tl:"last_name,omitempty:flags:2"`
	Username              *string             `tl:"username,omitempty:flags:3"`
	Phone                 *string             `tl:"phone,omitempty:flags:4"`
	Photo                 UserProfilePhoto    `tl:"photo,omitempty:flags:5"`
	Status                UserStatus          `tl:"status,omitempty:flags:6"`
	BotInfoVersion        *int32              `tl:"bot_info_version,omitempty:flags:14"`
	RestrictionReason     []RestrictionReason `tl:"restriction_reason,omitempty:flags:18"`
	BotInlinePlaceholder  *string             `tl:"bot_inline_placeholder,omitempty:flags:19"`
	LangCode              *string             `tl:"lang_code,omitempty:flags:22"`
	EmojiStatus           EmojiStatus         `tl:"emoji_status,omitempty:flags:30"`
	Usernames             []Username          `tl:"usernames,omitempty:flags2:0"`
	StoriesMaxID          *int32              `tl:"stories_max_id,omitempty:flags2:5"`
	Color                 PeerColor           `tl:"color,omitempty:flags2:8"`
	ProfileColor          PeerColor           `tl:"profile_color,omitempty:flags2:9"`
	BotActiveUsers        *int32              `tl:"bot_active_users,omitempty:flags2:12"`
}

func (*UserPredict) CRC() uint32 {
	return 0x83314fca
}
func (*UserPredict) _User() {}

type UserFull interface {
	tl.TLObject
	_UserFull()
}

var (
	_ UserFull = (*UserFullPredict)(nil)
)

type UserFullPredict struct {
	_                       struct{}                `tl:"flags,bitflag"`
	Blocked                 bool                    `tl:"blocked,omitempty:flags:0,implicit"`
	PhoneCallsAvailable     bool                    `tl:"phone_calls_available,omitempty:flags:4,implicit"`
	PhoneCallsPrivate       bool                    `tl:"phone_calls_private,omitempty:flags:5,implicit"`
	CanPinMessage           bool                    `tl:"can_pin_message,omitempty:flags:7,implicit"`
	HasScheduled            bool                    `tl:"has_scheduled,omitempty:flags:12,implicit"`
	VideoCallsAvailable     bool                    `tl:"video_calls_available,omitempty:flags:13,implicit"`
	VoiceMessagesForbidden  bool                    `tl:"voice_messages_forbidden,omitempty:flags:20,implicit"`
	TranslationsDisabled    bool                    `tl:"translations_disabled,omitempty:flags:23,implicit"`
	StoriesPinnedAvailable  bool                    `tl:"stories_pinned_available,omitempty:flags:26,implicit"`
	BlockedMyStoriesFrom    bool                    `tl:"blocked_my_stories_from,omitempty:flags:27,implicit"`
	WallpaperOverridden     bool                    `tl:"wallpaper_overridden,omitempty:flags:28,implicit"`
	ContactRequirePremium   bool                    `tl:"contact_require_premium,omitempty:flags:29,implicit"`
	ReadDatesPrivate        bool                    `tl:"read_dates_private,omitempty:flags:30,implicit"`
	_                       struct{}                `tl:"flags2,bitflag"`
	SponsoredEnabled        bool                    `tl:"sponsored_enabled,omitempty:flags2:7,implicit"`
	ID                      int64                   `tl:"id"`
	About                   *string                 `tl:"about,omitempty:flags:1"`
	Settings                PeerSettings            `tl:"settings"`
	PersonalPhoto           Photo                   `tl:"personal_photo,omitempty:flags:21"`
	ProfilePhoto            Photo                   `tl:"profile_photo,omitempty:flags:2"`
	FallbackPhoto           Photo                   `tl:"fallback_photo,omitempty:flags:22"`
	NotifySettings          PeerNotifySettings      `tl:"notify_settings"`
	BotInfo                 BotInfo                 `tl:"bot_info,omitempty:flags:3"`
	PinnedMsgID             *int32                  `tl:"pinned_msg_id,omitempty:flags:6"`
	CommonChatsCount        int32                   `tl:"common_chats_count"`
	FolderID                *int32                  `tl:"folder_id,omitempty:flags:11"`
	TTLPeriod               *int32                  `tl:"ttl_period,omitempty:flags:14"`
	ThemeEmoticon           *string                 `tl:"theme_emoticon,omitempty:flags:15"`
	PrivateForwardName      *string                 `tl:"private_forward_name,omitempty:flags:16"`
	BotGroupAdminRights     ChatAdminRights         `tl:"bot_group_admin_rights,omitempty:flags:17"`
	BotBroadcastAdminRights ChatAdminRights         `tl:"bot_broadcast_admin_rights,omitempty:flags:18"`
	PremiumGifts            []PremiumGiftOption     `tl:"premium_gifts,omitempty:flags:19"`
	Wallpaper               WallPaper               `tl:"wallpaper,omitempty:flags:24"`
	Stories                 PeerStories             `tl:"stories,omitempty:flags:25"`
	BusinessWorkHours       BusinessWorkHours       `tl:"business_work_hours,omitempty:flags2:0"`
	BusinessLocation        BusinessLocation        `tl:"business_location,omitempty:flags2:1"`
	BusinessGreetingMessage BusinessGreetingMessage `tl:"business_greeting_message,omitempty:flags2:2"`
	BusinessAwayMessage     BusinessAwayMessage     `tl:"business_away_message,omitempty:flags2:3"`
	BusinessIntro           BusinessIntro           `tl:"business_intro,omitempty:flags2:4"`
	Birthday                Birthday                `tl:"birthday,omitempty:flags2:5"`
	PersonalChannelID       *int64                  `tl:"personal_channel_id,omitempty:flags2:6"`
	PersonalChannelMessage  *int32                  `tl:"personal_channel_message,omitempty:flags2:6"`
}

func (*UserFullPredict) CRC() uint32 {
	return 0xcc997720
}
func (*UserFullPredict) _UserFull() {}

type UserProfilePhoto interface {
	tl.TLObject
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
	HasVideo      bool     `tl:"has_video,omitempty:flags:0,implicit"`
	Personal      bool     `tl:"personal,omitempty:flags:2,implicit"`
	PhotoID       int64    `tl:"photo_id"`
	StrippedThumb *[]byte  `tl:"stripped_thumb,omitempty:flags:1"`
	DcID          int32    `tl:"dc_id"`
}

func (*UserProfilePhotoPredict) CRC() uint32 {
	return 0x82d1f706
}
func (*UserProfilePhotoPredict) _UserProfilePhoto() {}

type UserStatus interface {
	tl.TLObject
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
	return 0x09d05049
}
func (*UserStatusEmptyPredict) _UserStatus() {}

type UserStatusOnlinePredict struct {
	Expires int32 `tl:"expires"`
}

func (*UserStatusOnlinePredict) CRC() uint32 {
	return 0xedb93949
}
func (*UserStatusOnlinePredict) _UserStatus() {}

type UserStatusOfflinePredict struct {
	WasOnline int32 `tl:"was_online"`
}

func (*UserStatusOfflinePredict) CRC() uint32 {
	return 0x008c703f
}
func (*UserStatusOfflinePredict) _UserStatus() {}

type UserStatusRecentlyPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ByMe bool     `tl:"by_me,omitempty:flags:0,implicit"`
}

func (*UserStatusRecentlyPredict) CRC() uint32 {
	return 0x7b197dc8
}
func (*UserStatusRecentlyPredict) _UserStatus() {}

type UserStatusLastWeekPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ByMe bool     `tl:"by_me,omitempty:flags:0,implicit"`
}

func (*UserStatusLastWeekPredict) CRC() uint32 {
	return 0x541a1d1a
}
func (*UserStatusLastWeekPredict) _UserStatus() {}

type UserStatusLastMonthPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ByMe bool     `tl:"by_me,omitempty:flags:0,implicit"`
}

func (*UserStatusLastMonthPredict) CRC() uint32 {
	return 0x65899777
}
func (*UserStatusLastMonthPredict) _UserStatus() {}

type Username interface {
	tl.TLObject
	_Username()
}

var (
	_ Username = (*UsernamePredict)(nil)
)

type UsernamePredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Editable bool     `tl:"editable,omitempty:flags:0,implicit"`
	Active   bool     `tl:"active,omitempty:flags:1,implicit"`
	Username string   `tl:"username"`
}

func (*UsernamePredict) CRC() uint32 {
	return 0xb4073647
}
func (*UsernamePredict) _Username() {}

type VideoSize interface {
	tl.TLObject
	_VideoSize()
}

var (
	_ VideoSize = (*VideoSizePredict)(nil)
	_ VideoSize = (*VideoSizeEmojiMarkupPredict)(nil)
	_ VideoSize = (*VideoSizeStickerMarkupPredict)(nil)
)

type VideoSizePredict struct {
	_            struct{} `tl:"flags,bitflag"`
	Type         string   `tl:"type"`
	W            int32    `tl:"w"`
	H            int32    `tl:"h"`
	Size         int32    `tl:"size"`
	VideoStartTs *float64 `tl:"video_start_ts,omitempty:flags:0"`
}

func (*VideoSizePredict) CRC() uint32 {
	return 0xde33b094
}
func (*VideoSizePredict) _VideoSize() {}

type VideoSizeEmojiMarkupPredict struct {
	EmojiID          int64   `tl:"emoji_id"`
	BackgroundColors []int32 `tl:"background_colors"`
}

func (*VideoSizeEmojiMarkupPredict) CRC() uint32 {
	return 0xf85c413c
}
func (*VideoSizeEmojiMarkupPredict) _VideoSize() {}

type VideoSizeStickerMarkupPredict struct {
	Stickerset       InputStickerSet `tl:"stickerset"`
	StickerID        int64           `tl:"sticker_id"`
	BackgroundColors []int32         `tl:"background_colors"`
}

func (*VideoSizeStickerMarkupPredict) CRC() uint32 {
	return 0x0da082fe
}
func (*VideoSizeStickerMarkupPredict) _VideoSize() {}

type WallPaper interface {
	tl.TLObject
	_WallPaper()
}

var (
	_ WallPaper = (*WallPaperPredict)(nil)
	_ WallPaper = (*WallPaperNoFilePredict)(nil)
)

type WallPaperPredict struct {
	ID         int64             `tl:"id"`
	_          struct{}          `tl:"flags,bitflag"`
	Creator    bool              `tl:"creator,omitempty:flags:0,implicit"`
	Default    bool              `tl:"default,omitempty:flags:1,implicit"`
	Pattern    bool              `tl:"pattern,omitempty:flags:3,implicit"`
	Dark       bool              `tl:"dark,omitempty:flags:4,implicit"`
	AccessHash int64             `tl:"access_hash"`
	Slug       string            `tl:"slug"`
	Document   Document          `tl:"document"`
	Settings   WallPaperSettings `tl:"settings,omitempty:flags:2"`
}

func (*WallPaperPredict) CRC() uint32 {
	return 0xa437c3ed
}
func (*WallPaperPredict) _WallPaper() {}

type WallPaperNoFilePredict struct {
	ID       int64             `tl:"id"`
	_        struct{}          `tl:"flags,bitflag"`
	Default  bool              `tl:"default,omitempty:flags:1,implicit"`
	Dark     bool              `tl:"dark,omitempty:flags:4,implicit"`
	Settings WallPaperSettings `tl:"settings,omitempty:flags:2"`
}

func (*WallPaperNoFilePredict) CRC() uint32 {
	return 0xe0804116
}
func (*WallPaperNoFilePredict) _WallPaper() {}

type WallPaperSettings interface {
	tl.TLObject
	_WallPaperSettings()
}

var (
	_ WallPaperSettings = (*WallPaperSettingsPredict)(nil)
)

type WallPaperSettingsPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	Blur                  bool     `tl:"blur,omitempty:flags:1,implicit"`
	Motion                bool     `tl:"motion,omitempty:flags:2,implicit"`
	BackgroundColor       *int32   `tl:"background_color,omitempty:flags:0"`
	SecondBackgroundColor *int32   `tl:"second_background_color,omitempty:flags:4"`
	ThirdBackgroundColor  *int32   `tl:"third_background_color,omitempty:flags:5"`
	FourthBackgroundColor *int32   `tl:"fourth_background_color,omitempty:flags:6"`
	Intensity             *int32   `tl:"intensity,omitempty:flags:3"`
	Rotation              *int32   `tl:"rotation,omitempty:flags:4"`
	Emoticon              *string  `tl:"emoticon,omitempty:flags:7"`
}

func (*WallPaperSettingsPredict) CRC() uint32 {
	return 0x372efcd0
}
func (*WallPaperSettingsPredict) _WallPaperSettings() {}

type WebAuthorization interface {
	tl.TLObject
	_WebAuthorization()
}

var (
	_ WebAuthorization = (*WebAuthorizationPredict)(nil)
)

type WebAuthorizationPredict struct {
	Hash        int64  `tl:"hash"`
	BotID       int64  `tl:"bot_id"`
	Domain      string `tl:"domain"`
	Browser     string `tl:"browser"`
	Platform    string `tl:"platform"`
	DateCreated int32  `tl:"date_created"`
	DateActive  int32  `tl:"date_active"`
	Ip          string `tl:"ip"`
	Region      string `tl:"region"`
}

func (*WebAuthorizationPredict) CRC() uint32 {
	return 0xa6f8f452
}
func (*WebAuthorizationPredict) _WebAuthorization() {}

type WebDocument interface {
	tl.TLObject
	_WebDocument()
}

var (
	_ WebDocument = (*WebDocumentPredict)(nil)
	_ WebDocument = (*WebDocumentNoProxyPredict)(nil)
)

type WebDocumentPredict struct {
	URL        string              `tl:"url"`
	AccessHash int64               `tl:"access_hash"`
	Size       int32               `tl:"size"`
	MimeType   string              `tl:"mime_type"`
	Attributes []DocumentAttribute `tl:"attributes"`
}

func (*WebDocumentPredict) CRC() uint32 {
	return 0x1c570ed1
}
func (*WebDocumentPredict) _WebDocument() {}

type WebDocumentNoProxyPredict struct {
	URL        string              `tl:"url"`
	Size       int32               `tl:"size"`
	MimeType   string              `tl:"mime_type"`
	Attributes []DocumentAttribute `tl:"attributes"`
}

func (*WebDocumentNoProxyPredict) CRC() uint32 {
	return 0xf9c8bcc6
}
func (*WebDocumentNoProxyPredict) _WebDocument() {}

type WebPage interface {
	tl.TLObject
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
	ID  int64    `tl:"id"`
	URL *string  `tl:"url,omitempty:flags:0"`
}

func (*WebPageEmptyPredict) CRC() uint32 {
	return 0x211a1788
}
func (*WebPageEmptyPredict) _WebPage() {}

type WebPagePendingPredict struct {
	_    struct{} `tl:"flags,bitflag"`
	ID   int64    `tl:"id"`
	URL  *string  `tl:"url,omitempty:flags:0"`
	Date int32    `tl:"date"`
}

func (*WebPagePendingPredict) CRC() uint32 {
	return 0xb0d13e47
}
func (*WebPagePendingPredict) _WebPage() {}

type WebPagePredict struct {
	_             struct{}           `tl:"flags,bitflag"`
	HasLargeMedia bool               `tl:"has_large_media,omitempty:flags:13,implicit"`
	ID            int64              `tl:"id"`
	URL           string             `tl:"url"`
	DisplayURL    string             `tl:"display_url"`
	Hash          int32              `tl:"hash"`
	Type          *string            `tl:"type,omitempty:flags:0"`
	SiteName      *string            `tl:"site_name,omitempty:flags:1"`
	Title         *string            `tl:"title,omitempty:flags:2"`
	Description   *string            `tl:"description,omitempty:flags:3"`
	Photo         Photo              `tl:"photo,omitempty:flags:4"`
	EmbedURL      *string            `tl:"embed_url,omitempty:flags:5"`
	EmbedType     *string            `tl:"embed_type,omitempty:flags:5"`
	EmbedWidth    *int32             `tl:"embed_width,omitempty:flags:6"`
	EmbedHeight   *int32             `tl:"embed_height,omitempty:flags:6"`
	Duration      *int32             `tl:"duration,omitempty:flags:7"`
	Author        *string            `tl:"author,omitempty:flags:8"`
	Document      Document           `tl:"document,omitempty:flags:9"`
	CachedPage    Page               `tl:"cached_page,omitempty:flags:10"`
	Attributes    []WebPageAttribute `tl:"attributes,omitempty:flags:12"`
}

func (*WebPagePredict) CRC() uint32 {
	return 0xe89c45b2
}
func (*WebPagePredict) _WebPage() {}

type WebPageNotModifiedPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	CachedPageViews *int32   `tl:"cached_page_views,omitempty:flags:0"`
}

func (*WebPageNotModifiedPredict) CRC() uint32 {
	return 0x7311ca11
}
func (*WebPageNotModifiedPredict) _WebPage() {}

type WebPageAttribute interface {
	tl.TLObject
	_WebPageAttribute()
}

var (
	_ WebPageAttribute = (*WebPageAttributeThemePredict)(nil)
	_ WebPageAttribute = (*WebPageAttributeStoryPredict)(nil)
	_ WebPageAttribute = (*WebPageAttributeStickerSetPredict)(nil)
)

type WebPageAttributeThemePredict struct {
	_         struct{}      `tl:"flags,bitflag"`
	Documents []Document    `tl:"documents,omitempty:flags:0"`
	Settings  ThemeSettings `tl:"settings,omitempty:flags:1"`
}

func (*WebPageAttributeThemePredict) CRC() uint32 {
	return 0x54b56617
}
func (*WebPageAttributeThemePredict) _WebPageAttribute() {}

type WebPageAttributeStoryPredict struct {
	_     struct{}  `tl:"flags,bitflag"`
	Peer  Peer      `tl:"peer"`
	ID    int32     `tl:"id"`
	Story StoryItem `tl:"story,omitempty:flags:0"`
}

func (*WebPageAttributeStoryPredict) CRC() uint32 {
	return 0x2e94c3e7
}
func (*WebPageAttributeStoryPredict) _WebPageAttribute() {}

type WebPageAttributeStickerSetPredict struct {
	_         struct{}   `tl:"flags,bitflag"`
	Emojis    bool       `tl:"emojis,omitempty:flags:0,implicit"`
	TextColor bool       `tl:"text_color,omitempty:flags:1,implicit"`
	Stickers  []Document `tl:"stickers"`
}

func (*WebPageAttributeStickerSetPredict) CRC() uint32 {
	return 0x50cc03d3
}
func (*WebPageAttributeStickerSetPredict) _WebPageAttribute() {}

type WebViewMessageSent interface {
	tl.TLObject
	_WebViewMessageSent()
}

var (
	_ WebViewMessageSent = (*WebViewMessageSentPredict)(nil)
)

type WebViewMessageSentPredict struct {
	_     struct{}                `tl:"flags,bitflag"`
	MsgID InputBotInlineMessageID `tl:"msg_id,omitempty:flags:0"`
}

func (*WebViewMessageSentPredict) CRC() uint32 {
	return 0x0c94511c
}
func (*WebViewMessageSentPredict) _WebViewMessageSent() {}

type WebViewResult interface {
	tl.TLObject
	_WebViewResult()
}

var (
	_ WebViewResult = (*WebViewResultURLPredict)(nil)
)

type WebViewResultURLPredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Fullsize bool     `tl:"fullsize,omitempty:flags:1,implicit"`
	QueryID  *int64   `tl:"query_id,omitempty:flags:0"`
	URL      string   `tl:"url"`
}

func (*WebViewResultURLPredict) CRC() uint32 {
	return 0x4d22ff98
}
func (*WebViewResultURLPredict) _WebViewResult() {}

type AccountAuthorizationForm interface {
	tl.TLObject
	_AccountAuthorizationForm()
}

var (
	_ AccountAuthorizationForm = (*AccountAuthorizationFormPredict)(nil)
)

type AccountAuthorizationFormPredict struct {
	_                struct{}             `tl:"flags,bitflag"`
	RequiredTypes    []SecureRequiredType `tl:"required_types"`
	Values           []SecureValue        `tl:"values"`
	Errors           []SecureValueError   `tl:"errors"`
	Users            []User               `tl:"users"`
	PrivacyPolicyURL *string              `tl:"privacy_policy_url,omitempty:flags:0"`
}

func (*AccountAuthorizationFormPredict) CRC() uint32 {
	return 0xad2e1cd8
}
func (*AccountAuthorizationFormPredict) _AccountAuthorizationForm() {}

type AccountAuthorizations interface {
	tl.TLObject
	_AccountAuthorizations()
}

var (
	_ AccountAuthorizations = (*AccountAuthorizationsPredict)(nil)
)

type AccountAuthorizationsPredict struct {
	AuthorizationTTLDays int32           `tl:"authorization_ttl_days"`
	Authorizations       []Authorization `tl:"authorizations"`
}

func (*AccountAuthorizationsPredict) CRC() uint32 {
	return 0x4bff8ea0
}
func (*AccountAuthorizationsPredict) _AccountAuthorizations() {}

type AccountAutoDownloadSettings interface {
	tl.TLObject
	_AccountAutoDownloadSettings()
}

var (
	_ AccountAutoDownloadSettings = (*AccountAutoDownloadSettingsPredict)(nil)
)

type AccountAutoDownloadSettingsPredict struct {
	Low    AutoDownloadSettings `tl:"low"`
	Medium AutoDownloadSettings `tl:"medium"`
	High   AutoDownloadSettings `tl:"high"`
}

func (*AccountAutoDownloadSettingsPredict) CRC() uint32 {
	return 0x63cacf26
}
func (*AccountAutoDownloadSettingsPredict) _AccountAutoDownloadSettings() {}

type AccountAutoSaveSettings interface {
	tl.TLObject
	_AccountAutoSaveSettings()
}

var (
	_ AccountAutoSaveSettings = (*AccountAutoSaveSettingsPredict)(nil)
)

type AccountAutoSaveSettingsPredict struct {
	UsersSettings      AutoSaveSettings    `tl:"users_settings"`
	ChatsSettings      AutoSaveSettings    `tl:"chats_settings"`
	BroadcastsSettings AutoSaveSettings    `tl:"broadcasts_settings"`
	Exceptions         []AutoSaveException `tl:"exceptions"`
	Chats              []Chat              `tl:"chats"`
	Users              []User              `tl:"users"`
}

func (*AccountAutoSaveSettingsPredict) CRC() uint32 {
	return 0x4c3e069d
}
func (*AccountAutoSaveSettingsPredict) _AccountAutoSaveSettings() {}

type AccountBusinessChatLinks interface {
	tl.TLObject
	_AccountBusinessChatLinks()
}

var (
	_ AccountBusinessChatLinks = (*AccountBusinessChatLinksPredict)(nil)
)

type AccountBusinessChatLinksPredict struct {
	Links []BusinessChatLink `tl:"links"`
	Chats []Chat             `tl:"chats"`
	Users []User             `tl:"users"`
}

func (*AccountBusinessChatLinksPredict) CRC() uint32 {
	return 0xec43a2d1
}
func (*AccountBusinessChatLinksPredict) _AccountBusinessChatLinks() {}

type AccountConnectedBots interface {
	tl.TLObject
	_AccountConnectedBots()
}

var (
	_ AccountConnectedBots = (*AccountConnectedBotsPredict)(nil)
)

type AccountConnectedBotsPredict struct {
	ConnectedBots []ConnectedBot `tl:"connected_bots"`
	Users         []User         `tl:"users"`
}

func (*AccountConnectedBotsPredict) CRC() uint32 {
	return 0x17d7f87b
}
func (*AccountConnectedBotsPredict) _AccountConnectedBots() {}

type AccountContentSettings interface {
	tl.TLObject
	_AccountContentSettings()
}

var (
	_ AccountContentSettings = (*AccountContentSettingsPredict)(nil)
)

type AccountContentSettingsPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	SensitiveEnabled   bool     `tl:"sensitive_enabled,omitempty:flags:0,implicit"`
	SensitiveCanChange bool     `tl:"sensitive_can_change,omitempty:flags:1,implicit"`
}

func (*AccountContentSettingsPredict) CRC() uint32 {
	return 0x57e28221
}
func (*AccountContentSettingsPredict) _AccountContentSettings() {}

type AccountEmailVerified interface {
	tl.TLObject
	_AccountEmailVerified()
}

var (
	_ AccountEmailVerified = (*AccountEmailVerifiedPredict)(nil)
	_ AccountEmailVerified = (*AccountEmailVerifiedLoginPredict)(nil)
)

type AccountEmailVerifiedPredict struct {
	Email string `tl:"email"`
}

func (*AccountEmailVerifiedPredict) CRC() uint32 {
	return 0x2b96cd1b
}
func (*AccountEmailVerifiedPredict) _AccountEmailVerified() {}

type AccountEmailVerifiedLoginPredict struct {
	Email    string       `tl:"email"`
	SentCode AuthSentCode `tl:"sent_code"`
}

func (*AccountEmailVerifiedLoginPredict) CRC() uint32 {
	return 0xe1bb0d61
}
func (*AccountEmailVerifiedLoginPredict) _AccountEmailVerified() {}

type AccountEmojiStatuses interface {
	tl.TLObject
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
	Hash     int64         `tl:"hash"`
	Statuses []EmojiStatus `tl:"statuses"`
}

func (*AccountEmojiStatusesPredict) CRC() uint32 {
	return 0x90c467d1
}
func (*AccountEmojiStatusesPredict) _AccountEmojiStatuses() {}

type AccountPassword interface {
	tl.TLObject
	_AccountPassword()
}

var (
	_ AccountPassword = (*AccountPasswordPredict)(nil)
)

type AccountPasswordPredict struct {
	_                       struct{}              `tl:"flags,bitflag"`
	HasRecovery             bool                  `tl:"has_recovery,omitempty:flags:0,implicit"`
	HasSecureValues         bool                  `tl:"has_secure_values,omitempty:flags:1,implicit"`
	HasPassword             bool                  `tl:"has_password,omitempty:flags:2,implicit"`
	CurrentAlgo             PasswordKdfAlgo       `tl:"current_algo,omitempty:flags:2"`
	SRPB                    *[]byte               `tl:"srp_B,omitempty:flags:2"`
	SRPID                   *int64                `tl:"srp_id,omitempty:flags:2"`
	Hint                    *string               `tl:"hint,omitempty:flags:3"`
	EmailUnconfirmedPattern *string               `tl:"email_unconfirmed_pattern,omitempty:flags:4"`
	NewAlgo                 PasswordKdfAlgo       `tl:"new_algo"`
	NewSecureAlgo           SecurePasswordKdfAlgo `tl:"new_secure_algo"`
	SecureRandom            []byte                `tl:"secure_random"`
	PendingResetDate        *int32                `tl:"pending_reset_date,omitempty:flags:5"`
	LoginEmailPattern       *string               `tl:"login_email_pattern,omitempty:flags:6"`
}

func (*AccountPasswordPredict) CRC() uint32 {
	return 0x957b50fb
}
func (*AccountPasswordPredict) _AccountPassword() {}

type AccountPasswordInputSettings interface {
	tl.TLObject
	_AccountPasswordInputSettings()
}

var (
	_ AccountPasswordInputSettings = (*AccountPasswordInputSettingsPredict)(nil)
)

type AccountPasswordInputSettingsPredict struct {
	_                 struct{}             `tl:"flags,bitflag"`
	NewAlgo           PasswordKdfAlgo      `tl:"new_algo,omitempty:flags:0"`
	NewPasswordHash   *[]byte              `tl:"new_password_hash,omitempty:flags:0"`
	Hint              *string              `tl:"hint,omitempty:flags:0"`
	Email             *string              `tl:"email,omitempty:flags:1"`
	NewSecureSettings SecureSecretSettings `tl:"new_secure_settings,omitempty:flags:2"`
}

func (*AccountPasswordInputSettingsPredict) CRC() uint32 {
	return 0xc23727c9
}
func (*AccountPasswordInputSettingsPredict) _AccountPasswordInputSettings() {}

type AccountPasswordSettings interface {
	tl.TLObject
	_AccountPasswordSettings()
}

var (
	_ AccountPasswordSettings = (*AccountPasswordSettingsPredict)(nil)
)

type AccountPasswordSettingsPredict struct {
	_              struct{}             `tl:"flags,bitflag"`
	Email          *string              `tl:"email,omitempty:flags:0"`
	SecureSettings SecureSecretSettings `tl:"secure_settings,omitempty:flags:1"`
}

func (*AccountPasswordSettingsPredict) CRC() uint32 {
	return 0x9a5c33e5
}
func (*AccountPasswordSettingsPredict) _AccountPasswordSettings() {}

type AccountPrivacyRules interface {
	tl.TLObject
	_AccountPrivacyRules()
}

var (
	_ AccountPrivacyRules = (*AccountPrivacyRulesPredict)(nil)
)

type AccountPrivacyRulesPredict struct {
	Rules []PrivacyRule `tl:"rules"`
	Chats []Chat        `tl:"chats"`
	Users []User        `tl:"users"`
}

func (*AccountPrivacyRulesPredict) CRC() uint32 {
	return 0x50a04e45
}
func (*AccountPrivacyRulesPredict) _AccountPrivacyRules() {}

type AccountResetPasswordResult interface {
	tl.TLObject
	_AccountResetPasswordResult()
}

var (
	_ AccountResetPasswordResult = (*AccountResetPasswordFailedWaitPredict)(nil)
	_ AccountResetPasswordResult = (*AccountResetPasswordRequestedWaitPredict)(nil)
	_ AccountResetPasswordResult = (*AccountResetPasswordOkPredict)(nil)
)

type AccountResetPasswordFailedWaitPredict struct {
	RetryDate int32 `tl:"retry_date"`
}

func (*AccountResetPasswordFailedWaitPredict) CRC() uint32 {
	return 0xe3779861
}
func (*AccountResetPasswordFailedWaitPredict) _AccountResetPasswordResult() {}

type AccountResetPasswordRequestedWaitPredict struct {
	UntilDate int32 `tl:"until_date"`
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
	tl.TLObject
	_AccountResolvedBusinessChatLinks()
}

var (
	_ AccountResolvedBusinessChatLinks = (*AccountResolvedBusinessChatLinksPredict)(nil)
)

type AccountResolvedBusinessChatLinksPredict struct {
	_        struct{}        `tl:"flags,bitflag"`
	Peer     Peer            `tl:"peer"`
	Message  string          `tl:"message"`
	Entities []MessageEntity `tl:"entities,omitempty:flags:0"`
	Chats    []Chat          `tl:"chats"`
	Users    []User          `tl:"users"`
}

func (*AccountResolvedBusinessChatLinksPredict) CRC() uint32 {
	return 0x9a23af21
}
func (*AccountResolvedBusinessChatLinksPredict) _AccountResolvedBusinessChatLinks() {}

type AccountSavedRingtone interface {
	tl.TLObject
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
	Document Document `tl:"document"`
}

func (*AccountSavedRingtoneConvertedPredict) CRC() uint32 {
	return 0x1f307eb7
}
func (*AccountSavedRingtoneConvertedPredict) _AccountSavedRingtone() {}

type AccountSavedRingtones interface {
	tl.TLObject
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
	Hash      int64      `tl:"hash"`
	Ringtones []Document `tl:"ringtones"`
}

func (*AccountSavedRingtonesPredict) CRC() uint32 {
	return 0xc1e92cc5
}
func (*AccountSavedRingtonesPredict) _AccountSavedRingtones() {}

type AccountSentEmailCode interface {
	tl.TLObject
	_AccountSentEmailCode()
}

var (
	_ AccountSentEmailCode = (*AccountSentEmailCodePredict)(nil)
)

type AccountSentEmailCodePredict struct {
	EmailPattern string `tl:"email_pattern"`
	Length       int32  `tl:"length"`
}

func (*AccountSentEmailCodePredict) CRC() uint32 {
	return 0x811f854f
}
func (*AccountSentEmailCodePredict) _AccountSentEmailCode() {}

type AccountTakeout interface {
	tl.TLObject
	_AccountTakeout()
}

var (
	_ AccountTakeout = (*AccountTakeoutPredict)(nil)
)

type AccountTakeoutPredict struct {
	ID int64 `tl:"id"`
}

func (*AccountTakeoutPredict) CRC() uint32 {
	return 0x4dba4501
}
func (*AccountTakeoutPredict) _AccountTakeout() {}

type AccountThemes interface {
	tl.TLObject
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
	Hash   int64   `tl:"hash"`
	Themes []Theme `tl:"themes"`
}

func (*AccountThemesPredict) CRC() uint32 {
	return 0x9a3d8c6d
}
func (*AccountThemesPredict) _AccountThemes() {}

type AccountTmpPassword interface {
	tl.TLObject
	_AccountTmpPassword()
}

var (
	_ AccountTmpPassword = (*AccountTmpPasswordPredict)(nil)
)

type AccountTmpPasswordPredict struct {
	TmpPassword []byte `tl:"tmp_password"`
	ValidUntil  int32  `tl:"valid_until"`
}

func (*AccountTmpPasswordPredict) CRC() uint32 {
	return 0xdb64fd34
}
func (*AccountTmpPasswordPredict) _AccountTmpPassword() {}

type AccountWallPapers interface {
	tl.TLObject
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
	Hash       int64       `tl:"hash"`
	Wallpapers []WallPaper `tl:"wallpapers"`
}

func (*AccountWallPapersPredict) CRC() uint32 {
	return 0xcdc3858c
}
func (*AccountWallPapersPredict) _AccountWallPapers() {}

type AccountWebAuthorizations interface {
	tl.TLObject
	_AccountWebAuthorizations()
}

var (
	_ AccountWebAuthorizations = (*AccountWebAuthorizationsPredict)(nil)
)

type AccountWebAuthorizationsPredict struct {
	Authorizations []WebAuthorization `tl:"authorizations"`
	Users          []User             `tl:"users"`
}

func (*AccountWebAuthorizationsPredict) CRC() uint32 {
	return 0xed56c9fc
}
func (*AccountWebAuthorizationsPredict) _AccountWebAuthorizations() {}

type AuthAuthorization interface {
	tl.TLObject
	_AuthAuthorization()
}

var (
	_ AuthAuthorization = (*AuthAuthorizationPredict)(nil)
	_ AuthAuthorization = (*AuthAuthorizationSignUpRequiredPredict)(nil)
)

type AuthAuthorizationPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	SetupPasswordRequired bool     `tl:"setup_password_required,omitempty:flags:1,implicit"`
	OtherwiseReloginDays  *int32   `tl:"otherwise_relogin_days,omitempty:flags:1"`
	TmpSessions           *int32   `tl:"tmp_sessions,omitempty:flags:0"`
	FutureAuthToken       *[]byte  `tl:"future_auth_token,omitempty:flags:2"`
	User                  User     `tl:"user"`
}

func (*AuthAuthorizationPredict) CRC() uint32 {
	return 0x2ea2c0d4
}
func (*AuthAuthorizationPredict) _AuthAuthorization() {}

type AuthAuthorizationSignUpRequiredPredict struct {
	_              struct{}           `tl:"flags,bitflag"`
	TermsOfService HelpTermsOfService `tl:"terms_of_service,omitempty:flags:0"`
}

func (*AuthAuthorizationSignUpRequiredPredict) CRC() uint32 {
	return 0x44747e9a
}
func (*AuthAuthorizationSignUpRequiredPredict) _AuthAuthorization() {}

type AuthCodeType interface {
	tl.TLObject
	_AuthCodeType()
}

var (
	_ AuthCodeType = (*AuthCodeTypeSmsPredict)(nil)
	_ AuthCodeType = (*AuthCodeTypeCallPredict)(nil)
	_ AuthCodeType = (*AuthCodeTypeFlashCallPredict)(nil)
	_ AuthCodeType = (*AuthCodeTypeMissedCallPredict)(nil)
	_ AuthCodeType = (*AuthCodeTypeFragmentSmsPredict)(nil)
)

type AuthCodeTypeSmsPredict struct{}

func (*AuthCodeTypeSmsPredict) CRC() uint32 {
	return 0x72a3158c
}
func (*AuthCodeTypeSmsPredict) _AuthCodeType() {}

type AuthCodeTypeCallPredict struct{}

func (*AuthCodeTypeCallPredict) CRC() uint32 {
	return 0x741cd3e3
}
func (*AuthCodeTypeCallPredict) _AuthCodeType() {}

type AuthCodeTypeFlashCallPredict struct{}

func (*AuthCodeTypeFlashCallPredict) CRC() uint32 {
	return 0x226ccefb
}
func (*AuthCodeTypeFlashCallPredict) _AuthCodeType() {}

type AuthCodeTypeMissedCallPredict struct{}

func (*AuthCodeTypeMissedCallPredict) CRC() uint32 {
	return 0xd61ad6ee
}
func (*AuthCodeTypeMissedCallPredict) _AuthCodeType() {}

type AuthCodeTypeFragmentSmsPredict struct{}

func (*AuthCodeTypeFragmentSmsPredict) CRC() uint32 {
	return 0x06ed998c
}
func (*AuthCodeTypeFragmentSmsPredict) _AuthCodeType() {}

type AuthExportedAuthorization interface {
	tl.TLObject
	_AuthExportedAuthorization()
}

var (
	_ AuthExportedAuthorization = (*AuthExportedAuthorizationPredict)(nil)
)

type AuthExportedAuthorizationPredict struct {
	ID    int64  `tl:"id"`
	Bytes []byte `tl:"bytes"`
}

func (*AuthExportedAuthorizationPredict) CRC() uint32 {
	return 0xb434e2b8
}
func (*AuthExportedAuthorizationPredict) _AuthExportedAuthorization() {}

type AuthLoggedOut interface {
	tl.TLObject
	_AuthLoggedOut()
}

var (
	_ AuthLoggedOut = (*AuthLoggedOutPredict)(nil)
)

type AuthLoggedOutPredict struct {
	_               struct{} `tl:"flags,bitflag"`
	FutureAuthToken *[]byte  `tl:"future_auth_token,omitempty:flags:0"`
}

func (*AuthLoggedOutPredict) CRC() uint32 {
	return 0xc3a2835f
}
func (*AuthLoggedOutPredict) _AuthLoggedOut() {}

type AuthLoginToken interface {
	tl.TLObject
	_AuthLoginToken()
}

var (
	_ AuthLoginToken = (*AuthLoginTokenPredict)(nil)
	_ AuthLoginToken = (*AuthLoginTokenMigrateToPredict)(nil)
	_ AuthLoginToken = (*AuthLoginTokenSuccessPredict)(nil)
)

type AuthLoginTokenPredict struct {
	Expires int32  `tl:"expires"`
	Token   []byte `tl:"token"`
}

func (*AuthLoginTokenPredict) CRC() uint32 {
	return 0x629f1980
}
func (*AuthLoginTokenPredict) _AuthLoginToken() {}

type AuthLoginTokenMigrateToPredict struct {
	DcID  int32  `tl:"dc_id"`
	Token []byte `tl:"token"`
}

func (*AuthLoginTokenMigrateToPredict) CRC() uint32 {
	return 0x068e9916
}
func (*AuthLoginTokenMigrateToPredict) _AuthLoginToken() {}

type AuthLoginTokenSuccessPredict struct {
	Authorization AuthAuthorization `tl:"authorization"`
}

func (*AuthLoginTokenSuccessPredict) CRC() uint32 {
	return 0x390d5c5e
}
func (*AuthLoginTokenSuccessPredict) _AuthLoginToken() {}

type AuthPasswordRecovery interface {
	tl.TLObject
	_AuthPasswordRecovery()
}

var (
	_ AuthPasswordRecovery = (*AuthPasswordRecoveryPredict)(nil)
)

type AuthPasswordRecoveryPredict struct {
	EmailPattern string `tl:"email_pattern"`
}

func (*AuthPasswordRecoveryPredict) CRC() uint32 {
	return 0x137948a5
}
func (*AuthPasswordRecoveryPredict) _AuthPasswordRecovery() {}

type AuthSentCode interface {
	tl.TLObject
	_AuthSentCode()
}

var (
	_ AuthSentCode = (*AuthSentCodePredict)(nil)
	_ AuthSentCode = (*AuthSentCodeSuccessPredict)(nil)
)

type AuthSentCodePredict struct {
	_             struct{}         `tl:"flags,bitflag"`
	Type          AuthSentCodeType `tl:"type"`
	PhoneCodeHash string           `tl:"phone_code_hash"`
	NextType      AuthCodeType     `tl:"next_type,omitempty:flags:1"`
	Timeout       *int32           `tl:"timeout,omitempty:flags:2"`
}

func (*AuthSentCodePredict) CRC() uint32 {
	return 0x5e002502
}
func (*AuthSentCodePredict) _AuthSentCode() {}

type AuthSentCodeSuccessPredict struct {
	Authorization AuthAuthorization `tl:"authorization"`
}

func (*AuthSentCodeSuccessPredict) CRC() uint32 {
	return 0x2390fe44
}
func (*AuthSentCodeSuccessPredict) _AuthSentCode() {}

type AuthSentCodeType interface {
	tl.TLObject
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
	Length int32 `tl:"length"`
}

func (*AuthSentCodeTypeAppPredict) CRC() uint32 {
	return 0x3dbb5986
}
func (*AuthSentCodeTypeAppPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSmsPredict struct {
	Length int32 `tl:"length"`
}

func (*AuthSentCodeTypeSmsPredict) CRC() uint32 {
	return 0xc000bba2
}
func (*AuthSentCodeTypeSmsPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeCallPredict struct {
	Length int32 `tl:"length"`
}

func (*AuthSentCodeTypeCallPredict) CRC() uint32 {
	return 0x5353e5a7
}
func (*AuthSentCodeTypeCallPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeFlashCallPredict struct {
	Pattern string `tl:"pattern"`
}

func (*AuthSentCodeTypeFlashCallPredict) CRC() uint32 {
	return 0xab03c6d9
}
func (*AuthSentCodeTypeFlashCallPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeMissedCallPredict struct {
	Prefix string `tl:"prefix"`
	Length int32  `tl:"length"`
}

func (*AuthSentCodeTypeMissedCallPredict) CRC() uint32 {
	return 0x82006484
}
func (*AuthSentCodeTypeMissedCallPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeEmailCodePredict struct {
	_                    struct{} `tl:"flags,bitflag"`
	AppleSigninAllowed   bool     `tl:"apple_signin_allowed,omitempty:flags:0,implicit"`
	GoogleSigninAllowed  bool     `tl:"google_signin_allowed,omitempty:flags:1,implicit"`
	EmailPattern         string   `tl:"email_pattern"`
	Length               int32    `tl:"length"`
	ResetAvailablePeriod *int32   `tl:"reset_available_period,omitempty:flags:3"`
	ResetPendingDate     *int32   `tl:"reset_pending_date,omitempty:flags:4"`
}

func (*AuthSentCodeTypeEmailCodePredict) CRC() uint32 {
	return 0xf450f59b
}
func (*AuthSentCodeTypeEmailCodePredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSetUpEmailRequiredPredict struct {
	_                   struct{} `tl:"flags,bitflag"`
	AppleSigninAllowed  bool     `tl:"apple_signin_allowed,omitempty:flags:0,implicit"`
	GoogleSigninAllowed bool     `tl:"google_signin_allowed,omitempty:flags:1,implicit"`
}

func (*AuthSentCodeTypeSetUpEmailRequiredPredict) CRC() uint32 {
	return 0xa5491dea
}
func (*AuthSentCodeTypeSetUpEmailRequiredPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeFragmentSmsPredict struct {
	URL    string `tl:"url"`
	Length int32  `tl:"length"`
}

func (*AuthSentCodeTypeFragmentSmsPredict) CRC() uint32 {
	return 0xd9565c39
}
func (*AuthSentCodeTypeFragmentSmsPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeFirebaseSmsPredict struct {
	_                      struct{} `tl:"flags,bitflag"`
	Nonce                  *[]byte  `tl:"nonce,omitempty:flags:0"`
	PlayIntegrityProjectID *int64   `tl:"play_integrity_project_id,omitempty:flags:2"`
	PlayIntegrityNonce     *[]byte  `tl:"play_integrity_nonce,omitempty:flags:2"`
	Receipt                *string  `tl:"receipt,omitempty:flags:1"`
	PushTimeout            *int32   `tl:"push_timeout,omitempty:flags:1"`
	Length                 int32    `tl:"length"`
}

func (*AuthSentCodeTypeFirebaseSmsPredict) CRC() uint32 {
	return 0x009fd736
}
func (*AuthSentCodeTypeFirebaseSmsPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSmsWordPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Beginning *string  `tl:"beginning,omitempty:flags:0"`
}

func (*AuthSentCodeTypeSmsWordPredict) CRC() uint32 {
	return 0xa416ac81
}
func (*AuthSentCodeTypeSmsWordPredict) _AuthSentCodeType() {}

type AuthSentCodeTypeSmsPhrasePredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Beginning *string  `tl:"beginning,omitempty:flags:0"`
}

func (*AuthSentCodeTypeSmsPhrasePredict) CRC() uint32 {
	return 0xb37794af
}
func (*AuthSentCodeTypeSmsPhrasePredict) _AuthSentCodeType() {}

type BotsBotInfo interface {
	tl.TLObject
	_BotsBotInfo()
}

var (
	_ BotsBotInfo = (*BotsBotInfoPredict)(nil)
)

type BotsBotInfoPredict struct {
	Name        string `tl:"name"`
	About       string `tl:"about"`
	Description string `tl:"description"`
}

func (*BotsBotInfoPredict) CRC() uint32 {
	return 0xe8a775b0
}
func (*BotsBotInfoPredict) _BotsBotInfo() {}

type BotsPopularAppBots interface {
	tl.TLObject
	_BotsPopularAppBots()
}

var (
	_ BotsPopularAppBots = (*BotsPopularAppBotsPredict)(nil)
)

type BotsPopularAppBotsPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	NextOffset *string  `tl:"next_offset,omitempty:flags:0"`
	Users      []User   `tl:"users"`
}

func (*BotsPopularAppBotsPredict) CRC() uint32 {
	return 0x1991b13b
}
func (*BotsPopularAppBotsPredict) _BotsPopularAppBots() {}

type BotsPreviewInfo interface {
	tl.TLObject
	_BotsPreviewInfo()
}

var (
	_ BotsPreviewInfo = (*BotsPreviewInfoPredict)(nil)
)

type BotsPreviewInfoPredict struct {
	Media     []BotPreviewMedia `tl:"media"`
	LangCodes []string          `tl:"lang_codes"`
}

func (*BotsPreviewInfoPredict) CRC() uint32 {
	return 0x0ca71d64
}
func (*BotsPreviewInfoPredict) _BotsPreviewInfo() {}

type ChannelsAdminLogResults interface {
	tl.TLObject
	_ChannelsAdminLogResults()
}

var (
	_ ChannelsAdminLogResults = (*ChannelsAdminLogResultsPredict)(nil)
)

type ChannelsAdminLogResultsPredict struct {
	Events []ChannelAdminLogEvent `tl:"events"`
	Chats  []Chat                 `tl:"chats"`
	Users  []User                 `tl:"users"`
}

func (*ChannelsAdminLogResultsPredict) CRC() uint32 {
	return 0xed8af74d
}
func (*ChannelsAdminLogResultsPredict) _ChannelsAdminLogResults() {}

type ChannelsChannelParticipant interface {
	tl.TLObject
	_ChannelsChannelParticipant()
}

var (
	_ ChannelsChannelParticipant = (*ChannelsChannelParticipantPredict)(nil)
)

type ChannelsChannelParticipantPredict struct {
	Participant ChannelParticipant `tl:"participant"`
	Chats       []Chat             `tl:"chats"`
	Users       []User             `tl:"users"`
}

func (*ChannelsChannelParticipantPredict) CRC() uint32 {
	return 0xdfb80317
}
func (*ChannelsChannelParticipantPredict) _ChannelsChannelParticipant() {}

type ChannelsChannelParticipants interface {
	tl.TLObject
	_ChannelsChannelParticipants()
}

var (
	_ ChannelsChannelParticipants = (*ChannelsChannelParticipantsPredict)(nil)
	_ ChannelsChannelParticipants = (*ChannelsChannelParticipantsNotModifiedPredict)(nil)
)

type ChannelsChannelParticipantsPredict struct {
	Count        int32                `tl:"count"`
	Participants []ChannelParticipant `tl:"participants"`
	Chats        []Chat               `tl:"chats"`
	Users        []User               `tl:"users"`
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
	tl.TLObject
	_ChannelsSendAsPeers()
}

var (
	_ ChannelsSendAsPeers = (*ChannelsSendAsPeersPredict)(nil)
)

type ChannelsSendAsPeersPredict struct {
	Peers []SendAsPeer `tl:"peers"`
	Chats []Chat       `tl:"chats"`
	Users []User       `tl:"users"`
}

func (*ChannelsSendAsPeersPredict) CRC() uint32 {
	return 0xf496b0c6
}
func (*ChannelsSendAsPeersPredict) _ChannelsSendAsPeers() {}

type ChannelsSponsoredMessageReportResult interface {
	tl.TLObject
	_ChannelsSponsoredMessageReportResult()
}

var (
	_ ChannelsSponsoredMessageReportResult = (*ChannelsSponsoredMessageReportResultChooseOptionPredict)(nil)
	_ ChannelsSponsoredMessageReportResult = (*ChannelsSponsoredMessageReportResultAdsHiddenPredict)(nil)
	_ ChannelsSponsoredMessageReportResult = (*ChannelsSponsoredMessageReportResultReportedPredict)(nil)
)

type ChannelsSponsoredMessageReportResultChooseOptionPredict struct {
	Title   string                         `tl:"title"`
	Options []SponsoredMessageReportOption `tl:"options"`
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
	tl.TLObject
	_ChatlistsChatlistInvite()
}

var (
	_ ChatlistsChatlistInvite = (*ChatlistsChatlistInviteAlreadyPredict)(nil)
	_ ChatlistsChatlistInvite = (*ChatlistsChatlistInvitePredict)(nil)
)

type ChatlistsChatlistInviteAlreadyPredict struct {
	FilterID     int32  `tl:"filter_id"`
	MissingPeers []Peer `tl:"missing_peers"`
	AlreadyPeers []Peer `tl:"already_peers"`
	Chats        []Chat `tl:"chats"`
	Users        []User `tl:"users"`
}

func (*ChatlistsChatlistInviteAlreadyPredict) CRC() uint32 {
	return 0xfa87f659
}
func (*ChatlistsChatlistInviteAlreadyPredict) _ChatlistsChatlistInvite() {}

type ChatlistsChatlistInvitePredict struct {
	_        struct{} `tl:"flags,bitflag"`
	Title    string   `tl:"title"`
	Emoticon *string  `tl:"emoticon,omitempty:flags:0"`
	Peers    []Peer   `tl:"peers"`
	Chats    []Chat   `tl:"chats"`
	Users    []User   `tl:"users"`
}

func (*ChatlistsChatlistInvitePredict) CRC() uint32 {
	return 0x1dcd839d
}
func (*ChatlistsChatlistInvitePredict) _ChatlistsChatlistInvite() {}

type ChatlistsChatlistUpdates interface {
	tl.TLObject
	_ChatlistsChatlistUpdates()
}

var (
	_ ChatlistsChatlistUpdates = (*ChatlistsChatlistUpdatesPredict)(nil)
)

type ChatlistsChatlistUpdatesPredict struct {
	MissingPeers []Peer `tl:"missing_peers"`
	Chats        []Chat `tl:"chats"`
	Users        []User `tl:"users"`
}

func (*ChatlistsChatlistUpdatesPredict) CRC() uint32 {
	return 0x93bd878d
}
func (*ChatlistsChatlistUpdatesPredict) _ChatlistsChatlistUpdates() {}

type ChatlistsExportedChatlistInvite interface {
	tl.TLObject
	_ChatlistsExportedChatlistInvite()
}

var (
	_ ChatlistsExportedChatlistInvite = (*ChatlistsExportedChatlistInvitePredict)(nil)
)

type ChatlistsExportedChatlistInvitePredict struct {
	Filter DialogFilter           `tl:"filter"`
	Invite ExportedChatlistInvite `tl:"invite"`
}

func (*ChatlistsExportedChatlistInvitePredict) CRC() uint32 {
	return 0x10e6e3a6
}
func (*ChatlistsExportedChatlistInvitePredict) _ChatlistsExportedChatlistInvite() {}

type ChatlistsExportedInvites interface {
	tl.TLObject
	_ChatlistsExportedInvites()
}

var (
	_ ChatlistsExportedInvites = (*ChatlistsExportedInvitesPredict)(nil)
)

type ChatlistsExportedInvitesPredict struct {
	Invites []ExportedChatlistInvite `tl:"invites"`
	Chats   []Chat                   `tl:"chats"`
	Users   []User                   `tl:"users"`
}

func (*ChatlistsExportedInvitesPredict) CRC() uint32 {
	return 0x10ab6dc7
}
func (*ChatlistsExportedInvitesPredict) _ChatlistsExportedInvites() {}

type ContactsBlocked interface {
	tl.TLObject
	_ContactsBlocked()
}

var (
	_ ContactsBlocked = (*ContactsBlockedPredict)(nil)
	_ ContactsBlocked = (*ContactsBlockedSlicePredict)(nil)
)

type ContactsBlockedPredict struct {
	Blocked []PeerBlocked `tl:"blocked"`
	Chats   []Chat        `tl:"chats"`
	Users   []User        `tl:"users"`
}

func (*ContactsBlockedPredict) CRC() uint32 {
	return 0x0ade1591
}
func (*ContactsBlockedPredict) _ContactsBlocked() {}

type ContactsBlockedSlicePredict struct {
	Count   int32         `tl:"count"`
	Blocked []PeerBlocked `tl:"blocked"`
	Chats   []Chat        `tl:"chats"`
	Users   []User        `tl:"users"`
}

func (*ContactsBlockedSlicePredict) CRC() uint32 {
	return 0xe1664194
}
func (*ContactsBlockedSlicePredict) _ContactsBlocked() {}

type ContactsContactBirthdays interface {
	tl.TLObject
	_ContactsContactBirthdays()
}

var (
	_ ContactsContactBirthdays = (*ContactsContactBirthdaysPredict)(nil)
)

type ContactsContactBirthdaysPredict struct {
	Contacts []ContactBirthday `tl:"contacts"`
	Users    []User            `tl:"users"`
}

func (*ContactsContactBirthdaysPredict) CRC() uint32 {
	return 0x114ff30d
}
func (*ContactsContactBirthdaysPredict) _ContactsContactBirthdays() {}

type ContactsContacts interface {
	tl.TLObject
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
	Contacts   []Contact `tl:"contacts"`
	SavedCount int32     `tl:"saved_count"`
	Users      []User    `tl:"users"`
}

func (*ContactsContactsPredict) CRC() uint32 {
	return 0xeae87e42
}
func (*ContactsContactsPredict) _ContactsContacts() {}

type ContactsFound interface {
	tl.TLObject
	_ContactsFound()
}

var (
	_ ContactsFound = (*ContactsFoundPredict)(nil)
)

type ContactsFoundPredict struct {
	MyResults []Peer `tl:"my_results"`
	Results   []Peer `tl:"results"`
	Chats     []Chat `tl:"chats"`
	Users     []User `tl:"users"`
}

func (*ContactsFoundPredict) CRC() uint32 {
	return 0xb3134d9d
}
func (*ContactsFoundPredict) _ContactsFound() {}

type ContactsImportedContacts interface {
	tl.TLObject
	_ContactsImportedContacts()
}

var (
	_ ContactsImportedContacts = (*ContactsImportedContactsPredict)(nil)
)

type ContactsImportedContactsPredict struct {
	Imported       []ImportedContact `tl:"imported"`
	PopularInvites []PopularContact  `tl:"popular_invites"`
	RetryContacts  []int64           `tl:"retry_contacts"`
	Users          []User            `tl:"users"`
}

func (*ContactsImportedContactsPredict) CRC() uint32 {
	return 0x77d01c3b
}
func (*ContactsImportedContactsPredict) _ContactsImportedContacts() {}

type ContactsResolvedPeer interface {
	tl.TLObject
	_ContactsResolvedPeer()
}

var (
	_ ContactsResolvedPeer = (*ContactsResolvedPeerPredict)(nil)
)

type ContactsResolvedPeerPredict struct {
	Peer  Peer   `tl:"peer"`
	Chats []Chat `tl:"chats"`
	Users []User `tl:"users"`
}

func (*ContactsResolvedPeerPredict) CRC() uint32 {
	return 0x7f077ad9
}
func (*ContactsResolvedPeerPredict) _ContactsResolvedPeer() {}

type ContactsTopPeers interface {
	tl.TLObject
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
	Categories []TopPeerCategoryPeers `tl:"categories"`
	Chats      []Chat                 `tl:"chats"`
	Users      []User                 `tl:"users"`
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
	tl.TLObject
	_FragmentCollectibleInfo()
}

var (
	_ FragmentCollectibleInfo = (*FragmentCollectibleInfoPredict)(nil)
)

type FragmentCollectibleInfoPredict struct {
	PurchaseDate   int32  `tl:"purchase_date"`
	Currency       string `tl:"currency"`
	Amount         int64  `tl:"amount"`
	CryptoCurrency string `tl:"crypto_currency"`
	CryptoAmount   int64  `tl:"crypto_amount"`
	URL            string `tl:"url"`
}

func (*FragmentCollectibleInfoPredict) CRC() uint32 {
	return 0x6ebdff91
}
func (*FragmentCollectibleInfoPredict) _FragmentCollectibleInfo() {}

type HelpAppConfig interface {
	tl.TLObject
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
	Hash   int32     `tl:"hash"`
	Config JSONValue `tl:"config"`
}

func (*HelpAppConfigPredict) CRC() uint32 {
	return 0xdd18782e
}
func (*HelpAppConfigPredict) _HelpAppConfig() {}

type HelpAppUpdate interface {
	tl.TLObject
	_HelpAppUpdate()
}

var (
	_ HelpAppUpdate = (*HelpAppUpdatePredict)(nil)
	_ HelpAppUpdate = (*HelpNoAppUpdatePredict)(nil)
)

type HelpAppUpdatePredict struct {
	_          struct{}        `tl:"flags,bitflag"`
	CanNotSkip bool            `tl:"can_not_skip,omitempty:flags:0,implicit"`
	ID         int32           `tl:"id"`
	Version    string          `tl:"version"`
	Text       string          `tl:"text"`
	Entities   []MessageEntity `tl:"entities"`
	Document   Document        `tl:"document,omitempty:flags:1"`
	URL        *string         `tl:"url,omitempty:flags:2"`
	Sticker    Document        `tl:"sticker,omitempty:flags:3"`
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
	tl.TLObject
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
	Countries []HelpCountry `tl:"countries"`
	Hash      int32         `tl:"hash"`
}

func (*HelpCountriesListPredict) CRC() uint32 {
	return 0x87d0759e
}
func (*HelpCountriesListPredict) _HelpCountriesList() {}

type HelpCountry interface {
	tl.TLObject
	_HelpCountry()
}

var (
	_ HelpCountry = (*HelpCountryPredict)(nil)
)

type HelpCountryPredict struct {
	_            struct{}          `tl:"flags,bitflag"`
	Hidden       bool              `tl:"hidden,omitempty:flags:0,implicit"`
	Iso2         string            `tl:"iso2"`
	DefaultName  string            `tl:"default_name"`
	Name         *string           `tl:"name,omitempty:flags:1"`
	CountryCodes []HelpCountryCode `tl:"country_codes"`
}

func (*HelpCountryPredict) CRC() uint32 {
	return 0xc3878e23
}
func (*HelpCountryPredict) _HelpCountry() {}

type HelpCountryCode interface {
	tl.TLObject
	_HelpCountryCode()
}

var (
	_ HelpCountryCode = (*HelpCountryCodePredict)(nil)
)

type HelpCountryCodePredict struct {
	_           struct{} `tl:"flags,bitflag"`
	CountryCode string   `tl:"country_code"`
	Prefixes    []string `tl:"prefixes,omitempty:flags:0"`
	Patterns    []string `tl:"patterns,omitempty:flags:1"`
}

func (*HelpCountryCodePredict) CRC() uint32 {
	return 0x4203c5ef
}
func (*HelpCountryCodePredict) _HelpCountryCode() {}

type HelpDeepLinkInfo interface {
	tl.TLObject
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
	_         struct{}        `tl:"flags,bitflag"`
	UpdateApp bool            `tl:"update_app,omitempty:flags:0,implicit"`
	Message   string          `tl:"message"`
	Entities  []MessageEntity `tl:"entities,omitempty:flags:1"`
}

func (*HelpDeepLinkInfoPredict) CRC() uint32 {
	return 0x6a4ee832
}
func (*HelpDeepLinkInfoPredict) _HelpDeepLinkInfo() {}

type HelpInviteText interface {
	tl.TLObject
	_HelpInviteText()
}

var (
	_ HelpInviteText = (*HelpInviteTextPredict)(nil)
)

type HelpInviteTextPredict struct {
	Message string `tl:"message"`
}

func (*HelpInviteTextPredict) CRC() uint32 {
	return 0x18cb9f78
}
func (*HelpInviteTextPredict) _HelpInviteText() {}

type HelpPassportConfig interface {
	tl.TLObject
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
	Hash           int32    `tl:"hash"`
	CountriesLangs DataJSON `tl:"countries_langs"`
}

func (*HelpPassportConfigPredict) CRC() uint32 {
	return 0xa098d6af
}
func (*HelpPassportConfigPredict) _HelpPassportConfig() {}

type HelpPeerColorOption interface {
	tl.TLObject
	_HelpPeerColorOption()
}

var (
	_ HelpPeerColorOption = (*HelpPeerColorOptionPredict)(nil)
)

type HelpPeerColorOptionPredict struct {
	_               struct{}         `tl:"flags,bitflag"`
	Hidden          bool             `tl:"hidden,omitempty:flags:0,implicit"`
	ColorID         int32            `tl:"color_id"`
	Colors          HelpPeerColorSet `tl:"colors,omitempty:flags:1"`
	DarkColors      HelpPeerColorSet `tl:"dark_colors,omitempty:flags:2"`
	ChannelMinLevel *int32           `tl:"channel_min_level,omitempty:flags:3"`
	GroupMinLevel   *int32           `tl:"group_min_level,omitempty:flags:4"`
}

func (*HelpPeerColorOptionPredict) CRC() uint32 {
	return 0xadec6ebe
}
func (*HelpPeerColorOptionPredict) _HelpPeerColorOption() {}

type HelpPeerColorSet interface {
	tl.TLObject
	_HelpPeerColorSet()
}

var (
	_ HelpPeerColorSet = (*HelpPeerColorSetPredict)(nil)
	_ HelpPeerColorSet = (*HelpPeerColorProfileSetPredict)(nil)
)

type HelpPeerColorSetPredict struct {
	Colors []int32 `tl:"colors"`
}

func (*HelpPeerColorSetPredict) CRC() uint32 {
	return 0x26219a58
}
func (*HelpPeerColorSetPredict) _HelpPeerColorSet() {}

type HelpPeerColorProfileSetPredict struct {
	PaletteColors []int32 `tl:"palette_colors"`
	BgColors      []int32 `tl:"bg_colors"`
	StoryColors   []int32 `tl:"story_colors"`
}

func (*HelpPeerColorProfileSetPredict) CRC() uint32 {
	return 0x767d61eb
}
func (*HelpPeerColorProfileSetPredict) _HelpPeerColorSet() {}

type HelpPeerColors interface {
	tl.TLObject
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
	Hash   int32                 `tl:"hash"`
	Colors []HelpPeerColorOption `tl:"colors"`
}

func (*HelpPeerColorsPredict) CRC() uint32 {
	return 0x00f8ed08
}
func (*HelpPeerColorsPredict) _HelpPeerColors() {}

type HelpPremiumPromo interface {
	tl.TLObject
	_HelpPremiumPromo()
}

var (
	_ HelpPremiumPromo = (*HelpPremiumPromoPredict)(nil)
)

type HelpPremiumPromoPredict struct {
	StatusText     string                      `tl:"status_text"`
	StatusEntities []MessageEntity             `tl:"status_entities"`
	VideoSections  []string                    `tl:"video_sections"`
	Videos         []Document                  `tl:"videos"`
	PeriodOptions  []PremiumSubscriptionOption `tl:"period_options"`
	Users          []User                      `tl:"users"`
}

func (*HelpPremiumPromoPredict) CRC() uint32 {
	return 0x5334759c
}
func (*HelpPremiumPromoPredict) _HelpPremiumPromo() {}

type HelpPromoData interface {
	tl.TLObject
	_HelpPromoData()
}

var (
	_ HelpPromoData = (*HelpPromoDataEmptyPredict)(nil)
	_ HelpPromoData = (*HelpPromoDataPredict)(nil)
)

type HelpPromoDataEmptyPredict struct {
	Expires int32 `tl:"expires"`
}

func (*HelpPromoDataEmptyPredict) CRC() uint32 {
	return 0x98f6ac75
}
func (*HelpPromoDataEmptyPredict) _HelpPromoData() {}

type HelpPromoDataPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Proxy      bool     `tl:"proxy,omitempty:flags:0,implicit"`
	Expires    int32    `tl:"expires"`
	Peer       Peer     `tl:"peer"`
	Chats      []Chat   `tl:"chats"`
	Users      []User   `tl:"users"`
	PsaType    *string  `tl:"psa_type,omitempty:flags:1"`
	PsaMessage *string  `tl:"psa_message,omitempty:flags:2"`
}

func (*HelpPromoDataPredict) CRC() uint32 {
	return 0x8c39793f
}
func (*HelpPromoDataPredict) _HelpPromoData() {}

type HelpRecentMeUrls interface {
	tl.TLObject
	_HelpRecentMeUrls()
}

var (
	_ HelpRecentMeUrls = (*HelpRecentMeUrlsPredict)(nil)
)

type HelpRecentMeUrlsPredict struct {
	Urls  []RecentMeURL `tl:"urls"`
	Chats []Chat        `tl:"chats"`
	Users []User        `tl:"users"`
}

func (*HelpRecentMeUrlsPredict) CRC() uint32 {
	return 0x0e0310d7
}
func (*HelpRecentMeUrlsPredict) _HelpRecentMeUrls() {}

type HelpSupport interface {
	tl.TLObject
	_HelpSupport()
}

var (
	_ HelpSupport = (*HelpSupportPredict)(nil)
)

type HelpSupportPredict struct {
	PhoneNumber string `tl:"phone_number"`
	User        User   `tl:"user"`
}

func (*HelpSupportPredict) CRC() uint32 {
	return 0x17c6b5f6
}
func (*HelpSupportPredict) _HelpSupport() {}

type HelpSupportName interface {
	tl.TLObject
	_HelpSupportName()
}

var (
	_ HelpSupportName = (*HelpSupportNamePredict)(nil)
)

type HelpSupportNamePredict struct {
	Name string `tl:"name"`
}

func (*HelpSupportNamePredict) CRC() uint32 {
	return 0x8c05f1c9
}
func (*HelpSupportNamePredict) _HelpSupportName() {}

type HelpTermsOfService interface {
	tl.TLObject
	_HelpTermsOfService()
}

var (
	_ HelpTermsOfService = (*HelpTermsOfServicePredict)(nil)
)

type HelpTermsOfServicePredict struct {
	_             struct{}        `tl:"flags,bitflag"`
	Popup         bool            `tl:"popup,omitempty:flags:0,implicit"`
	ID            DataJSON        `tl:"id"`
	Text          string          `tl:"text"`
	Entities      []MessageEntity `tl:"entities"`
	MinAgeConfirm *int32          `tl:"min_age_confirm,omitempty:flags:1"`
}

func (*HelpTermsOfServicePredict) CRC() uint32 {
	return 0x780a0310
}
func (*HelpTermsOfServicePredict) _HelpTermsOfService() {}

type HelpTermsOfServiceUpdate interface {
	tl.TLObject
	_HelpTermsOfServiceUpdate()
}

var (
	_ HelpTermsOfServiceUpdate = (*HelpTermsOfServiceUpdateEmptyPredict)(nil)
	_ HelpTermsOfServiceUpdate = (*HelpTermsOfServiceUpdatePredict)(nil)
)

type HelpTermsOfServiceUpdateEmptyPredict struct {
	Expires int32 `tl:"expires"`
}

func (*HelpTermsOfServiceUpdateEmptyPredict) CRC() uint32 {
	return 0xe3309f7f
}
func (*HelpTermsOfServiceUpdateEmptyPredict) _HelpTermsOfServiceUpdate() {}

type HelpTermsOfServiceUpdatePredict struct {
	Expires        int32              `tl:"expires"`
	TermsOfService HelpTermsOfService `tl:"terms_of_service"`
}

func (*HelpTermsOfServiceUpdatePredict) CRC() uint32 {
	return 0x28ecf961
}
func (*HelpTermsOfServiceUpdatePredict) _HelpTermsOfServiceUpdate() {}

type HelpTimezonesList interface {
	tl.TLObject
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
	Timezones []Timezone `tl:"timezones"`
	Hash      int32      `tl:"hash"`
}

func (*HelpTimezonesListPredict) CRC() uint32 {
	return 0x7b74ed71
}
func (*HelpTimezonesListPredict) _HelpTimezonesList() {}

type HelpUserInfo interface {
	tl.TLObject
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
	Message  string          `tl:"message"`
	Entities []MessageEntity `tl:"entities"`
	Author   string          `tl:"author"`
	Date     int32           `tl:"date"`
}

func (*HelpUserInfoPredict) CRC() uint32 {
	return 0x01eb3758
}
func (*HelpUserInfoPredict) _HelpUserInfo() {}

type MessagesAffectedFoundMessages interface {
	tl.TLObject
	_MessagesAffectedFoundMessages()
}

var (
	_ MessagesAffectedFoundMessages = (*MessagesAffectedFoundMessagesPredict)(nil)
)

type MessagesAffectedFoundMessagesPredict struct {
	Pts      int32   `tl:"pts"`
	PtsCount int32   `tl:"pts_count"`
	Offset   int32   `tl:"offset"`
	Messages []int32 `tl:"messages"`
}

func (*MessagesAffectedFoundMessagesPredict) CRC() uint32 {
	return 0xef8d3e6c
}
func (*MessagesAffectedFoundMessagesPredict) _MessagesAffectedFoundMessages() {}

type MessagesAffectedHistory interface {
	tl.TLObject
	_MessagesAffectedHistory()
}

var (
	_ MessagesAffectedHistory = (*MessagesAffectedHistoryPredict)(nil)
)

type MessagesAffectedHistoryPredict struct {
	Pts      int32 `tl:"pts"`
	PtsCount int32 `tl:"pts_count"`
	Offset   int32 `tl:"offset"`
}

func (*MessagesAffectedHistoryPredict) CRC() uint32 {
	return 0xb45c69d1
}
func (*MessagesAffectedHistoryPredict) _MessagesAffectedHistory() {}

type MessagesAffectedMessages interface {
	tl.TLObject
	_MessagesAffectedMessages()
}

var (
	_ MessagesAffectedMessages = (*MessagesAffectedMessagesPredict)(nil)
)

type MessagesAffectedMessagesPredict struct {
	Pts      int32 `tl:"pts"`
	PtsCount int32 `tl:"pts_count"`
}

func (*MessagesAffectedMessagesPredict) CRC() uint32 {
	return 0x84d19185
}
func (*MessagesAffectedMessagesPredict) _MessagesAffectedMessages() {}

type MessagesAllStickers interface {
	tl.TLObject
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
	Hash int64        `tl:"hash"`
	Sets []StickerSet `tl:"sets"`
}

func (*MessagesAllStickersPredict) CRC() uint32 {
	return 0xcdbbcebb
}
func (*MessagesAllStickersPredict) _MessagesAllStickers() {}

type MessagesArchivedStickers interface {
	tl.TLObject
	_MessagesArchivedStickers()
}

var (
	_ MessagesArchivedStickers = (*MessagesArchivedStickersPredict)(nil)
)

type MessagesArchivedStickersPredict struct {
	Count int32               `tl:"count"`
	Sets  []StickerSetCovered `tl:"sets"`
}

func (*MessagesArchivedStickersPredict) CRC() uint32 {
	return 0x4fcba9c8
}
func (*MessagesArchivedStickersPredict) _MessagesArchivedStickers() {}

type MessagesAvailableEffects interface {
	tl.TLObject
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
	Hash      int32             `tl:"hash"`
	Effects   []AvailableEffect `tl:"effects"`
	Documents []Document        `tl:"documents"`
}

func (*MessagesAvailableEffectsPredict) CRC() uint32 {
	return 0xbddb616e
}
func (*MessagesAvailableEffectsPredict) _MessagesAvailableEffects() {}

type MessagesAvailableReactions interface {
	tl.TLObject
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
	Hash      int32               `tl:"hash"`
	Reactions []AvailableReaction `tl:"reactions"`
}

func (*MessagesAvailableReactionsPredict) CRC() uint32 {
	return 0x768e3aad
}
func (*MessagesAvailableReactionsPredict) _MessagesAvailableReactions() {}

type MessagesBotApp interface {
	tl.TLObject
	_MessagesBotApp()
}

var (
	_ MessagesBotApp = (*MessagesBotAppPredict)(nil)
)

type MessagesBotAppPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	Inactive           bool     `tl:"inactive,omitempty:flags:0,implicit"`
	RequestWriteAccess bool     `tl:"request_write_access,omitempty:flags:1,implicit"`
	HasSettings        bool     `tl:"has_settings,omitempty:flags:2,implicit"`
	App                BotApp   `tl:"app"`
}

func (*MessagesBotAppPredict) CRC() uint32 {
	return 0xeb50adf5
}
func (*MessagesBotAppPredict) _MessagesBotApp() {}

type MessagesBotCallbackAnswer interface {
	tl.TLObject
	_MessagesBotCallbackAnswer()
}

var (
	_ MessagesBotCallbackAnswer = (*MessagesBotCallbackAnswerPredict)(nil)
)

type MessagesBotCallbackAnswerPredict struct {
	_         struct{} `tl:"flags,bitflag"`
	Alert     bool     `tl:"alert,omitempty:flags:1,implicit"`
	HasURL    bool     `tl:"has_url,omitempty:flags:3,implicit"`
	NativeUi  bool     `tl:"native_ui,omitempty:flags:4,implicit"`
	Message   *string  `tl:"message,omitempty:flags:0"`
	URL       *string  `tl:"url,omitempty:flags:2"`
	CacheTime int32    `tl:"cache_time"`
}

func (*MessagesBotCallbackAnswerPredict) CRC() uint32 {
	return 0x36585ea4
}
func (*MessagesBotCallbackAnswerPredict) _MessagesBotCallbackAnswer() {}

type MessagesBotResults interface {
	tl.TLObject
	_MessagesBotResults()
}

var (
	_ MessagesBotResults = (*MessagesBotResultsPredict)(nil)
)

type MessagesBotResultsPredict struct {
	_             struct{}          `tl:"flags,bitflag"`
	Gallery       bool              `tl:"gallery,omitempty:flags:0,implicit"`
	QueryID       int64             `tl:"query_id"`
	NextOffset    *string           `tl:"next_offset,omitempty:flags:1"`
	SwitchPm      InlineBotSwitchPm `tl:"switch_pm,omitempty:flags:2"`
	SwitchWebview InlineBotWebView  `tl:"switch_webview,omitempty:flags:3"`
	Results       []BotInlineResult `tl:"results"`
	CacheTime     int32             `tl:"cache_time"`
	Users         []User            `tl:"users"`
}

func (*MessagesBotResultsPredict) CRC() uint32 {
	return 0xe021f2f6
}
func (*MessagesBotResultsPredict) _MessagesBotResults() {}

type MessagesChatAdminsWithInvites interface {
	tl.TLObject
	_MessagesChatAdminsWithInvites()
}

var (
	_ MessagesChatAdminsWithInvites = (*MessagesChatAdminsWithInvitesPredict)(nil)
)

type MessagesChatAdminsWithInvitesPredict struct {
	Admins []ChatAdminWithInvites `tl:"admins"`
	Users  []User                 `tl:"users"`
}

func (*MessagesChatAdminsWithInvitesPredict) CRC() uint32 {
	return 0xb69b72d7
}
func (*MessagesChatAdminsWithInvitesPredict) _MessagesChatAdminsWithInvites() {}

type MessagesChatFull interface {
	tl.TLObject
	_MessagesChatFull()
}

var (
	_ MessagesChatFull = (*MessagesChatFullPredict)(nil)
)

type MessagesChatFullPredict struct {
	FullChat ChatFull `tl:"full_chat"`
	Chats    []Chat   `tl:"chats"`
	Users    []User   `tl:"users"`
}

func (*MessagesChatFullPredict) CRC() uint32 {
	return 0xe5d7d19c
}
func (*MessagesChatFullPredict) _MessagesChatFull() {}

type MessagesChatInviteImporters interface {
	tl.TLObject
	_MessagesChatInviteImporters()
}

var (
	_ MessagesChatInviteImporters = (*MessagesChatInviteImportersPredict)(nil)
)

type MessagesChatInviteImportersPredict struct {
	Count     int32                `tl:"count"`
	Importers []ChatInviteImporter `tl:"importers"`
	Users     []User               `tl:"users"`
}

func (*MessagesChatInviteImportersPredict) CRC() uint32 {
	return 0x81b6b00a
}
func (*MessagesChatInviteImportersPredict) _MessagesChatInviteImporters() {}

type MessagesChats interface {
	tl.TLObject
	_MessagesChats()
}

var (
	_ MessagesChats = (*MessagesChatsPredict)(nil)
	_ MessagesChats = (*MessagesChatsSlicePredict)(nil)
)

type MessagesChatsPredict struct {
	Chats []Chat `tl:"chats"`
}

func (*MessagesChatsPredict) CRC() uint32 {
	return 0x64ff9fd5
}
func (*MessagesChatsPredict) _MessagesChats() {}

type MessagesChatsSlicePredict struct {
	Count int32  `tl:"count"`
	Chats []Chat `tl:"chats"`
}

func (*MessagesChatsSlicePredict) CRC() uint32 {
	return 0x9cd81144
}
func (*MessagesChatsSlicePredict) _MessagesChats() {}

type MessagesCheckedHistoryImportPeer interface {
	tl.TLObject
	_MessagesCheckedHistoryImportPeer()
}

var (
	_ MessagesCheckedHistoryImportPeer = (*MessagesCheckedHistoryImportPeerPredict)(nil)
)

type MessagesCheckedHistoryImportPeerPredict struct {
	ConfirmText string `tl:"confirm_text"`
}

func (*MessagesCheckedHistoryImportPeerPredict) CRC() uint32 {
	return 0xa24de717
}
func (*MessagesCheckedHistoryImportPeerPredict) _MessagesCheckedHistoryImportPeer() {}

type MessagesDhConfig interface {
	tl.TLObject
	_MessagesDhConfig()
}

var (
	_ MessagesDhConfig = (*MessagesDhConfigNotModifiedPredict)(nil)
	_ MessagesDhConfig = (*MessagesDhConfigPredict)(nil)
)

type MessagesDhConfigNotModifiedPredict struct {
	Random []byte `tl:"random"`
}

func (*MessagesDhConfigNotModifiedPredict) CRC() uint32 {
	return 0xc0e24635
}
func (*MessagesDhConfigNotModifiedPredict) _MessagesDhConfig() {}

type MessagesDhConfigPredict struct {
	G       int32  `tl:"g"`
	P       []byte `tl:"p"`
	Version int32  `tl:"version"`
	Random  []byte `tl:"random"`
}

func (*MessagesDhConfigPredict) CRC() uint32 {
	return 0x2c221edd
}
func (*MessagesDhConfigPredict) _MessagesDhConfig() {}

type MessagesDialogFilters interface {
	tl.TLObject
	_MessagesDialogFilters()
}

var (
	_ MessagesDialogFilters = (*MessagesDialogFiltersPredict)(nil)
)

type MessagesDialogFiltersPredict struct {
	_           struct{}       `tl:"flags,bitflag"`
	TagsEnabled bool           `tl:"tags_enabled,omitempty:flags:0,implicit"`
	Filters     []DialogFilter `tl:"filters"`
}

func (*MessagesDialogFiltersPredict) CRC() uint32 {
	return 0x2ad93719
}
func (*MessagesDialogFiltersPredict) _MessagesDialogFilters() {}

type MessagesDialogs interface {
	tl.TLObject
	_MessagesDialogs()
}

var (
	_ MessagesDialogs = (*MessagesDialogsPredict)(nil)
	_ MessagesDialogs = (*MessagesDialogsSlicePredict)(nil)
	_ MessagesDialogs = (*MessagesDialogsNotModifiedPredict)(nil)
)

type MessagesDialogsPredict struct {
	Dialogs  []Dialog  `tl:"dialogs"`
	Messages []Message `tl:"messages"`
	Chats    []Chat    `tl:"chats"`
	Users    []User    `tl:"users"`
}

func (*MessagesDialogsPredict) CRC() uint32 {
	return 0x15ba6c40
}
func (*MessagesDialogsPredict) _MessagesDialogs() {}

type MessagesDialogsSlicePredict struct {
	Count    int32     `tl:"count"`
	Dialogs  []Dialog  `tl:"dialogs"`
	Messages []Message `tl:"messages"`
	Chats    []Chat    `tl:"chats"`
	Users    []User    `tl:"users"`
}

func (*MessagesDialogsSlicePredict) CRC() uint32 {
	return 0x71e094f3
}
func (*MessagesDialogsSlicePredict) _MessagesDialogs() {}

type MessagesDialogsNotModifiedPredict struct {
	Count int32 `tl:"count"`
}

func (*MessagesDialogsNotModifiedPredict) CRC() uint32 {
	return 0xf0e3e596
}
func (*MessagesDialogsNotModifiedPredict) _MessagesDialogs() {}

type MessagesDiscussionMessage interface {
	tl.TLObject
	_MessagesDiscussionMessage()
}

var (
	_ MessagesDiscussionMessage = (*MessagesDiscussionMessagePredict)(nil)
)

type MessagesDiscussionMessagePredict struct {
	_               struct{}  `tl:"flags,bitflag"`
	Messages        []Message `tl:"messages"`
	MaxID           *int32    `tl:"max_id,omitempty:flags:0"`
	ReadInboxMaxID  *int32    `tl:"read_inbox_max_id,omitempty:flags:1"`
	ReadOutboxMaxID *int32    `tl:"read_outbox_max_id,omitempty:flags:2"`
	UnreadCount     int32     `tl:"unread_count"`
	Chats           []Chat    `tl:"chats"`
	Users           []User    `tl:"users"`
}

func (*MessagesDiscussionMessagePredict) CRC() uint32 {
	return 0xa6341782
}
func (*MessagesDiscussionMessagePredict) _MessagesDiscussionMessage() {}

type MessagesEmojiGroups interface {
	tl.TLObject
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
	Hash   int32        `tl:"hash"`
	Groups []EmojiGroup `tl:"groups"`
}

func (*MessagesEmojiGroupsPredict) CRC() uint32 {
	return 0x881fb94b
}
func (*MessagesEmojiGroupsPredict) _MessagesEmojiGroups() {}

type MessagesExportedChatInvite interface {
	tl.TLObject
	_MessagesExportedChatInvite()
}

var (
	_ MessagesExportedChatInvite = (*MessagesExportedChatInvitePredict)(nil)
	_ MessagesExportedChatInvite = (*MessagesExportedChatInviteReplacedPredict)(nil)
)

type MessagesExportedChatInvitePredict struct {
	Invite ExportedChatInvite `tl:"invite"`
	Users  []User             `tl:"users"`
}

func (*MessagesExportedChatInvitePredict) CRC() uint32 {
	return 0x1871be50
}
func (*MessagesExportedChatInvitePredict) _MessagesExportedChatInvite() {}

type MessagesExportedChatInviteReplacedPredict struct {
	Invite    ExportedChatInvite `tl:"invite"`
	NewInvite ExportedChatInvite `tl:"new_invite"`
	Users     []User             `tl:"users"`
}

func (*MessagesExportedChatInviteReplacedPredict) CRC() uint32 {
	return 0x222600ef
}
func (*MessagesExportedChatInviteReplacedPredict) _MessagesExportedChatInvite() {}

type MessagesExportedChatInvites interface {
	tl.TLObject
	_MessagesExportedChatInvites()
}

var (
	_ MessagesExportedChatInvites = (*MessagesExportedChatInvitesPredict)(nil)
)

type MessagesExportedChatInvitesPredict struct {
	Count   int32                `tl:"count"`
	Invites []ExportedChatInvite `tl:"invites"`
	Users   []User               `tl:"users"`
}

func (*MessagesExportedChatInvitesPredict) CRC() uint32 {
	return 0xbdc62dcc
}
func (*MessagesExportedChatInvitesPredict) _MessagesExportedChatInvites() {}

type MessagesFavedStickers interface {
	tl.TLObject
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
	Hash     int64         `tl:"hash"`
	Packs    []StickerPack `tl:"packs"`
	Stickers []Document    `tl:"stickers"`
}

func (*MessagesFavedStickersPredict) CRC() uint32 {
	return 0x2cb51097
}
func (*MessagesFavedStickersPredict) _MessagesFavedStickers() {}

type MessagesFeaturedStickers interface {
	tl.TLObject
	_MessagesFeaturedStickers()
}

var (
	_ MessagesFeaturedStickers = (*MessagesFeaturedStickersNotModifiedPredict)(nil)
	_ MessagesFeaturedStickers = (*MessagesFeaturedStickersPredict)(nil)
)

type MessagesFeaturedStickersNotModifiedPredict struct {
	Count int32 `tl:"count"`
}

func (*MessagesFeaturedStickersNotModifiedPredict) CRC() uint32 {
	return 0xc6dc0c66
}
func (*MessagesFeaturedStickersNotModifiedPredict) _MessagesFeaturedStickers() {}

type MessagesFeaturedStickersPredict struct {
	_       struct{}            `tl:"flags,bitflag"`
	Premium bool                `tl:"premium,omitempty:flags:0,implicit"`
	Hash    int64               `tl:"hash"`
	Count   int32               `tl:"count"`
	Sets    []StickerSetCovered `tl:"sets"`
	Unread  []int64             `tl:"unread"`
}

func (*MessagesFeaturedStickersPredict) CRC() uint32 {
	return 0xbe382906
}
func (*MessagesFeaturedStickersPredict) _MessagesFeaturedStickers() {}

type MessagesForumTopics interface {
	tl.TLObject
	_MessagesForumTopics()
}

var (
	_ MessagesForumTopics = (*MessagesForumTopicsPredict)(nil)
)

type MessagesForumTopicsPredict struct {
	_                 struct{}     `tl:"flags,bitflag"`
	OrderByCreateDate bool         `tl:"order_by_create_date,omitempty:flags:0,implicit"`
	Count             int32        `tl:"count"`
	Topics            []ForumTopic `tl:"topics"`
	Messages          []Message    `tl:"messages"`
	Chats             []Chat       `tl:"chats"`
	Users             []User       `tl:"users"`
	Pts               int32        `tl:"pts"`
}

func (*MessagesForumTopicsPredict) CRC() uint32 {
	return 0x367617d3
}
func (*MessagesForumTopicsPredict) _MessagesForumTopics() {}

type MessagesFoundStickerSets interface {
	tl.TLObject
	_MessagesFoundStickerSets()
}

var (
	_ MessagesFoundStickerSets = (*MessagesFoundStickerSetsNotModifiedPredict)(nil)
	_ MessagesFoundStickerSets = (*MessagesFoundStickerSetsPredict)(nil)
)

type MessagesFoundStickerSetsNotModifiedPredict struct{}

func (*MessagesFoundStickerSetsNotModifiedPredict) CRC() uint32 {
	return 0x0d54b65d
}
func (*MessagesFoundStickerSetsNotModifiedPredict) _MessagesFoundStickerSets() {}

type MessagesFoundStickerSetsPredict struct {
	Hash int64               `tl:"hash"`
	Sets []StickerSetCovered `tl:"sets"`
}

func (*MessagesFoundStickerSetsPredict) CRC() uint32 {
	return 0x8af09dd2
}
func (*MessagesFoundStickerSetsPredict) _MessagesFoundStickerSets() {}

type MessagesHighScores interface {
	tl.TLObject
	_MessagesHighScores()
}

var (
	_ MessagesHighScores = (*MessagesHighScoresPredict)(nil)
)

type MessagesHighScoresPredict struct {
	Scores []HighScore `tl:"scores"`
	Users  []User      `tl:"users"`
}

func (*MessagesHighScoresPredict) CRC() uint32 {
	return 0x9a3bfd99
}
func (*MessagesHighScoresPredict) _MessagesHighScores() {}

type MessagesHistoryImport interface {
	tl.TLObject
	_MessagesHistoryImport()
}

var (
	_ MessagesHistoryImport = (*MessagesHistoryImportPredict)(nil)
)

type MessagesHistoryImportPredict struct {
	ID int64 `tl:"id"`
}

func (*MessagesHistoryImportPredict) CRC() uint32 {
	return 0x1662af0b
}
func (*MessagesHistoryImportPredict) _MessagesHistoryImport() {}

type MessagesHistoryImportParsed interface {
	tl.TLObject
	_MessagesHistoryImportParsed()
}

var (
	_ MessagesHistoryImportParsed = (*MessagesHistoryImportParsedPredict)(nil)
)

type MessagesHistoryImportParsedPredict struct {
	_     struct{} `tl:"flags,bitflag"`
	Pm    bool     `tl:"pm,omitempty:flags:0,implicit"`
	Group bool     `tl:"group,omitempty:flags:1,implicit"`
	Title *string  `tl:"title,omitempty:flags:2"`
}

func (*MessagesHistoryImportParsedPredict) CRC() uint32 {
	return 0x5e0fb7b9
}
func (*MessagesHistoryImportParsedPredict) _MessagesHistoryImportParsed() {}

type MessagesInactiveChats interface {
	tl.TLObject
	_MessagesInactiveChats()
}

var (
	_ MessagesInactiveChats = (*MessagesInactiveChatsPredict)(nil)
)

type MessagesInactiveChatsPredict struct {
	Dates []int32 `tl:"dates"`
	Chats []Chat  `tl:"chats"`
	Users []User  `tl:"users"`
}

func (*MessagesInactiveChatsPredict) CRC() uint32 {
	return 0xa927fec5
}
func (*MessagesInactiveChatsPredict) _MessagesInactiveChats() {}

type MessagesInvitedUsers interface {
	tl.TLObject
	_MessagesInvitedUsers()
}

var (
	_ MessagesInvitedUsers = (*MessagesInvitedUsersPredict)(nil)
)

type MessagesInvitedUsersPredict struct {
	Updates         Updates          `tl:"updates"`
	MissingInvitees []MissingInvitee `tl:"missing_invitees"`
}

func (*MessagesInvitedUsersPredict) CRC() uint32 {
	return 0x7f5defa6
}
func (*MessagesInvitedUsersPredict) _MessagesInvitedUsers() {}

type MessagesMessageEditData interface {
	tl.TLObject
	_MessagesMessageEditData()
}

var (
	_ MessagesMessageEditData = (*MessagesMessageEditDataPredict)(nil)
)

type MessagesMessageEditDataPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Caption bool     `tl:"caption,omitempty:flags:0,implicit"`
}

func (*MessagesMessageEditDataPredict) CRC() uint32 {
	return 0x26b5dde6
}
func (*MessagesMessageEditDataPredict) _MessagesMessageEditData() {}

type MessagesMessageReactionsList interface {
	tl.TLObject
	_MessagesMessageReactionsList()
}

var (
	_ MessagesMessageReactionsList = (*MessagesMessageReactionsListPredict)(nil)
)

type MessagesMessageReactionsListPredict struct {
	_          struct{}              `tl:"flags,bitflag"`
	Count      int32                 `tl:"count"`
	Reactions  []MessagePeerReaction `tl:"reactions"`
	Chats      []Chat                `tl:"chats"`
	Users      []User                `tl:"users"`
	NextOffset *string               `tl:"next_offset,omitempty:flags:0"`
}

func (*MessagesMessageReactionsListPredict) CRC() uint32 {
	return 0x31bd492d
}
func (*MessagesMessageReactionsListPredict) _MessagesMessageReactionsList() {}

type MessagesMessageViews interface {
	tl.TLObject
	_MessagesMessageViews()
}

var (
	_ MessagesMessageViews = (*MessagesMessageViewsPredict)(nil)
)

type MessagesMessageViewsPredict struct {
	Views []MessageViews `tl:"views"`
	Chats []Chat         `tl:"chats"`
	Users []User         `tl:"users"`
}

func (*MessagesMessageViewsPredict) CRC() uint32 {
	return 0xb6c4f543
}
func (*MessagesMessageViewsPredict) _MessagesMessageViews() {}

type MessagesMessages interface {
	tl.TLObject
	_MessagesMessages()
}

var (
	_ MessagesMessages = (*MessagesMessagesPredict)(nil)
	_ MessagesMessages = (*MessagesMessagesSlicePredict)(nil)
	_ MessagesMessages = (*MessagesChannelMessagesPredict)(nil)
	_ MessagesMessages = (*MessagesMessagesNotModifiedPredict)(nil)
)

type MessagesMessagesPredict struct {
	Messages []Message `tl:"messages"`
	Chats    []Chat    `tl:"chats"`
	Users    []User    `tl:"users"`
}

func (*MessagesMessagesPredict) CRC() uint32 {
	return 0x8c718e87
}
func (*MessagesMessagesPredict) _MessagesMessages() {}

type MessagesMessagesSlicePredict struct {
	_              struct{}  `tl:"flags,bitflag"`
	Inexact        bool      `tl:"inexact,omitempty:flags:1,implicit"`
	Count          int32     `tl:"count"`
	NextRate       *int32    `tl:"next_rate,omitempty:flags:0"`
	OffsetIDOffset *int32    `tl:"offset_id_offset,omitempty:flags:2"`
	Messages       []Message `tl:"messages"`
	Chats          []Chat    `tl:"chats"`
	Users          []User    `tl:"users"`
}

func (*MessagesMessagesSlicePredict) CRC() uint32 {
	return 0x3a54685e
}
func (*MessagesMessagesSlicePredict) _MessagesMessages() {}

type MessagesChannelMessagesPredict struct {
	_              struct{}     `tl:"flags,bitflag"`
	Inexact        bool         `tl:"inexact,omitempty:flags:1,implicit"`
	Pts            int32        `tl:"pts"`
	Count          int32        `tl:"count"`
	OffsetIDOffset *int32       `tl:"offset_id_offset,omitempty:flags:2"`
	Messages       []Message    `tl:"messages"`
	Topics         []ForumTopic `tl:"topics"`
	Chats          []Chat       `tl:"chats"`
	Users          []User       `tl:"users"`
}

func (*MessagesChannelMessagesPredict) CRC() uint32 {
	return 0xc776ba4e
}
func (*MessagesChannelMessagesPredict) _MessagesMessages() {}

type MessagesMessagesNotModifiedPredict struct {
	Count int32 `tl:"count"`
}

func (*MessagesMessagesNotModifiedPredict) CRC() uint32 {
	return 0x74535f21
}
func (*MessagesMessagesNotModifiedPredict) _MessagesMessages() {}

type MessagesMyStickers interface {
	tl.TLObject
	_MessagesMyStickers()
}

var (
	_ MessagesMyStickers = (*MessagesMyStickersPredict)(nil)
)

type MessagesMyStickersPredict struct {
	Count int32               `tl:"count"`
	Sets  []StickerSetCovered `tl:"sets"`
}

func (*MessagesMyStickersPredict) CRC() uint32 {
	return 0xfaff629d
}
func (*MessagesMyStickersPredict) _MessagesMyStickers() {}

type MessagesPeerDialogs interface {
	tl.TLObject
	_MessagesPeerDialogs()
}

var (
	_ MessagesPeerDialogs = (*MessagesPeerDialogsPredict)(nil)
)

type MessagesPeerDialogsPredict struct {
	Dialogs  []Dialog     `tl:"dialogs"`
	Messages []Message    `tl:"messages"`
	Chats    []Chat       `tl:"chats"`
	Users    []User       `tl:"users"`
	State    UpdatesState `tl:"state"`
}

func (*MessagesPeerDialogsPredict) CRC() uint32 {
	return 0x3371c354
}
func (*MessagesPeerDialogsPredict) _MessagesPeerDialogs() {}

type MessagesPeerSettings interface {
	tl.TLObject
	_MessagesPeerSettings()
}

var (
	_ MessagesPeerSettings = (*MessagesPeerSettingsPredict)(nil)
)

type MessagesPeerSettingsPredict struct {
	Settings PeerSettings `tl:"settings"`
	Chats    []Chat       `tl:"chats"`
	Users    []User       `tl:"users"`
}

func (*MessagesPeerSettingsPredict) CRC() uint32 {
	return 0x6880b94d
}
func (*MessagesPeerSettingsPredict) _MessagesPeerSettings() {}

type MessagesQuickReplies interface {
	tl.TLObject
	_MessagesQuickReplies()
}

var (
	_ MessagesQuickReplies = (*MessagesQuickRepliesPredict)(nil)
	_ MessagesQuickReplies = (*MessagesQuickRepliesNotModifiedPredict)(nil)
)

type MessagesQuickRepliesPredict struct {
	QuickReplies []QuickReply `tl:"quick_replies"`
	Messages     []Message    `tl:"messages"`
	Chats        []Chat       `tl:"chats"`
	Users        []User       `tl:"users"`
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
	tl.TLObject
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
	Hash      int64      `tl:"hash"`
	Reactions []Reaction `tl:"reactions"`
}

func (*MessagesReactionsPredict) CRC() uint32 {
	return 0xeafdf716
}
func (*MessagesReactionsPredict) _MessagesReactions() {}

type MessagesRecentStickers interface {
	tl.TLObject
	_MessagesRecentStickers()
}

var (
	_ MessagesRecentStickers = (*MessagesRecentStickersNotModifiedPredict)(nil)
	_ MessagesRecentStickers = (*MessagesRecentStickersPredict)(nil)
)

type MessagesRecentStickersNotModifiedPredict struct{}

func (*MessagesRecentStickersNotModifiedPredict) CRC() uint32 {
	return 0x0b17f890
}
func (*MessagesRecentStickersNotModifiedPredict) _MessagesRecentStickers() {}

type MessagesRecentStickersPredict struct {
	Hash     int64         `tl:"hash"`
	Packs    []StickerPack `tl:"packs"`
	Stickers []Document    `tl:"stickers"`
	Dates    []int32       `tl:"dates"`
}

func (*MessagesRecentStickersPredict) CRC() uint32 {
	return 0x88d37c56
}
func (*MessagesRecentStickersPredict) _MessagesRecentStickers() {}

type MessagesSavedDialogs interface {
	tl.TLObject
	_MessagesSavedDialogs()
}

var (
	_ MessagesSavedDialogs = (*MessagesSavedDialogsPredict)(nil)
	_ MessagesSavedDialogs = (*MessagesSavedDialogsSlicePredict)(nil)
	_ MessagesSavedDialogs = (*MessagesSavedDialogsNotModifiedPredict)(nil)
)

type MessagesSavedDialogsPredict struct {
	Dialogs  []SavedDialog `tl:"dialogs"`
	Messages []Message     `tl:"messages"`
	Chats    []Chat        `tl:"chats"`
	Users    []User        `tl:"users"`
}

func (*MessagesSavedDialogsPredict) CRC() uint32 {
	return 0xf83ae221
}
func (*MessagesSavedDialogsPredict) _MessagesSavedDialogs() {}

type MessagesSavedDialogsSlicePredict struct {
	Count    int32         `tl:"count"`
	Dialogs  []SavedDialog `tl:"dialogs"`
	Messages []Message     `tl:"messages"`
	Chats    []Chat        `tl:"chats"`
	Users    []User        `tl:"users"`
}

func (*MessagesSavedDialogsSlicePredict) CRC() uint32 {
	return 0x44ba9dd9
}
func (*MessagesSavedDialogsSlicePredict) _MessagesSavedDialogs() {}

type MessagesSavedDialogsNotModifiedPredict struct {
	Count int32 `tl:"count"`
}

func (*MessagesSavedDialogsNotModifiedPredict) CRC() uint32 {
	return 0xc01f6fe8
}
func (*MessagesSavedDialogsNotModifiedPredict) _MessagesSavedDialogs() {}

type MessagesSavedGifs interface {
	tl.TLObject
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
	Hash int64      `tl:"hash"`
	Gifs []Document `tl:"gifs"`
}

func (*MessagesSavedGifsPredict) CRC() uint32 {
	return 0x84a02a0d
}
func (*MessagesSavedGifsPredict) _MessagesSavedGifs() {}

type MessagesSavedReactionTags interface {
	tl.TLObject
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
	Tags []SavedReactionTag `tl:"tags"`
	Hash int64              `tl:"hash"`
}

func (*MessagesSavedReactionTagsPredict) CRC() uint32 {
	return 0x3259950a
}
func (*MessagesSavedReactionTagsPredict) _MessagesSavedReactionTags() {}

type MessagesSearchCounter interface {
	tl.TLObject
	_MessagesSearchCounter()
}

var (
	_ MessagesSearchCounter = (*MessagesSearchCounterPredict)(nil)
)

type MessagesSearchCounterPredict struct {
	_       struct{}       `tl:"flags,bitflag"`
	Inexact bool           `tl:"inexact,omitempty:flags:1,implicit"`
	Filter  MessagesFilter `tl:"filter"`
	Count   int32          `tl:"count"`
}

func (*MessagesSearchCounterPredict) CRC() uint32 {
	return 0xe844ebff
}
func (*MessagesSearchCounterPredict) _MessagesSearchCounter() {}

type MessagesSearchResultsCalendar interface {
	tl.TLObject
	_MessagesSearchResultsCalendar()
}

var (
	_ MessagesSearchResultsCalendar = (*MessagesSearchResultsCalendarPredict)(nil)
)

type MessagesSearchResultsCalendarPredict struct {
	_              struct{}                      `tl:"flags,bitflag"`
	Inexact        bool                          `tl:"inexact,omitempty:flags:0,implicit"`
	Count          int32                         `tl:"count"`
	MinDate        int32                         `tl:"min_date"`
	MinMsgID       int32                         `tl:"min_msg_id"`
	OffsetIDOffset *int32                        `tl:"offset_id_offset,omitempty:flags:1"`
	Periods        []SearchResultsCalendarPeriod `tl:"periods"`
	Messages       []Message                     `tl:"messages"`
	Chats          []Chat                        `tl:"chats"`
	Users          []User                        `tl:"users"`
}

func (*MessagesSearchResultsCalendarPredict) CRC() uint32 {
	return 0x147ee23c
}
func (*MessagesSearchResultsCalendarPredict) _MessagesSearchResultsCalendar() {}

type MessagesSearchResultsPositions interface {
	tl.TLObject
	_MessagesSearchResultsPositions()
}

var (
	_ MessagesSearchResultsPositions = (*MessagesSearchResultsPositionsPredict)(nil)
)

type MessagesSearchResultsPositionsPredict struct {
	Count     int32                   `tl:"count"`
	Positions []SearchResultsPosition `tl:"positions"`
}

func (*MessagesSearchResultsPositionsPredict) CRC() uint32 {
	return 0x53b22baf
}
func (*MessagesSearchResultsPositionsPredict) _MessagesSearchResultsPositions() {}

type MessagesSentEncryptedMessage interface {
	tl.TLObject
	_MessagesSentEncryptedMessage()
}

var (
	_ MessagesSentEncryptedMessage = (*MessagesSentEncryptedMessagePredict)(nil)
	_ MessagesSentEncryptedMessage = (*MessagesSentEncryptedFilePredict)(nil)
)

type MessagesSentEncryptedMessagePredict struct {
	Date int32 `tl:"date"`
}

func (*MessagesSentEncryptedMessagePredict) CRC() uint32 {
	return 0x560f8935
}
func (*MessagesSentEncryptedMessagePredict) _MessagesSentEncryptedMessage() {}

type MessagesSentEncryptedFilePredict struct {
	Date int32         `tl:"date"`
	File EncryptedFile `tl:"file"`
}

func (*MessagesSentEncryptedFilePredict) CRC() uint32 {
	return 0x9493ff32
}
func (*MessagesSentEncryptedFilePredict) _MessagesSentEncryptedMessage() {}

type MessagesSponsoredMessages interface {
	tl.TLObject
	_MessagesSponsoredMessages()
}

var (
	_ MessagesSponsoredMessages = (*MessagesSponsoredMessagesPredict)(nil)
	_ MessagesSponsoredMessages = (*MessagesSponsoredMessagesEmptyPredict)(nil)
)

type MessagesSponsoredMessagesPredict struct {
	_            struct{}           `tl:"flags,bitflag"`
	PostsBetween *int32             `tl:"posts_between,omitempty:flags:0"`
	Messages     []SponsoredMessage `tl:"messages"`
	Chats        []Chat             `tl:"chats"`
	Users        []User             `tl:"users"`
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
	tl.TLObject
	_MessagesStickerSet()
}

var (
	_ MessagesStickerSet = (*MessagesStickerSetPredict)(nil)
	_ MessagesStickerSet = (*MessagesStickerSetNotModifiedPredict)(nil)
)

type MessagesStickerSetPredict struct {
	Set       StickerSet       `tl:"set"`
	Packs     []StickerPack    `tl:"packs"`
	Keywords  []StickerKeyword `tl:"keywords"`
	Documents []Document       `tl:"documents"`
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
	tl.TLObject
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
	Sets []StickerSetCovered `tl:"sets"`
}

func (*MessagesStickerSetInstallResultArchivePredict) CRC() uint32 {
	return 0x35e410a8
}
func (*MessagesStickerSetInstallResultArchivePredict) _MessagesStickerSetInstallResult() {}

type MessagesStickers interface {
	tl.TLObject
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
	Hash     int64      `tl:"hash"`
	Stickers []Document `tl:"stickers"`
}

func (*MessagesStickersPredict) CRC() uint32 {
	return 0x30a6ec7e
}
func (*MessagesStickersPredict) _MessagesStickers() {}

type MessagesTranscribedAudio interface {
	tl.TLObject
	_MessagesTranscribedAudio()
}

var (
	_ MessagesTranscribedAudio = (*MessagesTranscribedAudioPredict)(nil)
)

type MessagesTranscribedAudioPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	Pending               bool     `tl:"pending,omitempty:flags:0,implicit"`
	TranscriptionID       int64    `tl:"transcription_id"`
	Text                  string   `tl:"text"`
	TrialRemainsNum       *int32   `tl:"trial_remains_num,omitempty:flags:1"`
	TrialRemainsUntilDate *int32   `tl:"trial_remains_until_date,omitempty:flags:1"`
}

func (*MessagesTranscribedAudioPredict) CRC() uint32 {
	return 0xcfb9d957
}
func (*MessagesTranscribedAudioPredict) _MessagesTranscribedAudio() {}

type MessagesTranslatedText interface {
	tl.TLObject
	_MessagesTranslatedText()
}

var (
	_ MessagesTranslatedText = (*MessagesTranslateResultPredict)(nil)
)

type MessagesTranslateResultPredict struct {
	Result []TextWithEntities `tl:"result"`
}

func (*MessagesTranslateResultPredict) CRC() uint32 {
	return 0x33db32f8
}
func (*MessagesTranslateResultPredict) _MessagesTranslatedText() {}

type MessagesVotesList interface {
	tl.TLObject
	_MessagesVotesList()
}

var (
	_ MessagesVotesList = (*MessagesVotesListPredict)(nil)
)

type MessagesVotesListPredict struct {
	_          struct{}          `tl:"flags,bitflag"`
	Count      int32             `tl:"count"`
	Votes      []MessagePeerVote `tl:"votes"`
	Chats      []Chat            `tl:"chats"`
	Users      []User            `tl:"users"`
	NextOffset *string           `tl:"next_offset,omitempty:flags:0"`
}

func (*MessagesVotesListPredict) CRC() uint32 {
	return 0x4899484e
}
func (*MessagesVotesListPredict) _MessagesVotesList() {}

type MessagesWebPage interface {
	tl.TLObject
	_MessagesWebPage()
}

var (
	_ MessagesWebPage = (*MessagesWebPagePredict)(nil)
)

type MessagesWebPagePredict struct {
	Webpage WebPage `tl:"webpage"`
	Chats   []Chat  `tl:"chats"`
	Users   []User  `tl:"users"`
}

func (*MessagesWebPagePredict) CRC() uint32 {
	return 0xfd5e12bd
}
func (*MessagesWebPagePredict) _MessagesWebPage() {}

type PaymentsBankCardData interface {
	tl.TLObject
	_PaymentsBankCardData()
}

var (
	_ PaymentsBankCardData = (*PaymentsBankCardDataPredict)(nil)
)

type PaymentsBankCardDataPredict struct {
	Title    string            `tl:"title"`
	OpenUrls []BankCardOpenURL `tl:"open_urls"`
}

func (*PaymentsBankCardDataPredict) CRC() uint32 {
	return 0x3e24e573
}
func (*PaymentsBankCardDataPredict) _PaymentsBankCardData() {}

type PaymentsCheckedGiftCode interface {
	tl.TLObject
	_PaymentsCheckedGiftCode()
}

var (
	_ PaymentsCheckedGiftCode = (*PaymentsCheckedGiftCodePredict)(nil)
)

type PaymentsCheckedGiftCodePredict struct {
	_             struct{} `tl:"flags,bitflag"`
	ViaGiveaway   bool     `tl:"via_giveaway,omitempty:flags:2,implicit"`
	FromID        Peer     `tl:"from_id,omitempty:flags:4"`
	GiveawayMsgID *int32   `tl:"giveaway_msg_id,omitempty:flags:3"`
	ToID          *int64   `tl:"to_id,omitempty:flags:0"`
	Date          int32    `tl:"date"`
	Months        int32    `tl:"months"`
	UsedDate      *int32   `tl:"used_date,omitempty:flags:1"`
	Chats         []Chat   `tl:"chats"`
	Users         []User   `tl:"users"`
}

func (*PaymentsCheckedGiftCodePredict) CRC() uint32 {
	return 0x284a1096
}
func (*PaymentsCheckedGiftCodePredict) _PaymentsCheckedGiftCode() {}

type PaymentsExportedInvoice interface {
	tl.TLObject
	_PaymentsExportedInvoice()
}

var (
	_ PaymentsExportedInvoice = (*PaymentsExportedInvoicePredict)(nil)
)

type PaymentsExportedInvoicePredict struct {
	URL string `tl:"url"`
}

func (*PaymentsExportedInvoicePredict) CRC() uint32 {
	return 0xaed0cbd9
}
func (*PaymentsExportedInvoicePredict) _PaymentsExportedInvoice() {}

type PaymentsGiveawayInfo interface {
	tl.TLObject
	_PaymentsGiveawayInfo()
}

var (
	_ PaymentsGiveawayInfo = (*PaymentsGiveawayInfoPredict)(nil)
	_ PaymentsGiveawayInfo = (*PaymentsGiveawayInfoResultsPredict)(nil)
)

type PaymentsGiveawayInfoPredict struct {
	_                     struct{} `tl:"flags,bitflag"`
	Participating         bool     `tl:"participating,omitempty:flags:0,implicit"`
	PreparingResults      bool     `tl:"preparing_results,omitempty:flags:3,implicit"`
	StartDate             int32    `tl:"start_date"`
	JoinedTooEarlyDate    *int32   `tl:"joined_too_early_date,omitempty:flags:1"`
	AdminDisallowedChatID *int64   `tl:"admin_disallowed_chat_id,omitempty:flags:2"`
	DisallowedCountry     *string  `tl:"disallowed_country,omitempty:flags:4"`
}

func (*PaymentsGiveawayInfoPredict) CRC() uint32 {
	return 0x4367daa0
}
func (*PaymentsGiveawayInfoPredict) _PaymentsGiveawayInfo() {}

type PaymentsGiveawayInfoResultsPredict struct {
	_              struct{} `tl:"flags,bitflag"`
	Winner         bool     `tl:"winner,omitempty:flags:0,implicit"`
	Refunded       bool     `tl:"refunded,omitempty:flags:1,implicit"`
	StartDate      int32    `tl:"start_date"`
	GiftCodeSlug   *string  `tl:"gift_code_slug,omitempty:flags:0"`
	FinishDate     int32    `tl:"finish_date"`
	WinnersCount   int32    `tl:"winners_count"`
	ActivatedCount int32    `tl:"activated_count"`
}

func (*PaymentsGiveawayInfoResultsPredict) CRC() uint32 {
	return 0x00cd5570
}
func (*PaymentsGiveawayInfoResultsPredict) _PaymentsGiveawayInfo() {}

type PaymentsPaymentForm interface {
	tl.TLObject
	_PaymentsPaymentForm()
}

var (
	_ PaymentsPaymentForm = (*PaymentsPaymentFormPredict)(nil)
	_ PaymentsPaymentForm = (*PaymentsPaymentFormStarsPredict)(nil)
)

type PaymentsPaymentFormPredict struct {
	_                  struct{}                  `tl:"flags,bitflag"`
	CanSaveCredentials bool                      `tl:"can_save_credentials,omitempty:flags:2,implicit"`
	PasswordMissing    bool                      `tl:"password_missing,omitempty:flags:3,implicit"`
	FormID             int64                     `tl:"form_id"`
	BotID              int64                     `tl:"bot_id"`
	Title              string                    `tl:"title"`
	Description        string                    `tl:"description"`
	Photo              WebDocument               `tl:"photo,omitempty:flags:5"`
	Invoice            Invoice                   `tl:"invoice"`
	ProviderID         int64                     `tl:"provider_id"`
	URL                string                    `tl:"url"`
	NativeProvider     *string                   `tl:"native_provider,omitempty:flags:4"`
	NativeParams       DataJSON                  `tl:"native_params,omitempty:flags:4"`
	AdditionalMethods  []PaymentFormMethod       `tl:"additional_methods,omitempty:flags:6"`
	SavedInfo          PaymentRequestedInfo      `tl:"saved_info,omitempty:flags:0"`
	SavedCredentials   []PaymentSavedCredentials `tl:"saved_credentials,omitempty:flags:1"`
	Users              []User                    `tl:"users"`
}

func (*PaymentsPaymentFormPredict) CRC() uint32 {
	return 0xa0058751
}
func (*PaymentsPaymentFormPredict) _PaymentsPaymentForm() {}

type PaymentsPaymentFormStarsPredict struct {
	_           struct{}    `tl:"flags,bitflag"`
	FormID      int64       `tl:"form_id"`
	BotID       int64       `tl:"bot_id"`
	Title       string      `tl:"title"`
	Description string      `tl:"description"`
	Photo       WebDocument `tl:"photo,omitempty:flags:5"`
	Invoice     Invoice     `tl:"invoice"`
	Users       []User      `tl:"users"`
}

func (*PaymentsPaymentFormStarsPredict) CRC() uint32 {
	return 0x7bf6b15c
}
func (*PaymentsPaymentFormStarsPredict) _PaymentsPaymentForm() {}

type PaymentsPaymentReceipt interface {
	tl.TLObject
	_PaymentsPaymentReceipt()
}

var (
	_ PaymentsPaymentReceipt = (*PaymentsPaymentReceiptPredict)(nil)
	_ PaymentsPaymentReceipt = (*PaymentsPaymentReceiptStarsPredict)(nil)
)

type PaymentsPaymentReceiptPredict struct {
	_                struct{}             `tl:"flags,bitflag"`
	Date             int32                `tl:"date"`
	BotID            int64                `tl:"bot_id"`
	ProviderID       int64                `tl:"provider_id"`
	Title            string               `tl:"title"`
	Description      string               `tl:"description"`
	Photo            WebDocument          `tl:"photo,omitempty:flags:2"`
	Invoice          Invoice              `tl:"invoice"`
	Info             PaymentRequestedInfo `tl:"info,omitempty:flags:0"`
	Shipping         ShippingOption       `tl:"shipping,omitempty:flags:1"`
	TipAmount        *int64               `tl:"tip_amount,omitempty:flags:3"`
	Currency         string               `tl:"currency"`
	TotalAmount      int64                `tl:"total_amount"`
	CredentialsTitle string               `tl:"credentials_title"`
	Users            []User               `tl:"users"`
}

func (*PaymentsPaymentReceiptPredict) CRC() uint32 {
	return 0x70c4fe03
}
func (*PaymentsPaymentReceiptPredict) _PaymentsPaymentReceipt() {}

type PaymentsPaymentReceiptStarsPredict struct {
	_             struct{}    `tl:"flags,bitflag"`
	Date          int32       `tl:"date"`
	BotID         int64       `tl:"bot_id"`
	Title         string      `tl:"title"`
	Description   string      `tl:"description"`
	Photo         WebDocument `tl:"photo,omitempty:flags:2"`
	Invoice       Invoice     `tl:"invoice"`
	Currency      string      `tl:"currency"`
	TotalAmount   int64       `tl:"total_amount"`
	TransactionID string      `tl:"transaction_id"`
	Users         []User      `tl:"users"`
}

func (*PaymentsPaymentReceiptStarsPredict) CRC() uint32 {
	return 0xdabbf83a
}
func (*PaymentsPaymentReceiptStarsPredict) _PaymentsPaymentReceipt() {}

type PaymentsPaymentResult interface {
	tl.TLObject
	_PaymentsPaymentResult()
}

var (
	_ PaymentsPaymentResult = (*PaymentsPaymentResultPredict)(nil)
	_ PaymentsPaymentResult = (*PaymentsPaymentVerificationNeededPredict)(nil)
)

type PaymentsPaymentResultPredict struct {
	Updates Updates `tl:"updates"`
}

func (*PaymentsPaymentResultPredict) CRC() uint32 {
	return 0x4e5f810d
}
func (*PaymentsPaymentResultPredict) _PaymentsPaymentResult() {}

type PaymentsPaymentVerificationNeededPredict struct {
	URL string `tl:"url"`
}

func (*PaymentsPaymentVerificationNeededPredict) CRC() uint32 {
	return 0xd8411139
}
func (*PaymentsPaymentVerificationNeededPredict) _PaymentsPaymentResult() {}

type PaymentsSavedInfo interface {
	tl.TLObject
	_PaymentsSavedInfo()
}

var (
	_ PaymentsSavedInfo = (*PaymentsSavedInfoPredict)(nil)
)

type PaymentsSavedInfoPredict struct {
	_                   struct{}             `tl:"flags,bitflag"`
	HasSavedCredentials bool                 `tl:"has_saved_credentials,omitempty:flags:1,implicit"`
	SavedInfo           PaymentRequestedInfo `tl:"saved_info,omitempty:flags:0"`
}

func (*PaymentsSavedInfoPredict) CRC() uint32 {
	return 0xfb8fe43c
}
func (*PaymentsSavedInfoPredict) _PaymentsSavedInfo() {}

type PaymentsStarsRevenueAdsAccountURL interface {
	tl.TLObject
	_PaymentsStarsRevenueAdsAccountURL()
}

var (
	_ PaymentsStarsRevenueAdsAccountURL = (*PaymentsStarsRevenueAdsAccountURLPredict)(nil)
)

type PaymentsStarsRevenueAdsAccountURLPredict struct {
	URL string `tl:"url"`
}

func (*PaymentsStarsRevenueAdsAccountURLPredict) CRC() uint32 {
	return 0x394e7f21
}
func (*PaymentsStarsRevenueAdsAccountURLPredict) _PaymentsStarsRevenueAdsAccountURL() {}

type PaymentsStarsRevenueStats interface {
	tl.TLObject
	_PaymentsStarsRevenueStats()
}

var (
	_ PaymentsStarsRevenueStats = (*PaymentsStarsRevenueStatsPredict)(nil)
)

type PaymentsStarsRevenueStatsPredict struct {
	RevenueGraph StatsGraph         `tl:"revenue_graph"`
	Status       StarsRevenueStatus `tl:"status"`
	UsdRate      float64            `tl:"usd_rate"`
}

func (*PaymentsStarsRevenueStatsPredict) CRC() uint32 {
	return 0xc92bb73b
}
func (*PaymentsStarsRevenueStatsPredict) _PaymentsStarsRevenueStats() {}

type PaymentsStarsRevenueWithdrawalURL interface {
	tl.TLObject
	_PaymentsStarsRevenueWithdrawalURL()
}

var (
	_ PaymentsStarsRevenueWithdrawalURL = (*PaymentsStarsRevenueWithdrawalURLPredict)(nil)
)

type PaymentsStarsRevenueWithdrawalURLPredict struct {
	URL string `tl:"url"`
}

func (*PaymentsStarsRevenueWithdrawalURLPredict) CRC() uint32 {
	return 0x1dab80b7
}
func (*PaymentsStarsRevenueWithdrawalURLPredict) _PaymentsStarsRevenueWithdrawalURL() {}

type PaymentsStarsStatus interface {
	tl.TLObject
	_PaymentsStarsStatus()
}

var (
	_ PaymentsStarsStatus = (*PaymentsStarsStatusPredict)(nil)
)

type PaymentsStarsStatusPredict struct {
	_          struct{}           `tl:"flags,bitflag"`
	Balance    int64              `tl:"balance"`
	History    []StarsTransaction `tl:"history"`
	NextOffset *string            `tl:"next_offset,omitempty:flags:0"`
	Chats      []Chat             `tl:"chats"`
	Users      []User             `tl:"users"`
}

func (*PaymentsStarsStatusPredict) CRC() uint32 {
	return 0x8cf4ee60
}
func (*PaymentsStarsStatusPredict) _PaymentsStarsStatus() {}

type PaymentsValidatedRequestedInfo interface {
	tl.TLObject
	_PaymentsValidatedRequestedInfo()
}

var (
	_ PaymentsValidatedRequestedInfo = (*PaymentsValidatedRequestedInfoPredict)(nil)
)

type PaymentsValidatedRequestedInfoPredict struct {
	_               struct{}         `tl:"flags,bitflag"`
	ID              *string          `tl:"id,omitempty:flags:0"`
	ShippingOptions []ShippingOption `tl:"shipping_options,omitempty:flags:1"`
}

func (*PaymentsValidatedRequestedInfoPredict) CRC() uint32 {
	return 0xd1451883
}
func (*PaymentsValidatedRequestedInfoPredict) _PaymentsValidatedRequestedInfo() {}

type PhoneExportedGroupCallInvite interface {
	tl.TLObject
	_PhoneExportedGroupCallInvite()
}

var (
	_ PhoneExportedGroupCallInvite = (*PhoneExportedGroupCallInvitePredict)(nil)
)

type PhoneExportedGroupCallInvitePredict struct {
	Link string `tl:"link"`
}

func (*PhoneExportedGroupCallInvitePredict) CRC() uint32 {
	return 0x204bd158
}
func (*PhoneExportedGroupCallInvitePredict) _PhoneExportedGroupCallInvite() {}

type PhoneGroupCall interface {
	tl.TLObject
	_PhoneGroupCall()
}

var (
	_ PhoneGroupCall = (*PhoneGroupCallPredict)(nil)
)

type PhoneGroupCallPredict struct {
	Call                   GroupCall              `tl:"call"`
	Participants           []GroupCallParticipant `tl:"participants"`
	ParticipantsNextOffset string                 `tl:"participants_next_offset"`
	Chats                  []Chat                 `tl:"chats"`
	Users                  []User                 `tl:"users"`
}

func (*PhoneGroupCallPredict) CRC() uint32 {
	return 0x9e727aad
}
func (*PhoneGroupCallPredict) _PhoneGroupCall() {}

type PhoneGroupCallStreamChannels interface {
	tl.TLObject
	_PhoneGroupCallStreamChannels()
}

var (
	_ PhoneGroupCallStreamChannels = (*PhoneGroupCallStreamChannelsPredict)(nil)
)

type PhoneGroupCallStreamChannelsPredict struct {
	Channels []GroupCallStreamChannel `tl:"channels"`
}

func (*PhoneGroupCallStreamChannelsPredict) CRC() uint32 {
	return 0xd0e482b2
}
func (*PhoneGroupCallStreamChannelsPredict) _PhoneGroupCallStreamChannels() {}

type PhoneGroupCallStreamRtmpURL interface {
	tl.TLObject
	_PhoneGroupCallStreamRtmpURL()
}

var (
	_ PhoneGroupCallStreamRtmpURL = (*PhoneGroupCallStreamRtmpURLPredict)(nil)
)

type PhoneGroupCallStreamRtmpURLPredict struct {
	URL string `tl:"url"`
	Key string `tl:"key"`
}

func (*PhoneGroupCallStreamRtmpURLPredict) CRC() uint32 {
	return 0x2dbf3432
}
func (*PhoneGroupCallStreamRtmpURLPredict) _PhoneGroupCallStreamRtmpURL() {}

type PhoneGroupParticipants interface {
	tl.TLObject
	_PhoneGroupParticipants()
}

var (
	_ PhoneGroupParticipants = (*PhoneGroupParticipantsPredict)(nil)
)

type PhoneGroupParticipantsPredict struct {
	Count        int32                  `tl:"count"`
	Participants []GroupCallParticipant `tl:"participants"`
	NextOffset   string                 `tl:"next_offset"`
	Chats        []Chat                 `tl:"chats"`
	Users        []User                 `tl:"users"`
	Version      int32                  `tl:"version"`
}

func (*PhoneGroupParticipantsPredict) CRC() uint32 {
	return 0xf47751b6
}
func (*PhoneGroupParticipantsPredict) _PhoneGroupParticipants() {}

type PhoneJoinAsPeers interface {
	tl.TLObject
	_PhoneJoinAsPeers()
}

var (
	_ PhoneJoinAsPeers = (*PhoneJoinAsPeersPredict)(nil)
)

type PhoneJoinAsPeersPredict struct {
	Peers []Peer `tl:"peers"`
	Chats []Chat `tl:"chats"`
	Users []User `tl:"users"`
}

func (*PhoneJoinAsPeersPredict) CRC() uint32 {
	return 0xafe5623f
}
func (*PhoneJoinAsPeersPredict) _PhoneJoinAsPeers() {}

type PhonePhoneCall interface {
	tl.TLObject
	_PhonePhoneCall()
}

var (
	_ PhonePhoneCall = (*PhonePhoneCallPredict)(nil)
)

type PhonePhoneCallPredict struct {
	PhoneCall PhoneCall `tl:"phone_call"`
	Users     []User    `tl:"users"`
}

func (*PhonePhoneCallPredict) CRC() uint32 {
	return 0xec82e140
}
func (*PhonePhoneCallPredict) _PhonePhoneCall() {}

type PhotosPhoto interface {
	tl.TLObject
	_PhotosPhoto()
}

var (
	_ PhotosPhoto = (*PhotosPhotoPredict)(nil)
)

type PhotosPhotoPredict struct {
	Photo Photo  `tl:"photo"`
	Users []User `tl:"users"`
}

func (*PhotosPhotoPredict) CRC() uint32 {
	return 0x20212ca8
}
func (*PhotosPhotoPredict) _PhotosPhoto() {}

type PhotosPhotos interface {
	tl.TLObject
	_PhotosPhotos()
}

var (
	_ PhotosPhotos = (*PhotosPhotosPredict)(nil)
	_ PhotosPhotos = (*PhotosPhotosSlicePredict)(nil)
)

type PhotosPhotosPredict struct {
	Photos []Photo `tl:"photos"`
	Users  []User  `tl:"users"`
}

func (*PhotosPhotosPredict) CRC() uint32 {
	return 0x8dca6aa5
}
func (*PhotosPhotosPredict) _PhotosPhotos() {}

type PhotosPhotosSlicePredict struct {
	Count  int32   `tl:"count"`
	Photos []Photo `tl:"photos"`
	Users  []User  `tl:"users"`
}

func (*PhotosPhotosSlicePredict) CRC() uint32 {
	return 0x15051f54
}
func (*PhotosPhotosSlicePredict) _PhotosPhotos() {}

type PremiumBoostsList interface {
	tl.TLObject
	_PremiumBoostsList()
}

var (
	_ PremiumBoostsList = (*PremiumBoostsListPredict)(nil)
)

type PremiumBoostsListPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	Count      int32    `tl:"count"`
	Boosts     []Boost  `tl:"boosts"`
	NextOffset *string  `tl:"next_offset,omitempty:flags:0"`
	Users      []User   `tl:"users"`
}

func (*PremiumBoostsListPredict) CRC() uint32 {
	return 0x86f8613c
}
func (*PremiumBoostsListPredict) _PremiumBoostsList() {}

type PremiumBoostsStatus interface {
	tl.TLObject
	_PremiumBoostsStatus()
}

var (
	_ PremiumBoostsStatus = (*PremiumBoostsStatusPredict)(nil)
)

type PremiumBoostsStatusPredict struct {
	_                  struct{}          `tl:"flags,bitflag"`
	MyBoost            bool              `tl:"my_boost,omitempty:flags:2,implicit"`
	Level              int32             `tl:"level"`
	CurrentLevelBoosts int32             `tl:"current_level_boosts"`
	Boosts             int32             `tl:"boosts"`
	GiftBoosts         *int32            `tl:"gift_boosts,omitempty:flags:4"`
	NextLevelBoosts    *int32            `tl:"next_level_boosts,omitempty:flags:0"`
	PremiumAudience    StatsPercentValue `tl:"premium_audience,omitempty:flags:1"`
	BoostURL           string            `tl:"boost_url"`
	PrepaidGiveaways   []PrepaidGiveaway `tl:"prepaid_giveaways,omitempty:flags:3"`
	MyBoostSlots       []int32           `tl:"my_boost_slots,omitempty:flags:2"`
}

func (*PremiumBoostsStatusPredict) CRC() uint32 {
	return 0x4959427a
}
func (*PremiumBoostsStatusPredict) _PremiumBoostsStatus() {}

type PremiumMyBoosts interface {
	tl.TLObject
	_PremiumMyBoosts()
}

var (
	_ PremiumMyBoosts = (*PremiumMyBoostsPredict)(nil)
)

type PremiumMyBoostsPredict struct {
	MyBoosts []MyBoost `tl:"my_boosts"`
	Chats    []Chat    `tl:"chats"`
	Users    []User    `tl:"users"`
}

func (*PremiumMyBoostsPredict) CRC() uint32 {
	return 0x9ae228e2
}
func (*PremiumMyBoostsPredict) _PremiumMyBoosts() {}

type SmsjobsEligibilityToJoin interface {
	tl.TLObject
	_SmsjobsEligibilityToJoin()
}

var (
	_ SmsjobsEligibilityToJoin = (*SmsjobsEligibleToJoinPredict)(nil)
)

type SmsjobsEligibleToJoinPredict struct {
	TermsURL       string `tl:"terms_url"`
	MonthlySentSms int32  `tl:"monthly_sent_sms"`
}

func (*SmsjobsEligibleToJoinPredict) CRC() uint32 {
	return 0xdc8b44cf
}
func (*SmsjobsEligibleToJoinPredict) _SmsjobsEligibilityToJoin() {}

type SmsjobsStatus interface {
	tl.TLObject
	_SmsjobsStatus()
}

var (
	_ SmsjobsStatus = (*SmsjobsStatusPredict)(nil)
)

type SmsjobsStatusPredict struct {
	_                  struct{} `tl:"flags,bitflag"`
	AllowInternational bool     `tl:"allow_international,omitempty:flags:0,implicit"`
	RecentSent         int32    `tl:"recent_sent"`
	RecentSince        int32    `tl:"recent_since"`
	RecentRemains      int32    `tl:"recent_remains"`
	TotalSent          int32    `tl:"total_sent"`
	TotalSince         int32    `tl:"total_since"`
	LastGiftSlug       *string  `tl:"last_gift_slug,omitempty:flags:1"`
	TermsURL           string   `tl:"terms_url"`
}

func (*SmsjobsStatusPredict) CRC() uint32 {
	return 0x2aee9191
}
func (*SmsjobsStatusPredict) _SmsjobsStatus() {}

type StatsBroadcastRevenueStats interface {
	tl.TLObject
	_StatsBroadcastRevenueStats()
}

var (
	_ StatsBroadcastRevenueStats = (*StatsBroadcastRevenueStatsPredict)(nil)
)

type StatsBroadcastRevenueStatsPredict struct {
	TopHoursGraph StatsGraph               `tl:"top_hours_graph"`
	RevenueGraph  StatsGraph               `tl:"revenue_graph"`
	Balances      BroadcastRevenueBalances `tl:"balances"`
	UsdRate       float64                  `tl:"usd_rate"`
}

func (*StatsBroadcastRevenueStatsPredict) CRC() uint32 {
	return 0x5407e297
}
func (*StatsBroadcastRevenueStatsPredict) _StatsBroadcastRevenueStats() {}

type StatsBroadcastRevenueTransactions interface {
	tl.TLObject
	_StatsBroadcastRevenueTransactions()
}

var (
	_ StatsBroadcastRevenueTransactions = (*StatsBroadcastRevenueTransactionsPredict)(nil)
)

type StatsBroadcastRevenueTransactionsPredict struct {
	Count        int32                         `tl:"count"`
	Transactions []BroadcastRevenueTransaction `tl:"transactions"`
}

func (*StatsBroadcastRevenueTransactionsPredict) CRC() uint32 {
	return 0x87158466
}
func (*StatsBroadcastRevenueTransactionsPredict) _StatsBroadcastRevenueTransactions() {}

type StatsBroadcastRevenueWithdrawalURL interface {
	tl.TLObject
	_StatsBroadcastRevenueWithdrawalURL()
}

var (
	_ StatsBroadcastRevenueWithdrawalURL = (*StatsBroadcastRevenueWithdrawalURLPredict)(nil)
)

type StatsBroadcastRevenueWithdrawalURLPredict struct {
	URL string `tl:"url"`
}

func (*StatsBroadcastRevenueWithdrawalURLPredict) CRC() uint32 {
	return 0xec659737
}
func (*StatsBroadcastRevenueWithdrawalURLPredict) _StatsBroadcastRevenueWithdrawalURL() {}

type StatsBroadcastStats interface {
	tl.TLObject
	_StatsBroadcastStats()
}

var (
	_ StatsBroadcastStats = (*StatsBroadcastStatsPredict)(nil)
)

type StatsBroadcastStatsPredict struct {
	Period                       StatsDateRangeDays        `tl:"period"`
	Followers                    StatsAbsValueAndPrev      `tl:"followers"`
	ViewsPerPost                 StatsAbsValueAndPrev      `tl:"views_per_post"`
	SharesPerPost                StatsAbsValueAndPrev      `tl:"shares_per_post"`
	ReactionsPerPost             StatsAbsValueAndPrev      `tl:"reactions_per_post"`
	ViewsPerStory                StatsAbsValueAndPrev      `tl:"views_per_story"`
	SharesPerStory               StatsAbsValueAndPrev      `tl:"shares_per_story"`
	ReactionsPerStory            StatsAbsValueAndPrev      `tl:"reactions_per_story"`
	EnabledNotifications         StatsPercentValue         `tl:"enabled_notifications"`
	GrowthGraph                  StatsGraph                `tl:"growth_graph"`
	FollowersGraph               StatsGraph                `tl:"followers_graph"`
	MuteGraph                    StatsGraph                `tl:"mute_graph"`
	TopHoursGraph                StatsGraph                `tl:"top_hours_graph"`
	InteractionsGraph            StatsGraph                `tl:"interactions_graph"`
	IvInteractionsGraph          StatsGraph                `tl:"iv_interactions_graph"`
	ViewsBySourceGraph           StatsGraph                `tl:"views_by_source_graph"`
	NewFollowersBySourceGraph    StatsGraph                `tl:"new_followers_by_source_graph"`
	LanguagesGraph               StatsGraph                `tl:"languages_graph"`
	ReactionsByEmotionGraph      StatsGraph                `tl:"reactions_by_emotion_graph"`
	StoryInteractionsGraph       StatsGraph                `tl:"story_interactions_graph"`
	StoryReactionsByEmotionGraph StatsGraph                `tl:"story_reactions_by_emotion_graph"`
	RecentPostsInteractions      []PostInteractionCounters `tl:"recent_posts_interactions"`
}

func (*StatsBroadcastStatsPredict) CRC() uint32 {
	return 0x396ca5fc
}
func (*StatsBroadcastStatsPredict) _StatsBroadcastStats() {}

type StatsMegagroupStats interface {
	tl.TLObject
	_StatsMegagroupStats()
}

var (
	_ StatsMegagroupStats = (*StatsMegagroupStatsPredict)(nil)
)

type StatsMegagroupStatsPredict struct {
	Period                  StatsDateRangeDays     `tl:"period"`
	Members                 StatsAbsValueAndPrev   `tl:"members"`
	Messages                StatsAbsValueAndPrev   `tl:"messages"`
	Viewers                 StatsAbsValueAndPrev   `tl:"viewers"`
	Posters                 StatsAbsValueAndPrev   `tl:"posters"`
	GrowthGraph             StatsGraph             `tl:"growth_graph"`
	MembersGraph            StatsGraph             `tl:"members_graph"`
	NewMembersBySourceGraph StatsGraph             `tl:"new_members_by_source_graph"`
	LanguagesGraph          StatsGraph             `tl:"languages_graph"`
	MessagesGraph           StatsGraph             `tl:"messages_graph"`
	ActionsGraph            StatsGraph             `tl:"actions_graph"`
	TopHoursGraph           StatsGraph             `tl:"top_hours_graph"`
	WeekdaysGraph           StatsGraph             `tl:"weekdays_graph"`
	TopPosters              []StatsGroupTopPoster  `tl:"top_posters"`
	TopAdmins               []StatsGroupTopAdmin   `tl:"top_admins"`
	TopInviters             []StatsGroupTopInviter `tl:"top_inviters"`
	Users                   []User                 `tl:"users"`
}

func (*StatsMegagroupStatsPredict) CRC() uint32 {
	return 0xef7ff916
}
func (*StatsMegagroupStatsPredict) _StatsMegagroupStats() {}

type StatsMessageStats interface {
	tl.TLObject
	_StatsMessageStats()
}

var (
	_ StatsMessageStats = (*StatsMessageStatsPredict)(nil)
)

type StatsMessageStatsPredict struct {
	ViewsGraph              StatsGraph `tl:"views_graph"`
	ReactionsByEmotionGraph StatsGraph `tl:"reactions_by_emotion_graph"`
}

func (*StatsMessageStatsPredict) CRC() uint32 {
	return 0x7fe91c14
}
func (*StatsMessageStatsPredict) _StatsMessageStats() {}

type StatsPublicForwards interface {
	tl.TLObject
	_StatsPublicForwards()
}

var (
	_ StatsPublicForwards = (*StatsPublicForwardsPredict)(nil)
)

type StatsPublicForwardsPredict struct {
	_          struct{}        `tl:"flags,bitflag"`
	Count      int32           `tl:"count"`
	Forwards   []PublicForward `tl:"forwards"`
	NextOffset *string         `tl:"next_offset,omitempty:flags:0"`
	Chats      []Chat          `tl:"chats"`
	Users      []User          `tl:"users"`
}

func (*StatsPublicForwardsPredict) CRC() uint32 {
	return 0x93037e20
}
func (*StatsPublicForwardsPredict) _StatsPublicForwards() {}

type StatsStoryStats interface {
	tl.TLObject
	_StatsStoryStats()
}

var (
	_ StatsStoryStats = (*StatsStoryStatsPredict)(nil)
)

type StatsStoryStatsPredict struct {
	ViewsGraph              StatsGraph `tl:"views_graph"`
	ReactionsByEmotionGraph StatsGraph `tl:"reactions_by_emotion_graph"`
}

func (*StatsStoryStatsPredict) CRC() uint32 {
	return 0x50cd067c
}
func (*StatsStoryStatsPredict) _StatsStoryStats() {}

type StickersSuggestedShortName interface {
	tl.TLObject
	_StickersSuggestedShortName()
}

var (
	_ StickersSuggestedShortName = (*StickersSuggestedShortNamePredict)(nil)
)

type StickersSuggestedShortNamePredict struct {
	ShortName string `tl:"short_name"`
}

func (*StickersSuggestedShortNamePredict) CRC() uint32 {
	return 0x85fea03f
}
func (*StickersSuggestedShortNamePredict) _StickersSuggestedShortName() {}

type StorageFileType interface {
	tl.TLObject
	_StorageFileType()
}

var (
	_ StorageFileType = (*StorageFileUnknownPredict)(nil)
	_ StorageFileType = (*StorageFilePartialPredict)(nil)
	_ StorageFileType = (*StorageFileJpegPredict)(nil)
	_ StorageFileType = (*StorageFileGifPredict)(nil)
	_ StorageFileType = (*StorageFilePngPredict)(nil)
	_ StorageFileType = (*StorageFilePdfPredict)(nil)
	_ StorageFileType = (*StorageFileMp3Predict)(nil)
	_ StorageFileType = (*StorageFileMovPredict)(nil)
	_ StorageFileType = (*StorageFileMp4Predict)(nil)
	_ StorageFileType = (*StorageFileWebpPredict)(nil)
)

type StorageFileUnknownPredict struct{}

func (*StorageFileUnknownPredict) CRC() uint32 {
	return 0xaa963b05
}
func (*StorageFileUnknownPredict) _StorageFileType() {}

type StorageFilePartialPredict struct{}

func (*StorageFilePartialPredict) CRC() uint32 {
	return 0x40bc6f52
}
func (*StorageFilePartialPredict) _StorageFileType() {}

type StorageFileJpegPredict struct{}

func (*StorageFileJpegPredict) CRC() uint32 {
	return 0x007efe0e
}
func (*StorageFileJpegPredict) _StorageFileType() {}

type StorageFileGifPredict struct{}

func (*StorageFileGifPredict) CRC() uint32 {
	return 0xcae1aadf
}
func (*StorageFileGifPredict) _StorageFileType() {}

type StorageFilePngPredict struct{}

func (*StorageFilePngPredict) CRC() uint32 {
	return 0x0a4f63c0
}
func (*StorageFilePngPredict) _StorageFileType() {}

type StorageFilePdfPredict struct{}

func (*StorageFilePdfPredict) CRC() uint32 {
	return 0xae1e508d
}
func (*StorageFilePdfPredict) _StorageFileType() {}

type StorageFileMp3Predict struct{}

func (*StorageFileMp3Predict) CRC() uint32 {
	return 0x528a0677
}
func (*StorageFileMp3Predict) _StorageFileType() {}

type StorageFileMovPredict struct{}

func (*StorageFileMovPredict) CRC() uint32 {
	return 0x4b09ebbc
}
func (*StorageFileMovPredict) _StorageFileType() {}

type StorageFileMp4Predict struct{}

func (*StorageFileMp4Predict) CRC() uint32 {
	return 0xb3cea0e4
}
func (*StorageFileMp4Predict) _StorageFileType() {}

type StorageFileWebpPredict struct{}

func (*StorageFileWebpPredict) CRC() uint32 {
	return 0x1081464c
}
func (*StorageFileWebpPredict) _StorageFileType() {}

type StoriesAllStories interface {
	tl.TLObject
	_StoriesAllStories()
}

var (
	_ StoriesAllStories = (*StoriesAllStoriesNotModifiedPredict)(nil)
	_ StoriesAllStories = (*StoriesAllStoriesPredict)(nil)
)

type StoriesAllStoriesNotModifiedPredict struct {
	_           struct{}           `tl:"flags,bitflag"`
	State       string             `tl:"state"`
	StealthMode StoriesStealthMode `tl:"stealth_mode"`
}

func (*StoriesAllStoriesNotModifiedPredict) CRC() uint32 {
	return 0x1158fe3e
}
func (*StoriesAllStoriesNotModifiedPredict) _StoriesAllStories() {}

type StoriesAllStoriesPredict struct {
	_           struct{}           `tl:"flags,bitflag"`
	HasMore     bool               `tl:"has_more,omitempty:flags:0,implicit"`
	Count       int32              `tl:"count"`
	State       string             `tl:"state"`
	PeerStories []PeerStories      `tl:"peer_stories"`
	Chats       []Chat             `tl:"chats"`
	Users       []User             `tl:"users"`
	StealthMode StoriesStealthMode `tl:"stealth_mode"`
}

func (*StoriesAllStoriesPredict) CRC() uint32 {
	return 0x6efc5e81
}
func (*StoriesAllStoriesPredict) _StoriesAllStories() {}

type StoriesFoundStories interface {
	tl.TLObject
	_StoriesFoundStories()
}

var (
	_ StoriesFoundStories = (*StoriesFoundStoriesPredict)(nil)
)

type StoriesFoundStoriesPredict struct {
	_          struct{}     `tl:"flags,bitflag"`
	Count      int32        `tl:"count"`
	Stories    []FoundStory `tl:"stories"`
	NextOffset *string      `tl:"next_offset,omitempty:flags:0"`
	Chats      []Chat       `tl:"chats"`
	Users      []User       `tl:"users"`
}

func (*StoriesFoundStoriesPredict) CRC() uint32 {
	return 0xe2de7737
}
func (*StoriesFoundStoriesPredict) _StoriesFoundStories() {}

type StoriesPeerStories interface {
	tl.TLObject
	_StoriesPeerStories()
}

var (
	_ StoriesPeerStories = (*StoriesPeerStoriesPredict)(nil)
)

type StoriesPeerStoriesPredict struct {
	Stories PeerStories `tl:"stories"`
	Chats   []Chat      `tl:"chats"`
	Users   []User      `tl:"users"`
}

func (*StoriesPeerStoriesPredict) CRC() uint32 {
	return 0xcae68768
}
func (*StoriesPeerStoriesPredict) _StoriesPeerStories() {}

type StoriesStories interface {
	tl.TLObject
	_StoriesStories()
}

var (
	_ StoriesStories = (*StoriesStoriesPredict)(nil)
)

type StoriesStoriesPredict struct {
	_           struct{}    `tl:"flags,bitflag"`
	Count       int32       `tl:"count"`
	Stories     []StoryItem `tl:"stories"`
	PinnedToTop []int32     `tl:"pinned_to_top,omitempty:flags:0"`
	Chats       []Chat      `tl:"chats"`
	Users       []User      `tl:"users"`
}

func (*StoriesStoriesPredict) CRC() uint32 {
	return 0x63c3dd0a
}
func (*StoriesStoriesPredict) _StoriesStories() {}

type StoriesStoryReactionsList interface {
	tl.TLObject
	_StoriesStoryReactionsList()
}

var (
	_ StoriesStoryReactionsList = (*StoriesStoryReactionsListPredict)(nil)
)

type StoriesStoryReactionsListPredict struct {
	_          struct{}        `tl:"flags,bitflag"`
	Count      int32           `tl:"count"`
	Reactions  []StoryReaction `tl:"reactions"`
	Chats      []Chat          `tl:"chats"`
	Users      []User          `tl:"users"`
	NextOffset *string         `tl:"next_offset,omitempty:flags:0"`
}

func (*StoriesStoryReactionsListPredict) CRC() uint32 {
	return 0xaa5f789c
}
func (*StoriesStoryReactionsListPredict) _StoriesStoryReactionsList() {}

type StoriesStoryViews interface {
	tl.TLObject
	_StoriesStoryViews()
}

var (
	_ StoriesStoryViews = (*StoriesStoryViewsPredict)(nil)
)

type StoriesStoryViewsPredict struct {
	Views []StoryViews `tl:"views"`
	Users []User       `tl:"users"`
}

func (*StoriesStoryViewsPredict) CRC() uint32 {
	return 0xde9eed1d
}
func (*StoriesStoryViewsPredict) _StoriesStoryViews() {}

type StoriesStoryViewsList interface {
	tl.TLObject
	_StoriesStoryViewsList()
}

var (
	_ StoriesStoryViewsList = (*StoriesStoryViewsListPredict)(nil)
)

type StoriesStoryViewsListPredict struct {
	_              struct{}    `tl:"flags,bitflag"`
	Count          int32       `tl:"count"`
	ViewsCount     int32       `tl:"views_count"`
	ForwardsCount  int32       `tl:"forwards_count"`
	ReactionsCount int32       `tl:"reactions_count"`
	Views          []StoryView `tl:"views"`
	Chats          []Chat      `tl:"chats"`
	Users          []User      `tl:"users"`
	NextOffset     *string     `tl:"next_offset,omitempty:flags:0"`
}

func (*StoriesStoryViewsListPredict) CRC() uint32 {
	return 0x59d78fc5
}
func (*StoriesStoryViewsListPredict) _StoriesStoryViewsList() {}

type UpdatesChannelDifference interface {
	tl.TLObject
	_UpdatesChannelDifference()
}

var (
	_ UpdatesChannelDifference = (*UpdatesChannelDifferenceEmptyPredict)(nil)
	_ UpdatesChannelDifference = (*UpdatesChannelDifferenceTooLongPredict)(nil)
	_ UpdatesChannelDifference = (*UpdatesChannelDifferencePredict)(nil)
)

type UpdatesChannelDifferenceEmptyPredict struct {
	_       struct{} `tl:"flags,bitflag"`
	Final   bool     `tl:"final,omitempty:flags:0,implicit"`
	Pts     int32    `tl:"pts"`
	Timeout *int32   `tl:"timeout,omitempty:flags:1"`
}

func (*UpdatesChannelDifferenceEmptyPredict) CRC() uint32 {
	return 0x3e11affb
}
func (*UpdatesChannelDifferenceEmptyPredict) _UpdatesChannelDifference() {}

type UpdatesChannelDifferenceTooLongPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Final    bool      `tl:"final,omitempty:flags:0,implicit"`
	Timeout  *int32    `tl:"timeout,omitempty:flags:1"`
	Dialog   Dialog    `tl:"dialog"`
	Messages []Message `tl:"messages"`
	Chats    []Chat    `tl:"chats"`
	Users    []User    `tl:"users"`
}

func (*UpdatesChannelDifferenceTooLongPredict) CRC() uint32 {
	return 0xa4bcc6fe
}
func (*UpdatesChannelDifferenceTooLongPredict) _UpdatesChannelDifference() {}

type UpdatesChannelDifferencePredict struct {
	_            struct{}  `tl:"flags,bitflag"`
	Final        bool      `tl:"final,omitempty:flags:0,implicit"`
	Pts          int32     `tl:"pts"`
	Timeout      *int32    `tl:"timeout,omitempty:flags:1"`
	NewMessages  []Message `tl:"new_messages"`
	OtherUpdates []Update  `tl:"other_updates"`
	Chats        []Chat    `tl:"chats"`
	Users        []User    `tl:"users"`
}

func (*UpdatesChannelDifferencePredict) CRC() uint32 {
	return 0x2064674e
}
func (*UpdatesChannelDifferencePredict) _UpdatesChannelDifference() {}

type UpdatesDifference interface {
	tl.TLObject
	_UpdatesDifference()
}

var (
	_ UpdatesDifference = (*UpdatesDifferenceEmptyPredict)(nil)
	_ UpdatesDifference = (*UpdatesDifferencePredict)(nil)
	_ UpdatesDifference = (*UpdatesDifferenceSlicePredict)(nil)
	_ UpdatesDifference = (*UpdatesDifferenceTooLongPredict)(nil)
)

type UpdatesDifferenceEmptyPredict struct {
	Date int32 `tl:"date"`
	Seq  int32 `tl:"seq"`
}

func (*UpdatesDifferenceEmptyPredict) CRC() uint32 {
	return 0x5d75a138
}
func (*UpdatesDifferenceEmptyPredict) _UpdatesDifference() {}

type UpdatesDifferencePredict struct {
	NewMessages          []Message          `tl:"new_messages"`
	NewEncryptedMessages []EncryptedMessage `tl:"new_encrypted_messages"`
	OtherUpdates         []Update           `tl:"other_updates"`
	Chats                []Chat             `tl:"chats"`
	Users                []User             `tl:"users"`
	State                UpdatesState       `tl:"state"`
}

func (*UpdatesDifferencePredict) CRC() uint32 {
	return 0x00f49ca0
}
func (*UpdatesDifferencePredict) _UpdatesDifference() {}

type UpdatesDifferenceSlicePredict struct {
	NewMessages          []Message          `tl:"new_messages"`
	NewEncryptedMessages []EncryptedMessage `tl:"new_encrypted_messages"`
	OtherUpdates         []Update           `tl:"other_updates"`
	Chats                []Chat             `tl:"chats"`
	Users                []User             `tl:"users"`
	IntermediateState    UpdatesState       `tl:"intermediate_state"`
}

func (*UpdatesDifferenceSlicePredict) CRC() uint32 {
	return 0xa8fb1981
}
func (*UpdatesDifferenceSlicePredict) _UpdatesDifference() {}

type UpdatesDifferenceTooLongPredict struct {
	Pts int32 `tl:"pts"`
}

func (*UpdatesDifferenceTooLongPredict) CRC() uint32 {
	return 0x4afe8f6d
}
func (*UpdatesDifferenceTooLongPredict) _UpdatesDifference() {}

type UpdatesState interface {
	tl.TLObject
	_UpdatesState()
}

var (
	_ UpdatesState = (*UpdatesStatePredict)(nil)
)

type UpdatesStatePredict struct {
	Pts         int32 `tl:"pts"`
	Qts         int32 `tl:"qts"`
	Date        int32 `tl:"date"`
	Seq         int32 `tl:"seq"`
	UnreadCount int32 `tl:"unread_count"`
}

func (*UpdatesStatePredict) CRC() uint32 {
	return 0xa56c2a3e
}
func (*UpdatesStatePredict) _UpdatesState() {}

type UploadCdnFile interface {
	tl.TLObject
	_UploadCdnFile()
}

var (
	_ UploadCdnFile = (*UploadCdnFileReuploadNeededPredict)(nil)
	_ UploadCdnFile = (*UploadCdnFilePredict)(nil)
)

type UploadCdnFileReuploadNeededPredict struct {
	RequestToken []byte `tl:"request_token"`
}

func (*UploadCdnFileReuploadNeededPredict) CRC() uint32 {
	return 0xeea8e46e
}
func (*UploadCdnFileReuploadNeededPredict) _UploadCdnFile() {}

type UploadCdnFilePredict struct {
	Bytes []byte `tl:"bytes"`
}

func (*UploadCdnFilePredict) CRC() uint32 {
	return 0xa99fca4f
}
func (*UploadCdnFilePredict) _UploadCdnFile() {}

type UploadFile interface {
	tl.TLObject
	_UploadFile()
}

var (
	_ UploadFile = (*UploadFilePredict)(nil)
	_ UploadFile = (*UploadFileCdnRedirectPredict)(nil)
)

type UploadFilePredict struct {
	Type  StorageFileType `tl:"type"`
	Mtime int32           `tl:"mtime"`
	Bytes []byte          `tl:"bytes"`
}

func (*UploadFilePredict) CRC() uint32 {
	return 0x096a18d5
}
func (*UploadFilePredict) _UploadFile() {}

type UploadFileCdnRedirectPredict struct {
	DcID          int32      `tl:"dc_id"`
	FileToken     []byte     `tl:"file_token"`
	EncryptionKey []byte     `tl:"encryption_key"`
	EncryptionIv  []byte     `tl:"encryption_iv"`
	FileHashes    []FileHash `tl:"file_hashes"`
}

func (*UploadFileCdnRedirectPredict) CRC() uint32 {
	return 0xf18cda44
}
func (*UploadFileCdnRedirectPredict) _UploadFile() {}

type UploadWebFile interface {
	tl.TLObject
	_UploadWebFile()
}

var (
	_ UploadWebFile = (*UploadWebFilePredict)(nil)
)

type UploadWebFilePredict struct {
	Size     int32           `tl:"size"`
	MimeType string          `tl:"mime_type"`
	FileType StorageFileType `tl:"file_type"`
	Mtime    int32           `tl:"mtime"`
	Bytes    []byte          `tl:"bytes"`
}

func (*UploadWebFilePredict) CRC() uint32 {
	return 0x21e753bc
}
func (*UploadWebFilePredict) _UploadWebFile() {}

type UsersUserFull interface {
	tl.TLObject
	_UsersUserFull()
}

var (
	_ UsersUserFull = (*UsersUserFullPredict)(nil)
)

type UsersUserFullPredict struct {
	FullUser UserFull `tl:"full_user"`
	Chats    []Chat   `tl:"chats"`
	Users    []User   `tl:"users"`
}

func (*UsersUserFullPredict) CRC() uint32 {
	return 0x3b6d152e
}
func (*UsersUserFullPredict) _UsersUserFull() {}

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

type InvokeAfterMsgRequestPredict[X tl.TLObject] struct {
	MsgID int64 `tl:"msg_id"`
	Query X     `tl:"query"`
}

func (*InvokeAfterMsgRequestPredict[X]) CRC() uint32 {
	return 0xcb9f372d
}

func InvokeAfterMsg[X tl.TLObject](ctx context.Context, m Requester, i InvokeAfterMsgRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeAfterMsgsRequestPredict[X tl.TLObject] struct {
	MsgIds []int64 `tl:"msg_ids"`
	Query  X       `tl:"query"`
}

func (*InvokeAfterMsgsRequestPredict[X]) CRC() uint32 {
	return 0x3dc4b4f0
}

func InvokeAfterMsgs[X tl.TLObject](ctx context.Context, m Requester, i InvokeAfterMsgsRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InitConnectionRequestPredict[X tl.TLObject] struct {
	_              struct{}         `tl:"flags,bitflag"`
	APIID          int32            `tl:"api_id"`
	DeviceModel    string           `tl:"device_model"`
	SystemVersion  string           `tl:"system_version"`
	AppVersion     string           `tl:"app_version"`
	SystemLangCode string           `tl:"system_lang_code"`
	LangPack       string           `tl:"lang_pack"`
	LangCode       string           `tl:"lang_code"`
	Proxy          InputClientProxy `tl:"proxy,omitempty:flags:0"`
	Params         JSONValue        `tl:"params,omitempty:flags:1"`
	Query          X                `tl:"query"`
}

func (*InitConnectionRequestPredict[X]) CRC() uint32 {
	return 0xc1cd5ea9
}

func InitConnection[X tl.TLObject](ctx context.Context, m Requester, i InitConnectionRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithLayerRequestPredict[X tl.TLObject] struct {
	Layer int32 `tl:"layer"`
	Query X     `tl:"query"`
}

func (*InvokeWithLayerRequestPredict[X]) CRC() uint32 {
	return 0xda9b0d0d
}

func InvokeWithLayer[X tl.TLObject](ctx context.Context, m Requester, i InvokeWithLayerRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithoutUpdatesRequestPredict[X tl.TLObject] struct {
	Query X `tl:"query"`
}

func (*InvokeWithoutUpdatesRequestPredict[X]) CRC() uint32 {
	return 0xbf9459b7
}

func InvokeWithoutUpdates[X tl.TLObject](ctx context.Context, m Requester, i InvokeWithoutUpdatesRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithMessagesRangeRequestPredict[X tl.TLObject] struct {
	Range MessageRange `tl:"range"`
	Query X            `tl:"query"`
}

func (*InvokeWithMessagesRangeRequestPredict[X]) CRC() uint32 {
	return 0x365275f2
}

func InvokeWithMessagesRange[X tl.TLObject](ctx context.Context, m Requester, i InvokeWithMessagesRangeRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithTakeoutRequestPredict[X tl.TLObject] struct {
	TakeoutID int64 `tl:"takeout_id"`
	Query     X     `tl:"query"`
}

func (*InvokeWithTakeoutRequestPredict[X]) CRC() uint32 {
	return 0xaca9fd2e
}

func InvokeWithTakeout[X tl.TLObject](ctx context.Context, m Requester, i InvokeWithTakeoutRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithBusinessConnectionRequestPredict[X tl.TLObject] struct {
	ConnectionID string `tl:"connection_id"`
	Query        X      `tl:"query"`
}

func (*InvokeWithBusinessConnectionRequestPredict[X]) CRC() uint32 {
	return 0xdd289f8e
}

func InvokeWithBusinessConnection[X tl.TLObject](ctx context.Context, m Requester, i InvokeWithBusinessConnectionRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithGooglePlayIntegrityRequestPredict[X tl.TLObject] struct {
	Nonce string `tl:"nonce"`
	Token string `tl:"token"`
	Query X      `tl:"query"`
}

func (*InvokeWithGooglePlayIntegrityRequestPredict[X]) CRC() uint32 {
	return 0x1df92984
}

func InvokeWithGooglePlayIntegrity[X tl.TLObject](ctx context.Context, m Requester, i InvokeWithGooglePlayIntegrityRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type InvokeWithApnsSecretRequestPredict[X tl.TLObject] struct {
	Nonce  string `tl:"nonce"`
	Secret string `tl:"secret"`
	Query  X      `tl:"query"`
}

func (*InvokeWithApnsSecretRequestPredict[X]) CRC() uint32 {
	return 0x0dae54f8
}

func InvokeWithApnsSecret[X tl.TLObject](ctx context.Context, m Requester, i InvokeWithApnsSecretRequestPredict[X]) (X, error) {
	var res X
	return res, request(ctx, m, &i, &res)
}

type AccountRegisterDeviceRequestPredict struct {
	_          struct{} `tl:"flags,bitflag"`
	NoMuted    bool     `tl:"no_muted,omitempty:flags:0,implicit"`
	TokenType  int32    `tl:"token_type"`
	Token      string   `tl:"token"`
	AppSandbox bool     `tl:"app_sandbox"`
	Secret     []byte   `tl:"secret"`
	OtherUids  []int64  `tl:"other_uids"`
}

func (*AccountRegisterDeviceRequestPredict) CRC() uint32 {
	return 0xec86017a
}

func AccountRegisterDevice(ctx context.Context, m Requester, i AccountRegisterDeviceRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUnregisterDeviceRequestPredict struct {
	TokenType int32   `tl:"token_type"`
	Token     string  `tl:"token"`
	OtherUids []int64 `tl:"other_uids"`
}

func (*AccountUnregisterDeviceRequestPredict) CRC() uint32 {
	return 0x6a0d3206
}

func AccountUnregisterDevice(ctx context.Context, m Requester, i AccountUnregisterDeviceRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateNotifySettingsRequestPredict struct {
	Peer     InputNotifyPeer         `tl:"peer"`
	Settings InputPeerNotifySettings `tl:"settings"`
}

func (*AccountUpdateNotifySettingsRequestPredict) CRC() uint32 {
	return 0x84be5b93
}

func AccountUpdateNotifySettings(ctx context.Context, m Requester, i AccountUpdateNotifySettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetNotifySettingsRequestPredict struct {
	Peer InputNotifyPeer `tl:"peer"`
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
	FirstName *string  `tl:"first_name,omitempty:flags:0"`
	LastName  *string  `tl:"last_name,omitempty:flags:1"`
	About     *string  `tl:"about,omitempty:flags:2"`
}

func (*AccountUpdateProfileRequestPredict) CRC() uint32 {
	return 0x78515775
}

func AccountUpdateProfile(ctx context.Context, m Requester, i AccountUpdateProfileRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateStatusRequestPredict struct {
	Offline bool `tl:"offline"`
}

func (*AccountUpdateStatusRequestPredict) CRC() uint32 {
	return 0x6628562c
}

func AccountUpdateStatus(ctx context.Context, m Requester, i AccountUpdateStatusRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetWallPapersRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*AccountGetWallPapersRequestPredict) CRC() uint32 {
	return 0x07967d36
}

func AccountGetWallPapers(ctx context.Context, m Requester, i AccountGetWallPapersRequestPredict) (AccountWallPapers, error) {
	var res AccountWallPapers
	return res, request(ctx, m, &i, &res)
}

type AccountReportPeerRequestPredict struct {
	Peer    InputPeer    `tl:"peer"`
	Reason  ReportReason `tl:"reason"`
	Message string       `tl:"message"`
}

func (*AccountReportPeerRequestPredict) CRC() uint32 {
	return 0xc5ba3d86
}

func AccountReportPeer(ctx context.Context, m Requester, i AccountReportPeerRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountCheckUsernameRequestPredict struct {
	Username string `tl:"username"`
}

func (*AccountCheckUsernameRequestPredict) CRC() uint32 {
	return 0x2714d86c
}

func AccountCheckUsername(ctx context.Context, m Requester, i AccountCheckUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateUsernameRequestPredict struct {
	Username string `tl:"username"`
}

func (*AccountUpdateUsernameRequestPredict) CRC() uint32 {
	return 0x3e0bdd7c
}

func AccountUpdateUsername(ctx context.Context, m Requester, i AccountUpdateUsernameRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type AccountGetPrivacyRequestPredict struct {
	Key InputPrivacyKey `tl:"key"`
}

func (*AccountGetPrivacyRequestPredict) CRC() uint32 {
	return 0xdadbc950
}

func AccountGetPrivacy(ctx context.Context, m Requester, i AccountGetPrivacyRequestPredict) (AccountPrivacyRules, error) {
	var res AccountPrivacyRules
	return res, request(ctx, m, &i, &res)
}

type AccountSetPrivacyRequestPredict struct {
	Key   InputPrivacyKey    `tl:"key"`
	Rules []InputPrivacyRule `tl:"rules"`
}

func (*AccountSetPrivacyRequestPredict) CRC() uint32 {
	return 0xc9f81ce8
}

func AccountSetPrivacy(ctx context.Context, m Requester, i AccountSetPrivacyRequestPredict) (AccountPrivacyRules, error) {
	var res AccountPrivacyRules
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteAccountRequestPredict struct {
	_        struct{}              `tl:"flags,bitflag"`
	Reason   string                `tl:"reason"`
	Password InputCheckPasswordSRP `tl:"password,omitempty:flags:0"`
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
	return 0x08fc711d
}

func AccountGetAccountTTL(ctx context.Context, m Requester, i AccountGetAccountTTLRequestPredict) (AccountDaysTTL, error) {
	var res AccountDaysTTL
	return res, request(ctx, m, &i, &res)
}

type AccountSetAccountTTLRequestPredict struct {
	TTL AccountDaysTTL `tl:"ttl"`
}

func (*AccountSetAccountTTLRequestPredict) CRC() uint32 {
	return 0x2442485e
}

func AccountSetAccountTTL(ctx context.Context, m Requester, i AccountSetAccountTTLRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendChangePhoneCodeRequestPredict struct {
	PhoneNumber string       `tl:"phone_number"`
	Settings    CodeSettings `tl:"settings"`
}

func (*AccountSendChangePhoneCodeRequestPredict) CRC() uint32 {
	return 0x82574ae5
}

func AccountSendChangePhoneCode(ctx context.Context, m Requester, i AccountSendChangePhoneCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountChangePhoneRequestPredict struct {
	PhoneNumber   string `tl:"phone_number"`
	PhoneCodeHash string `tl:"phone_code_hash"`
	PhoneCode     string `tl:"phone_code"`
}

func (*AccountChangePhoneRequestPredict) CRC() uint32 {
	return 0x70c32edb
}

func AccountChangePhone(ctx context.Context, m Requester, i AccountChangePhoneRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateDeviceLockedRequestPredict struct {
	Period int32 `tl:"period"`
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
	Hash int64 `tl:"hash"`
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
	Password InputCheckPasswordSRP `tl:"password"`
}

func (*AccountGetPasswordSettingsRequestPredict) CRC() uint32 {
	return 0x9cd4eaf9
}

func AccountGetPasswordSettings(ctx context.Context, m Requester, i AccountGetPasswordSettingsRequestPredict) (AccountPasswordSettings, error) {
	var res AccountPasswordSettings
	return res, request(ctx, m, &i, &res)
}

type AccountUpdatePasswordSettingsRequestPredict struct {
	Password    InputCheckPasswordSRP        `tl:"password"`
	NewSettings AccountPasswordInputSettings `tl:"new_settings"`
}

func (*AccountUpdatePasswordSettingsRequestPredict) CRC() uint32 {
	return 0xa59b102f
}

func AccountUpdatePasswordSettings(ctx context.Context, m Requester, i AccountUpdatePasswordSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendConfirmPhoneCodeRequestPredict struct {
	Hash     string       `tl:"hash"`
	Settings CodeSettings `tl:"settings"`
}

func (*AccountSendConfirmPhoneCodeRequestPredict) CRC() uint32 {
	return 0x1b3faa88
}

func AccountSendConfirmPhoneCode(ctx context.Context, m Requester, i AccountSendConfirmPhoneCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountConfirmPhoneRequestPredict struct {
	PhoneCodeHash string `tl:"phone_code_hash"`
	PhoneCode     string `tl:"phone_code"`
}

func (*AccountConfirmPhoneRequestPredict) CRC() uint32 {
	return 0x5f2178c3
}

func AccountConfirmPhone(ctx context.Context, m Requester, i AccountConfirmPhoneRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetTmpPasswordRequestPredict struct {
	Password InputCheckPasswordSRP `tl:"password"`
	Period   int32                 `tl:"period"`
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
	Hash int64 `tl:"hash"`
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
	Types []SecureValueType `tl:"types"`
}

func (*AccountGetSecureValueRequestPredict) CRC() uint32 {
	return 0x73665bc2
}

func AccountGetSecureValue(ctx context.Context, m Requester, i AccountGetSecureValueRequestPredict) ([]SecureValue, error) {
	var res []SecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountSaveSecureValueRequestPredict struct {
	Value          InputSecureValue `tl:"value"`
	SecureSecretID int64            `tl:"secure_secret_id"`
}

func (*AccountSaveSecureValueRequestPredict) CRC() uint32 {
	return 0x899fe31d
}

func AccountSaveSecureValue(ctx context.Context, m Requester, i AccountSaveSecureValueRequestPredict) (SecureValue, error) {
	var res SecureValue
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteSecureValueRequestPredict struct {
	Types []SecureValueType `tl:"types"`
}

func (*AccountDeleteSecureValueRequestPredict) CRC() uint32 {
	return 0xb880bc4b
}

func AccountDeleteSecureValue(ctx context.Context, m Requester, i AccountDeleteSecureValueRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetAuthorizationFormRequestPredict struct {
	BotID     int64  `tl:"bot_id"`
	Scope     string `tl:"scope"`
	PublicKey string `tl:"public_key"`
}

func (*AccountGetAuthorizationFormRequestPredict) CRC() uint32 {
	return 0xa929597a
}

func AccountGetAuthorizationForm(ctx context.Context, m Requester, i AccountGetAuthorizationFormRequestPredict) (AccountAuthorizationForm, error) {
	var res AccountAuthorizationForm
	return res, request(ctx, m, &i, &res)
}

type AccountAcceptAuthorizationRequestPredict struct {
	BotID       int64                      `tl:"bot_id"`
	Scope       string                     `tl:"scope"`
	PublicKey   string                     `tl:"public_key"`
	ValueHashes []SecureValueHash          `tl:"value_hashes"`
	Credentials SecureCredentialsEncrypted `tl:"credentials"`
}

func (*AccountAcceptAuthorizationRequestPredict) CRC() uint32 {
	return 0xf3ed4c73
}

func AccountAcceptAuthorization(ctx context.Context, m Requester, i AccountAcceptAuthorizationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendVerifyPhoneCodeRequestPredict struct {
	PhoneNumber string       `tl:"phone_number"`
	Settings    CodeSettings `tl:"settings"`
}

func (*AccountSendVerifyPhoneCodeRequestPredict) CRC() uint32 {
	return 0xa5a356f9
}

func AccountSendVerifyPhoneCode(ctx context.Context, m Requester, i AccountSendVerifyPhoneCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AccountVerifyPhoneRequestPredict struct {
	PhoneNumber   string `tl:"phone_number"`
	PhoneCodeHash string `tl:"phone_code_hash"`
	PhoneCode     string `tl:"phone_code"`
}

func (*AccountVerifyPhoneRequestPredict) CRC() uint32 {
	return 0x4dd3a7f6
}

func AccountVerifyPhone(ctx context.Context, m Requester, i AccountVerifyPhoneRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountSendVerifyEmailCodeRequestPredict struct {
	Purpose EmailVerifyPurpose `tl:"purpose"`
	Email   string             `tl:"email"`
}

func (*AccountSendVerifyEmailCodeRequestPredict) CRC() uint32 {
	return 0x98e037bb
}

func AccountSendVerifyEmailCode(ctx context.Context, m Requester, i AccountSendVerifyEmailCodeRequestPredict) (AccountSentEmailCode, error) {
	var res AccountSentEmailCode
	return res, request(ctx, m, &i, &res)
}

type AccountVerifyEmailRequestPredict struct {
	Purpose      EmailVerifyPurpose `tl:"purpose"`
	Verification EmailVerification  `tl:"verification"`
}

func (*AccountVerifyEmailRequestPredict) CRC() uint32 {
	return 0x032da4cf
}

func AccountVerifyEmail(ctx context.Context, m Requester, i AccountVerifyEmailRequestPredict) (AccountEmailVerified, error) {
	var res AccountEmailVerified
	return res, request(ctx, m, &i, &res)
}

type AccountInitTakeoutSessionRequestPredict struct {
	_                 struct{} `tl:"flags,bitflag"`
	Contacts          bool     `tl:"contacts,omitempty:flags:0,implicit"`
	MessageUsers      bool     `tl:"message_users,omitempty:flags:1,implicit"`
	MessageChats      bool     `tl:"message_chats,omitempty:flags:2,implicit"`
	MessageMegagroups bool     `tl:"message_megagroups,omitempty:flags:3,implicit"`
	MessageChannels   bool     `tl:"message_channels,omitempty:flags:4,implicit"`
	Files             bool     `tl:"files,omitempty:flags:5,implicit"`
	FileMaxSize       *int64   `tl:"file_max_size,omitempty:flags:5"`
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
	Success bool     `tl:"success,omitempty:flags:0,implicit"`
}

func (*AccountFinishTakeoutSessionRequestPredict) CRC() uint32 {
	return 0x1d2652ee
}

func AccountFinishTakeoutSession(ctx context.Context, m Requester, i AccountFinishTakeoutSessionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountConfirmPasswordEmailRequestPredict struct {
	Code string `tl:"code"`
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
	Silent bool `tl:"silent"`
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
	CompareSound   bool            `tl:"compare_sound,omitempty:flags:1,implicit"`
	CompareStories bool            `tl:"compare_stories,omitempty:flags:2,implicit"`
	Peer           InputNotifyPeer `tl:"peer,omitempty:flags:0"`
}

func (*AccountGetNotifyExceptionsRequestPredict) CRC() uint32 {
	return 0x53577479
}

func AccountGetNotifyExceptions(ctx context.Context, m Requester, i AccountGetNotifyExceptionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type AccountGetWallPaperRequestPredict struct {
	Wallpaper InputWallPaper `tl:"wallpaper"`
}

func (*AccountGetWallPaperRequestPredict) CRC() uint32 {
	return 0xfc8ddbea
}

func AccountGetWallPaper(ctx context.Context, m Requester, i AccountGetWallPaperRequestPredict) (WallPaper, error) {
	var res WallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountUploadWallPaperRequestPredict struct {
	_        struct{}          `tl:"flags,bitflag"`
	ForChat  bool              `tl:"for_chat,omitempty:flags:0,implicit"`
	File     InputFile         `tl:"file"`
	MimeType string            `tl:"mime_type"`
	Settings WallPaperSettings `tl:"settings"`
}

func (*AccountUploadWallPaperRequestPredict) CRC() uint32 {
	return 0xe39a8f03
}

func AccountUploadWallPaper(ctx context.Context, m Requester, i AccountUploadWallPaperRequestPredict) (WallPaper, error) {
	var res WallPaper
	return res, request(ctx, m, &i, &res)
}

type AccountSaveWallPaperRequestPredict struct {
	Wallpaper InputWallPaper    `tl:"wallpaper"`
	Unsave    bool              `tl:"unsave"`
	Settings  WallPaperSettings `tl:"settings"`
}

func (*AccountSaveWallPaperRequestPredict) CRC() uint32 {
	return 0x6c5a5b37
}

func AccountSaveWallPaper(ctx context.Context, m Requester, i AccountSaveWallPaperRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountInstallWallPaperRequestPredict struct {
	Wallpaper InputWallPaper    `tl:"wallpaper"`
	Settings  WallPaperSettings `tl:"settings"`
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
	_        struct{}             `tl:"flags,bitflag"`
	Low      bool                 `tl:"low,omitempty:flags:0,implicit"`
	High     bool                 `tl:"high,omitempty:flags:1,implicit"`
	Settings AutoDownloadSettings `tl:"settings"`
}

func (*AccountSaveAutoDownloadSettingsRequestPredict) CRC() uint32 {
	return 0x76f36233
}

func AccountSaveAutoDownloadSettings(ctx context.Context, m Requester, i AccountSaveAutoDownloadSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUploadThemeRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	File     InputFile `tl:"file"`
	Thumb    InputFile `tl:"thumb,omitempty:flags:0"`
	FileName string    `tl:"file_name"`
	MimeType string    `tl:"mime_type"`
}

func (*AccountUploadThemeRequestPredict) CRC() uint32 {
	return 0x1c3db333
}

func AccountUploadTheme(ctx context.Context, m Requester, i AccountUploadThemeRequestPredict) (Document, error) {
	var res Document
	return res, request(ctx, m, &i, &res)
}

type AccountCreateThemeRequestPredict struct {
	_        struct{}             `tl:"flags,bitflag"`
	Slug     string               `tl:"slug"`
	Title    string               `tl:"title"`
	Document InputDocument        `tl:"document,omitempty:flags:2"`
	Settings []InputThemeSettings `tl:"settings,omitempty:flags:3"`
}

func (*AccountCreateThemeRequestPredict) CRC() uint32 {
	return 0x652e4400
}

func AccountCreateTheme(ctx context.Context, m Requester, i AccountCreateThemeRequestPredict) (Theme, error) {
	var res Theme
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateThemeRequestPredict struct {
	_        struct{}             `tl:"flags,bitflag"`
	Format   string               `tl:"format"`
	Theme    InputTheme           `tl:"theme"`
	Slug     *string              `tl:"slug,omitempty:flags:0"`
	Title    *string              `tl:"title,omitempty:flags:1"`
	Document InputDocument        `tl:"document,omitempty:flags:2"`
	Settings []InputThemeSettings `tl:"settings,omitempty:flags:3"`
}

func (*AccountUpdateThemeRequestPredict) CRC() uint32 {
	return 0x2bf40ccc
}

func AccountUpdateTheme(ctx context.Context, m Requester, i AccountUpdateThemeRequestPredict) (Theme, error) {
	var res Theme
	return res, request(ctx, m, &i, &res)
}

type AccountSaveThemeRequestPredict struct {
	Theme  InputTheme `tl:"theme"`
	Unsave bool       `tl:"unsave"`
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
	Dark      bool       `tl:"dark,omitempty:flags:0,implicit"`
	Theme     InputTheme `tl:"theme,omitempty:flags:1"`
	Format    *string    `tl:"format,omitempty:flags:2"`
	BaseTheme BaseTheme  `tl:"base_theme,omitempty:flags:3"`
}

func (*AccountInstallThemeRequestPredict) CRC() uint32 {
	return 0xc727bb3b
}

func AccountInstallTheme(ctx context.Context, m Requester, i AccountInstallThemeRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetThemeRequestPredict struct {
	Format string     `tl:"format"`
	Theme  InputTheme `tl:"theme"`
}

func (*AccountGetThemeRequestPredict) CRC() uint32 {
	return 0x3a5869ec
}

func AccountGetTheme(ctx context.Context, m Requester, i AccountGetThemeRequestPredict) (Theme, error) {
	var res Theme
	return res, request(ctx, m, &i, &res)
}

type AccountGetThemesRequestPredict struct {
	Format string `tl:"format"`
	Hash   int64  `tl:"hash"`
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
	SensitiveEnabled bool     `tl:"sensitive_enabled,omitempty:flags:0,implicit"`
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
	Wallpapers []InputWallPaper `tl:"wallpapers"`
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
	Settings GlobalPrivacySettings `tl:"settings"`
}

func (*AccountSetGlobalPrivacySettingsRequestPredict) CRC() uint32 {
	return 0x1edaaac2
}

func AccountSetGlobalPrivacySettings(ctx context.Context, m Requester, i AccountSetGlobalPrivacySettingsRequestPredict) (GlobalPrivacySettings, error) {
	var res GlobalPrivacySettings
	return res, request(ctx, m, &i, &res)
}

type AccountReportProfilePhotoRequestPredict struct {
	Peer    InputPeer    `tl:"peer"`
	PhotoID InputPhoto   `tl:"photo_id"`
	Reason  ReportReason `tl:"reason"`
	Message string       `tl:"message"`
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
	Hash int64 `tl:"hash"`
}

func (*AccountGetChatThemesRequestPredict) CRC() uint32 {
	return 0xd638de89
}

func AccountGetChatThemes(ctx context.Context, m Requester, i AccountGetChatThemesRequestPredict) (AccountThemes, error) {
	var res AccountThemes
	return res, request(ctx, m, &i, &res)
}

type AccountSetAuthorizationTTLRequestPredict struct {
	AuthorizationTTLDays int32 `tl:"authorization_ttl_days"`
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
	Confirmed                 bool     `tl:"confirmed,omitempty:flags:3,implicit"`
	Hash                      int64    `tl:"hash"`
	EncryptedRequestsDisabled *bool    `tl:"encrypted_requests_disabled,omitempty:flags:0"`
	CallRequestsDisabled      *bool    `tl:"call_requests_disabled,omitempty:flags:1"`
}

func (*AccountChangeAuthorizationSettingsRequestPredict) CRC() uint32 {
	return 0x40f48462
}

func AccountChangeAuthorizationSettings(ctx context.Context, m Requester, i AccountChangeAuthorizationSettingsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetSavedRingtonesRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*AccountGetSavedRingtonesRequestPredict) CRC() uint32 {
	return 0xe1902288
}

func AccountGetSavedRingtones(ctx context.Context, m Requester, i AccountGetSavedRingtonesRequestPredict) (AccountSavedRingtones, error) {
	var res AccountSavedRingtones
	return res, request(ctx, m, &i, &res)
}

type AccountSaveRingtoneRequestPredict struct {
	ID     InputDocument `tl:"id"`
	Unsave bool          `tl:"unsave"`
}

func (*AccountSaveRingtoneRequestPredict) CRC() uint32 {
	return 0x3dea5b03
}

func AccountSaveRingtone(ctx context.Context, m Requester, i AccountSaveRingtoneRequestPredict) (AccountSavedRingtone, error) {
	var res AccountSavedRingtone
	return res, request(ctx, m, &i, &res)
}

type AccountUploadRingtoneRequestPredict struct {
	File     InputFile `tl:"file"`
	FileName string    `tl:"file_name"`
	MimeType string    `tl:"mime_type"`
}

func (*AccountUploadRingtoneRequestPredict) CRC() uint32 {
	return 0x831a83a2
}

func AccountUploadRingtone(ctx context.Context, m Requester, i AccountUploadRingtoneRequestPredict) (Document, error) {
	var res Document
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateEmojiStatusRequestPredict struct {
	EmojiStatus EmojiStatus `tl:"emoji_status"`
}

func (*AccountUpdateEmojiStatusRequestPredict) CRC() uint32 {
	return 0xfbd3de6b
}

func AccountUpdateEmojiStatus(ctx context.Context, m Requester, i AccountUpdateEmojiStatusRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultEmojiStatusesRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*AccountGetDefaultEmojiStatusesRequestPredict) CRC() uint32 {
	return 0xd6753386
}

func AccountGetDefaultEmojiStatuses(ctx context.Context, m Requester, i AccountGetDefaultEmojiStatusesRequestPredict) (AccountEmojiStatuses, error) {
	var res AccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountGetRecentEmojiStatusesRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*AccountGetRecentEmojiStatusesRequestPredict) CRC() uint32 {
	return 0x0f578105
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
	Order []string `tl:"order"`
}

func (*AccountReorderUsernamesRequestPredict) CRC() uint32 {
	return 0xef500eab
}

func AccountReorderUsernames(ctx context.Context, m Requester, i AccountReorderUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountToggleUsernameRequestPredict struct {
	Username string `tl:"username"`
	Active   bool   `tl:"active"`
}

func (*AccountToggleUsernameRequestPredict) CRC() uint32 {
	return 0x58d6b376
}

func AccountToggleUsername(ctx context.Context, m Requester, i AccountToggleUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultProfilePhotoEmojisRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*AccountGetDefaultProfilePhotoEmojisRequestPredict) CRC() uint32 {
	return 0xe2750328
}

func AccountGetDefaultProfilePhotoEmojis(ctx context.Context, m Requester, i AccountGetDefaultProfilePhotoEmojisRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultGroupPhotoEmojisRequestPredict struct {
	Hash int64 `tl:"hash"`
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
	_          struct{}         `tl:"flags,bitflag"`
	Users      bool             `tl:"users,omitempty:flags:0,implicit"`
	Chats      bool             `tl:"chats,omitempty:flags:1,implicit"`
	Broadcasts bool             `tl:"broadcasts,omitempty:flags:2,implicit"`
	Peer       InputPeer        `tl:"peer,omitempty:flags:3"`
	Settings   AutoSaveSettings `tl:"settings"`
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
	Codes []string `tl:"codes"`
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
	ForProfile        bool     `tl:"for_profile,omitempty:flags:1,implicit"`
	Color             *int32   `tl:"color,omitempty:flags:2"`
	BackgroundEmojiID *int64   `tl:"background_emoji_id,omitempty:flags:0"`
}

func (*AccountUpdateColorRequestPredict) CRC() uint32 {
	return 0x7cefa15d
}

func AccountUpdateColor(ctx context.Context, m Requester, i AccountUpdateColorRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountGetDefaultBackgroundEmojisRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*AccountGetDefaultBackgroundEmojisRequestPredict) CRC() uint32 {
	return 0xa60ab9ce
}

func AccountGetDefaultBackgroundEmojis(ctx context.Context, m Requester, i AccountGetDefaultBackgroundEmojisRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type AccountGetChannelDefaultEmojiStatusesRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*AccountGetChannelDefaultEmojiStatusesRequestPredict) CRC() uint32 {
	return 0x7727a7d5
}

func AccountGetChannelDefaultEmojiStatuses(ctx context.Context, m Requester, i AccountGetChannelDefaultEmojiStatusesRequestPredict) (AccountEmojiStatuses, error) {
	var res AccountEmojiStatuses
	return res, request(ctx, m, &i, &res)
}

type AccountGetChannelRestrictedStatusEmojisRequestPredict struct {
	Hash int64 `tl:"hash"`
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
	BusinessWorkHours BusinessWorkHours `tl:"business_work_hours,omitempty:flags:0"`
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
	GeoPoint InputGeoPoint `tl:"geo_point,omitempty:flags:1"`
	Address  *string       `tl:"address,omitempty:flags:0"`
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
	Message InputBusinessGreetingMessage `tl:"message,omitempty:flags:0"`
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
	Message InputBusinessAwayMessage `tl:"message,omitempty:flags:0"`
}

func (*AccountUpdateBusinessAwayMessageRequestPredict) CRC() uint32 {
	return 0xa26a7fa5
}

func AccountUpdateBusinessAwayMessage(ctx context.Context, m Requester, i AccountUpdateBusinessAwayMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountUpdateConnectedBotRequestPredict struct {
	_          struct{}                   `tl:"flags,bitflag"`
	CanReply   bool                       `tl:"can_reply,omitempty:flags:0,implicit"`
	Deleted    bool                       `tl:"deleted,omitempty:flags:1,implicit"`
	Bot        InputUser                  `tl:"bot"`
	Recipients InputBusinessBotRecipients `tl:"recipients"`
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
	ConnectionID string `tl:"connection_id"`
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
	Intro InputBusinessIntro `tl:"intro,omitempty:flags:0"`
}

func (*AccountUpdateBusinessIntroRequestPredict) CRC() uint32 {
	return 0xa614d034
}

func AccountUpdateBusinessIntro(ctx context.Context, m Requester, i AccountUpdateBusinessIntroRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountToggleConnectedBotPausedRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	Paused bool      `tl:"paused"`
}

func (*AccountToggleConnectedBotPausedRequestPredict) CRC() uint32 {
	return 0x646e1097
}

func AccountToggleConnectedBotPaused(ctx context.Context, m Requester, i AccountToggleConnectedBotPausedRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountDisablePeerConnectedBotRequestPredict struct {
	Peer InputPeer `tl:"peer"`
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
	Birthday Birthday `tl:"birthday,omitempty:flags:0"`
}

func (*AccountUpdateBirthdayRequestPredict) CRC() uint32 {
	return 0xcc6e0c11
}

func AccountUpdateBirthday(ctx context.Context, m Requester, i AccountUpdateBirthdayRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountCreateBusinessChatLinkRequestPredict struct {
	Link InputBusinessChatLink `tl:"link"`
}

func (*AccountCreateBusinessChatLinkRequestPredict) CRC() uint32 {
	return 0x8851e68e
}

func AccountCreateBusinessChatLink(ctx context.Context, m Requester, i AccountCreateBusinessChatLinkRequestPredict) (BusinessChatLink, error) {
	var res BusinessChatLink
	return res, request(ctx, m, &i, &res)
}

type AccountEditBusinessChatLinkRequestPredict struct {
	Slug string                `tl:"slug"`
	Link InputBusinessChatLink `tl:"link"`
}

func (*AccountEditBusinessChatLinkRequestPredict) CRC() uint32 {
	return 0x8c3410af
}

func AccountEditBusinessChatLink(ctx context.Context, m Requester, i AccountEditBusinessChatLinkRequestPredict) (BusinessChatLink, error) {
	var res BusinessChatLink
	return res, request(ctx, m, &i, &res)
}

type AccountDeleteBusinessChatLinkRequestPredict struct {
	Slug string `tl:"slug"`
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
	Slug string `tl:"slug"`
}

func (*AccountResolveBusinessChatLinkRequestPredict) CRC() uint32 {
	return 0x5492e5ee
}

func AccountResolveBusinessChatLink(ctx context.Context, m Requester, i AccountResolveBusinessChatLinkRequestPredict) (AccountResolvedBusinessChatLinks, error) {
	var res AccountResolvedBusinessChatLinks
	return res, request(ctx, m, &i, &res)
}

type AccountUpdatePersonalChannelRequestPredict struct {
	Channel InputChannel `tl:"channel"`
}

func (*AccountUpdatePersonalChannelRequestPredict) CRC() uint32 {
	return 0xd94305e0
}

func AccountUpdatePersonalChannel(ctx context.Context, m Requester, i AccountUpdatePersonalChannelRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AccountToggleSponsoredMessagesRequestPredict struct {
	Enabled bool `tl:"enabled"`
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
	return 0x06dd654c
}

func AccountGetReactionsNotifySettings(ctx context.Context, m Requester, i AccountGetReactionsNotifySettingsRequestPredict) (ReactionsNotifySettings, error) {
	var res ReactionsNotifySettings
	return res, request(ctx, m, &i, &res)
}

type AccountSetReactionsNotifySettingsRequestPredict struct {
	Settings ReactionsNotifySettings `tl:"settings"`
}

func (*AccountSetReactionsNotifySettingsRequestPredict) CRC() uint32 {
	return 0x316ce548
}

func AccountSetReactionsNotifySettings(ctx context.Context, m Requester, i AccountSetReactionsNotifySettingsRequestPredict) (ReactionsNotifySettings, error) {
	var res ReactionsNotifySettings
	return res, request(ctx, m, &i, &res)
}

type AuthSendCodeRequestPredict struct {
	PhoneNumber string       `tl:"phone_number"`
	APIID       int32        `tl:"api_id"`
	APIHash     string       `tl:"api_hash"`
	Settings    CodeSettings `tl:"settings"`
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
	NoJoinedNotifications bool     `tl:"no_joined_notifications,omitempty:flags:0,implicit"`
	PhoneNumber           string   `tl:"phone_number"`
	PhoneCodeHash         string   `tl:"phone_code_hash"`
	FirstName             string   `tl:"first_name"`
	LastName              string   `tl:"last_name"`
}

func (*AuthSignUpRequestPredict) CRC() uint32 {
	return 0xaac7b717
}

func AuthSignUp(ctx context.Context, m Requester, i AuthSignUpRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthSignInRequestPredict struct {
	_                 struct{}          `tl:"flags,bitflag"`
	PhoneNumber       string            `tl:"phone_number"`
	PhoneCodeHash     string            `tl:"phone_code_hash"`
	PhoneCode         *string           `tl:"phone_code,omitempty:flags:0"`
	EmailVerification EmailVerification `tl:"email_verification,omitempty:flags:1"`
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
	DcID int32 `tl:"dc_id"`
}

func (*AuthExportAuthorizationRequestPredict) CRC() uint32 {
	return 0xe5bfffcd
}

func AuthExportAuthorization(ctx context.Context, m Requester, i AuthExportAuthorizationRequestPredict) (AuthExportedAuthorization, error) {
	var res AuthExportedAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthImportAuthorizationRequestPredict struct {
	ID    int64  `tl:"id"`
	Bytes []byte `tl:"bytes"`
}

func (*AuthImportAuthorizationRequestPredict) CRC() uint32 {
	return 0xa57a7dad
}

func AuthImportAuthorization(ctx context.Context, m Requester, i AuthImportAuthorizationRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthBindTempAuthKeyRequestPredict struct {
	PermAuthKeyID    int64  `tl:"perm_auth_key_id"`
	Nonce            int64  `tl:"nonce"`
	ExpiresAt        int32  `tl:"expires_at"`
	EncryptedMessage []byte `tl:"encrypted_message"`
}

func (*AuthBindTempAuthKeyRequestPredict) CRC() uint32 {
	return 0xcdd42a05
}

func AuthBindTempAuthKey(ctx context.Context, m Requester, i AuthBindTempAuthKeyRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthImportBotAuthorizationRequestPredict struct {
	Flags        int32  `tl:"flags"`
	APIID        int32  `tl:"api_id"`
	APIHash      string `tl:"api_hash"`
	BotAuthToken string `tl:"bot_auth_token"`
}

func (*AuthImportBotAuthorizationRequestPredict) CRC() uint32 {
	return 0x67a3ff2c
}

func AuthImportBotAuthorization(ctx context.Context, m Requester, i AuthImportBotAuthorizationRequestPredict) (AuthAuthorization, error) {
	var res AuthAuthorization
	return res, request(ctx, m, &i, &res)
}

type AuthCheckPasswordRequestPredict struct {
	Password InputCheckPasswordSRP `tl:"password"`
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
	_           struct{}                     `tl:"flags,bitflag"`
	Code        string                       `tl:"code"`
	NewSettings AccountPasswordInputSettings `tl:"new_settings,omitempty:flags:0"`
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
	PhoneNumber   string   `tl:"phone_number"`
	PhoneCodeHash string   `tl:"phone_code_hash"`
	Reason        *string  `tl:"reason,omitempty:flags:0"`
}

func (*AuthResendCodeRequestPredict) CRC() uint32 {
	return 0xcae47523
}

func AuthResendCode(ctx context.Context, m Requester, i AuthResendCodeRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AuthCancelCodeRequestPredict struct {
	PhoneNumber   string `tl:"phone_number"`
	PhoneCodeHash string `tl:"phone_code_hash"`
}

func (*AuthCancelCodeRequestPredict) CRC() uint32 {
	return 0x1f040578
}

func AuthCancelCode(ctx context.Context, m Requester, i AuthCancelCodeRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthDropTempAuthKeysRequestPredict struct {
	ExceptAuthKeys []int64 `tl:"except_auth_keys"`
}

func (*AuthDropTempAuthKeysRequestPredict) CRC() uint32 {
	return 0x8e48a188
}

func AuthDropTempAuthKeys(ctx context.Context, m Requester, i AuthDropTempAuthKeysRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthExportLoginTokenRequestPredict struct {
	APIID     int32   `tl:"api_id"`
	APIHash   string  `tl:"api_hash"`
	ExceptIds []int64 `tl:"except_ids"`
}

func (*AuthExportLoginTokenRequestPredict) CRC() uint32 {
	return 0xb7e085fe
}

func AuthExportLoginToken(ctx context.Context, m Requester, i AuthExportLoginTokenRequestPredict) (AuthLoginToken, error) {
	var res AuthLoginToken
	return res, request(ctx, m, &i, &res)
}

type AuthImportLoginTokenRequestPredict struct {
	Token []byte `tl:"token"`
}

func (*AuthImportLoginTokenRequestPredict) CRC() uint32 {
	return 0x95ac5ce4
}

func AuthImportLoginToken(ctx context.Context, m Requester, i AuthImportLoginTokenRequestPredict) (AuthLoginToken, error) {
	var res AuthLoginToken
	return res, request(ctx, m, &i, &res)
}

type AuthAcceptLoginTokenRequestPredict struct {
	Token []byte `tl:"token"`
}

func (*AuthAcceptLoginTokenRequestPredict) CRC() uint32 {
	return 0xe894ad4d
}

func AuthAcceptLoginToken(ctx context.Context, m Requester, i AuthAcceptLoginTokenRequestPredict) (Authorization, error) {
	var res Authorization
	return res, request(ctx, m, &i, &res)
}

type AuthCheckRecoveryPasswordRequestPredict struct {
	Code string `tl:"code"`
}

func (*AuthCheckRecoveryPasswordRequestPredict) CRC() uint32 {
	return 0x0d36bf79
}

func AuthCheckRecoveryPassword(ctx context.Context, m Requester, i AuthCheckRecoveryPasswordRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthImportWebTokenAuthorizationRequestPredict struct {
	APIID        int32  `tl:"api_id"`
	APIHash      string `tl:"api_hash"`
	WebAuthToken string `tl:"web_auth_token"`
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
	PhoneNumber        string   `tl:"phone_number"`
	PhoneCodeHash      string   `tl:"phone_code_hash"`
	SafetyNetToken     *string  `tl:"safety_net_token,omitempty:flags:0"`
	PlayIntegrityToken *string  `tl:"play_integrity_token,omitempty:flags:2"`
	IosPushSecret      *string  `tl:"ios_push_secret,omitempty:flags:1"`
}

func (*AuthRequestFirebaseSmsRequestPredict) CRC() uint32 {
	return 0x8e39261e
}

func AuthRequestFirebaseSms(ctx context.Context, m Requester, i AuthRequestFirebaseSmsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type AuthResetLoginEmailRequestPredict struct {
	PhoneNumber   string `tl:"phone_number"`
	PhoneCodeHash string `tl:"phone_code_hash"`
}

func (*AuthResetLoginEmailRequestPredict) CRC() uint32 {
	return 0x7e960193
}

func AuthResetLoginEmail(ctx context.Context, m Requester, i AuthResetLoginEmailRequestPredict) (AuthSentCode, error) {
	var res AuthSentCode
	return res, request(ctx, m, &i, &res)
}

type AuthReportMissingCodeRequestPredict struct {
	PhoneNumber   string `tl:"phone_number"`
	PhoneCodeHash string `tl:"phone_code_hash"`
	Mnc           string `tl:"mnc"`
}

func (*AuthReportMissingCodeRequestPredict) CRC() uint32 {
	return 0xcb9deff6
}

func AuthReportMissingCode(ctx context.Context, m Requester, i AuthReportMissingCodeRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSendCustomRequestRequestPredict struct {
	CustomMethod string   `tl:"custom_method"`
	Params       DataJSON `tl:"params"`
}

func (*BotsSendCustomRequestRequestPredict) CRC() uint32 {
	return 0xaa2769ed
}

func BotsSendCustomRequest(ctx context.Context, m Requester, i BotsSendCustomRequestRequestPredict) (DataJSON, error) {
	var res DataJSON
	return res, request(ctx, m, &i, &res)
}

type BotsAnswerWebhookJSONQueryRequestPredict struct {
	QueryID int64    `tl:"query_id"`
	Data    DataJSON `tl:"data"`
}

func (*BotsAnswerWebhookJSONQueryRequestPredict) CRC() uint32 {
	return 0xe6213f4d
}

func BotsAnswerWebhookJSONQuery(ctx context.Context, m Requester, i BotsAnswerWebhookJSONQueryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotCommandsRequestPredict struct {
	Scope    BotCommandScope `tl:"scope"`
	LangCode string          `tl:"lang_code"`
	Commands []BotCommand    `tl:"commands"`
}

func (*BotsSetBotCommandsRequestPredict) CRC() uint32 {
	return 0x0517165a
}

func BotsSetBotCommands(ctx context.Context, m Requester, i BotsSetBotCommandsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsResetBotCommandsRequestPredict struct {
	Scope    BotCommandScope `tl:"scope"`
	LangCode string          `tl:"lang_code"`
}

func (*BotsResetBotCommandsRequestPredict) CRC() uint32 {
	return 0x3d8de0f9
}

func BotsResetBotCommands(ctx context.Context, m Requester, i BotsResetBotCommandsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotCommandsRequestPredict struct {
	Scope    BotCommandScope `tl:"scope"`
	LangCode string          `tl:"lang_code"`
}

func (*BotsGetBotCommandsRequestPredict) CRC() uint32 {
	return 0xe34c0dd6
}

func BotsGetBotCommands(ctx context.Context, m Requester, i BotsGetBotCommandsRequestPredict) ([]BotCommand, error) {
	var res []BotCommand
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotMenuButtonRequestPredict struct {
	UserID InputUser     `tl:"user_id"`
	Button BotMenuButton `tl:"button"`
}

func (*BotsSetBotMenuButtonRequestPredict) CRC() uint32 {
	return 0x4504d54f
}

func BotsSetBotMenuButton(ctx context.Context, m Requester, i BotsSetBotMenuButtonRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetBotMenuButtonRequestPredict struct {
	UserID InputUser `tl:"user_id"`
}

func (*BotsGetBotMenuButtonRequestPredict) CRC() uint32 {
	return 0x9c60eb28
}

func BotsGetBotMenuButton(ctx context.Context, m Requester, i BotsGetBotMenuButtonRequestPredict) (BotMenuButton, error) {
	var res BotMenuButton
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotBroadcastDefaultAdminRightsRequestPredict struct {
	AdminRights ChatAdminRights `tl:"admin_rights"`
}

func (*BotsSetBotBroadcastDefaultAdminRightsRequestPredict) CRC() uint32 {
	return 0x788464e1
}

func BotsSetBotBroadcastDefaultAdminRights(ctx context.Context, m Requester, i BotsSetBotBroadcastDefaultAdminRightsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsSetBotGroupDefaultAdminRightsRequestPredict struct {
	AdminRights ChatAdminRights `tl:"admin_rights"`
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
	Bot         InputUser `tl:"bot,omitempty:flags:2"`
	LangCode    string    `tl:"lang_code"`
	Name        *string   `tl:"name,omitempty:flags:3"`
	About       *string   `tl:"about,omitempty:flags:0"`
	Description *string   `tl:"description,omitempty:flags:1"`
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
	Bot      InputUser `tl:"bot,omitempty:flags:0"`
	LangCode string    `tl:"lang_code"`
}

func (*BotsGetBotInfoRequestPredict) CRC() uint32 {
	return 0xdcd914fd
}

func BotsGetBotInfo(ctx context.Context, m Requester, i BotsGetBotInfoRequestPredict) (BotsBotInfo, error) {
	var res BotsBotInfo
	return res, request(ctx, m, &i, &res)
}

type BotsReorderUsernamesRequestPredict struct {
	Bot   InputUser `tl:"bot"`
	Order []string  `tl:"order"`
}

func (*BotsReorderUsernamesRequestPredict) CRC() uint32 {
	return 0x9709b1c2
}

func BotsReorderUsernames(ctx context.Context, m Requester, i BotsReorderUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsToggleUsernameRequestPredict struct {
	Bot      InputUser `tl:"bot"`
	Username string    `tl:"username"`
	Active   bool      `tl:"active"`
}

func (*BotsToggleUsernameRequestPredict) CRC() uint32 {
	return 0x053ca973
}

func BotsToggleUsername(ctx context.Context, m Requester, i BotsToggleUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsCanSendMessageRequestPredict struct {
	Bot InputUser `tl:"bot"`
}

func (*BotsCanSendMessageRequestPredict) CRC() uint32 {
	return 0x1359f4e6
}

func BotsCanSendMessage(ctx context.Context, m Requester, i BotsCanSendMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsAllowSendMessageRequestPredict struct {
	Bot InputUser `tl:"bot"`
}

func (*BotsAllowSendMessageRequestPredict) CRC() uint32 {
	return 0xf132e3ef
}

func BotsAllowSendMessage(ctx context.Context, m Requester, i BotsAllowSendMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type BotsInvokeWebViewCustomMethodRequestPredict struct {
	Bot          InputUser `tl:"bot"`
	CustomMethod string    `tl:"custom_method"`
	Params       DataJSON  `tl:"params"`
}

func (*BotsInvokeWebViewCustomMethodRequestPredict) CRC() uint32 {
	return 0x087fc5e7
}

func BotsInvokeWebViewCustomMethod(ctx context.Context, m Requester, i BotsInvokeWebViewCustomMethodRequestPredict) (DataJSON, error) {
	var res DataJSON
	return res, request(ctx, m, &i, &res)
}

type BotsGetPopularAppBotsRequestPredict struct {
	Offset string `tl:"offset"`
	Limit  int32  `tl:"limit"`
}

func (*BotsGetPopularAppBotsRequestPredict) CRC() uint32 {
	return 0xc2510192
}

func BotsGetPopularAppBots(ctx context.Context, m Requester, i BotsGetPopularAppBotsRequestPredict) (BotsPopularAppBots, error) {
	var res BotsPopularAppBots
	return res, request(ctx, m, &i, &res)
}

type BotsAddPreviewMediaRequestPredict struct {
	Bot      InputUser  `tl:"bot"`
	LangCode string     `tl:"lang_code"`
	Media    InputMedia `tl:"media"`
}

func (*BotsAddPreviewMediaRequestPredict) CRC() uint32 {
	return 0x17aeb75a
}

func BotsAddPreviewMedia(ctx context.Context, m Requester, i BotsAddPreviewMediaRequestPredict) (BotPreviewMedia, error) {
	var res BotPreviewMedia
	return res, request(ctx, m, &i, &res)
}

type BotsEditPreviewMediaRequestPredict struct {
	Bot      InputUser  `tl:"bot"`
	LangCode string     `tl:"lang_code"`
	Media    InputMedia `tl:"media"`
	NewMedia InputMedia `tl:"new_media"`
}

func (*BotsEditPreviewMediaRequestPredict) CRC() uint32 {
	return 0x8525606f
}

func BotsEditPreviewMedia(ctx context.Context, m Requester, i BotsEditPreviewMediaRequestPredict) (BotPreviewMedia, error) {
	var res BotPreviewMedia
	return res, request(ctx, m, &i, &res)
}

type BotsDeletePreviewMediaRequestPredict struct {
	Bot      InputUser    `tl:"bot"`
	LangCode string       `tl:"lang_code"`
	Media    []InputMedia `tl:"media"`
}

func (*BotsDeletePreviewMediaRequestPredict) CRC() uint32 {
	return 0x2d0135b3
}

func BotsDeletePreviewMedia(ctx context.Context, m Requester, i BotsDeletePreviewMediaRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsReorderPreviewMediasRequestPredict struct {
	Bot      InputUser    `tl:"bot"`
	LangCode string       `tl:"lang_code"`
	Order    []InputMedia `tl:"order"`
}

func (*BotsReorderPreviewMediasRequestPredict) CRC() uint32 {
	return 0xb627f3aa
}

func BotsReorderPreviewMedias(ctx context.Context, m Requester, i BotsReorderPreviewMediasRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type BotsGetPreviewInfoRequestPredict struct {
	Bot      InputUser `tl:"bot"`
	LangCode string    `tl:"lang_code"`
}

func (*BotsGetPreviewInfoRequestPredict) CRC() uint32 {
	return 0x423ab3ad
}

func BotsGetPreviewInfo(ctx context.Context, m Requester, i BotsGetPreviewInfoRequestPredict) (BotsPreviewInfo, error) {
	var res BotsPreviewInfo
	return res, request(ctx, m, &i, &res)
}

type BotsGetPreviewMediasRequestPredict struct {
	Bot InputUser `tl:"bot"`
}

func (*BotsGetPreviewMediasRequestPredict) CRC() uint32 {
	return 0xa2a5594d
}

func BotsGetPreviewMedias(ctx context.Context, m Requester, i BotsGetPreviewMediasRequestPredict) ([]BotPreviewMedia, error) {
	var res []BotPreviewMedia
	return res, request(ctx, m, &i, &res)
}

type ChannelsReadHistoryRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	MaxID   int32        `tl:"max_id"`
}

func (*ChannelsReadHistoryRequestPredict) CRC() uint32 {
	return 0xcc104937
}

func ChannelsReadHistory(ctx context.Context, m Requester, i ChannelsReadHistoryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteMessagesRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	ID      []int32      `tl:"id"`
}

func (*ChannelsDeleteMessagesRequestPredict) CRC() uint32 {
	return 0x84c1fd4e
}

func ChannelsDeleteMessages(ctx context.Context, m Requester, i ChannelsDeleteMessagesRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportSpamRequestPredict struct {
	Channel     InputChannel `tl:"channel"`
	Participant InputPeer    `tl:"participant"`
	ID          []int32      `tl:"id"`
}

func (*ChannelsReportSpamRequestPredict) CRC() uint32 {
	return 0xf44a8315
}

func ChannelsReportSpam(ctx context.Context, m Requester, i ChannelsReportSpamRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetMessagesRequestPredict struct {
	Channel InputChannel   `tl:"channel"`
	ID      []InputMessage `tl:"id"`
}

func (*ChannelsGetMessagesRequestPredict) CRC() uint32 {
	return 0xad8c9a23
}

func ChannelsGetMessages(ctx context.Context, m Requester, i ChannelsGetMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetParticipantsRequestPredict struct {
	Channel InputChannel              `tl:"channel"`
	Filter  ChannelParticipantsFilter `tl:"filter"`
	Offset  int32                     `tl:"offset"`
	Limit   int32                     `tl:"limit"`
	Hash    int64                     `tl:"hash"`
}

func (*ChannelsGetParticipantsRequestPredict) CRC() uint32 {
	return 0x77ced9d0
}

func ChannelsGetParticipants(ctx context.Context, m Requester, i ChannelsGetParticipantsRequestPredict) (ChannelsChannelParticipants, error) {
	var res ChannelsChannelParticipants
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetParticipantRequestPredict struct {
	Channel     InputChannel `tl:"channel"`
	Participant InputPeer    `tl:"participant"`
}

func (*ChannelsGetParticipantRequestPredict) CRC() uint32 {
	return 0xa0ab6cc6
}

func ChannelsGetParticipant(ctx context.Context, m Requester, i ChannelsGetParticipantRequestPredict) (ChannelsChannelParticipant, error) {
	var res ChannelsChannelParticipant
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetChannelsRequestPredict struct {
	ID []InputChannel `tl:"id"`
}

func (*ChannelsGetChannelsRequestPredict) CRC() uint32 {
	return 0x0a7f6bbb
}

func ChannelsGetChannels(ctx context.Context, m Requester, i ChannelsGetChannelsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetFullChannelRequestPredict struct {
	Channel InputChannel `tl:"channel"`
}

func (*ChannelsGetFullChannelRequestPredict) CRC() uint32 {
	return 0x08736a09
}

func ChannelsGetFullChannel(ctx context.Context, m Requester, i ChannelsGetFullChannelRequestPredict) (MessagesChatFull, error) {
	var res MessagesChatFull
	return res, request(ctx, m, &i, &res)
}

type ChannelsCreateChannelRequestPredict struct {
	_         struct{}      `tl:"flags,bitflag"`
	Broadcast bool          `tl:"broadcast,omitempty:flags:0,implicit"`
	Megagroup bool          `tl:"megagroup,omitempty:flags:1,implicit"`
	ForImport bool          `tl:"for_import,omitempty:flags:3,implicit"`
	Forum     bool          `tl:"forum,omitempty:flags:5,implicit"`
	Title     string        `tl:"title"`
	About     string        `tl:"about"`
	GeoPoint  InputGeoPoint `tl:"geo_point,omitempty:flags:2"`
	Address   *string       `tl:"address,omitempty:flags:2"`
	TTLPeriod *int32        `tl:"ttl_period,omitempty:flags:4"`
}

func (*ChannelsCreateChannelRequestPredict) CRC() uint32 {
	return 0x91006707
}

func ChannelsCreateChannel(ctx context.Context, m Requester, i ChannelsCreateChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditAdminRequestPredict struct {
	Channel     InputChannel    `tl:"channel"`
	UserID      InputUser       `tl:"user_id"`
	AdminRights ChatAdminRights `tl:"admin_rights"`
	Rank        string          `tl:"rank"`
}

func (*ChannelsEditAdminRequestPredict) CRC() uint32 {
	return 0xd33c8902
}

func ChannelsEditAdmin(ctx context.Context, m Requester, i ChannelsEditAdminRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditTitleRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Title   string       `tl:"title"`
}

func (*ChannelsEditTitleRequestPredict) CRC() uint32 {
	return 0x566decd0
}

func ChannelsEditTitle(ctx context.Context, m Requester, i ChannelsEditTitleRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditPhotoRequestPredict struct {
	Channel InputChannel   `tl:"channel"`
	Photo   InputChatPhoto `tl:"photo"`
}

func (*ChannelsEditPhotoRequestPredict) CRC() uint32 {
	return 0xf12e57c9
}

func ChannelsEditPhoto(ctx context.Context, m Requester, i ChannelsEditPhotoRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsCheckUsernameRequestPredict struct {
	Channel  InputChannel `tl:"channel"`
	Username string       `tl:"username"`
}

func (*ChannelsCheckUsernameRequestPredict) CRC() uint32 {
	return 0x10e6bd2c
}

func ChannelsCheckUsername(ctx context.Context, m Requester, i ChannelsCheckUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateUsernameRequestPredict struct {
	Channel  InputChannel `tl:"channel"`
	Username string       `tl:"username"`
}

func (*ChannelsUpdateUsernameRequestPredict) CRC() uint32 {
	return 0x3514b3de
}

func ChannelsUpdateUsername(ctx context.Context, m Requester, i ChannelsUpdateUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsJoinChannelRequestPredict struct {
	Channel InputChannel `tl:"channel"`
}

func (*ChannelsJoinChannelRequestPredict) CRC() uint32 {
	return 0x24b524c5
}

func ChannelsJoinChannel(ctx context.Context, m Requester, i ChannelsJoinChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsLeaveChannelRequestPredict struct {
	Channel InputChannel `tl:"channel"`
}

func (*ChannelsLeaveChannelRequestPredict) CRC() uint32 {
	return 0xf836aa95
}

func ChannelsLeaveChannel(ctx context.Context, m Requester, i ChannelsLeaveChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsInviteToChannelRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Users   []InputUser  `tl:"users"`
}

func (*ChannelsInviteToChannelRequestPredict) CRC() uint32 {
	return 0xc9e33d54
}

func ChannelsInviteToChannel(ctx context.Context, m Requester, i ChannelsInviteToChannelRequestPredict) (MessagesInvitedUsers, error) {
	var res MessagesInvitedUsers
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteChannelRequestPredict struct {
	Channel InputChannel `tl:"channel"`
}

func (*ChannelsDeleteChannelRequestPredict) CRC() uint32 {
	return 0xc0111fe3
}

func ChannelsDeleteChannel(ctx context.Context, m Requester, i ChannelsDeleteChannelRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsExportMessageLinkRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Grouped bool         `tl:"grouped,omitempty:flags:0,implicit"`
	Thread  bool         `tl:"thread,omitempty:flags:1,implicit"`
	Channel InputChannel `tl:"channel"`
	ID      int32        `tl:"id"`
}

func (*ChannelsExportMessageLinkRequestPredict) CRC() uint32 {
	return 0xe63fadeb
}

func ChannelsExportMessageLink(ctx context.Context, m Requester, i ChannelsExportMessageLinkRequestPredict) (ExportedMessageLink, error) {
	var res ExportedMessageLink
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleSignaturesRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
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
	ByLocation  bool     `tl:"by_location,omitempty:flags:0,implicit"`
	CheckLimit  bool     `tl:"check_limit,omitempty:flags:1,implicit"`
	ForPersonal bool     `tl:"for_personal,omitempty:flags:2,implicit"`
}

func (*ChannelsGetAdminedPublicChannelsRequestPredict) CRC() uint32 {
	return 0xf8b036af
}

func ChannelsGetAdminedPublicChannels(ctx context.Context, m Requester, i ChannelsGetAdminedPublicChannelsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditBannedRequestPredict struct {
	Channel      InputChannel     `tl:"channel"`
	Participant  InputPeer        `tl:"participant"`
	BannedRights ChatBannedRights `tl:"banned_rights"`
}

func (*ChannelsEditBannedRequestPredict) CRC() uint32 {
	return 0x96e6cd81
}

func ChannelsEditBanned(ctx context.Context, m Requester, i ChannelsEditBannedRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetAdminLogRequestPredict struct {
	_            struct{}                    `tl:"flags,bitflag"`
	Channel      InputChannel                `tl:"channel"`
	Q            string                      `tl:"q"`
	EventsFilter ChannelAdminLogEventsFilter `tl:"events_filter,omitempty:flags:0"`
	Admins       []InputUser                 `tl:"admins,omitempty:flags:1"`
	MaxID        int64                       `tl:"max_id"`
	MinID        int64                       `tl:"min_id"`
	Limit        int32                       `tl:"limit"`
}

func (*ChannelsGetAdminLogRequestPredict) CRC() uint32 {
	return 0x33ddf480
}

func ChannelsGetAdminLog(ctx context.Context, m Requester, i ChannelsGetAdminLogRequestPredict) (ChannelsAdminLogResults, error) {
	var res ChannelsAdminLogResults
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetStickersRequestPredict struct {
	Channel    InputChannel    `tl:"channel"`
	Stickerset InputStickerSet `tl:"stickerset"`
}

func (*ChannelsSetStickersRequestPredict) CRC() uint32 {
	return 0xea8ca4f9
}

func ChannelsSetStickers(ctx context.Context, m Requester, i ChannelsSetStickersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsReadMessageContentsRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	ID      []int32      `tl:"id"`
}

func (*ChannelsReadMessageContentsRequestPredict) CRC() uint32 {
	return 0xeab5dc38
}

func ChannelsReadMessageContents(ctx context.Context, m Requester, i ChannelsReadMessageContentsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteHistoryRequestPredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	ForEveryone bool         `tl:"for_everyone,omitempty:flags:0,implicit"`
	Channel     InputChannel `tl:"channel"`
	MaxID       int32        `tl:"max_id"`
}

func (*ChannelsDeleteHistoryRequestPredict) CRC() uint32 {
	return 0x9baa9647
}

func ChannelsDeleteHistory(ctx context.Context, m Requester, i ChannelsDeleteHistoryRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsTogglePreHistoryHiddenRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
}

func (*ChannelsTogglePreHistoryHiddenRequestPredict) CRC() uint32 {
	return 0xeabbb94c
}

func ChannelsTogglePreHistoryHidden(ctx context.Context, m Requester, i ChannelsTogglePreHistoryHiddenRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetLeftChannelsRequestPredict struct {
	Offset int32 `tl:"offset"`
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
	Broadcast InputChannel `tl:"broadcast"`
	Group     InputChannel `tl:"group"`
}

func (*ChannelsSetDiscussionGroupRequestPredict) CRC() uint32 {
	return 0x40582bb2
}

func ChannelsSetDiscussionGroup(ctx context.Context, m Requester, i ChannelsSetDiscussionGroupRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditCreatorRequestPredict struct {
	Channel  InputChannel          `tl:"channel"`
	UserID   InputUser             `tl:"user_id"`
	Password InputCheckPasswordSRP `tl:"password"`
}

func (*ChannelsEditCreatorRequestPredict) CRC() uint32 {
	return 0x8f38cd1f
}

func ChannelsEditCreator(ctx context.Context, m Requester, i ChannelsEditCreatorRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditLocationRequestPredict struct {
	Channel  InputChannel  `tl:"channel"`
	GeoPoint InputGeoPoint `tl:"geo_point"`
	Address  string        `tl:"address"`
}

func (*ChannelsEditLocationRequestPredict) CRC() uint32 {
	return 0x58e63f6d
}

func ChannelsEditLocation(ctx context.Context, m Requester, i ChannelsEditLocationRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleSlowModeRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Seconds int32        `tl:"seconds"`
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
	Channel InputChannel `tl:"channel"`
}

func (*ChannelsConvertToGigagroupRequestPredict) CRC() uint32 {
	return 0x0b290c69
}

func ChannelsConvertToGigagroup(ctx context.Context, m Requester, i ChannelsConvertToGigagroupRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsViewSponsoredMessageRequestPredict struct {
	Channel  InputChannel `tl:"channel"`
	RandomID []byte       `tl:"random_id"`
}

func (*ChannelsViewSponsoredMessageRequestPredict) CRC() uint32 {
	return 0xbeaedb94
}

func ChannelsViewSponsoredMessage(ctx context.Context, m Requester, i ChannelsViewSponsoredMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetSponsoredMessagesRequestPredict struct {
	Channel InputChannel `tl:"channel"`
}

func (*ChannelsGetSponsoredMessagesRequestPredict) CRC() uint32 {
	return 0xec210fbf
}

func ChannelsGetSponsoredMessages(ctx context.Context, m Requester, i ChannelsGetSponsoredMessagesRequestPredict) (MessagesSponsoredMessages, error) {
	var res MessagesSponsoredMessages
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetSendAsRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*ChannelsGetSendAsRequestPredict) CRC() uint32 {
	return 0x0dc770ee
}

func ChannelsGetSendAs(ctx context.Context, m Requester, i ChannelsGetSendAsRequestPredict) (ChannelsSendAsPeers, error) {
	var res ChannelsSendAsPeers
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteParticipantHistoryRequestPredict struct {
	Channel     InputChannel `tl:"channel"`
	Participant InputPeer    `tl:"participant"`
}

func (*ChannelsDeleteParticipantHistoryRequestPredict) CRC() uint32 {
	return 0x367544db
}

func ChannelsDeleteParticipantHistory(ctx context.Context, m Requester, i ChannelsDeleteParticipantHistoryRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleJoinToSendRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
}

func (*ChannelsToggleJoinToSendRequestPredict) CRC() uint32 {
	return 0xe4cb9580
}

func ChannelsToggleJoinToSend(ctx context.Context, m Requester, i ChannelsToggleJoinToSendRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleJoinRequestRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
}

func (*ChannelsToggleJoinRequestRequestPredict) CRC() uint32 {
	return 0x4c2985b6
}

func ChannelsToggleJoinRequest(ctx context.Context, m Requester, i ChannelsToggleJoinRequestRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsReorderUsernamesRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Order   []string     `tl:"order"`
}

func (*ChannelsReorderUsernamesRequestPredict) CRC() uint32 {
	return 0xb45ced1d
}

func ChannelsReorderUsernames(ctx context.Context, m Requester, i ChannelsReorderUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleUsernameRequestPredict struct {
	Channel  InputChannel `tl:"channel"`
	Username string       `tl:"username"`
	Active   bool         `tl:"active"`
}

func (*ChannelsToggleUsernameRequestPredict) CRC() uint32 {
	return 0x50f24105
}

func ChannelsToggleUsername(ctx context.Context, m Requester, i ChannelsToggleUsernameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeactivateAllUsernamesRequestPredict struct {
	Channel InputChannel `tl:"channel"`
}

func (*ChannelsDeactivateAllUsernamesRequestPredict) CRC() uint32 {
	return 0x0a245dd3
}

func ChannelsDeactivateAllUsernames(ctx context.Context, m Requester, i ChannelsDeactivateAllUsernamesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleForumRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
}

func (*ChannelsToggleForumRequestPredict) CRC() uint32 {
	return 0xa4298b29
}

func ChannelsToggleForum(ctx context.Context, m Requester, i ChannelsToggleForumRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsCreateForumTopicRequestPredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	Channel     InputChannel `tl:"channel"`
	Title       string       `tl:"title"`
	IconColor   *int32       `tl:"icon_color,omitempty:flags:0"`
	IconEmojiID *int64       `tl:"icon_emoji_id,omitempty:flags:3"`
	RandomID    int64        `tl:"random_id"`
	SendAs      InputPeer    `tl:"send_as,omitempty:flags:2"`
}

func (*ChannelsCreateForumTopicRequestPredict) CRC() uint32 {
	return 0xf40c0224
}

func ChannelsCreateForumTopic(ctx context.Context, m Requester, i ChannelsCreateForumTopicRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetForumTopicsRequestPredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	Channel     InputChannel `tl:"channel"`
	Q           *string      `tl:"q,omitempty:flags:0"`
	OffsetDate  int32        `tl:"offset_date"`
	OffsetID    int32        `tl:"offset_id"`
	OffsetTopic int32        `tl:"offset_topic"`
	Limit       int32        `tl:"limit"`
}

func (*ChannelsGetForumTopicsRequestPredict) CRC() uint32 {
	return 0x0de560d1
}

func ChannelsGetForumTopics(ctx context.Context, m Requester, i ChannelsGetForumTopicsRequestPredict) (MessagesForumTopics, error) {
	var res MessagesForumTopics
	return res, request(ctx, m, &i, &res)
}

type ChannelsGetForumTopicsByIDRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Topics  []int32      `tl:"topics"`
}

func (*ChannelsGetForumTopicsByIDRequestPredict) CRC() uint32 {
	return 0xb0831eb9
}

func ChannelsGetForumTopicsByID(ctx context.Context, m Requester, i ChannelsGetForumTopicsByIDRequestPredict) (MessagesForumTopics, error) {
	var res MessagesForumTopics
	return res, request(ctx, m, &i, &res)
}

type ChannelsEditForumTopicRequestPredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	Channel     InputChannel `tl:"channel"`
	TopicID     int32        `tl:"topic_id"`
	Title       *string      `tl:"title,omitempty:flags:0"`
	IconEmojiID *int64       `tl:"icon_emoji_id,omitempty:flags:1"`
	Closed      *bool        `tl:"closed,omitempty:flags:2"`
	Hidden      *bool        `tl:"hidden,omitempty:flags:3"`
}

func (*ChannelsEditForumTopicRequestPredict) CRC() uint32 {
	return 0xf4dfa185
}

func ChannelsEditForumTopic(ctx context.Context, m Requester, i ChannelsEditForumTopicRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdatePinnedForumTopicRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	TopicID int32        `tl:"topic_id"`
	Pinned  bool         `tl:"pinned"`
}

func (*ChannelsUpdatePinnedForumTopicRequestPredict) CRC() uint32 {
	return 0x6c2d9026
}

func ChannelsUpdatePinnedForumTopic(ctx context.Context, m Requester, i ChannelsUpdatePinnedForumTopicRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsDeleteTopicHistoryRequestPredict struct {
	Channel  InputChannel `tl:"channel"`
	TopMsgID int32        `tl:"top_msg_id"`
}

func (*ChannelsDeleteTopicHistoryRequestPredict) CRC() uint32 {
	return 0x34435f2d
}

func ChannelsDeleteTopicHistory(ctx context.Context, m Requester, i ChannelsDeleteTopicHistoryRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type ChannelsReorderPinnedForumTopicsRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Force   bool         `tl:"force,omitempty:flags:0,implicit"`
	Channel InputChannel `tl:"channel"`
	Order   []int32      `tl:"order"`
}

func (*ChannelsReorderPinnedForumTopicsRequestPredict) CRC() uint32 {
	return 0x2950a18f
}

func ChannelsReorderPinnedForumTopics(ctx context.Context, m Requester, i ChannelsReorderPinnedForumTopicsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleAntiSpamRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
}

func (*ChannelsToggleAntiSpamRequestPredict) CRC() uint32 {
	return 0x68f3e4eb
}

func ChannelsToggleAntiSpam(ctx context.Context, m Requester, i ChannelsToggleAntiSpamRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportAntiSpamFalsePositiveRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	MsgID   int32        `tl:"msg_id"`
}

func (*ChannelsReportAntiSpamFalsePositiveRequestPredict) CRC() uint32 {
	return 0xa850a693
}

func ChannelsReportAntiSpamFalsePositive(ctx context.Context, m Requester, i ChannelsReportAntiSpamFalsePositiveRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleParticipantsHiddenRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
}

func (*ChannelsToggleParticipantsHiddenRequestPredict) CRC() uint32 {
	return 0x6a6e7854
}

func ChannelsToggleParticipantsHidden(ctx context.Context, m Requester, i ChannelsToggleParticipantsHiddenRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsClickSponsoredMessageRequestPredict struct {
	Channel  InputChannel `tl:"channel"`
	RandomID []byte       `tl:"random_id"`
}

func (*ChannelsClickSponsoredMessageRequestPredict) CRC() uint32 {
	return 0x18afbc93
}

func ChannelsClickSponsoredMessage(ctx context.Context, m Requester, i ChannelsClickSponsoredMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateColorRequestPredict struct {
	_                 struct{}     `tl:"flags,bitflag"`
	ForProfile        bool         `tl:"for_profile,omitempty:flags:1,implicit"`
	Channel           InputChannel `tl:"channel"`
	Color             *int32       `tl:"color,omitempty:flags:2"`
	BackgroundEmojiID *int64       `tl:"background_emoji_id,omitempty:flags:0"`
}

func (*ChannelsUpdateColorRequestPredict) CRC() uint32 {
	return 0xd8aa3671
}

func ChannelsUpdateColor(ctx context.Context, m Requester, i ChannelsUpdateColorRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsToggleViewForumAsMessagesRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Enabled bool         `tl:"enabled"`
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
	Channel InputChannel `tl:"channel,omitempty:flags:0"`
}

func (*ChannelsGetChannelRecommendationsRequestPredict) CRC() uint32 {
	return 0x25a71742
}

func ChannelsGetChannelRecommendations(ctx context.Context, m Requester, i ChannelsGetChannelRecommendationsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type ChannelsUpdateEmojiStatusRequestPredict struct {
	Channel     InputChannel `tl:"channel"`
	EmojiStatus EmojiStatus  `tl:"emoji_status"`
}

func (*ChannelsUpdateEmojiStatusRequestPredict) CRC() uint32 {
	return 0xf0d3e6a8
}

func ChannelsUpdateEmojiStatus(ctx context.Context, m Requester, i ChannelsUpdateEmojiStatusRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetBoostsToUnblockRestrictionsRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Boosts  int32        `tl:"boosts"`
}

func (*ChannelsSetBoostsToUnblockRestrictionsRequestPredict) CRC() uint32 {
	return 0xad399cee
}

func ChannelsSetBoostsToUnblockRestrictions(ctx context.Context, m Requester, i ChannelsSetBoostsToUnblockRestrictionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsSetEmojiStickersRequestPredict struct {
	Channel    InputChannel    `tl:"channel"`
	Stickerset InputStickerSet `tl:"stickerset"`
}

func (*ChannelsSetEmojiStickersRequestPredict) CRC() uint32 {
	return 0x3cd930b7
}

func ChannelsSetEmojiStickers(ctx context.Context, m Requester, i ChannelsSetEmojiStickersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChannelsReportSponsoredMessageRequestPredict struct {
	Channel  InputChannel `tl:"channel"`
	RandomID []byte       `tl:"random_id"`
	Option   []byte       `tl:"option"`
}

func (*ChannelsReportSponsoredMessageRequestPredict) CRC() uint32 {
	return 0xaf8ff6b9
}

func ChannelsReportSponsoredMessage(ctx context.Context, m Requester, i ChannelsReportSponsoredMessageRequestPredict) (ChannelsSponsoredMessageReportResult, error) {
	var res ChannelsSponsoredMessageReportResult
	return res, request(ctx, m, &i, &res)
}

type ChannelsRestrictSponsoredMessagesRequestPredict struct {
	Channel    InputChannel `tl:"channel"`
	Restricted bool         `tl:"restricted"`
}

func (*ChannelsRestrictSponsoredMessagesRequestPredict) CRC() uint32 {
	return 0x9ae91519
}

func ChannelsRestrictSponsoredMessages(ctx context.Context, m Requester, i ChannelsRestrictSponsoredMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChannelsSearchPostsRequestPredict struct {
	Hashtag    string    `tl:"hashtag"`
	OffsetRate int32     `tl:"offset_rate"`
	OffsetPeer InputPeer `tl:"offset_peer"`
	OffsetID   int32     `tl:"offset_id"`
	Limit      int32     `tl:"limit"`
}

func (*ChannelsSearchPostsRequestPredict) CRC() uint32 {
	return 0xd19f987b
}

func ChannelsSearchPosts(ctx context.Context, m Requester, i ChannelsSearchPostsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type ChatlistsExportChatlistInviteRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
	Title    string        `tl:"title"`
	Peers    []InputPeer   `tl:"peers"`
}

func (*ChatlistsExportChatlistInviteRequestPredict) CRC() uint32 {
	return 0x8472478e
}

func ChatlistsExportChatlistInvite(ctx context.Context, m Requester, i ChatlistsExportChatlistInviteRequestPredict) (ChatlistsExportedChatlistInvite, error) {
	var res ChatlistsExportedChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsDeleteExportedInviteRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
	Slug     string        `tl:"slug"`
}

func (*ChatlistsDeleteExportedInviteRequestPredict) CRC() uint32 {
	return 0x719c5c5e
}

func ChatlistsDeleteExportedInvite(ctx context.Context, m Requester, i ChatlistsDeleteExportedInviteRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChatlistsEditExportedInviteRequestPredict struct {
	_        struct{}      `tl:"flags,bitflag"`
	Chatlist InputChatlist `tl:"chatlist"`
	Slug     string        `tl:"slug"`
	Title    *string       `tl:"title,omitempty:flags:1"`
	Peers    []InputPeer   `tl:"peers,omitempty:flags:2"`
}

func (*ChatlistsEditExportedInviteRequestPredict) CRC() uint32 {
	return 0x653db63d
}

func ChatlistsEditExportedInvite(ctx context.Context, m Requester, i ChatlistsEditExportedInviteRequestPredict) (ExportedChatlistInvite, error) {
	var res ExportedChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetExportedInvitesRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
}

func (*ChatlistsGetExportedInvitesRequestPredict) CRC() uint32 {
	return 0xce03da83
}

func ChatlistsGetExportedInvites(ctx context.Context, m Requester, i ChatlistsGetExportedInvitesRequestPredict) (ChatlistsExportedInvites, error) {
	var res ChatlistsExportedInvites
	return res, request(ctx, m, &i, &res)
}

type ChatlistsCheckChatlistInviteRequestPredict struct {
	Slug string `tl:"slug"`
}

func (*ChatlistsCheckChatlistInviteRequestPredict) CRC() uint32 {
	return 0x41c10fff
}

func ChatlistsCheckChatlistInvite(ctx context.Context, m Requester, i ChatlistsCheckChatlistInviteRequestPredict) (ChatlistsChatlistInvite, error) {
	var res ChatlistsChatlistInvite
	return res, request(ctx, m, &i, &res)
}

type ChatlistsJoinChatlistInviteRequestPredict struct {
	Slug  string      `tl:"slug"`
	Peers []InputPeer `tl:"peers"`
}

func (*ChatlistsJoinChatlistInviteRequestPredict) CRC() uint32 {
	return 0xa6b1e39a
}

func ChatlistsJoinChatlistInvite(ctx context.Context, m Requester, i ChatlistsJoinChatlistInviteRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetChatlistUpdatesRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
}

func (*ChatlistsGetChatlistUpdatesRequestPredict) CRC() uint32 {
	return 0x89419521
}

func ChatlistsGetChatlistUpdates(ctx context.Context, m Requester, i ChatlistsGetChatlistUpdatesRequestPredict) (ChatlistsChatlistUpdates, error) {
	var res ChatlistsChatlistUpdates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsJoinChatlistUpdatesRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
	Peers    []InputPeer   `tl:"peers"`
}

func (*ChatlistsJoinChatlistUpdatesRequestPredict) CRC() uint32 {
	return 0xe089f8f5
}

func ChatlistsJoinChatlistUpdates(ctx context.Context, m Requester, i ChatlistsJoinChatlistUpdatesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ChatlistsHideChatlistUpdatesRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
}

func (*ChatlistsHideChatlistUpdatesRequestPredict) CRC() uint32 {
	return 0x66e486fb
}

func ChatlistsHideChatlistUpdates(ctx context.Context, m Requester, i ChatlistsHideChatlistUpdatesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ChatlistsGetLeaveChatlistSuggestionsRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
}

func (*ChatlistsGetLeaveChatlistSuggestionsRequestPredict) CRC() uint32 {
	return 0xfdbcd714
}

func ChatlistsGetLeaveChatlistSuggestions(ctx context.Context, m Requester, i ChatlistsGetLeaveChatlistSuggestionsRequestPredict) ([]Peer, error) {
	var res []Peer
	return res, request(ctx, m, &i, &res)
}

type ChatlistsLeaveChatlistRequestPredict struct {
	Chatlist InputChatlist `tl:"chatlist"`
	Peers    []InputPeer   `tl:"peers"`
}

func (*ChatlistsLeaveChatlistRequestPredict) CRC() uint32 {
	return 0x74fae13a
}

func ChatlistsLeaveChatlist(ctx context.Context, m Requester, i ChatlistsLeaveChatlistRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsGetContactIDsRequestPredict struct {
	Hash int64 `tl:"hash"`
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
	Hash int64 `tl:"hash"`
}

func (*ContactsGetContactsRequestPredict) CRC() uint32 {
	return 0x5dd69e12
}

func ContactsGetContacts(ctx context.Context, m Requester, i ContactsGetContactsRequestPredict) (ContactsContacts, error) {
	var res ContactsContacts
	return res, request(ctx, m, &i, &res)
}

type ContactsImportContactsRequestPredict struct {
	Contacts []InputContact `tl:"contacts"`
}

func (*ContactsImportContactsRequestPredict) CRC() uint32 {
	return 0x2c800be5
}

func ContactsImportContacts(ctx context.Context, m Requester, i ContactsImportContactsRequestPredict) (ContactsImportedContacts, error) {
	var res ContactsImportedContacts
	return res, request(ctx, m, &i, &res)
}

type ContactsDeleteContactsRequestPredict struct {
	ID []InputUser `tl:"id"`
}

func (*ContactsDeleteContactsRequestPredict) CRC() uint32 {
	return 0x096a0e00
}

func ContactsDeleteContacts(ctx context.Context, m Requester, i ContactsDeleteContactsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsDeleteByPhonesRequestPredict struct {
	Phones []string `tl:"phones"`
}

func (*ContactsDeleteByPhonesRequestPredict) CRC() uint32 {
	return 0x1013fd9e
}

func ContactsDeleteByPhones(ctx context.Context, m Requester, i ContactsDeleteByPhonesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsBlockRequestPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	MyStoriesFrom bool      `tl:"my_stories_from,omitempty:flags:0,implicit"`
	ID            InputPeer `tl:"id"`
}

func (*ContactsBlockRequestPredict) CRC() uint32 {
	return 0x2e2e8734
}

func ContactsBlock(ctx context.Context, m Requester, i ContactsBlockRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsUnblockRequestPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	MyStoriesFrom bool      `tl:"my_stories_from,omitempty:flags:0,implicit"`
	ID            InputPeer `tl:"id"`
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
	MyStoriesFrom bool     `tl:"my_stories_from,omitempty:flags:0,implicit"`
	Offset        int32    `tl:"offset"`
	Limit         int32    `tl:"limit"`
}

func (*ContactsGetBlockedRequestPredict) CRC() uint32 {
	return 0x9a868f80
}

func ContactsGetBlocked(ctx context.Context, m Requester, i ContactsGetBlockedRequestPredict) (ContactsBlocked, error) {
	var res ContactsBlocked
	return res, request(ctx, m, &i, &res)
}

type ContactsSearchRequestPredict struct {
	Q     string `tl:"q"`
	Limit int32  `tl:"limit"`
}

func (*ContactsSearchRequestPredict) CRC() uint32 {
	return 0x11f812d8
}

func ContactsSearch(ctx context.Context, m Requester, i ContactsSearchRequestPredict) (ContactsFound, error) {
	var res ContactsFound
	return res, request(ctx, m, &i, &res)
}

type ContactsResolveUsernameRequestPredict struct {
	Username string `tl:"username"`
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
	Correspondents bool     `tl:"correspondents,omitempty:flags:0,implicit"`
	BotsPm         bool     `tl:"bots_pm,omitempty:flags:1,implicit"`
	BotsInline     bool     `tl:"bots_inline,omitempty:flags:2,implicit"`
	PhoneCalls     bool     `tl:"phone_calls,omitempty:flags:3,implicit"`
	ForwardUsers   bool     `tl:"forward_users,omitempty:flags:4,implicit"`
	ForwardChats   bool     `tl:"forward_chats,omitempty:flags:5,implicit"`
	Groups         bool     `tl:"groups,omitempty:flags:10,implicit"`
	Channels       bool     `tl:"channels,omitempty:flags:15,implicit"`
	BotsApp        bool     `tl:"bots_app,omitempty:flags:16,implicit"`
	Offset         int32    `tl:"offset"`
	Limit          int32    `tl:"limit"`
	Hash           int64    `tl:"hash"`
}

func (*ContactsGetTopPeersRequestPredict) CRC() uint32 {
	return 0x973478b6
}

func ContactsGetTopPeers(ctx context.Context, m Requester, i ContactsGetTopPeersRequestPredict) (ContactsTopPeers, error) {
	var res ContactsTopPeers
	return res, request(ctx, m, &i, &res)
}

type ContactsResetTopPeerRatingRequestPredict struct {
	Category TopPeerCategory `tl:"category"`
	Peer     InputPeer       `tl:"peer"`
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
	Enabled bool `tl:"enabled"`
}

func (*ContactsToggleTopPeersRequestPredict) CRC() uint32 {
	return 0x8514bdda
}

func ContactsToggleTopPeers(ctx context.Context, m Requester, i ContactsToggleTopPeersRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsAddContactRequestPredict struct {
	_                        struct{}  `tl:"flags,bitflag"`
	AddPhonePrivacyException bool      `tl:"add_phone_privacy_exception,omitempty:flags:0,implicit"`
	ID                       InputUser `tl:"id"`
	FirstName                string    `tl:"first_name"`
	LastName                 string    `tl:"last_name"`
	Phone                    string    `tl:"phone"`
}

func (*ContactsAddContactRequestPredict) CRC() uint32 {
	return 0xe8f463d0
}

func ContactsAddContact(ctx context.Context, m Requester, i ContactsAddContactRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsAcceptContactRequestPredict struct {
	ID InputUser `tl:"id"`
}

func (*ContactsAcceptContactRequestPredict) CRC() uint32 {
	return 0xf831a20f
}

func ContactsAcceptContact(ctx context.Context, m Requester, i ContactsAcceptContactRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsGetLocatedRequestPredict struct {
	_           struct{}      `tl:"flags,bitflag"`
	Background  bool          `tl:"background,omitempty:flags:1,implicit"`
	GeoPoint    InputGeoPoint `tl:"geo_point"`
	SelfExpires *int32        `tl:"self_expires,omitempty:flags:0"`
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
	DeleteMessage bool     `tl:"delete_message,omitempty:flags:0,implicit"`
	DeleteHistory bool     `tl:"delete_history,omitempty:flags:1,implicit"`
	ReportSpam    bool     `tl:"report_spam,omitempty:flags:2,implicit"`
	MsgID         int32    `tl:"msg_id"`
}

func (*ContactsBlockFromRepliesRequestPredict) CRC() uint32 {
	return 0x29a8962c
}

func ContactsBlockFromReplies(ctx context.Context, m Requester, i ContactsBlockFromRepliesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type ContactsResolvePhoneRequestPredict struct {
	Phone string `tl:"phone"`
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
	Token string `tl:"token"`
}

func (*ContactsImportContactTokenRequestPredict) CRC() uint32 {
	return 0x13005788
}

func ContactsImportContactToken(ctx context.Context, m Requester, i ContactsImportContactTokenRequestPredict) (User, error) {
	var res User
	return res, request(ctx, m, &i, &res)
}

type ContactsEditCloseFriendsRequestPredict struct {
	ID []int64 `tl:"id"`
}

func (*ContactsEditCloseFriendsRequestPredict) CRC() uint32 {
	return 0xba6705f0
}

func ContactsEditCloseFriends(ctx context.Context, m Requester, i ContactsEditCloseFriendsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type ContactsSetBlockedRequestPredict struct {
	_             struct{}    `tl:"flags,bitflag"`
	MyStoriesFrom bool        `tl:"my_stories_from,omitempty:flags:0,implicit"`
	ID            []InputPeer `tl:"id"`
	Limit         int32       `tl:"limit"`
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
	FolderPeers []InputFolderPeer `tl:"folder_peers"`
}

func (*FoldersEditPeerFoldersRequestPredict) CRC() uint32 {
	return 0x6847d0ab
}

func FoldersEditPeerFolders(ctx context.Context, m Requester, i FoldersEditPeerFoldersRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type FragmentGetCollectibleInfoRequestPredict struct {
	Collectible InputCollectible `tl:"collectible"`
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
	Source string `tl:"source"`
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
	PendingUpdatesCount int32  `tl:"pending_updates_count"`
	Message             string `tl:"message"`
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
	Referer string `tl:"referer"`
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
	ID DataJSON `tl:"id"`
}

func (*HelpAcceptTermsOfServiceRequestPredict) CRC() uint32 {
	return 0xee72f79a
}

func HelpAcceptTermsOfService(ctx context.Context, m Requester, i HelpAcceptTermsOfServiceRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetDeepLinkInfoRequestPredict struct {
	Path string `tl:"path"`
}

func (*HelpGetDeepLinkInfoRequestPredict) CRC() uint32 {
	return 0x3fedc75f
}

func HelpGetDeepLinkInfo(ctx context.Context, m Requester, i HelpGetDeepLinkInfoRequestPredict) (HelpDeepLinkInfo, error) {
	var res HelpDeepLinkInfo
	return res, request(ctx, m, &i, &res)
}

type HelpGetAppConfigRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*HelpGetAppConfigRequestPredict) CRC() uint32 {
	return 0x61e3f854
}

func HelpGetAppConfig(ctx context.Context, m Requester, i HelpGetAppConfigRequestPredict) (HelpAppConfig, error) {
	var res HelpAppConfig
	return res, request(ctx, m, &i, &res)
}

type HelpSaveAppLogRequestPredict struct {
	Events []InputAppEvent `tl:"events"`
}

func (*HelpSaveAppLogRequestPredict) CRC() uint32 {
	return 0x6f02f748
}

func HelpSaveAppLog(ctx context.Context, m Requester, i HelpSaveAppLogRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetPassportConfigRequestPredict struct {
	Hash int32 `tl:"hash"`
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
	UserID InputUser `tl:"user_id"`
}

func (*HelpGetUserInfoRequestPredict) CRC() uint32 {
	return 0x038a08d3
}

func HelpGetUserInfo(ctx context.Context, m Requester, i HelpGetUserInfoRequestPredict) (HelpUserInfo, error) {
	var res HelpUserInfo
	return res, request(ctx, m, &i, &res)
}

type HelpEditUserInfoRequestPredict struct {
	UserID   InputUser       `tl:"user_id"`
	Message  string          `tl:"message"`
	Entities []MessageEntity `tl:"entities"`
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
	Peer InputPeer `tl:"peer"`
}

func (*HelpHidePromoDataRequestPredict) CRC() uint32 {
	return 0x1e251c95
}

func HelpHidePromoData(ctx context.Context, m Requester, i HelpHidePromoDataRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpDismissSuggestionRequestPredict struct {
	Peer       InputPeer `tl:"peer"`
	Suggestion string    `tl:"suggestion"`
}

func (*HelpDismissSuggestionRequestPredict) CRC() uint32 {
	return 0xf50dbaa1
}

func HelpDismissSuggestion(ctx context.Context, m Requester, i HelpDismissSuggestionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type HelpGetCountriesListRequestPredict struct {
	LangCode string `tl:"lang_code"`
	Hash     int32  `tl:"hash"`
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
	Hash int32 `tl:"hash"`
}

func (*HelpGetPeerColorsRequestPredict) CRC() uint32 {
	return 0xda80f42f
}

func HelpGetPeerColors(ctx context.Context, m Requester, i HelpGetPeerColorsRequestPredict) (HelpPeerColors, error) {
	var res HelpPeerColors
	return res, request(ctx, m, &i, &res)
}

type HelpGetPeerProfileColorsRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*HelpGetPeerProfileColorsRequestPredict) CRC() uint32 {
	return 0xabcfa9fd
}

func HelpGetPeerProfileColors(ctx context.Context, m Requester, i HelpGetPeerProfileColorsRequestPredict) (HelpPeerColors, error) {
	var res HelpPeerColors
	return res, request(ctx, m, &i, &res)
}

type HelpGetTimezonesListRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*HelpGetTimezonesListRequestPredict) CRC() uint32 {
	return 0x49b30240
}

func HelpGetTimezonesList(ctx context.Context, m Requester, i HelpGetTimezonesListRequestPredict) (HelpTimezonesList, error) {
	var res HelpTimezonesList
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLangPackRequestPredict struct {
	LangPack string `tl:"lang_pack"`
	LangCode string `tl:"lang_code"`
}

func (*LangpackGetLangPackRequestPredict) CRC() uint32 {
	return 0xf2f2330a
}

func LangpackGetLangPack(ctx context.Context, m Requester, i LangpackGetLangPackRequestPredict) (LangPackDifference, error) {
	var res LangPackDifference
	return res, request(ctx, m, &i, &res)
}

type LangpackGetStringsRequestPredict struct {
	LangPack string   `tl:"lang_pack"`
	LangCode string   `tl:"lang_code"`
	Keys     []string `tl:"keys"`
}

func (*LangpackGetStringsRequestPredict) CRC() uint32 {
	return 0xefea3803
}

func LangpackGetStrings(ctx context.Context, m Requester, i LangpackGetStringsRequestPredict) ([]LangPackString, error) {
	var res []LangPackString
	return res, request(ctx, m, &i, &res)
}

type LangpackGetDifferenceRequestPredict struct {
	LangPack    string `tl:"lang_pack"`
	LangCode    string `tl:"lang_code"`
	FromVersion int32  `tl:"from_version"`
}

func (*LangpackGetDifferenceRequestPredict) CRC() uint32 {
	return 0xcd984aa5
}

func LangpackGetDifference(ctx context.Context, m Requester, i LangpackGetDifferenceRequestPredict) (LangPackDifference, error) {
	var res LangPackDifference
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLanguagesRequestPredict struct {
	LangPack string `tl:"lang_pack"`
}

func (*LangpackGetLanguagesRequestPredict) CRC() uint32 {
	return 0x42c6978f
}

func LangpackGetLanguages(ctx context.Context, m Requester, i LangpackGetLanguagesRequestPredict) ([]LangPackLanguage, error) {
	var res []LangPackLanguage
	return res, request(ctx, m, &i, &res)
}

type LangpackGetLanguageRequestPredict struct {
	LangPack string `tl:"lang_pack"`
	LangCode string `tl:"lang_code"`
}

func (*LangpackGetLanguageRequestPredict) CRC() uint32 {
	return 0x6a596502
}

func LangpackGetLanguage(ctx context.Context, m Requester, i LangpackGetLanguageRequestPredict) (LangPackLanguage, error) {
	var res LangPackLanguage
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesRequestPredict struct {
	ID []InputMessage `tl:"id"`
}

func (*MessagesGetMessagesRequestPredict) CRC() uint32 {
	return 0x63c66506
}

func MessagesGetMessages(ctx context.Context, m Requester, i MessagesGetMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDialogsRequestPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	ExcludePinned bool      `tl:"exclude_pinned,omitempty:flags:0,implicit"`
	FolderID      *int32    `tl:"folder_id,omitempty:flags:1"`
	OffsetDate    int32     `tl:"offset_date"`
	OffsetID      int32     `tl:"offset_id"`
	OffsetPeer    InputPeer `tl:"offset_peer"`
	Limit         int32     `tl:"limit"`
	Hash          int64     `tl:"hash"`
}

func (*MessagesGetDialogsRequestPredict) CRC() uint32 {
	return 0xa0f4cb4f
}

func MessagesGetDialogs(ctx context.Context, m Requester, i MessagesGetDialogsRequestPredict) (MessagesDialogs, error) {
	var res MessagesDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesGetHistoryRequestPredict struct {
	Peer       InputPeer `tl:"peer"`
	OffsetID   int32     `tl:"offset_id"`
	OffsetDate int32     `tl:"offset_date"`
	AddOffset  int32     `tl:"add_offset"`
	Limit      int32     `tl:"limit"`
	MaxID      int32     `tl:"max_id"`
	MinID      int32     `tl:"min_id"`
	Hash       int64     `tl:"hash"`
}

func (*MessagesGetHistoryRequestPredict) CRC() uint32 {
	return 0x4423e6c5
}

func MessagesGetHistory(ctx context.Context, m Requester, i MessagesGetHistoryRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchRequestPredict struct {
	_             struct{}       `tl:"flags,bitflag"`
	Peer          InputPeer      `tl:"peer"`
	Q             string         `tl:"q"`
	FromID        InputPeer      `tl:"from_id,omitempty:flags:0"`
	SavedPeerID   InputPeer      `tl:"saved_peer_id,omitempty:flags:2"`
	SavedReaction []Reaction     `tl:"saved_reaction,omitempty:flags:3"`
	TopMsgID      *int32         `tl:"top_msg_id,omitempty:flags:1"`
	Filter        MessagesFilter `tl:"filter"`
	MinDate       int32          `tl:"min_date"`
	MaxDate       int32          `tl:"max_date"`
	OffsetID      int32          `tl:"offset_id"`
	AddOffset     int32          `tl:"add_offset"`
	Limit         int32          `tl:"limit"`
	MaxID         int32          `tl:"max_id"`
	MinID         int32          `tl:"min_id"`
	Hash          int64          `tl:"hash"`
}

func (*MessagesSearchRequestPredict) CRC() uint32 {
	return 0x29ee847a
}

func MessagesSearch(ctx context.Context, m Requester, i MessagesSearchRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadHistoryRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MaxID int32     `tl:"max_id"`
}

func (*MessagesReadHistoryRequestPredict) CRC() uint32 {
	return 0x0e306d3a
}

func MessagesReadHistory(ctx context.Context, m Requester, i MessagesReadHistoryRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteHistoryRequestPredict struct {
	_         struct{}  `tl:"flags,bitflag"`
	JustClear bool      `tl:"just_clear,omitempty:flags:0,implicit"`
	Revoke    bool      `tl:"revoke,omitempty:flags:1,implicit"`
	Peer      InputPeer `tl:"peer"`
	MaxID     int32     `tl:"max_id"`
	MinDate   *int32    `tl:"min_date,omitempty:flags:2"`
	MaxDate   *int32    `tl:"max_date,omitempty:flags:3"`
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
	Revoke bool     `tl:"revoke,omitempty:flags:0,implicit"`
	ID     []int32  `tl:"id"`
}

func (*MessagesDeleteMessagesRequestPredict) CRC() uint32 {
	return 0xe58e95d2
}

func MessagesDeleteMessages(ctx context.Context, m Requester, i MessagesDeleteMessagesRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReceivedMessagesRequestPredict struct {
	MaxID int32 `tl:"max_id"`
}

func (*MessagesReceivedMessagesRequestPredict) CRC() uint32 {
	return 0x05a954c0
}

func MessagesReceivedMessages(ctx context.Context, m Requester, i MessagesReceivedMessagesRequestPredict) ([]ReceivedNotifyMessage, error) {
	var res []ReceivedNotifyMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSetTypingRequestPredict struct {
	_        struct{}          `tl:"flags,bitflag"`
	Peer     InputPeer         `tl:"peer"`
	TopMsgID *int32            `tl:"top_msg_id,omitempty:flags:0"`
	Action   SendMessageAction `tl:"action"`
}

func (*MessagesSetTypingRequestPredict) CRC() uint32 {
	return 0x58943ee2
}

func MessagesSetTyping(ctx context.Context, m Requester, i MessagesSetTypingRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMessageRequestPredict struct {
	_                      struct{}                `tl:"flags,bitflag"`
	NoWebpage              bool                    `tl:"no_webpage,omitempty:flags:1,implicit"`
	Silent                 bool                    `tl:"silent,omitempty:flags:5,implicit"`
	Background             bool                    `tl:"background,omitempty:flags:6,implicit"`
	ClearDraft             bool                    `tl:"clear_draft,omitempty:flags:7,implicit"`
	Noforwards             bool                    `tl:"noforwards,omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool                    `tl:"update_stickersets_order,omitempty:flags:15,implicit"`
	InvertMedia            bool                    `tl:"invert_media,omitempty:flags:16,implicit"`
	Peer                   InputPeer               `tl:"peer"`
	ReplyTo                InputReplyTo            `tl:"reply_to,omitempty:flags:0"`
	Message                string                  `tl:"message"`
	RandomID               int64                   `tl:"random_id"`
	ReplyMarkup            ReplyMarkup             `tl:"reply_markup,omitempty:flags:2"`
	Entities               []MessageEntity         `tl:"entities,omitempty:flags:3"`
	ScheduleDate           *int32                  `tl:"schedule_date,omitempty:flags:10"`
	SendAs                 InputPeer               `tl:"send_as,omitempty:flags:13"`
	QuickReplyShortcut     InputQuickReplyShortcut `tl:"quick_reply_shortcut,omitempty:flags:17"`
	Effect                 *int64                  `tl:"effect,omitempty:flags:18"`
}

func (*MessagesSendMessageRequestPredict) CRC() uint32 {
	return 0x983f9745
}

func MessagesSendMessage(ctx context.Context, m Requester, i MessagesSendMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMediaRequestPredict struct {
	_                      struct{}                `tl:"flags,bitflag"`
	Silent                 bool                    `tl:"silent,omitempty:flags:5,implicit"`
	Background             bool                    `tl:"background,omitempty:flags:6,implicit"`
	ClearDraft             bool                    `tl:"clear_draft,omitempty:flags:7,implicit"`
	Noforwards             bool                    `tl:"noforwards,omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool                    `tl:"update_stickersets_order,omitempty:flags:15,implicit"`
	InvertMedia            bool                    `tl:"invert_media,omitempty:flags:16,implicit"`
	Peer                   InputPeer               `tl:"peer"`
	ReplyTo                InputReplyTo            `tl:"reply_to,omitempty:flags:0"`
	Media                  InputMedia              `tl:"media"`
	Message                string                  `tl:"message"`
	RandomID               int64                   `tl:"random_id"`
	ReplyMarkup            ReplyMarkup             `tl:"reply_markup,omitempty:flags:2"`
	Entities               []MessageEntity         `tl:"entities,omitempty:flags:3"`
	ScheduleDate           *int32                  `tl:"schedule_date,omitempty:flags:10"`
	SendAs                 InputPeer               `tl:"send_as,omitempty:flags:13"`
	QuickReplyShortcut     InputQuickReplyShortcut `tl:"quick_reply_shortcut,omitempty:flags:17"`
	Effect                 *int64                  `tl:"effect,omitempty:flags:18"`
}

func (*MessagesSendMediaRequestPredict) CRC() uint32 {
	return 0x7852834e
}

func MessagesSendMedia(ctx context.Context, m Requester, i MessagesSendMediaRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesForwardMessagesRequestPredict struct {
	_                  struct{}                `tl:"flags,bitflag"`
	Silent             bool                    `tl:"silent,omitempty:flags:5,implicit"`
	Background         bool                    `tl:"background,omitempty:flags:6,implicit"`
	WithMyScore        bool                    `tl:"with_my_score,omitempty:flags:8,implicit"`
	DropAuthor         bool                    `tl:"drop_author,omitempty:flags:11,implicit"`
	DropMediaCaptions  bool                    `tl:"drop_media_captions,omitempty:flags:12,implicit"`
	Noforwards         bool                    `tl:"noforwards,omitempty:flags:14,implicit"`
	FromPeer           InputPeer               `tl:"from_peer"`
	ID                 []int32                 `tl:"id"`
	RandomID           []int64                 `tl:"random_id"`
	ToPeer             InputPeer               `tl:"to_peer"`
	TopMsgID           *int32                  `tl:"top_msg_id,omitempty:flags:9"`
	ScheduleDate       *int32                  `tl:"schedule_date,omitempty:flags:10"`
	SendAs             InputPeer               `tl:"send_as,omitempty:flags:13"`
	QuickReplyShortcut InputQuickReplyShortcut `tl:"quick_reply_shortcut,omitempty:flags:17"`
}

func (*MessagesForwardMessagesRequestPredict) CRC() uint32 {
	return 0xd5039208
}

func MessagesForwardMessages(ctx context.Context, m Requester, i MessagesForwardMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesReportSpamRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*MessagesReportSpamRequestPredict) CRC() uint32 {
	return 0xcf1592db
}

func MessagesReportSpam(ctx context.Context, m Requester, i MessagesReportSpamRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPeerSettingsRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*MessagesGetPeerSettingsRequestPredict) CRC() uint32 {
	return 0xefd9a6a2
}

func MessagesGetPeerSettings(ctx context.Context, m Requester, i MessagesGetPeerSettingsRequestPredict) (MessagesPeerSettings, error) {
	var res MessagesPeerSettings
	return res, request(ctx, m, &i, &res)
}

type MessagesReportRequestPredict struct {
	Peer    InputPeer    `tl:"peer"`
	ID      []int32      `tl:"id"`
	Reason  ReportReason `tl:"reason"`
	Message string       `tl:"message"`
}

func (*MessagesReportRequestPredict) CRC() uint32 {
	return 0x8953ab4e
}

func MessagesReport(ctx context.Context, m Requester, i MessagesReportRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetChatsRequestPredict struct {
	ID []int64 `tl:"id"`
}

func (*MessagesGetChatsRequestPredict) CRC() uint32 {
	return 0x49e9528f
}

func MessagesGetChats(ctx context.Context, m Requester, i MessagesGetChatsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFullChatRequestPredict struct {
	ChatID int64 `tl:"chat_id"`
}

func (*MessagesGetFullChatRequestPredict) CRC() uint32 {
	return 0xaeb00b34
}

func MessagesGetFullChat(ctx context.Context, m Requester, i MessagesGetFullChatRequestPredict) (MessagesChatFull, error) {
	var res MessagesChatFull
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatTitleRequestPredict struct {
	ChatID int64  `tl:"chat_id"`
	Title  string `tl:"title"`
}

func (*MessagesEditChatTitleRequestPredict) CRC() uint32 {
	return 0x73783ffd
}

func MessagesEditChatTitle(ctx context.Context, m Requester, i MessagesEditChatTitleRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatPhotoRequestPredict struct {
	ChatID int64          `tl:"chat_id"`
	Photo  InputChatPhoto `tl:"photo"`
}

func (*MessagesEditChatPhotoRequestPredict) CRC() uint32 {
	return 0x35ddd674
}

func MessagesEditChatPhoto(ctx context.Context, m Requester, i MessagesEditChatPhotoRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesAddChatUserRequestPredict struct {
	ChatID   int64     `tl:"chat_id"`
	UserID   InputUser `tl:"user_id"`
	FwdLimit int32     `tl:"fwd_limit"`
}

func (*MessagesAddChatUserRequestPredict) CRC() uint32 {
	return 0xcbc6d107
}

func MessagesAddChatUser(ctx context.Context, m Requester, i MessagesAddChatUserRequestPredict) (MessagesInvitedUsers, error) {
	var res MessagesInvitedUsers
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteChatUserRequestPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	RevokeHistory bool      `tl:"revoke_history,omitempty:flags:0,implicit"`
	ChatID        int64     `tl:"chat_id"`
	UserID        InputUser `tl:"user_id"`
}

func (*MessagesDeleteChatUserRequestPredict) CRC() uint32 {
	return 0xa2185cab
}

func MessagesDeleteChatUser(ctx context.Context, m Requester, i MessagesDeleteChatUserRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesCreateChatRequestPredict struct {
	_         struct{}    `tl:"flags,bitflag"`
	Users     []InputUser `tl:"users"`
	Title     string      `tl:"title"`
	TTLPeriod *int32      `tl:"ttl_period,omitempty:flags:0"`
}

func (*MessagesCreateChatRequestPredict) CRC() uint32 {
	return 0x92ceddd4
}

func MessagesCreateChat(ctx context.Context, m Requester, i MessagesCreateChatRequestPredict) (MessagesInvitedUsers, error) {
	var res MessagesInvitedUsers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDhConfigRequestPredict struct {
	Version      int32 `tl:"version"`
	RandomLength int32 `tl:"random_length"`
}

func (*MessagesGetDhConfigRequestPredict) CRC() uint32 {
	return 0x26cf8950
}

func MessagesGetDhConfig(ctx context.Context, m Requester, i MessagesGetDhConfigRequestPredict) (MessagesDhConfig, error) {
	var res MessagesDhConfig
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestEncryptionRequestPredict struct {
	UserID   InputUser `tl:"user_id"`
	RandomID int32     `tl:"random_id"`
	GA       []byte    `tl:"g_a"`
}

func (*MessagesRequestEncryptionRequestPredict) CRC() uint32 {
	return 0xf64daf43
}

func MessagesRequestEncryption(ctx context.Context, m Requester, i MessagesRequestEncryptionRequestPredict) (EncryptedChat, error) {
	var res EncryptedChat
	return res, request(ctx, m, &i, &res)
}

type MessagesAcceptEncryptionRequestPredict struct {
	Peer           InputEncryptedChat `tl:"peer"`
	GB             []byte             `tl:"g_b"`
	KeyFingerprint int64              `tl:"key_fingerprint"`
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
	DeleteHistory bool     `tl:"delete_history,omitempty:flags:0,implicit"`
	ChatID        int32    `tl:"chat_id"`
}

func (*MessagesDiscardEncryptionRequestPredict) CRC() uint32 {
	return 0xf393aea0
}

func MessagesDiscardEncryption(ctx context.Context, m Requester, i MessagesDiscardEncryptionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSetEncryptedTypingRequestPredict struct {
	Peer   InputEncryptedChat `tl:"peer"`
	Typing bool               `tl:"typing"`
}

func (*MessagesSetEncryptedTypingRequestPredict) CRC() uint32 {
	return 0x791451ed
}

func MessagesSetEncryptedTyping(ctx context.Context, m Requester, i MessagesSetEncryptedTypingRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReadEncryptedHistoryRequestPredict struct {
	Peer    InputEncryptedChat `tl:"peer"`
	MaxDate int32              `tl:"max_date"`
}

func (*MessagesReadEncryptedHistoryRequestPredict) CRC() uint32 {
	return 0x7f4b690a
}

func MessagesReadEncryptedHistory(ctx context.Context, m Requester, i MessagesReadEncryptedHistoryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedRequestPredict struct {
	_        struct{}           `tl:"flags,bitflag"`
	Silent   bool               `tl:"silent,omitempty:flags:0,implicit"`
	Peer     InputEncryptedChat `tl:"peer"`
	RandomID int64              `tl:"random_id"`
	Data     []byte             `tl:"data"`
}

func (*MessagesSendEncryptedRequestPredict) CRC() uint32 {
	return 0x44fa7a15
}

func MessagesSendEncrypted(ctx context.Context, m Requester, i MessagesSendEncryptedRequestPredict) (MessagesSentEncryptedMessage, error) {
	var res MessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedFileRequestPredict struct {
	_        struct{}           `tl:"flags,bitflag"`
	Silent   bool               `tl:"silent,omitempty:flags:0,implicit"`
	Peer     InputEncryptedChat `tl:"peer"`
	RandomID int64              `tl:"random_id"`
	Data     []byte             `tl:"data"`
	File     InputEncryptedFile `tl:"file"`
}

func (*MessagesSendEncryptedFileRequestPredict) CRC() uint32 {
	return 0x5559481d
}

func MessagesSendEncryptedFile(ctx context.Context, m Requester, i MessagesSendEncryptedFileRequestPredict) (MessagesSentEncryptedMessage, error) {
	var res MessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesSendEncryptedServiceRequestPredict struct {
	Peer     InputEncryptedChat `tl:"peer"`
	RandomID int64              `tl:"random_id"`
	Data     []byte             `tl:"data"`
}

func (*MessagesSendEncryptedServiceRequestPredict) CRC() uint32 {
	return 0x32d439a4
}

func MessagesSendEncryptedService(ctx context.Context, m Requester, i MessagesSendEncryptedServiceRequestPredict) (MessagesSentEncryptedMessage, error) {
	var res MessagesSentEncryptedMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesReceivedQueueRequestPredict struct {
	MaxQts int32 `tl:"max_qts"`
}

func (*MessagesReceivedQueueRequestPredict) CRC() uint32 {
	return 0x55a5bb66
}

func MessagesReceivedQueue(ctx context.Context, m Requester, i MessagesReceivedQueueRequestPredict) ([]int64, error) {
	var res []int64
	return res, request(ctx, m, &i, &res)
}

type MessagesReportEncryptedSpamRequestPredict struct {
	Peer InputEncryptedChat `tl:"peer"`
}

func (*MessagesReportEncryptedSpamRequestPredict) CRC() uint32 {
	return 0x4b0c8c0f
}

func MessagesReportEncryptedSpam(ctx context.Context, m Requester, i MessagesReportEncryptedSpamRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReadMessageContentsRequestPredict struct {
	ID []int32 `tl:"id"`
}

func (*MessagesReadMessageContentsRequestPredict) CRC() uint32 {
	return 0x36a73f77
}

func MessagesReadMessageContents(ctx context.Context, m Requester, i MessagesReadMessageContentsRequestPredict) (MessagesAffectedMessages, error) {
	var res MessagesAffectedMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetStickersRequestPredict struct {
	Emoticon string `tl:"emoticon"`
	Hash     int64  `tl:"hash"`
}

func (*MessagesGetStickersRequestPredict) CRC() uint32 {
	return 0xd5a5d3a1
}

func MessagesGetStickers(ctx context.Context, m Requester, i MessagesGetStickersRequestPredict) (MessagesStickers, error) {
	var res MessagesStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAllStickersRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetAllStickersRequestPredict) CRC() uint32 {
	return 0xb8a0a1a8
}

func MessagesGetAllStickers(ctx context.Context, m Requester, i MessagesGetAllStickersRequestPredict) (MessagesAllStickers, error) {
	var res MessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetWebPagePreviewRequestPredict struct {
	_        struct{}        `tl:"flags,bitflag"`
	Message  string          `tl:"message"`
	Entities []MessageEntity `tl:"entities,omitempty:flags:3"`
}

func (*MessagesGetWebPagePreviewRequestPredict) CRC() uint32 {
	return 0x8b68b0cc
}

func MessagesGetWebPagePreview(ctx context.Context, m Requester, i MessagesGetWebPagePreviewRequestPredict) (MessageMedia, error) {
	var res MessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesExportChatInviteRequestPredict struct {
	_                     struct{}  `tl:"flags,bitflag"`
	LegacyRevokePermanent bool      `tl:"legacy_revoke_permanent,omitempty:flags:2,implicit"`
	RequestNeeded         bool      `tl:"request_needed,omitempty:flags:3,implicit"`
	Peer                  InputPeer `tl:"peer"`
	ExpireDate            *int32    `tl:"expire_date,omitempty:flags:0"`
	UsageLimit            *int32    `tl:"usage_limit,omitempty:flags:1"`
	Title                 *string   `tl:"title,omitempty:flags:4"`
}

func (*MessagesExportChatInviteRequestPredict) CRC() uint32 {
	return 0xa02ce5d5
}

func MessagesExportChatInvite(ctx context.Context, m Requester, i MessagesExportChatInviteRequestPredict) (ExportedChatInvite, error) {
	var res ExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckChatInviteRequestPredict struct {
	Hash string `tl:"hash"`
}

func (*MessagesCheckChatInviteRequestPredict) CRC() uint32 {
	return 0x3eadb1bb
}

func MessagesCheckChatInvite(ctx context.Context, m Requester, i MessagesCheckChatInviteRequestPredict) (ChatInvite, error) {
	var res ChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesImportChatInviteRequestPredict struct {
	Hash string `tl:"hash"`
}

func (*MessagesImportChatInviteRequestPredict) CRC() uint32 {
	return 0x6c50051c
}

func MessagesImportChatInvite(ctx context.Context, m Requester, i MessagesImportChatInviteRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetStickerSetRequestPredict struct {
	Stickerset InputStickerSet `tl:"stickerset"`
	Hash       int32           `tl:"hash"`
}

func (*MessagesGetStickerSetRequestPredict) CRC() uint32 {
	return 0xc8a0ec74
}

func MessagesGetStickerSet(ctx context.Context, m Requester, i MessagesGetStickerSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type MessagesInstallStickerSetRequestPredict struct {
	Stickerset InputStickerSet `tl:"stickerset"`
	Archived   bool            `tl:"archived"`
}

func (*MessagesInstallStickerSetRequestPredict) CRC() uint32 {
	return 0xc78fe460
}

func MessagesInstallStickerSet(ctx context.Context, m Requester, i MessagesInstallStickerSetRequestPredict) (MessagesStickerSetInstallResult, error) {
	var res MessagesStickerSetInstallResult
	return res, request(ctx, m, &i, &res)
}

type MessagesUninstallStickerSetRequestPredict struct {
	Stickerset InputStickerSet `tl:"stickerset"`
}

func (*MessagesUninstallStickerSetRequestPredict) CRC() uint32 {
	return 0xf96e55de
}

func MessagesUninstallStickerSet(ctx context.Context, m Requester, i MessagesUninstallStickerSetRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesStartBotRequestPredict struct {
	Bot        InputUser `tl:"bot"`
	Peer       InputPeer `tl:"peer"`
	RandomID   int64     `tl:"random_id"`
	StartParam string    `tl:"start_param"`
}

func (*MessagesStartBotRequestPredict) CRC() uint32 {
	return 0xe6df7378
}

func MessagesStartBot(ctx context.Context, m Requester, i MessagesStartBotRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesViewsRequestPredict struct {
	Peer      InputPeer `tl:"peer"`
	ID        []int32   `tl:"id"`
	Increment bool      `tl:"increment"`
}

func (*MessagesGetMessagesViewsRequestPredict) CRC() uint32 {
	return 0x5784d3e1
}

func MessagesGetMessagesViews(ctx context.Context, m Requester, i MessagesGetMessagesViewsRequestPredict) (MessagesMessageViews, error) {
	var res MessagesMessageViews
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatAdminRequestPredict struct {
	ChatID  int64     `tl:"chat_id"`
	UserID  InputUser `tl:"user_id"`
	IsAdmin bool      `tl:"is_admin"`
}

func (*MessagesEditChatAdminRequestPredict) CRC() uint32 {
	return 0xa85bd1c2
}

func MessagesEditChatAdmin(ctx context.Context, m Requester, i MessagesEditChatAdminRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesMigrateChatRequestPredict struct {
	ChatID int64 `tl:"chat_id"`
}

func (*MessagesMigrateChatRequestPredict) CRC() uint32 {
	return 0xa2875319
}

func MessagesMigrateChat(ctx context.Context, m Requester, i MessagesMigrateChatRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchGlobalRequestPredict struct {
	_              struct{}       `tl:"flags,bitflag"`
	BroadcastsOnly bool           `tl:"broadcasts_only,omitempty:flags:1,implicit"`
	FolderID       *int32         `tl:"folder_id,omitempty:flags:0"`
	Q              string         `tl:"q"`
	Filter         MessagesFilter `tl:"filter"`
	MinDate        int32          `tl:"min_date"`
	MaxDate        int32          `tl:"max_date"`
	OffsetRate     int32          `tl:"offset_rate"`
	OffsetPeer     InputPeer      `tl:"offset_peer"`
	OffsetID       int32          `tl:"offset_id"`
	Limit          int32          `tl:"limit"`
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
	Masks  bool     `tl:"masks,omitempty:flags:0,implicit"`
	Emojis bool     `tl:"emojis,omitempty:flags:1,implicit"`
	Order  []int64  `tl:"order"`
}

func (*MessagesReorderStickerSetsRequestPredict) CRC() uint32 {
	return 0x78337739
}

func MessagesReorderStickerSets(ctx context.Context, m Requester, i MessagesReorderStickerSetsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDocumentByHashRequestPredict struct {
	SHA256   []byte `tl:"sha256"`
	Size     int64  `tl:"size"`
	MimeType string `tl:"mime_type"`
}

func (*MessagesGetDocumentByHashRequestPredict) CRC() uint32 {
	return 0xb1f2061f
}

func MessagesGetDocumentByHash(ctx context.Context, m Requester, i MessagesGetDocumentByHashRequestPredict) (Document, error) {
	var res Document
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedGifsRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetSavedGifsRequestPredict) CRC() uint32 {
	return 0x5cf09635
}

func MessagesGetSavedGifs(ctx context.Context, m Requester, i MessagesGetSavedGifsRequestPredict) (MessagesSavedGifs, error) {
	var res MessagesSavedGifs
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveGifRequestPredict struct {
	ID     InputDocument `tl:"id"`
	Unsave bool          `tl:"unsave"`
}

func (*MessagesSaveGifRequestPredict) CRC() uint32 {
	return 0x327a30cb
}

func MessagesSaveGif(ctx context.Context, m Requester, i MessagesSaveGifRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetInlineBotResultsRequestPredict struct {
	_        struct{}      `tl:"flags,bitflag"`
	Bot      InputUser     `tl:"bot"`
	Peer     InputPeer     `tl:"peer"`
	GeoPoint InputGeoPoint `tl:"geo_point,omitempty:flags:0"`
	Query    string        `tl:"query"`
	Offset   string        `tl:"offset"`
}

func (*MessagesGetInlineBotResultsRequestPredict) CRC() uint32 {
	return 0x514e999d
}

func MessagesGetInlineBotResults(ctx context.Context, m Requester, i MessagesGetInlineBotResultsRequestPredict) (MessagesBotResults, error) {
	var res MessagesBotResults
	return res, request(ctx, m, &i, &res)
}

type MessagesSetInlineBotResultsRequestPredict struct {
	_             struct{}               `tl:"flags,bitflag"`
	Gallery       bool                   `tl:"gallery,omitempty:flags:0,implicit"`
	Private       bool                   `tl:"private,omitempty:flags:1,implicit"`
	QueryID       int64                  `tl:"query_id"`
	Results       []InputBotInlineResult `tl:"results"`
	CacheTime     int32                  `tl:"cache_time"`
	NextOffset    *string                `tl:"next_offset,omitempty:flags:2"`
	SwitchPm      InlineBotSwitchPm      `tl:"switch_pm,omitempty:flags:3"`
	SwitchWebview InlineBotWebView       `tl:"switch_webview,omitempty:flags:4"`
}

func (*MessagesSetInlineBotResultsRequestPredict) CRC() uint32 {
	return 0xbb12a419
}

func MessagesSetInlineBotResults(ctx context.Context, m Requester, i MessagesSetInlineBotResultsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendInlineBotResultRequestPredict struct {
	_                  struct{}                `tl:"flags,bitflag"`
	Silent             bool                    `tl:"silent,omitempty:flags:5,implicit"`
	Background         bool                    `tl:"background,omitempty:flags:6,implicit"`
	ClearDraft         bool                    `tl:"clear_draft,omitempty:flags:7,implicit"`
	HideVia            bool                    `tl:"hide_via,omitempty:flags:11,implicit"`
	Peer               InputPeer               `tl:"peer"`
	ReplyTo            InputReplyTo            `tl:"reply_to,omitempty:flags:0"`
	RandomID           int64                   `tl:"random_id"`
	QueryID            int64                   `tl:"query_id"`
	ID                 string                  `tl:"id"`
	ScheduleDate       *int32                  `tl:"schedule_date,omitempty:flags:10"`
	SendAs             InputPeer               `tl:"send_as,omitempty:flags:13"`
	QuickReplyShortcut InputQuickReplyShortcut `tl:"quick_reply_shortcut,omitempty:flags:17"`
}

func (*MessagesSendInlineBotResultRequestPredict) CRC() uint32 {
	return 0x3ebee86a
}

func MessagesSendInlineBotResult(ctx context.Context, m Requester, i MessagesSendInlineBotResultRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageEditDataRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   int32     `tl:"id"`
}

func (*MessagesGetMessageEditDataRequestPredict) CRC() uint32 {
	return 0xfda68d36
}

func MessagesGetMessageEditData(ctx context.Context, m Requester, i MessagesGetMessageEditDataRequestPredict) (MessagesMessageEditData, error) {
	var res MessagesMessageEditData
	return res, request(ctx, m, &i, &res)
}

type MessagesEditMessageRequestPredict struct {
	_                    struct{}        `tl:"flags,bitflag"`
	NoWebpage            bool            `tl:"no_webpage,omitempty:flags:1,implicit"`
	InvertMedia          bool            `tl:"invert_media,omitempty:flags:16,implicit"`
	Peer                 InputPeer       `tl:"peer"`
	ID                   int32           `tl:"id"`
	Message              *string         `tl:"message,omitempty:flags:11"`
	Media                InputMedia      `tl:"media,omitempty:flags:14"`
	ReplyMarkup          ReplyMarkup     `tl:"reply_markup,omitempty:flags:2"`
	Entities             []MessageEntity `tl:"entities,omitempty:flags:3"`
	ScheduleDate         *int32          `tl:"schedule_date,omitempty:flags:15"`
	QuickReplyShortcutID *int32          `tl:"quick_reply_shortcut_id,omitempty:flags:17"`
}

func (*MessagesEditMessageRequestPredict) CRC() uint32 {
	return 0xdfd14005
}

func MessagesEditMessage(ctx context.Context, m Requester, i MessagesEditMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesEditInlineBotMessageRequestPredict struct {
	_           struct{}                `tl:"flags,bitflag"`
	NoWebpage   bool                    `tl:"no_webpage,omitempty:flags:1,implicit"`
	InvertMedia bool                    `tl:"invert_media,omitempty:flags:16,implicit"`
	ID          InputBotInlineMessageID `tl:"id"`
	Message     *string                 `tl:"message,omitempty:flags:11"`
	Media       InputMedia              `tl:"media,omitempty:flags:14"`
	ReplyMarkup ReplyMarkup             `tl:"reply_markup,omitempty:flags:2"`
	Entities    []MessageEntity         `tl:"entities,omitempty:flags:3"`
}

func (*MessagesEditInlineBotMessageRequestPredict) CRC() uint32 {
	return 0x83557dba
}

func MessagesEditInlineBotMessage(ctx context.Context, m Requester, i MessagesEditInlineBotMessageRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetBotCallbackAnswerRequestPredict struct {
	_        struct{}              `tl:"flags,bitflag"`
	Game     bool                  `tl:"game,omitempty:flags:1,implicit"`
	Peer     InputPeer             `tl:"peer"`
	MsgID    int32                 `tl:"msg_id"`
	Data     *[]byte               `tl:"data,omitempty:flags:0"`
	Password InputCheckPasswordSRP `tl:"password,omitempty:flags:2"`
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
	Alert     bool     `tl:"alert,omitempty:flags:1,implicit"`
	QueryID   int64    `tl:"query_id"`
	Message   *string  `tl:"message,omitempty:flags:0"`
	URL       *string  `tl:"url,omitempty:flags:2"`
	CacheTime int32    `tl:"cache_time"`
}

func (*MessagesSetBotCallbackAnswerRequestPredict) CRC() uint32 {
	return 0xd58f130a
}

func MessagesSetBotCallbackAnswer(ctx context.Context, m Requester, i MessagesSetBotCallbackAnswerRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPeerDialogsRequestPredict struct {
	Peers []InputDialogPeer `tl:"peers"`
}

func (*MessagesGetPeerDialogsRequestPredict) CRC() uint32 {
	return 0xe470bcfd
}

func MessagesGetPeerDialogs(ctx context.Context, m Requester, i MessagesGetPeerDialogsRequestPredict) (MessagesPeerDialogs, error) {
	var res MessagesPeerDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveDraftRequestPredict struct {
	_           struct{}        `tl:"flags,bitflag"`
	NoWebpage   bool            `tl:"no_webpage,omitempty:flags:1,implicit"`
	InvertMedia bool            `tl:"invert_media,omitempty:flags:6,implicit"`
	ReplyTo     InputReplyTo    `tl:"reply_to,omitempty:flags:4"`
	Peer        InputPeer       `tl:"peer"`
	Message     string          `tl:"message"`
	Entities    []MessageEntity `tl:"entities,omitempty:flags:3"`
	Media       InputMedia      `tl:"media,omitempty:flags:5"`
	Effect      *int64          `tl:"effect,omitempty:flags:7"`
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
	Hash int64 `tl:"hash"`
}

func (*MessagesGetFeaturedStickersRequestPredict) CRC() uint32 {
	return 0x64780b14
}

func MessagesGetFeaturedStickers(ctx context.Context, m Requester, i MessagesGetFeaturedStickersRequestPredict) (MessagesFeaturedStickers, error) {
	var res MessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesReadFeaturedStickersRequestPredict struct {
	ID []int64 `tl:"id"`
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
	Attached bool     `tl:"attached,omitempty:flags:0,implicit"`
	Hash     int64    `tl:"hash"`
}

func (*MessagesGetRecentStickersRequestPredict) CRC() uint32 {
	return 0x9da9403b
}

func MessagesGetRecentStickers(ctx context.Context, m Requester, i MessagesGetRecentStickersRequestPredict) (MessagesRecentStickers, error) {
	var res MessagesRecentStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveRecentStickerRequestPredict struct {
	_        struct{}      `tl:"flags,bitflag"`
	Attached bool          `tl:"attached,omitempty:flags:0,implicit"`
	ID       InputDocument `tl:"id"`
	Unsave   bool          `tl:"unsave"`
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
	Attached bool     `tl:"attached,omitempty:flags:0,implicit"`
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
	Masks    bool     `tl:"masks,omitempty:flags:0,implicit"`
	Emojis   bool     `tl:"emojis,omitempty:flags:1,implicit"`
	OffsetID int64    `tl:"offset_id"`
	Limit    int32    `tl:"limit"`
}

func (*MessagesGetArchivedStickersRequestPredict) CRC() uint32 {
	return 0x57f17692
}

func MessagesGetArchivedStickers(ctx context.Context, m Requester, i MessagesGetArchivedStickersRequestPredict) (MessagesArchivedStickers, error) {
	var res MessagesArchivedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMaskStickersRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetMaskStickersRequestPredict) CRC() uint32 {
	return 0x640f82b8
}

func MessagesGetMaskStickers(ctx context.Context, m Requester, i MessagesGetMaskStickersRequestPredict) (MessagesAllStickers, error) {
	var res MessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachedStickersRequestPredict struct {
	Media InputStickeredMedia `tl:"media"`
}

func (*MessagesGetAttachedStickersRequestPredict) CRC() uint32 {
	return 0xcc5b67cc
}

func MessagesGetAttachedStickers(ctx context.Context, m Requester, i MessagesGetAttachedStickersRequestPredict) ([]StickerSetCovered, error) {
	var res []StickerSetCovered
	return res, request(ctx, m, &i, &res)
}

type MessagesSetGameScoreRequestPredict struct {
	_           struct{}  `tl:"flags,bitflag"`
	EditMessage bool      `tl:"edit_message,omitempty:flags:0,implicit"`
	Force       bool      `tl:"force,omitempty:flags:1,implicit"`
	Peer        InputPeer `tl:"peer"`
	ID          int32     `tl:"id"`
	UserID      InputUser `tl:"user_id"`
	Score       int32     `tl:"score"`
}

func (*MessagesSetGameScoreRequestPredict) CRC() uint32 {
	return 0x8ef8ecc0
}

func MessagesSetGameScore(ctx context.Context, m Requester, i MessagesSetGameScoreRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSetInlineGameScoreRequestPredict struct {
	_           struct{}                `tl:"flags,bitflag"`
	EditMessage bool                    `tl:"edit_message,omitempty:flags:0,implicit"`
	Force       bool                    `tl:"force,omitempty:flags:1,implicit"`
	ID          InputBotInlineMessageID `tl:"id"`
	UserID      InputUser               `tl:"user_id"`
	Score       int32                   `tl:"score"`
}

func (*MessagesSetInlineGameScoreRequestPredict) CRC() uint32 {
	return 0x15ad9f64
}

func MessagesSetInlineGameScore(ctx context.Context, m Requester, i MessagesSetInlineGameScoreRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetGameHighScoresRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	ID     int32     `tl:"id"`
	UserID InputUser `tl:"user_id"`
}

func (*MessagesGetGameHighScoresRequestPredict) CRC() uint32 {
	return 0xe822649d
}

func MessagesGetGameHighScores(ctx context.Context, m Requester, i MessagesGetGameHighScoresRequestPredict) (MessagesHighScores, error) {
	var res MessagesHighScores
	return res, request(ctx, m, &i, &res)
}

type MessagesGetInlineGameHighScoresRequestPredict struct {
	ID     InputBotInlineMessageID `tl:"id"`
	UserID InputUser               `tl:"user_id"`
}

func (*MessagesGetInlineGameHighScoresRequestPredict) CRC() uint32 {
	return 0x0f635e1b
}

func MessagesGetInlineGameHighScores(ctx context.Context, m Requester, i MessagesGetInlineGameHighScoresRequestPredict) (MessagesHighScores, error) {
	var res MessagesHighScores
	return res, request(ctx, m, &i, &res)
}

type MessagesGetCommonChatsRequestPredict struct {
	UserID InputUser `tl:"user_id"`
	MaxID  int64     `tl:"max_id"`
	Limit  int32     `tl:"limit"`
}

func (*MessagesGetCommonChatsRequestPredict) CRC() uint32 {
	return 0xe40ca104
}

func MessagesGetCommonChats(ctx context.Context, m Requester, i MessagesGetCommonChatsRequestPredict) (MessagesChats, error) {
	var res MessagesChats
	return res, request(ctx, m, &i, &res)
}

type MessagesGetWebPageRequestPredict struct {
	URL  string `tl:"url"`
	Hash int32  `tl:"hash"`
}

func (*MessagesGetWebPageRequestPredict) CRC() uint32 {
	return 0x8d9692a3
}

func MessagesGetWebPage(ctx context.Context, m Requester, i MessagesGetWebPageRequestPredict) (MessagesWebPage, error) {
	var res MessagesWebPage
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleDialogPinRequestPredict struct {
	_      struct{}        `tl:"flags,bitflag"`
	Pinned bool            `tl:"pinned,omitempty:flags:0,implicit"`
	Peer   InputDialogPeer `tl:"peer"`
}

func (*MessagesToggleDialogPinRequestPredict) CRC() uint32 {
	return 0xa731e257
}

func MessagesToggleDialogPin(ctx context.Context, m Requester, i MessagesToggleDialogPinRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderPinnedDialogsRequestPredict struct {
	_        struct{}          `tl:"flags,bitflag"`
	Force    bool              `tl:"force,omitempty:flags:0,implicit"`
	FolderID int32             `tl:"folder_id"`
	Order    []InputDialogPeer `tl:"order"`
}

func (*MessagesReorderPinnedDialogsRequestPredict) CRC() uint32 {
	return 0x3b1adf37
}

func MessagesReorderPinnedDialogs(ctx context.Context, m Requester, i MessagesReorderPinnedDialogsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPinnedDialogsRequestPredict struct {
	FolderID int32 `tl:"folder_id"`
}

func (*MessagesGetPinnedDialogsRequestPredict) CRC() uint32 {
	return 0xd6b94df2
}

func MessagesGetPinnedDialogs(ctx context.Context, m Requester, i MessagesGetPinnedDialogsRequestPredict) (MessagesPeerDialogs, error) {
	var res MessagesPeerDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesSetBotShippingResultsRequestPredict struct {
	_               struct{}         `tl:"flags,bitflag"`
	QueryID         int64            `tl:"query_id"`
	Error           *string          `tl:"error,omitempty:flags:0"`
	ShippingOptions []ShippingOption `tl:"shipping_options,omitempty:flags:1"`
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
	Success bool     `tl:"success,omitempty:flags:1,implicit"`
	QueryID int64    `tl:"query_id"`
	Error   *string  `tl:"error,omitempty:flags:0"`
}

func (*MessagesSetBotPrecheckoutResultsRequestPredict) CRC() uint32 {
	return 0x09c2dd95
}

func MessagesSetBotPrecheckoutResults(ctx context.Context, m Requester, i MessagesSetBotPrecheckoutResultsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadMediaRequestPredict struct {
	_                    struct{}   `tl:"flags,bitflag"`
	BusinessConnectionID *string    `tl:"business_connection_id,omitempty:flags:0"`
	Peer                 InputPeer  `tl:"peer"`
	Media                InputMedia `tl:"media"`
}

func (*MessagesUploadMediaRequestPredict) CRC() uint32 {
	return 0x14967978
}

func MessagesUploadMedia(ctx context.Context, m Requester, i MessagesUploadMediaRequestPredict) (MessageMedia, error) {
	var res MessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesSendScreenshotNotificationRequestPredict struct {
	Peer     InputPeer    `tl:"peer"`
	ReplyTo  InputReplyTo `tl:"reply_to"`
	RandomID int64        `tl:"random_id"`
}

func (*MessagesSendScreenshotNotificationRequestPredict) CRC() uint32 {
	return 0xa1405817
}

func MessagesSendScreenshotNotification(ctx context.Context, m Requester, i MessagesSendScreenshotNotificationRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFavedStickersRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetFavedStickersRequestPredict) CRC() uint32 {
	return 0x04f1aaa9
}

func MessagesGetFavedStickers(ctx context.Context, m Requester, i MessagesGetFavedStickersRequestPredict) (MessagesFavedStickers, error) {
	var res MessagesFavedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesFaveStickerRequestPredict struct {
	ID     InputDocument `tl:"id"`
	Unfave bool          `tl:"unfave"`
}

func (*MessagesFaveStickerRequestPredict) CRC() uint32 {
	return 0xb9ffc55b
}

func MessagesFaveSticker(ctx context.Context, m Requester, i MessagesFaveStickerRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetUnreadMentionsRequestPredict struct {
	_         struct{}  `tl:"flags,bitflag"`
	Peer      InputPeer `tl:"peer"`
	TopMsgID  *int32    `tl:"top_msg_id,omitempty:flags:0"`
	OffsetID  int32     `tl:"offset_id"`
	AddOffset int32     `tl:"add_offset"`
	Limit     int32     `tl:"limit"`
	MaxID     int32     `tl:"max_id"`
	MinID     int32     `tl:"min_id"`
}

func (*MessagesGetUnreadMentionsRequestPredict) CRC() uint32 {
	return 0xf107e790
}

func MessagesGetUnreadMentions(ctx context.Context, m Requester, i MessagesGetUnreadMentionsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadMentionsRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Peer     InputPeer `tl:"peer"`
	TopMsgID *int32    `tl:"top_msg_id,omitempty:flags:0"`
}

func (*MessagesReadMentionsRequestPredict) CRC() uint32 {
	return 0x36e5bf4d
}

func MessagesReadMentions(ctx context.Context, m Requester, i MessagesReadMentionsRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentLocationsRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	Limit int32     `tl:"limit"`
	Hash  int64     `tl:"hash"`
}

func (*MessagesGetRecentLocationsRequestPredict) CRC() uint32 {
	return 0x702a40e0
}

func MessagesGetRecentLocations(ctx context.Context, m Requester, i MessagesGetRecentLocationsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendMultiMediaRequestPredict struct {
	_                      struct{}                `tl:"flags,bitflag"`
	Silent                 bool                    `tl:"silent,omitempty:flags:5,implicit"`
	Background             bool                    `tl:"background,omitempty:flags:6,implicit"`
	ClearDraft             bool                    `tl:"clear_draft,omitempty:flags:7,implicit"`
	Noforwards             bool                    `tl:"noforwards,omitempty:flags:14,implicit"`
	UpdateStickersetsOrder bool                    `tl:"update_stickersets_order,omitempty:flags:15,implicit"`
	InvertMedia            bool                    `tl:"invert_media,omitempty:flags:16,implicit"`
	Peer                   InputPeer               `tl:"peer"`
	ReplyTo                InputReplyTo            `tl:"reply_to,omitempty:flags:0"`
	MultiMedia             []InputSingleMedia      `tl:"multi_media"`
	ScheduleDate           *int32                  `tl:"schedule_date,omitempty:flags:10"`
	SendAs                 InputPeer               `tl:"send_as,omitempty:flags:13"`
	QuickReplyShortcut     InputQuickReplyShortcut `tl:"quick_reply_shortcut,omitempty:flags:17"`
	Effect                 *int64                  `tl:"effect,omitempty:flags:18"`
}

func (*MessagesSendMultiMediaRequestPredict) CRC() uint32 {
	return 0x37b74355
}

func MessagesSendMultiMedia(ctx context.Context, m Requester, i MessagesSendMultiMediaRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadEncryptedFileRequestPredict struct {
	Peer InputEncryptedChat `tl:"peer"`
	File InputEncryptedFile `tl:"file"`
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
	ExcludeFeatured bool     `tl:"exclude_featured,omitempty:flags:0,implicit"`
	Q               string   `tl:"q"`
	Hash            int64    `tl:"hash"`
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
	_      struct{}        `tl:"flags,bitflag"`
	Unread bool            `tl:"unread,omitempty:flags:0,implicit"`
	Peer   InputDialogPeer `tl:"peer"`
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
	_         struct{}  `tl:"flags,bitflag"`
	Silent    bool      `tl:"silent,omitempty:flags:0,implicit"`
	Unpin     bool      `tl:"unpin,omitempty:flags:1,implicit"`
	PmOneside bool      `tl:"pm_oneside,omitempty:flags:2,implicit"`
	Peer      InputPeer `tl:"peer"`
	ID        int32     `tl:"id"`
}

func (*MessagesUpdatePinnedMessageRequestPredict) CRC() uint32 {
	return 0xd2aaf7ec
}

func MessagesUpdatePinnedMessage(ctx context.Context, m Requester, i MessagesUpdatePinnedMessageRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSendVoteRequestPredict struct {
	Peer    InputPeer `tl:"peer"`
	MsgID   int32     `tl:"msg_id"`
	Options [][]byte  `tl:"options"`
}

func (*MessagesSendVoteRequestPredict) CRC() uint32 {
	return 0x10ea6184
}

func MessagesSendVote(ctx context.Context, m Requester, i MessagesSendVoteRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPollResultsRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*MessagesGetPollResultsRequestPredict) CRC() uint32 {
	return 0x73bb643b
}

func MessagesGetPollResults(ctx context.Context, m Requester, i MessagesGetPollResultsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOnlinesRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*MessagesGetOnlinesRequestPredict) CRC() uint32 {
	return 0x6e2be050
}

func MessagesGetOnlines(ctx context.Context, m Requester, i MessagesGetOnlinesRequestPredict) (ChatOnlines, error) {
	var res ChatOnlines
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatAboutRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	About string    `tl:"about"`
}

func (*MessagesEditChatAboutRequestPredict) CRC() uint32 {
	return 0xdef60797
}

func MessagesEditChatAbout(ctx context.Context, m Requester, i MessagesEditChatAboutRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesEditChatDefaultBannedRightsRequestPredict struct {
	Peer         InputPeer        `tl:"peer"`
	BannedRights ChatBannedRights `tl:"banned_rights"`
}

func (*MessagesEditChatDefaultBannedRightsRequestPredict) CRC() uint32 {
	return 0xa5866b41
}

func MessagesEditChatDefaultBannedRights(ctx context.Context, m Requester, i MessagesEditChatDefaultBannedRightsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsRequestPredict struct {
	LangCode string `tl:"lang_code"`
}

func (*MessagesGetEmojiKeywordsRequestPredict) CRC() uint32 {
	return 0x35a0e062
}

func MessagesGetEmojiKeywords(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsRequestPredict) (EmojiKeywordsDifference, error) {
	var res EmojiKeywordsDifference
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsDifferenceRequestPredict struct {
	LangCode    string `tl:"lang_code"`
	FromVersion int32  `tl:"from_version"`
}

func (*MessagesGetEmojiKeywordsDifferenceRequestPredict) CRC() uint32 {
	return 0x1508b6af
}

func MessagesGetEmojiKeywordsDifference(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsDifferenceRequestPredict) (EmojiKeywordsDifference, error) {
	var res EmojiKeywordsDifference
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiKeywordsLanguagesRequestPredict struct {
	LangCodes []string `tl:"lang_codes"`
}

func (*MessagesGetEmojiKeywordsLanguagesRequestPredict) CRC() uint32 {
	return 0x4e9963b2
}

func MessagesGetEmojiKeywordsLanguages(ctx context.Context, m Requester, i MessagesGetEmojiKeywordsLanguagesRequestPredict) ([]EmojiLanguage, error) {
	var res []EmojiLanguage
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiURLRequestPredict struct {
	LangCode string `tl:"lang_code"`
}

func (*MessagesGetEmojiURLRequestPredict) CRC() uint32 {
	return 0xd5b10c26
}

func MessagesGetEmojiURL(ctx context.Context, m Requester, i MessagesGetEmojiURLRequestPredict) (EmojiURL, error) {
	var res EmojiURL
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchCountersRequestPredict struct {
	_           struct{}         `tl:"flags,bitflag"`
	Peer        InputPeer        `tl:"peer"`
	SavedPeerID InputPeer        `tl:"saved_peer_id,omitempty:flags:2"`
	TopMsgID    *int32           `tl:"top_msg_id,omitempty:flags:0"`
	Filters     []MessagesFilter `tl:"filters"`
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
	Peer     InputPeer `tl:"peer,omitempty:flags:1"`
	MsgID    *int32    `tl:"msg_id,omitempty:flags:1"`
	ButtonID *int32    `tl:"button_id,omitempty:flags:1"`
	URL      *string   `tl:"url,omitempty:flags:2"`
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
	WriteAllowed bool      `tl:"write_allowed,omitempty:flags:0,implicit"`
	Peer         InputPeer `tl:"peer,omitempty:flags:1"`
	MsgID        *int32    `tl:"msg_id,omitempty:flags:1"`
	ButtonID     *int32    `tl:"button_id,omitempty:flags:1"`
	URL          *string   `tl:"url,omitempty:flags:2"`
}

func (*MessagesAcceptURLAuthRequestPredict) CRC() uint32 {
	return 0xb12c7125
}

func MessagesAcceptURLAuth(ctx context.Context, m Requester, i MessagesAcceptURLAuthRequestPredict) (URLAuthResult, error) {
	var res URLAuthResult
	return res, request(ctx, m, &i, &res)
}

type MessagesHidePeerSettingsBarRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*MessagesHidePeerSettingsBarRequestPredict) CRC() uint32 {
	return 0x4facb138
}

func MessagesHidePeerSettingsBar(ctx context.Context, m Requester, i MessagesHidePeerSettingsBarRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetScheduledHistoryRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	Hash int64     `tl:"hash"`
}

func (*MessagesGetScheduledHistoryRequestPredict) CRC() uint32 {
	return 0xf516760b
}

func MessagesGetScheduledHistory(ctx context.Context, m Requester, i MessagesGetScheduledHistoryRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetScheduledMessagesRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*MessagesGetScheduledMessagesRequestPredict) CRC() uint32 {
	return 0xbdbb0464
}

func MessagesGetScheduledMessages(ctx context.Context, m Requester, i MessagesGetScheduledMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendScheduledMessagesRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*MessagesSendScheduledMessagesRequestPredict) CRC() uint32 {
	return 0xbd38850a
}

func MessagesSendScheduledMessages(ctx context.Context, m Requester, i MessagesSendScheduledMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteScheduledMessagesRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*MessagesDeleteScheduledMessagesRequestPredict) CRC() uint32 {
	return 0x59ae2b16
}

func MessagesDeleteScheduledMessages(ctx context.Context, m Requester, i MessagesDeleteScheduledMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetPollVotesRequestPredict struct {
	_      struct{}  `tl:"flags,bitflag"`
	Peer   InputPeer `tl:"peer"`
	ID     int32     `tl:"id"`
	Option *[]byte   `tl:"option,omitempty:flags:0"`
	Offset *string   `tl:"offset,omitempty:flags:1"`
	Limit  int32     `tl:"limit"`
}

func (*MessagesGetPollVotesRequestPredict) CRC() uint32 {
	return 0xb86e380e
}

func MessagesGetPollVotes(ctx context.Context, m Requester, i MessagesGetPollVotesRequestPredict) (MessagesVotesList, error) {
	var res MessagesVotesList
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleStickerSetsRequestPredict struct {
	_           struct{}          `tl:"flags,bitflag"`
	Uninstall   bool              `tl:"uninstall,omitempty:flags:0,implicit"`
	Archive     bool              `tl:"archive,omitempty:flags:1,implicit"`
	Unarchive   bool              `tl:"unarchive,omitempty:flags:2,implicit"`
	Stickersets []InputStickerSet `tl:"stickersets"`
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
	_      struct{}     `tl:"flags,bitflag"`
	ID     int32        `tl:"id"`
	Filter DialogFilter `tl:"filter,omitempty:flags:0"`
}

func (*MessagesUpdateDialogFilterRequestPredict) CRC() uint32 {
	return 0x1ad4a04a
}

func MessagesUpdateDialogFilter(ctx context.Context, m Requester, i MessagesUpdateDialogFilterRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUpdateDialogFiltersOrderRequestPredict struct {
	Order []int32 `tl:"order"`
}

func (*MessagesUpdateDialogFiltersOrderRequestPredict) CRC() uint32 {
	return 0xc563c1e4
}

func MessagesUpdateDialogFiltersOrder(ctx context.Context, m Requester, i MessagesUpdateDialogFiltersOrderRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOldFeaturedStickersRequestPredict struct {
	Offset int32 `tl:"offset"`
	Limit  int32 `tl:"limit"`
	Hash   int64 `tl:"hash"`
}

func (*MessagesGetOldFeaturedStickersRequestPredict) CRC() uint32 {
	return 0x7ed094a1
}

func MessagesGetOldFeaturedStickers(ctx context.Context, m Requester, i MessagesGetOldFeaturedStickersRequestPredict) (MessagesFeaturedStickers, error) {
	var res MessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRepliesRequestPredict struct {
	Peer       InputPeer `tl:"peer"`
	MsgID      int32     `tl:"msg_id"`
	OffsetID   int32     `tl:"offset_id"`
	OffsetDate int32     `tl:"offset_date"`
	AddOffset  int32     `tl:"add_offset"`
	Limit      int32     `tl:"limit"`
	MaxID      int32     `tl:"max_id"`
	MinID      int32     `tl:"min_id"`
	Hash       int64     `tl:"hash"`
}

func (*MessagesGetRepliesRequestPredict) CRC() uint32 {
	return 0x22ddd30c
}

func MessagesGetReplies(ctx context.Context, m Requester, i MessagesGetRepliesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDiscussionMessageRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*MessagesGetDiscussionMessageRequestPredict) CRC() uint32 {
	return 0x446972fd
}

func MessagesGetDiscussionMessage(ctx context.Context, m Requester, i MessagesGetDiscussionMessageRequestPredict) (MessagesDiscussionMessage, error) {
	var res MessagesDiscussionMessage
	return res, request(ctx, m, &i, &res)
}

type MessagesReadDiscussionRequestPredict struct {
	Peer      InputPeer `tl:"peer"`
	MsgID     int32     `tl:"msg_id"`
	ReadMaxID int32     `tl:"read_max_id"`
}

func (*MessagesReadDiscussionRequestPredict) CRC() uint32 {
	return 0xf731a9f4
}

func MessagesReadDiscussion(ctx context.Context, m Requester, i MessagesReadDiscussionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesUnpinAllMessagesRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Peer     InputPeer `tl:"peer"`
	TopMsgID *int32    `tl:"top_msg_id,omitempty:flags:0"`
}

func (*MessagesUnpinAllMessagesRequestPredict) CRC() uint32 {
	return 0xee22b9a8
}

func MessagesUnpinAllMessages(ctx context.Context, m Requester, i MessagesUnpinAllMessagesRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteChatRequestPredict struct {
	ChatID int64 `tl:"chat_id"`
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
	Revoke bool     `tl:"revoke,omitempty:flags:0,implicit"`
}

func (*MessagesDeletePhoneCallHistoryRequestPredict) CRC() uint32 {
	return 0xf9cbe409
}

func MessagesDeletePhoneCallHistory(ctx context.Context, m Requester, i MessagesDeletePhoneCallHistoryRequestPredict) (MessagesAffectedFoundMessages, error) {
	var res MessagesAffectedFoundMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckHistoryImportRequestPredict struct {
	ImportHead string `tl:"import_head"`
}

func (*MessagesCheckHistoryImportRequestPredict) CRC() uint32 {
	return 0x43fe19f3
}

func MessagesCheckHistoryImport(ctx context.Context, m Requester, i MessagesCheckHistoryImportRequestPredict) (MessagesHistoryImportParsed, error) {
	var res MessagesHistoryImportParsed
	return res, request(ctx, m, &i, &res)
}

type MessagesInitHistoryImportRequestPredict struct {
	Peer       InputPeer `tl:"peer"`
	File       InputFile `tl:"file"`
	MediaCount int32     `tl:"media_count"`
}

func (*MessagesInitHistoryImportRequestPredict) CRC() uint32 {
	return 0x34090c3b
}

func MessagesInitHistoryImport(ctx context.Context, m Requester, i MessagesInitHistoryImportRequestPredict) (MessagesHistoryImport, error) {
	var res MessagesHistoryImport
	return res, request(ctx, m, &i, &res)
}

type MessagesUploadImportedMediaRequestPredict struct {
	Peer     InputPeer  `tl:"peer"`
	ImportID int64      `tl:"import_id"`
	FileName string     `tl:"file_name"`
	Media    InputMedia `tl:"media"`
}

func (*MessagesUploadImportedMediaRequestPredict) CRC() uint32 {
	return 0x2a862092
}

func MessagesUploadImportedMedia(ctx context.Context, m Requester, i MessagesUploadImportedMediaRequestPredict) (MessageMedia, error) {
	var res MessageMedia
	return res, request(ctx, m, &i, &res)
}

type MessagesStartHistoryImportRequestPredict struct {
	Peer     InputPeer `tl:"peer"`
	ImportID int64     `tl:"import_id"`
}

func (*MessagesStartHistoryImportRequestPredict) CRC() uint32 {
	return 0xb43df344
}

func MessagesStartHistoryImport(ctx context.Context, m Requester, i MessagesStartHistoryImportRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExportedChatInvitesRequestPredict struct {
	_          struct{}  `tl:"flags,bitflag"`
	Revoked    bool      `tl:"revoked,omitempty:flags:3,implicit"`
	Peer       InputPeer `tl:"peer"`
	AdminID    InputUser `tl:"admin_id"`
	OffsetDate *int32    `tl:"offset_date,omitempty:flags:2"`
	OffsetLink *string   `tl:"offset_link,omitempty:flags:2"`
	Limit      int32     `tl:"limit"`
}

func (*MessagesGetExportedChatInvitesRequestPredict) CRC() uint32 {
	return 0xa2b5a3f6
}

func MessagesGetExportedChatInvites(ctx context.Context, m Requester, i MessagesGetExportedChatInvitesRequestPredict) (MessagesExportedChatInvites, error) {
	var res MessagesExportedChatInvites
	return res, request(ctx, m, &i, &res)
}

type MessagesGetExportedChatInviteRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	Link string    `tl:"link"`
}

func (*MessagesGetExportedChatInviteRequestPredict) CRC() uint32 {
	return 0x73746f5c
}

func MessagesGetExportedChatInvite(ctx context.Context, m Requester, i MessagesGetExportedChatInviteRequestPredict) (MessagesExportedChatInvite, error) {
	var res MessagesExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesEditExportedChatInviteRequestPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	Revoked       bool      `tl:"revoked,omitempty:flags:2,implicit"`
	Peer          InputPeer `tl:"peer"`
	Link          string    `tl:"link"`
	ExpireDate    *int32    `tl:"expire_date,omitempty:flags:0"`
	UsageLimit    *int32    `tl:"usage_limit,omitempty:flags:1"`
	RequestNeeded *bool     `tl:"request_needed,omitempty:flags:3"`
	Title         *string   `tl:"title,omitempty:flags:4"`
}

func (*MessagesEditExportedChatInviteRequestPredict) CRC() uint32 {
	return 0xbdca2f75
}

func MessagesEditExportedChatInvite(ctx context.Context, m Requester, i MessagesEditExportedChatInviteRequestPredict) (MessagesExportedChatInvite, error) {
	var res MessagesExportedChatInvite
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteRevokedExportedChatInvitesRequestPredict struct {
	Peer    InputPeer `tl:"peer"`
	AdminID InputUser `tl:"admin_id"`
}

func (*MessagesDeleteRevokedExportedChatInvitesRequestPredict) CRC() uint32 {
	return 0x56987bd5
}

func MessagesDeleteRevokedExportedChatInvites(ctx context.Context, m Requester, i MessagesDeleteRevokedExportedChatInvitesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteExportedChatInviteRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	Link string    `tl:"link"`
}

func (*MessagesDeleteExportedChatInviteRequestPredict) CRC() uint32 {
	return 0xd464a42b
}

func MessagesDeleteExportedChatInvite(ctx context.Context, m Requester, i MessagesDeleteExportedChatInviteRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAdminsWithInvitesRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*MessagesGetAdminsWithInvitesRequestPredict) CRC() uint32 {
	return 0x3920e6ef
}

func MessagesGetAdminsWithInvites(ctx context.Context, m Requester, i MessagesGetAdminsWithInvitesRequestPredict) (MessagesChatAdminsWithInvites, error) {
	var res MessagesChatAdminsWithInvites
	return res, request(ctx, m, &i, &res)
}

type MessagesGetChatInviteImportersRequestPredict struct {
	_          struct{}  `tl:"flags,bitflag"`
	Requested  bool      `tl:"requested,omitempty:flags:0,implicit"`
	Peer       InputPeer `tl:"peer"`
	Link       *string   `tl:"link,omitempty:flags:1"`
	Q          *string   `tl:"q,omitempty:flags:2"`
	OffsetDate int32     `tl:"offset_date"`
	OffsetUser InputUser `tl:"offset_user"`
	Limit      int32     `tl:"limit"`
}

func (*MessagesGetChatInviteImportersRequestPredict) CRC() uint32 {
	return 0xdf04dd4e
}

func MessagesGetChatInviteImporters(ctx context.Context, m Requester, i MessagesGetChatInviteImportersRequestPredict) (MessagesChatInviteImporters, error) {
	var res MessagesChatInviteImporters
	return res, request(ctx, m, &i, &res)
}

type MessagesSetHistoryTTLRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	Period int32     `tl:"period"`
}

func (*MessagesSetHistoryTTLRequestPredict) CRC() uint32 {
	return 0xb80e5fe4
}

func MessagesSetHistoryTTL(ctx context.Context, m Requester, i MessagesSetHistoryTTLRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckHistoryImportPeerRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*MessagesCheckHistoryImportPeerRequestPredict) CRC() uint32 {
	return 0x5dc60f03
}

func MessagesCheckHistoryImportPeer(ctx context.Context, m Requester, i MessagesCheckHistoryImportPeerRequestPredict) (MessagesCheckedHistoryImportPeer, error) {
	var res MessagesCheckedHistoryImportPeer
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatThemeRequestPredict struct {
	Peer     InputPeer `tl:"peer"`
	Emoticon string    `tl:"emoticon"`
}

func (*MessagesSetChatThemeRequestPredict) CRC() uint32 {
	return 0xe63be13f
}

func MessagesSetChatTheme(ctx context.Context, m Requester, i MessagesSetChatThemeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageReadParticipantsRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*MessagesGetMessageReadParticipantsRequestPredict) CRC() uint32 {
	return 0x31c1c44f
}

func MessagesGetMessageReadParticipants(ctx context.Context, m Requester, i MessagesGetMessageReadParticipantsRequestPredict) ([]ReadParticipantDate, error) {
	var res []ReadParticipantDate
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchResultsCalendarRequestPredict struct {
	_           struct{}       `tl:"flags,bitflag"`
	Peer        InputPeer      `tl:"peer"`
	SavedPeerID InputPeer      `tl:"saved_peer_id,omitempty:flags:2"`
	Filter      MessagesFilter `tl:"filter"`
	OffsetID    int32          `tl:"offset_id"`
	OffsetDate  int32          `tl:"offset_date"`
}

func (*MessagesGetSearchResultsCalendarRequestPredict) CRC() uint32 {
	return 0x6aa3f6bd
}

func MessagesGetSearchResultsCalendar(ctx context.Context, m Requester, i MessagesGetSearchResultsCalendarRequestPredict) (MessagesSearchResultsCalendar, error) {
	var res MessagesSearchResultsCalendar
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSearchResultsPositionsRequestPredict struct {
	_           struct{}       `tl:"flags,bitflag"`
	Peer        InputPeer      `tl:"peer"`
	SavedPeerID InputPeer      `tl:"saved_peer_id,omitempty:flags:2"`
	Filter      MessagesFilter `tl:"filter"`
	OffsetID    int32          `tl:"offset_id"`
	Limit       int32          `tl:"limit"`
}

func (*MessagesGetSearchResultsPositionsRequestPredict) CRC() uint32 {
	return 0x9c7f2f10
}

func MessagesGetSearchResultsPositions(ctx context.Context, m Requester, i MessagesGetSearchResultsPositionsRequestPredict) (MessagesSearchResultsPositions, error) {
	var res MessagesSearchResultsPositions
	return res, request(ctx, m, &i, &res)
}

type MessagesHideChatJoinRequestRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Approved bool      `tl:"approved,omitempty:flags:0,implicit"`
	Peer     InputPeer `tl:"peer"`
	UserID   InputUser `tl:"user_id"`
}

func (*MessagesHideChatJoinRequestRequestPredict) CRC() uint32 {
	return 0x7fe7e815
}

func MessagesHideChatJoinRequest(ctx context.Context, m Requester, i MessagesHideChatJoinRequestRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesHideAllChatJoinRequestsRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Approved bool      `tl:"approved,omitempty:flags:0,implicit"`
	Peer     InputPeer `tl:"peer"`
	Link     *string   `tl:"link,omitempty:flags:1"`
}

func (*MessagesHideAllChatJoinRequestsRequestPredict) CRC() uint32 {
	return 0xe085f4ea
}

func MessagesHideAllChatJoinRequests(ctx context.Context, m Requester, i MessagesHideAllChatJoinRequestsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleNoForwardsRequestPredict struct {
	Peer    InputPeer `tl:"peer"`
	Enabled bool      `tl:"enabled"`
}

func (*MessagesToggleNoForwardsRequestPredict) CRC() uint32 {
	return 0xb11eafa2
}

func MessagesToggleNoForwards(ctx context.Context, m Requester, i MessagesToggleNoForwardsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSaveDefaultSendAsRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	SendAs InputPeer `tl:"send_as"`
}

func (*MessagesSaveDefaultSendAsRequestPredict) CRC() uint32 {
	return 0xccfddf96
}

func MessagesSaveDefaultSendAs(ctx context.Context, m Requester, i MessagesSaveDefaultSendAsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesSendReactionRequestPredict struct {
	_           struct{}   `tl:"flags,bitflag"`
	Big         bool       `tl:"big,omitempty:flags:1,implicit"`
	AddToRecent bool       `tl:"add_to_recent,omitempty:flags:2,implicit"`
	Peer        InputPeer  `tl:"peer"`
	MsgID       int32      `tl:"msg_id"`
	Reaction    []Reaction `tl:"reaction,omitempty:flags:0"`
}

func (*MessagesSendReactionRequestPredict) CRC() uint32 {
	return 0xd30d78d4
}

func MessagesSendReaction(ctx context.Context, m Requester, i MessagesSendReactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessagesReactionsRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*MessagesGetMessagesReactionsRequestPredict) CRC() uint32 {
	return 0x8bba90e6
}

func MessagesGetMessagesReactions(ctx context.Context, m Requester, i MessagesGetMessagesReactionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMessageReactionsListRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Peer     InputPeer `tl:"peer"`
	ID       int32     `tl:"id"`
	Reaction Reaction  `tl:"reaction,omitempty:flags:0"`
	Offset   *string   `tl:"offset,omitempty:flags:1"`
	Limit    int32     `tl:"limit"`
}

func (*MessagesGetMessageReactionsListRequestPredict) CRC() uint32 {
	return 0x461b3f48
}

func MessagesGetMessageReactionsList(ctx context.Context, m Requester, i MessagesGetMessageReactionsListRequestPredict) (MessagesMessageReactionsList, error) {
	var res MessagesMessageReactionsList
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatAvailableReactionsRequestPredict struct {
	_                  struct{}      `tl:"flags,bitflag"`
	Peer               InputPeer     `tl:"peer"`
	AvailableReactions ChatReactions `tl:"available_reactions"`
	ReactionsLimit     *int32        `tl:"reactions_limit,omitempty:flags:0"`
}

func (*MessagesSetChatAvailableReactionsRequestPredict) CRC() uint32 {
	return 0x5a150bd4
}

func MessagesSetChatAvailableReactions(ctx context.Context, m Requester, i MessagesSetChatAvailableReactionsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAvailableReactionsRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*MessagesGetAvailableReactionsRequestPredict) CRC() uint32 {
	return 0x18dea0ac
}

func MessagesGetAvailableReactions(ctx context.Context, m Requester, i MessagesGetAvailableReactionsRequestPredict) (MessagesAvailableReactions, error) {
	var res MessagesAvailableReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesSetDefaultReactionRequestPredict struct {
	Reaction Reaction `tl:"reaction"`
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
	Peer   InputPeer          `tl:"peer,omitempty:flags:0"`
	ID     []int32            `tl:"id,omitempty:flags:0"`
	Text   []TextWithEntities `tl:"text,omitempty:flags:1"`
	ToLang string             `tl:"to_lang"`
}

func (*MessagesTranslateTextRequestPredict) CRC() uint32 {
	return 0x63183030
}

func MessagesTranslateText(ctx context.Context, m Requester, i MessagesTranslateTextRequestPredict) (MessagesTranslatedText, error) {
	var res MessagesTranslatedText
	return res, request(ctx, m, &i, &res)
}

type MessagesGetUnreadReactionsRequestPredict struct {
	_         struct{}  `tl:"flags,bitflag"`
	Peer      InputPeer `tl:"peer"`
	TopMsgID  *int32    `tl:"top_msg_id,omitempty:flags:0"`
	OffsetID  int32     `tl:"offset_id"`
	AddOffset int32     `tl:"add_offset"`
	Limit     int32     `tl:"limit"`
	MaxID     int32     `tl:"max_id"`
	MinID     int32     `tl:"min_id"`
}

func (*MessagesGetUnreadReactionsRequestPredict) CRC() uint32 {
	return 0x3223495b
}

func MessagesGetUnreadReactions(ctx context.Context, m Requester, i MessagesGetUnreadReactionsRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesReadReactionsRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Peer     InputPeer `tl:"peer"`
	TopMsgID *int32    `tl:"top_msg_id,omitempty:flags:0"`
}

func (*MessagesReadReactionsRequestPredict) CRC() uint32 {
	return 0x54aa7f8e
}

func MessagesReadReactions(ctx context.Context, m Requester, i MessagesReadReactionsRequestPredict) (MessagesAffectedHistory, error) {
	var res MessagesAffectedHistory
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchSentMediaRequestPredict struct {
	Q      string         `tl:"q"`
	Filter MessagesFilter `tl:"filter"`
	Limit  int32          `tl:"limit"`
}

func (*MessagesSearchSentMediaRequestPredict) CRC() uint32 {
	return 0x107e31a0
}

func MessagesSearchSentMedia(ctx context.Context, m Requester, i MessagesSearchSentMediaRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachMenuBotsRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetAttachMenuBotsRequestPredict) CRC() uint32 {
	return 0x16fcc2cb
}

func MessagesGetAttachMenuBots(ctx context.Context, m Requester, i MessagesGetAttachMenuBotsRequestPredict) (AttachMenuBots, error) {
	var res AttachMenuBots
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAttachMenuBotRequestPredict struct {
	Bot InputUser `tl:"bot"`
}

func (*MessagesGetAttachMenuBotRequestPredict) CRC() uint32 {
	return 0x77216192
}

func MessagesGetAttachMenuBot(ctx context.Context, m Requester, i MessagesGetAttachMenuBotRequestPredict) (AttachMenuBotsBot, error) {
	var res AttachMenuBotsBot
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleBotInAttachMenuRequestPredict struct {
	_            struct{}  `tl:"flags,bitflag"`
	WriteAllowed bool      `tl:"write_allowed,omitempty:flags:0,implicit"`
	Bot          InputUser `tl:"bot"`
	Enabled      bool      `tl:"enabled"`
}

func (*MessagesToggleBotInAttachMenuRequestPredict) CRC() uint32 {
	return 0x69f59d69
}

func MessagesToggleBotInAttachMenu(ctx context.Context, m Requester, i MessagesToggleBotInAttachMenuRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestWebViewRequestPredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	FromBotMenu bool         `tl:"from_bot_menu,omitempty:flags:4,implicit"`
	Silent      bool         `tl:"silent,omitempty:flags:5,implicit"`
	Compact     bool         `tl:"compact,omitempty:flags:7,implicit"`
	Peer        InputPeer    `tl:"peer"`
	Bot         InputUser    `tl:"bot"`
	URL         *string      `tl:"url,omitempty:flags:1"`
	StartParam  *string      `tl:"start_param,omitempty:flags:3"`
	ThemeParams DataJSON     `tl:"theme_params,omitempty:flags:2"`
	Platform    string       `tl:"platform"`
	ReplyTo     InputReplyTo `tl:"reply_to,omitempty:flags:0"`
	SendAs      InputPeer    `tl:"send_as,omitempty:flags:13"`
}

func (*MessagesRequestWebViewRequestPredict) CRC() uint32 {
	return 0x269dc2c1
}

func MessagesRequestWebView(ctx context.Context, m Requester, i MessagesRequestWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesProlongWebViewRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Silent  bool         `tl:"silent,omitempty:flags:5,implicit"`
	Peer    InputPeer    `tl:"peer"`
	Bot     InputUser    `tl:"bot"`
	QueryID int64        `tl:"query_id"`
	ReplyTo InputReplyTo `tl:"reply_to,omitempty:flags:0"`
	SendAs  InputPeer    `tl:"send_as,omitempty:flags:13"`
}

func (*MessagesProlongWebViewRequestPredict) CRC() uint32 {
	return 0xb0d81a83
}

func MessagesProlongWebView(ctx context.Context, m Requester, i MessagesProlongWebViewRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestSimpleWebViewRequestPredict struct {
	_                 struct{}  `tl:"flags,bitflag"`
	FromSwitchWebview bool      `tl:"from_switch_webview,omitempty:flags:1,implicit"`
	FromSideMenu      bool      `tl:"from_side_menu,omitempty:flags:2,implicit"`
	Compact           bool      `tl:"compact,omitempty:flags:7,implicit"`
	Bot               InputUser `tl:"bot"`
	URL               *string   `tl:"url,omitempty:flags:3"`
	StartParam        *string   `tl:"start_param,omitempty:flags:4"`
	ThemeParams       DataJSON  `tl:"theme_params,omitempty:flags:0"`
	Platform          string    `tl:"platform"`
}

func (*MessagesRequestSimpleWebViewRequestPredict) CRC() uint32 {
	return 0x413a3e73
}

func MessagesRequestSimpleWebView(ctx context.Context, m Requester, i MessagesRequestSimpleWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesSendWebViewResultMessageRequestPredict struct {
	BotQueryID string               `tl:"bot_query_id"`
	Result     InputBotInlineResult `tl:"result"`
}

func (*MessagesSendWebViewResultMessageRequestPredict) CRC() uint32 {
	return 0x0a4314f5
}

func MessagesSendWebViewResultMessage(ctx context.Context, m Requester, i MessagesSendWebViewResultMessageRequestPredict) (WebViewMessageSent, error) {
	var res WebViewMessageSent
	return res, request(ctx, m, &i, &res)
}

type MessagesSendWebViewDataRequestPredict struct {
	Bot        InputUser `tl:"bot"`
	RandomID   int64     `tl:"random_id"`
	ButtonText string    `tl:"button_text"`
	Data       string    `tl:"data"`
}

func (*MessagesSendWebViewDataRequestPredict) CRC() uint32 {
	return 0xdc0242c8
}

func MessagesSendWebViewData(ctx context.Context, m Requester, i MessagesSendWebViewDataRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesTranscribeAudioRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*MessagesTranscribeAudioRequestPredict) CRC() uint32 {
	return 0x269e9a49
}

func MessagesTranscribeAudio(ctx context.Context, m Requester, i MessagesTranscribeAudioRequestPredict) (MessagesTranscribedAudio, error) {
	var res MessagesTranscribedAudio
	return res, request(ctx, m, &i, &res)
}

type MessagesRateTranscribedAudioRequestPredict struct {
	Peer            InputPeer `tl:"peer"`
	MsgID           int32     `tl:"msg_id"`
	TranscriptionID int64     `tl:"transcription_id"`
	Good            bool      `tl:"good"`
}

func (*MessagesRateTranscribedAudioRequestPredict) CRC() uint32 {
	return 0x7f1d072f
}

func MessagesRateTranscribedAudio(ctx context.Context, m Requester, i MessagesRateTranscribedAudioRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetCustomEmojiDocumentsRequestPredict struct {
	DocumentID []int64 `tl:"document_id"`
}

func (*MessagesGetCustomEmojiDocumentsRequestPredict) CRC() uint32 {
	return 0xd9ab0f54
}

func MessagesGetCustomEmojiDocuments(ctx context.Context, m Requester, i MessagesGetCustomEmojiDocumentsRequestPredict) ([]Document, error) {
	var res []Document
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStickersRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetEmojiStickersRequestPredict) CRC() uint32 {
	return 0xfbfca18f
}

func MessagesGetEmojiStickers(ctx context.Context, m Requester, i MessagesGetEmojiStickersRequestPredict) (MessagesAllStickers, error) {
	var res MessagesAllStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFeaturedEmojiStickersRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetFeaturedEmojiStickersRequestPredict) CRC() uint32 {
	return 0x0ecf6736
}

func MessagesGetFeaturedEmojiStickers(ctx context.Context, m Requester, i MessagesGetFeaturedEmojiStickersRequestPredict) (MessagesFeaturedStickers, error) {
	var res MessagesFeaturedStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesReportReactionRequestPredict struct {
	Peer         InputPeer `tl:"peer"`
	ID           int32     `tl:"id"`
	ReactionPeer InputPeer `tl:"reaction_peer"`
}

func (*MessagesReportReactionRequestPredict) CRC() uint32 {
	return 0x3f64c076
}

func MessagesReportReaction(ctx context.Context, m Requester, i MessagesReportReactionRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetTopReactionsRequestPredict struct {
	Limit int32 `tl:"limit"`
	Hash  int64 `tl:"hash"`
}

func (*MessagesGetTopReactionsRequestPredict) CRC() uint32 {
	return 0xbb8125ba
}

func MessagesGetTopReactions(ctx context.Context, m Requester, i MessagesGetTopReactionsRequestPredict) (MessagesReactions, error) {
	var res MessagesReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesGetRecentReactionsRequestPredict struct {
	Limit int32 `tl:"limit"`
	Hash  int64 `tl:"hash"`
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
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*MessagesGetExtendedMediaRequestPredict) CRC() uint32 {
	return 0x84f80814
}

func MessagesGetExtendedMedia(ctx context.Context, m Requester, i MessagesGetExtendedMediaRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesSetDefaultHistoryTTLRequestPredict struct {
	Period int32 `tl:"period"`
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
	Peer           InputPeer   `tl:"peer"`
	MsgID          int32       `tl:"msg_id"`
	ButtonID       int32       `tl:"button_id"`
	RequestedPeers []InputPeer `tl:"requested_peers"`
}

func (*MessagesSendBotRequestedPeerRequestPredict) CRC() uint32 {
	return 0x91b2d060
}

func MessagesSendBotRequestedPeer(ctx context.Context, m Requester, i MessagesSendBotRequestedPeerRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiGroupsRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*MessagesGetEmojiGroupsRequestPredict) CRC() uint32 {
	return 0x7488ce5b
}

func MessagesGetEmojiGroups(ctx context.Context, m Requester, i MessagesGetEmojiGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStatusGroupsRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*MessagesGetEmojiStatusGroupsRequestPredict) CRC() uint32 {
	return 0x2ecd56cd
}

func MessagesGetEmojiStatusGroups(ctx context.Context, m Requester, i MessagesGetEmojiStatusGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiProfilePhotoGroupsRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*MessagesGetEmojiProfilePhotoGroupsRequestPredict) CRC() uint32 {
	return 0x21a548f3
}

func MessagesGetEmojiProfilePhotoGroups(ctx context.Context, m Requester, i MessagesGetEmojiProfilePhotoGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesSearchCustomEmojiRequestPredict struct {
	Emoticon string `tl:"emoticon"`
	Hash     int64  `tl:"hash"`
}

func (*MessagesSearchCustomEmojiRequestPredict) CRC() uint32 {
	return 0x2c11c0d7
}

func MessagesSearchCustomEmoji(ctx context.Context, m Requester, i MessagesSearchCustomEmojiRequestPredict) (EmojiList, error) {
	var res EmojiList
	return res, request(ctx, m, &i, &res)
}

type MessagesTogglePeerTranslationsRequestPredict struct {
	_        struct{}  `tl:"flags,bitflag"`
	Disabled bool      `tl:"disabled,omitempty:flags:0,implicit"`
	Peer     InputPeer `tl:"peer"`
}

func (*MessagesTogglePeerTranslationsRequestPredict) CRC() uint32 {
	return 0xe47cb579
}

func MessagesTogglePeerTranslations(ctx context.Context, m Requester, i MessagesTogglePeerTranslationsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetBotAppRequestPredict struct {
	App  InputBotApp `tl:"app"`
	Hash int64       `tl:"hash"`
}

func (*MessagesGetBotAppRequestPredict) CRC() uint32 {
	return 0x34fdc5c3
}

func MessagesGetBotApp(ctx context.Context, m Requester, i MessagesGetBotAppRequestPredict) (MessagesBotApp, error) {
	var res MessagesBotApp
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestAppWebViewRequestPredict struct {
	_            struct{}    `tl:"flags,bitflag"`
	WriteAllowed bool        `tl:"write_allowed,omitempty:flags:0,implicit"`
	Compact      bool        `tl:"compact,omitempty:flags:7,implicit"`
	Peer         InputPeer   `tl:"peer"`
	App          InputBotApp `tl:"app"`
	StartParam   *string     `tl:"start_param,omitempty:flags:1"`
	ThemeParams  DataJSON    `tl:"theme_params,omitempty:flags:2"`
	Platform     string      `tl:"platform"`
}

func (*MessagesRequestAppWebViewRequestPredict) CRC() uint32 {
	return 0x53618bce
}

func MessagesRequestAppWebView(ctx context.Context, m Requester, i MessagesRequestAppWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type MessagesSetChatWallPaperRequestPredict struct {
	_         struct{}          `tl:"flags,bitflag"`
	ForBoth   bool              `tl:"for_both,omitempty:flags:3,implicit"`
	Revert    bool              `tl:"revert,omitempty:flags:4,implicit"`
	Peer      InputPeer         `tl:"peer"`
	Wallpaper InputWallPaper    `tl:"wallpaper,omitempty:flags:0"`
	Settings  WallPaperSettings `tl:"settings,omitempty:flags:2"`
	ID        *int32            `tl:"id,omitempty:flags:1"`
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
	ExcludeFeatured bool     `tl:"exclude_featured,omitempty:flags:0,implicit"`
	Q               string   `tl:"q"`
	Hash            int64    `tl:"hash"`
}

func (*MessagesSearchEmojiStickerSetsRequestPredict) CRC() uint32 {
	return 0x92b4494c
}

func MessagesSearchEmojiStickerSets(ctx context.Context, m Requester, i MessagesSearchEmojiStickerSetsRequestPredict) (MessagesFoundStickerSets, error) {
	var res MessagesFoundStickerSets
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedDialogsRequestPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	ExcludePinned bool      `tl:"exclude_pinned,omitempty:flags:0,implicit"`
	OffsetDate    int32     `tl:"offset_date"`
	OffsetID      int32     `tl:"offset_id"`
	OffsetPeer    InputPeer `tl:"offset_peer"`
	Limit         int32     `tl:"limit"`
	Hash          int64     `tl:"hash"`
}

func (*MessagesGetSavedDialogsRequestPredict) CRC() uint32 {
	return 0x5381d21a
}

func MessagesGetSavedDialogs(ctx context.Context, m Requester, i MessagesGetSavedDialogsRequestPredict) (MessagesSavedDialogs, error) {
	var res MessagesSavedDialogs
	return res, request(ctx, m, &i, &res)
}

type MessagesGetSavedHistoryRequestPredict struct {
	Peer       InputPeer `tl:"peer"`
	OffsetID   int32     `tl:"offset_id"`
	OffsetDate int32     `tl:"offset_date"`
	AddOffset  int32     `tl:"add_offset"`
	Limit      int32     `tl:"limit"`
	MaxID      int32     `tl:"max_id"`
	MinID      int32     `tl:"min_id"`
	Hash       int64     `tl:"hash"`
}

func (*MessagesGetSavedHistoryRequestPredict) CRC() uint32 {
	return 0x3d9a414d
}

func MessagesGetSavedHistory(ctx context.Context, m Requester, i MessagesGetSavedHistoryRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteSavedHistoryRequestPredict struct {
	_       struct{}  `tl:"flags,bitflag"`
	Peer    InputPeer `tl:"peer"`
	MaxID   int32     `tl:"max_id"`
	MinDate *int32    `tl:"min_date,omitempty:flags:2"`
	MaxDate *int32    `tl:"max_date,omitempty:flags:3"`
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
	_      struct{}        `tl:"flags,bitflag"`
	Pinned bool            `tl:"pinned,omitempty:flags:0,implicit"`
	Peer   InputDialogPeer `tl:"peer"`
}

func (*MessagesToggleSavedDialogPinRequestPredict) CRC() uint32 {
	return 0xac81bbde
}

func MessagesToggleSavedDialogPin(ctx context.Context, m Requester, i MessagesToggleSavedDialogPinRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderPinnedSavedDialogsRequestPredict struct {
	_     struct{}          `tl:"flags,bitflag"`
	Force bool              `tl:"force,omitempty:flags:0,implicit"`
	Order []InputDialogPeer `tl:"order"`
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
	Peer InputPeer `tl:"peer,omitempty:flags:0"`
	Hash int64     `tl:"hash"`
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
	Reaction Reaction `tl:"reaction"`
	Title    *string  `tl:"title,omitempty:flags:0"`
}

func (*MessagesUpdateSavedReactionTagRequestPredict) CRC() uint32 {
	return 0x60297dec
}

func MessagesUpdateSavedReactionTag(ctx context.Context, m Requester, i MessagesUpdateSavedReactionTagRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetDefaultTagReactionsRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetDefaultTagReactionsRequestPredict) CRC() uint32 {
	return 0xbdf93428
}

func MessagesGetDefaultTagReactions(ctx context.Context, m Requester, i MessagesGetDefaultTagReactionsRequestPredict) (MessagesReactions, error) {
	var res MessagesReactions
	return res, request(ctx, m, &i, &res)
}

type MessagesGetOutboxReadDateRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*MessagesGetOutboxReadDateRequestPredict) CRC() uint32 {
	return 0x8c4bfe5d
}

func MessagesGetOutboxReadDate(ctx context.Context, m Requester, i MessagesGetOutboxReadDateRequestPredict) (OutboxReadDate, error) {
	var res OutboxReadDate
	return res, request(ctx, m, &i, &res)
}

type MessagesGetQuickRepliesRequestPredict struct {
	Hash int64 `tl:"hash"`
}

func (*MessagesGetQuickRepliesRequestPredict) CRC() uint32 {
	return 0xd483f2a8
}

func MessagesGetQuickReplies(ctx context.Context, m Requester, i MessagesGetQuickRepliesRequestPredict) (MessagesQuickReplies, error) {
	var res MessagesQuickReplies
	return res, request(ctx, m, &i, &res)
}

type MessagesReorderQuickRepliesRequestPredict struct {
	Order []int32 `tl:"order"`
}

func (*MessagesReorderQuickRepliesRequestPredict) CRC() uint32 {
	return 0x60331907
}

func MessagesReorderQuickReplies(ctx context.Context, m Requester, i MessagesReorderQuickRepliesRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesCheckQuickReplyShortcutRequestPredict struct {
	Shortcut string `tl:"shortcut"`
}

func (*MessagesCheckQuickReplyShortcutRequestPredict) CRC() uint32 {
	return 0xf1d0fbd3
}

func MessagesCheckQuickReplyShortcut(ctx context.Context, m Requester, i MessagesCheckQuickReplyShortcutRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesEditQuickReplyShortcutRequestPredict struct {
	ShortcutID int32  `tl:"shortcut_id"`
	Shortcut   string `tl:"shortcut"`
}

func (*MessagesEditQuickReplyShortcutRequestPredict) CRC() uint32 {
	return 0x5c003cef
}

func MessagesEditQuickReplyShortcut(ctx context.Context, m Requester, i MessagesEditQuickReplyShortcutRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteQuickReplyShortcutRequestPredict struct {
	ShortcutID int32 `tl:"shortcut_id"`
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
	ShortcutID int32    `tl:"shortcut_id"`
	ID         []int32  `tl:"id,omitempty:flags:0"`
	Hash       int64    `tl:"hash"`
}

func (*MessagesGetQuickReplyMessagesRequestPredict) CRC() uint32 {
	return 0x94a495c3
}

func MessagesGetQuickReplyMessages(ctx context.Context, m Requester, i MessagesGetQuickReplyMessagesRequestPredict) (MessagesMessages, error) {
	var res MessagesMessages
	return res, request(ctx, m, &i, &res)
}

type MessagesSendQuickReplyMessagesRequestPredict struct {
	Peer       InputPeer `tl:"peer"`
	ShortcutID int32     `tl:"shortcut_id"`
	ID         []int32   `tl:"id"`
	RandomID   []int64   `tl:"random_id"`
}

func (*MessagesSendQuickReplyMessagesRequestPredict) CRC() uint32 {
	return 0x6c750de1
}

func MessagesSendQuickReplyMessages(ctx context.Context, m Requester, i MessagesSendQuickReplyMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteQuickReplyMessagesRequestPredict struct {
	ShortcutID int32   `tl:"shortcut_id"`
	ID         []int32 `tl:"id"`
}

func (*MessagesDeleteQuickReplyMessagesRequestPredict) CRC() uint32 {
	return 0xe105e910
}

func MessagesDeleteQuickReplyMessages(ctx context.Context, m Requester, i MessagesDeleteQuickReplyMessagesRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesToggleDialogFilterTagsRequestPredict struct {
	Enabled bool `tl:"enabled"`
}

func (*MessagesToggleDialogFilterTagsRequestPredict) CRC() uint32 {
	return 0xfd2dda49
}

func MessagesToggleDialogFilterTags(ctx context.Context, m Requester, i MessagesToggleDialogFilterTagsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type MessagesGetMyStickersRequestPredict struct {
	OffsetID int64 `tl:"offset_id"`
	Limit    int32 `tl:"limit"`
}

func (*MessagesGetMyStickersRequestPredict) CRC() uint32 {
	return 0xd0b5e1fc
}

func MessagesGetMyStickers(ctx context.Context, m Requester, i MessagesGetMyStickersRequestPredict) (MessagesMyStickers, error) {
	var res MessagesMyStickers
	return res, request(ctx, m, &i, &res)
}

type MessagesGetEmojiStickerGroupsRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*MessagesGetEmojiStickerGroupsRequestPredict) CRC() uint32 {
	return 0x1dd840f5
}

func MessagesGetEmojiStickerGroups(ctx context.Context, m Requester, i MessagesGetEmojiStickerGroupsRequestPredict) (MessagesEmojiGroups, error) {
	var res MessagesEmojiGroups
	return res, request(ctx, m, &i, &res)
}

type MessagesGetAvailableEffectsRequestPredict struct {
	Hash int32 `tl:"hash"`
}

func (*MessagesGetAvailableEffectsRequestPredict) CRC() uint32 {
	return 0xdea20a39
}

func MessagesGetAvailableEffects(ctx context.Context, m Requester, i MessagesGetAvailableEffectsRequestPredict) (MessagesAvailableEffects, error) {
	var res MessagesAvailableEffects
	return res, request(ctx, m, &i, &res)
}

type MessagesEditFactCheckRequestPredict struct {
	Peer  InputPeer        `tl:"peer"`
	MsgID int32            `tl:"msg_id"`
	Text  TextWithEntities `tl:"text"`
}

func (*MessagesEditFactCheckRequestPredict) CRC() uint32 {
	return 0x0589ee75
}

func MessagesEditFactCheck(ctx context.Context, m Requester, i MessagesEditFactCheckRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesDeleteFactCheckRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*MessagesDeleteFactCheckRequestPredict) CRC() uint32 {
	return 0xd1da940c
}

func MessagesDeleteFactCheck(ctx context.Context, m Requester, i MessagesDeleteFactCheckRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type MessagesGetFactCheckRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID []int32   `tl:"msg_id"`
}

func (*MessagesGetFactCheckRequestPredict) CRC() uint32 {
	return 0xb9cdc5ee
}

func MessagesGetFactCheck(ctx context.Context, m Requester, i MessagesGetFactCheckRequestPredict) ([]FactCheck, error) {
	var res []FactCheck
	return res, request(ctx, m, &i, &res)
}

type MessagesRequestMainWebViewRequestPredict struct {
	_           struct{}  `tl:"flags,bitflag"`
	Compact     bool      `tl:"compact,omitempty:flags:7,implicit"`
	Peer        InputPeer `tl:"peer"`
	Bot         InputUser `tl:"bot"`
	StartParam  *string   `tl:"start_param,omitempty:flags:1"`
	ThemeParams DataJSON  `tl:"theme_params,omitempty:flags:0"`
	Platform    string    `tl:"platform"`
}

func (*MessagesRequestMainWebViewRequestPredict) CRC() uint32 {
	return 0xc9e01e7b
}

func MessagesRequestMainWebView(ctx context.Context, m Requester, i MessagesRequestMainWebViewRequestPredict) (WebViewResult, error) {
	var res WebViewResult
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPaymentFormRequestPredict struct {
	_           struct{}     `tl:"flags,bitflag"`
	Invoice     InputInvoice `tl:"invoice"`
	ThemeParams DataJSON     `tl:"theme_params,omitempty:flags:0"`
}

func (*PaymentsGetPaymentFormRequestPredict) CRC() uint32 {
	return 0x37148dbb
}

func PaymentsGetPaymentForm(ctx context.Context, m Requester, i PaymentsGetPaymentFormRequestPredict) (PaymentsPaymentForm, error) {
	var res PaymentsPaymentForm
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetPaymentReceiptRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*PaymentsGetPaymentReceiptRequestPredict) CRC() uint32 {
	return 0x2478d1cc
}

func PaymentsGetPaymentReceipt(ctx context.Context, m Requester, i PaymentsGetPaymentReceiptRequestPredict) (PaymentsPaymentReceipt, error) {
	var res PaymentsPaymentReceipt
	return res, request(ctx, m, &i, &res)
}

type PaymentsValidateRequestedInfoRequestPredict struct {
	_       struct{}             `tl:"flags,bitflag"`
	Save    bool                 `tl:"save,omitempty:flags:0,implicit"`
	Invoice InputInvoice         `tl:"invoice"`
	Info    PaymentRequestedInfo `tl:"info"`
}

func (*PaymentsValidateRequestedInfoRequestPredict) CRC() uint32 {
	return 0xb6c8f12b
}

func PaymentsValidateRequestedInfo(ctx context.Context, m Requester, i PaymentsValidateRequestedInfoRequestPredict) (PaymentsValidatedRequestedInfo, error) {
	var res PaymentsValidatedRequestedInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsSendPaymentFormRequestPredict struct {
	_                struct{}                `tl:"flags,bitflag"`
	FormID           int64                   `tl:"form_id"`
	Invoice          InputInvoice            `tl:"invoice"`
	RequestedInfoID  *string                 `tl:"requested_info_id,omitempty:flags:0"`
	ShippingOptionID *string                 `tl:"shipping_option_id,omitempty:flags:1"`
	Credentials      InputPaymentCredentials `tl:"credentials"`
	TipAmount        *int64                  `tl:"tip_amount,omitempty:flags:2"`
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
	Credentials bool     `tl:"credentials,omitempty:flags:0,implicit"`
	Info        bool     `tl:"info,omitempty:flags:1,implicit"`
}

func (*PaymentsClearSavedInfoRequestPredict) CRC() uint32 {
	return 0xd83d70c1
}

func PaymentsClearSavedInfo(ctx context.Context, m Requester, i PaymentsClearSavedInfoRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetBankCardDataRequestPredict struct {
	Number string `tl:"number"`
}

func (*PaymentsGetBankCardDataRequestPredict) CRC() uint32 {
	return 0x2e79d779
}

func PaymentsGetBankCardData(ctx context.Context, m Requester, i PaymentsGetBankCardDataRequestPredict) (PaymentsBankCardData, error) {
	var res PaymentsBankCardData
	return res, request(ctx, m, &i, &res)
}

type PaymentsExportInvoiceRequestPredict struct {
	InvoiceMedia InputMedia `tl:"invoice_media"`
}

func (*PaymentsExportInvoiceRequestPredict) CRC() uint32 {
	return 0x0f91b065
}

func PaymentsExportInvoice(ctx context.Context, m Requester, i PaymentsExportInvoiceRequestPredict) (PaymentsExportedInvoice, error) {
	var res PaymentsExportedInvoice
	return res, request(ctx, m, &i, &res)
}

type PaymentsAssignAppStoreTransactionRequestPredict struct {
	Receipt []byte                   `tl:"receipt"`
	Purpose InputStorePaymentPurpose `tl:"purpose"`
}

func (*PaymentsAssignAppStoreTransactionRequestPredict) CRC() uint32 {
	return 0x80ed747d
}

func PaymentsAssignAppStoreTransaction(ctx context.Context, m Requester, i PaymentsAssignAppStoreTransactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsAssignPlayMarketTransactionRequestPredict struct {
	Receipt DataJSON                 `tl:"receipt"`
	Purpose InputStorePaymentPurpose `tl:"purpose"`
}

func (*PaymentsAssignPlayMarketTransactionRequestPredict) CRC() uint32 {
	return 0xdffd50d3
}

func PaymentsAssignPlayMarketTransaction(ctx context.Context, m Requester, i PaymentsAssignPlayMarketTransactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsCanPurchasePremiumRequestPredict struct {
	Purpose InputStorePaymentPurpose `tl:"purpose"`
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
	BoostPeer InputPeer `tl:"boost_peer,omitempty:flags:0"`
}

func (*PaymentsGetPremiumGiftCodeOptionsRequestPredict) CRC() uint32 {
	return 0x2757ba54
}

func PaymentsGetPremiumGiftCodeOptions(ctx context.Context, m Requester, i PaymentsGetPremiumGiftCodeOptionsRequestPredict) ([]PremiumGiftCodeOption, error) {
	var res []PremiumGiftCodeOption
	return res, request(ctx, m, &i, &res)
}

type PaymentsCheckGiftCodeRequestPredict struct {
	Slug string `tl:"slug"`
}

func (*PaymentsCheckGiftCodeRequestPredict) CRC() uint32 {
	return 0x8e51b4c1
}

func PaymentsCheckGiftCode(ctx context.Context, m Requester, i PaymentsCheckGiftCodeRequestPredict) (PaymentsCheckedGiftCode, error) {
	var res PaymentsCheckedGiftCode
	return res, request(ctx, m, &i, &res)
}

type PaymentsApplyGiftCodeRequestPredict struct {
	Slug string `tl:"slug"`
}

func (*PaymentsApplyGiftCodeRequestPredict) CRC() uint32 {
	return 0xf6e26854
}

func PaymentsApplyGiftCode(ctx context.Context, m Requester, i PaymentsApplyGiftCodeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetGiveawayInfoRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MsgID int32     `tl:"msg_id"`
}

func (*PaymentsGetGiveawayInfoRequestPredict) CRC() uint32 {
	return 0xf4239425
}

func PaymentsGetGiveawayInfo(ctx context.Context, m Requester, i PaymentsGetGiveawayInfoRequestPredict) (PaymentsGiveawayInfo, error) {
	var res PaymentsGiveawayInfo
	return res, request(ctx, m, &i, &res)
}

type PaymentsLaunchPrepaidGiveawayRequestPredict struct {
	Peer       InputPeer                `tl:"peer"`
	GiveawayID int64                    `tl:"giveaway_id"`
	Purpose    InputStorePaymentPurpose `tl:"purpose"`
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
	Peer InputPeer `tl:"peer"`
}

func (*PaymentsGetStarsStatusRequestPredict) CRC() uint32 {
	return 0x104fcfa7
}

func PaymentsGetStarsStatus(ctx context.Context, m Requester, i PaymentsGetStarsStatusRequestPredict) (PaymentsStarsStatus, error) {
	var res PaymentsStarsStatus
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsTransactionsRequestPredict struct {
	_         struct{}  `tl:"flags,bitflag"`
	Inbound   bool      `tl:"inbound,omitempty:flags:0,implicit"`
	Outbound  bool      `tl:"outbound,omitempty:flags:1,implicit"`
	Ascending bool      `tl:"ascending,omitempty:flags:2,implicit"`
	Peer      InputPeer `tl:"peer"`
	Offset    string    `tl:"offset"`
	Limit     int32     `tl:"limit"`
}

func (*PaymentsGetStarsTransactionsRequestPredict) CRC() uint32 {
	return 0x97938d5a
}

func PaymentsGetStarsTransactions(ctx context.Context, m Requester, i PaymentsGetStarsTransactionsRequestPredict) (PaymentsStarsStatus, error) {
	var res PaymentsStarsStatus
	return res, request(ctx, m, &i, &res)
}

type PaymentsSendStarsFormRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	FormID  int64        `tl:"form_id"`
	Invoice InputInvoice `tl:"invoice"`
}

func (*PaymentsSendStarsFormRequestPredict) CRC() uint32 {
	return 0x02bb731d
}

func PaymentsSendStarsForm(ctx context.Context, m Requester, i PaymentsSendStarsFormRequestPredict) (PaymentsPaymentResult, error) {
	var res PaymentsPaymentResult
	return res, request(ctx, m, &i, &res)
}

type PaymentsRefundStarsChargeRequestPredict struct {
	UserID   InputUser `tl:"user_id"`
	ChargeID string    `tl:"charge_id"`
}

func (*PaymentsRefundStarsChargeRequestPredict) CRC() uint32 {
	return 0x25ae8f4a
}

func PaymentsRefundStarsCharge(ctx context.Context, m Requester, i PaymentsRefundStarsChargeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsRevenueStatsRequestPredict struct {
	_    struct{}  `tl:"flags,bitflag"`
	Dark bool      `tl:"dark,omitempty:flags:0,implicit"`
	Peer InputPeer `tl:"peer"`
}

func (*PaymentsGetStarsRevenueStatsRequestPredict) CRC() uint32 {
	return 0xd91ffad6
}

func PaymentsGetStarsRevenueStats(ctx context.Context, m Requester, i PaymentsGetStarsRevenueStatsRequestPredict) (PaymentsStarsRevenueStats, error) {
	var res PaymentsStarsRevenueStats
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsRevenueWithdrawalURLRequestPredict struct {
	Peer     InputPeer             `tl:"peer"`
	Stars    int64                 `tl:"stars"`
	Password InputCheckPasswordSRP `tl:"password"`
}

func (*PaymentsGetStarsRevenueWithdrawalURLRequestPredict) CRC() uint32 {
	return 0x13bbe8b3
}

func PaymentsGetStarsRevenueWithdrawalURL(ctx context.Context, m Requester, i PaymentsGetStarsRevenueWithdrawalURLRequestPredict) (PaymentsStarsRevenueWithdrawalURL, error) {
	var res PaymentsStarsRevenueWithdrawalURL
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsRevenueAdsAccountURLRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*PaymentsGetStarsRevenueAdsAccountURLRequestPredict) CRC() uint32 {
	return 0xd1d7efc5
}

func PaymentsGetStarsRevenueAdsAccountURL(ctx context.Context, m Requester, i PaymentsGetStarsRevenueAdsAccountURLRequestPredict) (PaymentsStarsRevenueAdsAccountURL, error) {
	var res PaymentsStarsRevenueAdsAccountURL
	return res, request(ctx, m, &i, &res)
}

type PaymentsGetStarsTransactionsByIDRequestPredict struct {
	Peer InputPeer               `tl:"peer"`
	ID   []InputStarsTransaction `tl:"id"`
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
	UserID InputUser `tl:"user_id,omitempty:flags:0"`
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
	_        struct{}          `tl:"flags,bitflag"`
	Video    bool              `tl:"video,omitempty:flags:0,implicit"`
	UserID   InputUser         `tl:"user_id"`
	RandomID int32             `tl:"random_id"`
	GAHash   []byte            `tl:"g_a_hash"`
	Protocol PhoneCallProtocol `tl:"protocol"`
}

func (*PhoneRequestCallRequestPredict) CRC() uint32 {
	return 0x42ff96ed
}

func PhoneRequestCall(ctx context.Context, m Requester, i PhoneRequestCallRequestPredict) (PhonePhoneCall, error) {
	var res PhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneAcceptCallRequestPredict struct {
	Peer     InputPhoneCall    `tl:"peer"`
	GB       []byte            `tl:"g_b"`
	Protocol PhoneCallProtocol `tl:"protocol"`
}

func (*PhoneAcceptCallRequestPredict) CRC() uint32 {
	return 0x3bd2b4a0
}

func PhoneAcceptCall(ctx context.Context, m Requester, i PhoneAcceptCallRequestPredict) (PhonePhoneCall, error) {
	var res PhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneConfirmCallRequestPredict struct {
	Peer           InputPhoneCall    `tl:"peer"`
	GA             []byte            `tl:"g_a"`
	KeyFingerprint int64             `tl:"key_fingerprint"`
	Protocol       PhoneCallProtocol `tl:"protocol"`
}

func (*PhoneConfirmCallRequestPredict) CRC() uint32 {
	return 0x2efe1722
}

func PhoneConfirmCall(ctx context.Context, m Requester, i PhoneConfirmCallRequestPredict) (PhonePhoneCall, error) {
	var res PhonePhoneCall
	return res, request(ctx, m, &i, &res)
}

type PhoneReceivedCallRequestPredict struct {
	Peer InputPhoneCall `tl:"peer"`
}

func (*PhoneReceivedCallRequestPredict) CRC() uint32 {
	return 0x17d54f61
}

func PhoneReceivedCall(ctx context.Context, m Requester, i PhoneReceivedCallRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneDiscardCallRequestPredict struct {
	_            struct{}               `tl:"flags,bitflag"`
	Video        bool                   `tl:"video,omitempty:flags:0,implicit"`
	Peer         InputPhoneCall         `tl:"peer"`
	Duration     int32                  `tl:"duration"`
	Reason       PhoneCallDiscardReason `tl:"reason"`
	ConnectionID int64                  `tl:"connection_id"`
}

func (*PhoneDiscardCallRequestPredict) CRC() uint32 {
	return 0xb2cbc1c0
}

func PhoneDiscardCall(ctx context.Context, m Requester, i PhoneDiscardCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneSetCallRatingRequestPredict struct {
	_              struct{}       `tl:"flags,bitflag"`
	UserInitiative bool           `tl:"user_initiative,omitempty:flags:0,implicit"`
	Peer           InputPhoneCall `tl:"peer"`
	Rating         int32          `tl:"rating"`
	Comment        string         `tl:"comment"`
}

func (*PhoneSetCallRatingRequestPredict) CRC() uint32 {
	return 0x59ead627
}

func PhoneSetCallRating(ctx context.Context, m Requester, i PhoneSetCallRatingRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveCallDebugRequestPredict struct {
	Peer  InputPhoneCall `tl:"peer"`
	Debug DataJSON       `tl:"debug"`
}

func (*PhoneSaveCallDebugRequestPredict) CRC() uint32 {
	return 0x277add7e
}

func PhoneSaveCallDebug(ctx context.Context, m Requester, i PhoneSaveCallDebugRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneSendSignalingDataRequestPredict struct {
	Peer InputPhoneCall `tl:"peer"`
	Data []byte         `tl:"data"`
}

func (*PhoneSendSignalingDataRequestPredict) CRC() uint32 {
	return 0xff7a9383
}

func PhoneSendSignalingData(ctx context.Context, m Requester, i PhoneSendSignalingDataRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneCreateGroupCallRequestPredict struct {
	_            struct{}  `tl:"flags,bitflag"`
	RtmpStream   bool      `tl:"rtmp_stream,omitempty:flags:2,implicit"`
	Peer         InputPeer `tl:"peer"`
	RandomID     int32     `tl:"random_id"`
	Title        *string   `tl:"title,omitempty:flags:0"`
	ScheduleDate *int32    `tl:"schedule_date,omitempty:flags:1"`
}

func (*PhoneCreateGroupCallRequestPredict) CRC() uint32 {
	return 0x48cdc6d8
}

func PhoneCreateGroupCall(ctx context.Context, m Requester, i PhoneCreateGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneJoinGroupCallRequestPredict struct {
	_            struct{}       `tl:"flags,bitflag"`
	Muted        bool           `tl:"muted,omitempty:flags:0,implicit"`
	VideoStopped bool           `tl:"video_stopped,omitempty:flags:2,implicit"`
	Call         InputGroupCall `tl:"call"`
	JoinAs       InputPeer      `tl:"join_as"`
	InviteHash   *string        `tl:"invite_hash,omitempty:flags:1"`
	Params       DataJSON       `tl:"params"`
}

func (*PhoneJoinGroupCallRequestPredict) CRC() uint32 {
	return 0xb132ff7b
}

func PhoneJoinGroupCall(ctx context.Context, m Requester, i PhoneJoinGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneLeaveGroupCallRequestPredict struct {
	Call   InputGroupCall `tl:"call"`
	Source int32          `tl:"source"`
}

func (*PhoneLeaveGroupCallRequestPredict) CRC() uint32 {
	return 0x500377f9
}

func PhoneLeaveGroupCall(ctx context.Context, m Requester, i PhoneLeaveGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneInviteToGroupCallRequestPredict struct {
	Call  InputGroupCall `tl:"call"`
	Users []InputUser    `tl:"users"`
}

func (*PhoneInviteToGroupCallRequestPredict) CRC() uint32 {
	return 0x7b393160
}

func PhoneInviteToGroupCall(ctx context.Context, m Requester, i PhoneInviteToGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneDiscardGroupCallRequestPredict struct {
	Call InputGroupCall `tl:"call"`
}

func (*PhoneDiscardGroupCallRequestPredict) CRC() uint32 {
	return 0x7a777135
}

func PhoneDiscardGroupCall(ctx context.Context, m Requester, i PhoneDiscardGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallSettingsRequestPredict struct {
	_               struct{}       `tl:"flags,bitflag"`
	ResetInviteHash bool           `tl:"reset_invite_hash,omitempty:flags:1,implicit"`
	Call            InputGroupCall `tl:"call"`
	JoinMuted       *bool          `tl:"join_muted,omitempty:flags:0"`
}

func (*PhoneToggleGroupCallSettingsRequestPredict) CRC() uint32 {
	return 0x74bbb43d
}

func PhoneToggleGroupCallSettings(ctx context.Context, m Requester, i PhoneToggleGroupCallSettingsRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallRequestPredict struct {
	Call  InputGroupCall `tl:"call"`
	Limit int32          `tl:"limit"`
}

func (*PhoneGetGroupCallRequestPredict) CRC() uint32 {
	return 0x041845db
}

func PhoneGetGroupCall(ctx context.Context, m Requester, i PhoneGetGroupCallRequestPredict) (PhoneGroupCall, error) {
	var res PhoneGroupCall
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupParticipantsRequestPredict struct {
	Call    InputGroupCall `tl:"call"`
	Ids     []InputPeer    `tl:"ids"`
	Sources []int32        `tl:"sources"`
	Offset  string         `tl:"offset"`
	Limit   int32          `tl:"limit"`
}

func (*PhoneGetGroupParticipantsRequestPredict) CRC() uint32 {
	return 0xc558d8ab
}

func PhoneGetGroupParticipants(ctx context.Context, m Requester, i PhoneGetGroupParticipantsRequestPredict) (PhoneGroupParticipants, error) {
	var res PhoneGroupParticipants
	return res, request(ctx, m, &i, &res)
}

type PhoneCheckGroupCallRequestPredict struct {
	Call    InputGroupCall `tl:"call"`
	Sources []int32        `tl:"sources"`
}

func (*PhoneCheckGroupCallRequestPredict) CRC() uint32 {
	return 0xb59cf977
}

func PhoneCheckGroupCall(ctx context.Context, m Requester, i PhoneCheckGroupCallRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallRecordRequestPredict struct {
	_             struct{}       `tl:"flags,bitflag"`
	Start         bool           `tl:"start,omitempty:flags:0,implicit"`
	Video         bool           `tl:"video,omitempty:flags:2,implicit"`
	Call          InputGroupCall `tl:"call"`
	Title         *string        `tl:"title,omitempty:flags:1"`
	VideoPortrait *bool          `tl:"video_portrait,omitempty:flags:2"`
}

func (*PhoneToggleGroupCallRecordRequestPredict) CRC() uint32 {
	return 0xf128c708
}

func PhoneToggleGroupCallRecord(ctx context.Context, m Requester, i PhoneToggleGroupCallRecordRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneEditGroupCallParticipantRequestPredict struct {
	_                  struct{}       `tl:"flags,bitflag"`
	Call               InputGroupCall `tl:"call"`
	Participant        InputPeer      `tl:"participant"`
	Muted              *bool          `tl:"muted,omitempty:flags:0"`
	Volume             *int32         `tl:"volume,omitempty:flags:1"`
	RaiseHand          *bool          `tl:"raise_hand,omitempty:flags:2"`
	VideoStopped       *bool          `tl:"video_stopped,omitempty:flags:3"`
	VideoPaused        *bool          `tl:"video_paused,omitempty:flags:4"`
	PresentationPaused *bool          `tl:"presentation_paused,omitempty:flags:5"`
}

func (*PhoneEditGroupCallParticipantRequestPredict) CRC() uint32 {
	return 0xa5273abf
}

func PhoneEditGroupCallParticipant(ctx context.Context, m Requester, i PhoneEditGroupCallParticipantRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneEditGroupCallTitleRequestPredict struct {
	Call  InputGroupCall `tl:"call"`
	Title string         `tl:"title"`
}

func (*PhoneEditGroupCallTitleRequestPredict) CRC() uint32 {
	return 0x1ca6ac0a
}

func PhoneEditGroupCallTitle(ctx context.Context, m Requester, i PhoneEditGroupCallTitleRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallJoinAsRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*PhoneGetGroupCallJoinAsRequestPredict) CRC() uint32 {
	return 0xef7c213a
}

func PhoneGetGroupCallJoinAs(ctx context.Context, m Requester, i PhoneGetGroupCallJoinAsRequestPredict) (PhoneJoinAsPeers, error) {
	var res PhoneJoinAsPeers
	return res, request(ctx, m, &i, &res)
}

type PhoneExportGroupCallInviteRequestPredict struct {
	_             struct{}       `tl:"flags,bitflag"`
	CanSelfUnmute bool           `tl:"can_self_unmute,omitempty:flags:0,implicit"`
	Call          InputGroupCall `tl:"call"`
}

func (*PhoneExportGroupCallInviteRequestPredict) CRC() uint32 {
	return 0xe6aa647f
}

func PhoneExportGroupCallInvite(ctx context.Context, m Requester, i PhoneExportGroupCallInviteRequestPredict) (PhoneExportedGroupCallInvite, error) {
	var res PhoneExportedGroupCallInvite
	return res, request(ctx, m, &i, &res)
}

type PhoneToggleGroupCallStartSubscriptionRequestPredict struct {
	Call       InputGroupCall `tl:"call"`
	Subscribed bool           `tl:"subscribed"`
}

func (*PhoneToggleGroupCallStartSubscriptionRequestPredict) CRC() uint32 {
	return 0x219c34e6
}

func PhoneToggleGroupCallStartSubscription(ctx context.Context, m Requester, i PhoneToggleGroupCallStartSubscriptionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneStartScheduledGroupCallRequestPredict struct {
	Call InputGroupCall `tl:"call"`
}

func (*PhoneStartScheduledGroupCallRequestPredict) CRC() uint32 {
	return 0x5680e342
}

func PhoneStartScheduledGroupCall(ctx context.Context, m Requester, i PhoneStartScheduledGroupCallRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveDefaultGroupCallJoinAsRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	JoinAs InputPeer `tl:"join_as"`
}

func (*PhoneSaveDefaultGroupCallJoinAsRequestPredict) CRC() uint32 {
	return 0x575e1f8c
}

func PhoneSaveDefaultGroupCallJoinAs(ctx context.Context, m Requester, i PhoneSaveDefaultGroupCallJoinAsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhoneJoinGroupCallPresentationRequestPredict struct {
	Call   InputGroupCall `tl:"call"`
	Params DataJSON       `tl:"params"`
}

func (*PhoneJoinGroupCallPresentationRequestPredict) CRC() uint32 {
	return 0xcbea6bc4
}

func PhoneJoinGroupCallPresentation(ctx context.Context, m Requester, i PhoneJoinGroupCallPresentationRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneLeaveGroupCallPresentationRequestPredict struct {
	Call InputGroupCall `tl:"call"`
}

func (*PhoneLeaveGroupCallPresentationRequestPredict) CRC() uint32 {
	return 0x1c50d144
}

func PhoneLeaveGroupCallPresentation(ctx context.Context, m Requester, i PhoneLeaveGroupCallPresentationRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallStreamChannelsRequestPredict struct {
	Call InputGroupCall `tl:"call"`
}

func (*PhoneGetGroupCallStreamChannelsRequestPredict) CRC() uint32 {
	return 0x1ab21940
}

func PhoneGetGroupCallStreamChannels(ctx context.Context, m Requester, i PhoneGetGroupCallStreamChannelsRequestPredict) (PhoneGroupCallStreamChannels, error) {
	var res PhoneGroupCallStreamChannels
	return res, request(ctx, m, &i, &res)
}

type PhoneGetGroupCallStreamRtmpURLRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	Revoke bool      `tl:"revoke"`
}

func (*PhoneGetGroupCallStreamRtmpURLRequestPredict) CRC() uint32 {
	return 0xdeb3abbf
}

func PhoneGetGroupCallStreamRtmpURL(ctx context.Context, m Requester, i PhoneGetGroupCallStreamRtmpURLRequestPredict) (PhoneGroupCallStreamRtmpURL, error) {
	var res PhoneGroupCallStreamRtmpURL
	return res, request(ctx, m, &i, &res)
}

type PhoneSaveCallLogRequestPredict struct {
	Peer InputPhoneCall `tl:"peer"`
	File InputFile      `tl:"file"`
}

func (*PhoneSaveCallLogRequestPredict) CRC() uint32 {
	return 0x41248786
}

func PhoneSaveCallLog(ctx context.Context, m Requester, i PhoneSaveCallLogRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type PhotosUpdateProfilePhotoRequestPredict struct {
	_        struct{}   `tl:"flags,bitflag"`
	Fallback bool       `tl:"fallback,omitempty:flags:0,implicit"`
	Bot      InputUser  `tl:"bot,omitempty:flags:1"`
	ID       InputPhoto `tl:"id"`
}

func (*PhotosUpdateProfilePhotoRequestPredict) CRC() uint32 {
	return 0x09e82039
}

func PhotosUpdateProfilePhoto(ctx context.Context, m Requester, i PhotosUpdateProfilePhotoRequestPredict) (PhotosPhoto, error) {
	var res PhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PhotosUploadProfilePhotoRequestPredict struct {
	_                struct{}  `tl:"flags,bitflag"`
	Fallback         bool      `tl:"fallback,omitempty:flags:3,implicit"`
	Bot              InputUser `tl:"bot,omitempty:flags:5"`
	File             InputFile `tl:"file,omitempty:flags:0"`
	Video            InputFile `tl:"video,omitempty:flags:1"`
	VideoStartTs     *float64  `tl:"video_start_ts,omitempty:flags:2"`
	VideoEmojiMarkup VideoSize `tl:"video_emoji_markup,omitempty:flags:4"`
}

func (*PhotosUploadProfilePhotoRequestPredict) CRC() uint32 {
	return 0x0388a3b5
}

func PhotosUploadProfilePhoto(ctx context.Context, m Requester, i PhotosUploadProfilePhotoRequestPredict) (PhotosPhoto, error) {
	var res PhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PhotosDeletePhotosRequestPredict struct {
	ID []InputPhoto `tl:"id"`
}

func (*PhotosDeletePhotosRequestPredict) CRC() uint32 {
	return 0x87cf7f2f
}

func PhotosDeletePhotos(ctx context.Context, m Requester, i PhotosDeletePhotosRequestPredict) ([]int64, error) {
	var res []int64
	return res, request(ctx, m, &i, &res)
}

type PhotosGetUserPhotosRequestPredict struct {
	UserID InputUser `tl:"user_id"`
	Offset int32     `tl:"offset"`
	MaxID  int64     `tl:"max_id"`
	Limit  int32     `tl:"limit"`
}

func (*PhotosGetUserPhotosRequestPredict) CRC() uint32 {
	return 0x91cd32a8
}

func PhotosGetUserPhotos(ctx context.Context, m Requester, i PhotosGetUserPhotosRequestPredict) (PhotosPhotos, error) {
	var res PhotosPhotos
	return res, request(ctx, m, &i, &res)
}

type PhotosUploadContactProfilePhotoRequestPredict struct {
	_                struct{}  `tl:"flags,bitflag"`
	Suggest          bool      `tl:"suggest,omitempty:flags:3,implicit"`
	Save             bool      `tl:"save,omitempty:flags:4,implicit"`
	UserID           InputUser `tl:"user_id"`
	File             InputFile `tl:"file,omitempty:flags:0"`
	Video            InputFile `tl:"video,omitempty:flags:1"`
	VideoStartTs     *float64  `tl:"video_start_ts,omitempty:flags:2"`
	VideoEmojiMarkup VideoSize `tl:"video_emoji_markup,omitempty:flags:5"`
}

func (*PhotosUploadContactProfilePhotoRequestPredict) CRC() uint32 {
	return 0xe14c4a71
}

func PhotosUploadContactProfilePhoto(ctx context.Context, m Requester, i PhotosUploadContactProfilePhotoRequestPredict) (PhotosPhoto, error) {
	var res PhotosPhoto
	return res, request(ctx, m, &i, &res)
}

type PremiumGetBoostsListRequestPredict struct {
	_      struct{}  `tl:"flags,bitflag"`
	Gifts  bool      `tl:"gifts,omitempty:flags:0,implicit"`
	Peer   InputPeer `tl:"peer"`
	Offset string    `tl:"offset"`
	Limit  int32     `tl:"limit"`
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
	return 0x0be77b4a
}

func PremiumGetMyBoosts(ctx context.Context, m Requester, i PremiumGetMyBoostsRequestPredict) (PremiumMyBoosts, error) {
	var res PremiumMyBoosts
	return res, request(ctx, m, &i, &res)
}

type PremiumApplyBoostRequestPredict struct {
	_     struct{}  `tl:"flags,bitflag"`
	Slots []int32   `tl:"slots,omitempty:flags:0"`
	Peer  InputPeer `tl:"peer"`
}

func (*PremiumApplyBoostRequestPredict) CRC() uint32 {
	return 0x6b7da746
}

func PremiumApplyBoost(ctx context.Context, m Requester, i PremiumApplyBoostRequestPredict) (PremiumMyBoosts, error) {
	var res PremiumMyBoosts
	return res, request(ctx, m, &i, &res)
}

type PremiumGetBoostsStatusRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*PremiumGetBoostsStatusRequestPredict) CRC() uint32 {
	return 0x042f1f61
}

func PremiumGetBoostsStatus(ctx context.Context, m Requester, i PremiumGetBoostsStatusRequestPredict) (PremiumBoostsStatus, error) {
	var res PremiumBoostsStatus
	return res, request(ctx, m, &i, &res)
}

type PremiumGetUserBoostsRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	UserID InputUser `tl:"user_id"`
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
	return 0x0edc39d0
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
	AllowInternational bool     `tl:"allow_international,omitempty:flags:0,implicit"`
}

func (*SmsjobsUpdateSettingsRequestPredict) CRC() uint32 {
	return 0x093fa0bf
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
	JobID string `tl:"job_id"`
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
	JobID string   `tl:"job_id"`
	Error *string  `tl:"error,omitempty:flags:0"`
}

func (*SmsjobsFinishJobRequestPredict) CRC() uint32 {
	return 0x4f1ebf24
}

func SmsjobsFinishJob(ctx context.Context, m Requester, i SmsjobsFinishJobRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastStatsRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Dark    bool         `tl:"dark,omitempty:flags:0,implicit"`
	Channel InputChannel `tl:"channel"`
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
	Token string   `tl:"token"`
	X     *int64   `tl:"x,omitempty:flags:0"`
}

func (*StatsLoadAsyncGraphRequestPredict) CRC() uint32 {
	return 0x621d5fa0
}

func StatsLoadAsyncGraph(ctx context.Context, m Requester, i StatsLoadAsyncGraphRequestPredict) (StatsGraph, error) {
	var res StatsGraph
	return res, request(ctx, m, &i, &res)
}

type StatsGetMegagroupStatsRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Dark    bool         `tl:"dark,omitempty:flags:0,implicit"`
	Channel InputChannel `tl:"channel"`
}

func (*StatsGetMegagroupStatsRequestPredict) CRC() uint32 {
	return 0xdcdf8607
}

func StatsGetMegagroupStats(ctx context.Context, m Requester, i StatsGetMegagroupStatsRequestPredict) (StatsMegagroupStats, error) {
	var res StatsMegagroupStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetMessagePublicForwardsRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	MsgID   int32        `tl:"msg_id"`
	Offset  string       `tl:"offset"`
	Limit   int32        `tl:"limit"`
}

func (*StatsGetMessagePublicForwardsRequestPredict) CRC() uint32 {
	return 0x5f150144
}

func StatsGetMessagePublicForwards(ctx context.Context, m Requester, i StatsGetMessagePublicForwardsRequestPredict) (StatsPublicForwards, error) {
	var res StatsPublicForwards
	return res, request(ctx, m, &i, &res)
}

type StatsGetMessageStatsRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Dark    bool         `tl:"dark,omitempty:flags:0,implicit"`
	Channel InputChannel `tl:"channel"`
	MsgID   int32        `tl:"msg_id"`
}

func (*StatsGetMessageStatsRequestPredict) CRC() uint32 {
	return 0xb6e0a3f5
}

func StatsGetMessageStats(ctx context.Context, m Requester, i StatsGetMessageStatsRequestPredict) (StatsMessageStats, error) {
	var res StatsMessageStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetStoryStatsRequestPredict struct {
	_    struct{}  `tl:"flags,bitflag"`
	Dark bool      `tl:"dark,omitempty:flags:0,implicit"`
	Peer InputPeer `tl:"peer"`
	ID   int32     `tl:"id"`
}

func (*StatsGetStoryStatsRequestPredict) CRC() uint32 {
	return 0x374fef40
}

func StatsGetStoryStats(ctx context.Context, m Requester, i StatsGetStoryStatsRequestPredict) (StatsStoryStats, error) {
	var res StatsStoryStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetStoryPublicForwardsRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	ID     int32     `tl:"id"`
	Offset string    `tl:"offset"`
	Limit  int32     `tl:"limit"`
}

func (*StatsGetStoryPublicForwardsRequestPredict) CRC() uint32 {
	return 0xa6437ef6
}

func StatsGetStoryPublicForwards(ctx context.Context, m Requester, i StatsGetStoryPublicForwardsRequestPredict) (StatsPublicForwards, error) {
	var res StatsPublicForwards
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastRevenueStatsRequestPredict struct {
	_       struct{}     `tl:"flags,bitflag"`
	Dark    bool         `tl:"dark,omitempty:flags:0,implicit"`
	Channel InputChannel `tl:"channel"`
}

func (*StatsGetBroadcastRevenueStatsRequestPredict) CRC() uint32 {
	return 0x75dfb671
}

func StatsGetBroadcastRevenueStats(ctx context.Context, m Requester, i StatsGetBroadcastRevenueStatsRequestPredict) (StatsBroadcastRevenueStats, error) {
	var res StatsBroadcastRevenueStats
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastRevenueWithdrawalURLRequestPredict struct {
	Channel  InputChannel          `tl:"channel"`
	Password InputCheckPasswordSRP `tl:"password"`
}

func (*StatsGetBroadcastRevenueWithdrawalURLRequestPredict) CRC() uint32 {
	return 0x2a65ef73
}

func StatsGetBroadcastRevenueWithdrawalURL(ctx context.Context, m Requester, i StatsGetBroadcastRevenueWithdrawalURLRequestPredict) (StatsBroadcastRevenueWithdrawalURL, error) {
	var res StatsBroadcastRevenueWithdrawalURL
	return res, request(ctx, m, &i, &res)
}

type StatsGetBroadcastRevenueTransactionsRequestPredict struct {
	Channel InputChannel `tl:"channel"`
	Offset  int32        `tl:"offset"`
	Limit   int32        `tl:"limit"`
}

func (*StatsGetBroadcastRevenueTransactionsRequestPredict) CRC() uint32 {
	return 0x0069280f
}

func StatsGetBroadcastRevenueTransactions(ctx context.Context, m Requester, i StatsGetBroadcastRevenueTransactionsRequestPredict) (StatsBroadcastRevenueTransactions, error) {
	var res StatsBroadcastRevenueTransactions
	return res, request(ctx, m, &i, &res)
}

type StickersCreateStickerSetRequestPredict struct {
	_         struct{}              `tl:"flags,bitflag"`
	Masks     bool                  `tl:"masks,omitempty:flags:0,implicit"`
	Emojis    bool                  `tl:"emojis,omitempty:flags:5,implicit"`
	TextColor bool                  `tl:"text_color,omitempty:flags:6,implicit"`
	UserID    InputUser             `tl:"user_id"`
	Title     string                `tl:"title"`
	ShortName string                `tl:"short_name"`
	Thumb     InputDocument         `tl:"thumb,omitempty:flags:2"`
	Stickers  []InputStickerSetItem `tl:"stickers"`
	Software  *string               `tl:"software,omitempty:flags:3"`
}

func (*StickersCreateStickerSetRequestPredict) CRC() uint32 {
	return 0x9021ab67
}

func StickersCreateStickerSet(ctx context.Context, m Requester, i StickersCreateStickerSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersRemoveStickerFromSetRequestPredict struct {
	Sticker InputDocument `tl:"sticker"`
}

func (*StickersRemoveStickerFromSetRequestPredict) CRC() uint32 {
	return 0xf7760f51
}

func StickersRemoveStickerFromSet(ctx context.Context, m Requester, i StickersRemoveStickerFromSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersChangeStickerPositionRequestPredict struct {
	Sticker  InputDocument `tl:"sticker"`
	Position int32         `tl:"position"`
}

func (*StickersChangeStickerPositionRequestPredict) CRC() uint32 {
	return 0xffb6d4ca
}

func StickersChangeStickerPosition(ctx context.Context, m Requester, i StickersChangeStickerPositionRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersAddStickerToSetRequestPredict struct {
	Stickerset InputStickerSet     `tl:"stickerset"`
	Sticker    InputStickerSetItem `tl:"sticker"`
}

func (*StickersAddStickerToSetRequestPredict) CRC() uint32 {
	return 0x8653febe
}

func StickersAddStickerToSet(ctx context.Context, m Requester, i StickersAddStickerToSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersSetStickerSetThumbRequestPredict struct {
	_               struct{}        `tl:"flags,bitflag"`
	Stickerset      InputStickerSet `tl:"stickerset"`
	Thumb           InputDocument   `tl:"thumb,omitempty:flags:0"`
	ThumbDocumentID *int64          `tl:"thumb_document_id,omitempty:flags:1"`
}

func (*StickersSetStickerSetThumbRequestPredict) CRC() uint32 {
	return 0xa76a5392
}

func StickersSetStickerSetThumb(ctx context.Context, m Requester, i StickersSetStickerSetThumbRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersCheckShortNameRequestPredict struct {
	ShortName string `tl:"short_name"`
}

func (*StickersCheckShortNameRequestPredict) CRC() uint32 {
	return 0x284b3639
}

func StickersCheckShortName(ctx context.Context, m Requester, i StickersCheckShortNameRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StickersSuggestShortNameRequestPredict struct {
	Title string `tl:"title"`
}

func (*StickersSuggestShortNameRequestPredict) CRC() uint32 {
	return 0x4dafc503
}

func StickersSuggestShortName(ctx context.Context, m Requester, i StickersSuggestShortNameRequestPredict) (StickersSuggestedShortName, error) {
	var res StickersSuggestedShortName
	return res, request(ctx, m, &i, &res)
}

type StickersChangeStickerRequestPredict struct {
	_          struct{}      `tl:"flags,bitflag"`
	Sticker    InputDocument `tl:"sticker"`
	Emoji      *string       `tl:"emoji,omitempty:flags:0"`
	MaskCoords MaskCoords    `tl:"mask_coords,omitempty:flags:1"`
	Keywords   *string       `tl:"keywords,omitempty:flags:2"`
}

func (*StickersChangeStickerRequestPredict) CRC() uint32 {
	return 0xf5537ebc
}

func StickersChangeSticker(ctx context.Context, m Requester, i StickersChangeStickerRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersRenameStickerSetRequestPredict struct {
	Stickerset InputStickerSet `tl:"stickerset"`
	Title      string          `tl:"title"`
}

func (*StickersRenameStickerSetRequestPredict) CRC() uint32 {
	return 0x124b1c00
}

func StickersRenameStickerSet(ctx context.Context, m Requester, i StickersRenameStickerSetRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StickersDeleteStickerSetRequestPredict struct {
	Stickerset InputStickerSet `tl:"stickerset"`
}

func (*StickersDeleteStickerSetRequestPredict) CRC() uint32 {
	return 0x87704394
}

func StickersDeleteStickerSet(ctx context.Context, m Requester, i StickersDeleteStickerSetRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StickersReplaceStickerRequestPredict struct {
	Sticker    InputDocument       `tl:"sticker"`
	NewSticker InputStickerSetItem `tl:"new_sticker"`
}

func (*StickersReplaceStickerRequestPredict) CRC() uint32 {
	return 0x4696459a
}

func StickersReplaceSticker(ctx context.Context, m Requester, i StickersReplaceStickerRequestPredict) (MessagesStickerSet, error) {
	var res MessagesStickerSet
	return res, request(ctx, m, &i, &res)
}

type StoriesCanSendStoryRequestPredict struct {
	Peer InputPeer `tl:"peer"`
}

func (*StoriesCanSendStoryRequestPredict) CRC() uint32 {
	return 0xc7dfdfdd
}

func StoriesCanSendStory(ctx context.Context, m Requester, i StoriesCanSendStoryRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesSendStoryRequestPredict struct {
	_            struct{}           `tl:"flags,bitflag"`
	Pinned       bool               `tl:"pinned,omitempty:flags:2,implicit"`
	Noforwards   bool               `tl:"noforwards,omitempty:flags:4,implicit"`
	FwdModified  bool               `tl:"fwd_modified,omitempty:flags:7,implicit"`
	Peer         InputPeer          `tl:"peer"`
	Media        InputMedia         `tl:"media"`
	MediaAreas   []MediaArea        `tl:"media_areas,omitempty:flags:5"`
	Caption      *string            `tl:"caption,omitempty:flags:0"`
	Entities     []MessageEntity    `tl:"entities,omitempty:flags:1"`
	PrivacyRules []InputPrivacyRule `tl:"privacy_rules"`
	RandomID     int64              `tl:"random_id"`
	Period       *int32             `tl:"period,omitempty:flags:3"`
	FwdFromID    InputPeer          `tl:"fwd_from_id,omitempty:flags:6"`
	FwdFromStory *int32             `tl:"fwd_from_story,omitempty:flags:6"`
}

func (*StoriesSendStoryRequestPredict) CRC() uint32 {
	return 0xe4e6694b
}

func StoriesSendStory(ctx context.Context, m Requester, i StoriesSendStoryRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesEditStoryRequestPredict struct {
	_            struct{}           `tl:"flags,bitflag"`
	Peer         InputPeer          `tl:"peer"`
	ID           int32              `tl:"id"`
	Media        InputMedia         `tl:"media,omitempty:flags:0"`
	MediaAreas   []MediaArea        `tl:"media_areas,omitempty:flags:3"`
	Caption      *string            `tl:"caption,omitempty:flags:1"`
	Entities     []MessageEntity    `tl:"entities,omitempty:flags:1"`
	PrivacyRules []InputPrivacyRule `tl:"privacy_rules,omitempty:flags:2"`
}

func (*StoriesEditStoryRequestPredict) CRC() uint32 {
	return 0xb583ba46
}

func StoriesEditStory(ctx context.Context, m Requester, i StoriesEditStoryRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesDeleteStoriesRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*StoriesDeleteStoriesRequestPredict) CRC() uint32 {
	return 0xae59db5f
}

func StoriesDeleteStories(ctx context.Context, m Requester, i StoriesDeleteStoriesRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesTogglePinnedRequestPredict struct {
	Peer   InputPeer `tl:"peer"`
	ID     []int32   `tl:"id"`
	Pinned bool      `tl:"pinned"`
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
	Next   bool     `tl:"next,omitempty:flags:1,implicit"`
	Hidden bool     `tl:"hidden,omitempty:flags:2,implicit"`
	State  *string  `tl:"state,omitempty:flags:0"`
}

func (*StoriesGetAllStoriesRequestPredict) CRC() uint32 {
	return 0xeeb0d625
}

func StoriesGetAllStories(ctx context.Context, m Requester, i StoriesGetAllStoriesRequestPredict) (StoriesAllStories, error) {
	var res StoriesAllStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPinnedStoriesRequestPredict struct {
	Peer     InputPeer `tl:"peer"`
	OffsetID int32     `tl:"offset_id"`
	Limit    int32     `tl:"limit"`
}

func (*StoriesGetPinnedStoriesRequestPredict) CRC() uint32 {
	return 0x5821a5dc
}

func StoriesGetPinnedStories(ctx context.Context, m Requester, i StoriesGetPinnedStoriesRequestPredict) (StoriesStories, error) {
	var res StoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesArchiveRequestPredict struct {
	Peer     InputPeer `tl:"peer"`
	OffsetID int32     `tl:"offset_id"`
	Limit    int32     `tl:"limit"`
}

func (*StoriesGetStoriesArchiveRequestPredict) CRC() uint32 {
	return 0xb4352016
}

func StoriesGetStoriesArchive(ctx context.Context, m Requester, i StoriesGetStoriesArchiveRequestPredict) (StoriesStories, error) {
	var res StoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesByIDRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*StoriesGetStoriesByIDRequestPredict) CRC() uint32 {
	return 0x5774ca74
}

func StoriesGetStoriesByID(ctx context.Context, m Requester, i StoriesGetStoriesByIDRequestPredict) (StoriesStories, error) {
	var res StoriesStories
	return res, request(ctx, m, &i, &res)
}

type StoriesToggleAllStoriesHiddenRequestPredict struct {
	Hidden bool `tl:"hidden"`
}

func (*StoriesToggleAllStoriesHiddenRequestPredict) CRC() uint32 {
	return 0x7c2557c4
}

func StoriesToggleAllStoriesHidden(ctx context.Context, m Requester, i StoriesToggleAllStoriesHiddenRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesReadStoriesRequestPredict struct {
	Peer  InputPeer `tl:"peer"`
	MaxID int32     `tl:"max_id"`
}

func (*StoriesReadStoriesRequestPredict) CRC() uint32 {
	return 0xa556dac8
}

func StoriesReadStories(ctx context.Context, m Requester, i StoriesReadStoriesRequestPredict) ([]int32, error) {
	var res []int32
	return res, request(ctx, m, &i, &res)
}

type StoriesIncrementStoryViewsRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*StoriesIncrementStoryViewsRequestPredict) CRC() uint32 {
	return 0xb2028afb
}

func StoriesIncrementStoryViews(ctx context.Context, m Requester, i StoriesIncrementStoryViewsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoryViewsListRequestPredict struct {
	_              struct{}  `tl:"flags,bitflag"`
	JustContacts   bool      `tl:"just_contacts,omitempty:flags:0,implicit"`
	ReactionsFirst bool      `tl:"reactions_first,omitempty:flags:2,implicit"`
	ForwardsFirst  bool      `tl:"forwards_first,omitempty:flags:3,implicit"`
	Peer           InputPeer `tl:"peer"`
	Q              *string   `tl:"q,omitempty:flags:1"`
	ID             int32     `tl:"id"`
	Offset         string    `tl:"offset"`
	Limit          int32     `tl:"limit"`
}

func (*StoriesGetStoryViewsListRequestPredict) CRC() uint32 {
	return 0x7ed23c57
}

func StoriesGetStoryViewsList(ctx context.Context, m Requester, i StoriesGetStoryViewsListRequestPredict) (StoriesStoryViewsList, error) {
	var res StoriesStoryViewsList
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoriesViewsRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*StoriesGetStoriesViewsRequestPredict) CRC() uint32 {
	return 0x28e16cc8
}

func StoriesGetStoriesViews(ctx context.Context, m Requester, i StoriesGetStoriesViewsRequestPredict) (StoriesStoryViews, error) {
	var res StoriesStoryViews
	return res, request(ctx, m, &i, &res)
}

type StoriesExportStoryLinkRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   int32     `tl:"id"`
}

func (*StoriesExportStoryLinkRequestPredict) CRC() uint32 {
	return 0x7b8def20
}

func StoriesExportStoryLink(ctx context.Context, m Requester, i StoriesExportStoryLinkRequestPredict) (ExportedStoryLink, error) {
	var res ExportedStoryLink
	return res, request(ctx, m, &i, &res)
}

type StoriesReportRequestPredict struct {
	Peer    InputPeer    `tl:"peer"`
	ID      []int32      `tl:"id"`
	Reason  ReportReason `tl:"reason"`
	Message string       `tl:"message"`
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
	Past   bool     `tl:"past,omitempty:flags:0,implicit"`
	Future bool     `tl:"future,omitempty:flags:1,implicit"`
}

func (*StoriesActivateStealthModeRequestPredict) CRC() uint32 {
	return 0x57bbd166
}

func StoriesActivateStealthMode(ctx context.Context, m Requester, i StoriesActivateStealthModeRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesSendReactionRequestPredict struct {
	_           struct{}  `tl:"flags,bitflag"`
	AddToRecent bool      `tl:"add_to_recent,omitempty:flags:0,implicit"`
	Peer        InputPeer `tl:"peer"`
	StoryID     int32     `tl:"story_id"`
	Reaction    Reaction  `tl:"reaction"`
}

func (*StoriesSendReactionRequestPredict) CRC() uint32 {
	return 0x7fd736b2
}

func StoriesSendReaction(ctx context.Context, m Requester, i StoriesSendReactionRequestPredict) (Updates, error) {
	var res Updates
	return res, request(ctx, m, &i, &res)
}

type StoriesGetPeerStoriesRequestPredict struct {
	Peer InputPeer `tl:"peer"`
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
	ID []InputPeer `tl:"id"`
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
	Peer   InputPeer `tl:"peer"`
	Hidden bool      `tl:"hidden"`
}

func (*StoriesTogglePeerStoriesHiddenRequestPredict) CRC() uint32 {
	return 0xbd0415c4
}

func StoriesTogglePeerStoriesHidden(ctx context.Context, m Requester, i StoriesTogglePeerStoriesHiddenRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesGetStoryReactionsListRequestPredict struct {
	_             struct{}  `tl:"flags,bitflag"`
	ForwardsFirst bool      `tl:"forwards_first,omitempty:flags:2,implicit"`
	Peer          InputPeer `tl:"peer"`
	ID            int32     `tl:"id"`
	Reaction      Reaction  `tl:"reaction,omitempty:flags:0"`
	Offset        *string   `tl:"offset,omitempty:flags:1"`
	Limit         int32     `tl:"limit"`
}

func (*StoriesGetStoryReactionsListRequestPredict) CRC() uint32 {
	return 0xb9b2881f
}

func StoriesGetStoryReactionsList(ctx context.Context, m Requester, i StoriesGetStoryReactionsListRequestPredict) (StoriesStoryReactionsList, error) {
	var res StoriesStoryReactionsList
	return res, request(ctx, m, &i, &res)
}

type StoriesTogglePinnedToTopRequestPredict struct {
	Peer InputPeer `tl:"peer"`
	ID   []int32   `tl:"id"`
}

func (*StoriesTogglePinnedToTopRequestPredict) CRC() uint32 {
	return 0x0b297e9b
}

func StoriesTogglePinnedToTop(ctx context.Context, m Requester, i StoriesTogglePinnedToTopRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type StoriesSearchPostsRequestPredict struct {
	_       struct{}  `tl:"flags,bitflag"`
	Hashtag *string   `tl:"hashtag,omitempty:flags:0"`
	Area    MediaArea `tl:"area,omitempty:flags:1"`
	Offset  string    `tl:"offset"`
	Limit   int32     `tl:"limit"`
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
	Pts           int32    `tl:"pts"`
	PtsLimit      *int32   `tl:"pts_limit,omitempty:flags:1"`
	PtsTotalLimit *int32   `tl:"pts_total_limit,omitempty:flags:0"`
	Date          int32    `tl:"date"`
	Qts           int32    `tl:"qts"`
	QtsLimit      *int32   `tl:"qts_limit,omitempty:flags:2"`
}

func (*UpdatesGetDifferenceRequestPredict) CRC() uint32 {
	return 0x19c2f763
}

func UpdatesGetDifference(ctx context.Context, m Requester, i UpdatesGetDifferenceRequestPredict) (UpdatesDifference, error) {
	var res UpdatesDifference
	return res, request(ctx, m, &i, &res)
}

type UpdatesGetChannelDifferenceRequestPredict struct {
	_       struct{}              `tl:"flags,bitflag"`
	Force   bool                  `tl:"force,omitempty:flags:0,implicit"`
	Channel InputChannel          `tl:"channel"`
	Filter  ChannelMessagesFilter `tl:"filter"`
	Pts     int32                 `tl:"pts"`
	Limit   int32                 `tl:"limit"`
}

func (*UpdatesGetChannelDifferenceRequestPredict) CRC() uint32 {
	return 0x03173d78
}

func UpdatesGetChannelDifference(ctx context.Context, m Requester, i UpdatesGetChannelDifferenceRequestPredict) (UpdatesChannelDifference, error) {
	var res UpdatesChannelDifference
	return res, request(ctx, m, &i, &res)
}

type UploadSaveFilePartRequestPredict struct {
	FileID   int64  `tl:"file_id"`
	FilePart int32  `tl:"file_part"`
	Bytes    []byte `tl:"bytes"`
}

func (*UploadSaveFilePartRequestPredict) CRC() uint32 {
	return 0xb304a621
}

func UploadSaveFilePart(ctx context.Context, m Requester, i UploadSaveFilePartRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UploadGetFileRequestPredict struct {
	_            struct{}          `tl:"flags,bitflag"`
	Precise      bool              `tl:"precise,omitempty:flags:0,implicit"`
	CdnSupported bool              `tl:"cdn_supported,omitempty:flags:1,implicit"`
	Location     InputFileLocation `tl:"location"`
	Offset       int64             `tl:"offset"`
	Limit        int32             `tl:"limit"`
}

func (*UploadGetFileRequestPredict) CRC() uint32 {
	return 0xbe5335be
}

func UploadGetFile(ctx context.Context, m Requester, i UploadGetFileRequestPredict) (UploadFile, error) {
	var res UploadFile
	return res, request(ctx, m, &i, &res)
}

type UploadSaveBigFilePartRequestPredict struct {
	FileID         int64  `tl:"file_id"`
	FilePart       int32  `tl:"file_part"`
	FileTotalParts int32  `tl:"file_total_parts"`
	Bytes          []byte `tl:"bytes"`
}

func (*UploadSaveBigFilePartRequestPredict) CRC() uint32 {
	return 0xde7b673d
}

func UploadSaveBigFilePart(ctx context.Context, m Requester, i UploadSaveBigFilePartRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UploadGetWebFileRequestPredict struct {
	Location InputWebFileLocation `tl:"location"`
	Offset   int32                `tl:"offset"`
	Limit    int32                `tl:"limit"`
}

func (*UploadGetWebFileRequestPredict) CRC() uint32 {
	return 0x24e6818d
}

func UploadGetWebFile(ctx context.Context, m Requester, i UploadGetWebFileRequestPredict) (UploadWebFile, error) {
	var res UploadWebFile
	return res, request(ctx, m, &i, &res)
}

type UploadGetCdnFileRequestPredict struct {
	FileToken []byte `tl:"file_token"`
	Offset    int64  `tl:"offset"`
	Limit     int32  `tl:"limit"`
}

func (*UploadGetCdnFileRequestPredict) CRC() uint32 {
	return 0x395f69da
}

func UploadGetCdnFile(ctx context.Context, m Requester, i UploadGetCdnFileRequestPredict) (UploadCdnFile, error) {
	var res UploadCdnFile
	return res, request(ctx, m, &i, &res)
}

type UploadReuploadCdnFileRequestPredict struct {
	FileToken    []byte `tl:"file_token"`
	RequestToken []byte `tl:"request_token"`
}

func (*UploadReuploadCdnFileRequestPredict) CRC() uint32 {
	return 0x9b2754a8
}

func UploadReuploadCdnFile(ctx context.Context, m Requester, i UploadReuploadCdnFileRequestPredict) ([]FileHash, error) {
	var res []FileHash
	return res, request(ctx, m, &i, &res)
}

type UploadGetCdnFileHashesRequestPredict struct {
	FileToken []byte `tl:"file_token"`
	Offset    int64  `tl:"offset"`
}

func (*UploadGetCdnFileHashesRequestPredict) CRC() uint32 {
	return 0x91dc3f31
}

func UploadGetCdnFileHashes(ctx context.Context, m Requester, i UploadGetCdnFileHashesRequestPredict) ([]FileHash, error) {
	var res []FileHash
	return res, request(ctx, m, &i, &res)
}

type UploadGetFileHashesRequestPredict struct {
	Location InputFileLocation `tl:"location"`
	Offset   int64             `tl:"offset"`
}

func (*UploadGetFileHashesRequestPredict) CRC() uint32 {
	return 0x9156982a
}

func UploadGetFileHashes(ctx context.Context, m Requester, i UploadGetFileHashesRequestPredict) ([]FileHash, error) {
	var res []FileHash
	return res, request(ctx, m, &i, &res)
}

type UsersGetUsersRequestPredict struct {
	ID []InputUser `tl:"id"`
}

func (*UsersGetUsersRequestPredict) CRC() uint32 {
	return 0x0d91a548
}

func UsersGetUsers(ctx context.Context, m Requester, i UsersGetUsersRequestPredict) ([]User, error) {
	var res []User
	return res, request(ctx, m, &i, &res)
}

type UsersGetFullUserRequestPredict struct {
	ID InputUser `tl:"id"`
}

func (*UsersGetFullUserRequestPredict) CRC() uint32 {
	return 0xb60f5918
}

func UsersGetFullUser(ctx context.Context, m Requester, i UsersGetFullUserRequestPredict) (UsersUserFull, error) {
	var res UsersUserFull
	return res, request(ctx, m, &i, &res)
}

type UsersSetSecureValueErrorsRequestPredict struct {
	ID     InputUser          `tl:"id"`
	Errors []SecureValueError `tl:"errors"`
}

func (*UsersSetSecureValueErrorsRequestPredict) CRC() uint32 {
	return 0x90c894b5
}

func UsersSetSecureValueErrors(ctx context.Context, m Requester, i UsersSetSecureValueErrorsRequestPredict) (bool, error) {
	var res bool
	return res, request(ctx, m, &i, &res)
}

type UsersGetIsPremiumRequiredToContactRequestPredict struct {
	ID []InputUser `tl:"id"`
}

func (*UsersGetIsPremiumRequiredToContactRequestPredict) CRC() uint32 {
	return 0xa622aa10
}

func UsersGetIsPremiumRequiredToContact(ctx context.Context, m Requester, i UsersGetIsPremiumRequiredToContactRequestPredict) ([]bool, error) {
	var res []bool
	return res, request(ctx, m, &i, &res)
}
