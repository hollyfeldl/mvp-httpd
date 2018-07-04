#!/bin/bash

# script to perform a rolling upgrade to markrank.net website

echo "**********************************"
echo "* Upgrade markrank.net           *"
echo "*                                *"
echo "**********************************"

# look for container name as the second parm and if not default it

if [ -z $3 ] 
then
	curRepController="markrank-net-httpd-rc"
	echo "WARN - No replication controller provided, default to" $curRepController
else
	$curRepController=$3
	echo "INFO - replication controller set to" $curRepController
fi

if [ -z $2 ] 
then
	curContainer="mvp-httpd"
	echo "WARN - No container provided, default to" $curContainer
else
	curContainer=$2
	echo "INFO - Container set to" $curContainer
fi

# look for the project id
if [ -z $1 ] 
then
	curProject="my-container-httpd"
	echo "WARN - No project provided, default to" $curProject
else
	curProject=$1
	echo "INFO - Project set to" $curProject
fi

#look what is in the SSL staging area
ls -als ./ssl

# clear it out
echo -n "Lets Encrypt Files There? (y/n)?"
read curClearSSL
case $curClearSSL in
	y) echo "Let's Go!";;
	Y) echo "Sure, same thing. Let's Go!";;
    *) echo "Put them there and try again."; exit 0;;
esac

# ask if we should continue
echo -n "Should we proceed (y/n)?"
read curPrompt
case $curPrompt in
	y) echo "Let's Go!";;
	Y) echo "Sure, same thing. Let's Go!";;
    *) echo "Well, better luck next time."; exit 0;;
esac

# look for the Dockerfile
if test -e "./Dockerfile"
then 
	echo "INFO - Using local Dockerfile"
	cat ./Dockerfile
else
	echo "ERROR - Cannot find Dockerfile to build container"
	exit 1
fi

# look for a private key
if test -e "./ssl/privkey.pem"
then 
	echo "INFO - Found the private key file"
else
	echo "ERROR - Cannot find private key file file"
	exit 1
fi

# look for cert full chain 
if test -e "./ssl/fullchain.pem"
then 
	echo "INFO - Found the cert fullchain file"
else
	echo "ERROR - Cannot find cert fullchain file"
	exit 1
fi

echo "INFO - Come up with unique container label"
curLabel=$(date --rfc-3339=date)
echo "INFO - Container will be "$curContainer":"$curLabel

echo "INFO - rebuild the container using the latest httpd" 
docker build -t gcr.io/$curProject/$curContainer:$curLabel --pull=true .

# some feedback
docker images | grep $curContainer

echo "INFO - push the container to gcloud repository"
gcloud docker -- push gcr.io/$curProject/$curContainer:$curLabel

echo "INFO - perform a rolling upgrade of " $curRepController " to the new container image"
kubectl rolling-update $curRepController --image=gcr.io/$curProject/$curContainer:$curLabel

echo "INFO - check the replication controller version (kubectl describe rc " $curRepController ")" 
kubectl describe rc $curRepController

# we are done
exit 0

