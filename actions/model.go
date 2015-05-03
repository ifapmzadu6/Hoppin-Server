package actions

type Actions struct {
	Actions []Action `json:"actions"`
}

type Action struct {
	VideoId string `json:"video_id"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Time    int    `json:"time"`
	Type    string `json:"type"`
}
