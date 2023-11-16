# Project `groupie-tracker`

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

```
git clone git@git.01.alem.school:aecheist/groupie-tracker.git && cd groupie-tracker/
```

### How to go run

```
go run cmd/api/main.go
```

and click on http://localhost:8080/ (ctrl + click)

### Audit link

* <a href="https://github.com/01-edu/public/tree/master/subjects/groupie-tracker/audit" target="_blank">groupie-tracker audit</a>

## Autors
Follow us on:
* <a href="https://github.com/GoArtyom" target="_blank">GitHub.com/GoArtyom</a>
* <a href="https://github.com/grenkoff" target="_blank">GitHub.com/grenkoff</a>