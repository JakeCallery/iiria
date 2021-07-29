package keymaps

const (
	APIkey = iota
)

var EnvKeyMap map[int]string = map[int]string{
	APIkey: "apikey",
}
