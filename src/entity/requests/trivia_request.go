package requests

type CreateTriviaRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	QuestionIDs []uint `json:"question_ids"`
	UserIDs     []uint `json:"user_ids"`
}

type SubmitAnswersRequest struct {
	UserID    uint            `json:"user_id"`
	Responses []AnswerRequest `json:"responses"`
}

type AnswerRequest struct {
	QuestionID     uint `json:"question_id"`
	SelectedOption uint `json:"selected_option"`
}
