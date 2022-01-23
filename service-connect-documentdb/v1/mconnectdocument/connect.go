package mconnectdocument

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://github.com/retgits/lambda-util/blob/master/s3.go

// connectionStringTemplate url conexion documentDB o mongoDB
// https://docs.aws.amazon.com/documentdb/latest/developerguide/connect-to-replica-set.html
const connectionStringTemplate = "mongodb://%s:%s@%s/%s?replicaSet=rs0&readpreference=%s"

// DefaultReadPreference preferencia de lectura en documentDB
// https://docs.aws.amazon.com/documentdb/latest/developerguide/how-it-works.html#durability-consistency-isolation
const DefaultReadPreference = "secondaryPreferred"

// DefaultTimeOut tiempo default para la conexion a documentDB
const DefaultTimeOut = 5 * time.Second

const catFile = "/var/task/bin/rds-combined-ca-bundle.pem"

// ParamConnectionDocumentDB parametros necesarios para la conexion a documentDB
type ParamConnectionDocumentDB struct {
	// tiempo de conexion para la base de datos en segundos
	ConnectTimeOut time.Duration
	// nombre de uri cluster de documentDB
	Cluster string
	// nos indica de donde se leeran las operacion de lectura en la base de datos
	ReadPreference string
	// usuario de la base de datos
	UserName string
	// password de la base de datos
	Password string
	// nombre de la base de datos
	DBName string
}

// DocumentDBAPI contiene todos los metodos para operar con la api de documt db
type DocumentDBAPI interface {
	GetContext() context.Context

	Find(ctx context.Context, nameCollection string, filter interface{}) (*mongo.Cursor, error)
	FindOne(ctx context.Context, nameCollection string, filter interface{}) *mongo.SingleResult
	InsertOne(ctx context.Context, nameCollection string, document interface{}) (*mongo.InsertOneResult, error)
	UpdateOne(
		ctx context.Context,
		nameCollection string,
		filter interface{},
		update interface{},
	) (*mongo.UpdateResult, error)
}

// documentDB struct que implementa la interface DocumentDBAPI
type documentDB struct {
	database *mongo.Database
	context  context.Context
}

// GetContext obtenemos el contexto padre de la conexion del cliente de document db
func (d *documentDB) GetContext() context.Context {
	return d.context
}

// Find busca una colección dado los parametros requeridos
// parametro filter es un objeto bson nativo de mongoDB ejemplo (bson.D{{"foo", "bar"})
func (d *documentDB) Find(ctx context.Context, nameCollection string, filter interface{}) (*mongo.Cursor, error) {
	return d.database.Collection(nameCollection).Find(ctx, filter)
}

// FindOne busca una colección dado los parametros requeridos
// parametro filter es un objeto bson nativo de mongoDB ejemplo (bson.D{{"foo", "bar"})
func (d *documentDB) FindOne(ctx context.Context, nameCollection string, filter interface{}) *mongo.SingleResult {
	return d.database.Collection(nameCollection).FindOne(ctx, filter)
}

// InsertOne inserta en una colección dado los parametros requeridos
// parametro document es un objeto bson nativo de mongoDB  ejemplo (bson.D{{"foo", "bar"})
func (d *documentDB) InsertOne(
	ctx context.Context,
	nameCollection string,
	document interface{},
) (*mongo.InsertOneResult, error) {
	return d.database.Collection(nameCollection).InsertOne(ctx, document)
}

// UpdateOne actualiza en una colección dado los parametros requeridos
// parametro filter es un objeto bson nativo de mongoDB ejemplo (bson.D{{"foo", "bar"})
func (d *documentDB) UpdateOne(
	ctx context.Context,
	nameCollection string,
	filter interface{},
	update interface{},
) (*mongo.UpdateResult, error) {
	return d.database.Collection(nameCollection).UpdateOne(ctx, filter, update)
}

// NewCreateConnectionDocumentDB obtiene un nuevo cliente de documentDB
func NewCreateConnectionDocumentDB(db *ParamConnectionDocumentDB) (DocumentDBAPI, error) {

	pwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	fmt.Println(pwd)

	connectURI := fmt.Sprintf(connectionStringTemplate,
		db.UserName,
		db.Password,
		db.Cluster,
		db.DBName,
		db.ReadPreference,
	)

	tlsConfig, err := getCustomTLSConfig(catFile)

	if err != nil {
		return nil, err
	}

	if _, err := os.Stat("/var/task/bin/rds-combined-ca-bundle.pem"); err == nil {
		fmt.Printf("File exists in normal ..\n")
	} else {
		fmt.Printf("File does not exist noraml .. \n")
	}

	url := "mongodb://gromo:Guillermo920517@docdb-2022-01-09-17-41-27.cluster-cgvrlgagsdcn.us-east-1.docdb.amazonaws.com:27017/prime?replicaSet=rs0&readPreference=secondaryPreferred"
	client, err := mongo.NewClient(options.Client().ApplyURI(url).SetTLSConfig(tlsConfig))

	if err != nil {
		fmt.Println("error in client")
		fmt.Println(err)
		fmt.Println(db)
		fmt.Println("error in client")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.ConnectTimeOut)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		fmt.Println("error in connect")
		fmt.Println(err)
		fmt.Println(db, connectURI)
		fmt.Println("error in connect")
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		if err != nil {
			fmt.Println("error in ping")
			fmt.Println(err)
			fmt.Println(url)
			fmt.Println("error out ping")
			return nil, err
		}
	}

	return &documentDB{
		context:  ctx,
		database: client.Database(db.DBName),
	}, nil
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("failed parsing pem file")
	}

	return tlsConfig, nil
}
