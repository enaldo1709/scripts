#!/bin/bash

user=$(whoami)
if [[ $user != "root" ]]; then
	echo "please run as root"
	exit 1
fi


if [[ "$1" == "enable" ]]; then
	echo "enabling docker service..."
	systemctl enable docker.service
	systemctl enable docker.socket

	echo "starting docker service..."
	systemctl start docker.service
	systemctl start docker.socket

	echo "check docker service status"
	systemctl status docker.service | cat

	echo "run container for test"
	docker run hello-world

elif [[ "$1" == "disable" ]]; then
	echo "stopping docker service"
	systemctl stop docker.service
	systemctl stop docker.socket

	echo "disabling docket service"
	systemctl disable docker.service
	systemctl disable docker.socket

	echo "check docker service status"
	systemctl status docker.service | cat

else
	echo "nothing to do"
fi

