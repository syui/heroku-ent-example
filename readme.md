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
