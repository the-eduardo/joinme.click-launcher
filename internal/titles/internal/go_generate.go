//go:build ignore

package internal

//go:generate mockgen -destination=mock_test.go -package=$GOPACKAGE -write_package_comment=false "github.com/cetteup/joinme.click-launcher/pkg/game_launcher" FileRepository
