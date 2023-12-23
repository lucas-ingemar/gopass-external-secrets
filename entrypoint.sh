#!/usr/bin/env sh

echo $GPG_SECRET | base64 -d | gpg --armor --import

mkdir -p /home/root/.config/gopass

git config --global user.email "$GIT_USER_EMAIL"
git config --global user.name "$GIT_USER_NAME"

echo "Cloning secrets repo"
git clone $SECRETS_REPO_URL /rundir/secrets

cat << EOF > /home/root/.config/gopass/config
[mounts]
	path = /rundir/secrets
EOF

PASSWORD_STORE_DIR=/rundir/secrets gopass init $GPG_KEY_ID

./gopass-external-secrets
