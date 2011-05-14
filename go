#!/usr/bin/env bash
#this will copy our public key to a remote host and ssh to it.

userhost=$1
keyfile=~/.ssh/id_rsa.pub
authkeyfile='~/.ssh/authorized_keys'


#if no ssh public key exists, create one in the default spot
if [ ! -e $keyfile ]
  then
    echo "Creating SSH key in $keyfile"
    ssh-keygen -t rsa  -f $keyfile -q -N ''
fi
#now get the key itself into a variable
mypubkey=`cat $keyfile`


#if no username is passed (like someuser@somehost), use root by default
if [[ ! "$userhost" =~ '@' ]]
  then
    userhost=root@$1
fi

#this keeps it to one time needed to enter the password,
#it'll create the .ssh directory with right perms, touch the key file,
#create a backup without our key (no dupes),
#and copy it back
ssh $userhost "mkdir -p .ssh; 
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
      scp -q $f $userhost:$f
    done
fi


#and finally, ssh to the host.
ssh $userhost $2
