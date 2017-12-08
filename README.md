# strdist [![Travis-CI](https://travis-ci.org/RadekD/strdist.svg)](https://travis-ci.org/RadekD/strdist) [![GoDoc](https://godoc.org/github.com/RadekD/strdist?status.svg)](https://godoc.org/github.com/RadekD/strdist)

Golang implementation of Levenshtein distance algorithm

More info: https://en.wikipedia.org/wiki/Levenshtein_distance


`go get github.com/RadekD/strdist`

## Usage
```go
import "github.com/RadekD/strdist"

func main() {
    distance := strdist.Levenshtein(source, target)
    if distance > 0 && distance < 3 {
        //typo?
    }
}
```

## Licence

GNU General Public License