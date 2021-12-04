//	Copyright 2015 Matt Ho
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.

package geoip2

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"net/http"

	"fmt"

	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
)

var sample = `
{
  "city":  {
    "confidence":  25,
    "geoname_id": 54321,
    "names":  {
      "de":    "Los Angeles",
      "en":    "Los Angeles",
      "es":    "Los Ángeles",
      "fr":    "Los Angeles",
      "ja":    "ロサンゼルス市",
      "pt-BR":  "Los Angeles",
      "ru":    "Лос-Анджелес",
      "zh-CN": "洛杉矶"
    }
  },
  "continent":  {
    "code":       "NA",
    "geoname_id": 123456,
    "names":  {
      "de":    "Nordamerika",
      "en":    "North America",
      "es":    "América del Norte",
      "fr":    "Amérique du Nord",
      "ja":    "北アメリカ",
      "pt-BR": "América do Norte",
      "ru":    "Северная Америка",
      "zh-CN": "北美洲"
    }
  },
  "country":  {
    "confidence":  75,
    "geoname_id":  6252001,
    "is_in_european_union": false,
    "iso_code":    "US",
    "names":  {
      "de":     "USA",
      "en":     "United States",
      "es":     "Estados Unidos",
      "fr":     "États-Unis",
      "ja":     "アメリカ合衆国",
      "pt-BR":  "Estados Unidos",
      "ru":     "США",
      "zh-CN":  "美国"
    }
  },
  "location":  {
    "accuracy_radius":     20,
    "average_income":      128321,
    "latitude":            37.6293,
    "longitude":           -122.1163,
    "metro_code":          807,
    "population_density":  7122,
    "time_zone":           "America/Los_Angeles"
  },
  "postal": {
    "code":       "90001",
    "confidence": 10
  },
  "registered_country":  {
    "geoname_id":  6252001,
    "is_in_european_union": false,
    "iso_code":    "US",
    "names":  {
      "de":     "USA",
      "en":     "United States",
      "es":     "Estados Unidos",
      "fr":     "États-Unis",
      "ja":     "アメリカ合衆国",
      "pt-BR":  "Estados Unidos",
      "ru":     "США",
      "zh-CN":  "美国"
    }
  },
  "represented_country":  {
    "geoname_id":  6252001,
    "is_in_european_union": false,
    "iso_code":    "US",
    "names":  {
      "de":     "USA",
      "en":     "United States",
      "es":     "Estados Unidos",
      "fr":     "États-Unis",
      "ja":     "アメリカ合衆国",
      "pt-BR":  "Estados Unidos",
      "ru":     "США",
      "zh-CN":  "美国"
    },
    "type": "military"
  },
  "subdivisions":  [
    {
      "confidence":  50,
      "geoname_id":  5332921,
      "iso_code":    "CA",
      "names":  {
        "de":    "Kalifornien",
        "en":    "California",
        "es":    "California",
        "fr":    "Californie",
        "ja":    "カリフォルニア",
        "ru":    "Калифорния",
        "zh-CN": "加州"
      }
    }
  ],
  "traits": {
    "ip_address":                     "1.2.3.4",
    "network":                        "1.2.3.0/24",
    "autonomous_system_number":       1239,
    "autonomous_system_organization": "Linkem IR WiMax Network",
    "domain":                         "example.com",
    "isp":                            "Linkem spa",
    "organization":                   "Linkem IR WiMax Network",
    "is_anonymous":                   true,
    "is_anonymous_proxy":             true,
    "is_anonymous_vpn":               true,
    "is_hosting_provider":            true,
    "is_public_proxy":                true,
    "is_residential_proxy":           true,
    "is_satellite_provider":          true,
    "is_tor_exit_node":               true,
    "static_ip_score":                1.5,
    "user_count":                     1,
    "user_type":                      "traveler"
  },
  "maxmind": {
    "queries_remaining": 54321
  }
}`

func TestDecodesJson(t *testing.T) {
	Convey("Given a complete maxmind response", t, func() {

		resp := Response{}
		err := json.NewDecoder(strings.NewReader(sample)).Decode(&resp)
		So(err, ShouldBeNil)

		data, err := json.Marshal(resp)
		So(err, ShouldBeNil)
		text := string(data)

		So(text, ShouldNotContainSubstring, `""`)

		So(normalize(string(data)), ShouldEqual, normalize(string(data)))
	})
}

func TestApi(t *testing.T) {
	Convey("Given an Api client", t, func() {
		api := New("blah-user-id", "blah-license-key")

		Convey("When I make a query that returns a valid result", func() {
			doFunc := func(context.Context, *http.Request) (*http.Response, error) {
				resp := &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(sample)),
				}
				return resp, nil
			}

			Convey("When I call #Country", func() {
				api = WithClientFunc(api, doFunc)
				resp, err := api.Country(nil, "1.2.3.4")

				Convey("I expect no errors", func() {
					So(err, ShouldBeNil)
					So(resp.City.Confidence, ShouldEqual, 25)
				})
			})

			Convey("When I call #City", func() {
				api = WithClientFunc(api, doFunc)
				resp, err := api.City(nil, "1.2.3.4")

				Convey("I expect no errors", func() {
					So(err, ShouldBeNil)
					So(resp.City.Confidence, ShouldEqual, 25)
				})
			})

			Convey("When I call #Insights", func() {
				api = WithClientFunc(api, doFunc)
				resp, err := api.Insights(nil, "1.2.3.4")

				Convey("I expect no errors", func() {
					So(err, ShouldBeNil)
					So(resp.City.Confidence, ShouldEqual, 25)
				})

				Convey("I expect Traits to be properly populated", func() {
					So(err, ShouldBeNil)
					So(resp.Traits.IpAddress, ShouldEqual, "1.2.3.4")
					So(resp.Traits.Network, ShouldEqual, "1.2.3.0/24")
					So(resp.Traits.AutonomousSystemNumber, ShouldEqual, 1239)
					So(resp.Traits.AutonomousSystemOrganization, ShouldEqual, "Linkem IR WiMax Network")
					So(resp.Traits.Domain, ShouldEqual, "example.com")
					So(resp.Traits.Isp, ShouldEqual, "Linkem spa")
					So(resp.Traits.Organization, ShouldEqual, "Linkem IR WiMax Network")
					So(resp.Traits.IsAnonymous, ShouldEqual, true)
					So(resp.Traits.IsAnonymousProxy, ShouldEqual, true)
					So(resp.Traits.IsAnonymousVpn, ShouldEqual, true)
					So(resp.Traits.IsHostingProvider, ShouldEqual, true)
					So(resp.Traits.IsPublicProxy, ShouldEqual, true)
					So(resp.Traits.IsResidentialProxy, ShouldEqual, true)
					So(resp.Traits.IsSatelliteProvider, ShouldEqual, true)
					So(resp.Traits.IsTorExitNode, ShouldEqual, true)
					So(resp.Traits.StaticIpScore, ShouldEqual, 1.5)
					So(resp.Traits.UserCount, ShouldEqual, 1)
					So(resp.Traits.UserType, ShouldEqual, "traveler")
				})

			})
		})

		Convey("When I make a query that returns an invalid result", func() {
			code := "IP_ADDRESS_REQUIRED"
			message := "You have not supplied an IP address, which is a required field."
			doFunc := func(context.Context, *http.Request) (*http.Response, error) {
				layout := `
                {
                    "code":"%s",
                    "error": "%s"
                }`
				text := fmt.Sprintf(layout, code, message)
				resp := &http.Response{
					StatusCode: 400,
					Body:       ioutil.NopCloser(strings.NewReader(text)),
				}
				return resp, nil
			}
			api = WithClientFunc(api, doFunc)
			_, err := api.City(nil, "1.2.3.4")

			Convey("I expect no errors", func() {
				So(err, ShouldNotBeNil)

				e := err.(Error)
				So(e.Code, ShouldEqual, code)
				So(e.Err, ShouldEqual, message)
			})
		})
	})
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	return s
}
