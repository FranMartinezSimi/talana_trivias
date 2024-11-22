package shared

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"talana_prueba_tecnica/src/entity/models"
)

var Env = GetEnvs()

func Init() *gorm.DB {
	var db *gorm.DB
	dsn := "host=" + Env["DB_HOST"] + " user=" + Env["DB_USER"] + " password=" + Env["DB_PASSWORD"] + " dbname=" + Env["DB_NAME"] + " port=" + Env["DB_PORT"] + " sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	migration(db)
	log.Println("Database connected")
	return db
}

func migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.UserModel{},
		&models.TriviaModel{},
		&models.QuestionModel{},
		&models.AnswerModel{},
		&models.ScoreModel{},
		&models.TriviaQuestionsModel{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database migrated")
}
