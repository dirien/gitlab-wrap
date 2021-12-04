package card

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"gitlab-wrap/internal/gitlab"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed assets/card.png
var card []byte

//go:embed assets/SourceSansPro-Regular.ttf
var font []byte

//go:embed assets/404.png
var notFound []byte

func CreateUserNotFound(username string) (*string, error) {
	err := os.WriteFile(fmt.Sprintf("%s/404.png", os.TempDir()), notFound, 0644)
	if err != nil {
		return nil, err
	}
	bgImage, err := gg.LoadImage(fmt.Sprintf("%s/404.png", os.TempDir()))
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(fmt.Sprintf("%s/SourceSansPro-Regular.ttf", os.TempDir()), 40.0); err != nil {
		return nil, err
	}
	printText(dc, fmt.Sprintf("Hooman, %s is not found!", username), 200, 70.0, true)
	outPutFileName := fmt.Sprintf("%s/%s.png", os.TempDir(), username)
	err = gg.SavePNG(outPutFileName, dc.Image())
	return &outPutFileName, nil
}

func CreateGitLabWrapCard(userCard *gitlab.WrapStats) (*string, error) {
	err := os.WriteFile(fmt.Sprintf("%s/card.png", os.TempDir()), card, 0644)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(fmt.Sprintf("%s/SourceSansPro-Regular.ttf", os.TempDir()), font, 0644)
	if err != nil {
		return nil, err
	}

	bgImage, err := gg.LoadImage(fmt.Sprintf("%s/card.png", os.TempDir()))
	if err != nil {
		return nil, err
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(fmt.Sprintf("%s/SourceSansPro-Regular.ttf", os.TempDir()), 40.0); err != nil {
		return nil, err
	}

	err = downloadAvatar(userCard.User.AvatarURL, fmt.Sprintf("%s/%d_avatar.png", os.TempDir(), userCard.User.ID))
	if err != nil {
		log.Printf("Error downloading avatar: %s", err)
	} else {
		avatar, err := gg.LoadImage(fmt.Sprintf("%s/%d_avatar.png", os.TempDir(), userCard.User.ID))
		if err != nil {
			return nil, err
		}
		dc.DrawCircle(1050, 140, 100)
		dc.Clip()
		dc.DrawImage(avatar, 950, 40)
		dc.ResetClip()
	}

	printText(dc, userCard.User.Username, 990, 315.0, false)
	printText(dc, userCard.User.CreatedAt.Format("02 Jan 2006"), 990, 400.0, false)
	printText(dc, fmt.Sprintf("%d", userCard.StarredProjects), 990, 490.0, false)
	printText(dc, fmt.Sprintf("%d", userCard.User.Followers), 990, 565.0, false)

	printText(dc, fmt.Sprintf("%d", userCard.ProjectSum), 570, 255.0, false)
	printText(dc, fmt.Sprintf("%d", userCard.IssuesSum), 570, 380.0, false)
	printText(dc, fmt.Sprintf("%d", userCard.MergeRequestsSum), 570, 500.0, false)

	outPutFileName := fmt.Sprintf("%s/%d.png", os.TempDir(), userCard.User.ID)
	err = gg.SavePNG(outPutFileName, dc.Image())
	return &outPutFileName, nil
}

func printText(dc *gg.Context, txt string, x, y float64, isWhite bool) {
	if isWhite == true {
		dc.SetRGB(1, 1, 1)
		dc.DrawStringWrapped(txt, x, y, 0.5, 0.5, 300, 1.5, gg.AlignLeft)
	} else {
		dc.SetColor(color.Gray16{Y: 0x9999})
		dc.DrawString(txt, x, y)
	}

}

func downloadAvatar(URL, fileName string) error {
	url := fmt.Sprintf("%s", URL)
	if strings.Contains(url, "s=80") {
		url = strings.Replace(url, "s=80", "s=200", -1)
	}
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("could not download avatar")
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
