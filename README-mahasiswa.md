# PinjamKampus API

- Nama dan NRP: Ararya Arka Anugraha | 5053241050
- Persona: A - Petugas gedung
- Cara menjalankan: `go run .` lalu buka http://localhost:8080

## Fitur
- [*] B1 Sumber daya
- [*] B2 Buat peminjaman
- [*] B3 Lihat peminjaman (daftar, filter, detail)
- [*] B4 Validasi dan error JSON
- [ ] B5 Deteksi bentrok (409)
- [*] B6 Ubah dan hapus
- [ ] B7 Aman untuk permintaan bersamaan
- [*] Fitur khas persona
- [*] Change request

## Yang belum selesai
endpoint yang pakai POST dan PUT beberapa belum ada validasi untuk input ( request body nya ), jadi request body apa aja bisa, untuk yang GET harusnya sudah aman semua, lalu fitur untuk bentrok begitu juga belum divalidasi, dan permintaan bersamaan juga belum saya belum implementasi mutex untuk struct nya sehingga akan race condition :v
