# Random project name generator

This is a random project name generator, written in Go. There
are many libraries for this task, but most only reproduce the
same [list of 50 adjectives and 50 words][list] from [this
gist](https://gist.github.com/afriggeri/1266756). This means there are only
2500 combinations, and you're sharing with every other project using those
lists. Other Go libraries store the list in a JSON file, which means end users
need to figure out how to find/load the JSON file at runtime.

This project exposes the [project-name-generator][project-name-generator] list,
which is much more extensive - 1300 adjectives and 870 nouns means there are
over a million possible combinations.

Usage:

```go
package main

import (
    "github.com/kevinburke/go-random-project-generator"
    "fmt"
)

func main() {
    name := random_project_generator.Generate()
    fmt.Println(name) // "tidy-boat"
    name = random_project_generator.GenerateNumber(4)
    fmt.Println(name) // "viable-action-2183"
}
```

You can bring your own list as well, overwrite `random_project_generator.Nouns`
or `random_project_generator.Adjectives` in your init function.

[project-name-generator]: https://github.com/aceakash/project-name-generator
[list]: https://github.com/usmanbashir/haikunator/blob/master/lib/haikunator.rb#L31-L53
