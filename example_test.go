package geoip2_test

import (
	"encoding/json"
	"os"
	"time"

	"code.google.com/p/go.net/context"
	"github.com/savaki/geoip2"
)

func ExampleCity() {
	api := geoip2.New(os.Getenv("MAXMIND_USER_ID"), os.Getenv("MAXMIND_LICENSE_KEY"))
	resp, _ := api.City(nil, "1.2.3.4")
	json.NewEncoder(os.Stdout).Encode(resp)
}

func ExampleContext() {
	api := geoip2.New(os.Getenv("MAXMIND_USER_ID"), os.Getenv("MAXMIND_LICENSE_KEY"))

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	resp, _ := api.City(ctx, "1.2.3.4")
	json.NewEncoder(os.Stdout).Encode(resp)
}
