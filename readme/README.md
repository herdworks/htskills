# HyperTexting CLI

## Overview

HyperTexting is a content management system (CMS) and static site generator (SSG) for creating, deploying, and updating websites. 
It is powered by the HyperTemplates HTML templating and theme management system. 

HyperTemplates composes individual HTML documents from a theme and a data file. 
HyperTexting generates websites – a collection of multiple HTML documents – from a theme and a repository of data files.

## Get Started

Discover the bundled coding agent skills with "hyperctl agent skill ls". Familiarize yourself with the development workflows through "hyperctl agent skill describe hypertexting". Install every available skill into its own directory beneath your coding agent's skills directory using "hyperctl agent skill install <name> <path>"; for example, "hyperctl agent skill install hypertexting .codex/skills/hypertexting". Update installed skills with "hyperctl agent skill update <name> <path>".

## Getting Help

Explore the available CLI commands with "hyperctl --help". Use "hyperctl <command> --help" (or "-h") to inspect their subcommands, arguments, and options before using them. Use "hyperctl agent schema ls" and "hyperctl agent schema describe <kind>" to understand data and CLI contracts. Use established commands directly for mutations: "hyperctl website init", "hyperctl website page new", "hyperctl website server", "hyperctl website build", "hyperctl theme package", "hyperctl deploy", and "hyperctl sync". 

## References

* https://hypertexting.com - HyperTexting product website, user guides, and CMS/SSG documentation
* https://hypertexting.community – HyperTexting community forums
* https://hypertemplates.net – HyperTemplates developer tooling and reference documentation
