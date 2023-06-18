package entity

type AmiiboDatabase struct {
	Amiibos    map[string]Amiibo `json:"amiibos"`
	Characters map[string]string `json:"characters"`
	GameSeries map[string]string `json:"game_series"`
	Types      map[string]string `json:"types"`
}
