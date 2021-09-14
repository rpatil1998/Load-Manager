import paramiko
from sshtunnel import SSHTunnelForwarder


# kuch nahi 

with SSHTunnelForwarder(
    ('<jump server public ip>', 22),
    ssh_username="ec2-user",
    ssh_pkey=r"C:\Users\CHINMOY\Desktop\aws_rounik.pem",
    #ssh_private_key_password="secret",
    remote_bind_address=('<private server private ip>',22), #inbound route = security grp of jump server 
    local_bind_address=('0.0.0.0', 10022)
) as tunnel:
    client = paramiko.SSHClient()
    client.load_system_host_keys()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect('127.0.0.1',10022 , username='ec2-user', key_filename=r'C:\Users\CHINMOY\Desktop\aws_rounik.pem')
    # do some operations with client session
    #use sleep to make tunnel alive for some time
    client.close() #closes the ssh tunnel

#print('FINISH!')