#coding: utf-8
from fabric.api import *

env.use_ssh_config = True
env.hosts = ['beatstorm']

def deploy():
  user = prompt("User for the machine?:", default="web")
  put('deploy/deploy.sh', '/opt/beatstorm/')
  put('deploy/start.sh', '/opt/beatstorm/')
  put('redis/*', '/opt/beatstorm/redis')

  run("sed -i 's/USERNAME/%s/' /opt/beatstorm/deploy.sh" % user)
  sudo('chmod +x /opt/beatstorm/deploy.sh')
  sudo('chmod +x /opt/beatstorm/start.sh')

  with cd('~/go/src/beatstorm'):
    result = -1
    while result != 0:
      result = run('git remote update && git reset --hard origin/master').return_code
    run("GOROOT=/opt/go GOPATH=/home/%s/go/ /opt/go/bin/go get -u ./app" % user)
  with cd('/opt/beatstorm/'):
    sudo('./deploy.sh')

def setup():
  with settings(warn_only = True):
    run("mkdir -p /opt/beatstorm/")
    run("mkdir /opt/beatstorm/redis")
    run("mkdir -p ~/go/src")
  with cd('~/go/src'):
    result = -1
    while result != 0:
      result = run('git clone git@github.com:beatstorm/beatstorm.git').return_code
  put('deploy/beatstorm.nginx', '/opt/nginx/sites-available/beatstorm', True)
  sudo('ln -s /opt/nginx/sites-available/beatstorm /opt/nginx/sites-enabled/beatstorm')
