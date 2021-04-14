package structs

// Вспомогательная функция, чтобы спрятать некрасивый код.
func GetTableTemplate() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8" />
    <style>
table {
font-family: "Lucida Sans Unicode", "Lucida Grande", Sans-Serif;
text-align: center;
border-collapse: separate;
border-spacing: 5px;
background: #ECE9E0;
color: #656665;
border: 16px solid #ECE9E0;
border-radius: 20px;
}
th {
font-size: 18px;
padding: 10px;
}
td {
background: #F5D7BF;
padding: 10px;
}
    </style>
</head>
<body>
<table style="width:80%"> <tr> <th>Название тарифа</th> <th>Стоимость</th> <th>Доступно мест</th> <th>Номер полета</th> <th>Авиакомпания</th></tr>`
}
