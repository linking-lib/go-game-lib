package module

type (
	options struct {
		name     string              // module name
		nameFunc func(string) string // rename handler name
	}

	// Option used to customize handler
	Option func(options *options)
)

// WithName used to rename module name
func WithName(name string) Option {
	return func(opt *options) {
		opt.name = name
	}
}

// WithNameFunc override handler name by specific function
// such as: strings.ToUpper/strings.ToLower
func WithNameFunc(fn func(string) string) Option {
	return func(opt *options) {
		opt.nameFunc = fn
	}
}
