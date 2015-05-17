package cassandracheck

import (
	"fmt"
	"github.com/pivotal-golang/lager"
	"os/exec"
	"strings"
)

type Cassandrachecker struct {
	bool
}

func New(logger lager.Logger) *Cassandrachecker {
	return &Cassandrachecker{}
}

func (c *Cassandrachecker) Check() bool {
	cmd := "nodetool"
	args := []string{"-h 127.0.0.1 status | tail -2 | head -1"}
	if out, err := exec.Command(cmd, args...).CombinedOutput(); err != nil {
		fmt.Println(fmt.Sprint(err))
		return false
	} else if strings.Contains(string(out), "UN") {
		fmt.Println("CASSANDRA RUNNING NORMAL. Nodetool Output: ", string(out))
		return true
	} else {
		fmt.Println("CASSANDRA DOWN. Nodetool Output: ", string(out))
		return false
	}
}
