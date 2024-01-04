#!/usr/bin/env sh

if [ "$LOG_LEVEL" == "debug" ]; then
	echo $GPG_SECRET | base64 -d | gpg --armor --import
else
	echo $GPG_SECRET | base64 -d | gpg --armor --import &> /dev/null
fi

mkdir -p /home/root/.config/gopass

cat << EOF > /home/root/.config/gopass/config
[mounts]
	path = /rundir/secrets
EOF

if [ "$LOG_LEVEL" == "debug" ]; then
	gopass clone --check-keys=false $SECRETS_REPO_URL
else
	gopass clone --check-keys=false $SECRETS_REPO_URL &> /dev/null
fi

./gopass-external-secrets
