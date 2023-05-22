package tactivities

import (
	"context"
	"go.temporal.io/sdk/activity"
	"os"
)

func (a *Activities) DownloadFileActivity(ctx context.Context, fileID string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Downloading file...", "FileID", fileID)
	data := a.BlobStore.downloadFile(fileID)

	tmpFile, err := saveToTmpFile(data)
	if err != nil {
		logger.Error("downloadFileActivity failed to save tmp file.", "Error", err)
		return "", err
	}
	fileName := tmpFile.Name()
	logger.Info("downloadFileActivity succeed.", "SavedFilePath", fileName)
	return fileName, nil
}

func (b *BlobStore) downloadFile(fileID string) []byte {
	// dummy downloader
	dummyContent := "dummy content for fileID:" + fileID
	return []byte(dummyContent)
}

func saveToTmpFile(data []byte) (f *os.File, err error) {
	tmpFile, err := os.CreateTemp("", "temporal_sample")
	if err != nil {
		return nil, err
	}
	_, err = tmpFile.Write(data)
	if err != nil {
		_ = os.Remove(tmpFile.Name())
		return nil, err
	}

	return tmpFile, nil
}
