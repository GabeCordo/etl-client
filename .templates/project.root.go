package main

import (
	"github.com/GabeCordo/etl/components/cluster"
	"github.com/GabeCordo/etl/core"
	"<project>/src"
)

// DO NOT TOUCH THIS FILE UNLESS YOU ARE CERTAIN ABOUT WHAT YOU ARE DOING

func main() {
	c := core.NewCore()

	// DEFINED CLUSTERS START

	m := src.Vector{} // A structure implementing the etl.Cluster.Cluster interface
	c.Cluster("vector", m, cluster.Config{Identifier: "vector"})

	c.Run()
}
