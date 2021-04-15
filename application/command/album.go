package command

type CreateAlbum struct {
	Public       bool
	Name         string
	Introduction string
}

type UpdateAlbumInfo struct {
	AlbumId      int
	Public       bool
	Name         string
	Introduction string
}

type GetKeyAndUrl struct {
	Names []string
}

type UploadSuccess struct {
	AlbumId int
	IsVideo bool
	Url     string
}
