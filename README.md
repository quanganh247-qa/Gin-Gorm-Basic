## Aruchitecture diagram

![image](https://github.com/user-attachments/assets/ca17c4d2-dc75-4cb4-8fa0-3870fba4a055)

``` eval `ssh-agent` ```
Public Key
: The public key is the one you provide to the virtual machine. It is installed on the VM when you create or provision it. This allows your local machine to authenticate with the VM using the corresponding private key.
 
Private Key: The private key remains on your local machine. When you attempt to SSH into the virtual machine, your local machine uses the private key to authenticate with the VM using the public key installed on the VM.

``` sudo chmod 644 /etc/rancher/k3s/k3s.yaml ```
``` cat /etc/rancher/k3s/k3s.yaml ```
``` clusterctl init --infrastructure vsphere --bootstrap k3s --control-plane k3s ```
![image](https://github.com/user-attachments/assets/3423f308-b468-4f03-beaf-62ca6ccceac3)

![image](https://github.com/user-attachments/assets/260f03fc-5ee9-4666-a74f-ec1769aa10c0)

![image](https://github.com/user-attachments/assets/d49d22dd-d580-4d34-93ad-66034d919fb4)
=> 
`export KUBECONFIG=/etc/rancher/k3s/k3s.yaml`
`sudo chmod 644 /etc/rancher/k3s/k3s.yaml`


![image](https://github.com/user-attachments/assets/59f916db-ecdc-4cdc-a3d6-cba3e9c50d93)

check logs pods:
```sudo k3s kubectl get pods -n capv-system```
``` sudo k3s kubectl logs <pod-name> -n capv-system ```

![image](https://github.com/user-attachments/assets/e1b20064-e311-4fd8-b62c-46645e74a44c)

 ```clusterctl delete --infrastructure vsphere```

I1027 16:22:31.756225       1 session.go:148] "Found active cached vSphere client session" controller="vspherecluster" controllerGroup="infrastructure.cluster.x-k8s.io" controllerKind="VSphereCluster" VSphereCluster="default/k3s-test" namespace="default" name="k3s-test" reconcileID="96468b46-24aa-465c-963f-c586e0268e85" Cluster="default/k3s-test" server="10.1.148.31" datacenter="" username="administrator@vsphere.local"

![Uploading image.png…]()

