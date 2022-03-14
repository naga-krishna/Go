package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func connectToK8s() *kubernetes.Clientset {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}

	configPath := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatalln("failed to create K8s config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create K8s clientset")
	}

	return clientset
}

func CreatePod(clientset *kubernetes.Clientset, jobName *string, image *string) {

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *jobName,
			Namespace: "go-lang-ns",
			Labels: map[string]string{
				"k8s-app": "kube-dns",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  *jobName,
					Image: *image,
					Ports: []v1.ContainerPort{
						{
							Name:          "http",
							Protocol:      v1.ProtocolTCP,
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}

	_, err := clientset.CoreV1().Pods("go-lang-ns").Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Failed to create K8s pod.")
	}

	//print job details
	log.Println("Created K8s pod successfully")
}

func main() {
	jobName := flag.String("jobname", "test-job", "The name of the job")
	containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")

	flag.Parse()

	clientset := connectToK8s()
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get namespaces:", err)
	}

	fmt.Println("Kube Namespaces: ")
	for i, ns := range namespaces.Items {
		fmt.Printf("[%d] %s\n", i, ns.GetName())
	}

	nsName := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "go-lang-ns",
		},
	}

	_, err = clientset.CoreV1().Namespaces().Create(context.Background(), nsName, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
	}
	CreatePod(clientset, jobName, containerImage)

	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{LabelSelector: "k8s-app"})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Println("Namespace: ", pod.Namespace, " Pod name: ", pod.Name)
	}

	err = clientset.CoreV1().Pods("go-lang-ns").Delete(context.Background(), "test", metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleting the test pod....")

	time.Sleep(5 * time.Second)
	newpods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{LabelSelector: "k8s-app"})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range newpods.Items {
		fmt.Println("Namespace: ", pod.Namespace, " Pod name: ", pod.Name)
	}
}
