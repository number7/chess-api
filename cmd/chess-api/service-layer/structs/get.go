package structs

type GetRequest struct {
	ID string
}

type GetResponse struct {
	ID        string
	Location  string
	Date      string
	Round     int
	WhiteName string
	BlackName string
	Moves     string
	Result    string
}
