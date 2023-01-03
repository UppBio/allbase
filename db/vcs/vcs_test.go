package vcs

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
)

// func TestConnectVCSDB(t *testing.T) {
// 	// Connect to the VCS Database
// 	db, err := ConnectVCSDB()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	defer db.Close()
// }

func TestInitialization(t *testing.T) {
	// Create a connection object to the Remote SynBio Database
	db, err := sqlx.Connect("sqlite3", "../../data/dev/retsynth/minimal.db")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	// Create a new VCS Database
	vcs := VCS{}
	vcs.Initialize(db)

	// Check if the VCS Database is initialized correctly by checking if the table records exist
	var initchange Change
	vcs.VCSDB.Get(&initchange, "SELECT * FROM change WHERE id = ?", 0)

	if reflect.DeepEqual(initchange, Change{
		ID:                  0,
		Action:              string(Initialize),
		RowPrimaryKeyColumn: "",
		RowPrimaryKeyValue:  "",
		Field:               "",
		OldData:             "",
		NewData:             "",
		ForkID:              0,
		TableName:           "",
		CreatedAt:           time.Now(),
	}) {
		t.Error("The VCS Database is not initialized correctly")
	}

	var initfork Fork
	vcs.VCSDB.Get(&initfork, "SELECT * FROM fork WHERE id = ?", 0)

	if reflect.DeepEqual(initfork, Fork{
		ID:        0,
		ParentID:  sql.NullInt64{Valid: false},
		Name:      "master",
		CreatedAt: time.Now(),
	}) {
		t.Error("The VCS Database is not initialized correctly")
	}

}

func TestFindExistingRow(t *testing.T) {
	// // Set the OS environment variable for the SynBio database
	// os.Setenv("SYNBIO_DB_PATH", "./synbio.db")

	// // Connect to the SynBio Database
	// db, err := ConnectSynBioDB()
	// if err != nil {
	// 	t.Error(err)
	// }
	// defer db.Close()

	// // Create a row that should exist in the database
	// row := Row{
	// 	TableName:        "reaction",
	// 	PrimaryKeyColumn: "ID",
	// 	PrimaryKeyValue:  "rxn10163_c0",
	// }

	// // Use the function to find the row
	// oldRow := findExistingRow(&row)

	// // Assert if the row exists by checking if the returned row is equal
	// if !reflect.DeepEqual(row, oldRow) {
	// 	t.Error("The row does not exist")
	// }

}

func TestTrackChanges(t *testing.T) {

}

func TestGetChanges(t *testing.T) {

}

func TestReadVCSState(t *testing.T) {

}

func TestWriteVCSState(t *testing.T) {

}

func TestCreateFork(t *testing.T) {

}
