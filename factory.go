package blade

func NewBlade(viewPath, cachePath string) *Blade {
	blade := new(Blade)
	blade.viewPath = viewPath
	blade.cachePath = cachePath
	blade.bootstrap()

	return blade
}
