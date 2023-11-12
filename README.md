# Go Web Templates - API

An Framework agnostic REST Api template. This repository is part of a collection of Golang web templates, for more information on how to use and modify these templates check the [main repository]() and the [wiki]().

## How to use this template:

Click in the the `Use this template` at the top of this repository and create a new repository with this one as
template.
![Screenshot from 2023-11-12 17-48-36](https://github.com/go-web-templates/api/assets/46976272/883a0682-dfc7-47d4-9d60-ca93f8f38982)

After creating a repository, you need to configure your own go module name. At the root of this repetition you can
find a script called `setup.sh`. This script will change the current module name to one of your choice.

```bash
# Clone your new repo
git clone git:@github.com/YOUR_USER/YOUR_REPO.git #for ssh

# Setup the go module:
cd ./YOUR_REPO
./setup.sh github.com/YOUR_USER/YOUR_REPO

# Commit the changes:
git add . && git commit -m "New go module name"
```
