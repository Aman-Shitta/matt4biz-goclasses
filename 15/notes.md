# GOPL Exercis e 4.12: 

The popular web comic xkcd has a JSON interface. For example, a request to
https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
that, using this index, prints the URL and transcript of each comic that matches a search term
provided on the command line.

# RAW data

{
  "month": "4",
  "num": 572,
  "link": "",
  "year": "2009",
  "news": "",
  "safe_title": "Together",
  "transcript": "[[A male and female are running in a field, holding hands. Another male and female stand in the background, next to a table.]]\n[[The man and woman are in a boat on a lake, very romantic. The man is speaking to the woman, illustrated with a heart.]]\n[[The man and woman sit together on a bench on a beach, watching the sunset.]]\n[[The man and woman stand at an altar. They have married.]]\n[[The man and woman, having grown old together, sit together on their doorstep, holding hands.]]\n[[The man begins walking away with his cane.]]\nWoman: Dear? Where are you--Come back!\n[[The man approaches the other couple from the first panel, who are now just as old.]]\n[[The man picks up a piece of paper from the table in the first panel and begins to write.]]\nMan: Okay,\n[[The paper is shown: a scavenger hunt list. \"Happiness\" has just been checked off.]]\nMan: What's next?\n[[Full list:\nSCAVENGER HUNT:\n[X] Indian-head penny\n[X] Snake skin\n[X] Happiness\n[  ] Four-leaf clover\n[  ] Shark tooth\n[...]\n]]\n{{Alt text: This scavenger hunt is getting boring. Let's go work on the treehouse!}}",
  "alt": "This scavenger hunt is getting boring.  Let's go work on the treehouse!",
  "img": "https://imgs.xkcd.com/comics/together.png",
  "title": "Together",
  "day": "22"
}


# Loader program

- Read until we get 2 404 responses in a row
- Each request will generate a JSON object as a string.
- We will bracket them with [ and ] and a comma in between.
- The result will be a file with JSOn array and metedata objects, so we won't need to decode
- Optinally take a filename from commmand line for output.

# Searcher Program

- require a dilename from CLI for input
- read in an decode slices of object
- take in mutliple search terms from cli
- we will slect comics that match all words by doing a quadratic search (nexted loops)
- we will compare the string in lower case.