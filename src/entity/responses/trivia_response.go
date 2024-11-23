package responses

type TriviaResponse struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Questions   []QuestionResponse `json:"questions"`
	Users       []UserResponse     `json:"users"`
}

type PlayTriviaResponse struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Questions   []QuestionResponse `json:"questions"`
}

type SubmitAnswersResponse struct {
	TriviaID       uint `json:"trivia_id"`
	UserID         uint `json:"user_id"`
	CorrectAnswers int  `json:"correct_answers"`
	TotalQuestions int  `json:"total_questions"`
	Score          int  `json:"score"`
}

type UserScoreResponse struct {
	TriviaID       uint `json:"trivia_id"`
	UserID         uint `json:"user_id"`
	Score          int  `json:"score"`
	CorrectAnswers int  `json:"correct_answers"`
	TotalQuestions int  `json:"total_questions"`
}
