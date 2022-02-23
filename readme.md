heroku ent example

- onconflict

```sh
$ curl -X 'POST' -H 'Content-Type: application/json' -d '{"name":"syui"}' "https://$APP.herokuapp.com/users"
...ok

$ !!
...err

$ heroku logs
```

```sh
$ curl "https://$APP.herokuapp.com/users/"
```

ref : 

- https://github.com/ent/ent/blob/master/dialect/sql/schema/postgres_test.go

- https://github.com/go-kratos/beer-shop/tree/main/app/catalog/service/internal/data/ent
