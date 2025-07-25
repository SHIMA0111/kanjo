# ROADMAP for development

## [0.1.0]

 - Supports CSV, Google Sheets

## [0.2.0]

 - Supports Microsoft Excel, RDBMS table

## [0.3.0]

 - Supports ER between the tables from sources (Not only the same source type)

## Future (Not specific planning)

 - Migrate the core engine to calculation to Apache Arrow-based Dataframe
   - Currently, there is no influent Dataframe library with Apache Arrow in the Go ecosystem.  
     ([gomem](https://github.com/gomem/gomem) is the Dataframe implementation with Apache Arrow, 
   but the maintenance seems to be stopped from 5 years ago for the dataframe directory.)
   - We may create a new Dataframe library with Apache Arrow for this project, but this is very hard work, 
   so currently we don't have any specific planning for this.