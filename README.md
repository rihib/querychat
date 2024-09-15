# QueryChat

With QueryChat, you can chat and ask what you want to know about a database and it will show you the data in its proper visualized form. Unlike BI tools like Looker or redash, there is no need to write SQL. You can use natural language to get information about your data.

## Example

If you ask "What are the monthly sales for 2013?", QueryChat will return following visualizable data:

```json
{
  "chart": {
    "type": "line",
    "x": "SaleMonth",
    "y": "TotalSales"
  },
  "data": {
    [
      {"SaleMonth": "2013-01", "TotalSales": 37.62},
      {"SaleMonth": "2013-02", "TotalSales": 27.72},
      {"SaleMonth": "2013-03", "TotalSales": 37.62},
      {"SaleMonth": "2013-04", "TotalSales": 33.66},
      {"SaleMonth": "2013-05", "TotalSales": 37.62},
      {"SaleMonth": "2013-06", "TotalSales": 37.62},
      {"SaleMonth": "2013-07", "TotalSales": 37.62},
      {"SaleMonth": "2013-08", "TotalSales": 37.62},
      {"SaleMonth": "2013-09", "TotalSales": 37.62},
      {"SaleMonth": "2013-10", "TotalSales": 37.62},
      {"SaleMonth": "2013-11", "TotalSales": 49.62},
      {"SaleMonth": "2013-12", "TotalSales": 38.62}
    ]
  }
}
```
