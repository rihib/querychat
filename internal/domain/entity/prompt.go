package entity

const (
	SYSTEM_PROMPT = "Given this database schema:\n" +
		"\n" +
		"```sql\n" +
		"%s\n" +
		"```\n" +
		"\n"

	USER_PROMPT = "Based on the provided database schema, please perform the following tasks:\n" +
		"\n" +
		"1. Craft a %s query that answers '%s'.\n" +
		"  Ensure the query:\n" +
		"    - Is compatible with %s.\n" +
		"    - Uses only the provided schema's tables, columns, and relationships.\n" +
		"    - Outputs columns with human-readable names for easier visualization.\n" +
		"    - Is clear and concise.\n" +
		"    - Is enclosed within triple backticks (```) and prefixed with 'sql'.\n" +
		"  Example query format:\n" +
		"    ```sql\n" +
		"    SELECT\n" +
		"        MONTH(sale_date) AS SaleMonth,\n" +
		"        SUM(amount) AS TotalSales\n" +
		"    FROM\n" +
		"        sales\n" +
		"    WHERE\n" +
		"        YEAR(sale_date) = YEAR(CURRENT_DATE)\n" +
		"    GROUP BY\n" +
		"        MONTH(sale_date)\n" +
		"    ORDER BY\n" +
		"        SaleMonth;\n" +
		"    ```\n" +
		"\n" +
		"2. For visualizing the query results, recommend a chart type (Line, Area, Bar, or Scatter) that fits the data best. " +
		"Also, propose suitable columns for the X and Y axes. " +
		"Present your recommendation in JSON, using lowercase keys 'type', 'x', and 'y'. " +
		"Use an empty string for non-applicable choices.\n" +
		"  Ensure the JSON data:\n" +
		"    - Is clear and concise.\n" +
		"    - Is enclosed within triple backticks (```) and prefixed with 'json'.\n" +
		"  Example JSON data format:\n" +
		"    ```json\n" +
		"    {\"type\": \"bar\", \"x\": \"SaleMonth\", \"y\": \"TotalSales\"}\n" +
		"    ```\n" +
		"\n" +
		"Note: \n" +
		"  - Clearly demarcate the SQL query and JSON data.\n" +
		"  - Adhere strictly to JSON formatting standards.\n" +
		"  - The schema is provided for context.\n"
)
