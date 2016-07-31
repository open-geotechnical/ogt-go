

import os
import sys

from fabric.api import env, local, run, lcd, cd, sudo, warn_only, prompt



env.hosts = [ 'ags.daffodil.uk.com' ]
env.user = "ags"
env.password = "using keys and ./.config"
env.shell = "/bin/sh -c"
env.use_ssh_config = True

HERE_PATH =  os.path.abspath( os.path.dirname( __file__ )	 )


AGS_DEF_GIT = "git@bitbucket.org:daf0dil/ags-def-json.git"


def ws_update():
	"""Update developer `machine` workspace back"""
	ws_dir =  HERE_PATH + "/workspace"
	if not os.path.exists(ws_dir):
		local("mkdir %s" % ws_dir )


	ags_def = ws_dir + "/ags-data-json"
	if not os.path.exists(ags_def):
		with lcd(ws_dir):
			local("git clone %s" % AGS_DEF_GIT)

def up_server():
	"""Update server at ags.daffodil.uk.com"""
	with cd(LIVE_DIR):
