package syahid

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "backend",
}
var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertKaryawan(t *testing.T) {
	dbname := "penggajian"
	karyawan := Karyawan{
		Nama:    "Ujang Saepudin",
		Status:  "Aktif",
		Jabatan: "Staff Administrasi",
		Gaji:    "RP 4.000.000",
	}
	insertedID := InsertKaryawan(dbname, karyawan)
	if insertedID == nil {
		t.Error("Failed to insert karyawan")
	}
}

func TestInsertHonor(t *testing.T) {
	dbname := "penggajian"
	honor := Honor{
		Nama:    "Asep Saepuloh",
		Status:  "Aktif",
		Jabatan: "Staff",
		Gaji:    "RP 1.500.000",
	}
	insertedID := InsertHonor(dbname, honor)
	if insertedID == nil {
		t.Error("Failed to insert honor")
	}
}

func TestInsertJob(t *testing.T) {
	dbname := "penggajian"
	job := Job{
		Namajob: "Staff Administrasi",
	}
	insertedID := InsertJob(dbname, job)
	if insertedID == nil {
		t.Error("Failed to insert job")
	}
}

func TestInsertTeam(t *testing.T) {
	dbname := "penggajian"
	team := Team{
		Nama:      "Uzumaki Memet",
		Deskripsi: "anjay mabar",
	}
	insertedID := InsertTeam(dbname, team)
	if insertedID == nil {
		t.Error("Failed to insert team")
	}
}

func TestInsertAbout(t *testing.T) {
	dbname := "penggajian"
	about := About{
		Isi_satu: "Ini isi satu",
		Isi_dua:  "Ini isi dua",
		Image:    "image.jpg",
	}
	insertedID := InsertAbout(dbname, about)
	if insertedID == nil {
		t.Error("Failed to insert about")
	}
}

func TestGetDataKaryawan(t *testing.T) {
	stats := "Aktif"
	data := GetDataKaryawan(stats)
	fmt.Println(data)
}
func TestGetDataHonor(t *testing.T) {
	stats := "Aktif"
	data := GetDataHonor(stats)
	fmt.Println(data)
}
func TestGetDataJob(t *testing.T) {
	stats := "Staff Administrasi"
	data := GetDataJob(stats)
	fmt.Println(data)
}
func TestGetDataTeam(t *testing.T) {
	stats := "Uzumaki Memet"
	data := GetDataTeam(stats)
	fmt.Println(data)
}
