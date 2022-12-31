package src

import (
	"github.com/GabeCordo/etl/components/channel"
)

//
//	ETL: <cluster>
//	Generated On: <date>
//  Generated By: <first-name> <last-name> <email>
//

type <cluster> struct {
}

func (<cluster-short> <cluster>) ExtractFunc(output channel.OutputChannel) {
	// output some data
	close(output)
}

func (<cluster-short> <cluster>) TransformFunc(input channel.InputChannel, output channel.OutputChannel) {
	for request := range input {
		// do something with the request
	}
	close(output)
}

func (<cluster-short> <cluster>) LoadFunc(input channel.InputChannel) {
	for request := range input {
		// do something with the request
	}
}