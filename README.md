1. config your github token in git config

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

go run main.go
