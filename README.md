## Aruchitecture diagram

![image](https://github.com/user-attachments/assets/ca17c4d2-dc75-4cb4-8fa0-3870fba4a055)

``` eval `ssh-agent` ```
Public Key
: The public key is the one you provide to the virtual machine. It is installed on the VM when you create or provision it. This allows your local machine to authenticate with the VM using the corresponding private key.
 
Private Key: The private key remains on your local machine. When you attempt to SSH into the virtual machine, your local machine uses the private key to authenticate with the VM using the public key installed on the VM.
