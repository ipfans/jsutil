package tailwind

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// TailwindElementsCSS is the tailwind elements css files.
var (
	TailwindElementsCSS = []string{
		"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css",
		"https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap",
		"https://cdn.jsdelivr.net/npm/tw-elements/dist/css/index.min.css",
	}
	TailwindElementsJS     = app.Script().Src("https://cdn.jsdelivr.net/npm/tw-elements/dist/js/index.min.js").Async(true)
	TailwindScript         = `<script src="https://cdn.tailwindcss.com"></script>`
	TailwindElementsConfig = `<script>
tailwind.config = {
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'sans-serif'],
      },
    }
  }
}
</script>`
)
