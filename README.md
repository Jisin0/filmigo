<p align="center">
    <a href="https://github.com/Jisin0/filmigo">
        <img src="https://github.com/Jisin0/filmigo/blob/main/filmigo-logo.png" alt="filmigo" width="512">
    </a>
    <br>
    <b>Golang API Wrapper for Imdb Omdb and JustWatch</b>
    <br>
    <a href="https://github.com/Jisin0/filmigo">
        Homepage
    </a>
    •
    <a href="https://github.com/Jisin0/filmigo/tags">
        Releases
    </a>
    •
    <a href="https://pkg.go.dev/github.com/Jisin0/filmigo">
        Documentation
    </a>
</p>

# filmigo
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/Jisin0/filmigo)](https://goreportcard.com/report/github.com/Jisin0/filmigo) [![Go Reference](https://pkg.go.dev/badge/github.com/Jisin0/filmigo.svg)](https://pkg.go.dev/github.com/Jisin0/filmigo)

**filmigo** is a library designed to make accessing movie databases as simple as possible in **Go**. 
It provides coverage of many api methods with close to no configuration and in-built configurable caching.

## Installation
Install the latest version of **filmigo** from github.
```bash
go get github.com/Jisin0/Filmigo@latest
```

## Quickstart
All the examples below fetch full data about the movie [Inception](https://www.imdb.com/title/tt1375666) with
it's imdbId or JustWatchId respectfully assuming there are no errors. Detailed examples can be found at the 
root of each package.

### Imdb
```go
import "github.com/Jisin0/filmigo/imdb"

client := imdb.NewClient()

func main() {
   movie, _ := client.GetMovie("t1375666")
   movie.PrettyPrint()
}
```
[_More Examples_](imdb/)

### Omdb
```go
import "github.com/Jisin0/filmigo/omdb"

client := omdb.NewClient("your_api_key")

func main() {
   movie, _ := client.GetMovie("t1375666")
   movie.PrettyPrint()
}
```
[_More Examples_](omdb/)

### JustWatch
The justwatch id of a title can inly be obtained from search results. You can alsu use the justwatch url
of the movie which is more common within justwatch.
```go
import "github.com/Jisin0/filmigo/justwatch"

client := justwatch.NewClient()

func main() {
   movie, _ := client.GetTitle("tm92641")
   movie.PrettyPrint()
}
```
[_More Examples_](justwatch/)

## Disclaimer
- This product is only for educational purposes and is not meant for commercial usage .
- This product uses the apis of imdb, omdb and juswatch but is by no means endorsed or certified by any of them.
- This product uses apis not intended for public use.
- This product **does not** use justwatch's official [partners api](https://www.justwatch.com/us/JustWatch-Streaming-API).
- This product comes with **no warranty**.

Please read the privacy policy of the respective platform before using this product.


## Supported Sites

- [IMDB](https://imdb.com)
- [JustWatch](https://justwatch.com)
- [OMDB](https://omdapi.com)

[TMDB](https://themoviedatabase.org) is not currently supported as there is already [go-tmdb](https://github.com/ryanbradynd05/go-tmdb/) 
which provides full coverage of all api methods.

## Contributing
**filmigo** is an open-source project and your contribution is very much appreciated. If you'd like to contribute, simply fork the repository, commit your changes and send a pull request. If you have any questions, feel free to ask.

## License
The MIT License (MIT)

(c) Jisin0

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.