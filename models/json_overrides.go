package models

func (mim *MaybeInaccessibleMessage) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[MaybeInaccessibleMessageType]{
		{
			TypeField:   &mim.Type,
			TypeValue:   MaybeInaccessibleMessageTypeMessage,
			TargetField: &mim.Message,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields(""),
			),
		},
		{
			TypeField:   &mim.Type,
			TypeValue:   MaybeInaccessibleMessageTypeInaccessibleMessage,
			TargetField: &mim.InaccessibleMessage,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("from", "business_connection_id", "forward_origin", "sender_chat", "sender_boost_count", "sender_business_bot", "message_thread_id", "direct_messages_topic"),
			),
		},
	})
}

func (mim *MaybeInaccessibleMessage) MarshalJSON() ([]byte, error) {
	return MarshalUnion(mim.Type, map[MaybeInaccessibleMessageType]any{
		MaybeInaccessibleMessageTypeMessage:             mim.Message,
		MaybeInaccessibleMessageTypeInaccessibleMessage: mim.InaccessibleMessage,
	})
}

func (mo *MessageOrigin) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[MessageOriginType]{
		{
			TypeField:   &mo.Type,
			TypeValue:   MessageOriginTypeMessageOriginUser,
			TargetField: &mo.MessageOriginUser,
			Rules: CombineRules(
				RequiredFields("sender_user"),
				ForbiddenFields("chat", "author_signature", "sender_chat", "sender_user_name", "message_id"),
			),
		},
		{
			TypeField:   &mo.Type,
			TypeValue:   MessageOriginTypeMessageOriginHiddenUser,
			TargetField: &mo.MessageOriginHiddenUser,
			Rules: CombineRules(
				RequiredFields("sender_user_name"),
				ForbiddenFields("chat", "sender_user", "author_signature", "sender_chat", "message_id"),
			),
		},
		{
			TypeField:   &mo.Type,
			TypeValue:   MessageOriginTypeMessageOriginChat,
			TargetField: &mo.MessageOriginChat,
			Rules: CombineRules(
				RequiredFields("sender_chat"),
				ForbiddenFields("sender_user", "chat", "sender_user_name", "message_id"),
			),
		},
		{
			TypeField:   &mo.Type,
			TypeValue:   MessageOriginTypeMessageOriginChannel,
			TargetField: &mo.MessageOriginChannel,
			Rules: CombineRules(
				RequiredFields("chat", "message_id"),
				ForbiddenFields("sender_user", "sender_user_name", "sender_chat"),
			),
		},
	})
}

func (mo *MessageOrigin) MarshalJSON() ([]byte, error) {
	return MarshalUnion(mo.Type, map[MessageOriginType]any{
		MessageOriginTypeMessageOriginUser:       mo.MessageOriginUser,
		MessageOriginTypeMessageOriginHiddenUser: mo.MessageOriginHiddenUser,
		MessageOriginTypeMessageOriginChat:       mo.MessageOriginChat,
		MessageOriginTypeMessageOriginChannel:    mo.MessageOriginChannel,
	})
}

func (pm *PaidMedia) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[PaidMediaType]{
		{
			TypeField:   &pm.Type,
			TypeValue:   PaidMediaTypePaidMediaPreview,
			TargetField: &pm.PaidMediaPreview,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("video", "photo"),
			),
		},
		{
			TypeField:   &pm.Type,
			TypeValue:   PaidMediaTypePaidMediaPhoto,
			TargetField: &pm.PaidMediaPhoto,
			Rules: CombineRules(
				RequiredFields("photo"),
				ForbiddenFields("video", "duration", "height", "width"),
			),
		},
		{
			TypeField:   &pm.Type,
			TypeValue:   PaidMediaTypePaidMediaVideo,
			TargetField: &pm.PaidMediaVideo,
			Rules: CombineRules(
				RequiredFields("video"),
				ForbiddenFields("duration", "height", "width", "photo"),
			),
		},
	})
}

func (pm *PaidMedia) MarshalJSON() ([]byte, error) {
	return MarshalUnion(pm.Type, map[PaidMediaType]any{
		PaidMediaTypePaidMediaPreview: pm.PaidMediaPreview,
		PaidMediaTypePaidMediaPhoto:   pm.PaidMediaPhoto,
		PaidMediaTypePaidMediaVideo:   pm.PaidMediaVideo,
	})
}

func (bf *BackgroundFill) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[BackgroundFillType]{
		{
			TypeField:   &bf.Type,
			TypeValue:   BackgroundFillTypeBackgroundFillSolid,
			TargetField: &bf.BackgroundFillSolid,
			Rules: CombineRules(
				RequiredFields("color"),
				ForbiddenFields("rotation_angle", "top_color", "bottom_color", "colors"),
			),
		},
		{
			TypeField:   &bf.Type,
			TypeValue:   BackgroundFillTypeBackgroundFillGradient,
			TargetField: &bf.BackgroundFillGradient,
			Rules: CombineRules(
				RequiredFields("rotation_angle", "top_color", "bottom_color"),
				ForbiddenFields("colors", "color"),
			),
		},
		{
			TypeField:   &bf.Type,
			TypeValue:   BackgroundFillTypeBackgroundFillFreeformGradient,
			TargetField: &bf.BackgroundFillFreeformGradient,
			Rules: CombineRules(
				RequiredFields("colors"),
				ForbiddenFields("rotation_angle", "top_color", "bottom_color", "color"),
			),
		},
	})
}

func (bf *BackgroundFill) MarshalJSON() ([]byte, error) {
	return MarshalUnion(bf.Type, map[BackgroundFillType]any{
		BackgroundFillTypeBackgroundFillSolid:            bf.BackgroundFillSolid,
		BackgroundFillTypeBackgroundFillGradient:         bf.BackgroundFillGradient,
		BackgroundFillTypeBackgroundFillFreeformGradient: bf.BackgroundFillFreeformGradient,
	})
}

func (bt *BackgroundType) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[BackgroundTypeType]{
		{
			TypeField:   &bt.Type,
			TypeValue:   BackgroundTypeTypeBackgroundTypeFill,
			TargetField: &bt.BackgroundTypeFill,
			Rules: CombineRules(
				RequiredFields("fill", "dark_theme_dimming"),
				ForbiddenFields("is_moving", "theme_name", "is_inverted", "document", "intensity", "is_blurred"),
			),
		},
		{
			TypeField:   &bt.Type,
			TypeValue:   BackgroundTypeTypeBackgroundTypeWallpaper,
			TargetField: &bt.BackgroundTypeWallpaper,
			Rules: CombineRules(
				RequiredFields("dark_theme_dimming", "document"),
				ForbiddenFields("theme_name", "fill", "is_inverted", "intensity"),
			),
		},
		{
			TypeField:   &bt.Type,
			TypeValue:   BackgroundTypeTypeBackgroundTypePattern,
			TargetField: &bt.BackgroundTypePattern,
			Rules: CombineRules(
				RequiredFields("fill", "document", "intensity"),
				ForbiddenFields("theme_name", "dark_theme_dimming", "is_blurred"),
			),
		},
		{
			TypeField:   &bt.Type,
			TypeValue:   BackgroundTypeTypeBackgroundTypeChatTheme,
			TargetField: &bt.BackgroundTypeChatTheme,
			Rules: CombineRules(
				RequiredFields("theme_name"),
				ForbiddenFields("is_moving", "fill", "document", "dark_theme_dimming", "is_blurred"),
			),
		},
	})
}

func (bt *BackgroundType) MarshalJSON() ([]byte, error) {
	return MarshalUnion(bt.Type, map[BackgroundTypeType]any{
		BackgroundTypeTypeBackgroundTypeFill:      bt.BackgroundTypeFill,
		BackgroundTypeTypeBackgroundTypeWallpaper: bt.BackgroundTypeWallpaper,
		BackgroundTypeTypeBackgroundTypePattern:   bt.BackgroundTypePattern,
		BackgroundTypeTypeBackgroundTypeChatTheme: bt.BackgroundTypeChatTheme,
	})
}

func (cm *ChatMember) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[ChatMemberType]{
		{
			TypeField:   &cm.Type,
			TypeValue:   ChatMemberTypeChatMemberOwner,
			TargetField: &cm.ChatMemberOwner,
			Rules: CombineRules(
				RequiredFields("is_anonymous"),
				ForbiddenFields("can_restrict_members", "can_be_edited", "can_manage_chat", "can_delete_messages", "can_manage_video_chats", "can_invite_users", "can_promote_members", "can_change_info"),
			),
		},
		{
			TypeField:   &cm.Type,
			TypeValue:   ChatMemberTypeChatMemberAdministrator,
			TargetField: &cm.ChatMemberAdministrator,
			Rules: CombineRules(
				RequiredFields("can_restrict_members", "can_edit_stories", "can_be_edited", "can_manage_chat", "can_manage_video_chats", "can_delete_messages", "is_anonymous", "can_post_stories", "can_invite_users", "can_promote_members", "can_delete_stories", "can_change_info"),
				ForbiddenFields("until_date", "can_send_photos", "is_member", "can_send_videos", "can_send_video_notes", "can_send_messages", "can_send_audios", "can_send_documents"),
			),
		},
		{
			TypeField:   &cm.Type,
			TypeValue:   ChatMemberTypeChatMemberMember,
			TargetField: &cm.ChatMemberMember,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("can_restrict_members", "can_be_edited", "can_manage_chat", "can_delete_messages", "can_manage_video_chats", "custom_title", "is_anonymous"),
			),
		},
		{
			TypeField:   &cm.Type,
			TypeValue:   ChatMemberTypeChatMemberRestricted,
			TargetField: &cm.ChatMemberRestricted,
			Rules: CombineRules(
				RequiredFields("until_date", "can_send_voice_notes", "can_send_photos", "is_member", "can_send_videos", "can_send_video_notes", "can_invite_users", "can_manage_topics", "can_send_messages", "can_send_polls", "can_send_audios", "can_send_documents", "can_pin_messages", "can_send_other_messages", "can_change_info", "can_add_web_page_previews"),
				ForbiddenFields("can_restrict_members", "can_be_edited", "can_manage_chat", "can_delete_messages", "can_manage_video_chats", "custom_title", "is_anonymous"),
			),
		},
		{
			TypeField:   &cm.Type,
			TypeValue:   ChatMemberTypeChatMemberLeft,
			TargetField: &cm.ChatMemberLeft,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("can_restrict_members", "can_be_edited", "can_manage_chat", "can_delete_messages", "can_manage_video_chats", "custom_title", "is_anonymous"),
			),
		},
		{
			TypeField:   &cm.Type,
			TypeValue:   ChatMemberTypeChatMemberBanned,
			TargetField: &cm.ChatMemberBanned,
			Rules: CombineRules(
				RequiredFields("until_date"),
				ForbiddenFields("can_restrict_members", "can_be_edited", "can_manage_chat", "can_delete_messages", "can_manage_video_chats", "custom_title", "is_anonymous"),
			),
		},
	})
}

func (cm *ChatMember) MarshalJSON() ([]byte, error) {
	return MarshalUnion(cm.Type, map[ChatMemberType]any{
		ChatMemberTypeChatMemberOwner:         cm.ChatMemberOwner,
		ChatMemberTypeChatMemberAdministrator: cm.ChatMemberAdministrator,
		ChatMemberTypeChatMemberMember:        cm.ChatMemberMember,
		ChatMemberTypeChatMemberRestricted:    cm.ChatMemberRestricted,
		ChatMemberTypeChatMemberLeft:          cm.ChatMemberLeft,
		ChatMemberTypeChatMemberBanned:        cm.ChatMemberBanned,
	})
}

func (sat *StoryAreaType) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[StoryAreaTypeType]{
		{
			TypeField:   &sat.Type,
			TypeValue:   StoryAreaTypeTypeStoryAreaTypeLocation,
			TargetField: &sat.StoryAreaTypeLocation,
			Rules: CombineRules(
				RequiredFields("latitude", "longitude"),
				ForbiddenFields("is_dark", "background_color", "name", "reaction_type", "temperature", "emoji", "url", "is_flipped"),
			),
		},
		{
			TypeField:   &sat.Type,
			TypeValue:   StoryAreaTypeTypeStoryAreaTypeSuggestedReaction,
			TargetField: &sat.StoryAreaTypeSuggestedReaction,
			Rules: CombineRules(
				RequiredFields("reaction_type"),
				ForbiddenFields("longitude", "background_color", "name", "latitude", "address", "temperature", "emoji", "url"),
			),
		},
		{
			TypeField:   &sat.Type,
			TypeValue:   StoryAreaTypeTypeStoryAreaTypeLink,
			TargetField: &sat.StoryAreaTypeLink,
			Rules: CombineRules(
				RequiredFields("url"),
				ForbiddenFields("is_dark", "longitude", "latitude", "address", "reaction_type", "temperature", "emoji", "is_flipped"),
			),
		},
		{
			TypeField:   &sat.Type,
			TypeValue:   StoryAreaTypeTypeStoryAreaTypeWeather,
			TargetField: &sat.StoryAreaTypeWeather,
			Rules: CombineRules(
				RequiredFields("background_color", "temperature", "emoji"),
				ForbiddenFields("is_dark", "longitude", "name", "latitude", "address", "reaction_type", "url", "is_flipped"),
			),
		},
		{
			TypeField:   &sat.Type,
			TypeValue:   StoryAreaTypeTypeStoryAreaTypeUniqueGift,
			TargetField: &sat.StoryAreaTypeUniqueGift,
			Rules: CombineRules(
				RequiredFields("name"),
				ForbiddenFields("is_dark", "longitude", "latitude", "address", "reaction_type", "temperature", "url", "is_flipped"),
			),
		},
	})
}

func (sat *StoryAreaType) MarshalJSON() ([]byte, error) {
	return MarshalUnion(sat.Type, map[StoryAreaTypeType]any{
		StoryAreaTypeTypeStoryAreaTypeLocation:          sat.StoryAreaTypeLocation,
		StoryAreaTypeTypeStoryAreaTypeSuggestedReaction: sat.StoryAreaTypeSuggestedReaction,
		StoryAreaTypeTypeStoryAreaTypeLink:              sat.StoryAreaTypeLink,
		StoryAreaTypeTypeStoryAreaTypeWeather:           sat.StoryAreaTypeWeather,
		StoryAreaTypeTypeStoryAreaTypeUniqueGift:        sat.StoryAreaTypeUniqueGift,
	})
}

func (rt *ReactionType) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[ReactionTypeType]{
		{
			TypeField:   &rt.Type,
			TypeValue:   ReactionTypeTypeReactionTypeEmoji,
			TargetField: &rt.ReactionTypeEmoji,
			Rules: CombineRules(
				RequiredFields("emoji"),
				ForbiddenFields("custom_emoji_id"),
			),
		},
		{
			TypeField:   &rt.Type,
			TypeValue:   ReactionTypeTypeReactionTypeCustomEmoji,
			TargetField: &rt.ReactionTypeCustomEmoji,
			Rules: CombineRules(
				RequiredFields("custom_emoji_id"),
				ForbiddenFields("emoji"),
			),
		},
		{
			TypeField:   &rt.Type,
			TypeValue:   ReactionTypeTypeReactionTypePaid,
			TargetField: &rt.ReactionTypePaid,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("custom_emoji_id", "emoji"),
			),
		},
	})
}

func (rt *ReactionType) MarshalJSON() ([]byte, error) {
	return MarshalUnion(rt.Type, map[ReactionTypeType]any{
		ReactionTypeTypeReactionTypeEmoji:       rt.ReactionTypeEmoji,
		ReactionTypeTypeReactionTypeCustomEmoji: rt.ReactionTypeCustomEmoji,
		ReactionTypeTypeReactionTypePaid:        rt.ReactionTypePaid,
	})
}

func (og *OwnedGift) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[OwnedGiftType]{
		{
			TypeField:   &og.Type,
			TypeValue:   OwnedGiftTypeOwnedGiftRegular,
			TargetField: &og.OwnedGiftRegular,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("transfer_star_count", "can_be_transferred", "next_transfer_date"),
			),
		},
		{
			TypeField:   &og.Type,
			TypeValue:   OwnedGiftTypeOwnedGiftUnique,
			TargetField: &og.OwnedGiftUnique,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("prepaid_upgrade_star_count", "was_refunded", "is_private", "can_be_upgraded", "convert_star_count", "entities", "text"),
			),
		},
	})
}

func (og *OwnedGift) MarshalJSON() ([]byte, error) {
	return MarshalUnion(og.Type, map[OwnedGiftType]any{
		OwnedGiftTypeOwnedGiftRegular: og.OwnedGiftRegular,
		OwnedGiftTypeOwnedGiftUnique:  og.OwnedGiftUnique,
	})
}

func (bcs *BotCommandScope) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[BotCommandScopeType]{
		{
			TypeField:   &bcs.Type,
			TypeValue:   BotCommandScopeTypeBotCommandScopeDefault,
			TargetField: &bcs.BotCommandScopeDefault,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("user_id", "chat_id"),
			),
		},
		{
			TypeField:   &bcs.Type,
			TypeValue:   BotCommandScopeTypeBotCommandScopeAllPrivateChats,
			TargetField: &bcs.BotCommandScopeAllPrivateChats,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("user_id", "chat_id"),
			),
		},
		{
			TypeField:   &bcs.Type,
			TypeValue:   BotCommandScopeTypeBotCommandScopeAllGroupChats,
			TargetField: &bcs.BotCommandScopeAllGroupChats,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("user_id", "chat_id"),
			),
		},
		{
			TypeField:   &bcs.Type,
			TypeValue:   BotCommandScopeTypeBotCommandScopeAllChatAdministrators,
			TargetField: &bcs.BotCommandScopeAllChatAdministrators,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("user_id", "chat_id"),
			),
		},
		{
			TypeField:   &bcs.Type,
			TypeValue:   BotCommandScopeTypeBotCommandScopeChat,
			TargetField: &bcs.BotCommandScopeChat,
			Rules: CombineRules(
				RequiredFields("chat_id"),
				ForbiddenFields("user_id"),
			),
		},
		{
			TypeField:   &bcs.Type,
			TypeValue:   BotCommandScopeTypeBotCommandScopeChatAdministrators,
			TargetField: &bcs.BotCommandScopeChatAdministrators,
			Rules: CombineRules(
				RequiredFields("chat_id"),
				ForbiddenFields("user_id"),
			),
		},
		{
			TypeField:   &bcs.Type,
			TypeValue:   BotCommandScopeTypeBotCommandScopeChatMember,
			TargetField: &bcs.BotCommandScopeChatMember,
			Rules: CombineRules(
				RequiredFields("user_id", "chat_id"),
				ForbiddenFields(""),
			),
		},
	})
}

func (bcs *BotCommandScope) MarshalJSON() ([]byte, error) {
	return MarshalUnion(bcs.Type, map[BotCommandScopeType]any{
		BotCommandScopeTypeBotCommandScopeDefault:               bcs.BotCommandScopeDefault,
		BotCommandScopeTypeBotCommandScopeAllPrivateChats:       bcs.BotCommandScopeAllPrivateChats,
		BotCommandScopeTypeBotCommandScopeAllGroupChats:         bcs.BotCommandScopeAllGroupChats,
		BotCommandScopeTypeBotCommandScopeAllChatAdministrators: bcs.BotCommandScopeAllChatAdministrators,
		BotCommandScopeTypeBotCommandScopeChat:                  bcs.BotCommandScopeChat,
		BotCommandScopeTypeBotCommandScopeChatAdministrators:    bcs.BotCommandScopeChatAdministrators,
		BotCommandScopeTypeBotCommandScopeChatMember:            bcs.BotCommandScopeChatMember,
	})
}

func (mb *MenuButton) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[MenuButtonType]{
		{
			TypeField:   &mb.Type,
			TypeValue:   MenuButtonTypeMenuButtonCommands,
			TargetField: &mb.MenuButtonCommands,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("web_app", "text"),
			),
		},
		{
			TypeField:   &mb.Type,
			TypeValue:   MenuButtonTypeMenuButtonWebApp,
			TargetField: &mb.MenuButtonWebApp,
			Rules: CombineRules(
				RequiredFields("web_app", "text"),
				ForbiddenFields(""),
			),
		},
		{
			TypeField:   &mb.Type,
			TypeValue:   MenuButtonTypeMenuButtonDefault,
			TargetField: &mb.MenuButtonDefault,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("web_app", "text"),
			),
		},
	})
}

func (mb *MenuButton) MarshalJSON() ([]byte, error) {
	return MarshalUnion(mb.Type, map[MenuButtonType]any{
		MenuButtonTypeMenuButtonCommands: mb.MenuButtonCommands,
		MenuButtonTypeMenuButtonWebApp:   mb.MenuButtonWebApp,
		MenuButtonTypeMenuButtonDefault:  mb.MenuButtonDefault,
	})
}

func (cbs *ChatBoostSource) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[ChatBoostSourceType]{
		{
			TypeField:   &cbs.Type,
			TypeValue:   ChatBoostSourceTypeChatBoostSourcePremium,
			TargetField: &cbs.ChatBoostSourcePremium,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("prize_star_count", "is_unclaimed", "giveaway_message_id"),
			),
		},
		{
			TypeField:   &cbs.Type,
			TypeValue:   ChatBoostSourceTypeChatBoostSourceGiftCode,
			TargetField: &cbs.ChatBoostSourceGiftCode,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("prize_star_count", "is_unclaimed", "giveaway_message_id"),
			),
		},
		{
			TypeField:   &cbs.Type,
			TypeValue:   ChatBoostSourceTypeChatBoostSourceGiveaway,
			TargetField: &cbs.ChatBoostSourceGiveaway,
			Rules: CombineRules(
				RequiredFields("giveaway_message_id"),
				ForbiddenFields(""),
			),
		},
	})
}

func (cbs *ChatBoostSource) MarshalJSON() ([]byte, error) {
	return MarshalUnion(cbs.Type, map[ChatBoostSourceType]any{
		ChatBoostSourceTypeChatBoostSourcePremium:  cbs.ChatBoostSourcePremium,
		ChatBoostSourceTypeChatBoostSourceGiftCode: cbs.ChatBoostSourceGiftCode,
		ChatBoostSourceTypeChatBoostSourceGiveaway: cbs.ChatBoostSourceGiveaway,
	})
}

func (im *InputMedia) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[InputMediaType]{
		{
			TypeField:   &im.Type,
			TypeValue:   InputMediaTypeInputMediaAnimation,
			TargetField: &im.InputMediaAnimation,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("performer", "supports_streaming", "cover", "start_timestamp", "title", "disable_content_type_detection"),
			),
		},
		{
			TypeField:   &im.Type,
			TypeValue:   InputMediaTypeInputMediaDocument,
			TargetField: &im.InputMediaDocument,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("performer", "show_caption_above_media", "height", "width", "title", "duration", "has_spoiler"),
			),
		},
		{
			TypeField:   &im.Type,
			TypeValue:   InputMediaTypeInputMediaAudio,
			TargetField: &im.InputMediaAudio,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("cover", "show_caption_above_media", "height", "width", "has_spoiler", "disable_content_type_detection"),
			),
		},
		{
			TypeField:   &im.Type,
			TypeValue:   InputMediaTypeInputMediaPhoto,
			TargetField: &im.InputMediaPhoto,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("thumbnail", "height", "width", "duration", "disable_content_type_detection"),
			),
		},
		{
			TypeField:   &im.Type,
			TypeValue:   InputMediaTypeInputMediaVideo,
			TargetField: &im.InputMediaVideo,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("title", "disable_content_type_detection", "performer"),
			),
		},
	})
}

func (im *InputMedia) MarshalJSON() ([]byte, error) {
	return MarshalUnion(im.Type, map[InputMediaType]any{
		InputMediaTypeInputMediaAnimation: im.InputMediaAnimation,
		InputMediaTypeInputMediaDocument:  im.InputMediaDocument,
		InputMediaTypeInputMediaAudio:     im.InputMediaAudio,
		InputMediaTypeInputMediaPhoto:     im.InputMediaPhoto,
		InputMediaTypeInputMediaVideo:     im.InputMediaVideo,
	})
}

func (ipm *InputPaidMedia) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[InputPaidMediaType]{
		{
			TypeField:   &ipm.Type,
			TypeValue:   InputPaidMediaTypeInputPaidMediaPhoto,
			TargetField: &ipm.InputPaidMediaPhoto,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("supports_streaming", "cover", "thumbnail", "start_timestamp", "height", "width", "duration"),
			),
		},
		{
			TypeField:   &ipm.Type,
			TypeValue:   InputPaidMediaTypeInputPaidMediaVideo,
			TargetField: &ipm.InputPaidMediaVideo,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields(""),
			),
		},
	})
}

func (ipm *InputPaidMedia) MarshalJSON() ([]byte, error) {
	return MarshalUnion(ipm.Type, map[InputPaidMediaType]any{
		InputPaidMediaTypeInputPaidMediaPhoto: ipm.InputPaidMediaPhoto,
		InputPaidMediaTypeInputPaidMediaVideo: ipm.InputPaidMediaVideo,
	})
}

func (ipp *InputProfilePhoto) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[InputProfilePhotoType]{
		{
			TypeField:   &ipp.Type,
			TypeValue:   InputProfilePhotoTypeInputProfilePhotoStatic,
			TargetField: &ipp.InputProfilePhotoStatic,
			Rules: CombineRules(
				RequiredFields("photo"),
				ForbiddenFields("main_frame_timestamp", "animation"),
			),
		},
		{
			TypeField:   &ipp.Type,
			TypeValue:   InputProfilePhotoTypeInputProfilePhotoAnimated,
			TargetField: &ipp.InputProfilePhotoAnimated,
			Rules: CombineRules(
				RequiredFields("animation"),
				ForbiddenFields("photo"),
			),
		},
	})
}

func (ipp *InputProfilePhoto) MarshalJSON() ([]byte, error) {
	return MarshalUnion(ipp.Type, map[InputProfilePhotoType]any{
		InputProfilePhotoTypeInputProfilePhotoStatic:   ipp.InputProfilePhotoStatic,
		InputProfilePhotoTypeInputProfilePhotoAnimated: ipp.InputProfilePhotoAnimated,
	})
}

func (isc *InputStoryContent) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[InputStoryContentType]{
		{
			TypeField:   &isc.Type,
			TypeValue:   InputStoryContentTypeInputStoryContentPhoto,
			TargetField: &isc.InputStoryContentPhoto,
			Rules: CombineRules(
				RequiredFields("photo"),
				ForbiddenFields("video", "duration", "cover_frame_timestamp", "is_animation"),
			),
		},
		{
			TypeField:   &isc.Type,
			TypeValue:   InputStoryContentTypeInputStoryContentVideo,
			TargetField: &isc.InputStoryContentVideo,
			Rules: CombineRules(
				RequiredFields("video"),
				ForbiddenFields("photo"),
			),
		},
	})
}

func (isc *InputStoryContent) MarshalJSON() ([]byte, error) {
	return MarshalUnion(isc.Type, map[InputStoryContentType]any{
		InputStoryContentTypeInputStoryContentPhoto: isc.InputStoryContentPhoto,
		InputStoryContentTypeInputStoryContentVideo: isc.InputStoryContentVideo,
	})
}

func (iqr *InlineQueryResult) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[InlineQueryResultType]{
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedAudio,
			TargetField: &iqr.InlineQueryResultCachedAudio,
			Rules: CombineRules(
				RequiredFields("audio_file_id"),
				ForbiddenFields("mpeg4_file_id", "show_caption_above_media", "title", "description", "document_file_id", "gif_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedDocument,
			TargetField: &iqr.InlineQueryResultCachedDocument,
			Rules: CombineRules(
				RequiredFields("document_file_id", "title"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "sticker_file_id", "show_caption_above_media", "gif_file_id", "photo_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedGif,
			TargetField: &iqr.InlineQueryResultCachedGif,
			Rules: CombineRules(
				RequiredFields("gif_file_id"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "sticker_file_id", "video_file_id", "description", "document_file_id", "photo_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedMpeg4Gif,
			TargetField: &iqr.InlineQueryResultCachedMpeg4Gif,
			Rules: CombineRules(
				RequiredFields("mpeg4_file_id"),
				ForbiddenFields("audio_file_id", "sticker_file_id", "video_file_id", "description", "document_file_id", "photo_file_id", "gif_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedPhoto,
			TargetField: &iqr.InlineQueryResultCachedPhoto,
			Rules: CombineRules(
				RequiredFields("photo_file_id"),
				ForbiddenFields("voice_file_id", "audio_file_id", "mpeg4_file_id", "sticker_file_id", "video_file_id", "gif_file_id", "url", "document_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedSticker,
			TargetField: &iqr.InlineQueryResultCachedSticker,
			Rules: CombineRules(
				RequiredFields("sticker_file_id"),
				ForbiddenFields("parse_mode", "audio_file_id", "caption", "caption_entities", "title", "description", "document_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedVideo,
			TargetField: &iqr.InlineQueryResultCachedVideo,
			Rules: CombineRules(
				RequiredFields("video_file_id", "title"),
				ForbiddenFields("voice_file_id", "audio_file_id", "mpeg4_file_id", "sticker_file_id", "gif_file_id", "url", "document_file_id", "photo_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultCachedVoice,
			TargetField: &iqr.InlineQueryResultCachedVoice,
			Rules: CombineRules(
				RequiredFields("title", "voice_file_id"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "show_caption_above_media", "description", "document_file_id", "photo_file_id", "gif_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultArticle,
			TargetField: &iqr.InlineQueryResultArticle,
			Rules: CombineRules(
				RequiredFields("input_message_content", "title"),
				ForbiddenFields("parse_mode", "audio_file_id", "caption", "caption_entities", "document_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultAudio,
			TargetField: &iqr.InlineQueryResultAudio,
			Rules: CombineRules(
				RequiredFields("audio_url", "title"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "show_caption_above_media", "description", "document_file_id", "photo_file_id", "gif_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultContact,
			TargetField: &iqr.InlineQueryResultContact,
			Rules: CombineRules(
				RequiredFields("first_name", "phone_number"),
				ForbiddenFields("parse_mode", "audio_file_id", "caption", "caption_entities", "title", "description", "document_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultGame,
			TargetField: &iqr.InlineQueryResultGame,
			Rules: CombineRules(
				RequiredFields("game_short_name"),
				ForbiddenFields("parse_mode", "audio_file_id", "input_message_content", "caption", "caption_entities", "title", "description", "document_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultDocument,
			TargetField: &iqr.InlineQueryResultDocument,
			Rules: CombineRules(
				RequiredFields("document_url", "mime_type", "title"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "show_caption_above_media", "gif_file_id", "document_file_id", "photo_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultGif,
			TargetField: &iqr.InlineQueryResultGif,
			Rules: CombineRules(
				RequiredFields("thumbnail_url", "gif_url"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "sticker_file_id", "description", "document_file_id", "photo_file_id", "gif_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultLocation,
			TargetField: &iqr.InlineQueryResultLocation,
			Rules: CombineRules(
				RequiredFields("latitude", "longitude", "title"),
				ForbiddenFields("parse_mode", "audio_file_id", "caption", "caption_entities", "description", "document_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultMpeg4Gif,
			TargetField: &iqr.InlineQueryResultMpeg4Gif,
			Rules: CombineRules(
				RequiredFields("thumbnail_url", "mpeg4_url"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "sticker_file_id", "description", "document_file_id", "photo_file_id", "gif_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultPhoto,
			TargetField: &iqr.InlineQueryResultPhoto,
			Rules: CombineRules(
				RequiredFields("thumbnail_url", "photo_url"),
				ForbiddenFields("voice_file_id", "audio_file_id", "mpeg4_file_id", "sticker_file_id", "video_file_id", "gif_file_id", "document_file_id", "photo_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultVenue,
			TargetField: &iqr.InlineQueryResultVenue,
			Rules: CombineRules(
				RequiredFields("address", "latitude", "longitude", "title"),
				ForbiddenFields("parse_mode", "audio_file_id", "caption", "caption_entities", "description", "document_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultVideo,
			TargetField: &iqr.InlineQueryResultVideo,
			Rules: CombineRules(
				RequiredFields("thumbnail_url", "video_url", "title", "mime_type"),
				ForbiddenFields("voice_file_id", "audio_file_id", "mpeg4_file_id", "sticker_file_id", "video_file_id", "gif_file_id", "document_file_id", "photo_file_id"),
			),
		},
		{
			TypeField:   &iqr.Type,
			TypeValue:   InlineQueryResultTypeInlineQueryResultVoice,
			TargetField: &iqr.InlineQueryResultVoice,
			Rules: CombineRules(
				RequiredFields("voice_url", "title"),
				ForbiddenFields("audio_file_id", "mpeg4_file_id", "show_caption_above_media", "description", "document_file_id", "photo_file_id", "gif_file_id"),
			),
		},
	})
}

func (iqr *InlineQueryResult) MarshalJSON() ([]byte, error) {
	return MarshalUnion(iqr.Type, map[InlineQueryResultType]any{
		InlineQueryResultTypeInlineQueryResultCachedAudio:    iqr.InlineQueryResultCachedAudio,
		InlineQueryResultTypeInlineQueryResultCachedDocument: iqr.InlineQueryResultCachedDocument,
		InlineQueryResultTypeInlineQueryResultCachedGif:      iqr.InlineQueryResultCachedGif,
		InlineQueryResultTypeInlineQueryResultCachedMpeg4Gif: iqr.InlineQueryResultCachedMpeg4Gif,
		InlineQueryResultTypeInlineQueryResultCachedPhoto:    iqr.InlineQueryResultCachedPhoto,
		InlineQueryResultTypeInlineQueryResultCachedSticker:  iqr.InlineQueryResultCachedSticker,
		InlineQueryResultTypeInlineQueryResultCachedVideo:    iqr.InlineQueryResultCachedVideo,
		InlineQueryResultTypeInlineQueryResultCachedVoice:    iqr.InlineQueryResultCachedVoice,
		InlineQueryResultTypeInlineQueryResultArticle:        iqr.InlineQueryResultArticle,
		InlineQueryResultTypeInlineQueryResultAudio:          iqr.InlineQueryResultAudio,
		InlineQueryResultTypeInlineQueryResultContact:        iqr.InlineQueryResultContact,
		InlineQueryResultTypeInlineQueryResultGame:           iqr.InlineQueryResultGame,
		InlineQueryResultTypeInlineQueryResultDocument:       iqr.InlineQueryResultDocument,
		InlineQueryResultTypeInlineQueryResultGif:            iqr.InlineQueryResultGif,
		InlineQueryResultTypeInlineQueryResultLocation:       iqr.InlineQueryResultLocation,
		InlineQueryResultTypeInlineQueryResultMpeg4Gif:       iqr.InlineQueryResultMpeg4Gif,
		InlineQueryResultTypeInlineQueryResultPhoto:          iqr.InlineQueryResultPhoto,
		InlineQueryResultTypeInlineQueryResultVenue:          iqr.InlineQueryResultVenue,
		InlineQueryResultTypeInlineQueryResultVideo:          iqr.InlineQueryResultVideo,
		InlineQueryResultTypeInlineQueryResultVoice:          iqr.InlineQueryResultVoice,
	})
}

func (imc *InputMessageContent) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[InputMessageContentType]{
		{
			TypeField:   &imc.Type,
			TypeValue:   InputMessageContentTypeInputTextMessageContent,
			TargetField: &imc.InputTextMessageContent,
			Rules: CombineRules(
				RequiredFields("message_text"),
				ForbiddenFields("proximity_alert_radius", "longitude", "horizontal_accuracy", "heading", "latitude", "live_period"),
			),
		},
		{
			TypeField:   &imc.Type,
			TypeValue:   InputMessageContentTypeInputLocationMessageContent,
			TargetField: &imc.InputLocationMessageContent,
			Rules: CombineRules(
				RequiredFields("latitude", "longitude"),
				ForbiddenFields("link_preview_options", "parse_mode", "message_text", "foursquare_type", "foursquare_id", "address", "title", "entities"),
			),
		},
		{
			TypeField:   &imc.Type,
			TypeValue:   InputMessageContentTypeInputVenueMessageContent,
			TargetField: &imc.InputVenueMessageContent,
			Rules: CombineRules(
				RequiredFields("address", "latitude", "longitude", "title"),
				ForbiddenFields("link_preview_options", "parse_mode", "message_text", "proximity_alert_radius", "horizontal_accuracy", "heading", "entities", "live_period"),
			),
		},
		{
			TypeField:   &imc.Type,
			TypeValue:   InputMessageContentTypeInputContactMessageContent,
			TargetField: &imc.InputContactMessageContent,
			Rules: CombineRules(
				RequiredFields("first_name", "phone_number"),
				ForbiddenFields("link_preview_options", "parse_mode", "message_text", "longitude", "horizontal_accuracy", "latitude", "entities", "live_period"),
			),
		},
		{
			TypeField:   &imc.Type,
			TypeValue:   InputMessageContentTypeInputInvoiceMessageContent,
			TargetField: &imc.InputInvoiceMessageContent,
			Rules: CombineRules(
				RequiredFields("currency", "payload", "title", "description", "prices"),
				ForbiddenFields("link_preview_options", "parse_mode", "message_text", "longitude", "horizontal_accuracy", "latitude", "entities", "live_period"),
			),
		},
	})
}

func (imc *InputMessageContent) MarshalJSON() ([]byte, error) {
	return MarshalUnion(imc.Type, map[InputMessageContentType]any{
		InputMessageContentTypeInputTextMessageContent:     imc.InputTextMessageContent,
		InputMessageContentTypeInputLocationMessageContent: imc.InputLocationMessageContent,
		InputMessageContentTypeInputVenueMessageContent:    imc.InputVenueMessageContent,
		InputMessageContentTypeInputContactMessageContent:  imc.InputContactMessageContent,
		InputMessageContentTypeInputInvoiceMessageContent:  imc.InputInvoiceMessageContent,
	})
}

func (rws *RevenueWithdrawalState) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[RevenueWithdrawalStateType]{
		{
			TypeField:   &rws.Type,
			TypeValue:   RevenueWithdrawalStateTypeRevenueWithdrawalStatePending,
			TargetField: &rws.RevenueWithdrawalStatePending,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("date", "url"),
			),
		},
		{
			TypeField:   &rws.Type,
			TypeValue:   RevenueWithdrawalStateTypeRevenueWithdrawalStateSucceeded,
			TargetField: &rws.RevenueWithdrawalStateSucceeded,
			Rules: CombineRules(
				RequiredFields("date", "url"),
				ForbiddenFields(""),
			),
		},
		{
			TypeField:   &rws.Type,
			TypeValue:   RevenueWithdrawalStateTypeRevenueWithdrawalStateFailed,
			TargetField: &rws.RevenueWithdrawalStateFailed,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("date", "url"),
			),
		},
	})
}

func (rws *RevenueWithdrawalState) MarshalJSON() ([]byte, error) {
	return MarshalUnion(rws.Type, map[RevenueWithdrawalStateType]any{
		RevenueWithdrawalStateTypeRevenueWithdrawalStatePending:   rws.RevenueWithdrawalStatePending,
		RevenueWithdrawalStateTypeRevenueWithdrawalStateSucceeded: rws.RevenueWithdrawalStateSucceeded,
		RevenueWithdrawalStateTypeRevenueWithdrawalStateFailed:    rws.RevenueWithdrawalStateFailed,
	})
}

func (tp *TransactionPartner) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[TransactionPartnerType]{
		{
			TypeField:   &tp.Type,
			TypeValue:   TransactionPartnerTypeTransactionPartnerUser,
			TargetField: &tp.TransactionPartnerUser,
			Rules: CombineRules(
				RequiredFields("user", "transaction_type"),
				ForbiddenFields("sponsor_user", "chat", "commission_per_mille", "withdrawal_state", "request_count"),
			),
		},
		{
			TypeField:   &tp.Type,
			TypeValue:   TransactionPartnerTypeTransactionPartnerChat,
			TargetField: &tp.TransactionPartnerChat,
			Rules: CombineRules(
				RequiredFields("chat"),
				ForbiddenFields("invoice_payload", "user", "subscription_period", "affiliate", "transaction_type", "paid_media", "premium_subscription_duration", "paid_media_payload"),
			),
		},
		{
			TypeField:   &tp.Type,
			TypeValue:   TransactionPartnerTypeTransactionPartnerAffiliateProgram,
			TargetField: &tp.TransactionPartnerAffiliateProgram,
			Rules: CombineRules(
				RequiredFields("commission_per_mille"),
				ForbiddenFields("invoice_payload", "user", "subscription_period", "affiliate", "transaction_type", "paid_media", "paid_media_payload", "gift"),
			),
		},
		{
			TypeField:   &tp.Type,
			TypeValue:   TransactionPartnerTypeTransactionPartnerFragment,
			TargetField: &tp.TransactionPartnerFragment,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("invoice_payload", "user", "subscription_period", "affiliate", "transaction_type", "paid_media", "paid_media_payload", "gift"),
			),
		},
		{
			TypeField:   &tp.Type,
			TypeValue:   TransactionPartnerTypeTransactionPartnerTelegramAds,
			TargetField: &tp.TransactionPartnerTelegramAds,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("invoice_payload", "user", "subscription_period", "affiliate", "transaction_type", "paid_media", "paid_media_payload", "gift"),
			),
		},
		{
			TypeField:   &tp.Type,
			TypeValue:   TransactionPartnerTypeTransactionPartnerTelegramApi,
			TargetField: &tp.TransactionPartnerTelegramApi,
			Rules: CombineRules(
				RequiredFields("request_count"),
				ForbiddenFields("invoice_payload", "user", "subscription_period", "affiliate", "transaction_type", "paid_media", "paid_media_payload", "gift"),
			),
		},
		{
			TypeField:   &tp.Type,
			TypeValue:   TransactionPartnerTypeTransactionPartnerOther,
			TargetField: &tp.TransactionPartnerOther,
			Rules: CombineRules(
				RequiredFields(""),
				ForbiddenFields("invoice_payload", "user", "subscription_period", "affiliate", "transaction_type", "paid_media", "paid_media_payload", "gift"),
			),
		},
	})
}

func (tp *TransactionPartner) MarshalJSON() ([]byte, error) {
	return MarshalUnion(tp.Type, map[TransactionPartnerType]any{
		TransactionPartnerTypeTransactionPartnerUser:             tp.TransactionPartnerUser,
		TransactionPartnerTypeTransactionPartnerChat:             tp.TransactionPartnerChat,
		TransactionPartnerTypeTransactionPartnerAffiliateProgram: tp.TransactionPartnerAffiliateProgram,
		TransactionPartnerTypeTransactionPartnerFragment:         tp.TransactionPartnerFragment,
		TransactionPartnerTypeTransactionPartnerTelegramAds:      tp.TransactionPartnerTelegramAds,
		TransactionPartnerTypeTransactionPartnerTelegramApi:      tp.TransactionPartnerTelegramApi,
		TransactionPartnerTypeTransactionPartnerOther:            tp.TransactionPartnerOther,
	})
}

func (pee *PassportElementError) UnmarshalJSON(data []byte) error {
	return UnmarshalUnion(data, []UnionConfig[PassportElementErrorType]{
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorDataField,
			TargetField: &pee.PassportElementErrorDataField,
			Rules: CombineRules(
				RequiredFields("data_hash", "field_name"),
				ForbiddenFields("file_hash", "file_hashes", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorFrontSide,
			TargetField: &pee.PassportElementErrorFrontSide,
			Rules: CombineRules(
				RequiredFields("file_hash"),
				ForbiddenFields("file_hashes", "data_hash", "field_name", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorReverseSide,
			TargetField: &pee.PassportElementErrorReverseSide,
			Rules: CombineRules(
				RequiredFields("file_hash"),
				ForbiddenFields("file_hashes", "data_hash", "field_name", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorSelfie,
			TargetField: &pee.PassportElementErrorSelfie,
			Rules: CombineRules(
				RequiredFields("file_hash"),
				ForbiddenFields("file_hashes", "data_hash", "field_name", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorFile,
			TargetField: &pee.PassportElementErrorFile,
			Rules: CombineRules(
				RequiredFields("file_hash"),
				ForbiddenFields("file_hashes", "data_hash", "field_name", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorFiles,
			TargetField: &pee.PassportElementErrorFiles,
			Rules: CombineRules(
				RequiredFields("file_hashes"),
				ForbiddenFields("file_hash", "data_hash", "field_name", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorTranslationFile,
			TargetField: &pee.PassportElementErrorTranslationFile,
			Rules: CombineRules(
				RequiredFields("file_hash"),
				ForbiddenFields("file_hashes", "data_hash", "field_name", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorTranslationFiles,
			TargetField: &pee.PassportElementErrorTranslationFiles,
			Rules: CombineRules(
				RequiredFields("file_hashes"),
				ForbiddenFields("file_hash", "data_hash", "field_name", "element_hash"),
			),
		},
		{
			TypeField:   &pee.Type,
			TypeValue:   PassportElementErrorTypePassportElementErrorUnspecified,
			TargetField: &pee.PassportElementErrorUnspecified,
			Rules: CombineRules(
				RequiredFields("element_hash"),
				ForbiddenFields("file_hash", "file_hashes", "data_hash", "field_name"),
			),
		},
	})
}

func (pee *PassportElementError) MarshalJSON() ([]byte, error) {
	return MarshalUnion(pee.Type, map[PassportElementErrorType]any{
		PassportElementErrorTypePassportElementErrorDataField:        pee.PassportElementErrorDataField,
		PassportElementErrorTypePassportElementErrorFrontSide:        pee.PassportElementErrorFrontSide,
		PassportElementErrorTypePassportElementErrorReverseSide:      pee.PassportElementErrorReverseSide,
		PassportElementErrorTypePassportElementErrorSelfie:           pee.PassportElementErrorSelfie,
		PassportElementErrorTypePassportElementErrorFile:             pee.PassportElementErrorFile,
		PassportElementErrorTypePassportElementErrorFiles:            pee.PassportElementErrorFiles,
		PassportElementErrorTypePassportElementErrorTranslationFile:  pee.PassportElementErrorTranslationFile,
		PassportElementErrorTypePassportElementErrorTranslationFiles: pee.PassportElementErrorTranslationFiles,
		PassportElementErrorTypePassportElementErrorUnspecified:      pee.PassportElementErrorUnspecified,
	})
}
