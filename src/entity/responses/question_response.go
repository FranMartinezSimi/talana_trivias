package responses

type QuestionResponse struct {
	ID            uint   `json:"id"`
	Question      string `json:"question"`
	CorrectOption uint   `json:"correct_option"`
	Options       []struct {
		ID     uint   `json:"id"`
		Option string `json:"option"`
	} `json:"options"`
	Difficulty string `json:"difficulty"`
}
