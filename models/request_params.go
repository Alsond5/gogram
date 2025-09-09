package models

type GetUpdatesRequest struct {
	Offset         int64    `json:"offset,omitempty"`
	Limit          int64    `json:"limit,omitempty"`
	Timeout        int64    `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type SetWebhookRequest struct {
	Url                string     `json:"url"`
	Certificate        *InputFile `json:"certificate,omitempty"`
	IpAddress          string     `json:"ip_address,omitempty"`
	MaxConnections     int64      `json:"max_connections,omitempty"`
	AllowedUpdates     []string   `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool       `json:"drop_pending_updates,omitempty"`
	SecretToken        string     `json:"secret_token,omitempty"`
}

type DeleteWebhookRequest struct {
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

type SendMessageRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Text                    string                   `json:"text"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	Entities                []MessageEntity          `json:"entities,omitempty"`
	LinkPreviewOptions      *LinkPreviewOptions      `json:"link_preview_options,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type ForwardMessageRequest struct {
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	FromChatId              any                      `json:"from_chat_id"`
	VideoStartTimestamp     int64                    `json:"video_start_timestamp,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	MessageId               int64                    `json:"message_id"`
}

type ForwardMessagesRequest struct {
	ChatId                any     `json:"chat_id"`
	MessageThreadId       int64   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId int64   `json:"direct_messages_topic_id,omitempty"`
	FromChatId            any     `json:"from_chat_id"`
	MessageIds            []int64 `json:"message_ids"`
	DisableNotification   bool    `json:"disable_notification,omitempty"`
	ProtectContent        bool    `json:"protect_content,omitempty"`
}

type CopyMessageRequest struct {
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	FromChatId              any                      `json:"from_chat_id"`
	MessageId               int64                    `json:"message_id"`
	VideoStartTimestamp     int64                    `json:"video_start_timestamp,omitempty"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                     `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type CopyMessagesRequest struct {
	ChatId                any     `json:"chat_id"`
	MessageThreadId       int64   `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId int64   `json:"direct_messages_topic_id,omitempty"`
	FromChatId            any     `json:"from_chat_id"`
	MessageIds            []int64 `json:"message_ids"`
	DisableNotification   bool    `json:"disable_notification,omitempty"`
	ProtectContent        bool    `json:"protect_content,omitempty"`
	RemoveCaption         bool    `json:"remove_caption,omitempty"`
}

type SendPhotoRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Photo                   any                      `json:"photo"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                     `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                     `json:"has_spoiler,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendAudioRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Audio                   any                      `json:"audio"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity          `json:"caption_entities,omitempty"`
	Duration                int64                    `json:"duration,omitempty"`
	Performer               string                   `json:"performer,omitempty"`
	Title                   string                   `json:"title,omitempty"`
	Thumbnail               any                      `json:"thumbnail,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendDocumentRequest struct {
	BusinessConnectionId        string                   `json:"business_connection_id,omitempty"`
	ChatId                      any                      `json:"chat_id"`
	MessageThreadId             int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId       int64                    `json:"direct_messages_topic_id,omitempty"`
	Document                    any                      `json:"document"`
	Thumbnail                   any                      `json:"thumbnail,omitempty"`
	Caption                     string                   `json:"caption,omitempty"`
	ParseMode                   string                   `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity          `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                     `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                     `json:"disable_notification,omitempty"`
	ProtectContent              bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast          bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId             string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters     *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters             *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup                 any                      `json:"reply_markup,omitempty"`
}

type SendVideoRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Video                   any                      `json:"video"`
	Duration                int64                    `json:"duration,omitempty"`
	Width                   int64                    `json:"width,omitempty"`
	Height                  int64                    `json:"height,omitempty"`
	Thumbnail               any                      `json:"thumbnail,omitempty"`
	Cover                   any                      `json:"cover,omitempty"`
	StartTimestamp          int64                    `json:"start_timestamp,omitempty"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                     `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                     `json:"has_spoiler,omitempty"`
	SupportsStreaming       bool                     `json:"supports_streaming,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendAnimationRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Animation               any                      `json:"animation"`
	Duration                int64                    `json:"duration,omitempty"`
	Width                   int64                    `json:"width,omitempty"`
	Height                  int64                    `json:"height,omitempty"`
	Thumbnail               any                      `json:"thumbnail,omitempty"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                     `json:"show_caption_above_media,omitempty"`
	HasSpoiler              bool                     `json:"has_spoiler,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendVoiceRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Voice                   any                      `json:"voice"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity          `json:"caption_entities,omitempty"`
	Duration                int64                    `json:"duration,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendVideoNoteRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	VideoNote               any                      `json:"video_note"`
	Duration                int64                    `json:"duration,omitempty"`
	Length                  int64                    `json:"length,omitempty"`
	Thumbnail               any                      `json:"thumbnail,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendPaidMediaRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	StarCount               int64                    `json:"star_count"`
	Media                   []InputPaidMedia         `json:"media"`
	Payload                 string                   `json:"payload,omitempty"`
	Caption                 string                   `json:"caption,omitempty"`
	ParseMode               string                   `json:"parse_mode,omitempty"`
	CaptionEntities         []MessageEntity          `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia   bool                     `json:"show_caption_above_media,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendMediaGroupRequest struct {
	BusinessConnectionId  string           `json:"business_connection_id,omitempty"`
	ChatId                any              `json:"chat_id"`
	MessageThreadId       int64            `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId int64            `json:"direct_messages_topic_id,omitempty"`
	Media                 any              `json:"media"`
	DisableNotification   bool             `json:"disable_notification,omitempty"`
	ProtectContent        bool             `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool             `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId       string           `json:"message_effect_id,omitempty"`
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`
}

type SendLocationRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Latitude                float64                  `json:"latitude"`
	Longitude               float64                  `json:"longitude"`
	HorizontalAccuracy      float64                  `json:"horizontal_accuracy,omitempty"`
	LivePeriod              int64                    `json:"live_period,omitempty"`
	Heading                 int64                    `json:"heading,omitempty"`
	ProximityAlertRadius    int64                    `json:"proximity_alert_radius,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendVenueRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Latitude                float64                  `json:"latitude"`
	Longitude               float64                  `json:"longitude"`
	Title                   string                   `json:"title"`
	Address                 string                   `json:"address"`
	FoursquareId            string                   `json:"foursquare_id,omitempty"`
	FoursquareType          string                   `json:"foursquare_type,omitempty"`
	GooglePlaceId           string                   `json:"google_place_id,omitempty"`
	GooglePlaceType         string                   `json:"google_place_type,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendContactRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	PhoneNumber             string                   `json:"phone_number"`
	FirstName               string                   `json:"first_name"`
	LastName                string                   `json:"last_name,omitempty"`
	Vcard                   string                   `json:"vcard,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendPollRequest struct {
	BusinessConnectionId  string            `json:"business_connection_id,omitempty"`
	ChatId                any               `json:"chat_id"`
	MessageThreadId       int64             `json:"message_thread_id,omitempty"`
	Question              string            `json:"question"`
	QuestionParseMode     string            `json:"question_parse_mode,omitempty"`
	QuestionEntities      []MessageEntity   `json:"question_entities,omitempty"`
	Options               []InputPollOption `json:"options"`
	IsAnonymous           bool              `json:"is_anonymous,omitempty"`
	Type                  string            `json:"type,omitempty"`
	AllowsMultipleAnswers bool              `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId       int64             `json:"correct_option_id,omitempty"`
	Explanation           string            `json:"explanation,omitempty"`
	ExplanationParseMode  string            `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities   []MessageEntity   `json:"explanation_entities,omitempty"`
	OpenPeriod            int64             `json:"open_period,omitempty"`
	CloseDate             int64             `json:"close_date,omitempty"`
	IsClosed              bool              `json:"is_closed,omitempty"`
	DisableNotification   bool              `json:"disable_notification,omitempty"`
	ProtectContent        bool              `json:"protect_content,omitempty"`
	AllowPaidBroadcast    bool              `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId       string            `json:"message_effect_id,omitempty"`
	ReplyParameters       *ReplyParameters  `json:"reply_parameters,omitempty"`
	ReplyMarkup           any               `json:"reply_markup,omitempty"`
}

type SendChecklistRequest struct {
	BusinessConnectionId string                `json:"business_connection_id"`
	ChatId               int64                 `json:"chat_id"`
	Checklist            *InputChecklist       `json:"checklist"`
	DisableNotification  bool                  `json:"disable_notification,omitempty"`
	ProtectContent       bool                  `json:"protect_content,omitempty"`
	MessageEffectId      string                `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SendDiceRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Emoji                   string                   `json:"emoji,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type SendChatActionRequest struct {
	BusinessConnectionId string `json:"business_connection_id,omitempty"`
	ChatId               any    `json:"chat_id"`
	MessageThreadId      int64  `json:"message_thread_id,omitempty"`
	Action               string `json:"action"`
}

type SetMessageReactionRequest struct {
	ChatId    any            `json:"chat_id"`
	MessageId int64          `json:"message_id"`
	Reaction  []ReactionType `json:"reaction,omitempty"`
	IsBig     bool           `json:"is_big,omitempty"`
}

type GetUserProfilePhotosRequest struct {
	UserId int64 `json:"user_id"`
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

type SetUserEmojiStatusRequest struct {
	UserId                    int64  `json:"user_id"`
	EmojiStatusCustomEmojiId  string `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate int64  `json:"emoji_status_expiration_date,omitempty"`
}

type GetFileRequest struct {
	FileId string `json:"file_id"`
}

type BanChatMemberRequest struct {
	ChatId         any   `json:"chat_id"`
	UserId         int64 `json:"user_id"`
	UntilDate      int64 `json:"until_date,omitempty"`
	RevokeMessages bool  `json:"revoke_messages,omitempty"`
}

type UnbanChatMemberRequest struct {
	ChatId       any   `json:"chat_id"`
	UserId       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

type RestrictChatMemberRequest struct {
	ChatId                        any              `json:"chat_id"`
	UserId                        int64            `json:"user_id"`
	Permissions                   *ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64            `json:"until_date,omitempty"`
}

type PromoteChatMemberRequest struct {
	ChatId                  any   `json:"chat_id"`
	UserId                  int64 `json:"user_id"`
	IsAnonymous             bool  `json:"is_anonymous,omitempty"`
	CanManageChat           bool  `json:"can_manage_chat,omitempty"`
	CanDeleteMessages       bool  `json:"can_delete_messages,omitempty"`
	CanManageVideoChats     bool  `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers      bool  `json:"can_restrict_members,omitempty"`
	CanPromoteMembers       bool  `json:"can_promote_members,omitempty"`
	CanChangeInfo           bool  `json:"can_change_info,omitempty"`
	CanInviteUsers          bool  `json:"can_invite_users,omitempty"`
	CanPostStories          bool  `json:"can_post_stories,omitempty"`
	CanEditStories          bool  `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool  `json:"can_delete_stories,omitempty"`
	CanPostMessages         bool  `json:"can_post_messages,omitempty"`
	CanEditMessages         bool  `json:"can_edit_messages,omitempty"`
	CanPinMessages          bool  `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool  `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages bool  `json:"can_manage_direct_messages,omitempty"`
}

type SetChatAdministratorCustomTitleRequest struct {
	ChatId      any    `json:"chat_id"`
	UserId      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

type BanChatSenderChatRequest struct {
	ChatId       any   `json:"chat_id"`
	SenderChatId int64 `json:"sender_chat_id"`
}

type UnbanChatSenderChatRequest struct {
	ChatId       any   `json:"chat_id"`
	SenderChatId int64 `json:"sender_chat_id"`
}

type SetChatPermissionsRequest struct {
	ChatId                        any              `json:"chat_id"`
	Permissions                   *ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool             `json:"use_independent_chat_permissions,omitempty"`
}

type ExportChatInviteLinkRequest struct {
	ChatId any `json:"chat_id"`
}

type CreateChatInviteLinkRequest struct {
	ChatId             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type EditChatInviteLinkRequest struct {
	ChatId             any    `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int64  `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type CreateChatSubscriptionInviteLinkRequest struct {
	ChatId             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	SubscriptionPeriod int64  `json:"subscription_period"`
	SubscriptionPrice  int64  `json:"subscription_price"`
}

type EditChatSubscriptionInviteLinkRequest struct {
	ChatId     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
	Name       string `json:"name,omitempty"`
}

type RevokeChatInviteLinkRequest struct {
	ChatId     any    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

type ApproveChatJoinRequestRequest struct {
	ChatId any   `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

type DeclineChatJoinRequestRequest struct {
	ChatId any   `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

type SetChatPhotoRequest struct {
	ChatId any        `json:"chat_id"`
	Photo  *InputFile `json:"photo"`
}

type DeleteChatPhotoRequest struct {
	ChatId any `json:"chat_id"`
}

type SetChatTitleRequest struct {
	ChatId any    `json:"chat_id"`
	Title  string `json:"title"`
}

type SetChatDescriptionRequest struct {
	ChatId      any    `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

type PinChatMessageRequest struct {
	BusinessConnectionId string `json:"business_connection_id,omitempty"`
	ChatId               any    `json:"chat_id"`
	MessageId            int64  `json:"message_id"`
	DisableNotification  bool   `json:"disable_notification,omitempty"`
}

type UnpinChatMessageRequest struct {
	BusinessConnectionId string `json:"business_connection_id,omitempty"`
	ChatId               any    `json:"chat_id"`
	MessageId            int64  `json:"message_id,omitempty"`
}

type UnpinAllChatMessagesRequest struct {
	ChatId any `json:"chat_id"`
}

type LeaveChatRequest struct {
	ChatId any `json:"chat_id"`
}

type GetChatRequest struct {
	ChatId any `json:"chat_id"`
}

type GetChatAdministratorsRequest struct {
	ChatId any `json:"chat_id"`
}

type GetChatMemberCountRequest struct {
	ChatId any `json:"chat_id"`
}

type GetChatMemberRequest struct {
	ChatId any   `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

type SetChatStickerSetRequest struct {
	ChatId         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type DeleteChatStickerSetRequest struct {
	ChatId any `json:"chat_id"`
}

type CreateForumTopicRequest struct {
	ChatId            any    `json:"chat_id"`
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type EditForumTopicRequest struct {
	ChatId            any    `json:"chat_id"`
	MessageThreadId   int64  `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}

type CloseForumTopicRequest struct {
	ChatId          any   `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

type ReopenForumTopicRequest struct {
	ChatId          any   `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

type DeleteForumTopicRequest struct {
	ChatId          any   `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

type UnpinAllForumTopicMessagesRequest struct {
	ChatId          any   `json:"chat_id"`
	MessageThreadId int64 `json:"message_thread_id"`
}

type EditGeneralForumTopicRequest struct {
	ChatId any    `json:"chat_id"`
	Name   string `json:"name"`
}

type CloseGeneralForumTopicRequest struct {
	ChatId any `json:"chat_id"`
}

type ReopenGeneralForumTopicRequest struct {
	ChatId any `json:"chat_id"`
}

type HideGeneralForumTopicRequest struct {
	ChatId any `json:"chat_id"`
}

type UnhideGeneralForumTopicRequest struct {
	ChatId any `json:"chat_id"`
}

type UnpinAllGeneralForumTopicMessagesRequest struct {
	ChatId any `json:"chat_id"`
}

type AnswerCallbackQueryRequest struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       int64  `json:"cache_time,omitempty"`
}

type GetUserChatBoostsRequest struct {
	ChatId any   `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

type GetBusinessConnectionRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
}

type SetMyCommandsRequest struct {
	Commands     []BotCommand     `json:"commands"`
	Scope        *BotCommandScope `json:"scope,omitempty"`
	LanguageCode string           `json:"language_code,omitempty"`
}

type DeleteMyCommandsRequest struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`
	LanguageCode string           `json:"language_code,omitempty"`
}

type GetMyCommandsRequest struct {
	Scope        *BotCommandScope `json:"scope,omitempty"`
	LanguageCode string           `json:"language_code,omitempty"`
}

type SetMyNameRequest struct {
	Name         string `json:"name,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyNameRequest struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyDescriptionRequest struct {
	Description  string `json:"description,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type GetMyDescriptionRequest struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetMyShortDescriptionRequest struct {
	ShortDescription string `json:"short_description,omitempty"`
	LanguageCode     string `json:"language_code,omitempty"`
}

type GetMyShortDescriptionRequest struct {
	LanguageCode string `json:"language_code,omitempty"`
}

type SetChatMenuButtonRequest struct {
	ChatId     int64       `json:"chat_id,omitempty"`
	MenuButton *MenuButton `json:"menu_button,omitempty"`
}

type GetChatMenuButtonRequest struct {
	ChatId int64 `json:"chat_id,omitempty"`
}

type SetMyDefaultAdministratorRightsRequest struct {
	Rights      *ChatAdministratorRights `json:"rights,omitempty"`
	ForChannels bool                     `json:"for_channels,omitempty"`
}

type GetMyDefaultAdministratorRightsRequest struct {
	ForChannels bool `json:"for_channels,omitempty"`
}

type SendGiftRequest struct {
	UserId        int64           `json:"user_id,omitempty"`
	ChatId        any             `json:"chat_id,omitempty"`
	GiftId        string          `json:"gift_id"`
	PayForUpgrade bool            `json:"pay_for_upgrade,omitempty"`
	Text          string          `json:"text,omitempty"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

type GiftPremiumSubscriptionRequest struct {
	UserId        int64           `json:"user_id"`
	MonthCount    int64           `json:"month_count"`
	StarCount     int64           `json:"star_count"`
	Text          string          `json:"text,omitempty"`
	TextParseMode string          `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

type VerifyUserRequest struct {
	UserId            int64  `json:"user_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

type VerifyChatRequest struct {
	ChatId            any    `json:"chat_id"`
	CustomDescription string `json:"custom_description,omitempty"`
}

type RemoveUserVerificationRequest struct {
	UserId int64 `json:"user_id"`
}

type RemoveChatVerificationRequest struct {
	ChatId any `json:"chat_id"`
}

type ReadBusinessMessageRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	ChatId               int64  `json:"chat_id"`
	MessageId            int64  `json:"message_id"`
}

type DeleteBusinessMessagesRequest struct {
	BusinessConnectionId string  `json:"business_connection_id"`
	MessageIds           []int64 `json:"message_ids"`
}

type SetBusinessAccountNameRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name,omitempty"`
}

type SetBusinessAccountUsernameRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	Username             string `json:"username,omitempty"`
}

type SetBusinessAccountBioRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	Bio                  string `json:"bio,omitempty"`
}

type SetBusinessAccountProfilePhotoRequest struct {
	BusinessConnectionId string             `json:"business_connection_id"`
	Photo                *InputProfilePhoto `json:"photo"`
	IsPublic             bool               `json:"is_public,omitempty"`
}

type RemoveBusinessAccountProfilePhotoRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	IsPublic             bool   `json:"is_public,omitempty"`
}

type SetBusinessAccountGiftSettingsRequest struct {
	BusinessConnectionId string             `json:"business_connection_id"`
	ShowGiftButton       bool               `json:"show_gift_button"`
	AcceptedGiftTypes    *AcceptedGiftTypes `json:"accepted_gift_types"`
}

type GetBusinessAccountStarBalanceRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
}

type TransferBusinessAccountStarsRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	StarCount            int64  `json:"star_count"`
}

type GetBusinessAccountGiftsRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	ExcludeUnsaved       bool   `json:"exclude_unsaved,omitempty"`
	ExcludeSaved         bool   `json:"exclude_saved,omitempty"`
	ExcludeUnlimited     bool   `json:"exclude_unlimited,omitempty"`
	ExcludeLimited       bool   `json:"exclude_limited,omitempty"`
	ExcludeUnique        bool   `json:"exclude_unique,omitempty"`
	SortByPrice          bool   `json:"sort_by_price,omitempty"`
	Offset               string `json:"offset,omitempty"`
	Limit                int64  `json:"limit,omitempty"`
}

type ConvertGiftToStarsRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	OwnedGiftId          string `json:"owned_gift_id"`
}

type UpgradeGiftRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	OwnedGiftId          string `json:"owned_gift_id"`
	KeepOriginalDetails  bool   `json:"keep_original_details,omitempty"`
	StarCount            int64  `json:"star_count,omitempty"`
}

type TransferGiftRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	OwnedGiftId          string `json:"owned_gift_id"`
	NewOwnerChatId       int64  `json:"new_owner_chat_id"`
	StarCount            int64  `json:"star_count,omitempty"`
}

type PostStoryRequest struct {
	BusinessConnectionId string             `json:"business_connection_id"`
	Content              *InputStoryContent `json:"content"`
	ActivePeriod         int64              `json:"active_period"`
	Caption              string             `json:"caption,omitempty"`
	ParseMode            string             `json:"parse_mode,omitempty"`
	CaptionEntities      []MessageEntity    `json:"caption_entities,omitempty"`
	Areas                []StoryArea        `json:"areas,omitempty"`
	PostToChatPage       bool               `json:"post_to_chat_page,omitempty"`
	ProtectContent       bool               `json:"protect_content,omitempty"`
}

type EditStoryRequest struct {
	BusinessConnectionId string             `json:"business_connection_id"`
	StoryId              int64              `json:"story_id"`
	Content              *InputStoryContent `json:"content"`
	Caption              string             `json:"caption,omitempty"`
	ParseMode            string             `json:"parse_mode,omitempty"`
	CaptionEntities      []MessageEntity    `json:"caption_entities,omitempty"`
	Areas                []StoryArea        `json:"areas,omitempty"`
}

type DeleteStoryRequest struct {
	BusinessConnectionId string `json:"business_connection_id"`
	StoryId              int64  `json:"story_id"`
}

type EditMessageTextRequest struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"`
	ChatId               any                   `json:"chat_id,omitempty"`
	MessageId            int64                 `json:"message_id,omitempty"`
	InlineMessageId      string                `json:"inline_message_id,omitempty"`
	Text                 string                `json:"text"`
	ParseMode            string                `json:"parse_mode,omitempty"`
	Entities             []MessageEntity       `json:"entities,omitempty"`
	LinkPreviewOptions   *LinkPreviewOptions   `json:"link_preview_options,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageCaptionRequest struct {
	BusinessConnectionId  string                `json:"business_connection_id,omitempty"`
	ChatId                any                   `json:"chat_id,omitempty"`
	MessageId             int64                 `json:"message_id,omitempty"`
	InlineMessageId       string                `json:"inline_message_id,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageMediaRequest struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"`
	ChatId               any                   `json:"chat_id,omitempty"`
	MessageId            int64                 `json:"message_id,omitempty"`
	InlineMessageId      string                `json:"inline_message_id,omitempty"`
	Media                *InputMedia           `json:"media"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageLiveLocationRequest struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"`
	ChatId               any                   `json:"chat_id,omitempty"`
	MessageId            int64                 `json:"message_id,omitempty"`
	InlineMessageId      string                `json:"inline_message_id,omitempty"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	LivePeriod           int64                 `json:"live_period,omitempty"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	Heading              int64                 `json:"heading,omitempty"`
	ProximityAlertRadius int64                 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type StopMessageLiveLocationRequest struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"`
	ChatId               any                   `json:"chat_id,omitempty"`
	MessageId            int64                 `json:"message_id,omitempty"`
	InlineMessageId      string                `json:"inline_message_id,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageChecklistRequest struct {
	BusinessConnectionId string                `json:"business_connection_id"`
	ChatId               int64                 `json:"chat_id"`
	MessageId            int64                 `json:"message_id"`
	Checklist            *InputChecklist       `json:"checklist"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type EditMessageReplyMarkupRequest struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"`
	ChatId               any                   `json:"chat_id,omitempty"`
	MessageId            int64                 `json:"message_id,omitempty"`
	InlineMessageId      string                `json:"inline_message_id,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type StopPollRequest struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"`
	ChatId               any                   `json:"chat_id"`
	MessageId            int64                 `json:"message_id"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type ApproveSuggestedPostRequest struct {
	ChatId    int64 `json:"chat_id"`
	MessageId int64 `json:"message_id"`
	SendDate  int64 `json:"send_date,omitempty"`
}

type DeclineSuggestedPostRequest struct {
	ChatId    int64  `json:"chat_id"`
	MessageId int64  `json:"message_id"`
	Comment   string `json:"comment,omitempty"`
}

type DeleteMessageRequest struct {
	ChatId    any   `json:"chat_id"`
	MessageId int64 `json:"message_id"`
}

type DeleteMessagesRequest struct {
	ChatId     any     `json:"chat_id"`
	MessageIds []int64 `json:"message_ids"`
}

type SendStickerRequest struct {
	BusinessConnectionId    string                   `json:"business_connection_id,omitempty"`
	ChatId                  any                      `json:"chat_id"`
	MessageThreadId         int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId   int64                    `json:"direct_messages_topic_id,omitempty"`
	Sticker                 any                      `json:"sticker"`
	Emoji                   string                   `json:"emoji,omitempty"`
	DisableNotification     bool                     `json:"disable_notification,omitempty"`
	ProtectContent          bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast      bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId         string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters         *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup             any                      `json:"reply_markup,omitempty"`
}

type GetStickerSetRequest struct {
	Name string `json:"name"`
}

type GetCustomEmojiStickersRequest struct {
	CustomEmojiIds []string `json:"custom_emoji_ids"`
}

type UploadStickerFileRequest struct {
	UserId        int64      `json:"user_id"`
	Sticker       *InputFile `json:"sticker"`
	StickerFormat string     `json:"sticker_format"`
}

type CreateNewStickerSetRequest struct {
	UserId          int64          `json:"user_id"`
	Name            string         `json:"name"`
	Title           string         `json:"title"`
	Stickers        []InputSticker `json:"stickers"`
	StickerType     string         `json:"sticker_type,omitempty"`
	NeedsRepainting bool           `json:"needs_repainting,omitempty"`
}

type AddStickerToSetRequest struct {
	UserId  int64         `json:"user_id"`
	Name    string        `json:"name"`
	Sticker *InputSticker `json:"sticker"`
}

type SetStickerPositionInSetRequest struct {
	Sticker  string `json:"sticker"`
	Position int64  `json:"position"`
}

type DeleteStickerFromSetRequest struct {
	Sticker string `json:"sticker"`
}

type ReplaceStickerInSetRequest struct {
	UserId     int64         `json:"user_id"`
	Name       string        `json:"name"`
	OldSticker string        `json:"old_sticker"`
	Sticker    *InputSticker `json:"sticker"`
}

type SetStickerEmojiListRequest struct {
	Sticker   string   `json:"sticker"`
	EmojiList []string `json:"emoji_list"`
}

type SetStickerKeywordsRequest struct {
	Sticker  string   `json:"sticker"`
	Keywords []string `json:"keywords,omitempty"`
}

type SetStickerMaskPositionRequest struct {
	Sticker      string        `json:"sticker"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

type SetStickerSetTitleRequest struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type SetStickerSetThumbnailRequest struct {
	Name      string `json:"name"`
	UserId    int64  `json:"user_id"`
	Thumbnail any    `json:"thumbnail,omitempty"`
	Format    string `json:"format"`
}

type SetCustomEmojiStickerSetThumbnailRequest struct {
	Name          string `json:"name"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

type DeleteStickerSetRequest struct {
	Name string `json:"name"`
}

type AnswerInlineQueryRequest struct {
	InlineQueryId string                    `json:"inline_query_id"`
	Results       []InlineQueryResult       `json:"results"`
	CacheTime     int64                     `json:"cache_time,omitempty"`
	IsPersonal    bool                      `json:"is_personal,omitempty"`
	NextOffset    string                    `json:"next_offset,omitempty"`
	Button        *InlineQueryResultsButton `json:"button,omitempty"`
}

type AnswerWebAppQueryRequest struct {
	WebAppQueryId string             `json:"web_app_query_id"`
	Result        *InlineQueryResult `json:"result"`
}

type SavePreparedInlineMessageRequest struct {
	UserId            int64              `json:"user_id"`
	Result            *InlineQueryResult `json:"result"`
	AllowUserChats    bool               `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool               `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool               `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool               `json:"allow_channel_chats,omitempty"`
}

type SendInvoiceRequest struct {
	ChatId                    any                      `json:"chat_id"`
	MessageThreadId           int64                    `json:"message_thread_id,omitempty"`
	DirectMessagesTopicId     int64                    `json:"direct_messages_topic_id,omitempty"`
	Title                     string                   `json:"title"`
	Description               string                   `json:"description"`
	Payload                   string                   `json:"payload"`
	ProviderToken             string                   `json:"provider_token,omitempty"`
	Currency                  string                   `json:"currency"`
	Prices                    []LabeledPrice           `json:"prices"`
	MaxTipAmount              int64                    `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int64                  `json:"suggested_tip_amounts,omitempty"`
	StartParameter            string                   `json:"start_parameter,omitempty"`
	ProviderData              string                   `json:"provider_data,omitempty"`
	PhotoUrl                  string                   `json:"photo_url,omitempty"`
	PhotoSize                 int64                    `json:"photo_size,omitempty"`
	PhotoWidth                int64                    `json:"photo_width,omitempty"`
	PhotoHeight               int64                    `json:"photo_height,omitempty"`
	NeedName                  bool                     `json:"need_name,omitempty"`
	NeedPhoneNumber           bool                     `json:"need_phone_number,omitempty"`
	NeedEmail                 bool                     `json:"need_email,omitempty"`
	NeedShippingAddress       bool                     `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool                     `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool                     `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool                     `json:"is_flexible,omitempty"`
	DisableNotification       bool                     `json:"disable_notification,omitempty"`
	ProtectContent            bool                     `json:"protect_content,omitempty"`
	AllowPaidBroadcast        bool                     `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId           string                   `json:"message_effect_id,omitempty"`
	SuggestedPostParameters   *SuggestedPostParameters `json:"suggested_post_parameters,omitempty"`
	ReplyParameters           *ReplyParameters         `json:"reply_parameters,omitempty"`
	ReplyMarkup               *InlineKeyboardMarkup    `json:"reply_markup,omitempty"`
}

type CreateInvoiceLinkRequest struct {
	BusinessConnectionId      string         `json:"business_connection_id,omitempty"`
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	SubscriptionPeriod        int64          `json:"subscription_period,omitempty"`
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

type AnswerShippingQueryRequest struct {
	ShippingQueryId string           `json:"shipping_query_id"`
	Ok              bool             `json:"ok"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
}

type AnswerPreCheckoutQueryRequest struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

type GetStarTransactionsRequest struct {
	Offset int64 `json:"offset,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
}

type RefundStarPaymentRequest struct {
	UserId                  int64  `json:"user_id"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
}

type EditUserStarSubscriptionRequest struct {
	UserId                  int64  `json:"user_id"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
}

type SetPassportDataErrorsRequest struct {
	UserId int64                  `json:"user_id"`
	Errors []PassportElementError `json:"errors"`
}

type SendGameRequest struct {
	BusinessConnectionId string                `json:"business_connection_id,omitempty"`
	ChatId               int64                 `json:"chat_id"`
	MessageThreadId      int64                 `json:"message_thread_id,omitempty"`
	GameShortName        string                `json:"game_short_name"`
	DisableNotification  bool                  `json:"disable_notification,omitempty"`
	ProtectContent       bool                  `json:"protect_content,omitempty"`
	AllowPaidBroadcast   bool                  `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId      string                `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters      `json:"reply_parameters,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type SetGameScoreRequest struct {
	UserId             int64  `json:"user_id"`
	Score              int64  `json:"score"`
	Force              bool   `json:"force,omitempty"`
	DisableEditMessage bool   `json:"disable_edit_message,omitempty"`
	ChatId             int64  `json:"chat_id,omitempty"`
	MessageId          int64  `json:"message_id,omitempty"`
	InlineMessageId    string `json:"inline_message_id,omitempty"`
}

type GetGameHighScoresRequest struct {
	UserId          int64  `json:"user_id"`
	ChatId          int64  `json:"chat_id,omitempty"`
	MessageId       int64  `json:"message_id,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
}
