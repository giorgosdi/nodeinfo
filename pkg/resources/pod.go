package resources

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/giorgosdi/nodeinfo/pkg/logger"
	"github.com/giorgosdi/nodeinfo/pkg/options"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

type podLogger struct {
	pods   []P
	header int
	body   int
}

type Container struct {
	name  string
	creq  int64
	clim  int64
	mreq  int64
	mlim  int64
	cutil int64
	mutil int64
}

type P struct {
	kind       string
	name       string
	namespace  string
	containers []Container
	creqSum    int64
	climSum    int64
	mreqSum    int64
	mlimSum    int64
	cutilSum   int64
	mutilSum   int64
}

func loop(pods *v1.PodList, node string, metricsClient metricsv.Interface) []P {
	var listOfPods []P
	var p P
	for _, pod := range pods.Items {
		if pod.Spec.NodeName == node {
			p.name = pod.GetName()
			p.namespace = pod.GetNamespace()
			p.kind = pod.TypeMeta.Kind
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
	listWithFullMetrics := utilMetrics(listOfPods, metricsClient)

	aggregatedPodMetrics := aggregateMetrics(listWithFullMetrics)
	return aggregatedPodMetrics
}

func utilMetrics(pods []P, metricsClient metricsv.Interface) []P {
	var p []P
	for _, pod := range pods {
		pod = pod.podMetrics(metricsClient)
		p = append(p, pod)
	}
	return p
}

// TODO: use pass by reference rather than pass by value
func (pod P) podMetrics(metricsClient metricsv.Interface) P {
	podMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses(pod.namespace).Get(context.TODO(), pod.name, metav1.GetOptions{})
	var pd P
	var cn []Container
	if err != nil {
		fmt.Printf("Could not get utilization metrics for : %s", pod.name)
	}
	pd.name = pod.name
	pd.namespace = pod.namespace
	pd.kind = pod.kind
	// FIX: optmize the loop
	for _, liveMetrics := range podMetricsList.Containers {
		for _, cntr := range pod.containers {
			if cntr.name == liveMetrics.Name {
				cntr.cutil = liveMetrics.Usage.Cpu().MilliValue()
				cntr.mutil = liveMetrics.Usage.Memory().Value()
				cn = append(cn, cntr)
			}
		}
	}
	pd.containers = cn
	return pd
}

func (pod P) sumResource() P {
	for _, cntr := range pod.containers {
		pod.creqSum += cntr.creq
		pod.mreqSum += cntr.mreq
		pod.climSum += cntr.clim
		pod.mlimSum += cntr.mlim
		pod.cutilSum += cntr.cutil
		pod.mutilSum += cntr.mutil
	}
	pod.mlimSum = pod.mlimSum / 1024 / 1024
	pod.mreqSum = pod.mreqSum / 1024 / 1024
	pod.mutilSum = pod.mutilSum / 1024 / 1024
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

func (pLogger podLogger) GetHeader(w *tabwriter.Writer) {
	_, err := fmt.Fprintf(w, "POD\tNAMESPACE\tCPU REQUEST\tCPU LIMIT\tCPU UTIL\tMEM REQUEST\tMEM LIMIT\tMEM UTIL\n")
	if err != nil {
		fmt.Println("Could not fetch header. Hint: maybe the tabwriter is facing issues")
	}
}

func (pLogger podLogger) GetBody(w *tabwriter.Writer) {
	for _, pod := range pLogger.pods {
		_, err := fmt.Fprintf(w, "%s\t%s\t%dm\t%dm\t%dm\t%dMi\t%dMi\t%dMi\n", pod.name, pod.namespace, pod.creqSum, pod.climSum, pod.cutilSum, pod.mreqSum, pod.mlimSum, pod.mutilSum)
		if err != nil {
			fmt.Println("Could not fetch body. Hint: maybe the tabwriter is facing issues or the pod struct is missing a field")
		}
	}
}

func (pLogger podLogger) Log(w *tabwriter.Writer) {
	w.Flush()
}

func GetPodInfo(o *options.NodeInfoOptions, client *kubernetes.Clientset, metricsClient metricsv.Interface) {
	var pLogger podLogger
	pods, _ := client.CoreV1().Pods(o.Namespace).List(context.TODO(), metav1.ListOptions{})
	listOfPods := loop(pods, o.Args[0], metricsClient)
	if len(listOfPods) == 0 {
		fmt.Println("Node does not exist or has no pods in the current namespace(or at all)")
		os.Exit(0)
	}
	pLogger.pods = listOfPods
	logger.TableLogger(pLogger)

}
