package responses

type QuestionResponse struct {
	ID            uint             `json:"id"`
	Question      string           `json:"question"`
	CorrectOption uint             `json:"correct_option"`
	Options       []OptionResponse `json:"options"`
	Difficulty    string           `json:"difficulty"`
}
