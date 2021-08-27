# go_examples
Go Lang Examples


## GitLab Setup
Setup your Persnal Access Token in GitLab
* Settings => Developer Settings =>  Personal Access Tocken.
* Select Generate Token and save the token 

From command line git setup your global settings
$ git config --global user.name "your_github_username"
$ git config --global user.email "your_github_email"
$ git config -l

git clone <your repo>

Clone your repo, provide your user name and Personal Access Token then 
cache your credentials.

git clone <your repo>
git config --global credential.helper cache

To unset you credentials do the following

$ git config --global --unset credential.helper
$ git config --system --unset credential.helper

verify your credentials 
git pull -v

Go Examples
---
[hello.go](./bin/hello)
