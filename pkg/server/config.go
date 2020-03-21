package server

import (
	"flag"

	"github.com/spf13/pflag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

type Configuration struct {
	BindSocket     string
	KubeConfigFile string
	KubeClient     kubernetes.Interface
}

// ParseFlags ...
// TODO: validate configuration
func ParseFlags() (*Configuration, error) {
	var (
		argBindSocket     = pflag.String("bind-socket", "/var/run/cniserver.sock", "The socket daemon bind to.")
		argKubeConfigFile = pflag.String("kubeconfig", "", "Path to kubeconfig file with authorization and master location information. If not set use the inCluster token.")
	)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	// Init for glog calls in kubernetes packages
	flag.CommandLine.Parse(make([]string, 0))

	config := &Configuration{
		BindSocket:     *argBindSocket,
		KubeConfigFile: *argKubeConfigFile,
	}
	err := config.initKubeClient()
	if err != nil {
		return nil, err
	}
	klog.Infof("bind socket: %s", config.BindSocket)
	return config, nil
}

func (config *Configuration) initKubeClient() error {
	var cfg *rest.Config
	var err error
	if config.KubeConfigFile == "" {
		klog.Infof("no --kubeconfig, use in-cluster kubernetes config")
		cfg, err = rest.InClusterConfig()
		if err != nil {
			klog.Errorf("use in cluster config failed %v", err)
			return err
		}
	} else {
		cfg, err = clientcmd.BuildConfigFromFlags("", config.KubeConfigFile)
		if err != nil {
			klog.Errorf("use --kubeconfig %s failed %v", config.KubeConfigFile, err)
			return err
		}
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Errorf("init kubernetes client failed %v", err)
		return err
	}

	config.KubeClient = kubeClient
	return nil
}
