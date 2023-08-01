package services

import (
	"encoding/json"
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/SenselessA/w2w_backend/internal/repository"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type List struct {
	NextPage string         `json:"next_page"`
	PrevPage string         `json:"prev_page"`
	Results  []models.Movie `json:"results"`
	Time     string         `json:"time"`
	Total    uint32         `json:"total"`
}

type ServiceKodik struct {
	repo   *repository.RepoMovies
	client *http.Client
}

func initKodik(repo *repository.RepoMovies) *ServiceKodik {
	return &ServiceKodik{repo: repo, client: http.DefaultClient}
}

func (s *ServiceKodik) StartParseMovies(apikey string) error {
	uri, err := url.Parse("https://kodikapi.com/list")
	if err != nil {
		log.Fatal(err)
	}

	params := uri.Query()
	params.Set("token", apikey)
	params.Set("limit", "100")
	params.Set("with_material_data", "true")
	params.Set("types", "anime-serial")
	params.Set("sort", "updated_at")

	channel := make(chan []models.Movie)

	uri.RawQuery = params.Encode()
	targetUrl := uri.String()

	go s.GetAllList(&targetUrl, &channel)

	go func() {
		for movies := range channel {
			err := s.repo.BatchInsertAndUpdate(&movies)
			if err != nil {
				log.Fatal("BatchInsertAndUpdate:", err)
			}
		}
	}()

	return nil
}

func (s *ServiceKodik) GetAllList(targetUrlCur *string, channelCur *chan []models.Movie) {
	targetUrl := *targetUrlCur
	channel := *channelCur

	lastUpdated, err := s.repo.GetLastUpdatedAt()
	if err != nil {
		log.Println("lastUpdated error: ", err)
	}

	if lastUpdated != nil {
		lastUpdatedAt := *lastUpdated
		lastUpdatedAtTime, err := time.Parse(time.RFC3339, lastUpdatedAt)
		log.Println("lastUpdatedAt: ", lastUpdatedAt)
		if err != nil {
			log.Println("lastUpdatedAtTime error: ", err)
		}

		for {
			// parse until we have next_page in response
			if targetUrl != "" {
				r, err := s.MakeRequest(targetUrl)
				if err != nil {
					log.Fatal(err)
				}

				log.Println("parsing...")
				targetUrl = r.NextPage

				// вставляем ток до lastUpdatedAt
				for _, movie := range r.Results {
					movieUpdatedAt, err := time.Parse(time.RFC3339, movie.UpdatedAt)
					if err != nil {
						log.Println("movieUpdatedAt error: ", err)
					}
					if lastUpdatedAtTime.Before(movieUpdatedAt) || lastUpdatedAtTime.Equal(movieUpdatedAt) {
					} else {
						targetUrl = ""
						break
					}
				}

				channel <- r.Results
			} else {
				log.Println("parsing done")
				break
			}
		}
	} else {
		// а ниже фул парсинг, если lastUpdatedAt нет (база movies пустая)
		for {
			// parse until we have next_page in response
			if targetUrl != "" {
				r, err := s.MakeRequest(targetUrl)
				if err != nil {
					log.Fatal("MakeRequest: ", err)
				}

				log.Println("parsing...")

				channel <- r.Results

				targetUrl = r.NextPage
			} else {
				break
			}
		}
	}

	close(channel)
}

func (s *ServiceKodik) MakeRequest(targetUrl string) (List, error) {
	var r List

	resp, err := s.client.Get(targetUrl)
	if err != nil {
		return r, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return r, err
	}

	return r, nil
}
