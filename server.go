package main

import (
  "database/sql"
  _ "embed"
  "log"
  "net/http"
  "strconv"

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
  e.GET("/", h.getHome)

  e.GET("/tournament/new", h.getNewTournament)
  e.GET("/tournament/:tournamentId", h.getTournament)
  e.POST("/tournament/new", h.postNewTournament)

  e.GET("tournament/:tournamentId/newPlayer", h.getNewPlayer)
  e.POST("tournament/:tournamentId/newPlayer", h.postNewPlayer)

  e.GET("/dashboard", h.getDashboard)

  e.Use(middleware.Logger())
  e.Logger.Fatal(e.Start(":5500"))
}

type Handler struct {
  DB *sql.DB
}

func (h *Handler) getHome(c echo.Context) error {
  return c.HTML(http.StatusOK, templates.Home())
}

func (h *Handler) getNewTournament(c echo.Context) error {
  return c.HTML(http.StatusOK, templates.NewTournament())
}

func (h *Handler) postNewTournament(c echo.Context) error {
  var id int
  err := h.DB.QueryRow(`
INSERT INTO "tournament" (
  "title",
  "subtitle"
)
VALUES ($1, $2)
RETURNING "tournament_id"
`, c.FormValue("title"), c.FormValue("subtitle")).Scan(&id)

  if err != nil {
    c.Error(err)
  }

  return c.Redirect(http.StatusSeeOther, strconv.Itoa(id))
}

func (h *Handler) getNewPlayer(c echo.Context) error {
  tournamentId, err := strconv.Atoi(c.Param("tournamentId"))
  _, err = retrieveTournamentById(h.DB, tournamentId)
  if err != nil {
    if err == sql.ErrNoRows {
      return echo.NewHTTPError(http.StatusNotFound, "Tournament not found")
    } else {
      c.Error(err)
    }
  }

  return c.HTML(http.StatusOK, templates.NewPlayer())
}

func (h *Handler) postNewPlayer(c echo.Context) error {
  tournamentId, err := strconv.Atoi(c.Param("tournamentId"))
  _, err = retrieveTournamentById(h.DB, tournamentId)
  if err != nil {
    if err == sql.ErrNoRows {
      return echo.NewHTTPError(http.StatusNotFound, "Tournament not found")
    } else {
      c.Error(err)
    }
  }

  _, err = h.DB.Exec(`
INSERT INTO "player"(
  "tournament_id"
  ,"name"
)
VALUES ($1, $2)
`, tournamentId, c.FormValue("name"))
  if err != nil {
    c.Error(err)
  }

  return c.Redirect(http.StatusSeeOther, c.Request().Header.Get("referer"))
}

type Tournament struct {
  Title    string
  Subtitle string
}

// @CLEANUP this shouldn't exist
func retrieveTournamentById(db *sql.DB, tournamentId int) (*Tournament, error) {
  tournament := Tournament{}
  err := db.QueryRow(`
SELECT "tournament"."title"
      ,"tournament"."subtitle"
  FROM "tournament"
 WHERE "tournament"."tournament_id" = $1
`, tournamentId).Scan(&tournament.Title, &tournament.Subtitle)

  return &tournament, err
}

func (h *Handler) getTournament(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("tournamentId"))
  if err != nil {
    c.Error(err)
  }

  tournament, err := retrieveTournamentById(h.DB, id)

  if err != nil {
    if err == sql.ErrNoRows {
      return echo.NewHTTPError(http.StatusNotFound, "Not Found")
    } else {
      c.Error(err)
    }
  }

  return c.HTML(http.StatusOK, templates.Tournament(tournament.Title, tournament.Subtitle))
}

func (h *Handler) getDashboard(c echo.Context) error {
  roundParam, err := strconv.Atoi(c.QueryParam("round"))
  if err != nil {
    c.Error(err)
  }

  model := templates.DashboardPage{
    Pairings:      []templates.PairingModel{},
    Standings:     []templates.StandingModel{},
    SelectedRound: roundParam,
  }

  tDetailsAndRoundsRows, err := h.DB.Query(`
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
 WHERE  "tournament"."tournament_id" = 1
   AND  "pairing"."round_id" = $1;`, roundParam)
  if err != nil {
    return err
  }
  defer tDetailsAndRoundsRows.Close()

  for tDetailsAndRoundsRows.Next() {
    p := templates.PairingModel{}
    err = tDetailsAndRoundsRows.Scan(
      &model.TournamentTitle,
      &model.TournamentSubtitle,
      &model.RoundCount,
      &p.Player1Name,
      &p.Player1Score,
      &p.Player2Name,
      &p.Player2Score)
    if err == nil {
      model.Pairings = append(model.Pairings, p)
    }
  }

  // @INCOMPLETE - throwing in a random not actually working query until result-keeping is designed
  standingsRows, err := h.DB.Query(`
   SELECT "player"."name"
         ,COUNT(*) AS "score"
     FROM "player"
     JOIN "pairing"
     ON ("pairing"."first_player_id" = "player"."player_id" 
         AND "pairing"."first_player_score" > "pairing"."second_player_score")
     OR ("pairing"."second_player_id" = "player"."player_id"
         AND "pairing"."second_player_score" > "pairing"."first_player_score")
GROUP BY "player"."player_id"
ORDER BY "score" DESC
         ,"player"."name" ASC;`)
  if err != nil {
    return err
  }
  defer standingsRows.Close()

  for standingsRows.Next() {
    s := templates.StandingModel{}
    err = standingsRows.Scan(&s.PlayerName, &s.Points)
    if err == nil {
      model.Standings = append(model.Standings, s)
    }
  }

  return c.HTML(http.StatusOK, templates.Dashboard(&model))
}
