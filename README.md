# go_examples
Go Lang Examples

## Setup GOlang
Go to [golang.org](https://golang.org/) to download and install go.  See the instructions
on the site.  But you should do something like the following.
```
wget https://golang.org/dl/go1.17.linux-amd64.tar.gz
mkdir /usr/local/go 
sudo tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz
echo export PATH=/usr/local/go/bin:$PATH >> ~/.bashrc 
```

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

## Pro GOlang
* Compiled
* Concurrency built in
* designed to be readable and usable
* can compile for different OS
* philosophy is "one problem one solution"


