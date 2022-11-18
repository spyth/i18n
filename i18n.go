package i18n

func Localize(opts ...Option) GinI18n {
	// init ins
	ins := &ginI18nImpl{}

	// set ins property from opts
	for _, opt := range opts {
		opt(ins)
	}

	// 	if bundle isn't constructed then assign it from default
	if ins.bundle == nil {
		ins.setBundle(defaultBundleConfig)
	}

	return ins
}
