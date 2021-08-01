package keymaps

const (
	APIkey  = iota
	LatLong = iota
	BaseURL = iota
)

var EnvKeyMap map[int]string = map[int]string{
	APIkey:  "apikey",
	LatLong: "latlong",
	BaseURL: "baseurl",
}
