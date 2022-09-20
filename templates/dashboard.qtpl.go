// Code generated by qtc from "dashboard.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/dashboard.qtpl:1
package templates

//line templates/dashboard.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/dashboard.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/dashboard.qtpl:3
type StandingModel struct {
	PlayerName string
	Points     int
}

type PairingModel struct {
	Player1Name  string
	Player1Score int
	Player2Name  string
	Player2Score int
}

type DashboardPage struct {
	TournamentTitle    string
	TournamentSubtitle string
	SelectedRound      int
	RoundCount         int
	Pairings           []PairingModel
	Standings          []StandingModel
}

//line templates/dashboard.qtpl:26
func StreamDashboard(qw422016 *qt422016.Writer, p *DashboardPage) {
//line templates/dashboard.qtpl:26
	qw422016.N().S(`
<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
  <title>Puchar</title>

  <style>
    .badge-win,
    .badge-defeat,
    .badge-draw {
      color: #FFFFFF;
      border-radius: 3px;
      text-align: center;
      width: 25px;
    }

    .badge-win {
      background: #4BB543;
    }

    .badge-defeat {
      background: #FF0033;
    }

    .badge-draw {
      background: #888888;
    }
  </style>

</head>


<body>

<div class="container">

<section class="hero">
  <div class="hero-body">
    <p class="title">
    `)
//line templates/dashboard.qtpl:69
	qw422016.E().S(p.TournamentTitle)
//line templates/dashboard.qtpl:69
	qw422016.N().S(`
    </p>
    <p class="subtitle">
    `)
//line templates/dashboard.qtpl:72
	qw422016.E().S(p.TournamentSubtitle)
//line templates/dashboard.qtpl:72
	qw422016.N().S(`
    </p>
  </div>
</section>
<section>
  <div class="columns">
    <div class="column is-three-quarters" style="padding-right: 50px;">
    <nav class="pagination is-small">
      <ul class="pagination-list">
        `)
//line templates/dashboard.qtpl:81
	for i := 1; i <= p.RoundCount; i++ {
//line templates/dashboard.qtpl:81
		qw422016.N().S(`
          `)
//line templates/dashboard.qtpl:82
		if i == p.SelectedRound {
//line templates/dashboard.qtpl:82
			qw422016.N().S(`
            <li><a href="/dashboard?round=`)
//line templates/dashboard.qtpl:83
			qw422016.N().D(i)
//line templates/dashboard.qtpl:83
			qw422016.N().S(`" class="pagination-link is-current">`)
//line templates/dashboard.qtpl:83
			qw422016.N().D(i)
//line templates/dashboard.qtpl:83
			qw422016.N().S(`</a></li>
          `)
//line templates/dashboard.qtpl:84
		} else {
//line templates/dashboard.qtpl:84
			qw422016.N().S(`
            <li><a href="/dashboard?round=`)
//line templates/dashboard.qtpl:85
			qw422016.N().D(i)
//line templates/dashboard.qtpl:85
			qw422016.N().S(`" class="pagination-link">`)
//line templates/dashboard.qtpl:85
			qw422016.N().D(i)
//line templates/dashboard.qtpl:85
			qw422016.N().S(`</a></li>
          `)
//line templates/dashboard.qtpl:86
		}
//line templates/dashboard.qtpl:86
		qw422016.N().S(`
        `)
//line templates/dashboard.qtpl:87
	}
//line templates/dashboard.qtpl:87
	qw422016.N().S(`
      </ul>
    </nav>
    <table class="table is-fullwidth">
      <thead>
        <tr>
          <th style="width: 50px;"><abbr title="Board">#</abbr></th>
          <th class="has-text-left">Name</th>
          <th class="has-text-centered">Result</th>
          <th class="has-text-right">Name</th>
        </tr>
      </thead>
      <tbody>
        `)
//line templates/dashboard.qtpl:100
	for _, pairing := range p.Pairings {
//line templates/dashboard.qtpl:100
		qw422016.N().S(`
          <tr>
            <th>?</th>
            <td class="has-text-left">`)
//line templates/dashboard.qtpl:103
		qw422016.E().S(pairing.Player1Name)
//line templates/dashboard.qtpl:103
		qw422016.N().S(`</td>
            <td class="is-flex is-justify-content-center">
              `)
//line templates/dashboard.qtpl:105
		if pairing.Player1Score > pairing.Player2Score {
//line templates/dashboard.qtpl:105
			qw422016.N().S(`
                <span class="badge-win mr-1">`)
//line templates/dashboard.qtpl:106
			qw422016.N().D(pairing.Player1Score)
//line templates/dashboard.qtpl:106
			qw422016.N().S(`</span>
                <span class="badge-defeat">`)
//line templates/dashboard.qtpl:107
			qw422016.N().D(pairing.Player2Score)
//line templates/dashboard.qtpl:107
			qw422016.N().S(`</span>
              `)
//line templates/dashboard.qtpl:108
		} else {
//line templates/dashboard.qtpl:108
			qw422016.N().S(`
                `)
//line templates/dashboard.qtpl:109
			if pairing.Player1Score < pairing.Player2Score {
//line templates/dashboard.qtpl:109
				qw422016.N().S(`
                  <span class="badge-defeat mr-1">`)
//line templates/dashboard.qtpl:110
				qw422016.N().D(pairing.Player1Score)
//line templates/dashboard.qtpl:110
				qw422016.N().S(`</span>
                  <span class="badge-win">`)
//line templates/dashboard.qtpl:111
				qw422016.N().D(pairing.Player2Score)
//line templates/dashboard.qtpl:111
				qw422016.N().S(`</span>
                `)
//line templates/dashboard.qtpl:112
			} else {
//line templates/dashboard.qtpl:112
				qw422016.N().S(`
                  <span class="badge-draw mr-1">`)
//line templates/dashboard.qtpl:113
				qw422016.N().D(pairing.Player1Score)
//line templates/dashboard.qtpl:113
				qw422016.N().S(`</span>
                  <span class="badge-draw">`)
//line templates/dashboard.qtpl:114
				qw422016.N().D(pairing.Player2Score)
//line templates/dashboard.qtpl:114
				qw422016.N().S(`</span>
                `)
//line templates/dashboard.qtpl:115
			}
//line templates/dashboard.qtpl:115
			qw422016.N().S(`
              `)
//line templates/dashboard.qtpl:116
		}
//line templates/dashboard.qtpl:116
		qw422016.N().S(`
            </td>
            <td class="has-text-right">`)
//line templates/dashboard.qtpl:118
		qw422016.E().S(pairing.Player2Name)
//line templates/dashboard.qtpl:118
		qw422016.N().S(`</td>
          </tr>
        `)
//line templates/dashboard.qtpl:120
	}
//line templates/dashboard.qtpl:120
	qw422016.N().S(`
      </tbody>
    </table>
    </div>

    <div class="column">
      <table class="table is-fullwidth">
        <thead>
          <tr>
            <th style="width: 50px;"><abbr title="Place">#</abbr></th>
            <th>Player</th>
            <th>Points</th>
          </tr>
        </thead>
        <tbody>
          `)
//line templates/dashboard.qtpl:135
	for i, s := range p.Standings {
//line templates/dashboard.qtpl:135
		qw422016.N().S(`
            <tr>
              <td>`)
//line templates/dashboard.qtpl:137
		qw422016.N().D(i + 1)
//line templates/dashboard.qtpl:137
		qw422016.N().S(`</td>
              <td>`)
//line templates/dashboard.qtpl:138
		qw422016.E().S(s.PlayerName)
//line templates/dashboard.qtpl:138
		qw422016.N().S(`</td>
              <td>`)
//line templates/dashboard.qtpl:139
		qw422016.N().D(s.Points)
//line templates/dashboard.qtpl:139
		qw422016.N().S(`</td>
            </tr>
          `)
//line templates/dashboard.qtpl:141
	}
//line templates/dashboard.qtpl:141
	qw422016.N().S(`
        </tbody>
      </table>
    </div>
  </div>
</section>

</div>

</body>

</html>
`)
//line templates/dashboard.qtpl:153
}

//line templates/dashboard.qtpl:153
func WriteDashboard(qq422016 qtio422016.Writer, p *DashboardPage) {
//line templates/dashboard.qtpl:153
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/dashboard.qtpl:153
	StreamDashboard(qw422016, p)
//line templates/dashboard.qtpl:153
	qt422016.ReleaseWriter(qw422016)
//line templates/dashboard.qtpl:153
}

//line templates/dashboard.qtpl:153
func Dashboard(p *DashboardPage) string {
//line templates/dashboard.qtpl:153
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/dashboard.qtpl:153
	WriteDashboard(qb422016, p)
//line templates/dashboard.qtpl:153
	qs422016 := string(qb422016.B)
//line templates/dashboard.qtpl:153
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/dashboard.qtpl:153
	return qs422016
//line templates/dashboard.qtpl:153
}
