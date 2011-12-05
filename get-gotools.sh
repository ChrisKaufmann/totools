#!/usr/bin/env bash
location='https://raw.github.com/ChrisKaufmann/gotools/master/'
for thing in ami background didi ducks go gox killthe
do
	curl $location/$thing -o /usr/local/bin/$thing
	chmod +x /usr/local/bin/$thing
done
