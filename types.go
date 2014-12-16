package playlyfe

type PlayersData struct {
	Data []Player
}
type Player struct {
	Id      string `json:"id,omitempty"`
	Alias   string `json:"alias,omitempty"`
	Email   string `json:"email,omitempty"`
	Enabled *bool  `json:"enabled,omitempty"`
}
