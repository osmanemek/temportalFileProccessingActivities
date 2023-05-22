package tactivities

import (
	"context"
	"go.temporal.io/sdk/activity"
	"os"
	"strings"
	"time"
)

func (a *Activities) ProcessFileActivity(ctx context.Context, fileName string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("processFileActivity started.", "FileName", fileName)

	defer func() { _ = os.Remove(fileName) }() // cleanup temp file

	// read downloaded file
	data, err := os.ReadFile(fileName)
	if err != nil {
		logger.Error("processFileActivity failed to read file.", "FileName", fileName, "Error", err)
		return "", err
	}

	// process the file
	transData := transcodeData(ctx, data)
	tmpFile, err := saveToTmpFile(transData)
	if err != nil {
		logger.Error("processFileActivity failed to save tmp file.", "Error", err)
		return "", err
	}

	processedFileName := tmpFile.Name()
	logger.Info("processFileActivity succeed.", "SavedFilePath", processedFileName)
	return processedFileName, nil
}

func transcodeData(ctx context.Context, data []byte) []byte {
	// dummy file processor, just do upper case for the data.
	// in real world case, you would want to avoid load entire file content into memory at once.
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		// Demonstrates that heartbeat accepts progress data.
		// In case of a heartbeat timeout it is included into the error.
		activity.RecordHeartbeat(ctx, i)
	}
	return []byte(strings.ToUpper(string(data)))
}
