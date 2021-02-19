package math

import (
	"fmt"
	"image/color"
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vdobler/chart"

	"github.com/theovidal/bacbot/lib"
)

func PlotCommand() lib.Command {
	return lib.Command{
		Description: fmt.Sprintf("Tracer des graphiques riches et complets. Vous pouvez tracer plusieurs fonctions en séparant leurs expressions par une esperluette `&`.\n%s\n\n%s", dataDocumentation, calcDisclaimer),
		Flags: map[string]lib.Flag{
			"x_min":   {"Valeur minimale de `x` à afficher", -10.0},
			"x_max":   {"Valeur maximale de `x` à afficher", 10.0},
			"x_scale": {"Pas pour l'abscisse", 1.0},
			"y_min":   {"Valeur minimale de `y` à afficher", -10.0},
			"y_max":   {"Valeur maximale de `y` à afficher", 10.0},
			"y_scale": {"Pas pour l'ordonnée", 1.0},

			// "color":      {"Couleur de la courbe : `red`, `pink`, `purple`, `indigo`, `blue`, `light_blue`, `cyan`, `teal`, `green`, `light_green`, `lime`, `yellow`, `amber`, `orange`, `brown`.", "red"},
			"line_width": {"Épaisseur de la courbe (en pixels)", 1},

			"grid": {"Afficher la grille sur le graphique (0 ou 1)", 1},
		},
		Execute: func(bot *telegram.BotAPI, update *telegram.Update, args []string, flags map[string]interface{}) error {
			/* lineColor, exists := lib.Colors[flags["color"].(string)]
			if !exists {
				return lib.Error(bot, update, "La couleur spécifiée n'existe pas. Vérifiez la liste des couleurs disponibles sur la page d'aide de la commande.")
			} */

			grid := flags["grid"].(int)

			raw := strings.Join(args, " ")
			functions := strings.Split(raw, "&")

			plot := chart.ScatterChart{}
			plot.XRange.MinMode.Fixed, plot.XRange.MaxMode.Fixed = true, true
			plot.XRange.MinMode.Value, plot.XRange.MaxMode.Value = flags["x_min"].(float64), flags["x_max"].(float64)
			if grid == 1 {
				plot.XRange.TicSetting.Grid = chart.GridLines
			}

			plot.YRange.MinMode.Fixed, plot.YRange.MaxMode.Fixed = true, true
			plot.YRange.MinMode.Value, plot.YRange.MaxMode.Value = flags["y_min"].(float64), flags["y_max"].(float64)
			if grid == 1 {
				plot.YRange.TicSetting.Grid = chart.GridLines
			}

			plot.XRange.TicSetting.Delta = flags["x_scale"].(float64)
			plot.YRange.TicSetting.Delta = flags["y_scale"].(float64)

			plot.XRange.TicSetting.Mirror = -1
			plot.YRange.TicSetting.Mirror = -1

			style := chart.Style{
				Symbol:    'o',
				FillColor: color.NRGBA{0xff, 0x80, 0x80, 0xff},
				LineStyle: chart.SolidLine,
				LineWidth: flags["line_width"].(int),
			}

			for _, function := range functions {
				msg := CheckExpression(function)
				if msg != "" {
					return lib.Error(bot, update, msg)
				}
			}

			config := telegram.NewMessage(update.Message.Chat.ID, "_Génération du graphique en cours..._")
			config.ParseMode = "Markdown"
			waiter, _ := bot.Send(config)

			colorNumber := 0
			for _, function := range functions {
				current := strings.TrimSpace(function)
				if colorNumber == len(lib.Colors) {
					colorNumber = 0
				}
				style.LineColor = lib.Colors[colorNumber]
				colorNumber++

				plot.AddFunc(current, func(x float64) float64 {
					return Evaluate(current, x)
				}, chart.PlotStyleLines, style)
			}

			file := lib.Plot(&plot, "function_plot")
			photo := telegram.NewPhotoUpload(update.Message.Chat.ID, file)
			_, err := bot.Send(photo)
			bot.DeleteMessage(telegram.DeleteMessageConfig{
				ChatID:    waiter.Chat.ID,
				MessageID: waiter.MessageID,
			})
			return err
		},
	}
}
