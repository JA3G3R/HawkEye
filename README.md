# HawkEye
An extensive and profitable reconnaissance tool
=======
# Documentation of HawkEye

## Overview

HawkEye is an extensive recon-framework designed by keeping in mind simplicity, efficiency and the one-click-does-all mindset.


## Functionaliy

### Adding Target to the database

Since HawkEye is all about reconnaissance at scale, it is imperative to have an efficient system in place to track all the targets. For this reason the tool includes two scripts to add targets manually or automatically. 

**MANUAL**

To add the targets manually, use the script `newtarget` with the following cmdline arguments:

- --name: The name of the new target
- --platform: Bug Bounty platform on which the bug bounty program of this new target is hosted.
- --root-domains: A comma separated list of root domains included in the scope of the target.
- --excluded-domains: A comma separated list of root domains related to the target to exclude.

