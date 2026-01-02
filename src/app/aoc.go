package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var INPUT_FILES = []string{
	"example",
	"input",
}

type AOCApplet interface {
	Part1(string) error
	Part2(string) error
}

type AOCApp struct {
	WorkDirectory string
	LogFile       *os.File
	Log           *log.Logger
	LogPath       string
	Applets       []AOCApplet
	AppName       string
	Part          string
	UseAllInputs  bool
	InputFile     string
}

func NewApp() AOCApp {
	wd, err := os.Getwd()
	handleError(err)
	runner := AOCApp{
		WorkDirectory: wd,
		LogFile:       nil,
		Log:           nil,
		LogPath:       "",
		Applets:       []AOCApplet{},
		AppName:       fmt.Sprintf("day_%s", os.Args[1]),
		Part:          os.Args[2],
		UseAllInputs:  len(os.Args) == 4,
		InputFile:     os.Args[len(os.Args)-1],
	}
	logPath := fmt.Sprintf("%s/logs/%s_%s.log", runner.WorkDirectory, runner.AppName, runner.Part)
	if _, err := os.Stat(logPath); err == nil {
		handleError(os.Remove(logPath))
	}
	logFile, err := os.Create(logPath)
	handleError(err)

	runner.LogPath = logPath
	runner.LogFile = logFile
	runner.Log = log.New(logFile, "", log.Lmsgprefix)

	runner.Applets = append(runner.Applets, NewDay1(runner.Log))
	runner.Applets = append(runner.Applets, NewDay2(runner.Log))
	runner.Applets = append(runner.Applets, NewDay3(runner.Log))
	runner.Applets = append(runner.Applets, NewDay4(runner.Log))
	runner.Applets = append(runner.Applets, NewDay5(runner.Log))
	runner.Applets = append(runner.Applets, NewDay6(runner.Log))
	runner.Applets = append(runner.Applets, NewDay7(runner.Log))
	runner.Applets = append(runner.Applets, NewDay8(runner.Log))
	runner.Applets = append(runner.Applets, NewDay9(runner.Log))
	runner.Applets = append(runner.Applets, NewDay10(runner.Log))
	runner.Applets = append(runner.Applets, NewDay11(runner.Log))
	runner.Applets = append(runner.Applets, NewDay12(runner.Log))
	return runner
}

func (r AOCApp) getApplet() (AOCApplet, error) {
	appletId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return nil, err
	}
	appletId -= 1
	if appletId > len(r.Applets) {
		return nil, errors.New("no applet found with applet id: " + os.Args[1])
	}
	r.Log.Println("Running app:", appletId + 1)
	return r.Applets[appletId], nil
}

func (r AOCApp) getInputFilePath(inputFile string) (path string, err error) {
	path1 := fmt.Sprintf("%s/inputs/%s_%s", r.WorkDirectory, r.AppName, inputFile)
	if _, err = os.Stat(path1); !errors.Is(err, os.ErrNotExist) {
		return path1, nil
	}
	path2 := fmt.Sprintf("%s/inputs/%s_%s_%s", r.WorkDirectory, r.AppName, inputFile, r.Part)
	if _, err = os.Stat(path2); !errors.Is(err, os.ErrNotExist) {
		return path2, nil
	}
	return "", fmt.Errorf("can't open '%s' or '%s' due to '%s'", path1, path2, err.Error())
}

func (r AOCApp) Run() {
	applet, err := r.getApplet()
	handleError(err)
	for _, inputFile := range INPUT_FILES {
		if r.UseAllInputs && inputFile != r.InputFile {
			continue
		}
		inputFilePath, err := r.getInputFilePath(inputFile)
		handleError(err)
		r.Log.Println("Using input:", inputFilePath)
		inputFile, err := os.Open(inputFilePath)
		handleError(err)
		bs, err := io.ReadAll(inputFile)
		handleError(err)
		content := string(bs)
		if os.Args[2] == "1" {
			r.Log.Println("Running Part 1")
			handleError(applet.Part1(content))
		} else {
			r.Log.Println("Running Part 2")
			handleError(applet.Part2(content))
		}
	}
	fmt.Println("Result in", r.LogPath)
	handleError(r.LogFile.Close())
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
