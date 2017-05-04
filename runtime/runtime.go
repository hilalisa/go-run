// Package runtime provides high level functions for runtimes
package runtime

import (
	"github.com/go-log/log/log"
	"github.com/micro/go-run"
)

var (
	// should this even be here?
	Logger = log.New()
)

// Run manages the lifecycle of an application.
// This is a blocking call until the process exits.
func Run(r run.Runtime, url string) error {
	// get the source
	Logger.Logf("fetching %s\n", url)
	src, err := r.Fetch(url)
	if err != nil {
		return err
	}

	// build the binary
	Logger.Logf("building %s\n", url)
	bin, err := r.Build(src)
	if err != nil {
		return err
	}

	// execute the binary
	Logger.Logf("executing %s\n", url)
	proc, err := r.Exec(bin)
	if err != nil {
		return err
	}

	// wait till exit
	Logger.Logf("running %s\n", url)
	return r.Wait(proc)
}
