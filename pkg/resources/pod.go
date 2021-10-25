package resources

import (
	"context"
	"fmt"
	"os"
	"sort"

	"text/tabwriter"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Resp struct {
	pod  string
	creq string
	clim string
	mreq string
	mlim string
}

type Container struct {
	name string
	creq int64
	clim int64
	mreq int64
	mlim int64
}

type P struct {
	name       string
	containers []Container
	creqSum    int64
	climSum    int64
	mreqSum    int64
	mlimSum    int64
}

func getWriter() *tabwriter.Writer {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 5, '\t', 0)
	return w
}

func (pod P) sumResource() P {
	for _, cntr := range pod.containers {
		pod.creqSum += cntr.creq
		pod.mreqSum += cntr.mreq
		pod.climSum += cntr.clim
		pod.mlimSum += cntr.mlim
	}
	pod.mlimSum = pod.mlimSum / 1024 / 1024
	pod.mreqSum = pod.mreqSum / 1024 / 1024
	return pod
}

func GetPodInfo(ns, node string, client *kubernetes.Clientset) {
	var listOfPods []P
	var p P
	w := getWriter()
	defer w.Flush()
	pods, _ := client.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	fmt.Printf("Node : %s\n", node)
	for _, pod := range pods.Items {
		if pod.Spec.NodeName == node {
			//fmt.Printf("POD: %s\n", pod.ObjectMeta.Name)
			p.name = pod.ObjectMeta.Name
			p.containers = make([]Container, len(pod.Spec.Containers))
			for idx, container := range pod.Spec.Containers {
				//fmt.Printf("Container: %s\n", container.Name)
				p.containers[idx].name = container.Name
				p.containers[idx].creq = container.Resources.Requests.Cpu().MilliValue()
				p.containers[idx].mreq = container.Resources.Requests.Memory().Value()
				p.containers[idx].clim = container.Resources.Limits.Cpu().MilliValue()
				p.containers[idx].mlim = container.Resources.Limits.Memory().Value()
			}
			listOfPods = append(listOfPods, p)
		}
	}
	fmt.Fprintf(w, "\nPOD\tCPU REQUEST\tCPU LIMIT\tMEM REQUEST\tMEM LIMIT")
	for _, pod := range listOfPods {
		pod = pod.sumResource()
		fmt.Fprintf(w, "\n%s\t%dm\t%dm\t%dMi\t%dMi", pod.name, pod.creqSum, pod.climSum, pod.mreqSum, pod.mlimSum)
	}
	fmt.Fprintf(w, "\n")
}
