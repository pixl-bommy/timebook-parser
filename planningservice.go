package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"timebook/internal"
	"timebook/utils"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type PlanningService struct {
	log *log.Logger

	// filename of the timebook file (without extension or path)
	filename string
	cached   *internal.PlanningPeriod
}

// @override
func (p *PlanningService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	p.log = log.New(os.Stdout, "[PlanningService] ", log.LstdFlags)

	// TODO: to avoid errors later, let's make sure data/ folder exists at this point

	///////////////////////////////////////////////////////////////////////////
	// TODO: remove later
	p.filename = "cookies"
	p.cached = internal.NewPlanningPeriod()
	p.cached.AddWorkEntry(internal.DailyTask, "Something to do")
	p.cached.AddWorkEntry(internal.DailyTask, "Something other")
	p.cached.AddWorkEntry(internal.WorkBreak, "")
	p.cached.AddWorkEntry(internal.DailyTask, "Something to do after")
	p.cached.AddWorkEntry(internal.ClosingTime, "Something to do")

	p.LoadFile("cookies")
	// TODO: END remove later
	///////////////////////////////////////////////////////////////////////////

	p.log.Println("ServiceStartup done")
	return nil
}

// @override
func (p *PlanningService) ServiceShutdown() error {
	p.SaveFile()

	p.log.Println("ServiceShutdown done")
	return nil
}

// Try to load a timebook file.
//
// Tries to load a file named {filename}.json, which should be located
// in current data folder.
//
// If no file with expected name is found, or if loading the file fails,
// the newly cached data is empty.
//
// If current cached data is not empty when loading a new one, it is
// saved before loading the new data.
func (p *PlanningService) LoadFile(filename string) {
	// maybe there is data cached before, so let's save it first
	p.SaveFile()

	// make sure service is cleared before any other actions
	p.filename = ""
	p.cached = nil

	// TODO: load and parse file
	json, err := utils.LoadJsonFile[internal.PlanningPeriod](buildWorkbookPath(filename))
	if err != nil {
		p.log.Printf("LoadFile(%s): TODO: something failed: %s", p.filename, err.Error())
		return
	}
	p.filename = filename
	p.cached = json

	// there is no cached data at this point, so loading must have failed
	if p.filename == "" || p.cached == nil {
		p.log.Printf("LoadFile(%s): TODO: failed to load", p.filename)
		return
	}

	// at this point loading succeeded
	p.log.Printf("LoadFile(%s): loading succeeded", p.filename)
}

// Try to save a timebook file.
//
// This will save existing cached data to a file named {p.filename}.json
// to data/ folder. If the file already exists it will be overwritten.
//
// If filename is empty or there is no cached data, nothing will happen.
func (p *PlanningService) SaveFile() {
	// if there is nothing to save, do nothing
	if p.filename == "" || p.cached == nil {
		p.log.Printf("SaveFile(): TODO: Nothing to write out")
		return
	}

	filePath := buildWorkbookPath(p.filename)

	// try to save current cached data
	err := utils.StoreJsonFile(filePath, p.cached)
	if err != nil {
		p.log.Printf("SaveFile(): failed to write file %s: %s", filePath, err.Error())
		return
	}

	p.log.Printf("SaveFile(): stored file %s", filePath)
}

func (p *PlanningService) AddWork() {
	if p.cached == nil {
		p.log.Printf("AddWork(): TODO: there is no cached period to add to")
		return
	}
	p.log.Printf("AddWork(): TODO: not implemented")
}

func (p *PlanningService) AddWorkBreak() {
	if p.cached == nil {
		p.log.Printf("AddWorkBreak(): TODO: there is no cached period to add to")
		return
	}
	p.log.Printf("AddWorkBreak(): TODO: not implemented")
}

func (p *PlanningService) AddWorkEnd() {
	if p.cached == nil {
		p.log.Printf("AddWorkEnd(): TODO: there is no cached period to add to")
		return
	}
	p.log.Printf("AddWorkEnd(): TODO: not implemented")
}

func buildWorkbookPath(filename string) string {
	if filename == "" {
		return ""
	}

	return fmt.Sprintf("./data/%s.json", filename)
}
