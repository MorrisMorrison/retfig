package payload

import "encoding/json"

func GetCreateVoteJsonPayload(voteType string) string {
	payload, err := json.Marshal(map[string]interface{}{
		"voteType": voteType,
	})

	if err != nil {
		return ""
	}

	return string(payload)
}

func GetClaimPresentJsonPayload(username string) string {
	payload, err := json.Marshal(map[string]interface{}{
		"username": username,
	})

	if err != nil {
		return ""
	}

	return string(payload)
}
