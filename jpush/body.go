package jpush

const (
	PLATFORM_ALL       = "all"
	PLATFORM_ANDROID   = "android"
	PLATFORM_IOS       = "ios"
	PLATFORM_HARMONYOS = "hmos"
	PLATFORM_QUICKAPP  = "quickapp" // VIP用户功能
)

type Body struct {
	// "all" || []string
	Platform any `json:"platform"`
	// "all" || *Audience
	Audience     any           `json:"audience"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	Options      *Option       `json:"options,omitempty"` // iOS不可省略
}

type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationId []string `json:"registration_id,omitempty"`
}

type Notification struct {
	Alert   string   `json:"alert,omitempty"`
	Android *Android `json:"android,omitempty"`
	Ios     *Ios     `json:"ios,omitempty"`
	Homs    *Homs    `json:"homs,omitempty"`
}

type Android struct {
	Alert       string         `json:"alert"`
	Title       string         `json:"title,omitempty"`
	ChannelId   int            `json:"channel_id,omitempty"`
	Extras      map[string]any `json:"extras,omitempty"`
	LargeIcon   string         `json:"large_icon,omitempty"`
	BadgeAddNum int            `json:"badge_add_num,omitempty"`
	BadgeSetNum int            `json:"badge_set_num,omitempty"`
	Sound       string         `json:"sound,omitempty"`
}

type Ios struct {
	// string || *IosAlert
	Alert            any            `json:"alert"`
	Sound            string         `json:"sound,omitempty"`
	Badge            string         `json:"badge,omitempty"`
	ContentAvailable bool           `json:"content-available,omitempty"`
	MutableContent   bool           `json:"mutable-content,omitempty"`
	Category         string         `json:"category,omitempty"`
	Extras           map[string]any `json:"extras,omitempty"`
}

type IosAlert struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

type Homs struct {
	Alert       string         `json:"alert"`
	Title       string         `json:"title,omitempty"`
	Category    string         `json:"category"`
	LargeIcon   string         `json:"large_icon,omitempty"`
	Intent      map[string]any `json:"intent"`
	BadgeAddNum int            `json:"badge_add_num,omitempty"`
	BadgeSetNum int            `json:"badge_set_num,omitempty"`
	Extras      map[string]any `json:"extras,omitempty"`
}

type Message struct {
	Content     string         `json:"msg_content"`
	Title       string         `json:"title,omitempty"`
	ContentType string         `json:"content_type,omitempty"`
	Extras      map[string]any `json:"extras,omitempty"`
}

type Option struct {
	SendNo          int   `json:"sendno,omitempty"`
	TimeToLive      int   `json:"time_to_live,omitempty"`
	OverrideMsgId   int64 `json:"override_msg_id,omitempty"`
	ApnsProduction  bool  `json:"apns_production"`
	BigPushDuration int   `json:"big_push_duration,omitempty"`
}
