package state

var (
	maxDownloadedBlockID int64
)

// GetMaxDownloadedBlockID gets the value of maxDownloadedBlockID
func GetMaxDownloadedBlockID() int64 {
	return maxDownloadedBlockID
}

// SetMaxDownloadedBlockID sets the value of maxDownloadedBlockID
func SetMaxDownloadedBlockID(v int64) {
	maxDownloadedBlockID = v
}
