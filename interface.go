package i18n

// GinI18n ...
type GinI18n interface {
	GetMessage(param interface{}, lng string) (string, error)
	MustGetMessage(param interface{}, lng string) string
	setBundle(cfg *BundleCfg)
}
