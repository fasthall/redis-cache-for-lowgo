import yaml
import requests
import socket
import time

with open('/config/config.yaml', 'r') as ifile:
    obj = yaml.load(ifile)
    print('Reporting to ' + obj['controller'])
    ip = socket.gethostbyname(socket.gethostname())
    done = False
    while not done:
        try:
            r = requests.post('http://' + obj['controller'] + '/rediscache?host=' + ip + ':6379')
            print(r.text)
            done = True
        except Exception as e:
            print(e)
            print('Retry in 2 seconds')
            time.sleep(2)
    print('Address reported')
