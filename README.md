# Go Weather API and UI Tests

This project contains Go tests for a weather API and a UI site.

## API Tests

The API tests verify the functionality of a weather API by making a GET request to `https://wttr.in/` and checking for a 200 OK status.

To run the API tests, navigate to the project root and execute:
```bash
go test ./weather/api
```

## UI Tests

The UI tests use `chromedp` to navigate to `https://wttr.in/` and check the page title.

To run the UI tests, navigate to the project root and execute:
```bash
go test ./weather/ui
```

**Note:** The UI test `TestWeatherUINavigation` currently fails because the expected title "wttr.in" does not match the actual title "Weather report: Batumi, Georgia". This is expected as `wttr.in` dynamically generates titles based on location.