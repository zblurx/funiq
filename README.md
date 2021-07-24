# funiq

Make a uniq on a file without f*cking up with its content

```bash
$ wc -l userlist 
99 userlist
$ funiq userlist 
$ wc -l userlist
30 userlist
```