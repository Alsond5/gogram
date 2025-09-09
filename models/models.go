package models

type Update struct {
	UpdateId                int64                        `json:"update_id"`
	Message                 *Message                     `json:"message,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
}

type WebhookInfo struct {
	Url                          string   `json:"url"`
	HasCustomCertificate         bool     `json:"has_custom_certificate"`
	PendingUpdateCount           int64    `json:"pending_update_count"`
	IpAddress                    string   `json:"ip_address,omitempty"`
	LastErrorDate                int64    `json:"last_error_date,omitempty"`
	LastErrorMessage             string   `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate int64    `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               int64    `json:"max_connections,omitempty"`
	AllowedUpdates               []string `json:"allowed_updates,omitempty"`
}

type User struct {
	Id                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`
}

type Chat struct {
	Id               int64  `json:"id"`
	Type             string `json:"type"`
	Title            string `json:"title,omitempty"`
	Username         string `json:"username,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	IsForum          bool   `json:"is_forum,omitempty"`
	IsDirectMessages bool   `json:"is_direct_messages,omitempty"`
}

type ChatFullInfo struct {
	Id                                 int64                 `json:"id"`
	Type                               string                `json:"type"`
	Title                              string                `json:"title,omitempty"`
	Username                           string                `json:"username,omitempty"`
	FirstName                          string                `json:"first_name,omitempty"`
	LastName                           string                `json:"last_name,omitempty"`
	IsForum                            bool                  `json:"is_forum,omitempty"`
	IsDirectMessages                   bool                  `json:"is_direct_messages,omitempty"`
	AccentColorId                      int64                 `json:"accent_color_id"`
	MaxReactionCount                   int64                 `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	ParentChat                         *Chat                 `json:"parent_chat,omitempty"`
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiId            string                `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorId               int64                 `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiId     string                `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiId           string                `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int64                 `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string                `json:"bio,omitempty"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`
	Description                        string                `json:"description,omitempty"`
	InviteLink                         string                `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	AcceptedGiftTypes                  *AcceptedGiftTypes    `json:"accepted_gift_types"`
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      int64                 `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               int64                 `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              int64                 `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatId                       int64                 `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
}

type Message struct {
	MessageId                     int64                          `json:"message_id"`
	MessageThreadId               int64                          `json:"message_thread_id,omitempty"`
	DirectMessagesTopic           *DirectMessagesTopic           `json:"direct_messages_topic,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	SenderBoostCount              int64                          `json:"sender_boost_count,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	Date                          int64                          `json:"date"`
	BusinessConnectionId          string                         `json:"business_connection_id,omitempty"`
	Chat                          *Chat                          `json:"chat"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ReplyToChecklistTaskId        int64                          `json:"reply_to_checklist_task_id,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int64                          `json:"edit_date,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`
	IsPaidPost                    bool                           `json:"is_paid_post,omitempty"`
	MediaGroupId                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	PaidStarCount                 int64                          `json:"paid_star_count,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	SuggestedPostInfo             *SuggestedPostInfo             `json:"suggested_post_info,omitempty"`
	EffectId                      string                         `json:"effect_id,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`
	Checklist                     *Checklist                     `json:"checklist,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatId               int64                          `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId             int64                          `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	Gift                          *GiftInfo                      `json:"gift,omitempty"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	ChecklistTasksDone            *ChecklistTasksDone            `json:"checklist_tasks_done,omitempty"`
	ChecklistTasksAdded           *ChecklistTasksAdded           `json:"checklist_tasks_added,omitempty"`
	DirectMessagePriceChanged     *DirectMessagePriceChanged     `json:"direct_message_price_changed,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed,omitempty"`
	SuggestedPostApproved         *SuggestedPostApproved         `json:"suggested_post_approved,omitempty"`
	SuggestedPostApprovalFailed   *SuggestedPostApprovalFailed   `json:"suggested_post_approval_failed,omitempty"`
	SuggestedPostDeclined         *SuggestedPostDeclined         `json:"suggested_post_declined,omitempty"`
	SuggestedPostPaid             *SuggestedPostPaid             `json:"suggested_post_paid,omitempty"`
	SuggestedPostRefunded         *SuggestedPostRefunded         `json:"suggested_post_refunded,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

type MessageId struct {
	MessageId int64 `json:"message_id"`
}

type InaccessibleMessage struct {
	Chat      *Chat `json:"chat"`
	MessageId int64 `json:"message_id"`
	Date      int64 `json:"date"`
}

// contains subtypes
type MaybeInaccessibleMessage struct {
	Type MaybeInaccessibleMessageType

	Message             *Message
	InaccessibleMessage *InaccessibleMessage
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int64  `json:"offset"`
	Length        int64  `json:"length"`
	Url           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities,omitempty"`
	Position int64           `json:"position"`
	IsManual bool            `json:"is_manual,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"origin"`
	Chat               *Chat               `json:"chat,omitempty"`
	MessageId          int64               `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              []PhotoSize         `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`
	Checklist          *Checklist          `json:"checklist,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

type ReplyParameters struct {
	MessageId                int64           `json:"message_id"`
	ChatId                   any             `json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	Quote                    string          `json:"quote,omitempty"`
	QuoteParseMode           string          `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            int64           `json:"quote_position,omitempty"`
	ChecklistTaskId          int64           `json:"checklist_task_id,omitempty"`
}

// contains subtypes
type MessageOrigin struct {
	Type MessageOriginType

	MessageOriginUser       *MessageOriginUser
	MessageOriginHiddenUser *MessageOriginHiddenUser
	MessageOriginChat       *MessageOriginChat
	MessageOriginChannel    *MessageOriginChannel
}

type MessageOriginUser struct {
	Type       string `json:"type"`
	Date       int64  `json:"date"`
	SenderUser *User  `json:"sender_user"`
}

type MessageOriginHiddenUser struct {
	Type           string `json:"type"`
	Date           int64  `json:"date"`
	SenderUserName string `json:"sender_user_name"`
}

type MessageOriginChat struct {
	Type            string `json:"type"`
	Date            int64  `json:"date"`
	SenderChat      *Chat  `json:"sender_chat"`
	AuthorSignature string `json:"author_signature,omitempty"`
}

type MessageOriginChannel struct {
	Type            string `json:"type"`
	Date            int64  `json:"date"`
	Chat            *Chat  `json:"chat"`
	MessageId       int64  `json:"message_id"`
	AuthorSignature string `json:"author_signature,omitempty"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int64      `json:"width"`
	Height       int64      `json:"height"`
	Duration     int64      `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int64      `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Story struct {
	Chat *Chat `json:"chat"`
	Id   int64 `json:"id"`
}

type Video struct {
	FileId         string      `json:"file_id"`
	FileUniqueId   string      `json:"file_unique_id"`
	Width          int64       `json:"width"`
	Height         int64       `json:"height"`
	Duration       int64       `json:"duration"`
	Thumbnail      *PhotoSize  `json:"thumbnail,omitempty"`
	Cover          []PhotoSize `json:"cover,omitempty"`
	StartTimestamp int64       `json:"start_timestamp,omitempty"`
	FileName       string      `json:"file_name,omitempty"`
	MimeType       string      `json:"mime_type,omitempty"`
	FileSize       int64       `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	Duration     int64      `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int64  `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type PaidMediaInfo struct {
	StarCount int64       `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

// contains subtypes
type PaidMedia struct {
	Type PaidMediaType

	PaidMediaPreview *PaidMediaPreview
	PaidMediaPhoto   *PaidMediaPhoto
	PaidMediaVideo   *PaidMediaVideo
}

type PaidMediaPreview struct {
	Type     string `json:"type"`
	Width    int64  `json:"width,omitempty"`
	Height   int64  `json:"height,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}

type PaidMediaPhoto struct {
	Type  string      `json:"type"`
	Photo []PhotoSize `json:"photo"`
}

type PaidMediaVideo struct {
	Type  string `json:"type"`
	Video *Video `json:"video"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int64  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int64  `json:"value"`
}

type PollOption struct {
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int64           `json:"voter_count"`
}

type InputPollOption struct {
	Text          string          `json:"text"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

type PollAnswer struct {
	PollId    string  `json:"poll_id"`
	VoterChat *Chat   `json:"voter_chat,omitempty"`
	User      *User   `json:"user,omitempty"`
	OptionIds []int64 `json:"option_ids"`
}

type Poll struct {
	Id                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEntities      []MessageEntity `json:"question_entities,omitempty"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int64           `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionId       int64           `json:"correct_option_id,omitempty"`
	Explanation           string          `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int64           `json:"open_period,omitempty"`
	CloseDate             int64           `json:"close_date,omitempty"`
}

type ChecklistTask struct {
	Id              int64           `json:"id"`
	Text            string          `json:"text"`
	TextEntities    []MessageEntity `json:"text_entities,omitempty"`
	CompletedByUser *User           `json:"completed_by_user,omitempty"`
	CompletionDate  int64           `json:"completion_date,omitempty"`
}

type Checklist struct {
	Title                    string          `json:"title"`
	TitleEntities            []MessageEntity `json:"title_entities,omitempty"`
	Tasks                    []ChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool            `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool            `json:"others_can_mark_tasks_as_done,omitempty"`
}

type InputChecklistTask struct {
	Id           int64           `json:"id"`
	Text         string          `json:"text"`
	ParseMode    string          `json:"parse_mode,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                string               `json:"parse_mode,omitempty"`
	TitleEntities            []MessageEntity      `json:"title_entities,omitempty"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done,omitempty"`
}

type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message,omitempty"`
	MarkedAsDoneTaskIds    []int64  `json:"marked_as_done_task_ids,omitempty"`
	MarkedAsNotDoneTaskIds []int64  `json:"marked_as_not_done_task_ids,omitempty"`
}

type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message,omitempty"`
	Tasks            []ChecklistTask `json:"tasks"`
}

type Location struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64   `json:"live_period,omitempty"`
	Heading              int64   `json:"heading,omitempty"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`
}

type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareId    string    `json:"foursquare_id,omitempty"`
	FoursquareType  string    `json:"foursquare_type,omitempty"`
	GooglePlaceId   string    `json:"google_place_id,omitempty"`
	GooglePlaceType string    `json:"google_place_type,omitempty"`
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int64 `json:"distance"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`
}

type ChatBoostAdded struct {
	BoostCount int64 `json:"boost_count"`
}

// contains subtypes
type BackgroundFill struct {
	Type BackgroundFillType

	BackgroundFillSolid            *BackgroundFillSolid
	BackgroundFillGradient         *BackgroundFillGradient
	BackgroundFillFreeformGradient *BackgroundFillFreeformGradient
}

type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int64  `json:"color"`
}

type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int64  `json:"top_color"`
	BottomColor   int64  `json:"bottom_color"`
	RotationAngle int64  `json:"rotation_angle"`
}

type BackgroundFillFreeformGradient struct {
	Type   string  `json:"type"`
	Colors []int64 `json:"colors"`
}

// contains subtypes
type BackgroundType struct {
	Type BackgroundTypeType

	BackgroundTypeFill      *BackgroundTypeFill
	BackgroundTypeWallpaper *BackgroundTypeWallpaper
	BackgroundTypePattern   *BackgroundTypePattern
	BackgroundTypeChatTheme *BackgroundTypeChatTheme
}

type BackgroundTypeFill struct {
	Type             string          `json:"type"`
	Fill             *BackgroundFill `json:"fill"`
	DarkThemeDimming int64           `json:"dark_theme_dimming"`
}

type BackgroundTypeWallpaper struct {
	Type             string    `json:"type"`
	Document         *Document `json:"document"`
	DarkThemeDimming int64     `json:"dark_theme_dimming"`
	IsBlurred        bool      `json:"is_blurred,omitempty"`
	IsMoving         bool      `json:"is_moving,omitempty"`
}

type BackgroundTypePattern struct {
	Type       string          `json:"type"`
	Document   *Document       `json:"document"`
	Fill       *BackgroundFill `json:"fill"`
	Intensity  int64           `json:"intensity"`
	IsInverted bool            `json:"is_inverted,omitempty"`
	IsMoving   bool            `json:"is_moving,omitempty"`
}

type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

type ChatBackground struct {
	Type *BackgroundType `json:"type"`
}

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicClosed struct{}

type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicReopened struct{}

type GeneralForumTopicHidden struct{}

type GeneralForumTopicUnhidden struct{}

type SharedUser struct {
	UserId    int64       `json:"user_id"`
	FirstName string      `json:"first_name,omitempty"`
	LastName  string      `json:"last_name,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

type UsersShared struct {
	RequestId int64        `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

type ChatShared struct {
	RequestId int64       `json:"request_id"`
	ChatId    int64       `json:"chat_id"`
	Title     string      `json:"title,omitempty"`
	Username  string      `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

type VideoChatStarted struct{}

type VideoChatEnded struct {
	Duration int64 `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

type PaidMessagePriceChanged struct {
	PaidMessageStarCount int64 `json:"paid_message_star_count"`
}

type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool  `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   int64 `json:"direct_message_star_count,omitempty"`
}

type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
	SendDate             int64               `json:"send_date"`
}

type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price"`
}

type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Comment              string   `json:"comment,omitempty"`
}

type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message,omitempty"`
	Currency             string      `json:"currency"`
	Amount               int64       `json:"amount,omitempty"`
	StarAmount           *StarAmount `json:"star_amount,omitempty"`
}

type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Reason               string   `json:"reason"`
}

type GiveawayCreated struct {
	PrizeStarCount int64 `json:"prize_star_count,omitempty"`
}

type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int64    `json:"winners_selection_date"`
	WinnerCount                   int64    `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PrizeStarCount                int64    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int64    `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          *Chat  `json:"chat"`
	GiveawayMessageId             int64  `json:"giveaway_message_id"`
	WinnersSelectionDate          int64  `json:"winners_selection_date"`
	WinnerCount                   int64  `json:"winner_count"`
	Winners                       []User `json:"winners"`
	AdditionalChatCount           int64  `json:"additional_chat_count,omitempty"`
	PrizeStarCount                int64  `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int64  `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int64  `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`
	WasRefunded                   bool   `json:"was_refunded,omitempty"`
	PrizeDescription              string `json:"prize_description,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int64    `json:"winner_count"`
	UnclaimedPrizeCount int64    `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`
}

type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	Url              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

type SuggestedPostInfo struct {
	State    string              `json:"state"`
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date,omitempty"`
}

type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date,omitempty"`
}

type DirectMessagesTopic struct {
	TopicId int64 `json:"topic_id"`
	User    *User `json:"user,omitempty"`
}

type UserProfilePhotos struct {
	TotalCount int64         `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type WebAppInfo struct {
	Url string `json:"url"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
}

type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  bool                        `json:"request_contact,omitempty"`
	RequestLocation bool                        `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
}

type KeyboardButtonRequestUsers struct {
	RequestId       int64 `json:"request_id"`
	UserIsBot       bool  `json:"user_is_bot,omitempty"`
	UserIsPremium   bool  `json:"user_is_premium,omitempty"`
	MaxQuantity     int64 `json:"max_quantity,omitempty"`
	RequestName     bool  `json:"request_name,omitempty"`
	RequestUsername bool  `json:"request_username,omitempty"`
	RequestPhoto    bool  `json:"request_photo,omitempty"`
}

type KeyboardButtonRequestChat struct {
	RequestId               int64                    `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`
	RequestTitle            bool                     `json:"request_title,omitempty"`
	RequestUsername         bool                     `json:"request_username,omitempty"`
	RequestPhoto            bool                     `json:"request_photo,omitempty"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	Url                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

type CopyTextButton struct {
	Text string `json:"text"`
}

type CallbackQuery struct {
	Id              string                    `json:"id"`
	From            *User                     `json:"from"`
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageId string                    `json:"inline_message_id,omitempty"`
	ChatInstance    string                    `json:"chat_instance"`
	Data            string                    `json:"data,omitempty"`
	GameShortName   string                    `json:"game_short_name,omitempty"`
}

type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 *User  `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name,omitempty"`
	ExpireDate              int64  `json:"expire_date,omitempty"`
	MemberLimit             int64  `json:"member_limit,omitempty"`
	PendingJoinRequestCount int64  `json:"pending_join_request_count,omitempty"`
	SubscriptionPeriod      int64  `json:"subscription_period,omitempty"`
	SubscriptionPrice       int64  `json:"subscription_price,omitempty"`
}

type ChatAdministratorRights struct {
	IsAnonymous             bool `json:"is_anonymous"`
	CanManageChat           bool `json:"can_manage_chat"`
	CanDeleteMessages       bool `json:"can_delete_messages"`
	CanManageVideoChats     bool `json:"can_manage_video_chats"`
	CanRestrictMembers      bool `json:"can_restrict_members"`
	CanPromoteMembers       bool `json:"can_promote_members"`
	CanChangeInfo           bool `json:"can_change_info"`
	CanInviteUsers          bool `json:"can_invite_users"`
	CanPostStories          bool `json:"can_post_stories"`
	CanEditStories          bool `json:"can_edit_stories"`
	CanDeleteStories        bool `json:"can_delete_stories"`
	CanPostMessages         bool `json:"can_post_messages,omitempty"`
	CanEditMessages         bool `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool `json:"can_manage_direct_messages,omitempty"`
}

type ChatMemberUpdated struct {
	Chat                    *Chat           `json:"chat"`
	From                    *User           `json:"from"`
	Date                    int64           `json:"date"`
	OldChatMember           *ChatMember     `json:"old_chat_member"`
	NewChatMember           *ChatMember     `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
}

// contains subtypes
type ChatMember struct {
	Type ChatMemberType

	ChatMemberOwner         *ChatMemberOwner
	ChatMemberAdministrator *ChatMemberAdministrator
	ChatMemberMember        *ChatMemberMember
	ChatMemberRestricted    *ChatMemberRestricted
	ChatMemberLeft          *ChatMemberLeft
	ChatMemberBanned        *ChatMemberBanned
}

type ChatMemberOwner struct {
	Status      string `json:"status"`
	User        *User  `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title,omitempty"`
}

type ChatMemberAdministrator struct {
	Status                  string `json:"status"`
	User                    *User  `json:"user"`
	CanBeEdited             bool   `json:"can_be_edited"`
	IsAnonymous             bool   `json:"is_anonymous"`
	CanManageChat           bool   `json:"can_manage_chat"`
	CanDeleteMessages       bool   `json:"can_delete_messages"`
	CanManageVideoChats     bool   `json:"can_manage_video_chats"`
	CanRestrictMembers      bool   `json:"can_restrict_members"`
	CanPromoteMembers       bool   `json:"can_promote_members"`
	CanChangeInfo           bool   `json:"can_change_info"`
	CanInviteUsers          bool   `json:"can_invite_users"`
	CanPostStories          bool   `json:"can_post_stories"`
	CanEditStories          bool   `json:"can_edit_stories"`
	CanDeleteStories        bool   `json:"can_delete_stories"`
	CanPostMessages         bool   `json:"can_post_messages,omitempty"`
	CanEditMessages         bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool   `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool   `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool   `json:"can_manage_direct_messages,omitempty"`
	CustomTitle             string `json:"custom_title,omitempty"`
}

type ChatMemberMember struct {
	Status    string `json:"status"`
	User      *User  `json:"user"`
	UntilDate int64  `json:"until_date,omitempty"`
}

type ChatMemberRestricted struct {
	Status                string `json:"status"`
	User                  *User  `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	UntilDate             int64  `json:"until_date"`
}

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   *User  `json:"user"`
}

type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      *User  `json:"user"`
	UntilDate int64  `json:"until_date"`
}

type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`
	From       *User           `json:"from"`
	UserChatId int64           `json:"user_chat_id"`
	Date       int64           `json:"date"`
	Bio        string          `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

type Birthdate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year,omitempty"`
}

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message string   `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int64 `json:"opening_minute"`
	ClosingMinute int64 `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

type LocationAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
}

// contains subtypes
type StoryAreaType struct {
	Type StoryAreaTypeType

	StoryAreaTypeLocation          *StoryAreaTypeLocation
	StoryAreaTypeSuggestedReaction *StoryAreaTypeSuggestedReaction
	StoryAreaTypeLink              *StoryAreaTypeLink
	StoryAreaTypeWeather           *StoryAreaTypeWeather
	StoryAreaTypeUniqueGift        *StoryAreaTypeUniqueGift
}

type StoryAreaTypeLocation struct {
	Type      string           `json:"type"`
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
	Address   *LocationAddress `json:"address,omitempty"`
}

type StoryAreaTypeSuggestedReaction struct {
	Type         string        `json:"type"`
	ReactionType *ReactionType `json:"reaction_type"`
	IsDark       bool          `json:"is_dark,omitempty"`
	IsFlipped    bool          `json:"is_flipped,omitempty"`
}

type StoryAreaTypeLink struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type StoryAreaTypeWeather struct {
	Type            string  `json:"type"`
	Temperature     float64 `json:"temperature"`
	Emoji           string  `json:"emoji"`
	BackgroundColor int64   `json:"background_color"`
}

type StoryAreaTypeUniqueGift struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type StoryArea struct {
	Position *StoryAreaPosition `json:"position"`
	Type     *StoryAreaType     `json:"type"`
}

type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

// contains subtypes
type ReactionType struct {
	Type ReactionTypeType

	ReactionTypeEmoji       *ReactionTypeEmoji
	ReactionTypeCustomEmoji *ReactionTypeCustomEmoji
	ReactionTypePaid        *ReactionTypePaid
}

type ReactionTypeEmoji struct {
	Type  string `json:"type"`
	Emoji string `json:"emoji"`
}

type ReactionTypeCustomEmoji struct {
	Type          string `json:"type"`
	CustomEmojiId string `json:"custom_emoji_id"`
}

type ReactionTypePaid struct {
	Type string `json:"type"`
}

type ReactionCount struct {
	Type       *ReactionType `json:"type"`
	TotalCount int64         `json:"total_count"`
}

type MessageReactionUpdated struct {
	Chat        *Chat          `json:"chat"`
	MessageId   int64          `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int64          `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

type MessageReactionCountUpdated struct {
	Chat      *Chat           `json:"chat"`
	MessageId int64           `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

type ForumTopic struct {
	MessageThreadId   int64  `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type Gift struct {
	Id               string   `json:"id"`
	Sticker          *Sticker `json:"sticker"`
	StarCount        int64    `json:"star_count"`
	UpgradeStarCount int64    `json:"upgrade_star_count,omitempty"`
	TotalCount       int64    `json:"total_count,omitempty"`
	RemainingCount   int64    `json:"remaining_count,omitempty"`
	PublisherChat    *Chat    `json:"publisher_chat,omitempty"`
}

type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

type UniqueGiftModel struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int64    `json:"rarity_per_mille"`
}

type UniqueGiftSymbol struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int64    `json:"rarity_per_mille"`
}

type UniqueGiftBackdropColors struct {
	CenterColor int64 `json:"center_color"`
	EdgeColor   int64 `json:"edge_color"`
	SymbolColor int64 `json:"symbol_color"`
	TextColor   int64 `json:"text_color"`
}

type UniqueGiftBackdrop struct {
	Name           string                    `json:"name"`
	Colors         *UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int64                     `json:"rarity_per_mille"`
}

type UniqueGift struct {
	BaseName      string              `json:"base_name"`
	Name          string              `json:"name"`
	Number        int64               `json:"number"`
	Model         *UniqueGiftModel    `json:"model"`
	Symbol        *UniqueGiftSymbol   `json:"symbol"`
	Backdrop      *UniqueGiftBackdrop `json:"backdrop"`
	PublisherChat *Chat               `json:"publisher_chat,omitempty"`
}

type GiftInfo struct {
	Gift                    *Gift           `json:"gift"`
	OwnedGiftId             string          `json:"owned_gift_id,omitempty"`
	ConvertStarCount        int64           `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int64           `json:"prepaid_upgrade_star_count,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
}

type UniqueGiftInfo struct {
	Gift                *UniqueGift `json:"gift"`
	Origin              string      `json:"origin"`
	LastResaleStarCount int64       `json:"last_resale_star_count,omitempty"`
	OwnedGiftId         string      `json:"owned_gift_id,omitempty"`
	TransferStarCount   int64       `json:"transfer_star_count,omitempty"`
	NextTransferDate    int64       `json:"next_transfer_date,omitempty"`
}

// contains subtypes
type OwnedGift struct {
	Type OwnedGiftType

	OwnedGiftRegular *OwnedGiftRegular
	OwnedGiftUnique  *OwnedGiftUnique
}

type OwnedGiftRegular struct {
	Type                    string          `json:"type"`
	Gift                    *Gift           `json:"gift"`
	OwnedGiftId             string          `json:"owned_gift_id,omitempty"`
	SenderUser              *User           `json:"sender_user,omitempty"`
	SendDate                int64           `json:"send_date"`
	Text                    string          `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool            `json:"is_private,omitempty"`
	IsSaved                 bool            `json:"is_saved,omitempty"`
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`
	WasRefunded             bool            `json:"was_refunded,omitempty"`
	ConvertStarCount        int64           `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int64           `json:"prepaid_upgrade_star_count,omitempty"`
}

type OwnedGiftUnique struct {
	Type              string      `json:"type"`
	Gift              *UniqueGift `json:"gift"`
	OwnedGiftId       string      `json:"owned_gift_id,omitempty"`
	SenderUser        *User       `json:"sender_user,omitempty"`
	SendDate          int64       `json:"send_date"`
	IsSaved           bool        `json:"is_saved,omitempty"`
	CanBeTransferred  bool        `json:"can_be_transferred,omitempty"`
	TransferStarCount int64       `json:"transfer_star_count,omitempty"`
	NextTransferDate  int64       `json:"next_transfer_date,omitempty"`
}

type OwnedGifts struct {
	TotalCount int64       `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset string      `json:"next_offset,omitempty"`
}

type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
}

type StarAmount struct {
	Amount         int64 `json:"amount"`
	NanostarAmount int64 `json:"nanostar_amount,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// contains subtypes
type BotCommandScope struct {
	Type BotCommandScopeType

	BotCommandScopeDefault               *BotCommandScopeDefault
	BotCommandScopeAllPrivateChats       *BotCommandScopeAllPrivateChats
	BotCommandScopeAllGroupChats         *BotCommandScopeAllGroupChats
	BotCommandScopeAllChatAdministrators *BotCommandScopeAllChatAdministrators
	BotCommandScopeChat                  *BotCommandScopeChat
	BotCommandScopeChatAdministrators    *BotCommandScopeChatAdministrators
	BotCommandScopeChatMember            *BotCommandScopeChatMember
}

type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

type BotCommandScopeAllPrivateChats struct {
	Type string `json:"type"`
}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

type BotCommandScopeAllChatAdministrators struct {
	Type string `json:"type"`
}

type BotCommandScopeChat struct {
	Type   string `json:"type"`
	ChatId any    `json:"chat_id"`
}

type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatId any    `json:"chat_id"`
}

type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatId any    `json:"chat_id"`
	UserId int64  `json:"user_id"`
}

type BotName struct {
	Name string `json:"name"`
}

type BotDescription struct {
	Description string `json:"description"`
}

type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// contains subtypes
type MenuButton struct {
	Type MenuButtonType

	MenuButtonCommands *MenuButtonCommands
	MenuButtonWebApp   *MenuButtonWebApp
	MenuButtonDefault  *MenuButtonDefault
}

type MenuButtonCommands struct {
	Type string `json:"type"`
}

type MenuButtonWebApp struct {
	Type   string      `json:"type"`
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app"`
}

type MenuButtonDefault struct {
	Type string `json:"type"`
}

// contains subtypes
type ChatBoostSource struct {
	Type ChatBoostSourceType

	ChatBoostSourcePremium  *ChatBoostSourcePremium
	ChatBoostSourceGiftCode *ChatBoostSourceGiftCode
	ChatBoostSourceGiveaway *ChatBoostSourceGiveaway
}

type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   *User  `json:"user"`
}

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   *User  `json:"user"`
}

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`
	GiveawayMessageId int64  `json:"giveaway_message_id"`
	User              *User  `json:"user,omitempty"`
	PrizeStarCount    int64  `json:"prize_star_count,omitempty"`
	IsUnclaimed       bool   `json:"is_unclaimed,omitempty"`
}

type ChatBoost struct {
	BoostId        string           `json:"boost_id"`
	AddDate        int64            `json:"add_date"`
	ExpirationDate int64            `json:"expiration_date"`
	Source         *ChatBoostSource `json:"source"`
}

type ChatBoostUpdated struct {
	Chat  *Chat      `json:"chat"`
	Boost *ChatBoost `json:"boost"`
}

type ChatBoostRemoved struct {
	Chat       *Chat            `json:"chat"`
	BoostId    string           `json:"boost_id"`
	RemoveDate int64            `json:"remove_date"`
	Source     *ChatBoostSource `json:"source"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply,omitempty"`
	CanReadMessages            bool `json:"can_read_messages,omitempty"`
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages,omitempty"`
	CanDeleteAllMessages       bool `json:"can_delete_all_messages,omitempty"`
	CanEditName                bool `json:"can_edit_name,omitempty"`
	CanEditBio                 bool `json:"can_edit_bio,omitempty"`
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo,omitempty"`
	CanEditUsername            bool `json:"can_edit_username,omitempty"`
	CanChangeGiftSettings      bool `json:"can_change_gift_settings,omitempty"`
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars,omitempty"`
	CanConvertGiftsToStars     bool `json:"can_convert_gifts_to_stars,omitempty"`
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	CanTransferStars           bool `json:"can_transfer_stars,omitempty"`
	CanManageStories           bool `json:"can_manage_stories,omitempty"`
}

type BusinessConnection struct {
	Id         string             `json:"id"`
	User       *User              `json:"user"`
	UserChatId int64              `json:"user_chat_id"`
	Date       int64              `json:"date"`
	Rights     *BusinessBotRights `json:"rights,omitempty"`
	IsEnabled  bool               `json:"is_enabled"`
}

type BusinessMessagesDeleted struct {
	BusinessConnectionId string  `json:"business_connection_id"`
	Chat                 *Chat   `json:"chat"`
	MessageIds           []int64 `json:"message_ids"`
}

type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int64 `json:"retry_after,omitempty"`
}

// contains subtypes
type InputMedia struct {
	Type InputMediaType

	InputMediaAnimation *InputMediaAnimation
	InputMediaDocument  *InputMediaDocument
	InputMediaAudio     *InputMediaAudio
	InputMediaPhoto     *InputMediaPhoto
	InputMediaVideo     *InputMediaVideo
}

type InputMediaPhoto struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

type InputMediaVideo struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             string          `json:"thumbnail,omitempty"`
	Cover                 string          `json:"cover,omitempty"`
	StartTimestamp        int64           `json:"start_timestamp,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int64           `json:"width,omitempty"`
	Height                int64           `json:"height,omitempty"`
	Duration              int64           `json:"duration,omitempty"`
	SupportsStreaming     bool            `json:"supports_streaming,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

type InputMediaAnimation struct {
	Type                  string          `json:"type"`
	Media                 string          `json:"media"`
	Thumbnail             string          `json:"thumbnail,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             string          `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int64           `json:"width,omitempty"`
	Height                int64           `json:"height,omitempty"`
	Duration              int64           `json:"duration,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

type InputMediaAudio struct {
	Type            string          `json:"type"`
	Media           string          `json:"media"`
	Thumbnail       string          `json:"thumbnail,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int64           `json:"duration,omitempty"`
	Performer       string          `json:"performer,omitempty"`
	Title           string          `json:"title,omitempty"`
}

type InputMediaDocument struct {
	Type                        string          `json:"type"`
	Media                       string          `json:"media"`
	Thumbnail                   string          `json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   string          `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}

type InputFile struct{}

// contains subtypes
type InputPaidMedia struct {
	Type InputPaidMediaType

	InputPaidMediaPhoto *InputPaidMediaPhoto
	InputPaidMediaVideo *InputPaidMediaVideo
}

type InputPaidMediaPhoto struct {
	Type  string `json:"type"`
	Media string `json:"media"`
}

type InputPaidMediaVideo struct {
	Type              string `json:"type"`
	Media             string `json:"media"`
	Thumbnail         string `json:"thumbnail,omitempty"`
	Cover             string `json:"cover,omitempty"`
	StartTimestamp    int64  `json:"start_timestamp,omitempty"`
	Width             int64  `json:"width,omitempty"`
	Height            int64  `json:"height,omitempty"`
	Duration          int64  `json:"duration,omitempty"`
	SupportsStreaming bool   `json:"supports_streaming,omitempty"`
}

// contains subtypes
type InputProfilePhoto struct {
	Type InputProfilePhotoType

	InputProfilePhotoStatic   *InputProfilePhotoStatic
	InputProfilePhotoAnimated *InputProfilePhotoAnimated
}

type InputProfilePhotoStatic struct {
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

type InputProfilePhotoAnimated struct {
	Type               string  `json:"type"`
	Animation          string  `json:"animation"`
	MainFrameTimestamp float64 `json:"main_frame_timestamp,omitempty"`
}

// contains subtypes
type InputStoryContent struct {
	Type InputStoryContentType

	InputStoryContentPhoto *InputStoryContentPhoto
	InputStoryContentVideo *InputStoryContentVideo
}

type InputStoryContentPhoto struct {
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

type InputStoryContentVideo struct {
	Type                string  `json:"type"`
	Video               string  `json:"video"`
	Duration            float64 `json:"duration,omitempty"`
	CoverFrameTimestamp float64 `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         bool    `json:"is_animation,omitempty"`
}

type Sticker struct {
	FileId           string        `json:"file_id"`
	FileUniqueId     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int64         `json:"width"`
	Height           int64         `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string        `json:"emoji,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiId    string        `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool          `json:"needs_repainting,omitempty"`
	FileSize         int64         `json:"file_size,omitempty"`
}

type StickerSet struct {
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	StickerType string     `json:"sticker_type"`
	Stickers    []Sticker  `json:"stickers"`
	Thumbnail   *PhotoSize `json:"thumbnail,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type InputSticker struct {
	Sticker      string        `json:"sticker"`
	Format       string        `json:"format"`
	EmojiList    []string      `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     []string      `json:"keywords,omitempty"`
}

type InlineQuery struct {
	Id       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter string      `json:"start_parameter,omitempty"`
}

// contains subtypes
type InlineQueryResult struct {
	Type InlineQueryResultType

	InlineQueryResultCachedAudio    *InlineQueryResultCachedAudio
	InlineQueryResultCachedDocument *InlineQueryResultCachedDocument
	InlineQueryResultCachedGif      *InlineQueryResultCachedGif
	InlineQueryResultCachedMpeg4Gif *InlineQueryResultCachedMpeg4Gif
	InlineQueryResultCachedPhoto    *InlineQueryResultCachedPhoto
	InlineQueryResultCachedSticker  *InlineQueryResultCachedSticker
	InlineQueryResultCachedVideo    *InlineQueryResultCachedVideo
	InlineQueryResultCachedVoice    *InlineQueryResultCachedVoice
	InlineQueryResultArticle        *InlineQueryResultArticle
	InlineQueryResultAudio          *InlineQueryResultAudio
	InlineQueryResultContact        *InlineQueryResultContact
	InlineQueryResultGame           *InlineQueryResultGame
	InlineQueryResultDocument       *InlineQueryResultDocument
	InlineQueryResultGif            *InlineQueryResultGif
	InlineQueryResultLocation       *InlineQueryResultLocation
	InlineQueryResultMpeg4Gif       *InlineQueryResultMpeg4Gif
	InlineQueryResultPhoto          *InlineQueryResultPhoto
	InlineQueryResultVenue          *InlineQueryResultVenue
	InlineQueryResultVideo          *InlineQueryResultVideo
	InlineQueryResultVoice          *InlineQueryResultVoice
}

type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 string                `json:"url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultPhoto struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	PhotoUrl              string                `json:"photo_url"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	PhotoWidth            int64                 `json:"photo_width,omitempty"`
	PhotoHeight           int64                 `json:"photo_height,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultGif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	GifUrl                string                `json:"gif_url"`
	GifWidth              int64                 `json:"gif_width,omitempty"`
	GifHeight             int64                 `json:"gif_height,omitempty"`
	GifDuration           int64                 `json:"gif_duration,omitempty"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	Mpeg4Url              string                `json:"mpeg4_url"`
	Mpeg4Width            int64                 `json:"mpeg4_width,omitempty"`
	Mpeg4Height           int64                 `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         int64                 `json:"mpeg4_duration,omitempty"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	ThumbnailMimeType     string                `json:"thumbnail_mime_type,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	VideoUrl              string                `json:"video_url"`
	MimeType              string                `json:"mime_type"`
	ThumbnailUrl          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	VideoWidth            int64                 `json:"video_width,omitempty"`
	VideoHeight           int64                 `json:"video_height,omitempty"`
	VideoDuration         int64                 `json:"video_duration,omitempty"`
	Description           string                `json:"description,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	AudioUrl            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int64                 `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	VoiceUrl            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	VoiceDuration       int64                 `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	DocumentUrl         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	Id                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64                 `json:"live_period,omitempty"`
	Heading              int64                 `json:"heading,omitempty"`
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl         string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int64                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int64                 `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareId        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	GooglePlaceId       string                `json:"google_place_id,omitempty"`
	GooglePlaceType     string                `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	Vcard               string                `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailUrl        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int64                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int64                 `json:"thumbnail_height,omitempty"`
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	Id            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	PhotoFileId           string                `json:"photo_file_id"`
	Title                 string                `json:"title,omitempty"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedGif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	GifFileId             string                `json:"gif_file_id"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	Mpeg4FileId           string                `json:"mpeg4_file_id"`
	Title                 string                `json:"title,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	StickerFileId       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileId      string                `json:"document_file_id"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Type                  string                `json:"type"`
	Id                    string                `json:"id"`
	VideoFileId           string                `json:"video_file_id"`
	Title                 string                `json:"title"`
	Description           string                `json:"description,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	VoiceFileId         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	AudioFileId         string                `json:"audio_file_id"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// contains subtypes
type InputMessageContent struct {
	Type InputMessageContentType

	InputTextMessageContent     *InputTextMessageContent
	InputLocationMessageContent *InputLocationMessageContent
	InputVenueMessageContent    *InputVenueMessageContent
	InputContactMessageContent  *InputContactMessageContent
	InputInvoiceMessageContent  *InputInvoiceMessageContent
}

type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          string              `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int64   `json:"live_period,omitempty"`
	Heading              int64   `json:"heading,omitempty"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius,omitempty"`
}

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareId    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceId   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int64          `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64        `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoUrl                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int64          `json:"photo_size,omitempty"`
	PhotoWidth                int64          `json:"photo_width,omitempty"`
	PhotoHeight               int64          `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageId string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

type SentWebAppMessage struct {
	InlineMessageId string `json:"inline_message_id,omitempty"`
}

type PreparedInlineMessage struct {
	Id             string `json:"id"`
	ExpirationDate int64  `json:"expiration_date"`
}

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int64  `json:"amount"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int64  `json:"total_amount"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type ShippingOption struct {
	Id     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`
	TotalAmount                int64      `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate int64      `json:"subscription_expiration_date,omitempty"`
	IsRecurring                bool       `json:"is_recurring,omitempty"`
	IsFirstRecurring           bool       `json:"is_first_recurring,omitempty"`
	ShippingOptionId           string     `json:"shipping_option_id,omitempty"`
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeId    string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId    string     `json:"provider_payment_charge_id"`
}

type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int64  `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string `json:"provider_payment_charge_id,omitempty"`
}

type ShippingQuery struct {
	Id              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	Id               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionId string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

type PaidMediaPurchased struct {
	From             *User  `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

// contains subtypes
type RevenueWithdrawalState struct {
	Type RevenueWithdrawalStateType

	RevenueWithdrawalStatePending   *RevenueWithdrawalStatePending
	RevenueWithdrawalStateSucceeded *RevenueWithdrawalStateSucceeded
	RevenueWithdrawalStateFailed    *RevenueWithdrawalStateFailed
}

type RevenueWithdrawalStatePending struct {
	Type string `json:"type"`
}

type RevenueWithdrawalStateSucceeded struct {
	Type string `json:"type"`
	Date int64  `json:"date"`
	Url  string `json:"url"`
}

type RevenueWithdrawalStateFailed struct {
	Type string `json:"type"`
}

type AffiliateInfo struct {
	AffiliateUser      *User `json:"affiliate_user,omitempty"`
	AffiliateChat      *Chat `json:"affiliate_chat,omitempty"`
	CommissionPerMille int64 `json:"commission_per_mille"`
	Amount             int64 `json:"amount"`
	NanostarAmount     int64 `json:"nanostar_amount,omitempty"`
}

// contains subtypes
type TransactionPartner struct {
	Type TransactionPartnerType

	TransactionPartnerUser             *TransactionPartnerUser
	TransactionPartnerChat             *TransactionPartnerChat
	TransactionPartnerAffiliateProgram *TransactionPartnerAffiliateProgram
	TransactionPartnerFragment         *TransactionPartnerFragment
	TransactionPartnerTelegramAds      *TransactionPartnerTelegramAds
	TransactionPartnerTelegramApi      *TransactionPartnerTelegramApi
	TransactionPartnerOther            *TransactionPartnerOther
}

type TransactionPartnerUser struct {
	Type                        string         `json:"type"`
	TransactionType             string         `json:"transaction_type"`
	User                        *User          `json:"user"`
	Affiliate                   *AffiliateInfo `json:"affiliate,omitempty"`
	InvoicePayload              string         `json:"invoice_payload,omitempty"`
	SubscriptionPeriod          int64          `json:"subscription_period,omitempty"`
	PaidMedia                   []PaidMedia    `json:"paid_media,omitempty"`
	PaidMediaPayload            string         `json:"paid_media_payload,omitempty"`
	Gift                        *Gift          `json:"gift,omitempty"`
	PremiumSubscriptionDuration int64          `json:"premium_subscription_duration,omitempty"`
}

type TransactionPartnerChat struct {
	Type string `json:"type"`
	Chat *Chat  `json:"chat"`
	Gift *Gift  `json:"gift,omitempty"`
}

type TransactionPartnerAffiliateProgram struct {
	Type               string `json:"type"`
	SponsorUser        *User  `json:"sponsor_user,omitempty"`
	CommissionPerMille int64  `json:"commission_per_mille"`
}

type TransactionPartnerFragment struct {
	Type            string                  `json:"type"`
	WithdrawalState *RevenueWithdrawalState `json:"withdrawal_state,omitempty"`
}

type TransactionPartnerTelegramAds struct {
	Type string `json:"type"`
}

type TransactionPartnerTelegramApi struct {
	Type         string `json:"type"`
	RequestCount int64  `json:"request_count"`
}

type TransactionPartnerOther struct {
	Type string `json:"type"`
}

type StarTransaction struct {
	Id             string              `json:"id"`
	Amount         int64               `json:"amount"`
	NanostarAmount int64               `json:"nanostar_amount,omitempty"`
	Date           int64               `json:"date"`
	Source         *TransactionPartner `json:"source,omitempty"`
	Receiver       *TransactionPartner `json:"receiver,omitempty"`
}

type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials      `json:"credentials"`
}

type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// contains subtypes
type PassportElementError struct {
	Type PassportElementErrorType

	PassportElementErrorDataField        *PassportElementErrorDataField
	PassportElementErrorFrontSide        *PassportElementErrorFrontSide
	PassportElementErrorReverseSide      *PassportElementErrorReverseSide
	PassportElementErrorSelfie           *PassportElementErrorSelfie
	PassportElementErrorFile             *PassportElementErrorFile
	PassportElementErrorFiles            *PassportElementErrorFiles
	PassportElementErrorTranslationFile  *PassportElementErrorTranslationFile
	PassportElementErrorTranslationFiles *PassportElementErrorTranslationFiles
	PassportElementErrorUnspecified      *PassportElementErrorUnspecified
}

type PassportElementErrorDataField struct {
	Source    string `json:"source"`
	Type      string `json:"type"`
	FieldName string `json:"field_name"`
	DataHash  string `json:"data_hash"`
	Message   string `json:"message"`
}

type PassportElementErrorFrontSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorReverseSide struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorSelfie struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

type PassportElementErrorTranslationFile struct {
	Source   string `json:"source"`
	Type     string `json:"type"`
	FileHash string `json:"file_hash"`
	Message  string `json:"message"`
}

type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source"`
	Type       string   `json:"type"`
	FileHashes []string `json:"file_hashes"`
	Message    string   `json:"message"`
}

type PassportElementErrorUnspecified struct {
	Source      string `json:"source"`
	Type        string `json:"type"`
	ElementHash string `json:"element_hash"`
	Message     string `json:"message"`
}

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

type CallbackGame struct{}

type GameHighScore struct {
	Position int64 `json:"position"`
	User     *User `json:"user"`
	Score    int64 `json:"score"`
}
