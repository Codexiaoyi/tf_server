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
	Name string
}

type UploadSuccess struct {
	AlbumId int
	IsVideo bool
	Url     string
}

type SetUserAlbumCover struct {
	AlbumId int
	MediaId int
}
