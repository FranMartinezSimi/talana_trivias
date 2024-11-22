CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       name VARCHAR(50) NOT NULL,
                       email VARCHAR(30) NOT NULL UNIQUE,
                       password VARCHAR(50) NOT NULL,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       deleted_at DATETIME DEFAULT NULL
);

CREATE TABLE trivias (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         title VARCHAR(50) NOT NULL,
                         enabled BOOLEAN DEFAULT TRUE,
                         user_id INTEGER NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         deleted_at DATETIME DEFAULT NULL,
                         FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE questions (
                           id INTEGER PRIMARY KEY AUTOINCREMENT,
                           question TEXT NOT NULL,
                           enabled BOOLEAN DEFAULT TRUE,
                           created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                           updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                           deleted_at DATETIME DEFAULT NULL
);

CREATE TABLE answers (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         value TEXT NOT NULL,
                         is_correct BOOLEAN DEFAULT FALSE,
                         question_id INTEGER NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         deleted_at DATETIME DEFAULT NULL,
                         FOREIGN KEY (question_id) REFERENCES questions(id)
);

CREATE TABLE trivia_questions (
                                  trivia_id INTEGER NOT NULL,
                                  question_id INTEGER NOT NULL,
                                  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                  PRIMARY KEY (trivia_id, question_id),
                                  FOREIGN KEY (trivia_id) REFERENCES trivias(id),
                                  FOREIGN KEY (question_id) REFERENCES questions(id)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_trivias_user_id ON trivias(user_id);
CREATE INDEX idx_answers_question_id ON answers(question_id);
CREATE INDEX idx_trivia_questions_trivia_id ON trivia_questions(trivia_id);
CREATE INDEX idx_trivia_questions_question_id ON trivia_questions(question_id);