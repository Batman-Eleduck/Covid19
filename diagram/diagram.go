package diagram

import (
	"fmt"
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/aws"
)

func GetDiagram() *diagram.Diagram {
	workerCount := 5

	d, err := diagram.New(
		diagram.Label("Workers"),
		diagram.Filename("workers"),
		diagram.Direction("TB"),
	)

	if err != nil {
		log.Fatal(err)
	}

	lb := aws.Network.ElasticLoadBalancing(diagram.NodeLabel("nlb"))
	d.Add(lb)

	db := aws.Database.Database(diagram.NodeLabel("db"))
	d.Add(db)

	workers := make([]*diagram.Node, workerCount)

	for i := 0; i < workerCount; i++ {
		label := fmt.Sprintf("worker %d", i+1)
		workers[i] = aws.Compute.Ec2(diagram.NodeLabel(label))
	}

	d.Group(diagram.NewGroup("workers").
		Add(workers...).
		ConnectAllTo(db.ID()).
		ConnectAllFrom(lb.ID()),
	)
	return d
}
