package configure

import (
	log "github.com/kataras/golog"
)

func Logging() {
	log.SetTimeFormat("[15:04:05]")
	// Level defaults to "info",
	// but you can change it:
	log.SetLevel("debug")
}
