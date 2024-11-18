package main

import (
	"github.com/charmbracelet/huh"
)

func main() {
	var base *huh.Theme = huh.ThemeCatppuccin()
	var file string
	var energy string
	var number_density string
	var wavelength string
	var sigma01 string
	var sigma12 string
	var sigma34 string
	var t10 string
	var t21 string
	var t13 string
	var t30 string
	var t43 string


	huh.NewForm(
		huh.NewGroup(
		
			huh.NewFilePicker().
				Title("Filename").
				Description("Select your file for fitting.").
				AllowedTypes([]string{".all"}).
				Value(&file),

			huh.NewInput().
				Title("Laser Energy").
				Placeholder("211.8e-9").
				Description("Enter laser energy in joules.").
				Value(&energy),

			huh.NewInput().
				Title("Molecule Density").
				Placeholder("2.08e17").
				Description("Enter Molecule Density in molecules/cm3.").
				Value(&number_density),

			huh.NewInput().
				Title("Laser Wavelength").
				Placeholder("532e-9").
				Description("Enter laser wavelength in meters.").
				Value(&wavelength),

			huh.NewInput().
				Title("Cross Section of S0-S1.").
				Placeholder("5.48e-18").
				Description("Enter Cross-Section of S0-S1 in cm2.").
				Value(&sigma01),

			huh.NewInput().
				Title("Cross Section of S1-S2.").
				Placeholder("5.48e-18").
				Description("Enter Cross-Section of S1-S2 in cm2.").
				Value(&sigma12),

			huh.NewInput().
				Title("Cross Section of T1-T2.").
				Placeholder("5.48e-18").
				Description("Enter Cross-Section of T1-T2 in cm2.").
				Value(&sigma34),

			huh.NewInput().
				Title("Lifetime of S1-S0.").
				Placeholder("1.0e-12").
				Description("Enter lifetime of S1-S0 in seconds").
				Value(&t10),

			huh.NewInput().
				Title("Lifetime of S2-S1.").
				Placeholder("1.0e-12").
				Description("Enter lifetime of S2-S1 in seconds").
				Value(&t21),
			
			huh.NewInput().
				Title("Lifetime of intersystem crossing.").
				Placeholder("1.0e-12").
				Description("Enter lifetime of intersystem crossing in seconds").
				Value(&t13),
			
			huh.NewInput().
				Title("Lifetime of T1-S0.").
				Placeholder("1.0e-12").
				Description("Enter lifetime of T1-S0 in seconds").
				Value(&t30),

			huh.NewInput().
				Title("Lifetime of T2-T1.").
				Placeholder("1.0e-12").
				Description("Enter lifetime of T2-T1 in seconds").
				Value(&t43),

		),
	).WithShowHelp(true).WithTheme(base).Run()
}
