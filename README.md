## A Simple ToDo App built with a-h/templ, Go-Chi, Sqlite3 and Htmx

I was curious about templating using a-h/templ and htmx (because of all the hype obviously).

This is a CRUD app without the **"U"** mainly because I got bored.

To run the app on your own machine machine you must have **sqlite3** installed in your machine, and create the necessary table. Make sure to run the following command

```bash
sqlite3 test.db < table.sql
```
### You Want To Edit the files
Make sure to follow instructions in the `a-h/templ` in [templ.guide](https://templ.guide/quick-start/installation) You can use the `watcher.sh` to run `templ generate`, every time there's a change *.templ* file.

```bash
chmod +x watcher.sh
```

then run (I use **linux**, it maybe different on your device).
```bash
./watcher.sh
```
