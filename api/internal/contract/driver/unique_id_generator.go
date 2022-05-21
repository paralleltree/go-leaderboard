//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
package driver

type UniqueIdGenerator interface {
	GenerateNewId() (string, error)
}
