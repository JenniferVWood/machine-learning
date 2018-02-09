package types

// Users of this project will have to make their own config file.
// Mine is not included, because it contains my openei api key.
//
// format:
//    {  "key": "your-openei-api-key" }
//
// subject to change.  Definitive doc is the struct below.
type Configuration struct {
	Key string
}
