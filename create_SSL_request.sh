#!/bin/bash

usage()
{
	echo "Usage $0 -e 'user@host.com' -o 'organization name' -c 'Country code' -s 'State code' -l 'Locality' [-p 'key pass' -d '/path/to/desired/folder' -u 'organizational unit']"
	exit
}

while getopts "u:e:o:c:s:l:p:d:h?" opt; do
	case $opt in
		e)
			EMAIL="$OPTARG"
			;;
		o)
			ORG="$OPTARG"
			;;
		c)
			COUNTRY="$OPTARG"
			;;
		s)
			STATE="$OPTARG"
			;;
		l)
			LOC="$OPTARG"
			;;
		p)
			PASS="$OPTARG"
			;;
		d)
			DESTDIR="$OPTARG"
			;;
		u)
			OU="/OU=$OPTARG"
			;;
		h)
			usage
			;;
		?)
			usage
			;;
	esac
done
shift $(($OPTIND -1))
DN=$1

echo "loc=$LOC 
email=$EMAIL
org=$ORG
country=$COUNTRY 
state=$STATE
destdir=$DESTDIR
ou=$OU
dn=$DN"
if [[ ! $LOC  || ! $EMAIL || ! $ORG || ! $COUNTRY || ! $STATE || ! $DN ]]
	then
		usage
fi
if [[ ! $PASS ]]
	then
		PASS=`pwgen -n 24`
		echo "passphrase is $PASS"
fi
if [[ ! $DESTDIR ]]
	then
		DESTDIR=$DN
fi

if [[ ! -e $DESTDIR ]]
	then
		echo "Creating dir $DESTDIR"
		mkdir -p $DESTDIR
fi
pushd $DESTDIR
echo $PASS >> passphrase
openssl genrsa -aes256 -passout pass:$PASS -out ${DN}.encrypt.key 2048
openssl rsa -in ${DN}.encrypt.key -passin pass:$PASS -out ${DN}.key
openssl req -new -sha256 -key ${DN}.key -out ${DN}.csr -subj "/CN=${DN}/emailAddress=$EMAIL/O=$ORG/C=$COUNTRY/ST=$STATE/L=$LOC"
popd
cat ${DESTDIR}/${DN}.csr
