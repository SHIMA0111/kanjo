# Kanjo: Fetch and Filter and Aggregate the data in Spreadsheets
Kanjo fetches and transforms data (filter, aggregation, and column merge) 
from spreadsheets like CSV, Google Sheets, Excel (planned in 0.2) easily. 

## Kanjo
Kanjo means `勘定(方)` in Japanese refers to officials in historical Japan 
who were responsible for financial affairs, similar to 
what we call accountants or treasurers today.  
**Kanjo** supports you to transform and aggregate your data like the `勘定(方)`.

## Requirements

 - Go 1.23.0+ (toolchain 1.23.2)
 - Google OAuth2 credentials (If you use Google Sheets)

## Limitation

 - Supports only one sheet per operation (Plan to support ER after 0.3)