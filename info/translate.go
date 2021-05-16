package info

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/theovidal/bacbot/lib"
)

// DeeplUrl is the endpoint to the translation API
const DeeplUrl = "https://api-free.deepl.com/v2/translate"

// DeeplResponse stores the response from DeepL API
type DeeplResponse struct {
	Translations []Translation
}

// Translation represents the translation of a sentence or more from a certain language
type Translation struct {
	// The language from which the text was translated (usually automatically determined by DeepL)
	SourceLanguage string `json:"detected_source_language"`
	// The translated text
	Text string
}

func TranslateCommand() lib.Command {
	return lib.Command{
		Name:        "translate",
		Description: "Traduire un texte entier avec contexte (DeepL)",
		Flags: map[string]lib.Flag{
			"source": {Description: "Manuellement inscrire la langue source", Value: ""},
		},
		Execute: func(bot *telegram.BotAPI, update *telegram.Update, args []string, flags map[string]interface{}) (err error) {
			if len(args) < 2 {
				return lib.Error(bot, update, "Indiquez la langue cible ainsi que le texte à traduire.")
			}
			from := strings.ToUpper(flags["source"].(string))
			to := strings.ToUpper(args[0])
			text := strings.Join(args[1:], " ")

			if _, exists := sourceLanguages[from]; !exists && from != "" {
				return lib.Error(bot, update, "La langue source est invalide. Veuillez choisir parmi "+sourceLanguagesDoc)
			}

			if _, exists := targetLanguages[to]; !exists {
				return lib.Error(bot, update, "La langue cible est invalide. Veuillez choisir parmi "+targetLanguagesDoc)
			}

			request, _ := http.NewRequest(
				"GET",
				lib.EncodeURL(DeeplUrl, map[string]string{
					"auth_key":    os.Getenv("DEEPL_KEY"),
					"text":        text,
					"source_lang": from,
					"target_lang": to,
				}),
				nil,
			)

			response, err := lib.DoRequest(request)
			if err != nil {
				return
			}

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return
			}
			response.Body.Close()

			var result DeeplResponse
			err = json.Unmarshal(body, &result)
			if err != nil {
				return
			}

			var content string
			for _, translation := range result.Translations {
				content += fmt.Sprintf(
					"*―――――― 💱 %s → %s ――――――*\n%s",
					sourceLanguages[translation.SourceLanguage],
					targetLanguages[to],
					translation.Text,
				)
			}

			msg := telegram.NewMessage(update.Message.Chat.ID, content)
			msg.ParseMode = "Markdown"
			_, err = bot.Send(msg)
			return
		},
	}
}

// sourceLanguages stores keys and flags for supported languages as sources
var sourceLanguages = map[string]string{
	"BG": "🇧🇬",
	"CS": "🇨🇿",
	"DA": "🇩🇰",
	"DE": "🇩🇪",
	"EL": "🇬🇷",
	"EN": "🇬🇧",
	"ES": "🇪🇸",
	"ET": "🇪🇪",
	"FI": "🇫🇮",
	"FR": "🇫🇷",
	"HU": "🇭🇺",
	"IT": "🇮🇹",
	"JA": "🇯🇵",
	"LT": "🇱🇹",
	"LV": "🇱🇻",
	"NL": "🇳🇱",
	"PL": "🇵🇱",
	"PT": "🇵🇹",
	"RO": "🇷🇴",
	"RU": "🇷🇺",
	"SK": "🇸🇰",
	"SL": "🇸🇮",
	"SV": "🇸🇪",
	"ZH": "🇨🇳",
}

// targetLanguages stores keys and flags for supported languages as targets
var targetLanguages = func() (output map[string]string) {
	output = map[string]string{
		"EN-GB": "🇬🇧",
		"EN-US": "🇺🇸",
		"PT-PT": "🇵🇹",
		"PT-BR": "🇧🇷",
	}
	for language, flag := range sourceLanguages {
		output[language] = flag
	}
	return
}()

var sourceLanguagesDoc = generateDoc(sourceLanguages)
var targetLanguagesDoc = generateDoc(targetLanguages)

// generateDoc generates the documentation text for languages list
func generateDoc(input map[string]string) string {
	var languages []string
	for lang := range input {
		languages = append(languages, lang)
	}
	return strings.Join(languages, ", ")
}
