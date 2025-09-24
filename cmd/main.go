package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	gravatar_recon "github.com/anotherhadi/gravatar-recon"
	"github.com/anotherhadi/gravatar-recon/utils"
	"github.com/charmbracelet/log"

	flag "github.com/spf13/pflag"
)

func isEmailValid(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "gravatar-recon [flags] <target email>\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	flag.CommandLine.SortFlags = false

	silent := flag.BoolP("silent", "s", false, "Suppress all non-essential output")
	printAvatar := flag.BoolP("print-avatar", "a", true, "Show the avatar in the output")
	jsonOutput := flag.StringP("json", "j", "", "Write results to specified JSON file")

	flag.Parse()

	// TAIL

	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 1 {
		log.Error("Please provide only one target (email)")
		flag.Usage()
		os.Exit(1)
	} else if len(nonFlagArgs) == 0 {
		log.Error("Please provide a target (email)")
		flag.Usage()
		os.Exit(1)
	}

	email := strings.TrimSpace(flag.Arg(0))

	if !isEmailValid(email) {
		log.Fatal("Invalid email address", "email", email)
		return
	}

	profiles, err := gravatar_recon.GetGravatarProfiles(email)
	if err != nil {
		log.Fatal("Failed to get Gravatar profile", "err", err)
	}

	if !*silent {
		utils.Header()
		for _, profile := range *profiles {
			utils.PrintTitle(profile.PreferredUser + ":")
			if *printAvatar {
				utils.PrintAvatar(profile.ThumbnailURL)
			}
			utils.PrintStruct(profile, 0)
		}
	}
	writeJson(*jsonOutput, profiles)
}

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{".", "_"}
	to := "-"
	for _, sep := range from {
		name = strings.ReplaceAll(name, sep, to)
	}
	return flag.NormalizedName(name)
}
