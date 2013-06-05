#!/usr/bin/env bash
location='https://raw.github.com/ChrisKaufmann/totools/master/'
for thing in ami amp background didi ducks to tox killthe killport pwgen loop
do
	curl $location/$thing -o /usr/local/bin/$thing
	chmod +x /usr/local/bin/$thing
done
