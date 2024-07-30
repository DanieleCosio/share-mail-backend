dev:
	node build/esbuild.js
	postcss internal/app/components/css/app.css --output ./web/static/styles.css --env development --verbose