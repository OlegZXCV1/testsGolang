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

## AI Tests

The AI tests fetch weather data from `wttr.in` and then use the Gemini API to generate a haiku about the weather.

To run the AI tests, you need to set the `GEMINI_API_KEY` environment variable with your Gemini API key (obtainable from Google AI Studio).

Once the environment variable is set, navigate to the project root and execute:
```bash
export GEMINI_API_KEY="YOUR_API_KEY"
go test ./weather/ai
```

## Multi-Modal UI AI Tests

The multi-modal UI AI tests use `chromedp` to take a screenshot of `https://wttr.in/` and then use the Gemini API to generate a description of the image.

To run the multi-modal UI AI tests, you need to set the `GEMINI_API_KEY` environment variable with your Gemini API key (obtainable from Google AI Studio).

Once the environment variable is set, navigate to the project root and execute:
```bash
export GEMINI_API_KEY="YOUR_API_KEY"
go test ./weather/ui_mcp
```