package twitterstream

type User struct {
	ContributorsEnabled       *bool  `json:"contributors_enabled"`
	CreatedAt                 string `json:"created_at"`
	Description               *string
	FavoritesCount            uint32 `json:"favorites_count"`
	FollowersCount            uint32 `json:"followers_count"`
	FriendsCount              uint32 `json:"friends_count"`
	GeoEnabled                bool   `json:"geo_enabled"`
	Id                        uint64
	Lang                      *string
	Location                  *string
	Name                      *string
	Notifications             *bool
	ProfileBackgroundColor    *string `json:"profile_background_color"`
	ProfileBackgroundImageUrl *string `json:"profile_background_image_url;"`
	ProfileBackgroundTile     *bool   `json:"profile_background_tile"`
	ProfileImageUrl           *string `json:"profile_image_url"`
	ProfileLinkColor          *string `json:"profile_link_color"`
	ProfileSidebarBorderColor *string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor   *string `json:"profile_sidebar_fill_color"`
	ProfileTextColor          *string `json:"profile_text_color"`
	ProfileUseBackgroundImage bool    `json:"profile_use_background_image"`
	Protected                 bool
	Status                    *Tweet
	StatusesCount             uint32  `json:"statuses_count"`
	TimeZone                  *string `json:"time_zone"`
	Url                       *string
	UtcOffset                 int32
	Verified                  bool
}

type Url struct {
	DisplayUrl  *string `json:"display_url"`
	ExpandedUrl *string `json:"expanded_url"`
	Indices     []uint32
	Url         *string
}

type UserMention struct {
	Id         uint64
	IdStr      string `json:"id_str"`
	Indices    []uint32
	Name       string
	ScreenName string `json:"screen_name"`
}

type BoundingBox struct {
	Type        string
	Coordinates [][][]float32
}

type Entities struct {
	Hashtags     []Hashtag
	Media        []Media
	Urls         []Url
	UserMentions []UserMention `json:"user_mentions"`
}

type Geo struct {
	Type        string
	Coordinates []float32
}

type Hashtag struct {
	Indices []uint32
	Text    string
}

type Media struct {
	DisplayUrl    string
	ExpandedUrl   string
	Id            uint64
	IdStr         string `json:"id_str"`
	Indices       []uint32
	MediaUrl      string `json:"media_url"`
	MediaUrlHttps string `json:"media_url_https"`
	Sizes         map[string]MediaSize
	Type          string
	Url           string
}

type MediaSize struct {
	Height uint32 `json:"h"`
	Resize string
	Width  uint32 `json:"w"`
}

type Place struct {
	Attributres map[string]string
	BoundingBox *BoundingBox `json:"bounding_box"`
	Country     string
	CountryCode string `json:"country_code"`
	FullName    string `json:"full_name"`
	Id          string
	Name        string
	PlaceType   string `json:place_type"`
	Url         string
}

type Tweet struct {
	Annotations          *string
	CreatedAt            string `json:"created_at"`
	Contributors         *string
	Coordinates          *Geo
	Entities             *Entities
	Favorited            bool
	Geo                  *Geo
	Id                   uint64
	IdStr                *string `json:"id_str"`
	InReplyToScreenName  *string `json:"in_reply_to_screen_name"`
	InReplyToStatusId    *uint64 `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr *string `json:"in_reply_to_status_id_str"`
	InReplyToUserId      *uint64 `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   *string `json:"in_reply_to_user_id_str"`
	Place                *Place
	RetweetCount         uint32 `json:"retweet_count"`
	Retweeted            bool
	Source               string
	Text                 string
	Truncated            bool
	User                 *User

}

type SiteStreamMessage struct {
    For_user int64
    Message  Tweet
}

type Event struct {
    Target     User
    Source     User
    Created_at string
    Event      string
}

type FriendList struct {
    Friends []int64
}
