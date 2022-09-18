INSERT INTO "player" ("name")
VALUES
  ("Player Primary"),
  ("Byk Byczy"),
  ("Char Mander"),
  ("Krakers Pies");

INSERT INTO "tournament" ("title", "subtitle")
VALUES
  ("Exceptional Championship", "The best of the best face each other");

INSERT INTO "pairing" (
  "tournament_id",
  "round_id",
  "table_id",
  "first_player_id",
  "first_player_score",
  "second_player_id",
  "second_player_score")
VALUES
--t_id | r_id | t_id | fp_id | fp_s | sp_id | sp_s
  (  1,     1,     1,     1,     2,      4,     0),
  (  1,     1,     2,     2,     2,      3,     0),
  (  1,     2,     1,     1,     0,      2,     2),
  (  1,     2,     2,     3,     1,      4,     1),
  (  1,     3,     1,     4,     2,      1,     0),
  (  1,     3,     2,     3,     2,      2,     0),
  (  1,     4,     1,     2,     0,      1,     2),
  (  1,     4,     2,     4,     1,      3,     1);
