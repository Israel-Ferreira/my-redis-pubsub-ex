package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/my-redis-pubsub/pkg/config"
	"github.com/Israel-Ferreira/my-redis-pubsub/pkg/models"
)

func HandleCreateCarRequest(rw http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var carDto *models.CarDTO

	if err := json.NewDecoder(r.Body).Decode(&carDto); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	car := carDto.ToCar()

	rdb := config.ConfigRedis()

	json_msg, err := json.Marshal(car)

	if err != nil {
		log.Printf("Error on Marshaling msg: %s", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Sending message to redis pub/sub....")

	if err := rdb.Publish(context.Background(), "cars", json_msg).Err(); err != nil {
		log.Printf("Error on publish message: %s", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("message published sucessfully")

	rw.WriteHeader(http.StatusOK)
}
