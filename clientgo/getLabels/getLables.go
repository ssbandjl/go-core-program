package operator

import (
	"flag"
	extensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/util/logs"
	clientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"os"
	"testing"
)

type OperatorManagerServer struct {
	Master     string
	Kubeconfig string
}

func NewOMServer() *OperatorManagerServer {
	s := OperatorManagerServer{}
	return &s
}

var s *OperatorManagerServer

func init() {

	s = NewOMServer()
	flag.StringVar(&s.Master, "master", s.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	flag.StringVar(&s.Kubeconfig, "kubeconfig", s.Kubeconfig, "Path to kubeconfig file with authorization and master location information.")
	//初始化klog等flag
	logs.InitLogs()
	flag.Parse()
}

func Test_DeleteCollection(t *testing.T) {
	if err := Run(s); err != nil {
		t.Fatalf("%v\n", err)
		os.Exit(1)
	}
}

func Run(s *OperatorManagerServer) error {

	var (
		generalLabelKey       = "harmonycloud.cn/statefulset"
		redisClusterName      = "redis-ll-1010"
		redisClusterNamespace = "kube-system"
	)

	kubeClient, _, _, err := createClients(s)

	if err != nil {
		return err
	}

	//根据label批量删除pod
	labelPod := labels.SelectorFromSet(labels.Set(map[string]string{generalLabelKey: redisClusterName}))
	listPodOptions := metav1.ListOptions{
		LabelSelector: labelPod.String(),
	}
	err = kubeClient.CoreV1().Pods(redisClusterNamespace).DeleteCollection(&metav1.DeleteOptions{}, listPodOptions)
	if err != nil {
		if !errors.IsNotFound(err) {
			klog.Errorf("Drop RedisCluster: %v/%v pod error: %v", redisClusterNamespace, redisClusterName, err)
			return err
		}
	}

	//根据label批量删除pvc
	labelPvc := labels.SelectorFromSet(labels.Set(map[string]string{"app": redisClusterName}))
	listPvcOptions := metav1.ListOptions{
		LabelSelector: labelPvc.String(),
	}
	err = kubeClient.CoreV1().PersistentVolumeClaims(redisClusterNamespace).DeleteCollection(&metav1.DeleteOptions{}, listPvcOptions)
	if err != nil {
		if !errors.IsNotFound(err) {
			klog.Errorf("Drop RedisCluster: %v/%v pvc error: %v", redisClusterNamespace, redisClusterName, err)
			return err
		}
	}

	//如果pv没有删除掉,则删除
	labelPv := labels.SelectorFromSet(labels.Set(map[string]string{generalLabelKey: redisClusterName}))
	listPvOptions := metav1.ListOptions{
		LabelSelector: labelPv.String(),
	}
	err = kubeClient.CoreV1().PersistentVolumes().DeleteCollection(&metav1.DeleteOptions{}, listPvOptions)

	if err != nil {
		if !errors.IsNotFound(err) {
			klog.Errorf("Drop RedisCluster: %v/%v pv error: %v", redisClusterNamespace, redisClusterName, err)
			return err
		}
	}

	return nil
}

//根据kubeconfig文件创建客户端
func createClients(s *OperatorManagerServer) (*clientset.Clientset, *extensionsclient.Clientset, *restclient.Config, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags(s.Master, s.Kubeconfig)
	if err != nil {
		return nil, nil, nil, err
	}

	kubeconfig.QPS = 100
	kubeconfig.Burst = 100

	kubeClient, err := clientset.NewForConfig(restclient.AddUserAgent(kubeconfig, "operator-manager"))
	if err != nil {
		klog.Fatalf("Invalid API configuration: %v", err)
	}

	extensionClient, err := extensionsclient.NewForConfig(restclient.AddUserAgent(kubeconfig, "operator-manager"))
	if err != nil {
		klog.Fatalf("Invalid API configuration: %v", err)
	}

	return kubeClient, extensionClient, kubeconfig, nil
}
