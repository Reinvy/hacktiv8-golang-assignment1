package models

type Person struct {
	ID        uint
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var ListPerson = []Person{
	{
		ID:        1,
		Nama:      "Budi",
		Alamat:    "Jakarta",
		Pekerjaan: "Programmer",
		Alasan:    "Memilih kelas Golang karena ingin memperdalam pemrograman backend",
	},
	{
		ID:        2,
		Nama:      "Citra",
		Alamat:    "Bandung",
		Pekerjaan: "Designer",
		Alasan:    "Bersemangat dalam kreativitas dan tertarik dengan fitur-fitur bahasa Golang",
	},
	{
		ID:        3,
		Nama:      "Dewi",
		Alamat:    "Surabaya",
		Pekerjaan: "Teacher",
		Alasan:    "Senang bekerja dengan anak-anak dan ingin menjelajahi Golang untuk tujuan pendidikan",
	},
	{
		ID:        4,
		Nama:      "Eko",
		Alamat:    "Yogyakarta",
		Pekerjaan: "Writer",
		Alasan:    "Mengekspresikan pemikiran melalui tulisan dan penasaran dengan kesederhanaan Golang",
	},
	{
		ID:        5,
		Nama:      "Fajar",
		Alamat:    "Medan",
		Pekerjaan: "Sales",
		Alasan:    "Menikmati bertemu orang baru dan percaya Golang dapat meningkatkan pengembangan backend",
	},
	{
		ID:        6,
		Nama:      "Gita",
		Alamat:    "Semarang",
		Pekerjaan: "Accountant",
		Alasan:    "Detail-oriented dan ingin menjelajahi kemampuan kinerja Golang",
	},
	{
		ID:        7,
		Nama:      "Hadi",
		Alamat:    "Makassar",
		Pekerjaan: "Engineer",
		Alasan:    "Bersemangat tentang teknologi dan tertarik dengan fitur konkurensi Golang",
	},
	{
		ID:        8,
		Nama:      "Indra",
		Alamat:    "Palembang",
		Pekerjaan: "Marketing",
		Alasan:    "Pemikir kreatif dan ingin menjelajahi Golang untuk membangun API yang efisien",
	},
	{
		ID:        9,
		Nama:      "Joko",
		Alamat:    "Denpasar",
		Pekerjaan: "Entrepreneur",
		Alasan:    "Ingin menciptakan sesuatu yang berdampak dan percaya Golang dapat membantu mencapai itu",
	},
	{
		ID:        10,
		Nama:      "Kartika",
		Alamat:    "Bali",
		Pekerjaan: "Photographer",
		Alasan:    "Bersemangat dalam menangkap momen dan penasaran dengan kerangka kerja web Golang",
	},
}
