# Lazysh

Ever opened your terminal and thought your shell starts slow?

Well, that probably happens because of the enormous amounts of version managers, etc. which are loaded in your *rc-file.

The solution: lazy load that stuff!!1!

> [!NOTE]
> Currently only works with [bash](https://www.gnu.org/software/bash/) and [fish](https://fishshell.com/). On support for other shells is being worked on.

## Installation

### From source

To install the `lazysh`-binary on your system, simply run:

~~~sh
sudo make clean install
~~~

## Getting started

Now, in your *rc-file (`.config/fish/config.fish`, `.zshrc`, `.bashrc`, ...), you have to paste a small snippet.

### Bash

~~~sh
source $(echo '
<init commands>
' | lazysh bash)
~~~

### Fish

~~~sh
source $(echo '
<init commands>
' | lazysh fish)
~~~

and replace `<init commands>` with your init commands, like `rbenv init - | source`, each placed on a new line.

It could look like:

~~~sh
source $(echo '
eval "$(zoxide init bash)"
eval "$(rbenv init - bash)"
' | lazysh bash)
~~~

(or a fish equivalent):
~~~sh
source $(echo '
zoxide init fish | source
rbenv init - | source
' | lazysh fish)
~~~

> [!WARNING]
> It should be common sense to **only** paste commands you trust in there.

That's it!
