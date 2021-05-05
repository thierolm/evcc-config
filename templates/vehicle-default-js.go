package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "vehicle",
		Type:   "custom",
		Name:   "Generic EV without SoC (Javascript)",
		Sample: `title: My electric vehicle # display name for UI
capacity: 10 # kWh
charge:
  source: js
  script: 95 // vehicle SoC in %`,
	}

	registry.Add(template)
}
