# todoapp
A simple starting boilerplate code that uses gorm and echo

```bash
# To run the app
$ go run .

# To test the app
$ go test .

# To test the app from curl
curl -XGET    -H 'Content-Type:application/json'                                         localhost:8080/api/v1/todos
curl -XPOST   -H 'Content-Type:application/json' -d '{"title": "hello1"}'                localhost:8080/api/v1/todos
curl -XPOST   -H 'Content-Type:application/json' -d '{"title": "hello2", "done": false}' localhost:8080/api/v1/todos
curl -XPUT    -H 'Content-Type:application/json' -d '{"title": "hello2", "done": true}'  localhost:8080/api/v1/todos/2
curl -XDELETE -H 'Content-Type:application/json'                                         localhost:8080/api/v1/todos/2
```
