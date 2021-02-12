package websocket

import (
	"n.eko.moe/neko/internal/types"
	"n.eko.moe/neko/internal/types/event"
	"n.eko.moe/neko/internal/types/message"
)

func (h *MessageHandler) controlRelease(id string, session types.Session) error {

	// check if session is host
	if !h.sessions.IsHost(session.Name()) {
		h.logger.Debug().Str("id", id).Msg("is not the host")
		return nil
	}

	// release host
	h.logger.Debug().Str("id", id).Msgf("host called %s", event.CONTROL_RELEASE)
	h.sessions.ClearHost()

	// tell everyone
	if err := h.sessions.Broadcast(
		message.Control{
			Event: event.CONTROL_RELEASE,
			Name:  session.Name(),
		}, nil); err != nil {
		h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.CONTROL_RELEASE)
		return err
	}

	return nil
}

func (h *MessageHandler) controlRequest(id string, session types.Session) error {

	if session.GetRights() != 2 {
		h.logger.Debug().Msg(session.Name() + " is not an admin")
		return nil
	}
	//Simply give host, and tell everyone controls have been taken
	// set host
	h.sessions.SetHost(session.Name())

	// let everyone know
	if err := h.sessions.Broadcast(
		message.Control{
			Event: event.CONTROL_LOCKED,
			Name:  session.Name(),
		}, nil); err != nil {
		h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.CONTROL_LOCKED)
		return err
	}

	return nil
}

func (h *MessageHandler) controlGive(id string, session types.Session, payload *message.Control) error {
	// check if session is host
	if !h.sessions.IsHost(session.Name()) {
		h.logger.Debug().Msg(session.Name() + " is not the host")
		return nil
	}

	if !h.sessions.Has(payload.Name) {
		h.logger.Debug().Msg(payload.Name + " user does not exist")
		return nil
	}

	// set host
	h.sessions.SetHost(payload.Name)

	// let everyone know
	if err := h.sessions.Broadcast(
		message.ControlTarget{
			Event:  event.CONTROL_GIVE,
			Name:   session.Name(),
			Target: payload.Name,
		}, nil); err != nil {
		h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.CONTROL_LOCKED)
		return err
	}

	return nil
}

func (h *MessageHandler) controlClipboard(id string, session types.Session, payload *message.Clipboard) error {
	// check if session is host
	if !h.sessions.IsHost(session.Name()) {
		h.logger.Debug().Str("id", id).Msg("is not the host")
		return nil
	}

	h.remote.WriteClipboard(payload.Text)
	return nil
}

func (h *MessageHandler) controlKeyboard(id string, session types.Session, payload *message.Keyboard) error {
	// check if session is host
	if !h.sessions.IsHost(session.Name()) {
		h.logger.Debug().Str("id", id).Msg("is not the host")
		return nil
	}

	// change layout
	if payload.Layout != nil {
		h.remote.SetKeyboardLayout(*payload.Layout)
	}

	// set num lock
	var NumLock = 0
	if payload.NumLock == nil {
		NumLock = -1
	} else if *payload.NumLock {
		NumLock = 1
	}

	// set caps lock
	var CapsLock = 0
	if payload.CapsLock == nil {
		CapsLock = -1
	} else if *payload.CapsLock {
		CapsLock = 1
	}

	// set scroll lock
	var ScrollLock = 0
	if payload.ScrollLock == nil {
		ScrollLock = -1
	} else if *payload.ScrollLock {
		ScrollLock = 1
	}

	h.logger.Debug().
		Int("NumLock", NumLock).
		Int("CapsLock", CapsLock).
		Int("ScrollLock", ScrollLock).
		Msg("setting keyboard modifiers")

	h.remote.SetKeyboardModifiers(NumLock, CapsLock, ScrollLock)
	return nil
}
