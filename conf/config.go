package conf

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	filePerm = 0644
	dirPerm  = 0755
)

var Conf *config

type config struct {
	Dev bool `yaml:"dev"`
	Listen string `yaml:"listen"`
	DB     struct {
		MongoUri string `yaml:"mongoUri"`
		DbName   string `yaml:"dbName"`
	} `yaml:"db"`
}

func init() {
	var err error

	Conf = new(config)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	err = setupLogDir()
	if err != nil {
		logrus.Fatalln(err)
	}

	err = setupLogOutput()
	if err != nil {
		logrus.Fatalln(err)
	}

	yamlFileBytes, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		logrus.Fatalln(err)
	}

	err = yaml.Unmarshal(yamlFileBytes, Conf)
	if err != nil {
		logrus.Fatalln(err)
	}

	err = setupGinLog()
	if err != nil {
		logrus.Fatalln(err)
	}
}

func setupLogDir() error {
	var err error
	if _, err = os.Stat("./logs/"); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir("./logs/", dirPerm)
	}
	return err
}

func setupLogOutput() error {
	var err error

	logFileName := time.Now().Format("2006-01-02")
	logFile, err := os.OpenFile("./logs/"+logFileName+"-app-all.log", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, filePerm)
	if err != nil {
		return err
	}
	logOut := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(logOut)

	return err
}

func setupGinLog() error {
	var err error
	logErrorFileName := time.Now().Format("2006-01-02")
	logErrorFile, err := os.OpenFile("./logs/"+logErrorFileName+"-gin-error.log", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, filePerm)
	if err != nil {
		return err
	}
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, logErrorFile)
	logInfoFileName := time.Now().Format("2006-01-02")
	logInfoFile, err := os.OpenFile("./logs/"+logInfoFileName+"-gin-info.log", syscall.O_CREAT|syscall.O_RDWR|syscall.O_APPEND, filePerm)
	if err != nil {
		return err
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logInfoFile)
	return err
}
