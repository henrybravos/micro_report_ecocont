## Installation

Install micro-report with tidy

```bash
  go mod tidy
  go run cmd/server/main.go
  create .env file
  put fonts
      root/fonts
      - ARIAL.TTF //normal
      - ARIALBD.TTF //BOLD
```

## API Reference

### Sales

#### Get pdf sales report

```http
  GET /api/sales-pdf?businessID=bf4336e4-b9b7-11ec-b4c3-00505605deef&period=2024-01
```

| Parameter    | Type     | Description                                                                |
|:-------------|:---------|:---------------------------------------------------------------------------|
| `period`     | `string` | **Required**. Period for retrieve, eg: 2024-01                             |
| `businessID` | `string` | **Required**. Company for filter, eg: bf4336e4-b9b7-11ec-b4c3-00505605deef |

**return:** binary stream pdf

#### Get excel sales report

```http
  GET /api/sales-excel?businessID=bf4336e4-b9b7-11ec-b4c3-00505605deef&period=2024-01
```

| Parameter    | Type     | Description                                                                |
|:-------------|:---------|:---------------------------------------------------------------------------|
| `period`     | `string` | **Required**. Period for retrieve, eg: 2024-01                             |
| `businessID` | `string` | **Required**. Company for filter, eg: bf4336e4-b9b7-11ec-b4c3-00505605deef |

**return:** binary stream excel

#### Get sales report by pagination

```http
  GET /api/sales-paginated?businessID=bf4336e4-b9b7-11ec-b4c3-00505605deef&period=2024-01&page=1162&pageSize=30
```

| Parameter    | Type     | Description                                                                |
|:-------------|:---------|:---------------------------------------------------------------------------|
| `period`     | `string` | **Required**. Period for retrieve, eg: 2024-01                             |
| `businessID` | `string` | **Required**. Company for filter, eg: bf4336e4-b9b7-11ec-b4c3-00505605deef |
| `page`       | `int`    | **Required**. Page number, eg: 1162                                        |
| `pageSize`   | `int`    | **Required**. Page size, eg: 30                                            |

**return:** array of Sales Report object with pagination metadata
