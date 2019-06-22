# TokenRequest

Minimal code to use the Kubernetes TokenRequest API to create an audience scoped token for a ServiceAccount.

```bash
$ kubectl create -f rbac.yaml
$ kubectl run -i -t tokenrequest --image=quay.io/brancz/tokenrequest:v0.0.1 --restart=Never
```
