package mconnectv2

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connectionStringTemplate url conexion documentDB o mongoDB
// https://docs.aws.amazon.com/documentdb/latest/developerguide/connect-to-replica-set.html
const connectionStringTemplate = "mongodb://%s:%s@%s:%d/%s?replicaSet=rs0&readpreference=%s"

// connectionStringTemplateTLS nueva url si la conexion necesita certificado
const connectionStringTemplateTLS = "mongodb://%s:%s@%s:%d/%s?tls=true&replicaSet=rs0&readpreference=%s"

// DefaultReadPreference preferencia de lectura en documentDB
// https://docs.aws.amazon.com/documentdb/latest/developerguide/how-it-works.html#durability-consistency-isolation
const DefaultReadPreference = "secondaryPreferred"

// DefaultTimeOut tiempo default para la conexion a documentDB
const DefaultTimeOut = 5 * time.Second

const pathCertificatesTLS = "./cert/"

// DefaultPort port default documentDB
const DefaultPort = 27017

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
	// puerto connexion base de datos
	Port int
	// indica si la conexion tendra un certificado tls
	TLS bool
	// no es necesario si la TLS es false
	NameCertificate string
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
	connectURI := getURIConnection(db)
	fmt.Println(connectURI)

	clientOptions := options.Client().ApplyURI(connectURI)

	if db.TLS {
		stl, err := getTLSCertificate(db)

		if err != nil {
			return nil, err
		}
		clientOptions.SetTLSConfig(stl)
	}

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		fmt.Println(333)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.ConnectTimeOut)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return &documentDB{
		context:  ctx,
		database: client.Database(db.DBName),
	}, nil
}

func getURIConnection(db *ParamConnectionDocumentDB) string {
	uri := connectionStringTemplateTLS
	if !db.TLS {
		uri = connectionStringTemplate
	}

	return fmt.Sprintf(uri,
		db.UserName,
		db.Password,
		db.Cluster,
		db.Port,
		db.DBName,
		db.ReadPreference,
	)
}

// getTLSCertificate new instance tls certificate
func getTLSCertificate(db *ParamConnectionDocumentDB) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	pathCertificate := fmt.Sprintf("%s%s", pathCertificatesTLS, db.NameCertificate)
	certs, err := ioutil.ReadFile(pathCertificate)

	if err != nil {
		return nil, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return nil, err
	}

	return tlsConfig, nil
}
