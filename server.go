package main

import (
	"database/sql"
	"log"
	"net/http"
  _ "embed"

	"github.com/Jlll1/Puchar/templates"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

//go:generate qtc -dir=templates

//go:embed db/schema.sql
var schema string
//go:embed db/init.sql
var initialData string

func main() {
	e := echo.New()

  db, err := sql.Open("sqlite3", ":memory:")
  if err != nil {
    log.Fatal(err)
  }

  if _, err := db.Exec(schema); err != nil {
    log.Fatal(err)
  }
  if _, err := db.Exec(initialData); err != nil {
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
       ,(SELECT MAX("pairing"."round_id")
           FROM "pairing"
          WHERE "pairing"."tournament_id" = 1)
       ,"player1"."name"
       ,"pairing"."first_player_score"
       ,"player2"."name"
       ,"pairing"."second_player_score"
  FROM  "tournament"
  JOIN  "pairing"
    ON  "pairing"."tournament_id" = "tournament"."tournament_id"
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
