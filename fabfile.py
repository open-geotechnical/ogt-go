

import os
import sys

from fabric.api import env, local, run, lcd, cd, sudo, warn_only, prompt
from fabric.context_managers import shell_env



env.hosts = [ 'ags' ]
env.user = "ags"
env.password = "using keys and ./.config"
env.shell = "/bin/sh -c"
env.use_ssh_config = True

HERE_PATH =  os.path.abspath( os.path.dirname( __file__ )	 )


AGS_DEF_GIT = "git@bitbucket.org:daf0dil/ags-def-json.git"


LIVE_DIR = "/home/ags/ags2go"

def ws_update():
	"""Update developer `machine` workspace back"""
	ws_dir =  HERE_PATH + "/workspace"
	if not os.path.exists(ws_dir):
		local("mkdir %s" % ws_dir )


	ags_def = ws_dir + "/ags-data-json"
	if not os.path.exists(ags_def):
		with lcd(ws_dir):
			local("git clone %s" % AGS_DEF_GIT)

def test_ssh():
	"""Logins into remote servers and print basics"""
	run("whoami")
	run("pwd")


def up_server():
	"""Update server at ags.daffodil.uk.com"""
	local("git push origin master")
	with cd(LIVE_DIR):
		run("git pull origin master")
	r_build()
	r_run()

def r_build():
	"""Build app on remote server"""
	with shell_env(GOPATH="/home/ags"):
		with cd(LIVE_DIR):
			run("go build -v")


def r_run():
	"""Start/Reboot remote server"""
	sudo("/usr/local/bin/supervisorctl restart ags", pty=True)