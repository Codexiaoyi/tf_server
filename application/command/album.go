package command

type CreateAlbum struct {
	Public       bool
	Name         string
	Introduction string
	Cover        string
}

type UpdateAlbumInfo struct {
	AlbumId      int
	Public       bool
	Name         string
	Introduction string
	Cover        string
}
