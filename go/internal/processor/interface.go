package processor

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jlaffaye/ftp"
)

type Processor interface {
	Process(inputFilePath string, dateStart, dateStop time.Time) error
	io.Closer
}

func NewHub(FTPLogin, FTPPassword string, dbUrl string) (Processor, error) {
	serverConn, err := connectToFTP(FTPLogin, FTPPassword)
	if err != nil {
		return nil, err
	}

	connPool, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	return &processor{
		ftpClient:        serverConn,
		wg:               &sync.WaitGroup{},
		dbConn:           connPool,
		dbChan:           make(chan *XmlFile),
		filesChannel:     make(chan *ftp.Entry, 2),
		ftpPassword:      FTPPassword,
		ftpLogin:         FTPLogin,
		numOfDBWorkers:   1,
		numOfFileWorkers: 1,
		fileWorkers:      make(map[int]*fileProcessor),
		dbWorkers:        make(map[int]*postgresProcessor),
	}, nil
}

func connectToFTP(login, password string) (*ftp.ServerConn, error) {
	serverConn, err := ftp.Dial("ftp.zakupki.gov.ru:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	err = serverConn.Login(login, password)
	if err != nil {
		return nil, err
	}

	return serverConn, nil
}

var startDateForContracts = time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC)

func ShouldSkipFile(entry *ftp.Entry) bool {
	if entry.Type != ftp.EntryTypeFile {
		return true
	}

	if entry.Time.Before(startDateForContracts) {
		return true
	}

	return false
}
