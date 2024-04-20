package views

import "encoding/json"

func GetCreateVoteJsonPayload(eventId string, presentId string, voteType string, username string) string {
	payload, err := json.Marshal(map[string]interface{}{
		"eventId":   eventId,
		"presentId": presentId,
		"voteType":  voteType,
		"username":  username,
	})

	if err != nil {
		return ""
	}

	return string(payload)
}
