// Code generated by gotestmd DO NOT EDIT.
package single_cluster

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
)

type Suite struct {
	base.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/spire/single_cluster")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete crd spiffeids.spiffeid.spiffe.io` + "\n" + `kubectl delete ns spire`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/spire/single_cluster/?ref=d70bb72458b27706a7d13b08ca9a5fac57b35d0f`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-server`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-agent`)
}
func (s *Suite) Test() {}
