# kubectl nodeinfo

A `kubectl` plugin that provides you information for a given node. It will tell you what pods live in the a node and their utilization metrics and limits/requests

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


## LICENSE

Apache 2.0. See [LICENSE](./LICENSE).
