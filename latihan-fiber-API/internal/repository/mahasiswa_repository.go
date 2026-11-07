package repository

import (
	"gorm.io/gorm"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

type FilterMahasiswa struct {
	Cari       string
	Angkatan   int
	Urut       string
	Halaman    int
	PerHalaman int
}

// Struct untuk menampung hasil query program studi dan jumlah mahasiswa.
type ProgramStudiDenganJumlah struct {
	ID              uint
	Kode            string
	Nama            string
	Jenjang         string
	JumlahMahasiswa int64
}

// Ambil mengambil data mahasiswa dengan filter dan pagination.
func Ambil(filter FilterMahasiswa) ([]model.Mahasiswa, int64, error) {
	var mahasiswa []model.Mahasiswa
	var total int64

	query := config.DB.Model(&model.Mahasiswa{})

	if filter.Cari != "" {
		pola := "%" + filter.Cari + "%"
		query = query.Where(
			"nama LIKE ? OR nim LIKE ?",
			pola,
			pola,
		)
	}

	if filter.Angkatan > 0 {
		query = query.Where("angkatan = ?", filter.Angkatan)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	switch filter.Urut {
	case "nama":
		query = query.Order("nama ASC")
	case "ipk":
		query = query.Order("ipk ASC")
	default:
		query = query.Order("id ASC")
	}

	offset := (filter.Halaman - 1) * filter.PerHalaman

	err := query.
		Preload("ProgramStudi").
		Limit(filter.PerHalaman).
		Offset(offset).
		Find(&mahasiswa).Error

	if err != nil {
		return nil, 0, err
	}

	return mahasiswa, total, nil
}

// AmbilPerId mengambil satu mahasiswa berdasarkan ID.
func AmbilPerId(id uint) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.
		Preload("ProgramStudi").
		First(&mahasiswa, id).Error

	return mahasiswa, err
}

// Simpan menyimpan mahasiswa baru.
func Simpan(mahasiswa *model.Mahasiswa) error {
	return config.DB.Create(mahasiswa).Error
}

// Perbarui memperbarui data mahasiswa.
func Perbarui(mahasiswa *model.Mahasiswa) error {
	return config.DB.Save(mahasiswa).Error
}

// Hapus melakukan soft delete pada mahasiswa.
func Hapus(id uint) error {
	hasil := config.DB.Delete(&model.Mahasiswa{}, id)

	if hasil.Error != nil {
		return hasil.Error
	}

	if hasil.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func Pulihkan(id uint) error {
	var mahasiswa model.Mahasiswa

	err := config.DB.Unscoped().First(&mahasiswa, id).Error
	if err != nil {
		return err
	}

	if !mahasiswa.DeletedAt.Valid {
		return gorm.ErrInvalidData
	}

	return config.DB.Unscoped().
		Model(&mahasiswa).
		Update("deleted_at", nil).Error
}

// AmbilPerNim mengambil mahasiswa berdasarkan NIM.
func AmbilPerNim(nim string) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa

	err := config.DB.
		Preload("ProgramStudi").
		Where("nim = ?", nim).
		First(&mahasiswa).Error

	return mahasiswa, err
}

// AmbilProgramStudiDenganJumlah mengambil program studi
// beserta jumlah mahasiswa yang terdaftar.
func AmbilProgramStudiDenganJumlah() ([]ProgramStudiDenganJumlah, error) {
	var hasil []ProgramStudiDenganJumlah

	err := config.DB.
		Model(&model.ProgramStudi{}).
		Select(`
			program_studis.id,
			program_studis.kode,
			program_studis.nama,
			program_studis.jenjang,
			COUNT(mahasiswas.id) AS jumlah_mahasiswa
		`).
		Joins("LEFT JOIN mahasiswas ON mahasiswas.program_studi_id = program_studis.id AND mahasiswas.deleted_at IS NULL").
		Group("program_studis.id").
		Scan(&hasil).Error

	return hasil, err
}
