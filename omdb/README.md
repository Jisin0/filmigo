# Omdb
The omdb package is an api wrapper for omdbapi.com.
[_Generated Docs_](https://pkg.go.dev/github.com/Jisin0/filmigo/omdb)

## Table Of Content
- [Setup](#setup)
- [Search](#search)
- [Fetch Movie](#getmovie)

## Guide
Here's a short guide of the available methods and it's usage. All options are passed in the optional field of each function.

### Setup
Let's start by importing the omdb package
```go
import "github.com/Jisin0/filmigo/omdb"
```

Now let's create a new omdb client with your api key. Get you api key [here](https://www.omdbapi.com/apikey.aspx).
```go
client := omdb.NewClient("your_api_key")
```
**Options**
- DisableCaching : Indicates wether data should not be cached.
- CacheExpiration : Duration for which cache is valid. Defaults to 5 hours.

### Search
You can search for titles i.e Movies and Shows using the Search method.
```go
client.Search("inception")
```
**Options**
- Type : Type of result to return either "movie", "series" or "episode".
- Year : Year of release of the movie.
- Page : Results page to return.

### GetMovie
You can fetch a movie by it's imdb id or it's exact title. Either ID or Title field must be set in request options.
```go
opts := omdb.GetMovieOpts{
    Title: "inception",
}
client.GetMovie(&opts)
```
**Options**
- ID : Imdb id of the title.
- Title : Exact title of the movie.
- Type : Type of result to return either "movie", "series" or "episode".
- Year : Year of release of the movie.
- Plot : Length of plot to return, "short" for a short plot or "full" for the full plot.