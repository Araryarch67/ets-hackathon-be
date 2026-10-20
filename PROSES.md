# PROSES

## 1. Persona dan masalahnya
ada date yang dimana untuk validasi date nya harus diparsing dulu pakai library time

## 2. Rancangan struktur data
```json
type Peminjaman struct {
	ID           int       `json:"id"`
	SumberDayaID int       `json:"sumber_daya_id"`
	Peminjam     string    `json:"peminjam"`
	Tanggal      string    `json:"tanggal"`
	SlotMulai    int       `json:"slot_mulai"`
	SlotSelesai  int       `json:"slot_selesai"`
	Jumlah       int       `json:"jumlah"`
	Status       string    `json:"status"`
}
```

```json
type Ringkasan struct {
	Tanggal          string                `json:"tanggal"`
	JumlahPeminjaman int                   `json:"jumlah_peminjaman"`
	TotalUnitSlot    int                   `json:"total_unit_slot"`
	PerSumberDaya    []RingkasanSumberDaya `json:"per_sumber_daya"`
}

type RingkasanSumberDaya struct {
	SumberDayaID     int    `json:"sumber_daya_id"`
	Nama             string `json:"nama"`
	JumlahPeminjaman int    `json:"jumlah_peminjaman"`
}

type SlotKetersediaan struct {
	Slot     int  `json:"slot"`
	Sisa     int  `json:"sisa"`
	Tersedia bool `json:"tersedia"`
}

type Ketersediaan struct {
	SumberDayaID int                `json:"sumber_daya_id"`
	Tanggal      string             `json:"tanggal"`
	Slot         []SlotKetersediaan `json:"slot"`
}

type SumberDaya struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Jenis string `json:"jenis"`
	Stok  int    `json:"stok"`
}
```

## 3. Satu keputusan desain
saya memilih untuk memisah filenya agar lebih modular dan ga terlalu pusing karena bisa baca dengan maximal di masing masing code nya

## 4. Bug atau kendala (minimal dua)
### Bug 1
- Gejala: sulit untuk membuat perssist data dengan struct dan sempat menggunakan dummy data untuk sementara waktu
- Cara menemukan: ketika pertama mengerjakan ada fitur memory dan pakai dummy data dulu untuk cek endpoint
- Penyebab dan solusi: penyebabnya karena saya kurang explore di perssist local data dengan struct, solusinya saya tanya AI untuk membuat perssist data dengan struct, dengan cara menggangap struct itu perssist sejak awal saya kirain awalnya kalau data di local abis di delete via endpoint data nya ga hilang karena local ternyata kesalahpahaman

### Bug 2
- Gejala: sulit memvalidasi time baik yang di path endpoint maupun time yang biasa
- Cara menemukan: ketika dapat task yang ada timenya
- Penyebab dan solusi: penyebab kurang tau soal library time, solusinya adalah saya tanya AI tentang time dan validasi time di golang

## 5. Change request
saya membuat struct baru untuk ringkasan di /store/peminjaman.go untuk ringkasan

```json
type Ringkasan struct {
	Tanggal          string                `json:"tanggal"`
	JumlahPeminjaman int                   `json:"jumlah_peminjaman"`
	TotalUnitSlot    int                   `json:"total_unit_slot"`
	PerSumberDaya    []RingkasanSumberDaya `json:"per_sumber_daya"`
}

type RingkasanSumberDaya struct {
	SumberDayaID     int    `json:"sumber_daya_id"`
	Nama             string `json:"nama"`
	JumlahPeminjaman int    `json:"jumlah_peminjaman"`
}
```

lalu saya buat `GetRingkasan` juga agar lebih enak tinggal di call di controllers yang saya buat
