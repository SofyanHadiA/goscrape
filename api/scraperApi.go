package api

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"

	scrape "github.com/SofyanHadiA/goscrape/services/scraper"
	"github.com/SofyanHadiA/linqcore/api"
	"github.com/SofyanHadiA/linqcore/database"
)

type scraperAPI struct {
	db database.IDB
}

func NewScraperAPI(db database.IDB) scraperAPI {
	return scraperAPI{db: db}
}

type ScrapperResultVM struct {
	Name         string `json:"Name"`
	Demographics string `json:"demographics"`
	Organization string `json:"organization"`
	Description  string `json:"description"`
}

// ScraperHandler handle scraper operations
func (ctrl scraperAPI) ScraperHandler(w http.ResponseWriter, r *http.Request) {

	respWriter := api.ApiService(w, r)

	liName := r.FormValue("LiName")

	result := doScrape(liName)

	respWriter.ReturnJson(result)
}

func doScrape(liName string) ScrapperResultVM {
	url := fmt.Sprintf("https://www.linkedin.com/in/%s/", liName)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		panic(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		panic(err)
	}

	name, _ := scrape.Find(root, scrape.ById("name"))
	demographics, _ := scrape.Find(root, scrape.ById("demographics"))

	extraInfo, _ := scrape.Find(root, scrape.ByClass("extra-info"))
	org, _ := scrape.Find(extraInfo, scrape.ByClass("org"))

	description, _ := scrape.Find(root, scrape.ByClass("description"))

	return ScrapperResultVM{
		Name:         scrape.Text(name),
		Demographics: scrape.Text(demographics),
		Organization: scrape.Text(org),
		Description:  scrape.Text(description),
	}
}
