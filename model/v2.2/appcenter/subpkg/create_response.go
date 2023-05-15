package subpkg

type CreateResponse struct {
	// CreativeID 创意ID
	item []Item
}

type Item struct {
	PackageId       int    `json:"package_id"`
	BuildStatus     int    `json:"build_status"`
	ParentPackageId int64  `json:"parent_package_id"`
	ChannelId       string `json:"channel_id"`
}

