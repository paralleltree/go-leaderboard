//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package driver

type ScoringStrategy interface {
	ComposeScore(time, score int64) (int64, error)
	ExtractScore(rawScore int64) int64
}
