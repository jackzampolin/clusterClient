package main

import (
	"flag"
	// "fmt"
	"log"

	"github.com/jackzampolin/clusterClient/cmd"
	"github.com/kubernetes/client-go/1.5/kubernetes"
	"google.golang.org/api/container/v1"
)

var (
	projectID *string
	zone      *string
	update    *string
)

func init() {
	projectID = flag.String("projectID", "jackzampolin-web", "project id for google compute cloud")
	zone = flag.String("zone", "us-east1-b", "zone for google compute cloud")
	update = flag.String("update", "foobar-1", "cluster to update")
	flag.Parse()
}

func main() {
	cl, err := containerClient.NewClusterClient(*projectID, *zone)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cl.ProjectsSvc.Zones.Clusters.List(*projectID, *zone).Do()
	if err != nil {
		panic(err)
	}
	var workingCluster *container.Cluster
	for _, cl := range resp.Clusters {
		if cl.Name == *update {
			workingCluster = cl
		}
	}

}
