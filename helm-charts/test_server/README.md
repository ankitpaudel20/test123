# Some improvements and other usecase based optimizations

* this uses basic routing and deployment.
* more advanced routing can be achived using CRDs or scripting depending upon which ingress controller we choose to use.
* moreover if our current app is expected to grow into very large custer cluster, capable of handling a lot of traffic, then using istio or other service meshes, along with flagger for better releases.
    * if that is not the case, the usual ingress-nginx or something like contour can and should be used, decreasing the complexity of cluster along with not using much of cluster resource.
* If we are using self hosted database, depending upon cases, I would try to separate that from the original application cluster or get a managed one if possible.