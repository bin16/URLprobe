package dotenv

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func init() {
	filepath := ".env"
	file, err := os.Open(filepath)
	if os.IsNotExist(err) {
		log.Printf("Warning: .env file not found, please check or rmeove the package.\n")
		return
	} else if err != nil {
		log.Printf("Warning: failed to read .env file, %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "=")
		key := items[0]
		values := strings.Join(items[1:], "=")
		os.Setenv(key, values)
	}
}
