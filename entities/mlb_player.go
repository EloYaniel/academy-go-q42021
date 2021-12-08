package entities

// MLBPlayer struct has MLB Player business info.
type MLBPlayer struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Team     string  `json:"team"`
	Position string  `json:"position"`
	Height   int     `json:"height_inches"`
	Weight   float32 `json:"weight_lbs"`
	Age      float32 `json:"age"`
}
