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
