package websocket

import (
	"regexp"

	"n.eko.moe/neko/internal/types"
	"n.eko.moe/neko/internal/types/event"
	"n.eko.moe/neko/internal/types/message"
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

	if err := h.sessions.Broadcast(
		message.ChatSend{
			Event:   event.CHAT_MESSAGE,
			Content: content,
			Name:    session.Name(),
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
			".(com)|(ca)|(co.uk)"
		},
		Replace: nil,
	},
}

func (h *MessageHandler) censorChat(content string, session types.Session) string {

	for _, censor := range CENSORED {
		for _, regex := range censor.Regex {
			match, err := regexp.MatchString(regex, content)

			if err != nil || !match {
				return ""
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
