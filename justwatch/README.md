# Justwatch
The justwatch package offers methods to browse justwatch and get data about entities using it's graphql api.

## Table Of Content
- [Setup](https://github.com/Jisin0/filmigo/justwatch#setup)
- [Search](https://github.com/Jisin0/filmigo/justwatch#search)
- [Fetch Movie](https://github.com/Jisin0/filmigo/justwatch#getmovie)
- [Fetch Title Offers](https://github.com/Jisin0/filmigo/justwatch#getoffers)

## Guide
Here's a short guide of the available methods and it's usage. All options are passed in the optional field of each function.

### Setup
Let's start by importing the justwatch package
```go
import "github.com/Jisin0/filmigo/justwatch"
```

Now let's create a new justwatch client, all methods are called through this client.
```go
client := justwatch.NewClient()
```
**Options**
- Country : Country code for the source country deafults to US.
- LangCode : Language code defaults to en.

### Search
You can search for titles i.e Movies and Shows using the SearchTitles method.
```go
client.SearchTitles("inception")
```
**Options**
- Limit : Maximimum number of results to return defaults to 5.
- NoTitlesWithoutURL : Indicates wether titles without url should not be returned.
- Country : Use a country code for the specific request (uses client's country by default).
- LangCode : Use a language code for the specific request (uses client's LangCode by default).

### Fetch Movie {#getmovie}
You can fetch a title by it's id or it's url. Justwatch ids are only used internally unlike imdb that exclusively use this to identify titles.

#### By ID
Use this function to get a movie using it's justwatch id. We'll use [tm820952](https://www.justwatch.com/us/movie/inception) for this example.
```go
client.GetTitle("tt1375666")
```
**Options**
- EpisodeMaxLimit : Maximimum number of episodes to return for a show season defaluts to 20.
- Country : Use a country code for the specific request (uses client's country by default).
- LangCode : Use a language code for the specific request (uses client's LangCode by default).

#### By URL
Use this function to get a movie using it's justwatch link. We'll use https://www.justwatch.com/us/movie/inception for this example.
The function also accepts just the url path i.e /us/movie/inception which is the norm.
```go
client.GetTitleFromURL("https://www.justwatch.com/us/movie/inception")
```
**Options**
- EpisodeMaxLimit : Maximimum number of episodes to return for a show season defaluts to 20.
- Country : Use a country code for the specific request (uses client's country by default).
- LangCode : Use a language code for the specific request (uses client's LangCode by default).

### Get Offers {#getoffers}
Use this function to get offers for a title using it's justwatch id. We'll use [tm820952](https://www.justwatch.com/us/movie/inception) for this example.
```go
client.GetTitleOffers("tm820952")
```
**Options**
- Country : Use a country code for the specific request (uses client's country by default).
- LangCode : Use a language code for the specific request (uses client's LangCode by default).