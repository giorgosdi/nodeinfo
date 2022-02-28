# kubectl nodeinfo

A `kubectl` plugin that provides you information for a given node. It will tell you what pods live in the a node and their utilization metrics and limits/requests. You can also find the pods of a node for a particular namespace rather than querying in all namespaces.

## Installation

You should use the [krew](https://krew.sigs.k8s.io/) plugin manager to install this plugin.

_[PR](https://github.com/kubernetes-sigs/krew-index/pull/2029) for the plugin to be added in the krew-index is still open so the below command wont work for now._
```
kubectl krew install nodeinfo
```
If you want to install it locally, you can pull the repo and run `go build .`. Alternatively you can utilize the make rules.

_dont forget to move the executable in your `/usr/local/bin`_


## Usage

`kubectl nodeinfo <node>`

example:

![nodeinfo example](assets/nodeinfo_example.png)

## Upcoming features

* Show the capacity and labels of the node itself
* Make metrics show as an option instead of having them by default
* Get information about the containers inside a pod
* Get the Status of a pod (Running, CrashLoopBackOff etc)
* (maybe) output to different sources like json and csv. Not really sure if that is useful since this is focusing on the pods of a node.


## Similar plugins

There are a couple of similar plugins in the krew-index that you can use.

* [viewnode](https://github.com/NTTDATA-DACH/viewnode). The difference with `nodeinfo` is that `viewnode` is looking at all node running in your cluster rather than passing a specific node in `nodeinfo`. There are also some other information that you are getting like the state of the pod (running etc). There are no metrics included. (at least to my knowledge)
* [kubectl-view-allocations](https://github.com/davidB/kubectl-view-allocations). The difference with `nodeinfo` is that `view-allocations` focuses heavily on the CPU/Mem allocations in nodes. It also provides information about pods grouped in resources like `CPU/Memory`. `nodeinfo`'s output is less populated with a main focus on pods running in a given node. `view-allocations` offers a variety of flags for different things like grouping by namespace and outputting to csv which looks really handy. I think that the purpose of `view-allocations` is different than `nodeinfo` but I included it to the list because they share functionality.


## LICENSE

Apache 2.0. See [LICENSE](./LICENSE).
