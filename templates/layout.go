package templates

const (
	ErrorPageLayout = `
	<!DOCTYPE html>
	<html lang="en">

	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{ .Title }}</title>
		<script src="https://cdn.tailwindcss.com"></script>
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@shadcn/ui@0.1.1/dist/shad.css">
	</head>

	<body class="bg-gray-100 flex items-center justify-center min-h-screen">
		<div class="container text-center p-6 md:p-12">
			<div class="bg-white shadow-md rounded-lg p-8 max-w-lg mx-auto">
				<h1 class="text-4xl md:text-5xl font-bold text-red-600 mb-6">404 Error</h1>
				<p class="text-gray-600 text-lg mb-6">{{ .Message }}</p>
				<p class="text-gray-400 mb-5">RequestID: {{ .TraceId }}</p>
			</div>
		</div>
	</body>

	</html>`
)
