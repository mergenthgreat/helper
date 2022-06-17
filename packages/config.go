package packages

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetToken() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Yuklenirken Hata oldu")
	}
	return os.Getenv("TOKEN")
}

func GetMong() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Yuklenirken Hata oldu")
	}
	return os.Getenv("MONGO_URL")
}
