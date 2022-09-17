package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Jlll1/Puchar/templates"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

  db, err := sql.Open("sqlite3", "puchar.db")
  if err != nil {
    log.Fatal(err)
  }
  h := Handler{DB: db}

  e.GET("/", h.getDashboard)
  e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":5500"))
}

type Handler struct {
  DB *sql.DB
}

func (h *Handler) getDashboard(c echo.Context) error {
  rows, err := h.DB.Query(`
SELECT  "tournament"."title"
       ,"tournament"."subtitle"
       ,(SELECT MAX("round_pairing_xref"."round_id")
           FROM "round_pairing_xref"
          WHERE "round_pairing_xref"."tournament_id" = 1)
       ,"player1"."name"
       ,"pairing"."first_player_score"
       ,"player2"."name"
       ,"pairing"."second_player_score"
  FROM  "tournament"
  JOIN  "round_pairing_xref"
    ON  "round_pairing_xref"."tournament_id" = "tournament"."tournament_id"
  JOIN  "pairing"
    ON  "pairing"."pairing_id" = "round_pairing_xref"."pairing_id"
  JOIN  "player" AS "player1"
    ON  "player1"."player_id" = "pairing"."first_player_id"
  JOIN  "player" AS "player2"
    ON  "player2"."player_id" = "pairing"."second_player_id"
 WHERE  "tournament"."tournament_id" = 1`)
  if err != nil {
    return err
  }
  defer rows.Close()

  model := templates.DashboardPage{Pairings: []templates.PairingModel{}}
  for rows.Next() {
    p := templates.PairingModel{}
    err = rows.Scan(
      &model.TournamentTitle,
      &model.TournamentSubtitle,
      &model.RoundCount,
      &p.Player1Name,
      &p.Player1Score,
      &p.Player2Name,
      &p.Player2Score)
    model.Pairings = append(model.Pairings, p)
  }
  return c.HTML(http.StatusOK, templates.Dashboard(&model))
}
