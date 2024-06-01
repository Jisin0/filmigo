# Imdb

The imdb package offers methods to browse imdb and get data about entities.
[_Generated Docs_](https://pkg.go.dev/github.com/Jisin0/filmigo/imdb)

## Table Of Content
- [Setup](#setup)
- [Search](#search)
- [Fetch Movie](#getmovie)
- [Fetch Person](#getperson)
- [Advanced Search](#advanced-search)

## Guide
Here's a short guide of the available methods and it's usage. All options are passed in the optional field of each function.

### Setup
Let's start by importing the imdb package
```go
import "github.com/Jisin0/filmigo/imdb"
```

Now let's create a new imdb client, all methods are called through this client.
```go
client := imdb.NewClient()
```
**Options**
- DisableCaching : Indicates wether data should not be cached.
- CacheExpiration : Duration for which cache is valid. Defaults to 5 hours.

### Search
These functions use imdb's api and are superfast. You can search for titles. names or both.

#### Titles
You can search for titles i.e Movies and Shows using the SearchTitles method.
```go
client.SearchTitles("inception")
```

#### Names
You can search for names using the SearchNames method.
```go
client.SearchNames("keanu")
```

#### Both
You can search for titles and names using the SearchAll method.
```go
client.SearchAll("mad")
```

### GetMovie
Use this function to get a movie using it's imdb id. We'll use [tt1375666](https://www.imdb.com/title/tt1375666) for this example.
```go
client.GetMovie("tt1375666")
```

### GetPerson
Use this function to get a person using their imdb id. We'll use [nm0000206](https://www.imdb.com/name/nm0000206) for this example.
```go
client.GetPerson("nm0000206")
```

### Advanced Search
Imdb offers an advanced search page that allows filtering people and titles based on a wide variety of parameters.

#### Titles
We'll search for all action movies in this examples.

First lets import the constants package.
```go
import "github.com/Jisin0/filmigo/imdb/constants"
```

Now lets create our request options.
```go
opts := imdb.AdvancedSearchTitleOpts{
    Types: []string{constants.TitleTypeMovie},
    Genres: []string{constants.TitleGenreAction},
}
```

Then make the request.
```go
results, err := client.AdvancedSearchTitle(&opts)
if err!=nil{
    panic(err)
}
```

#### Names
We'll search for people who starred in inception using it's id.

Lets create our request options.
```go
opts := imdb.AdvancedSearchNameOpts{
    Titles: []string{"tt1375666"},
}
```

Then make the request.
```go
results, err := client.AdvancedSearchName(&opts)
if err!=nil{
    panic(err)
}
```