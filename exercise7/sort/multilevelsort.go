package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// テンプレート用
type TrackViewData struct {
	Tracks  []*Track
	PrevKey string
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"The Gospel", "Alicia Keys", "Here", 2016, length("3m01s")},
	{"fuga", "Moby", "hoge", 1991, length("3m37s")},
	{"bar", "Martin Solveig", "foo", 2016, length("1m00s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

var trackList = template.Must(template.New("tracklist").Parse(`
<h1> tracklist </h1>
<table>
<tr style='text-align: left'>
	<th><a href='http://localhost:8000/?sort=title&prev={{.PrevKey}}'>Title</a></th>
	<th><a href='http://localhost:8000/?sort=artist&prev={{.PrevKey}}'>Artist</a></th>
	<th><a href='http://localhost:8000/?sort=album&prev={{.PrevKey}}'>Album</a></th>
	<th><a href='http://localhost:8000/?sort=year&prev={{.PrevKey}}'>Year</a></th>
	<th><a href='http://localhost:8000/?sort=length&prev={{.PrevKey}}'>Length</a></th>
</tr>
{{range .Tracks}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byAlbum []*Track

func (x byAlbum) Len() int           { return len(x) }
func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }
func (x byAlbum) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// 多段ソート用の構造体
type mlSort struct {
	t       []*Track
	key     sort.Interface
	prevKey sort.Interface
}

func (x mlSort) Len() int {
	return len(x.t)
}

func (x mlSort) Less(i, j int) bool {
	if x.t[i] == x.t[j] && x.prevKey != nil {
		return x.prevKey.Less(i, j)
	}

	return x.key.Less(i, j)
}

func (x mlSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func main() {
	//printTracksForStdout()
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("sort")
	var sortKey sort.Interface
	if value != "" {
		switch value {
		case "title":
			sortKey = byTitle(tracks)
		case "artist":
			sortKey = byArtist(tracks)
		case "album":
			sortKey = byAlbum(tracks)
		case "year":
			sortKey = byYear(tracks)
		case "length":
		default:
			sortKey = nil
		}

		if sortKey != nil {
			sort.Sort(mlSort{tracks, sortKey, nil})
		}
	}

	data := TrackViewData{tracks, value}
	if err := trackList.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func printTracksForStdout() {
	printTracks(tracks)

	//sort.Sort(mlSort{tracks, byYear(tracks), nil})
	sort.Sort(mlSort{tracks, byArtist(tracks), nil})
	print("\n======================\n\n")
	printTracks(tracks)

	//sort.Sort(mlSort{tracks, byArtist(tracks), byYear(tracks)})
	sort.Sort(mlSort{tracks, byYear(tracks), byArtist(tracks)})
	print("\n======================\n\n")
	printTracks(tracks)
}
