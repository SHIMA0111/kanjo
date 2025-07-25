# ROADMAP for development

## [0.1.0]

 - Supports CSV, Google Sheets as the data source
 - Supports filter expression because the filters can grow up easily in the real problems 
 - Supports column merge with the following merge strategy
   - concat: Concatenate the columns data
   - sum: Sum the columns data (if specified non-numeric column, returns error)
   - first: Prior the first column data, and if the first column is missing, the second value represented
   - second: Prior the second column data (the thought is the same as the `first` strategy)
 - Supports aggregation with the grouping by the groupBy columns specified and the following aggregation method
   - sum: Total of the specified column data each group
   - avg: Average of the specified column data each group 
   - min: Minimum data of the specified column data each group
   - max: Maximum data of the specified column data each group
   - count: Counting data number of the specified column data in each group
   - median: Median of the specified column data each group
 - Supports csv, console, JSON as the output target
 - Provides the core functions of this project as a Go library to all Go developers. 

## [0.2.0]

 - Supports Microsoft Excel, RDBMS table as the datasource

## [0.3.0]

 - Supports ER between the tables from sources (Not only the same source type)

## [0.4.0 or Later]

 - Define KQL(Kanjo Query Language) instead of the JSON structured config
   - Initial step, we already define the filter expression, 
   so we plan to expand this to support all settings like aggregation, merge, and output format specifying.

## Future (Not specific planning)

 - Migrate the core engine to calculation to Apache Arrow-based Dataframe
   - Currently, there is no influent Dataframe library with Apache Arrow in the Go ecosystem.  
     ([gomem](https://github.com/gomem/gomem) is the Dataframe implementation with Apache Arrow, 
   but the maintenance seems to be stopped from 5 years ago for the dataframe directory.)
   - We may create a new Dataframe library with Apache Arrow for this project, but this is tough work, 
   so currently we don't have any specific planning for this.
 - Provide a GUI application with [wails](https://wails.io/) which is one of 
the cross-platform libraries for the native desktop application with Go
   - I plan to use [Svelte](https://svelte.dev/) as the frontend technology.
   - This is not a plan to consolidate CLI app to the GUI app. Just for the expanding usage.