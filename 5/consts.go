package main

import "regexp"

const (
	SEEDS = iota
	SEED_TO_SOIL
	SOIL_TO_FERTILIZER
	FERTILIZER_TO_WATER
	WATER_TO_LIGHT
	LIGHT_TO_TEMPERATURE
	TEMPERATURE_TO_HUMIDITY
	HUMIDITY_TO_LOCATION
)

var STATE_TRANSITION = map[int]int{
	SEEDS:                   SEED_TO_SOIL,
	SEED_TO_SOIL:            SOIL_TO_FERTILIZER,
	SOIL_TO_FERTILIZER:      FERTILIZER_TO_WATER,
	FERTILIZER_TO_WATER:     WATER_TO_LIGHT,
	WATER_TO_LIGHT:          LIGHT_TO_TEMPERATURE,
	LIGHT_TO_TEMPERATURE:    TEMPERATURE_TO_HUMIDITY,
	TEMPERATURE_TO_HUMIDITY: HUMIDITY_TO_LOCATION,
}

var STATE_TO_PREFIX = map[int]string{
	SEEDS:                   "seed-to-soil",
	SEED_TO_SOIL:            "soil-to-fertilizer",
	SOIL_TO_FERTILIZER:      "fertilizer-to-water",
	FERTILIZER_TO_WATER:     "water-to-light",
	WATER_TO_LIGHT:          "light-to-temperature",
	LIGHT_TO_TEMPERATURE:    "temperature-to-humidity",
	TEMPERATURE_TO_HUMIDITY: "humidity-to-location",
}

var r = regexp.MustCompile("[0-9]+")
