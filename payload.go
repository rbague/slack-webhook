package webhook

// Predefined slack attachment colors
const (
	GoodColor    = "good"
	WarningColor = "warning"
	DangerColor  = "danger"
)

// Payload what to send to the Incoming Webhook API
type Payload struct {
	Text string `json:"text"`

	// The channel where to send the payload to, or the configured channel
	// Can be both a channel '#other-channel' or a direct message '@username'
	Channel     string        `json:"channel,omitempty"`
	UserName    string        `json:"username,omitempty"`
	IconURL     string        `json:"icon_url,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"` // :ghost:
	UnfurlLinks bool          `json:"unfurl_links,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`

	// Markdown used to disable markdown formatting on the text field
	Markdown *bool `json:"mrkdwn,omitempty"`
}

// Attachment adds more context to a message, making them more usable
type Attachment struct {
	// Fallback is a plain-text summary.
	//
	// Will be used in clients that do not show formatted text,
	// such as IRC, mobile and desktop notifications
	Fallback string `json:"fallback,omitempty"`

	// Use one of the predefined constants (*Color) or pass an hex color
	Color      string  `json:"color,omitempty"`
	Pretext    string  `json:"pretext,omitempty"`
	AuthorName string  `json:"author_name,omitempty"`
	AuthorLink string  `json:"author_link,omitempty"` // author_name required
	AuthorIcon string  `json:"author_icon,omitempty"` // author_name required
	Title      string  `json:"title,omitempty"`
	TitleLink  string  `json:"title_link,omitempty"` // title required
	Text       string  `json:"text,omitempty"`
	Fields     []Field `json:"fields,omitempty"`
	ImageURL   string  `json:"image_url,omitempty"`
	ThumbURL   string  `json:"thumb_url,omitempty"`
	Footer     string  `json:"footer,omitempty"`
	FooterIcon string  `json:"footer_icon,omitempty"` // footer required

	// Timestamp an epoch timestamp that will be displayed in
	// the attachment footer (time.Unix())
	Timestamp int64 `json:"ts,omitempty"`

	// MarkdownIn an array with field names to apply markdown formatting to
	MarkdownIn []string `json:"mrkdwn_in,omitempty"`
}

// Field key-value pairs displayed in a table-like way
type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	// Whether it is short enough to be displayed side-by-side with other fields
	Short bool `json:"short,omitempty"`
}
