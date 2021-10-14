//package main

//import (
//"html/template"
//"log"
//"net/http"
//"regexp"
//"strings"
//)

//const (
//index = `<h1>Changelog aggregator</h1>
//<ul>
//{{ range . }}
//<li>{{ .Name }}: <a href="/{{ .Name | lower }}@{{ .Version | lower }}">{{ .Version }}</a></li>
//{{ end }}
//</ul>
//<a href="/new"><button type="button"/>new entry</a>
//`
//change = `<h1>{{ .Name }} - {{ .Version }}</h1><hr/>
//{{ .Content }}
//<br />
//<a href="{{ .Source }}">source</a>`

//newEntry = `<h1>New entry<h1><hr>
//<form>

//</form>
//`
//)

//// changelog entry are like /linux@5.10.67
//var changelog = regexp.MustCompile(`.*@.*`)

//type Software struct {
//Name    string
//Version string
//Content string
//Source  string
//}

//type Softwares []Software

//func main() {
//// TODO: fetch subscribed softwares from DB
//softwares := Softwares{
//Software{
//Name:    "Linux",
//Version: "5.10.67",
//Content: "I'm announcing the release of the 5.10.67 kernel.",
//Source:  "https://lwn.net/Articles/869749/",
//},
//Software{
//Name:    "OpenSSL",
//Version: "3.0.0",
//Source:  "https://www.openssl.org/blog/blog/2021/09/07/OpenSSL3.Final/",
//Content: "After 3 years of development work, 17 alpha releases, 2 beta releases, over 7,500 commits and contributions from over 350 different authors we have finally released OpenSSL 3.0!",
//},
//}

//funcMap := template.FuncMap{
//"lower": strings.ToLower,
//}

//http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
//w.Write([]byte("add new software"))
//})

//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//switch {
//case changelog.MatchString(r.URL.Path):
//tmpl, _ := template.New("changelog").Funcs(funcMap).Parse(change)
//soft := softwares[1]
//tmpl.Execute(w, soft)
//default:
//tmpl, _ := template.New("index").Funcs(funcMap).Parse(index)
//tmpl.Execute(w, softwares)
//}
//})

//log.Fatal(http.ListenAndServe(":8080", nil))
//}

package main

import (
	"github.com/flatcar-linux/flog/app"
	"github.com/flatcar-linux/flog/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
