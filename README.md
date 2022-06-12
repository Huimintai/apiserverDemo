# apiserverDemo
This demo for writing a kubernetes apiserver with apiserver-builder
# how to build a apiserver demo with apiserver-builder
```
apiserver-boot init repo --domain hm.com --module-name hm
apiserver-boot create group version resource --group animal --version v1alpha1 --kind Cat --non-namespaced=false
make all
apiserver-boot run in-cluster --image=registry.tke.com/library/hm-demo:0.0.1 --name=hm-demo --namespace=default
```
