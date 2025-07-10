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

type Pokemon struct {
	BaseHappiness  int      `json:"base_happiness"`
	CaptureRate    int      `json:"capture_rate"`
	Color          result   `json:"color"`
	EggGroups      []result `json:"egg_groups"`
	EvolutionChain struct {
		URL string `json:"url"`
	} `json:"evolution_chain"`
	EvolvesFromSpecies result `json:"evolves_from_species"`
	FlavorTextEntries  []struct {
		FlavorText string `json:"flavor_text"`
		Language   result `json:"language"`
		Version    result `json:"version"`
	} `json:"flavor_text_entries"`
	FormDescriptions []interface{} `json:"form_descriptions"`
	FormsSwitchable  bool          `json:"forms_switchable"`
	GenderRate       int           `json:"gender_rate"`
	Genera           []struct {
		Genus    string `json:"genus"`
		Language result `json:"language"`
	} `json:"genera"`
	Generation           result `json:"generation"`
	GrowthRate           result `json:"growth_rate"`
	Habitat              result `json:"habitat"`
	HasGenderDifferences bool   `json:"has_gender_differences"`
	HatchCounter         int    `json:"hatch_counter"`
	ID                   int    `json:"id"`
	IsBaby               bool   `json:"is_baby"`
	IsLegendary          bool   `json:"is_legendary"`
	IsMythical           bool   `json:"is_mythical"`
	Name                 string `json:"name"`
	Names                []struct {
		Language result `json:"language"`
		Name     string `json:"name"`
	} `json:"names"`
	Order             int `json:"order"`
	PalParkEncounters []struct {
		Area      result `json:"area"`
		BaseScore int    `json:"base_score"`
		Rate      int    `json:"rate"`
	} `json:"pal_park_encounters"`
	PokedexNumbers []struct {
		EntryNumber int    `json:"entry_number"`
		Pokedex     result `json:"pokedex"`
	} `json:"pokedex_numbers"`
	Shape     result `json:"shape"`
	Varieties []struct {
		IsDefault bool   `json:"is_default"`
		Pokemon   result `json:"pokemon"`
	} `json:"varieties"`
}
