package webserver

type discordOauth struct {
	RedirectURL  string `json:"redirect_url,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

type discordUser struct {
	ID               string      `json:"id,omitempty"`
	Username         string      `json:"username,omitempty"`
	Avatar           string      `json:"avatar,omitempty"`
	AvatarDecoration interface{} `json:"avatar_decoration,omitempty"`
	Discriminator    string      `json:"discriminator,omitempty"`
	PublicFlags      int         `json:"public_flags,omitempty"`
	Flags            int         `json:"flags,omitempty"`
	Banner           interface{} `json:"banner,omitempty"`
	BannerColor      interface{} `json:"banner_color,omitempty"`
	AccentColor      interface{} `json:"accent_color,omitempty"`
	Locale           string      `json:"locale,omitempty"`
	MfaEnabled       bool        `json:"mfa_enabled,omitempty"`
	PremiumType      int         `json:"premium_type,omitempty"`
}

type discordUserGuilds struct {
	ID             string   `json:"id,omitempty"`
	Name           string   `json:"name,omitempty"`
	Icon           string   `json:"icon,omitempty"`
	Owner          bool     `json:"owner,omitempty"`
	Permissions    int      `json:"permissions,omitempty"`
	Features       []string `json:"features,omitempty"`
	PermissionsNew string   `json:"permissions_new,omitempty"`
}

type discordUserRole struct {
	ID           string      `json:"id,omitempty"`
	Name         string      `json:"name,omitempty"`
	Color        int         `json:"color,omitempty"`
	Hoist        bool        `json:"hoist,omitempty"`
	Icon         string      `json:"icon,omitempty"`
	UnicodeEmoji interface{} `json:"unicode_emoji,omitempty"`
	Position     int         `json:"position,omitempty"`
	Permissions  string      `json:"permissions,omitempty"`
	Managed      bool        `json:"managed,omitempty"`
	Mentionable  bool        `json:"mentionable,omitempty"`
}
