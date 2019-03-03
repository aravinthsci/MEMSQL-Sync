package database

import (
	. "datasync/constants"
	. "datasync/lib"

	"golang.org/x/crypto/ssh"

	/*	"io/ioutil"
		"os/user"
		"path/filepath"
		"strings"*/
	"log"
)

type DBFetcher interface {
	Fetch() error
}

type DBInserter interface {
	Clean() error
	Insert() error
}

type DBConnector struct {
	SSHClient        *ssh.Client
	Host             string
	ManagementSystem string
	Name             string
	User             string
	Password         string
	IsContainer      bool
}

func CreateFetcher(dbConf Database, sshConf SSH) (fetcher DBFetcher, err error) {
	// Connect to the host of the data soruce.
	log.Print("user:" + sshConf.User)
	log.Print("host:" + sshConf.Host)
	log.Print("key:" + sshConf.Key)
	config := LoadSrcSSHConf(sshConf.User, sshConf.Key)
	log.Print(config.User)
	srcHostConn, err := ssh.Dial("tcp", sshConf.Host+":"+sshConf.Port, config)
	if err != nil {
		return nil, err
	}

	switch dbConf.ManagementSystem {
	case "memsql":
		return &MySQLFetcher{
			SSHClient:   srcHostConn,
			Host:        dbConf.Host,
			Name:        dbConf.Name,
			User:        dbConf.User,
			Password:    dbConf.Password,
			IsContainer: dbConf.IsContainer,
		}, nil
	default:
		return nil, nil
	}
}

func CreateInserter(dbConf Database, sshConf SSH) (inserter DBInserter, err error) {
	config, err := generateSSHSign(sshConf)
	if err != nil {
		return nil, err
	}
	var dstHostConn *ssh.Client
	if sshConf.Host == "localhost" || sshConf.Host == "127.0.0.1" {
		dstHostConn = nil
	} else {
		log.Print("host:" + sshConf.Host)
		log.Print("port:" + sshConf.Port)
		log.Print("user:" + config.User)
		dstHostConn, err = ssh.Dial("tcp", sshConf.Host+":"+sshConf.Port, config)
		if err != nil {
			return nil, err
		}
	}

	switch dbConf.ManagementSystem {
	case "memsql":
		return &MySQLInserter{
			SSHClient:   dstHostConn,
			Host:        dbConf.Host,
			Name:        dbConf.Name,
			User:        dbConf.User,
			Password:    dbConf.Password,
			IsContainer: dbConf.IsContainer,
		}, nil
	default:
		return nil, nil
	}
}

func generateSSHSign(sshConf SSH) (*ssh.ClientConfig, error) {
	if sshConf.Host == "localhost" || sshConf.Host == "127.0.0.1" {
		return nil, nil
	}
	//usr, _ := user.Current()
	//keypathString := strings.Replace(sshConf.Key, "~", usr.HomeDir, 1)
	//keypath, _ := filepath.Abs(keypathString)
	//var password = ""
	//key, err := ioutil.ReadFile(keypath)
	/*if err != nil {
		return nil, err
	}

	//signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}*/

	config := &ssh.ClientConfig{
		User: sshConf.User,
		Auth: []ssh.AuthMethod{
			ssh.Password("*********"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}

func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	for n := range questions {
		answers[n] = "**********"
	}

	return answers, nil
}
