## Installation

Install micro-report with tidy

```bash
  go mod tidy
  create .env file
  put the fonts in the root/fonts
      root/fonts
      - ARIAL.TTF //normal
      - ARIALBD.TTF //BOLD
  
  make directories
        root/tmp
        root/tmp/pdf
        root/tmp/excel
  
  go run cmd/server/main.go
```
## Connect GRPC
### Requeriments
- Install [buf](https://docs.buf.build/installation)

### Generate pb

```bash
  cd protos
  buf lint
  buf generate
```

## API Reference

### Sales

#### Get pdf sales report

```http
  POST /v1.SalesService/RetrieveSalesResourceReport
```

| Parameter    | Type     | Description                                                                |
|:-------------|:---------|:---------------------------------------------------------------------------|
| `period`     | `string` | **Required**. Period for retrieve, eg: 2024-01                             |
| `businessId` | `string` | **Required**. Company for filter, eg: bf4336e4-b9b7-11ec-b4c3-00505605deef |
| `type`       | `number` | **Required**. 1 PDF, 2 XLSX                                                |

**return:** `string` path of file, use: http://localhost:8080/tmp/pdf/2024000.pdf
available for 5 minutes

#### Get excel sales report

```http
  POST /v1.SalesService/RetrieveSalesPaginatedReport
```

| Parameter    | Type     | Description                                                                |
|:-------------|:---------|:---------------------------------------------------------------------------|
| `period`     | `string` | **Required**. Period for retrieve, eg: 2024-01                             |
| `businessId` | `string` | **Required**. Company for filter, eg: bf4336e4-b9b7-11ec-b4c3-00505605deef |
| `page`       | `int`    | **Required**. Page number, eg: 1162                                        |
| `pageSize`   | `int`    | **Required**. Page size, eg: 30                                            |

**return:** proto []SalesReport in json
