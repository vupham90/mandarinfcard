package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/vupham90/mandarinfcard"
)

func main() {
	app := &cli.App{
		Name:  "Mandarin Flashcard",
		Usage: "Just something to learn Mandarin daily",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "question",
				Aliases: []string{"q"},
			},
			&cli.StringFlag{
				Name:     "slack_url",
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.IntFlag{
				Name:    "time",
				Aliases: []string{"t"},
			},
		},
		Action: Run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Run(ctx *cli.Context) error {
	file, err := os.Open("data.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	dict := mandarinfcard.ReadAll(file)
	picked := mandarinfcard.Pick(time.Now(), int(ctx.Int("time")), len(dict))

	msg := GenMsg(dict[picked], ctx.IsSet("question"))
	err = mandarinfcard.NotiSend(ctx.Context, ctx.String("slack_url"), msg)
	if err != nil {
		return err
	}
	return nil
}

func GenMsg(w mandarinfcard.Word, q bool) string {
	if q {
		return w.Mandarin
	}
	return fmt.Sprintf("%s; %s; %s", w.Mandarin, w.Pinyin, w.English)
}
