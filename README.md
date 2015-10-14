
# go-utils

Currently this is just a wrapper around github.com/segmentio's 2 fast case parses:

[github.com/segmentio/go-camelcase](https://github.com/segmentio/go-camelcase)

[github.com/segmentio/go-snakecase](https://github.com/segmentio/go-snakecase)

With an added Pascal Case function:

```
func CamelCase(str string) string
func SnakeCase(str string) string
func PascalCase(str string) string
```

Thank you @tj for switching to Go just before we did! ;)
