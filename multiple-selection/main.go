package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
)

type item struct {
	ID         string
	IsSelected bool
}

func main() {
	items := []*item{
		{
			ID: "Cat",
		},
		{
			ID: "Dog",
		},
		{
			ID: "Horse",
		},
		{
			ID: "Parrot",
		},
		{
			ID: "Zebra",
		},
	}
	selected, err := selectItems(0, items)
	if err != nil {
		log.Fatal(err)
	}

	// Print selected items
	fmt.Println("Selected:")
	for _, s := range selected {
		fmt.Println(s.ID)
	}
}

// selectItems() prompts user to select one or more items in the given slice
func selectItems(selectedPos int, allItems []*item) ([]*item, error) {
	// Always prepend a "Done" item to the slice if it doesn't
	// already exist.
	const doneID = "Done"
	if len(allItems) > 0 && allItems[0].ID != doneID {
		var items = []*item{
			{
				ID: doneID,
			},
		}

		allItems = append(items, allItems...)
	}

	// Define promptui template
	templates := &promptui.SelectTemplates{
		Label: `{{if .IsSelected}}
					✔
				{{end}} {{ .ID }} - label`,
		Active:   "→ {{if .IsSelected}}✔ {{end}}{{ .ID | cyan }}",
		Inactive: "{{if .IsSelected}}✔ {{end}}{{ .ID | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Item",
		Items:     allItems,
		Templates: templates,
		Size:      5,
		// Start the cursor at the currently selected index
		CursorPos:    selectedPos,
		HideSelected: true,
	}

	selectionIdx, _, err := prompt.Run()
	if err != nil {
		return nil, fmt.Errorf("prompt failed: %w", err)
	}

	chosenItem := allItems[selectionIdx]

	if chosenItem.ID != doneID {
		// If the user selected something other than "Done",
		// toggle selection on this item and run the function again.
		chosenItem.IsSelected = !chosenItem.IsSelected
		return selectItems(selectionIdx, allItems)
	}

	// If the user selected the "Done" item, return
	// all selected items.
	var selectedItems []*item
	for _, i := range allItems {
		if i.IsSelected {
			selectedItems = append(selectedItems, i)
		}
	}
	return selectedItems, nil
}
