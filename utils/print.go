package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
	gopixels "github.com/saran13raj/go-pixels"
)

var (
	grey  = lipgloss.Color("#7d7d7d")
	green = lipgloss.Color("#a6e3a1")
	blue  = lipgloss.Color("#7287fd")

	greyStyle  = lipgloss.NewStyle().Foreground(grey)
	greenStyle = lipgloss.NewStyle().Foreground(green)
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(blue)
)

func PrintStruct(s any, indent int) {
	prefix := strings.Repeat("  ", indent)

	v := reflect.ValueOf(s)
	if !v.IsValid() {
		return
	}
	t := reflect.TypeOf(s)

	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return
		}
		v = v.Elem()
		t = v.Type()
	}

	switch v.Kind() {
	case reflect.Struct:
		if v.NumField() == 0 {
			fmt.Println(prefix + greyStyle.Render("No data found"))
			fmt.Println("")
			return
		}

		printed := 0
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i).Name
			value := v.Field(i)
			if !value.CanInterface() {
				continue
			}
			if !value.IsValid() || (value.Kind() == reflect.String && value.String() == "") {
				continue
			}
			if value.Kind() == reflect.String && value.String() == "0001-01-01 00:00:00 +0000 UTC" {
				continue
			}
			printed++

			switch value.Kind() {
			case reflect.Struct, reflect.Slice, reflect.Array, reflect.Ptr, reflect.Map, reflect.Interface:
				fmt.Println(prefix + greyStyle.Render(field+":"))
				PrintStruct(value.Interface(), indent+1)

			case reflect.String:
				fmt.Printf("%s%s %s\n", prefix, greyStyle.Render(field+":"), greenStyle.Render(fmt.Sprintf("%q", value.Interface())))

			default:
				fmt.Printf("%s%s %s\n", prefix, greyStyle.Render(field+":"), greenStyle.Render(fmt.Sprintf("%v", value.Interface())))
			}
		}
		if printed == 0 {
			fmt.Println(prefix + greyStyle.Render("No data found"))
		}
		fmt.Println("")

	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {
			fmt.Println(prefix + greyStyle.Render("No data found"))
			fmt.Println("")
			return
		}
		for i := 0; i < v.Len(); i++ {
			PrintStruct(v.Index(i).Interface(), indent)
		}

	case reflect.Map:
		if v.Len() == 0 {
			fmt.Println(prefix + greyStyle.Render("No data found"))
			return
		}

		keys := v.MapKeys()
		keyStrs := make([]string, len(keys))
		for i, k := range keys {
			keyStrs[i] = fmt.Sprintf("%v", k.Interface())
		}
		sort.Strings(keyStrs)

		for _, keyStr := range keyStrs {
			for _, k := range keys {
				if fmt.Sprintf("%v", k.Interface()) == keyStr {
					val := v.MapIndex(k)
					fmt.Println(prefix + greyStyle.Render(fmt.Sprintf("%v:", k.Interface())))
					PrintStruct(val.Interface(), indent+1)
				}
			}
		}

	default:
		fmt.Println(prefix + greenStyle.Render(fmt.Sprintf("%v", v.Interface())))
	}
}

func Header() {
	asciiArt := " _____                 _                                   \n|   __|___ ___ _ _ ___| |_ ___ ___ ___ ___ ___ ___ ___ ___ \n|  |  |  _| .'| | | .'|  _| .'|  _|___|  _| -_|  _| . |   |\n|_____|_| |__,|\\_/|__,|_| |__,|_|     |_| |___|___|___|_|_|"

	grey := lipgloss.Color("#7d7d7d")

	greyStyle := lipgloss.NewStyle().Foreground(grey)
	fmt.Println(
		greyStyle.Render(lipgloss.JoinVertical(lipgloss.Right, asciiArt, "@anotherhadi\n")),
	)
}

func PrintTitle(title string) {
	fmt.Println(titleStyle.Render(title) + "\n")
}

func PrintAvatar(url string) {
	if url == "" {
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	tmpfile, err := os.CreateTemp("", "avatar-*.png")
	if err != nil {
		return
	}
	defer os.Remove(tmpfile.Name())

	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return
	}

	output, err := gopixels.FromImagePath(tmpfile.Name(), 30, 25, "halfcell", true)

	if err != nil {
		return
	}
	fmt.Println(output + "\n")
}
