CREATE TABLE "player" (
  "player_id"     INTEGER NOT NULL,
  "tournament_id" INTEGER NOT NULL,
  "name"          TEXT    NOT NULL,
  PRIMARY KEY ("player_id" AUTOINCREMENT),
  FOREIGN KEY ("tournament_id") REFERENCES "tournament"("tournament_id")
);

CREATE TABLE "tournament" (
  "tournament_id" INTEGER NOT NULL,
  "title"         TEXT    NOT NULL,
  "subtitle"      TEXT    NOT NULL,
  PRIMARY KEY ("tournament_id" AUTOINCREMENT)
);

CREATE TABLE "pairing" (
  "tournament_id"       INTEGER NOT NULL,
  "round_id"            INTEGER NOT NULL,
  "table_id"            INTEGER NOT NULL,
  "first_player_id"     INTEGER NOT NULL,
  "first_player_score"  INTEGER NOT NULL,
  "second_player_id"    INTEGER NOT NULL,
  "second_player_score" INTEGER NOT NULL,
  PRIMARY KEY ("tournament_id", "round_id", "table_id"),
  FOREIGN KEY ("tournament_id")     REFERENCES "tournament"("tournament_id"),
  FOREIGN KEY ("first_player_id")   REFERENCES "player"("player_id"),
  FOREIGN KEY ("second_player_id")  REFERENCES "player"("player_id")
);
