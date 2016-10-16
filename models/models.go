package models

import (
	"time"
)

// GameType -
type GameType struct {
	ID           int64    `datastore:"-"`
	Name         string   `json:"name"`
	Version      string   `json:"version" datastore:",noindex"`
	BotTypes     []int64  `json:"botTypes" datastore:",noindex"`
	ItemTypes    []int64  `json:"itemTypes" datastore:",noindex"`
	TerrainTypes []int64  `json:"terrainTypes" datastore:",noindex"`
	ApiFuncs     []string `json:"apiFuncs" datastore:",noindex"`
	NumTeams     int      `json:"numTeams" datastore:",noindex"`
	CreatorID    int64    `json:"creatorID" datastore:",noindex"`
	Description  string   `json:"description" datastore:",noindex"`
	MapIDs       []int64  `json:"mapIDs" datastore:",noindex"`
}

// User -
type User struct {
	ID         int64   `datastore:"-"`
	ProjectIDs []int64 `json:"projectIDs" datastore:",noindex"`
}

// Project -
type Project struct {
	ID         int64    `datastore:"-"`
	Name       string   `json:"name"`
	Language   string   `json:"language" datastore:",noindex"`
	GameTypeID int64    `json:"gameTypeID" datastore:",noindex"`
	UserID     int64    `json:"userID" datastore:",noindex"`
	FileNames  []string `json:"fileNames" datastore:",noindex"`
	GameIDs    []int64  `json:"gameIDs" datastore:",noindex"`
}

// Game -
type Game struct {
	ID          int64     `datastore:"-"`
	Created     time.Time `json:"created" datastore:",noindex"`
	Finished    time.Time `json:"finished" datastore:",noindex"`
	GameTypeID  int64     `json:"gameTypeID" datastore:",noindex"`
	MapID       int64     `json:"mapID" datastore:",noindex"`
	ProjectIDs  []int64   `json:"projectIDs" datastore:",noindex"`
	WinningTeam int64     `json:"winningTeam" datastore:",noindex"`
	Reason      string    `json:"reason" datastore:",noindex"`
	Complete    bool      `json:"complete" datastore:",noindex"`
}

// Map -
type Map struct {
	ID         int64  `datastore:"-"`
	Name       string `json:"name"`
	GameTypeID int64  `json:"gameTypeID" datastore:",noindex"`
	RoundLimit int    `json:"roundLimit" datastore:",noindex"`
}
