package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/heyyakash/realtime-weather-aggregator/configs"
	"github.com/heyyakash/realtime-weather-aggregator/modals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetHistoricalData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	city := r.URL.Query().Get("city")
	var res []modals.WeatherEvent
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "dt", Value: 1}})
	cur, err := configs.WeatherDataCollection.Find(context.TODO(), bson.M{"city": city}, findOptions)
	if err != nil {
		log.Print(err)
		return
	}
	for cur.Next(context.TODO()) {
		var data modals.WeatherEvent
		if err := cur.Decode(&data); err != nil {
			log.Print(err)
			return
		}
		res = append(res, data)
	}
	resMarshal, err := json.Marshal(res)
	w.Write(resMarshal)
}
