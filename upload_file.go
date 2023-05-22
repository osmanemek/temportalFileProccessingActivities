package tactivities

import (
	"context"
	"go.temporal.io/sdk/activity"
	"os"
	"time"
)

func (a *Activities) UploadFileActivity(ctx context.Context, fileName string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("uploadFileActivity begin.", "UploadedFileName", fileName)

	defer func() { _ = os.Remove(fileName) }() // cleanup temp file

	err := a.BlobStore.uploadFile(ctx, fileName)
	if err != nil {
		logger.Error("uploadFileActivity uploading failed.", "Error", err)
		return err
	}
	logger.Info("uploadFileActivity succeed.", "UploadedFileName", fileName)
	return nil
}

func (b *BlobStore) uploadFile(ctx context.Context, filename string) error {
	// dummy uploader
	_, err := os.ReadFile(filename)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		// Demonstrates that heartbeat accepts progress data.
		// In case of a heartbeat timeout it is included into the error.
		activity.RecordHeartbeat(ctx, i)
	}
	if err != nil {
		return err
	}
	return nil
}
