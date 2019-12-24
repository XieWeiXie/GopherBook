package Apexcharts

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type point struct {
	X string `json:"x"`
	Y int64  `json:"y"`
}

func generateData(length int, min int64, max int64) []point {
	var result []point
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < length+1; i++ {
		var one point
		n := seed.Int63n(int64(max-min)) + int64(min)

		one = point{
			X: strconv.Itoa(i),
			Y: n,
		}
		result = append(result, one)
	}
	return result
}

type OneData struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func TestHeatMap(tests *testing.T) {
	heatMap := NewHeatMap("Example HeatMap")
	var result []interface{}
	for i := 0; i < 15; i++ {
		var one OneData
		one.Name = fmt.Sprintf("W%d", i)
		one.Data = generateData(8, 0, 90)
		result = append(result, one)
	}
	heatMap.Series = Series{
		Data: result,
	}
	heatMap.SetColors([]string{"#F3B415", "#F27036", "#663F59", "#6A6E94", "#4E88B4", "#00A7C6", "#18D8D8", "#A9D794", "#46AF78", "#A93F55", "#8C5E58", "#2176FF", "#33A1FD", "#7A918D", "#BAFF29"})
	heatMap.SetCategories([]string{"10:00", "10:30", "11:00", "11:30", "12:00", "12:30", "01:00", "01:30"})
	http.HandleFunc("/", heatMap.Plot)
	log.Println("http://127.0.0.1:7878")
	log.Fatal(http.ListenAndServe(":7878", nil))
}

func TestData(tests *testing.T) {
	for i := 0; i < 15; i++ {
		fmt.Println(generateData(8, 0, 90))
	}
}
