A simple reproduction of the golang antlr performance issue

A simple rule such as:

```
1 EQ 2 OR
1 EQ 2 OR
1 EQ 2 OR
1 EQ 2 OR
1 EQ 2
```

takes exponentially longer to parse the more `1 EQ 2 OR` clauses there are. This does not happen in python.

On my machine, # of lines vs parse time:

```
11: 0.5s
12: 1.2s
13: 3.2s
14: 8.1s
15: 21.9s
16: 57.5s
```

Antlr generation run in evaluate/simpleboolean:
`java -jar antlr-4.10.1-complete.jar -visitor -Dlanguage=Go SimpleBoolean.g4 -package simpleboolean`
