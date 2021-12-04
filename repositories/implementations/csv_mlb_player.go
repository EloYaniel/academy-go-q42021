package repositories

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	e "github.com/EloYaniel/academy-go-q42021/entities"
)

// CSVMLBPlayerRepository implements MLBPlayerRepository interface
type CSVMLBPlayerRepository struct {
	filePath string
}

func NewCSVMLBPlayerRepository(filePath string) *CSVMLBPlayerRepository {
	return &CSVMLBPlayerRepository{filePath: filePath}
}

func (repo *CSVMLBPlayerRepository) GetMLBPlayers() ([]e.MLBPlayer, error) {
	f, err := os.Open(repo.filePath)

	if err != nil {
		return nil, errors.New("error opening the file")
	}
	defer f.Close()
	data, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return nil, errors.New("error reading the file")
	}
	var players []e.MLBPlayer
	for i, line := range data {
		if i != 0 {
			id, err := strconv.Atoi(line[0])

			if err != nil {
				return nil, errors.New("error casting ID")
			}
			height, err := strconv.Atoi(line[4])

			if err != nil {
				return nil, errors.New("error casting Height")
			}
			weight, err := strconv.ParseFloat(line[5], 32)
			if err != nil {
				return nil, errors.New("error casting Weight")
			}
			age, err := strconv.ParseFloat(line[6], 32)

			if err != nil {
				return nil, errors.New("error casting Age")
			}

			players = append(players, e.MLBPlayer{
				ID:       id,
				Name:     line[1],
				Team:     line[2],
				Position: line[3],
				Height:   height,
				Weight:   float32(weight),
				Age:      float32(age),
			})
		}
	}

	return players, nil
}

func (repo *CSVMLBPlayerRepository) GetMLBPlayerByID(id int) (*e.MLBPlayer, error) {
	players, err := repo.GetMLBPlayers()

	if err != nil {
		return nil, errors.New("error getting player")
	}

	for _, p := range players {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, nil
}

func (repo *CSVMLBPlayerRepository) GetMLBPlayerDesired(filterType string, totalItems int, itemsPerWorker int) ([]e.MLBPlayer, error) {
	f, err := os.Open(repo.filePath)

	if err != nil {
		return nil, errors.New("error opening the file")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Read()

	pCount := 0
	m := sync.Mutex{}
	job := func(jobID int) (*e.MLBPlayer, bool) {
		pCount++
		m.Lock()
		data, err := reader.Read()
		m.Unlock()
		if err == io.EOF {
			log.Println("JOB", jobID, "reached end of file")

			return nil, true
		}

		if err != nil {
			log.Println("JOB", jobID, "error reading line: ", err)

			return nil, false
		}
		id, err := strconv.Atoi(data[0])

		if err != nil {
			return nil, false
		}
		height, err := strconv.Atoi(data[4])

		if err != nil {
			return nil, false
		}
		weight, err := strconv.ParseFloat(data[5], 32)
		if err != nil {
			return nil, false
		}
		age, err := strconv.ParseFloat(data[6], 32)

		if err != nil {
			return nil, false
		}

		return &e.MLBPlayer{
			ID:       id,
			Name:     data[1],
			Team:     data[2],
			Position: data[3],
			Height:   height,
			Weight:   float32(weight),
			Age:      float32(age),
		}, false
	}

	jobs := make(chan int, totalItems)
	results := make(chan *e.MLBPlayer, totalItems)
	workersCount := totalItems / itemsPerWorker

	log.Println(workersCount)
	for i := 0; i < workersCount; i++ {
		go func(wokerID int, jobs <-chan int, results chan<- *e.MLBPlayer) {
			for j := range jobs {
				log.Println("Worker", wokerID, "executed job", j)
				p, f := job(j)
				if f {
					close(results)
					break
				}
				results <- p
				if pCount == totalItems {
					close(results)
					break
				}
			}
		}(i, jobs, results)
	}

	for i := 1; i <= totalItems; i++ {
		jobs <- i
	}
	close(jobs)

	var players []e.MLBPlayer
	for r := range results {
		players = append(players, *r)
	}

	return players, nil
}
