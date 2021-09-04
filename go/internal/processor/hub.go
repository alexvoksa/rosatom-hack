package processor

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jlaffaye/ftp"
)

const minimalNumOfWorkers = 5

type processor struct {
	numOfFileWorkers, numOfDBWorkers int
	ftpLogin, ftpPassword            string
	ftpClient                        *ftp.ServerConn
	dbConn                           *pgxpool.Pool
	wg                               *sync.WaitGroup
	filesChannel                     chan *ftp.Entry
	fileWorkers                      map[int]*fileProcessor
	dbWorkers                        map[int]*postgresProcessor
	dbChan                           chan *Tender
}

func (p *processor) Process(directoryPath string, dateStart, dateStop time.Time) error {
	err := p.ftpClient.ChangeDir(directoryPath)
	if err != nil {
		return fmt.Errorf("failed to change dir to %s due to %p", directoryPath, err)
	}

	err = p.initWorkers(directoryPath)
	if err != nil {
		return err
	}

	// воркеры запущены, начинаем писать в канал
	filesList, err := p.ftpClient.List(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range filesList {
		if ShouldSkipFile(entry) {
			log.Printf("skipping something not a file: %s, type: %d\n, time: %v", entry.Name, entry.Type, entry.Time)
			continue
		}

		fmt.Println("sending file: ", entry.Name, entry.Time)
		p.filesChannel <- entry
	}

	p.wg.Wait()

	return nil
}

func (p *processor) initWorkers(directoryPath string) error {
	for i := 0; i < p.numOfFileWorkers; i++ {
		serverConn, err := connectToFTP(p.ftpLogin, p.ftpPassword)
		if err != nil {
			return err
		}

		err = serverConn.ChangeDir(directoryPath)
		if err != nil {
			return err
		}

		p.fileWorkers[i] = &fileProcessor{
			workerID:  i,
			inputChan: p.filesChannel,
			client:    serverConn,
			wg:        p.wg,
			dbChannel: p.dbChan,
		}
	}

	for i := 0; i < p.numOfDBWorkers; i++ {
		conn, err := p.dbConn.Acquire(context.Background())
		if err != nil {
			return err
		}

		p.dbWorkers[i] = &postgresProcessor{
			// чтобы удобнее отличать бд от фтп
			workerID:   i + 100,
			postgresDB: conn,
			wg:         p.wg,
			inputChan:  p.dbChan,
		}
	}

	p.wg.Add(p.numOfDBWorkers + p.numOfFileWorkers)

	for i := 0; i < p.numOfFileWorkers; i++ {
		go p.fileWorkers[i].startProcessing()
	}

	for i := 0; i < p.numOfDBWorkers; i++ {
		go p.dbWorkers[i].startProcessing()
	}

	return nil
}

func (p *processor) Close() error {

	close(p.filesChannel)
	close(p.dbChan)
	p.dbConn.Close()

	return nil
}
