package conf

import "sync"

var (
	flagConfBaseDir *string = nil
	once                    = sync.Once{}
)

// Util, conf util
type Util struct{}

// SetBaseDir, set base dir
func (u *Util) SetBaseDir(baseDir string) {
	once.Do(func() {
		var dir = baseDir
		flagConfBaseDir = &dir
	})
}

// GetBaseDir, get base dir
func (u *Util) GetBaseDir() string {
	if flagConfBaseDir == nil {
		panic("You must set base directory first")
	}
	return *flagConfBaseDir
}
