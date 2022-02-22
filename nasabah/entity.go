package nasabah

import "time"

type Nasabah struct {
	IdNasabahPerorangan int
	DEntry              time.Time
	IsBlackList         bool
	NamaNasabah         string
	NoIdentitas         string
	NamaIdentas         string
	AlamatIdentitas     string
}
