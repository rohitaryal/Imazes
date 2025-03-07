package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/rohitaryal/imazes/internal/argv"
	"github.com/rohitaryal/imazes/pkg/imazes"
)

func main() {
	argv.Init("rohitaryal", "imazen", "Generate AI images from your terminal")

	argv.AddArg("prompt", "p", "Prompt string that describes an image", "A dumb person who forgot to give prompt to AI")
	argv.AddArg("negative", "n", "Avoid generating specific elements or characteristics", "Deformed hands and faces, blurry, ugly")
	argv.AddArg("style", "s", "Name of unique visual characteristics and aesthetic", "SDXL 1.0")
	argv.AddArg("count", "c", "Number of images you want to generate", "1")
	argv.AddArg("steps", "k", "Number of iterative denoising processes the model performs", "40")
	argv.AddArg("ratio", "r", "Aspect ratio or size of generated image", "1:1")

	if slices.Contains(os.Args, "--help") {
		argv.PrintHelp()
		return
	}

	prompt, _ := argv.GetArg("prompt")
	negative, _ := argv.GetArg("negative")
	style, _ := argv.GetArg("style")
	count, _ := argv.GetArg("count")
	steps, _ := argv.GetArg("steps")
	ratio, _ := argv.GetArg("ratio")

	if !slices.Contains(imazes.Styles, style) {
		printStyles()
		return
	}

	parsedCount, _ := strconv.Atoi(count)
	if parsedCount > 20 {
		fmt.Println("[!] Image count must be less than ")
		return
	}

	parsedSteps, _ := strconv.Atoi(steps)
	if parsedSteps > 50 {
		fmt.Println("[!] Step count must be less than 50")
		return
	}

	if !slices.Contains(imazes.Ratios, ratio) {
		printRatio()
		return
	}

	imageDescription := imazes.Image{
		Prompt:   prompt,
		Negative: negative,
		Style:    style,
		Count:    count,
		Steps:    steps,
		Ratio:    ratio,
	}

	token := imazes.GenerateToken().IdToken
	fmt.Println("[+] Obtained Token")
	generator := imazes.GenerateImage(imageDescription, token, true)
	fmt.Println("[+] Status ID: ", generator.RecordID)
	for {
		status := imazes.GetImage(generator.RecordID, token)
		fmt.Println("[+] Current Status: ", status.Status)
		if status.Status != "DONE" {
			time.Sleep(10 * time.Second)
			continue
		} else {
			var images []imazes.ImageResponse = status.Response
			for _, image := range images {
				fmt.Println("[+] Image URL: ", image.URL)
			}
			return
		}
	}
}

func printStyles() {
	fmt.Printf("[!] Invalid style. Available Styles include:\n\n")
	for index, style := range imazes.Styles {
		fmt.Printf("%-20s", style)
		if index%5 == 0 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
}

func printRatio() {
	fmt.Printf("[!] Invalid ratio. Available ratio include:\n\n")
	for index, ratio := range imazes.Ratios {
		fmt.Printf("%-5s", ratio)
		if index%5 == 0 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
}
