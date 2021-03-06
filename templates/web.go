package templates

const (
	Web = `<!DOCTYPE html>
<html>
	<head>
		<link href="{{.web_host}}" rel="canonical">
		<base href="{{.web_host}}">
		<meta http-equiv="content-type" content="text/html; charset=utf-8">
		<meta name="keywords" content="micro, microgen, pthethanh, thanh pham">
		<meta name="author" content="Thanh Pham">
		<meta property="og:url" content="{{.web_host}}">
		<meta property="og:title" content="microgen">
		<meta property="og:type" content="article">
		<meta property="og:image" content="https://blog.golang.org/lib/godoc/images/go-logo-blue.svg">
		<meta name="image" content="https://blog.golang.org/lib/godoc/images/go-logo-blue.svg">
		<meta property="og:description" content="micro: simple tool kit for building microservices">
		<meta name="description" content="micro: simple tool kit for building microservices">
		
		<title>micro: simple tool kit for building microservices</title>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
		<!-- CSS -->
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css" integrity="sha384-9aIt2nRpC12Uk9gS9baDl411NQApFmC26EwAOH8WgZl5MYYxFfc+NcPb1dKGj7Sk" crossorigin="anonymous"/>
		<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Libre+Baskerville:400,400i,700" />
		<link rel="stylesheet" href="static/css/main.css" />
		<!-- Favicon -->
		<link rel="icon" type="image/png" sizes="32x32" href="static/img/favicon.ico" />
		<link rel="icon" type="image/png" sizes="16x16" href="static/img/favicon.ico" />
		<link rel="apple-touch-icon" sizes="180x180" href="static/img/favicon.ico" />
	</head>
	<body>
	<div class="container-fluid">
		<div>
			<h1>micro</h1>
			<h3>Just a simple tool kit for building microservices</h3>
			<div><span>Github: <a href="https://github.com/pthethanh/micro">https://github.com/pthethanh/micro</a></span></div>
		</div>
	</div>
	<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
	<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/js/bootstrap.min.js" integrity="sha384-OgVRvuATP1z7JjHLkuOU7Xw704+h835Lr+6QL9UvYjZE3Ipu6Tp75j7Bh/kR0JKI" crossorigin="anonymous"></script>
	</body>
</html>
	`
)
