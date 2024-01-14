![01-edu-system-blue](https://raw.githubusercontent.com/GoArtyom/study/1a66b22c5b511ccce94b582481a45dfd7f001d3a/alem.svg)

# Project `groupie-tracker`

![01-edu-system-blue](./web/static/img/preview.gif)

## Description

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

It will be given an <a href="https://groupietrackers.herokuapp.com/api" target="_blank">API</a>, that consists in four parts:

The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

The second one, `locations`, consists in their last and/or upcoming concert locations.

The third one, `dates`, consists in their last and/or upcoming concert dates.

And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

Given all this you should build a user friendly website where you can display the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc). It is up to you to decide how you will display it.

## Usage

### How to install

```bash
git clone git@github.com:ArtEmerged/groupie-tracker.git && cd groupie-tracker/
```

### How to go run

```bash
go run cmd/main.go
```

and click on http://localhost:8081/ (ctrl + click)

## Autors
Follow us on:
* <a href="https://github.com/ArtEmerged" target="_blank">GitHub.com/ArtEmerged</a>
* <a href="https://github.com/NordStream777" target="_blank">GitHub.com/NordStream777</a>
* <a href="https://github.com/grenkoff" target="_blank">GitHub.com/grenkoff</a>