package rogu

type Options struct {
	EnableDate bool
	EnableLog bool
	EnableError bool
	EnableWarning bool
	EnableStack bool
}

func DefaultOptions() Options {
	options := Options{}
	options.EnableDate = true
	options.EnableLog = true
	options.EnableError = true
	options.EnableWarning = true
	options.EnableStack = true
	return options
}

func ProductionOptions() Options {
	options := Options{}
	options.EnableDate = true
	options.EnableError = true
	options.EnableWarning = true
	return options
}

func ProductionOptionsWithStack() Options {
	options := Options{}
	options.EnableDate = true
	options.EnableError = true
	options.EnableWarning = true
	options.EnableStack = true
	return options
}