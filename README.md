
## Installation

Install micro-report with tidy

```bash
  go mod tidy
  go run cmd/server/main.go
```

## API Reference

### Sales

#### Get all Sales

```http
  GET /api/sales-excel?companyID=bf4336e4-b9b7-11ec-b4c3-00505605deef&period=2024-01
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `period` | `string` | **Required**. Period for retrieve, eg: 2024-01|
| `companyID` | `string` | **Required**. Company for filter, eg: bf4336e4-b9b7-11ec-b4c3-00505605deef|

**return:** binary stream excel 