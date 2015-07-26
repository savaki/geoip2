package geoip2_test

import (
	"encoding/json"
	"os"
	"time"

	"github.com/savaki/geoip2"
	"golang.org/x/net/context"
)

func ExampleCity() {
	userId := os.Getenv("MAXMIND_USER_ID")
	licenseKey := os.Getenv("MAXMIND_LICENSE_KEY")
	api := geoip2.New(userId, licenseKey)

	resp, _ := api.City(nil, "1.2.3.4")
	json.NewEncoder(os.Stdout).Encode(resp)
}

func ExampleCountry() {
	userId := os.Getenv("MAXMIND_USER_ID")
	licenseKey := os.Getenv("MAXMIND_LICENSE_KEY")
	api := geoip2.New(userId, licenseKey)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	resp, _ := api.Country(ctx, "1.2.3.4")
	json.NewEncoder(os.Stdout).Encode(resp)
}

func ExampleInsights() {
	userId := os.Getenv("MAXMIND_USER_ID")
	licenseKey := os.Getenv("MAXMIND_LICENSE_KEY")
	api := geoip2.New(userId, licenseKey)

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	resp, _ := api.Insights(ctx, "1.2.3.4")
	json.NewEncoder(os.Stdout).Encode(resp)
}
