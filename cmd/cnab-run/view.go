package main

import (
	"os"

	"github.com/docker/app/internal/packager"
	appview "github.com/docker/app/internal/view"
)

func viewAction(instanceName string) error {
	app, err := packager.Extract("")
	// todo: merge additional compose file
	if err != nil {
		return err
	}
	defer app.Cleanup()

	imageMap, err := getBundleImageMap()
	if err != nil {
		return err
	}

	parameters := packager.ExtractCNABParametersValues(packager.ExtractCNABParameterMapping(app.Parameters()), os.Environ())
	return appview.View(os.Stdout, app, parameters, imageMap)
}
