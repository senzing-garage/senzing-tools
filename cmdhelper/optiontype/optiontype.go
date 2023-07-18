package optiontype

type OptionType uint

const (
	Bool OptionType = iota
	Int
	String
	StringSlice
)
