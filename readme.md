### Prueba tecnica talana - Backend

## Descripción
Este proyecto es una API REST que permite realizar operaciones en un juego de trivias

## Autor
* Nombre: Juan Martinez
* correo: juanfmartinezsimi@outlook.com | jua.martinez@outlook.es
* telefono: +56 9 6787 2837

- Golang 1.23
- PostgreSQL
- gorm ORM
- docker

## Instalación

repositorio url: [Repositorio de github](https://github.com/FranMartinezSimi/talana_trivias)
rama para evaluacion: ***MAIN***


variables de entorno:

| Variable | valor | Descripción |
| ------ | ------ | ------ |
| DB_HOST | localhost | Host de la base de datos |
| DB_USER | trivia_user | Usuario de la base de datos |
| DB_PASSWORD | trivia_password | Contraseña de la base de datos |
| DB_NAME | trivia_db | Nombre de la base de datos |
| DB_PORT | 5432 | Puerto de la base de datos |
| PORT | 8080 | Puerto de la aplicación |

#### SQL para la creación de la base de datos
```sql
-- Create user_models table
CREATE TABLE user_models (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

-- Create trivia table
CREATE TABLE trivia (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

-- Create questions table
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    question VARCHAR(255) NOT NULL,
    correct_option INTEGER NOT NULL,
    difficulty VARCHAR(50) NOT NULL,
    points INTEGER NOT NULL
);

-- Create options table
CREATE TABLE options (
    id SERIAL PRIMARY KEY,
    text VARCHAR(255) NOT NULL,
    question_id INTEGER NOT NULL,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

-- Create trivia_users table (junction table between trivia and user_models)
CREATE TABLE trivia_users (
    trivia_id INTEGER NOT NULL,
    user_model_id INTEGER NOT NULL,
    PRIMARY KEY (trivia_id, user_model_id),
    FOREIGN KEY (trivia_id) REFERENCES trivia(id) ON DELETE CASCADE,
    FOREIGN KEY (user_model_id) REFERENCES user_models(id) ON DELETE CASCADE
);

-- Create trivia_questions table (junction table between trivia and questions)
CREATE TABLE trivia_questions (
    trivia_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    PRIMARY KEY (trivia_id, question_id),
    FOREIGN KEY (trivia_id) REFERENCES trivia(id) ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

-- Create participations table
CREATE TABLE participations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    trivia_id INTEGER NOT NULL,
    score INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES user_models(id) ON DELETE CASCADE,
    FOREIGN KEY (trivia_id) REFERENCES trivia(id) ON DELETE CASCADE
);

-- Create answers table
CREATE TABLE answers (
    id SERIAL PRIMARY KEY,
    participation_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    selected_option INTEGER NOT NULL,
    is_correct BOOLEAN NOT NULL,
    FOREIGN KEY (participation_id) REFERENCES participations(id) ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

-- Add indexes for better query performance
CREATE INDEX idx_user_models_email ON user_models(email);
CREATE INDEX idx_questions_difficulty ON questions(difficulty);
CREATE INDEX idx_trivia_questions_trivia ON trivia_questions(trivia_id);
CREATE INDEX idx_trivia_questions_question ON trivia_questions(question_id);
CREATE INDEX idx_answers_participation ON answers(participation_id);
CREATE INDEX idx_participations_user ON participations(user_id);
CREATE INDEX idx_participations_trivia ON participations(trivia_id);

-- Add Full Text Search index for questions
CREATE INDEX question_text_idx ON questions USING GIN (to_tsvector('english', question));

```

## Endpoints
### swagger url: [Swagger](http://localhost:8080/swagger/index.html#/Trivias/post_trivias)


## Instalacion con docker
```sh

docker-compose up --build
    
```

