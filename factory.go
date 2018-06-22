package blade

import "github.com/fatrbaby/go-blade/cache"

func New(viewPath string, cache cache.Driver) *Blade {
	blade := new(Blade)
	blade.ViewPath = viewPath
	blade.cache = cache
	blade.bootstrap()

	return blade
}
