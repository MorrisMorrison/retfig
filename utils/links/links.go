package links

import "fmt"

func BuildInvitationLink(eventId string) string {
	return fmt.Sprintf("http://localhost:8080/events/%s/invitations", eventId)
}

func BuildGetEventLink(eventId string) string {
	return fmt.Sprintf("/events/%s", eventId)
}
