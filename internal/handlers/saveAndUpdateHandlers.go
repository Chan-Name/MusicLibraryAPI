package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"library/internal/storage"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Library struct {
	Storage *storage.Storage
}

type SongToSave struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// @Summary Save a new song
// @Description Save a new song by group and title
// @Accept json
// @Produce json
// @Param song body SongToSave true "Song to save"
// @Success 200 {object} gin.H {"info": "song saved"}
// @Failure 400 {object} gin.H {"error": "invalid request"}
// @Failure 404 {object} gin.H {"error": "song not found"}
// @Failure 500 {object} gin.H {"error": "error saving song"}
// @Router /songs [post]

func (s *Library) SaveSong(c *gin.Context) {
	var req SongToSave

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		slog.Error("ERROR", slog.Any("err", err))
		return
	}

	songDetail, err := decodeApiRequest(c, req.Group, req.Song)
	if err != nil {
		slog.Error("ERROR", slog.Any("err", err))
		return
	}

	err = s.Storage.SaveSongToDb(req.Group, req.Song, songDetail.Text,
		songDetail.Link, songDetail.ReleaseDate)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
			slog.Error("ERROR", slog.Any("err", err))
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		slog.Error("ERROR", slog.Any("err", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"info": "song saved"})
}

// @Summary Update a song
// @Description Update the song details by ID
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param group path string true "Song group"
// @Param song path string true "Song title"
// @Param songDetail body SongDetail true "Updated song details"
// @Success 200 {object} gin.H {"info": "song updated"}
// @Failure 400 {object} gin.H {"error": "invalid request"}
// @Failure 404 {object} gin.H {"error": "song not found"}
// @Failure 500 {object} gin.H {"error": "error updating song"}
// @Router /songs/{group}/{song}/{id} [put]

func (s *Library) UpdateSong(c *gin.Context) {
	id := c.Param("id")
	group := c.Param("group")
	song := c.Param("song")

	songDetail, err := decodeApiRequest(c, group, song)
	if err != nil {
		slog.Error("ERROR", slog.Any("err", err))
		return
	}

	if err := s.Storage.UpdateSongFromIdToDb(id, group, song, songDetail.Text,
		songDetail.Link, songDetail.ReleaseDate); err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
			slog.Error("ERROR", slog.Any("err", err))
			return

		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error update song"})
		slog.Error("ERROR", slog.Any("err", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "song updated"})
}

func decodeApiRequest(c *gin.Context, group, song string) (*SongDetail, error) {
	url := fmt.Sprintf("http://localhost:8080/info/%s/%s", group, song)

	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to call external API"})
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		c.JSON(response.StatusCode, gin.H{"error": "failed to get song details"})
		return nil, err
	}

	var songDetail SongDetail

	if err := json.NewDecoder(response.Body).Decode(&songDetail); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode response"})
		return nil, err
	}
	return &songDetail, nil
}
