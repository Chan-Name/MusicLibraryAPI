package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Deletes a song by ID
// @Description Deletes a song based on the ID provided
// @Param id path integer true "Song ID"
// @Success 200 {object} gin.H {"info": "song deleted"}
// @Failure 404 {object} gin.H {"error": "song not found"}
// @Failure 500 {object} gin.H {"error": "error deleting song"}
// @Router /songs/{id} [delete]
func (s *Library) DeleteSong(c *gin.Context) {
	id := c.Param("id")

	err := s.Storage.DeleteSongToDb(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
			slog.Error("ERROR", slog.Any("err", err))
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting song"})
		slog.Error("ERROR", slog.Any("err", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"info": "song deleted"})
}

// @Summary Deletes the entire library
// @Description Delete all songs in the library
// @Success 200 {object} gin.H {"info": "db deleted"}
// @Failure 500 {object} gin.H {"error": "error deleting db"}
// @Router /library [delete]
func (s *Library) DeleteDb(c *gin.Context) {
	err := s.Storage.Del()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting db"})
		slog.Error("ERROR", slog.Any("err", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"info": "db deleted"})
}

// @Summary Retrieves a song by group and name
// @Description Get a song by its group and name
// @Param group path string true "Group name"
// @Param song path string true "Song name"
// @Success 200 {object} gin.H {"info": "song text"}
// @Failure 404 {object} gin.H {"error": "song not found"}
// @Failure 500 {object} gin.H {"error": "error return song with name"}
// @Router /library/{group}/{song} [get]
func (s *Library) ReturnSongWithName(c *gin.Context) {
	group := c.Param("group")
	song := c.Param("song")

	songStruct, err := s.Storage.ReturnSongDateWithName(group, song)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
			slog.Error("ERROR", slog.Any("err", err))
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error return song with name"})
		slog.Error("ERROR", slog.Any("err", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"info": songStruct.SongText})
}

// @Summary Returns the entire library
// @Description Get all songs in the library
// @Success 200 {object} gin.H {"text": "library data"}
// @Failure 500 {object} gin.H {"error": "error return library"}
// @Router /library [get]
func (s *Library) ReturnLibrary(c *gin.Context) {
	songStruct, err := s.Storage.ReturnLibraryDate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error return library"})
		slog.Error("ERROR", slog.Any("err", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"text": songStruct})
}
