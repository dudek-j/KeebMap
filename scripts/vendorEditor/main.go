package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	// Disable timpestamps in logger
	configureLogger()

	// Grab vendorList path from arguemnts and valide it
	path := parseArguments()
	path, err := verifyPath(path)

	if err != nil {
		if err == errNoListInDirectory {
			promptForNewListFile(&path)
		} else {
			log.Fatal("Incorrect Path:", err)
		}
	}

	in, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed opening vendorData.json:", err)
	}

	// Read all data before unmarshalling
	data, err := ioutil.ReadAll(in)
	if err != nil {
		log.Fatal("Data read error:", err)
	}
	in.Close()

	var vendorList VendorList

	err = json.Unmarshal(data, &vendorList)
	if err != nil {
		log.Fatal("failed when parsing json data:", err)
	}

	// Start interactive loop
	for {
		operation := promptForOperation(len(vendorList))
		handleOpertion(operation, &vendorList, path)
	}

}

type VendorList []Vendor

type Vendor struct {
	Name    string
	Url     string
	Country string
	Region  string
	Hide    bool `json:"hide,omitempty"`
}

func promptForOperation(vendorCount int) string {

	promptForOperation := promptui.Select{
		Label: fmt.Sprintf("Imported %v vendors", vendorCount),
		Items: []string{
			"Add",
			"Verify Sorted",
			"Search and Edit",
			"Save and Exit",
			"Exit without Changes",
		},
		HideSelected: true,
	}

	_, result, err := promptForOperation.Run()

	if err != nil {
		log.Fatal("internal prompt error:", err)
	}

	return result
}

func handleOpertion(operation string, vendorList *VendorList, path string) {
	switch operation {
	case "Add":
		addPromptCoordinator(vendorList)
	case "Verify Sorted":
		verifyList(*vendorList)
	case "Search and Edit":
		promptSearchAndEditCoordinator(vendorList)
	case "Save and Exit":
		writeListToFile(path, *vendorList)
		os.Exit(0)
	case "Exit without Changes":
		log.Println("Exited without saving")
		os.Exit(0)
	default:
		log.Fatal("Internal Error, incorrect operation type")
	}
}

func promptSearchAndEditCoordinator(list *VendorList) {
	for {
		index := searchVendorPrompt(*list)
		err := editVendorPrompt(list, index)

		if err == errBackSelected {
			continue
		}

		return
	}
}

var errBackSelected = errors.New("back to list selected")

func editVendorPrompt(list *VendorList, index int) error {

	selectVendorPtr := &(*list)[index]

	promptForOperation := promptui.Select{
		Label:        fmt.Sprintf("Selected %v ", selectVendorPtr.Name),
		Items:        []string{"Edit Vendor", "Delete", "Toggle Hide", "Back to List", "Exit to Start"},
		HideSelected: true,
	}

	_, res, err := promptForOperation.Run()

	if err != nil {
		log.Fatal("internal prompt error:", err)
	}

	switch res {
	case "Edit Vendor":
		vendorCopy := *selectVendorPtr
		addVendorPromptSequence(&vendorCopy)
		removeVendorFromList(list, index)
		insertIntoSortedList(list, vendorCopy)
		fmt.Println("Modified vendor:", selectVendorPtr.Name)
	case "Delete":
		name := selectVendorPtr.Name
		removeVendorFromList(list, index)
		fmt.Println("Deleted vendor from list:", name)
	case "Toggle Hide":
		selectVendorPtr.Hide = !selectVendorPtr.Hide
		fmt.Println("Toggled hidden:", selectVendorPtr.Name)
	case "Back to List":
		return errBackSelected
	case "Exit to Start":
		return nil
	default:
		log.Fatal("internal error, unknown prompt option")
	}

	return nil

}

func removeVendorFromList(list *VendorList, index int) {
	*list = append((*list)[:index], (*list)[index+1:]...)
}

func searchVendorPrompt(list VendorList) int {

	searcher := func(input string, index int) bool {
		vendor := list[index]
		name := strings.Replace(strings.ToLower(vendor.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	promptForOperation := promptui.Select{
		Label:             "Search for Vendor",
		Items:             list,
		HideSelected:      true,
		Size:              10,
		Searcher:          searcher,
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Inactive: "{{ .Name }}",
			Active:   promptui.IconSelect + " {{ .Name | underline }}",
			Details: `
--------- Vendor Info ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "URL:" | faint }}	{{ .Url }}
{{ "Country:" | faint }}	{{ .Country }}
{{ "Region:" | faint }}	{{ .Region }}
{{ "Hidden:" | faint }}	{{ .Hide }}`,
		},
	}

	index, _, err := promptForOperation.Run()

	if err != nil {
		log.Fatal("prmo")
	}

	return index
}

func writeListToFile(path string, list VendorList) {
	// Opens and clears the file
	in, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		log.Fatal("error opening file", err)
	}
	defer in.Close()
	// Marshall the data
	json, _ := json.MarshalIndent(list, "", " ")

	//Write
	_, err = in.Write(json)
	if err != nil {
		log.Fatal("error writing file", err)
	}

}

func addPromptCoordinator(vendorList *VendorList) {
	newVendor := &Vendor{}
	sorted := isVendorListSorted(*vendorList)

	if !sorted {
		log.Fatal("List not sorted, can't add entries")
	}

	for {
		addVendorPromptSequence(newVendor)
		res, err := confirmationPrompt(*newVendor)

		if err != nil {
			return
		}

		if res {
			insertIntoSortedList(vendorList, *newVendor)
			fmt.Println("Vendor Added to List:", newVendor.Name)
			return
		}
	}
}

func insertIntoSortedList(list *VendorList, item Vendor) {
	insertIndex := 0

	for ; insertIndex < len(*list); insertIndex++ {
		if compareVendorNames((*list)[insertIndex], item) == 1 {
			break
		}
	}

	*list = append(*list, Vendor{})
	copy((*list)[insertIndex+1:], (*list)[insertIndex:])
	(*list)[insertIndex] = item
}

func verifyList(list VendorList) {
	if isVendorListSorted(list) {
		fmt.Println("Verified List: Sorted Correctly")
		return
	}
}

func isVendorListSorted(list VendorList) bool {
	for i := 1; i < len(list); i++ {
		if compareVendorNames(list[i-1], list[i]) != -1 {
			fmt.Println("Not sorted at:", list[i])
			return false
		}
	}
	return true
}

func compareVendorNames(a Vendor, b Vendor) int {
	return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
}

var errConfiramtionCancel = errors.New("user canceled action during the confirmation prompt")

func confirmationPrompt(changes Vendor) (bool, error) {

	prompt := promptui.Select{
		Label:        "Confirm Data Below",
		Items:        []string{"Looks Good", "Try Again", "Cancel"},
		HideSelected: true,
		Templates: &promptui.SelectTemplates{
			Details: fmt.Sprintf(`
------ New Vendor Info ------
{{ "Name:" | faint }}	%v
{{ "URL:" | faint }}	%v
{{ "Country:" | faint }}	%v
{{ "Region:" | faint }}	%v
{{ "Hidden:" | faint }}	%v`,
				changes.Name, changes.Url, changes.Country, changes.Region, changes.Hide)},
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal("internal prompt error:", err)
	}

	switch result {
	case "Looks Good":
		return true, nil
	case "Try Again":
		return false, nil
	case "Cancel":
		return false, errConfiramtionCancel
	default:
		log.Fatal("internal error, unknown prompt option")
	}

	return false, nil
}

func formatVendor(v Vendor) string {
	return fmt.Sprintf("{\n\tname: %v,\n\turl: %v,\n\tcountry: %v,\n\tregion: %v,\n\thide: %v,\n}", v.Name, v.Url, v.Country, v.Region, v.Hide)
}

func addVendorPromptSequence(newVendor *Vendor) {
	promptForStringField("Name", &newVendor.Name)
	promptForStringField("URL", &newVendor.Url)
	promptForStringField("Country", &newVendor.Country)
	selectRegion(&newVendor.Region)
}

func promptForStringField(fieldName string, fieldPtr *string) {
	prompt := promptui.Prompt{
		Label:       fmt.Sprintf("Enter Vendor %v", fieldName),
		Default:     *fieldPtr,
		AllowEdit:   true,
		HideEntered: true,
	}

	res, err := prompt.Run()

	if err != nil {
		log.Fatal("internal prompt error:", err)
	}

	*fieldPtr = res
}

func selectRegion(fieldPtr *string) {
	regions := []string{"asia", "africa", "canada", "china", "europe", "latin america", "ocenia", "uk", "us"}

	initPos := 0

	// Set initial postion of cursor
	if *fieldPtr != "" {
		for i, v := range regions {
			if v == *fieldPtr {
				initPos = i
			}
		}
	}

	prompt := promptui.Select{
		Label:        "Select Region",
		Items:        regions,
		Size:         9,
		CursorPos:    initPos,
		HideSelected: true,
	}
	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal("internal prompt error:", err)
	}

	*fieldPtr = res
}

func promptForNewListFile(pathString *string) {
	prompt := promptui.Select{
		Label: "The provided path is a directory. No vendorData.json found, would you like to create one?",
		Items: []string{"Yes", "No"},

		HideSelected: true,
	}
	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal("Prompt Failed, exiting.", err)
	}

	if result == "No" {
		log.Fatal("No vendorList file provided, exiting")
	}

	// Change path string to point to new file and create it
	*pathString = *pathString + "/vendorData.json"

	createVendorListFile(*pathString)
}

func createVendorListFile(filePath string) {
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal("Failed to create file:", err)
	}

	file.Write([]byte("[]"))
	file.Close()
}

var errNoListInDirectory = errors.New("provided path is a directory, no 'vendorData.json' found")

func verifyPath(pathString string) (string, error) {
	//Check if path exists
	info, err := os.Stat(pathString)

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Incorrect argument, path does not exist")
		}
	}

	//Check if path is a directory
	if info.IsDir() {
		//Check if list in directory
		filePath := filepath.Clean(pathString) + "/vendorData.json"
		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil
		} else {
			return pathString, errNoListInDirectory
		}
	}

	return pathString, nil
}

func parseArguments() string {
	if len(os.Args) == 1 {
		log.Fatal("No path provided, will exit now")
	}

	return os.Args[1]
}

func configureLogger() {
	log.SetFlags(0)
}
