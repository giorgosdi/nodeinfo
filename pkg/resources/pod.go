package resources

import (
	"context"
	"fmt"

	//"text/tabwritter"

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

type Aggregator struct {
	creq []int64
	clim []int64
	mreq []int64
	mlim []int64
}

type P struct {
	name string
	Aggregator
}

//func getWriter() tabwritter.Writter {
//	w := new(tabwritter.Writter)
//	w.Init(os.Stdout, 8, 8, 5, '\t', 0)
//	return w
//}
//

func sum(array []int64) int64 {
	var result int64
	for _, v := range array {
		result += v
	}
	return result
}

func GetPodInfo(ns, node string, client *kubernetes.Clientset) {
	var listOfPods []P
	var p P
	//w := getWriter()
	pods, _ := client.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	fmt.Printf("Node : %s\n", node)
	for _, pod := range pods.Items {
		p = P{}
		if pod.Spec.NodeName == node {
			fmt.Printf("POD: %s\n", pod.ObjectMeta.Name)
			p.name = pod.ObjectMeta.Name
			for _, container := range pod.Spec.Containers {
				fmt.Printf("Container: %s\n", container.Name)
				cpureq := container.Resources.Requests.Cpu().MilliValue()
				//memreq := container.Resources.Requests.Memory().ScaledValue(6)
				////cpulim := container.Resources.Limits.Cpu().MilliValue()
				//////memlim := container.Resources.Limits.Memory().ScaledValue(6)
				p.creq = append(p.creq, cpureq)
				////p.clim = append(p.creq, cpulim)
				//p.mreq = append(p.creq, memreq)
				//p.mlim = append(p.creq, memlim)
			}
			fmt.Println(p)
			listOfPods = append(listOfPods, p)
		}
	}
	fmt.Println(listOfPods)
	//fmt.Fprintf(w, "\nPod\tCPU req/limit\tMem req/limit")

}
