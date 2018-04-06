package mongodb

import (
	"context"
	"fmt"
	"strings"

	 "fmt"
    // Import the Radix.v2 redis package.
    "github.com/mediocregopher/radix.v2/redis"
    "log"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-redis")

const (
	methodGet     = "GET"
	
	ivConnectionURI = "uri"
	
	ovOutput = "output"
	

}

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

/*
Integration with MongoDb
inputs: {uri}
outputs: {output}
*/
type RedisDbActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &RedisDbActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *RedisDbActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - MongoDb integration
func (a *RedisDbActivity) Eval(ctx activity.Context) (done bool, err error) {

	

    fmt.Println("Electric Ladyland added!")
}