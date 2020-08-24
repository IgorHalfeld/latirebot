<p align="center">
  <img src="latire.png" width="150" />
  <h1 align="center">Latire</h1>
  <p align="center">
    Um CRON feito em Golang que procura roupas masculinas na riachuelo <br />
    e manda mensagem no telegram quando tem promoção
  </p>
</p>


### Todo

[x] Aceitar mais de uma loja (Renner, Dafiti)
[x] Aceitar roupas femininas também (foi masculinas inicialmente pq é pra mim haha)
[] Aceitar o usuário setar o tamanho do desconto
[x] Mandar pro usuário que mandou o `/start` no bot
[] Testes

### Var env

```sh
TELEGRAM_KEY=MY_KEY
```

### Como rodar 

```sh
make migration-up

make run
```
