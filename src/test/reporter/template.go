/*
 * SPACE-AICOS is pleased to support the open source community by making
 * SPACE-AICOS available.
 * Copyright (C) 2026 China Academy of Information and Communications Technology (CAICT).
 * All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package reporter

const (
	// HtmlTemplate TODO
	HtmlTemplate = `
<html>
	<head>
		<title>{{.Name}}</title>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
		<style type="text/css">
		/* gridtable */
			table.gridtable {
				font-family: verdana,arial,sans-serif;
				font-size:12px;
				color:#333333;
				border-width: 1px;
				border-color: #666666;
				border-collapse: collapse;
			}
			table.gridtable th {
				border-width: 1px;
				padding: 8px;
				border-style: solid;
				border-color: #666666;
				background-color: #dedede;
			}
			table.gridtable td {
				border-width: 1px;
				padding: 8px;
				border-style: solid;
				border-color: #666666;
				background-color: #ffffff;
			}
		/* /gridtable */
		</style>
	</head>
	<body>
		<div>
			<h1>{{.Name}}</h1>
			<p><strong>RunTime: </strong>{{.RunTime}}s</p>
			<p><strong>Total Tests: </strong>{{.TotalNum}}</p>
			<p><strong>Failed Tests: </strong>{{.FailedNum}}</p>
		</div>
		<table class="gridtable">
			<tr>
				<th>Name</th>
				<th>State</th>
				<th>RunTime</th>
				<th>Detail</th>
			</tr>
			{{ range .FailedTestCases }}
			<tr style='color:red;'>
				<td>{{.Name}}</td>
				<td>{{.State}}</td>
				<td>{{.RunTime}}s</td>
				<td>{{.Detail}}</td>
			</tr>
			{{ end }}
			{{ range .OtherTestCases }}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.State}}</td>
				<td>{{.RunTime}}s</td>
				<td>{{.Detail}}</td>
			</tr>
			{{ end }}
		</table>
	</body>
</html>
`
	// SummaryHtmlTemplate TODO
	SummaryHtmlTemplate = `
<html>
<head>
    <title>Summary of Test Results</title>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
	<style type="text/css">
		/* gridtable */
		table.gridtable {
			font-family: verdana,arial,sans-serif;
			font-size:12px;
			color:#333333;
			border-width: 1px;
			border-color: #666666;
			border-collapse: collapse;
		}
		table.gridtable th {
			border-width: 1px;
			padding: 8px;
			border-style: solid;
			border-color: #666666;
			background-color: #dedede;
		}
		table.gridtable td {
			border-width: 1px;
			padding: 8px;
			border-style: solid;
			border-color: #666666;
			background-color: #ffffff;
		}
	/* /gridtable */
	</style>
</head>
<body>
    <h2>Summary of Test Results</h2>
	<table class="gridtable">
		<tr>
			<th>Test Name</th>
			<th>Test Status</th>
			<th>Total Check Num</th>
			<th>Successful Check Num</th>
			<th>Failed Check Num</th>
			<th>Html Url</th>
		</tr>
	</table>
</body>
</html>
`

	// SummaryTemplate TODO
	SummaryTemplate = `
		<tr>
			<td>{{.Name}}</td>
			<td>{{.State}}</td>
			<td>{{.TotalNum}}</td>
			<td>{{.SuccessNum}}</td>
			<td>{{ if ne .State "Passed" }}<a href="#{{.Name}}">{{ end }}{{.FailedNum}}{{ if ne .State "Passed" }}</a>{{ end }}</td>
			<td><a href='{{.Url}}'>html</a></td>
		</tr>
`

	// FailedTemplate TODO
	FailedTemplate = `
	{{ if ne .State "Passed" }}
	<div id="{{.Name}}">
		<h3>{{.Name}}</h3>
		<table class="gridtable">
			<tr>
				<th>Name</th>
				<th>State</th>
				<th>RunTime</th>
				<th>Detail</th>
			</tr>
			{{ range .FailedTestCases }}
			<tr style='color:red;'>
				<td>{{.Name}}</td>
				<td>{{.State}}</td>
				<td>{{.RunTime}}s</td>
				<td>{{.Detail}}</td>
			</tr>
			{{ end }}
		</table>
	</div>
	{{ end }}
`
)
