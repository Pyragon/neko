package websocket

import (
	"fmt"
	"regexp"
	"time"

	"n.eko.moe/neko/internal/types"
	"n.eko.moe/neko/internal/types/event"
	"n.eko.moe/neko/internal/types/message"
	"n.eko.moe/neko/internal/utils"
)

func (h *MessageHandler) chat(id string, session types.Session, payload *message.ChatReceive) error {
	if session.Muted() {
		return nil
	}

	content := payload.Content

	content = h.censorChat(content, session)

	if content == "" {
		return nil
	}

	//is session.lastMessage less than 1 second ago?
	currentMillis := time.Now().UnixNano() / int64(time.Millisecond)

	if session.GetLastMessage() > currentMillis-1000 {
		return nil
	}

	session.SetLastMessage(currentMillis)

	chatID, err := utils.NewUID(32)

	if err != nil {
		h.logger.Info().Msg("Error creating new ID: " + err.Error())
		return nil
	}

	chatMessage := &types.ChatMessage{
		ID:      chatID,
		Author:  session.Name(),
		Content: content,
	}

	fmt.Println(chatMessage)

	if err := h.sessions.Broadcast(
		message.ChatSend{
			Event:   event.CHAT_MESSAGE,
			Name:    session.Name(),
			Content: content,
			ChatID:  chatID,
		}, nil); err != nil {
		h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.CONTROL_RELEASE)
		return err
	}
	return nil
}

var CENSORED = []*types.CensorType{
	{
		Regex: []string{
			"https?://",
			".(com)|(ca)|(co.uk)",
		},
		Replace: nil,
	},
}

func (h *MessageHandler) censorChat(content string, session types.Session) string {

	if session.GetRights() == 2 {
		return content
	}

	for _, censor := range CENSORED {
		for _, regex := range censor.Regex {
			match, err := regexp.MatchString(regex, content)

			if err != nil {
				return ""
			}

			if !match {
				continue
			}

			if censor.Replace == nil {
				return ""
			}

			re := regexp.MustCompile(regex)

			content = re.ReplaceAllString(regex, content)
		}
	}
	return content
}

func (h *MessageHandler) chatEmote(id string, session types.Session, payload *message.EmoteReceive) error {
	if session.Muted() {
		return nil
	}

	if err := h.sessions.Broadcast(
		message.EmoteSend{
			Event: event.CHAT_EMOTE,
			Emote: payload.Emote,
			ID:    id,
		}, nil); err != nil {
		h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.CONTROL_RELEASE)
		return err
	}
	return nil
}
