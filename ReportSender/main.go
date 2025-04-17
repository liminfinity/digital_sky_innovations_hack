package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/kurochkinivan/ReportSender/truncate"
	"github.com/kurochkinivan/ReportSender/upload"
	"github.com/pkg/errors"
)

func main() {
	log.Println("starting program...")

	flags, err := ParseFlags()
	if err != nil {
		log.Fatalln("failed to parse flags:", err)
	}

	log.Println("loading xml template...")
	err = upload.LoadXMLTemplate(filepath.Join(flags.PathToFilesDirectory, "pid.xml"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("processing files...")
	err = upload.ProcessFiles(flags.PathToFilesDirectory)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("uploading files to github...")
	err = upload.UploadFilesInSingleCommit(flags.PathToFilesDirectory, flags.Owner, flags.Repo, flags.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connecting to database...")
	db, err := truncate.ConnectToDB(flags.PathToDB)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("truncating pids table...")
	result, err := truncate.TruncatePIDs(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

type Flags struct {
	PathToFilesDirectory string
	PathToDB             string
	DBName               string
	GithubData
}

type GithubData struct {
	Owner       string
	Repo        string
	AccessToken string
}

func ParseFlags() (*Flags, error) {
	flags := new(Flags)

	flag.StringVar(&flags.PathToFilesDirectory, "files", "", "path to directory that contains files")
	flag.StringVar(&flags.PathToDB, "db", "", "path to sqlite database")
	flag.StringVar(&flags.DBName, "name", "app.db", "name of database")
	flag.StringVar(&flags.Owner, "owner", "", "owner of the github repository which is used to store reports")
	flag.StringVar(&flags.Repo, "repo", "", "repository where to store reports")
	flag.StringVar(&flags.AccessToken, "token", "", "github access token")
	flag.Parse()

	if flags.PathToFilesDirectory == "" {
		return nil, errors.New("Path to files directory is empty")
	}

	if flags.PathToDB == "" {
		return nil, errors.New("Path to database is empty")
	}

	if flags.Owner == "" {
		return nil, errors.New("Owner is empty")
	}

	if flags.Repo == "" {
		return nil, errors.New("Repo is empty")
	}

	if flags.AccessToken == "" {
		return nil, errors.New("Access token is empty")
	}

	return flags, nil
}
