package keymaps

const (
	APIkey    = iota
	LatLong   = iota
	BaseURL   = iota
	LocalOnly = iota
)

var EnvKeyMap map[int]string = map[int]string{
	APIkey:    "apikey",
	LatLong:   "latlong",
	BaseURL:   "baseurl",
	LocalOnly: "localonly",
}
