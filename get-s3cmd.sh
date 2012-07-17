#!/usr/bin/env bash
pushd /etc/yum.repos.d
	wget http://s3tools.org/repo/RHEL_6/s3tools.repo
	yum -y install s3cmd
popd
