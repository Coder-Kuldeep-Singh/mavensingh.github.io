package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Password Strength Checker")

	passwordEntry := widget.NewEntry()
	passwordEntry.SetPlaceHolder("Enter Password")
	passwordStrengthLabel := widget.NewLabel("")
	timeToCrackLabel := widget.NewLabel("")
	var checkButton *widget.Button
	checkButton = widget.NewButton("Check", func() {
		passwordStrengthLabel.SetText("")
		timeToCrackLabel.SetText("")
		progressBar := widget.NewProgressBarInfinite()
		progressBarLabel := widget.NewLabel("Password strength checking started: Making Brute Force Attack")
		progressBarContainer := container.NewVBox(progressBarLabel, progressBar)
		w.SetContent(container.NewVBox(container.NewHBox(passwordEntry, checkButton), passwordStrengthLabel, timeToCrackLabel, progressBarContainer))

		go func() {
			password := passwordEntry.Text
			timeToCrack := calculateTimeToCrack(calculatePasswordStrength(password))
			progressBarContainer.Hide()
			passwordStrengthLabel.SetText(fmt.Sprintf("Password Strength: %s", calculatePasswordStrength(password)))
			timeToCrackLabel.SetText(fmt.Sprintf("Time to Crack: %s", timeToCrack))
		}()
	})

	w.SetContent(container.NewVBox(container.NewHBox(passwordEntry, checkButton), passwordStrengthLabel, timeToCrackLabel))
	w.ShowAndRun()
}

func calculatePasswordStrength(password string) string {
	length := len(password)
	complexity := calculatePasswordComplexity(password)
	score := int(math.Floor(float64(length)*complexity/2.5 - 5.5))
	if score < 0 {
		score = 0
	} else if score > 4 {
		score = 4
	}
	strengths := []string{"Very Weak", "Weak", "Moderate", "Strong", "Very Strong"}
	return strengths[score]
}

func calculatePasswordComplexity(password string) float64 {
	length := len(password)
	sets := []string{"abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "0123456789", "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"}
	diversity := 0
	for _, set := range sets {
		if strings.ContainsAny(password, set) {
			diversity++
		}
	}
	complexity := float64(length) * math.Pow(float64(diversity), 1.5) / math.Pow(float64(len(sets)), 1.5)
	return complexity
}

func calculateTimeToCrack(strength string) string {
	var timeToCrack time.Duration
	switch strings.ToLower(strength) {
	case "very weak":
		timeToCrack = time.Duration(math.Pow(62, 8)) * time.Second / 1e9
	case "weak":
		timeToCrack = time.Duration(math.Pow(62, 10)) * time.Second / 1e9
	case "moderate":
		timeToCrack = time.Duration(math.Pow(62, 12)) * time.Second / 1e9
	case "strong":
		timeToCrack = time.Duration(math.Pow(62, 14)) * time.Second / 1e9
	case "very strong":
		timeToCrack = time.Duration(math.Pow(62, 16)) * time.Second / 1e9
	}
	return timeToCrack.String()
}
