package minecraft

import (
	"time"
)

//Player store information about players
type Player struct {
	name     string
	uuid     string
	deaths   int
	lastSeen time.Time
}
