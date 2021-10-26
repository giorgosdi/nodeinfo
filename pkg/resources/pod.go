package resources

import (
	"context"
	"fmt"
	logger "nodeinfo/pkg/logger"
	"text/tabwriter"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type podLogger struct {
	pods   []P
	header int
	body   int
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

func (pLogger podLogger) GetHeader(w *tabwriter.Writer) (*tabwriter.Writer, int) {
	header, err := fmt.Fprintf(w, "\nPOD\tCPU REQUEST\tCPU LIMIT\tMEM REQUEST\tMEM LIMIT")
	if err != nil {
		fmt.Println("Could not fetch header. Hint: maybe the tabwriter is facing issues")
	}
	return w, header
}

func (pLogger podLogger) GetBody(w *tabwriter.Writer, pod P) int {
	body, err := fmt.Fprintf(w, "\n%s\t%dm\t%dm\t%dMi\t%dMi", pod.name, pod.creqSum, pod.climSum, pod.mreqSum, pod.mlimSum)
	if err != nil {
		fmt.Println("Could not fetch body. Hint: maybe the tabwriter is facing issues or the pod struct is missing a field")
	}
	return body
}

func (pLogger podLogger) Log(w *tabwriter.Writer) (*tabwriter.Writer, int) {
	var val int
	for _, pod := range pLogger.pods {
		val = pLogger.GetBody(w, pod)
	}
	return w, val
}

func loop(pods *v1.PodList, node string) []P {
	var listOfPods []P
	var p P
	for _, pod := range pods.Items {
		if pod.Spec.NodeName == node {
			p.name = pod.ObjectMeta.Name
			p.containers = make([]Container, len(pod.Spec.Containers))
			for idx, container := range pod.Spec.Containers {
				p.containers[idx].name = container.Name
				p.containers[idx].creq = container.Resources.Requests.Cpu().MilliValue()
				p.containers[idx].mreq = container.Resources.Requests.Memory().Value()
				p.containers[idx].clim = container.Resources.Limits.Cpu().MilliValue()
				p.containers[idx].mlim = container.Resources.Limits.Memory().Value()
			}
			listOfPods = append(listOfPods, p)
		}
	}
	aggregatedPodMetrics := aggregateMetrics(listOfPods)
	return aggregatedPodMetrics
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

func aggregateMetrics(pods []P) []P {
	listOfPods := []P{}
	for _, pod := range pods {
		pod = pod.sumResource()
		listOfPods = append(listOfPods, pod)
	}
	return listOfPods
}

func GetPodInfo(ns, node string, client *kubernetes.Clientset) {
	var pLogger podLogger
	pods, _ := client.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	fmt.Printf("Node : %s\n", node)
	listOfPods := loop(pods, node)
	pLogger.pods = listOfPods
	logger.TableLogger(pLogger)
}
