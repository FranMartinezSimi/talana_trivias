package requests

type CreateQuestionRequest struct {
	Question      string   `json:"question"`
	Difficulty    string   `json:"difficulty"`
	Points        int      `json:"points"`
	Options       []string `json:"options"`
	CorrectOption int      `json:"correct_option"`
}
