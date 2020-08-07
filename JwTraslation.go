package jwParser

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func init() {
	getGenreFromJW()
	getProviderFromJW()
}

func (cw *Optimus)JwTranslate(data []byte) {
	//if offers tag present then on do all this, if not just skip the content
	if !gjson.GetBytes(data, "offers").Exists() {
		return
	}

	//make media
	makingJWMedia(data, cw.Media)
	println("MEDIA**************************************************")
	println(cw.Media.String())
	println("\n\n\n")
	// make content
	makingJwContent(data, cw.Content)
	println("CONTENT**************************************************")
	println(cw.Content.String())
	println("\n\n\n")

	//make metadata
	makingJwMetadata(data, cw.Metadata)

	println("METADATA**************************************************")
	println(cw.Metadata.String())
	println("\n\n\n")

	//make content Avaliable
	makingJWContentAvliable(data, &cw.ContentAvailable)

	println("Content Avalible**************************************************")
	for _, aval := range cw.ContentAvailable {
		println("\n")
		println(aval.String())
	}
}

//perfect and freez
func makingJWMedia(data []byte, media *Media){
	//making media
	//init
	landscape := []string{}
	portrait := []string{}
	backdrop := []string{}
	banner := []string{}
	video := []string{}

	// landscape
	if gjson.GetBytes(data, "poster").Exists() {
		poster := gjson.GetBytes(data, "poster").String()
		poster = strings.Replace(poster, "{profile}", "s592/movie.webp", -1)
		portrait = append(portrait, "https://images.justwatch.com"+poster)
	}

	//backdrop landscape
	if gjson.GetBytes(data, "backdrops.#.backdrop_url").Exists() {
		for _, r := range gjson.GetBytes(data, "backdrops.#.backdrop_url").Array() {
			//get the name of the content
			nameOfTile := gjson.GetBytes(data, "full_path").String()
			nameOfTile = strings.Replace(nameOfTile, "/in/movie", "", -1)
			nameOfTile = strings.Replace(nameOfTile, "/in/tv-show", "", -1)
			image := strings.Replace(r.String(), "{profile}", "s1440"+nameOfTile+".webp", -1)
			backdrop = append(backdrop, "https://images.justwatch.com"+image)
			landscape = append(landscape, "https://images.justwatch.com"+image)
		}
	}

	//video
	for _, r := range gjson.GetBytes(data, `clips.#(provider="youtube").external_id`).Array() {
		video = append(video, "https://www.youtube.com/watch?v="+r.String())
	}

	media.Landscape = landscape
	media.Portrait = portrait
	media.Backdrop = backdrop
	media.Banner = banner
	media.Video  = video
}

//perfect and freez
func makingJwContent(data []byte, content *Content){
	var sources []string
	for _, r := range gjson.GetBytes(data, `offers.#.provider_id`).Array() {
		for k, v := range jwProvidersMap {
			if v == r.Int() {
				contains := false
				tmp := jwProviderSourceMap[k]
				for _, s := range sources {
					if s == tmp {
						contains  = true
						break
					}
				}
				if !contains{
					sources = append(sources, tmp)
				}
				break
			}
		}
	}
	content.Sources = sources
	content.DetailPage = true
	content.PublishState = true
}

// perfect and freeze
func makingJwMetadata(data []byte, metadata *Metadata){

	metadata.Title = gjson.GetBytes(data, `original_title`).String()
	metadata.Synopsis = gjson.GetBytes(data, `short_description`).String()
	metadata.Year = int32(gjson.GetBytes(data, "original_release_year").Int())
	metadata.Mood = []int32{}
	metadata.ReleaseDate = fmt.Sprintf("05-01-%d", metadata.GetYear())
	metadata.ImdbId = gjson.GetBytes(data, "external_ids.#(provider=imdb).external_id").String()
	metadata.Country = []string{"INDIA"}
	metadata.Rating = gjson.GetBytes(data, `tmdb_popularity`).Float()
	metadata.Episode = -1
	metadata.Part = -1
	metadata.Season = -1
	metadata.Runtime = gjson.GetBytes(data, `runtime`).String()
	metadata.ViewCount = 0.0


	//prefs
	metadata.Categories = []string{}
	metadata.Languages = []string{}
	metadata.Genre = []string{}
	metadata.Directors = []string{}
	metadata.Cast = []string{}



	//genre
	for _, r := range gjson.GetBytes(data, "genre_ids").Array() {
		metadata.Genre = append(metadata.Genre, jwGenre[r.Int()])
	}

	//Director
	for _, r := range gjson.GetBytes(data, `credits.#(role="DIRECTOR")#.name`).Array() {
		metadata.Directors = append(metadata.Directors, r.String())
	}

	//actor
	for _, r := range gjson.GetBytes(data, `credits.#(role="ACTOR")#.name`).Array() {
		metadata.Cast = append(metadata.Cast, r.String())
	}

	//content type
	for _, r := range gjson.GetBytes(data, `object_type`).Array() {
		if r.String() == "movie" {
			metadata.Categories = append(metadata.Categories, "Movie")
		}else {
			metadata.Categories = append(metadata.Categories, "Tv Show")
		}
	}

	// languages
	if gjson.GetBytes(data, `offers.#.audio_languages`).Exists() {
		for _, r := range gjson.GetBytes(data, `offers.#.audio_languages`).Array() {
			metadata.Categories = append(metadata.Categories, r.String())
		}
	}

	//kidsSafe
	if len(metadata.GetGenre()) > 0 {
		for _, genre := range metadata.GetGenre() {
			if genre == "Kids & Family" {
				metadata.KidsSafe = true
			} else {
				metadata.KidsSafe = false
			}
		}
	}

	//merging genre and categories
	metadata.Tags = append([]string{}, append(metadata.Genre, metadata.Categories...)...)
}

// perfect and freeze
func makingJWContentAvliable(data []byte , avaliable *[]*ContentAvailable) {

	for _, r := range gjson.GetBytes(data, "offers").Array() {

		var ca ContentAvailable
		ca.Type = "CW_THIRDPARTY"
		ca.Monetize = Monetize_Free
		ca.TargetId = gjson.GetBytes(data, "jw_entity_id").String()
		ca.Target = r.Get("urls.standard_web").String()


		switch r.Get("monetization_type").String() {
							case "buy":
							case "flatrate":
								{
									ca.Monetize = Monetize_Paid
								}
								break
							case "rent":
								{
									ca.Monetize = Monetize_Rent
								}
								break
							}

		for k, v := range jwProvidersMap {
			if v == r.Get("provider_id").Int() {
				ca.Source = jwProviderSourceMap[k]
				ca.Package = GetPackageNameForJW(k)
				break
			}
		}
		if !itemExists(*avaliable, &ca) {
			*avaliable = append(*avaliable, &ca)
		}
	}
}








