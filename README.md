# iTrak Machine Maintenance

System for managing maintenance activities on machines in the manufacturing plant.


# Architecture

## SQL Database

- 
- Holds stored procedures / business logic for interfacing with that data

## WebApp

- Single Page App, with login validation
- Roles for Machine Workers
  - View restricted to own machines
  - Raise stoppage event
  - View machine manuals
  - View machine reports (work orders, photos, notes)
- Roles for Maintenance Manager 
  - Receive stoppage event
  - Create work orders for stoppage event
  - Allocate work orders to maintenance worker
  - Manually escalate a work order
  - Run maintenance reports
    - Outstanding work orders
    - Preventative vs Breakdown over a period
    - Spare parts used over a period
    - Spare parts inventory report
- Roles for Maintenance Workers
  - View restricted to machine related to work orders
  - Receive work order
  - Acknowledge work order
  - View machine manuals
  - View machine reports
  - Attach notes / photos to work order
  - Close off work order
- Roles for Stock Control Workers
  - View restricted to own machines
  - View spare parts inventory
  - Receive spare parts into inventory
  - Remove spare parts from inventory
  - Stock Take report + physical stocktake reconcilliation

## Server

- Lightweight server (golang) to provide REST services to the SPA
- Runs a watcher task to generate escalation events
- Uses eternal service to broadcast SMS alerts