
Like uniq, but for logs.

# Install

```
go install github.com/psykhi/uno@latest
```

Or download from the [releases page](https://github.com/psykhi/uno/releases).

# Usage

```bash
cat my_log_file.txt | uno
uno my_log_file.txt
```

`uno` behaves like `uniq` and writes unique lines to the standard output. Unlike `uniq`, `uno` uses fuzzy matching
to determine if two lines are similar or not. 

# Options

To see all options, use`uno -h`

### -d

The distance option `-d` can be used to specify how different a new line must be from the others we've seen to be deemed
new/unique. It can take a value between `0` and `1`. The default is `0.3` (30% difference)

### -all

To see all input lines and highlight the new ones in red, use `-all`

```bash
cat my_log_file.txt | uno -all
uno -all my_log_file.txt
```
