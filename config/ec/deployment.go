package ec

import (
	"github.com/crossplane/terrajet/pkg/config"
)

const (
	deployment = "ec_deployment"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(deployment, func(r *config.Resource) {
		r.Kind = "ECDeployment"
	})
}
