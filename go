#!/usr/bin/env bash
#this will copy our public key to a remote host and ssh to it.

userhost=$1
keyfile=~/.ssh/id_rsa.pub
authkeyfile='~/.ssh/authorized_keys'

#if no userhost is specified, print usage and exit
if [ ! $userhost ]
  then
    echo "usage: $0 [user@]hostname[:port] <defaults to root@>"
	exit
fi

#if no ssh public key exists, create one in the default spot
if [ ! -e $keyfile ]
  then
    echo "Error, no public key file,please run ssh-keygen"
	exit
fi
#now get the key itself into a variable
mypubkey=`cat $keyfile`

#if no username is passed (like someuser@somehost), use root by default
if [[ ! "$userhost" =~ '@' ]]
  then
    userhost=root@$1
fi

#see if there's a port on the end, if so create $p with the port
if [[ $userhost =~ (.*@.*):(.*) ]]
  then
    userhost="${BASH_REMATCH[1]}"
  	port=" -p ${BASH_REMATCH[2]} "
  else
    port=" "
fi

#this keeps it to one time needed to enter the password,
#it'll create the .ssh directory with right perms, touch the key file,
#create a backup without our key (no dupes),
#and copy it back
ssh $port $userhost "mkdir -p .ssh; 
  chmod 700 .ssh;
  touch $authkeyfile; 
  cp $authkeyfile ${authkeyfile}.bak;
  grep -v '$mypubkey' ${authkeyfile}.bak > $authkeyfile; 
  echo '$mypubkey' >> $authkeyfile"

#if no username is passed (like someuser@somehost), use root by default
if [[  "$userhost" == 'root@'$1 ]]
  then
    userhost=root@$1
    for f in /usr/bin/ami /usr/bin/didi /usr/bin/killthe /usr/bin/redo /usr/bin/ducks
    do
	  if [ -e $f ]
	    then
          scp $port -q $f $userhost:$f
	  fi
    done
fi

#and finally, ssh to the host.
ssh $port $userhost $2
