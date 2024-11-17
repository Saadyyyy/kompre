package models

import "gorm.io/gorm"

type KinerjaCrud struct {
	gorm.Model
	Idkary                string `json:"id_kary" gorm:"primaryKey"`
	Nama                  string `json:"nama"`
	Kehadiran             int    `json:"kehadiran"`
	HasilKerja            int    `json:"jumlah_kinerjaCrud"`
	Inisiatif             int    `json:"inisiatif"`
	TeamWork              int    `json:"team_work"`
	IndikatorKinerjaUtama int    `json:"indikator_kinerja_utama"`
	Penilaian             string `json:"penilaian"`
}
