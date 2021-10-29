package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var count = 1

func init() {
	flag.IntVar(&count, "c", 1, "number of names to generate")
}

func main() {
	flag.Parse()

	foods, err := loadFood()
	if err != nil {
		log.Fatal(err)
	}

	adjs, err := loadAdj()
	if err != nil {
		log.Fatal(err)
	}

	names := make([]string, count)

	for i := 0; i < count; i++ {
		n := fmt.Sprintf("%s %s", adjs.Random(), foods.Random())
		names[i] = strings.Title(n)
	}

	fmt.Println(strings.Join(names, "\n"))
}

type Foods []Food

func (f Foods) Random() string {

	idx := r.Intn(len(f))
	sel := f[idx]

	item := r.Intn(len(sel.FoodItems))
	return sel.FoodItems[item].FoodName
}

type Food struct {
	Restaurant string     `json:"restaurant"`
	FoodItems  []FoodItem `json:"foodItems"`
}

type FoodItem struct {
	FoodName     string  `json:"foodName"`
	FoodType     string  `json:"foodType,omitempty"`
	Protein      *string `json:"protein,omitempty"`
	Calories     int64   `json:"calories"`
	SideItem     *bool   `json:"sideItem,omitempty"`
	DressingItem *bool   `json:"dressingItem,omitempty"`
}

func loadFood() (Foods, error) {
	f, err := os.ReadFile("./foods.json")
	if err != nil {
		return nil, err
	}

	var foods Foods
	if err := json.Unmarshal(f, &foods); err != nil {
		return nil, err
	}

	return foods, nil
}

type Adjectives []string

func (a Adjectives) Random() string {
	aidx := r.Intn(len(a))
	return a[aidx]
}

func loadAdj() (Adjectives, error) {
	f, err := os.ReadFile("./adj.txt")
	if err != nil {
		return nil, err
	}

	return strings.Split(string(f), "\n"), nil
}
