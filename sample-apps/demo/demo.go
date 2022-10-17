package main

import (
	"log"

	goandroid "github.com/supaplextor/adbGoAutomation"
)

func main() {
	// Create a new android manager with 60 seconds adb time out and
	// take adb executable path from system path.
	android_manager := goandroid.GetNewAndroidManager(60, "adb")
	android_manager.Endpoint.Debug = true

	// Use the sn from '$ adb devices' on the shell.
	// Create an android device instance with following serial
	android := android_manager.GetNewAndroidDevice("R9PN711W1ZJ")

	x, y, err := android.Display.GetDisplaySize()
	if nil != err {
		log.Printf("android.Display.GetDisplaySize() = %d %d %v\n", x, y, err)
	}
	log.Printf("x=%d y=%d\n", x, y)

	// Start settings activity
	err = android.Activity.StartActivity("com.android.settings")
	if nil != err {
		log.Printf("android.Activity.StartActivity(...) = %v\n", err)
	}
	// Wait for settings activity to get focused and displayed on screen
	// with 10 seconds timeout
	err = android.Activity.WaitForActivityToFocus("com.android.settings", 10)
	if nil != err {
		log.Panicf("android.Activity.WaitForActivityToFocus(...) = %v\n", err)
	}

	// Scroll down to "About phone"
	err = android.View.ScrollDownToText("About phone", 0, 10)
	if nil != err {
		log.Panicf("android.View.ScrollDownToText(...) = %v\n", err)
	}
	// Now click "About phone" settings item
	err = android.View.ClickText("About phone", 0, 5)
	if nil != err {
		log.Panicf("android.View.ClickText(...) = %v\n", err)
	}

	// Now scroll down to "Build number"
	err = android.View.ScrollDownToText("Build number", 0, 10)
	if nil != err {
		log.Panicf("android.View.ScrollDownToText(...) = %v\n", err)
	}

	// Now for faster click operation, we are going to get the view for "Build number" text
	view, err := android.View.GetViewForText("Build number", 0, 15)
	if nil != err {
		log.Panicf("android.View.GetViewForText(...) = %v\n", err)
	}

	// Now we will click the text 10 times
	for i := 0; i < 10; i++ {
		err = android.Input.TouchScreen.Tap(view.Center.X, view.Center.Y)
		if nil != err {
			log.Printf("android.Input.TouchScreen.Tap(x=%d,y=%d)) = %v\n", view.Center.X, view.Center.Y, err)
		}
	}

	// Now go back to main settings page
	err = android.Input.Key.PressBack(1)
	if nil != err {
		log.Panicf("android.Input.Key.PressBack(...) = %v\n", err)
	}
	// Click developer options
	err = android.View.ClickText("Developer options", 0, 10)
	if nil != err {
		log.Panicf("android.View.ClickText(...) = %v\n", err)
	}

	// Now scroll down to "Show CPU Usage" and enable it
	err = android.View.ScrollDownToMatchingText("show cpu", 0, 10)
	if nil != err {
		log.Panicf("android.View.ScrollDownToMatchingText(...) = %v\n", err)
	}
	err = android.View.ClickMatchingText("show cpu", 0, 10)
	if nil != err {
		log.Panicf("android.View.ClickMatchingText(...) = %v\n", err)
	}
}
