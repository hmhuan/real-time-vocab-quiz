package messages

type ParticipantsMessage struct {
	Base[ParticipantList]
}

type ParticipantList struct {
	Participants []Participant `json:"participants"`
}

type Participant struct {
	UserId string `json:"user_id"`
}

type JoinMessage struct {
	Base[Join]
}

type Join struct {
	UserId string `json:"user_id"`
}

type LeaveMessage struct {
	Base[Leave]
}

type Leave struct {
	UserId string `json:"user_id"`
}
