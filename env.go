package env

import (
	"os"
	"strings"
	"sync"
)

type getOption func(key string) string

var (
	mu   sync.Mutex
	vars map[string]string
)

func init() {
	all := os.Environ()
	vars = make(map[string]string, len(all))

	mu.Lock()
	defer mu.Unlock()
	for _, v := range all {
		parts := strings.Split(v, "=")
		key := parts[0]
		value := parts[1]
		vars[key] = value
	}
}

func Get(env string, opts ...getOption) string {
	mu.Lock()
	defer mu.Unlock()

	v := vars[env]
	for _, opt := range opts {
		v = opt(env)
	}
	return v
}

func Set(env, value string) {
	mu.Lock()
	defer mu.Unlock()

	vars[env] = value
}

func Default(def string) getOption {
	return Transform(func(v string, found bool) string {
		if !found {
			return def
		}
		return v
	})
}

func Transform(fn func(string, bool) string) getOption {
	return func(key string) string {
		v, found := vars[key]
		return fn(v, found)
	}
}
