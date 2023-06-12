package cmdhelper

type ContextBool struct {
	Default bool   `json:"default"`
	Envar   string `json:"envar"`
	Help    string `json:"help"`
	Option  string `json:"option"`
}

type ContextInt struct {
	Default int    `json:"default"`
	Envar   string `json:"envar"`
	Help    string `json:"help"`
	Option  string `json:"option"`
}

type ContextString struct {
	Default string `json:"default"`
	Envar   string `json:"envar"`
	Help    string `json:"help"`
	Option  string `json:"option"`
}

type ContextStringSlice struct {
	Default []string `json:"default"`
	Envar   string   `json:"envar"`
	Help    string   `json:"help"`
	Option  string   `json:"option"`
}

type ContextVariables struct {
	Bools        []ContextBool        `json:"bools"`
	Ints         []ContextInt         `json:"ints"`
	Strings      []ContextString      `json:"strings"`
	StringSlices []ContextStringSlice `json:"stringSlices"`
}
