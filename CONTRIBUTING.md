## Vendor requirements

The following requirements have to be fulfilled in order to for a vendor website to be qualified for being part of Keebmap:

1. Vendor should be focused on the custom mechanical keyboard hobby, *i.e not a generic electronics store*
2. Website **should not** only be focused on particular group buy run with limited availability
3. Vendor should carry at least one in stock item related to the mechanical keyboard hobby

The goal of these requirements is to keep the list useful and make it so the users don't have to click through a bunch of irrelevant links.

## Using the command line tool

In order to facilitate the modification of the list of vendors, you can use the CLI tool via the following command: `npm run edit`

ℹ️ You must have Go installed (preferably in version 1.16), [*See the download page*](https://go.dev/dl/).

The tool will process the data, apply the alphabetical sorting, and apply the changes to the listing (its use is therefore optional).
It's possible to disable a vendor instead of deleting it, this can be useful in case the vendor seems not to be active anymore but has not given any information.

## Run development environment

`npm install` and `npm run start`, that is it! You should be up and running.
