heroku open-api ent example

- go-module-name : t

- onconflict

```sh
$ curl -X 'POST' -H 'Content-Type: application/json' -d '{"name":"syui"}' "https://$APP.herokuapp.com/users"
...ok

$ !!
...err

$ heroku logs
```

```sh
$ curl https://$APP.herokuapp.com/users/1
{"user":"syui", "draw":1}

# card draw
$ curl https://$APP.herokuapp.com/users/1/d
{"user":"syui", "draw":2}
```

ref : 

- https://github.com/ent/ent/blob/master/dialect/sql/schema/postgres_test.go

- https://github.com/go-kratos/beer-shop/tree/main/app/catalog/service/internal/data/ent

- https://entgo.io/ja/blog/2022/02/15/generate-rest-crud-with-ent-and-ogen/

- https://github.com/ariga/ogent/blob/main/example/todo/ent/entc.go
