package injector

type Options struct {
	Alias string
	Scope string
}

type Option func(*Options)

func WithAlias(alias string) Option {
	return func(o *Options) {
		o.Alias = alias
	}
}

func WithScope(scope string) Option {
	return func(o *Options) {
		o.Scope = scope
	}
}
