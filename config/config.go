package config

import (
	"encoding/json"
	"io"
	"os"
	"peertubeupload/logger"
)

type Config struct {
	APIConfig struct {
		URL             string `json:"url"`
		Port            string `json:"port"`
		Username        string `json:"username"`
		Password        string `json:"password"`
		ChannelID       int    `json:"channelId"`
		DownloadEnabled bool   `json:"downloadEnabled"`
		CommentsPolicy  int    `json:"commentsPolicy"`
		Privacy         int    `json:"privacy"`
		Category        int    `json:"category"`
		Language	string `json:"language"`
		WaitTranscoding bool   `json:"waitTranscoding"`
	} `json:"apiConfig"`
	LoadType struct {
		LoadPathFromDB     bool     `json:"loadPathFromDB"`
		LoadFromFolder     bool     `json:"loadFromFolder"`
		SpecificExtensions bool     `json:"specificextensions"`
		Extensions         []string `json:"extensions"`
		ConvertAudioToMp3  bool     `json:"convertAudioToMp3"`
		TempFolder         string   `json:"tempFolder"`
		LogType            string   `json:"logType"`
	} `json:"loadType"`
	FolderConfig struct {
		Path string `json:"path"`
	} `json:"folderConfig"`
	DBConfig struct {
		DBType           string   `json:"dbType"`
		Username         string   `json:"username"`
		Password         string   `json:"password"`
		Port             string   `json:"port"`
		Host             string   `json:"host"`
		Dbname           string   `json:"dbname"`
		MediaIdentifier  []string `json:"media_identifier"`
		TableName        string   `json:"table_name"`
		Title            string   `json:"title"`
		Description      string   `json:"description"`
		FilePath         string   `json:"file_path"`
		ReferenceColumns []string `json:"reference_columns"`
	} `json:"dbConfig"`
	ProccessConfig struct {
		Threads int `json:"threads"`
	}
}

func (c *Config) LoadConfiguration(file string) {
	configFile, err := os.Open(file)

	// If file not found, create a sample config file
	if os.IsNotExist(err) {
		*c = Config{
			APIConfig: struct {
				URL             string `json:"url"`
				Port            string `json:"port"`
				Username        string `json:"username"`
				Password        string `json:"password"`
				ChannelID       int    `json:"channelId"`
				DownloadEnabled bool   `json:"downloadEnabled"`
				CommentsPolicy  int    `json:"commentsPolicy"`
				Privacy         int    `json:"privacy"`
				Category	int    `json:"category"`
				Language	string `json:"language"`
				WaitTranscoding bool   `json:"waitTranscoding"`
			}{
				URL:             "http://peertube.localhost",
				Port:            "9000",
				Username:        "root",
				Password:        "ali12345",
				ChannelID:       1,
				DownloadEnabled: false,
				CommentsPolicy:  2,
				Privacy:         2,
				Category:	 1,
				Language:	 "en",
				WaitTranscoding: true,
			},
			LoadType: struct {
				LoadPathFromDB     bool     `json:"loadPathFromDB"`
				LoadFromFolder     bool     `json:"loadFromFolder"`
				SpecificExtensions bool     `json:"specificextensions"`
				Extensions         []string `json:"extensions"`
				ConvertAudioToMp3  bool     `json:"convertAudioToMp3"`
				TempFolder         string   `json:"tempFolder"`
				LogType            string   `json:"logType"`
			}{
				LoadPathFromDB:     false,
				LoadFromFolder:     true,
				SpecificExtensions: true,
				Extensions:         []string{".mp4", ".wmv"},
				ConvertAudioToMp3:  true,
				TempFolder:         "./tmp/",
				LogType:            "db , file or none",
			},
			DBConfig: struct {
				DBType           string   `json:"dbType"`
				Username         string   `json:"username"`
				Password         string   `json:"password"`
				Port             string   `json:"port"`
				Host             string   `json:"host"`
				Dbname           string   `json:"dbname"`
				MediaIdentifier  []string `json:"media_identifier"`
				TableName        string   `json:"table_name"`
				Title            string   `json:"title"`
				Description      string   `json:"description"`
				FilePath         string   `json:"file_path"`
				ReferenceColumns []string `json:"reference_columns"`
			}{
				DBType:           "postgres or oracle",
				Username:         "user",
				Password:         "password",
				Port:             "5432 or 1521",
				Host:             "localhost",
				Dbname:           "dbname",
				TableName:        "media_table",
				MediaIdentifier:  []string{"id", "sub_id"},
				Title:            "title_column",
				Description:      "description_column",
				FilePath:         "file_path_column",
				ReferenceColumns: []string{"peertube_id", "uuid", "shortuuid", "file_path"},
			},
			FolderConfig: struct {
				Path string `json:"path"`
			}{
				Path: "./videos/",
			},
			ProccessConfig: struct {
				Threads int `json:"threads"`
			}{
				Threads: 1,
			},
		}
		configJSON, _ := json.MarshalIndent(*c, "", " ")
		_ = os.WriteFile(file, configJSON, 0644)
		logger.LogInfo("No config file is there , created one for you .. please re-run the script after modifing the config.json file", nil)
		if _, err := os.Stat(c.LoadType.TempFolder); os.IsNotExist(err) {
			// If the directory does not exist, create it
			errDir := os.MkdirAll(c.LoadType.TempFolder, 0755)
			if errDir != nil {
				panic(err)
			}
			logger.LogInfo("Directory created", nil)

		} else {
			logger.LogInfo("Directory already exists", nil)

		}
		os.Exit(0)
	} else {

		defer configFile.Close()
		byteValue, _ := io.ReadAll(configFile)
		err := json.Unmarshal(byteValue, c)
		if err != nil {
			panic("not able to read config.json ")
		}
		if _, err := os.Stat(c.LoadType.TempFolder); os.IsNotExist(err) {
			// If the directory does not exist, create it
			errDir := os.MkdirAll(c.LoadType.TempFolder, 0755)
			if errDir != nil {
				panic(err)
			}
			logger.LogInfo("TEMP Directory created", nil)

		} else {
			logger.LogInfo("TEMP Directory exists", nil)
		}

	}
}
