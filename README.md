# What is this project for?

This is for github open source repositories to generate weekly/monthly reports, into markdown.

The markdown template is currently hardcoded and is in Chinese.

# Usage

1. Config your github token in git config

Before run this program, you need to config `personal_access_token` in your command line git config:

```
git config \
  --global \
  url."https://${user}:${personal_access_token}@github.com".insteadOf \
  https://github.com
```

The program will read token from git config from this line of '.insteadOf'

```
git config -l
```


2. run progress
```
go run main.go
```
