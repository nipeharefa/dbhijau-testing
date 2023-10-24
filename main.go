package demo

import (
	"context"

	"github.com/arangodb/go-driver"
	drviverHttp "github.com/arangodb/go-driver/http"
)

// ExampleNewClient shows how to create the simple client with a single endpoint
// By default, the client will use the http.DefaultTransport configuration
func ExampleNewClient() (string, error) {
	// Create an HTTP connection to the database
	conn, err := drviverHttp.NewConnection(drviverHttp.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		// log.Fatalf("Failed to create HTTP connection: %v", err)
		return "", err
	}
	// Create a client
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		return "", err
	}
	// Ask the version of the server
	versionInfo, err := c.Version(context.TODO())
	if err != nil {
		// log.Fatalf("Failed to get version info: %v", err)
		return "", err
	}
	// fmt.Printf("Database has version '%s' and license '%s'\n", versionInfo.Version, versionInfo.License)
	return versionInfo.Server, nil
}
