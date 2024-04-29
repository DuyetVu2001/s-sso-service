go - v1.22.2
sqlc - v1.26.0
migrate - 4.17.1


```sh
# install on window

Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression

scoop install migrate

go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0
```

![alt text](image.png)

TODO: gRPC

- register
- login
- logout
- verify token
- refresh token
- getUserInfoByToken

Questions:

- bscrypt: hash, cost, salt
- base 64