# calc
Fork of goadesign/examples basic (calc) example, with added swagger ui

On end of design.go this line

```
Files("/swagger.json", "../../gen/http/openapi.json")
```

is changed with this lines

```
Files("/ui/swagger.json", "./gen/http/openapi.json")
// Serve swagger ui
Files("/ui/{*filepath}", "ui/")
```