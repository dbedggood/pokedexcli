package main

type result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreas struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []result `json:"results"`
}

type LocationAreaDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod result `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int    `json:"rate"`
			Version result `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int    `json:"game_index"`
	ID        int    `json:"id"`
	Location  result `json:"location"`
	Name      string `json:"name"`
	Names     []struct {
		Language result `json:"language"`
		Name     string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        result `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int    `json:"chance"`
				ConditionValues []any  `json:"condition_values"`
				MaxLevel        int    `json:"max_level"`
				Method          result `json:"method"`
				MinLevel        int    `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int    `json:"max_chance"`
			Version   result `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
