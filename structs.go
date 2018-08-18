package mastodon

// Account holds informations about an account.
type Account struct {
	ID          string `json:"id"`              // The ID of the account
	Username    string `json:"username"`        // The username of the account
	Acct        string `json:"acct"`            // Equals username for local users, includes @domain for remote ones
	DisplayName string `json:"display_name"`    // The account's display name
	Note        string `json:"note"`            // Biography of user
	URL         string `json:"url"`             // URL of the user's profile page (can be remote)
	Avatar      string `json:"avatar"`          // URL to the avatar image
	Header      string `json:"header"`          // URL to the header image
	Locked      bool   `json:"locked"`          // Boolean for when the account cannot be followed without waiting for approval first
	CreatedAt   string `json:"created_at"`      // The time the account was created
	Followers   int    `json:"followers_count"` // The number of followers for the account
	Following   int    `json:"following_count"` // The number of accounts the given account is following
	Statuses    int    `json:"statuses_count"`  // The number of statuses the account has made
}

// Application holds informations about an application.
type Application struct {
	ID           string `json:"id"`
	Name         string `json:"name"`    // Name of the app
	Website      string `json:"website"` // Homepage URL of the app
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Attachment holds informations about an attachment.
type Attachment struct {
	ID         string `json:"id"`          // ID of the attachment
	Type       string `json:"type"`        // One of: "image", "video", "gifv"
	URL        string `json:"url"`         // URL of the locally hosted version of the image
	RemoteURL  string `json:"remote_url"`  // For remote images, the remote URL of the original image
	PreviewURL string `json:"preview_url"` // URL of the preview image
	TextURL    string `json:"text_url"`    // Shorter URL for the image, for insertion into text (only present on local images)
}

// Card holds informations about a card.
type Card struct {
	URL         string `json:"url"`         // The url associated with the card
	Title       string `json:"title"`       // The title of the card
	Description string `json:"description"` // The card description
	Image       string `json:"image"`       // The image associated with the card, if any
}

// Context holds informations about a statuses context.
type Context struct {
	Ancestors   []Status `json:"ancestors"`   // The ancestors of the status in the conversation, as a list of Statuses
	Descendants []Status `json:"descendants"` // The descendants of the status in the conversation, as a list of Statuses
}

// Error holds informations about an error.
type Error struct {
	Error string `json:"error"` // A textual description of the error
}

// Instance holds informations about an instance.
type Instance struct {
	URI         string `json:"uri"`         // URI of the current instance
	Title       string `json:"title"`       // The instance's title
	Description string `json:"description"` // A description for the instance
	Email       string `json:"email"`       // An email address which can be used to contact the instance administrator
}

// Mention holds informations about a mention.
type Mention struct {
	ID       string `json:"id"`       // Account ID
	URL      string `json:"url"`      // URL of user's profile (can be remote)
	Username string `json:"username"` // The username of the account
	Acct     string `json:"acct"`     // Equals username for local users, includes @domain for remote ones
}

// Notification holds informations about a notification.
type Notification struct {
	ID        string   `json:"id"`         // The notification ID
	Type      string   `json:"type"`       // One of: "mention", "reblog", "favourite", "follow"
	CreatedAt string   `json:"created_at"` // The time the notification was created
	Account   *Account `json:"account"`    // The Account sending the notification to the user
	Status    *Status  `json:"status"`     // The Status associated with the notification, if applicable
}

// Relationship holds informations about a relationship.
type Relationship struct {
	Following  bool `json:"following"`   // Whether the user is currently following the account
	FollowedBy bool `json:"followed_by"` // Whether the user is currently being followed by the account
	Blocking   bool `json:"blocking"`    // Whether the user is currently blocking the account
	Muting     bool `json:"muting"`      // Whether the user is currently muting the account
	Requested  bool `json:"requested"`   // Whether the user has requested to follow the account
}

// Report holds informations about a report.
type Report struct {
	ID          string `json:"id"`           // The ID of the report
	ActionTaken string `json:"action_taken"` // The action taken in response to the report
}

// Results holds informations about results.
type Results struct {
	Accounts []Account `json:"accounts"` // An array of matched Accounts
	Statuses []Status  `json:"statuses"` // An array of matchhed Statuses
	Hashtags []string  `json:"hashtags"` // An array of matched hashtags, as strings
}

// Status holds informations about a status.
type Status struct {
	ID                 string       `json:"id"`                     // The ID of the status
	URI                string       `json:"uri"`                    // A Fediverse-unique resource ID
	URL                string       `json:"url"`                    // URL to the status page (can be remote)
	Account            *Account     `json:"account"`                // The Account which posted the status
	InReplyToID        string       `json:"in_reply_to_id"`         // null or the ID of the status it replies to
	InReplyToAccountID string       `json:"in_reply_to_account_id"` // null or the ID of the account it replies to
	Reblog             *Status      `json:"reblog"`                 // null or the reblogged Status
	Content            string       `json:"content"`                // Body of the status; this will contain HTML (remote HTML already sanitized)
	CreatedAt          string       `json:"created_at"`             // The time the status was created
	Reblogs            int          `json:"reblogs_count"`          // The number of reblogs for the status
	Favourites         int          `json:"favourites_count"`       // The number of favourites for the status
	Reblogged          bool         `json:"reblogged"`              // Whether the authenticated user has reblogged the status
	Favourited         bool         `json:"favourited"`             // Whether the authenticated user has favourited the status
	Sensitive          bool         `json:"sensitive"`              // Whether media attachments should be hidden by default
	SpoilerText        string       `json:"spoiler_text"`           // If not empty, warning text that should be displayed before the actual content
	Visibility         string       `json:"visibility"`             // One of: public, unlisted, private, direct
	MediaAttachments   []Attachment `json:"media_attachments"`      // An array of Attachments
	Mentions           []Mention    `json:"mentions"`               // An array of Mentions
	Tags               []Tag        `json:"tags"`                   // An array of Tags
	Application        *Application `json:"application"`            // Application from which the status was posted
}

// Tag holds informations about a tag.
type Tag struct {
	Name string `json:"name"` // The hashtag, not including the preceding #
	URL  string `json:"url"`  // The URL of the hashtag
}
