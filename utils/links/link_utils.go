package links

import (
	"fmt"

	"github.com/MorrisMorrison/retfig/infrastructure/config"
)

const (
	API_PREFIX  = "api"
	API_TYPE    = "htmx"
	API_VERSION = "v1"

	EVENTS_API       = "events"
	PRESENTS_API     = "presents"
	COMMENTS_API     = "comments"
	PARTICIPANTS_API = "participants"
	VOTES_API        = "votes"
)

func BuildBaseUrl() string {
	return fmt.Sprintf("%s/%s/%s", API_PREFIX, API_TYPE, config.CONFIG.ApiVersion)
}

func BuildInvitationLink(eventId string) string {
	return fmt.Sprintf("http://%s:%s/%s/%s/invitations", config.CONFIG.Host, config.CONFIG.Port, EVENTS_API, eventId)
}

func BuildGetEventLink(eventId string) string {
	return fmt.Sprintf("/%s/%s", EVENTS_API, eventId)
}

func BuildGetCommentsLink(eventId string, presentId string) string {
	return fmt.Sprintf("/%s/%s/%s/%s/%s/%s", BuildBaseUrl(), EVENTS_API, eventId, PRESENTS_API, presentId, COMMENTS_API)
}

func BuildGetInvitationLink(eventId string) string {
	return fmt.Sprintf("/%s/%s/%s/%s", BuildBaseUrl(), EVENTS_API, eventId, PARTICIPANTS_API)
}

func BuildCreateEventLink() string {
	return fmt.Sprintf("/%s/%s", BuildBaseUrl(), EVENTS_API)
}

func BuildCreatePresentLink(eventId string) string {
	return fmt.Sprintf("/%s/%s/%s/%s", BuildBaseUrl(), EVENTS_API, eventId, PRESENTS_API)
}

func BuildCreateCommentLink(eventId string, presentId string) string {
	return fmt.Sprintf("/%s/%s/%s/%s/%s/%s", BuildBaseUrl(), EVENTS_API, eventId, PRESENTS_API, presentId, COMMENTS_API)
}

func BuildCreateVoteLink(eventId string, presentId string) string {
	return fmt.Sprintf("/%s/%s/%s/%s/%s/%s", BuildBaseUrl(), EVENTS_API, eventId, PRESENTS_API, presentId, VOTES_API)
}

func BuildRootLink() string {
	return fmt.Sprintf("http://%s:%s/", config.CONFIG.Host, config.CONFIG.Port)
}
