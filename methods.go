package gogram

import (
	"context"

	"github.com/Alsond5/gogram/models"
)

func (b *Bot) GetUpdates(ctx context.Context, params *models.GetUpdatesRequest) ([]models.Update, error) {
	var result []models.Update

	err := request(b, ctx, "getUpdates", params, &result)
	return result, err
}

func (b *Bot) SetWebhook(ctx context.Context, params *models.SetWebhookRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setWebhook", params, &result)
	return result, err
}

func (b *Bot) DeleteWebhook(ctx context.Context, params *models.DeleteWebhookRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteWebhook", params, &result)
	return result, err
}

func (b *Bot) GetWebhookInfo(ctx context.Context) (*models.WebhookInfo, error) {
	result := new(models.WebhookInfo)

	err := request[any](b, ctx, "getWebhookInfo", nil, result)
	return result, err
}

func (b *Bot) GetMe(ctx context.Context) (*models.User, error) {
	result := new(models.User)

	err := request[any](b, ctx, "getMe", nil, result)
	return result, err
}

func (b *Bot) LogOut(ctx context.Context) (bool, error) {
	var result bool

	err := request[any](b, ctx, "logOut", nil, &result)
	return result, err
}

func (b *Bot) Close(ctx context.Context) (bool, error) {
	var result bool

	err := request[any](b, ctx, "close", nil, &result)
	return result, err
}

func (b *Bot) SendMessage(ctx context.Context, params *models.SendMessageRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendMessage", params, result)
	return result, err
}

func (b *Bot) ForwardMessage(ctx context.Context, params *models.ForwardMessageRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "forwardMessage", params, result)
	return result, err
}

func (b *Bot) ForwardMessages(ctx context.Context, params *models.ForwardMessagesRequest) ([]models.MessageId, error) {
	var result []models.MessageId

	err := request(b, ctx, "forwardMessages", params, &result)
	return result, err
}

func (b *Bot) CopyMessage(ctx context.Context, params *models.CopyMessageRequest) (*models.MessageId, error) {
	result := new(models.MessageId)

	err := request(b, ctx, "copyMessage", params, result)
	return result, err
}

func (b *Bot) CopyMessages(ctx context.Context, params *models.CopyMessagesRequest) ([]models.MessageId, error) {
	var result []models.MessageId

	err := request(b, ctx, "copyMessages", params, &result)
	return result, err
}

func (b *Bot) SendPhoto(ctx context.Context, params *models.SendPhotoRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendPhoto", params, result)
	return result, err
}

func (b *Bot) SendAudio(ctx context.Context, params *models.SendAudioRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendAudio", params, result)
	return result, err
}

func (b *Bot) SendDocument(ctx context.Context, params *models.SendDocumentRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendDocument", params, result)
	return result, err
}

func (b *Bot) SendVideo(ctx context.Context, params *models.SendVideoRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendVideo", params, result)
	return result, err
}

func (b *Bot) SendAnimation(ctx context.Context, params *models.SendAnimationRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendAnimation", params, result)
	return result, err
}

func (b *Bot) SendVoice(ctx context.Context, params *models.SendVoiceRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendVoice", params, result)
	return result, err
}

func (b *Bot) SendVideoNote(ctx context.Context, params *models.SendVideoNoteRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendVideoNote", params, result)
	return result, err
}

func (b *Bot) SendPaidMedia(ctx context.Context, params *models.SendPaidMediaRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendPaidMedia", params, result)
	return result, err
}

func (b *Bot) SendMediaGroup(ctx context.Context, params *models.SendMediaGroupRequest) ([]models.Message, error) {
	var result []models.Message

	err := request(b, ctx, "sendMediaGroup", params, &result)
	return result, err
}

func (b *Bot) SendLocation(ctx context.Context, params *models.SendLocationRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendLocation", params, result)
	return result, err
}

func (b *Bot) SendVenue(ctx context.Context, params *models.SendVenueRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendVenue", params, result)
	return result, err
}

func (b *Bot) SendContact(ctx context.Context, params *models.SendContactRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendContact", params, result)
	return result, err
}

func (b *Bot) SendPoll(ctx context.Context, params *models.SendPollRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendPoll", params, result)
	return result, err
}

func (b *Bot) SendChecklist(ctx context.Context, params *models.SendChecklistRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendChecklist", params, result)
	return result, err
}

func (b *Bot) SendDice(ctx context.Context, params *models.SendDiceRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendDice", params, result)
	return result, err
}

func (b *Bot) SendChatAction(ctx context.Context, params *models.SendChatActionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "sendChatAction", params, &result)
	return result, err
}

func (b *Bot) SetMessageReaction(ctx context.Context, params *models.SetMessageReactionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setMessageReaction", params, &result)
	return result, err
}

func (b *Bot) GetUserProfilePhotos(ctx context.Context, params *models.GetUserProfilePhotosRequest) (*models.UserProfilePhotos, error) {
	result := new(models.UserProfilePhotos)

	err := request(b, ctx, "getUserProfilePhotos", params, result)
	return result, err
}

func (b *Bot) SetUserEmojiStatus(ctx context.Context, params *models.SetUserEmojiStatusRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setUserEmojiStatus", params, &result)
	return result, err
}

func (b *Bot) GetFile(ctx context.Context, params *models.GetFileRequest) (*models.File, error) {
	result := new(models.File)

	err := request(b, ctx, "getFile", params, result)
	return result, err
}

func (b *Bot) BanChatMember(ctx context.Context, params *models.BanChatMemberRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "banChatMember", params, &result)
	return result, err
}

func (b *Bot) UnbanChatMember(ctx context.Context, params *models.UnbanChatMemberRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "unbanChatMember", params, &result)
	return result, err
}

func (b *Bot) RestrictChatMember(ctx context.Context, params *models.RestrictChatMemberRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "restrictChatMember", params, &result)
	return result, err
}

func (b *Bot) PromoteChatMember(ctx context.Context, params *models.PromoteChatMemberRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "promoteChatMember", params, &result)
	return result, err
}

func (b *Bot) SetChatAdministratorCustomTitle(ctx context.Context, params *models.SetChatAdministratorCustomTitleRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setChatAdministratorCustomTitle", params, &result)
	return result, err
}

func (b *Bot) BanChatSenderChat(ctx context.Context, params *models.BanChatSenderChatRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "banChatSenderChat", params, &result)
	return result, err
}

func (b *Bot) UnbanChatSenderChat(ctx context.Context, params *models.UnbanChatSenderChatRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "unbanChatSenderChat", params, &result)
	return result, err
}

func (b *Bot) SetChatPermissions(ctx context.Context, params *models.SetChatPermissionsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setChatPermissions", params, &result)
	return result, err
}

func (b *Bot) ExportChatInviteLink(ctx context.Context, params *models.ExportChatInviteLinkRequest) (string, error) {
	var result string

	err := request(b, ctx, "exportChatInviteLink", params, &result)
	return result, err
}

func (b *Bot) CreateChatInviteLink(ctx context.Context, params *models.CreateChatInviteLinkRequest) (*models.ChatInviteLink, error) {
	result := new(models.ChatInviteLink)

	err := request(b, ctx, "createChatInviteLink", params, result)
	return result, err
}

func (b *Bot) EditChatInviteLink(ctx context.Context, params *models.EditChatInviteLinkRequest) (*models.ChatInviteLink, error) {
	result := new(models.ChatInviteLink)

	err := request(b, ctx, "editChatInviteLink", params, result)
	return result, err
}

func (b *Bot) CreateChatSubscriptionInviteLink(ctx context.Context, params *models.CreateChatSubscriptionInviteLinkRequest) (*models.ChatInviteLink, error) {
	result := new(models.ChatInviteLink)

	err := request(b, ctx, "createChatSubscriptionInviteLink", params, result)
	return result, err
}

func (b *Bot) EditChatSubscriptionInviteLink(ctx context.Context, params *models.EditChatSubscriptionInviteLinkRequest) (*models.ChatInviteLink, error) {
	result := new(models.ChatInviteLink)

	err := request(b, ctx, "editChatSubscriptionInviteLink", params, result)
	return result, err
}

func (b *Bot) RevokeChatInviteLink(ctx context.Context, params *models.RevokeChatInviteLinkRequest) (*models.ChatInviteLink, error) {
	result := new(models.ChatInviteLink)

	err := request(b, ctx, "revokeChatInviteLink", params, result)
	return result, err
}

func (b *Bot) ApproveChatJoinRequest(ctx context.Context, params *models.ApproveChatJoinRequestRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "approveChatJoinRequest", params, &result)
	return result, err
}

func (b *Bot) DeclineChatJoinRequest(ctx context.Context, params *models.DeclineChatJoinRequestRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "declineChatJoinRequest", params, &result)
	return result, err
}

func (b *Bot) SetChatPhoto(ctx context.Context, params *models.SetChatPhotoRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setChatPhoto", params, &result)
	return result, err
}

func (b *Bot) DeleteChatPhoto(ctx context.Context, params *models.DeleteChatPhotoRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteChatPhoto", params, &result)
	return result, err
}

func (b *Bot) SetChatTitle(ctx context.Context, params *models.SetChatTitleRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setChatTitle", params, &result)
	return result, err
}

func (b *Bot) SetChatDescription(ctx context.Context, params *models.SetChatDescriptionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setChatDescription", params, &result)
	return result, err
}

func (b *Bot) PinChatMessage(ctx context.Context, params *models.PinChatMessageRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "pinChatMessage", params, &result)
	return result, err
}

func (b *Bot) UnpinChatMessage(ctx context.Context, params *models.UnpinChatMessageRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "unpinChatMessage", params, &result)
	return result, err
}

func (b *Bot) UnpinAllChatMessages(ctx context.Context, params *models.UnpinAllChatMessagesRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "unpinAllChatMessages", params, &result)
	return result, err
}

func (b *Bot) LeaveChat(ctx context.Context, params *models.LeaveChatRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "leaveChat", params, &result)
	return result, err
}

func (b *Bot) GetChat(ctx context.Context, params *models.GetChatRequest) (*models.ChatFullInfo, error) {
	result := new(models.ChatFullInfo)

	err := request(b, ctx, "getChat", params, result)
	return result, err
}

func (b *Bot) GetChatAdministrators(ctx context.Context, params *models.GetChatAdministratorsRequest) ([]models.ChatMember, error) {
	var result []models.ChatMember

	err := request(b, ctx, "getChatAdministrators", params, &result)
	return result, err
}

func (b *Bot) GetChatMemberCount(ctx context.Context, params *models.GetChatMemberCountRequest) (int64, error) {
	var result int64

	err := request(b, ctx, "getChatMemberCount", params, &result)
	return result, err
}

func (b *Bot) GetChatMember(ctx context.Context, params *models.GetChatMemberRequest) (*models.ChatMember, error) {
	result := new(models.ChatMember)

	err := request(b, ctx, "getChatMember", params, result)
	return result, err
}

func (b *Bot) SetChatStickerSet(ctx context.Context, params *models.SetChatStickerSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setChatStickerSet", params, &result)
	return result, err
}

func (b *Bot) DeleteChatStickerSet(ctx context.Context, params *models.DeleteChatStickerSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteChatStickerSet", params, &result)
	return result, err
}

func (b *Bot) GetForumTopicIconStickers(ctx context.Context) ([]models.Sticker, error) {
	var result []models.Sticker

	err := request[any](b, ctx, "getForumTopicIconStickers", nil, &result)
	return result, err
}

func (b *Bot) CreateForumTopic(ctx context.Context, params *models.CreateForumTopicRequest) (*models.ForumTopic, error) {
	result := new(models.ForumTopic)

	err := request(b, ctx, "createForumTopic", params, result)
	return result, err
}

func (b *Bot) EditForumTopic(ctx context.Context, params *models.EditForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "editForumTopic", params, &result)
	return result, err
}

func (b *Bot) CloseForumTopic(ctx context.Context, params *models.CloseForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "closeForumTopic", params, &result)
	return result, err
}

func (b *Bot) ReopenForumTopic(ctx context.Context, params *models.ReopenForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "reopenForumTopic", params, &result)
	return result, err
}

func (b *Bot) DeleteForumTopic(ctx context.Context, params *models.DeleteForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteForumTopic", params, &result)
	return result, err
}

func (b *Bot) UnpinAllForumTopicMessages(ctx context.Context, params *models.UnpinAllForumTopicMessagesRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "unpinAllForumTopicMessages", params, &result)
	return result, err
}

func (b *Bot) EditGeneralForumTopic(ctx context.Context, params *models.EditGeneralForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "editGeneralForumTopic", params, &result)
	return result, err
}

func (b *Bot) CloseGeneralForumTopic(ctx context.Context, params *models.CloseGeneralForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "closeGeneralForumTopic", params, &result)
	return result, err
}

func (b *Bot) ReopenGeneralForumTopic(ctx context.Context, params *models.ReopenGeneralForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "reopenGeneralForumTopic", params, &result)
	return result, err
}

func (b *Bot) HideGeneralForumTopic(ctx context.Context, params *models.HideGeneralForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "hideGeneralForumTopic", params, &result)
	return result, err
}

func (b *Bot) UnhideGeneralForumTopic(ctx context.Context, params *models.UnhideGeneralForumTopicRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "unhideGeneralForumTopic", params, &result)
	return result, err
}

func (b *Bot) UnpinAllGeneralForumTopicMessages(ctx context.Context, params *models.UnpinAllGeneralForumTopicMessagesRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "unpinAllGeneralForumTopicMessages", params, &result)
	return result, err
}

func (b *Bot) AnswerCallbackQuery(ctx context.Context, params *models.AnswerCallbackQueryRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "answerCallbackQuery", params, &result)
	return result, err
}

func (b *Bot) GetUserChatBoosts(ctx context.Context, params *models.GetUserChatBoostsRequest) (*models.UserChatBoosts, error) {
	result := new(models.UserChatBoosts)

	err := request(b, ctx, "getUserChatBoosts", params, result)
	return result, err
}

func (b *Bot) GetBusinessConnection(ctx context.Context, params *models.GetBusinessConnectionRequest) (*models.BusinessConnection, error) {
	result := new(models.BusinessConnection)

	err := request(b, ctx, "getBusinessConnection", params, result)
	return result, err
}

func (b *Bot) SetMyCommands(ctx context.Context, params *models.SetMyCommandsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setMyCommands", params, &result)
	return result, err
}

func (b *Bot) DeleteMyCommands(ctx context.Context, params *models.DeleteMyCommandsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteMyCommands", params, &result)
	return result, err
}

func (b *Bot) GetMyCommands(ctx context.Context, params *models.GetMyCommandsRequest) ([]models.BotCommand, error) {
	var result []models.BotCommand

	err := request(b, ctx, "getMyCommands", params, &result)
	return result, err
}

func (b *Bot) SetMyName(ctx context.Context, params *models.SetMyNameRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setMyName", params, &result)
	return result, err
}

func (b *Bot) GetMyName(ctx context.Context, params *models.GetMyNameRequest) (*models.BotName, error) {
	result := new(models.BotName)

	err := request(b, ctx, "getMyName", params, result)
	return result, err
}

func (b *Bot) SetMyDescription(ctx context.Context, params *models.SetMyDescriptionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setMyDescription", params, &result)
	return result, err
}

func (b *Bot) GetMyDescription(ctx context.Context, params *models.GetMyDescriptionRequest) (*models.BotDescription, error) {
	result := new(models.BotDescription)

	err := request(b, ctx, "getMyDescription", params, result)
	return result, err
}

func (b *Bot) SetMyShortDescription(ctx context.Context, params *models.SetMyShortDescriptionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setMyShortDescription", params, &result)
	return result, err
}

func (b *Bot) GetMyShortDescription(ctx context.Context, params *models.GetMyShortDescriptionRequest) (*models.BotShortDescription, error) {
	result := new(models.BotShortDescription)

	err := request(b, ctx, "getMyShortDescription", params, result)
	return result, err
}

func (b *Bot) SetChatMenuButton(ctx context.Context, params *models.SetChatMenuButtonRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setChatMenuButton", params, &result)
	return result, err
}

func (b *Bot) GetChatMenuButton(ctx context.Context, params *models.GetChatMenuButtonRequest) (*models.MenuButton, error) {
	result := new(models.MenuButton)

	err := request(b, ctx, "getChatMenuButton", params, result)
	return result, err
}

func (b *Bot) SetMyDefaultAdministratorRights(ctx context.Context, params *models.SetMyDefaultAdministratorRightsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setMyDefaultAdministratorRights", params, &result)
	return result, err
}

func (b *Bot) GetMyDefaultAdministratorRights(ctx context.Context, params *models.GetMyDefaultAdministratorRightsRequest) (*models.ChatAdministratorRights, error) {
	result := new(models.ChatAdministratorRights)

	err := request(b, ctx, "getMyDefaultAdministratorRights", params, result)
	return result, err
}

func (b *Bot) GetAvailableGifts(ctx context.Context) (*models.Gifts, error) {
	result := new(models.Gifts)

	err := request[any](b, ctx, "getAvailableGifts", nil, result)
	return result, err
}

func (b *Bot) SendGift(ctx context.Context, params *models.SendGiftRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "sendGift", params, &result)
	return result, err
}

func (b *Bot) GiftPremiumSubscription(ctx context.Context, params *models.GiftPremiumSubscriptionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "giftPremiumSubscription", params, &result)
	return result, err
}

func (b *Bot) VerifyUser(ctx context.Context, params *models.VerifyUserRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "verifyUser", params, &result)
	return result, err
}

func (b *Bot) VerifyChat(ctx context.Context, params *models.VerifyChatRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "verifyChat", params, &result)
	return result, err
}

func (b *Bot) RemoveUserVerification(ctx context.Context, params *models.RemoveUserVerificationRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "removeUserVerification", params, &result)
	return result, err
}

func (b *Bot) RemoveChatVerification(ctx context.Context, params *models.RemoveChatVerificationRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "removeChatVerification", params, &result)
	return result, err
}

func (b *Bot) ReadBusinessMessage(ctx context.Context, params *models.ReadBusinessMessageRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "readBusinessMessage", params, &result)
	return result, err
}

func (b *Bot) DeleteBusinessMessages(ctx context.Context, params *models.DeleteBusinessMessagesRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteBusinessMessages", params, &result)
	return result, err
}

func (b *Bot) SetBusinessAccountName(ctx context.Context, params *models.SetBusinessAccountNameRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setBusinessAccountName", params, &result)
	return result, err
}

func (b *Bot) SetBusinessAccountUsername(ctx context.Context, params *models.SetBusinessAccountUsernameRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setBusinessAccountUsername", params, &result)
	return result, err
}

func (b *Bot) SetBusinessAccountBio(ctx context.Context, params *models.SetBusinessAccountBioRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setBusinessAccountBio", params, &result)
	return result, err
}

func (b *Bot) SetBusinessAccountProfilePhoto(ctx context.Context, params *models.SetBusinessAccountProfilePhotoRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setBusinessAccountProfilePhoto", params, &result)
	return result, err
}

func (b *Bot) RemoveBusinessAccountProfilePhoto(ctx context.Context, params *models.RemoveBusinessAccountProfilePhotoRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "removeBusinessAccountProfilePhoto", params, &result)
	return result, err
}

func (b *Bot) SetBusinessAccountGiftSettings(ctx context.Context, params *models.SetBusinessAccountGiftSettingsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setBusinessAccountGiftSettings", params, &result)
	return result, err
}

func (b *Bot) GetBusinessAccountStarBalance(ctx context.Context, params *models.GetBusinessAccountStarBalanceRequest) (*models.StarAmount, error) {
	result := new(models.StarAmount)

	err := request(b, ctx, "getBusinessAccountStarBalance", params, result)
	return result, err
}

func (b *Bot) TransferBusinessAccountStars(ctx context.Context, params *models.TransferBusinessAccountStarsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "transferBusinessAccountStars", params, &result)
	return result, err
}

func (b *Bot) GetBusinessAccountGifts(ctx context.Context, params *models.GetBusinessAccountGiftsRequest) (*models.OwnedGifts, error) {
	result := new(models.OwnedGifts)

	err := request(b, ctx, "getBusinessAccountGifts", params, result)
	return result, err
}

func (b *Bot) ConvertGiftToStars(ctx context.Context, params *models.ConvertGiftToStarsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "convertGiftToStars", params, &result)
	return result, err
}

func (b *Bot) UpgradeGift(ctx context.Context, params *models.UpgradeGiftRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "upgradeGift", params, &result)
	return result, err
}

func (b *Bot) TransferGift(ctx context.Context, params *models.TransferGiftRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "transferGift", params, &result)
	return result, err
}

func (b *Bot) PostStory(ctx context.Context, params *models.PostStoryRequest) (*models.Story, error) {
	result := new(models.Story)

	err := request(b, ctx, "postStory", params, result)
	return result, err
}

func (b *Bot) EditStory(ctx context.Context, params *models.EditStoryRequest) (*models.Story, error) {
	result := new(models.Story)

	err := request(b, ctx, "editStory", params, result)
	return result, err
}

func (b *Bot) DeleteStory(ctx context.Context, params *models.DeleteStoryRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteStory", params, &result)
	return result, err
}

func (b *Bot) EditMessageText(ctx context.Context, params *models.EditMessageTextRequest) (interface{}, error) {
	var result interface{}

	err := request(b, ctx, "editMessageText", params, &result)
	return result, err
}

func (b *Bot) EditMessageCaption(ctx context.Context, params *models.EditMessageCaptionRequest) (interface{}, error) {
	var result interface{}

	err := request(b, ctx, "editMessageCaption", params, &result)
	return result, err
}

func (b *Bot) EditMessageMedia(ctx context.Context, params *models.EditMessageMediaRequest) (interface{}, error) {
	var result interface{}

	err := request(b, ctx, "editMessageMedia", params, &result)
	return result, err
}

func (b *Bot) EditMessageLiveLocation(ctx context.Context, params *models.EditMessageLiveLocationRequest) (interface{}, error) {
	var result interface{}

	err := request(b, ctx, "editMessageLiveLocation", params, &result)
	return result, err
}

func (b *Bot) StopMessageLiveLocation(ctx context.Context, params *models.StopMessageLiveLocationRequest) (interface{}, error) {
	var result interface{}

	err := request(b, ctx, "stopMessageLiveLocation", params, &result)
	return result, err
}

func (b *Bot) EditMessageChecklist(ctx context.Context, params *models.EditMessageChecklistRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "editMessageChecklist", params, result)
	return result, err
}

func (b *Bot) EditMessageReplyMarkup(ctx context.Context, params *models.EditMessageReplyMarkupRequest) (interface{}, error) {
	var result interface{}

	err := request(b, ctx, "editMessageReplyMarkup", params, &result)
	return result, err
}

func (b *Bot) StopPoll(ctx context.Context, params *models.StopPollRequest) (*models.Poll, error) {
	result := new(models.Poll)

	err := request(b, ctx, "stopPoll", params, result)
	return result, err
}

func (b *Bot) ApproveSuggestedPost(ctx context.Context, params *models.ApproveSuggestedPostRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "approveSuggestedPost", params, &result)
	return result, err
}

func (b *Bot) DeclineSuggestedPost(ctx context.Context, params *models.DeclineSuggestedPostRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "declineSuggestedPost", params, &result)
	return result, err
}

func (b *Bot) DeleteMessage(ctx context.Context, params *models.DeleteMessageRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteMessage", params, &result)
	return result, err
}

func (b *Bot) DeleteMessages(ctx context.Context, params *models.DeleteMessagesRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteMessages", params, &result)
	return result, err
}

func (b *Bot) SendSticker(ctx context.Context, params *models.SendStickerRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendSticker", params, result)
	return result, err
}

func (b *Bot) GetStickerSet(ctx context.Context, params *models.GetStickerSetRequest) (*models.StickerSet, error) {
	result := new(models.StickerSet)

	err := request(b, ctx, "getStickerSet", params, result)
	return result, err
}

func (b *Bot) GetCustomEmojiStickers(ctx context.Context, params *models.GetCustomEmojiStickersRequest) ([]models.Sticker, error) {
	var result []models.Sticker

	err := request(b, ctx, "getCustomEmojiStickers", params, &result)
	return result, err
}

func (b *Bot) UploadStickerFile(ctx context.Context, params *models.UploadStickerFileRequest) (*models.File, error) {
	result := new(models.File)

	err := request(b, ctx, "uploadStickerFile", params, result)
	return result, err
}

func (b *Bot) CreateNewStickerSet(ctx context.Context, params *models.CreateNewStickerSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "createNewStickerSet", params, &result)
	return result, err
}

func (b *Bot) AddStickerToSet(ctx context.Context, params *models.AddStickerToSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "addStickerToSet", params, &result)
	return result, err
}

func (b *Bot) SetStickerPositionInSet(ctx context.Context, params *models.SetStickerPositionInSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setStickerPositionInSet", params, &result)
	return result, err
}

func (b *Bot) DeleteStickerFromSet(ctx context.Context, params *models.DeleteStickerFromSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteStickerFromSet", params, &result)
	return result, err
}

func (b *Bot) ReplaceStickerInSet(ctx context.Context, params *models.ReplaceStickerInSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "replaceStickerInSet", params, &result)
	return result, err
}

func (b *Bot) SetStickerEmojiList(ctx context.Context, params *models.SetStickerEmojiListRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setStickerEmojiList", params, &result)
	return result, err
}

func (b *Bot) SetStickerKeywords(ctx context.Context, params *models.SetStickerKeywordsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setStickerKeywords", params, &result)
	return result, err
}

func (b *Bot) SetStickerMaskPosition(ctx context.Context, params *models.SetStickerMaskPositionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setStickerMaskPosition", params, &result)
	return result, err
}

func (b *Bot) SetStickerSetTitle(ctx context.Context, params *models.SetStickerSetTitleRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setStickerSetTitle", params, &result)
	return result, err
}

func (b *Bot) SetStickerSetThumbnail(ctx context.Context, params *models.SetStickerSetThumbnailRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setStickerSetThumbnail", params, &result)
	return result, err
}

func (b *Bot) SetCustomEmojiStickerSetThumbnail(ctx context.Context, params *models.SetCustomEmojiStickerSetThumbnailRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setCustomEmojiStickerSetThumbnail", params, &result)
	return result, err
}

func (b *Bot) DeleteStickerSet(ctx context.Context, params *models.DeleteStickerSetRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "deleteStickerSet", params, &result)
	return result, err
}

func (b *Bot) AnswerInlineQuery(ctx context.Context, params *models.AnswerInlineQueryRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "answerInlineQuery", params, &result)
	return result, err
}

func (b *Bot) AnswerWebAppQuery(ctx context.Context, params *models.AnswerWebAppQueryRequest) (*models.SentWebAppMessage, error) {
	result := new(models.SentWebAppMessage)

	err := request(b, ctx, "answerWebAppQuery", params, result)
	return result, err
}

func (b *Bot) SavePreparedInlineMessage(ctx context.Context, params *models.SavePreparedInlineMessageRequest) (*models.PreparedInlineMessage, error) {
	result := new(models.PreparedInlineMessage)

	err := request(b, ctx, "savePreparedInlineMessage", params, result)
	return result, err
}

func (b *Bot) SendInvoice(ctx context.Context, params *models.SendInvoiceRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendInvoice", params, result)
	return result, err
}

func (b *Bot) CreateInvoiceLink(ctx context.Context, params *models.CreateInvoiceLinkRequest) (string, error) {
	var result string

	err := request(b, ctx, "createInvoiceLink", params, &result)
	return result, err
}

func (b *Bot) AnswerShippingQuery(ctx context.Context, params *models.AnswerShippingQueryRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "answerShippingQuery", params, &result)
	return result, err
}

func (b *Bot) AnswerPreCheckoutQuery(ctx context.Context, params *models.AnswerPreCheckoutQueryRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "answerPreCheckoutQuery", params, &result)
	return result, err
}

func (b *Bot) GetMyStarBalance(ctx context.Context) (*models.StarAmount, error) {
	result := new(models.StarAmount)

	err := request[any](b, ctx, "getMyStarBalance", nil, result)
	return result, err
}

func (b *Bot) GetStarTransactions(ctx context.Context, params *models.GetStarTransactionsRequest) (*models.StarTransactions, error) {
	result := new(models.StarTransactions)

	err := request(b, ctx, "getStarTransactions", params, result)
	return result, err
}

func (b *Bot) RefundStarPayment(ctx context.Context, params *models.RefundStarPaymentRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "refundStarPayment", params, &result)
	return result, err
}

func (b *Bot) EditUserStarSubscription(ctx context.Context, params *models.EditUserStarSubscriptionRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "editUserStarSubscription", params, &result)
	return result, err
}

func (b *Bot) SetPassportDataErrors(ctx context.Context, params *models.SetPassportDataErrorsRequest) (bool, error) {
	var result bool

	err := request(b, ctx, "setPassportDataErrors", params, &result)
	return result, err
}

func (b *Bot) SendGame(ctx context.Context, params *models.SendGameRequest) (*models.Message, error) {
	result := new(models.Message)

	err := request(b, ctx, "sendGame", params, result)
	return result, err
}

func (b *Bot) SetGameScore(ctx context.Context, params *models.SetGameScoreRequest) (interface{}, error) {
	var result interface{}

	err := request(b, ctx, "setGameScore", params, &result)
	return result, err
}

func (b *Bot) GetGameHighScores(ctx context.Context, params *models.GetGameHighScoresRequest) ([]models.GameHighScore, error) {
	var result []models.GameHighScore

	err := request(b, ctx, "getGameHighScores", params, &result)
	return result, err
}
