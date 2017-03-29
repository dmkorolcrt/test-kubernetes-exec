package main

import (
"bytes"
"fmt"
"flag"
"k8s.io/kubernetes/pkg/api"

kubeletcmd "k8s.io/kubernetes/pkg/kubelet/server/remotecommand"
"k8s.io/kubernetes/pkg/client/unversioned/remotecommand"
client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"

)
var (
	kubeconfig = flag.String("kubeconfig", "/Users/dmitriy_korol/GoglandProjects/kube-exec/src/k8s.io/client-go/examples/config/config", "absolute path to the kubeconfig file")
)
func main () {

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}



	conn, err := client.New(config)
	if err != nil {
		fmt.Println(err)

	}
//	var t *fromkubernetes.r
//get list of pods
//	clientset, err := kubernetes.NewForConfig( config)
//	if err != nil {
//		panic(err.Error())
//	}
/*
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
*/
	req := conn.RESTClient.Post().
		Resource("pods").
		Name("mongo-1448603570-10d6d").
		Namespace("default").
		SubResource("exec")
	req.VersionedParams(&api.PodExecOptions{
		Container: "mongo",
		Command: []string{"/bin/bash", "-c", "ss"},
		Stdout:  true,
		Stderr:  true,
	}, api.ParameterCodec)

	// Create SPDY connection
	exec, err := remotecommand.NewExecutor(config, "POST", req.URL())
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to setup a session with mongo" )
	}
	var b bytes.Buffer
	var berr bytes.Buffer

	// Excute command
	err = exec.Stream(remotecommand.StreamOptions{
		SupportedProtocols: kubeletcmd.SupportedStreamingProtocols,
		Stdout:             &b,
		Stderr:             &berr,
	})


	s := b.String()
	fmt.Println(s)
}

