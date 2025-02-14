// Code generated by gotestmd DO NOT EDIT.
package ipsec

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/three_cluster_configuration/ipsec"
)

type Suite struct {
	base.Suite
	ipsecSuite ipsec.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.ipsecSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
}
func (s *Suite) TestFloating_Kernel2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-floating-kernel2ip2kernel`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-floating-kernel2ip2kernel`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns ns-floating-kernel2ip2kernel`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Kernel/cluster3?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Kernel/cluster2?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-kernel2ip2kernel`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Kernel/cluster1?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-floating-kernel2ip2kernel`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-kernel2ip2kernel -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-kernel2ip2kernel -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestFloating_Kernel2IP2Memif() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-floating-kernel2ip2memif`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-floating-kernel2ip2memif`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns ns-floating-kernel2ip2memif`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Memif/cluster3?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Memif/cluster2?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=2m pod -l app=nse-memif -n ns-floating-kernel2ip2memif`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Kernel2IP2Memif/cluster1?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-floating-kernel2ip2memif`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-kernel2ip2memif -- ping -c 4 172.16.1.2`)
	r.Run(`result=$(kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-memif -n "ns-floating-kernel2ip2memif" -- vppctl ping 172.16.1.3 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
func (s *Suite) TestFloating_Memif2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-floating-memif2ip2kernel`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-floating-memif2ip2kernel`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns ns-floating-memif2ip2kernel`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Kernel/cluster3?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Kernel/cluster2?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-memif2ip2kernel`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Kernel/cluster1?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=2m pod -l app=nsc-memif -n ns-floating-memif2ip2kernel`)
	r.Run(`result=$(kubectl --kubeconfig=$KUBECONFIG1 exec deployments/nsc-memif -n "ns-floating-memif2ip2kernel" -- vppctl ping 172.16.1.2 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-memif2ip2kernel -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestFloating_Memif2IP2Memif() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-floating-memif2ip2memif`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-floating-memif2ip2memif`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns ns-floating-memif2ip2memif`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Memif/cluster3?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Memif/cluster2?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=2m pod -l app=nse-memif -n ns-floating-memif2ip2memif`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/floating_Memif2IP2Memif/cluster1?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=2m pod -l app=nsc-memif -n ns-floating-memif2ip2memif`)
	r.Run(`result=$(kubectl --kubeconfig=$KUBECONFIG1 exec deployments/nsc-memif -n "ns-floating-memif2ip2memif" -- vppctl ping 172.16.1.2 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`result=$(kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-memif -n "ns-floating-memif2ip2memif" -- vppctl ping 172.16.1.3 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
func (s *Suite) TestInterdomain_Kernel2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/interdomain_Kernel2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-interdomain-kernel2ip2kernel`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-interdomain-kernel2ip2kernel`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/interdomain_Kernel2IP2Kernel/cluster2?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-interdomain-kernel2ip2kernel`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/interdomain_Kernel2IP2Kernel/cluster1?ref=v0.1.8`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-interdomain-kernel2ip2kernel`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-interdomain-kernel2ip2kernel -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-interdomain-kernel2ip2kernel -- ping -c 4 172.16.1.3`)
}
