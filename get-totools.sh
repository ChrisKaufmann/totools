#!/usr/bin/env bash
location='https://raw.githubusercontent.com/ChrisKaufmann/totools/master'
for thing in ami amp background didi ducks to tox killthe killport pwgen loop radduser spork
do
	curl $location/$thing -o /usr/local/bin/$thing
	chmod +x /usr/local/bin/$thing
done
curl https://raw.githubusercontent.com/ChrisKaufmann/Threadeach/master/threadeach.pm -o /usr/local/lib64/perl5/
