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

import "fmt"

type Error struct {
	Code string `json:"code,omitempty"`
	Err  string `json:"error,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Err)
}

type City struct {
	Confidence int               `json:"confidence,omitempty"`
	GeoNameId  int               `json:"geoname_id,omitempty"`
	Names      map[string]string `json:"names,omitempty"`
}

type Continent struct {
	Code      string            `json:"code,omitempty"`
	GeoNameId int               `json:"geoname_id,omitempty"`
	Names     map[string]string `json:"names,omitempty"`
}

type Country struct {
	Confidence        int               `json:"confidence,omitempty"`
	GeoNameId         int               `json:"geoname_id,omitempty"`
	IsInEuropeanUnion bool              `json:"is_in_european_union,omitempty"`
	IsoCode           string            `json:"iso_code,omitempty"`
	Names             map[string]string `json:"names,omitempty"`
}

type Location struct {
	AccuracyRadius    int     `json:"accuracy_radius,omitempty"`
	AverageIncome     int     `json:"average_income,omitempty"`
	Latitude          float64 `json:"latitude,omitempty"`
	Longitude         float64 `json:"longitude,omitempty"`
	MetroCode         int     `json:"metro_code,omitempty"`
	PopulationDensity int     `json:"population_density,omitempty"`
	TimeZone          string  `json:"time_zone,omitempty"`
}

type Postal struct {
	Code       string `json:"code,omitempty"`
	Confidence int    `json:"confidence,omitempty"`
}

type RegisteredCountry struct {
	GeoNameId         int               `json:"geoname_id,omitempty"`
	IsInEuropeanUnion bool              `json:"is_in_european_union,omitempty"`
	IsoCode           string            `json:"iso_code,omitempty"`
	Names             map[string]string `json:"names,omitempty"`
}

type RepresentedCountry struct {
	GeoNameId         int               `json:"geoname_id,omitempty"`
	IsInEuropeanUnion bool              `json:"is_in_european_union,omitempty"`
	IsoCode           string            `json:"iso_code,omitempty"`
	Names             map[string]string `json:"names,omitempty"`
	Type              string            `json:"type,omitempty"`
}

type Subdivision struct {
	Confidence int               `json:"confidence,omitempty"`
	GeoNameId  int               `json:"geoname_id,omitempty"`
	IsoCode    string            `json:"iso_code,omitempty"`
	Names      map[string]string `json:"names,omitempty"`
}

type Traits struct {
	AutonomousSystemNumber       int     `json:"autonomous_system_number,omitempty"`
	AutonomousSystemOrganization string  `json:"autonomous_system_organization,omitempty"`
	Domain                       string  `json:"domain,omitempty"`
	IsAnonymous                  bool    `json:"is_anonymous,omitempty"`
	IsAnonymousProxy             bool    `json:"is_anonymous_proxy,omitempty"` // Deprecated
	IsAnonymousVpn               bool    `json:"is_anonymous_vpn,omitempty"`
	IsHostingProvider            bool    `json:"is_hosting_provider,omitempty"`
	IsPublicProxy                bool    `json:"is_public_proxy,omitempty"`
	IsResidentialProxy           bool    `json:"is_residential_proxy,omitempty"`
	IsSatelliteProvider          bool    `json:"is_satellite_provider,omitempty"` // Deprecated
	IsTorExitNode                bool    `json:"is_tor_exit_node,omitempty"`
	Isp                          string  `json:"isp,omitempty"`
	IpAddress                    string  `json:"ip_address,omitempty"`
	Network                      string  `json:"network,omitempty"`
	Organization                 string  `json:"organization,omitempty"`
	StaticIpScore                float64 `json:"static_ip_score,omitempty"`
	UserCount                    int     `json:"user_count,omitempty"`
	UserType                     string  `json:"user_type,omitempty"`
}

type MaxMind struct {
	QueriesRemaining int `json:"queries_remaining,omitempty"`
}

type Response struct {
	City               City               `json:"city,omitempty"`
	Continent          Continent          `json:"continent,omitempty"`
	Country            Country            `json:"country,omitempty"`
	Location           Location           `json:"location,omitempty"`
	Postal             Postal             `json:"postal,omitempty"`
	RegisteredCountry  RegisteredCountry  `json:"registered_country,omitempty"`
	RepresentedCountry RepresentedCountry `json:"represented_country,omitempty"`
	Subdivisions       []Subdivision      `json:"subdivisions,omitempty"`
	Traits             Traits             `json:"traits,omitempty"`
	MaxMind            MaxMind            `json:"maxmind,omitempty"`
}
