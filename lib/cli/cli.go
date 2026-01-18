package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	util "raha_tracker/lib"
	"raha_tracker/lib/iniutil"
	"strings"

	"github.com/manifoldco/promptui"
)

func promptCategoryChoice(items []string) {
	prompt := promptui.Select{
		Label:  "Kategooria tehingule puudub, lisa kategooria?",
		Items:  []string{"yes", "no"},
		Stdout: os.Stdout,
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed\n", err)
	}
	fmt.Println(result)
}

func promptCategoryUpdate(items []string) {
	prompt := promptui.Select{
		Label:  "Maksele lisatud mitu kategooriat, palun vali milline kategooria on rohkem sobilik",
		Items:  items,
		Stdout: os.Stdout,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal("Prompt failed\n", err)
	}

	fmt.Printf("You choose %q\n", result)
}

func promptConfigUpdate() string {
	prompt := promptui.Select{
		Label:  "Tehingu kaupmees puudub, salvesta kaupmees?",
		Items:  []string{"jah", "ei"},
		Stdout: os.Stdout,
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed\n", err)
	}
	return result
}

func RunPrompts(r *util.Record) {
	// log for the user
	fmt.Println("\033[32m", "TEHING:", r.Timestamp, ":", r.Supplier, "on kandnud raha. Summa:", r.Amount, r.Currency, "\033[0m")

	// check if transaction vendor in config
	if !iniutil.SupplierInConfig(r.Supplier) {
		choice := promptConfigUpdate()
		switch choice {
		case "jah":
			r.SectionName = PromptSectionName()
			r.Category = PromptCategoryName()
			if r.SectionName != "" && r.Category != "" {
				iniutil.AddConfigEntry(r.SectionName, r.Category)
			}
		case "ei":
			r.Category = PromptSingleCategory()
		}
	}

	items := iniutil.FindCategories(r.Supplier)

	if len(items) > 1 {
		promptCategoryUpdate(items)
	}

	if len(items) == 1 {
		r.Category = items[0]
	}

	if len(items) == 0 {

	}
}

func PromptSectionName() string {
	// using bufio reader here due to promptui bugging with back to back prompts
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Sisesta kaupmehe nimi")
	fmt.Println("Näide: kaupmees:RIMI/TARTU LOUNAKESKUS --> rimi")

	fmt.Print("> ")

	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Read failed: %v\n", err)
		return ""
	}

	result = strings.TrimSpace(result)

	return result
}

func PromptCategoryName() string {
	// using bufio reader here due to promptui bugging with back to back prompts
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Sisesta kategooria mille alla salvestada\n")
	fmt.Println("Et lisada mitu kategooriat, pane kategooriate vahele koma NÄIDE: söök,meelelahutus")

	fmt.Print("> ")

	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Read failed: %v\n", err)
		return ""
	}

	result = strings.TrimSpace(result)

	return result
}

func PromptSingleCategory() string {
	// using bufio reader here due to promptui bugging with back to back prompts
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Sisesta kategooria mille alla tehing salvestada (max 1)\n")
	fmt.Print("> ")

	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Read failed: %v\n", err)
		return ""
	}

	result = strings.TrimSpace(result)

	return result
}
