# Simple database mirror

> The goal here is to provide a low-cost simple database mirror from Oracle databases (and SQL-Server)
> to Postgres and Mongodb within a specific use-case that one of my clients is facing in (PoC).  
> In essence, the problem that needs to be tackled is to
> avoid launching reporting queries against the production database.
> So, all reporting queries can be executed over the mirrored databases
> without creating problems for the users in production time.
>
## IMPORTANT:
> During the mirroring the routines executes some ETL that are specific for databases used for IwCare systems 

## It isn't production-ready
> It uses actual database connections to execute all the processes. 
> It must be executed in time periods where the users actually aren't using the system. (Generally during the night)
> For production cases I recommend using Kafka-connect (or something like it)
> that offers connectors that execute all the processes using database-transactions-log
> to accomplish the same goal without creates any database degradation in production-time and as well allows the replication of data to occur synchronously as soon as each transaction takes place
> 
