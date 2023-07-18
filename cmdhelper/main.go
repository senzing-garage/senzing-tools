package cmdhelper

import "github.com/senzing/senzing-tools/cmdhelper/optiontype"

// type ContextBool struct {
// 	Default bool   `json:"default"`
// 	Envar   string `json:"envar"`
// 	Help    string `json:"help"`
// 	Option  string `json:"option"`
// }

// type ContextInt struct {
// 	Default int    `json:"default"`
// 	Envar   string `json:"envar"`
// 	Help    string `json:"help"`
// 	Option  string `json:"option"`
// }

// type ContextString struct {
// 	Default string `json:"default"`
// 	Envar   string `json:"envar"`
// 	Help    string `json:"help"`
// 	Option  string `json:"option"`
// }

// type ContextStringSlice struct {
// 	Default []string `json:"default"`
// 	Envar   string   `json:"envar"`
// 	Help    string   `json:"help"`
// 	Option  string   `json:"option"`
// }

type ContextVariable struct {
	Default any                   `json:"default"`
	Envar   string                `json:"envar"`
	Help    string                `json:"help"`
	Arg     string                `json:"option"`
	Type    optiontype.OptionType `json:"optiontype"`
}

// type ContextVariables struct {
// 	Bools        []ContextBool        `json:"bools"`
// 	Ints         []ContextInt         `json:"ints"`
// 	Strings      []ContextString      `json:"strings"`
// 	StringSlices []ContextStringSlice `json:"stringSlices"`
// }
