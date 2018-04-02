// database migration
// all migration modules must be declared here
package migration

func Run() {
	users.migrate()
}

func Rollback() {

}
