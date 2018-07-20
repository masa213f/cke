package cke

import (
	"context"
	"net"
	"sync"

	"github.com/cybozu-go/cmd"
	"github.com/pkg/errors"
)

// ClusterStatus represents the working cluster status.
// The structure reflects Cluster, of course.
type ClusterStatus struct {
	Name          string
	NodeStatuses  map[string]*NodeStatus // keys are IP address strings.
	Agents        map[string]Agent       // ditto.
	ServiceSubnet *net.IPNet
	RBAC          bool // true if RBAC is enabled
	// TODO:
	// CoreDNS will be deployed as k8s Pods.
	// We probably need to use k8s API to query CoreDNS service status.
}

// Destroy calls Close for all agents.
func (cs *ClusterStatus) Destroy() {
	for _, a := range cs.Agents {
		a.Close()
	}
	cs.Agents = nil
}

// NodeStatus status of a node.
type NodeStatus struct {
	Etcd       EtcdStatus
	APIServer  ServiceStatus
	Controller ServiceStatus
	Scheduler  ServiceStatus
	Proxy      ServiceStatus
	Kubelet    KubeletStatus
	Labels     map[string]string // are labels for k8s Node resource.
}

// IsControlPlane returns true if the node has been configured as a control plane.
func (ns *NodeStatus) IsControlPlane() bool {
	return ns.Etcd.HasData
}

// ServiceStatus represents statuses of a service.
//
// If Running is false, the service is not running on the node.
// ExtraXX are extra parameters of the running service, if any.
type ServiceStatus struct {
	Running        bool
	Image          string
	ExtraArguments []string
	ExtraBinds     []Mount
	ExtraEnvvar    map[string]string
}


// KubeletStatus is the status of kubelet.
type KubeletStatus struct {
	ServiceStatus
	Domain    string
	AllowSwap bool
}

// GetClusterStatus consults the whole cluster and constructs *ClusterStatus.
func GetAgents(ctx context.Context, cluster *Cluster) (map[string]Agent, error) {
	var mu sync.Mutex
	needClose := true
	agents := make(map[string]Agent)
	defer func() {
		if !needClose {
			return
		}
			for _, a := range agents {
				a.Close()
			}
	}()

	env := cmd.NewEnvironment(ctx)
	for _, n := range cluster.Nodes {
		n := n
		env.Go(func(ctx context.Context) error {
			a, err := SSHAgent(n)
			if err != nil {
				return errors.Wrap(err, n.Address)
			}
			mu.Lock()
			agents[n.Address] = a
			mu.Unlock()
			return nil
		})
	}
	env.Stop()
	err := env.Wait()
	if err != nil {
		return nil, err
	}

	needClose = false
	return agents, nil
}

