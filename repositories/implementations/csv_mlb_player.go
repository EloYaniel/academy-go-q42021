package repositories

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"sync"

	e "github.com/EloYaniel/academy-go-q42021/entities"
)

// CSVMLBPlayerRepository struct implements MLBPlayerRepository interface
type CSVMLBPlayerRepository struct {
	filePath string
}

// NewCSVMLBPlayerRepository function creates a new instance of type CSVMLBPlayerRepository.
func NewCSVMLBPlayerRepository(filePath string) *CSVMLBPlayerRepository {
	return &CSVMLBPlayerRepository{filePath: filePath}
}

// GetMLBPlayers gets all MLB Players from the file.
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
			player, err := parsePlayer(line)

			if err != nil {
				return nil, err
			}

			players = append(players, *player)
		}
	}

	return players, nil
}

// GetMLBPlayerByID get a Player by its ID
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

// GetMLBPlayerDesired gets MLB Players from the file concurrently and filetered by its params.
func (repo *CSVMLBPlayerRepository) GetMLBPlayerDesired(filterType string, totalItems int, itemsPerWorker int) ([]e.MLBPlayer, error) {
	f, err := os.Open(repo.filePath)

	if err != nil {
		return nil, errors.New("error opening the file")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Read()
	m := new(sync.Mutex)
	jobs := make(chan int)
	intermediateChan := make(chan int)
	workersCount := totalItems / itemsPerWorker
	done := make(chan struct{})
	wg := new(sync.WaitGroup)
	totalItemsCount := 1

	go func() {
		for {
			select {
			case ev := <-intermediateChan:
				jobs <- ev
			case <-done:
				close(jobs)
				return
			}
		}
	}()

	wg.Add(workersCount)
	var players []e.MLBPlayer

	for i := 0; i < workersCount; i++ {
		go func(workerID int, wg *sync.WaitGroup) {
			defer wg.Done()
			itemsCount := 0
			for j := range jobs {
				if totalItemsCount > totalItems {
					return
				}

				p, err := job(j, filterType, m, reader)

				if err != nil {
					done <- struct{}{}
				}

				if p != nil {
					players = append(players, *p)
					totalItemsCount++
					itemsCount++
				}

				if itemsCount == itemsPerWorker {
					return
				}
			}
		}(i, wg)
	}

	go func() {
		for i := 1; i <= totalItems; i++ {
			intermediateChan <- i
		}
	}()
	wg.Wait()

	return players, nil
}

func job(jobID int, filter string, m *sync.Mutex, reader *csv.Reader) (*e.MLBPlayer, error) {
	m.Lock()
	data, err := reader.Read()
	m.Unlock()
	if err == io.EOF {
		return nil, errors.New(fmt.Sprint("JOB", jobID, "reached end of file"))
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.New("error casting ID")
	}

	if (math.Remainder(float64(id), 2) == 0) != (filter == "even") {
		m.Lock()
		data, err = reader.Read()
		m.Unlock()
		if err == io.EOF {
			return nil, errors.New(fmt.Sprint("JOB", jobID, "reached end of file"))
		}
	}

	return parsePlayer(data)
}

func parsePlayer(line []string) (*e.MLBPlayer, error) {
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

	return &e.MLBPlayer{
		ID:       id,
		Name:     line[1],
		Team:     line[2],
		Position: line[3],
		Height:   height,
		Weight:   float32(weight),
		Age:      float32(age),
	}, nil
}
