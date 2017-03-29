#!/usr/bin/env bash

wget -c -O - https://bintray.com/user/downloadSubjectPublicKey?username=bintray | apt-key add -
apt-get -qq update
apt-get -y install apt-transport-https
echo "deb https://dl.bintray.com/like-all/deb jessie-testing main" > /etc/apt/sources.list.d/bintray.list
apt-get -qq update
apt-get -y install noop
rm -f /sbin/init
ln -s /sbin/noop /sbin/init
