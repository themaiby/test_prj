package migration

import (
	"github.com/fatih/color"
	log "github.com/kataras/golog"
)

var table = "users"

var users usersTable

type usersTable struct{}

func (this *usersTable) migrate() {

	/*yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()*/
	green := color.New(color.FgGreen).SprintFunc()

	log.Warn("==== RUNNING MIGRATIONS ====")
	log.Println("Table [" + green(table) + "] created")
}
