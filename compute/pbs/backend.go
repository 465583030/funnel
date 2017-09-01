package pbs

import (
	"github.com/ohsu-comp-bio/funnel/compute"
	"github.com/ohsu-comp-bio/funnel/config"
)

// NewBackend returns a new Grid Engine HPCBackend instance.
func NewBackend(conf config.Config) *compute.HPCBackend {
	return compute.NewHPCBackend("pbs", "qsub", conf, conf.Backends.PBS.Template)
}
