#!/usr/bin/env sh

echo $GPG_SECRET | base64 -d | gpg --armor --import
which gpg

# echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
# systemctl restart sshd

# Authorize SSH Host
# mkdir -p /home/root/.ssh
# chmod 0700 /home/root/.ssh
# ssh-keyscan git2.borje.zone > /home/root/.ssh/known_hosts

# Use auth key instted
#
#
#
# Add the keys and set permissions
# echo $SSH_PRIVATE_KEY | base64 -d > /home/root/.ssh/id_rsa
# echo $SSH_PUBLIC_KEY | base64 -d > /home/root/.ssh/id_rsa.pub
# chmod 600 /home/root/.ssh/id_rsa
# chmod 600 /home/root/.ssh/id_rsa.pub
# ssh-add -k /root/.ssh/id_rsa

# eval `ssh-agent -s`

echo $SSH_PRIVATE_KEY | base64 -d | ssh-add -

# ls /home/root/.ssh

# alias gopass=./gopass
# alias "gopass-jsonapi"=./gopass-jsonapi

mkdir -p /home/root/.config/gopass

echo "KLON"
git clone $SECRETS_REPO_URL /rundir/secrets
echo "SLUT KLON"

cat << EOF > /home/root/.config/gopass/config
[mounts]
	path = /rundir/secrets
EOF

PASSWORD_STORE_DIR=/rundir/secrets gopass init $GPG_KEY_ID
# ./gopass ls
echo "hejsan"
gopass-jsonapi configure --help
gopass-jsonapi configure --browser chrome --path . --global false << 'EOF'
y
EOF
# ./gopass-jsonapi listen
ls -la .
./gopass_wrapper.sh
#
# cat ./gopass_wrapper.sh
# gopass-jsonapi << 'EOF'
# y
# EOF
top
