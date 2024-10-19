package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/heyyakash/realtime-weather-aggregator/configs"
	"github.com/heyyakash/realtime-weather-aggregator/modals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAggregateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	city := r.URL.Query().Get("city")
	res, err := GetDailyStats(time.Now(), city)
	if err != nil {
		log.Print(err)
		w.Write([]byte("error"))
		return
	}

	go func(stat *modals.DailyStats) {
		if err := UpdateDailyAggregate(stat); err != nil {
			log.Print(err)
		}
	}(res)

	rm, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
		w.Write([]byte("error"))
		return
	}

	w.Write(rm)
}

func UpdateDailyAggregate(stat *modals.DailyStats) error {
	filter := bson.M{
		"city": stat.City,
		"day":  stat.Day,
	}
	update := bson.M{
		"$set": stat,
	}
	opts := options.Update().SetUpsert(true)

	_, err := configs.CalculatedDataCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func GetDailyStats(date time.Time, city string) (*modals.DailyStats, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDayUnix := startOfDay.Add(24 * time.Hour).Unix()
	startOfDayUnix := startOfDay.Unix()

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{
			"dt": bson.M{
				"$gte": startOfDayUnix,
				"$lt":  endOfDayUnix,
			},
			"city": city,
		}}},
		{{Key: "$addFields", Value: bson.M{
			"temperatureConverted": bson.M{"$toDouble": "$temperature"},
		}}},
		{{Key: "$facet", Value: bson.M{
			"tempStats": bson.A{
				bson.M{
					"$group": bson.M{
						"_id":     nil,
						"avgTemp": bson.M{"$avg": "$temperatureConverted"},
						"maxTemp": bson.M{"$max": "$temperatureConverted"},
						"minTemp": bson.M{"$min": "$temperatureConverted"},
					},
				},
			},
			"weatherStats": bson.A{
				bson.M{
					"$group": bson.M{
						"_id":   "$description",
						"count": bson.M{"$sum": 1},
					},
				},
				bson.M{"$sort": bson.M{"count": -1}},
				bson.M{"$limit": 1},
			},
		}}},
		{{Key: "$project", Value: bson.M{
			"avgTemp":         bson.M{"$arrayElemAt": []interface{}{"$tempStats.avgTemp", 0}},
			"maxTemp":         bson.M{"$arrayElemAt": []interface{}{"$tempStats.maxTemp", 0}},
			"minTemp":         bson.M{"$arrayElemAt": []interface{}{"$tempStats.minTemp", 0}},
			"dominantWeather": bson.M{"$arrayElemAt": []interface{}{"$weatherStats._id", 0}},
		}}},
	}

	cursor, err := configs.WeatherDataCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []modals.DailyStats
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return &modals.DailyStats{}, nil
	}

	results[0].Day = startOfDay
	results[0].City = city
	return &results[0], nil
}
