package db

import (
	"fmt"
)

var Db DbHandler

type DbObject struct {
	Method          string
	ReplyReason     string
	Ruri            string
	RuriUser        string
	RuriDomain      string
	FromUser        string
	FromDomain      string
	FromTag         string
	ToUser          string
	ToDomain        string
	ToTag           string
	PidUser         string
	ContactUser     string
	AuthUser        string
	CallId          string
	CallIdAleg      string
	Via             string
	ViaBranch       string
	Cseq            string
	Diversion       string
	Reason          string
	ContentType     string
	Auth            string
	UserAgent       string
	SourceIp        string
	SourcePort      uint16
	DestinationIp   string
	DestinationPort uint16
	ContactIp       string
	ContactPort     uint16
	OriginatorIp    string
	OriginatorPort  uint16
	Proto           uint
	Family          uint
	RtpStat         string
	Type            uint
	Node            string
	Msg             string
}

type Filter struct {
	StartDate string
	EndDate   string
	Equals    map[string]string
	Like      map[string]string
}

type Options struct {
	Sort  []string
	Limit uint
}

type DbHandler interface {
	Connect(connectString string) error
	Insert(dbObject *DbObject) error
	Find(filter *Filter, options *Options, result *[]DbObject) error
}

func RegisterHandler(dbHandler DbHandler) {
	Db = dbHandler
	err := Db.Connect("localhost")
	if err != nil {
		fmt.Println(err)
	}
}

func NewDbObject() *DbObject {
	newDbObject := &DbObject{}
	return newDbObject
}

func (d *DbObject) Save() error {
	err := Db.Insert(d)
	if err != nil {
		return err
	}
	return nil
}
