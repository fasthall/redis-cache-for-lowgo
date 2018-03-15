import yaml
import requests
import socket

with open('/config/config.yaml', 'r') as ifile:
    obj = yaml.load(ifile)
    print('Reporting to ' + obj['controller'])
    ip = socket.gethostbyname(socket.gethostname())
    r = requests.post('http://' + obj['controller'] + '/addRedisCache?host=' + ip + ':6379')
    print(r.text)
