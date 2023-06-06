package main

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func cncpts() {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("https://google.com")
	assertErrorToNilf("could not goto: %w", err)
	time.Sleep(2 * time.Second)

	_, err = page.Goto("https://cncpts.com/collections/mens-sneakers")
	assertErrorToNilf("could not goto: %w", err)


	//page.Click("text=" + "MENS")
	time.Sleep(2 * time.Second)
	page.Click("text=" + "NIKE SB")
	time.Sleep(2 * time.Second)
	labelSize := []string{"9"}
	selectSize,_ := page.QuerySelector(`select[data-name="Size"]`)
	selectSize.SelectOption(playwright.SelectOptionValues{Labels: &labelSize})
	page.Click(`span[class="btn__text"]`)
 	time.Sleep(2 * time.Second)
 	
	page.Click("text=" + "CART (1)")
	time.Sleep(2 * time.Second)

	page.Click("text=" + "CONTINUE TO")
	time.Sleep(2 * time.Second)
	
	page.Click(`input[name="add"]`)
 	time.Sleep(2 * time.Second)
	
	assertErrorToNilf("could not goto: %w", err)
	
	time.Sleep(2 * time.Second)
	// labelCountry := []string{}
	// selectCountry,_ := page.QuerySelector(`select[name="countryCode"]`)
	// selectCountry.SelectOption(playwright.SelectOptionValues{Labels: &labelCountry})

	
	// labelCounty := []string{}
	
	// selectCounty,_ := page.QuerySelector(`select[name="zone"]`)
	// selectCounty.SelectOption(playwright.SelectOptionValues{Labels: &labelCounty})
	
	
	// page.Fill(`input[name="firstName"]`, "")
 	// page.Fill(`input[name="lastName"]`, "")
 	// page.Fill(`input[name="email"]`, "")
 	// page.Fill(`input[name="address1"]`, "")
 	// page.Fill(`input[name="address2"]`, "")
 	// page.Fill(`input[name="postalCode"]`, "")
 	// page.Fill(`input[name="city"]`, "")
 	// page.Fill(`input[name="phone"]`, "")
	
 	// page.Click("text=" + "card number")
	// page.Keyboard().InsertText("")
	// page.Click("text=" + "name on card")
 	// page.Keyboard().InsertText("")
	// page.Click("text=" + "expiration date")
 	// page.Keyboard().InsertText("")
	// page.Click("text=" + "security code")
 	// page.Keyboard().InsertText("")
	// page.Click(`input[id="accept_tos"]`) 
	// time.Sleep(10 * time.Second)
	// context.Close()
	// browser.Close()
}


func size() {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("https://www.size.co.uk/campaign/New+In/?facet:new=latest&sort=latest")
	assertErrorToNilf("could not goto: %w", err)
	time.Sleep(2 * time.Second)
	
	page.Click(`button[class="btn btn-level1 accept-all-cookies"]`)
	time.Sleep(2 * time.Second)


	time.Sleep(2 * time.Second)
	page.Click("text=" + "Nike Air Force 1 Low '07")
	time.Sleep(2 * time.Second)
	//labelSize := []string{"9.5"}
	//selectSize,_ := page.QuerySelector(`select[data-name="Size"]`)
	//selectSize.SelectOption(playwright.SelectOptionValues{Labels: &labelSize})
 	time.Sleep(1 * time.Second)
	page.Click(`button[title="Select Your UK Size 9.5"]`)
 
	//page.Click("text="+" L ")
	//time.Sleep(3 * time.Second)

	page.Click("text=" + "Add to Bag")
	time.Sleep(2 * time.Second)

	page.Click("text=" + "Checkout securely")
	time.Sleep(2 * time.Second)

	page.Click("text=" + " Checkout securely ")
	time.Sleep(4 * time.Second)




	//page.Click(`input[name="add"]`)
 	//time.Sleep(2 * time.Second)
	
	assertErrorToNilf("could not goto: %w", err)
	
	time.Sleep(2 * time.Second)
	// labelCountry := []string{}
	// selectCountry,_ := page.QuerySelector(`select[name="countryCode"]`)
	// selectCountry.SelectOption(playwright.SelectOptionValues{Labels: &labelCountry})

	
	// labelCounty := []string{}
	
	// selectCounty,_ := page.QuerySelector(`select[name="zone"]`)
	// selectCounty.SelectOption(playwright.SelectOptionValues{Labels: &labelCounty})
	
	
	// page.Fill(`input[name="firstName"]`, "")
 	// page.Fill(`input[name="lastName"]`, "")
 	// page.Fill(`input[name="email"]`, "")
 	// page.Fill(`input[name="address1"]`, "")
 	// page.Fill(`input[name="address2"]`, "")
 	// page.Fill(`input[name="postalCode"]`, "")
 	// page.Fill(`input[name="city"]`, "")
 	// page.Fill(`input[name="phone"]`, "")
	
 	// page.Click("text=" + "card number")
	// page.Keyboard().InsertText("")
	// page.Click("text=" + "name on card")
 	// page.Keyboard().InsertText("")
	// page.Click("text=" + "expiration date")
 	// page.Keyboard().InsertText("")
	// page.Click("text=" + "security code")
 	// page.Keyboard().InsertText("")
	// page.Click(`input[id="accept_tos"]`) 
	// time.Sleep(10 * time.Second)
	// context.Close()
	// browser.Close()
}

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func undefeated() {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("https://undefeated.com/collections/all")
	assertErrorToNilf("could not goto: %w", err)
	page.Click("text=" + "Continue")
	time.Sleep(2 * time.Second)
	page.Click(`img[title='JORDAN FLIGHT FLEECE PULLOVER HOODIE']`)
	time.Sleep(2 * time.Second)
	labelSize := []string{"L"}
	selectSize,_ := page.QuerySelector(`select[data-name="Size"]`)
	selectSize.SelectOption(playwright.SelectOptionValues{Labels: &labelSize})
	page.Click(`span[class="btn__text"]`)
 	time.Sleep(2 * time.Second)
 	
	page.Click("text=" + "CART (1)")
	time.Sleep(2 * time.Second)

	page.Click("text=" + "CONTINUE TO")
	time.Sleep(2 * time.Second)
	
	page.Click(`input[name="add"]`)
 	time.Sleep(2 * time.Second)
	
	assertErrorToNilf("could not goto: %w", err)
	
	time.Sleep(2 * time.Second)
	// labelCountry := []string{}
	// selectCountry,_ := page.QuerySelector(`select[name="countryCode"]`)
	// selectCountry.SelectOption(playwright.SelectOptionValues{Labels: &labelCountry})

	
	// labelCounty := []string{}
	
	// selectCounty,_ := page.QuerySelector(`select[name="zone"]`)
	// selectCounty.SelectOption(playwright.SelectOptionValues{Labels: &labelCounty})
	
	
	// page.Fill(`input[name="firstName"]`, "")
 	// page.Fill(`input[name="lastName"]`, "")
 	// page.Fill(`input[name="email"]`, "")
 	// page.Fill(`input[name="address1"]`, "")
 	// page.Fill(`input[name="address2"]`, "")
 	// page.Fill(`input[name="postalCode"]`, "")
 	// page.Fill(`input[name="city"]`, "")
 	// page.Fill(`input[name="phone"]`, "")
	
 	// page.Click("text=" + "card number")
	// page.Keyboard().InsertText("")
	// page.Click("text=" + "name on card")
 	// page.Keyboard().InsertText("")
	// page.Click("text=" + "expiration date")
 	// page.Keyboard().InsertText("")
	// page.Click("text=" + "security code")
 	// page.Keyboard().InsertText("")
	// page.Click(`input[id="accept_tos"]`) 
	// time.Sleep(10 * time.Second)
	// context.Close()
	// browser.Close()
}


func palace(Task ReadyTask) {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("https://shop-eu.palaceskateboards.com")
	assertErrorToNilf("could not goto: %w", err)
	time.Sleep(2 * time.Second)
	page.Click("text=" + Task.Task.Category)

	time.Sleep(2 * time.Second)
	page.Click("text=" + Task.Task.ProductName)

	time.Sleep(2 * time.Second)
	
	if Task.Task.Size != "One Size"{
		labelSize := []string{Task.Task.Size}
		selectSize,_ := page.QuerySelector(`select[id=product-select]`)
		selectSize.SelectOption(playwright.SelectOptionValues{Labels: &labelSize})
	}
	
	time.Sleep(2 * time.Second)
	page.Click("text="+"ADD TO CART")
	
	time.Sleep(2 * time.Second)
	page.Click("text="+"CART")

	
	time.Sleep(2 * time.Second)
	page.Click(`div[class="checkbox-control"]`)
	time.Sleep(2 * time.Second)
	page.Click("text="+"CHECKOUT")
	time.Sleep(2 * time.Second)
	page.Fill(`input[name="checkout[email]"]`,Task.Shipment.Email)
	page.Fill(`input[name="checkout[shipping_address][country]"]`,Task.Shipment.Country)
	page.Fill(`input[name="checkout[shipping_address][first_name]"]`,Task.Shipment.Name)
	page.Fill(`input[name="checkout[shipping_address][last_name]"]`,Task.Shipment.Surname)
	page.Fill(`input[name="checkout[shipping_address][address1]"]`,Task.Shipment.Address)
	page.Fill(`input[name="checkout[shipping_address][address2]"]`,Task.Shipment.Address2)
	page.Fill(`input[name="checkout[shipping_address][zip]"]`, Task.Shipment.Zip)
	page.Fill(`input[name="checkout[shipping_address][city]"]`, Task.Shipment.City)
	page.Fill(`input[name="checkout[shipping_address][province]"]`, Task.Shipment.County)
	page.Fill(`input[name="checkout[shipping_address][phone]"]`,Task.Shipment.Phone)
	time.Sleep(2 * time.Second)
	page.Click("text="+"Continue to shipping")
	time.Sleep(2 * time.Second)
	page.Click("text="+"Continue to payment")
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Card number")
	page.Keyboard().InsertText(Task.Card.CardNumber)
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Name on card")
 	page.Keyboard().InsertText(Task.Card.CardOwner)
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Expiration date")
 	page.Keyboard().InsertText(Task.Card.ExpirationDate)
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Security code")
 	page.Keyboard().InsertText(Task.Card.Cvv)
	time.Sleep(5 * time.Second)
	context.Close()
	browser.Close()
}

func supreme(Task ReadyTask) {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("https://www.supremenewyork.com/shop/all")
	assertErrorToNilf("could not goto: %w", err)

	page.Click("text=" + Task.Task.Category)
	time.Sleep(2 * time.Second)
 	page.Click("text=" + Task.Task.ProductName + " " + Task.Task.Color)
 	time.Sleep(2 * time.Second)

 	time.Sleep(2 * time.Second)
	if Task.Task.Size != "One Size"{
		labelSize := []string{Task.Task.Size}
		selectSize,_ := page.QuerySelector(`select[data-cy="size-selector"]`)
		selectSize.SelectOption(playwright.SelectOptionValues{Labels: &labelSize})
	}
 	
	
	page.Click(`input[name="add"]`)
 	time.Sleep(2 * time.Second)
	//url := page.URL()
	_,err = page.Goto("https://eu.supreme.com/checkout")
	
	assertErrorToNilf("could not goto: %w", err)

	labelCountry := []string{Task.Shipment.Country}
	selectCountry,_ := page.QuerySelector(`select[name="countryCode"]`)
	selectCountry.SelectOption(playwright.SelectOptionValues{Labels: &labelCountry})

	
	labelCounty := []string{Task.Shipment.County}
	fmt.Println(Task.Shipment.County)
	selectCounty,_ := page.QuerySelector(`select[name="zone"]`)
	selectCounty.SelectOption(playwright.SelectOptionValues{Labels: &labelCounty})
	
	
	page.Fill(`input[name="firstName"]`, Task.Shipment.Name)
 	page.Fill(`input[name="lastName"]`, Task.Shipment.Surname)
 	page.Fill(`input[name="email"]`, Task.Shipment.Email)
 	page.Fill(`input[name="address1"]`, Task.Shipment.Address)
 	page.Fill(`input[name="address2"]`, Task.Shipment.Address2)
 	page.Fill(`input[name="postalCode"]`, Task.Shipment.Zip)
 	page.Fill(`input[name="city"]`, Task.Shipment.City)
 	page.Fill(`input[name="phone"]`, Task.Shipment.Phone)
	
 	page.Click("text=" + "card number")
	page.Keyboard().InsertText(Task.Card.CardNumber)
	page.Click("text=" + "name on card")
 	page.Keyboard().InsertText(Task.Card.CardOwner)
	page.Click("text=" + "expiration date")
 	page.Keyboard().InsertText(Task.Card.ExpirationDate)
	page.Click("text=" + "security code")
 	page.Keyboard().InsertText(Task.Card.CardNumber)
	page.Click(`input[id="accept_tos"]`) 
	time.Sleep(10 * time.Second)
	context.Close()
	browser.Close()
}