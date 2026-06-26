package main
import(
	"log"
	"movie-recommendation/internal/app"
	"movie-recommendation/internal/config"
	"movie-recommendation/internal/database"
)

func main(){
	cfg := config.LoadConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	router := app.New(db)
	log.Printf("Server running on port %s", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}

