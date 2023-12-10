package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitializeMongoDB() (*mongo.Client, error) {
	logger = GetLogger("mongodb")

	// Configurar as opções de conexão
	uri := "mongodb+srv://twenvweb:admin@cluster0.amyx2ob.mongodb.net/?retryWrites=true&w=majority"

	// Configurar o contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri) // Substitua pela sua URL de conexão

	// Conectar ao servidor MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// Verificar se a conexão está ativa
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Info("Conectado ao MongoDB!")
	return client, nil
}
