package cmd

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	startupDelay                int32  = 0
	readyTime                   int32  = 0
	endTime                     int32  = 0
	startupMarkerFilePath       string = "/startup-marker"
	readinessMarkerFilePath     string = "/readiness-marker"
	wg                          sync.WaitGroup
	markerRootPath              string = "/tmp"
	startupMarkerFileRootPath   string
	readinessMarkerFileRootPath string

	rootCmd = &cobra.Command{
		Use:   "healthprobe-for-app",
		Short: "k8s health probe study task",
		Long: `I'm learning how to write a health probe for k8s application.
`,
		Run: run,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int32VarP(&startupDelay, "startup", "s", 0, "Startup delay for application")
	rootCmd.PersistentFlags().Int32VarP(&readyTime, "ready", "r", 0, "Ready time for application")
	rootCmd.PersistentFlags().Int32VarP(&endTime, "end", "e", -1, "End time for application")
	rootCmd.PersistentFlags().StringVarP(&markerRootPath, "path", "p", "/tmp", "Marker file path")
}

func run(cmd *cobra.Command, args []string) {
	log.Printf("startupDelay: %d, readyTime: %d, endTime: %d, path: %s\n", startupDelay, readyTime, endTime, markerRootPath)
	startupMarkerFileRootPath = markerRootPath + startupMarkerFilePath
	readinessMarkerFileRootPath = markerRootPath + readinessMarkerFilePath
	deleteFile(startupMarkerFileRootPath) // JIC
	deleteFile(readinessMarkerFileRootPath) // JIC
	wg.Add(3)
	go processStartupDelay()
	go processApplicationReadiness()
	go processApplicationEnd()
	wg.Wait()
}

func processStartupDelay() {
	defer wg.Done()
	if startupDelay <= 0 {
		return
	}

	log.Printf("Startup is delayed for %ds\n", startupDelay)
	time.Sleep(time.Duration(startupDelay) * time.Second)

	createMarkerFile(startupMarkerFileRootPath)
	log.Printf("Startup marker file %s is successfully created\n", startupMarkerFileRootPath)
}

func processApplicationReadiness() {
	defer wg.Done()
	if readyTime <= 0 {
		return
	}

	log.Printf("Application will be ready in %ds\n", readyTime)
	time.Sleep(time.Duration(readyTime) * time.Second)
	createMarkerFile(readinessMarkerFileRootPath)
	log.Printf("Readiness marker file %s is successfully created\n", readinessMarkerFileRootPath)
}

func processApplicationEnd() {
	defer wg.Done()
	if endTime <= 0 {
		endTime = startupDelay + readyTime + 1
	}

	log.Printf("Application will be terminated in %ds\n", endTime)
	time.Sleep(time.Duration(endTime) * time.Second)
	deleteFile(startupMarkerFilePath)
	deleteFile(readinessMarkerFilePath)
	log.Println("Application is terminated")
}

func createMarkerFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}

func deleteFile(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			return err
		}

		log.Printf("File %s is successfully deleted\n", filePath)
	} else if os.IsNotExist(err) {
		// log.Printf("File %s does not exist\n", filePath)
	} else {
		return err
	}

	return nil
}
