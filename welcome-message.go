package main

import (
	"fmt"

	"github.com/fatih/color"
)

func welcomeMessage(CONFIG_PATH string) {
	fmt.Fprint(color.Output, color.CyanString(`
  __          ________ _      _____ ____  __  __ ______  
  \ \        / /  ____| |    / ____/ __ \|  \/  |  ____| 
   \ \  /\  / /| |__  | |   | |   | |  | | \  / | |__    
    \ \/  \/ / |  __| | |   | |   | |  | | |\/| |  __|   
     \  /\  /  | |____| |___| |___| |__| | |  | | |____  
      \/  \/   |______|______\_____\____/|_|  |_|______|       
`))

	fmt.Fprintf(
		color.Output,
		"\n\nThis is the "+color.GreenString("Auto Class launcher")+" project! Opens up your class links based on your timetable 5 minutes before\n",
	)

	fmt.Println("\nThis project works based on a config file stored in your computer. Your timetable and links are there only. ")

	fmt.Fprintf(color.Output, "Your config file is stored at "+color.GreenString(CONFIG_PATH)+"\n")

	fmt.Fprint(color.Output, color.YellowString("\nPlease modify the file for your own purposes.\n"))

	fmt.Fprintf(color.Output, "\nTo read about how to modify the file and its format, go to "+color.RedString("https://github.com/PuruVJ/auto-class-launcher\n"))

}
