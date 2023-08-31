package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func setKeyHandler(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса на установку ключа
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Прочитать JSON из тела запроса
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	key, val := data["key"], data["val"]
	if key == "" || val == "" {
		http.Error(w, "Missing key or value", http.StatusBadRequest)
		return
	}

	// Ваш код для записи значения в Redis
	fmt.Fprintf(w, "Key %s set to value %s", key, val)
}

// Обработка запроса на получение значения по ключу
func getKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key parameter", http.StatusBadRequest)
		return
	}

	// Получения значения из Redis
	// Если ключ не существует, вернуть 404
	val, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving value", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Value for key %s: %s", key, val)
}

// Обработка запроса на удаление пары ключ-значение
func delKeyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	key := data["key"]
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	// Удаления ключа из Redis
	err = redisClient.Del(key).Err()
	if err != nil {
		http.Error(w, "Error deleting key", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Key %s deleted", key)
}

var redisClient *redis.Client

func main() {
	redisAddr := "redis:6379"

	// Подключение к Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	http.HandleFunc("/set_key", setKeyHandler)
	http.HandleFunc("/get_key", getKeyHandler)
	http.HandleFunc("/del_key", delKeyHandler)

	port := ":8089"
	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
