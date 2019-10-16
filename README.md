# test

$ gem install travis
$ ssh-keygen -t rsa -b 4096 -C "ah@be.com"

Add to public key to Deploy Keys on the repo

password=$(openssl rand -hex 32)
cat travis_key | openssl aes-256-cbc -k "$password" -md sha256 -a  > travis_key.enc

travis encrypt travis_key_password=$password --add env.matrix

git add travis_key.enc
git add .travis.yml
git commit -m ...

then in build script

openssl aes-256-cbc -k "$travis_key_password" -d -md sha256 -a -in travis_key.enc -out travis_key
echo "Host github.com" > ~/.ssh/config
echo "  IdentityFile  $(pwd)/travis_key" >> ~/.ssh/config
chmod 400 travis_key
