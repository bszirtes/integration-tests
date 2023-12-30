// Code generated by gotestmd DO NOT EDIT.
package jaeger_and_prometheus

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/observability/nsm_system"
)

type Suite struct {
	base.Suite
	nsm_systemSuite nsm_system.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.nsm_systemSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/observability/jaeger_and_prometheus")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-jaeger-and-prometheus`)
		r.Run(`kubectl describe pods -n observability` + "\n" + `kubectl delete ns observability` + "\n" + `pkill -f "port-forward"`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/observability/jaeger_and_prometheus?ref=27125b7983e1a154874b2eda590c8424bdda41d1`)
	r.Run(`kubectl wait -n observability --timeout=1m --for=condition=ready pod -l app=opentelemetry`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/observability/jaeger_and_prometheus/example?ref=27125b7983e1a154874b2eda590c8424bdda41d1`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-jaeger-and-prometheus`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-jaeger-and-prometheus`)
	r.Run(`kubectl exec pods/alpine -n ns-jaeger-and-prometheus -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-jaeger-and-prometheus -- ping -c 4 172.16.1.101`)
	r.Run(`NSE_NODE=$(kubectl get pods -l app=nse-kernel -n ns-jaeger-and-prometheus --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')` + "\n" + `FORWARDER=$(kubectl get pods -l app=forwarder-vpp --field-selector spec.nodeName==${NSE_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl port-forward service/jaeger -n observability 16686:16686 2>&1 > /dev/null &` + "\n" + `kubectl port-forward service/prometheus -n observability 9090:9090 2>&1 > /dev/null &`)
	r.Run(`result=$(curl -X GET localhost:16686/api/traces?service=${FORWARDER}&lookback=5m&limit=1)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -q "forwarder"`)
	r.Run(`FORWARDER=${FORWARDER//-/_}`)
	r.Run(`result=$(curl -X GET localhost:9090/api/v1/query?query="${FORWARDER}_server_tx_bytes")` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -q "forwarder"`)
}
func (s *Suite) Test() {}
