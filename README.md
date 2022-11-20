# PNC-DB
Project Neural Cloud - Database

## Overiew
This repository is intended to be a database for the PNC discord bot, HK416. We use YAML for its readability, and as our database is small, there is little impacts on performance.

To help us, fork the repo and directly work on the YAML files. You need [go](https://go.dev/) installed to run go code to update meta.yaml.

After installation, you can run `go run .` in the project directory to generates a new meta.yaml file. Finally, open a PR so we can review your contribution and merge them, if approved.

## File Structure
- Asset folder contains icon thumbnails for dolls and alorithms.
    - Use PNG format.

- Data directory contains YAML data files to all relevant data for Neural Cloud. Stucture as follows:

### Doll file format:
#### Required
```
name: string
bio:
    rarity: byte
    class: string
    model: string
    manufacturer: string
    career: string
    birthday: string
    release: date (ISO8601)
    voice: string
skills:
    passive:
        name: string
        desc: string
    auto:
        name: string
        desc: string
    ultimate:
        name: string
        desc: string
```
#### Optional
```
algorithm:
    set: string
    main: string
    sub: string
    image: string

analysis:
    rating: string
    detail: string
links:
    wiki: string
    video: string
```
