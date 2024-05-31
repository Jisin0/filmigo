// (c) Jisin0
// Object types and interfaces.

package justwatch

import "strings"

// // TitleContent implements an interface between the different title type's content.
// type TitleContent interface {
// 	// Title of the movie or show .
// 	Title() string
// 	// Full url path of the movie or show .
// 	FullPath() string
// 	// Original relase year of the movie or show.
// 	OriginalReleaseYear() string
// 	// Poster url format of the movie or show.
// 	Poster() PosterURL
// }

// Full details of a url obtained from using GetTitleFromURL()
type URLDetails struct {
	// graphql object id.
	ID string `json:"id,omitempty"`
	// Meta description is a short sentence with the available streaming provider names and a greeting.
	MetaDescripting string `json:"metaDescription,omitempty"`
	// Meta keyword about the title concatanated with commas.
	MetaKeywords string `json:"metaKeywords,omitempty"`
	// Meta title.
	MetaTitle string `json:"metaTitle,omitempty"`
	// Primary heading.
	Heading1 string `json:"heading1,omitempty"`
	// Secondary heading.
	Heading2 string `json:"heading2,omitempty"`
	// Raw html content with description and heading.
	HTMLContent string `json:"htmlContent,omitempty"`
	// Full data on the entity.
	Data *Title `json:"node,omitempty"`
}

// Data about any title on justwatch.
type Title struct {
	// Justwatch id of the title. for ex: ts20233.
	ID string `json:"id,omitempty"`
	// Type of title either MOVIE, SHOW, SHOW_SEASON or SHOW_EPISODE.
	Type string `json:"objectType,omitempty"`
	// Numeric id of the title for ex: ts20233 becomes 202333.
	NumericID int `json:"objectID,omitempty"`
	// Data or content about the title.
	Content *TitleContent `json:"content,omitempty"`
	// Total number of available offers.
	OfferCount int `json:"offerCount,omitempty"`
	// Total number of unique offers.
	UniqueOfferCount int `json:"uniqueOfferCount,omitempty"`
	// All available offers for a title.
	Offers []*Offer `json:"offers,omitempty"`
	// WatchNowOffer is a direct view offer.
	WatchNowOffer *Offer `json:"watchNowOffer,omitempty"`
	// Promoted Bundles.
	PromotedBundles []*PromotedBundle `json:"promotedBundles,omitempty"`
	// Availability of offers.
	AvailableTo []*AvailableTo `json:"availableTo,omitempty"`
	// Offers for the title on appleTV/Itunes.
	AppleOffers []*Offer `json:"appleOffers,omitempty"`
	// Offers for the title on plexplayer.
	PlexPlayerOffers []*Offer `json:"plexPlayerOffers,omitempty"`
	// Full timestamp of when offers were updated.
	MaxOfferUpdateAt string `json:"maxOfferUpdatedAt,omitempty"`
	// Number of offers availbe on disney.
	DisneyOffersCount int `json:"disneyOffersCount,omitempty"`
	// Number of offers availbe on hotstar.
	StarOffersCount int `json:"starOffersCount,omitempty"`
	// Popularity rank of the title.
	PopularityRank *PopularityRank `json:"popularityRank,omitempty"`
	// Streaming charts and trends of the movie.
	StreamingCharts struct {
		Edges []*StreamingChartInfo `json:"edges,omitempty"`
	} `json:"streamingCharts,omitempty"`
	// Total number of seasons only for shows.
	TotalSeasonCount int `json:"totalSeasonCount,omitempty"`
	// Total Episode count only retirned for show seasons.
	TotalEpisodeCount int `json:"totalEpisodeCount,omitempty"`
	// Details about the show only present for season type.
	Show struct {
		// Justwatch id of the show.
		ID string `json:"id,omitempty"`
		// Numeric value of the id.
		NumericID int `json:"objectID,omitempty"`
		// Type of the object in this case always SHOW.
		Type string `json:"objectType,omitempty"`
		// Content of the show.
		Content struct {
			Title string `json:"title,omitempty"`
		} `json:"content,omitempty"`
	} `json:"show,omitempty"`
	// Seasons of a show only returned for type SHOW.
	Seasons []*SeasonPreview `json:"seasons,omitempty"`
	// Recent episode of a show only returned for SHOW or SHOW_SEASON types.
	RecentEpisodes []*EpisodePreview `json:"recentEpisodes,omitempty"`
}

// Data about the actual content of the title like it's genres, poster, runtime etc.
type TitleContent struct {
	// Name of the title.
	Title string `json:"title,omitempty"`
	// Original name of the title.
	OriginalTitle string `json:"originalTitle,omitempty"`
	// URL path of the title.
	URLPath string `json:"fullPath,omitempty"`
	// A short description of the title .
	Description string `json:"shortDescription,omitempty"`
	// Year in which the title was released.
	ReleaseYear int `json:"originalReleaseYear,omitempty"`
	// Date on which the title was first released in the format yyyy-mm-dd
	ReleaseDate string `json:"originalReleaseDate,omitempty"`
	// Genres of the title like comedy, romance etc.
	Genres *Genres `json:"genres,omitempty"`
	// Unformatted poster url path.
	Poster *PosterURL `json:"posterURL,omitempty"`
	// Formatted poster url path.
	FullPoster *PosterURL `json:"fullPosterURL,omitempty"`
	// Runtime or duration of the title in minutes.
	Runtime int `json:"runtime,omitempty"`
	// Age certification for ex: TV-MA.
	AgeCertification string `json:"ageCertification,omitempty"`
	// Scores and ratings for the title on different platforms
	Scores *Scoring `json:"scoring,omitempty"`
	// IsReleased indicates wether the title is released in theaters or online.
	IsReleased bool `json:"isReleased,omitempty"`
	// Credits to people in the title.
	Credits []*Credit `json:"credits,omitempty"`
	// Interactions details like likes and dislikes.
	Interactions *Interactions `json:"interactions,omitempty"`
	// Season number for shows and episodes.
	SeasonNumer int `json:"seasonNumber,omitempty"`
	// Episode number only for show episodes.
	EpisodeNumber int `json:"episodeNumber,omitempty"`
	// Backdrop or banner images for the title.
	Backdrops []*Backdrop `json:"backdrops,omitempty"`
	// Full backdrops returns the fully formatted url path of backdrops.
	FullBackdrops []*Backdrop `json:"fullBackdrops,omitempty"`
	// Video clips of the title, usually only yt links are provided in this field.
	Clips []*Clip `json:"clips,omitempty"`
	// Clips from videobuster.com.
	VideobusterClips []*Clip `json:"videobusterClips,omitempty"`
	// Clips from dailmotion.com . Justwatch gives priority to these on their own website.
	DailymotionClips []*Clip `json:"dailymotionClips,omitempty"`
	// External ids of the title on imdb and tmdb.
	ExteranlIDs *ExteranlIDs `json:"externalIDs,omitempty"`
	// Countries in which the movie was produced.
	ProductionCountries []string `json:"productionCountries,omitempty"`
}

// A single episode of a show.
type Episode struct {
	// Justwatch id of the episode. for ex: tse7685834.
	ID string `json:"id,omitempty"`
	// Numeric value of the id.
	NumericID int `json:"objectID,omitempty"`
	// Content of the episode.
	Content struct {
		// Title of the episode.
		Title string `json:"episode,omitempty"`
		// Description of the episode.
		Description string `json:"shortDescription,omitempty"`
		// Season number.
		SeasonNumber int `json:"seasonNumber,omitempty"`
		// Number of the episode.
		EpisodeNumber int `json:"episodeNumber,omitempty"`
		// Indicates wether the episode is released.
		IsReleased bool `json:"isReleased,omitempty"`
	} `json:"content,omitempty"`
}

// Results from GetTitleOffers() query.
type GetTitleOffersResult struct {
	// Justwatch ID of the title.
	ID string `json:"id"`
	// Type name either Movie or Show.
	TypeName string `json:"__typename"`
	// Number of offers available for the title.
	OfferCount int `json:"offerCount"`
	// Full timestamp of the last time offers were updated.
	MaxOfferUpdateAt string `json:"maxOfferUpdatedAt"`
	// Flatrate offers for the title.
	Flatrate []*Offer `json:"flatrate"`
	// Offers to buy the movie/show.
	Buy []*Offer `json:"buy"`
	// Offers to rent the title for a period.
	Rent []*Offer `json:"rent"`
	// Offers to watch the title for free.
	Free []*Offer `json:"free"`
	// Fast-fiew offers.
	Fast []*Offer `json:"fast"`
}

// Any streaming provider's offer for a movie or show.
type Offer struct {
	// Graphql id of the object.
	ID string `json:"id,omitempty"`
	// Direct web url to the offer package.
	URL string `json:"standardWebURL,omitempty"`
	// Type indicates the type of offer .
	Type string `json:"type,omitempty"`
	// Retail price of the provider. The string is a well formatted price tag with the currency of the request country.
	RetailPrice string `json:"retailPrice,omitempty"`
	// MonetizationType indicates how you pay for the movie/show. Values are BUY, FLATRATE, RENT or FREE.
	MonetizationType string `json:"monetizationType,omitempty"`
	// RetailPriceValue is the actual numeric value without the currency tag.
	RetailPriceValue float32 `json:"retailPriceValue,omitempty"`
	// Currency of the offer. for ex: USD, GBP, INR etc.
	Currency string `json:"currency,omitempty"`
	// PresentationType is the type of video quality either HD or SD.
	PresentationType string `json:"presentationType,omitempty"`
	// LastChangeRetailPriceValue is the numeric value of the price before the last change in it.
	LastChangeRetailPriceValue float32 `json:"lastChangeRetailPriceValue,omitempty"`
	// Details about the offer provider.
	Package *Package `json:"package,omitempty"`
	// Elements count in the results.
	ElemCount int `json:"elementCount,omitempty"`
	// Available to

	// Deeplink path from the justwatch site to the offer.
	Deeplink string `json:"deeplinkRoku,omitempty"`
}

// Basic details about an offer package.
type Package struct {
	// Graphql id of the type
	ID string `json:"id,omitempty"`
	// ID of the package.
	PackageID int `json:"packageID,omitempty"`
	// Clear user friendly name of the package for ex: Apple TV.
	ClearName string `json:"clearName,omitempty"`
	// Technical name of the package for ex: itunes.
	TechnicalName string `json:"technicalName,omitempty"`
	// URL path to the icon of the package.
	Icon *Icon `json:"icon,omitempty"`
	// Shortname of the package.
	ShortName string `json:"shortName,omitempty"`
}

// Basic data about a season that's included in the full result of a show.
type SeasonPreview struct {
	// Justwatch id of the season for ex: tss414472.
	ID string `json:"id,omitempty"`
	// Numeric id.
	NumericID int `json:"objectID,omitempty"`
	// Type will be SHOW_SEASON.
	Type string `json:"objectType,omitempty"`
	// Number of episodes in the season.
	TotalEpisodeCount int `json:"totalEpisodeCount,omitempty"`
	// Availability.
	AvailableTo []*AvailableTo `json:"availableTo,omitempty"`
	// Content of the season, only few fields are populated.
	Content *TitleContent `json:"content,omitempty"`
}

// Basic data about an episode that's included in the full result of a show or season.
type EpisodePreview struct {
	// Justwatch id of the season for ex: tss414472.
	ID string `json:"id,omitempty"`
	// Numeric id.
	NumericID int `json:"objectID,omitempty"`
	// Number of episodes in the season.
	TotalEpisodeCount int `json:"totalEpisodeCount,omitempty"`
	// Availability.
	AvailableTo []*AvailableTo `json:"availableTo,omitempty"`
	// Content of the episode, only few fields are populated.
	Content *TitleContent `json:"content,omitempty"`
}

// Credit for a role in a title.
type Credit struct {
	// role of the person in the title.
	Role string `json:"role,omitempty"`
	// Name of the person.
	Name string `json:"name,omitempty"`
	// Name of the character played by the person.
	CharacterName string `json:"characterName,omitempty"`
	// Justwatch ID of the person
	ID int `json:"personID,omitempty"`
}

// Video clip of the title.
type Clip struct {
	// URL to the video .
	URL string `json:"sourceURL,omitempty"`
	// Name of the clip.
	Name string `json:"name,omitempty"`
	// Provider of the clip probably YOUTUBE or DAILYMOTION.
	Provider string `json:"provider,omitempty"`
	// Exteranl id of the clip on their respective platforms.
	ExteranlID string `json:"externalID,omitempty"`
}

// Sores or ratings for the title on various platforms.
type Scoring struct {
	// Imdb rating out of 10.
	ImdbRating float32 `json:"imdbScore,omitempty"`
	// Votes received on imdb
	ImdbVotes float32 `json:"imdbVotes,omitempty"`
	// Popularity rating on tmdb.
	TmdbPopularity float32 `json:"tmdbPopularity,omitempty"`
	// Rating on tmdb out of 10.
	TmdbRating float32 `json:"tmdbScore,omitempty"`
	// Rating percentage on justwatch for ex: 0.8778742982 indicates 87.8%
	JustwatchRating float32 `json:"jwRating,omitempty"`
}

// Interactions with the title like likes and dislikes details.
type Interactions struct {
	// Dislikes number.
	Dislikes int `json:"dislikelistAdditions,omitempty"`
	// Likes number.
	Likes int `json:"likelistAdditions,omitempty"`
	// Number of votes (likes + dislikes).
	Votes int `json:"votesNumber,omitempty"`
}

// Promoted bundle.
type PromotedBundle struct {
	// Direct URL to the bundle.
	URL string `json:"promotionURL,omitempty"`
}

// Availability of the movie or show.
type AvailableTo struct {
	// Countdown integer.
	AvailableCountdown int `json:"availableCountDown,omitempty"`
	// Date upto which it's available in the format yyyy-mm-dd.
	AvalableToDate string `json:"availableToDate,omitempty"`
	// Package details with only shortname populated.
	Package *Package `json:"package,omitempty"`
}

// Popularity rank of a title.
type PopularityRank struct {
	// Rank number.
	Rank int `json:"rank,omitempty"`
	// Trend indicates the stability of the rank.
	Trend string `json:"STABLE,omitempty"`
	// Trend difference.
	TrendDifference int `json:"trendDifference,omitempty"`
}

// Info about a title on a streaming chart.
type StreamingChartInfo struct {
	// Rank on the streaming chart.
	Rank int `json:"rank,omitempty"`
	// Trend indicates the stability of the rank either STABLE, UP or DOWN.
	Trend string `json:"trend,omitempty"`
	// Net difference in the trend.
	TrendDifference int `json:"trendDifference,omitempty"`
	// Full timestamp of the last time the rank was updated.
	UpadtedAt string `json:"updatedAt,omitempty"`
	// Maximum rank the title has ever achieved.
	TopRank int `json:"topRank,omitempty"`
	// Days for which the title was in the top 3.
	DaysInTop3 int `json:"daysInTop3,omitempty"`
	// Days for which the title was in the top 10.
	DaysInTop10 int `json:"daysInTop10,omitempty"`
	// Days for which the title was in the top 100.
	DaysInTop100 int `json:"daysInTop100,omitempty"`
	// Days for which the title was in the top 1000.
	DaysInTop1000 int `json:"daysInTop1000,omitempty"`
}

// External IDs for the movie on imdb, tmdb etc.
type ExteranlIDs struct {
	// Imdb id.
	ImdbID string `json:"imdbID,omitempty"`
	// Tmdb id.
	TmdbID string `json:"tmdbID,omitempty"`
}

type Icon string

// returns the direct url to the icon.
func (i Icon) FullURL() string {
	return imageBaseURL + strings.TrimSuffix(string(i), ".{format}")
}
