package gomatrix

// RespError is the standard JSON error response from Homeservers. It also implements the Golang "error" interface.
// See http://matrix.org/docs/spec/client_server/r0.2.0.html#api-standards
type RespError struct {
	ErrCode string `json:"errcode"`
	Err     string `json:"error"`
}

// Error returns the errcode and error message.
func (e RespError) Error() string {
	return e.ErrCode + ": " + e.Err
}

// RespCreateFilter is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-user-userid-filter
type RespCreateFilter struct {
	FilterID string `json:"filter_id"`
}

// RespVersions is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-versions
type RespVersions struct {
	Versions []string `json:"versions"`
}

// RespPublicRooms is the JSON response for http://matrix.org/speculator/spec/HEAD/client_server/unstable.html#get-matrix-client-unstable-publicrooms
type RespPublicRooms struct {
	TotalRoomCountEstimate int          `json:"total_room_count_estimate"`
	PrevBatch              string       `json:"prev_batch"`
	NextBatch              string       `json:"next_batch"`
	Chunk                  []PublicRoom `json:"chunk"`
}

// PublicRooms is a part of RespPublicRooms representing the disclosed information regarding a public room
type PublicRoom struct {
	CanonicalAlias   string   `json:"canonical_alias"`
	Name             string   `json:"name"`
	WorldReadable    bool     `json:"world_readable"`
	Topic            string   `json:"topic"`
	NumJoinedMembers int      `json:"num_joined_members"`
	AvatarURL        string   `json:"avatar_url"`
	RoomID           string   `json:"room_id"`
	GuestCanJoin     bool     `json:"guest_can_join"`
	Aliases          []string `json:"aliases"`
}

// RespJoinRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-join
type RespJoinRoom struct {
	RoomID string `json:"room_id"`
}

// RespLeaveRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-leave
type RespLeaveRoom struct{}

// RespForgetRoom is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-forget
type RespForgetRoom struct{}

// RespInviteUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-invite
type RespInviteUser struct{}

// RespKickUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-kick
type RespKickUser struct{}

// RespBanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-ban
type RespBanUser struct{}

// RespUnbanUser is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-rooms-roomid-unban
type RespUnbanUser struct{}

// RespTyping is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-typing-userid
type RespTyping struct{}

// RespJoinedRooms is the JSON response for TODO-SPEC https://github.com/matrix-org/synapse/pull/1680
type RespJoinedRooms struct {
	JoinedRooms []string `json:"joined_rooms"`
}

// RespJoinedMembers is the JSON response for TODO-SPEC https://github.com/matrix-org/synapse/pull/1680
type RespJoinedMembers struct {
	Joined map[string]struct {
		DisplayName *string `json:"display_name"`
		AvatarURL   *string `json:"avatar_url"`
	} `json:"joined"`
}

// RespMessages is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-rooms-roomid-messages
type RespMessages struct {
	Start string  `json:"start"`
	Chunk []Event `json:"chunk"`
	End   string  `json:"end"`
}

// RespSendEvent is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#put-matrix-client-r0-rooms-roomid-send-eventtype-txnid
type RespSendEvent struct {
	EventID string `json:"event_id"`
}

// RespMediaUpload is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-media-r0-upload
type RespMediaUpload struct {
	ContentURI string `json:"content_uri"`
}

// RespUserInteractive is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#user-interactive-authentication-api
type RespUserInteractive struct {
	Flows []struct {
		Stages []string `json:"stages"`
	} `json:"flows"`
	Params    map[string]interface{} `json:"params"`
	Session   string                 `json:"session"`
	Completed []string               `json:"completed"`
	ErrCode   string                 `json:"errcode"`
	Error     string                 `json:"error"`
}

// HasSingleStageFlow returns true if there exists at least 1 Flow with a single stage of stageName.
func (r RespUserInteractive) HasSingleStageFlow(stageName string) bool {
	for _, f := range r.Flows {
		if len(f.Stages) == 1 && f.Stages[0] == stageName {
			return true
		}
	}
	return false
}

// RespUserDisplayName is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-profile-userid-displayname
type RespUserDisplayName struct {
	DisplayName string `json:"displayname"`
}

// RespRegister is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-register
type RespRegister struct {
	AccessToken  string `json:"access_token"`
	DeviceID     string `json:"device_id"`
	HomeServer   string `json:"home_server"`
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
}

// RespLogin is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-login
type RespLogin struct {
	AccessToken string `json:"access_token"`
	DeviceID    string `json:"device_id"`
	HomeServer  string `json:"home_server"`
	UserID      string `json:"user_id"`
}

// RespLogout is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-logout
type RespLogout struct{}

// RespCreateRoom is the JSON response for https://matrix.org/docs/spec/client_server/r0.2.0.html#post-matrix-client-r0-createroom
type RespCreateRoom struct {
	RoomID string `json:"room_id"`
}

// RespSync is the JSON response for http://matrix.org/docs/spec/client_server/r0.2.0.html#get-matrix-client-r0-sync
type RespSync struct {
	NextBatch   string `json:"next_batch"`
	AccountData struct {
		Events []Event `json:"events"`
	} `json:"account_data"`
	Presence struct {
		Events []Event `json:"events"`
	} `json:"presence"`
	Rooms struct {
		Leave map[string]struct {
			State struct {
				Events []Event `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []Event `json:"events"`
				Limited   bool    `json:"limited"`
				PrevBatch string  `json:"prev_batch"`
			} `json:"timeline"`
		} `json:"leave"`
		Join map[string]struct {
			State struct {
				Events []Event `json:"events"`
			} `json:"state"`
			Timeline struct {
				Events    []Event `json:"events"`
				Limited   bool    `json:"limited"`
				PrevBatch string  `json:"prev_batch"`
			} `json:"timeline"`
		} `json:"join"`
		Invite map[string]struct {
			State struct {
				Events []Event
			} `json:"invite_state"`
		} `json:"invite"`
	} `json:"rooms"`
}

// RespTurnServer is the JSON response from a Turn Server
type RespTurnServer struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	TTL      int      `json:"ttl"`
	URIs     []string `json:"uris"`
}
